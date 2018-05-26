
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ComboBox_Create(AOwner: TComponent): TComboBox; stdcall;
begin
  Result :=  TComboBox.Create(AOwner);
end;

procedure ComboBox_Free(AObj: TComboBox); stdcall;
begin
  AObj.Free;
end;

procedure ComboBox_AddItem(AObj: TComboBox; Item: PChar; AObject: TObject); stdcall;
begin
  AObj.AddItem(Item, AObject);
end;

procedure ComboBox_Clear(AObj: TComboBox); stdcall;
begin
  AObj.Clear;
end;

procedure ComboBox_ClearSelection(AObj: TComboBox); stdcall;
begin
  AObj.ClearSelection;
end;

procedure ComboBox_DeleteSelected(AObj: TComboBox); stdcall;
begin
  AObj.DeleteSelected;
end;

function ComboBox_Focused(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

procedure ComboBox_SelectAll(AObj: TComboBox); stdcall;
begin
  AObj.SelectAll;
end;

function ComboBox_CanFocus(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure ComboBox_FlipChildren(AObj: TComboBox; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function ComboBox_HandleAllocated(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ComboBox_Invalidate(AObj: TComboBox); stdcall;
begin
  AObj.Invalidate;
end;

procedure ComboBox_Realign(AObj: TComboBox); stdcall;
begin
  AObj.Realign;
end;

procedure ComboBox_Repaint(AObj: TComboBox); stdcall;
begin
  AObj.Repaint;
end;

procedure ComboBox_ScaleBy(AObj: TComboBox; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure ComboBox_SetBounds(AObj: TComboBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ComboBox_SetFocus(AObj: TComboBox); stdcall;
begin
  AObj.SetFocus;
end;

procedure ComboBox_Update(AObj: TComboBox); stdcall;
begin
  AObj.Update;
end;

procedure ComboBox_BringToFront(AObj: TComboBox); stdcall;
begin
  AObj.BringToFront;
end;

function ComboBox_HasParent(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ComboBox_Hide(AObj: TComboBox); stdcall;
begin
  AObj.Hide;
end;

function ComboBox_Perform(AObj: TComboBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ComboBox_Refresh(AObj: TComboBox); stdcall;
begin
  AObj.Refresh;
end;

procedure ComboBox_SendToBack(AObj: TComboBox); stdcall;
begin
  AObj.SendToBack;
end;

procedure ComboBox_Show(AObj: TComboBox); stdcall;
begin
  AObj.Show;
end;

function ComboBox_GetTextBuf(AObj: TComboBox; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ComboBox_GetTextLen(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function ComboBox_FindComponent(AObj: TComboBox; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ComboBox_GetNamePath(AObj: TComboBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ComboBox_Assign(AObj: TComboBox; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ComboBox_ClassName(AObj: TComboBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ComboBox_Equals(AObj: TComboBox; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ComboBox_GetHashCode(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ComboBox_ToString(AObj: TComboBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ComboBox_GetAlign(AObj: TComboBox): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ComboBox_SetAlign(AObj: TComboBox; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ComboBox_GetAutoComplete(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.AutoComplete;
end;

procedure ComboBox_SetAutoComplete(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.AutoComplete := AValue;
end;

function ComboBox_GetAutoCompleteDelay(AObj: TComboBox): Cardinal; stdcall;
begin
  Result :=  AObj.AutoCompleteDelay;
end;

procedure ComboBox_SetAutoCompleteDelay(AObj: TComboBox; AValue: Cardinal); stdcall;
begin
  AObj.AutoCompleteDelay := AValue;
end;

function ComboBox_GetAutoDropDown(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.AutoDropDown;
end;

procedure ComboBox_SetAutoDropDown(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.AutoDropDown := AValue;
end;

function ComboBox_GetAutoCloseUp(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.AutoCloseUp;
end;

procedure ComboBox_SetAutoCloseUp(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.AutoCloseUp := AValue;
end;

function ComboBox_GetBevelEdges(AObj: TComboBox): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure ComboBox_SetBevelEdges(AObj: TComboBox; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function ComboBox_GetBevelInner(AObj: TComboBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure ComboBox_SetBevelInner(AObj: TComboBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function ComboBox_GetBevelKind(AObj: TComboBox): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure ComboBox_SetBevelKind(AObj: TComboBox; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function ComboBox_GetBevelOuter(AObj: TComboBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure ComboBox_SetBevelOuter(AObj: TComboBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function ComboBox_GetStyle(AObj: TComboBox): TComboBoxStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure ComboBox_SetStyle(AObj: TComboBox; AValue: TComboBoxStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function ComboBox_GetAnchors(AObj: TComboBox): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ComboBox_SetAnchors(AObj: TComboBox; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ComboBox_GetBiDiMode(AObj: TComboBox): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ComboBox_SetBiDiMode(AObj: TComboBox; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function ComboBox_GetColor(AObj: TComboBox): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure ComboBox_SetColor(AObj: TComboBox; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function ComboBox_GetDoubleBuffered(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ComboBox_SetDoubleBuffered(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function ComboBox_GetDropDownCount(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.DropDownCount;
end;

procedure ComboBox_SetDropDownCount(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.DropDownCount := AValue;
end;

function ComboBox_GetEnabled(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ComboBox_SetEnabled(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ComboBox_GetFont(AObj: TComboBox): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure ComboBox_SetFont(AObj: TComboBox; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function ComboBox_GetItemHeight(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ItemHeight;
end;

procedure ComboBox_SetItemHeight(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.ItemHeight := AValue;
end;

function ComboBox_GetItemIndex(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ItemIndex;
end;

procedure ComboBox_SetItemIndex(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.ItemIndex := AValue;
end;

function ComboBox_GetMaxLength(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.MaxLength;
end;

procedure ComboBox_SetMaxLength(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.MaxLength := AValue;
end;

function ComboBox_GetParentColor(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure ComboBox_SetParentColor(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function ComboBox_GetParentCtl3D(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure ComboBox_SetParentCtl3D(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function ComboBox_GetParentDoubleBuffered(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure ComboBox_SetParentDoubleBuffered(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function ComboBox_GetParentFont(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure ComboBox_SetParentFont(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function ComboBox_GetParentShowHint(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ComboBox_SetParentShowHint(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ComboBox_GetPopupMenu(AObj: TComboBox): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ComboBox_SetPopupMenu(AObj: TComboBox; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ComboBox_GetShowHint(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ComboBox_SetShowHint(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ComboBox_GetSorted(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.Sorted;
end;

procedure ComboBox_SetSorted(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.Sorted := AValue;
end;

function ComboBox_GetTabOrder(AObj: TComboBox): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure ComboBox_SetTabOrder(AObj: TComboBox; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function ComboBox_GetTabStop(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure ComboBox_SetTabStop(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function ComboBox_GetText(AObj: TComboBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Text);
end;

procedure ComboBox_SetText(AObj: TComboBox; AValue: PChar); stdcall;
begin
  AObj.Text := AValue;
end;

function ComboBox_GetTextHint(AObj: TComboBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.TextHint);
end;

procedure ComboBox_SetTextHint(AObj: TComboBox; AValue: PChar); stdcall;
begin
  AObj.TextHint := AValue;
end;

function ComboBox_GetVisible(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ComboBox_SetVisible(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function ComboBox_GetStyleElements(AObj: TComboBox): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ComboBox_SetStyleElements(AObj: TComboBox; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure ComboBox_SetOnChange(AObj: TComboBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, geChange, AEventId);
end;

procedure ComboBox_SetOnClick(AObj: TComboBox; AEventId: NativeUInt); stdcall;
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

procedure ComboBox_SetOnDblClick(AObj: TComboBox; AEventId: NativeUInt); stdcall;
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

procedure ComboBox_SetOnDrawItem(AObj: TComboBox; AEventId: NativeUInt); stdcall;
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

procedure ComboBox_SetOnEnter(AObj: TComboBox; AEventId: NativeUInt); stdcall;
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

procedure ComboBox_SetOnExit(AObj: TComboBox; AEventId: NativeUInt); stdcall;
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

procedure ComboBox_SetOnKeyDown(AObj: TComboBox; AEventId: NativeUInt); stdcall;
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

procedure ComboBox_SetOnKeyPress(AObj: TComboBox; AEventId: NativeUInt); stdcall;
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

procedure ComboBox_SetOnKeyUp(AObj: TComboBox; AEventId: NativeUInt); stdcall;
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

procedure ComboBox_SetOnMouseEnter(AObj: TComboBox; AEventId: NativeUInt); stdcall;
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

procedure ComboBox_SetOnMouseLeave(AObj: TComboBox; AEventId: NativeUInt); stdcall;
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

function ComboBox_GetItems(AObj: TComboBox): TStrings; stdcall;
begin
  Result :=  AObj.Items;
end;

procedure ComboBox_SetItems(AObj: TComboBox; AValue: TStrings); stdcall;
begin
  AObj.Items := AValue;
end;

function ComboBox_GetSelText(AObj: TComboBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure ComboBox_SetSelText(AObj: TComboBox; AValue: PChar); stdcall;
begin
  AObj.SelText := AValue;
end;

function ComboBox_GetCanvas(AObj: TComboBox): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function ComboBox_GetDroppedDown(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.DroppedDown;
end;

procedure ComboBox_SetDroppedDown(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.DroppedDown := AValue;
end;

function ComboBox_GetSelLength(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.SelLength;
end;

procedure ComboBox_SetSelLength(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.SelLength := AValue;
end;

function ComboBox_GetSelStart(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.SelStart;
end;

procedure ComboBox_SetSelStart(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.SelStart := AValue;
end;

function ComboBox_GetBrush(AObj: TComboBox): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function ComboBox_GetControlCount(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function ComboBox_GetHandle(AObj: TComboBox): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ComboBox_GetParentWindow(AObj: TComboBox): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ComboBox_SetParentWindow(AObj: TComboBox; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function ComboBox_GetAction(AObj: TComboBox): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ComboBox_SetAction(AObj: TComboBox; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure ComboBox_GetBoundsRect(AObj: TComboBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ComboBox_SetBoundsRect(AObj: TComboBox; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ComboBox_GetClientHeight(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ComboBox_SetClientHeight(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ComboBox_GetClientRect(AObj: TComboBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ComboBox_GetClientWidth(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ComboBox_SetClientWidth(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ComboBox_GetExplicitLeft(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ComboBox_GetExplicitTop(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ComboBox_GetExplicitWidth(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ComboBox_GetExplicitHeight(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ComboBox_GetParent(AObj: TComboBox): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ComboBox_SetParent(AObj: TComboBox; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ComboBox_GetAlignWithMargins(AObj: TComboBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ComboBox_SetAlignWithMargins(AObj: TComboBox; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ComboBox_GetLeft(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ComboBox_SetLeft(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ComboBox_GetTop(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ComboBox_SetTop(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ComboBox_GetWidth(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ComboBox_SetWidth(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ComboBox_GetHeight(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ComboBox_SetHeight(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ComboBox_GetCursor(AObj: TComboBox): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ComboBox_SetCursor(AObj: TComboBox; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ComboBox_GetHint(AObj: TComboBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ComboBox_SetHint(AObj: TComboBox; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ComboBox_GetMargins(AObj: TComboBox): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ComboBox_SetMargins(AObj: TComboBox; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ComboBox_GetCustomHint(AObj: TComboBox): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ComboBox_SetCustomHint(AObj: TComboBox; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ComboBox_GetComponentCount(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ComboBox_GetComponentIndex(AObj: TComboBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ComboBox_SetComponentIndex(AObj: TComboBox; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ComboBox_GetOwner(AObj: TComboBox): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ComboBox_GetName(AObj: TComboBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ComboBox_SetName(AObj: TComboBox; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ComboBox_GetTag(AObj: TComboBox): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ComboBox_SetTag(AObj: TComboBox; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ComboBox_GetControls(AObj: TComboBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ComboBox_GetComponents(AObj: TComboBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ComboBox_Create,
  ComboBox_Free,
  ComboBox_AddItem,
  ComboBox_Clear,
  ComboBox_ClearSelection,
  ComboBox_DeleteSelected,
  ComboBox_Focused,
  ComboBox_SelectAll,
  ComboBox_CanFocus,
  ComboBox_FlipChildren,
  ComboBox_HandleAllocated,
  ComboBox_Invalidate,
  ComboBox_Realign,
  ComboBox_Repaint,
  ComboBox_ScaleBy,
  ComboBox_SetBounds,
  ComboBox_SetFocus,
  ComboBox_Update,
  ComboBox_BringToFront,
  ComboBox_HasParent,
  ComboBox_Hide,
  ComboBox_Perform,
  ComboBox_Refresh,
  ComboBox_SendToBack,
  ComboBox_Show,
  ComboBox_GetTextBuf,
  ComboBox_GetTextLen,
  ComboBox_FindComponent,
  ComboBox_GetNamePath,
  ComboBox_Assign,
  ComboBox_ClassName,
  ComboBox_Equals,
  ComboBox_GetHashCode,
  ComboBox_ToString,
  ComboBox_GetAlign,
  ComboBox_SetAlign,
  ComboBox_GetAutoComplete,
  ComboBox_SetAutoComplete,
  ComboBox_GetAutoCompleteDelay,
  ComboBox_SetAutoCompleteDelay,
  ComboBox_GetAutoDropDown,
  ComboBox_SetAutoDropDown,
  ComboBox_GetAutoCloseUp,
  ComboBox_SetAutoCloseUp,
  ComboBox_GetBevelEdges,
  ComboBox_SetBevelEdges,
  ComboBox_GetBevelInner,
  ComboBox_SetBevelInner,
  ComboBox_GetBevelKind,
  ComboBox_SetBevelKind,
  ComboBox_GetBevelOuter,
  ComboBox_SetBevelOuter,
  ComboBox_GetStyle,
  ComboBox_SetStyle,
  ComboBox_GetAnchors,
  ComboBox_SetAnchors,
  ComboBox_GetBiDiMode,
  ComboBox_SetBiDiMode,
  ComboBox_GetColor,
  ComboBox_SetColor,
  ComboBox_GetDoubleBuffered,
  ComboBox_SetDoubleBuffered,
  ComboBox_GetDropDownCount,
  ComboBox_SetDropDownCount,
  ComboBox_GetEnabled,
  ComboBox_SetEnabled,
  ComboBox_GetFont,
  ComboBox_SetFont,
  ComboBox_GetItemHeight,
  ComboBox_SetItemHeight,
  ComboBox_GetItemIndex,
  ComboBox_SetItemIndex,
  ComboBox_GetMaxLength,
  ComboBox_SetMaxLength,
  ComboBox_GetParentColor,
  ComboBox_SetParentColor,
  ComboBox_GetParentCtl3D,
  ComboBox_SetParentCtl3D,
  ComboBox_GetParentDoubleBuffered,
  ComboBox_SetParentDoubleBuffered,
  ComboBox_GetParentFont,
  ComboBox_SetParentFont,
  ComboBox_GetParentShowHint,
  ComboBox_SetParentShowHint,
  ComboBox_GetPopupMenu,
  ComboBox_SetPopupMenu,
  ComboBox_GetShowHint,
  ComboBox_SetShowHint,
  ComboBox_GetSorted,
  ComboBox_SetSorted,
  ComboBox_GetTabOrder,
  ComboBox_SetTabOrder,
  ComboBox_GetTabStop,
  ComboBox_SetTabStop,
  ComboBox_GetText,
  ComboBox_SetText,
  ComboBox_GetTextHint,
  ComboBox_SetTextHint,
  ComboBox_GetVisible,
  ComboBox_SetVisible,
  ComboBox_GetStyleElements,
  ComboBox_SetStyleElements,
  ComboBox_SetOnChange,
  ComboBox_SetOnClick,
  ComboBox_SetOnDblClick,
  ComboBox_SetOnDrawItem,
  ComboBox_SetOnEnter,
  ComboBox_SetOnExit,
  ComboBox_SetOnKeyDown,
  ComboBox_SetOnKeyPress,
  ComboBox_SetOnKeyUp,
  ComboBox_SetOnMouseEnter,
  ComboBox_SetOnMouseLeave,
  ComboBox_GetItems,
  ComboBox_SetItems,
  ComboBox_GetSelText,
  ComboBox_SetSelText,
  ComboBox_GetCanvas,
  ComboBox_GetDroppedDown,
  ComboBox_SetDroppedDown,
  ComboBox_GetSelLength,
  ComboBox_SetSelLength,
  ComboBox_GetSelStart,
  ComboBox_SetSelStart,
  ComboBox_GetBrush,
  ComboBox_GetControlCount,
  ComboBox_GetHandle,
  ComboBox_GetParentWindow,
  ComboBox_SetParentWindow,
  ComboBox_GetAction,
  ComboBox_SetAction,
  ComboBox_GetBoundsRect,
  ComboBox_SetBoundsRect,
  ComboBox_GetClientHeight,
  ComboBox_SetClientHeight,
  ComboBox_GetClientRect,
  ComboBox_GetClientWidth,
  ComboBox_SetClientWidth,
  ComboBox_GetExplicitLeft,
  ComboBox_GetExplicitTop,
  ComboBox_GetExplicitWidth,
  ComboBox_GetExplicitHeight,
  ComboBox_GetParent,
  ComboBox_SetParent,
  ComboBox_GetAlignWithMargins,
  ComboBox_SetAlignWithMargins,
  ComboBox_GetLeft,
  ComboBox_SetLeft,
  ComboBox_GetTop,
  ComboBox_SetTop,
  ComboBox_GetWidth,
  ComboBox_SetWidth,
  ComboBox_GetHeight,
  ComboBox_SetHeight,
  ComboBox_GetCursor,
  ComboBox_SetCursor,
  ComboBox_GetHint,
  ComboBox_SetHint,
  ComboBox_GetMargins,
  ComboBox_SetMargins,
  ComboBox_GetCustomHint,
  ComboBox_SetCustomHint,
  ComboBox_GetComponentCount,
  ComboBox_GetComponentIndex,
  ComboBox_SetComponentIndex,
  ComboBox_GetOwner,
  ComboBox_GetName,
  ComboBox_SetName,
  ComboBox_GetTag,
  ComboBox_SetTag,
  ComboBox_GetControls,
  ComboBox_GetComponents;

