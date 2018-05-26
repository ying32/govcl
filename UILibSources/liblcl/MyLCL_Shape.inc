
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Shape_Create(AOwner: TComponent): TShape; extdecl;
begin
  Result :=  TShape.Create(AOwner);
end;

procedure Shape_Free(AObj: TShape); extdecl;
begin
  AObj.Free;
end;

procedure Shape_BringToFront(AObj: TShape); extdecl;
begin
  AObj.BringToFront;
end;

function Shape_HasParent(AObj: TShape): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Shape_Hide(AObj: TShape); extdecl;
begin
  AObj.Hide;
end;

procedure Shape_Invalidate(AObj: TShape); extdecl;
begin
  AObj.Invalidate;
end;

function Shape_Perform(AObj: TShape; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Shape_Refresh(AObj: TShape); extdecl;
begin
  AObj.Refresh;
end;

procedure Shape_Repaint(AObj: TShape); extdecl;
begin
  AObj.Repaint;
end;

procedure Shape_SendToBack(AObj: TShape); extdecl;
begin
  AObj.SendToBack;
end;

procedure Shape_SetBounds(AObj: TShape; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Shape_Show(AObj: TShape); extdecl;
begin
  AObj.Show;
end;

procedure Shape_Update(AObj: TShape); extdecl;
begin
  AObj.Update;
end;

function Shape_GetTextBuf(AObj: TShape; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Shape_GetTextLen(AObj: TShape): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function Shape_FindComponent(AObj: TShape; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Shape_GetNamePath(AObj: TShape): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Shape_Assign(AObj: TShape; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Shape_ClassName(AObj: TShape): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Shape_Equals(AObj: TShape; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Shape_GetHashCode(AObj: TShape): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Shape_ToString(AObj: TShape): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Shape_GetAlign(AObj: TShape): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure Shape_SetAlign(AObj: TShape; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function Shape_GetAnchors(AObj: TShape): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure Shape_SetAnchors(AObj: TShape; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function Shape_GetBrush(AObj: TShape): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

procedure Shape_SetBrush(AObj: TShape; AValue: TBrush); extdecl;
begin
  AObj.Brush := AValue;
end;

function Shape_GetEnabled(AObj: TShape): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Shape_SetEnabled(AObj: TShape; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Shape_GetParentShowHint(AObj: TShape): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Shape_SetParentShowHint(AObj: TShape; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function Shape_GetPen(AObj: TShape): TPen; extdecl;
begin
  Result :=  AObj.Pen;
end;

procedure Shape_SetPen(AObj: TShape; AValue: TPen); extdecl;
begin
  AObj.Pen := AValue;
end;

function Shape_GetShape(AObj: TShape): TShapeType; extdecl;
begin
  Result :=  AObj.Shape;
end;

procedure Shape_SetShape(AObj: TShape; AValue: TShapeType); extdecl;
begin
  AObj.Shape := AValue;
end;

function Shape_GetShowHint(AObj: TShape): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure Shape_SetShowHint(AObj: TShape; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function Shape_GetVisible(AObj: TShape): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure Shape_SetVisible(AObj: TShape; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure Shape_SetOnMouseDown(AObj: TShape; AEventId: NativeUInt); extdecl;
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

procedure Shape_SetOnMouseEnter(AObj: TShape; AEventId: NativeUInt); extdecl;
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

procedure Shape_SetOnMouseLeave(AObj: TShape; AEventId: NativeUInt); extdecl;
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

procedure Shape_SetOnMouseMove(AObj: TShape; AEventId: NativeUInt); extdecl;
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

procedure Shape_SetOnMouseUp(AObj: TShape; AEventId: NativeUInt); extdecl;
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

function Shape_GetAction(AObj: TShape): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure Shape_SetAction(AObj: TShape; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function Shape_GetBiDiMode(AObj: TShape): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Shape_SetBiDiMode(AObj: TShape; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure Shape_GetBoundsRect(AObj: TShape; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Shape_SetBoundsRect(AObj: TShape; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function Shape_GetClientHeight(AObj: TShape): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Shape_SetClientHeight(AObj: TShape; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure Shape_GetClientRect(AObj: TShape; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function Shape_GetClientWidth(AObj: TShape): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Shape_SetClientWidth(AObj: TShape; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function Shape_GetParent(AObj: TShape): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure Shape_SetParent(AObj: TShape; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function Shape_GetLeft(AObj: TShape): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure Shape_SetLeft(AObj: TShape; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function Shape_GetTop(AObj: TShape): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure Shape_SetTop(AObj: TShape; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function Shape_GetWidth(AObj: TShape): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Shape_SetWidth(AObj: TShape; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function Shape_GetHeight(AObj: TShape): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Shape_SetHeight(AObj: TShape; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Shape_GetCursor(AObj: TShape): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure Shape_SetCursor(AObj: TShape; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function Shape_GetHint(AObj: TShape): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Shape_SetHint(AObj: TShape; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function Shape_GetComponentCount(AObj: TShape): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Shape_GetComponentIndex(AObj: TShape): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Shape_SetComponentIndex(AObj: TShape; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Shape_GetOwner(AObj: TShape): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Shape_GetName(AObj: TShape): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Shape_SetName(AObj: TShape; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Shape_GetTag(AObj: TShape): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Shape_SetTag(AObj: TShape; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Shape_GetComponents(AObj: TShape; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Shape_Create,
  Shape_Free,
  Shape_BringToFront,
  Shape_HasParent,
  Shape_Hide,
  Shape_Invalidate,
  Shape_Perform,
  Shape_Refresh,
  Shape_Repaint,
  Shape_SendToBack,
  Shape_SetBounds,
  Shape_Show,
  Shape_Update,
  Shape_GetTextBuf,
  Shape_GetTextLen,
  Shape_FindComponent,
  Shape_GetNamePath,
  Shape_Assign,
  Shape_ClassName,
  Shape_Equals,
  Shape_GetHashCode,
  Shape_ToString,
  Shape_GetAlign,
  Shape_SetAlign,
  Shape_GetAnchors,
  Shape_SetAnchors,
  Shape_GetBrush,
  Shape_SetBrush,
  Shape_GetEnabled,
  Shape_SetEnabled,
  Shape_GetParentShowHint,
  Shape_SetParentShowHint,
  Shape_GetPen,
  Shape_SetPen,
  Shape_GetShape,
  Shape_SetShape,
  Shape_GetShowHint,
  Shape_SetShowHint,
  Shape_GetVisible,
  Shape_SetVisible,
  Shape_SetOnMouseDown,
  Shape_SetOnMouseEnter,
  Shape_SetOnMouseLeave,
  Shape_SetOnMouseMove,
  Shape_SetOnMouseUp,
  Shape_GetAction,
  Shape_SetAction,
  Shape_GetBiDiMode,
  Shape_SetBiDiMode,
  Shape_GetBoundsRect,
  Shape_SetBoundsRect,
  Shape_GetClientHeight,
  Shape_SetClientHeight,
  Shape_GetClientRect,
  Shape_GetClientWidth,
  Shape_SetClientWidth,
  Shape_GetParent,
  Shape_SetParent,
  Shape_GetLeft,
  Shape_SetLeft,
  Shape_GetTop,
  Shape_SetTop,
  Shape_GetWidth,
  Shape_SetWidth,
  Shape_GetHeight,
  Shape_SetHeight,
  Shape_GetCursor,
  Shape_SetCursor,
  Shape_GetHint,
  Shape_SetHint,
  Shape_GetComponentCount,
  Shape_GetComponentIndex,
  Shape_SetComponentIndex,
  Shape_GetOwner,
  Shape_GetName,
  Shape_SetName,
  Shape_GetTag,
  Shape_SetTag,
  Shape_GetComponents;

