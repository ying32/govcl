
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function CollectionItem_Create(AOwner: TCollection): TCollectionItem; extdecl;
begin
  Result :=  TCollectionItem.Create(AOwner);
end;

procedure CollectionItem_Free(AObj: TCollectionItem); extdecl;
begin
  AObj.Free;
end;

function CollectionItem_GetNamePath(AObj: TCollectionItem): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure CollectionItem_Assign(AObj: TCollectionItem; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function CollectionItem_ClassName(AObj: TCollectionItem): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function CollectionItem_Equals(AObj: TCollectionItem; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function CollectionItem_GetHashCode(AObj: TCollectionItem): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function CollectionItem_ToString(AObj: TCollectionItem): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function CollectionItem_GetIndex(AObj: TCollectionItem): Integer; extdecl;
begin
  Result :=  AObj.Index;
end;

procedure CollectionItem_SetIndex(AObj: TCollectionItem; AValue: Integer); extdecl;
begin
  AObj.Index := AValue;
end;


exports
  CollectionItem_Create,
  CollectionItem_Free,
  CollectionItem_GetNamePath,
  CollectionItem_Assign,
  CollectionItem_ClassName,
  CollectionItem_Equals,
  CollectionItem_GetHashCode,
  CollectionItem_ToString,
  CollectionItem_GetIndex,
  CollectionItem_SetIndex;

