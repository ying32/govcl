
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function CheckBox_Create(AOwner: TComponent): TCheckBox; stdcall;
begin
  Result :=  TCheckBox.Create(AOwner);
end;

procedure CheckBox_Free(AObj: TCheckBox); stdcall;
begin
  AObj.Free;
end;

function CheckBox_CanFocus(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure CheckBox_FlipChildren(AObj: TCheckBox; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function CheckBox_Focused(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function CheckBox_HandleAllocated(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure CheckBox_Invalidate(AObj: TCheckBox); stdcall;
begin
  AObj.Invalidate;
end;

procedure CheckBox_Realign(AObj: TCheckBox); stdcall;
begin
  AObj.Realign;
end;

procedure CheckBox_Repaint(AObj: TCheckBox); stdcall;
begin
  AObj.Repaint;
end;

procedure CheckBox_ScaleBy(AObj: TCheckBox; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure CheckBox_SetBounds(AObj: TCheckBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure CheckBox_SetFocus(AObj: TCheckBox); stdcall;
begin
  AObj.SetFocus;
end;

procedure CheckBox_Update(AObj: TCheckBox); stdcall;
begin
  AObj.Update;
end;

procedure CheckBox_BringToFront(AObj: TCheckBox); stdcall;
begin
  AObj.BringToFront;
end;

procedure CheckBox_ClientToScreen(AObj: TCheckBox; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure CheckBox_ClientToParent(AObj: TCheckBox; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function CheckBox_Dragging(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function CheckBox_HasParent(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure CheckBox_Hide(AObj: TCheckBox); stdcall;
begin
  AObj.Hide;
end;

function CheckBox_Perform(AObj: TCheckBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure CheckBox_Refresh(AObj: TCheckBox); stdcall;
begin
  AObj.Refresh;
end;

procedure CheckBox_ScreenToClient(AObj: TCheckBox; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure CheckBox_ParentToClient(AObj: TCheckBox; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure CheckBox_SendToBack(AObj: TCheckBox); stdcall;
begin
  AObj.SendToBack;
end;

procedure CheckBox_Show(AObj: TCheckBox); stdcall;
begin
  AObj.Show;
end;

function CheckBox_GetTextBuf(AObj: TCheckBox; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function CheckBox_GetTextLen(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function CheckBox_FindComponent(AObj: TCheckBox; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function CheckBox_GetNamePath(AObj: TCheckBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure CheckBox_Assign(AObj: TCheckBox; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure CheckBox_DisposeOf(AObj: TCheckBox); stdcall;
begin
  AObj.DisposeOf;
end;

function CheckBox_ClassType(AObj: TCheckBox): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function CheckBox_ClassName(AObj: TCheckBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function CheckBox_InstanceSize(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function CheckBox_InheritsFrom(AObj: TCheckBox; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function CheckBox_Equals(AObj: TCheckBox; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function CheckBox_GetHashCode(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function CheckBox_ToString(AObj: TCheckBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function CheckBox_GetAction(AObj: TCheckBox): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure CheckBox_SetAction(AObj: TCheckBox; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function CheckBox_GetAlign(AObj: TCheckBox): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure CheckBox_SetAlign(AObj: TCheckBox; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function CheckBox_GetAlignment(AObj: TCheckBox): TLeftRight; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure CheckBox_SetAlignment(AObj: TCheckBox; AValue: TLeftRight); stdcall;
begin
  AObj.Alignment := AValue;
end;

function CheckBox_GetAllowGrayed(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.AllowGrayed;
end;

procedure CheckBox_SetAllowGrayed(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.AllowGrayed := AValue;
end;

function CheckBox_GetAnchors(AObj: TCheckBox): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure CheckBox_SetAnchors(AObj: TCheckBox; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function CheckBox_GetBiDiMode(AObj: TCheckBox): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure CheckBox_SetBiDiMode(AObj: TCheckBox; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function CheckBox_GetCaption(AObj: TCheckBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure CheckBox_SetCaption(AObj: TCheckBox; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function CheckBox_GetChecked(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.Checked;
end;

procedure CheckBox_SetChecked(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.Checked := AValue;
end;

function CheckBox_GetColor(AObj: TCheckBox): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure CheckBox_SetColor(AObj: TCheckBox; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function CheckBox_GetDoubleBuffered(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure CheckBox_SetDoubleBuffered(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function CheckBox_GetDragCursor(AObj: TCheckBox): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure CheckBox_SetDragCursor(AObj: TCheckBox; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function CheckBox_GetDragKind(AObj: TCheckBox): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure CheckBox_SetDragKind(AObj: TCheckBox; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function CheckBox_GetDragMode(AObj: TCheckBox): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure CheckBox_SetDragMode(AObj: TCheckBox; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function CheckBox_GetEnabled(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure CheckBox_SetEnabled(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function CheckBox_GetFont(AObj: TCheckBox): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure CheckBox_SetFont(AObj: TCheckBox; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function CheckBox_GetParentColor(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure CheckBox_SetParentColor(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function CheckBox_GetParentCtl3D(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure CheckBox_SetParentCtl3D(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function CheckBox_GetParentDoubleBuffered(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure CheckBox_SetParentDoubleBuffered(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function CheckBox_GetParentFont(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure CheckBox_SetParentFont(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function CheckBox_GetParentShowHint(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure CheckBox_SetParentShowHint(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function CheckBox_GetPopupMenu(AObj: TCheckBox): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure CheckBox_SetPopupMenu(AObj: TCheckBox; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function CheckBox_GetShowHint(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure CheckBox_SetShowHint(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function CheckBox_GetState(AObj: TCheckBox): TCheckBoxState; stdcall;
begin
  Result :=  AObj.State;
end;

procedure CheckBox_SetState(AObj: TCheckBox; AValue: TCheckBoxState); stdcall;
begin
  AObj.State := AValue;
end;

function CheckBox_GetTabOrder(AObj: TCheckBox): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure CheckBox_SetTabOrder(AObj: TCheckBox; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function CheckBox_GetTabStop(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure CheckBox_SetTabStop(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function CheckBox_GetVisible(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure CheckBox_SetVisible(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function CheckBox_GetWordWrap(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.WordWrap;
end;

procedure CheckBox_SetWordWrap(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.WordWrap := AValue;
end;

function CheckBox_GetStyleElements(AObj: TCheckBox): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure CheckBox_SetStyleElements(AObj: TCheckBox; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure CheckBox_SetOnClick(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    TEventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, geClick, AEventId);
end;

procedure CheckBox_SetOnContextPopup(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnContextPopup := nil;
    TEventClass.Remove(AObj, geContextPopup);
    Exit;
  end;
  AObj.OnContextPopup := TEventClass.OnContextPopup;
  TEventClass.Add(AObj, geContextPopup, AEventId);
end;

procedure CheckBox_SetOnDragDrop(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragDrop := nil;
    TEventClass.Remove(AObj, geDragDrop);
    Exit;
  end;
  AObj.OnDragDrop := TEventClass.OnDragDrop;
  TEventClass.Add(AObj, geDragDrop, AEventId);
end;

procedure CheckBox_SetOnDragOver(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragOver := nil;
    TEventClass.Remove(AObj, geDragOver);
    Exit;
  end;
  AObj.OnDragOver := TEventClass.OnDragOver;
  TEventClass.Add(AObj, geDragOver, AEventId);
end;

procedure CheckBox_SetOnEndDock(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDock := nil;
    TEventClass.Remove(AObj, geEndDock);
    Exit;
  end;
  AObj.OnEndDock := TEventClass.OnEndDock;
  TEventClass.Add(AObj, geEndDock, AEventId);
end;

procedure CheckBox_SetOnEndDrag(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDrag := nil;
    TEventClass.Remove(AObj, geEndDrag);
    Exit;
  end;
  AObj.OnEndDrag := TEventClass.OnEndDrag;
  TEventClass.Add(AObj, geEndDrag, AEventId);
end;

procedure CheckBox_SetOnEnter(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
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

procedure CheckBox_SetOnExit(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
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

procedure CheckBox_SetOnKeyDown(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyDown := nil;
    TEventClass.Remove(AObj, geKeyDown);
    Exit;
  end;
  AObj.OnKeyDown := TEventClass.OnKeyDown;
  TEventClass.Add(AObj, geKeyDown, AEventId);
end;

procedure CheckBox_SetOnKeyPress(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyPress := nil;
    TEventClass.Remove(AObj, geKeyPress);
    Exit;
  end;
  AObj.OnKeyPress := TEventClass.OnKeyPress;
  TEventClass.Add(AObj, geKeyPress, AEventId);
end;

procedure CheckBox_SetOnKeyUp(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyUp := nil;
    TEventClass.Remove(AObj, geKeyUp);
    Exit;
  end;
  AObj.OnKeyUp := TEventClass.OnKeyUp;
  TEventClass.Add(AObj, geKeyUp, AEventId);
end;

procedure CheckBox_SetOnMouseDown(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
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

procedure CheckBox_SetOnMouseEnter(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
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

procedure CheckBox_SetOnMouseLeave(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
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

procedure CheckBox_SetOnMouseMove(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
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

procedure CheckBox_SetOnMouseUp(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
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

procedure CheckBox_SetOnStartDock(AObj: TCheckBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnStartDock := nil;
    TEventClass.Remove(AObj, geStartDock);
    Exit;
  end;
  AObj.OnStartDock := TEventClass.OnStartDock;
  TEventClass.Add(AObj, geStartDock, AEventId);
end;

function CheckBox_GetDockSite(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure CheckBox_SetDockSite(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function CheckBox_GetBrush(AObj: TCheckBox): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function CheckBox_GetControlCount(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function CheckBox_GetHandle(AObj: TCheckBox): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function CheckBox_GetParentWindow(AObj: TCheckBox): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure CheckBox_SetParentWindow(AObj: TCheckBox; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function CheckBox_GetUseDockManager(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure CheckBox_SetUseDockManager(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

procedure CheckBox_GetBoundsRect(AObj: TCheckBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure CheckBox_SetBoundsRect(AObj: TCheckBox; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function CheckBox_GetClientHeight(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure CheckBox_SetClientHeight(AObj: TCheckBox; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure CheckBox_GetClientRect(AObj: TCheckBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function CheckBox_GetClientWidth(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure CheckBox_SetClientWidth(AObj: TCheckBox; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function CheckBox_GetExplicitLeft(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function CheckBox_GetExplicitTop(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function CheckBox_GetExplicitWidth(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function CheckBox_GetExplicitHeight(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function CheckBox_GetFloating(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function CheckBox_GetParent(AObj: TCheckBox): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure CheckBox_SetParent(AObj: TCheckBox; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function CheckBox_GetAlignWithMargins(AObj: TCheckBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure CheckBox_SetAlignWithMargins(AObj: TCheckBox; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function CheckBox_GetLeft(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure CheckBox_SetLeft(AObj: TCheckBox; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function CheckBox_GetTop(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure CheckBox_SetTop(AObj: TCheckBox; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function CheckBox_GetWidth(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure CheckBox_SetWidth(AObj: TCheckBox; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function CheckBox_GetHeight(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure CheckBox_SetHeight(AObj: TCheckBox; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function CheckBox_GetCursor(AObj: TCheckBox): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure CheckBox_SetCursor(AObj: TCheckBox; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function CheckBox_GetHint(AObj: TCheckBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure CheckBox_SetHint(AObj: TCheckBox; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function CheckBox_GetMargins(AObj: TCheckBox): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure CheckBox_SetMargins(AObj: TCheckBox; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function CheckBox_GetCustomHint(AObj: TCheckBox): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure CheckBox_SetCustomHint(AObj: TCheckBox; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function CheckBox_GetComponentCount(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function CheckBox_GetComponentIndex(AObj: TCheckBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure CheckBox_SetComponentIndex(AObj: TCheckBox; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function CheckBox_GetOwner(AObj: TCheckBox): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function CheckBox_GetName(AObj: TCheckBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure CheckBox_SetName(AObj: TCheckBox; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function CheckBox_GetTag(AObj: TCheckBox): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure CheckBox_SetTag(AObj: TCheckBox; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function CheckBox_GetControls(AObj: TCheckBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function CheckBox_GetComponents(AObj: TCheckBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function CheckBox_StaticClassType: TClass; stdcall;
begin
  Result := TCheckBox;
end;

exports
  CheckBox_Create,
  CheckBox_Free,
  CheckBox_CanFocus,
  CheckBox_FlipChildren,
  CheckBox_Focused,
  CheckBox_HandleAllocated,
  CheckBox_Invalidate,
  CheckBox_Realign,
  CheckBox_Repaint,
  CheckBox_ScaleBy,
  CheckBox_SetBounds,
  CheckBox_SetFocus,
  CheckBox_Update,
  CheckBox_BringToFront,
  CheckBox_ClientToScreen,
  CheckBox_ClientToParent,
  CheckBox_Dragging,
  CheckBox_HasParent,
  CheckBox_Hide,
  CheckBox_Perform,
  CheckBox_Refresh,
  CheckBox_ScreenToClient,
  CheckBox_ParentToClient,
  CheckBox_SendToBack,
  CheckBox_Show,
  CheckBox_GetTextBuf,
  CheckBox_GetTextLen,
  CheckBox_FindComponent,
  CheckBox_GetNamePath,
  CheckBox_Assign,
  CheckBox_DisposeOf,
  CheckBox_ClassType,
  CheckBox_ClassName,
  CheckBox_InstanceSize,
  CheckBox_InheritsFrom,
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
  CheckBox_GetDragCursor,
  CheckBox_SetDragCursor,
  CheckBox_GetDragKind,
  CheckBox_SetDragKind,
  CheckBox_GetDragMode,
  CheckBox_SetDragMode,
  CheckBox_GetEnabled,
  CheckBox_SetEnabled,
  CheckBox_GetFont,
  CheckBox_SetFont,
  CheckBox_GetParentColor,
  CheckBox_SetParentColor,
  CheckBox_GetParentCtl3D,
  CheckBox_SetParentCtl3D,
  CheckBox_GetParentDoubleBuffered,
  CheckBox_SetParentDoubleBuffered,
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
  CheckBox_GetWordWrap,
  CheckBox_SetWordWrap,
  CheckBox_GetStyleElements,
  CheckBox_SetStyleElements,
  CheckBox_SetOnClick,
  CheckBox_SetOnContextPopup,
  CheckBox_SetOnDragDrop,
  CheckBox_SetOnDragOver,
  CheckBox_SetOnEndDock,
  CheckBox_SetOnEndDrag,
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
  CheckBox_SetOnStartDock,
  CheckBox_GetDockSite,
  CheckBox_SetDockSite,
  CheckBox_GetBrush,
  CheckBox_GetControlCount,
  CheckBox_GetHandle,
  CheckBox_GetParentWindow,
  CheckBox_SetParentWindow,
  CheckBox_GetUseDockManager,
  CheckBox_SetUseDockManager,
  CheckBox_GetBoundsRect,
  CheckBox_SetBoundsRect,
  CheckBox_GetClientHeight,
  CheckBox_SetClientHeight,
  CheckBox_GetClientRect,
  CheckBox_GetClientWidth,
  CheckBox_SetClientWidth,
  CheckBox_GetExplicitLeft,
  CheckBox_GetExplicitTop,
  CheckBox_GetExplicitWidth,
  CheckBox_GetExplicitHeight,
  CheckBox_GetFloating,
  CheckBox_GetParent,
  CheckBox_SetParent,
  CheckBox_GetAlignWithMargins,
  CheckBox_SetAlignWithMargins,
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
  CheckBox_GetMargins,
  CheckBox_SetMargins,
  CheckBox_GetCustomHint,
  CheckBox_SetCustomHint,
  CheckBox_GetComponentCount,
  CheckBox_GetComponentIndex,
  CheckBox_SetComponentIndex,
  CheckBox_GetOwner,
  CheckBox_GetName,
  CheckBox_SetName,
  CheckBox_GetTag,
  CheckBox_SetTag,
  CheckBox_GetControls,
  CheckBox_GetComponents,
  CheckBox_StaticClassType;

