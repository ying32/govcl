
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function PopupMenu_Create(AOwner: TComponent): TPopupMenu; extdecl;
begin
  Result :=  TPopupMenu.Create(AOwner);
end;

procedure PopupMenu_Free(AObj: TPopupMenu); extdecl;
begin
  AObj.Free;
end;

procedure PopupMenu_Popup(AObj: TPopupMenu; X: Integer; Y: Integer); extdecl;
begin
  AObj.Popup(X, Y);
end;

function PopupMenu_FindComponent(AObj: TPopupMenu; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function PopupMenu_GetNamePath(AObj: TPopupMenu): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function PopupMenu_HasParent(AObj: TPopupMenu): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure PopupMenu_Assign(AObj: TPopupMenu; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function PopupMenu_ClassName(AObj: TPopupMenu): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function PopupMenu_Equals(AObj: TPopupMenu; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function PopupMenu_GetHashCode(AObj: TPopupMenu): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function PopupMenu_ToString(AObj: TPopupMenu): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function PopupMenu_GetAlignment(AObj: TPopupMenu): TPopupAlignment; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure PopupMenu_SetAlignment(AObj: TPopupMenu; AValue: TPopupAlignment); extdecl;
begin
  AObj.Alignment := AValue;
end;

function PopupMenu_GetBiDiMode(AObj: TPopupMenu): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure PopupMenu_SetBiDiMode(AObj: TPopupMenu; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function PopupMenu_GetImages(AObj: TPopupMenu): TCustomImageList; extdecl;
begin
  Result :=  AObj.Images;
end;

procedure PopupMenu_SetImages(AObj: TPopupMenu; AValue: TCustomImageList); extdecl;
begin
  AObj.Images := AValue;
end;

procedure PopupMenu_SetOnPopup(AObj: TPopupMenu; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnPopup := nil;
    EventClass.Remove(AObj, gePopup);
    Exit;
  end;
  AObj.OnPopup := @EventClass.OnPopup;
  EventClass.Add(AObj, gePopup, AEventId);
end;

function PopupMenu_GetHandle(AObj: TPopupMenu): HMENU; extdecl;
begin
  Result :=  AObj.Handle;
end;

function PopupMenu_GetItems(AObj: TPopupMenu): TMenuItem; extdecl;
begin
  Result :=  AObj.Items;
end;

function PopupMenu_GetComponentCount(AObj: TPopupMenu): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function PopupMenu_GetComponentIndex(AObj: TPopupMenu): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure PopupMenu_SetComponentIndex(AObj: TPopupMenu; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function PopupMenu_GetOwner(AObj: TPopupMenu): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function PopupMenu_GetName(AObj: TPopupMenu): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure PopupMenu_SetName(AObj: TPopupMenu; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function PopupMenu_GetTag(AObj: TPopupMenu): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure PopupMenu_SetTag(AObj: TPopupMenu; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function PopupMenu_GetComponents(AObj: TPopupMenu; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  PopupMenu_Create,
  PopupMenu_Free,
  PopupMenu_Popup,
  PopupMenu_FindComponent,
  PopupMenu_GetNamePath,
  PopupMenu_HasParent,
  PopupMenu_Assign,
  PopupMenu_ClassName,
  PopupMenu_Equals,
  PopupMenu_GetHashCode,
  PopupMenu_ToString,
  PopupMenu_GetAlignment,
  PopupMenu_SetAlignment,
  PopupMenu_GetBiDiMode,
  PopupMenu_SetBiDiMode,
  PopupMenu_GetImages,
  PopupMenu_SetImages,
  PopupMenu_SetOnPopup,
  PopupMenu_GetHandle,
  PopupMenu_GetItems,
  PopupMenu_GetComponentCount,
  PopupMenu_GetComponentIndex,
  PopupMenu_SetComponentIndex,
  PopupMenu_GetOwner,
  PopupMenu_GetName,
  PopupMenu_SetName,
  PopupMenu_GetTag,
  PopupMenu_SetTag,
  PopupMenu_GetComponents;

