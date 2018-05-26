
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function StaticText_Create(AOwner: TComponent): TStaticText; extdecl;
begin
  Result :=  TStaticText.Create(AOwner);
end;

procedure StaticText_Free(AObj: TStaticText); extdecl;
begin
  AObj.Free;
end;

function StaticText_CanFocus(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure StaticText_FlipChildren(AObj: TStaticText; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function StaticText_Focused(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure StaticText_Invalidate(AObj: TStaticText); extdecl;
begin
  AObj.Invalidate;
end;

procedure StaticText_Realign(AObj: TStaticText); extdecl;
begin
  AObj.Realign;
end;

procedure StaticText_Repaint(AObj: TStaticText); extdecl;
begin
  AObj.Repaint;
end;

procedure StaticText_ScaleBy(AObj: TStaticText; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure StaticText_SetBounds(AObj: TStaticText; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure StaticText_SetFocus(AObj: TStaticText); extdecl;
begin
  AObj.SetFocus;
end;

procedure StaticText_Update(AObj: TStaticText); extdecl;
begin
  AObj.Update;
end;

procedure StaticText_BringToFront(AObj: TStaticText); extdecl;
begin
  AObj.BringToFront;
end;

function StaticText_HasParent(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure StaticText_Hide(AObj: TStaticText); extdecl;
begin
  AObj.Hide;
end;

function StaticText_Perform(AObj: TStaticText; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure StaticText_Refresh(AObj: TStaticText); extdecl;
begin
  AObj.Refresh;
end;

procedure StaticText_SendToBack(AObj: TStaticText); extdecl;
begin
  AObj.SendToBack;
end;

procedure StaticText_Show(AObj: TStaticText); extdecl;
begin
  AObj.Show;
end;

function StaticText_GetTextBuf(AObj: TStaticText; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function StaticText_GetTextLen(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function StaticText_FindComponent(AObj: TStaticText; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function StaticText_GetNamePath(AObj: TStaticText): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure StaticText_Assign(AObj: TStaticText; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function StaticText_ClassName(AObj: TStaticText): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function StaticText_Equals(AObj: TStaticText; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function StaticText_GetHashCode(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function StaticText_ToString(AObj: TStaticText): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function StaticText_GetAlign(AObj: TStaticText): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure StaticText_SetAlign(AObj: TStaticText; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function StaticText_GetAlignment(AObj: TStaticText): TAlignment; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure StaticText_SetAlignment(AObj: TStaticText; AValue: TAlignment); extdecl;
begin
  AObj.Alignment := AValue;
end;

function StaticText_GetAnchors(AObj: TStaticText): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure StaticText_SetAnchors(AObj: TStaticText; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function StaticText_GetAutoSize(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.AutoSize;
end;

procedure StaticText_SetAutoSize(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.AutoSize := AValue;
end;

function StaticText_GetBiDiMode(AObj: TStaticText): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure StaticText_SetBiDiMode(AObj: TStaticText; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function StaticText_GetBorderStyle(AObj: TStaticText): TStaticBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure StaticText_SetBorderStyle(AObj: TStaticText; AValue: TStaticBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function StaticText_GetCaption(AObj: TStaticText): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure StaticText_SetCaption(AObj: TStaticText; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function StaticText_GetColor(AObj: TStaticText): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure StaticText_SetColor(AObj: TStaticText; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function StaticText_GetDoubleBuffered(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure StaticText_SetDoubleBuffered(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function StaticText_GetEnabled(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure StaticText_SetEnabled(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function StaticText_GetFont(AObj: TStaticText): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure StaticText_SetFont(AObj: TStaticText; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function StaticText_GetParentColor(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure StaticText_SetParentColor(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function StaticText_GetParentFont(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure StaticText_SetParentFont(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function StaticText_GetParentShowHint(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure StaticText_SetParentShowHint(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function StaticText_GetPopupMenu(AObj: TStaticText): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure StaticText_SetPopupMenu(AObj: TStaticText; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function StaticText_GetShowAccelChar(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.ShowAccelChar;
end;

procedure StaticText_SetShowAccelChar(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.ShowAccelChar := AValue;
end;

function StaticText_GetShowHint(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure StaticText_SetShowHint(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function StaticText_GetTabOrder(AObj: TStaticText): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure StaticText_SetTabOrder(AObj: TStaticText; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function StaticText_GetTabStop(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure StaticText_SetTabStop(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function StaticText_GetTransparent(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.Transparent;
end;

procedure StaticText_SetTransparent(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.Transparent := AValue;
end;

function StaticText_GetVisible(AObj: TStaticText): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure StaticText_SetVisible(AObj: TStaticText; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure StaticText_SetOnClick(AObj: TStaticText; AEventId: NativeUInt); extdecl;
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

procedure StaticText_SetOnDblClick(AObj: TStaticText; AEventId: NativeUInt); extdecl;
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

procedure StaticText_SetOnMouseDown(AObj: TStaticText; AEventId: NativeUInt); extdecl;
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

procedure StaticText_SetOnMouseEnter(AObj: TStaticText; AEventId: NativeUInt); extdecl;
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

procedure StaticText_SetOnMouseLeave(AObj: TStaticText; AEventId: NativeUInt); extdecl;
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

procedure StaticText_SetOnMouseMove(AObj: TStaticText; AEventId: NativeUInt); extdecl;
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

procedure StaticText_SetOnMouseUp(AObj: TStaticText; AEventId: NativeUInt); extdecl;
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

function StaticText_GetBrush(AObj: TStaticText): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function StaticText_GetControlCount(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function StaticText_GetHandle(AObj: TStaticText): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function StaticText_GetParentWindow(AObj: TStaticText): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure StaticText_SetParentWindow(AObj: TStaticText; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function StaticText_GetAction(AObj: TStaticText): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure StaticText_SetAction(AObj: TStaticText; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure StaticText_GetBoundsRect(AObj: TStaticText; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure StaticText_SetBoundsRect(AObj: TStaticText; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function StaticText_GetClientHeight(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure StaticText_SetClientHeight(AObj: TStaticText; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure StaticText_GetClientRect(AObj: TStaticText; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function StaticText_GetClientWidth(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure StaticText_SetClientWidth(AObj: TStaticText; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function StaticText_GetParent(AObj: TStaticText): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure StaticText_SetParent(AObj: TStaticText; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function StaticText_GetLeft(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure StaticText_SetLeft(AObj: TStaticText; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function StaticText_GetTop(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure StaticText_SetTop(AObj: TStaticText; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function StaticText_GetWidth(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure StaticText_SetWidth(AObj: TStaticText; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function StaticText_GetHeight(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure StaticText_SetHeight(AObj: TStaticText; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function StaticText_GetCursor(AObj: TStaticText): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure StaticText_SetCursor(AObj: TStaticText; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function StaticText_GetHint(AObj: TStaticText): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure StaticText_SetHint(AObj: TStaticText; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function StaticText_GetComponentCount(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function StaticText_GetComponentIndex(AObj: TStaticText): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure StaticText_SetComponentIndex(AObj: TStaticText; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function StaticText_GetOwner(AObj: TStaticText): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function StaticText_GetName(AObj: TStaticText): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure StaticText_SetName(AObj: TStaticText; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function StaticText_GetTag(AObj: TStaticText): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure StaticText_SetTag(AObj: TStaticText; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function StaticText_GetControls(AObj: TStaticText; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function StaticText_GetComponents(AObj: TStaticText; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  StaticText_Create,
  StaticText_Free,
  StaticText_CanFocus,
  StaticText_FlipChildren,
  StaticText_Focused,
  StaticText_Invalidate,
  StaticText_Realign,
  StaticText_Repaint,
  StaticText_ScaleBy,
  StaticText_SetBounds,
  StaticText_SetFocus,
  StaticText_Update,
  StaticText_BringToFront,
  StaticText_HasParent,
  StaticText_Hide,
  StaticText_Perform,
  StaticText_Refresh,
  StaticText_SendToBack,
  StaticText_Show,
  StaticText_GetTextBuf,
  StaticText_GetTextLen,
  StaticText_FindComponent,
  StaticText_GetNamePath,
  StaticText_Assign,
  StaticText_ClassName,
  StaticText_Equals,
  StaticText_GetHashCode,
  StaticText_ToString,
  StaticText_GetAlign,
  StaticText_SetAlign,
  StaticText_GetAlignment,
  StaticText_SetAlignment,
  StaticText_GetAnchors,
  StaticText_SetAnchors,
  StaticText_GetAutoSize,
  StaticText_SetAutoSize,
  StaticText_GetBiDiMode,
  StaticText_SetBiDiMode,
  StaticText_GetBorderStyle,
  StaticText_SetBorderStyle,
  StaticText_GetCaption,
  StaticText_SetCaption,
  StaticText_GetColor,
  StaticText_SetColor,
  StaticText_GetDoubleBuffered,
  StaticText_SetDoubleBuffered,
  StaticText_GetEnabled,
  StaticText_SetEnabled,
  StaticText_GetFont,
  StaticText_SetFont,
  StaticText_GetParentColor,
  StaticText_SetParentColor,
  StaticText_GetParentFont,
  StaticText_SetParentFont,
  StaticText_GetParentShowHint,
  StaticText_SetParentShowHint,
  StaticText_GetPopupMenu,
  StaticText_SetPopupMenu,
  StaticText_GetShowAccelChar,
  StaticText_SetShowAccelChar,
  StaticText_GetShowHint,
  StaticText_SetShowHint,
  StaticText_GetTabOrder,
  StaticText_SetTabOrder,
  StaticText_GetTabStop,
  StaticText_SetTabStop,
  StaticText_GetTransparent,
  StaticText_SetTransparent,
  StaticText_GetVisible,
  StaticText_SetVisible,
  StaticText_SetOnClick,
  StaticText_SetOnDblClick,
  StaticText_SetOnMouseDown,
  StaticText_SetOnMouseEnter,
  StaticText_SetOnMouseLeave,
  StaticText_SetOnMouseMove,
  StaticText_SetOnMouseUp,
  StaticText_GetBrush,
  StaticText_GetControlCount,
  StaticText_GetHandle,
  StaticText_GetParentWindow,
  StaticText_SetParentWindow,
  StaticText_GetAction,
  StaticText_SetAction,
  StaticText_GetBoundsRect,
  StaticText_SetBoundsRect,
  StaticText_GetClientHeight,
  StaticText_SetClientHeight,
  StaticText_GetClientRect,
  StaticText_GetClientWidth,
  StaticText_SetClientWidth,
  StaticText_GetParent,
  StaticText_SetParent,
  StaticText_GetLeft,
  StaticText_SetLeft,
  StaticText_GetTop,
  StaticText_SetTop,
  StaticText_GetWidth,
  StaticText_SetWidth,
  StaticText_GetHeight,
  StaticText_SetHeight,
  StaticText_GetCursor,
  StaticText_SetCursor,
  StaticText_GetHint,
  StaticText_SetHint,
  StaticText_GetComponentCount,
  StaticText_GetComponentIndex,
  StaticText_SetComponentIndex,
  StaticText_GetOwner,
  StaticText_GetName,
  StaticText_SetName,
  StaticText_GetTag,
  StaticText_SetTag,
  StaticText_GetControls,
  StaticText_GetComponents;

