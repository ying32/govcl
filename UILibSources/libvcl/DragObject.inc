
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function DragObject_Create: TDragObject; stdcall;
begin
  Result :=  TDragObject.Create;
end;

procedure DragObject_Free(AObj: TDragObject); stdcall;
begin
  AObj.Free;
end;

procedure DragObject_Assign(AObj: TDragObject; Source: TDragObject); stdcall;
begin
  AObj.Assign(Source);
end;

procedure DragObject_HideDragImage(AObj: TDragObject); stdcall;
begin
  AObj.HideDragImage;
end;

procedure DragObject_ShowDragImage(AObj: TDragObject); stdcall;
begin
  AObj.ShowDragImage;
end;

procedure DragObject_DisposeOf(AObj: TDragObject); stdcall;
begin
  AObj.DisposeOf;
end;

function DragObject_ClassType(AObj: TDragObject): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function DragObject_ClassName(AObj: TDragObject): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function DragObject_InstanceSize(AObj: TDragObject): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function DragObject_InheritsFrom(AObj: TDragObject; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function DragObject_Equals(AObj: TDragObject; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function DragObject_GetHashCode(AObj: TDragObject): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function DragObject_ToString(AObj: TDragObject): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function DragObject_GetAlwaysShowDragImages(AObj: TDragObject): LongBool; stdcall;
begin
  Result :=  AObj.AlwaysShowDragImages;
end;

procedure DragObject_SetAlwaysShowDragImages(AObj: TDragObject; AValue: LongBool); stdcall;
begin
  AObj.AlwaysShowDragImages := AValue;
end;

function DragObject_GetCancelling(AObj: TDragObject): LongBool; stdcall;
begin
  Result :=  AObj.Cancelling;
end;

procedure DragObject_SetCancelling(AObj: TDragObject; AValue: LongBool); stdcall;
begin
  AObj.Cancelling := AValue;
end;

function DragObject_GetDragHandle(AObj: TDragObject): HWND; stdcall;
begin
  Result :=  AObj.DragHandle;
end;

procedure DragObject_SetDragHandle(AObj: TDragObject; AValue: HWND); stdcall;
begin
  AObj.DragHandle := AValue;
end;

procedure DragObject_GetDragPos(AObj: TDragObject; var Result: TPoint); stdcall;
begin
  Result :=  AObj.DragPos;
end;

procedure DragObject_SetDragPos(AObj: TDragObject; var AValue: TPoint); stdcall;
begin
  AObj.DragPos := AValue;
end;

function DragObject_GetDragTarget(AObj: TDragObject): Pointer; stdcall;
begin
  Result :=  AObj.DragTarget;
end;

procedure DragObject_SetDragTarget(AObj: TDragObject; AValue: Pointer); stdcall;
begin
  AObj.DragTarget := AValue;
end;

procedure DragObject_GetDragTargetPos(AObj: TDragObject; var Result: TPoint); stdcall;
begin
  Result :=  AObj.DragTargetPos;
end;

procedure DragObject_SetDragTargetPos(AObj: TDragObject; var AValue: TPoint); stdcall;
begin
  AObj.DragTargetPos := AValue;
end;

function DragObject_GetDropped(AObj: TDragObject): LongBool; stdcall;
begin
  Result :=  AObj.Dropped;
end;

procedure DragObject_GetMouseDeltaX(AObj: TDragObject; var Result: Double); stdcall;
begin
  Result :=  AObj.MouseDeltaX;
end;

procedure DragObject_GetMouseDeltaY(AObj: TDragObject; var Result: Double); stdcall;
begin
  Result :=  AObj.MouseDeltaY;
end;

function DragObject_GetRightClickCancels(AObj: TDragObject): LongBool; stdcall;
begin
  Result :=  AObj.RightClickCancels;
end;

procedure DragObject_SetRightClickCancels(AObj: TDragObject; AValue: LongBool); stdcall;
begin
  AObj.RightClickCancels := AValue;
end;

function DragObject_StaticClassType: TClass; stdcall;
begin
  Result := TDragObject;
end;

exports
  DragObject_Create,
  DragObject_Free,
  DragObject_Assign,
  DragObject_HideDragImage,
  DragObject_ShowDragImage,
  DragObject_DisposeOf,
  DragObject_ClassType,
  DragObject_ClassName,
  DragObject_InstanceSize,
  DragObject_InheritsFrom,
  DragObject_Equals,
  DragObject_GetHashCode,
  DragObject_ToString,
  DragObject_GetAlwaysShowDragImages,
  DragObject_SetAlwaysShowDragImages,
  DragObject_GetCancelling,
  DragObject_SetCancelling,
  DragObject_GetDragHandle,
  DragObject_SetDragHandle,
  DragObject_GetDragPos,
  DragObject_SetDragPos,
  DragObject_GetDragTarget,
  DragObject_SetDragTarget,
  DragObject_GetDragTargetPos,
  DragObject_SetDragTargetPos,
  DragObject_GetDropped,
  DragObject_GetMouseDeltaX,
  DragObject_GetMouseDeltaY,
  DragObject_GetRightClickCancels,
  DragObject_SetRightClickCancels,
  DragObject_StaticClassType;

