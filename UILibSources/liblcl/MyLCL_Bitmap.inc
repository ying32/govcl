
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Bitmap_Create: TBitmap; extdecl;
begin
  Result :=  TBitmap.Create;
end;

procedure Bitmap_Free(AObj: TBitmap); extdecl;
begin
  AObj.Free;
end;

procedure Bitmap_Assign(AObj: TBitmap; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure Bitmap_LoadFromStream(AObj: TBitmap; Stream: TStream); extdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure Bitmap_SaveToStream(AObj: TBitmap; Stream: TStream); extdecl;
begin
  AObj.SaveToStream(Stream);
end;

procedure Bitmap_SetSize(AObj: TBitmap; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetSize(AWidth, AHeight);
end;

procedure Bitmap_LoadFromResourceName(AObj: TBitmap; Instance: NativeUInt; ResName: PChar); extdecl;
begin
  AObj.LoadFromResourceName(Instance, ResName);
end;

procedure Bitmap_LoadFromResourceID(AObj: TBitmap; Instance: NativeUInt; ResID: Integer); extdecl;
begin
{$IFDEF MSWINDOWS}
  AObj.LoadFromResourceID(Instance, ResID);
{$ENDIF MSWINDOWS}
end;

function Bitmap_Equals(AObj: TBitmap; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure Bitmap_LoadFromFile(AObj: TBitmap; Filename: PChar); extdecl;
begin
  AObj.LoadFromFile(Filename);
end;

procedure Bitmap_SaveToFile(AObj: TBitmap; Filename: PChar); extdecl;
begin
  AObj.SaveToFile(Filename);
end;

function Bitmap_GetNamePath(AObj: TBitmap): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Bitmap_ClassName(AObj: TBitmap): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Bitmap_GetHashCode(AObj: TBitmap): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Bitmap_ToString(AObj: TBitmap): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Bitmap_GetCanvas(AObj: TBitmap): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function Bitmap_GetHandle(AObj: TBitmap): HBITMAP; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure Bitmap_SetHandle(AObj: TBitmap; AValue: HBITMAP); extdecl;
begin
  AObj.Handle := AValue;
end;

function Bitmap_GetPixelFormat(AObj: TBitmap): TPixelFormat; extdecl;
begin
  Result :=  AObj.PixelFormat;
end;

procedure Bitmap_SetPixelFormat(AObj: TBitmap; AValue: TPixelFormat); extdecl;
begin
  AObj.PixelFormat := AValue;
end;

function Bitmap_GetEmpty(AObj: TBitmap): LongBool; extdecl;
begin
  Result :=  AObj.Empty;
end;

function Bitmap_GetHeight(AObj: TBitmap): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Bitmap_SetHeight(AObj: TBitmap; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Bitmap_GetModified(AObj: TBitmap): LongBool; extdecl;
begin
  Result :=  AObj.Modified;
end;

procedure Bitmap_SetModified(AObj: TBitmap; AValue: LongBool); extdecl;
begin
  AObj.Modified := AValue;
end;

function Bitmap_GetPaletteModified(AObj: TBitmap): LongBool; extdecl;
begin
  Result :=  AObj.PaletteModified;
end;

procedure Bitmap_SetPaletteModified(AObj: TBitmap; AValue: LongBool); extdecl;
begin
  AObj.PaletteModified := AValue;
end;

function Bitmap_GetTransparent(AObj: TBitmap): LongBool; extdecl;
begin
  Result :=  AObj.Transparent;
end;

procedure Bitmap_SetTransparent(AObj: TBitmap; AValue: LongBool); extdecl;
begin
  AObj.Transparent := AValue;
end;

function Bitmap_GetWidth(AObj: TBitmap): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Bitmap_SetWidth(AObj: TBitmap; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

procedure Bitmap_SetOnChange(AObj: TBitmap; AEventId: NativeUInt); extdecl;
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

function Bitmap_GetScanLine(AObj: TBitmap; Row: Integer): Pointer; extdecl;
begin
  Result :=  AObj.ScanLine[Row];
end;


exports
  Bitmap_Create,
  Bitmap_Free,
  Bitmap_Assign,
  Bitmap_LoadFromStream,
  Bitmap_SaveToStream,
  Bitmap_SetSize,
  Bitmap_LoadFromResourceName,
  Bitmap_LoadFromResourceID,
  Bitmap_Equals,
  Bitmap_LoadFromFile,
  Bitmap_SaveToFile,
  Bitmap_GetNamePath,
  Bitmap_ClassName,
  Bitmap_GetHashCode,
  Bitmap_ToString,
  Bitmap_GetCanvas,
  Bitmap_GetHandle,
  Bitmap_SetHandle,
  Bitmap_GetPixelFormat,
  Bitmap_SetPixelFormat,
  Bitmap_GetEmpty,
  Bitmap_GetHeight,
  Bitmap_SetHeight,
  Bitmap_GetModified,
  Bitmap_SetModified,
  Bitmap_GetPaletteModified,
  Bitmap_SetPaletteModified,
  Bitmap_GetTransparent,
  Bitmap_SetTransparent,
  Bitmap_GetWidth,
  Bitmap_SetWidth,
  Bitmap_SetOnChange,
  Bitmap_GetScanLine;

