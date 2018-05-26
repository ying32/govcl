
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Object_Create: TObject; extdecl;
begin
  Result :=  TObject.Create;
end;

procedure Object_Free(AObj: TObject); extdecl;
begin
  AObj.Free;
end;

function Object_ClassName(AObj: TObject): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Object_Equals(AObj: TObject; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Object_GetHashCode(AObj: TObject): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Object_ToString(AObj: TObject): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;


exports
  Object_Create,
  Object_Free,
  Object_ClassName,
  Object_Equals,
  Object_GetHashCode,
  Object_ToString;

