
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

function HotKey_ContainsControl(AObj: THotKey; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function HotKey_ControlAtPos(AObj: THotKey; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure HotKey_DisableAlign(AObj: THotKey); stdcall;
begin
  AObj.DisableAlign;
end;

procedure HotKey_EnableAlign(AObj: THotKey); stdcall;
begin
  AObj.EnableAlign;
end;

function HotKey_FindChildControl(AObj: THotKey; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
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

procedure HotKey_InsertControl(AObj: THotKey; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure HotKey_Invalidate(AObj: THotKey); stdcall;
begin
  AObj.Invalidate;
end;

procedure HotKey_PaintTo(AObj: THotKey; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure HotKey_RemoveControl(AObj: THotKey; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
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

procedure HotKey_ScrollBy(AObj: THotKey; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
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

procedure HotKey_UpdateControlState(AObj: THotKey); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure HotKey_BringToFront(AObj: THotKey); stdcall;
begin
  AObj.BringToFront;
end;

procedure HotKey_ClientToScreen(AObj: THotKey; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure HotKey_ClientToParent(AObj: THotKey; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function HotKey_Dragging(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
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

procedure HotKey_ScreenToClient(AObj: THotKey; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure HotKey_ParentToClient(AObj: THotKey; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
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

procedure HotKey_SetTextBuf(AObj: THotKey; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
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

procedure HotKey_DisposeOf(AObj: THotKey); stdcall;
begin
  AObj.DisposeOf;
end;

function HotKey_ClassType(AObj: THotKey): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function HotKey_ClassName(AObj: THotKey): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function HotKey_InstanceSize(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function HotKey_InheritsFrom(AObj: THotKey; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
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
    TEventClass.Remove(AObj, @TEventClass.OnChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, @TEventClass.OnChange, AEventId);
end;

procedure HotKey_SetOnContextPopup(AObj: THotKey; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnContextPopup := nil;
    TEventClass.Remove(AObj, @TEventClass.OnContextPopup);
    Exit;
  end;
  AObj.OnContextPopup := TEventClass.OnContextPopup;
  TEventClass.Add(AObj, @TEventClass.OnContextPopup, AEventId);
end;

procedure HotKey_SetOnEnter(AObj: THotKey; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEnter := nil;
    TEventClass.Remove(AObj, @TEventClass.OnEnter);
    Exit;
  end;
  AObj.OnEnter := TEventClass.OnEnter;
  TEventClass.Add(AObj, @TEventClass.OnEnter, AEventId);
end;

procedure HotKey_SetOnExit(AObj: THotKey; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnExit := nil;
    TEventClass.Remove(AObj, @TEventClass.OnExit);
    Exit;
  end;
  AObj.OnExit := TEventClass.OnExit;
  TEventClass.Add(AObj, @TEventClass.OnExit, AEventId);
end;

procedure HotKey_SetOnGesture(AObj: THotKey; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGesture := nil;
    TEventClass.Remove(AObj, @TEventClass.OnGesture);
    Exit;
  end;
  AObj.OnGesture := TEventClass.OnGesture;
  TEventClass.Add(AObj, @TEventClass.OnGesture, AEventId);
end;

procedure HotKey_SetOnMouseActivate(AObj: THotKey; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseActivate := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseActivate);
    Exit;
  end;
  AObj.OnMouseActivate := TEventClass.OnMouseActivate;
  TEventClass.Add(AObj, @TEventClass.OnMouseActivate, AEventId);
end;

procedure HotKey_SetOnMouseDown(AObj: THotKey; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseDown := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseDown);
    Exit;
  end;
  AObj.OnMouseDown := TEventClass.OnMouseDown;
  TEventClass.Add(AObj, @TEventClass.OnMouseDown, AEventId);
end;

procedure HotKey_SetOnMouseEnter(AObj: THotKey; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseEnter := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseEnter);
    Exit;
  end;
  AObj.OnMouseEnter := TEventClass.OnMouseEnter;
  TEventClass.Add(AObj, @TEventClass.OnMouseEnter, AEventId);
end;

procedure HotKey_SetOnMouseLeave(AObj: THotKey; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseLeave := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseLeave);
    Exit;
  end;
  AObj.OnMouseLeave := TEventClass.OnMouseLeave;
  TEventClass.Add(AObj, @TEventClass.OnMouseLeave, AEventId);
end;

procedure HotKey_SetOnMouseMove(AObj: THotKey; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseMove := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseMove);
    Exit;
  end;
  AObj.OnMouseMove := TEventClass.OnMouseMove;
  TEventClass.Add(AObj, @TEventClass.OnMouseMove, AEventId);
end;

procedure HotKey_SetOnMouseUp(AObj: THotKey; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseUp := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseUp);
    Exit;
  end;
  AObj.OnMouseUp := TEventClass.OnMouseUp;
  TEventClass.Add(AObj, @TEventClass.OnMouseUp, AEventId);
end;

function HotKey_GetDockClientCount(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function HotKey_GetDockSite(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure HotKey_SetDockSite(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function HotKey_GetDoubleBuffered(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure HotKey_SetDoubleBuffered(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function HotKey_GetAlignDisabled(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function HotKey_GetMouseInClient(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function HotKey_GetVisibleDockClientCount(AObj: THotKey): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
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

function HotKey_GetUseDockManager(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure HotKey_SetUseDockManager(AObj: THotKey; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
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

procedure HotKey_GetClientOrigin(AObj: THotKey; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
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

function HotKey_GetControlState(AObj: THotKey): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure HotKey_SetControlState(AObj: THotKey; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function HotKey_GetControlStyle(AObj: THotKey): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure HotKey_SetControlStyle(AObj: THotKey; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
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

function HotKey_GetFloating(AObj: THotKey): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
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

function HotKey_GetDockClients(AObj: THotKey; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function HotKey_GetControls(AObj: THotKey; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function HotKey_GetComponents(AObj: THotKey; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function HotKey_StaticClassType: TClass; stdcall;
begin
  Result := THotKey;
end;

exports
  HotKey_Create,
  HotKey_Free,
  HotKey_CanFocus,
  HotKey_ContainsControl,
  HotKey_ControlAtPos,
  HotKey_DisableAlign,
  HotKey_EnableAlign,
  HotKey_FindChildControl,
  HotKey_FlipChildren,
  HotKey_Focused,
  HotKey_HandleAllocated,
  HotKey_InsertControl,
  HotKey_Invalidate,
  HotKey_PaintTo,
  HotKey_RemoveControl,
  HotKey_Realign,
  HotKey_Repaint,
  HotKey_ScaleBy,
  HotKey_ScrollBy,
  HotKey_SetBounds,
  HotKey_SetFocus,
  HotKey_Update,
  HotKey_UpdateControlState,
  HotKey_BringToFront,
  HotKey_ClientToScreen,
  HotKey_ClientToParent,
  HotKey_Dragging,
  HotKey_HasParent,
  HotKey_Hide,
  HotKey_Perform,
  HotKey_Refresh,
  HotKey_ScreenToClient,
  HotKey_ParentToClient,
  HotKey_SendToBack,
  HotKey_Show,
  HotKey_GetTextBuf,
  HotKey_GetTextLen,
  HotKey_SetTextBuf,
  HotKey_FindComponent,
  HotKey_GetNamePath,
  HotKey_Assign,
  HotKey_DisposeOf,
  HotKey_ClassType,
  HotKey_ClassName,
  HotKey_InstanceSize,
  HotKey_InheritsFrom,
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
  HotKey_SetOnContextPopup,
  HotKey_SetOnEnter,
  HotKey_SetOnExit,
  HotKey_SetOnGesture,
  HotKey_SetOnMouseActivate,
  HotKey_SetOnMouseDown,
  HotKey_SetOnMouseEnter,
  HotKey_SetOnMouseLeave,
  HotKey_SetOnMouseMove,
  HotKey_SetOnMouseUp,
  HotKey_GetDockClientCount,
  HotKey_GetDockSite,
  HotKey_SetDockSite,
  HotKey_GetDoubleBuffered,
  HotKey_SetDoubleBuffered,
  HotKey_GetAlignDisabled,
  HotKey_GetMouseInClient,
  HotKey_GetVisibleDockClientCount,
  HotKey_GetBrush,
  HotKey_GetControlCount,
  HotKey_GetHandle,
  HotKey_GetParentDoubleBuffered,
  HotKey_SetParentDoubleBuffered,
  HotKey_GetParentWindow,
  HotKey_SetParentWindow,
  HotKey_GetUseDockManager,
  HotKey_SetUseDockManager,
  HotKey_GetAction,
  HotKey_SetAction,
  HotKey_GetAlign,
  HotKey_SetAlign,
  HotKey_GetBoundsRect,
  HotKey_SetBoundsRect,
  HotKey_GetClientHeight,
  HotKey_SetClientHeight,
  HotKey_GetClientOrigin,
  HotKey_GetClientRect,
  HotKey_GetClientWidth,
  HotKey_SetClientWidth,
  HotKey_GetControlState,
  HotKey_SetControlState,
  HotKey_GetControlStyle,
  HotKey_SetControlStyle,
  HotKey_GetExplicitLeft,
  HotKey_GetExplicitTop,
  HotKey_GetExplicitWidth,
  HotKey_GetExplicitHeight,
  HotKey_GetFloating,
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
  HotKey_GetDockClients,
  HotKey_GetControls,
  HotKey_GetComponents,
  HotKey_StaticClassType;

