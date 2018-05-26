procedure IconOptions_Assign(AObj: TIconOptions; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function IconOptions_GetNamePath(AObj: TIconOptions): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function IconOptions_ClassName(AObj: TIconOptions): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function IconOptions_Equals(AObj: TIconOptions; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function IconOptions_GetHashCode(AObj: TIconOptions): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function IconOptions_ToString(AObj: TIconOptions): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function IconOptions_GetArrangement(AObj: TIconOptions): TIconArrangement; extdecl;
begin
  Result :=  AObj.Arrangement;
end;

procedure IconOptions_SetArrangement(AObj: TIconOptions; AValue: TIconArrangement); extdecl;
begin
  AObj.Arrangement := AValue;
end;

function IconOptions_GetAutoArrange(AObj: TIconOptions): LongBool; extdecl;
begin
  Result :=  AObj.AutoArrange;
end;

procedure IconOptions_SetAutoArrange(AObj: TIconOptions; AValue: LongBool); extdecl;
begin
  AObj.AutoArrange := AValue;
end;


exports
  IconOptions_Assign,
  IconOptions_GetNamePath,
  IconOptions_ClassName,
  IconOptions_Equals,
  IconOptions_GetHashCode,
  IconOptions_ToString,
  IconOptions_GetArrangement,
  IconOptions_SetArrangement,
  IconOptions_GetAutoArrange,
  IconOptions_SetAutoArrange;

