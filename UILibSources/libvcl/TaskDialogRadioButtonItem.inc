
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function TaskDialogRadioButtonItem_Create(AOwner: TCollection): TTaskDialogRadioButtonItem; stdcall;
begin
  Result :=  TTaskDialogRadioButtonItem.Create(AOwner);
end;

procedure TaskDialogRadioButtonItem_Free(AObj: TTaskDialogRadioButtonItem); stdcall;
begin
  AObj.Free;
end;

procedure TaskDialogRadioButtonItem_Click(AObj: TTaskDialogRadioButtonItem); stdcall;
begin
  AObj.Click;
end;

function TaskDialogRadioButtonItem_GetNamePath(AObj: TTaskDialogRadioButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TaskDialogRadioButtonItem_Assign(AObj: TTaskDialogRadioButtonItem; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure TaskDialogRadioButtonItem_DisposeOf(AObj: TTaskDialogRadioButtonItem); stdcall;
begin
  AObj.DisposeOf;
end;

function TaskDialogRadioButtonItem_ClassType(AObj: TTaskDialogRadioButtonItem): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function TaskDialogRadioButtonItem_ClassName(AObj: TTaskDialogRadioButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TaskDialogRadioButtonItem_InstanceSize(AObj: TTaskDialogRadioButtonItem): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function TaskDialogRadioButtonItem_InheritsFrom(AObj: TTaskDialogRadioButtonItem; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function TaskDialogRadioButtonItem_Equals(AObj: TTaskDialogRadioButtonItem; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function TaskDialogRadioButtonItem_GetHashCode(AObj: TTaskDialogRadioButtonItem): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function TaskDialogRadioButtonItem_ToString(AObj: TTaskDialogRadioButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function TaskDialogRadioButtonItem_GetModalResult(AObj: TTaskDialogRadioButtonItem): TModalResult; stdcall;
begin
  Result :=  AObj.ModalResult;
end;

procedure TaskDialogRadioButtonItem_SetModalResult(AObj: TTaskDialogRadioButtonItem; AValue: TModalResult); stdcall;
begin
  AObj.ModalResult := AValue;
end;

function TaskDialogRadioButtonItem_GetCaption(AObj: TTaskDialogRadioButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure TaskDialogRadioButtonItem_SetCaption(AObj: TTaskDialogRadioButtonItem; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function TaskDialogRadioButtonItem_GetDefault(AObj: TTaskDialogRadioButtonItem): LongBool; stdcall;
begin
  Result :=  AObj.Default;
end;

procedure TaskDialogRadioButtonItem_SetDefault(AObj: TTaskDialogRadioButtonItem; AValue: LongBool); stdcall;
begin
  AObj.Default := AValue;
end;

function TaskDialogRadioButtonItem_GetEnabled(AObj: TTaskDialogRadioButtonItem): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure TaskDialogRadioButtonItem_SetEnabled(AObj: TTaskDialogRadioButtonItem; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function TaskDialogRadioButtonItem_GetCollection(AObj: TTaskDialogRadioButtonItem): TCollection; stdcall;
begin
  Result :=  AObj.Collection;
end;

procedure TaskDialogRadioButtonItem_SetCollection(AObj: TTaskDialogRadioButtonItem; AValue: TCollection); stdcall;
begin
  AObj.Collection := AValue;
end;

function TaskDialogRadioButtonItem_GetIndex(AObj: TTaskDialogRadioButtonItem): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure TaskDialogRadioButtonItem_SetIndex(AObj: TTaskDialogRadioButtonItem; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;

function TaskDialogRadioButtonItem_GetDisplayName(AObj: TTaskDialogRadioButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.DisplayName);
end;

procedure TaskDialogRadioButtonItem_SetDisplayName(AObj: TTaskDialogRadioButtonItem; AValue: PChar); stdcall;
begin
  AObj.DisplayName := AValue;
end;

function TaskDialogRadioButtonItem_StaticClassType: TClass; stdcall;
begin
  Result := TTaskDialogRadioButtonItem;
end;

exports
  TaskDialogRadioButtonItem_Create,
  TaskDialogRadioButtonItem_Free,
  TaskDialogRadioButtonItem_Click,
  TaskDialogRadioButtonItem_GetNamePath,
  TaskDialogRadioButtonItem_Assign,
  TaskDialogRadioButtonItem_DisposeOf,
  TaskDialogRadioButtonItem_ClassType,
  TaskDialogRadioButtonItem_ClassName,
  TaskDialogRadioButtonItem_InstanceSize,
  TaskDialogRadioButtonItem_InheritsFrom,
  TaskDialogRadioButtonItem_Equals,
  TaskDialogRadioButtonItem_GetHashCode,
  TaskDialogRadioButtonItem_ToString,
  TaskDialogRadioButtonItem_GetModalResult,
  TaskDialogRadioButtonItem_SetModalResult,
  TaskDialogRadioButtonItem_GetCaption,
  TaskDialogRadioButtonItem_SetCaption,
  TaskDialogRadioButtonItem_GetDefault,
  TaskDialogRadioButtonItem_SetDefault,
  TaskDialogRadioButtonItem_GetEnabled,
  TaskDialogRadioButtonItem_SetEnabled,
  TaskDialogRadioButtonItem_GetCollection,
  TaskDialogRadioButtonItem_SetCollection,
  TaskDialogRadioButtonItem_GetIndex,
  TaskDialogRadioButtonItem_SetIndex,
  TaskDialogRadioButtonItem_GetDisplayName,
  TaskDialogRadioButtonItem_SetDisplayName,
  TaskDialogRadioButtonItem_StaticClassType;

