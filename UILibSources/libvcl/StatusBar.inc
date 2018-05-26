
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function StatusBar_Create(AOwner: TComponent): TStatusBar; stdcall;
begin
  Result :=  TStatusBar.Create(AOwner);
end;

procedure StatusBar_Free(AObj: TStatusBar); stdcall;
begin
  AObj.Free;
end;

procedure StatusBar_FlipChildren(AObj: TStatusBar; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

procedure StatusBar_SetBounds(AObj: TStatusBar; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

function StatusBar_CanFocus(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

function StatusBar_Focused(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function StatusBar_HandleAllocated(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure StatusBar_Invalidate(AObj: TStatusBar); stdcall;
begin
  AObj.Invalidate;
end;

procedure StatusBar_Realign(AObj: TStatusBar); stdcall;
begin
  AObj.Realign;
end;

procedure StatusBar_Repaint(AObj: TStatusBar); stdcall;
begin
  AObj.Repaint;
end;

procedure StatusBar_ScaleBy(AObj: TStatusBar; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure StatusBar_SetFocus(AObj: TStatusBar); stdcall;
begin
  AObj.SetFocus;
end;

procedure StatusBar_Update(AObj: TStatusBar); stdcall;
begin
  AObj.Update;
end;

procedure StatusBar_BringToFront(AObj: TStatusBar); stdcall;
begin
  AObj.BringToFront;
end;

function StatusBar_HasParent(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure StatusBar_Hide(AObj: TStatusBar); stdcall;
begin
  AObj.Hide;
end;

function StatusBar_Perform(AObj: TStatusBar; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure StatusBar_Refresh(AObj: TStatusBar); stdcall;
begin
  AObj.Refresh;
end;

procedure StatusBar_SendToBack(AObj: TStatusBar); stdcall;
begin
  AObj.SendToBack;
end;

procedure StatusBar_Show(AObj: TStatusBar); stdcall;
begin
  AObj.Show;
end;

function StatusBar_GetTextBuf(AObj: TStatusBar; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function StatusBar_GetTextLen(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function StatusBar_FindComponent(AObj: TStatusBar; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function StatusBar_GetNamePath(AObj: TStatusBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure StatusBar_Assign(AObj: TStatusBar; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function StatusBar_ClassName(AObj: TStatusBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function StatusBar_Equals(AObj: TStatusBar; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function StatusBar_GetHashCode(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function StatusBar_ToString(AObj: TStatusBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function StatusBar_GetAction(AObj: TStatusBar): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure StatusBar_SetAction(AObj: TStatusBar; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function StatusBar_GetAutoHint(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.AutoHint;
end;

procedure StatusBar_SetAutoHint(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.AutoHint := AValue;
end;

function StatusBar_GetAlign(AObj: TStatusBar): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure StatusBar_SetAlign(AObj: TStatusBar; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function StatusBar_GetAnchors(AObj: TStatusBar): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure StatusBar_SetAnchors(AObj: TStatusBar; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function StatusBar_GetBiDiMode(AObj: TStatusBar): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure StatusBar_SetBiDiMode(AObj: TStatusBar; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function StatusBar_GetBorderWidth(AObj: TStatusBar): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure StatusBar_SetBorderWidth(AObj: TStatusBar; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function StatusBar_GetColor(AObj: TStatusBar): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure StatusBar_SetColor(AObj: TStatusBar; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function StatusBar_GetDoubleBuffered(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure StatusBar_SetDoubleBuffered(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function StatusBar_GetEnabled(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure StatusBar_SetEnabled(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function StatusBar_GetFont(AObj: TStatusBar): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure StatusBar_SetFont(AObj: TStatusBar; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function StatusBar_GetPanels(AObj: TStatusBar): TStatusPanels; stdcall;
begin
  Result :=  AObj.Panels;
end;

procedure StatusBar_SetPanels(AObj: TStatusBar; AValue: TStatusPanels); stdcall;
begin
  AObj.Panels := AValue;
end;

function StatusBar_GetParentColor(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure StatusBar_SetParentColor(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function StatusBar_GetParentDoubleBuffered(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure StatusBar_SetParentDoubleBuffered(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function StatusBar_GetParentFont(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure StatusBar_SetParentFont(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function StatusBar_GetParentShowHint(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure StatusBar_SetParentShowHint(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function StatusBar_GetPopupMenu(AObj: TStatusBar): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure StatusBar_SetPopupMenu(AObj: TStatusBar; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function StatusBar_GetShowHint(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure StatusBar_SetShowHint(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function StatusBar_GetSimplePanel(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.SimplePanel;
end;

procedure StatusBar_SetSimplePanel(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.SimplePanel := AValue;
end;

function StatusBar_GetSimpleText(AObj: TStatusBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.SimpleText);
end;

procedure StatusBar_SetSimpleText(AObj: TStatusBar; AValue: PChar); stdcall;
begin
  AObj.SimpleText := AValue;
end;

function StatusBar_GetSizeGrip(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.SizeGrip;
end;

procedure StatusBar_SetSizeGrip(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.SizeGrip := AValue;
end;

function StatusBar_GetUseSystemFont(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.UseSystemFont;
end;

procedure StatusBar_SetUseSystemFont(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.UseSystemFont := AValue;
end;

function StatusBar_GetVisible(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure StatusBar_SetVisible(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function StatusBar_GetStyleElements(AObj: TStatusBar): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure StatusBar_SetStyleElements(AObj: TStatusBar; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure StatusBar_SetOnClick(AObj: TStatusBar; AEventId: NativeUInt); stdcall;
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

procedure StatusBar_SetOnDblClick(AObj: TStatusBar; AEventId: NativeUInt); stdcall;
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

procedure StatusBar_SetOnHint(AObj: TStatusBar; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnHint := nil;
    TEventClass.Remove(AObj, geHint);
    Exit;
  end;
  AObj.OnHint := TEventClass.OnHint;
  TEventClass.Add(AObj, geHint, AEventId);
end;

procedure StatusBar_SetOnMouseDown(AObj: TStatusBar; AEventId: NativeUInt); stdcall;
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

procedure StatusBar_SetOnMouseEnter(AObj: TStatusBar; AEventId: NativeUInt); stdcall;
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

procedure StatusBar_SetOnMouseLeave(AObj: TStatusBar; AEventId: NativeUInt); stdcall;
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

procedure StatusBar_SetOnMouseMove(AObj: TStatusBar; AEventId: NativeUInt); stdcall;
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

procedure StatusBar_SetOnMouseUp(AObj: TStatusBar; AEventId: NativeUInt); stdcall;
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

procedure StatusBar_SetOnResize(AObj: TStatusBar; AEventId: NativeUInt); stdcall;
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

function StatusBar_GetCanvas(AObj: TStatusBar): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function StatusBar_GetBrush(AObj: TStatusBar): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function StatusBar_GetControlCount(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function StatusBar_GetHandle(AObj: TStatusBar): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function StatusBar_GetParentWindow(AObj: TStatusBar): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure StatusBar_SetParentWindow(AObj: TStatusBar; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function StatusBar_GetTabOrder(AObj: TStatusBar): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure StatusBar_SetTabOrder(AObj: TStatusBar; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function StatusBar_GetTabStop(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure StatusBar_SetTabStop(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

procedure StatusBar_GetBoundsRect(AObj: TStatusBar; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure StatusBar_SetBoundsRect(AObj: TStatusBar; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function StatusBar_GetClientHeight(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure StatusBar_SetClientHeight(AObj: TStatusBar; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure StatusBar_GetClientRect(AObj: TStatusBar; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function StatusBar_GetClientWidth(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure StatusBar_SetClientWidth(AObj: TStatusBar; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function StatusBar_GetExplicitLeft(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function StatusBar_GetExplicitTop(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function StatusBar_GetExplicitWidth(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function StatusBar_GetExplicitHeight(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function StatusBar_GetParent(AObj: TStatusBar): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure StatusBar_SetParent(AObj: TStatusBar; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function StatusBar_GetAlignWithMargins(AObj: TStatusBar): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure StatusBar_SetAlignWithMargins(AObj: TStatusBar; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function StatusBar_GetLeft(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure StatusBar_SetLeft(AObj: TStatusBar; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function StatusBar_GetTop(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure StatusBar_SetTop(AObj: TStatusBar; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function StatusBar_GetWidth(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure StatusBar_SetWidth(AObj: TStatusBar; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function StatusBar_GetHeight(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure StatusBar_SetHeight(AObj: TStatusBar; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function StatusBar_GetCursor(AObj: TStatusBar): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure StatusBar_SetCursor(AObj: TStatusBar; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function StatusBar_GetHint(AObj: TStatusBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure StatusBar_SetHint(AObj: TStatusBar; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function StatusBar_GetMargins(AObj: TStatusBar): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure StatusBar_SetMargins(AObj: TStatusBar; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function StatusBar_GetCustomHint(AObj: TStatusBar): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure StatusBar_SetCustomHint(AObj: TStatusBar; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function StatusBar_GetComponentCount(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function StatusBar_GetComponentIndex(AObj: TStatusBar): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure StatusBar_SetComponentIndex(AObj: TStatusBar; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function StatusBar_GetOwner(AObj: TStatusBar): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function StatusBar_GetName(AObj: TStatusBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure StatusBar_SetName(AObj: TStatusBar; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function StatusBar_GetTag(AObj: TStatusBar): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure StatusBar_SetTag(AObj: TStatusBar; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function StatusBar_GetControls(AObj: TStatusBar; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function StatusBar_GetComponents(AObj: TStatusBar; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  StatusBar_Create,
  StatusBar_Free,
  StatusBar_FlipChildren,
  StatusBar_SetBounds,
  StatusBar_CanFocus,
  StatusBar_Focused,
  StatusBar_HandleAllocated,
  StatusBar_Invalidate,
  StatusBar_Realign,
  StatusBar_Repaint,
  StatusBar_ScaleBy,
  StatusBar_SetFocus,
  StatusBar_Update,
  StatusBar_BringToFront,
  StatusBar_HasParent,
  StatusBar_Hide,
  StatusBar_Perform,
  StatusBar_Refresh,
  StatusBar_SendToBack,
  StatusBar_Show,
  StatusBar_GetTextBuf,
  StatusBar_GetTextLen,
  StatusBar_FindComponent,
  StatusBar_GetNamePath,
  StatusBar_Assign,
  StatusBar_ClassName,
  StatusBar_Equals,
  StatusBar_GetHashCode,
  StatusBar_ToString,
  StatusBar_GetAction,
  StatusBar_SetAction,
  StatusBar_GetAutoHint,
  StatusBar_SetAutoHint,
  StatusBar_GetAlign,
  StatusBar_SetAlign,
  StatusBar_GetAnchors,
  StatusBar_SetAnchors,
  StatusBar_GetBiDiMode,
  StatusBar_SetBiDiMode,
  StatusBar_GetBorderWidth,
  StatusBar_SetBorderWidth,
  StatusBar_GetColor,
  StatusBar_SetColor,
  StatusBar_GetDoubleBuffered,
  StatusBar_SetDoubleBuffered,
  StatusBar_GetEnabled,
  StatusBar_SetEnabled,
  StatusBar_GetFont,
  StatusBar_SetFont,
  StatusBar_GetPanels,
  StatusBar_SetPanels,
  StatusBar_GetParentColor,
  StatusBar_SetParentColor,
  StatusBar_GetParentDoubleBuffered,
  StatusBar_SetParentDoubleBuffered,
  StatusBar_GetParentFont,
  StatusBar_SetParentFont,
  StatusBar_GetParentShowHint,
  StatusBar_SetParentShowHint,
  StatusBar_GetPopupMenu,
  StatusBar_SetPopupMenu,
  StatusBar_GetShowHint,
  StatusBar_SetShowHint,
  StatusBar_GetSimplePanel,
  StatusBar_SetSimplePanel,
  StatusBar_GetSimpleText,
  StatusBar_SetSimpleText,
  StatusBar_GetSizeGrip,
  StatusBar_SetSizeGrip,
  StatusBar_GetUseSystemFont,
  StatusBar_SetUseSystemFont,
  StatusBar_GetVisible,
  StatusBar_SetVisible,
  StatusBar_GetStyleElements,
  StatusBar_SetStyleElements,
  StatusBar_SetOnClick,
  StatusBar_SetOnDblClick,
  StatusBar_SetOnHint,
  StatusBar_SetOnMouseDown,
  StatusBar_SetOnMouseEnter,
  StatusBar_SetOnMouseLeave,
  StatusBar_SetOnMouseMove,
  StatusBar_SetOnMouseUp,
  StatusBar_SetOnResize,
  StatusBar_GetCanvas,
  StatusBar_GetBrush,
  StatusBar_GetControlCount,
  StatusBar_GetHandle,
  StatusBar_GetParentWindow,
  StatusBar_SetParentWindow,
  StatusBar_GetTabOrder,
  StatusBar_SetTabOrder,
  StatusBar_GetTabStop,
  StatusBar_SetTabStop,
  StatusBar_GetBoundsRect,
  StatusBar_SetBoundsRect,
  StatusBar_GetClientHeight,
  StatusBar_SetClientHeight,
  StatusBar_GetClientRect,
  StatusBar_GetClientWidth,
  StatusBar_SetClientWidth,
  StatusBar_GetExplicitLeft,
  StatusBar_GetExplicitTop,
  StatusBar_GetExplicitWidth,
  StatusBar_GetExplicitHeight,
  StatusBar_GetParent,
  StatusBar_SetParent,
  StatusBar_GetAlignWithMargins,
  StatusBar_SetAlignWithMargins,
  StatusBar_GetLeft,
  StatusBar_SetLeft,
  StatusBar_GetTop,
  StatusBar_SetTop,
  StatusBar_GetWidth,
  StatusBar_SetWidth,
  StatusBar_GetHeight,
  StatusBar_SetHeight,
  StatusBar_GetCursor,
  StatusBar_SetCursor,
  StatusBar_GetHint,
  StatusBar_SetHint,
  StatusBar_GetMargins,
  StatusBar_SetMargins,
  StatusBar_GetCustomHint,
  StatusBar_SetCustomHint,
  StatusBar_GetComponentCount,
  StatusBar_GetComponentIndex,
  StatusBar_SetComponentIndex,
  StatusBar_GetOwner,
  StatusBar_GetName,
  StatusBar_SetName,
  StatusBar_GetTag,
  StatusBar_SetTag,
  StatusBar_GetControls,
  StatusBar_GetComponents;

