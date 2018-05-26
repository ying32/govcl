
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ListColumns_Create(AOwner: TCustomListView): TListColumns; extdecl;
begin
  Result :=  TListColumns.Create(AOwner);
end;

procedure ListColumns_Free(AObj: TListColumns); extdecl;
begin
  AObj.Free;
end;

function ListColumns_Add(AObj: TListColumns): TListColumn; extdecl;
begin
  Result :=  AObj.Add;
end;

function ListColumns_Owner(AObj: TListColumns): TCustomListView; extdecl;
begin
  Result :=  AObj.Owner;
end;

procedure ListColumns_Assign(AObj: TListColumns; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

procedure ListColumns_BeginUpdate(AObj: TListColumns); extdecl;
begin
  AObj.BeginUpdate;
end;

procedure ListColumns_Clear(AObj: TListColumns); extdecl;
begin
  AObj.Clear;
end;

procedure ListColumns_Delete(AObj: TListColumns; Index: Integer); extdecl;
begin
  AObj.Delete(Index);
end;

procedure ListColumns_EndUpdate(AObj: TListColumns); extdecl;
begin
  AObj.EndUpdate;
end;

function ListColumns_GetNamePath(AObj: TListColumns): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ListColumns_Insert(AObj: TListColumns; Index: Integer): TCollectionItem; extdecl;
begin
  Result :=  AObj.Insert(Index);
end;

function ListColumns_ClassName(AObj: TListColumns): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ListColumns_Equals(AObj: TListColumns; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListColumns_GetHashCode(AObj: TListColumns): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ListColumns_ToString(AObj: TListColumns): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListColumns_GetCount(AObj: TListColumns): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

function ListColumns_GetItems(AObj: TListColumns; Index: Integer): TListColumn; extdecl;
begin
  Result :=  AObj.Items[Index];
end;

procedure ListColumns_SetItems(AObj: TListColumns; Index: Integer; AValue: TListColumn); extdecl;
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
  ListColumns_GetCount,
  ListColumns_GetItems,
  ListColumns_SetItems;

