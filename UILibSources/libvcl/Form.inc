
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function Form_Create(AOwner: TComponent): TForm; stdcall;
begin
  Result :=  TGoForm.CreateNew(AOwner);
end;

procedure Form_Free(AObj: TForm); stdcall;
begin
  AObj.Free;
end;

procedure Form_Close(AObj: TForm); stdcall;
begin
  AObj.Close;
end;

procedure Form_Hide(AObj: TForm); stdcall;
begin
  AObj.Hide;
end;

procedure Form_Print(AObj: TForm); stdcall;
begin
  AObj.Print;
end;

procedure Form_SetFocus(AObj: TForm); stdcall;
begin
  AObj.SetFocus;
end;

procedure Form_Show(AObj: TForm); stdcall;
begin
  AObj.Show;
end;

function Form_ShowModal(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.ShowModal;
end;

function Form_CanFocus(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure Form_FlipChildren(AObj: TForm; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function Form_Focused(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function Form_HandleAllocated(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure Form_Invalidate(AObj: TForm); stdcall;
begin
  AObj.Invalidate;
end;

procedure Form_Realign(AObj: TForm); stdcall;
begin
  AObj.Realign;
end;

procedure Form_Repaint(AObj: TForm); stdcall;
begin
  AObj.Repaint;
end;

procedure Form_ScaleBy(AObj: TForm; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure Form_SetBounds(AObj: TForm; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure Form_Update(AObj: TForm); stdcall;
begin
  AObj.Update;
end;

procedure Form_BringToFront(AObj: TForm); stdcall;
begin
  AObj.BringToFront;
end;

procedure Form_ClientToScreen(AObj: TForm; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToScreen(Point);
end;

procedure Form_ClientToParent(AObj: TForm; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ClientToParent(Point, AParent);
end;

function Form_Dragging(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.Dragging;
end;

function Form_HasParent(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

function Form_Perform(AObj: TForm; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure Form_Refresh(AObj: TForm); stdcall;
begin
  AObj.Refresh;
end;

procedure Form_ScreenToClient(AObj: TForm; var Point: TPoint; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ScreenToClient(Point);
end;

procedure Form_ParentToClient(AObj: TForm; var Point: TPoint; AParent: TWinControl; var Result: TPoint); stdcall;
begin
  Result :=  AObj.ParentToClient(Point, AParent);
end;

procedure Form_SendToBack(AObj: TForm); stdcall;
begin
  AObj.SendToBack;
end;

function Form_GetTextBuf(AObj: TForm; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function Form_GetTextLen(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.GetTextLen;
end;

function Form_FindComponent(AObj: TForm; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function Form_GetNamePath(AObj: TForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure Form_Assign(AObj: TForm; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

procedure Form_DisposeOf(AObj: TForm); stdcall;
begin
  AObj.DisposeOf;
end;

function Form_ClassType(AObj: TForm): TClass; stdcall;
begin
  Result :=  AObj.ClassType;
end;

function Form_ClassName(AObj: TForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function Form_InstanceSize(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.InstanceSize;
end;

function Form_InheritsFrom(AObj: TForm; AClass: TClass): LongBool; stdcall;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function Form_Equals(AObj: TForm; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function Form_GetHashCode(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function Form_ToString(AObj: TForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function Form_GetAction(AObj: TForm): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure Form_SetAction(AObj: TForm; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function Form_GetAlign(AObj: TForm): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure Form_SetAlign(AObj: TForm; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function Form_GetAlphaBlend(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.AlphaBlend;
end;

procedure Form_SetAlphaBlend(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.AlphaBlend := AValue;
end;

function Form_GetAlphaBlendValue(AObj: TForm): Byte; stdcall;
begin
  Result :=  AObj.AlphaBlendValue;
end;

procedure Form_SetAlphaBlendValue(AObj: TForm; AValue: Byte); stdcall;
begin
  AObj.AlphaBlendValue := AValue;
end;

function Form_GetAnchors(AObj: TForm): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure Form_SetAnchors(AObj: TForm; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function Form_GetAutoSize(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.AutoSize;
end;

procedure Form_SetAutoSize(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.AutoSize := AValue;
end;

function Form_GetBiDiMode(AObj: TForm): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure Form_SetBiDiMode(AObj: TForm; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

function Form_GetBorderIcons(AObj: TForm): TBorderIcons; stdcall;
begin
  Result :=  AObj.BorderIcons;
end;

procedure Form_SetBorderIcons(AObj: TForm; AValue: TBorderIcons); stdcall;
begin
  AObj.BorderIcons := AValue;
end;

function Form_GetBorderStyle(AObj: TForm): TFormBorderStyle; stdcall;
begin
  Result :=  AObj.BorderStyle;
end;

procedure Form_SetBorderStyle(AObj: TForm; AValue: TFormBorderStyle); stdcall;
begin
  AObj.BorderStyle := AValue;
end;

function Form_GetBorderWidth(AObj: TForm): TBorderWidth; stdcall;
begin
  Result :=  AObj.BorderWidth;
end;

procedure Form_SetBorderWidth(AObj: TForm; AValue: TBorderWidth); stdcall;
begin
  AObj.BorderWidth := AValue;
end;

function Form_GetCaption(AObj: TForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.Caption);
end;

procedure Form_SetCaption(AObj: TForm; AValue: PChar); stdcall;
begin
  AObj.Caption := AValue;
end;

function Form_GetClientHeight(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure Form_SetClientHeight(AObj: TForm; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

function Form_GetClientWidth(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure Form_SetClientWidth(AObj: TForm; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function Form_GetColor(AObj: TForm): TColor; stdcall;
begin
  Result :=  AObj.Color;
end;

procedure Form_SetColor(AObj: TForm; AValue: TColor); stdcall;
begin
  AObj.Color := AValue;
end;

function Form_GetTransparentColor(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.TransparentColor;
end;

procedure Form_SetTransparentColor(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.TransparentColor := AValue;
end;

function Form_GetTransparentColorValue(AObj: TForm): TColor; stdcall;
begin
  Result :=  AObj.TransparentColorValue;
end;

procedure Form_SetTransparentColorValue(AObj: TForm; AValue: TColor); stdcall;
begin
  AObj.TransparentColorValue := AValue;
end;

function Form_GetUseDockManager(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.UseDockManager;
end;

procedure Form_SetUseDockManager(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.UseDockManager := AValue;
end;

function Form_GetDockSite(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.DockSite;
end;

procedure Form_SetDockSite(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.DockSite := AValue;
end;

function Form_GetDoubleBuffered(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure Form_SetDoubleBuffered(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function Form_GetDragKind(AObj: TForm): TDragKind; stdcall;
begin
  Result :=  AObj.DragKind;
end;

procedure Form_SetDragKind(AObj: TForm; AValue: TDragKind); stdcall;
begin
  AObj.DragKind := AValue;
end;

function Form_GetDragMode(AObj: TForm): TDragMode; stdcall;
begin
  Result :=  AObj.DragMode;
end;

procedure Form_SetDragMode(AObj: TForm; AValue: TDragMode); stdcall;
begin
  AObj.DragMode := AValue;
end;

function Form_GetEnabled(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure Form_SetEnabled(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function Form_GetParentFont(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.ParentFont;
end;

procedure Form_SetParentFont(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.ParentFont := AValue;
end;

function Form_GetFont(AObj: TForm): TFont; stdcall;
begin
  Result :=  AObj.Font;
end;

procedure Form_SetFont(AObj: TForm; AValue: TFont); stdcall;
begin
  AObj.Font := AValue;
end;

function Form_GetFormStyle(AObj: TForm): TFormStyle; stdcall;
begin
  Result :=  AObj.FormStyle;
end;

procedure Form_SetFormStyle(AObj: TForm; AValue: TFormStyle); stdcall;
begin
  AObj.FormStyle := AValue;
end;

function Form_GetHeight(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure Form_SetHeight(AObj: TForm; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function Form_GetIcon(AObj: TForm): TIcon; stdcall;
begin
  Result :=  AObj.Icon;
end;

procedure Form_SetIcon(AObj: TForm; AValue: TIcon); stdcall;
begin
  AObj.Icon := AValue;
end;

function Form_GetKeyPreview(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.KeyPreview;
end;

procedure Form_SetKeyPreview(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.KeyPreview := AValue;
end;

function Form_GetMenu(AObj: TForm): TMainMenu; stdcall;
begin
  Result :=  AObj.Menu;
end;

procedure Form_SetMenu(AObj: TForm; AValue: TMainMenu); stdcall;
begin
  AObj.Menu := AValue;
end;

function Form_GetPixelsPerInch(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.PixelsPerInch;
end;

procedure Form_SetPixelsPerInch(AObj: TForm; AValue: Integer); stdcall;
begin
  AObj.PixelsPerInch := AValue;
end;

function Form_GetPopupMenu(AObj: TForm): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure Form_SetPopupMenu(AObj: TForm; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function Form_GetPosition(AObj: TForm): TPosition; stdcall;
begin
  Result :=  AObj.Position;
end;

procedure Form_SetPosition(AObj: TForm; AValue: TPosition); stdcall;
begin
  AObj.Position := AValue;
end;

function Form_GetScaled(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.Scaled;
end;

procedure Form_SetScaled(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.Scaled := AValue;
end;

function Form_GetShowHint(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure Form_SetShowHint(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function Form_GetVisible(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure Form_SetVisible(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function Form_GetWidth(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure Form_SetWidth(AObj: TForm; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function Form_GetWindowState(AObj: TForm): TWindowState; stdcall;
begin
  Result :=  AObj.WindowState;
end;

procedure Form_SetWindowState(AObj: TForm; AValue: TWindowState); stdcall;
begin
  AObj.WindowState := AValue;
end;

function Form_GetStyleElements(AObj: TForm): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure Form_SetStyleElements(AObj: TForm; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure Form_SetOnClick(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnClose(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnCloseQuery(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnContextPopup(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnContextPopup := nil;
    TEventClass.Remove(AObj, geContextPopup);
    Exit;
  end;
  AObj.OnContextPopup := TEventClass.OnContextPopup;
  TEventClass.Add(AObj, geContextPopup, AEventId);
end;

procedure Form_SetOnDblClick(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnDockDrop(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDockDrop := nil;
    TEventClass.Remove(AObj, geDockDrop);
    Exit;
  end;
  AObj.OnDockDrop := TEventClass.OnDockDrop;
  TEventClass.Add(AObj, geDockDrop, AEventId);
end;

procedure Form_SetOnDragDrop(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragDrop := nil;
    TEventClass.Remove(AObj, geDragDrop);
    Exit;
  end;
  AObj.OnDragDrop := TEventClass.OnDragDrop;
  TEventClass.Add(AObj, geDragDrop, AEventId);
end;

procedure Form_SetOnDragOver(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnDragOver := nil;
    TEventClass.Remove(AObj, geDragOver);
    Exit;
  end;
  AObj.OnDragOver := TEventClass.OnDragOver;
  TEventClass.Add(AObj, geDragOver, AEventId);
end;

procedure Form_SetOnEndDock(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnEndDock := nil;
    TEventClass.Remove(AObj, geEndDock);
    Exit;
  end;
  AObj.OnEndDock := TEventClass.OnEndDock;
  TEventClass.Add(AObj, geEndDock, AEventId);
end;

procedure Form_SetOnGetSiteInfo(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnGetSiteInfo := nil;
    TEventClass.Remove(AObj, geGetSiteInfo);
    Exit;
  end;
  AObj.OnGetSiteInfo := TEventClass.OnGetSiteInfo;
  TEventClass.Add(AObj, geGetSiteInfo, AEventId);
end;

procedure Form_SetOnHide(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnHelp(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnHelp := nil;
    TEventClass.Remove(AObj, geHelp);
    Exit;
  end;
  AObj.OnHelp := TEventClass.OnHelp;
  TEventClass.Add(AObj, geHelp, AEventId);
end;

procedure Form_SetOnKeyDown(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnKeyPress(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnKeyUp(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnMouseDown(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnMouseEnter(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnMouseLeave(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnMouseMove(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnMouseUp(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnMouseWheel(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnMouseWheelDown(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheelDown := nil;
    TEventClass.Remove(AObj, geMouseWheelDown);
    Exit;
  end;
  AObj.OnMouseWheelDown := TEventClass.OnMouseWheelDown;
  TEventClass.Add(AObj, geMouseWheelDown, AEventId);
end;

procedure Form_SetOnMouseWheelUp(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseWheelUp := nil;
    TEventClass.Remove(AObj, geMouseWheelUp);
    Exit;
  end;
  AObj.OnMouseWheelUp := TEventClass.OnMouseWheelUp;
  TEventClass.Add(AObj, geMouseWheelUp, AEventId);
end;

procedure Form_SetOnPaint(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnResize(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnShortCut(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnShortCut := nil;
    TEventClass.Remove(AObj, geShortCut);
    Exit;
  end;
  AObj.OnShortCut := TEventClass.OnShortCut;
  TEventClass.Add(AObj, geShortCut, AEventId);
end;

procedure Form_SetOnShow(AObj: TForm; AEventId: NativeUInt); stdcall;
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

procedure Form_SetOnStartDock(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnStartDock := nil;
    TEventClass.Remove(AObj, geStartDock);
    Exit;
  end;
  AObj.OnStartDock := TEventClass.OnStartDock;
  TEventClass.Add(AObj, geStartDock, AEventId);
end;

procedure Form_SetOnUnDock(AObj: TForm; AEventId: NativeUInt); stdcall;
begin
  if AEventId = 0 then
  begin
    AObj.OnUnDock := nil;
    TEventClass.Remove(AObj, geUnDock);
    Exit;
  end;
  AObj.OnUnDock := TEventClass.OnUnDock;
  TEventClass.Add(AObj, geUnDock, AEventId);
end;

function Form_GetCanvas(AObj: TForm): TCanvas; stdcall;
begin
  Result :=  AObj.Canvas;
end;

function Form_GetDropTarget(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.DropTarget;
end;

procedure Form_SetDropTarget(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.DropTarget := AValue;
end;

function Form_GetModalResult(AObj: TForm): TModalResult; stdcall;
begin
  Result :=  AObj.ModalResult;
end;

procedure Form_SetModalResult(AObj: TForm; AValue: TModalResult); stdcall;
begin
  AObj.ModalResult := AValue;
end;

function Form_GetLeft(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure Form_SetLeft(AObj: TForm; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function Form_GetTop(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure Form_SetTop(AObj: TForm; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function Form_GetBrush(AObj: TForm): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function Form_GetControlCount(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function Form_GetHandle(AObj: TForm): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function Form_GetParentDoubleBuffered(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure Form_SetParentDoubleBuffered(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function Form_GetParentWindow(AObj: TForm): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure Form_SetParentWindow(AObj: TForm; AValue: HWND); stdcall;
begin
  AObj.ParentWindow := AValue;
end;

function Form_GetTabOrder(AObj: TForm): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure Form_SetTabOrder(AObj: TForm; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function Form_GetTabStop(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure Form_SetTabStop(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

procedure Form_GetBoundsRect(AObj: TForm; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure Form_SetBoundsRect(AObj: TForm; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

procedure Form_GetClientRect(AObj: TForm; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function Form_GetExplicitLeft(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function Form_GetExplicitTop(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function Form_GetExplicitWidth(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function Form_GetExplicitHeight(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function Form_GetFloating(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.Floating;
end;

function Form_GetParent(AObj: TForm): TWinControl; stdcall;
begin
  Result :=  AObj.Parent;
end;

procedure Form_SetParent(AObj: TForm; AValue: TWinControl); stdcall;
begin
  AObj.Parent := AValue;
end;

function Form_GetAlignWithMargins(AObj: TForm): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure Form_SetAlignWithMargins(AObj: TForm; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function Form_GetCursor(AObj: TForm): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure Form_SetCursor(AObj: TForm; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function Form_GetHint(AObj: TForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure Form_SetHint(AObj: TForm; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function Form_GetMargins(AObj: TForm): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure Form_SetMargins(AObj: TForm; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function Form_GetCustomHint(AObj: TForm): TCustomHint; stdcall;
begin
  Result :=  AObj.CustomHint;
end;

procedure Form_SetCustomHint(AObj: TForm; AValue: TCustomHint); stdcall;
begin
  AObj.CustomHint := AValue;
end;

function Form_GetComponentCount(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function Form_GetComponentIndex(AObj: TForm): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure Form_SetComponentIndex(AObj: TForm; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function Form_GetOwner(AObj: TForm): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function Form_GetName(AObj: TForm): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

procedure Form_SetName(AObj: TForm; AValue: PChar); stdcall;
begin
  AObj.Name := AValue;
end;

function Form_GetTag(AObj: TForm): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure Form_SetTag(AObj: TForm; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function Form_GetControls(AObj: TForm; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function Form_GetComponents(AObj: TForm; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

function Form_StaticClassType: TClass; stdcall;
begin
  Result := TForm;
end;

exports
  Form_Create,
  Form_Free,
  Form_Close,
  Form_Hide,
  Form_Print,
  Form_SetFocus,
  Form_Show,
  Form_ShowModal,
  Form_CanFocus,
  Form_FlipChildren,
  Form_Focused,
  Form_HandleAllocated,
  Form_Invalidate,
  Form_Realign,
  Form_Repaint,
  Form_ScaleBy,
  Form_SetBounds,
  Form_Update,
  Form_BringToFront,
  Form_ClientToScreen,
  Form_ClientToParent,
  Form_Dragging,
  Form_HasParent,
  Form_Perform,
  Form_Refresh,
  Form_ScreenToClient,
  Form_ParentToClient,
  Form_SendToBack,
  Form_GetTextBuf,
  Form_GetTextLen,
  Form_FindComponent,
  Form_GetNamePath,
  Form_Assign,
  Form_DisposeOf,
  Form_ClassType,
  Form_ClassName,
  Form_InstanceSize,
  Form_InheritsFrom,
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
  Form_GetTransparentColor,
  Form_SetTransparentColor,
  Form_GetTransparentColorValue,
  Form_SetTransparentColorValue,
  Form_GetUseDockManager,
  Form_SetUseDockManager,
  Form_GetDockSite,
  Form_SetDockSite,
  Form_GetDoubleBuffered,
  Form_SetDoubleBuffered,
  Form_GetDragKind,
  Form_SetDragKind,
  Form_GetDragMode,
  Form_SetDragMode,
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
  Form_GetScaled,
  Form_SetScaled,
  Form_GetShowHint,
  Form_SetShowHint,
  Form_GetVisible,
  Form_SetVisible,
  Form_GetWidth,
  Form_SetWidth,
  Form_GetWindowState,
  Form_SetWindowState,
  Form_GetStyleElements,
  Form_SetStyleElements,
  Form_SetOnClick,
  Form_SetOnClose,
  Form_SetOnCloseQuery,
  Form_SetOnContextPopup,
  Form_SetOnDblClick,
  Form_SetOnDockDrop,
  Form_SetOnDragDrop,
  Form_SetOnDragOver,
  Form_SetOnEndDock,
  Form_SetOnGetSiteInfo,
  Form_SetOnHide,
  Form_SetOnHelp,
  Form_SetOnKeyDown,
  Form_SetOnKeyPress,
  Form_SetOnKeyUp,
  Form_SetOnMouseDown,
  Form_SetOnMouseEnter,
  Form_SetOnMouseLeave,
  Form_SetOnMouseMove,
  Form_SetOnMouseUp,
  Form_SetOnMouseWheel,
  Form_SetOnMouseWheelDown,
  Form_SetOnMouseWheelUp,
  Form_SetOnPaint,
  Form_SetOnResize,
  Form_SetOnShortCut,
  Form_SetOnShow,
  Form_SetOnStartDock,
  Form_SetOnUnDock,
  Form_GetCanvas,
  Form_GetDropTarget,
  Form_SetDropTarget,
  Form_GetModalResult,
  Form_SetModalResult,
  Form_GetLeft,
  Form_SetLeft,
  Form_GetTop,
  Form_SetTop,
  Form_GetBrush,
  Form_GetControlCount,
  Form_GetHandle,
  Form_GetParentDoubleBuffered,
  Form_SetParentDoubleBuffered,
  Form_GetParentWindow,
  Form_SetParentWindow,
  Form_GetTabOrder,
  Form_SetTabOrder,
  Form_GetTabStop,
  Form_SetTabStop,
  Form_GetBoundsRect,
  Form_SetBoundsRect,
  Form_GetClientRect,
  Form_GetExplicitLeft,
  Form_GetExplicitTop,
  Form_GetExplicitWidth,
  Form_GetExplicitHeight,
  Form_GetFloating,
  Form_GetParent,
  Form_SetParent,
  Form_GetAlignWithMargins,
  Form_SetAlignWithMargins,
  Form_GetCursor,
  Form_SetCursor,
  Form_GetHint,
  Form_SetHint,
  Form_GetMargins,
  Form_SetMargins,
  Form_GetCustomHint,
  Form_SetCustomHint,
  Form_GetComponentCount,
  Form_GetComponentIndex,
  Form_SetComponentIndex,
  Form_GetOwner,
  Form_GetName,
  Form_SetName,
  Form_GetTag,
  Form_SetTag,
  Form_GetControls,
  Form_GetComponents,
  Form_StaticClassType;

