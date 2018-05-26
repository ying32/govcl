
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function RadioGroup_Create(AOwner: TComponent): TRadioGroup; stdcall;
begin
  Result :=  TRadioGroup.Create(AOwner);
end;

procedure RadioGroup_Free(AObj: TRadioGroup); stdcall;
begin
  AObj.Free;
end;

procedure RadioGroup_FlipChildren(AObj: TRadioGroup; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function RadioGroup_CanFocus(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function RadioGroup_Focused(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function RadioGroup_HandleAllocated(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure RadioGroup_Invalidate(AObj: TRadioGroup); stdcall;
begin
  AObj.Invalidate;
end;

procedure RadioGroup_Realign(AObj: TRadioGroup); stdcall;
begin
  AObj.Realign;
end;

procedure RadioGroup_Repaint(AObj: TRadioGroup); stdcall;
begin
  AObj.Repaint;
end;

procedure RadioGroup_ScaleBy(AObj: TRadioGroup; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure RadioGroup_SetBounds(AObj: TRadioGroup; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure RadioGroup_SetFocus(AObj: TRadioGroup); stdcall;
begin
  AObj.SetFocus;
end;

procedure RadioGroup_Update(AObj: TRadioGroup); stdcall;
begin
  AObj.Update;
end;

procedure RadioGroup_BringToFront(AObj: TRadioGroup); stdcall;
begin
  AObj.BringToFront;
end;

function RadioGroup_HasParent(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure RadioGroup_Hide(AObj: TRadioGroup); stdcall;
begin
  AObj.Hide;
end;

function RadioGroup_Perform(AObj: TRadioGroup; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure RadioGroup_Refresh(AObj: TRadioGroup); stdcall;
begin
  AObj.Refresh;
end;

procedure RadioGroup_SendToBack(AObj: TRadioGroup); stdcall;
begin
  AObj.SendToBack;
end;

procedure RadioGroup_Show(AObj: TRadioGroup); stdcall;
begin
  AObj.Show;
end;

function RadioGroup_GetTextBuf(AObj: TRadioGroup; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function RadioGroup_GetTextLen(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function RadioGroup_FindComponent(AObj: TRadioGroup; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function RadioGroup_GetNamePath(AObj: TRadioGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure RadioGroup_Assign(AObj: TRadioGroup; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function RadioGroup_ClassName(AObj: TRadioGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function RadioGroup_Equals(AObj: TRadioGroup; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function RadioGroup_GetHashCode(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function RadioGroup_ToString(AObj: TRadioGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function RadioGroup_GetAlign(AObj: TRadioGroup): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure RadioGroup_SetAlign(AObj: TRadioGroup; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function RadioGroup_GetAnchors(AObj: TRadioGroup): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure RadioGroup_SetAnchors(AObj: TRadioGroup; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function RadioGroup_GetBiDiMode(AObj: TRadioGroup): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure RadioGroup_SetBiDiMode(AObj: TRadioGroup; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function RadioGroup_GetCaption(AObj: TRadioGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure RadioGroup_SetCaption(AObj: TRadioGroup; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function RadioGroup_GetColor(AObj: TRadioGroup): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure RadioGroup_SetColor(AObj: TRadioGroup; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function RadioGroup_GetColumns(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.Columns;
end;

procedure RadioGroup_SetColumns(AObj: TRadioGroup; AValue: Integer); stdcall;
begin
  AObj.Columns := AValue;
end;

function RadioGroup_GetDoubleBuffered(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure RadioGroup_SetDoubleBuffered(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function RadioGroup_GetEnabled(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure RadioGroup_SetEnabled(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function RadioGroup_GetFont(AObj: TRadioGroup): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure RadioGroup_SetFont(AObj: TRadioGroup; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function RadioGroup_GetItemIndex(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.ItemIndex;
end;

procedure RadioGroup_SetItemIndex(AObj: TRadioGroup; AValue: Integer); stdcall;
begin
  AObj.ItemIndex := AValue;
end;

function RadioGroup_GetItems(AObj: TRadioGroup): TStrings; stdcall;
begin
  Result :=  AObj.Items;
end;

procedure RadioGroup_SetItems(AObj: TRadioGroup; AValue: TStrings); stdcall;
begin
  AObj.Items := AValue;
end;

function RadioGroup_GetParentBackground(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentBackground;
end;

procedure RadioGroup_SetParentBackground(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentBackground := AValue;
end;

function RadioGroup_GetParentColor(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure RadioGroup_SetParentColor(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function RadioGroup_GetParentCtl3D(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure RadioGroup_SetParentCtl3D(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function RadioGroup_GetParentDoubleBuffered(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure RadioGroup_SetParentDoubleBuffered(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function RadioGroup_GetParentFont(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure RadioGroup_SetParentFont(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function RadioGroup_GetParentShowHint(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure RadioGroup_SetParentShowHint(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function RadioGroup_GetPopupMenu(AObj: TRadioGroup): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure RadioGroup_SetPopupMenu(AObj: TRadioGroup; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function RadioGroup_GetShowHint(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure RadioGroup_SetShowHint(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function RadioGroup_GetTabOrder(AObj: TRadioGroup): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure RadioGroup_SetTabOrder(AObj: TRadioGroup; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function RadioGroup_GetTabStop(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure RadioGroup_SetTabStop(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function RadioGroup_GetVisible(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure RadioGroup_SetVisible(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function RadioGroup_GetStyleElements(AObj: TRadioGroup): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure RadioGroup_SetStyleElements(AObj: TRadioGroup; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function RadioGroup_GetWordWrap(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.WordWrap;
end;

procedure RadioGroup_SetWordWrap(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.WordWrap := AValue;
end;

procedure RadioGroup_SetOnClick(AObj: TRadioGroup; AEventId: NativeUInt); stdcall;
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

procedure RadioGroup_SetOnEnter(AObj: TRadioGroup; AEventId: NativeUInt); stdcall;
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

procedure RadioGroup_SetOnExit(AObj: TRadioGroup; AEventId: NativeUInt); stdcall;
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

function RadioGroup_GetBrush(AObj: TRadioGroup): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function RadioGroup_GetControlCount(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function RadioGroup_GetHandle(AObj: TRadioGroup): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function RadioGroup_GetParentWindow(AObj: TRadioGroup): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure RadioGroup_SetParentWindow(AObj: TRadioGroup; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function RadioGroup_GetAction(AObj: TRadioGroup): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure RadioGroup_SetAction(AObj: TRadioGroup; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure RadioGroup_GetBoundsRect(AObj: TRadioGroup; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure RadioGroup_SetBoundsRect(AObj: TRadioGroup; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function RadioGroup_GetClientHeight(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure RadioGroup_SetClientHeight(AObj: TRadioGroup; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure RadioGroup_GetClientRect(AObj: TRadioGroup; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function RadioGroup_GetClientWidth(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure RadioGroup_SetClientWidth(AObj: TRadioGroup; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function RadioGroup_GetExplicitLeft(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function RadioGroup_GetExplicitTop(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function RadioGroup_GetExplicitWidth(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function RadioGroup_GetExplicitHeight(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function RadioGroup_GetParent(AObj: TRadioGroup): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure RadioGroup_SetParent(AObj: TRadioGroup; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function RadioGroup_GetAlignWithMargins(AObj: TRadioGroup): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure RadioGroup_SetAlignWithMargins(AObj: TRadioGroup; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function RadioGroup_GetLeft(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure RadioGroup_SetLeft(AObj: TRadioGroup; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function RadioGroup_GetTop(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure RadioGroup_SetTop(AObj: TRadioGroup; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function RadioGroup_GetWidth(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure RadioGroup_SetWidth(AObj: TRadioGroup; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function RadioGroup_GetHeight(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure RadioGroup_SetHeight(AObj: TRadioGroup; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function RadioGroup_GetCursor(AObj: TRadioGroup): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure RadioGroup_SetCursor(AObj: TRadioGroup; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function RadioGroup_GetHint(AObj: TRadioGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure RadioGroup_SetHint(AObj: TRadioGroup; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function RadioGroup_GetMargins(AObj: TRadioGroup): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure RadioGroup_SetMargins(AObj: TRadioGroup; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function RadioGroup_GetCustomHint(AObj: TRadioGroup): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure RadioGroup_SetCustomHint(AObj: TRadioGroup; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function RadioGroup_GetComponentCount(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function RadioGroup_GetComponentIndex(AObj: TRadioGroup): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure RadioGroup_SetComponentIndex(AObj: TRadioGroup; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function RadioGroup_GetOwner(AObj: TRadioGroup): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function RadioGroup_GetName(AObj: TRadioGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure RadioGroup_SetName(AObj: TRadioGroup; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function RadioGroup_GetTag(AObj: TRadioGroup): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure RadioGroup_SetTag(AObj: TRadioGroup; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function RadioGroup_GetButtons(AObj: TRadioGroup; Index: Integer): TRadioButton; stdcall;
begin
  Result :=  AObj.Buttons[Index];
end;

function RadioGroup_GetControls(AObj: TRadioGroup; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function RadioGroup_GetComponents(AObj: TRadioGroup; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  RadioGroup_Create,
  RadioGroup_Free,
  RadioGroup_FlipChildren,
  RadioGroup_CanFocus,
  RadioGroup_Focused,
  RadioGroup_HandleAllocated,
  RadioGroup_Invalidate,
  RadioGroup_Realign,
  RadioGroup_Repaint,
  RadioGroup_ScaleBy,
  RadioGroup_SetBounds,
  RadioGroup_SetFocus,
  RadioGroup_Update,
  RadioGroup_BringToFront,
  RadioGroup_HasParent,
  RadioGroup_Hide,
  RadioGroup_Perform,
  RadioGroup_Refresh,
  RadioGroup_SendToBack,
  RadioGroup_Show,
  RadioGroup_GetTextBuf,
  RadioGroup_GetTextLen,
  RadioGroup_FindComponent,
  RadioGroup_GetNamePath,
  RadioGroup_Assign,
  RadioGroup_ClassName,
  RadioGroup_Equals,
  RadioGroup_GetHashCode,
  RadioGroup_ToString,
  RadioGroup_GetAlign,
  RadioGroup_SetAlign,
  RadioGroup_GetAnchors,
  RadioGroup_SetAnchors,
  RadioGroup_GetBiDiMode,
  RadioGroup_SetBiDiMode,
  RadioGroup_GetCaption,
  RadioGroup_SetCaption,
  RadioGroup_GetColor,
  RadioGroup_SetColor,
  RadioGroup_GetColumns,
  RadioGroup_SetColumns,
  RadioGroup_GetDoubleBuffered,
  RadioGroup_SetDoubleBuffered,
  RadioGroup_GetEnabled,
  RadioGroup_SetEnabled,
  RadioGroup_GetFont,
  RadioGroup_SetFont,
  RadioGroup_GetItemIndex,
  RadioGroup_SetItemIndex,
  RadioGroup_GetItems,
  RadioGroup_SetItems,
  RadioGroup_GetParentBackground,
  RadioGroup_SetParentBackground,
  RadioGroup_GetParentColor,
  RadioGroup_SetParentColor,
  RadioGroup_GetParentCtl3D,
  RadioGroup_SetParentCtl3D,
  RadioGroup_GetParentDoubleBuffered,
  RadioGroup_SetParentDoubleBuffered,
  RadioGroup_GetParentFont,
  RadioGroup_SetParentFont,
  RadioGroup_GetParentShowHint,
  RadioGroup_SetParentShowHint,
  RadioGroup_GetPopupMenu,
  RadioGroup_SetPopupMenu,
  RadioGroup_GetShowHint,
  RadioGroup_SetShowHint,
  RadioGroup_GetTabOrder,
  RadioGroup_SetTabOrder,
  RadioGroup_GetTabStop,
  RadioGroup_SetTabStop,
  RadioGroup_GetVisible,
  RadioGroup_SetVisible,
  RadioGroup_GetStyleElements,
  RadioGroup_SetStyleElements,
  RadioGroup_GetWordWrap,
  RadioGroup_SetWordWrap,
  RadioGroup_SetOnClick,
  RadioGroup_SetOnEnter,
  RadioGroup_SetOnExit,
  RadioGroup_GetBrush,
  RadioGroup_GetControlCount,
  RadioGroup_GetHandle,
  RadioGroup_GetParentWindow,
  RadioGroup_SetParentWindow,
  RadioGroup_GetAction,
  RadioGroup_SetAction,
  RadioGroup_GetBoundsRect,
  RadioGroup_SetBoundsRect,
  RadioGroup_GetClientHeight,
  RadioGroup_SetClientHeight,
  RadioGroup_GetClientRect,
  RadioGroup_GetClientWidth,
  RadioGroup_SetClientWidth,
  RadioGroup_GetExplicitLeft,
  RadioGroup_GetExplicitTop,
  RadioGroup_GetExplicitWidth,
  RadioGroup_GetExplicitHeight,
  RadioGroup_GetParent,
  RadioGroup_SetParent,
  RadioGroup_GetAlignWithMargins,
  RadioGroup_SetAlignWithMargins,
  RadioGroup_GetLeft,
  RadioGroup_SetLeft,
  RadioGroup_GetTop,
  RadioGroup_SetTop,
  RadioGroup_GetWidth,
  RadioGroup_SetWidth,
  RadioGroup_GetHeight,
  RadioGroup_SetHeight,
  RadioGroup_GetCursor,
  RadioGroup_SetCursor,
  RadioGroup_GetHint,
  RadioGroup_SetHint,
  RadioGroup_GetMargins,
  RadioGroup_SetMargins,
  RadioGroup_GetCustomHint,
  RadioGroup_SetCustomHint,
  RadioGroup_GetComponentCount,
  RadioGroup_GetComponentIndex,
  RadioGroup_SetComponentIndex,
  RadioGroup_GetOwner,
  RadioGroup_GetName,
  RadioGroup_SetName,
  RadioGroup_GetTag,
  RadioGroup_SetTag,
  RadioGroup_GetButtons,
  RadioGroup_GetControls,
  RadioGroup_GetComponents;

