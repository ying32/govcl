//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

// 设计器gfm文件格式，这里是加密格式，至于为啥使用加密的，一
// 开始是做了外部gfm文件加载，防止项目发布使用了外部文件随意被修改，所以没有公开相关的


unit uFormDesignerFile;

interface

uses
  System.SysUtils,
  System.Classes,
  System.Zlib,
  Vcl.Forms;

type

  TFormResFile = class
  public const
    VERSION = 1;
    // GOVCLFORMZ
    HEADER: array[0..9] of byte = ($47, $4F, $56, $43, $4C, $46, $4F, $52, $4D, $5A);
    // TPF0
    HEADERTPF0: array[0..3] of byte = ($54, $50, $46, $30);
  public
    class function Decrypt(AInStream: TStream; AOutStream: TStream): Boolean;
    class procedure XorStream(AStream: TMemoryStream);
  end;

implementation

uses
  Winapi.Windows,
  uComponents;

{ TFormResFile }

class function TFormResFile.Decrypt(AInStream: TStream; AOutStream: TStream): Boolean;
var
  LVer: Word;
  LMem, LZipMem: TMemoryStream;
  LHeader:array[0..High(HEADER)] of Byte;
  LBytes: TBytes;
begin
  Result := False;
  AInStream.Position := 0;
  if AInStream.Size > Length(HEADER) + SizeOf(Word) then
  begin
    AInStream.Read(LHeader[0], Length(LHeader));
    if CompareMem(@LHeader[0], @HEADER[0], Length(HEADER)) then
    begin
      AInStream.Read(LVer, SizeOf(Word));
      if LVer >= VERSION then
      begin
        LMem := TMemoryStream.Create;
        try
          SetLength(LBytes, AInStream.Size - Length(HEADER) - SizeOf(Word));
          AInStream.Read(LBytes, 0, Length(LBytes));
          LMem.Write(LBytes, 0, Length(LBytes));
          LMem.Position := 0;
          XorStream(LMem);
          LMem.Position := 0;
          LZipMem := TMemoryStream.Create;
          try
            ZDecompressStream(LMem, LZipMem);
            AOutStream.Position := 0;
            AOutStream.Write(LZipMem.Memory^, LZipMem.Size);
            AOutStream.Position := 0;
            Result := True;
          finally
            LZipMem.Free;
          end;
        finally
          LMem.Free;
        end;
      end;
    end else
    // 原生格式
    if CompareMem(@LHeader[0], @HEADERTPF0[0], Length(HEADERTPF0)) then
    begin
      AInStream.Position := 0;
      AOutStream.Position := 0;
      AOutStream.CopyFrom(AInStream, AInStream.Size);
      AOutStream.Position := 0;
      Result := True;
    end;
  end;
end;

class procedure TFormResFile.XorStream(AStream: TMemoryStream);
var
  EncodedBytes, ZipBytes:array of Byte; //编码字节
  BytesLength: Integer;
  I, L:integer;
const
  EncKey: array[0..15] of Byte = (
	$EB, $D4, $93, $48, $97, $0E, $38, $5B, $89, $B8, $2F, $D2, $31, $11, $AD, $BC);

begin
  AStream.Position := 0;
  BytesLength := AStream.Size;
  SetLength(EncodedBytes, BytesLength);
  SetLength(ZipBytes, BytesLength);
  AStream.Read(EncodedBytes[0], BytesLength);
  for I := 0 to BytesLength - 1 do
  begin
    L := I mod 16;
    ZipBytes[i] := EncodedBytes[i] xor EncKey[L];
  end;
  AStream.Clear; //清除以前的流
  AStream.Write(ZipBytes[0], BytesLength); //将异或后的字节集写入原流中
  AStream.Position := 0;
end;

type
  TResForm = class
  public
    procedure OnFindComponentClass(Reader: TReader; const AClassName: string; var ComponentClass: TComponentClass);
    procedure OnAncestorNotFound(Reader: TReader; const ComponentName: string; ComponentClass: TPersistentClass; var Component: TComponent);
    procedure OnReaderError(Reader: TReader; const Message: string; var Handled: Boolean);
    procedure OnFindMethod(Reader: TReader; const MethodName: string; var Address: Pointer; var Error: Boolean);
    procedure OnFindMethodInstance(Reader: TReader; const MethodName: string; var AMethod: TMethod; var Error: Boolean);

    procedure LoadFromFile(const AFileName: string; ARoot: TComponent);
    procedure LoadFromStream(AStream: TStream; ARoot: TComponent);
    procedure LoadFromResourceName(AInstance: NativeUInt; AName: string; ARoot: TComponent);
  end;


procedure TResForm.OnFindComponentClass(Reader: TReader; const AClassName: string;
  var ComponentClass: TComponentClass);
var
  LB: TClass;
begin
  if uClassLists.TryGetValue(AClassName, LB) then
    ComponentClass := TComponentClass(LB);
end;

procedure TResForm.OnAncestorNotFound(Reader: TReader; const ComponentName: string;
   ComponentClass: TPersistentClass; var Component: TComponent);
begin

end;

procedure TResForm.OnReaderError(Reader: TReader; const Message: string; var Handled: Boolean);
begin
{$IFDEF DEBUG}
  OutputDebugString(PChar('ReaderError: ' + Message));
{$ENDIF}
  Handled := True; // 跳过不显示错误。
end;

procedure TResForm.OnFindMethod(Reader: TReader; const MethodName: string; var Address: Pointer; var Error: Boolean);
begin
{$IFDEF DEBUG}
  OutputDebugString(PChar('FindMethod: ' + MethodName));
{$ENDIF}
  Error := False;
end;

procedure TResForm.OnFindMethodInstance(Reader: TReader; const MethodName: string; var AMethod: TMethod; var Error: Boolean);
begin
{$IFDEF DEBUG}
  OutputDebugString(PChar('FindMethodInstance: ' + MethodName));
{$ENDIF}
  Error := False;
end;




function InitInheritedComponent(AReader: TReader; Instance: TComponent; RootAncestor: TClass): Boolean;

  function InitComponent(ClassType: TClass): Boolean;
  begin
    Result := False;
    if (ClassType = TComponent) or (ClassType = RootAncestor) then Exit;
    Result := InitComponent(ClassType.ClassParent);
    AReader.ReadRootComponent(Instance);
  end;

var
  LocalizeLoading: Boolean;
begin
  GlobalNameSpace.BeginWrite;  // hold lock across all ancestor loads (performance)
  try
    LocalizeLoading := (Instance.ComponentState * [csInline, csLoading]) = [];
    if LocalizeLoading then BeginGlobalLoading;  // push new loadlist onto stack
    try
      Result := InitComponent(Instance.ClassType);
      //if Result then Instance.ReadDeltaState;
      if LocalizeLoading then NotifyGlobalLoading;  // call Loaded
    finally
      if LocalizeLoading then EndGlobalLoading;  // pop loadlist off stack
    end;
  finally
    GlobalNameSpace.EndWrite;
  end;
end;

type
   TFormPatch = class(TForm)
   public
     procedure FormStateIncludeCreating;
     procedure FormStateExcludeCreating;
   end;

procedure TFormPatch.FormStateIncludeCreating;
begin
  Include(FFormState, fsCreating);
end;

procedure TFormPatch.FormStateExcludeCreating;
begin
  Exclude(FFormState, fsCreating);
end;

procedure TResForm.LoadFromStream(AStream: TStream; ARoot: TComponent);
var
  LR: TReader;
  LMem: TMemoryStream;
begin
  if (AStream = nil) or (ARoot = nil) then
    Exit;
  AStream.Position := 0;
  LMem := TMemoryStream.Create;
  try
    TFormResFile.Decrypt(AStream, LMem);
    LR := TReader.Create(LMem, 4096);
    try
      LR.OnFindComponentClass := OnFindComponentClass;
      LR.OnError := OnReaderError;
      LR.OnFindMethod := OnFindMethod;
      LR.OnAncestorNotFound := OnAncestorNotFound;
      LR.OnFindMethodInstance := OnFindMethodInstance;

      GlobalNameSpace.BeginWrite;
      try
        if (ARoot.ClassType <> TForm) and not (csDesigning in ARoot.ComponentState) then
        begin
          TFormPatch(ARoot).FormStateIncludeCreating;
          try
            // 这里换地方处理吧
             InitInheritedComponent(LR, ARoot, TForm);
          finally
            TFormPatch(ARoot).FormStateExcludeCreating;
          end;
        end;
      finally
        GlobalNameSpace.EndWrite;
      end;
    finally
      LR.Free;
    end;
  finally
    LMem.Free;
  end;
end;

procedure TResForm.LoadFromFile(const AFileName: string; ARoot: TComponent);
var
  LFileStream: TFileStream;
begin
  LFileStream := TFileStream.Create(AFileName, fmOpenRead);
  try
    LoadFromStream(LFileStream, ARoot);
  finally
    LFileStream.Free;
  end;
end;

procedure TResForm.LoadFromResourceName(AInstance: NativeUInt; AName: string; ARoot: TComponent);
var
  LRes: TResourceStream;
begin
  LRes := TResourceStream.Create(AInstance, PChar(AName), RT_RCDATA);
  try
    LoadFromStream(LRes, ARoot);
  finally
    LRes.Free;
  end;
end;

//-----------------------------------------------------------------------------

procedure ResFormLoadFromStream(AStream: TStream; ARoot: TComponent); stdcall;
var
  LObj: TResForm;
begin
  LObj := TResForm.Create;
  try
    try
      LObj.LoadFromStream(AStream, ARoot);
    except
      on E: Exception do
        OutputDebugString(PChar(E.Message));
    end;
  finally
    LObj.Free;
  end;
end;

procedure ResFormLoadFromFile(AFileName: PChar; ARoot: TComponent); stdcall;
var
  LObj: TResForm;
begin
  LObj := TResForm.Create;
  try
    try
      LObj.LoadFromFile(AFileName, ARoot);
    except
      on E: Exception do
        OutputDebugString(PChar(E.Message));
    end;
  finally
    LObj.Free;
  end;
end;

procedure ResFormLoadFromResourceName(AInstance: NativeUInt; AResName: PChar; ARoot: TComponent); stdcall;
var
  LObj: TResForm;
begin
  LObj := TResForm.Create;
  try
    try
      LObj.LoadFromResourceName(AInstance, AResName, ARoot);
    except
      on E: Exception do
        OutputDebugString(PChar(E.Message));
    end;
  finally
    LObj.Free;
  end;
end;

exports
  ResFormLoadFromStream,
  ResFormLoadFromFile,
  ResFormLoadFromResourceName;


end.
