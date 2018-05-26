
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Form_Create(AOwner: TComponent): TGoForm; stdcall;
begin
  Result :=  TGoForm.Create(AOwner);
end;

procedure Form_Free(AObj: TGoForm); stdcall;
begin
  AObj.Free;
end;

procedure Form_Close(AObj: TGoForm); stdcall;
begin
  AObj.Close;
end;

procedure Form_Hide(AObj: TGoForm); stdcall;
begin
  AObj.Hide;
end;

procedure Form_Print(AObj: TGoForm); stdcall;
begin
  AObj.Print;
end;

procedure Form_SetFocus(AObj: TGoForm); stdcall;
begin
  AObj.SetFocus;
end;

procedure Form_Show(AObj: TGoForm); stdcall;
begin
  AObj.Show;
end;

function Form_ShowModal(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.ShowModal;
end;

function Form_CanFocus(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure Form_FlipChildren(AObj: TGoForm; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function Form_Focused(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function Form_HandleAllocated(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure Form_Invalidate(AObj: TGoForm); stdcall;
begin
  AObj.Invalidate;
end;

procedure Form_Realign(AObj: TGoForm); stdcall;
begin
  AObj.Realign;
end;

procedure Form_Repaint(AObj: TGoForm); stdcall;
begin
  AObj.Repaint;
end;

procedure Form_ScaleBy(AObj: TGoForm; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure Form_SetBounds(AObj: TGoForm; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Form_Update(AObj: TGoForm); stdcall;
begin
  AObj.Update;
end;

procedure Form_BringToFront(AObj: TGoForm); stdcall;
begin
  AObj.BringToFront;
end;

function Form_HasParent(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

function Form_Perform(AObj: TGoForm; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Form_Refresh(AObj: TGoForm); stdcall;
begin
  AObj.Refresh;
end;

procedure Form_SendToBack(AObj: TGoForm); stdcall;
begin
  AObj.SendToBack;
end;

function Form_GetTextBuf(AObj: TGoForm; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Form_FindComponent(AObj: TGoForm; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Form_GetNamePath(AObj: TGoForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Form_Assign(AObj: TGoForm; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function Form_ClassName(AObj: TGoForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Form_Equals(AObj: TGoForm; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Form_GetHashCode(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Form_ToString(AObj: TGoForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Form_GetAction(AObj: TGoForm): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Form_SetAction(AObj: TGoForm; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function Form_GetAlign(AObj: TGoForm): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Form_SetAlign(AObj: TGoForm; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Form_GetAlphaBlend(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.AlphaBlend;
end;

procedure Form_SetAlphaBlend(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.AlphaBlend := AValue;
end;

function Form_GetAlphaBlendValue(AObj: TGoForm): Byte; stdcall;
begin
  Result :=  AObj.AlphaBlendValue;
end;

procedure Form_SetAlphaBlendValue(AObj: TGoForm; AValue: Byte); stdcall;
begin
  AObj.AlphaBlendValue := AValue;
end;

function Form_GetAnchors(AObj: TGoForm): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Form_SetAnchors(AObj: TGoForm; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Form_GetAutoSize(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure Form_SetAutoSize(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function Form_GetBiDiMode(AObj: TGoForm): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Form_SetBiDiMode(AObj: TGoForm; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function Form_GetBorderIcons(AObj: TGoForm): TBorderIcons; stdcall;
begin
  Result :=  AObj.BorderIcons;
end;

procedure Form_SetBorderIcons(AObj: TGoForm; AValue: TBorderIcons); stdcall;
begin
  AObj.BorderIcons := AValue;
end;

function Form_GetBorderStyle(AObj: TGoForm): TFormBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure Form_SetBorderStyle(AObj: TGoForm; AValue: TFormBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function Form_GetBorderWidth(AObj: TGoForm): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure Form_SetBorderWidth(AObj: TGoForm; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function Form_GetCaption(AObj: TGoForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure Form_SetCaption(AObj: TGoForm; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function Form_GetClientHeight(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Form_SetClientHeight(AObj: TGoForm; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

function Form_GetClientWidth(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Form_SetClientWidth(AObj: TGoForm; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Form_GetColor(AObj: TGoForm): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure Form_SetColor(AObj: TGoForm; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function Form_GetTransparentColor(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.TransparentColor;
end;

procedure Form_SetTransparentColor(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.TransparentColor := AValue;
end;

function Form_GetTransparentColorValue(AObj: TGoForm): TColor; stdcall;
begin
  Result :=  AObj.TransparentColorValue;
end;

procedure Form_SetTransparentColorValue(AObj: TGoForm; AValue: TColor); stdcall;
begin
  AObj.TransparentColorValue := AValue;
end;

function Form_GetDoubleBuffered(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure Form_SetDoubleBuffered(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function Form_GetEnabled(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Form_SetEnabled(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Form_GetParentFont(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure Form_SetParentFont(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function Form_GetFont(AObj: TGoForm): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure Form_SetFont(AObj: TGoForm; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function Form_GetFormStyle(AObj: TGoForm): TFormStyle; stdcall;
begin
  Result :=  AObj.FormStyle;
end;

procedure Form_SetFormStyle(AObj: TGoForm; AValue: TFormStyle); stdcall;
begin
  AObj.FormStyle := AValue;
end;

function Form_GetHeight(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Form_SetHeight(AObj: TGoForm; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Form_GetIcon(AObj: TGoForm): TIcon; stdcall;
begin
  Result :=  AObj.Icon;
end;

procedure Form_SetIcon(AObj: TGoForm; AValue: TIcon); stdcall;
begin
  AObj.Icon := AValue;
end;

function Form_GetKeyPreview(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.KeyPreview;
end;

procedure Form_SetKeyPreview(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.KeyPreview := AValue;
end;

function Form_GetMenu(AObj: TGoForm): TMainMenu; stdcall;
begin
  Result :=  AObj.Menu;
end;

procedure Form_SetMenu(AObj: TGoForm; AValue: TMainMenu); stdcall;
begin
  AObj.Menu := AValue;
end;

function Form_GetPixelsPerInch(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.PixelsPerInch;
end;

procedure Form_SetPixelsPerInch(AObj: TGoForm; AValue: Integer); stdcall;
begin
  AObj.PixelsPerInch := AValue;
end;

function Form_GetPopupMenu(AObj: TGoForm): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Form_SetPopupMenu(AObj: TGoForm; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function Form_GetPosition(AObj: TGoForm): TPosition; stdcall;
begin
  Result :=  AObj.Position;
end;

procedure Form_SetPosition(AObj: TGoForm; AValue: TPosition); stdcall;
begin
  AObj.Position := AValue;
end;

function Form_GetScaled(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.Scaled;
end;

procedure Form_SetScaled(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.Scaled := AValue;
end;

function Form_GetShowHint(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Form_SetShowHint(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Form_GetVisible(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Form_SetVisible(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Form_GetWidth(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Form_SetWidth(AObj: TGoForm; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Form_GetWindowState(AObj: TGoForm): TWindowState; stdcall;
begin
  Result :=  AObj.WindowState;
end;

procedure Form_SetWindowState(AObj: TGoForm; AValue: TWindowState); stdcall;
begin
  AObj.WindowState := AValue;
end;

function Form_GetStyleElements(AObj: TGoForm): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Form_SetStyleElements(AObj: TGoForm; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure Form_SetOnClick(AObj: TGoForm; AEventId: NativeUInt); stdcall
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

procedure Form_SetOnClose(AObj: TGoForm; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnClose := nil;
    TEventClass.Remove(AObj, geFormClose);
    Exit;
  end;
  AObj.OnClose := TEventClass.FormOnClose;
  TEventClass.Add(AObj, geFormClose, AEventId);
end;

procedure Form_SetOnCloseQuery(AObj: TGoForm; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnCloseQuery := nil;
    TEventClass.Remove(AObj, geFormCloseQuery);
    Exit;
  end;
  AObj.OnCloseQuery := TEventClass.FormOnCloseQuery;
  TEventClass.Add(AObj, geFormCloseQuery, AEventId);
end;

procedure Form_SetOnDblClick(AObj: TGoForm; AEventId: NativeUInt); stdcall
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

procedure Form_SetOnHide(AObj: TGoForm; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnHide := nil;
    TEventClass.Remove(AObj, geHide);
    Exit;
  end;
  AObj.OnHide := TEventClass.OnHide;
  TEventClass.Add(AObj, geHide, AEventId);
end;

procedure Form_SetOnKeyDown(AObj: TGoForm; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyDown := nil;
    TEventClass.Remove(AObj, geKeyDown);
    Exit;
  end;
  AObj.OnKeyDown := TEventClass.OnKeyDown;
  TEventClass.Add(AObj, geKeyDown, AEventId);
end;

procedure Form_SetOnKeyPress(AObj: TGoForm; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyPress := nil;
    TEventClass.Remove(AObj, geKeyPress);
    Exit;
  end;
  AObj.OnKeyPress := TEventClass.OnKeyPress;
  TEventClass.Add(AObj, geKeyPress, AEventId);
end;

procedure Form_SetOnKeyUp(AObj: TGoForm; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyUp := nil;
    TEventClass.Remove(AObj, geKeyUp);
    Exit;
  end;
  AObj.OnKeyUp := TEventClass.OnKeyUp;
  TEventClass.Add(AObj, geKeyUp, AEventId);
end;

procedure Form_SetOnMouseDown(AObj: TGoForm; AEventId: NativeUInt); stdcall
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

procedure Form_SetOnMouseEnter(AObj: TGoForm; AEventId: NativeUInt); stdcall
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

procedure Form_SetOnMouseLeave(AObj: TGoForm; AEventId: NativeUInt); stdcall
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

procedure Form_SetOnMouseMove(AObj: TGoForm; AEventId: NativeUInt); stdcall
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

procedure Form_SetOnMouseUp(AObj: TGoForm; AEventId: NativeUInt); stdcall
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

procedure Form_SetOnMouseWheel(AObj: TGoForm; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheel := nil;
    TEventClass.Remove(AObj, geMouseWheel);
    Exit;
  end;
  AObj.OnMouseWheel := TEventClass.OnMouseWheel;
  TEventClass.Add(AObj, geMouseWheel, AEventId);
end;

procedure Form_SetOnPaint(AObj: TGoForm; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnPaint := nil;
    TEventClass.Remove(AObj, gePaint);
    Exit;
  end;
  AObj.OnPaint := TEventClass.OnPaint;
  TEventClass.Add(AObj, gePaint, AEventId);
end;

procedure Form_SetOnResize(AObj: TGoForm; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnResize := nil;
    TEventClass.Remove(AObj, geResize);
    Exit;
  end;
  AObj.OnResize := TEventClass.OnResize;
  TEventClass.Add(AObj, geResize, AEventId);
end;

procedure Form_SetOnShow(AObj: TGoForm; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnShow := nil;
    TEventClass.Remove(AObj, geShow);
    Exit;
  end;
  AObj.OnShow := TEventClass.OnShow;
  TEventClass.Add(AObj, geShow, AEventId);
end;

function Form_GetCanvas(AObj: TGoForm): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function Form_GetDropTarget(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.DropTarget;
end;

procedure Form_SetDropTarget(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.DropTarget := AValue;
end;

function Form_GetModalResult(AObj: TGoForm): TModalResult; stdcall;
begin
  Result :=  AObj.ModalResult;
end;

procedure Form_SetModalResult(AObj: TGoForm; AValue: TModalResult); stdcall;
begin
  AObj.ModalResult := AValue;
end;

function Form_GetLeft(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Form_SetLeft(AObj: TGoForm; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Form_GetTop(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Form_SetTop(AObj: TGoForm; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Form_GetBrush(AObj: TGoForm): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function Form_GetControlCount(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function Form_GetHandle(AObj: TGoForm): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function Form_GetParentDoubleBuffered(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure Form_SetParentDoubleBuffered(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function Form_GetParentWindow(AObj: TGoForm): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure Form_SetParentWindow(AObj: TGoForm; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function Form_GetTabOrder(AObj: TGoForm): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure Form_SetTabOrder(AObj: TGoForm; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function Form_GetTabStop(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure Form_SetTabStop(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

procedure Form_GetBoundsRect(AObj: TGoForm; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Form_SetBoundsRect(AObj: TGoForm; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

procedure Form_GetClientRect(AObj: TGoForm; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Form_GetExplicitLeft(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Form_GetExplicitTop(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Form_GetExplicitWidth(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Form_GetExplicitHeight(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Form_GetParent(AObj: TGoForm): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Form_SetParent(AObj: TGoForm; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Form_GetAlignWithMargins(AObj: TGoForm): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Form_SetAlignWithMargins(AObj: TGoForm; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Form_GetCursor(AObj: TGoForm): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Form_SetCursor(AObj: TGoForm; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Form_GetHint(AObj: TGoForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Form_SetHint(AObj: TGoForm; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Form_GetMargins(AObj: TGoForm): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Form_SetMargins(AObj: TGoForm; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Form_GetComponentCount(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Form_GetComponentIndex(AObj: TGoForm): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Form_SetComponentIndex(AObj: TGoForm; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Form_GetOwner(AObj: TGoForm): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Form_GetName(AObj: TGoForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Form_SetName(AObj: TGoForm; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Form_GetTag(AObj: TGoForm): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Form_SetTag(AObj: TGoForm; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Form_GetControls(AObj: TGoForm; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function Form_GetComponents(AObj: TGoForm; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  Form_Create {$IFNDEF MSWINDOWS}name '_Form_Create'{$ENDIF},
  Form_Free {$IFNDEF MSWINDOWS}name '_Form_Free'{$ENDIF},
  Form_Close {$IFNDEF MSWINDOWS}name '_Form_Close'{$ENDIF},
  Form_Hide {$IFNDEF MSWINDOWS}name '_Form_Hide'{$ENDIF},
  Form_Print {$IFNDEF MSWINDOWS}name '_Form_Print'{$ENDIF},
  Form_SetFocus {$IFNDEF MSWINDOWS}name '_Form_SetFocus'{$ENDIF},
  Form_Show {$IFNDEF MSWINDOWS}name '_Form_Show'{$ENDIF},
  Form_ShowModal {$IFNDEF MSWINDOWS}name '_Form_ShowModal'{$ENDIF},
  Form_CanFocus {$IFNDEF MSWINDOWS}name '_Form_CanFocus'{$ENDIF},
  Form_FlipChildren {$IFNDEF MSWINDOWS}name '_Form_FlipChildren'{$ENDIF},
  Form_Focused {$IFNDEF MSWINDOWS}name '_Form_Focused'{$ENDIF},
  Form_HandleAllocated {$IFNDEF MSWINDOWS}name '_Form_HandleAllocated'{$ENDIF},
  Form_Invalidate {$IFNDEF MSWINDOWS}name '_Form_Invalidate'{$ENDIF},
  Form_Realign {$IFNDEF MSWINDOWS}name '_Form_Realign'{$ENDIF},
  Form_Repaint {$IFNDEF MSWINDOWS}name '_Form_Repaint'{$ENDIF},
  Form_ScaleBy {$IFNDEF MSWINDOWS}name '_Form_ScaleBy'{$ENDIF},
  Form_SetBounds {$IFNDEF MSWINDOWS}name '_Form_SetBounds'{$ENDIF},
  Form_Update {$IFNDEF MSWINDOWS}name '_Form_Update'{$ENDIF},
  Form_BringToFront {$IFNDEF MSWINDOWS}name '_Form_BringToFront'{$ENDIF},
  Form_HasParent {$IFNDEF MSWINDOWS}name '_Form_HasParent'{$ENDIF},
  Form_Perform {$IFNDEF MSWINDOWS}name '_Form_Perform'{$ENDIF},
  Form_Refresh {$IFNDEF MSWINDOWS}name '_Form_Refresh'{$ENDIF},
  Form_SendToBack {$IFNDEF MSWINDOWS}name '_Form_SendToBack'{$ENDIF},
  Form_GetTextBuf {$IFNDEF MSWINDOWS}name '_Form_GetTextBuf'{$ENDIF},
  Form_FindComponent {$IFNDEF MSWINDOWS}name '_Form_FindComponent'{$ENDIF},
  Form_GetNamePath {$IFNDEF MSWINDOWS}name '_Form_GetNamePath'{$ENDIF},
  Form_Assign {$IFNDEF MSWINDOWS}name '_Form_Assign'{$ENDIF},
  Form_ClassName {$IFNDEF MSWINDOWS}name '_Form_ClassName'{$ENDIF},
  Form_Equals {$IFNDEF MSWINDOWS}name '_Form_Equals'{$ENDIF},
  Form_GetHashCode {$IFNDEF MSWINDOWS}name '_Form_GetHashCode'{$ENDIF},
  Form_ToString {$IFNDEF MSWINDOWS}name '_Form_ToString'{$ENDIF},
  Form_GetAction {$IFNDEF MSWINDOWS}name '_Form_GetAction'{$ENDIF},
  Form_SetAction {$IFNDEF MSWINDOWS}name '_Form_SetAction'{$ENDIF},
  Form_GetAlign {$IFNDEF MSWINDOWS}name '_Form_GetAlign'{$ENDIF},
  Form_SetAlign {$IFNDEF MSWINDOWS}name '_Form_SetAlign'{$ENDIF},
  Form_GetAlphaBlend {$IFNDEF MSWINDOWS}name '_Form_GetAlphaBlend'{$ENDIF},
  Form_SetAlphaBlend {$IFNDEF MSWINDOWS}name '_Form_SetAlphaBlend'{$ENDIF},
  Form_GetAlphaBlendValue {$IFNDEF MSWINDOWS}name '_Form_GetAlphaBlendValue'{$ENDIF},
  Form_SetAlphaBlendValue {$IFNDEF MSWINDOWS}name '_Form_SetAlphaBlendValue'{$ENDIF},
  Form_GetAnchors {$IFNDEF MSWINDOWS}name '_Form_GetAnchors'{$ENDIF},
  Form_SetAnchors {$IFNDEF MSWINDOWS}name '_Form_SetAnchors'{$ENDIF},
  Form_GetAutoSize {$IFNDEF MSWINDOWS}name '_Form_GetAutoSize'{$ENDIF},
  Form_SetAutoSize {$IFNDEF MSWINDOWS}name '_Form_SetAutoSize'{$ENDIF},
  Form_GetBiDiMode {$IFNDEF MSWINDOWS}name '_Form_GetBiDiMode'{$ENDIF},
  Form_SetBiDiMode {$IFNDEF MSWINDOWS}name '_Form_SetBiDiMode'{$ENDIF},
  Form_GetBorderIcons {$IFNDEF MSWINDOWS}name '_Form_GetBorderIcons'{$ENDIF},
  Form_SetBorderIcons {$IFNDEF MSWINDOWS}name '_Form_SetBorderIcons'{$ENDIF},
  Form_GetBorderStyle {$IFNDEF MSWINDOWS}name '_Form_GetBorderStyle'{$ENDIF},
  Form_SetBorderStyle {$IFNDEF MSWINDOWS}name '_Form_SetBorderStyle'{$ENDIF},
  Form_GetBorderWidth {$IFNDEF MSWINDOWS}name '_Form_GetBorderWidth'{$ENDIF},
  Form_SetBorderWidth {$IFNDEF MSWINDOWS}name '_Form_SetBorderWidth'{$ENDIF},
  Form_GetCaption {$IFNDEF MSWINDOWS}name '_Form_GetCaption'{$ENDIF},
  Form_SetCaption {$IFNDEF MSWINDOWS}name '_Form_SetCaption'{$ENDIF},
  Form_GetClientHeight {$IFNDEF MSWINDOWS}name '_Form_GetClientHeight'{$ENDIF},
  Form_SetClientHeight {$IFNDEF MSWINDOWS}name '_Form_SetClientHeight'{$ENDIF},
  Form_GetClientWidth {$IFNDEF MSWINDOWS}name '_Form_GetClientWidth'{$ENDIF},
  Form_SetClientWidth {$IFNDEF MSWINDOWS}name '_Form_SetClientWidth'{$ENDIF},
  Form_GetColor {$IFNDEF MSWINDOWS}name '_Form_GetColor'{$ENDIF},
  Form_SetColor {$IFNDEF MSWINDOWS}name '_Form_SetColor'{$ENDIF},
  Form_GetTransparentColor {$IFNDEF MSWINDOWS}name '_Form_GetTransparentColor'{$ENDIF},
  Form_SetTransparentColor {$IFNDEF MSWINDOWS}name '_Form_SetTransparentColor'{$ENDIF},
  Form_GetTransparentColorValue {$IFNDEF MSWINDOWS}name '_Form_GetTransparentColorValue'{$ENDIF},
  Form_SetTransparentColorValue {$IFNDEF MSWINDOWS}name '_Form_SetTransparentColorValue'{$ENDIF},
  Form_GetDoubleBuffered {$IFNDEF MSWINDOWS}name '_Form_GetDoubleBuffered'{$ENDIF},
  Form_SetDoubleBuffered {$IFNDEF MSWINDOWS}name '_Form_SetDoubleBuffered'{$ENDIF},
  Form_GetEnabled {$IFNDEF MSWINDOWS}name '_Form_GetEnabled'{$ENDIF},
  Form_SetEnabled {$IFNDEF MSWINDOWS}name '_Form_SetEnabled'{$ENDIF},
  Form_GetParentFont {$IFNDEF MSWINDOWS}name '_Form_GetParentFont'{$ENDIF},
  Form_SetParentFont {$IFNDEF MSWINDOWS}name '_Form_SetParentFont'{$ENDIF},
  Form_GetFont {$IFNDEF MSWINDOWS}name '_Form_GetFont'{$ENDIF},
  Form_SetFont {$IFNDEF MSWINDOWS}name '_Form_SetFont'{$ENDIF},
  Form_GetFormStyle {$IFNDEF MSWINDOWS}name '_Form_GetFormStyle'{$ENDIF},
  Form_SetFormStyle {$IFNDEF MSWINDOWS}name '_Form_SetFormStyle'{$ENDIF},
  Form_GetHeight {$IFNDEF MSWINDOWS}name '_Form_GetHeight'{$ENDIF},
  Form_SetHeight {$IFNDEF MSWINDOWS}name '_Form_SetHeight'{$ENDIF},
  Form_GetIcon {$IFNDEF MSWINDOWS}name '_Form_GetIcon'{$ENDIF},
  Form_SetIcon {$IFNDEF MSWINDOWS}name '_Form_SetIcon'{$ENDIF},
  Form_GetKeyPreview {$IFNDEF MSWINDOWS}name '_Form_GetKeyPreview'{$ENDIF},
  Form_SetKeyPreview {$IFNDEF MSWINDOWS}name '_Form_SetKeyPreview'{$ENDIF},
  Form_GetMenu {$IFNDEF MSWINDOWS}name '_Form_GetMenu'{$ENDIF},
  Form_SetMenu {$IFNDEF MSWINDOWS}name '_Form_SetMenu'{$ENDIF},
  Form_GetPixelsPerInch {$IFNDEF MSWINDOWS}name '_Form_GetPixelsPerInch'{$ENDIF},
  Form_SetPixelsPerInch {$IFNDEF MSWINDOWS}name '_Form_SetPixelsPerInch'{$ENDIF},
  Form_GetPopupMenu {$IFNDEF MSWINDOWS}name '_Form_GetPopupMenu'{$ENDIF},
  Form_SetPopupMenu {$IFNDEF MSWINDOWS}name '_Form_SetPopupMenu'{$ENDIF},
  Form_GetPosition {$IFNDEF MSWINDOWS}name '_Form_GetPosition'{$ENDIF},
  Form_SetPosition {$IFNDEF MSWINDOWS}name '_Form_SetPosition'{$ENDIF},
  Form_GetScaled {$IFNDEF MSWINDOWS}name '_Form_GetScaled'{$ENDIF},
  Form_SetScaled {$IFNDEF MSWINDOWS}name '_Form_SetScaled'{$ENDIF},
  Form_GetShowHint {$IFNDEF MSWINDOWS}name '_Form_GetShowHint'{$ENDIF},
  Form_SetShowHint {$IFNDEF MSWINDOWS}name '_Form_SetShowHint'{$ENDIF},
  Form_GetVisible {$IFNDEF MSWINDOWS}name '_Form_GetVisible'{$ENDIF},
  Form_SetVisible {$IFNDEF MSWINDOWS}name '_Form_SetVisible'{$ENDIF},
  Form_GetWidth {$IFNDEF MSWINDOWS}name '_Form_GetWidth'{$ENDIF},
  Form_SetWidth {$IFNDEF MSWINDOWS}name '_Form_SetWidth'{$ENDIF},
  Form_GetWindowState {$IFNDEF MSWINDOWS}name '_Form_GetWindowState'{$ENDIF},
  Form_SetWindowState {$IFNDEF MSWINDOWS}name '_Form_SetWindowState'{$ENDIF},
  Form_GetStyleElements {$IFNDEF MSWINDOWS}name '_Form_GetStyleElements'{$ENDIF},
  Form_SetStyleElements {$IFNDEF MSWINDOWS}name '_Form_SetStyleElements'{$ENDIF},
  Form_SetOnClick {$IFNDEF MSWINDOWS}name '_Form_SetOnClick'{$ENDIF},
  Form_SetOnClose {$IFNDEF MSWINDOWS}name '_Form_SetOnClose'{$ENDIF},
  Form_SetOnCloseQuery {$IFNDEF MSWINDOWS}name '_Form_SetOnCloseQuery'{$ENDIF},
  Form_SetOnDblClick {$IFNDEF MSWINDOWS}name '_Form_SetOnDblClick'{$ENDIF},
  Form_SetOnHide {$IFNDEF MSWINDOWS}name '_Form_SetOnHide'{$ENDIF},
  Form_SetOnKeyDown {$IFNDEF MSWINDOWS}name '_Form_SetOnKeyDown'{$ENDIF},
  Form_SetOnKeyPress {$IFNDEF MSWINDOWS}name '_Form_SetOnKeyPress'{$ENDIF},
  Form_SetOnKeyUp {$IFNDEF MSWINDOWS}name '_Form_SetOnKeyUp'{$ENDIF},
  Form_SetOnMouseDown {$IFNDEF MSWINDOWS}name '_Form_SetOnMouseDown'{$ENDIF},
  Form_SetOnMouseEnter {$IFNDEF MSWINDOWS}name '_Form_SetOnMouseEnter'{$ENDIF},
  Form_SetOnMouseLeave {$IFNDEF MSWINDOWS}name '_Form_SetOnMouseLeave'{$ENDIF},
  Form_SetOnMouseMove {$IFNDEF MSWINDOWS}name '_Form_SetOnMouseMove'{$ENDIF},
  Form_SetOnMouseUp {$IFNDEF MSWINDOWS}name '_Form_SetOnMouseUp'{$ENDIF},
  Form_SetOnMouseWheel {$IFNDEF MSWINDOWS}name '_Form_SetOnMouseWheel'{$ENDIF},
  Form_SetOnPaint {$IFNDEF MSWINDOWS}name '_Form_SetOnPaint'{$ENDIF},
  Form_SetOnResize {$IFNDEF MSWINDOWS}name '_Form_SetOnResize'{$ENDIF},
  Form_SetOnShow {$IFNDEF MSWINDOWS}name '_Form_SetOnShow'{$ENDIF},
  Form_GetCanvas {$IFNDEF MSWINDOWS}name '_Form_GetCanvas'{$ENDIF},
  Form_GetDropTarget {$IFNDEF MSWINDOWS}name '_Form_GetDropTarget'{$ENDIF},
  Form_SetDropTarget {$IFNDEF MSWINDOWS}name '_Form_SetDropTarget'{$ENDIF},
  Form_GetModalResult {$IFNDEF MSWINDOWS}name '_Form_GetModalResult'{$ENDIF},
  Form_SetModalResult {$IFNDEF MSWINDOWS}name '_Form_SetModalResult'{$ENDIF},
  Form_GetLeft {$IFNDEF MSWINDOWS}name '_Form_GetLeft'{$ENDIF},
  Form_SetLeft {$IFNDEF MSWINDOWS}name '_Form_SetLeft'{$ENDIF},
  Form_GetTop {$IFNDEF MSWINDOWS}name '_Form_GetTop'{$ENDIF},
  Form_SetTop {$IFNDEF MSWINDOWS}name '_Form_SetTop'{$ENDIF},
  Form_GetBrush {$IFNDEF MSWINDOWS}name '_Form_GetBrush'{$ENDIF},
  Form_GetControlCount {$IFNDEF MSWINDOWS}name '_Form_GetControlCount'{$ENDIF},
  Form_GetHandle {$IFNDEF MSWINDOWS}name '_Form_GetHandle'{$ENDIF},
  Form_GetParentDoubleBuffered {$IFNDEF MSWINDOWS}name '_Form_GetParentDoubleBuffered'{$ENDIF},
  Form_SetParentDoubleBuffered {$IFNDEF MSWINDOWS}name '_Form_SetParentDoubleBuffered'{$ENDIF},
  Form_GetParentWindow {$IFNDEF MSWINDOWS}name '_Form_GetParentWindow'{$ENDIF},
  Form_SetParentWindow {$IFNDEF MSWINDOWS}name '_Form_SetParentWindow'{$ENDIF},
  Form_GetTabOrder {$IFNDEF MSWINDOWS}name '_Form_GetTabOrder'{$ENDIF},
  Form_SetTabOrder {$IFNDEF MSWINDOWS}name '_Form_SetTabOrder'{$ENDIF},
  Form_GetTabStop {$IFNDEF MSWINDOWS}name '_Form_GetTabStop'{$ENDIF},
  Form_SetTabStop {$IFNDEF MSWINDOWS}name '_Form_SetTabStop'{$ENDIF},
  Form_GetBoundsRect {$IFNDEF MSWINDOWS}name '_Form_GetBoundsRect'{$ENDIF},
  Form_SetBoundsRect {$IFNDEF MSWINDOWS}name '_Form_SetBoundsRect'{$ENDIF},
  Form_GetClientRect {$IFNDEF MSWINDOWS}name '_Form_GetClientRect'{$ENDIF},
  Form_GetExplicitLeft {$IFNDEF MSWINDOWS}name '_Form_GetExplicitLeft'{$ENDIF},
  Form_GetExplicitTop {$IFNDEF MSWINDOWS}name '_Form_GetExplicitTop'{$ENDIF},
  Form_GetExplicitWidth {$IFNDEF MSWINDOWS}name '_Form_GetExplicitWidth'{$ENDIF},
  Form_GetExplicitHeight {$IFNDEF MSWINDOWS}name '_Form_GetExplicitHeight'{$ENDIF},
  Form_GetParent {$IFNDEF MSWINDOWS}name '_Form_GetParent'{$ENDIF},
  Form_SetParent {$IFNDEF MSWINDOWS}name '_Form_SetParent'{$ENDIF},
  Form_GetAlignWithMargins {$IFNDEF MSWINDOWS}name '_Form_GetAlignWithMargins'{$ENDIF},
  Form_SetAlignWithMargins {$IFNDEF MSWINDOWS}name '_Form_SetAlignWithMargins'{$ENDIF},
  Form_GetCursor {$IFNDEF MSWINDOWS}name '_Form_GetCursor'{$ENDIF},
  Form_SetCursor {$IFNDEF MSWINDOWS}name '_Form_SetCursor'{$ENDIF},
  Form_GetHint {$IFNDEF MSWINDOWS}name '_Form_GetHint'{$ENDIF},
  Form_SetHint {$IFNDEF MSWINDOWS}name '_Form_SetHint'{$ENDIF},
  Form_GetMargins {$IFNDEF MSWINDOWS}name '_Form_GetMargins'{$ENDIF},
  Form_SetMargins {$IFNDEF MSWINDOWS}name '_Form_SetMargins'{$ENDIF},
  Form_GetComponentCount {$IFNDEF MSWINDOWS}name '_Form_GetComponentCount'{$ENDIF},
  Form_GetComponentIndex {$IFNDEF MSWINDOWS}name '_Form_GetComponentIndex'{$ENDIF},
  Form_SetComponentIndex {$IFNDEF MSWINDOWS}name '_Form_SetComponentIndex'{$ENDIF},
  Form_GetOwner {$IFNDEF MSWINDOWS}name '_Form_GetOwner'{$ENDIF},
  Form_GetName {$IFNDEF MSWINDOWS}name '_Form_GetName'{$ENDIF},
  Form_SetName {$IFNDEF MSWINDOWS}name '_Form_SetName'{$ENDIF},
  Form_GetTag {$IFNDEF MSWINDOWS}name '_Form_GetTag'{$ENDIF},
  Form_SetTag {$IFNDEF MSWINDOWS}name '_Form_SetTag'{$ENDIF},
  Form_GetControls {$IFNDEF MSWINDOWS}name '_Form_GetControls'{$ENDIF},
  Form_GetComponents {$IFNDEF MSWINDOWS}name '_Form_GetComponents'{$ENDIF};

