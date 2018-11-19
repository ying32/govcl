
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Mouse_Create: TMouse; stdcall;
begin
  Result :=  TMouse.Create;
end;

procedure Mouse_Free(AObj: TMouse); stdcall;
begin
  AObj.Free;
end;

procedure Mouse_DisposeOf(AObj: TMouse); stdcall;
begin
  AObj.DisposeOf;
end;

function Mouse_ClassType(AObj: TMouse): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function Mouse_ClassName(AObj: TMouse): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Mouse_InstanceSize(AObj: TMouse): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function Mouse_InheritsFrom(AObj: TMouse; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function Mouse_Equals(AObj: TMouse; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Mouse_GetHashCode(AObj: TMouse): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Mouse_ToString(AObj: TMouse): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Mouse_GetCapture(AObj: TMouse): HWND; stdcall;
begin
  Result :=  AObj.Capture;
end;

procedure Mouse_SetCapture(AObj: TMouse; AValue: HWND); stdcall;
begin
  AObj.Capture := AValue;
end;

procedure Mouse_GetCursorPos(AObj: TMouse; var Result: TPoint); stdcall;
begin
  Result :=  AObj.CursorPos;
end;

procedure Mouse_SetCursorPos(AObj: TMouse; var AValue: TPoint); stdcall;
begin
  AObj.CursorPos := AValue;
end;

function Mouse_GetIsDragging(AObj: TMouse): LongBool; stdcall;
begin
  Result :=  AObj.IsDragging;
end;

function Mouse_GetIsPanning(AObj: TMouse): LongBool; stdcall;
begin
  Result :=  AObj.IsPanning;
end;

function Mouse_GetWheelPresent(AObj: TMouse): LongBool; stdcall;
begin
  Result :=  AObj.WheelPresent;
end;

function Mouse_GetWheelScrollLines(AObj: TMouse): Integer; stdcall;
begin
  Result :=  AObj.WheelScrollLines;
end;

function Mouse_StaticClassType: TClass; stdcall;
begin
  Result := TMouse;
end;

exports
  Mouse_Create,
  Mouse_Free,
  Mouse_DisposeOf,
  Mouse_ClassType,
  Mouse_ClassName,
  Mouse_InstanceSize,
  Mouse_InheritsFrom,
  Mouse_Equals,
  Mouse_GetHashCode,
  Mouse_ToString,
  Mouse_GetCapture,
  Mouse_SetCapture,
  Mouse_GetCursorPos,
  Mouse_SetCursorPos,
  Mouse_GetIsDragging,
  Mouse_GetIsPanning,
  Mouse_GetWheelPresent,
  Mouse_GetWheelScrollLines,
  Mouse_StaticClassType;

