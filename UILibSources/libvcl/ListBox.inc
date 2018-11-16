
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ListBox_Create(AOwner: TComponent): TListBox; stdcall;
begin
  Result :=  TListBox.Create(AOwner);
end;

procedure ListBox_Free(AObj: TListBox); stdcall;
begin
  AObj.Free;
end;

procedure ListBox_AddItem(AObj: TListBox; Item: PChar; AObject: TObject); stdcall;
begin
  AObj.AddItem(Item, AObject);
end;

procedure ListBox_Clear(AObj: TListBox); stdcall;
begin
  AObj.Clear;
end;

procedure ListBox_ClearSelection(AObj: TListBox); stdcall;
begin
  AObj.ClearSelection;
end;

procedure ListBox_DeleteSelected(AObj: TListBox); stdcall;
begin
  AObj.DeleteSelected;
end;

procedure ListBox_SelectAll(AObj: TListBox); stdcall;
begin
  AObj.SelectAll;
end;

function ListBox_CanFocus(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function ListBox_ContainsControl(AObj: TListBox; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function ListBox_ControlAtPos(AObj: TListBox; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure ListBox_DisableAlign(AObj: TListBox); stdcall;
begin
  AObj.DisableAlign;
end;

procedure ListBox_EnableAlign(AObj: TListBox); stdcall;
begin
  AObj.EnableAlign;
end;

function ListBox_FindChildControl(AObj: TListBox; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure ListBox_FlipChildren(AObj: TListBox; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function ListBox_Focused(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function ListBox_HandleAllocated(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ListBox_InsertControl(AObj: TListBox; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure ListBox_Invalidate(AObj: TListBox); stdcall;
begin
  AObj.Invalidate;
end;

procedure ListBox_PaintTo(AObj: TListBox; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure ListBox_RemoveControl(AObj: TListBox; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure ListBox_Realign(AObj: TListBox); stdcall;
begin
  AObj.Realign;
end;

procedure ListBox_Repaint(AObj: TListBox); stdcall;
begin
  AObj.Repaint;
end;

procedure ListBox_ScaleBy(AObj: TListBox; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure ListBox_ScrollBy(AObj: TListBox; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure ListBox_SetBounds(AObj: TListBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ListBox_SetFocus(AObj: TListBox); stdcall;
begin
  AObj.SetFocus;
end;

procedure ListBox_Update(AObj: TListBox); stdcall;
begin
  AObj.Update;
end;

procedure ListBox_UpdateControlState(AObj: TListBox); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure ListBox_BringToFront(AObj: TListBox); stdcall;
begin
  AObj.BringToFront;
end;

procedure ListBox_ClientToScreen(AObj: TListBox; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure ListBox_ClientToParent(AObj: TListBox; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function ListBox_Dragging(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function ListBox_HasParent(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ListBox_Hide(AObj: TListBox); stdcall;
begin
  AObj.Hide;
end;

function ListBox_Perform(AObj: TListBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ListBox_Refresh(AObj: TListBox); stdcall;
begin
  AObj.Refresh;
end;

procedure ListBox_ScreenToClient(AObj: TListBox; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure ListBox_ParentToClient(AObj: TListBox; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure ListBox_SendToBack(AObj: TListBox); stdcall;
begin
  AObj.SendToBack;
end;

procedure ListBox_Show(AObj: TListBox); stdcall;
begin
  AObj.Show;
end;

function ListBox_GetTextBuf(AObj: TListBox; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ListBox_GetTextLen(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure ListBox_SetTextBuf(AObj: TListBox; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function ListBox_FindComponent(AObj: TListBox; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ListBox_GetNamePath(AObj: TListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ListBox_Assign(AObj: TListBox; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ListBox_DisposeOf(AObj: TListBox); stdcall;
begin
  AObj.DisposeOf;
end;

function ListBox_ClassType(AObj: TListBox): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function ListBox_ClassName(AObj: TListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ListBox_InstanceSize(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function ListBox_InheritsFrom(AObj: TListBox; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function ListBox_Equals(AObj: TListBox; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListBox_GetHashCode(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ListBox_ToString(AObj: TListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListBox_GetStyle(AObj: TListBox): TListBoxStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure ListBox_SetStyle(AObj: TListBox; AValue: TListBoxStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function ListBox_GetAutoComplete(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.AutoComplete;
end;

procedure ListBox_SetAutoComplete(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.AutoComplete := AValue;
end;

function ListBox_GetAutoCompleteDelay(AObj: TListBox): Cardinal; stdcall;
begin
  Result :=  AObj.AutoCompleteDelay;
end;

procedure ListBox_SetAutoCompleteDelay(AObj: TListBox; AValue: Cardinal); stdcall;
begin
  AObj.AutoCompleteDelay := AValue;
end;

function ListBox_GetAlign(AObj: TListBox): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ListBox_SetAlign(AObj: TListBox; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ListBox_GetAnchors(AObj: TListBox): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ListBox_SetAnchors(AObj: TListBox; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ListBox_GetBevelEdges(AObj: TListBox): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure ListBox_SetBevelEdges(AObj: TListBox; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function ListBox_GetBevelInner(AObj: TListBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure ListBox_SetBevelInner(AObj: TListBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function ListBox_GetBevelKind(AObj: TListBox): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure ListBox_SetBevelKind(AObj: TListBox; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function ListBox_GetBevelOuter(AObj: TListBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure ListBox_SetBevelOuter(AObj: TListBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function ListBox_GetBiDiMode(AObj: TListBox): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ListBox_SetBiDiMode(AObj: TListBox; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function ListBox_GetBorderStyle(AObj: TListBox): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure ListBox_SetBorderStyle(AObj: TListBox; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function ListBox_GetColor(AObj: TListBox): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure ListBox_SetColor(AObj: TListBox; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function ListBox_GetColumns(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.Columns;
end;

procedure ListBox_SetColumns(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.Columns := AValue;
end;

function ListBox_GetDoubleBuffered(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ListBox_SetDoubleBuffered(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function ListBox_GetDragCursor(AObj: TListBox): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure ListBox_SetDragCursor(AObj: TListBox; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function ListBox_GetDragKind(AObj: TListBox): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure ListBox_SetDragKind(AObj: TListBox; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function ListBox_GetDragMode(AObj: TListBox): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure ListBox_SetDragMode(AObj: TListBox; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function ListBox_GetEnabled(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ListBox_SetEnabled(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ListBox_GetFont(AObj: TListBox): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure ListBox_SetFont(AObj: TListBox; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function ListBox_GetItemHeight(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ItemHeight;
end;

procedure ListBox_SetItemHeight(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.ItemHeight := AValue;
end;

function ListBox_GetItems(AObj: TListBox): TStrings; stdcall;
begin
  Result :=  AObj.Items;
end;

procedure ListBox_SetItems(AObj: TListBox; AValue: TStrings); stdcall;
begin
  AObj.Items := AValue;
end;

function ListBox_GetMultiSelect(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.MultiSelect;
end;

procedure ListBox_SetMultiSelect(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.MultiSelect := AValue;
end;

function ListBox_GetParentColor(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure ListBox_SetParentColor(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function ListBox_GetParentCtl3D(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure ListBox_SetParentCtl3D(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function ListBox_GetParentDoubleBuffered(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure ListBox_SetParentDoubleBuffered(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function ListBox_GetParentFont(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure ListBox_SetParentFont(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function ListBox_GetParentShowHint(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ListBox_SetParentShowHint(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ListBox_GetPopupMenu(AObj: TListBox): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ListBox_SetPopupMenu(AObj: TListBox; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ListBox_GetShowHint(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ListBox_SetShowHint(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ListBox_GetSorted(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.Sorted;
end;

procedure ListBox_SetSorted(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.Sorted := AValue;
end;

function ListBox_GetTabOrder(AObj: TListBox): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure ListBox_SetTabOrder(AObj: TListBox; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function ListBox_GetTabStop(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure ListBox_SetTabStop(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function ListBox_GetTabWidth(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.TabWidth;
end;

procedure ListBox_SetTabWidth(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.TabWidth := AValue;
end;

function ListBox_GetVisible(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ListBox_SetVisible(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function ListBox_GetStyleElements(AObj: TListBox): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ListBox_SetStyleElements(AObj: TListBox; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure ListBox_SetOnClick(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnContextPopup(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnData(AObj: TListBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnData := nil;
    TEventClass.Remove(AObj, @TEventClass.ListBoxOnData);
    Exit;
  end;
  AObj.OnData := TEventClass.ListBoxOnData;
  TEventClass.Add(AObj, @TEventClass.ListBoxOnData, AEventId);
end;

procedure ListBox_SetOnDataFind(AObj: TListBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDataFind := nil;
    TEventClass.Remove(AObj, @TEventClass.ListBoxOnDataFind);
    Exit;
  end;
  AObj.OnDataFind := TEventClass.ListBoxOnDataFind;
  TEventClass.Add(AObj, @TEventClass.ListBoxOnDataFind, AEventId);
end;

procedure ListBox_SetOnDataObject(AObj: TListBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDataObject := nil;
    TEventClass.Remove(AObj, @TEventClass.ListBoxOnDataObject);
    Exit;
  end;
  AObj.OnDataObject := TEventClass.ListBoxOnDataObject;
  TEventClass.Add(AObj, @TEventClass.ListBoxOnDataObject, AEventId);
end;

procedure ListBox_SetOnDblClick(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnDragDrop(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnDragOver(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnDrawItem(AObj: TListBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDrawItem := nil;
    TEventClass.Remove(AObj, @TEventClass.ListBoxOnDrawItem);
    Exit;
  end;
  AObj.OnDrawItem := TEventClass.ListBoxOnDrawItem;
  TEventClass.Add(AObj, @TEventClass.ListBoxOnDrawItem, AEventId);
end;

procedure ListBox_SetOnEndDock(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnEndDrag(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnEnter(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnExit(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnGesture(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnKeyDown(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnKeyPress(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnKeyUp(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMeasureItem(AObj: TListBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMeasureItem := nil;
    TEventClass.Remove(AObj, @TEventClass.ListBoxOnMeasureItem);
    Exit;
  end;
  AObj.OnMeasureItem := TEventClass.ListBoxOnMeasureItem;
  TEventClass.Add(AObj, @TEventClass.ListBoxOnMeasureItem, AEventId);
end;

procedure ListBox_SetOnMouseActivate(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMouseDown(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMouseEnter(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMouseLeave(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMouseMove(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMouseUp(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnStartDock(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

function ListBox_GetCanvas(AObj: TListBox): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function ListBox_GetCount(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

procedure ListBox_SetCount(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.Count := AValue;
end;

function ListBox_GetSelCount(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.SelCount;
end;

function ListBox_GetItemIndex(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ItemIndex;
end;

procedure ListBox_SetItemIndex(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.ItemIndex := AValue;
end;

function ListBox_GetDockClientCount(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function ListBox_GetDockSite(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure ListBox_SetDockSite(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function ListBox_GetAlignDisabled(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function ListBox_GetMouseInClient(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function ListBox_GetVisibleDockClientCount(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function ListBox_GetBrush(AObj: TListBox): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function ListBox_GetControlCount(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function ListBox_GetHandle(AObj: TListBox): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ListBox_GetParentWindow(AObj: TListBox): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ListBox_SetParentWindow(AObj: TListBox; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function ListBox_GetUseDockManager(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure ListBox_SetUseDockManager(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function ListBox_GetAction(AObj: TListBox): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ListBox_SetAction(AObj: TListBox; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure ListBox_GetBoundsRect(AObj: TListBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ListBox_SetBoundsRect(AObj: TListBox; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ListBox_GetClientHeight(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ListBox_SetClientHeight(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ListBox_GetClientOrigin(AObj: TListBox; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure ListBox_GetClientRect(AObj: TListBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ListBox_GetClientWidth(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ListBox_SetClientWidth(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ListBox_GetControlState(AObj: TListBox): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure ListBox_SetControlState(AObj: TListBox; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function ListBox_GetControlStyle(AObj: TListBox): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure ListBox_SetControlStyle(AObj: TListBox; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function ListBox_GetExplicitLeft(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ListBox_GetExplicitTop(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ListBox_GetExplicitWidth(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ListBox_GetExplicitHeight(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ListBox_GetFloating(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function ListBox_GetParent(AObj: TListBox): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ListBox_SetParent(AObj: TListBox; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ListBox_GetAlignWithMargins(AObj: TListBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ListBox_SetAlignWithMargins(AObj: TListBox; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ListBox_GetLeft(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ListBox_SetLeft(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ListBox_GetTop(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ListBox_SetTop(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ListBox_GetWidth(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ListBox_SetWidth(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ListBox_GetHeight(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ListBox_SetHeight(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ListBox_GetCursor(AObj: TListBox): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ListBox_SetCursor(AObj: TListBox; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ListBox_GetHint(AObj: TListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ListBox_SetHint(AObj: TListBox; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ListBox_GetMargins(AObj: TListBox): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ListBox_SetMargins(AObj: TListBox; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ListBox_GetCustomHint(AObj: TListBox): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ListBox_SetCustomHint(AObj: TListBox; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ListBox_GetComponentCount(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ListBox_GetComponentIndex(AObj: TListBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ListBox_SetComponentIndex(AObj: TListBox; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ListBox_GetOwner(AObj: TListBox): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ListBox_GetName(AObj: TListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ListBox_SetName(AObj: TListBox; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ListBox_GetTag(AObj: TListBox): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ListBox_SetTag(AObj: TListBox; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ListBox_GetSelected(AObj: TListBox; Index: Integer): LongBool; stdcall;
begin
  Result :=  AObj.Selected[Index];
end;

procedure ListBox_SetSelected(AObj: TListBox; Index: Integer; AValue: LongBool); stdcall;
begin
  AObj.Selected[Index] := AValue;
end;

function ListBox_GetDockClients(AObj: TListBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function ListBox_GetControls(AObj: TListBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ListBox_GetComponents(AObj: TListBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function ListBox_StaticClassType: TClass; stdcall;
begin
  Result := TListBox;
end;

exports
  ListBox_Create,
  ListBox_Free,
  ListBox_AddItem,
  ListBox_Clear,
  ListBox_ClearSelection,
  ListBox_DeleteSelected,
  ListBox_SelectAll,
  ListBox_CanFocus,
  ListBox_ContainsControl,
  ListBox_ControlAtPos,
  ListBox_DisableAlign,
  ListBox_EnableAlign,
  ListBox_FindChildControl,
  ListBox_FlipChildren,
  ListBox_Focused,
  ListBox_HandleAllocated,
  ListBox_InsertControl,
  ListBox_Invalidate,
  ListBox_PaintTo,
  ListBox_RemoveControl,
  ListBox_Realign,
  ListBox_Repaint,
  ListBox_ScaleBy,
  ListBox_ScrollBy,
  ListBox_SetBounds,
  ListBox_SetFocus,
  ListBox_Update,
  ListBox_UpdateControlState,
  ListBox_BringToFront,
  ListBox_ClientToScreen,
  ListBox_ClientToParent,
  ListBox_Dragging,
  ListBox_HasParent,
  ListBox_Hide,
  ListBox_Perform,
  ListBox_Refresh,
  ListBox_ScreenToClient,
  ListBox_ParentToClient,
  ListBox_SendToBack,
  ListBox_Show,
  ListBox_GetTextBuf,
  ListBox_GetTextLen,
  ListBox_SetTextBuf,
  ListBox_FindComponent,
  ListBox_GetNamePath,
  ListBox_Assign,
  ListBox_DisposeOf,
  ListBox_ClassType,
  ListBox_ClassName,
  ListBox_InstanceSize,
  ListBox_InheritsFrom,
  ListBox_Equals,
  ListBox_GetHashCode,
  ListBox_ToString,
  ListBox_GetStyle,
  ListBox_SetStyle,
  ListBox_GetAutoComplete,
  ListBox_SetAutoComplete,
  ListBox_GetAutoCompleteDelay,
  ListBox_SetAutoCompleteDelay,
  ListBox_GetAlign,
  ListBox_SetAlign,
  ListBox_GetAnchors,
  ListBox_SetAnchors,
  ListBox_GetBevelEdges,
  ListBox_SetBevelEdges,
  ListBox_GetBevelInner,
  ListBox_SetBevelInner,
  ListBox_GetBevelKind,
  ListBox_SetBevelKind,
  ListBox_GetBevelOuter,
  ListBox_SetBevelOuter,
  ListBox_GetBiDiMode,
  ListBox_SetBiDiMode,
  ListBox_GetBorderStyle,
  ListBox_SetBorderStyle,
  ListBox_GetColor,
  ListBox_SetColor,
  ListBox_GetColumns,
  ListBox_SetColumns,
  ListBox_GetDoubleBuffered,
  ListBox_SetDoubleBuffered,
  ListBox_GetDragCursor,
  ListBox_SetDragCursor,
  ListBox_GetDragKind,
  ListBox_SetDragKind,
  ListBox_GetDragMode,
  ListBox_SetDragMode,
  ListBox_GetEnabled,
  ListBox_SetEnabled,
  ListBox_GetFont,
  ListBox_SetFont,
  ListBox_GetItemHeight,
  ListBox_SetItemHeight,
  ListBox_GetItems,
  ListBox_SetItems,
  ListBox_GetMultiSelect,
  ListBox_SetMultiSelect,
  ListBox_GetParentColor,
  ListBox_SetParentColor,
  ListBox_GetParentCtl3D,
  ListBox_SetParentCtl3D,
  ListBox_GetParentDoubleBuffered,
  ListBox_SetParentDoubleBuffered,
  ListBox_GetParentFont,
  ListBox_SetParentFont,
  ListBox_GetParentShowHint,
  ListBox_SetParentShowHint,
  ListBox_GetPopupMenu,
  ListBox_SetPopupMenu,
  ListBox_GetShowHint,
  ListBox_SetShowHint,
  ListBox_GetSorted,
  ListBox_SetSorted,
  ListBox_GetTabOrder,
  ListBox_SetTabOrder,
  ListBox_GetTabStop,
  ListBox_SetTabStop,
  ListBox_GetTabWidth,
  ListBox_SetTabWidth,
  ListBox_GetVisible,
  ListBox_SetVisible,
  ListBox_GetStyleElements,
  ListBox_SetStyleElements,
  ListBox_SetOnClick,
  ListBox_SetOnContextPopup,
  ListBox_SetOnData,
  ListBox_SetOnDataFind,
  ListBox_SetOnDataObject,
  ListBox_SetOnDblClick,
  ListBox_SetOnDragDrop,
  ListBox_SetOnDragOver,
  ListBox_SetOnDrawItem,
  ListBox_SetOnEndDock,
  ListBox_SetOnEndDrag,
  ListBox_SetOnEnter,
  ListBox_SetOnExit,
  ListBox_SetOnGesture,
  ListBox_SetOnKeyDown,
  ListBox_SetOnKeyPress,
  ListBox_SetOnKeyUp,
  ListBox_SetOnMeasureItem,
  ListBox_SetOnMouseActivate,
  ListBox_SetOnMouseDown,
  ListBox_SetOnMouseEnter,
  ListBox_SetOnMouseLeave,
  ListBox_SetOnMouseMove,
  ListBox_SetOnMouseUp,
  ListBox_SetOnStartDock,
  ListBox_GetCanvas,
  ListBox_GetCount,
  ListBox_SetCount,
  ListBox_GetSelCount,
  ListBox_GetItemIndex,
  ListBox_SetItemIndex,
  ListBox_GetDockClientCount,
  ListBox_GetDockSite,
  ListBox_SetDockSite,
  ListBox_GetAlignDisabled,
  ListBox_GetMouseInClient,
  ListBox_GetVisibleDockClientCount,
  ListBox_GetBrush,
  ListBox_GetControlCount,
  ListBox_GetHandle,
  ListBox_GetParentWindow,
  ListBox_SetParentWindow,
  ListBox_GetUseDockManager,
  ListBox_SetUseDockManager,
  ListBox_GetAction,
  ListBox_SetAction,
  ListBox_GetBoundsRect,
  ListBox_SetBoundsRect,
  ListBox_GetClientHeight,
  ListBox_SetClientHeight,
  ListBox_GetClientOrigin,
  ListBox_GetClientRect,
  ListBox_GetClientWidth,
  ListBox_SetClientWidth,
  ListBox_GetControlState,
  ListBox_SetControlState,
  ListBox_GetControlStyle,
  ListBox_SetControlStyle,
  ListBox_GetExplicitLeft,
  ListBox_GetExplicitTop,
  ListBox_GetExplicitWidth,
  ListBox_GetExplicitHeight,
  ListBox_GetFloating,
  ListBox_GetParent,
  ListBox_SetParent,
  ListBox_GetAlignWithMargins,
  ListBox_SetAlignWithMargins,
  ListBox_GetLeft,
  ListBox_SetLeft,
  ListBox_GetTop,
  ListBox_SetTop,
  ListBox_GetWidth,
  ListBox_SetWidth,
  ListBox_GetHeight,
  ListBox_SetHeight,
  ListBox_GetCursor,
  ListBox_SetCursor,
  ListBox_GetHint,
  ListBox_SetHint,
  ListBox_GetMargins,
  ListBox_SetMargins,
  ListBox_GetCustomHint,
  ListBox_SetCustomHint,
  ListBox_GetComponentCount,
  ListBox_GetComponentIndex,
  ListBox_SetComponentIndex,
  ListBox_GetOwner,
  ListBox_GetName,
  ListBox_SetName,
  ListBox_GetTag,
  ListBox_SetTag,
  ListBox_GetSelected,
  ListBox_SetSelected,
  ListBox_GetDockClients,
  ListBox_GetControls,
  ListBox_GetComponents,
  ListBox_StaticClassType;

