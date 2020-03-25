unit Unit1;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ExtCtrls, StdCtrls;

type

  { TMainForm }

  TMainForm = class(TForm)
    BtnShowAbout: TButton;
    BtnShowFrame1: TButton;
    BtnShowFrame2: TButton;
    Panel1: TPanel;
    Panel2: TPanel;
    Splitter1: TSplitter;
    procedure BtnShowAboutClick(Sender: TObject);
    procedure BtnShowFrame1Click(Sender: TObject);
    procedure BtnShowFrame2Click(Sender: TObject);
    procedure FormCreate(Sender: TObject);
  private
    procedure OnFindComponentClass(AReader: TReader; const AClassName: string; var AComponentClass: TComponentClass);
  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

uses
  Unit2;

{ TMainForm }

procedure TMainForm.BtnShowAboutClick(Sender: TObject);
begin

end;

procedure TMainForm.BtnShowFrame1Click(Sender: TObject);
begin

end;

procedure TMainForm.BtnShowFrame2Click(Sender: TObject);
begin

end;

procedure TMainForm.OnFindComponentClass(AReader: TReader; const AClassName: string;
  var AComponentClass: TComponentClass);
var
  LB: TClass;
begin
  if AClassName = 'TButton' then
   AComponentClass := TButton;
end;

procedure TMainForm.FormCreate(Sender: TObject);
var
   LFrame: TFrame;
   LReader: TReader;
   LStream: TFileStream;
begin
  LFrame := TFrame.Create(Self);
  LFrame.Parent := Panel2;
  LStream := TFileStream.Create('.\gocode\Frame1.gfm', fmOpenRead);
  LReader := TReader.Create(LStream, 4096);
  LReader.OnFindComponentClass := @OnFindComponentClass;
  LReader.ReadRootComponent(LFrame);
  LStream.free;
  //with TFrame1.Create(Self) do
  //  Parent := Panel2;
end;

end.

