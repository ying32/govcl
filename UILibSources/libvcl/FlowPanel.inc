
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function FlowPanel_Create(AOwner: TComponent): TFlowPanel; stdcall;
begin
  Result :=  TFlowPanel.Create(AOwner);
end;

procedure FlowPanel_Free(AObj: TFlowPanel); stdcall;
begin
  AObj.Free;
end;

function FlowPanel_GetControlIndex(AObj: TFlowPanel; AControl: TControl): Integer; stdcall;
begin
  Result :=  AObj.GetControlIndex(AControl);
end;

procedure FlowPanel_SetControlIndex(AObj: TFlowPanel; AControl: TControl; Index: Integer); stdcall;
begin
  AObj.SetControlIndex(AControl, Index);
end;

function FlowPanel_CanFocus(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function FlowPanel_ContainsControl(AObj: TFlowPanel; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function FlowPanel_ControlAtPos(AObj: TFlowPanel; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure FlowPanel_DisableAlign(AObj: TFlowPanel); stdcall;
begin
  AObj.DisableAlign;
end;

procedure FlowPanel_EnableAlign(AObj: TFlowPanel); stdcall;
begin
  AObj.EnableAlign;
end;

function FlowPanel_FindChildControl(AObj: TFlowPanel; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure FlowPanel_FlipChildren(AObj: TFlowPanel; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function FlowPanel_Focused(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function FlowPanel_HandleAllocated(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure FlowPanel_InsertControl(AObj: TFlowPanel; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure FlowPanel_Invalidate(AObj: TFlowPanel); stdcall;
begin
  AObj.Invalidate;
end;

procedure FlowPanel_PaintTo(AObj: TFlowPanel; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure FlowPanel_RemoveControl(AObj: TFlowPanel; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure FlowPanel_Realign(AObj: TFlowPanel); stdcall;
begin
  AObj.Realign;
end;

procedure FlowPanel_Repaint(AObj: TFlowPanel); stdcall;
begin
  AObj.Repaint;
end;

procedure FlowPanel_ScaleBy(AObj: TFlowPanel; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure FlowPanel_ScrollBy(AObj: TFlowPanel; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure FlowPanel_SetBounds(AObj: TFlowPanel; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure FlowPanel_SetFocus(AObj: TFlowPanel); stdcall;
begin
  AObj.SetFocus;
end;

procedure FlowPanel_Update(AObj: TFlowPanel); stdcall;
begin
  AObj.Update;
end;

procedure FlowPanel_UpdateControlState(AObj: TFlowPanel); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure FlowPanel_BringToFront(AObj: TFlowPanel); stdcall;
begin
  AObj.BringToFront;
end;

procedure FlowPanel_ClientToScreen(AObj: TFlowPanel; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure FlowPanel_ClientToParent(AObj: TFlowPanel; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function FlowPanel_Dragging(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function FlowPanel_HasParent(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure FlowPanel_Hide(AObj: TFlowPanel); stdcall;
begin
  AObj.Hide;
end;

function FlowPanel_Perform(AObj: TFlowPanel; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure FlowPanel_Refresh(AObj: TFlowPanel); stdcall;
begin
  AObj.Refresh;
end;

procedure FlowPanel_ScreenToClient(AObj: TFlowPanel; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure FlowPanel_ParentToClient(AObj: TFlowPanel; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure FlowPanel_SendToBack(AObj: TFlowPanel); stdcall;
begin
  AObj.SendToBack;
end;

procedure FlowPanel_Show(AObj: TFlowPanel); stdcall;
begin
  AObj.Show;
end;

function FlowPanel_GetTextBuf(AObj: TFlowPanel; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function FlowPanel_GetTextLen(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure FlowPanel_SetTextBuf(AObj: TFlowPanel; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function FlowPanel_FindComponent(AObj: TFlowPanel; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function FlowPanel_GetNamePath(AObj: TFlowPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure FlowPanel_Assign(AObj: TFlowPanel; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure FlowPanel_DisposeOf(AObj: TFlowPanel); stdcall;
begin
  AObj.DisposeOf;
end;

function FlowPanel_ClassType(AObj: TFlowPanel): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function FlowPanel_ClassName(AObj: TFlowPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function FlowPanel_InstanceSize(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function FlowPanel_InheritsFrom(AObj: TFlowPanel; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function FlowPanel_Equals(AObj: TFlowPanel; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function FlowPanel_GetHashCode(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function FlowPanel_ToString(AObj: TFlowPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function FlowPanel_GetAlign(AObj: TFlowPanel): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure FlowPanel_SetAlign(AObj: TFlowPanel; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function FlowPanel_GetAlignment(AObj: TFlowPanel): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure FlowPanel_SetAlignment(AObj: TFlowPanel; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function FlowPanel_GetAnchors(AObj: TFlowPanel): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure FlowPanel_SetAnchors(AObj: TFlowPanel; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function FlowPanel_GetAutoSize(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure FlowPanel_SetAutoSize(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function FlowPanel_GetAutoWrap(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.AutoWrap;
end;

procedure FlowPanel_SetAutoWrap(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.AutoWrap := AValue;
end;

function FlowPanel_GetBevelEdges(AObj: TFlowPanel): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure FlowPanel_SetBevelEdges(AObj: TFlowPanel; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function FlowPanel_GetBevelInner(AObj: TFlowPanel): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure FlowPanel_SetBevelInner(AObj: TFlowPanel; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function FlowPanel_GetBevelKind(AObj: TFlowPanel): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure FlowPanel_SetBevelKind(AObj: TFlowPanel; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function FlowPanel_GetBevelOuter(AObj: TFlowPanel): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure FlowPanel_SetBevelOuter(AObj: TFlowPanel; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function FlowPanel_GetBiDiMode(AObj: TFlowPanel): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure FlowPanel_SetBiDiMode(AObj: TFlowPanel; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function FlowPanel_GetBorderWidth(AObj: TFlowPanel): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure FlowPanel_SetBorderWidth(AObj: TFlowPanel; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function FlowPanel_GetBorderStyle(AObj: TFlowPanel): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure FlowPanel_SetBorderStyle(AObj: TFlowPanel; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function FlowPanel_GetCaption(AObj: TFlowPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure FlowPanel_SetCaption(AObj: TFlowPanel; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function FlowPanel_GetColor(AObj: TFlowPanel): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure FlowPanel_SetColor(AObj: TFlowPanel; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function FlowPanel_GetConstraints(AObj: TFlowPanel): TSizeConstraints; stdcall;
begin
  Result :=  AObj.Constraints;
end;

procedure FlowPanel_SetConstraints(AObj: TFlowPanel; AValue: TSizeConstraints); stdcall;
begin
  AObj.Constraints := AValue;
end;

function FlowPanel_GetCtl3D(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.Ctl3D;
end;

procedure FlowPanel_SetCtl3D(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.Ctl3D := AValue;
end;

function FlowPanel_GetUseDockManager(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure FlowPanel_SetUseDockManager(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function FlowPanel_GetDockSite(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure FlowPanel_SetDockSite(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function FlowPanel_GetDoubleBuffered(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure FlowPanel_SetDoubleBuffered(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function FlowPanel_GetDragCursor(AObj: TFlowPanel): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure FlowPanel_SetDragCursor(AObj: TFlowPanel; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function FlowPanel_GetDragKind(AObj: TFlowPanel): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure FlowPanel_SetDragKind(AObj: TFlowPanel; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function FlowPanel_GetDragMode(AObj: TFlowPanel): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure FlowPanel_SetDragMode(AObj: TFlowPanel; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function FlowPanel_GetEnabled(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure FlowPanel_SetEnabled(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function FlowPanel_GetFlowStyle(AObj: TFlowPanel): TFlowStyle; stdcall;
begin
  Result :=  AObj.FlowStyle;
end;

procedure FlowPanel_SetFlowStyle(AObj: TFlowPanel; AValue: TFlowStyle); stdcall;
begin
  AObj.FlowStyle := AValue;
end;

function FlowPanel_GetFullRepaint(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.FullRepaint;
end;

procedure FlowPanel_SetFullRepaint(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.FullRepaint := AValue;
end;

function FlowPanel_GetFont(AObj: TFlowPanel): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure FlowPanel_SetFont(AObj: TFlowPanel; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function FlowPanel_GetLocked(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.Locked;
end;

procedure FlowPanel_SetLocked(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.Locked := AValue;
end;

function FlowPanel_GetParentBackground(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentBackground;
end;

procedure FlowPanel_SetParentBackground(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentBackground := AValue;
end;

function FlowPanel_GetParentColor(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure FlowPanel_SetParentColor(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function FlowPanel_GetParentCtl3D(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure FlowPanel_SetParentCtl3D(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function FlowPanel_GetParentDoubleBuffered(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure FlowPanel_SetParentDoubleBuffered(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function FlowPanel_GetParentFont(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure FlowPanel_SetParentFont(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function FlowPanel_GetParentShowHint(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure FlowPanel_SetParentShowHint(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function FlowPanel_GetPopupMenu(AObj: TFlowPanel): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure FlowPanel_SetPopupMenu(AObj: TFlowPanel; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function FlowPanel_GetShowCaption(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.ShowCaption;
end;

procedure FlowPanel_SetShowCaption(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.ShowCaption := AValue;
end;

function FlowPanel_GetShowHint(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure FlowPanel_SetShowHint(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function FlowPanel_GetTabOrder(AObj: TFlowPanel): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure FlowPanel_SetTabOrder(AObj: TFlowPanel; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function FlowPanel_GetTabStop(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure FlowPanel_SetTabStop(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function FlowPanel_GetVisible(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure FlowPanel_SetVisible(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function FlowPanel_GetStyleElements(AObj: TFlowPanel): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure FlowPanel_SetStyleElements(AObj: TFlowPanel; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure FlowPanel_SetOnClick(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnContextPopup(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnDockDrop(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnDblClick(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnDragDrop(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnDragOver(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnEndDock(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnEndDrag(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnEnter(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnExit(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnGesture(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnGetSiteInfo(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnMouseActivate(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnMouseDown(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnMouseEnter(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnMouseLeave(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnMouseMove(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnMouseUp(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnResize(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnStartDock(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

procedure FlowPanel_SetOnUnDock(AObj: TFlowPanel; AEventId: NativeUInt); stdcall;
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

function FlowPanel_GetDockClientCount(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function FlowPanel_GetAlignDisabled(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function FlowPanel_GetMouseInClient(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function FlowPanel_GetVisibleDockClientCount(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function FlowPanel_GetBrush(AObj: TFlowPanel): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function FlowPanel_GetControlCount(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function FlowPanel_GetHandle(AObj: TFlowPanel): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function FlowPanel_GetParentWindow(AObj: TFlowPanel): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure FlowPanel_SetParentWindow(AObj: TFlowPanel; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function FlowPanel_GetAction(AObj: TFlowPanel): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure FlowPanel_SetAction(AObj: TFlowPanel; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure FlowPanel_GetBoundsRect(AObj: TFlowPanel; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure FlowPanel_SetBoundsRect(AObj: TFlowPanel; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function FlowPanel_GetClientHeight(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure FlowPanel_SetClientHeight(AObj: TFlowPanel; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure FlowPanel_GetClientOrigin(AObj: TFlowPanel; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure FlowPanel_GetClientRect(AObj: TFlowPanel; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function FlowPanel_GetClientWidth(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure FlowPanel_SetClientWidth(AObj: TFlowPanel; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function FlowPanel_GetControlState(AObj: TFlowPanel): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure FlowPanel_SetControlState(AObj: TFlowPanel; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function FlowPanel_GetControlStyle(AObj: TFlowPanel): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure FlowPanel_SetControlStyle(AObj: TFlowPanel; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function FlowPanel_GetExplicitLeft(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function FlowPanel_GetExplicitTop(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function FlowPanel_GetExplicitWidth(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function FlowPanel_GetExplicitHeight(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function FlowPanel_GetFloating(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function FlowPanel_GetParent(AObj: TFlowPanel): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure FlowPanel_SetParent(AObj: TFlowPanel; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function FlowPanel_GetAlignWithMargins(AObj: TFlowPanel): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure FlowPanel_SetAlignWithMargins(AObj: TFlowPanel; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function FlowPanel_GetLeft(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure FlowPanel_SetLeft(AObj: TFlowPanel; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function FlowPanel_GetTop(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure FlowPanel_SetTop(AObj: TFlowPanel; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function FlowPanel_GetWidth(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure FlowPanel_SetWidth(AObj: TFlowPanel; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function FlowPanel_GetHeight(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure FlowPanel_SetHeight(AObj: TFlowPanel; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function FlowPanel_GetCursor(AObj: TFlowPanel): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure FlowPanel_SetCursor(AObj: TFlowPanel; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function FlowPanel_GetHint(AObj: TFlowPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure FlowPanel_SetHint(AObj: TFlowPanel; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function FlowPanel_GetMargins(AObj: TFlowPanel): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure FlowPanel_SetMargins(AObj: TFlowPanel; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function FlowPanel_GetCustomHint(AObj: TFlowPanel): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure FlowPanel_SetCustomHint(AObj: TFlowPanel; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function FlowPanel_GetComponentCount(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function FlowPanel_GetComponentIndex(AObj: TFlowPanel): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure FlowPanel_SetComponentIndex(AObj: TFlowPanel; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function FlowPanel_GetOwner(AObj: TFlowPanel): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function FlowPanel_GetName(AObj: TFlowPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure FlowPanel_SetName(AObj: TFlowPanel; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function FlowPanel_GetTag(AObj: TFlowPanel): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure FlowPanel_SetTag(AObj: TFlowPanel; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function FlowPanel_GetDockClients(AObj: TFlowPanel; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function FlowPanel_GetControls(AObj: TFlowPanel; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function FlowPanel_GetComponents(AObj: TFlowPanel; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function FlowPanel_StaticClassType: TClass; stdcall;
begin
  Result := TFlowPanel;
end;

exports
  FlowPanel_Create,
  FlowPanel_Free,
  FlowPanel_GetControlIndex,
  FlowPanel_SetControlIndex,
  FlowPanel_CanFocus,
  FlowPanel_ContainsControl,
  FlowPanel_ControlAtPos,
  FlowPanel_DisableAlign,
  FlowPanel_EnableAlign,
  FlowPanel_FindChildControl,
  FlowPanel_FlipChildren,
  FlowPanel_Focused,
  FlowPanel_HandleAllocated,
  FlowPanel_InsertControl,
  FlowPanel_Invalidate,
  FlowPanel_PaintTo,
  FlowPanel_RemoveControl,
  FlowPanel_Realign,
  FlowPanel_Repaint,
  FlowPanel_ScaleBy,
  FlowPanel_ScrollBy,
  FlowPanel_SetBounds,
  FlowPanel_SetFocus,
  FlowPanel_Update,
  FlowPanel_UpdateControlState,
  FlowPanel_BringToFront,
  FlowPanel_ClientToScreen,
  FlowPanel_ClientToParent,
  FlowPanel_Dragging,
  FlowPanel_HasParent,
  FlowPanel_Hide,
  FlowPanel_Perform,
  FlowPanel_Refresh,
  FlowPanel_ScreenToClient,
  FlowPanel_ParentToClient,
  FlowPanel_SendToBack,
  FlowPanel_Show,
  FlowPanel_GetTextBuf,
  FlowPanel_GetTextLen,
  FlowPanel_SetTextBuf,
  FlowPanel_FindComponent,
  FlowPanel_GetNamePath,
  FlowPanel_Assign,
  FlowPanel_DisposeOf,
  FlowPanel_ClassType,
  FlowPanel_ClassName,
  FlowPanel_InstanceSize,
  FlowPanel_InheritsFrom,
  FlowPanel_Equals,
  FlowPanel_GetHashCode,
  FlowPanel_ToString,
  FlowPanel_GetAlign,
  FlowPanel_SetAlign,
  FlowPanel_GetAlignment,
  FlowPanel_SetAlignment,
  FlowPanel_GetAnchors,
  FlowPanel_SetAnchors,
  FlowPanel_GetAutoSize,
  FlowPanel_SetAutoSize,
  FlowPanel_GetAutoWrap,
  FlowPanel_SetAutoWrap,
  FlowPanel_GetBevelEdges,
  FlowPanel_SetBevelEdges,
  FlowPanel_GetBevelInner,
  FlowPanel_SetBevelInner,
  FlowPanel_GetBevelKind,
  FlowPanel_SetBevelKind,
  FlowPanel_GetBevelOuter,
  FlowPanel_SetBevelOuter,
  FlowPanel_GetBiDiMode,
  FlowPanel_SetBiDiMode,
  FlowPanel_GetBorderWidth,
  FlowPanel_SetBorderWidth,
  FlowPanel_GetBorderStyle,
  FlowPanel_SetBorderStyle,
  FlowPanel_GetCaption,
  FlowPanel_SetCaption,
  FlowPanel_GetColor,
  FlowPanel_SetColor,
  FlowPanel_GetConstraints,
  FlowPanel_SetConstraints,
  FlowPanel_GetCtl3D,
  FlowPanel_SetCtl3D,
  FlowPanel_GetUseDockManager,
  FlowPanel_SetUseDockManager,
  FlowPanel_GetDockSite,
  FlowPanel_SetDockSite,
  FlowPanel_GetDoubleBuffered,
  FlowPanel_SetDoubleBuffered,
  FlowPanel_GetDragCursor,
  FlowPanel_SetDragCursor,
  FlowPanel_GetDragKind,
  FlowPanel_SetDragKind,
  FlowPanel_GetDragMode,
  FlowPanel_SetDragMode,
  FlowPanel_GetEnabled,
  FlowPanel_SetEnabled,
  FlowPanel_GetFlowStyle,
  FlowPanel_SetFlowStyle,
  FlowPanel_GetFullRepaint,
  FlowPanel_SetFullRepaint,
  FlowPanel_GetFont,
  FlowPanel_SetFont,
  FlowPanel_GetLocked,
  FlowPanel_SetLocked,
  FlowPanel_GetParentBackground,
  FlowPanel_SetParentBackground,
  FlowPanel_GetParentColor,
  FlowPanel_SetParentColor,
  FlowPanel_GetParentCtl3D,
  FlowPanel_SetParentCtl3D,
  FlowPanel_GetParentDoubleBuffered,
  FlowPanel_SetParentDoubleBuffered,
  FlowPanel_GetParentFont,
  FlowPanel_SetParentFont,
  FlowPanel_GetParentShowHint,
  FlowPanel_SetParentShowHint,
  FlowPanel_GetPopupMenu,
  FlowPanel_SetPopupMenu,
  FlowPanel_GetShowCaption,
  FlowPanel_SetShowCaption,
  FlowPanel_GetShowHint,
  FlowPanel_SetShowHint,
  FlowPanel_GetTabOrder,
  FlowPanel_SetTabOrder,
  FlowPanel_GetTabStop,
  FlowPanel_SetTabStop,
  FlowPanel_GetVisible,
  FlowPanel_SetVisible,
  FlowPanel_GetStyleElements,
  FlowPanel_SetStyleElements,
  FlowPanel_SetOnClick,
  FlowPanel_SetOnContextPopup,
  FlowPanel_SetOnDockDrop,
  FlowPanel_SetOnDblClick,
  FlowPanel_SetOnDragDrop,
  FlowPanel_SetOnDragOver,
  FlowPanel_SetOnEndDock,
  FlowPanel_SetOnEndDrag,
  FlowPanel_SetOnEnter,
  FlowPanel_SetOnExit,
  FlowPanel_SetOnGesture,
  FlowPanel_SetOnGetSiteInfo,
  FlowPanel_SetOnMouseActivate,
  FlowPanel_SetOnMouseDown,
  FlowPanel_SetOnMouseEnter,
  FlowPanel_SetOnMouseLeave,
  FlowPanel_SetOnMouseMove,
  FlowPanel_SetOnMouseUp,
  FlowPanel_SetOnResize,
  FlowPanel_SetOnStartDock,
  FlowPanel_SetOnUnDock,
  FlowPanel_GetDockClientCount,
  FlowPanel_GetAlignDisabled,
  FlowPanel_GetMouseInClient,
  FlowPanel_GetVisibleDockClientCount,
  FlowPanel_GetBrush,
  FlowPanel_GetControlCount,
  FlowPanel_GetHandle,
  FlowPanel_GetParentWindow,
  FlowPanel_SetParentWindow,
  FlowPanel_GetAction,
  FlowPanel_SetAction,
  FlowPanel_GetBoundsRect,
  FlowPanel_SetBoundsRect,
  FlowPanel_GetClientHeight,
  FlowPanel_SetClientHeight,
  FlowPanel_GetClientOrigin,
  FlowPanel_GetClientRect,
  FlowPanel_GetClientWidth,
  FlowPanel_SetClientWidth,
  FlowPanel_GetControlState,
  FlowPanel_SetControlState,
  FlowPanel_GetControlStyle,
  FlowPanel_SetControlStyle,
  FlowPanel_GetExplicitLeft,
  FlowPanel_GetExplicitTop,
  FlowPanel_GetExplicitWidth,
  FlowPanel_GetExplicitHeight,
  FlowPanel_GetFloating,
  FlowPanel_GetParent,
  FlowPanel_SetParent,
  FlowPanel_GetAlignWithMargins,
  FlowPanel_SetAlignWithMargins,
  FlowPanel_GetLeft,
  FlowPanel_SetLeft,
  FlowPanel_GetTop,
  FlowPanel_SetTop,
  FlowPanel_GetWidth,
  FlowPanel_SetWidth,
  FlowPanel_GetHeight,
  FlowPanel_SetHeight,
  FlowPanel_GetCursor,
  FlowPanel_SetCursor,
  FlowPanel_GetHint,
  FlowPanel_SetHint,
  FlowPanel_GetMargins,
  FlowPanel_SetMargins,
  FlowPanel_GetCustomHint,
  FlowPanel_SetCustomHint,
  FlowPanel_GetComponentCount,
  FlowPanel_GetComponentIndex,
  FlowPanel_SetComponentIndex,
  FlowPanel_GetOwner,
  FlowPanel_GetName,
  FlowPanel_SetName,
  FlowPanel_GetTag,
  FlowPanel_SetTag,
  FlowPanel_GetDockClients,
  FlowPanel_GetControls,
  FlowPanel_GetComponents,
  FlowPanel_StaticClassType;

