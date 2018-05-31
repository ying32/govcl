
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Memo_Create(AOwner: TComponent): TMemo; extdecl;
begin
  Result :=  TMemo.Create(AOwner);
end;

procedure Memo_Free(AObj: TMemo); extdecl;
begin
  AObj.Free;
end;

procedure Memo_Clear(AObj: TMemo); extdecl;
begin
  AObj.Clear;
end;

procedure Memo_ClearSelection(AObj: TMemo); extdecl;
begin
  AObj.ClearSelection;
end;

procedure Memo_CopyToClipboard(AObj: TMemo); extdecl;
begin
  AObj.CopyToClipboard;
end;

procedure Memo_CutToClipboard(AObj: TMemo); extdecl;
begin
  AObj.CutToClipboard;
end;

procedure Memo_PasteFromClipboard(AObj: TMemo); extdecl;
begin
  AObj.PasteFromClipboard;
end;

procedure Memo_Undo(AObj: TMemo); extdecl;
begin
  AObj.Undo;
end;

procedure Memo_SelectAll(AObj: TMemo); extdecl;
begin
  AObj.SelectAll;
end;

function Memo_CanFocus(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure Memo_FlipChildren(AObj: TMemo; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function Memo_Focused(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

function Memo_HandleAllocated(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure Memo_Invalidate(AObj: TMemo); extdecl;
begin
  AObj.Invalidate;
end;

procedure Memo_Realign(AObj: TMemo); extdecl;
begin
  AObj.Realign;
end;

procedure Memo_Repaint(AObj: TMemo); extdecl;
begin
  AObj.Repaint;
end;

procedure Memo_ScaleBy(AObj: TMemo; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure Memo_SetBounds(AObj: TMemo; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Memo_SetFocus(AObj: TMemo); extdecl;
begin
  AObj.SetFocus;
end;

procedure Memo_Update(AObj: TMemo); extdecl;
begin
  AObj.Update;
end;

procedure Memo_BringToFront(AObj: TMemo); extdecl;
begin
  AObj.BringToFront;
end;

function Memo_Dragging(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.Dragging;
end;

function Memo_HasParent(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Memo_Hide(AObj: TMemo); extdecl;
begin
  AObj.Hide;
end;

function Memo_Perform(AObj: TMemo; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Memo_Refresh(AObj: TMemo); extdecl;
begin
  AObj.Refresh;
end;

procedure Memo_SendToBack(AObj: TMemo); extdecl;
begin
  AObj.SendToBack;
end;

procedure Memo_Show(AObj: TMemo); extdecl;
begin
  AObj.Show;
end;

function Memo_GetTextBuf(AObj: TMemo; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Memo_GetTextLen(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function Memo_FindComponent(AObj: TMemo; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Memo_GetNamePath(AObj: TMemo): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Memo_Assign(AObj: TMemo; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Memo_ClassName(AObj: TMemo): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Memo_Equals(AObj: TMemo; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Memo_GetHashCode(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Memo_ToString(AObj: TMemo): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Memo_GetAlign(AObj: TMemo): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure Memo_SetAlign(AObj: TMemo; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function Memo_GetAlignment(AObj: TMemo): TAlignment; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure Memo_SetAlignment(AObj: TMemo; AValue: TAlignment); extdecl;
begin
  AObj.Alignment := AValue;
end;

function Memo_GetAnchors(AObj: TMemo): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure Memo_SetAnchors(AObj: TMemo; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function Memo_GetBiDiMode(AObj: TMemo): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Memo_SetBiDiMode(AObj: TMemo; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function Memo_GetBorderStyle(AObj: TMemo): TBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure Memo_SetBorderStyle(AObj: TMemo; AValue: TBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function Memo_GetColor(AObj: TMemo): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure Memo_SetColor(AObj: TMemo; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function Memo_GetDoubleBuffered(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure Memo_SetDoubleBuffered(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function Memo_GetDragCursor(AObj: TMemo): TCursor; extdecl;
begin
  Result :=  AObj.DragCursor;
end;

procedure Memo_SetDragCursor(AObj: TMemo; AValue: TCursor); extdecl;
begin
  AObj.DragCursor := AValue;
end;

function Memo_GetDragKind(AObj: TMemo): TDragKind; extdecl;
begin
  Result :=  AObj.DragKind;
end;

procedure Memo_SetDragKind(AObj: TMemo; AValue: TDragKind); extdecl;
begin
  AObj.DragKind := AValue;
end;

function Memo_GetDragMode(AObj: TMemo): TDragMode; extdecl;
begin
  Result :=  AObj.DragMode;
end;

procedure Memo_SetDragMode(AObj: TMemo; AValue: TDragMode); extdecl;
begin
  AObj.DragMode := AValue;
end;

function Memo_GetEnabled(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Memo_SetEnabled(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Memo_GetFont(AObj: TMemo): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure Memo_SetFont(AObj: TMemo; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function Memo_GetHideSelection(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.HideSelection;
end;

procedure Memo_SetHideSelection(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.HideSelection := AValue;
end;

function Memo_GetLines(AObj: TMemo): TStrings; extdecl;
begin
  Result :=  AObj.Lines;
end;

procedure Memo_SetLines(AObj: TMemo; AValue: TStrings); extdecl;
begin
  AObj.Lines := AValue;
end;

function Memo_GetMaxLength(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.MaxLength;
end;

procedure Memo_SetMaxLength(AObj: TMemo; AValue: Integer); extdecl;
begin
  AObj.MaxLength := AValue;
end;

function Memo_GetParentColor(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure Memo_SetParentColor(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function Memo_GetParentFont(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure Memo_SetParentFont(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function Memo_GetParentShowHint(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Memo_SetParentShowHint(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function Memo_GetPopupMenu(AObj: TMemo): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Memo_SetPopupMenu(AObj: TMemo; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function Memo_GetReadOnly(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.ReadOnly;
end;

procedure Memo_SetReadOnly(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.ReadOnly := AValue;
end;

function Memo_GetScrollBars(AObj: TMemo): TScrollStyle; extdecl;
begin
  Result :=  AObj.ScrollBars;
end;

procedure Memo_SetScrollBars(AObj: TMemo; AValue: TScrollStyle); extdecl;
begin
  AObj.ScrollBars := AValue;
end;

function Memo_GetShowHint(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure Memo_SetShowHint(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function Memo_GetTabOrder(AObj: TMemo): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure Memo_SetTabOrder(AObj: TMemo; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function Memo_GetTabStop(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure Memo_SetTabStop(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function Memo_GetVisible(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure Memo_SetVisible(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function Memo_GetWantReturns(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.WantReturns;
end;

procedure Memo_SetWantReturns(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.WantReturns := AValue;
end;

function Memo_GetWantTabs(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.WantTabs;
end;

procedure Memo_SetWantTabs(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.WantTabs := AValue;
end;

function Memo_GetWordWrap(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.WordWrap;
end;

procedure Memo_SetWordWrap(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.WordWrap := AValue;
end;

procedure Memo_SetOnChange(AObj: TMemo; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    EventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := @EventClass.OnChange;
  EventClass.Add(AObj, geChange, AEventId);
end;

procedure Memo_SetOnClick(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnContextPopup(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnDblClick(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnDragDrop(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnDragOver(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnEndDrag(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnEnter(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnExit(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnKeyDown(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnKeyPress(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnKeyUp(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnMouseDown(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnMouseEnter(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnMouseLeave(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnMouseMove(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_SetOnMouseUp(AObj: TMemo; AEventId: NativeUInt); extdecl;
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

procedure Memo_GetCaretPos(AObj: TMemo; var Result: TPoint); extdecl;
begin
  Result :=  AObj.CaretPos;
end;

procedure Memo_SetCaretPos(AObj: TMemo; var AValue: TPoint); extdecl;
begin
  AObj.CaretPos := AValue;
end;

function Memo_GetCanUndo(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.CanUndo;
end;

function Memo_GetModified(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.Modified;
end;

procedure Memo_SetModified(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.Modified := AValue;
end;

function Memo_GetSelLength(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.SelLength;
end;

procedure Memo_SetSelLength(AObj: TMemo; AValue: Integer); extdecl;
begin
  AObj.SelLength := AValue;
end;

function Memo_GetSelStart(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.SelStart;
end;

procedure Memo_SetSelStart(AObj: TMemo; AValue: Integer); extdecl;
begin
  AObj.SelStart := AValue;
end;

function Memo_GetSelText(AObj: TMemo): PChar; extdecl;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure Memo_SetSelText(AObj: TMemo; AValue: PChar); extdecl;
begin
  AObj.SelText := AValue;
end;

function Memo_GetText(AObj: TMemo): PChar; extdecl;
begin
  Result :=  PChar(AObj.Text);
end;

procedure Memo_SetText(AObj: TMemo; AValue: PChar); extdecl;
begin
  AObj.Text := AValue;
end;

function Memo_GetTextHint(AObj: TMemo): PChar; extdecl;
begin
  Result :=  PChar(AObj.TextHint);
end;

procedure Memo_SetTextHint(AObj: TMemo; AValue: PChar); extdecl;
begin
  AObj.TextHint := AValue;
end;

function Memo_GetDockSite(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.DockSite;
end;

procedure Memo_SetDockSite(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.DockSite := AValue;
end;

function Memo_GetBrush(AObj: TMemo): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function Memo_GetControlCount(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function Memo_GetHandle(AObj: TMemo): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function Memo_GetParentWindow(AObj: TMemo): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure Memo_SetParentWindow(AObj: TMemo; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function Memo_GetUseDockManager(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.UseDockManager;
end;

procedure Memo_SetUseDockManager(AObj: TMemo; AValue: LongBool); extdecl;
begin
  AObj.UseDockManager := AValue;
end;

function Memo_GetAction(AObj: TMemo): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure Memo_SetAction(AObj: TMemo; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure Memo_GetBoundsRect(AObj: TMemo; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Memo_SetBoundsRect(AObj: TMemo; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function Memo_GetClientHeight(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Memo_SetClientHeight(AObj: TMemo; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure Memo_GetClientRect(AObj: TMemo; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function Memo_GetClientWidth(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Memo_SetClientWidth(AObj: TMemo; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function Memo_GetFloating(AObj: TMemo): LongBool; extdecl;
begin
  Result :=  AObj.Floating;
end;

function Memo_GetParent(AObj: TMemo): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure Memo_SetParent(AObj: TMemo; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function Memo_GetLeft(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure Memo_SetLeft(AObj: TMemo; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function Memo_GetTop(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure Memo_SetTop(AObj: TMemo; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function Memo_GetWidth(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Memo_SetWidth(AObj: TMemo; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function Memo_GetHeight(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Memo_SetHeight(AObj: TMemo; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Memo_GetCursor(AObj: TMemo): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure Memo_SetCursor(AObj: TMemo; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function Memo_GetHint(AObj: TMemo): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Memo_SetHint(AObj: TMemo; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function Memo_GetComponentCount(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Memo_GetComponentIndex(AObj: TMemo): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Memo_SetComponentIndex(AObj: TMemo; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Memo_GetOwner(AObj: TMemo): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Memo_GetName(AObj: TMemo): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Memo_SetName(AObj: TMemo; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Memo_GetTag(AObj: TMemo): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Memo_SetTag(AObj: TMemo; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Memo_GetControls(AObj: TMemo; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function Memo_GetComponents(AObj: TMemo; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Memo_Create,
  Memo_Free,
  Memo_Clear,
  Memo_ClearSelection,
  Memo_CopyToClipboard,
  Memo_CutToClipboard,
  Memo_PasteFromClipboard,
  Memo_Undo,
  Memo_SelectAll,
  Memo_CanFocus,
  Memo_FlipChildren,
  Memo_Focused,
  Memo_HandleAllocated,
  Memo_Invalidate,
  Memo_Realign,
  Memo_Repaint,
  Memo_ScaleBy,
  Memo_SetBounds,
  Memo_SetFocus,
  Memo_Update,
  Memo_BringToFront,
  Memo_Dragging,
  Memo_HasParent,
  Memo_Hide,
  Memo_Perform,
  Memo_Refresh,
  Memo_SendToBack,
  Memo_Show,
  Memo_GetTextBuf,
  Memo_GetTextLen,
  Memo_FindComponent,
  Memo_GetNamePath,
  Memo_Assign,
  Memo_ClassName,
  Memo_Equals,
  Memo_GetHashCode,
  Memo_ToString,
  Memo_GetAlign,
  Memo_SetAlign,
  Memo_GetAlignment,
  Memo_SetAlignment,
  Memo_GetAnchors,
  Memo_SetAnchors,
  Memo_GetBiDiMode,
  Memo_SetBiDiMode,
  Memo_GetBorderStyle,
  Memo_SetBorderStyle,
  Memo_GetColor,
  Memo_SetColor,
  Memo_GetDoubleBuffered,
  Memo_SetDoubleBuffered,
  Memo_GetDragCursor,
  Memo_SetDragCursor,
  Memo_GetDragKind,
  Memo_SetDragKind,
  Memo_GetDragMode,
  Memo_SetDragMode,
  Memo_GetEnabled,
  Memo_SetEnabled,
  Memo_GetFont,
  Memo_SetFont,
  Memo_GetHideSelection,
  Memo_SetHideSelection,
  Memo_GetLines,
  Memo_SetLines,
  Memo_GetMaxLength,
  Memo_SetMaxLength,
  Memo_GetParentColor,
  Memo_SetParentColor,
  Memo_GetParentFont,
  Memo_SetParentFont,
  Memo_GetParentShowHint,
  Memo_SetParentShowHint,
  Memo_GetPopupMenu,
  Memo_SetPopupMenu,
  Memo_GetReadOnly,
  Memo_SetReadOnly,
  Memo_GetScrollBars,
  Memo_SetScrollBars,
  Memo_GetShowHint,
  Memo_SetShowHint,
  Memo_GetTabOrder,
  Memo_SetTabOrder,
  Memo_GetTabStop,
  Memo_SetTabStop,
  Memo_GetVisible,
  Memo_SetVisible,
  Memo_GetWantReturns,
  Memo_SetWantReturns,
  Memo_GetWantTabs,
  Memo_SetWantTabs,
  Memo_GetWordWrap,
  Memo_SetWordWrap,
  Memo_SetOnChange,
  Memo_SetOnClick,
  Memo_SetOnContextPopup,
  Memo_SetOnDblClick,
  Memo_SetOnDragDrop,
  Memo_SetOnDragOver,
  Memo_SetOnEndDrag,
  Memo_SetOnEnter,
  Memo_SetOnExit,
  Memo_SetOnKeyDown,
  Memo_SetOnKeyPress,
  Memo_SetOnKeyUp,
  Memo_SetOnMouseDown,
  Memo_SetOnMouseEnter,
  Memo_SetOnMouseLeave,
  Memo_SetOnMouseMove,
  Memo_SetOnMouseUp,
  Memo_GetCaretPos,
  Memo_SetCaretPos,
  Memo_GetCanUndo,
  Memo_GetModified,
  Memo_SetModified,
  Memo_GetSelLength,
  Memo_SetSelLength,
  Memo_GetSelStart,
  Memo_SetSelStart,
  Memo_GetSelText,
  Memo_SetSelText,
  Memo_GetText,
  Memo_SetText,
  Memo_GetTextHint,
  Memo_SetTextHint,
  Memo_GetDockSite,
  Memo_SetDockSite,
  Memo_GetBrush,
  Memo_GetControlCount,
  Memo_GetHandle,
  Memo_GetParentWindow,
  Memo_SetParentWindow,
  Memo_GetUseDockManager,
  Memo_SetUseDockManager,
  Memo_GetAction,
  Memo_SetAction,
  Memo_GetBoundsRect,
  Memo_SetBoundsRect,
  Memo_GetClientHeight,
  Memo_SetClientHeight,
  Memo_GetClientRect,
  Memo_GetClientWidth,
  Memo_SetClientWidth,
  Memo_GetFloating,
  Memo_GetParent,
  Memo_SetParent,
  Memo_GetLeft,
  Memo_SetLeft,
  Memo_GetTop,
  Memo_SetTop,
  Memo_GetWidth,
  Memo_SetWidth,
  Memo_GetHeight,
  Memo_SetHeight,
  Memo_GetCursor,
  Memo_SetCursor,
  Memo_GetHint,
  Memo_SetHint,
  Memo_GetComponentCount,
  Memo_GetComponentIndex,
  Memo_SetComponentIndex,
  Memo_GetOwner,
  Memo_GetName,
  Memo_SetName,
  Memo_GetTag,
  Memo_SetTag,
  Memo_GetControls,
  Memo_GetComponents;

