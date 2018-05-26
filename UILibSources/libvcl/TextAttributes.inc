procedure TextAttributes_Assign(AObj: TTextAttributes; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function TextAttributes_GetNamePath(AObj: TTextAttributes): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function TextAttributes_ClassName(AObj: TTextAttributes): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TextAttributes_Equals(AObj: TTextAttributes; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function TextAttributes_GetHashCode(AObj: TTextAttributes): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function TextAttributes_ToString(AObj: TTextAttributes): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function TextAttributes_GetCharset(AObj: TTextAttributes): TFontCharset; stdcall;
begin
  Result :=  AObj.Charset;
end;

procedure TextAttributes_SetCharset(AObj: TTextAttributes; AValue: TFontCharset); stdcall;
begin
  AObj.Charset := AValue;
end;

function TextAttributes_GetColor(AObj: TTextAttributes): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure TextAttributes_SetColor(AObj: TTextAttributes; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function TextAttributes_GetConsistentAttributes(AObj: TTextAttributes): TConsistentAttributes; stdcall;
begin
  Result :=  AObj.ConsistentAttributes;
end;

function TextAttributes_GetName(AObj: TTextAttributes): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TextAttributes_SetName(AObj: TTextAttributes; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function TextAttributes_GetPitch(AObj: TTextAttributes): TFontPitch; stdcall;
begin
  Result :=  AObj.Pitch;
end;

procedure TextAttributes_SetPitch(AObj: TTextAttributes; AValue: TFontPitch); stdcall;
begin
  AObj.Pitch := AValue;
end;

function TextAttributes_GetProtected(AObj: TTextAttributes): LongBool; stdcall;
begin
  Result :=  AObj.Protected;
end;

procedure TextAttributes_SetProtected(AObj: TTextAttributes; AValue: LongBool); stdcall;
begin
  AObj.Protected := AValue;
end;

function TextAttributes_GetSize(AObj: TTextAttributes): Integer; stdcall;
begin
  Result :=  AObj.Size;
end;

procedure TextAttributes_SetSize(AObj: TTextAttributes; AValue: Integer); stdcall;
begin
  AObj.Size := AValue;
end;

function TextAttributes_GetStyle(AObj: TTextAttributes): TFontStyles; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure TextAttributes_SetStyle(AObj: TTextAttributes; AValue: TFontStyles); stdcall;
begin
  AObj.Style := AValue;
end;

function TextAttributes_GetHeight(AObj: TTextAttributes): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure TextAttributes_SetHeight(AObj: TTextAttributes; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;


exports
  TextAttributes_Assign,
  TextAttributes_GetNamePath,
  TextAttributes_ClassName,
  TextAttributes_Equals,
  TextAttributes_GetHashCode,
  TextAttributes_ToString,
  TextAttributes_GetCharset,
  TextAttributes_SetCharset,
  TextAttributes_GetColor,
  TextAttributes_SetColor,
  TextAttributes_GetConsistentAttributes,
  TextAttributes_GetName,
  TextAttributes_SetName,
  TextAttributes_GetPitch,
  TextAttributes_SetPitch,
  TextAttributes_GetProtected,
  TextAttributes_SetProtected,
  TextAttributes_GetSize,
  TextAttributes_SetSize,
  TextAttributes_GetStyle,
  TextAttributes_SetStyle,
  TextAttributes_GetHeight,
  TextAttributes_SetHeight;

