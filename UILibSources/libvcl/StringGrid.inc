
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function StringGrid_Create(AOwner: TComponent): TStringGrid; stdcall;
begin
  Result :=  TStringGrid.Create(AOwner);
end;

procedure StringGrid_Free(AObj: TStringGrid); stdcall;
begin
  AObj.Free;
end;

procedure StringGrid_CellRect(AObj: TStringGrid; ACol: Integer; ARow: Integer; var Result: TRect); stdcall;
begin
  Result :=  AObj.CellRect(ACol, ARow);
end;

procedure StringGrid_MouseToCell(AObj: TStringGrid; X: Integer; Y: Integer; var ACol: Integer; var ARow: Integer); stdcall;
begin
  AObj.MouseToCell(X, Y, ACol, ARow);
end;

procedure StringGrid_MouseCoord(AObj: TStringGrid; X: Integer; Y: Integer; var Result: TGridCoord); stdcall;
begin
  Result :=  AObj.MouseCoord(X, Y);
end;

function StringGrid_CanFocus(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function StringGrid_ContainsControl(AObj: TStringGrid; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function StringGrid_ControlAtPos(AObj: TStringGrid; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure StringGrid_DisableAlign(AObj: TStringGrid); stdcall;
begin
  AObj.DisableAlign;
end;

procedure StringGrid_EnableAlign(AObj: TStringGrid); stdcall;
begin
  AObj.EnableAlign;
end;

function StringGrid_FindChildControl(AObj: TStringGrid; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure StringGrid_FlipChildren(AObj: TStringGrid; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function StringGrid_Focused(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function StringGrid_HandleAllocated(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure StringGrid_InsertControl(AObj: TStringGrid; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure StringGrid_Invalidate(AObj: TStringGrid); stdcall;
begin
  AObj.Invalidate;
end;

procedure StringGrid_PaintTo(AObj: TStringGrid; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure StringGrid_RemoveControl(AObj: TStringGrid; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure StringGrid_Realign(AObj: TStringGrid); stdcall;
begin
  AObj.Realign;
end;

procedure StringGrid_Repaint(AObj: TStringGrid); stdcall;
begin
  AObj.Repaint;
end;

procedure StringGrid_ScaleBy(AObj: TStringGrid; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure StringGrid_ScrollBy(AObj: TStringGrid; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure StringGrid_SetBounds(AObj: TStringGrid; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure StringGrid_SetFocus(AObj: TStringGrid); stdcall;
begin
  AObj.SetFocus;
end;

procedure StringGrid_Update(AObj: TStringGrid); stdcall;
begin
  AObj.Update;
end;

procedure StringGrid_UpdateControlState(AObj: TStringGrid); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure StringGrid_BringToFront(AObj: TStringGrid); stdcall;
begin
  AObj.BringToFront;
end;

procedure StringGrid_ClientToScreen(AObj: TStringGrid; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure StringGrid_ClientToParent(AObj: TStringGrid; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function StringGrid_Dragging(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function StringGrid_HasParent(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure StringGrid_Hide(AObj: TStringGrid); stdcall;
begin
  AObj.Hide;
end;

function StringGrid_Perform(AObj: TStringGrid; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure StringGrid_Refresh(AObj: TStringGrid); stdcall;
begin
  AObj.Refresh;
end;

procedure StringGrid_ScreenToClient(AObj: TStringGrid; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure StringGrid_ParentToClient(AObj: TStringGrid; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure StringGrid_SendToBack(AObj: TStringGrid); stdcall;
begin
  AObj.SendToBack;
end;

procedure StringGrid_Show(AObj: TStringGrid); stdcall;
begin
  AObj.Show;
end;

function StringGrid_GetTextBuf(AObj: TStringGrid; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function StringGrid_GetTextLen(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure StringGrid_SetTextBuf(AObj: TStringGrid; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function StringGrid_FindComponent(AObj: TStringGrid; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function StringGrid_GetNamePath(AObj: TStringGrid): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure StringGrid_Assign(AObj: TStringGrid; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure StringGrid_DisposeOf(AObj: TStringGrid); stdcall;
begin
  AObj.DisposeOf;
end;

function StringGrid_ClassType(AObj: TStringGrid): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function StringGrid_ClassName(AObj: TStringGrid): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function StringGrid_InstanceSize(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function StringGrid_InheritsFrom(AObj: TStringGrid; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function StringGrid_Equals(AObj: TStringGrid; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function StringGrid_GetHashCode(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function StringGrid_ToString(AObj: TStringGrid): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function StringGrid_GetAlign(AObj: TStringGrid): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure StringGrid_SetAlign(AObj: TStringGrid; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function StringGrid_GetAnchors(AObj: TStringGrid): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure StringGrid_SetAnchors(AObj: TStringGrid; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function StringGrid_GetBevelEdges(AObj: TStringGrid): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure StringGrid_SetBevelEdges(AObj: TStringGrid; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function StringGrid_GetBevelInner(AObj: TStringGrid): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure StringGrid_SetBevelInner(AObj: TStringGrid; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function StringGrid_GetBevelKind(AObj: TStringGrid): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure StringGrid_SetBevelKind(AObj: TStringGrid; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function StringGrid_GetBevelOuter(AObj: TStringGrid): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure StringGrid_SetBevelOuter(AObj: TStringGrid; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function StringGrid_GetBiDiMode(AObj: TStringGrid): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure StringGrid_SetBiDiMode(AObj: TStringGrid; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function StringGrid_GetBorderStyle(AObj: TStringGrid): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure StringGrid_SetBorderStyle(AObj: TStringGrid; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function StringGrid_GetColor(AObj: TStringGrid): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure StringGrid_SetColor(AObj: TStringGrid; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function StringGrid_GetColCount(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.ColCount;
end;

procedure StringGrid_SetColCount(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.ColCount := AValue;
end;

function StringGrid_GetDefaultColWidth(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.DefaultColWidth;
end;

procedure StringGrid_SetDefaultColWidth(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.DefaultColWidth := AValue;
end;

function StringGrid_GetDefaultRowHeight(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.DefaultRowHeight;
end;

procedure StringGrid_SetDefaultRowHeight(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.DefaultRowHeight := AValue;
end;

function StringGrid_GetDefaultDrawing(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.DefaultDrawing;
end;

procedure StringGrid_SetDefaultDrawing(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.DefaultDrawing := AValue;
end;

function StringGrid_GetDoubleBuffered(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure StringGrid_SetDoubleBuffered(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function StringGrid_GetDragCursor(AObj: TStringGrid): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure StringGrid_SetDragCursor(AObj: TStringGrid; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function StringGrid_GetDragKind(AObj: TStringGrid): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure StringGrid_SetDragKind(AObj: TStringGrid; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function StringGrid_GetDragMode(AObj: TStringGrid): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure StringGrid_SetDragMode(AObj: TStringGrid; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function StringGrid_GetDrawingStyle(AObj: TStringGrid): TGridDrawingStyle; stdcall;
begin
  Result :=  AObj.DrawingStyle;
end;

procedure StringGrid_SetDrawingStyle(AObj: TStringGrid; AValue: TGridDrawingStyle); stdcall;
begin
  AObj.DrawingStyle := AValue;
end;

function StringGrid_GetEnabled(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure StringGrid_SetEnabled(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function StringGrid_GetFixedColor(AObj: TStringGrid): TColor; stdcall;
begin
  Result :=  AObj.FixedColor;
end;

procedure StringGrid_SetFixedColor(AObj: TStringGrid; AValue: TColor); stdcall;
begin
  AObj.FixedColor := AValue;
end;

function StringGrid_GetFixedCols(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.FixedCols;
end;

procedure StringGrid_SetFixedCols(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.FixedCols := AValue;
end;

function StringGrid_GetRowCount(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.RowCount;
end;

procedure StringGrid_SetRowCount(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.RowCount := AValue;
end;

function StringGrid_GetFixedRows(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.FixedRows;
end;

procedure StringGrid_SetFixedRows(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.FixedRows := AValue;
end;

function StringGrid_GetFont(AObj: TStringGrid): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure StringGrid_SetFont(AObj: TStringGrid; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function StringGrid_GetGradientEndColor(AObj: TStringGrid): TColor; stdcall;
begin
  Result :=  AObj.GradientEndColor;
end;

procedure StringGrid_SetGradientEndColor(AObj: TStringGrid; AValue: TColor); stdcall;
begin
  AObj.GradientEndColor := AValue;
end;

function StringGrid_GetGradientStartColor(AObj: TStringGrid): TColor; stdcall;
begin
  Result :=  AObj.GradientStartColor;
end;

procedure StringGrid_SetGradientStartColor(AObj: TStringGrid; AValue: TColor); stdcall;
begin
  AObj.GradientStartColor := AValue;
end;

function StringGrid_GetGridLineWidth(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.GridLineWidth;
end;

procedure StringGrid_SetGridLineWidth(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.GridLineWidth := AValue;
end;

function StringGrid_GetOptions(AObj: TStringGrid): TGridOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure StringGrid_SetOptions(AObj: TStringGrid; AValue: TGridOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function StringGrid_GetParentColor(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure StringGrid_SetParentColor(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function StringGrid_GetParentCtl3D(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure StringGrid_SetParentCtl3D(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function StringGrid_GetParentDoubleBuffered(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure StringGrid_SetParentDoubleBuffered(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function StringGrid_GetParentFont(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure StringGrid_SetParentFont(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function StringGrid_GetParentShowHint(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure StringGrid_SetParentShowHint(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function StringGrid_GetPopupMenu(AObj: TStringGrid): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure StringGrid_SetPopupMenu(AObj: TStringGrid; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function StringGrid_GetScrollBars(AObj: TStringGrid): TScrollStyle; stdcall;
begin
  Result :=  AObj.ScrollBars;
end;

procedure StringGrid_SetScrollBars(AObj: TStringGrid; AValue: TScrollStyle); stdcall;
begin
  AObj.ScrollBars := AValue;
end;

function StringGrid_GetShowHint(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure StringGrid_SetShowHint(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function StringGrid_GetTabOrder(AObj: TStringGrid): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure StringGrid_SetTabOrder(AObj: TStringGrid; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function StringGrid_GetVisible(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure StringGrid_SetVisible(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function StringGrid_GetStyleElements(AObj: TStringGrid): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure StringGrid_SetStyleElements(AObj: TStringGrid; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function StringGrid_GetVisibleColCount(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.VisibleColCount;
end;

function StringGrid_GetVisibleRowCount(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.VisibleRowCount;
end;

procedure StringGrid_SetOnClick(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnColumnMoved(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnContextPopup(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnDblClick(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnDragDrop(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnDragOver(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnDrawCell(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnEndDock(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnEndDrag(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnEnter(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnExit(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnFixedCellClick(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnGesture(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnGetEditMask(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnGetEditText(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnKeyDown(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnKeyPress(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnKeyUp(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnMouseActivate(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnMouseDown(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnMouseEnter(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnMouseLeave(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnMouseMove(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnMouseUp(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnMouseWheelDown(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnMouseWheelUp(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnRowMoved(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnSelectCell(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnSetEditText(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnStartDock(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

procedure StringGrid_SetOnTopLeftChanged(AObj: TStringGrid; AEventId: NativeUInt); stdcall;
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

function StringGrid_GetCanvas(AObj: TStringGrid): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function StringGrid_GetCol(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.Col;
end;

procedure StringGrid_SetCol(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.Col := AValue;
end;

function StringGrid_GetEditorMode(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.EditorMode;
end;

procedure StringGrid_SetEditorMode(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.EditorMode := AValue;
end;

function StringGrid_GetGridHeight(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.GridHeight;
end;

function StringGrid_GetGridWidth(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.GridWidth;
end;

function StringGrid_GetLeftCol(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.LeftCol;
end;

procedure StringGrid_SetLeftCol(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.LeftCol := AValue;
end;

procedure StringGrid_GetSelection(AObj: TStringGrid; var Result: TGridRect); stdcall;
begin
  Result :=  AObj.Selection;
end;

procedure StringGrid_SetSelection(AObj: TStringGrid; var AValue: TGridRect); stdcall;
begin
  AObj.Selection := AValue;
end;

function StringGrid_GetRow(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.Row;
end;

procedure StringGrid_SetRow(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.Row := AValue;
end;

function StringGrid_GetTopRow(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.TopRow;
end;

procedure StringGrid_SetTopRow(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.TopRow := AValue;
end;

function StringGrid_GetTabStop(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure StringGrid_SetTabStop(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function StringGrid_GetDockClientCount(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function StringGrid_GetDockSite(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure StringGrid_SetDockSite(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function StringGrid_GetAlignDisabled(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function StringGrid_GetMouseInClient(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function StringGrid_GetVisibleDockClientCount(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function StringGrid_GetBrush(AObj: TStringGrid): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function StringGrid_GetControlCount(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function StringGrid_GetHandle(AObj: TStringGrid): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function StringGrid_GetParentWindow(AObj: TStringGrid): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure StringGrid_SetParentWindow(AObj: TStringGrid; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function StringGrid_GetUseDockManager(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure StringGrid_SetUseDockManager(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function StringGrid_GetAction(AObj: TStringGrid): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure StringGrid_SetAction(AObj: TStringGrid; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure StringGrid_GetBoundsRect(AObj: TStringGrid; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure StringGrid_SetBoundsRect(AObj: TStringGrid; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function StringGrid_GetClientHeight(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure StringGrid_SetClientHeight(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure StringGrid_GetClientOrigin(AObj: TStringGrid; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure StringGrid_GetClientRect(AObj: TStringGrid; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function StringGrid_GetClientWidth(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure StringGrid_SetClientWidth(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function StringGrid_GetControlState(AObj: TStringGrid): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure StringGrid_SetControlState(AObj: TStringGrid; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function StringGrid_GetControlStyle(AObj: TStringGrid): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure StringGrid_SetControlStyle(AObj: TStringGrid; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function StringGrid_GetExplicitLeft(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function StringGrid_GetExplicitTop(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function StringGrid_GetExplicitWidth(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function StringGrid_GetExplicitHeight(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function StringGrid_GetFloating(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function StringGrid_GetParent(AObj: TStringGrid): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure StringGrid_SetParent(AObj: TStringGrid; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function StringGrid_GetAlignWithMargins(AObj: TStringGrid): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure StringGrid_SetAlignWithMargins(AObj: TStringGrid; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function StringGrid_GetLeft(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure StringGrid_SetLeft(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function StringGrid_GetTop(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure StringGrid_SetTop(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function StringGrid_GetWidth(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure StringGrid_SetWidth(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function StringGrid_GetHeight(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure StringGrid_SetHeight(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function StringGrid_GetCursor(AObj: TStringGrid): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure StringGrid_SetCursor(AObj: TStringGrid; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function StringGrid_GetHint(AObj: TStringGrid): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure StringGrid_SetHint(AObj: TStringGrid; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function StringGrid_GetMargins(AObj: TStringGrid): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure StringGrid_SetMargins(AObj: TStringGrid; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function StringGrid_GetCustomHint(AObj: TStringGrid): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure StringGrid_SetCustomHint(AObj: TStringGrid; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function StringGrid_GetComponentCount(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function StringGrid_GetComponentIndex(AObj: TStringGrid): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure StringGrid_SetComponentIndex(AObj: TStringGrid; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function StringGrid_GetOwner(AObj: TStringGrid): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function StringGrid_GetName(AObj: TStringGrid): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure StringGrid_SetName(AObj: TStringGrid; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function StringGrid_GetTag(AObj: TStringGrid): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure StringGrid_SetTag(AObj: TStringGrid; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function StringGrid_GetCells(AObj: TStringGrid; ACol: Integer; ARow: Integer): PChar; stdcall;
begin
  Result :=  PChar(AObj.Cells[ACol, ARow]);
end;

procedure StringGrid_SetCells(AObj: TStringGrid; ACol: Integer; ARow: Integer; AValue: PChar); stdcall;
begin
  AObj.Cells[ACol, ARow] := AValue;
end;

function StringGrid_GetCols(AObj: TStringGrid; Index: Integer): TStrings; stdcall;
begin
  Result :=  AObj.Cols[Index];
end;

procedure StringGrid_SetCols(AObj: TStringGrid; Index: Integer; AValue: TStrings); stdcall;
begin
  AObj.Cols[Index] := AValue;
end;

function StringGrid_GetObjects(AObj: TStringGrid; ACol: Integer; ARow: Integer): TObject; stdcall;
begin
  Result :=  AObj.Objects[ACol, ARow];
end;

procedure StringGrid_SetObjects(AObj: TStringGrid; ACol: Integer; ARow: Integer; AValue: TObject); stdcall;
begin
  AObj.Objects[ACol, ARow] := AValue;
end;

function StringGrid_GetRows(AObj: TStringGrid; Index: Integer): TStrings; stdcall;
begin
  Result :=  AObj.Rows[Index];
end;

procedure StringGrid_SetRows(AObj: TStringGrid; Index: Integer; AValue: TStrings); stdcall;
begin
  AObj.Rows[Index] := AValue;
end;

function StringGrid_GetColWidths(AObj: TStringGrid; Index: Integer): Integer; stdcall;
begin
  Result :=  AObj.ColWidths[Index];
end;

procedure StringGrid_SetColWidths(AObj: TStringGrid; Index: Integer; AValue: Integer); stdcall;
begin
  AObj.ColWidths[Index] := AValue;
end;

function StringGrid_GetRowHeights(AObj: TStringGrid; Index: Integer): Integer; stdcall;
begin
  Result :=  AObj.RowHeights[Index];
end;

procedure StringGrid_SetRowHeights(AObj: TStringGrid; Index: Integer; AValue: Integer); stdcall;
begin
  AObj.RowHeights[Index] := AValue;
end;

function StringGrid_GetTabStops(AObj: TStringGrid; Index: Integer): LongBool; stdcall;
begin
  Result :=  AObj.TabStops[Index];
end;

procedure StringGrid_SetTabStops(AObj: TStringGrid; Index: Integer; AValue: LongBool); stdcall;
begin
  AObj.TabStops[Index] := AValue;
end;

function StringGrid_GetDockClients(AObj: TStringGrid; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function StringGrid_GetControls(AObj: TStringGrid; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function StringGrid_GetComponents(AObj: TStringGrid; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function StringGrid_StaticClassType: TClass; stdcall;
begin
  Result := TStringGrid;
end;

exports
  StringGrid_Create,
  StringGrid_Free,
  StringGrid_CellRect,
  StringGrid_MouseToCell,
  StringGrid_MouseCoord,
  StringGrid_CanFocus,
  StringGrid_ContainsControl,
  StringGrid_ControlAtPos,
  StringGrid_DisableAlign,
  StringGrid_EnableAlign,
  StringGrid_FindChildControl,
  StringGrid_FlipChildren,
  StringGrid_Focused,
  StringGrid_HandleAllocated,
  StringGrid_InsertControl,
  StringGrid_Invalidate,
  StringGrid_PaintTo,
  StringGrid_RemoveControl,
  StringGrid_Realign,
  StringGrid_Repaint,
  StringGrid_ScaleBy,
  StringGrid_ScrollBy,
  StringGrid_SetBounds,
  StringGrid_SetFocus,
  StringGrid_Update,
  StringGrid_UpdateControlState,
  StringGrid_BringToFront,
  StringGrid_ClientToScreen,
  StringGrid_ClientToParent,
  StringGrid_Dragging,
  StringGrid_HasParent,
  StringGrid_Hide,
  StringGrid_Perform,
  StringGrid_Refresh,
  StringGrid_ScreenToClient,
  StringGrid_ParentToClient,
  StringGrid_SendToBack,
  StringGrid_Show,
  StringGrid_GetTextBuf,
  StringGrid_GetTextLen,
  StringGrid_SetTextBuf,
  StringGrid_FindComponent,
  StringGrid_GetNamePath,
  StringGrid_Assign,
  StringGrid_DisposeOf,
  StringGrid_ClassType,
  StringGrid_ClassName,
  StringGrid_InstanceSize,
  StringGrid_InheritsFrom,
  StringGrid_Equals,
  StringGrid_GetHashCode,
  StringGrid_ToString,
  StringGrid_GetAlign,
  StringGrid_SetAlign,
  StringGrid_GetAnchors,
  StringGrid_SetAnchors,
  StringGrid_GetBevelEdges,
  StringGrid_SetBevelEdges,
  StringGrid_GetBevelInner,
  StringGrid_SetBevelInner,
  StringGrid_GetBevelKind,
  StringGrid_SetBevelKind,
  StringGrid_GetBevelOuter,
  StringGrid_SetBevelOuter,
  StringGrid_GetBiDiMode,
  StringGrid_SetBiDiMode,
  StringGrid_GetBorderStyle,
  StringGrid_SetBorderStyle,
  StringGrid_GetColor,
  StringGrid_SetColor,
  StringGrid_GetColCount,
  StringGrid_SetColCount,
  StringGrid_GetDefaultColWidth,
  StringGrid_SetDefaultColWidth,
  StringGrid_GetDefaultRowHeight,
  StringGrid_SetDefaultRowHeight,
  StringGrid_GetDefaultDrawing,
  StringGrid_SetDefaultDrawing,
  StringGrid_GetDoubleBuffered,
  StringGrid_SetDoubleBuffered,
  StringGrid_GetDragCursor,
  StringGrid_SetDragCursor,
  StringGrid_GetDragKind,
  StringGrid_SetDragKind,
  StringGrid_GetDragMode,
  StringGrid_SetDragMode,
  StringGrid_GetDrawingStyle,
  StringGrid_SetDrawingStyle,
  StringGrid_GetEnabled,
  StringGrid_SetEnabled,
  StringGrid_GetFixedColor,
  StringGrid_SetFixedColor,
  StringGrid_GetFixedCols,
  StringGrid_SetFixedCols,
  StringGrid_GetRowCount,
  StringGrid_SetRowCount,
  StringGrid_GetFixedRows,
  StringGrid_SetFixedRows,
  StringGrid_GetFont,
  StringGrid_SetFont,
  StringGrid_GetGradientEndColor,
  StringGrid_SetGradientEndColor,
  StringGrid_GetGradientStartColor,
  StringGrid_SetGradientStartColor,
  StringGrid_GetGridLineWidth,
  StringGrid_SetGridLineWidth,
  StringGrid_GetOptions,
  StringGrid_SetOptions,
  StringGrid_GetParentColor,
  StringGrid_SetParentColor,
  StringGrid_GetParentCtl3D,
  StringGrid_SetParentCtl3D,
  StringGrid_GetParentDoubleBuffered,
  StringGrid_SetParentDoubleBuffered,
  StringGrid_GetParentFont,
  StringGrid_SetParentFont,
  StringGrid_GetParentShowHint,
  StringGrid_SetParentShowHint,
  StringGrid_GetPopupMenu,
  StringGrid_SetPopupMenu,
  StringGrid_GetScrollBars,
  StringGrid_SetScrollBars,
  StringGrid_GetShowHint,
  StringGrid_SetShowHint,
  StringGrid_GetTabOrder,
  StringGrid_SetTabOrder,
  StringGrid_GetVisible,
  StringGrid_SetVisible,
  StringGrid_GetStyleElements,
  StringGrid_SetStyleElements,
  StringGrid_GetVisibleColCount,
  StringGrid_GetVisibleRowCount,
  StringGrid_SetOnClick,
  StringGrid_SetOnColumnMoved,
  StringGrid_SetOnContextPopup,
  StringGrid_SetOnDblClick,
  StringGrid_SetOnDragDrop,
  StringGrid_SetOnDragOver,
  StringGrid_SetOnDrawCell,
  StringGrid_SetOnEndDock,
  StringGrid_SetOnEndDrag,
  StringGrid_SetOnEnter,
  StringGrid_SetOnExit,
  StringGrid_SetOnFixedCellClick,
  StringGrid_SetOnGesture,
  StringGrid_SetOnGetEditMask,
  StringGrid_SetOnGetEditText,
  StringGrid_SetOnKeyDown,
  StringGrid_SetOnKeyPress,
  StringGrid_SetOnKeyUp,
  StringGrid_SetOnMouseActivate,
  StringGrid_SetOnMouseDown,
  StringGrid_SetOnMouseEnter,
  StringGrid_SetOnMouseLeave,
  StringGrid_SetOnMouseMove,
  StringGrid_SetOnMouseUp,
  StringGrid_SetOnMouseWheelDown,
  StringGrid_SetOnMouseWheelUp,
  StringGrid_SetOnRowMoved,
  StringGrid_SetOnSelectCell,
  StringGrid_SetOnSetEditText,
  StringGrid_SetOnStartDock,
  StringGrid_SetOnTopLeftChanged,
  StringGrid_GetCanvas,
  StringGrid_GetCol,
  StringGrid_SetCol,
  StringGrid_GetEditorMode,
  StringGrid_SetEditorMode,
  StringGrid_GetGridHeight,
  StringGrid_GetGridWidth,
  StringGrid_GetLeftCol,
  StringGrid_SetLeftCol,
  StringGrid_GetSelection,
  StringGrid_SetSelection,
  StringGrid_GetRow,
  StringGrid_SetRow,
  StringGrid_GetTopRow,
  StringGrid_SetTopRow,
  StringGrid_GetTabStop,
  StringGrid_SetTabStop,
  StringGrid_GetDockClientCount,
  StringGrid_GetDockSite,
  StringGrid_SetDockSite,
  StringGrid_GetAlignDisabled,
  StringGrid_GetMouseInClient,
  StringGrid_GetVisibleDockClientCount,
  StringGrid_GetBrush,
  StringGrid_GetControlCount,
  StringGrid_GetHandle,
  StringGrid_GetParentWindow,
  StringGrid_SetParentWindow,
  StringGrid_GetUseDockManager,
  StringGrid_SetUseDockManager,
  StringGrid_GetAction,
  StringGrid_SetAction,
  StringGrid_GetBoundsRect,
  StringGrid_SetBoundsRect,
  StringGrid_GetClientHeight,
  StringGrid_SetClientHeight,
  StringGrid_GetClientOrigin,
  StringGrid_GetClientRect,
  StringGrid_GetClientWidth,
  StringGrid_SetClientWidth,
  StringGrid_GetControlState,
  StringGrid_SetControlState,
  StringGrid_GetControlStyle,
  StringGrid_SetControlStyle,
  StringGrid_GetExplicitLeft,
  StringGrid_GetExplicitTop,
  StringGrid_GetExplicitWidth,
  StringGrid_GetExplicitHeight,
  StringGrid_GetFloating,
  StringGrid_GetParent,
  StringGrid_SetParent,
  StringGrid_GetAlignWithMargins,
  StringGrid_SetAlignWithMargins,
  StringGrid_GetLeft,
  StringGrid_SetLeft,
  StringGrid_GetTop,
  StringGrid_SetTop,
  StringGrid_GetWidth,
  StringGrid_SetWidth,
  StringGrid_GetHeight,
  StringGrid_SetHeight,
  StringGrid_GetCursor,
  StringGrid_SetCursor,
  StringGrid_GetHint,
  StringGrid_SetHint,
  StringGrid_GetMargins,
  StringGrid_SetMargins,
  StringGrid_GetCustomHint,
  StringGrid_SetCustomHint,
  StringGrid_GetComponentCount,
  StringGrid_GetComponentIndex,
  StringGrid_SetComponentIndex,
  StringGrid_GetOwner,
  StringGrid_GetName,
  StringGrid_SetName,
  StringGrid_GetTag,
  StringGrid_SetTag,
  StringGrid_GetCells,
  StringGrid_SetCells,
  StringGrid_GetCols,
  StringGrid_SetCols,
  StringGrid_GetObjects,
  StringGrid_SetObjects,
  StringGrid_GetRows,
  StringGrid_SetRows,
  StringGrid_GetColWidths,
  StringGrid_SetColWidths,
  StringGrid_GetRowHeights,
  StringGrid_SetRowHeights,
  StringGrid_GetTabStops,
  StringGrid_SetTabStops,
  StringGrid_GetDockClients,
  StringGrid_GetControls,
  StringGrid_GetComponents,
  StringGrid_StaticClassType;

