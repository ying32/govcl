// +build windows

package miniblink

import (
	"unsafe"

	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

type TMiniBlinkWebview struct {
	wkePtr WkeWebView
}

func NewMiniBlinkWebview(hWnd types.HWND) *TMiniBlinkWebview {
	w := new(TMiniBlinkWebview)
	r, _ := win.GetClientRect2(hWnd)
	w.wkePtr = wkeCreateWebWindow(WKE_WINDOW_TYPE_CONTROL, hWnd, 0, 0, int(r.Width()), int(r.Height()))
	ptr := uintptr(unsafe.Pointer(w))
	wkeOnCreateView(w.wkePtr, WkeCreateViewCallback(_wkeCreateViewCallback), ptr)
	return w
}

func (w *TMiniBlinkWebview) IsValid() bool {
	return w.wkePtr != 0
}

func (w *TMiniBlinkWebview) Free() {
	if w.IsValid() {
		wkeDestroyWebWindow(w.wkePtr)
		w.wkePtr = 0
	}
}

func (w *TMiniBlinkWebview) MoveWindow(x, y, width, height int) {
	if w.IsValid() {
		wkeMoveWindow(w.wkePtr, x, y, width, height)
	}
}

func (w *TMiniBlinkWebview) Show(val bool) {
	if w.IsValid() {
		wkeShowWindow(w.wkePtr, val)
	}
}

func (w *TMiniBlinkWebview) LoadURL(url string) {
	if w.IsValid() {
		wkeLoadURL(w.wkePtr, url)
	}
}

func (w *TMiniBlinkWebview) LoadURLW(url string) {
	if w.IsValid() {
		wkeLoadURLW(w.wkePtr, url)
	}
}

func (w *TMiniBlinkWebview) LoadHTML(html string) {
	if w.IsValid() {
		wkeLoadHTML(w.wkePtr, html)
	}
}

func (w *TMiniBlinkWebview) LoadHTMLW(html string) {
	if w.IsValid() {
		wkeLoadHTMLW(w.wkePtr, html)
	}
}

func (w *TMiniBlinkWebview) LoadFile(html string) {
	if w.IsValid() {
		wkeLoadFile(w.wkePtr, html)
	}
}

func (w *TMiniBlinkWebview) LoadFileW(html string) {
	if w.IsValid() {
		wkeLoadFileW(w.wkePtr, html)
	}
}

func (w *TMiniBlinkWebview) Reload() {
	if w.IsValid() {
		wkeReload(w.wkePtr)
	}
}

func (w *TMiniBlinkWebview) GoBack() {
	if w.IsValid() {
		wkeGoBack(w.wkePtr)
	}
}

func (w *TMiniBlinkWebview) GoForward() {
	if w.IsValid() {
		wkeGoForward(w.wkePtr)
	}
}
