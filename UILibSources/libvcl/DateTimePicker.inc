
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function DateTimePicker_Create(AOwner: TComponent): TDateTimePicker; stdcall;
begin
  Result :=  TDateTimePicker.Create(AOwner);
end;

procedure DateTimePicker_Free(AObj: TDateTimePicker); stdcall;
begin
  AObj.Free;
end;

function DateTimePicker_CanFocus(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure DateTimePicker_FlipChildren(AObj: TDateTimePicker; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function DateTimePicker_Focused(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function DateTimePicker_HandleAllocated(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure DateTimePicker_Invalidate(AObj: TDateTimePicker); stdcall;
begin
  AObj.Invalidate;
end;

procedure DateTimePicker_Realign(AObj: TDateTimePicker); stdcall;
begin
  AObj.Realign;
end;

procedure DateTimePicker_Repaint(AObj: TDateTimePicker); stdcall;
begin
  AObj.Repaint;
end;

procedure DateTimePicker_ScaleBy(AObj: TDateTimePicker; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure DateTimePicker_SetBounds(AObj: TDateTimePicker; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure DateTimePicker_SetFocus(AObj: TDateTimePicker); stdcall;
begin
  AObj.SetFocus;
end;

procedure DateTimePicker_Update(AObj: TDateTimePicker); stdcall;
begin
  AObj.Update;
end;

procedure DateTimePicker_BringToFront(AObj: TDateTimePicker); stdcall;
begin
  AObj.BringToFront;
end;

function DateTimePicker_HasParent(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure DateTimePicker_Hide(AObj: TDateTimePicker); stdcall;
begin
  AObj.Hide;
end;

function DateTimePicker_Perform(AObj: TDateTimePicker; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure DateTimePicker_Refresh(AObj: TDateTimePicker); stdcall;
begin
  AObj.Refresh;
end;

procedure DateTimePicker_SendToBack(AObj: TDateTimePicker); stdcall;
begin
  AObj.SendToBack;
end;

procedure DateTimePicker_Show(AObj: TDateTimePicker); stdcall;
begin
  AObj.Show;
end;

function DateTimePicker_GetTextBuf(AObj: TDateTimePicker; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function DateTimePicker_GetTextLen(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function DateTimePicker_FindComponent(AObj: TDateTimePicker; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function DateTimePicker_GetNamePath(AObj: TDateTimePicker): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure DateTimePicker_Assign(AObj: TDateTimePicker; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function DateTimePicker_ClassName(AObj: TDateTimePicker): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function DateTimePicker_Equals(AObj: TDateTimePicker; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function DateTimePicker_GetHashCode(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function DateTimePicker_ToString(AObj: TDateTimePicker): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function DateTimePicker_GetDateTime(AObj: TDateTimePicker): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.DateTime);
end;

procedure DateTimePicker_SetDateTime(AObj: TDateTimePicker; AValue: TUnixDateTime); stdcall;
begin
  AObj.DateTime := UnixToTime(AValue);
end;

function DateTimePicker_GetDroppedDown(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.DroppedDown;
end;

function DateTimePicker_GetAlign(AObj: TDateTimePicker): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure DateTimePicker_SetAlign(AObj: TDateTimePicker; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function DateTimePicker_GetAnchors(AObj: TDateTimePicker): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure DateTimePicker_SetAnchors(AObj: TDateTimePicker; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function DateTimePicker_GetBevelEdges(AObj: TDateTimePicker): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure DateTimePicker_SetBevelEdges(AObj: TDateTimePicker; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function DateTimePicker_GetBevelInner(AObj: TDateTimePicker): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure DateTimePicker_SetBevelInner(AObj: TDateTimePicker; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function DateTimePicker_GetBevelOuter(AObj: TDateTimePicker): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure DateTimePicker_SetBevelOuter(AObj: TDateTimePicker; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function DateTimePicker_GetBevelKind(AObj: TDateTimePicker): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure DateTimePicker_SetBevelKind(AObj: TDateTimePicker; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function DateTimePicker_GetBiDiMode(AObj: TDateTimePicker): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure DateTimePicker_SetBiDiMode(AObj: TDateTimePicker; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function DateTimePicker_GetCalAlignment(AObj: TDateTimePicker): TDTCalAlignment; stdcall;
begin
  Result :=  AObj.CalAlignment;
end;

procedure DateTimePicker_SetCalAlignment(AObj: TDateTimePicker; AValue: TDTCalAlignment); stdcall;
begin
  AObj.CalAlignment := AValue;
end;

function DateTimePicker_GetCalColors(AObj: TDateTimePicker): TMonthCalColors; stdcall;
begin
  Result :=  AObj.CalColors;
end;

procedure DateTimePicker_SetCalColors(AObj: TDateTimePicker; AValue: TMonthCalColors); stdcall;
begin
  AObj.CalColors := AValue;
end;

function DateTimePicker_GetDate(AObj: TDateTimePicker): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.Date);
end;

procedure DateTimePicker_SetDate(AObj: TDateTimePicker; AValue: TUnixDateTime); stdcall;
begin
  AObj.Date := UnixToTime(AValue);
end;

function DateTimePicker_GetFormat(AObj: TDateTimePicker): PChar; stdcall;
begin
  Result :=  PChar(AObj.Format);
end;

procedure DateTimePicker_SetFormat(AObj: TDateTimePicker; AValue: PChar); stdcall;
begin
  AObj.Format := AValue;
end;

function DateTimePicker_GetTime(AObj: TDateTimePicker): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.Time);
end;

procedure DateTimePicker_SetTime(AObj: TDateTimePicker; AValue: TUnixDateTime); stdcall;
begin
  AObj.Time := UnixToTime(AValue);
end;

function DateTimePicker_GetChecked(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.Checked;
end;

procedure DateTimePicker_SetChecked(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.Checked := AValue;
end;

function DateTimePicker_GetColor(AObj: TDateTimePicker): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure DateTimePicker_SetColor(AObj: TDateTimePicker; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function DateTimePicker_GetDateFormat(AObj: TDateTimePicker): TDTDateFormat; stdcall;
begin
  Result :=  AObj.DateFormat;
end;

procedure DateTimePicker_SetDateFormat(AObj: TDateTimePicker; AValue: TDTDateFormat); stdcall;
begin
  AObj.DateFormat := AValue;
end;

function DateTimePicker_GetDateMode(AObj: TDateTimePicker): TDTDateMode; stdcall;
begin
  Result :=  AObj.DateMode;
end;

procedure DateTimePicker_SetDateMode(AObj: TDateTimePicker; AValue: TDTDateMode); stdcall;
begin
  AObj.DateMode := AValue;
end;

function DateTimePicker_GetDoubleBuffered(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure DateTimePicker_SetDoubleBuffered(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function DateTimePicker_GetEnabled(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure DateTimePicker_SetEnabled(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function DateTimePicker_GetFont(AObj: TDateTimePicker): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure DateTimePicker_SetFont(AObj: TDateTimePicker; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function DateTimePicker_GetKind(AObj: TDateTimePicker): TDateTimeKind; stdcall;
begin
  Result :=  AObj.Kind;
end;

procedure DateTimePicker_SetKind(AObj: TDateTimePicker; AValue: TDateTimeKind); stdcall;
begin
  AObj.Kind := AValue;
end;

function DateTimePicker_GetMaxDate(AObj: TDateTimePicker): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.MaxDate);
end;

procedure DateTimePicker_SetMaxDate(AObj: TDateTimePicker; AValue: TUnixDateTime); stdcall;
begin
  AObj.MaxDate := UnixToTime(AValue);
end;

function DateTimePicker_GetMinDate(AObj: TDateTimePicker): TUnixDateTime; stdcall;
begin
  Result :=  ToUnixTime(AObj.MinDate);
end;

procedure DateTimePicker_SetMinDate(AObj: TDateTimePicker; AValue: TUnixDateTime); stdcall;
begin
  AObj.MinDate := UnixToTime(AValue);
end;

function DateTimePicker_GetParseInput(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.ParseInput;
end;

procedure DateTimePicker_SetParseInput(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.ParseInput := AValue;
end;

function DateTimePicker_GetParentColor(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure DateTimePicker_SetParentColor(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function DateTimePicker_GetParentDoubleBuffered(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure DateTimePicker_SetParentDoubleBuffered(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function DateTimePicker_GetParentFont(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure DateTimePicker_SetParentFont(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function DateTimePicker_GetParentShowHint(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure DateTimePicker_SetParentShowHint(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function DateTimePicker_GetPopupMenu(AObj: TDateTimePicker): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure DateTimePicker_SetPopupMenu(AObj: TDateTimePicker; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function DateTimePicker_GetShowHint(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure DateTimePicker_SetShowHint(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function DateTimePicker_GetTabOrder(AObj: TDateTimePicker): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure DateTimePicker_SetTabOrder(AObj: TDateTimePicker; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function DateTimePicker_GetTabStop(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure DateTimePicker_SetTabStop(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function DateTimePicker_GetVisible(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure DateTimePicker_SetVisible(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function DateTimePicker_GetStyleElements(AObj: TDateTimePicker): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure DateTimePicker_SetStyleElements(AObj: TDateTimePicker; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure DateTimePicker_SetOnClick(AObj: TDateTimePicker; AEventId: NativeUInt); stdcall;
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

procedure DateTimePicker_SetOnChange(AObj: TDateTimePicker; AEventId: NativeUInt); stdcall;
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

procedure DateTimePicker_SetOnEnter(AObj: TDateTimePicker; AEventId: NativeUInt); stdcall;
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

procedure DateTimePicker_SetOnExit(AObj: TDateTimePicker; AEventId: NativeUInt); stdcall;
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

procedure DateTimePicker_SetOnKeyDown(AObj: TDateTimePicker; AEventId: NativeUInt); stdcall;
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

procedure DateTimePicker_SetOnKeyPress(AObj: TDateTimePicker; AEventId: NativeUInt); stdcall;
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

procedure DateTimePicker_SetOnKeyUp(AObj: TDateTimePicker; AEventId: NativeUInt); stdcall;
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

procedure DateTimePicker_SetOnMouseEnter(AObj: TDateTimePicker; AEventId: NativeUInt); stdcall;
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

procedure DateTimePicker_SetOnMouseLeave(AObj: TDateTimePicker; AEventId: NativeUInt); stdcall;
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

function DateTimePicker_GetBrush(AObj: TDateTimePicker): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function DateTimePicker_GetControlCount(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function DateTimePicker_GetHandle(AObj: TDateTimePicker): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function DateTimePicker_GetParentWindow(AObj: TDateTimePicker): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure DateTimePicker_SetParentWindow(AObj: TDateTimePicker; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function DateTimePicker_GetAction(AObj: TDateTimePicker): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure DateTimePicker_SetAction(AObj: TDateTimePicker; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure DateTimePicker_GetBoundsRect(AObj: TDateTimePicker; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure DateTimePicker_SetBoundsRect(AObj: TDateTimePicker; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function DateTimePicker_GetClientHeight(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure DateTimePicker_SetClientHeight(AObj: TDateTimePicker; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure DateTimePicker_GetClientRect(AObj: TDateTimePicker; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function DateTimePicker_GetClientWidth(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure DateTimePicker_SetClientWidth(AObj: TDateTimePicker; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function DateTimePicker_GetExplicitLeft(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function DateTimePicker_GetExplicitTop(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function DateTimePicker_GetExplicitWidth(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function DateTimePicker_GetExplicitHeight(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function DateTimePicker_GetParent(AObj: TDateTimePicker): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure DateTimePicker_SetParent(AObj: TDateTimePicker; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function DateTimePicker_GetAlignWithMargins(AObj: TDateTimePicker): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure DateTimePicker_SetAlignWithMargins(AObj: TDateTimePicker; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function DateTimePicker_GetLeft(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure DateTimePicker_SetLeft(AObj: TDateTimePicker; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function DateTimePicker_GetTop(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure DateTimePicker_SetTop(AObj: TDateTimePicker; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function DateTimePicker_GetWidth(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure DateTimePicker_SetWidth(AObj: TDateTimePicker; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function DateTimePicker_GetHeight(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure DateTimePicker_SetHeight(AObj: TDateTimePicker; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function DateTimePicker_GetCursor(AObj: TDateTimePicker): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure DateTimePicker_SetCursor(AObj: TDateTimePicker; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function DateTimePicker_GetHint(AObj: TDateTimePicker): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure DateTimePicker_SetHint(AObj: TDateTimePicker; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function DateTimePicker_GetMargins(AObj: TDateTimePicker): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure DateTimePicker_SetMargins(AObj: TDateTimePicker; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function DateTimePicker_GetCustomHint(AObj: TDateTimePicker): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure DateTimePicker_SetCustomHint(AObj: TDateTimePicker; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function DateTimePicker_GetComponentCount(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function DateTimePicker_GetComponentIndex(AObj: TDateTimePicker): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure DateTimePicker_SetComponentIndex(AObj: TDateTimePicker; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function DateTimePicker_GetOwner(AObj: TDateTimePicker): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function DateTimePicker_GetName(AObj: TDateTimePicker): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure DateTimePicker_SetName(AObj: TDateTimePicker; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function DateTimePicker_GetTag(AObj: TDateTimePicker): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure DateTimePicker_SetTag(AObj: TDateTimePicker; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function DateTimePicker_GetControls(AObj: TDateTimePicker; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function DateTimePicker_GetComponents(AObj: TDateTimePicker; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  DateTimePicker_Create,
  DateTimePicker_Free,
  DateTimePicker_CanFocus,
  DateTimePicker_FlipChildren,
  DateTimePicker_Focused,
  DateTimePicker_HandleAllocated,
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
  DateTimePicker_GetBevelEdges,
  DateTimePicker_SetBevelEdges,
  DateTimePicker_GetBevelInner,
  DateTimePicker_SetBevelInner,
  DateTimePicker_GetBevelOuter,
  DateTimePicker_SetBevelOuter,
  DateTimePicker_GetBevelKind,
  DateTimePicker_SetBevelKind,
  DateTimePicker_GetBiDiMode,
  DateTimePicker_SetBiDiMode,
  DateTimePicker_GetCalAlignment,
  DateTimePicker_SetCalAlignment,
  DateTimePicker_GetCalColors,
  DateTimePicker_SetCalColors,
  DateTimePicker_GetDate,
  DateTimePicker_SetDate,
  DateTimePicker_GetFormat,
  DateTimePicker_SetFormat,
  DateTimePicker_GetTime,
  DateTimePicker_SetTime,
  DateTimePicker_GetChecked,
  DateTimePicker_SetChecked,
  DateTimePicker_GetColor,
  DateTimePicker_SetColor,
  DateTimePicker_GetDateFormat,
  DateTimePicker_SetDateFormat,
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
  DateTimePicker_GetParseInput,
  DateTimePicker_SetParseInput,
  DateTimePicker_GetParentColor,
  DateTimePicker_SetParentColor,
  DateTimePicker_GetParentDoubleBuffered,
  DateTimePicker_SetParentDoubleBuffered,
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
  DateTimePicker_GetStyleElements,
  DateTimePicker_SetStyleElements,
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
  DateTimePicker_GetExplicitLeft,
  DateTimePicker_GetExplicitTop,
  DateTimePicker_GetExplicitWidth,
  DateTimePicker_GetExplicitHeight,
  DateTimePicker_GetParent,
  DateTimePicker_SetParent,
  DateTimePicker_GetAlignWithMargins,
  DateTimePicker_SetAlignWithMargins,
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
  DateTimePicker_GetMargins,
  DateTimePicker_SetMargins,
  DateTimePicker_GetCustomHint,
  DateTimePicker_SetCustomHint,
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

