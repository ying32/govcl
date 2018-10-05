program res2go;

{$mode objfpc}{$H+}

{$APPTYPE CONSOLE}

uses
  {$IFDEF UNIX}{$IFDEF UseCThreads}
  cthreads,
  {$ENDIF}{$ENDIF}
  uResourceFormToGo;


{$R *.res}

begin
  //Getenv('HEAPTRC', 'disabled');
  ConvertAll;
end.

