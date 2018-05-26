
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function MenuItem_Create(AOwner: TComponent): TMenuItem; stdcall;
begin
  Result :=  TMenuItem.Create(AOwner);
end;

procedure MenuItem_Free(AObj: TMenuItem); stdcall;
begin
  AObj.Free;
end;

procedure MenuItem_Insert(AObj: TMenuItem; Index: Integer; Item: TMenuItem); stdcall;
begin
  AObj.Insert(Index, Item);
end;

procedure MenuItem_Delete(AObj: TMenuItem; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure MenuItem_Clear(AObj: TMenuItem); stdcall;
begin
  AObj.Clear;
end;

procedure MenuItem_Click(AObj: TMenuItem); stdcall;
begin
  AObj.Click;
end;

function MenuItem_IndexOf(AObj: TMenuItem; Item: TMenuItem): Integer; stdcall;
begin
  Result :=  AObj.IndexOf(Item);
end;

function MenuItem_HasParent(AObj: TMenuItem): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure MenuItem_Add(AObj: TMenuItem; Item: TMenuItem); stdcall;
begin
  AObj.Add(Item);
end;

function MenuItem_FindComponent(AObj: TMenuItem; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function MenuItem_GetNamePath(AObj: TMenuItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure MenuItem_Assign(AObj: TMenuItem; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function MenuItem_ClassName(AObj: TMenuItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function MenuItem_Equals(AObj: TMenuItem; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function MenuItem_GetHashCode(AObj: TMenuItem): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function MenuItem_ToString(AObj: TMenuItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function MenuItem_GetHandle(AObj: TMenuItem): HMENU; stdcall;
begin
  Result :=  AObj.Handle;
end;

function MenuItem_GetCount(AObj: TMenuItem): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function MenuItem_GetParent(AObj: TMenuItem): TMenuItem; stdcall;
begin
  Result :=  AObj.Parent;
end;

function MenuItem_GetAction(AObj: TMenuItem): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure MenuItem_SetAction(AObj: TMenuItem; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function MenuItem_GetAutoHotkeys(AObj: TMenuItem): TMenuItemAutoFlag; stdcall;
begin
  Result :=  AObj.AutoHotkeys;
end;

procedure MenuItem_SetAutoHotkeys(AObj: TMenuItem; AValue: TMenuItemAutoFlag); stdcall;
begin
  AObj.AutoHotkeys := AValue;
end;

function MenuItem_GetBitmap(AObj: TMenuItem): TBitmap; stdcall;
begin
  Result :=  AObj.Bitmap;
end;

procedure MenuItem_SetBitmap(AObj: TMenuItem; AValue: TBitmap); stdcall;
begin
  AObj.Bitmap := AValue;
end;

function MenuItem_GetCaption(AObj: TMenuItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure MenuItem_SetCaption(AObj: TMenuItem; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function MenuItem_GetChecked(AObj: TMenuItem): LongBool; stdcall;
begin
  Result :=  AObj.Checked;
end;

procedure MenuItem_SetChecked(AObj: TMenuItem; AValue: LongBool); stdcall;
begin
  AObj.Checked := AValue;
end;

function MenuItem_GetDefault(AObj: TMenuItem): LongBool; stdcall;
begin
  Result :=  AObj.Default;
end;

procedure MenuItem_SetDefault(AObj: TMenuItem; AValue: LongBool); stdcall;
begin
  AObj.Default := AValue;
end;

function MenuItem_GetEnabled(AObj: TMenuItem): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure MenuItem_SetEnabled(AObj: TMenuItem; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function MenuItem_GetGroupIndex(AObj: TMenuItem): Byte; stdcall;
begin
  Result :=  AObj.GroupIndex;
end;

procedure MenuItem_SetGroupIndex(AObj: TMenuItem; AValue: Byte); stdcall;
begin
  AObj.GroupIndex := AValue;
end;

function MenuItem_GetHint(AObj: TMenuItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure MenuItem_SetHint(AObj: TMenuItem; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function MenuItem_GetImageIndex(AObj: TMenuItem): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure MenuItem_SetImageIndex(AObj: TMenuItem; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function MenuItem_GetShortCut(AObj: TMenuItem): TShortCut; stdcall;
begin
  Result :=  AObj.ShortCut;
end;

procedure MenuItem_SetShortCut(AObj: TMenuItem; AValue: TShortCut); stdcall;
begin
  AObj.ShortCut := AValue;
end;

function MenuItem_GetVisible(AObj: TMenuItem): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure MenuItem_SetVisible(AObj: TMenuItem; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

procedure MenuItem_SetOnClick(AObj: TMenuItem; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    TEventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, geClick, AEventId);
end;

procedure MenuItem_SetOnDrawItem(AObj: TMenuItem; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDrawItem := nil;
    TEventClass.Remove(AObj, geMenuItemDrawItem);
    Exit;
  end;
  AObj.OnDrawItem := TEventClass.MenuItemOnDrawItem;
  TEventClass.Add(AObj, geMenuItemDrawItem, AEventId);
end;

function MenuItem_GetComponentCount(AObj: TMenuItem): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function MenuItem_GetComponentIndex(AObj: TMenuItem): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure MenuItem_SetComponentIndex(AObj: TMenuItem; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function MenuItem_GetOwner(AObj: TMenuItem): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function MenuItem_GetName(AObj: TMenuItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure MenuItem_SetName(AObj: TMenuItem; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function MenuItem_GetTag(AObj: TMenuItem): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure MenuItem_SetTag(AObj: TMenuItem; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function MenuItem_GetItems(AObj: TMenuItem; Index: Integer): TMenuItem; stdcall;
begin
  Result :=  AObj.Items[Index];
end;

function MenuItem_GetComponents(AObj: TMenuItem; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  MenuItem_Create,
  MenuItem_Free,
  MenuItem_Insert,
  MenuItem_Delete,
  MenuItem_Clear,
  MenuItem_Click,
  MenuItem_IndexOf,
  MenuItem_HasParent,
  MenuItem_Add,
  MenuItem_FindComponent,
  MenuItem_GetNamePath,
  MenuItem_Assign,
  MenuItem_ClassName,
  MenuItem_Equals,
  MenuItem_GetHashCode,
  MenuItem_ToString,
  MenuItem_GetHandle,
  MenuItem_GetCount,
  MenuItem_GetParent,
  MenuItem_GetAction,
  MenuItem_SetAction,
  MenuItem_GetAutoHotkeys,
  MenuItem_SetAutoHotkeys,
  MenuItem_GetBitmap,
  MenuItem_SetBitmap,
  MenuItem_GetCaption,
  MenuItem_SetCaption,
  MenuItem_GetChecked,
  MenuItem_SetChecked,
  MenuItem_GetDefault,
  MenuItem_SetDefault,
  MenuItem_GetEnabled,
  MenuItem_SetEnabled,
  MenuItem_GetGroupIndex,
  MenuItem_SetGroupIndex,
  MenuItem_GetHint,
  MenuItem_SetHint,
  MenuItem_GetImageIndex,
  MenuItem_SetImageIndex,
  MenuItem_GetShortCut,
  MenuItem_SetShortCut,
  MenuItem_GetVisible,
  MenuItem_SetVisible,
  MenuItem_SetOnClick,
  MenuItem_SetOnDrawItem,
  MenuItem_GetComponentCount,
  MenuItem_GetComponentIndex,
  MenuItem_SetComponentIndex,
  MenuItem_GetOwner,
  MenuItem_GetName,
  MenuItem_SetName,
  MenuItem_GetTag,
  MenuItem_SetTag,
  MenuItem_GetItems,
  MenuItem_GetComponents;

