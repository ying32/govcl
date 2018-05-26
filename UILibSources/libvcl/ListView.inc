
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ListView_Create(AOwner: TComponent): TListView; stdcall;
begin
  Result :=  TListView.Create(AOwner);
end;

procedure ListView_Free(AObj: TListView); stdcall;
begin
  AObj.Free;
end;

procedure ListView_AddItem(AObj: TListView; Item: PChar; AObject: TObject); stdcall;
begin
  AObj.AddItem(Item, AObject);
end;

function ListView_AlphaSort(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.AlphaSort;
end;

procedure ListView_Clear(AObj: TListView); stdcall;
begin
  AObj.Clear;
end;

procedure ListView_ClearSelection(AObj: TListView); stdcall;
begin
  AObj.ClearSelection;
end;

procedure ListView_DeleteSelected(AObj: TListView); stdcall;
begin
  AObj.DeleteSelected;
end;

function ListView_GetSearchString(AObj: TListView): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetSearchString);
end;

function ListView_IsEditing(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.IsEditing;
end;

procedure ListView_SelectAll(AObj: TListView); stdcall;
begin
  AObj.SelectAll;
end;

procedure ListView_Scroll(AObj: TListView; DX: Integer; DY: Integer); stdcall;
begin
  AObj.Scroll(DX, DY);
end;

function ListView_CustomSort(AObj: TListView; SortProc: PFNLVCOMPARE; lParam: NativeInt): LongBool; stdcall;
begin
  Result :=  AObj.CustomSort(SortProc, lParam);
end;

function ListView_CanFocus(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure ListView_FlipChildren(AObj: TListView; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function ListView_Focused(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function ListView_HandleAllocated(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ListView_Invalidate(AObj: TListView); stdcall;
begin
  AObj.Invalidate;
end;

procedure ListView_Realign(AObj: TListView); stdcall;
begin
  AObj.Realign;
end;

procedure ListView_Repaint(AObj: TListView); stdcall;
begin
  AObj.Repaint;
end;

procedure ListView_ScaleBy(AObj: TListView; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure ListView_SetBounds(AObj: TListView; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ListView_SetFocus(AObj: TListView); stdcall;
begin
  AObj.SetFocus;
end;

procedure ListView_Update(AObj: TListView); stdcall;
begin
  AObj.Update;
end;

procedure ListView_BringToFront(AObj: TListView); stdcall;
begin
  AObj.BringToFront;
end;

function ListView_HasParent(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ListView_Hide(AObj: TListView); stdcall;
begin
  AObj.Hide;
end;

function ListView_Perform(AObj: TListView; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ListView_Refresh(AObj: TListView); stdcall;
begin
  AObj.Refresh;
end;

procedure ListView_SendToBack(AObj: TListView); stdcall;
begin
  AObj.SendToBack;
end;

procedure ListView_Show(AObj: TListView); stdcall;
begin
  AObj.Show;
end;

function ListView_GetTextBuf(AObj: TListView; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ListView_GetTextLen(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function ListView_FindComponent(AObj: TListView; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ListView_GetNamePath(AObj: TListView): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ListView_Assign(AObj: TListView; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ListView_ClassName(AObj: TListView): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ListView_Equals(AObj: TListView; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListView_GetHashCode(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ListView_ToString(AObj: TListView): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListView_GetAction(AObj: TListView): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ListView_SetAction(AObj: TListView; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function ListView_GetAlign(AObj: TListView): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ListView_SetAlign(AObj: TListView; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ListView_GetAllocBy(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.AllocBy;
end;

procedure ListView_SetAllocBy(AObj: TListView; AValue: Integer); stdcall;
begin
  AObj.AllocBy := AValue;
end;

function ListView_GetAnchors(AObj: TListView): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ListView_SetAnchors(AObj: TListView; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ListView_GetBevelEdges(AObj: TListView): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure ListView_SetBevelEdges(AObj: TListView; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function ListView_GetBevelInner(AObj: TListView): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure ListView_SetBevelInner(AObj: TListView; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function ListView_GetBevelOuter(AObj: TListView): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure ListView_SetBevelOuter(AObj: TListView; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function ListView_GetBevelKind(AObj: TListView): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure ListView_SetBevelKind(AObj: TListView; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function ListView_GetBiDiMode(AObj: TListView): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ListView_SetBiDiMode(AObj: TListView; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function ListView_GetBorderStyle(AObj: TListView): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure ListView_SetBorderStyle(AObj: TListView; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function ListView_GetBorderWidth(AObj: TListView): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure ListView_SetBorderWidth(AObj: TListView; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function ListView_GetCheckboxes(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.Checkboxes;
end;

procedure ListView_SetCheckboxes(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.Checkboxes := AValue;
end;

function ListView_GetColor(AObj: TListView): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure ListView_SetColor(AObj: TListView; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function ListView_GetColumns(AObj: TListView): TListColumns; stdcall;
begin
  Result :=  AObj.Columns;
end;

procedure ListView_SetColumns(AObj: TListView; AValue: TListColumns); stdcall;
begin
  AObj.Columns := AValue;
end;

function ListView_GetColumnClick(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.ColumnClick;
end;

procedure ListView_SetColumnClick(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.ColumnClick := AValue;
end;

function ListView_GetDoubleBuffered(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ListView_SetDoubleBuffered(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function ListView_GetEnabled(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ListView_SetEnabled(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ListView_GetFont(AObj: TListView): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure ListView_SetFont(AObj: TListView; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function ListView_GetFlatScrollBars(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.FlatScrollBars;
end;

procedure ListView_SetFlatScrollBars(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.FlatScrollBars := AValue;
end;

function ListView_GetFullDrag(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.FullDrag;
end;

procedure ListView_SetFullDrag(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.FullDrag := AValue;
end;

function ListView_GetGridLines(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.GridLines;
end;

procedure ListView_SetGridLines(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.GridLines := AValue;
end;

function ListView_GetGroups(AObj: TListView): TListGroups; stdcall;
begin
  Result :=  AObj.Groups;
end;

procedure ListView_SetGroups(AObj: TListView; AValue: TListGroups); stdcall;
begin
  AObj.Groups := AValue;
end;

function ListView_GetHideSelection(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.HideSelection;
end;

procedure ListView_SetHideSelection(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.HideSelection := AValue;
end;

function ListView_GetHotTrack(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.HotTrack;
end;

procedure ListView_SetHotTrack(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.HotTrack := AValue;
end;

function ListView_GetHoverTime(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.HoverTime;
end;

procedure ListView_SetHoverTime(AObj: TListView; AValue: Integer); stdcall;
begin
  AObj.HoverTime := AValue;
end;

function ListView_GetIconOptions(AObj: TListView): TIconOptions; stdcall;
begin
  Result :=  AObj.IconOptions;
end;

procedure ListView_SetIconOptions(AObj: TListView; AValue: TIconOptions); stdcall;
begin
  AObj.IconOptions := AValue;
end;

function ListView_GetItems(AObj: TListView): TListItems; stdcall;
begin
  Result :=  AObj.Items;
end;

procedure ListView_SetItems(AObj: TListView; AValue: TListItems); stdcall;
begin
  AObj.Items := AValue;
end;

function ListView_GetLargeImages(AObj: TListView): TCustomImageList; stdcall;
begin
  Result :=  AObj.LargeImages;
end;

procedure ListView_SetLargeImages(AObj: TListView; AValue: TCustomImageList); stdcall;
begin
  AObj.LargeImages := AValue;
end;

function ListView_GetMultiSelect(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.MultiSelect;
end;

procedure ListView_SetMultiSelect(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.MultiSelect := AValue;
end;

function ListView_GetStyleElements(AObj: TListView): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ListView_SetStyleElements(AObj: TListView; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function ListView_GetGroupHeaderImages(AObj: TListView): TCustomImageList; stdcall;
begin
  Result :=  AObj.GroupHeaderImages;
end;

procedure ListView_SetGroupHeaderImages(AObj: TListView; AValue: TCustomImageList); stdcall;
begin
  AObj.GroupHeaderImages := AValue;
end;

function ListView_GetGroupView(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.GroupView;
end;

procedure ListView_SetGroupView(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.GroupView := AValue;
end;

function ListView_GetReadOnly(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.ReadOnly;
end;

procedure ListView_SetReadOnly(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.ReadOnly := AValue;
end;

function ListView_GetRowSelect(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.RowSelect;
end;

procedure ListView_SetRowSelect(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.RowSelect := AValue;
end;

function ListView_GetParentColor(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure ListView_SetParentColor(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function ListView_GetParentDoubleBuffered(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure ListView_SetParentDoubleBuffered(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function ListView_GetParentFont(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure ListView_SetParentFont(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function ListView_GetParentShowHint(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ListView_SetParentShowHint(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ListView_GetPopupMenu(AObj: TListView): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ListView_SetPopupMenu(AObj: TListView; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ListView_GetShowColumnHeaders(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.ShowColumnHeaders;
end;

procedure ListView_SetShowColumnHeaders(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.ShowColumnHeaders := AValue;
end;

function ListView_GetShowWorkAreas(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.ShowWorkAreas;
end;

procedure ListView_SetShowWorkAreas(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.ShowWorkAreas := AValue;
end;

function ListView_GetShowHint(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ListView_SetShowHint(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ListView_GetSmallImages(AObj: TListView): TCustomImageList; stdcall;
begin
  Result :=  AObj.SmallImages;
end;

procedure ListView_SetSmallImages(AObj: TListView; AValue: TCustomImageList); stdcall;
begin
  AObj.SmallImages := AValue;
end;

function ListView_GetSortType(AObj: TListView): TSortType; stdcall;
begin
  Result :=  AObj.SortType;
end;

procedure ListView_SetSortType(AObj: TListView; AValue: TSortType); stdcall;
begin
  AObj.SortType := AValue;
end;

function ListView_GetStateImages(AObj: TListView): TCustomImageList; stdcall;
begin
  Result :=  AObj.StateImages;
end;

procedure ListView_SetStateImages(AObj: TListView; AValue: TCustomImageList); stdcall;
begin
  AObj.StateImages := AValue;
end;

function ListView_GetTabOrder(AObj: TListView): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure ListView_SetTabOrder(AObj: TListView; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function ListView_GetTabStop(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure ListView_SetTabStop(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function ListView_GetViewStyle(AObj: TListView): TViewStyle; stdcall;
begin
  Result :=  AObj.ViewStyle;
end;

procedure ListView_SetViewStyle(AObj: TListView; AValue: TViewStyle); stdcall;
begin
  AObj.ViewStyle := AValue;
end;

function ListView_GetVisible(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ListView_SetVisible(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

procedure ListView_SetOnAdvancedCustomDraw(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDraw := nil;
    TEventClass.Remove(AObj, geListViewAdvancedCustomDraw);
    Exit;
  end;
  AObj.OnAdvancedCustomDraw := TEventClass.ListViewOnAdvancedCustomDraw;
  TEventClass.Add(AObj, geListViewAdvancedCustomDraw, AEventId);
end;

procedure ListView_SetOnAdvancedCustomDrawItem(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDrawItem := nil;
    TEventClass.Remove(AObj, geListViewAdvancedCustomDrawItem);
    Exit;
  end;
  AObj.OnAdvancedCustomDrawItem := TEventClass.ListViewOnAdvancedCustomDrawItem;
  TEventClass.Add(AObj, geListViewAdvancedCustomDrawItem, AEventId);
end;

procedure ListView_SetOnAdvancedCustomDrawSubItem(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDrawSubItem := nil;
    TEventClass.Remove(AObj, geListViewAdvancedCustomDrawSubItem);
    Exit;
  end;
  AObj.OnAdvancedCustomDrawSubItem := TEventClass.ListViewOnAdvancedCustomDrawSubItem;
  TEventClass.Add(AObj, geListViewAdvancedCustomDrawSubItem, AEventId);
end;

procedure ListView_SetOnChange(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, geListViewChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.ListViewOnChange;
  TEventClass.Add(AObj, geListViewChange, AEventId);
end;

procedure ListView_SetOnClick(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnColumnClick(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnColumnClick := nil;
    TEventClass.Remove(AObj, geListViewColumnClick);
    Exit;
  end;
  AObj.OnColumnClick := TEventClass.ListViewOnColumnClick;
  TEventClass.Add(AObj, geListViewColumnClick, AEventId);
end;

procedure ListView_SetOnColumnRightClick(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnColumnRightClick := nil;
    TEventClass.Remove(AObj, geListViewColumnRightClick);
    Exit;
  end;
  AObj.OnColumnRightClick := TEventClass.ListViewOnColumnRightClick;
  TEventClass.Add(AObj, geListViewColumnRightClick, AEventId);
end;

procedure ListView_SetOnCompare(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnCompare := nil;
    TEventClass.Remove(AObj, geListViewCompare);
    Exit;
  end;
  AObj.OnCompare := TEventClass.ListViewOnCompare;
  TEventClass.Add(AObj, geListViewCompare, AEventId);
end;

procedure ListView_SetOnDblClick(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnEnter(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnExit(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnGetImageIndex(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetImageIndex := nil;
    TEventClass.Remove(AObj, geListViewGetImageIndex);
    Exit;
  end;
  AObj.OnGetImageIndex := TEventClass.ListViewOnGetImageIndex;
  TEventClass.Add(AObj, geListViewGetImageIndex, AEventId);
end;

procedure ListView_SetOnKeyDown(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnKeyPress(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnKeyUp(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnMouseDown(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnMouseEnter(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnMouseLeave(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnMouseMove(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnMouseUp(AObj: TListView; AEventId: NativeUInt); stdcall;
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

procedure ListView_SetOnResize(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnResize := nil;
    TEventClass.Remove(AObj, geResize);
    Exit;
  end;
  AObj.OnResize := TEventClass.OnResize;
  TEventClass.Add(AObj, geResize, AEventId);
end;

procedure ListView_SetOnSelectItem(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnSelectItem := nil;
    TEventClass.Remove(AObj, geListViewSelectItem);
    Exit;
  end;
  AObj.OnSelectItem := TEventClass.ListViewOnSelectItem;
  TEventClass.Add(AObj, geListViewSelectItem, AEventId);
end;

procedure ListView_SetOnItemChecked(AObj: TListView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnItemChecked := nil;
    TEventClass.Remove(AObj, geListViewItemChecked);
    Exit;
  end;
  AObj.OnItemChecked := TEventClass.ListViewOnItemChecked;
  TEventClass.Add(AObj, geListViewItemChecked, AEventId);
end;

function ListView_GetCanvas(AObj: TListView): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function ListView_GetDropTarget(AObj: TListView): TListItem; stdcall;
begin
  Result :=  AObj.DropTarget;
end;

procedure ListView_SetDropTarget(AObj: TListView; AValue: TListItem); stdcall;
begin
  AObj.DropTarget := AValue;
end;

function ListView_GetItemFocused(AObj: TListView): TListItem; stdcall;
begin
  Result :=  AObj.ItemFocused;
end;

procedure ListView_SetItemFocused(AObj: TListView; AValue: TListItem); stdcall;
begin
  AObj.ItemFocused := AValue;
end;

function ListView_GetSelCount(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.SelCount;
end;

function ListView_GetSelected(AObj: TListView): TListItem; stdcall;
begin
  Result :=  AObj.Selected;
end;

procedure ListView_SetSelected(AObj: TListView; AValue: TListItem); stdcall;
begin
  AObj.Selected := AValue;
end;

function ListView_GetTopItem(AObj: TListView): TListItem; stdcall;
begin
  Result :=  AObj.TopItem;
end;

function ListView_GetVisibleRowCount(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.VisibleRowCount;
end;

function ListView_GetItemIndex(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.ItemIndex;
end;

procedure ListView_SetItemIndex(AObj: TListView; AValue: Integer); stdcall;
begin
  AObj.ItemIndex := AValue;
end;

function ListView_GetBrush(AObj: TListView): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function ListView_GetControlCount(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function ListView_GetHandle(AObj: TListView): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ListView_GetParentWindow(AObj: TListView): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ListView_SetParentWindow(AObj: TListView; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

procedure ListView_GetBoundsRect(AObj: TListView; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ListView_SetBoundsRect(AObj: TListView; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ListView_GetClientHeight(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ListView_SetClientHeight(AObj: TListView; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ListView_GetClientRect(AObj: TListView; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ListView_GetClientWidth(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ListView_SetClientWidth(AObj: TListView; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ListView_GetExplicitLeft(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ListView_GetExplicitTop(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ListView_GetExplicitWidth(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ListView_GetExplicitHeight(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ListView_GetParent(AObj: TListView): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ListView_SetParent(AObj: TListView; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ListView_GetAlignWithMargins(AObj: TListView): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ListView_SetAlignWithMargins(AObj: TListView; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ListView_GetLeft(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ListView_SetLeft(AObj: TListView; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ListView_GetTop(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ListView_SetTop(AObj: TListView; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ListView_GetWidth(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ListView_SetWidth(AObj: TListView; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ListView_GetHeight(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ListView_SetHeight(AObj: TListView; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ListView_GetCursor(AObj: TListView): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ListView_SetCursor(AObj: TListView; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ListView_GetHint(AObj: TListView): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ListView_SetHint(AObj: TListView; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ListView_GetMargins(AObj: TListView): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ListView_SetMargins(AObj: TListView; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ListView_GetCustomHint(AObj: TListView): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ListView_SetCustomHint(AObj: TListView; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ListView_GetComponentCount(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ListView_GetComponentIndex(AObj: TListView): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ListView_SetComponentIndex(AObj: TListView; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ListView_GetOwner(AObj: TListView): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ListView_GetName(AObj: TListView): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ListView_SetName(AObj: TListView; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ListView_GetTag(AObj: TListView): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ListView_SetTag(AObj: TListView; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ListView_GetColumn(AObj: TListView; Index: Integer): TListColumn; stdcall;
begin
  Result :=  AObj.Column[Index];
end;

function ListView_GetControls(AObj: TListView; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ListView_GetComponents(AObj: TListView; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ListView_Create,
  ListView_Free,
  ListView_AddItem,
  ListView_AlphaSort,
  ListView_Clear,
  ListView_ClearSelection,
  ListView_DeleteSelected,
  ListView_GetSearchString,
  ListView_IsEditing,
  ListView_SelectAll,
  ListView_Scroll,
  ListView_CustomSort,
  ListView_CanFocus,
  ListView_FlipChildren,
  ListView_Focused,
  ListView_HandleAllocated,
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
  ListView_GetBevelEdges,
  ListView_SetBevelEdges,
  ListView_GetBevelInner,
  ListView_SetBevelInner,
  ListView_GetBevelOuter,
  ListView_SetBevelOuter,
  ListView_GetBevelKind,
  ListView_SetBevelKind,
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
  ListView_GetGroups,
  ListView_SetGroups,
  ListView_GetHideSelection,
  ListView_SetHideSelection,
  ListView_GetHotTrack,
  ListView_SetHotTrack,
  ListView_GetHoverTime,
  ListView_SetHoverTime,
  ListView_GetIconOptions,
  ListView_SetIconOptions,
  ListView_GetItems,
  ListView_SetItems,
  ListView_GetLargeImages,
  ListView_SetLargeImages,
  ListView_GetMultiSelect,
  ListView_SetMultiSelect,
  ListView_GetStyleElements,
  ListView_SetStyleElements,
  ListView_GetGroupHeaderImages,
  ListView_SetGroupHeaderImages,
  ListView_GetGroupView,
  ListView_SetGroupView,
  ListView_GetReadOnly,
  ListView_SetReadOnly,
  ListView_GetRowSelect,
  ListView_SetRowSelect,
  ListView_GetParentColor,
  ListView_SetParentColor,
  ListView_GetParentDoubleBuffered,
  ListView_SetParentDoubleBuffered,
  ListView_GetParentFont,
  ListView_SetParentFont,
  ListView_GetParentShowHint,
  ListView_SetParentShowHint,
  ListView_GetPopupMenu,
  ListView_SetPopupMenu,
  ListView_GetShowColumnHeaders,
  ListView_SetShowColumnHeaders,
  ListView_GetShowWorkAreas,
  ListView_SetShowWorkAreas,
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
  ListView_SetOnColumnRightClick,
  ListView_SetOnCompare,
  ListView_SetOnDblClick,
  ListView_SetOnEnter,
  ListView_SetOnExit,
  ListView_SetOnGetImageIndex,
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
  ListView_GetDropTarget,
  ListView_SetDropTarget,
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
  ListView_GetExplicitLeft,
  ListView_GetExplicitTop,
  ListView_GetExplicitWidth,
  ListView_GetExplicitHeight,
  ListView_GetParent,
  ListView_SetParent,
  ListView_GetAlignWithMargins,
  ListView_SetAlignWithMargins,
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
  ListView_GetMargins,
  ListView_SetMargins,
  ListView_GetCustomHint,
  ListView_SetCustomHint,
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

