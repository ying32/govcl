unit uNewBaiduTranslateEngine;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls;

type

  { TNewBaiduTranslateEngine }

  TNewBaiduTranslateEngine = class(TForm)
    BtnOK: TButton;
    Button2: TButton;
    CbbType: TComboBox;
    EditAppSecret: TEdit;
    EdtAppId: TEdit;
    EdtName: TEdit;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    Label4: TLabel;
    Label5: TLabel;
    procedure BtnOKClick(Sender: TObject);
  private

  public

  end;

var
  NewBaiduTranslateEngine: TNewBaiduTranslateEngine;

implementation

{$R *.lfm}

{ TNewBaiduTranslateEngine }

procedure TNewBaiduTranslateEngine.BtnOKClick(Sender: TObject);
begin
  //
end;

end.

