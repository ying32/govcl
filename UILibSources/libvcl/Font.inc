
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Font_Create: TFont; stdcall;
begin
  Result :=  TFont.Create;
end;

procedure Font_Free(AObj: TFont); stdcall;
begin
  AObj.Free;
end;

procedure Font_Assign(AObj: TFont; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Font_HandleAllocated(AObj: TFont): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

function Font_GetNamePath(AObj: TFont): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Font_ClassName(AObj: TFont): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Font_Equals(AObj: TFont; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Font_GetHashCode(AObj: TFont): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Font_ToString(AObj: TFont): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Font_GetHandle(AObj: TFont): HFONT; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure Font_SetHandle(AObj: TFont; AValue: HFONT); stdcall;
begin
  AObj.Handle := AValue;
end;

function Font_GetPixelsPerInch(AObj: TFont): Integer; stdcall;
begin
  Result :=  AObj.PixelsPerInch;
end;

procedure Font_SetPixelsPerInch(AObj: TFont; AValue: Integer); stdcall;
begin
  AObj.PixelsPerInch := AValue;
end;

function Font_GetCharset(AObj: TFont): TFontCharset; stdcall;
begin
  Result :=  AObj.Charset;
end;

procedure Font_SetCharset(AObj: TFont; AValue: TFontCharset); stdcall;
begin
  AObj.Charset := AValue;
end;

function Font_GetColor(AObj: TFont): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure Font_SetColor(AObj: TFont; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function Font_GetHeight(AObj: TFont): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Font_SetHeight(AObj: TFont; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Font_GetName(AObj: TFont): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Font_SetName(AObj: TFont; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Font_GetOrientation(AObj: TFont): Integer; stdcall;
begin
  Result :=  AObj.Orientation;
end;

procedure Font_SetOrientation(AObj: TFont; AValue: Integer); stdcall;
begin
  AObj.Orientation := AValue;
end;

function Font_GetPitch(AObj: TFont): TFontPitch; stdcall;
begin
  Result :=  AObj.Pitch;
end;

procedure Font_SetPitch(AObj: TFont; AValue: TFontPitch); stdcall;
begin
  AObj.Pitch := AValue;
end;

function Font_GetSize(AObj: TFont): Integer; stdcall;
begin
  Result :=  AObj.Size;
end;

procedure Font_SetSize(AObj: TFont; AValue: Integer); stdcall;
begin
  AObj.Size := AValue;
end;

function Font_GetStyle(AObj: TFont): TFontStyles; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure Font_SetStyle(AObj: TFont; AValue: TFontStyles); stdcall;
begin
  AObj.Style := AValue;
end;

function Font_GetQuality(AObj: TFont): TFontQuality; stdcall;
begin
  Result :=  AObj.Quality;
end;

procedure Font_SetQuality(AObj: TFont; AValue: TFontQuality); stdcall;
begin
  AObj.Quality := AValue;
end;

procedure Font_SetOnChange(AObj: TFont; AEventId: NativeUInt); stdcall;
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
  Font_Create,
  Font_Free,
  Font_Assign,
  Font_HandleAllocated,
  Font_GetNamePath,
  Font_ClassName,
  Font_Equals,
  Font_GetHashCode,
  Font_ToString,
  Font_GetHandle,
  Font_SetHandle,
  Font_GetPixelsPerInch,
  Font_SetPixelsPerInch,
  Font_GetCharset,
  Font_SetCharset,
  Font_GetColor,
  Font_SetColor,
  Font_GetHeight,
  Font_SetHeight,
  Font_GetName,
  Font_SetName,
  Font_GetOrientation,
  Font_SetOrientation,
  Font_GetPitch,
  Font_SetPitch,
  Font_GetSize,
  Font_SetSize,
  Font_GetStyle,
  Font_SetStyle,
  Font_GetQuality,
  Font_SetQuality,
  Font_SetOnChange;

