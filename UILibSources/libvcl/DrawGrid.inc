
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function DrawGrid_Create(AOwner: TComponent): TDrawGrid; stdcall;
begin
  Result :=  TDrawGrid.Create(AOwner);
end;

procedure DrawGrid_Free(AObj: TDrawGrid); stdcall;
begin
  AObj.Free;
end;

procedure DrawGrid_CellRect(AObj: TDrawGrid; ACol: Integer; ARow: Integer; var Result: TRect); stdcall;
begin
  Result :=  AObj.CellRect(ACol, ARow);
end;

procedure DrawGrid_MouseToCell(AObj: TDrawGrid; X: Integer; Y: Integer; var ACol: Integer; var ARow: Integer); stdcall;
begin
  AObj.MouseToCell(X, Y, ACol, ARow);
end;

procedure DrawGrid_MouseCoord(AObj: TDrawGrid; X: Integer; Y: Integer; var Result: TGridCoord); stdcall;
begin
  Result :=  AObj.MouseCoord(X, Y);
end;

function DrawGrid_CanFocus(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function DrawGrid_ContainsControl(AObj: TDrawGrid; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function DrawGrid_ControlAtPos(AObj: TDrawGrid; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure DrawGrid_DisableAlign(AObj: TDrawGrid); stdcall;
begin
  AObj.DisableAlign;
end;

procedure DrawGrid_EnableAlign(AObj: TDrawGrid); stdcall;
begin
  AObj.EnableAlign;
end;

function DrawGrid_FindChildControl(AObj: TDrawGrid; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure DrawGrid_FlipChildren(AObj: TDrawGrid; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function DrawGrid_Focused(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function DrawGrid_HandleAllocated(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure DrawGrid_InsertControl(AObj: TDrawGrid; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure DrawGrid_Invalidate(AObj: TDrawGrid); stdcall;
begin
  AObj.Invalidate;
end;

procedure DrawGrid_PaintTo(AObj: TDrawGrid; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure DrawGrid_RemoveControl(AObj: TDrawGrid; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure DrawGrid_Realign(AObj: TDrawGrid); stdcall;
begin
  AObj.Realign;
end;

procedure DrawGrid_Repaint(AObj: TDrawGrid); stdcall;
begin
  AObj.Repaint;
end;

procedure DrawGrid_ScaleBy(AObj: TDrawGrid; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure DrawGrid_ScrollBy(AObj: TDrawGrid; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure DrawGrid_SetBounds(AObj: TDrawGrid; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure DrawGrid_SetFocus(AObj: TDrawGrid); stdcall;
begin
  AObj.SetFocus;
end;

procedure DrawGrid_Update(AObj: TDrawGrid); stdcall;
begin
  AObj.Update;
end;

procedure DrawGrid_UpdateControlState(AObj: TDrawGrid); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure DrawGrid_BringToFront(AObj: TDrawGrid); stdcall;
begin
  AObj.BringToFront;
end;

procedure DrawGrid_ClientToScreen(AObj: TDrawGrid; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure DrawGrid_ClientToParent(AObj: TDrawGrid; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function DrawGrid_Dragging(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function DrawGrid_HasParent(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure DrawGrid_Hide(AObj: TDrawGrid); stdcall;
begin
  AObj.Hide;
end;

function DrawGrid_Perform(AObj: TDrawGrid; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure DrawGrid_Refresh(AObj: TDrawGrid); stdcall;
begin
  AObj.Refresh;
end;

procedure DrawGrid_ScreenToClient(AObj: TDrawGrid; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure DrawGrid_ParentToClient(AObj: TDrawGrid; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure DrawGrid_SendToBack(AObj: TDrawGrid); stdcall;
begin
  AObj.SendToBack;
end;

procedure DrawGrid_Show(AObj: TDrawGrid); stdcall;
begin
  AObj.Show;
end;

function DrawGrid_GetTextBuf(AObj: TDrawGrid; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function DrawGrid_GetTextLen(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure DrawGrid_SetTextBuf(AObj: TDrawGrid; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function DrawGrid_FindComponent(AObj: TDrawGrid; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function DrawGrid_GetNamePath(AObj: TDrawGrid): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure DrawGrid_Assign(AObj: TDrawGrid; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure DrawGrid_DisposeOf(AObj: TDrawGrid); stdcall;
begin
  AObj.DisposeOf;
end;

function DrawGrid_ClassType(AObj: TDrawGrid): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function DrawGrid_ClassName(AObj: TDrawGrid): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function DrawGrid_InstanceSize(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function DrawGrid_InheritsFrom(AObj: TDrawGrid; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function DrawGrid_Equals(AObj: TDrawGrid; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function DrawGrid_GetHashCode(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function DrawGrid_ToString(AObj: TDrawGrid): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function DrawGrid_GetAlign(AObj: TDrawGrid): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure DrawGrid_SetAlign(AObj: TDrawGrid; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function DrawGrid_GetAnchors(AObj: TDrawGrid): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure DrawGrid_SetAnchors(AObj: TDrawGrid; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function DrawGrid_GetBevelEdges(AObj: TDrawGrid): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure DrawGrid_SetBevelEdges(AObj: TDrawGrid; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function DrawGrid_GetBevelInner(AObj: TDrawGrid): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure DrawGrid_SetBevelInner(AObj: TDrawGrid; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function DrawGrid_GetBevelKind(AObj: TDrawGrid): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure DrawGrid_SetBevelKind(AObj: TDrawGrid; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function DrawGrid_GetBevelOuter(AObj: TDrawGrid): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure DrawGrid_SetBevelOuter(AObj: TDrawGrid; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function DrawGrid_GetBiDiMode(AObj: TDrawGrid): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure DrawGrid_SetBiDiMode(AObj: TDrawGrid; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function DrawGrid_GetBorderStyle(AObj: TDrawGrid): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure DrawGrid_SetBorderStyle(AObj: TDrawGrid; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function DrawGrid_GetColor(AObj: TDrawGrid): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure DrawGrid_SetColor(AObj: TDrawGrid; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function DrawGrid_GetColCount(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.ColCount;
end;

procedure DrawGrid_SetColCount(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.ColCount := AValue;
end;

function DrawGrid_GetDefaultColWidth(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.DefaultColWidth;
end;

procedure DrawGrid_SetDefaultColWidth(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.DefaultColWidth := AValue;
end;

function DrawGrid_GetDefaultRowHeight(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.DefaultRowHeight;
end;

procedure DrawGrid_SetDefaultRowHeight(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.DefaultRowHeight := AValue;
end;

function DrawGrid_GetDefaultDrawing(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.DefaultDrawing;
end;

procedure DrawGrid_SetDefaultDrawing(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.DefaultDrawing := AValue;
end;

function DrawGrid_GetDoubleBuffered(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure DrawGrid_SetDoubleBuffered(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function DrawGrid_GetDragCursor(AObj: TDrawGrid): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure DrawGrid_SetDragCursor(AObj: TDrawGrid; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function DrawGrid_GetDragKind(AObj: TDrawGrid): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure DrawGrid_SetDragKind(AObj: TDrawGrid; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function DrawGrid_GetDragMode(AObj: TDrawGrid): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure DrawGrid_SetDragMode(AObj: TDrawGrid; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function DrawGrid_GetDrawingStyle(AObj: TDrawGrid): TGridDrawingStyle; stdcall;
begin
  Result :=  AObj.DrawingStyle;
end;

procedure DrawGrid_SetDrawingStyle(AObj: TDrawGrid; AValue: TGridDrawingStyle); stdcall;
begin
  AObj.DrawingStyle := AValue;
end;

function DrawGrid_GetEnabled(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure DrawGrid_SetEnabled(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function DrawGrid_GetFixedColor(AObj: TDrawGrid): TColor; stdcall;
begin
  Result :=  AObj.FixedColor;
end;

procedure DrawGrid_SetFixedColor(AObj: TDrawGrid; AValue: TColor); stdcall;
begin
  AObj.FixedColor := AValue;
end;

function DrawGrid_GetFixedCols(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.FixedCols;
end;

procedure DrawGrid_SetFixedCols(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.FixedCols := AValue;
end;

function DrawGrid_GetRowCount(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.RowCount;
end;

procedure DrawGrid_SetRowCount(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.RowCount := AValue;
end;

function DrawGrid_GetFixedRows(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.FixedRows;
end;

procedure DrawGrid_SetFixedRows(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.FixedRows := AValue;
end;

function DrawGrid_GetFont(AObj: TDrawGrid): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure DrawGrid_SetFont(AObj: TDrawGrid; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function DrawGrid_GetGradientEndColor(AObj: TDrawGrid): TColor; stdcall;
begin
  Result :=  AObj.GradientEndColor;
end;

procedure DrawGrid_SetGradientEndColor(AObj: TDrawGrid; AValue: TColor); stdcall;
begin
  AObj.GradientEndColor := AValue;
end;

function DrawGrid_GetGradientStartColor(AObj: TDrawGrid): TColor; stdcall;
begin
  Result :=  AObj.GradientStartColor;
end;

procedure DrawGrid_SetGradientStartColor(AObj: TDrawGrid; AValue: TColor); stdcall;
begin
  AObj.GradientStartColor := AValue;
end;

function DrawGrid_GetGridLineWidth(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.GridLineWidth;
end;

procedure DrawGrid_SetGridLineWidth(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.GridLineWidth := AValue;
end;

function DrawGrid_GetOptions(AObj: TDrawGrid): TGridOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure DrawGrid_SetOptions(AObj: TDrawGrid; AValue: TGridOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function DrawGrid_GetParentColor(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure DrawGrid_SetParentColor(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function DrawGrid_GetParentCtl3D(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure DrawGrid_SetParentCtl3D(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function DrawGrid_GetParentDoubleBuffered(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure DrawGrid_SetParentDoubleBuffered(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function DrawGrid_GetParentFont(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure DrawGrid_SetParentFont(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function DrawGrid_GetParentShowHint(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure DrawGrid_SetParentShowHint(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function DrawGrid_GetPopupMenu(AObj: TDrawGrid): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure DrawGrid_SetPopupMenu(AObj: TDrawGrid; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function DrawGrid_GetScrollBars(AObj: TDrawGrid): TScrollStyle; stdcall;
begin
  Result :=  AObj.ScrollBars;
end;

procedure DrawGrid_SetScrollBars(AObj: TDrawGrid; AValue: TScrollStyle); stdcall;
begin
  AObj.ScrollBars := AValue;
end;

function DrawGrid_GetShowHint(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure DrawGrid_SetShowHint(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function DrawGrid_GetTabOrder(AObj: TDrawGrid): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure DrawGrid_SetTabOrder(AObj: TDrawGrid; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function DrawGrid_GetVisible(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure DrawGrid_SetVisible(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function DrawGrid_GetStyleElements(AObj: TDrawGrid): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure DrawGrid_SetStyleElements(AObj: TDrawGrid; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function DrawGrid_GetVisibleColCount(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.VisibleColCount;
end;

function DrawGrid_GetVisibleRowCount(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.VisibleRowCount;
end;

procedure DrawGrid_SetOnClick(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnColumnMoved(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnColumnMoved := nil;
    TEventClass.Remove(AObj, @TEventClass.OnColumnMoved);
    Exit;
  end;
  AObj.OnColumnMoved := TEventClass.OnColumnMoved;
  TEventClass.Add(AObj, @TEventClass.OnColumnMoved, AEventId);
end;

procedure DrawGrid_SetOnContextPopup(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnContextPopup := nil;
    TEventClass.Remove(AObj, @TEventClass.OnContextPopup);
    Exit;
  end;
  AObj.OnContextPopup := TEventClass.OnContextPopup;
  TEventClass.Add(AObj, @TEventClass.OnContextPopup, AEventId);
end;

procedure DrawGrid_SetOnDblClick(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnDragDrop(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnDragOver(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnDrawCell(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDrawCell := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDrawCell);
    Exit;
  end;
  AObj.OnDrawCell := TEventClass.OnDrawCell;
  TEventClass.Add(AObj, @TEventClass.OnDrawCell, AEventId);
end;

procedure DrawGrid_SetOnEndDock(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDock := nil;
    TEventClass.Remove(AObj, @TEventClass.OnEndDock);
    Exit;
  end;
  AObj.OnEndDock := TEventClass.OnEndDock;
  TEventClass.Add(AObj, @TEventClass.OnEndDock, AEventId);
end;

procedure DrawGrid_SetOnEndDrag(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnEnter(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnExit(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnFixedCellClick(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnFixedCellClick := nil;
    TEventClass.Remove(AObj, @TEventClass.OnFixedCellClick);
    Exit;
  end;
  AObj.OnFixedCellClick := TEventClass.OnFixedCellClick;
  TEventClass.Add(AObj, @TEventClass.OnFixedCellClick, AEventId);
end;

procedure DrawGrid_SetOnGesture(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnGetEditMask(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetEditMask := nil;
    TEventClass.Remove(AObj, @TEventClass.OnGetEditMask);
    Exit;
  end;
  AObj.OnGetEditMask := TEventClass.OnGetEditMask;
  TEventClass.Add(AObj, @TEventClass.OnGetEditMask, AEventId);
end;

procedure DrawGrid_SetOnGetEditText(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetEditText := nil;
    TEventClass.Remove(AObj, @TEventClass.OnGetEditText);
    Exit;
  end;
  AObj.OnGetEditText := TEventClass.OnGetEditText;
  TEventClass.Add(AObj, @TEventClass.OnGetEditText, AEventId);
end;

procedure DrawGrid_SetOnKeyDown(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnKeyPress(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnKeyUp(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnMouseActivate(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseActivate := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseActivate);
    Exit;
  end;
  AObj.OnMouseActivate := TEventClass.OnMouseActivate;
  TEventClass.Add(AObj, @TEventClass.OnMouseActivate, AEventId);
end;

procedure DrawGrid_SetOnMouseDown(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnMouseEnter(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseEnter := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseEnter);
    Exit;
  end;
  AObj.OnMouseEnter := TEventClass.OnMouseEnter;
  TEventClass.Add(AObj, @TEventClass.OnMouseEnter, AEventId);
end;

procedure DrawGrid_SetOnMouseLeave(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseLeave := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseLeave);
    Exit;
  end;
  AObj.OnMouseLeave := TEventClass.OnMouseLeave;
  TEventClass.Add(AObj, @TEventClass.OnMouseLeave, AEventId);
end;

procedure DrawGrid_SetOnMouseMove(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnMouseUp(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
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

procedure DrawGrid_SetOnMouseWheelDown(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheelDown := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseWheelDown);
    Exit;
  end;
  AObj.OnMouseWheelDown := TEventClass.OnMouseWheelDown;
  TEventClass.Add(AObj, @TEventClass.OnMouseWheelDown, AEventId);
end;

procedure DrawGrid_SetOnMouseWheelUp(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheelUp := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseWheelUp);
    Exit;
  end;
  AObj.OnMouseWheelUp := TEventClass.OnMouseWheelUp;
  TEventClass.Add(AObj, @TEventClass.OnMouseWheelUp, AEventId);
end;

procedure DrawGrid_SetOnRowMoved(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnRowMoved := nil;
    TEventClass.Remove(AObj, @TEventClass.OnRowMoved);
    Exit;
  end;
  AObj.OnRowMoved := TEventClass.OnRowMoved;
  TEventClass.Add(AObj, @TEventClass.OnRowMoved, AEventId);
end;

procedure DrawGrid_SetOnSelectCell(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnSelectCell := nil;
    TEventClass.Remove(AObj, @TEventClass.OnSelectCell);
    Exit;
  end;
  AObj.OnSelectCell := TEventClass.OnSelectCell;
  TEventClass.Add(AObj, @TEventClass.OnSelectCell, AEventId);
end;

procedure DrawGrid_SetOnSetEditText(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnSetEditText := nil;
    TEventClass.Remove(AObj, @TEventClass.OnSetEditText);
    Exit;
  end;
  AObj.OnSetEditText := TEventClass.OnSetEditText;
  TEventClass.Add(AObj, @TEventClass.OnSetEditText, AEventId);
end;

procedure DrawGrid_SetOnStartDock(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnStartDock := nil;
    TEventClass.Remove(AObj, @TEventClass.OnStartDock);
    Exit;
  end;
  AObj.OnStartDock := TEventClass.OnStartDock;
  TEventClass.Add(AObj, @TEventClass.OnStartDock, AEventId);
end;

procedure DrawGrid_SetOnTopLeftChanged(AObj: TDrawGrid; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnTopLeftChanged := nil;
    TEventClass.Remove(AObj, @TEventClass.OnTopLeftChanged);
    Exit;
  end;
  AObj.OnTopLeftChanged := TEventClass.OnTopLeftChanged;
  TEventClass.Add(AObj, @TEventClass.OnTopLeftChanged, AEventId);
end;

function DrawGrid_GetCanvas(AObj: TDrawGrid): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function DrawGrid_GetCol(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.Col;
end;

procedure DrawGrid_SetCol(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.Col := AValue;
end;

function DrawGrid_GetEditorMode(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.EditorMode;
end;

procedure DrawGrid_SetEditorMode(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.EditorMode := AValue;
end;

function DrawGrid_GetGridHeight(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.GridHeight;
end;

function DrawGrid_GetGridWidth(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.GridWidth;
end;

function DrawGrid_GetLeftCol(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.LeftCol;
end;

procedure DrawGrid_SetLeftCol(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.LeftCol := AValue;
end;

procedure DrawGrid_GetSelection(AObj: TDrawGrid; var Result: TGridRect); stdcall;
begin
  Result :=  AObj.Selection;
end;

procedure DrawGrid_SetSelection(AObj: TDrawGrid; var AValue: TGridRect); stdcall;
begin
  AObj.Selection := AValue;
end;

function DrawGrid_GetRow(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.Row;
end;

procedure DrawGrid_SetRow(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.Row := AValue;
end;

function DrawGrid_GetTopRow(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.TopRow;
end;

procedure DrawGrid_SetTopRow(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.TopRow := AValue;
end;

function DrawGrid_GetTabStop(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure DrawGrid_SetTabStop(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function DrawGrid_GetDockClientCount(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function DrawGrid_GetDockSite(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure DrawGrid_SetDockSite(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function DrawGrid_GetAlignDisabled(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function DrawGrid_GetMouseInClient(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function DrawGrid_GetVisibleDockClientCount(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function DrawGrid_GetBrush(AObj: TDrawGrid): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function DrawGrid_GetControlCount(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function DrawGrid_GetHandle(AObj: TDrawGrid): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function DrawGrid_GetParentWindow(AObj: TDrawGrid): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure DrawGrid_SetParentWindow(AObj: TDrawGrid; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function DrawGrid_GetUseDockManager(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure DrawGrid_SetUseDockManager(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function DrawGrid_GetAction(AObj: TDrawGrid): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure DrawGrid_SetAction(AObj: TDrawGrid; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure DrawGrid_GetBoundsRect(AObj: TDrawGrid; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure DrawGrid_SetBoundsRect(AObj: TDrawGrid; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function DrawGrid_GetClientHeight(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure DrawGrid_SetClientHeight(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure DrawGrid_GetClientOrigin(AObj: TDrawGrid; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure DrawGrid_GetClientRect(AObj: TDrawGrid; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function DrawGrid_GetClientWidth(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure DrawGrid_SetClientWidth(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function DrawGrid_GetControlState(AObj: TDrawGrid): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure DrawGrid_SetControlState(AObj: TDrawGrid; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function DrawGrid_GetControlStyle(AObj: TDrawGrid): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure DrawGrid_SetControlStyle(AObj: TDrawGrid; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function DrawGrid_GetExplicitLeft(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function DrawGrid_GetExplicitTop(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function DrawGrid_GetExplicitWidth(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function DrawGrid_GetExplicitHeight(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function DrawGrid_GetFloating(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function DrawGrid_GetParent(AObj: TDrawGrid): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure DrawGrid_SetParent(AObj: TDrawGrid; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function DrawGrid_GetAlignWithMargins(AObj: TDrawGrid): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure DrawGrid_SetAlignWithMargins(AObj: TDrawGrid; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function DrawGrid_GetLeft(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure DrawGrid_SetLeft(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function DrawGrid_GetTop(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure DrawGrid_SetTop(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function DrawGrid_GetWidth(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure DrawGrid_SetWidth(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function DrawGrid_GetHeight(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure DrawGrid_SetHeight(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function DrawGrid_GetCursor(AObj: TDrawGrid): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure DrawGrid_SetCursor(AObj: TDrawGrid; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function DrawGrid_GetHint(AObj: TDrawGrid): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure DrawGrid_SetHint(AObj: TDrawGrid; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function DrawGrid_GetMargins(AObj: TDrawGrid): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure DrawGrid_SetMargins(AObj: TDrawGrid; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function DrawGrid_GetCustomHint(AObj: TDrawGrid): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure DrawGrid_SetCustomHint(AObj: TDrawGrid; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function DrawGrid_GetComponentCount(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function DrawGrid_GetComponentIndex(AObj: TDrawGrid): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure DrawGrid_SetComponentIndex(AObj: TDrawGrid; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function DrawGrid_GetOwner(AObj: TDrawGrid): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function DrawGrid_GetName(AObj: TDrawGrid): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure DrawGrid_SetName(AObj: TDrawGrid; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function DrawGrid_GetTag(AObj: TDrawGrid): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure DrawGrid_SetTag(AObj: TDrawGrid; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function DrawGrid_GetColWidths(AObj: TDrawGrid; Index: Integer): Integer; stdcall;
begin
  Result :=  AObj.ColWidths[Index];
end;

procedure DrawGrid_SetColWidths(AObj: TDrawGrid; Index: Integer; AValue: Integer); stdcall;
begin
  AObj.ColWidths[Index] := AValue;
end;

function DrawGrid_GetRowHeights(AObj: TDrawGrid; Index: Integer): Integer; stdcall;
begin
  Result :=  AObj.RowHeights[Index];
end;

procedure DrawGrid_SetRowHeights(AObj: TDrawGrid; Index: Integer; AValue: Integer); stdcall;
begin
  AObj.RowHeights[Index] := AValue;
end;

function DrawGrid_GetTabStops(AObj: TDrawGrid; Index: Integer): LongBool; stdcall;
begin
  Result :=  AObj.TabStops[Index];
end;

procedure DrawGrid_SetTabStops(AObj: TDrawGrid; Index: Integer; AValue: LongBool); stdcall;
begin
  AObj.TabStops[Index] := AValue;
end;

function DrawGrid_GetDockClients(AObj: TDrawGrid; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function DrawGrid_GetControls(AObj: TDrawGrid; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function DrawGrid_GetComponents(AObj: TDrawGrid; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function DrawGrid_StaticClassType: TClass; stdcall;
begin
  Result := TDrawGrid;
end;

exports
  DrawGrid_Create,
  DrawGrid_Free,
  DrawGrid_CellRect,
  DrawGrid_MouseToCell,
  DrawGrid_MouseCoord,
  DrawGrid_CanFocus,
  DrawGrid_ContainsControl,
  DrawGrid_ControlAtPos,
  DrawGrid_DisableAlign,
  DrawGrid_EnableAlign,
  DrawGrid_FindChildControl,
  DrawGrid_FlipChildren,
  DrawGrid_Focused,
  DrawGrid_HandleAllocated,
  DrawGrid_InsertControl,
  DrawGrid_Invalidate,
  DrawGrid_PaintTo,
  DrawGrid_RemoveControl,
  DrawGrid_Realign,
  DrawGrid_Repaint,
  DrawGrid_ScaleBy,
  DrawGrid_ScrollBy,
  DrawGrid_SetBounds,
  DrawGrid_SetFocus,
  DrawGrid_Update,
  DrawGrid_UpdateControlState,
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
  DrawGrid_SetTextBuf,
  DrawGrid_FindComponent,
  DrawGrid_GetNamePath,
  DrawGrid_Assign,
  DrawGrid_DisposeOf,
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
  DrawGrid_GetBevelEdges,
  DrawGrid_SetBevelEdges,
  DrawGrid_GetBevelInner,
  DrawGrid_SetBevelInner,
  DrawGrid_GetBevelKind,
  DrawGrid_SetBevelKind,
  DrawGrid_GetBevelOuter,
  DrawGrid_SetBevelOuter,
  DrawGrid_GetBiDiMode,
  DrawGrid_SetBiDiMode,
  DrawGrid_GetBorderStyle,
  DrawGrid_SetBorderStyle,
  DrawGrid_GetColor,
  DrawGrid_SetColor,
  DrawGrid_GetColCount,
  DrawGrid_SetColCount,
  DrawGrid_GetDefaultColWidth,
  DrawGrid_SetDefaultColWidth,
  DrawGrid_GetDefaultRowHeight,
  DrawGrid_SetDefaultRowHeight,
  DrawGrid_GetDefaultDrawing,
  DrawGrid_SetDefaultDrawing,
  DrawGrid_GetDoubleBuffered,
  DrawGrid_SetDoubleBuffered,
  DrawGrid_GetDragCursor,
  DrawGrid_SetDragCursor,
  DrawGrid_GetDragKind,
  DrawGrid_SetDragKind,
  DrawGrid_GetDragMode,
  DrawGrid_SetDragMode,
  DrawGrid_GetDrawingStyle,
  DrawGrid_SetDrawingStyle,
  DrawGrid_GetEnabled,
  DrawGrid_SetEnabled,
  DrawGrid_GetFixedColor,
  DrawGrid_SetFixedColor,
  DrawGrid_GetFixedCols,
  DrawGrid_SetFixedCols,
  DrawGrid_GetRowCount,
  DrawGrid_SetRowCount,
  DrawGrid_GetFixedRows,
  DrawGrid_SetFixedRows,
  DrawGrid_GetFont,
  DrawGrid_SetFont,
  DrawGrid_GetGradientEndColor,
  DrawGrid_SetGradientEndColor,
  DrawGrid_GetGradientStartColor,
  DrawGrid_SetGradientStartColor,
  DrawGrid_GetGridLineWidth,
  DrawGrid_SetGridLineWidth,
  DrawGrid_GetOptions,
  DrawGrid_SetOptions,
  DrawGrid_GetParentColor,
  DrawGrid_SetParentColor,
  DrawGrid_GetParentCtl3D,
  DrawGrid_SetParentCtl3D,
  DrawGrid_GetParentDoubleBuffered,
  DrawGrid_SetParentDoubleBuffered,
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
  DrawGrid_GetStyleElements,
  DrawGrid_SetStyleElements,
  DrawGrid_GetVisibleColCount,
  DrawGrid_GetVisibleRowCount,
  DrawGrid_SetOnClick,
  DrawGrid_SetOnColumnMoved,
  DrawGrid_SetOnContextPopup,
  DrawGrid_SetOnDblClick,
  DrawGrid_SetOnDragDrop,
  DrawGrid_SetOnDragOver,
  DrawGrid_SetOnDrawCell,
  DrawGrid_SetOnEndDock,
  DrawGrid_SetOnEndDrag,
  DrawGrid_SetOnEnter,
  DrawGrid_SetOnExit,
  DrawGrid_SetOnFixedCellClick,
  DrawGrid_SetOnGesture,
  DrawGrid_SetOnGetEditMask,
  DrawGrid_SetOnGetEditText,
  DrawGrid_SetOnKeyDown,
  DrawGrid_SetOnKeyPress,
  DrawGrid_SetOnKeyUp,
  DrawGrid_SetOnMouseActivate,
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
  DrawGrid_GetCol,
  DrawGrid_SetCol,
  DrawGrid_GetEditorMode,
  DrawGrid_SetEditorMode,
  DrawGrid_GetGridHeight,
  DrawGrid_GetGridWidth,
  DrawGrid_GetLeftCol,
  DrawGrid_SetLeftCol,
  DrawGrid_GetSelection,
  DrawGrid_SetSelection,
  DrawGrid_GetRow,
  DrawGrid_SetRow,
  DrawGrid_GetTopRow,
  DrawGrid_SetTopRow,
  DrawGrid_GetTabStop,
  DrawGrid_SetTabStop,
  DrawGrid_GetDockClientCount,
  DrawGrid_GetDockSite,
  DrawGrid_SetDockSite,
  DrawGrid_GetAlignDisabled,
  DrawGrid_GetMouseInClient,
  DrawGrid_GetVisibleDockClientCount,
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
  DrawGrid_GetClientOrigin,
  DrawGrid_GetClientRect,
  DrawGrid_GetClientWidth,
  DrawGrid_SetClientWidth,
  DrawGrid_GetControlState,
  DrawGrid_SetControlState,
  DrawGrid_GetControlStyle,
  DrawGrid_SetControlStyle,
  DrawGrid_GetExplicitLeft,
  DrawGrid_GetExplicitTop,
  DrawGrid_GetExplicitWidth,
  DrawGrid_GetExplicitHeight,
  DrawGrid_GetFloating,
  DrawGrid_GetParent,
  DrawGrid_SetParent,
  DrawGrid_GetAlignWithMargins,
  DrawGrid_SetAlignWithMargins,
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
  DrawGrid_GetMargins,
  DrawGrid_SetMargins,
  DrawGrid_GetCustomHint,
  DrawGrid_SetCustomHint,
  DrawGrid_GetComponentCount,
  DrawGrid_GetComponentIndex,
  DrawGrid_SetComponentIndex,
  DrawGrid_GetOwner,
  DrawGrid_GetName,
  DrawGrid_SetName,
  DrawGrid_GetTag,
  DrawGrid_SetTag,
  DrawGrid_GetColWidths,
  DrawGrid_SetColWidths,
  DrawGrid_GetRowHeights,
  DrawGrid_SetRowHeights,
  DrawGrid_GetTabStops,
  DrawGrid_SetTabStops,
  DrawGrid_GetDockClients,
  DrawGrid_GetControls,
  DrawGrid_GetComponents,
  DrawGrid_StaticClassType;

