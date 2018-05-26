
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function RadioButton_Create(AOwner: TComponent): TRadioButton; extdecl;
begin
  Result :=  TRadioButton.Create(AOwner);
end;

procedure RadioButton_Free(AObj: TRadioButton); extdecl;
begin
  AObj.Free;
end;

function RadioButton_CanFocus(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure RadioButton_FlipChildren(AObj: TRadioButton; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function RadioButton_Focused(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure RadioButton_Invalidate(AObj: TRadioButton); extdecl;
begin
  AObj.Invalidate;
end;

procedure RadioButton_Realign(AObj: TRadioButton); extdecl;
begin
  AObj.Realign;
end;

procedure RadioButton_Repaint(AObj: TRadioButton); extdecl;
begin
  AObj.Repaint;
end;

procedure RadioButton_ScaleBy(AObj: TRadioButton; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure RadioButton_SetBounds(AObj: TRadioButton; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure RadioButton_SetFocus(AObj: TRadioButton); extdecl;
begin
  AObj.SetFocus;
end;

procedure RadioButton_Update(AObj: TRadioButton); extdecl;
begin
  AObj.Update;
end;

procedure RadioButton_BringToFront(AObj: TRadioButton); extdecl;
begin
  AObj.BringToFront;
end;

function RadioButton_HasParent(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure RadioButton_Hide(AObj: TRadioButton); extdecl;
begin
  AObj.Hide;
end;

function RadioButton_Perform(AObj: TRadioButton; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure RadioButton_Refresh(AObj: TRadioButton); extdecl;
begin
  AObj.Refresh;
end;

procedure RadioButton_SendToBack(AObj: TRadioButton); extdecl;
begin
  AObj.SendToBack;
end;

procedure RadioButton_Show(AObj: TRadioButton); extdecl;
begin
  AObj.Show;
end;

function RadioButton_GetTextBuf(AObj: TRadioButton; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function RadioButton_GetTextLen(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function RadioButton_FindComponent(AObj: TRadioButton; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function RadioButton_GetNamePath(AObj: TRadioButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure RadioButton_Assign(AObj: TRadioButton; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function RadioButton_ClassName(AObj: TRadioButton): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function RadioButton_Equals(AObj: TRadioButton; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function RadioButton_GetHashCode(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function RadioButton_ToString(AObj: TRadioButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function RadioButton_GetAction(AObj: TRadioButton): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure RadioButton_SetAction(AObj: TRadioButton; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function RadioButton_GetAlign(AObj: TRadioButton): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure RadioButton_SetAlign(AObj: TRadioButton; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function RadioButton_GetAlignment(AObj: TRadioButton): TLeftRight; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure RadioButton_SetAlignment(AObj: TRadioButton; AValue: TLeftRight); extdecl;
begin
  AObj.Alignment := AValue;
end;

function RadioButton_GetAnchors(AObj: TRadioButton): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure RadioButton_SetAnchors(AObj: TRadioButton; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function RadioButton_GetBiDiMode(AObj: TRadioButton): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure RadioButton_SetBiDiMode(AObj: TRadioButton; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function RadioButton_GetCaption(AObj: TRadioButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure RadioButton_SetCaption(AObj: TRadioButton; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function RadioButton_GetChecked(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.Checked;
end;

procedure RadioButton_SetChecked(AObj: TRadioButton; AValue: LongBool); extdecl;
begin
  AObj.Checked := AValue;
end;

function RadioButton_GetColor(AObj: TRadioButton): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure RadioButton_SetColor(AObj: TRadioButton; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function RadioButton_GetDoubleBuffered(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure RadioButton_SetDoubleBuffered(AObj: TRadioButton; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function RadioButton_GetEnabled(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure RadioButton_SetEnabled(AObj: TRadioButton; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function RadioButton_GetFont(AObj: TRadioButton): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure RadioButton_SetFont(AObj: TRadioButton; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function RadioButton_GetParentColor(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure RadioButton_SetParentColor(AObj: TRadioButton; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function RadioButton_GetParentFont(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure RadioButton_SetParentFont(AObj: TRadioButton; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function RadioButton_GetParentShowHint(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure RadioButton_SetParentShowHint(AObj: TRadioButton; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function RadioButton_GetPopupMenu(AObj: TRadioButton): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure RadioButton_SetPopupMenu(AObj: TRadioButton; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function RadioButton_GetShowHint(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure RadioButton_SetShowHint(AObj: TRadioButton; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function RadioButton_GetTabOrder(AObj: TRadioButton): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure RadioButton_SetTabOrder(AObj: TRadioButton; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function RadioButton_GetTabStop(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure RadioButton_SetTabStop(AObj: TRadioButton; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function RadioButton_GetVisible(AObj: TRadioButton): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure RadioButton_SetVisible(AObj: TRadioButton; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure RadioButton_SetOnClick(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
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

procedure RadioButton_SetOnEnter(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
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

procedure RadioButton_SetOnExit(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
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

procedure RadioButton_SetOnKeyDown(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
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

procedure RadioButton_SetOnKeyPress(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
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

procedure RadioButton_SetOnKeyUp(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
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

procedure RadioButton_SetOnMouseDown(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseDown := nil;
    EventClass.Remove(AObj, geMouseDown);
    Exit;
  end;
  AObj.OnMouseDown := @EventClass.OnMouseDown;
  EventClass.Add(AObj, geMouseDown, AEventId);
end;

procedure RadioButton_SetOnMouseEnter(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
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

procedure RadioButton_SetOnMouseLeave(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
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

procedure RadioButton_SetOnMouseMove(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseMove := nil;
    EventClass.Remove(AObj, geMouseMove);
    Exit;
  end;
  AObj.OnMouseMove := @EventClass.OnMouseMove;
  EventClass.Add(AObj, geMouseMove, AEventId);
end;

procedure RadioButton_SetOnMouseUp(AObj: TRadioButton; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseUp := nil;
    EventClass.Remove(AObj, geMouseUp);
    Exit;
  end;
  AObj.OnMouseUp := @EventClass.OnMouseUp;
  EventClass.Add(AObj, geMouseUp, AEventId);
end;

function RadioButton_GetBrush(AObj: TRadioButton): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function RadioButton_GetControlCount(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function RadioButton_GetHandle(AObj: TRadioButton): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function RadioButton_GetParentWindow(AObj: TRadioButton): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure RadioButton_SetParentWindow(AObj: TRadioButton; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

procedure RadioButton_GetBoundsRect(AObj: TRadioButton; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure RadioButton_SetBoundsRect(AObj: TRadioButton; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function RadioButton_GetClientHeight(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure RadioButton_SetClientHeight(AObj: TRadioButton; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure RadioButton_GetClientRect(AObj: TRadioButton; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function RadioButton_GetClientWidth(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure RadioButton_SetClientWidth(AObj: TRadioButton; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function RadioButton_GetParent(AObj: TRadioButton): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure RadioButton_SetParent(AObj: TRadioButton; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function RadioButton_GetLeft(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure RadioButton_SetLeft(AObj: TRadioButton; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function RadioButton_GetTop(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure RadioButton_SetTop(AObj: TRadioButton; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function RadioButton_GetWidth(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure RadioButton_SetWidth(AObj: TRadioButton; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function RadioButton_GetHeight(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure RadioButton_SetHeight(AObj: TRadioButton; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function RadioButton_GetCursor(AObj: TRadioButton): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure RadioButton_SetCursor(AObj: TRadioButton; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function RadioButton_GetHint(AObj: TRadioButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure RadioButton_SetHint(AObj: TRadioButton; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function RadioButton_GetComponentCount(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function RadioButton_GetComponentIndex(AObj: TRadioButton): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure RadioButton_SetComponentIndex(AObj: TRadioButton; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function RadioButton_GetOwner(AObj: TRadioButton): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function RadioButton_GetName(AObj: TRadioButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure RadioButton_SetName(AObj: TRadioButton; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function RadioButton_GetTag(AObj: TRadioButton): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure RadioButton_SetTag(AObj: TRadioButton; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function RadioButton_GetControls(AObj: TRadioButton; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function RadioButton_GetComponents(AObj: TRadioButton; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  RadioButton_Create,
  RadioButton_Free,
  RadioButton_CanFocus,
  RadioButton_FlipChildren,
  RadioButton_Focused,
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
  RadioButton_SetOnClick,
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
  RadioButton_GetParent,
  RadioButton_SetParent,
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

