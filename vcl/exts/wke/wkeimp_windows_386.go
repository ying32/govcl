// govcl wke测试头文件，by: ying32
// https://gitee.com/ying32/ying32

package wke

import (
	"syscall"

	"unsafe"

	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/types"
)

type WindowType int32

const (
	WKE_WINDOW_TYPE_POPUP WindowType = iota + 0
	WKE_WINDOW_TYPE_TRANSPARENT
	WKE_WINDOW_TYPE_CONTROL
)

type NavigationType int32

const (
	WKE_NAVIGATION_TYPE_LINKCLICK NavigationType = iota + 0
	WKE_NAVIGATION_TYPE_FORMSUBMITTE
	WKE_NAVIGATION_TYPE_BACKFORWARD
	WKE_NAVIGATION_TYPE_RELOAD
	WKE_NAVIGATION_TYPE_FORMRESUBMITT
	WKE_NAVIGATION_TYPE_OTHER
)

type LoadingResult int32

const (
	WKE_LOADING_SUCCEEDED LoadingResult = iota + 0
	WKE_LOADING_FAILED
	WKE_LOADING_CANCELED
)

type DocumentReadyInfo struct {
	url              uintptr
	frameJSState     uintptr
	mainFrameJSState uintptr
}

/*
  PwkeNewViewInfo = ^wkeNewViewInfo;
  wkeNewViewInfo = record
    navigationType: wkeNavigationType;
    url: wkeString;
    target: wkeString;

    x: Integer;
    y: Integer;
    width: Integer;
    height: Integer;
    menuBarVisible: Boolean;
    statusBarVisible: bool ;
    toolBarVisible: Boolean;
    locationBarVisible: Boolean;
    scrollbarsVisible: Boolean;
    resizable: Boolean;
    fullscreen: Boolean;
  end;
  TwkeNewViewInfo = wkeNewViewInfo;
*/

var (
	wkedll               = syscall.NewLazyDLL("wke.dll")
	_wkeInitialize       = wkedll.NewProc("wkeInitialize")
	_wkeFinalize         = wkedll.NewProc("wkeFinalize")
	_wkeCreateWebWindow  = wkedll.NewProc("wkeCreateWebWindow")
	_wkeShowWindow       = wkedll.NewProc("wkeShowWindow")
	_wkeDestroyWebWindow = wkedll.NewProc("wkeDestroyWebWindow")
	_wkeMoveWindow       = wkedll.NewProc("wkeMoveWindow")
	_wkeLoadW            = wkedll.NewProc("wkeLoadW")
	_wkeRepaintAllNeeded = wkedll.NewProc("wkeRepaintAllNeeded")
	_wkeReload           = wkedll.NewProc("wkeReload")
	_wkeCanGoBack        = wkedll.NewProc("wkeCanGoBack")
	_wkeGoBack           = wkedll.NewProc("wkeGoBack")
	_wkeCanGoForward     = wkedll.NewProc("wkeCanGoForward")
	_wkeGoForward        = wkedll.NewProc("wkeGoForward")
	_wkeOnTitleChanged   = wkedll.NewProc("wkeOnTitleChanged")
	_wkeOnURLChanged     = wkedll.NewProc("wkeOnURLChanged")
	_wkeGetStringW       = wkedll.NewProc("wkeGetStringW")

	// 事件设置
	//procedure wkeOnPaintUpdated; external wkedll name 'wkeOnPaintUpdated';
	//procedure wkeOnAlertBox; external wkedll name 'wkeOnAlertBox';
	//procedure wkeOnConfirmBox; external wkedll name 'wkeOnConfirmBox';
	//procedure wkeOnPromptBox; external wkedll name 'wkeOnPromptBox';
	//procedure wkeOnConsoleMessage; external wkedll name 'wkeOnConsoleMessage';
	_wkeOnNavigation    = wkedll.NewProc("wkeOnNavigation")
	_wkeOnCreateView    = wkedll.NewProc("wkeOnCreateView")
	_wkeOnDocumentReady = wkedll.NewProc("wkeOnDocumentReady")
	_wkeOnLoadingFinish = wkedll.NewProc("wkeOnLoadingFinish")

	// callback
	_wkeTitleChangedCallback  = syscall.NewCallbackCDecl(wkeTitleChangedCallback)
	_wkeURLChangedCallback    = syscall.NewCallbackCDecl(wkeURLChangedCallback)
	_wkeNavigationCallback    = syscall.NewCallbackCDecl(wkeNavigationCallback)
	_wkeLoadingFinishCallback = syscall.NewCallbackCDecl(wkeLoadingFinishCallback)
	_wkeDocumentReadyCallback = syscall.NewCallbackCDecl(wkeDocumentReadyCallback)
	_wkeCreateViewCallback    = syscall.NewCallbackCDecl(wkeCreateViewCallback)
)

func wkeTitleChangedCallback(webView, param, title uintptr) uintptr {
	p := (*TWkeWebBrowser)(unsafe.Pointer(param)).onTitleChnaged
	if p != nil {
		p(wkeGetStringW(title))
	}
	return 1
}

func wkeURLChangedCallback(webView, param, url uintptr) uintptr {
	p := (*TWkeWebBrowser)(unsafe.Pointer(param)).onURLChanged
	if p != nil {
		p(wkeGetStringW(url))
	}
	return 1
}

func wkeNavigationCallback(webView, param uintptr, navigationType NavigationType, url uintptr) uintptr {
	p := (*TWkeWebBrowser)(unsafe.Pointer(param)).onNavigation
	if p != nil {
		p(navigationType, wkeGetStringW(url))
	}
	return 1
}

func wkeLoadingFinishCallback(webView, param, url uintptr, result LoadingResult, failedReason uintptr) uintptr {
	p := (*TWkeWebBrowser)(unsafe.Pointer(param)).onLoadingFinish
	if p != nil {
		p(wkeGetStringW(url), result, wkeGetStringW(failedReason))
	}
	return 1
}

func wkeDocumentReadyCallback(webView, param, info uintptr) uintptr {
	p := (*TWkeWebBrowser)(unsafe.Pointer(param)).onDocumentReady
	if p != nil {
		p(info)
	}
	return 1
}

func wkeCreateViewCallback(webView, param, info uintptr) uintptr {
	p := (*TWkeWebBrowser)(unsafe.Pointer(param)).onCreateView
	if p != nil {
		p(info)
	}
	return param
}

//wkeAlertBoxCallback = procedure(webView: wkeWebView; param: Pointer; msg: wkeString); cdecl;
//wkeConfirmBoxCallback = function(webView: wkeWebView; param: Pointer; msg: wkeString): Boolean; cdecl;
//wkePromptBoxCallback = function(webView: wkeWebView; param: Pointer; msg: wkeString; defaultResult: wkeString; result: wkeString): Boolean; cdecl;
//wkeDocumentReadyCallback = procedure(webView: wkeWebView; param: Pointer; info: PwkeDocumentReadyInfo); cdecl;
//wkeWindowClosingCallback = function(webWindow: wkeWebView; param: Pointer): Boolean; cdecl;
//wkeWindowDestroyCallback = procedure(webWindow: wkeWebView; param: Pointer); cdecl;

//--------------------------------------------------------

func Initialize() {
	_wkeInitialize.Call()
}

func Finalize() {
	_wkeFinalize.Call()
}

func wkeCreateWebWindow(aType WindowType, hWd types.HWND, x, y, width, height int32) uintptr {
	r, _, _ := _wkeCreateWebWindow.Call(uintptr(aType), uintptr(hWd), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return r
}

func wkeShowWindow(webWindow uintptr, show bool) {
	_wkeShowWindow.Call(webWindow, api.GoBoolToDBool(show))
}

func wkeDestroyWebWindow(webWindow uintptr) {
	_wkeDestroyWebWindow.Call(webWindow)
}

func wkeMoveWindow(webWindow uintptr, x, y, width, height int32) {
	_wkeMoveWindow.Call(webWindow, uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func wkeLoadW(webWindow uintptr, url string) {
	_wkeLoadW.Call(webWindow, api.GoStrToDStr(url))
}

func RepaintAllNeeded() bool {
	r, _, _ := _wkeRepaintAllNeeded.Call()
	return api.DBoolToGoBool(r)

}

func wkeReload(webWindow uintptr) {
	_wkeReload.Call(webWindow)
}

func wkeCanGoBack(webWindow uintptr) bool {
	r, _, _ := _wkeCanGoBack.Call(webWindow)
	return api.DBoolToGoBool(r)
}

func wkeGoBack(webWindow uintptr) {
	_wkeGoBack.Call(webWindow)
}

func wkeCanGoForward(webWindow uintptr) bool {
	r, _, _ := _wkeCanGoForward.Call(webWindow)
	return api.DBoolToGoBool(r)
}

func wkeGoForward(webWindow uintptr) {
	_wkeGoForward.Call(webWindow)
}

func wkeGetStringW(wkeStr uintptr) string {
	r, _, _ := _wkeGetStringW.Call(wkeStr)
	return api.DStrToGoStr(r)
}

func wkeOnTitleChanged(webView, callback, callbackParam uintptr) {
	_wkeOnTitleChanged.Call(webView, callback, callbackParam)
}

func wkeOnURLChanged(webView, callback, callbackParam uintptr) {
	_wkeOnURLChanged.Call(webView, callback, callbackParam)
}

func wkeOnNavigation(webView, callback, callbackParam uintptr) {
	_wkeOnNavigation.Call(webView, callback, callbackParam)
}

func wkeOnLoadingFinish(webView, callback, callbackParam uintptr) {
	_wkeOnLoadingFinish.Call(webView, callback, callbackParam)
}

func wkeOnDocumentReady(webView, callback, callbackParam uintptr) {
	_wkeOnDocumentReady.Call(webView, callback, callbackParam)
}

func wkeOnCreateView(webView, callback, callbackParam uintptr) {
	_wkeOnCreateView.Call(webView, callback, callbackParam)
}

//procedure wkeOnPaintUpdated(webView: wkeWebView; callback: wkePaintUpdatedCallback; callbackParam: Pointer); cdecl;
//procedure wkeOnAlertBox(webView: wkeWebView; callback: wkeAlertBoxCallback; callbackParam: Pointer); cdecl;
//procedure wkeOnConfirmBox(webView: wkeWebView; callback: wkeConfirmBoxCallback; callbackParam: Pointer); cdecl;
//procedure wkeOnPromptBox(webView: wkeWebView; callback: wkePromptBoxCallback; callbackParam: Pointer); cdecl;
//procedure wkeOnCreateView(webView: wkeWebView; callback: wkeCreateViewCallback; param: Pointer); cdecl;
//procedure wkeOnConsoleMessage(webView: wkeWebView; callback: wkeConsoleMessageCallback; callbackParam: Pointer); cdecl;
