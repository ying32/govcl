
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function List_Create: TList; extdecl;
begin
  Result :=  TList.Create;
end;

procedure List_Free(AObj: TList); extdecl;
begin
  AObj.Free;
end;

function List_Add(AObj: TList; Item: Pointer): Integer; extdecl;
begin
  Result :=  AObj.Add(Item);
end;

procedure List_Clear(AObj: TList); extdecl;
begin
  AObj.Clear;
end;

procedure List_Delete(AObj: TList; Index: Integer); extdecl;
begin
  AObj.Delete(Index);
end;

function List_Expand(AObj: TList): TList; extdecl;
begin
  Result :=  AObj.Expand;
end;

function List_IndexOf(AObj: TList; Item: Pointer): Integer; extdecl;
begin
  Result :=  AObj.IndexOf(Item);
end;

procedure List_Insert(AObj: TList; Index: Integer; Item: Pointer); extdecl;
begin
  AObj.Insert(Index, Item);
end;

procedure List_Move(AObj: TList; CurIndex: Integer; NewIndex: Integer); extdecl;
begin
  AObj.Move(CurIndex, NewIndex);
end;

function List_ClassName(AObj: TList): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function List_Equals(AObj: TList; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function List_GetHashCode(AObj: TList): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function List_ToString(AObj: TList): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function List_GetCount(AObj: TList): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

procedure List_SetCount(AObj: TList; AValue: Integer); extdecl;
begin
  AObj.Count := AValue;
end;

function List_GetList(AObj: TList): PPointerList; extdecl;
begin
  Result :=  AObj.List;
end;

function List_GetItems(AObj: TList; Index: Integer): Pointer; extdecl;
begin
  Result :=  AObj.Items[Index];
end;

procedure List_SetItems(AObj: TList; Index: Integer; AValue: Pointer); extdecl;
begin
  AObj.Items[Index] := AValue;
end;


exports
  List_Create,
  List_Free,
  List_Add,
  List_Clear,
  List_Delete,
  List_Expand,
  List_IndexOf,
  List_Insert,
  List_Move,
  List_ClassName,
  List_Equals,
  List_GetHashCode,
  List_ToString,
  List_GetCount,
  List_SetCount,
  List_GetList,
  List_GetItems,
  List_SetItems;

