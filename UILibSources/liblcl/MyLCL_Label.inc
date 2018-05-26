
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Label_Create(AOwner: TComponent): TLabel; extdecl;
begin
  Result :=  TLabel.Create(AOwner);
end;

procedure Label_Free(AObj: TLabel); extdecl;
begin
  AObj.Free;
end;

procedure Label_BringToFront(AObj: TLabel); extdecl;
begin
  AObj.BringToFront;
end;

function Label_HasParent(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure Label_Hide(AObj: TLabel); extdecl;
begin
  AObj.Hide;
end;

procedure Label_Invalidate(AObj: TLabel); extdecl;
begin
  AObj.Invalidate;
end;

function Label_Perform(AObj: TLabel; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Label_Refresh(AObj: TLabel); extdecl;
begin
  AObj.Refresh;
end;

procedure Label_Repaint(AObj: TLabel); extdecl;
begin
  AObj.Repaint;
end;

procedure Label_SendToBack(AObj: TLabel); extdecl;
begin
  AObj.SendToBack;
end;

procedure Label_SetBounds(AObj: TLabel; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Label_Show(AObj: TLabel); extdecl;
begin
  AObj.Show;
end;

procedure Label_Update(AObj: TLabel); extdecl;
begin
  AObj.Update;
end;

function Label_GetTextBuf(AObj: TLabel; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Label_GetTextLen(AObj: TLabel): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function Label_FindComponent(AObj: TLabel; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Label_GetNamePath(AObj: TLabel): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Label_Assign(AObj: TLabel; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Label_ClassName(AObj: TLabel): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Label_Equals(AObj: TLabel; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Label_GetHashCode(AObj: TLabel): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Label_ToString(AObj: TLabel): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Label_GetAlign(AObj: TLabel): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure Label_SetAlign(AObj: TLabel; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function Label_GetAlignment(AObj: TLabel): TAlignment; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure Label_SetAlignment(AObj: TLabel; AValue: TAlignment); extdecl;
begin
  AObj.Alignment := AValue;
end;

function Label_GetAnchors(AObj: TLabel): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure Label_SetAnchors(AObj: TLabel; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function Label_GetAutoSize(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.AutoSize;
end;

procedure Label_SetAutoSize(AObj: TLabel; AValue: LongBool); extdecl;
begin
  AObj.AutoSize := AValue;
end;

function Label_GetBiDiMode(AObj: TLabel): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Label_SetBiDiMode(AObj: TLabel; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function Label_GetCaption(AObj: TLabel): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure Label_SetCaption(AObj: TLabel; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function Label_GetColor(AObj: TLabel): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure Label_SetColor(AObj: TLabel; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function Label_GetEnabled(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Label_SetEnabled(AObj: TLabel; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Label_GetFont(AObj: TLabel): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure Label_SetFont(AObj: TLabel; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function Label_GetParentColor(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure Label_SetParentColor(AObj: TLabel; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function Label_GetParentFont(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure Label_SetParentFont(AObj: TLabel; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function Label_GetParentShowHint(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Label_SetParentShowHint(AObj: TLabel; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function Label_GetPopupMenu(AObj: TLabel): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Label_SetPopupMenu(AObj: TLabel; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function Label_GetShowAccelChar(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.ShowAccelChar;
end;

procedure Label_SetShowAccelChar(AObj: TLabel; AValue: LongBool); extdecl;
begin
  AObj.ShowAccelChar := AValue;
end;

function Label_GetShowHint(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure Label_SetShowHint(AObj: TLabel; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function Label_GetTransparent(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.Transparent;
end;

procedure Label_SetTransparent(AObj: TLabel; AValue: LongBool); extdecl;
begin
  AObj.Transparent := AValue;
end;

function Label_GetLayout(AObj: TLabel): TTextLayout; extdecl;
begin
  Result :=  AObj.Layout;
end;

procedure Label_SetLayout(AObj: TLabel; AValue: TTextLayout); extdecl;
begin
  AObj.Layout := AValue;
end;

function Label_GetVisible(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure Label_SetVisible(AObj: TLabel; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function Label_GetWordWrap(AObj: TLabel): LongBool; extdecl;
begin
  Result :=  AObj.WordWrap;
end;

procedure Label_SetWordWrap(AObj: TLabel; AValue: LongBool); extdecl;
begin
  AObj.WordWrap := AValue;
end;

procedure Label_SetOnClick(AObj: TLabel; AEventId: NativeUInt); extdecl;
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

procedure Label_SetOnDblClick(AObj: TLabel; AEventId: NativeUInt); extdecl;
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

procedure Label_SetOnMouseDown(AObj: TLabel; AEventId: NativeUInt); extdecl;
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

procedure Label_SetOnMouseMove(AObj: TLabel; AEventId: NativeUInt); extdecl;
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

procedure Label_SetOnMouseUp(AObj: TLabel; AEventId: NativeUInt); extdecl;
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

procedure Label_SetOnMouseEnter(AObj: TLabel; AEventId: NativeUInt); extdecl;
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

procedure Label_SetOnMouseLeave(AObj: TLabel; AEventId: NativeUInt); extdecl;
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

function Label_GetCanvas(AObj: TLabel): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function Label_GetAction(AObj: TLabel): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure Label_SetAction(AObj: TLabel; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure Label_GetBoundsRect(AObj: TLabel; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Label_SetBoundsRect(AObj: TLabel; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function Label_GetClientHeight(AObj: TLabel): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Label_SetClientHeight(AObj: TLabel; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure Label_GetClientRect(AObj: TLabel; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function Label_GetClientWidth(AObj: TLabel): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Label_SetClientWidth(AObj: TLabel; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function Label_GetParent(AObj: TLabel): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure Label_SetParent(AObj: TLabel; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function Label_GetLeft(AObj: TLabel): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure Label_SetLeft(AObj: TLabel; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function Label_GetTop(AObj: TLabel): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure Label_SetTop(AObj: TLabel; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function Label_GetWidth(AObj: TLabel): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Label_SetWidth(AObj: TLabel; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function Label_GetHeight(AObj: TLabel): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Label_SetHeight(AObj: TLabel; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Label_GetCursor(AObj: TLabel): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure Label_SetCursor(AObj: TLabel; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function Label_GetHint(AObj: TLabel): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Label_SetHint(AObj: TLabel; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function Label_GetComponentCount(AObj: TLabel): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Label_GetComponentIndex(AObj: TLabel): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Label_SetComponentIndex(AObj: TLabel; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Label_GetOwner(AObj: TLabel): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Label_GetName(AObj: TLabel): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Label_SetName(AObj: TLabel; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Label_GetTag(AObj: TLabel): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Label_SetTag(AObj: TLabel; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Label_GetComponents(AObj: TLabel; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Label_Create,
  Label_Free,
  Label_BringToFront,
  Label_HasParent,
  Label_Hide,
  Label_Invalidate,
  Label_Perform,
  Label_Refresh,
  Label_Repaint,
  Label_SendToBack,
  Label_SetBounds,
  Label_Show,
  Label_Update,
  Label_GetTextBuf,
  Label_GetTextLen,
  Label_FindComponent,
  Label_GetNamePath,
  Label_Assign,
  Label_ClassName,
  Label_Equals,
  Label_GetHashCode,
  Label_ToString,
  Label_GetAlign,
  Label_SetAlign,
  Label_GetAlignment,
  Label_SetAlignment,
  Label_GetAnchors,
  Label_SetAnchors,
  Label_GetAutoSize,
  Label_SetAutoSize,
  Label_GetBiDiMode,
  Label_SetBiDiMode,
  Label_GetCaption,
  Label_SetCaption,
  Label_GetColor,
  Label_SetColor,
  Label_GetEnabled,
  Label_SetEnabled,
  Label_GetFont,
  Label_SetFont,
  Label_GetParentColor,
  Label_SetParentColor,
  Label_GetParentFont,
  Label_SetParentFont,
  Label_GetParentShowHint,
  Label_SetParentShowHint,
  Label_GetPopupMenu,
  Label_SetPopupMenu,
  Label_GetShowAccelChar,
  Label_SetShowAccelChar,
  Label_GetShowHint,
  Label_SetShowHint,
  Label_GetTransparent,
  Label_SetTransparent,
  Label_GetLayout,
  Label_SetLayout,
  Label_GetVisible,
  Label_SetVisible,
  Label_GetWordWrap,
  Label_SetWordWrap,
  Label_SetOnClick,
  Label_SetOnDblClick,
  Label_SetOnMouseDown,
  Label_SetOnMouseMove,
  Label_SetOnMouseUp,
  Label_SetOnMouseEnter,
  Label_SetOnMouseLeave,
  Label_GetCanvas,
  Label_GetAction,
  Label_SetAction,
  Label_GetBoundsRect,
  Label_SetBoundsRect,
  Label_GetClientHeight,
  Label_SetClientHeight,
  Label_GetClientRect,
  Label_GetClientWidth,
  Label_SetClientWidth,
  Label_GetParent,
  Label_SetParent,
  Label_GetLeft,
  Label_SetLeft,
  Label_GetTop,
  Label_SetTop,
  Label_GetWidth,
  Label_SetWidth,
  Label_GetHeight,
  Label_SetHeight,
  Label_GetCursor,
  Label_SetCursor,
  Label_GetHint,
  Label_SetHint,
  Label_GetComponentCount,
  Label_GetComponentIndex,
  Label_SetComponentIndex,
  Label_GetOwner,
  Label_GetName,
  Label_SetName,
  Label_GetTag,
  Label_SetTag,
  Label_GetComponents;

