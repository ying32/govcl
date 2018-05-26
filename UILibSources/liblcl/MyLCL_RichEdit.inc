
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function RichEdit_Create(AOwner: TComponent): TRichEdit; extdecl;
begin
  Result :=  TRichEdit.Create(AOwner);
end;

procedure RichEdit_Free(AObj: TRichEdit); extdecl;
begin
  AObj.Free;
end;

procedure RichEdit_Clear(AObj: TRichEdit); extdecl;
begin
  AObj.Clear;
end;

function RichEdit_FindText(AObj: TRichEdit; SearchStr: PChar; StartPos: Integer; Length: Integer; Options: TSearchTypes): Integer; extdecl;
begin
  Result :=  AObj.FindText(SearchStr, StartPos, Length, Options);
end;

procedure RichEdit_ClearSelection(AObj: TRichEdit); extdecl;
begin
  AObj.ClearSelection;
end;

procedure RichEdit_CopyToClipboard(AObj: TRichEdit); extdecl;
begin
  AObj.CopyToClipboard;
end;

procedure RichEdit_CutToClipboard(AObj: TRichEdit); extdecl;
begin
  AObj.CutToClipboard;
end;

procedure RichEdit_PasteFromClipboard(AObj: TRichEdit); extdecl;
begin
  AObj.PasteFromClipboard;
end;

procedure RichEdit_SelectAll(AObj: TRichEdit); extdecl;
begin
  AObj.SelectAll;
end;

function RichEdit_CanFocus(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure RichEdit_FlipChildren(AObj: TRichEdit; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function RichEdit_Focused(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

function RichEdit_HandleAllocated(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure RichEdit_Invalidate(AObj: TRichEdit); extdecl;
begin
  AObj.Invalidate;
end;

procedure RichEdit_Realign(AObj: TRichEdit); extdecl;
begin
  AObj.Realign;
end;

procedure RichEdit_Repaint(AObj: TRichEdit); extdecl;
begin
  AObj.Repaint;
end;

procedure RichEdit_ScaleBy(AObj: TRichEdit; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure RichEdit_SetBounds(AObj: TRichEdit; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure RichEdit_SetFocus(AObj: TRichEdit); extdecl;
begin
  AObj.SetFocus;
end;

procedure RichEdit_Update(AObj: TRichEdit); extdecl;
begin
  AObj.Update;
end;

procedure RichEdit_BringToFront(AObj: TRichEdit); extdecl;
begin
  AObj.BringToFront;
end;

function RichEdit_HasParent(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure RichEdit_Hide(AObj: TRichEdit); extdecl;
begin
  AObj.Hide;
end;

function RichEdit_Perform(AObj: TRichEdit; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure RichEdit_Refresh(AObj: TRichEdit); extdecl;
begin
  AObj.Refresh;
end;

procedure RichEdit_SendToBack(AObj: TRichEdit); extdecl;
begin
  AObj.SendToBack;
end;

procedure RichEdit_Show(AObj: TRichEdit); extdecl;
begin
  AObj.Show;
end;

function RichEdit_GetTextBuf(AObj: TRichEdit; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function RichEdit_FindComponent(AObj: TRichEdit; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function RichEdit_GetNamePath(AObj: TRichEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure RichEdit_Assign(AObj: TRichEdit; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function RichEdit_Equals(AObj: TRichEdit; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function RichEdit_GetHashCode(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function RichEdit_ToString(AObj: TRichEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function RichEdit_GetAlign(AObj: TRichEdit): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure RichEdit_SetAlign(AObj: TRichEdit; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function RichEdit_GetAlignment(AObj: TRichEdit): TAlignment; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure RichEdit_SetAlignment(AObj: TRichEdit; AValue: TAlignment); extdecl;
begin
  AObj.Alignment := AValue;
end;

function RichEdit_GetAnchors(AObj: TRichEdit): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure RichEdit_SetAnchors(AObj: TRichEdit; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function RichEdit_GetBiDiMode(AObj: TRichEdit): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure RichEdit_SetBiDiMode(AObj: TRichEdit; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function RichEdit_GetBorderStyle(AObj: TRichEdit): TBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure RichEdit_SetBorderStyle(AObj: TRichEdit; AValue: TBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function RichEdit_GetBorderWidth(AObj: TRichEdit): TBorderWidth; extdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure RichEdit_SetBorderWidth(AObj: TRichEdit; AValue: TBorderWidth); extdecl;
begin
  AObj.BorderWidth := AValue;
end;

function RichEdit_GetColor(AObj: TRichEdit): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure RichEdit_SetColor(AObj: TRichEdit; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function RichEdit_GetEnabled(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure RichEdit_SetEnabled(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function RichEdit_GetFont(AObj: TRichEdit): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure RichEdit_SetFont(AObj: TRichEdit; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function RichEdit_GetHideSelection(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.HideSelection;
end;

procedure RichEdit_SetHideSelection(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.HideSelection := AValue;
end;

function RichEdit_GetHideScrollBars(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.HideScrollBars;
end;

procedure RichEdit_SetHideScrollBars(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.HideScrollBars := AValue;
end;

function RichEdit_GetLines(AObj: TRichEdit): TStrings; extdecl;
begin
  Result :=  AObj.Lines;
end;

procedure RichEdit_SetLines(AObj: TRichEdit; AValue: TStrings); extdecl;
begin
  AObj.Lines := AValue;
end;

function RichEdit_GetMaxLength(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.MaxLength;
end;

procedure RichEdit_SetMaxLength(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.MaxLength := AValue;
end;

function RichEdit_GetParentColor(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure RichEdit_SetParentColor(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function RichEdit_GetParentFont(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure RichEdit_SetParentFont(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function RichEdit_GetParentShowHint(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure RichEdit_SetParentShowHint(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function RichEdit_GetPlainText(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.PlainText;
end;

procedure RichEdit_SetPlainText(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.PlainText := AValue;
end;

function RichEdit_GetPopupMenu(AObj: TRichEdit): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure RichEdit_SetPopupMenu(AObj: TRichEdit; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function RichEdit_GetReadOnly(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.ReadOnly;
end;

procedure RichEdit_SetReadOnly(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.ReadOnly := AValue;
end;

function RichEdit_GetScrollBars(AObj: TRichEdit): TScrollStyle; extdecl;
begin
  Result :=  AObj.ScrollBars;
end;

procedure RichEdit_SetScrollBars(AObj: TRichEdit; AValue: TScrollStyle); extdecl;
begin
  AObj.ScrollBars := AValue;
end;

function RichEdit_GetShowHint(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure RichEdit_SetShowHint(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function RichEdit_GetTabOrder(AObj: TRichEdit): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure RichEdit_SetTabOrder(AObj: TRichEdit; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function RichEdit_GetTabStop(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure RichEdit_SetTabStop(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function RichEdit_GetVisible(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure RichEdit_SetVisible(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function RichEdit_GetWantTabs(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.WantTabs;
end;

procedure RichEdit_SetWantTabs(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.WantTabs := AValue;
end;

function RichEdit_GetWantReturns(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.WantReturns;
end;

procedure RichEdit_SetWantReturns(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.WantReturns := AValue;
end;

function RichEdit_GetWordWrap(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.WordWrap;
end;

procedure RichEdit_SetWordWrap(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.WordWrap := AValue;
end;

function RichEdit_GetZoom(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.Zoom;
end;

procedure RichEdit_SetZoom(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.Zoom := AValue;
end;

procedure RichEdit_SetOnChange(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnClick(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnDblClick(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnEnter(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnExit(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnKeyDown(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnKeyPress(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnKeyUp(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnMouseDown(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnMouseEnter(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnMouseLeave(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnMouseMove(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnMouseUp(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
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

procedure RichEdit_SetOnMouseWheel(AObj: TRichEdit; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheel := nil;
    EventClass.Remove(AObj, geMouseWheel);
    Exit;
  end;
  AObj.OnMouseWheel := @EventClass.OnMouseWheel;
  EventClass.Add(AObj, geMouseWheel, AEventId);
end;

function RichEdit_GetActiveLineNo(AObj: TRichEdit): Cardinal; extdecl;
begin
  Result :=  AObj.ActiveLineNo;
end;

function RichEdit_GetDefAttributes(AObj: TRichEdit): TTextAttributes; extdecl;
begin
  Result :=  AObj.DefAttributes;
end;

procedure RichEdit_SetDefAttributes(AObj: TRichEdit; AValue: TTextAttributes); extdecl;
begin
  AObj.DefAttributes := AValue;
end;

function RichEdit_GetSelAttributes(AObj: TRichEdit): TTextAttributes; extdecl;
begin
  Result :=  AObj.SelAttributes;
end;

procedure RichEdit_SetSelAttributes(AObj: TRichEdit; AValue: TTextAttributes); extdecl;
begin
  AObj.SelAttributes := AValue;
end;

procedure RichEdit_GetPageRect(AObj: TRichEdit; var Result: TRect); extdecl;
begin
  Result :=  AObj.PageRect;
end;

procedure RichEdit_SetPageRect(AObj: TRichEdit; var AValue: TRect); extdecl;
begin
  AObj.PageRect := AValue;
end;

function RichEdit_GetParagraph(AObj: TRichEdit): TParaAttributes; extdecl;
begin
  Result :=  AObj.Paragraph;
end;

procedure RichEdit_GetCaretPos(AObj: TRichEdit; var Result: TPoint); extdecl;
begin
  Result :=  AObj.CaretPos;
end;

procedure RichEdit_SetCaretPos(AObj: TRichEdit; var AValue: TPoint); extdecl;
begin
  AObj.CaretPos := AValue;
end;

function RichEdit_GetModified(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.Modified;
end;

procedure RichEdit_SetModified(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.Modified := AValue;
end;

function RichEdit_GetSelLength(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.SelLength;
end;

procedure RichEdit_SetSelLength(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.SelLength := AValue;
end;

function RichEdit_GetSelStart(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.SelStart;
end;

procedure RichEdit_SetSelStart(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.SelStart := AValue;
end;

function RichEdit_GetSelText(AObj: TRichEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure RichEdit_SetSelText(AObj: TRichEdit; AValue: PChar); extdecl;
begin
  AObj.SelText := AValue;
end;

function RichEdit_GetText(AObj: TRichEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.Text);
end;

procedure RichEdit_SetText(AObj: TRichEdit; AValue: PChar); extdecl;
begin
  AObj.Text := AValue;
end;

function RichEdit_GetDoubleBuffered(AObj: TRichEdit): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure RichEdit_SetDoubleBuffered(AObj: TRichEdit; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function RichEdit_GetBrush(AObj: TRichEdit): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function RichEdit_GetControlCount(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function RichEdit_GetHandle(AObj: TRichEdit): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function RichEdit_GetParentWindow(AObj: TRichEdit): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure RichEdit_SetParentWindow(AObj: TRichEdit; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function RichEdit_GetAction(AObj: TRichEdit): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure RichEdit_SetAction(AObj: TRichEdit; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure RichEdit_GetBoundsRect(AObj: TRichEdit; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure RichEdit_SetBoundsRect(AObj: TRichEdit; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function RichEdit_GetClientHeight(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure RichEdit_SetClientHeight(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure RichEdit_GetClientRect(AObj: TRichEdit; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function RichEdit_GetClientWidth(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure RichEdit_SetClientWidth(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function RichEdit_GetParent(AObj: TRichEdit): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure RichEdit_SetParent(AObj: TRichEdit; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function RichEdit_GetLeft(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure RichEdit_SetLeft(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function RichEdit_GetTop(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure RichEdit_SetTop(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function RichEdit_GetWidth(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure RichEdit_SetWidth(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function RichEdit_GetHeight(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure RichEdit_SetHeight(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function RichEdit_GetCursor(AObj: TRichEdit): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure RichEdit_SetCursor(AObj: TRichEdit; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function RichEdit_GetHint(AObj: TRichEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure RichEdit_SetHint(AObj: TRichEdit; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function RichEdit_GetComponentCount(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function RichEdit_GetComponentIndex(AObj: TRichEdit): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure RichEdit_SetComponentIndex(AObj: TRichEdit; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function RichEdit_GetOwner(AObj: TRichEdit): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function RichEdit_GetName(AObj: TRichEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure RichEdit_SetName(AObj: TRichEdit; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function RichEdit_GetTag(AObj: TRichEdit): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure RichEdit_SetTag(AObj: TRichEdit; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function RichEdit_GetControls(AObj: TRichEdit; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function RichEdit_GetComponents(AObj: TRichEdit; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  RichEdit_Create,
  RichEdit_Free,
  RichEdit_Clear,
  RichEdit_FindText,
  RichEdit_ClearSelection,
  RichEdit_CopyToClipboard,
  RichEdit_CutToClipboard,
  RichEdit_PasteFromClipboard,
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
  RichEdit_HasParent,
  RichEdit_Hide,
  RichEdit_Perform,
  RichEdit_Refresh,
  RichEdit_SendToBack,
  RichEdit_Show,
  RichEdit_GetTextBuf,
  RichEdit_FindComponent,
  RichEdit_GetNamePath,
  RichEdit_Assign,
  RichEdit_Equals,
  RichEdit_GetHashCode,
  RichEdit_ToString,
  RichEdit_GetAlign,
  RichEdit_SetAlign,
  RichEdit_GetAlignment,
  RichEdit_SetAlignment,
  RichEdit_GetAnchors,
  RichEdit_SetAnchors,
  RichEdit_GetBiDiMode,
  RichEdit_SetBiDiMode,
  RichEdit_GetBorderStyle,
  RichEdit_SetBorderStyle,
  RichEdit_GetBorderWidth,
  RichEdit_SetBorderWidth,
  RichEdit_GetColor,
  RichEdit_SetColor,
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
  RichEdit_GetZoom,
  RichEdit_SetZoom,
  RichEdit_SetOnChange,
  RichEdit_SetOnClick,
  RichEdit_SetOnDblClick,
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
  RichEdit_GetDoubleBuffered,
  RichEdit_SetDoubleBuffered,
  RichEdit_GetBrush,
  RichEdit_GetControlCount,
  RichEdit_GetHandle,
  RichEdit_GetParentWindow,
  RichEdit_SetParentWindow,
  RichEdit_GetAction,
  RichEdit_SetAction,
  RichEdit_GetBoundsRect,
  RichEdit_SetBoundsRect,
  RichEdit_GetClientHeight,
  RichEdit_SetClientHeight,
  RichEdit_GetClientRect,
  RichEdit_GetClientWidth,
  RichEdit_SetClientWidth,
  RichEdit_GetParent,
  RichEdit_SetParent,
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

