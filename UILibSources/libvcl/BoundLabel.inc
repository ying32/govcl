
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function BoundLabel_Create(AOwner: TComponent): TBoundLabel; stdcall;
begin
  Result :=  TBoundLabel.Create(AOwner);
end;

procedure BoundLabel_Free(AObj: TBoundLabel); stdcall;
begin
  AObj.Free;
end;

procedure BoundLabel_BringToFront(AObj: TBoundLabel); stdcall;
begin
  AObj.BringToFront;
end;

procedure BoundLabel_ClientToScreen(AObj: TBoundLabel; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure BoundLabel_ClientToParent(AObj: TBoundLabel; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function BoundLabel_Dragging(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function BoundLabel_HasParent(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure BoundLabel_Hide(AObj: TBoundLabel); stdcall;
begin
  AObj.Hide;
end;

procedure BoundLabel_Invalidate(AObj: TBoundLabel); stdcall;
begin
  AObj.Invalidate;
end;

function BoundLabel_Perform(AObj: TBoundLabel; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure BoundLabel_Refresh(AObj: TBoundLabel); stdcall;
begin
  AObj.Refresh;
end;

procedure BoundLabel_Repaint(AObj: TBoundLabel); stdcall;
begin
  AObj.Repaint;
end;

procedure BoundLabel_ScreenToClient(AObj: TBoundLabel; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure BoundLabel_ParentToClient(AObj: TBoundLabel; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure BoundLabel_SendToBack(AObj: TBoundLabel); stdcall;
begin
  AObj.SendToBack;
end;

procedure BoundLabel_SetBounds(AObj: TBoundLabel; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure BoundLabel_Show(AObj: TBoundLabel); stdcall;
begin
  AObj.Show;
end;

procedure BoundLabel_Update(AObj: TBoundLabel); stdcall;
begin
  AObj.Update;
end;

function BoundLabel_GetTextBuf(AObj: TBoundLabel; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function BoundLabel_GetTextLen(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure BoundLabel_SetTextBuf(AObj: TBoundLabel; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function BoundLabel_FindComponent(AObj: TBoundLabel; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function BoundLabel_GetNamePath(AObj: TBoundLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure BoundLabel_Assign(AObj: TBoundLabel; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure BoundLabel_DisposeOf(AObj: TBoundLabel); stdcall;
begin
  AObj.DisposeOf;
end;

function BoundLabel_ClassType(AObj: TBoundLabel): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function BoundLabel_ClassName(AObj: TBoundLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function BoundLabel_InstanceSize(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function BoundLabel_InheritsFrom(AObj: TBoundLabel; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function BoundLabel_Equals(AObj: TBoundLabel; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function BoundLabel_GetHashCode(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function BoundLabel_ToString(AObj: TBoundLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function BoundLabel_GetBiDiMode(AObj: TBoundLabel): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure BoundLabel_SetBiDiMode(AObj: TBoundLabel; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function BoundLabel_GetCaption(AObj: TBoundLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure BoundLabel_SetCaption(AObj: TBoundLabel; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function BoundLabel_GetColor(AObj: TBoundLabel): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure BoundLabel_SetColor(AObj: TBoundLabel; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function BoundLabel_GetDragCursor(AObj: TBoundLabel): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure BoundLabel_SetDragCursor(AObj: TBoundLabel; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function BoundLabel_GetDragKind(AObj: TBoundLabel): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure BoundLabel_SetDragKind(AObj: TBoundLabel; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function BoundLabel_GetDragMode(AObj: TBoundLabel): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure BoundLabel_SetDragMode(AObj: TBoundLabel; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function BoundLabel_GetFont(AObj: TBoundLabel): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure BoundLabel_SetFont(AObj: TBoundLabel; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function BoundLabel_GetHeight(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure BoundLabel_SetHeight(AObj: TBoundLabel; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function BoundLabel_GetLeft(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

function BoundLabel_GetParentColor(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure BoundLabel_SetParentColor(AObj: TBoundLabel; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function BoundLabel_GetParentFont(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure BoundLabel_SetParentFont(AObj: TBoundLabel; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function BoundLabel_GetParentShowHint(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure BoundLabel_SetParentShowHint(AObj: TBoundLabel; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function BoundLabel_GetPopupMenu(AObj: TBoundLabel): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure BoundLabel_SetPopupMenu(AObj: TBoundLabel; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function BoundLabel_GetShowAccelChar(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.ShowAccelChar;
end;

procedure BoundLabel_SetShowAccelChar(AObj: TBoundLabel; AValue: LongBool); stdcall;
begin
  AObj.ShowAccelChar := AValue;
end;

function BoundLabel_GetShowHint(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure BoundLabel_SetShowHint(AObj: TBoundLabel; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function BoundLabel_GetTop(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

function BoundLabel_GetTransparent(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure BoundLabel_SetTransparent(AObj: TBoundLabel; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function BoundLabel_GetLayout(AObj: TBoundLabel): TTextLayout; stdcall;
begin
  Result :=  AObj.Layout;
end;

procedure BoundLabel_SetLayout(AObj: TBoundLabel; AValue: TTextLayout); stdcall;
begin
  AObj.Layout := AValue;
end;

function BoundLabel_GetWordWrap(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.WordWrap;
end;

procedure BoundLabel_SetWordWrap(AObj: TBoundLabel; AValue: LongBool); stdcall;
begin
  AObj.WordWrap := AValue;
end;

function BoundLabel_GetWidth(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure BoundLabel_SetWidth(AObj: TBoundLabel; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

procedure BoundLabel_SetOnClick(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
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

procedure BoundLabel_SetOnContextPopup(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
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

procedure BoundLabel_SetOnDblClick(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDblClick := nil;
    TEventClass.Remove(AObj, geDblClick);
    Exit;
  end;
  AObj.OnDblClick := TEventClass.OnDblClick;
  TEventClass.Add(AObj, geDblClick, AEventId);
end;

procedure BoundLabel_SetOnDragDrop(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
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

procedure BoundLabel_SetOnDragOver(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
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

procedure BoundLabel_SetOnEndDock(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
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

procedure BoundLabel_SetOnEndDrag(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
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

procedure BoundLabel_SetOnGesture(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGesture := nil;
    TEventClass.Remove(AObj, geGesture);
    Exit;
  end;
  AObj.OnGesture := TEventClass.OnGesture;
  TEventClass.Add(AObj, geGesture, AEventId);
end;

procedure BoundLabel_SetOnMouseActivate(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseActivate := nil;
    TEventClass.Remove(AObj, geMouseActivate);
    Exit;
  end;
  AObj.OnMouseActivate := TEventClass.OnMouseActivate;
  TEventClass.Add(AObj, geMouseActivate, AEventId);
end;

procedure BoundLabel_SetOnMouseDown(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
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

procedure BoundLabel_SetOnMouseMove(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
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

procedure BoundLabel_SetOnMouseUp(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
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

procedure BoundLabel_SetOnStartDock(AObj: TBoundLabel; AEventId: NativeUInt); stdcall;
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

function BoundLabel_GetCanvas(AObj: TBoundLabel): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function BoundLabel_GetGlowSize(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.GlowSize;
end;

procedure BoundLabel_SetGlowSize(AObj: TBoundLabel; AValue: Integer); stdcall;
begin
  AObj.GlowSize := AValue;
end;

function BoundLabel_GetEnabled(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure BoundLabel_SetEnabled(AObj: TBoundLabel; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function BoundLabel_GetAction(AObj: TBoundLabel): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure BoundLabel_SetAction(AObj: TBoundLabel; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function BoundLabel_GetAlign(AObj: TBoundLabel): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure BoundLabel_SetAlign(AObj: TBoundLabel; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function BoundLabel_GetAnchors(AObj: TBoundLabel): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure BoundLabel_SetAnchors(AObj: TBoundLabel; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

procedure BoundLabel_GetBoundsRect(AObj: TBoundLabel; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure BoundLabel_SetBoundsRect(AObj: TBoundLabel; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function BoundLabel_GetClientHeight(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure BoundLabel_SetClientHeight(AObj: TBoundLabel; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure BoundLabel_GetClientOrigin(AObj: TBoundLabel; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure BoundLabel_GetClientRect(AObj: TBoundLabel; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function BoundLabel_GetClientWidth(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure BoundLabel_SetClientWidth(AObj: TBoundLabel; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function BoundLabel_GetControlState(AObj: TBoundLabel): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure BoundLabel_SetControlState(AObj: TBoundLabel; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function BoundLabel_GetControlStyle(AObj: TBoundLabel): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure BoundLabel_SetControlStyle(AObj: TBoundLabel; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function BoundLabel_GetExplicitLeft(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function BoundLabel_GetExplicitTop(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function BoundLabel_GetExplicitWidth(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function BoundLabel_GetExplicitHeight(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function BoundLabel_GetFloating(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function BoundLabel_GetVisible(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure BoundLabel_SetVisible(AObj: TBoundLabel; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function BoundLabel_GetParent(AObj: TBoundLabel): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure BoundLabel_SetParent(AObj: TBoundLabel; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function BoundLabel_GetStyleElements(AObj: TBoundLabel): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure BoundLabel_SetStyleElements(AObj: TBoundLabel; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function BoundLabel_GetAlignWithMargins(AObj: TBoundLabel): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure BoundLabel_SetAlignWithMargins(AObj: TBoundLabel; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function BoundLabel_GetCursor(AObj: TBoundLabel): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure BoundLabel_SetCursor(AObj: TBoundLabel; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function BoundLabel_GetHint(AObj: TBoundLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure BoundLabel_SetHint(AObj: TBoundLabel; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function BoundLabel_GetMargins(AObj: TBoundLabel): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure BoundLabel_SetMargins(AObj: TBoundLabel; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function BoundLabel_GetCustomHint(AObj: TBoundLabel): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure BoundLabel_SetCustomHint(AObj: TBoundLabel; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function BoundLabel_GetComponentCount(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function BoundLabel_GetComponentIndex(AObj: TBoundLabel): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure BoundLabel_SetComponentIndex(AObj: TBoundLabel; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function BoundLabel_GetOwner(AObj: TBoundLabel): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function BoundLabel_GetName(AObj: TBoundLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure BoundLabel_SetName(AObj: TBoundLabel; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function BoundLabel_GetTag(AObj: TBoundLabel): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure BoundLabel_SetTag(AObj: TBoundLabel; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function BoundLabel_GetComponents(AObj: TBoundLabel; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function BoundLabel_StaticClassType: TClass; stdcall;
begin
  Result := TBoundLabel;
end;

exports
  BoundLabel_Create,
  BoundLabel_Free,
  BoundLabel_BringToFront,
  BoundLabel_ClientToScreen,
  BoundLabel_ClientToParent,
  BoundLabel_Dragging,
  BoundLabel_HasParent,
  BoundLabel_Hide,
  BoundLabel_Invalidate,
  BoundLabel_Perform,
  BoundLabel_Refresh,
  BoundLabel_Repaint,
  BoundLabel_ScreenToClient,
  BoundLabel_ParentToClient,
  BoundLabel_SendToBack,
  BoundLabel_SetBounds,
  BoundLabel_Show,
  BoundLabel_Update,
  BoundLabel_GetTextBuf,
  BoundLabel_GetTextLen,
  BoundLabel_SetTextBuf,
  BoundLabel_FindComponent,
  BoundLabel_GetNamePath,
  BoundLabel_Assign,
  BoundLabel_DisposeOf,
  BoundLabel_ClassType,
  BoundLabel_ClassName,
  BoundLabel_InstanceSize,
  BoundLabel_InheritsFrom,
  BoundLabel_Equals,
  BoundLabel_GetHashCode,
  BoundLabel_ToString,
  BoundLabel_GetBiDiMode,
  BoundLabel_SetBiDiMode,
  BoundLabel_GetCaption,
  BoundLabel_SetCaption,
  BoundLabel_GetColor,
  BoundLabel_SetColor,
  BoundLabel_GetDragCursor,
  BoundLabel_SetDragCursor,
  BoundLabel_GetDragKind,
  BoundLabel_SetDragKind,
  BoundLabel_GetDragMode,
  BoundLabel_SetDragMode,
  BoundLabel_GetFont,
  BoundLabel_SetFont,
  BoundLabel_GetHeight,
  BoundLabel_SetHeight,
  BoundLabel_GetLeft,
  BoundLabel_GetParentColor,
  BoundLabel_SetParentColor,
  BoundLabel_GetParentFont,
  BoundLabel_SetParentFont,
  BoundLabel_GetParentShowHint,
  BoundLabel_SetParentShowHint,
  BoundLabel_GetPopupMenu,
  BoundLabel_SetPopupMenu,
  BoundLabel_GetShowAccelChar,
  BoundLabel_SetShowAccelChar,
  BoundLabel_GetShowHint,
  BoundLabel_SetShowHint,
  BoundLabel_GetTop,
  BoundLabel_GetTransparent,
  BoundLabel_SetTransparent,
  BoundLabel_GetLayout,
  BoundLabel_SetLayout,
  BoundLabel_GetWordWrap,
  BoundLabel_SetWordWrap,
  BoundLabel_GetWidth,
  BoundLabel_SetWidth,
  BoundLabel_SetOnClick,
  BoundLabel_SetOnContextPopup,
  BoundLabel_SetOnDblClick,
  BoundLabel_SetOnDragDrop,
  BoundLabel_SetOnDragOver,
  BoundLabel_SetOnEndDock,
  BoundLabel_SetOnEndDrag,
  BoundLabel_SetOnGesture,
  BoundLabel_SetOnMouseActivate,
  BoundLabel_SetOnMouseDown,
  BoundLabel_SetOnMouseMove,
  BoundLabel_SetOnMouseUp,
  BoundLabel_SetOnStartDock,
  BoundLabel_GetCanvas,
  BoundLabel_GetGlowSize,
  BoundLabel_SetGlowSize,
  BoundLabel_GetEnabled,
  BoundLabel_SetEnabled,
  BoundLabel_GetAction,
  BoundLabel_SetAction,
  BoundLabel_GetAlign,
  BoundLabel_SetAlign,
  BoundLabel_GetAnchors,
  BoundLabel_SetAnchors,
  BoundLabel_GetBoundsRect,
  BoundLabel_SetBoundsRect,
  BoundLabel_GetClientHeight,
  BoundLabel_SetClientHeight,
  BoundLabel_GetClientOrigin,
  BoundLabel_GetClientRect,
  BoundLabel_GetClientWidth,
  BoundLabel_SetClientWidth,
  BoundLabel_GetControlState,
  BoundLabel_SetControlState,
  BoundLabel_GetControlStyle,
  BoundLabel_SetControlStyle,
  BoundLabel_GetExplicitLeft,
  BoundLabel_GetExplicitTop,
  BoundLabel_GetExplicitWidth,
  BoundLabel_GetExplicitHeight,
  BoundLabel_GetFloating,
  BoundLabel_GetVisible,
  BoundLabel_SetVisible,
  BoundLabel_GetParent,
  BoundLabel_SetParent,
  BoundLabel_GetStyleElements,
  BoundLabel_SetStyleElements,
  BoundLabel_GetAlignWithMargins,
  BoundLabel_SetAlignWithMargins,
  BoundLabel_GetCursor,
  BoundLabel_SetCursor,
  BoundLabel_GetHint,
  BoundLabel_SetHint,
  BoundLabel_GetMargins,
  BoundLabel_SetMargins,
  BoundLabel_GetCustomHint,
  BoundLabel_SetCustomHint,
  BoundLabel_GetComponentCount,
  BoundLabel_GetComponentIndex,
  BoundLabel_SetComponentIndex,
  BoundLabel_GetOwner,
  BoundLabel_GetName,
  BoundLabel_SetName,
  BoundLabel_GetTag,
  BoundLabel_SetTag,
  BoundLabel_GetComponents,
  BoundLabel_StaticClassType;

