
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Edit_Create(AOwner: TComponent): TEdit; stdcall;
begin
  Result :=  TEdit.Create(AOwner);
end;

procedure Edit_Free(AObj: TEdit); stdcall;
begin
  AObj.Free;
end;

procedure Edit_Clear(AObj: TEdit); stdcall;
begin
  AObj.Clear;
end;

procedure Edit_ClearSelection(AObj: TEdit); stdcall;
begin
  AObj.ClearSelection;
end;

procedure Edit_CopyToClipboard(AObj: TEdit); stdcall;
begin
  AObj.CopyToClipboard;
end;

procedure Edit_CutToClipboard(AObj: TEdit); stdcall;
begin
  AObj.CutToClipboard;
end;

procedure Edit_PasteFromClipboard(AObj: TEdit); stdcall;
begin
  AObj.PasteFromClipboard;
end;

procedure Edit_Undo(AObj: TEdit); stdcall;
begin
  AObj.Undo;
end;

procedure Edit_ClearUndo(AObj: TEdit); stdcall;
begin
  AObj.ClearUndo;
end;

procedure Edit_SelectAll(AObj: TEdit); stdcall;
begin
  AObj.SelectAll;
end;

function Edit_GetSelTextBuf(AObj: TEdit; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetSelTextBuf(Buffer, BufSize);
end;

function Edit_CanFocus(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure Edit_FlipChildren(AObj: TEdit; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function Edit_Focused(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function Edit_HandleAllocated(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure Edit_Invalidate(AObj: TEdit); stdcall;
begin
  AObj.Invalidate;
end;

procedure Edit_Realign(AObj: TEdit); stdcall;
begin
  AObj.Realign;
end;

procedure Edit_Repaint(AObj: TEdit); stdcall;
begin
  AObj.Repaint;
end;

procedure Edit_ScaleBy(AObj: TEdit; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure Edit_SetBounds(AObj: TEdit; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Edit_SetFocus(AObj: TEdit); stdcall;
begin
  AObj.SetFocus;
end;

procedure Edit_Update(AObj: TEdit); stdcall;
begin
  AObj.Update;
end;

procedure Edit_BringToFront(AObj: TEdit); stdcall;
begin
  AObj.BringToFront;
end;

procedure Edit_ClientToScreen(AObj: TEdit; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure Edit_ClientToParent(AObj: TEdit; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function Edit_Dragging(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function Edit_HasParent(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Edit_Hide(AObj: TEdit); stdcall;
begin
  AObj.Hide;
end;

function Edit_Perform(AObj: TEdit; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Edit_Refresh(AObj: TEdit); stdcall;
begin
  AObj.Refresh;
end;

procedure Edit_ScreenToClient(AObj: TEdit; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure Edit_ParentToClient(AObj: TEdit; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure Edit_SendToBack(AObj: TEdit); stdcall;
begin
  AObj.SendToBack;
end;

procedure Edit_Show(AObj: TEdit); stdcall;
begin
  AObj.Show;
end;

function Edit_GetTextBuf(AObj: TEdit; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Edit_GetTextLen(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function Edit_FindComponent(AObj: TEdit; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Edit_GetNamePath(AObj: TEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Edit_Assign(AObj: TEdit; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure Edit_DisposeOf(AObj: TEdit); stdcall;
begin
  AObj.DisposeOf;
end;

function Edit_ClassType(AObj: TEdit): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function Edit_ClassName(AObj: TEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Edit_InstanceSize(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function Edit_InheritsFrom(AObj: TEdit; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function Edit_Equals(AObj: TEdit; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Edit_GetHashCode(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Edit_ToString(AObj: TEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Edit_GetAlign(AObj: TEdit): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Edit_SetAlign(AObj: TEdit; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Edit_GetAlignment(AObj: TEdit): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure Edit_SetAlignment(AObj: TEdit; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function Edit_GetAnchors(AObj: TEdit): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Edit_SetAnchors(AObj: TEdit; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Edit_GetAutoSelect(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.AutoSelect;
end;

procedure Edit_SetAutoSelect(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.AutoSelect := AValue;
end;

function Edit_GetAutoSize(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure Edit_SetAutoSize(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function Edit_GetBevelEdges(AObj: TEdit): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure Edit_SetBevelEdges(AObj: TEdit; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function Edit_GetBevelInner(AObj: TEdit): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure Edit_SetBevelInner(AObj: TEdit; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function Edit_GetBevelKind(AObj: TEdit): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure Edit_SetBevelKind(AObj: TEdit; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function Edit_GetBevelOuter(AObj: TEdit): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure Edit_SetBevelOuter(AObj: TEdit; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function Edit_GetBiDiMode(AObj: TEdit): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Edit_SetBiDiMode(AObj: TEdit; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function Edit_GetBorderStyle(AObj: TEdit): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure Edit_SetBorderStyle(AObj: TEdit; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function Edit_GetColor(AObj: TEdit): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure Edit_SetColor(AObj: TEdit; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function Edit_GetDoubleBuffered(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure Edit_SetDoubleBuffered(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function Edit_GetDragCursor(AObj: TEdit): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure Edit_SetDragCursor(AObj: TEdit; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function Edit_GetDragKind(AObj: TEdit): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure Edit_SetDragKind(AObj: TEdit; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function Edit_GetDragMode(AObj: TEdit): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure Edit_SetDragMode(AObj: TEdit; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function Edit_GetEnabled(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Edit_SetEnabled(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Edit_GetFont(AObj: TEdit): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure Edit_SetFont(AObj: TEdit; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function Edit_GetHideSelection(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.HideSelection;
end;

procedure Edit_SetHideSelection(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.HideSelection := AValue;
end;

function Edit_GetMaxLength(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.MaxLength;
end;

procedure Edit_SetMaxLength(AObj: TEdit; AValue: Integer); stdcall;
begin
  AObj.MaxLength := AValue;
end;

function Edit_GetNumbersOnly(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.NumbersOnly;
end;

procedure Edit_SetNumbersOnly(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.NumbersOnly := AValue;
end;

function Edit_GetParentColor(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure Edit_SetParentColor(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function Edit_GetParentCtl3D(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure Edit_SetParentCtl3D(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function Edit_GetParentDoubleBuffered(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure Edit_SetParentDoubleBuffered(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function Edit_GetParentFont(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure Edit_SetParentFont(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function Edit_GetParentShowHint(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Edit_SetParentShowHint(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function Edit_GetPasswordChar(AObj: TEdit): Char; stdcall;
begin
  Result :=  AObj.PasswordChar;
end;

procedure Edit_SetPasswordChar(AObj: TEdit; AValue: Char); stdcall;
begin
  AObj.PasswordChar := AValue;
end;

function Edit_GetPopupMenu(AObj: TEdit): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Edit_SetPopupMenu(AObj: TEdit; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function Edit_GetReadOnly(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.ReadOnly;
end;

procedure Edit_SetReadOnly(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.ReadOnly := AValue;
end;

function Edit_GetShowHint(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Edit_SetShowHint(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Edit_GetTabOrder(AObj: TEdit): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure Edit_SetTabOrder(AObj: TEdit; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function Edit_GetTabStop(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure Edit_SetTabStop(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function Edit_GetText(AObj: TEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure Edit_SetText(AObj: TEdit; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function Edit_GetTextHint(AObj: TEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.TextHint);
end;

procedure Edit_SetTextHint(AObj: TEdit; AValue: PChar); stdcall;
begin
  AObj.TextHint := AValue;
end;

function Edit_GetVisible(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Edit_SetVisible(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Edit_GetStyleElements(AObj: TEdit): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Edit_SetStyleElements(AObj: TEdit; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure Edit_SetOnChange(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnClick(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnContextPopup(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnDblClick(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnDragDrop(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnDragOver(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnEndDock(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnEndDrag(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnEnter(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnExit(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnKeyDown(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnKeyPress(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnKeyUp(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnMouseDown(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnMouseEnter(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnMouseLeave(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnMouseMove(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnMouseUp(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

procedure Edit_SetOnStartDock(AObj: TEdit; AEventId: NativeUInt); stdcall;
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

function Edit_GetCanUndo(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.CanUndo;
end;

function Edit_GetModified(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure Edit_SetModified(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function Edit_GetSelLength(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.SelLength;
end;

procedure Edit_SetSelLength(AObj: TEdit; AValue: Integer); stdcall;
begin
  AObj.SelLength := AValue;
end;

function Edit_GetSelStart(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.SelStart;
end;

procedure Edit_SetSelStart(AObj: TEdit; AValue: Integer); stdcall;
begin
  AObj.SelStart := AValue;
end;

function Edit_GetSelText(AObj: TEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure Edit_SetSelText(AObj: TEdit; AValue: PChar); stdcall;
begin
  AObj.SelText := AValue;
end;

function Edit_GetDockSite(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure Edit_SetDockSite(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function Edit_GetBrush(AObj: TEdit): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function Edit_GetControlCount(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function Edit_GetHandle(AObj: TEdit): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function Edit_GetParentWindow(AObj: TEdit): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure Edit_SetParentWindow(AObj: TEdit; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function Edit_GetUseDockManager(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure Edit_SetUseDockManager(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function Edit_GetAction(AObj: TEdit): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Edit_SetAction(AObj: TEdit; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure Edit_GetBoundsRect(AObj: TEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Edit_SetBoundsRect(AObj: TEdit; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Edit_GetClientHeight(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Edit_SetClientHeight(AObj: TEdit; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Edit_GetClientRect(AObj: TEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Edit_GetClientWidth(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Edit_SetClientWidth(AObj: TEdit; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Edit_GetExplicitLeft(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Edit_GetExplicitTop(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Edit_GetExplicitWidth(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Edit_GetExplicitHeight(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Edit_GetFloating(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function Edit_GetParent(AObj: TEdit): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Edit_SetParent(AObj: TEdit; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Edit_GetAlignWithMargins(AObj: TEdit): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Edit_SetAlignWithMargins(AObj: TEdit; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Edit_GetLeft(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Edit_SetLeft(AObj: TEdit; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Edit_GetTop(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Edit_SetTop(AObj: TEdit; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Edit_GetWidth(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Edit_SetWidth(AObj: TEdit; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Edit_GetHeight(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Edit_SetHeight(AObj: TEdit; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Edit_GetCursor(AObj: TEdit): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Edit_SetCursor(AObj: TEdit; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Edit_GetHint(AObj: TEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Edit_SetHint(AObj: TEdit; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Edit_GetMargins(AObj: TEdit): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Edit_SetMargins(AObj: TEdit; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Edit_GetCustomHint(AObj: TEdit): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Edit_SetCustomHint(AObj: TEdit; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Edit_GetComponentCount(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Edit_GetComponentIndex(AObj: TEdit): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Edit_SetComponentIndex(AObj: TEdit; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Edit_GetOwner(AObj: TEdit): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Edit_GetName(AObj: TEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Edit_SetName(AObj: TEdit; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Edit_GetTag(AObj: TEdit): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Edit_SetTag(AObj: TEdit; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Edit_GetControls(AObj: TEdit; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function Edit_GetComponents(AObj: TEdit; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function Edit_StaticClassType: TClass; stdcall;
begin
  Result := TEdit;
end;

exports
  Edit_Create,
  Edit_Free,
  Edit_Clear,
  Edit_ClearSelection,
  Edit_CopyToClipboard,
  Edit_CutToClipboard,
  Edit_PasteFromClipboard,
  Edit_Undo,
  Edit_ClearUndo,
  Edit_SelectAll,
  Edit_GetSelTextBuf,
  Edit_CanFocus,
  Edit_FlipChildren,
  Edit_Focused,
  Edit_HandleAllocated,
  Edit_Invalidate,
  Edit_Realign,
  Edit_Repaint,
  Edit_ScaleBy,
  Edit_SetBounds,
  Edit_SetFocus,
  Edit_Update,
  Edit_BringToFront,
  Edit_ClientToScreen,
  Edit_ClientToParent,
  Edit_Dragging,
  Edit_HasParent,
  Edit_Hide,
  Edit_Perform,
  Edit_Refresh,
  Edit_ScreenToClient,
  Edit_ParentToClient,
  Edit_SendToBack,
  Edit_Show,
  Edit_GetTextBuf,
  Edit_GetTextLen,
  Edit_FindComponent,
  Edit_GetNamePath,
  Edit_Assign,
  Edit_DisposeOf,
  Edit_ClassType,
  Edit_ClassName,
  Edit_InstanceSize,
  Edit_InheritsFrom,
  Edit_Equals,
  Edit_GetHashCode,
  Edit_ToString,
  Edit_GetAlign,
  Edit_SetAlign,
  Edit_GetAlignment,
  Edit_SetAlignment,
  Edit_GetAnchors,
  Edit_SetAnchors,
  Edit_GetAutoSelect,
  Edit_SetAutoSelect,
  Edit_GetAutoSize,
  Edit_SetAutoSize,
  Edit_GetBevelEdges,
  Edit_SetBevelEdges,
  Edit_GetBevelInner,
  Edit_SetBevelInner,
  Edit_GetBevelKind,
  Edit_SetBevelKind,
  Edit_GetBevelOuter,
  Edit_SetBevelOuter,
  Edit_GetBiDiMode,
  Edit_SetBiDiMode,
  Edit_GetBorderStyle,
  Edit_SetBorderStyle,
  Edit_GetColor,
  Edit_SetColor,
  Edit_GetDoubleBuffered,
  Edit_SetDoubleBuffered,
  Edit_GetDragCursor,
  Edit_SetDragCursor,
  Edit_GetDragKind,
  Edit_SetDragKind,
  Edit_GetDragMode,
  Edit_SetDragMode,
  Edit_GetEnabled,
  Edit_SetEnabled,
  Edit_GetFont,
  Edit_SetFont,
  Edit_GetHideSelection,
  Edit_SetHideSelection,
  Edit_GetMaxLength,
  Edit_SetMaxLength,
  Edit_GetNumbersOnly,
  Edit_SetNumbersOnly,
  Edit_GetParentColor,
  Edit_SetParentColor,
  Edit_GetParentCtl3D,
  Edit_SetParentCtl3D,
  Edit_GetParentDoubleBuffered,
  Edit_SetParentDoubleBuffered,
  Edit_GetParentFont,
  Edit_SetParentFont,
  Edit_GetParentShowHint,
  Edit_SetParentShowHint,
  Edit_GetPasswordChar,
  Edit_SetPasswordChar,
  Edit_GetPopupMenu,
  Edit_SetPopupMenu,
  Edit_GetReadOnly,
  Edit_SetReadOnly,
  Edit_GetShowHint,
  Edit_SetShowHint,
  Edit_GetTabOrder,
  Edit_SetTabOrder,
  Edit_GetTabStop,
  Edit_SetTabStop,
  Edit_GetText,
  Edit_SetText,
  Edit_GetTextHint,
  Edit_SetTextHint,
  Edit_GetVisible,
  Edit_SetVisible,
  Edit_GetStyleElements,
  Edit_SetStyleElements,
  Edit_SetOnChange,
  Edit_SetOnClick,
  Edit_SetOnContextPopup,
  Edit_SetOnDblClick,
  Edit_SetOnDragDrop,
  Edit_SetOnDragOver,
  Edit_SetOnEndDock,
  Edit_SetOnEndDrag,
  Edit_SetOnEnter,
  Edit_SetOnExit,
  Edit_SetOnKeyDown,
  Edit_SetOnKeyPress,
  Edit_SetOnKeyUp,
  Edit_SetOnMouseDown,
  Edit_SetOnMouseEnter,
  Edit_SetOnMouseLeave,
  Edit_SetOnMouseMove,
  Edit_SetOnMouseUp,
  Edit_SetOnStartDock,
  Edit_GetCanUndo,
  Edit_GetModified,
  Edit_SetModified,
  Edit_GetSelLength,
  Edit_SetSelLength,
  Edit_GetSelStart,
  Edit_SetSelStart,
  Edit_GetSelText,
  Edit_SetSelText,
  Edit_GetDockSite,
  Edit_SetDockSite,
  Edit_GetBrush,
  Edit_GetControlCount,
  Edit_GetHandle,
  Edit_GetParentWindow,
  Edit_SetParentWindow,
  Edit_GetUseDockManager,
  Edit_SetUseDockManager,
  Edit_GetAction,
  Edit_SetAction,
  Edit_GetBoundsRect,
  Edit_SetBoundsRect,
  Edit_GetClientHeight,
  Edit_SetClientHeight,
  Edit_GetClientRect,
  Edit_GetClientWidth,
  Edit_SetClientWidth,
  Edit_GetExplicitLeft,
  Edit_GetExplicitTop,
  Edit_GetExplicitWidth,
  Edit_GetExplicitHeight,
  Edit_GetFloating,
  Edit_GetParent,
  Edit_SetParent,
  Edit_GetAlignWithMargins,
  Edit_SetAlignWithMargins,
  Edit_GetLeft,
  Edit_SetLeft,
  Edit_GetTop,
  Edit_SetTop,
  Edit_GetWidth,
  Edit_SetWidth,
  Edit_GetHeight,
  Edit_SetHeight,
  Edit_GetCursor,
  Edit_SetCursor,
  Edit_GetHint,
  Edit_SetHint,
  Edit_GetMargins,
  Edit_SetMargins,
  Edit_GetCustomHint,
  Edit_SetCustomHint,
  Edit_GetComponentCount,
  Edit_GetComponentIndex,
  Edit_SetComponentIndex,
  Edit_GetOwner,
  Edit_GetName,
  Edit_SetName,
  Edit_GetTag,
  Edit_SetTag,
  Edit_GetControls,
  Edit_GetComponents,
  Edit_StaticClassType;

