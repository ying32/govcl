
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ImageButton_Create(AOwner: TComponent): TImageButton; stdcall;
begin
  Result :=  TImageButton.Create(AOwner);
end;

procedure ImageButton_Free(AObj: TImageButton); stdcall;
begin
  AObj.Free;
end;

procedure ImageButton_Click(AObj: TImageButton); stdcall;
begin
  AObj.Click;
end;

procedure ImageButton_BringToFront(AObj: TImageButton); stdcall;
begin
  AObj.BringToFront;
end;

function ImageButton_HasParent(AObj: TImageButton): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ImageButton_Hide(AObj: TImageButton); stdcall;
begin
  AObj.Hide;
end;

procedure ImageButton_Invalidate(AObj: TImageButton); stdcall;
begin
  AObj.Invalidate;
end;

function ImageButton_Perform(AObj: TImageButton; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ImageButton_Refresh(AObj: TImageButton); stdcall;
begin
  AObj.Refresh;
end;

procedure ImageButton_Repaint(AObj: TImageButton); stdcall;
begin
  AObj.Repaint;
end;

procedure ImageButton_SendToBack(AObj: TImageButton); stdcall;
begin
  AObj.SendToBack;
end;

procedure ImageButton_SetBounds(AObj: TImageButton; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ImageButton_Show(AObj: TImageButton); stdcall;
begin
  AObj.Show;
end;

procedure ImageButton_Update(AObj: TImageButton); stdcall;
begin
  AObj.Update;
end;

function ImageButton_GetTextBuf(AObj: TImageButton; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ImageButton_GetTextLen(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function ImageButton_FindComponent(AObj: TImageButton; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ImageButton_GetNamePath(AObj: TImageButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ImageButton_Assign(AObj: TImageButton; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ImageButton_ClassName(AObj: TImageButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ImageButton_Equals(AObj: TImageButton; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ImageButton_GetHashCode(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ImageButton_ToString(AObj: TImageButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ImageButton_GetAction(AObj: TImageButton): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ImageButton_SetAction(AObj: TImageButton; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function ImageButton_GetAlign(AObj: TImageButton): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ImageButton_SetAlign(AObj: TImageButton; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ImageButton_GetAnchors(AObj: TImageButton): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ImageButton_SetAnchors(AObj: TImageButton; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ImageButton_GetAutoSize(AObj: TImageButton): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure ImageButton_SetAutoSize(AObj: TImageButton; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function ImageButton_GetCaption(AObj: TImageButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure ImageButton_SetCaption(AObj: TImageButton; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function ImageButton_GetEnabled(AObj: TImageButton): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ImageButton_SetEnabled(AObj: TImageButton; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ImageButton_GetFont(AObj: TImageButton): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure ImageButton_SetFont(AObj: TImageButton; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function ImageButton_GetImageCount(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.ImageCount;
end;

procedure ImageButton_SetImageCount(AObj: TImageButton; AValue: Integer); stdcall;
begin
  AObj.ImageCount := AValue;
end;

function ImageButton_GetModalResult(AObj: TImageButton): TModalResult; stdcall;
begin
  Result :=  AObj.ModalResult;
end;

procedure ImageButton_SetModalResult(AObj: TImageButton; AValue: TModalResult); stdcall;
begin
  AObj.ModalResult := AValue;
end;

function ImageButton_GetParentShowHint(AObj: TImageButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ImageButton_SetParentShowHint(AObj: TImageButton; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ImageButton_GetParentFont(AObj: TImageButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure ImageButton_SetParentFont(AObj: TImageButton; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function ImageButton_GetPicture(AObj: TImageButton): TPicture; stdcall;
begin
  Result :=  AObj.Picture;
end;

procedure ImageButton_SetPicture(AObj: TImageButton; AValue: TPicture); stdcall;
begin
  AObj.Picture := AValue;
end;

function ImageButton_GetPopupMenu(AObj: TImageButton): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ImageButton_SetPopupMenu(AObj: TImageButton; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ImageButton_GetShowHint(AObj: TImageButton): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ImageButton_SetShowHint(AObj: TImageButton; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ImageButton_GetShowCaption(AObj: TImageButton): LongBool; stdcall;
begin
  Result :=  AObj.ShowCaption;
end;

procedure ImageButton_SetShowCaption(AObj: TImageButton; AValue: LongBool); stdcall;
begin
  AObj.ShowCaption := AValue;
end;

function ImageButton_GetVisible(AObj: TImageButton): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ImageButton_SetVisible(AObj: TImageButton; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

procedure ImageButton_SetOnClick(AObj: TImageButton; AEventId: NativeUInt); stdcall;
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

procedure ImageButton_SetOnDblClick(AObj: TImageButton; AEventId: NativeUInt); stdcall;
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

procedure ImageButton_SetOnMouseDown(AObj: TImageButton; AEventId: NativeUInt); stdcall;
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

procedure ImageButton_SetOnMouseEnter(AObj: TImageButton; AEventId: NativeUInt); stdcall;
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

procedure ImageButton_SetOnMouseLeave(AObj: TImageButton; AEventId: NativeUInt); stdcall;
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

procedure ImageButton_SetOnMouseMove(AObj: TImageButton; AEventId: NativeUInt); stdcall;
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

procedure ImageButton_SetOnMouseUp(AObj: TImageButton; AEventId: NativeUInt); stdcall;
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

function ImageButton_GetBiDiMode(AObj: TImageButton): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ImageButton_SetBiDiMode(AObj: TImageButton; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure ImageButton_GetBoundsRect(AObj: TImageButton; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ImageButton_SetBoundsRect(AObj: TImageButton; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ImageButton_GetClientHeight(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ImageButton_SetClientHeight(AObj: TImageButton; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ImageButton_GetClientRect(AObj: TImageButton; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ImageButton_GetClientWidth(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ImageButton_SetClientWidth(AObj: TImageButton; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ImageButton_GetExplicitLeft(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ImageButton_GetExplicitTop(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ImageButton_GetExplicitWidth(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ImageButton_GetExplicitHeight(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ImageButton_GetParent(AObj: TImageButton): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ImageButton_SetParent(AObj: TImageButton; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ImageButton_GetStyleElements(AObj: TImageButton): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ImageButton_SetStyleElements(AObj: TImageButton; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function ImageButton_GetAlignWithMargins(AObj: TImageButton): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ImageButton_SetAlignWithMargins(AObj: TImageButton; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ImageButton_GetLeft(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ImageButton_SetLeft(AObj: TImageButton; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ImageButton_GetTop(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ImageButton_SetTop(AObj: TImageButton; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ImageButton_GetWidth(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ImageButton_SetWidth(AObj: TImageButton; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function ImageButton_GetHeight(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ImageButton_SetHeight(AObj: TImageButton; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ImageButton_GetCursor(AObj: TImageButton): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ImageButton_SetCursor(AObj: TImageButton; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ImageButton_GetHint(AObj: TImageButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ImageButton_SetHint(AObj: TImageButton; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ImageButton_GetMargins(AObj: TImageButton): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ImageButton_SetMargins(AObj: TImageButton; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ImageButton_GetCustomHint(AObj: TImageButton): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ImageButton_SetCustomHint(AObj: TImageButton; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ImageButton_GetComponentCount(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ImageButton_GetComponentIndex(AObj: TImageButton): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ImageButton_SetComponentIndex(AObj: TImageButton; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ImageButton_GetOwner(AObj: TImageButton): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ImageButton_GetName(AObj: TImageButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ImageButton_SetName(AObj: TImageButton; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ImageButton_GetTag(AObj: TImageButton): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ImageButton_SetTag(AObj: TImageButton; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ImageButton_GetComponents(AObj: TImageButton; AIndex: Integer): TComponent; stdcall;
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
  ImageButton_GetExplicitLeft,
  ImageButton_GetExplicitTop,
  ImageButton_GetExplicitWidth,
  ImageButton_GetExplicitHeight,
  ImageButton_GetParent,
  ImageButton_SetParent,
  ImageButton_GetStyleElements,
  ImageButton_SetStyleElements,
  ImageButton_GetAlignWithMargins,
  ImageButton_SetAlignWithMargins,
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
  ImageButton_GetMargins,
  ImageButton_SetMargins,
  ImageButton_GetCustomHint,
  ImageButton_SetCustomHint,
  ImageButton_GetComponentCount,
  ImageButton_GetComponentIndex,
  ImageButton_SetComponentIndex,
  ImageButton_GetOwner,
  ImageButton_GetName,
  ImageButton_SetName,
  ImageButton_GetTag,
  ImageButton_SetTag,
  ImageButton_GetComponents;

