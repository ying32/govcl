
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function RadioButton_Create(AOwner: TComponent): TRadioButton; stdcall;
begin
  Result :=  TRadioButton.Create(AOwner);
end;

procedure RadioButton_Free(AObj: TRadioButton); stdcall;
begin
  AObj.Free;
end;

function RadioButton_CanFocus(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure RadioButton_FlipChildren(AObj: TRadioButton; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function RadioButton_Focused(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function RadioButton_HandleAllocated(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure RadioButton_Invalidate(AObj: TRadioButton); stdcall;
begin
  AObj.Invalidate;
end;

procedure RadioButton_Realign(AObj: TRadioButton); stdcall;
begin
  AObj.Realign;
end;

procedure RadioButton_Repaint(AObj: TRadioButton); stdcall;
begin
  AObj.Repaint;
end;

procedure RadioButton_ScaleBy(AObj: TRadioButton; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure RadioButton_SetBounds(AObj: TRadioButton; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure RadioButton_SetFocus(AObj: TRadioButton); stdcall;
begin
  AObj.SetFocus;
end;

procedure RadioButton_Update(AObj: TRadioButton); stdcall;
begin
  AObj.Update;
end;

procedure RadioButton_BringToFront(AObj: TRadioButton); stdcall;
begin
  AObj.BringToFront;
end;

function RadioButton_HasParent(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure RadioButton_Hide(AObj: TRadioButton); stdcall;
begin
  AObj.Hide;
end;

function RadioButton_Perform(AObj: TRadioButton; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure RadioButton_Refresh(AObj: TRadioButton); stdcall;
begin
  AObj.Refresh;
end;

procedure RadioButton_SendToBack(AObj: TRadioButton); stdcall;
begin
  AObj.SendToBack;
end;

procedure RadioButton_Show(AObj: TRadioButton); stdcall;
begin
  AObj.Show;
end;

function RadioButton_GetTextBuf(AObj: TRadioButton; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function RadioButton_GetTextLen(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function RadioButton_FindComponent(AObj: TRadioButton; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function RadioButton_GetNamePath(AObj: TRadioButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure RadioButton_Assign(AObj: TRadioButton; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function RadioButton_ClassName(AObj: TRadioButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function RadioButton_Equals(AObj: TRadioButton; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function RadioButton_GetHashCode(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function RadioButton_ToString(AObj: TRadioButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function RadioButton_GetAction(AObj: TRadioButton): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure RadioButton_SetAction(AObj: TRadioButton; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function RadioButton_GetAlign(AObj: TRadioButton): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure RadioButton_SetAlign(AObj: TRadioButton; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function RadioButton_GetAlignment(AObj: TRadioButton): TLeftRight; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure RadioButton_SetAlignment(AObj: TRadioButton; AValue: TLeftRight); stdcall;
begin
  AObj.Alignment := AValue;
end;

function RadioButton_GetAnchors(AObj: TRadioButton): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure RadioButton_SetAnchors(AObj: TRadioButton; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function RadioButton_GetBiDiMode(AObj: TRadioButton): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure RadioButton_SetBiDiMode(AObj: TRadioButton; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function RadioButton_GetCaption(AObj: TRadioButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure RadioButton_SetCaption(AObj: TRadioButton; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function RadioButton_GetChecked(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.Checked;
end;

procedure RadioButton_SetChecked(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.Checked := AValue;
end;

function RadioButton_GetColor(AObj: TRadioButton): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure RadioButton_SetColor(AObj: TRadioButton; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function RadioButton_GetDoubleBuffered(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure RadioButton_SetDoubleBuffered(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function RadioButton_GetEnabled(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure RadioButton_SetEnabled(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function RadioButton_GetFont(AObj: TRadioButton): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure RadioButton_SetFont(AObj: TRadioButton; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function RadioButton_GetParentColor(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure RadioButton_SetParentColor(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function RadioButton_GetParentCtl3D(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure RadioButton_SetParentCtl3D(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function RadioButton_GetParentDoubleBuffered(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure RadioButton_SetParentDoubleBuffered(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function RadioButton_GetParentFont(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure RadioButton_SetParentFont(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function RadioButton_GetParentShowHint(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure RadioButton_SetParentShowHint(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function RadioButton_GetPopupMenu(AObj: TRadioButton): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure RadioButton_SetPopupMenu(AObj: TRadioButton; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function RadioButton_GetShowHint(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure RadioButton_SetShowHint(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function RadioButton_GetTabOrder(AObj: TRadioButton): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure RadioButton_SetTabOrder(AObj: TRadioButton; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function RadioButton_GetTabStop(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure RadioButton_SetTabStop(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function RadioButton_GetVisible(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure RadioButton_SetVisible(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function RadioButton_GetWordWrap(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.WordWrap;
end;

procedure RadioButton_SetWordWrap(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.WordWrap := AValue;
end;

function RadioButton_GetStyleElements(AObj: TRadioButton): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure RadioButton_SetStyleElements(AObj: TRadioButton; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure RadioButton_SetOnClick(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnDblClick(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnEnter(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnExit(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnKeyDown(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnKeyPress(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnKeyUp(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnMouseDown(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnMouseEnter(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnMouseLeave(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnMouseMove(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

procedure RadioButton_SetOnMouseUp(AObj: TRadioButton; AEventId: NativeUInt); stdcall;
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

function RadioButton_GetBrush(AObj: TRadioButton): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function RadioButton_GetControlCount(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function RadioButton_GetHandle(AObj: TRadioButton): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function RadioButton_GetParentWindow(AObj: TRadioButton): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure RadioButton_SetParentWindow(AObj: TRadioButton; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

procedure RadioButton_GetBoundsRect(AObj: TRadioButton; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure RadioButton_SetBoundsRect(AObj: TRadioButton; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function RadioButton_GetClientHeight(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure RadioButton_SetClientHeight(AObj: TRadioButton; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure RadioButton_GetClientRect(AObj: TRadioButton; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function RadioButton_GetClientWidth(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure RadioButton_SetClientWidth(AObj: TRadioButton; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function RadioButton_GetExplicitLeft(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function RadioButton_GetExplicitTop(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function RadioButton_GetExplicitWidth(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function RadioButton_GetExplicitHeight(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function RadioButton_GetParent(AObj: TRadioButton): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure RadioButton_SetParent(AObj: TRadioButton; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function RadioButton_GetAlignWithMargins(AObj: TRadioButton): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure RadioButton_SetAlignWithMargins(AObj: TRadioButton; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function RadioButton_GetLeft(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure RadioButton_SetLeft(AObj: TRadioButton; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function RadioButton_GetTop(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure RadioButton_SetTop(AObj: TRadioButton; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function RadioButton_GetWidth(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure RadioButton_SetWidth(AObj: TRadioButton; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function RadioButton_GetHeight(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure RadioButton_SetHeight(AObj: TRadioButton; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function RadioButton_GetCursor(AObj: TRadioButton): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure RadioButton_SetCursor(AObj: TRadioButton; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function RadioButton_GetHint(AObj: TRadioButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure RadioButton_SetHint(AObj: TRadioButton; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function RadioButton_GetMargins(AObj: TRadioButton): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure RadioButton_SetMargins(AObj: TRadioButton; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function RadioButton_GetCustomHint(AObj: TRadioButton): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure RadioButton_SetCustomHint(AObj: TRadioButton; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function RadioButton_GetComponentCount(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function RadioButton_GetComponentIndex(AObj: TRadioButton): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure RadioButton_SetComponentIndex(AObj: TRadioButton; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function RadioButton_GetOwner(AObj: TRadioButton): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function RadioButton_GetName(AObj: TRadioButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure RadioButton_SetName(AObj: TRadioButton; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function RadioButton_GetTag(AObj: TRadioButton): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure RadioButton_SetTag(AObj: TRadioButton; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function RadioButton_GetControls(AObj: TRadioButton; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function RadioButton_GetComponents(AObj: TRadioButton; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  RadioButton_Create,
  RadioButton_Free,
  RadioButton_CanFocus,
  RadioButton_FlipChildren,
  RadioButton_Focused,
  RadioButton_HandleAllocated,
  RadioButton_Invalidate,
  RadioButton_Realign,
  RadioButton_Repaint,
  RadioButton_ScaleBy,
  RadioButton_SetBounds,
  RadioButton_SetFocus,
  RadioButton_Update,
  RadioButton_BringToFront,
  RadioButton_HasParent,
  RadioButton_Hide,
  RadioButton_Perform,
  RadioButton_Refresh,
  RadioButton_SendToBack,
  RadioButton_Show,
  RadioButton_GetTextBuf,
  RadioButton_GetTextLen,
  RadioButton_FindComponent,
  RadioButton_GetNamePath,
  RadioButton_Assign,
  RadioButton_ClassName,
  RadioButton_Equals,
  RadioButton_GetHashCode,
  RadioButton_ToString,
  RadioButton_GetAction,
  RadioButton_SetAction,
  RadioButton_GetAlign,
  RadioButton_SetAlign,
  RadioButton_GetAlignment,
  RadioButton_SetAlignment,
  RadioButton_GetAnchors,
  RadioButton_SetAnchors,
  RadioButton_GetBiDiMode,
  RadioButton_SetBiDiMode,
  RadioButton_GetCaption,
  RadioButton_SetCaption,
  RadioButton_GetChecked,
  RadioButton_SetChecked,
  RadioButton_GetColor,
  RadioButton_SetColor,
  RadioButton_GetDoubleBuffered,
  RadioButton_SetDoubleBuffered,
  RadioButton_GetEnabled,
  RadioButton_SetEnabled,
  RadioButton_GetFont,
  RadioButton_SetFont,
  RadioButton_GetParentColor,
  RadioButton_SetParentColor,
  RadioButton_GetParentCtl3D,
  RadioButton_SetParentCtl3D,
  RadioButton_GetParentDoubleBuffered,
  RadioButton_SetParentDoubleBuffered,
  RadioButton_GetParentFont,
  RadioButton_SetParentFont,
  RadioButton_GetParentShowHint,
  RadioButton_SetParentShowHint,
  RadioButton_GetPopupMenu,
  RadioButton_SetPopupMenu,
  RadioButton_GetShowHint,
  RadioButton_SetShowHint,
  RadioButton_GetTabOrder,
  RadioButton_SetTabOrder,
  RadioButton_GetTabStop,
  RadioButton_SetTabStop,
  RadioButton_GetVisible,
  RadioButton_SetVisible,
  RadioButton_GetWordWrap,
  RadioButton_SetWordWrap,
  RadioButton_GetStyleElements,
  RadioButton_SetStyleElements,
  RadioButton_SetOnClick,
  RadioButton_SetOnDblClick,
  RadioButton_SetOnEnter,
  RadioButton_SetOnExit,
  RadioButton_SetOnKeyDown,
  RadioButton_SetOnKeyPress,
  RadioButton_SetOnKeyUp,
  RadioButton_SetOnMouseDown,
  RadioButton_SetOnMouseEnter,
  RadioButton_SetOnMouseLeave,
  RadioButton_SetOnMouseMove,
  RadioButton_SetOnMouseUp,
  RadioButton_GetBrush,
  RadioButton_GetControlCount,
  RadioButton_GetHandle,
  RadioButton_GetParentWindow,
  RadioButton_SetParentWindow,
  RadioButton_GetBoundsRect,
  RadioButton_SetBoundsRect,
  RadioButton_GetClientHeight,
  RadioButton_SetClientHeight,
  RadioButton_GetClientRect,
  RadioButton_GetClientWidth,
  RadioButton_SetClientWidth,
  RadioButton_GetExplicitLeft,
  RadioButton_GetExplicitTop,
  RadioButton_GetExplicitWidth,
  RadioButton_GetExplicitHeight,
  RadioButton_GetParent,
  RadioButton_SetParent,
  RadioButton_GetAlignWithMargins,
  RadioButton_SetAlignWithMargins,
  RadioButton_GetLeft,
  RadioButton_SetLeft,
  RadioButton_GetTop,
  RadioButton_SetTop,
  RadioButton_GetWidth,
  RadioButton_SetWidth,
  RadioButton_GetHeight,
  RadioButton_SetHeight,
  RadioButton_GetCursor,
  RadioButton_SetCursor,
  RadioButton_GetHint,
  RadioButton_SetHint,
  RadioButton_GetMargins,
  RadioButton_SetMargins,
  RadioButton_GetCustomHint,
  RadioButton_SetCustomHint,
  RadioButton_GetComponentCount,
  RadioButton_GetComponentIndex,
  RadioButton_SetComponentIndex,
  RadioButton_GetOwner,
  RadioButton_GetName,
  RadioButton_SetName,
  RadioButton_GetTag,
  RadioButton_SetTag,
  RadioButton_GetControls,
  RadioButton_GetComponents;

