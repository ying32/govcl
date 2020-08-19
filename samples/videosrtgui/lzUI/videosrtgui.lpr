program videosrtgui;

{$mode objfpc}{$H+}

uses
  {$IFDEF UNIX}{$IFDEF UseCThreads}
  cthreads,
  {$ENDIF}{$ENDIF}
  Interfaces, // this includes the LCL widgetset
  Forms, uMain, uuAppSettings, uNewBaiduTranslateEngine,
  uNewTencentTranslateEngine, uOSSSaveSettings
  { you can add units after this };

{$R *.res}

begin
  RequireDerivedFormResource:=True;
  Application.Scaled:=True;
  Application.Initialize;
  Application.CreateForm(TMainForm, MainForm);
  Application.CreateForm(TAppSettings, AppSettings);
  Application.CreateForm(TNewBaiduTranslateEngine, NewBaiduTranslateEngine);
  Application.CreateForm(TNewTencentTranslateEngine, NewTencentTranslateEngine);
  Application.CreateForm(TOSSSaveSettings, OSSSaveSettings);
  Application.Run;
end.

