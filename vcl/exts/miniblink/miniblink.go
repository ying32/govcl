// +build windows

package miniblink

import (
	//	"unsafe"

	"unsafe"

	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

type TMiniBlinkWebview struct {
	Webview         WkeWebView
	OnCreateView    TOnCreateViewEvent
	OnTitleChanged  TOnTitleChangedEvent
	OnURLChanged    TOnURLChangedEvent
	OnNavigation    TOnNavigationEvent
	OnLoadingFinish TOnLoadingFinishEvent
	OnDocumentReady TOnDocumentReadyEvent
	OnConsole       TOnConsoleEvent
}

func NewMiniBlinkWebview(hWnd types.HWND) *TMiniBlinkWebview {
	w := new(TMiniBlinkWebview)
	r, _ := win.GetClientRect2(hWnd)
	w.Webview = wkeCreateWebWindow(WKE_WINDOW_TYPE_CONTROL, hWnd, 0, 0, int(r.Width()), int(r.Height()))
	ptr := uintptr(unsafe.Pointer(w))
	wkeOnCreateView(w.Webview, WkeCreateViewCallback(_wkeCreateViewCallback), ptr)
	wkeOnTitleChanged(w.Webview, WkeTitleChangedCallback(_wkeTitleChangedCallback), ptr)
	wkeOnURLChanged(w.Webview, WkeURLChangedCallback(_wkeURLChangedCallback), ptr)
	wkeOnNavigation(w.Webview, WkeNavigationCallback(_wkeNavigationCallback), ptr)
	wkeOnLoadingFinish(w.Webview, WkeLoadingFinishCallback(_wkeLoadingFinishCallback), ptr)
	wkeOnDocumentReady(w.Webview, WkeDocumentReadyCallback(_wkeDocumentReadyCallback), ptr)
	wkeOnConsole(w.Webview, WkeConsoleCallback(_wkeConsoleCallback), ptr)
	return w
}

func (w *TMiniBlinkWebview) IsValid() bool {
	return w.Webview != 0
}

func (w *TMiniBlinkWebview) Free() {
	w.DestroyWebWindow()
}

//--------------

func (w *TMiniBlinkWebview) DestroyWebWindow() {
	if w.IsValid() {
		wkeDestroyWebWindow(w.Webview)
		w.Webview = 0
	}
}

func (w *TMiniBlinkWebview) MoveWindow(x, y, width, height int) {
	if w.IsValid() {
		wkeMoveWindow(w.Webview, x, y, width, height)
	}
}

func (w *TMiniBlinkWebview) Show(val bool) {
	if w.IsValid() {
		wkeShowWindow(w.Webview, val)
	}
}

func (w *TMiniBlinkWebview) Load(url string) {
	if w.IsValid() {
		wkeLoadW(w.Webview, url)
	}
}

func (w *TMiniBlinkWebview) LoadUrl(url string) {
	if w.IsValid() {
		wkeLoadURLW(w.Webview, url)
	}
}
func (w *TMiniBlinkWebview) SetZoomFactor(factor float32) {
	if w.IsValid() {
		wkeSetZoomFactor(w.Webview, factor)
	}
}
func (w *TMiniBlinkWebview) SetDeviceParameter(device string, paramStr string, paramInt int, paramFloat float32) {
	if w.IsValid() {
		wkeSetDeviceParameter(w.Webview, device, paramStr, paramInt, paramFloat)
	}
}
func (w *TMiniBlinkWebview) SetUserAgent(userAgent string) {
	if w.IsValid() {
		wkeSetUserAgentW(w.Webview, userAgent)
	}
}
