
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Button_Create(AOwner: TComponent): TButton; stdcall;
begin
  Result :=  TButton.Create(AOwner);
end;

procedure Button_Free(AObj: TButton); stdcall;
begin
  AObj.Free;
end;

procedure Button_Click(AObj: TButton); stdcall;
begin
  AObj.Click;
end;

function Button_CanFocus(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure Button_FlipChildren(AObj: TButton; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function Button_Focused(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function Button_HandleAllocated(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure Button_Invalidate(AObj: TButton); stdcall;
begin
  AObj.Invalidate;
end;

procedure Button_Realign(AObj: TButton); stdcall;
begin
  AObj.Realign;
end;

procedure Button_Repaint(AObj: TButton); stdcall;
begin
  AObj.Repaint;
end;

procedure Button_ScaleBy(AObj: TButton; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure Button_SetBounds(AObj: TButton; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Button_SetFocus(AObj: TButton); stdcall;
begin
  AObj.SetFocus;
end;

procedure Button_Update(AObj: TButton); stdcall;
begin
  AObj.Update;
end;

procedure Button_BringToFront(AObj: TButton); stdcall;
begin
  AObj.BringToFront;
end;

function Button_HasParent(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Button_Hide(AObj: TButton); stdcall;
begin
  AObj.Hide;
end;

function Button_Perform(AObj: TButton; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Button_Refresh(AObj: TButton); stdcall;
begin
  AObj.Refresh;
end;

procedure Button_SendToBack(AObj: TButton); stdcall;
begin
  AObj.SendToBack;
end;

procedure Button_Show(AObj: TButton); stdcall;
begin
  AObj.Show;
end;

function Button_GetTextBuf(AObj: TButton; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Button_GetTextLen(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function Button_FindComponent(AObj: TButton; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Button_GetNamePath(AObj: TButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Button_Assign(AObj: TButton; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Button_ClassName(AObj: TButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Button_Equals(AObj: TButton; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Button_GetHashCode(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Button_ToString(AObj: TButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Button_GetAction(AObj: TButton): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Button_SetAction(AObj: TButton; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function Button_GetAlign(AObj: TButton): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Button_SetAlign(AObj: TButton; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Button_GetAnchors(AObj: TButton): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Button_SetAnchors(AObj: TButton; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Button_GetBiDiMode(AObj: TButton): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Button_SetBiDiMode(AObj: TButton; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function Button_GetCancel(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.Cancel;
end;

procedure Button_SetCancel(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.Cancel := AValue;
end;

function Button_GetCaption(AObj: TButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure Button_SetCaption(AObj: TButton; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function Button_GetCommandLinkHint(AObj: TButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.CommandLinkHint);
end;

procedure Button_SetCommandLinkHint(AObj: TButton; AValue: PChar); stdcall;
begin
  AObj.CommandLinkHint := AValue;
end;

function Button_GetDefault(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.Default;
end;

procedure Button_SetDefault(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.Default := AValue;
end;

function Button_GetDisabledImageIndex(AObj: TButton): TImageIndex; stdcall;
begin
  Result :=  AObj.DisabledImageIndex;
end;

procedure Button_SetDisabledImageIndex(AObj: TButton; AValue: TImageIndex); stdcall;
begin
  AObj.DisabledImageIndex := AValue;
end;

function Button_GetDoubleBuffered(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure Button_SetDoubleBuffered(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function Button_GetElevationRequired(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.ElevationRequired;
end;

procedure Button_SetElevationRequired(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.ElevationRequired := AValue;
end;

function Button_GetEnabled(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Button_SetEnabled(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Button_GetFont(AObj: TButton): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure Button_SetFont(AObj: TButton; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function Button_GetHotImageIndex(AObj: TButton): TImageIndex; stdcall;
begin
  Result :=  AObj.HotImageIndex;
end;

procedure Button_SetHotImageIndex(AObj: TButton; AValue: TImageIndex); stdcall;
begin
  AObj.HotImageIndex := AValue;
end;

function Button_GetImageAlignment(AObj: TButton): TImageAlignment; stdcall;
begin
  Result :=  AObj.ImageAlignment;
end;

procedure Button_SetImageAlignment(AObj: TButton; AValue: TImageAlignment); stdcall;
begin
  AObj.ImageAlignment := AValue;
end;

function Button_GetImageIndex(AObj: TButton): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure Button_SetImageIndex(AObj: TButton; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function Button_GetImages(AObj: TButton): TCustomImageList; stdcall;
begin
  Result :=  AObj.Images;
end;

procedure Button_SetImages(AObj: TButton; AValue: TCustomImageList); stdcall;
begin
  AObj.Images := AValue;
end;

function Button_GetModalResult(AObj: TButton): TModalResult; stdcall;
begin
  Result :=  AObj.ModalResult;
end;

procedure Button_SetModalResult(AObj: TButton; AValue: TModalResult); stdcall;
begin
  AObj.ModalResult := AValue;
end;

function Button_GetParentDoubleBuffered(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure Button_SetParentDoubleBuffered(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function Button_GetParentFont(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure Button_SetParentFont(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function Button_GetParentShowHint(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Button_SetParentShowHint(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function Button_GetPopupMenu(AObj: TButton): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Button_SetPopupMenu(AObj: TButton; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function Button_GetPressedImageIndex(AObj: TButton): TImageIndex; stdcall;
begin
  Result :=  AObj.PressedImageIndex;
end;

procedure Button_SetPressedImageIndex(AObj: TButton; AValue: TImageIndex); stdcall;
begin
  AObj.PressedImageIndex := AValue;
end;

function Button_GetSelectedImageIndex(AObj: TButton): TImageIndex; stdcall;
begin
  Result :=  AObj.SelectedImageIndex;
end;

procedure Button_SetSelectedImageIndex(AObj: TButton; AValue: TImageIndex); stdcall;
begin
  AObj.SelectedImageIndex := AValue;
end;

function Button_GetShowHint(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Button_SetShowHint(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Button_GetStyle(AObj: TButton): TCustomButton.TButtonStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure Button_SetStyle(AObj: TButton; AValue: TCustomButton.TButtonStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function Button_GetStylusHotImageIndex(AObj: TButton): TImageIndex; stdcall;
begin
  Result :=  AObj.StylusHotImageIndex;
end;

procedure Button_SetStylusHotImageIndex(AObj: TButton; AValue: TImageIndex); stdcall;
begin
  AObj.StylusHotImageIndex := AValue;
end;

function Button_GetTabOrder(AObj: TButton): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure Button_SetTabOrder(AObj: TButton; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function Button_GetTabStop(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure Button_SetTabStop(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function Button_GetVisible(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Button_SetVisible(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Button_GetWordWrap(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.WordWrap;
end;

procedure Button_SetWordWrap(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.WordWrap := AValue;
end;

function Button_GetStyleElements(AObj: TButton): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Button_SetStyleElements(AObj: TButton; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure Button_SetOnClick(AObj: TButton; AEventId: NativeUInt); stdcall;
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

procedure Button_SetOnEnter(AObj: TButton; AEventId: NativeUInt); stdcall;
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

procedure Button_SetOnExit(AObj: TButton; AEventId: NativeUInt); stdcall;
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

procedure Button_SetOnKeyDown(AObj: TButton; AEventId: NativeUInt); stdcall;
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

procedure Button_SetOnKeyPress(AObj: TButton; AEventId: NativeUInt); stdcall;
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

procedure Button_SetOnKeyUp(AObj: TButton; AEventId: NativeUInt); stdcall;
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

procedure Button_SetOnMouseDown(AObj: TButton; AEventId: NativeUInt); stdcall;
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

procedure Button_SetOnMouseEnter(AObj: TButton; AEventId: NativeUInt); stdcall;
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

procedure Button_SetOnMouseLeave(AObj: TButton; AEventId: NativeUInt); stdcall;
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

procedure Button_SetOnMouseMove(AObj: TButton; AEventId: NativeUInt); stdcall;
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

procedure Button_SetOnMouseUp(AObj: TButton; AEventId: NativeUInt); stdcall;
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

function Button_GetBrush(AObj: TButton): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function Button_GetControlCount(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function Button_GetHandle(AObj: TButton): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function Button_GetParentWindow(AObj: TButton): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure Button_SetParentWindow(AObj: TButton; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

procedure Button_GetBoundsRect(AObj: TButton; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Button_SetBoundsRect(AObj: TButton; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Button_GetClientHeight(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Button_SetClientHeight(AObj: TButton; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Button_GetClientRect(AObj: TButton; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Button_GetClientWidth(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Button_SetClientWidth(AObj: TButton; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Button_GetExplicitLeft(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Button_GetExplicitTop(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Button_GetExplicitWidth(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Button_GetExplicitHeight(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Button_GetParent(AObj: TButton): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Button_SetParent(AObj: TButton; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Button_GetAlignWithMargins(AObj: TButton): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Button_SetAlignWithMargins(AObj: TButton; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Button_GetLeft(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Button_SetLeft(AObj: TButton; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Button_GetTop(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Button_SetTop(AObj: TButton; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Button_GetWidth(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Button_SetWidth(AObj: TButton; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Button_GetHeight(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Button_SetHeight(AObj: TButton; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Button_GetCursor(AObj: TButton): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Button_SetCursor(AObj: TButton; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Button_GetHint(AObj: TButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Button_SetHint(AObj: TButton; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Button_GetMargins(AObj: TButton): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Button_SetMargins(AObj: TButton; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Button_GetCustomHint(AObj: TButton): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Button_SetCustomHint(AObj: TButton; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Button_GetComponentCount(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Button_GetComponentIndex(AObj: TButton): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Button_SetComponentIndex(AObj: TButton; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Button_GetOwner(AObj: TButton): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Button_GetName(AObj: TButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Button_SetName(AObj: TButton; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Button_GetTag(AObj: TButton): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Button_SetTag(AObj: TButton; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Button_GetControls(AObj: TButton; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function Button_GetComponents(AObj: TButton; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Button_Create,
  Button_Free,
  Button_Click,
  Button_CanFocus,
  Button_FlipChildren,
  Button_Focused,
  Button_HandleAllocated,
  Button_Invalidate,
  Button_Realign,
  Button_Repaint,
  Button_ScaleBy,
  Button_SetBounds,
  Button_SetFocus,
  Button_Update,
  Button_BringToFront,
  Button_HasParent,
  Button_Hide,
  Button_Perform,
  Button_Refresh,
  Button_SendToBack,
  Button_Show,
  Button_GetTextBuf,
  Button_GetTextLen,
  Button_FindComponent,
  Button_GetNamePath,
  Button_Assign,
  Button_ClassName,
  Button_Equals,
  Button_GetHashCode,
  Button_ToString,
  Button_GetAction,
  Button_SetAction,
  Button_GetAlign,
  Button_SetAlign,
  Button_GetAnchors,
  Button_SetAnchors,
  Button_GetBiDiMode,
  Button_SetBiDiMode,
  Button_GetCancel,
  Button_SetCancel,
  Button_GetCaption,
  Button_SetCaption,
  Button_GetCommandLinkHint,
  Button_SetCommandLinkHint,
  Button_GetDefault,
  Button_SetDefault,
  Button_GetDisabledImageIndex,
  Button_SetDisabledImageIndex,
  Button_GetDoubleBuffered,
  Button_SetDoubleBuffered,
  Button_GetElevationRequired,
  Button_SetElevationRequired,
  Button_GetEnabled,
  Button_SetEnabled,
  Button_GetFont,
  Button_SetFont,
  Button_GetHotImageIndex,
  Button_SetHotImageIndex,
  Button_GetImageAlignment,
  Button_SetImageAlignment,
  Button_GetImageIndex,
  Button_SetImageIndex,
  Button_GetImages,
  Button_SetImages,
  Button_GetModalResult,
  Button_SetModalResult,
  Button_GetParentDoubleBuffered,
  Button_SetParentDoubleBuffered,
  Button_GetParentFont,
  Button_SetParentFont,
  Button_GetParentShowHint,
  Button_SetParentShowHint,
  Button_GetPopupMenu,
  Button_SetPopupMenu,
  Button_GetPressedImageIndex,
  Button_SetPressedImageIndex,
  Button_GetSelectedImageIndex,
  Button_SetSelectedImageIndex,
  Button_GetShowHint,
  Button_SetShowHint,
  Button_GetStyle,
  Button_SetStyle,
  Button_GetStylusHotImageIndex,
  Button_SetStylusHotImageIndex,
  Button_GetTabOrder,
  Button_SetTabOrder,
  Button_GetTabStop,
  Button_SetTabStop,
  Button_GetVisible,
  Button_SetVisible,
  Button_GetWordWrap,
  Button_SetWordWrap,
  Button_GetStyleElements,
  Button_SetStyleElements,
  Button_SetOnClick,
  Button_SetOnEnter,
  Button_SetOnExit,
  Button_SetOnKeyDown,
  Button_SetOnKeyPress,
  Button_SetOnKeyUp,
  Button_SetOnMouseDown,
  Button_SetOnMouseEnter,
  Button_SetOnMouseLeave,
  Button_SetOnMouseMove,
  Button_SetOnMouseUp,
  Button_GetBrush,
  Button_GetControlCount,
  Button_GetHandle,
  Button_GetParentWindow,
  Button_SetParentWindow,
  Button_GetBoundsRect,
  Button_SetBoundsRect,
  Button_GetClientHeight,
  Button_SetClientHeight,
  Button_GetClientRect,
  Button_GetClientWidth,
  Button_SetClientWidth,
  Button_GetExplicitLeft,
  Button_GetExplicitTop,
  Button_GetExplicitWidth,
  Button_GetExplicitHeight,
  Button_GetParent,
  Button_SetParent,
  Button_GetAlignWithMargins,
  Button_SetAlignWithMargins,
  Button_GetLeft,
  Button_SetLeft,
  Button_GetTop,
  Button_SetTop,
  Button_GetWidth,
  Button_SetWidth,
  Button_GetHeight,
  Button_SetHeight,
  Button_GetCursor,
  Button_SetCursor,
  Button_GetHint,
  Button_SetHint,
  Button_GetMargins,
  Button_SetMargins,
  Button_GetCustomHint,
  Button_SetCustomHint,
  Button_GetComponentCount,
  Button_GetComponentIndex,
  Button_SetComponentIndex,
  Button_GetOwner,
  Button_GetName,
  Button_SetName,
  Button_GetTag,
  Button_SetTag,
  Button_GetControls,
  Button_GetComponents;

