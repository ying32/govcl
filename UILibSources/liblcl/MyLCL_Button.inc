
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Button_Create(AOwner: TComponent): TButton; extdecl;
begin
  Result :=  TButton.Create(AOwner);
end;

procedure Button_Free(AObj: TButton); extdecl;
begin
  AObj.Free;
end;

procedure Button_Click(AObj: TButton); extdecl;
begin
  AObj.Click;
end;

function Button_CanFocus(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure Button_FlipChildren(AObj: TButton; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function Button_Focused(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure Button_Invalidate(AObj: TButton); extdecl;
begin
  AObj.Invalidate;
end;

procedure Button_Realign(AObj: TButton); extdecl;
begin
  AObj.Realign;
end;

procedure Button_Repaint(AObj: TButton); extdecl;
begin
  AObj.Repaint;
end;

procedure Button_ScaleBy(AObj: TButton; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure Button_SetBounds(AObj: TButton; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Button_SetFocus(AObj: TButton); extdecl;
begin
  AObj.SetFocus;
end;

procedure Button_Update(AObj: TButton); extdecl;
begin
  AObj.Update;
end;

procedure Button_BringToFront(AObj: TButton); extdecl;
begin
  AObj.BringToFront;
end;

function Button_HasParent(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Button_Hide(AObj: TButton); extdecl;
begin
  AObj.Hide;
end;

function Button_Perform(AObj: TButton; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Button_Refresh(AObj: TButton); extdecl;
begin
  AObj.Refresh;
end;

procedure Button_SendToBack(AObj: TButton); extdecl;
begin
  AObj.SendToBack;
end;

procedure Button_Show(AObj: TButton); extdecl;
begin
  AObj.Show;
end;

function Button_GetTextBuf(AObj: TButton; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Button_GetTextLen(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function Button_FindComponent(AObj: TButton; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Button_GetNamePath(AObj: TButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Button_Assign(AObj: TButton; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Button_ClassName(AObj: TButton): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Button_Equals(AObj: TButton; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Button_GetHashCode(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Button_ToString(AObj: TButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Button_GetAction(AObj: TButton): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure Button_SetAction(AObj: TButton; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function Button_GetAlign(AObj: TButton): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure Button_SetAlign(AObj: TButton; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function Button_GetAnchors(AObj: TButton): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure Button_SetAnchors(AObj: TButton; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function Button_GetBiDiMode(AObj: TButton): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Button_SetBiDiMode(AObj: TButton; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function Button_GetCancel(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.Cancel;
end;

procedure Button_SetCancel(AObj: TButton; AValue: LongBool); extdecl;
begin
  AObj.Cancel := AValue;
end;

function Button_GetCaption(AObj: TButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure Button_SetCaption(AObj: TButton; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function Button_GetDefault(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.Default;
end;

procedure Button_SetDefault(AObj: TButton; AValue: LongBool); extdecl;
begin
  AObj.Default := AValue;
end;

function Button_GetDoubleBuffered(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure Button_SetDoubleBuffered(AObj: TButton; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function Button_GetEnabled(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Button_SetEnabled(AObj: TButton; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Button_GetFont(AObj: TButton): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure Button_SetFont(AObj: TButton; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function Button_GetModalResult(AObj: TButton): TModalResult; extdecl;
begin
  Result :=  AObj.ModalResult;
end;

procedure Button_SetModalResult(AObj: TButton; AValue: TModalResult); extdecl;
begin
  AObj.ModalResult := AValue;
end;

function Button_GetParentFont(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure Button_SetParentFont(AObj: TButton; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function Button_GetParentShowHint(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Button_SetParentShowHint(AObj: TButton; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function Button_GetPopupMenu(AObj: TButton): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Button_SetPopupMenu(AObj: TButton; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function Button_GetShowHint(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure Button_SetShowHint(AObj: TButton; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function Button_GetTabOrder(AObj: TButton): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure Button_SetTabOrder(AObj: TButton; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function Button_GetTabStop(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure Button_SetTabStop(AObj: TButton; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function Button_GetVisible(AObj: TButton): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure Button_SetVisible(AObj: TButton; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure Button_SetOnClick(AObj: TButton; AEventId: NativeUInt); extdecl;
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

procedure Button_SetOnEnter(AObj: TButton; AEventId: NativeUInt); extdecl;
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

procedure Button_SetOnExit(AObj: TButton; AEventId: NativeUInt); extdecl;
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

procedure Button_SetOnKeyDown(AObj: TButton; AEventId: NativeUInt); extdecl;
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

procedure Button_SetOnKeyPress(AObj: TButton; AEventId: NativeUInt); extdecl;
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

procedure Button_SetOnKeyUp(AObj: TButton; AEventId: NativeUInt); extdecl;
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

procedure Button_SetOnMouseDown(AObj: TButton; AEventId: NativeUInt); extdecl;
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

procedure Button_SetOnMouseEnter(AObj: TButton; AEventId: NativeUInt); extdecl;
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

procedure Button_SetOnMouseLeave(AObj: TButton; AEventId: NativeUInt); extdecl;
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

procedure Button_SetOnMouseMove(AObj: TButton; AEventId: NativeUInt); extdecl;
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

procedure Button_SetOnMouseUp(AObj: TButton; AEventId: NativeUInt); extdecl;
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

function Button_GetBrush(AObj: TButton): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function Button_GetControlCount(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function Button_GetHandle(AObj: TButton): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function Button_GetParentWindow(AObj: TButton): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure Button_SetParentWindow(AObj: TButton; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

procedure Button_GetBoundsRect(AObj: TButton; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Button_SetBoundsRect(AObj: TButton; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function Button_GetClientHeight(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Button_SetClientHeight(AObj: TButton; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure Button_GetClientRect(AObj: TButton; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function Button_GetClientWidth(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Button_SetClientWidth(AObj: TButton; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function Button_GetParent(AObj: TButton): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure Button_SetParent(AObj: TButton; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function Button_GetLeft(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure Button_SetLeft(AObj: TButton; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function Button_GetTop(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure Button_SetTop(AObj: TButton; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function Button_GetWidth(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Button_SetWidth(AObj: TButton; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function Button_GetHeight(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Button_SetHeight(AObj: TButton; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Button_GetCursor(AObj: TButton): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure Button_SetCursor(AObj: TButton; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function Button_GetHint(AObj: TButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Button_SetHint(AObj: TButton; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function Button_GetComponentCount(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Button_GetComponentIndex(AObj: TButton): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Button_SetComponentIndex(AObj: TButton; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Button_GetOwner(AObj: TButton): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Button_GetName(AObj: TButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Button_SetName(AObj: TButton; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Button_GetTag(AObj: TButton): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Button_SetTag(AObj: TButton; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Button_GetControls(AObj: TButton; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function Button_GetComponents(AObj: TButton; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Button_Create,
  Button_Free,
  Button_Click,
  Button_CanFocus,
  Button_FlipChildren,
  Button_Focused,
  Button_Invalidate,
  Button_Realign,
  Button_Repaint,
  Button_ScaleBy,
  Button_SetBounds,
  Button_SetFocus,
  Button_Update,
  Button_BringToFront,
  Button_HasParent,
  Button_Hide,
  Button_Perform,
  Button_Refresh,
  Button_SendToBack,
  Button_Show,
  Button_GetTextBuf,
  Button_GetTextLen,
  Button_FindComponent,
  Button_GetNamePath,
  Button_Assign,
  Button_ClassName,
  Button_Equals,
  Button_GetHashCode,
  Button_ToString,
  Button_GetAction,
  Button_SetAction,
  Button_GetAlign,
  Button_SetAlign,
  Button_GetAnchors,
  Button_SetAnchors,
  Button_GetBiDiMode,
  Button_SetBiDiMode,
  Button_GetCancel,
  Button_SetCancel,
  Button_GetCaption,
  Button_SetCaption,
  Button_GetDefault,
  Button_SetDefault,
  Button_GetDoubleBuffered,
  Button_SetDoubleBuffered,
  Button_GetEnabled,
  Button_SetEnabled,
  Button_GetFont,
  Button_SetFont,
  Button_GetModalResult,
  Button_SetModalResult,
  Button_GetParentFont,
  Button_SetParentFont,
  Button_GetParentShowHint,
  Button_SetParentShowHint,
  Button_GetPopupMenu,
  Button_SetPopupMenu,
  Button_GetShowHint,
  Button_SetShowHint,
  Button_GetTabOrder,
  Button_SetTabOrder,
  Button_GetTabStop,
  Button_SetTabStop,
  Button_GetVisible,
  Button_SetVisible,
  Button_SetOnClick,
  Button_SetOnEnter,
  Button_SetOnExit,
  Button_SetOnKeyDown,
  Button_SetOnKeyPress,
  Button_SetOnKeyUp,
  Button_SetOnMouseDown,
  Button_SetOnMouseEnter,
  Button_SetOnMouseLeave,
  Button_SetOnMouseMove,
  Button_SetOnMouseUp,
  Button_GetBrush,
  Button_GetControlCount,
  Button_GetHandle,
  Button_GetParentWindow,
  Button_SetParentWindow,
  Button_GetBoundsRect,
  Button_SetBoundsRect,
  Button_GetClientHeight,
  Button_SetClientHeight,
  Button_GetClientRect,
  Button_GetClientWidth,
  Button_SetClientWidth,
  Button_GetParent,
  Button_SetParent,
  Button_GetLeft,
  Button_SetLeft,
  Button_GetTop,
  Button_SetTop,
  Button_GetWidth,
  Button_SetWidth,
  Button_GetHeight,
  Button_SetHeight,
  Button_GetCursor,
  Button_SetCursor,
  Button_GetHint,
  Button_SetHint,
  Button_GetComponentCount,
  Button_GetComponentIndex,
  Button_SetComponentIndex,
  Button_GetOwner,
  Button_GetName,
  Button_SetName,
  Button_GetTag,
  Button_SetTag,
  Button_GetControls,
  Button_GetComponents;

