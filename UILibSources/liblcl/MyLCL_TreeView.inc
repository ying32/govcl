
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function TreeView_Create(AOwner: TComponent): TTreeView; extdecl;
begin
  Result :=  TTreeView.Create(AOwner);
  Result.ScrollBars := ssAutoBoth;
end;

procedure TreeView_Free(AObj: TTreeView); extdecl;
begin
  AObj.Free;
end;

procedure TreeView_FullCollapse(AObj: TTreeView); extdecl;
begin
  AObj.FullCollapse;
end;

procedure TreeView_FullExpand(AObj: TTreeView); extdecl;
begin
  AObj.FullExpand;
end;

function TreeView_GetNodeAt(AObj: TTreeView; X: Integer; Y: Integer): TTreeNode; extdecl;
begin
  Result :=  AObj.GetNodeAt(X, Y);
end;

function TreeView_IsEditing(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.IsEditing;
end;

procedure TreeView_LoadFromFile(AObj: TTreeView; FileName: PChar); extdecl;
begin
  AObj.LoadFromFile(FileName);
end;

procedure TreeView_LoadFromStream(AObj: TTreeView; Stream: TStream); extdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure TreeView_SaveToFile(AObj: TTreeView; FileName: PChar); extdecl;
begin
  AObj.SaveToFile(FileName);
end;

procedure TreeView_SaveToStream(AObj: TTreeView; Stream: TStream); extdecl;
begin
  AObj.SaveToStream(Stream);
end;

procedure TreeView_ClearSelection(AObj: TTreeView; KeepPrimary: LongBool); extdecl;
begin
  AObj.ClearSelection(KeepPrimary);
end;

function TreeView_CanFocus(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure TreeView_FlipChildren(AObj: TTreeView; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function TreeView_Focused(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure TreeView_Invalidate(AObj: TTreeView); extdecl;
begin
  AObj.Invalidate;
end;

procedure TreeView_Realign(AObj: TTreeView); extdecl;
begin
  AObj.Realign;
end;

procedure TreeView_Repaint(AObj: TTreeView); extdecl;
begin
  AObj.Repaint;
end;

procedure TreeView_ScaleBy(AObj: TTreeView; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure TreeView_SetBounds(AObj: TTreeView; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure TreeView_SetFocus(AObj: TTreeView); extdecl;
begin
  AObj.SetFocus;
end;

procedure TreeView_Update(AObj: TTreeView); extdecl;
begin
  AObj.Update;
end;

procedure TreeView_BringToFront(AObj: TTreeView); extdecl;
begin
  AObj.BringToFront;
end;

function TreeView_HasParent(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure TreeView_Hide(AObj: TTreeView); extdecl;
begin
  AObj.Hide;
end;

function TreeView_Perform(AObj: TTreeView; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure TreeView_Refresh(AObj: TTreeView); extdecl;
begin
  AObj.Refresh;
end;

procedure TreeView_SendToBack(AObj: TTreeView); extdecl;
begin
  AObj.SendToBack;
end;

procedure TreeView_Show(AObj: TTreeView); extdecl;
begin
  AObj.Show;
end;

function TreeView_GetTextBuf(AObj: TTreeView; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function TreeView_GetTextLen(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function TreeView_FindComponent(AObj: TTreeView; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function TreeView_GetNamePath(AObj: TTreeView): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TreeView_Assign(AObj: TTreeView; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function TreeView_ClassName(AObj: TTreeView): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function TreeView_Equals(AObj: TTreeView; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TreeView_GetHashCode(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TreeView_ToString(AObj: TTreeView): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function TreeView_GetAlign(AObj: TTreeView): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure TreeView_SetAlign(AObj: TTreeView; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function TreeView_GetAnchors(AObj: TTreeView): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure TreeView_SetAnchors(AObj: TTreeView; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function TreeView_GetAutoExpand(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.AutoExpand;
end;

procedure TreeView_SetAutoExpand(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.AutoExpand := AValue;
end;

function TreeView_GetBiDiMode(AObj: TTreeView): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure TreeView_SetBiDiMode(AObj: TTreeView; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function TreeView_GetBorderStyle(AObj: TTreeView): TBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure TreeView_SetBorderStyle(AObj: TTreeView; AValue: TBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function TreeView_GetBorderWidth(AObj: TTreeView): TBorderWidth; extdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure TreeView_SetBorderWidth(AObj: TTreeView; AValue: TBorderWidth); extdecl;
begin
  AObj.BorderWidth := AValue;
end;

function TreeView_GetColor(AObj: TTreeView): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure TreeView_SetColor(AObj: TTreeView; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function TreeView_GetDoubleBuffered(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure TreeView_SetDoubleBuffered(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function TreeView_GetEnabled(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure TreeView_SetEnabled(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function TreeView_GetFont(AObj: TTreeView): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure TreeView_SetFont(AObj: TTreeView; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function TreeView_GetHideSelection(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.HideSelection;
end;

procedure TreeView_SetHideSelection(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.HideSelection := AValue;
end;

function TreeView_GetHotTrack(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.HotTrack;
end;

procedure TreeView_SetHotTrack(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.HotTrack := AValue;
end;

function TreeView_GetImages(AObj: TTreeView): TCustomImageList; extdecl;
begin
  Result :=  AObj.Images;
end;

procedure TreeView_SetImages(AObj: TTreeView; AValue: TCustomImageList); extdecl;
begin
  AObj.Images := AValue;
end;

function TreeView_GetIndent(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.Indent;
end;

procedure TreeView_SetIndent(AObj: TTreeView; AValue: Integer); extdecl;
begin
  AObj.Indent := AValue;
end;

function TreeView_GetMultiSelect(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.MultiSelect;
end;

procedure TreeView_SetMultiSelect(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.MultiSelect := AValue;
end;

function TreeView_GetMultiSelectStyle(AObj: TTreeView): TMultiSelectStyle; extdecl;
begin
  Result :=  AObj.MultiSelectStyle;
end;

procedure TreeView_SetMultiSelectStyle(AObj: TTreeView; AValue: TMultiSelectStyle); extdecl;
begin
  AObj.MultiSelectStyle := AValue;
end;

function TreeView_GetParentColor(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure TreeView_SetParentColor(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function TreeView_GetParentFont(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure TreeView_SetParentFont(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function TreeView_GetParentShowHint(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure TreeView_SetParentShowHint(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function TreeView_GetPopupMenu(AObj: TTreeView): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure TreeView_SetPopupMenu(AObj: TTreeView; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function TreeView_GetReadOnly(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.ReadOnly;
end;

procedure TreeView_SetReadOnly(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.ReadOnly := AValue;
end;

function TreeView_GetRightClickSelect(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.RightClickSelect;
end;

procedure TreeView_SetRightClickSelect(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.RightClickSelect := AValue;
end;

function TreeView_GetRowSelect(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.RowSelect;
end;

procedure TreeView_SetRowSelect(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.RowSelect := AValue;
end;

function TreeView_GetShowButtons(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.ShowButtons;
end;

procedure TreeView_SetShowButtons(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.ShowButtons := AValue;
end;

function TreeView_GetShowHint(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure TreeView_SetShowHint(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function TreeView_GetShowLines(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.ShowLines;
end;

procedure TreeView_SetShowLines(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.ShowLines := AValue;
end;

function TreeView_GetShowRoot(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.ShowRoot;
end;

procedure TreeView_SetShowRoot(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.ShowRoot := AValue;
end;

function TreeView_GetSortType(AObj: TTreeView): TSortType; extdecl;
begin
  Result :=  AObj.SortType;
end;

procedure TreeView_SetSortType(AObj: TTreeView; AValue: TSortType); extdecl;
begin
  AObj.SortType := AValue;
end;

function TreeView_GetStateImages(AObj: TTreeView): TCustomImageList; extdecl;
begin
  Result :=  AObj.StateImages;
end;

procedure TreeView_SetStateImages(AObj: TTreeView; AValue: TCustomImageList); extdecl;
begin
  AObj.StateImages := AValue;
end;

function TreeView_GetTabOrder(AObj: TTreeView): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure TreeView_SetTabOrder(AObj: TTreeView; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function TreeView_GetTabStop(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure TreeView_SetTabStop(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function TreeView_GetToolTips(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.ToolTips;
end;

procedure TreeView_SetToolTips(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.ToolTips := AValue;
end;

function TreeView_GetVisible(AObj: TTreeView): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure TreeView_SetVisible(AObj: TTreeView; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure TreeView_SetOnAdvancedCustomDraw(AObj: TTreeView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDraw := nil;
    EventClass.Remove(AObj, geTreeViewAdvancedCustomDraw);
    Exit;
  end;
  AObj.OnAdvancedCustomDraw := @EventClass.TreeViewOnAdvancedCustomDraw;
  EventClass.Add(AObj, geTreeViewAdvancedCustomDraw, AEventId);
end;

procedure TreeView_SetOnAdvancedCustomDrawItem(AObj: TTreeView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDrawItem := nil;
    EventClass.Remove(AObj, geTreeViewAdvancedCustomDrawItem);
    Exit;
  end;
  AObj.OnAdvancedCustomDrawItem := @EventClass.TreeViewOnAdvancedCustomDrawItem;
  EventClass.Add(AObj, geTreeViewAdvancedCustomDrawItem, AEventId);
end;

procedure TreeView_SetOnChange(AObj: TTreeView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    EventClass.Remove(AObj, geTreeViewChange);
    Exit;
  end;
  AObj.OnChange := @EventClass.TreeViewOnChange;
  EventClass.Add(AObj, geTreeViewChange, AEventId);
end;

procedure TreeView_SetOnClick(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnCompare(AObj: TTreeView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnCompare := nil;
    EventClass.Remove(AObj, geTreeViewCompare);
    Exit;
  end;
  AObj.OnCompare := @EventClass.TreeViewOnCompare;
  EventClass.Add(AObj, geTreeViewCompare, AEventId);
end;

procedure TreeView_SetOnDblClick(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnEnter(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnExit(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnGetSelectedIndex(AObj: TTreeView; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetSelectedIndex := nil;
    EventClass.Remove(AObj, geTreeViewGetSelectedIndex);
    Exit;
  end;
  AObj.OnGetSelectedIndex := @EventClass.TreeViewOnGetSelectedIndex;
  EventClass.Add(AObj, geTreeViewGetSelectedIndex, AEventId);
end;

procedure TreeView_SetOnKeyDown(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnKeyPress(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnKeyUp(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnMouseDown(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnMouseEnter(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnMouseLeave(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnMouseMove(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

procedure TreeView_SetOnMouseUp(AObj: TTreeView; AEventId: NativeUInt); extdecl;
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

function TreeView_GetItems(AObj: TTreeView): TTreeNodes; extdecl;
begin
  Result :=  AObj.Items;
end;

procedure TreeView_SetItems(AObj: TTreeView; AValue: TTreeNodes); extdecl;
begin
  AObj.Items := AValue;
end;

function TreeView_GetCanvas(AObj: TTreeView): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function TreeView_GetSelected(AObj: TTreeView): TTreeNode; extdecl;
begin
  Result :=  AObj.Selected;
end;

procedure TreeView_SetSelected(AObj: TTreeView; AValue: TTreeNode); extdecl;
begin
  AObj.Selected := AValue;
end;

function TreeView_GetTopItem(AObj: TTreeView): TTreeNode; extdecl;
begin
  Result :=  AObj.TopItem;
end;

procedure TreeView_SetTopItem(AObj: TTreeView; AValue: TTreeNode); extdecl;
begin
  AObj.TopItem := AValue;
end;

function TreeView_GetSelectionCount(AObj: TTreeView): Cardinal; extdecl;
begin
  Result :=  AObj.SelectionCount;
end;

function TreeView_GetBrush(AObj: TTreeView): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function TreeView_GetControlCount(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function TreeView_GetHandle(AObj: TTreeView): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function TreeView_GetParentWindow(AObj: TTreeView): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure TreeView_SetParentWindow(AObj: TTreeView; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function TreeView_GetAction(AObj: TTreeView): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure TreeView_SetAction(AObj: TTreeView; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure TreeView_GetBoundsRect(AObj: TTreeView; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure TreeView_SetBoundsRect(AObj: TTreeView; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function TreeView_GetClientHeight(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure TreeView_SetClientHeight(AObj: TTreeView; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure TreeView_GetClientRect(AObj: TTreeView; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function TreeView_GetClientWidth(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure TreeView_SetClientWidth(AObj: TTreeView; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function TreeView_GetParent(AObj: TTreeView): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure TreeView_SetParent(AObj: TTreeView; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function TreeView_GetLeft(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure TreeView_SetLeft(AObj: TTreeView; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function TreeView_GetTop(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure TreeView_SetTop(AObj: TTreeView; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function TreeView_GetWidth(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure TreeView_SetWidth(AObj: TTreeView; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function TreeView_GetHeight(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure TreeView_SetHeight(AObj: TTreeView; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function TreeView_GetCursor(AObj: TTreeView): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure TreeView_SetCursor(AObj: TTreeView; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function TreeView_GetHint(AObj: TTreeView): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure TreeView_SetHint(AObj: TTreeView; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function TreeView_GetComponentCount(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function TreeView_GetComponentIndex(AObj: TTreeView): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure TreeView_SetComponentIndex(AObj: TTreeView; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function TreeView_GetOwner(AObj: TTreeView): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function TreeView_GetName(AObj: TTreeView): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TreeView_SetName(AObj: TTreeView; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function TreeView_GetTag(AObj: TTreeView): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure TreeView_SetTag(AObj: TTreeView; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function TreeView_GetSelections(AObj: TTreeView; Index: Integer): TTreeNode; extdecl;
begin
  Result :=  AObj.Selections[Index];
end;

function TreeView_GetControls(AObj: TTreeView; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function TreeView_GetComponents(AObj: TTreeView; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  TreeView_Create,
  TreeView_Free,
  TreeView_FullCollapse,
  TreeView_FullExpand,
  TreeView_GetNodeAt,
  TreeView_IsEditing,
  TreeView_LoadFromFile,
  TreeView_LoadFromStream,
  TreeView_SaveToFile,
  TreeView_SaveToStream,
  TreeView_ClearSelection,
  TreeView_CanFocus,
  TreeView_FlipChildren,
  TreeView_Focused,
  TreeView_Invalidate,
  TreeView_Realign,
  TreeView_Repaint,
  TreeView_ScaleBy,
  TreeView_SetBounds,
  TreeView_SetFocus,
  TreeView_Update,
  TreeView_BringToFront,
  TreeView_HasParent,
  TreeView_Hide,
  TreeView_Perform,
  TreeView_Refresh,
  TreeView_SendToBack,
  TreeView_Show,
  TreeView_GetTextBuf,
  TreeView_GetTextLen,
  TreeView_FindComponent,
  TreeView_GetNamePath,
  TreeView_Assign,
  TreeView_ClassName,
  TreeView_Equals,
  TreeView_GetHashCode,
  TreeView_ToString,
  TreeView_GetAlign,
  TreeView_SetAlign,
  TreeView_GetAnchors,
  TreeView_SetAnchors,
  TreeView_GetAutoExpand,
  TreeView_SetAutoExpand,
  TreeView_GetBiDiMode,
  TreeView_SetBiDiMode,
  TreeView_GetBorderStyle,
  TreeView_SetBorderStyle,
  TreeView_GetBorderWidth,
  TreeView_SetBorderWidth,
  TreeView_GetColor,
  TreeView_SetColor,
  TreeView_GetDoubleBuffered,
  TreeView_SetDoubleBuffered,
  TreeView_GetEnabled,
  TreeView_SetEnabled,
  TreeView_GetFont,
  TreeView_SetFont,
  TreeView_GetHideSelection,
  TreeView_SetHideSelection,
  TreeView_GetHotTrack,
  TreeView_SetHotTrack,
  TreeView_GetImages,
  TreeView_SetImages,
  TreeView_GetIndent,
  TreeView_SetIndent,
  TreeView_GetMultiSelect,
  TreeView_SetMultiSelect,
  TreeView_GetMultiSelectStyle,
  TreeView_SetMultiSelectStyle,
  TreeView_GetParentColor,
  TreeView_SetParentColor,
  TreeView_GetParentFont,
  TreeView_SetParentFont,
  TreeView_GetParentShowHint,
  TreeView_SetParentShowHint,
  TreeView_GetPopupMenu,
  TreeView_SetPopupMenu,
  TreeView_GetReadOnly,
  TreeView_SetReadOnly,
  TreeView_GetRightClickSelect,
  TreeView_SetRightClickSelect,
  TreeView_GetRowSelect,
  TreeView_SetRowSelect,
  TreeView_GetShowButtons,
  TreeView_SetShowButtons,
  TreeView_GetShowHint,
  TreeView_SetShowHint,
  TreeView_GetShowLines,
  TreeView_SetShowLines,
  TreeView_GetShowRoot,
  TreeView_SetShowRoot,
  TreeView_GetSortType,
  TreeView_SetSortType,
  TreeView_GetStateImages,
  TreeView_SetStateImages,
  TreeView_GetTabOrder,
  TreeView_SetTabOrder,
  TreeView_GetTabStop,
  TreeView_SetTabStop,
  TreeView_GetToolTips,
  TreeView_SetToolTips,
  TreeView_GetVisible,
  TreeView_SetVisible,
  TreeView_SetOnAdvancedCustomDraw,
  TreeView_SetOnAdvancedCustomDrawItem,
  TreeView_SetOnChange,
  TreeView_SetOnClick,
  TreeView_SetOnCompare,
  TreeView_SetOnDblClick,
  TreeView_SetOnEnter,
  TreeView_SetOnExit,
  TreeView_SetOnGetSelectedIndex,
  TreeView_SetOnKeyDown,
  TreeView_SetOnKeyPress,
  TreeView_SetOnKeyUp,
  TreeView_SetOnMouseDown,
  TreeView_SetOnMouseEnter,
  TreeView_SetOnMouseLeave,
  TreeView_SetOnMouseMove,
  TreeView_SetOnMouseUp,
  TreeView_GetItems,
  TreeView_SetItems,
  TreeView_GetCanvas,
  TreeView_GetSelected,
  TreeView_SetSelected,
  TreeView_GetTopItem,
  TreeView_SetTopItem,
  TreeView_GetSelectionCount,
  TreeView_GetBrush,
  TreeView_GetControlCount,
  TreeView_GetHandle,
  TreeView_GetParentWindow,
  TreeView_SetParentWindow,
  TreeView_GetAction,
  TreeView_SetAction,
  TreeView_GetBoundsRect,
  TreeView_SetBoundsRect,
  TreeView_GetClientHeight,
  TreeView_SetClientHeight,
  TreeView_GetClientRect,
  TreeView_GetClientWidth,
  TreeView_SetClientWidth,
  TreeView_GetParent,
  TreeView_SetParent,
  TreeView_GetLeft,
  TreeView_SetLeft,
  TreeView_GetTop,
  TreeView_SetTop,
  TreeView_GetWidth,
  TreeView_SetWidth,
  TreeView_GetHeight,
  TreeView_SetHeight,
  TreeView_GetCursor,
  TreeView_SetCursor,
  TreeView_GetHint,
  TreeView_SetHint,
  TreeView_GetComponentCount,
  TreeView_GetComponentIndex,
  TreeView_SetComponentIndex,
  TreeView_GetOwner,
  TreeView_GetName,
  TreeView_SetName,
  TreeView_GetTag,
  TreeView_SetTag,
  TreeView_GetSelections,
  TreeView_GetControls,
  TreeView_GetComponents;

