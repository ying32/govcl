
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function GIFImage_Create: TGIFImage; extdecl;
begin
  Result :=  TGIFImage.Create;
end;

procedure GIFImage_Free(AObj: TGIFImage); extdecl;
begin
  AObj.Free;
end;

procedure GIFImage_SaveToStream(AObj: TGIFImage; Stream: TStream); extdecl;
begin
  AObj.SaveToStream(Stream);
end;

procedure GIFImage_LoadFromStream(AObj: TGIFImage; Stream: TStream); extdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure GIFImage_Clear(AObj: TGIFImage); extdecl;
begin
  AObj.Clear;
end;

procedure GIFImage_Assign(AObj: TGIFImage; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function GIFImage_Equals(AObj: TGIFImage; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure GIFImage_LoadFromFile(AObj: TGIFImage; Filename: PChar); extdecl;
begin
  AObj.LoadFromFile(Filename);
end;

procedure GIFImage_SaveToFile(AObj: TGIFImage; Filename: PChar); extdecl;
begin
  AObj.SaveToFile(Filename);
end;

procedure GIFImage_SetSize(AObj: TGIFImage; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetSize(AWidth, AHeight);
end;

function GIFImage_GetNamePath(AObj: TGIFImage): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function GIFImage_ClassName(AObj: TGIFImage): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function GIFImage_GetHashCode(AObj: TGIFImage): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function GIFImage_ToString(AObj: TGIFImage): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function GIFImage_GetBitsPerPixel(AObj: TGIFImage): Integer; extdecl;
begin
  Result :=  AObj.BitsPerPixel;
end;

function GIFImage_GetEmpty(AObj: TGIFImage): LongBool; extdecl;
begin
  Result :=  AObj.Empty;
end;

function GIFImage_GetHeight(AObj: TGIFImage): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure GIFImage_SetHeight(AObj: TGIFImage; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function GIFImage_GetModified(AObj: TGIFImage): LongBool; extdecl;
begin
  Result :=  AObj.Modified;
end;

procedure GIFImage_SetModified(AObj: TGIFImage; AValue: LongBool); extdecl;
begin
  AObj.Modified := AValue;
end;

function GIFImage_GetPaletteModified(AObj: TGIFImage): LongBool; extdecl;
begin
  Result :=  AObj.PaletteModified;
end;

procedure GIFImage_SetPaletteModified(AObj: TGIFImage; AValue: LongBool); extdecl;
begin
  AObj.PaletteModified := AValue;
end;

function GIFImage_GetTransparent(AObj: TGIFImage): LongBool; extdecl;
begin
  Result :=  AObj.Transparent;
end;

function GIFImage_GetWidth(AObj: TGIFImage): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure GIFImage_SetWidth(AObj: TGIFImage; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

procedure GIFImage_SetOnChange(AObj: TGIFImage; AEventId: NativeUInt); extdecl;
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
  GIFImage_Create,
  GIFImage_Free,
  GIFImage_SaveToStream,
  GIFImage_LoadFromStream,
  GIFImage_Clear,
  GIFImage_Assign,
  GIFImage_Equals,
  GIFImage_LoadFromFile,
  GIFImage_SaveToFile,
  GIFImage_SetSize,
  GIFImage_GetNamePath,
  GIFImage_ClassName,
  GIFImage_GetHashCode,
  GIFImage_ToString,
  GIFImage_GetBitsPerPixel,
  GIFImage_GetEmpty,
  GIFImage_GetHeight,
  GIFImage_SetHeight,
  GIFImage_GetModified,
  GIFImage_SetModified,
  GIFImage_GetPaletteModified,
  GIFImage_SetPaletteModified,
  GIFImage_GetTransparent,
  GIFImage_GetWidth,
  GIFImage_SetWidth,
  GIFImage_SetOnChange;

