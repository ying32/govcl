
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function SpeedButton_Create(AOwner: TComponent): TSpeedButton; stdcall;
begin
  Result :=  TSpeedButton.Create(AOwner);
end;

procedure SpeedButton_Free(AObj: TSpeedButton); stdcall;
begin
  AObj.Free;
end;

procedure SpeedButton_Click(AObj: TSpeedButton); stdcall;
begin
  AObj.Click;
end;

procedure SpeedButton_BringToFront(AObj: TSpeedButton); stdcall;
begin
  AObj.BringToFront;
end;

function SpeedButton_HasParent(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure SpeedButton_Hide(AObj: TSpeedButton); stdcall;
begin
  AObj.Hide;
end;

procedure SpeedButton_Invalidate(AObj: TSpeedButton); stdcall;
begin
  AObj.Invalidate;
end;

function SpeedButton_Perform(AObj: TSpeedButton; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure SpeedButton_Refresh(AObj: TSpeedButton); stdcall;
begin
  AObj.Refresh;
end;

procedure SpeedButton_Repaint(AObj: TSpeedButton); stdcall;
begin
  AObj.Repaint;
end;

procedure SpeedButton_SendToBack(AObj: TSpeedButton); stdcall;
begin
  AObj.SendToBack;
end;

procedure SpeedButton_SetBounds(AObj: TSpeedButton; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure SpeedButton_Show(AObj: TSpeedButton); stdcall;
begin
  AObj.Show;
end;

procedure SpeedButton_Update(AObj: TSpeedButton); stdcall;
begin
  AObj.Update;
end;

function SpeedButton_GetTextBuf(AObj: TSpeedButton; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function SpeedButton_GetTextLen(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function SpeedButton_FindComponent(AObj: TSpeedButton; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function SpeedButton_GetNamePath(AObj: TSpeedButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure SpeedButton_Assign(AObj: TSpeedButton; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function SpeedButton_ClassName(AObj: TSpeedButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function SpeedButton_Equals(AObj: TSpeedButton; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function SpeedButton_GetHashCode(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function SpeedButton_ToString(AObj: TSpeedButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function SpeedButton_GetAction(AObj: TSpeedButton): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure SpeedButton_SetAction(AObj: TSpeedButton; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function SpeedButton_GetAlign(AObj: TSpeedButton): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure SpeedButton_SetAlign(AObj: TSpeedButton; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function SpeedButton_GetAllowAllUp(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.AllowAllUp;
end;

procedure SpeedButton_SetAllowAllUp(AObj: TSpeedButton; AValue: LongBool); stdcall;
begin
  AObj.AllowAllUp := AValue;
end;

function SpeedButton_GetAnchors(AObj: TSpeedButton): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure SpeedButton_SetAnchors(AObj: TSpeedButton; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function SpeedButton_GetBiDiMode(AObj: TSpeedButton): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure SpeedButton_SetBiDiMode(AObj: TSpeedButton; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function SpeedButton_GetGroupIndex(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.GroupIndex;
end;

procedure SpeedButton_SetGroupIndex(AObj: TSpeedButton; AValue: Integer); stdcall;
begin
  AObj.GroupIndex := AValue;
end;

function SpeedButton_GetDown(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.Down;
end;

procedure SpeedButton_SetDown(AObj: TSpeedButton; AValue: LongBool); stdcall;
begin
  AObj.Down := AValue;
end;

function SpeedButton_GetCaption(AObj: TSpeedButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure SpeedButton_SetCaption(AObj: TSpeedButton; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function SpeedButton_GetEnabled(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure SpeedButton_SetEnabled(AObj: TSpeedButton; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function SpeedButton_GetFlat(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.Flat;
end;

procedure SpeedButton_SetFlat(AObj: TSpeedButton; AValue: LongBool); stdcall;
begin
  AObj.Flat := AValue;
end;

function SpeedButton_GetFont(AObj: TSpeedButton): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure SpeedButton_SetFont(AObj: TSpeedButton; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function SpeedButton_GetGlyph(AObj: TSpeedButton): TBitmap; stdcall;
begin
  Result :=  AObj.Glyph;
end;

procedure SpeedButton_SetGlyph(AObj: TSpeedButton; AValue: TBitmap); stdcall;
begin
  AObj.Glyph := AValue;
end;

function SpeedButton_GetLayout(AObj: TSpeedButton): TButtonLayout; stdcall;
begin
  Result :=  AObj.Layout;
end;

procedure SpeedButton_SetLayout(AObj: TSpeedButton; AValue: TButtonLayout); stdcall;
begin
  AObj.Layout := AValue;
end;

function SpeedButton_GetNumGlyphs(AObj: TSpeedButton): TNumGlyphs; stdcall;
begin
  Result :=  AObj.NumGlyphs;
end;

procedure SpeedButton_SetNumGlyphs(AObj: TSpeedButton; AValue: TNumGlyphs); stdcall;
begin
  AObj.NumGlyphs := AValue;
end;

function SpeedButton_GetParentFont(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure SpeedButton_SetParentFont(AObj: TSpeedButton; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function SpeedButton_GetParentShowHint(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure SpeedButton_SetParentShowHint(AObj: TSpeedButton; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function SpeedButton_GetPopupMenu(AObj: TSpeedButton): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure SpeedButton_SetPopupMenu(AObj: TSpeedButton; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function SpeedButton_GetShowHint(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure SpeedButton_SetShowHint(AObj: TSpeedButton; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function SpeedButton_GetSpacing(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.Spacing;
end;

procedure SpeedButton_SetSpacing(AObj: TSpeedButton; AValue: Integer); stdcall;
begin
  AObj.Spacing := AValue;
end;

function SpeedButton_GetTransparent(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure SpeedButton_SetTransparent(AObj: TSpeedButton; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function SpeedButton_GetVisible(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure SpeedButton_SetVisible(AObj: TSpeedButton; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function SpeedButton_GetStyleElements(AObj: TSpeedButton): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure SpeedButton_SetStyleElements(AObj: TSpeedButton; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure SpeedButton_SetOnClick(AObj: TSpeedButton; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    TEventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, geClick, AEventId);
end;

procedure SpeedButton_SetOnDblClick(AObj: TSpeedButton; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDblClick := nil;
    TEventClass.Remove(AObj, geDblClick);
    Exit;
  end;
  AObj.OnDblClick := TEventClass.OnDblClick;
  TEventClass.Add(AObj, geDblClick, AEventId);
end;

procedure SpeedButton_SetOnMouseDown(AObj: TSpeedButton; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseDown := nil;
    TEventClass.Remove(AObj, geMouseDown);
    Exit;
  end;
  AObj.OnMouseDown := TEventClass.OnMouseDown;
  TEventClass.Add(AObj, geMouseDown, AEventId);
end;

procedure SpeedButton_SetOnMouseEnter(AObj: TSpeedButton; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseEnter := nil;
    TEventClass.Remove(AObj, geMouseEnter);
    Exit;
  end;
  AObj.OnMouseEnter := TEventClass.OnMouseEnter;
  TEventClass.Add(AObj, geMouseEnter, AEventId);
end;

procedure SpeedButton_SetOnMouseLeave(AObj: TSpeedButton; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseLeave := nil;
    TEventClass.Remove(AObj, geMouseLeave);
    Exit;
  end;
  AObj.OnMouseLeave := TEventClass.OnMouseLeave;
  TEventClass.Add(AObj, geMouseLeave, AEventId);
end;

procedure SpeedButton_SetOnMouseMove(AObj: TSpeedButton; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseMove := nil;
    TEventClass.Remove(AObj, geMouseMove);
    Exit;
  end;
  AObj.OnMouseMove := TEventClass.OnMouseMove;
  TEventClass.Add(AObj, geMouseMove, AEventId);
end;

procedure SpeedButton_SetOnMouseUp(AObj: TSpeedButton; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseUp := nil;
    TEventClass.Remove(AObj, geMouseUp);
    Exit;
  end;
  AObj.OnMouseUp := TEventClass.OnMouseUp;
  TEventClass.Add(AObj, geMouseUp, AEventId);
end;

procedure SpeedButton_GetBoundsRect(AObj: TSpeedButton; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure SpeedButton_SetBoundsRect(AObj: TSpeedButton; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function SpeedButton_GetClientHeight(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure SpeedButton_SetClientHeight(AObj: TSpeedButton; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure SpeedButton_GetClientRect(AObj: TSpeedButton; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function SpeedButton_GetClientWidth(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure SpeedButton_SetClientWidth(AObj: TSpeedButton; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function SpeedButton_GetExplicitLeft(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function SpeedButton_GetExplicitTop(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function SpeedButton_GetExplicitWidth(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function SpeedButton_GetExplicitHeight(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function SpeedButton_GetParent(AObj: TSpeedButton): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure SpeedButton_SetParent(AObj: TSpeedButton; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function SpeedButton_GetAlignWithMargins(AObj: TSpeedButton): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure SpeedButton_SetAlignWithMargins(AObj: TSpeedButton; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function SpeedButton_GetLeft(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure SpeedButton_SetLeft(AObj: TSpeedButton; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function SpeedButton_GetTop(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure SpeedButton_SetTop(AObj: TSpeedButton; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function SpeedButton_GetWidth(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure SpeedButton_SetWidth(AObj: TSpeedButton; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function SpeedButton_GetHeight(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure SpeedButton_SetHeight(AObj: TSpeedButton; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function SpeedButton_GetCursor(AObj: TSpeedButton): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure SpeedButton_SetCursor(AObj: TSpeedButton; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function SpeedButton_GetHint(AObj: TSpeedButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure SpeedButton_SetHint(AObj: TSpeedButton; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function SpeedButton_GetMargins(AObj: TSpeedButton): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure SpeedButton_SetMargins(AObj: TSpeedButton; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function SpeedButton_GetCustomHint(AObj: TSpeedButton): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure SpeedButton_SetCustomHint(AObj: TSpeedButton; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function SpeedButton_GetComponentCount(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function SpeedButton_GetComponentIndex(AObj: TSpeedButton): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure SpeedButton_SetComponentIndex(AObj: TSpeedButton; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function SpeedButton_GetOwner(AObj: TSpeedButton): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function SpeedButton_GetName(AObj: TSpeedButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure SpeedButton_SetName(AObj: TSpeedButton; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function SpeedButton_GetTag(AObj: TSpeedButton): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure SpeedButton_SetTag(AObj: TSpeedButton; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function SpeedButton_GetComponents(AObj: TSpeedButton; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  SpeedButton_Create,
  SpeedButton_Free,
  SpeedButton_Click,
  SpeedButton_BringToFront,
  SpeedButton_HasParent,
  SpeedButton_Hide,
  SpeedButton_Invalidate,
  SpeedButton_Perform,
  SpeedButton_Refresh,
  SpeedButton_Repaint,
  SpeedButton_SendToBack,
  SpeedButton_SetBounds,
  SpeedButton_Show,
  SpeedButton_Update,
  SpeedButton_GetTextBuf,
  SpeedButton_GetTextLen,
  SpeedButton_FindComponent,
  SpeedButton_GetNamePath,
  SpeedButton_Assign,
  SpeedButton_ClassName,
  SpeedButton_Equals,
  SpeedButton_GetHashCode,
  SpeedButton_ToString,
  SpeedButton_GetAction,
  SpeedButton_SetAction,
  SpeedButton_GetAlign,
  SpeedButton_SetAlign,
  SpeedButton_GetAllowAllUp,
  SpeedButton_SetAllowAllUp,
  SpeedButton_GetAnchors,
  SpeedButton_SetAnchors,
  SpeedButton_GetBiDiMode,
  SpeedButton_SetBiDiMode,
  SpeedButton_GetGroupIndex,
  SpeedButton_SetGroupIndex,
  SpeedButton_GetDown,
  SpeedButton_SetDown,
  SpeedButton_GetCaption,
  SpeedButton_SetCaption,
  SpeedButton_GetEnabled,
  SpeedButton_SetEnabled,
  SpeedButton_GetFlat,
  SpeedButton_SetFlat,
  SpeedButton_GetFont,
  SpeedButton_SetFont,
  SpeedButton_GetGlyph,
  SpeedButton_SetGlyph,
  SpeedButton_GetLayout,
  SpeedButton_SetLayout,
  SpeedButton_GetNumGlyphs,
  SpeedButton_SetNumGlyphs,
  SpeedButton_GetParentFont,
  SpeedButton_SetParentFont,
  SpeedButton_GetParentShowHint,
  SpeedButton_SetParentShowHint,
  SpeedButton_GetPopupMenu,
  SpeedButton_SetPopupMenu,
  SpeedButton_GetShowHint,
  SpeedButton_SetShowHint,
  SpeedButton_GetSpacing,
  SpeedButton_SetSpacing,
  SpeedButton_GetTransparent,
  SpeedButton_SetTransparent,
  SpeedButton_GetVisible,
  SpeedButton_SetVisible,
  SpeedButton_GetStyleElements,
  SpeedButton_SetStyleElements,
  SpeedButton_SetOnClick,
  SpeedButton_SetOnDblClick,
  SpeedButton_SetOnMouseDown,
  SpeedButton_SetOnMouseEnter,
  SpeedButton_SetOnMouseLeave,
  SpeedButton_SetOnMouseMove,
  SpeedButton_SetOnMouseUp,
  SpeedButton_GetBoundsRect,
  SpeedButton_SetBoundsRect,
  SpeedButton_GetClientHeight,
  SpeedButton_SetClientHeight,
  SpeedButton_GetClientRect,
  SpeedButton_GetClientWidth,
  SpeedButton_SetClientWidth,
  SpeedButton_GetExplicitLeft,
  SpeedButton_GetExplicitTop,
  SpeedButton_GetExplicitWidth,
  SpeedButton_GetExplicitHeight,
  SpeedButton_GetParent,
  SpeedButton_SetParent,
  SpeedButton_GetAlignWithMargins,
  SpeedButton_SetAlignWithMargins,
  SpeedButton_GetLeft,
  SpeedButton_SetLeft,
  SpeedButton_GetTop,
  SpeedButton_SetTop,
  SpeedButton_GetWidth,
  SpeedButton_SetWidth,
  SpeedButton_GetHeight,
  SpeedButton_SetHeight,
  SpeedButton_GetCursor,
  SpeedButton_SetCursor,
  SpeedButton_GetHint,
  SpeedButton_SetHint,
  SpeedButton_GetMargins,
  SpeedButton_SetMargins,
  SpeedButton_GetCustomHint,
  SpeedButton_SetCustomHint,
  SpeedButton_GetComponentCount,
  SpeedButton_GetComponentIndex,
  SpeedButton_SetComponentIndex,
  SpeedButton_GetOwner,
  SpeedButton_GetName,
  SpeedButton_SetName,
  SpeedButton_GetTag,
  SpeedButton_SetTag,
  SpeedButton_GetComponents;

