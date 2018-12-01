
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ThumbBarButton_Create(AOwner: TCollection): TThumbBarButton; stdcall;
begin
  Result :=  TThumbBarButton.Create(AOwner);
end;

procedure ThumbBarButton_Free(AObj: TThumbBarButton); stdcall;
begin
  AObj.Free;
end;

function ThumbBarButton_GetNamePath(AObj: TThumbBarButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ThumbBarButton_Assign(AObj: TThumbBarButton; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ThumbBarButton_DisposeOf(AObj: TThumbBarButton); stdcall;
begin
  AObj.DisposeOf;
end;

function ThumbBarButton_ClassType(AObj: TThumbBarButton): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function ThumbBarButton_ClassName(AObj: TThumbBarButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ThumbBarButton_InstanceSize(AObj: TThumbBarButton): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function ThumbBarButton_InheritsFrom(AObj: TThumbBarButton; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function ThumbBarButton_Equals(AObj: TThumbBarButton; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ThumbBarButton_GetHashCode(AObj: TThumbBarButton): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ThumbBarButton_ToString(AObj: TThumbBarButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ThumbBarButton_GetAction(AObj: TThumbBarButton): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ThumbBarButton_SetAction(AObj: TThumbBarButton; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function ThumbBarButton_GetButtonState(AObj: TThumbBarButton): TThumbButtonStates; stdcall;
begin
  Result :=  AObj.ButtonState;
end;

procedure ThumbBarButton_SetButtonState(AObj: TThumbBarButton; AValue: TThumbButtonStates); stdcall;
begin
  AObj.ButtonState := AValue;
end;

function ThumbBarButton_GetHint(AObj: TThumbBarButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ThumbBarButton_SetHint(AObj: TThumbBarButton; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ThumbBarButton_GetIcon(AObj: TThumbBarButton): TIcon; stdcall;
begin
  Result :=  AObj.Icon;
end;

procedure ThumbBarButton_SetIcon(AObj: TThumbBarButton; AValue: TIcon); stdcall;
begin
  AObj.Icon := AValue;
end;

function ThumbBarButton_GetCollection(AObj: TThumbBarButton): TCollection; stdcall;
begin
  Result :=  AObj.Collection;
end;

procedure ThumbBarButton_SetCollection(AObj: TThumbBarButton; AValue: TCollection); stdcall;
begin
  AObj.Collection := AValue;
end;

function ThumbBarButton_GetIndex(AObj: TThumbBarButton): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure ThumbBarButton_SetIndex(AObj: TThumbBarButton; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;

function ThumbBarButton_StaticClassType: TClass; stdcall;
begin
  Result := TThumbBarButton;
end;

exports
  ThumbBarButton_Create,
  ThumbBarButton_Free,
  ThumbBarButton_GetNamePath,
  ThumbBarButton_Assign,
  ThumbBarButton_DisposeOf,
  ThumbBarButton_ClassType,
  ThumbBarButton_ClassName,
  ThumbBarButton_InstanceSize,
  ThumbBarButton_InheritsFrom,
  ThumbBarButton_Equals,
  ThumbBarButton_GetHashCode,
  ThumbBarButton_ToString,
  ThumbBarButton_GetAction,
  ThumbBarButton_SetAction,
  ThumbBarButton_GetButtonState,
  ThumbBarButton_SetButtonState,
  ThumbBarButton_GetHint,
  ThumbBarButton_SetHint,
  ThumbBarButton_GetIcon,
  ThumbBarButton_SetIcon,
  ThumbBarButton_GetCollection,
  ThumbBarButton_SetCollection,
  ThumbBarButton_GetIndex,
  ThumbBarButton_SetIndex,
  ThumbBarButton_StaticClassType;

