
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function TaskDialogButtons_Create(AOwner: TPersistent; ItemClass: TCollectionItemClass): TTaskDialogButtons; stdcall;
begin
  Result :=  TTaskDialogButtons.Create(AOwner, ItemClass);
end;

procedure TaskDialogButtons_Free(AObj: TTaskDialogButtons); stdcall;
begin
  AObj.Free;
end;

function TaskDialogButtons_Add(AObj: TTaskDialogButtons): TTaskDialogBaseButtonItem; stdcall;
begin
  Result :=  AObj.Add;
end;

function TaskDialogButtons_Buttons(AObj: TTaskDialogButtons): PTaskDialogButton; stdcall;
begin
  Result :=  AObj.Buttons;
end;

function TaskDialogButtons_Owner(AObj: TTaskDialogButtons): TPersistent; stdcall;
begin
  Result :=  AObj.Owner;
end;

procedure TaskDialogButtons_Assign(AObj: TTaskDialogButtons; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure TaskDialogButtons_BeginUpdate(AObj: TTaskDialogButtons); stdcall;
begin
  AObj.BeginUpdate;
end;

procedure TaskDialogButtons_Clear(AObj: TTaskDialogButtons); stdcall;
begin
  AObj.Clear;
end;

procedure TaskDialogButtons_ClearAndResetID(AObj: TTaskDialogButtons); stdcall;
begin
  AObj.ClearAndResetID;
end;

procedure TaskDialogButtons_Delete(AObj: TTaskDialogButtons; Index: Integer); stdcall;
begin
  AObj.Delete(Index);
end;

procedure TaskDialogButtons_EndUpdate(AObj: TTaskDialogButtons); stdcall;
begin
  AObj.EndUpdate;
end;

function TaskDialogButtons_FindItemID(AObj: TTaskDialogButtons; ID: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.FindItemID(ID);
end;

function TaskDialogButtons_GetNamePath(AObj: TTaskDialogButtons): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function TaskDialogButtons_Insert(AObj: TTaskDialogButtons; Index: Integer): TCollectionItem; stdcall;
begin
  Result :=  AObj.Insert(Index);
end;

procedure TaskDialogButtons_DisposeOf(AObj: TTaskDialogButtons); stdcall;
begin
  AObj.DisposeOf;
end;

function TaskDialogButtons_ClassType(AObj: TTaskDialogButtons): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function TaskDialogButtons_ClassName(AObj: TTaskDialogButtons): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TaskDialogButtons_InstanceSize(AObj: TTaskDialogButtons): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function TaskDialogButtons_InheritsFrom(AObj: TTaskDialogButtons; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function TaskDialogButtons_Equals(AObj: TTaskDialogButtons; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function TaskDialogButtons_GetHashCode(AObj: TTaskDialogButtons): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function TaskDialogButtons_ToString(AObj: TTaskDialogButtons): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function TaskDialogButtons_GetDefaultButton(AObj: TTaskDialogButtons): TTaskDialogBaseButtonItem; stdcall;
begin
  Result :=  AObj.DefaultButton;
end;

procedure TaskDialogButtons_SetDefaultButton(AObj: TTaskDialogButtons; AValue: TTaskDialogBaseButtonItem); stdcall;
begin
  AObj.DefaultButton := AValue;
end;

function TaskDialogButtons_GetCapacity(AObj: TTaskDialogButtons): Integer; stdcall;
begin
  Result :=  AObj.Capacity;
end;

procedure TaskDialogButtons_SetCapacity(AObj: TTaskDialogButtons; AValue: Integer); stdcall;
begin
  AObj.Capacity := AValue;
end;

function TaskDialogButtons_GetCount(AObj: TTaskDialogButtons): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

function TaskDialogButtons_GetItems(AObj: TTaskDialogButtons; Index: Integer): TTaskDialogBaseButtonItem; stdcall;
begin
  Result :=  AObj.Items[Index];
end;

procedure TaskDialogButtons_SetItems(AObj: TTaskDialogButtons; Index: Integer; AValue: TTaskDialogBaseButtonItem); stdcall;
begin
  AObj.Items[Index] := AValue;
end;

function TaskDialogButtons_StaticClassType: TClass; stdcall;
begin
  Result := TTaskDialogButtons;
end;

exports
  TaskDialogButtons_Create,
  TaskDialogButtons_Free,
  TaskDialogButtons_Add,
  TaskDialogButtons_Buttons,
  TaskDialogButtons_Owner,
  TaskDialogButtons_Assign,
  TaskDialogButtons_BeginUpdate,
  TaskDialogButtons_Clear,
  TaskDialogButtons_ClearAndResetID,
  TaskDialogButtons_Delete,
  TaskDialogButtons_EndUpdate,
  TaskDialogButtons_FindItemID,
  TaskDialogButtons_GetNamePath,
  TaskDialogButtons_Insert,
  TaskDialogButtons_DisposeOf,
  TaskDialogButtons_ClassType,
  TaskDialogButtons_ClassName,
  TaskDialogButtons_InstanceSize,
  TaskDialogButtons_InheritsFrom,
  TaskDialogButtons_Equals,
  TaskDialogButtons_GetHashCode,
  TaskDialogButtons_ToString,
  TaskDialogButtons_GetDefaultButton,
  TaskDialogButtons_SetDefaultButton,
  TaskDialogButtons_GetCapacity,
  TaskDialogButtons_SetCapacity,
  TaskDialogButtons_GetCount,
  TaskDialogButtons_GetItems,
  TaskDialogButtons_SetItems,
  TaskDialogButtons_StaticClassType;

