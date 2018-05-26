
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function JPEGImage_Create: TJPEGImage; stdcall;
begin
  Result :=  TJPEGImage.Create;
end;

procedure JPEGImage_Free(AObj: TJPEGImage); stdcall;
begin
  AObj.Free;
end;

procedure JPEGImage_Assign(AObj: TJPEGImage; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure JPEGImage_LoadFromStream(AObj: TJPEGImage; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure JPEGImage_SaveToStream(AObj: TJPEGImage; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

function JPEGImage_Equals(AObj: TJPEGImage; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure JPEGImage_LoadFromFile(AObj: TJPEGImage; Filename: PChar); stdcall;
begin
  AObj.LoadFromFile(Filename);
end;

procedure JPEGImage_SaveToFile(AObj: TJPEGImage; Filename: PChar); stdcall;
begin
  AObj.SaveToFile(Filename);
end;

procedure JPEGImage_SetSize(AObj: TJPEGImage; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetSize(AWidth, AHeight);
end;

function JPEGImage_GetNamePath(AObj: TJPEGImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function JPEGImage_ClassName(AObj: TJPEGImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function JPEGImage_GetHashCode(AObj: TJPEGImage): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function JPEGImage_ToString(AObj: TJPEGImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function JPEGImage_GetPixelFormat(AObj: TJPEGImage): TJPEGPixelFormat; stdcall;
begin
  Result :=  AObj.PixelFormat;
end;

procedure JPEGImage_SetPixelFormat(AObj: TJPEGImage; AValue: TJPEGPixelFormat); stdcall;
begin
  AObj.PixelFormat := AValue;
end;

function JPEGImage_GetProgressiveDisplay(AObj: TJPEGImage): LongBool; stdcall;
begin
  Result :=  AObj.ProgressiveDisplay;
end;

procedure JPEGImage_SetProgressiveDisplay(AObj: TJPEGImage; AValue: LongBool); stdcall;
begin
  AObj.ProgressiveDisplay := AValue;
end;

function JPEGImage_GetPerformance(AObj: TJPEGImage): TJPEGPerformance; stdcall;
begin
  Result :=  AObj.Performance;
end;

procedure JPEGImage_SetPerformance(AObj: TJPEGImage; AValue: TJPEGPerformance); stdcall;
begin
  AObj.Performance := AValue;
end;

function JPEGImage_GetScale(AObj: TJPEGImage): TJPEGScale; stdcall;
begin
  Result :=  AObj.Scale;
end;

procedure JPEGImage_SetScale(AObj: TJPEGImage; AValue: TJPEGScale); stdcall;
begin
  AObj.Scale := AValue;
end;

function JPEGImage_GetSmoothing(AObj: TJPEGImage): LongBool; stdcall;
begin
  Result :=  AObj.Smoothing;
end;

procedure JPEGImage_SetSmoothing(AObj: TJPEGImage; AValue: LongBool); stdcall;
begin
  AObj.Smoothing := AValue;
end;

function JPEGImage_GetCanvas(AObj: TJPEGImage): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function JPEGImage_GetEmpty(AObj: TJPEGImage): LongBool; stdcall;
begin
  Result :=  AObj.Empty;
end;

function JPEGImage_GetHeight(AObj: TJPEGImage): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure JPEGImage_SetHeight(AObj: TJPEGImage; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function JPEGImage_GetModified(AObj: TJPEGImage): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure JPEGImage_SetModified(AObj: TJPEGImage; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function JPEGImage_GetPaletteModified(AObj: TJPEGImage): LongBool; stdcall;
begin
  Result :=  AObj.PaletteModified;
end;

procedure JPEGImage_SetPaletteModified(AObj: TJPEGImage; AValue: LongBool); stdcall;
begin
  AObj.PaletteModified := AValue;
end;

function JPEGImage_GetTransparent(AObj: TJPEGImage): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure JPEGImage_SetTransparent(AObj: TJPEGImage; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function JPEGImage_GetWidth(AObj: TJPEGImage): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure JPEGImage_SetWidth(AObj: TJPEGImage; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

procedure JPEGImage_SetOnChange(AObj: TJPEGImage; AEventId: NativeUInt); stdcall;
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
  JPEGImage_GetProgressiveDisplay,
  JPEGImage_SetProgressiveDisplay,
  JPEGImage_GetPerformance,
  JPEGImage_SetPerformance,
  JPEGImage_GetScale,
  JPEGImage_SetScale,
  JPEGImage_GetSmoothing,
  JPEGImage_SetSmoothing,
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

