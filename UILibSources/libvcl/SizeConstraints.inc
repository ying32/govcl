procedure SizeConstraints_Assign(AObj: TSizeConstraints; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function SizeConstraints_GetNamePath(AObj: TSizeConstraints): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure SizeConstraints_DisposeOf(AObj: TSizeConstraints); stdcall;
begin
  AObj.DisposeOf;
end;

function SizeConstraints_ClassType(AObj: TSizeConstraints): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function SizeConstraints_ClassName(AObj: TSizeConstraints): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function SizeConstraints_InstanceSize(AObj: TSizeConstraints): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function SizeConstraints_InheritsFrom(AObj: TSizeConstraints; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function SizeConstraints_Equals(AObj: TSizeConstraints; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function SizeConstraints_GetHashCode(AObj: TSizeConstraints): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function SizeConstraints_ToString(AObj: TSizeConstraints): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

procedure SizeConstraints_SetOnChange(AObj: TSizeConstraints; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, @TEventClass.OnChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, @TEventClass.OnChange, AEventId);
end;

function SizeConstraints_GetMaxHeight(AObj: TSizeConstraints): TConstraintSize; stdcall;
begin
  Result :=  AObj.MaxHeight;
end;

procedure SizeConstraints_SetMaxHeight(AObj: TSizeConstraints; AValue: TConstraintSize); stdcall;
begin
  AObj.MaxHeight := AValue;
end;

function SizeConstraints_GetMaxWidth(AObj: TSizeConstraints): TConstraintSize; stdcall;
begin
  Result :=  AObj.MaxWidth;
end;

procedure SizeConstraints_SetMaxWidth(AObj: TSizeConstraints; AValue: TConstraintSize); stdcall;
begin
  AObj.MaxWidth := AValue;
end;

function SizeConstraints_GetMinHeight(AObj: TSizeConstraints): TConstraintSize; stdcall;
begin
  Result :=  AObj.MinHeight;
end;

procedure SizeConstraints_SetMinHeight(AObj: TSizeConstraints; AValue: TConstraintSize); stdcall;
begin
  AObj.MinHeight := AValue;
end;

function SizeConstraints_GetMinWidth(AObj: TSizeConstraints): TConstraintSize; stdcall;
begin
  Result :=  AObj.MinWidth;
end;

procedure SizeConstraints_SetMinWidth(AObj: TSizeConstraints; AValue: TConstraintSize); stdcall;
begin
  AObj.MinWidth := AValue;
end;

function SizeConstraints_StaticClassType: TClass; stdcall;
begin
  Result := TSizeConstraints;
end;

exports
  SizeConstraints_Assign,
  SizeConstraints_GetNamePath,
  SizeConstraints_DisposeOf,
  SizeConstraints_ClassType,
  SizeConstraints_ClassName,
  SizeConstraints_InstanceSize,
  SizeConstraints_InheritsFrom,
  SizeConstraints_Equals,
  SizeConstraints_GetHashCode,
  SizeConstraints_ToString,
  SizeConstraints_SetOnChange,
  SizeConstraints_GetMaxHeight,
  SizeConstraints_SetMaxHeight,
  SizeConstraints_GetMaxWidth,
  SizeConstraints_SetMaxWidth,
  SizeConstraints_GetMinHeight,
  SizeConstraints_SetMinHeight,
  SizeConstraints_GetMinWidth,
  SizeConstraints_SetMinWidth,
  SizeConstraints_StaticClassType;

