function Exception_ToString(AObj: Exception): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Exception_ClassName(AObj: Exception): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Exception_Equals(AObj: Exception; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Exception_GetHashCode(AObj: Exception): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Exception_GetBaseException(AObj: Exception): Exception; stdcall;
begin
  Result :=  AObj.BaseException;
end;

function Exception_GetInnerException(AObj: Exception): Exception; stdcall;
begin
  Result :=  AObj.InnerException;
end;

function Exception_GetMessage(AObj: Exception): PChar; stdcall;
begin
  Result :=  PChar(AObj.Message);
end;

procedure Exception_SetMessage(AObj: Exception; AValue: PChar); stdcall;
begin
  AObj.Message := AValue;
end;

function Exception_GetStackTrace(AObj: Exception): PChar; stdcall;
begin
  Result :=  PChar(AObj.StackTrace);
end;

function Exception_GetStackInfo(AObj: Exception): Pointer; stdcall;
begin
  Result :=  AObj.StackInfo;
end;


exports
  Exception_ToString,
  Exception_ClassName,
  Exception_Equals,
  Exception_GetHashCode,
  Exception_GetBaseException,
  Exception_GetInnerException,
  Exception_GetMessage,
  Exception_SetMessage,
  Exception_GetStackTrace,
  Exception_GetStackInfo;

