unit uOSSSaveSettings;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls;

type

  { TOSSSaveSettings }

  TOSSSaveSettings = class(TForm)
    BtnCancel: TButton;
    BtnSave: TButton;
    EdtAccessKeyId: TEdit;
    EdtAccessKeySecret: TEdit;
    EdtBucketDomain: TEdit;
    EdtBucketName: TEdit;
    EdtEndPoint: TEdit;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    Label4: TLabel;
    Label5: TLabel;
    Label6: TLabel;
    procedure BtnSaveClick(Sender: TObject);
  private

  public

  end;

var
  OSSSaveSettings: TOSSSaveSettings;

implementation

{$R *.lfm}

{ TOSSSaveSettings }

procedure TOSSSaveSettings.BtnSaveClick(Sender: TObject);
begin
  //
end;

end.

