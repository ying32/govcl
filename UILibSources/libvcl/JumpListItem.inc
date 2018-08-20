
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function JumpListItem_Create(AOwner: TCollection): TJumpListItem; stdcall;
begin
  Result :=  TJumpListItem.Create(AOwner);
end;

procedure JumpListItem_Free(AObj: TJumpListItem); stdcall;
begin
  AObj.Free;
end;

function JumpListItem_GetNamePath(AObj: TJumpListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure JumpListItem_Assign(AObj: TJumpListItem; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure JumpListItem_DisposeOf(AObj: TJumpListItem); stdcall;
begin
  AObj.DisposeOf;
end;

function JumpListItem_ClassType(AObj: TJumpListItem): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function JumpListItem_ClassName(AObj: TJumpListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function JumpListItem_InstanceSize(AObj: TJumpListItem): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function JumpListItem_InheritsFrom(AObj: TJumpListItem; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function JumpListItem_Equals(AObj: TJumpListItem; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function JumpListItem_GetHashCode(AObj: TJumpListItem): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function JumpListItem_ToString(AObj: TJumpListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function JumpListItem_GetIsSeparator(AObj: TJumpListItem): LongBool; stdcall;
begin
  Result :=  AObj.IsSeparator;
end;

procedure JumpListItem_SetIsSeparator(AObj: TJumpListItem; AValue: LongBool); stdcall;
begin
  AObj.IsSeparator := AValue;
end;

function JumpListItem_GetIcon(AObj: TJumpListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.Icon);
end;

procedure JumpListItem_SetIcon(AObj: TJumpListItem; AValue: PChar); stdcall;
begin
  AObj.Icon := AValue;
end;

function JumpListItem_GetPath(AObj: TJumpListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.Path);
end;

procedure JumpListItem_SetPath(AObj: TJumpListItem; AValue: PChar); stdcall;
begin
  AObj.Path := AValue;
end;

function JumpListItem_GetFriendlyName(AObj: TJumpListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.FriendlyName);
end;

procedure JumpListItem_SetFriendlyName(AObj: TJumpListItem; AValue: PChar); stdcall;
begin
  AObj.FriendlyName := AValue;
end;

function JumpListItem_GetArguments(AObj: TJumpListItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.Arguments);
end;

procedure JumpListItem_SetArguments(AObj: TJumpListItem; AValue: PChar); stdcall;
begin
  AObj.Arguments := AValue;
end;

function JumpListItem_GetVisible(AObj: TJumpListItem): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure JumpListItem_SetVisible(AObj: TJumpListItem; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function JumpListItem_GetCollection(AObj: TJumpListItem): TCollection; stdcall;
begin
  Result :=  AObj.Collection;
end;

procedure JumpListItem_SetCollection(AObj: TJumpListItem; AValue: TCollection); stdcall;
begin
  AObj.Collection := AValue;
end;

function JumpListItem_GetIndex(AObj: TJumpListItem): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure JumpListItem_SetIndex(AObj: TJumpListItem; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;

function JumpListItem_StaticClassType: TClass; stdcall;
begin
  Result := TJumpListItem;
end;

exports
  JumpListItem_Create,
  JumpListItem_Free,
  JumpListItem_GetNamePath,
  JumpListItem_Assign,
  JumpListItem_DisposeOf,
  JumpListItem_ClassType,
  JumpListItem_ClassName,
  JumpListItem_InstanceSize,
  JumpListItem_InheritsFrom,
  JumpListItem_Equals,
  JumpListItem_GetHashCode,
  JumpListItem_ToString,
  JumpListItem_GetIsSeparator,
  JumpListItem_SetIsSeparator,
  JumpListItem_GetIcon,
  JumpListItem_SetIcon,
  JumpListItem_GetPath,
  JumpListItem_SetPath,
  JumpListItem_GetFriendlyName,
  JumpListItem_SetFriendlyName,
  JumpListItem_GetArguments,
  JumpListItem_SetArguments,
  JumpListItem_GetVisible,
  JumpListItem_SetVisible,
  JumpListItem_GetCollection,
  JumpListItem_SetCollection,
  JumpListItem_GetIndex,
  JumpListItem_SetIndex,
  JumpListItem_StaticClassType;

