
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function RichEdit_Create(AOwner: TComponent): TRichEdit; stdcall;
begin
  Result :=  TRichEdit.Create(AOwner);
end;

procedure RichEdit_Free(AObj: TRichEdit); stdcall;
begin
  AObj.Free;
end;

procedure RichEdit_Clear(AObj: TRichEdit); stdcall;
begin
  AObj.Clear;
end;

function RichEdit_FindText(AObj: TRichEdit; SearchStr: PChar; StartPos: Integer; Length: Integer; Options: TSearchTypes): Integer; stdcall;
begin
  Result :=  AObj.FindText(SearchStr, StartPos, Length, Options);
end;

procedure RichEdit_Print(AObj: TRichEdit; Caption: PChar); stdcall;
begin
  AObj.Print(Caption);
end;

function RichEdit_GetSelTextBuf(AObj: TRichEdit; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetSelTextBuf(Buffer, BufSize);
end;

procedure RichEdit_ClearSelection(AObj: TRichEdit); stdcall;
begin
  AObj.ClearSelection;
end;

procedure RichEdit_CopyToClipboard(AObj: TRichEdit); stdcall;
begin
  AObj.CopyToClipboard;
end;

procedure RichEdit_CutToClipboard(AObj: TRichEdit); stdcall;
begin
  AObj.CutToClipboard;
end;

procedure RichEdit_PasteFromClipboard(AObj: TRichEdit); stdcall;
begin
  AObj.PasteFromClipboard;
end;

procedure RichEdit_Undo(AObj: TRichEdit); stdcall;
begin
  AObj.Undo;
end;

procedure RichEdit_ClearUndo(AObj: TRichEdit); stdcall;
begin
  AObj.ClearUndo;
end;

procedure RichEdit_SelectAll(AObj: TRichEdit); stdcall;
begin
  AObj.SelectAll;
end;

function RichEdit_CanFocus(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure RichEdit_FlipChildren(AObj: TRichEdit; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function RichEdit_Focused(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function RichEdit_HandleAllocated(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure RichEdit_Invalidate(AObj: TRichEdit); stdcall;
begin
  AObj.Invalidate;
end;

procedure RichEdit_Realign(AObj: TRichEdit); stdcall;
begin
  AObj.Realign;
end;

procedure RichEdit_Repaint(AObj: TRichEdit); stdcall;
begin
  AObj.Repaint;
end;

procedure RichEdit_ScaleBy(AObj: TRichEdit; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure RichEdit_SetBounds(AObj: TRichEdit; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure RichEdit_SetFocus(AObj: TRichEdit); stdcall;
begin
  AObj.SetFocus;
end;

procedure RichEdit_Update(AObj: TRichEdit); stdcall;
begin
  AObj.Update;
end;

procedure RichEdit_BringToFront(AObj: TRichEdit); stdcall;
begin
  AObj.BringToFront;
end;

function RichEdit_Dragging(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function RichEdit_HasParent(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure RichEdit_Hide(AObj: TRichEdit); stdcall;
begin
  AObj.Hide;
end;

function RichEdit_Perform(AObj: TRichEdit; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure RichEdit_Refresh(AObj: TRichEdit); stdcall;
begin
  AObj.Refresh;
end;

procedure RichEdit_SendToBack(AObj: TRichEdit); stdcall;
begin
  AObj.SendToBack;
end;

procedure RichEdit_Show(AObj: TRichEdit); stdcall;
begin
  AObj.Show;
end;

function RichEdit_GetTextBuf(AObj: TRichEdit; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function RichEdit_GetTextLen(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function RichEdit_FindComponent(AObj: TRichEdit; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function RichEdit_GetNamePath(AObj: TRichEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure RichEdit_Assign(AObj: TRichEdit; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function RichEdit_ClassName(AObj: TRichEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function RichEdit_Equals(AObj: TRichEdit; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function RichEdit_GetHashCode(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function RichEdit_ToString(AObj: TRichEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function RichEdit_GetAlign(AObj: TRichEdit): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure RichEdit_SetAlign(AObj: TRichEdit; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function RichEdit_GetAlignment(AObj: TRichEdit): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure RichEdit_SetAlignment(AObj: TRichEdit; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function RichEdit_GetAnchors(AObj: TRichEdit): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure RichEdit_SetAnchors(AObj: TRichEdit; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function RichEdit_GetBevelEdges(AObj: TRichEdit): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure RichEdit_SetBevelEdges(AObj: TRichEdit; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function RichEdit_GetBevelInner(AObj: TRichEdit): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure RichEdit_SetBevelInner(AObj: TRichEdit; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function RichEdit_GetBevelOuter(AObj: TRichEdit): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure RichEdit_SetBevelOuter(AObj: TRichEdit; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function RichEdit_GetBevelKind(AObj: TRichEdit): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure RichEdit_SetBevelKind(AObj: TRichEdit; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function RichEdit_GetBiDiMode(AObj: TRichEdit): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure RichEdit_SetBiDiMode(AObj: TRichEdit; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function RichEdit_GetBorderStyle(AObj: TRichEdit): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure RichEdit_SetBorderStyle(AObj: TRichEdit; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function RichEdit_GetBorderWidth(AObj: TRichEdit): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure RichEdit_SetBorderWidth(AObj: TRichEdit; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function RichEdit_GetColor(AObj: TRichEdit): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure RichEdit_SetColor(AObj: TRichEdit; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function RichEdit_GetDragCursor(AObj: TRichEdit): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure RichEdit_SetDragCursor(AObj: TRichEdit; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function RichEdit_GetDragKind(AObj: TRichEdit): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure RichEdit_SetDragKind(AObj: TRichEdit; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function RichEdit_GetDragMode(AObj: TRichEdit): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure RichEdit_SetDragMode(AObj: TRichEdit; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function RichEdit_GetEnabled(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure RichEdit_SetEnabled(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function RichEdit_GetFont(AObj: TRichEdit): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure RichEdit_SetFont(AObj: TRichEdit; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function RichEdit_GetHideSelection(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.HideSelection;
end;

procedure RichEdit_SetHideSelection(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.HideSelection := AValue;
end;

function RichEdit_GetHideScrollBars(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.HideScrollBars;
end;

procedure RichEdit_SetHideScrollBars(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.HideScrollBars := AValue;
end;

function RichEdit_GetLines(AObj: TRichEdit): TStrings; stdcall;
begin
  Result :=  AObj.Lines;
end;

procedure RichEdit_SetLines(AObj: TRichEdit; AValue: TStrings); stdcall;
begin
  AObj.Lines := AValue;
end;

function RichEdit_GetMaxLength(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.MaxLength;
end;

procedure RichEdit_SetMaxLength(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.MaxLength := AValue;
end;

function RichEdit_GetParentColor(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure RichEdit_SetParentColor(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function RichEdit_GetParentCtl3D(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure RichEdit_SetParentCtl3D(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function RichEdit_GetParentFont(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure RichEdit_SetParentFont(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function RichEdit_GetParentShowHint(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure RichEdit_SetParentShowHint(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function RichEdit_GetPlainText(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.PlainText;
end;

procedure RichEdit_SetPlainText(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.PlainText := AValue;
end;

function RichEdit_GetPopupMenu(AObj: TRichEdit): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure RichEdit_SetPopupMenu(AObj: TRichEdit; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function RichEdit_GetReadOnly(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.ReadOnly;
end;

procedure RichEdit_SetReadOnly(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.ReadOnly := AValue;
end;

function RichEdit_GetScrollBars(AObj: TRichEdit): TScrollStyle; stdcall;
begin
  Result :=  AObj.ScrollBars;
end;

procedure RichEdit_SetScrollBars(AObj: TRichEdit; AValue: TScrollStyle); stdcall;
begin
  AObj.ScrollBars := AValue;
end;

function RichEdit_GetShowHint(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure RichEdit_SetShowHint(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function RichEdit_GetTabOrder(AObj: TRichEdit): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure RichEdit_SetTabOrder(AObj: TRichEdit; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function RichEdit_GetTabStop(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure RichEdit_SetTabStop(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function RichEdit_GetVisible(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure RichEdit_SetVisible(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function RichEdit_GetWantTabs(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.WantTabs;
end;

procedure RichEdit_SetWantTabs(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.WantTabs := AValue;
end;

function RichEdit_GetWantReturns(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.WantReturns;
end;

procedure RichEdit_SetWantReturns(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.WantReturns := AValue;
end;

function RichEdit_GetWordWrap(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.WordWrap;
end;

procedure RichEdit_SetWordWrap(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.WordWrap := AValue;
end;

function RichEdit_GetStyleElements(AObj: TRichEdit): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure RichEdit_SetStyleElements(AObj: TRichEdit; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function RichEdit_GetZoom(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.Zoom;
end;

procedure RichEdit_SetZoom(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.Zoom := AValue;
end;

procedure RichEdit_SetOnChange(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnClick(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnContextPopup(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnDblClick(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnDragDrop(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnDragOver(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnEndDock(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnEndDrag(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnEnter(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnExit(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnKeyDown(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnKeyPress(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnKeyUp(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnMouseDown(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnMouseEnter(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnMouseLeave(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnMouseMove(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnMouseUp(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

procedure RichEdit_SetOnMouseWheel(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheel := nil;
    TEventClass.Remove(AObj, geMouseWheel);
    Exit;
  end;
  AObj.OnMouseWheel := TEventClass.OnMouseWheel;
  TEventClass.Add(AObj, geMouseWheel, AEventId);
end;

procedure RichEdit_SetOnMouseWheelDown(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheelDown := nil;
    TEventClass.Remove(AObj, geMouseWheelDown);
    Exit;
  end;
  AObj.OnMouseWheelDown := TEventClass.OnMouseWheelDown;
  TEventClass.Add(AObj, geMouseWheelDown, AEventId);
end;

procedure RichEdit_SetOnMouseWheelUp(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheelUp := nil;
    TEventClass.Remove(AObj, geMouseWheelUp);
    Exit;
  end;
  AObj.OnMouseWheelUp := TEventClass.OnMouseWheelUp;
  TEventClass.Add(AObj, geMouseWheelUp, AEventId);
end;

procedure RichEdit_SetOnStartDock(AObj: TRichEdit; AEventId: NativeUInt); stdcall;
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

function RichEdit_GetActiveLineNo(AObj: TRichEdit): Cardinal; stdcall;
begin
  Result :=  AObj.ActiveLineNo;
end;

function RichEdit_GetDefAttributes(AObj: TRichEdit): TTextAttributes; stdcall;
begin
  Result :=  AObj.DefAttributes;
end;

procedure RichEdit_SetDefAttributes(AObj: TRichEdit; AValue: TTextAttributes); stdcall;
begin
  AObj.DefAttributes := AValue;
end;

function RichEdit_GetSelAttributes(AObj: TRichEdit): TTextAttributes; stdcall;
begin
  Result :=  AObj.SelAttributes;
end;

procedure RichEdit_SetSelAttributes(AObj: TRichEdit; AValue: TTextAttributes); stdcall;
begin
  AObj.SelAttributes := AValue;
end;

procedure RichEdit_GetPageRect(AObj: TRichEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.PageRect;
end;

procedure RichEdit_SetPageRect(AObj: TRichEdit; var AValue: TRect); stdcall;
begin
  AObj.PageRect := AValue;
end;

function RichEdit_GetParagraph(AObj: TRichEdit): TParaAttributes; stdcall;
begin
  Result :=  AObj.Paragraph;
end;

procedure RichEdit_GetCaretPos(AObj: TRichEdit; var Result: TPoint); stdcall;
begin
  Result :=  AObj.CaretPos;
end;

procedure RichEdit_SetCaretPos(AObj: TRichEdit; var AValue: TPoint); stdcall;
begin
  AObj.CaretPos := AValue;
end;

function RichEdit_GetCanUndo(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.CanUndo;
end;

function RichEdit_GetModified(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure RichEdit_SetModified(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function RichEdit_GetSelLength(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.SelLength;
end;

procedure RichEdit_SetSelLength(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.SelLength := AValue;
end;

function RichEdit_GetSelStart(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.SelStart;
end;

procedure RichEdit_SetSelStart(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.SelStart := AValue;
end;

function RichEdit_GetSelText(AObj: TRichEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure RichEdit_SetSelText(AObj: TRichEdit; AValue: PChar); stdcall;
begin
  AObj.SelText := AValue;
end;

function RichEdit_GetText(AObj: TRichEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure RichEdit_SetText(AObj: TRichEdit; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function RichEdit_GetTextHint(AObj: TRichEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.TextHint);
end;

procedure RichEdit_SetTextHint(AObj: TRichEdit; AValue: PChar); stdcall;
begin
  AObj.TextHint := AValue;
end;

function RichEdit_GetDockSite(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure RichEdit_SetDockSite(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function RichEdit_GetDoubleBuffered(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure RichEdit_SetDoubleBuffered(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function RichEdit_GetBrush(AObj: TRichEdit): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function RichEdit_GetControlCount(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function RichEdit_GetHandle(AObj: TRichEdit): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function RichEdit_GetParentDoubleBuffered(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure RichEdit_SetParentDoubleBuffered(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function RichEdit_GetParentWindow(AObj: TRichEdit): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure RichEdit_SetParentWindow(AObj: TRichEdit; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function RichEdit_GetUseDockManager(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure RichEdit_SetUseDockManager(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function RichEdit_GetAction(AObj: TRichEdit): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure RichEdit_SetAction(AObj: TRichEdit; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure RichEdit_GetBoundsRect(AObj: TRichEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure RichEdit_SetBoundsRect(AObj: TRichEdit; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function RichEdit_GetClientHeight(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure RichEdit_SetClientHeight(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure RichEdit_GetClientRect(AObj: TRichEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function RichEdit_GetClientWidth(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure RichEdit_SetClientWidth(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function RichEdit_GetExplicitLeft(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function RichEdit_GetExplicitTop(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function RichEdit_GetExplicitWidth(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function RichEdit_GetExplicitHeight(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function RichEdit_GetFloating(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function RichEdit_GetParent(AObj: TRichEdit): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure RichEdit_SetParent(AObj: TRichEdit; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function RichEdit_GetAlignWithMargins(AObj: TRichEdit): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure RichEdit_SetAlignWithMargins(AObj: TRichEdit; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function RichEdit_GetLeft(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure RichEdit_SetLeft(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function RichEdit_GetTop(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure RichEdit_SetTop(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function RichEdit_GetWidth(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure RichEdit_SetWidth(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function RichEdit_GetHeight(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure RichEdit_SetHeight(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function RichEdit_GetCursor(AObj: TRichEdit): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure RichEdit_SetCursor(AObj: TRichEdit; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function RichEdit_GetHint(AObj: TRichEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure RichEdit_SetHint(AObj: TRichEdit; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function RichEdit_GetMargins(AObj: TRichEdit): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure RichEdit_SetMargins(AObj: TRichEdit; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function RichEdit_GetCustomHint(AObj: TRichEdit): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure RichEdit_SetCustomHint(AObj: TRichEdit; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function RichEdit_GetComponentCount(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function RichEdit_GetComponentIndex(AObj: TRichEdit): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure RichEdit_SetComponentIndex(AObj: TRichEdit; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function RichEdit_GetOwner(AObj: TRichEdit): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function RichEdit_GetName(AObj: TRichEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure RichEdit_SetName(AObj: TRichEdit; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function RichEdit_GetTag(AObj: TRichEdit): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure RichEdit_SetTag(AObj: TRichEdit; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function RichEdit_GetControls(AObj: TRichEdit; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function RichEdit_GetComponents(AObj: TRichEdit; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  RichEdit_Create,
  RichEdit_Free,
  RichEdit_Clear,
  RichEdit_FindText,
  RichEdit_Print,
  RichEdit_GetSelTextBuf,
  RichEdit_ClearSelection,
  RichEdit_CopyToClipboard,
  RichEdit_CutToClipboard,
  RichEdit_PasteFromClipboard,
  RichEdit_Undo,
  RichEdit_ClearUndo,
  RichEdit_SelectAll,
  RichEdit_CanFocus,
  RichEdit_FlipChildren,
  RichEdit_Focused,
  RichEdit_HandleAllocated,
  RichEdit_Invalidate,
  RichEdit_Realign,
  RichEdit_Repaint,
  RichEdit_ScaleBy,
  RichEdit_SetBounds,
  RichEdit_SetFocus,
  RichEdit_Update,
  RichEdit_BringToFront,
  RichEdit_Dragging,
  RichEdit_HasParent,
  RichEdit_Hide,
  RichEdit_Perform,
  RichEdit_Refresh,
  RichEdit_SendToBack,
  RichEdit_Show,
  RichEdit_GetTextBuf,
  RichEdit_GetTextLen,
  RichEdit_FindComponent,
  RichEdit_GetNamePath,
  RichEdit_Assign,
  RichEdit_ClassName,
  RichEdit_Equals,
  RichEdit_GetHashCode,
  RichEdit_ToString,
  RichEdit_GetAlign,
  RichEdit_SetAlign,
  RichEdit_GetAlignment,
  RichEdit_SetAlignment,
  RichEdit_GetAnchors,
  RichEdit_SetAnchors,
  RichEdit_GetBevelEdges,
  RichEdit_SetBevelEdges,
  RichEdit_GetBevelInner,
  RichEdit_SetBevelInner,
  RichEdit_GetBevelOuter,
  RichEdit_SetBevelOuter,
  RichEdit_GetBevelKind,
  RichEdit_SetBevelKind,
  RichEdit_GetBiDiMode,
  RichEdit_SetBiDiMode,
  RichEdit_GetBorderStyle,
  RichEdit_SetBorderStyle,
  RichEdit_GetBorderWidth,
  RichEdit_SetBorderWidth,
  RichEdit_GetColor,
  RichEdit_SetColor,
  RichEdit_GetDragCursor,
  RichEdit_SetDragCursor,
  RichEdit_GetDragKind,
  RichEdit_SetDragKind,
  RichEdit_GetDragMode,
  RichEdit_SetDragMode,
  RichEdit_GetEnabled,
  RichEdit_SetEnabled,
  RichEdit_GetFont,
  RichEdit_SetFont,
  RichEdit_GetHideSelection,
  RichEdit_SetHideSelection,
  RichEdit_GetHideScrollBars,
  RichEdit_SetHideScrollBars,
  RichEdit_GetLines,
  RichEdit_SetLines,
  RichEdit_GetMaxLength,
  RichEdit_SetMaxLength,
  RichEdit_GetParentColor,
  RichEdit_SetParentColor,
  RichEdit_GetParentCtl3D,
  RichEdit_SetParentCtl3D,
  RichEdit_GetParentFont,
  RichEdit_SetParentFont,
  RichEdit_GetParentShowHint,
  RichEdit_SetParentShowHint,
  RichEdit_GetPlainText,
  RichEdit_SetPlainText,
  RichEdit_GetPopupMenu,
  RichEdit_SetPopupMenu,
  RichEdit_GetReadOnly,
  RichEdit_SetReadOnly,
  RichEdit_GetScrollBars,
  RichEdit_SetScrollBars,
  RichEdit_GetShowHint,
  RichEdit_SetShowHint,
  RichEdit_GetTabOrder,
  RichEdit_SetTabOrder,
  RichEdit_GetTabStop,
  RichEdit_SetTabStop,
  RichEdit_GetVisible,
  RichEdit_SetVisible,
  RichEdit_GetWantTabs,
  RichEdit_SetWantTabs,
  RichEdit_GetWantReturns,
  RichEdit_SetWantReturns,
  RichEdit_GetWordWrap,
  RichEdit_SetWordWrap,
  RichEdit_GetStyleElements,
  RichEdit_SetStyleElements,
  RichEdit_GetZoom,
  RichEdit_SetZoom,
  RichEdit_SetOnChange,
  RichEdit_SetOnClick,
  RichEdit_SetOnContextPopup,
  RichEdit_SetOnDblClick,
  RichEdit_SetOnDragDrop,
  RichEdit_SetOnDragOver,
  RichEdit_SetOnEndDock,
  RichEdit_SetOnEndDrag,
  RichEdit_SetOnEnter,
  RichEdit_SetOnExit,
  RichEdit_SetOnKeyDown,
  RichEdit_SetOnKeyPress,
  RichEdit_SetOnKeyUp,
  RichEdit_SetOnMouseDown,
  RichEdit_SetOnMouseEnter,
  RichEdit_SetOnMouseLeave,
  RichEdit_SetOnMouseMove,
  RichEdit_SetOnMouseUp,
  RichEdit_SetOnMouseWheel,
  RichEdit_SetOnMouseWheelDown,
  RichEdit_SetOnMouseWheelUp,
  RichEdit_SetOnStartDock,
  RichEdit_GetActiveLineNo,
  RichEdit_GetDefAttributes,
  RichEdit_SetDefAttributes,
  RichEdit_GetSelAttributes,
  RichEdit_SetSelAttributes,
  RichEdit_GetPageRect,
  RichEdit_SetPageRect,
  RichEdit_GetParagraph,
  RichEdit_GetCaretPos,
  RichEdit_SetCaretPos,
  RichEdit_GetCanUndo,
  RichEdit_GetModified,
  RichEdit_SetModified,
  RichEdit_GetSelLength,
  RichEdit_SetSelLength,
  RichEdit_GetSelStart,
  RichEdit_SetSelStart,
  RichEdit_GetSelText,
  RichEdit_SetSelText,
  RichEdit_GetText,
  RichEdit_SetText,
  RichEdit_GetTextHint,
  RichEdit_SetTextHint,
  RichEdit_GetDockSite,
  RichEdit_SetDockSite,
  RichEdit_GetDoubleBuffered,
  RichEdit_SetDoubleBuffered,
  RichEdit_GetBrush,
  RichEdit_GetControlCount,
  RichEdit_GetHandle,
  RichEdit_GetParentDoubleBuffered,
  RichEdit_SetParentDoubleBuffered,
  RichEdit_GetParentWindow,
  RichEdit_SetParentWindow,
  RichEdit_GetUseDockManager,
  RichEdit_SetUseDockManager,
  RichEdit_GetAction,
  RichEdit_SetAction,
  RichEdit_GetBoundsRect,
  RichEdit_SetBoundsRect,
  RichEdit_GetClientHeight,
  RichEdit_SetClientHeight,
  RichEdit_GetClientRect,
  RichEdit_GetClientWidth,
  RichEdit_SetClientWidth,
  RichEdit_GetExplicitLeft,
  RichEdit_GetExplicitTop,
  RichEdit_GetExplicitWidth,
  RichEdit_GetExplicitHeight,
  RichEdit_GetFloating,
  RichEdit_GetParent,
  RichEdit_SetParent,
  RichEdit_GetAlignWithMargins,
  RichEdit_SetAlignWithMargins,
  RichEdit_GetLeft,
  RichEdit_SetLeft,
  RichEdit_GetTop,
  RichEdit_SetTop,
  RichEdit_GetWidth,
  RichEdit_SetWidth,
  RichEdit_GetHeight,
  RichEdit_SetHeight,
  RichEdit_GetCursor,
  RichEdit_SetCursor,
  RichEdit_GetHint,
  RichEdit_SetHint,
  RichEdit_GetMargins,
  RichEdit_SetMargins,
  RichEdit_GetCustomHint,
  RichEdit_SetCustomHint,
  RichEdit_GetComponentCount,
  RichEdit_GetComponentIndex,
  RichEdit_SetComponentIndex,
  RichEdit_GetOwner,
  RichEdit_GetName,
  RichEdit_SetName,
  RichEdit_GetTag,
  RichEdit_SetTag,
  RichEdit_GetControls,
  RichEdit_GetComponents;

