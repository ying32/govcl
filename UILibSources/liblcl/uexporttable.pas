unit uExportTable;

// 因在非Windows下无法在单元中导出函数，只能换种解决方法了。

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, fgl;

  function GetProcAddr(const AName: string): Pointer;
  procedure AddToExportTable(AName: string; AProc: Pointer);
implementation

type
  TExports = specialize TFPGMap<string, Pointer>;

var uExports: TExports;

function GetProcAddr(const AName: string): Pointer;
begin
  if not uExports.TryGetData(AName, Result) then
    Result := nil;
end;

procedure AddToExportTable(AName: string; AProc: Pointer);
begin
  uExports.AddOrSetData(AName, AProc);
end;



initialization
  uExports := TExports.Create;


finalization
  uExports.Free;

end.

