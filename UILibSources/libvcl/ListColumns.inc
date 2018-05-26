
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ListColumns_Create(AOwner: TCustomListView): TListColumns; stdcall;
begin
  Result :=  TListColumns.Create(AOwner);
end;

procedure ListColumns_Free(AObj: TListColumns); stdcall;
begin
  AObj.Free;
end;

function ListColumns_Add(AObj: TListColumns): TListColumn; stdcall;
begin
  Result :=  AObj.Add;
end;

function ListColumns_Owner(AObj: TListColumns): TCustomListView; stdcall;
begin
  Result :=  AObj.Owner;
end;

procedure ListColumns_Assign(AObj: TListColumns; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ListColumns_BeginUpdate(AObj: TListColumns); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure ListColumns_Clear(AObj: TListColumns); stdcall;
begin
  AObj.Clear;
end;

procedure ListColumns_Delete(AObj: TListColumns; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure ListColumns_EndUpdate(AObj: TListColumns); stdcall;
begin
  AObj.EndUpdate;
end;

function ListColumns_GetNamePath(AObj: TListColumns): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ListColumns_Insert(AObj: TListColumns; Index: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.Insert(Index);
end;

function ListColumns_ClassName(AObj: TListColumns): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ListColumns_Equals(AObj: TListColumns; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListColumns_GetHashCode(AObj: TListColumns): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ListColumns_ToString(AObj: TListColumns): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListColumns_GetCapacity(AObj: TListColumns): Integer; stdcall;
begin
  Result :=  AObj.Capacity;
end;

procedure ListColumns_SetCapacity(AObj: TListColumns; AValue: Integer); stdcall;
begin
  AObj.Capacity := AValue;
end;

function ListColumns_GetCount(AObj: TListColumns): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function ListColumns_GetItems(AObj: TListColumns; Index: Integer): TListColumn; stdcall;
begin
  Result :=  AObj.Items[Index];
end;

procedure ListColumns_SetItems(AObj: TListColumns; Index: Integer; AValue: TListColumn); stdcall;
begin
  AObj.Items[Index] := AValue;
end;


exports
  ListColumns_Create,
  ListColumns_Free,
  ListColumns_Add,
  ListColumns_Owner,
  ListColumns_Assign,
  ListColumns_BeginUpdate,
  ListColumns_Clear,
  ListColumns_Delete,
  ListColumns_EndUpdate,
  ListColumns_GetNamePath,
  ListColumns_Insert,
  ListColumns_ClassName,
  ListColumns_Equals,
  ListColumns_GetHashCode,
  ListColumns_ToString,
  ListColumns_GetCapacity,
  ListColumns_SetCapacity,
  ListColumns_GetCount,
  ListColumns_GetItems,
  ListColumns_SetItems;

