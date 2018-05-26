
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ProgressBar_Create(AOwner: TComponent): TProgressBar; extdecl;
begin
  Result :=  TProgressBar.Create(AOwner);
end;

procedure ProgressBar_Free(AObj: TProgressBar); extdecl;
begin
  AObj.Free;
end;

procedure ProgressBar_StepIt(AObj: TProgressBar); extdecl;
begin
  AObj.StepIt;
end;

procedure ProgressBar_StepBy(AObj: TProgressBar; Delta: Integer); extdecl;
begin
  AObj.StepBy(Delta);
end;

function ProgressBar_CanFocus(AObj: TProgressBar): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure ProgressBar_FlipChildren(AObj: TProgressBar; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function ProgressBar_Focused(AObj: TProgressBar): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure ProgressBar_Invalidate(AObj: TProgressBar); extdecl;
begin
  AObj.Invalidate;
end;

procedure ProgressBar_Realign(AObj: TProgressBar); extdecl;
begin
  AObj.Realign;
end;

procedure ProgressBar_Repaint(AObj: TProgressBar); extdecl;
begin
  AObj.Repaint;
end;

procedure ProgressBar_ScaleBy(AObj: TProgressBar; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure ProgressBar_SetBounds(AObj: TProgressBar; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ProgressBar_SetFocus(AObj: TProgressBar); extdecl;
begin
  AObj.SetFocus;
end;

procedure ProgressBar_Update(AObj: TProgressBar); extdecl;
begin
  AObj.Update;
end;

procedure ProgressBar_BringToFront(AObj: TProgressBar); extdecl;
begin
  AObj.BringToFront;
end;

function ProgressBar_HasParent(AObj: TProgressBar): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ProgressBar_Hide(AObj: TProgressBar); extdecl;
begin
  AObj.Hide;
end;

function ProgressBar_Perform(AObj: TProgressBar; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ProgressBar_Refresh(AObj: TProgressBar); extdecl;
begin
  AObj.Refresh;
end;

procedure ProgressBar_SendToBack(AObj: TProgressBar); extdecl;
begin
  AObj.SendToBack;
end;

procedure ProgressBar_Show(AObj: TProgressBar); extdecl;
begin
  AObj.Show;
end;

function ProgressBar_GetTextBuf(AObj: TProgressBar; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ProgressBar_GetTextLen(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function ProgressBar_FindComponent(AObj: TProgressBar; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ProgressBar_GetNamePath(AObj: TProgressBar): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ProgressBar_Assign(AObj: TProgressBar; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ProgressBar_ClassName(AObj: TProgressBar): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ProgressBar_Equals(AObj: TProgressBar; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ProgressBar_GetHashCode(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ProgressBar_ToString(AObj: TProgressBar): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ProgressBar_GetAlign(AObj: TProgressBar): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure ProgressBar_SetAlign(AObj: TProgressBar; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function ProgressBar_GetAnchors(AObj: TProgressBar): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure ProgressBar_SetAnchors(AObj: TProgressBar; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function ProgressBar_GetBorderWidth(AObj: TProgressBar): TBorderWidth; extdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure ProgressBar_SetBorderWidth(AObj: TProgressBar; AValue: TBorderWidth); extdecl;
begin
  AObj.BorderWidth := AValue;
end;

function ProgressBar_GetDoubleBuffered(AObj: TProgressBar): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure ProgressBar_SetDoubleBuffered(AObj: TProgressBar; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function ProgressBar_GetEnabled(AObj: TProgressBar): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure ProgressBar_SetEnabled(AObj: TProgressBar; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function ProgressBar_GetHint(AObj: TProgressBar): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ProgressBar_SetHint(AObj: TProgressBar; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function ProgressBar_GetMin(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.Min;
end;

procedure ProgressBar_SetMin(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.Min := AValue;
end;

function ProgressBar_GetMax(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.Max;
end;

procedure ProgressBar_SetMax(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.Max := AValue;
end;

function ProgressBar_GetOrientation(AObj: TProgressBar): TProgressBarOrientation; extdecl;
begin
  Result :=  AObj.Orientation;
end;

procedure ProgressBar_SetOrientation(AObj: TProgressBar; AValue: TProgressBarOrientation); extdecl;
begin
  AObj.Orientation := AValue;
end;

function ProgressBar_GetParentShowHint(AObj: TProgressBar): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ProgressBar_SetParentShowHint(AObj: TProgressBar; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function ProgressBar_GetPopupMenu(AObj: TProgressBar): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ProgressBar_SetPopupMenu(AObj: TProgressBar; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function ProgressBar_GetPosition(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.Position;
end;

procedure ProgressBar_SetPosition(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.Position := AValue;
end;

function ProgressBar_GetSmooth(AObj: TProgressBar): LongBool; extdecl;
begin
  Result :=  AObj.Smooth;
end;

procedure ProgressBar_SetSmooth(AObj: TProgressBar; AValue: LongBool); extdecl;
begin
  AObj.Smooth := AValue;
end;

function ProgressBar_GetStyle(AObj: TProgressBar): TProgressBarStyle; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure ProgressBar_SetStyle(AObj: TProgressBar; AValue: TProgressBarStyle); extdecl;
begin
  AObj.Style := AValue;
end;

function ProgressBar_GetStep(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.Step;
end;

procedure ProgressBar_SetStep(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.Step := AValue;
end;

function ProgressBar_GetShowHint(AObj: TProgressBar): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure ProgressBar_SetShowHint(AObj: TProgressBar; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function ProgressBar_GetTabOrder(AObj: TProgressBar): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure ProgressBar_SetTabOrder(AObj: TProgressBar; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function ProgressBar_GetTabStop(AObj: TProgressBar): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure ProgressBar_SetTabStop(AObj: TProgressBar; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function ProgressBar_GetVisible(AObj: TProgressBar): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure ProgressBar_SetVisible(AObj: TProgressBar; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure ProgressBar_SetOnEnter(AObj: TProgressBar; AEventId: NativeUInt); extdecl;
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

procedure ProgressBar_SetOnExit(AObj: TProgressBar; AEventId: NativeUInt); extdecl;
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

procedure ProgressBar_SetOnMouseDown(AObj: TProgressBar; AEventId: NativeUInt); extdecl;
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

procedure ProgressBar_SetOnMouseEnter(AObj: TProgressBar; AEventId: NativeUInt); extdecl;
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

procedure ProgressBar_SetOnMouseLeave(AObj: TProgressBar; AEventId: NativeUInt); extdecl;
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

procedure ProgressBar_SetOnMouseMove(AObj: TProgressBar; AEventId: NativeUInt); extdecl;
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

procedure ProgressBar_SetOnMouseUp(AObj: TProgressBar; AEventId: NativeUInt); extdecl;
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

function ProgressBar_GetBrush(AObj: TProgressBar): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function ProgressBar_GetControlCount(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function ProgressBar_GetHandle(AObj: TProgressBar): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function ProgressBar_GetParentWindow(AObj: TProgressBar): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure ProgressBar_SetParentWindow(AObj: TProgressBar; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function ProgressBar_GetAction(AObj: TProgressBar): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure ProgressBar_SetAction(AObj: TProgressBar; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function ProgressBar_GetBiDiMode(AObj: TProgressBar): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ProgressBar_SetBiDiMode(AObj: TProgressBar; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure ProgressBar_GetBoundsRect(AObj: TProgressBar; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ProgressBar_SetBoundsRect(AObj: TProgressBar; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function ProgressBar_GetClientHeight(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ProgressBar_SetClientHeight(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure ProgressBar_GetClientRect(AObj: TProgressBar; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function ProgressBar_GetClientWidth(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ProgressBar_SetClientWidth(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function ProgressBar_GetParent(AObj: TProgressBar): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure ProgressBar_SetParent(AObj: TProgressBar; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function ProgressBar_GetLeft(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ProgressBar_SetLeft(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ProgressBar_GetTop(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ProgressBar_SetTop(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ProgressBar_GetWidth(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ProgressBar_SetWidth(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function ProgressBar_GetHeight(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ProgressBar_SetHeight(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ProgressBar_GetCursor(AObj: TProgressBar): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure ProgressBar_SetCursor(AObj: TProgressBar; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function ProgressBar_GetComponentCount(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ProgressBar_GetComponentIndex(AObj: TProgressBar): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ProgressBar_SetComponentIndex(AObj: TProgressBar; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ProgressBar_GetOwner(AObj: TProgressBar): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ProgressBar_GetName(AObj: TProgressBar): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ProgressBar_SetName(AObj: TProgressBar; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ProgressBar_GetTag(AObj: TProgressBar): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ProgressBar_SetTag(AObj: TProgressBar; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ProgressBar_GetControls(AObj: TProgressBar; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function ProgressBar_GetComponents(AObj: TProgressBar; AIndex: Integer): TComponent; extdecl;
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
  ProgressBar_GetStep,
  ProgressBar_SetStep,
  ProgressBar_GetShowHint,
  ProgressBar_SetShowHint,
  ProgressBar_GetTabOrder,
  ProgressBar_SetTabOrder,
  ProgressBar_GetTabStop,
  ProgressBar_SetTabStop,
  ProgressBar_GetVisible,
  ProgressBar_SetVisible,
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
  ProgressBar_GetParent,
  ProgressBar_SetParent,
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

