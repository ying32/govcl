
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Shape_Create(AOwner: TComponent): TShape; stdcall;
begin
  Result :=  TShape.Create(AOwner);
end;

procedure Shape_Free(AObj: TShape); stdcall;
begin
  AObj.Free;
end;

procedure Shape_BringToFront(AObj: TShape); stdcall;
begin
  AObj.BringToFront;
end;

procedure Shape_ClientToScreen(AObj: TShape; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure Shape_ClientToParent(AObj: TShape; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function Shape_Dragging(AObj: TShape): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function Shape_HasParent(AObj: TShape): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Shape_Hide(AObj: TShape); stdcall;
begin
  AObj.Hide;
end;

procedure Shape_Invalidate(AObj: TShape); stdcall;
begin
  AObj.Invalidate;
end;

function Shape_Perform(AObj: TShape; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Shape_Refresh(AObj: TShape); stdcall;
begin
  AObj.Refresh;
end;

procedure Shape_Repaint(AObj: TShape); stdcall;
begin
  AObj.Repaint;
end;

procedure Shape_ScreenToClient(AObj: TShape; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure Shape_ParentToClient(AObj: TShape; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure Shape_SendToBack(AObj: TShape); stdcall;
begin
  AObj.SendToBack;
end;

procedure Shape_SetBounds(AObj: TShape; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Shape_Show(AObj: TShape); stdcall;
begin
  AObj.Show;
end;

procedure Shape_Update(AObj: TShape); stdcall;
begin
  AObj.Update;
end;

function Shape_GetTextBuf(AObj: TShape; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Shape_GetTextLen(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure Shape_SetTextBuf(AObj: TShape; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function Shape_FindComponent(AObj: TShape; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Shape_GetNamePath(AObj: TShape): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Shape_Assign(AObj: TShape; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure Shape_DisposeOf(AObj: TShape); stdcall;
begin
  AObj.DisposeOf;
end;

function Shape_ClassType(AObj: TShape): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function Shape_ClassName(AObj: TShape): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Shape_InstanceSize(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function Shape_InheritsFrom(AObj: TShape; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function Shape_Equals(AObj: TShape; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Shape_GetHashCode(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Shape_ToString(AObj: TShape): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Shape_GetAlign(AObj: TShape): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Shape_SetAlign(AObj: TShape; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Shape_GetAnchors(AObj: TShape): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Shape_SetAnchors(AObj: TShape; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Shape_GetBrush(AObj: TShape): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

procedure Shape_SetBrush(AObj: TShape; AValue: TBrush); stdcall;
begin
  AObj.Brush := AValue;
end;

function Shape_GetDragCursor(AObj: TShape): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure Shape_SetDragCursor(AObj: TShape; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function Shape_GetDragKind(AObj: TShape): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure Shape_SetDragKind(AObj: TShape; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function Shape_GetDragMode(AObj: TShape): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure Shape_SetDragMode(AObj: TShape; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function Shape_GetEnabled(AObj: TShape): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Shape_SetEnabled(AObj: TShape; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Shape_GetParentShowHint(AObj: TShape): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Shape_SetParentShowHint(AObj: TShape; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function Shape_GetPen(AObj: TShape): TPen; stdcall;
begin
  Result :=  AObj.Pen;
end;

procedure Shape_SetPen(AObj: TShape; AValue: TPen); stdcall;
begin
  AObj.Pen := AValue;
end;

function Shape_GetShape(AObj: TShape): TShapeType; stdcall;
begin
  Result :=  AObj.Shape;
end;

procedure Shape_SetShape(AObj: TShape; AValue: TShapeType); stdcall;
begin
  AObj.Shape := AValue;
end;

function Shape_GetShowHint(AObj: TShape): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Shape_SetShowHint(AObj: TShape; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Shape_GetVisible(AObj: TShape): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Shape_SetVisible(AObj: TShape; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

procedure Shape_SetOnContextPopup(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnDragDrop(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnDragOver(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnEndDock(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnEndDrag(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnMouseActivate(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnMouseDown(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnMouseEnter(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnMouseLeave(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnMouseMove(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnMouseUp(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnGesture(AObj: TShape; AEventId: NativeUInt); stdcall;
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

procedure Shape_SetOnStartDock(AObj: TShape; AEventId: NativeUInt); stdcall;
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

function Shape_GetAction(AObj: TShape): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Shape_SetAction(AObj: TShape; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function Shape_GetBiDiMode(AObj: TShape): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Shape_SetBiDiMode(AObj: TShape; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure Shape_GetBoundsRect(AObj: TShape; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Shape_SetBoundsRect(AObj: TShape; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Shape_GetClientHeight(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Shape_SetClientHeight(AObj: TShape; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Shape_GetClientOrigin(AObj: TShape; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure Shape_GetClientRect(AObj: TShape; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Shape_GetClientWidth(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Shape_SetClientWidth(AObj: TShape; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Shape_GetControlState(AObj: TShape): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure Shape_SetControlState(AObj: TShape; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function Shape_GetControlStyle(AObj: TShape): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure Shape_SetControlStyle(AObj: TShape; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function Shape_GetExplicitLeft(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Shape_GetExplicitTop(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Shape_GetExplicitWidth(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Shape_GetExplicitHeight(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Shape_GetFloating(AObj: TShape): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function Shape_GetParent(AObj: TShape): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Shape_SetParent(AObj: TShape; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Shape_GetStyleElements(AObj: TShape): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Shape_SetStyleElements(AObj: TShape; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function Shape_GetAlignWithMargins(AObj: TShape): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Shape_SetAlignWithMargins(AObj: TShape; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Shape_GetLeft(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Shape_SetLeft(AObj: TShape; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Shape_GetTop(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Shape_SetTop(AObj: TShape; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Shape_GetWidth(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Shape_SetWidth(AObj: TShape; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Shape_GetHeight(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Shape_SetHeight(AObj: TShape; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Shape_GetCursor(AObj: TShape): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Shape_SetCursor(AObj: TShape; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Shape_GetHint(AObj: TShape): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Shape_SetHint(AObj: TShape; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Shape_GetMargins(AObj: TShape): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Shape_SetMargins(AObj: TShape; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Shape_GetCustomHint(AObj: TShape): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Shape_SetCustomHint(AObj: TShape; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Shape_GetComponentCount(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Shape_GetComponentIndex(AObj: TShape): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Shape_SetComponentIndex(AObj: TShape; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Shape_GetOwner(AObj: TShape): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Shape_GetName(AObj: TShape): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Shape_SetName(AObj: TShape; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Shape_GetTag(AObj: TShape): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Shape_SetTag(AObj: TShape; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Shape_GetComponents(AObj: TShape; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function Shape_StaticClassType: TClass; stdcall;
begin
  Result := TShape;
end;

exports
  Shape_Create,
  Shape_Free,
  Shape_BringToFront,
  Shape_ClientToScreen,
  Shape_ClientToParent,
  Shape_Dragging,
  Shape_HasParent,
  Shape_Hide,
  Shape_Invalidate,
  Shape_Perform,
  Shape_Refresh,
  Shape_Repaint,
  Shape_ScreenToClient,
  Shape_ParentToClient,
  Shape_SendToBack,
  Shape_SetBounds,
  Shape_Show,
  Shape_Update,
  Shape_GetTextBuf,
  Shape_GetTextLen,
  Shape_SetTextBuf,
  Shape_FindComponent,
  Shape_GetNamePath,
  Shape_Assign,
  Shape_DisposeOf,
  Shape_ClassType,
  Shape_ClassName,
  Shape_InstanceSize,
  Shape_InheritsFrom,
  Shape_Equals,
  Shape_GetHashCode,
  Shape_ToString,
  Shape_GetAlign,
  Shape_SetAlign,
  Shape_GetAnchors,
  Shape_SetAnchors,
  Shape_GetBrush,
  Shape_SetBrush,
  Shape_GetDragCursor,
  Shape_SetDragCursor,
  Shape_GetDragKind,
  Shape_SetDragKind,
  Shape_GetDragMode,
  Shape_SetDragMode,
  Shape_GetEnabled,
  Shape_SetEnabled,
  Shape_GetParentShowHint,
  Shape_SetParentShowHint,
  Shape_GetPen,
  Shape_SetPen,
  Shape_GetShape,
  Shape_SetShape,
  Shape_GetShowHint,
  Shape_SetShowHint,
  Shape_GetVisible,
  Shape_SetVisible,
  Shape_SetOnContextPopup,
  Shape_SetOnDragDrop,
  Shape_SetOnDragOver,
  Shape_SetOnEndDock,
  Shape_SetOnEndDrag,
  Shape_SetOnMouseActivate,
  Shape_SetOnMouseDown,
  Shape_SetOnMouseEnter,
  Shape_SetOnMouseLeave,
  Shape_SetOnMouseMove,
  Shape_SetOnMouseUp,
  Shape_SetOnGesture,
  Shape_SetOnStartDock,
  Shape_GetAction,
  Shape_SetAction,
  Shape_GetBiDiMode,
  Shape_SetBiDiMode,
  Shape_GetBoundsRect,
  Shape_SetBoundsRect,
  Shape_GetClientHeight,
  Shape_SetClientHeight,
  Shape_GetClientOrigin,
  Shape_GetClientRect,
  Shape_GetClientWidth,
  Shape_SetClientWidth,
  Shape_GetControlState,
  Shape_SetControlState,
  Shape_GetControlStyle,
  Shape_SetControlStyle,
  Shape_GetExplicitLeft,
  Shape_GetExplicitTop,
  Shape_GetExplicitWidth,
  Shape_GetExplicitHeight,
  Shape_GetFloating,
  Shape_GetParent,
  Shape_SetParent,
  Shape_GetStyleElements,
  Shape_SetStyleElements,
  Shape_GetAlignWithMargins,
  Shape_SetAlignWithMargins,
  Shape_GetLeft,
  Shape_SetLeft,
  Shape_GetTop,
  Shape_SetTop,
  Shape_GetWidth,
  Shape_SetWidth,
  Shape_GetHeight,
  Shape_SetHeight,
  Shape_GetCursor,
  Shape_SetCursor,
  Shape_GetHint,
  Shape_SetHint,
  Shape_GetMargins,
  Shape_SetMargins,
  Shape_GetCustomHint,
  Shape_SetCustomHint,
  Shape_GetComponentCount,
  Shape_GetComponentIndex,
  Shape_SetComponentIndex,
  Shape_GetOwner,
  Shape_GetName,
  Shape_SetName,
  Shape_GetTag,
  Shape_SetTag,
  Shape_GetComponents,
  Shape_StaticClassType;

