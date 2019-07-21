function ComboExItems_Add(AObj: TComboExItems): TComboExItem; stdcall;
begin
  Result :=  AObj.Add;
end;

function ComboExItems_AddItem(AObj: TComboExItems; Caption: PChar; ImageIndex: Integer; SelectedImageIndex: Integer; OverlayImageIndex: Integer; Indent: Integer; Data: Pointer): TComboExItem; stdcall;
begin
  Result :=  AObj.AddItem(Caption, ImageIndex, SelectedImageIndex, OverlayImageIndex, Indent, Data);
end;

function ComboExItems_Insert(AObj: TComboExItems; Index: Integer): TComboExItem; stdcall;
begin
  Result :=  AObj.Insert(Index);
end;

function ComboExItems_Owner(AObj: TComboExItems): TPersistent; stdcall;
begin
  Result :=  AObj.Owner;
end;

procedure ComboExItems_Assign(AObj: TComboExItems; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ComboExItems_BeginUpdate(AObj: TComboExItems); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure ComboExItems_Clear(AObj: TComboExItems); stdcall;
begin
  AObj.Clear;
end;

procedure ComboExItems_ClearAndResetID(AObj: TComboExItems); stdcall;
begin
  AObj.ClearAndResetID;
end;

procedure ComboExItems_Delete(AObj: TComboExItems; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure ComboExItems_EndUpdate(AObj: TComboExItems); stdcall;
begin
  AObj.EndUpdate;
end;

function ComboExItems_FindItemID(AObj: TComboExItems; ID: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.FindItemID(ID);
end;

function ComboExItems_GetNamePath(AObj: TComboExItems): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ComboExItems_DisposeOf(AObj: TComboExItems); stdcall;
begin
  AObj.DisposeOf;
end;

function ComboExItems_ClassType(AObj: TComboExItems): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function ComboExItems_ClassName(AObj: TComboExItems): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ComboExItems_InstanceSize(AObj: TComboExItems): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function ComboExItems_InheritsFrom(AObj: TComboExItems; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function ComboExItems_Equals(AObj: TComboExItems; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ComboExItems_GetHashCode(AObj: TComboExItems): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ComboExItems_ToString(AObj: TComboExItems): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ComboExItems_GetCapacity(AObj: TComboExItems): Integer; stdcall;
begin
  Result :=  AObj.Capacity;
end;

procedure ComboExItems_SetCapacity(AObj: TComboExItems; AValue: Integer); stdcall;
begin
  AObj.Capacity := AValue;
end;

function ComboExItems_GetCount(AObj: TComboExItems): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function ComboExItems_GetComboItems(AObj: TComboExItems; Index: Integer): TComboExItem; stdcall;
begin
  Result :=  AObj.ComboItems[Index];
end;

function ComboExItems_StaticClassType: TClass; stdcall;
begin
  Result := TComboExItems;
end;

exports
  ComboExItems_Add,
  ComboExItems_AddItem,
  ComboExItems_Insert,
  ComboExItems_Owner,
  ComboExItems_Assign,
  ComboExItems_BeginUpdate,
  ComboExItems_Clear,
  ComboExItems_ClearAndResetID,
  ComboExItems_Delete,
  ComboExItems_EndUpdate,
  ComboExItems_FindItemID,
  ComboExItems_GetNamePath,
  ComboExItems_DisposeOf,
  ComboExItems_ClassType,
  ComboExItems_ClassName,
  ComboExItems_InstanceSize,
  ComboExItems_InheritsFrom,
  ComboExItems_Equals,
  ComboExItems_GetHashCode,
  ComboExItems_ToString,
  ComboExItems_GetCapacity,
  ComboExItems_SetCapacity,
  ComboExItems_GetCount,
  ComboExItems_GetComboItems,
  ComboExItems_StaticClassType;

