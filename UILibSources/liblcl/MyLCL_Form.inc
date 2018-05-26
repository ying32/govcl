
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function Form_Create(AOwner: TComponent): TForm; extdecl;
begin
  Result :=  TGoForm.CreateNew(AOwner);
end;

procedure Form_Free(AObj: TGoForm); extdecl;
begin
  AObj.Free;
end;

procedure Form_Close(AObj: TGoForm); extdecl;
begin
  AObj.Close;
end;

procedure Form_Hide(AObj: TGoForm); extdecl;
begin
  AObj.Hide;
end;

procedure Form_SetFocus(AObj: TGoForm); extdecl;
begin
  AObj.SetFocus;
end;

procedure Form_Show(AObj: TGoForm); extdecl;
begin
  AObj.Show;
end;

function Form_ShowModal(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.ShowModal;
end;

function Form_CanFocus(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure Form_FlipChildren(AObj: TGoForm; AllLevels: LongBool); extdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function Form_Focused(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.Focused;
end;

procedure Form_Invalidate(AObj: TGoForm); extdecl;
begin
  AObj.Invalidate;
end;

procedure Form_Realign(AObj: TGoForm); extdecl;
begin
  AObj.Realign;
end;

procedure Form_Repaint(AObj: TGoForm); extdecl;
begin
  AObj.Repaint;
end;

procedure Form_ScaleBy(AObj: TGoForm; M: Integer; D: Integer); extdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure Form_SetBounds(AObj: TGoForm; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); extdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Form_Update(AObj: TGoForm); extdecl;
begin
  AObj.Update;
end;

procedure Form_BringToFront(AObj: TGoForm); extdecl;
begin
  AObj.BringToFront;
end;

function Form_HasParent(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.HasParent;
end;

function Form_Perform(AObj: TGoForm; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; extdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Form_Refresh(AObj: TGoForm); extdecl;
begin
  AObj.Refresh;
end;

procedure Form_SendToBack(AObj: TGoForm); extdecl;
begin
  AObj.SendToBack;
end;

function Form_GetTextBuf(AObj: TGoForm; Buffer: PChar; BufSize: Integer): Integer; extdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Form_GetTextLen(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.GetTextLen;
end;

function Form_FindComponent(AObj: TGoForm; AName: PChar): TComponent; extdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Form_GetNamePath(AObj: TGoForm): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Form_Assign(AObj: TGoForm; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function Form_ClassName(AObj: TGoForm): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function Form_Equals(AObj: TGoForm; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function Form_GetHashCode(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function Form_ToString(AObj: TGoForm): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function Form_GetAction(AObj: TGoForm): TBasicAction; extdecl;
begin
  Result :=  AObj.Action;
end;

procedure Form_SetAction(AObj: TGoForm; AValue: TBasicAction); extdecl;
begin
  AObj.Action := AValue;
end;

function Form_GetAlign(AObj: TGoForm): TAlign; extdecl;
begin
  Result :=  AObj.Align;
end;

procedure Form_SetAlign(AObj: TGoForm; AValue: TAlign); extdecl;
begin
  AObj.Align := AValue;
end;

function Form_GetAlphaBlend(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.AlphaBlend;
end;

procedure Form_SetAlphaBlend(AObj: TGoForm; AValue: LongBool); extdecl;
begin
  AObj.AlphaBlend := AValue;
end;

function Form_GetAlphaBlendValue(AObj: TGoForm): Byte; extdecl;
begin
  Result :=  AObj.AlphaBlendValue;
end;

procedure Form_SetAlphaBlendValue(AObj: TGoForm; AValue: Byte); extdecl;
begin
  AObj.AlphaBlendValue := AValue;
end;

function Form_GetAnchors(AObj: TGoForm): TAnchors; extdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure Form_SetAnchors(AObj: TGoForm; AValue: TAnchors); extdecl;
begin
  AObj.Anchors := AValue;
end;

function Form_GetAutoSize(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.AutoSize;
end;

procedure Form_SetAutoSize(AObj: TGoForm; AValue: LongBool); extdecl;
begin
  AObj.AutoSize := AValue;
end;

function Form_GetBiDiMode(AObj: TGoForm): TBiDiMode; extdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Form_SetBiDiMode(AObj: TGoForm; AValue: TBiDiMode); extdecl;
begin
  AObj.BiDiMode := AValue;
end;

function Form_GetBorderIcons(AObj: TGoForm): TBorderIcons; extdecl;
begin
  Result :=  AObj.BorderIcons;
end;

procedure Form_SetBorderIcons(AObj: TGoForm; AValue: TBorderIcons); extdecl;
begin
  AObj.BorderIcons := AValue;
end;

function Form_GetBorderStyle(AObj: TGoForm): TFormBorderStyle; extdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure Form_SetBorderStyle(AObj: TGoForm; AValue: TFormBorderStyle); extdecl;
begin
  AObj.BorderStyle := AValue;
end;

function Form_GetBorderWidth(AObj: TGoForm): TBorderWidth; extdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure Form_SetBorderWidth(AObj: TGoForm; AValue: TBorderWidth); extdecl;
begin
  AObj.BorderWidth := AValue;
end;

function Form_GetCaption(AObj: TGoForm): PChar; extdecl;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure Form_SetCaption(AObj: TGoForm; AValue: PChar); extdecl;
begin
  AObj.Caption := AValue;
end;

function Form_GetClientHeight(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Form_SetClientHeight(AObj: TGoForm; AValue: Integer); extdecl;
begin
  AObj.ClientHeight := AValue;
end;

function Form_GetClientWidth(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Form_SetClientWidth(AObj: TGoForm; AValue: Integer); extdecl;
begin
  AObj.ClientWidth := AValue;
end;

function Form_GetColor(AObj: TGoForm): TColor; extdecl;
begin
  Result :=  AObj.Color;
end;

procedure Form_SetColor(AObj: TGoForm; AValue: TColor); extdecl;
begin
  AObj.Color := AValue;
end;

function Form_GetDoubleBuffered(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure Form_SetDoubleBuffered(AObj: TGoForm; AValue: LongBool); extdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function Form_GetEnabled(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure Form_SetEnabled(AObj: TGoForm; AValue: LongBool); extdecl;
begin
  AObj.Enabled := AValue;
end;

function Form_GetParentFont(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure Form_SetParentFont(AObj: TGoForm; AValue: LongBool); extdecl;
begin
  AObj.ParentFont := AValue;
end;

function Form_GetFont(AObj: TGoForm): TFont; extdecl;
begin
  Result :=  AObj.Font;
end;

procedure Form_SetFont(AObj: TGoForm; AValue: TFont); extdecl;
begin
  AObj.Font := AValue;
end;

function Form_GetFormStyle(AObj: TGoForm): TFormStyle; extdecl;
begin
  Result :=  AObj.FormStyle;
end;

procedure Form_SetFormStyle(AObj: TGoForm; AValue: TFormStyle); extdecl;
begin
  AObj.FormStyle := AValue;
end;

function Form_GetHeight(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.Height;
end;

procedure Form_SetHeight(AObj: TGoForm; AValue: Integer); extdecl;
begin
  AObj.Height := AValue;
end;

function Form_GetIcon(AObj: TGoForm): TIcon; extdecl;
begin
  Result :=  AObj.Icon;
end;

procedure Form_SetIcon(AObj: TGoForm; AValue: TIcon); extdecl;
begin
  AObj.Icon := AValue;
end;

function Form_GetKeyPreview(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.KeyPreview;
end;

procedure Form_SetKeyPreview(AObj: TGoForm; AValue: LongBool); extdecl;
begin
  AObj.KeyPreview := AValue;
end;

function Form_GetMenu(AObj: TGoForm): TMainMenu; extdecl;
begin
  Result :=  AObj.Menu;
end;

procedure Form_SetMenu(AObj: TGoForm; AValue: TMainMenu); extdecl;
begin
  AObj.Menu := AValue;
end;

function Form_GetPixelsPerInch(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.PixelsPerInch;
end;

procedure Form_SetPixelsPerInch(AObj: TGoForm; AValue: Integer); extdecl;
begin
  AObj.PixelsPerInch := AValue;
end;

function Form_GetPopupMenu(AObj: TGoForm): TPopupMenu; extdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Form_SetPopupMenu(AObj: TGoForm; AValue: TPopupMenu); extdecl;
begin
  AObj.PopupMenu := AValue;
end;

function Form_GetPosition(AObj: TGoForm): TPosition; extdecl;
begin
  Result :=  AObj.Position;
end;

procedure Form_SetPosition(AObj: TGoForm; AValue: TPosition); extdecl;
begin
  AObj.Position := AValue;
end;

function Form_GetShowHint(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure Form_SetShowHint(AObj: TGoForm; AValue: LongBool); extdecl;
begin
  AObj.ShowHint := AValue;
end;

function Form_GetVisible(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.Visible;
end;

procedure Form_SetVisible(AObj: TGoForm; AValue: LongBool); extdecl;
begin
  AObj.Visible := AValue;
end;

function Form_GetWidth(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.Width;
end;

procedure Form_SetWidth(AObj: TGoForm; AValue: Integer); extdecl;
begin
  AObj.Width := AValue;
end;

function Form_GetWindowState(AObj: TGoForm): TWindowState; extdecl;
begin
  Result :=  AObj.WindowState;
end;

procedure Form_SetWindowState(AObj: TGoForm; AValue: TWindowState); extdecl;
begin
  AObj.WindowState := AValue;
end;

procedure Form_SetOnClick(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnClose(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnClose := nil;
    EventClass.Remove(AObj, geFormClose);
    Exit;
  end;
  AObj.OnClose := @EventClass.FormOnClose;
  EventClass.Add(AObj, geFormClose, AEventId);
end;

procedure Form_SetOnCloseQuery(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnCloseQuery := nil;
    EventClass.Remove(AObj, geFormCloseQuery);
    Exit;
  end;
  AObj.OnCloseQuery := @EventClass.FormOnCloseQuery;
  EventClass.Add(AObj, geFormCloseQuery, AEventId);
end;

procedure Form_SetOnDblClick(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnHide(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnHide := nil;
    EventClass.Remove(AObj, geHide);
    Exit;
  end;
  AObj.OnHide := @EventClass.OnHide;
  EventClass.Add(AObj, geHide, AEventId);
end;

procedure Form_SetOnKeyDown(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnKeyPress(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnKeyUp(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnMouseDown(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnMouseEnter(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnMouseLeave(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnMouseMove(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnMouseUp(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnMouseWheel(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheel := nil;
    EventClass.Remove(AObj, geMouseWheel);
    Exit;
  end;
  AObj.OnMouseWheel := @EventClass.OnMouseWheel;
  EventClass.Add(AObj, geMouseWheel, AEventId);
end;

procedure Form_SetOnPaint(AObj: TGoForm; AEventId: NativeUInt); extdecl;
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

procedure Form_SetOnResize(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnResize := nil;
    EventClass.Remove(AObj, geResize);
    Exit;
  end;
  AObj.OnResize := @EventClass.OnResize;
  EventClass.Add(AObj, geResize, AEventId);
end;

procedure Form_SetOnShow(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnShow := nil;
    EventClass.Remove(AObj, geShow);
    Exit;
  end;
  AObj.OnShow := @EventClass.OnShow;
  EventClass.Add(AObj, geShow, AEventId);
end;

function Form_GetCanvas(AObj: TGoForm): TCanvas; extdecl;
begin
  Result :=  AObj.Canvas;
end;

function Form_GetModalResult(AObj: TGoForm): TModalResult; extdecl;
begin
  Result :=  AObj.ModalResult;
end;

procedure Form_SetModalResult(AObj: TGoForm; AValue: TModalResult); extdecl;
begin
  AObj.ModalResult := AValue;
end;

function Form_GetLeft(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.Left;
end;

procedure Form_SetLeft(AObj: TGoForm; AValue: Integer); extdecl;
begin
  AObj.Left := AValue;
end;

function Form_GetTop(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.Top;
end;

procedure Form_SetTop(AObj: TGoForm; AValue: Integer); extdecl;
begin
  AObj.Top := AValue;
end;

function Form_GetBrush(AObj: TGoForm): TBrush; extdecl;
begin
  Result :=  AObj.Brush;
end;

function Form_GetControlCount(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.ControlCount;
end;

function Form_GetHandle(AObj: TGoForm): HWND; extdecl;
begin
  Result :=  AObj.Handle;
end;

function Form_GetParentWindow(AObj: TGoForm): HWND; extdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure Form_SetParentWindow(AObj: TGoForm; AValue: HWND); extdecl;
begin
  AObj.ParentWindow := AValue;
end;

function Form_GetTabOrder(AObj: TGoForm): TTabOrder; extdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure Form_SetTabOrder(AObj: TGoForm; AValue: TTabOrder); extdecl;
begin
  AObj.TabOrder := AValue;
end;

function Form_GetTabStop(AObj: TGoForm): LongBool; extdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure Form_SetTabStop(AObj: TGoForm; AValue: LongBool); extdecl;
begin
  AObj.TabStop := AValue;
end;

procedure Form_GetBoundsRect(AObj: TGoForm; var Result: TRect); extdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Form_SetBoundsRect(AObj: TGoForm; var AValue: TRect); extdecl;
begin
  AObj.BoundsRect := AValue;
end;

procedure Form_GetClientRect(AObj: TGoForm; var Result: TRect); extdecl;
begin
  Result :=  AObj.ClientRect;
end;

function Form_GetParent(AObj: TGoForm): TWinControl; extdecl;
begin
  Result :=  AObj.Parent;
end;

procedure Form_SetParent(AObj: TGoForm; AValue: TWinControl); extdecl;
begin
  AObj.Parent := AValue;
end;

function Form_GetCursor(AObj: TGoForm): TCursor; extdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure Form_SetCursor(AObj: TGoForm; AValue: TCursor); extdecl;
begin
  AObj.Cursor := AValue;
end;

function Form_GetHint(AObj: TGoForm): PChar; extdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Form_SetHint(AObj: TGoForm; AValue: PChar); extdecl;
begin
  AObj.Hint := AValue;
end;

function Form_GetComponentCount(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function Form_GetComponentIndex(AObj: TGoForm): Integer; extdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Form_SetComponentIndex(AObj: TGoForm; AValue: Integer); extdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function Form_GetOwner(AObj: TGoForm): TComponent; extdecl;
begin
  Result :=  AObj.Owner;
end;

function Form_GetName(AObj: TGoForm): PChar; extdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Form_SetName(AObj: TGoForm; AValue: PChar); extdecl;
begin
  AObj.Name := AValue;
end;

function Form_GetTag(AObj: TGoForm): NativeInt; extdecl;
begin
  Result :=  AObj.Tag;
end;

procedure Form_SetTag(AObj: TGoForm; AValue: NativeInt); extdecl;
begin
  AObj.Tag := AValue;
end;

function Form_GetControls(AObj: TGoForm; Index: Integer): TControl; extdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function Form_GetComponents(AObj: TGoForm; AIndex: Integer): TComponent; extdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Form_Create,
  Form_Free,
  Form_Close,
  Form_Hide,
  Form_SetFocus,
  Form_Show,
  Form_ShowModal,
  Form_CanFocus,
  Form_FlipChildren,
  Form_Focused,
  Form_Invalidate,
  Form_Realign,
  Form_Repaint,
  Form_ScaleBy,
  Form_SetBounds,
  Form_Update,
  Form_BringToFront,
  Form_HasParent,
  Form_Perform,
  Form_Refresh,
  Form_SendToBack,
  Form_GetTextBuf,
  Form_GetTextLen,
  Form_FindComponent,
  Form_GetNamePath,
  Form_Assign,
  Form_ClassName,
  Form_Equals,
  Form_GetHashCode,
  Form_ToString,
  Form_GetAction,
  Form_SetAction,
  Form_GetAlign,
  Form_SetAlign,
  Form_GetAlphaBlend,
  Form_SetAlphaBlend,
  Form_GetAlphaBlendValue,
  Form_SetAlphaBlendValue,
  Form_GetAnchors,
  Form_SetAnchors,
  Form_GetAutoSize,
  Form_SetAutoSize,
  Form_GetBiDiMode,
  Form_SetBiDiMode,
  Form_GetBorderIcons,
  Form_SetBorderIcons,
  Form_GetBorderStyle,
  Form_SetBorderStyle,
  Form_GetBorderWidth,
  Form_SetBorderWidth,
  Form_GetCaption,
  Form_SetCaption,
  Form_GetClientHeight,
  Form_SetClientHeight,
  Form_GetClientWidth,
  Form_SetClientWidth,
  Form_GetColor,
  Form_SetColor,
  Form_GetDoubleBuffered,
  Form_SetDoubleBuffered,
  Form_GetEnabled,
  Form_SetEnabled,
  Form_GetParentFont,
  Form_SetParentFont,
  Form_GetFont,
  Form_SetFont,
  Form_GetFormStyle,
  Form_SetFormStyle,
  Form_GetHeight,
  Form_SetHeight,
  Form_GetIcon,
  Form_SetIcon,
  Form_GetKeyPreview,
  Form_SetKeyPreview,
  Form_GetMenu,
  Form_SetMenu,
  Form_GetPixelsPerInch,
  Form_SetPixelsPerInch,
  Form_GetPopupMenu,
  Form_SetPopupMenu,
  Form_GetPosition,
  Form_SetPosition,
  Form_GetShowHint,
  Form_SetShowHint,
  Form_GetVisible,
  Form_SetVisible,
  Form_GetWidth,
  Form_SetWidth,
  Form_GetWindowState,
  Form_SetWindowState,
  Form_SetOnClick,
  Form_SetOnClose,
  Form_SetOnCloseQuery,
  Form_SetOnDblClick,
  Form_SetOnHide,
  Form_SetOnKeyDown,
  Form_SetOnKeyPress,
  Form_SetOnKeyUp,
  Form_SetOnMouseDown,
  Form_SetOnMouseEnter,
  Form_SetOnMouseLeave,
  Form_SetOnMouseMove,
  Form_SetOnMouseUp,
  Form_SetOnMouseWheel,
  Form_SetOnPaint,
  Form_SetOnResize,
  Form_SetOnShow,
  Form_GetCanvas,
  Form_GetModalResult,
  Form_SetModalResult,
  Form_GetLeft,
  Form_SetLeft,
  Form_GetTop,
  Form_SetTop,
  Form_GetBrush,
  Form_GetControlCount,
  Form_GetHandle,
  Form_GetParentWindow,
  Form_SetParentWindow,
  Form_GetTabOrder,
  Form_SetTabOrder,
  Form_GetTabStop,
  Form_SetTabStop,
  Form_GetBoundsRect,
  Form_SetBoundsRect,
  Form_GetClientRect,
  Form_GetParent,
  Form_SetParent,
  Form_GetCursor,
  Form_SetCursor,
  Form_GetHint,
  Form_SetHint,
  Form_GetComponentCount,
  Form_GetComponentIndex,
  Form_SetComponentIndex,
  Form_GetOwner,
  Form_GetName,
  Form_SetName,
  Form_GetTag,
  Form_SetTag,
  Form_GetControls,
  Form_GetComponents;

