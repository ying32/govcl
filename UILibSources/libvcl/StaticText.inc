
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function StaticText_Create(AOwner: TComponent): TStaticText; stdcall;
begin
  Result :=  TStaticText.Create(AOwner);
end;

procedure StaticText_Free(AObj: TStaticText); stdcall;
begin
  AObj.Free;
end;

function StaticText_CanFocus(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function StaticText_ContainsControl(AObj: TStaticText; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function StaticText_ControlAtPos(AObj: TStaticText; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure StaticText_DisableAlign(AObj: TStaticText); stdcall;
begin
  AObj.DisableAlign;
end;

procedure StaticText_EnableAlign(AObj: TStaticText); stdcall;
begin
  AObj.EnableAlign;
end;

function StaticText_FindChildControl(AObj: TStaticText; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure StaticText_FlipChildren(AObj: TStaticText; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function StaticText_Focused(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function StaticText_HandleAllocated(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure StaticText_InsertControl(AObj: TStaticText; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure StaticText_Invalidate(AObj: TStaticText); stdcall;
begin
  AObj.Invalidate;
end;

procedure StaticText_PaintTo(AObj: TStaticText; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure StaticText_RemoveControl(AObj: TStaticText; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure StaticText_Realign(AObj: TStaticText); stdcall;
begin
  AObj.Realign;
end;

procedure StaticText_Repaint(AObj: TStaticText); stdcall;
begin
  AObj.Repaint;
end;

procedure StaticText_ScaleBy(AObj: TStaticText; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure StaticText_ScrollBy(AObj: TStaticText; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure StaticText_SetBounds(AObj: TStaticText; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure StaticText_SetFocus(AObj: TStaticText); stdcall;
begin
  AObj.SetFocus;
end;

procedure StaticText_Update(AObj: TStaticText); stdcall;
begin
  AObj.Update;
end;

procedure StaticText_UpdateControlState(AObj: TStaticText); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure StaticText_BringToFront(AObj: TStaticText); stdcall;
begin
  AObj.BringToFront;
end;

procedure StaticText_ClientToScreen(AObj: TStaticText; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure StaticText_ClientToParent(AObj: TStaticText; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function StaticText_Dragging(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function StaticText_HasParent(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure StaticText_Hide(AObj: TStaticText); stdcall;
begin
  AObj.Hide;
end;

function StaticText_Perform(AObj: TStaticText; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure StaticText_Refresh(AObj: TStaticText); stdcall;
begin
  AObj.Refresh;
end;

procedure StaticText_ScreenToClient(AObj: TStaticText; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure StaticText_ParentToClient(AObj: TStaticText; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure StaticText_SendToBack(AObj: TStaticText); stdcall;
begin
  AObj.SendToBack;
end;

procedure StaticText_Show(AObj: TStaticText); stdcall;
begin
  AObj.Show;
end;

function StaticText_GetTextBuf(AObj: TStaticText; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function StaticText_GetTextLen(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure StaticText_SetTextBuf(AObj: TStaticText; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function StaticText_FindComponent(AObj: TStaticText; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function StaticText_GetNamePath(AObj: TStaticText): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure StaticText_Assign(AObj: TStaticText; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure StaticText_DisposeOf(AObj: TStaticText); stdcall;
begin
  AObj.DisposeOf;
end;

function StaticText_ClassType(AObj: TStaticText): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function StaticText_ClassName(AObj: TStaticText): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function StaticText_InstanceSize(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function StaticText_InheritsFrom(AObj: TStaticText; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function StaticText_Equals(AObj: TStaticText; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function StaticText_GetHashCode(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function StaticText_ToString(AObj: TStaticText): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function StaticText_GetAlign(AObj: TStaticText): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure StaticText_SetAlign(AObj: TStaticText; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function StaticText_GetAlignment(AObj: TStaticText): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure StaticText_SetAlignment(AObj: TStaticText; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function StaticText_GetAnchors(AObj: TStaticText): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure StaticText_SetAnchors(AObj: TStaticText; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function StaticText_GetAutoSize(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure StaticText_SetAutoSize(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function StaticText_GetBevelEdges(AObj: TStaticText): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure StaticText_SetBevelEdges(AObj: TStaticText; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function StaticText_GetBevelInner(AObj: TStaticText): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure StaticText_SetBevelInner(AObj: TStaticText; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function StaticText_GetBevelKind(AObj: TStaticText): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure StaticText_SetBevelKind(AObj: TStaticText; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function StaticText_GetBevelOuter(AObj: TStaticText): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure StaticText_SetBevelOuter(AObj: TStaticText; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function StaticText_GetBiDiMode(AObj: TStaticText): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure StaticText_SetBiDiMode(AObj: TStaticText; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function StaticText_GetBorderStyle(AObj: TStaticText): TStaticBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure StaticText_SetBorderStyle(AObj: TStaticText; AValue: TStaticBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function StaticText_GetCaption(AObj: TStaticText): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure StaticText_SetCaption(AObj: TStaticText; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function StaticText_GetColor(AObj: TStaticText): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure StaticText_SetColor(AObj: TStaticText; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function StaticText_GetDoubleBuffered(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure StaticText_SetDoubleBuffered(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function StaticText_GetDragCursor(AObj: TStaticText): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure StaticText_SetDragCursor(AObj: TStaticText; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function StaticText_GetDragKind(AObj: TStaticText): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure StaticText_SetDragKind(AObj: TStaticText; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function StaticText_GetDragMode(AObj: TStaticText): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure StaticText_SetDragMode(AObj: TStaticText; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function StaticText_GetEnabled(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure StaticText_SetEnabled(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function StaticText_GetFont(AObj: TStaticText): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure StaticText_SetFont(AObj: TStaticText; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function StaticText_GetParentColor(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure StaticText_SetParentColor(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function StaticText_GetParentDoubleBuffered(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure StaticText_SetParentDoubleBuffered(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function StaticText_GetParentFont(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure StaticText_SetParentFont(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function StaticText_GetParentShowHint(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure StaticText_SetParentShowHint(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function StaticText_GetPopupMenu(AObj: TStaticText): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure StaticText_SetPopupMenu(AObj: TStaticText; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function StaticText_GetShowAccelChar(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.ShowAccelChar;
end;

procedure StaticText_SetShowAccelChar(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.ShowAccelChar := AValue;
end;

function StaticText_GetShowHint(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure StaticText_SetShowHint(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function StaticText_GetTabOrder(AObj: TStaticText): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure StaticText_SetTabOrder(AObj: TStaticText; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function StaticText_GetTabStop(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure StaticText_SetTabStop(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function StaticText_GetTransparent(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure StaticText_SetTransparent(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function StaticText_GetVisible(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure StaticText_SetVisible(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function StaticText_GetStyleElements(AObj: TStaticText): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure StaticText_SetStyleElements(AObj: TStaticText; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure StaticText_SetOnClick(AObj: TStaticText; AEventId: NativeUInt); stdcall;
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

procedure StaticText_SetOnContextPopup(AObj: TStaticText; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnContextPopup := nil;
    TEventClass.Remove(AObj, geContextPopup);
    Exit;
  end;
  AObj.OnContextPopup := TEventClass.OnContextPopup;
  TEventClass.Add(AObj, geContextPopup, AEventId);
end;

procedure StaticText_SetOnDblClick(AObj: TStaticText; AEventId: NativeUInt); stdcall;
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

procedure StaticText_SetOnDragDrop(AObj: TStaticText; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragDrop := nil;
    TEventClass.Remove(AObj, geDragDrop);
    Exit;
  end;
  AObj.OnDragDrop := TEventClass.OnDragDrop;
  TEventClass.Add(AObj, geDragDrop, AEventId);
end;

procedure StaticText_SetOnDragOver(AObj: TStaticText; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragOver := nil;
    TEventClass.Remove(AObj, geDragOver);
    Exit;
  end;
  AObj.OnDragOver := TEventClass.OnDragOver;
  TEventClass.Add(AObj, geDragOver, AEventId);
end;

procedure StaticText_SetOnEndDock(AObj: TStaticText; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDock := nil;
    TEventClass.Remove(AObj, geEndDock);
    Exit;
  end;
  AObj.OnEndDock := TEventClass.OnEndDock;
  TEventClass.Add(AObj, geEndDock, AEventId);
end;

procedure StaticText_SetOnEndDrag(AObj: TStaticText; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDrag := nil;
    TEventClass.Remove(AObj, geEndDrag);
    Exit;
  end;
  AObj.OnEndDrag := TEventClass.OnEndDrag;
  TEventClass.Add(AObj, geEndDrag, AEventId);
end;

procedure StaticText_SetOnGesture(AObj: TStaticText; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGesture := nil;
    TEventClass.Remove(AObj, geGesture);
    Exit;
  end;
  AObj.OnGesture := TEventClass.OnGesture;
  TEventClass.Add(AObj, geGesture, AEventId);
end;

procedure StaticText_SetOnMouseActivate(AObj: TStaticText; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseActivate := nil;
    TEventClass.Remove(AObj, geMouseActivate);
    Exit;
  end;
  AObj.OnMouseActivate := TEventClass.OnMouseActivate;
  TEventClass.Add(AObj, geMouseActivate, AEventId);
end;

procedure StaticText_SetOnMouseDown(AObj: TStaticText; AEventId: NativeUInt); stdcall;
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

procedure StaticText_SetOnMouseEnter(AObj: TStaticText; AEventId: NativeUInt); stdcall;
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

procedure StaticText_SetOnMouseLeave(AObj: TStaticText; AEventId: NativeUInt); stdcall;
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

procedure StaticText_SetOnMouseMove(AObj: TStaticText; AEventId: NativeUInt); stdcall;
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

procedure StaticText_SetOnMouseUp(AObj: TStaticText; AEventId: NativeUInt); stdcall;
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

procedure StaticText_SetOnStartDock(AObj: TStaticText; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnStartDock := nil;
    TEventClass.Remove(AObj, geStartDock);
    Exit;
  end;
  AObj.OnStartDock := TEventClass.OnStartDock;
  TEventClass.Add(AObj, geStartDock, AEventId);
end;

function StaticText_GetDockClientCount(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function StaticText_GetDockSite(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure StaticText_SetDockSite(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function StaticText_GetAlignDisabled(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function StaticText_GetMouseInClient(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function StaticText_GetVisibleDockClientCount(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function StaticText_GetBrush(AObj: TStaticText): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function StaticText_GetControlCount(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function StaticText_GetHandle(AObj: TStaticText): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function StaticText_GetParentWindow(AObj: TStaticText): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure StaticText_SetParentWindow(AObj: TStaticText; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function StaticText_GetUseDockManager(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure StaticText_SetUseDockManager(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function StaticText_GetAction(AObj: TStaticText): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure StaticText_SetAction(AObj: TStaticText; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure StaticText_GetBoundsRect(AObj: TStaticText; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure StaticText_SetBoundsRect(AObj: TStaticText; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function StaticText_GetClientHeight(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure StaticText_SetClientHeight(AObj: TStaticText; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure StaticText_GetClientOrigin(AObj: TStaticText; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure StaticText_GetClientRect(AObj: TStaticText; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function StaticText_GetClientWidth(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure StaticText_SetClientWidth(AObj: TStaticText; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function StaticText_GetControlState(AObj: TStaticText): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure StaticText_SetControlState(AObj: TStaticText; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function StaticText_GetControlStyle(AObj: TStaticText): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure StaticText_SetControlStyle(AObj: TStaticText; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function StaticText_GetExplicitLeft(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function StaticText_GetExplicitTop(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function StaticText_GetExplicitWidth(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function StaticText_GetExplicitHeight(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function StaticText_GetFloating(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function StaticText_GetParent(AObj: TStaticText): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure StaticText_SetParent(AObj: TStaticText; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function StaticText_GetAlignWithMargins(AObj: TStaticText): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure StaticText_SetAlignWithMargins(AObj: TStaticText; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function StaticText_GetLeft(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure StaticText_SetLeft(AObj: TStaticText; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function StaticText_GetTop(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure StaticText_SetTop(AObj: TStaticText; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function StaticText_GetWidth(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure StaticText_SetWidth(AObj: TStaticText; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function StaticText_GetHeight(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure StaticText_SetHeight(AObj: TStaticText; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function StaticText_GetCursor(AObj: TStaticText): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure StaticText_SetCursor(AObj: TStaticText; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function StaticText_GetHint(AObj: TStaticText): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure StaticText_SetHint(AObj: TStaticText; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function StaticText_GetMargins(AObj: TStaticText): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure StaticText_SetMargins(AObj: TStaticText; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function StaticText_GetCustomHint(AObj: TStaticText): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure StaticText_SetCustomHint(AObj: TStaticText; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function StaticText_GetComponentCount(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function StaticText_GetComponentIndex(AObj: TStaticText): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure StaticText_SetComponentIndex(AObj: TStaticText; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function StaticText_GetOwner(AObj: TStaticText): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function StaticText_GetName(AObj: TStaticText): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure StaticText_SetName(AObj: TStaticText; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function StaticText_GetTag(AObj: TStaticText): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure StaticText_SetTag(AObj: TStaticText; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function StaticText_GetDockClients(AObj: TStaticText; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function StaticText_GetControls(AObj: TStaticText; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function StaticText_GetComponents(AObj: TStaticText; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function StaticText_StaticClassType: TClass; stdcall;
begin
  Result := TStaticText;
end;

exports
  StaticText_Create,
  StaticText_Free,
  StaticText_CanFocus,
  StaticText_ContainsControl,
  StaticText_ControlAtPos,
  StaticText_DisableAlign,
  StaticText_EnableAlign,
  StaticText_FindChildControl,
  StaticText_FlipChildren,
  StaticText_Focused,
  StaticText_HandleAllocated,
  StaticText_InsertControl,
  StaticText_Invalidate,
  StaticText_PaintTo,
  StaticText_RemoveControl,
  StaticText_Realign,
  StaticText_Repaint,
  StaticText_ScaleBy,
  StaticText_ScrollBy,
  StaticText_SetBounds,
  StaticText_SetFocus,
  StaticText_Update,
  StaticText_UpdateControlState,
  StaticText_BringToFront,
  StaticText_ClientToScreen,
  StaticText_ClientToParent,
  StaticText_Dragging,
  StaticText_HasParent,
  StaticText_Hide,
  StaticText_Perform,
  StaticText_Refresh,
  StaticText_ScreenToClient,
  StaticText_ParentToClient,
  StaticText_SendToBack,
  StaticText_Show,
  StaticText_GetTextBuf,
  StaticText_GetTextLen,
  StaticText_SetTextBuf,
  StaticText_FindComponent,
  StaticText_GetNamePath,
  StaticText_Assign,
  StaticText_DisposeOf,
  StaticText_ClassType,
  StaticText_ClassName,
  StaticText_InstanceSize,
  StaticText_InheritsFrom,
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
  StaticText_GetBevelEdges,
  StaticText_SetBevelEdges,
  StaticText_GetBevelInner,
  StaticText_SetBevelInner,
  StaticText_GetBevelKind,
  StaticText_SetBevelKind,
  StaticText_GetBevelOuter,
  StaticText_SetBevelOuter,
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
  StaticText_GetDragCursor,
  StaticText_SetDragCursor,
  StaticText_GetDragKind,
  StaticText_SetDragKind,
  StaticText_GetDragMode,
  StaticText_SetDragMode,
  StaticText_GetEnabled,
  StaticText_SetEnabled,
  StaticText_GetFont,
  StaticText_SetFont,
  StaticText_GetParentColor,
  StaticText_SetParentColor,
  StaticText_GetParentDoubleBuffered,
  StaticText_SetParentDoubleBuffered,
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
  StaticText_GetStyleElements,
  StaticText_SetStyleElements,
  StaticText_SetOnClick,
  StaticText_SetOnContextPopup,
  StaticText_SetOnDblClick,
  StaticText_SetOnDragDrop,
  StaticText_SetOnDragOver,
  StaticText_SetOnEndDock,
  StaticText_SetOnEndDrag,
  StaticText_SetOnGesture,
  StaticText_SetOnMouseActivate,
  StaticText_SetOnMouseDown,
  StaticText_SetOnMouseEnter,
  StaticText_SetOnMouseLeave,
  StaticText_SetOnMouseMove,
  StaticText_SetOnMouseUp,
  StaticText_SetOnStartDock,
  StaticText_GetDockClientCount,
  StaticText_GetDockSite,
  StaticText_SetDockSite,
  StaticText_GetAlignDisabled,
  StaticText_GetMouseInClient,
  StaticText_GetVisibleDockClientCount,
  StaticText_GetBrush,
  StaticText_GetControlCount,
  StaticText_GetHandle,
  StaticText_GetParentWindow,
  StaticText_SetParentWindow,
  StaticText_GetUseDockManager,
  StaticText_SetUseDockManager,
  StaticText_GetAction,
  StaticText_SetAction,
  StaticText_GetBoundsRect,
  StaticText_SetBoundsRect,
  StaticText_GetClientHeight,
  StaticText_SetClientHeight,
  StaticText_GetClientOrigin,
  StaticText_GetClientRect,
  StaticText_GetClientWidth,
  StaticText_SetClientWidth,
  StaticText_GetControlState,
  StaticText_SetControlState,
  StaticText_GetControlStyle,
  StaticText_SetControlStyle,
  StaticText_GetExplicitLeft,
  StaticText_GetExplicitTop,
  StaticText_GetExplicitWidth,
  StaticText_GetExplicitHeight,
  StaticText_GetFloating,
  StaticText_GetParent,
  StaticText_SetParent,
  StaticText_GetAlignWithMargins,
  StaticText_SetAlignWithMargins,
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
  StaticText_GetMargins,
  StaticText_SetMargins,
  StaticText_GetCustomHint,
  StaticText_SetCustomHint,
  StaticText_GetComponentCount,
  StaticText_GetComponentIndex,
  StaticText_SetComponentIndex,
  StaticText_GetOwner,
  StaticText_GetName,
  StaticText_SetName,
  StaticText_GetTag,
  StaticText_SetTag,
  StaticText_GetDockClients,
  StaticText_GetControls,
  StaticText_GetComponents,
  StaticText_StaticClassType;

