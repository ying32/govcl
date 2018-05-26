
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function SpeedButton_Create(AOwner: TComponent): TSpeedButton; extdecl;
begin
  Result :=  TSpeedButton.Create(AOwner);
end;

procedure SpeedButton_Free(AObj: TSpeedButton); extdecl;
begin
  AObj.Free;
end;

procedure SpeedButton_Click(AObj: TSpeedButton); extdecl;
begin
  AObj.Click;
end;

procedure SpeedButton_BringToFront(AObj: TSpeedButton); extdecl;
begin
  AObj.BringToFront;
end;

function SpeedButton_HasParent(AObj: TSpeedButton): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure SpeedButton_Hide(AObj: TSpeedButton); extdecl;
begin
  AObj.Hide;
end;

procedure SpeedButton_Invalidate(AObj: TSpeedButton); extdecl;
begin
  AObj.Invalidate;
end;

function SpeedButton_Perform(AObj: TSpeedButton; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure SpeedButton_Refresh(AObj: TSpeedButton); extdecl;
begin
  AObj.Refresh;
end;

procedure SpeedButton_Repaint(AObj: TSpeedButton); extdecl;
begin
  AObj.Repaint;
end;

procedure SpeedButton_SendToBack(AObj: TSpeedButton); extdecl;
begin
  AObj.SendToBack;
end;

procedure SpeedButton_SetBounds(AObj: TSpeedButton; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure SpeedButton_Show(AObj: TSpeedButton); extdecl;
begin
  AObj.Show;
end;

procedure SpeedButton_Update(AObj: TSpeedButton); extdecl;
begin
  AObj.Update;
end;

function SpeedButton_GetTextBuf(AObj: TSpeedButton; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function SpeedButton_GetTextLen(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function SpeedButton_FindComponent(AObj: TSpeedButton; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function SpeedButton_GetNamePath(AObj: TSpeedButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure SpeedButton_Assign(AObj: TSpeedButton; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function SpeedButton_ClassName(AObj: TSpeedButton): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function SpeedButton_Equals(AObj: TSpeedButton; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function SpeedButton_GetHashCode(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function SpeedButton_ToString(AObj: TSpeedButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function SpeedButton_GetAction(AObj: TSpeedButton): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure SpeedButton_SetAction(AObj: TSpeedButton; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function SpeedButton_GetAlign(AObj: TSpeedButton): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure SpeedButton_SetAlign(AObj: TSpeedButton; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function SpeedButton_GetAllowAllUp(AObj: TSpeedButton): LongBool; extdecl;
begin
  Result :=  AObj.AllowAllUp;
end;

procedure SpeedButton_SetAllowAllUp(AObj: TSpeedButton; AValue: LongBool); extdecl;
begin
  AObj.AllowAllUp := AValue;
end;

function SpeedButton_GetAnchors(AObj: TSpeedButton): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure SpeedButton_SetAnchors(AObj: TSpeedButton; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function SpeedButton_GetBiDiMode(AObj: TSpeedButton): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure SpeedButton_SetBiDiMode(AObj: TSpeedButton; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function SpeedButton_GetGroupIndex(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.GroupIndex;
end;

procedure SpeedButton_SetGroupIndex(AObj: TSpeedButton; AValue: Integer); extdecl;
begin
  AObj.GroupIndex := AValue;
end;

function SpeedButton_GetDown(AObj: TSpeedButton): LongBool; extdecl;
begin
  Result :=  AObj.Down;
end;

procedure SpeedButton_SetDown(AObj: TSpeedButton; AValue: LongBool); extdecl;
begin
  AObj.Down := AValue;
end;

function SpeedButton_GetCaption(AObj: TSpeedButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure SpeedButton_SetCaption(AObj: TSpeedButton; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function SpeedButton_GetEnabled(AObj: TSpeedButton): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure SpeedButton_SetEnabled(AObj: TSpeedButton; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function SpeedButton_GetFlat(AObj: TSpeedButton): LongBool; extdecl;
begin
  Result :=  AObj.Flat;
end;

procedure SpeedButton_SetFlat(AObj: TSpeedButton; AValue: LongBool); extdecl;
begin
  AObj.Flat := AValue;
end;

function SpeedButton_GetFont(AObj: TSpeedButton): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure SpeedButton_SetFont(AObj: TSpeedButton; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function SpeedButton_GetGlyph(AObj: TSpeedButton): TBitmap; extdecl;
begin
  Result :=  AObj.Glyph;
end;

procedure SpeedButton_SetGlyph(AObj: TSpeedButton; AValue: TBitmap); extdecl;
begin
  AObj.Glyph := AValue;
end;

function SpeedButton_GetLayout(AObj: TSpeedButton): TButtonLayout; extdecl;
begin
  Result :=  AObj.Layout;
end;

procedure SpeedButton_SetLayout(AObj: TSpeedButton; AValue: TButtonLayout); extdecl;
begin
  AObj.Layout := AValue;
end;

function SpeedButton_GetNumGlyphs(AObj: TSpeedButton): TNumGlyphs; extdecl;
begin
  Result :=  AObj.NumGlyphs;
end;

procedure SpeedButton_SetNumGlyphs(AObj: TSpeedButton; AValue: TNumGlyphs); extdecl;
begin
  AObj.NumGlyphs := AValue;
end;

function SpeedButton_GetParentFont(AObj: TSpeedButton): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure SpeedButton_SetParentFont(AObj: TSpeedButton; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function SpeedButton_GetParentShowHint(AObj: TSpeedButton): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure SpeedButton_SetParentShowHint(AObj: TSpeedButton; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function SpeedButton_GetPopupMenu(AObj: TSpeedButton): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure SpeedButton_SetPopupMenu(AObj: TSpeedButton; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function SpeedButton_GetShowHint(AObj: TSpeedButton): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure SpeedButton_SetShowHint(AObj: TSpeedButton; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function SpeedButton_GetSpacing(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.Spacing;
end;

procedure SpeedButton_SetSpacing(AObj: TSpeedButton; AValue: Integer); extdecl;
begin
  AObj.Spacing := AValue;
end;

function SpeedButton_GetTransparent(AObj: TSpeedButton): LongBool; extdecl;
begin
  Result :=  AObj.Transparent;
end;

procedure SpeedButton_SetTransparent(AObj: TSpeedButton; AValue: LongBool); extdecl;
begin
  AObj.Transparent := AValue;
end;

function SpeedButton_GetVisible(AObj: TSpeedButton): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure SpeedButton_SetVisible(AObj: TSpeedButton; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure SpeedButton_SetOnClick(AObj: TSpeedButton; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    EventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := @EventClass.OnClick;
  EventClass.Add(AObj, geClick, AEventId);
end;

procedure SpeedButton_SetOnDblClick(AObj: TSpeedButton; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnDblClick := nil;
    EventClass.Remove(AObj, geDblClick);
    Exit;
  end;
  AObj.OnDblClick := @EventClass.OnDblClick;
  EventClass.Add(AObj, geDblClick, AEventId);
end;

procedure SpeedButton_SetOnMouseDown(AObj: TSpeedButton; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseDown := nil;
    EventClass.Remove(AObj, geMouseDown);
    Exit;
  end;
  AObj.OnMouseDown := @EventClass.OnMouseDown;
  EventClass.Add(AObj, geMouseDown, AEventId);
end;

procedure SpeedButton_SetOnMouseEnter(AObj: TSpeedButton; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseEnter := nil;
    EventClass.Remove(AObj, geMouseEnter);
    Exit;
  end;
  AObj.OnMouseEnter := @EventClass.OnMouseEnter;
  EventClass.Add(AObj, geMouseEnter, AEventId);
end;

procedure SpeedButton_SetOnMouseLeave(AObj: TSpeedButton; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseLeave := nil;
    EventClass.Remove(AObj, geMouseLeave);
    Exit;
  end;
  AObj.OnMouseLeave := @EventClass.OnMouseLeave;
  EventClass.Add(AObj, geMouseLeave, AEventId);
end;

procedure SpeedButton_SetOnMouseMove(AObj: TSpeedButton; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseMove := nil;
    EventClass.Remove(AObj, geMouseMove);
    Exit;
  end;
  AObj.OnMouseMove := @EventClass.OnMouseMove;
  EventClass.Add(AObj, geMouseMove, AEventId);
end;

procedure SpeedButton_SetOnMouseUp(AObj: TSpeedButton; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseUp := nil;
    EventClass.Remove(AObj, geMouseUp);
    Exit;
  end;
  AObj.OnMouseUp := @EventClass.OnMouseUp;
  EventClass.Add(AObj, geMouseUp, AEventId);
end;

procedure SpeedButton_GetBoundsRect(AObj: TSpeedButton; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure SpeedButton_SetBoundsRect(AObj: TSpeedButton; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function SpeedButton_GetClientHeight(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure SpeedButton_SetClientHeight(AObj: TSpeedButton; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure SpeedButton_GetClientRect(AObj: TSpeedButton; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function SpeedButton_GetClientWidth(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure SpeedButton_SetClientWidth(AObj: TSpeedButton; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function SpeedButton_GetParent(AObj: TSpeedButton): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure SpeedButton_SetParent(AObj: TSpeedButton; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function SpeedButton_GetLeft(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure SpeedButton_SetLeft(AObj: TSpeedButton; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function SpeedButton_GetTop(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure SpeedButton_SetTop(AObj: TSpeedButton; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function SpeedButton_GetWidth(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure SpeedButton_SetWidth(AObj: TSpeedButton; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function SpeedButton_GetHeight(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure SpeedButton_SetHeight(AObj: TSpeedButton; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function SpeedButton_GetCursor(AObj: TSpeedButton): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure SpeedButton_SetCursor(AObj: TSpeedButton; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function SpeedButton_GetHint(AObj: TSpeedButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure SpeedButton_SetHint(AObj: TSpeedButton; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function SpeedButton_GetComponentCount(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function SpeedButton_GetComponentIndex(AObj: TSpeedButton): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure SpeedButton_SetComponentIndex(AObj: TSpeedButton; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function SpeedButton_GetOwner(AObj: TSpeedButton): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function SpeedButton_GetName(AObj: TSpeedButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure SpeedButton_SetName(AObj: TSpeedButton; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function SpeedButton_GetTag(AObj: TSpeedButton): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure SpeedButton_SetTag(AObj: TSpeedButton; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function SpeedButton_GetComponents(AObj: TSpeedButton; AIndex: Integer): TComponent; extdecl;
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
  SpeedButton_GetParent,
  SpeedButton_SetParent,
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
  SpeedButton_GetComponentCount,
  SpeedButton_GetComponentIndex,
  SpeedButton_SetComponentIndex,
  SpeedButton_GetOwner,
  SpeedButton_GetName,
  SpeedButton_SetName,
  SpeedButton_GetTag,
  SpeedButton_SetTag,
  SpeedButton_GetComponents;

