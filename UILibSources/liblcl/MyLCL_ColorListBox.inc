
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ColorListBox_Create(AOwner: TComponent): TColorListBox; extdecl;
begin
  Result :=  TColorListBox.Create(AOwner);
end;

procedure ColorListBox_Free(AObj: TColorListBox); extdecl;
begin
  AObj.Free;
end;

procedure ColorListBox_Clear(AObj: TColorListBox); extdecl;
begin
  AObj.Clear;
end;

procedure ColorListBox_ClearSelection(AObj: TColorListBox); extdecl;
begin
  AObj.ClearSelection;
end;

procedure ColorListBox_SelectAll(AObj: TColorListBox); extdecl;
begin
  AObj.SelectAll;
end;

function ColorListBox_CanFocus(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure ColorListBox_FlipChildren(AObj: TColorListBox; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function ColorListBox_Focused(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure ColorListBox_Invalidate(AObj: TColorListBox); extdecl;
begin
  AObj.Invalidate;
end;

procedure ColorListBox_Realign(AObj: TColorListBox); extdecl;
begin
  AObj.Realign;
end;

procedure ColorListBox_Repaint(AObj: TColorListBox); extdecl;
begin
  AObj.Repaint;
end;

procedure ColorListBox_ScaleBy(AObj: TColorListBox; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure ColorListBox_SetBounds(AObj: TColorListBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ColorListBox_SetFocus(AObj: TColorListBox); extdecl;
begin
  AObj.SetFocus;
end;

procedure ColorListBox_Update(AObj: TColorListBox); extdecl;
begin
  AObj.Update;
end;

procedure ColorListBox_BringToFront(AObj: TColorListBox); extdecl;
begin
  AObj.BringToFront;
end;

function ColorListBox_HasParent(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ColorListBox_Hide(AObj: TColorListBox); extdecl;
begin
  AObj.Hide;
end;

function ColorListBox_Perform(AObj: TColorListBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ColorListBox_Refresh(AObj: TColorListBox); extdecl;
begin
  AObj.Refresh;
end;

procedure ColorListBox_SendToBack(AObj: TColorListBox); extdecl;
begin
  AObj.SendToBack;
end;

procedure ColorListBox_Show(AObj: TColorListBox); extdecl;
begin
  AObj.Show;
end;

function ColorListBox_GetTextBuf(AObj: TColorListBox; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ColorListBox_GetTextLen(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function ColorListBox_FindComponent(AObj: TColorListBox; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ColorListBox_GetNamePath(AObj: TColorListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ColorListBox_Assign(AObj: TColorListBox; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ColorListBox_ClassName(AObj: TColorListBox): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ColorListBox_Equals(AObj: TColorListBox; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ColorListBox_GetHashCode(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ColorListBox_ToString(AObj: TColorListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ColorListBox_GetAlign(AObj: TColorListBox): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure ColorListBox_SetAlign(AObj: TColorListBox; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function ColorListBox_GetDefaultColorColor(AObj: TColorListBox): TColor; extdecl;
begin
  Result :=  AObj.DefaultColorColor;
end;

procedure ColorListBox_SetDefaultColorColor(AObj: TColorListBox; AValue: TColor); extdecl;
begin
  AObj.DefaultColorColor := AValue;
end;

function ColorListBox_GetNoneColorColor(AObj: TColorListBox): TColor; extdecl;
begin
  Result :=  AObj.NoneColorColor;
end;

procedure ColorListBox_SetNoneColorColor(AObj: TColorListBox; AValue: TColor); extdecl;
begin
  AObj.NoneColorColor := AValue;
end;

function ColorListBox_GetSelected(AObj: TColorListBox): TColor; extdecl;
begin
  Result :=  AObj.Selected;
end;

procedure ColorListBox_SetSelected(AObj: TColorListBox; AValue: TColor); extdecl;
begin
  AObj.Selected := AValue;
end;

function ColorListBox_GetStyle(AObj: TColorListBox): TColorBoxStyle; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure ColorListBox_SetStyle(AObj: TColorListBox; AValue: TColorBoxStyle); extdecl;
begin
  AObj.Style := AValue;
end;

function ColorListBox_GetAnchors(AObj: TColorListBox): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure ColorListBox_SetAnchors(AObj: TColorListBox; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function ColorListBox_GetBiDiMode(AObj: TColorListBox): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ColorListBox_SetBiDiMode(AObj: TColorListBox; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function ColorListBox_GetColor(AObj: TColorListBox): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure ColorListBox_SetColor(AObj: TColorListBox; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function ColorListBox_GetDoubleBuffered(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ColorListBox_SetDoubleBuffered(AObj: TColorListBox; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function ColorListBox_GetEnabled(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure ColorListBox_SetEnabled(AObj: TColorListBox; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function ColorListBox_GetFont(AObj: TColorListBox): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure ColorListBox_SetFont(AObj: TColorListBox; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function ColorListBox_GetItemHeight(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.ItemHeight;
end;

procedure ColorListBox_SetItemHeight(AObj: TColorListBox; AValue: Integer); extdecl;
begin
  AObj.ItemHeight := AValue;
end;

function ColorListBox_GetParentColor(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure ColorListBox_SetParentColor(AObj: TColorListBox; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function ColorListBox_GetParentFont(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure ColorListBox_SetParentFont(AObj: TColorListBox; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function ColorListBox_GetParentShowHint(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ColorListBox_SetParentShowHint(AObj: TColorListBox; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function ColorListBox_GetPopupMenu(AObj: TColorListBox): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ColorListBox_SetPopupMenu(AObj: TColorListBox; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function ColorListBox_GetShowHint(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure ColorListBox_SetShowHint(AObj: TColorListBox; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function ColorListBox_GetTabOrder(AObj: TColorListBox): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure ColorListBox_SetTabOrder(AObj: TColorListBox; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function ColorListBox_GetTabStop(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure ColorListBox_SetTabStop(AObj: TColorListBox; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function ColorListBox_GetVisible(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure ColorListBox_SetVisible(AObj: TColorListBox; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure ColorListBox_SetOnClick(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnDblClick(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnEnter(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnExit(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnKeyDown(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnKeyPress(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnKeyUp(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnMouseDown(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnMouseEnter(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnMouseLeave(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnMouseMove(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

procedure ColorListBox_SetOnMouseUp(AObj: TColorListBox; AEventId: NativeUInt); extdecl;
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

function ColorListBox_GetCanvas(AObj: TColorListBox): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function ColorListBox_GetCount(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.Count;
end;

function ColorListBox_GetItems(AObj: TColorListBox): TStrings; extdecl;
begin
  Result :=  AObj.Items;
end;

procedure ColorListBox_SetItems(AObj: TColorListBox; AValue: TStrings); extdecl;
begin
  AObj.Items := AValue;
end;

function ColorListBox_GetMultiSelect(AObj: TColorListBox): LongBool; extdecl;
begin
  Result :=  AObj.MultiSelect;
end;

procedure ColorListBox_SetMultiSelect(AObj: TColorListBox; AValue: LongBool); extdecl;
begin
  AObj.MultiSelect := AValue;
end;

function ColorListBox_GetSelCount(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.SelCount;
end;

function ColorListBox_GetItemIndex(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.ItemIndex;
end;

procedure ColorListBox_SetItemIndex(AObj: TColorListBox; AValue: Integer); extdecl;
begin
  AObj.ItemIndex := AValue;
end;

function ColorListBox_GetBrush(AObj: TColorListBox): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function ColorListBox_GetControlCount(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function ColorListBox_GetHandle(AObj: TColorListBox): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function ColorListBox_GetParentWindow(AObj: TColorListBox): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ColorListBox_SetParentWindow(AObj: TColorListBox; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function ColorListBox_GetAction(AObj: TColorListBox): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure ColorListBox_SetAction(AObj: TColorListBox; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure ColorListBox_GetBoundsRect(AObj: TColorListBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ColorListBox_SetBoundsRect(AObj: TColorListBox; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function ColorListBox_GetClientHeight(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ColorListBox_SetClientHeight(AObj: TColorListBox; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure ColorListBox_GetClientRect(AObj: TColorListBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function ColorListBox_GetClientWidth(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ColorListBox_SetClientWidth(AObj: TColorListBox; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function ColorListBox_GetParent(AObj: TColorListBox): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure ColorListBox_SetParent(AObj: TColorListBox; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function ColorListBox_GetLeft(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ColorListBox_SetLeft(AObj: TColorListBox; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ColorListBox_GetTop(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ColorListBox_SetTop(AObj: TColorListBox; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ColorListBox_GetWidth(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ColorListBox_SetWidth(AObj: TColorListBox; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function ColorListBox_GetHeight(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ColorListBox_SetHeight(AObj: TColorListBox; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ColorListBox_GetCursor(AObj: TColorListBox): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure ColorListBox_SetCursor(AObj: TColorListBox; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function ColorListBox_GetHint(AObj: TColorListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ColorListBox_SetHint(AObj: TColorListBox; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function ColorListBox_GetComponentCount(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ColorListBox_GetComponentIndex(AObj: TColorListBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ColorListBox_SetComponentIndex(AObj: TColorListBox; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ColorListBox_GetOwner(AObj: TColorListBox): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ColorListBox_GetName(AObj: TColorListBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ColorListBox_SetName(AObj: TColorListBox; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ColorListBox_GetTag(AObj: TColorListBox): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ColorListBox_SetTag(AObj: TColorListBox; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ColorListBox_GetColors(AObj: TColorListBox; Index: Integer): TColor; extdecl;
begin
  Result :=  AObj.Colors[Index];
end;

function ColorListBox_GetColorNames(AObj: TColorListBox; Index: Integer): PChar; extdecl;
begin
  Result :=  PChar(AObj.ColorNames[Index]);
end;

function ColorListBox_GetControls(AObj: TColorListBox; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function ColorListBox_GetComponents(AObj: TColorListBox; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ColorListBox_Create,
  ColorListBox_Free,
  ColorListBox_Clear,
  ColorListBox_ClearSelection,
  ColorListBox_SelectAll,
  ColorListBox_CanFocus,
  ColorListBox_FlipChildren,
  ColorListBox_Focused,
  ColorListBox_Invalidate,
  ColorListBox_Realign,
  ColorListBox_Repaint,
  ColorListBox_ScaleBy,
  ColorListBox_SetBounds,
  ColorListBox_SetFocus,
  ColorListBox_Update,
  ColorListBox_BringToFront,
  ColorListBox_HasParent,
  ColorListBox_Hide,
  ColorListBox_Perform,
  ColorListBox_Refresh,
  ColorListBox_SendToBack,
  ColorListBox_Show,
  ColorListBox_GetTextBuf,
  ColorListBox_GetTextLen,
  ColorListBox_FindComponent,
  ColorListBox_GetNamePath,
  ColorListBox_Assign,
  ColorListBox_ClassName,
  ColorListBox_Equals,
  ColorListBox_GetHashCode,
  ColorListBox_ToString,
  ColorListBox_GetAlign,
  ColorListBox_SetAlign,
  ColorListBox_GetDefaultColorColor,
  ColorListBox_SetDefaultColorColor,
  ColorListBox_GetNoneColorColor,
  ColorListBox_SetNoneColorColor,
  ColorListBox_GetSelected,
  ColorListBox_SetSelected,
  ColorListBox_GetStyle,
  ColorListBox_SetStyle,
  ColorListBox_GetAnchors,
  ColorListBox_SetAnchors,
  ColorListBox_GetBiDiMode,
  ColorListBox_SetBiDiMode,
  ColorListBox_GetColor,
  ColorListBox_SetColor,
  ColorListBox_GetDoubleBuffered,
  ColorListBox_SetDoubleBuffered,
  ColorListBox_GetEnabled,
  ColorListBox_SetEnabled,
  ColorListBox_GetFont,
  ColorListBox_SetFont,
  ColorListBox_GetItemHeight,
  ColorListBox_SetItemHeight,
  ColorListBox_GetParentColor,
  ColorListBox_SetParentColor,
  ColorListBox_GetParentFont,
  ColorListBox_SetParentFont,
  ColorListBox_GetParentShowHint,
  ColorListBox_SetParentShowHint,
  ColorListBox_GetPopupMenu,
  ColorListBox_SetPopupMenu,
  ColorListBox_GetShowHint,
  ColorListBox_SetShowHint,
  ColorListBox_GetTabOrder,
  ColorListBox_SetTabOrder,
  ColorListBox_GetTabStop,
  ColorListBox_SetTabStop,
  ColorListBox_GetVisible,
  ColorListBox_SetVisible,
  ColorListBox_SetOnClick,
  ColorListBox_SetOnDblClick,
  ColorListBox_SetOnEnter,
  ColorListBox_SetOnExit,
  ColorListBox_SetOnKeyDown,
  ColorListBox_SetOnKeyPress,
  ColorListBox_SetOnKeyUp,
  ColorListBox_SetOnMouseDown,
  ColorListBox_SetOnMouseEnter,
  ColorListBox_SetOnMouseLeave,
  ColorListBox_SetOnMouseMove,
  ColorListBox_SetOnMouseUp,
  ColorListBox_GetCanvas,
  ColorListBox_GetCount,
  ColorListBox_GetItems,
  ColorListBox_SetItems,
  ColorListBox_GetMultiSelect,
  ColorListBox_SetMultiSelect,
  ColorListBox_GetSelCount,
  ColorListBox_GetItemIndex,
  ColorListBox_SetItemIndex,
  ColorListBox_GetBrush,
  ColorListBox_GetControlCount,
  ColorListBox_GetHandle,
  ColorListBox_GetParentWindow,
  ColorListBox_SetParentWindow,
  ColorListBox_GetAction,
  ColorListBox_SetAction,
  ColorListBox_GetBoundsRect,
  ColorListBox_SetBoundsRect,
  ColorListBox_GetClientHeight,
  ColorListBox_SetClientHeight,
  ColorListBox_GetClientRect,
  ColorListBox_GetClientWidth,
  ColorListBox_SetClientWidth,
  ColorListBox_GetParent,
  ColorListBox_SetParent,
  ColorListBox_GetLeft,
  ColorListBox_SetLeft,
  ColorListBox_GetTop,
  ColorListBox_SetTop,
  ColorListBox_GetWidth,
  ColorListBox_SetWidth,
  ColorListBox_GetHeight,
  ColorListBox_SetHeight,
  ColorListBox_GetCursor,
  ColorListBox_SetCursor,
  ColorListBox_GetHint,
  ColorListBox_SetHint,
  ColorListBox_GetComponentCount,
  ColorListBox_GetComponentIndex,
  ColorListBox_SetComponentIndex,
  ColorListBox_GetOwner,
  ColorListBox_GetName,
  ColorListBox_SetName,
  ColorListBox_GetTag,
  ColorListBox_SetTag,
  ColorListBox_GetColors,
  ColorListBox_GetColorNames,
  ColorListBox_GetControls,
  ColorListBox_GetComponents;

