
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function TaskDialogBaseButtonItem_Create(AOwner: TCollection): TTaskDialogBaseButtonItem; stdcall;
begin
  Result :=  TTaskDialogBaseButtonItem.Create(AOwner);
end;

procedure TaskDialogBaseButtonItem_Free(AObj: TTaskDialogBaseButtonItem); stdcall;
begin
  AObj.Free;
end;

procedure TaskDialogBaseButtonItem_Click(AObj: TTaskDialogBaseButtonItem); stdcall;
begin
  AObj.Click;
end;

function TaskDialogBaseButtonItem_GetNamePath(AObj: TTaskDialogBaseButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TaskDialogBaseButtonItem_Assign(AObj: TTaskDialogBaseButtonItem; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure TaskDialogBaseButtonItem_DisposeOf(AObj: TTaskDialogBaseButtonItem); stdcall;
begin
  AObj.DisposeOf;
end;

function TaskDialogBaseButtonItem_ClassType(AObj: TTaskDialogBaseButtonItem): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function TaskDialogBaseButtonItem_ClassName(AObj: TTaskDialogBaseButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TaskDialogBaseButtonItem_InstanceSize(AObj: TTaskDialogBaseButtonItem): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function TaskDialogBaseButtonItem_InheritsFrom(AObj: TTaskDialogBaseButtonItem; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function TaskDialogBaseButtonItem_Equals(AObj: TTaskDialogBaseButtonItem; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function TaskDialogBaseButtonItem_GetHashCode(AObj: TTaskDialogBaseButtonItem): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function TaskDialogBaseButtonItem_ToString(AObj: TTaskDialogBaseButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function TaskDialogBaseButtonItem_GetModalResult(AObj: TTaskDialogBaseButtonItem): TModalResult; stdcall;
begin
  Result :=  AObj.ModalResult;
end;

procedure TaskDialogBaseButtonItem_SetModalResult(AObj: TTaskDialogBaseButtonItem; AValue: TModalResult); stdcall;
begin
  AObj.ModalResult := AValue;
end;

function TaskDialogBaseButtonItem_GetCaption(AObj: TTaskDialogBaseButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure TaskDialogBaseButtonItem_SetCaption(AObj: TTaskDialogBaseButtonItem; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function TaskDialogBaseButtonItem_GetDefault(AObj: TTaskDialogBaseButtonItem): LongBool; stdcall;
begin
  Result :=  AObj.Default;
end;

procedure TaskDialogBaseButtonItem_SetDefault(AObj: TTaskDialogBaseButtonItem; AValue: LongBool); stdcall;
begin
  AObj.Default := AValue;
end;

function TaskDialogBaseButtonItem_GetEnabled(AObj: TTaskDialogBaseButtonItem): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure TaskDialogBaseButtonItem_SetEnabled(AObj: TTaskDialogBaseButtonItem; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function TaskDialogBaseButtonItem_GetCollection(AObj: TTaskDialogBaseButtonItem): TCollection; stdcall;
begin
  Result :=  AObj.Collection;
end;

procedure TaskDialogBaseButtonItem_SetCollection(AObj: TTaskDialogBaseButtonItem; AValue: TCollection); stdcall;
begin
  AObj.Collection := AValue;
end;

function TaskDialogBaseButtonItem_GetIndex(AObj: TTaskDialogBaseButtonItem): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure TaskDialogBaseButtonItem_SetIndex(AObj: TTaskDialogBaseButtonItem; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;

function TaskDialogBaseButtonItem_GetDisplayName(AObj: TTaskDialogBaseButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.DisplayName);
end;

procedure TaskDialogBaseButtonItem_SetDisplayName(AObj: TTaskDialogBaseButtonItem; AValue: PChar); stdcall;
begin
  AObj.DisplayName := AValue;
end;

function TaskDialogBaseButtonItem_StaticClassType: TClass; stdcall;
begin
  Result := TTaskDialogBaseButtonItem;
end;

exports
  TaskDialogBaseButtonItem_Create,
  TaskDialogBaseButtonItem_Free,
  TaskDialogBaseButtonItem_Click,
  TaskDialogBaseButtonItem_GetNamePath,
  TaskDialogBaseButtonItem_Assign,
  TaskDialogBaseButtonItem_DisposeOf,
  TaskDialogBaseButtonItem_ClassType,
  TaskDialogBaseButtonItem_ClassName,
  TaskDialogBaseButtonItem_InstanceSize,
  TaskDialogBaseButtonItem_InheritsFrom,
  TaskDialogBaseButtonItem_Equals,
  TaskDialogBaseButtonItem_GetHashCode,
  TaskDialogBaseButtonItem_ToString,
  TaskDialogBaseButtonItem_GetModalResult,
  TaskDialogBaseButtonItem_SetModalResult,
  TaskDialogBaseButtonItem_GetCaption,
  TaskDialogBaseButtonItem_SetCaption,
  TaskDialogBaseButtonItem_GetDefault,
  TaskDialogBaseButtonItem_SetDefault,
  TaskDialogBaseButtonItem_GetEnabled,
  TaskDialogBaseButtonItem_SetEnabled,
  TaskDialogBaseButtonItem_GetCollection,
  TaskDialogBaseButtonItem_SetCollection,
  TaskDialogBaseButtonItem_GetIndex,
  TaskDialogBaseButtonItem_SetIndex,
  TaskDialogBaseButtonItem_GetDisplayName,
  TaskDialogBaseButtonItem_SetDisplayName,
  TaskDialogBaseButtonItem_StaticClassType;

