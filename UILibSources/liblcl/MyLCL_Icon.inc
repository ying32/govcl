
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Icon_Create: TIcon; extdecl;
begin
  Result :=  TIcon.Create;
end;

procedure Icon_Free(AObj: TIcon); extdecl;
begin
  AObj.Free;
end;

procedure Icon_Assign(AObj: TIcon; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure Icon_LoadFromStream(AObj: TIcon; Stream: TStream); extdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure Icon_SaveToStream(AObj: TIcon; Stream: TStream); extdecl;
begin
  AObj.SaveToStream(Stream);
end;

procedure Icon_SetSize(AObj: TIcon; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetSize(AWidth, AHeight);
end;

procedure Icon_LoadFromResourceName(AObj: TIcon; Instance: NativeUInt; ResName: PChar); extdecl;
begin
  AObj.LoadFromResourceName(Instance, ResName);
end;

procedure Icon_LoadFromResourceID(AObj: TIcon; Instance: NativeUInt; ResID: Integer); extdecl;
begin
{$IFDEF MSWINDOWS}
  AObj.LoadFromResourceID(Instance, ResID);
{$ENDIF MSWINDOWS}
end;

function Icon_Equals(AObj: TIcon; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure Icon_LoadFromFile(AObj: TIcon; Filename: PChar); extdecl;
begin
  AObj.LoadFromFile(Filename);
end;

procedure Icon_SaveToFile(AObj: TIcon; Filename: PChar); extdecl;
begin
  AObj.SaveToFile(Filename);
end;

function Icon_GetNamePath(AObj: TIcon): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Icon_ClassName(AObj: TIcon): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Icon_GetHashCode(AObj: TIcon): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Icon_ToString(AObj: TIcon): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Icon_GetHandle(AObj: TIcon): HICON; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure Icon_SetHandle(AObj: TIcon; AValue: HICON); extdecl;
begin
  AObj.Handle := AValue;
end;

function Icon_GetEmpty(AObj: TIcon): LongBool; extdecl;
begin
  Result :=  AObj.Empty;
end;

function Icon_GetHeight(AObj: TIcon): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Icon_SetHeight(AObj: TIcon; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Icon_GetModified(AObj: TIcon): LongBool; extdecl;
begin
  Result :=  AObj.Modified;
end;

procedure Icon_SetModified(AObj: TIcon; AValue: LongBool); extdecl;
begin
  AObj.Modified := AValue;
end;

function Icon_GetPaletteModified(AObj: TIcon): LongBool; extdecl;
begin
  Result :=  AObj.PaletteModified;
end;

procedure Icon_SetPaletteModified(AObj: TIcon; AValue: LongBool); extdecl;
begin
  AObj.PaletteModified := AValue;
end;

function Icon_GetTransparent(AObj: TIcon): LongBool; extdecl;
begin
  Result :=  AObj.Transparent;
end;

procedure Icon_SetTransparent(AObj: TIcon; AValue: LongBool); extdecl;
begin
  AObj.Transparent := AValue;
end;

function Icon_GetWidth(AObj: TIcon): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Icon_SetWidth(AObj: TIcon; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

procedure Icon_SetOnChange(AObj: TIcon; AEventId: NativeUInt); extdecl;
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
  Icon_Create,
  Icon_Free,
  Icon_Assign,
  Icon_LoadFromStream,
  Icon_SaveToStream,
  Icon_SetSize,
  Icon_LoadFromResourceName,
  Icon_LoadFromResourceID,
  Icon_Equals,
  Icon_LoadFromFile,
  Icon_SaveToFile,
  Icon_GetNamePath,
  Icon_ClassName,
  Icon_GetHashCode,
  Icon_ToString,
  Icon_GetHandle,
  Icon_SetHandle,
  Icon_GetEmpty,
  Icon_GetHeight,
  Icon_SetHeight,
  Icon_GetModified,
  Icon_SetModified,
  Icon_GetPaletteModified,
  Icon_SetPaletteModified,
  Icon_GetTransparent,
  Icon_SetTransparent,
  Icon_GetWidth,
  Icon_SetWidth,
  Icon_SetOnChange;

