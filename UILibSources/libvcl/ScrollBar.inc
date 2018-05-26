
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ScrollBar_Create(AOwner: TComponent): TScrollBar; stdcall;
begin
  Result :=  TScrollBar.Create(AOwner);
end;

procedure ScrollBar_Free(AObj: TScrollBar); stdcall;
begin
  AObj.Free;
end;

procedure ScrollBar_SetParams(AObj: TScrollBar; APosition: Integer; AMin: Integer; AMax: Integer); stdcall;
begin
  AObj.SetParams(APosition, AMin, AMax);
end;

function ScrollBar_CanFocus(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure ScrollBar_FlipChildren(AObj: TScrollBar; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function ScrollBar_Focused(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function ScrollBar_HandleAllocated(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ScrollBar_Invalidate(AObj: TScrollBar); stdcall;
begin
  AObj.Invalidate;
end;

procedure ScrollBar_Realign(AObj: TScrollBar); stdcall;
begin
  AObj.Realign;
end;

procedure ScrollBar_Repaint(AObj: TScrollBar); stdcall;
begin
  AObj.Repaint;
end;

procedure ScrollBar_ScaleBy(AObj: TScrollBar; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure ScrollBar_SetBounds(AObj: TScrollBar; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ScrollBar_SetFocus(AObj: TScrollBar); stdcall;
begin
  AObj.SetFocus;
end;

procedure ScrollBar_Update(AObj: TScrollBar); stdcall;
begin
  AObj.Update;
end;

procedure ScrollBar_BringToFront(AObj: TScrollBar); stdcall;
begin
  AObj.BringToFront;
end;

function ScrollBar_HasParent(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ScrollBar_Hide(AObj: TScrollBar); stdcall;
begin
  AObj.Hide;
end;

function ScrollBar_Perform(AObj: TScrollBar; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ScrollBar_Refresh(AObj: TScrollBar); stdcall;
begin
  AObj.Refresh;
end;

procedure ScrollBar_SendToBack(AObj: TScrollBar); stdcall;
begin
  AObj.SendToBack;
end;

procedure ScrollBar_Show(AObj: TScrollBar); stdcall;
begin
  AObj.Show;
end;

function ScrollBar_GetTextBuf(AObj: TScrollBar; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ScrollBar_GetTextLen(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function ScrollBar_FindComponent(AObj: TScrollBar; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ScrollBar_GetNamePath(AObj: TScrollBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ScrollBar_Assign(AObj: TScrollBar; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ScrollBar_ClassName(AObj: TScrollBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ScrollBar_Equals(AObj: TScrollBar; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ScrollBar_GetHashCode(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ScrollBar_ToString(AObj: TScrollBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ScrollBar_GetAlign(AObj: TScrollBar): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ScrollBar_SetAlign(AObj: TScrollBar; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ScrollBar_GetAnchors(AObj: TScrollBar): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ScrollBar_SetAnchors(AObj: TScrollBar; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ScrollBar_GetBiDiMode(AObj: TScrollBar): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ScrollBar_SetBiDiMode(AObj: TScrollBar; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function ScrollBar_GetDoubleBuffered(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ScrollBar_SetDoubleBuffered(AObj: TScrollBar; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function ScrollBar_GetEnabled(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ScrollBar_SetEnabled(AObj: TScrollBar; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ScrollBar_GetKind(AObj: TScrollBar): TScrollBarKind; stdcall;
begin
  Result :=  AObj.Kind;
end;

procedure ScrollBar_SetKind(AObj: TScrollBar; AValue: TScrollBarKind); stdcall;
begin
  AObj.Kind := AValue;
end;

function ScrollBar_GetLargeChange(AObj: TScrollBar): TScrollBarInc; stdcall;
begin
  Result :=  AObj.LargeChange;
end;

procedure ScrollBar_SetLargeChange(AObj: TScrollBar; AValue: TScrollBarInc); stdcall;
begin
  AObj.LargeChange := AValue;
end;

function ScrollBar_GetMax(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.Max;
end;

procedure ScrollBar_SetMax(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.Max := AValue;
end;

function ScrollBar_GetMin(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.Min;
end;

procedure ScrollBar_SetMin(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.Min := AValue;
end;

function ScrollBar_GetPageSize(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.PageSize;
end;

procedure ScrollBar_SetPageSize(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.PageSize := AValue;
end;

function ScrollBar_GetParentCtl3D(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure ScrollBar_SetParentCtl3D(AObj: TScrollBar; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function ScrollBar_GetParentDoubleBuffered(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure ScrollBar_SetParentDoubleBuffered(AObj: TScrollBar; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function ScrollBar_GetParentShowHint(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ScrollBar_SetParentShowHint(AObj: TScrollBar; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ScrollBar_GetPopupMenu(AObj: TScrollBar): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ScrollBar_SetPopupMenu(AObj: TScrollBar; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ScrollBar_GetPosition(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.Position;
end;

procedure ScrollBar_SetPosition(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.Position := AValue;
end;

function ScrollBar_GetShowHint(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ScrollBar_SetShowHint(AObj: TScrollBar; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ScrollBar_GetSmallChange(AObj: TScrollBar): TScrollBarInc; stdcall;
begin
  Result :=  AObj.SmallChange;
end;

procedure ScrollBar_SetSmallChange(AObj: TScrollBar; AValue: TScrollBarInc); stdcall;
begin
  AObj.SmallChange := AValue;
end;

function ScrollBar_GetTabOrder(AObj: TScrollBar): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure ScrollBar_SetTabOrder(AObj: TScrollBar; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function ScrollBar_GetTabStop(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure ScrollBar_SetTabStop(AObj: TScrollBar; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function ScrollBar_GetVisible(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ScrollBar_SetVisible(AObj: TScrollBar; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function ScrollBar_GetStyleElements(AObj: TScrollBar): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ScrollBar_SetStyleElements(AObj: TScrollBar; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure ScrollBar_SetOnChange(AObj: TScrollBar; AEventId: NativeUInt); stdcall;
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

procedure ScrollBar_SetOnEnter(AObj: TScrollBar; AEventId: NativeUInt); stdcall;
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

procedure ScrollBar_SetOnExit(AObj: TScrollBar; AEventId: NativeUInt); stdcall;
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

procedure ScrollBar_SetOnKeyDown(AObj: TScrollBar; AEventId: NativeUInt); stdcall;
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

procedure ScrollBar_SetOnKeyPress(AObj: TScrollBar; AEventId: NativeUInt); stdcall;
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

procedure ScrollBar_SetOnKeyUp(AObj: TScrollBar; AEventId: NativeUInt); stdcall;
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

procedure ScrollBar_SetOnMouseEnter(AObj: TScrollBar; AEventId: NativeUInt); stdcall;
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

procedure ScrollBar_SetOnMouseLeave(AObj: TScrollBar; AEventId: NativeUInt); stdcall;
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

function ScrollBar_GetBrush(AObj: TScrollBar): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function ScrollBar_GetControlCount(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function ScrollBar_GetHandle(AObj: TScrollBar): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ScrollBar_GetParentWindow(AObj: TScrollBar): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ScrollBar_SetParentWindow(AObj: TScrollBar; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function ScrollBar_GetAction(AObj: TScrollBar): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ScrollBar_SetAction(AObj: TScrollBar; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure ScrollBar_GetBoundsRect(AObj: TScrollBar; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ScrollBar_SetBoundsRect(AObj: TScrollBar; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ScrollBar_GetClientHeight(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ScrollBar_SetClientHeight(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ScrollBar_GetClientRect(AObj: TScrollBar; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ScrollBar_GetClientWidth(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ScrollBar_SetClientWidth(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ScrollBar_GetExplicitLeft(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ScrollBar_GetExplicitTop(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ScrollBar_GetExplicitWidth(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ScrollBar_GetExplicitHeight(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ScrollBar_GetParent(AObj: TScrollBar): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ScrollBar_SetParent(AObj: TScrollBar; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ScrollBar_GetAlignWithMargins(AObj: TScrollBar): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ScrollBar_SetAlignWithMargins(AObj: TScrollBar; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ScrollBar_GetLeft(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ScrollBar_SetLeft(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ScrollBar_GetTop(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ScrollBar_SetTop(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ScrollBar_GetWidth(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ScrollBar_SetWidth(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ScrollBar_GetHeight(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ScrollBar_SetHeight(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ScrollBar_GetCursor(AObj: TScrollBar): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ScrollBar_SetCursor(AObj: TScrollBar; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ScrollBar_GetHint(AObj: TScrollBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ScrollBar_SetHint(AObj: TScrollBar; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ScrollBar_GetMargins(AObj: TScrollBar): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ScrollBar_SetMargins(AObj: TScrollBar; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ScrollBar_GetCustomHint(AObj: TScrollBar): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ScrollBar_SetCustomHint(AObj: TScrollBar; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ScrollBar_GetComponentCount(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ScrollBar_GetComponentIndex(AObj: TScrollBar): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ScrollBar_SetComponentIndex(AObj: TScrollBar; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ScrollBar_GetOwner(AObj: TScrollBar): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ScrollBar_GetName(AObj: TScrollBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ScrollBar_SetName(AObj: TScrollBar; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ScrollBar_GetTag(AObj: TScrollBar): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ScrollBar_SetTag(AObj: TScrollBar; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ScrollBar_GetControls(AObj: TScrollBar; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ScrollBar_GetComponents(AObj: TScrollBar; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ScrollBar_Create,
  ScrollBar_Free,
  ScrollBar_SetParams,
  ScrollBar_CanFocus,
  ScrollBar_FlipChildren,
  ScrollBar_Focused,
  ScrollBar_HandleAllocated,
  ScrollBar_Invalidate,
  ScrollBar_Realign,
  ScrollBar_Repaint,
  ScrollBar_ScaleBy,
  ScrollBar_SetBounds,
  ScrollBar_SetFocus,
  ScrollBar_Update,
  ScrollBar_BringToFront,
  ScrollBar_HasParent,
  ScrollBar_Hide,
  ScrollBar_Perform,
  ScrollBar_Refresh,
  ScrollBar_SendToBack,
  ScrollBar_Show,
  ScrollBar_GetTextBuf,
  ScrollBar_GetTextLen,
  ScrollBar_FindComponent,
  ScrollBar_GetNamePath,
  ScrollBar_Assign,
  ScrollBar_ClassName,
  ScrollBar_Equals,
  ScrollBar_GetHashCode,
  ScrollBar_ToString,
  ScrollBar_GetAlign,
  ScrollBar_SetAlign,
  ScrollBar_GetAnchors,
  ScrollBar_SetAnchors,
  ScrollBar_GetBiDiMode,
  ScrollBar_SetBiDiMode,
  ScrollBar_GetDoubleBuffered,
  ScrollBar_SetDoubleBuffered,
  ScrollBar_GetEnabled,
  ScrollBar_SetEnabled,
  ScrollBar_GetKind,
  ScrollBar_SetKind,
  ScrollBar_GetLargeChange,
  ScrollBar_SetLargeChange,
  ScrollBar_GetMax,
  ScrollBar_SetMax,
  ScrollBar_GetMin,
  ScrollBar_SetMin,
  ScrollBar_GetPageSize,
  ScrollBar_SetPageSize,
  ScrollBar_GetParentCtl3D,
  ScrollBar_SetParentCtl3D,
  ScrollBar_GetParentDoubleBuffered,
  ScrollBar_SetParentDoubleBuffered,
  ScrollBar_GetParentShowHint,
  ScrollBar_SetParentShowHint,
  ScrollBar_GetPopupMenu,
  ScrollBar_SetPopupMenu,
  ScrollBar_GetPosition,
  ScrollBar_SetPosition,
  ScrollBar_GetShowHint,
  ScrollBar_SetShowHint,
  ScrollBar_GetSmallChange,
  ScrollBar_SetSmallChange,
  ScrollBar_GetTabOrder,
  ScrollBar_SetTabOrder,
  ScrollBar_GetTabStop,
  ScrollBar_SetTabStop,
  ScrollBar_GetVisible,
  ScrollBar_SetVisible,
  ScrollBar_GetStyleElements,
  ScrollBar_SetStyleElements,
  ScrollBar_SetOnChange,
  ScrollBar_SetOnEnter,
  ScrollBar_SetOnExit,
  ScrollBar_SetOnKeyDown,
  ScrollBar_SetOnKeyPress,
  ScrollBar_SetOnKeyUp,
  ScrollBar_SetOnMouseEnter,
  ScrollBar_SetOnMouseLeave,
  ScrollBar_GetBrush,
  ScrollBar_GetControlCount,
  ScrollBar_GetHandle,
  ScrollBar_GetParentWindow,
  ScrollBar_SetParentWindow,
  ScrollBar_GetAction,
  ScrollBar_SetAction,
  ScrollBar_GetBoundsRect,
  ScrollBar_SetBoundsRect,
  ScrollBar_GetClientHeight,
  ScrollBar_SetClientHeight,
  ScrollBar_GetClientRect,
  ScrollBar_GetClientWidth,
  ScrollBar_SetClientWidth,
  ScrollBar_GetExplicitLeft,
  ScrollBar_GetExplicitTop,
  ScrollBar_GetExplicitWidth,
  ScrollBar_GetExplicitHeight,
  ScrollBar_GetParent,
  ScrollBar_SetParent,
  ScrollBar_GetAlignWithMargins,
  ScrollBar_SetAlignWithMargins,
  ScrollBar_GetLeft,
  ScrollBar_SetLeft,
  ScrollBar_GetTop,
  ScrollBar_SetTop,
  ScrollBar_GetWidth,
  ScrollBar_SetWidth,
  ScrollBar_GetHeight,
  ScrollBar_SetHeight,
  ScrollBar_GetCursor,
  ScrollBar_SetCursor,
  ScrollBar_GetHint,
  ScrollBar_SetHint,
  ScrollBar_GetMargins,
  ScrollBar_SetMargins,
  ScrollBar_GetCustomHint,
  ScrollBar_SetCustomHint,
  ScrollBar_GetComponentCount,
  ScrollBar_GetComponentIndex,
  ScrollBar_SetComponentIndex,
  ScrollBar_GetOwner,
  ScrollBar_GetName,
  ScrollBar_SetName,
  ScrollBar_GetTag,
  ScrollBar_SetTag,
  ScrollBar_GetControls,
  ScrollBar_GetComponents;

