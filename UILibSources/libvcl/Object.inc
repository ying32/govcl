
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Object_Create: TObject; stdcall;
begin
  Result :=  TObject.Create;
end;

procedure Object_Free(AObj: TObject); stdcall;
begin
  AObj.Free;
end;

procedure Object_DisposeOf(AObj: TObject); stdcall;
begin
  AObj.DisposeOf;
end;

function Object_ClassType(AObj: TObject): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function Object_ClassName(AObj: TObject): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Object_InstanceSize(AObj: TObject): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function Object_InheritsFrom(AObj: TObject; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function Object_Equals(AObj: TObject; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Object_GetHashCode(AObj: TObject): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Object_ToString(AObj: TObject): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Object_StaticClassType: TClass; stdcall;
begin
  Result := TObject;
end;

exports
  Object_Create,
  Object_Free,
  Object_DisposeOf,
  Object_ClassType,
  Object_ClassName,
  Object_InstanceSize,
  Object_InheritsFrom,
  Object_Equals,
  Object_GetHashCode,
  Object_ToString,
  Object_StaticClassType;

