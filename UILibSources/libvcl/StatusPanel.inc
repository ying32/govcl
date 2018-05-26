
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function StatusPanel_Create(AOwner: TCollection): TStatusPanel; stdcall;
begin
  Result :=  TStatusPanel.Create(AOwner);
end;

procedure StatusPanel_Free(AObj: TStatusPanel); stdcall;
begin
  AObj.Free;
end;

procedure StatusPanel_Assign(AObj: TStatusPanel; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function StatusPanel_GetNamePath(AObj: TStatusPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function StatusPanel_ClassName(AObj: TStatusPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function StatusPanel_Equals(AObj: TStatusPanel; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function StatusPanel_GetHashCode(AObj: TStatusPanel): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function StatusPanel_ToString(AObj: TStatusPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function StatusPanel_GetAlignment(AObj: TStatusPanel): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure StatusPanel_SetAlignment(AObj: TStatusPanel; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function StatusPanel_GetBiDiMode(AObj: TStatusPanel): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure StatusPanel_SetBiDiMode(AObj: TStatusPanel; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function StatusPanel_GetStyle(AObj: TStatusPanel): TStatusPanelStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure StatusPanel_SetStyle(AObj: TStatusPanel; AValue: TStatusPanelStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function StatusPanel_GetText(AObj: TStatusPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure StatusPanel_SetText(AObj: TStatusPanel; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function StatusPanel_GetWidth(AObj: TStatusPanel): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure StatusPanel_SetWidth(AObj: TStatusPanel; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function StatusPanel_GetIndex(AObj: TStatusPanel): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

procedure StatusPanel_SetIndex(AObj: TStatusPanel; AValue: Integer); stdcall;
begin
  AObj.Index := AValue;
end;


exports
  StatusPanel_Create,
  StatusPanel_Free,
  StatusPanel_Assign,
  StatusPanel_GetNamePath,
  StatusPanel_ClassName,
  StatusPanel_Equals,
  StatusPanel_GetHashCode,
  StatusPanel_ToString,
  StatusPanel_GetAlignment,
  StatusPanel_SetAlignment,
  StatusPanel_GetBiDiMode,
  StatusPanel_SetBiDiMode,
  StatusPanel_GetStyle,
  StatusPanel_SetStyle,
  StatusPanel_GetText,
  StatusPanel_SetText,
  StatusPanel_GetWidth,
  StatusPanel_SetWidth,
  StatusPanel_GetIndex,
  StatusPanel_SetIndex;

