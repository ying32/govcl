
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

function Object_ClassName(AObj: TObject): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
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


exports
  Object_Create,
  Object_Free,
  Object_ClassName,
  Object_Equals,
  Object_GetHashCode,
  Object_ToString;

