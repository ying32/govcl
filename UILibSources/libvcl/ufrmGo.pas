unit ufrmGo;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Winapi.ActiveX;

type
  TDropFilesEvent = procedure(Sender: TObject; const FileNames: array of String) of object;

  TWndProcEvent = procedure(Sender: TObject; var AMsg: TMessage; var AHandled: Boolean) of object;

  TGoForm = class(TForm)
  private
    FOnDropFiles: TDropFilesEvent;
    FOnStyleChanged: TNotifyEvent;
    FOnWndProc: TWndProcEvent;
    FAllowDropFiles: Boolean;
    procedure SetAllowDropFiles(const Value: Boolean);
    procedure WMDropFiles(var Msg: TWMDropFiles); message WM_DROPFILES;
    procedure CMStyleChanged(var Msg: TMessage); message CM_STYLECHANGED;
    function GetAllowDropFiles: Boolean;
  protected
    procedure InitializeNewForm; override;
  public
    constructor Create(AOwner: TComponent); override;
    procedure WndProc(var AMsg: TMessage); override;
  published
    property AllowDropFiles: Boolean read GetAllowDropFiles write SetAllowDropFiles;
    property OnDropFiles: TDropFilesEvent read FOnDropFiles write FOnDropFiles;
    property OnStyleChanged: TNotifyEvent read FOnStyleChanged write FOnStyleChanged;
    property OnWndProc: TWndProcEvent read FOnWndProc write FOnWndProc;
  end;

  procedure LockInitScale;
  procedure UnLockInitScale;
  procedure SetInitScale(AValue: Boolean);
implementation

//{$R *.dfm}

uses
  uComponents,
  Winapi.ShellAPI;

var
  uLockObj: TObject;
  uInitScale: Boolean;

procedure LockInitScale;
begin
  System.TMonitor.Enter(uLockObj);
end;

procedure UnLockInitScale;
begin
  System.TMonitor.Exit(uLockObj);
end;

procedure SetInitScale(AValue: Boolean);
begin
  uInitScale := AValue;
end;

procedure Form_ScaleForPPI(AObj: TGoForm; ANewPPI: Integer); stdcall;
begin
  AObj.ScaleForPPI(ANewPPI);
end;

procedure Form_ScaleControlsForDpi(AObj: TGoForm; ANewPPI: Integer); stdcall;
begin
  AObj.ScaleControlsForDpi(ANewPPI);
end;


{ TGoForm }

procedure TGoForm.CMStyleChanged(var Msg: TMessage);
begin
  inherited;
  // 修复样式造成的问题，如果设置了允许拖放，但实际窗口风格中已经不存在了，则
  // 重新设置。
  if FAllowDropFiles and not AllowDropFiles then
    AllowDropFiles := True;
  if Assigned(FOnStyleChanged) then
    FOnStyleChanged(Self);
end;

constructor TGoForm.Create(AOwner: TComponent);
var
  LPPI: Integer;
begin
  try
    // 这里需要屏蔽对资源查找的错误
    inherited Create(AOwner);
  except
  end;
  if OldCreateOrder then
    DoCreate;
//  Create(AOwner, 0);
  if uInitScale and GetGlobalFormScaled then
  begin
    LPPI := Screen.PixelsPerInch;
    ClientWidth := MulDiv(ClientWidth, LPPI, 96);
    ClientHeight := MulDiv(ClientHeight, LPPI, 96);
    ScaleForPPI(LPPI);
  end;
  ControlStyle := ControlStyle + [csPaintBlackOpaqueOnGlass];
end;

function TGoForm.GetAllowDropFiles: Boolean;
begin
  Result := (GetWindowLong(Handle, GWL_EXSTYLE) and WS_EX_ACCEPTFILES) <> 0;
end;

procedure TGoForm.InitializeNewForm;
begin
  inherited InitializeNewForm;
  Self.ClientHeight := 321;
  Self.ClientWidth := 678;
  if GetGlobalFormScaled then
    Self.PixelsPerInch := 96
  else Self.PixelsPerInch := Screen.PixelsPerInch;
  Scaled := False;
end;

procedure TGoForm.SetAllowDropFiles(const Value: Boolean);
begin
  FAllowDropFiles := Value;
  if AllowDropFiles <> Value then
    DragAcceptFiles(Handle, Value);
end;

procedure TGoForm.WMDropFiles(var Msg: TWMDropFiles);
var
  LCount, I: Cardinal;
  LFileName: array[0..MAX_PATH-1] of Char;
  LFileNames: array of string;
begin
  inherited;
  if Assigned(FOnDropFiles) then
  begin
    LCount := DragQueryFile(Msg.Drop, INFINITE, nil, 0);
    if LCount > 0 then
    begin
      SetLength(LFileNames, LCount);
      for I := 0 to LCount - 1 do
      begin
        FillChar(LFileName, SizeOf(LFileName), #0);
        DragQueryFile(Msg.Drop, I, LFileName, SizeOf(LFileName));
        LFileNames[I] := string(LFileName);
      end;
      FOnDropFiles(Self, LFileNames);
    end;
  end;
end;

procedure TGoForm.WndProc(var AMsg: TMessage);
var
  LHandled: Boolean;
begin
  LHandled := True;
  if Assigned(FOnWndProc) then
    FOnWndProc(Self, AMsg, LHandled);
  if LHandled then
    inherited;
end;

exports
   Form_ScaleForPPI,
   Form_ScaleControlsForDpi;


initialization
  uLockObj := TObject.Create;
  CoInitialize(nil);

finalization
  CoUninitialize;
  uLockObj.Free;


end.
