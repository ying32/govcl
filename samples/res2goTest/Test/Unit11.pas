unit Unit11;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms, Vcl.Dialogs, Vcl.StdCtrls, Vcl.ExtCtrls,
  Vcl.ComCtrls, Vcl.ToolWin;

type
  TMainForm = class(TForm)
    Splitter1: TSplitter;
    Panel1: TPanel;
    Panel2: TPanel;
    BtnShowAbout: TButton;
    BtnShowFrame1: TButton;
    BtnShowFrame2: TButton;
    procedure FormCreate(Sender: TObject);
    procedure BtnShowAboutClick(Sender: TObject);
    procedure BtnShowFrame1Click(Sender: TObject);
    procedure BtnShowFrame2Click(Sender: TObject);
  private
    { Private declarations }
  public
    { Public declarations }
  end;

var
  MainForm: TMainForm;

implementation

{$R *.dfm}

procedure TMainForm.BtnShowAboutClick(Sender: TObject);
begin
//
end;

procedure TMainForm.BtnShowFrame1Click(Sender: TObject);
begin
//
end;

procedure TMainForm.BtnShowFrame2Click(Sender: TObject);
begin
//
end;

procedure TMainForm.FormCreate(Sender: TObject);
begin
  //
end;

end.
