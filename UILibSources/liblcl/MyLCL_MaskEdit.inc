
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function MaskEdit_Create(AOwner: TComponent): TMaskEdit; extdecl;
begin
  Result :=  TMaskEdit.Create(AOwner);
end;

procedure MaskEdit_Free(AObj: TMaskEdit); extdecl;
begin
  AObj.Free;
end;

procedure MaskEdit_ValidateEdit(AObj: TMaskEdit); extdecl;
begin
  AObj.ValidateEdit;
end;

procedure MaskEdit_Clear(AObj: TMaskEdit); extdecl;
begin
  AObj.Clear;
end;

function MaskEdit_GetTextLen(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

procedure MaskEdit_ClearSelection(AObj: TMaskEdit); extdecl;
begin
  AObj.ClearSelection;
end;

procedure MaskEdit_CopyToClipboard(AObj: TMaskEdit); extdecl;
begin
  AObj.CopyToClipboard;
end;

procedure MaskEdit_CutToClipboard(AObj: TMaskEdit); extdecl;
begin
  AObj.CutToClipboard;
end;

procedure MaskEdit_PasteFromClipboard(AObj: TMaskEdit); extdecl;
begin
  AObj.PasteFromClipboard;
end;

procedure MaskEdit_Undo(AObj: TMaskEdit); extdecl;
begin
  AObj.Undo;
end;

procedure MaskEdit_SelectAll(AObj: TMaskEdit); extdecl;
begin
  AObj.SelectAll;
end;

function MaskEdit_CanFocus(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure MaskEdit_FlipChildren(AObj: TMaskEdit; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function MaskEdit_Focused(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure MaskEdit_Invalidate(AObj: TMaskEdit); extdecl;
begin
  AObj.Invalidate;
end;

procedure MaskEdit_Realign(AObj: TMaskEdit); extdecl;
begin
  AObj.Realign;
end;

procedure MaskEdit_Repaint(AObj: TMaskEdit); extdecl;
begin
  AObj.Repaint;
end;

procedure MaskEdit_ScaleBy(AObj: TMaskEdit; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure MaskEdit_SetBounds(AObj: TMaskEdit; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure MaskEdit_SetFocus(AObj: TMaskEdit); extdecl;
begin
  AObj.SetFocus;
end;

procedure MaskEdit_Update(AObj: TMaskEdit); extdecl;
begin
  AObj.Update;
end;

procedure MaskEdit_BringToFront(AObj: TMaskEdit); extdecl;
begin
  AObj.BringToFront;
end;

function MaskEdit_HasParent(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure MaskEdit_Hide(AObj: TMaskEdit); extdecl;
begin
  AObj.Hide;
end;

function MaskEdit_Perform(AObj: TMaskEdit; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure MaskEdit_Refresh(AObj: TMaskEdit); extdecl;
begin
  AObj.Refresh;
end;

procedure MaskEdit_SendToBack(AObj: TMaskEdit); extdecl;
begin
  AObj.SendToBack;
end;

procedure MaskEdit_Show(AObj: TMaskEdit); extdecl;
begin
  AObj.Show;
end;

function MaskEdit_GetTextBuf(AObj: TMaskEdit; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function MaskEdit_FindComponent(AObj: TMaskEdit; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function MaskEdit_GetNamePath(AObj: TMaskEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure MaskEdit_Assign(AObj: TMaskEdit; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function MaskEdit_ClassName(AObj: TMaskEdit): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function MaskEdit_Equals(AObj: TMaskEdit; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function MaskEdit_GetHashCode(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function MaskEdit_ToString(AObj: TMaskEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function MaskEdit_GetAlign(AObj: TMaskEdit): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure MaskEdit_SetAlign(AObj: TMaskEdit; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function MaskEdit_GetAlignment(AObj: TMaskEdit): TAlignment; extdecl;
begin
  Result :=  AObj.Alignment;
end;

procedure MaskEdit_SetAlignment(AObj: TMaskEdit; AValue: TAlignment); extdecl;
begin
  AObj.Alignment := AValue;
end;

function MaskEdit_GetAnchors(AObj: TMaskEdit): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure MaskEdit_SetAnchors(AObj: TMaskEdit; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function MaskEdit_GetAutoSelect(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.AutoSelect;
end;

procedure MaskEdit_SetAutoSelect(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.AutoSelect := AValue;
end;

function MaskEdit_GetAutoSize(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.AutoSize;
end;

procedure MaskEdit_SetAutoSize(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.AutoSize := AValue;
end;

function MaskEdit_GetBiDiMode(AObj: TMaskEdit): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure MaskEdit_SetBiDiMode(AObj: TMaskEdit; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function MaskEdit_GetBorderStyle(AObj: TMaskEdit): TBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure MaskEdit_SetBorderStyle(AObj: TMaskEdit; AValue: TBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function MaskEdit_GetColor(AObj: TMaskEdit): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure MaskEdit_SetColor(AObj: TMaskEdit; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function MaskEdit_GetDoubleBuffered(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure MaskEdit_SetDoubleBuffered(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function MaskEdit_GetEnabled(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure MaskEdit_SetEnabled(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function MaskEdit_GetFont(AObj: TMaskEdit): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure MaskEdit_SetFont(AObj: TMaskEdit; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function MaskEdit_GetMaxLength(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.MaxLength;
end;

procedure MaskEdit_SetMaxLength(AObj: TMaskEdit; AValue: Integer); extdecl;
begin
  AObj.MaxLength := AValue;
end;

function MaskEdit_GetParentColor(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure MaskEdit_SetParentColor(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.ParentColor := AValue;
end;

function MaskEdit_GetParentFont(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure MaskEdit_SetParentFont(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function MaskEdit_GetParentShowHint(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure MaskEdit_SetParentShowHint(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function MaskEdit_GetPasswordChar(AObj: TMaskEdit): Char; extdecl;
begin
  Result :=  AObj.PasswordChar;
end;

procedure MaskEdit_SetPasswordChar(AObj: TMaskEdit; AValue: Char); extdecl;
begin
  AObj.PasswordChar := AValue;
end;

function MaskEdit_GetPopupMenu(AObj: TMaskEdit): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure MaskEdit_SetPopupMenu(AObj: TMaskEdit; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function MaskEdit_GetReadOnly(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.ReadOnly;
end;

procedure MaskEdit_SetReadOnly(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.ReadOnly := AValue;
end;

function MaskEdit_GetShowHint(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure MaskEdit_SetShowHint(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function MaskEdit_GetTabOrder(AObj: TMaskEdit): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure MaskEdit_SetTabOrder(AObj: TMaskEdit; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function MaskEdit_GetTabStop(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure MaskEdit_SetTabStop(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

function MaskEdit_GetText(AObj: TMaskEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.Text);
end;

procedure MaskEdit_SetText(AObj: TMaskEdit; AValue: PChar); extdecl;
begin
  AObj.Text := AValue;
end;

function MaskEdit_GetVisible(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure MaskEdit_SetVisible(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

procedure MaskEdit_SetOnChange(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    EventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := @EventClass.OnChange;
  EventClass.Add(AObj, geChange, AEventId);
end;

procedure MaskEdit_SetOnClick(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
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

procedure MaskEdit_SetOnDblClick(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
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

procedure MaskEdit_SetOnEnter(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
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

procedure MaskEdit_SetOnExit(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
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

procedure MaskEdit_SetOnKeyDown(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyDown := nil;
    EventClass.Remove(AObj, geKeyDown);
    Exit;
  end;
  AObj.OnKeyDown := @EventClass.OnKeyDown;
  EventClass.Add(AObj, geKeyDown, AEventId);
end;

procedure MaskEdit_SetOnKeyPress(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyPress := nil;
    EventClass.Remove(AObj, geKeyPress);
    Exit;
  end;
  AObj.OnKeyPress := @EventClass.OnKeyPress;
  EventClass.Add(AObj, geKeyPress, AEventId);
end;

procedure MaskEdit_SetOnKeyUp(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyUp := nil;
    EventClass.Remove(AObj, geKeyUp);
    Exit;
  end;
  AObj.OnKeyUp := @EventClass.OnKeyUp;
  EventClass.Add(AObj, geKeyUp, AEventId);
end;

procedure MaskEdit_SetOnMouseDown(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
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

procedure MaskEdit_SetOnMouseEnter(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
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

procedure MaskEdit_SetOnMouseLeave(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
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

procedure MaskEdit_SetOnMouseMove(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
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

procedure MaskEdit_SetOnMouseUp(AObj: TMaskEdit; AEventId: NativeUInt); extdecl;
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

function MaskEdit_GetIsMasked(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.IsMasked;
end;

function MaskEdit_GetEditText(AObj: TMaskEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.EditText);
end;

procedure MaskEdit_SetEditText(AObj: TMaskEdit; AValue: PChar); extdecl;
begin
  AObj.EditText := AValue;
end;

function MaskEdit_GetCanUndo(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.CanUndo;
end;

function MaskEdit_GetModified(AObj: TMaskEdit): LongBool; extdecl;
begin
  Result :=  AObj.Modified;
end;

procedure MaskEdit_SetModified(AObj: TMaskEdit; AValue: LongBool); extdecl;
begin
  AObj.Modified := AValue;
end;

function MaskEdit_GetSelLength(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.SelLength;
end;

procedure MaskEdit_SetSelLength(AObj: TMaskEdit; AValue: Integer); extdecl;
begin
  AObj.SelLength := AValue;
end;

function MaskEdit_GetSelStart(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.SelStart;
end;

procedure MaskEdit_SetSelStart(AObj: TMaskEdit; AValue: Integer); extdecl;
begin
  AObj.SelStart := AValue;
end;

function MaskEdit_GetSelText(AObj: TMaskEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.SelText);
end;

procedure MaskEdit_SetSelText(AObj: TMaskEdit; AValue: PChar); extdecl;
begin
  AObj.SelText := AValue;
end;

function MaskEdit_GetBrush(AObj: TMaskEdit): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function MaskEdit_GetControlCount(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function MaskEdit_GetHandle(AObj: TMaskEdit): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function MaskEdit_GetParentWindow(AObj: TMaskEdit): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure MaskEdit_SetParentWindow(AObj: TMaskEdit; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function MaskEdit_GetAction(AObj: TMaskEdit): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure MaskEdit_SetAction(AObj: TMaskEdit; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

procedure MaskEdit_GetBoundsRect(AObj: TMaskEdit; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure MaskEdit_SetBoundsRect(AObj: TMaskEdit; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function MaskEdit_GetClientHeight(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure MaskEdit_SetClientHeight(AObj: TMaskEdit; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure MaskEdit_GetClientRect(AObj: TMaskEdit; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function MaskEdit_GetClientWidth(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure MaskEdit_SetClientWidth(AObj: TMaskEdit; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function MaskEdit_GetParent(AObj: TMaskEdit): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure MaskEdit_SetParent(AObj: TMaskEdit; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function MaskEdit_GetLeft(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure MaskEdit_SetLeft(AObj: TMaskEdit; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function MaskEdit_GetTop(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure MaskEdit_SetTop(AObj: TMaskEdit; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function MaskEdit_GetWidth(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure MaskEdit_SetWidth(AObj: TMaskEdit; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function MaskEdit_GetHeight(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure MaskEdit_SetHeight(AObj: TMaskEdit; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function MaskEdit_GetCursor(AObj: TMaskEdit): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure MaskEdit_SetCursor(AObj: TMaskEdit; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function MaskEdit_GetHint(AObj: TMaskEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure MaskEdit_SetHint(AObj: TMaskEdit; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function MaskEdit_GetComponentCount(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function MaskEdit_GetComponentIndex(AObj: TMaskEdit): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure MaskEdit_SetComponentIndex(AObj: TMaskEdit; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function MaskEdit_GetOwner(AObj: TMaskEdit): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function MaskEdit_GetName(AObj: TMaskEdit): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure MaskEdit_SetName(AObj: TMaskEdit; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function MaskEdit_GetTag(AObj: TMaskEdit): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure MaskEdit_SetTag(AObj: TMaskEdit; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function MaskEdit_GetControls(AObj: TMaskEdit; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function MaskEdit_GetComponents(AObj: TMaskEdit; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  MaskEdit_Create,
  MaskEdit_Free,
  MaskEdit_ValidateEdit,
  MaskEdit_Clear,
  MaskEdit_GetTextLen,
  MaskEdit_ClearSelection,
  MaskEdit_CopyToClipboard,
  MaskEdit_CutToClipboard,
  MaskEdit_PasteFromClipboard,
  MaskEdit_Undo,
  MaskEdit_SelectAll,
  MaskEdit_CanFocus,
  MaskEdit_FlipChildren,
  MaskEdit_Focused,
  MaskEdit_Invalidate,
  MaskEdit_Realign,
  MaskEdit_Repaint,
  MaskEdit_ScaleBy,
  MaskEdit_SetBounds,
  MaskEdit_SetFocus,
  MaskEdit_Update,
  MaskEdit_BringToFront,
  MaskEdit_HasParent,
  MaskEdit_Hide,
  MaskEdit_Perform,
  MaskEdit_Refresh,
  MaskEdit_SendToBack,
  MaskEdit_Show,
  MaskEdit_GetTextBuf,
  MaskEdit_FindComponent,
  MaskEdit_GetNamePath,
  MaskEdit_Assign,
  MaskEdit_ClassName,
  MaskEdit_Equals,
  MaskEdit_GetHashCode,
  MaskEdit_ToString,
  MaskEdit_GetAlign,
  MaskEdit_SetAlign,
  MaskEdit_GetAlignment,
  MaskEdit_SetAlignment,
  MaskEdit_GetAnchors,
  MaskEdit_SetAnchors,
  MaskEdit_GetAutoSelect,
  MaskEdit_SetAutoSelect,
  MaskEdit_GetAutoSize,
  MaskEdit_SetAutoSize,
  MaskEdit_GetBiDiMode,
  MaskEdit_SetBiDiMode,
  MaskEdit_GetBorderStyle,
  MaskEdit_SetBorderStyle,
  MaskEdit_GetColor,
  MaskEdit_SetColor,
  MaskEdit_GetDoubleBuffered,
  MaskEdit_SetDoubleBuffered,
  MaskEdit_GetEnabled,
  MaskEdit_SetEnabled,
  MaskEdit_GetFont,
  MaskEdit_SetFont,
  MaskEdit_GetMaxLength,
  MaskEdit_SetMaxLength,
  MaskEdit_GetParentColor,
  MaskEdit_SetParentColor,
  MaskEdit_GetParentFont,
  MaskEdit_SetParentFont,
  MaskEdit_GetParentShowHint,
  MaskEdit_SetParentShowHint,
  MaskEdit_GetPasswordChar,
  MaskEdit_SetPasswordChar,
  MaskEdit_GetPopupMenu,
  MaskEdit_SetPopupMenu,
  MaskEdit_GetReadOnly,
  MaskEdit_SetReadOnly,
  MaskEdit_GetShowHint,
  MaskEdit_SetShowHint,
  MaskEdit_GetTabOrder,
  MaskEdit_SetTabOrder,
  MaskEdit_GetTabStop,
  MaskEdit_SetTabStop,
  MaskEdit_GetText,
  MaskEdit_SetText,
  MaskEdit_GetVisible,
  MaskEdit_SetVisible,
  MaskEdit_SetOnChange,
  MaskEdit_SetOnClick,
  MaskEdit_SetOnDblClick,
  MaskEdit_SetOnEnter,
  MaskEdit_SetOnExit,
  MaskEdit_SetOnKeyDown,
  MaskEdit_SetOnKeyPress,
  MaskEdit_SetOnKeyUp,
  MaskEdit_SetOnMouseDown,
  MaskEdit_SetOnMouseEnter,
  MaskEdit_SetOnMouseLeave,
  MaskEdit_SetOnMouseMove,
  MaskEdit_SetOnMouseUp,
  MaskEdit_GetIsMasked,
  MaskEdit_GetEditText,
  MaskEdit_SetEditText,
  MaskEdit_GetCanUndo,
  MaskEdit_GetModified,
  MaskEdit_SetModified,
  MaskEdit_GetSelLength,
  MaskEdit_SetSelLength,
  MaskEdit_GetSelStart,
  MaskEdit_SetSelStart,
  MaskEdit_GetSelText,
  MaskEdit_SetSelText,
  MaskEdit_GetBrush,
  MaskEdit_GetControlCount,
  MaskEdit_GetHandle,
  MaskEdit_GetParentWindow,
  MaskEdit_SetParentWindow,
  MaskEdit_GetAction,
  MaskEdit_SetAction,
  MaskEdit_GetBoundsRect,
  MaskEdit_SetBoundsRect,
  MaskEdit_GetClientHeight,
  MaskEdit_SetClientHeight,
  MaskEdit_GetClientRect,
  MaskEdit_GetClientWidth,
  MaskEdit_SetClientWidth,
  MaskEdit_GetParent,
  MaskEdit_SetParent,
  MaskEdit_GetLeft,
  MaskEdit_SetLeft,
  MaskEdit_GetTop,
  MaskEdit_SetTop,
  MaskEdit_GetWidth,
  MaskEdit_SetWidth,
  MaskEdit_GetHeight,
  MaskEdit_SetHeight,
  MaskEdit_GetCursor,
  MaskEdit_SetCursor,
  MaskEdit_GetHint,
  MaskEdit_SetHint,
  MaskEdit_GetComponentCount,
  MaskEdit_GetComponentIndex,
  MaskEdit_SetComponentIndex,
  MaskEdit_GetOwner,
  MaskEdit_GetName,
  MaskEdit_SetName,
  MaskEdit_GetTag,
  MaskEdit_SetTag,
  MaskEdit_GetControls,
  MaskEdit_GetComponents;

