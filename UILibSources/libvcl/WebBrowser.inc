
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function WebBrowser_Create(AOwner: TComponent): TWebBrowser; stdcall;
begin
//  CoInitializeEx(nil, 0);
  Result :=  TWebBrowser.Create(AOwner);
end;

procedure WebBrowser_Free(AObj: TWebBrowser); stdcall;
begin
  AObj.Free;
//  CoUninitialize;
end;

procedure WebBrowser_GoBack(AObj: TWebBrowser); stdcall;
begin
  AObj.GoBack;
end;

procedure WebBrowser_GoForward(AObj: TWebBrowser); stdcall;
begin
  AObj.GoForward;
end;

procedure WebBrowser_GoHome(AObj: TWebBrowser); stdcall;
begin
  AObj.GoHome;
end;

procedure WebBrowser_GoSearch(AObj: TWebBrowser); stdcall;
begin
  AObj.GoSearch;
end;

procedure WebBrowser_Navigate(AObj: TWebBrowser; URL: PChar); stdcall;
begin
  AObj.Navigate(URL);
end;

procedure WebBrowser_Refresh(AObj: TWebBrowser); stdcall;
begin
  AObj.Refresh;
end;

procedure WebBrowser_Refresh2(AObj: TWebBrowser); stdcall;
begin
  AObj.Refresh2;
end;

procedure WebBrowser_Stop(AObj: TWebBrowser); stdcall;
begin
  AObj.Stop;
end;

procedure WebBrowser_Quit(AObj: TWebBrowser); stdcall;
begin
  AObj.Quit;
end;

procedure WebBrowser_SetBounds(AObj: TWebBrowser; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

function WebBrowser_CanFocus(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

procedure WebBrowser_FlipChildren(AObj: TWebBrowser; AllLevels: LongBool); stdcall;
begin
  AObj.FlipChildren(AllLevels);
end;

function WebBrowser_Focused(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.Focused;
end;

function WebBrowser_HandleAllocated(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure WebBrowser_Invalidate(AObj: TWebBrowser); stdcall;
begin
  AObj.Invalidate;
end;

procedure WebBrowser_Realign(AObj: TWebBrowser); stdcall;
begin
  AObj.Realign;
end;

procedure WebBrowser_Repaint(AObj: TWebBrowser); stdcall;
begin
  AObj.Repaint;
end;

procedure WebBrowser_ScaleBy(AObj: TWebBrowser; M: Integer; D: Integer); stdcall;
begin
  AObj.ScaleBy(M, D);
end;

procedure WebBrowser_SetFocus(AObj: TWebBrowser); stdcall;
begin
  AObj.SetFocus;
end;

procedure WebBrowser_Update(AObj: TWebBrowser); stdcall;
begin
  AObj.Update;
end;

procedure WebBrowser_BringToFront(AObj: TWebBrowser); stdcall;
begin
  AObj.BringToFront;
end;

function WebBrowser_HasParent(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.HasParent;
end;

procedure WebBrowser_Hide(AObj: TWebBrowser); stdcall;
begin
  AObj.Hide;
end;

function WebBrowser_Perform(AObj: TWebBrowser; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; stdcall;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure WebBrowser_SendToBack(AObj: TWebBrowser); stdcall;
begin
  AObj.SendToBack;
end;

procedure WebBrowser_Show(AObj: TWebBrowser); stdcall;
begin
  AObj.Show;
end;

function WebBrowser_GetTextBuf(AObj: TWebBrowser; Buffer: PWideChar; BufSize: Integer): Integer; stdcall;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function WebBrowser_FindComponent(AObj: TWebBrowser; AName: PChar): TComponent; stdcall;
begin
  Result :=  AObj.FindComponent(AName);
end;

function WebBrowser_GetNamePath(AObj: TWebBrowser): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure WebBrowser_Assign(AObj: TWebBrowser; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function WebBrowser_ClassName(AObj: TWebBrowser): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function WebBrowser_Equals(AObj: TWebBrowser; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function WebBrowser_GetHashCode(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function WebBrowser_ToString(AObj: TWebBrowser): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;

function WebBrowser_GetControlInterface(AObj: TWebBrowser): IWebBrowser2; stdcall;
begin
  Result :=  AObj.ControlInterface;
end;

function WebBrowser_GetDefaultInterface(AObj: TWebBrowser): IWebBrowser2; stdcall;
begin
  Result :=  AObj.DefaultInterface;
end;

function WebBrowser_GetApplication(AObj: TWebBrowser): IDispatch; stdcall;
begin
  Result :=  AObj.Application;
end;

function WebBrowser_GetParent(AObj: TWebBrowser): IDispatch; stdcall;
begin
  Result :=  AObj.Parent;
end;

function WebBrowser_GetDocument(AObj: TWebBrowser): IDispatch; stdcall;
begin
  Result :=  AObj.Document;
end;

function WebBrowser_GetLocationName(AObj: TWebBrowser): PChar; stdcall;
begin
  Result :=  PChar(AObj.LocationName);
end;

function WebBrowser_GetLocationURL(AObj: TWebBrowser): PChar; stdcall;
begin
  Result :=  PChar(AObj.LocationURL);
end;

function WebBrowser_GetBusy(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.Busy;
end;

function WebBrowser_GetName(AObj: TWebBrowser): PChar; stdcall;
begin
  Result :=  PChar(AObj.Name);
end;

function WebBrowser_GetFullName(AObj: TWebBrowser): PChar; stdcall;
begin
  Result :=  PChar(AObj.FullName);
end;

function WebBrowser_GetPath(AObj: TWebBrowser): PChar; stdcall;
begin
  Result :=  PChar(AObj.Path);
end;

function WebBrowser_GetReadyState(AObj: TWebBrowser): TOleEnum; stdcall;
begin
  Result :=  AObj.ReadyState;
end;

function WebBrowser_GetAnchors(AObj: TWebBrowser): TAnchors; stdcall;
begin
  Result :=  AObj.Anchors;
end;

procedure WebBrowser_SetAnchors(AObj: TWebBrowser; AValue: TAnchors); stdcall;
begin
  AObj.Anchors := AValue;
end;

function WebBrowser_GetTabStop(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.TabStop;
end;

procedure WebBrowser_SetTabStop(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.TabStop := AValue;
end;

function WebBrowser_GetAlign(AObj: TWebBrowser): TAlign; stdcall;
begin
  Result :=  AObj.Align;
end;

procedure WebBrowser_SetAlign(AObj: TWebBrowser; AValue: TAlign); stdcall;
begin
  AObj.Align := AValue;
end;

function WebBrowser_GetParentShowHint(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure WebBrowser_SetParentShowHint(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.ParentShowHint := AValue;
end;

function WebBrowser_GetPopupMenu(AObj: TWebBrowser): TPopupMenu; stdcall;
begin
  Result :=  AObj.PopupMenu;
end;

procedure WebBrowser_SetPopupMenu(AObj: TWebBrowser; AValue: TPopupMenu); stdcall;
begin
  AObj.PopupMenu := AValue;
end;

function WebBrowser_GetShowHint(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.ShowHint;
end;

procedure WebBrowser_SetShowHint(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.ShowHint := AValue;
end;

function WebBrowser_GetTabOrder(AObj: TWebBrowser): TTabOrder; stdcall;
begin
  Result :=  AObj.TabOrder;
end;

procedure WebBrowser_SetTabOrder(AObj: TWebBrowser; AValue: TTabOrder); stdcall;
begin
  AObj.TabOrder := AValue;
end;

function WebBrowser_GetStyleElements(AObj: TWebBrowser): TStyleElements; stdcall;
begin
  Result :=  AObj.StyleElements;
end;

procedure WebBrowser_SetStyleElements(AObj: TWebBrowser; AValue: TStyleElements); stdcall;
begin
  AObj.StyleElements := AValue;
end;

procedure WebBrowser_SetOnEnter(AObj: TWebBrowser; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnEnter := nil;
    TEventClass.Remove(AObj, geEnter);
    Exit;
  end;
  AObj.OnEnter := TEventClass.OnEnter;
  TEventClass.Add(AObj, geEnter, AEventId);
end;

procedure WebBrowser_SetOnExit(AObj: TWebBrowser; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnExit := nil;
    TEventClass.Remove(AObj, geExit);
    Exit;
  end;
  AObj.OnExit := TEventClass.OnExit;
  TEventClass.Add(AObj, geExit, AEventId);
end;

function WebBrowser_GetVisible(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.Visible;
end;

procedure WebBrowser_SetVisible(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.Visible := AValue;
end;

function WebBrowser_GetFullScreen(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.FullScreen;
end;

procedure WebBrowser_SetFullScreen(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.FullScreen := AValue;
end;

function WebBrowser_GetOffline(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.Offline;
end;

procedure WebBrowser_SetOffline(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.Offline := AValue;
end;

function WebBrowser_GetSilent(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.Silent;
end;

procedure WebBrowser_SetSilent(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.Silent := AValue;
end;

function WebBrowser_GetDoubleBuffered(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure WebBrowser_SetDoubleBuffered(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.DoubleBuffered := AValue;
end;

function WebBrowser_GetBrush(AObj: TWebBrowser): TBrush; stdcall;
begin
  Result :=  AObj.Brush;
end;

function WebBrowser_GetControlCount(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.ControlCount;
end;

function WebBrowser_GetHandle(AObj: TWebBrowser): HWND; stdcall;
begin
  Result :=  AObj.Handle;
end;

function WebBrowser_GetParentDoubleBuffered(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure WebBrowser_SetParentDoubleBuffered(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function WebBrowser_GetParentWindow(AObj: TWebBrowser): HWND; stdcall;
begin
  Result :=  AObj.ParentWindow;
end;

procedure WebBrowser_SetParentWindow(AObj: TWebBrowser; AValue: HWND); stdcall;
begin
  try
  AObj.ParentWindow := AValue;
  except
    on E: Exception do
      writeln(E.message);
  end;
end;

function WebBrowser_GetEnabled(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.Enabled;
end;

procedure WebBrowser_SetEnabled(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.Enabled := AValue;
end;

function WebBrowser_GetAction(AObj: TWebBrowser): TBasicAction; stdcall;
begin
  Result :=  AObj.Action;
end;

procedure WebBrowser_SetAction(AObj: TWebBrowser; AValue: TBasicAction); stdcall;
begin
  AObj.Action := AValue;
end;

function WebBrowser_GetBiDiMode(AObj: TWebBrowser): TBiDiMode; stdcall;
begin
  Result :=  AObj.BiDiMode;
end;

procedure WebBrowser_SetBiDiMode(AObj: TWebBrowser; AValue: TBiDiMode); stdcall;
begin
  AObj.BiDiMode := AValue;
end;

procedure WebBrowser_GetBoundsRect(AObj: TWebBrowser; var Result: TRect); stdcall;
begin
  Result :=  AObj.BoundsRect;
end;

procedure WebBrowser_SetBoundsRect(AObj: TWebBrowser; var AValue: TRect); stdcall;
begin
  AObj.BoundsRect := AValue;
end;

function WebBrowser_GetClientHeight(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.ClientHeight;
end;

procedure WebBrowser_SetClientHeight(AObj: TWebBrowser; AValue: Integer); stdcall;
begin
  AObj.ClientHeight := AValue;
end;

procedure WebBrowser_GetClientRect(AObj: TWebBrowser; var Result: TRect); stdcall;
begin
  Result :=  AObj.ClientRect;
end;

function WebBrowser_GetClientWidth(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.ClientWidth;
end;

procedure WebBrowser_SetClientWidth(AObj: TWebBrowser; AValue: Integer); stdcall;
begin
  AObj.ClientWidth := AValue;
end;

function WebBrowser_GetExplicitLeft(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.ExplicitLeft;
end;

function WebBrowser_GetExplicitTop(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.ExplicitTop;
end;

function WebBrowser_GetExplicitWidth(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.ExplicitWidth;
end;

function WebBrowser_GetExplicitHeight(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.ExplicitHeight;
end;

function WebBrowser_GetAlignWithMargins(AObj: TWebBrowser): LongBool; stdcall;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure WebBrowser_SetAlignWithMargins(AObj: TWebBrowser; AValue: LongBool); stdcall;
begin
  AObj.AlignWithMargins := AValue;
end;

function WebBrowser_GetLeft(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.Left;
end;

procedure WebBrowser_SetLeft(AObj: TWebBrowser; AValue: Integer); stdcall;
begin
  AObj.Left := AValue;
end;

function WebBrowser_GetTop(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.Top;
end;

procedure WebBrowser_SetTop(AObj: TWebBrowser; AValue: Integer); stdcall;
begin
  AObj.Top := AValue;
end;

function WebBrowser_GetWidth(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.Width;
end;

procedure WebBrowser_SetWidth(AObj: TWebBrowser; AValue: Integer); stdcall;
begin
  AObj.Width := AValue;
end;

function WebBrowser_GetHeight(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.Height;
end;

procedure WebBrowser_SetHeight(AObj: TWebBrowser; AValue: Integer); stdcall;
begin
  AObj.Height := AValue;
end;

function WebBrowser_GetCursor(AObj: TWebBrowser): TCursor; stdcall;
begin
  Result :=  AObj.Cursor;
end;

procedure WebBrowser_SetCursor(AObj: TWebBrowser; AValue: TCursor); stdcall;
begin
  AObj.Cursor := AValue;
end;

function WebBrowser_GetHint(AObj: TWebBrowser): PChar; stdcall;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure WebBrowser_SetHint(AObj: TWebBrowser; AValue: PChar); stdcall;
begin
  AObj.Hint := AValue;
end;

function WebBrowser_GetMargins(AObj: TWebBrowser): TMargins; stdcall;
begin
  Result :=  AObj.Margins;
end;

procedure WebBrowser_SetMargins(AObj: TWebBrowser; AValue: TMargins); stdcall;
begin
  AObj.Margins := AValue;
end;

function WebBrowser_GetComponentCount(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.ComponentCount;
end;

function WebBrowser_GetComponentIndex(AObj: TWebBrowser): Integer; stdcall;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure WebBrowser_SetComponentIndex(AObj: TWebBrowser; AValue: Integer); stdcall;
begin
  AObj.ComponentIndex := AValue;
end;

function WebBrowser_GetOwner(AObj: TWebBrowser): TComponent; stdcall;
begin
  Result :=  AObj.Owner;
end;

function WebBrowser_GetTag(AObj: TWebBrowser): NativeInt; stdcall;
begin
  Result :=  AObj.Tag;
end;

procedure WebBrowser_SetTag(AObj: TWebBrowser; AValue: NativeInt); stdcall;
begin
  AObj.Tag := AValue;
end;

function WebBrowser_GetControls(AObj: TWebBrowser; Index: Integer): TControl; stdcall;
begin
  Result :=  AObj.Controls[Index];
end;

function WebBrowser_GetComponents(AObj: TWebBrowser; AIndex: Integer): TComponent; stdcall;
begin
  Result :=  AObj.Components[AIndex];
end;

procedure Webbrowser_SetParent(AObj: TWebBrowser; AParent: TWinControl); stdcall;
begin
  //TWinControl(AObj.Parent) := AParent;
end;



exports
  WebBrowser_Create {$IFNDEF MSWINDOWS}name '_WebBrowser_Create'{$ENDIF},
  WebBrowser_Free {$IFNDEF MSWINDOWS}name '_WebBrowser_Free'{$ENDIF},
  WebBrowser_GoBack {$IFNDEF MSWINDOWS}name '_WebBrowser_GoBack'{$ENDIF},
  WebBrowser_GoForward {$IFNDEF MSWINDOWS}name '_WebBrowser_GoForward'{$ENDIF},
  WebBrowser_GoHome {$IFNDEF MSWINDOWS}name '_WebBrowser_GoHome'{$ENDIF},
  WebBrowser_GoSearch {$IFNDEF MSWINDOWS}name '_WebBrowser_GoSearch'{$ENDIF},
  WebBrowser_Navigate {$IFNDEF MSWINDOWS}name '_WebBrowser_Navigate'{$ENDIF},
  WebBrowser_Refresh {$IFNDEF MSWINDOWS}name '_WebBrowser_Refresh'{$ENDIF},
  WebBrowser_Refresh2 {$IFNDEF MSWINDOWS}name '_WebBrowser_Refresh2'{$ENDIF},
  WebBrowser_Stop {$IFNDEF MSWINDOWS}name '_WebBrowser_Stop'{$ENDIF},
  WebBrowser_Quit {$IFNDEF MSWINDOWS}name '_WebBrowser_Quit'{$ENDIF},
  WebBrowser_SetBounds {$IFNDEF MSWINDOWS}name '_WebBrowser_SetBounds'{$ENDIF},
  WebBrowser_CanFocus {$IFNDEF MSWINDOWS}name '_WebBrowser_CanFocus'{$ENDIF},
  WebBrowser_FlipChildren {$IFNDEF MSWINDOWS}name '_WebBrowser_FlipChildren'{$ENDIF},
  WebBrowser_Focused {$IFNDEF MSWINDOWS}name '_WebBrowser_Focused'{$ENDIF},
  WebBrowser_HandleAllocated {$IFNDEF MSWINDOWS}name '_WebBrowser_HandleAllocated'{$ENDIF},
  WebBrowser_Invalidate {$IFNDEF MSWINDOWS}name '_WebBrowser_Invalidate'{$ENDIF},
  WebBrowser_Realign {$IFNDEF MSWINDOWS}name '_WebBrowser_Realign'{$ENDIF},
  WebBrowser_Repaint {$IFNDEF MSWINDOWS}name '_WebBrowser_Repaint'{$ENDIF},
  WebBrowser_ScaleBy {$IFNDEF MSWINDOWS}name '_WebBrowser_ScaleBy'{$ENDIF},
  WebBrowser_SetFocus {$IFNDEF MSWINDOWS}name '_WebBrowser_SetFocus'{$ENDIF},
  WebBrowser_Update {$IFNDEF MSWINDOWS}name '_WebBrowser_Update'{$ENDIF},
  WebBrowser_BringToFront {$IFNDEF MSWINDOWS}name '_WebBrowser_BringToFront'{$ENDIF},
  WebBrowser_HasParent {$IFNDEF MSWINDOWS}name '_WebBrowser_HasParent'{$ENDIF},
  WebBrowser_Hide {$IFNDEF MSWINDOWS}name '_WebBrowser_Hide'{$ENDIF},
  WebBrowser_Perform {$IFNDEF MSWINDOWS}name '_WebBrowser_Perform'{$ENDIF},
  WebBrowser_SendToBack {$IFNDEF MSWINDOWS}name '_WebBrowser_SendToBack'{$ENDIF},
  WebBrowser_Show {$IFNDEF MSWINDOWS}name '_WebBrowser_Show'{$ENDIF},
  WebBrowser_GetTextBuf {$IFNDEF MSWINDOWS}name '_WebBrowser_GetTextBuf'{$ENDIF},
  WebBrowser_FindComponent {$IFNDEF MSWINDOWS}name '_WebBrowser_FindComponent'{$ENDIF},
  WebBrowser_GetNamePath {$IFNDEF MSWINDOWS}name '_WebBrowser_GetNamePath'{$ENDIF},
  WebBrowser_Assign {$IFNDEF MSWINDOWS}name '_WebBrowser_Assign'{$ENDIF},
  WebBrowser_ClassName {$IFNDEF MSWINDOWS}name '_WebBrowser_ClassName'{$ENDIF},
  WebBrowser_Equals {$IFNDEF MSWINDOWS}name '_WebBrowser_Equals'{$ENDIF},
  WebBrowser_GetHashCode {$IFNDEF MSWINDOWS}name '_WebBrowser_GetHashCode'{$ENDIF},
  WebBrowser_ToString {$IFNDEF MSWINDOWS}name '_WebBrowser_ToString'{$ENDIF},
  WebBrowser_GetControlInterface {$IFNDEF MSWINDOWS}name '_WebBrowser_GetControlInterface'{$ENDIF},
  WebBrowser_GetDefaultInterface {$IFNDEF MSWINDOWS}name '_WebBrowser_GetDefaultInterface'{$ENDIF},
  WebBrowser_GetApplication {$IFNDEF MSWINDOWS}name '_WebBrowser_GetApplication'{$ENDIF},
  WebBrowser_GetParent {$IFNDEF MSWINDOWS}name '_WebBrowser_GetParent'{$ENDIF},
  WebBrowser_GetDocument {$IFNDEF MSWINDOWS}name '_WebBrowser_GetDocument'{$ENDIF},
  WebBrowser_GetLocationName {$IFNDEF MSWINDOWS}name '_WebBrowser_GetLocationName'{$ENDIF},
  WebBrowser_GetLocationURL {$IFNDEF MSWINDOWS}name '_WebBrowser_GetLocationURL'{$ENDIF},
  WebBrowser_GetBusy {$IFNDEF MSWINDOWS}name '_WebBrowser_GetBusy'{$ENDIF},
  WebBrowser_GetName {$IFNDEF MSWINDOWS}name '_WebBrowser_GetName'{$ENDIF},
  WebBrowser_GetFullName {$IFNDEF MSWINDOWS}name '_WebBrowser_GetFullName'{$ENDIF},
  WebBrowser_GetPath {$IFNDEF MSWINDOWS}name '_WebBrowser_GetPath'{$ENDIF},
  WebBrowser_GetReadyState {$IFNDEF MSWINDOWS}name '_WebBrowser_GetReadyState'{$ENDIF},
  WebBrowser_GetAnchors {$IFNDEF MSWINDOWS}name '_WebBrowser_GetAnchors'{$ENDIF},
  WebBrowser_SetAnchors {$IFNDEF MSWINDOWS}name '_WebBrowser_SetAnchors'{$ENDIF},
  WebBrowser_GetTabStop {$IFNDEF MSWINDOWS}name '_WebBrowser_GetTabStop'{$ENDIF},
  WebBrowser_SetTabStop {$IFNDEF MSWINDOWS}name '_WebBrowser_SetTabStop'{$ENDIF},
  WebBrowser_GetAlign {$IFNDEF MSWINDOWS}name '_WebBrowser_GetAlign'{$ENDIF},
  WebBrowser_SetAlign {$IFNDEF MSWINDOWS}name '_WebBrowser_SetAlign'{$ENDIF},
  WebBrowser_GetParentShowHint {$IFNDEF MSWINDOWS}name '_WebBrowser_GetParentShowHint'{$ENDIF},
  WebBrowser_SetParentShowHint {$IFNDEF MSWINDOWS}name '_WebBrowser_SetParentShowHint'{$ENDIF},
  WebBrowser_GetPopupMenu {$IFNDEF MSWINDOWS}name '_WebBrowser_GetPopupMenu'{$ENDIF},
  WebBrowser_SetPopupMenu {$IFNDEF MSWINDOWS}name '_WebBrowser_SetPopupMenu'{$ENDIF},
  WebBrowser_GetShowHint {$IFNDEF MSWINDOWS}name '_WebBrowser_GetShowHint'{$ENDIF},
  WebBrowser_SetShowHint {$IFNDEF MSWINDOWS}name '_WebBrowser_SetShowHint'{$ENDIF},
  WebBrowser_GetTabOrder {$IFNDEF MSWINDOWS}name '_WebBrowser_GetTabOrder'{$ENDIF},
  WebBrowser_SetTabOrder {$IFNDEF MSWINDOWS}name '_WebBrowser_SetTabOrder'{$ENDIF},
  WebBrowser_GetStyleElements {$IFNDEF MSWINDOWS}name '_WebBrowser_GetStyleElements'{$ENDIF},
  WebBrowser_SetStyleElements {$IFNDEF MSWINDOWS}name '_WebBrowser_SetStyleElements'{$ENDIF},
  WebBrowser_SetOnEnter {$IFNDEF MSWINDOWS}name '_WebBrowser_SetOnEnter'{$ENDIF},
  WebBrowser_SetOnExit {$IFNDEF MSWINDOWS}name '_WebBrowser_SetOnExit'{$ENDIF},
  WebBrowser_GetVisible {$IFNDEF MSWINDOWS}name '_WebBrowser_GetVisible'{$ENDIF},
  WebBrowser_SetVisible {$IFNDEF MSWINDOWS}name '_WebBrowser_SetVisible'{$ENDIF},
  WebBrowser_GetFullScreen {$IFNDEF MSWINDOWS}name '_WebBrowser_GetFullScreen'{$ENDIF},
  WebBrowser_SetFullScreen {$IFNDEF MSWINDOWS}name '_WebBrowser_SetFullScreen'{$ENDIF},
  WebBrowser_GetOffline {$IFNDEF MSWINDOWS}name '_WebBrowser_GetOffline'{$ENDIF},
  WebBrowser_SetOffline {$IFNDEF MSWINDOWS}name '_WebBrowser_SetOffline'{$ENDIF},
  WebBrowser_GetSilent {$IFNDEF MSWINDOWS}name '_WebBrowser_GetSilent'{$ENDIF},
  WebBrowser_SetSilent {$IFNDEF MSWINDOWS}name '_WebBrowser_SetSilent'{$ENDIF},
  WebBrowser_GetDoubleBuffered {$IFNDEF MSWINDOWS}name '_WebBrowser_GetDoubleBuffered'{$ENDIF},
  WebBrowser_SetDoubleBuffered {$IFNDEF MSWINDOWS}name '_WebBrowser_SetDoubleBuffered'{$ENDIF},
  WebBrowser_GetBrush {$IFNDEF MSWINDOWS}name '_WebBrowser_GetBrush'{$ENDIF},
  WebBrowser_GetControlCount {$IFNDEF MSWINDOWS}name '_WebBrowser_GetControlCount'{$ENDIF},
  WebBrowser_GetHandle {$IFNDEF MSWINDOWS}name '_WebBrowser_GetHandle'{$ENDIF},
  WebBrowser_GetParentDoubleBuffered {$IFNDEF MSWINDOWS}name '_WebBrowser_GetParentDoubleBuffered'{$ENDIF},
  WebBrowser_SetParentDoubleBuffered {$IFNDEF MSWINDOWS}name '_WebBrowser_SetParentDoubleBuffered'{$ENDIF},
  WebBrowser_GetParentWindow {$IFNDEF MSWINDOWS}name '_WebBrowser_GetParentWindow'{$ENDIF},
  WebBrowser_SetParentWindow {$IFNDEF MSWINDOWS}name '_WebBrowser_SetParentWindow'{$ENDIF},
  WebBrowser_GetEnabled {$IFNDEF MSWINDOWS}name '_WebBrowser_GetEnabled'{$ENDIF},
  WebBrowser_SetEnabled {$IFNDEF MSWINDOWS}name '_WebBrowser_SetEnabled'{$ENDIF},
  WebBrowser_GetAction {$IFNDEF MSWINDOWS}name '_WebBrowser_GetAction'{$ENDIF},
  WebBrowser_SetAction {$IFNDEF MSWINDOWS}name '_WebBrowser_SetAction'{$ENDIF},
  WebBrowser_GetBiDiMode {$IFNDEF MSWINDOWS}name '_WebBrowser_GetBiDiMode'{$ENDIF},
  WebBrowser_SetBiDiMode {$IFNDEF MSWINDOWS}name '_WebBrowser_SetBiDiMode'{$ENDIF},
  WebBrowser_GetBoundsRect {$IFNDEF MSWINDOWS}name '_WebBrowser_GetBoundsRect'{$ENDIF},
  WebBrowser_SetBoundsRect {$IFNDEF MSWINDOWS}name '_WebBrowser_SetBoundsRect'{$ENDIF},
  WebBrowser_GetClientHeight {$IFNDEF MSWINDOWS}name '_WebBrowser_GetClientHeight'{$ENDIF},
  WebBrowser_SetClientHeight {$IFNDEF MSWINDOWS}name '_WebBrowser_SetClientHeight'{$ENDIF},
  WebBrowser_GetClientRect {$IFNDEF MSWINDOWS}name '_WebBrowser_GetClientRect'{$ENDIF},
  WebBrowser_GetClientWidth {$IFNDEF MSWINDOWS}name '_WebBrowser_GetClientWidth'{$ENDIF},
  WebBrowser_SetClientWidth {$IFNDEF MSWINDOWS}name '_WebBrowser_SetClientWidth'{$ENDIF},
  WebBrowser_GetExplicitLeft {$IFNDEF MSWINDOWS}name '_WebBrowser_GetExplicitLeft'{$ENDIF},
  WebBrowser_GetExplicitTop {$IFNDEF MSWINDOWS}name '_WebBrowser_GetExplicitTop'{$ENDIF},
  WebBrowser_GetExplicitWidth {$IFNDEF MSWINDOWS}name '_WebBrowser_GetExplicitWidth'{$ENDIF},
  WebBrowser_GetExplicitHeight {$IFNDEF MSWINDOWS}name '_WebBrowser_GetExplicitHeight'{$ENDIF},
  WebBrowser_GetAlignWithMargins {$IFNDEF MSWINDOWS}name '_WebBrowser_GetAlignWithMargins'{$ENDIF},
  WebBrowser_SetAlignWithMargins {$IFNDEF MSWINDOWS}name '_WebBrowser_SetAlignWithMargins'{$ENDIF},
  WebBrowser_GetLeft {$IFNDEF MSWINDOWS}name '_WebBrowser_GetLeft'{$ENDIF},
  WebBrowser_SetLeft {$IFNDEF MSWINDOWS}name '_WebBrowser_SetLeft'{$ENDIF},
  WebBrowser_GetTop {$IFNDEF MSWINDOWS}name '_WebBrowser_GetTop'{$ENDIF},
  WebBrowser_SetTop {$IFNDEF MSWINDOWS}name '_WebBrowser_SetTop'{$ENDIF},
  WebBrowser_GetWidth {$IFNDEF MSWINDOWS}name '_WebBrowser_GetWidth'{$ENDIF},
  WebBrowser_SetWidth {$IFNDEF MSWINDOWS}name '_WebBrowser_SetWidth'{$ENDIF},
  WebBrowser_GetHeight {$IFNDEF MSWINDOWS}name '_WebBrowser_GetHeight'{$ENDIF},
  WebBrowser_SetHeight {$IFNDEF MSWINDOWS}name '_WebBrowser_SetHeight'{$ENDIF},
  WebBrowser_GetCursor {$IFNDEF MSWINDOWS}name '_WebBrowser_GetCursor'{$ENDIF},
  WebBrowser_SetCursor {$IFNDEF MSWINDOWS}name '_WebBrowser_SetCursor'{$ENDIF},
  WebBrowser_GetHint {$IFNDEF MSWINDOWS}name '_WebBrowser_GetHint'{$ENDIF},
  WebBrowser_SetHint {$IFNDEF MSWINDOWS}name '_WebBrowser_SetHint'{$ENDIF},
  WebBrowser_GetMargins {$IFNDEF MSWINDOWS}name '_WebBrowser_GetMargins'{$ENDIF},
  WebBrowser_SetMargins {$IFNDEF MSWINDOWS}name '_WebBrowser_SetMargins'{$ENDIF},
  WebBrowser_GetComponentCount {$IFNDEF MSWINDOWS}name '_WebBrowser_GetComponentCount'{$ENDIF},
  WebBrowser_GetComponentIndex {$IFNDEF MSWINDOWS}name '_WebBrowser_GetComponentIndex'{$ENDIF},
  WebBrowser_SetComponentIndex {$IFNDEF MSWINDOWS}name '_WebBrowser_SetComponentIndex'{$ENDIF},
  WebBrowser_GetOwner {$IFNDEF MSWINDOWS}name '_WebBrowser_GetOwner'{$ENDIF},
  WebBrowser_GetTag {$IFNDEF MSWINDOWS}name '_WebBrowser_GetTag'{$ENDIF},
  WebBrowser_SetTag {$IFNDEF MSWINDOWS}name '_WebBrowser_SetTag'{$ENDIF},
  WebBrowser_GetControls {$IFNDEF MSWINDOWS}name '_WebBrowser_GetControls'{$ENDIF},
  WebBrowser_GetComponents {$IFNDEF MSWINDOWS}name '_WebBrowser_GetComponents'{$ENDIF},
  Webbrowser_SetParent {$IFNDEF MSWINDOWS}name '_Webbrowser_SetParent'{$ENDIF};

