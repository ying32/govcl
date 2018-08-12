
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function CoolBar_Create(AOwner: TComponent): TCoolBar; stdcall;
begin
  Result :=  TCoolBar.Create(AOwner);
end;

procedure CoolBar_Free(AObj: TCoolBar); stdcall;
begin
  AObj.Free;
end;

procedure CoolBar_FlipChildren(AObj: TCoolBar; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function CoolBar_CanFocus(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function CoolBar_ContainsControl(AObj: TCoolBar; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function CoolBar_ControlAtPos(AObj: TCoolBar; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure CoolBar_DisableAlign(AObj: TCoolBar); stdcall;
begin
  AObj.DisableAlign;
end;

procedure CoolBar_EnableAlign(AObj: TCoolBar); stdcall;
begin
  AObj.EnableAlign;
end;

function CoolBar_FindChildControl(AObj: TCoolBar; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

function CoolBar_Focused(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function CoolBar_HandleAllocated(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure CoolBar_InsertControl(AObj: TCoolBar; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure CoolBar_Invalidate(AObj: TCoolBar); stdcall;
begin
  AObj.Invalidate;
end;

procedure CoolBar_PaintTo(AObj: TCoolBar; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure CoolBar_RemoveControl(AObj: TCoolBar; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure CoolBar_Realign(AObj: TCoolBar); stdcall;
begin
  AObj.Realign;
end;

procedure CoolBar_Repaint(AObj: TCoolBar); stdcall;
begin
  AObj.Repaint;
end;

procedure CoolBar_ScaleBy(AObj: TCoolBar; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure CoolBar_ScrollBy(AObj: TCoolBar; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure CoolBar_SetBounds(AObj: TCoolBar; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure CoolBar_SetFocus(AObj: TCoolBar); stdcall;
begin
  AObj.SetFocus;
end;

procedure CoolBar_Update(AObj: TCoolBar); stdcall;
begin
  AObj.Update;
end;

procedure CoolBar_UpdateControlState(AObj: TCoolBar); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure CoolBar_BringToFront(AObj: TCoolBar); stdcall;
begin
  AObj.BringToFront;
end;

procedure CoolBar_ClientToScreen(AObj: TCoolBar; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure CoolBar_ClientToParent(AObj: TCoolBar; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function CoolBar_Dragging(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function CoolBar_HasParent(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure CoolBar_Hide(AObj: TCoolBar); stdcall;
begin
  AObj.Hide;
end;

function CoolBar_Perform(AObj: TCoolBar; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure CoolBar_Refresh(AObj: TCoolBar); stdcall;
begin
  AObj.Refresh;
end;

procedure CoolBar_ScreenToClient(AObj: TCoolBar; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure CoolBar_ParentToClient(AObj: TCoolBar; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure CoolBar_SendToBack(AObj: TCoolBar); stdcall;
begin
  AObj.SendToBack;
end;

procedure CoolBar_Show(AObj: TCoolBar); stdcall;
begin
  AObj.Show;
end;

function CoolBar_GetTextBuf(AObj: TCoolBar; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function CoolBar_GetTextLen(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure CoolBar_SetTextBuf(AObj: TCoolBar; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function CoolBar_FindComponent(AObj: TCoolBar; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function CoolBar_GetNamePath(AObj: TCoolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure CoolBar_Assign(AObj: TCoolBar; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure CoolBar_DisposeOf(AObj: TCoolBar); stdcall;
begin
  AObj.DisposeOf;
end;

function CoolBar_ClassType(AObj: TCoolBar): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function CoolBar_ClassName(AObj: TCoolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function CoolBar_InstanceSize(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function CoolBar_InheritsFrom(AObj: TCoolBar; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function CoolBar_Equals(AObj: TCoolBar; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function CoolBar_GetHashCode(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function CoolBar_ToString(AObj: TCoolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function CoolBar_GetAlign(AObj: TCoolBar): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure CoolBar_SetAlign(AObj: TCoolBar; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function CoolBar_GetAnchors(AObj: TCoolBar): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure CoolBar_SetAnchors(AObj: TCoolBar; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function CoolBar_GetAutoSize(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure CoolBar_SetAutoSize(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function CoolBar_GetBandBorderStyle(AObj: TCoolBar): TBorderStyle; stdcall;
begin
  Result :=  AObj.BandBorderStyle;
end;

procedure CoolBar_SetBandBorderStyle(AObj: TCoolBar; AValue: TBorderStyle); stdcall;
begin
  AObj.BandBorderStyle := AValue;
end;

function CoolBar_GetBandMaximize(AObj: TCoolBar): TCoolBandMaximize; stdcall;
begin
  Result :=  AObj.BandMaximize;
end;

procedure CoolBar_SetBandMaximize(AObj: TCoolBar; AValue: TCoolBandMaximize); stdcall;
begin
  AObj.BandMaximize := AValue;
end;

function CoolBar_GetBands(AObj: TCoolBar): TCoolBands; stdcall;
begin
  Result :=  AObj.Bands;
end;

procedure CoolBar_SetBands(AObj: TCoolBar; AValue: TCoolBands); stdcall;
begin
  AObj.Bands := AValue;
end;

function CoolBar_GetBorderWidth(AObj: TCoolBar): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure CoolBar_SetBorderWidth(AObj: TCoolBar; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function CoolBar_GetColor(AObj: TCoolBar): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure CoolBar_SetColor(AObj: TCoolBar; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function CoolBar_GetDockSite(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure CoolBar_SetDockSite(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function CoolBar_GetDoubleBuffered(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure CoolBar_SetDoubleBuffered(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function CoolBar_GetDragCursor(AObj: TCoolBar): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure CoolBar_SetDragCursor(AObj: TCoolBar; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function CoolBar_GetDragKind(AObj: TCoolBar): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure CoolBar_SetDragKind(AObj: TCoolBar; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function CoolBar_GetDragMode(AObj: TCoolBar): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure CoolBar_SetDragMode(AObj: TCoolBar; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function CoolBar_GetEdgeBorders(AObj: TCoolBar): TEdgeBorders; stdcall;
begin
  Result :=  AObj.EdgeBorders;
end;

procedure CoolBar_SetEdgeBorders(AObj: TCoolBar; AValue: TEdgeBorders); stdcall;
begin
  AObj.EdgeBorders := AValue;
end;

function CoolBar_GetEdgeInner(AObj: TCoolBar): TEdgeStyle; stdcall;
begin
  Result :=  AObj.EdgeInner;
end;

procedure CoolBar_SetEdgeInner(AObj: TCoolBar; AValue: TEdgeStyle); stdcall;
begin
  AObj.EdgeInner := AValue;
end;

function CoolBar_GetEdgeOuter(AObj: TCoolBar): TEdgeStyle; stdcall;
begin
  Result :=  AObj.EdgeOuter;
end;

procedure CoolBar_SetEdgeOuter(AObj: TCoolBar; AValue: TEdgeStyle); stdcall;
begin
  AObj.EdgeOuter := AValue;
end;

function CoolBar_GetEnabled(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure CoolBar_SetEnabled(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function CoolBar_GetFixedSize(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.FixedSize;
end;

procedure CoolBar_SetFixedSize(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.FixedSize := AValue;
end;

function CoolBar_GetFixedOrder(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.FixedOrder;
end;

procedure CoolBar_SetFixedOrder(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.FixedOrder := AValue;
end;

function CoolBar_GetFont(AObj: TCoolBar): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure CoolBar_SetFont(AObj: TCoolBar; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function CoolBar_GetImages(AObj: TCoolBar): TCustomImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure CoolBar_SetImages(AObj: TCoolBar; AValue: TCustomImageList); stdcall;
begin
  AObj.Images := AValue;
end;

function CoolBar_GetParentColor(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure CoolBar_SetParentColor(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function CoolBar_GetParentDoubleBuffered(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure CoolBar_SetParentDoubleBuffered(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function CoolBar_GetParentFont(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure CoolBar_SetParentFont(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function CoolBar_GetParentShowHint(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure CoolBar_SetParentShowHint(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function CoolBar_GetBitmap(AObj: TCoolBar): TBitmap; stdcall;
begin
  Result :=  AObj.Bitmap;
end;

procedure CoolBar_SetBitmap(AObj: TCoolBar; AValue: TBitmap); stdcall;
begin
  AObj.Bitmap := AValue;
end;

function CoolBar_GetPopupMenu(AObj: TCoolBar): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure CoolBar_SetPopupMenu(AObj: TCoolBar; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function CoolBar_GetShowHint(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure CoolBar_SetShowHint(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function CoolBar_GetShowText(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.ShowText;
end;

procedure CoolBar_SetShowText(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.ShowText := AValue;
end;

function CoolBar_GetVertical(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.Vertical;
end;

procedure CoolBar_SetVertical(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.Vertical := AValue;
end;

function CoolBar_GetVisible(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure CoolBar_SetVisible(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function CoolBar_GetStyleElements(AObj: TCoolBar): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure CoolBar_SetStyleElements(AObj: TCoolBar; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure CoolBar_SetOnChange(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, geChange, AEventId);
end;

procedure CoolBar_SetOnClick(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnContextPopup(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnDblClick(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnDockDrop(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnDragDrop(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnDragOver(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnEndDock(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnEndDrag(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnGetSiteInfo(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnMouseDown(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnMouseEnter(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnMouseLeave(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnMouseMove(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnMouseUp(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnResize(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnStartDock(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

procedure CoolBar_SetOnUnDock(AObj: TCoolBar; AEventId: NativeUInt); stdcall;
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

function CoolBar_GetDockClientCount(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function CoolBar_GetAlignDisabled(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function CoolBar_GetMouseInClient(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function CoolBar_GetVisibleDockClientCount(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function CoolBar_GetBrush(AObj: TCoolBar): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function CoolBar_GetControlCount(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function CoolBar_GetHandle(AObj: TCoolBar): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function CoolBar_GetParentWindow(AObj: TCoolBar): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure CoolBar_SetParentWindow(AObj: TCoolBar; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function CoolBar_GetTabOrder(AObj: TCoolBar): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure CoolBar_SetTabOrder(AObj: TCoolBar; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function CoolBar_GetTabStop(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure CoolBar_SetTabStop(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function CoolBar_GetUseDockManager(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure CoolBar_SetUseDockManager(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function CoolBar_GetAction(AObj: TCoolBar): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure CoolBar_SetAction(AObj: TCoolBar; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function CoolBar_GetBiDiMode(AObj: TCoolBar): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure CoolBar_SetBiDiMode(AObj: TCoolBar; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure CoolBar_GetBoundsRect(AObj: TCoolBar; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure CoolBar_SetBoundsRect(AObj: TCoolBar; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function CoolBar_GetClientHeight(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure CoolBar_SetClientHeight(AObj: TCoolBar; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure CoolBar_GetClientOrigin(AObj: TCoolBar; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure CoolBar_GetClientRect(AObj: TCoolBar; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function CoolBar_GetClientWidth(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure CoolBar_SetClientWidth(AObj: TCoolBar; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function CoolBar_GetControlState(AObj: TCoolBar): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure CoolBar_SetControlState(AObj: TCoolBar; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function CoolBar_GetControlStyle(AObj: TCoolBar): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure CoolBar_SetControlStyle(AObj: TCoolBar; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function CoolBar_GetExplicitLeft(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function CoolBar_GetExplicitTop(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function CoolBar_GetExplicitWidth(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function CoolBar_GetExplicitHeight(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function CoolBar_GetFloating(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function CoolBar_GetParent(AObj: TCoolBar): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure CoolBar_SetParent(AObj: TCoolBar; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function CoolBar_GetAlignWithMargins(AObj: TCoolBar): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure CoolBar_SetAlignWithMargins(AObj: TCoolBar; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function CoolBar_GetLeft(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure CoolBar_SetLeft(AObj: TCoolBar; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function CoolBar_GetTop(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure CoolBar_SetTop(AObj: TCoolBar; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function CoolBar_GetWidth(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure CoolBar_SetWidth(AObj: TCoolBar; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function CoolBar_GetHeight(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure CoolBar_SetHeight(AObj: TCoolBar; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function CoolBar_GetCursor(AObj: TCoolBar): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure CoolBar_SetCursor(AObj: TCoolBar; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function CoolBar_GetHint(AObj: TCoolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure CoolBar_SetHint(AObj: TCoolBar; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function CoolBar_GetMargins(AObj: TCoolBar): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure CoolBar_SetMargins(AObj: TCoolBar; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function CoolBar_GetCustomHint(AObj: TCoolBar): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure CoolBar_SetCustomHint(AObj: TCoolBar; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function CoolBar_GetComponentCount(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function CoolBar_GetComponentIndex(AObj: TCoolBar): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure CoolBar_SetComponentIndex(AObj: TCoolBar; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function CoolBar_GetOwner(AObj: TCoolBar): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function CoolBar_GetName(AObj: TCoolBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure CoolBar_SetName(AObj: TCoolBar; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function CoolBar_GetTag(AObj: TCoolBar): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure CoolBar_SetTag(AObj: TCoolBar; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function CoolBar_GetDockClients(AObj: TCoolBar; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function CoolBar_GetControls(AObj: TCoolBar; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function CoolBar_GetComponents(AObj: TCoolBar; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function CoolBar_StaticClassType: TClass; stdcall;
begin
  Result := TCoolBar;
end;

exports
  CoolBar_Create,
  CoolBar_Free,
  CoolBar_FlipChildren,
  CoolBar_CanFocus,
  CoolBar_ContainsControl,
  CoolBar_ControlAtPos,
  CoolBar_DisableAlign,
  CoolBar_EnableAlign,
  CoolBar_FindChildControl,
  CoolBar_Focused,
  CoolBar_HandleAllocated,
  CoolBar_InsertControl,
  CoolBar_Invalidate,
  CoolBar_PaintTo,
  CoolBar_RemoveControl,
  CoolBar_Realign,
  CoolBar_Repaint,
  CoolBar_ScaleBy,
  CoolBar_ScrollBy,
  CoolBar_SetBounds,
  CoolBar_SetFocus,
  CoolBar_Update,
  CoolBar_UpdateControlState,
  CoolBar_BringToFront,
  CoolBar_ClientToScreen,
  CoolBar_ClientToParent,
  CoolBar_Dragging,
  CoolBar_HasParent,
  CoolBar_Hide,
  CoolBar_Perform,
  CoolBar_Refresh,
  CoolBar_ScreenToClient,
  CoolBar_ParentToClient,
  CoolBar_SendToBack,
  CoolBar_Show,
  CoolBar_GetTextBuf,
  CoolBar_GetTextLen,
  CoolBar_SetTextBuf,
  CoolBar_FindComponent,
  CoolBar_GetNamePath,
  CoolBar_Assign,
  CoolBar_DisposeOf,
  CoolBar_ClassType,
  CoolBar_ClassName,
  CoolBar_InstanceSize,
  CoolBar_InheritsFrom,
  CoolBar_Equals,
  CoolBar_GetHashCode,
  CoolBar_ToString,
  CoolBar_GetAlign,
  CoolBar_SetAlign,
  CoolBar_GetAnchors,
  CoolBar_SetAnchors,
  CoolBar_GetAutoSize,
  CoolBar_SetAutoSize,
  CoolBar_GetBandBorderStyle,
  CoolBar_SetBandBorderStyle,
  CoolBar_GetBandMaximize,
  CoolBar_SetBandMaximize,
  CoolBar_GetBands,
  CoolBar_SetBands,
  CoolBar_GetBorderWidth,
  CoolBar_SetBorderWidth,
  CoolBar_GetColor,
  CoolBar_SetColor,
  CoolBar_GetDockSite,
  CoolBar_SetDockSite,
  CoolBar_GetDoubleBuffered,
  CoolBar_SetDoubleBuffered,
  CoolBar_GetDragCursor,
  CoolBar_SetDragCursor,
  CoolBar_GetDragKind,
  CoolBar_SetDragKind,
  CoolBar_GetDragMode,
  CoolBar_SetDragMode,
  CoolBar_GetEdgeBorders,
  CoolBar_SetEdgeBorders,
  CoolBar_GetEdgeInner,
  CoolBar_SetEdgeInner,
  CoolBar_GetEdgeOuter,
  CoolBar_SetEdgeOuter,
  CoolBar_GetEnabled,
  CoolBar_SetEnabled,
  CoolBar_GetFixedSize,
  CoolBar_SetFixedSize,
  CoolBar_GetFixedOrder,
  CoolBar_SetFixedOrder,
  CoolBar_GetFont,
  CoolBar_SetFont,
  CoolBar_GetImages,
  CoolBar_SetImages,
  CoolBar_GetParentColor,
  CoolBar_SetParentColor,
  CoolBar_GetParentDoubleBuffered,
  CoolBar_SetParentDoubleBuffered,
  CoolBar_GetParentFont,
  CoolBar_SetParentFont,
  CoolBar_GetParentShowHint,
  CoolBar_SetParentShowHint,
  CoolBar_GetBitmap,
  CoolBar_SetBitmap,
  CoolBar_GetPopupMenu,
  CoolBar_SetPopupMenu,
  CoolBar_GetShowHint,
  CoolBar_SetShowHint,
  CoolBar_GetShowText,
  CoolBar_SetShowText,
  CoolBar_GetVertical,
  CoolBar_SetVertical,
  CoolBar_GetVisible,
  CoolBar_SetVisible,
  CoolBar_GetStyleElements,
  CoolBar_SetStyleElements,
  CoolBar_SetOnChange,
  CoolBar_SetOnClick,
  CoolBar_SetOnContextPopup,
  CoolBar_SetOnDblClick,
  CoolBar_SetOnDockDrop,
  CoolBar_SetOnDragDrop,
  CoolBar_SetOnDragOver,
  CoolBar_SetOnEndDock,
  CoolBar_SetOnEndDrag,
  CoolBar_SetOnGetSiteInfo,
  CoolBar_SetOnMouseDown,
  CoolBar_SetOnMouseEnter,
  CoolBar_SetOnMouseLeave,
  CoolBar_SetOnMouseMove,
  CoolBar_SetOnMouseUp,
  CoolBar_SetOnResize,
  CoolBar_SetOnStartDock,
  CoolBar_SetOnUnDock,
  CoolBar_GetDockClientCount,
  CoolBar_GetAlignDisabled,
  CoolBar_GetMouseInClient,
  CoolBar_GetVisibleDockClientCount,
  CoolBar_GetBrush,
  CoolBar_GetControlCount,
  CoolBar_GetHandle,
  CoolBar_GetParentWindow,
  CoolBar_SetParentWindow,
  CoolBar_GetTabOrder,
  CoolBar_SetTabOrder,
  CoolBar_GetTabStop,
  CoolBar_SetTabStop,
  CoolBar_GetUseDockManager,
  CoolBar_SetUseDockManager,
  CoolBar_GetAction,
  CoolBar_SetAction,
  CoolBar_GetBiDiMode,
  CoolBar_SetBiDiMode,
  CoolBar_GetBoundsRect,
  CoolBar_SetBoundsRect,
  CoolBar_GetClientHeight,
  CoolBar_SetClientHeight,
  CoolBar_GetClientOrigin,
  CoolBar_GetClientRect,
  CoolBar_GetClientWidth,
  CoolBar_SetClientWidth,
  CoolBar_GetControlState,
  CoolBar_SetControlState,
  CoolBar_GetControlStyle,
  CoolBar_SetControlStyle,
  CoolBar_GetExplicitLeft,
  CoolBar_GetExplicitTop,
  CoolBar_GetExplicitWidth,
  CoolBar_GetExplicitHeight,
  CoolBar_GetFloating,
  CoolBar_GetParent,
  CoolBar_SetParent,
  CoolBar_GetAlignWithMargins,
  CoolBar_SetAlignWithMargins,
  CoolBar_GetLeft,
  CoolBar_SetLeft,
  CoolBar_GetTop,
  CoolBar_SetTop,
  CoolBar_GetWidth,
  CoolBar_SetWidth,
  CoolBar_GetHeight,
  CoolBar_SetHeight,
  CoolBar_GetCursor,
  CoolBar_SetCursor,
  CoolBar_GetHint,
  CoolBar_SetHint,
  CoolBar_GetMargins,
  CoolBar_SetMargins,
  CoolBar_GetCustomHint,
  CoolBar_SetCustomHint,
  CoolBar_GetComponentCount,
  CoolBar_GetComponentIndex,
  CoolBar_SetComponentIndex,
  CoolBar_GetOwner,
  CoolBar_GetName,
  CoolBar_SetName,
  CoolBar_GetTag,
  CoolBar_SetTag,
  CoolBar_GetDockClients,
  CoolBar_GetControls,
  CoolBar_GetComponents,
  CoolBar_StaticClassType;

