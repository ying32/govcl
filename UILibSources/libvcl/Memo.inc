
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Memo_Create(AOwner: TComponent): TMemo; stdcall;
begin
  Result :=  TMemo.Create(AOwner);
end;

procedure Memo_Free(AObj: TMemo); stdcall;
begin
  AObj.Free;
end;

procedure Memo_Clear(AObj: TMemo); stdcall;
begin
  AObj.Clear;
end;

procedure Memo_ClearSelection(AObj: TMemo); stdcall;
begin
  AObj.ClearSelection;
end;

procedure Memo_CopyToClipboard(AObj: TMemo); stdcall;
begin
  AObj.CopyToClipboard;
end;

procedure Memo_CutToClipboard(AObj: TMemo); stdcall;
begin
  AObj.CutToClipboard;
end;

procedure Memo_PasteFromClipboard(AObj: TMemo); stdcall;
begin
  AObj.PasteFromClipboard;
end;

procedure Memo_Undo(AObj: TMemo); stdcall;
begin
  AObj.Undo;
end;

procedure Memo_ClearUndo(AObj: TMemo); stdcall;
begin
  AObj.ClearUndo;
end;

procedure Memo_SelectAll(AObj: TMemo); stdcall;
begin
  AObj.SelectAll;
end;

function Memo_GetSelTextBuf(AObj: TMemo; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetSelTextBuf(Buffer, BufSize);
end;

function Memo_CanFocus(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure Memo_FlipChildren(AObj: TMemo; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function Memo_Focused(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function Memo_HandleAllocated(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure Memo_Invalidate(AObj: TMemo); stdcall;
begin
  AObj.Invalidate;
end;

procedure Memo_Realign(AObj: TMemo); stdcall;
begin
  AObj.Realign;
end;

procedure Memo_Repaint(AObj: TMemo); stdcall;
begin
  AObj.Repaint;
end;

procedure Memo_ScaleBy(AObj: TMemo; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure Memo_SetBounds(AObj: TMemo; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Memo_SetFocus(AObj: TMemo); stdcall;
begin
  AObj.SetFocus;
end;

procedure Memo_Update(AObj: TMemo); stdcall;
begin
  AObj.Update;
end;

procedure Memo_BringToFront(AObj: TMemo); stdcall;
begin
  AObj.BringToFront;
end;

procedure Memo_ClientToScreen(AObj: TMemo; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure Memo_ClientToParent(AObj: TMemo; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function Memo_Dragging(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function Memo_HasParent(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Memo_Hide(AObj: TMemo); stdcall;
begin
  AObj.Hide;
end;

function Memo_Perform(AObj: TMemo; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Memo_Refresh(AObj: TMemo); stdcall;
begin
  AObj.Refresh;
end;

procedure Memo_ScreenToClient(AObj: TMemo; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure Memo_ParentToClient(AObj: TMemo; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure Memo_SendToBack(AObj: TMemo); stdcall;
begin
  AObj.SendToBack;
end;

procedure Memo_Show(AObj: TMemo); stdcall;
begin
  AObj.Show;
end;

function Memo_GetTextBuf(AObj: TMemo; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Memo_GetTextLen(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function Memo_FindComponent(AObj: TMemo; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Memo_GetNamePath(AObj: TMemo): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Memo_Assign(AObj: TMemo; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Memo_ClassName(AObj: TMemo): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Memo_Equals(AObj: TMemo; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Memo_GetHashCode(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Memo_ToString(AObj: TMemo): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Memo_GetAlign(AObj: TMemo): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Memo_SetAlign(AObj: TMemo; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Memo_GetAlignment(AObj: TMemo): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure Memo_SetAlignment(AObj: TMemo; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function Memo_GetAnchors(AObj: TMemo): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Memo_SetAnchors(AObj: TMemo; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Memo_GetBevelEdges(AObj: TMemo): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure Memo_SetBevelEdges(AObj: TMemo; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function Memo_GetBevelInner(AObj: TMemo): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure Memo_SetBevelInner(AObj: TMemo; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function Memo_GetBevelKind(AObj: TMemo): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure Memo_SetBevelKind(AObj: TMemo; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function Memo_GetBevelOuter(AObj: TMemo): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure Memo_SetBevelOuter(AObj: TMemo; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function Memo_GetBiDiMode(AObj: TMemo): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Memo_SetBiDiMode(AObj: TMemo; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function Memo_GetBorderStyle(AObj: TMemo): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure Memo_SetBorderStyle(AObj: TMemo; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function Memo_GetColor(AObj: TMemo): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure Memo_SetColor(AObj: TMemo; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function Memo_GetDoubleBuffered(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure Memo_SetDoubleBuffered(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function Memo_GetDragCursor(AObj: TMemo): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure Memo_SetDragCursor(AObj: TMemo; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function Memo_GetDragKind(AObj: TMemo): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure Memo_SetDragKind(AObj: TMemo; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function Memo_GetDragMode(AObj: TMemo): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure Memo_SetDragMode(AObj: TMemo; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function Memo_GetEnabled(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Memo_SetEnabled(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Memo_GetFont(AObj: TMemo): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure Memo_SetFont(AObj: TMemo; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function Memo_GetHideSelection(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.HideSelection;
end;

procedure Memo_SetHideSelection(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.HideSelection := AValue;
end;

function Memo_GetLines(AObj: TMemo): TStrings; stdcall;
begin
  Result :=  AObj.Lines;
end;

procedure Memo_SetLines(AObj: TMemo; AValue: TStrings); stdcall;
begin
  AObj.Lines := AValue;
end;

function Memo_GetMaxLength(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.MaxLength;
end;

procedure Memo_SetMaxLength(AObj: TMemo; AValue: Integer); stdcall;
begin
  AObj.MaxLength := AValue;
end;

function Memo_GetParentColor(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure Memo_SetParentColor(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function Memo_GetParentCtl3D(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure Memo_SetParentCtl3D(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function Memo_GetParentDoubleBuffered(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure Memo_SetParentDoubleBuffered(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function Memo_GetParentFont(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure Memo_SetParentFont(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function Memo_GetParentShowHint(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Memo_SetParentShowHint(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function Memo_GetPopupMenu(AObj: TMemo): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Memo_SetPopupMenu(AObj: TMemo; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function Memo_GetReadOnly(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.ReadOnly;
end;

procedure Memo_SetReadOnly(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.ReadOnly := AValue;
end;

function Memo_GetScrollBars(AObj: TMemo): TScrollStyle; stdcall;
begin
  Result :=  AObj.ScrollBars;
end;

procedure Memo_SetScrollBars(AObj: TMemo; AValue: TScrollStyle); stdcall;
begin
  AObj.ScrollBars := AValue;
end;

function Memo_GetShowHint(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Memo_SetShowHint(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Memo_GetTabOrder(AObj: TMemo): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure Memo_SetTabOrder(AObj: TMemo; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function Memo_GetTabStop(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure Memo_SetTabStop(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function Memo_GetVisible(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Memo_SetVisible(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Memo_GetWantReturns(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.WantReturns;
end;

procedure Memo_SetWantReturns(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.WantReturns := AValue;
end;

function Memo_GetWantTabs(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.WantTabs;
end;

procedure Memo_SetWantTabs(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.WantTabs := AValue;
end;

function Memo_GetWordWrap(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.WordWrap;
end;

procedure Memo_SetWordWrap(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.WordWrap := AValue;
end;

function Memo_GetStyleElements(AObj: TMemo): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Memo_SetStyleElements(AObj: TMemo; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure Memo_SetOnChange(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnClick(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnContextPopup(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnDblClick(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnDragDrop(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnDragOver(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnEndDock(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnEndDrag(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnEnter(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnExit(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnKeyDown(AObj: TMemo; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyDown := nil;
    TEventClass.Remove(AObj, geKeyDown);
    Exit;
  end;
  AObj.OnKeyDown := TEventClass.OnKeyDown;
  TEventClass.Add(AObj, geKeyDown, AEventId);
end;

procedure Memo_SetOnKeyPress(AObj: TMemo; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyPress := nil;
    TEventClass.Remove(AObj, geKeyPress);
    Exit;
  end;
  AObj.OnKeyPress := TEventClass.OnKeyPress;
  TEventClass.Add(AObj, geKeyPress, AEventId);
end;

procedure Memo_SetOnKeyUp(AObj: TMemo; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyUp := nil;
    TEventClass.Remove(AObj, geKeyUp);
    Exit;
  end;
  AObj.OnKeyUp := TEventClass.OnKeyUp;
  TEventClass.Add(AObj, geKeyUp, AEventId);
end;

procedure Memo_SetOnMouseDown(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnMouseEnter(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnMouseLeave(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnMouseMove(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnMouseUp(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_SetOnStartDock(AObj: TMemo; AEventId: NativeUInt); stdcall;
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

procedure Memo_GetCaretPos(AObj: TMemo; var Result: TPoint); stdcall;
begin
  Result :=  AObj.CaretPos;
end;

procedure Memo_SetCaretPos(AObj: TMemo; var AValue: TPoint); stdcall;
begin
  AObj.CaretPos := AValue;
end;

function Memo_GetCanUndo(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.CanUndo;
end;

function Memo_GetModified(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure Memo_SetModified(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function Memo_GetSelLength(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.SelLength;
end;

procedure Memo_SetSelLength(AObj: TMemo; AValue: Integer); stdcall;
begin
  AObj.SelLength := AValue;
end;

function Memo_GetSelStart(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.SelStart;
end;

procedure Memo_SetSelStart(AObj: TMemo; AValue: Integer); stdcall;
begin
  AObj.SelStart := AValue;
end;

function Memo_GetSelText(AObj: TMemo): PChar; stdcall;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure Memo_SetSelText(AObj: TMemo; AValue: PChar); stdcall;
begin
  AObj.SelText := AValue;
end;

function Memo_GetText(AObj: TMemo): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure Memo_SetText(AObj: TMemo; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function Memo_GetTextHint(AObj: TMemo): PChar; stdcall;
begin
  Result :=  PChar(AObj.TextHint);
end;

procedure Memo_SetTextHint(AObj: TMemo; AValue: PChar); stdcall;
begin
  AObj.TextHint := AValue;
end;

function Memo_GetDockSite(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure Memo_SetDockSite(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function Memo_GetBrush(AObj: TMemo): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function Memo_GetControlCount(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function Memo_GetHandle(AObj: TMemo): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function Memo_GetParentWindow(AObj: TMemo): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure Memo_SetParentWindow(AObj: TMemo; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function Memo_GetUseDockManager(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure Memo_SetUseDockManager(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function Memo_GetAction(AObj: TMemo): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Memo_SetAction(AObj: TMemo; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure Memo_GetBoundsRect(AObj: TMemo; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Memo_SetBoundsRect(AObj: TMemo; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Memo_GetClientHeight(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Memo_SetClientHeight(AObj: TMemo; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Memo_GetClientRect(AObj: TMemo; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Memo_GetClientWidth(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Memo_SetClientWidth(AObj: TMemo; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Memo_GetExplicitLeft(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Memo_GetExplicitTop(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Memo_GetExplicitWidth(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Memo_GetExplicitHeight(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Memo_GetFloating(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function Memo_GetParent(AObj: TMemo): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Memo_SetParent(AObj: TMemo; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Memo_GetAlignWithMargins(AObj: TMemo): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Memo_SetAlignWithMargins(AObj: TMemo; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Memo_GetLeft(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Memo_SetLeft(AObj: TMemo; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Memo_GetTop(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Memo_SetTop(AObj: TMemo; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Memo_GetWidth(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Memo_SetWidth(AObj: TMemo; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Memo_GetHeight(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Memo_SetHeight(AObj: TMemo; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Memo_GetCursor(AObj: TMemo): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Memo_SetCursor(AObj: TMemo; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Memo_GetHint(AObj: TMemo): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Memo_SetHint(AObj: TMemo; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Memo_GetMargins(AObj: TMemo): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Memo_SetMargins(AObj: TMemo; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Memo_GetCustomHint(AObj: TMemo): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Memo_SetCustomHint(AObj: TMemo; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Memo_GetComponentCount(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Memo_GetComponentIndex(AObj: TMemo): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Memo_SetComponentIndex(AObj: TMemo; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Memo_GetOwner(AObj: TMemo): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Memo_GetName(AObj: TMemo): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Memo_SetName(AObj: TMemo; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Memo_GetTag(AObj: TMemo): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Memo_SetTag(AObj: TMemo; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Memo_GetControls(AObj: TMemo; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function Memo_GetComponents(AObj: TMemo; AIndex: Integer): TComponent; stdcall;
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
  Memo_ClearUndo,
  Memo_SelectAll,
  Memo_GetSelTextBuf,
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
  Memo_ClientToScreen,
  Memo_ClientToParent,
  Memo_Dragging,
  Memo_HasParent,
  Memo_Hide,
  Memo_Perform,
  Memo_Refresh,
  Memo_ScreenToClient,
  Memo_ParentToClient,
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
  Memo_GetBevelEdges,
  Memo_SetBevelEdges,
  Memo_GetBevelInner,
  Memo_SetBevelInner,
  Memo_GetBevelKind,
  Memo_SetBevelKind,
  Memo_GetBevelOuter,
  Memo_SetBevelOuter,
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
  Memo_GetParentCtl3D,
  Memo_SetParentCtl3D,
  Memo_GetParentDoubleBuffered,
  Memo_SetParentDoubleBuffered,
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
  Memo_GetStyleElements,
  Memo_SetStyleElements,
  Memo_SetOnChange,
  Memo_SetOnClick,
  Memo_SetOnContextPopup,
  Memo_SetOnDblClick,
  Memo_SetOnDragDrop,
  Memo_SetOnDragOver,
  Memo_SetOnEndDock,
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
  Memo_SetOnStartDock,
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
  Memo_GetExplicitLeft,
  Memo_GetExplicitTop,
  Memo_GetExplicitWidth,
  Memo_GetExplicitHeight,
  Memo_GetFloating,
  Memo_GetParent,
  Memo_SetParent,
  Memo_GetAlignWithMargins,
  Memo_SetAlignWithMargins,
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
  Memo_GetMargins,
  Memo_SetMargins,
  Memo_GetCustomHint,
  Memo_SetCustomHint,
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

