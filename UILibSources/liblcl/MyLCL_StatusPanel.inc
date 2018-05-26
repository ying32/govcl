
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function StatusPanel_Create(AOwner: TCollection): TStatusPanel; extdecl;
begin
  Result :=  TStatusPanel.Create(AOwner);
end;

procedure StatusPanel_Free(AObj: TStatusPanel); extdecl;
begin
  AObj.Free;
end;

procedure StatusPanel_Assign(AObj: TStatusPanel; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function StatusPanel_GetNamePath(AObj: TStatusPanel): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function StatusPanel_ClassName(AObj: TStatusPanel): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function StatusPanel_Equals(AObj: TStatusPanel; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function StatusPanel_GetHashCode(AObj: TStatusPanel): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function StatusPanel_ToString(AObj: TStatusPanel): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function StatusPanel_GetAlignment(AObj: TStatusPanel): TAlignment; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure StatusPanel_SetAlignment(AObj: TStatusPanel; AValue: TAlignment); extdecl;
begin
  AObj.Alignment := AValue;
end;

function StatusPanel_GetBiDiMode(AObj: TStatusPanel): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure StatusPanel_SetBiDiMode(AObj: TStatusPanel; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function StatusPanel_GetStyle(AObj: TStatusPanel): TStatusPanelStyle; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure StatusPanel_SetStyle(AObj: TStatusPanel; AValue: TStatusPanelStyle); extdecl;
begin
  AObj.Style := AValue;
end;

function StatusPanel_GetText(AObj: TStatusPanel): PChar; extdecl;
begin
  Result :=  PChar(AObj.Text);
end;

procedure StatusPanel_SetText(AObj: TStatusPanel; AValue: PChar); extdecl;
begin
  AObj.Text := AValue;
end;

function StatusPanel_GetWidth(AObj: TStatusPanel): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure StatusPanel_SetWidth(AObj: TStatusPanel; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function StatusPanel_GetIndex(AObj: TStatusPanel): Integer; extdecl;
begin
  Result :=  AObj.Index;
end;

procedure StatusPanel_SetIndex(AObj: TStatusPanel; AValue: Integer); extdecl;
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

