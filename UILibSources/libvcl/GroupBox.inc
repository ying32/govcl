
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function GroupBox_Create(AOwner: TComponent): TGroupBox; stdcall;
begin
  Result :=  TGroupBox.Create(AOwner);
end;

procedure GroupBox_Free(AObj: TGroupBox); stdcall;
begin
  AObj.Free;
end;

function GroupBox_CanFocus(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function GroupBox_ContainsControl(AObj: TGroupBox; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function GroupBox_ControlAtPos(AObj: TGroupBox; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure GroupBox_DisableAlign(AObj: TGroupBox); stdcall;
begin
  AObj.DisableAlign;
end;

procedure GroupBox_EnableAlign(AObj: TGroupBox); stdcall;
begin
  AObj.EnableAlign;
end;

function GroupBox_FindChildControl(AObj: TGroupBox; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure GroupBox_FlipChildren(AObj: TGroupBox; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function GroupBox_Focused(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function GroupBox_HandleAllocated(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure GroupBox_InsertControl(AObj: TGroupBox; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure GroupBox_Invalidate(AObj: TGroupBox); stdcall;
begin
  AObj.Invalidate;
end;

procedure GroupBox_PaintTo(AObj: TGroupBox; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure GroupBox_RemoveControl(AObj: TGroupBox; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure GroupBox_Realign(AObj: TGroupBox); stdcall;
begin
  AObj.Realign;
end;

procedure GroupBox_Repaint(AObj: TGroupBox); stdcall;
begin
  AObj.Repaint;
end;

procedure GroupBox_ScaleBy(AObj: TGroupBox; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure GroupBox_ScrollBy(AObj: TGroupBox; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure GroupBox_SetBounds(AObj: TGroupBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure GroupBox_SetFocus(AObj: TGroupBox); stdcall;
begin
  AObj.SetFocus;
end;

procedure GroupBox_Update(AObj: TGroupBox); stdcall;
begin
  AObj.Update;
end;

procedure GroupBox_UpdateControlState(AObj: TGroupBox); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure GroupBox_BringToFront(AObj: TGroupBox); stdcall;
begin
  AObj.BringToFront;
end;

procedure GroupBox_ClientToScreen(AObj: TGroupBox; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure GroupBox_ClientToParent(AObj: TGroupBox; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function GroupBox_Dragging(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function GroupBox_HasParent(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure GroupBox_Hide(AObj: TGroupBox); stdcall;
begin
  AObj.Hide;
end;

function GroupBox_Perform(AObj: TGroupBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure GroupBox_Refresh(AObj: TGroupBox); stdcall;
begin
  AObj.Refresh;
end;

procedure GroupBox_ScreenToClient(AObj: TGroupBox; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure GroupBox_ParentToClient(AObj: TGroupBox; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure GroupBox_SendToBack(AObj: TGroupBox); stdcall;
begin
  AObj.SendToBack;
end;

procedure GroupBox_Show(AObj: TGroupBox); stdcall;
begin
  AObj.Show;
end;

function GroupBox_GetTextBuf(AObj: TGroupBox; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function GroupBox_GetTextLen(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure GroupBox_SetTextBuf(AObj: TGroupBox; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function GroupBox_FindComponent(AObj: TGroupBox; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function GroupBox_GetNamePath(AObj: TGroupBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure GroupBox_Assign(AObj: TGroupBox; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure GroupBox_DisposeOf(AObj: TGroupBox); stdcall;
begin
  AObj.DisposeOf;
end;

function GroupBox_ClassType(AObj: TGroupBox): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function GroupBox_ClassName(AObj: TGroupBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function GroupBox_InstanceSize(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function GroupBox_InheritsFrom(AObj: TGroupBox; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function GroupBox_Equals(AObj: TGroupBox; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function GroupBox_GetHashCode(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function GroupBox_ToString(AObj: TGroupBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function GroupBox_GetAlign(AObj: TGroupBox): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure GroupBox_SetAlign(AObj: TGroupBox; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function GroupBox_GetAnchors(AObj: TGroupBox): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure GroupBox_SetAnchors(AObj: TGroupBox; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function GroupBox_GetBiDiMode(AObj: TGroupBox): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure GroupBox_SetBiDiMode(AObj: TGroupBox; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function GroupBox_GetCaption(AObj: TGroupBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure GroupBox_SetCaption(AObj: TGroupBox; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function GroupBox_GetColor(AObj: TGroupBox): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure GroupBox_SetColor(AObj: TGroupBox; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function GroupBox_GetDockSite(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure GroupBox_SetDockSite(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function GroupBox_GetDoubleBuffered(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure GroupBox_SetDoubleBuffered(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function GroupBox_GetDragCursor(AObj: TGroupBox): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure GroupBox_SetDragCursor(AObj: TGroupBox; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function GroupBox_GetDragKind(AObj: TGroupBox): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure GroupBox_SetDragKind(AObj: TGroupBox; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function GroupBox_GetDragMode(AObj: TGroupBox): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure GroupBox_SetDragMode(AObj: TGroupBox; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function GroupBox_GetEnabled(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure GroupBox_SetEnabled(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function GroupBox_GetFont(AObj: TGroupBox): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure GroupBox_SetFont(AObj: TGroupBox; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function GroupBox_GetParentBackground(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentBackground;
end;

procedure GroupBox_SetParentBackground(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.ParentBackground := AValue;
end;

function GroupBox_GetParentColor(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure GroupBox_SetParentColor(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function GroupBox_GetParentCtl3D(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure GroupBox_SetParentCtl3D(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function GroupBox_GetParentDoubleBuffered(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure GroupBox_SetParentDoubleBuffered(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function GroupBox_GetParentFont(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure GroupBox_SetParentFont(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function GroupBox_GetParentShowHint(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure GroupBox_SetParentShowHint(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function GroupBox_GetPopupMenu(AObj: TGroupBox): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure GroupBox_SetPopupMenu(AObj: TGroupBox; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function GroupBox_GetShowHint(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure GroupBox_SetShowHint(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function GroupBox_GetTabOrder(AObj: TGroupBox): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure GroupBox_SetTabOrder(AObj: TGroupBox; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function GroupBox_GetTabStop(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure GroupBox_SetTabStop(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function GroupBox_GetVisible(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure GroupBox_SetVisible(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function GroupBox_GetStyleElements(AObj: TGroupBox): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure GroupBox_SetStyleElements(AObj: TGroupBox; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure GroupBox_SetOnClick(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnContextPopup(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnDblClick(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnDragDrop(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnDockDrop(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnDragOver(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnEndDock(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnEndDrag(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnEnter(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnExit(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnGesture(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnGetSiteInfo(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnMouseActivate(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnMouseDown(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnMouseEnter(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnMouseLeave(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnMouseMove(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnMouseUp(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnStartDock(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

procedure GroupBox_SetOnUnDock(AObj: TGroupBox; AEventId: NativeUInt); stdcall;
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

function GroupBox_GetDockClientCount(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function GroupBox_GetAlignDisabled(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function GroupBox_GetMouseInClient(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function GroupBox_GetVisibleDockClientCount(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function GroupBox_GetBrush(AObj: TGroupBox): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function GroupBox_GetControlCount(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function GroupBox_GetHandle(AObj: TGroupBox): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function GroupBox_GetParentWindow(AObj: TGroupBox): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure GroupBox_SetParentWindow(AObj: TGroupBox; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function GroupBox_GetUseDockManager(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure GroupBox_SetUseDockManager(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function GroupBox_GetAction(AObj: TGroupBox): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure GroupBox_SetAction(AObj: TGroupBox; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure GroupBox_GetBoundsRect(AObj: TGroupBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure GroupBox_SetBoundsRect(AObj: TGroupBox; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function GroupBox_GetClientHeight(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure GroupBox_SetClientHeight(AObj: TGroupBox; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure GroupBox_GetClientOrigin(AObj: TGroupBox; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure GroupBox_GetClientRect(AObj: TGroupBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function GroupBox_GetClientWidth(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure GroupBox_SetClientWidth(AObj: TGroupBox; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function GroupBox_GetControlState(AObj: TGroupBox): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure GroupBox_SetControlState(AObj: TGroupBox; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function GroupBox_GetControlStyle(AObj: TGroupBox): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure GroupBox_SetControlStyle(AObj: TGroupBox; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function GroupBox_GetExplicitLeft(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function GroupBox_GetExplicitTop(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function GroupBox_GetExplicitWidth(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function GroupBox_GetExplicitHeight(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function GroupBox_GetFloating(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function GroupBox_GetParent(AObj: TGroupBox): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure GroupBox_SetParent(AObj: TGroupBox; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function GroupBox_GetAlignWithMargins(AObj: TGroupBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure GroupBox_SetAlignWithMargins(AObj: TGroupBox; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function GroupBox_GetLeft(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure GroupBox_SetLeft(AObj: TGroupBox; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function GroupBox_GetTop(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure GroupBox_SetTop(AObj: TGroupBox; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function GroupBox_GetWidth(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure GroupBox_SetWidth(AObj: TGroupBox; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function GroupBox_GetHeight(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure GroupBox_SetHeight(AObj: TGroupBox; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function GroupBox_GetCursor(AObj: TGroupBox): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure GroupBox_SetCursor(AObj: TGroupBox; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function GroupBox_GetHint(AObj: TGroupBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure GroupBox_SetHint(AObj: TGroupBox; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function GroupBox_GetMargins(AObj: TGroupBox): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure GroupBox_SetMargins(AObj: TGroupBox; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function GroupBox_GetCustomHint(AObj: TGroupBox): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure GroupBox_SetCustomHint(AObj: TGroupBox; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function GroupBox_GetComponentCount(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function GroupBox_GetComponentIndex(AObj: TGroupBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure GroupBox_SetComponentIndex(AObj: TGroupBox; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function GroupBox_GetOwner(AObj: TGroupBox): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function GroupBox_GetName(AObj: TGroupBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure GroupBox_SetName(AObj: TGroupBox; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function GroupBox_GetTag(AObj: TGroupBox): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure GroupBox_SetTag(AObj: TGroupBox; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function GroupBox_GetDockClients(AObj: TGroupBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function GroupBox_GetControls(AObj: TGroupBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function GroupBox_GetComponents(AObj: TGroupBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function GroupBox_StaticClassType: TClass; stdcall;
begin
  Result := TGroupBox;
end;

exports
  GroupBox_Create,
  GroupBox_Free,
  GroupBox_CanFocus,
  GroupBox_ContainsControl,
  GroupBox_ControlAtPos,
  GroupBox_DisableAlign,
  GroupBox_EnableAlign,
  GroupBox_FindChildControl,
  GroupBox_FlipChildren,
  GroupBox_Focused,
  GroupBox_HandleAllocated,
  GroupBox_InsertControl,
  GroupBox_Invalidate,
  GroupBox_PaintTo,
  GroupBox_RemoveControl,
  GroupBox_Realign,
  GroupBox_Repaint,
  GroupBox_ScaleBy,
  GroupBox_ScrollBy,
  GroupBox_SetBounds,
  GroupBox_SetFocus,
  GroupBox_Update,
  GroupBox_UpdateControlState,
  GroupBox_BringToFront,
  GroupBox_ClientToScreen,
  GroupBox_ClientToParent,
  GroupBox_Dragging,
  GroupBox_HasParent,
  GroupBox_Hide,
  GroupBox_Perform,
  GroupBox_Refresh,
  GroupBox_ScreenToClient,
  GroupBox_ParentToClient,
  GroupBox_SendToBack,
  GroupBox_Show,
  GroupBox_GetTextBuf,
  GroupBox_GetTextLen,
  GroupBox_SetTextBuf,
  GroupBox_FindComponent,
  GroupBox_GetNamePath,
  GroupBox_Assign,
  GroupBox_DisposeOf,
  GroupBox_ClassType,
  GroupBox_ClassName,
  GroupBox_InstanceSize,
  GroupBox_InheritsFrom,
  GroupBox_Equals,
  GroupBox_GetHashCode,
  GroupBox_ToString,
  GroupBox_GetAlign,
  GroupBox_SetAlign,
  GroupBox_GetAnchors,
  GroupBox_SetAnchors,
  GroupBox_GetBiDiMode,
  GroupBox_SetBiDiMode,
  GroupBox_GetCaption,
  GroupBox_SetCaption,
  GroupBox_GetColor,
  GroupBox_SetColor,
  GroupBox_GetDockSite,
  GroupBox_SetDockSite,
  GroupBox_GetDoubleBuffered,
  GroupBox_SetDoubleBuffered,
  GroupBox_GetDragCursor,
  GroupBox_SetDragCursor,
  GroupBox_GetDragKind,
  GroupBox_SetDragKind,
  GroupBox_GetDragMode,
  GroupBox_SetDragMode,
  GroupBox_GetEnabled,
  GroupBox_SetEnabled,
  GroupBox_GetFont,
  GroupBox_SetFont,
  GroupBox_GetParentBackground,
  GroupBox_SetParentBackground,
  GroupBox_GetParentColor,
  GroupBox_SetParentColor,
  GroupBox_GetParentCtl3D,
  GroupBox_SetParentCtl3D,
  GroupBox_GetParentDoubleBuffered,
  GroupBox_SetParentDoubleBuffered,
  GroupBox_GetParentFont,
  GroupBox_SetParentFont,
  GroupBox_GetParentShowHint,
  GroupBox_SetParentShowHint,
  GroupBox_GetPopupMenu,
  GroupBox_SetPopupMenu,
  GroupBox_GetShowHint,
  GroupBox_SetShowHint,
  GroupBox_GetTabOrder,
  GroupBox_SetTabOrder,
  GroupBox_GetTabStop,
  GroupBox_SetTabStop,
  GroupBox_GetVisible,
  GroupBox_SetVisible,
  GroupBox_GetStyleElements,
  GroupBox_SetStyleElements,
  GroupBox_SetOnClick,
  GroupBox_SetOnContextPopup,
  GroupBox_SetOnDblClick,
  GroupBox_SetOnDragDrop,
  GroupBox_SetOnDockDrop,
  GroupBox_SetOnDragOver,
  GroupBox_SetOnEndDock,
  GroupBox_SetOnEndDrag,
  GroupBox_SetOnEnter,
  GroupBox_SetOnExit,
  GroupBox_SetOnGesture,
  GroupBox_SetOnGetSiteInfo,
  GroupBox_SetOnMouseActivate,
  GroupBox_SetOnMouseDown,
  GroupBox_SetOnMouseEnter,
  GroupBox_SetOnMouseLeave,
  GroupBox_SetOnMouseMove,
  GroupBox_SetOnMouseUp,
  GroupBox_SetOnStartDock,
  GroupBox_SetOnUnDock,
  GroupBox_GetDockClientCount,
  GroupBox_GetAlignDisabled,
  GroupBox_GetMouseInClient,
  GroupBox_GetVisibleDockClientCount,
  GroupBox_GetBrush,
  GroupBox_GetControlCount,
  GroupBox_GetHandle,
  GroupBox_GetParentWindow,
  GroupBox_SetParentWindow,
  GroupBox_GetUseDockManager,
  GroupBox_SetUseDockManager,
  GroupBox_GetAction,
  GroupBox_SetAction,
  GroupBox_GetBoundsRect,
  GroupBox_SetBoundsRect,
  GroupBox_GetClientHeight,
  GroupBox_SetClientHeight,
  GroupBox_GetClientOrigin,
  GroupBox_GetClientRect,
  GroupBox_GetClientWidth,
  GroupBox_SetClientWidth,
  GroupBox_GetControlState,
  GroupBox_SetControlState,
  GroupBox_GetControlStyle,
  GroupBox_SetControlStyle,
  GroupBox_GetExplicitLeft,
  GroupBox_GetExplicitTop,
  GroupBox_GetExplicitWidth,
  GroupBox_GetExplicitHeight,
  GroupBox_GetFloating,
  GroupBox_GetParent,
  GroupBox_SetParent,
  GroupBox_GetAlignWithMargins,
  GroupBox_SetAlignWithMargins,
  GroupBox_GetLeft,
  GroupBox_SetLeft,
  GroupBox_GetTop,
  GroupBox_SetTop,
  GroupBox_GetWidth,
  GroupBox_SetWidth,
  GroupBox_GetHeight,
  GroupBox_SetHeight,
  GroupBox_GetCursor,
  GroupBox_SetCursor,
  GroupBox_GetHint,
  GroupBox_SetHint,
  GroupBox_GetMargins,
  GroupBox_SetMargins,
  GroupBox_GetCustomHint,
  GroupBox_SetCustomHint,
  GroupBox_GetComponentCount,
  GroupBox_GetComponentIndex,
  GroupBox_SetComponentIndex,
  GroupBox_GetOwner,
  GroupBox_GetName,
  GroupBox_SetName,
  GroupBox_GetTag,
  GroupBox_SetTag,
  GroupBox_GetDockClients,
  GroupBox_GetControls,
  GroupBox_GetComponents,
  GroupBox_StaticClassType;

