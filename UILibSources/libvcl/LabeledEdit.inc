
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function LabeledEdit_Create(AOwner: TComponent): TLabeledEdit; stdcall;
begin
  Result :=  TLabeledEdit.Create(AOwner);
end;

procedure LabeledEdit_Free(AObj: TLabeledEdit); stdcall;
begin
  AObj.Free;
end;

procedure LabeledEdit_SetBounds(AObj: TLabeledEdit; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure LabeledEdit_SetupInternalLabel(AObj: TLabeledEdit); stdcall;
begin
  AObj.SetupInternalLabel;
end;

procedure LabeledEdit_Clear(AObj: TLabeledEdit); stdcall;
begin
  AObj.Clear;
end;

procedure LabeledEdit_ClearSelection(AObj: TLabeledEdit); stdcall;
begin
  AObj.ClearSelection;
end;

procedure LabeledEdit_CopyToClipboard(AObj: TLabeledEdit); stdcall;
begin
  AObj.CopyToClipboard;
end;

procedure LabeledEdit_CutToClipboard(AObj: TLabeledEdit); stdcall;
begin
  AObj.CutToClipboard;
end;

procedure LabeledEdit_PasteFromClipboard(AObj: TLabeledEdit); stdcall;
begin
  AObj.PasteFromClipboard;
end;

procedure LabeledEdit_Undo(AObj: TLabeledEdit); stdcall;
begin
  AObj.Undo;
end;

procedure LabeledEdit_ClearUndo(AObj: TLabeledEdit); stdcall;
begin
  AObj.ClearUndo;
end;

procedure LabeledEdit_SelectAll(AObj: TLabeledEdit); stdcall;
begin
  AObj.SelectAll;
end;

function LabeledEdit_GetSelTextBuf(AObj: TLabeledEdit; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetSelTextBuf(Buffer, BufSize);
end;

function LabeledEdit_CanFocus(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure LabeledEdit_FlipChildren(AObj: TLabeledEdit; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function LabeledEdit_Focused(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function LabeledEdit_HandleAllocated(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure LabeledEdit_Invalidate(AObj: TLabeledEdit); stdcall;
begin
  AObj.Invalidate;
end;

procedure LabeledEdit_Realign(AObj: TLabeledEdit); stdcall;
begin
  AObj.Realign;
end;

procedure LabeledEdit_Repaint(AObj: TLabeledEdit); stdcall;
begin
  AObj.Repaint;
end;

procedure LabeledEdit_ScaleBy(AObj: TLabeledEdit; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure LabeledEdit_SetFocus(AObj: TLabeledEdit); stdcall;
begin
  AObj.SetFocus;
end;

procedure LabeledEdit_Update(AObj: TLabeledEdit); stdcall;
begin
  AObj.Update;
end;

procedure LabeledEdit_BringToFront(AObj: TLabeledEdit); stdcall;
begin
  AObj.BringToFront;
end;

procedure LabeledEdit_ClientToScreen(AObj: TLabeledEdit; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure LabeledEdit_ClientToParent(AObj: TLabeledEdit; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function LabeledEdit_Dragging(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function LabeledEdit_HasParent(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure LabeledEdit_Hide(AObj: TLabeledEdit); stdcall;
begin
  AObj.Hide;
end;

function LabeledEdit_Perform(AObj: TLabeledEdit; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure LabeledEdit_Refresh(AObj: TLabeledEdit); stdcall;
begin
  AObj.Refresh;
end;

procedure LabeledEdit_ScreenToClient(AObj: TLabeledEdit; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure LabeledEdit_ParentToClient(AObj: TLabeledEdit; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure LabeledEdit_SendToBack(AObj: TLabeledEdit); stdcall;
begin
  AObj.SendToBack;
end;

procedure LabeledEdit_Show(AObj: TLabeledEdit); stdcall;
begin
  AObj.Show;
end;

function LabeledEdit_GetTextBuf(AObj: TLabeledEdit; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function LabeledEdit_GetTextLen(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function LabeledEdit_FindComponent(AObj: TLabeledEdit; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function LabeledEdit_GetNamePath(AObj: TLabeledEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure LabeledEdit_Assign(AObj: TLabeledEdit; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure LabeledEdit_DisposeOf(AObj: TLabeledEdit); stdcall;
begin
  AObj.DisposeOf;
end;

function LabeledEdit_ClassType(AObj: TLabeledEdit): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function LabeledEdit_ClassName(AObj: TLabeledEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function LabeledEdit_InstanceSize(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function LabeledEdit_InheritsFrom(AObj: TLabeledEdit; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function LabeledEdit_Equals(AObj: TLabeledEdit; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function LabeledEdit_GetHashCode(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function LabeledEdit_ToString(AObj: TLabeledEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function LabeledEdit_GetAlignment(AObj: TLabeledEdit): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure LabeledEdit_SetAlignment(AObj: TLabeledEdit; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function LabeledEdit_GetAnchors(AObj: TLabeledEdit): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure LabeledEdit_SetAnchors(AObj: TLabeledEdit; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function LabeledEdit_GetAutoSelect(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.AutoSelect;
end;

procedure LabeledEdit_SetAutoSelect(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.AutoSelect := AValue;
end;

function LabeledEdit_GetAutoSize(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure LabeledEdit_SetAutoSize(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function LabeledEdit_GetBevelEdges(AObj: TLabeledEdit): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure LabeledEdit_SetBevelEdges(AObj: TLabeledEdit; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function LabeledEdit_GetBevelInner(AObj: TLabeledEdit): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure LabeledEdit_SetBevelInner(AObj: TLabeledEdit; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function LabeledEdit_GetBevelKind(AObj: TLabeledEdit): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure LabeledEdit_SetBevelKind(AObj: TLabeledEdit; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function LabeledEdit_GetBevelOuter(AObj: TLabeledEdit): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure LabeledEdit_SetBevelOuter(AObj: TLabeledEdit; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function LabeledEdit_GetBiDiMode(AObj: TLabeledEdit): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure LabeledEdit_SetBiDiMode(AObj: TLabeledEdit; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function LabeledEdit_GetBorderStyle(AObj: TLabeledEdit): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure LabeledEdit_SetBorderStyle(AObj: TLabeledEdit; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function LabeledEdit_GetCharCase(AObj: TLabeledEdit): TEditCharCase; stdcall;
begin
  Result :=  AObj.CharCase;
end;

procedure LabeledEdit_SetCharCase(AObj: TLabeledEdit; AValue: TEditCharCase); stdcall;
begin
  AObj.CharCase := AValue;
end;

function LabeledEdit_GetColor(AObj: TLabeledEdit): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure LabeledEdit_SetColor(AObj: TLabeledEdit; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function LabeledEdit_GetDoubleBuffered(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure LabeledEdit_SetDoubleBuffered(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function LabeledEdit_GetDragCursor(AObj: TLabeledEdit): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure LabeledEdit_SetDragCursor(AObj: TLabeledEdit; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function LabeledEdit_GetDragKind(AObj: TLabeledEdit): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure LabeledEdit_SetDragKind(AObj: TLabeledEdit; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function LabeledEdit_GetDragMode(AObj: TLabeledEdit): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure LabeledEdit_SetDragMode(AObj: TLabeledEdit; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function LabeledEdit_GetEditLabel(AObj: TLabeledEdit): TBoundLabel; stdcall;
begin
  Result :=  AObj.EditLabel;
end;

function LabeledEdit_GetEnabled(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure LabeledEdit_SetEnabled(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function LabeledEdit_GetFont(AObj: TLabeledEdit): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure LabeledEdit_SetFont(AObj: TLabeledEdit; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function LabeledEdit_GetHideSelection(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.HideSelection;
end;

procedure LabeledEdit_SetHideSelection(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.HideSelection := AValue;
end;

function LabeledEdit_GetLabelPosition(AObj: TLabeledEdit): TLabelPosition; stdcall;
begin
  Result :=  AObj.LabelPosition;
end;

procedure LabeledEdit_SetLabelPosition(AObj: TLabeledEdit; AValue: TLabelPosition); stdcall;
begin
  AObj.LabelPosition := AValue;
end;

function LabeledEdit_GetLabelSpacing(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.LabelSpacing;
end;

procedure LabeledEdit_SetLabelSpacing(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.LabelSpacing := AValue;
end;

function LabeledEdit_GetMaxLength(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.MaxLength;
end;

procedure LabeledEdit_SetMaxLength(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.MaxLength := AValue;
end;

function LabeledEdit_GetNumbersOnly(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.NumbersOnly;
end;

procedure LabeledEdit_SetNumbersOnly(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.NumbersOnly := AValue;
end;

function LabeledEdit_GetParentColor(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure LabeledEdit_SetParentColor(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function LabeledEdit_GetParentCtl3D(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure LabeledEdit_SetParentCtl3D(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function LabeledEdit_GetParentDoubleBuffered(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure LabeledEdit_SetParentDoubleBuffered(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function LabeledEdit_GetParentFont(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure LabeledEdit_SetParentFont(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function LabeledEdit_GetParentShowHint(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure LabeledEdit_SetParentShowHint(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function LabeledEdit_GetPasswordChar(AObj: TLabeledEdit): Char; stdcall;
begin
  Result :=  AObj.PasswordChar;
end;

procedure LabeledEdit_SetPasswordChar(AObj: TLabeledEdit; AValue: Char); stdcall;
begin
  AObj.PasswordChar := AValue;
end;

function LabeledEdit_GetPopupMenu(AObj: TLabeledEdit): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure LabeledEdit_SetPopupMenu(AObj: TLabeledEdit; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function LabeledEdit_GetReadOnly(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.ReadOnly;
end;

procedure LabeledEdit_SetReadOnly(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.ReadOnly := AValue;
end;

function LabeledEdit_GetShowHint(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure LabeledEdit_SetShowHint(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function LabeledEdit_GetTabOrder(AObj: TLabeledEdit): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure LabeledEdit_SetTabOrder(AObj: TLabeledEdit; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function LabeledEdit_GetTabStop(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure LabeledEdit_SetTabStop(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function LabeledEdit_GetText(AObj: TLabeledEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure LabeledEdit_SetText(AObj: TLabeledEdit; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function LabeledEdit_GetTextHint(AObj: TLabeledEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.TextHint);
end;

procedure LabeledEdit_SetTextHint(AObj: TLabeledEdit; AValue: PChar); stdcall;
begin
  AObj.TextHint := AValue;
end;

function LabeledEdit_GetVisible(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure LabeledEdit_SetVisible(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function LabeledEdit_GetStyleElements(AObj: TLabeledEdit): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure LabeledEdit_SetStyleElements(AObj: TLabeledEdit; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure LabeledEdit_SetOnChange(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnClick(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnContextPopup(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnDblClick(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnDragDrop(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnDragOver(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnEndDock(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnEndDrag(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnEnter(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnExit(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnKeyDown(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnKeyPress(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnKeyUp(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnMouseDown(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnMouseEnter(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnMouseLeave(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnMouseMove(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnMouseUp(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

procedure LabeledEdit_SetOnStartDock(AObj: TLabeledEdit; AEventId: NativeUInt); stdcall;
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

function LabeledEdit_GetCanUndo(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.CanUndo;
end;

function LabeledEdit_GetModified(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure LabeledEdit_SetModified(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function LabeledEdit_GetSelLength(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.SelLength;
end;

procedure LabeledEdit_SetSelLength(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.SelLength := AValue;
end;

function LabeledEdit_GetSelStart(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.SelStart;
end;

procedure LabeledEdit_SetSelStart(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.SelStart := AValue;
end;

function LabeledEdit_GetSelText(AObj: TLabeledEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure LabeledEdit_SetSelText(AObj: TLabeledEdit; AValue: PChar); stdcall;
begin
  AObj.SelText := AValue;
end;

function LabeledEdit_GetDockSite(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure LabeledEdit_SetDockSite(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function LabeledEdit_GetBrush(AObj: TLabeledEdit): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function LabeledEdit_GetControlCount(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function LabeledEdit_GetHandle(AObj: TLabeledEdit): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function LabeledEdit_GetParentWindow(AObj: TLabeledEdit): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure LabeledEdit_SetParentWindow(AObj: TLabeledEdit; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function LabeledEdit_GetUseDockManager(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure LabeledEdit_SetUseDockManager(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function LabeledEdit_GetAction(AObj: TLabeledEdit): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure LabeledEdit_SetAction(AObj: TLabeledEdit; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function LabeledEdit_GetAlign(AObj: TLabeledEdit): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure LabeledEdit_SetAlign(AObj: TLabeledEdit; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

procedure LabeledEdit_GetBoundsRect(AObj: TLabeledEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure LabeledEdit_SetBoundsRect(AObj: TLabeledEdit; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function LabeledEdit_GetClientHeight(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure LabeledEdit_SetClientHeight(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure LabeledEdit_GetClientRect(AObj: TLabeledEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function LabeledEdit_GetClientWidth(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure LabeledEdit_SetClientWidth(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function LabeledEdit_GetExplicitLeft(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function LabeledEdit_GetExplicitTop(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function LabeledEdit_GetExplicitWidth(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function LabeledEdit_GetExplicitHeight(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function LabeledEdit_GetFloating(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function LabeledEdit_GetParent(AObj: TLabeledEdit): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure LabeledEdit_SetParent(AObj: TLabeledEdit; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function LabeledEdit_GetAlignWithMargins(AObj: TLabeledEdit): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure LabeledEdit_SetAlignWithMargins(AObj: TLabeledEdit; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function LabeledEdit_GetLeft(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure LabeledEdit_SetLeft(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function LabeledEdit_GetTop(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure LabeledEdit_SetTop(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function LabeledEdit_GetWidth(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure LabeledEdit_SetWidth(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function LabeledEdit_GetHeight(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure LabeledEdit_SetHeight(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function LabeledEdit_GetCursor(AObj: TLabeledEdit): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure LabeledEdit_SetCursor(AObj: TLabeledEdit; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function LabeledEdit_GetHint(AObj: TLabeledEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure LabeledEdit_SetHint(AObj: TLabeledEdit; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function LabeledEdit_GetMargins(AObj: TLabeledEdit): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure LabeledEdit_SetMargins(AObj: TLabeledEdit; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function LabeledEdit_GetCustomHint(AObj: TLabeledEdit): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure LabeledEdit_SetCustomHint(AObj: TLabeledEdit; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function LabeledEdit_GetComponentCount(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function LabeledEdit_GetComponentIndex(AObj: TLabeledEdit): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure LabeledEdit_SetComponentIndex(AObj: TLabeledEdit; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function LabeledEdit_GetOwner(AObj: TLabeledEdit): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function LabeledEdit_GetName(AObj: TLabeledEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure LabeledEdit_SetName(AObj: TLabeledEdit; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function LabeledEdit_GetTag(AObj: TLabeledEdit): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure LabeledEdit_SetTag(AObj: TLabeledEdit; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function LabeledEdit_GetControls(AObj: TLabeledEdit; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function LabeledEdit_GetComponents(AObj: TLabeledEdit; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function LabeledEdit_StaticClassType: TClass; stdcall;
begin
  Result := TLabeledEdit;
end;

exports
  LabeledEdit_Create,
  LabeledEdit_Free,
  LabeledEdit_SetBounds,
  LabeledEdit_SetupInternalLabel,
  LabeledEdit_Clear,
  LabeledEdit_ClearSelection,
  LabeledEdit_CopyToClipboard,
  LabeledEdit_CutToClipboard,
  LabeledEdit_PasteFromClipboard,
  LabeledEdit_Undo,
  LabeledEdit_ClearUndo,
  LabeledEdit_SelectAll,
  LabeledEdit_GetSelTextBuf,
  LabeledEdit_CanFocus,
  LabeledEdit_FlipChildren,
  LabeledEdit_Focused,
  LabeledEdit_HandleAllocated,
  LabeledEdit_Invalidate,
  LabeledEdit_Realign,
  LabeledEdit_Repaint,
  LabeledEdit_ScaleBy,
  LabeledEdit_SetFocus,
  LabeledEdit_Update,
  LabeledEdit_BringToFront,
  LabeledEdit_ClientToScreen,
  LabeledEdit_ClientToParent,
  LabeledEdit_Dragging,
  LabeledEdit_HasParent,
  LabeledEdit_Hide,
  LabeledEdit_Perform,
  LabeledEdit_Refresh,
  LabeledEdit_ScreenToClient,
  LabeledEdit_ParentToClient,
  LabeledEdit_SendToBack,
  LabeledEdit_Show,
  LabeledEdit_GetTextBuf,
  LabeledEdit_GetTextLen,
  LabeledEdit_FindComponent,
  LabeledEdit_GetNamePath,
  LabeledEdit_Assign,
  LabeledEdit_DisposeOf,
  LabeledEdit_ClassType,
  LabeledEdit_ClassName,
  LabeledEdit_InstanceSize,
  LabeledEdit_InheritsFrom,
  LabeledEdit_Equals,
  LabeledEdit_GetHashCode,
  LabeledEdit_ToString,
  LabeledEdit_GetAlignment,
  LabeledEdit_SetAlignment,
  LabeledEdit_GetAnchors,
  LabeledEdit_SetAnchors,
  LabeledEdit_GetAutoSelect,
  LabeledEdit_SetAutoSelect,
  LabeledEdit_GetAutoSize,
  LabeledEdit_SetAutoSize,
  LabeledEdit_GetBevelEdges,
  LabeledEdit_SetBevelEdges,
  LabeledEdit_GetBevelInner,
  LabeledEdit_SetBevelInner,
  LabeledEdit_GetBevelKind,
  LabeledEdit_SetBevelKind,
  LabeledEdit_GetBevelOuter,
  LabeledEdit_SetBevelOuter,
  LabeledEdit_GetBiDiMode,
  LabeledEdit_SetBiDiMode,
  LabeledEdit_GetBorderStyle,
  LabeledEdit_SetBorderStyle,
  LabeledEdit_GetCharCase,
  LabeledEdit_SetCharCase,
  LabeledEdit_GetColor,
  LabeledEdit_SetColor,
  LabeledEdit_GetDoubleBuffered,
  LabeledEdit_SetDoubleBuffered,
  LabeledEdit_GetDragCursor,
  LabeledEdit_SetDragCursor,
  LabeledEdit_GetDragKind,
  LabeledEdit_SetDragKind,
  LabeledEdit_GetDragMode,
  LabeledEdit_SetDragMode,
  LabeledEdit_GetEditLabel,
  LabeledEdit_GetEnabled,
  LabeledEdit_SetEnabled,
  LabeledEdit_GetFont,
  LabeledEdit_SetFont,
  LabeledEdit_GetHideSelection,
  LabeledEdit_SetHideSelection,
  LabeledEdit_GetLabelPosition,
  LabeledEdit_SetLabelPosition,
  LabeledEdit_GetLabelSpacing,
  LabeledEdit_SetLabelSpacing,
  LabeledEdit_GetMaxLength,
  LabeledEdit_SetMaxLength,
  LabeledEdit_GetNumbersOnly,
  LabeledEdit_SetNumbersOnly,
  LabeledEdit_GetParentColor,
  LabeledEdit_SetParentColor,
  LabeledEdit_GetParentCtl3D,
  LabeledEdit_SetParentCtl3D,
  LabeledEdit_GetParentDoubleBuffered,
  LabeledEdit_SetParentDoubleBuffered,
  LabeledEdit_GetParentFont,
  LabeledEdit_SetParentFont,
  LabeledEdit_GetParentShowHint,
  LabeledEdit_SetParentShowHint,
  LabeledEdit_GetPasswordChar,
  LabeledEdit_SetPasswordChar,
  LabeledEdit_GetPopupMenu,
  LabeledEdit_SetPopupMenu,
  LabeledEdit_GetReadOnly,
  LabeledEdit_SetReadOnly,
  LabeledEdit_GetShowHint,
  LabeledEdit_SetShowHint,
  LabeledEdit_GetTabOrder,
  LabeledEdit_SetTabOrder,
  LabeledEdit_GetTabStop,
  LabeledEdit_SetTabStop,
  LabeledEdit_GetText,
  LabeledEdit_SetText,
  LabeledEdit_GetTextHint,
  LabeledEdit_SetTextHint,
  LabeledEdit_GetVisible,
  LabeledEdit_SetVisible,
  LabeledEdit_GetStyleElements,
  LabeledEdit_SetStyleElements,
  LabeledEdit_SetOnChange,
  LabeledEdit_SetOnClick,
  LabeledEdit_SetOnContextPopup,
  LabeledEdit_SetOnDblClick,
  LabeledEdit_SetOnDragDrop,
  LabeledEdit_SetOnDragOver,
  LabeledEdit_SetOnEndDock,
  LabeledEdit_SetOnEndDrag,
  LabeledEdit_SetOnEnter,
  LabeledEdit_SetOnExit,
  LabeledEdit_SetOnKeyDown,
  LabeledEdit_SetOnKeyPress,
  LabeledEdit_SetOnKeyUp,
  LabeledEdit_SetOnMouseDown,
  LabeledEdit_SetOnMouseEnter,
  LabeledEdit_SetOnMouseLeave,
  LabeledEdit_SetOnMouseMove,
  LabeledEdit_SetOnMouseUp,
  LabeledEdit_SetOnStartDock,
  LabeledEdit_GetCanUndo,
  LabeledEdit_GetModified,
  LabeledEdit_SetModified,
  LabeledEdit_GetSelLength,
  LabeledEdit_SetSelLength,
  LabeledEdit_GetSelStart,
  LabeledEdit_SetSelStart,
  LabeledEdit_GetSelText,
  LabeledEdit_SetSelText,
  LabeledEdit_GetDockSite,
  LabeledEdit_SetDockSite,
  LabeledEdit_GetBrush,
  LabeledEdit_GetControlCount,
  LabeledEdit_GetHandle,
  LabeledEdit_GetParentWindow,
  LabeledEdit_SetParentWindow,
  LabeledEdit_GetUseDockManager,
  LabeledEdit_SetUseDockManager,
  LabeledEdit_GetAction,
  LabeledEdit_SetAction,
  LabeledEdit_GetAlign,
  LabeledEdit_SetAlign,
  LabeledEdit_GetBoundsRect,
  LabeledEdit_SetBoundsRect,
  LabeledEdit_GetClientHeight,
  LabeledEdit_SetClientHeight,
  LabeledEdit_GetClientRect,
  LabeledEdit_GetClientWidth,
  LabeledEdit_SetClientWidth,
  LabeledEdit_GetExplicitLeft,
  LabeledEdit_GetExplicitTop,
  LabeledEdit_GetExplicitWidth,
  LabeledEdit_GetExplicitHeight,
  LabeledEdit_GetFloating,
  LabeledEdit_GetParent,
  LabeledEdit_SetParent,
  LabeledEdit_GetAlignWithMargins,
  LabeledEdit_SetAlignWithMargins,
  LabeledEdit_GetLeft,
  LabeledEdit_SetLeft,
  LabeledEdit_GetTop,
  LabeledEdit_SetTop,
  LabeledEdit_GetWidth,
  LabeledEdit_SetWidth,
  LabeledEdit_GetHeight,
  LabeledEdit_SetHeight,
  LabeledEdit_GetCursor,
  LabeledEdit_SetCursor,
  LabeledEdit_GetHint,
  LabeledEdit_SetHint,
  LabeledEdit_GetMargins,
  LabeledEdit_SetMargins,
  LabeledEdit_GetCustomHint,
  LabeledEdit_SetCustomHint,
  LabeledEdit_GetComponentCount,
  LabeledEdit_GetComponentIndex,
  LabeledEdit_SetComponentIndex,
  LabeledEdit_GetOwner,
  LabeledEdit_GetName,
  LabeledEdit_SetName,
  LabeledEdit_GetTag,
  LabeledEdit_SetTag,
  LabeledEdit_GetControls,
  LabeledEdit_GetComponents,
  LabeledEdit_StaticClassType;

