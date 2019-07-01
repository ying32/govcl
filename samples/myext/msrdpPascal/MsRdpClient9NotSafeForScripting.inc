
//----------------------------------------
// 代码由GenlibVcl子工具GenCode自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function MsRdpClient9NotSafeForScripting_Create(AOwner: TComponent): TMsRdpClient9NotSafeForScripting; stdcall;
begin
  Result :=  TMsRdpClient9NotSafeForScripting.Create(AOwner);
end;

procedure MsRdpClient9NotSafeForScripting_Free(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.Free;
end;

procedure MsRdpClient9NotSafeForScripting_SetBounds(AObj: TMsRdpClient9NotSafeForScripting; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

function MsRdpClient9NotSafeForScripting_CanFocus(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function MsRdpClient9NotSafeForScripting_ContainsControl(AObj: TMsRdpClient9NotSafeForScripting; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function MsRdpClient9NotSafeForScripting_ControlAtPos(AObj: TMsRdpClient9NotSafeForScripting; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure MsRdpClient9NotSafeForScripting_DisableAlign(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.DisableAlign;
end;

procedure MsRdpClient9NotSafeForScripting_EnableAlign(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.EnableAlign;
end;

function MsRdpClient9NotSafeForScripting_FindChildControl(AObj: TMsRdpClient9NotSafeForScripting; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure MsRdpClient9NotSafeForScripting_FlipChildren(AObj: TMsRdpClient9NotSafeForScripting; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function MsRdpClient9NotSafeForScripting_Focused(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function MsRdpClient9NotSafeForScripting_HandleAllocated(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure MsRdpClient9NotSafeForScripting_InsertControl(AObj: TMsRdpClient9NotSafeForScripting; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure MsRdpClient9NotSafeForScripting_Invalidate(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.Invalidate;
end;

procedure MsRdpClient9NotSafeForScripting_PaintTo(AObj: TMsRdpClient9NotSafeForScripting; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure MsRdpClient9NotSafeForScripting_RemoveControl(AObj: TMsRdpClient9NotSafeForScripting; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure MsRdpClient9NotSafeForScripting_Realign(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.Realign;
end;

procedure MsRdpClient9NotSafeForScripting_Repaint(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.Repaint;
end;

procedure MsRdpClient9NotSafeForScripting_ScaleBy(AObj: TMsRdpClient9NotSafeForScripting; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure MsRdpClient9NotSafeForScripting_ScrollBy(AObj: TMsRdpClient9NotSafeForScripting; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure MsRdpClient9NotSafeForScripting_SetFocus(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.SetFocus;
end;

procedure MsRdpClient9NotSafeForScripting_Update(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.Update;
end;

procedure MsRdpClient9NotSafeForScripting_UpdateControlState(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure MsRdpClient9NotSafeForScripting_BringToFront(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.BringToFront;
end;

procedure MsRdpClient9NotSafeForScripting_ClientToScreen(AObj: TMsRdpClient9NotSafeForScripting; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure MsRdpClient9NotSafeForScripting_ClientToParent(AObj: TMsRdpClient9NotSafeForScripting; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function MsRdpClient9NotSafeForScripting_Dragging(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function MsRdpClient9NotSafeForScripting_HasParent(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure MsRdpClient9NotSafeForScripting_Hide(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.Hide;
end;

function MsRdpClient9NotSafeForScripting_Perform(AObj: TMsRdpClient9NotSafeForScripting; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure MsRdpClient9NotSafeForScripting_Refresh(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.Refresh;
end;

procedure MsRdpClient9NotSafeForScripting_ScreenToClient(AObj: TMsRdpClient9NotSafeForScripting; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure MsRdpClient9NotSafeForScripting_ParentToClient(AObj: TMsRdpClient9NotSafeForScripting; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure MsRdpClient9NotSafeForScripting_SendToBack(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.SendToBack;
end;

procedure MsRdpClient9NotSafeForScripting_Show(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.Show;
end;

function MsRdpClient9NotSafeForScripting_GetTextBuf(AObj: TMsRdpClient9NotSafeForScripting; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function MsRdpClient9NotSafeForScripting_GetTextLen(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure MsRdpClient9NotSafeForScripting_SetTextBuf(AObj: TMsRdpClient9NotSafeForScripting; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function MsRdpClient9NotSafeForScripting_FindComponent(AObj: TMsRdpClient9NotSafeForScripting; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function MsRdpClient9NotSafeForScripting_GetNamePath(AObj: TMsRdpClient9NotSafeForScripting): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure MsRdpClient9NotSafeForScripting_Assign(AObj: TMsRdpClient9NotSafeForScripting; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure MsRdpClient9NotSafeForScripting_DisposeOf(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.DisposeOf;
end;

function MsRdpClient9NotSafeForScripting_ClassType(AObj: TMsRdpClient9NotSafeForScripting): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function MsRdpClient9NotSafeForScripting_ClassName(AObj: TMsRdpClient9NotSafeForScripting): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function MsRdpClient9NotSafeForScripting_InstanceSize(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function MsRdpClient9NotSafeForScripting_InheritsFrom(AObj: TMsRdpClient9NotSafeForScripting; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function MsRdpClient9NotSafeForScripting_Equals(AObj: TMsRdpClient9NotSafeForScripting; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function MsRdpClient9NotSafeForScripting_GetHashCode(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function MsRdpClient9NotSafeForScripting_ToString(AObj: TMsRdpClient9NotSafeForScripting): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function MsRdpClient9NotSafeForScripting_GetControlInterface(AObj: TMsRdpClient9NotSafeForScripting): Pointer; stdcall;
begin
  Result :=  AObj.ControlInterface;
end;

function MsRdpClient9NotSafeForScripting_GetDefaultInterface(AObj: TMsRdpClient9NotSafeForScripting): Pointer; stdcall;
begin
  Result :=  AObj.DefaultInterface;
end;

function MsRdpClient9NotSafeForScripting_GetVersion(AObj: TMsRdpClient9NotSafeForScripting): PChar; stdcall;
begin
  Result :=  PChar(AObj.Version);
end;

function MsRdpClient9NotSafeForScripting_GetAdvancedSettings9(AObj: TMsRdpClient9NotSafeForScripting): Pointer; stdcall;
begin
  Result :=  AObj.AdvancedSettings9;
end;

function MsRdpClient9NotSafeForScripting_GetAnchors(AObj: TMsRdpClient9NotSafeForScripting): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure MsRdpClient9NotSafeForScripting_SetAnchors(AObj: TMsRdpClient9NotSafeForScripting; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetTabStop(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure MsRdpClient9NotSafeForScripting_SetTabStop(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetAlign(AObj: TMsRdpClient9NotSafeForScripting): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure MsRdpClient9NotSafeForScripting_SetAlign(AObj: TMsRdpClient9NotSafeForScripting; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetDragCursor(AObj: TMsRdpClient9NotSafeForScripting): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure MsRdpClient9NotSafeForScripting_SetDragCursor(AObj: TMsRdpClient9NotSafeForScripting; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetDragMode(AObj: TMsRdpClient9NotSafeForScripting): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure MsRdpClient9NotSafeForScripting_SetDragMode(AObj: TMsRdpClient9NotSafeForScripting; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetParentShowHint(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure MsRdpClient9NotSafeForScripting_SetParentShowHint(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetPopupMenu(AObj: TMsRdpClient9NotSafeForScripting): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure MsRdpClient9NotSafeForScripting_SetPopupMenu(AObj: TMsRdpClient9NotSafeForScripting; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetShowHint(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure MsRdpClient9NotSafeForScripting_SetShowHint(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetTabOrder(AObj: TMsRdpClient9NotSafeForScripting): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure MsRdpClient9NotSafeForScripting_SetTabOrder(AObj: TMsRdpClient9NotSafeForScripting; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetVisible(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure MsRdpClient9NotSafeForScripting_SetVisible(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

procedure MsRdpClient9NotSafeForScripting_SetOnDragDrop(AObj: TMsRdpClient9NotSafeForScripting; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragDrop := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDragDrop);
    Exit;
  end;
  AObj.OnDragDrop := TEventClass.OnDragDrop;
  TEventClass.Add(AObj, @TEventClass.OnDragDrop, AEventId);
end;

procedure MsRdpClient9NotSafeForScripting_SetOnDragOver(AObj: TMsRdpClient9NotSafeForScripting; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragOver := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDragOver);
    Exit;
  end;
  AObj.OnDragOver := TEventClass.OnDragOver;
  TEventClass.Add(AObj, @TEventClass.OnDragOver, AEventId);
end;

procedure MsRdpClient9NotSafeForScripting_SetOnEndDrag(AObj: TMsRdpClient9NotSafeForScripting; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDrag := nil;
    TEventClass.Remove(AObj, @TEventClass.OnEndDrag);
    Exit;
  end;
  AObj.OnEndDrag := TEventClass.OnEndDrag;
  TEventClass.Add(AObj, @TEventClass.OnEndDrag, AEventId);
end;

procedure MsRdpClient9NotSafeForScripting_SetOnEnter(AObj: TMsRdpClient9NotSafeForScripting; AEventId: NativeUInt); stdcall;
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

procedure MsRdpClient9NotSafeForScripting_SetOnExit(AObj: TMsRdpClient9NotSafeForScripting; AEventId: NativeUInt); stdcall;
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

function MsRdpClient9NotSafeForScripting_GetDesktopWidth(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.DesktopWidth;
end;

procedure MsRdpClient9NotSafeForScripting_SetDesktopWidth(AObj: TMsRdpClient9NotSafeForScripting; AValue: Integer); stdcall;
begin
  AObj.DesktopWidth := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetDesktopHeight(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.DesktopHeight;
end;

procedure MsRdpClient9NotSafeForScripting_SetDesktopHeight(AObj: TMsRdpClient9NotSafeForScripting; AValue: Integer); stdcall;
begin
  AObj.DesktopHeight := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetColorDepth(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.ColorDepth;
end;

procedure MsRdpClient9NotSafeForScripting_SetColorDepth(AObj: TMsRdpClient9NotSafeForScripting; AValue: Integer); stdcall;
begin
  AObj.ColorDepth := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetFullScreen(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.FullScreen;
end;

procedure MsRdpClient9NotSafeForScripting_SetFullScreen(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.FullScreen := AValue;
end;

procedure MsRdpClient9NotSafeForScripting_SetOnDisconnected(AObj: TMsRdpClient9NotSafeForScripting; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDisconnected := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDisconnected);
    Exit;
  end;
  AObj.OnDisconnected := TEventClass.OnDisconnected;
  TEventClass.Add(AObj, @TEventClass.OnDisconnected, AEventId);
end;

function MsRdpClient9NotSafeForScripting_GetDockClientCount(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function MsRdpClient9NotSafeForScripting_GetDockSite(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure MsRdpClient9NotSafeForScripting_SetDockSite(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetDoubleBuffered(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure MsRdpClient9NotSafeForScripting_SetDoubleBuffered(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetAlignDisabled(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function MsRdpClient9NotSafeForScripting_GetMouseInClient(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function MsRdpClient9NotSafeForScripting_GetVisibleDockClientCount(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function MsRdpClient9NotSafeForScripting_GetBrush(AObj: TMsRdpClient9NotSafeForScripting): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function MsRdpClient9NotSafeForScripting_GetControlCount(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function MsRdpClient9NotSafeForScripting_GetHandle(AObj: TMsRdpClient9NotSafeForScripting): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function MsRdpClient9NotSafeForScripting_GetParentDoubleBuffered(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure MsRdpClient9NotSafeForScripting_SetParentDoubleBuffered(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetParentWindow(AObj: TMsRdpClient9NotSafeForScripting): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure MsRdpClient9NotSafeForScripting_SetParentWindow(AObj: TMsRdpClient9NotSafeForScripting; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetUseDockManager(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure MsRdpClient9NotSafeForScripting_SetUseDockManager(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetEnabled(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure MsRdpClient9NotSafeForScripting_SetEnabled(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetAction(AObj: TMsRdpClient9NotSafeForScripting): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure MsRdpClient9NotSafeForScripting_SetAction(AObj: TMsRdpClient9NotSafeForScripting; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetBiDiMode(AObj: TMsRdpClient9NotSafeForScripting): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure MsRdpClient9NotSafeForScripting_SetBiDiMode(AObj: TMsRdpClient9NotSafeForScripting; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure MsRdpClient9NotSafeForScripting_GetBoundsRect(AObj: TMsRdpClient9NotSafeForScripting; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure MsRdpClient9NotSafeForScripting_SetBoundsRect(AObj: TMsRdpClient9NotSafeForScripting; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetClientHeight(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure MsRdpClient9NotSafeForScripting_SetClientHeight(AObj: TMsRdpClient9NotSafeForScripting; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure MsRdpClient9NotSafeForScripting_GetClientOrigin(AObj: TMsRdpClient9NotSafeForScripting; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure MsRdpClient9NotSafeForScripting_GetClientRect(AObj: TMsRdpClient9NotSafeForScripting; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function MsRdpClient9NotSafeForScripting_GetClientWidth(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure MsRdpClient9NotSafeForScripting_SetClientWidth(AObj: TMsRdpClient9NotSafeForScripting; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetControlState(AObj: TMsRdpClient9NotSafeForScripting): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure MsRdpClient9NotSafeForScripting_SetControlState(AObj: TMsRdpClient9NotSafeForScripting; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetControlStyle(AObj: TMsRdpClient9NotSafeForScripting): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure MsRdpClient9NotSafeForScripting_SetControlStyle(AObj: TMsRdpClient9NotSafeForScripting; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetExplicitLeft(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function MsRdpClient9NotSafeForScripting_GetExplicitTop(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function MsRdpClient9NotSafeForScripting_GetExplicitWidth(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function MsRdpClient9NotSafeForScripting_GetExplicitHeight(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function MsRdpClient9NotSafeForScripting_GetFloating(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function MsRdpClient9NotSafeForScripting_GetParent(AObj: TMsRdpClient9NotSafeForScripting): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure MsRdpClient9NotSafeForScripting_SetParent(AObj: TMsRdpClient9NotSafeForScripting; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetStyleElements(AObj: TMsRdpClient9NotSafeForScripting): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure MsRdpClient9NotSafeForScripting_SetStyleElements(AObj: TMsRdpClient9NotSafeForScripting; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure MsRdpClient9NotSafeForScripting_SetOnGesture(AObj: TMsRdpClient9NotSafeForScripting; AEventId: NativeUInt); stdcall;
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

function MsRdpClient9NotSafeForScripting_GetAlignWithMargins(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure MsRdpClient9NotSafeForScripting_SetAlignWithMargins(AObj: TMsRdpClient9NotSafeForScripting; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetLeft(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure MsRdpClient9NotSafeForScripting_SetLeft(AObj: TMsRdpClient9NotSafeForScripting; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetTop(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure MsRdpClient9NotSafeForScripting_SetTop(AObj: TMsRdpClient9NotSafeForScripting; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetWidth(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure MsRdpClient9NotSafeForScripting_SetWidth(AObj: TMsRdpClient9NotSafeForScripting; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetHeight(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure MsRdpClient9NotSafeForScripting_SetHeight(AObj: TMsRdpClient9NotSafeForScripting; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetCursor(AObj: TMsRdpClient9NotSafeForScripting): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure MsRdpClient9NotSafeForScripting_SetCursor(AObj: TMsRdpClient9NotSafeForScripting; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetHint(AObj: TMsRdpClient9NotSafeForScripting): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure MsRdpClient9NotSafeForScripting_SetHint(AObj: TMsRdpClient9NotSafeForScripting; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetMargins(AObj: TMsRdpClient9NotSafeForScripting): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure MsRdpClient9NotSafeForScripting_SetMargins(AObj: TMsRdpClient9NotSafeForScripting; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetCustomHint(AObj: TMsRdpClient9NotSafeForScripting): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure MsRdpClient9NotSafeForScripting_SetCustomHint(AObj: TMsRdpClient9NotSafeForScripting; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetComponentCount(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function MsRdpClient9NotSafeForScripting_GetComponentIndex(AObj: TMsRdpClient9NotSafeForScripting): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure MsRdpClient9NotSafeForScripting_SetComponentIndex(AObj: TMsRdpClient9NotSafeForScripting; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetOwner(AObj: TMsRdpClient9NotSafeForScripting): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function MsRdpClient9NotSafeForScripting_GetName(AObj: TMsRdpClient9NotSafeForScripting): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure MsRdpClient9NotSafeForScripting_SetName(AObj: TMsRdpClient9NotSafeForScripting; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetTag(AObj: TMsRdpClient9NotSafeForScripting): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure MsRdpClient9NotSafeForScripting_SetTag(AObj: TMsRdpClient9NotSafeForScripting; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function MsRdpClient9NotSafeForScripting_GetDockClients(AObj: TMsRdpClient9NotSafeForScripting; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function MsRdpClient9NotSafeForScripting_GetControls(AObj: TMsRdpClient9NotSafeForScripting; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function MsRdpClient9NotSafeForScripting_GetComponents(AObj: TMsRdpClient9NotSafeForScripting; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function MsRdpClient9NotSafeForScripting_StaticClassType: TClass; stdcall;
begin
  Result := TMsRdpClient9NotSafeForScripting;
end;

exports
  MsRdpClient9NotSafeForScripting_Create,
  MsRdpClient9NotSafeForScripting_Free,
  MsRdpClient9NotSafeForScripting_SetBounds,
  MsRdpClient9NotSafeForScripting_CanFocus,
  MsRdpClient9NotSafeForScripting_ContainsControl,
  MsRdpClient9NotSafeForScripting_ControlAtPos,
  MsRdpClient9NotSafeForScripting_DisableAlign,
  MsRdpClient9NotSafeForScripting_EnableAlign,
  MsRdpClient9NotSafeForScripting_FindChildControl,
  MsRdpClient9NotSafeForScripting_FlipChildren,
  MsRdpClient9NotSafeForScripting_Focused,
  MsRdpClient9NotSafeForScripting_HandleAllocated,
  MsRdpClient9NotSafeForScripting_InsertControl,
  MsRdpClient9NotSafeForScripting_Invalidate,
  MsRdpClient9NotSafeForScripting_PaintTo,
  MsRdpClient9NotSafeForScripting_RemoveControl,
  MsRdpClient9NotSafeForScripting_Realign,
  MsRdpClient9NotSafeForScripting_Repaint,
  MsRdpClient9NotSafeForScripting_ScaleBy,
  MsRdpClient9NotSafeForScripting_ScrollBy,
  MsRdpClient9NotSafeForScripting_SetFocus,
  MsRdpClient9NotSafeForScripting_Update,
  MsRdpClient9NotSafeForScripting_UpdateControlState,
  MsRdpClient9NotSafeForScripting_BringToFront,
  MsRdpClient9NotSafeForScripting_ClientToScreen,
  MsRdpClient9NotSafeForScripting_ClientToParent,
  MsRdpClient9NotSafeForScripting_Dragging,
  MsRdpClient9NotSafeForScripting_HasParent,
  MsRdpClient9NotSafeForScripting_Hide,
  MsRdpClient9NotSafeForScripting_Perform,
  MsRdpClient9NotSafeForScripting_Refresh,
  MsRdpClient9NotSafeForScripting_ScreenToClient,
  MsRdpClient9NotSafeForScripting_ParentToClient,
  MsRdpClient9NotSafeForScripting_SendToBack,
  MsRdpClient9NotSafeForScripting_Show,
  MsRdpClient9NotSafeForScripting_GetTextBuf,
  MsRdpClient9NotSafeForScripting_GetTextLen,
  MsRdpClient9NotSafeForScripting_SetTextBuf,
  MsRdpClient9NotSafeForScripting_FindComponent,
  MsRdpClient9NotSafeForScripting_GetNamePath,
  MsRdpClient9NotSafeForScripting_Assign,
  MsRdpClient9NotSafeForScripting_DisposeOf,
  MsRdpClient9NotSafeForScripting_ClassType,
  MsRdpClient9NotSafeForScripting_ClassName,
  MsRdpClient9NotSafeForScripting_InstanceSize,
  MsRdpClient9NotSafeForScripting_InheritsFrom,
  MsRdpClient9NotSafeForScripting_Equals,
  MsRdpClient9NotSafeForScripting_GetHashCode,
  MsRdpClient9NotSafeForScripting_ToString,
  MsRdpClient9NotSafeForScripting_GetControlInterface,
  MsRdpClient9NotSafeForScripting_GetDefaultInterface,
  MsRdpClient9NotSafeForScripting_GetVersion,
  MsRdpClient9NotSafeForScripting_GetAdvancedSettings9,
  MsRdpClient9NotSafeForScripting_GetAnchors,
  MsRdpClient9NotSafeForScripting_SetAnchors,
  MsRdpClient9NotSafeForScripting_GetTabStop,
  MsRdpClient9NotSafeForScripting_SetTabStop,
  MsRdpClient9NotSafeForScripting_GetAlign,
  MsRdpClient9NotSafeForScripting_SetAlign,
  MsRdpClient9NotSafeForScripting_GetDragCursor,
  MsRdpClient9NotSafeForScripting_SetDragCursor,
  MsRdpClient9NotSafeForScripting_GetDragMode,
  MsRdpClient9NotSafeForScripting_SetDragMode,
  MsRdpClient9NotSafeForScripting_GetParentShowHint,
  MsRdpClient9NotSafeForScripting_SetParentShowHint,
  MsRdpClient9NotSafeForScripting_GetPopupMenu,
  MsRdpClient9NotSafeForScripting_SetPopupMenu,
  MsRdpClient9NotSafeForScripting_GetShowHint,
  MsRdpClient9NotSafeForScripting_SetShowHint,
  MsRdpClient9NotSafeForScripting_GetTabOrder,
  MsRdpClient9NotSafeForScripting_SetTabOrder,
  MsRdpClient9NotSafeForScripting_GetVisible,
  MsRdpClient9NotSafeForScripting_SetVisible,
  MsRdpClient9NotSafeForScripting_SetOnDragDrop,
  MsRdpClient9NotSafeForScripting_SetOnDragOver,
  MsRdpClient9NotSafeForScripting_SetOnEndDrag,
  MsRdpClient9NotSafeForScripting_SetOnEnter,
  MsRdpClient9NotSafeForScripting_SetOnExit,
  MsRdpClient9NotSafeForScripting_GetDesktopWidth,
  MsRdpClient9NotSafeForScripting_SetDesktopWidth,
  MsRdpClient9NotSafeForScripting_GetDesktopHeight,
  MsRdpClient9NotSafeForScripting_SetDesktopHeight,
  MsRdpClient9NotSafeForScripting_GetColorDepth,
  MsRdpClient9NotSafeForScripting_SetColorDepth,
  MsRdpClient9NotSafeForScripting_GetFullScreen,
  MsRdpClient9NotSafeForScripting_SetFullScreen,
  MsRdpClient9NotSafeForScripting_SetOnDisconnected,
  MsRdpClient9NotSafeForScripting_GetDockClientCount,
  MsRdpClient9NotSafeForScripting_GetDockSite,
  MsRdpClient9NotSafeForScripting_SetDockSite,
  MsRdpClient9NotSafeForScripting_GetDoubleBuffered,
  MsRdpClient9NotSafeForScripting_SetDoubleBuffered,
  MsRdpClient9NotSafeForScripting_GetAlignDisabled,
  MsRdpClient9NotSafeForScripting_GetMouseInClient,
  MsRdpClient9NotSafeForScripting_GetVisibleDockClientCount,
  MsRdpClient9NotSafeForScripting_GetBrush,
  MsRdpClient9NotSafeForScripting_GetControlCount,
  MsRdpClient9NotSafeForScripting_GetHandle,
  MsRdpClient9NotSafeForScripting_GetParentDoubleBuffered,
  MsRdpClient9NotSafeForScripting_SetParentDoubleBuffered,
  MsRdpClient9NotSafeForScripting_GetParentWindow,
  MsRdpClient9NotSafeForScripting_SetParentWindow,
  MsRdpClient9NotSafeForScripting_GetUseDockManager,
  MsRdpClient9NotSafeForScripting_SetUseDockManager,
  MsRdpClient9NotSafeForScripting_GetEnabled,
  MsRdpClient9NotSafeForScripting_SetEnabled,
  MsRdpClient9NotSafeForScripting_GetAction,
  MsRdpClient9NotSafeForScripting_SetAction,
  MsRdpClient9NotSafeForScripting_GetBiDiMode,
  MsRdpClient9NotSafeForScripting_SetBiDiMode,
  MsRdpClient9NotSafeForScripting_GetBoundsRect,
  MsRdpClient9NotSafeForScripting_SetBoundsRect,
  MsRdpClient9NotSafeForScripting_GetClientHeight,
  MsRdpClient9NotSafeForScripting_SetClientHeight,
  MsRdpClient9NotSafeForScripting_GetClientOrigin,
  MsRdpClient9NotSafeForScripting_GetClientRect,
  MsRdpClient9NotSafeForScripting_GetClientWidth,
  MsRdpClient9NotSafeForScripting_SetClientWidth,
  MsRdpClient9NotSafeForScripting_GetControlState,
  MsRdpClient9NotSafeForScripting_SetControlState,
  MsRdpClient9NotSafeForScripting_GetControlStyle,
  MsRdpClient9NotSafeForScripting_SetControlStyle,
  MsRdpClient9NotSafeForScripting_GetExplicitLeft,
  MsRdpClient9NotSafeForScripting_GetExplicitTop,
  MsRdpClient9NotSafeForScripting_GetExplicitWidth,
  MsRdpClient9NotSafeForScripting_GetExplicitHeight,
  MsRdpClient9NotSafeForScripting_GetFloating,
  MsRdpClient9NotSafeForScripting_GetParent,
  MsRdpClient9NotSafeForScripting_SetParent,
  MsRdpClient9NotSafeForScripting_GetStyleElements,
  MsRdpClient9NotSafeForScripting_SetStyleElements,
  MsRdpClient9NotSafeForScripting_SetOnGesture,
  MsRdpClient9NotSafeForScripting_GetAlignWithMargins,
  MsRdpClient9NotSafeForScripting_SetAlignWithMargins,
  MsRdpClient9NotSafeForScripting_GetLeft,
  MsRdpClient9NotSafeForScripting_SetLeft,
  MsRdpClient9NotSafeForScripting_GetTop,
  MsRdpClient9NotSafeForScripting_SetTop,
  MsRdpClient9NotSafeForScripting_GetWidth,
  MsRdpClient9NotSafeForScripting_SetWidth,
  MsRdpClient9NotSafeForScripting_GetHeight,
  MsRdpClient9NotSafeForScripting_SetHeight,
  MsRdpClient9NotSafeForScripting_GetCursor,
  MsRdpClient9NotSafeForScripting_SetCursor,
  MsRdpClient9NotSafeForScripting_GetHint,
  MsRdpClient9NotSafeForScripting_SetHint,
  MsRdpClient9NotSafeForScripting_GetMargins,
  MsRdpClient9NotSafeForScripting_SetMargins,
  MsRdpClient9NotSafeForScripting_GetCustomHint,
  MsRdpClient9NotSafeForScripting_SetCustomHint,
  MsRdpClient9NotSafeForScripting_GetComponentCount,
  MsRdpClient9NotSafeForScripting_GetComponentIndex,
  MsRdpClient9NotSafeForScripting_SetComponentIndex,
  MsRdpClient9NotSafeForScripting_GetOwner,
  MsRdpClient9NotSafeForScripting_GetName,
  MsRdpClient9NotSafeForScripting_SetName,
  MsRdpClient9NotSafeForScripting_GetTag,
  MsRdpClient9NotSafeForScripting_SetTag,
  MsRdpClient9NotSafeForScripting_GetDockClients,
  MsRdpClient9NotSafeForScripting_GetControls,
  MsRdpClient9NotSafeForScripting_GetComponents,
  MsRdpClient9NotSafeForScripting_StaticClassType;

