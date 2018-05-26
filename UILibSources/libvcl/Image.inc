
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Image_Create(AOwner: TComponent): TImage; stdcall;
begin
  Result :=  TImage.Create(AOwner);
end;

procedure Image_Free(AObj: TImage); stdcall;
begin
  AObj.Free;
end;

procedure Image_BringToFront(AObj: TImage); stdcall;
begin
  AObj.BringToFront;
end;

function Image_HasParent(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Image_Hide(AObj: TImage); stdcall;
begin
  AObj.Hide;
end;

procedure Image_Invalidate(AObj: TImage); stdcall;
begin
  AObj.Invalidate;
end;

function Image_Perform(AObj: TImage; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Image_Refresh(AObj: TImage); stdcall;
begin
  AObj.Refresh;
end;

procedure Image_Repaint(AObj: TImage); stdcall;
begin
  AObj.Repaint;
end;

procedure Image_SendToBack(AObj: TImage); stdcall;
begin
  AObj.SendToBack;
end;

procedure Image_SetBounds(AObj: TImage; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Image_Show(AObj: TImage); stdcall;
begin
  AObj.Show;
end;

procedure Image_Update(AObj: TImage); stdcall;
begin
  AObj.Update;
end;

function Image_GetTextBuf(AObj: TImage; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Image_GetTextLen(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function Image_FindComponent(AObj: TImage; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Image_GetNamePath(AObj: TImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Image_Assign(AObj: TImage; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Image_ClassName(AObj: TImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Image_Equals(AObj: TImage; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Image_GetHashCode(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Image_ToString(AObj: TImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Image_GetCanvas(AObj: TImage): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function Image_GetAlign(AObj: TImage): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Image_SetAlign(AObj: TImage; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Image_GetAnchors(AObj: TImage): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Image_SetAnchors(AObj: TImage; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Image_GetAutoSize(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure Image_SetAutoSize(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function Image_GetCenter(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.Center;
end;

procedure Image_SetCenter(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.Center := AValue;
end;

function Image_GetEnabled(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Image_SetEnabled(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Image_GetIncrementalDisplay(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.IncrementalDisplay;
end;

procedure Image_SetIncrementalDisplay(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.IncrementalDisplay := AValue;
end;

function Image_GetParentShowHint(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Image_SetParentShowHint(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function Image_GetPicture(AObj: TImage): TPicture; stdcall;
begin
  Result :=  AObj.Picture;
end;

procedure Image_SetPicture(AObj: TImage; AValue: TPicture); stdcall;
begin
  AObj.Picture := AValue;
end;

function Image_GetPopupMenu(AObj: TImage): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Image_SetPopupMenu(AObj: TImage; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function Image_GetProportional(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.Proportional;
end;

procedure Image_SetProportional(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.Proportional := AValue;
end;

function Image_GetShowHint(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Image_SetShowHint(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Image_GetStretch(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.Stretch;
end;

procedure Image_SetStretch(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.Stretch := AValue;
end;

function Image_GetTransparent(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure Image_SetTransparent(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function Image_GetVisible(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Image_SetVisible(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

procedure Image_SetOnClick(AObj: TImage; AEventId: NativeUInt); stdcall;
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

procedure Image_SetOnDblClick(AObj: TImage; AEventId: NativeUInt); stdcall;
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

procedure Image_SetOnMouseDown(AObj: TImage; AEventId: NativeUInt); stdcall;
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

procedure Image_SetOnMouseEnter(AObj: TImage; AEventId: NativeUInt); stdcall;
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

procedure Image_SetOnMouseLeave(AObj: TImage; AEventId: NativeUInt); stdcall;
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

procedure Image_SetOnMouseMove(AObj: TImage; AEventId: NativeUInt); stdcall;
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

procedure Image_SetOnMouseUp(AObj: TImage; AEventId: NativeUInt); stdcall;
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

function Image_GetAction(AObj: TImage): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Image_SetAction(AObj: TImage; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function Image_GetBiDiMode(AObj: TImage): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Image_SetBiDiMode(AObj: TImage; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure Image_GetBoundsRect(AObj: TImage; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Image_SetBoundsRect(AObj: TImage; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Image_GetClientHeight(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Image_SetClientHeight(AObj: TImage; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Image_GetClientRect(AObj: TImage; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Image_GetClientWidth(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Image_SetClientWidth(AObj: TImage; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Image_GetExplicitLeft(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Image_GetExplicitTop(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Image_GetExplicitWidth(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Image_GetExplicitHeight(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Image_GetParent(AObj: TImage): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Image_SetParent(AObj: TImage; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Image_GetStyleElements(AObj: TImage): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Image_SetStyleElements(AObj: TImage; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function Image_GetAlignWithMargins(AObj: TImage): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Image_SetAlignWithMargins(AObj: TImage; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Image_GetLeft(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Image_SetLeft(AObj: TImage; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Image_GetTop(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Image_SetTop(AObj: TImage; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Image_GetWidth(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Image_SetWidth(AObj: TImage; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Image_GetHeight(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Image_SetHeight(AObj: TImage; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Image_GetCursor(AObj: TImage): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Image_SetCursor(AObj: TImage; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Image_GetHint(AObj: TImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Image_SetHint(AObj: TImage; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Image_GetMargins(AObj: TImage): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Image_SetMargins(AObj: TImage; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Image_GetCustomHint(AObj: TImage): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Image_SetCustomHint(AObj: TImage; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Image_GetComponentCount(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Image_GetComponentIndex(AObj: TImage): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Image_SetComponentIndex(AObj: TImage; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Image_GetOwner(AObj: TImage): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Image_GetName(AObj: TImage): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Image_SetName(AObj: TImage; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Image_GetTag(AObj: TImage): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Image_SetTag(AObj: TImage; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Image_GetComponents(AObj: TImage; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Image_Create,
  Image_Free,
  Image_BringToFront,
  Image_HasParent,
  Image_Hide,
  Image_Invalidate,
  Image_Perform,
  Image_Refresh,
  Image_Repaint,
  Image_SendToBack,
  Image_SetBounds,
  Image_Show,
  Image_Update,
  Image_GetTextBuf,
  Image_GetTextLen,
  Image_FindComponent,
  Image_GetNamePath,
  Image_Assign,
  Image_ClassName,
  Image_Equals,
  Image_GetHashCode,
  Image_ToString,
  Image_GetCanvas,
  Image_GetAlign,
  Image_SetAlign,
  Image_GetAnchors,
  Image_SetAnchors,
  Image_GetAutoSize,
  Image_SetAutoSize,
  Image_GetCenter,
  Image_SetCenter,
  Image_GetEnabled,
  Image_SetEnabled,
  Image_GetIncrementalDisplay,
  Image_SetIncrementalDisplay,
  Image_GetParentShowHint,
  Image_SetParentShowHint,
  Image_GetPicture,
  Image_SetPicture,
  Image_GetPopupMenu,
  Image_SetPopupMenu,
  Image_GetProportional,
  Image_SetProportional,
  Image_GetShowHint,
  Image_SetShowHint,
  Image_GetStretch,
  Image_SetStretch,
  Image_GetTransparent,
  Image_SetTransparent,
  Image_GetVisible,
  Image_SetVisible,
  Image_SetOnClick,
  Image_SetOnDblClick,
  Image_SetOnMouseDown,
  Image_SetOnMouseEnter,
  Image_SetOnMouseLeave,
  Image_SetOnMouseMove,
  Image_SetOnMouseUp,
  Image_GetAction,
  Image_SetAction,
  Image_GetBiDiMode,
  Image_SetBiDiMode,
  Image_GetBoundsRect,
  Image_SetBoundsRect,
  Image_GetClientHeight,
  Image_SetClientHeight,
  Image_GetClientRect,
  Image_GetClientWidth,
  Image_SetClientWidth,
  Image_GetExplicitLeft,
  Image_GetExplicitTop,
  Image_GetExplicitWidth,
  Image_GetExplicitHeight,
  Image_GetParent,
  Image_SetParent,
  Image_GetStyleElements,
  Image_SetStyleElements,
  Image_GetAlignWithMargins,
  Image_SetAlignWithMargins,
  Image_GetLeft,
  Image_SetLeft,
  Image_GetTop,
  Image_SetTop,
  Image_GetWidth,
  Image_SetWidth,
  Image_GetHeight,
  Image_SetHeight,
  Image_GetCursor,
  Image_SetCursor,
  Image_GetHint,
  Image_SetHint,
  Image_GetMargins,
  Image_SetMargins,
  Image_GetCustomHint,
  Image_SetCustomHint,
  Image_GetComponentCount,
  Image_GetComponentIndex,
  Image_SetComponentIndex,
  Image_GetOwner,
  Image_GetName,
  Image_SetName,
  Image_GetTag,
  Image_SetTag,
  Image_GetComponents;

