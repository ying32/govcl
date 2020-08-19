unit uNewAliyunAudioEngine;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls;

type

  { TNewAliyunAudioEngine }

  TNewAliyunAudioEngine = class(TForm)
    BtnOK: TButton;
    Button2: TButton;
    CbbArea: TComboBox;
    EdtAccessKeyId: TEdit;
    EdtAccessKeySecret: TEdit;
    EdtAppKey: TEdit;
    EdtName: TEdit;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    Label4: TLabel;
    Label5: TLabel;
    Label6: TLabel;
    procedure BtnOKClick(Sender: TObject);
  private

  public

  end;

var
  NewAliyunAudioEngine: TNewAliyunAudioEngine;

implementation

{$R *.lfm}

{ TNewAliyunAudioEngine }

procedure TNewAliyunAudioEngine.BtnOKClick(Sender: TObject);
begin
  //
end;

end.

