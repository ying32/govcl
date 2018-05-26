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
  published
    property AllowDropFiles: Boolean read FAllowDropFiles write SetAllowDropFiles;
    property OnDropFiles: TDropFilesEvent read FOnDropFiles write FOnDropFiles;
  end;

implementation

{$R *.dfm}

uses
  Winapi.ShellAPI;

var
  // 默认设置为False，但会显示为True，原因估计是加载
  // dfm文件问题
  uGlobalFormScaled: Boolean = False;


procedure SetGlobalFormScaled(AValue: LongBool); stdcall;
begin
  uGlobalFormScaled := AValue;
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

procedure TGoForm.InitializeNewForm;
begin
  inherited InitializeNewForm;
  Scaled := uGlobalFormScaled;
  if not Scaled then
    PixelsPerInch := 96;
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
   SetGlobalFormScaled,
   Form_ScaleForPPI,
   Form_ScaleControlsForDpi;





end.
