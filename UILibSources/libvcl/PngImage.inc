
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function PngImage_Create: TPngImage; stdcall;
begin
  Result :=  TPngImage.Create;
end;

procedure PngImage_Free(AObj: TPngImage); stdcall;
begin
  AObj.Free;
end;

procedure PngImage_Assign(AObj: TPngImage; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure PngImage_LoadFromStream(AObj: TPngImage; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure PngImage_SaveToStream(AObj: TPngImage; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

procedure PngImage_LoadFromResourceName(AObj: TPngImage; Instance: NativeUInt; Name: PChar); stdcall;
begin
  AObj.LoadFromResourceName(Instance, Name);
end;

procedure PngImage_LoadFromResourceID(AObj: TPngImage; Instance: NativeUInt; ResID: Integer); stdcall;
begin
{$IFDEF MSWINDOWS}
  AObj.LoadFromResourceID(Instance, ResID);
{$ENDIF MSWINDOWS}
end;

function PngImage_Equals(AObj: TPngImage; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure PngImage_LoadFromFile(AObj: TPngImage; Filename: PChar); stdcall;
begin
  AObj.LoadFromFile(Filename);
end;

procedure PngImage_SaveToFile(AObj: TPngImage; Filename: PChar); stdcall;
begin
  AObj.SaveToFile(Filename);
end;

procedure PngImage_SetSize(AObj: TPngImage; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetSize(AWidth, AHeight);
end;

function PngImage_GetNamePath(AObj: TPngImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function PngImage_ClassName(AObj: TPngImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function PngImage_GetHashCode(AObj: TPngImage): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function PngImage_ToString(AObj: TPngImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function PngImage_GetTransparentColor(AObj: TPngImage): TColor; stdcall;
begin
  Result :=  AObj.TransparentColor;
end;

procedure PngImage_SetTransparentColor(AObj: TPngImage; AValue: TColor); stdcall;
begin
  AObj.TransparentColor := AValue;
end;

function PngImage_GetCanvas(AObj: TPngImage): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function PngImage_GetWidth(AObj: TPngImage): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

function PngImage_GetHeight(AObj: TPngImage): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

function PngImage_GetMaxIdatSize(AObj: TPngImage): Integer; stdcall;
begin
  Result :=  AObj.MaxIdatSize;
end;

procedure PngImage_SetMaxIdatSize(AObj: TPngImage; AValue: Integer); stdcall;
begin
  AObj.MaxIdatSize := AValue;
end;

function PngImage_GetEmpty(AObj: TPngImage): LongBool; stdcall;
begin
  Result :=  AObj.Empty;
end;

function PngImage_GetCompressionLevel(AObj: TPngImage): TCompressionLevel; stdcall;
begin
  Result :=  AObj.CompressionLevel;
end;

procedure PngImage_SetCompressionLevel(AObj: TPngImage; AValue: TCompressionLevel); stdcall;
begin
  AObj.CompressionLevel := AValue;
end;

function PngImage_GetVersion(AObj: TPngImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.Version);
end;

function PngImage_GetModified(AObj: TPngImage): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure PngImage_SetModified(AObj: TPngImage; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function PngImage_GetPaletteModified(AObj: TPngImage): LongBool; stdcall;
begin
  Result :=  AObj.PaletteModified;
end;

procedure PngImage_SetPaletteModified(AObj: TPngImage; AValue: LongBool); stdcall;
begin
  AObj.PaletteModified := AValue;
end;

function PngImage_GetTransparent(AObj: TPngImage): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure PngImage_SetTransparent(AObj: TPngImage; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

procedure PngImage_SetOnChange(AObj: TPngImage; AEventId: NativeUInt); stdcall;
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
  PngImage_Create,
  PngImage_Free,
  PngImage_Assign,
  PngImage_LoadFromStream,
  PngImage_SaveToStream,
  PngImage_LoadFromResourceName,
  PngImage_LoadFromResourceID,
  PngImage_Equals,
  PngImage_LoadFromFile,
  PngImage_SaveToFile,
  PngImage_SetSize,
  PngImage_GetNamePath,
  PngImage_ClassName,
  PngImage_GetHashCode,
  PngImage_ToString,
  PngImage_GetTransparentColor,
  PngImage_SetTransparentColor,
  PngImage_GetCanvas,
  PngImage_GetWidth,
  PngImage_GetHeight,
  PngImage_GetMaxIdatSize,
  PngImage_SetMaxIdatSize,
  PngImage_GetEmpty,
  PngImage_GetCompressionLevel,
  PngImage_SetCompressionLevel,
  PngImage_GetVersion,
  PngImage_GetModified,
  PngImage_SetModified,
  PngImage_GetPaletteModified,
  PngImage_SetPaletteModified,
  PngImage_GetTransparent,
  PngImage_SetTransparent,
  PngImage_SetOnChange;

