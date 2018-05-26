
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function PaintBox_Create(AOwner: TComponent): TPaintBox; extdecl;
begin
  Result :=  TPaintBox.Create(AOwner);
end;

procedure PaintBox_Free(AObj: TPaintBox); extdecl;
begin
  AObj.Free;
end;

procedure PaintBox_BringToFront(AObj: TPaintBox); extdecl;
begin
  AObj.BringToFront;
end;

function PaintBox_HasParent(AObj: TPaintBox): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure PaintBox_Hide(AObj: TPaintBox); extdecl;
begin
  AObj.Hide;
end;

procedure PaintBox_Invalidate(AObj: TPaintBox); extdecl;
begin
  AObj.Invalidate;
end;

function PaintBox_Perform(AObj: TPaintBox; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure PaintBox_Refresh(AObj: TPaintBox); extdecl;
begin
  AObj.Refresh;
end;

procedure PaintBox_Repaint(AObj: TPaintBox); extdecl;
begin
  AObj.Repaint;
end;

procedure PaintBox_SendToBack(AObj: TPaintBox); extdecl;
begin
  AObj.SendToBack;
end;

procedure PaintBox_SetBounds(AObj: TPaintBox; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure PaintBox_Show(AObj: TPaintBox); extdecl;
begin
  AObj.Show;
end;

procedure PaintBox_Update(AObj: TPaintBox); extdecl;
begin
  AObj.Update;
end;

function PaintBox_GetTextBuf(AObj: TPaintBox; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function PaintBox_GetTextLen(AObj: TPaintBox): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function PaintBox_FindComponent(AObj: TPaintBox; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function PaintBox_GetNamePath(AObj: TPaintBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure PaintBox_Assign(AObj: TPaintBox; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function PaintBox_ClassName(AObj: TPaintBox): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function PaintBox_Equals(AObj: TPaintBox; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function PaintBox_GetHashCode(AObj: TPaintBox): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function PaintBox_ToString(AObj: TPaintBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function PaintBox_GetCanvas(AObj: TPaintBox): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function PaintBox_GetAlign(AObj: TPaintBox): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure PaintBox_SetAlign(AObj: TPaintBox; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function PaintBox_GetAnchors(AObj: TPaintBox): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure PaintBox_SetAnchors(AObj: TPaintBox; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function PaintBox_GetColor(AObj: TPaintBox): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure PaintBox_SetColor(AObj: TPaintBox; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function PaintBox_GetEnabled(AObj: TPaintBox): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure PaintBox_SetEnabled(AObj: TPaintBox; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function PaintBox_GetFont(AObj: TPaintBox): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure PaintBox_SetFont(AObj: TPaintBox; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function PaintBox_GetParentColor(AObj: TPaintBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure PaintBox_SetParentColor(AObj: TPaintBox; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function PaintBox_GetParentFont(AObj: TPaintBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure PaintBox_SetParentFont(AObj: TPaintBox; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function PaintBox_GetParentShowHint(AObj: TPaintBox): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure PaintBox_SetParentShowHint(AObj: TPaintBox; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function PaintBox_GetPopupMenu(AObj: TPaintBox): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure PaintBox_SetPopupMenu(AObj: TPaintBox; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function PaintBox_GetShowHint(AObj: TPaintBox): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure PaintBox_SetShowHint(AObj: TPaintBox; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function PaintBox_GetVisible(AObj: TPaintBox): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure PaintBox_SetVisible(AObj: TPaintBox; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure PaintBox_SetOnClick(AObj: TPaintBox; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    EventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := @EventClass.OnClick;
  EventClass.Add(AObj, geClick, AEventId);
end;

procedure PaintBox_SetOnDblClick(AObj: TPaintBox; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnDblClick := nil;
    EventClass.Remove(AObj, geDblClick);
    Exit;
  end;
  AObj.OnDblClick := @EventClass.OnDblClick;
  EventClass.Add(AObj, geDblClick, AEventId);
end;

procedure PaintBox_SetOnMouseDown(AObj: TPaintBox; AEventId: NativeUInt); extdecl;
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

procedure PaintBox_SetOnMouseEnter(AObj: TPaintBox; AEventId: NativeUInt); extdecl;
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

procedure PaintBox_SetOnMouseLeave(AObj: TPaintBox; AEventId: NativeUInt); extdecl;
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

procedure PaintBox_SetOnMouseMove(AObj: TPaintBox; AEventId: NativeUInt); extdecl;
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

procedure PaintBox_SetOnMouseUp(AObj: TPaintBox; AEventId: NativeUInt); extdecl;
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

procedure PaintBox_SetOnPaint(AObj: TPaintBox; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnPaint := nil;
    EventClass.Remove(AObj, gePaint);
    Exit;
  end;
  AObj.OnPaint := @EventClass.OnPaint;
  EventClass.Add(AObj, gePaint, AEventId);
end;

function PaintBox_GetAction(AObj: TPaintBox): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure PaintBox_SetAction(AObj: TPaintBox; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function PaintBox_GetBiDiMode(AObj: TPaintBox): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure PaintBox_SetBiDiMode(AObj: TPaintBox; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure PaintBox_GetBoundsRect(AObj: TPaintBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure PaintBox_SetBoundsRect(AObj: TPaintBox; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function PaintBox_GetClientHeight(AObj: TPaintBox): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure PaintBox_SetClientHeight(AObj: TPaintBox; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure PaintBox_GetClientRect(AObj: TPaintBox; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function PaintBox_GetClientWidth(AObj: TPaintBox): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure PaintBox_SetClientWidth(AObj: TPaintBox; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function PaintBox_GetParent(AObj: TPaintBox): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure PaintBox_SetParent(AObj: TPaintBox; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function PaintBox_GetLeft(AObj: TPaintBox): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure PaintBox_SetLeft(AObj: TPaintBox; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function PaintBox_GetTop(AObj: TPaintBox): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure PaintBox_SetTop(AObj: TPaintBox; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function PaintBox_GetWidth(AObj: TPaintBox): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure PaintBox_SetWidth(AObj: TPaintBox; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function PaintBox_GetHeight(AObj: TPaintBox): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure PaintBox_SetHeight(AObj: TPaintBox; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function PaintBox_GetCursor(AObj: TPaintBox): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure PaintBox_SetCursor(AObj: TPaintBox; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function PaintBox_GetHint(AObj: TPaintBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure PaintBox_SetHint(AObj: TPaintBox; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function PaintBox_GetComponentCount(AObj: TPaintBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function PaintBox_GetComponentIndex(AObj: TPaintBox): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure PaintBox_SetComponentIndex(AObj: TPaintBox; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function PaintBox_GetOwner(AObj: TPaintBox): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function PaintBox_GetName(AObj: TPaintBox): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure PaintBox_SetName(AObj: TPaintBox; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function PaintBox_GetTag(AObj: TPaintBox): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure PaintBox_SetTag(AObj: TPaintBox; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function PaintBox_GetComponents(AObj: TPaintBox; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  PaintBox_Create,
  PaintBox_Free,
  PaintBox_BringToFront,
  PaintBox_HasParent,
  PaintBox_Hide,
  PaintBox_Invalidate,
  PaintBox_Perform,
  PaintBox_Refresh,
  PaintBox_Repaint,
  PaintBox_SendToBack,
  PaintBox_SetBounds,
  PaintBox_Show,
  PaintBox_Update,
  PaintBox_GetTextBuf,
  PaintBox_GetTextLen,
  PaintBox_FindComponent,
  PaintBox_GetNamePath,
  PaintBox_Assign,
  PaintBox_ClassName,
  PaintBox_Equals,
  PaintBox_GetHashCode,
  PaintBox_ToString,
  PaintBox_GetCanvas,
  PaintBox_GetAlign,
  PaintBox_SetAlign,
  PaintBox_GetAnchors,
  PaintBox_SetAnchors,
  PaintBox_GetColor,
  PaintBox_SetColor,
  PaintBox_GetEnabled,
  PaintBox_SetEnabled,
  PaintBox_GetFont,
  PaintBox_SetFont,
  PaintBox_GetParentColor,
  PaintBox_SetParentColor,
  PaintBox_GetParentFont,
  PaintBox_SetParentFont,
  PaintBox_GetParentShowHint,
  PaintBox_SetParentShowHint,
  PaintBox_GetPopupMenu,
  PaintBox_SetPopupMenu,
  PaintBox_GetShowHint,
  PaintBox_SetShowHint,
  PaintBox_GetVisible,
  PaintBox_SetVisible,
  PaintBox_SetOnClick,
  PaintBox_SetOnDblClick,
  PaintBox_SetOnMouseDown,
  PaintBox_SetOnMouseEnter,
  PaintBox_SetOnMouseLeave,
  PaintBox_SetOnMouseMove,
  PaintBox_SetOnMouseUp,
  PaintBox_SetOnPaint,
  PaintBox_GetAction,
  PaintBox_SetAction,
  PaintBox_GetBiDiMode,
  PaintBox_SetBiDiMode,
  PaintBox_GetBoundsRect,
  PaintBox_SetBoundsRect,
  PaintBox_GetClientHeight,
  PaintBox_SetClientHeight,
  PaintBox_GetClientRect,
  PaintBox_GetClientWidth,
  PaintBox_SetClientWidth,
  PaintBox_GetParent,
  PaintBox_SetParent,
  PaintBox_GetLeft,
  PaintBox_SetLeft,
  PaintBox_GetTop,
  PaintBox_SetTop,
  PaintBox_GetWidth,
  PaintBox_SetWidth,
  PaintBox_GetHeight,
  PaintBox_SetHeight,
  PaintBox_GetCursor,
  PaintBox_SetCursor,
  PaintBox_GetHint,
  PaintBox_SetHint,
  PaintBox_GetComponentCount,
  PaintBox_GetComponentIndex,
  PaintBox_SetComponentIndex,
  PaintBox_GetOwner,
  PaintBox_GetName,
  PaintBox_SetName,
  PaintBox_GetTag,
  PaintBox_SetTag,
  PaintBox_GetComponents;

