
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ToolBar_Create(AOwner: TComponent): TToolBar; extdecl;
begin
  Result :=  TToolBar.Create(AOwner);
end;

procedure ToolBar_Free(AObj: TToolBar); extdecl;
begin
  AObj.Free;
end;

procedure ToolBar_FlipChildren(AObj: TToolBar; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function ToolBar_CanFocus(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

function ToolBar_Focused(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure ToolBar_Invalidate(AObj: TToolBar); extdecl;
begin
  AObj.Invalidate;
end;

procedure ToolBar_Realign(AObj: TToolBar); extdecl;
begin
  AObj.Realign;
end;

procedure ToolBar_Repaint(AObj: TToolBar); extdecl;
begin
  AObj.Repaint;
end;

procedure ToolBar_ScaleBy(AObj: TToolBar; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure ToolBar_SetBounds(AObj: TToolBar; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ToolBar_SetFocus(AObj: TToolBar); extdecl;
begin
  AObj.SetFocus;
end;

procedure ToolBar_Update(AObj: TToolBar); extdecl;
begin
  AObj.Update;
end;

procedure ToolBar_BringToFront(AObj: TToolBar); extdecl;
begin
  AObj.BringToFront;
end;

function ToolBar_HasParent(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ToolBar_Hide(AObj: TToolBar); extdecl;
begin
  AObj.Hide;
end;

function ToolBar_Perform(AObj: TToolBar; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ToolBar_Refresh(AObj: TToolBar); extdecl;
begin
  AObj.Refresh;
end;

procedure ToolBar_SendToBack(AObj: TToolBar); extdecl;
begin
  AObj.SendToBack;
end;

procedure ToolBar_Show(AObj: TToolBar); extdecl;
begin
  AObj.Show;
end;

function ToolBar_GetTextBuf(AObj: TToolBar; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ToolBar_GetTextLen(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function ToolBar_FindComponent(AObj: TToolBar; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ToolBar_GetNamePath(AObj: TToolBar): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ToolBar_Assign(AObj: TToolBar; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ToolBar_ClassName(AObj: TToolBar): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ToolBar_Equals(AObj: TToolBar; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ToolBar_GetHashCode(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ToolBar_ToString(AObj: TToolBar): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ToolBar_GetButtonCount(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.ButtonCount;
end;

function ToolBar_GetCanvas(AObj: TToolBar): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function ToolBar_GetRowCount(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.RowCount;
end;

function ToolBar_GetAlign(AObj: TToolBar): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure ToolBar_SetAlign(AObj: TToolBar; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function ToolBar_GetAnchors(AObj: TToolBar): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure ToolBar_SetAnchors(AObj: TToolBar; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function ToolBar_GetAutoSize(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.AutoSize;
end;

procedure ToolBar_SetAutoSize(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.AutoSize := AValue;
end;

function ToolBar_GetBorderWidth(AObj: TToolBar): TBorderWidth; extdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure ToolBar_SetBorderWidth(AObj: TToolBar; AValue: TBorderWidth); extdecl;
begin
  AObj.BorderWidth := AValue;
end;

function ToolBar_GetButtonHeight(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.ButtonHeight;
end;

procedure ToolBar_SetButtonHeight(AObj: TToolBar; AValue: Integer); extdecl;
begin
  AObj.ButtonHeight := AValue;
end;

function ToolBar_GetButtonWidth(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.ButtonWidth;
end;

procedure ToolBar_SetButtonWidth(AObj: TToolBar; AValue: Integer); extdecl;
begin
  AObj.ButtonWidth := AValue;
end;

function ToolBar_GetCaption(AObj: TToolBar): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure ToolBar_SetCaption(AObj: TToolBar; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function ToolBar_GetColor(AObj: TToolBar): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure ToolBar_SetColor(AObj: TToolBar; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function ToolBar_GetDoubleBuffered(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ToolBar_SetDoubleBuffered(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function ToolBar_GetEnabled(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure ToolBar_SetEnabled(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function ToolBar_GetFlat(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.Flat;
end;

procedure ToolBar_SetFlat(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.Flat := AValue;
end;

function ToolBar_GetFont(AObj: TToolBar): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure ToolBar_SetFont(AObj: TToolBar; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function ToolBar_GetHeight(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ToolBar_SetHeight(AObj: TToolBar; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ToolBar_GetHotImages(AObj: TToolBar): TCustomImageList; extdecl;
begin
  Result :=  AObj.HotImages;
end;

procedure ToolBar_SetHotImages(AObj: TToolBar; AValue: TCustomImageList); extdecl;
begin
  AObj.HotImages := AValue;
end;

function ToolBar_GetImages(AObj: TToolBar): TCustomImageList; extdecl;
begin
  Result :=  AObj.Images;
end;

procedure ToolBar_SetImages(AObj: TToolBar; AValue: TCustomImageList); extdecl;
begin
  AObj.Images := AValue;
end;

function ToolBar_GetIndent(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.Indent;
end;

procedure ToolBar_SetIndent(AObj: TToolBar; AValue: Integer); extdecl;
begin
  AObj.Indent := AValue;
end;

function ToolBar_GetList(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.List;
end;

procedure ToolBar_SetList(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.List := AValue;
end;

function ToolBar_GetParentColor(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure ToolBar_SetParentColor(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function ToolBar_GetParentFont(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure ToolBar_SetParentFont(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function ToolBar_GetParentShowHint(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ToolBar_SetParentShowHint(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function ToolBar_GetPopupMenu(AObj: TToolBar): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ToolBar_SetPopupMenu(AObj: TToolBar; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function ToolBar_GetShowCaptions(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.ShowCaptions;
end;

procedure ToolBar_SetShowCaptions(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.ShowCaptions := AValue;
end;

function ToolBar_GetShowHint(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure ToolBar_SetShowHint(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function ToolBar_GetTabOrder(AObj: TToolBar): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure ToolBar_SetTabOrder(AObj: TToolBar; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function ToolBar_GetTabStop(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure ToolBar_SetTabStop(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function ToolBar_GetTransparent(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.Transparent;
end;

procedure ToolBar_SetTransparent(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.Transparent := AValue;
end;

function ToolBar_GetVisible(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure ToolBar_SetVisible(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function ToolBar_GetWrapable(AObj: TToolBar): LongBool; extdecl;
begin
  Result :=  AObj.Wrapable;
end;

procedure ToolBar_SetWrapable(AObj: TToolBar; AValue: LongBool); extdecl;
begin
  AObj.Wrapable := AValue;
end;

procedure ToolBar_SetOnClick(AObj: TToolBar; AEventId: NativeUInt); extdecl;
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

procedure ToolBar_SetOnDblClick(AObj: TToolBar; AEventId: NativeUInt); extdecl;
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

procedure ToolBar_SetOnEnter(AObj: TToolBar; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnEnter := nil;
    EventClass.Remove(AObj, geEnter);
    Exit;
  end;
  AObj.OnEnter := @EventClass.OnEnter;
  EventClass.Add(AObj, geEnter, AEventId);
end;

procedure ToolBar_SetOnExit(AObj: TToolBar; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnExit := nil;
    EventClass.Remove(AObj, geExit);
    Exit;
  end;
  AObj.OnExit := @EventClass.OnExit;
  EventClass.Add(AObj, geExit, AEventId);
end;

procedure ToolBar_SetOnMouseDown(AObj: TToolBar; AEventId: NativeUInt); extdecl;
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

procedure ToolBar_SetOnMouseEnter(AObj: TToolBar; AEventId: NativeUInt); extdecl;
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

procedure ToolBar_SetOnMouseLeave(AObj: TToolBar; AEventId: NativeUInt); extdecl;
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

procedure ToolBar_SetOnMouseMove(AObj: TToolBar; AEventId: NativeUInt); extdecl;
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

procedure ToolBar_SetOnMouseUp(AObj: TToolBar; AEventId: NativeUInt); extdecl;
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

procedure ToolBar_SetOnResize(AObj: TToolBar; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnResize := nil;
    EventClass.Remove(AObj, geResize);
    Exit;
  end;
  AObj.OnResize := @EventClass.OnResize;
  EventClass.Add(AObj, geResize, AEventId);
end;

function ToolBar_GetBrush(AObj: TToolBar): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function ToolBar_GetControlCount(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function ToolBar_GetHandle(AObj: TToolBar): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function ToolBar_GetParentWindow(AObj: TToolBar): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ToolBar_SetParentWindow(AObj: TToolBar; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function ToolBar_GetAction(AObj: TToolBar): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure ToolBar_SetAction(AObj: TToolBar; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function ToolBar_GetBiDiMode(AObj: TToolBar): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ToolBar_SetBiDiMode(AObj: TToolBar; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure ToolBar_GetBoundsRect(AObj: TToolBar; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ToolBar_SetBoundsRect(AObj: TToolBar; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function ToolBar_GetClientHeight(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ToolBar_SetClientHeight(AObj: TToolBar; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure ToolBar_GetClientRect(AObj: TToolBar; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function ToolBar_GetClientWidth(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ToolBar_SetClientWidth(AObj: TToolBar; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function ToolBar_GetParent(AObj: TToolBar): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure ToolBar_SetParent(AObj: TToolBar; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function ToolBar_GetLeft(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ToolBar_SetLeft(AObj: TToolBar; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ToolBar_GetTop(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ToolBar_SetTop(AObj: TToolBar; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ToolBar_GetWidth(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ToolBar_SetWidth(AObj: TToolBar; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function ToolBar_GetCursor(AObj: TToolBar): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure ToolBar_SetCursor(AObj: TToolBar; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function ToolBar_GetHint(AObj: TToolBar): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ToolBar_SetHint(AObj: TToolBar; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function ToolBar_GetComponentCount(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ToolBar_GetComponentIndex(AObj: TToolBar): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ToolBar_SetComponentIndex(AObj: TToolBar; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ToolBar_GetOwner(AObj: TToolBar): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ToolBar_GetName(AObj: TToolBar): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ToolBar_SetName(AObj: TToolBar; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ToolBar_GetTag(AObj: TToolBar): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ToolBar_SetTag(AObj: TToolBar; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ToolBar_GetButtons(AObj: TToolBar; Index: Integer): TToolButton; extdecl;
begin
  Result :=  AObj.Buttons[Index];
end;

function ToolBar_GetControls(AObj: TToolBar; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function ToolBar_GetComponents(AObj: TToolBar; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ToolBar_Create,
  ToolBar_Free,
  ToolBar_FlipChildren,
  ToolBar_CanFocus,
  ToolBar_Focused,
  ToolBar_Invalidate,
  ToolBar_Realign,
  ToolBar_Repaint,
  ToolBar_ScaleBy,
  ToolBar_SetBounds,
  ToolBar_SetFocus,
  ToolBar_Update,
  ToolBar_BringToFront,
  ToolBar_HasParent,
  ToolBar_Hide,
  ToolBar_Perform,
  ToolBar_Refresh,
  ToolBar_SendToBack,
  ToolBar_Show,
  ToolBar_GetTextBuf,
  ToolBar_GetTextLen,
  ToolBar_FindComponent,
  ToolBar_GetNamePath,
  ToolBar_Assign,
  ToolBar_ClassName,
  ToolBar_Equals,
  ToolBar_GetHashCode,
  ToolBar_ToString,
  ToolBar_GetButtonCount,
  ToolBar_GetCanvas,
  ToolBar_GetRowCount,
  ToolBar_GetAlign,
  ToolBar_SetAlign,
  ToolBar_GetAnchors,
  ToolBar_SetAnchors,
  ToolBar_GetAutoSize,
  ToolBar_SetAutoSize,
  ToolBar_GetBorderWidth,
  ToolBar_SetBorderWidth,
  ToolBar_GetButtonHeight,
  ToolBar_SetButtonHeight,
  ToolBar_GetButtonWidth,
  ToolBar_SetButtonWidth,
  ToolBar_GetCaption,
  ToolBar_SetCaption,
  ToolBar_GetColor,
  ToolBar_SetColor,
  ToolBar_GetDoubleBuffered,
  ToolBar_SetDoubleBuffered,
  ToolBar_GetEnabled,
  ToolBar_SetEnabled,
  ToolBar_GetFlat,
  ToolBar_SetFlat,
  ToolBar_GetFont,
  ToolBar_SetFont,
  ToolBar_GetHeight,
  ToolBar_SetHeight,
  ToolBar_GetHotImages,
  ToolBar_SetHotImages,
  ToolBar_GetImages,
  ToolBar_SetImages,
  ToolBar_GetIndent,
  ToolBar_SetIndent,
  ToolBar_GetList,
  ToolBar_SetList,
  ToolBar_GetParentColor,
  ToolBar_SetParentColor,
  ToolBar_GetParentFont,
  ToolBar_SetParentFont,
  ToolBar_GetParentShowHint,
  ToolBar_SetParentShowHint,
  ToolBar_GetPopupMenu,
  ToolBar_SetPopupMenu,
  ToolBar_GetShowCaptions,
  ToolBar_SetShowCaptions,
  ToolBar_GetShowHint,
  ToolBar_SetShowHint,
  ToolBar_GetTabOrder,
  ToolBar_SetTabOrder,
  ToolBar_GetTabStop,
  ToolBar_SetTabStop,
  ToolBar_GetTransparent,
  ToolBar_SetTransparent,
  ToolBar_GetVisible,
  ToolBar_SetVisible,
  ToolBar_GetWrapable,
  ToolBar_SetWrapable,
  ToolBar_SetOnClick,
  ToolBar_SetOnDblClick,
  ToolBar_SetOnEnter,
  ToolBar_SetOnExit,
  ToolBar_SetOnMouseDown,
  ToolBar_SetOnMouseEnter,
  ToolBar_SetOnMouseLeave,
  ToolBar_SetOnMouseMove,
  ToolBar_SetOnMouseUp,
  ToolBar_SetOnResize,
  ToolBar_GetBrush,
  ToolBar_GetControlCount,
  ToolBar_GetHandle,
  ToolBar_GetParentWindow,
  ToolBar_SetParentWindow,
  ToolBar_GetAction,
  ToolBar_SetAction,
  ToolBar_GetBiDiMode,
  ToolBar_SetBiDiMode,
  ToolBar_GetBoundsRect,
  ToolBar_SetBoundsRect,
  ToolBar_GetClientHeight,
  ToolBar_SetClientHeight,
  ToolBar_GetClientRect,
  ToolBar_GetClientWidth,
  ToolBar_SetClientWidth,
  ToolBar_GetParent,
  ToolBar_SetParent,
  ToolBar_GetLeft,
  ToolBar_SetLeft,
  ToolBar_GetTop,
  ToolBar_SetTop,
  ToolBar_GetWidth,
  ToolBar_SetWidth,
  ToolBar_GetCursor,
  ToolBar_SetCursor,
  ToolBar_GetHint,
  ToolBar_SetHint,
  ToolBar_GetComponentCount,
  ToolBar_GetComponentIndex,
  ToolBar_SetComponentIndex,
  ToolBar_GetOwner,
  ToolBar_GetName,
  ToolBar_SetName,
  ToolBar_GetTag,
  ToolBar_SetTag,
  ToolBar_GetButtons,
  ToolBar_GetControls,
  ToolBar_GetComponents;

