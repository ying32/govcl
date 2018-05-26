
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Gauge_Create(AOwner: TComponent): TGauge; stdcall;
begin
  Result :=  TGauge.Create(AOwner);
end;

procedure Gauge_Free(AObj: TGauge); stdcall;
begin
  AObj.Free;
end;

procedure Gauge_AddProgress(AObj: TGauge; Value: Integer); stdcall;
begin
  AObj.AddProgress(Value);
end;

procedure Gauge_BringToFront(AObj: TGauge); stdcall;
begin
  AObj.BringToFront;
end;

function Gauge_HasParent(AObj: TGauge): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Gauge_Hide(AObj: TGauge); stdcall;
begin
  AObj.Hide;
end;

procedure Gauge_Invalidate(AObj: TGauge); stdcall;
begin
  AObj.Invalidate;
end;

function Gauge_Perform(AObj: TGauge; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Gauge_Refresh(AObj: TGauge); stdcall;
begin
  AObj.Refresh;
end;

procedure Gauge_Repaint(AObj: TGauge); stdcall;
begin
  AObj.Repaint;
end;

procedure Gauge_SendToBack(AObj: TGauge); stdcall;
begin
  AObj.SendToBack;
end;

procedure Gauge_SetBounds(AObj: TGauge; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Gauge_Show(AObj: TGauge); stdcall;
begin
  AObj.Show;
end;

procedure Gauge_Update(AObj: TGauge); stdcall;
begin
  AObj.Update;
end;

function Gauge_GetTextBuf(AObj: TGauge; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Gauge_GetTextLen(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function Gauge_FindComponent(AObj: TGauge; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Gauge_GetNamePath(AObj: TGauge): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Gauge_Assign(AObj: TGauge; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Gauge_ClassName(AObj: TGauge): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Gauge_Equals(AObj: TGauge; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Gauge_GetHashCode(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Gauge_ToString(AObj: TGauge): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Gauge_GetPercentDone(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.PercentDone;
end;

function Gauge_GetAlign(AObj: TGauge): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Gauge_SetAlign(AObj: TGauge; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Gauge_GetAnchors(AObj: TGauge): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Gauge_SetAnchors(AObj: TGauge; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Gauge_GetBackColor(AObj: TGauge): TColor; stdcall;
begin
  Result :=  AObj.BackColor;
end;

procedure Gauge_SetBackColor(AObj: TGauge; AValue: TColor); stdcall;
begin
  AObj.BackColor := AValue;
end;

function Gauge_GetBorderStyle(AObj: TGauge): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure Gauge_SetBorderStyle(AObj: TGauge; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function Gauge_GetColor(AObj: TGauge): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure Gauge_SetColor(AObj: TGauge; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function Gauge_GetEnabled(AObj: TGauge): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Gauge_SetEnabled(AObj: TGauge; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Gauge_GetForeColor(AObj: TGauge): TColor; stdcall;
begin
  Result :=  AObj.ForeColor;
end;

procedure Gauge_SetForeColor(AObj: TGauge; AValue: TColor); stdcall;
begin
  AObj.ForeColor := AValue;
end;

function Gauge_GetFont(AObj: TGauge): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure Gauge_SetFont(AObj: TGauge; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function Gauge_GetKind(AObj: TGauge): TGaugeKind; stdcall;
begin
  Result :=  AObj.Kind;
end;

procedure Gauge_SetKind(AObj: TGauge; AValue: TGaugeKind); stdcall;
begin
  AObj.Kind := AValue;
end;

function Gauge_GetMinValue(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.MinValue;
end;

procedure Gauge_SetMinValue(AObj: TGauge; AValue: Integer); stdcall;
begin
  AObj.MinValue := AValue;
end;

function Gauge_GetMaxValue(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.MaxValue;
end;

procedure Gauge_SetMaxValue(AObj: TGauge; AValue: Integer); stdcall;
begin
  AObj.MaxValue := AValue;
end;

function Gauge_GetParentColor(AObj: TGauge): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure Gauge_SetParentColor(AObj: TGauge; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function Gauge_GetParentFont(AObj: TGauge): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure Gauge_SetParentFont(AObj: TGauge; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function Gauge_GetParentShowHint(AObj: TGauge): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Gauge_SetParentShowHint(AObj: TGauge; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function Gauge_GetPopupMenu(AObj: TGauge): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Gauge_SetPopupMenu(AObj: TGauge; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function Gauge_GetProgress(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.Progress;
end;

procedure Gauge_SetProgress(AObj: TGauge; AValue: Integer); stdcall;
begin
  AObj.Progress := AValue;
end;

function Gauge_GetShowHint(AObj: TGauge): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Gauge_SetShowHint(AObj: TGauge; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Gauge_GetShowText(AObj: TGauge): LongBool; stdcall;
begin
  Result :=  AObj.ShowText;
end;

procedure Gauge_SetShowText(AObj: TGauge; AValue: LongBool); stdcall;
begin
  AObj.ShowText := AValue;
end;

function Gauge_GetVisible(AObj: TGauge): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Gauge_SetVisible(AObj: TGauge; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Gauge_GetAction(AObj: TGauge): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Gauge_SetAction(AObj: TGauge; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function Gauge_GetBiDiMode(AObj: TGauge): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Gauge_SetBiDiMode(AObj: TGauge; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure Gauge_GetBoundsRect(AObj: TGauge; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Gauge_SetBoundsRect(AObj: TGauge; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Gauge_GetClientHeight(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Gauge_SetClientHeight(AObj: TGauge; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Gauge_GetClientRect(AObj: TGauge; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Gauge_GetClientWidth(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Gauge_SetClientWidth(AObj: TGauge; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Gauge_GetExplicitLeft(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Gauge_GetExplicitTop(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Gauge_GetExplicitWidth(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Gauge_GetExplicitHeight(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Gauge_GetParent(AObj: TGauge): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Gauge_SetParent(AObj: TGauge; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Gauge_GetStyleElements(AObj: TGauge): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Gauge_SetStyleElements(AObj: TGauge; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function Gauge_GetAlignWithMargins(AObj: TGauge): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Gauge_SetAlignWithMargins(AObj: TGauge; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Gauge_GetLeft(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Gauge_SetLeft(AObj: TGauge; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Gauge_GetTop(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Gauge_SetTop(AObj: TGauge; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Gauge_GetWidth(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Gauge_SetWidth(AObj: TGauge; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Gauge_GetHeight(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Gauge_SetHeight(AObj: TGauge; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Gauge_GetCursor(AObj: TGauge): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Gauge_SetCursor(AObj: TGauge; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Gauge_GetHint(AObj: TGauge): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Gauge_SetHint(AObj: TGauge; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Gauge_GetMargins(AObj: TGauge): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Gauge_SetMargins(AObj: TGauge; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Gauge_GetCustomHint(AObj: TGauge): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Gauge_SetCustomHint(AObj: TGauge; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Gauge_GetComponentCount(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Gauge_GetComponentIndex(AObj: TGauge): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Gauge_SetComponentIndex(AObj: TGauge; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Gauge_GetOwner(AObj: TGauge): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Gauge_GetName(AObj: TGauge): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Gauge_SetName(AObj: TGauge; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Gauge_GetTag(AObj: TGauge): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Gauge_SetTag(AObj: TGauge; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Gauge_GetComponents(AObj: TGauge; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Gauge_Create,
  Gauge_Free,
  Gauge_AddProgress,
  Gauge_BringToFront,
  Gauge_HasParent,
  Gauge_Hide,
  Gauge_Invalidate,
  Gauge_Perform,
  Gauge_Refresh,
  Gauge_Repaint,
  Gauge_SendToBack,
  Gauge_SetBounds,
  Gauge_Show,
  Gauge_Update,
  Gauge_GetTextBuf,
  Gauge_GetTextLen,
  Gauge_FindComponent,
  Gauge_GetNamePath,
  Gauge_Assign,
  Gauge_ClassName,
  Gauge_Equals,
  Gauge_GetHashCode,
  Gauge_ToString,
  Gauge_GetPercentDone,
  Gauge_GetAlign,
  Gauge_SetAlign,
  Gauge_GetAnchors,
  Gauge_SetAnchors,
  Gauge_GetBackColor,
  Gauge_SetBackColor,
  Gauge_GetBorderStyle,
  Gauge_SetBorderStyle,
  Gauge_GetColor,
  Gauge_SetColor,
  Gauge_GetEnabled,
  Gauge_SetEnabled,
  Gauge_GetForeColor,
  Gauge_SetForeColor,
  Gauge_GetFont,
  Gauge_SetFont,
  Gauge_GetKind,
  Gauge_SetKind,
  Gauge_GetMinValue,
  Gauge_SetMinValue,
  Gauge_GetMaxValue,
  Gauge_SetMaxValue,
  Gauge_GetParentColor,
  Gauge_SetParentColor,
  Gauge_GetParentFont,
  Gauge_SetParentFont,
  Gauge_GetParentShowHint,
  Gauge_SetParentShowHint,
  Gauge_GetPopupMenu,
  Gauge_SetPopupMenu,
  Gauge_GetProgress,
  Gauge_SetProgress,
  Gauge_GetShowHint,
  Gauge_SetShowHint,
  Gauge_GetShowText,
  Gauge_SetShowText,
  Gauge_GetVisible,
  Gauge_SetVisible,
  Gauge_GetAction,
  Gauge_SetAction,
  Gauge_GetBiDiMode,
  Gauge_SetBiDiMode,
  Gauge_GetBoundsRect,
  Gauge_SetBoundsRect,
  Gauge_GetClientHeight,
  Gauge_SetClientHeight,
  Gauge_GetClientRect,
  Gauge_GetClientWidth,
  Gauge_SetClientWidth,
  Gauge_GetExplicitLeft,
  Gauge_GetExplicitTop,
  Gauge_GetExplicitWidth,
  Gauge_GetExplicitHeight,
  Gauge_GetParent,
  Gauge_SetParent,
  Gauge_GetStyleElements,
  Gauge_SetStyleElements,
  Gauge_GetAlignWithMargins,
  Gauge_SetAlignWithMargins,
  Gauge_GetLeft,
  Gauge_SetLeft,
  Gauge_GetTop,
  Gauge_SetTop,
  Gauge_GetWidth,
  Gauge_SetWidth,
  Gauge_GetHeight,
  Gauge_SetHeight,
  Gauge_GetCursor,
  Gauge_SetCursor,
  Gauge_GetHint,
  Gauge_SetHint,
  Gauge_GetMargins,
  Gauge_SetMargins,
  Gauge_GetCustomHint,
  Gauge_SetCustomHint,
  Gauge_GetComponentCount,
  Gauge_GetComponentIndex,
  Gauge_SetComponentIndex,
  Gauge_GetOwner,
  Gauge_GetName,
  Gauge_SetName,
  Gauge_GetTag,
  Gauge_SetTag,
  Gauge_GetComponents;

