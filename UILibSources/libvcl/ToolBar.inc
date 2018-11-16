
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ToolBar_Create(AOwner: TComponent): TToolBar; stdcall;
begin
  Result :=  TToolBar.Create(AOwner);
end;

procedure ToolBar_Free(AObj: TToolBar); stdcall;
begin
  AObj.Free;
end;

procedure ToolBar_FlipChildren(AObj: TToolBar; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function ToolBar_CanFocus(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function ToolBar_ContainsControl(AObj: TToolBar; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function ToolBar_ControlAtPos(AObj: TToolBar; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure ToolBar_DisableAlign(AObj: TToolBar); stdcall;
begin
  AObj.DisableAlign;
end;

procedure ToolBar_EnableAlign(AObj: TToolBar); stdcall;
begin
  AObj.EnableAlign;
end;

function ToolBar_FindChildControl(AObj: TToolBar; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

function ToolBar_Focused(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function ToolBar_HandleAllocated(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ToolBar_InsertControl(AObj: TToolBar; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure ToolBar_Invalidate(AObj: TToolBar); stdcall;
begin
  AObj.Invalidate;
end;

procedure ToolBar_PaintTo(AObj: TToolBar; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure ToolBar_RemoveControl(AObj: TToolBar; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure ToolBar_Realign(AObj: TToolBar); stdcall;
begin
  AObj.Realign;
end;

procedure ToolBar_Repaint(AObj: TToolBar); stdcall;
begin
  AObj.Repaint;
end;

procedure ToolBar_ScaleBy(AObj: TToolBar; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure ToolBar_ScrollBy(AObj: TToolBar; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure ToolBar_SetBounds(AObj: TToolBar; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ToolBar_SetFocus(AObj: TToolBar); stdcall;
begin
  AObj.SetFocus;
end;

procedure ToolBar_Update(AObj: TToolBar); stdcall;
begin
  AObj.Update;
end;

procedure ToolBar_UpdateControlState(AObj: TToolBar); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure ToolBar_BringToFront(AObj: TToolBar); stdcall;
begin
  AObj.BringToFront;
end;

procedure ToolBar_ClientToScreen(AObj: TToolBar; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure ToolBar_ClientToParent(AObj: TToolBar; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function ToolBar_Dragging(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function ToolBar_HasParent(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ToolBar_Hide(AObj: TToolBar); stdcall;
begin
  AObj.Hide;
end;

function ToolBar_Perform(AObj: TToolBar; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ToolBar_Refresh(AObj: TToolBar); stdcall;
begin
  AObj.Refresh;
end;

procedure ToolBar_ScreenToClient(AObj: TToolBar; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure ToolBar_ParentToClient(AObj: TToolBar; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure ToolBar_SendToBack(AObj: TToolBar); stdcall;
begin
  AObj.SendToBack;
end;

procedure ToolBar_Show(AObj: TToolBar); stdcall;
begin
  AObj.Show;
end;

function ToolBar_GetTextBuf(AObj: TToolBar; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ToolBar_GetTextLen(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure ToolBar_SetTextBuf(AObj: TToolBar; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function ToolBar_FindComponent(AObj: TToolBar; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ToolBar_GetNamePath(AObj: TToolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ToolBar_Assign(AObj: TToolBar; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ToolBar_DisposeOf(AObj: TToolBar); stdcall;
begin
  AObj.DisposeOf;
end;

function ToolBar_ClassType(AObj: TToolBar): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function ToolBar_ClassName(AObj: TToolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ToolBar_InstanceSize(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function ToolBar_InheritsFrom(AObj: TToolBar; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function ToolBar_Equals(AObj: TToolBar; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ToolBar_GetHashCode(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ToolBar_ToString(AObj: TToolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ToolBar_GetButtonCount(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ButtonCount;
end;

function ToolBar_GetCanvas(AObj: TToolBar): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function ToolBar_GetCustomizeKeyName(AObj: TToolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.CustomizeKeyName);
end;

procedure ToolBar_SetCustomizeKeyName(AObj: TToolBar; AValue: PChar); stdcall;
begin
  AObj.CustomizeKeyName := AValue;
end;

function ToolBar_GetCustomizeValueName(AObj: TToolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.CustomizeValueName);
end;

procedure ToolBar_SetCustomizeValueName(AObj: TToolBar; AValue: PChar); stdcall;
begin
  AObj.CustomizeValueName := AValue;
end;

function ToolBar_GetRowCount(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.RowCount;
end;

function ToolBar_GetAlign(AObj: TToolBar): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ToolBar_SetAlign(AObj: TToolBar; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ToolBar_GetAnchors(AObj: TToolBar): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ToolBar_SetAnchors(AObj: TToolBar; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ToolBar_GetAutoSize(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure ToolBar_SetAutoSize(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function ToolBar_GetBorderWidth(AObj: TToolBar): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure ToolBar_SetBorderWidth(AObj: TToolBar; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function ToolBar_GetButtonHeight(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ButtonHeight;
end;

procedure ToolBar_SetButtonHeight(AObj: TToolBar; AValue: Integer); stdcall;
begin
  AObj.ButtonHeight := AValue;
end;

function ToolBar_GetButtonWidth(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ButtonWidth;
end;

procedure ToolBar_SetButtonWidth(AObj: TToolBar; AValue: Integer); stdcall;
begin
  AObj.ButtonWidth := AValue;
end;

function ToolBar_GetCaption(AObj: TToolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure ToolBar_SetCaption(AObj: TToolBar; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function ToolBar_GetColor(AObj: TToolBar): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure ToolBar_SetColor(AObj: TToolBar; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function ToolBar_GetDoubleBuffered(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ToolBar_SetDoubleBuffered(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function ToolBar_GetDockSite(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure ToolBar_SetDockSite(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function ToolBar_GetDragCursor(AObj: TToolBar): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure ToolBar_SetDragCursor(AObj: TToolBar; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function ToolBar_GetDragKind(AObj: TToolBar): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure ToolBar_SetDragKind(AObj: TToolBar; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function ToolBar_GetDragMode(AObj: TToolBar): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure ToolBar_SetDragMode(AObj: TToolBar; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function ToolBar_GetDrawingStyle(AObj: TToolBar): TTBDrawingStyle; stdcall;
begin
  Result :=  AObj.DrawingStyle;
end;

procedure ToolBar_SetDrawingStyle(AObj: TToolBar; AValue: TTBDrawingStyle); stdcall;
begin
  AObj.DrawingStyle := AValue;
end;

function ToolBar_GetEdgeBorders(AObj: TToolBar): TEdgeBorders; stdcall;
begin
  Result :=  AObj.EdgeBorders;
end;

procedure ToolBar_SetEdgeBorders(AObj: TToolBar; AValue: TEdgeBorders); stdcall;
begin
  AObj.EdgeBorders := AValue;
end;

function ToolBar_GetEdgeInner(AObj: TToolBar): TEdgeStyle; stdcall;
begin
  Result :=  AObj.EdgeInner;
end;

procedure ToolBar_SetEdgeInner(AObj: TToolBar; AValue: TEdgeStyle); stdcall;
begin
  AObj.EdgeInner := AValue;
end;

function ToolBar_GetEdgeOuter(AObj: TToolBar): TEdgeStyle; stdcall;
begin
  Result :=  AObj.EdgeOuter;
end;

procedure ToolBar_SetEdgeOuter(AObj: TToolBar; AValue: TEdgeStyle); stdcall;
begin
  AObj.EdgeOuter := AValue;
end;

function ToolBar_GetEnabled(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ToolBar_SetEnabled(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ToolBar_GetFlat(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.Flat;
end;

procedure ToolBar_SetFlat(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.Flat := AValue;
end;

function ToolBar_GetFont(AObj: TToolBar): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure ToolBar_SetFont(AObj: TToolBar; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function ToolBar_GetGradientEndColor(AObj: TToolBar): TColor; stdcall;
begin
  Result :=  AObj.GradientEndColor;
end;

procedure ToolBar_SetGradientEndColor(AObj: TToolBar; AValue: TColor); stdcall;
begin
  AObj.GradientEndColor := AValue;
end;

function ToolBar_GetGradientStartColor(AObj: TToolBar): TColor; stdcall;
begin
  Result :=  AObj.GradientStartColor;
end;

procedure ToolBar_SetGradientStartColor(AObj: TToolBar; AValue: TColor); stdcall;
begin
  AObj.GradientStartColor := AValue;
end;

function ToolBar_GetHeight(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ToolBar_SetHeight(AObj: TToolBar; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ToolBar_GetHideClippedButtons(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.HideClippedButtons;
end;

procedure ToolBar_SetHideClippedButtons(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.HideClippedButtons := AValue;
end;

function ToolBar_GetHotImages(AObj: TToolBar): TCustomImageList; stdcall;
begin
  Result :=  AObj.HotImages;
end;

procedure ToolBar_SetHotImages(AObj: TToolBar; AValue: TCustomImageList); stdcall;
begin
  AObj.HotImages := AValue;
end;

function ToolBar_GetImages(AObj: TToolBar): TCustomImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure ToolBar_SetImages(AObj: TToolBar; AValue: TCustomImageList); stdcall;
begin
  AObj.Images := AValue;
end;

function ToolBar_GetIndent(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.Indent;
end;

procedure ToolBar_SetIndent(AObj: TToolBar; AValue: Integer); stdcall;
begin
  AObj.Indent := AValue;
end;

function ToolBar_GetList(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.List;
end;

procedure ToolBar_SetList(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.List := AValue;
end;

function ToolBar_GetMenu(AObj: TToolBar): TMainMenu; stdcall;
begin
  Result :=  AObj.Menu;
end;

procedure ToolBar_SetMenu(AObj: TToolBar; AValue: TMainMenu); stdcall;
begin
  AObj.Menu := AValue;
end;

function ToolBar_GetGradientDirection(AObj: TToolBar): TGradientDirection; stdcall;
begin
  Result :=  AObj.GradientDirection;
end;

procedure ToolBar_SetGradientDirection(AObj: TToolBar; AValue: TGradientDirection); stdcall;
begin
  AObj.GradientDirection := AValue;
end;

function ToolBar_GetGradientDrawingOptions(AObj: TToolBar): TTBGradientDrawingOptions; stdcall;
begin
  Result :=  AObj.GradientDrawingOptions;
end;

procedure ToolBar_SetGradientDrawingOptions(AObj: TToolBar; AValue: TTBGradientDrawingOptions); stdcall;
begin
  AObj.GradientDrawingOptions := AValue;
end;

function ToolBar_GetParentColor(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure ToolBar_SetParentColor(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function ToolBar_GetParentDoubleBuffered(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure ToolBar_SetParentDoubleBuffered(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function ToolBar_GetParentFont(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure ToolBar_SetParentFont(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function ToolBar_GetParentShowHint(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ToolBar_SetParentShowHint(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ToolBar_GetPopupMenu(AObj: TToolBar): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ToolBar_SetPopupMenu(AObj: TToolBar; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ToolBar_GetShowCaptions(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.ShowCaptions;
end;

procedure ToolBar_SetShowCaptions(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.ShowCaptions := AValue;
end;

function ToolBar_GetShowHint(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ToolBar_SetShowHint(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ToolBar_GetTabOrder(AObj: TToolBar): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure ToolBar_SetTabOrder(AObj: TToolBar; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function ToolBar_GetTabStop(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure ToolBar_SetTabStop(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function ToolBar_GetTransparent(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure ToolBar_SetTransparent(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function ToolBar_GetVisible(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ToolBar_SetVisible(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function ToolBar_GetStyleElements(AObj: TToolBar): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ToolBar_SetStyleElements(AObj: TToolBar; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function ToolBar_GetWrapable(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.Wrapable;
end;

procedure ToolBar_SetWrapable(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.Wrapable := AValue;
end;

procedure ToolBar_SetOnAdvancedCustomDraw(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDraw := nil;
    TEventClass.Remove(AObj, @TEventClass.ToolBarOnAdvancedCustomDraw);
    Exit;
  end;
  AObj.OnAdvancedCustomDraw := TEventClass.ToolBarOnAdvancedCustomDraw;
  TEventClass.Add(AObj, @TEventClass.ToolBarOnAdvancedCustomDraw, AEventId);
end;

procedure ToolBar_SetOnAdvancedCustomDrawButton(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDrawButton := nil;
    TEventClass.Remove(AObj, @TEventClass.ToolBarOnAdvancedCustomDrawButton);
    Exit;
  end;
  AObj.OnAdvancedCustomDrawButton := TEventClass.ToolBarOnAdvancedCustomDrawButton;
  TEventClass.Add(AObj, @TEventClass.ToolBarOnAdvancedCustomDrawButton, AEventId);
end;

procedure ToolBar_SetOnClick(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    TEventClass.Remove(AObj, @TEventClass.OnClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, @TEventClass.OnClick, AEventId);
end;

procedure ToolBar_SetOnContextPopup(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnContextPopup := nil;
    TEventClass.Remove(AObj, @TEventClass.OnContextPopup);
    Exit;
  end;
  AObj.OnContextPopup := TEventClass.OnContextPopup;
  TEventClass.Add(AObj, @TEventClass.OnContextPopup, AEventId);
end;

procedure ToolBar_SetOnDblClick(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDblClick := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDblClick);
    Exit;
  end;
  AObj.OnDblClick := TEventClass.OnDblClick;
  TEventClass.Add(AObj, @TEventClass.OnDblClick, AEventId);
end;

procedure ToolBar_SetOnDockDrop(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDockDrop := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDockDrop);
    Exit;
  end;
  AObj.OnDockDrop := TEventClass.OnDockDrop;
  TEventClass.Add(AObj, @TEventClass.OnDockDrop, AEventId);
end;

procedure ToolBar_SetOnDragDrop(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragDrop := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDragDrop);
    Exit;
  end;
  AObj.OnDragDrop := TEventClass.OnDragDrop;
  TEventClass.Add(AObj, @TEventClass.OnDragDrop, AEventId);
end;

procedure ToolBar_SetOnDragOver(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragOver := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDragOver);
    Exit;
  end;
  AObj.OnDragOver := TEventClass.OnDragOver;
  TEventClass.Add(AObj, @TEventClass.OnDragOver, AEventId);
end;

procedure ToolBar_SetOnEndDock(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDock := nil;
    TEventClass.Remove(AObj, @TEventClass.OnEndDock);
    Exit;
  end;
  AObj.OnEndDock := TEventClass.OnEndDock;
  TEventClass.Add(AObj, @TEventClass.OnEndDock, AEventId);
end;

procedure ToolBar_SetOnEndDrag(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDrag := nil;
    TEventClass.Remove(AObj, @TEventClass.OnEndDrag);
    Exit;
  end;
  AObj.OnEndDrag := TEventClass.OnEndDrag;
  TEventClass.Add(AObj, @TEventClass.OnEndDrag, AEventId);
end;

procedure ToolBar_SetOnEnter(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEnter := nil;
    TEventClass.Remove(AObj, @TEventClass.OnEnter);
    Exit;
  end;
  AObj.OnEnter := TEventClass.OnEnter;
  TEventClass.Add(AObj, @TEventClass.OnEnter, AEventId);
end;

procedure ToolBar_SetOnExit(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnExit := nil;
    TEventClass.Remove(AObj, @TEventClass.OnExit);
    Exit;
  end;
  AObj.OnExit := TEventClass.OnExit;
  TEventClass.Add(AObj, @TEventClass.OnExit, AEventId);
end;

procedure ToolBar_SetOnGesture(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGesture := nil;
    TEventClass.Remove(AObj, @TEventClass.OnGesture);
    Exit;
  end;
  AObj.OnGesture := TEventClass.OnGesture;
  TEventClass.Add(AObj, @TEventClass.OnGesture, AEventId);
end;

procedure ToolBar_SetOnGetSiteInfo(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetSiteInfo := nil;
    TEventClass.Remove(AObj, @TEventClass.OnGetSiteInfo);
    Exit;
  end;
  AObj.OnGetSiteInfo := TEventClass.OnGetSiteInfo;
  TEventClass.Add(AObj, @TEventClass.OnGetSiteInfo, AEventId);
end;

procedure ToolBar_SetOnMouseActivate(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseActivate := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseActivate);
    Exit;
  end;
  AObj.OnMouseActivate := TEventClass.OnMouseActivate;
  TEventClass.Add(AObj, @TEventClass.OnMouseActivate, AEventId);
end;

procedure ToolBar_SetOnMouseDown(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseDown := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseDown);
    Exit;
  end;
  AObj.OnMouseDown := TEventClass.OnMouseDown;
  TEventClass.Add(AObj, @TEventClass.OnMouseDown, AEventId);
end;

procedure ToolBar_SetOnMouseEnter(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseEnter := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseEnter);
    Exit;
  end;
  AObj.OnMouseEnter := TEventClass.OnMouseEnter;
  TEventClass.Add(AObj, @TEventClass.OnMouseEnter, AEventId);
end;

procedure ToolBar_SetOnMouseLeave(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseLeave := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseLeave);
    Exit;
  end;
  AObj.OnMouseLeave := TEventClass.OnMouseLeave;
  TEventClass.Add(AObj, @TEventClass.OnMouseLeave, AEventId);
end;

procedure ToolBar_SetOnMouseMove(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseMove := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseMove);
    Exit;
  end;
  AObj.OnMouseMove := TEventClass.OnMouseMove;
  TEventClass.Add(AObj, @TEventClass.OnMouseMove, AEventId);
end;

procedure ToolBar_SetOnMouseUp(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseUp := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseUp);
    Exit;
  end;
  AObj.OnMouseUp := TEventClass.OnMouseUp;
  TEventClass.Add(AObj, @TEventClass.OnMouseUp, AEventId);
end;

procedure ToolBar_SetOnResize(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnResize := nil;
    TEventClass.Remove(AObj, @TEventClass.OnResize);
    Exit;
  end;
  AObj.OnResize := TEventClass.OnResize;
  TEventClass.Add(AObj, @TEventClass.OnResize, AEventId);
end;

procedure ToolBar_SetOnStartDock(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnStartDock := nil;
    TEventClass.Remove(AObj, @TEventClass.OnStartDock);
    Exit;
  end;
  AObj.OnStartDock := TEventClass.OnStartDock;
  TEventClass.Add(AObj, @TEventClass.OnStartDock, AEventId);
end;

procedure ToolBar_SetOnUnDock(AObj: TToolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnUnDock := nil;
    TEventClass.Remove(AObj, @TEventClass.OnUnDock);
    Exit;
  end;
  AObj.OnUnDock := TEventClass.OnUnDock;
  TEventClass.Add(AObj, @TEventClass.OnUnDock, AEventId);
end;

function ToolBar_GetDockClientCount(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function ToolBar_GetAlignDisabled(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function ToolBar_GetMouseInClient(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function ToolBar_GetVisibleDockClientCount(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function ToolBar_GetBrush(AObj: TToolBar): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function ToolBar_GetControlCount(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function ToolBar_GetHandle(AObj: TToolBar): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ToolBar_GetParentWindow(AObj: TToolBar): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ToolBar_SetParentWindow(AObj: TToolBar; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function ToolBar_GetUseDockManager(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure ToolBar_SetUseDockManager(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function ToolBar_GetAction(AObj: TToolBar): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ToolBar_SetAction(AObj: TToolBar; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function ToolBar_GetBiDiMode(AObj: TToolBar): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ToolBar_SetBiDiMode(AObj: TToolBar; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure ToolBar_GetBoundsRect(AObj: TToolBar; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ToolBar_SetBoundsRect(AObj: TToolBar; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ToolBar_GetClientHeight(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ToolBar_SetClientHeight(AObj: TToolBar; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ToolBar_GetClientOrigin(AObj: TToolBar; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure ToolBar_GetClientRect(AObj: TToolBar; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ToolBar_GetClientWidth(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ToolBar_SetClientWidth(AObj: TToolBar; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ToolBar_GetControlState(AObj: TToolBar): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure ToolBar_SetControlState(AObj: TToolBar; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function ToolBar_GetControlStyle(AObj: TToolBar): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure ToolBar_SetControlStyle(AObj: TToolBar; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function ToolBar_GetExplicitLeft(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ToolBar_GetExplicitTop(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ToolBar_GetExplicitWidth(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ToolBar_GetExplicitHeight(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ToolBar_GetFloating(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function ToolBar_GetParent(AObj: TToolBar): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ToolBar_SetParent(AObj: TToolBar; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ToolBar_GetAlignWithMargins(AObj: TToolBar): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ToolBar_SetAlignWithMargins(AObj: TToolBar; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ToolBar_GetLeft(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ToolBar_SetLeft(AObj: TToolBar; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ToolBar_GetTop(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ToolBar_SetTop(AObj: TToolBar; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ToolBar_GetWidth(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ToolBar_SetWidth(AObj: TToolBar; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ToolBar_GetCursor(AObj: TToolBar): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ToolBar_SetCursor(AObj: TToolBar; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ToolBar_GetHint(AObj: TToolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ToolBar_SetHint(AObj: TToolBar; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ToolBar_GetMargins(AObj: TToolBar): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ToolBar_SetMargins(AObj: TToolBar; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ToolBar_GetCustomHint(AObj: TToolBar): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ToolBar_SetCustomHint(AObj: TToolBar; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ToolBar_GetComponentCount(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ToolBar_GetComponentIndex(AObj: TToolBar): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ToolBar_SetComponentIndex(AObj: TToolBar; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ToolBar_GetOwner(AObj: TToolBar): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ToolBar_GetName(AObj: TToolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ToolBar_SetName(AObj: TToolBar; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ToolBar_GetTag(AObj: TToolBar): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ToolBar_SetTag(AObj: TToolBar; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ToolBar_GetButtons(AObj: TToolBar; Index: Integer): TToolButton; stdcall;
begin
  Result :=  AObj.Buttons[Index];
end;

function ToolBar_GetDockClients(AObj: TToolBar; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function ToolBar_GetControls(AObj: TToolBar; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ToolBar_GetComponents(AObj: TToolBar; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function ToolBar_StaticClassType: TClass; stdcall;
begin
  Result := TToolBar;
end;

exports
  ToolBar_Create,
  ToolBar_Free,
  ToolBar_FlipChildren,
  ToolBar_CanFocus,
  ToolBar_ContainsControl,
  ToolBar_ControlAtPos,
  ToolBar_DisableAlign,
  ToolBar_EnableAlign,
  ToolBar_FindChildControl,
  ToolBar_Focused,
  ToolBar_HandleAllocated,
  ToolBar_InsertControl,
  ToolBar_Invalidate,
  ToolBar_PaintTo,
  ToolBar_RemoveControl,
  ToolBar_Realign,
  ToolBar_Repaint,
  ToolBar_ScaleBy,
  ToolBar_ScrollBy,
  ToolBar_SetBounds,
  ToolBar_SetFocus,
  ToolBar_Update,
  ToolBar_UpdateControlState,
  ToolBar_BringToFront,
  ToolBar_ClientToScreen,
  ToolBar_ClientToParent,
  ToolBar_Dragging,
  ToolBar_HasParent,
  ToolBar_Hide,
  ToolBar_Perform,
  ToolBar_Refresh,
  ToolBar_ScreenToClient,
  ToolBar_ParentToClient,
  ToolBar_SendToBack,
  ToolBar_Show,
  ToolBar_GetTextBuf,
  ToolBar_GetTextLen,
  ToolBar_SetTextBuf,
  ToolBar_FindComponent,
  ToolBar_GetNamePath,
  ToolBar_Assign,
  ToolBar_DisposeOf,
  ToolBar_ClassType,
  ToolBar_ClassName,
  ToolBar_InstanceSize,
  ToolBar_InheritsFrom,
  ToolBar_Equals,
  ToolBar_GetHashCode,
  ToolBar_ToString,
  ToolBar_GetButtonCount,
  ToolBar_GetCanvas,
  ToolBar_GetCustomizeKeyName,
  ToolBar_SetCustomizeKeyName,
  ToolBar_GetCustomizeValueName,
  ToolBar_SetCustomizeValueName,
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
  ToolBar_GetDockSite,
  ToolBar_SetDockSite,
  ToolBar_GetDragCursor,
  ToolBar_SetDragCursor,
  ToolBar_GetDragKind,
  ToolBar_SetDragKind,
  ToolBar_GetDragMode,
  ToolBar_SetDragMode,
  ToolBar_GetDrawingStyle,
  ToolBar_SetDrawingStyle,
  ToolBar_GetEdgeBorders,
  ToolBar_SetEdgeBorders,
  ToolBar_GetEdgeInner,
  ToolBar_SetEdgeInner,
  ToolBar_GetEdgeOuter,
  ToolBar_SetEdgeOuter,
  ToolBar_GetEnabled,
  ToolBar_SetEnabled,
  ToolBar_GetFlat,
  ToolBar_SetFlat,
  ToolBar_GetFont,
  ToolBar_SetFont,
  ToolBar_GetGradientEndColor,
  ToolBar_SetGradientEndColor,
  ToolBar_GetGradientStartColor,
  ToolBar_SetGradientStartColor,
  ToolBar_GetHeight,
  ToolBar_SetHeight,
  ToolBar_GetHideClippedButtons,
  ToolBar_SetHideClippedButtons,
  ToolBar_GetHotImages,
  ToolBar_SetHotImages,
  ToolBar_GetImages,
  ToolBar_SetImages,
  ToolBar_GetIndent,
  ToolBar_SetIndent,
  ToolBar_GetList,
  ToolBar_SetList,
  ToolBar_GetMenu,
  ToolBar_SetMenu,
  ToolBar_GetGradientDirection,
  ToolBar_SetGradientDirection,
  ToolBar_GetGradientDrawingOptions,
  ToolBar_SetGradientDrawingOptions,
  ToolBar_GetParentColor,
  ToolBar_SetParentColor,
  ToolBar_GetParentDoubleBuffered,
  ToolBar_SetParentDoubleBuffered,
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
  ToolBar_GetStyleElements,
  ToolBar_SetStyleElements,
  ToolBar_GetWrapable,
  ToolBar_SetWrapable,
  ToolBar_SetOnAdvancedCustomDraw,
  ToolBar_SetOnAdvancedCustomDrawButton,
  ToolBar_SetOnClick,
  ToolBar_SetOnContextPopup,
  ToolBar_SetOnDblClick,
  ToolBar_SetOnDockDrop,
  ToolBar_SetOnDragDrop,
  ToolBar_SetOnDragOver,
  ToolBar_SetOnEndDock,
  ToolBar_SetOnEndDrag,
  ToolBar_SetOnEnter,
  ToolBar_SetOnExit,
  ToolBar_SetOnGesture,
  ToolBar_SetOnGetSiteInfo,
  ToolBar_SetOnMouseActivate,
  ToolBar_SetOnMouseDown,
  ToolBar_SetOnMouseEnter,
  ToolBar_SetOnMouseLeave,
  ToolBar_SetOnMouseMove,
  ToolBar_SetOnMouseUp,
  ToolBar_SetOnResize,
  ToolBar_SetOnStartDock,
  ToolBar_SetOnUnDock,
  ToolBar_GetDockClientCount,
  ToolBar_GetAlignDisabled,
  ToolBar_GetMouseInClient,
  ToolBar_GetVisibleDockClientCount,
  ToolBar_GetBrush,
  ToolBar_GetControlCount,
  ToolBar_GetHandle,
  ToolBar_GetParentWindow,
  ToolBar_SetParentWindow,
  ToolBar_GetUseDockManager,
  ToolBar_SetUseDockManager,
  ToolBar_GetAction,
  ToolBar_SetAction,
  ToolBar_GetBiDiMode,
  ToolBar_SetBiDiMode,
  ToolBar_GetBoundsRect,
  ToolBar_SetBoundsRect,
  ToolBar_GetClientHeight,
  ToolBar_SetClientHeight,
  ToolBar_GetClientOrigin,
  ToolBar_GetClientRect,
  ToolBar_GetClientWidth,
  ToolBar_SetClientWidth,
  ToolBar_GetControlState,
  ToolBar_SetControlState,
  ToolBar_GetControlStyle,
  ToolBar_SetControlStyle,
  ToolBar_GetExplicitLeft,
  ToolBar_GetExplicitTop,
  ToolBar_GetExplicitWidth,
  ToolBar_GetExplicitHeight,
  ToolBar_GetFloating,
  ToolBar_GetParent,
  ToolBar_SetParent,
  ToolBar_GetAlignWithMargins,
  ToolBar_SetAlignWithMargins,
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
  ToolBar_GetMargins,
  ToolBar_SetMargins,
  ToolBar_GetCustomHint,
  ToolBar_SetCustomHint,
  ToolBar_GetComponentCount,
  ToolBar_GetComponentIndex,
  ToolBar_SetComponentIndex,
  ToolBar_GetOwner,
  ToolBar_GetName,
  ToolBar_SetName,
  ToolBar_GetTag,
  ToolBar_SetTag,
  ToolBar_GetButtons,
  ToolBar_GetDockClients,
  ToolBar_GetControls,
  ToolBar_GetComponents,
  ToolBar_StaticClassType;

