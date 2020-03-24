// 不使用自动生成的
// 这里主要是对自动生成做兼容和补充相关

function ToUnixTime(ADateTime: TDateTime): TUnixDateTime; inline;
begin
  Result := DateTimeToUnix(LocalTimeToUniversal(ADateTime));
end;

function UnixToTime(ADateTime: TUnixDateTime): TDateTime; inline;
begin
  Result := UniversalTimeToLocal(UnixToDateTime(ADateTime));
end;

function ToPChar(AStr: string): PChar; inline;
begin
  Result := PChar(AStr);
end;

// 一些关于库的信息
type
  TStringEncoding = (seUnknown, seANSI, seUnicode, seUTF8);

function DLibStringEncoding: TStringEncoding; extdecl;
begin
  Result := seUTF8;
end;

function DLibVersion: Cardinal; extdecl;
begin
  // 共8位，2位2位的
  Result := $02000000;
end;

// 一些初始
procedure InitLazarusDef;
begin
  // 注册与VCL中类名一致的TPngImage
  TPicture.RegisterFileFormat(TPngImage.GetFileExtensions, rsPortableNetworkGraphic, TPngImage);
end;


// SysLocale  14byte
procedure DSysLocale(var AInfo: TSysLocale); extdecl;
begin
  AInfo := SysLocale;
end;


// TApplication
function Application_Instance: TApplication; extdecl;
begin
  Result := Application;
end;

{$IFDEF WINDOWS}
// 当IsLibrary = True时，是不会调用WidgetSet.CreateAppHandle方法
// 这时就需要手动去伪造一个 WidgetSet.AppHandle := CreateWindowsAppHandle;
// 这里的代码来自win32object.inc单元中 CreateAppHandle方法。
function CreateWindowsAppHandle: HWND;
var
  SysMenu: HMENU;
begin
  Result := CreateWindowW(@ClsNameW,
    PWideChar(UTF8ToUTF16(Application.Title)),
    WS_POPUP or WS_CLIPSIBLINGS or WS_SYSMENU or WS_MINIMIZEBOX, 0,
    0, 0, 0, 0, 0, HInstance, nil);
  if Result = 0 then
    Exit(0);
  AllocWindowInfo(Result);
  // remove useless menuitems from sysmenu
  SysMenu := Windows.GetSystemMenu(Result, False);
  Windows.DeleteMenu(SysMenu, SC_MAXIMIZE, MF_BYCOMMAND);
  Windows.DeleteMenu(SysMenu, SC_SIZE, MF_BYCOMMAND);
  Windows.DeleteMenu(SysMenu, SC_MOVE, MF_BYCOMMAND);
end;
{$ENDIF}

// 移动来自 MyLCL_Application.inc
procedure Application_Initialize(AObj: TApplication); extdecl;
begin
  AObj.Initialize;
{$IFDEF WINDOWS}
  WidgetSet.ThemeServices.UpdateThemes;
  WidgetSet.AppHandle := CreateWindowsAppHandle;
{$ENDIF}
end;

function Application_CreateForm(App: TApplication; AInitScale: LongBool): TForm; extdecl;
begin
  App.CreateForm(TGoForm, Result);
end;

procedure RemoveAllForms;
var
  C: TComponent;
  I: Integer;
begin
  for I := Application.ComponentCount - 1 downto 0 do
  begin
    C := Application.Components[I];
    if C is TCustomForm then
      C.Free;
  end;
end;

procedure Application_Run(AObj: TApplication); extdecl;
begin
// 先不用这里的了
//{$IFDEF WINDOWS}
  //if (AObj.MainForm <> nil) and AObj.ShowMainForm then AObj.MainForm.Show;
     //WidgetSet.AppRun(AObj.MyRunLoop);
//{$ELSE}
  AObj.Run;
//{$ENDIF}
  // 在Run结束后手动释放Application中的Form，主要是考虑在非Delphi运行环境中造成的一些资源清理问题
  RemoveAllForms;
end;

function Application_GetHandle(AObj: TApplication): THandle; extdecl;
begin
{$IFDEF WINDOWS}
   Result := WidgetSet.AppHandle;
{$ELSE}
  Result := 0;
{$ENDIF}
end;

procedure Application_SetHandle(AObj: TApplication; AValue: THandle); extdecl;
begin
  // 不允许设置
end;

// TForm

function Form_Create2(AOwner: TComponent; AInitScale: LongBool): TGoForm; extdecl;
begin
  Result := TGoForm.Create(AOwner);
end;

procedure Form_EnabledMaximize(AObj: TGoForm; AValue: Boolean); extdecl;
begin
  if AValue then
  begin
    if not(biMaximize in AObj.BorderIcons) then
      AObj.BorderIcons := AObj.BorderIcons + [biMaximize]
  end else
  begin
    if biMaximize in AObj.BorderIcons then
      AObj.BorderIcons := AObj.BorderIcons - [biMaximize]
  end;
end;

procedure Form_EnabledMinimize(AObj: TGoForm; AValue: Boolean); extdecl;
begin
  if AValue then
  begin
    if not(biMinimize in AObj.BorderIcons) then
      AObj.BorderIcons := AObj.BorderIcons + [biMinimize]
  end else
  begin
    if biMinimize in AObj.BorderIcons then
      AObj.BorderIcons := AObj.BorderIcons - [biMinimize]
  end;
end;

procedure Form_EnabledSystemMenu(AObj: TGoForm; AValue: Boolean); extdecl;
begin
  if AValue then
  begin
    if not(biSystemMenu in AObj.BorderIcons) then
      AObj.BorderIcons := AObj.BorderIcons + [biSystemMenu]
  end else
  begin
    if biSystemMenu in AObj.BorderIcons then
      AObj.BorderIcons := AObj.BorderIcons - [biSystemMenu]
  end;
end;

procedure Form_SetAllowDropFiles(AObj: TGoForm; AValue: LongBool); extdecl;
begin
  AObj.AllowDropFiles := AValue;
end;

function Form_GetAllowDropFiles(AObj: TGoForm): LongBool; extdecl;
begin
  Result := AObj.AllowDropFiles;
end;

procedure Form_SetOnDropFiles(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnDropFiles}
  {$DEFINE EventMethodName := FormOnDropFiles}
  EventMethodCode
end;

procedure Form_SetOnDestroy(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnDestroy}
  {$DEFINE EventMethodName := OnDestroy}
  EventMethodCode
end;

procedure Form_SetOnConstrainedResize(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnConstrainedResize}
  {$DEFINE EventMethodName := OnConstrainedResize}
  EventMethodCode
end;

procedure Form_SetOnDeactivate(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnDeactivate}
  {$DEFINE EventMethodName := OnDeactivate}
  EventMethodCode
end;

procedure Form_SetOnActivate(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnActivate}
  {$DEFINE EventMethodName := OnActivate}
  EventMethodCode
end;

procedure Form_SetOnWndProc(AObj: TGoForm; AEventId: NativeUInt); extdecl;
begin
  if AEventId = 0 then
  begin
    AObj.OnWndProc := nil;
    TMessageEventClass.Remove(AObj);
    Exit;
  end;
  AObj.OnWndProc := TMessageEventClass.OnWndProc;
  TMessageEventClass.Add(AObj, AEventId);
end;

procedure Form_SetShowInTaskBar(AObj: TGoForm; AVal: TShowInTaskbar); extdecl;
begin
  AObj.ShowInTaskBar := AVal;
end;

function Form_ShowInTaskBar(AObj: TGoForm): TShowInTaskbar; extdecl;
begin
  Result := AObj.ShowInTaskBar;
end;

procedure Form_ScaleForPPI(AObj: TGoForm; ANewPPI: Integer); extdecl;
begin
  AObj.ScaleForPPI(ANewPPI);
end;

procedure Form_ScaleControlsForDpi(AObj: TGoForm; ANewPPI: Integer); extdecl;
begin
  // no support
end;

procedure Form_ScaleForCurrentDpi(AObj: TGoForm); extdecl;
begin
  AObj.ScaleForCurrentDpi;
end;

procedure Form_InheritedWndProc(AObj: TGoForm; var AMsg: TLMessage); extdecl;
begin
  AObj.InheritedWndProc(AMsg);
end;

// mouse
function Mouse_Instance: TMouse; extdecl;
begin
  Result := Mouse;
end;

// screen
function Screen_Instance: TScreen; extdecl;
begin
  Result := Screen;
end;

function DTextToShortCut(AText: PChar): TShortCut; extdecl;
begin
  Result := TextToShortCut(AText);
end;

function DShortCutToText(AVal: TShortCut): PChar; extdecl;
begin
  Result := PChar(ShortCutToText(AVal));
end;

// TClipboard
function Clipboard_Instance: TClipboard; extdecl;
begin
  Result := Clipboard;
end;

function Clipboard_SetClipboard(NewClipboard: TClipboard): TClipboard; extdecl;
begin
  Result := SetClipboard(NewClipboard);
end;

// TMemoryStream
// function Write(const Buffer; Count: Longint): Longint;
function MemoryStream_Write(AObj: TMemoryStream; Buffer: Pointer; Count: Longint): Longint; extdecl;
begin
  Result := AObj.Write(Buffer^, Count);
end;

// function Read(var Buffer; Count: Longint): Longint;
function MemoryStream_Read(AObj: TMemoryStream; Buffer: Pointer; Count: Longint): Longint; extdecl;
begin
  Result := AObj.Read(Buffer^, Count);
end;

// TCanvas
// Canvas_BrushCopy
// Canvas_CopyRect
// Canvas_Draw1
// Canvas_Draw2
// Canvas_DrawFocusRect
// Canvas_FillRect
// Canvas_FrameRect
// Canvas_StretchDraw
// Canvas_TextRect1
// Canvas_TextRect2

procedure Canvas_BrushCopy(AObj: TCanvas; var Dest: TRect; Bitmap: TBitmap;
      var Source: TRect; Color: TColor); extdecl;
begin
  AObj.BrushCopy(Dest, Bitmap, Source, Color);
end;

procedure Canvas_CopyRect(AObj: TCanvas; var Dest: TRect; Canvas: TCanvas;
  var Source: TRect); extdecl;
begin
  AObj.CopyRect(Dest, Canvas, Source);
end;

procedure Canvas_Draw1(AObj: TCanvas; X, Y: Integer; Graphic: TGraphic); extdecl;
begin
  AObj.Draw(X, Y, Graphic);
end;

procedure Canvas_Draw2(AObj: TCanvas; X, Y: Integer; Graphic: TGraphic; Opacity: Byte); extdecl;
begin
  AObj.Draw(X, Y, Graphic);
end;

procedure Canvas_DrawFocusRect(AObj: TCanvas; var ARect: TRect); extdecl;
begin
  AObj.DrawFocusRect(ARect);
end;

procedure Canvas_FillRect(AObj: TCanvas; var Rect: TRect); extdecl;
begin
  AObj.FillRect(Rect);
end;

procedure Canvas_FrameRect(AObj: TCanvas; var Rect: TRect); extdecl;
begin
  AObj.FrameRect(Rect);
end;

procedure Canvas_StretchDraw(AObj: TCanvas; var Rect: TRect; Graphic: TGraphic); extdecl;
begin
  AObj.StretchDraw(Rect, Graphic);
end;

procedure Canvas_TextRect1(AObj: TCanvas; var Rect: TRect; X, Y: Integer; const Text: PChar); extdecl;
begin
  AObj.TextRect(Rect, X, Y, Text);
end;


type
// Note: tfComposited only supported by ThemeServices.DrawText
  TTextFormats = (tfBottom, tfCalcRect, tfCenter, tfEditControl, tfEndEllipsis,
    tfPathEllipsis, tfExpandTabs, tfExternalLeading, tfLeft, tfModifyString,
    tfNoClip, tfNoPrefix, tfRight, tfRtlReading, tfSingleLine, tfTop,
    tfVerticalCenter, tfWordBreak, tfHidePrefix, tfNoFullWidthCharBreak,
    tfPrefixOnly, tfTabStop, tfWordEllipsis, tfComposited);
  TTextFormat = set of TTextFormats;

(*

  TTextLayout = (tlTop, tlCenter, tlBottom);
  TTextStyle = packed record
    Alignment : TAlignment;  // TextRect Only: horizontal alignment
    Layout    : TTextLayout; // TextRect Only: vertical alignment
    SingleLine: boolean;     // If WordBreak is false then process #13, #10 as
    Clipping  : boolean;     // TextRect Only: Clip Text to passed Rectangle
    ExpandTabs: boolean;     // Replace #9 by apropriate amount of spaces (default is usually 8)
    ShowPrefix: boolean;     // TextRect Only: Process first single '&' per
    Wordbreak : boolean;     // TextRect Only: If line of text is too long
    Opaque    : boolean;     // TextRect: Fills background with current Brush
    SystemFont: Boolean;     // Use the system font instead of Canvas Font
    RightToLeft: Boolean;    //For RightToLeft text reading (Text Direction)
    EndEllipsis: Boolean;    // TextRect Only: If line of text is too long
  end;

  TAlignment = (taLeftJustify, taRightJustify, taCenter);

  TLeftRight = taLeftJustify..taRightJustify;

*)

function Canvas_TextRect2(AObj: TCanvas; var Rect: TRect; Text: PChar; var AOutStr: PChar; TextFormat: TTextFormat): Integer; extdecl;
var
  LStyle: TTextStyle;
begin
  Result := 0;
  LStyle := AObj.TextStyle;
  if TTextFormats.tfSingleLine in TextFormat then
    LStyle.SingleLine := True;
  if TTextFormats.tfTop in TextFormat then
    LStyle.Layout := TTextLayout.tlTop
  else if TTextFormats.tfVerticalCenter in TextFormat then
    LStyle.Layout := TTextLayout.tlCenter
  else if TTextFormats.tfBottom in TextFormat then
    LStyle.Layout := TTextLayout.tlBottom;
  if TTextFormats.tfNoClip in TextFormat then
    LStyle.Clipping := True;
  if TTextFormats.tfExpandTabs in TextFormat then
    LStyle.ExpandTabs:= True;
  if not(tfHidePrefix in TextFormat) and not(tfNoPrefix in TextFormat) then
    LStyle.ShowPrefix := True;
  if TTextFormats.tfWordBreak in TextFormat then
    LStyle.Wordbreak:= True;
  if tfEndEllipsis in TextFormat then
    LStyle.EndEllipsis:= True;
  if TTextFormats.tfLeft in TextFormat then
    LStyle.Alignment:= taLeftJustify
  else if TTextFormats.tfCenter in TextFormat then
    LStyle.Alignment:=taCenter
  else if TTextFormats.tfRight in TextFormat then
    LStyle.Alignment:= taRightJustify;
  AObj.TextRect(Rect, Rect.Left, Rect.Top, Text, LStyle);
end;


//  Canvas.Polygon();
//  Canvas.Polyline();
//  canvas.PolyBezier();
//  Canvas.PolyBezierTo();

procedure Canvas_Polygon(AObj: TCanvas; APoints: PPoint; ALen: Integer); extdecl;
begin
  AObj.Polygon(APoints, ALen);
end;

procedure Canvas_Polyline(AObj: TCanvas; APoints: PPoint; ALen: Integer); extdecl;
begin
  AObj.Polyline(APoints, ALen);
end;

procedure Canvas_PolyBezier(AObj: TCanvas; APoints: PPoint; ALen: Integer); extdecl;
begin
  AObj.PolyBezier(APoints, ALen);
end;

//
function Canvas_Pixels(AObj: TCanvas; X, Y: Integer): TColor; extdecl;
begin
  Result := AObj.Pixels[X, Y];
end;

procedure Canvas_SetPixels(AObj: TCanvas; X, Y: Integer; AColor: TColor); extdecl;
begin
  AObj.Pixels[X, Y] := AColor;
end;


// TImageList
// ImageList_Draw1
// ImageList_Draw2
// ImageList_DrawOverlay1
// ImageList_DrawOverlay2
// ImageList_GetIcon1
// ImageList_GetIcon2

procedure ImageList_Draw1(AObj: TImageList; Canvas: TCanvas; X, Y, Index: Integer; Enabled: LongBool); extdecl;
begin
  AObj.Draw(Canvas, X, Y, Index, Enabled);
end;

//procedure ImageList_Draw2(AObj: TImageList; Canvas: TCanvas; X, Y, Index: Integer;
//  ADrawingStyle: TDrawingStyle; AImageType: TImageType;
//  Enabled: LongBool); extdecl;
//begin
//  AObj.Draw(Canvas, X, Y, Index, ADrawingStyle, AImageType, Enabled);
//end;

procedure ImageList_DrawOverlay1(AObj: TImageList; Canvas: TCanvas; X, Y: Integer;
  ImageIndex: Integer; Overlay: TOverlay; Enabled: LongBool); extdecl;
begin
  AObj.DrawOverlay(Canvas, X, Y, ImageIndex, Overlay, Enabled);
end;
  
//procedure ImageList_DrawOverlay2(AObj: TImageList; Canvas: TCanvas; X, Y: Integer;
//  ImageIndex: Integer; Overlay: TOverlay; ADrawingStyle: TDrawingStyle;
//  AImageType: TImageType; Enabled: LongBool); extdecl;
//begin
//  AObj.DrawOverlay(Canvas, X, Y, ImageIndex, Overlay, ADrawingStyle, AImageType, Enabled);
//end;
  
procedure ImageList_GetIcon1(AObj: TImageList; Index: Integer; Image: TIcon); extdecl;
begin
  AObj.GetIcon(Index, Image);
end;

//procedure ImageList_GetIcon2(AObj: TImageList; Index: Integer; Image: TIcon; ADrawingStyle: TDrawingStyle; AImageType: TImageType); extdecl;
//begin
// // AObj.GetIcon(Index, Image, ADrawingStyle, AImageType);
//end;



// TListView
function ListView_GetScrollBars(AObj: TListView): TScrollStyle; extdecl;
begin
  Result :=  AObj.ScrollBars;
end;

procedure ListView_SetScrollBars(AObj: TListView; AValue: TScrollStyle); extdecl;
begin
  AObj.ScrollBars := AValue;
end;

//function ListView_CustomSort(AObj: TListView; SortProc: TLVCompare; lParam: NativeInt): LongBool; extdecl;
//begin
//  Result :=  AObj.CustomSort2(SortProc, lParam);
//end;

// TTreeView
function TreeView_GetScrollBars(AObj: TTreeView): TScrollStyle; extdecl;
begin
  Result :=  AObj.ScrollBars;
end;

procedure TreeView_SetScrollBars(AObj: TTreeView; AValue: TScrollStyle); extdecl;
begin
  AObj.ScrollBars := AValue;
end;

// TBitmap
procedure Bitmap_Clear(AObj: TBitmap); extdecl;
begin
  AObj.Clear;
end;

procedure Bitmap_BeginUpdate(AObj: TBitmap; ACanvasOnly: LongBool); extdecl;
begin
  AObj.BeginUpdate(ACanvasOnly);
end;

procedure Bitmap_EndUpdate(AObj: TBitmap; AStreamIsValid: LongBool); extdecl;
begin
  AObj.EndUpdate(AStreamIsValid);
end;

procedure Bitmap_LoadFromDevice(AObj: TBitmap; ADc: HDC); extdecl;
begin
  AObj.LoadFromDevice(ADc);
end;

// win api
{$IFNDEF WINDOWS}
function DSendMessage(hWd: HWND; msg: Cardinal; wParam: WParam; lParam: LParam): LResult; extdecl;
begin
  Result := LCLIntf.SendMessage(hWd, msg, wParam, lParam);
end;

function DPostMessage(hWd: HWND; msg: Cardinal; wParam: WParam; lParam: LParam): Boolean; extdecl;
begin
  Result := LCLIntf.PostMessage(hWd, msg, wParam, lParam);
end;

function DIsIconic(hWnd: HWND): LongBool; extdecl;
begin
  Result := LCLIntf.IsIconic(hWnd);
end;

function DIsWindow(hWnd: HWND): LongBool; extdecl;
begin
  Result := LCLIntf.IsWindow(hWnd);
end;

function DIsZoomed(hWnd: HWND): LongBool; extdecl;
begin
  Result := LCLIntf.IsZoomed(hWnd);
end;

function DIsWindowVisible(hWnd: HWND): LongBool; extdecl;
begin
  Result := LCLIntf.IsWindowVisible(hWnd);
end;

function DGetDC(hWnd: HWND): HDC; extdecl;
begin
  Result := LCLIntf.GetDC(hWnd);
end;

function DReleaseDC(hWnd: HWND; dc: HDC): Integer; extdecl;
begin
  Result := LCLIntf.ReleaseDC(hWnd, dc);
end;

function DSetForegroundWindow(hWnd: HWND): LongBool; extdecl;
begin
  Result := LCLIntf.SetForegroundWindow(hWnd);
end;

function DWindowFromPoint(point: PPoint): HWND; extdecl;
begin
  if point = nil then
    Exit(0);
  Result := LCLIntf.WindowFromPoint(PPoint(point)^);
end;
{$ENDIF}


// Other

procedure SetEventCallback(APtr: Pointer); extdecl;
begin
  GEventCallbackPtr := TEventCallbackPtr(APtr);
end;

procedure SetMessageCallback(APtr: Pointer); extdecl;
begin
  GMessageCallbackPtr := TMessageCallbackPtr(APtr);
end;

procedure SetThreadSyncCallback(APtr: Pointer); extdecl;
begin
  GThreadSyncCallbackPtr := APtr;
end;

// 从string数组中获取成员
function DGetStringArrOf(P: Pointer; AIndex: NativeInt): PChar; extdecl;
begin
  Result := PChar(PString(PByte(P) + AIndex * SizeOf(string))^);
end;

function DStrLen(p: PChar): NativeInt; extdecl;
begin
  Result := Length(p);
end;

procedure DMove(Src, Dest: Pointer; Len: NativeInt); extdecl;
begin
  Move(Src^, Dest^, Len);
end;

function DGetClassName(Obj: TObject): PChar; extdecl;
begin
  Result := PChar(string(Obj.ClassName));
end;

procedure DShowMessage(AMsg: PChar); extdecl;
begin
  ShowMessage(AMsg);
end;

function DGetMainInstance: HINST; extdecl;
begin
{$IFDEF WINDOWS}
  Result := GetModuleHandle(nil);
{$ELSE}
  Result := 0;
{$ENDIF}
end;

function DMessageDlg(const Msg: PChar; DlgType: TMsgDlgType;
  Buttons: TMsgDlgButtons; HelpCtx: Longint): Integer; extdecl;
begin
  Result := MessageDlg(Msg, DlgType, Buttons, HelpCtx);
end;

// 在主线程中运行
procedure DSynchronize(AUseMsg: LongBool); extdecl;
begin
  TThread.Synchronize(nil, TEventClass.ThreadProc);
end;

procedure DSysOpen(FileName: PChar); extdecl;
begin
 {$IFDEF WINDOWS}
   ShellExecute(0, nil, FileName, nil, nil, SW_SHOW);
 {$ELSE}
   OpenURL(FileName);
 {$ENDIF}
end;

// 不想导入包，所以干脆用delphi的得了
function DExtractFilePath(AFileName: PChar): PChar; extdecl;
begin
  Result := PChar(ExtractFilePath(AFileName));
end;

function DFileExists(AFileName: PChar): LongBool; extdecl;
begin
  Result := FileExists(AFileName);
end;

function DSelectDirectory1(var Directory: PChar;
  Options: TSelectDirOpts; HelpCtx: Longint): LongBool; extdecl;
var
  LDir: string;
begin
  Result := SelectDirectory(LDir, Options, HelpCtx);
  if Result then
    Directory := PChar(LDir);
end;

function DSelectDirectory2(Caption, Root: PChar;
  var Directory: PChar; Options: Cardinal{无效参数}; Parent: TWinControl{无效参数}): LongBool;  extdecl;
var
  LDir: string;
begin
  Result := SelectDirectory(Caption, Root, LDir, False, 0);
  if Result then
    Directory := PChar(LDir);
end;



function DInputBox(ACaption, APrompt, ADefault: PChar): PChar; extdecl;
begin
  Result := PChar(InputBox(ACaption, APrompt, ADefault));
end;

function DInputQuery(ACaption, APrompt: PChar; Value: PChar; out AOut: PChar): LongBool; extdecl;
var
  S: string;
begin
  S := Value;
  Result := InputQuery(string(ACaption), string(APrompt), S);
  if Result then
    AOut := PChar(S);
end;


procedure SetGlobalFormScaled(AVal: LongBool); extdecl;
begin
  // 先移除掉，待后我有2k屏再说吧。
  // Application.Scaled := AVal;
end;


{$IFDEF WINDOWS}
// 路径合并
function PathCombine(A1, A2: string): string; inline;
begin
  if A1[Length(A1)] <> PathDelim then
    A1 := A1 + PathDelim;
  if (Length(A2) > 0) and (A2[1] = PathDelim) then
    Delete(A2, 1, 1);
  Result := A1 + A2;
end;

// 创建URL快捷方式
procedure DCreateURLShortCut(ADestPath, AShortCutName, AURL: PChar); extdecl;
var
  Ini: TIniFile;
begin
  Ini := TIniFile.Create(PathCombine(ADestPath, AShortCutName) + '.url');
  try
    Ini.WriteString('{000214A0-0000-0000-C000-000000000046}', 'Prop3', '19,2');
    Ini.WriteString('InternetShortcut', 'URL', AURL);
  finally
    Ini.Free;
  end;
end;

// 创建快捷方式
function DCreateShortCut(ADestPath, AShortCutName, ASrcFileName, AIconFileName, ADescription, ACmdArgs: PChar): LongBool; extdecl;
var
  LObject : IUnknown;
  LPLink : IShellLink;
  LPPFile : IPersistFile;
  LinkFilename : WideString;
begin
  CoInitializeEx(nil, 0);
  try
    LObject := CreateComObject(CLSID_ShellLink);
    LPLink := LObject as IShellLink;
    LPPFile := LObject as IPersistFile;
    LPLink.SetPath(ASrcFileName);
    LPLink.SetWorkingDirectory(PChar(ExtractFilePath(ASrcFileName)));
    LPLink.SetDescription(ADescription);
    if AIconFileName = '' then
      LPLink.SetIconLocation(ASrcFileName, 0)
    else LPLink.SetIconLocation(AIconFileName, 0);
    LPLink.SetArguments(ACmdArgs);
    LinkFilename := WideString(PathCombine(ADestPath, AShortCutName) + '.lnk');
    Result := LPPFile.Save(PWideChar(LinkFilename), FALSE) = S_OK;
  finally
    CoUninitialize;
  end;
end;

{$ENDIF}


// 属性设置，这里主要提供给多语言使用
procedure DSetPropertyValue(Instance: TObject; PropName, Value: PChar); extdecl;
var
  LPropInfo: PPropInfo;
begin
  LPropInfo := GetPropInfo(Instance, PropName);
  if LPropInfo <> nil then
  begin
    case LPropInfo^.PropType^.Kind of
      tkInteger, tkInt64:
        SetInt64Prop(Instance, LPropInfo, StrToIntDef(Value, -1));
      tkChar, tkSString, tkLString, tkAString, tkWString, tkWChar, tkUString,tkUChar:
        SetStrProp(Instance, LPropInfo, Value);
    end;
  end;
end;

procedure DSetPropertySecValue(Instance: TObject; PropName, SecPropName, Value: PChar); extdecl;
type
  TCallIndexGetClassProc = function(AData: Pointer; AIndex: Integer): Pointer;
var
  LPropInfo: PPropInfo;
  LObj: TObject;
  LPropName: string;
  LP, LIndex: Integer;
begin
  LPropName := PropName;
  LIndex := -1;
  LP := LPropName.IndexOf('[');
  if LP > 0 then
  begin
    LIndex := StrToIntDef(LPropName.Substring(LP + 1, LPropName.IndexOf(']') - LP - 1), -1);
    LPropName := LPropName.Substring(0, LP);
  end;
  LPropInfo := GetPropInfo(Instance, LPropName);
  if LPropInfo <> nil then
  begin
    LObj := GetObjectProp(Instance, LPropInfo);
    if LObj <> nil then
    begin
      if LIndex <> -1 then
      begin
        if LPropInfo^.PropType^.Kind = tkClass then
        begin
          if LObj is TListColumns then
          begin
            if (LIndex >= 0) and (LIndex < TListColumns(LObj).Count) and string(SecPropName).Equals('Caption') then
              TListColumns(LObj)[LIndex].Caption := Value;
          end else
          if LObj is TStrings then
          begin
            if (LIndex >= 0) and (LIndex < TStrings(LObj).Count) and (string(SecPropName).Equals('String') or string(SecPropName).IsEmpty) then
              TStrings(LObj)[LIndex] := Value;
          end;
        end;
      end else
        DSetPropertyValue(LObj, SecPropName, Value);
    end;
  end;
end;


//
//GUIDToString
function DGUIDToString(var AGUID: TGUID): PChar; extdecl;
begin
  Result := PChar(GUIDToString(AGUID));
end;
//StringToGUID
procedure DStringToGUID(AGUIDStr: PChar; out AGUID: TGUID); extdecl;
begin
  AGUID := StringToGUID(AGUIDStr);
end;

//CreateGUID
procedure DCreateGUID(out AGUID: TGUID); extdecl;
begin
  CreateGUID(AGUID);//TGUID.NewGuid;
end;

// Printer
function Printer_Instance: TPrinter; extdecl;
begin
  Result := Printer;
end;

//function DPrinter_SetPrinter(NewPrinter: TPrinter): TPrinter; extdecl;
//begin
//  Result := SetPrinter(NewPrinter);
//end;
//
// 因参数与Delphi中不一致，这里要伪造下参数，后面  ADriver, APort, ADeviceMode 三个参数无效
procedure Printer_SetPrinter(Obj: TPrinter; aName, ADriver, APort: PChar; ADeviceMode: NativeUInt); extdecl;
begin
  Obj.SetPrinter(aName);
end;



//----------------------------------------------------
// 一些lib的资源
type
  TResItem = record
    Name: PChar;
    ValuePtr: Pointer;
  end;

const
  GLibResouces: array[0..40] of TResItem = (
     // 27
    (Name: 'SOpenFileTitle'; ValuePtr: @rsmbOpen),
    (Name: 'SOKButton'; ValuePtr: @rsMbOK),
    (Name: 'SCancelButton'; ValuePtr: @rsMbCancel),
    (Name: 'SYesButton'; ValuePtr: @rsMbYes),
    (Name: 'SNoButton'; ValuePtr: @rsMbNo),
    (Name: 'SHelpButton'; ValuePtr: @rsMbHelp),
    (Name: 'SCloseButton'; ValuePtr: @rsMbClose),
    (Name: 'SIgnoreButton'; ValuePtr: @rsMbIgnore),
    (Name: 'SRetryButton'; ValuePtr: @rsMbRetry),
    (Name: 'SAbortButton'; ValuePtr: @rsMbAbort),
    (Name: 'SAllButton'; ValuePtr: @rsMbAll),
    (Name: 'SMsgDlgWarning'; ValuePtr: @rsMtWarning),
    (Name: 'SMsgDlgError'; ValuePtr: @rsMtError),
    (Name: 'SMsgDlgInformation'; ValuePtr: @rsMtInformation),
    (Name: 'SMsgDlgConfirm'; ValuePtr: @rsMtConfirmation),
    (Name: 'SMsgDlgYes'; ValuePtr: @rsMbYes),
    (Name: 'SMsgDlgNo'; ValuePtr: @rsMbNo),
    (Name: 'SMsgDlgOK'; ValuePtr: @rsMbOK),
    (Name: 'SMsgDlgCancel'; ValuePtr: @rsMbCancel),
    (Name: 'SMsgDlgHelp'; ValuePtr: @rsMbHelp),
    //(Name: 'SMsgDlgHelpNone'; ValuePtr: @SMsgDlgHelpNone),
    //(Name: 'SMsgDlgHelpHelp'; ValuePtr: @SMsgDlgHelpHelp),
    (Name: 'SMsgDlgAbort'; ValuePtr: @rsMbAbort),
    (Name: 'SMsgDlgRetry'; ValuePtr: @rsMbRetry),
    (Name: 'SMsgDlgIgnore'; ValuePtr: @rsMbIgnore),
    (Name: 'SMsgDlgAll'; ValuePtr: @rsMbAll),
    (Name: 'SMsgDlgNoToAll'; ValuePtr: @rsMbNoToAll),
    (Name: 'SMsgDlgYesToAll'; ValuePtr: @rsMbYesToAll),
    (Name: 'SMsgDlgClose'; ValuePtr: @rsMbClose),

     // TFindDialog And TReplaceDialog
     // 9
    (Name: 'rsFind'; ValuePtr: @rsFind),
    (Name: 'rsHelp'; ValuePtr: @rsHelp),
    (Name: 'rsWholeWordsOnly'; ValuePtr: @rsWholeWordsOnly),
    (Name: 'rsCaseSensitive'; ValuePtr: @rsCaseSensitive),
    (Name: 'rsEntireScope'; ValuePtr: @rsEntireScope),
    (Name: 'rsText'; ValuePtr: @rsText),
    (Name: 'rsDirection'; ValuePtr: @rsDirection),
    (Name: 'rsForward'; ValuePtr: @rsForward),
    (Name: 'rsBackward'; ValuePtr: @rsBackward),

    (Name: 'rsFindMore'; ValuePtr: @rsFindMore),
    (Name: 'rsReplace'; ValuePtr: @rsReplace),
    (Name: 'rsReplaceAll'; ValuePtr: @rsReplaceAll),

    (Name: 'rsSelectcolorTitle'; ValuePtr: @rsSelectcolorTitle),
    (Name: 'rsSelectFontTitle'; ValuePtr: @rsSelectFontTitle)

  );

function DGetLibResouceCount: Integer; extdecl;
begin
  Result := Length(GLibResouces);
end;

procedure DGetLibResouceItem(AIndex: Integer; out AResult: TResItem); extdecl;
begin
  if (AIndex >= 0)  and (AIndex < Length(GLibResouces)) then
    AResult := GLibResouces[AIndex];
end;

// 说不好有没有内存泄露，啊，先这样呗
procedure DModifyLibResouce(APtr: Pointer; AValue: PChar); extdecl;
begin
  if APtr <> nil then
    PResStringRec(APtr)^ := string(AValue);
end;


type TLibType = (ltVCL, ltLCL);

function DGetLibType: TLibType; extdecl;
begin
  Result := ltLCL;
end;


// dylib
type
  TSyscall0  = function: UInt64; extdecl;
  TSyscall1  = function(A1: Pointer): UInt64; extdecl;
  TSyscall2  = function(A1, A2: Pointer): UInt64; extdecl;
  TSyscall3  = function(A1, A2, A3: Pointer): UInt64; extdecl;
  TSyscall4  = function(A1, A2, A3, A4: Pointer): UInt64; extdecl;
  TSyscall5  = function(A1, A2, A3, A4, A5: Pointer): UInt64; extdecl;
  TSyscall6  = function(A1, A2, A3, A4, A5, A6: Pointer): UInt64; extdecl;
  TSyscall7  = function(A1, A2, A3, A4, A5, A6, A7: Pointer): UInt64; extdecl;
  TSyscall8  = function(A1, A2, A3, A4, A5, A6, A7, A8: Pointer): UInt64; extdecl;
  TSyscall9  = function(A1, A2, A3, A4, A5, A6, A7, A8, A9: Pointer): UInt64; extdecl;
  TSyscall10 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10: Pointer): UInt64; extdecl;
  TSyscall11 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11: Pointer): UInt64; extdecl;
  TSyscall12 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12: Pointer): UInt64; extdecl;

function MySyscall(AProc:Pointer; ALen: NativeInt; A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12: Pointer): UInt64; extdecl;
begin
  Result := 0;
  if AProc = nil then Exit;
  try
    case ALen of
      0 : Result := TSyscall0(AProc)();
      1 : Result := TSyscall1(AProc) (A1);
      2 : Result := TSyscall2(AProc) (A1, A2);
      3 : Result := TSyscall3(AProc) (A1, A2, A3);
      4 : Result := TSyscall4(AProc) (A1, A2, A3, A4);
      5 : Result := TSyscall5(AProc) (A1, A2, A3, A4, A5);
      6 : Result := TSyscall6(AProc) (A1, A2, A3, A4, A5, A6);
      7 : Result := TSyscall7(AProc) (A1, A2, A3, A4, A5, A6, A7);
      8 : Result := TSyscall8(AProc) (A1, A2, A3, A4, A5, A6, A7, A8);
      9 : Result := TSyscall9(AProc) (A1, A2, A3, A4, A5, A6, A7, A8, A9);
      10: Result := TSyscall10(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10);
      11: Result := TSyscall11(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11);
      12: Result := TSyscall12(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12);
    else
      Exit;
    end;
  except
    on E: Exception do
    begin
      if Assigned(Application) then
      begin
        if Assigned(Application.OnException) then
          Application.OnException(Application, E)
        else
          MessageDlg(E.Message,  mtError, [mbOK], 0);
      end else
        DebugLn('Syscall' + ALen.ToString + ' Error: ' + E.Message);
    end;
  end;
end;



exports
  // 导出动态调用的
  MySyscall,

  // SysLocaled
  DSysLocale,

  // 库的信息
  DLibStringEncoding,
  DLibVersion,

  // TApplication
  Application_Instance,
  Application_CreateForm,
  Application_Run,
  SetGlobalFormScaled,
  Application_Initialize,
  Application_GetHandle,
  Application_SetHandle,

  // TForm
  Form_Create2,
  Form_EnabledMaximize,
  Form_EnabledMinimize,
  Form_EnabledSystemMenu,
  Form_SetAllowDropFiles,
  Form_GetAllowDropFiles,
  Form_SetOnDropFiles,
  Form_SetOnDestroy,
  Form_SetOnConstrainedResize,
  Form_SetOnDeactivate,
  Form_SetOnActivate,
  Form_SetOnWndProc,
  Form_SetShowInTaskBar,
  Form_ShowInTaskBar,
  Form_ScaleForPPI,
  Form_ScaleControlsForDpi,
  Form_ScaleForCurrentDpi,
  Form_InheritedWndProc,


  // TMouse
  Mouse_Instance,
  Screen_Instance,

{$IFNDEF WINDOWS}
  DSendMessage,
  DPostMessage,
  DIsIconic,
  DIsWindow,
  DIsZoomed,
  DIsWindowVisible,
  DGetDC,
  DReleaseDC,
  DSetForegroundWindow,
  DWindowFromPoint,
{$ENDIF}

{$IFDEF LCLCocoa}
  NSWindow_FromForm,
  NSWindow_titleVisibility,
  NSWindow_setTitleVisibility,
  NSWindow_titlebarAppearsTransparent,
  NSWindow_setTitlebarAppearsTransparent,
  NSWindow_styleMask,
  NSWindow_setStyleMask,
  NSWindow_setRepresentedURL,
  NSWindow_release,
{$ENDIF}

  // Other
  SetEventCallback,
  SetMessageCallback,
  SetThreadSyncCallback,

  DGetStringArrOf,
  DStrLen,
  DMove,
  DGetClassName,
  DShowMessage,
  DGetMainInstance,
  DSynchronize,
  DSysOpen,
  DTextToShortCut,
  DShortCutToText,
  DMessageDlg,
  DExtractFilePath,
  DFileExists,

  DSelectDirectory1,
  DSelectDirectory2,
  DInputBox,
  DInputQuery,

{$IFDEF WINDOWS}
  // ShortCut
  DCreateURLShortCut,
  DCreateShortCut,
{$ENDIF}

  // SetProperty
  DSetPropertyValue,
  DSetPropertySecValue,

  // TMemoryStream
  MemoryStream_Write,
  MemoryStream_Read,

  // TClipboard
  Clipboard_Instance,
  Clipboard_SetClipboard,

  // TCanavs
  Canvas_BrushCopy,
  Canvas_CopyRect,
  Canvas_Draw1,
  Canvas_Draw2,
  Canvas_DrawFocusRect,
  Canvas_FillRect,
  Canvas_FrameRect,
  Canvas_StretchDraw,
  Canvas_TextRect1,
  Canvas_TextRect2,
  Canvas_Polygon,
  Canvas_Polyline,
  canvas_PolyBezier,
  //Canvas_PolyBezierTo,
  Canvas_Pixels,
  Canvas_SetPixels,

  // TImageList
  ImageList_Draw1,
  //ImageList_Draw2,
  ImageList_DrawOverlay1,
  //ImageList_DrawOverlay2,
  ImageList_GetIcon1,
  //ImageList_GetIcon2,

  // TListView
  ListView_GetScrollBars,
  ListView_SetScrollBars,
  //ListView_CustomSort,

  // TTreeView
  TreeView_GetScrollBars,
  TreeView_SetScrollBars,

  // TBitmap
  Bitmap_Clear,
  Bitmap_BeginUpdate,
  Bitmap_EndUpdate,
  Bitmap_LoadFromDevice,

  // Form uFormDesignerFile.pas
  ResFormLoadFromStream,
  ResFormLoadFromFile,
  ResFormLoadFromResourceName,

  // TGUID
  DGUIDToString,
  DStringToGUID,
  DCreateGUID,

  // Printer
  Printer_Instance,
  Printer_SetPrinter,

  // libResouces
  DGetLibResouceCount,
  DGetLibResouceItem,
  DModifyLibResouce,

{$IFDEF LINUX}
  GdkWindow_GetXId,
  GtkWidget_GetGtkFixed,
  GdkWindow_FromForm,
{$ENDIF}
  DGetLibType
;
