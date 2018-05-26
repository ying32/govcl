
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function CheckListBox_Create(AOwner: TComponent): TCheckListBox; stdcall;
begin
  Result :=  TCheckListBox.Create(AOwner);
end;

procedure CheckListBox_Free(AObj: TCheckListBox); stdcall;
begin
  AObj.Free;
end;

procedure CheckListBox_CheckAll(AObj: TCheckListBox; AState: TCheckBoxState; AllowGrayed: LongBool; AllowDisabled: LongBool); stdcall;
begin
  AObj.CheckAll(AState, AllowGrayed, AllowDisabled);
end;

procedure CheckListBox_AddItem(AObj: TCheckListBox; Item: PChar; AObject: TObject); stdcall;
begin
  AObj.AddItem(Item, AObject);
end;

procedure CheckListBox_Clear(AObj: TCheckListBox); stdcall;
begin
  AObj.Clear;
end;

procedure CheckListBox_ClearSelection(AObj: TCheckListBox); stdcall;
begin
  AObj.ClearSelection;
end;

procedure CheckListBox_DeleteSelected(AObj: TCheckListBox); stdcall;
begin
  AObj.DeleteSelected;
end;

procedure CheckListBox_SelectAll(AObj: TCheckListBox); stdcall;
begin
  AObj.SelectAll;
end;

function CheckListBox_CanFocus(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure CheckListBox_FlipChildren(AObj: TCheckListBox; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function CheckListBox_Focused(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function CheckListBox_HandleAllocated(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure CheckListBox_Invalidate(AObj: TCheckListBox); stdcall;
begin
  AObj.Invalidate;
end;

procedure CheckListBox_Realign(AObj: TCheckListBox); stdcall;
begin
  AObj.Realign;
end;

procedure CheckListBox_Repaint(AObj: TCheckListBox); stdcall;
begin
  AObj.Repaint;
end;

procedure CheckListBox_ScaleBy(AObj: TCheckListBox; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure CheckListBox_SetBounds(AObj: TCheckListBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure CheckListBox_SetFocus(AObj: TCheckListBox); stdcall;
begin
  AObj.SetFocus;
end;

procedure CheckListBox_Update(AObj: TCheckListBox); stdcall;
begin
  AObj.Update;
end;

procedure CheckListBox_BringToFront(AObj: TCheckListBox); stdcall;
begin
  AObj.BringToFront;
end;

function CheckListBox_HasParent(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure CheckListBox_Hide(AObj: TCheckListBox); stdcall;
begin
  AObj.Hide;
end;

function CheckListBox_Perform(AObj: TCheckListBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure CheckListBox_Refresh(AObj: TCheckListBox); stdcall;
begin
  AObj.Refresh;
end;

procedure CheckListBox_SendToBack(AObj: TCheckListBox); stdcall;
begin
  AObj.SendToBack;
end;

procedure CheckListBox_Show(AObj: TCheckListBox); stdcall;
begin
  AObj.Show;
end;

function CheckListBox_GetTextBuf(AObj: TCheckListBox; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function CheckListBox_GetTextLen(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function CheckListBox_FindComponent(AObj: TCheckListBox; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function CheckListBox_GetNamePath(AObj: TCheckListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure CheckListBox_Assign(AObj: TCheckListBox; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function CheckListBox_ClassName(AObj: TCheckListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function CheckListBox_Equals(AObj: TCheckListBox; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function CheckListBox_GetHashCode(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function CheckListBox_ToString(AObj: TCheckListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

procedure CheckListBox_SetOnClickCheck(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClickCheck := nil;
    TEventClass.Remove(AObj, geClickCheck);
    Exit;
  end;
  AObj.OnClickCheck := TEventClass.OnClickCheck;
  TEventClass.Add(AObj, geClickCheck, AEventId);
end;

function CheckListBox_GetAlign(AObj: TCheckListBox): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure CheckListBox_SetAlign(AObj: TCheckListBox; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function CheckListBox_GetAllowGrayed(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.AllowGrayed;
end;

procedure CheckListBox_SetAllowGrayed(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.AllowGrayed := AValue;
end;

function CheckListBox_GetAnchors(AObj: TCheckListBox): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure CheckListBox_SetAnchors(AObj: TCheckListBox; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function CheckListBox_GetAutoComplete(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.AutoComplete;
end;

procedure CheckListBox_SetAutoComplete(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.AutoComplete := AValue;
end;

function CheckListBox_GetBevelEdges(AObj: TCheckListBox): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure CheckListBox_SetBevelEdges(AObj: TCheckListBox; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function CheckListBox_GetBevelInner(AObj: TCheckListBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure CheckListBox_SetBevelInner(AObj: TCheckListBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function CheckListBox_GetBevelOuter(AObj: TCheckListBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure CheckListBox_SetBevelOuter(AObj: TCheckListBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function CheckListBox_GetBevelKind(AObj: TCheckListBox): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure CheckListBox_SetBevelKind(AObj: TCheckListBox; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function CheckListBox_GetBiDiMode(AObj: TCheckListBox): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure CheckListBox_SetBiDiMode(AObj: TCheckListBox; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function CheckListBox_GetBorderStyle(AObj: TCheckListBox): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure CheckListBox_SetBorderStyle(AObj: TCheckListBox; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function CheckListBox_GetColor(AObj: TCheckListBox): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure CheckListBox_SetColor(AObj: TCheckListBox; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function CheckListBox_GetColumns(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.Columns;
end;

procedure CheckListBox_SetColumns(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.Columns := AValue;
end;

function CheckListBox_GetDoubleBuffered(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure CheckListBox_SetDoubleBuffered(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function CheckListBox_GetEnabled(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure CheckListBox_SetEnabled(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function CheckListBox_GetFlat(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.Flat;
end;

procedure CheckListBox_SetFlat(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.Flat := AValue;
end;

function CheckListBox_GetFont(AObj: TCheckListBox): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure CheckListBox_SetFont(AObj: TCheckListBox; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function CheckListBox_GetHeaderColor(AObj: TCheckListBox): TColor; stdcall;
begin
  Result :=  AObj.HeaderColor;
end;

procedure CheckListBox_SetHeaderColor(AObj: TCheckListBox; AValue: TColor); stdcall;
begin
  AObj.HeaderColor := AValue;
end;

function CheckListBox_GetHeaderBackgroundColor(AObj: TCheckListBox): TColor; stdcall;
begin
  Result :=  AObj.HeaderBackgroundColor;
end;

procedure CheckListBox_SetHeaderBackgroundColor(AObj: TCheckListBox; AValue: TColor); stdcall;
begin
  AObj.HeaderBackgroundColor := AValue;
end;

function CheckListBox_GetItemHeight(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ItemHeight;
end;

procedure CheckListBox_SetItemHeight(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.ItemHeight := AValue;
end;

function CheckListBox_GetItems(AObj: TCheckListBox): TStrings; stdcall;
begin
  Result :=  AObj.Items;
end;

procedure CheckListBox_SetItems(AObj: TCheckListBox; AValue: TStrings); stdcall;
begin
  AObj.Items := AValue;
end;

function CheckListBox_GetParentColor(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure CheckListBox_SetParentColor(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function CheckListBox_GetParentCtl3D(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure CheckListBox_SetParentCtl3D(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function CheckListBox_GetParentDoubleBuffered(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure CheckListBox_SetParentDoubleBuffered(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function CheckListBox_GetParentFont(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure CheckListBox_SetParentFont(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function CheckListBox_GetParentShowHint(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure CheckListBox_SetParentShowHint(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function CheckListBox_GetPopupMenu(AObj: TCheckListBox): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure CheckListBox_SetPopupMenu(AObj: TCheckListBox; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function CheckListBox_GetShowHint(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure CheckListBox_SetShowHint(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function CheckListBox_GetSorted(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.Sorted;
end;

procedure CheckListBox_SetSorted(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.Sorted := AValue;
end;

function CheckListBox_GetStyle(AObj: TCheckListBox): TListBoxStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure CheckListBox_SetStyle(AObj: TCheckListBox; AValue: TListBoxStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function CheckListBox_GetTabOrder(AObj: TCheckListBox): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure CheckListBox_SetTabOrder(AObj: TCheckListBox; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function CheckListBox_GetTabStop(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure CheckListBox_SetTabStop(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function CheckListBox_GetTabWidth(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.TabWidth;
end;

procedure CheckListBox_SetTabWidth(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.TabWidth := AValue;
end;

function CheckListBox_GetVisible(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure CheckListBox_SetVisible(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function CheckListBox_GetStyleElements(AObj: TCheckListBox): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure CheckListBox_SetStyleElements(AObj: TCheckListBox; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure CheckListBox_SetOnClick(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnDblClick(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnEnter(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnExit(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnKeyDown(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnKeyPress(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnKeyUp(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnMouseDown(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnMouseEnter(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnMouseLeave(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnMouseMove(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

procedure CheckListBox_SetOnMouseUp(AObj: TCheckListBox; AEventId: NativeUInt); stdcall;
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

function CheckListBox_GetAutoCompleteDelay(AObj: TCheckListBox): Cardinal; stdcall;
begin
  Result :=  AObj.AutoCompleteDelay;
end;

procedure CheckListBox_SetAutoCompleteDelay(AObj: TCheckListBox; AValue: Cardinal); stdcall;
begin
  AObj.AutoCompleteDelay := AValue;
end;

function CheckListBox_GetCanvas(AObj: TCheckListBox): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function CheckListBox_GetCount(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

procedure CheckListBox_SetCount(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.Count := AValue;
end;

function CheckListBox_GetMultiSelect(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.MultiSelect;
end;

procedure CheckListBox_SetMultiSelect(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.MultiSelect := AValue;
end;

function CheckListBox_GetSelCount(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.SelCount;
end;

function CheckListBox_GetItemIndex(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ItemIndex;
end;

procedure CheckListBox_SetItemIndex(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.ItemIndex := AValue;
end;

function CheckListBox_GetBrush(AObj: TCheckListBox): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function CheckListBox_GetControlCount(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function CheckListBox_GetHandle(AObj: TCheckListBox): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function CheckListBox_GetParentWindow(AObj: TCheckListBox): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure CheckListBox_SetParentWindow(AObj: TCheckListBox; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function CheckListBox_GetAction(AObj: TCheckListBox): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure CheckListBox_SetAction(AObj: TCheckListBox; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure CheckListBox_GetBoundsRect(AObj: TCheckListBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure CheckListBox_SetBoundsRect(AObj: TCheckListBox; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function CheckListBox_GetClientHeight(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure CheckListBox_SetClientHeight(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure CheckListBox_GetClientRect(AObj: TCheckListBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function CheckListBox_GetClientWidth(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure CheckListBox_SetClientWidth(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function CheckListBox_GetExplicitLeft(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function CheckListBox_GetExplicitTop(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function CheckListBox_GetExplicitWidth(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function CheckListBox_GetExplicitHeight(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function CheckListBox_GetParent(AObj: TCheckListBox): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure CheckListBox_SetParent(AObj: TCheckListBox; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function CheckListBox_GetAlignWithMargins(AObj: TCheckListBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure CheckListBox_SetAlignWithMargins(AObj: TCheckListBox; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function CheckListBox_GetLeft(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure CheckListBox_SetLeft(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function CheckListBox_GetTop(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure CheckListBox_SetTop(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function CheckListBox_GetWidth(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure CheckListBox_SetWidth(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function CheckListBox_GetHeight(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure CheckListBox_SetHeight(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function CheckListBox_GetCursor(AObj: TCheckListBox): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure CheckListBox_SetCursor(AObj: TCheckListBox; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function CheckListBox_GetHint(AObj: TCheckListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure CheckListBox_SetHint(AObj: TCheckListBox; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function CheckListBox_GetMargins(AObj: TCheckListBox): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure CheckListBox_SetMargins(AObj: TCheckListBox; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function CheckListBox_GetCustomHint(AObj: TCheckListBox): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure CheckListBox_SetCustomHint(AObj: TCheckListBox; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function CheckListBox_GetComponentCount(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function CheckListBox_GetComponentIndex(AObj: TCheckListBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure CheckListBox_SetComponentIndex(AObj: TCheckListBox; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function CheckListBox_GetOwner(AObj: TCheckListBox): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function CheckListBox_GetName(AObj: TCheckListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure CheckListBox_SetName(AObj: TCheckListBox; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function CheckListBox_GetTag(AObj: TCheckListBox): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure CheckListBox_SetTag(AObj: TCheckListBox; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function CheckListBox_GetChecked(AObj: TCheckListBox; Index: Integer): LongBool; stdcall;
begin
  Result :=  AObj.Checked[Index];
end;

procedure CheckListBox_SetChecked(AObj: TCheckListBox; Index: Integer; AValue: LongBool); stdcall;
begin
  AObj.Checked[Index] := AValue;
end;

function CheckListBox_GetItemEnabled(AObj: TCheckListBox; Index: Integer): LongBool; stdcall;
begin
  Result :=  AObj.ItemEnabled[Index];
end;

procedure CheckListBox_SetItemEnabled(AObj: TCheckListBox; Index: Integer; AValue: LongBool); stdcall;
begin
  AObj.ItemEnabled[Index] := AValue;
end;

function CheckListBox_GetState(AObj: TCheckListBox; Index: Integer): TCheckBoxState; stdcall;
begin
  Result :=  AObj.State[Index];
end;

procedure CheckListBox_SetState(AObj: TCheckListBox; Index: Integer; AValue: TCheckBoxState); stdcall;
begin
  AObj.State[Index] := AValue;
end;

function CheckListBox_GetHeader(AObj: TCheckListBox; Index: Integer): LongBool; stdcall;
begin
  Result :=  AObj.Header[Index];
end;

procedure CheckListBox_SetHeader(AObj: TCheckListBox; Index: Integer; AValue: LongBool); stdcall;
begin
  AObj.Header[Index] := AValue;
end;

function CheckListBox_GetSelected(AObj: TCheckListBox; Index: Integer): LongBool; stdcall;
begin
  Result :=  AObj.Selected[Index];
end;

procedure CheckListBox_SetSelected(AObj: TCheckListBox; Index: Integer; AValue: LongBool); stdcall;
begin
  AObj.Selected[Index] := AValue;
end;

function CheckListBox_GetControls(AObj: TCheckListBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function CheckListBox_GetComponents(AObj: TCheckListBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  CheckListBox_Create,
  CheckListBox_Free,
  CheckListBox_CheckAll,
  CheckListBox_AddItem,
  CheckListBox_Clear,
  CheckListBox_ClearSelection,
  CheckListBox_DeleteSelected,
  CheckListBox_SelectAll,
  CheckListBox_CanFocus,
  CheckListBox_FlipChildren,
  CheckListBox_Focused,
  CheckListBox_HandleAllocated,
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
  CheckListBox_GetAutoComplete,
  CheckListBox_SetAutoComplete,
  CheckListBox_GetBevelEdges,
  CheckListBox_SetBevelEdges,
  CheckListBox_GetBevelInner,
  CheckListBox_SetBevelInner,
  CheckListBox_GetBevelOuter,
  CheckListBox_SetBevelOuter,
  CheckListBox_GetBevelKind,
  CheckListBox_SetBevelKind,
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
  CheckListBox_GetFlat,
  CheckListBox_SetFlat,
  CheckListBox_GetFont,
  CheckListBox_SetFont,
  CheckListBox_GetHeaderColor,
  CheckListBox_SetHeaderColor,
  CheckListBox_GetHeaderBackgroundColor,
  CheckListBox_SetHeaderBackgroundColor,
  CheckListBox_GetItemHeight,
  CheckListBox_SetItemHeight,
  CheckListBox_GetItems,
  CheckListBox_SetItems,
  CheckListBox_GetParentColor,
  CheckListBox_SetParentColor,
  CheckListBox_GetParentCtl3D,
  CheckListBox_SetParentCtl3D,
  CheckListBox_GetParentDoubleBuffered,
  CheckListBox_SetParentDoubleBuffered,
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
  CheckListBox_GetTabWidth,
  CheckListBox_SetTabWidth,
  CheckListBox_GetVisible,
  CheckListBox_SetVisible,
  CheckListBox_GetStyleElements,
  CheckListBox_SetStyleElements,
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
  CheckListBox_GetAutoCompleteDelay,
  CheckListBox_SetAutoCompleteDelay,
  CheckListBox_GetCanvas,
  CheckListBox_GetCount,
  CheckListBox_SetCount,
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
  CheckListBox_GetExplicitLeft,
  CheckListBox_GetExplicitTop,
  CheckListBox_GetExplicitWidth,
  CheckListBox_GetExplicitHeight,
  CheckListBox_GetParent,
  CheckListBox_SetParent,
  CheckListBox_GetAlignWithMargins,
  CheckListBox_SetAlignWithMargins,
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
  CheckListBox_GetMargins,
  CheckListBox_SetMargins,
  CheckListBox_GetCustomHint,
  CheckListBox_SetCustomHint,
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

