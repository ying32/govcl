
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function CollectionItem_Create(AOwner: TCollection): TCollectionItem; stdcall;
begin
  Result :=  TCollectionItem.Create(AOwner);
end;

procedure CollectionItem_Free(AObj: TCollectionItem); stdcall;
begin
  AObj.Free;
end;

function CollectionItem_GetNamePath(AObj: TCollectionItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure CollectionItem_Assign(AObj: TCollectionItem; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure CollectionItem_DisposeOf(AObj: TCollectionItem); stdcall;
begin
  AObj.DisposeOf;
end;

function CollectionItem_ClassType(AObj: TCollectionItem): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function CollectionItem_ClassName(AObj: TCollectionItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function CollectionItem_InstanceSize(AObj: TCollectionItem): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function CollectionItem_InheritsFrom(AObj: TCollectionItem; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function CollectionItem_Equals(AObj: TCollectionItem; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function CollectionItem_GetHashCode(AObj: TCollectionItem): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function CollectionItem_ToString(AObj: TCollectionItem): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function CollectionItem_GetCollection(AObj: TCollectionItem): TCollection; stdcall;
begin
  Result :=  AObj.Collection;
end;

procedure CollectionItem_SetCollection(AObj: TCollectionItem; AValue: TCollection); stdcall;
begin
  AObj.Collection := AValue;
end;

function CollectionItem_GetIndex(AObj: TCollectionItem): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure CollectionItem_SetIndex(AObj: TCollectionItem; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;

function CollectionItem_StaticClassType: TClass; stdcall;
begin
  Result := TCollectionItem;
end;

exports
  CollectionItem_Create,
  CollectionItem_Free,
  CollectionItem_GetNamePath,
  CollectionItem_Assign,
  CollectionItem_DisposeOf,
  CollectionItem_ClassType,
  CollectionItem_ClassName,
  CollectionItem_InstanceSize,
  CollectionItem_InheritsFrom,
  CollectionItem_Equals,
  CollectionItem_GetHashCode,
  CollectionItem_ToString,
  CollectionItem_GetCollection,
  CollectionItem_SetCollection,
  CollectionItem_GetIndex,
  CollectionItem_SetIndex,
  CollectionItem_StaticClassType;

