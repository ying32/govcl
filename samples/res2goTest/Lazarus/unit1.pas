unit Unit1;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls,
  ExtCtrls;

type

  { TMainForm }

  TMainForm = class(TForm)
    Button1: TButton;
    Button2: TButton;
    Button3: TButton;
    Button4: TButton;
    Label1: TLabel;
    Panel1: TPanel;
    Panel2: TPanel;
    procedure Button1Click(Sender: TObject);
    procedure Button2Click(Sender: TObject);
    procedure Button2Resize(Sender: TObject);
    procedure FormCreate(Sender: TObject);
    procedure Panel1MouseEnter(Sender: TObject);
    procedure Panel1MouseLeave(Sender: TObject);
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

{ TMainForm }

procedure TMainForm.Button1Click(Sender: TObject);
begin
  //
end;

procedure TMainForm.Button2Click(Sender: TObject);
begin
  //
end;

procedure TMainForm.Button2Resize(Sender: TObject);
begin
  //
end;

procedure TMainForm.FormCreate(Sender: TObject);
begin
  //
end;

procedure TMainForm.Panel1MouseEnter(Sender: TObject);
begin
  //
end;

procedure TMainForm.Panel1MouseLeave(Sender: TObject);
begin
  //
end;

end.

