
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function ToolButton_Create(AOwner: TComponent): TToolButton; extdecl;
begin
  Result :=  TToolButton.Create(AOwner);
end;

procedure ToolButton_Free(AObj: TToolButton); extdecl;
begin
  AObj.Free;
end;

function ToolButton_CheckMenuDropdown(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.CheckMenuDropdown;
end;

procedure ToolButton_Click(AObj: TToolButton); extdecl;
begin
  AObj.Click;
end;

procedure ToolButton_SetBounds(AObj: TToolButton; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ToolButton_BringToFront(AObj: TToolButton); extdecl;
begin
  AObj.BringToFront;
end;

function ToolButton_HasParent(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure ToolButton_Hide(AObj: TToolButton); extdecl;
begin
  AObj.Hide;
end;

procedure ToolButton_Invalidate(AObj: TToolButton); extdecl;
begin
  AObj.Invalidate;
end;

function ToolButton_Perform(AObj: TToolButton; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ToolButton_Refresh(AObj: TToolButton); extdecl;
begin
  AObj.Refresh;
end;

procedure ToolButton_Repaint(AObj: TToolButton); extdecl;
begin
  AObj.Repaint;
end;

procedure ToolButton_SendToBack(AObj: TToolButton); extdecl;
begin
  AObj.SendToBack;
end;

procedure ToolButton_Show(AObj: TToolButton); extdecl;
begin
  AObj.Show;
end;

procedure ToolButton_Update(AObj: TToolButton); extdecl;
begin
  AObj.Update;
end;

function ToolButton_GetTextBuf(AObj: TToolButton; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ToolButton_GetTextLen(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function ToolButton_FindComponent(AObj: TToolButton; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ToolButton_GetNamePath(AObj: TToolButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ToolButton_Assign(AObj: TToolButton; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function ToolButton_ClassName(AObj: TToolButton): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function ToolButton_Equals(AObj: TToolButton; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function ToolButton_GetHashCode(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function ToolButton_ToString(AObj: TToolButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function ToolButton_GetIndex(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.Index;
end;

function ToolButton_GetAction(AObj: TToolButton): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure ToolButton_SetAction(AObj: TToolButton; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function ToolButton_GetAllowAllUp(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.AllowAllUp;
end;

procedure ToolButton_SetAllowAllUp(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.AllowAllUp := AValue;
end;

function ToolButton_GetAutoSize(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.AutoSize;
end;

procedure ToolButton_SetAutoSize(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.AutoSize := AValue;
end;

function ToolButton_GetCaption(AObj: TToolButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure ToolButton_SetCaption(AObj: TToolButton; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function ToolButton_GetDown(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.Down;
end;

procedure ToolButton_SetDown(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.Down := AValue;
end;

function ToolButton_GetDropdownMenu(AObj: TToolButton): TPopupMenu; extdecl;
begin
  Result :=  AObj.DropdownMenu;
end;

procedure ToolButton_SetDropdownMenu(AObj: TToolButton; AValue: TPopupMenu); extdecl;
begin
  AObj.DropdownMenu := AValue;
end;

function ToolButton_GetEnabled(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure ToolButton_SetEnabled(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function ToolButton_GetGrouped(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.Grouped;
end;

procedure ToolButton_SetGrouped(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.Grouped := AValue;
end;

function ToolButton_GetHeight(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure ToolButton_SetHeight(AObj: TToolButton; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function ToolButton_GetImageIndex(AObj: TToolButton): TImageIndex; extdecl;
begin
  Result :=  AObj.ImageIndex;
end;

procedure ToolButton_SetImageIndex(AObj: TToolButton; AValue: TImageIndex); extdecl;
begin
  AObj.ImageIndex := AValue;
end;

function ToolButton_GetIndeterminate(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.Indeterminate;
end;

procedure ToolButton_SetIndeterminate(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.Indeterminate := AValue;
end;

function ToolButton_GetMarked(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.Marked;
end;

procedure ToolButton_SetMarked(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.Marked := AValue;
end;

function ToolButton_GetParentShowHint(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ToolButton_SetParentShowHint(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function ToolButton_GetPopupMenu(AObj: TToolButton): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ToolButton_SetPopupMenu(AObj: TToolButton; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function ToolButton_GetWrap(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.Wrap;
end;

procedure ToolButton_SetWrap(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.Wrap := AValue;
end;

function ToolButton_GetShowHint(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure ToolButton_SetShowHint(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function ToolButton_GetStyle(AObj: TToolButton): TToolButtonStyle; extdecl;
begin
  Result :=  AObj.Style;
end;

procedure ToolButton_SetStyle(AObj: TToolButton; AValue: TToolButtonStyle); extdecl;
begin
  AObj.Style := AValue;
end;

function ToolButton_GetVisible(AObj: TToolButton): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure ToolButton_SetVisible(AObj: TToolButton; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function ToolButton_GetWidth(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure ToolButton_SetWidth(AObj: TToolButton; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

procedure ToolButton_SetOnClick(AObj: TToolButton; AEventId: NativeUInt); extdecl;
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

procedure ToolButton_SetOnMouseDown(AObj: TToolButton; AEventId: NativeUInt); extdecl;
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

procedure ToolButton_SetOnMouseEnter(AObj: TToolButton; AEventId: NativeUInt); extdecl;
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

procedure ToolButton_SetOnMouseLeave(AObj: TToolButton; AEventId: NativeUInt); extdecl;
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

procedure ToolButton_SetOnMouseMove(AObj: TToolButton; AEventId: NativeUInt); extdecl;
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

procedure ToolButton_SetOnMouseUp(AObj: TToolButton; AEventId: NativeUInt); extdecl;
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

function ToolButton_GetAlign(AObj: TToolButton): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure ToolButton_SetAlign(AObj: TToolButton; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function ToolButton_GetAnchors(AObj: TToolButton): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure ToolButton_SetAnchors(AObj: TToolButton; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function ToolButton_GetBiDiMode(AObj: TToolButton): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ToolButton_SetBiDiMode(AObj: TToolButton; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure ToolButton_GetBoundsRect(AObj: TToolButton; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ToolButton_SetBoundsRect(AObj: TToolButton; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

function ToolButton_GetClientHeight(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ToolButton_SetClientHeight(AObj: TToolButton; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure ToolButton_GetClientRect(AObj: TToolButton; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function ToolButton_GetClientWidth(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ToolButton_SetClientWidth(AObj: TToolButton; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function ToolButton_GetParent(AObj: TToolButton): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure ToolButton_SetParent(AObj: TToolButton; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function ToolButton_GetLeft(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure ToolButton_SetLeft(AObj: TToolButton; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function ToolButton_GetTop(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure ToolButton_SetTop(AObj: TToolButton; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function ToolButton_GetCursor(AObj: TToolButton): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure ToolButton_SetCursor(AObj: TToolButton; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function ToolButton_GetHint(AObj: TToolButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ToolButton_SetHint(AObj: TToolButton; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function ToolButton_GetComponentCount(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function ToolButton_GetComponentIndex(AObj: TToolButton): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ToolButton_SetComponentIndex(AObj: TToolButton; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function ToolButton_GetOwner(AObj: TToolButton): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function ToolButton_GetName(AObj: TToolButton): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ToolButton_SetName(AObj: TToolButton; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function ToolButton_GetTag(AObj: TToolButton): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure ToolButton_SetTag(AObj: TToolButton; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function ToolButton_GetComponents(AObj: TToolButton; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  ToolButton_Create,
  ToolButton_Free,
  ToolButton_CheckMenuDropdown,
  ToolButton_Click,
  ToolButton_SetBounds,
  ToolButton_BringToFront,
  ToolButton_HasParent,
  ToolButton_Hide,
  ToolButton_Invalidate,
  ToolButton_Perform,
  ToolButton_Refresh,
  ToolButton_Repaint,
  ToolButton_SendToBack,
  ToolButton_Show,
  ToolButton_Update,
  ToolButton_GetTextBuf,
  ToolButton_GetTextLen,
  ToolButton_FindComponent,
  ToolButton_GetNamePath,
  ToolButton_Assign,
  ToolButton_ClassName,
  ToolButton_Equals,
  ToolButton_GetHashCode,
  ToolButton_ToString,
  ToolButton_GetIndex,
  ToolButton_GetAction,
  ToolButton_SetAction,
  ToolButton_GetAllowAllUp,
  ToolButton_SetAllowAllUp,
  ToolButton_GetAutoSize,
  ToolButton_SetAutoSize,
  ToolButton_GetCaption,
  ToolButton_SetCaption,
  ToolButton_GetDown,
  ToolButton_SetDown,
  ToolButton_GetDropdownMenu,
  ToolButton_SetDropdownMenu,
  ToolButton_GetEnabled,
  ToolButton_SetEnabled,
  ToolButton_GetGrouped,
  ToolButton_SetGrouped,
  ToolButton_GetHeight,
  ToolButton_SetHeight,
  ToolButton_GetImageIndex,
  ToolButton_SetImageIndex,
  ToolButton_GetIndeterminate,
  ToolButton_SetIndeterminate,
  ToolButton_GetMarked,
  ToolButton_SetMarked,
  ToolButton_GetParentShowHint,
  ToolButton_SetParentShowHint,
  ToolButton_GetPopupMenu,
  ToolButton_SetPopupMenu,
  ToolButton_GetWrap,
  ToolButton_SetWrap,
  ToolButton_GetShowHint,
  ToolButton_SetShowHint,
  ToolButton_GetStyle,
  ToolButton_SetStyle,
  ToolButton_GetVisible,
  ToolButton_SetVisible,
  ToolButton_GetWidth,
  ToolButton_SetWidth,
  ToolButton_SetOnClick,
  ToolButton_SetOnMouseDown,
  ToolButton_SetOnMouseEnter,
  ToolButton_SetOnMouseLeave,
  ToolButton_SetOnMouseMove,
  ToolButton_SetOnMouseUp,
  ToolButton_GetAlign,
  ToolButton_SetAlign,
  ToolButton_GetAnchors,
  ToolButton_SetAnchors,
  ToolButton_GetBiDiMode,
  ToolButton_SetBiDiMode,
  ToolButton_GetBoundsRect,
  ToolButton_SetBoundsRect,
  ToolButton_GetClientHeight,
  ToolButton_SetClientHeight,
  ToolButton_GetClientRect,
  ToolButton_GetClientWidth,
  ToolButton_SetClientWidth,
  ToolButton_GetParent,
  ToolButton_SetParent,
  ToolButton_GetLeft,
  ToolButton_SetLeft,
  ToolButton_GetTop,
  ToolButton_SetTop,
  ToolButton_GetCursor,
  ToolButton_SetCursor,
  ToolButton_GetHint,
  ToolButton_SetHint,
  ToolButton_GetComponentCount,
  ToolButton_GetComponentIndex,
  ToolButton_SetComponentIndex,
  ToolButton_GetOwner,
  ToolButton_GetName,
  ToolButton_SetName,
  ToolButton_GetTag,
  ToolButton_SetTag,
  ToolButton_GetComponents;

