
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function TaskDialogButtonItem_Create(AOwner: TCollection): TTaskDialogButtonItem; stdcall;
begin
  Result :=  TTaskDialogButtonItem.Create(AOwner);
end;

procedure TaskDialogButtonItem_Free(AObj: TTaskDialogButtonItem); stdcall;
begin
  AObj.Free;
end;

procedure TaskDialogButtonItem_Click(AObj: TTaskDialogButtonItem); stdcall;
begin
  AObj.Click;
end;

function TaskDialogButtonItem_GetNamePath(AObj: TTaskDialogButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TaskDialogButtonItem_Assign(AObj: TTaskDialogButtonItem; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure TaskDialogButtonItem_DisposeOf(AObj: TTaskDialogButtonItem); stdcall;
begin
  AObj.DisposeOf;
end;

function TaskDialogButtonItem_ClassType(AObj: TTaskDialogButtonItem): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function TaskDialogButtonItem_ClassName(AObj: TTaskDialogButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TaskDialogButtonItem_InstanceSize(AObj: TTaskDialogButtonItem): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function TaskDialogButtonItem_InheritsFrom(AObj: TTaskDialogButtonItem; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function TaskDialogButtonItem_Equals(AObj: TTaskDialogButtonItem; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function TaskDialogButtonItem_GetHashCode(AObj: TTaskDialogButtonItem): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function TaskDialogButtonItem_ToString(AObj: TTaskDialogButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function TaskDialogButtonItem_GetCommandLinkHint(AObj: TTaskDialogButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.CommandLinkHint);
end;

procedure TaskDialogButtonItem_SetCommandLinkHint(AObj: TTaskDialogButtonItem; AValue: PChar); stdcall;
begin
  AObj.CommandLinkHint := AValue;
end;

function TaskDialogButtonItem_GetElevationRequired(AObj: TTaskDialogButtonItem): LongBool; stdcall;
begin
  Result :=  AObj.ElevationRequired;
end;

procedure TaskDialogButtonItem_SetElevationRequired(AObj: TTaskDialogButtonItem; AValue: LongBool); stdcall;
begin
  AObj.ElevationRequired := AValue;
end;

function TaskDialogButtonItem_GetModalResult(AObj: TTaskDialogButtonItem): TModalResult; stdcall;
begin
  Result :=  AObj.ModalResult;
end;

procedure TaskDialogButtonItem_SetModalResult(AObj: TTaskDialogButtonItem; AValue: TModalResult); stdcall;
begin
  AObj.ModalResult := AValue;
end;

function TaskDialogButtonItem_GetCaption(AObj: TTaskDialogButtonItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure TaskDialogButtonItem_SetCaption(AObj: TTaskDialogButtonItem; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function TaskDialogButtonItem_GetDefault(AObj: TTaskDialogButtonItem): LongBool; stdcall;
begin
  Result :=  AObj.Default;
end;

procedure TaskDialogButtonItem_SetDefault(AObj: TTaskDialogButtonItem; AValue: LongBool); stdcall;
begin
  AObj.Default := AValue;
end;

function TaskDialogButtonItem_GetEnabled(AObj: TTaskDialogButtonItem): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure TaskDialogButtonItem_SetEnabled(AObj: TTaskDialogButtonItem; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function TaskDialogButtonItem_GetCollection(AObj: TTaskDialogButtonItem): TCollection; stdcall;
begin
  Result :=  AObj.Collection;
end;

procedure TaskDialogButtonItem_SetCollection(AObj: TTaskDialogButtonItem; AValue: TCollection); stdcall;
begin
  AObj.Collection := AValue;
end;

function TaskDialogButtonItem_GetIndex(AObj: TTaskDialogButtonItem): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure TaskDialogButtonItem_SetIndex(AObj: TTaskDialogButtonItem; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;

function TaskDialogButtonItem_StaticClassType: TClass; stdcall;
begin
  Result := TTaskDialogButtonItem;
end;

exports
  TaskDialogButtonItem_Create,
  TaskDialogButtonItem_Free,
  TaskDialogButtonItem_Click,
  TaskDialogButtonItem_GetNamePath,
  TaskDialogButtonItem_Assign,
  TaskDialogButtonItem_DisposeOf,
  TaskDialogButtonItem_ClassType,
  TaskDialogButtonItem_ClassName,
  TaskDialogButtonItem_InstanceSize,
  TaskDialogButtonItem_InheritsFrom,
  TaskDialogButtonItem_Equals,
  TaskDialogButtonItem_GetHashCode,
  TaskDialogButtonItem_ToString,
  TaskDialogButtonItem_GetCommandLinkHint,
  TaskDialogButtonItem_SetCommandLinkHint,
  TaskDialogButtonItem_GetElevationRequired,
  TaskDialogButtonItem_SetElevationRequired,
  TaskDialogButtonItem_GetModalResult,
  TaskDialogButtonItem_SetModalResult,
  TaskDialogButtonItem_GetCaption,
  TaskDialogButtonItem_SetCaption,
  TaskDialogButtonItem_GetDefault,
  TaskDialogButtonItem_SetDefault,
  TaskDialogButtonItem_GetEnabled,
  TaskDialogButtonItem_SetEnabled,
  TaskDialogButtonItem_GetCollection,
  TaskDialogButtonItem_SetCollection,
  TaskDialogButtonItem_GetIndex,
  TaskDialogButtonItem_SetIndex,
  TaskDialogButtonItem_StaticClassType;

