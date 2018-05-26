
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
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

procedure ListBox_Invalidate(AObj: TListBox); stdcall;
begin
  AObj.Invalidate;
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

procedure ListBox_BringToFront(AObj: TListBox); stdcall;
begin
  AObj.BringToFront;
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

function ListBox_ClassName(AObj: TListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
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
    TEventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, geClick, AEventId);
end;

procedure ListBox_SetOnDblClick(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnDrawItem(AObj: TListBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDrawItem := nil;
    TEventClass.Remove(AObj, geListBoxDrawItem);
    Exit;
  end;
  AObj.OnDrawItem := TEventClass.ListBoxOnDrawItem;
  TEventClass.Add(AObj, geListBoxDrawItem, AEventId);
end;

procedure ListBox_SetOnEnter(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnExit(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnKeyDown(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnKeyPress(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnKeyUp(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMouseDown(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMouseEnter(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMouseLeave(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMouseMove(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

procedure ListBox_SetOnMouseUp(AObj: TListBox; AEventId: NativeUInt); stdcall;
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

function ListBox_GetControls(AObj: TListBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ListBox_GetComponents(AObj: TListBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
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
  ListBox_FlipChildren,
  ListBox_Focused,
  ListBox_HandleAllocated,
  ListBox_Invalidate,
  ListBox_Realign,
  ListBox_Repaint,
  ListBox_ScaleBy,
  ListBox_SetBounds,
  ListBox_SetFocus,
  ListBox_Update,
  ListBox_BringToFront,
  ListBox_HasParent,
  ListBox_Hide,
  ListBox_Perform,
  ListBox_Refresh,
  ListBox_SendToBack,
  ListBox_Show,
  ListBox_GetTextBuf,
  ListBox_GetTextLen,
  ListBox_FindComponent,
  ListBox_GetNamePath,
  ListBox_Assign,
  ListBox_ClassName,
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
  ListBox_SetOnDblClick,
  ListBox_SetOnDrawItem,
  ListBox_SetOnEnter,
  ListBox_SetOnExit,
  ListBox_SetOnKeyDown,
  ListBox_SetOnKeyPress,
  ListBox_SetOnKeyUp,
  ListBox_SetOnMouseDown,
  ListBox_SetOnMouseEnter,
  ListBox_SetOnMouseLeave,
  ListBox_SetOnMouseMove,
  ListBox_SetOnMouseUp,
  ListBox_GetCanvas,
  ListBox_GetCount,
  ListBox_SetCount,
  ListBox_GetSelCount,
  ListBox_GetItemIndex,
  ListBox_SetItemIndex,
  ListBox_GetBrush,
  ListBox_GetControlCount,
  ListBox_GetHandle,
  ListBox_GetParentWindow,
  ListBox_SetParentWindow,
  ListBox_GetAction,
  ListBox_SetAction,
  ListBox_GetBoundsRect,
  ListBox_SetBoundsRect,
  ListBox_GetClientHeight,
  ListBox_SetClientHeight,
  ListBox_GetClientRect,
  ListBox_GetClientWidth,
  ListBox_SetClientWidth,
  ListBox_GetExplicitLeft,
  ListBox_GetExplicitTop,
  ListBox_GetExplicitWidth,
  ListBox_GetExplicitHeight,
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
  ListBox_GetControls,
  ListBox_GetComponents;

