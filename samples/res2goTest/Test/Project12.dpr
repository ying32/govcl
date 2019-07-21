program Project12;

uses
  Vcl.Forms,
  Unit11 in 'Unit11.pas' {MainForm},
  Unit1 in 'Unit1.pas' {About},
  Unit2 in 'Unit2.pas' {Frame2: TFrame},
  Unit3 in 'Unit3.pas' {Frame3: TFrame};

{$R *.res}

begin
  Application.Initialize;
  Application.MainFormOnTaskbar := True;
  Application.CreateForm(TMainForm, MainForm);
  Application.Run;
end.
