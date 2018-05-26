
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function CheckListBox_Create(AOwner: TComponent): TCheckListBox; extdecl;
begin
  Result :=  TCheckListBox.Create(AOwner);
end;

procedure CheckListBox_Free(AObj: TCheckListBox); extdecl;
begin
  AObj.Free;
end;

procedure CheckListBox_CheckAll(AObj: TCheckListBox; AState: TCheckBoxState; AllowGrayed: LongBool; AllowDisabled: LongBool); extdecl;
begin
  AObj.CheckAll(AState, AllowGrayed, AllowDisabled);
end;

procedure CheckListBox_Clear(AObj: TCheckListBox); extdecl;
begin
  AObj.Clear;
end;

procedure CheckListBox_ClearSelection(AObj: TCheckListBox); extdecl;
begin
  AObj.ClearSelection;
end;

procedure CheckListBox_SelectAll(AObj: TCheckListBox); extdecl;
begin
  AObj.SelectAll;
end;

function CheckListBox_CanFocus(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure CheckListBox_FlipChildren(AObj: TCheckListBox; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function CheckListBox_Focused(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure CheckListBox_Invalidate(AObj: TCheckListBox); extdecl;
begin
  AObj.Invalidate;
end;

procedure CheckListBox_Realign(AObj: TCheckListBox); extdecl;
begin
  AObj.Realign;
end;

procedure CheckListBox_Repaint(AObj: TCheckListBox); extdecl;
begin
  AObj.Repaint;
end;

procedure CheckListBox_ScaleBy(AObj: TCheckListBox; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure CheckListBox_SetBounds(AObj: TCheckListBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure CheckListBox_SetFocus(AObj: TCheckListBox); extdecl;
begin
  AObj.SetFocus;
end;

procedure CheckListBox_Update(AObj: TCheckListBox); extdecl;
begin
  AObj.Update;
end;

procedure CheckListBox_BringToFront(AObj: TCheckListBox); extdecl;
begin
  AObj.BringToFront;
end;

function CheckListBox_HasParent(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure CheckListBox_Hide(AObj: TCheckListBox); extdecl;
begin
  AObj.Hide;
end;

function CheckListBox_Perform(AObj: TCheckListBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure CheckListBox_Refresh(AObj: TCheckListBox); extdecl;
begin
  AObj.Refresh;
end;

procedure CheckListBox_SendToBack(AObj: TCheckListBox); extdecl;
begin
  AObj.SendToBack;
end;

procedure CheckListBox_Show(AObj: TCheckListBox); extdecl;
begin
  AObj.Show;
end;

function CheckListBox_GetTextBuf(AObj: TCheckListBox; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function CheckListBox_GetTextLen(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function CheckListBox_FindComponent(AObj: TCheckListBox; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function CheckListBox_GetNamePath(AObj: TCheckListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure CheckListBox_Assign(AObj: TCheckListBox; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function CheckListBox_ClassName(AObj: TCheckListBox): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function CheckListBox_Equals(AObj: TCheckListBox; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function CheckListBox_GetHashCode(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function CheckListBox_ToString(AObj: TCheckListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

procedure CheckListBox_SetOnClickCheck(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnClickCheck := nil;
    EventClass.Remove(AObj, geClickCheck);
    Exit;
  end;
  AObj.OnClickCheck := @EventClass.OnClickCheck;
  EventClass.Add(AObj, geClickCheck, AEventId);
end;

function CheckListBox_GetAlign(AObj: TCheckListBox): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure CheckListBox_SetAlign(AObj: TCheckListBox; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function CheckListBox_GetAllowGrayed(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.AllowGrayed;
end;

procedure CheckListBox_SetAllowGrayed(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.AllowGrayed := AValue;
end;

function CheckListBox_GetAnchors(AObj: TCheckListBox): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure CheckListBox_SetAnchors(AObj: TCheckListBox; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function CheckListBox_GetBiDiMode(AObj: TCheckListBox): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure CheckListBox_SetBiDiMode(AObj: TCheckListBox; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function CheckListBox_GetBorderStyle(AObj: TCheckListBox): TBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure CheckListBox_SetBorderStyle(AObj: TCheckListBox; AValue: TBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function CheckListBox_GetColor(AObj: TCheckListBox): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure CheckListBox_SetColor(AObj: TCheckListBox; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function CheckListBox_GetColumns(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.Columns;
end;

procedure CheckListBox_SetColumns(AObj: TCheckListBox; AValue: Integer); extdecl;
begin
  AObj.Columns := AValue;
end;

function CheckListBox_GetDoubleBuffered(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure CheckListBox_SetDoubleBuffered(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function CheckListBox_GetEnabled(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure CheckListBox_SetEnabled(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function CheckListBox_GetFont(AObj: TCheckListBox): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure CheckListBox_SetFont(AObj: TCheckListBox; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function CheckListBox_GetItemHeight(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.ItemHeight;
end;

procedure CheckListBox_SetItemHeight(AObj: TCheckListBox; AValue: Integer); extdecl;
begin
  AObj.ItemHeight := AValue;
end;

function CheckListBox_GetItems(AObj: TCheckListBox): TStrings; extdecl;
begin
  Result :=  AObj.Items;
end;

procedure CheckListBox_SetItems(AObj: TCheckListBox; AValue: TStrings); extdecl;
begin
  AObj.Items := AValue;
end;

function CheckListBox_GetParentColor(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure CheckListBox_SetParentColor(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function CheckListBox_GetParentFont(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure CheckListBox_SetParentFont(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function CheckListBox_GetParentShowHint(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure CheckListBox_SetParentShowHint(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function CheckListBox_GetPopupMenu(AObj: TCheckListBox): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure CheckListBox_SetPopupMenu(AObj: TCheckListBox; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function CheckListBox_GetShowHint(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure CheckListBox_SetShowHint(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function CheckListBox_GetSorted(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.Sorted;
end;

procedure CheckListBox_SetSorted(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.Sorted := AValue;
end;

function CheckListBox_GetStyle(AObj: TCheckListBox): TListBoxStyle; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure CheckListBox_SetStyle(AObj: TCheckListBox; AValue: TListBoxStyle); extdecl;
begin
  AObj.Style := AValue;
end;

function CheckListBox_GetTabOrder(AObj: TCheckListBox): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure CheckListBox_SetTabOrder(AObj: TCheckListBox; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function CheckListBox_GetTabStop(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure CheckListBox_SetTabStop(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function CheckListBox_GetVisible(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure CheckListBox_SetVisible(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure CheckListBox_SetOnClick(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnDblClick(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnEnter(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnExit(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnKeyDown(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnKeyPress(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnKeyUp(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnMouseDown(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnMouseEnter(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnMouseLeave(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnMouseMove(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

procedure CheckListBox_SetOnMouseUp(AObj: TCheckListBox; AEventId: NativeUInt); extdecl;
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

function CheckListBox_GetCanvas(AObj: TCheckListBox): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function CheckListBox_GetCount(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

function CheckListBox_GetMultiSelect(AObj: TCheckListBox): LongBool; extdecl;
begin
  Result :=  AObj.MultiSelect;
end;

procedure CheckListBox_SetMultiSelect(AObj: TCheckListBox; AValue: LongBool); extdecl;
begin
  AObj.MultiSelect := AValue;
end;

function CheckListBox_GetSelCount(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.SelCount;
end;

function CheckListBox_GetItemIndex(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.ItemIndex;
end;

procedure CheckListBox_SetItemIndex(AObj: TCheckListBox; AValue: Integer); extdecl;
begin
  AObj.ItemIndex := AValue;
end;

function CheckListBox_GetBrush(AObj: TCheckListBox): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function CheckListBox_GetControlCount(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function CheckListBox_GetHandle(AObj: TCheckListBox): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function CheckListBox_GetParentWindow(AObj: TCheckListBox): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure CheckListBox_SetParentWindow(AObj: TCheckListBox; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function CheckListBox_GetAction(AObj: TCheckListBox): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure CheckListBox_SetAction(AObj: TCheckListBox; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure CheckListBox_GetBoundsRect(AObj: TCheckListBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure CheckListBox_SetBoundsRect(AObj: TCheckListBox; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function CheckListBox_GetClientHeight(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure CheckListBox_SetClientHeight(AObj: TCheckListBox; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure CheckListBox_GetClientRect(AObj: TCheckListBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function CheckListBox_GetClientWidth(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure CheckListBox_SetClientWidth(AObj: TCheckListBox; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function CheckListBox_GetParent(AObj: TCheckListBox): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure CheckListBox_SetParent(AObj: TCheckListBox; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function CheckListBox_GetLeft(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure CheckListBox_SetLeft(AObj: TCheckListBox; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function CheckListBox_GetTop(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure CheckListBox_SetTop(AObj: TCheckListBox; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function CheckListBox_GetWidth(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure CheckListBox_SetWidth(AObj: TCheckListBox; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function CheckListBox_GetHeight(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure CheckListBox_SetHeight(AObj: TCheckListBox; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function CheckListBox_GetCursor(AObj: TCheckListBox): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure CheckListBox_SetCursor(AObj: TCheckListBox; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function CheckListBox_GetHint(AObj: TCheckListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure CheckListBox_SetHint(AObj: TCheckListBox; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function CheckListBox_GetComponentCount(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function CheckListBox_GetComponentIndex(AObj: TCheckListBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure CheckListBox_SetComponentIndex(AObj: TCheckListBox; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function CheckListBox_GetOwner(AObj: TCheckListBox): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function CheckListBox_GetName(AObj: TCheckListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure CheckListBox_SetName(AObj: TCheckListBox; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function CheckListBox_GetTag(AObj: TCheckListBox): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure CheckListBox_SetTag(AObj: TCheckListBox; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function CheckListBox_GetChecked(AObj: TCheckListBox; Index: Integer): LongBool; extdecl;
begin
  Result :=  AObj.Checked[Index];
end;

procedure CheckListBox_SetChecked(AObj: TCheckListBox; Index: Integer; AValue: LongBool); extdecl;
begin
  AObj.Checked[Index] := AValue;
end;

function CheckListBox_GetItemEnabled(AObj: TCheckListBox; Index: Integer): LongBool; extdecl;
begin
  Result :=  AObj.ItemEnabled[Index];
end;

procedure CheckListBox_SetItemEnabled(AObj: TCheckListBox; Index: Integer; AValue: LongBool); extdecl;
begin
  AObj.ItemEnabled[Index] := AValue;
end;

function CheckListBox_GetState(AObj: TCheckListBox; Index: Integer): TCheckBoxState; extdecl;
begin
  Result :=  AObj.State[Index];
end;

procedure CheckListBox_SetState(AObj: TCheckListBox; Index: Integer; AValue: TCheckBoxState); extdecl;
begin
  AObj.State[Index] := AValue;
end;

function CheckListBox_GetHeader(AObj: TCheckListBox; Index: Integer): LongBool; extdecl;
begin
  Result :=  AObj.Header[Index];
end;

procedure CheckListBox_SetHeader(AObj: TCheckListBox; Index: Integer; AValue: LongBool); extdecl;
begin
  AObj.Header[Index] := AValue;
end;

function CheckListBox_GetSelected(AObj: TCheckListBox; Index: Integer): LongBool; extdecl;
begin
  Result :=  AObj.Selected[Index];
end;

procedure CheckListBox_SetSelected(AObj: TCheckListBox; Index: Integer; AValue: LongBool); extdecl;
begin
  AObj.Selected[Index] := AValue;
end;

function CheckListBox_GetControls(AObj: TCheckListBox; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function CheckListBox_GetComponents(AObj: TCheckListBox; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  CheckListBox_Create,
  CheckListBox_Free,
  CheckListBox_CheckAll,
  CheckListBox_Clear,
  CheckListBox_ClearSelection,
  CheckListBox_SelectAll,
  CheckListBox_CanFocus,
  CheckListBox_FlipChildren,
  CheckListBox_Focused,
  CheckListBox_Invalidate,
  CheckListBox_Realign,
  CheckListBox_Repaint,
  CheckListBox_ScaleBy,
  CheckListBox_SetBounds,
  CheckListBox_SetFocus,
  CheckListBox_Update,
  CheckListBox_BringToFront,
  CheckListBox_HasParent,
  CheckListBox_Hide,
  CheckListBox_Perform,
  CheckListBox_Refresh,
  CheckListBox_SendToBack,
  CheckListBox_Show,
  CheckListBox_GetTextBuf,
  CheckListBox_GetTextLen,
  CheckListBox_FindComponent,
  CheckListBox_GetNamePath,
  CheckListBox_Assign,
  CheckListBox_ClassName,
  CheckListBox_Equals,
  CheckListBox_GetHashCode,
  CheckListBox_ToString,
  CheckListBox_SetOnClickCheck,
  CheckListBox_GetAlign,
  CheckListBox_SetAlign,
  CheckListBox_GetAllowGrayed,
  CheckListBox_SetAllowGrayed,
  CheckListBox_GetAnchors,
  CheckListBox_SetAnchors,
  CheckListBox_GetBiDiMode,
  CheckListBox_SetBiDiMode,
  CheckListBox_GetBorderStyle,
  CheckListBox_SetBorderStyle,
  CheckListBox_GetColor,
  CheckListBox_SetColor,
  CheckListBox_GetColumns,
  CheckListBox_SetColumns,
  CheckListBox_GetDoubleBuffered,
  CheckListBox_SetDoubleBuffered,
  CheckListBox_GetEnabled,
  CheckListBox_SetEnabled,
  CheckListBox_GetFont,
  CheckListBox_SetFont,
  CheckListBox_GetItemHeight,
  CheckListBox_SetItemHeight,
  CheckListBox_GetItems,
  CheckListBox_SetItems,
  CheckListBox_GetParentColor,
  CheckListBox_SetParentColor,
  CheckListBox_GetParentFont,
  CheckListBox_SetParentFont,
  CheckListBox_GetParentShowHint,
  CheckListBox_SetParentShowHint,
  CheckListBox_GetPopupMenu,
  CheckListBox_SetPopupMenu,
  CheckListBox_GetShowHint,
  CheckListBox_SetShowHint,
  CheckListBox_GetSorted,
  CheckListBox_SetSorted,
  CheckListBox_GetStyle,
  CheckListBox_SetStyle,
  CheckListBox_GetTabOrder,
  CheckListBox_SetTabOrder,
  CheckListBox_GetTabStop,
  CheckListBox_SetTabStop,
  CheckListBox_GetVisible,
  CheckListBox_SetVisible,
  CheckListBox_SetOnClick,
  CheckListBox_SetOnDblClick,
  CheckListBox_SetOnEnter,
  CheckListBox_SetOnExit,
  CheckListBox_SetOnKeyDown,
  CheckListBox_SetOnKeyPress,
  CheckListBox_SetOnKeyUp,
  CheckListBox_SetOnMouseDown,
  CheckListBox_SetOnMouseEnter,
  CheckListBox_SetOnMouseLeave,
  CheckListBox_SetOnMouseMove,
  CheckListBox_SetOnMouseUp,
  CheckListBox_GetCanvas,
  CheckListBox_GetCount,
  CheckListBox_GetMultiSelect,
  CheckListBox_SetMultiSelect,
  CheckListBox_GetSelCount,
  CheckListBox_GetItemIndex,
  CheckListBox_SetItemIndex,
  CheckListBox_GetBrush,
  CheckListBox_GetControlCount,
  CheckListBox_GetHandle,
  CheckListBox_GetParentWindow,
  CheckListBox_SetParentWindow,
  CheckListBox_GetAction,
  CheckListBox_SetAction,
  CheckListBox_GetBoundsRect,
  CheckListBox_SetBoundsRect,
  CheckListBox_GetClientHeight,
  CheckListBox_SetClientHeight,
  CheckListBox_GetClientRect,
  CheckListBox_GetClientWidth,
  CheckListBox_SetClientWidth,
  CheckListBox_GetParent,
  CheckListBox_SetParent,
  CheckListBox_GetLeft,
  CheckListBox_SetLeft,
  CheckListBox_GetTop,
  CheckListBox_SetTop,
  CheckListBox_GetWidth,
  CheckListBox_SetWidth,
  CheckListBox_GetHeight,
  CheckListBox_SetHeight,
  CheckListBox_GetCursor,
  CheckListBox_SetCursor,
  CheckListBox_GetHint,
  CheckListBox_SetHint,
  CheckListBox_GetComponentCount,
  CheckListBox_GetComponentIndex,
  CheckListBox_SetComponentIndex,
  CheckListBox_GetOwner,
  CheckListBox_GetName,
  CheckListBox_SetName,
  CheckListBox_GetTag,
  CheckListBox_SetTag,
  CheckListBox_GetChecked,
  CheckListBox_SetChecked,
  CheckListBox_GetItemEnabled,
  CheckListBox_SetItemEnabled,
  CheckListBox_GetState,
  CheckListBox_SetState,
  CheckListBox_GetHeader,
  CheckListBox_SetHeader,
  CheckListBox_GetSelected,
  CheckListBox_SetSelected,
  CheckListBox_GetControls,
  CheckListBox_GetComponents;

