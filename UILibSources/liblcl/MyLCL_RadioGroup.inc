
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function RadioGroup_Create(AOwner: TComponent): TRadioGroup; extdecl;
begin
  Result :=  TRadioGroup.Create(AOwner);
end;

procedure RadioGroup_Free(AObj: TRadioGroup); extdecl;
begin
  AObj.Free;
end;

procedure RadioGroup_FlipChildren(AObj: TRadioGroup; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function RadioGroup_CanFocus(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

function RadioGroup_Focused(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure RadioGroup_Invalidate(AObj: TRadioGroup); extdecl;
begin
  AObj.Invalidate;
end;

procedure RadioGroup_Realign(AObj: TRadioGroup); extdecl;
begin
  AObj.Realign;
end;

procedure RadioGroup_Repaint(AObj: TRadioGroup); extdecl;
begin
  AObj.Repaint;
end;

procedure RadioGroup_ScaleBy(AObj: TRadioGroup; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure RadioGroup_SetBounds(AObj: TRadioGroup; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure RadioGroup_SetFocus(AObj: TRadioGroup); extdecl;
begin
  AObj.SetFocus;
end;

procedure RadioGroup_Update(AObj: TRadioGroup); extdecl;
begin
  AObj.Update;
end;

procedure RadioGroup_BringToFront(AObj: TRadioGroup); extdecl;
begin
  AObj.BringToFront;
end;

function RadioGroup_HasParent(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure RadioGroup_Hide(AObj: TRadioGroup); extdecl;
begin
  AObj.Hide;
end;

function RadioGroup_Perform(AObj: TRadioGroup; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure RadioGroup_Refresh(AObj: TRadioGroup); extdecl;
begin
  AObj.Refresh;
end;

procedure RadioGroup_SendToBack(AObj: TRadioGroup); extdecl;
begin
  AObj.SendToBack;
end;

procedure RadioGroup_Show(AObj: TRadioGroup); extdecl;
begin
  AObj.Show;
end;

function RadioGroup_GetTextBuf(AObj: TRadioGroup; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function RadioGroup_GetTextLen(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function RadioGroup_FindComponent(AObj: TRadioGroup; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function RadioGroup_GetNamePath(AObj: TRadioGroup): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure RadioGroup_Assign(AObj: TRadioGroup; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function RadioGroup_ClassName(AObj: TRadioGroup): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function RadioGroup_Equals(AObj: TRadioGroup; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function RadioGroup_GetHashCode(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function RadioGroup_ToString(AObj: TRadioGroup): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function RadioGroup_GetAlign(AObj: TRadioGroup): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure RadioGroup_SetAlign(AObj: TRadioGroup; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function RadioGroup_GetAnchors(AObj: TRadioGroup): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure RadioGroup_SetAnchors(AObj: TRadioGroup; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function RadioGroup_GetBiDiMode(AObj: TRadioGroup): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure RadioGroup_SetBiDiMode(AObj: TRadioGroup; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function RadioGroup_GetCaption(AObj: TRadioGroup): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure RadioGroup_SetCaption(AObj: TRadioGroup; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function RadioGroup_GetColor(AObj: TRadioGroup): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure RadioGroup_SetColor(AObj: TRadioGroup; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function RadioGroup_GetColumns(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.Columns;
end;

procedure RadioGroup_SetColumns(AObj: TRadioGroup; AValue: Integer); extdecl;
begin
  AObj.Columns := AValue;
end;

function RadioGroup_GetDoubleBuffered(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure RadioGroup_SetDoubleBuffered(AObj: TRadioGroup; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function RadioGroup_GetEnabled(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure RadioGroup_SetEnabled(AObj: TRadioGroup; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function RadioGroup_GetFont(AObj: TRadioGroup): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure RadioGroup_SetFont(AObj: TRadioGroup; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function RadioGroup_GetItemIndex(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.ItemIndex;
end;

procedure RadioGroup_SetItemIndex(AObj: TRadioGroup; AValue: Integer); extdecl;
begin
  AObj.ItemIndex := AValue;
end;

function RadioGroup_GetItems(AObj: TRadioGroup): TStrings; extdecl;
begin
  Result :=  AObj.Items;
end;

procedure RadioGroup_SetItems(AObj: TRadioGroup; AValue: TStrings); extdecl;
begin
  AObj.Items := AValue;
end;

function RadioGroup_GetParentColor(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure RadioGroup_SetParentColor(AObj: TRadioGroup; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function RadioGroup_GetParentFont(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure RadioGroup_SetParentFont(AObj: TRadioGroup; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function RadioGroup_GetParentShowHint(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure RadioGroup_SetParentShowHint(AObj: TRadioGroup; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function RadioGroup_GetPopupMenu(AObj: TRadioGroup): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure RadioGroup_SetPopupMenu(AObj: TRadioGroup; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function RadioGroup_GetShowHint(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure RadioGroup_SetShowHint(AObj: TRadioGroup; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function RadioGroup_GetTabOrder(AObj: TRadioGroup): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure RadioGroup_SetTabOrder(AObj: TRadioGroup; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function RadioGroup_GetTabStop(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure RadioGroup_SetTabStop(AObj: TRadioGroup; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function RadioGroup_GetVisible(AObj: TRadioGroup): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure RadioGroup_SetVisible(AObj: TRadioGroup; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure RadioGroup_SetOnClick(AObj: TRadioGroup; AEventId: NativeUInt); extdecl;
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

procedure RadioGroup_SetOnEnter(AObj: TRadioGroup; AEventId: NativeUInt); extdecl;
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

procedure RadioGroup_SetOnExit(AObj: TRadioGroup; AEventId: NativeUInt); extdecl;
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

function RadioGroup_GetBrush(AObj: TRadioGroup): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function RadioGroup_GetControlCount(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function RadioGroup_GetHandle(AObj: TRadioGroup): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function RadioGroup_GetParentWindow(AObj: TRadioGroup): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure RadioGroup_SetParentWindow(AObj: TRadioGroup; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function RadioGroup_GetAction(AObj: TRadioGroup): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure RadioGroup_SetAction(AObj: TRadioGroup; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure RadioGroup_GetBoundsRect(AObj: TRadioGroup; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure RadioGroup_SetBoundsRect(AObj: TRadioGroup; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function RadioGroup_GetClientHeight(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure RadioGroup_SetClientHeight(AObj: TRadioGroup; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure RadioGroup_GetClientRect(AObj: TRadioGroup; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function RadioGroup_GetClientWidth(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure RadioGroup_SetClientWidth(AObj: TRadioGroup; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function RadioGroup_GetParent(AObj: TRadioGroup): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure RadioGroup_SetParent(AObj: TRadioGroup; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function RadioGroup_GetLeft(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure RadioGroup_SetLeft(AObj: TRadioGroup; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function RadioGroup_GetTop(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure RadioGroup_SetTop(AObj: TRadioGroup; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function RadioGroup_GetWidth(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure RadioGroup_SetWidth(AObj: TRadioGroup; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function RadioGroup_GetHeight(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure RadioGroup_SetHeight(AObj: TRadioGroup; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function RadioGroup_GetCursor(AObj: TRadioGroup): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure RadioGroup_SetCursor(AObj: TRadioGroup; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function RadioGroup_GetHint(AObj: TRadioGroup): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure RadioGroup_SetHint(AObj: TRadioGroup; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function RadioGroup_GetComponentCount(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function RadioGroup_GetComponentIndex(AObj: TRadioGroup): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure RadioGroup_SetComponentIndex(AObj: TRadioGroup; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function RadioGroup_GetOwner(AObj: TRadioGroup): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function RadioGroup_GetName(AObj: TRadioGroup): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure RadioGroup_SetName(AObj: TRadioGroup; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function RadioGroup_GetTag(AObj: TRadioGroup): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure RadioGroup_SetTag(AObj: TRadioGroup; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function RadioGroup_GetControls(AObj: TRadioGroup; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function RadioGroup_GetComponents(AObj: TRadioGroup; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  RadioGroup_Create,
  RadioGroup_Free,
  RadioGroup_FlipChildren,
  RadioGroup_CanFocus,
  RadioGroup_Focused,
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
  RadioGroup_GetParentColor,
  RadioGroup_SetParentColor,
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
  RadioGroup_GetParent,
  RadioGroup_SetParent,
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
  RadioGroup_GetComponentCount,
  RadioGroup_GetComponentIndex,
  RadioGroup_SetComponentIndex,
  RadioGroup_GetOwner,
  RadioGroup_GetName,
  RadioGroup_SetName,
  RadioGroup_GetTag,
  RadioGroup_SetTag,
  RadioGroup_GetControls,
  RadioGroup_GetComponents;

