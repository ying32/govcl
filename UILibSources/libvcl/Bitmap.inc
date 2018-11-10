
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Bitmap_Create: TBitmap; stdcall;
begin
  Result :=  TBitmap.Create;
end;

procedure Bitmap_Free(AObj: TBitmap); stdcall;
begin
  AObj.Free;
end;

procedure Bitmap_Assign(AObj: TBitmap; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure Bitmap_FreeImage(AObj: TBitmap); stdcall;
begin
  AObj.FreeImage;
end;

function Bitmap_HandleAllocated(AObj: TBitmap): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure Bitmap_LoadFromStream(AObj: TBitmap; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure Bitmap_SaveToStream(AObj: TBitmap; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

procedure Bitmap_SetSize(AObj: TBitmap; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetSize(AWidth, AHeight);
end;

procedure Bitmap_LoadFromResourceName(AObj: TBitmap; Instance: NativeUInt; ResName: PChar); stdcall;
begin
  AObj.LoadFromResourceName(Instance, ResName);
end;

procedure Bitmap_LoadFromResourceID(AObj: TBitmap; Instance: NativeUInt; ResID: Integer); stdcall;
begin
{$IFDEF MSWINDOWS}
  AObj.LoadFromResourceID(Instance, ResID);
{$ENDIF MSWINDOWS}
end;

function Bitmap_Equals(AObj: TBitmap; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure Bitmap_LoadFromFile(AObj: TBitmap; Filename: PChar); stdcall;
begin
  AObj.LoadFromFile(Filename);
end;

procedure Bitmap_SaveToFile(AObj: TBitmap; Filename: PChar); stdcall;
begin
  AObj.SaveToFile(Filename);
end;

function Bitmap_GetNamePath(AObj: TBitmap): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Bitmap_DisposeOf(AObj: TBitmap); stdcall;
begin
  AObj.DisposeOf;
end;

function Bitmap_ClassType(AObj: TBitmap): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function Bitmap_ClassName(AObj: TBitmap): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Bitmap_InstanceSize(AObj: TBitmap): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function Bitmap_InheritsFrom(AObj: TBitmap; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function Bitmap_GetHashCode(AObj: TBitmap): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Bitmap_ToString(AObj: TBitmap): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Bitmap_GetCanvas(AObj: TBitmap): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function Bitmap_GetHandle(AObj: TBitmap): HBITMAP; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure Bitmap_SetHandle(AObj: TBitmap; AValue: HBITMAP); stdcall;
begin
  AObj.Handle := AValue;
end;

function Bitmap_GetHandleType(AObj: TBitmap): TBitmapHandleType; stdcall;
begin
  Result :=  AObj.HandleType;
end;

procedure Bitmap_SetHandleType(AObj: TBitmap; AValue: TBitmapHandleType); stdcall;
begin
  AObj.HandleType := AValue;
end;

function Bitmap_GetMaskHandle(AObj: TBitmap): HBITMAP; stdcall;
begin
  Result :=  AObj.MaskHandle;
end;

procedure Bitmap_SetMaskHandle(AObj: TBitmap; AValue: HBITMAP); stdcall;
begin
  AObj.MaskHandle := AValue;
end;

function Bitmap_GetPixelFormat(AObj: TBitmap): TPixelFormat; stdcall;
begin
  Result :=  AObj.PixelFormat;
end;

procedure Bitmap_SetPixelFormat(AObj: TBitmap; AValue: TPixelFormat); stdcall;
begin
  AObj.PixelFormat := AValue;
end;

function Bitmap_GetTransparentColor(AObj: TBitmap): TColor; stdcall;
begin
  Result :=  AObj.TransparentColor;
end;

procedure Bitmap_SetTransparentColor(AObj: TBitmap; AValue: TColor); stdcall;
begin
  AObj.TransparentColor := AValue;
end;

function Bitmap_GetEmpty(AObj: TBitmap): LongBool; stdcall;
begin
  Result :=  AObj.Empty;
end;

function Bitmap_GetHeight(AObj: TBitmap): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Bitmap_SetHeight(AObj: TBitmap; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Bitmap_GetModified(AObj: TBitmap): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure Bitmap_SetModified(AObj: TBitmap; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function Bitmap_GetPalette(AObj: TBitmap): HPALETTE; stdcall;
begin
  Result :=  AObj.Palette;
end;

procedure Bitmap_SetPalette(AObj: TBitmap; AValue: HPALETTE); stdcall;
begin
  AObj.Palette := AValue;
end;

function Bitmap_GetPaletteModified(AObj: TBitmap): LongBool; stdcall;
begin
  Result :=  AObj.PaletteModified;
end;

procedure Bitmap_SetPaletteModified(AObj: TBitmap; AValue: LongBool); stdcall;
begin
  AObj.PaletteModified := AValue;
end;

function Bitmap_GetTransparent(AObj: TBitmap): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure Bitmap_SetTransparent(AObj: TBitmap; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function Bitmap_GetWidth(AObj: TBitmap): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Bitmap_SetWidth(AObj: TBitmap; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

procedure Bitmap_SetOnChange(AObj: TBitmap; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, @TEventClass.OnChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, @TEventClass.OnChange, AEventId);
end;

function Bitmap_GetScanLine(AObj: TBitmap; Row: Integer): Pointer; stdcall;
begin
  Result :=  AObj.ScanLine[Row];
end;

function Bitmap_StaticClassType: TClass; stdcall;
begin
  Result := TBitmap;
end;

exports
  Bitmap_Create,
  Bitmap_Free,
  Bitmap_Assign,
  Bitmap_FreeImage,
  Bitmap_HandleAllocated,
  Bitmap_LoadFromStream,
  Bitmap_SaveToStream,
  Bitmap_SetSize,
  Bitmap_LoadFromResourceName,
  Bitmap_LoadFromResourceID,
  Bitmap_Equals,
  Bitmap_LoadFromFile,
  Bitmap_SaveToFile,
  Bitmap_GetNamePath,
  Bitmap_DisposeOf,
  Bitmap_ClassType,
  Bitmap_ClassName,
  Bitmap_InstanceSize,
  Bitmap_InheritsFrom,
  Bitmap_GetHashCode,
  Bitmap_ToString,
  Bitmap_GetCanvas,
  Bitmap_GetHandle,
  Bitmap_SetHandle,
  Bitmap_GetHandleType,
  Bitmap_SetHandleType,
  Bitmap_GetMaskHandle,
  Bitmap_SetMaskHandle,
  Bitmap_GetPixelFormat,
  Bitmap_SetPixelFormat,
  Bitmap_GetTransparentColor,
  Bitmap_SetTransparentColor,
  Bitmap_GetEmpty,
  Bitmap_GetHeight,
  Bitmap_SetHeight,
  Bitmap_GetModified,
  Bitmap_SetModified,
  Bitmap_GetPalette,
  Bitmap_SetPalette,
  Bitmap_GetPaletteModified,
  Bitmap_SetPaletteModified,
  Bitmap_GetTransparent,
  Bitmap_SetTransparent,
  Bitmap_GetWidth,
  Bitmap_SetWidth,
  Bitmap_SetOnChange,
  Bitmap_GetScanLine,
  Bitmap_StaticClassType;

