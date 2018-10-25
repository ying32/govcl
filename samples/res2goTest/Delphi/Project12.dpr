program Project12;

uses
  Vcl.Forms,
  Unit11 in 'Unit11.pas' {MainForm},
  Unit1 in 'Unit1.pas' {About};

{$R *.res}

begin
  Application.Initialize;
  Application.MainFormOnTaskbar := True;
  Application.CreateForm(TMainForm, MainForm);
  Application.CreateForm(TAbout, About);
  Application.Run;
end.
