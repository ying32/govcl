
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ProgressBar_Create(AOwner: TComponent): TProgressBar; stdcall;
begin
  Result :=  TProgressBar.Create(AOwner);
end;

procedure ProgressBar_Free(AObj: TProgressBar); stdcall;
begin
  AObj.Free;
end;

procedure ProgressBar_StepIt(AObj: TProgressBar); stdcall;
begin
  AObj.StepIt;
end;

procedure ProgressBar_StepBy(AObj: TProgressBar; Delta: Integer); stdcall;
begin
  AObj.StepBy(Delta);
end;

function ProgressBar_CanFocus(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure ProgressBar_FlipChildren(AObj: TProgressBar; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function ProgressBar_Focused(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function ProgressBar_HandleAllocated(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure ProgressBar_Invalidate(AObj: TProgressBar); stdcall;
begin
  AObj.Invalidate;
end;

procedure ProgressBar_Realign(AObj: TProgressBar); stdcall;
begin
  AObj.Realign;
end;

procedure ProgressBar_Repaint(AObj: TProgressBar); stdcall;
begin
  AObj.Repaint;
end;

procedure ProgressBar_ScaleBy(AObj: TProgressBar; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure ProgressBar_SetBounds(AObj: TProgressBar; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ProgressBar_SetFocus(AObj: TProgressBar); stdcall;
begin
  AObj.SetFocus;
end;

procedure ProgressBar_Update(AObj: TProgressBar); stdcall;
begin
  AObj.Update;
end;

procedure ProgressBar_BringToFront(AObj: TProgressBar); stdcall;
begin
  AObj.BringToFront;
end;

function ProgressBar_HasParent(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ProgressBar_Hide(AObj: TProgressBar); stdcall;
begin
  AObj.Hide;
end;

function ProgressBar_Perform(AObj: TProgressBar; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ProgressBar_Refresh(AObj: TProgressBar); stdcall;
begin
  AObj.Refresh;
end;

procedure ProgressBar_SendToBack(AObj: TProgressBar); stdcall;
begin
  AObj.SendToBack;
end;

procedure ProgressBar_Show(AObj: TProgressBar); stdcall;
begin
  AObj.Show;
end;

function ProgressBar_GetTextBuf(AObj: TProgressBar; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ProgressBar_GetTextLen(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function ProgressBar_FindComponent(AObj: TProgressBar; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ProgressBar_GetNamePath(AObj: TProgressBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ProgressBar_Assign(AObj: TProgressBar; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ProgressBar_ClassName(AObj: TProgressBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ProgressBar_Equals(AObj: TProgressBar; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ProgressBar_GetHashCode(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ProgressBar_ToString(AObj: TProgressBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ProgressBar_GetAlign(AObj: TProgressBar): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ProgressBar_SetAlign(AObj: TProgressBar; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ProgressBar_GetAnchors(AObj: TProgressBar): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ProgressBar_SetAnchors(AObj: TProgressBar; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ProgressBar_GetBorderWidth(AObj: TProgressBar): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure ProgressBar_SetBorderWidth(AObj: TProgressBar; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function ProgressBar_GetDoubleBuffered(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ProgressBar_SetDoubleBuffered(AObj: TProgressBar; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function ProgressBar_GetEnabled(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ProgressBar_SetEnabled(AObj: TProgressBar; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ProgressBar_GetHint(AObj: TProgressBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ProgressBar_SetHint(AObj: TProgressBar; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ProgressBar_GetMin(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Min;
end;

procedure ProgressBar_SetMin(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.Min := AValue;
end;

function ProgressBar_GetMax(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Max;
end;

procedure ProgressBar_SetMax(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.Max := AValue;
end;

function ProgressBar_GetOrientation(AObj: TProgressBar): TProgressBarOrientation; stdcall;
begin
  Result :=  AObj.Orientation;
end;

procedure ProgressBar_SetOrientation(AObj: TProgressBar; AValue: TProgressBarOrientation); stdcall;
begin
  AObj.Orientation := AValue;
end;

function ProgressBar_GetParentDoubleBuffered(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure ProgressBar_SetParentDoubleBuffered(AObj: TProgressBar; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function ProgressBar_GetParentShowHint(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ProgressBar_SetParentShowHint(AObj: TProgressBar; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ProgressBar_GetPopupMenu(AObj: TProgressBar): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ProgressBar_SetPopupMenu(AObj: TProgressBar; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ProgressBar_GetPosition(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Position;
end;

procedure ProgressBar_SetPosition(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.Position := AValue;
end;

function ProgressBar_GetSmooth(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.Smooth;
end;

procedure ProgressBar_SetSmooth(AObj: TProgressBar; AValue: LongBool); stdcall;
begin
  AObj.Smooth := AValue;
end;

function ProgressBar_GetStyle(AObj: TProgressBar): TProgressBarStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure ProgressBar_SetStyle(AObj: TProgressBar; AValue: TProgressBarStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function ProgressBar_GetMarqueeInterval(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.MarqueeInterval;
end;

procedure ProgressBar_SetMarqueeInterval(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.MarqueeInterval := AValue;
end;

function ProgressBar_GetBarColor(AObj: TProgressBar): TColor; stdcall;
begin
  Result :=  AObj.BarColor;
end;

procedure ProgressBar_SetBarColor(AObj: TProgressBar; AValue: TColor); stdcall;
begin
  AObj.BarColor := AValue;
end;

function ProgressBar_GetBackgroundColor(AObj: TProgressBar): TColor; stdcall;
begin
  Result :=  AObj.BackgroundColor;
end;

procedure ProgressBar_SetBackgroundColor(AObj: TProgressBar; AValue: TColor); stdcall;
begin
  AObj.BackgroundColor := AValue;
end;

function ProgressBar_GetSmoothReverse(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.SmoothReverse;
end;

procedure ProgressBar_SetSmoothReverse(AObj: TProgressBar; AValue: LongBool); stdcall;
begin
  AObj.SmoothReverse := AValue;
end;

function ProgressBar_GetStep(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Step;
end;

procedure ProgressBar_SetStep(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.Step := AValue;
end;

function ProgressBar_GetState(AObj: TProgressBar): TProgressBarState; stdcall;
begin
  Result :=  AObj.State;
end;

procedure ProgressBar_SetState(AObj: TProgressBar; AValue: TProgressBarState); stdcall;
begin
  AObj.State := AValue;
end;

function ProgressBar_GetShowHint(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ProgressBar_SetShowHint(AObj: TProgressBar; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ProgressBar_GetTabOrder(AObj: TProgressBar): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure ProgressBar_SetTabOrder(AObj: TProgressBar; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function ProgressBar_GetTabStop(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure ProgressBar_SetTabStop(AObj: TProgressBar; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function ProgressBar_GetVisible(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ProgressBar_SetVisible(AObj: TProgressBar; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function ProgressBar_GetStyleElements(AObj: TProgressBar): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ProgressBar_SetStyleElements(AObj: TProgressBar; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure ProgressBar_SetOnEnter(AObj: TProgressBar; AEventId: NativeUInt); stdcall;
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

procedure ProgressBar_SetOnExit(AObj: TProgressBar; AEventId: NativeUInt); stdcall;
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

procedure ProgressBar_SetOnMouseDown(AObj: TProgressBar; AEventId: NativeUInt); stdcall;
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

procedure ProgressBar_SetOnMouseEnter(AObj: TProgressBar; AEventId: NativeUInt); stdcall;
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

procedure ProgressBar_SetOnMouseLeave(AObj: TProgressBar; AEventId: NativeUInt); stdcall;
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

procedure ProgressBar_SetOnMouseMove(AObj: TProgressBar; AEventId: NativeUInt); stdcall;
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

procedure ProgressBar_SetOnMouseUp(AObj: TProgressBar; AEventId: NativeUInt); stdcall;
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

function ProgressBar_GetBrush(AObj: TProgressBar): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function ProgressBar_GetControlCount(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function ProgressBar_GetHandle(AObj: TProgressBar): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function ProgressBar_GetParentWindow(AObj: TProgressBar): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ProgressBar_SetParentWindow(AObj: TProgressBar; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function ProgressBar_GetAction(AObj: TProgressBar): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ProgressBar_SetAction(AObj: TProgressBar; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function ProgressBar_GetBiDiMode(AObj: TProgressBar): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ProgressBar_SetBiDiMode(AObj: TProgressBar; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure ProgressBar_GetBoundsRect(AObj: TProgressBar; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ProgressBar_SetBoundsRect(AObj: TProgressBar; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ProgressBar_GetClientHeight(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ProgressBar_SetClientHeight(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ProgressBar_GetClientRect(AObj: TProgressBar; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ProgressBar_GetClientWidth(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ProgressBar_SetClientWidth(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ProgressBar_GetExplicitLeft(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ProgressBar_GetExplicitTop(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ProgressBar_GetExplicitWidth(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ProgressBar_GetExplicitHeight(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ProgressBar_GetParent(AObj: TProgressBar): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ProgressBar_SetParent(AObj: TProgressBar; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ProgressBar_GetAlignWithMargins(AObj: TProgressBar): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ProgressBar_SetAlignWithMargins(AObj: TProgressBar; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ProgressBar_GetLeft(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ProgressBar_SetLeft(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ProgressBar_GetTop(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ProgressBar_SetTop(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ProgressBar_GetWidth(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ProgressBar_SetWidth(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ProgressBar_GetHeight(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ProgressBar_SetHeight(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ProgressBar_GetCursor(AObj: TProgressBar): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ProgressBar_SetCursor(AObj: TProgressBar; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ProgressBar_GetMargins(AObj: TProgressBar): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ProgressBar_SetMargins(AObj: TProgressBar; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ProgressBar_GetCustomHint(AObj: TProgressBar): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ProgressBar_SetCustomHint(AObj: TProgressBar; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ProgressBar_GetComponentCount(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ProgressBar_GetComponentIndex(AObj: TProgressBar): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ProgressBar_SetComponentIndex(AObj: TProgressBar; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ProgressBar_GetOwner(AObj: TProgressBar): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ProgressBar_GetName(AObj: TProgressBar): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ProgressBar_SetName(AObj: TProgressBar; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ProgressBar_GetTag(AObj: TProgressBar): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ProgressBar_SetTag(AObj: TProgressBar; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ProgressBar_GetControls(AObj: TProgressBar; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function ProgressBar_GetComponents(AObj: TProgressBar; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ProgressBar_Create,
  ProgressBar_Free,
  ProgressBar_StepIt,
  ProgressBar_StepBy,
  ProgressBar_CanFocus,
  ProgressBar_FlipChildren,
  ProgressBar_Focused,
  ProgressBar_HandleAllocated,
  ProgressBar_Invalidate,
  ProgressBar_Realign,
  ProgressBar_Repaint,
  ProgressBar_ScaleBy,
  ProgressBar_SetBounds,
  ProgressBar_SetFocus,
  ProgressBar_Update,
  ProgressBar_BringToFront,
  ProgressBar_HasParent,
  ProgressBar_Hide,
  ProgressBar_Perform,
  ProgressBar_Refresh,
  ProgressBar_SendToBack,
  ProgressBar_Show,
  ProgressBar_GetTextBuf,
  ProgressBar_GetTextLen,
  ProgressBar_FindComponent,
  ProgressBar_GetNamePath,
  ProgressBar_Assign,
  ProgressBar_ClassName,
  ProgressBar_Equals,
  ProgressBar_GetHashCode,
  ProgressBar_ToString,
  ProgressBar_GetAlign,
  ProgressBar_SetAlign,
  ProgressBar_GetAnchors,
  ProgressBar_SetAnchors,
  ProgressBar_GetBorderWidth,
  ProgressBar_SetBorderWidth,
  ProgressBar_GetDoubleBuffered,
  ProgressBar_SetDoubleBuffered,
  ProgressBar_GetEnabled,
  ProgressBar_SetEnabled,
  ProgressBar_GetHint,
  ProgressBar_SetHint,
  ProgressBar_GetMin,
  ProgressBar_SetMin,
  ProgressBar_GetMax,
  ProgressBar_SetMax,
  ProgressBar_GetOrientation,
  ProgressBar_SetOrientation,
  ProgressBar_GetParentDoubleBuffered,
  ProgressBar_SetParentDoubleBuffered,
  ProgressBar_GetParentShowHint,
  ProgressBar_SetParentShowHint,
  ProgressBar_GetPopupMenu,
  ProgressBar_SetPopupMenu,
  ProgressBar_GetPosition,
  ProgressBar_SetPosition,
  ProgressBar_GetSmooth,
  ProgressBar_SetSmooth,
  ProgressBar_GetStyle,
  ProgressBar_SetStyle,
  ProgressBar_GetMarqueeInterval,
  ProgressBar_SetMarqueeInterval,
  ProgressBar_GetBarColor,
  ProgressBar_SetBarColor,
  ProgressBar_GetBackgroundColor,
  ProgressBar_SetBackgroundColor,
  ProgressBar_GetSmoothReverse,
  ProgressBar_SetSmoothReverse,
  ProgressBar_GetStep,
  ProgressBar_SetStep,
  ProgressBar_GetState,
  ProgressBar_SetState,
  ProgressBar_GetShowHint,
  ProgressBar_SetShowHint,
  ProgressBar_GetTabOrder,
  ProgressBar_SetTabOrder,
  ProgressBar_GetTabStop,
  ProgressBar_SetTabStop,
  ProgressBar_GetVisible,
  ProgressBar_SetVisible,
  ProgressBar_GetStyleElements,
  ProgressBar_SetStyleElements,
  ProgressBar_SetOnEnter,
  ProgressBar_SetOnExit,
  ProgressBar_SetOnMouseDown,
  ProgressBar_SetOnMouseEnter,
  ProgressBar_SetOnMouseLeave,
  ProgressBar_SetOnMouseMove,
  ProgressBar_SetOnMouseUp,
  ProgressBar_GetBrush,
  ProgressBar_GetControlCount,
  ProgressBar_GetHandle,
  ProgressBar_GetParentWindow,
  ProgressBar_SetParentWindow,
  ProgressBar_GetAction,
  ProgressBar_SetAction,
  ProgressBar_GetBiDiMode,
  ProgressBar_SetBiDiMode,
  ProgressBar_GetBoundsRect,
  ProgressBar_SetBoundsRect,
  ProgressBar_GetClientHeight,
  ProgressBar_SetClientHeight,
  ProgressBar_GetClientRect,
  ProgressBar_GetClientWidth,
  ProgressBar_SetClientWidth,
  ProgressBar_GetExplicitLeft,
  ProgressBar_GetExplicitTop,
  ProgressBar_GetExplicitWidth,
  ProgressBar_GetExplicitHeight,
  ProgressBar_GetParent,
  ProgressBar_SetParent,
  ProgressBar_GetAlignWithMargins,
  ProgressBar_SetAlignWithMargins,
  ProgressBar_GetLeft,
  ProgressBar_SetLeft,
  ProgressBar_GetTop,
  ProgressBar_SetTop,
  ProgressBar_GetWidth,
  ProgressBar_SetWidth,
  ProgressBar_GetHeight,
  ProgressBar_SetHeight,
  ProgressBar_GetCursor,
  ProgressBar_SetCursor,
  ProgressBar_GetMargins,
  ProgressBar_SetMargins,
  ProgressBar_GetCustomHint,
  ProgressBar_SetCustomHint,
  ProgressBar_GetComponentCount,
  ProgressBar_GetComponentIndex,
  ProgressBar_SetComponentIndex,
  ProgressBar_GetOwner,
  ProgressBar_GetName,
  ProgressBar_SetName,
  ProgressBar_GetTag,
  ProgressBar_SetTag,
  ProgressBar_GetControls,
  ProgressBar_GetComponents;

