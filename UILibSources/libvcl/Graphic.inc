
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Graphic_Create: TGraphic; stdcall;
begin
  Result :=  TGraphic.Create;
end;

procedure Graphic_Free(AObj: TGraphic); stdcall;
begin
  AObj.Free;
end;

function Graphic_Equals(AObj: TGraphic; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure Graphic_LoadFromFile(AObj: TGraphic; Filename: PChar); stdcall;
begin
  AObj.LoadFromFile(Filename);
end;

procedure Graphic_SaveToFile(AObj: TGraphic; Filename: PChar); stdcall;
begin
  AObj.SaveToFile(Filename);
end;

procedure Graphic_LoadFromStream(AObj: TGraphic; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure Graphic_SaveToStream(AObj: TGraphic; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

procedure Graphic_SetSize(AObj: TGraphic; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetSize(AWidth, AHeight);
end;

procedure Graphic_Assign(AObj: TGraphic; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Graphic_GetNamePath(AObj: TGraphic): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Graphic_ClassName(AObj: TGraphic): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Graphic_GetHashCode(AObj: TGraphic): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Graphic_ToString(AObj: TGraphic): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Graphic_GetEmpty(AObj: TGraphic): LongBool; stdcall;
begin
  Result :=  AObj.Empty;
end;

function Graphic_GetHeight(AObj: TGraphic): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Graphic_SetHeight(AObj: TGraphic; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Graphic_GetModified(AObj: TGraphic): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure Graphic_SetModified(AObj: TGraphic; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function Graphic_GetPaletteModified(AObj: TGraphic): LongBool; stdcall;
begin
  Result :=  AObj.PaletteModified;
end;

procedure Graphic_SetPaletteModified(AObj: TGraphic; AValue: LongBool); stdcall;
begin
  AObj.PaletteModified := AValue;
end;

function Graphic_GetTransparent(AObj: TGraphic): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure Graphic_SetTransparent(AObj: TGraphic; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function Graphic_GetWidth(AObj: TGraphic): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Graphic_SetWidth(AObj: TGraphic; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

procedure Graphic_SetOnChange(AObj: TGraphic; AEventId: NativeUInt); stdcall;
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
  Graphic_Create,
  Graphic_Free,
  Graphic_Equals,
  Graphic_LoadFromFile,
  Graphic_SaveToFile,
  Graphic_LoadFromStream,
  Graphic_SaveToStream,
  Graphic_SetSize,
  Graphic_Assign,
  Graphic_GetNamePath,
  Graphic_ClassName,
  Graphic_GetHashCode,
  Graphic_ToString,
  Graphic_GetEmpty,
  Graphic_GetHeight,
  Graphic_SetHeight,
  Graphic_GetModified,
  Graphic_SetModified,
  Graphic_GetPaletteModified,
  Graphic_SetPaletteModified,
  Graphic_GetTransparent,
  Graphic_SetTransparent,
  Graphic_GetWidth,
  Graphic_SetWidth,
  Graphic_SetOnChange;

