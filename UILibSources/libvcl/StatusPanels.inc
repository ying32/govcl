
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function StatusPanels_Create(AOwner: TCustomStatusBar): TStatusPanels; stdcall;
begin
  Result :=  TStatusPanels.Create(AOwner);
end;

procedure StatusPanels_Free(AObj: TStatusPanels); stdcall;
begin
  AObj.Free;
end;

function StatusPanels_Add(AObj: TStatusPanels): TStatusPanel; stdcall;
begin
  Result :=  AObj.Add;
end;

function StatusPanels_AddItem(AObj: TStatusPanels; Item: TStatusPanel; Index: Integer): TStatusPanel; stdcall;
begin
  Result :=  AObj.AddItem(Item, Index);
end;

function StatusPanels_Insert(AObj: TStatusPanels; Index: Integer): TStatusPanel; stdcall;
begin
  Result :=  AObj.Insert(Index);
end;

function StatusPanels_Owner(AObj: TStatusPanels): TPersistent; stdcall;
begin
  Result :=  AObj.Owner;
end;

procedure StatusPanels_Assign(AObj: TStatusPanels; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure StatusPanels_BeginUpdate(AObj: TStatusPanels); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure StatusPanels_Clear(AObj: TStatusPanels); stdcall;
begin
  AObj.Clear;
end;

procedure StatusPanels_Delete(AObj: TStatusPanels; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure StatusPanels_EndUpdate(AObj: TStatusPanels); stdcall;
begin
  AObj.EndUpdate;
end;

function StatusPanels_GetNamePath(AObj: TStatusPanels): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function StatusPanels_ClassName(AObj: TStatusPanels): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function StatusPanels_Equals(AObj: TStatusPanels; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function StatusPanels_GetHashCode(AObj: TStatusPanels): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function StatusPanels_ToString(AObj: TStatusPanels): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function StatusPanels_GetCapacity(AObj: TStatusPanels): Integer; stdcall;
begin
  Result :=  AObj.Capacity;
end;

procedure StatusPanels_SetCapacity(AObj: TStatusPanels; AValue: Integer); stdcall;
begin
  AObj.Capacity := AValue;
end;

function StatusPanels_GetCount(AObj: TStatusPanels): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function StatusPanels_GetItems(AObj: TStatusPanels; Index: Integer): TStatusPanel; stdcall;
begin
  Result :=  AObj.Items[Index];
end;

procedure StatusPanels_SetItems(AObj: TStatusPanels; Index: Integer; AValue: TStatusPanel); stdcall;
begin
  AObj.Items[Index] := AValue;
end;


exports
  StatusPanels_Create,
  StatusPanels_Free,
  StatusPanels_Add,
  StatusPanels_AddItem,
  StatusPanels_Insert,
  StatusPanels_Owner,
  StatusPanels_Assign,
  StatusPanels_BeginUpdate,
  StatusPanels_Clear,
  StatusPanels_Delete,
  StatusPanels_EndUpdate,
  StatusPanels_GetNamePath,
  StatusPanels_ClassName,
  StatusPanels_Equals,
  StatusPanels_GetHashCode,
  StatusPanels_ToString,
  StatusPanels_GetCapacity,
  StatusPanels_SetCapacity,
  StatusPanels_GetCount,
  StatusPanels_GetItems,
  StatusPanels_SetItems;

