
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Screen_Create(AOwner: TComponent): TScreen; stdcall;
begin
  Result :=  TScreen.Create(AOwner);
end;

procedure Screen_Free(AObj: TScreen); stdcall;
begin
  AObj.Free;
end;

procedure Screen_Realign(AObj: TScreen); stdcall;
begin
  AObj.Realign;
end;

function Screen_FindComponent(AObj: TScreen; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Screen_GetNamePath(AObj: TScreen): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Screen_HasParent(AObj: TScreen): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Screen_Assign(AObj: TScreen; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Screen_ClassName(AObj: TScreen): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Screen_Equals(AObj: TScreen; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Screen_GetHashCode(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Screen_ToString(AObj: TScreen): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Screen_GetActiveForm(AObj: TScreen): TForm; stdcall;
begin
  Result :=  AObj.ActiveForm;
end;

function Screen_GetCustomFormCount(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.CustomFormCount;
end;

function Screen_GetCursorCount(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.CursorCount;
end;

function Screen_GetCursor(AObj: TScreen): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Screen_SetCursor(AObj: TScreen; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Screen_GetFocusedForm(AObj: TScreen): TCustomForm; stdcall;
begin
  Result :=  AObj.FocusedForm;
end;

procedure Screen_SetFocusedForm(AObj: TScreen; AValue: TCustomForm); stdcall;
begin
  AObj.FocusedForm := AValue;
end;

function Screen_GetMonitorCount(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.MonitorCount;
end;

procedure Screen_GetDesktopRect(AObj: TScreen; var Result: TRect); stdcall;
begin
  Result :=  AObj.DesktopRect;
end;

function Screen_GetDesktopHeight(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.DesktopHeight;
end;

function Screen_GetDesktopLeft(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.DesktopLeft;
end;

function Screen_GetDesktopTop(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.DesktopTop;
end;

function Screen_GetDesktopWidth(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.DesktopWidth;
end;

procedure Screen_GetWorkAreaRect(AObj: TScreen; var Result: TRect); stdcall;
begin
  Result :=  AObj.WorkAreaRect;
end;

function Screen_GetWorkAreaHeight(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.WorkAreaHeight;
end;

function Screen_GetWorkAreaLeft(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.WorkAreaLeft;
end;

function Screen_GetWorkAreaTop(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.WorkAreaTop;
end;

function Screen_GetWorkAreaWidth(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.WorkAreaWidth;
end;

function Screen_GetFonts(AObj: TScreen): TStrings; stdcall;
begin
  Result :=  AObj.Fonts;
end;

function Screen_GetFormCount(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.FormCount;
end;

function Screen_GetImes(AObj: TScreen): TStrings; stdcall;
begin
  Result :=  AObj.Imes;
end;

function Screen_GetDefaultIme(AObj: TScreen): PChar; stdcall;
begin
  Result :=  PChar(AObj.DefaultIme);
end;

function Screen_GetHeight(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

function Screen_GetPixelsPerInch(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.PixelsPerInch;
end;

function Screen_GetPrimaryMonitor(AObj: TScreen): TMonitor; stdcall;
begin
  Result :=  AObj.PrimaryMonitor;
end;

function Screen_GetWidth(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

function Screen_GetComponentCount(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Screen_GetComponentIndex(AObj: TScreen): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Screen_SetComponentIndex(AObj: TScreen; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Screen_GetOwner(AObj: TScreen): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Screen_GetName(AObj: TScreen): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Screen_SetName(AObj: TScreen; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Screen_GetTag(AObj: TScreen): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Screen_SetTag(AObj: TScreen; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Screen_GetCursors(AObj: TScreen; Index: Integer): HICON; stdcall;
begin
  Result :=  AObj.Cursors[Index];
end;

procedure Screen_SetCursors(AObj: TScreen; Index: Integer; AValue: HICON); stdcall;
begin
  AObj.Cursors[Index] := AValue;
end;

function Screen_GetMonitors(AObj: TScreen; Index: Integer): TMonitor; stdcall;
begin
  Result :=  AObj.Monitors[Index];
end;

function Screen_GetForms(AObj: TScreen; Index: Integer): TForm; stdcall;
begin
  Result :=  AObj.Forms[Index];
end;

function Screen_GetComponents(AObj: TScreen; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Screen_Create,
  Screen_Free,
  Screen_Realign,
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
  Screen_GetCursorCount,
  Screen_GetCursor,
  Screen_SetCursor,
  Screen_GetFocusedForm,
  Screen_SetFocusedForm,
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
  Screen_GetFormCount,
  Screen_GetImes,
  Screen_GetDefaultIme,
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

