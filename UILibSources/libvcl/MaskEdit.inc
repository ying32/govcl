
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function MaskEdit_Create(AOwner: TComponent): TMaskEdit; stdcall;
begin
  Result :=  TMaskEdit.Create(AOwner);
end;

procedure MaskEdit_Free(AObj: TMaskEdit); stdcall;
begin
  AObj.Free;
end;

procedure MaskEdit_ValidateEdit(AObj: TMaskEdit); stdcall;
begin
  AObj.ValidateEdit;
end;

procedure MaskEdit_Clear(AObj: TMaskEdit); stdcall;
begin
  AObj.Clear;
end;

function MaskEdit_GetTextLen(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure MaskEdit_ClearSelection(AObj: TMaskEdit); stdcall;
begin
  AObj.ClearSelection;
end;

procedure MaskEdit_CopyToClipboard(AObj: TMaskEdit); stdcall;
begin
  AObj.CopyToClipboard;
end;

procedure MaskEdit_CutToClipboard(AObj: TMaskEdit); stdcall;
begin
  AObj.CutToClipboard;
end;

procedure MaskEdit_PasteFromClipboard(AObj: TMaskEdit); stdcall;
begin
  AObj.PasteFromClipboard;
end;

procedure MaskEdit_Undo(AObj: TMaskEdit); stdcall;
begin
  AObj.Undo;
end;

procedure MaskEdit_ClearUndo(AObj: TMaskEdit); stdcall;
begin
  AObj.ClearUndo;
end;

procedure MaskEdit_SelectAll(AObj: TMaskEdit); stdcall;
begin
  AObj.SelectAll;
end;

function MaskEdit_GetSelTextBuf(AObj: TMaskEdit; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetSelTextBuf(Buffer, BufSize);
end;

function MaskEdit_CanFocus(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure MaskEdit_FlipChildren(AObj: TMaskEdit; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function MaskEdit_Focused(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function MaskEdit_HandleAllocated(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure MaskEdit_Invalidate(AObj: TMaskEdit); stdcall;
begin
  AObj.Invalidate;
end;

procedure MaskEdit_Realign(AObj: TMaskEdit); stdcall;
begin
  AObj.Realign;
end;

procedure MaskEdit_Repaint(AObj: TMaskEdit); stdcall;
begin
  AObj.Repaint;
end;

procedure MaskEdit_ScaleBy(AObj: TMaskEdit; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure MaskEdit_SetBounds(AObj: TMaskEdit; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure MaskEdit_SetFocus(AObj: TMaskEdit); stdcall;
begin
  AObj.SetFocus;
end;

procedure MaskEdit_Update(AObj: TMaskEdit); stdcall;
begin
  AObj.Update;
end;

procedure MaskEdit_BringToFront(AObj: TMaskEdit); stdcall;
begin
  AObj.BringToFront;
end;

procedure MaskEdit_ClientToScreen(AObj: TMaskEdit; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure MaskEdit_ClientToParent(AObj: TMaskEdit; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function MaskEdit_Dragging(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function MaskEdit_HasParent(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure MaskEdit_Hide(AObj: TMaskEdit); stdcall;
begin
  AObj.Hide;
end;

function MaskEdit_Perform(AObj: TMaskEdit; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure MaskEdit_Refresh(AObj: TMaskEdit); stdcall;
begin
  AObj.Refresh;
end;

procedure MaskEdit_ScreenToClient(AObj: TMaskEdit; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure MaskEdit_ParentToClient(AObj: TMaskEdit; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure MaskEdit_SendToBack(AObj: TMaskEdit); stdcall;
begin
  AObj.SendToBack;
end;

procedure MaskEdit_Show(AObj: TMaskEdit); stdcall;
begin
  AObj.Show;
end;

function MaskEdit_GetTextBuf(AObj: TMaskEdit; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function MaskEdit_FindComponent(AObj: TMaskEdit; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function MaskEdit_GetNamePath(AObj: TMaskEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure MaskEdit_Assign(AObj: TMaskEdit; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure MaskEdit_DisposeOf(AObj: TMaskEdit); stdcall;
begin
  AObj.DisposeOf;
end;

function MaskEdit_ClassType(AObj: TMaskEdit): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function MaskEdit_ClassName(AObj: TMaskEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function MaskEdit_InstanceSize(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function MaskEdit_InheritsFrom(AObj: TMaskEdit; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function MaskEdit_Equals(AObj: TMaskEdit; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function MaskEdit_GetHashCode(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function MaskEdit_ToString(AObj: TMaskEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function MaskEdit_GetAlign(AObj: TMaskEdit): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure MaskEdit_SetAlign(AObj: TMaskEdit; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function MaskEdit_GetAlignment(AObj: TMaskEdit): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure MaskEdit_SetAlignment(AObj: TMaskEdit; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function MaskEdit_GetAnchors(AObj: TMaskEdit): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure MaskEdit_SetAnchors(AObj: TMaskEdit; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function MaskEdit_GetAutoSelect(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.AutoSelect;
end;

procedure MaskEdit_SetAutoSelect(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.AutoSelect := AValue;
end;

function MaskEdit_GetAutoSize(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure MaskEdit_SetAutoSize(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function MaskEdit_GetBevelEdges(AObj: TMaskEdit): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure MaskEdit_SetBevelEdges(AObj: TMaskEdit; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function MaskEdit_GetBevelInner(AObj: TMaskEdit): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure MaskEdit_SetBevelInner(AObj: TMaskEdit; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function MaskEdit_GetBevelOuter(AObj: TMaskEdit): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure MaskEdit_SetBevelOuter(AObj: TMaskEdit; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function MaskEdit_GetBevelKind(AObj: TMaskEdit): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure MaskEdit_SetBevelKind(AObj: TMaskEdit; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function MaskEdit_GetBiDiMode(AObj: TMaskEdit): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure MaskEdit_SetBiDiMode(AObj: TMaskEdit; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function MaskEdit_GetBorderStyle(AObj: TMaskEdit): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure MaskEdit_SetBorderStyle(AObj: TMaskEdit; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function MaskEdit_GetCharCase(AObj: TMaskEdit): TEditCharCase; stdcall;
begin
  Result :=  AObj.CharCase;
end;

procedure MaskEdit_SetCharCase(AObj: TMaskEdit; AValue: TEditCharCase); stdcall;
begin
  AObj.CharCase := AValue;
end;

function MaskEdit_GetColor(AObj: TMaskEdit): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure MaskEdit_SetColor(AObj: TMaskEdit; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function MaskEdit_GetDoubleBuffered(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure MaskEdit_SetDoubleBuffered(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function MaskEdit_GetDragCursor(AObj: TMaskEdit): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure MaskEdit_SetDragCursor(AObj: TMaskEdit; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function MaskEdit_GetDragKind(AObj: TMaskEdit): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure MaskEdit_SetDragKind(AObj: TMaskEdit; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function MaskEdit_GetDragMode(AObj: TMaskEdit): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure MaskEdit_SetDragMode(AObj: TMaskEdit; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function MaskEdit_GetEnabled(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure MaskEdit_SetEnabled(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function MaskEdit_GetFont(AObj: TMaskEdit): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure MaskEdit_SetFont(AObj: TMaskEdit; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function MaskEdit_GetMaxLength(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.MaxLength;
end;

procedure MaskEdit_SetMaxLength(AObj: TMaskEdit; AValue: Integer); stdcall;
begin
  AObj.MaxLength := AValue;
end;

function MaskEdit_GetParentColor(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure MaskEdit_SetParentColor(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function MaskEdit_GetParentCtl3D(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure MaskEdit_SetParentCtl3D(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function MaskEdit_GetParentDoubleBuffered(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure MaskEdit_SetParentDoubleBuffered(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function MaskEdit_GetParentFont(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure MaskEdit_SetParentFont(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function MaskEdit_GetParentShowHint(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure MaskEdit_SetParentShowHint(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function MaskEdit_GetPasswordChar(AObj: TMaskEdit): Char; stdcall;
begin
  Result :=  AObj.PasswordChar;
end;

procedure MaskEdit_SetPasswordChar(AObj: TMaskEdit; AValue: Char); stdcall;
begin
  AObj.PasswordChar := AValue;
end;

function MaskEdit_GetPopupMenu(AObj: TMaskEdit): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure MaskEdit_SetPopupMenu(AObj: TMaskEdit; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function MaskEdit_GetReadOnly(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.ReadOnly;
end;

procedure MaskEdit_SetReadOnly(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.ReadOnly := AValue;
end;

function MaskEdit_GetShowHint(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure MaskEdit_SetShowHint(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function MaskEdit_GetTabOrder(AObj: TMaskEdit): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure MaskEdit_SetTabOrder(AObj: TMaskEdit; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function MaskEdit_GetTabStop(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure MaskEdit_SetTabStop(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function MaskEdit_GetText(AObj: TMaskEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure MaskEdit_SetText(AObj: TMaskEdit; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function MaskEdit_GetTextHint(AObj: TMaskEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.TextHint);
end;

procedure MaskEdit_SetTextHint(AObj: TMaskEdit; AValue: PChar); stdcall;
begin
  AObj.TextHint := AValue;
end;

function MaskEdit_GetVisible(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure MaskEdit_SetVisible(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function MaskEdit_GetStyleElements(AObj: TMaskEdit): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure MaskEdit_SetStyleElements(AObj: TMaskEdit; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure MaskEdit_SetOnChange(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnClick(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnDblClick(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnDragDrop(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnDragOver(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnEndDock(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnEndDrag(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnEnter(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnExit(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnKeyDown(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnKeyPress(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnKeyUp(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnMouseDown(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnMouseEnter(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnMouseLeave(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnMouseMove(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnMouseUp(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

procedure MaskEdit_SetOnStartDock(AObj: TMaskEdit; AEventId: NativeUInt); stdcall;
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

function MaskEdit_GetIsMasked(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.IsMasked;
end;

function MaskEdit_GetEditText(AObj: TMaskEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.EditText);
end;

procedure MaskEdit_SetEditText(AObj: TMaskEdit; AValue: PChar); stdcall;
begin
  AObj.EditText := AValue;
end;

function MaskEdit_GetCanUndo(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.CanUndo;
end;

function MaskEdit_GetModified(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure MaskEdit_SetModified(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function MaskEdit_GetSelLength(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.SelLength;
end;

procedure MaskEdit_SetSelLength(AObj: TMaskEdit; AValue: Integer); stdcall;
begin
  AObj.SelLength := AValue;
end;

function MaskEdit_GetSelStart(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.SelStart;
end;

procedure MaskEdit_SetSelStart(AObj: TMaskEdit; AValue: Integer); stdcall;
begin
  AObj.SelStart := AValue;
end;

function MaskEdit_GetSelText(AObj: TMaskEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure MaskEdit_SetSelText(AObj: TMaskEdit; AValue: PChar); stdcall;
begin
  AObj.SelText := AValue;
end;

function MaskEdit_GetDockSite(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure MaskEdit_SetDockSite(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function MaskEdit_GetBrush(AObj: TMaskEdit): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function MaskEdit_GetControlCount(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function MaskEdit_GetHandle(AObj: TMaskEdit): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function MaskEdit_GetParentWindow(AObj: TMaskEdit): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure MaskEdit_SetParentWindow(AObj: TMaskEdit; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function MaskEdit_GetUseDockManager(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure MaskEdit_SetUseDockManager(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function MaskEdit_GetAction(AObj: TMaskEdit): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure MaskEdit_SetAction(AObj: TMaskEdit; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure MaskEdit_GetBoundsRect(AObj: TMaskEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure MaskEdit_SetBoundsRect(AObj: TMaskEdit; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function MaskEdit_GetClientHeight(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure MaskEdit_SetClientHeight(AObj: TMaskEdit; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure MaskEdit_GetClientRect(AObj: TMaskEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function MaskEdit_GetClientWidth(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure MaskEdit_SetClientWidth(AObj: TMaskEdit; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function MaskEdit_GetExplicitLeft(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function MaskEdit_GetExplicitTop(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function MaskEdit_GetExplicitWidth(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function MaskEdit_GetExplicitHeight(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function MaskEdit_GetFloating(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function MaskEdit_GetParent(AObj: TMaskEdit): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure MaskEdit_SetParent(AObj: TMaskEdit; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function MaskEdit_GetAlignWithMargins(AObj: TMaskEdit): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure MaskEdit_SetAlignWithMargins(AObj: TMaskEdit; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function MaskEdit_GetLeft(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure MaskEdit_SetLeft(AObj: TMaskEdit; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function MaskEdit_GetTop(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure MaskEdit_SetTop(AObj: TMaskEdit; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function MaskEdit_GetWidth(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure MaskEdit_SetWidth(AObj: TMaskEdit; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function MaskEdit_GetHeight(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure MaskEdit_SetHeight(AObj: TMaskEdit; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function MaskEdit_GetCursor(AObj: TMaskEdit): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure MaskEdit_SetCursor(AObj: TMaskEdit; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function MaskEdit_GetHint(AObj: TMaskEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure MaskEdit_SetHint(AObj: TMaskEdit; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function MaskEdit_GetMargins(AObj: TMaskEdit): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure MaskEdit_SetMargins(AObj: TMaskEdit; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function MaskEdit_GetCustomHint(AObj: TMaskEdit): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure MaskEdit_SetCustomHint(AObj: TMaskEdit; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function MaskEdit_GetComponentCount(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function MaskEdit_GetComponentIndex(AObj: TMaskEdit): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure MaskEdit_SetComponentIndex(AObj: TMaskEdit; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function MaskEdit_GetOwner(AObj: TMaskEdit): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function MaskEdit_GetName(AObj: TMaskEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure MaskEdit_SetName(AObj: TMaskEdit; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function MaskEdit_GetTag(AObj: TMaskEdit): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure MaskEdit_SetTag(AObj: TMaskEdit; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function MaskEdit_GetControls(AObj: TMaskEdit; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function MaskEdit_GetComponents(AObj: TMaskEdit; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function MaskEdit_StaticClassType: TClass; stdcall;
begin
  Result := TMaskEdit;
end;

exports
  MaskEdit_Create,
  MaskEdit_Free,
  MaskEdit_ValidateEdit,
  MaskEdit_Clear,
  MaskEdit_GetTextLen,
  MaskEdit_ClearSelection,
  MaskEdit_CopyToClipboard,
  MaskEdit_CutToClipboard,
  MaskEdit_PasteFromClipboard,
  MaskEdit_Undo,
  MaskEdit_ClearUndo,
  MaskEdit_SelectAll,
  MaskEdit_GetSelTextBuf,
  MaskEdit_CanFocus,
  MaskEdit_FlipChildren,
  MaskEdit_Focused,
  MaskEdit_HandleAllocated,
  MaskEdit_Invalidate,
  MaskEdit_Realign,
  MaskEdit_Repaint,
  MaskEdit_ScaleBy,
  MaskEdit_SetBounds,
  MaskEdit_SetFocus,
  MaskEdit_Update,
  MaskEdit_BringToFront,
  MaskEdit_ClientToScreen,
  MaskEdit_ClientToParent,
  MaskEdit_Dragging,
  MaskEdit_HasParent,
  MaskEdit_Hide,
  MaskEdit_Perform,
  MaskEdit_Refresh,
  MaskEdit_ScreenToClient,
  MaskEdit_ParentToClient,
  MaskEdit_SendToBack,
  MaskEdit_Show,
  MaskEdit_GetTextBuf,
  MaskEdit_FindComponent,
  MaskEdit_GetNamePath,
  MaskEdit_Assign,
  MaskEdit_DisposeOf,
  MaskEdit_ClassType,
  MaskEdit_ClassName,
  MaskEdit_InstanceSize,
  MaskEdit_InheritsFrom,
  MaskEdit_Equals,
  MaskEdit_GetHashCode,
  MaskEdit_ToString,
  MaskEdit_GetAlign,
  MaskEdit_SetAlign,
  MaskEdit_GetAlignment,
  MaskEdit_SetAlignment,
  MaskEdit_GetAnchors,
  MaskEdit_SetAnchors,
  MaskEdit_GetAutoSelect,
  MaskEdit_SetAutoSelect,
  MaskEdit_GetAutoSize,
  MaskEdit_SetAutoSize,
  MaskEdit_GetBevelEdges,
  MaskEdit_SetBevelEdges,
  MaskEdit_GetBevelInner,
  MaskEdit_SetBevelInner,
  MaskEdit_GetBevelOuter,
  MaskEdit_SetBevelOuter,
  MaskEdit_GetBevelKind,
  MaskEdit_SetBevelKind,
  MaskEdit_GetBiDiMode,
  MaskEdit_SetBiDiMode,
  MaskEdit_GetBorderStyle,
  MaskEdit_SetBorderStyle,
  MaskEdit_GetCharCase,
  MaskEdit_SetCharCase,
  MaskEdit_GetColor,
  MaskEdit_SetColor,
  MaskEdit_GetDoubleBuffered,
  MaskEdit_SetDoubleBuffered,
  MaskEdit_GetDragCursor,
  MaskEdit_SetDragCursor,
  MaskEdit_GetDragKind,
  MaskEdit_SetDragKind,
  MaskEdit_GetDragMode,
  MaskEdit_SetDragMode,
  MaskEdit_GetEnabled,
  MaskEdit_SetEnabled,
  MaskEdit_GetFont,
  MaskEdit_SetFont,
  MaskEdit_GetMaxLength,
  MaskEdit_SetMaxLength,
  MaskEdit_GetParentColor,
  MaskEdit_SetParentColor,
  MaskEdit_GetParentCtl3D,
  MaskEdit_SetParentCtl3D,
  MaskEdit_GetParentDoubleBuffered,
  MaskEdit_SetParentDoubleBuffered,
  MaskEdit_GetParentFont,
  MaskEdit_SetParentFont,
  MaskEdit_GetParentShowHint,
  MaskEdit_SetParentShowHint,
  MaskEdit_GetPasswordChar,
  MaskEdit_SetPasswordChar,
  MaskEdit_GetPopupMenu,
  MaskEdit_SetPopupMenu,
  MaskEdit_GetReadOnly,
  MaskEdit_SetReadOnly,
  MaskEdit_GetShowHint,
  MaskEdit_SetShowHint,
  MaskEdit_GetTabOrder,
  MaskEdit_SetTabOrder,
  MaskEdit_GetTabStop,
  MaskEdit_SetTabStop,
  MaskEdit_GetText,
  MaskEdit_SetText,
  MaskEdit_GetTextHint,
  MaskEdit_SetTextHint,
  MaskEdit_GetVisible,
  MaskEdit_SetVisible,
  MaskEdit_GetStyleElements,
  MaskEdit_SetStyleElements,
  MaskEdit_SetOnChange,
  MaskEdit_SetOnClick,
  MaskEdit_SetOnDblClick,
  MaskEdit_SetOnDragDrop,
  MaskEdit_SetOnDragOver,
  MaskEdit_SetOnEndDock,
  MaskEdit_SetOnEndDrag,
  MaskEdit_SetOnEnter,
  MaskEdit_SetOnExit,
  MaskEdit_SetOnKeyDown,
  MaskEdit_SetOnKeyPress,
  MaskEdit_SetOnKeyUp,
  MaskEdit_SetOnMouseDown,
  MaskEdit_SetOnMouseEnter,
  MaskEdit_SetOnMouseLeave,
  MaskEdit_SetOnMouseMove,
  MaskEdit_SetOnMouseUp,
  MaskEdit_SetOnStartDock,
  MaskEdit_GetIsMasked,
  MaskEdit_GetEditText,
  MaskEdit_SetEditText,
  MaskEdit_GetCanUndo,
  MaskEdit_GetModified,
  MaskEdit_SetModified,
  MaskEdit_GetSelLength,
  MaskEdit_SetSelLength,
  MaskEdit_GetSelStart,
  MaskEdit_SetSelStart,
  MaskEdit_GetSelText,
  MaskEdit_SetSelText,
  MaskEdit_GetDockSite,
  MaskEdit_SetDockSite,
  MaskEdit_GetBrush,
  MaskEdit_GetControlCount,
  MaskEdit_GetHandle,
  MaskEdit_GetParentWindow,
  MaskEdit_SetParentWindow,
  MaskEdit_GetUseDockManager,
  MaskEdit_SetUseDockManager,
  MaskEdit_GetAction,
  MaskEdit_SetAction,
  MaskEdit_GetBoundsRect,
  MaskEdit_SetBoundsRect,
  MaskEdit_GetClientHeight,
  MaskEdit_SetClientHeight,
  MaskEdit_GetClientRect,
  MaskEdit_GetClientWidth,
  MaskEdit_SetClientWidth,
  MaskEdit_GetExplicitLeft,
  MaskEdit_GetExplicitTop,
  MaskEdit_GetExplicitWidth,
  MaskEdit_GetExplicitHeight,
  MaskEdit_GetFloating,
  MaskEdit_GetParent,
  MaskEdit_SetParent,
  MaskEdit_GetAlignWithMargins,
  MaskEdit_SetAlignWithMargins,
  MaskEdit_GetLeft,
  MaskEdit_SetLeft,
  MaskEdit_GetTop,
  MaskEdit_SetTop,
  MaskEdit_GetWidth,
  MaskEdit_SetWidth,
  MaskEdit_GetHeight,
  MaskEdit_SetHeight,
  MaskEdit_GetCursor,
  MaskEdit_SetCursor,
  MaskEdit_GetHint,
  MaskEdit_SetHint,
  MaskEdit_GetMargins,
  MaskEdit_SetMargins,
  MaskEdit_GetCustomHint,
  MaskEdit_SetCustomHint,
  MaskEdit_GetComponentCount,
  MaskEdit_GetComponentIndex,
  MaskEdit_SetComponentIndex,
  MaskEdit_GetOwner,
  MaskEdit_GetName,
  MaskEdit_SetName,
  MaskEdit_GetTag,
  MaskEdit_SetTag,
  MaskEdit_GetControls,
  MaskEdit_GetComponents,
  MaskEdit_StaticClassType;

