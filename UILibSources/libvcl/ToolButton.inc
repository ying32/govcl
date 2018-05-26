
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function ToolButton_Create(AOwner: TComponent): TToolButton; stdcall;
begin
  Result :=  TToolButton.Create(AOwner);
end;

procedure ToolButton_Free(AObj: TToolButton); stdcall;
begin
  AObj.Free;
end;

function ToolButton_CheckMenuDropdown(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.CheckMenuDropdown;
end;

procedure ToolButton_Click(AObj: TToolButton); stdcall;
begin
  AObj.Click;
end;

procedure ToolButton_SetBounds(AObj: TToolButton; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure ToolButton_BringToFront(AObj: TToolButton); stdcall;
begin
  AObj.BringToFront;
end;

function ToolButton_HasParent(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure ToolButton_Hide(AObj: TToolButton); stdcall;
begin
  AObj.Hide;
end;

procedure ToolButton_Invalidate(AObj: TToolButton); stdcall;
begin
  AObj.Invalidate;
end;

function ToolButton_Perform(AObj: TToolButton; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure ToolButton_Refresh(AObj: TToolButton); stdcall;
begin
  AObj.Refresh;
end;

procedure ToolButton_Repaint(AObj: TToolButton); stdcall;
begin
  AObj.Repaint;
end;

procedure ToolButton_SendToBack(AObj: TToolButton); stdcall;
begin
  AObj.SendToBack;
end;

procedure ToolButton_Show(AObj: TToolButton); stdcall;
begin
  AObj.Show;
end;

procedure ToolButton_Update(AObj: TToolButton); stdcall;
begin
  AObj.Update;
end;

function ToolButton_GetTextBuf(AObj: TToolButton; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function ToolButton_GetTextLen(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function ToolButton_FindComponent(AObj: TToolButton; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function ToolButton_GetNamePath(AObj: TToolButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure ToolButton_Assign(AObj: TToolButton; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function ToolButton_ClassName(AObj: TToolButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function ToolButton_Equals(AObj: TToolButton; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function ToolButton_GetHashCode(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function ToolButton_ToString(AObj: TToolButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function ToolButton_GetIndex(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.Index;
end;

function ToolButton_GetAction(AObj: TToolButton): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure ToolButton_SetAction(AObj: TToolButton; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function ToolButton_GetAllowAllUp(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.AllowAllUp;
end;

procedure ToolButton_SetAllowAllUp(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.AllowAllUp := AValue;
end;

function ToolButton_GetAutoSize(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure ToolButton_SetAutoSize(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function ToolButton_GetCaption(AObj: TToolButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure ToolButton_SetCaption(AObj: TToolButton; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function ToolButton_GetDown(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.Down;
end;

procedure ToolButton_SetDown(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.Down := AValue;
end;

function ToolButton_GetDropdownMenu(AObj: TToolButton): TPopupMenu; stdcall;
begin
  Result :=  AObj.DropdownMenu;
end;

procedure ToolButton_SetDropdownMenu(AObj: TToolButton; AValue: TPopupMenu); stdcall;
begin
  AObj.DropdownMenu := AValue;
end;

function ToolButton_GetEnabled(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure ToolButton_SetEnabled(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function ToolButton_GetEnableDropdown(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.EnableDropdown;
end;

procedure ToolButton_SetEnableDropdown(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.EnableDropdown := AValue;
end;

function ToolButton_GetGrouped(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.Grouped;
end;

procedure ToolButton_SetGrouped(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.Grouped := AValue;
end;

function ToolButton_GetHeight(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure ToolButton_SetHeight(AObj: TToolButton; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function ToolButton_GetImageIndex(AObj: TToolButton): TImageIndex; stdcall;
begin
  Result :=  AObj.ImageIndex;
end;

procedure ToolButton_SetImageIndex(AObj: TToolButton; AValue: TImageIndex); stdcall;
begin
  AObj.ImageIndex := AValue;
end;

function ToolButton_GetIndeterminate(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.Indeterminate;
end;

procedure ToolButton_SetIndeterminate(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.Indeterminate := AValue;
end;

function ToolButton_GetMarked(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.Marked;
end;

procedure ToolButton_SetMarked(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.Marked := AValue;
end;

function ToolButton_GetParentShowHint(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure ToolButton_SetParentShowHint(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function ToolButton_GetPopupMenu(AObj: TToolButton): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure ToolButton_SetPopupMenu(AObj: TToolButton; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function ToolButton_GetWrap(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.Wrap;
end;

procedure ToolButton_SetWrap(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.Wrap := AValue;
end;

function ToolButton_GetShowHint(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure ToolButton_SetShowHint(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function ToolButton_GetStyle(AObj: TToolButton): TToolButtonStyle; stdcall;
begin
  Result :=  AObj.Style;
end;

procedure ToolButton_SetStyle(AObj: TToolButton; AValue: TToolButtonStyle); stdcall;
begin
  AObj.Style := AValue;
end;

function ToolButton_GetVisible(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure ToolButton_SetVisible(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function ToolButton_GetWidth(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure ToolButton_SetWidth(AObj: TToolButton; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

procedure ToolButton_SetOnClick(AObj: TToolButton; AEventId: NativeUInt); stdcall;
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

procedure ToolButton_SetOnMouseDown(AObj: TToolButton; AEventId: NativeUInt); stdcall;
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

procedure ToolButton_SetOnMouseEnter(AObj: TToolButton; AEventId: NativeUInt); stdcall;
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

procedure ToolButton_SetOnMouseLeave(AObj: TToolButton; AEventId: NativeUInt); stdcall;
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

procedure ToolButton_SetOnMouseMove(AObj: TToolButton; AEventId: NativeUInt); stdcall;
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

procedure ToolButton_SetOnMouseUp(AObj: TToolButton; AEventId: NativeUInt); stdcall;
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

function ToolButton_GetAlign(AObj: TToolButton): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure ToolButton_SetAlign(AObj: TToolButton; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function ToolButton_GetAnchors(AObj: TToolButton): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure ToolButton_SetAnchors(AObj: TToolButton; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function ToolButton_GetBiDiMode(AObj: TToolButton): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure ToolButton_SetBiDiMode(AObj: TToolButton; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure ToolButton_GetBoundsRect(AObj: TToolButton; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure ToolButton_SetBoundsRect(AObj: TToolButton; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function ToolButton_GetClientHeight(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure ToolButton_SetClientHeight(AObj: TToolButton; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure ToolButton_GetClientRect(AObj: TToolButton; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function ToolButton_GetClientWidth(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure ToolButton_SetClientWidth(AObj: TToolButton; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function ToolButton_GetExplicitLeft(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function ToolButton_GetExplicitTop(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function ToolButton_GetExplicitWidth(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function ToolButton_GetExplicitHeight(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function ToolButton_GetParent(AObj: TToolButton): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure ToolButton_SetParent(AObj: TToolButton; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function ToolButton_GetStyleElements(AObj: TToolButton): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure ToolButton_SetStyleElements(AObj: TToolButton; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

function ToolButton_GetAlignWithMargins(AObj: TToolButton): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure ToolButton_SetAlignWithMargins(AObj: TToolButton; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function ToolButton_GetLeft(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure ToolButton_SetLeft(AObj: TToolButton; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function ToolButton_GetTop(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure ToolButton_SetTop(AObj: TToolButton; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function ToolButton_GetCursor(AObj: TToolButton): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure ToolButton_SetCursor(AObj: TToolButton; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function ToolButton_GetHint(AObj: TToolButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure ToolButton_SetHint(AObj: TToolButton; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function ToolButton_GetMargins(AObj: TToolButton): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure ToolButton_SetMargins(AObj: TToolButton; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function ToolButton_GetCustomHint(AObj: TToolButton): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure ToolButton_SetCustomHint(AObj: TToolButton; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function ToolButton_GetComponentCount(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function ToolButton_GetComponentIndex(AObj: TToolButton): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure ToolButton_SetComponentIndex(AObj: TToolButton; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function ToolButton_GetOwner(AObj: TToolButton): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function ToolButton_GetName(AObj: TToolButton): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure ToolButton_SetName(AObj: TToolButton; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function ToolButton_GetTag(AObj: TToolButton): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure ToolButton_SetTag(AObj: TToolButton; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function ToolButton_GetComponents(AObj: TToolButton; AIndex: Integer): TComponent; stdcall;
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
  ToolButton_GetEnableDropdown,
  ToolButton_SetEnableDropdown,
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
  ToolButton_GetExplicitLeft,
  ToolButton_GetExplicitTop,
  ToolButton_GetExplicitWidth,
  ToolButton_GetExplicitHeight,
  ToolButton_GetParent,
  ToolButton_SetParent,
  ToolButton_GetStyleElements,
  ToolButton_SetStyleElements,
  ToolButton_GetAlignWithMargins,
  ToolButton_SetAlignWithMargins,
  ToolButton_GetLeft,
  ToolButton_SetLeft,
  ToolButton_GetTop,
  ToolButton_SetTop,
  ToolButton_GetCursor,
  ToolButton_SetCursor,
  ToolButton_GetHint,
  ToolButton_SetHint,
  ToolButton_GetMargins,
  ToolButton_SetMargins,
  ToolButton_GetCustomHint,
  ToolButton_SetCustomHint,
  ToolButton_GetComponentCount,
  ToolButton_GetComponentIndex,
  ToolButton_SetComponentIndex,
  ToolButton_GetOwner,
  ToolButton_GetName,
  ToolButton_SetName,
  ToolButton_GetTag,
  ToolButton_SetTag,
  ToolButton_GetComponents;

