
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ValueListEditor_Create(AOwner: TComponent): TValueListEditor; stdcall;
begin
  Result :=  TValueListEditor.Create(AOwner);
end;

procedure ValueListEditor_Free(AObj: TValueListEditor); stdcall;
begin
  AObj.Free;
end;

procedure ValueListEditor_Refresh(AObj: TValueListEditor); stdcall;
begin
  AObj.Refresh;
end;

procedure ValueListEditor_CellRect(AObj: TValueListEditor; ACol: Integer; ARow: Integer; var Result: TRect); stdcall;
begin
  Result :=  AObj.CellRect(ACol, ARow);
end;

procedure ValueListEditor_MouseToCell(AObj: TValueListEditor; X: Integer; Y: Integer; var ACol: Integer; var ARow: Integer); stdcall;
begin
  AObj.MouseToCell(X, Y, ACol, ARow);
end;

procedure ValueListEditor_MouseCoord(AObj: TValueListEditor; X: Integer; Y: Integer; var Result: TGridCoord); stdcall;
begin
  Result :=  AObj.MouseCoord(X, Y);
end;

function ValueListEditor_CanFocus(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function ValueListEditor_ContainsControl(AObj: TValueListEditor; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function ValueListEditor_ControlAtPos(AObj: TValueListEditor; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure ValueListEditor_DisableAlign(AObj: TValueListEditor); stdcall;
begin
  AObj.DisableAlign;
end;

procedure ValueListEditor_EnableAlign(AObj: TValueListEditor); stdcall;
begin
  AObj.EnableAlign;
end;

function ValueListEditor_FindChildControl(AObj: TValueListEditor; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure ValueListEditor_FlipChildren(AObj: TValueListEditor; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function ValueListEditor_Focused(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function ValueListEditor_HandleAllocated(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ValueListEditor_InsertControl(AObj: TValueListEditor; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure ValueListEditor_Invalidate(AObj: TValueListEditor); stdcall;
begin
  AObj.Invalidate;
end;

procedure ValueListEditor_PaintTo(AObj: TValueListEditor; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure ValueListEditor_RemoveControl(AObj: TValueListEditor; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure ValueListEditor_Realign(AObj: TValueListEditor); stdcall;
begin
  AObj.Realign;
end;

procedure ValueListEditor_Repaint(AObj: TValueListEditor); stdcall;
begin
  AObj.Repaint;
end;

procedure ValueListEditor_ScaleBy(AObj: TValueListEditor; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure ValueListEditor_ScrollBy(AObj: TValueListEditor; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure ValueListEditor_SetBounds(AObj: TValueListEditor; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ValueListEditor_SetFocus(AObj: TValueListEditor); stdcall;
begin
  AObj.SetFocus;
end;

procedure ValueListEditor_Update(AObj: TValueListEditor); stdcall;
begin
  AObj.Update;
end;

procedure ValueListEditor_UpdateControlState(AObj: TValueListEditor); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure ValueListEditor_BringToFront(AObj: TValueListEditor); stdcall;
begin
  AObj.BringToFront;
end;

procedure ValueListEditor_ClientToScreen(AObj: TValueListEditor; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure ValueListEditor_ClientToParent(AObj: TValueListEditor; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function ValueListEditor_Dragging(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function ValueListEditor_HasParent(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ValueListEditor_Hide(AObj: TValueListEditor); stdcall;
begin
  AObj.Hide;
end;

function ValueListEditor_Perform(AObj: TValueListEditor; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ValueListEditor_ScreenToClient(AObj: TValueListEditor; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure ValueListEditor_ParentToClient(AObj: TValueListEditor; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure ValueListEditor_SendToBack(AObj: TValueListEditor); stdcall;
begin
  AObj.SendToBack;
end;

procedure ValueListEditor_Show(AObj: TValueListEditor); stdcall;
begin
  AObj.Show;
end;

function ValueListEditor_GetTextBuf(AObj: TValueListEditor; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ValueListEditor_GetTextLen(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure ValueListEditor_SetTextBuf(AObj: TValueListEditor; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function ValueListEditor_FindComponent(AObj: TValueListEditor; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ValueListEditor_GetNamePath(AObj: TValueListEditor): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ValueListEditor_Assign(AObj: TValueListEditor; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ValueListEditor_DisposeOf(AObj: TValueListEditor); stdcall;
begin
  AObj.DisposeOf;
end;

function ValueListEditor_ClassType(AObj: TValueListEditor): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function ValueListEditor_ClassName(AObj: TValueListEditor): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ValueListEditor_InstanceSize(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function ValueListEditor_InheritsFrom(AObj: TValueListEditor; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function ValueListEditor_Equals(AObj: TValueListEditor; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ValueListEditor_GetHashCode(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ValueListEditor_ToString(AObj: TValueListEditor): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ValueListEditor_GetColCount(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.ColCount;
end;

procedure ValueListEditor_SetColCount(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.ColCount := AValue;
end;

function ValueListEditor_GetRowCount(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.RowCount;
end;

function ValueListEditor_GetVisibleColCount(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.VisibleColCount;
end;

function ValueListEditor_GetVisibleRowCount(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.VisibleRowCount;
end;

function ValueListEditor_GetAlign(AObj: TValueListEditor): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ValueListEditor_SetAlign(AObj: TValueListEditor; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ValueListEditor_GetAnchors(AObj: TValueListEditor): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ValueListEditor_SetAnchors(AObj: TValueListEditor; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ValueListEditor_GetBiDiMode(AObj: TValueListEditor): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ValueListEditor_SetBiDiMode(AObj: TValueListEditor; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function ValueListEditor_GetBorderStyle(AObj: TValueListEditor): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure ValueListEditor_SetBorderStyle(AObj: TValueListEditor; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function ValueListEditor_GetColor(AObj: TValueListEditor): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure ValueListEditor_SetColor(AObj: TValueListEditor; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function ValueListEditor_GetConstraints(AObj: TValueListEditor): TSizeConstraints; stdcall;
begin
  Result :=  AObj.Constraints;
end;

procedure ValueListEditor_SetConstraints(AObj: TValueListEditor; AValue: TSizeConstraints); stdcall;
begin
  AObj.Constraints := AValue;
end;

function ValueListEditor_GetCtl3D(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.Ctl3D;
end;

procedure ValueListEditor_SetCtl3D(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.Ctl3D := AValue;
end;

function ValueListEditor_GetDefaultColWidth(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.DefaultColWidth;
end;

procedure ValueListEditor_SetDefaultColWidth(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.DefaultColWidth := AValue;
end;

function ValueListEditor_GetDefaultDrawing(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.DefaultDrawing;
end;

procedure ValueListEditor_SetDefaultDrawing(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.DefaultDrawing := AValue;
end;

function ValueListEditor_GetDefaultRowHeight(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.DefaultRowHeight;
end;

procedure ValueListEditor_SetDefaultRowHeight(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.DefaultRowHeight := AValue;
end;

function ValueListEditor_GetDoubleBuffered(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ValueListEditor_SetDoubleBuffered(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function ValueListEditor_GetDragCursor(AObj: TValueListEditor): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure ValueListEditor_SetDragCursor(AObj: TValueListEditor; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function ValueListEditor_GetDragKind(AObj: TValueListEditor): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure ValueListEditor_SetDragKind(AObj: TValueListEditor; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function ValueListEditor_GetDragMode(AObj: TValueListEditor): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure ValueListEditor_SetDragMode(AObj: TValueListEditor; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function ValueListEditor_GetDrawingStyle(AObj: TValueListEditor): TGridDrawingStyle; stdcall;
begin
  Result :=  AObj.DrawingStyle;
end;

procedure ValueListEditor_SetDrawingStyle(AObj: TValueListEditor; AValue: TGridDrawingStyle); stdcall;
begin
  AObj.DrawingStyle := AValue;
end;

function ValueListEditor_GetEnabled(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ValueListEditor_SetEnabled(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ValueListEditor_GetFixedColor(AObj: TValueListEditor): TColor; stdcall;
begin
  Result :=  AObj.FixedColor;
end;

procedure ValueListEditor_SetFixedColor(AObj: TValueListEditor; AValue: TColor); stdcall;
begin
  AObj.FixedColor := AValue;
end;

function ValueListEditor_GetFixedCols(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.FixedCols;
end;

procedure ValueListEditor_SetFixedCols(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.FixedCols := AValue;
end;

function ValueListEditor_GetFont(AObj: TValueListEditor): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure ValueListEditor_SetFont(AObj: TValueListEditor; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function ValueListEditor_GetGradientEndColor(AObj: TValueListEditor): TColor; stdcall;
begin
  Result :=  AObj.GradientEndColor;
end;

procedure ValueListEditor_SetGradientEndColor(AObj: TValueListEditor; AValue: TColor); stdcall;
begin
  AObj.GradientEndColor := AValue;
end;

function ValueListEditor_GetGradientStartColor(AObj: TValueListEditor): TColor; stdcall;
begin
  Result :=  AObj.GradientStartColor;
end;

procedure ValueListEditor_SetGradientStartColor(AObj: TValueListEditor; AValue: TColor); stdcall;
begin
  AObj.GradientStartColor := AValue;
end;

function ValueListEditor_GetGridLineWidth(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.GridLineWidth;
end;

procedure ValueListEditor_SetGridLineWidth(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.GridLineWidth := AValue;
end;

function ValueListEditor_GetOptions(AObj: TValueListEditor): TGridOptions; stdcall;
begin
  Result :=  AObj.Options;
end;

procedure ValueListEditor_SetOptions(AObj: TValueListEditor; AValue: TGridOptions); stdcall;
begin
  AObj.Options := AValue;
end;

function ValueListEditor_GetParentColor(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure ValueListEditor_SetParentColor(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function ValueListEditor_GetParentCtl3D(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure ValueListEditor_SetParentCtl3D(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function ValueListEditor_GetParentDoubleBuffered(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure ValueListEditor_SetParentDoubleBuffered(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function ValueListEditor_GetParentFont(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure ValueListEditor_SetParentFont(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function ValueListEditor_GetParentShowHint(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ValueListEditor_SetParentShowHint(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ValueListEditor_GetPopupMenu(AObj: TValueListEditor): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ValueListEditor_SetPopupMenu(AObj: TValueListEditor; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ValueListEditor_GetScrollBars(AObj: TValueListEditor): TScrollStyle; stdcall;
begin
  Result :=  AObj.ScrollBars;
end;

procedure ValueListEditor_SetScrollBars(AObj: TValueListEditor; AValue: TScrollStyle); stdcall;
begin
  AObj.ScrollBars := AValue;
end;

function ValueListEditor_GetShowHint(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ValueListEditor_SetShowHint(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ValueListEditor_GetStrings(AObj: TValueListEditor): TStrings; stdcall;
begin
  Result :=  AObj.Strings;
end;

procedure ValueListEditor_SetStrings(AObj: TValueListEditor; AValue: TStrings); stdcall;
begin
  AObj.Strings := AValue;
end;

function ValueListEditor_GetTabOrder(AObj: TValueListEditor): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure ValueListEditor_SetTabOrder(AObj: TValueListEditor; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function ValueListEditor_GetVisible(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ValueListEditor_SetVisible(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function ValueListEditor_GetStyleElements(AObj: TValueListEditor): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ValueListEditor_SetStyleElements(AObj: TValueListEditor; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure ValueListEditor_SetOnClick(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnContextPopup(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnDblClick(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnDragDrop(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnDragOver(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnDrawCell(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnEndDock(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnEndDrag(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnEnter(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnExit(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnGesture(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnGetEditMask(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnGetEditText(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnKeyDown(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnKeyPress(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnKeyUp(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnMouseActivate(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnMouseDown(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnMouseEnter(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnMouseLeave(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnMouseMove(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnMouseUp(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnMouseWheelDown(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnMouseWheelUp(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnRowMoved(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnSelectCell(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnSetEditText(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnStartDock(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

procedure ValueListEditor_SetOnTopLeftChanged(AObj: TValueListEditor; AEventId: NativeUInt); stdcall;
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

function ValueListEditor_GetCanvas(AObj: TValueListEditor): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function ValueListEditor_GetCol(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.Col;
end;

procedure ValueListEditor_SetCol(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.Col := AValue;
end;

function ValueListEditor_GetEditorMode(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.EditorMode;
end;

procedure ValueListEditor_SetEditorMode(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.EditorMode := AValue;
end;

function ValueListEditor_GetGridHeight(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.GridHeight;
end;

function ValueListEditor_GetGridWidth(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.GridWidth;
end;

function ValueListEditor_GetLeftCol(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.LeftCol;
end;

procedure ValueListEditor_SetLeftCol(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.LeftCol := AValue;
end;

procedure ValueListEditor_GetSelection(AObj: TValueListEditor; var Result: TGridRect); stdcall;
begin
  Result :=  AObj.Selection;
end;

procedure ValueListEditor_SetSelection(AObj: TValueListEditor; var AValue: TGridRect); stdcall;
begin
  AObj.Selection := AValue;
end;

function ValueListEditor_GetRow(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.Row;
end;

procedure ValueListEditor_SetRow(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.Row := AValue;
end;

function ValueListEditor_GetTopRow(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.TopRow;
end;

procedure ValueListEditor_SetTopRow(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.TopRow := AValue;
end;

function ValueListEditor_GetTabStop(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure ValueListEditor_SetTabStop(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function ValueListEditor_GetDockClientCount(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function ValueListEditor_GetDockSite(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure ValueListEditor_SetDockSite(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function ValueListEditor_GetAlignDisabled(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function ValueListEditor_GetMouseInClient(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function ValueListEditor_GetVisibleDockClientCount(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function ValueListEditor_GetBrush(AObj: TValueListEditor): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function ValueListEditor_GetControlCount(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function ValueListEditor_GetHandle(AObj: TValueListEditor): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ValueListEditor_GetParentWindow(AObj: TValueListEditor): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ValueListEditor_SetParentWindow(AObj: TValueListEditor; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function ValueListEditor_GetUseDockManager(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure ValueListEditor_SetUseDockManager(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function ValueListEditor_GetAction(AObj: TValueListEditor): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ValueListEditor_SetAction(AObj: TValueListEditor; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure ValueListEditor_GetBoundsRect(AObj: TValueListEditor; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ValueListEditor_SetBoundsRect(AObj: TValueListEditor; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ValueListEditor_GetClientHeight(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ValueListEditor_SetClientHeight(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ValueListEditor_GetClientOrigin(AObj: TValueListEditor; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure ValueListEditor_GetClientRect(AObj: TValueListEditor; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ValueListEditor_GetClientWidth(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ValueListEditor_SetClientWidth(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ValueListEditor_GetControlState(AObj: TValueListEditor): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure ValueListEditor_SetControlState(AObj: TValueListEditor; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function ValueListEditor_GetControlStyle(AObj: TValueListEditor): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure ValueListEditor_SetControlStyle(AObj: TValueListEditor; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function ValueListEditor_GetExplicitLeft(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ValueListEditor_GetExplicitTop(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ValueListEditor_GetExplicitWidth(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ValueListEditor_GetExplicitHeight(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ValueListEditor_GetFloating(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function ValueListEditor_GetParent(AObj: TValueListEditor): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ValueListEditor_SetParent(AObj: TValueListEditor; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ValueListEditor_GetAlignWithMargins(AObj: TValueListEditor): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ValueListEditor_SetAlignWithMargins(AObj: TValueListEditor; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ValueListEditor_GetLeft(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ValueListEditor_SetLeft(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ValueListEditor_GetTop(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ValueListEditor_SetTop(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ValueListEditor_GetWidth(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ValueListEditor_SetWidth(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ValueListEditor_GetHeight(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ValueListEditor_SetHeight(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ValueListEditor_GetCursor(AObj: TValueListEditor): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ValueListEditor_SetCursor(AObj: TValueListEditor; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ValueListEditor_GetHint(AObj: TValueListEditor): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ValueListEditor_SetHint(AObj: TValueListEditor; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ValueListEditor_GetMargins(AObj: TValueListEditor): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ValueListEditor_SetMargins(AObj: TValueListEditor; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ValueListEditor_GetCustomHint(AObj: TValueListEditor): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ValueListEditor_SetCustomHint(AObj: TValueListEditor; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ValueListEditor_GetComponentCount(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ValueListEditor_GetComponentIndex(AObj: TValueListEditor): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ValueListEditor_SetComponentIndex(AObj: TValueListEditor; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ValueListEditor_GetOwner(AObj: TValueListEditor): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ValueListEditor_GetName(AObj: TValueListEditor): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ValueListEditor_SetName(AObj: TValueListEditor; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ValueListEditor_GetTag(AObj: TValueListEditor): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ValueListEditor_SetTag(AObj: TValueListEditor; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ValueListEditor_GetCells(AObj: TValueListEditor; ACol: Integer; ARow: Integer): PChar; stdcall;
begin
  Result :=  PChar(AObj.Cells[ACol, ARow]);
end;

procedure ValueListEditor_SetCells(AObj: TValueListEditor; ACol: Integer; ARow: Integer; AValue: PChar); stdcall;
begin
  AObj.Cells[ACol, ARow] := AValue;
end;

function ValueListEditor_GetValues(AObj: TValueListEditor; Key: PChar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Values[Key]);
end;

procedure ValueListEditor_SetValues(AObj: TValueListEditor; Key: string; AValue: PChar); stdcall;
begin
  AObj.Values[Key] := AValue;
end;

function ValueListEditor_GetColWidths(AObj: TValueListEditor; Index: Integer): Integer; stdcall;
begin
  Result :=  AObj.ColWidths[Index];
end;

procedure ValueListEditor_SetColWidths(AObj: TValueListEditor; Index: Integer; AValue: Integer); stdcall;
begin
  AObj.ColWidths[Index] := AValue;
end;

function ValueListEditor_GetRowHeights(AObj: TValueListEditor; Index: Integer): Integer; stdcall;
begin
  Result :=  AObj.RowHeights[Index];
end;

procedure ValueListEditor_SetRowHeights(AObj: TValueListEditor; Index: Integer; AValue: Integer); stdcall;
begin
  AObj.RowHeights[Index] := AValue;
end;

function ValueListEditor_GetTabStops(AObj: TValueListEditor; Index: Integer): LongBool; stdcall;
begin
  Result :=  AObj.TabStops[Index];
end;

procedure ValueListEditor_SetTabStops(AObj: TValueListEditor; Index: Integer; AValue: LongBool); stdcall;
begin
  AObj.TabStops[Index] := AValue;
end;

function ValueListEditor_GetDockClients(AObj: TValueListEditor; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function ValueListEditor_GetControls(AObj: TValueListEditor; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ValueListEditor_GetComponents(AObj: TValueListEditor; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function ValueListEditor_StaticClassType: TClass; stdcall;
begin
  Result := TValueListEditor;
end;

exports
  ValueListEditor_Create,
  ValueListEditor_Free,
  ValueListEditor_Refresh,
  ValueListEditor_CellRect,
  ValueListEditor_MouseToCell,
  ValueListEditor_MouseCoord,
  ValueListEditor_CanFocus,
  ValueListEditor_ContainsControl,
  ValueListEditor_ControlAtPos,
  ValueListEditor_DisableAlign,
  ValueListEditor_EnableAlign,
  ValueListEditor_FindChildControl,
  ValueListEditor_FlipChildren,
  ValueListEditor_Focused,
  ValueListEditor_HandleAllocated,
  ValueListEditor_InsertControl,
  ValueListEditor_Invalidate,
  ValueListEditor_PaintTo,
  ValueListEditor_RemoveControl,
  ValueListEditor_Realign,
  ValueListEditor_Repaint,
  ValueListEditor_ScaleBy,
  ValueListEditor_ScrollBy,
  ValueListEditor_SetBounds,
  ValueListEditor_SetFocus,
  ValueListEditor_Update,
  ValueListEditor_UpdateControlState,
  ValueListEditor_BringToFront,
  ValueListEditor_ClientToScreen,
  ValueListEditor_ClientToParent,
  ValueListEditor_Dragging,
  ValueListEditor_HasParent,
  ValueListEditor_Hide,
  ValueListEditor_Perform,
  ValueListEditor_ScreenToClient,
  ValueListEditor_ParentToClient,
  ValueListEditor_SendToBack,
  ValueListEditor_Show,
  ValueListEditor_GetTextBuf,
  ValueListEditor_GetTextLen,
  ValueListEditor_SetTextBuf,
  ValueListEditor_FindComponent,
  ValueListEditor_GetNamePath,
  ValueListEditor_Assign,
  ValueListEditor_DisposeOf,
  ValueListEditor_ClassType,
  ValueListEditor_ClassName,
  ValueListEditor_InstanceSize,
  ValueListEditor_InheritsFrom,
  ValueListEditor_Equals,
  ValueListEditor_GetHashCode,
  ValueListEditor_ToString,
  ValueListEditor_GetColCount,
  ValueListEditor_SetColCount,
  ValueListEditor_GetRowCount,
  ValueListEditor_GetVisibleColCount,
  ValueListEditor_GetVisibleRowCount,
  ValueListEditor_GetAlign,
  ValueListEditor_SetAlign,
  ValueListEditor_GetAnchors,
  ValueListEditor_SetAnchors,
  ValueListEditor_GetBiDiMode,
  ValueListEditor_SetBiDiMode,
  ValueListEditor_GetBorderStyle,
  ValueListEditor_SetBorderStyle,
  ValueListEditor_GetColor,
  ValueListEditor_SetColor,
  ValueListEditor_GetConstraints,
  ValueListEditor_SetConstraints,
  ValueListEditor_GetCtl3D,
  ValueListEditor_SetCtl3D,
  ValueListEditor_GetDefaultColWidth,
  ValueListEditor_SetDefaultColWidth,
  ValueListEditor_GetDefaultDrawing,
  ValueListEditor_SetDefaultDrawing,
  ValueListEditor_GetDefaultRowHeight,
  ValueListEditor_SetDefaultRowHeight,
  ValueListEditor_GetDoubleBuffered,
  ValueListEditor_SetDoubleBuffered,
  ValueListEditor_GetDragCursor,
  ValueListEditor_SetDragCursor,
  ValueListEditor_GetDragKind,
  ValueListEditor_SetDragKind,
  ValueListEditor_GetDragMode,
  ValueListEditor_SetDragMode,
  ValueListEditor_GetDrawingStyle,
  ValueListEditor_SetDrawingStyle,
  ValueListEditor_GetEnabled,
  ValueListEditor_SetEnabled,
  ValueListEditor_GetFixedColor,
  ValueListEditor_SetFixedColor,
  ValueListEditor_GetFixedCols,
  ValueListEditor_SetFixedCols,
  ValueListEditor_GetFont,
  ValueListEditor_SetFont,
  ValueListEditor_GetGradientEndColor,
  ValueListEditor_SetGradientEndColor,
  ValueListEditor_GetGradientStartColor,
  ValueListEditor_SetGradientStartColor,
  ValueListEditor_GetGridLineWidth,
  ValueListEditor_SetGridLineWidth,
  ValueListEditor_GetOptions,
  ValueListEditor_SetOptions,
  ValueListEditor_GetParentColor,
  ValueListEditor_SetParentColor,
  ValueListEditor_GetParentCtl3D,
  ValueListEditor_SetParentCtl3D,
  ValueListEditor_GetParentDoubleBuffered,
  ValueListEditor_SetParentDoubleBuffered,
  ValueListEditor_GetParentFont,
  ValueListEditor_SetParentFont,
  ValueListEditor_GetParentShowHint,
  ValueListEditor_SetParentShowHint,
  ValueListEditor_GetPopupMenu,
  ValueListEditor_SetPopupMenu,
  ValueListEditor_GetScrollBars,
  ValueListEditor_SetScrollBars,
  ValueListEditor_GetShowHint,
  ValueListEditor_SetShowHint,
  ValueListEditor_GetStrings,
  ValueListEditor_SetStrings,
  ValueListEditor_GetTabOrder,
  ValueListEditor_SetTabOrder,
  ValueListEditor_GetVisible,
  ValueListEditor_SetVisible,
  ValueListEditor_GetStyleElements,
  ValueListEditor_SetStyleElements,
  ValueListEditor_SetOnClick,
  ValueListEditor_SetOnContextPopup,
  ValueListEditor_SetOnDblClick,
  ValueListEditor_SetOnDragDrop,
  ValueListEditor_SetOnDragOver,
  ValueListEditor_SetOnDrawCell,
  ValueListEditor_SetOnEndDock,
  ValueListEditor_SetOnEndDrag,
  ValueListEditor_SetOnEnter,
  ValueListEditor_SetOnExit,
  ValueListEditor_SetOnGesture,
  ValueListEditor_SetOnGetEditMask,
  ValueListEditor_SetOnGetEditText,
  ValueListEditor_SetOnKeyDown,
  ValueListEditor_SetOnKeyPress,
  ValueListEditor_SetOnKeyUp,
  ValueListEditor_SetOnMouseActivate,
  ValueListEditor_SetOnMouseDown,
  ValueListEditor_SetOnMouseEnter,
  ValueListEditor_SetOnMouseLeave,
  ValueListEditor_SetOnMouseMove,
  ValueListEditor_SetOnMouseUp,
  ValueListEditor_SetOnMouseWheelDown,
  ValueListEditor_SetOnMouseWheelUp,
  ValueListEditor_SetOnRowMoved,
  ValueListEditor_SetOnSelectCell,
  ValueListEditor_SetOnSetEditText,
  ValueListEditor_SetOnStartDock,
  ValueListEditor_SetOnTopLeftChanged,
  ValueListEditor_GetCanvas,
  ValueListEditor_GetCol,
  ValueListEditor_SetCol,
  ValueListEditor_GetEditorMode,
  ValueListEditor_SetEditorMode,
  ValueListEditor_GetGridHeight,
  ValueListEditor_GetGridWidth,
  ValueListEditor_GetLeftCol,
  ValueListEditor_SetLeftCol,
  ValueListEditor_GetSelection,
  ValueListEditor_SetSelection,
  ValueListEditor_GetRow,
  ValueListEditor_SetRow,
  ValueListEditor_GetTopRow,
  ValueListEditor_SetTopRow,
  ValueListEditor_GetTabStop,
  ValueListEditor_SetTabStop,
  ValueListEditor_GetDockClientCount,
  ValueListEditor_GetDockSite,
  ValueListEditor_SetDockSite,
  ValueListEditor_GetAlignDisabled,
  ValueListEditor_GetMouseInClient,
  ValueListEditor_GetVisibleDockClientCount,
  ValueListEditor_GetBrush,
  ValueListEditor_GetControlCount,
  ValueListEditor_GetHandle,
  ValueListEditor_GetParentWindow,
  ValueListEditor_SetParentWindow,
  ValueListEditor_GetUseDockManager,
  ValueListEditor_SetUseDockManager,
  ValueListEditor_GetAction,
  ValueListEditor_SetAction,
  ValueListEditor_GetBoundsRect,
  ValueListEditor_SetBoundsRect,
  ValueListEditor_GetClientHeight,
  ValueListEditor_SetClientHeight,
  ValueListEditor_GetClientOrigin,
  ValueListEditor_GetClientRect,
  ValueListEditor_GetClientWidth,
  ValueListEditor_SetClientWidth,
  ValueListEditor_GetControlState,
  ValueListEditor_SetControlState,
  ValueListEditor_GetControlStyle,
  ValueListEditor_SetControlStyle,
  ValueListEditor_GetExplicitLeft,
  ValueListEditor_GetExplicitTop,
  ValueListEditor_GetExplicitWidth,
  ValueListEditor_GetExplicitHeight,
  ValueListEditor_GetFloating,
  ValueListEditor_GetParent,
  ValueListEditor_SetParent,
  ValueListEditor_GetAlignWithMargins,
  ValueListEditor_SetAlignWithMargins,
  ValueListEditor_GetLeft,
  ValueListEditor_SetLeft,
  ValueListEditor_GetTop,
  ValueListEditor_SetTop,
  ValueListEditor_GetWidth,
  ValueListEditor_SetWidth,
  ValueListEditor_GetHeight,
  ValueListEditor_SetHeight,
  ValueListEditor_GetCursor,
  ValueListEditor_SetCursor,
  ValueListEditor_GetHint,
  ValueListEditor_SetHint,
  ValueListEditor_GetMargins,
  ValueListEditor_SetMargins,
  ValueListEditor_GetCustomHint,
  ValueListEditor_SetCustomHint,
  ValueListEditor_GetComponentCount,
  ValueListEditor_GetComponentIndex,
  ValueListEditor_SetComponentIndex,
  ValueListEditor_GetOwner,
  ValueListEditor_GetName,
  ValueListEditor_SetName,
  ValueListEditor_GetTag,
  ValueListEditor_SetTag,
  ValueListEditor_GetCells,
  ValueListEditor_SetCells,
  ValueListEditor_GetValues,
  ValueListEditor_SetValues,
  ValueListEditor_GetColWidths,
  ValueListEditor_SetColWidths,
  ValueListEditor_GetRowHeights,
  ValueListEditor_SetRowHeights,
  ValueListEditor_GetTabStops,
  ValueListEditor_SetTabStops,
  ValueListEditor_GetDockClients,
  ValueListEditor_GetControls,
  ValueListEditor_GetComponents,
  ValueListEditor_StaticClassType;

