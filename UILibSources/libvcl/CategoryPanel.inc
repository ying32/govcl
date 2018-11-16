
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function CategoryPanel_Create(AOwner: TComponent): TCategoryPanel; stdcall;
begin
  Result :=  TCategoryPanel.Create(AOwner);
end;

procedure CategoryPanel_Free(AObj: TCategoryPanel); stdcall;
begin
  AObj.Free;
end;

procedure CategoryPanel_Collapse(AObj: TCategoryPanel); stdcall;
begin
  AObj.Collapse;
end;

procedure CategoryPanel_Expand(AObj: TCategoryPanel); stdcall;
begin
  AObj.Expand;
end;

procedure CategoryPanel_SetBounds(AObj: TCategoryPanel; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

function CategoryPanel_CanFocus(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function CategoryPanel_ContainsControl(AObj: TCategoryPanel; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function CategoryPanel_ControlAtPos(AObj: TCategoryPanel; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure CategoryPanel_DisableAlign(AObj: TCategoryPanel); stdcall;
begin
  AObj.DisableAlign;
end;

procedure CategoryPanel_EnableAlign(AObj: TCategoryPanel); stdcall;
begin
  AObj.EnableAlign;
end;

function CategoryPanel_FindChildControl(AObj: TCategoryPanel; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure CategoryPanel_FlipChildren(AObj: TCategoryPanel; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function CategoryPanel_Focused(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function CategoryPanel_HandleAllocated(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure CategoryPanel_InsertControl(AObj: TCategoryPanel; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure CategoryPanel_Invalidate(AObj: TCategoryPanel); stdcall;
begin
  AObj.Invalidate;
end;

procedure CategoryPanel_PaintTo(AObj: TCategoryPanel; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure CategoryPanel_RemoveControl(AObj: TCategoryPanel; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure CategoryPanel_Realign(AObj: TCategoryPanel); stdcall;
begin
  AObj.Realign;
end;

procedure CategoryPanel_Repaint(AObj: TCategoryPanel); stdcall;
begin
  AObj.Repaint;
end;

procedure CategoryPanel_ScaleBy(AObj: TCategoryPanel; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure CategoryPanel_ScrollBy(AObj: TCategoryPanel; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure CategoryPanel_SetFocus(AObj: TCategoryPanel); stdcall;
begin
  AObj.SetFocus;
end;

procedure CategoryPanel_Update(AObj: TCategoryPanel); stdcall;
begin
  AObj.Update;
end;

procedure CategoryPanel_UpdateControlState(AObj: TCategoryPanel); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure CategoryPanel_BringToFront(AObj: TCategoryPanel); stdcall;
begin
  AObj.BringToFront;
end;

procedure CategoryPanel_ClientToScreen(AObj: TCategoryPanel; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure CategoryPanel_ClientToParent(AObj: TCategoryPanel; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function CategoryPanel_Dragging(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function CategoryPanel_HasParent(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure CategoryPanel_Hide(AObj: TCategoryPanel); stdcall;
begin
  AObj.Hide;
end;

function CategoryPanel_Perform(AObj: TCategoryPanel; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure CategoryPanel_Refresh(AObj: TCategoryPanel); stdcall;
begin
  AObj.Refresh;
end;

procedure CategoryPanel_ScreenToClient(AObj: TCategoryPanel; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure CategoryPanel_ParentToClient(AObj: TCategoryPanel; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure CategoryPanel_SendToBack(AObj: TCategoryPanel); stdcall;
begin
  AObj.SendToBack;
end;

procedure CategoryPanel_Show(AObj: TCategoryPanel); stdcall;
begin
  AObj.Show;
end;

function CategoryPanel_GetTextBuf(AObj: TCategoryPanel; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function CategoryPanel_GetTextLen(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure CategoryPanel_SetTextBuf(AObj: TCategoryPanel; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function CategoryPanel_FindComponent(AObj: TCategoryPanel; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function CategoryPanel_GetNamePath(AObj: TCategoryPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure CategoryPanel_Assign(AObj: TCategoryPanel; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure CategoryPanel_DisposeOf(AObj: TCategoryPanel); stdcall;
begin
  AObj.DisposeOf;
end;

function CategoryPanel_ClassType(AObj: TCategoryPanel): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function CategoryPanel_ClassName(AObj: TCategoryPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function CategoryPanel_InstanceSize(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function CategoryPanel_InheritsFrom(AObj: TCategoryPanel; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function CategoryPanel_Equals(AObj: TCategoryPanel; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function CategoryPanel_GetHashCode(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function CategoryPanel_ToString(AObj: TCategoryPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function CategoryPanel_GetBiDiMode(AObj: TCategoryPanel): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure CategoryPanel_SetBiDiMode(AObj: TCategoryPanel; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function CategoryPanel_GetCaption(AObj: TCategoryPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure CategoryPanel_SetCaption(AObj: TCategoryPanel; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function CategoryPanel_GetColor(AObj: TCategoryPanel): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure CategoryPanel_SetColor(AObj: TCategoryPanel; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function CategoryPanel_GetCollapsed(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.Collapsed;
end;

procedure CategoryPanel_SetCollapsed(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.Collapsed := AValue;
end;

function CategoryPanel_GetCollapsedHotImageIndex(AObj: TCategoryPanel): TImageIndex; stdcall;
begin
  Result :=  AObj.CollapsedHotImageIndex;
end;

procedure CategoryPanel_SetCollapsedHotImageIndex(AObj: TCategoryPanel; AValue: TImageIndex); stdcall;
begin
  AObj.CollapsedHotImageIndex := AValue;
end;

function CategoryPanel_GetCollapsedImageIndex(AObj: TCategoryPanel): TImageIndex; stdcall;
begin
  Result :=  AObj.CollapsedImageIndex;
end;

procedure CategoryPanel_SetCollapsedImageIndex(AObj: TCategoryPanel; AValue: TImageIndex); stdcall;
begin
  AObj.CollapsedImageIndex := AValue;
end;

function CategoryPanel_GetCollapsedPressedImageIndex(AObj: TCategoryPanel): TImageIndex; stdcall;
begin
  Result :=  AObj.CollapsedPressedImageIndex;
end;

procedure CategoryPanel_SetCollapsedPressedImageIndex(AObj: TCategoryPanel; AValue: TImageIndex); stdcall;
begin
  AObj.CollapsedPressedImageIndex := AValue;
end;

function CategoryPanel_GetUseDockManager(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure CategoryPanel_SetUseDockManager(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function CategoryPanel_GetDockSite(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure CategoryPanel_SetDockSite(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function CategoryPanel_GetDoubleBuffered(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure CategoryPanel_SetDoubleBuffered(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function CategoryPanel_GetDragCursor(AObj: TCategoryPanel): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure CategoryPanel_SetDragCursor(AObj: TCategoryPanel; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function CategoryPanel_GetDragKind(AObj: TCategoryPanel): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure CategoryPanel_SetDragKind(AObj: TCategoryPanel; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function CategoryPanel_GetDragMode(AObj: TCategoryPanel): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure CategoryPanel_SetDragMode(AObj: TCategoryPanel; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function CategoryPanel_GetEnabled(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure CategoryPanel_SetEnabled(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function CategoryPanel_GetExpandedHotImageIndex(AObj: TCategoryPanel): TImageIndex; stdcall;
begin
  Result :=  AObj.ExpandedHotImageIndex;
end;

procedure CategoryPanel_SetExpandedHotImageIndex(AObj: TCategoryPanel; AValue: TImageIndex); stdcall;
begin
  AObj.ExpandedHotImageIndex := AValue;
end;

function CategoryPanel_GetExpandedImageIndex(AObj: TCategoryPanel): TImageIndex; stdcall;
begin
  Result :=  AObj.ExpandedImageIndex;
end;

procedure CategoryPanel_SetExpandedImageIndex(AObj: TCategoryPanel; AValue: TImageIndex); stdcall;
begin
  AObj.ExpandedImageIndex := AValue;
end;

function CategoryPanel_GetExpandedPressedImageIndex(AObj: TCategoryPanel): TImageIndex; stdcall;
begin
  Result :=  AObj.ExpandedPressedImageIndex;
end;

procedure CategoryPanel_SetExpandedPressedImageIndex(AObj: TCategoryPanel; AValue: TImageIndex); stdcall;
begin
  AObj.ExpandedPressedImageIndex := AValue;
end;

function CategoryPanel_GetFullRepaint(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.FullRepaint;
end;

procedure CategoryPanel_SetFullRepaint(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.FullRepaint := AValue;
end;

function CategoryPanel_GetFont(AObj: TCategoryPanel): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure CategoryPanel_SetFont(AObj: TCategoryPanel; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function CategoryPanel_GetHeight(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure CategoryPanel_SetHeight(AObj: TCategoryPanel; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function CategoryPanel_GetLeft(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure CategoryPanel_SetLeft(AObj: TCategoryPanel; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function CategoryPanel_GetLocked(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.Locked;
end;

procedure CategoryPanel_SetLocked(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.Locked := AValue;
end;

function CategoryPanel_GetParentBackground(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentBackground;
end;

procedure CategoryPanel_SetParentBackground(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentBackground := AValue;
end;

function CategoryPanel_GetParentColor(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure CategoryPanel_SetParentColor(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function CategoryPanel_GetParentCtl3D(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure CategoryPanel_SetParentCtl3D(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function CategoryPanel_GetParentDoubleBuffered(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure CategoryPanel_SetParentDoubleBuffered(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function CategoryPanel_GetParentFont(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure CategoryPanel_SetParentFont(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function CategoryPanel_GetParentShowHint(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure CategoryPanel_SetParentShowHint(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function CategoryPanel_GetPopupMenu(AObj: TCategoryPanel): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure CategoryPanel_SetPopupMenu(AObj: TCategoryPanel; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function CategoryPanel_GetShowHint(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure CategoryPanel_SetShowHint(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function CategoryPanel_GetTabOrder(AObj: TCategoryPanel): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure CategoryPanel_SetTabOrder(AObj: TCategoryPanel; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function CategoryPanel_GetTabStop(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure CategoryPanel_SetTabStop(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function CategoryPanel_GetVisible(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure CategoryPanel_SetVisible(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function CategoryPanel_GetStyleElements(AObj: TCategoryPanel): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure CategoryPanel_SetStyleElements(AObj: TCategoryPanel; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function CategoryPanel_GetWidth(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure CategoryPanel_SetWidth(AObj: TCategoryPanel; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

procedure CategoryPanel_SetOnClick(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnContextPopup(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnDockDrop(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnDblClick(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnDragDrop(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnDragOver(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnEndDock(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnEndDrag(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnEnter(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnExit(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnGesture(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnGetSiteInfo(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnMouseActivate(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnMouseDown(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnMouseEnter(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnMouseLeave(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnMouseMove(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnMouseUp(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnStartDock(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanel_SetOnUnDock(AObj: TCategoryPanel; AEventId: NativeUInt); stdcall;
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

function CategoryPanel_GetPanelGroup(AObj: TCategoryPanel): TCustomCategoryPanelGroup; stdcall;
begin
  Result :=  AObj.PanelGroup;
end;

procedure CategoryPanel_SetPanelGroup(AObj: TCategoryPanel; AValue: TCustomCategoryPanelGroup); stdcall;
begin
  AObj.PanelGroup := AValue;
end;

function CategoryPanel_GetDockClientCount(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function CategoryPanel_GetAlignDisabled(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function CategoryPanel_GetMouseInClient(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function CategoryPanel_GetVisibleDockClientCount(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function CategoryPanel_GetBrush(AObj: TCategoryPanel): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function CategoryPanel_GetControlCount(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function CategoryPanel_GetHandle(AObj: TCategoryPanel): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function CategoryPanel_GetParentWindow(AObj: TCategoryPanel): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure CategoryPanel_SetParentWindow(AObj: TCategoryPanel; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function CategoryPanel_GetAction(AObj: TCategoryPanel): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure CategoryPanel_SetAction(AObj: TCategoryPanel; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function CategoryPanel_GetAlign(AObj: TCategoryPanel): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure CategoryPanel_SetAlign(AObj: TCategoryPanel; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function CategoryPanel_GetAnchors(AObj: TCategoryPanel): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure CategoryPanel_SetAnchors(AObj: TCategoryPanel; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

procedure CategoryPanel_GetBoundsRect(AObj: TCategoryPanel; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure CategoryPanel_SetBoundsRect(AObj: TCategoryPanel; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function CategoryPanel_GetClientHeight(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure CategoryPanel_SetClientHeight(AObj: TCategoryPanel; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure CategoryPanel_GetClientOrigin(AObj: TCategoryPanel; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure CategoryPanel_GetClientRect(AObj: TCategoryPanel; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function CategoryPanel_GetClientWidth(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure CategoryPanel_SetClientWidth(AObj: TCategoryPanel; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function CategoryPanel_GetControlState(AObj: TCategoryPanel): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure CategoryPanel_SetControlState(AObj: TCategoryPanel; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function CategoryPanel_GetControlStyle(AObj: TCategoryPanel): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure CategoryPanel_SetControlStyle(AObj: TCategoryPanel; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function CategoryPanel_GetExplicitLeft(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function CategoryPanel_GetExplicitTop(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function CategoryPanel_GetExplicitWidth(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function CategoryPanel_GetExplicitHeight(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function CategoryPanel_GetFloating(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function CategoryPanel_GetParent(AObj: TCategoryPanel): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure CategoryPanel_SetParent(AObj: TCategoryPanel; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function CategoryPanel_GetAlignWithMargins(AObj: TCategoryPanel): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure CategoryPanel_SetAlignWithMargins(AObj: TCategoryPanel; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function CategoryPanel_GetTop(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure CategoryPanel_SetTop(AObj: TCategoryPanel; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function CategoryPanel_GetCursor(AObj: TCategoryPanel): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure CategoryPanel_SetCursor(AObj: TCategoryPanel; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function CategoryPanel_GetHint(AObj: TCategoryPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure CategoryPanel_SetHint(AObj: TCategoryPanel; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function CategoryPanel_GetMargins(AObj: TCategoryPanel): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure CategoryPanel_SetMargins(AObj: TCategoryPanel; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function CategoryPanel_GetCustomHint(AObj: TCategoryPanel): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure CategoryPanel_SetCustomHint(AObj: TCategoryPanel; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function CategoryPanel_GetComponentCount(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function CategoryPanel_GetComponentIndex(AObj: TCategoryPanel): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure CategoryPanel_SetComponentIndex(AObj: TCategoryPanel; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function CategoryPanel_GetOwner(AObj: TCategoryPanel): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function CategoryPanel_GetName(AObj: TCategoryPanel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure CategoryPanel_SetName(AObj: TCategoryPanel; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function CategoryPanel_GetTag(AObj: TCategoryPanel): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure CategoryPanel_SetTag(AObj: TCategoryPanel; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function CategoryPanel_GetDockClients(AObj: TCategoryPanel; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function CategoryPanel_GetControls(AObj: TCategoryPanel; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function CategoryPanel_GetComponents(AObj: TCategoryPanel; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function CategoryPanel_StaticClassType: TClass; stdcall;
begin
  Result := TCategoryPanel;
end;

exports
  CategoryPanel_Create,
  CategoryPanel_Free,
  CategoryPanel_Collapse,
  CategoryPanel_Expand,
  CategoryPanel_SetBounds,
  CategoryPanel_CanFocus,
  CategoryPanel_ContainsControl,
  CategoryPanel_ControlAtPos,
  CategoryPanel_DisableAlign,
  CategoryPanel_EnableAlign,
  CategoryPanel_FindChildControl,
  CategoryPanel_FlipChildren,
  CategoryPanel_Focused,
  CategoryPanel_HandleAllocated,
  CategoryPanel_InsertControl,
  CategoryPanel_Invalidate,
  CategoryPanel_PaintTo,
  CategoryPanel_RemoveControl,
  CategoryPanel_Realign,
  CategoryPanel_Repaint,
  CategoryPanel_ScaleBy,
  CategoryPanel_ScrollBy,
  CategoryPanel_SetFocus,
  CategoryPanel_Update,
  CategoryPanel_UpdateControlState,
  CategoryPanel_BringToFront,
  CategoryPanel_ClientToScreen,
  CategoryPanel_ClientToParent,
  CategoryPanel_Dragging,
  CategoryPanel_HasParent,
  CategoryPanel_Hide,
  CategoryPanel_Perform,
  CategoryPanel_Refresh,
  CategoryPanel_ScreenToClient,
  CategoryPanel_ParentToClient,
  CategoryPanel_SendToBack,
  CategoryPanel_Show,
  CategoryPanel_GetTextBuf,
  CategoryPanel_GetTextLen,
  CategoryPanel_SetTextBuf,
  CategoryPanel_FindComponent,
  CategoryPanel_GetNamePath,
  CategoryPanel_Assign,
  CategoryPanel_DisposeOf,
  CategoryPanel_ClassType,
  CategoryPanel_ClassName,
  CategoryPanel_InstanceSize,
  CategoryPanel_InheritsFrom,
  CategoryPanel_Equals,
  CategoryPanel_GetHashCode,
  CategoryPanel_ToString,
  CategoryPanel_GetBiDiMode,
  CategoryPanel_SetBiDiMode,
  CategoryPanel_GetCaption,
  CategoryPanel_SetCaption,
  CategoryPanel_GetColor,
  CategoryPanel_SetColor,
  CategoryPanel_GetCollapsed,
  CategoryPanel_SetCollapsed,
  CategoryPanel_GetCollapsedHotImageIndex,
  CategoryPanel_SetCollapsedHotImageIndex,
  CategoryPanel_GetCollapsedImageIndex,
  CategoryPanel_SetCollapsedImageIndex,
  CategoryPanel_GetCollapsedPressedImageIndex,
  CategoryPanel_SetCollapsedPressedImageIndex,
  CategoryPanel_GetUseDockManager,
  CategoryPanel_SetUseDockManager,
  CategoryPanel_GetDockSite,
  CategoryPanel_SetDockSite,
  CategoryPanel_GetDoubleBuffered,
  CategoryPanel_SetDoubleBuffered,
  CategoryPanel_GetDragCursor,
  CategoryPanel_SetDragCursor,
  CategoryPanel_GetDragKind,
  CategoryPanel_SetDragKind,
  CategoryPanel_GetDragMode,
  CategoryPanel_SetDragMode,
  CategoryPanel_GetEnabled,
  CategoryPanel_SetEnabled,
  CategoryPanel_GetExpandedHotImageIndex,
  CategoryPanel_SetExpandedHotImageIndex,
  CategoryPanel_GetExpandedImageIndex,
  CategoryPanel_SetExpandedImageIndex,
  CategoryPanel_GetExpandedPressedImageIndex,
  CategoryPanel_SetExpandedPressedImageIndex,
  CategoryPanel_GetFullRepaint,
  CategoryPanel_SetFullRepaint,
  CategoryPanel_GetFont,
  CategoryPanel_SetFont,
  CategoryPanel_GetHeight,
  CategoryPanel_SetHeight,
  CategoryPanel_GetLeft,
  CategoryPanel_SetLeft,
  CategoryPanel_GetLocked,
  CategoryPanel_SetLocked,
  CategoryPanel_GetParentBackground,
  CategoryPanel_SetParentBackground,
  CategoryPanel_GetParentColor,
  CategoryPanel_SetParentColor,
  CategoryPanel_GetParentCtl3D,
  CategoryPanel_SetParentCtl3D,
  CategoryPanel_GetParentDoubleBuffered,
  CategoryPanel_SetParentDoubleBuffered,
  CategoryPanel_GetParentFont,
  CategoryPanel_SetParentFont,
  CategoryPanel_GetParentShowHint,
  CategoryPanel_SetParentShowHint,
  CategoryPanel_GetPopupMenu,
  CategoryPanel_SetPopupMenu,
  CategoryPanel_GetShowHint,
  CategoryPanel_SetShowHint,
  CategoryPanel_GetTabOrder,
  CategoryPanel_SetTabOrder,
  CategoryPanel_GetTabStop,
  CategoryPanel_SetTabStop,
  CategoryPanel_GetVisible,
  CategoryPanel_SetVisible,
  CategoryPanel_GetStyleElements,
  CategoryPanel_SetStyleElements,
  CategoryPanel_GetWidth,
  CategoryPanel_SetWidth,
  CategoryPanel_SetOnClick,
  CategoryPanel_SetOnContextPopup,
  CategoryPanel_SetOnDockDrop,
  CategoryPanel_SetOnDblClick,
  CategoryPanel_SetOnDragDrop,
  CategoryPanel_SetOnDragOver,
  CategoryPanel_SetOnEndDock,
  CategoryPanel_SetOnEndDrag,
  CategoryPanel_SetOnEnter,
  CategoryPanel_SetOnExit,
  CategoryPanel_SetOnGesture,
  CategoryPanel_SetOnGetSiteInfo,
  CategoryPanel_SetOnMouseActivate,
  CategoryPanel_SetOnMouseDown,
  CategoryPanel_SetOnMouseEnter,
  CategoryPanel_SetOnMouseLeave,
  CategoryPanel_SetOnMouseMove,
  CategoryPanel_SetOnMouseUp,
  CategoryPanel_SetOnStartDock,
  CategoryPanel_SetOnUnDock,
  CategoryPanel_GetPanelGroup,
  CategoryPanel_SetPanelGroup,
  CategoryPanel_GetDockClientCount,
  CategoryPanel_GetAlignDisabled,
  CategoryPanel_GetMouseInClient,
  CategoryPanel_GetVisibleDockClientCount,
  CategoryPanel_GetBrush,
  CategoryPanel_GetControlCount,
  CategoryPanel_GetHandle,
  CategoryPanel_GetParentWindow,
  CategoryPanel_SetParentWindow,
  CategoryPanel_GetAction,
  CategoryPanel_SetAction,
  CategoryPanel_GetAlign,
  CategoryPanel_SetAlign,
  CategoryPanel_GetAnchors,
  CategoryPanel_SetAnchors,
  CategoryPanel_GetBoundsRect,
  CategoryPanel_SetBoundsRect,
  CategoryPanel_GetClientHeight,
  CategoryPanel_SetClientHeight,
  CategoryPanel_GetClientOrigin,
  CategoryPanel_GetClientRect,
  CategoryPanel_GetClientWidth,
  CategoryPanel_SetClientWidth,
  CategoryPanel_GetControlState,
  CategoryPanel_SetControlState,
  CategoryPanel_GetControlStyle,
  CategoryPanel_SetControlStyle,
  CategoryPanel_GetExplicitLeft,
  CategoryPanel_GetExplicitTop,
  CategoryPanel_GetExplicitWidth,
  CategoryPanel_GetExplicitHeight,
  CategoryPanel_GetFloating,
  CategoryPanel_GetParent,
  CategoryPanel_SetParent,
  CategoryPanel_GetAlignWithMargins,
  CategoryPanel_SetAlignWithMargins,
  CategoryPanel_GetTop,
  CategoryPanel_SetTop,
  CategoryPanel_GetCursor,
  CategoryPanel_SetCursor,
  CategoryPanel_GetHint,
  CategoryPanel_SetHint,
  CategoryPanel_GetMargins,
  CategoryPanel_SetMargins,
  CategoryPanel_GetCustomHint,
  CategoryPanel_SetCustomHint,
  CategoryPanel_GetComponentCount,
  CategoryPanel_GetComponentIndex,
  CategoryPanel_SetComponentIndex,
  CategoryPanel_GetOwner,
  CategoryPanel_GetName,
  CategoryPanel_SetName,
  CategoryPanel_GetTag,
  CategoryPanel_SetTag,
  CategoryPanel_GetDockClients,
  CategoryPanel_GetControls,
  CategoryPanel_GetComponents,
  CategoryPanel_StaticClassType;

