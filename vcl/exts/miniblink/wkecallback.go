// +build windows

package miniblink

import (
	"syscall"
	"unsafe"
)

var (
	_wkeCreateViewCallback    = syscall.NewCallbackCDecl(fnwkeCreateViewCallback)
	_wkeTitleChangedCallback  = syscall.NewCallbackCDecl(fnwkeTitleChangedCallback)
	_wkeURLChangedCallback    = syscall.NewCallbackCDecl(fnwkeURLChangedCallback)
	_wkeNavigationCallback    = syscall.NewCallbackCDecl(fnwkeNavigationCallback)
	_wkeLoadingFinishCallback = syscall.NewCallbackCDecl(fnwkeLoadingFinishCallback)
	_wkeDocumentReadyCallback = syscall.NewCallbackCDecl(fnwkeDocumentReadyCallback)
)

type TOnCreateViewEvent func(sender *TMiniBlinkWebview, navigationType WkeNavigationType, url WkeString, windowFeatures *WkeWindowFeatures, result *WkeWebView)
type TOnTitleChangedEvent func(sender *TMiniBlinkWebview, title string)
type TOnURLChangedEvent func(sender *TMiniBlinkWebview, url string)
type TOnNavigationEvent func(sender *TMiniBlinkWebview, navigationType WkeNavigationType, url string)
type TOnLoadingFinishEvent func(sender *TMiniBlinkWebview, url string, result WkeLoadingResult, failedReason string)
type TOnDocumentReadyEvent func(sender *TMiniBlinkWebview)

func getObj(u uintptr) *TMiniBlinkWebview {
	return (*TMiniBlinkWebview)(unsafe.Pointer(u))
}

/// typedef wkeWebView(WKE_CALL_TYPE*wkeCreateViewCallback)(wkeWebView webView, void* param, wkeNavigationType navigationType, const wkeString url, const wkeWindowFeatures* windowFeatures);
func fnwkeCreateViewCallback(webView WkeWebView, param uintptr, navigationType WkeNavigationType, url WkeString, windowFeatures *WkeWindowFeatures) WkeWebView {
	ret := webView
	if param != 0 {
		obj := getObj(param)
		proc := obj.OnCreateView
		if proc != nil {
			proc(obj, navigationType, url, windowFeatures, &ret)
		}
	}
	return ret
}

/// typedef void(WKE_CALL_TYPE*wkeTitleChangedCallback)(wkeWebView webView, void* param, const wkeString title);

/// export fnwkeTitleChangedCallback
func fnwkeTitleChangedCallback(webView unsafe.Pointer, param unsafe.Pointer, title unsafe.Pointer) uintptr {
	if param != nil {
		obj := getObj(uintptr(param))
		proc := obj.OnTitleChanged
		if proc != nil {
			if isLcl {
				proc(obj, wkeGetString(WkeString(title)))
			} else {
				proc(obj, wkeGetStringW(WkeString(title)))
			}
		}
	}
	return 0
}

/// 	typedef void(WKE_CALL_TYPE*wkeURLChangedCallback)(wkeWebView webView, void* param, const wkeString url);
func fnwkeURLChangedCallback(view WkeWebView, param uintptr, url WkeString) uintptr {
	if param != 0 {
		obj := getObj(param)
		proc := obj.OnURLChanged
		if proc != nil {
			if isLcl {
				proc(obj, wkeGetString(url))
			} else {
				proc(obj, wkeGetStringW(url))
			}
		}
	}
	return 1
}

/// typedef bool(WKE_CALL_TYPE*wkeNavigationCallback)(wkeWebView webView, void* param, wkeNavigationType navigationType, wkeString url);
func fnwkeNavigationCallback(view WkeWebView, param uintptr, navigationType WkeNavigationType, url WkeString) uintptr {
	if param != 0 {
		obj := getObj(param)
		proc := obj.OnNavigation
		if proc != nil {
			if isLcl {
				proc(obj, navigationType, wkeGetString(url))
			} else {
				proc(obj, navigationType, wkeGetStringW(url))
			}
		}
	}
	return 1
}

/// typedef void(WKE_CALL_TYPE*wkeLoadingFinishCallback)(wkeWebView webView, void* param, const wkeString url, wkeLoadingResult result, const wkeString failedReason);
func fnwkeLoadingFinishCallback(view WkeWebView, param uintptr, url WkeString, result WkeLoadingResult, failedReason WkeString) uintptr {
	if param != 0 {
		obj := getObj(param)
		proc := obj.OnLoadingFinish
		if proc != nil {
			if isLcl {
				proc(obj, wkeGetString(url), result, wkeGetString(failedReason))
			} else {
				proc(obj, wkeGetStringW(url), result, wkeGetStringW(failedReason))
			}
		}
	}
	return 1
}

/// typedef void(WKE_CALL_TYPE*wkeDocumentReadyCallback)(wkeWebView webView, void* param);
func fnwkeDocumentReadyCallback(view WkeWebView, param uintptr) uintptr {
	if param != 0 {
		obj := getObj(param)
		proc := obj.OnDocumentReady
		if proc != nil {
			proc(obj)
		}
	}
	return 1
}
