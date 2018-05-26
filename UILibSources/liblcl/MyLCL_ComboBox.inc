
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ComboBox_Create(AOwner: TComponent): TComboBox; extdecl;
begin
  Result :=  TComboBox.Create(AOwner);
end;

procedure ComboBox_Free(AObj: TComboBox); extdecl;
begin
  AObj.Free;
end;

procedure ComboBox_Clear(AObj: TComboBox); extdecl;
begin
  AObj.Clear;
end;

procedure ComboBox_ClearSelection(AObj: TComboBox); extdecl;
begin
  AObj.ClearSelection;
end;

function ComboBox_Focused(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure ComboBox_SelectAll(AObj: TComboBox); extdecl;
begin
  AObj.SelectAll;
end;

function ComboBox_CanFocus(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure ComboBox_FlipChildren(AObj: TComboBox; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

procedure ComboBox_Invalidate(AObj: TComboBox); extdecl;
begin
  AObj.Invalidate;
end;

procedure ComboBox_Realign(AObj: TComboBox); extdecl;
begin
  AObj.Realign;
end;

procedure ComboBox_Repaint(AObj: TComboBox); extdecl;
begin
  AObj.Repaint;
end;

procedure ComboBox_ScaleBy(AObj: TComboBox; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure ComboBox_SetBounds(AObj: TComboBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ComboBox_SetFocus(AObj: TComboBox); extdecl;
begin
  AObj.SetFocus;
end;

procedure ComboBox_Update(AObj: TComboBox); extdecl;
begin
  AObj.Update;
end;

procedure ComboBox_BringToFront(AObj: TComboBox); extdecl;
begin
  AObj.BringToFront;
end;

function ComboBox_HasParent(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ComboBox_Hide(AObj: TComboBox); extdecl;
begin
  AObj.Hide;
end;

function ComboBox_Perform(AObj: TComboBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ComboBox_Refresh(AObj: TComboBox); extdecl;
begin
  AObj.Refresh;
end;

procedure ComboBox_SendToBack(AObj: TComboBox); extdecl;
begin
  AObj.SendToBack;
end;

procedure ComboBox_Show(AObj: TComboBox); extdecl;
begin
  AObj.Show;
end;

function ComboBox_GetTextBuf(AObj: TComboBox; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ComboBox_GetTextLen(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function ComboBox_FindComponent(AObj: TComboBox; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ComboBox_GetNamePath(AObj: TComboBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ComboBox_Assign(AObj: TComboBox; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ComboBox_ClassName(AObj: TComboBox): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ComboBox_Equals(AObj: TComboBox; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ComboBox_GetHashCode(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ComboBox_ToString(AObj: TComboBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ComboBox_GetAlign(AObj: TComboBox): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure ComboBox_SetAlign(AObj: TComboBox; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function ComboBox_GetAutoComplete(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.AutoComplete;
end;

procedure ComboBox_SetAutoComplete(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.AutoComplete := AValue;
end;

function ComboBox_GetAutoDropDown(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.AutoDropDown;
end;

procedure ComboBox_SetAutoDropDown(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.AutoDropDown := AValue;
end;

function ComboBox_GetStyle(AObj: TComboBox): TComboBoxStyle; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure ComboBox_SetStyle(AObj: TComboBox; AValue: TComboBoxStyle); extdecl;
begin
  AObj.Style := AValue;
end;

function ComboBox_GetAnchors(AObj: TComboBox): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure ComboBox_SetAnchors(AObj: TComboBox; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function ComboBox_GetBiDiMode(AObj: TComboBox): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ComboBox_SetBiDiMode(AObj: TComboBox; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function ComboBox_GetColor(AObj: TComboBox): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure ComboBox_SetColor(AObj: TComboBox; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function ComboBox_GetDoubleBuffered(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ComboBox_SetDoubleBuffered(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function ComboBox_GetDropDownCount(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.DropDownCount;
end;

procedure ComboBox_SetDropDownCount(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.DropDownCount := AValue;
end;

function ComboBox_GetEnabled(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure ComboBox_SetEnabled(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function ComboBox_GetFont(AObj: TComboBox): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure ComboBox_SetFont(AObj: TComboBox; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function ComboBox_GetItemHeight(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.ItemHeight;
end;

procedure ComboBox_SetItemHeight(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.ItemHeight := AValue;
end;

function ComboBox_GetItemIndex(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.ItemIndex;
end;

procedure ComboBox_SetItemIndex(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.ItemIndex := AValue;
end;

function ComboBox_GetMaxLength(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.MaxLength;
end;

procedure ComboBox_SetMaxLength(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.MaxLength := AValue;
end;

function ComboBox_GetParentColor(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure ComboBox_SetParentColor(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function ComboBox_GetParentFont(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure ComboBox_SetParentFont(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function ComboBox_GetParentShowHint(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ComboBox_SetParentShowHint(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function ComboBox_GetPopupMenu(AObj: TComboBox): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ComboBox_SetPopupMenu(AObj: TComboBox; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function ComboBox_GetShowHint(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure ComboBox_SetShowHint(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function ComboBox_GetSorted(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.Sorted;
end;

procedure ComboBox_SetSorted(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.Sorted := AValue;
end;

function ComboBox_GetTabOrder(AObj: TComboBox): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure ComboBox_SetTabOrder(AObj: TComboBox; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function ComboBox_GetTabStop(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure ComboBox_SetTabStop(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function ComboBox_GetText(AObj: TComboBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Text);
end;

procedure ComboBox_SetText(AObj: TComboBox; AValue: PChar); extdecl;
begin
  AObj.Text := AValue;
end;

function ComboBox_GetVisible(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure ComboBox_SetVisible(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure ComboBox_SetOnChange(AObj: TComboBox; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    EventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := @EventClass.OnChange;
  EventClass.Add(AObj, geChange, AEventId);
end;

procedure ComboBox_SetOnClick(AObj: TComboBox; AEventId: NativeUInt); extdecl;
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

procedure ComboBox_SetOnDblClick(AObj: TComboBox; AEventId: NativeUInt); extdecl;
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

procedure ComboBox_SetOnDrawItem(AObj: TComboBox; AEventId: NativeUInt); extdecl;
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

procedure ComboBox_SetOnEnter(AObj: TComboBox; AEventId: NativeUInt); extdecl;
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

procedure ComboBox_SetOnExit(AObj: TComboBox; AEventId: NativeUInt); extdecl;
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

procedure ComboBox_SetOnKeyDown(AObj: TComboBox; AEventId: NativeUInt); extdecl;
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

procedure ComboBox_SetOnKeyPress(AObj: TComboBox; AEventId: NativeUInt); extdecl;
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

procedure ComboBox_SetOnKeyUp(AObj: TComboBox; AEventId: NativeUInt); extdecl;
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

procedure ComboBox_SetOnMouseEnter(AObj: TComboBox; AEventId: NativeUInt); extdecl;
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

procedure ComboBox_SetOnMouseLeave(AObj: TComboBox; AEventId: NativeUInt); extdecl;
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

function ComboBox_GetItems(AObj: TComboBox): TStrings; extdecl;
begin
  Result :=  AObj.Items;
end;

procedure ComboBox_SetItems(AObj: TComboBox; AValue: TStrings); extdecl;
begin
  AObj.Items := AValue;
end;

function ComboBox_GetSelText(AObj: TComboBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure ComboBox_SetSelText(AObj: TComboBox; AValue: PChar); extdecl;
begin
  AObj.SelText := AValue;
end;

function ComboBox_GetCanvas(AObj: TComboBox): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function ComboBox_GetDroppedDown(AObj: TComboBox): LongBool; extdecl;
begin
  Result :=  AObj.DroppedDown;
end;

procedure ComboBox_SetDroppedDown(AObj: TComboBox; AValue: LongBool); extdecl;
begin
  AObj.DroppedDown := AValue;
end;

function ComboBox_GetSelLength(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.SelLength;
end;

procedure ComboBox_SetSelLength(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.SelLength := AValue;
end;

function ComboBox_GetSelStart(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.SelStart;
end;

procedure ComboBox_SetSelStart(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.SelStart := AValue;
end;

function ComboBox_GetBrush(AObj: TComboBox): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function ComboBox_GetControlCount(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function ComboBox_GetHandle(AObj: TComboBox): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function ComboBox_GetParentWindow(AObj: TComboBox): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ComboBox_SetParentWindow(AObj: TComboBox; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function ComboBox_GetAction(AObj: TComboBox): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure ComboBox_SetAction(AObj: TComboBox; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure ComboBox_GetBoundsRect(AObj: TComboBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ComboBox_SetBoundsRect(AObj: TComboBox; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function ComboBox_GetClientHeight(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ComboBox_SetClientHeight(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure ComboBox_GetClientRect(AObj: TComboBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function ComboBox_GetClientWidth(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ComboBox_SetClientWidth(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function ComboBox_GetParent(AObj: TComboBox): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure ComboBox_SetParent(AObj: TComboBox; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function ComboBox_GetLeft(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ComboBox_SetLeft(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ComboBox_GetTop(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ComboBox_SetTop(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ComboBox_GetWidth(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ComboBox_SetWidth(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function ComboBox_GetHeight(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ComboBox_SetHeight(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ComboBox_GetCursor(AObj: TComboBox): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure ComboBox_SetCursor(AObj: TComboBox; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function ComboBox_GetHint(AObj: TComboBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ComboBox_SetHint(AObj: TComboBox; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function ComboBox_GetComponentCount(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ComboBox_GetComponentIndex(AObj: TComboBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ComboBox_SetComponentIndex(AObj: TComboBox; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ComboBox_GetOwner(AObj: TComboBox): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ComboBox_GetName(AObj: TComboBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ComboBox_SetName(AObj: TComboBox; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ComboBox_GetTag(AObj: TComboBox): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ComboBox_SetTag(AObj: TComboBox; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ComboBox_GetControls(AObj: TComboBox; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function ComboBox_GetComponents(AObj: TComboBox; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ComboBox_Create,
  ComboBox_Free,
  ComboBox_Clear,
  ComboBox_ClearSelection,
  ComboBox_Focused,
  ComboBox_SelectAll,
  ComboBox_CanFocus,
  ComboBox_FlipChildren,
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
  ComboBox_GetAutoDropDown,
  ComboBox_SetAutoDropDown,
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
  ComboBox_GetVisible,
  ComboBox_SetVisible,
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
  ComboBox_GetParent,
  ComboBox_SetParent,
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

