
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function FontDialog_Create(AOwner: TComponent): TFontDialog; extdecl;
begin
  Result :=  TFontDialog.Create(AOwner);
end;

procedure FontDialog_Free(AObj: TFontDialog); extdecl;
begin
  AObj.Free;
end;

function FontDialog_Execute(AObj: TFontDialog): LongBool; extdecl;
begin
  Result :=  AObj.Execute;
end;

function FontDialog_FindComponent(AObj: TFontDialog; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function FontDialog_GetNamePath(AObj: TFontDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function FontDialog_HasParent(AObj: TFontDialog): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure FontDialog_Assign(AObj: TFontDialog; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function FontDialog_ClassName(AObj: TFontDialog): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function FontDialog_Equals(AObj: TFontDialog; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function FontDialog_GetHashCode(AObj: TFontDialog): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function FontDialog_ToString(AObj: TFontDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function FontDialog_GetFont(AObj: TFontDialog): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure FontDialog_SetFont(AObj: TFontDialog; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function FontDialog_GetOptions(AObj: TFontDialog): TFontDialogOptions; extdecl;
begin
  Result :=  AObj.Options;
end;

procedure FontDialog_SetOptions(AObj: TFontDialog; AValue: TFontDialogOptions); extdecl;
begin
  AObj.Options := AValue;
end;

function FontDialog_GetHandle(AObj: TFontDialog): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

procedure FontDialog_SetOnClose(AObj: TFontDialog; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnClose := nil;
    EventClass.Remove(AObj, geClose);
    Exit;
  end;
  AObj.OnClose := @EventClass.OnClose;
  EventClass.Add(AObj, geClose, AEventId);
end;

procedure FontDialog_SetOnShow(AObj: TFontDialog; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnShow := nil;
    EventClass.Remove(AObj, geShow);
    Exit;
  end;
  AObj.OnShow := @EventClass.OnShow;
  EventClass.Add(AObj, geShow, AEventId);
end;

function FontDialog_GetComponentCount(AObj: TFontDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function FontDialog_GetComponentIndex(AObj: TFontDialog): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure FontDialog_SetComponentIndex(AObj: TFontDialog; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function FontDialog_GetOwner(AObj: TFontDialog): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function FontDialog_GetName(AObj: TFontDialog): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure FontDialog_SetName(AObj: TFontDialog; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function FontDialog_GetTag(AObj: TFontDialog): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure FontDialog_SetTag(AObj: TFontDialog; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function FontDialog_GetComponents(AObj: TFontDialog; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  FontDialog_Create,
  FontDialog_Free,
  FontDialog_Execute,
  FontDialog_FindComponent,
  FontDialog_GetNamePath,
  FontDialog_HasParent,
  FontDialog_Assign,
  FontDialog_ClassName,
  FontDialog_Equals,
  FontDialog_GetHashCode,
  FontDialog_ToString,
  FontDialog_GetFont,
  FontDialog_SetFont,
  FontDialog_GetOptions,
  FontDialog_SetOptions,
  FontDialog_GetHandle,
  FontDialog_SetOnClose,
  FontDialog_SetOnShow,
  FontDialog_GetComponentCount,
  FontDialog_GetComponentIndex,
  FontDialog_SetComponentIndex,
  FontDialog_GetOwner,
  FontDialog_GetName,
  FontDialog_SetName,
  FontDialog_GetTag,
  FontDialog_SetTag,
  FontDialog_GetComponents;

