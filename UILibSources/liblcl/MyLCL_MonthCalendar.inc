
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function MonthCalendar_Create(AOwner: TComponent): TMonthCalendar; extdecl;
begin
  Result :=  TMonthCalendar.Create(AOwner);
end;

procedure MonthCalendar_Free(AObj: TMonthCalendar); extdecl;
begin
  AObj.Free;
end;

function MonthCalendar_CanFocus(AObj: TMonthCalendar): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure MonthCalendar_FlipChildren(AObj: TMonthCalendar; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function MonthCalendar_Focused(AObj: TMonthCalendar): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure MonthCalendar_Invalidate(AObj: TMonthCalendar); extdecl;
begin
  AObj.Invalidate;
end;

procedure MonthCalendar_Realign(AObj: TMonthCalendar); extdecl;
begin
  AObj.Realign;
end;

procedure MonthCalendar_Repaint(AObj: TMonthCalendar); extdecl;
begin
  AObj.Repaint;
end;

procedure MonthCalendar_ScaleBy(AObj: TMonthCalendar; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure MonthCalendar_SetBounds(AObj: TMonthCalendar; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure MonthCalendar_SetFocus(AObj: TMonthCalendar); extdecl;
begin
  AObj.SetFocus;
end;

procedure MonthCalendar_Update(AObj: TMonthCalendar); extdecl;
begin
  AObj.Update;
end;

procedure MonthCalendar_BringToFront(AObj: TMonthCalendar); extdecl;
begin
  AObj.BringToFront;
end;

function MonthCalendar_HasParent(AObj: TMonthCalendar): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure MonthCalendar_Hide(AObj: TMonthCalendar); extdecl;
begin
  AObj.Hide;
end;

function MonthCalendar_Perform(AObj: TMonthCalendar; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure MonthCalendar_Refresh(AObj: TMonthCalendar); extdecl;
begin
  AObj.Refresh;
end;

procedure MonthCalendar_SendToBack(AObj: TMonthCalendar); extdecl;
begin
  AObj.SendToBack;
end;

procedure MonthCalendar_Show(AObj: TMonthCalendar); extdecl;
begin
  AObj.Show;
end;

function MonthCalendar_GetTextBuf(AObj: TMonthCalendar; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function MonthCalendar_GetTextLen(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function MonthCalendar_FindComponent(AObj: TMonthCalendar; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function MonthCalendar_GetNamePath(AObj: TMonthCalendar): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure MonthCalendar_Assign(AObj: TMonthCalendar; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function MonthCalendar_ClassName(AObj: TMonthCalendar): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function MonthCalendar_Equals(AObj: TMonthCalendar; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function MonthCalendar_GetHashCode(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function MonthCalendar_ToString(AObj: TMonthCalendar): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function MonthCalendar_GetAlign(AObj: TMonthCalendar): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure MonthCalendar_SetAlign(AObj: TMonthCalendar; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function MonthCalendar_GetAnchors(AObj: TMonthCalendar): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure MonthCalendar_SetAnchors(AObj: TMonthCalendar; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function MonthCalendar_GetAutoSize(AObj: TMonthCalendar): LongBool; extdecl;
begin
  Result :=  AObj.AutoSize;
end;

procedure MonthCalendar_SetAutoSize(AObj: TMonthCalendar; AValue: LongBool); extdecl;
begin
  AObj.AutoSize := AValue;
end;

function MonthCalendar_GetBorderWidth(AObj: TMonthCalendar): TBorderWidth; extdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure MonthCalendar_SetBorderWidth(AObj: TMonthCalendar; AValue: TBorderWidth); extdecl;
begin
  AObj.BorderWidth := AValue;
end;

function MonthCalendar_GetBiDiMode(AObj: TMonthCalendar): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure MonthCalendar_SetBiDiMode(AObj: TMonthCalendar; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function MonthCalendar_GetDoubleBuffered(AObj: TMonthCalendar): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure MonthCalendar_SetDoubleBuffered(AObj: TMonthCalendar; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function MonthCalendar_GetEnabled(AObj: TMonthCalendar): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure MonthCalendar_SetEnabled(AObj: TMonthCalendar; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function MonthCalendar_GetFont(AObj: TMonthCalendar): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure MonthCalendar_SetFont(AObj: TMonthCalendar; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function MonthCalendar_GetPopupMenu(AObj: TMonthCalendar): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure MonthCalendar_SetPopupMenu(AObj: TMonthCalendar; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function MonthCalendar_GetShowHint(AObj: TMonthCalendar): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure MonthCalendar_SetShowHint(AObj: TMonthCalendar; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function MonthCalendar_GetTabOrder(AObj: TMonthCalendar): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure MonthCalendar_SetTabOrder(AObj: TMonthCalendar; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function MonthCalendar_GetTabStop(AObj: TMonthCalendar): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure MonthCalendar_SetTabStop(AObj: TMonthCalendar; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function MonthCalendar_GetVisible(AObj: TMonthCalendar): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure MonthCalendar_SetVisible(AObj: TMonthCalendar; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure MonthCalendar_SetOnClick(AObj: TMonthCalendar; AEventId: NativeUInt); extdecl;
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

procedure MonthCalendar_SetOnDblClick(AObj: TMonthCalendar; AEventId: NativeUInt); extdecl;
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

procedure MonthCalendar_SetOnEnter(AObj: TMonthCalendar; AEventId: NativeUInt); extdecl;
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

procedure MonthCalendar_SetOnExit(AObj: TMonthCalendar; AEventId: NativeUInt); extdecl;
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

procedure MonthCalendar_SetOnKeyDown(AObj: TMonthCalendar; AEventId: NativeUInt); extdecl;
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

procedure MonthCalendar_SetOnKeyPress(AObj: TMonthCalendar; AEventId: NativeUInt); extdecl;
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

procedure MonthCalendar_SetOnKeyUp(AObj: TMonthCalendar; AEventId: NativeUInt); extdecl;
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

procedure MonthCalendar_SetOnMouseEnter(AObj: TMonthCalendar; AEventId: NativeUInt); extdecl;
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

procedure MonthCalendar_SetOnMouseLeave(AObj: TMonthCalendar; AEventId: NativeUInt); extdecl;
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

function MonthCalendar_GetBrush(AObj: TMonthCalendar): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function MonthCalendar_GetControlCount(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function MonthCalendar_GetHandle(AObj: TMonthCalendar): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function MonthCalendar_GetParentWindow(AObj: TMonthCalendar): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure MonthCalendar_SetParentWindow(AObj: TMonthCalendar; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function MonthCalendar_GetAction(AObj: TMonthCalendar): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure MonthCalendar_SetAction(AObj: TMonthCalendar; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure MonthCalendar_GetBoundsRect(AObj: TMonthCalendar; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure MonthCalendar_SetBoundsRect(AObj: TMonthCalendar; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function MonthCalendar_GetClientHeight(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure MonthCalendar_SetClientHeight(AObj: TMonthCalendar; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure MonthCalendar_GetClientRect(AObj: TMonthCalendar; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function MonthCalendar_GetClientWidth(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure MonthCalendar_SetClientWidth(AObj: TMonthCalendar; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function MonthCalendar_GetParent(AObj: TMonthCalendar): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure MonthCalendar_SetParent(AObj: TMonthCalendar; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function MonthCalendar_GetLeft(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure MonthCalendar_SetLeft(AObj: TMonthCalendar; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function MonthCalendar_GetTop(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure MonthCalendar_SetTop(AObj: TMonthCalendar; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function MonthCalendar_GetWidth(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure MonthCalendar_SetWidth(AObj: TMonthCalendar; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function MonthCalendar_GetHeight(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure MonthCalendar_SetHeight(AObj: TMonthCalendar; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function MonthCalendar_GetCursor(AObj: TMonthCalendar): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure MonthCalendar_SetCursor(AObj: TMonthCalendar; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function MonthCalendar_GetHint(AObj: TMonthCalendar): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure MonthCalendar_SetHint(AObj: TMonthCalendar; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function MonthCalendar_GetComponentCount(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function MonthCalendar_GetComponentIndex(AObj: TMonthCalendar): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure MonthCalendar_SetComponentIndex(AObj: TMonthCalendar; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function MonthCalendar_GetOwner(AObj: TMonthCalendar): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function MonthCalendar_GetName(AObj: TMonthCalendar): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure MonthCalendar_SetName(AObj: TMonthCalendar; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function MonthCalendar_GetTag(AObj: TMonthCalendar): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure MonthCalendar_SetTag(AObj: TMonthCalendar; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function MonthCalendar_GetControls(AObj: TMonthCalendar; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function MonthCalendar_GetComponents(AObj: TMonthCalendar; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  MonthCalendar_Create,
  MonthCalendar_Free,
  MonthCalendar_CanFocus,
  MonthCalendar_FlipChildren,
  MonthCalendar_Focused,
  MonthCalendar_Invalidate,
  MonthCalendar_Realign,
  MonthCalendar_Repaint,
  MonthCalendar_ScaleBy,
  MonthCalendar_SetBounds,
  MonthCalendar_SetFocus,
  MonthCalendar_Update,
  MonthCalendar_BringToFront,
  MonthCalendar_HasParent,
  MonthCalendar_Hide,
  MonthCalendar_Perform,
  MonthCalendar_Refresh,
  MonthCalendar_SendToBack,
  MonthCalendar_Show,
  MonthCalendar_GetTextBuf,
  MonthCalendar_GetTextLen,
  MonthCalendar_FindComponent,
  MonthCalendar_GetNamePath,
  MonthCalendar_Assign,
  MonthCalendar_ClassName,
  MonthCalendar_Equals,
  MonthCalendar_GetHashCode,
  MonthCalendar_ToString,
  MonthCalendar_GetAlign,
  MonthCalendar_SetAlign,
  MonthCalendar_GetAnchors,
  MonthCalendar_SetAnchors,
  MonthCalendar_GetAutoSize,
  MonthCalendar_SetAutoSize,
  MonthCalendar_GetBorderWidth,
  MonthCalendar_SetBorderWidth,
  MonthCalendar_GetBiDiMode,
  MonthCalendar_SetBiDiMode,
  MonthCalendar_GetDoubleBuffered,
  MonthCalendar_SetDoubleBuffered,
  MonthCalendar_GetEnabled,
  MonthCalendar_SetEnabled,
  MonthCalendar_GetFont,
  MonthCalendar_SetFont,
  MonthCalendar_GetPopupMenu,
  MonthCalendar_SetPopupMenu,
  MonthCalendar_GetShowHint,
  MonthCalendar_SetShowHint,
  MonthCalendar_GetTabOrder,
  MonthCalendar_SetTabOrder,
  MonthCalendar_GetTabStop,
  MonthCalendar_SetTabStop,
  MonthCalendar_GetVisible,
  MonthCalendar_SetVisible,
  MonthCalendar_SetOnClick,
  MonthCalendar_SetOnDblClick,
  MonthCalendar_SetOnEnter,
  MonthCalendar_SetOnExit,
  MonthCalendar_SetOnKeyDown,
  MonthCalendar_SetOnKeyPress,
  MonthCalendar_SetOnKeyUp,
  MonthCalendar_SetOnMouseEnter,
  MonthCalendar_SetOnMouseLeave,
  MonthCalendar_GetBrush,
  MonthCalendar_GetControlCount,
  MonthCalendar_GetHandle,
  MonthCalendar_GetParentWindow,
  MonthCalendar_SetParentWindow,
  MonthCalendar_GetAction,
  MonthCalendar_SetAction,
  MonthCalendar_GetBoundsRect,
  MonthCalendar_SetBoundsRect,
  MonthCalendar_GetClientHeight,
  MonthCalendar_SetClientHeight,
  MonthCalendar_GetClientRect,
  MonthCalendar_GetClientWidth,
  MonthCalendar_SetClientWidth,
  MonthCalendar_GetParent,
  MonthCalendar_SetParent,
  MonthCalendar_GetLeft,
  MonthCalendar_SetLeft,
  MonthCalendar_GetTop,
  MonthCalendar_SetTop,
  MonthCalendar_GetWidth,
  MonthCalendar_SetWidth,
  MonthCalendar_GetHeight,
  MonthCalendar_SetHeight,
  MonthCalendar_GetCursor,
  MonthCalendar_SetCursor,
  MonthCalendar_GetHint,
  MonthCalendar_SetHint,
  MonthCalendar_GetComponentCount,
  MonthCalendar_GetComponentIndex,
  MonthCalendar_SetComponentIndex,
  MonthCalendar_GetOwner,
  MonthCalendar_GetName,
  MonthCalendar_SetName,
  MonthCalendar_GetTag,
  MonthCalendar_SetTag,
  MonthCalendar_GetControls,
  MonthCalendar_GetComponents;

