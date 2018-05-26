
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function HotKey_Create(AOwner: TComponent): THotKey; stdcall;
begin
  Result :=  THotKey.Create(AOwner);
end;

procedure HotKey_Free(AObj: THotKey); stdcall;
begin
  AObj.Free;
end;

function HotKey_CanFocus(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure HotKey_FlipChildren(AObj: THotKey; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function HotKey_Focused(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function HotKey_HandleAllocated(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure HotKey_Invalidate(AObj: THotKey); stdcall;
begin
  AObj.Invalidate;
end;

procedure HotKey_Realign(AObj: THotKey); stdcall;
begin
  AObj.Realign;
end;

procedure HotKey_Repaint(AObj: THotKey); stdcall;
begin
  AObj.Repaint;
end;

procedure HotKey_ScaleBy(AObj: THotKey; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure HotKey_SetBounds(AObj: THotKey; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure HotKey_SetFocus(AObj: THotKey); stdcall;
begin
  AObj.SetFocus;
end;

procedure HotKey_Update(AObj: THotKey); stdcall;
begin
  AObj.Update;
end;

procedure HotKey_BringToFront(AObj: THotKey); stdcall;
begin
  AObj.BringToFront;
end;

function HotKey_HasParent(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure HotKey_Hide(AObj: THotKey); stdcall;
begin
  AObj.Hide;
end;

function HotKey_Perform(AObj: THotKey; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure HotKey_Refresh(AObj: THotKey); stdcall;
begin
  AObj.Refresh;
end;

procedure HotKey_SendToBack(AObj: THotKey); stdcall;
begin
  AObj.SendToBack;
end;

procedure HotKey_Show(AObj: THotKey); stdcall;
begin
  AObj.Show;
end;

function HotKey_GetTextBuf(AObj: THotKey; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function HotKey_GetTextLen(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function HotKey_FindComponent(AObj: THotKey; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function HotKey_GetNamePath(AObj: THotKey): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure HotKey_Assign(AObj: THotKey; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function HotKey_ClassName(AObj: THotKey): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function HotKey_Equals(AObj: THotKey; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function HotKey_GetHashCode(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function HotKey_ToString(AObj: THotKey): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function HotKey_GetAnchors(AObj: THotKey): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure HotKey_SetAnchors(AObj: THotKey; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function HotKey_GetAutoSize(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure HotKey_SetAutoSize(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function HotKey_GetBiDiMode(AObj: THotKey): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure HotKey_SetBiDiMode(AObj: THotKey; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function HotKey_GetEnabled(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure HotKey_SetEnabled(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function HotKey_GetHint(AObj: THotKey): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure HotKey_SetHint(AObj: THotKey; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function HotKey_GetParentShowHint(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure HotKey_SetParentShowHint(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function HotKey_GetPopupMenu(AObj: THotKey): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure HotKey_SetPopupMenu(AObj: THotKey; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function HotKey_GetShowHint(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure HotKey_SetShowHint(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function HotKey_GetTabOrder(AObj: THotKey): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure HotKey_SetTabOrder(AObj: THotKey; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function HotKey_GetTabStop(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure HotKey_SetTabStop(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function HotKey_GetVisible(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure HotKey_SetVisible(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function HotKey_GetStyleElements(AObj: THotKey): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure HotKey_SetStyleElements(AObj: THotKey; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure HotKey_SetOnChange(AObj: THotKey; AEventId: NativeUInt); stdcall;
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

procedure HotKey_SetOnEnter(AObj: THotKey; AEventId: NativeUInt); stdcall;
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

procedure HotKey_SetOnExit(AObj: THotKey; AEventId: NativeUInt); stdcall;
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

procedure HotKey_SetOnMouseDown(AObj: THotKey; AEventId: NativeUInt); stdcall;
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

procedure HotKey_SetOnMouseEnter(AObj: THotKey; AEventId: NativeUInt); stdcall;
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

procedure HotKey_SetOnMouseLeave(AObj: THotKey; AEventId: NativeUInt); stdcall;
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

procedure HotKey_SetOnMouseMove(AObj: THotKey; AEventId: NativeUInt); stdcall;
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

procedure HotKey_SetOnMouseUp(AObj: THotKey; AEventId: NativeUInt); stdcall;
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

function HotKey_GetDoubleBuffered(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure HotKey_SetDoubleBuffered(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function HotKey_GetBrush(AObj: THotKey): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function HotKey_GetControlCount(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function HotKey_GetHandle(AObj: THotKey): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function HotKey_GetParentDoubleBuffered(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure HotKey_SetParentDoubleBuffered(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function HotKey_GetParentWindow(AObj: THotKey): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure HotKey_SetParentWindow(AObj: THotKey; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function HotKey_GetAction(AObj: THotKey): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure HotKey_SetAction(AObj: THotKey; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function HotKey_GetAlign(AObj: THotKey): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure HotKey_SetAlign(AObj: THotKey; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

procedure HotKey_GetBoundsRect(AObj: THotKey; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure HotKey_SetBoundsRect(AObj: THotKey; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function HotKey_GetClientHeight(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure HotKey_SetClientHeight(AObj: THotKey; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure HotKey_GetClientRect(AObj: THotKey; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function HotKey_GetClientWidth(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure HotKey_SetClientWidth(AObj: THotKey; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function HotKey_GetExplicitLeft(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function HotKey_GetExplicitTop(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function HotKey_GetExplicitWidth(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function HotKey_GetExplicitHeight(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function HotKey_GetParent(AObj: THotKey): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure HotKey_SetParent(AObj: THotKey; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function HotKey_GetAlignWithMargins(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure HotKey_SetAlignWithMargins(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function HotKey_GetLeft(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure HotKey_SetLeft(AObj: THotKey; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function HotKey_GetTop(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure HotKey_SetTop(AObj: THotKey; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function HotKey_GetWidth(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure HotKey_SetWidth(AObj: THotKey; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function HotKey_GetHeight(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure HotKey_SetHeight(AObj: THotKey; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function HotKey_GetCursor(AObj: THotKey): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure HotKey_SetCursor(AObj: THotKey; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function HotKey_GetMargins(AObj: THotKey): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure HotKey_SetMargins(AObj: THotKey; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function HotKey_GetCustomHint(AObj: THotKey): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure HotKey_SetCustomHint(AObj: THotKey; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function HotKey_GetComponentCount(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function HotKey_GetComponentIndex(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure HotKey_SetComponentIndex(AObj: THotKey; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function HotKey_GetOwner(AObj: THotKey): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function HotKey_GetName(AObj: THotKey): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure HotKey_SetName(AObj: THotKey; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function HotKey_GetTag(AObj: THotKey): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure HotKey_SetTag(AObj: THotKey; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function HotKey_GetControls(AObj: THotKey; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function HotKey_GetComponents(AObj: THotKey; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  HotKey_Create,
  HotKey_Free,
  HotKey_CanFocus,
  HotKey_FlipChildren,
  HotKey_Focused,
  HotKey_HandleAllocated,
  HotKey_Invalidate,
  HotKey_Realign,
  HotKey_Repaint,
  HotKey_ScaleBy,
  HotKey_SetBounds,
  HotKey_SetFocus,
  HotKey_Update,
  HotKey_BringToFront,
  HotKey_HasParent,
  HotKey_Hide,
  HotKey_Perform,
  HotKey_Refresh,
  HotKey_SendToBack,
  HotKey_Show,
  HotKey_GetTextBuf,
  HotKey_GetTextLen,
  HotKey_FindComponent,
  HotKey_GetNamePath,
  HotKey_Assign,
  HotKey_ClassName,
  HotKey_Equals,
  HotKey_GetHashCode,
  HotKey_ToString,
  HotKey_GetAnchors,
  HotKey_SetAnchors,
  HotKey_GetAutoSize,
  HotKey_SetAutoSize,
  HotKey_GetBiDiMode,
  HotKey_SetBiDiMode,
  HotKey_GetEnabled,
  HotKey_SetEnabled,
  HotKey_GetHint,
  HotKey_SetHint,
  HotKey_GetParentShowHint,
  HotKey_SetParentShowHint,
  HotKey_GetPopupMenu,
  HotKey_SetPopupMenu,
  HotKey_GetShowHint,
  HotKey_SetShowHint,
  HotKey_GetTabOrder,
  HotKey_SetTabOrder,
  HotKey_GetTabStop,
  HotKey_SetTabStop,
  HotKey_GetVisible,
  HotKey_SetVisible,
  HotKey_GetStyleElements,
  HotKey_SetStyleElements,
  HotKey_SetOnChange,
  HotKey_SetOnEnter,
  HotKey_SetOnExit,
  HotKey_SetOnMouseDown,
  HotKey_SetOnMouseEnter,
  HotKey_SetOnMouseLeave,
  HotKey_SetOnMouseMove,
  HotKey_SetOnMouseUp,
  HotKey_GetDoubleBuffered,
  HotKey_SetDoubleBuffered,
  HotKey_GetBrush,
  HotKey_GetControlCount,
  HotKey_GetHandle,
  HotKey_GetParentDoubleBuffered,
  HotKey_SetParentDoubleBuffered,
  HotKey_GetParentWindow,
  HotKey_SetParentWindow,
  HotKey_GetAction,
  HotKey_SetAction,
  HotKey_GetAlign,
  HotKey_SetAlign,
  HotKey_GetBoundsRect,
  HotKey_SetBoundsRect,
  HotKey_GetClientHeight,
  HotKey_SetClientHeight,
  HotKey_GetClientRect,
  HotKey_GetClientWidth,
  HotKey_SetClientWidth,
  HotKey_GetExplicitLeft,
  HotKey_GetExplicitTop,
  HotKey_GetExplicitWidth,
  HotKey_GetExplicitHeight,
  HotKey_GetParent,
  HotKey_SetParent,
  HotKey_GetAlignWithMargins,
  HotKey_SetAlignWithMargins,
  HotKey_GetLeft,
  HotKey_SetLeft,
  HotKey_GetTop,
  HotKey_SetTop,
  HotKey_GetWidth,
  HotKey_SetWidth,
  HotKey_GetHeight,
  HotKey_SetHeight,
  HotKey_GetCursor,
  HotKey_SetCursor,
  HotKey_GetMargins,
  HotKey_SetMargins,
  HotKey_GetCustomHint,
  HotKey_SetCustomHint,
  HotKey_GetComponentCount,
  HotKey_GetComponentIndex,
  HotKey_SetComponentIndex,
  HotKey_GetOwner,
  HotKey_GetName,
  HotKey_SetName,
  HotKey_GetTag,
  HotKey_SetTag,
  HotKey_GetControls,
  HotKey_GetComponents;

