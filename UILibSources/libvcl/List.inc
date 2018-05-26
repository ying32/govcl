
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function List_Create: TList; stdcall;
begin
  Result :=  TList.Create;
end;

procedure List_Free(AObj: TList); stdcall;
begin
  AObj.Free;
end;

function List_Add(AObj: TList; Item: Pointer): Integer; stdcall;
begin
  Result :=  AObj.Add(Item);
end;

procedure List_Clear(AObj: TList); stdcall;
begin
  AObj.Clear;
end;

procedure List_Delete(AObj: TList; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

function List_Expand(AObj: TList): TList; stdcall;
begin
  Result :=  AObj.Expand;
end;

function List_IndexOf(AObj: TList; Item: Pointer): Integer; stdcall;
begin
  Result :=  AObj.IndexOf(Item);
end;

procedure List_Insert(AObj: TList; Index: Integer; Item: Pointer); stdcall;
begin
  AObj.Insert(Index, Item);
end;

procedure List_Move(AObj: TList; CurIndex: Integer; NewIndex: Integer); stdcall;
begin
  AObj.Move(CurIndex, NewIndex);
end;

function List_ClassName(AObj: TList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function List_Equals(AObj: TList; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function List_GetHashCode(AObj: TList): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function List_ToString(AObj: TList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function List_GetCapacity(AObj: TList): Integer; stdcall;
begin
  Result :=  AObj.Capacity;
end;

procedure List_SetCapacity(AObj: TList; AValue: Integer); stdcall;
begin
  AObj.Capacity := AValue;
end;

function List_GetCount(AObj: TList): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

procedure List_SetCount(AObj: TList; AValue: Integer); stdcall;
begin
  AObj.Count := AValue;
end;

function List_GetList(AObj: TList): TPointerList; stdcall;
begin
  Result :=  AObj.List;
end;

function List_GetItems(AObj: TList; Index: Integer): Pointer; stdcall;
begin
  Result :=  AObj.Items[Index];
end;

procedure List_SetItems(AObj: TList; Index: Integer; AValue: Pointer); stdcall;
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
  List_GetCapacity,
  List_SetCapacity,
  List_GetCount,
  List_SetCount,
  List_GetList,
  List_GetItems,
  List_SetItems;

