// +build windows

package miniblink

// WkeWebview

func (w WkeWebView) Name() string {
	return wkeGetName(w)
}

func (w WkeWebView) SetName(name string) {
	wkeSetName(w, name)
}

func (w WkeWebView) IsTransparent() bool {
	return wkeIsTransparent(w)
}

func (w WkeWebView) SetTransparent(val bool) {
	wkeSetTransparent(w, val)
}

func (w WkeWebView) LoadURL(url string) {
	if isLcl {
		wkeLoadURL(w, url)
	} else {
		wkeLoadURLW(w, url)
	}
}

func (w WkeWebView) LoadHTML(html string) {
	if isLcl {
		wkeLoadHTML(w, html)
	} else {
		wkeLoadHTMLW(w, html)
	}
}

func (w WkeWebView) LoadFile(html string) {
	if isLcl {
		wkeLoadFile(w, html)
	} else {
		wkeLoadFileW(w, html)
	}
}

func (w WkeWebView) URL() string {
	return wkeGetURL(w)
}

func (w WkeWebView) IsLoading() bool {
	return wkeIsLoading(w)
}

func (w WkeWebView) IsLoadingFailed() bool {
	return wkeIsLoadingFailed(w)
}

func (w WkeWebView) IsLoadingSucceeded() bool {
	return wkeIsLoadingSucceeded(w)
}

func (w WkeWebView) IsDocumentReady() bool {
	return wkeIsDocumentReady(w)
}

func (w WkeWebView) StopLoading() {
	wkeStopLoading(w)
}

func (w WkeWebView) Reload() {
	wkeReload(w)
}

func (w WkeWebView) Title() string {
	if isLcl {
		return wkeTitle(w)
	} else {
		return wkeTitleW(w)
	}
}

func (w WkeWebView) Resize(width, height int) {
	wkeResize(w, width, height)
}

func (w WkeWebView) Width() int {
	return wkeWidth(w)
}

func (w WkeWebView) Height() int {
	return wkeHeight(w)
}

func (w WkeWebView) ContentsWidth() int {
	return wkeContentsWidth(w)
}

func (w WkeWebView) ContentsHeight() int {
	return wkeContentsHeight(w)
}

func (w WkeWebView) SetDirty(dirty bool) {
	wkeSetDirty(w, dirty)
}

func (w WkeWebView) IsDirty() bool {
	return wkeIsDirty(w)
}

func (w WkeWebView) AddDirtyArea(x, y, width, height int) {
	wkeAddDirtyArea(w, x, y, width, height)
}

func (w WkeWebView) LayoutIfNeeded() {
	wkeLayoutIfNeeded(w)
}

func (w WkeWebView) Paint(bits uintptr, pitch int) {
	wkePaint(w, bits, pitch)
}

func (w WkeWebView) CanGoBack() bool {
	return wkeCanGoBack(w)
}

func (w WkeWebView) GoBack() {
	wkeGoBack(w)
}

func (w WkeWebView) CanGoForward() bool {
	return wkeCanGoForward(w)
}

func (w WkeWebView) GoForward() {
	wkeGoForward(w)
}

func (w WkeWebView) EditorSelectAll() {
	wkeEditorSelectAll(w)
}

func (w WkeWebView) EditorUnSelect() {
	wkeEditorUnSelect(w)
}

func (w WkeWebView) EditorCopy() {
	wkeEditorCopy(w)
}

func (w WkeWebView) EditorCut() {
	wkeEditorCut(w)
}

func (w WkeWebView) EditorPaste() {
	wkeEditorPaste(w)
}

func (w WkeWebView) EditorDelete() {
	wkeEditorDelete(w)
}

func (w WkeWebView) EditorUndo() {
	wkeEditorUndo(w)
}

func (w WkeWebView) EditorRedo() {
	wkeEditorRedo(w)
}

func (w WkeWebView) SetCookieEnabled(enable bool) {
	wkeSetCookieEnabled(w, enable)
}

func (w WkeWebView) IsCookieEnabled() bool {
	return wkeIsCookieEnabled(w)
}

func (w WkeWebView) SetMediaVolume(volume float32) {
	wkeSetMediaVolume(w, volume)
}

func (w WkeWebView) MediaVolume() float32 {
	return wkeMediaVolume(w)
}

func (w WkeWebView) FireMouseEvent(message uint, x, y int, flags uint) bool {
	return wkeFireMouseEvent(w, message, x, y, flags)
}

func (w WkeWebView) FireContextMenuEvent(x, y int, flags uint) bool {
	return wkeFireContextMenuEvent(w, x, y, flags)
}

func (w WkeWebView) FireMouseWheelEvent(x, y, delta int, flags uint) bool {
	return wkeFireMouseWheelEvent(w, x, y, delta, flags)
}

func (w WkeWebView) FireKeyUpEvent(virtualKeyCode, flags uint, systemKey bool) bool {
	return wkeFireKeyUpEvent(w, virtualKeyCode, flags, systemKey)
}

func (w WkeWebView) FireKeyDownEvent(virtualKeyCode, flags uint, systemKey bool) bool {
	return wkeFireKeyDownEvent(w, virtualKeyCode, flags, systemKey)
}

func (w WkeWebView) FireKeyPressEvent(virtualKeyCode, flags uint, systemKey bool) bool {
	return wkeFireKeyPressEvent(w, virtualKeyCode, flags, systemKey)
}

func (w WkeWebView) SetFocus() {
	wkeSetFocus(w)
}

func (w WkeWebView) KillFocus() {
	wkeKillFocus(w)
}

func (w WkeWebView) RunJS(script string) JsValue {
	if isLcl {
		return wkeRunJS(w, script)
	} else {
		return wkeRunJSW(w, script)
	}
}

func (w WkeWebView) GlobalExec() JsExecState {
	return wkeGlobalExec(w)
}

func (w WkeWebView) Sleep() {
	wkeSleep(w)
}

func (w WkeWebView) Wake() {
	wkeWake(w)
}

func (w WkeWebView) IsAwake() bool {
	return wkeIsAwake(w)
}

func (w WkeWebView) SetZoomFactor(factor float32) {
	wkeSetZoomFactor(w, factor)
}

func (w WkeWebView) ZoomFactor() float32 {
	return wkeZoomFactor(w)
}

func (w WkeWebView) SetEditable(editable bool) {
	wkeSetEditable(w, editable)
}

func (w WkeWebView) SetClientHandler(handler *WkeClientHandler) {
	wkeSetClientHandler(w, handler)
}

func (w WkeWebView) GetClientHandler() *WkeClientHandler {
	return wkeGetClientHandler(w)
}
