
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function PaintBox_Create(AOwner: TComponent): TPaintBox; stdcall;
begin
  Result :=  TPaintBox.Create(AOwner);
end;

procedure PaintBox_Free(AObj: TPaintBox); stdcall;
begin
  AObj.Free;
end;

procedure PaintBox_BringToFront(AObj: TPaintBox); stdcall;
begin
  AObj.BringToFront;
end;

procedure PaintBox_ClientToScreen(AObj: TPaintBox; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure PaintBox_ClientToParent(AObj: TPaintBox; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function PaintBox_Dragging(AObj: TPaintBox): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function PaintBox_HasParent(AObj: TPaintBox): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure PaintBox_Hide(AObj: TPaintBox); stdcall;
begin
  AObj.Hide;
end;

procedure PaintBox_Invalidate(AObj: TPaintBox); stdcall;
begin
  AObj.Invalidate;
end;

function PaintBox_Perform(AObj: TPaintBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure PaintBox_Refresh(AObj: TPaintBox); stdcall;
begin
  AObj.Refresh;
end;

procedure PaintBox_Repaint(AObj: TPaintBox); stdcall;
begin
  AObj.Repaint;
end;

procedure PaintBox_ScreenToClient(AObj: TPaintBox; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure PaintBox_ParentToClient(AObj: TPaintBox; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure PaintBox_SendToBack(AObj: TPaintBox); stdcall;
begin
  AObj.SendToBack;
end;

procedure PaintBox_SetBounds(AObj: TPaintBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure PaintBox_Show(AObj: TPaintBox); stdcall;
begin
  AObj.Show;
end;

procedure PaintBox_Update(AObj: TPaintBox); stdcall;
begin
  AObj.Update;
end;

function PaintBox_GetTextBuf(AObj: TPaintBox; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function PaintBox_GetTextLen(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

procedure PaintBox_SetTextBuf(AObj: TPaintBox; Buffer: PWideChar); stdcall;
begin
  AObj.SetTextBuf(Buffer);
end;

function PaintBox_FindComponent(AObj: TPaintBox; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function PaintBox_GetNamePath(AObj: TPaintBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure PaintBox_Assign(AObj: TPaintBox; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure PaintBox_DisposeOf(AObj: TPaintBox); stdcall;
begin
  AObj.DisposeOf;
end;

function PaintBox_ClassType(AObj: TPaintBox): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function PaintBox_ClassName(AObj: TPaintBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function PaintBox_InstanceSize(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function PaintBox_InheritsFrom(AObj: TPaintBox; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function PaintBox_Equals(AObj: TPaintBox; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function PaintBox_GetHashCode(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function PaintBox_ToString(AObj: TPaintBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function PaintBox_GetCanvas(AObj: TPaintBox): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function PaintBox_GetAlign(AObj: TPaintBox): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure PaintBox_SetAlign(AObj: TPaintBox; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function PaintBox_GetAnchors(AObj: TPaintBox): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure PaintBox_SetAnchors(AObj: TPaintBox; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function PaintBox_GetColor(AObj: TPaintBox): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure PaintBox_SetColor(AObj: TPaintBox; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function PaintBox_GetDragCursor(AObj: TPaintBox): TCursor; stdcall;
begin
  Result :=  AObj.DragCursor;
end;

procedure PaintBox_SetDragCursor(AObj: TPaintBox; AValue: TCursor); stdcall;
begin
  AObj.DragCursor := AValue;
end;

function PaintBox_GetDragKind(AObj: TPaintBox): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure PaintBox_SetDragKind(AObj: TPaintBox; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function PaintBox_GetDragMode(AObj: TPaintBox): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure PaintBox_SetDragMode(AObj: TPaintBox; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function PaintBox_GetEnabled(AObj: TPaintBox): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure PaintBox_SetEnabled(AObj: TPaintBox; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function PaintBox_GetFont(AObj: TPaintBox): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure PaintBox_SetFont(AObj: TPaintBox; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function PaintBox_GetParentColor(AObj: TPaintBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure PaintBox_SetParentColor(AObj: TPaintBox; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function PaintBox_GetParentFont(AObj: TPaintBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure PaintBox_SetParentFont(AObj: TPaintBox; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function PaintBox_GetParentShowHint(AObj: TPaintBox): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure PaintBox_SetParentShowHint(AObj: TPaintBox; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function PaintBox_GetPopupMenu(AObj: TPaintBox): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure PaintBox_SetPopupMenu(AObj: TPaintBox; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function PaintBox_GetShowHint(AObj: TPaintBox): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure PaintBox_SetShowHint(AObj: TPaintBox; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function PaintBox_GetVisible(AObj: TPaintBox): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure PaintBox_SetVisible(AObj: TPaintBox; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

procedure PaintBox_SetOnClick(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnContextPopup(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnDblClick(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnDragDrop(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnDragOver(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnEndDock(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnEndDrag(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnGesture(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnMouseActivate(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnMouseDown(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnMouseEnter(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnMouseLeave(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnMouseMove(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnMouseUp(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

procedure PaintBox_SetOnPaint(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnPaint := nil;
    TEventClass.Remove(AObj, @TEventClass.OnPaint);
    Exit;
  end;
  AObj.OnPaint := TEventClass.OnPaint;
  TEventClass.Add(AObj, @TEventClass.OnPaint, AEventId);
end;

procedure PaintBox_SetOnStartDock(AObj: TPaintBox; AEventId: NativeUInt); stdcall;
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

function PaintBox_GetAction(AObj: TPaintBox): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure PaintBox_SetAction(AObj: TPaintBox; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function PaintBox_GetBiDiMode(AObj: TPaintBox): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure PaintBox_SetBiDiMode(AObj: TPaintBox; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure PaintBox_GetBoundsRect(AObj: TPaintBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure PaintBox_SetBoundsRect(AObj: TPaintBox; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function PaintBox_GetClientHeight(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure PaintBox_SetClientHeight(AObj: TPaintBox; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure PaintBox_GetClientOrigin(AObj: TPaintBox; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientOrigin;
end;

procedure PaintBox_GetClientRect(AObj: TPaintBox; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function PaintBox_GetClientWidth(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure PaintBox_SetClientWidth(AObj: TPaintBox; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function PaintBox_GetControlState(AObj: TPaintBox): TControlState; stdcall;
begin
  Result :=  AObj.ControlState;
end;

procedure PaintBox_SetControlState(AObj: TPaintBox; AValue: TControlState); stdcall;
begin
  AObj.ControlState := AValue;
end;

function PaintBox_GetControlStyle(AObj: TPaintBox): TControlStyle; stdcall;
begin
  Result :=  AObj.ControlStyle;
end;

procedure PaintBox_SetControlStyle(AObj: TPaintBox; AValue: TControlStyle); stdcall;
begin
  AObj.ControlStyle := AValue;
end;

function PaintBox_GetExplicitLeft(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function PaintBox_GetExplicitTop(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function PaintBox_GetExplicitWidth(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function PaintBox_GetExplicitHeight(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function PaintBox_GetFloating(AObj: TPaintBox): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function PaintBox_GetParent(AObj: TPaintBox): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure PaintBox_SetParent(AObj: TPaintBox; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function PaintBox_GetStyleElements(AObj: TPaintBox): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure PaintBox_SetStyleElements(AObj: TPaintBox; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function PaintBox_GetAlignWithMargins(AObj: TPaintBox): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure PaintBox_SetAlignWithMargins(AObj: TPaintBox; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function PaintBox_GetLeft(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure PaintBox_SetLeft(AObj: TPaintBox; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function PaintBox_GetTop(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure PaintBox_SetTop(AObj: TPaintBox; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function PaintBox_GetWidth(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure PaintBox_SetWidth(AObj: TPaintBox; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function PaintBox_GetHeight(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure PaintBox_SetHeight(AObj: TPaintBox; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function PaintBox_GetCursor(AObj: TPaintBox): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure PaintBox_SetCursor(AObj: TPaintBox; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function PaintBox_GetHint(AObj: TPaintBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure PaintBox_SetHint(AObj: TPaintBox; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function PaintBox_GetMargins(AObj: TPaintBox): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure PaintBox_SetMargins(AObj: TPaintBox; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function PaintBox_GetCustomHint(AObj: TPaintBox): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure PaintBox_SetCustomHint(AObj: TPaintBox; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function PaintBox_GetComponentCount(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function PaintBox_GetComponentIndex(AObj: TPaintBox): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure PaintBox_SetComponentIndex(AObj: TPaintBox; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function PaintBox_GetOwner(AObj: TPaintBox): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function PaintBox_GetName(AObj: TPaintBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure PaintBox_SetName(AObj: TPaintBox; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function PaintBox_GetTag(AObj: TPaintBox): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure PaintBox_SetTag(AObj: TPaintBox; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function PaintBox_GetComponents(AObj: TPaintBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function PaintBox_StaticClassType: TClass; stdcall;
begin
  Result := TPaintBox;
end;

exports
  PaintBox_Create,
  PaintBox_Free,
  PaintBox_BringToFront,
  PaintBox_ClientToScreen,
  PaintBox_ClientToParent,
  PaintBox_Dragging,
  PaintBox_HasParent,
  PaintBox_Hide,
  PaintBox_Invalidate,
  PaintBox_Perform,
  PaintBox_Refresh,
  PaintBox_Repaint,
  PaintBox_ScreenToClient,
  PaintBox_ParentToClient,
  PaintBox_SendToBack,
  PaintBox_SetBounds,
  PaintBox_Show,
  PaintBox_Update,
  PaintBox_GetTextBuf,
  PaintBox_GetTextLen,
  PaintBox_SetTextBuf,
  PaintBox_FindComponent,
  PaintBox_GetNamePath,
  PaintBox_Assign,
  PaintBox_DisposeOf,
  PaintBox_ClassType,
  PaintBox_ClassName,
  PaintBox_InstanceSize,
  PaintBox_InheritsFrom,
  PaintBox_Equals,
  PaintBox_GetHashCode,
  PaintBox_ToString,
  PaintBox_GetCanvas,
  PaintBox_GetAlign,
  PaintBox_SetAlign,
  PaintBox_GetAnchors,
  PaintBox_SetAnchors,
  PaintBox_GetColor,
  PaintBox_SetColor,
  PaintBox_GetDragCursor,
  PaintBox_SetDragCursor,
  PaintBox_GetDragKind,
  PaintBox_SetDragKind,
  PaintBox_GetDragMode,
  PaintBox_SetDragMode,
  PaintBox_GetEnabled,
  PaintBox_SetEnabled,
  PaintBox_GetFont,
  PaintBox_SetFont,
  PaintBox_GetParentColor,
  PaintBox_SetParentColor,
  PaintBox_GetParentFont,
  PaintBox_SetParentFont,
  PaintBox_GetParentShowHint,
  PaintBox_SetParentShowHint,
  PaintBox_GetPopupMenu,
  PaintBox_SetPopupMenu,
  PaintBox_GetShowHint,
  PaintBox_SetShowHint,
  PaintBox_GetVisible,
  PaintBox_SetVisible,
  PaintBox_SetOnClick,
  PaintBox_SetOnContextPopup,
  PaintBox_SetOnDblClick,
  PaintBox_SetOnDragDrop,
  PaintBox_SetOnDragOver,
  PaintBox_SetOnEndDock,
  PaintBox_SetOnEndDrag,
  PaintBox_SetOnGesture,
  PaintBox_SetOnMouseActivate,
  PaintBox_SetOnMouseDown,
  PaintBox_SetOnMouseEnter,
  PaintBox_SetOnMouseLeave,
  PaintBox_SetOnMouseMove,
  PaintBox_SetOnMouseUp,
  PaintBox_SetOnPaint,
  PaintBox_SetOnStartDock,
  PaintBox_GetAction,
  PaintBox_SetAction,
  PaintBox_GetBiDiMode,
  PaintBox_SetBiDiMode,
  PaintBox_GetBoundsRect,
  PaintBox_SetBoundsRect,
  PaintBox_GetClientHeight,
  PaintBox_SetClientHeight,
  PaintBox_GetClientOrigin,
  PaintBox_GetClientRect,
  PaintBox_GetClientWidth,
  PaintBox_SetClientWidth,
  PaintBox_GetControlState,
  PaintBox_SetControlState,
  PaintBox_GetControlStyle,
  PaintBox_SetControlStyle,
  PaintBox_GetExplicitLeft,
  PaintBox_GetExplicitTop,
  PaintBox_GetExplicitWidth,
  PaintBox_GetExplicitHeight,
  PaintBox_GetFloating,
  PaintBox_GetParent,
  PaintBox_SetParent,
  PaintBox_GetStyleElements,
  PaintBox_SetStyleElements,
  PaintBox_GetAlignWithMargins,
  PaintBox_SetAlignWithMargins,
  PaintBox_GetLeft,
  PaintBox_SetLeft,
  PaintBox_GetTop,
  PaintBox_SetTop,
  PaintBox_GetWidth,
  PaintBox_SetWidth,
  PaintBox_GetHeight,
  PaintBox_SetHeight,
  PaintBox_GetCursor,
  PaintBox_SetCursor,
  PaintBox_GetHint,
  PaintBox_SetHint,
  PaintBox_GetMargins,
  PaintBox_SetMargins,
  PaintBox_GetCustomHint,
  PaintBox_SetCustomHint,
  PaintBox_GetComponentCount,
  PaintBox_GetComponentIndex,
  PaintBox_SetComponentIndex,
  PaintBox_GetOwner,
  PaintBox_GetName,
  PaintBox_SetName,
  PaintBox_GetTag,
  PaintBox_SetTag,
  PaintBox_GetComponents,
  PaintBox_StaticClassType;

