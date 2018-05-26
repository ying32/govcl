
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Component_Create(AOwner: TComponent): TComponent; extdecl;
begin
  Result :=  TComponent.Create(AOwner);
end;

procedure Component_Free(AObj: TComponent); extdecl;
begin
  AObj.Free;
end;

function Component_FindComponent(AObj: TComponent; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Component_GetNamePath(AObj: TComponent): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function Component_HasParent(AObj: TComponent): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Component_Assign(AObj: TComponent; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Component_ClassName(AObj: TComponent): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Component_Equals(AObj: TComponent; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Component_GetHashCode(AObj: TComponent): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Component_ToString(AObj: TComponent): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Component_GetComponentCount(AObj: TComponent): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Component_GetComponentIndex(AObj: TComponent): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Component_SetComponentIndex(AObj: TComponent; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Component_GetOwner(AObj: TComponent): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Component_GetName(AObj: TComponent): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Component_SetName(AObj: TComponent; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Component_GetTag(AObj: TComponent): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Component_SetTag(AObj: TComponent; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Component_GetComponents(AObj: TComponent; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Component_Create,
  Component_Free,
  Component_FindComponent,
  Component_GetNamePath,
  Component_HasParent,
  Component_Assign,
  Component_ClassName,
  Component_Equals,
  Component_GetHashCode,
  Component_ToString,
  Component_GetComponentCount,
  Component_GetComponentIndex,
  Component_SetComponentIndex,
  Component_GetOwner,
  Component_GetName,
  Component_SetName,
  Component_GetTag,
  Component_SetTag,
  Component_GetComponents;

