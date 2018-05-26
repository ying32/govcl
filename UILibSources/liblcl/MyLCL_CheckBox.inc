
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function CheckBox_Create(AOwner: TComponent): TCheckBox; extdecl;
begin
  Result :=  TCheckBox.Create(AOwner);
end;

procedure CheckBox_Free(AObj: TCheckBox); extdecl;
begin
  AObj.Free;
end;

function CheckBox_CanFocus(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure CheckBox_FlipChildren(AObj: TCheckBox; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function CheckBox_Focused(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure CheckBox_Invalidate(AObj: TCheckBox); extdecl;
begin
  AObj.Invalidate;
end;

procedure CheckBox_Realign(AObj: TCheckBox); extdecl;
begin
  AObj.Realign;
end;

procedure CheckBox_Repaint(AObj: TCheckBox); extdecl;
begin
  AObj.Repaint;
end;

procedure CheckBox_ScaleBy(AObj: TCheckBox; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure CheckBox_SetBounds(AObj: TCheckBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure CheckBox_SetFocus(AObj: TCheckBox); extdecl;
begin
  AObj.SetFocus;
end;

procedure CheckBox_Update(AObj: TCheckBox); extdecl;
begin
  AObj.Update;
end;

procedure CheckBox_BringToFront(AObj: TCheckBox); extdecl;
begin
  AObj.BringToFront;
end;

function CheckBox_HasParent(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure CheckBox_Hide(AObj: TCheckBox); extdecl;
begin
  AObj.Hide;
end;

function CheckBox_Perform(AObj: TCheckBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure CheckBox_Refresh(AObj: TCheckBox); extdecl;
begin
  AObj.Refresh;
end;

procedure CheckBox_SendToBack(AObj: TCheckBox); extdecl;
begin
  AObj.SendToBack;
end;

procedure CheckBox_Show(AObj: TCheckBox); extdecl;
begin
  AObj.Show;
end;

function CheckBox_GetTextBuf(AObj: TCheckBox; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function CheckBox_GetTextLen(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function CheckBox_FindComponent(AObj: TCheckBox; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function CheckBox_GetNamePath(AObj: TCheckBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure CheckBox_Assign(AObj: TCheckBox; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function CheckBox_ClassName(AObj: TCheckBox): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function CheckBox_Equals(AObj: TCheckBox; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function CheckBox_GetHashCode(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function CheckBox_ToString(AObj: TCheckBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function CheckBox_GetAction(AObj: TCheckBox): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure CheckBox_SetAction(AObj: TCheckBox; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function CheckBox_GetAlign(AObj: TCheckBox): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure CheckBox_SetAlign(AObj: TCheckBox; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function CheckBox_GetAlignment(AObj: TCheckBox): TLeftRight; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure CheckBox_SetAlignment(AObj: TCheckBox; AValue: TLeftRight); extdecl;
begin
  AObj.Alignment := AValue;
end;

function CheckBox_GetAllowGrayed(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.AllowGrayed;
end;

procedure CheckBox_SetAllowGrayed(AObj: TCheckBox; AValue: LongBool); extdecl;
begin
  AObj.AllowGrayed := AValue;
end;

function CheckBox_GetAnchors(AObj: TCheckBox): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure CheckBox_SetAnchors(AObj: TCheckBox; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function CheckBox_GetBiDiMode(AObj: TCheckBox): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure CheckBox_SetBiDiMode(AObj: TCheckBox; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function CheckBox_GetCaption(AObj: TCheckBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure CheckBox_SetCaption(AObj: TCheckBox; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function CheckBox_GetChecked(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.Checked;
end;

procedure CheckBox_SetChecked(AObj: TCheckBox; AValue: LongBool); extdecl;
begin
  AObj.Checked := AValue;
end;

function CheckBox_GetColor(AObj: TCheckBox): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure CheckBox_SetColor(AObj: TCheckBox; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function CheckBox_GetDoubleBuffered(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure CheckBox_SetDoubleBuffered(AObj: TCheckBox; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function CheckBox_GetEnabled(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure CheckBox_SetEnabled(AObj: TCheckBox; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function CheckBox_GetFont(AObj: TCheckBox): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure CheckBox_SetFont(AObj: TCheckBox; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function CheckBox_GetParentColor(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure CheckBox_SetParentColor(AObj: TCheckBox; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function CheckBox_GetParentFont(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure CheckBox_SetParentFont(AObj: TCheckBox; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function CheckBox_GetParentShowHint(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure CheckBox_SetParentShowHint(AObj: TCheckBox; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function CheckBox_GetPopupMenu(AObj: TCheckBox): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure CheckBox_SetPopupMenu(AObj: TCheckBox; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function CheckBox_GetShowHint(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure CheckBox_SetShowHint(AObj: TCheckBox; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function CheckBox_GetState(AObj: TCheckBox): TCheckBoxState; extdecl;
begin
  Result :=  AObj.State;
end;

procedure CheckBox_SetState(AObj: TCheckBox; AValue: TCheckBoxState); extdecl;
begin
  AObj.State := AValue;
end;

function CheckBox_GetTabOrder(AObj: TCheckBox): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure CheckBox_SetTabOrder(AObj: TCheckBox; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function CheckBox_GetTabStop(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure CheckBox_SetTabStop(AObj: TCheckBox; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function CheckBox_GetVisible(AObj: TCheckBox): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure CheckBox_SetVisible(AObj: TCheckBox; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure CheckBox_SetOnClick(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

procedure CheckBox_SetOnEnter(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

procedure CheckBox_SetOnExit(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

procedure CheckBox_SetOnKeyDown(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

procedure CheckBox_SetOnKeyPress(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

procedure CheckBox_SetOnKeyUp(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

procedure CheckBox_SetOnMouseDown(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

procedure CheckBox_SetOnMouseEnter(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

procedure CheckBox_SetOnMouseLeave(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

procedure CheckBox_SetOnMouseMove(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

procedure CheckBox_SetOnMouseUp(AObj: TCheckBox; AEventId: NativeUInt); extdecl;
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

function CheckBox_GetBrush(AObj: TCheckBox): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function CheckBox_GetControlCount(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function CheckBox_GetHandle(AObj: TCheckBox): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function CheckBox_GetParentWindow(AObj: TCheckBox): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure CheckBox_SetParentWindow(AObj: TCheckBox; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

procedure CheckBox_GetBoundsRect(AObj: TCheckBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure CheckBox_SetBoundsRect(AObj: TCheckBox; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function CheckBox_GetClientHeight(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure CheckBox_SetClientHeight(AObj: TCheckBox; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure CheckBox_GetClientRect(AObj: TCheckBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function CheckBox_GetClientWidth(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure CheckBox_SetClientWidth(AObj: TCheckBox; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function CheckBox_GetParent(AObj: TCheckBox): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure CheckBox_SetParent(AObj: TCheckBox; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function CheckBox_GetLeft(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure CheckBox_SetLeft(AObj: TCheckBox; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function CheckBox_GetTop(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure CheckBox_SetTop(AObj: TCheckBox; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function CheckBox_GetWidth(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure CheckBox_SetWidth(AObj: TCheckBox; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function CheckBox_GetHeight(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure CheckBox_SetHeight(AObj: TCheckBox; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function CheckBox_GetCursor(AObj: TCheckBox): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure CheckBox_SetCursor(AObj: TCheckBox; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function CheckBox_GetHint(AObj: TCheckBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure CheckBox_SetHint(AObj: TCheckBox; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function CheckBox_GetComponentCount(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function CheckBox_GetComponentIndex(AObj: TCheckBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure CheckBox_SetComponentIndex(AObj: TCheckBox; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function CheckBox_GetOwner(AObj: TCheckBox): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function CheckBox_GetName(AObj: TCheckBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure CheckBox_SetName(AObj: TCheckBox; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function CheckBox_GetTag(AObj: TCheckBox): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure CheckBox_SetTag(AObj: TCheckBox; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function CheckBox_GetControls(AObj: TCheckBox; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function CheckBox_GetComponents(AObj: TCheckBox; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  CheckBox_Create,
  CheckBox_Free,
  CheckBox_CanFocus,
  CheckBox_FlipChildren,
  CheckBox_Focused,
  CheckBox_Invalidate,
  CheckBox_Realign,
  CheckBox_Repaint,
  CheckBox_ScaleBy,
  CheckBox_SetBounds,
  CheckBox_SetFocus,
  CheckBox_Update,
  CheckBox_BringToFront,
  CheckBox_HasParent,
  CheckBox_Hide,
  CheckBox_Perform,
  CheckBox_Refresh,
  CheckBox_SendToBack,
  CheckBox_Show,
  CheckBox_GetTextBuf,
  CheckBox_GetTextLen,
  CheckBox_FindComponent,
  CheckBox_GetNamePath,
  CheckBox_Assign,
  CheckBox_ClassName,
  CheckBox_Equals,
  CheckBox_GetHashCode,
  CheckBox_ToString,
  CheckBox_GetAction,
  CheckBox_SetAction,
  CheckBox_GetAlign,
  CheckBox_SetAlign,
  CheckBox_GetAlignment,
  CheckBox_SetAlignment,
  CheckBox_GetAllowGrayed,
  CheckBox_SetAllowGrayed,
  CheckBox_GetAnchors,
  CheckBox_SetAnchors,
  CheckBox_GetBiDiMode,
  CheckBox_SetBiDiMode,
  CheckBox_GetCaption,
  CheckBox_SetCaption,
  CheckBox_GetChecked,
  CheckBox_SetChecked,
  CheckBox_GetColor,
  CheckBox_SetColor,
  CheckBox_GetDoubleBuffered,
  CheckBox_SetDoubleBuffered,
  CheckBox_GetEnabled,
  CheckBox_SetEnabled,
  CheckBox_GetFont,
  CheckBox_SetFont,
  CheckBox_GetParentColor,
  CheckBox_SetParentColor,
  CheckBox_GetParentFont,
  CheckBox_SetParentFont,
  CheckBox_GetParentShowHint,
  CheckBox_SetParentShowHint,
  CheckBox_GetPopupMenu,
  CheckBox_SetPopupMenu,
  CheckBox_GetShowHint,
  CheckBox_SetShowHint,
  CheckBox_GetState,
  CheckBox_SetState,
  CheckBox_GetTabOrder,
  CheckBox_SetTabOrder,
  CheckBox_GetTabStop,
  CheckBox_SetTabStop,
  CheckBox_GetVisible,
  CheckBox_SetVisible,
  CheckBox_SetOnClick,
  CheckBox_SetOnEnter,
  CheckBox_SetOnExit,
  CheckBox_SetOnKeyDown,
  CheckBox_SetOnKeyPress,
  CheckBox_SetOnKeyUp,
  CheckBox_SetOnMouseDown,
  CheckBox_SetOnMouseEnter,
  CheckBox_SetOnMouseLeave,
  CheckBox_SetOnMouseMove,
  CheckBox_SetOnMouseUp,
  CheckBox_GetBrush,
  CheckBox_GetControlCount,
  CheckBox_GetHandle,
  CheckBox_GetParentWindow,
  CheckBox_SetParentWindow,
  CheckBox_GetBoundsRect,
  CheckBox_SetBoundsRect,
  CheckBox_GetClientHeight,
  CheckBox_SetClientHeight,
  CheckBox_GetClientRect,
  CheckBox_GetClientWidth,
  CheckBox_SetClientWidth,
  CheckBox_GetParent,
  CheckBox_SetParent,
  CheckBox_GetLeft,
  CheckBox_SetLeft,
  CheckBox_GetTop,
  CheckBox_SetTop,
  CheckBox_GetWidth,
  CheckBox_SetWidth,
  CheckBox_GetHeight,
  CheckBox_SetHeight,
  CheckBox_GetCursor,
  CheckBox_SetCursor,
  CheckBox_GetHint,
  CheckBox_SetHint,
  CheckBox_GetComponentCount,
  CheckBox_GetComponentIndex,
  CheckBox_SetComponentIndex,
  CheckBox_GetOwner,
  CheckBox_GetName,
  CheckBox_SetName,
  CheckBox_GetTag,
  CheckBox_SetTag,
  CheckBox_GetControls,
  CheckBox_GetComponents;

