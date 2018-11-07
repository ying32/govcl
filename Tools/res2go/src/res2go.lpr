program res2go;

{$mode objfpc}{$H+}

{$APPTYPE CONSOLE}

uses
  uresourceformtogo;

{$IFDEF WINDOWS}
  {$R *.res}
{$ENDIF}

begin
  //Getenv('HEAPTRC', 'disabled');
  ConvertAll;
end.

