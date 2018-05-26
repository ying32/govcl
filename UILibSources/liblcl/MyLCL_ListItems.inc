
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ListItems_Create(AOwner: TCustomListView): TListItems; extdecl;
begin
  Result :=  TListItems.Create(AOwner);
end;

procedure ListItems_Free(AObj: TListItems); extdecl;
begin
  AObj.Free;
end;

function ListItems_Add(AObj: TListItems): TListItem; extdecl;
begin
  Result :=  AObj.Add;
end;

procedure ListItems_Assign(AObj: TListItems; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure ListItems_BeginUpdate(AObj: TListItems); extdecl;
begin
  AObj.BeginUpdate;
end;

procedure ListItems_Clear(AObj: TListItems); extdecl;
begin
  AObj.Clear;
end;

procedure ListItems_Delete(AObj: TListItems; Index: Integer); extdecl;
begin
  AObj.Delete(Index);
end;

procedure ListItems_EndUpdate(AObj: TListItems); extdecl;
begin
  AObj.EndUpdate;
end;

function ListItems_IndexOf(AObj: TListItems; Value: TListItem): Integer; extdecl;
begin
  Result :=  AObj.IndexOf(Value);
end;

function ListItems_Insert(AObj: TListItems; Index: Integer): TListItem; extdecl;
begin
  Result :=  AObj.Insert(Index);
end;

function ListItems_GetNamePath(AObj: TListItems): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ListItems_ClassName(AObj: TListItems): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ListItems_Equals(AObj: TListItems; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListItems_GetHashCode(AObj: TListItems): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ListItems_ToString(AObj: TListItems): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListItems_GetCount(AObj: TListItems): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

procedure ListItems_SetCount(AObj: TListItems; AValue: Integer); extdecl;
begin
  AObj.Count := AValue;
end;

function ListItems_GetOwner(AObj: TListItems): TCustomListView; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ListItems_GetItem(AObj: TListItems; Index: Integer): TListItem; extdecl;
begin
  Result :=  AObj.Item[Index];
end;

procedure ListItems_SetItem(AObj: TListItems; Index: Integer; AValue: TListItem); extdecl;
begin
  AObj.Item[Index] := AValue;
end;


exports
  ListItems_Create,
  ListItems_Free,
  ListItems_Add,
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
  ListItems_GetOwner,
  ListItems_GetItem,
  ListItems_SetItem;

