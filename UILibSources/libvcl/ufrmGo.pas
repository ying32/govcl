//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

unit ufrmGo;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Winapi.ActiveX;

type
  TDropFilesEvent = procedure(Sender: TObject; const FileNames: array of String) of object;

  TWndProcEvent = procedure(Sender: TObject; var AMsg: TMessage; var AHandled: Boolean) of object;

  // 兼容Lazarus的
  TShowInTaskbar = (
    stDefault,  // use default rules for showing taskbar item
    stAlways,   // always show taskbar item for the form
    stNever     // never show taskbar item for the form
  );

  TGoForm = class(TForm)
  private
    FOnDropFiles: TDropFilesEvent;
    FOnStyleChanged: TNotifyEvent;
    FOnWndProc: TWndProcEvent;
    FAllowDropFiles: Boolean;
    FShowInTaskBar: TShowInTaskbar;
    procedure SetAllowDropFiles(const Value: Boolean);
    procedure WMDropFiles(var Msg: TWMDropFiles); message WM_DROPFILES;
    procedure CMStyleChanged(var Msg: TMessage); message CM_STYLECHANGED;
    function GetAllowDropFiles: Boolean;
    procedure SetShowInTaskBar(const Value: TShowInTaskbar);
    procedure UpdateShowInTaskBar;
  public
    constructor Create(AOwner: TComponent); override;
    procedure WndProc(var AMsg: TMessage); override;

    /// <summary>
    /// Checks if there is a change in dpi and perform the necessary changes to scale all
    /// the controls for the new dpi
    /// </summary>
    procedure ScaleForCurrentDpi; override;
  published
    property AllowDropFiles: Boolean read GetAllowDropFiles write SetAllowDropFiles;
    property OnDropFiles: TDropFilesEvent read FOnDropFiles write FOnDropFiles;
    property OnStyleChanged: TNotifyEvent read FOnStyleChanged write FOnStyleChanged;
    property OnWndProc: TWndProcEvent read FOnWndProc write FOnWndProc;
    property ShowInTaskBar: TShowInTaskbar read FShowInTaskBar write SetShowInTaskBar default stDefault;
  end;

//  procedure LockInitScale;
//  procedure UnLockInitScale;
//  procedure SetInitScale(AValue: Boolean);
implementation

//{$R *.dfm}

uses
  uComponents,
  Winapi.ShellAPI;

//var
//  uLockObj: TObject;
//  uInitScale: Boolean;

//procedure LockInitScale;
//begin
//  System.TMonitor.Enter(uLockObj);
//end;

//procedure UnLockInitScale;
//begin
//  System.TMonitor.Exit(uLockObj);
//end;

//procedure SetInitScale(AValue: Boolean);
//begin
//  uInitScale := AValue;
//end;

procedure Form_ScaleForPPI(AObj: TGoForm; ANewPPI: Integer); stdcall;
begin
  AObj.ScaleForPPI(ANewPPI);
end;

procedure Form_ScaleControlsForDpi(AObj: TGoForm; ANewPPI: Integer); stdcall;
begin
  AObj.ScaleControlsForDpi(ANewPPI);
end;

procedure Form_ScaleForCurrentDpi(AObj: TGoForm); stdcall;
begin
  AObj.ScaleForCurrentDpi;
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
begin
  try
    inherited Create(AOwner);
  except
    // 不处理这个异常
  end;
  ControlStyle := ControlStyle + [csPaintBlackOpaqueOnGlass];
  FShowInTaskBar := stDefault;
end;

function TGoForm.GetAllowDropFiles: Boolean;
begin
  Result := (GetWindowLong(Handle, GWL_EXSTYLE) and WS_EX_ACCEPTFILES) <> 0;
end;

procedure TGoForm.ScaleForCurrentDpi;
begin
  inherited;
end;

procedure TGoForm.SetAllowDropFiles(const Value: Boolean);
begin
  FAllowDropFiles := Value;
  if AllowDropFiles <> Value then
    DragAcceptFiles(Handle, Value);
end;

procedure TGoForm.SetShowInTaskBar(const Value: TShowInTaskbar);
begin
  if FShowInTaskBar <> Value then
  begin
    FShowInTaskBar := Value;
    UpdateShowInTaskBar;
  end;
end;

procedure TGoForm.UpdateShowInTaskBar;
  function IsSetVal: Boolean;
  begin
    Result := (GetWindowLong(Handle, GWL_EXSTYLE) and WS_EX_ACCEPTFILES) <> 0;
  end;
begin
  if (Assigned(Application) and (Application.MainForm = Self)) or
   (not HandleAllocated) or Assigned(Parent) or
   (FormStyle = fsMDIChild){ or not Showing} then
    Exit;
  if FShowInTaskBar = stAlways then
  begin
    if not IsSetVal then
      SetWindowLong(Handle, GWL_EXSTYLE, GetWindowLong(Handle, GWL_EXSTYLE) or WS_EX_APPWINDOW);
  end else
  begin
    if IsSetVal then
      SetWindowLong(Handle, GWL_EXSTYLE, GetWindowLong(Handle, GWL_EXSTYLE) and not WS_EX_APPWINDOW);
  end;
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
   Form_ScaleControlsForDpi,
   Form_ScaleForCurrentDpi;


initialization
  CoInitialize(nil);

finalization
  CoUninitialize;

end.
