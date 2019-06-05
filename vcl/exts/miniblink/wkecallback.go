// +build windows

package miniblink

import (
	"syscall"
	"unsafe"
)

var (
	_wkeCreateViewCallback   = syscall.NewCallbackCDecl(fnwkeCreateViewCallback)
	_wkeTitleChangedCallback = syscall.NewCallbackCDecl(fnwkeTitleChangedCallback)
)

type TOnCreateViewEvent func(sender *TMiniBlinkWebview, navigationType WkeNavigationType, url WkeString, windowFeatures *WkeWindowFeatures, result *WkeWebView)
type TOnTitleChangedEvent func(sender *TMiniBlinkWebview, title string)

func getObj(u uintptr) *TMiniBlinkWebview {
	return (*TMiniBlinkWebview)(unsafe.Pointer(u))
}

// typedef wkeWebView(WKE_CALL_TYPE*wkeCreateViewCallback)(wkeWebView webView, void* param, wkeNavigationType navigationType, const wkeString url, const wkeWindowFeatures* windowFeatures);
func fnwkeCreateViewCallback(webView WkeWebView, param uintptr, navigationType WkeNavigationType, url WkeString, windowFeatures *WkeWindowFeatures) WkeWebView {
	ret := webView
	if param != 0 {
		obj := getObj(param)
		proc := obj.OnCreateView
		if proc != nil {
			obj.OnCreateView(obj, navigationType, url, windowFeatures, &ret)
		}
	}
	return ret
}

// typedef void(WKE_CALL_TYPE*wkeTitleChangedCallback)(wkeWebView webView, void* param, const wkeString title);
func fnwkeTitleChangedCallback(webView WkeWebView, param uintptr, title WkeString) uintptr {
	if param != 0 {
		obj := getObj(param)
		proc := obj.OnTitleChanged
		if proc != nil {
			var str string
			if isLcl {
				str = wkeGetString(title)
			} else {
				str = wkeGetStringW(title)
			}
			obj.OnTitleChanged(obj, str)
		}
	}
	return 1
}
