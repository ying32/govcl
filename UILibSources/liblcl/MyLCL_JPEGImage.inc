
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function JPEGImage_Create: TJPEGImage; extdecl;
begin
  Result :=  TJPEGImage.Create;
end;

procedure JPEGImage_Free(AObj: TJPEGImage); extdecl;
begin
  AObj.Free;
end;

procedure JPEGImage_Assign(AObj: TJPEGImage; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure JPEGImage_LoadFromStream(AObj: TJPEGImage; Stream: TStream); extdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure JPEGImage_SaveToStream(AObj: TJPEGImage; Stream: TStream); extdecl;
begin
  AObj.SaveToStream(Stream);
end;

function JPEGImage_Equals(AObj: TJPEGImage; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure JPEGImage_LoadFromFile(AObj: TJPEGImage; Filename: PChar); extdecl;
begin
  AObj.LoadFromFile(Filename);
end;

procedure JPEGImage_SaveToFile(AObj: TJPEGImage; Filename: PChar); extdecl;
begin
  AObj.SaveToFile(Filename);
end;

procedure JPEGImage_SetSize(AObj: TJPEGImage; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetSize(AWidth, AHeight);
end;

function JPEGImage_GetNamePath(AObj: TJPEGImage): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function JPEGImage_ClassName(AObj: TJPEGImage): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function JPEGImage_GetHashCode(AObj: TJPEGImage): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function JPEGImage_ToString(AObj: TJPEGImage): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function JPEGImage_GetPixelFormat(AObj: TJPEGImage): TPixelFormat; extdecl;
begin
  Result :=  AObj.PixelFormat;
end;

procedure JPEGImage_SetPixelFormat(AObj: TJPEGImage; AValue: TJPEGPixelFormat); extdecl;
begin
  AObj.PixelFormat := AValue;
end;

function JPEGImage_GetPerformance(AObj: TJPEGImage): TJPEGPerformance; extdecl;
begin
  Result :=  AObj.Performance;
end;

procedure JPEGImage_SetPerformance(AObj: TJPEGImage; AValue: TJPEGPerformance); extdecl;
begin
  AObj.Performance := AValue;
end;

function JPEGImage_GetCanvas(AObj: TJPEGImage): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function JPEGImage_GetEmpty(AObj: TJPEGImage): LongBool; extdecl;
begin
  Result :=  AObj.Empty;
end;

function JPEGImage_GetHeight(AObj: TJPEGImage): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure JPEGImage_SetHeight(AObj: TJPEGImage; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function JPEGImage_GetModified(AObj: TJPEGImage): LongBool; extdecl;
begin
  Result :=  AObj.Modified;
end;

procedure JPEGImage_SetModified(AObj: TJPEGImage; AValue: LongBool); extdecl;
begin
  AObj.Modified := AValue;
end;

function JPEGImage_GetPaletteModified(AObj: TJPEGImage): LongBool; extdecl;
begin
  Result :=  AObj.PaletteModified;
end;

procedure JPEGImage_SetPaletteModified(AObj: TJPEGImage; AValue: LongBool); extdecl;
begin
  AObj.PaletteModified := AValue;
end;

function JPEGImage_GetTransparent(AObj: TJPEGImage): LongBool; extdecl;
begin
  Result :=  AObj.Transparent;
end;

procedure JPEGImage_SetTransparent(AObj: TJPEGImage; AValue: LongBool); extdecl;
begin
  AObj.Transparent := AValue;
end;

function JPEGImage_GetWidth(AObj: TJPEGImage): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure JPEGImage_SetWidth(AObj: TJPEGImage; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

procedure JPEGImage_SetOnChange(AObj: TJPEGImage; AEventId: NativeUInt); extdecl;
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
  JPEGImage_Create,
  JPEGImage_Free,
  JPEGImage_Assign,
  JPEGImage_LoadFromStream,
  JPEGImage_SaveToStream,
  JPEGImage_Equals,
  JPEGImage_LoadFromFile,
  JPEGImage_SaveToFile,
  JPEGImage_SetSize,
  JPEGImage_GetNamePath,
  JPEGImage_ClassName,
  JPEGImage_GetHashCode,
  JPEGImage_ToString,
  JPEGImage_GetPixelFormat,
  JPEGImage_SetPixelFormat,
  JPEGImage_GetPerformance,
  JPEGImage_SetPerformance,
  JPEGImage_GetCanvas,
  JPEGImage_GetEmpty,
  JPEGImage_GetHeight,
  JPEGImage_SetHeight,
  JPEGImage_GetModified,
  JPEGImage_SetModified,
  JPEGImage_GetPaletteModified,
  JPEGImage_SetPaletteModified,
  JPEGImage_GetTransparent,
  JPEGImage_SetTransparent,
  JPEGImage_GetWidth,
  JPEGImage_SetWidth,
  JPEGImage_SetOnChange;

