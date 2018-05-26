
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function TreeView_Create(AOwner: TComponent): TTreeView; stdcall;
begin
  Result :=  TTreeView.Create(AOwner);
end;

procedure TreeView_Free(AObj: TTreeView); stdcall;
begin
  AObj.Free;
end;

function TreeView_AlphaSort(AObj: TTreeView; ARecurse: LongBool): LongBool; stdcall;
begin
  Result :=  AObj.AlphaSort(ARecurse);
end;

procedure TreeView_FullCollapse(AObj: TTreeView); stdcall;
begin
  AObj.FullCollapse;
end;

procedure TreeView_FullExpand(AObj: TTreeView); stdcall;
begin
  AObj.FullExpand;
end;

function TreeView_GetNodeAt(AObj: TTreeView; X: Integer; Y: Integer): TTreeNode; stdcall;
begin
  Result :=  AObj.GetNodeAt(X, Y);
end;

function TreeView_IsEditing(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.IsEditing;
end;

procedure TreeView_LoadFromFile(AObj: TTreeView; FileName: PChar); stdcall;
begin
  AObj.LoadFromFile(FileName);
end;

procedure TreeView_LoadFromStream(AObj: TTreeView; Stream: TStream); stdcall;
begin
  AObj.LoadFromStream(Stream);
end;

procedure TreeView_SaveToFile(AObj: TTreeView; FileName: PChar); stdcall;
begin
  AObj.SaveToFile(FileName);
end;

procedure TreeView_SaveToStream(AObj: TTreeView; Stream: TStream); stdcall;
begin
  AObj.SaveToStream(Stream);
end;

procedure TreeView_Deselect(AObj: TTreeView; Node: TTreeNode); stdcall;
begin
  AObj.Deselect(Node);
end;

procedure TreeView_Subselect(AObj: TTreeView; Node: TTreeNode; Validate: LongBool); stdcall;
begin
  AObj.Subselect(Node, Validate);
end;

procedure TreeView_ClearSelection(AObj: TTreeView; KeepPrimary: LongBool); stdcall;
begin
  AObj.ClearSelection(KeepPrimary);
end;

function TreeView_FindNextToSelect(AObj: TTreeView): TTreeNode; stdcall;
begin
  Result :=  AObj.FindNextToSelect;
end;

function TreeView_CustomSort(AObj: TTreeView; SortProc: PFNTVCOMPARE; Data: NativeInt; ARecurse: LongBool): LongBool; stdcall;
begin
  Result :=  AObj.CustomSort(SortProc, Data, ARecurse);
end;

function TreeView_CanFocus(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure TreeView_FlipChildren(AObj: TTreeView; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function TreeView_Focused(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function TreeView_HandleAllocated(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure TreeView_Invalidate(AObj: TTreeView); stdcall;
begin
  AObj.Invalidate;
end;

procedure TreeView_Realign(AObj: TTreeView); stdcall;
begin
  AObj.Realign;
end;

procedure TreeView_Repaint(AObj: TTreeView); stdcall;
begin
  AObj.Repaint;
end;

procedure TreeView_ScaleBy(AObj: TTreeView; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure TreeView_SetBounds(AObj: TTreeView; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure TreeView_SetFocus(AObj: TTreeView); stdcall;
begin
  AObj.SetFocus;
end;

procedure TreeView_Update(AObj: TTreeView); stdcall;
begin
  AObj.Update;
end;

procedure TreeView_BringToFront(AObj: TTreeView); stdcall;
begin
  AObj.BringToFront;
end;

function TreeView_HasParent(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure TreeView_Hide(AObj: TTreeView); stdcall;
begin
  AObj.Hide;
end;

function TreeView_Perform(AObj: TTreeView; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure TreeView_Refresh(AObj: TTreeView); stdcall;
begin
  AObj.Refresh;
end;

procedure TreeView_SendToBack(AObj: TTreeView); stdcall;
begin
  AObj.SendToBack;
end;

procedure TreeView_Show(AObj: TTreeView); stdcall;
begin
  AObj.Show;
end;

function TreeView_GetTextBuf(AObj: TTreeView; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function TreeView_GetTextLen(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function TreeView_FindComponent(AObj: TTreeView; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function TreeView_GetNamePath(AObj: TTreeView): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TreeView_Assign(AObj: TTreeView; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function TreeView_ClassName(AObj: TTreeView): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TreeView_Equals(AObj: TTreeView; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function TreeView_GetHashCode(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function TreeView_ToString(AObj: TTreeView): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function TreeView_GetAlign(AObj: TTreeView): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure TreeView_SetAlign(AObj: TTreeView; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function TreeView_GetAnchors(AObj: TTreeView): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure TreeView_SetAnchors(AObj: TTreeView; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function TreeView_GetAutoExpand(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.AutoExpand;
end;

procedure TreeView_SetAutoExpand(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.AutoExpand := AValue;
end;

function TreeView_GetBevelEdges(AObj: TTreeView): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure TreeView_SetBevelEdges(AObj: TTreeView; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function TreeView_GetBevelInner(AObj: TTreeView): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure TreeView_SetBevelInner(AObj: TTreeView; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function TreeView_GetBevelOuter(AObj: TTreeView): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure TreeView_SetBevelOuter(AObj: TTreeView; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function TreeView_GetBevelKind(AObj: TTreeView): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure TreeView_SetBevelKind(AObj: TTreeView; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function TreeView_GetBiDiMode(AObj: TTreeView): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure TreeView_SetBiDiMode(AObj: TTreeView; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function TreeView_GetBorderStyle(AObj: TTreeView): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure TreeView_SetBorderStyle(AObj: TTreeView; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function TreeView_GetBorderWidth(AObj: TTreeView): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure TreeView_SetBorderWidth(AObj: TTreeView; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function TreeView_GetChangeDelay(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.ChangeDelay;
end;

procedure TreeView_SetChangeDelay(AObj: TTreeView; AValue: Integer); stdcall;
begin
  AObj.ChangeDelay := AValue;
end;

function TreeView_GetColor(AObj: TTreeView): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure TreeView_SetColor(AObj: TTreeView; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function TreeView_GetDoubleBuffered(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure TreeView_SetDoubleBuffered(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function TreeView_GetEnabled(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure TreeView_SetEnabled(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function TreeView_GetFont(AObj: TTreeView): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure TreeView_SetFont(AObj: TTreeView; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function TreeView_GetHideSelection(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.HideSelection;
end;

procedure TreeView_SetHideSelection(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.HideSelection := AValue;
end;

function TreeView_GetHotTrack(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.HotTrack;
end;

procedure TreeView_SetHotTrack(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.HotTrack := AValue;
end;

function TreeView_GetImages(AObj: TTreeView): TCustomImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure TreeView_SetImages(AObj: TTreeView; AValue: TCustomImageList); stdcall;
begin
  AObj.Images := AValue;
end;

function TreeView_GetIndent(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.Indent;
end;

procedure TreeView_SetIndent(AObj: TTreeView; AValue: Integer); stdcall;
begin
  AObj.Indent := AValue;
end;

function TreeView_GetMultiSelect(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.MultiSelect;
end;

procedure TreeView_SetMultiSelect(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.MultiSelect := AValue;
end;

function TreeView_GetMultiSelectStyle(AObj: TTreeView): TMultiSelectStyle; stdcall;
begin
  Result :=  AObj.MultiSelectStyle;
end;

procedure TreeView_SetMultiSelectStyle(AObj: TTreeView; AValue: TMultiSelectStyle); stdcall;
begin
  AObj.MultiSelectStyle := AValue;
end;

function TreeView_GetParentColor(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure TreeView_SetParentColor(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function TreeView_GetParentCtl3D(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure TreeView_SetParentCtl3D(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function TreeView_GetParentDoubleBuffered(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure TreeView_SetParentDoubleBuffered(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function TreeView_GetParentFont(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure TreeView_SetParentFont(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function TreeView_GetParentShowHint(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure TreeView_SetParentShowHint(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function TreeView_GetPopupMenu(AObj: TTreeView): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure TreeView_SetPopupMenu(AObj: TTreeView; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function TreeView_GetReadOnly(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ReadOnly;
end;

procedure TreeView_SetReadOnly(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ReadOnly := AValue;
end;

function TreeView_GetRightClickSelect(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.RightClickSelect;
end;

procedure TreeView_SetRightClickSelect(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.RightClickSelect := AValue;
end;

function TreeView_GetRowSelect(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.RowSelect;
end;

procedure TreeView_SetRowSelect(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.RowSelect := AValue;
end;

function TreeView_GetShowButtons(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ShowButtons;
end;

procedure TreeView_SetShowButtons(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ShowButtons := AValue;
end;

function TreeView_GetShowHint(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure TreeView_SetShowHint(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function TreeView_GetShowLines(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ShowLines;
end;

procedure TreeView_SetShowLines(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ShowLines := AValue;
end;

function TreeView_GetShowRoot(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ShowRoot;
end;

procedure TreeView_SetShowRoot(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ShowRoot := AValue;
end;

function TreeView_GetSortType(AObj: TTreeView): TSortType; stdcall;
begin
  Result :=  AObj.SortType;
end;

procedure TreeView_SetSortType(AObj: TTreeView; AValue: TSortType); stdcall;
begin
  AObj.SortType := AValue;
end;

function TreeView_GetStateImages(AObj: TTreeView): TCustomImageList; stdcall;
begin
  Result :=  AObj.StateImages;
end;

procedure TreeView_SetStateImages(AObj: TTreeView; AValue: TCustomImageList); stdcall;
begin
  AObj.StateImages := AValue;
end;

function TreeView_GetTabOrder(AObj: TTreeView): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure TreeView_SetTabOrder(AObj: TTreeView; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function TreeView_GetTabStop(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure TreeView_SetTabStop(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function TreeView_GetToolTips(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.ToolTips;
end;

procedure TreeView_SetToolTips(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.ToolTips := AValue;
end;

function TreeView_GetVisible(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure TreeView_SetVisible(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function TreeView_GetStyleElements(AObj: TTreeView): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure TreeView_SetStyleElements(AObj: TTreeView; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure TreeView_SetOnAdvancedCustomDraw(AObj: TTreeView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDraw := nil;
    TEventClass.Remove(AObj, geTreeViewAdvancedCustomDraw);
    Exit;
  end;
  AObj.OnAdvancedCustomDraw := TEventClass.TreeViewOnAdvancedCustomDraw;
  TEventClass.Add(AObj, geTreeViewAdvancedCustomDraw, AEventId);
end;

procedure TreeView_SetOnAdvancedCustomDrawItem(AObj: TTreeView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnAdvancedCustomDrawItem := nil;
    TEventClass.Remove(AObj, geTreeViewAdvancedCustomDrawItem);
    Exit;
  end;
  AObj.OnAdvancedCustomDrawItem := TEventClass.TreeViewOnAdvancedCustomDrawItem;
  TEventClass.Add(AObj, geTreeViewAdvancedCustomDrawItem, AEventId);
end;

procedure TreeView_SetOnChange(AObj: TTreeView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, geTreeViewChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.TreeViewOnChange;
  TEventClass.Add(AObj, geTreeViewChange, AEventId);
end;

procedure TreeView_SetOnClick(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnCompare(AObj: TTreeView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnCompare := nil;
    TEventClass.Remove(AObj, geTreeViewCompare);
    Exit;
  end;
  AObj.OnCompare := TEventClass.TreeViewOnCompare;
  TEventClass.Add(AObj, geTreeViewCompare, AEventId);
end;

procedure TreeView_SetOnDblClick(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnEnter(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnExit(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnGetImageIndex(AObj: TTreeView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetImageIndex := nil;
    TEventClass.Remove(AObj, geTreeViewGetImageIndex);
    Exit;
  end;
  AObj.OnGetImageIndex := TEventClass.TreeViewOnGetImageIndex;
  TEventClass.Add(AObj, geTreeViewGetImageIndex, AEventId);
end;

procedure TreeView_SetOnGetSelectedIndex(AObj: TTreeView; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetSelectedIndex := nil;
    TEventClass.Remove(AObj, geTreeViewGetSelectedIndex);
    Exit;
  end;
  AObj.OnGetSelectedIndex := TEventClass.TreeViewOnGetSelectedIndex;
  TEventClass.Add(AObj, geTreeViewGetSelectedIndex, AEventId);
end;

procedure TreeView_SetOnKeyDown(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnKeyPress(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnKeyUp(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnMouseDown(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnMouseEnter(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnMouseLeave(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnMouseMove(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

procedure TreeView_SetOnMouseUp(AObj: TTreeView; AEventId: NativeUInt); stdcall;
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

function TreeView_GetItems(AObj: TTreeView): TTreeNodes; stdcall;
begin
  Result :=  AObj.Items;
end;

procedure TreeView_SetItems(AObj: TTreeView; AValue: TTreeNodes); stdcall;
begin
  AObj.Items := AValue;
end;

function TreeView_GetCanvas(AObj: TTreeView): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function TreeView_GetDropTarget(AObj: TTreeView): TTreeNode; stdcall;
begin
  Result :=  AObj.DropTarget;
end;

procedure TreeView_SetDropTarget(AObj: TTreeView; AValue: TTreeNode); stdcall;
begin
  AObj.DropTarget := AValue;
end;

function TreeView_GetSelected(AObj: TTreeView): TTreeNode; stdcall;
begin
  Result :=  AObj.Selected;
end;

procedure TreeView_SetSelected(AObj: TTreeView; AValue: TTreeNode); stdcall;
begin
  AObj.Selected := AValue;
end;

function TreeView_GetTopItem(AObj: TTreeView): TTreeNode; stdcall;
begin
  Result :=  AObj.TopItem;
end;

procedure TreeView_SetTopItem(AObj: TTreeView; AValue: TTreeNode); stdcall;
begin
  AObj.TopItem := AValue;
end;

function TreeView_GetSelectionCount(AObj: TTreeView): Cardinal; stdcall;
begin
  Result :=  AObj.SelectionCount;
end;

function TreeView_GetBrush(AObj: TTreeView): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function TreeView_GetControlCount(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function TreeView_GetHandle(AObj: TTreeView): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function TreeView_GetParentWindow(AObj: TTreeView): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure TreeView_SetParentWindow(AObj: TTreeView; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function TreeView_GetAction(AObj: TTreeView): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure TreeView_SetAction(AObj: TTreeView; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure TreeView_GetBoundsRect(AObj: TTreeView; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure TreeView_SetBoundsRect(AObj: TTreeView; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function TreeView_GetClientHeight(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure TreeView_SetClientHeight(AObj: TTreeView; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure TreeView_GetClientRect(AObj: TTreeView; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function TreeView_GetClientWidth(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure TreeView_SetClientWidth(AObj: TTreeView; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function TreeView_GetExplicitLeft(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function TreeView_GetExplicitTop(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function TreeView_GetExplicitWidth(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function TreeView_GetExplicitHeight(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function TreeView_GetParent(AObj: TTreeView): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure TreeView_SetParent(AObj: TTreeView; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function TreeView_GetAlignWithMargins(AObj: TTreeView): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure TreeView_SetAlignWithMargins(AObj: TTreeView; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function TreeView_GetLeft(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure TreeView_SetLeft(AObj: TTreeView; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function TreeView_GetTop(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure TreeView_SetTop(AObj: TTreeView; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function TreeView_GetWidth(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure TreeView_SetWidth(AObj: TTreeView; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function TreeView_GetHeight(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure TreeView_SetHeight(AObj: TTreeView; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function TreeView_GetCursor(AObj: TTreeView): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure TreeView_SetCursor(AObj: TTreeView; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function TreeView_GetHint(AObj: TTreeView): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure TreeView_SetHint(AObj: TTreeView; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function TreeView_GetMargins(AObj: TTreeView): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure TreeView_SetMargins(AObj: TTreeView; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function TreeView_GetCustomHint(AObj: TTreeView): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure TreeView_SetCustomHint(AObj: TTreeView; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function TreeView_GetComponentCount(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function TreeView_GetComponentIndex(AObj: TTreeView): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure TreeView_SetComponentIndex(AObj: TTreeView; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function TreeView_GetOwner(AObj: TTreeView): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function TreeView_GetName(AObj: TTreeView): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TreeView_SetName(AObj: TTreeView; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function TreeView_GetTag(AObj: TTreeView): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure TreeView_SetTag(AObj: TTreeView; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function TreeView_GetSelections(AObj: TTreeView; Index: Integer): TTreeNode; stdcall;
begin
  Result :=  AObj.Selections[Index];
end;

function TreeView_GetControls(AObj: TTreeView; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function TreeView_GetComponents(AObj: TTreeView; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  TreeView_Create,
  TreeView_Free,
  TreeView_AlphaSort,
  TreeView_FullCollapse,
  TreeView_FullExpand,
  TreeView_GetNodeAt,
  TreeView_IsEditing,
  TreeView_LoadFromFile,
  TreeView_LoadFromStream,
  TreeView_SaveToFile,
  TreeView_SaveToStream,
  TreeView_Deselect,
  TreeView_Subselect,
  TreeView_ClearSelection,
  TreeView_FindNextToSelect,
  TreeView_CustomSort,
  TreeView_CanFocus,
  TreeView_FlipChildren,
  TreeView_Focused,
  TreeView_HandleAllocated,
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
  TreeView_GetBevelEdges,
  TreeView_SetBevelEdges,
  TreeView_GetBevelInner,
  TreeView_SetBevelInner,
  TreeView_GetBevelOuter,
  TreeView_SetBevelOuter,
  TreeView_GetBevelKind,
  TreeView_SetBevelKind,
  TreeView_GetBiDiMode,
  TreeView_SetBiDiMode,
  TreeView_GetBorderStyle,
  TreeView_SetBorderStyle,
  TreeView_GetBorderWidth,
  TreeView_SetBorderWidth,
  TreeView_GetChangeDelay,
  TreeView_SetChangeDelay,
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
  TreeView_GetParentCtl3D,
  TreeView_SetParentCtl3D,
  TreeView_GetParentDoubleBuffered,
  TreeView_SetParentDoubleBuffered,
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
  TreeView_GetStyleElements,
  TreeView_SetStyleElements,
  TreeView_SetOnAdvancedCustomDraw,
  TreeView_SetOnAdvancedCustomDrawItem,
  TreeView_SetOnChange,
  TreeView_SetOnClick,
  TreeView_SetOnCompare,
  TreeView_SetOnDblClick,
  TreeView_SetOnEnter,
  TreeView_SetOnExit,
  TreeView_SetOnGetImageIndex,
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
  TreeView_GetDropTarget,
  TreeView_SetDropTarget,
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
  TreeView_GetExplicitLeft,
  TreeView_GetExplicitTop,
  TreeView_GetExplicitWidth,
  TreeView_GetExplicitHeight,
  TreeView_GetParent,
  TreeView_SetParent,
  TreeView_GetAlignWithMargins,
  TreeView_SetAlignWithMargins,
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
  TreeView_GetMargins,
  TreeView_SetMargins,
  TreeView_GetCustomHint,
  TreeView_SetCustomHint,
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

