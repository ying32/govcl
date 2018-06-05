
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ValueListEditor_Create(AOwner: TComponent): TValueListEditor; extdecl;
begin
  Result :=  TValueListEditor.Create(AOwner);
end;

procedure ValueListEditor_Free(AObj: TValueListEditor); extdecl;
begin
  AObj.Free;
end;

procedure ValueListEditor_Refresh(AObj: TValueListEditor); extdecl;
begin
  AObj.Refresh;
end;

function ValueListEditor_CanFocus(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure ValueListEditor_FlipChildren(AObj: TValueListEditor; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function ValueListEditor_Focused(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

function ValueListEditor_HandleAllocated(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ValueListEditor_Invalidate(AObj: TValueListEditor); extdecl;
begin
  AObj.Invalidate;
end;

procedure ValueListEditor_Realign(AObj: TValueListEditor); extdecl;
begin
  AObj.Realign;
end;

procedure ValueListEditor_Repaint(AObj: TValueListEditor); extdecl;
begin
  AObj.Repaint;
end;

procedure ValueListEditor_ScaleBy(AObj: TValueListEditor; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure ValueListEditor_SetBounds(AObj: TValueListEditor; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ValueListEditor_SetFocus(AObj: TValueListEditor); extdecl;
begin
  AObj.SetFocus;
end;

procedure ValueListEditor_Update(AObj: TValueListEditor); extdecl;
begin
  AObj.Update;
end;

procedure ValueListEditor_BringToFront(AObj: TValueListEditor); extdecl;
begin
  AObj.BringToFront;
end;

procedure ValueListEditor_ClientToScreen(AObj: TValueListEditor; var Point: TPoint; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure ValueListEditor_ClientToParent(AObj: TValueListEditor; var Point: TPoint; AParent: TWinControl; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function ValueListEditor_Dragging(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.Dragging;
end;

function ValueListEditor_HasParent(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ValueListEditor_Hide(AObj: TValueListEditor); extdecl;
begin
  AObj.Hide;
end;

function ValueListEditor_Perform(AObj: TValueListEditor; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ValueListEditor_ScreenToClient(AObj: TValueListEditor; var Point: TPoint; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure ValueListEditor_ParentToClient(AObj: TValueListEditor; var Point: TPoint; AParent: TWinControl; var Result: TPoint); extdecl;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure ValueListEditor_SendToBack(AObj: TValueListEditor); extdecl;
begin
  AObj.SendToBack;
end;

procedure ValueListEditor_Show(AObj: TValueListEditor); extdecl;
begin
  AObj.Show;
end;

function ValueListEditor_GetTextBuf(AObj: TValueListEditor; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ValueListEditor_GetTextLen(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function ValueListEditor_FindComponent(AObj: TValueListEditor; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ValueListEditor_GetNamePath(AObj: TValueListEditor): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ValueListEditor_Assign(AObj: TValueListEditor; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ValueListEditor_ClassType(AObj: TValueListEditor): TClass; extdecl;
begin
  Result :=  AObj.ClassType;
end;

function ValueListEditor_ClassName(AObj: TValueListEditor): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ValueListEditor_InstanceSize(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.InstanceSize;
end;

function ValueListEditor_InheritsFrom(AObj: TValueListEditor; AClass: TClass): LongBool; extdecl;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function ValueListEditor_Equals(AObj: TValueListEditor; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ValueListEditor_GetHashCode(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ValueListEditor_ToString(AObj: TValueListEditor): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ValueListEditor_GetRowCount(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.RowCount;
end;

function ValueListEditor_GetVisibleRowCount(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.VisibleRowCount;
end;

function ValueListEditor_GetAlign(AObj: TValueListEditor): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure ValueListEditor_SetAlign(AObj: TValueListEditor; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function ValueListEditor_GetAnchors(AObj: TValueListEditor): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure ValueListEditor_SetAnchors(AObj: TValueListEditor; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function ValueListEditor_GetBiDiMode(AObj: TValueListEditor): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ValueListEditor_SetBiDiMode(AObj: TValueListEditor; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function ValueListEditor_GetBorderStyle(AObj: TValueListEditor): TBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure ValueListEditor_SetBorderStyle(AObj: TValueListEditor; AValue: TBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function ValueListEditor_GetColor(AObj: TValueListEditor): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure ValueListEditor_SetColor(AObj: TValueListEditor; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function ValueListEditor_GetDoubleBuffered(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ValueListEditor_SetDoubleBuffered(AObj: TValueListEditor; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function ValueListEditor_GetDragCursor(AObj: TValueListEditor): TCursor; extdecl;
begin
  Result :=  AObj.DragCursor;
end;

procedure ValueListEditor_SetDragCursor(AObj: TValueListEditor; AValue: TCursor); extdecl;
begin
  AObj.DragCursor := AValue;
end;

function ValueListEditor_GetDragKind(AObj: TValueListEditor): TDragKind; extdecl;
begin
  Result :=  AObj.DragKind;
end;

procedure ValueListEditor_SetDragKind(AObj: TValueListEditor; AValue: TDragKind); extdecl;
begin
  AObj.DragKind := AValue;
end;

function ValueListEditor_GetDragMode(AObj: TValueListEditor): TDragMode; extdecl;
begin
  Result :=  AObj.DragMode;
end;

procedure ValueListEditor_SetDragMode(AObj: TValueListEditor; AValue: TDragMode); extdecl;
begin
  AObj.DragMode := AValue;
end;

function ValueListEditor_GetEnabled(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure ValueListEditor_SetEnabled(AObj: TValueListEditor; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function ValueListEditor_GetFont(AObj: TValueListEditor): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure ValueListEditor_SetFont(AObj: TValueListEditor; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function ValueListEditor_GetOptions(AObj: TValueListEditor): TGridOptions; extdecl;
begin
  Result :=  AObj.Options;
end;

procedure ValueListEditor_SetOptions(AObj: TValueListEditor; AValue: TGridOptions); extdecl;
begin
  AObj.Options := AValue;
end;

function ValueListEditor_GetParentColor(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure ValueListEditor_SetParentColor(AObj: TValueListEditor; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function ValueListEditor_GetParentFont(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure ValueListEditor_SetParentFont(AObj: TValueListEditor; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function ValueListEditor_GetParentShowHint(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ValueListEditor_SetParentShowHint(AObj: TValueListEditor; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function ValueListEditor_GetPopupMenu(AObj: TValueListEditor): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ValueListEditor_SetPopupMenu(AObj: TValueListEditor; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function ValueListEditor_GetScrollBars(AObj: TValueListEditor): TScrollStyle; extdecl;
begin
  Result :=  AObj.ScrollBars;
end;

procedure ValueListEditor_SetScrollBars(AObj: TValueListEditor; AValue: TScrollStyle); extdecl;
begin
  AObj.ScrollBars := AValue;
end;

function ValueListEditor_GetShowHint(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure ValueListEditor_SetShowHint(AObj: TValueListEditor; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function ValueListEditor_GetStrings(AObj: TValueListEditor): TStrings; extdecl;
begin
  Result :=  AObj.Strings as TValueListStrings;
end;

procedure ValueListEditor_SetStrings(AObj: TValueListEditor; AValue: TStrings); extdecl;
begin
  AObj.Strings := AValue as TValueListStrings;
end;

function ValueListEditor_GetTabOrder(AObj: TValueListEditor): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure ValueListEditor_SetTabOrder(AObj: TValueListEditor; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function ValueListEditor_GetVisible(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure ValueListEditor_SetVisible(AObj: TValueListEditor; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure ValueListEditor_SetOnClick(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnContextPopup(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnDblClick(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnDragDrop(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnDragOver(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnEndDock(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnEndDrag(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnEnter(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnExit(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnGetEditMask(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnGetEditText(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnKeyDown(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnKeyPress(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnKeyUp(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnMouseDown(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnMouseEnter(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnMouseLeave(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnMouseMove(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnMouseUp(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnMouseWheelDown(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnMouseWheelUp(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnRowMoved(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnSelectCell(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnSetEditText(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnStartDock(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

procedure ValueListEditor_SetOnTopLeftChanged(AObj: TValueListEditor; AEventId: NativeUInt); extdecl;
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

function ValueListEditor_GetCanvas(AObj: TValueListEditor): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function ValueListEditor_GetTabStop(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure ValueListEditor_SetTabStop(AObj: TValueListEditor; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function ValueListEditor_GetDockSite(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.DockSite;
end;

procedure ValueListEditor_SetDockSite(AObj: TValueListEditor; AValue: LongBool); extdecl;
begin
  AObj.DockSite := AValue;
end;

function ValueListEditor_GetBrush(AObj: TValueListEditor): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function ValueListEditor_GetControlCount(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function ValueListEditor_GetHandle(AObj: TValueListEditor): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function ValueListEditor_GetParentWindow(AObj: TValueListEditor): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ValueListEditor_SetParentWindow(AObj: TValueListEditor; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function ValueListEditor_GetUseDockManager(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.UseDockManager;
end;

procedure ValueListEditor_SetUseDockManager(AObj: TValueListEditor; AValue: LongBool); extdecl;
begin
  AObj.UseDockManager := AValue;
end;

function ValueListEditor_GetAction(AObj: TValueListEditor): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure ValueListEditor_SetAction(AObj: TValueListEditor; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure ValueListEditor_GetBoundsRect(AObj: TValueListEditor; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ValueListEditor_SetBoundsRect(AObj: TValueListEditor; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function ValueListEditor_GetClientHeight(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ValueListEditor_SetClientHeight(AObj: TValueListEditor; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure ValueListEditor_GetClientRect(AObj: TValueListEditor; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function ValueListEditor_GetClientWidth(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ValueListEditor_SetClientWidth(AObj: TValueListEditor; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function ValueListEditor_GetFloating(AObj: TValueListEditor): LongBool; extdecl;
begin
  Result :=  AObj.Floating;
end;

function ValueListEditor_GetParent(AObj: TValueListEditor): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure ValueListEditor_SetParent(AObj: TValueListEditor; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function ValueListEditor_GetLeft(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ValueListEditor_SetLeft(AObj: TValueListEditor; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ValueListEditor_GetTop(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ValueListEditor_SetTop(AObj: TValueListEditor; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ValueListEditor_GetWidth(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ValueListEditor_SetWidth(AObj: TValueListEditor; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function ValueListEditor_GetHeight(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ValueListEditor_SetHeight(AObj: TValueListEditor; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ValueListEditor_GetCursor(AObj: TValueListEditor): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure ValueListEditor_SetCursor(AObj: TValueListEditor; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function ValueListEditor_GetHint(AObj: TValueListEditor): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ValueListEditor_SetHint(AObj: TValueListEditor; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function ValueListEditor_GetComponentCount(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ValueListEditor_GetComponentIndex(AObj: TValueListEditor): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ValueListEditor_SetComponentIndex(AObj: TValueListEditor; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ValueListEditor_GetOwner(AObj: TValueListEditor): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ValueListEditor_GetName(AObj: TValueListEditor): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ValueListEditor_SetName(AObj: TValueListEditor; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ValueListEditor_GetTag(AObj: TValueListEditor): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ValueListEditor_SetTag(AObj: TValueListEditor; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ValueListEditor_GetValues(AObj: TValueListEditor; Key: PChar): PChar; extdecl;
begin
  Result :=  PChar(AObj.Values[Key]);
end;

procedure ValueListEditor_SetValues(AObj: TValueListEditor; Key: string; AValue: PChar); extdecl;
begin
  AObj.Values[Key] := AValue;
end;

function ValueListEditor_GetControls(AObj: TValueListEditor; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function ValueListEditor_GetComponents(AObj: TValueListEditor; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;

function ValueListEditor_StaticClassType: TClass; extdecl;
begin
  Result := TValueListEditor;
end;

exports
  ValueListEditor_Create,
  ValueListEditor_Free,
  ValueListEditor_Refresh,
  ValueListEditor_CanFocus,
  ValueListEditor_FlipChildren,
  ValueListEditor_Focused,
  ValueListEditor_HandleAllocated,
  ValueListEditor_Invalidate,
  ValueListEditor_Realign,
  ValueListEditor_Repaint,
  ValueListEditor_ScaleBy,
  ValueListEditor_SetBounds,
  ValueListEditor_SetFocus,
  ValueListEditor_Update,
  ValueListEditor_BringToFront,
  ValueListEditor_ClientToScreen,
  ValueListEditor_ClientToParent,
  ValueListEditor_Dragging,
  ValueListEditor_HasParent,
  ValueListEditor_Hide,
  ValueListEditor_Perform,
  ValueListEditor_ScreenToClient,
  ValueListEditor_ParentToClient,
  ValueListEditor_SendToBack,
  ValueListEditor_Show,
  ValueListEditor_GetTextBuf,
  ValueListEditor_GetTextLen,
  ValueListEditor_FindComponent,
  ValueListEditor_GetNamePath,
  ValueListEditor_Assign,
  ValueListEditor_ClassType,
  ValueListEditor_ClassName,
  ValueListEditor_InstanceSize,
  ValueListEditor_InheritsFrom,
  ValueListEditor_Equals,
  ValueListEditor_GetHashCode,
  ValueListEditor_ToString,
  ValueListEditor_GetRowCount,
  ValueListEditor_GetVisibleRowCount,
  ValueListEditor_GetAlign,
  ValueListEditor_SetAlign,
  ValueListEditor_GetAnchors,
  ValueListEditor_SetAnchors,
  ValueListEditor_GetBiDiMode,
  ValueListEditor_SetBiDiMode,
  ValueListEditor_GetBorderStyle,
  ValueListEditor_SetBorderStyle,
  ValueListEditor_GetColor,
  ValueListEditor_SetColor,
  ValueListEditor_GetDoubleBuffered,
  ValueListEditor_SetDoubleBuffered,
  ValueListEditor_GetDragCursor,
  ValueListEditor_SetDragCursor,
  ValueListEditor_GetDragKind,
  ValueListEditor_SetDragKind,
  ValueListEditor_GetDragMode,
  ValueListEditor_SetDragMode,
  ValueListEditor_GetEnabled,
  ValueListEditor_SetEnabled,
  ValueListEditor_GetFont,
  ValueListEditor_SetFont,
  ValueListEditor_GetOptions,
  ValueListEditor_SetOptions,
  ValueListEditor_GetParentColor,
  ValueListEditor_SetParentColor,
  ValueListEditor_GetParentFont,
  ValueListEditor_SetParentFont,
  ValueListEditor_GetParentShowHint,
  ValueListEditor_SetParentShowHint,
  ValueListEditor_GetPopupMenu,
  ValueListEditor_SetPopupMenu,
  ValueListEditor_GetScrollBars,
  ValueListEditor_SetScrollBars,
  ValueListEditor_GetShowHint,
  ValueListEditor_SetShowHint,
  ValueListEditor_GetStrings,
  ValueListEditor_SetStrings,
  ValueListEditor_GetTabOrder,
  ValueListEditor_SetTabOrder,
  ValueListEditor_GetVisible,
  ValueListEditor_SetVisible,
  ValueListEditor_SetOnClick,
  ValueListEditor_SetOnContextPopup,
  ValueListEditor_SetOnDblClick,
  ValueListEditor_SetOnDragDrop,
  ValueListEditor_SetOnDragOver,
  ValueListEditor_SetOnEndDock,
  ValueListEditor_SetOnEndDrag,
  ValueListEditor_SetOnEnter,
  ValueListEditor_SetOnExit,
  ValueListEditor_SetOnGetEditMask,
  ValueListEditor_SetOnGetEditText,
  ValueListEditor_SetOnKeyDown,
  ValueListEditor_SetOnKeyPress,
  ValueListEditor_SetOnKeyUp,
  ValueListEditor_SetOnMouseDown,
  ValueListEditor_SetOnMouseEnter,
  ValueListEditor_SetOnMouseLeave,
  ValueListEditor_SetOnMouseMove,
  ValueListEditor_SetOnMouseUp,
  ValueListEditor_SetOnMouseWheelDown,
  ValueListEditor_SetOnMouseWheelUp,
  ValueListEditor_SetOnRowMoved,
  ValueListEditor_SetOnSelectCell,
  ValueListEditor_SetOnSetEditText,
  ValueListEditor_SetOnStartDock,
  ValueListEditor_SetOnTopLeftChanged,
  ValueListEditor_GetCanvas,
  ValueListEditor_GetTabStop,
  ValueListEditor_SetTabStop,
  ValueListEditor_GetDockSite,
  ValueListEditor_SetDockSite,
  ValueListEditor_GetBrush,
  ValueListEditor_GetControlCount,
  ValueListEditor_GetHandle,
  ValueListEditor_GetParentWindow,
  ValueListEditor_SetParentWindow,
  ValueListEditor_GetUseDockManager,
  ValueListEditor_SetUseDockManager,
  ValueListEditor_GetAction,
  ValueListEditor_SetAction,
  ValueListEditor_GetBoundsRect,
  ValueListEditor_SetBoundsRect,
  ValueListEditor_GetClientHeight,
  ValueListEditor_SetClientHeight,
  ValueListEditor_GetClientRect,
  ValueListEditor_GetClientWidth,
  ValueListEditor_SetClientWidth,
  ValueListEditor_GetFloating,
  ValueListEditor_GetParent,
  ValueListEditor_SetParent,
  ValueListEditor_GetLeft,
  ValueListEditor_SetLeft,
  ValueListEditor_GetTop,
  ValueListEditor_SetTop,
  ValueListEditor_GetWidth,
  ValueListEditor_SetWidth,
  ValueListEditor_GetHeight,
  ValueListEditor_SetHeight,
  ValueListEditor_GetCursor,
  ValueListEditor_SetCursor,
  ValueListEditor_GetHint,
  ValueListEditor_SetHint,
  ValueListEditor_GetComponentCount,
  ValueListEditor_GetComponentIndex,
  ValueListEditor_SetComponentIndex,
  ValueListEditor_GetOwner,
  ValueListEditor_GetName,
  ValueListEditor_SetName,
  ValueListEditor_GetTag,
  ValueListEditor_SetTag,
  ValueListEditor_GetValues,
  ValueListEditor_SetValues,
  ValueListEditor_GetControls,
  ValueListEditor_GetComponents,
  ValueListEditor_StaticClassType;

