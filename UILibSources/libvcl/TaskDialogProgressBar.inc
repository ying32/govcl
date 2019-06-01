
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function TaskDialogProgressBar_Create(AOwner: TCustomTaskDialog): TTaskDialogProgressBar; stdcall;
begin
  Result :=  TTaskDialogProgressBar.Create(AOwner);
end;

procedure TaskDialogProgressBar_Free(AObj: TTaskDialogProgressBar); stdcall;
begin
  AObj.Free;
end;

procedure TaskDialogProgressBar_Assign(AObj: TTaskDialogProgressBar; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function TaskDialogProgressBar_GetNamePath(AObj: TTaskDialogProgressBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TaskDialogProgressBar_DisposeOf(AObj: TTaskDialogProgressBar); stdcall;
begin
  AObj.DisposeOf;
end;

function TaskDialogProgressBar_ClassType(AObj: TTaskDialogProgressBar): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function TaskDialogProgressBar_ClassName(AObj: TTaskDialogProgressBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TaskDialogProgressBar_InstanceSize(AObj: TTaskDialogProgressBar): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function TaskDialogProgressBar_InheritsFrom(AObj: TTaskDialogProgressBar; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function TaskDialogProgressBar_Equals(AObj: TTaskDialogProgressBar; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function TaskDialogProgressBar_GetHashCode(AObj: TTaskDialogProgressBar): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function TaskDialogProgressBar_ToString(AObj: TTaskDialogProgressBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function TaskDialogProgressBar_GetMax(AObj: TTaskDialogProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Max;
end;

procedure TaskDialogProgressBar_SetMax(AObj: TTaskDialogProgressBar; AValue: Integer); stdcall;
begin
  AObj.Max := AValue;
end;

function TaskDialogProgressBar_GetMin(AObj: TTaskDialogProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Min;
end;

procedure TaskDialogProgressBar_SetMin(AObj: TTaskDialogProgressBar; AValue: Integer); stdcall;
begin
  AObj.Min := AValue;
end;

function TaskDialogProgressBar_GetPosition(AObj: TTaskDialogProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Position;
end;

procedure TaskDialogProgressBar_SetPosition(AObj: TTaskDialogProgressBar; AValue: Integer); stdcall;
begin
  AObj.Position := AValue;
end;

function TaskDialogProgressBar_GetState(AObj: TTaskDialogProgressBar): TProgressBarState; stdcall;
begin
  Result :=  AObj.State;
end;

procedure TaskDialogProgressBar_SetState(AObj: TTaskDialogProgressBar; AValue: TProgressBarState); stdcall;
begin
  AObj.State := AValue;
end;

function TaskDialogProgressBar_StaticClassType: TClass; stdcall;
begin
  Result := TTaskDialogProgressBar;
end;

exports
  TaskDialogProgressBar_Create,
  TaskDialogProgressBar_Free,
  TaskDialogProgressBar_Assign,
  TaskDialogProgressBar_GetNamePath,
  TaskDialogProgressBar_DisposeOf,
  TaskDialogProgressBar_ClassType,
  TaskDialogProgressBar_ClassName,
  TaskDialogProgressBar_InstanceSize,
  TaskDialogProgressBar_InheritsFrom,
  TaskDialogProgressBar_Equals,
  TaskDialogProgressBar_GetHashCode,
  TaskDialogProgressBar_ToString,
  TaskDialogProgressBar_GetMax,
  TaskDialogProgressBar_SetMax,
  TaskDialogProgressBar_GetMin,
  TaskDialogProgressBar_SetMin,
  TaskDialogProgressBar_GetPosition,
  TaskDialogProgressBar_SetPosition,
  TaskDialogProgressBar_GetState,
  TaskDialogProgressBar_SetState,
  TaskDialogProgressBar_StaticClassType;

