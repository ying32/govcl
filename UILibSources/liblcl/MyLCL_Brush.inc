
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Brush_Create: TBrush; extdecl;
begin
  Result :=  TBrush.Create;
end;

procedure Brush_Free(AObj: TBrush); extdecl;
begin
  AObj.Free;
end;

procedure Brush_Assign(AObj: TBrush; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Brush_GetNamePath(AObj: TBrush): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Brush_ClassName(AObj: TBrush): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Brush_Equals(AObj: TBrush; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Brush_GetHashCode(AObj: TBrush): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Brush_ToString(AObj: TBrush): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Brush_GetBitmap(AObj: TBrush): TBitmap; extdecl;
begin
  Result :=  AObj.Bitmap as TBitmap;
end;

procedure Brush_SetBitmap(AObj: TBrush; AValue: TBitmap); extdecl;
begin
  AObj.Bitmap := AValue;
end;

function Brush_GetHandle(AObj: TBrush): HBRUSH; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure Brush_SetHandle(AObj: TBrush; AValue: HBRUSH); extdecl;
begin
  AObj.Handle := AValue;
end;

function Brush_GetColor(AObj: TBrush): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure Brush_SetColor(AObj: TBrush; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function Brush_GetStyle(AObj: TBrush): TBrushStyle; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure Brush_SetStyle(AObj: TBrush; AValue: TBrushStyle); extdecl;
begin
  AObj.Style := AValue;
end;

procedure Brush_SetOnChange(AObj: TBrush; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    EventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := @EventClass.OnChange;
  EventClass.Add(AObj, geChange, AEventId);
end;


exports
  Brush_Create,
  Brush_Free,
  Brush_Assign,
  Brush_GetNamePath,
  Brush_ClassName,
  Brush_Equals,
  Brush_GetHashCode,
  Brush_ToString,
  Brush_GetBitmap,
  Brush_SetBitmap,
  Brush_GetHandle,
  Brush_SetHandle,
  Brush_GetColor,
  Brush_SetColor,
  Brush_GetStyle,
  Brush_SetStyle,
  Brush_SetOnChange;

