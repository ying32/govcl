
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ThumbBarButtonList_Create(AOwner: TComponent): TThumbBarButtonList; stdcall;
begin
  Result :=  TThumbBarButtonList.Create(AOwner);
end;

procedure ThumbBarButtonList_Free(AObj: TThumbBarButtonList); stdcall;
begin
  AObj.Free;
end;

function ThumbBarButtonList_Add(AObj: TThumbBarButtonList): TThumbBarButton; stdcall;
begin
  Result :=  AObj.Add;
end;

function ThumbBarButtonList_Owner(AObj: TThumbBarButtonList): TPersistent; stdcall;
begin
  Result :=  AObj.Owner;
end;

procedure ThumbBarButtonList_Assign(AObj: TThumbBarButtonList; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ThumbBarButtonList_BeginUpdate(AObj: TThumbBarButtonList); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure ThumbBarButtonList_Clear(AObj: TThumbBarButtonList); stdcall;
begin
  AObj.Clear;
end;

procedure ThumbBarButtonList_ClearAndResetID(AObj: TThumbBarButtonList); stdcall;
begin
  AObj.ClearAndResetID;
end;

procedure ThumbBarButtonList_Delete(AObj: TThumbBarButtonList; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure ThumbBarButtonList_EndUpdate(AObj: TThumbBarButtonList); stdcall;
begin
  AObj.EndUpdate;
end;

function ThumbBarButtonList_FindItemID(AObj: TThumbBarButtonList; ID: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.FindItemID(ID);
end;

function ThumbBarButtonList_GetNamePath(AObj: TThumbBarButtonList): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ThumbBarButtonList_Insert(AObj: TThumbBarButtonList; Index: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.Insert(Index);
end;

procedure ThumbBarButtonList_DisposeOf(AObj: TThumbBarButtonList); stdcall;
begin
  AObj.DisposeOf;
end;

function ThumbBarButtonList_ClassType(AObj: TThumbBarButtonList): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function ThumbBarButtonList_ClassName(AObj: TThumbBarButtonList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ThumbBarButtonList_InstanceSize(AObj: TThumbBarButtonList): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function ThumbBarButtonList_InheritsFrom(AObj: TThumbBarButtonList; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function ThumbBarButtonList_Equals(AObj: TThumbBarButtonList; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ThumbBarButtonList_GetHashCode(AObj: TThumbBarButtonList): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ThumbBarButtonList_ToString(AObj: TThumbBarButtonList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

procedure ThumbBarButtonList_SetOnChange(AObj: TThumbBarButtonList; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, @TEventClass.OnChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, @TEventClass.OnChange, AEventId);
end;

function ThumbBarButtonList_GetCapacity(AObj: TThumbBarButtonList): Integer; stdcall;
begin
  Result :=  AObj.Capacity;
end;

procedure ThumbBarButtonList_SetCapacity(AObj: TThumbBarButtonList; AValue: Integer); stdcall;
begin
  AObj.Capacity := AValue;
end;

function ThumbBarButtonList_GetCount(AObj: TThumbBarButtonList): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function ThumbBarButtonList_GetItems(AObj: TThumbBarButtonList; Index: Integer): TThumbBarButton; stdcall;
begin
  Result :=  AObj.Items[Index];
end;

procedure ThumbBarButtonList_SetItems(AObj: TThumbBarButtonList; Index: Integer; AValue: TThumbBarButton); stdcall;
begin
  AObj.Items[Index] := AValue;
end;

function ThumbBarButtonList_StaticClassType: TClass; stdcall;
begin
  Result := TThumbBarButtonList;
end;

exports
  ThumbBarButtonList_Create,
  ThumbBarButtonList_Free,
  ThumbBarButtonList_Add,
  ThumbBarButtonList_Owner,
  ThumbBarButtonList_Assign,
  ThumbBarButtonList_BeginUpdate,
  ThumbBarButtonList_Clear,
  ThumbBarButtonList_ClearAndResetID,
  ThumbBarButtonList_Delete,
  ThumbBarButtonList_EndUpdate,
  ThumbBarButtonList_FindItemID,
  ThumbBarButtonList_GetNamePath,
  ThumbBarButtonList_Insert,
  ThumbBarButtonList_DisposeOf,
  ThumbBarButtonList_ClassType,
  ThumbBarButtonList_ClassName,
  ThumbBarButtonList_InstanceSize,
  ThumbBarButtonList_InheritsFrom,
  ThumbBarButtonList_Equals,
  ThumbBarButtonList_GetHashCode,
  ThumbBarButtonList_ToString,
  ThumbBarButtonList_SetOnChange,
  ThumbBarButtonList_GetCapacity,
  ThumbBarButtonList_SetCapacity,
  ThumbBarButtonList_GetCount,
  ThumbBarButtonList_GetItems,
  ThumbBarButtonList_SetItems,
  ThumbBarButtonList_StaticClassType;

