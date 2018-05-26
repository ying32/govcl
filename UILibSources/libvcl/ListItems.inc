
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ListItems_Create(AOwner: TCustomListView): TListItems; stdcall;
begin
  Result :=  TListItems.Create(AOwner);
end;

procedure ListItems_Free(AObj: TListItems); stdcall;
begin
  AObj.Free;
end;

function ListItems_Add(AObj: TListItems): TListItem; stdcall;
begin
  Result :=  AObj.Add;
end;

function ListItems_AddItem(AObj: TListItems; Item: TListItem; Index: Integer): TListItem; stdcall;
begin
  Result :=  AObj.AddItem(Item, Index);
end;

procedure ListItems_Assign(AObj: TListItems; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ListItems_BeginUpdate(AObj: TListItems); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure ListItems_Clear(AObj: TListItems); stdcall;
begin
  AObj.Clear;
end;

procedure ListItems_Delete(AObj: TListItems; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure ListItems_EndUpdate(AObj: TListItems); stdcall;
begin
  AObj.EndUpdate;
end;

function ListItems_IndexOf(AObj: TListItems; Value: TListItem): Integer; stdcall;
begin
  Result :=  AObj.IndexOf(Value);
end;

function ListItems_Insert(AObj: TListItems; Index: Integer): TListItem; stdcall;
begin
  Result :=  AObj.Insert(Index);
end;

function ListItems_GetNamePath(AObj: TListItems): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ListItems_ClassName(AObj: TListItems): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ListItems_Equals(AObj: TListItems; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListItems_GetHashCode(AObj: TListItems): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ListItems_ToString(AObj: TListItems): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListItems_GetCount(AObj: TListItems): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

procedure ListItems_SetCount(AObj: TListItems; AValue: Integer); stdcall;
begin
  AObj.Count := AValue;
end;

function ListItems_GetHandle(AObj: TListItems): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ListItems_GetOwner(AObj: TListItems): TCustomListView; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ListItems_GetItem(AObj: TListItems; Index: Integer): TListItem; stdcall;
begin
  Result :=  AObj.Item[Index];
end;

procedure ListItems_SetItem(AObj: TListItems; Index: Integer; AValue: TListItem); stdcall;
begin
  AObj.Item[Index] := AValue;
end;


exports
  ListItems_Create,
  ListItems_Free,
  ListItems_Add,
  ListItems_AddItem,
  ListItems_Assign,
  ListItems_BeginUpdate,
  ListItems_Clear,
  ListItems_Delete,
  ListItems_EndUpdate,
  ListItems_IndexOf,
  ListItems_Insert,
  ListItems_GetNamePath,
  ListItems_ClassName,
  ListItems_Equals,
  ListItems_GetHashCode,
  ListItems_ToString,
  ListItems_GetCount,
  ListItems_SetCount,
  ListItems_GetHandle,
  ListItems_GetOwner,
  ListItems_GetItem,
  ListItems_SetItem;

