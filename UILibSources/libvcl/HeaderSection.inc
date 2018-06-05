
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function HeaderSection_Create(AOwner: TCollection): THeaderSection; stdcall;
begin
  Result :=  THeaderSection.Create(AOwner);
end;

procedure HeaderSection_Free(AObj: THeaderSection); stdcall;
begin
  AObj.Free;
end;

procedure HeaderSection_Assign(AObj: THeaderSection; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function HeaderSection_GetNamePath(AObj: THeaderSection): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure HeaderSection_DisposeOf(AObj: THeaderSection); stdcall;
begin
  AObj.DisposeOf;
end;

function HeaderSection_ClassType(AObj: THeaderSection): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function HeaderSection_ClassName(AObj: THeaderSection): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function HeaderSection_InstanceSize(AObj: THeaderSection): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function HeaderSection_InheritsFrom(AObj: THeaderSection; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function HeaderSection_Equals(AObj: THeaderSection; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function HeaderSection_GetHashCode(AObj: THeaderSection): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function HeaderSection_ToString(AObj: THeaderSection): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function HeaderSection_GetLeft(AObj: THeaderSection): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

function HeaderSection_GetRight(AObj: THeaderSection): Integer; stdcall;
begin
  Result :=  AObj.Right;
end;

function HeaderSection_GetAlignment(AObj: THeaderSection): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure HeaderSection_SetAlignment(AObj: THeaderSection; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function HeaderSection_GetAllowClick(AObj: THeaderSection): LongBool; stdcall;
begin
  Result :=  AObj.AllowClick;
end;

procedure HeaderSection_SetAllowClick(AObj: THeaderSection; AValue: LongBool); stdcall;
begin
  AObj.AllowClick := AValue;
end;

function HeaderSection_GetAutoSize(AObj: THeaderSection): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure HeaderSection_SetAutoSize(AObj: THeaderSection; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function HeaderSection_GetBiDiMode(AObj: THeaderSection): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure HeaderSection_SetBiDiMode(AObj: THeaderSection; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function HeaderSection_GetCheckBox(AObj: THeaderSection): LongBool; stdcall;
begin
  Result :=  AObj.CheckBox;
end;

procedure HeaderSection_SetCheckBox(AObj: THeaderSection; AValue: LongBool); stdcall;
begin
  AObj.CheckBox := AValue;
end;

function HeaderSection_GetChecked(AObj: THeaderSection): LongBool; stdcall;
begin
  Result :=  AObj.Checked;
end;

procedure HeaderSection_SetChecked(AObj: THeaderSection; AValue: LongBool); stdcall;
begin
  AObj.Checked := AValue;
end;

function HeaderSection_GetFixedWidth(AObj: THeaderSection): LongBool; stdcall;
begin
  Result :=  AObj.FixedWidth;
end;

procedure HeaderSection_SetFixedWidth(AObj: THeaderSection; AValue: LongBool); stdcall;
begin
  AObj.FixedWidth := AValue;
end;

function HeaderSection_GetImageIndex(AObj: THeaderSection): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure HeaderSection_SetImageIndex(AObj: THeaderSection; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function HeaderSection_GetMaxWidth(AObj: THeaderSection): Integer; stdcall;
begin
  Result :=  AObj.MaxWidth;
end;

procedure HeaderSection_SetMaxWidth(AObj: THeaderSection; AValue: Integer); stdcall;
begin
  AObj.MaxWidth := AValue;
end;

function HeaderSection_GetMinWidth(AObj: THeaderSection): Integer; stdcall;
begin
  Result :=  AObj.MinWidth;
end;

procedure HeaderSection_SetMinWidth(AObj: THeaderSection; AValue: Integer); stdcall;
begin
  AObj.MinWidth := AValue;
end;

function HeaderSection_GetStyle(AObj: THeaderSection): THeaderSectionStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure HeaderSection_SetStyle(AObj: THeaderSection; AValue: THeaderSectionStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function HeaderSection_GetText(AObj: THeaderSection): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure HeaderSection_SetText(AObj: THeaderSection; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function HeaderSection_GetWidth(AObj: THeaderSection): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure HeaderSection_SetWidth(AObj: THeaderSection; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function HeaderSection_GetIndex(AObj: THeaderSection): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure HeaderSection_SetIndex(AObj: THeaderSection; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;

function HeaderSection_StaticClassType: TClass; stdcall;
begin
  Result := THeaderSection;
end;

exports
  HeaderSection_Create,
  HeaderSection_Free,
  HeaderSection_Assign,
  HeaderSection_GetNamePath,
  HeaderSection_DisposeOf,
  HeaderSection_ClassType,
  HeaderSection_ClassName,
  HeaderSection_InstanceSize,
  HeaderSection_InheritsFrom,
  HeaderSection_Equals,
  HeaderSection_GetHashCode,
  HeaderSection_ToString,
  HeaderSection_GetLeft,
  HeaderSection_GetRight,
  HeaderSection_GetAlignment,
  HeaderSection_SetAlignment,
  HeaderSection_GetAllowClick,
  HeaderSection_SetAllowClick,
  HeaderSection_GetAutoSize,
  HeaderSection_SetAutoSize,
  HeaderSection_GetBiDiMode,
  HeaderSection_SetBiDiMode,
  HeaderSection_GetCheckBox,
  HeaderSection_SetCheckBox,
  HeaderSection_GetChecked,
  HeaderSection_SetChecked,
  HeaderSection_GetFixedWidth,
  HeaderSection_SetFixedWidth,
  HeaderSection_GetImageIndex,
  HeaderSection_SetImageIndex,
  HeaderSection_GetMaxWidth,
  HeaderSection_SetMaxWidth,
  HeaderSection_GetMinWidth,
  HeaderSection_SetMinWidth,
  HeaderSection_GetStyle,
  HeaderSection_SetStyle,
  HeaderSection_GetText,
  HeaderSection_SetText,
  HeaderSection_GetWidth,
  HeaderSection_SetWidth,
  HeaderSection_GetIndex,
  HeaderSection_SetIndex,
  HeaderSection_StaticClassType;

