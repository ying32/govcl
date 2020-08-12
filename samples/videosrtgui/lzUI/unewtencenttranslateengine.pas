unit uNewTencentTranslateEngine;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls;

type

  { TNewTencentTranslateEngine }

  TNewTencentTranslateEngine = class(TForm)
    BtnOK: TButton;
    Button2: TButton;
    EditSecretKey: TEdit;
    EdtName: TEdit;
    EdtSecretId: TEdit;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    Label5: TLabel;
    procedure BtnOKClick(Sender: TObject);
  private

  public

  end;

var
  NewTencentTranslateEngine: TNewTencentTranslateEngine;

implementation

{$R *.lfm}

{ TNewTencentTranslateEngine }

procedure TNewTencentTranslateEngine.BtnOKClick(Sender: TObject);
begin
  //
end;

end.

