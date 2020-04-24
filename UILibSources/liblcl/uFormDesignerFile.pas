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

{$I ExtDecl.inc}

interface

uses
  Forms,
  sysutils,
  Classes,
  zbase,
  typinfo,
  paszlib,
  Types;

type
  EZLibError = class(Exception);
  EZCompressionError = class(EZLibError);
  EZDecompressionError = class(EZLibError);

  { TFormResFile }

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
    class procedure ZDecompressStream(inStream, outStream: TStream);
    class function ZDecompressCheck(code: Integer): Integer;
    class function ZDecompressCheckWithoutBufferError(code: Integer): Integer;
  end;

  procedure ResFormLoadFromResourceName(AInstance: NativeUInt; AResName: PChar; ARoot: TComponent); extdecl;
  procedure ResFormLoadFromFile(AFileName: PChar; ARoot: TComponent); extdecl;
  procedure ResFormLoadFromStream(AStream: TStream; ARoot: TComponent); extdecl;

implementation

uses
  uComponents;


const
  {** return code messages **************************************************}
  _z_errmsg: array [0..9] of PAnsiChar = (
    'need dictionary',      // Z_NEED_DICT      (2)  //do not localize
    'stream end',           // Z_STREAM_END     (1)  //do not localize
    '',                     // Z_OK             (0)  //do not localize
    'file error',           // Z_ERRNO          (-1) //do not localize
    'stream error',         // Z_STREAM_ERROR   (-2) //do not localize
    'data error',           // Z_DATA_ERROR     (-3) //do not localize
    'insufficient memory',  // Z_MEM_ERROR      (-4) //do not localize
    'buffer error',         // Z_BUF_ERROR      (-5) //do not localize
    'incompatible version', // Z_VERSION_ERROR  (-6) //do not localize
    ''                                               //do not localize
    );

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
          AInStream.Read(LBytes[0], Length(LBytes));
          LMem.Write(LBytes[0], Length(LBytes));
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

class procedure TFormResFile.ZDecompressStream(inStream, outStream: TStream);
const
  bufferSize = 32768;
var
  zstream: TZStream;
  zresult: Integer;
  inBuffer: TBytes;
  outBuffer: TBytes;
  inSize: Integer;
  outSize: Integer;
begin
  SetLength(inBuffer, BufferSize);
  SetLength(outBuffer, BufferSize);
  FillChar(zstream, SizeOf(TZStream), 0);
  ZDecompressCheck(InflateInit(zstream));

  try
    inSize := inStream.Read(inBuffer[0], bufferSize);
    while inSize > 0 do
    begin
      zstream.next_in := @inBuffer[0];
      zstream.avail_in := inSize;

      repeat
        zstream.next_out := @outBuffer[0];
        zstream.avail_out := bufferSize;

        ZDecompressCheckWithoutBufferError(inflate(zstream, Z_NO_FLUSH));

        outSize := bufferSize - zstream.avail_out;

        outStream.Write(outBuffer[0], outSize);
      until (zstream.avail_in = 0) and (zstream.avail_out > 0);

      inSize := inStream.Read(inBuffer[0], bufferSize);
    end;

    repeat
      zstream.next_out := @outBuffer[0];
      zstream.avail_out := bufferSize;

      zresult := ZDecompressCheckWithoutBufferError(inflate(zstream, Z_FINISH));

      outSize := bufferSize - zstream.avail_out;

      outStream.Write(outBuffer[0], outSize);
    until (zresult = Z_STREAM_END) and (zstream.avail_out > 0);
  finally
    ZDecompressCheck(inflateEnd(zstream));
  end;
end;

class function TFormResFile.ZDecompressCheck(code: Integer): Integer;
begin
  Result := code;
  if code < 0 then
    raise EZDecompressionError.Create(string(_z_errmsg[2 - code]));
end;

class function TFormResFile.ZDecompressCheckWithoutBufferError(code: Integer): Integer;
begin
  Result := code;
  if code < 0 then
    if (code <> Z_BUF_ERROR) then
     raise EZDecompressionError.Create(string(_z_errmsg[2 - code]));
end;

//------------------------------------------------------------------------------


type

  { TResForm }

  TResForm = class
  public
    procedure OnFindComponentClass(AReader: TReader; const AClassName: string;
       var AComponentClass: TComponentClass);
    procedure OnPropertyNotFound(AReader: TReader; AInstance: TPersistent;
       var APropName: string; AIsPath: boolean; var AHandled, ASkip: Boolean);

    procedure OnAncestorNotFound(Reader: TReader; const ComponentName: string;
      ComponentClass: TPersistentClass; var Component: TComponent);
    procedure OnReaderError(Reader: TReader; const Message: string;
      var Handled: Boolean);
    procedure OnCreateComponentEvent(Reader: TReader;
      ComponentClass: TComponentClass; var Component: TComponent);

    procedure OnReadWriteStringProperty(Sender:TObject;
      const Instance: TPersistent; PropInfo: PPropInfo;
      var Content:string);

    procedure LoadFromFile(const AFileName: string; ARoot: TComponent);
    procedure LoadFromStream(AStream: TStream; ARoot: TComponent);
    procedure LoadFromResourceName(AInstance: NativeUInt; AName: string; ARoot: TComponent);
  end;


procedure TResForm.OnFindComponentClass(AReader: TReader; const AClassName: string;
  var AComponentClass: TComponentClass);
var
  LB: TClass;
begin
  if uClassLists.TryGetData(AClassName, LB) then
    AComponentClass := TComponentClass(LB);
end;

procedure TResForm.OnPropertyNotFound(AReader: TReader; AInstance: TPersistent;
  var APropName: string; AIsPath: boolean; var AHandled, ASkip: Boolean);
begin
   ASkip := True;
end;

procedure TResForm.OnAncestorNotFound(Reader: TReader;
  const ComponentName: string; ComponentClass: TPersistentClass;
  var Component: TComponent);
begin

end;

procedure TResForm.OnReaderError(Reader: TReader; const Message: string;
  var Handled: Boolean);
begin
  Handled := True;
end;

procedure TResForm.OnCreateComponentEvent(Reader: TReader;
  ComponentClass: TComponentClass; var Component: TComponent);
begin
//{$IFDEF WINDOWS}
//  if ComponentClass.ClassName = 'TStatusBar' then
//    Component := TComponent.Create(Reader.Owner);
//{$ENDIF}
end;

procedure TResForm.OnReadWriteStringProperty(Sender: TObject;
  const Instance: TPersistent; PropInfo: PPropInfo; var Content: string);
begin

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

function InitLazResourceComponent(AReader: TReader; Instance: TComponent; RootAncestor: TClass): Boolean;

  function InitComponent(ClassType: TClass): Boolean;
  begin
    Result := False;
    if (ClassType = TComponent) or (ClassType = RootAncestor) then
      Exit;
    if Assigned(ClassType.ClassParent) then
      Result := InitComponent(ClassType.ClassParent);
    AReader.ReadRootComponent(Instance);
    Result := True;
  end;

begin
  if Instance.ComponentState * [csLoading, csInline] <> []
  then begin
    // global loading not needed
    Result := InitComponent(Instance.ClassType);
  end
  else try
    BeginGlobalLoading;
    Result := InitComponent(Instance.ClassType);
    NotifyGlobalLoading;
  finally
    EndGlobalLoading;
  end;
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
    LMem.Position := 0;
    LR := TReader.Create(LMem, 4096);
    try
      LR.OnFindComponentClass := @OnFindComponentClass;
      LR.OnPropertyNotFound := @OnPropertyNotFound;
      LR.OnAncestorNotFound := @OnAncestorNotFound;
      LR.OnError:= @OnReaderError;
      LR.OnReadStringProperty := @OnReadWriteStringProperty;
      LR.OnCreateComponent := @OnCreateComponentEvent;

      GlobalNameSpace.BeginWrite;
      try
        if (ARoot.ClassType <> TForm) and not (csDesigning in ARoot.ComponentState) then
        begin
          if ARoot is TForm then
          begin
            TFormPatch(ARoot).FormStateIncludeCreating;
            try
              InitLazResourceComponent(LR, ARoot, TForm);
            finally
              TFormPatch(ARoot).FormStateExcludeCreating;
            end;
          end else if ARoot is TFrame then
          begin
            //ControlStyle := [csAcceptsControls, csCaptureMouse, csClickEvents, csSetCaption,
            //                 csDoubleClicks, csParentBackground];
            if (ARoot.ClassType = TFrame) and ([csDesignInstance, csDesigning] * ARoot.ComponentState = []) then
            begin
              LR.ReadRootComponent(ARoot);
              // InitLazResourceComponent(LR, ARoot, TFrame);
            end
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

procedure ResFormLoadFromStream(AStream: TStream; ARoot: TComponent); extdecl;
var
  LObj: TResForm;
begin
  LObj := TResForm.Create;
  try
    try
      LObj.LoadFromStream(AStream, ARoot);
    except
      on E: Exception do
        Writeln(E.Message);
    end;
  finally
    LObj.Free;
  end;
end;

procedure ResFormLoadFromFile(AFileName: PChar; ARoot: TComponent); extdecl;
var
  LObj: TResForm;
begin
  LObj := TResForm.Create;
  try
    try
      LObj.LoadFromFile(AFileName, ARoot);
    except
      on E: Exception do
        Writeln(E.Message);
    end;
  finally
    LObj.Free;
  end;
end;

procedure ResFormLoadFromResourceName(AInstance: NativeUInt; AResName: PChar; ARoot: TComponent); extdecl;
var
  LObj: TResForm;
begin
  LObj := TResForm.Create;
  try
    try
      LObj.LoadFromResourceName(AInstance, AResName, ARoot);
    except
      on E: Exception do
        Writeln(E.Message);
    end;
  finally
    LObj.Free;
  end;
end;


//exports
//  ResFormLoadFromStream,
//  ResFormLoadFromFile,
//  ResFormLoadFromResourceName;
//
end.
