
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ColorListBox_Create(AOwner: TComponent): TColorListBox; stdcall;
begin
  Result :=  TColorListBox.Create(AOwner);
end;

procedure ColorListBox_Free(AObj: TColorListBox); stdcall;
begin
  AObj.Free;
end;

procedure ColorListBox_AddItem(AObj: TColorListBox; Item: PChar; AObject: TObject); stdcall;
begin
  AObj.AddItem(Item, AObject);
end;

procedure ColorListBox_Clear(AObj: TColorListBox); stdcall;
begin
  AObj.Clear;
end;

procedure ColorListBox_ClearSelection(AObj: TColorListBox); stdcall;
begin
  AObj.ClearSelection;
end;

procedure ColorListBox_DeleteSelected(AObj: TColorListBox); stdcall;
begin
  AObj.DeleteSelected;
end;

procedure ColorListBox_SelectAll(AObj: TColorListBox); stdcall;
begin
  AObj.SelectAll;
end;

function ColorListBox_CanFocus(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure ColorListBox_FlipChildren(AObj: TColorListBox; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function ColorListBox_Focused(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function ColorListBox_HandleAllocated(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ColorListBox_Invalidate(AObj: TColorListBox); stdcall;
begin
  AObj.Invalidate;
end;

procedure ColorListBox_Realign(AObj: TColorListBox); stdcall;
begin
  AObj.Realign;
end;

procedure ColorListBox_Repaint(AObj: TColorListBox); stdcall;
begin
  AObj.Repaint;
end;

procedure ColorListBox_ScaleBy(AObj: TColorListBox; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure ColorListBox_SetBounds(AObj: TColorListBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ColorListBox_SetFocus(AObj: TColorListBox); stdcall;
begin
  AObj.SetFocus;
end;

procedure ColorListBox_Update(AObj: TColorListBox); stdcall;
begin
  AObj.Update;
end;

procedure ColorListBox_BringToFront(AObj: TColorListBox); stdcall;
begin
  AObj.BringToFront;
end;

function ColorListBox_HasParent(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ColorListBox_Hide(AObj: TColorListBox); stdcall;
begin
  AObj.Hide;
end;

function ColorListBox_Perform(AObj: TColorListBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ColorListBox_Refresh(AObj: TColorListBox); stdcall;
begin
  AObj.Refresh;
end;

procedure ColorListBox_SendToBack(AObj: TColorListBox); stdcall;
begin
  AObj.SendToBack;
end;

procedure ColorListBox_Show(AObj: TColorListBox); stdcall;
begin
  AObj.Show;
end;

function ColorListBox_GetTextBuf(AObj: TColorListBox; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ColorListBox_GetTextLen(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function ColorListBox_FindComponent(AObj: TColorListBox; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ColorListBox_GetNamePath(AObj: TColorListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ColorListBox_Assign(AObj: TColorListBox; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ColorListBox_ClassName(AObj: TColorListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ColorListBox_Equals(AObj: TColorListBox; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ColorListBox_GetHashCode(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ColorListBox_ToString(AObj: TColorListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ColorListBox_GetAlign(AObj: TColorListBox): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ColorListBox_SetAlign(AObj: TColorListBox; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ColorListBox_GetAutoComplete(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.AutoComplete;
end;

procedure ColorListBox_SetAutoComplete(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.AutoComplete := AValue;
end;

function ColorListBox_GetDefaultColorColor(AObj: TColorListBox): TColor; stdcall;
begin
  Result :=  AObj.DefaultColorColor;
end;

procedure ColorListBox_SetDefaultColorColor(AObj: TColorListBox; AValue: TColor); stdcall;
begin
  AObj.DefaultColorColor := AValue;
end;

function ColorListBox_GetNoneColorColor(AObj: TColorListBox): TColor; stdcall;
begin
  Result :=  AObj.NoneColorColor;
end;

procedure ColorListBox_SetNoneColorColor(AObj: TColorListBox; AValue: TColor); stdcall;
begin
  AObj.NoneColorColor := AValue;
end;

function ColorListBox_GetSelected(AObj: TColorListBox): TColor; stdcall;
begin
  Result :=  AObj.Selected;
end;

procedure ColorListBox_SetSelected(AObj: TColorListBox; AValue: TColor); stdcall;
begin
  AObj.Selected := AValue;
end;

function ColorListBox_GetStyle(AObj: TColorListBox): TColorBoxStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure ColorListBox_SetStyle(AObj: TColorListBox; AValue: TColorBoxStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function ColorListBox_GetAnchors(AObj: TColorListBox): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ColorListBox_SetAnchors(AObj: TColorListBox; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ColorListBox_GetBevelEdges(AObj: TColorListBox): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure ColorListBox_SetBevelEdges(AObj: TColorListBox; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function ColorListBox_GetBevelInner(AObj: TColorListBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure ColorListBox_SetBevelInner(AObj: TColorListBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function ColorListBox_GetBevelKind(AObj: TColorListBox): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure ColorListBox_SetBevelKind(AObj: TColorListBox; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function ColorListBox_GetBevelOuter(AObj: TColorListBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure ColorListBox_SetBevelOuter(AObj: TColorListBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function ColorListBox_GetBiDiMode(AObj: TColorListBox): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ColorListBox_SetBiDiMode(AObj: TColorListBox; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function ColorListBox_GetColor(AObj: TColorListBox): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure ColorListBox_SetColor(AObj: TColorListBox; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function ColorListBox_GetDoubleBuffered(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ColorListBox_SetDoubleBuffered(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function ColorListBox_GetEnabled(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ColorListBox_SetEnabled(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ColorListBox_GetFont(AObj: TColorListBox): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure ColorListBox_SetFont(AObj: TColorListBox; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function ColorListBox_GetItemHeight(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ItemHeight;
end;

procedure ColorListBox_SetItemHeight(AObj: TColorListBox; AValue: Integer); stdcall;
begin
  AObj.ItemHeight := AValue;
end;

function ColorListBox_GetParentColor(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure ColorListBox_SetParentColor(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function ColorListBox_GetParentCtl3D(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure ColorListBox_SetParentCtl3D(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function ColorListBox_GetParentDoubleBuffered(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure ColorListBox_SetParentDoubleBuffered(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function ColorListBox_GetParentFont(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure ColorListBox_SetParentFont(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function ColorListBox_GetParentShowHint(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ColorListBox_SetParentShowHint(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ColorListBox_GetPopupMenu(AObj: TColorListBox): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ColorListBox_SetPopupMenu(AObj: TColorListBox; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ColorListBox_GetShowHint(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ColorListBox_SetShowHint(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ColorListBox_GetTabOrder(AObj: TColorListBox): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure ColorListBox_SetTabOrder(AObj: TColorListBox; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function ColorListBox_GetTabStop(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure ColorListBox_SetTabStop(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function ColorListBox_GetVisible(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ColorListBox_SetVisible(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function ColorListBox_GetStyleElements(AObj: TColorListBox): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ColorListBox_SetStyleElements(AObj: TColorListBox; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure ColorListBox_SetOnClick(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnDblClick(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnEnter(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnExit(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnKeyDown(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnKeyPress(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnKeyUp(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnMouseDown(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnMouseEnter(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnMouseLeave(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnMouseMove(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

procedure ColorListBox_SetOnMouseUp(AObj: TColorListBox; AEventId: NativeUInt); stdcall;
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

function ColorListBox_GetAutoCompleteDelay(AObj: TColorListBox): Cardinal; stdcall;
begin
  Result :=  AObj.AutoCompleteDelay;
end;

procedure ColorListBox_SetAutoCompleteDelay(AObj: TColorListBox; AValue: Cardinal); stdcall;
begin
  AObj.AutoCompleteDelay := AValue;
end;

function ColorListBox_GetCanvas(AObj: TColorListBox): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function ColorListBox_GetCount(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.Count;
end;

procedure ColorListBox_SetCount(AObj: TColorListBox; AValue: Integer); stdcall;
begin
  AObj.Count := AValue;
end;

function ColorListBox_GetItems(AObj: TColorListBox): TStrings; stdcall;
begin
  Result :=  AObj.Items;
end;

procedure ColorListBox_SetItems(AObj: TColorListBox; AValue: TStrings); stdcall;
begin
  AObj.Items := AValue;
end;

function ColorListBox_GetMultiSelect(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.MultiSelect;
end;

procedure ColorListBox_SetMultiSelect(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.MultiSelect := AValue;
end;

function ColorListBox_GetSelCount(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.SelCount;
end;

function ColorListBox_GetItemIndex(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ItemIndex;
end;

procedure ColorListBox_SetItemIndex(AObj: TColorListBox; AValue: Integer); stdcall;
begin
  AObj.ItemIndex := AValue;
end;

function ColorListBox_GetBrush(AObj: TColorListBox): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function ColorListBox_GetControlCount(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function ColorListBox_GetHandle(AObj: TColorListBox): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ColorListBox_GetParentWindow(AObj: TColorListBox): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ColorListBox_SetParentWindow(AObj: TColorListBox; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function ColorListBox_GetAction(AObj: TColorListBox): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ColorListBox_SetAction(AObj: TColorListBox; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure ColorListBox_GetBoundsRect(AObj: TColorListBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ColorListBox_SetBoundsRect(AObj: TColorListBox; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ColorListBox_GetClientHeight(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ColorListBox_SetClientHeight(AObj: TColorListBox; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ColorListBox_GetClientRect(AObj: TColorListBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ColorListBox_GetClientWidth(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ColorListBox_SetClientWidth(AObj: TColorListBox; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ColorListBox_GetExplicitLeft(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ColorListBox_GetExplicitTop(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ColorListBox_GetExplicitWidth(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ColorListBox_GetExplicitHeight(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ColorListBox_GetParent(AObj: TColorListBox): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ColorListBox_SetParent(AObj: TColorListBox; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ColorListBox_GetAlignWithMargins(AObj: TColorListBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ColorListBox_SetAlignWithMargins(AObj: TColorListBox; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ColorListBox_GetLeft(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ColorListBox_SetLeft(AObj: TColorListBox; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ColorListBox_GetTop(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ColorListBox_SetTop(AObj: TColorListBox; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ColorListBox_GetWidth(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ColorListBox_SetWidth(AObj: TColorListBox; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ColorListBox_GetHeight(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ColorListBox_SetHeight(AObj: TColorListBox; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ColorListBox_GetCursor(AObj: TColorListBox): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ColorListBox_SetCursor(AObj: TColorListBox; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ColorListBox_GetHint(AObj: TColorListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ColorListBox_SetHint(AObj: TColorListBox; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ColorListBox_GetMargins(AObj: TColorListBox): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ColorListBox_SetMargins(AObj: TColorListBox; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ColorListBox_GetCustomHint(AObj: TColorListBox): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ColorListBox_SetCustomHint(AObj: TColorListBox; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ColorListBox_GetComponentCount(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ColorListBox_GetComponentIndex(AObj: TColorListBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ColorListBox_SetComponentIndex(AObj: TColorListBox; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ColorListBox_GetOwner(AObj: TColorListBox): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ColorListBox_GetName(AObj: TColorListBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ColorListBox_SetName(AObj: TColorListBox; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ColorListBox_GetTag(AObj: TColorListBox): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ColorListBox_SetTag(AObj: TColorListBox; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ColorListBox_GetColors(AObj: TColorListBox; Index: Integer): TColor; stdcall;
begin
  Result :=  AObj.Colors[Index];
end;

function ColorListBox_GetColorNames(AObj: TColorListBox; Index: Integer): PChar; stdcall;
begin
  Result :=  PChar(AObj.ColorNames[Index]);
end;

function ColorListBox_GetControls(AObj: TColorListBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ColorListBox_GetComponents(AObj: TColorListBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ColorListBox_Create,
  ColorListBox_Free,
  ColorListBox_AddItem,
  ColorListBox_Clear,
  ColorListBox_ClearSelection,
  ColorListBox_DeleteSelected,
  ColorListBox_SelectAll,
  ColorListBox_CanFocus,
  ColorListBox_FlipChildren,
  ColorListBox_Focused,
  ColorListBox_HandleAllocated,
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
  ColorListBox_GetAutoComplete,
  ColorListBox_SetAutoComplete,
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
  ColorListBox_GetBevelEdges,
  ColorListBox_SetBevelEdges,
  ColorListBox_GetBevelInner,
  ColorListBox_SetBevelInner,
  ColorListBox_GetBevelKind,
  ColorListBox_SetBevelKind,
  ColorListBox_GetBevelOuter,
  ColorListBox_SetBevelOuter,
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
  ColorListBox_GetParentCtl3D,
  ColorListBox_SetParentCtl3D,
  ColorListBox_GetParentDoubleBuffered,
  ColorListBox_SetParentDoubleBuffered,
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
  ColorListBox_GetStyleElements,
  ColorListBox_SetStyleElements,
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
  ColorListBox_GetAutoCompleteDelay,
  ColorListBox_SetAutoCompleteDelay,
  ColorListBox_GetCanvas,
  ColorListBox_GetCount,
  ColorListBox_SetCount,
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
  ColorListBox_GetExplicitLeft,
  ColorListBox_GetExplicitTop,
  ColorListBox_GetExplicitWidth,
  ColorListBox_GetExplicitHeight,
  ColorListBox_GetParent,
  ColorListBox_SetParent,
  ColorListBox_GetAlignWithMargins,
  ColorListBox_SetAlignWithMargins,
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
  ColorListBox_GetMargins,
  ColorListBox_SetMargins,
  ColorListBox_GetCustomHint,
  ColorListBox_SetCustomHint,
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

