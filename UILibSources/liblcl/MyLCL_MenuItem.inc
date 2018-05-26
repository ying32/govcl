
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function MenuItem_Create(AOwner: TComponent): TMenuItem; extdecl;
begin
  Result :=  TMenuItem.Create(AOwner);
end;

procedure MenuItem_Free(AObj: TMenuItem); extdecl;
begin
  AObj.Free;
end;

procedure MenuItem_Insert(AObj: TMenuItem; Index: Integer; Item: TMenuItem); extdecl;
begin
  AObj.Insert(Index, Item);
end;

procedure MenuItem_Delete(AObj: TMenuItem; Index: Integer); extdecl;
begin
  AObj.Delete(Index);
end;

procedure MenuItem_Clear(AObj: TMenuItem); extdecl;
begin
  AObj.Clear;
end;

procedure MenuItem_Click(AObj: TMenuItem); extdecl;
begin
  AObj.Click;
end;

function MenuItem_IndexOf(AObj: TMenuItem; Item: TMenuItem): Integer; extdecl;
begin
  Result :=  AObj.IndexOf(Item);
end;

function MenuItem_HasParent(AObj: TMenuItem): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure MenuItem_Add(AObj: TMenuItem; Item: TMenuItem); extdecl;
begin
  AObj.Add(Item);
end;

function MenuItem_FindComponent(AObj: TMenuItem; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function MenuItem_GetNamePath(AObj: TMenuItem): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure MenuItem_Assign(AObj: TMenuItem; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function MenuItem_ClassName(AObj: TMenuItem): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function MenuItem_Equals(AObj: TMenuItem; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function MenuItem_GetHashCode(AObj: TMenuItem): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function MenuItem_ToString(AObj: TMenuItem): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function MenuItem_GetHandle(AObj: TMenuItem): HMENU; extdecl;
begin
  Result :=  AObj.Handle;
end;

function MenuItem_GetCount(AObj: TMenuItem): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

function MenuItem_GetParent(AObj: TMenuItem): TMenuItem; extdecl;
begin
  Result :=  AObj.Parent;
end;

function MenuItem_GetAction(AObj: TMenuItem): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure MenuItem_SetAction(AObj: TMenuItem; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function MenuItem_GetBitmap(AObj: TMenuItem): TBitmap; extdecl;
begin
  Result :=  AObj.Bitmap;
end;

procedure MenuItem_SetBitmap(AObj: TMenuItem; AValue: TBitmap); extdecl;
begin
  AObj.Bitmap := AValue;
end;

function MenuItem_GetCaption(AObj: TMenuItem): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure MenuItem_SetCaption(AObj: TMenuItem; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function MenuItem_GetChecked(AObj: TMenuItem): LongBool; extdecl;
begin
  Result :=  AObj.Checked;
end;

procedure MenuItem_SetChecked(AObj: TMenuItem; AValue: LongBool); extdecl;
begin
  AObj.Checked := AValue;
end;

function MenuItem_GetDefault(AObj: TMenuItem): LongBool; extdecl;
begin
  Result :=  AObj.Default;
end;

procedure MenuItem_SetDefault(AObj: TMenuItem; AValue: LongBool); extdecl;
begin
  AObj.Default := AValue;
end;

function MenuItem_GetEnabled(AObj: TMenuItem): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure MenuItem_SetEnabled(AObj: TMenuItem; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function MenuItem_GetGroupIndex(AObj: TMenuItem): Byte; extdecl;
begin
  Result :=  AObj.GroupIndex;
end;

procedure MenuItem_SetGroupIndex(AObj: TMenuItem; AValue: Byte); extdecl;
begin
  AObj.GroupIndex := AValue;
end;

function MenuItem_GetHint(AObj: TMenuItem): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure MenuItem_SetHint(AObj: TMenuItem; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function MenuItem_GetImageIndex(AObj: TMenuItem): TImageIndex; extdecl;
begin
  Result :=  AObj.ImageIndex;
end;

procedure MenuItem_SetImageIndex(AObj: TMenuItem; AValue: TImageIndex); extdecl;
begin
  AObj.ImageIndex := AValue;
end;

function MenuItem_GetShortCut(AObj: TMenuItem): TShortCut; extdecl;
begin
  Result :=  AObj.ShortCut;
end;

procedure MenuItem_SetShortCut(AObj: TMenuItem; AValue: TShortCut); extdecl;
begin
  AObj.ShortCut := AValue;
end;

function MenuItem_GetVisible(AObj: TMenuItem): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure MenuItem_SetVisible(AObj: TMenuItem; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure MenuItem_SetOnClick(AObj: TMenuItem; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    EventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := @EventClass.OnClick;
  EventClass.Add(AObj, geClick, AEventId);
end;

function MenuItem_GetComponentCount(AObj: TMenuItem): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function MenuItem_GetComponentIndex(AObj: TMenuItem): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure MenuItem_SetComponentIndex(AObj: TMenuItem; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function MenuItem_GetOwner(AObj: TMenuItem): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function MenuItem_GetName(AObj: TMenuItem): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure MenuItem_SetName(AObj: TMenuItem; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function MenuItem_GetTag(AObj: TMenuItem): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure MenuItem_SetTag(AObj: TMenuItem; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function MenuItem_GetItems(AObj: TMenuItem; Index: Integer): TMenuItem; extdecl;
begin
  Result :=  AObj.Items[Index];
end;

function MenuItem_GetComponents(AObj: TMenuItem; AIndex: Integer): TComponent; extdecl;
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

