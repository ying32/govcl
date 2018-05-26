
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ListBox_Create(AOwner: TComponent): TListBox; extdecl;
begin
  Result :=  TListBox.Create(AOwner);
end;

procedure ListBox_Free(AObj: TListBox); extdecl;
begin
  AObj.Free;
end;

procedure ListBox_Clear(AObj: TListBox); extdecl;
begin
  AObj.Clear;
end;

procedure ListBox_ClearSelection(AObj: TListBox); extdecl;
begin
  AObj.ClearSelection;
end;

procedure ListBox_SelectAll(AObj: TListBox); extdecl;
begin
  AObj.SelectAll;
end;

function ListBox_CanFocus(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure ListBox_FlipChildren(AObj: TListBox; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function ListBox_Focused(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure ListBox_Invalidate(AObj: TListBox); extdecl;
begin
  AObj.Invalidate;
end;

procedure ListBox_Realign(AObj: TListBox); extdecl;
begin
  AObj.Realign;
end;

procedure ListBox_Repaint(AObj: TListBox); extdecl;
begin
  AObj.Repaint;
end;

procedure ListBox_ScaleBy(AObj: TListBox; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure ListBox_SetBounds(AObj: TListBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ListBox_SetFocus(AObj: TListBox); extdecl;
begin
  AObj.SetFocus;
end;

procedure ListBox_Update(AObj: TListBox); extdecl;
begin
  AObj.Update;
end;

procedure ListBox_BringToFront(AObj: TListBox); extdecl;
begin
  AObj.BringToFront;
end;

function ListBox_HasParent(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ListBox_Hide(AObj: TListBox); extdecl;
begin
  AObj.Hide;
end;

function ListBox_Perform(AObj: TListBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ListBox_Refresh(AObj: TListBox); extdecl;
begin
  AObj.Refresh;
end;

procedure ListBox_SendToBack(AObj: TListBox); extdecl;
begin
  AObj.SendToBack;
end;

procedure ListBox_Show(AObj: TListBox); extdecl;
begin
  AObj.Show;
end;

function ListBox_GetTextBuf(AObj: TListBox; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ListBox_GetTextLen(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function ListBox_FindComponent(AObj: TListBox; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ListBox_GetNamePath(AObj: TListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ListBox_Assign(AObj: TListBox; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ListBox_ClassName(AObj: TListBox): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ListBox_Equals(AObj: TListBox; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ListBox_GetHashCode(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ListBox_ToString(AObj: TListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ListBox_GetStyle(AObj: TListBox): TListBoxStyle; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure ListBox_SetStyle(AObj: TListBox; AValue: TListBoxStyle); extdecl;
begin
  AObj.Style := AValue;
end;

function ListBox_GetAlign(AObj: TListBox): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure ListBox_SetAlign(AObj: TListBox; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function ListBox_GetAnchors(AObj: TListBox): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure ListBox_SetAnchors(AObj: TListBox; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function ListBox_GetBiDiMode(AObj: TListBox): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ListBox_SetBiDiMode(AObj: TListBox; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function ListBox_GetBorderStyle(AObj: TListBox): TBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure ListBox_SetBorderStyle(AObj: TListBox; AValue: TBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function ListBox_GetColor(AObj: TListBox): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure ListBox_SetColor(AObj: TListBox; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function ListBox_GetColumns(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.Columns;
end;

procedure ListBox_SetColumns(AObj: TListBox; AValue: Integer); extdecl;
begin
  AObj.Columns := AValue;
end;

function ListBox_GetDoubleBuffered(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ListBox_SetDoubleBuffered(AObj: TListBox; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function ListBox_GetEnabled(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure ListBox_SetEnabled(AObj: TListBox; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function ListBox_GetFont(AObj: TListBox): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure ListBox_SetFont(AObj: TListBox; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function ListBox_GetItemHeight(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.ItemHeight;
end;

procedure ListBox_SetItemHeight(AObj: TListBox; AValue: Integer); extdecl;
begin
  AObj.ItemHeight := AValue;
end;

function ListBox_GetItems(AObj: TListBox): TStrings; extdecl;
begin
  Result :=  AObj.Items;
end;

procedure ListBox_SetItems(AObj: TListBox; AValue: TStrings); extdecl;
begin
  AObj.Items := AValue;
end;

function ListBox_GetMultiSelect(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.MultiSelect;
end;

procedure ListBox_SetMultiSelect(AObj: TListBox; AValue: LongBool); extdecl;
begin
  AObj.MultiSelect := AValue;
end;

function ListBox_GetParentColor(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure ListBox_SetParentColor(AObj: TListBox; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function ListBox_GetParentFont(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure ListBox_SetParentFont(AObj: TListBox; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function ListBox_GetParentShowHint(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ListBox_SetParentShowHint(AObj: TListBox; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function ListBox_GetPopupMenu(AObj: TListBox): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ListBox_SetPopupMenu(AObj: TListBox; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function ListBox_GetShowHint(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure ListBox_SetShowHint(AObj: TListBox; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function ListBox_GetSorted(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.Sorted;
end;

procedure ListBox_SetSorted(AObj: TListBox; AValue: LongBool); extdecl;
begin
  AObj.Sorted := AValue;
end;

function ListBox_GetTabOrder(AObj: TListBox): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure ListBox_SetTabOrder(AObj: TListBox; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function ListBox_GetTabStop(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure ListBox_SetTabStop(AObj: TListBox; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function ListBox_GetVisible(AObj: TListBox): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure ListBox_SetVisible(AObj: TListBox; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure ListBox_SetOnClick(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnDblClick(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnDrawItem(AObj: TListBox; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnDrawItem := nil;
    EventClass.Remove(AObj, geListBoxDrawItem);
    Exit;
  end;
  AObj.OnDrawItem := @EventClass.ListBoxOnDrawItem;
  EventClass.Add(AObj, geListBoxDrawItem, AEventId);
end;

procedure ListBox_SetOnEnter(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnExit(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnKeyDown(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnKeyPress(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnKeyUp(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnMouseDown(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnMouseEnter(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnMouseLeave(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnMouseMove(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

procedure ListBox_SetOnMouseUp(AObj: TListBox; AEventId: NativeUInt); extdecl;
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

function ListBox_GetCanvas(AObj: TListBox): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function ListBox_GetCount(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

function ListBox_GetSelCount(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.SelCount;
end;

function ListBox_GetItemIndex(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.ItemIndex;
end;

procedure ListBox_SetItemIndex(AObj: TListBox; AValue: Integer); extdecl;
begin
  AObj.ItemIndex := AValue;
end;

function ListBox_GetBrush(AObj: TListBox): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function ListBox_GetControlCount(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function ListBox_GetHandle(AObj: TListBox): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function ListBox_GetParentWindow(AObj: TListBox): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ListBox_SetParentWindow(AObj: TListBox; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function ListBox_GetAction(AObj: TListBox): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure ListBox_SetAction(AObj: TListBox; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure ListBox_GetBoundsRect(AObj: TListBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ListBox_SetBoundsRect(AObj: TListBox; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function ListBox_GetClientHeight(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ListBox_SetClientHeight(AObj: TListBox; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure ListBox_GetClientRect(AObj: TListBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function ListBox_GetClientWidth(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ListBox_SetClientWidth(AObj: TListBox; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function ListBox_GetParent(AObj: TListBox): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure ListBox_SetParent(AObj: TListBox; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function ListBox_GetLeft(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ListBox_SetLeft(AObj: TListBox; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ListBox_GetTop(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ListBox_SetTop(AObj: TListBox; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ListBox_GetWidth(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ListBox_SetWidth(AObj: TListBox; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function ListBox_GetHeight(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ListBox_SetHeight(AObj: TListBox; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ListBox_GetCursor(AObj: TListBox): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure ListBox_SetCursor(AObj: TListBox; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function ListBox_GetHint(AObj: TListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ListBox_SetHint(AObj: TListBox; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function ListBox_GetComponentCount(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ListBox_GetComponentIndex(AObj: TListBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ListBox_SetComponentIndex(AObj: TListBox; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ListBox_GetOwner(AObj: TListBox): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ListBox_GetName(AObj: TListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ListBox_SetName(AObj: TListBox; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ListBox_GetTag(AObj: TListBox): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ListBox_SetTag(AObj: TListBox; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ListBox_GetSelected(AObj: TListBox; Index: Integer): LongBool; extdecl;
begin
  Result :=  AObj.Selected[Index];
end;

procedure ListBox_SetSelected(AObj: TListBox; Index: Integer; AValue: LongBool); extdecl;
begin
  AObj.Selected[Index] := AValue;
end;

function ListBox_GetControls(AObj: TListBox; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function ListBox_GetComponents(AObj: TListBox; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ListBox_Create,
  ListBox_Free,
  ListBox_Clear,
  ListBox_ClearSelection,
  ListBox_SelectAll,
  ListBox_CanFocus,
  ListBox_FlipChildren,
  ListBox_Focused,
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
  ListBox_GetAlign,
  ListBox_SetAlign,
  ListBox_GetAnchors,
  ListBox_SetAnchors,
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
  ListBox_GetVisible,
  ListBox_SetVisible,
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
  ListBox_GetParent,
  ListBox_SetParent,
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

