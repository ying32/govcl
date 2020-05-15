unit uNewConnection;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ExtCtrls, StdCtrls;

type

  { TNewConnection }

  TNewConnection = class(TForm)
    BtnTestConnection: TButton;
    BtnOk: TButton;
    BtnCancel: TButton;
    EdtConnName: TEdit;
    EdtHost: TEdit;
    EdtPort: TEdit;
    EdtPassword: TEdit;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    Label5: TLabel;
    Panel1: TPanel;
    Panel2: TPanel;
    procedure BtnTestConnectionClick(Sender: TObject);
    procedure FormCreate(Sender: TObject);
  private

  public

  end;

var
  NewConnection: TNewConnection;

implementation

{$R *.lfm}

{ TNewConnection }

procedure TNewConnection.FormCreate(Sender: TObject);
begin
  //
end;

procedure TNewConnection.BtnTestConnectionClick(Sender: TObject);
begin
  //
end;

end.

