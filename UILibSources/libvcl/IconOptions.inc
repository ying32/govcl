procedure IconOptions_Assign(AObj: TIconOptions; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function IconOptions_GetNamePath(AObj: TIconOptions): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function IconOptions_ClassName(AObj: TIconOptions): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function IconOptions_Equals(AObj: TIconOptions; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function IconOptions_GetHashCode(AObj: TIconOptions): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function IconOptions_ToString(AObj: TIconOptions): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function IconOptions_GetArrangement(AObj: TIconOptions): TIconArrangement; stdcall;
begin
  Result :=  AObj.Arrangement;
end;

procedure IconOptions_SetArrangement(AObj: TIconOptions; AValue: TIconArrangement); stdcall;
begin
  AObj.Arrangement := AValue;
end;

function IconOptions_GetAutoArrange(AObj: TIconOptions): LongBool; stdcall;
begin
  Result :=  AObj.AutoArrange;
end;

procedure IconOptions_SetAutoArrange(AObj: TIconOptions; AValue: LongBool); stdcall;
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

