
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Icon_Create: TIcon; stdcall;
begin
  Result :=  TIcon.Create;
end;

procedure Icon_Free(AObj: TIcon); stdcall;
begin
  AObj.Free;
end;

procedure Icon_Assign(AObj: TIcon; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Icon_HandleAllocated(AObj: TIcon): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure Icon_LoadFromStream(AObj: TIcon; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure Icon_SaveToStream(AObj: TIcon; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

procedure Icon_SetSize(AObj: TIcon; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetSize(AWidth, AHeight);
end;

procedure Icon_LoadFromResourceName(AObj: TIcon; Instance: NativeUInt; ResName: PChar); stdcall;
begin
  AObj.LoadFromResourceName(Instance, ResName);
end;

procedure Icon_LoadFromResourceID(AObj: TIcon; Instance: NativeUInt; ResID: Integer); stdcall;
begin
{$IFDEF MSWINDOWS}
  AObj.LoadFromResourceID(Instance, ResID);
{$ENDIF MSWINDOWS}
end;

function Icon_Equals(AObj: TIcon; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure Icon_LoadFromFile(AObj: TIcon; Filename: PChar); stdcall;
begin
  AObj.LoadFromFile(Filename);
end;

procedure Icon_SaveToFile(AObj: TIcon; Filename: PChar); stdcall;
begin
  AObj.SaveToFile(Filename);
end;

function Icon_GetNamePath(AObj: TIcon): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Icon_ClassName(AObj: TIcon): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Icon_GetHashCode(AObj: TIcon): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Icon_ToString(AObj: TIcon): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Icon_GetHandle(AObj: TIcon): HICON; stdcall;
begin
  Result :=  AObj.Handle;
end;

procedure Icon_SetHandle(AObj: TIcon; AValue: HICON); stdcall;
begin
  AObj.Handle := AValue;
end;

function Icon_GetEmpty(AObj: TIcon): LongBool; stdcall;
begin
  Result :=  AObj.Empty;
end;

function Icon_GetHeight(AObj: TIcon): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Icon_SetHeight(AObj: TIcon; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Icon_GetModified(AObj: TIcon): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure Icon_SetModified(AObj: TIcon; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function Icon_GetPaletteModified(AObj: TIcon): LongBool; stdcall;
begin
  Result :=  AObj.PaletteModified;
end;

procedure Icon_SetPaletteModified(AObj: TIcon; AValue: LongBool); stdcall;
begin
  AObj.PaletteModified := AValue;
end;

function Icon_GetTransparent(AObj: TIcon): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure Icon_SetTransparent(AObj: TIcon; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function Icon_GetWidth(AObj: TIcon): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Icon_SetWidth(AObj: TIcon; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

procedure Icon_SetOnChange(AObj: TIcon; AEventId: NativeUInt); stdcall;
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
  Icon_Create,
  Icon_Free,
  Icon_Assign,
  Icon_HandleAllocated,
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

