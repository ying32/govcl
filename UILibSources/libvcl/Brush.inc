
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Brush_Create: TBrush; stdcall;
begin
  Result :=  TBrush.Create;
end;

procedure Brush_Free(AObj: TBrush); stdcall;
begin
  AObj.Free;
end;

procedure Brush_Assign(AObj: TBrush; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Brush_HandleAllocated(AObj: TBrush): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

function Brush_GetNamePath(AObj: TBrush): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Brush_ClassName(AObj: TBrush): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Brush_Equals(AObj: TBrush; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Brush_GetHashCode(AObj: TBrush): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Brush_ToString(AObj: TBrush): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Brush_GetBitmap(AObj: TBrush): TBitmap; stdcall;
begin
  Result :=  AObj.Bitmap;
end;

procedure Brush_SetBitmap(AObj: TBrush; AValue: TBitmap); stdcall;
begin
  AObj.Bitmap := AValue;
end;

function Brush_GetHandle(AObj: TBrush): HBRUSH; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure Brush_SetHandle(AObj: TBrush; AValue: HBRUSH); stdcall;
begin
  AObj.Handle := AValue;
end;

function Brush_GetColor(AObj: TBrush): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure Brush_SetColor(AObj: TBrush; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function Brush_GetStyle(AObj: TBrush): TBrushStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure Brush_SetStyle(AObj: TBrush; AValue: TBrushStyle); stdcall;
begin
  AObj.Style := AValue;
end;

procedure Brush_SetOnChange(AObj: TBrush; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, geChange, AEventId);
end;


exports
  Brush_Create,
  Brush_Free,
  Brush_Assign,
  Brush_HandleAllocated,
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

