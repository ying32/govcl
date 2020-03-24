program Project12;

uses
{$IFDEF UNIX}{$IFDEF UseCThreads}
  cthreads,
{$ENDIF}{$ENDIF}
  Interfaces, // this includes the LCL widgetset,
  Forms,
  Unit11 in 'Unit11.pas' {Form11};

{$R *.res}

begin
  Application.Initialize;
  Application.MainFormOnTaskbar := True;
  Application.CreateForm(TForm11, Form11);
  Application.Run;
end.
