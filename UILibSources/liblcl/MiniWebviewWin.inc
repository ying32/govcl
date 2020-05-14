

{*******************************************************}
{                                                       }
{       govcl windows WebBrowser                        }
{                                                       }
{       根据VCL原有的TWebBrowser精简的一个，用于在      }
{       Lazarus中也能使用                               }
{                                                       }
{       版权所有 (C) ying32                             }
{                                                       }
{*******************************************************}

{$IFDEF MSWINDOWS}

{$IFDEF FPC}
  {$mode delphi}
{$ENDIF}

interface

uses
  Windows, Messages, SysUtils, Variants, Classes, Graphics,
  Controls, Forms, ActiveX, Menus, ComObj, MSHTML, WebJSExternal_TLB;

const

  DOCHOSTUIFLAG_DIALOG                     = $00000001;
  DOCHOSTUIFLAG_DISABLE_HELP_MENU          = $00000002;
  DOCHOSTUIFLAG_NO3DBORDER                 = $00000004;
  DOCHOSTUIFLAG_SCROLL_NO                  = $00000008;
  DOCHOSTUIFLAG_DISABLE_SCRIPT_INACTIVE    = $00000010;
  DOCHOSTUIFLAG_OPENNEWWIN                 = $00000020;
  DOCHOSTUIFLAG_DISABLE_OFFSCREEN          = $00000040;
  DOCHOSTUIFLAG_FLAT_SCROLLBAR             = $00000080;
  DOCHOSTUIFLAG_DIV_BLOCKDEFAULT           = $00000100;
  DOCHOSTUIFLAG_ACTIVATE_CLIENTHIT_ONLY    = $00000200;
  DOCHOSTUIFLAG_OVERRIDEBEHAVIORFACTORY    = $00000400;
  DOCHOSTUIFLAG_CODEPAGELINKEDFONTS        = $00000800;
  DOCHOSTUIFLAG_URL_ENCODING_DISABLE_UTF8  = $00001000;
  DOCHOSTUIFLAG_URL_ENCODING_ENABLE_UTF8   = $00002000;
  DOCHOSTUIFLAG_ENABLE_FORMS_AUTOCOMPLETE  = $00004000;
  DOCHOSTUIFLAG_ENABLE_INPLACE_NAVIGATION  = $00010000;
  DOCHOSTUIFLAG_IME_ENABLE_RECONVERSION    = $00020000;
  DOCHOSTUIFLAG_THEME                      = $00040000;
  DOCHOSTUIFLAG_NOTHEME                    = $00080000;
  DOCHOSTUIFLAG_NOPICS                     = $00100000;
  DOCHOSTUIFLAG_NO3DOUTERBORDER            = $00200000;
  DOCHOSTUIFLAG_DISABLE_EDIT_NS_FIXUP      = $1;
  DOCHOSTUIFLAG_LOCAL_MACHINE_ACCESS_CHECK = $1;
  DOCHOSTUIFLAG_DISABLE_UNTRUSTEDPROTOCOL  = $1;
  DOCHOSTUIDBLCLK_DEFAULT                  = 0;
  DOCHOSTUIDBLCLK_SHOWPROPERTIES           = 1;
  DOCHOSTUIDBLCLK_SHOWCODE                 = 2;
  DOCHOSTUITYPE_BROWSE                     = 0;
  DOCHOSTUITYPE_AUTHOR                     = 1;


type


  OLECMDID = TOleEnum;
  OLECMDF = TOleEnum;
  OLECMDEXECOPT = TOleEnum;
  tagREADYSTATE = TOleEnum;

  TReadyState = (rsUninitialized, rsLoading, rsLoaded, rsInterActive, rsComplete);
  
  TWebStatusTextChangeEvent = procedure(Sender: TObject; const AText: string) of object;
  TWebTitleChangeEvent = procedure(Sender: TObject; const AText: string) of object;
  TWebDocumentCompleteEvent = procedure(Sender: TObject; const pDisp: Pointer; const AURL: string) of object;
  TWebWindowClosingEvent = procedure(Sender: TObject; AIsChildWindow: Boolean; var ACancel: Boolean) of object;
  TWebJSExternalEvent = procedure(Sender: TObject; const AFunc: string; const AArgs: WideString; var ARetval: WideString) of object;

  
// *********************************************************************//
// Interface: IWebBrowser
// Flags:     (4432) Hidden Dual OleAutomation Dispatchable
// GUID:      {EAB22AC1-30C1-11CF-A7EB-0000C05BAE0B}
// *********************************************************************//
  IWebBrowser = interface(IDispatch)
    ['{EAB22AC1-30C1-11CF-A7EB-0000C05BAE0B}']
    procedure GoBack; safecall;
    procedure GoForward; safecall;
    procedure GoHome; safecall;
    procedure GoSearch; safecall;
    procedure Navigate(const URL: WideString; const Flags: OleVariant;
                       const TargetFrameName: OleVariant; const PostData: OleVariant;
                       const Headers: OleVariant); safecall;
    procedure Refresh; safecall;
    procedure Refresh2(const Level: OleVariant); safecall;
    procedure Stop; safecall;
    function Get_Application: IDispatch; safecall;
    function Get_Parent: IDispatch; safecall;
    function Get_Container: IDispatch; safecall;
    function Get_Document: IDispatch; safecall;
    function Get_TopLevelContainer: WordBool; safecall;
    function Get_type_: WideString; safecall;
    function Get_Left: Integer; safecall;
    procedure Set_Left(pl: Integer); safecall;
    function Get_Top: Integer; safecall;
    procedure Set_Top(pl: Integer); safecall;
    function Get_Width: Integer; safecall;
    procedure Set_Width(pl: Integer); safecall;
    function Get_Height: Integer; safecall;
    procedure Set_Height(pl: Integer); safecall;
    function Get_LocationName: WideString; safecall;
    function Get_LocationURL: WideString; safecall;
    function Get_Busy: WordBool; safecall;
    property Application: IDispatch read Get_Application;
    property Parent: IDispatch read Get_Parent;
    property Container: IDispatch read Get_Container;
    property Document: IDispatch read Get_Document;
    property TopLevelContainer: WordBool read Get_TopLevelContainer;
    property type_: WideString read Get_type_;
    property Left: Integer read Get_Left write Set_Left;
    property Top: Integer read Get_Top write Set_Top;
    property Width: Integer read Get_Width write Set_Width;
    property Height: Integer read Get_Height write Set_Height;
    property LocationName: WideString read Get_LocationName;
    property LocationURL: WideString read Get_LocationURL;
    property Busy: WordBool read Get_Busy;
  end;

// *********************************************************************//
// Interface: IWebBrowserApp
// Flags:     (4432) Hidden Dual OleAutomation Dispatchable
// GUID:      {0002DF05-0000-0000-C000-000000000046}
// *********************************************************************//
  IWebBrowserApp = interface(IWebBrowser)
    ['{0002DF05-0000-0000-C000-000000000046}']
    procedure Quit; safecall;
    procedure ClientToWindow(var pcx: SYSINT; var pcy: SYSINT); safecall;
    procedure PutProperty(const Property_: WideString; vtValue: OleVariant); safecall;
    function GetProperty(const Property_: WideString): OleVariant; safecall;
    function Get_Name: WideString; safecall;
    function Get_HWND: HWND; safecall;
    function Get_FullName: WideString; safecall;
    function Get_Path: WideString; safecall;
    function Get_Visible: WordBool; safecall;
    procedure Set_Visible(pBool: WordBool); safecall;
    function Get_StatusBar: WordBool; safecall;
    procedure Set_StatusBar(pBool: WordBool); safecall;
    function Get_StatusText: WideString; safecall;
    procedure Set_StatusText(const StatusText: WideString); safecall;
    function Get_ToolBar: SYSINT; safecall;
    procedure Set_ToolBar(Value: SYSINT); safecall;
    function Get_MenuBar: WordBool; safecall;
    procedure Set_MenuBar(Value: WordBool); safecall;
    function Get_FullScreen: WordBool; safecall;
    procedure Set_FullScreen(pbFullScreen: WordBool); safecall;
    property Name: WideString read Get_Name;
    property HWND: HWND read Get_HWND;
    property FullName: WideString read Get_FullName;
    property Path: WideString read Get_Path;
    property Visible: WordBool read Get_Visible write Set_Visible;
    property StatusBar: WordBool read Get_StatusBar write Set_StatusBar;
    property StatusText: WideString read Get_StatusText write Set_StatusText;
    property ToolBar: SYSINT read Get_ToolBar write Set_ToolBar;
    property MenuBar: WordBool read Get_MenuBar write Set_MenuBar;
    property FullScreen: WordBool read Get_FullScreen write Set_FullScreen;
  end;
// *********************************************************************//
// Interface: IWebBrowser2
// Flags:     (4432) Hidden Dual OleAutomation Dispatchable
// GUID:      {D30C1661-CDAF-11D0-8A3E-00C04FC9E26E}
// *********************************************************************//
  IWebBrowser2 = interface(IWebBrowserApp)
    ['{D30C1661-CDAF-11D0-8A3E-00C04FC9E26E}']
    procedure Navigate2(const URL: OleVariant; const Flags: OleVariant;
                        const TargetFrameName: OleVariant; const PostData: OleVariant;
                        const Headers: OleVariant); safecall;
    function QueryStatusWB(cmdID: OLECMDID): OLECMDF; safecall;
    procedure ExecWB(cmdID: OLECMDID; cmdexecopt: OLECMDEXECOPT; const pvaIn: OleVariant;
                     var pvaOut: OleVariant); safecall;
    procedure ShowBrowserBar(const pvaClsid: OleVariant; const pvarShow: OleVariant;
                             const pvarSize: OleVariant); safecall;
    function Get_ReadyState: tagREADYSTATE; safecall;
    function Get_Offline: WordBool; safecall;
    procedure Set_Offline(pbOffline: WordBool); safecall;
    function Get_Silent: WordBool; safecall;
    procedure Set_Silent(pbSilent: WordBool); safecall;
    function Get_RegisterAsBrowser: WordBool; safecall;
    procedure Set_RegisterAsBrowser(pbRegister: WordBool); safecall;
    function Get_RegisterAsDropTarget: WordBool; safecall;
    procedure Set_RegisterAsDropTarget(pbRegister: WordBool); safecall;
    function Get_TheaterMode: WordBool; safecall;
    procedure Set_TheaterMode(pbRegister: WordBool); safecall;
    function Get_AddressBar: WordBool; safecall;
    procedure Set_AddressBar(Value: WordBool); safecall;
    function Get_Resizable: WordBool; safecall;
    procedure Set_Resizable(Value: WordBool); safecall;
    property ReadyState: tagREADYSTATE read Get_ReadyState;
    property Offline: WordBool read Get_Offline write Set_Offline;
    property Silent: WordBool read Get_Silent write Set_Silent;
    property RegisterAsBrowser: WordBool read Get_RegisterAsBrowser write Set_RegisterAsBrowser;
    property RegisterAsDropTarget: WordBool read Get_RegisterAsDropTarget write Set_RegisterAsDropTarget;
    property TheaterMode: WordBool read Get_TheaterMode write Set_TheaterMode;
    property AddressBar: WordBool read Get_AddressBar write Set_AddressBar;
    property Resizable: WordBool read Get_Resizable write Set_Resizable;
  end;

// *********************************************************************//
// DispIntf:  DWebBrowserEvents2
// Flags:     (4112) Hidden Dispatchable
// GUID:      {34A715A0-6587-11D0-924A-0020AFC7AC4D}
// *********************************************************************//
  DWebBrowserEvents2 = dispinterface
    ['{34A715A0-6587-11D0-924A-0020AFC7AC4D}']
    procedure StatusTextChange(const Text: WideString); dispid 102;
    procedure ProgressChange(Progress: Integer; ProgressMax: Integer); dispid 108;
    procedure CommandStateChange(Command: Integer; Enable: WordBool); dispid 105;
    procedure DownloadBegin; dispid 106;
    procedure DownloadComplete; dispid 104;
    procedure TitleChange(const Text: WideString); dispid 113;
    procedure PropertyChange(const szProperty: WideString); dispid 112;
    procedure BeforeNavigate2(const pDisp: IDispatch; const URL: OleVariant;
                              const Flags: OleVariant; const TargetFrameName: OleVariant;
                              const PostData: OleVariant; const Headers: OleVariant;
                              var Cancel: WordBool); dispid 250;
    procedure NewWindow2(var ppDisp: IDispatch; var Cancel: WordBool); dispid 251;
    procedure NavigateComplete2(const pDisp: IDispatch; const URL: OleVariant); dispid 252;
    procedure DocumentComplete(const pDisp: IDispatch; const URL: OleVariant); dispid 259;
    procedure OnQuit; dispid 253;
    procedure OnVisible(Visible: WordBool); dispid 254;
    procedure OnToolBar(ToolBar: WordBool); dispid 255;
    procedure OnMenuBar(MenuBar: WordBool); dispid 256;
    procedure OnStatusBar(StatusBar: WordBool); dispid 257;
    procedure OnFullScreen(FullScreen: WordBool); dispid 258;
    procedure OnTheaterMode(TheaterMode: WordBool); dispid 260;
    procedure WindowSetResizable(Resizable: WordBool); dispid 262;
    procedure WindowSetLeft(Left: Integer); dispid 264;
    procedure WindowSetTop(Top: Integer); dispid 265;
    procedure WindowSetWidth(Width: Integer); dispid 266;
    procedure WindowSetHeight(Height: Integer); dispid 267;
    procedure WindowClosing(IsChildWindow: WordBool; var Cancel: WordBool); dispid 263;
    procedure ClientToHostWindow(var CX: Integer; var CY: Integer); dispid 268;
    procedure SetSecureLockIcon(SecureLockIcon: Integer); dispid 269;
    procedure FileDownload(ActiveDocument: WordBool; var Cancel: WordBool); dispid 270;
    procedure NavigateError(const pDisp: IDispatch; const URL: OleVariant; const Frame: OleVariant;
                            const StatusCode: OleVariant; var Cancel: WordBool); dispid 271;
    procedure PrintTemplateInstantiation(const pDisp: IDispatch); dispid 225;
    procedure PrintTemplateTeardown(const pDisp: IDispatch); dispid 226;
    procedure UpdatePageStatus(const pDisp: IDispatch; const nPage: OleVariant;
                               const fDone: OleVariant); dispid 227;
    procedure PrivacyImpactedStateChange(bImpacted: WordBool); dispid 272;
    procedure NewWindow3(var ppDisp: IDispatch; var Cancel: WordBool; dwFlags: LongWord;
                         const bstrUrlContext: WideString; const bstrUrl: WideString); dispid 273;
    procedure SetPhishingFilterStatus(PhishingFilterStatus: Integer); dispid 282;
    procedure WindowStateChanged(dwWindowStateFlags: LongWord; dwValidFlagsMask: LongWord); dispid 283;
    procedure NewProcess(lCauseFlag: Integer; const pWB2: IDispatch; var Cancel: WordBool); dispid 284;
    procedure ThirdPartyUrlBlocked(const URL: OleVariant; dwCount: LongWord); dispid 285;
    procedure RedirectXDomainBlocked(const pDisp: IDispatch; const StartURL: OleVariant;
                                     const RedirectURL: OleVariant; const Frame: OleVariant;
                                     const StatusCode: OleVariant); dispid 286;
    procedure BeforeScriptExecute(const pDispWindow: IDispatch); dispid 290;
    procedure WebWorkerStarted(dwUniqueID: LongWord; const bstrWorkerLabel: WideString); dispid 288;
    procedure WebWorkerFinsihed(dwUniqueID: LongWord); dispid 289;
  end;


  TDocHostUIInfo = record
    cbSize: ULONG;
    dwFlags: DWORD;
    dwDoubleClick: DWORD;
    pchHostCss: PWChar;
    pchHostNS: PWChar;
  end;

  IDocHostUIHandler  = interface(IUnknown)
    ['{BD3F23C0-D43E-11CF-893B-00AA00BDCE1A}']
    function ShowContextMenu(const dwID: DWORD; const ppt: PPOINT; const pcmdtReserved: IUnknown; const pdispReserved: IDispatch): HRESULT; stdcall;
    function GetHostInfo(var pInfo: TDocHostUIInfo): HRESULT; stdcall;
    function ShowUI(const dwID: DWORD; const pActiveObject: IOleInPlaceActiveObject; const pCommandTarget: IOleCommandTarget; const pFrame: IOleInPlaceFrame; const pDoc: IOleInPlaceUIWindow): HRESULT; stdcall;
    function HideUI: HRESULT; stdcall;
    function UpdateUI: HRESULT; stdcall;
    function EnableModeless(const fEnable: BOOL): HRESULT; stdcall;
    function OnDocWindowActivate(const fActivate: BOOL): HRESULT; stdcall;
    function OnFrameWindowActivate(const fActivate: BOOL): HRESULT; stdcall;
    function ResizeBorder(const prcBorder: PRECT; const pUIWindow: IOleInPlaceUIWindow; const FrameWindow: BOOL): HRESULT; stdcall;
    function TranslateAccelerator(const lpMsg: PMSG; const pguidCmdGroup:  PGUID; const nCmdID: DWORD): HRESULT; stdcall;
    function GetOptionKeyPath(var pchKey: POLESTR; const dw: DWORD): HRESULT; stdcall;
    function GetDropTarget(const pDropTarget: IDropTarget; out ppDropTarget: IDropTarget): HRESULT; stdcall;
    function GetExternal(out ppDispatch: IDispatch): HRESULT; stdcall;
    function TranslateUrl(const dwTranslate: DWORD; const pchURLIn: POLESTR; var ppchURLOut: POLESTR): HRESULT; stdcall;
    function FilterDataObject(const pDO: IDataObject; out ppDORet: IDataObject): HRESULT; stdcall;
  end;


  TJsExternal = class(TAutoIntfObject, IWebJSExternal, IDispatch)
  private
    FWebForm: Pointer;
  protected
    function JSExternal(const func: WideString; const args: WideString): WideString; safecall;
  public
    constructor Create(AWebForm: Pointer);
    destructor Destroy; override;
  end;

  { TMiniWebview }

  TMiniWebview = class(TCustomControl, IDispatch, IOleClientSite, IOleControlSite, IOleInPlaceSite, IOleInPlaceFrame, IDocHostUIHandler)
  private
    FWebBrowser: IWebBrowser2;
    FOleInPlaceObject: IOleInPlaceObject;
    FOleInPlaceActiveObject: IOleInPlaceActiveObject;
    FOleObject: IOleObject;
    FURL: string;
    FConnected: Boolean;
    FDwCookie: {$IFDEF FPC} DWord {$ELSE} Integer {$ENDIF};
    FConPpoint: IConnectionPoint;
    FExternalObj: IDispatch;

    FOnStatusTextChange: TWebStatusTextChangeEvent;
    FOnWindowClosing: TWebWindowClosingEvent;
    FOnOnDocumentComplete: TWebDocumentCompleteEvent;
    FOnTitleChange: TWebTitleChangeEvent;
    FOnJSExternal: TWebJSExternalEvent;

    function GetReadyState: TReadyState;
    function GetMainMenu: TMainMenu;
    procedure DisconnectPpoint;
  protected
    procedure VisibleChanging; override;
    // IOleClientSite
    function SaveObject: HResult; stdcall;
    function GetMoniker(dwAssign: Longint; dwWhichMoniker: Longint;
      out mk: IMoniker): HResult; stdcall;
    function GetContainer(out container: IOleContainer): HResult; stdcall;
    function ShowObject: HResult; stdcall;
    function OnShowWindow(fShow: BOOL): HResult; stdcall;
    function RequestNewObjectLayout: HResult; stdcall;

    // IOleInPlaceSite
    function CanInPlaceActivate: HResult; stdcall;
    function OnInPlaceActivate: HResult; stdcall;
    function OnUIActivate: HResult; stdcall;
  {$IFDEF FPC}
    function GetWindowContext(out ppframe:IOleInPlaceFrame;out ppdoc:IOleInPlaceUIWindow;
      lprcposrect:LPRECT;lprccliprect:LPRECT;lpframeinfo:LPOLEINPLACEFRAMEINFO):hresult; stdcall;
    function Scroll(scrollExtant:TSIZE):hresult; stdcall;
    function OnPosRectChange(lprcPosRect:LPRect):hresult; stdcall;
  {$ELSE}
    function GetWindowContext(out frame: IOleInPlaceFrame;
      out doc: IOleInPlaceUIWindow; out rcPosRect: TRect;
      out rcClipRect: TRect; out frameInfo: TOleInPlaceFrameInfo): HResult;
      stdcall;
    function Scroll(scrollExtent: TPoint): HResult; stdcall;
    function OnPosRectChange(const rcPosRect: TRect): HResult; stdcall;
  {$ENDIF}

    function OnUIDeactivate(fUndoable: BOOL): HResult; stdcall;
    function OnInPlaceDeactivate: HResult; stdcall;
    function DiscardUndoState: HResult; stdcall;
    function DeactivateAndUndo: HResult; stdcall;


    // IOleWindow
//    function GetWindow(out wnd: HWnd): HResult; stdcall;
    function IOleInPlaceSite.GetWindow = OleInPlaceSite_GetWindow;
    function OleInPlaceSite_GetWindow(out wnd: HWnd): HResult; stdcall;

    function ContextSensitiveHelp(fEnterMode: BOOL): HResult; stdcall;

    // IOleControlSite
    function OnControlInfoChanged: HResult; stdcall;
    function LockInPlaceActive(fLock: BOOL): HResult; stdcall;
    function GetExtendedControl(out disp: IDispatch): HResult; stdcall;

  {$IFDEF FPC}
    function TransformCoords(var ptlHimetric: _POINTL; var ptfContainer: tagPOINTF; flags: LongWord): HResult; stdcall;
  {$ELSE}
    function TransformCoords(var ptlHimetric: TPoint; var ptfContainer: TPointF; flags: Longint): HResult; stdcall;
  {$ENDIF}
    function IOleControlSite.TranslateAccelerator = OleControlSite_TranslateAccelerator;
  {$IFDEF FPC}
    function OleControlSite_TranslateAccelerator(var pMsg:tagMSG;grfModifiers:LongWord):HRESULT;stdcall;
  {$ELSE}
    function OleControlSite_TranslateAccelerator(msg: PMsg; grfModifiers: Longint): HResult; stdcall;
  {$ENDIF}
    function OnFocus(fGotFocus: BOOL): HResult; stdcall;
    function ShowPropertyFrame: HResult; stdcall;

    // IOleInPlaceFrame
    function InsertMenus(hmenuShared: HMenu;
      var menuWidths: TOleMenuGroupWidths): HResult; stdcall;
    function SetMenu(hmenuShared: HMenu; holemenu: HMenu;
      hwndActiveObject: HWnd): HResult; stdcall;
    function RemoveMenus(hmenuShared: HMenu): HResult; stdcall;
    function SetStatusText(pszStatusText: POleStr): HResult; stdcall;
    function EnableModeless(fEnable: BOOL): HResult; stdcall;
    function IOleInPlaceFrame.TranslateAccelerator = OleInPlaceFrame_TranslateAccelerator;
    function OleInPlaceFrame_TranslateAccelerator(var msg: TMsg;
      wID: Word): HResult; stdcall;

    function IOleInPlaceFrame.GetWindow = OleInPlaceFrame_GetWindow;
    function OleInPlaceFrame_GetWindow(out wnd: HWnd): HResult; stdcall;

    // IOleInPlaceUIWindow
    function GetBorder(out rectBorder: TRect): HResult; stdcall;
    function RequestBorderSpace(const borderwidths: TRect): HResult; stdcall;
  {$IFDEF FPC}
    function SetBorderSpace(const borderwidths: TRect):HResult;StdCall;
  {$ELSE}
    function SetBorderSpace(pborderwidths: PRect): HResult; stdcall;
  {$ENDIF}
    function SetActiveObject(const activeObject: IOleInPlaceActiveObject;
      pszObjName: POleStr): HResult; stdcall;


    // IDocHostUIHandler
      function ShowContextMenu(const dwID: DWORD; const ppt: PPOINT; const pcmdtReserved: IUnknown; const pdispReserved: IDispatch): HRESULT; stdcall;
      function GetHostInfo(var pInfo: TDocHostUIInfo): HRESULT; stdcall;
      function ShowUI(const dwID: DWORD; const pActiveObject: IOleInPlaceActiveObject; const pCommandTarget: IOleCommandTarget; const pFrame: IOleInPlaceFrame; const pDoc: IOleInPlaceUIWindow): HRESULT; stdcall;
      function HideUI: HRESULT; stdcall;
      function UpdateUI: HRESULT; stdcall;

      function IDocHostUIHandler_EnableModeless(const fEnable: BOOL): HRESULT; stdcall;
      function IDocHostUIHandler.EnableModeless = IDocHostUIHandler_EnableModeless;

      function OnDocWindowActivate(const fActivate: BOOL): HRESULT; stdcall;
      function OnFrameWindowActivate(const fActivate: BOOL): HRESULT; stdcall;
      function ResizeBorder(const prcBorder: PRECT; const pUIWindow: IOleInPlaceUIWindow; const FrameWindow: BOOL): HRESULT; stdcall;
      function TranslateAccelerator(const lpMsg: PMSG; const pguidCmdGroup:  PGUID; const nCmdID: DWORD): HRESULT; stdcall;
      function GetOptionKeyPath(var pchKey: POLESTR; const dw: DWORD): HRESULT; stdcall;
      function GetDropTarget(const pDropTarget: IDropTarget; out ppDropTarget: IDropTarget): HRESULT; stdcall;
      function GetExternal(out ppDispatch: IDispatch): HRESULT; stdcall;
      function TranslateUrl(const dwTranslate: DWORD; const pchURLIn: POLESTR; var ppchURLOut: POLESTR): HRESULT; stdcall;
      function FilterDataObject(const pDO: IDataObject; out ppDORet: IDataObject): HRESULT; stdcall;



    procedure SetUIActive(Active: Boolean);
    procedure Resize; override;
    procedure SetParent(AParent: TWinControl); override;

    function Invoke(DispID: Integer; const IID: TGUID; LocaleID: Integer;
      Flags: Word; var Params; VarResult, ExcepInfo, ArgErr: Pointer): HResult; stdcall;

    procedure UpdateWebbrowserSize;
    procedure Paint; override;
  public
     procedure CreateWnd; override;
     procedure WndProc(var Message: TMessage); override;
  public
    constructor Create(AOwner: TComponent); override;
    destructor Destroy; override;
    procedure Navigate(const AURL: string);
    procedure GoBack;
    procedure GoForward;
    procedure GoHome;
    procedure GoSearch;
    procedure Refresh;
    procedure Stop;
    procedure SetBounds(ALeft, ATop, AWidth, AHeight: Integer); override;

    procedure ExecuteScript(AScriptText, AScriptType: string);
    procedure ExecuteJS(AScriptText: string);

    procedure LoadHTML(const AStr: string);

    // propertys
    property ReadyState: TReadyState read GetReadyState;
  published
    property Align;
    property Anchors;
  {$ifdef fpc}
    property BorderSpacing;
    property ChildSizing;
  {$endif}
    property Constraints;
    property Enabled;
    property Visible;
    // events
    property OnStatusTextChange: TWebStatusTextChangeEvent read FOnStatusTextChange write FOnStatusTextChange;
    property OnTitleChange: TWebTitleChangeEvent read FOnTitleChange write FOnTitleChange;
    property OnDocumentComplete: TWebDocumentCompleteEvent read FOnOnDocumentComplete write FOnOnDocumentComplete;
    property OnWindowClosing: TWebWindowClosingEvent read FOnWindowClosing write FOnWindowClosing;
    property OnJSExternal: TWebJSExternalEvent read FOnJSExternal write FOnJSExternal;
  end;

var
  // 特殊用处，用于消息处理，没有浏览器就不管去处理
  GWebviewCount: Integer = 0;

implementation

const
  CLSID_WebBrowser: TGUID = '{8856F961-340A-11D0-A96B-00C04FD705A2}';
  IID_IOleObject: TGUID = '{00000112-0000-0000-C000-000000000046}';
  IID_IWebBrowser2: TGUID = '{D30C1661-CDAF-11D0-8A3E-00C04FC9E26E}';

{$R WebBrowser.tlb}

{ TJsExternal }

constructor TJsExternal.Create(AWebForm: Pointer);
var
  TypeLib: ITypeLib;
  ExeName: WideString;
begin
  ExeName := GetModuleName(HInstance);
  FWebForm := AWebForm;
  OleCheck(LoadTypeLib(PWideChar(ExeName), TypeLib));
  inherited Create(TypeLib, IWebJSExternal);
end;

destructor TJsExternal.Destroy;
begin
  inherited;
end;

function TJsExternal.JSExternal(const func: WideString; const args: WideString): WideString; safecall;
begin
  Result := EmptyWideStr;

  if Assigned(FWebForm) and Assigned(TMiniWebview(FWebForm).OnJSExternal) then
    TMiniWebview(FWebForm).OnJSExternal(FWebForm, func, args, Result);
end;

{ TMiniWebview }

constructor TMiniWebview.Create(AOwner: TComponent);
begin
  inherited Create(AOwner);
  ControlStyle := ControlStyle - [csAcceptsControls, csSetCaption];
  Width := 300;
  Height := 200;
  if not(csDesigning in ComponentState) then
  begin
    OleCheck(CoCreateInstance(CLSID_WebBrowser, nil, CLSCTX_INPROC, IID_IOleObject, FOleObject));
    FExternalObj := TJsExternal.Create(Self);
    Inc(GWebviewCount);
  end;
end;


// 没有考虑设计时的，所以尽量不要安装成控件使用
procedure TMiniWebview.CreateWnd;
var
  con: IConnectionPointContainer;
begin
  inherited CreateWnd;
  if not(csDesigning in ComponentState) then
  begin
    if FOleObject <> nil then
    begin
      DisconnectPpoint;
      OnInPlaceDeactivate;
      FOleObject.Close(OLECLOSE_NOSAVE);
      FOleObject.SetClientSite(nil);
      FWebBrowser := nil;
    end;
    if FWebBrowser = nil then
    begin
      OleCheck(FOleObject.SetClientSite(Self));
      OleCheck(FOleObject.DoVerb(OLEIVERB_INPLACEACTIVATE, nil, Self, 0, Handle, PRect(nil)^));
      OleCheck(FOleObject.QueryInterface(IID_IWebBrowser2, FWebBrowser));
      FWebBrowser.Set_Silent(true);
      FConnected := Succeeded(FWebBrowser.QueryInterface(IConnectionPointContainer, con)) and
                    Succeeded(con.FindConnectionPoint(DWebBrowserEvents2, FConPpoint))  and
                    Succeeded(FConPpoint.Advise(Self, FDwCookie));
      con := nil;
      if FURL <> '' then
        Navigate(FURL);
    end;
  end;
end;

// Vcl.OleCtrls
procedure TMiniWebview.WndProc(var Message: TMessage);
var
  WinMsg: TMsg;
begin
  if (Message.Msg >= CN_BASE + WM_KEYFIRST) and
    (Message.Msg <= CN_BASE + WM_KEYLAST) and
    (FOleInPlaceActiveObject <> nil) then
  begin
    WinMsg.HWnd := Handle;
    WinMsg.Message := Message.Msg - CN_BASE;
    WinMsg.WParam := Message.WParam;
    WinMsg.LParam := Message.LParam;
    WinMsg.Time := GetMessageTime;
    WinMsg.Pt.X := $115DE1F1;
    WinMsg.Pt.Y := $115DE1F1;
    if FOleInPlaceActiveObject.TranslateAccelerator(WinMsg) = S_OK then
    begin
      Message.Result := 1;
      Exit;
    end;
  end;
  inherited WndProc(Message);
end;

destructor TMiniWebview.Destroy;
begin
  if not(csDesigning in ComponentState) then
  begin
    DisconnectPpoint;
    if FOleObject <> nil then
    begin
      FOleObject.Close(OLECLOSE_NOSAVE);
      FOleObject.SetClientSite(nil);
    end;
    FExternalObj := nil;
    FOleObject := nil;
    FWebBrowser := nil;
    Dec(GWebviewCount);
  end;
  inherited;
end;


procedure TMiniWebview.Navigate(const AURL: string);
var
  LPostData: OleVariant;
begin
  FURL := AURL;
  LPostData := EmptyParam;
  if FWebBrowser <> nil then
    FWebBrowser.Navigate(FURL, EmptyParam, EmptyParam, LPostData, EmptyParam);
end;

procedure TMiniWebview.GoBack;
begin
  if Assigned(FWebBrowser) then
  try
    FWebBrowser.GoBack;
  except
  end;
end;

procedure TMiniWebview.GoForward;
begin
  if Assigned(FWebBrowser) then
  try
    FWebBrowser.GoForward;
  except
  end;
end;

procedure TMiniWebview.GoHome;
begin
  if Assigned(FWebBrowser) then
  try
    FWebBrowser.GoHome;
  except
  end;
end;

procedure TMiniWebview.GoSearch;
begin
  if Assigned(FWebBrowser) then
  try
    FWebBrowser.GoSearch;
  except
  end;
end;

procedure TMiniWebview.Refresh;
begin
  if Assigned(FWebBrowser) then
  try
    FWebBrowser.Refresh;
  except
  end;
end;

procedure TMiniWebview.Stop;
begin
  if Assigned(FWebBrowser) then
  try
    FWebBrowser.Stop;
  except
  end;
end;

procedure TMiniWebview.SetParent(AParent: TWinControl);
begin
  inherited SetParent(AParent);
end;

procedure TMiniWebview.Resize;
begin
  inherited;
end;

//----------------------------------------------------------------------

function TMiniWebview.Invoke(DispID: Integer; const IID: TGUID; LocaleID: Integer;
  Flags: Word; var Params; VarResult, ExcepInfo, ArgErr: Pointer): HResult; stdcall;
var
  LVarParams : array of OleVariant;
  I: Integer;
  LPVarArgIn: PVariantArg;
  LvPDispParams: TDispParams;
  LB: Boolean;
begin
  LvPDispParams := PDispParams(@Params)^;
  SetLength(LVarParams, LvPDispParams.cArgs);

  if LvPDispParams.cNamedArgs > 0 then
    for I := Low(LVarParams) to High(LVarParams) do
    begin
      LPVarArgIn := @LvPDispParams.rgvarg[i];
      LVarParams[LvPDispParams.rgdispidNamedArgs[i]] := POleVariant(LPVarArgIn)^;
    end
  else
  begin
    for I := Low(LVarParams) to High(LVarParams) do
    begin
      LPVarArgIn := @LvPDispParams.rgvarg[I];
      LVarParams[High(LVarParams) - I] := POleVariant(LPVarArgIn)^;
    end;
  end;


  case DispId of
    102:  // StatusTextChange
     begin
       if Assigned(FOnStatusTextChange) then
         FOnStatusTextChange(Self, LVarParams[0]);
     end;
    113:   // TitleChange
      begin
       if Assigned(FOnTitleChange) then
         FOnTitleChange(Self, LVarParams[0]);
      end;

    259:  // DocumentComplete
      begin
//        if Assigned(FOnOnDocumentComplete) then
//          FOnOnDocumentComplete(Self, (TVarData(VarAsType(Params[0], varAny)).varAny, Params[1]);
      end;

    263: // WindowClosing
      begin
        if Assigned(FOnWindowClosing) then
        begin
          LB := LVarParams[1];
          FOnWindowClosing(Self, LVarParams[0], LB);
          LVarParams[1] := LB;
        end;
      end;
  end;
  SetLength(LVarParams, 0);
  Result := S_OK;

  //procedure TitleChange(const Text: WideString); dispid 113;
  //procedure NewWindow3(var ppDisp: IDispatch; var Cancel: WordBool; dwFlags: LongWord;
  // const bstrUrlContext: WideString; const bstrUrl: WideString); dispid 273;
  //procedure NavigateComplete2(const pDisp: IDispatch; const URL: OleVariant); dispid 252;
  //procedure DocumentComplete(const pDisp: IDispatch; const URL: OleVariant); dispid 259;
  //procedure StatusTextChange(const Text: WideString); dispid 102;
  //procedure WindowClosing(IsChildWindow: WordBool; var Cancel: WordBool); dispid 263;

end;

procedure TMiniWebview.UpdateWebbrowserSize;
begin
  if FWebBrowser <> nil then
   begin
     FWebBrowser.Left := 0;
     FWebBrowser.Top := 0;
     FWebBrowser.Width := Self.Width;
     FWebBrowser.Height := Self.Height;
   end;
end;

procedure TMiniWebview.Paint;
begin
  inherited Paint;
  if csDesigning in ComponentState then
  begin
    Canvas.Brush.Color:=clWhite;
    Canvas.FillRect(ClientRect);
  end;
end;


function TMiniWebview.LockInPlaceActive(fLock: BOOL): HResult; stdcall;
begin
  Result := E_NOTIMPL;
end;

function TMiniWebview.CanInPlaceActivate: HResult; stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.ContextSensitiveHelp(fEnterMode: BOOL): HResult; stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.DeactivateAndUndo: HResult; stdcall;
begin
  FOleInPlaceObject.UIDeactivate;
  Result := S_OK;
end;

function TMiniWebview.DiscardUndoState: HResult; stdcall;
begin
  Result := E_NOTIMPL;
end;

procedure TMiniWebview.DisconnectPpoint;
begin
  if FConnected then
  begin
    if  Assigned(FConPpoint) then
      FConPpoint.Unadvise(FDwCookie);
    FConPpoint := nil;
    FConnected := False;
  end;
end;

function TMiniWebview.EnableModeless(fEnable: BOOL): HResult; stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.GetBorder(out rectBorder: TRect): HResult; stdcall;
begin
  Result := INPLACE_E_NOTOOLSPACE;
end;

function TMiniWebview.GetContainer(out container: IOleContainer): HResult;
  stdcall;
begin
  Result := E_NOINTERFACE;
end;

function TMiniWebview.GetExtendedControl(out disp: IDispatch): HResult; stdcall;
begin
  Result := E_NOTIMPL;
end;

function TMiniWebview.GetMainMenu: TMainMenu;
var
  Form: TCustomForm;
begin
  Result := nil;
  Form := GetParentForm(Self);
  if Form <> nil then
    if (Form is TForm) and (TForm(Form).FormStyle <> fsMDIChild) then
      Result := Form.Menu
    else
      if Application.MainForm <> nil then
        Result := Application.MainForm.Menu;
end;

function TMiniWebview.GetMoniker(dwAssign: Longint; dwWhichMoniker: Longint;
  out mk: IMoniker): HResult; stdcall;
begin
  Result := E_NOTIMPL;
end;

function TMiniWebview.GetReadyState: TReadyState;
begin
  Result := rsUninitialized;
  if Assigned(FWebBrowser) then
    Result := TReadyState(FWebBrowser.ReadyState);
end;

function TMiniWebview.OleInPlaceSite_GetWindow(out wnd: HWnd): HResult; stdcall;
begin
  Result := S_OK;
  wnd := Handle;
  if wnd = 0 then Result := E_FAIL;
end;

{$IFDEF FPC}
function TMiniWebview.GetWindowContext(out ppframe: IOleInPlaceFrame; out
  ppdoc: IOleInPlaceUIWindow; lprcposrect: LPRECT; lprccliprect: LPRECT;
  lpframeinfo: LPOLEINPLACEFRAMEINFO): hresult; stdcall;
begin
  ppframe := Self;
  ppdoc := nil;
  lprcposrect^ := BoundsRect;
  SetRect(lprccliprect^, 0, 0, 32767, 32767);
  with lpframeinfo^ do
  begin
    fMDIApp := False;
    hWndFrame := Handle;
    hAccel := 0;
    cAccelEntries := 0;
  end;
  Result := S_OK;
end;
{$ELSE}
function TMiniWebview.GetWindowContext(out frame: IOleInPlaceFrame;
    out doc: IOleInPlaceUIWindow; out rcPosRect: TRect;
    out rcClipRect: TRect; out frameInfo: TOleInPlaceFrameInfo): HResult;
    stdcall;
begin
  frame := Self;
  doc := nil;
  rcPosRect := BoundsRect;
  SetRect(rcClipRect, 0, 0, 32767, 32767);
  with frameInfo do
  begin
    fMDIApp := False;
    hWndFrame := Handle;
    hAccel := 0;
    cAccelEntries := 0;
  end;
  Result := S_OK;
end;
{$ENDIF}


function TMiniWebview.InsertMenus(hmenuShared: HMenu;
  var menuWidths: TOleMenuGroupWidths): HResult; stdcall;
{$IFDEF FPC}
begin
  Result := E_NOTIMPL;
end;
{$ELSE}
var
  Menu: TMainMenu;
begin
  Menu := GetMainMenu;
  if Menu <> nil then
    Menu.PopulateOle2Menu(hmenuShared, [0, 2, 4], menuWidths.width);
  Result := S_OK;
end;
{$ENDIF}

{$IFDEF FPC}
function TMiniWebview.OleControlSite_TranslateAccelerator(var pMsg: tagMSG;
  grfModifiers: LongWord): HRESULT; stdcall;
{$ELSE}
function TMiniWebview.OleControlSite_TranslateAccelerator(msg: PMsg; grfModifiers: Longint): HResult;
{$ENDIF}
begin
  Result := E_NOTIMPL;
end;

function TMiniWebview.OleInPlaceFrame_GetWindow(out wnd: HWnd): HResult;
  stdcall;
begin
  wnd := GetTopParentHandle;
  Result := S_OK;
end;

function TMiniWebview.OleInPlaceFrame_TranslateAccelerator(var msg: TMsg;
  wID: Word): HResult; stdcall;
begin
  Result := S_FALSE;
end;

function TMiniWebview.OnControlInfoChanged: HResult; stdcall;
begin
  Result := E_NOTIMPL;
end;

function TMiniWebview.OnFocus(fGotFocus: BOOL): HResult; stdcall;
begin
  Result := E_NOTIMPL;
end;

function TMiniWebview.OnInPlaceActivate: HResult; stdcall;
begin
  FOleObject.QueryInterface(IOleInPlaceObject, FOleInPlaceObject);
  FOleObject.QueryInterface(IOleInPlaceActiveObject, FOleInPlaceActiveObject);
  Result := S_OK;
end;

function TMiniWebview.OnInPlaceDeactivate: HResult; stdcall;
begin
  FOleInPlaceObject := nil;
  FOleInPlaceActiveObject := nil;
  Result := S_OK;
end;

{$IFDEF FPC}
function TMiniWebview.OnPosRectChange(lprcPosRect: LPRect): hresult; stdcall;
var
  LR: TRect;
begin
  LR := Rect(0, 0, 32767, 32767);
  FOleInPlaceObject.SetObjectRects(lprcPosRect, @LR);
  Result := S_OK;
end;
{$ELSE}
function TMiniWebview.OnPosRectChange(const rcPosRect: TRect): HResult;
begin
  FOleInPlaceObject.SetObjectRects(rcPosRect, Rect(0, 0, 32767, 32767));
  Result := S_OK;
end;
{$ENDIF}



function TMiniWebview.OnShowWindow(fShow: BOOL): HResult; stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.OnUIActivate: HResult; stdcall;
begin
  SetUIActive(True);
  Result := S_OK;
end;

function TMiniWebview.OnUIDeactivate(fUndoable: BOOL): HResult; stdcall;
begin
  SetMenu(0, 0, 0);
  SetUIActive(False);
  Result := S_OK;
end;

function TMiniWebview.RemoveMenus(hmenuShared: HMenu): HResult; stdcall;
begin
  while GetMenuItemCount(hmenuShared) > 0 do
    RemoveMenu(hmenuShared, 0, MF_BYPOSITION);
  Result := S_OK;
end;

function TMiniWebview.RequestBorderSpace(const borderwidths: TRect): HResult;
  stdcall;
begin
  Result := INPLACE_E_NOTOOLSPACE;
end;

function TMiniWebview.RequestNewObjectLayout: HResult; stdcall;
begin
  Result := E_NOTIMPL;
end;

function TMiniWebview.SaveObject: HResult; stdcall;
begin
  Result := S_OK;
end;

{$IFDEF FPC}
function TMiniWebview.Scroll(scrollExtant: TSIZE): hresult; stdcall;
{$ELSE}
function TMiniWebview.Scroll(scrollExtent: TPoint): HResult;
{$ENDIF}
begin
  Result := E_NOTIMPL;
end;

function TMiniWebview.SetActiveObject(
  const activeObject: IOleInPlaceActiveObject; pszObjName: POleStr): HResult;
  stdcall;
begin
  Result := S_OK;
end;

{$IFDEF FPC}
function TMiniWebview.SetBorderSpace(const borderwidths: TRect):HResult;StdCall;
{$ELSE}
function TMiniWebview.SetBorderSpace(pborderwidths: PRect): HResult;
{$ENDIF}
begin
  Result := E_NOTIMPL;
end;

procedure TMiniWebview.SetBounds(ALeft, ATop, AWidth, AHeight: Integer);
begin
  inherited SetBounds(ALeft, ATop, AWidth, AHeight);
  if not(csDesigning in ComponentState) then
    Self.UpdateWebbrowserSize;
end;

procedure TMiniWebview.ExecuteScript(AScriptText, AScriptType: string);
var
  LDocument: IHTMLDocument2;
  LParent: IHTMLWindow2;
begin
  if not HandleAllocated then
    HandleNeeded;
  if FWebBrowser = nil then
    Exit;
  try
    LDocument := FWebBrowser.Document as IHTMLDocument2;
    if Assigned(LDocument) then
    begin
      LParent := LDocument.parentWindow;
      if Assigned(LParent) then
        LParent.execScript(AScriptText, AScriptType);
    end;
  except
  end;
end;

procedure TMiniWebview.ExecuteJS(AScriptText: string);
begin
  ExecuteScript(AScriptText, 'javascript');
end;

procedure TMiniWebview.LoadHTML(const AStr: string);
var
  LMem: TMemoryStream;
  LBs: TBytes;
begin
  if not HandleAllocated then
    HandleNeeded;
  if FWebBrowser = nil then
    Exit;
  LMem := TMemoryStream.Create;
  try
    LBs := BytesOf(AStr);
    LMem.Write(LBs[0], Length(LBs));
    LMem.Position := 0;
    Navigate('about:blank');
    while ReadyState < rsInterActive do
    begin
      Application.ProcessMessages;
      Sleep(1);
    end;
    if Assigned(FWebBrowser.Document) then
      (FWebBrowser.Document as IPersistStreamInit).Load(TStreamAdapter.Create(LMem));
  finally
    LMem.Free;;
  end;
end;

function TMiniWebview.SetMenu(hmenuShared: HMenu; holemenu: HMenu;
  hwndActiveObject: HWnd): HResult; stdcall;
{$IFDEF FPC}
begin
  Result := E_NOTIMPL;
end;
{$ELSE}
var
  Menu: TMainMenu;
begin
  Menu := GetMainMenu;
  Result := S_OK;
  if Menu <> nil then
  begin
    Menu.SetOle2MenuHandle(hmenuShared);
    Result := OleSetMenuDescriptor(holemenu, Menu.WindowHandle,
      hwndActiveObject, nil, nil);
  end;
end;
{$ENDIF}

function TMiniWebview.SetStatusText(pszStatusText: POleStr): HResult; stdcall;
begin
  Result := S_OK;
end;


procedure TMiniWebview.SetUIActive(Active: Boolean);
{$IFDEF FPC}
var
  Form: TCustomForm;
begin
  Form := GetParentForm(Self);
  if Form <> nil then
    if Active then
    begin
      if (Form.ActiveControl <> nil) and
        (Form.ActiveControl <> Self) then
        Form.ActiveControl.Perform(CM_UIDEACTIVATE, 0, 0);
      Form.ActiveControl := Self;
    end else
      if Form.ActiveControl = Self then Form.ActiveControl := nil;
end;
{$ELSE}
var
  Form: TCustomForm;
begin
  Form := GetParentForm(Self);
  if Form <> nil then
    if Active then
    begin
      if (Form.ActiveOleControl <> nil) and
        (Form.ActiveOleControl <> Self) then
        Form.ActiveOleControl.Perform(CM_UIDEACTIVATE, 0, 0);
      Form.ActiveOleControl := Self;
    end else
      if Form.ActiveOleControl = Self then Form.ActiveOleControl := nil;
end;
{$ENDIF}

function TMiniWebview.ShowObject: HResult; stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.ShowPropertyFrame: HResult; stdcall;
begin
  Result := E_NOTIMPL;
end;

{$IFDEF FPC}
function TMiniWebview.TransformCoords(var ptlHimetric: _POINTL;
  var ptfContainer: tagPOINTF; flags: LongWord): HResult; stdcall;
{$ELSE}
function TMiniWebview.TransformCoords(var ptlHimetric: TPoint; var ptfContainer: TPointF; flags: Longint): HResult;
{$ENDIF}
begin
  if flags and XFORMCOORDS_HIMETRICTOCONTAINER <> 0 then
  begin
    ptfContainer.X := MulDiv(ptlHimetric.X, Screen.PixelsPerInch, 2540);
    ptfContainer.Y := MulDiv(ptlHimetric.Y, Screen.PixelsPerInch, 2540);
  end else
  begin
    ptlHimetric.X := Integer(Round(ptfContainer.X * 2540 / Screen.PixelsPerInch));
    ptlHimetric.Y := Integer(Round(ptfContainer.Y * 2540 / Screen.PixelsPerInch));
  end;
  Result := S_OK;
end;

procedure TMiniWebview.VisibleChanging;
var
  Flag: Integer;
begin
  inherited;
  if not(csDesigning in ComponentState) then
  begin
    if FOleObject = nil then
      Exit;
    Flag := OLEIVERB_HIDE;
    if Visible then
      Flag := OLEIVERB_SHOW;
    FOleObject.DoVerb(Flag, nil, Self, -1, Handle, ClientRect);
  end;
end;

function TMiniWebview.ShowContextMenu(const dwID: DWORD; const ppt: PPOINT;
  const pcmdtReserved: IUnknown; const pdispReserved: IDispatch): HRESULT;
  stdcall;
begin
  Result := S_FALSE;
end;

function TMiniWebview.GetHostInfo(var pInfo: TDocHostUIInfo): HRESULT; stdcall;
begin
  pInfo.cbSize := sizeof(TDocHostUIInfo);
  pInfo.dwFlags := 0;
  pInfo.dwFlags :={ DOCHOSTUIFLAG_DIALOG or }DOCHOSTUIFLAG_DISABLE_HELP_MENU or
    DOCHOSTUIFLAG_NO3DBORDER or DOCHOSTUIFLAG_THEME;
  // 这里暂时不能禁用滚动啊，不然鼠标没法用了
  //pInfo.dwFlags := pInfo.dwFlags or DOCHOSTUIFLAG_SCROLL_NO;

  pInfo.dwDoubleClick := DOCHOSTUIDBLCLK_DEFAULT;
  Result := S_OK;
end;

function TMiniWebview.ShowUI(const dwID: DWORD;
  const pActiveObject: IOleInPlaceActiveObject;
  const pCommandTarget: IOleCommandTarget; const pFrame: IOleInPlaceFrame;
  const pDoc: IOleInPlaceUIWindow): HRESULT; stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.HideUI: HRESULT; stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.UpdateUI: HRESULT; stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.IDocHostUIHandler_EnableModeless(const fEnable: BOOL
  ): HRESULT; stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.OnDocWindowActivate(const fActivate: BOOL): HRESULT;
  stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.OnFrameWindowActivate(const fActivate: BOOL): HRESULT;
  stdcall;
begin
  Result := S_OK;
end;

function TMiniWebview.ResizeBorder(const prcBorder: PRECT;
  const pUIWindow: IOleInPlaceUIWindow; const FrameWindow: BOOL): HRESULT;
  stdcall;
begin
  Result := S_FALSE;
end;

function TMiniWebview.TranslateAccelerator(const lpMsg: PMSG;
  const pguidCmdGroup: PGUID; const nCmdID: DWORD): HRESULT; stdcall;
begin
  Result := S_FALSE;
end;

function TMiniWebview.GetOptionKeyPath(var pchKey: POLESTR; const dw: DWORD
  ): HRESULT; stdcall;
begin
  Result := E_FAIL;
end;

function TMiniWebview.GetDropTarget(const pDropTarget: IDropTarget; out
  ppDropTarget: IDropTarget): HRESULT; stdcall;
begin
  ppDropTarget := nil;
  Result := E_FAIL;
end;

function TMiniWebview.GetExternal(out ppDispatch: IDispatch): HRESULT; stdcall;
begin
  ppDispatch := FExternalObj;
  Result := S_OK;
end;

function TMiniWebview.TranslateUrl(const dwTranslate: DWORD;
  const pchURLIn: POLESTR; var ppchURLOut: POLESTR): HRESULT; stdcall;
begin
  Result := S_FALSE;
end;

function TMiniWebview.FilterDataObject(const pDO: IDataObject; out
  ppDORet: IDataObject): HRESULT; stdcall;
begin
  ppDORet := nil;
  Result := S_FALSE;
end;

{$IFDEF WINDOWS}
{$IFDEF FPC}
var
  uOldFPU: Word;

initialization
  OleInitialize(nil);
  // 正式就可以免除这个
  uOldFPU := GetMXCSR;
  SetMXCSR(uOldFPU or $0080);

finalization
  SetMXCSR(uOldFPU);

  OleUninitialize;
{$ENDIF}
{$ENDIF WINDOWS}

{$ENDIF MSWINDOWS}
