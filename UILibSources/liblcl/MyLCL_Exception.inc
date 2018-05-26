function Exception_ToString(AObj: Exception): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Exception_ClassName(AObj: Exception): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Exception_Equals(AObj: Exception; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Exception_GetHashCode(AObj: Exception): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Exception_GetMessage(AObj: Exception): PChar; extdecl;
begin
  Result :=  PChar(AObj.Message);
end;

procedure Exception_SetMessage(AObj: Exception; AValue: PChar); extdecl;
begin
  AObj.Message := AValue;
end;


exports
  Exception_ToString,
  Exception_ClassName,
  Exception_Equals,
  Exception_GetHashCode,
  Exception_GetMessage,
  Exception_SetMessage;

