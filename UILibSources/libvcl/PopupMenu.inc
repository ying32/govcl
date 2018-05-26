
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function PopupMenu_Create(AOwner: TComponent): TPopupMenu; stdcall;
begin
  Result :=  TPopupMenu.Create(AOwner);
end;

procedure PopupMenu_Free(AObj: TPopupMenu); stdcall;
begin
  AObj.Free;
end;

procedure PopupMenu_CloseMenu(AObj: TPopupMenu); stdcall;
begin
  AObj.CloseMenu;
end;

procedure PopupMenu_Popup(AObj: TPopupMenu; X: Integer; Y: Integer); stdcall;
begin
  AObj.Popup(X, Y);
end;

function PopupMenu_FindComponent(AObj: TPopupMenu; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function PopupMenu_GetNamePath(AObj: TPopupMenu): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function PopupMenu_HasParent(AObj: TPopupMenu): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure PopupMenu_Assign(AObj: TPopupMenu; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function PopupMenu_ClassName(AObj: TPopupMenu): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function PopupMenu_Equals(AObj: TPopupMenu; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function PopupMenu_GetHashCode(AObj: TPopupMenu): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function PopupMenu_ToString(AObj: TPopupMenu): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function PopupMenu_GetAlignment(AObj: TPopupMenu): TPopupAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure PopupMenu_SetAlignment(AObj: TPopupMenu; AValue: TPopupAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function PopupMenu_GetAutoHotkeys(AObj: TPopupMenu): TMenuAutoFlag; stdcall;
begin
  Result :=  AObj.AutoHotkeys;
end;

procedure PopupMenu_SetAutoHotkeys(AObj: TPopupMenu; AValue: TMenuAutoFlag); stdcall;
begin
  AObj.AutoHotkeys := AValue;
end;

function PopupMenu_GetBiDiMode(AObj: TPopupMenu): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure PopupMenu_SetBiDiMode(AObj: TPopupMenu; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function PopupMenu_GetImages(AObj: TPopupMenu): TCustomImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure PopupMenu_SetImages(AObj: TPopupMenu; AValue: TCustomImageList); stdcall;
begin
  AObj.Images := AValue;
end;

procedure PopupMenu_SetOnChange(AObj: TPopupMenu; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, geMenuChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.MenuOnChange;
  TEventClass.Add(AObj, geMenuChange, AEventId);
end;

procedure PopupMenu_SetOnPopup(AObj: TPopupMenu; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnPopup := nil;
    TEventClass.Remove(AObj, gePopup);
    Exit;
  end;
  AObj.OnPopup := TEventClass.OnPopup;
  TEventClass.Add(AObj, gePopup, AEventId);
end;

function PopupMenu_GetHandle(AObj: TPopupMenu): HMENU; stdcall;
begin
  Result :=  AObj.Handle;
end;

function PopupMenu_GetWindowHandle(AObj: TPopupMenu): HWND; stdcall;
begin
  Result :=  AObj.WindowHandle;
end;

procedure PopupMenu_SetWindowHandle(AObj: TPopupMenu; AValue: HWND); stdcall;
begin
  AObj.WindowHandle := AValue;
end;

function PopupMenu_GetItems(AObj: TPopupMenu): TMenuItem; stdcall;
begin
  Result :=  AObj.Items;
end;

function PopupMenu_GetComponentCount(AObj: TPopupMenu): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function PopupMenu_GetComponentIndex(AObj: TPopupMenu): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure PopupMenu_SetComponentIndex(AObj: TPopupMenu; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function PopupMenu_GetOwner(AObj: TPopupMenu): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function PopupMenu_GetName(AObj: TPopupMenu): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure PopupMenu_SetName(AObj: TPopupMenu; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function PopupMenu_GetTag(AObj: TPopupMenu): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure PopupMenu_SetTag(AObj: TPopupMenu; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function PopupMenu_GetComponents(AObj: TPopupMenu; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  PopupMenu_Create,
  PopupMenu_Free,
  PopupMenu_CloseMenu,
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
  PopupMenu_GetAutoHotkeys,
  PopupMenu_SetAutoHotkeys,
  PopupMenu_GetBiDiMode,
  PopupMenu_SetBiDiMode,
  PopupMenu_GetImages,
  PopupMenu_SetImages,
  PopupMenu_SetOnChange,
  PopupMenu_SetOnPopup,
  PopupMenu_GetHandle,
  PopupMenu_GetWindowHandle,
  PopupMenu_SetWindowHandle,
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

