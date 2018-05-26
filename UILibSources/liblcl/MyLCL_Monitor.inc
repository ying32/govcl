
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Monitor_Create: TMonitor; extdecl;
begin
  Result :=  TMonitor.Create;
end;

procedure Monitor_Free(AObj: TMonitor); extdecl;
begin
  AObj.Free;
end;

function Monitor_ClassName(AObj: TMonitor): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Monitor_Equals(AObj: TMonitor; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Monitor_GetHashCode(AObj: TMonitor): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Monitor_ToString(AObj: TMonitor): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Monitor_GetHandle(AObj: TMonitor): HMONITOR; extdecl;
begin
  Result :=  AObj.Handle;
end;

function Monitor_GetMonitorNum(AObj: TMonitor): Integer; extdecl;
begin
  Result :=  AObj.MonitorNum;
end;

function Monitor_GetLeft(AObj: TMonitor): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

function Monitor_GetHeight(AObj: TMonitor): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

function Monitor_GetTop(AObj: TMonitor): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

function Monitor_GetWidth(AObj: TMonitor): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Monitor_GetBoundsRect(AObj: TMonitor; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Monitor_GetWorkareaRect(AObj: TMonitor; var Result: TRect); extdecl;
begin
  Result :=  AObj.WorkareaRect;
end;

function Monitor_GetPrimary(AObj: TMonitor): LongBool; extdecl;
begin
  Result :=  AObj.Primary;
end;

function Monitor_GetPixelsPerInch(AObj: TMonitor): Integer; extdecl;
begin
  Result :=  AObj.PixelsPerInch;
end;


exports
  Monitor_Create,
  Monitor_Free,
  Monitor_ClassName,
  Monitor_Equals,
  Monitor_GetHashCode,
  Monitor_ToString,
  Monitor_GetHandle,
  Monitor_GetMonitorNum,
  Monitor_GetLeft,
  Monitor_GetHeight,
  Monitor_GetTop,
  Monitor_GetWidth,
  Monitor_GetBoundsRect,
  Monitor_GetWorkareaRect,
  Monitor_GetPrimary,
  Monitor_GetPixelsPerInch;

