
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ListGroups_Create(AOwner: TCustomListView): TListGroups; stdcall;
begin
  Result :=  TListGroups.Create(AOwner);
end;

procedure ListGroups_Free(AObj: TListGroups); stdcall;
begin
  AObj.Free;
end;

function ListGroups_Add(AObj: TListGroups): TListGroup; stdcall;
begin
  Result :=  AObj.Add;
end;

function ListGroups_Owner(AObj: TListGroups): TCustomListView; stdcall;
begin
  Result :=  AObj.Owner;
end;

procedure ListGroups_Assign(AObj: TListGroups; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ListGroups_BeginUpdate(AObj: TListGroups); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure ListGroups_Clear(AObj: TListGroups); stdcall;
begin
  AObj.Clear;
end;

procedure ListGroups_Delete(AObj: TListGroups; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure ListGroups_EndUpdate(AObj: TListGroups); stdcall;
begin
  AObj.EndUpdate;
end;

function ListGroups_GetNamePath(AObj: TListGroups): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ListGroups_Insert(AObj: TListGroups; Index: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.Insert(Index);
end;

function ListGroups_ClassName(AObj: TListGroups): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ListGroups_Equals(AObj: TListGroups; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListGroups_GetHashCode(AObj: TListGroups): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ListGroups_ToString(AObj: TListGroups): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListGroups_GetCapacity(AObj: TListGroups): Integer; stdcall;
begin
  Result :=  AObj.Capacity;
end;

procedure ListGroups_SetCapacity(AObj: TListGroups; AValue: Integer); stdcall;
begin
  AObj.Capacity := AValue;
end;

function ListGroups_GetCount(AObj: TListGroups): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function ListGroups_GetItems(AObj: TListGroups; Index: Integer): TListGroup; stdcall;
begin
  Result :=  AObj.Items[Index];
end;

procedure ListGroups_SetItems(AObj: TListGroups; Index: Integer; AValue: TListGroup); stdcall;
begin
  AObj.Items[Index] := AValue;
end;


exports
  ListGroups_Create,
  ListGroups_Free,
  ListGroups_Add,
  ListGroups_Owner,
  ListGroups_Assign,
  ListGroups_BeginUpdate,
  ListGroups_Clear,
  ListGroups_Delete,
  ListGroups_EndUpdate,
  ListGroups_GetNamePath,
  ListGroups_Insert,
  ListGroups_ClassName,
  ListGroups_Equals,
  ListGroups_GetHashCode,
  ListGroups_ToString,
  ListGroups_GetCapacity,
  ListGroups_SetCapacity,
  ListGroups_GetCount,
  ListGroups_GetItems,
  ListGroups_SetItems;

