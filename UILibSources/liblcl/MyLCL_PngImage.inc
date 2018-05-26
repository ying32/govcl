
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function PngImage_Create: TPngImage; extdecl;
begin
  Result :=  TPngImage.Create;
end;

procedure PngImage_Free(AObj: TPngImage); extdecl;
begin
  AObj.Free;
end;

procedure PngImage_Assign(AObj: TPngImage; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure PngImage_LoadFromStream(AObj: TPngImage; Stream: TStream); extdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure PngImage_SaveToStream(AObj: TPngImage; Stream: TStream); extdecl;
begin
  AObj.SaveToStream(Stream);
end;

procedure PngImage_LoadFromResourceName(AObj: TPngImage; Instance: NativeUInt; Name: PChar); extdecl;
begin
  AObj.LoadFromResourceName(Instance, Name);
end;

procedure PngImage_LoadFromResourceID(AObj: TPngImage; Instance: NativeUInt; ResID: Integer); extdecl;
begin
{$IFDEF MSWINDOWS}
  AObj.LoadFromResourceID(Instance, ResID);
{$ENDIF MSWINDOWS}
end;

function PngImage_Equals(AObj: TPngImage; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure PngImage_LoadFromFile(AObj: TPngImage; Filename: PChar); extdecl;
begin
  AObj.LoadFromFile(Filename);
end;

procedure PngImage_SaveToFile(AObj: TPngImage; Filename: PChar); extdecl;
begin
  AObj.SaveToFile(Filename);
end;

procedure PngImage_SetSize(AObj: TPngImage; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetSize(AWidth, AHeight);
end;

function PngImage_GetNamePath(AObj: TPngImage): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function PngImage_ClassName(AObj: TPngImage): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function PngImage_GetHashCode(AObj: TPngImage): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function PngImage_ToString(AObj: TPngImage): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function PngImage_GetCanvas(AObj: TPngImage): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function PngImage_GetWidth(AObj: TPngImage): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

function PngImage_GetHeight(AObj: TPngImage): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

function PngImage_GetEmpty(AObj: TPngImage): LongBool; extdecl;
begin
  Result :=  AObj.Empty;
end;

function PngImage_GetModified(AObj: TPngImage): LongBool; extdecl;
begin
  Result :=  AObj.Modified;
end;

procedure PngImage_SetModified(AObj: TPngImage; AValue: LongBool); extdecl;
begin
  AObj.Modified := AValue;
end;

function PngImage_GetPaletteModified(AObj: TPngImage): LongBool; extdecl;
begin
  Result :=  AObj.PaletteModified;
end;

procedure PngImage_SetPaletteModified(AObj: TPngImage; AValue: LongBool); extdecl;
begin
  AObj.PaletteModified := AValue;
end;

function PngImage_GetTransparent(AObj: TPngImage): LongBool; extdecl;
begin
  Result :=  AObj.Transparent;
end;

procedure PngImage_SetTransparent(AObj: TPngImage; AValue: LongBool); extdecl;
begin
  AObj.Transparent := AValue;
end;

procedure PngImage_SetOnChange(AObj: TPngImage; AEventId: NativeUInt); extdecl;
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
  PngImage_GetCanvas,
  PngImage_GetWidth,
  PngImage_GetHeight,
  PngImage_GetEmpty,
  PngImage_GetModified,
  PngImage_SetModified,
  PngImage_GetPaletteModified,
  PngImage_SetPaletteModified,
  PngImage_GetTransparent,
  PngImage_SetTransparent,
  PngImage_SetOnChange;

