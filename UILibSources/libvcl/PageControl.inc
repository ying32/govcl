
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function PageControl_Create(AOwner: TComponent): TPageControl; stdcall;
begin
  Result :=  TPageControl.Create(AOwner);
end;

procedure PageControl_Free(AObj: TPageControl); stdcall;
begin
  AObj.Free;
end;

procedure PageControl_SelectNextPage(AObj: TPageControl; GoForward: LongBool; CheckTabVisible: LongBool); stdcall;
begin
  AObj.SelectNextPage(GoForward, CheckTabVisible);
end;

function PageControl_RowCount(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.RowCount;
end;

function PageControl_CanFocus(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure PageControl_FlipChildren(AObj: TPageControl; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function PageControl_Focused(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function PageControl_HandleAllocated(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure PageControl_Invalidate(AObj: TPageControl); stdcall;
begin
  AObj.Invalidate;
end;

procedure PageControl_Realign(AObj: TPageControl); stdcall;
begin
  AObj.Realign;
end;

procedure PageControl_Repaint(AObj: TPageControl); stdcall;
begin
  AObj.Repaint;
end;

procedure PageControl_ScaleBy(AObj: TPageControl; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure PageControl_SetBounds(AObj: TPageControl; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure PageControl_SetFocus(AObj: TPageControl); stdcall;
begin
  AObj.SetFocus;
end;

procedure PageControl_Update(AObj: TPageControl); stdcall;
begin
  AObj.Update;
end;

procedure PageControl_BringToFront(AObj: TPageControl); stdcall;
begin
  AObj.BringToFront;
end;

function PageControl_HasParent(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure PageControl_Hide(AObj: TPageControl); stdcall;
begin
  AObj.Hide;
end;

function PageControl_Perform(AObj: TPageControl; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure PageControl_Refresh(AObj: TPageControl); stdcall;
begin
  AObj.Refresh;
end;

procedure PageControl_SendToBack(AObj: TPageControl); stdcall;
begin
  AObj.SendToBack;
end;

procedure PageControl_Show(AObj: TPageControl); stdcall;
begin
  AObj.Show;
end;

function PageControl_GetTextBuf(AObj: TPageControl; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function PageControl_GetTextLen(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function PageControl_FindComponent(AObj: TPageControl; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function PageControl_GetNamePath(AObj: TPageControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure PageControl_Assign(AObj: TPageControl; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function PageControl_ClassName(AObj: TPageControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function PageControl_Equals(AObj: TPageControl; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function PageControl_GetHashCode(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function PageControl_ToString(AObj: TPageControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function PageControl_GetActivePageIndex(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.ActivePageIndex;
end;

procedure PageControl_SetActivePageIndex(AObj: TPageControl; AValue: Integer); stdcall;
begin
  AObj.ActivePageIndex := AValue;
end;

function PageControl_GetPageCount(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.PageCount;
end;

function PageControl_GetAlign(AObj: TPageControl): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure PageControl_SetAlign(AObj: TPageControl; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function PageControl_GetAnchors(AObj: TPageControl): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure PageControl_SetAnchors(AObj: TPageControl; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function PageControl_GetBiDiMode(AObj: TPageControl): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure PageControl_SetBiDiMode(AObj: TPageControl; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function PageControl_GetDoubleBuffered(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure PageControl_SetDoubleBuffered(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function PageControl_GetEnabled(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure PageControl_SetEnabled(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function PageControl_GetFont(AObj: TPageControl): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure PageControl_SetFont(AObj: TPageControl; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function PageControl_GetHotTrack(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.HotTrack;
end;

procedure PageControl_SetHotTrack(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.HotTrack := AValue;
end;

function PageControl_GetImages(AObj: TPageControl): TCustomImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure PageControl_SetImages(AObj: TPageControl; AValue: TCustomImageList); stdcall;
begin
  AObj.Images := AValue;
end;

function PageControl_GetMultiLine(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.MultiLine;
end;

procedure PageControl_SetMultiLine(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.MultiLine := AValue;
end;

function PageControl_GetParentDoubleBuffered(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure PageControl_SetParentDoubleBuffered(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function PageControl_GetParentFont(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure PageControl_SetParentFont(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function PageControl_GetParentShowHint(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure PageControl_SetParentShowHint(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function PageControl_GetPopupMenu(AObj: TPageControl): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure PageControl_SetPopupMenu(AObj: TPageControl; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function PageControl_GetShowHint(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure PageControl_SetShowHint(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function PageControl_GetStyle(AObj: TPageControl): TTabStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure PageControl_SetStyle(AObj: TPageControl; AValue: TTabStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function PageControl_GetTabHeight(AObj: TPageControl): SmallInt; stdcall;
begin
  Result :=  AObj.TabHeight;
end;

procedure PageControl_SetTabHeight(AObj: TPageControl; AValue: SmallInt); stdcall;
begin
  AObj.TabHeight := AValue;
end;

function PageControl_GetTabIndex(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.TabIndex;
end;

procedure PageControl_SetTabIndex(AObj: TPageControl; AValue: Integer); stdcall;
begin
  AObj.TabIndex := AValue;
end;

function PageControl_GetTabOrder(AObj: TPageControl): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure PageControl_SetTabOrder(AObj: TPageControl; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function PageControl_GetTabPosition(AObj: TPageControl): TTabPosition; stdcall;
begin
  Result :=  AObj.TabPosition;
end;

procedure PageControl_SetTabPosition(AObj: TPageControl; AValue: TTabPosition); stdcall;
begin
  AObj.TabPosition := AValue;
end;

function PageControl_GetTabStop(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure PageControl_SetTabStop(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function PageControl_GetTabWidth(AObj: TPageControl): SmallInt; stdcall;
begin
  Result :=  AObj.TabWidth;
end;

procedure PageControl_SetTabWidth(AObj: TPageControl; AValue: SmallInt); stdcall;
begin
  AObj.TabWidth := AValue;
end;

function PageControl_GetVisible(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure PageControl_SetVisible(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function PageControl_GetStyleElements(AObj: TPageControl): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure PageControl_SetStyleElements(AObj: TPageControl; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure PageControl_SetOnChange(AObj: TPageControl; AEventId: NativeUInt); stdcall;
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

procedure PageControl_SetOnEnter(AObj: TPageControl; AEventId: NativeUInt); stdcall;
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

procedure PageControl_SetOnExit(AObj: TPageControl; AEventId: NativeUInt); stdcall;
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

procedure PageControl_SetOnGetImageIndex(AObj: TPageControl; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetImageIndex := nil;
    TEventClass.Remove(AObj, gePageControlGetImageIndex);
    Exit;
  end;
  AObj.OnGetImageIndex := TEventClass.PageControlOnGetImageIndex;
  TEventClass.Add(AObj, gePageControlGetImageIndex, AEventId);
end;

procedure PageControl_SetOnMouseDown(AObj: TPageControl; AEventId: NativeUInt); stdcall;
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

procedure PageControl_SetOnMouseEnter(AObj: TPageControl; AEventId: NativeUInt); stdcall;
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

procedure PageControl_SetOnMouseLeave(AObj: TPageControl; AEventId: NativeUInt); stdcall;
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

procedure PageControl_SetOnMouseMove(AObj: TPageControl; AEventId: NativeUInt); stdcall;
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

procedure PageControl_SetOnMouseUp(AObj: TPageControl; AEventId: NativeUInt); stdcall;
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

procedure PageControl_SetOnResize(AObj: TPageControl; AEventId: NativeUInt); stdcall;
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

function PageControl_GetCanvas(AObj: TPageControl): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function PageControl_GetBrush(AObj: TPageControl): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function PageControl_GetControlCount(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function PageControl_GetHandle(AObj: TPageControl): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function PageControl_GetParentWindow(AObj: TPageControl): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure PageControl_SetParentWindow(AObj: TPageControl; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function PageControl_GetAction(AObj: TPageControl): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure PageControl_SetAction(AObj: TPageControl; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure PageControl_GetBoundsRect(AObj: TPageControl; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure PageControl_SetBoundsRect(AObj: TPageControl; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function PageControl_GetClientHeight(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure PageControl_SetClientHeight(AObj: TPageControl; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure PageControl_GetClientRect(AObj: TPageControl; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function PageControl_GetClientWidth(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure PageControl_SetClientWidth(AObj: TPageControl; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function PageControl_GetExplicitLeft(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function PageControl_GetExplicitTop(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function PageControl_GetExplicitWidth(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function PageControl_GetExplicitHeight(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function PageControl_GetParent(AObj: TPageControl): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure PageControl_SetParent(AObj: TPageControl; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function PageControl_GetAlignWithMargins(AObj: TPageControl): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure PageControl_SetAlignWithMargins(AObj: TPageControl; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function PageControl_GetLeft(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure PageControl_SetLeft(AObj: TPageControl; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function PageControl_GetTop(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure PageControl_SetTop(AObj: TPageControl; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function PageControl_GetWidth(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure PageControl_SetWidth(AObj: TPageControl; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function PageControl_GetHeight(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure PageControl_SetHeight(AObj: TPageControl; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function PageControl_GetCursor(AObj: TPageControl): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure PageControl_SetCursor(AObj: TPageControl; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function PageControl_GetHint(AObj: TPageControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure PageControl_SetHint(AObj: TPageControl; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function PageControl_GetMargins(AObj: TPageControl): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure PageControl_SetMargins(AObj: TPageControl; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function PageControl_GetCustomHint(AObj: TPageControl): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure PageControl_SetCustomHint(AObj: TPageControl; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function PageControl_GetComponentCount(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function PageControl_GetComponentIndex(AObj: TPageControl): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure PageControl_SetComponentIndex(AObj: TPageControl; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function PageControl_GetOwner(AObj: TPageControl): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function PageControl_GetName(AObj: TPageControl): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure PageControl_SetName(AObj: TPageControl; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function PageControl_GetTag(AObj: TPageControl): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure PageControl_SetTag(AObj: TPageControl; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function PageControl_GetPages(AObj: TPageControl; Index: Integer): TTabSheet; stdcall;
begin
  Result :=  AObj.Pages[Index];
end;

function PageControl_GetControls(AObj: TPageControl; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function PageControl_GetComponents(AObj: TPageControl; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  PageControl_Create,
  PageControl_Free,
  PageControl_SelectNextPage,
  PageControl_RowCount,
  PageControl_CanFocus,
  PageControl_FlipChildren,
  PageControl_Focused,
  PageControl_HandleAllocated,
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
  PageControl_GetHotTrack,
  PageControl_SetHotTrack,
  PageControl_GetImages,
  PageControl_SetImages,
  PageControl_GetMultiLine,
  PageControl_SetMultiLine,
  PageControl_GetParentDoubleBuffered,
  PageControl_SetParentDoubleBuffered,
  PageControl_GetParentFont,
  PageControl_SetParentFont,
  PageControl_GetParentShowHint,
  PageControl_SetParentShowHint,
  PageControl_GetPopupMenu,
  PageControl_SetPopupMenu,
  PageControl_GetShowHint,
  PageControl_SetShowHint,
  PageControl_GetStyle,
  PageControl_SetStyle,
  PageControl_GetTabHeight,
  PageControl_SetTabHeight,
  PageControl_GetTabIndex,
  PageControl_SetTabIndex,
  PageControl_GetTabOrder,
  PageControl_SetTabOrder,
  PageControl_GetTabPosition,
  PageControl_SetTabPosition,
  PageControl_GetTabStop,
  PageControl_SetTabStop,
  PageControl_GetTabWidth,
  PageControl_SetTabWidth,
  PageControl_GetVisible,
  PageControl_SetVisible,
  PageControl_GetStyleElements,
  PageControl_SetStyleElements,
  PageControl_SetOnChange,
  PageControl_SetOnEnter,
  PageControl_SetOnExit,
  PageControl_SetOnGetImageIndex,
  PageControl_SetOnMouseDown,
  PageControl_SetOnMouseEnter,
  PageControl_SetOnMouseLeave,
  PageControl_SetOnMouseMove,
  PageControl_SetOnMouseUp,
  PageControl_SetOnResize,
  PageControl_GetCanvas,
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
  PageControl_GetExplicitLeft,
  PageControl_GetExplicitTop,
  PageControl_GetExplicitWidth,
  PageControl_GetExplicitHeight,
  PageControl_GetParent,
  PageControl_SetParent,
  PageControl_GetAlignWithMargins,
  PageControl_SetAlignWithMargins,
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
  PageControl_GetMargins,
  PageControl_SetMargins,
  PageControl_GetCustomHint,
  PageControl_SetCustomHint,
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

