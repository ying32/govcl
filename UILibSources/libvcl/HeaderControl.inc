
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function HeaderControl_Create(AOwner: TComponent): THeaderControl; stdcall;
begin
  Result :=  THeaderControl.Create(AOwner);
end;

procedure HeaderControl_Free(AObj: THeaderControl); stdcall;
begin
  AObj.Free;
end;

procedure HeaderControl_FlipChildren(AObj: THeaderControl; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function HeaderControl_CanFocus(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function HeaderControl_ContainsControl(AObj: THeaderControl; Control: TControl): LongBool; stdcall;
begin
  Result :=  AObj.ContainsControl(Control);
end;

function HeaderControl_ControlAtPos(AObj: THeaderControl; var Pos: TPoint; AllowDisabled: LongBool; AllowWinControls: LongBool; AllLevels: LongBool): TControl; stdcall;
begin
  Result :=  AObj.ControlAtPos(Pos, AllowDisabled, AllowWinControls, AllLevels);
end;

procedure HeaderControl_DisableAlign(AObj: THeaderControl); stdcall;
begin
  AObj.DisableAlign;
end;

procedure HeaderControl_EnableAlign(AObj: THeaderControl); stdcall;
begin
  AObj.EnableAlign;
end;

function HeaderControl_FindChildControl(AObj: THeaderControl; ControlName: PChar): TControl; stdcall;
begin
  Result :=  AObj.FindChildControl(ControlName);
end;

function HeaderControl_Focused(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function HeaderControl_HandleAllocated(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure HeaderControl_InsertControl(AObj: THeaderControl; AControl: TControl); stdcall;
begin
  AObj.InsertControl(AControl);
end;

procedure HeaderControl_Invalidate(AObj: THeaderControl); stdcall;
begin
  AObj.Invalidate;
end;

procedure HeaderControl_PaintTo(AObj: THeaderControl; DC: HDC; X: Integer; Y: Integer); stdcall;
begin
  AObj.PaintTo(DC, X, Y);
end;

procedure HeaderControl_RemoveControl(AObj: THeaderControl; AControl: TControl); stdcall;
begin
  AObj.RemoveControl(AControl);
end;

procedure HeaderControl_Realign(AObj: THeaderControl); stdcall;
begin
  AObj.Realign;
end;

procedure HeaderControl_Repaint(AObj: THeaderControl); stdcall;
begin
  AObj.Repaint;
end;

procedure HeaderControl_ScaleBy(AObj: THeaderControl; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure HeaderControl_ScrollBy(AObj: THeaderControl; DeltaX: Integer; DeltaY: Integer); stdcall;
begin
  AObj.ScrollBy(DeltaX, DeltaY);
end;

procedure HeaderControl_SetBounds(AObj: THeaderControl; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure HeaderControl_SetFocus(AObj: THeaderControl); stdcall;
begin
  AObj.SetFocus;
end;

procedure HeaderControl_Update(AObj: THeaderControl); stdcall;
begin
  AObj.Update;
end;

procedure HeaderControl_UpdateControlState(AObj: THeaderControl); stdcall;
begin
  AObj.UpdateControlState;
end;

procedure HeaderControl_BringToFront(AObj: THeaderControl); stdcall;
begin
  AObj.BringToFront;
end;

procedure HeaderControl_ClientToScreen(AObj: THeaderControl; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure HeaderControl_ClientToParent(AObj: THeaderControl; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function HeaderControl_Dragging(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function HeaderControl_HasParent(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure HeaderControl_Hide(AObj: THeaderControl); stdcall;
begin
  AObj.Hide;
end;

function HeaderControl_Perform(AObj: THeaderControl; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure HeaderControl_Refresh(AObj: THeaderControl); stdcall;
begin
  AObj.Refresh;
end;

procedure HeaderControl_ScreenToClient(AObj: THeaderControl; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure HeaderControl_ParentToClient(AObj: THeaderControl; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure HeaderControl_SendToBack(AObj: THeaderControl); stdcall;
begin
  AObj.SendToBack;
end;

procedure HeaderControl_Show(AObj: THeaderControl); stdcall;
begin
  AObj.Show;
end;

function HeaderControl_GetTextBuf(AObj: THeaderControl; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function HeaderControl_GetTextLen(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure HeaderControl_SetTextBuf(AObj: THeaderControl; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function HeaderControl_FindComponent(AObj: THeaderControl; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function HeaderControl_GetNamePath(AObj: THeaderControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure HeaderControl_Assign(AObj: THeaderControl; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure HeaderControl_DisposeOf(AObj: THeaderControl); stdcall;
begin
  AObj.DisposeOf;
end;

function HeaderControl_ClassType(AObj: THeaderControl): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function HeaderControl_ClassName(AObj: THeaderControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function HeaderControl_InstanceSize(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function HeaderControl_InheritsFrom(AObj: THeaderControl; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function HeaderControl_Equals(AObj: THeaderControl; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function HeaderControl_GetHashCode(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function HeaderControl_ToString(AObj: THeaderControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function HeaderControl_GetAlign(AObj: THeaderControl): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure HeaderControl_SetAlign(AObj: THeaderControl; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function HeaderControl_GetAnchors(AObj: THeaderControl): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure HeaderControl_SetAnchors(AObj: THeaderControl; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function HeaderControl_GetBiDiMode(AObj: THeaderControl): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure HeaderControl_SetBiDiMode(AObj: THeaderControl; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function HeaderControl_GetBorderWidth(AObj: THeaderControl): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure HeaderControl_SetBorderWidth(AObj: THeaderControl; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function HeaderControl_GetDoubleBuffered(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure HeaderControl_SetDoubleBuffered(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function HeaderControl_GetDragCursor(AObj: THeaderControl): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure HeaderControl_SetDragCursor(AObj: THeaderControl; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function HeaderControl_GetDragKind(AObj: THeaderControl): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure HeaderControl_SetDragKind(AObj: THeaderControl; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function HeaderControl_GetDragMode(AObj: THeaderControl): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure HeaderControl_SetDragMode(AObj: THeaderControl; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function HeaderControl_GetEnabled(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure HeaderControl_SetEnabled(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function HeaderControl_GetFont(AObj: THeaderControl): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure HeaderControl_SetFont(AObj: THeaderControl; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function HeaderControl_GetFullDrag(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.FullDrag;
end;

procedure HeaderControl_SetFullDrag(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.FullDrag := AValue;
end;

function HeaderControl_GetHotTrack(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.HotTrack;
end;

procedure HeaderControl_SetHotTrack(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.HotTrack := AValue;
end;

function HeaderControl_GetImages(AObj: THeaderControl): TCustomImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure HeaderControl_SetImages(AObj: THeaderControl; AValue: TCustomImageList); stdcall;
begin
  AObj.Images := AValue;
end;

function HeaderControl_GetShowHint(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure HeaderControl_SetShowHint(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function HeaderControl_GetStyle(AObj: THeaderControl): THeaderStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure HeaderControl_SetStyle(AObj: THeaderControl; AValue: THeaderStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function HeaderControl_GetParentDoubleBuffered(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure HeaderControl_SetParentDoubleBuffered(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function HeaderControl_GetParentFont(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure HeaderControl_SetParentFont(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function HeaderControl_GetParentShowHint(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure HeaderControl_SetParentShowHint(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function HeaderControl_GetPopupMenu(AObj: THeaderControl): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure HeaderControl_SetPopupMenu(AObj: THeaderControl; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function HeaderControl_GetVisible(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure HeaderControl_SetVisible(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function HeaderControl_GetStyleElements(AObj: THeaderControl): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure HeaderControl_SetStyleElements(AObj: THeaderControl; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure HeaderControl_SetOnContextPopup(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnDragDrop(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnDragOver(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnEndDock(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnEndDrag(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnGesture(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnMouseActivate(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnMouseDown(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnMouseEnter(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnMouseLeave(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnMouseMove(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnMouseUp(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnResize(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

procedure HeaderControl_SetOnDrawSection(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDrawSection := nil;
    TEventClass.Remove(AObj, @TEventClass.OnDrawSection);
    Exit;
  end;
  AObj.OnDrawSection := TEventClass.OnDrawSection;
  TEventClass.Add(AObj, @TEventClass.OnDrawSection, AEventId);
end;

procedure HeaderControl_SetOnSectionClick(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnSectionClick := nil;
    TEventClass.Remove(AObj, @TEventClass.OnSectionClick);
    Exit;
  end;
  AObj.OnSectionClick := TEventClass.OnSectionClick;
  TEventClass.Add(AObj, @TEventClass.OnSectionClick, AEventId);
end;

procedure HeaderControl_SetOnSectionResize(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnSectionResize := nil;
    TEventClass.Remove(AObj, @TEventClass.OnSectionResize);
    Exit;
  end;
  AObj.OnSectionResize := TEventClass.OnSectionResize;
  TEventClass.Add(AObj, @TEventClass.OnSectionResize, AEventId);
end;

procedure HeaderControl_SetOnSectionTrack(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnSectionTrack := nil;
    TEventClass.Remove(AObj, @TEventClass.OnSectionTrack);
    Exit;
  end;
  AObj.OnSectionTrack := TEventClass.OnSectionTrack;
  TEventClass.Add(AObj, @TEventClass.OnSectionTrack, AEventId);
end;

procedure HeaderControl_SetOnSectionDrag(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnSectionDrag := nil;
    TEventClass.Remove(AObj, @TEventClass.OnSectionDrag);
    Exit;
  end;
  AObj.OnSectionDrag := TEventClass.OnSectionDrag;
  TEventClass.Add(AObj, @TEventClass.OnSectionDrag, AEventId);
end;

procedure HeaderControl_SetOnSectionEndDrag(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnSectionEndDrag := nil;
    TEventClass.Remove(AObj, @TEventClass.OnSectionEndDrag);
    Exit;
  end;
  AObj.OnSectionEndDrag := TEventClass.OnSectionEndDrag;
  TEventClass.Add(AObj, @TEventClass.OnSectionEndDrag, AEventId);
end;

procedure HeaderControl_SetOnStartDock(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
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

function HeaderControl_GetCanvas(AObj: THeaderControl): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

procedure HeaderControl_SetOnSectionCheck(AObj: THeaderControl; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnSectionCheck := nil;
    TEventClass.Remove(AObj, @TEventClass.OnSectionCheck);
    Exit;
  end;
  AObj.OnSectionCheck := TEventClass.OnSectionCheck;
  TEventClass.Add(AObj, @TEventClass.OnSectionCheck, AEventId);
end;

function HeaderControl_GetDockClientCount(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.DockClientCount;
end;

function HeaderControl_GetDockSite(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure HeaderControl_SetDockSite(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function HeaderControl_GetAlignDisabled(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.AlignDisabled;
end;

function HeaderControl_GetMouseInClient(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.MouseInClient;
end;

function HeaderControl_GetVisibleDockClientCount(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.VisibleDockClientCount;
end;

function HeaderControl_GetBrush(AObj: THeaderControl): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function HeaderControl_GetControlCount(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function HeaderControl_GetHandle(AObj: THeaderControl): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function HeaderControl_GetParentWindow(AObj: THeaderControl): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure HeaderControl_SetParentWindow(AObj: THeaderControl; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function HeaderControl_GetTabOrder(AObj: THeaderControl): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure HeaderControl_SetTabOrder(AObj: THeaderControl; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function HeaderControl_GetTabStop(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure HeaderControl_SetTabStop(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function HeaderControl_GetUseDockManager(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure HeaderControl_SetUseDockManager(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function HeaderControl_GetAction(AObj: THeaderControl): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure HeaderControl_SetAction(AObj: THeaderControl; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure HeaderControl_GetBoundsRect(AObj: THeaderControl; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure HeaderControl_SetBoundsRect(AObj: THeaderControl; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function HeaderControl_GetClientHeight(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure HeaderControl_SetClientHeight(AObj: THeaderControl; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure HeaderControl_GetClientOrigin(AObj: THeaderControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure HeaderControl_GetClientRect(AObj: THeaderControl; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function HeaderControl_GetClientWidth(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure HeaderControl_SetClientWidth(AObj: THeaderControl; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function HeaderControl_GetControlState(AObj: THeaderControl): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure HeaderControl_SetControlState(AObj: THeaderControl; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function HeaderControl_GetControlStyle(AObj: THeaderControl): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure HeaderControl_SetControlStyle(AObj: THeaderControl; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function HeaderControl_GetExplicitLeft(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function HeaderControl_GetExplicitTop(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function HeaderControl_GetExplicitWidth(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function HeaderControl_GetExplicitHeight(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function HeaderControl_GetFloating(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function HeaderControl_GetParent(AObj: THeaderControl): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure HeaderControl_SetParent(AObj: THeaderControl; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function HeaderControl_GetAlignWithMargins(AObj: THeaderControl): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure HeaderControl_SetAlignWithMargins(AObj: THeaderControl; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function HeaderControl_GetLeft(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure HeaderControl_SetLeft(AObj: THeaderControl; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function HeaderControl_GetTop(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure HeaderControl_SetTop(AObj: THeaderControl; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function HeaderControl_GetWidth(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure HeaderControl_SetWidth(AObj: THeaderControl; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function HeaderControl_GetHeight(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure HeaderControl_SetHeight(AObj: THeaderControl; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function HeaderControl_GetCursor(AObj: THeaderControl): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure HeaderControl_SetCursor(AObj: THeaderControl; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function HeaderControl_GetHint(AObj: THeaderControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure HeaderControl_SetHint(AObj: THeaderControl; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function HeaderControl_GetMargins(AObj: THeaderControl): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure HeaderControl_SetMargins(AObj: THeaderControl; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function HeaderControl_GetCustomHint(AObj: THeaderControl): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure HeaderControl_SetCustomHint(AObj: THeaderControl; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function HeaderControl_GetComponentCount(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function HeaderControl_GetComponentIndex(AObj: THeaderControl): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure HeaderControl_SetComponentIndex(AObj: THeaderControl; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function HeaderControl_GetOwner(AObj: THeaderControl): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function HeaderControl_GetName(AObj: THeaderControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure HeaderControl_SetName(AObj: THeaderControl; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function HeaderControl_GetTag(AObj: THeaderControl): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure HeaderControl_SetTag(AObj: THeaderControl; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function HeaderControl_GetDockClients(AObj: THeaderControl; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.DockClients[Index];
end;

function HeaderControl_GetControls(AObj: THeaderControl; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function HeaderControl_GetComponents(AObj: THeaderControl; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function HeaderControl_StaticClassType: TClass; stdcall;
begin
  Result := THeaderControl;
end;

exports
  HeaderControl_Create,
  HeaderControl_Free,
  HeaderControl_FlipChildren,
  HeaderControl_CanFocus,
  HeaderControl_ContainsControl,
  HeaderControl_ControlAtPos,
  HeaderControl_DisableAlign,
  HeaderControl_EnableAlign,
  HeaderControl_FindChildControl,
  HeaderControl_Focused,
  HeaderControl_HandleAllocated,
  HeaderControl_InsertControl,
  HeaderControl_Invalidate,
  HeaderControl_PaintTo,
  HeaderControl_RemoveControl,
  HeaderControl_Realign,
  HeaderControl_Repaint,
  HeaderControl_ScaleBy,
  HeaderControl_ScrollBy,
  HeaderControl_SetBounds,
  HeaderControl_SetFocus,
  HeaderControl_Update,
  HeaderControl_UpdateControlState,
  HeaderControl_BringToFront,
  HeaderControl_ClientToScreen,
  HeaderControl_ClientToParent,
  HeaderControl_Dragging,
  HeaderControl_HasParent,
  HeaderControl_Hide,
  HeaderControl_Perform,
  HeaderControl_Refresh,
  HeaderControl_ScreenToClient,
  HeaderControl_ParentToClient,
  HeaderControl_SendToBack,
  HeaderControl_Show,
  HeaderControl_GetTextBuf,
  HeaderControl_GetTextLen,
  HeaderControl_SetTextBuf,
  HeaderControl_FindComponent,
  HeaderControl_GetNamePath,
  HeaderControl_Assign,
  HeaderControl_DisposeOf,
  HeaderControl_ClassType,
  HeaderControl_ClassName,
  HeaderControl_InstanceSize,
  HeaderControl_InheritsFrom,
  HeaderControl_Equals,
  HeaderControl_GetHashCode,
  HeaderControl_ToString,
  HeaderControl_GetAlign,
  HeaderControl_SetAlign,
  HeaderControl_GetAnchors,
  HeaderControl_SetAnchors,
  HeaderControl_GetBiDiMode,
  HeaderControl_SetBiDiMode,
  HeaderControl_GetBorderWidth,
  HeaderControl_SetBorderWidth,
  HeaderControl_GetDoubleBuffered,
  HeaderControl_SetDoubleBuffered,
  HeaderControl_GetDragCursor,
  HeaderControl_SetDragCursor,
  HeaderControl_GetDragKind,
  HeaderControl_SetDragKind,
  HeaderControl_GetDragMode,
  HeaderControl_SetDragMode,
  HeaderControl_GetEnabled,
  HeaderControl_SetEnabled,
  HeaderControl_GetFont,
  HeaderControl_SetFont,
  HeaderControl_GetFullDrag,
  HeaderControl_SetFullDrag,
  HeaderControl_GetHotTrack,
  HeaderControl_SetHotTrack,
  HeaderControl_GetImages,
  HeaderControl_SetImages,
  HeaderControl_GetShowHint,
  HeaderControl_SetShowHint,
  HeaderControl_GetStyle,
  HeaderControl_SetStyle,
  HeaderControl_GetParentDoubleBuffered,
  HeaderControl_SetParentDoubleBuffered,
  HeaderControl_GetParentFont,
  HeaderControl_SetParentFont,
  HeaderControl_GetParentShowHint,
  HeaderControl_SetParentShowHint,
  HeaderControl_GetPopupMenu,
  HeaderControl_SetPopupMenu,
  HeaderControl_GetVisible,
  HeaderControl_SetVisible,
  HeaderControl_GetStyleElements,
  HeaderControl_SetStyleElements,
  HeaderControl_SetOnContextPopup,
  HeaderControl_SetOnDragDrop,
  HeaderControl_SetOnDragOver,
  HeaderControl_SetOnEndDock,
  HeaderControl_SetOnEndDrag,
  HeaderControl_SetOnGesture,
  HeaderControl_SetOnMouseActivate,
  HeaderControl_SetOnMouseDown,
  HeaderControl_SetOnMouseEnter,
  HeaderControl_SetOnMouseLeave,
  HeaderControl_SetOnMouseMove,
  HeaderControl_SetOnMouseUp,
  HeaderControl_SetOnResize,
  HeaderControl_SetOnDrawSection,
  HeaderControl_SetOnSectionClick,
  HeaderControl_SetOnSectionResize,
  HeaderControl_SetOnSectionTrack,
  HeaderControl_SetOnSectionDrag,
  HeaderControl_SetOnSectionEndDrag,
  HeaderControl_SetOnStartDock,
  HeaderControl_GetCanvas,
  HeaderControl_SetOnSectionCheck,
  HeaderControl_GetDockClientCount,
  HeaderControl_GetDockSite,
  HeaderControl_SetDockSite,
  HeaderControl_GetAlignDisabled,
  HeaderControl_GetMouseInClient,
  HeaderControl_GetVisibleDockClientCount,
  HeaderControl_GetBrush,
  HeaderControl_GetControlCount,
  HeaderControl_GetHandle,
  HeaderControl_GetParentWindow,
  HeaderControl_SetParentWindow,
  HeaderControl_GetTabOrder,
  HeaderControl_SetTabOrder,
  HeaderControl_GetTabStop,
  HeaderControl_SetTabStop,
  HeaderControl_GetUseDockManager,
  HeaderControl_SetUseDockManager,
  HeaderControl_GetAction,
  HeaderControl_SetAction,
  HeaderControl_GetBoundsRect,
  HeaderControl_SetBoundsRect,
  HeaderControl_GetClientHeight,
  HeaderControl_SetClientHeight,
  HeaderControl_GetClientOrigin,
  HeaderControl_GetClientRect,
  HeaderControl_GetClientWidth,
  HeaderControl_SetClientWidth,
  HeaderControl_GetControlState,
  HeaderControl_SetControlState,
  HeaderControl_GetControlStyle,
  HeaderControl_SetControlStyle,
  HeaderControl_GetExplicitLeft,
  HeaderControl_GetExplicitTop,
  HeaderControl_GetExplicitWidth,
  HeaderControl_GetExplicitHeight,
  HeaderControl_GetFloating,
  HeaderControl_GetParent,
  HeaderControl_SetParent,
  HeaderControl_GetAlignWithMargins,
  HeaderControl_SetAlignWithMargins,
  HeaderControl_GetLeft,
  HeaderControl_SetLeft,
  HeaderControl_GetTop,
  HeaderControl_SetTop,
  HeaderControl_GetWidth,
  HeaderControl_SetWidth,
  HeaderControl_GetHeight,
  HeaderControl_SetHeight,
  HeaderControl_GetCursor,
  HeaderControl_SetCursor,
  HeaderControl_GetHint,
  HeaderControl_SetHint,
  HeaderControl_GetMargins,
  HeaderControl_SetMargins,
  HeaderControl_GetCustomHint,
  HeaderControl_SetCustomHint,
  HeaderControl_GetComponentCount,
  HeaderControl_GetComponentIndex,
  HeaderControl_SetComponentIndex,
  HeaderControl_GetOwner,
  HeaderControl_GetName,
  HeaderControl_SetName,
  HeaderControl_GetTag,
  HeaderControl_SetTag,
  HeaderControl_GetDockClients,
  HeaderControl_GetControls,
  HeaderControl_GetComponents,
  HeaderControl_StaticClassType;

