
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function SpinEdit_Create(AOwner: TComponent): TSpinEdit; stdcall;
begin
  Result :=  TSpinEdit.Create(AOwner);
end;

procedure SpinEdit_Free(AObj: TSpinEdit); stdcall;
begin
  AObj.Free;
end;

procedure SpinEdit_Clear(AObj: TSpinEdit); stdcall;
begin
  AObj.Clear;
end;

procedure SpinEdit_ClearSelection(AObj: TSpinEdit); stdcall;
begin
  AObj.ClearSelection;
end;

procedure SpinEdit_CopyToClipboard(AObj: TSpinEdit); stdcall;
begin
  AObj.CopyToClipboard;
end;

procedure SpinEdit_CutToClipboard(AObj: TSpinEdit); stdcall;
begin
  AObj.CutToClipboard;
end;

procedure SpinEdit_PasteFromClipboard(AObj: TSpinEdit); stdcall;
begin
  AObj.PasteFromClipboard;
end;

procedure SpinEdit_Undo(AObj: TSpinEdit); stdcall;
begin
  AObj.Undo;
end;

procedure SpinEdit_ClearUndo(AObj: TSpinEdit); stdcall;
begin
  AObj.ClearUndo;
end;

procedure SpinEdit_SelectAll(AObj: TSpinEdit); stdcall;
begin
  AObj.SelectAll;
end;

function SpinEdit_GetSelTextBuf(AObj: TSpinEdit; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetSelTextBuf(Buffer, BufSize);
end;

function SpinEdit_CanFocus(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function SpinEdit_ContainsControl(AObj: TSpinEdit; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function SpinEdit_ControlAtPos(AObj: TSpinEdit; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure SpinEdit_DisableAlign(AObj: TSpinEdit); stdcall;
begin
  AObj.DisableAlign;
end;

procedure SpinEdit_EnableAlign(AObj: TSpinEdit); stdcall;
begin
  AObj.EnableAlign;
end;

function SpinEdit_FindChildControl(AObj: TSpinEdit; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure SpinEdit_FlipChildren(AObj: TSpinEdit; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function SpinEdit_Focused(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function SpinEdit_HandleAllocated(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure SpinEdit_InsertControl(AObj: TSpinEdit; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure SpinEdit_Invalidate(AObj: TSpinEdit); stdcall;
begin
  AObj.Invalidate;
end;

procedure SpinEdit_PaintTo(AObj: TSpinEdit; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure SpinEdit_RemoveControl(AObj: TSpinEdit; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure SpinEdit_Realign(AObj: TSpinEdit); stdcall;
begin
  AObj.Realign;
end;

procedure SpinEdit_Repaint(AObj: TSpinEdit); stdcall;
begin
  AObj.Repaint;
end;

procedure SpinEdit_ScaleBy(AObj: TSpinEdit; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure SpinEdit_ScrollBy(AObj: TSpinEdit; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure SpinEdit_SetBounds(AObj: TSpinEdit; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure SpinEdit_SetFocus(AObj: TSpinEdit); stdcall;
begin
  AObj.SetFocus;
end;

procedure SpinEdit_Update(AObj: TSpinEdit); stdcall;
begin
  AObj.Update;
end;

procedure SpinEdit_UpdateControlState(AObj: TSpinEdit); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure SpinEdit_BringToFront(AObj: TSpinEdit); stdcall;
begin
  AObj.BringToFront;
end;

procedure SpinEdit_ClientToScreen(AObj: TSpinEdit; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure SpinEdit_ClientToParent(AObj: TSpinEdit; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function SpinEdit_Dragging(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function SpinEdit_HasParent(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure SpinEdit_Hide(AObj: TSpinEdit); stdcall;
begin
  AObj.Hide;
end;

function SpinEdit_Perform(AObj: TSpinEdit; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure SpinEdit_Refresh(AObj: TSpinEdit); stdcall;
begin
  AObj.Refresh;
end;

procedure SpinEdit_ScreenToClient(AObj: TSpinEdit; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure SpinEdit_ParentToClient(AObj: TSpinEdit; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure SpinEdit_SendToBack(AObj: TSpinEdit); stdcall;
begin
  AObj.SendToBack;
end;

procedure SpinEdit_Show(AObj: TSpinEdit); stdcall;
begin
  AObj.Show;
end;

function SpinEdit_GetTextBuf(AObj: TSpinEdit; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function SpinEdit_GetTextLen(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure SpinEdit_SetTextBuf(AObj: TSpinEdit; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function SpinEdit_FindComponent(AObj: TSpinEdit; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function SpinEdit_GetNamePath(AObj: TSpinEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure SpinEdit_Assign(AObj: TSpinEdit; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure SpinEdit_DisposeOf(AObj: TSpinEdit); stdcall;
begin
  AObj.DisposeOf;
end;

function SpinEdit_ClassType(AObj: TSpinEdit): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function SpinEdit_ClassName(AObj: TSpinEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function SpinEdit_InstanceSize(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function SpinEdit_InheritsFrom(AObj: TSpinEdit; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function SpinEdit_Equals(AObj: TSpinEdit; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function SpinEdit_GetHashCode(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function SpinEdit_ToString(AObj: TSpinEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function SpinEdit_GetAnchors(AObj: TSpinEdit): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure SpinEdit_SetAnchors(AObj: TSpinEdit; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function SpinEdit_GetAutoSelect(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.AutoSelect;
end;

procedure SpinEdit_SetAutoSelect(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.AutoSelect := AValue;
end;

function SpinEdit_GetAutoSize(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure SpinEdit_SetAutoSize(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function SpinEdit_GetColor(AObj: TSpinEdit): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure SpinEdit_SetColor(AObj: TSpinEdit; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function SpinEdit_GetDragCursor(AObj: TSpinEdit): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure SpinEdit_SetDragCursor(AObj: TSpinEdit; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function SpinEdit_GetDragMode(AObj: TSpinEdit): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure SpinEdit_SetDragMode(AObj: TSpinEdit; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function SpinEdit_GetEditorEnabled(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.EditorEnabled;
end;

procedure SpinEdit_SetEditorEnabled(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.EditorEnabled := AValue;
end;

function SpinEdit_GetEnabled(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure SpinEdit_SetEnabled(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function SpinEdit_GetFont(AObj: TSpinEdit): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure SpinEdit_SetFont(AObj: TSpinEdit; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function SpinEdit_GetIncrement(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.Increment;
end;

procedure SpinEdit_SetIncrement(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.Increment := AValue;
end;

function SpinEdit_GetMaxLength(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.MaxLength;
end;

procedure SpinEdit_SetMaxLength(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.MaxLength := AValue;
end;

function SpinEdit_GetMaxValue(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.MaxValue;
end;

procedure SpinEdit_SetMaxValue(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.MaxValue := AValue;
end;

function SpinEdit_GetMinValue(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.MinValue;
end;

procedure SpinEdit_SetMinValue(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.MinValue := AValue;
end;

function SpinEdit_GetParentColor(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure SpinEdit_SetParentColor(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function SpinEdit_GetParentCtl3D(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure SpinEdit_SetParentCtl3D(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function SpinEdit_GetParentFont(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure SpinEdit_SetParentFont(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function SpinEdit_GetParentShowHint(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure SpinEdit_SetParentShowHint(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function SpinEdit_GetPopupMenu(AObj: TSpinEdit): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure SpinEdit_SetPopupMenu(AObj: TSpinEdit; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function SpinEdit_GetReadOnly(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.ReadOnly;
end;

procedure SpinEdit_SetReadOnly(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.ReadOnly := AValue;
end;

function SpinEdit_GetShowHint(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure SpinEdit_SetShowHint(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function SpinEdit_GetTabOrder(AObj: TSpinEdit): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure SpinEdit_SetTabOrder(AObj: TSpinEdit; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function SpinEdit_GetTabStop(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure SpinEdit_SetTabStop(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function SpinEdit_GetValue(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.Value;
end;

procedure SpinEdit_SetValue(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.Value := AValue;
end;

function SpinEdit_GetVisible(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure SpinEdit_SetVisible(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

procedure SpinEdit_SetOnChange(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, @TEventClass.OnChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, @TEventClass.OnChange, AEventId);
end;

procedure SpinEdit_SetOnClick(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    TEventClass.Remove(AObj, @TEventClass.OnClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, @TEventClass.OnClick, AEventId);
end;

procedure SpinEdit_SetOnDblClick(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDblClick := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDblClick);
    Exit;
  end;
  AObj.OnDblClick := TEventClass.OnDblClick;
  TEventClass.Add(AObj, @TEventClass.OnDblClick, AEventId);
end;

procedure SpinEdit_SetOnDragDrop(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragDrop := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDragDrop);
    Exit;
  end;
  AObj.OnDragDrop := TEventClass.OnDragDrop;
  TEventClass.Add(AObj, @TEventClass.OnDragDrop, AEventId);
end;

procedure SpinEdit_SetOnDragOver(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragOver := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDragOver);
    Exit;
  end;
  AObj.OnDragOver := TEventClass.OnDragOver;
  TEventClass.Add(AObj, @TEventClass.OnDragOver, AEventId);
end;

procedure SpinEdit_SetOnEndDrag(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDrag := nil;
    TEventClass.Remove(AObj, @TEventClass.OnEndDrag);
    Exit;
  end;
  AObj.OnEndDrag := TEventClass.OnEndDrag;
  TEventClass.Add(AObj, @TEventClass.OnEndDrag, AEventId);
end;

procedure SpinEdit_SetOnEnter(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEnter := nil;
    TEventClass.Remove(AObj, @TEventClass.OnEnter);
    Exit;
  end;
  AObj.OnEnter := TEventClass.OnEnter;
  TEventClass.Add(AObj, @TEventClass.OnEnter, AEventId);
end;

procedure SpinEdit_SetOnExit(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnExit := nil;
    TEventClass.Remove(AObj, @TEventClass.OnExit);
    Exit;
  end;
  AObj.OnExit := TEventClass.OnExit;
  TEventClass.Add(AObj, @TEventClass.OnExit, AEventId);
end;

procedure SpinEdit_SetOnKeyDown(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyDown := nil;
    TEventClass.Remove(AObj, @TEventClass.OnKeyDown);
    Exit;
  end;
  AObj.OnKeyDown := TEventClass.OnKeyDown;
  TEventClass.Add(AObj, @TEventClass.OnKeyDown, AEventId);
end;

procedure SpinEdit_SetOnKeyPress(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyPress := nil;
    TEventClass.Remove(AObj, @TEventClass.OnKeyPress);
    Exit;
  end;
  AObj.OnKeyPress := TEventClass.OnKeyPress;
  TEventClass.Add(AObj, @TEventClass.OnKeyPress, AEventId);
end;

procedure SpinEdit_SetOnKeyUp(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyUp := nil;
    TEventClass.Remove(AObj, @TEventClass.OnKeyUp);
    Exit;
  end;
  AObj.OnKeyUp := TEventClass.OnKeyUp;
  TEventClass.Add(AObj, @TEventClass.OnKeyUp, AEventId);
end;

procedure SpinEdit_SetOnMouseDown(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseDown := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseDown);
    Exit;
  end;
  AObj.OnMouseDown := TEventClass.OnMouseDown;
  TEventClass.Add(AObj, @TEventClass.OnMouseDown, AEventId);
end;

procedure SpinEdit_SetOnMouseMove(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseMove := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseMove);
    Exit;
  end;
  AObj.OnMouseMove := TEventClass.OnMouseMove;
  TEventClass.Add(AObj, @TEventClass.OnMouseMove, AEventId);
end;

procedure SpinEdit_SetOnMouseUp(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseUp := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseUp);
    Exit;
  end;
  AObj.OnMouseUp := TEventClass.OnMouseUp;
  TEventClass.Add(AObj, @TEventClass.OnMouseUp, AEventId);
end;

function SpinEdit_GetAlignment(AObj: TSpinEdit): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure SpinEdit_SetAlignment(AObj: TSpinEdit; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function SpinEdit_GetCanUndo(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.CanUndo;
end;

function SpinEdit_GetModified(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.Modified;
end;

procedure SpinEdit_SetModified(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.Modified := AValue;
end;

function SpinEdit_GetSelLength(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.SelLength;
end;

procedure SpinEdit_SetSelLength(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.SelLength := AValue;
end;

function SpinEdit_GetSelStart(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.SelStart;
end;

procedure SpinEdit_SetSelStart(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.SelStart := AValue;
end;

function SpinEdit_GetSelText(AObj: TSpinEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure SpinEdit_SetSelText(AObj: TSpinEdit; AValue: PChar); stdcall;
begin
  AObj.SelText := AValue;
end;

function SpinEdit_GetText(AObj: TSpinEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure SpinEdit_SetText(AObj: TSpinEdit; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function SpinEdit_GetTextHint(AObj: TSpinEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.TextHint);
end;

procedure SpinEdit_SetTextHint(AObj: TSpinEdit; AValue: PChar); stdcall;
begin
  AObj.TextHint := AValue;
end;

function SpinEdit_GetDockClientCount(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function SpinEdit_GetDockSite(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure SpinEdit_SetDockSite(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function SpinEdit_GetDoubleBuffered(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure SpinEdit_SetDoubleBuffered(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function SpinEdit_GetAlignDisabled(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function SpinEdit_GetMouseInClient(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function SpinEdit_GetVisibleDockClientCount(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function SpinEdit_GetBrush(AObj: TSpinEdit): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function SpinEdit_GetControlCount(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function SpinEdit_GetHandle(AObj: TSpinEdit): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function SpinEdit_GetParentDoubleBuffered(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure SpinEdit_SetParentDoubleBuffered(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function SpinEdit_GetParentWindow(AObj: TSpinEdit): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure SpinEdit_SetParentWindow(AObj: TSpinEdit; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function SpinEdit_GetUseDockManager(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure SpinEdit_SetUseDockManager(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function SpinEdit_GetAction(AObj: TSpinEdit): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure SpinEdit_SetAction(AObj: TSpinEdit; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function SpinEdit_GetAlign(AObj: TSpinEdit): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure SpinEdit_SetAlign(AObj: TSpinEdit; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function SpinEdit_GetBiDiMode(AObj: TSpinEdit): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure SpinEdit_SetBiDiMode(AObj: TSpinEdit; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure SpinEdit_GetBoundsRect(AObj: TSpinEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure SpinEdit_SetBoundsRect(AObj: TSpinEdit; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function SpinEdit_GetClientHeight(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure SpinEdit_SetClientHeight(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure SpinEdit_GetClientOrigin(AObj: TSpinEdit; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure SpinEdit_GetClientRect(AObj: TSpinEdit; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function SpinEdit_GetClientWidth(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure SpinEdit_SetClientWidth(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function SpinEdit_GetControlState(AObj: TSpinEdit): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure SpinEdit_SetControlState(AObj: TSpinEdit; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function SpinEdit_GetControlStyle(AObj: TSpinEdit): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure SpinEdit_SetControlStyle(AObj: TSpinEdit; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function SpinEdit_GetExplicitLeft(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function SpinEdit_GetExplicitTop(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function SpinEdit_GetExplicitWidth(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function SpinEdit_GetExplicitHeight(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function SpinEdit_GetFloating(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function SpinEdit_GetParent(AObj: TSpinEdit): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure SpinEdit_SetParent(AObj: TSpinEdit; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function SpinEdit_GetStyleElements(AObj: TSpinEdit): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure SpinEdit_SetStyleElements(AObj: TSpinEdit; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure SpinEdit_SetOnGesture(AObj: TSpinEdit; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGesture := nil;
    TEventClass.Remove(AObj, @TEventClass.OnGesture);
    Exit;
  end;
  AObj.OnGesture := TEventClass.OnGesture;
  TEventClass.Add(AObj, @TEventClass.OnGesture, AEventId);
end;

function SpinEdit_GetAlignWithMargins(AObj: TSpinEdit): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure SpinEdit_SetAlignWithMargins(AObj: TSpinEdit; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function SpinEdit_GetLeft(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure SpinEdit_SetLeft(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function SpinEdit_GetTop(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure SpinEdit_SetTop(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function SpinEdit_GetWidth(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure SpinEdit_SetWidth(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function SpinEdit_GetHeight(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure SpinEdit_SetHeight(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function SpinEdit_GetCursor(AObj: TSpinEdit): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure SpinEdit_SetCursor(AObj: TSpinEdit; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function SpinEdit_GetHint(AObj: TSpinEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure SpinEdit_SetHint(AObj: TSpinEdit; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function SpinEdit_GetMargins(AObj: TSpinEdit): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure SpinEdit_SetMargins(AObj: TSpinEdit; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function SpinEdit_GetCustomHint(AObj: TSpinEdit): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure SpinEdit_SetCustomHint(AObj: TSpinEdit; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function SpinEdit_GetComponentCount(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function SpinEdit_GetComponentIndex(AObj: TSpinEdit): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure SpinEdit_SetComponentIndex(AObj: TSpinEdit; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function SpinEdit_GetOwner(AObj: TSpinEdit): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function SpinEdit_GetName(AObj: TSpinEdit): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure SpinEdit_SetName(AObj: TSpinEdit; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function SpinEdit_GetTag(AObj: TSpinEdit): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure SpinEdit_SetTag(AObj: TSpinEdit; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function SpinEdit_GetDockClients(AObj: TSpinEdit; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function SpinEdit_GetControls(AObj: TSpinEdit; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function SpinEdit_GetComponents(AObj: TSpinEdit; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function SpinEdit_StaticClassType: TClass; stdcall;
begin
  Result := TSpinEdit;
end;

exports
  SpinEdit_Create,
  SpinEdit_Free,
  SpinEdit_Clear,
  SpinEdit_ClearSelection,
  SpinEdit_CopyToClipboard,
  SpinEdit_CutToClipboard,
  SpinEdit_PasteFromClipboard,
  SpinEdit_Undo,
  SpinEdit_ClearUndo,
  SpinEdit_SelectAll,
  SpinEdit_GetSelTextBuf,
  SpinEdit_CanFocus,
  SpinEdit_ContainsControl,
  SpinEdit_ControlAtPos,
  SpinEdit_DisableAlign,
  SpinEdit_EnableAlign,
  SpinEdit_FindChildControl,
  SpinEdit_FlipChildren,
  SpinEdit_Focused,
  SpinEdit_HandleAllocated,
  SpinEdit_InsertControl,
  SpinEdit_Invalidate,
  SpinEdit_PaintTo,
  SpinEdit_RemoveControl,
  SpinEdit_Realign,
  SpinEdit_Repaint,
  SpinEdit_ScaleBy,
  SpinEdit_ScrollBy,
  SpinEdit_SetBounds,
  SpinEdit_SetFocus,
  SpinEdit_Update,
  SpinEdit_UpdateControlState,
  SpinEdit_BringToFront,
  SpinEdit_ClientToScreen,
  SpinEdit_ClientToParent,
  SpinEdit_Dragging,
  SpinEdit_HasParent,
  SpinEdit_Hide,
  SpinEdit_Perform,
  SpinEdit_Refresh,
  SpinEdit_ScreenToClient,
  SpinEdit_ParentToClient,
  SpinEdit_SendToBack,
  SpinEdit_Show,
  SpinEdit_GetTextBuf,
  SpinEdit_GetTextLen,
  SpinEdit_SetTextBuf,
  SpinEdit_FindComponent,
  SpinEdit_GetNamePath,
  SpinEdit_Assign,
  SpinEdit_DisposeOf,
  SpinEdit_ClassType,
  SpinEdit_ClassName,
  SpinEdit_InstanceSize,
  SpinEdit_InheritsFrom,
  SpinEdit_Equals,
  SpinEdit_GetHashCode,
  SpinEdit_ToString,
  SpinEdit_GetAnchors,
  SpinEdit_SetAnchors,
  SpinEdit_GetAutoSelect,
  SpinEdit_SetAutoSelect,
  SpinEdit_GetAutoSize,
  SpinEdit_SetAutoSize,
  SpinEdit_GetColor,
  SpinEdit_SetColor,
  SpinEdit_GetDragCursor,
  SpinEdit_SetDragCursor,
  SpinEdit_GetDragMode,
  SpinEdit_SetDragMode,
  SpinEdit_GetEditorEnabled,
  SpinEdit_SetEditorEnabled,
  SpinEdit_GetEnabled,
  SpinEdit_SetEnabled,
  SpinEdit_GetFont,
  SpinEdit_SetFont,
  SpinEdit_GetIncrement,
  SpinEdit_SetIncrement,
  SpinEdit_GetMaxLength,
  SpinEdit_SetMaxLength,
  SpinEdit_GetMaxValue,
  SpinEdit_SetMaxValue,
  SpinEdit_GetMinValue,
  SpinEdit_SetMinValue,
  SpinEdit_GetParentColor,
  SpinEdit_SetParentColor,
  SpinEdit_GetParentCtl3D,
  SpinEdit_SetParentCtl3D,
  SpinEdit_GetParentFont,
  SpinEdit_SetParentFont,
  SpinEdit_GetParentShowHint,
  SpinEdit_SetParentShowHint,
  SpinEdit_GetPopupMenu,
  SpinEdit_SetPopupMenu,
  SpinEdit_GetReadOnly,
  SpinEdit_SetReadOnly,
  SpinEdit_GetShowHint,
  SpinEdit_SetShowHint,
  SpinEdit_GetTabOrder,
  SpinEdit_SetTabOrder,
  SpinEdit_GetTabStop,
  SpinEdit_SetTabStop,
  SpinEdit_GetValue,
  SpinEdit_SetValue,
  SpinEdit_GetVisible,
  SpinEdit_SetVisible,
  SpinEdit_SetOnChange,
  SpinEdit_SetOnClick,
  SpinEdit_SetOnDblClick,
  SpinEdit_SetOnDragDrop,
  SpinEdit_SetOnDragOver,
  SpinEdit_SetOnEndDrag,
  SpinEdit_SetOnEnter,
  SpinEdit_SetOnExit,
  SpinEdit_SetOnKeyDown,
  SpinEdit_SetOnKeyPress,
  SpinEdit_SetOnKeyUp,
  SpinEdit_SetOnMouseDown,
  SpinEdit_SetOnMouseMove,
  SpinEdit_SetOnMouseUp,
  SpinEdit_GetAlignment,
  SpinEdit_SetAlignment,
  SpinEdit_GetCanUndo,
  SpinEdit_GetModified,
  SpinEdit_SetModified,
  SpinEdit_GetSelLength,
  SpinEdit_SetSelLength,
  SpinEdit_GetSelStart,
  SpinEdit_SetSelStart,
  SpinEdit_GetSelText,
  SpinEdit_SetSelText,
  SpinEdit_GetText,
  SpinEdit_SetText,
  SpinEdit_GetTextHint,
  SpinEdit_SetTextHint,
  SpinEdit_GetDockClientCount,
  SpinEdit_GetDockSite,
  SpinEdit_SetDockSite,
  SpinEdit_GetDoubleBuffered,
  SpinEdit_SetDoubleBuffered,
  SpinEdit_GetAlignDisabled,
  SpinEdit_GetMouseInClient,
  SpinEdit_GetVisibleDockClientCount,
  SpinEdit_GetBrush,
  SpinEdit_GetControlCount,
  SpinEdit_GetHandle,
  SpinEdit_GetParentDoubleBuffered,
  SpinEdit_SetParentDoubleBuffered,
  SpinEdit_GetParentWindow,
  SpinEdit_SetParentWindow,
  SpinEdit_GetUseDockManager,
  SpinEdit_SetUseDockManager,
  SpinEdit_GetAction,
  SpinEdit_SetAction,
  SpinEdit_GetAlign,
  SpinEdit_SetAlign,
  SpinEdit_GetBiDiMode,
  SpinEdit_SetBiDiMode,
  SpinEdit_GetBoundsRect,
  SpinEdit_SetBoundsRect,
  SpinEdit_GetClientHeight,
  SpinEdit_SetClientHeight,
  SpinEdit_GetClientOrigin,
  SpinEdit_GetClientRect,
  SpinEdit_GetClientWidth,
  SpinEdit_SetClientWidth,
  SpinEdit_GetControlState,
  SpinEdit_SetControlState,
  SpinEdit_GetControlStyle,
  SpinEdit_SetControlStyle,
  SpinEdit_GetExplicitLeft,
  SpinEdit_GetExplicitTop,
  SpinEdit_GetExplicitWidth,
  SpinEdit_GetExplicitHeight,
  SpinEdit_GetFloating,
  SpinEdit_GetParent,
  SpinEdit_SetParent,
  SpinEdit_GetStyleElements,
  SpinEdit_SetStyleElements,
  SpinEdit_SetOnGesture,
  SpinEdit_GetAlignWithMargins,
  SpinEdit_SetAlignWithMargins,
  SpinEdit_GetLeft,
  SpinEdit_SetLeft,
  SpinEdit_GetTop,
  SpinEdit_SetTop,
  SpinEdit_GetWidth,
  SpinEdit_SetWidth,
  SpinEdit_GetHeight,
  SpinEdit_SetHeight,
  SpinEdit_GetCursor,
  SpinEdit_SetCursor,
  SpinEdit_GetHint,
  SpinEdit_SetHint,
  SpinEdit_GetMargins,
  SpinEdit_SetMargins,
  SpinEdit_GetCustomHint,
  SpinEdit_SetCustomHint,
  SpinEdit_GetComponentCount,
  SpinEdit_GetComponentIndex,
  SpinEdit_SetComponentIndex,
  SpinEdit_GetOwner,
  SpinEdit_GetName,
  SpinEdit_SetName,
  SpinEdit_GetTag,
  SpinEdit_SetTag,
  SpinEdit_GetDockClients,
  SpinEdit_GetControls,
  SpinEdit_GetComponents,
  SpinEdit_StaticClassType;

