
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function TabSheet_Create(AOwner: TComponent): TTabSheet; stdcall;
begin
  Result :=  TTabSheet.Create(AOwner);
end;

procedure TabSheet_Free(AObj: TTabSheet); stdcall;
begin
  AObj.Free;
end;

function TabSheet_CanFocus(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure TabSheet_FlipChildren(AObj: TTabSheet; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function TabSheet_Focused(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function TabSheet_HandleAllocated(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure TabSheet_Invalidate(AObj: TTabSheet); stdcall;
begin
  AObj.Invalidate;
end;

procedure TabSheet_Realign(AObj: TTabSheet); stdcall;
begin
  AObj.Realign;
end;

procedure TabSheet_Repaint(AObj: TTabSheet); stdcall;
begin
  AObj.Repaint;
end;

procedure TabSheet_ScaleBy(AObj: TTabSheet; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure TabSheet_SetBounds(AObj: TTabSheet; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure TabSheet_SetFocus(AObj: TTabSheet); stdcall;
begin
  AObj.SetFocus;
end;

procedure TabSheet_Update(AObj: TTabSheet); stdcall;
begin
  AObj.Update;
end;

procedure TabSheet_BringToFront(AObj: TTabSheet); stdcall;
begin
  AObj.BringToFront;
end;

function TabSheet_HasParent(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure TabSheet_Hide(AObj: TTabSheet); stdcall;
begin
  AObj.Hide;
end;

function TabSheet_Perform(AObj: TTabSheet; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure TabSheet_Refresh(AObj: TTabSheet); stdcall;
begin
  AObj.Refresh;
end;

procedure TabSheet_SendToBack(AObj: TTabSheet); stdcall;
begin
  AObj.SendToBack;
end;

procedure TabSheet_Show(AObj: TTabSheet); stdcall;
begin
  AObj.Show;
end;

function TabSheet_GetTextBuf(AObj: TTabSheet; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function TabSheet_GetTextLen(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function TabSheet_FindComponent(AObj: TTabSheet; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function TabSheet_GetNamePath(AObj: TTabSheet): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TabSheet_Assign(AObj: TTabSheet; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function TabSheet_ClassName(AObj: TTabSheet): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TabSheet_Equals(AObj: TTabSheet; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function TabSheet_GetHashCode(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function TabSheet_ToString(AObj: TTabSheet): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function TabSheet_GetPageControl(AObj: TTabSheet): TPageControl; stdcall;
begin
  Result :=  AObj.PageControl;
end;

procedure TabSheet_SetPageControl(AObj: TTabSheet; AValue: TPageControl); stdcall;
begin
  AObj.PageControl := AValue;
end;

function TabSheet_GetTabIndex(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.TabIndex;
end;

function TabSheet_GetBorderWidth(AObj: TTabSheet): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure TabSheet_SetBorderWidth(AObj: TTabSheet; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function TabSheet_GetCaption(AObj: TTabSheet): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure TabSheet_SetCaption(AObj: TTabSheet; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function TabSheet_GetDoubleBuffered(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure TabSheet_SetDoubleBuffered(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function TabSheet_GetEnabled(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure TabSheet_SetEnabled(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function TabSheet_GetFont(AObj: TTabSheet): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure TabSheet_SetFont(AObj: TTabSheet; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function TabSheet_GetHeight(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure TabSheet_SetHeight(AObj: TTabSheet; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function TabSheet_GetHighlighted(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.Highlighted;
end;

procedure TabSheet_SetHighlighted(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.Highlighted := AValue;
end;

function TabSheet_GetImageIndex(AObj: TTabSheet): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure TabSheet_SetImageIndex(AObj: TTabSheet; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function TabSheet_GetLeft(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure TabSheet_SetLeft(AObj: TTabSheet; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function TabSheet_GetPageIndex(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.PageIndex;
end;

procedure TabSheet_SetPageIndex(AObj: TTabSheet; AValue: Integer); stdcall;
begin
  AObj.PageIndex := AValue;
end;

function TabSheet_GetParentDoubleBuffered(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure TabSheet_SetParentDoubleBuffered(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function TabSheet_GetParentFont(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure TabSheet_SetParentFont(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function TabSheet_GetParentShowHint(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure TabSheet_SetParentShowHint(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function TabSheet_GetPopupMenu(AObj: TTabSheet): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure TabSheet_SetPopupMenu(AObj: TTabSheet; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function TabSheet_GetShowHint(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure TabSheet_SetShowHint(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function TabSheet_GetTabVisible(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.TabVisible;
end;

procedure TabSheet_SetTabVisible(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.TabVisible := AValue;
end;

function TabSheet_GetTop(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure TabSheet_SetTop(AObj: TTabSheet; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function TabSheet_GetVisible(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure TabSheet_SetVisible(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function TabSheet_GetWidth(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure TabSheet_SetWidth(AObj: TTabSheet; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

procedure TabSheet_SetOnEnter(AObj: TTabSheet; AEventId: NativeUInt); stdcall;
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

procedure TabSheet_SetOnExit(AObj: TTabSheet; AEventId: NativeUInt); stdcall;
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

procedure TabSheet_SetOnHide(AObj: TTabSheet; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnHide := nil;
    TEventClass.Remove(AObj, geHide);
    Exit;
  end;
  AObj.OnHide := TEventClass.OnHide;
  TEventClass.Add(AObj, geHide, AEventId);
end;

procedure TabSheet_SetOnMouseDown(AObj: TTabSheet; AEventId: NativeUInt); stdcall;
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

procedure TabSheet_SetOnMouseEnter(AObj: TTabSheet; AEventId: NativeUInt); stdcall;
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

procedure TabSheet_SetOnMouseLeave(AObj: TTabSheet; AEventId: NativeUInt); stdcall;
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

procedure TabSheet_SetOnMouseMove(AObj: TTabSheet; AEventId: NativeUInt); stdcall;
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

procedure TabSheet_SetOnMouseUp(AObj: TTabSheet; AEventId: NativeUInt); stdcall;
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

procedure TabSheet_SetOnResize(AObj: TTabSheet; AEventId: NativeUInt); stdcall;
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

procedure TabSheet_SetOnShow(AObj: TTabSheet; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnShow := nil;
    TEventClass.Remove(AObj, geShow);
    Exit;
  end;
  AObj.OnShow := TEventClass.OnShow;
  TEventClass.Add(AObj, geShow, AEventId);
end;

function TabSheet_GetBrush(AObj: TTabSheet): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function TabSheet_GetControlCount(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function TabSheet_GetHandle(AObj: TTabSheet): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function TabSheet_GetParentWindow(AObj: TTabSheet): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure TabSheet_SetParentWindow(AObj: TTabSheet; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function TabSheet_GetTabOrder(AObj: TTabSheet): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure TabSheet_SetTabOrder(AObj: TTabSheet; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function TabSheet_GetTabStop(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure TabSheet_SetTabStop(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function TabSheet_GetAction(AObj: TTabSheet): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure TabSheet_SetAction(AObj: TTabSheet; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function TabSheet_GetAlign(AObj: TTabSheet): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure TabSheet_SetAlign(AObj: TTabSheet; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function TabSheet_GetAnchors(AObj: TTabSheet): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure TabSheet_SetAnchors(AObj: TTabSheet; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function TabSheet_GetBiDiMode(AObj: TTabSheet): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure TabSheet_SetBiDiMode(AObj: TTabSheet; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure TabSheet_GetBoundsRect(AObj: TTabSheet; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure TabSheet_SetBoundsRect(AObj: TTabSheet; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function TabSheet_GetClientHeight(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure TabSheet_SetClientHeight(AObj: TTabSheet; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure TabSheet_GetClientRect(AObj: TTabSheet; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function TabSheet_GetClientWidth(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure TabSheet_SetClientWidth(AObj: TTabSheet; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function TabSheet_GetExplicitLeft(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function TabSheet_GetExplicitTop(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function TabSheet_GetExplicitWidth(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function TabSheet_GetExplicitHeight(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function TabSheet_GetParent(AObj: TTabSheet): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure TabSheet_SetParent(AObj: TTabSheet; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function TabSheet_GetStyleElements(AObj: TTabSheet): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure TabSheet_SetStyleElements(AObj: TTabSheet; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function TabSheet_GetAlignWithMargins(AObj: TTabSheet): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure TabSheet_SetAlignWithMargins(AObj: TTabSheet; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function TabSheet_GetCursor(AObj: TTabSheet): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure TabSheet_SetCursor(AObj: TTabSheet; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function TabSheet_GetHint(AObj: TTabSheet): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure TabSheet_SetHint(AObj: TTabSheet; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function TabSheet_GetMargins(AObj: TTabSheet): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure TabSheet_SetMargins(AObj: TTabSheet; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function TabSheet_GetCustomHint(AObj: TTabSheet): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure TabSheet_SetCustomHint(AObj: TTabSheet; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function TabSheet_GetComponentCount(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function TabSheet_GetComponentIndex(AObj: TTabSheet): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure TabSheet_SetComponentIndex(AObj: TTabSheet; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function TabSheet_GetOwner(AObj: TTabSheet): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function TabSheet_GetName(AObj: TTabSheet): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TabSheet_SetName(AObj: TTabSheet; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function TabSheet_GetTag(AObj: TTabSheet): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure TabSheet_SetTag(AObj: TTabSheet; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function TabSheet_GetControls(AObj: TTabSheet; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function TabSheet_GetComponents(AObj: TTabSheet; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  TabSheet_Create,
  TabSheet_Free,
  TabSheet_CanFocus,
  TabSheet_FlipChildren,
  TabSheet_Focused,
  TabSheet_HandleAllocated,
  TabSheet_Invalidate,
  TabSheet_Realign,
  TabSheet_Repaint,
  TabSheet_ScaleBy,
  TabSheet_SetBounds,
  TabSheet_SetFocus,
  TabSheet_Update,
  TabSheet_BringToFront,
  TabSheet_HasParent,
  TabSheet_Hide,
  TabSheet_Perform,
  TabSheet_Refresh,
  TabSheet_SendToBack,
  TabSheet_Show,
  TabSheet_GetTextBuf,
  TabSheet_GetTextLen,
  TabSheet_FindComponent,
  TabSheet_GetNamePath,
  TabSheet_Assign,
  TabSheet_ClassName,
  TabSheet_Equals,
  TabSheet_GetHashCode,
  TabSheet_ToString,
  TabSheet_GetPageControl,
  TabSheet_SetPageControl,
  TabSheet_GetTabIndex,
  TabSheet_GetBorderWidth,
  TabSheet_SetBorderWidth,
  TabSheet_GetCaption,
  TabSheet_SetCaption,
  TabSheet_GetDoubleBuffered,
  TabSheet_SetDoubleBuffered,
  TabSheet_GetEnabled,
  TabSheet_SetEnabled,
  TabSheet_GetFont,
  TabSheet_SetFont,
  TabSheet_GetHeight,
  TabSheet_SetHeight,
  TabSheet_GetHighlighted,
  TabSheet_SetHighlighted,
  TabSheet_GetImageIndex,
  TabSheet_SetImageIndex,
  TabSheet_GetLeft,
  TabSheet_SetLeft,
  TabSheet_GetPageIndex,
  TabSheet_SetPageIndex,
  TabSheet_GetParentDoubleBuffered,
  TabSheet_SetParentDoubleBuffered,
  TabSheet_GetParentFont,
  TabSheet_SetParentFont,
  TabSheet_GetParentShowHint,
  TabSheet_SetParentShowHint,
  TabSheet_GetPopupMenu,
  TabSheet_SetPopupMenu,
  TabSheet_GetShowHint,
  TabSheet_SetShowHint,
  TabSheet_GetTabVisible,
  TabSheet_SetTabVisible,
  TabSheet_GetTop,
  TabSheet_SetTop,
  TabSheet_GetVisible,
  TabSheet_SetVisible,
  TabSheet_GetWidth,
  TabSheet_SetWidth,
  TabSheet_SetOnEnter,
  TabSheet_SetOnExit,
  TabSheet_SetOnHide,
  TabSheet_SetOnMouseDown,
  TabSheet_SetOnMouseEnter,
  TabSheet_SetOnMouseLeave,
  TabSheet_SetOnMouseMove,
  TabSheet_SetOnMouseUp,
  TabSheet_SetOnResize,
  TabSheet_SetOnShow,
  TabSheet_GetBrush,
  TabSheet_GetControlCount,
  TabSheet_GetHandle,
  TabSheet_GetParentWindow,
  TabSheet_SetParentWindow,
  TabSheet_GetTabOrder,
  TabSheet_SetTabOrder,
  TabSheet_GetTabStop,
  TabSheet_SetTabStop,
  TabSheet_GetAction,
  TabSheet_SetAction,
  TabSheet_GetAlign,
  TabSheet_SetAlign,
  TabSheet_GetAnchors,
  TabSheet_SetAnchors,
  TabSheet_GetBiDiMode,
  TabSheet_SetBiDiMode,
  TabSheet_GetBoundsRect,
  TabSheet_SetBoundsRect,
  TabSheet_GetClientHeight,
  TabSheet_SetClientHeight,
  TabSheet_GetClientRect,
  TabSheet_GetClientWidth,
  TabSheet_SetClientWidth,
  TabSheet_GetExplicitLeft,
  TabSheet_GetExplicitTop,
  TabSheet_GetExplicitWidth,
  TabSheet_GetExplicitHeight,
  TabSheet_GetParent,
  TabSheet_SetParent,
  TabSheet_GetStyleElements,
  TabSheet_SetStyleElements,
  TabSheet_GetAlignWithMargins,
  TabSheet_SetAlignWithMargins,
  TabSheet_GetCursor,
  TabSheet_SetCursor,
  TabSheet_GetHint,
  TabSheet_SetHint,
  TabSheet_GetMargins,
  TabSheet_SetMargins,
  TabSheet_GetCustomHint,
  TabSheet_SetCustomHint,
  TabSheet_GetComponentCount,
  TabSheet_GetComponentIndex,
  TabSheet_SetComponentIndex,
  TabSheet_GetOwner,
  TabSheet_GetName,
  TabSheet_SetName,
  TabSheet_GetTag,
  TabSheet_SetTag,
  TabSheet_GetControls,
  TabSheet_GetComponents;

