
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function DrawGrid_Create(AOwner: TComponent): TDrawGrid; extdecl;
begin
  Result :=  TDrawGrid.Create(AOwner);
end;

procedure DrawGrid_Free(AObj: TDrawGrid); extdecl;
begin
  AObj.Free;
end;

function DrawGrid_CanFocus(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure DrawGrid_FlipChildren(AObj: TDrawGrid; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function DrawGrid_Focused(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

function DrawGrid_HandleAllocated(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure DrawGrid_Invalidate(AObj: TDrawGrid); extdecl;
begin
  AObj.Invalidate;
end;

procedure DrawGrid_Realign(AObj: TDrawGrid); extdecl;
begin
  AObj.Realign;
end;

procedure DrawGrid_Repaint(AObj: TDrawGrid); extdecl;
begin
  AObj.Repaint;
end;

procedure DrawGrid_ScaleBy(AObj: TDrawGrid; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure DrawGrid_SetBounds(AObj: TDrawGrid; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure DrawGrid_SetFocus(AObj: TDrawGrid); extdecl;
begin
  AObj.SetFocus;
end;

procedure DrawGrid_Update(AObj: TDrawGrid); extdecl;
begin
  AObj.Update;
end;

procedure DrawGrid_BringToFront(AObj: TDrawGrid); extdecl;
begin
  AObj.BringToFront;
end;

procedure DrawGrid_ClientToScreen(AObj: TDrawGrid; var Point: TPoint; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure DrawGrid_ClientToParent(AObj: TDrawGrid; var Point: TPoint; AParent: TWinControl; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function DrawGrid_Dragging(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.Dragging;
end;

function DrawGrid_HasParent(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure DrawGrid_Hide(AObj: TDrawGrid); extdecl;
begin
  AObj.Hide;
end;

function DrawGrid_Perform(AObj: TDrawGrid; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure DrawGrid_Refresh(AObj: TDrawGrid); extdecl;
begin
  AObj.Refresh;
end;

procedure DrawGrid_ScreenToClient(AObj: TDrawGrid; var Point: TPoint; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure DrawGrid_ParentToClient(AObj: TDrawGrid; var Point: TPoint; AParent: TWinControl; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure DrawGrid_SendToBack(AObj: TDrawGrid); extdecl;
begin
  AObj.SendToBack;
end;

procedure DrawGrid_Show(AObj: TDrawGrid); extdecl;
begin
  AObj.Show;
end;

function DrawGrid_GetTextBuf(AObj: TDrawGrid; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function DrawGrid_GetTextLen(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function DrawGrid_FindComponent(AObj: TDrawGrid; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function DrawGrid_GetNamePath(AObj: TDrawGrid): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure DrawGrid_Assign(AObj: TDrawGrid; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function DrawGrid_ClassType(AObj: TDrawGrid): TClass; extdecl;
begin
  Result :=  AObj.ClassType;
end;

function DrawGrid_ClassName(AObj: TDrawGrid): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function DrawGrid_InstanceSize(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.InstanceSize;
end;

function DrawGrid_InheritsFrom(AObj: TDrawGrid; AClass: TClass): LongBool; extdecl;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function DrawGrid_Equals(AObj: TDrawGrid; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function DrawGrid_GetHashCode(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function DrawGrid_ToString(AObj: TDrawGrid): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function DrawGrid_GetAlign(AObj: TDrawGrid): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure DrawGrid_SetAlign(AObj: TDrawGrid; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function DrawGrid_GetAnchors(AObj: TDrawGrid): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure DrawGrid_SetAnchors(AObj: TDrawGrid; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function DrawGrid_GetBiDiMode(AObj: TDrawGrid): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure DrawGrid_SetBiDiMode(AObj: TDrawGrid; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function DrawGrid_GetBorderStyle(AObj: TDrawGrid): TBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure DrawGrid_SetBorderStyle(AObj: TDrawGrid; AValue: TBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function DrawGrid_GetColor(AObj: TDrawGrid): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure DrawGrid_SetColor(AObj: TDrawGrid; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function DrawGrid_GetDoubleBuffered(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure DrawGrid_SetDoubleBuffered(AObj: TDrawGrid; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function DrawGrid_GetDragCursor(AObj: TDrawGrid): TCursor; extdecl;
begin
  Result :=  AObj.DragCursor;
end;

procedure DrawGrid_SetDragCursor(AObj: TDrawGrid; AValue: TCursor); extdecl;
begin
  AObj.DragCursor := AValue;
end;

function DrawGrid_GetDragKind(AObj: TDrawGrid): TDragKind; extdecl;
begin
  Result :=  AObj.DragKind;
end;

procedure DrawGrid_SetDragKind(AObj: TDrawGrid; AValue: TDragKind); extdecl;
begin
  AObj.DragKind := AValue;
end;

function DrawGrid_GetDragMode(AObj: TDrawGrid): TDragMode; extdecl;
begin
  Result :=  AObj.DragMode;
end;

procedure DrawGrid_SetDragMode(AObj: TDrawGrid; AValue: TDragMode); extdecl;
begin
  AObj.DragMode := AValue;
end;

function DrawGrid_GetEnabled(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure DrawGrid_SetEnabled(AObj: TDrawGrid; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function DrawGrid_GetRowCount(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.RowCount;
end;

procedure DrawGrid_SetRowCount(AObj: TDrawGrid; AValue: Integer); extdecl;
begin
  AObj.RowCount := AValue;
end;

function DrawGrid_GetFont(AObj: TDrawGrid): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure DrawGrid_SetFont(AObj: TDrawGrid; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function DrawGrid_GetOptions(AObj: TDrawGrid): TGridOptions; extdecl;
begin
  Result :=  AObj.Options;
end;

procedure DrawGrid_SetOptions(AObj: TDrawGrid; AValue: TGridOptions); extdecl;
begin
  AObj.Options := AValue;
end;

function DrawGrid_GetParentColor(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure DrawGrid_SetParentColor(AObj: TDrawGrid; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function DrawGrid_GetParentFont(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure DrawGrid_SetParentFont(AObj: TDrawGrid; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function DrawGrid_GetParentShowHint(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure DrawGrid_SetParentShowHint(AObj: TDrawGrid; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function DrawGrid_GetPopupMenu(AObj: TDrawGrid): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure DrawGrid_SetPopupMenu(AObj: TDrawGrid; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function DrawGrid_GetScrollBars(AObj: TDrawGrid): TScrollStyle; extdecl;
begin
  Result :=  AObj.ScrollBars;
end;

procedure DrawGrid_SetScrollBars(AObj: TDrawGrid; AValue: TScrollStyle); extdecl;
begin
  AObj.ScrollBars := AValue;
end;

function DrawGrid_GetShowHint(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure DrawGrid_SetShowHint(AObj: TDrawGrid; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function DrawGrid_GetTabOrder(AObj: TDrawGrid): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure DrawGrid_SetTabOrder(AObj: TDrawGrid; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function DrawGrid_GetVisible(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure DrawGrid_SetVisible(AObj: TDrawGrid; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function DrawGrid_GetVisibleRowCount(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.VisibleRowCount;
end;

procedure DrawGrid_SetOnClick(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
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

procedure DrawGrid_SetOnColumnMoved(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnColRowMoved := nil;
    EventClass.Remove(AObj, geColumnMoved);
    Exit;
  end;
  AObj.OnColRowMoved := @EventClass.OnColumnMoved;
  EventClass.Add(AObj, geColumnMoved, AEventId);
end;

procedure DrawGrid_SetOnContextPopup(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnContextPopup := nil;
    EventClass.Remove(AObj, geContextPopup);
    Exit;
  end;
  AObj.OnContextPopup := @EventClass.OnContextPopup;
  EventClass.Add(AObj, geContextPopup, AEventId);
end;

procedure DrawGrid_SetOnDblClick(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
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

procedure DrawGrid_SetOnDragDrop(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragDrop := nil;
    EventClass.Remove(AObj, geDragDrop);
    Exit;
  end;
  AObj.OnDragDrop := @EventClass.OnDragDrop;
  EventClass.Add(AObj, geDragDrop, AEventId);
end;

procedure DrawGrid_SetOnDragOver(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragOver := nil;
    EventClass.Remove(AObj, geDragOver);
    Exit;
  end;
  AObj.OnDragOver := @EventClass.OnDragOver;
  EventClass.Add(AObj, geDragOver, AEventId);
end;

procedure DrawGrid_SetOnEndDock(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDock := nil;
    EventClass.Remove(AObj, geEndDock);
    Exit;
  end;
  AObj.OnEndDock := @EventClass.OnEndDock;
  EventClass.Add(AObj, geEndDock, AEventId);
end;

procedure DrawGrid_SetOnEndDrag(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDrag := nil;
    EventClass.Remove(AObj, geEndDrag);
    Exit;
  end;
  AObj.OnEndDrag := @EventClass.OnEndDrag;
  EventClass.Add(AObj, geEndDrag, AEventId);
end;

procedure DrawGrid_SetOnEnter(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
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

procedure DrawGrid_SetOnExit(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
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

procedure DrawGrid_SetOnGetEditMask(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetEditMask := nil;
    EventClass.Remove(AObj, geGetEditMask);
    Exit;
  end;
  AObj.OnGetEditMask := @EventClass.OnGetEditMask;
  EventClass.Add(AObj, geGetEditMask, AEventId);
end;

procedure DrawGrid_SetOnGetEditText(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetEditText := nil;
    EventClass.Remove(AObj, geGetEditText);
    Exit;
  end;
  AObj.OnGetEditText := @EventClass.OnGetEditText;
  EventClass.Add(AObj, geGetEditText, AEventId);
end;

procedure DrawGrid_SetOnKeyDown(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyDown := nil;
    EventClass.Remove(AObj, geKeyDown);
    Exit;
  end;
  AObj.OnKeyDown := @EventClass.OnKeyDown;
  EventClass.Add(AObj, geKeyDown, AEventId);
end;

procedure DrawGrid_SetOnKeyPress(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyPress := nil;
    EventClass.Remove(AObj, geKeyPress);
    Exit;
  end;
  AObj.OnKeyPress := @EventClass.OnKeyPress;
  EventClass.Add(AObj, geKeyPress, AEventId);
end;

procedure DrawGrid_SetOnKeyUp(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyUp := nil;
    EventClass.Remove(AObj, geKeyUp);
    Exit;
  end;
  AObj.OnKeyUp := @EventClass.OnKeyUp;
  EventClass.Add(AObj, geKeyUp, AEventId);
end;

procedure DrawGrid_SetOnMouseDown(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
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

procedure DrawGrid_SetOnMouseEnter(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
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

procedure DrawGrid_SetOnMouseLeave(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
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

procedure DrawGrid_SetOnMouseMove(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
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

procedure DrawGrid_SetOnMouseUp(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
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

procedure DrawGrid_SetOnMouseWheelDown(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheelDown := nil;
    EventClass.Remove(AObj, geMouseWheelDown);
    Exit;
  end;
  AObj.OnMouseWheelDown := @EventClass.OnMouseWheelDown;
  EventClass.Add(AObj, geMouseWheelDown, AEventId);
end;

procedure DrawGrid_SetOnMouseWheelUp(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheelUp := nil;
    EventClass.Remove(AObj, geMouseWheelUp);
    Exit;
  end;
  AObj.OnMouseWheelUp := @EventClass.OnMouseWheelUp;
  EventClass.Add(AObj, geMouseWheelUp, AEventId);
end;

procedure DrawGrid_SetOnRowMoved(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnColRowMoved := nil;
    EventClass.Remove(AObj, geRowMoved);
    Exit;
  end;
  AObj.OnColRowMoved := @EventClass.OnRowMoved;
  EventClass.Add(AObj, geRowMoved, AEventId);
end;

procedure DrawGrid_SetOnSelectCell(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnSelectCell := nil;
    EventClass.Remove(AObj, geSelectCell);
    Exit;
  end;
  AObj.OnSelectCell := @EventClass.OnSelectCell;
  EventClass.Add(AObj, geSelectCell, AEventId);
end;

procedure DrawGrid_SetOnSetEditText(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnSetEditText := nil;
    EventClass.Remove(AObj, geSetEditText);
    Exit;
  end;
  AObj.OnSetEditText := @EventClass.OnSetEditText;
  EventClass.Add(AObj, geSetEditText, AEventId);
end;

procedure DrawGrid_SetOnStartDock(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnStartDock := nil;
    EventClass.Remove(AObj, geStartDock);
    Exit;
  end;
  AObj.OnStartDock := @EventClass.OnStartDock;
  EventClass.Add(AObj, geStartDock, AEventId);
end;

procedure DrawGrid_SetOnTopLeftChanged(AObj: TDrawGrid; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnTopLeftChanged := nil;
    EventClass.Remove(AObj, geTopLeftChanged);
    Exit;
  end;
  AObj.OnTopLeftChanged := @EventClass.OnTopLeftChanged;
  EventClass.Add(AObj, geTopLeftChanged, AEventId);
end;

function DrawGrid_GetCanvas(AObj: TDrawGrid): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function DrawGrid_GetTabStop(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure DrawGrid_SetTabStop(AObj: TDrawGrid; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function DrawGrid_GetDockSite(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.DockSite;
end;

procedure DrawGrid_SetDockSite(AObj: TDrawGrid; AValue: LongBool); extdecl;
begin
  AObj.DockSite := AValue;
end;

function DrawGrid_GetBrush(AObj: TDrawGrid): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function DrawGrid_GetControlCount(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function DrawGrid_GetHandle(AObj: TDrawGrid): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function DrawGrid_GetParentWindow(AObj: TDrawGrid): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure DrawGrid_SetParentWindow(AObj: TDrawGrid; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function DrawGrid_GetUseDockManager(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.UseDockManager;
end;

procedure DrawGrid_SetUseDockManager(AObj: TDrawGrid; AValue: LongBool); extdecl;
begin
  AObj.UseDockManager := AValue;
end;

function DrawGrid_GetAction(AObj: TDrawGrid): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure DrawGrid_SetAction(AObj: TDrawGrid; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure DrawGrid_GetBoundsRect(AObj: TDrawGrid; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure DrawGrid_SetBoundsRect(AObj: TDrawGrid; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function DrawGrid_GetClientHeight(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure DrawGrid_SetClientHeight(AObj: TDrawGrid; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure DrawGrid_GetClientRect(AObj: TDrawGrid; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function DrawGrid_GetClientWidth(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure DrawGrid_SetClientWidth(AObj: TDrawGrid; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function DrawGrid_GetFloating(AObj: TDrawGrid): LongBool; extdecl;
begin
  Result :=  AObj.Floating;
end;

function DrawGrid_GetParent(AObj: TDrawGrid): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure DrawGrid_SetParent(AObj: TDrawGrid; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function DrawGrid_GetLeft(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure DrawGrid_SetLeft(AObj: TDrawGrid; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function DrawGrid_GetTop(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure DrawGrid_SetTop(AObj: TDrawGrid; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function DrawGrid_GetWidth(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure DrawGrid_SetWidth(AObj: TDrawGrid; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function DrawGrid_GetHeight(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure DrawGrid_SetHeight(AObj: TDrawGrid; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function DrawGrid_GetCursor(AObj: TDrawGrid): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure DrawGrid_SetCursor(AObj: TDrawGrid; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function DrawGrid_GetHint(AObj: TDrawGrid): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure DrawGrid_SetHint(AObj: TDrawGrid; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function DrawGrid_GetComponentCount(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function DrawGrid_GetComponentIndex(AObj: TDrawGrid): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure DrawGrid_SetComponentIndex(AObj: TDrawGrid; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function DrawGrid_GetOwner(AObj: TDrawGrid): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function DrawGrid_GetName(AObj: TDrawGrid): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure DrawGrid_SetName(AObj: TDrawGrid; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function DrawGrid_GetTag(AObj: TDrawGrid): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure DrawGrid_SetTag(AObj: TDrawGrid; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function DrawGrid_GetControls(AObj: TDrawGrid; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function DrawGrid_GetComponents(AObj: TDrawGrid; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;

function DrawGrid_StaticClassType: TClass; extdecl;
begin
  Result := TDrawGrid;
end;

exports
  DrawGrid_Create,
  DrawGrid_Free,
  DrawGrid_CanFocus,
  DrawGrid_FlipChildren,
  DrawGrid_Focused,
  DrawGrid_HandleAllocated,
  DrawGrid_Invalidate,
  DrawGrid_Realign,
  DrawGrid_Repaint,
  DrawGrid_ScaleBy,
  DrawGrid_SetBounds,
  DrawGrid_SetFocus,
  DrawGrid_Update,
  DrawGrid_BringToFront,
  DrawGrid_ClientToScreen,
  DrawGrid_ClientToParent,
  DrawGrid_Dragging,
  DrawGrid_HasParent,
  DrawGrid_Hide,
  DrawGrid_Perform,
  DrawGrid_Refresh,
  DrawGrid_ScreenToClient,
  DrawGrid_ParentToClient,
  DrawGrid_SendToBack,
  DrawGrid_Show,
  DrawGrid_GetTextBuf,
  DrawGrid_GetTextLen,
  DrawGrid_FindComponent,
  DrawGrid_GetNamePath,
  DrawGrid_Assign,
  DrawGrid_ClassType,
  DrawGrid_ClassName,
  DrawGrid_InstanceSize,
  DrawGrid_InheritsFrom,
  DrawGrid_Equals,
  DrawGrid_GetHashCode,
  DrawGrid_ToString,
  DrawGrid_GetAlign,
  DrawGrid_SetAlign,
  DrawGrid_GetAnchors,
  DrawGrid_SetAnchors,
  DrawGrid_GetBiDiMode,
  DrawGrid_SetBiDiMode,
  DrawGrid_GetBorderStyle,
  DrawGrid_SetBorderStyle,
  DrawGrid_GetColor,
  DrawGrid_SetColor,
  DrawGrid_GetDoubleBuffered,
  DrawGrid_SetDoubleBuffered,
  DrawGrid_GetDragCursor,
  DrawGrid_SetDragCursor,
  DrawGrid_GetDragKind,
  DrawGrid_SetDragKind,
  DrawGrid_GetDragMode,
  DrawGrid_SetDragMode,
  DrawGrid_GetEnabled,
  DrawGrid_SetEnabled,
  DrawGrid_GetRowCount,
  DrawGrid_SetRowCount,
  DrawGrid_GetFont,
  DrawGrid_SetFont,
  DrawGrid_GetOptions,
  DrawGrid_SetOptions,
  DrawGrid_GetParentColor,
  DrawGrid_SetParentColor,
  DrawGrid_GetParentFont,
  DrawGrid_SetParentFont,
  DrawGrid_GetParentShowHint,
  DrawGrid_SetParentShowHint,
  DrawGrid_GetPopupMenu,
  DrawGrid_SetPopupMenu,
  DrawGrid_GetScrollBars,
  DrawGrid_SetScrollBars,
  DrawGrid_GetShowHint,
  DrawGrid_SetShowHint,
  DrawGrid_GetTabOrder,
  DrawGrid_SetTabOrder,
  DrawGrid_GetVisible,
  DrawGrid_SetVisible,
  DrawGrid_GetVisibleRowCount,
  DrawGrid_SetOnClick,
  DrawGrid_SetOnColumnMoved,
  DrawGrid_SetOnContextPopup,
  DrawGrid_SetOnDblClick,
  DrawGrid_SetOnDragDrop,
  DrawGrid_SetOnDragOver,
  DrawGrid_SetOnEndDock,
  DrawGrid_SetOnEndDrag,
  DrawGrid_SetOnEnter,
  DrawGrid_SetOnExit,
  DrawGrid_SetOnGetEditMask,
  DrawGrid_SetOnGetEditText,
  DrawGrid_SetOnKeyDown,
  DrawGrid_SetOnKeyPress,
  DrawGrid_SetOnKeyUp,
  DrawGrid_SetOnMouseDown,
  DrawGrid_SetOnMouseEnter,
  DrawGrid_SetOnMouseLeave,
  DrawGrid_SetOnMouseMove,
  DrawGrid_SetOnMouseUp,
  DrawGrid_SetOnMouseWheelDown,
  DrawGrid_SetOnMouseWheelUp,
  DrawGrid_SetOnRowMoved,
  DrawGrid_SetOnSelectCell,
  DrawGrid_SetOnSetEditText,
  DrawGrid_SetOnStartDock,
  DrawGrid_SetOnTopLeftChanged,
  DrawGrid_GetCanvas,
  DrawGrid_GetTabStop,
  DrawGrid_SetTabStop,
  DrawGrid_GetDockSite,
  DrawGrid_SetDockSite,
  DrawGrid_GetBrush,
  DrawGrid_GetControlCount,
  DrawGrid_GetHandle,
  DrawGrid_GetParentWindow,
  DrawGrid_SetParentWindow,
  DrawGrid_GetUseDockManager,
  DrawGrid_SetUseDockManager,
  DrawGrid_GetAction,
  DrawGrid_SetAction,
  DrawGrid_GetBoundsRect,
  DrawGrid_SetBoundsRect,
  DrawGrid_GetClientHeight,
  DrawGrid_SetClientHeight,
  DrawGrid_GetClientRect,
  DrawGrid_GetClientWidth,
  DrawGrid_SetClientWidth,
  DrawGrid_GetFloating,
  DrawGrid_GetParent,
  DrawGrid_SetParent,
  DrawGrid_GetLeft,
  DrawGrid_SetLeft,
  DrawGrid_GetTop,
  DrawGrid_SetTop,
  DrawGrid_GetWidth,
  DrawGrid_SetWidth,
  DrawGrid_GetHeight,
  DrawGrid_SetHeight,
  DrawGrid_GetCursor,
  DrawGrid_SetCursor,
  DrawGrid_GetHint,
  DrawGrid_SetHint,
  DrawGrid_GetComponentCount,
  DrawGrid_GetComponentIndex,
  DrawGrid_SetComponentIndex,
  DrawGrid_GetOwner,
  DrawGrid_GetName,
  DrawGrid_SetName,
  DrawGrid_GetTag,
  DrawGrid_SetTag,
  DrawGrid_GetControls,
  DrawGrid_GetComponents,
  DrawGrid_StaticClassType;

