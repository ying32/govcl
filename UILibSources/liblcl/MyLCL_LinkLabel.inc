
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function LinkLabel_Create(AOwner: TComponent): TLinkLabel; extdecl;
begin
  Result :=  TLinkLabel.Create(AOwner);
end;

procedure LinkLabel_Free(AObj: TLinkLabel); extdecl;
begin
  AObj.Free;
end;

procedure LinkLabel_Invalidate(AObj: TLinkLabel); extdecl;
begin
  AObj.Invalidate;
end;

procedure LinkLabel_Repaint(AObj: TLinkLabel); extdecl;
begin
  AObj.Repaint;
end;

procedure LinkLabel_SetBounds(AObj: TLinkLabel; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure LinkLabel_Update(AObj: TLinkLabel); extdecl;
begin
  AObj.Update;
end;

procedure LinkLabel_BringToFront(AObj: TLinkLabel); extdecl;
begin
  AObj.BringToFront;
end;

function LinkLabel_HasParent(AObj: TLinkLabel): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure LinkLabel_Hide(AObj: TLinkLabel); extdecl;
begin
  AObj.Hide;
end;

function LinkLabel_Perform(AObj: TLinkLabel; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure LinkLabel_Refresh(AObj: TLinkLabel); extdecl;
begin
  AObj.Refresh;
end;

procedure LinkLabel_SendToBack(AObj: TLinkLabel); extdecl;
begin
  AObj.SendToBack;
end;

procedure LinkLabel_Show(AObj: TLinkLabel); extdecl;
begin
  AObj.Show;
end;

function LinkLabel_GetTextBuf(AObj: TLinkLabel; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function LinkLabel_GetTextLen(AObj: TLinkLabel): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function LinkLabel_FindComponent(AObj: TLinkLabel; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function LinkLabel_GetNamePath(AObj: TLinkLabel): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure LinkLabel_Assign(AObj: TLinkLabel; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function LinkLabel_ClassName(AObj: TLinkLabel): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function LinkLabel_Equals(AObj: TLinkLabel; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function LinkLabel_GetHashCode(AObj: TLinkLabel): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function LinkLabel_ToString(AObj: TLinkLabel): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function LinkLabel_GetAlign(AObj: TLinkLabel): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure LinkLabel_SetAlign(AObj: TLinkLabel; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function LinkLabel_GetAlignment(AObj: TLinkLabel): TAlignment; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure LinkLabel_SetAlignment(AObj: TLinkLabel; AValue: TAlignment); extdecl;
begin
  AObj.Alignment := AValue;
end;

function LinkLabel_GetAnchors(AObj: TLinkLabel): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure LinkLabel_SetAnchors(AObj: TLinkLabel; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function LinkLabel_GetAutoSize(AObj: TLinkLabel): LongBool; extdecl;
begin
  Result :=  AObj.AutoSize;
end;

procedure LinkLabel_SetAutoSize(AObj: TLinkLabel; AValue: LongBool); extdecl;
begin
  AObj.AutoSize := AValue;
end;

function LinkLabel_GetCaption(AObj: TLinkLabel): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure LinkLabel_SetCaption(AObj: TLinkLabel; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function LinkLabel_GetColor(AObj: TLinkLabel): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure LinkLabel_SetColor(AObj: TLinkLabel; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function LinkLabel_GetEnabled(AObj: TLinkLabel): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure LinkLabel_SetEnabled(AObj: TLinkLabel; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function LinkLabel_GetFont(AObj: TLinkLabel): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure LinkLabel_SetFont(AObj: TLinkLabel; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function LinkLabel_GetParentColor(AObj: TLinkLabel): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure LinkLabel_SetParentColor(AObj: TLinkLabel; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function LinkLabel_GetParentFont(AObj: TLinkLabel): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure LinkLabel_SetParentFont(AObj: TLinkLabel; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function LinkLabel_GetParentShowHint(AObj: TLinkLabel): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure LinkLabel_SetParentShowHint(AObj: TLinkLabel; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function LinkLabel_GetPopupMenu(AObj: TLinkLabel): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure LinkLabel_SetPopupMenu(AObj: TLinkLabel; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function LinkLabel_GetShowHint(AObj: TLinkLabel): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure LinkLabel_SetShowHint(AObj: TLinkLabel; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function LinkLabel_GetVisible(AObj: TLinkLabel): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure LinkLabel_SetVisible(AObj: TLinkLabel; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure LinkLabel_SetOnClick(AObj: TLinkLabel; AEventId: NativeUInt); extdecl;
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

procedure LinkLabel_SetOnDblClick(AObj: TLinkLabel; AEventId: NativeUInt); extdecl;
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

procedure LinkLabel_SetOnMouseDown(AObj: TLinkLabel; AEventId: NativeUInt); extdecl;
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

procedure LinkLabel_SetOnMouseEnter(AObj: TLinkLabel; AEventId: NativeUInt); extdecl;
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

procedure LinkLabel_SetOnMouseLeave(AObj: TLinkLabel; AEventId: NativeUInt); extdecl;
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

procedure LinkLabel_SetOnMouseMove(AObj: TLinkLabel; AEventId: NativeUInt); extdecl;
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

procedure LinkLabel_SetOnMouseUp(AObj: TLinkLabel; AEventId: NativeUInt); extdecl;
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

procedure LinkLabel_SetOnLinkClick(AObj: TLinkLabel; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnLinkClick := nil;
    EventClass.Remove(AObj, geLinkClick);
    Exit;
  end;
  AObj.OnLinkClick := @EventClass.OnLinkClick;
  EventClass.Add(AObj, geLinkClick, AEventId);
end;

function LinkLabel_GetAction(AObj: TLinkLabel): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure LinkLabel_SetAction(AObj: TLinkLabel; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function LinkLabel_GetBiDiMode(AObj: TLinkLabel): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure LinkLabel_SetBiDiMode(AObj: TLinkLabel; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure LinkLabel_GetBoundsRect(AObj: TLinkLabel; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure LinkLabel_SetBoundsRect(AObj: TLinkLabel; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function LinkLabel_GetClientHeight(AObj: TLinkLabel): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure LinkLabel_SetClientHeight(AObj: TLinkLabel; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure LinkLabel_GetClientRect(AObj: TLinkLabel; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function LinkLabel_GetClientWidth(AObj: TLinkLabel): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure LinkLabel_SetClientWidth(AObj: TLinkLabel; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function LinkLabel_GetParent(AObj: TLinkLabel): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure LinkLabel_SetParent(AObj: TLinkLabel; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function LinkLabel_GetLeft(AObj: TLinkLabel): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure LinkLabel_SetLeft(AObj: TLinkLabel; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function LinkLabel_GetTop(AObj: TLinkLabel): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure LinkLabel_SetTop(AObj: TLinkLabel; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function LinkLabel_GetWidth(AObj: TLinkLabel): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure LinkLabel_SetWidth(AObj: TLinkLabel; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function LinkLabel_GetHeight(AObj: TLinkLabel): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure LinkLabel_SetHeight(AObj: TLinkLabel; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function LinkLabel_GetHint(AObj: TLinkLabel): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure LinkLabel_SetHint(AObj: TLinkLabel; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function LinkLabel_GetComponentCount(AObj: TLinkLabel): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function LinkLabel_GetComponentIndex(AObj: TLinkLabel): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure LinkLabel_SetComponentIndex(AObj: TLinkLabel; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function LinkLabel_GetOwner(AObj: TLinkLabel): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function LinkLabel_GetName(AObj: TLinkLabel): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure LinkLabel_SetName(AObj: TLinkLabel; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function LinkLabel_GetTag(AObj: TLinkLabel): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure LinkLabel_SetTag(AObj: TLinkLabel; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function LinkLabel_GetComponents(AObj: TLinkLabel; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  LinkLabel_Create,
  LinkLabel_Free,
  LinkLabel_Invalidate,
  LinkLabel_Repaint,
  LinkLabel_SetBounds,
  LinkLabel_Update,
  LinkLabel_BringToFront,
  LinkLabel_HasParent,
  LinkLabel_Hide,
  LinkLabel_Perform,
  LinkLabel_Refresh,
  LinkLabel_SendToBack,
  LinkLabel_Show,
  LinkLabel_GetTextBuf,
  LinkLabel_GetTextLen,
  LinkLabel_FindComponent,
  LinkLabel_GetNamePath,
  LinkLabel_Assign,
  LinkLabel_ClassName,
  LinkLabel_Equals,
  LinkLabel_GetHashCode,
  LinkLabel_ToString,
  LinkLabel_GetAlign,
  LinkLabel_SetAlign,
  LinkLabel_GetAlignment,
  LinkLabel_SetAlignment,
  LinkLabel_GetAnchors,
  LinkLabel_SetAnchors,
  LinkLabel_GetAutoSize,
  LinkLabel_SetAutoSize,
  LinkLabel_GetCaption,
  LinkLabel_SetCaption,
  LinkLabel_GetColor,
  LinkLabel_SetColor,
  LinkLabel_GetEnabled,
  LinkLabel_SetEnabled,
  LinkLabel_GetFont,
  LinkLabel_SetFont,
  LinkLabel_GetParentColor,
  LinkLabel_SetParentColor,
  LinkLabel_GetParentFont,
  LinkLabel_SetParentFont,
  LinkLabel_GetParentShowHint,
  LinkLabel_SetParentShowHint,
  LinkLabel_GetPopupMenu,
  LinkLabel_SetPopupMenu,
  LinkLabel_GetShowHint,
  LinkLabel_SetShowHint,
  LinkLabel_GetVisible,
  LinkLabel_SetVisible,
  LinkLabel_SetOnClick,
  LinkLabel_SetOnDblClick,
  LinkLabel_SetOnMouseDown,
  LinkLabel_SetOnMouseEnter,
  LinkLabel_SetOnMouseLeave,
  LinkLabel_SetOnMouseMove,
  LinkLabel_SetOnMouseUp,
  LinkLabel_SetOnLinkClick,
  LinkLabel_GetAction,
  LinkLabel_SetAction,
  LinkLabel_GetBiDiMode,
  LinkLabel_SetBiDiMode,
  LinkLabel_GetBoundsRect,
  LinkLabel_SetBoundsRect,
  LinkLabel_GetClientHeight,
  LinkLabel_SetClientHeight,
  LinkLabel_GetClientRect,
  LinkLabel_GetClientWidth,
  LinkLabel_SetClientWidth,
  LinkLabel_GetParent,
  LinkLabel_SetParent,
  LinkLabel_GetLeft,
  LinkLabel_SetLeft,
  LinkLabel_GetTop,
  LinkLabel_SetTop,
  LinkLabel_GetWidth,
  LinkLabel_SetWidth,
  LinkLabel_GetHeight,
  LinkLabel_SetHeight,
  LinkLabel_GetHint,
  LinkLabel_SetHint,
  LinkLabel_GetComponentCount,
  LinkLabel_GetComponentIndex,
  LinkLabel_SetComponentIndex,
  LinkLabel_GetOwner,
  LinkLabel_GetName,
  LinkLabel_SetName,
  LinkLabel_GetTag,
  LinkLabel_SetTag,
  LinkLabel_GetComponents;

