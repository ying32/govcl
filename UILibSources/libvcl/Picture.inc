
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Picture_Create: TPicture; stdcall;
begin
  Result :=  TPicture.Create;
end;

procedure Picture_Free(AObj: TPicture); stdcall;
begin
  AObj.Free;
end;

procedure Picture_LoadFromFile(AObj: TPicture; Filename: PChar); stdcall;
begin
  AObj.LoadFromFile(Filename);
end;

procedure Picture_SaveToFile(AObj: TPicture; Filename: PChar); stdcall;
begin
  AObj.SaveToFile(Filename);
end;

procedure Picture_LoadFromStream(AObj: TPicture; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure Picture_SaveToStream(AObj: TPicture; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

procedure Picture_Assign(AObj: TPicture; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Picture_GetNamePath(AObj: TPicture): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Picture_ClassName(AObj: TPicture): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Picture_Equals(AObj: TPicture; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Picture_GetHashCode(AObj: TPicture): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Picture_ToString(AObj: TPicture): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Picture_GetBitmap(AObj: TPicture): TBitmap; stdcall;
begin
  Result :=  AObj.Bitmap;
end;

procedure Picture_SetBitmap(AObj: TPicture; AValue: TBitmap); stdcall;
begin
  AObj.Bitmap := AValue;
end;

function Picture_GetGraphic(AObj: TPicture): TGraphic; stdcall;
begin
  Result :=  AObj.Graphic;
end;

procedure Picture_SetGraphic(AObj: TPicture; AValue: TGraphic); stdcall;
begin
  AObj.Graphic := AValue;
end;

function Picture_GetHeight(AObj: TPicture): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

function Picture_GetIcon(AObj: TPicture): TIcon; stdcall;
begin
  Result :=  AObj.Icon;
end;

procedure Picture_SetIcon(AObj: TPicture; AValue: TIcon); stdcall;
begin
  AObj.Icon := AValue;
end;

function Picture_GetWidth(AObj: TPicture): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Picture_SetOnChange(AObj: TPicture; AEventId: NativeUInt); stdcall;
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
  Picture_Create,
  Picture_Free,
  Picture_LoadFromFile,
  Picture_SaveToFile,
  Picture_LoadFromStream,
  Picture_SaveToStream,
  Picture_Assign,
  Picture_GetNamePath,
  Picture_ClassName,
  Picture_Equals,
  Picture_GetHashCode,
  Picture_ToString,
  Picture_GetBitmap,
  Picture_SetBitmap,
  Picture_GetGraphic,
  Picture_SetGraphic,
  Picture_GetHeight,
  Picture_GetIcon,
  Picture_SetIcon,
  Picture_GetWidth,
  Picture_SetOnChange;

