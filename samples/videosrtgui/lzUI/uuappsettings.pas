unit uuAppSettings;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, Spin;

type

  { TAppSettings }

  TAppSettings = class(TForm)
    BtnCancel: TButton;
    BtnSave: TButton;
    ChkCloseAutoUpdate: TCheckBox;
    ChkCloseOSS: TCheckBox;
    EdtOutputFileDir: TEdit;
    EdtTaskParallelCount: TSpinEdit;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    Label4: TLabel;
    Label5: TLabel;
  private

  public

  end;

var
  AppSettings: TAppSettings;

implementation

{$R *.lfm}

{ TAppSettings }



end.

