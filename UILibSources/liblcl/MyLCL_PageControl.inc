
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function PageControl_Create(AOwner: TComponent): TPageControl; extdecl;
begin
  Result :=  TPageControl.Create(AOwner);
end;

procedure PageControl_Free(AObj: TPageControl); extdecl;
begin
  AObj.Free;
end;

procedure PageControl_SelectNextPage(AObj: TPageControl; GoForward: LongBool; CheckTabVisible: LongBool); extdecl;
begin
  AObj.SelectNextPage(GoForward, CheckTabVisible);
end;

function PageControl_CanFocus(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure PageControl_FlipChildren(AObj: TPageControl; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function PageControl_Focused(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure PageControl_Invalidate(AObj: TPageControl); extdecl;
begin
  AObj.Invalidate;
end;

procedure PageControl_Realign(AObj: TPageControl); extdecl;
begin
  AObj.Realign;
end;

procedure PageControl_Repaint(AObj: TPageControl); extdecl;
begin
  AObj.Repaint;
end;

procedure PageControl_ScaleBy(AObj: TPageControl; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure PageControl_SetBounds(AObj: TPageControl; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure PageControl_SetFocus(AObj: TPageControl); extdecl;
begin
  AObj.SetFocus;
end;

procedure PageControl_Update(AObj: TPageControl); extdecl;
begin
  AObj.Update;
end;

procedure PageControl_BringToFront(AObj: TPageControl); extdecl;
begin
  AObj.BringToFront;
end;

function PageControl_HasParent(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure PageControl_Hide(AObj: TPageControl); extdecl;
begin
  AObj.Hide;
end;

function PageControl_Perform(AObj: TPageControl; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure PageControl_Refresh(AObj: TPageControl); extdecl;
begin
  AObj.Refresh;
end;

procedure PageControl_SendToBack(AObj: TPageControl); extdecl;
begin
  AObj.SendToBack;
end;

procedure PageControl_Show(AObj: TPageControl); extdecl;
begin
  AObj.Show;
end;

function PageControl_GetTextBuf(AObj: TPageControl; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function PageControl_GetTextLen(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function PageControl_FindComponent(AObj: TPageControl; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function PageControl_GetNamePath(AObj: TPageControl): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure PageControl_Assign(AObj: TPageControl; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function PageControl_ClassName(AObj: TPageControl): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function PageControl_Equals(AObj: TPageControl; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function PageControl_GetHashCode(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function PageControl_ToString(AObj: TPageControl): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function PageControl_GetActivePageIndex(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.ActivePageIndex;
end;

procedure PageControl_SetActivePageIndex(AObj: TPageControl; AValue: Integer); extdecl;
begin
  AObj.ActivePageIndex := AValue;
end;

function PageControl_GetPageCount(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.PageCount;
end;

function PageControl_GetAlign(AObj: TPageControl): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure PageControl_SetAlign(AObj: TPageControl; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function PageControl_GetAnchors(AObj: TPageControl): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure PageControl_SetAnchors(AObj: TPageControl; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function PageControl_GetBiDiMode(AObj: TPageControl): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure PageControl_SetBiDiMode(AObj: TPageControl; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function PageControl_GetDoubleBuffered(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure PageControl_SetDoubleBuffered(AObj: TPageControl; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function PageControl_GetEnabled(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure PageControl_SetEnabled(AObj: TPageControl; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function PageControl_GetFont(AObj: TPageControl): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure PageControl_SetFont(AObj: TPageControl; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function PageControl_GetImages(AObj: TPageControl): TCustomImageList; extdecl;
begin
  Result :=  AObj.Images;
end;

procedure PageControl_SetImages(AObj: TPageControl; AValue: TCustomImageList); extdecl;
begin
  AObj.Images := AValue;
end;

function PageControl_GetMultiLine(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.MultiLine;
end;

procedure PageControl_SetMultiLine(AObj: TPageControl; AValue: LongBool); extdecl;
begin
  AObj.MultiLine := AValue;
end;

function PageControl_GetParentFont(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure PageControl_SetParentFont(AObj: TPageControl; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function PageControl_GetParentShowHint(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure PageControl_SetParentShowHint(AObj: TPageControl; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function PageControl_GetPopupMenu(AObj: TPageControl): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure PageControl_SetPopupMenu(AObj: TPageControl; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function PageControl_GetShowHint(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure PageControl_SetShowHint(AObj: TPageControl; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function PageControl_GetTabIndex(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.TabIndex;
end;

procedure PageControl_SetTabIndex(AObj: TPageControl; AValue: Integer); extdecl;
begin
  AObj.TabIndex := AValue;
end;

function PageControl_GetTabOrder(AObj: TPageControl): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure PageControl_SetTabOrder(AObj: TPageControl; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function PageControl_GetTabPosition(AObj: TPageControl): TTabPosition; extdecl;
begin
  Result :=  AObj.TabPosition;
end;

procedure PageControl_SetTabPosition(AObj: TPageControl; AValue: TTabPosition); extdecl;
begin
  AObj.TabPosition := AValue;
end;

function PageControl_GetTabStop(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure PageControl_SetTabStop(AObj: TPageControl; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function PageControl_GetVisible(AObj: TPageControl): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure PageControl_SetVisible(AObj: TPageControl; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure PageControl_SetOnChange(AObj: TPageControl; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    EventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := @EventClass.OnChange;
  EventClass.Add(AObj, geChange, AEventId);
end;

procedure PageControl_SetOnEnter(AObj: TPageControl; AEventId: NativeUInt); extdecl;
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

procedure PageControl_SetOnExit(AObj: TPageControl; AEventId: NativeUInt); extdecl;
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

procedure PageControl_SetOnMouseDown(AObj: TPageControl; AEventId: NativeUInt); extdecl;
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

procedure PageControl_SetOnMouseEnter(AObj: TPageControl; AEventId: NativeUInt); extdecl;
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

procedure PageControl_SetOnMouseLeave(AObj: TPageControl; AEventId: NativeUInt); extdecl;
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

procedure PageControl_SetOnMouseMove(AObj: TPageControl; AEventId: NativeUInt); extdecl;
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

procedure PageControl_SetOnMouseUp(AObj: TPageControl; AEventId: NativeUInt); extdecl;
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

procedure PageControl_SetOnResize(AObj: TPageControl; AEventId: NativeUInt); extdecl;
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

function PageControl_GetBrush(AObj: TPageControl): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function PageControl_GetControlCount(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function PageControl_GetHandle(AObj: TPageControl): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function PageControl_GetParentWindow(AObj: TPageControl): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure PageControl_SetParentWindow(AObj: TPageControl; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function PageControl_GetAction(AObj: TPageControl): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure PageControl_SetAction(AObj: TPageControl; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure PageControl_GetBoundsRect(AObj: TPageControl; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure PageControl_SetBoundsRect(AObj: TPageControl; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function PageControl_GetClientHeight(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure PageControl_SetClientHeight(AObj: TPageControl; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure PageControl_GetClientRect(AObj: TPageControl; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function PageControl_GetClientWidth(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure PageControl_SetClientWidth(AObj: TPageControl; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function PageControl_GetParent(AObj: TPageControl): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure PageControl_SetParent(AObj: TPageControl; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function PageControl_GetLeft(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure PageControl_SetLeft(AObj: TPageControl; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function PageControl_GetTop(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure PageControl_SetTop(AObj: TPageControl; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function PageControl_GetWidth(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure PageControl_SetWidth(AObj: TPageControl; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function PageControl_GetHeight(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure PageControl_SetHeight(AObj: TPageControl; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function PageControl_GetCursor(AObj: TPageControl): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure PageControl_SetCursor(AObj: TPageControl; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function PageControl_GetHint(AObj: TPageControl): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure PageControl_SetHint(AObj: TPageControl; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function PageControl_GetComponentCount(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function PageControl_GetComponentIndex(AObj: TPageControl): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure PageControl_SetComponentIndex(AObj: TPageControl; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function PageControl_GetOwner(AObj: TPageControl): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function PageControl_GetName(AObj: TPageControl): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure PageControl_SetName(AObj: TPageControl; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function PageControl_GetTag(AObj: TPageControl): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure PageControl_SetTag(AObj: TPageControl; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function PageControl_GetPages(AObj: TPageControl; Index: Integer): TTabSheet; extdecl;
begin
  Result :=  AObj.Pages[Index];
end;

function PageControl_GetControls(AObj: TPageControl; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function PageControl_GetComponents(AObj: TPageControl; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  PageControl_Create,
  PageControl_Free,
  PageControl_SelectNextPage,
  PageControl_CanFocus,
  PageControl_FlipChildren,
  PageControl_Focused,
  PageControl_Invalidate,
  PageControl_Realign,
  PageControl_Repaint,
  PageControl_ScaleBy,
  PageControl_SetBounds,
  PageControl_SetFocus,
  PageControl_Update,
  PageControl_BringToFront,
  PageControl_HasParent,
  PageControl_Hide,
  PageControl_Perform,
  PageControl_Refresh,
  PageControl_SendToBack,
  PageControl_Show,
  PageControl_GetTextBuf,
  PageControl_GetTextLen,
  PageControl_FindComponent,
  PageControl_GetNamePath,
  PageControl_Assign,
  PageControl_ClassName,
  PageControl_Equals,
  PageControl_GetHashCode,
  PageControl_ToString,
  PageControl_GetActivePageIndex,
  PageControl_SetActivePageIndex,
  PageControl_GetPageCount,
  PageControl_GetAlign,
  PageControl_SetAlign,
  PageControl_GetAnchors,
  PageControl_SetAnchors,
  PageControl_GetBiDiMode,
  PageControl_SetBiDiMode,
  PageControl_GetDoubleBuffered,
  PageControl_SetDoubleBuffered,
  PageControl_GetEnabled,
  PageControl_SetEnabled,
  PageControl_GetFont,
  PageControl_SetFont,
  PageControl_GetImages,
  PageControl_SetImages,
  PageControl_GetMultiLine,
  PageControl_SetMultiLine,
  PageControl_GetParentFont,
  PageControl_SetParentFont,
  PageControl_GetParentShowHint,
  PageControl_SetParentShowHint,
  PageControl_GetPopupMenu,
  PageControl_SetPopupMenu,
  PageControl_GetShowHint,
  PageControl_SetShowHint,
  PageControl_GetTabIndex,
  PageControl_SetTabIndex,
  PageControl_GetTabOrder,
  PageControl_SetTabOrder,
  PageControl_GetTabPosition,
  PageControl_SetTabPosition,
  PageControl_GetTabStop,
  PageControl_SetTabStop,
  PageControl_GetVisible,
  PageControl_SetVisible,
  PageControl_SetOnChange,
  PageControl_SetOnEnter,
  PageControl_SetOnExit,
  PageControl_SetOnMouseDown,
  PageControl_SetOnMouseEnter,
  PageControl_SetOnMouseLeave,
  PageControl_SetOnMouseMove,
  PageControl_SetOnMouseUp,
  PageControl_SetOnResize,
  PageControl_GetBrush,
  PageControl_GetControlCount,
  PageControl_GetHandle,
  PageControl_GetParentWindow,
  PageControl_SetParentWindow,
  PageControl_GetAction,
  PageControl_SetAction,
  PageControl_GetBoundsRect,
  PageControl_SetBoundsRect,
  PageControl_GetClientHeight,
  PageControl_SetClientHeight,
  PageControl_GetClientRect,
  PageControl_GetClientWidth,
  PageControl_SetClientWidth,
  PageControl_GetParent,
  PageControl_SetParent,
  PageControl_GetLeft,
  PageControl_SetLeft,
  PageControl_GetTop,
  PageControl_SetTop,
  PageControl_GetWidth,
  PageControl_SetWidth,
  PageControl_GetHeight,
  PageControl_SetHeight,
  PageControl_GetCursor,
  PageControl_SetCursor,
  PageControl_GetHint,
  PageControl_SetHint,
  PageControl_GetComponentCount,
  PageControl_GetComponentIndex,
  PageControl_SetComponentIndex,
  PageControl_GetOwner,
  PageControl_GetName,
  PageControl_SetName,
  PageControl_GetTag,
  PageControl_SetTag,
  PageControl_GetPages,
  PageControl_GetControls,
  PageControl_GetComponents;

