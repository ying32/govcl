// govcl wke，by: ying32
// https://gitee.com/ying32/ying32

package wke

import (
	"unsafe"

	"gitee.com/ying32/govcl/vcl/types"
	"gitee.com/ying32/govcl/vcl/win"
)

type TOnTitleChangedEvent func(title string)
type TOnURLChangedEvent func(url string)
type TOnNavigationEvent func(navigationType NavigationType, url string)
type TOnLoadingFinishEvent func(url string, result LoadingResult, failedReason string)
type TOnDocumentReadyEvent func(info uintptr)
type TOnCreateViewEvent func(info uintptr)

type TWkeWebBrowser struct {
	wkePtr          uintptr
	onTitleChnaged  TOnTitleChangedEvent
	onURLChanged    TOnURLChangedEvent
	onNavigation    TOnNavigationEvent
	onLoadingFinish TOnLoadingFinishEvent
	onDocumentReady TOnDocumentReadyEvent
	onCreateView    TOnCreateViewEvent
}

func NewWkeWebBrowser(hWnd types.HWND) *TWkeWebBrowser {
	w := new(TWkeWebBrowser)
	r := win.GetClientRect(hWnd)
	w.wkePtr = wkeCreateWebWindow(WKE_WINDOW_TYPE_CONTROL, hWnd, 0, 0, r.Right-r.Left, r.Bottom-r.Top)
	ptr := uintptr(unsafe.Pointer(w))
	wkeOnTitleChanged(w.wkePtr, _wkeTitleChangedCallback, ptr)
	wkeOnURLChanged(w.wkePtr, _wkeURLChangedCallback, ptr)
	wkeOnNavigation(w.wkePtr, _wkeNavigationCallback, ptr)
	wkeOnLoadingFinish(w.wkePtr, _wkeLoadingFinishCallback, ptr)
	wkeOnDocumentReady(w.wkePtr, _wkeDocumentReadyCallback, ptr)
	wkeOnCreateView(w.wkePtr, _wkeCreateViewCallback, ptr)

	return w
}

func (w *TWkeWebBrowser) IsVaild() bool {
	return w.wkePtr != 0
}

func (w *TWkeWebBrowser) Show(val bool) {
	if w.IsVaild() {
		wkeShowWindow(w.wkePtr, val)
	}
}

func (w *TWkeWebBrowser) Free() {
	if w.IsVaild() {
		wkeDestroyWebWindow(w.wkePtr)
		w.wkePtr = 0
	}
}

func (w *TWkeWebBrowser) MoveWindow(x, y, width, height int32) {
	if w.IsVaild() {
		wkeMoveWindow(w.wkePtr, x, y, width, height)
	}
}

func (w *TWkeWebBrowser) Load(url string) {
	if w.IsVaild() {
		wkeLoadW(w.wkePtr, url)
	}
}

func (w *TWkeWebBrowser) Reload() {
	if w.IsVaild() {
		wkeReload(w.wkePtr)
	}
}

func (w *TWkeWebBrowser) GoBack() {
	if w.IsVaild() {
		wkeGoBack(w.wkePtr)
	}
}

func (w *TWkeWebBrowser) GoForward() {
	if w.IsVaild() {
		wkeGoForward(w.wkePtr)
	}
}

func (w *TWkeWebBrowser) SetOnTitleChanged(e TOnTitleChangedEvent) {
	w.onTitleChnaged = e
}

func (w *TWkeWebBrowser) SetOnURLChanged(e TOnURLChangedEvent) {
	w.onURLChanged = e
}

func (w *TWkeWebBrowser) SetOnNavigation(e TOnNavigationEvent) {
	w.onNavigation = e
}

func (w *TWkeWebBrowser) SetOnLoadingFinish(e TOnLoadingFinishEvent) {
	w.onLoadingFinish = e
}

func (w *TWkeWebBrowser) SetOnDocumentReady(e TOnDocumentReadyEvent) {
	w.onDocumentReady = e
}

func (w *TWkeWebBrowser) SetOnCreateView(e TOnCreateViewEvent) {
	w.onCreateView = e
}
