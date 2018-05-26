
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ColorBox_Create(AOwner: TComponent): TColorBox; stdcall;
begin
  Result :=  TColorBox.Create(AOwner);
end;

procedure ColorBox_Free(AObj: TColorBox); stdcall;
begin
  AObj.Free;
end;

procedure ColorBox_AddItem(AObj: TColorBox; Item: PChar; AObject: TObject); stdcall;
begin
  AObj.AddItem(Item, AObject);
end;

procedure ColorBox_Clear(AObj: TColorBox); stdcall;
begin
  AObj.Clear;
end;

procedure ColorBox_ClearSelection(AObj: TColorBox); stdcall;
begin
  AObj.ClearSelection;
end;

procedure ColorBox_DeleteSelected(AObj: TColorBox); stdcall;
begin
  AObj.DeleteSelected;
end;

function ColorBox_Focused(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

procedure ColorBox_SelectAll(AObj: TColorBox); stdcall;
begin
  AObj.SelectAll;
end;

function ColorBox_CanFocus(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure ColorBox_FlipChildren(AObj: TColorBox; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function ColorBox_HandleAllocated(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ColorBox_Invalidate(AObj: TColorBox); stdcall;
begin
  AObj.Invalidate;
end;

procedure ColorBox_Realign(AObj: TColorBox); stdcall;
begin
  AObj.Realign;
end;

procedure ColorBox_Repaint(AObj: TColorBox); stdcall;
begin
  AObj.Repaint;
end;

procedure ColorBox_ScaleBy(AObj: TColorBox; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure ColorBox_SetBounds(AObj: TColorBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ColorBox_SetFocus(AObj: TColorBox); stdcall;
begin
  AObj.SetFocus;
end;

procedure ColorBox_Update(AObj: TColorBox); stdcall;
begin
  AObj.Update;
end;

procedure ColorBox_BringToFront(AObj: TColorBox); stdcall;
begin
  AObj.BringToFront;
end;

function ColorBox_HasParent(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ColorBox_Hide(AObj: TColorBox); stdcall;
begin
  AObj.Hide;
end;

function ColorBox_Perform(AObj: TColorBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ColorBox_Refresh(AObj: TColorBox); stdcall;
begin
  AObj.Refresh;
end;

procedure ColorBox_SendToBack(AObj: TColorBox); stdcall;
begin
  AObj.SendToBack;
end;

procedure ColorBox_Show(AObj: TColorBox); stdcall;
begin
  AObj.Show;
end;

function ColorBox_GetTextBuf(AObj: TColorBox; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ColorBox_GetTextLen(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function ColorBox_FindComponent(AObj: TColorBox; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ColorBox_GetNamePath(AObj: TColorBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ColorBox_Assign(AObj: TColorBox; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ColorBox_ClassName(AObj: TColorBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ColorBox_Equals(AObj: TColorBox; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ColorBox_GetHashCode(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ColorBox_ToString(AObj: TColorBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ColorBox_GetAlign(AObj: TColorBox): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ColorBox_SetAlign(AObj: TColorBox; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ColorBox_GetAutoComplete(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.AutoComplete;
end;

procedure ColorBox_SetAutoComplete(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.AutoComplete := AValue;
end;

function ColorBox_GetAutoDropDown(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.AutoDropDown;
end;

procedure ColorBox_SetAutoDropDown(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.AutoDropDown := AValue;
end;

function ColorBox_GetDefaultColorColor(AObj: TColorBox): TColor; stdcall;
begin
  Result :=  AObj.DefaultColorColor;
end;

procedure ColorBox_SetDefaultColorColor(AObj: TColorBox; AValue: TColor); stdcall;
begin
  AObj.DefaultColorColor := AValue;
end;

function ColorBox_GetNoneColorColor(AObj: TColorBox): TColor; stdcall;
begin
  Result :=  AObj.NoneColorColor;
end;

procedure ColorBox_SetNoneColorColor(AObj: TColorBox; AValue: TColor); stdcall;
begin
  AObj.NoneColorColor := AValue;
end;

function ColorBox_GetSelected(AObj: TColorBox): TColor; stdcall;
begin
  Result :=  AObj.Selected;
end;

procedure ColorBox_SetSelected(AObj: TColorBox; AValue: TColor); stdcall;
begin
  AObj.Selected := AValue;
end;

function ColorBox_GetStyle(AObj: TColorBox): TColorBoxStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure ColorBox_SetStyle(AObj: TColorBox; AValue: TColorBoxStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function ColorBox_GetAnchors(AObj: TColorBox): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ColorBox_SetAnchors(AObj: TColorBox; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ColorBox_GetBevelEdges(AObj: TColorBox): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure ColorBox_SetBevelEdges(AObj: TColorBox; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function ColorBox_GetBevelInner(AObj: TColorBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure ColorBox_SetBevelInner(AObj: TColorBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function ColorBox_GetBevelKind(AObj: TColorBox): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure ColorBox_SetBevelKind(AObj: TColorBox; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function ColorBox_GetBevelOuter(AObj: TColorBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure ColorBox_SetBevelOuter(AObj: TColorBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function ColorBox_GetBiDiMode(AObj: TColorBox): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ColorBox_SetBiDiMode(AObj: TColorBox; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function ColorBox_GetColor(AObj: TColorBox): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure ColorBox_SetColor(AObj: TColorBox; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function ColorBox_GetDoubleBuffered(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ColorBox_SetDoubleBuffered(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function ColorBox_GetDropDownCount(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.DropDownCount;
end;

procedure ColorBox_SetDropDownCount(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.DropDownCount := AValue;
end;

function ColorBox_GetEnabled(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ColorBox_SetEnabled(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ColorBox_GetFont(AObj: TColorBox): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure ColorBox_SetFont(AObj: TColorBox; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function ColorBox_GetItemHeight(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ItemHeight;
end;

procedure ColorBox_SetItemHeight(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.ItemHeight := AValue;
end;

function ColorBox_GetParentColor(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure ColorBox_SetParentColor(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function ColorBox_GetParentCtl3D(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure ColorBox_SetParentCtl3D(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function ColorBox_GetParentDoubleBuffered(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure ColorBox_SetParentDoubleBuffered(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function ColorBox_GetParentFont(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure ColorBox_SetParentFont(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function ColorBox_GetParentShowHint(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ColorBox_SetParentShowHint(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ColorBox_GetPopupMenu(AObj: TColorBox): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ColorBox_SetPopupMenu(AObj: TColorBox; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ColorBox_GetShowHint(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ColorBox_SetShowHint(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ColorBox_GetTabOrder(AObj: TColorBox): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure ColorBox_SetTabOrder(AObj: TColorBox; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function ColorBox_GetTabStop(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure ColorBox_SetTabStop(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function ColorBox_GetVisible(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ColorBox_SetVisible(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function ColorBox_GetStyleElements(AObj: TColorBox): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ColorBox_SetStyleElements(AObj: TColorBox; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure ColorBox_SetOnChange(AObj: TColorBox; AEventId: NativeUInt); stdcall;
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

procedure ColorBox_SetOnClick(AObj: TColorBox; AEventId: NativeUInt); stdcall;
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

procedure ColorBox_SetOnEnter(AObj: TColorBox; AEventId: NativeUInt); stdcall;
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

procedure ColorBox_SetOnExit(AObj: TColorBox; AEventId: NativeUInt); stdcall;
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

procedure ColorBox_SetOnKeyDown(AObj: TColorBox; AEventId: NativeUInt); stdcall;
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

procedure ColorBox_SetOnKeyPress(AObj: TColorBox; AEventId: NativeUInt); stdcall;
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

procedure ColorBox_SetOnKeyUp(AObj: TColorBox; AEventId: NativeUInt); stdcall;
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

procedure ColorBox_SetOnMouseEnter(AObj: TColorBox; AEventId: NativeUInt); stdcall;
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

procedure ColorBox_SetOnMouseLeave(AObj: TColorBox; AEventId: NativeUInt); stdcall;
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

function ColorBox_GetAutoCompleteDelay(AObj: TColorBox): Cardinal; stdcall;
begin
  Result :=  AObj.AutoCompleteDelay;
end;

procedure ColorBox_SetAutoCompleteDelay(AObj: TColorBox; AValue: Cardinal); stdcall;
begin
  AObj.AutoCompleteDelay := AValue;
end;

function ColorBox_GetAutoCloseUp(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.AutoCloseUp;
end;

procedure ColorBox_SetAutoCloseUp(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.AutoCloseUp := AValue;
end;

function ColorBox_GetSelText(AObj: TColorBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure ColorBox_SetSelText(AObj: TColorBox; AValue: PChar); stdcall;
begin
  AObj.SelText := AValue;
end;

function ColorBox_GetTextHint(AObj: TColorBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.TextHint);
end;

procedure ColorBox_SetTextHint(AObj: TColorBox; AValue: PChar); stdcall;
begin
  AObj.TextHint := AValue;
end;

function ColorBox_GetCanvas(AObj: TColorBox): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function ColorBox_GetDroppedDown(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.DroppedDown;
end;

procedure ColorBox_SetDroppedDown(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.DroppedDown := AValue;
end;

function ColorBox_GetItems(AObj: TColorBox): TStrings; stdcall;
begin
  Result :=  AObj.Items;
end;

procedure ColorBox_SetItems(AObj: TColorBox; AValue: TStrings); stdcall;
begin
  AObj.Items := AValue;
end;

function ColorBox_GetSelLength(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.SelLength;
end;

procedure ColorBox_SetSelLength(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.SelLength := AValue;
end;

function ColorBox_GetSelStart(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.SelStart;
end;

procedure ColorBox_SetSelStart(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.SelStart := AValue;
end;

function ColorBox_GetItemIndex(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ItemIndex;
end;

procedure ColorBox_SetItemIndex(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.ItemIndex := AValue;
end;

function ColorBox_GetBrush(AObj: TColorBox): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function ColorBox_GetControlCount(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function ColorBox_GetHandle(AObj: TColorBox): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ColorBox_GetParentWindow(AObj: TColorBox): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ColorBox_SetParentWindow(AObj: TColorBox; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function ColorBox_GetAction(AObj: TColorBox): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ColorBox_SetAction(AObj: TColorBox; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure ColorBox_GetBoundsRect(AObj: TColorBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ColorBox_SetBoundsRect(AObj: TColorBox; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ColorBox_GetClientHeight(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ColorBox_SetClientHeight(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ColorBox_GetClientRect(AObj: TColorBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ColorBox_GetClientWidth(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ColorBox_SetClientWidth(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ColorBox_GetExplicitLeft(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ColorBox_GetExplicitTop(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ColorBox_GetExplicitWidth(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ColorBox_GetExplicitHeight(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ColorBox_GetParent(AObj: TColorBox): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ColorBox_SetParent(AObj: TColorBox; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ColorBox_GetAlignWithMargins(AObj: TColorBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ColorBox_SetAlignWithMargins(AObj: TColorBox; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ColorBox_GetLeft(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ColorBox_SetLeft(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ColorBox_GetTop(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ColorBox_SetTop(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ColorBox_GetWidth(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ColorBox_SetWidth(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ColorBox_GetHeight(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ColorBox_SetHeight(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ColorBox_GetCursor(AObj: TColorBox): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ColorBox_SetCursor(AObj: TColorBox; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ColorBox_GetHint(AObj: TColorBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ColorBox_SetHint(AObj: TColorBox; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ColorBox_GetMargins(AObj: TColorBox): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ColorBox_SetMargins(AObj: TColorBox; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ColorBox_GetCustomHint(AObj: TColorBox): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ColorBox_SetCustomHint(AObj: TColorBox; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ColorBox_GetComponentCount(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ColorBox_GetComponentIndex(AObj: TColorBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ColorBox_SetComponentIndex(AObj: TColorBox; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ColorBox_GetOwner(AObj: TColorBox): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ColorBox_GetName(AObj: TColorBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ColorBox_SetName(AObj: TColorBox; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ColorBox_GetTag(AObj: TColorBox): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ColorBox_SetTag(AObj: TColorBox; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ColorBox_GetColors(AObj: TColorBox; Index: Integer): TColor; stdcall;
begin
  Result :=  AObj.Colors[Index];
end;

function ColorBox_GetColorNames(AObj: TColorBox; Index: Integer): PChar; stdcall;
begin
  Result :=  PChar(AObj.ColorNames[Index]);
end;

function ColorBox_GetControls(AObj: TColorBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ColorBox_GetComponents(AObj: TColorBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ColorBox_Create,
  ColorBox_Free,
  ColorBox_AddItem,
  ColorBox_Clear,
  ColorBox_ClearSelection,
  ColorBox_DeleteSelected,
  ColorBox_Focused,
  ColorBox_SelectAll,
  ColorBox_CanFocus,
  ColorBox_FlipChildren,
  ColorBox_HandleAllocated,
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
  ColorBox_GetBevelEdges,
  ColorBox_SetBevelEdges,
  ColorBox_GetBevelInner,
  ColorBox_SetBevelInner,
  ColorBox_GetBevelKind,
  ColorBox_SetBevelKind,
  ColorBox_GetBevelOuter,
  ColorBox_SetBevelOuter,
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
  ColorBox_GetParentCtl3D,
  ColorBox_SetParentCtl3D,
  ColorBox_GetParentDoubleBuffered,
  ColorBox_SetParentDoubleBuffered,
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
  ColorBox_GetStyleElements,
  ColorBox_SetStyleElements,
  ColorBox_SetOnChange,
  ColorBox_SetOnClick,
  ColorBox_SetOnEnter,
  ColorBox_SetOnExit,
  ColorBox_SetOnKeyDown,
  ColorBox_SetOnKeyPress,
  ColorBox_SetOnKeyUp,
  ColorBox_SetOnMouseEnter,
  ColorBox_SetOnMouseLeave,
  ColorBox_GetAutoCompleteDelay,
  ColorBox_SetAutoCompleteDelay,
  ColorBox_GetAutoCloseUp,
  ColorBox_SetAutoCloseUp,
  ColorBox_GetSelText,
  ColorBox_SetSelText,
  ColorBox_GetTextHint,
  ColorBox_SetTextHint,
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
  ColorBox_GetExplicitLeft,
  ColorBox_GetExplicitTop,
  ColorBox_GetExplicitWidth,
  ColorBox_GetExplicitHeight,
  ColorBox_GetParent,
  ColorBox_SetParent,
  ColorBox_GetAlignWithMargins,
  ColorBox_SetAlignWithMargins,
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
  ColorBox_GetMargins,
  ColorBox_SetMargins,
  ColorBox_GetCustomHint,
  ColorBox_SetCustomHint,
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

