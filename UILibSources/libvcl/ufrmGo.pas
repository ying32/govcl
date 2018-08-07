unit ufrmGo;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms;

type
  TDropFilesEvent = procedure(Sender: TObject; const FileNames: array of String) of object;

  TGoForm = class(TForm)
  private
    FAllowDropFiles: Boolean;
    FOnDropFiles: TDropFilesEvent;
    procedure SetAllowDropFiles(const Value: Boolean);
    procedure WMDropFiles(var Msg: TWMDropFiles); message WM_DROPFILES;
  protected
    procedure InitializeNewForm; override;
  public
    constructor Create(AOwner: TComponent); override;
  published
    property AllowDropFiles: Boolean read FAllowDropFiles write SetAllowDropFiles;
    property OnDropFiles: TDropFilesEvent read FOnDropFiles write FOnDropFiles;
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

constructor TGoForm.Create(AOwner: TComponent);
var
  LPPI: Integer;
  LR: TRect;
begin
  CreateNew(AOwner, 0);
  if uInitScale and GetGlobalFormScaled then
  begin
    LPPI := Screen.PixelsPerInch;
    ClientWidth := MulDiv(ClientWidth, LPPI, 96);
    ClientHeight := MulDiv(ClientHeight, LPPI, 96);
    ScaleForPPI(LPPI);
  end;
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
  if FAllowDropFiles <> Value then
  begin
    FAllowDropFiles := Value;
    DragAcceptFiles(Handle, Value);
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

exports
   Form_ScaleForPPI,
   Form_ScaleControlsForDpi;


initialization
  uLockObj := TObject.Create;

finalization
  uLockObj.Free;


end.
