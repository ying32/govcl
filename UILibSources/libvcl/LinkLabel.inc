
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function LinkLabel_Create(AOwner: TComponent): TLinkLabel; stdcall;
begin
  Result :=  TLinkLabel.Create(AOwner);
end;

procedure LinkLabel_Free(AObj: TLinkLabel); stdcall;
begin
  AObj.Free;
end;

function LinkLabel_CanFocus(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure LinkLabel_FlipChildren(AObj: TLinkLabel; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function LinkLabel_Focused(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function LinkLabel_HandleAllocated(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure LinkLabel_Invalidate(AObj: TLinkLabel); stdcall;
begin
  AObj.Invalidate;
end;

procedure LinkLabel_Realign(AObj: TLinkLabel); stdcall;
begin
  AObj.Realign;
end;

procedure LinkLabel_Repaint(AObj: TLinkLabel); stdcall;
begin
  AObj.Repaint;
end;

procedure LinkLabel_ScaleBy(AObj: TLinkLabel; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure LinkLabel_SetBounds(AObj: TLinkLabel; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure LinkLabel_SetFocus(AObj: TLinkLabel); stdcall;
begin
  AObj.SetFocus;
end;

procedure LinkLabel_Update(AObj: TLinkLabel); stdcall;
begin
  AObj.Update;
end;

procedure LinkLabel_BringToFront(AObj: TLinkLabel); stdcall;
begin
  AObj.BringToFront;
end;

function LinkLabel_HasParent(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure LinkLabel_Hide(AObj: TLinkLabel); stdcall;
begin
  AObj.Hide;
end;

function LinkLabel_Perform(AObj: TLinkLabel; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure LinkLabel_Refresh(AObj: TLinkLabel); stdcall;
begin
  AObj.Refresh;
end;

procedure LinkLabel_SendToBack(AObj: TLinkLabel); stdcall;
begin
  AObj.SendToBack;
end;

procedure LinkLabel_Show(AObj: TLinkLabel); stdcall;
begin
  AObj.Show;
end;

function LinkLabel_GetTextBuf(AObj: TLinkLabel; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function LinkLabel_GetTextLen(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function LinkLabel_FindComponent(AObj: TLinkLabel; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function LinkLabel_GetNamePath(AObj: TLinkLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure LinkLabel_Assign(AObj: TLinkLabel; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function LinkLabel_ClassName(AObj: TLinkLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function LinkLabel_Equals(AObj: TLinkLabel; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function LinkLabel_GetHashCode(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function LinkLabel_ToString(AObj: TLinkLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function LinkLabel_GetAlign(AObj: TLinkLabel): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure LinkLabel_SetAlign(AObj: TLinkLabel; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function LinkLabel_GetAlignment(AObj: TLinkLabel): TCustomLinkLabel.TLinkAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure LinkLabel_SetAlignment(AObj: TLinkLabel; AValue: TCustomLinkLabel.TLinkAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function LinkLabel_GetAnchors(AObj: TLinkLabel): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure LinkLabel_SetAnchors(AObj: TLinkLabel; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function LinkLabel_GetAutoSize(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure LinkLabel_SetAutoSize(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function LinkLabel_GetBevelEdges(AObj: TLinkLabel): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure LinkLabel_SetBevelEdges(AObj: TLinkLabel; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function LinkLabel_GetBevelInner(AObj: TLinkLabel): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure LinkLabel_SetBevelInner(AObj: TLinkLabel; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function LinkLabel_GetBevelKind(AObj: TLinkLabel): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure LinkLabel_SetBevelKind(AObj: TLinkLabel; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function LinkLabel_GetBevelOuter(AObj: TLinkLabel): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure LinkLabel_SetBevelOuter(AObj: TLinkLabel; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function LinkLabel_GetCaption(AObj: TLinkLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure LinkLabel_SetCaption(AObj: TLinkLabel; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function LinkLabel_GetColor(AObj: TLinkLabel): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure LinkLabel_SetColor(AObj: TLinkLabel; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function LinkLabel_GetEnabled(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure LinkLabel_SetEnabled(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function LinkLabel_GetFont(AObj: TLinkLabel): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure LinkLabel_SetFont(AObj: TLinkLabel; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function LinkLabel_GetParentColor(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure LinkLabel_SetParentColor(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function LinkLabel_GetParentFont(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure LinkLabel_SetParentFont(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function LinkLabel_GetParentShowHint(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure LinkLabel_SetParentShowHint(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function LinkLabel_GetPopupMenu(AObj: TLinkLabel): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure LinkLabel_SetPopupMenu(AObj: TLinkLabel; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function LinkLabel_GetShowHint(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure LinkLabel_SetShowHint(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function LinkLabel_GetTabOrder(AObj: TLinkLabel): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure LinkLabel_SetTabOrder(AObj: TLinkLabel; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function LinkLabel_GetTabStop(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure LinkLabel_SetTabStop(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function LinkLabel_GetUseVisualStyle(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.UseVisualStyle;
end;

procedure LinkLabel_SetUseVisualStyle(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.UseVisualStyle := AValue;
end;

function LinkLabel_GetVisible(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure LinkLabel_SetVisible(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

procedure LinkLabel_SetOnClick(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnDblClick(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseDown(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseEnter(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseLeave(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseMove(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseUp(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnLinkClick(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnLinkClick := nil;
    TEventClass.Remove(AObj, geLinkClick);
    Exit;
  end;
  AObj.OnLinkClick := TEventClass.OnLinkClick;
  TEventClass.Add(AObj, geLinkClick, AEventId);
end;

function LinkLabel_GetDoubleBuffered(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure LinkLabel_SetDoubleBuffered(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function LinkLabel_GetBrush(AObj: TLinkLabel): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function LinkLabel_GetControlCount(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function LinkLabel_GetHandle(AObj: TLinkLabel): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function LinkLabel_GetParentDoubleBuffered(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure LinkLabel_SetParentDoubleBuffered(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function LinkLabel_GetParentWindow(AObj: TLinkLabel): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure LinkLabel_SetParentWindow(AObj: TLinkLabel; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function LinkLabel_GetAction(AObj: TLinkLabel): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure LinkLabel_SetAction(AObj: TLinkLabel; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function LinkLabel_GetBiDiMode(AObj: TLinkLabel): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure LinkLabel_SetBiDiMode(AObj: TLinkLabel; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure LinkLabel_GetBoundsRect(AObj: TLinkLabel; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure LinkLabel_SetBoundsRect(AObj: TLinkLabel; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function LinkLabel_GetClientHeight(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure LinkLabel_SetClientHeight(AObj: TLinkLabel; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure LinkLabel_GetClientRect(AObj: TLinkLabel; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function LinkLabel_GetClientWidth(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure LinkLabel_SetClientWidth(AObj: TLinkLabel; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function LinkLabel_GetExplicitLeft(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function LinkLabel_GetExplicitTop(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function LinkLabel_GetExplicitWidth(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function LinkLabel_GetExplicitHeight(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function LinkLabel_GetParent(AObj: TLinkLabel): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure LinkLabel_SetParent(AObj: TLinkLabel; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function LinkLabel_GetStyleElements(AObj: TLinkLabel): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure LinkLabel_SetStyleElements(AObj: TLinkLabel; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function LinkLabel_GetAlignWithMargins(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure LinkLabel_SetAlignWithMargins(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function LinkLabel_GetLeft(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure LinkLabel_SetLeft(AObj: TLinkLabel; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function LinkLabel_GetTop(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure LinkLabel_SetTop(AObj: TLinkLabel; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function LinkLabel_GetWidth(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure LinkLabel_SetWidth(AObj: TLinkLabel; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function LinkLabel_GetHeight(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure LinkLabel_SetHeight(AObj: TLinkLabel; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function LinkLabel_GetCursor(AObj: TLinkLabel): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure LinkLabel_SetCursor(AObj: TLinkLabel; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function LinkLabel_GetHint(AObj: TLinkLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure LinkLabel_SetHint(AObj: TLinkLabel; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function LinkLabel_GetMargins(AObj: TLinkLabel): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure LinkLabel_SetMargins(AObj: TLinkLabel; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function LinkLabel_GetCustomHint(AObj: TLinkLabel): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure LinkLabel_SetCustomHint(AObj: TLinkLabel; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function LinkLabel_GetComponentCount(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function LinkLabel_GetComponentIndex(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure LinkLabel_SetComponentIndex(AObj: TLinkLabel; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function LinkLabel_GetOwner(AObj: TLinkLabel): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function LinkLabel_GetName(AObj: TLinkLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure LinkLabel_SetName(AObj: TLinkLabel; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function LinkLabel_GetTag(AObj: TLinkLabel): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure LinkLabel_SetTag(AObj: TLinkLabel; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function LinkLabel_GetControls(AObj: TLinkLabel; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function LinkLabel_GetComponents(AObj: TLinkLabel; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  LinkLabel_Create,
  LinkLabel_Free,
  LinkLabel_CanFocus,
  LinkLabel_FlipChildren,
  LinkLabel_Focused,
  LinkLabel_HandleAllocated,
  LinkLabel_Invalidate,
  LinkLabel_Realign,
  LinkLabel_Repaint,
  LinkLabel_ScaleBy,
  LinkLabel_SetBounds,
  LinkLabel_SetFocus,
  LinkLabel_Update,
  LinkLabel_BringToFront,
  LinkLabel_HasParent,
  LinkLabel_Hide,
  LinkLabel_Perform,
  LinkLabel_Refresh,
  LinkLabel_SendToBack,
  LinkLabel_Show,
  LinkLabel_GetTextBuf,
  LinkLabel_GetTextLen,
  LinkLabel_FindComponent,
  LinkLabel_GetNamePath,
  LinkLabel_Assign,
  LinkLabel_ClassName,
  LinkLabel_Equals,
  LinkLabel_GetHashCode,
  LinkLabel_ToString,
  LinkLabel_GetAlign,
  LinkLabel_SetAlign,
  LinkLabel_GetAlignment,
  LinkLabel_SetAlignment,
  LinkLabel_GetAnchors,
  LinkLabel_SetAnchors,
  LinkLabel_GetAutoSize,
  LinkLabel_SetAutoSize,
  LinkLabel_GetBevelEdges,
  LinkLabel_SetBevelEdges,
  LinkLabel_GetBevelInner,
  LinkLabel_SetBevelInner,
  LinkLabel_GetBevelKind,
  LinkLabel_SetBevelKind,
  LinkLabel_GetBevelOuter,
  LinkLabel_SetBevelOuter,
  LinkLabel_GetCaption,
  LinkLabel_SetCaption,
  LinkLabel_GetColor,
  LinkLabel_SetColor,
  LinkLabel_GetEnabled,
  LinkLabel_SetEnabled,
  LinkLabel_GetFont,
  LinkLabel_SetFont,
  LinkLabel_GetParentColor,
  LinkLabel_SetParentColor,
  LinkLabel_GetParentFont,
  LinkLabel_SetParentFont,
  LinkLabel_GetParentShowHint,
  LinkLabel_SetParentShowHint,
  LinkLabel_GetPopupMenu,
  LinkLabel_SetPopupMenu,
  LinkLabel_GetShowHint,
  LinkLabel_SetShowHint,
  LinkLabel_GetTabOrder,
  LinkLabel_SetTabOrder,
  LinkLabel_GetTabStop,
  LinkLabel_SetTabStop,
  LinkLabel_GetUseVisualStyle,
  LinkLabel_SetUseVisualStyle,
  LinkLabel_GetVisible,
  LinkLabel_SetVisible,
  LinkLabel_SetOnClick,
  LinkLabel_SetOnDblClick,
  LinkLabel_SetOnMouseDown,
  LinkLabel_SetOnMouseEnter,
  LinkLabel_SetOnMouseLeave,
  LinkLabel_SetOnMouseMove,
  LinkLabel_SetOnMouseUp,
  LinkLabel_SetOnLinkClick,
  LinkLabel_GetDoubleBuffered,
  LinkLabel_SetDoubleBuffered,
  LinkLabel_GetBrush,
  LinkLabel_GetControlCount,
  LinkLabel_GetHandle,
  LinkLabel_GetParentDoubleBuffered,
  LinkLabel_SetParentDoubleBuffered,
  LinkLabel_GetParentWindow,
  LinkLabel_SetParentWindow,
  LinkLabel_GetAction,
  LinkLabel_SetAction,
  LinkLabel_GetBiDiMode,
  LinkLabel_SetBiDiMode,
  LinkLabel_GetBoundsRect,
  LinkLabel_SetBoundsRect,
  LinkLabel_GetClientHeight,
  LinkLabel_SetClientHeight,
  LinkLabel_GetClientRect,
  LinkLabel_GetClientWidth,
  LinkLabel_SetClientWidth,
  LinkLabel_GetExplicitLeft,
  LinkLabel_GetExplicitTop,
  LinkLabel_GetExplicitWidth,
  LinkLabel_GetExplicitHeight,
  LinkLabel_GetParent,
  LinkLabel_SetParent,
  LinkLabel_GetStyleElements,
  LinkLabel_SetStyleElements,
  LinkLabel_GetAlignWithMargins,
  LinkLabel_SetAlignWithMargins,
  LinkLabel_GetLeft,
  LinkLabel_SetLeft,
  LinkLabel_GetTop,
  LinkLabel_SetTop,
  LinkLabel_GetWidth,
  LinkLabel_SetWidth,
  LinkLabel_GetHeight,
  LinkLabel_SetHeight,
  LinkLabel_GetCursor,
  LinkLabel_SetCursor,
  LinkLabel_GetHint,
  LinkLabel_SetHint,
  LinkLabel_GetMargins,
  LinkLabel_SetMargins,
  LinkLabel_GetCustomHint,
  LinkLabel_SetCustomHint,
  LinkLabel_GetComponentCount,
  LinkLabel_GetComponentIndex,
  LinkLabel_SetComponentIndex,
  LinkLabel_GetOwner,
  LinkLabel_GetName,
  LinkLabel_SetName,
  LinkLabel_GetTag,
  LinkLabel_SetTag,
  LinkLabel_GetControls,
  LinkLabel_GetComponents;

