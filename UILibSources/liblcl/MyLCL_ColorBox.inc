
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ColorBox_Create(AOwner: TComponent): TColorBox; extdecl;
begin
  Result :=  TColorBox.Create(AOwner);
end;

procedure ColorBox_Free(AObj: TColorBox); extdecl;
begin
  AObj.Free;
end;

procedure ColorBox_Clear(AObj: TColorBox); extdecl;
begin
  AObj.Clear;
end;

procedure ColorBox_ClearSelection(AObj: TColorBox); extdecl;
begin
  AObj.ClearSelection;
end;

function ColorBox_Focused(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure ColorBox_SelectAll(AObj: TColorBox); extdecl;
begin
  AObj.SelectAll;
end;

function ColorBox_CanFocus(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure ColorBox_FlipChildren(AObj: TColorBox; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

procedure ColorBox_Invalidate(AObj: TColorBox); extdecl;
begin
  AObj.Invalidate;
end;

procedure ColorBox_Realign(AObj: TColorBox); extdecl;
begin
  AObj.Realign;
end;

procedure ColorBox_Repaint(AObj: TColorBox); extdecl;
begin
  AObj.Repaint;
end;

procedure ColorBox_ScaleBy(AObj: TColorBox; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure ColorBox_SetBounds(AObj: TColorBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ColorBox_SetFocus(AObj: TColorBox); extdecl;
begin
  AObj.SetFocus;
end;

procedure ColorBox_Update(AObj: TColorBox); extdecl;
begin
  AObj.Update;
end;

procedure ColorBox_BringToFront(AObj: TColorBox); extdecl;
begin
  AObj.BringToFront;
end;

function ColorBox_HasParent(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ColorBox_Hide(AObj: TColorBox); extdecl;
begin
  AObj.Hide;
end;

function ColorBox_Perform(AObj: TColorBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ColorBox_Refresh(AObj: TColorBox); extdecl;
begin
  AObj.Refresh;
end;

procedure ColorBox_SendToBack(AObj: TColorBox); extdecl;
begin
  AObj.SendToBack;
end;

procedure ColorBox_Show(AObj: TColorBox); extdecl;
begin
  AObj.Show;
end;

function ColorBox_GetTextBuf(AObj: TColorBox; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ColorBox_GetTextLen(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function ColorBox_FindComponent(AObj: TColorBox; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ColorBox_GetNamePath(AObj: TColorBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ColorBox_Assign(AObj: TColorBox; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ColorBox_ClassName(AObj: TColorBox): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ColorBox_Equals(AObj: TColorBox; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ColorBox_GetHashCode(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ColorBox_ToString(AObj: TColorBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ColorBox_GetAlign(AObj: TColorBox): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure ColorBox_SetAlign(AObj: TColorBox; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function ColorBox_GetAutoComplete(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.AutoComplete;
end;

procedure ColorBox_SetAutoComplete(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.AutoComplete := AValue;
end;

function ColorBox_GetAutoDropDown(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.AutoDropDown;
end;

procedure ColorBox_SetAutoDropDown(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.AutoDropDown := AValue;
end;

function ColorBox_GetDefaultColorColor(AObj: TColorBox): TColor; extdecl;
begin
  Result :=  AObj.DefaultColorColor;
end;

procedure ColorBox_SetDefaultColorColor(AObj: TColorBox; AValue: TColor); extdecl;
begin
  AObj.DefaultColorColor := AValue;
end;

function ColorBox_GetNoneColorColor(AObj: TColorBox): TColor; extdecl;
begin
  Result :=  AObj.NoneColorColor;
end;

procedure ColorBox_SetNoneColorColor(AObj: TColorBox; AValue: TColor); extdecl;
begin
  AObj.NoneColorColor := AValue;
end;

function ColorBox_GetSelected(AObj: TColorBox): TColor; extdecl;
begin
  Result :=  AObj.Selected;
end;

procedure ColorBox_SetSelected(AObj: TColorBox; AValue: TColor); extdecl;
begin
  AObj.Selected := AValue;
end;

function ColorBox_GetStyle(AObj: TColorBox): TColorBoxStyle; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure ColorBox_SetStyle(AObj: TColorBox; AValue: TColorBoxStyle); extdecl;
begin
  AObj.Style := AValue;
end;

function ColorBox_GetAnchors(AObj: TColorBox): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure ColorBox_SetAnchors(AObj: TColorBox; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function ColorBox_GetBiDiMode(AObj: TColorBox): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ColorBox_SetBiDiMode(AObj: TColorBox; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function ColorBox_GetColor(AObj: TColorBox): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure ColorBox_SetColor(AObj: TColorBox; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function ColorBox_GetDoubleBuffered(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ColorBox_SetDoubleBuffered(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function ColorBox_GetDropDownCount(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.DropDownCount;
end;

procedure ColorBox_SetDropDownCount(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.DropDownCount := AValue;
end;

function ColorBox_GetEnabled(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure ColorBox_SetEnabled(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function ColorBox_GetFont(AObj: TColorBox): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure ColorBox_SetFont(AObj: TColorBox; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function ColorBox_GetItemHeight(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.ItemHeight;
end;

procedure ColorBox_SetItemHeight(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.ItemHeight := AValue;
end;

function ColorBox_GetParentColor(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure ColorBox_SetParentColor(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function ColorBox_GetParentFont(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure ColorBox_SetParentFont(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function ColorBox_GetParentShowHint(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ColorBox_SetParentShowHint(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function ColorBox_GetPopupMenu(AObj: TColorBox): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ColorBox_SetPopupMenu(AObj: TColorBox; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function ColorBox_GetShowHint(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure ColorBox_SetShowHint(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function ColorBox_GetTabOrder(AObj: TColorBox): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure ColorBox_SetTabOrder(AObj: TColorBox; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function ColorBox_GetTabStop(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure ColorBox_SetTabStop(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function ColorBox_GetVisible(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure ColorBox_SetVisible(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure ColorBox_SetOnChange(AObj: TColorBox; AEventId: NativeUInt); extdecl;
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

procedure ColorBox_SetOnClick(AObj: TColorBox; AEventId: NativeUInt); extdecl;
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

procedure ColorBox_SetOnEnter(AObj: TColorBox; AEventId: NativeUInt); extdecl;
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

procedure ColorBox_SetOnExit(AObj: TColorBox; AEventId: NativeUInt); extdecl;
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

procedure ColorBox_SetOnKeyDown(AObj: TColorBox; AEventId: NativeUInt); extdecl;
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

procedure ColorBox_SetOnKeyPress(AObj: TColorBox; AEventId: NativeUInt); extdecl;
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

procedure ColorBox_SetOnKeyUp(AObj: TColorBox; AEventId: NativeUInt); extdecl;
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

procedure ColorBox_SetOnMouseEnter(AObj: TColorBox; AEventId: NativeUInt); extdecl;
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

procedure ColorBox_SetOnMouseLeave(AObj: TColorBox; AEventId: NativeUInt); extdecl;
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

function ColorBox_GetSelText(AObj: TColorBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure ColorBox_SetSelText(AObj: TColorBox; AValue: PChar); extdecl;
begin
  AObj.SelText := AValue;
end;

function ColorBox_GetCanvas(AObj: TColorBox): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function ColorBox_GetDroppedDown(AObj: TColorBox): LongBool; extdecl;
begin
  Result :=  AObj.DroppedDown;
end;

procedure ColorBox_SetDroppedDown(AObj: TColorBox; AValue: LongBool); extdecl;
begin
  AObj.DroppedDown := AValue;
end;

function ColorBox_GetItems(AObj: TColorBox): TStrings; extdecl;
begin
  Result :=  AObj.Items;
end;

procedure ColorBox_SetItems(AObj: TColorBox; AValue: TStrings); extdecl;
begin
  AObj.Items := AValue;
end;

function ColorBox_GetSelLength(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.SelLength;
end;

procedure ColorBox_SetSelLength(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.SelLength := AValue;
end;

function ColorBox_GetSelStart(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.SelStart;
end;

procedure ColorBox_SetSelStart(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.SelStart := AValue;
end;

function ColorBox_GetItemIndex(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.ItemIndex;
end;

procedure ColorBox_SetItemIndex(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.ItemIndex := AValue;
end;

function ColorBox_GetBrush(AObj: TColorBox): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function ColorBox_GetControlCount(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function ColorBox_GetHandle(AObj: TColorBox): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function ColorBox_GetParentWindow(AObj: TColorBox): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ColorBox_SetParentWindow(AObj: TColorBox; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function ColorBox_GetAction(AObj: TColorBox): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure ColorBox_SetAction(AObj: TColorBox; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure ColorBox_GetBoundsRect(AObj: TColorBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ColorBox_SetBoundsRect(AObj: TColorBox; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function ColorBox_GetClientHeight(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ColorBox_SetClientHeight(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure ColorBox_GetClientRect(AObj: TColorBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function ColorBox_GetClientWidth(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ColorBox_SetClientWidth(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function ColorBox_GetParent(AObj: TColorBox): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure ColorBox_SetParent(AObj: TColorBox; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function ColorBox_GetLeft(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ColorBox_SetLeft(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ColorBox_GetTop(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ColorBox_SetTop(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ColorBox_GetWidth(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ColorBox_SetWidth(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function ColorBox_GetHeight(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ColorBox_SetHeight(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ColorBox_GetCursor(AObj: TColorBox): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure ColorBox_SetCursor(AObj: TColorBox; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function ColorBox_GetHint(AObj: TColorBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ColorBox_SetHint(AObj: TColorBox; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function ColorBox_GetComponentCount(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ColorBox_GetComponentIndex(AObj: TColorBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ColorBox_SetComponentIndex(AObj: TColorBox; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ColorBox_GetOwner(AObj: TColorBox): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ColorBox_GetName(AObj: TColorBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ColorBox_SetName(AObj: TColorBox; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ColorBox_GetTag(AObj: TColorBox): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ColorBox_SetTag(AObj: TColorBox; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ColorBox_GetColors(AObj: TColorBox; Index: Integer): TColor; extdecl;
begin
  Result :=  AObj.Colors[Index];
end;

function ColorBox_GetColorNames(AObj: TColorBox; Index: Integer): PChar; extdecl;
begin
  Result :=  PChar(AObj.ColorNames[Index]);
end;

function ColorBox_GetControls(AObj: TColorBox; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function ColorBox_GetComponents(AObj: TColorBox; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ColorBox_Create,
  ColorBox_Free,
  ColorBox_Clear,
  ColorBox_ClearSelection,
  ColorBox_Focused,
  ColorBox_SelectAll,
  ColorBox_CanFocus,
  ColorBox_FlipChildren,
  ColorBox_Invalidate,
  ColorBox_Realign,
  ColorBox_Repaint,
  ColorBox_ScaleBy,
  ColorBox_SetBounds,
  ColorBox_SetFocus,
  ColorBox_Update,
  ColorBox_BringToFront,
  ColorBox_HasParent,
  ColorBox_Hide,
  ColorBox_Perform,
  ColorBox_Refresh,
  ColorBox_SendToBack,
  ColorBox_Show,
  ColorBox_GetTextBuf,
  ColorBox_GetTextLen,
  ColorBox_FindComponent,
  ColorBox_GetNamePath,
  ColorBox_Assign,
  ColorBox_ClassName,
  ColorBox_Equals,
  ColorBox_GetHashCode,
  ColorBox_ToString,
  ColorBox_GetAlign,
  ColorBox_SetAlign,
  ColorBox_GetAutoComplete,
  ColorBox_SetAutoComplete,
  ColorBox_GetAutoDropDown,
  ColorBox_SetAutoDropDown,
  ColorBox_GetDefaultColorColor,
  ColorBox_SetDefaultColorColor,
  ColorBox_GetNoneColorColor,
  ColorBox_SetNoneColorColor,
  ColorBox_GetSelected,
  ColorBox_SetSelected,
  ColorBox_GetStyle,
  ColorBox_SetStyle,
  ColorBox_GetAnchors,
  ColorBox_SetAnchors,
  ColorBox_GetBiDiMode,
  ColorBox_SetBiDiMode,
  ColorBox_GetColor,
  ColorBox_SetColor,
  ColorBox_GetDoubleBuffered,
  ColorBox_SetDoubleBuffered,
  ColorBox_GetDropDownCount,
  ColorBox_SetDropDownCount,
  ColorBox_GetEnabled,
  ColorBox_SetEnabled,
  ColorBox_GetFont,
  ColorBox_SetFont,
  ColorBox_GetItemHeight,
  ColorBox_SetItemHeight,
  ColorBox_GetParentColor,
  ColorBox_SetParentColor,
  ColorBox_GetParentFont,
  ColorBox_SetParentFont,
  ColorBox_GetParentShowHint,
  ColorBox_SetParentShowHint,
  ColorBox_GetPopupMenu,
  ColorBox_SetPopupMenu,
  ColorBox_GetShowHint,
  ColorBox_SetShowHint,
  ColorBox_GetTabOrder,
  ColorBox_SetTabOrder,
  ColorBox_GetTabStop,
  ColorBox_SetTabStop,
  ColorBox_GetVisible,
  ColorBox_SetVisible,
  ColorBox_SetOnChange,
  ColorBox_SetOnClick,
  ColorBox_SetOnEnter,
  ColorBox_SetOnExit,
  ColorBox_SetOnKeyDown,
  ColorBox_SetOnKeyPress,
  ColorBox_SetOnKeyUp,
  ColorBox_SetOnMouseEnter,
  ColorBox_SetOnMouseLeave,
  ColorBox_GetSelText,
  ColorBox_SetSelText,
  ColorBox_GetCanvas,
  ColorBox_GetDroppedDown,
  ColorBox_SetDroppedDown,
  ColorBox_GetItems,
  ColorBox_SetItems,
  ColorBox_GetSelLength,
  ColorBox_SetSelLength,
  ColorBox_GetSelStart,
  ColorBox_SetSelStart,
  ColorBox_GetItemIndex,
  ColorBox_SetItemIndex,
  ColorBox_GetBrush,
  ColorBox_GetControlCount,
  ColorBox_GetHandle,
  ColorBox_GetParentWindow,
  ColorBox_SetParentWindow,
  ColorBox_GetAction,
  ColorBox_SetAction,
  ColorBox_GetBoundsRect,
  ColorBox_SetBoundsRect,
  ColorBox_GetClientHeight,
  ColorBox_SetClientHeight,
  ColorBox_GetClientRect,
  ColorBox_GetClientWidth,
  ColorBox_SetClientWidth,
  ColorBox_GetParent,
  ColorBox_SetParent,
  ColorBox_GetLeft,
  ColorBox_SetLeft,
  ColorBox_GetTop,
  ColorBox_SetTop,
  ColorBox_GetWidth,
  ColorBox_SetWidth,
  ColorBox_GetHeight,
  ColorBox_SetHeight,
  ColorBox_GetCursor,
  ColorBox_SetCursor,
  ColorBox_GetHint,
  ColorBox_SetHint,
  ColorBox_GetComponentCount,
  ColorBox_GetComponentIndex,
  ColorBox_SetComponentIndex,
  ColorBox_GetOwner,
  ColorBox_GetName,
  ColorBox_SetName,
  ColorBox_GetTag,
  ColorBox_SetTag,
  ColorBox_GetColors,
  ColorBox_GetColorNames,
  ColorBox_GetControls,
  ColorBox_GetComponents;

