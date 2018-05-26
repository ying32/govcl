
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function GroupBox_Create(AOwner: TComponent): TGroupBox; extdecl;
begin
  Result :=  TGroupBox.Create(AOwner);
end;

procedure GroupBox_Free(AObj: TGroupBox); extdecl;
begin
  AObj.Free;
end;

function GroupBox_CanFocus(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure GroupBox_FlipChildren(AObj: TGroupBox; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function GroupBox_Focused(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure GroupBox_Invalidate(AObj: TGroupBox); extdecl;
begin
  AObj.Invalidate;
end;

procedure GroupBox_Realign(AObj: TGroupBox); extdecl;
begin
  AObj.Realign;
end;

procedure GroupBox_Repaint(AObj: TGroupBox); extdecl;
begin
  AObj.Repaint;
end;

procedure GroupBox_ScaleBy(AObj: TGroupBox; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure GroupBox_SetBounds(AObj: TGroupBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure GroupBox_SetFocus(AObj: TGroupBox); extdecl;
begin
  AObj.SetFocus;
end;

procedure GroupBox_Update(AObj: TGroupBox); extdecl;
begin
  AObj.Update;
end;

procedure GroupBox_BringToFront(AObj: TGroupBox); extdecl;
begin
  AObj.BringToFront;
end;

function GroupBox_HasParent(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure GroupBox_Hide(AObj: TGroupBox); extdecl;
begin
  AObj.Hide;
end;

function GroupBox_Perform(AObj: TGroupBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure GroupBox_Refresh(AObj: TGroupBox); extdecl;
begin
  AObj.Refresh;
end;

procedure GroupBox_SendToBack(AObj: TGroupBox); extdecl;
begin
  AObj.SendToBack;
end;

procedure GroupBox_Show(AObj: TGroupBox); extdecl;
begin
  AObj.Show;
end;

function GroupBox_GetTextBuf(AObj: TGroupBox; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function GroupBox_GetTextLen(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function GroupBox_FindComponent(AObj: TGroupBox; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function GroupBox_GetNamePath(AObj: TGroupBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure GroupBox_Assign(AObj: TGroupBox; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function GroupBox_ClassName(AObj: TGroupBox): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function GroupBox_Equals(AObj: TGroupBox; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function GroupBox_GetHashCode(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function GroupBox_ToString(AObj: TGroupBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function GroupBox_GetAlign(AObj: TGroupBox): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure GroupBox_SetAlign(AObj: TGroupBox; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function GroupBox_GetAnchors(AObj: TGroupBox): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure GroupBox_SetAnchors(AObj: TGroupBox; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function GroupBox_GetBiDiMode(AObj: TGroupBox): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure GroupBox_SetBiDiMode(AObj: TGroupBox; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function GroupBox_GetCaption(AObj: TGroupBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure GroupBox_SetCaption(AObj: TGroupBox; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function GroupBox_GetColor(AObj: TGroupBox): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure GroupBox_SetColor(AObj: TGroupBox; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function GroupBox_GetDoubleBuffered(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure GroupBox_SetDoubleBuffered(AObj: TGroupBox; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function GroupBox_GetEnabled(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure GroupBox_SetEnabled(AObj: TGroupBox; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function GroupBox_GetFont(AObj: TGroupBox): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure GroupBox_SetFont(AObj: TGroupBox; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function GroupBox_GetParentColor(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure GroupBox_SetParentColor(AObj: TGroupBox; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function GroupBox_GetParentFont(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure GroupBox_SetParentFont(AObj: TGroupBox; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function GroupBox_GetParentShowHint(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure GroupBox_SetParentShowHint(AObj: TGroupBox; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function GroupBox_GetPopupMenu(AObj: TGroupBox): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure GroupBox_SetPopupMenu(AObj: TGroupBox; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function GroupBox_GetShowHint(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure GroupBox_SetShowHint(AObj: TGroupBox; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function GroupBox_GetTabOrder(AObj: TGroupBox): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure GroupBox_SetTabOrder(AObj: TGroupBox; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function GroupBox_GetTabStop(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure GroupBox_SetTabStop(AObj: TGroupBox; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function GroupBox_GetVisible(AObj: TGroupBox): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure GroupBox_SetVisible(AObj: TGroupBox; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure GroupBox_SetOnClick(AObj: TGroupBox; AEventId: NativeUInt); extdecl;
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

procedure GroupBox_SetOnDblClick(AObj: TGroupBox; AEventId: NativeUInt); extdecl;
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

procedure GroupBox_SetOnEnter(AObj: TGroupBox; AEventId: NativeUInt); extdecl;
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

procedure GroupBox_SetOnExit(AObj: TGroupBox; AEventId: NativeUInt); extdecl;
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

procedure GroupBox_SetOnMouseDown(AObj: TGroupBox; AEventId: NativeUInt); extdecl;
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

procedure GroupBox_SetOnMouseEnter(AObj: TGroupBox; AEventId: NativeUInt); extdecl;
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

procedure GroupBox_SetOnMouseLeave(AObj: TGroupBox; AEventId: NativeUInt); extdecl;
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

procedure GroupBox_SetOnMouseMove(AObj: TGroupBox; AEventId: NativeUInt); extdecl;
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

procedure GroupBox_SetOnMouseUp(AObj: TGroupBox; AEventId: NativeUInt); extdecl;
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

function GroupBox_GetBrush(AObj: TGroupBox): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function GroupBox_GetControlCount(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function GroupBox_GetHandle(AObj: TGroupBox): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function GroupBox_GetParentWindow(AObj: TGroupBox): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure GroupBox_SetParentWindow(AObj: TGroupBox; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function GroupBox_GetAction(AObj: TGroupBox): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure GroupBox_SetAction(AObj: TGroupBox; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure GroupBox_GetBoundsRect(AObj: TGroupBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure GroupBox_SetBoundsRect(AObj: TGroupBox; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function GroupBox_GetClientHeight(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure GroupBox_SetClientHeight(AObj: TGroupBox; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure GroupBox_GetClientRect(AObj: TGroupBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function GroupBox_GetClientWidth(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure GroupBox_SetClientWidth(AObj: TGroupBox; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function GroupBox_GetParent(AObj: TGroupBox): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure GroupBox_SetParent(AObj: TGroupBox; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function GroupBox_GetLeft(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure GroupBox_SetLeft(AObj: TGroupBox; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function GroupBox_GetTop(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure GroupBox_SetTop(AObj: TGroupBox; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function GroupBox_GetWidth(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure GroupBox_SetWidth(AObj: TGroupBox; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function GroupBox_GetHeight(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure GroupBox_SetHeight(AObj: TGroupBox; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function GroupBox_GetCursor(AObj: TGroupBox): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure GroupBox_SetCursor(AObj: TGroupBox; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function GroupBox_GetHint(AObj: TGroupBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure GroupBox_SetHint(AObj: TGroupBox; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function GroupBox_GetComponentCount(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function GroupBox_GetComponentIndex(AObj: TGroupBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure GroupBox_SetComponentIndex(AObj: TGroupBox; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function GroupBox_GetOwner(AObj: TGroupBox): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function GroupBox_GetName(AObj: TGroupBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure GroupBox_SetName(AObj: TGroupBox; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function GroupBox_GetTag(AObj: TGroupBox): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure GroupBox_SetTag(AObj: TGroupBox; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function GroupBox_GetControls(AObj: TGroupBox; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function GroupBox_GetComponents(AObj: TGroupBox; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  GroupBox_Create,
  GroupBox_Free,
  GroupBox_CanFocus,
  GroupBox_FlipChildren,
  GroupBox_Focused,
  GroupBox_Invalidate,
  GroupBox_Realign,
  GroupBox_Repaint,
  GroupBox_ScaleBy,
  GroupBox_SetBounds,
  GroupBox_SetFocus,
  GroupBox_Update,
  GroupBox_BringToFront,
  GroupBox_HasParent,
  GroupBox_Hide,
  GroupBox_Perform,
  GroupBox_Refresh,
  GroupBox_SendToBack,
  GroupBox_Show,
  GroupBox_GetTextBuf,
  GroupBox_GetTextLen,
  GroupBox_FindComponent,
  GroupBox_GetNamePath,
  GroupBox_Assign,
  GroupBox_ClassName,
  GroupBox_Equals,
  GroupBox_GetHashCode,
  GroupBox_ToString,
  GroupBox_GetAlign,
  GroupBox_SetAlign,
  GroupBox_GetAnchors,
  GroupBox_SetAnchors,
  GroupBox_GetBiDiMode,
  GroupBox_SetBiDiMode,
  GroupBox_GetCaption,
  GroupBox_SetCaption,
  GroupBox_GetColor,
  GroupBox_SetColor,
  GroupBox_GetDoubleBuffered,
  GroupBox_SetDoubleBuffered,
  GroupBox_GetEnabled,
  GroupBox_SetEnabled,
  GroupBox_GetFont,
  GroupBox_SetFont,
  GroupBox_GetParentColor,
  GroupBox_SetParentColor,
  GroupBox_GetParentFont,
  GroupBox_SetParentFont,
  GroupBox_GetParentShowHint,
  GroupBox_SetParentShowHint,
  GroupBox_GetPopupMenu,
  GroupBox_SetPopupMenu,
  GroupBox_GetShowHint,
  GroupBox_SetShowHint,
  GroupBox_GetTabOrder,
  GroupBox_SetTabOrder,
  GroupBox_GetTabStop,
  GroupBox_SetTabStop,
  GroupBox_GetVisible,
  GroupBox_SetVisible,
  GroupBox_SetOnClick,
  GroupBox_SetOnDblClick,
  GroupBox_SetOnEnter,
  GroupBox_SetOnExit,
  GroupBox_SetOnMouseDown,
  GroupBox_SetOnMouseEnter,
  GroupBox_SetOnMouseLeave,
  GroupBox_SetOnMouseMove,
  GroupBox_SetOnMouseUp,
  GroupBox_GetBrush,
  GroupBox_GetControlCount,
  GroupBox_GetHandle,
  GroupBox_GetParentWindow,
  GroupBox_SetParentWindow,
  GroupBox_GetAction,
  GroupBox_SetAction,
  GroupBox_GetBoundsRect,
  GroupBox_SetBoundsRect,
  GroupBox_GetClientHeight,
  GroupBox_SetClientHeight,
  GroupBox_GetClientRect,
  GroupBox_GetClientWidth,
  GroupBox_SetClientWidth,
  GroupBox_GetParent,
  GroupBox_SetParent,
  GroupBox_GetLeft,
  GroupBox_SetLeft,
  GroupBox_GetTop,
  GroupBox_SetTop,
  GroupBox_GetWidth,
  GroupBox_SetWidth,
  GroupBox_GetHeight,
  GroupBox_SetHeight,
  GroupBox_GetCursor,
  GroupBox_SetCursor,
  GroupBox_GetHint,
  GroupBox_SetHint,
  GroupBox_GetComponentCount,
  GroupBox_GetComponentIndex,
  GroupBox_SetComponentIndex,
  GroupBox_GetOwner,
  GroupBox_GetName,
  GroupBox_SetName,
  GroupBox_GetTag,
  GroupBox_SetTag,
  GroupBox_GetControls,
  GroupBox_GetComponents;

