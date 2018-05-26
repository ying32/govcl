
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
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

procedure ScrollBox_Invalidate(AObj: TScrollBox); stdcall;
begin
  AObj.Invalidate;
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

procedure ScrollBox_BringToFront(AObj: TScrollBox); stdcall;
begin
  AObj.BringToFront;
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

function ScrollBox_ClassName(AObj: TScrollBox): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
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

function ScrollBox_GetDoubleBuffered(AObj: TScrollBox): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ScrollBox_SetDoubleBuffered(AObj: TScrollBox; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
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
    TEventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, geClick, AEventId);
end;

procedure ScrollBox_SetOnDblClick(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnEnter(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnExit(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseDown(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseEnter(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseLeave(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseMove(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseUp(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
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

procedure ScrollBox_SetOnMouseWheel(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheel := nil;
    TEventClass.Remove(AObj, geMouseWheel);
    Exit;
  end;
  AObj.OnMouseWheel := TEventClass.OnMouseWheel;
  TEventClass.Add(AObj, geMouseWheel, AEventId);
end;

procedure ScrollBox_SetOnResize(AObj: TScrollBox; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnResize := nil;
    TEventClass.Remove(AObj, geResize);
    Exit;
  end;
  AObj.OnResize := TEventClass.OnResize;
  TEventClass.Add(AObj, geResize, AEventId);
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

function ScrollBox_GetControls(AObj: TScrollBox; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ScrollBox_GetComponents(AObj: TScrollBox; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ScrollBox_Create,
  ScrollBox_Free,
  ScrollBox_CanFocus,
  ScrollBox_FlipChildren,
  ScrollBox_Focused,
  ScrollBox_HandleAllocated,
  ScrollBox_Invalidate,
  ScrollBox_Realign,
  ScrollBox_Repaint,
  ScrollBox_ScaleBy,
  ScrollBox_SetBounds,
  ScrollBox_SetFocus,
  ScrollBox_Update,
  ScrollBox_BringToFront,
  ScrollBox_HasParent,
  ScrollBox_Hide,
  ScrollBox_Perform,
  ScrollBox_Refresh,
  ScrollBox_SendToBack,
  ScrollBox_Show,
  ScrollBox_GetTextBuf,
  ScrollBox_GetTextLen,
  ScrollBox_FindComponent,
  ScrollBox_GetNamePath,
  ScrollBox_Assign,
  ScrollBox_ClassName,
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
  ScrollBox_GetDoubleBuffered,
  ScrollBox_SetDoubleBuffered,
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
  ScrollBox_SetOnDblClick,
  ScrollBox_SetOnEnter,
  ScrollBox_SetOnExit,
  ScrollBox_SetOnMouseDown,
  ScrollBox_SetOnMouseEnter,
  ScrollBox_SetOnMouseLeave,
  ScrollBox_SetOnMouseMove,
  ScrollBox_SetOnMouseUp,
  ScrollBox_SetOnMouseWheel,
  ScrollBox_SetOnResize,
  ScrollBox_GetBrush,
  ScrollBox_GetControlCount,
  ScrollBox_GetHandle,
  ScrollBox_GetParentWindow,
  ScrollBox_SetParentWindow,
  ScrollBox_GetAction,
  ScrollBox_SetAction,
  ScrollBox_GetBoundsRect,
  ScrollBox_SetBoundsRect,
  ScrollBox_GetClientHeight,
  ScrollBox_SetClientHeight,
  ScrollBox_GetClientRect,
  ScrollBox_GetClientWidth,
  ScrollBox_SetClientWidth,
  ScrollBox_GetExplicitLeft,
  ScrollBox_GetExplicitTop,
  ScrollBox_GetExplicitWidth,
  ScrollBox_GetExplicitHeight,
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
  ScrollBox_GetControls,
  ScrollBox_GetComponents;

