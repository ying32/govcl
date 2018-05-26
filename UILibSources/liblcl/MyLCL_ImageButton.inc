
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ImageButton_Create(AOwner: TComponent): TImageButton; extdecl;
begin
  Result :=  TImageButton.Create(AOwner);
end;

procedure ImageButton_Free(AObj: TImageButton); extdecl;
begin
  AObj.Free;
end;

procedure ImageButton_Click(AObj: TImageButton); extdecl;
begin
  AObj.Click;
end;

procedure ImageButton_BringToFront(AObj: TImageButton); extdecl;
begin
  AObj.BringToFront;
end;

function ImageButton_HasParent(AObj: TImageButton): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ImageButton_Hide(AObj: TImageButton); extdecl;
begin
  AObj.Hide;
end;

procedure ImageButton_Invalidate(AObj: TImageButton); extdecl;
begin
  AObj.Invalidate;
end;

function ImageButton_Perform(AObj: TImageButton; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ImageButton_Refresh(AObj: TImageButton); extdecl;
begin
  AObj.Refresh;
end;

procedure ImageButton_Repaint(AObj: TImageButton); extdecl;
begin
  AObj.Repaint;
end;

procedure ImageButton_SendToBack(AObj: TImageButton); extdecl;
begin
  AObj.SendToBack;
end;

procedure ImageButton_SetBounds(AObj: TImageButton; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ImageButton_Show(AObj: TImageButton); extdecl;
begin
  AObj.Show;
end;

procedure ImageButton_Update(AObj: TImageButton); extdecl;
begin
  AObj.Update;
end;

function ImageButton_GetTextBuf(AObj: TImageButton; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ImageButton_GetTextLen(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function ImageButton_FindComponent(AObj: TImageButton; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ImageButton_GetNamePath(AObj: TImageButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ImageButton_Assign(AObj: TImageButton; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ImageButton_ClassName(AObj: TImageButton): PChar; extdecl;
begin
  Result :=  PChar(string(AObj.ClassName));
end;

function ImageButton_Equals(AObj: TImageButton; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ImageButton_GetHashCode(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ImageButton_ToString(AObj: TImageButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ImageButton_GetAction(AObj: TImageButton): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure ImageButton_SetAction(AObj: TImageButton; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function ImageButton_GetAlign(AObj: TImageButton): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure ImageButton_SetAlign(AObj: TImageButton; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function ImageButton_GetAnchors(AObj: TImageButton): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure ImageButton_SetAnchors(AObj: TImageButton; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function ImageButton_GetAutoSize(AObj: TImageButton): LongBool; extdecl;
begin
  Result :=  AObj.AutoSize;
end;

procedure ImageButton_SetAutoSize(AObj: TImageButton; AValue: LongBool); extdecl;
begin
  AObj.AutoSize := AValue;
end;

function ImageButton_GetCaption(AObj: TImageButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure ImageButton_SetCaption(AObj: TImageButton; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function ImageButton_GetEnabled(AObj: TImageButton): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure ImageButton_SetEnabled(AObj: TImageButton; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function ImageButton_GetFont(AObj: TImageButton): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure ImageButton_SetFont(AObj: TImageButton; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function ImageButton_GetImageCount(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.ImageCount;
end;

procedure ImageButton_SetImageCount(AObj: TImageButton; AValue: Integer); extdecl;
begin
  AObj.ImageCount := AValue;
end;

function ImageButton_GetModalResult(AObj: TImageButton): TModalResult; extdecl;
begin
  Result :=  AObj.ModalResult;
end;

procedure ImageButton_SetModalResult(AObj: TImageButton; AValue: TModalResult); extdecl;
begin
  AObj.ModalResult := AValue;
end;

function ImageButton_GetParentShowHint(AObj: TImageButton): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ImageButton_SetParentShowHint(AObj: TImageButton; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function ImageButton_GetParentFont(AObj: TImageButton): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure ImageButton_SetParentFont(AObj: TImageButton; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function ImageButton_GetPicture(AObj: TImageButton): TPicture; extdecl;
begin
  Result :=  AObj.Picture;
end;

procedure ImageButton_SetPicture(AObj: TImageButton; AValue: TPicture); extdecl;
begin
  AObj.Picture := AValue;
end;

function ImageButton_GetPopupMenu(AObj: TImageButton): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ImageButton_SetPopupMenu(AObj: TImageButton; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function ImageButton_GetShowHint(AObj: TImageButton): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure ImageButton_SetShowHint(AObj: TImageButton; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function ImageButton_GetShowCaption(AObj: TImageButton): LongBool; extdecl;
begin
  Result :=  AObj.ShowCaption;
end;

procedure ImageButton_SetShowCaption(AObj: TImageButton; AValue: LongBool); extdecl;
begin
  AObj.ShowCaption := AValue;
end;

function ImageButton_GetVisible(AObj: TImageButton): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure ImageButton_SetVisible(AObj: TImageButton; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure ImageButton_SetOnClick(AObj: TImageButton; AEventId: NativeUInt); extdecl;
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

procedure ImageButton_SetOnDblClick(AObj: TImageButton; AEventId: NativeUInt); extdecl;
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

procedure ImageButton_SetOnMouseDown(AObj: TImageButton; AEventId: NativeUInt); extdecl;
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

procedure ImageButton_SetOnMouseEnter(AObj: TImageButton; AEventId: NativeUInt); extdecl;
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

procedure ImageButton_SetOnMouseLeave(AObj: TImageButton; AEventId: NativeUInt); extdecl;
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

procedure ImageButton_SetOnMouseMove(AObj: TImageButton; AEventId: NativeUInt); extdecl;
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

procedure ImageButton_SetOnMouseUp(AObj: TImageButton; AEventId: NativeUInt); extdecl;
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

function ImageButton_GetBiDiMode(AObj: TImageButton): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ImageButton_SetBiDiMode(AObj: TImageButton; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure ImageButton_GetBoundsRect(AObj: TImageButton; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ImageButton_SetBoundsRect(AObj: TImageButton; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function ImageButton_GetClientHeight(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ImageButton_SetClientHeight(AObj: TImageButton; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure ImageButton_GetClientRect(AObj: TImageButton; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function ImageButton_GetClientWidth(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ImageButton_SetClientWidth(AObj: TImageButton; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function ImageButton_GetParent(AObj: TImageButton): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure ImageButton_SetParent(AObj: TImageButton; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function ImageButton_GetLeft(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ImageButton_SetLeft(AObj: TImageButton; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ImageButton_GetTop(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ImageButton_SetTop(AObj: TImageButton; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ImageButton_GetWidth(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ImageButton_SetWidth(AObj: TImageButton; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function ImageButton_GetHeight(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ImageButton_SetHeight(AObj: TImageButton; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ImageButton_GetCursor(AObj: TImageButton): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure ImageButton_SetCursor(AObj: TImageButton; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function ImageButton_GetHint(AObj: TImageButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ImageButton_SetHint(AObj: TImageButton; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function ImageButton_GetComponentCount(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ImageButton_GetComponentIndex(AObj: TImageButton): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ImageButton_SetComponentIndex(AObj: TImageButton; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ImageButton_GetOwner(AObj: TImageButton): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ImageButton_GetName(AObj: TImageButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ImageButton_SetName(AObj: TImageButton; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ImageButton_GetTag(AObj: TImageButton): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ImageButton_SetTag(AObj: TImageButton; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ImageButton_GetComponents(AObj: TImageButton; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ImageButton_Create,
  ImageButton_Free,
  ImageButton_Click,
  ImageButton_BringToFront,
  ImageButton_HasParent,
  ImageButton_Hide,
  ImageButton_Invalidate,
  ImageButton_Perform,
  ImageButton_Refresh,
  ImageButton_Repaint,
  ImageButton_SendToBack,
  ImageButton_SetBounds,
  ImageButton_Show,
  ImageButton_Update,
  ImageButton_GetTextBuf,
  ImageButton_GetTextLen,
  ImageButton_FindComponent,
  ImageButton_GetNamePath,
  ImageButton_Assign,
  ImageButton_ClassName,
  ImageButton_Equals,
  ImageButton_GetHashCode,
  ImageButton_ToString,
  ImageButton_GetAction,
  ImageButton_SetAction,
  ImageButton_GetAlign,
  ImageButton_SetAlign,
  ImageButton_GetAnchors,
  ImageButton_SetAnchors,
  ImageButton_GetAutoSize,
  ImageButton_SetAutoSize,
  ImageButton_GetCaption,
  ImageButton_SetCaption,
  ImageButton_GetEnabled,
  ImageButton_SetEnabled,
  ImageButton_GetFont,
  ImageButton_SetFont,
  ImageButton_GetImageCount,
  ImageButton_SetImageCount,
  ImageButton_GetModalResult,
  ImageButton_SetModalResult,
  ImageButton_GetParentShowHint,
  ImageButton_SetParentShowHint,
  ImageButton_GetParentFont,
  ImageButton_SetParentFont,
  ImageButton_GetPicture,
  ImageButton_SetPicture,
  ImageButton_GetPopupMenu,
  ImageButton_SetPopupMenu,
  ImageButton_GetShowHint,
  ImageButton_SetShowHint,
  ImageButton_GetShowCaption,
  ImageButton_SetShowCaption,
  ImageButton_GetVisible,
  ImageButton_SetVisible,
  ImageButton_SetOnClick,
  ImageButton_SetOnDblClick,
  ImageButton_SetOnMouseDown,
  ImageButton_SetOnMouseEnter,
  ImageButton_SetOnMouseLeave,
  ImageButton_SetOnMouseMove,
  ImageButton_SetOnMouseUp,
  ImageButton_GetBiDiMode,
  ImageButton_SetBiDiMode,
  ImageButton_GetBoundsRect,
  ImageButton_SetBoundsRect,
  ImageButton_GetClientHeight,
  ImageButton_SetClientHeight,
  ImageButton_GetClientRect,
  ImageButton_GetClientWidth,
  ImageButton_SetClientWidth,
  ImageButton_GetParent,
  ImageButton_SetParent,
  ImageButton_GetLeft,
  ImageButton_SetLeft,
  ImageButton_GetTop,
  ImageButton_SetTop,
  ImageButton_GetWidth,
  ImageButton_SetWidth,
  ImageButton_GetHeight,
  ImageButton_SetHeight,
  ImageButton_GetCursor,
  ImageButton_SetCursor,
  ImageButton_GetHint,
  ImageButton_SetHint,
  ImageButton_GetComponentCount,
  ImageButton_GetComponentIndex,
  ImageButton_SetComponentIndex,
  ImageButton_GetOwner,
  ImageButton_GetName,
  ImageButton_SetName,
  ImageButton_GetTag,
  ImageButton_SetTag,
  ImageButton_GetComponents;

