
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Screen_Create(AOwner: TComponent): TScreen; extdecl;
begin
  Result :=  TScreen.Create(AOwner);
end;

procedure Screen_Free(AObj: TScreen); extdecl;
begin
  AObj.Free;
end;

function Screen_FindComponent(AObj: TScreen; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Screen_GetNamePath(AObj: TScreen): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Screen_HasParent(AObj: TScreen): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Screen_Assign(AObj: TScreen; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Screen_ClassName(AObj: TScreen): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Screen_Equals(AObj: TScreen; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Screen_GetHashCode(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Screen_ToString(AObj: TScreen): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Screen_GetActiveForm(AObj: TScreen): TForm; extdecl;
begin
  Result :=  AObj.ActiveForm;
end;

function Screen_GetCustomFormCount(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.CustomFormCount;
end;

function Screen_GetCursor(AObj: TScreen): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure Screen_SetCursor(AObj: TScreen; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function Screen_GetFocusedForm(AObj: TScreen): TCustomForm; extdecl;
begin
  Result :=  AObj.FocusedForm;
end;

function Screen_GetMonitorCount(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.MonitorCount;
end;

procedure Screen_GetDesktopRect(AObj: TScreen; var Result: TRect); extdecl;
begin
  Result :=  AObj.DesktopRect;
end;

function Screen_GetDesktopHeight(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.DesktopHeight;
end;

function Screen_GetDesktopLeft(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.DesktopLeft;
end;

function Screen_GetDesktopTop(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.DesktopTop;
end;

function Screen_GetDesktopWidth(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.DesktopWidth;
end;

procedure Screen_GetWorkAreaRect(AObj: TScreen; var Result: TRect); extdecl;
begin
  Result :=  AObj.WorkAreaRect;
end;

function Screen_GetWorkAreaHeight(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.WorkAreaHeight;
end;

function Screen_GetWorkAreaLeft(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.WorkAreaLeft;
end;

function Screen_GetWorkAreaTop(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.WorkAreaTop;
end;

function Screen_GetWorkAreaWidth(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.WorkAreaWidth;
end;

function Screen_GetFonts(AObj: TScreen): TStrings; extdecl;
begin
  Result :=  AObj.Fonts;
end;

function Screen_GetHeight(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

function Screen_GetPixelsPerInch(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.PixelsPerInch;
end;

function Screen_GetPrimaryMonitor(AObj: TScreen): TMonitor; extdecl;
begin
  Result :=  AObj.PrimaryMonitor;
end;

function Screen_GetWidth(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

function Screen_GetComponentCount(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Screen_GetComponentIndex(AObj: TScreen): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Screen_SetComponentIndex(AObj: TScreen; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Screen_GetOwner(AObj: TScreen): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Screen_GetName(AObj: TScreen): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Screen_SetName(AObj: TScreen; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Screen_GetTag(AObj: TScreen): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Screen_SetTag(AObj: TScreen; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Screen_GetCursors(AObj: TScreen; Index: Integer): HICON; extdecl;
begin
  Result :=  AObj.Cursors[Index];
end;

procedure Screen_SetCursors(AObj: TScreen; Index: Integer; AValue: HICON); extdecl;
begin
  AObj.Cursors[Index] := AValue;
end;

function Screen_GetMonitors(AObj: TScreen; Index: Integer): TMonitor; extdecl;
begin
  Result :=  AObj.Monitors[Index];
end;

function Screen_GetForms(AObj: TScreen; Index: Integer): TForm; extdecl;
begin
  Result :=  AObj.Forms[Index];
end;

function Screen_GetComponents(AObj: TScreen; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Screen_Create,
  Screen_Free,
  Screen_FindComponent,
  Screen_GetNamePath,
  Screen_HasParent,
  Screen_Assign,
  Screen_ClassName,
  Screen_Equals,
  Screen_GetHashCode,
  Screen_ToString,
  Screen_GetActiveForm,
  Screen_GetCustomFormCount,
  Screen_GetCursor,
  Screen_SetCursor,
  Screen_GetFocusedForm,
  Screen_GetMonitorCount,
  Screen_GetDesktopRect,
  Screen_GetDesktopHeight,
  Screen_GetDesktopLeft,
  Screen_GetDesktopTop,
  Screen_GetDesktopWidth,
  Screen_GetWorkAreaRect,
  Screen_GetWorkAreaHeight,
  Screen_GetWorkAreaLeft,
  Screen_GetWorkAreaTop,
  Screen_GetWorkAreaWidth,
  Screen_GetFonts,
  Screen_GetHeight,
  Screen_GetPixelsPerInch,
  Screen_GetPrimaryMonitor,
  Screen_GetWidth,
  Screen_GetComponentCount,
  Screen_GetComponentIndex,
  Screen_SetComponentIndex,
  Screen_GetOwner,
  Screen_GetName,
  Screen_SetName,
  Screen_GetTag,
  Screen_SetTag,
  Screen_GetCursors,
  Screen_SetCursors,
  Screen_GetMonitors,
  Screen_GetForms,
  Screen_GetComponents;

