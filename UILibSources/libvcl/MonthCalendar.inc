
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function MonthCalendar_Create(AOwner: TComponent): TMonthCalendar; stdcall;
begin
  Result :=  TMonthCalendar.Create(AOwner);
end;

procedure MonthCalendar_Free(AObj: TMonthCalendar); stdcall;
begin
  AObj.Free;
end;

function MonthCalendar_CanFocus(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure MonthCalendar_FlipChildren(AObj: TMonthCalendar; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function MonthCalendar_Focused(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function MonthCalendar_HandleAllocated(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure MonthCalendar_Invalidate(AObj: TMonthCalendar); stdcall;
begin
  AObj.Invalidate;
end;

procedure MonthCalendar_Realign(AObj: TMonthCalendar); stdcall;
begin
  AObj.Realign;
end;

procedure MonthCalendar_Repaint(AObj: TMonthCalendar); stdcall;
begin
  AObj.Repaint;
end;

procedure MonthCalendar_ScaleBy(AObj: TMonthCalendar; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure MonthCalendar_SetBounds(AObj: TMonthCalendar; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure MonthCalendar_SetFocus(AObj: TMonthCalendar); stdcall;
begin
  AObj.SetFocus;
end;

procedure MonthCalendar_Update(AObj: TMonthCalendar); stdcall;
begin
  AObj.Update;
end;

procedure MonthCalendar_BringToFront(AObj: TMonthCalendar); stdcall;
begin
  AObj.BringToFront;
end;

function MonthCalendar_HasParent(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure MonthCalendar_Hide(AObj: TMonthCalendar); stdcall;
begin
  AObj.Hide;
end;

function MonthCalendar_Perform(AObj: TMonthCalendar; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure MonthCalendar_Refresh(AObj: TMonthCalendar); stdcall;
begin
  AObj.Refresh;
end;

procedure MonthCalendar_SendToBack(AObj: TMonthCalendar); stdcall;
begin
  AObj.SendToBack;
end;

procedure MonthCalendar_Show(AObj: TMonthCalendar); stdcall;
begin
  AObj.Show;
end;

function MonthCalendar_GetTextBuf(AObj: TMonthCalendar; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function MonthCalendar_GetTextLen(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function MonthCalendar_FindComponent(AObj: TMonthCalendar; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function MonthCalendar_GetNamePath(AObj: TMonthCalendar): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure MonthCalendar_Assign(AObj: TMonthCalendar; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function MonthCalendar_ClassName(AObj: TMonthCalendar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function MonthCalendar_Equals(AObj: TMonthCalendar; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function MonthCalendar_GetHashCode(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function MonthCalendar_ToString(AObj: TMonthCalendar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function MonthCalendar_GetAlign(AObj: TMonthCalendar): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure MonthCalendar_SetAlign(AObj: TMonthCalendar; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function MonthCalendar_GetAnchors(AObj: TMonthCalendar): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure MonthCalendar_SetAnchors(AObj: TMonthCalendar; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function MonthCalendar_GetAutoSize(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure MonthCalendar_SetAutoSize(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function MonthCalendar_GetBorderWidth(AObj: TMonthCalendar): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure MonthCalendar_SetBorderWidth(AObj: TMonthCalendar; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function MonthCalendar_GetBiDiMode(AObj: TMonthCalendar): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure MonthCalendar_SetBiDiMode(AObj: TMonthCalendar; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function MonthCalendar_GetCalColors(AObj: TMonthCalendar): TMonthCalColors; stdcall;
begin
  Result :=  AObj.CalColors;
end;

procedure MonthCalendar_SetCalColors(AObj: TMonthCalendar; AValue: TMonthCalColors); stdcall;
begin
  AObj.CalColors := AValue;
end;

function MonthCalendar_GetMultiSelect(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.MultiSelect;
end;

procedure MonthCalendar_SetMultiSelect(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.MultiSelect := AValue;
end;

function MonthCalendar_GetDate(AObj: TMonthCalendar): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.Date);
end;

procedure MonthCalendar_SetDate(AObj: TMonthCalendar; AValue: TUnixDateTime); stdcall;
begin
  AObj.Date := UnixToTime(AValue);
end;

function MonthCalendar_GetDoubleBuffered(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure MonthCalendar_SetDoubleBuffered(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function MonthCalendar_GetEnabled(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure MonthCalendar_SetEnabled(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function MonthCalendar_GetFirstDayOfWeek(AObj: TMonthCalendar): TCalDayOfWeek; stdcall;
begin
  Result :=  AObj.FirstDayOfWeek;
end;

procedure MonthCalendar_SetFirstDayOfWeek(AObj: TMonthCalendar; AValue: TCalDayOfWeek); stdcall;
begin
  AObj.FirstDayOfWeek := AValue;
end;

function MonthCalendar_GetFont(AObj: TMonthCalendar): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure MonthCalendar_SetFont(AObj: TMonthCalendar; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function MonthCalendar_GetMaxDate(AObj: TMonthCalendar): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.MaxDate);
end;

procedure MonthCalendar_SetMaxDate(AObj: TMonthCalendar; AValue: TUnixDateTime); stdcall;
begin
  AObj.MaxDate := UnixToTime(AValue);
end;

function MonthCalendar_GetMaxSelectRange(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.MaxSelectRange;
end;

procedure MonthCalendar_SetMaxSelectRange(AObj: TMonthCalendar; AValue: Integer); stdcall;
begin
  AObj.MaxSelectRange := AValue;
end;

function MonthCalendar_GetMinDate(AObj: TMonthCalendar): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.MinDate);
end;

procedure MonthCalendar_SetMinDate(AObj: TMonthCalendar; AValue: TUnixDateTime); stdcall;
begin
  AObj.MinDate := UnixToTime(AValue);
end;

function MonthCalendar_GetParentDoubleBuffered(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure MonthCalendar_SetParentDoubleBuffered(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function MonthCalendar_GetParentFont(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure MonthCalendar_SetParentFont(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function MonthCalendar_GetParentShowHint(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure MonthCalendar_SetParentShowHint(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function MonthCalendar_GetPopupMenu(AObj: TMonthCalendar): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure MonthCalendar_SetPopupMenu(AObj: TMonthCalendar; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function MonthCalendar_GetShowHint(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure MonthCalendar_SetShowHint(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function MonthCalendar_GetShowToday(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.ShowToday;
end;

procedure MonthCalendar_SetShowToday(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.ShowToday := AValue;
end;

function MonthCalendar_GetShowTodayCircle(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.ShowTodayCircle;
end;

procedure MonthCalendar_SetShowTodayCircle(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.ShowTodayCircle := AValue;
end;

function MonthCalendar_GetTabOrder(AObj: TMonthCalendar): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure MonthCalendar_SetTabOrder(AObj: TMonthCalendar; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function MonthCalendar_GetTabStop(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure MonthCalendar_SetTabStop(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function MonthCalendar_GetVisible(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure MonthCalendar_SetVisible(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function MonthCalendar_GetWeekNumbers(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.WeekNumbers;
end;

procedure MonthCalendar_SetWeekNumbers(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.WeekNumbers := AValue;
end;

procedure MonthCalendar_SetOnClick(AObj: TMonthCalendar; AEventId: NativeUInt); stdcall;
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

procedure MonthCalendar_SetOnDblClick(AObj: TMonthCalendar; AEventId: NativeUInt); stdcall;
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

procedure MonthCalendar_SetOnEnter(AObj: TMonthCalendar; AEventId: NativeUInt); stdcall;
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

procedure MonthCalendar_SetOnExit(AObj: TMonthCalendar; AEventId: NativeUInt); stdcall;
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

procedure MonthCalendar_SetOnKeyDown(AObj: TMonthCalendar; AEventId: NativeUInt); stdcall;
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

procedure MonthCalendar_SetOnKeyPress(AObj: TMonthCalendar; AEventId: NativeUInt); stdcall;
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

procedure MonthCalendar_SetOnKeyUp(AObj: TMonthCalendar; AEventId: NativeUInt); stdcall;
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

procedure MonthCalendar_SetOnMouseEnter(AObj: TMonthCalendar; AEventId: NativeUInt); stdcall;
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

procedure MonthCalendar_SetOnMouseLeave(AObj: TMonthCalendar; AEventId: NativeUInt); stdcall;
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

function MonthCalendar_GetBrush(AObj: TMonthCalendar): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function MonthCalendar_GetControlCount(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function MonthCalendar_GetHandle(AObj: TMonthCalendar): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function MonthCalendar_GetParentWindow(AObj: TMonthCalendar): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure MonthCalendar_SetParentWindow(AObj: TMonthCalendar; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function MonthCalendar_GetAction(AObj: TMonthCalendar): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure MonthCalendar_SetAction(AObj: TMonthCalendar; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure MonthCalendar_GetBoundsRect(AObj: TMonthCalendar; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure MonthCalendar_SetBoundsRect(AObj: TMonthCalendar; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function MonthCalendar_GetClientHeight(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure MonthCalendar_SetClientHeight(AObj: TMonthCalendar; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure MonthCalendar_GetClientRect(AObj: TMonthCalendar; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function MonthCalendar_GetClientWidth(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure MonthCalendar_SetClientWidth(AObj: TMonthCalendar; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function MonthCalendar_GetExplicitLeft(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function MonthCalendar_GetExplicitTop(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function MonthCalendar_GetExplicitWidth(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function MonthCalendar_GetExplicitHeight(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function MonthCalendar_GetParent(AObj: TMonthCalendar): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure MonthCalendar_SetParent(AObj: TMonthCalendar; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function MonthCalendar_GetStyleElements(AObj: TMonthCalendar): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure MonthCalendar_SetStyleElements(AObj: TMonthCalendar; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function MonthCalendar_GetAlignWithMargins(AObj: TMonthCalendar): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure MonthCalendar_SetAlignWithMargins(AObj: TMonthCalendar; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function MonthCalendar_GetLeft(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure MonthCalendar_SetLeft(AObj: TMonthCalendar; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function MonthCalendar_GetTop(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure MonthCalendar_SetTop(AObj: TMonthCalendar; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function MonthCalendar_GetWidth(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure MonthCalendar_SetWidth(AObj: TMonthCalendar; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function MonthCalendar_GetHeight(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure MonthCalendar_SetHeight(AObj: TMonthCalendar; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function MonthCalendar_GetCursor(AObj: TMonthCalendar): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure MonthCalendar_SetCursor(AObj: TMonthCalendar; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function MonthCalendar_GetHint(AObj: TMonthCalendar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure MonthCalendar_SetHint(AObj: TMonthCalendar; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function MonthCalendar_GetMargins(AObj: TMonthCalendar): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure MonthCalendar_SetMargins(AObj: TMonthCalendar; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function MonthCalendar_GetCustomHint(AObj: TMonthCalendar): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure MonthCalendar_SetCustomHint(AObj: TMonthCalendar; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function MonthCalendar_GetComponentCount(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function MonthCalendar_GetComponentIndex(AObj: TMonthCalendar): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure MonthCalendar_SetComponentIndex(AObj: TMonthCalendar; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function MonthCalendar_GetOwner(AObj: TMonthCalendar): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function MonthCalendar_GetName(AObj: TMonthCalendar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure MonthCalendar_SetName(AObj: TMonthCalendar; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function MonthCalendar_GetTag(AObj: TMonthCalendar): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure MonthCalendar_SetTag(AObj: TMonthCalendar; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function MonthCalendar_GetControls(AObj: TMonthCalendar; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function MonthCalendar_GetComponents(AObj: TMonthCalendar; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  MonthCalendar_Create,
  MonthCalendar_Free,
  MonthCalendar_CanFocus,
  MonthCalendar_FlipChildren,
  MonthCalendar_Focused,
  MonthCalendar_HandleAllocated,
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
  MonthCalendar_GetCalColors,
  MonthCalendar_SetCalColors,
  MonthCalendar_GetMultiSelect,
  MonthCalendar_SetMultiSelect,
  MonthCalendar_GetDate,
  MonthCalendar_SetDate,
  MonthCalendar_GetDoubleBuffered,
  MonthCalendar_SetDoubleBuffered,
  MonthCalendar_GetEnabled,
  MonthCalendar_SetEnabled,
  MonthCalendar_GetFirstDayOfWeek,
  MonthCalendar_SetFirstDayOfWeek,
  MonthCalendar_GetFont,
  MonthCalendar_SetFont,
  MonthCalendar_GetMaxDate,
  MonthCalendar_SetMaxDate,
  MonthCalendar_GetMaxSelectRange,
  MonthCalendar_SetMaxSelectRange,
  MonthCalendar_GetMinDate,
  MonthCalendar_SetMinDate,
  MonthCalendar_GetParentDoubleBuffered,
  MonthCalendar_SetParentDoubleBuffered,
  MonthCalendar_GetParentFont,
  MonthCalendar_SetParentFont,
  MonthCalendar_GetParentShowHint,
  MonthCalendar_SetParentShowHint,
  MonthCalendar_GetPopupMenu,
  MonthCalendar_SetPopupMenu,
  MonthCalendar_GetShowHint,
  MonthCalendar_SetShowHint,
  MonthCalendar_GetShowToday,
  MonthCalendar_SetShowToday,
  MonthCalendar_GetShowTodayCircle,
  MonthCalendar_SetShowTodayCircle,
  MonthCalendar_GetTabOrder,
  MonthCalendar_SetTabOrder,
  MonthCalendar_GetTabStop,
  MonthCalendar_SetTabStop,
  MonthCalendar_GetVisible,
  MonthCalendar_SetVisible,
  MonthCalendar_GetWeekNumbers,
  MonthCalendar_SetWeekNumbers,
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
  MonthCalendar_GetExplicitLeft,
  MonthCalendar_GetExplicitTop,
  MonthCalendar_GetExplicitWidth,
  MonthCalendar_GetExplicitHeight,
  MonthCalendar_GetParent,
  MonthCalendar_SetParent,
  MonthCalendar_GetStyleElements,
  MonthCalendar_SetStyleElements,
  MonthCalendar_GetAlignWithMargins,
  MonthCalendar_SetAlignWithMargins,
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
  MonthCalendar_GetMargins,
  MonthCalendar_SetMargins,
  MonthCalendar_GetCustomHint,
  MonthCalendar_SetCustomHint,
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

