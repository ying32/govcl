
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Panel_Create(AOwner: TComponent): TPanel; stdcall;
begin
  Result :=  TPanel.Create(AOwner);
end;

procedure Panel_Free(AObj: TPanel); stdcall;
begin
  AObj.Free;
end;

function Panel_CanFocus(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function Panel_ContainsControl(AObj: TPanel; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function Panel_ControlAtPos(AObj: TPanel; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure Panel_DisableAlign(AObj: TPanel); stdcall;
begin
  AObj.DisableAlign;
end;

procedure Panel_EnableAlign(AObj: TPanel); stdcall;
begin
  AObj.EnableAlign;
end;

function Panel_FindChildControl(AObj: TPanel; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure Panel_FlipChildren(AObj: TPanel; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function Panel_Focused(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function Panel_HandleAllocated(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure Panel_InsertControl(AObj: TPanel; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure Panel_Invalidate(AObj: TPanel); stdcall;
begin
  AObj.Invalidate;
end;

procedure Panel_PaintTo(AObj: TPanel; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure Panel_RemoveControl(AObj: TPanel; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure Panel_Realign(AObj: TPanel); stdcall;
begin
  AObj.Realign;
end;

procedure Panel_Repaint(AObj: TPanel); stdcall;
begin
  AObj.Repaint;
end;

procedure Panel_ScaleBy(AObj: TPanel; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure Panel_ScrollBy(AObj: TPanel; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure Panel_SetBounds(AObj: TPanel; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Panel_SetFocus(AObj: TPanel); stdcall;
begin
  AObj.SetFocus;
end;

procedure Panel_Update(AObj: TPanel); stdcall;
begin
  AObj.Update;
end;

procedure Panel_UpdateControlState(AObj: TPanel); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure Panel_BringToFront(AObj: TPanel); stdcall;
begin
  AObj.BringToFront;
end;

procedure Panel_ClientToScreen(AObj: TPanel; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure Panel_ClientToParent(AObj: TPanel; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function Panel_Dragging(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function Panel_HasParent(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Panel_Hide(AObj: TPanel); stdcall;
begin
  AObj.Hide;
end;

function Panel_Perform(AObj: TPanel; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Panel_Refresh(AObj: TPanel); stdcall;
begin
  AObj.Refresh;
end;

procedure Panel_ScreenToClient(AObj: TPanel; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure Panel_ParentToClient(AObj: TPanel; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure Panel_SendToBack(AObj: TPanel); stdcall;
begin
  AObj.SendToBack;
end;

procedure Panel_Show(AObj: TPanel); stdcall;
begin
  AObj.Show;
end;

function Panel_GetTextBuf(AObj: TPanel; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Panel_GetTextLen(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure Panel_SetTextBuf(AObj: TPanel; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function Panel_FindComponent(AObj: TPanel; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Panel_GetNamePath(AObj: TPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Panel_Assign(AObj: TPanel; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure Panel_DisposeOf(AObj: TPanel); stdcall;
begin
  AObj.DisposeOf;
end;

function Panel_ClassType(AObj: TPanel): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function Panel_ClassName(AObj: TPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Panel_InstanceSize(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function Panel_InheritsFrom(AObj: TPanel; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function Panel_Equals(AObj: TPanel; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Panel_GetHashCode(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Panel_ToString(AObj: TPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Panel_GetAlign(AObj: TPanel): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Panel_SetAlign(AObj: TPanel; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Panel_GetAlignment(AObj: TPanel): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure Panel_SetAlignment(AObj: TPanel; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function Panel_GetAnchors(AObj: TPanel): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Panel_SetAnchors(AObj: TPanel; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Panel_GetAutoSize(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure Panel_SetAutoSize(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function Panel_GetBevelEdges(AObj: TPanel): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure Panel_SetBevelEdges(AObj: TPanel; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function Panel_GetBevelInner(AObj: TPanel): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure Panel_SetBevelInner(AObj: TPanel; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function Panel_GetBevelKind(AObj: TPanel): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure Panel_SetBevelKind(AObj: TPanel; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function Panel_GetBevelOuter(AObj: TPanel): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure Panel_SetBevelOuter(AObj: TPanel; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function Panel_GetBiDiMode(AObj: TPanel): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Panel_SetBiDiMode(AObj: TPanel; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function Panel_GetBorderWidth(AObj: TPanel): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure Panel_SetBorderWidth(AObj: TPanel; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function Panel_GetBorderStyle(AObj: TPanel): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure Panel_SetBorderStyle(AObj: TPanel; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function Panel_GetCaption(AObj: TPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure Panel_SetCaption(AObj: TPanel; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function Panel_GetColor(AObj: TPanel): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure Panel_SetColor(AObj: TPanel; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function Panel_GetUseDockManager(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure Panel_SetUseDockManager(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function Panel_GetDockSite(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure Panel_SetDockSite(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function Panel_GetDoubleBuffered(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure Panel_SetDoubleBuffered(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function Panel_GetDragCursor(AObj: TPanel): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure Panel_SetDragCursor(AObj: TPanel; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function Panel_GetDragKind(AObj: TPanel): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure Panel_SetDragKind(AObj: TPanel; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function Panel_GetDragMode(AObj: TPanel): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure Panel_SetDragMode(AObj: TPanel; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function Panel_GetEnabled(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Panel_SetEnabled(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Panel_GetFullRepaint(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.FullRepaint;
end;

procedure Panel_SetFullRepaint(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.FullRepaint := AValue;
end;

function Panel_GetFont(AObj: TPanel): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure Panel_SetFont(AObj: TPanel; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function Panel_GetLocked(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.Locked;
end;

procedure Panel_SetLocked(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.Locked := AValue;
end;

function Panel_GetParentBackground(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentBackground;
end;

procedure Panel_SetParentBackground(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentBackground := AValue;
end;

function Panel_GetParentColor(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure Panel_SetParentColor(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function Panel_GetParentCtl3D(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure Panel_SetParentCtl3D(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function Panel_GetParentDoubleBuffered(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure Panel_SetParentDoubleBuffered(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function Panel_GetParentFont(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure Panel_SetParentFont(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function Panel_GetParentShowHint(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Panel_SetParentShowHint(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function Panel_GetPopupMenu(AObj: TPanel): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Panel_SetPopupMenu(AObj: TPanel; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function Panel_GetShowCaption(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.ShowCaption;
end;

procedure Panel_SetShowCaption(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.ShowCaption := AValue;
end;

function Panel_GetShowHint(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Panel_SetShowHint(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Panel_GetTabOrder(AObj: TPanel): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure Panel_SetTabOrder(AObj: TPanel; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function Panel_GetTabStop(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure Panel_SetTabStop(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function Panel_GetVisible(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Panel_SetVisible(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Panel_GetStyleElements(AObj: TPanel): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Panel_SetStyleElements(AObj: TPanel; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure Panel_SetOnClick(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnContextPopup(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnDockDrop(AObj: TPanel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDockDrop := nil;
    TEventClass.Remove(AObj, geDockDrop);
    Exit;
  end;
  AObj.OnDockDrop := TEventClass.OnDockDrop;
  TEventClass.Add(AObj, geDockDrop, AEventId);
end;

procedure Panel_SetOnDblClick(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnDragDrop(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnDragOver(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnEndDock(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnEndDrag(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnEnter(AObj: TPanel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEnter := nil;
    TEventClass.Remove(AObj, geEnter);
    Exit;
  end;
  AObj.OnEnter := TEventClass.OnEnter;
  TEventClass.Add(AObj, geEnter, AEventId);
end;

procedure Panel_SetOnExit(AObj: TPanel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnExit := nil;
    TEventClass.Remove(AObj, geExit);
    Exit;
  end;
  AObj.OnExit := TEventClass.OnExit;
  TEventClass.Add(AObj, geExit, AEventId);
end;

procedure Panel_SetOnGetSiteInfo(AObj: TPanel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetSiteInfo := nil;
    TEventClass.Remove(AObj, geGetSiteInfo);
    Exit;
  end;
  AObj.OnGetSiteInfo := TEventClass.OnGetSiteInfo;
  TEventClass.Add(AObj, geGetSiteInfo, AEventId);
end;

procedure Panel_SetOnMouseDown(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnMouseEnter(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnMouseLeave(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnMouseMove(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnMouseUp(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnResize(AObj: TPanel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnResize := nil;
    TEventClass.Remove(AObj, geResize);
    Exit;
  end;
  AObj.OnResize := TEventClass.OnResize;
  TEventClass.Add(AObj, geResize, AEventId);
end;

procedure Panel_SetOnStartDock(AObj: TPanel; AEventId: NativeUInt); stdcall;
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

procedure Panel_SetOnUnDock(AObj: TPanel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnUnDock := nil;
    TEventClass.Remove(AObj, geUnDock);
    Exit;
  end;
  AObj.OnUnDock := TEventClass.OnUnDock;
  TEventClass.Add(AObj, geUnDock, AEventId);
end;

function Panel_GetDockClientCount(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function Panel_GetAlignDisabled(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function Panel_GetMouseInClient(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function Panel_GetVisibleDockClientCount(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function Panel_GetBrush(AObj: TPanel): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function Panel_GetControlCount(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function Panel_GetHandle(AObj: TPanel): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function Panel_GetParentWindow(AObj: TPanel): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure Panel_SetParentWindow(AObj: TPanel; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function Panel_GetAction(AObj: TPanel): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Panel_SetAction(AObj: TPanel; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure Panel_GetBoundsRect(AObj: TPanel; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Panel_SetBoundsRect(AObj: TPanel; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Panel_GetClientHeight(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Panel_SetClientHeight(AObj: TPanel; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Panel_GetClientOrigin(AObj: TPanel; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure Panel_GetClientRect(AObj: TPanel; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Panel_GetClientWidth(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Panel_SetClientWidth(AObj: TPanel; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Panel_GetControlState(AObj: TPanel): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure Panel_SetControlState(AObj: TPanel; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function Panel_GetControlStyle(AObj: TPanel): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure Panel_SetControlStyle(AObj: TPanel; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function Panel_GetExplicitLeft(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Panel_GetExplicitTop(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Panel_GetExplicitWidth(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Panel_GetExplicitHeight(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Panel_GetFloating(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function Panel_GetParent(AObj: TPanel): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Panel_SetParent(AObj: TPanel; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Panel_GetAlignWithMargins(AObj: TPanel): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Panel_SetAlignWithMargins(AObj: TPanel; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Panel_GetLeft(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Panel_SetLeft(AObj: TPanel; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Panel_GetTop(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Panel_SetTop(AObj: TPanel; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Panel_GetWidth(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Panel_SetWidth(AObj: TPanel; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Panel_GetHeight(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Panel_SetHeight(AObj: TPanel; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Panel_GetCursor(AObj: TPanel): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Panel_SetCursor(AObj: TPanel; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Panel_GetHint(AObj: TPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Panel_SetHint(AObj: TPanel; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Panel_GetMargins(AObj: TPanel): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Panel_SetMargins(AObj: TPanel; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Panel_GetCustomHint(AObj: TPanel): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Panel_SetCustomHint(AObj: TPanel; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Panel_GetComponentCount(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Panel_GetComponentIndex(AObj: TPanel): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Panel_SetComponentIndex(AObj: TPanel; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Panel_GetOwner(AObj: TPanel): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Panel_GetName(AObj: TPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Panel_SetName(AObj: TPanel; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Panel_GetTag(AObj: TPanel): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Panel_SetTag(AObj: TPanel; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Panel_GetDockClients(AObj: TPanel; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function Panel_GetControls(AObj: TPanel; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function Panel_GetComponents(AObj: TPanel; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function Panel_StaticClassType: TClass; stdcall;
begin
  Result := TPanel;
end;

exports
  Panel_Create,
  Panel_Free,
  Panel_CanFocus,
  Panel_ContainsControl,
  Panel_ControlAtPos,
  Panel_DisableAlign,
  Panel_EnableAlign,
  Panel_FindChildControl,
  Panel_FlipChildren,
  Panel_Focused,
  Panel_HandleAllocated,
  Panel_InsertControl,
  Panel_Invalidate,
  Panel_PaintTo,
  Panel_RemoveControl,
  Panel_Realign,
  Panel_Repaint,
  Panel_ScaleBy,
  Panel_ScrollBy,
  Panel_SetBounds,
  Panel_SetFocus,
  Panel_Update,
  Panel_UpdateControlState,
  Panel_BringToFront,
  Panel_ClientToScreen,
  Panel_ClientToParent,
  Panel_Dragging,
  Panel_HasParent,
  Panel_Hide,
  Panel_Perform,
  Panel_Refresh,
  Panel_ScreenToClient,
  Panel_ParentToClient,
  Panel_SendToBack,
  Panel_Show,
  Panel_GetTextBuf,
  Panel_GetTextLen,
  Panel_SetTextBuf,
  Panel_FindComponent,
  Panel_GetNamePath,
  Panel_Assign,
  Panel_DisposeOf,
  Panel_ClassType,
  Panel_ClassName,
  Panel_InstanceSize,
  Panel_InheritsFrom,
  Panel_Equals,
  Panel_GetHashCode,
  Panel_ToString,
  Panel_GetAlign,
  Panel_SetAlign,
  Panel_GetAlignment,
  Panel_SetAlignment,
  Panel_GetAnchors,
  Panel_SetAnchors,
  Panel_GetAutoSize,
  Panel_SetAutoSize,
  Panel_GetBevelEdges,
  Panel_SetBevelEdges,
  Panel_GetBevelInner,
  Panel_SetBevelInner,
  Panel_GetBevelKind,
  Panel_SetBevelKind,
  Panel_GetBevelOuter,
  Panel_SetBevelOuter,
  Panel_GetBiDiMode,
  Panel_SetBiDiMode,
  Panel_GetBorderWidth,
  Panel_SetBorderWidth,
  Panel_GetBorderStyle,
  Panel_SetBorderStyle,
  Panel_GetCaption,
  Panel_SetCaption,
  Panel_GetColor,
  Panel_SetColor,
  Panel_GetUseDockManager,
  Panel_SetUseDockManager,
  Panel_GetDockSite,
  Panel_SetDockSite,
  Panel_GetDoubleBuffered,
  Panel_SetDoubleBuffered,
  Panel_GetDragCursor,
  Panel_SetDragCursor,
  Panel_GetDragKind,
  Panel_SetDragKind,
  Panel_GetDragMode,
  Panel_SetDragMode,
  Panel_GetEnabled,
  Panel_SetEnabled,
  Panel_GetFullRepaint,
  Panel_SetFullRepaint,
  Panel_GetFont,
  Panel_SetFont,
  Panel_GetLocked,
  Panel_SetLocked,
  Panel_GetParentBackground,
  Panel_SetParentBackground,
  Panel_GetParentColor,
  Panel_SetParentColor,
  Panel_GetParentCtl3D,
  Panel_SetParentCtl3D,
  Panel_GetParentDoubleBuffered,
  Panel_SetParentDoubleBuffered,
  Panel_GetParentFont,
  Panel_SetParentFont,
  Panel_GetParentShowHint,
  Panel_SetParentShowHint,
  Panel_GetPopupMenu,
  Panel_SetPopupMenu,
  Panel_GetShowCaption,
  Panel_SetShowCaption,
  Panel_GetShowHint,
  Panel_SetShowHint,
  Panel_GetTabOrder,
  Panel_SetTabOrder,
  Panel_GetTabStop,
  Panel_SetTabStop,
  Panel_GetVisible,
  Panel_SetVisible,
  Panel_GetStyleElements,
  Panel_SetStyleElements,
  Panel_SetOnClick,
  Panel_SetOnContextPopup,
  Panel_SetOnDockDrop,
  Panel_SetOnDblClick,
  Panel_SetOnDragDrop,
  Panel_SetOnDragOver,
  Panel_SetOnEndDock,
  Panel_SetOnEndDrag,
  Panel_SetOnEnter,
  Panel_SetOnExit,
  Panel_SetOnGetSiteInfo,
  Panel_SetOnMouseDown,
  Panel_SetOnMouseEnter,
  Panel_SetOnMouseLeave,
  Panel_SetOnMouseMove,
  Panel_SetOnMouseUp,
  Panel_SetOnResize,
  Panel_SetOnStartDock,
  Panel_SetOnUnDock,
  Panel_GetDockClientCount,
  Panel_GetAlignDisabled,
  Panel_GetMouseInClient,
  Panel_GetVisibleDockClientCount,
  Panel_GetBrush,
  Panel_GetControlCount,
  Panel_GetHandle,
  Panel_GetParentWindow,
  Panel_SetParentWindow,
  Panel_GetAction,
  Panel_SetAction,
  Panel_GetBoundsRect,
  Panel_SetBoundsRect,
  Panel_GetClientHeight,
  Panel_SetClientHeight,
  Panel_GetClientOrigin,
  Panel_GetClientRect,
  Panel_GetClientWidth,
  Panel_SetClientWidth,
  Panel_GetControlState,
  Panel_SetControlState,
  Panel_GetControlStyle,
  Panel_SetControlStyle,
  Panel_GetExplicitLeft,
  Panel_GetExplicitTop,
  Panel_GetExplicitWidth,
  Panel_GetExplicitHeight,
  Panel_GetFloating,
  Panel_GetParent,
  Panel_SetParent,
  Panel_GetAlignWithMargins,
  Panel_SetAlignWithMargins,
  Panel_GetLeft,
  Panel_SetLeft,
  Panel_GetTop,
  Panel_SetTop,
  Panel_GetWidth,
  Panel_SetWidth,
  Panel_GetHeight,
  Panel_SetHeight,
  Panel_GetCursor,
  Panel_SetCursor,
  Panel_GetHint,
  Panel_SetHint,
  Panel_GetMargins,
  Panel_SetMargins,
  Panel_GetCustomHint,
  Panel_SetCustomHint,
  Panel_GetComponentCount,
  Panel_GetComponentIndex,
  Panel_SetComponentIndex,
  Panel_GetOwner,
  Panel_GetName,
  Panel_SetName,
  Panel_GetTag,
  Panel_SetTag,
  Panel_GetDockClients,
  Panel_GetControls,
  Panel_GetComponents,
  Panel_StaticClassType;

