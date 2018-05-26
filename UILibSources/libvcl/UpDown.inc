
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function UpDown_Create(AOwner: TComponent): TUpDown; stdcall;
begin
  Result :=  TUpDown.Create(AOwner);
end;

procedure UpDown_Free(AObj: TUpDown); stdcall;
begin
  AObj.Free;
end;

function UpDown_CanFocus(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure UpDown_FlipChildren(AObj: TUpDown; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function UpDown_Focused(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function UpDown_HandleAllocated(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure UpDown_Invalidate(AObj: TUpDown); stdcall;
begin
  AObj.Invalidate;
end;

procedure UpDown_Realign(AObj: TUpDown); stdcall;
begin
  AObj.Realign;
end;

procedure UpDown_Repaint(AObj: TUpDown); stdcall;
begin
  AObj.Repaint;
end;

procedure UpDown_ScaleBy(AObj: TUpDown; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure UpDown_SetBounds(AObj: TUpDown; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure UpDown_SetFocus(AObj: TUpDown); stdcall;
begin
  AObj.SetFocus;
end;

procedure UpDown_Update(AObj: TUpDown); stdcall;
begin
  AObj.Update;
end;

procedure UpDown_BringToFront(AObj: TUpDown); stdcall;
begin
  AObj.BringToFront;
end;

function UpDown_HasParent(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure UpDown_Hide(AObj: TUpDown); stdcall;
begin
  AObj.Hide;
end;

function UpDown_Perform(AObj: TUpDown; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure UpDown_Refresh(AObj: TUpDown); stdcall;
begin
  AObj.Refresh;
end;

procedure UpDown_SendToBack(AObj: TUpDown); stdcall;
begin
  AObj.SendToBack;
end;

procedure UpDown_Show(AObj: TUpDown); stdcall;
begin
  AObj.Show;
end;

function UpDown_GetTextBuf(AObj: TUpDown; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function UpDown_GetTextLen(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function UpDown_FindComponent(AObj: TUpDown; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function UpDown_GetNamePath(AObj: TUpDown): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure UpDown_Assign(AObj: TUpDown; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function UpDown_ClassName(AObj: TUpDown): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function UpDown_Equals(AObj: TUpDown; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function UpDown_GetHashCode(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function UpDown_ToString(AObj: TUpDown): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function UpDown_GetAnchors(AObj: TUpDown): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure UpDown_SetAnchors(AObj: TUpDown; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function UpDown_GetDoubleBuffered(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure UpDown_SetDoubleBuffered(AObj: TUpDown; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function UpDown_GetEnabled(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure UpDown_SetEnabled(AObj: TUpDown; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function UpDown_GetHint(AObj: TUpDown): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure UpDown_SetHint(AObj: TUpDown; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function UpDown_GetMin(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.Min;
end;

procedure UpDown_SetMin(AObj: TUpDown; AValue: Integer); stdcall;
begin
  AObj.Min := AValue;
end;

function UpDown_GetMax(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.Max;
end;

procedure UpDown_SetMax(AObj: TUpDown; AValue: Integer); stdcall;
begin
  AObj.Max := AValue;
end;

function UpDown_GetOrientation(AObj: TUpDown): TUDOrientation; stdcall;
begin
  Result :=  AObj.Orientation;
end;

procedure UpDown_SetOrientation(AObj: TUpDown; AValue: TUDOrientation); stdcall;
begin
  AObj.Orientation := AValue;
end;

function UpDown_GetParentDoubleBuffered(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure UpDown_SetParentDoubleBuffered(AObj: TUpDown; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function UpDown_GetParentShowHint(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure UpDown_SetParentShowHint(AObj: TUpDown; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function UpDown_GetPopupMenu(AObj: TUpDown): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure UpDown_SetPopupMenu(AObj: TUpDown; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function UpDown_GetPosition(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.Position;
end;

procedure UpDown_SetPosition(AObj: TUpDown; AValue: Integer); stdcall;
begin
  AObj.Position := AValue;
end;

function UpDown_GetShowHint(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure UpDown_SetShowHint(AObj: TUpDown; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function UpDown_GetTabOrder(AObj: TUpDown): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure UpDown_SetTabOrder(AObj: TUpDown; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function UpDown_GetTabStop(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure UpDown_SetTabStop(AObj: TUpDown; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function UpDown_GetVisible(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure UpDown_SetVisible(AObj: TUpDown; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function UpDown_GetWrap(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.Wrap;
end;

procedure UpDown_SetWrap(AObj: TUpDown; AValue: LongBool); stdcall;
begin
  AObj.Wrap := AValue;
end;

function UpDown_GetStyleElements(AObj: TUpDown): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure UpDown_SetStyleElements(AObj: TUpDown; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure UpDown_SetOnClick(AObj: TUpDown; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    TEventClass.Remove(AObj, geUpDownClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.UpDownOnClick;
  TEventClass.Add(AObj, geUpDownClick, AEventId);
end;

procedure UpDown_SetOnEnter(AObj: TUpDown; AEventId: NativeUInt); stdcall;
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

procedure UpDown_SetOnExit(AObj: TUpDown; AEventId: NativeUInt); stdcall;
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

procedure UpDown_SetOnMouseDown(AObj: TUpDown; AEventId: NativeUInt); stdcall;
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

procedure UpDown_SetOnMouseEnter(AObj: TUpDown; AEventId: NativeUInt); stdcall;
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

procedure UpDown_SetOnMouseLeave(AObj: TUpDown; AEventId: NativeUInt); stdcall;
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

procedure UpDown_SetOnMouseMove(AObj: TUpDown; AEventId: NativeUInt); stdcall;
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

procedure UpDown_SetOnMouseUp(AObj: TUpDown; AEventId: NativeUInt); stdcall;
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

function UpDown_GetBrush(AObj: TUpDown): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function UpDown_GetControlCount(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function UpDown_GetHandle(AObj: TUpDown): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function UpDown_GetParentWindow(AObj: TUpDown): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure UpDown_SetParentWindow(AObj: TUpDown; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function UpDown_GetAction(AObj: TUpDown): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure UpDown_SetAction(AObj: TUpDown; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function UpDown_GetAlign(AObj: TUpDown): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure UpDown_SetAlign(AObj: TUpDown; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function UpDown_GetBiDiMode(AObj: TUpDown): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure UpDown_SetBiDiMode(AObj: TUpDown; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure UpDown_GetBoundsRect(AObj: TUpDown; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure UpDown_SetBoundsRect(AObj: TUpDown; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function UpDown_GetClientHeight(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure UpDown_SetClientHeight(AObj: TUpDown; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure UpDown_GetClientRect(AObj: TUpDown; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function UpDown_GetClientWidth(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure UpDown_SetClientWidth(AObj: TUpDown; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function UpDown_GetExplicitLeft(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function UpDown_GetExplicitTop(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function UpDown_GetExplicitWidth(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function UpDown_GetExplicitHeight(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function UpDown_GetParent(AObj: TUpDown): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure UpDown_SetParent(AObj: TUpDown; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function UpDown_GetAlignWithMargins(AObj: TUpDown): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure UpDown_SetAlignWithMargins(AObj: TUpDown; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function UpDown_GetLeft(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure UpDown_SetLeft(AObj: TUpDown; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function UpDown_GetTop(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure UpDown_SetTop(AObj: TUpDown; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function UpDown_GetWidth(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure UpDown_SetWidth(AObj: TUpDown; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function UpDown_GetHeight(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure UpDown_SetHeight(AObj: TUpDown; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function UpDown_GetCursor(AObj: TUpDown): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure UpDown_SetCursor(AObj: TUpDown; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function UpDown_GetMargins(AObj: TUpDown): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure UpDown_SetMargins(AObj: TUpDown; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function UpDown_GetCustomHint(AObj: TUpDown): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure UpDown_SetCustomHint(AObj: TUpDown; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function UpDown_GetComponentCount(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function UpDown_GetComponentIndex(AObj: TUpDown): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure UpDown_SetComponentIndex(AObj: TUpDown; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function UpDown_GetOwner(AObj: TUpDown): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function UpDown_GetName(AObj: TUpDown): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure UpDown_SetName(AObj: TUpDown; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function UpDown_GetTag(AObj: TUpDown): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure UpDown_SetTag(AObj: TUpDown; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function UpDown_GetControls(AObj: TUpDown; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function UpDown_GetComponents(AObj: TUpDown; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  UpDown_Create,
  UpDown_Free,
  UpDown_CanFocus,
  UpDown_FlipChildren,
  UpDown_Focused,
  UpDown_HandleAllocated,
  UpDown_Invalidate,
  UpDown_Realign,
  UpDown_Repaint,
  UpDown_ScaleBy,
  UpDown_SetBounds,
  UpDown_SetFocus,
  UpDown_Update,
  UpDown_BringToFront,
  UpDown_HasParent,
  UpDown_Hide,
  UpDown_Perform,
  UpDown_Refresh,
  UpDown_SendToBack,
  UpDown_Show,
  UpDown_GetTextBuf,
  UpDown_GetTextLen,
  UpDown_FindComponent,
  UpDown_GetNamePath,
  UpDown_Assign,
  UpDown_ClassName,
  UpDown_Equals,
  UpDown_GetHashCode,
  UpDown_ToString,
  UpDown_GetAnchors,
  UpDown_SetAnchors,
  UpDown_GetDoubleBuffered,
  UpDown_SetDoubleBuffered,
  UpDown_GetEnabled,
  UpDown_SetEnabled,
  UpDown_GetHint,
  UpDown_SetHint,
  UpDown_GetMin,
  UpDown_SetMin,
  UpDown_GetMax,
  UpDown_SetMax,
  UpDown_GetOrientation,
  UpDown_SetOrientation,
  UpDown_GetParentDoubleBuffered,
  UpDown_SetParentDoubleBuffered,
  UpDown_GetParentShowHint,
  UpDown_SetParentShowHint,
  UpDown_GetPopupMenu,
  UpDown_SetPopupMenu,
  UpDown_GetPosition,
  UpDown_SetPosition,
  UpDown_GetShowHint,
  UpDown_SetShowHint,
  UpDown_GetTabOrder,
  UpDown_SetTabOrder,
  UpDown_GetTabStop,
  UpDown_SetTabStop,
  UpDown_GetVisible,
  UpDown_SetVisible,
  UpDown_GetWrap,
  UpDown_SetWrap,
  UpDown_GetStyleElements,
  UpDown_SetStyleElements,
  UpDown_SetOnClick,
  UpDown_SetOnEnter,
  UpDown_SetOnExit,
  UpDown_SetOnMouseDown,
  UpDown_SetOnMouseEnter,
  UpDown_SetOnMouseLeave,
  UpDown_SetOnMouseMove,
  UpDown_SetOnMouseUp,
  UpDown_GetBrush,
  UpDown_GetControlCount,
  UpDown_GetHandle,
  UpDown_GetParentWindow,
  UpDown_SetParentWindow,
  UpDown_GetAction,
  UpDown_SetAction,
  UpDown_GetAlign,
  UpDown_SetAlign,
  UpDown_GetBiDiMode,
  UpDown_SetBiDiMode,
  UpDown_GetBoundsRect,
  UpDown_SetBoundsRect,
  UpDown_GetClientHeight,
  UpDown_SetClientHeight,
  UpDown_GetClientRect,
  UpDown_GetClientWidth,
  UpDown_SetClientWidth,
  UpDown_GetExplicitLeft,
  UpDown_GetExplicitTop,
  UpDown_GetExplicitWidth,
  UpDown_GetExplicitHeight,
  UpDown_GetParent,
  UpDown_SetParent,
  UpDown_GetAlignWithMargins,
  UpDown_SetAlignWithMargins,
  UpDown_GetLeft,
  UpDown_SetLeft,
  UpDown_GetTop,
  UpDown_SetTop,
  UpDown_GetWidth,
  UpDown_SetWidth,
  UpDown_GetHeight,
  UpDown_SetHeight,
  UpDown_GetCursor,
  UpDown_SetCursor,
  UpDown_GetMargins,
  UpDown_SetMargins,
  UpDown_GetCustomHint,
  UpDown_SetCustomHint,
  UpDown_GetComponentCount,
  UpDown_GetComponentIndex,
  UpDown_SetComponentIndex,
  UpDown_GetOwner,
  UpDown_GetName,
  UpDown_SetName,
  UpDown_GetTag,
  UpDown_SetTag,
  UpDown_GetControls,
  UpDown_GetComponents;

