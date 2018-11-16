
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ScrollBox_Create(AOwner: TComponent): TScrollBox; stdcall;
begin
  Result :=  TScrollBox.Create(AOwner);
end;

procedure ScrollBox_Free(AObj: TScrollBox); stdcall;
begin
  AObj.Free;
end;

function ScrollBox_CanFocus(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function ScrollBox_ContainsControl(AObj: TScrollBox; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function ScrollBox_ControlAtPos(AObj: TScrollBox; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure ScrollBox_DisableAlign(AObj: TScrollBox); stdcall;
begin
  AObj.DisableAlign;
end;

procedure ScrollBox_EnableAlign(AObj: TScrollBox); stdcall;
begin
  AObj.EnableAlign;
end;

function ScrollBox_FindChildControl(AObj: TScrollBox; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

procedure ScrollBox_FlipChildren(AObj: TScrollBox; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function ScrollBox_Focused(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function ScrollBox_HandleAllocated(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ScrollBox_InsertControl(AObj: TScrollBox; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure ScrollBox_Invalidate(AObj: TScrollBox); stdcall;
begin
  AObj.Invalidate;
end;

procedure ScrollBox_PaintTo(AObj: TScrollBox; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure ScrollBox_RemoveControl(AObj: TScrollBox; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure ScrollBox_Realign(AObj: TScrollBox); stdcall;
begin
  AObj.Realign;
end;

procedure ScrollBox_Repaint(AObj: TScrollBox); stdcall;
begin
  AObj.Repaint;
end;

procedure ScrollBox_ScaleBy(AObj: TScrollBox; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure ScrollBox_ScrollBy(AObj: TScrollBox; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure ScrollBox_SetBounds(AObj: TScrollBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ScrollBox_SetFocus(AObj: TScrollBox); stdcall;
begin
  AObj.SetFocus;
end;

procedure ScrollBox_Update(AObj: TScrollBox); stdcall;
begin
  AObj.Update;
end;

procedure ScrollBox_UpdateControlState(AObj: TScrollBox); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure ScrollBox_BringToFront(AObj: TScrollBox); stdcall;
begin
  AObj.BringToFront;
end;

procedure ScrollBox_ClientToScreen(AObj: TScrollBox; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure ScrollBox_ClientToParent(AObj: TScrollBox; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function ScrollBox_Dragging(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function ScrollBox_HasParent(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ScrollBox_Hide(AObj: TScrollBox); stdcall;
begin
  AObj.Hide;
end;

function ScrollBox_Perform(AObj: TScrollBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ScrollBox_Refresh(AObj: TScrollBox); stdcall;
begin
  AObj.Refresh;
end;

procedure ScrollBox_ScreenToClient(AObj: TScrollBox; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure ScrollBox_ParentToClient(AObj: TScrollBox; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure ScrollBox_SendToBack(AObj: TScrollBox); stdcall;
begin
  AObj.SendToBack;
end;

procedure ScrollBox_Show(AObj: TScrollBox); stdcall;
begin
  AObj.Show;
end;

function ScrollBox_GetTextBuf(AObj: TScrollBox; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ScrollBox_GetTextLen(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure ScrollBox_SetTextBuf(AObj: TScrollBox; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function ScrollBox_FindComponent(AObj: TScrollBox; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ScrollBox_GetNamePath(AObj: TScrollBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ScrollBox_Assign(AObj: TScrollBox; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure ScrollBox_DisposeOf(AObj: TScrollBox); stdcall;
begin
  AObj.DisposeOf;
end;

function ScrollBox_ClassType(AObj: TScrollBox): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function ScrollBox_ClassName(AObj: TScrollBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ScrollBox_InstanceSize(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function ScrollBox_InheritsFrom(AObj: TScrollBox; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function ScrollBox_Equals(AObj: TScrollBox; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ScrollBox_GetHashCode(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ScrollBox_ToString(AObj: TScrollBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ScrollBox_GetAlign(AObj: TScrollBox): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ScrollBox_SetAlign(AObj: TScrollBox; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ScrollBox_GetAnchors(AObj: TScrollBox): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ScrollBox_SetAnchors(AObj: TScrollBox; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ScrollBox_GetAutoSize(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure ScrollBox_SetAutoSize(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function ScrollBox_GetBevelEdges(AObj: TScrollBox): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure ScrollBox_SetBevelEdges(AObj: TScrollBox; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function ScrollBox_GetBevelInner(AObj: TScrollBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure ScrollBox_SetBevelInner(AObj: TScrollBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function ScrollBox_GetBevelOuter(AObj: TScrollBox): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure ScrollBox_SetBevelOuter(AObj: TScrollBox; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function ScrollBox_GetBevelKind(AObj: TScrollBox): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure ScrollBox_SetBevelKind(AObj: TScrollBox; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function ScrollBox_GetBiDiMode(AObj: TScrollBox): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ScrollBox_SetBiDiMode(AObj: TScrollBox; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function ScrollBox_GetBorderStyle(AObj: TScrollBox): TBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure ScrollBox_SetBorderStyle(AObj: TScrollBox; AValue: TBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function ScrollBox_GetDockSite(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure ScrollBox_SetDockSite(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function ScrollBox_GetDoubleBuffered(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ScrollBox_SetDoubleBuffered(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function ScrollBox_GetDragCursor(AObj: TScrollBox): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure ScrollBox_SetDragCursor(AObj: TScrollBox; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function ScrollBox_GetDragKind(AObj: TScrollBox): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure ScrollBox_SetDragKind(AObj: TScrollBox; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function ScrollBox_GetDragMode(AObj: TScrollBox): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure ScrollBox_SetDragMode(AObj: TScrollBox; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function ScrollBox_GetEnabled(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ScrollBox_SetEnabled(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ScrollBox_GetColor(AObj: TScrollBox): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure ScrollBox_SetColor(AObj: TScrollBox; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function ScrollBox_GetFont(AObj: TScrollBox): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure ScrollBox_SetFont(AObj: TScrollBox; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function ScrollBox_GetParentBackground(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentBackground;
end;

procedure ScrollBox_SetParentBackground(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.ParentBackground := AValue;
end;

function ScrollBox_GetParentColor(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure ScrollBox_SetParentColor(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function ScrollBox_GetParentCtl3D(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure ScrollBox_SetParentCtl3D(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function ScrollBox_GetParentDoubleBuffered(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure ScrollBox_SetParentDoubleBuffered(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function ScrollBox_GetParentFont(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure ScrollBox_SetParentFont(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function ScrollBox_GetParentShowHint(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ScrollBox_SetParentShowHint(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ScrollBox_GetPopupMenu(AObj: TScrollBox): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ScrollBox_SetPopupMenu(AObj: TScrollBox; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ScrollBox_GetShowHint(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ScrollBox_SetShowHint(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ScrollBox_GetTabOrder(AObj: TScrollBox): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure ScrollBox_SetTabOrder(AObj: TScrollBox; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function ScrollBox_GetTabStop(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure ScrollBox_SetTabStop(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function ScrollBox_GetVisible(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ScrollBox_SetVisible(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function ScrollBox_GetStyleElements(AObj: TScrollBox): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ScrollBox_SetStyleElements(AObj: TScrollBox; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure ScrollBox_SetOnClick(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    TEventClass.Remove(AObj, @TEventClass.OnClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, @TEventClass.OnClick, AEventId);
end;

procedure ScrollBox_SetOnContextPopup(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnDblClick(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDblClick := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDblClick);
    Exit;
  end;
  AObj.OnDblClick := TEventClass.OnDblClick;
  TEventClass.Add(AObj, @TEventClass.OnDblClick, AEventId);
end;

procedure ScrollBox_SetOnDockDrop(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDockDrop := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDockDrop);
    Exit;
  end;
  AObj.OnDockDrop := TEventClass.OnDockDrop;
  TEventClass.Add(AObj, @TEventClass.OnDockDrop, AEventId);
end;

procedure ScrollBox_SetOnDragDrop(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnDragOver(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnEndDock(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDock := nil;
    TEventClass.Remove(AObj, @TEventClass.OnEndDock);
    Exit;
  end;
  AObj.OnEndDock := TEventClass.OnEndDock;
  TEventClass.Add(AObj, @TEventClass.OnEndDock, AEventId);
end;

procedure ScrollBox_SetOnEndDrag(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnEnter(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnExit(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnGesture(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnGetSiteInfo(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetSiteInfo := nil;
    TEventClass.Remove(AObj, @TEventClass.OnGetSiteInfo);
    Exit;
  end;
  AObj.OnGetSiteInfo := TEventClass.OnGetSiteInfo;
  TEventClass.Add(AObj, @TEventClass.OnGetSiteInfo, AEventId);
end;

procedure ScrollBox_SetOnMouseActivate(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseDown(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseEnter(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseLeave(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseMove(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseUp(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseWheel(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheel := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseWheel);
    Exit;
  end;
  AObj.OnMouseWheel := TEventClass.OnMouseWheel;
  TEventClass.Add(AObj, @TEventClass.OnMouseWheel, AEventId);
end;

procedure ScrollBox_SetOnMouseWheelDown(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheelDown := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseWheelDown);
    Exit;
  end;
  AObj.OnMouseWheelDown := TEventClass.OnMouseWheelDown;
  TEventClass.Add(AObj, @TEventClass.OnMouseWheelDown, AEventId);
end;

procedure ScrollBox_SetOnMouseWheelUp(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheelUp := nil;
    TEventClass.Remove(AObj, @TEventClass.OnMouseWheelUp);
    Exit;
  end;
  AObj.OnMouseWheelUp := TEventClass.OnMouseWheelUp;
  TEventClass.Add(AObj, @TEventClass.OnMouseWheelUp, AEventId);
end;

procedure ScrollBox_SetOnResize(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnResize := nil;
    TEventClass.Remove(AObj, @TEventClass.OnResize);
    Exit;
  end;
  AObj.OnResize := TEventClass.OnResize;
  TEventClass.Add(AObj, @TEventClass.OnResize, AEventId);
end;

procedure ScrollBox_SetOnStartDock(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnStartDock := nil;
    TEventClass.Remove(AObj, @TEventClass.OnStartDock);
    Exit;
  end;
  AObj.OnStartDock := TEventClass.OnStartDock;
  TEventClass.Add(AObj, @TEventClass.OnStartDock, AEventId);
end;

procedure ScrollBox_SetOnUnDock(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnUnDock := nil;
    TEventClass.Remove(AObj, @TEventClass.OnUnDock);
    Exit;
  end;
  AObj.OnUnDock := TEventClass.OnUnDock;
  TEventClass.Add(AObj, @TEventClass.OnUnDock, AEventId);
end;

function ScrollBox_GetDockClientCount(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function ScrollBox_GetAlignDisabled(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function ScrollBox_GetMouseInClient(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function ScrollBox_GetVisibleDockClientCount(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function ScrollBox_GetBrush(AObj: TScrollBox): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function ScrollBox_GetControlCount(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function ScrollBox_GetHandle(AObj: TScrollBox): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ScrollBox_GetParentWindow(AObj: TScrollBox): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ScrollBox_SetParentWindow(AObj: TScrollBox; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function ScrollBox_GetUseDockManager(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure ScrollBox_SetUseDockManager(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function ScrollBox_GetAction(AObj: TScrollBox): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ScrollBox_SetAction(AObj: TScrollBox; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure ScrollBox_GetBoundsRect(AObj: TScrollBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ScrollBox_SetBoundsRect(AObj: TScrollBox; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ScrollBox_GetClientHeight(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ScrollBox_SetClientHeight(AObj: TScrollBox; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ScrollBox_GetClientOrigin(AObj: TScrollBox; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure ScrollBox_GetClientRect(AObj: TScrollBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ScrollBox_GetClientWidth(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ScrollBox_SetClientWidth(AObj: TScrollBox; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ScrollBox_GetControlState(AObj: TScrollBox): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure ScrollBox_SetControlState(AObj: TScrollBox; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function ScrollBox_GetControlStyle(AObj: TScrollBox): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure ScrollBox_SetControlStyle(AObj: TScrollBox; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function ScrollBox_GetExplicitLeft(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ScrollBox_GetExplicitTop(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ScrollBox_GetExplicitWidth(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ScrollBox_GetExplicitHeight(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ScrollBox_GetFloating(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function ScrollBox_GetParent(AObj: TScrollBox): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ScrollBox_SetParent(AObj: TScrollBox; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ScrollBox_GetAlignWithMargins(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ScrollBox_SetAlignWithMargins(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ScrollBox_GetLeft(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ScrollBox_SetLeft(AObj: TScrollBox; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ScrollBox_GetTop(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ScrollBox_SetTop(AObj: TScrollBox; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ScrollBox_GetWidth(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ScrollBox_SetWidth(AObj: TScrollBox; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ScrollBox_GetHeight(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ScrollBox_SetHeight(AObj: TScrollBox; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ScrollBox_GetCursor(AObj: TScrollBox): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ScrollBox_SetCursor(AObj: TScrollBox; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ScrollBox_GetHint(AObj: TScrollBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ScrollBox_SetHint(AObj: TScrollBox; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ScrollBox_GetMargins(AObj: TScrollBox): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ScrollBox_SetMargins(AObj: TScrollBox; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ScrollBox_GetCustomHint(AObj: TScrollBox): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ScrollBox_SetCustomHint(AObj: TScrollBox; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ScrollBox_GetComponentCount(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ScrollBox_GetComponentIndex(AObj: TScrollBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ScrollBox_SetComponentIndex(AObj: TScrollBox; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ScrollBox_GetOwner(AObj: TScrollBox): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ScrollBox_GetName(AObj: TScrollBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ScrollBox_SetName(AObj: TScrollBox; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ScrollBox_GetTag(AObj: TScrollBox): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ScrollBox_SetTag(AObj: TScrollBox; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ScrollBox_GetDockClients(AObj: TScrollBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function ScrollBox_GetControls(AObj: TScrollBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ScrollBox_GetComponents(AObj: TScrollBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function ScrollBox_StaticClassType: TClass; stdcall;
begin
  Result := TScrollBox;
end;

exports
  ScrollBox_Create,
  ScrollBox_Free,
  ScrollBox_CanFocus,
  ScrollBox_ContainsControl,
  ScrollBox_ControlAtPos,
  ScrollBox_DisableAlign,
  ScrollBox_EnableAlign,
  ScrollBox_FindChildControl,
  ScrollBox_FlipChildren,
  ScrollBox_Focused,
  ScrollBox_HandleAllocated,
  ScrollBox_InsertControl,
  ScrollBox_Invalidate,
  ScrollBox_PaintTo,
  ScrollBox_RemoveControl,
  ScrollBox_Realign,
  ScrollBox_Repaint,
  ScrollBox_ScaleBy,
  ScrollBox_ScrollBy,
  ScrollBox_SetBounds,
  ScrollBox_SetFocus,
  ScrollBox_Update,
  ScrollBox_UpdateControlState,
  ScrollBox_BringToFront,
  ScrollBox_ClientToScreen,
  ScrollBox_ClientToParent,
  ScrollBox_Dragging,
  ScrollBox_HasParent,
  ScrollBox_Hide,
  ScrollBox_Perform,
  ScrollBox_Refresh,
  ScrollBox_ScreenToClient,
  ScrollBox_ParentToClient,
  ScrollBox_SendToBack,
  ScrollBox_Show,
  ScrollBox_GetTextBuf,
  ScrollBox_GetTextLen,
  ScrollBox_SetTextBuf,
  ScrollBox_FindComponent,
  ScrollBox_GetNamePath,
  ScrollBox_Assign,
  ScrollBox_DisposeOf,
  ScrollBox_ClassType,
  ScrollBox_ClassName,
  ScrollBox_InstanceSize,
  ScrollBox_InheritsFrom,
  ScrollBox_Equals,
  ScrollBox_GetHashCode,
  ScrollBox_ToString,
  ScrollBox_GetAlign,
  ScrollBox_SetAlign,
  ScrollBox_GetAnchors,
  ScrollBox_SetAnchors,
  ScrollBox_GetAutoSize,
  ScrollBox_SetAutoSize,
  ScrollBox_GetBevelEdges,
  ScrollBox_SetBevelEdges,
  ScrollBox_GetBevelInner,
  ScrollBox_SetBevelInner,
  ScrollBox_GetBevelOuter,
  ScrollBox_SetBevelOuter,
  ScrollBox_GetBevelKind,
  ScrollBox_SetBevelKind,
  ScrollBox_GetBiDiMode,
  ScrollBox_SetBiDiMode,
  ScrollBox_GetBorderStyle,
  ScrollBox_SetBorderStyle,
  ScrollBox_GetDockSite,
  ScrollBox_SetDockSite,
  ScrollBox_GetDoubleBuffered,
  ScrollBox_SetDoubleBuffered,
  ScrollBox_GetDragCursor,
  ScrollBox_SetDragCursor,
  ScrollBox_GetDragKind,
  ScrollBox_SetDragKind,
  ScrollBox_GetDragMode,
  ScrollBox_SetDragMode,
  ScrollBox_GetEnabled,
  ScrollBox_SetEnabled,
  ScrollBox_GetColor,
  ScrollBox_SetColor,
  ScrollBox_GetFont,
  ScrollBox_SetFont,
  ScrollBox_GetParentBackground,
  ScrollBox_SetParentBackground,
  ScrollBox_GetParentColor,
  ScrollBox_SetParentColor,
  ScrollBox_GetParentCtl3D,
  ScrollBox_SetParentCtl3D,
  ScrollBox_GetParentDoubleBuffered,
  ScrollBox_SetParentDoubleBuffered,
  ScrollBox_GetParentFont,
  ScrollBox_SetParentFont,
  ScrollBox_GetParentShowHint,
  ScrollBox_SetParentShowHint,
  ScrollBox_GetPopupMenu,
  ScrollBox_SetPopupMenu,
  ScrollBox_GetShowHint,
  ScrollBox_SetShowHint,
  ScrollBox_GetTabOrder,
  ScrollBox_SetTabOrder,
  ScrollBox_GetTabStop,
  ScrollBox_SetTabStop,
  ScrollBox_GetVisible,
  ScrollBox_SetVisible,
  ScrollBox_GetStyleElements,
  ScrollBox_SetStyleElements,
  ScrollBox_SetOnClick,
  ScrollBox_SetOnContextPopup,
  ScrollBox_SetOnDblClick,
  ScrollBox_SetOnDockDrop,
  ScrollBox_SetOnDragDrop,
  ScrollBox_SetOnDragOver,
  ScrollBox_SetOnEndDock,
  ScrollBox_SetOnEndDrag,
  ScrollBox_SetOnEnter,
  ScrollBox_SetOnExit,
  ScrollBox_SetOnGesture,
  ScrollBox_SetOnGetSiteInfo,
  ScrollBox_SetOnMouseActivate,
  ScrollBox_SetOnMouseDown,
  ScrollBox_SetOnMouseEnter,
  ScrollBox_SetOnMouseLeave,
  ScrollBox_SetOnMouseMove,
  ScrollBox_SetOnMouseUp,
  ScrollBox_SetOnMouseWheel,
  ScrollBox_SetOnMouseWheelDown,
  ScrollBox_SetOnMouseWheelUp,
  ScrollBox_SetOnResize,
  ScrollBox_SetOnStartDock,
  ScrollBox_SetOnUnDock,
  ScrollBox_GetDockClientCount,
  ScrollBox_GetAlignDisabled,
  ScrollBox_GetMouseInClient,
  ScrollBox_GetVisibleDockClientCount,
  ScrollBox_GetBrush,
  ScrollBox_GetControlCount,
  ScrollBox_GetHandle,
  ScrollBox_GetParentWindow,
  ScrollBox_SetParentWindow,
  ScrollBox_GetUseDockManager,
  ScrollBox_SetUseDockManager,
  ScrollBox_GetAction,
  ScrollBox_SetAction,
  ScrollBox_GetBoundsRect,
  ScrollBox_SetBoundsRect,
  ScrollBox_GetClientHeight,
  ScrollBox_SetClientHeight,
  ScrollBox_GetClientOrigin,
  ScrollBox_GetClientRect,
  ScrollBox_GetClientWidth,
  ScrollBox_SetClientWidth,
  ScrollBox_GetControlState,
  ScrollBox_SetControlState,
  ScrollBox_GetControlStyle,
  ScrollBox_SetControlStyle,
  ScrollBox_GetExplicitLeft,
  ScrollBox_GetExplicitTop,
  ScrollBox_GetExplicitWidth,
  ScrollBox_GetExplicitHeight,
  ScrollBox_GetFloating,
  ScrollBox_GetParent,
  ScrollBox_SetParent,
  ScrollBox_GetAlignWithMargins,
  ScrollBox_SetAlignWithMargins,
  ScrollBox_GetLeft,
  ScrollBox_SetLeft,
  ScrollBox_GetTop,
  ScrollBox_SetTop,
  ScrollBox_GetWidth,
  ScrollBox_SetWidth,
  ScrollBox_GetHeight,
  ScrollBox_SetHeight,
  ScrollBox_GetCursor,
  ScrollBox_SetCursor,
  ScrollBox_GetHint,
  ScrollBox_SetHint,
  ScrollBox_GetMargins,
  ScrollBox_SetMargins,
  ScrollBox_GetCustomHint,
  ScrollBox_SetCustomHint,
  ScrollBox_GetComponentCount,
  ScrollBox_GetComponentIndex,
  ScrollBox_SetComponentIndex,
  ScrollBox_GetOwner,
  ScrollBox_GetName,
  ScrollBox_SetName,
  ScrollBox_GetTag,
  ScrollBox_SetTag,
  ScrollBox_GetDockClients,
  ScrollBox_GetControls,
  ScrollBox_GetComponents,
  ScrollBox_StaticClassType;

