// +build windows

package miniblink

import (
	"syscall"
)

var (
	_wkeCreateViewCallback = syscall.NewCallbackCDecl(fnwkeCreateViewCallback)
)

// typedef wkeWebView(WKE_CALL_TYPE*wkeCreateViewCallback)(wkeWebView webView, void* param, wkeNavigationType navigationType, const wkeString url, const wkeWindowFeatures* windowFeatures);
func fnwkeCreateViewCallback(webView WkeWebView, param uintptr, navigationType WkeNavigationType, url WkeString, windowFeatures *WkeWindowFeatures) WkeWebView {
	//p := (*TMiniBlinkWebview)(unsafe.Pointer(param)).onCreateView
	//if p != nil {
	//	p(info)
	//}
	return webView
}
