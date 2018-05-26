
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Font_Create: TFont; extdecl;
begin
  Result :=  TFont.Create;
end;

procedure Font_Free(AObj: TFont); extdecl;
begin
  AObj.Free;
end;

procedure Font_Assign(AObj: TFont; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Font_GetNamePath(AObj: TFont): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Font_ClassName(AObj: TFont): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Font_Equals(AObj: TFont; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Font_GetHashCode(AObj: TFont): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Font_ToString(AObj: TFont): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Font_GetHandle(AObj: TFont): HFONT; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure Font_SetHandle(AObj: TFont; AValue: HFONT); extdecl;
begin
  AObj.Handle := AValue;
end;

function Font_GetPixelsPerInch(AObj: TFont): Integer; extdecl;
begin
  Result :=  AObj.PixelsPerInch;
end;

procedure Font_SetPixelsPerInch(AObj: TFont; AValue: Integer); extdecl;
begin
  AObj.PixelsPerInch := AValue;
end;

function Font_GetCharset(AObj: TFont): TFontCharset; extdecl;
begin
  Result :=  AObj.Charset;
end;

procedure Font_SetCharset(AObj: TFont; AValue: TFontCharset); extdecl;
begin
  AObj.Charset := AValue;
end;

function Font_GetColor(AObj: TFont): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure Font_SetColor(AObj: TFont; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function Font_GetHeight(AObj: TFont): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Font_SetHeight(AObj: TFont; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Font_GetName(AObj: TFont): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Font_SetName(AObj: TFont; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Font_GetOrientation(AObj: TFont): Integer; extdecl;
begin
  Result :=  AObj.Orientation;
end;

procedure Font_SetOrientation(AObj: TFont; AValue: Integer); extdecl;
begin
  AObj.Orientation := AValue;
end;

function Font_GetPitch(AObj: TFont): TFontPitch; extdecl;
begin
  Result :=  AObj.Pitch;
end;

procedure Font_SetPitch(AObj: TFont; AValue: TFontPitch); extdecl;
begin
  AObj.Pitch := AValue;
end;

function Font_GetSize(AObj: TFont): Integer; extdecl;
begin
  Result :=  AObj.Size;
end;

procedure Font_SetSize(AObj: TFont; AValue: Integer); extdecl;
begin
  AObj.Size := AValue;
end;

function Font_GetStyle(AObj: TFont): TFontStyles; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure Font_SetStyle(AObj: TFont; AValue: TFontStyles); extdecl;
begin
  AObj.Style := AValue;
end;

function Font_GetQuality(AObj: TFont): TFontQuality; extdecl;
begin
  Result :=  AObj.Quality;
end;

procedure Font_SetQuality(AObj: TFont; AValue: TFontQuality); extdecl;
begin
  AObj.Quality := AValue;
end;

procedure Font_SetOnChange(AObj: TFont; AEventId: NativeUInt); extdecl;
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
  Font_Create,
  Font_Free,
  Font_Assign,
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

