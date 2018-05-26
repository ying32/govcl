
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function CategoryPanelGroup_Create(AOwner: TComponent): TCategoryPanelGroup; stdcall;
begin
  Result :=  TCategoryPanelGroup.Create(AOwner);
end;

procedure CategoryPanelGroup_Free(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.Free;
end;

procedure CategoryPanelGroup_CollapseAll(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.CollapseAll;
end;

procedure CategoryPanelGroup_ExpandAll(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.ExpandAll;
end;

function CategoryPanelGroup_CanFocus(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure CategoryPanelGroup_FlipChildren(AObj: TCategoryPanelGroup; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function CategoryPanelGroup_Focused(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function CategoryPanelGroup_HandleAllocated(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure CategoryPanelGroup_Invalidate(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.Invalidate;
end;

procedure CategoryPanelGroup_Realign(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.Realign;
end;

procedure CategoryPanelGroup_Repaint(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.Repaint;
end;

procedure CategoryPanelGroup_ScaleBy(AObj: TCategoryPanelGroup; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure CategoryPanelGroup_SetBounds(AObj: TCategoryPanelGroup; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure CategoryPanelGroup_SetFocus(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.SetFocus;
end;

procedure CategoryPanelGroup_Update(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.Update;
end;

procedure CategoryPanelGroup_BringToFront(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.BringToFront;
end;

function CategoryPanelGroup_HasParent(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure CategoryPanelGroup_Hide(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.Hide;
end;

function CategoryPanelGroup_Perform(AObj: TCategoryPanelGroup; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure CategoryPanelGroup_Refresh(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.Refresh;
end;

procedure CategoryPanelGroup_SendToBack(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.SendToBack;
end;

procedure CategoryPanelGroup_Show(AObj: TCategoryPanelGroup); stdcall;
begin
  AObj.Show;
end;

function CategoryPanelGroup_GetTextBuf(AObj: TCategoryPanelGroup; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function CategoryPanelGroup_GetTextLen(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function CategoryPanelGroup_FindComponent(AObj: TCategoryPanelGroup; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function CategoryPanelGroup_GetNamePath(AObj: TCategoryPanelGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure CategoryPanelGroup_Assign(AObj: TCategoryPanelGroup; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function CategoryPanelGroup_ClassName(AObj: TCategoryPanelGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function CategoryPanelGroup_Equals(AObj: TCategoryPanelGroup; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function CategoryPanelGroup_GetHashCode(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function CategoryPanelGroup_ToString(AObj: TCategoryPanelGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function CategoryPanelGroup_GetAlign(AObj: TCategoryPanelGroup): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure CategoryPanelGroup_SetAlign(AObj: TCategoryPanelGroup; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function CategoryPanelGroup_GetAnchors(AObj: TCategoryPanelGroup): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure CategoryPanelGroup_SetAnchors(AObj: TCategoryPanelGroup; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function CategoryPanelGroup_GetBevelEdges(AObj: TCategoryPanelGroup): TBevelEdges; stdcall;
begin
  Result :=  AObj.BevelEdges;
end;

procedure CategoryPanelGroup_SetBevelEdges(AObj: TCategoryPanelGroup; AValue: TBevelEdges); stdcall;
begin
  AObj.BevelEdges := AValue;
end;

function CategoryPanelGroup_GetBevelInner(AObj: TCategoryPanelGroup): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelInner;
end;

procedure CategoryPanelGroup_SetBevelInner(AObj: TCategoryPanelGroup; AValue: TBevelCut); stdcall;
begin
  AObj.BevelInner := AValue;
end;

function CategoryPanelGroup_GetBevelOuter(AObj: TCategoryPanelGroup): TBevelCut; stdcall;
begin
  Result :=  AObj.BevelOuter;
end;

procedure CategoryPanelGroup_SetBevelOuter(AObj: TCategoryPanelGroup; AValue: TBevelCut); stdcall;
begin
  AObj.BevelOuter := AValue;
end;

function CategoryPanelGroup_GetBevelKind(AObj: TCategoryPanelGroup): TBevelKind; stdcall;
begin
  Result :=  AObj.BevelKind;
end;

procedure CategoryPanelGroup_SetBevelKind(AObj: TCategoryPanelGroup; AValue: TBevelKind); stdcall;
begin
  AObj.BevelKind := AValue;
end;

function CategoryPanelGroup_GetBiDiMode(AObj: TCategoryPanelGroup): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure CategoryPanelGroup_SetBiDiMode(AObj: TCategoryPanelGroup; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function CategoryPanelGroup_GetChevronAlignment(AObj: TCategoryPanelGroup): TAlignment; stdcall;
begin
  Result :=  AObj.ChevronAlignment;
end;

procedure CategoryPanelGroup_SetChevronAlignment(AObj: TCategoryPanelGroup; AValue: TAlignment); stdcall;
begin
  AObj.ChevronAlignment := AValue;
end;

function CategoryPanelGroup_GetChevronColor(AObj: TCategoryPanelGroup): TColor; stdcall;
begin
  Result :=  AObj.ChevronColor;
end;

procedure CategoryPanelGroup_SetChevronColor(AObj: TCategoryPanelGroup; AValue: TColor); stdcall;
begin
  AObj.ChevronColor := AValue;
end;

function CategoryPanelGroup_GetChevronHotColor(AObj: TCategoryPanelGroup): TColor; stdcall;
begin
  Result :=  AObj.ChevronHotColor;
end;

procedure CategoryPanelGroup_SetChevronHotColor(AObj: TCategoryPanelGroup; AValue: TColor); stdcall;
begin
  AObj.ChevronHotColor := AValue;
end;

function CategoryPanelGroup_GetDoubleBuffered(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure CategoryPanelGroup_SetDoubleBuffered(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function CategoryPanelGroup_GetEnabled(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure CategoryPanelGroup_SetEnabled(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function CategoryPanelGroup_GetColor(AObj: TCategoryPanelGroup): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure CategoryPanelGroup_SetColor(AObj: TCategoryPanelGroup; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function CategoryPanelGroup_GetFont(AObj: TCategoryPanelGroup): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure CategoryPanelGroup_SetFont(AObj: TCategoryPanelGroup; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function CategoryPanelGroup_GetGradientBaseColor(AObj: TCategoryPanelGroup): TColor; stdcall;
begin
  Result :=  AObj.GradientBaseColor;
end;

procedure CategoryPanelGroup_SetGradientBaseColor(AObj: TCategoryPanelGroup; AValue: TColor); stdcall;
begin
  AObj.GradientBaseColor := AValue;
end;

function CategoryPanelGroup_GetGradientColor(AObj: TCategoryPanelGroup): TColor; stdcall;
begin
  Result :=  AObj.GradientColor;
end;

procedure CategoryPanelGroup_SetGradientColor(AObj: TCategoryPanelGroup; AValue: TColor); stdcall;
begin
  AObj.GradientColor := AValue;
end;

function CategoryPanelGroup_GetGradientDirection(AObj: TCategoryPanelGroup): TGradientDirection; stdcall;
begin
  Result :=  AObj.GradientDirection;
end;

procedure CategoryPanelGroup_SetGradientDirection(AObj: TCategoryPanelGroup; AValue: TGradientDirection); stdcall;
begin
  AObj.GradientDirection := AValue;
end;

function CategoryPanelGroup_GetHeaderAlignment(AObj: TCategoryPanelGroup): TAlignment; stdcall;
begin
  Result :=  AObj.HeaderAlignment;
end;

procedure CategoryPanelGroup_SetHeaderAlignment(AObj: TCategoryPanelGroup; AValue: TAlignment); stdcall;
begin
  AObj.HeaderAlignment := AValue;
end;

function CategoryPanelGroup_GetHeaderFont(AObj: TCategoryPanelGroup): TFont; stdcall;
begin
  Result :=  AObj.HeaderFont;
end;

procedure CategoryPanelGroup_SetHeaderFont(AObj: TCategoryPanelGroup; AValue: TFont); stdcall;
begin
  AObj.HeaderFont := AValue;
end;

function CategoryPanelGroup_GetHeaderHeight(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.HeaderHeight;
end;

procedure CategoryPanelGroup_SetHeaderHeight(AObj: TCategoryPanelGroup; AValue: Integer); stdcall;
begin
  AObj.HeaderHeight := AValue;
end;

function CategoryPanelGroup_GetHeaderImage(AObj: TCategoryPanelGroup): TPicture; stdcall;
begin
  Result :=  AObj.HeaderImage;
end;

procedure CategoryPanelGroup_SetHeaderImage(AObj: TCategoryPanelGroup; AValue: TPicture); stdcall;
begin
  AObj.HeaderImage := AValue;
end;

function CategoryPanelGroup_GetHeaderStyle(AObj: TCategoryPanelGroup): TCustomCategoryPanelGroup.THeaderStyle; stdcall;
begin
  Result :=  AObj.HeaderStyle;
end;

procedure CategoryPanelGroup_SetHeaderStyle(AObj: TCategoryPanelGroup; AValue: TCustomCategoryPanelGroup.THeaderStyle); stdcall;
begin
  AObj.HeaderStyle := AValue;
end;

function CategoryPanelGroup_GetHeight(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure CategoryPanelGroup_SetHeight(AObj: TCategoryPanelGroup; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function CategoryPanelGroup_GetImages(AObj: TCategoryPanelGroup): TCustomImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure CategoryPanelGroup_SetImages(AObj: TCategoryPanelGroup; AValue: TCustomImageList); stdcall;
begin
  AObj.Images := AValue;
end;

function CategoryPanelGroup_GetParentBackground(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentBackground;
end;

procedure CategoryPanelGroup_SetParentBackground(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentBackground := AValue;
end;

function CategoryPanelGroup_GetParentColor(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure CategoryPanelGroup_SetParentColor(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function CategoryPanelGroup_GetParentCtl3D(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure CategoryPanelGroup_SetParentCtl3D(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentCtl3D := AValue;
end;

function CategoryPanelGroup_GetParentDoubleBuffered(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure CategoryPanelGroup_SetParentDoubleBuffered(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function CategoryPanelGroup_GetParentFont(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure CategoryPanelGroup_SetParentFont(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function CategoryPanelGroup_GetParentShowHint(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure CategoryPanelGroup_SetParentShowHint(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function CategoryPanelGroup_GetPopupMenu(AObj: TCategoryPanelGroup): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure CategoryPanelGroup_SetPopupMenu(AObj: TCategoryPanelGroup; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function CategoryPanelGroup_GetShowHint(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure CategoryPanelGroup_SetShowHint(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function CategoryPanelGroup_GetTabOrder(AObj: TCategoryPanelGroup): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure CategoryPanelGroup_SetTabOrder(AObj: TCategoryPanelGroup; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function CategoryPanelGroup_GetTabStop(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure CategoryPanelGroup_SetTabStop(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function CategoryPanelGroup_GetVisible(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure CategoryPanelGroup_SetVisible(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function CategoryPanelGroup_GetStyleElements(AObj: TCategoryPanelGroup): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure CategoryPanelGroup_SetStyleElements(AObj: TCategoryPanelGroup; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function CategoryPanelGroup_GetWidth(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure CategoryPanelGroup_SetWidth(AObj: TCategoryPanelGroup; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

procedure CategoryPanelGroup_SetOnClick(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanelGroup_SetOnDblClick(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanelGroup_SetOnEnter(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanelGroup_SetOnExit(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanelGroup_SetOnMouseDown(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanelGroup_SetOnMouseEnter(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanelGroup_SetOnMouseLeave(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanelGroup_SetOnMouseMove(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanelGroup_SetOnMouseUp(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanelGroup_SetOnMouseWheel(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

procedure CategoryPanelGroup_SetOnResize(AObj: TCategoryPanelGroup; AEventId: NativeUInt); stdcall;
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

function CategoryPanelGroup_GetPanels(AObj: TCategoryPanelGroup): TList; stdcall;
begin
  Result :=  AObj.Panels;
end;

function CategoryPanelGroup_GetBrush(AObj: TCategoryPanelGroup): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function CategoryPanelGroup_GetControlCount(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function CategoryPanelGroup_GetHandle(AObj: TCategoryPanelGroup): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function CategoryPanelGroup_GetParentWindow(AObj: TCategoryPanelGroup): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure CategoryPanelGroup_SetParentWindow(AObj: TCategoryPanelGroup; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function CategoryPanelGroup_GetAction(AObj: TCategoryPanelGroup): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure CategoryPanelGroup_SetAction(AObj: TCategoryPanelGroup; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure CategoryPanelGroup_GetBoundsRect(AObj: TCategoryPanelGroup; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure CategoryPanelGroup_SetBoundsRect(AObj: TCategoryPanelGroup; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function CategoryPanelGroup_GetClientHeight(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure CategoryPanelGroup_SetClientHeight(AObj: TCategoryPanelGroup; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure CategoryPanelGroup_GetClientRect(AObj: TCategoryPanelGroup; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function CategoryPanelGroup_GetClientWidth(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure CategoryPanelGroup_SetClientWidth(AObj: TCategoryPanelGroup; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function CategoryPanelGroup_GetExplicitLeft(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function CategoryPanelGroup_GetExplicitTop(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function CategoryPanelGroup_GetExplicitWidth(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function CategoryPanelGroup_GetExplicitHeight(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function CategoryPanelGroup_GetParent(AObj: TCategoryPanelGroup): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure CategoryPanelGroup_SetParent(AObj: TCategoryPanelGroup; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function CategoryPanelGroup_GetAlignWithMargins(AObj: TCategoryPanelGroup): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure CategoryPanelGroup_SetAlignWithMargins(AObj: TCategoryPanelGroup; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function CategoryPanelGroup_GetLeft(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure CategoryPanelGroup_SetLeft(AObj: TCategoryPanelGroup; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function CategoryPanelGroup_GetTop(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure CategoryPanelGroup_SetTop(AObj: TCategoryPanelGroup; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function CategoryPanelGroup_GetCursor(AObj: TCategoryPanelGroup): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure CategoryPanelGroup_SetCursor(AObj: TCategoryPanelGroup; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function CategoryPanelGroup_GetHint(AObj: TCategoryPanelGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure CategoryPanelGroup_SetHint(AObj: TCategoryPanelGroup; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function CategoryPanelGroup_GetMargins(AObj: TCategoryPanelGroup): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure CategoryPanelGroup_SetMargins(AObj: TCategoryPanelGroup; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function CategoryPanelGroup_GetCustomHint(AObj: TCategoryPanelGroup): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure CategoryPanelGroup_SetCustomHint(AObj: TCategoryPanelGroup; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function CategoryPanelGroup_GetComponentCount(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function CategoryPanelGroup_GetComponentIndex(AObj: TCategoryPanelGroup): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure CategoryPanelGroup_SetComponentIndex(AObj: TCategoryPanelGroup; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function CategoryPanelGroup_GetOwner(AObj: TCategoryPanelGroup): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function CategoryPanelGroup_GetName(AObj: TCategoryPanelGroup): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure CategoryPanelGroup_SetName(AObj: TCategoryPanelGroup; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function CategoryPanelGroup_GetTag(AObj: TCategoryPanelGroup): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure CategoryPanelGroup_SetTag(AObj: TCategoryPanelGroup; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function CategoryPanelGroup_GetControls(AObj: TCategoryPanelGroup; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function CategoryPanelGroup_GetComponents(AObj: TCategoryPanelGroup; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  CategoryPanelGroup_Create,
  CategoryPanelGroup_Free,
  CategoryPanelGroup_CollapseAll,
  CategoryPanelGroup_ExpandAll,
  CategoryPanelGroup_CanFocus,
  CategoryPanelGroup_FlipChildren,
  CategoryPanelGroup_Focused,
  CategoryPanelGroup_HandleAllocated,
  CategoryPanelGroup_Invalidate,
  CategoryPanelGroup_Realign,
  CategoryPanelGroup_Repaint,
  CategoryPanelGroup_ScaleBy,
  CategoryPanelGroup_SetBounds,
  CategoryPanelGroup_SetFocus,
  CategoryPanelGroup_Update,
  CategoryPanelGroup_BringToFront,
  CategoryPanelGroup_HasParent,
  CategoryPanelGroup_Hide,
  CategoryPanelGroup_Perform,
  CategoryPanelGroup_Refresh,
  CategoryPanelGroup_SendToBack,
  CategoryPanelGroup_Show,
  CategoryPanelGroup_GetTextBuf,
  CategoryPanelGroup_GetTextLen,
  CategoryPanelGroup_FindComponent,
  CategoryPanelGroup_GetNamePath,
  CategoryPanelGroup_Assign,
  CategoryPanelGroup_ClassName,
  CategoryPanelGroup_Equals,
  CategoryPanelGroup_GetHashCode,
  CategoryPanelGroup_ToString,
  CategoryPanelGroup_GetAlign,
  CategoryPanelGroup_SetAlign,
  CategoryPanelGroup_GetAnchors,
  CategoryPanelGroup_SetAnchors,
  CategoryPanelGroup_GetBevelEdges,
  CategoryPanelGroup_SetBevelEdges,
  CategoryPanelGroup_GetBevelInner,
  CategoryPanelGroup_SetBevelInner,
  CategoryPanelGroup_GetBevelOuter,
  CategoryPanelGroup_SetBevelOuter,
  CategoryPanelGroup_GetBevelKind,
  CategoryPanelGroup_SetBevelKind,
  CategoryPanelGroup_GetBiDiMode,
  CategoryPanelGroup_SetBiDiMode,
  CategoryPanelGroup_GetChevronAlignment,
  CategoryPanelGroup_SetChevronAlignment,
  CategoryPanelGroup_GetChevronColor,
  CategoryPanelGroup_SetChevronColor,
  CategoryPanelGroup_GetChevronHotColor,
  CategoryPanelGroup_SetChevronHotColor,
  CategoryPanelGroup_GetDoubleBuffered,
  CategoryPanelGroup_SetDoubleBuffered,
  CategoryPanelGroup_GetEnabled,
  CategoryPanelGroup_SetEnabled,
  CategoryPanelGroup_GetColor,
  CategoryPanelGroup_SetColor,
  CategoryPanelGroup_GetFont,
  CategoryPanelGroup_SetFont,
  CategoryPanelGroup_GetGradientBaseColor,
  CategoryPanelGroup_SetGradientBaseColor,
  CategoryPanelGroup_GetGradientColor,
  CategoryPanelGroup_SetGradientColor,
  CategoryPanelGroup_GetGradientDirection,
  CategoryPanelGroup_SetGradientDirection,
  CategoryPanelGroup_GetHeaderAlignment,
  CategoryPanelGroup_SetHeaderAlignment,
  CategoryPanelGroup_GetHeaderFont,
  CategoryPanelGroup_SetHeaderFont,
  CategoryPanelGroup_GetHeaderHeight,
  CategoryPanelGroup_SetHeaderHeight,
  CategoryPanelGroup_GetHeaderImage,
  CategoryPanelGroup_SetHeaderImage,
  CategoryPanelGroup_GetHeaderStyle,
  CategoryPanelGroup_SetHeaderStyle,
  CategoryPanelGroup_GetHeight,
  CategoryPanelGroup_SetHeight,
  CategoryPanelGroup_GetImages,
  CategoryPanelGroup_SetImages,
  CategoryPanelGroup_GetParentBackground,
  CategoryPanelGroup_SetParentBackground,
  CategoryPanelGroup_GetParentColor,
  CategoryPanelGroup_SetParentColor,
  CategoryPanelGroup_GetParentCtl3D,
  CategoryPanelGroup_SetParentCtl3D,
  CategoryPanelGroup_GetParentDoubleBuffered,
  CategoryPanelGroup_SetParentDoubleBuffered,
  CategoryPanelGroup_GetParentFont,
  CategoryPanelGroup_SetParentFont,
  CategoryPanelGroup_GetParentShowHint,
  CategoryPanelGroup_SetParentShowHint,
  CategoryPanelGroup_GetPopupMenu,
  CategoryPanelGroup_SetPopupMenu,
  CategoryPanelGroup_GetShowHint,
  CategoryPanelGroup_SetShowHint,
  CategoryPanelGroup_GetTabOrder,
  CategoryPanelGroup_SetTabOrder,
  CategoryPanelGroup_GetTabStop,
  CategoryPanelGroup_SetTabStop,
  CategoryPanelGroup_GetVisible,
  CategoryPanelGroup_SetVisible,
  CategoryPanelGroup_GetStyleElements,
  CategoryPanelGroup_SetStyleElements,
  CategoryPanelGroup_GetWidth,
  CategoryPanelGroup_SetWidth,
  CategoryPanelGroup_SetOnClick,
  CategoryPanelGroup_SetOnDblClick,
  CategoryPanelGroup_SetOnEnter,
  CategoryPanelGroup_SetOnExit,
  CategoryPanelGroup_SetOnMouseDown,
  CategoryPanelGroup_SetOnMouseEnter,
  CategoryPanelGroup_SetOnMouseLeave,
  CategoryPanelGroup_SetOnMouseMove,
  CategoryPanelGroup_SetOnMouseUp,
  CategoryPanelGroup_SetOnMouseWheel,
  CategoryPanelGroup_SetOnResize,
  CategoryPanelGroup_GetPanels,
  CategoryPanelGroup_GetBrush,
  CategoryPanelGroup_GetControlCount,
  CategoryPanelGroup_GetHandle,
  CategoryPanelGroup_GetParentWindow,
  CategoryPanelGroup_SetParentWindow,
  CategoryPanelGroup_GetAction,
  CategoryPanelGroup_SetAction,
  CategoryPanelGroup_GetBoundsRect,
  CategoryPanelGroup_SetBoundsRect,
  CategoryPanelGroup_GetClientHeight,
  CategoryPanelGroup_SetClientHeight,
  CategoryPanelGroup_GetClientRect,
  CategoryPanelGroup_GetClientWidth,
  CategoryPanelGroup_SetClientWidth,
  CategoryPanelGroup_GetExplicitLeft,
  CategoryPanelGroup_GetExplicitTop,
  CategoryPanelGroup_GetExplicitWidth,
  CategoryPanelGroup_GetExplicitHeight,
  CategoryPanelGroup_GetParent,
  CategoryPanelGroup_SetParent,
  CategoryPanelGroup_GetAlignWithMargins,
  CategoryPanelGroup_SetAlignWithMargins,
  CategoryPanelGroup_GetLeft,
  CategoryPanelGroup_SetLeft,
  CategoryPanelGroup_GetTop,
  CategoryPanelGroup_SetTop,
  CategoryPanelGroup_GetCursor,
  CategoryPanelGroup_SetCursor,
  CategoryPanelGroup_GetHint,
  CategoryPanelGroup_SetHint,
  CategoryPanelGroup_GetMargins,
  CategoryPanelGroup_SetMargins,
  CategoryPanelGroup_GetCustomHint,
  CategoryPanelGroup_SetCustomHint,
  CategoryPanelGroup_GetComponentCount,
  CategoryPanelGroup_GetComponentIndex,
  CategoryPanelGroup_SetComponentIndex,
  CategoryPanelGroup_GetOwner,
  CategoryPanelGroup_GetName,
  CategoryPanelGroup_SetName,
  CategoryPanelGroup_GetTag,
  CategoryPanelGroup_SetTag,
  CategoryPanelGroup_GetControls,
  CategoryPanelGroup_GetComponents;

