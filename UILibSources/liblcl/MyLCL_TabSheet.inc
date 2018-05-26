
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function TabSheet_Create(AOwner: TComponent): TTabSheet; extdecl;
begin
  Result :=  TTabSheet.Create(AOwner);
end;

procedure TabSheet_Free(AObj: TTabSheet); extdecl;
begin
  AObj.Free;
end;

function TabSheet_CanFocus(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure TabSheet_FlipChildren(AObj: TTabSheet; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function TabSheet_Focused(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure TabSheet_Invalidate(AObj: TTabSheet); extdecl;
begin
  AObj.Invalidate;
end;

procedure TabSheet_Realign(AObj: TTabSheet); extdecl;
begin
  AObj.Realign;
end;

procedure TabSheet_Repaint(AObj: TTabSheet); extdecl;
begin
  AObj.Repaint;
end;

procedure TabSheet_ScaleBy(AObj: TTabSheet; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure TabSheet_SetBounds(AObj: TTabSheet; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure TabSheet_SetFocus(AObj: TTabSheet); extdecl;
begin
  AObj.SetFocus;
end;

procedure TabSheet_Update(AObj: TTabSheet); extdecl;
begin
  AObj.Update;
end;

procedure TabSheet_BringToFront(AObj: TTabSheet); extdecl;
begin
  AObj.BringToFront;
end;

function TabSheet_HasParent(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure TabSheet_Hide(AObj: TTabSheet); extdecl;
begin
  AObj.Hide;
end;

function TabSheet_Perform(AObj: TTabSheet; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure TabSheet_Refresh(AObj: TTabSheet); extdecl;
begin
  AObj.Refresh;
end;

procedure TabSheet_SendToBack(AObj: TTabSheet); extdecl;
begin
  AObj.SendToBack;
end;

procedure TabSheet_Show(AObj: TTabSheet); extdecl;
begin
  AObj.Show;
end;

function TabSheet_GetTextBuf(AObj: TTabSheet; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function TabSheet_GetTextLen(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function TabSheet_FindComponent(AObj: TTabSheet; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function TabSheet_GetNamePath(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TabSheet_Assign(AObj: TTabSheet; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function TabSheet_ClassName(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function TabSheet_Equals(AObj: TTabSheet; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TabSheet_GetHashCode(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TabSheet_ToString(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function TabSheet_GetPageControl(AObj: TTabSheet): TPageControl; extdecl;
begin
  Result :=  AObj.PageControl;
end;

procedure TabSheet_SetPageControl(AObj: TTabSheet; AValue: TPageControl); extdecl;
begin
  AObj.PageControl := AValue;
end;

function TabSheet_GetTabIndex(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.TabIndex;
end;

function TabSheet_GetBorderWidth(AObj: TTabSheet): TBorderWidth; extdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure TabSheet_SetBorderWidth(AObj: TTabSheet; AValue: TBorderWidth); extdecl;
begin
  AObj.BorderWidth := AValue;
end;

function TabSheet_GetCaption(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure TabSheet_SetCaption(AObj: TTabSheet; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function TabSheet_GetDoubleBuffered(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure TabSheet_SetDoubleBuffered(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function TabSheet_GetEnabled(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure TabSheet_SetEnabled(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function TabSheet_GetFont(AObj: TTabSheet): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure TabSheet_SetFont(AObj: TTabSheet; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function TabSheet_GetHeight(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure TabSheet_SetHeight(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function TabSheet_GetImageIndex(AObj: TTabSheet): TImageIndex; extdecl;
begin
  Result :=  AObj.ImageIndex;
end;

procedure TabSheet_SetImageIndex(AObj: TTabSheet; AValue: TImageIndex); extdecl;
begin
  AObj.ImageIndex := AValue;
end;

function TabSheet_GetLeft(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure TabSheet_SetLeft(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function TabSheet_GetPageIndex(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.PageIndex;
end;

procedure TabSheet_SetPageIndex(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.PageIndex := AValue;
end;

function TabSheet_GetParentFont(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure TabSheet_SetParentFont(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function TabSheet_GetParentShowHint(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure TabSheet_SetParentShowHint(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function TabSheet_GetPopupMenu(AObj: TTabSheet): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure TabSheet_SetPopupMenu(AObj: TTabSheet; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function TabSheet_GetShowHint(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure TabSheet_SetShowHint(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function TabSheet_GetTabVisible(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.TabVisible;
end;

procedure TabSheet_SetTabVisible(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.TabVisible := AValue;
end;

function TabSheet_GetTop(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure TabSheet_SetTop(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function TabSheet_GetVisible(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure TabSheet_SetVisible(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function TabSheet_GetWidth(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure TabSheet_SetWidth(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

procedure TabSheet_SetOnEnter(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnEnter := nil;
    EventClass.Remove(AObj, geEnter);
    Exit;
  end;
  AObj.OnEnter := @EventClass.OnEnter;
  EventClass.Add(AObj, geEnter, AEventId);
end;

procedure TabSheet_SetOnExit(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnExit := nil;
    EventClass.Remove(AObj, geExit);
    Exit;
  end;
  AObj.OnExit := @EventClass.OnExit;
  EventClass.Add(AObj, geExit, AEventId);
end;

procedure TabSheet_SetOnHide(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnHide := nil;
    EventClass.Remove(AObj, geHide);
    Exit;
  end;
  AObj.OnHide := @EventClass.OnHide;
  EventClass.Add(AObj, geHide, AEventId);
end;

procedure TabSheet_SetOnMouseDown(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseDown := nil;
    EventClass.Remove(AObj, geMouseDown);
    Exit;
  end;
  AObj.OnMouseDown := @EventClass.OnMouseDown;
  EventClass.Add(AObj, geMouseDown, AEventId);
end;

procedure TabSheet_SetOnMouseEnter(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseEnter := nil;
    EventClass.Remove(AObj, geMouseEnter);
    Exit;
  end;
  AObj.OnMouseEnter := @EventClass.OnMouseEnter;
  EventClass.Add(AObj, geMouseEnter, AEventId);
end;

procedure TabSheet_SetOnMouseLeave(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseLeave := nil;
    EventClass.Remove(AObj, geMouseLeave);
    Exit;
  end;
  AObj.OnMouseLeave := @EventClass.OnMouseLeave;
  EventClass.Add(AObj, geMouseLeave, AEventId);
end;

procedure TabSheet_SetOnMouseMove(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseMove := nil;
    EventClass.Remove(AObj, geMouseMove);
    Exit;
  end;
  AObj.OnMouseMove := @EventClass.OnMouseMove;
  EventClass.Add(AObj, geMouseMove, AEventId);
end;

procedure TabSheet_SetOnMouseUp(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseUp := nil;
    EventClass.Remove(AObj, geMouseUp);
    Exit;
  end;
  AObj.OnMouseUp := @EventClass.OnMouseUp;
  EventClass.Add(AObj, geMouseUp, AEventId);
end;

procedure TabSheet_SetOnResize(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnResize := nil;
    EventClass.Remove(AObj, geResize);
    Exit;
  end;
  AObj.OnResize := @EventClass.OnResize;
  EventClass.Add(AObj, geResize, AEventId);
end;

procedure TabSheet_SetOnShow(AObj: TTabSheet; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnShow := nil;
    EventClass.Remove(AObj, geShow);
    Exit;
  end;
  AObj.OnShow := @EventClass.OnShow;
  EventClass.Add(AObj, geShow, AEventId);
end;

function TabSheet_GetBrush(AObj: TTabSheet): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function TabSheet_GetControlCount(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function TabSheet_GetHandle(AObj: TTabSheet): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function TabSheet_GetParentWindow(AObj: TTabSheet): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure TabSheet_SetParentWindow(AObj: TTabSheet; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function TabSheet_GetTabOrder(AObj: TTabSheet): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure TabSheet_SetTabOrder(AObj: TTabSheet; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function TabSheet_GetTabStop(AObj: TTabSheet): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure TabSheet_SetTabStop(AObj: TTabSheet; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function TabSheet_GetAction(AObj: TTabSheet): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure TabSheet_SetAction(AObj: TTabSheet; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function TabSheet_GetAlign(AObj: TTabSheet): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure TabSheet_SetAlign(AObj: TTabSheet; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function TabSheet_GetAnchors(AObj: TTabSheet): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure TabSheet_SetAnchors(AObj: TTabSheet; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function TabSheet_GetBiDiMode(AObj: TTabSheet): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure TabSheet_SetBiDiMode(AObj: TTabSheet; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure TabSheet_GetBoundsRect(AObj: TTabSheet; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure TabSheet_SetBoundsRect(AObj: TTabSheet; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function TabSheet_GetClientHeight(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure TabSheet_SetClientHeight(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure TabSheet_GetClientRect(AObj: TTabSheet; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function TabSheet_GetClientWidth(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure TabSheet_SetClientWidth(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function TabSheet_GetParent(AObj: TTabSheet): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure TabSheet_SetParent(AObj: TTabSheet; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function TabSheet_GetCursor(AObj: TTabSheet): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure TabSheet_SetCursor(AObj: TTabSheet; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function TabSheet_GetHint(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure TabSheet_SetHint(AObj: TTabSheet; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function TabSheet_GetComponentCount(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function TabSheet_GetComponentIndex(AObj: TTabSheet): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure TabSheet_SetComponentIndex(AObj: TTabSheet; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function TabSheet_GetOwner(AObj: TTabSheet): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function TabSheet_GetName(AObj: TTabSheet): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TabSheet_SetName(AObj: TTabSheet; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function TabSheet_GetTag(AObj: TTabSheet): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure TabSheet_SetTag(AObj: TTabSheet; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function TabSheet_GetControls(AObj: TTabSheet; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function TabSheet_GetComponents(AObj: TTabSheet; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  TabSheet_Create,
  TabSheet_Free,
  TabSheet_CanFocus,
  TabSheet_FlipChildren,
  TabSheet_Focused,
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
  TabSheet_GetImageIndex,
  TabSheet_SetImageIndex,
  TabSheet_GetLeft,
  TabSheet_SetLeft,
  TabSheet_GetPageIndex,
  TabSheet_SetPageIndex,
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
  TabSheet_GetParent,
  TabSheet_SetParent,
  TabSheet_GetCursor,
  TabSheet_SetCursor,
  TabSheet_GetHint,
  TabSheet_SetHint,
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

