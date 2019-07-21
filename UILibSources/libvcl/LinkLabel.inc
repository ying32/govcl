
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
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

function LinkLabel_ContainsControl(AObj: TLinkLabel; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function LinkLabel_ControlAtPos(AObj: TLinkLabel; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure LinkLabel_DisableAlign(AObj: TLinkLabel); stdcall;
begin
  AObj.DisableAlign;
end;

procedure LinkLabel_EnableAlign(AObj: TLinkLabel); stdcall;
begin
  AObj.EnableAlign;
end;

function LinkLabel_FindChildControl(AObj: TLinkLabel; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
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

procedure LinkLabel_InsertControl(AObj: TLinkLabel; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure LinkLabel_Invalidate(AObj: TLinkLabel); stdcall;
begin
  AObj.Invalidate;
end;

procedure LinkLabel_PaintTo(AObj: TLinkLabel; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure LinkLabel_RemoveControl(AObj: TLinkLabel; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
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

procedure LinkLabel_ScrollBy(AObj: TLinkLabel; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
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

procedure LinkLabel_UpdateControlState(AObj: TLinkLabel); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure LinkLabel_BringToFront(AObj: TLinkLabel); stdcall;
begin
  AObj.BringToFront;
end;

procedure LinkLabel_ClientToScreen(AObj: TLinkLabel; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure LinkLabel_ClientToParent(AObj: TLinkLabel; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function LinkLabel_Dragging(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
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

procedure LinkLabel_ScreenToClient(AObj: TLinkLabel; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure LinkLabel_ParentToClient(AObj: TLinkLabel; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
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

procedure LinkLabel_SetTextBuf(AObj: TLinkLabel; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
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

procedure LinkLabel_DisposeOf(AObj: TLinkLabel); stdcall;
begin
  AObj.DisposeOf;
end;

function LinkLabel_ClassType(AObj: TLinkLabel): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function LinkLabel_ClassName(AObj: TLinkLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function LinkLabel_InstanceSize(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function LinkLabel_InheritsFrom(AObj: TLinkLabel; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
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

function LinkLabel_GetConstraints(AObj: TLinkLabel): TSizeConstraints; stdcall;
begin
  Result :=  AObj.Constraints;
end;

procedure LinkLabel_SetConstraints(AObj: TLinkLabel; AValue: TSizeConstraints); stdcall;
begin
  AObj.Constraints := AValue;
end;

function LinkLabel_GetDragCursor(AObj: TLinkLabel): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure LinkLabel_SetDragCursor(AObj: TLinkLabel; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function LinkLabel_GetDragKind(AObj: TLinkLabel): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure LinkLabel_SetDragKind(AObj: TLinkLabel; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function LinkLabel_GetDragMode(AObj: TLinkLabel): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure LinkLabel_SetDragMode(AObj: TLinkLabel; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
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
    TEventClass.Remove(AObj, @TEventClass.OnClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, @TEventClass.OnClick, AEventId);
end;

procedure LinkLabel_SetOnContextPopup(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnDblClick(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnDragDrop(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnDragOver(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnEndDock(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnEndDrag(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnGesture(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseActivate(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseDown(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseEnter(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseLeave(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseMove(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnMouseUp(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnStartDock(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
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

procedure LinkLabel_SetOnLinkClick(AObj: TLinkLabel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnLinkClick := nil;
    TEventClass.Remove(AObj, @TEventClass.OnLinkClick);
    Exit;
  end;
  AObj.OnLinkClick := TEventClass.OnLinkClick;
  TEventClass.Add(AObj, @TEventClass.OnLinkClick, AEventId);
end;

function LinkLabel_GetDockClientCount(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function LinkLabel_GetDockSite(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure LinkLabel_SetDockSite(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function LinkLabel_GetDoubleBuffered(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure LinkLabel_SetDoubleBuffered(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function LinkLabel_GetAlignDisabled(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function LinkLabel_GetMouseInClient(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function LinkLabel_GetVisibleDockClientCount(AObj: TLinkLabel): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
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

function LinkLabel_GetUseDockManager(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure LinkLabel_SetUseDockManager(AObj: TLinkLabel; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
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

procedure LinkLabel_GetClientOrigin(AObj: TLinkLabel; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
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

function LinkLabel_GetControlState(AObj: TLinkLabel): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure LinkLabel_SetControlState(AObj: TLinkLabel; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function LinkLabel_GetControlStyle(AObj: TLinkLabel): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure LinkLabel_SetControlStyle(AObj: TLinkLabel; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
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

function LinkLabel_GetFloating(AObj: TLinkLabel): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
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

function LinkLabel_GetDockClients(AObj: TLinkLabel; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function LinkLabel_GetControls(AObj: TLinkLabel; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function LinkLabel_GetComponents(AObj: TLinkLabel; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function LinkLabel_StaticClassType: TClass; stdcall;
begin
  Result := TLinkLabel;
end;

exports
  LinkLabel_Create,
  LinkLabel_Free,
  LinkLabel_CanFocus,
  LinkLabel_ContainsControl,
  LinkLabel_ControlAtPos,
  LinkLabel_DisableAlign,
  LinkLabel_EnableAlign,
  LinkLabel_FindChildControl,
  LinkLabel_FlipChildren,
  LinkLabel_Focused,
  LinkLabel_HandleAllocated,
  LinkLabel_InsertControl,
  LinkLabel_Invalidate,
  LinkLabel_PaintTo,
  LinkLabel_RemoveControl,
  LinkLabel_Realign,
  LinkLabel_Repaint,
  LinkLabel_ScaleBy,
  LinkLabel_ScrollBy,
  LinkLabel_SetBounds,
  LinkLabel_SetFocus,
  LinkLabel_Update,
  LinkLabel_UpdateControlState,
  LinkLabel_BringToFront,
  LinkLabel_ClientToScreen,
  LinkLabel_ClientToParent,
  LinkLabel_Dragging,
  LinkLabel_HasParent,
  LinkLabel_Hide,
  LinkLabel_Perform,
  LinkLabel_Refresh,
  LinkLabel_ScreenToClient,
  LinkLabel_ParentToClient,
  LinkLabel_SendToBack,
  LinkLabel_Show,
  LinkLabel_GetTextBuf,
  LinkLabel_GetTextLen,
  LinkLabel_SetTextBuf,
  LinkLabel_FindComponent,
  LinkLabel_GetNamePath,
  LinkLabel_Assign,
  LinkLabel_DisposeOf,
  LinkLabel_ClassType,
  LinkLabel_ClassName,
  LinkLabel_InstanceSize,
  LinkLabel_InheritsFrom,
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
  LinkLabel_GetConstraints,
  LinkLabel_SetConstraints,
  LinkLabel_GetDragCursor,
  LinkLabel_SetDragCursor,
  LinkLabel_GetDragKind,
  LinkLabel_SetDragKind,
  LinkLabel_GetDragMode,
  LinkLabel_SetDragMode,
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
  LinkLabel_SetOnContextPopup,
  LinkLabel_SetOnDblClick,
  LinkLabel_SetOnDragDrop,
  LinkLabel_SetOnDragOver,
  LinkLabel_SetOnEndDock,
  LinkLabel_SetOnEndDrag,
  LinkLabel_SetOnGesture,
  LinkLabel_SetOnMouseActivate,
  LinkLabel_SetOnMouseDown,
  LinkLabel_SetOnMouseEnter,
  LinkLabel_SetOnMouseLeave,
  LinkLabel_SetOnMouseMove,
  LinkLabel_SetOnMouseUp,
  LinkLabel_SetOnStartDock,
  LinkLabel_SetOnLinkClick,
  LinkLabel_GetDockClientCount,
  LinkLabel_GetDockSite,
  LinkLabel_SetDockSite,
  LinkLabel_GetDoubleBuffered,
  LinkLabel_SetDoubleBuffered,
  LinkLabel_GetAlignDisabled,
  LinkLabel_GetMouseInClient,
  LinkLabel_GetVisibleDockClientCount,
  LinkLabel_GetBrush,
  LinkLabel_GetControlCount,
  LinkLabel_GetHandle,
  LinkLabel_GetParentDoubleBuffered,
  LinkLabel_SetParentDoubleBuffered,
  LinkLabel_GetParentWindow,
  LinkLabel_SetParentWindow,
  LinkLabel_GetUseDockManager,
  LinkLabel_SetUseDockManager,
  LinkLabel_GetAction,
  LinkLabel_SetAction,
  LinkLabel_GetBiDiMode,
  LinkLabel_SetBiDiMode,
  LinkLabel_GetBoundsRect,
  LinkLabel_SetBoundsRect,
  LinkLabel_GetClientHeight,
  LinkLabel_SetClientHeight,
  LinkLabel_GetClientOrigin,
  LinkLabel_GetClientRect,
  LinkLabel_GetClientWidth,
  LinkLabel_SetClientWidth,
  LinkLabel_GetControlState,
  LinkLabel_SetControlState,
  LinkLabel_GetControlStyle,
  LinkLabel_SetControlStyle,
  LinkLabel_GetExplicitLeft,
  LinkLabel_GetExplicitTop,
  LinkLabel_GetExplicitWidth,
  LinkLabel_GetExplicitHeight,
  LinkLabel_GetFloating,
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
  LinkLabel_GetDockClients,
  LinkLabel_GetControls,
  LinkLabel_GetComponents,
  LinkLabel_StaticClassType;

