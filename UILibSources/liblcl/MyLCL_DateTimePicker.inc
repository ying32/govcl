
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function DateTimePicker_Create(AOwner: TComponent): TDateTimePicker; extdecl;
begin
  Result :=  TDateTimePicker.Create(AOwner);
end;

procedure DateTimePicker_Free(AObj: TDateTimePicker); extdecl;
begin
  AObj.Free;
end;

function DateTimePicker_CanFocus(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure DateTimePicker_FlipChildren(AObj: TDateTimePicker; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function DateTimePicker_Focused(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure DateTimePicker_Invalidate(AObj: TDateTimePicker); extdecl;
begin
  AObj.Invalidate;
end;

procedure DateTimePicker_Realign(AObj: TDateTimePicker); extdecl;
begin
  AObj.Realign;
end;

procedure DateTimePicker_Repaint(AObj: TDateTimePicker); extdecl;
begin
  AObj.Repaint;
end;

procedure DateTimePicker_ScaleBy(AObj: TDateTimePicker; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure DateTimePicker_SetBounds(AObj: TDateTimePicker; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure DateTimePicker_SetFocus(AObj: TDateTimePicker); extdecl;
begin
  AObj.SetFocus;
end;

procedure DateTimePicker_Update(AObj: TDateTimePicker); extdecl;
begin
  AObj.Update;
end;

procedure DateTimePicker_BringToFront(AObj: TDateTimePicker); extdecl;
begin
  AObj.BringToFront;
end;

function DateTimePicker_HasParent(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure DateTimePicker_Hide(AObj: TDateTimePicker); extdecl;
begin
  AObj.Hide;
end;

function DateTimePicker_Perform(AObj: TDateTimePicker; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure DateTimePicker_Refresh(AObj: TDateTimePicker); extdecl;
begin
  AObj.Refresh;
end;

procedure DateTimePicker_SendToBack(AObj: TDateTimePicker); extdecl;
begin
  AObj.SendToBack;
end;

procedure DateTimePicker_Show(AObj: TDateTimePicker); extdecl;
begin
  AObj.Show;
end;

function DateTimePicker_GetTextBuf(AObj: TDateTimePicker; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function DateTimePicker_GetTextLen(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function DateTimePicker_FindComponent(AObj: TDateTimePicker; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function DateTimePicker_GetNamePath(AObj: TDateTimePicker): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure DateTimePicker_Assign(AObj: TDateTimePicker; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function DateTimePicker_ClassName(AObj: TDateTimePicker): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function DateTimePicker_Equals(AObj: TDateTimePicker; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function DateTimePicker_GetHashCode(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function DateTimePicker_ToString(AObj: TDateTimePicker): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function DateTimePicker_GetDateTime(AObj: TDateTimePicker): TUnixDateTime; extdecl;
begin
  Result :=  ToUnixTime(AObj.DateTime);
end;

procedure DateTimePicker_SetDateTime(AObj: TDateTimePicker; AValue: TUnixDateTime); extdecl;
begin
  AObj.DateTime := UnixToTime(AValue);
end;

function DateTimePicker_GetDroppedDown(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.DroppedDown;
end;

function DateTimePicker_GetAlign(AObj: TDateTimePicker): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure DateTimePicker_SetAlign(AObj: TDateTimePicker; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function DateTimePicker_GetAnchors(AObj: TDateTimePicker): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure DateTimePicker_SetAnchors(AObj: TDateTimePicker; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function DateTimePicker_GetBiDiMode(AObj: TDateTimePicker): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure DateTimePicker_SetBiDiMode(AObj: TDateTimePicker; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function DateTimePicker_GetCalAlignment(AObj: TDateTimePicker): TDTCalAlignment; extdecl;
begin
  Result :=  AObj.CalAlignment;
end;

procedure DateTimePicker_SetCalAlignment(AObj: TDateTimePicker; AValue: TDTCalAlignment); extdecl;
begin
  AObj.CalAlignment := AValue;
end;

function DateTimePicker_GetDate(AObj: TDateTimePicker): TUnixDateTime; extdecl;
begin
  Result :=  ToUnixTime(AObj.Date);
end;

procedure DateTimePicker_SetDate(AObj: TDateTimePicker; AValue: TUnixDateTime); extdecl;
begin
  AObj.Date := UnixToTime(AValue);
end;

function DateTimePicker_GetTime(AObj: TDateTimePicker): TUnixDateTime; extdecl;
begin
  Result :=  ToUnixTime(AObj.Time);
end;

procedure DateTimePicker_SetTime(AObj: TDateTimePicker; AValue: TUnixDateTime); extdecl;
begin
  AObj.Time := UnixToTime(AValue);
end;

function DateTimePicker_GetChecked(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.Checked;
end;

procedure DateTimePicker_SetChecked(AObj: TDateTimePicker; AValue: LongBool); extdecl;
begin
  AObj.Checked := AValue;
end;

function DateTimePicker_GetColor(AObj: TDateTimePicker): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure DateTimePicker_SetColor(AObj: TDateTimePicker; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function DateTimePicker_GetDateMode(AObj: TDateTimePicker): TDTDateMode; extdecl;
begin
  Result :=  AObj.DateMode;
end;

procedure DateTimePicker_SetDateMode(AObj: TDateTimePicker; AValue: TDTDateMode); extdecl;
begin
  AObj.DateMode := AValue;
end;

function DateTimePicker_GetDoubleBuffered(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure DateTimePicker_SetDoubleBuffered(AObj: TDateTimePicker; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function DateTimePicker_GetEnabled(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure DateTimePicker_SetEnabled(AObj: TDateTimePicker; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function DateTimePicker_GetFont(AObj: TDateTimePicker): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure DateTimePicker_SetFont(AObj: TDateTimePicker; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function DateTimePicker_GetKind(AObj: TDateTimePicker): TDateTimeKind; extdecl;
begin
  Result :=  AObj.Kind;
end;

procedure DateTimePicker_SetKind(AObj: TDateTimePicker; AValue: TDateTimeKind); extdecl;
begin
  AObj.Kind := AValue;
end;

function DateTimePicker_GetMaxDate(AObj: TDateTimePicker): TUnixDateTime; extdecl;
begin
  Result :=  ToUnixTime(AObj.MaxDate);
end;

procedure DateTimePicker_SetMaxDate(AObj: TDateTimePicker; AValue: TUnixDateTime); extdecl;
begin
  AObj.MaxDate := UnixToTime(AValue);
end;

function DateTimePicker_GetMinDate(AObj: TDateTimePicker): TUnixDateTime; extdecl;
begin
  Result :=  ToUnixTime(AObj.MinDate);
end;

procedure DateTimePicker_SetMinDate(AObj: TDateTimePicker; AValue: TUnixDateTime); extdecl;
begin
  AObj.MinDate := UnixToTime(AValue);
end;

function DateTimePicker_GetParentColor(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure DateTimePicker_SetParentColor(AObj: TDateTimePicker; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function DateTimePicker_GetParentFont(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure DateTimePicker_SetParentFont(AObj: TDateTimePicker; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function DateTimePicker_GetParentShowHint(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure DateTimePicker_SetParentShowHint(AObj: TDateTimePicker; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function DateTimePicker_GetPopupMenu(AObj: TDateTimePicker): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure DateTimePicker_SetPopupMenu(AObj: TDateTimePicker; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function DateTimePicker_GetShowHint(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure DateTimePicker_SetShowHint(AObj: TDateTimePicker; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function DateTimePicker_GetTabOrder(AObj: TDateTimePicker): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure DateTimePicker_SetTabOrder(AObj: TDateTimePicker; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function DateTimePicker_GetTabStop(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure DateTimePicker_SetTabStop(AObj: TDateTimePicker; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function DateTimePicker_GetVisible(AObj: TDateTimePicker): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure DateTimePicker_SetVisible(AObj: TDateTimePicker; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure DateTimePicker_SetOnClick(AObj: TDateTimePicker; AEventId: NativeUInt); extdecl;
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

procedure DateTimePicker_SetOnChange(AObj: TDateTimePicker; AEventId: NativeUInt); extdecl;
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

procedure DateTimePicker_SetOnEnter(AObj: TDateTimePicker; AEventId: NativeUInt); extdecl;
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

procedure DateTimePicker_SetOnExit(AObj: TDateTimePicker; AEventId: NativeUInt); extdecl;
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

procedure DateTimePicker_SetOnKeyDown(AObj: TDateTimePicker; AEventId: NativeUInt); extdecl;
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

procedure DateTimePicker_SetOnKeyPress(AObj: TDateTimePicker; AEventId: NativeUInt); extdecl;
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

procedure DateTimePicker_SetOnKeyUp(AObj: TDateTimePicker; AEventId: NativeUInt); extdecl;
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

procedure DateTimePicker_SetOnMouseEnter(AObj: TDateTimePicker; AEventId: NativeUInt); extdecl;
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

procedure DateTimePicker_SetOnMouseLeave(AObj: TDateTimePicker; AEventId: NativeUInt); extdecl;
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

function DateTimePicker_GetBrush(AObj: TDateTimePicker): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function DateTimePicker_GetControlCount(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function DateTimePicker_GetHandle(AObj: TDateTimePicker): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function DateTimePicker_GetParentWindow(AObj: TDateTimePicker): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure DateTimePicker_SetParentWindow(AObj: TDateTimePicker; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function DateTimePicker_GetAction(AObj: TDateTimePicker): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure DateTimePicker_SetAction(AObj: TDateTimePicker; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure DateTimePicker_GetBoundsRect(AObj: TDateTimePicker; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure DateTimePicker_SetBoundsRect(AObj: TDateTimePicker; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function DateTimePicker_GetClientHeight(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure DateTimePicker_SetClientHeight(AObj: TDateTimePicker; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure DateTimePicker_GetClientRect(AObj: TDateTimePicker; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function DateTimePicker_GetClientWidth(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure DateTimePicker_SetClientWidth(AObj: TDateTimePicker; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function DateTimePicker_GetParent(AObj: TDateTimePicker): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure DateTimePicker_SetParent(AObj: TDateTimePicker; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function DateTimePicker_GetLeft(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure DateTimePicker_SetLeft(AObj: TDateTimePicker; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function DateTimePicker_GetTop(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure DateTimePicker_SetTop(AObj: TDateTimePicker; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function DateTimePicker_GetWidth(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure DateTimePicker_SetWidth(AObj: TDateTimePicker; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function DateTimePicker_GetHeight(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure DateTimePicker_SetHeight(AObj: TDateTimePicker; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function DateTimePicker_GetCursor(AObj: TDateTimePicker): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure DateTimePicker_SetCursor(AObj: TDateTimePicker; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function DateTimePicker_GetHint(AObj: TDateTimePicker): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure DateTimePicker_SetHint(AObj: TDateTimePicker; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function DateTimePicker_GetComponentCount(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function DateTimePicker_GetComponentIndex(AObj: TDateTimePicker): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure DateTimePicker_SetComponentIndex(AObj: TDateTimePicker; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function DateTimePicker_GetOwner(AObj: TDateTimePicker): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function DateTimePicker_GetName(AObj: TDateTimePicker): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure DateTimePicker_SetName(AObj: TDateTimePicker; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function DateTimePicker_GetTag(AObj: TDateTimePicker): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure DateTimePicker_SetTag(AObj: TDateTimePicker; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function DateTimePicker_GetControls(AObj: TDateTimePicker; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function DateTimePicker_GetComponents(AObj: TDateTimePicker; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  DateTimePicker_Create,
  DateTimePicker_Free,
  DateTimePicker_CanFocus,
  DateTimePicker_FlipChildren,
  DateTimePicker_Focused,
  DateTimePicker_Invalidate,
  DateTimePicker_Realign,
  DateTimePicker_Repaint,
  DateTimePicker_ScaleBy,
  DateTimePicker_SetBounds,
  DateTimePicker_SetFocus,
  DateTimePicker_Update,
  DateTimePicker_BringToFront,
  DateTimePicker_HasParent,
  DateTimePicker_Hide,
  DateTimePicker_Perform,
  DateTimePicker_Refresh,
  DateTimePicker_SendToBack,
  DateTimePicker_Show,
  DateTimePicker_GetTextBuf,
  DateTimePicker_GetTextLen,
  DateTimePicker_FindComponent,
  DateTimePicker_GetNamePath,
  DateTimePicker_Assign,
  DateTimePicker_ClassName,
  DateTimePicker_Equals,
  DateTimePicker_GetHashCode,
  DateTimePicker_ToString,
  DateTimePicker_GetDateTime,
  DateTimePicker_SetDateTime,
  DateTimePicker_GetDroppedDown,
  DateTimePicker_GetAlign,
  DateTimePicker_SetAlign,
  DateTimePicker_GetAnchors,
  DateTimePicker_SetAnchors,
  DateTimePicker_GetBiDiMode,
  DateTimePicker_SetBiDiMode,
  DateTimePicker_GetCalAlignment,
  DateTimePicker_SetCalAlignment,
  DateTimePicker_GetDate,
  DateTimePicker_SetDate,
  DateTimePicker_GetTime,
  DateTimePicker_SetTime,
  DateTimePicker_GetChecked,
  DateTimePicker_SetChecked,
  DateTimePicker_GetColor,
  DateTimePicker_SetColor,
  DateTimePicker_GetDateMode,
  DateTimePicker_SetDateMode,
  DateTimePicker_GetDoubleBuffered,
  DateTimePicker_SetDoubleBuffered,
  DateTimePicker_GetEnabled,
  DateTimePicker_SetEnabled,
  DateTimePicker_GetFont,
  DateTimePicker_SetFont,
  DateTimePicker_GetKind,
  DateTimePicker_SetKind,
  DateTimePicker_GetMaxDate,
  DateTimePicker_SetMaxDate,
  DateTimePicker_GetMinDate,
  DateTimePicker_SetMinDate,
  DateTimePicker_GetParentColor,
  DateTimePicker_SetParentColor,
  DateTimePicker_GetParentFont,
  DateTimePicker_SetParentFont,
  DateTimePicker_GetParentShowHint,
  DateTimePicker_SetParentShowHint,
  DateTimePicker_GetPopupMenu,
  DateTimePicker_SetPopupMenu,
  DateTimePicker_GetShowHint,
  DateTimePicker_SetShowHint,
  DateTimePicker_GetTabOrder,
  DateTimePicker_SetTabOrder,
  DateTimePicker_GetTabStop,
  DateTimePicker_SetTabStop,
  DateTimePicker_GetVisible,
  DateTimePicker_SetVisible,
  DateTimePicker_SetOnClick,
  DateTimePicker_SetOnChange,
  DateTimePicker_SetOnEnter,
  DateTimePicker_SetOnExit,
  DateTimePicker_SetOnKeyDown,
  DateTimePicker_SetOnKeyPress,
  DateTimePicker_SetOnKeyUp,
  DateTimePicker_SetOnMouseEnter,
  DateTimePicker_SetOnMouseLeave,
  DateTimePicker_GetBrush,
  DateTimePicker_GetControlCount,
  DateTimePicker_GetHandle,
  DateTimePicker_GetParentWindow,
  DateTimePicker_SetParentWindow,
  DateTimePicker_GetAction,
  DateTimePicker_SetAction,
  DateTimePicker_GetBoundsRect,
  DateTimePicker_SetBoundsRect,
  DateTimePicker_GetClientHeight,
  DateTimePicker_SetClientHeight,
  DateTimePicker_GetClientRect,
  DateTimePicker_GetClientWidth,
  DateTimePicker_SetClientWidth,
  DateTimePicker_GetParent,
  DateTimePicker_SetParent,
  DateTimePicker_GetLeft,
  DateTimePicker_SetLeft,
  DateTimePicker_GetTop,
  DateTimePicker_SetTop,
  DateTimePicker_GetWidth,
  DateTimePicker_SetWidth,
  DateTimePicker_GetHeight,
  DateTimePicker_SetHeight,
  DateTimePicker_GetCursor,
  DateTimePicker_SetCursor,
  DateTimePicker_GetHint,
  DateTimePicker_SetHint,
  DateTimePicker_GetComponentCount,
  DateTimePicker_GetComponentIndex,
  DateTimePicker_SetComponentIndex,
  DateTimePicker_GetOwner,
  DateTimePicker_GetName,
  DateTimePicker_SetName,
  DateTimePicker_GetTag,
  DateTimePicker_SetTag,
  DateTimePicker_GetControls,
  DateTimePicker_GetComponents;

