
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ListView_Create(AOwner: TComponent): TListView; extdecl;
begin
  Result :=  TListView.Create(AOwner);
  Result.ScrollBars := ssAutoBoth;
end;

procedure ListView_Free(AObj: TListView); extdecl;
begin
  AObj.Free;
end;

procedure ListView_Clear(AObj: TListView); extdecl;
begin
  AObj.Clear;
end;

procedure ListView_ClearSelection(AObj: TListView); extdecl;
begin
  AObj.ClearSelection;
end;

function ListView_IsEditing(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.IsEditing;
end;

procedure ListView_SelectAll(AObj: TListView); extdecl;
begin
  AObj.SelectAll;
end;

function ListView_CanFocus(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure ListView_FlipChildren(AObj: TListView; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function ListView_Focused(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure ListView_Invalidate(AObj: TListView); extdecl;
begin
  AObj.Invalidate;
end;

procedure ListView_Realign(AObj: TListView); extdecl;
begin
  AObj.Realign;
end;

procedure ListView_Repaint(AObj: TListView); extdecl;
begin
  AObj.Repaint;
end;

procedure ListView_ScaleBy(AObj: TListView; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure ListView_SetBounds(AObj: TListView; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ListView_SetFocus(AObj: TListView); extdecl;
begin
  AObj.SetFocus;
end;

procedure ListView_Update(AObj: TListView); extdecl;
begin
  AObj.Update;
end;

procedure ListView_BringToFront(AObj: TListView); extdecl;
begin
  AObj.BringToFront;
end;

function ListView_HasParent(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ListView_Hide(AObj: TListView); extdecl;
begin
  AObj.Hide;
end;

function ListView_Perform(AObj: TListView; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ListView_Refresh(AObj: TListView); extdecl;
begin
  AObj.Refresh;
end;

procedure ListView_SendToBack(AObj: TListView); extdecl;
begin
  AObj.SendToBack;
end;

procedure ListView_Show(AObj: TListView); extdecl;
begin
  AObj.Show;
end;

function ListView_GetTextBuf(AObj: TListView; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ListView_GetTextLen(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function ListView_FindComponent(AObj: TListView; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ListView_GetNamePath(AObj: TListView): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ListView_Assign(AObj: TListView; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ListView_ClassName(AObj: TListView): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ListView_Equals(AObj: TListView; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListView_GetHashCode(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ListView_ToString(AObj: TListView): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListView_GetAction(AObj: TListView): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure ListView_SetAction(AObj: TListView; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function ListView_GetAlign(AObj: TListView): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure ListView_SetAlign(AObj: TListView; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function ListView_GetAllocBy(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.AllocBy;
end;

procedure ListView_SetAllocBy(AObj: TListView; AValue: Integer); extdecl;
begin
  AObj.AllocBy := AValue;
end;

function ListView_GetAnchors(AObj: TListView): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure ListView_SetAnchors(AObj: TListView; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function ListView_GetBiDiMode(AObj: TListView): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ListView_SetBiDiMode(AObj: TListView; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function ListView_GetBorderStyle(AObj: TListView): TBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure ListView_SetBorderStyle(AObj: TListView; AValue: TBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function ListView_GetBorderWidth(AObj: TListView): TBorderWidth; extdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure ListView_SetBorderWidth(AObj: TListView; AValue: TBorderWidth); extdecl;
begin
  AObj.BorderWidth := AValue;
end;

function ListView_GetCheckboxes(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.Checkboxes;
end;

procedure ListView_SetCheckboxes(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.Checkboxes := AValue;
end;

function ListView_GetColor(AObj: TListView): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure ListView_SetColor(AObj: TListView; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function ListView_GetColumns(AObj: TListView): TListColumns; extdecl;
begin
  Result :=  AObj.Columns;
end;

procedure ListView_SetColumns(AObj: TListView; AValue: TListColumns); extdecl;
begin
  AObj.Columns := AValue;
end;

function ListView_GetColumnClick(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.ColumnClick;
end;

procedure ListView_SetColumnClick(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.ColumnClick := AValue;
end;

function ListView_GetDoubleBuffered(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ListView_SetDoubleBuffered(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function ListView_GetEnabled(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure ListView_SetEnabled(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function ListView_GetFont(AObj: TListView): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure ListView_SetFont(AObj: TListView; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function ListView_GetFlatScrollBars(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.FlatScrollBars;
end;

procedure ListView_SetFlatScrollBars(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.FlatScrollBars := AValue;
end;

function ListView_GetFullDrag(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.FullDrag;
end;

procedure ListView_SetFullDrag(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.FullDrag := AValue;
end;

function ListView_GetGridLines(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.GridLines;
end;

procedure ListView_SetGridLines(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.GridLines := AValue;
end;

function ListView_GetHideSelection(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.HideSelection;
end;

procedure ListView_SetHideSelection(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.HideSelection := AValue;
end;

function ListView_GetHotTrack(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.HotTrack;
end;

procedure ListView_SetHotTrack(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.HotTrack := AValue;
end;

function ListView_GetIconOptions(AObj: TListView): TIconOptions; extdecl;
begin
  Result :=  AObj.IconOptions;
end;

procedure ListView_SetIconOptions(AObj: TListView; AValue: TIconOptions); extdecl;
begin
  AObj.IconOptions := AValue;
end;

function ListView_GetItems(AObj: TListView): TListItems; extdecl;
begin
  Result :=  AObj.Items;
end;

procedure ListView_SetItems(AObj: TListView; AValue: TListItems); extdecl;
begin
  AObj.Items := AValue;
end;

function ListView_GetLargeImages(AObj: TListView): TCustomImageList; extdecl;
begin
  Result :=  AObj.LargeImages;
end;

procedure ListView_SetLargeImages(AObj: TListView; AValue: TCustomImageList); extdecl;
begin
  AObj.LargeImages := AValue;
end;

function ListView_GetMultiSelect(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.MultiSelect;
end;

procedure ListView_SetMultiSelect(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.MultiSelect := AValue;
end;

function ListView_GetReadOnly(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.ReadOnly;
end;

procedure ListView_SetReadOnly(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.ReadOnly := AValue;
end;

function ListView_GetRowSelect(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.RowSelect;
end;

procedure ListView_SetRowSelect(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.RowSelect := AValue;
end;

function ListView_GetParentColor(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure ListView_SetParentColor(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function ListView_GetParentFont(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure ListView_SetParentFont(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function ListView_GetParentShowHint(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ListView_SetParentShowHint(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function ListView_GetPopupMenu(AObj: TListView): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ListView_SetPopupMenu(AObj: TListView; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function ListView_GetShowColumnHeaders(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.ShowColumnHeaders;
end;

procedure ListView_SetShowColumnHeaders(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.ShowColumnHeaders := AValue;
end;

function ListView_GetShowHint(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure ListView_SetShowHint(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function ListView_GetSmallImages(AObj: TListView): TCustomImageList; extdecl;
begin
  Result :=  AObj.SmallImages;
end;

procedure ListView_SetSmallImages(AObj: TListView; AValue: TCustomImageList); extdecl;
begin
  AObj.SmallImages := AValue;
end;

function ListView_GetSortType(AObj: TListView): TSortType; extdecl;
begin
  Result :=  AObj.SortType;
end;

procedure ListView_SetSortType(AObj: TListView; AValue: TSortType); extdecl;
begin
  AObj.SortType := AValue;
end;

function ListView_GetStateImages(AObj: TListView): TCustomImageList; extdecl;
begin
  Result :=  AObj.StateImages;
end;

procedure ListView_SetStateImages(AObj: TListView; AValue: TCustomImageList); extdecl;
begin
  AObj.StateImages := AValue;
end;

function ListView_GetTabOrder(AObj: TListView): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure ListView_SetTabOrder(AObj: TListView; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function ListView_GetTabStop(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure ListView_SetTabStop(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function ListView_GetViewStyle(AObj: TListView): TViewStyle; extdecl;
begin
  Result :=  AObj.ViewStyle;
end;

procedure ListView_SetViewStyle(AObj: TListView; AValue: TViewStyle); extdecl;
begin
  AObj.ViewStyle := AValue;
end;

function ListView_GetVisible(AObj: TListView): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure ListView_SetVisible(AObj: TListView; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure ListView_SetOnAdvancedCustomDraw(AObj: TListView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDraw := nil;
    EventClass.Remove(AObj, geListViewAdvancedCustomDraw);
    Exit;
  end;
  AObj.OnAdvancedCustomDraw := @EventClass.ListViewOnAdvancedCustomDraw;
  EventClass.Add(AObj, geListViewAdvancedCustomDraw, AEventId);
end;

procedure ListView_SetOnAdvancedCustomDrawItem(AObj: TListView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDrawItem := nil;
    EventClass.Remove(AObj, geListViewAdvancedCustomDrawItem);
    Exit;
  end;
  AObj.OnAdvancedCustomDrawItem := @EventClass.ListViewOnAdvancedCustomDrawItem;
  EventClass.Add(AObj, geListViewAdvancedCustomDrawItem, AEventId);
end;

procedure ListView_SetOnAdvancedCustomDrawSubItem(AObj: TListView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDrawSubItem := nil;
    EventClass.Remove(AObj, geListViewAdvancedCustomDrawSubItem);
    Exit;
  end;
  AObj.OnAdvancedCustomDrawSubItem := @EventClass.ListViewOnAdvancedCustomDrawSubItem;
  EventClass.Add(AObj, geListViewAdvancedCustomDrawSubItem, AEventId);
end;

procedure ListView_SetOnChange(AObj: TListView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    EventClass.Remove(AObj, geListViewChange);
    Exit;
  end;
  AObj.OnChange := @EventClass.ListViewOnChange;
  EventClass.Add(AObj, geListViewChange, AEventId);
end;

procedure ListView_SetOnClick(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnColumnClick(AObj: TListView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnColumnClick := nil;
    EventClass.Remove(AObj, geListViewColumnClick);
    Exit;
  end;
  AObj.OnColumnClick := @EventClass.ListViewOnColumnClick;
  EventClass.Add(AObj, geListViewColumnClick, AEventId);
end;

procedure ListView_SetOnCompare(AObj: TListView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnCompare := nil;
    EventClass.Remove(AObj, geListViewCompare);
    Exit;
  end;
  AObj.OnCompare := @EventClass.ListViewOnCompare;
  EventClass.Add(AObj, geListViewCompare, AEventId);
end;

procedure ListView_SetOnDblClick(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnEnter(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnExit(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnKeyDown(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnKeyPress(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnKeyUp(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnMouseDown(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnMouseEnter(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnMouseLeave(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnMouseMove(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnMouseUp(AObj: TListView; AEventId: NativeUInt); extdecl;
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

procedure ListView_SetOnResize(AObj: TListView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnResize := nil;
    EventClass.Remove(AObj, geResize);
    Exit;
  end;
  AObj.OnResize := @EventClass.OnResize;
  EventClass.Add(AObj, geResize, AEventId);
end;

procedure ListView_SetOnSelectItem(AObj: TListView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnSelectItem := nil;
    EventClass.Remove(AObj, geListViewSelectItem);
    Exit;
  end;
  AObj.OnSelectItem := @EventClass.ListViewOnSelectItem;
  EventClass.Add(AObj, geListViewSelectItem, AEventId);
end;

procedure ListView_SetOnItemChecked(AObj: TListView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnItemChecked := nil;
    EventClass.Remove(AObj, geListViewItemChecked);
    Exit;
  end;
  AObj.OnItemChecked := @EventClass.ListViewOnItemChecked;
  EventClass.Add(AObj, geListViewItemChecked, AEventId);
end;

function ListView_GetCanvas(AObj: TListView): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function ListView_GetItemFocused(AObj: TListView): TListItem; extdecl;
begin
  Result :=  AObj.ItemFocused;
end;

procedure ListView_SetItemFocused(AObj: TListView; AValue: TListItem); extdecl;
begin
  AObj.ItemFocused := AValue;
end;

function ListView_GetSelCount(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.SelCount;
end;

function ListView_GetSelected(AObj: TListView): TListItem; extdecl;
begin
  Result :=  AObj.Selected;
end;

procedure ListView_SetSelected(AObj: TListView; AValue: TListItem); extdecl;
begin
  AObj.Selected := AValue;
end;

function ListView_GetTopItem(AObj: TListView): TListItem; extdecl;
begin
  Result :=  AObj.TopItem;
end;

function ListView_GetVisibleRowCount(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.VisibleRowCount;
end;

function ListView_GetItemIndex(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.ItemIndex;
end;

procedure ListView_SetItemIndex(AObj: TListView; AValue: Integer); extdecl;
begin
  AObj.ItemIndex := AValue;
end;

function ListView_GetBrush(AObj: TListView): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function ListView_GetControlCount(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function ListView_GetHandle(AObj: TListView): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function ListView_GetParentWindow(AObj: TListView): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ListView_SetParentWindow(AObj: TListView; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

procedure ListView_GetBoundsRect(AObj: TListView; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ListView_SetBoundsRect(AObj: TListView; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function ListView_GetClientHeight(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ListView_SetClientHeight(AObj: TListView; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure ListView_GetClientRect(AObj: TListView; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function ListView_GetClientWidth(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ListView_SetClientWidth(AObj: TListView; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function ListView_GetParent(AObj: TListView): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure ListView_SetParent(AObj: TListView; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function ListView_GetLeft(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ListView_SetLeft(AObj: TListView; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ListView_GetTop(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ListView_SetTop(AObj: TListView; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ListView_GetWidth(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ListView_SetWidth(AObj: TListView; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function ListView_GetHeight(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ListView_SetHeight(AObj: TListView; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ListView_GetCursor(AObj: TListView): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure ListView_SetCursor(AObj: TListView; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function ListView_GetHint(AObj: TListView): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ListView_SetHint(AObj: TListView; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function ListView_GetComponentCount(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ListView_GetComponentIndex(AObj: TListView): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ListView_SetComponentIndex(AObj: TListView; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ListView_GetOwner(AObj: TListView): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ListView_GetName(AObj: TListView): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ListView_SetName(AObj: TListView; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ListView_GetTag(AObj: TListView): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ListView_SetTag(AObj: TListView; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ListView_GetColumn(AObj: TListView; Index: Integer): TListColumn; extdecl;
begin
  Result :=  AObj.Column[Index];
end;

function ListView_GetControls(AObj: TListView; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function ListView_GetComponents(AObj: TListView; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ListView_Create,
  ListView_Free,
  ListView_Clear,
  ListView_ClearSelection,
  ListView_IsEditing,
  ListView_SelectAll,
  ListView_CanFocus,
  ListView_FlipChildren,
  ListView_Focused,
  ListView_Invalidate,
  ListView_Realign,
  ListView_Repaint,
  ListView_ScaleBy,
  ListView_SetBounds,
  ListView_SetFocus,
  ListView_Update,
  ListView_BringToFront,
  ListView_HasParent,
  ListView_Hide,
  ListView_Perform,
  ListView_Refresh,
  ListView_SendToBack,
  ListView_Show,
  ListView_GetTextBuf,
  ListView_GetTextLen,
  ListView_FindComponent,
  ListView_GetNamePath,
  ListView_Assign,
  ListView_ClassName,
  ListView_Equals,
  ListView_GetHashCode,
  ListView_ToString,
  ListView_GetAction,
  ListView_SetAction,
  ListView_GetAlign,
  ListView_SetAlign,
  ListView_GetAllocBy,
  ListView_SetAllocBy,
  ListView_GetAnchors,
  ListView_SetAnchors,
  ListView_GetBiDiMode,
  ListView_SetBiDiMode,
  ListView_GetBorderStyle,
  ListView_SetBorderStyle,
  ListView_GetBorderWidth,
  ListView_SetBorderWidth,
  ListView_GetCheckboxes,
  ListView_SetCheckboxes,
  ListView_GetColor,
  ListView_SetColor,
  ListView_GetColumns,
  ListView_SetColumns,
  ListView_GetColumnClick,
  ListView_SetColumnClick,
  ListView_GetDoubleBuffered,
  ListView_SetDoubleBuffered,
  ListView_GetEnabled,
  ListView_SetEnabled,
  ListView_GetFont,
  ListView_SetFont,
  ListView_GetFlatScrollBars,
  ListView_SetFlatScrollBars,
  ListView_GetFullDrag,
  ListView_SetFullDrag,
  ListView_GetGridLines,
  ListView_SetGridLines,
  ListView_GetHideSelection,
  ListView_SetHideSelection,
  ListView_GetHotTrack,
  ListView_SetHotTrack,
  ListView_GetIconOptions,
  ListView_SetIconOptions,
  ListView_GetItems,
  ListView_SetItems,
  ListView_GetLargeImages,
  ListView_SetLargeImages,
  ListView_GetMultiSelect,
  ListView_SetMultiSelect,
  ListView_GetReadOnly,
  ListView_SetReadOnly,
  ListView_GetRowSelect,
  ListView_SetRowSelect,
  ListView_GetParentColor,
  ListView_SetParentColor,
  ListView_GetParentFont,
  ListView_SetParentFont,
  ListView_GetParentShowHint,
  ListView_SetParentShowHint,
  ListView_GetPopupMenu,
  ListView_SetPopupMenu,
  ListView_GetShowColumnHeaders,
  ListView_SetShowColumnHeaders,
  ListView_GetShowHint,
  ListView_SetShowHint,
  ListView_GetSmallImages,
  ListView_SetSmallImages,
  ListView_GetSortType,
  ListView_SetSortType,
  ListView_GetStateImages,
  ListView_SetStateImages,
  ListView_GetTabOrder,
  ListView_SetTabOrder,
  ListView_GetTabStop,
  ListView_SetTabStop,
  ListView_GetViewStyle,
  ListView_SetViewStyle,
  ListView_GetVisible,
  ListView_SetVisible,
  ListView_SetOnAdvancedCustomDraw,
  ListView_SetOnAdvancedCustomDrawItem,
  ListView_SetOnAdvancedCustomDrawSubItem,
  ListView_SetOnChange,
  ListView_SetOnClick,
  ListView_SetOnColumnClick,
  ListView_SetOnCompare,
  ListView_SetOnDblClick,
  ListView_SetOnEnter,
  ListView_SetOnExit,
  ListView_SetOnKeyDown,
  ListView_SetOnKeyPress,
  ListView_SetOnKeyUp,
  ListView_SetOnMouseDown,
  ListView_SetOnMouseEnter,
  ListView_SetOnMouseLeave,
  ListView_SetOnMouseMove,
  ListView_SetOnMouseUp,
  ListView_SetOnResize,
  ListView_SetOnSelectItem,
  ListView_SetOnItemChecked,
  ListView_GetCanvas,
  ListView_GetItemFocused,
  ListView_SetItemFocused,
  ListView_GetSelCount,
  ListView_GetSelected,
  ListView_SetSelected,
  ListView_GetTopItem,
  ListView_GetVisibleRowCount,
  ListView_GetItemIndex,
  ListView_SetItemIndex,
  ListView_GetBrush,
  ListView_GetControlCount,
  ListView_GetHandle,
  ListView_GetParentWindow,
  ListView_SetParentWindow,
  ListView_GetBoundsRect,
  ListView_SetBoundsRect,
  ListView_GetClientHeight,
  ListView_SetClientHeight,
  ListView_GetClientRect,
  ListView_GetClientWidth,
  ListView_SetClientWidth,
  ListView_GetParent,
  ListView_SetParent,
  ListView_GetLeft,
  ListView_SetLeft,
  ListView_GetTop,
  ListView_SetTop,
  ListView_GetWidth,
  ListView_SetWidth,
  ListView_GetHeight,
  ListView_SetHeight,
  ListView_GetCursor,
  ListView_SetCursor,
  ListView_GetHint,
  ListView_SetHint,
  ListView_GetComponentCount,
  ListView_GetComponentIndex,
  ListView_SetComponentIndex,
  ListView_GetOwner,
  ListView_GetName,
  ListView_SetName,
  ListView_GetTag,
  ListView_SetTag,
  ListView_GetColumn,
  ListView_GetControls,
  ListView_GetComponents;

