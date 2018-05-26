procedure ParaAttributes_Assign(AObj: TParaAttributes; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ParaAttributes_GetNamePath(AObj: TParaAttributes): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function ParaAttributes_ClassName(AObj: TParaAttributes): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ParaAttributes_Equals(AObj: TParaAttributes; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ParaAttributes_GetHashCode(AObj: TParaAttributes): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ParaAttributes_ToString(AObj: TParaAttributes): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ParaAttributes_GetAlignment(AObj: TParaAttributes): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure ParaAttributes_SetAlignment(AObj: TParaAttributes; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function ParaAttributes_GetFirstIndent(AObj: TParaAttributes): Integer; stdcall;
begin
  Result :=  AObj.FirstIndent;
end;

procedure ParaAttributes_SetFirstIndent(AObj: TParaAttributes; AValue: Integer); stdcall;
begin
  AObj.FirstIndent := AValue;
end;

function ParaAttributes_GetLeftIndent(AObj: TParaAttributes): Integer; stdcall;
begin
  Result :=  AObj.LeftIndent;
end;

procedure ParaAttributes_SetLeftIndent(AObj: TParaAttributes; AValue: Integer); stdcall;
begin
  AObj.LeftIndent := AValue;
end;

function ParaAttributes_GetNumbering(AObj: TParaAttributes): TNumberingStyle; stdcall;
begin
  Result :=  AObj.Numbering;
end;

procedure ParaAttributes_SetNumbering(AObj: TParaAttributes; AValue: TNumberingStyle); stdcall;
begin
  AObj.Numbering := AValue;
end;

function ParaAttributes_GetRightIndent(AObj: TParaAttributes): Integer; stdcall;
begin
  Result :=  AObj.RightIndent;
end;

procedure ParaAttributes_SetRightIndent(AObj: TParaAttributes; AValue: Integer); stdcall;
begin
  AObj.RightIndent := AValue;
end;

function ParaAttributes_GetTabCount(AObj: TParaAttributes): Integer; stdcall;
begin
  Result :=  AObj.TabCount;
end;

procedure ParaAttributes_SetTabCount(AObj: TParaAttributes; AValue: Integer); stdcall;
begin
  AObj.TabCount := AValue;
end;

function ParaAttributes_GetTab(AObj: TParaAttributes; Index: Byte): Integer; stdcall;
begin
  Result :=  AObj.Tab[Index];
end;

procedure ParaAttributes_SetTab(AObj: TParaAttributes; Index: Byte; AValue: Integer); stdcall;
begin
  AObj.Tab[Index] := AValue;
end;


exports
  ParaAttributes_Assign,
  ParaAttributes_GetNamePath,
  ParaAttributes_ClassName,
  ParaAttributes_Equals,
  ParaAttributes_GetHashCode,
  ParaAttributes_ToString,
  ParaAttributes_GetAlignment,
  ParaAttributes_SetAlignment,
  ParaAttributes_GetFirstIndent,
  ParaAttributes_SetFirstIndent,
  ParaAttributes_GetLeftIndent,
  ParaAttributes_SetLeftIndent,
  ParaAttributes_GetNumbering,
  ParaAttributes_SetNumbering,
  ParaAttributes_GetRightIndent,
  ParaAttributes_SetRightIndent,
  ParaAttributes_GetTabCount,
  ParaAttributes_SetTabCount,
  ParaAttributes_GetTab,
  ParaAttributes_SetTab;

