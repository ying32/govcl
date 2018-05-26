
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Mouse_Create: TMouse; extdecl;
begin
  Result :=  TMouse.Create;
end;

procedure Mouse_Free(AObj: TMouse); extdecl;
begin
  AObj.Free;
end;

function Mouse_ClassName(AObj: TMouse): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Mouse_Equals(AObj: TMouse; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Mouse_GetHashCode(AObj: TMouse): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Mouse_ToString(AObj: TMouse): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Mouse_GetCapture(AObj: TMouse): HWND; extdecl;
begin
  Result :=  AObj.Capture;
end;

procedure Mouse_SetCapture(AObj: TMouse; AValue: HWND); extdecl;
begin
  AObj.Capture := AValue;
end;

procedure Mouse_GetCursorPos(AObj: TMouse; var Result: TPoint); extdecl;
begin
  Result :=  AObj.CursorPos;
end;

procedure Mouse_SetCursorPos(AObj: TMouse; var AValue: TPoint); extdecl;
begin
  AObj.CursorPos := AValue;
end;

function Mouse_GetIsDragging(AObj: TMouse): LongBool; extdecl;
begin
  Result :=  AObj.IsDragging;
end;

function Mouse_GetWheelScrollLines(AObj: TMouse): Integer; extdecl;
begin
  Result :=  AObj.WheelScrollLines;
end;


exports
  Mouse_Create,
  Mouse_Free,
  Mouse_ClassName,
  Mouse_Equals,
  Mouse_GetHashCode,
  Mouse_ToString,
  Mouse_GetCapture,
  Mouse_SetCapture,
  Mouse_GetCursorPos,
  Mouse_SetCursorPos,
  Mouse_GetIsDragging,
  Mouse_GetWheelScrollLines;

