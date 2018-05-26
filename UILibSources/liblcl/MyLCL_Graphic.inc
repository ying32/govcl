
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Graphic_Create: TGraphic; extdecl;
begin
  Result :=  TGraphic.Create;
end;

procedure Graphic_Free(AObj: TGraphic); extdecl;
begin
  AObj.Free;
end;

function Graphic_Equals(AObj: TGraphic; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure Graphic_LoadFromFile(AObj: TGraphic; Filename: PChar); extdecl;
begin
  AObj.LoadFromFile(Filename);
end;

procedure Graphic_SaveToFile(AObj: TGraphic; Filename: PChar); extdecl;
begin
  AObj.SaveToFile(Filename);
end;

procedure Graphic_LoadFromStream(AObj: TGraphic; Stream: TStream); extdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure Graphic_SaveToStream(AObj: TGraphic; Stream: TStream); extdecl;
begin
  AObj.SaveToStream(Stream);
end;

procedure Graphic_Assign(AObj: TGraphic; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Graphic_GetNamePath(AObj: TGraphic): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Graphic_ClassName(AObj: TGraphic): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Graphic_GetHashCode(AObj: TGraphic): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Graphic_ToString(AObj: TGraphic): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Graphic_GetEmpty(AObj: TGraphic): LongBool; extdecl;
begin
  Result :=  AObj.Empty;
end;

function Graphic_GetHeight(AObj: TGraphic): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Graphic_SetHeight(AObj: TGraphic; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Graphic_GetModified(AObj: TGraphic): LongBool; extdecl;
begin
  Result :=  AObj.Modified;
end;

procedure Graphic_SetModified(AObj: TGraphic; AValue: LongBool); extdecl;
begin
  AObj.Modified := AValue;
end;

function Graphic_GetPaletteModified(AObj: TGraphic): LongBool; extdecl;
begin
  Result :=  AObj.PaletteModified;
end;

procedure Graphic_SetPaletteModified(AObj: TGraphic; AValue: LongBool); extdecl;
begin
  AObj.PaletteModified := AValue;
end;

function Graphic_GetTransparent(AObj: TGraphic): LongBool; extdecl;
begin
  Result :=  AObj.Transparent;
end;

procedure Graphic_SetTransparent(AObj: TGraphic; AValue: LongBool); extdecl;
begin
  AObj.Transparent := AValue;
end;

function Graphic_GetWidth(AObj: TGraphic): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Graphic_SetWidth(AObj: TGraphic; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

procedure Graphic_SetOnChange(AObj: TGraphic; AEventId: NativeUInt); extdecl;
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
  Graphic_Create,
  Graphic_Free,
  Graphic_Equals,
  Graphic_LoadFromFile,
  Graphic_SaveToFile,
  Graphic_LoadFromStream,
  Graphic_SaveToStream,
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

