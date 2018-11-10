
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function JumpList_Create(AOwner: TComponent): TJumpList; stdcall;
begin
  Result :=  TJumpList.Create(AOwner);
end;

procedure JumpList_Free(AObj: TJumpList); stdcall;
begin
  AObj.Free;
end;

procedure JumpList_AddToRecent(AObj: TJumpList; Path: PChar); stdcall;
begin
  AObj.AddToRecent(Path);
end;

function JumpList_RemoveFromRecent(AObj: TJumpList; Path: PChar; AppModelID: PChar): LongBool; stdcall;
begin
  Result :=  AObj.RemoveFromRecent(Path, AppModelID);
end;

function JumpList_RemoveAllFromRecent(AObj: TJumpList; AppModelID: PChar): LongBool; stdcall;
begin
  Result :=  AObj.RemoveAllFromRecent(AppModelID);
end;

function JumpList_AddCategory(AObj: TJumpList; CategoryName: PChar): Integer; stdcall;
begin
  Result :=  AObj.AddCategory(CategoryName);
end;

function JumpList_AddTask(AObj: TJumpList; FriendlyName: PChar; Path: PChar; Arguments: PChar; Icon: PChar): TJumpListItem; stdcall;
begin
  Result :=  AObj.AddTask(FriendlyName, Path, Arguments, Icon);
end;

function JumpList_AddTaskSeparator(AObj: TJumpList): TJumpListItem; stdcall;
begin
  Result :=  AObj.AddTaskSeparator;
end;

function JumpList_AddItemToCategory(AObj: TJumpList; CategoryIndex: Integer; FriendlyName: PChar; Path: PChar; Arguments: PChar; Icon: PChar): TJumpListItem; stdcall;
begin
  Result :=  AObj.AddItemToCategory(CategoryIndex, FriendlyName, Path, Arguments, Icon);
end;

function JumpList_UpdateList(AObj: TJumpList): LongBool; stdcall;
begin
  Result :=  AObj.UpdateList;
end;

function JumpList_DeleteList(AObj: TJumpList): LongBool; stdcall;
begin
  Result :=  AObj.DeleteList;
end;

function JumpList_FindComponent(AObj: TJumpList; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function JumpList_GetNamePath(AObj: TJumpList): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function JumpList_HasParent(AObj: TJumpList): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure JumpList_Assign(AObj: TJumpList; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure JumpList_DisposeOf(AObj: TJumpList); stdcall;
begin
  AObj.DisposeOf;
end;

function JumpList_ClassType(AObj: TJumpList): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function JumpList_ClassName(AObj: TJumpList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function JumpList_InstanceSize(AObj: TJumpList): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function JumpList_InheritsFrom(AObj: TJumpList; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function JumpList_Equals(AObj: TJumpList; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function JumpList_GetHashCode(AObj: TJumpList): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function JumpList_ToString(AObj: TJumpList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function JumpList_GetAutoRefresh(AObj: TJumpList): LongBool; stdcall;
begin
  Result :=  AObj.AutoRefresh;
end;

procedure JumpList_SetAutoRefresh(AObj: TJumpList; AValue: LongBool); stdcall;
begin
  AObj.AutoRefresh := AValue;
end;

function JumpList_GetEnabled(AObj: TJumpList): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure JumpList_SetEnabled(AObj: TJumpList; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function JumpList_GetApplicationID(AObj: TJumpList): PChar; stdcall;
begin
  Result :=  PChar(AObj.ApplicationID);
end;

procedure JumpList_SetApplicationID(AObj: TJumpList; AValue: PChar); stdcall;
begin
  AObj.ApplicationID := AValue;
end;

function JumpList_GetCustomCategories(AObj: TJumpList): TJumpCategories; stdcall;
begin
  Result :=  AObj.CustomCategories;
end;

procedure JumpList_SetCustomCategories(AObj: TJumpList; AValue: TJumpCategories); stdcall;
begin
  AObj.CustomCategories := AValue;
end;

function JumpList_GetShowRecent(AObj: TJumpList): LongBool; stdcall;
begin
  Result :=  AObj.ShowRecent;
end;

procedure JumpList_SetShowRecent(AObj: TJumpList; AValue: LongBool); stdcall;
begin
  AObj.ShowRecent := AValue;
end;

function JumpList_GetShowFrequent(AObj: TJumpList): LongBool; stdcall;
begin
  Result :=  AObj.ShowFrequent;
end;

procedure JumpList_SetShowFrequent(AObj: TJumpList; AValue: LongBool); stdcall;
begin
  AObj.ShowFrequent := AValue;
end;

function JumpList_GetTaskList(AObj: TJumpList): TJumpListCollection; stdcall;
begin
  Result :=  AObj.TaskList;
end;

procedure JumpList_SetTaskList(AObj: TJumpList; AValue: TJumpListCollection); stdcall;
begin
  AObj.TaskList := AValue;
end;

procedure JumpList_SetOnItemDeleted(AObj: TJumpList; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnItemDeleted := nil;
    TEventClass.Remove(AObj, @TEventClass.JumpListOnItemDeleted);
    Exit;
  end;
  AObj.OnItemDeleted := TEventClass.JumpListOnItemDeleted;
  TEventClass.Add(AObj, @TEventClass.JumpListOnItemDeleted, AEventId);
end;

procedure JumpList_SetOnListUpdateError(AObj: TJumpList; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnListUpdateError := nil;
    TEventClass.Remove(AObj, @TEventClass.JumpListOnListUpdateError);
    Exit;
  end;
  AObj.OnListUpdateError := TEventClass.JumpListOnListUpdateError;
  TEventClass.Add(AObj, @TEventClass.JumpListOnListUpdateError, AEventId);
end;

procedure JumpList_SetOnItemsLoaded(AObj: TJumpList; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnItemsLoaded := nil;
    TEventClass.Remove(AObj, @TEventClass.JumpListOnItemsLoaded);
    Exit;
  end;
  AObj.OnItemsLoaded := TEventClass.JumpListOnItemsLoaded;
  TEventClass.Add(AObj, @TEventClass.JumpListOnItemsLoaded, AEventId);
end;

function JumpList_GetComponentCount(AObj: TJumpList): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function JumpList_GetComponentIndex(AObj: TJumpList): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure JumpList_SetComponentIndex(AObj: TJumpList; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function JumpList_GetOwner(AObj: TJumpList): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function JumpList_GetName(AObj: TJumpList): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure JumpList_SetName(AObj: TJumpList; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function JumpList_GetTag(AObj: TJumpList): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure JumpList_SetTag(AObj: TJumpList; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function JumpList_GetComponents(AObj: TJumpList; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function JumpList_StaticClassType: TClass; stdcall;
begin
  Result := TJumpList;
end;

exports
  JumpList_Create,
  JumpList_Free,
  JumpList_AddToRecent,
  JumpList_RemoveFromRecent,
  JumpList_RemoveAllFromRecent,
  JumpList_AddCategory,
  JumpList_AddTask,
  JumpList_AddTaskSeparator,
  JumpList_AddItemToCategory,
  JumpList_UpdateList,
  JumpList_DeleteList,
  JumpList_FindComponent,
  JumpList_GetNamePath,
  JumpList_HasParent,
  JumpList_Assign,
  JumpList_DisposeOf,
  JumpList_ClassType,
  JumpList_ClassName,
  JumpList_InstanceSize,
  JumpList_InheritsFrom,
  JumpList_Equals,
  JumpList_GetHashCode,
  JumpList_ToString,
  JumpList_GetAutoRefresh,
  JumpList_SetAutoRefresh,
  JumpList_GetEnabled,
  JumpList_SetEnabled,
  JumpList_GetApplicationID,
  JumpList_SetApplicationID,
  JumpList_GetCustomCategories,
  JumpList_SetCustomCategories,
  JumpList_GetShowRecent,
  JumpList_SetShowRecent,
  JumpList_GetShowFrequent,
  JumpList_SetShowFrequent,
  JumpList_GetTaskList,
  JumpList_SetTaskList,
  JumpList_SetOnItemDeleted,
  JumpList_SetOnListUpdateError,
  JumpList_SetOnItemsLoaded,
  JumpList_GetComponentCount,
  JumpList_GetComponentIndex,
  JumpList_SetComponentIndex,
  JumpList_GetOwner,
  JumpList_GetName,
  JumpList_SetName,
  JumpList_GetTag,
  JumpList_SetTag,
  JumpList_GetComponents,
  JumpList_StaticClassType;

