
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Label_Create(AOwner: TComponent): TLabel; stdcall;
begin
  Result :=  TLabel.Create(AOwner);
end;

procedure Label_Free(AObj: TLabel); stdcall;
begin
  AObj.Free;
end;

procedure Label_BringToFront(AObj: TLabel); stdcall;
begin
  AObj.BringToFront;
end;

function Label_HasParent(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure Label_Hide(AObj: TLabel); stdcall;
begin
  AObj.Hide;
end;

procedure Label_Invalidate(AObj: TLabel); stdcall;
begin
  AObj.Invalidate;
end;

function Label_Perform(AObj: TLabel; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Label_Refresh(AObj: TLabel); stdcall;
begin
  AObj.Refresh;
end;

procedure Label_Repaint(AObj: TLabel); stdcall;
begin
  AObj.Repaint;
end;

procedure Label_SendToBack(AObj: TLabel); stdcall;
begin
  AObj.SendToBack;
end;

procedure Label_SetBounds(AObj: TLabel; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Label_Show(AObj: TLabel); stdcall;
begin
  AObj.Show;
end;

procedure Label_Update(AObj: TLabel); stdcall;
begin
  AObj.Update;
end;

function Label_GetTextBuf(AObj: TLabel; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Label_GetTextLen(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function Label_FindComponent(AObj: TLabel; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Label_GetNamePath(AObj: TLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Label_Assign(AObj: TLabel; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Label_ClassName(AObj: TLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Label_Equals(AObj: TLabel; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Label_GetHashCode(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Label_ToString(AObj: TLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Label_GetAlign(AObj: TLabel): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Label_SetAlign(AObj: TLabel; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Label_GetAlignment(AObj: TLabel): TAlignment; stdcall;
begin
  Result :=  AObj.Alignment;
end;

procedure Label_SetAlignment(AObj: TLabel; AValue: TAlignment); stdcall;
begin
  AObj.Alignment := AValue;
end;

function Label_GetAnchors(AObj: TLabel): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Label_SetAnchors(AObj: TLabel; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Label_GetAutoSize(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure Label_SetAutoSize(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function Label_GetBiDiMode(AObj: TLabel): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Label_SetBiDiMode(AObj: TLabel; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function Label_GetCaption(AObj: TLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure Label_SetCaption(AObj: TLabel; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function Label_GetColor(AObj: TLabel): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure Label_SetColor(AObj: TLabel; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function Label_GetEllipsisPosition(AObj: TLabel): TEllipsisPosition; stdcall;
begin
  Result :=  AObj.EllipsisPosition;
end;

procedure Label_SetEllipsisPosition(AObj: TLabel; AValue: TEllipsisPosition); stdcall;
begin
  AObj.EllipsisPosition := AValue;
end;

function Label_GetEnabled(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Label_SetEnabled(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Label_GetFont(AObj: TLabel): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure Label_SetFont(AObj: TLabel; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function Label_GetGlowSize(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.GlowSize;
end;

procedure Label_SetGlowSize(AObj: TLabel; AValue: Integer); stdcall;
begin
  AObj.GlowSize := AValue;
end;

function Label_GetParentColor(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.ParentColor;
end;

procedure Label_SetParentColor(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.ParentColor := AValue;
end;

function Label_GetParentFont(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure Label_SetParentFont(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function Label_GetParentShowHint(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure Label_SetParentShowHint(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function Label_GetPopupMenu(AObj: TLabel): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Label_SetPopupMenu(AObj: TLabel; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function Label_GetShowAccelChar(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.ShowAccelChar;
end;

procedure Label_SetShowAccelChar(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.ShowAccelChar := AValue;
end;

function Label_GetShowHint(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Label_SetShowHint(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Label_GetTransparent(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.Transparent;
end;

procedure Label_SetTransparent(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.Transparent := AValue;
end;

function Label_GetLayout(AObj: TLabel): TTextLayout; stdcall;
begin
  Result :=  AObj.Layout;
end;

procedure Label_SetLayout(AObj: TLabel; AValue: TTextLayout); stdcall;
begin
  AObj.Layout := AValue;
end;

function Label_GetVisible(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Label_SetVisible(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Label_GetWordWrap(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.WordWrap;
end;

procedure Label_SetWordWrap(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.WordWrap := AValue;
end;

function Label_GetStyleElements(AObj: TLabel): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Label_SetStyleElements(AObj: TLabel; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure Label_SetOnClick(AObj: TLabel; AEventId: NativeUInt); stdcall;
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

procedure Label_SetOnDblClick(AObj: TLabel; AEventId: NativeUInt); stdcall;
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

procedure Label_SetOnMouseDown(AObj: TLabel; AEventId: NativeUInt); stdcall;
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

procedure Label_SetOnMouseMove(AObj: TLabel; AEventId: NativeUInt); stdcall;
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

procedure Label_SetOnMouseUp(AObj: TLabel; AEventId: NativeUInt); stdcall;
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

procedure Label_SetOnMouseEnter(AObj: TLabel; AEventId: NativeUInt); stdcall;
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

procedure Label_SetOnMouseLeave(AObj: TLabel; AEventId: NativeUInt); stdcall;
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

function Label_GetCanvas(AObj: TLabel): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function Label_GetAction(AObj: TLabel): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Label_SetAction(AObj: TLabel; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

procedure Label_GetBoundsRect(AObj: TLabel; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Label_SetBoundsRect(AObj: TLabel; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function Label_GetClientHeight(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Label_SetClientHeight(AObj: TLabel; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure Label_GetClientRect(AObj: TLabel; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Label_GetClientWidth(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Label_SetClientWidth(AObj: TLabel; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Label_GetExplicitLeft(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Label_GetExplicitTop(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Label_GetExplicitWidth(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Label_GetExplicitHeight(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Label_GetParent(AObj: TLabel): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Label_SetParent(AObj: TLabel; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Label_GetAlignWithMargins(AObj: TLabel): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Label_SetAlignWithMargins(AObj: TLabel; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Label_GetLeft(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Label_SetLeft(AObj: TLabel; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Label_GetTop(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Label_SetTop(AObj: TLabel; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Label_GetWidth(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Label_SetWidth(AObj: TLabel; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Label_GetHeight(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Label_SetHeight(AObj: TLabel; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Label_GetCursor(AObj: TLabel): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Label_SetCursor(AObj: TLabel; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Label_GetHint(AObj: TLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Label_SetHint(AObj: TLabel; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Label_GetMargins(AObj: TLabel): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Label_SetMargins(AObj: TLabel; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Label_GetCustomHint(AObj: TLabel): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Label_SetCustomHint(AObj: TLabel; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Label_GetComponentCount(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Label_GetComponentIndex(AObj: TLabel): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Label_SetComponentIndex(AObj: TLabel; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Label_GetOwner(AObj: TLabel): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Label_GetName(AObj: TLabel): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Label_SetName(AObj: TLabel; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Label_GetTag(AObj: TLabel): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Label_SetTag(AObj: TLabel; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Label_GetComponents(AObj: TLabel; AIndex: Integer): TComponent; stdcall;
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
  Label_GetEllipsisPosition,
  Label_SetEllipsisPosition,
  Label_GetEnabled,
  Label_SetEnabled,
  Label_GetFont,
  Label_SetFont,
  Label_GetGlowSize,
  Label_SetGlowSize,
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
  Label_GetStyleElements,
  Label_SetStyleElements,
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
  Label_GetExplicitLeft,
  Label_GetExplicitTop,
  Label_GetExplicitWidth,
  Label_GetExplicitHeight,
  Label_GetParent,
  Label_SetParent,
  Label_GetAlignWithMargins,
  Label_SetAlignWithMargins,
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
  Label_GetMargins,
  Label_SetMargins,
  Label_GetCustomHint,
  Label_SetCustomHint,
  Label_GetComponentCount,
  Label_GetComponentIndex,
  Label_SetComponentIndex,
  Label_GetOwner,
  Label_GetName,
  Label_SetName,
  Label_GetTag,
  Label_SetTag,
  Label_GetComponents;

