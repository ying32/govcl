// +build windows

// miniblink及wke头文件导入
// 由ying32翻译，应用于govcl，因为没有完整测试，所以不保证100%正确

package miniblink

import (
	"github.com/ying32/govcl/vcl/types"

	"github.com/ying32/dylib/floatpatch"

	"unsafe"
)

var (
	_wkeInitializeEx                      = wkedll.NewProc("wkeInitializeEx")
	_wkeShutdown                          = wkedll.NewProc("wkeShutdown")
	_wkeVersion                           = wkedll.NewProc("wkeVersion")
	_wkeVersionString                     = wkedll.NewProc("wkeVersionString")
	_wkeGC                                = wkedll.NewProc("wkeGC")
	_wkeSetResourceGc                     = wkedll.NewProc("wkeSetResourceGc")
	_wkeSetFileSystem                     = wkedll.NewProc("wkeSetFileSystem")
	_wkeWebViewName                       = wkedll.NewProc("wkeWebViewName")
	_wkeSetWebViewName                    = wkedll.NewProc("wkeSetWebViewName")
	_wkeIsLoaded                          = wkedll.NewProc("wkeIsLoaded")
	_wkeIsLoadFailed                      = wkedll.NewProc("wkeIsLoadFailed")
	_wkeIsLoadComplete                    = wkedll.NewProc("wkeIsLoadComplete")
	_wkeGetSource                         = wkedll.NewProc("wkeGetSource")
	_wkeTitle                             = wkedll.NewProc("wkeTitle")
	_wkeTitleW                            = wkedll.NewProc("wkeTitleW")
	_wkeWidth                             = wkedll.NewProc("wkeWidth")
	_wkeHeight                            = wkedll.NewProc("wkeHeight")
	_wkeContentsWidth                     = wkedll.NewProc("wkeContentsWidth")
	_wkeContentsHeight                    = wkedll.NewProc("wkeContentsHeight")
	_wkeSelectAll                         = wkedll.NewProc("wkeSelectAll")
	_wkeCopy                              = wkedll.NewProc("wkeCopy")
	_wkeCut                               = wkedll.NewProc("wkeCut")
	_wkePaste                             = wkedll.NewProc("wkePaste")
	_wkeDelete                            = wkedll.NewProc("wkeDelete")
	_wkeCookieEnabled                     = wkedll.NewProc("wkeCookieEnabled")
	_wkeMediaVolume                       = wkedll.NewProc("wkeMediaVolume")
	_wkeMouseEvent                        = wkedll.NewProc("wkeMouseEvent")
	_wkeContextMenuEvent                  = wkedll.NewProc("wkeContextMenuEvent")
	_wkeMouseWheel                        = wkedll.NewProc("wkeMouseWheel")
	_wkeKeyUp                             = wkedll.NewProc("wkeKeyUp")
	_wkeKeyDown                           = wkedll.NewProc("wkeKeyDown")
	_wkeKeyPress                          = wkedll.NewProc("wkeKeyPress")
	_wkeFocus                             = wkedll.NewProc("wkeFocus")
	_wkeUnfocus                           = wkedll.NewProc("wkeUnfocus")
	_wkeAwaken                            = wkedll.NewProc("wkeAwaken")
	_wkeZoomFactor                        = wkedll.NewProc("wkeZoomFactor")
	_wkeSetClientHandler                  = wkedll.NewProc("wkeSetClientHandler")
	_wkeGetClientHandler                  = wkedll.NewProc("wkeGetClientHandler")
	_wkeToString                          = wkedll.NewProc("wkeToString")
	_wkeToStringW                         = wkedll.NewProc("wkeToStringW")
	_jsToString                           = wkedll.NewProc("jsToString")
	_jsToStringW                          = wkedll.NewProc("jsToStringW")
	_wkeConfigure                         = wkedll.NewProc("wkeConfigure")
	_wkeIsInitialize                      = wkedll.NewProc("wkeIsInitialize")
	_wkeSetViewSettings                   = wkedll.NewProc("wkeSetViewSettings")
	_wkeSetDebugConfig                    = wkedll.NewProc("wkeSetDebugConfig")
	_wkeGetDebugConfig                    = wkedll.NewProc("wkeGetDebugConfig")
	_wkeFinalize                          = wkedll.NewProc("wkeFinalize")
	_wkeUpdate                            = wkedll.NewProc("wkeUpdate")
	_wkeGetVersion                        = wkedll.NewProc("wkeGetVersion")
	_wkeGetVersionString                  = wkedll.NewProc("wkeGetVersionString")
	_wkeCreateWebView                     = wkedll.NewProc("wkeCreateWebView")
	_wkeDestroyWebView                    = wkedll.NewProc("wkeDestroyWebView")
	_wkeSetMemoryCacheEnable              = wkedll.NewProc("wkeSetMemoryCacheEnable")
	_wkeSetMouseEnabled                   = wkedll.NewProc("wkeSetMouseEnabled")
	_wkeSetTouchEnabled                   = wkedll.NewProc("wkeSetTouchEnabled")
	_wkeSetContextMenuEnabled             = wkedll.NewProc("wkeSetContextMenuEnabled")
	_wkeSetNavigationToNewWindowEnable    = wkedll.NewProc("wkeSetNavigationToNewWindowEnable")
	_wkeSetCspCheckEnable                 = wkedll.NewProc("wkeSetCspCheckEnable")
	_wkeSetNpapiPluginsEnabled            = wkedll.NewProc("wkeSetNpapiPluginsEnabled")
	_wkeSetHeadlessEnabled                = wkedll.NewProc("wkeSetHeadlessEnabled")
	_wkeSetDragEnable                     = wkedll.NewProc("wkeSetDragEnable")
	_wkeSetDragDropEnable                 = wkedll.NewProc("wkeSetDragDropEnable")
	_wkeSetContextMenuItemShow            = wkedll.NewProc("wkeSetContextMenuItemShow")
	_wkeSetLanguage                       = wkedll.NewProc("wkeSetLanguage")
	_wkeSetViewNetInterface               = wkedll.NewProc("wkeSetViewNetInterface")
	_wkeSetProxy                          = wkedll.NewProc("wkeSetProxy")
	_wkeSetViewProxy                      = wkedll.NewProc("wkeSetViewProxy")
	_wkeGetName                           = wkedll.NewProc("wkeGetName")
	_wkeSetName                           = wkedll.NewProc("wkeSetName")
	_wkeSetHandle                         = wkedll.NewProc("wkeSetHandle")
	_wkeSetHandleOffset                   = wkedll.NewProc("wkeSetHandleOffset")
	_wkeIsTransparent                     = wkedll.NewProc("wkeIsTransparent")
	_wkeSetTransparent                    = wkedll.NewProc("wkeSetTransparent")
	_wkeSetUserAgent                      = wkedll.NewProc("wkeSetUserAgent")
	_wkeGetUserAgent                      = wkedll.NewProc("wkeGetUserAgent")
	_wkeSetUserAgentW                     = wkedll.NewProc("wkeSetUserAgentW")
	_wkeShowDevtools                      = wkedll.NewProc("wkeShowDevtools")
	_wkeLoadW                             = wkedll.NewProc("wkeLoadW")
	_wkeLoadURL                           = wkedll.NewProc("wkeLoadURL")
	_wkeLoadURLW                          = wkedll.NewProc("wkeLoadURLW")
	_wkePostURL                           = wkedll.NewProc("wkePostURL")
	_wkePostURLW                          = wkedll.NewProc("wkePostURLW")
	_wkeLoadHTML                          = wkedll.NewProc("wkeLoadHTML")
	_wkeLoadHtmlWithBaseUrl               = wkedll.NewProc("wkeLoadHtmlWithBaseUrl")
	_wkeLoadHTMLW                         = wkedll.NewProc("wkeLoadHTMLW")
	_wkeLoadFile                          = wkedll.NewProc("wkeLoadFile")
	_wkeLoadFileW                         = wkedll.NewProc("wkeLoadFileW")
	_wkeGetURL                            = wkedll.NewProc("wkeGetURL")
	_wkeGetFrameUrl                       = wkedll.NewProc("wkeGetFrameUrl")
	_wkeIsLoading                         = wkedll.NewProc("wkeIsLoading")
	_wkeIsLoadingSucceeded                = wkedll.NewProc("wkeIsLoadingSucceeded")
	_wkeIsLoadingFailed                   = wkedll.NewProc("wkeIsLoadingFailed")
	_wkeIsLoadingCompleted                = wkedll.NewProc("wkeIsLoadingCompleted")
	_wkeIsDocumentReady                   = wkedll.NewProc("wkeIsDocumentReady")
	_wkeStopLoading                       = wkedll.NewProc("wkeStopLoading")
	_wkeReload                            = wkedll.NewProc("wkeReload")
	_wkeGoToOffset                        = wkedll.NewProc("wkeGoToOffset")
	_wkeGoToIndex                         = wkedll.NewProc("wkeGoToIndex")
	_wkeGetWebviewId                      = wkedll.NewProc("wkeGetWebviewId")
	_wkeIsWebviewAlive                    = wkedll.NewProc("wkeIsWebviewAlive")
	_wkeGetDocumentCompleteURL            = wkedll.NewProc("wkeGetDocumentCompleteURL")
	_wkeCreateMemBuf                      = wkedll.NewProc("wkeCreateMemBuf")
	_wkeFreeMemBuf                        = wkedll.NewProc("wkeFreeMemBuf")
	_wkeGetTitle                          = wkedll.NewProc("wkeGetTitle")
	_wkeGetTitleW                         = wkedll.NewProc("wkeGetTitleW")
	_wkeResize                            = wkedll.NewProc("wkeResize")
	_wkeGetWidth                          = wkedll.NewProc("wkeGetWidth")
	_wkeGetHeight                         = wkedll.NewProc("wkeGetHeight")
	_wkeGetContentWidth                   = wkedll.NewProc("wkeGetContentWidth")
	_wkeGetContentHeight                  = wkedll.NewProc("wkeGetContentHeight")
	_wkeSetDirty                          = wkedll.NewProc("wkeSetDirty")
	_wkeIsDirty                           = wkedll.NewProc("wkeIsDirty")
	_wkeAddDirtyArea                      = wkedll.NewProc("wkeAddDirtyArea")
	_wkeLayoutIfNeeded                    = wkedll.NewProc("wkeLayoutIfNeeded")
	_wkePaint2                            = wkedll.NewProc("wkePaint2")
	_wkePaint                             = wkedll.NewProc("wkePaint")
	_wkeRepaintIfNeeded                   = wkedll.NewProc("wkeRepaintIfNeeded")
	_wkeGetViewDC                         = wkedll.NewProc("wkeGetViewDC")
	_wkeGetHostHWND                       = wkedll.NewProc("wkeGetHostHWND")
	_wkeCanGoBack                         = wkedll.NewProc("wkeCanGoBack")
	_wkeGoBack                            = wkedll.NewProc("wkeGoBack")
	_wkeCanGoForward                      = wkedll.NewProc("wkeCanGoForward")
	_wkeGoForward                         = wkedll.NewProc("wkeGoForward")
	_wkeEditorSelectAll                   = wkedll.NewProc("wkeEditorSelectAll")
	_wkeEditorUnSelect                    = wkedll.NewProc("wkeEditorUnSelect")
	_wkeEditorCopy                        = wkedll.NewProc("wkeEditorCopy")
	_wkeEditorCut                         = wkedll.NewProc("wkeEditorCut")
	_wkeEditorPaste                       = wkedll.NewProc("wkeEditorPaste")
	_wkeEditorDelete                      = wkedll.NewProc("wkeEditorDelete")
	_wkeEditorUndo                        = wkedll.NewProc("wkeEditorUndo")
	_wkeEditorRedo                        = wkedll.NewProc("wkeEditorRedo")
	_wkeGetCookieW                        = wkedll.NewProc("wkeGetCookieW")
	_wkeGetCookie                         = wkedll.NewProc("wkeGetCookie")
	_wkeSetCookie                         = wkedll.NewProc("wkeSetCookie")
	_wkeVisitAllCookie                    = wkedll.NewProc("wkeVisitAllCookie")
	_wkePerformCookieCommand              = wkedll.NewProc("wkePerformCookieCommand")
	_wkeSetCookieEnabled                  = wkedll.NewProc("wkeSetCookieEnabled")
	_wkeIsCookieEnabled                   = wkedll.NewProc("wkeIsCookieEnabled")
	_wkeSetCookieJarPath                  = wkedll.NewProc("wkeSetCookieJarPath")
	_wkeSetCookieJarFullPath              = wkedll.NewProc("wkeSetCookieJarFullPath")
	_wkeSetLocalStorageFullPath           = wkedll.NewProc("wkeSetLocalStorageFullPath")
	_wkeAddPluginDirectory                = wkedll.NewProc("wkeAddPluginDirectory")
	_wkeSetMediaVolume                    = wkedll.NewProc("wkeSetMediaVolume")
	_wkeGetMediaVolume                    = wkedll.NewProc("wkeGetMediaVolume")
	_wkeFireMouseEvent                    = wkedll.NewProc("wkeFireMouseEvent")
	_wkeFireContextMenuEvent              = wkedll.NewProc("wkeFireContextMenuEvent")
	_wkeFireMouseWheelEvent               = wkedll.NewProc("wkeFireMouseWheelEvent")
	_wkeFireKeyUpEvent                    = wkedll.NewProc("wkeFireKeyUpEvent")
	_wkeFireKeyDownEvent                  = wkedll.NewProc("wkeFireKeyDownEvent")
	_wkeFireKeyPressEvent                 = wkedll.NewProc("wkeFireKeyPressEvent")
	_wkeFireWindowsMessage                = wkedll.NewProc("wkeFireWindowsMessage")
	_wkeSetFocus                          = wkedll.NewProc("wkeSetFocus")
	_wkeKillFocus                         = wkedll.NewProc("wkeKillFocus")
	_wkeRunJS                             = wkedll.NewProc("wkeRunJS")
	_wkeRunJSW                            = wkedll.NewProc("wkeRunJSW")
	_wkeGlobalExec                        = wkedll.NewProc("wkeGlobalExec")
	_wkeGetGlobalExecByFrame              = wkedll.NewProc("wkeGetGlobalExecByFrame")
	_wkeSleep                             = wkedll.NewProc("wkeSleep")
	_wkeWake                              = wkedll.NewProc("wkeWake")
	_wkeIsAwake                           = wkedll.NewProc("wkeIsAwake")
	_wkeSetZoomFactor                     = wkedll.NewProc("wkeSetZoomFactor")
	_wkeGetZoomFactor                     = wkedll.NewProc("wkeGetZoomFactor")
	_wkeSetEditable                       = wkedll.NewProc("wkeSetEditable")
	_wkeGetString                         = wkedll.NewProc("wkeGetString")
	_wkeGetStringW                        = wkedll.NewProc("wkeGetStringW")
	_wkeSetString                         = wkedll.NewProc("wkeSetString")
	_wkeSetStringW                        = wkedll.NewProc("wkeSetStringW")
	_wkeCreateString                      = wkedll.NewProc("wkeCreateString")
	_wkeCreateStringW                     = wkedll.NewProc("wkeCreateStringW")
	_wkeDeleteString                      = wkedll.NewProc("wkeDeleteString")
	_wkeGetWebViewForCurrentContext       = wkedll.NewProc("wkeGetWebViewForCurrentContext")
	_wkeSetUserKeyValue                   = wkedll.NewProc("wkeSetUserKeyValue")
	_wkeGetUserKeyValue                   = wkedll.NewProc("wkeGetUserKeyValue")
	_wkeGetCursorInfoType                 = wkedll.NewProc("wkeGetCursorInfoType")
	_wkeSetCursorInfoType                 = wkedll.NewProc("wkeSetCursorInfoType")
	_wkeSetDragFiles                      = wkedll.NewProc("wkeSetDragFiles")
	_wkeSetDeviceParameter                = wkedll.NewProc("wkeSetDeviceParameter")
	_wkeGetTempCallbackInfo               = wkedll.NewProc("wkeGetTempCallbackInfo")
	_wkeOnMouseOverUrlChanged             = wkedll.NewProc("wkeOnMouseOverUrlChanged")
	_wkeOnTitleChanged                    = wkedll.NewProc("wkeOnTitleChanged")
	_wkeOnURLChanged                      = wkedll.NewProc("wkeOnURLChanged")
	_wkeOnURLChanged2                     = wkedll.NewProc("wkeOnURLChanged2")
	_wkeOnPaintUpdated                    = wkedll.NewProc("wkeOnPaintUpdated")
	_wkeOnPaintBitUpdated                 = wkedll.NewProc("wkeOnPaintBitUpdated")
	_wkeOnAlertBox                        = wkedll.NewProc("wkeOnAlertBox")
	_wkeOnConfirmBox                      = wkedll.NewProc("wkeOnConfirmBox")
	_wkeOnPromptBox                       = wkedll.NewProc("wkeOnPromptBox")
	_wkeOnNavigation                      = wkedll.NewProc("wkeOnNavigation")
	_wkeOnCreateView                      = wkedll.NewProc("wkeOnCreateView")
	_wkeOnDocumentReady                   = wkedll.NewProc("wkeOnDocumentReady")
	_wkeOnDocumentReady2                  = wkedll.NewProc("wkeOnDocumentReady2")
	_wkeOnLoadingFinish                   = wkedll.NewProc("wkeOnLoadingFinish")
	_wkeOnDownload                        = wkedll.NewProc("wkeOnDownload")
	_wkeOnDownload2                       = wkedll.NewProc("wkeOnDownload2")
	_wkeOnConsole                         = wkedll.NewProc("wkeOnConsole")
	_wkeSetUIThreadCallback               = wkedll.NewProc("wkeSetUIThreadCallback")
	_wkeOnLoadUrlBegin                    = wkedll.NewProc("wkeOnLoadUrlBegin")
	_wkeOnLoadUrlEnd                      = wkedll.NewProc("wkeOnLoadUrlEnd")
	_wkeOnLoadUrlFail                     = wkedll.NewProc("wkeOnLoadUrlFail")
	_wkeOnDidCreateScriptContext          = wkedll.NewProc("wkeOnDidCreateScriptContext")
	_wkeOnWillReleaseScriptContext        = wkedll.NewProc("wkeOnWillReleaseScriptContext")
	_wkeOnWindowClosing                   = wkedll.NewProc("wkeOnWindowClosing")
	_wkeOnWindowDestroy                   = wkedll.NewProc("wkeOnWindowDestroy")
	_wkeOnDraggableRegionsChanged         = wkedll.NewProc("wkeOnDraggableRegionsChanged")
	_wkeOnWillMediaLoad                   = wkedll.NewProc("wkeOnWillMediaLoad")
	_wkeOnStartDragging                   = wkedll.NewProc("wkeOnStartDragging")
	_wkeOnPrint                           = wkedll.NewProc("wkeOnPrint")
	_wkeOnOtherLoad                       = wkedll.NewProc("wkeOnOtherLoad")
	_wkeOnContextMenuItemClick            = wkedll.NewProc("wkeOnContextMenuItemClick")
	_wkeIsProcessingUserGesture           = wkedll.NewProc("wkeIsProcessingUserGesture")
	_wkeNetSetMIMEType                    = wkedll.NewProc("wkeNetSetMIMEType")
	_wkeNetGetMIMEType                    = wkedll.NewProc("wkeNetGetMIMEType")
	_wkeNetSetHTTPHeaderField             = wkedll.NewProc("wkeNetSetHTTPHeaderField")
	_wkeNetGetHTTPHeaderField             = wkedll.NewProc("wkeNetGetHTTPHeaderField")
	_wkeNetGetHTTPHeaderFieldFromResponse = wkedll.NewProc("wkeNetGetHTTPHeaderFieldFromResponse")
	_wkeNetSetData                        = wkedll.NewProc("wkeNetSetData")
	_wkeNetHookRequest                    = wkedll.NewProc("wkeNetHookRequest")
	_wkeNetOnResponse                     = wkedll.NewProc("wkeNetOnResponse")
	_wkeNetGetRequestMethod               = wkedll.NewProc("wkeNetGetRequestMethod")
	_wkeNetGetFavicon                     = wkedll.NewProc("wkeNetGetFavicon")
	_wkeNetContinueJob                    = wkedll.NewProc("wkeNetContinueJob")
	_wkeNetGetUrlByJob                    = wkedll.NewProc("wkeNetGetUrlByJob")
	_wkeNetCancelRequest                  = wkedll.NewProc("wkeNetCancelRequest")
	_wkeNetHoldJobToAsynCommit            = wkedll.NewProc("wkeNetHoldJobToAsynCommit")
	_wkeNetChangeRequestUrl               = wkedll.NewProc("wkeNetChangeRequestUrl")
	_wkeNetCreateWebUrlRequest            = wkedll.NewProc("wkeNetCreateWebUrlRequest")
	_wkeNetCreateWebUrlRequest2           = wkedll.NewProc("wkeNetCreateWebUrlRequest2")
	_wkeNetCopyWebUrlRequest              = wkedll.NewProc("wkeNetCopyWebUrlRequest")
	_wkeNetDeleteBlinkWebURLRequestPtr    = wkedll.NewProc("wkeNetDeleteBlinkWebURLRequestPtr")
	_wkeNetAddHTTPHeaderFieldToUrlRequest = wkedll.NewProc("wkeNetAddHTTPHeaderFieldToUrlRequest")
	_wkeNetStartUrlRequest                = wkedll.NewProc("wkeNetStartUrlRequest")
	_wkeNetGetHttpStatusCode              = wkedll.NewProc("wkeNetGetHttpStatusCode")
	_wkeNetGetExpectedContentLength       = wkedll.NewProc("wkeNetGetExpectedContentLength")
	_wkeNetGetResponseUrl                 = wkedll.NewProc("wkeNetGetResponseUrl")
	_wkeNetCancelWebUrlRequest            = wkedll.NewProc("wkeNetCancelWebUrlRequest")
	_wkeNetGetPostBody                    = wkedll.NewProc("wkeNetGetPostBody")
	_wkeNetCreatePostBodyElements         = wkedll.NewProc("wkeNetCreatePostBodyElements")
	_wkeNetFreePostBodyElements           = wkedll.NewProc("wkeNetFreePostBodyElements")
	_wkeNetCreatePostBodyElement          = wkedll.NewProc("wkeNetCreatePostBodyElement")
	_wkeNetFreePostBodyElement            = wkedll.NewProc("wkeNetFreePostBodyElement")
	_wkeIsMainFrame                       = wkedll.NewProc("wkeIsMainFrame")
	_wkeIsWebRemoteFrame                  = wkedll.NewProc("wkeIsWebRemoteFrame")
	_wkeWebFrameGetMainFrame              = wkedll.NewProc("wkeWebFrameGetMainFrame")
	_wkeRunJsByFrame                      = wkedll.NewProc("wkeRunJsByFrame")
	_wkeInsertCSSByFrame                  = wkedll.NewProc("wkeInsertCSSByFrame")
	_wkeWebFrameGetMainWorldScriptContext = wkedll.NewProc("wkeWebFrameGetMainWorldScriptContext")
	_wkeGetBlinkMainThreadIsolate         = wkedll.NewProc("wkeGetBlinkMainThreadIsolate")
	_wkeCreateWebWindow                   = wkedll.NewProc("wkeCreateWebWindow")
	_wkeCreateWebCustomWindow             = wkedll.NewProc("wkeCreateWebCustomWindow")
	_wkeDestroyWebWindow                  = wkedll.NewProc("wkeDestroyWebWindow")
	_wkeGetWindowHandle                   = wkedll.NewProc("wkeGetWindowHandle")
	_wkeShowWindow                        = wkedll.NewProc("wkeShowWindow")
	_wkeEnableWindow                      = wkedll.NewProc("wkeEnableWindow")
	_wkeMoveWindow                        = wkedll.NewProc("wkeMoveWindow")
	_wkeMoveToCenter                      = wkedll.NewProc("wkeMoveToCenter")
	_wkeResizeWindow                      = wkedll.NewProc("wkeResizeWindow")
	_wkeDragTargetDragEnter               = wkedll.NewProc("wkeDragTargetDragEnter")
	_wkeDragTargetDragOver                = wkedll.NewProc("wkeDragTargetDragOver")
	_wkeDragTargetDragLeave               = wkedll.NewProc("wkeDragTargetDragLeave")
	_wkeDragTargetDrop                    = wkedll.NewProc("wkeDragTargetDrop")
	_wkeDragTargetEnd                     = wkedll.NewProc("wkeDragTargetEnd")
	_wkeUtilSetUiCallback                 = wkedll.NewProc("wkeUtilSetUiCallback")
	_wkeUtilSerializeToMHTML              = wkedll.NewProc("wkeUtilSerializeToMHTML")
	_wkeUtilPrintToPdf                    = wkedll.NewProc("wkeUtilPrintToPdf")
	_wkePrintToBitmap                     = wkedll.NewProc("wkePrintToBitmap")
	_wkeUtilRelasePrintPdfDatas           = wkedll.NewProc("wkeUtilRelasePrintPdfDatas")
	_wkeSetWindowTitle                    = wkedll.NewProc("wkeSetWindowTitle")
	_wkeSetWindowTitleW                   = wkedll.NewProc("wkeSetWindowTitleW")
	_wkeNodeOnCreateProcess               = wkedll.NewProc("wkeNodeOnCreateProcess")
	_wkeOnPluginFind                      = wkedll.NewProc("wkeOnPluginFind")
	_wkeAddNpapiPlugin                    = wkedll.NewProc("wkeAddNpapiPlugin")
	_wkeGetWebViewByNData                 = wkedll.NewProc("wkeGetWebViewByNData")
	_wkeRegisterEmbedderCustomElement     = wkedll.NewProc("wkeRegisterEmbedderCustomElement")
	_wkeSetMediaPlayerFactory             = wkedll.NewProc("wkeSetMediaPlayerFactory")
	_wkeGetContentAsMarkup                = wkedll.NewProc("wkeGetContentAsMarkup")
	_wkeUtilDecodeURLEscape               = wkedll.NewProc("wkeUtilDecodeURLEscape")
	_wkeUtilEncodeURLEscape               = wkedll.NewProc("wkeUtilEncodeURLEscape")
	_wkeUtilBase64Encode                  = wkedll.NewProc("wkeUtilBase64Encode")
	_wkeUtilBase64Decode                  = wkedll.NewProc("wkeUtilBase64Decode")
	_wkeUtilCreateV8Snapshot              = wkedll.NewProc("wkeUtilCreateV8Snapshot")
	_wkeRunMessageLoop                    = wkedll.NewProc("wkeRunMessageLoop")
	_wkeSaveMemoryCache                   = wkedll.NewProc("wkeSaveMemoryCache")
)

func wkeShutdown() {
	_wkeShutdown.Call()
}

func wkeVersion() uint {
	r, _, _ := _wkeVersion.Call()
	return uint(r)
}

func wkeVersionString() string {
	r, _, _ := _wkeVersionString.Call()
	return GoAStr(r)
}

func wkeGC(webView WkeWebView, intervalSec int32) {
	_wkeGC.Call(uintptr(webView), uintptr(intervalSec))
}

func wkeSetResourceGc(webView WkeWebView, intervalSec int32) {
	_wkeSetResourceGc.Call(uintptr(webView), uintptr(intervalSec))
}

func wkeSetFileSystem(pfnOpen WKE_FILE_OPEN, pfnClose WKE_FILE_CLOSE, pfnSize WKE_FILE_SIZE, pfnRead WKE_FILE_READ, pfnSeek WKE_FILE_SEEK) {
	_wkeSetFileSystem.Call(uintptr(pfnOpen), uintptr(pfnClose), uintptr(pfnSize), uintptr(pfnRead), uintptr(pfnSeek))
}

func wkeWebViewName(webView WkeWebView) string {
	r, _, _ := _wkeWebViewName.Call(uintptr(webView))
	return GoAStr(r)
}

func wkeSetWebViewName(webView WkeWebView, name string) {
	_wkeSetWebViewName.Call(uintptr(webView), CAStr(name))
}

func wkeIsLoaded(webView WkeWebView) bool {
	r, _, _ := _wkeIsLoaded.Call(uintptr(webView))
	return GoBool(r)
}

func wkeIsLoadFailed(webView WkeWebView) bool {
	r, _, _ := _wkeIsLoadFailed.Call(uintptr(webView))
	return GoBool(r)
}

func wkeIsLoadComplete(webView WkeWebView) bool {
	r, _, _ := _wkeIsLoadComplete.Call(uintptr(webView))
	return GoBool(r)
}

func wkeGetSource(webView WkeWebView) string {
	r, _, _ := _wkeGetSource.Call(uintptr(webView))
	return GoAStr(r)
}

func wkeTitle(webView WkeWebView) string {
	r, _, _ := _wkeTitle.Call(uintptr(webView))
	return GoAStr(r)
}

func wkeTitleW(webView WkeWebView) string {
	r, _, _ := _wkeTitleW.Call(uintptr(webView))
	return GoWStr(r)
}

func wkeWidth(webView WkeWebView) int {
	r, _, _ := _wkeWidth.Call(uintptr(webView))
	return int(r)
}

func wkeHeight(webView WkeWebView) int {
	r, _, _ := _wkeHeight.Call(uintptr(webView))
	return int(r)
}

func wkeContentsWidth(webView WkeWebView) int {
	r, _, _ := _wkeContentsWidth.Call(uintptr(webView))
	return int(r)
}

func wkeContentsHeight(webView WkeWebView) int {
	r, _, _ := _wkeContentsHeight.Call(uintptr(webView))
	return int(r)
}

func wkeSelectAll(webView WkeWebView) {
	_wkeSelectAll.Call(uintptr(webView))
}

func wkeCopy(webView WkeWebView) {
	_wkeCopy.Call(uintptr(webView))
}

func wkeCut(webView WkeWebView) {
	_wkeCut.Call(uintptr(webView))
}

func wkePaste(webView WkeWebView) {
	_wkePaste.Call(uintptr(webView))
}

func wkeDelete(webView WkeWebView) {
	_wkeDelete.Call(uintptr(webView))
}

func wkeCookieEnabled(webView WkeWebView) bool {
	r, _, _ := _wkeCookieEnabled.Call(uintptr(webView))
	return GoBool(r)
}

func wkeMediaVolume(webView WkeWebView) float32 {
	_wkeMediaVolume.Call(uintptr(webView))
	return floatpatch.Getfloat32()
}

func wkeMouseEvent(webView WkeWebView, message uint, x int, y int, flags uint) bool {
	r, _, _ := _wkeMouseEvent.Call(uintptr(webView), uintptr(message), uintptr(x), uintptr(y), uintptr(flags))
	return GoBool(r)
}

func wkeContextMenuEvent(webView WkeWebView, x int, y int, flags uint) bool {
	r, _, _ := _wkeContextMenuEvent.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(flags))
	return GoBool(r)
}

func wkeMouseWheel(webView WkeWebView, x int, y int, delta int, flags uint) bool {
	r, _, _ := _wkeMouseWheel.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(delta), uintptr(flags))
	return GoBool(r)
}

func wkeKeyUp(webView WkeWebView, virtualKeyCode uint, flags uint, systemKey bool) bool {
	r, _, _ := _wkeKeyUp.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), CBool(systemKey))
	return GoBool(r)
}

func wkeKeyDown(webView WkeWebView, virtualKeyCode uint, flags uint, systemKey bool) bool {
	r, _, _ := _wkeKeyDown.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), CBool(systemKey))
	return GoBool(r)
}

func wkeKeyPress(webView WkeWebView, virtualKeyCode uint, flags uint, systemKey bool) bool {
	r, _, _ := _wkeKeyPress.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), CBool(systemKey))
	return GoBool(r)
}

func wkeFocus(webView WkeWebView) {
	_wkeFocus.Call(uintptr(webView))
}

func wkeUnfocus(webView WkeWebView) {
	_wkeUnfocus.Call(uintptr(webView))
}

func wkeAwaken(webView WkeWebView) {
	_wkeAwaken.Call(uintptr(webView))
}

func wkeZoomFactor(webView WkeWebView) float32 {
	_wkeZoomFactor.Call(uintptr(webView))
	return floatpatch.Getfloat32()
}

func wkeSetClientHandler(webView WkeWebView, handler *WkeClientHandler) {
	_wkeSetClientHandler.Call(uintptr(webView), uintptr(unsafe.Pointer(handler)))
}

func wkeGetClientHandler(webView WkeWebView) *WkeClientHandler {
	r, _, _ := _wkeGetClientHandler.Call(uintptr(webView))
	return (*WkeClientHandler)(unsafe.Pointer(r))
}

func wkeToString(string WkeString) string {
	r, _, _ := _wkeToString.Call(uintptr(string))
	return GoAStr(r)
}

func wkeToStringW(string WkeString) string {
	r, _, _ := _wkeToStringW.Call(uintptr(string))
	return GoWStr(r)
}

func jsToString(es JsExecState, v JsValue) string {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsToString.Call(uintptr(es), v1, v2)
	} else {
		r, _, _ = _jsToString.Call(uintptr(es), uintptr(v))
	}
	return GoAStr(r)
}

func jsToStringW(es JsExecState, v JsValue) string {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsToStringW.Call(uintptr(es), v1, v2)
	} else {
		r, _, _ = _jsToStringW.Call(uintptr(es), uintptr(v))
	}
	return GoWStr(r)
}

func wkeConfigure(settings *WkeSettings) {
	_wkeConfigure.Call(uintptr(unsafe.Pointer(settings)))
}

func wkeIsInitialize() bool {
	r, _, _ := _wkeIsInitialize.Call()
	return GoBool(r)
}

func wkeSetViewSettings(webView WkeWebView, settings *WkeViewSettings) {
	_wkeSetViewSettings.Call(uintptr(webView), uintptr(unsafe.Pointer(settings)))
}

func wkeSetDebugConfig(webView WkeWebView, debugString string, param string) {
	_wkeSetDebugConfig.Call(uintptr(webView), CAStr(debugString), CAStr(param))
}

func wkeGetDebugConfig(webView WkeWebView, debugString string) uintptr {
	r, _, _ := _wkeGetDebugConfig.Call(uintptr(webView), CAStr(debugString))
	return r
}

func wkeFinalize() {
	_wkeFinalize.Call()
}

func wkeUpdate() {
	_wkeUpdate.Call()
}

func wkeGetVersion() uint {
	r, _, _ := _wkeGetVersion.Call()
	return uint(r)
}

func wkeGetVersionString() string {
	r, _, _ := _wkeGetVersionString.Call()
	return GoAStr(r)
}

func wkeCreateWebView() WkeWebView {
	r, _, _ := _wkeCreateWebView.Call()
	return WkeWebView(r)
}

func wkeDestroyWebView(webView WkeWebView) {
	_wkeDestroyWebView.Call(uintptr(webView))
}

func wkeSetMemoryCacheEnable(webView WkeWebView, b bool) {
	_wkeSetMemoryCacheEnable.Call(uintptr(webView), CBool(b))
}

func wkeSetMouseEnabled(webView WkeWebView, b bool) {
	_wkeSetMouseEnabled.Call(uintptr(webView), CBool(b))
}

func wkeSetTouchEnabled(webView WkeWebView, b bool) {
	_wkeSetTouchEnabled.Call(uintptr(webView), CBool(b))
}

func wkeSetContextMenuEnabled(webView WkeWebView, b bool) {
	_wkeSetContextMenuEnabled.Call(uintptr(webView), CBool(b))
}

func wkeSetNavigationToNewWindowEnable(webView WkeWebView, b bool) {
	_wkeSetNavigationToNewWindowEnable.Call(uintptr(webView), CBool(b))
}

func wkeSetCspCheckEnable(webView WkeWebView, b bool) {
	_wkeSetCspCheckEnable.Call(uintptr(webView), CBool(b))
}

func wkeSetNpapiPluginsEnabled(webView WkeWebView, b bool) {
	_wkeSetNpapiPluginsEnabled.Call(uintptr(webView), CBool(b))
}

func wkeSetHeadlessEnabled(webView WkeWebView, b bool) {
	_wkeSetHeadlessEnabled.Call(uintptr(webView), CBool(b))
}

func wkeSetDragEnable(webView WkeWebView, b bool) {
	_wkeSetDragEnable.Call(uintptr(webView), CBool(b))
}

func wkeSetDragDropEnable(webView WkeWebView, b bool) {
	_wkeSetDragDropEnable.Call(uintptr(webView), CBool(b))
}

func wkeSetContextMenuItemShow(webView WkeWebView, item WkeMenuItemId, isShow bool) {
	_wkeSetContextMenuItemShow.Call(uintptr(webView), uintptr(item), CBool(isShow))
}

func wkeSetLanguage(webView WkeWebView, language string) {
	_wkeSetLanguage.Call(uintptr(webView), CAStr(language))
}

func wkeSetViewNetInterface(webView WkeWebView, netInterface string) {
	_wkeSetViewNetInterface.Call(uintptr(webView), CAStr(netInterface))
}

func wkeSetProxy(proxy *WkeProxy) {
	_wkeSetProxy.Call(uintptr(unsafe.Pointer(proxy)))
}

func wkeSetViewProxy(webView WkeWebView, proxy *WkeProxy) {
	_wkeSetViewProxy.Call(uintptr(webView), uintptr(unsafe.Pointer(proxy)))
}

func wkeGetName(webView WkeWebView) string {
	r, _, _ := _wkeGetName.Call(uintptr(webView))
	return GoAStr(r)
}

func wkeSetName(webView WkeWebView, name string) {
	_wkeSetName.Call(uintptr(webView), CAStr(name))
}

func wkeSetHandle(webView WkeWebView, wnd types.HWND) {
	_wkeSetHandle.Call(uintptr(webView), uintptr(wnd))
}

func wkeSetHandleOffset(webView WkeWebView, x int, y int) {
	_wkeSetHandleOffset.Call(uintptr(webView), uintptr(x), uintptr(y))
}

func wkeIsTransparent(webView WkeWebView) bool {
	r, _, _ := _wkeIsTransparent.Call(uintptr(webView))
	return GoBool(r)
}

func wkeSetTransparent(webView WkeWebView, transparent bool) {
	_wkeSetTransparent.Call(uintptr(webView), CBool(transparent))
}

func wkeSetUserAgent(webView WkeWebView, userAgent string) {
	_wkeSetUserAgent.Call(uintptr(webView), CAStr(userAgent))
}

func wkeGetUserAgent(webView WkeWebView) string {
	r, _, _ := _wkeGetUserAgent.Call(uintptr(webView))
	return GoAStr(r)
}

func wkeSetUserAgentW(webView WkeWebView, userAgent string) {
	_wkeSetUserAgentW.Call(uintptr(webView), CWStr(userAgent))
}

func wkeShowDevtools(webView WkeWebView, path string, callback WkeOnShowDevtoolsCallback, param uintptr) {
	_wkeShowDevtools.Call(uintptr(webView), CAStr(path), uintptr(callback), param)
}

func wkeLoadW(webView WkeWebView, url string) {
	_wkeLoadW.Call(uintptr(webView), CWStr(url))
}

func wkeLoadURL(webView WkeWebView, url string) {
	_wkeLoadURL.Call(uintptr(webView), CAStr(url))
}

func wkeLoadURLW(webView WkeWebView, url string) {
	_wkeLoadURLW.Call(uintptr(webView), CWStr(url))
}

func wkePostURL(wkeView WkeWebView, url string, postData string, postLen int) {
	_wkePostURL.Call(uintptr(wkeView), CAStr(url), CAStr(postData), uintptr(postLen))
}

func wkePostURLW(wkeView WkeWebView, url string, postData string, postLen int) {
	_wkePostURLW.Call(uintptr(wkeView), CWStr(url), CWStr(postData), uintptr(postLen))
}

func wkeLoadHTML(webView WkeWebView, html string) {
	_wkeLoadHTML.Call(uintptr(webView), CAStr(html))
}

func wkeLoadHtmlWithBaseUrl(webView WkeWebView, html string, baseUrl string) {
	_wkeLoadHtmlWithBaseUrl.Call(uintptr(webView), CAStr(html), CAStr(baseUrl))
}

func wkeLoadHTMLW(webView WkeWebView, html string) {
	_wkeLoadHTMLW.Call(uintptr(webView), CWStr(html))
}

func wkeLoadFile(webView WkeWebView, filename string) {
	_wkeLoadFile.Call(uintptr(webView), CAStr(filename))
}

func wkeLoadFileW(webView WkeWebView, filename string) {
	_wkeLoadFileW.Call(uintptr(webView), CWStr(filename))
}

func wkeGetURL(webView WkeWebView) string {
	r, _, _ := _wkeGetURL.Call(uintptr(webView))
	return GoAStr(r)
}

func wkeGetFrameUrl(webView WkeWebView, frameId WkeWebFrameHandle) string {
	r, _, _ := _wkeGetFrameUrl.Call(uintptr(webView), uintptr(frameId))
	return GoAStr(r)
}

func wkeIsLoading(webView WkeWebView) bool {
	r, _, _ := _wkeIsLoading.Call(uintptr(webView))
	return GoBool(r)
}

func wkeIsLoadingSucceeded(webView WkeWebView) bool {
	r, _, _ := _wkeIsLoadingSucceeded.Call(uintptr(webView))
	return GoBool(r)
}

func wkeIsLoadingFailed(webView WkeWebView) bool {
	r, _, _ := _wkeIsLoadingFailed.Call(uintptr(webView))
	return GoBool(r)
}

func wkeIsLoadingCompleted(webView WkeWebView) bool {
	r, _, _ := _wkeIsLoadingCompleted.Call(uintptr(webView))
	return GoBool(r)
}

func wkeIsDocumentReady(webView WkeWebView) bool {
	r, _, _ := _wkeIsDocumentReady.Call(uintptr(webView))
	return GoBool(r)
}

func wkeStopLoading(webView WkeWebView) {
	_wkeStopLoading.Call(uintptr(webView))
}

func wkeReload(webView WkeWebView) {
	_wkeReload.Call(uintptr(webView))
}

func wkeGoToOffset(webView WkeWebView, offset int) {
	_wkeGoToOffset.Call(uintptr(webView), uintptr(offset))
}

func wkeGoToIndex(webView WkeWebView, index int) {
	_wkeGoToIndex.Call(uintptr(webView), uintptr(index))
}

func wkeGetWebviewId(webView WkeWebView) int {
	r, _, _ := _wkeGetWebviewId.Call(uintptr(webView))
	return int(r)
}

func wkeIsWebviewAlive(id int) bool {
	r, _, _ := _wkeIsWebviewAlive.Call(uintptr(id))
	return GoBool(r)
}

func wkeGetDocumentCompleteURL(webView WkeWebView, frameId WkeWebFrameHandle, partialURL string) string {
	r, _, _ := _wkeGetDocumentCompleteURL.Call(uintptr(webView), uintptr(frameId), CAStr(partialURL))
	return GoAStr(r)
}

func wkeCreateMemBuf(webView WkeWebView, buf uintptr, length uint) *WkeMemBuf {
	r, _, _ := _wkeCreateMemBuf.Call(uintptr(webView), buf, uintptr(length))
	return (*WkeMemBuf)(unsafe.Pointer(r))
}

func wkeFreeMemBuf(buf *WkeMemBuf) {
	_wkeFreeMemBuf.Call(uintptr(unsafe.Pointer(buf)))
}

func wkeGetTitle(webView WkeWebView) string {
	r, _, _ := _wkeGetTitle.Call(uintptr(webView))
	return GoAStr(r)
}

func wkeGetTitleW(webView WkeWebView) string {
	r, _, _ := _wkeGetTitleW.Call(uintptr(webView))
	return GoWStr(r)
}

func wkeResize(webView WkeWebView, w int, h int) {
	_wkeResize.Call(uintptr(webView), uintptr(w), uintptr(h))
}

func wkeGetWidth(webView WkeWebView) int {
	r, _, _ := _wkeGetWidth.Call(uintptr(webView))
	return int(r)
}

func wkeGetHeight(webView WkeWebView) int {
	r, _, _ := _wkeGetHeight.Call(uintptr(webView))
	return int(r)
}

func wkeGetContentWidth(webView WkeWebView) int {
	r, _, _ := _wkeGetContentWidth.Call(uintptr(webView))
	return int(r)
}

func wkeGetContentHeight(webView WkeWebView) int {
	r, _, _ := _wkeGetContentHeight.Call(uintptr(webView))
	return int(r)
}

func wkeSetDirty(webView WkeWebView, dirty bool) {
	_wkeSetDirty.Call(uintptr(webView), CBool(dirty))
}

func wkeIsDirty(webView WkeWebView) bool {
	r, _, _ := _wkeIsDirty.Call(uintptr(webView))
	return GoBool(r)
}

func wkeAddDirtyArea(webView WkeWebView, x int, y int, w int, h int) {
	_wkeAddDirtyArea.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(w), uintptr(h))
}

func wkeLayoutIfNeeded(webView WkeWebView) {
	_wkeLayoutIfNeeded.Call(uintptr(webView))
}

func wkePaint2(webView WkeWebView, bits uintptr, bufWid int, bufHei int, xDst int, yDst int, w int, h int, xSrc int, ySrc int, bCopyAlpha bool) {
	_wkePaint2.Call(uintptr(webView), bits, uintptr(bufWid), uintptr(bufHei), uintptr(xDst), uintptr(yDst), uintptr(w), uintptr(h), uintptr(xSrc), uintptr(ySrc), CBool(bCopyAlpha))
}

func wkePaint(webView WkeWebView, bits uintptr, pitch int) {
	_wkePaint.Call(uintptr(webView), bits, uintptr(pitch))
}

func wkeRepaintIfNeeded(webView WkeWebView) {
	_wkeRepaintIfNeeded.Call(uintptr(webView))
}

func wkeGetViewDC(webView WkeWebView) types.HDC {
	r, _, _ := _wkeGetViewDC.Call(uintptr(webView))
	return types.HDC(r)
}

func wkeGetHostHWND(webView WkeWebView) types.HWND {
	r, _, _ := _wkeGetHostHWND.Call(uintptr(webView))
	return types.HWND(r)
}

func wkeCanGoBack(webView WkeWebView) bool {
	r, _, _ := _wkeCanGoBack.Call(uintptr(webView))
	return GoBool(r)
}

func wkeGoBack(webView WkeWebView) bool {
	r, _, _ := _wkeGoBack.Call(uintptr(webView))
	return GoBool(r)
}

func wkeCanGoForward(webView WkeWebView) bool {
	r, _, _ := _wkeCanGoForward.Call(uintptr(webView))
	return GoBool(r)
}

func wkeGoForward(webView WkeWebView) bool {
	r, _, _ := _wkeGoForward.Call(uintptr(webView))
	return GoBool(r)
}

func wkeEditorSelectAll(webView WkeWebView) {
	_wkeEditorSelectAll.Call(uintptr(webView))
}

func wkeEditorUnSelect(webView WkeWebView) {
	_wkeEditorUnSelect.Call(uintptr(webView))
}

func wkeEditorCopy(webView WkeWebView) {
	_wkeEditorCopy.Call(uintptr(webView))
}

func wkeEditorCut(webView WkeWebView) {
	_wkeEditorCut.Call(uintptr(webView))
}

func wkeEditorPaste(webView WkeWebView) {
	_wkeEditorPaste.Call(uintptr(webView))
}

func wkeEditorDelete(webView WkeWebView) {
	_wkeEditorDelete.Call(uintptr(webView))
}

func wkeEditorUndo(webView WkeWebView) {
	_wkeEditorUndo.Call(uintptr(webView))
}

func wkeEditorRedo(webView WkeWebView) {
	_wkeEditorRedo.Call(uintptr(webView))
}

func wkeGetCookieW(webView WkeWebView) string {
	r, _, _ := _wkeGetCookieW.Call(uintptr(webView))
	return GoWStr(r)
}

func wkeGetCookie(webView WkeWebView) string {
	r, _, _ := _wkeGetCookie.Call(uintptr(webView))
	return GoAStr(r)
}

func wkeSetCookie(webView WkeWebView, url string, cookie string) {
	_wkeSetCookie.Call(uintptr(webView), CAStr(url), CAStr(cookie))
}

func wkeVisitAllCookie(webView WkeWebView, params uintptr, visitor WkeCookieVisitor) {
	_wkeVisitAllCookie.Call(uintptr(webView), params, uintptr(visitor))
}

func wkePerformCookieCommand(webView WkeWebView, command WkeCookieCommand) {
	_wkePerformCookieCommand.Call(uintptr(webView), uintptr(command))
}

func wkeSetCookieEnabled(webView WkeWebView, enable bool) {
	_wkeSetCookieEnabled.Call(uintptr(webView), CBool(enable))
}

func wkeIsCookieEnabled(webView WkeWebView) bool {
	r, _, _ := _wkeIsCookieEnabled.Call(uintptr(webView))
	return GoBool(r)
}

func wkeSetCookieJarPath(webView WkeWebView, path string) {
	_wkeSetCookieJarPath.Call(uintptr(webView), CAStr(path))
}

func wkeSetCookieJarFullPath(webView WkeWebView, path string) {
	_wkeSetCookieJarFullPath.Call(uintptr(webView), CAStr(path))
}

func wkeSetLocalStorageFullPath(webView WkeWebView, path string) {
	_wkeSetLocalStorageFullPath.Call(uintptr(webView), CAStr(path))
}

func wkeAddPluginDirectory(webView WkeWebView, path string) {
	_wkeAddPluginDirectory.Call(uintptr(webView), CAStr(path))
}

func wkeSetMediaVolume(webView WkeWebView, volume float32) {
	_wkeSetMediaVolume.Call(uintptr(webView), uintptr(volume))
}

func wkeGetMediaVolume(webView WkeWebView) float32 {
	_wkeGetMediaVolume.Call(uintptr(webView))
	return floatpatch.Getfloat32()
}

func wkeFireMouseEvent(webView WkeWebView, message uint, x int, y int, flags uint) bool {
	r, _, _ := _wkeFireMouseEvent.Call(uintptr(webView), uintptr(message), uintptr(x), uintptr(y), uintptr(flags))
	return GoBool(r)
}

func wkeFireContextMenuEvent(webView WkeWebView, x int, y int, flags uint) bool {
	r, _, _ := _wkeFireContextMenuEvent.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(flags))
	return GoBool(r)
}

func wkeFireMouseWheelEvent(webView WkeWebView, x int, y int, delta int, flags uint) bool {
	r, _, _ := _wkeFireMouseWheelEvent.Call(uintptr(webView), uintptr(x), uintptr(y), uintptr(delta), uintptr(flags))
	return GoBool(r)
}

func wkeFireKeyUpEvent(webView WkeWebView, virtualKeyCode uint, flags uint, systemKey bool) bool {
	r, _, _ := _wkeFireKeyUpEvent.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), CBool(systemKey))
	return GoBool(r)
}

func wkeFireKeyDownEvent(webView WkeWebView, virtualKeyCode uint, flags uint, systemKey bool) bool {
	r, _, _ := _wkeFireKeyDownEvent.Call(uintptr(webView), uintptr(virtualKeyCode), uintptr(flags), CBool(systemKey))
	return GoBool(r)
}

func wkeFireKeyPressEvent(webView WkeWebView, charCode uint, flags uint, systemKey bool) bool {
	r, _, _ := _wkeFireKeyPressEvent.Call(uintptr(webView), uintptr(charCode), uintptr(flags), CBool(systemKey))
	return GoBool(r)
}

func wkeFireWindowsMessage(webView WkeWebView, hWnd types.HWND, message uint, wParam uintptr, lParam uintptr, result *uintptr) bool {
	r, _, _ := _wkeFireWindowsMessage.Call(uintptr(webView), uintptr(hWnd), uintptr(message), wParam, lParam, uintptr(unsafe.Pointer(result)))
	return GoBool(r)
}

func wkeSetFocus(webView WkeWebView) {
	_wkeSetFocus.Call(uintptr(webView))
}

func wkeKillFocus(webView WkeWebView) {
	_wkeKillFocus.Call(uintptr(webView))
}

func wkeRunJS(webView WkeWebView, script string) JsValue {
	r, r2, _ := _wkeRunJS.Call(uintptr(webView), CAStr(script))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func wkeRunJSW(webView WkeWebView, script string) JsValue {
	r, r2, _ := _wkeRunJSW.Call(uintptr(webView), CWStr(script))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func wkeGlobalExec(webView WkeWebView) JsExecState {
	r, _, _ := _wkeGlobalExec.Call(uintptr(webView))
	return JsExecState(r)
}

func wkeGetGlobalExecByFrame(webView WkeWebView, frameId WkeWebFrameHandle) JsExecState {
	r, _, _ := _wkeGetGlobalExecByFrame.Call(uintptr(webView), uintptr(frameId))
	return JsExecState(r)
}

func wkeSleep(webView WkeWebView) {
	_wkeSleep.Call(uintptr(webView))
}

func wkeWake(webView WkeWebView) {
	_wkeWake.Call(uintptr(webView))
}

func wkeIsAwake(webView WkeWebView) bool {
	r, _, _ := _wkeIsAwake.Call(uintptr(webView))
	return GoBool(r)
}

func wkeSetZoomFactor(webView WkeWebView, factor float32) {
	_wkeSetZoomFactor.Call(uintptr(webView), uintptr(factor))
}

func wkeGetZoomFactor(webView WkeWebView) float32 {
	_wkeGetZoomFactor.Call(uintptr(webView))
	return floatpatch.Getfloat32()
}

func wkeSetEditable(webView WkeWebView, editable bool) {
	_wkeSetEditable.Call(uintptr(webView), CBool(editable))
}

func wkeGetString(str WkeString) string {
	r, _, _ := _wkeGetString.Call(uintptr(str))
	return GoAStr(r)
}

func wkeGetStringW(str WkeString) string {
	r, _, _ := _wkeGetStringW.Call(uintptr(str))
	return GoWStr(r)
}

func wkeSetString(string WkeString, str string, len uint) {
	_wkeSetString.Call(uintptr(string), CAStr(str), uintptr(len))
}

func wkeSetStringW(string WkeString, str string, len uint) {
	_wkeSetStringW.Call(uintptr(string), CWStr(str), uintptr(len))
}

func wkeCreateString(str string, len uint) WkeString {
	r, _, _ := _wkeCreateString.Call(CAStr(str), uintptr(len))
	return WkeString(r)
}

func wkeCreateStringW(str string, len uint) WkeString {
	r, _, _ := _wkeCreateStringW.Call(CWStr(str), uintptr(len))
	return WkeString(r)
}

func wkeDeleteString(str WkeString) {
	_wkeDeleteString.Call(uintptr(str))
}

func wkeGetWebViewForCurrentContext() WkeWebView {
	r, _, _ := _wkeGetWebViewForCurrentContext.Call()
	return WkeWebView(r)
}

func wkeSetUserKeyValue(webView WkeWebView, key string, value uintptr) {
	_wkeSetUserKeyValue.Call(uintptr(webView), CAStr(key), value)
}

func wkeGetUserKeyValue(webView WkeWebView, key string) uintptr {
	r, _, _ := _wkeGetUserKeyValue.Call(uintptr(webView), CAStr(key))
	return r
}

func wkeGetCursorInfoType(webView WkeWebView) int {
	r, _, _ := _wkeGetCursorInfoType.Call(uintptr(webView))
	return int(r)
}

func wkeSetCursorInfoType(webView WkeWebView, aType int) {
	_wkeSetCursorInfoType.Call(uintptr(webView), uintptr(aType))
}

func wkeSetDragFiles(webView WkeWebView, clintPos *types.TPoint, screenPos *types.TPoint, files []WkeString, filesCount int) {
	_wkeSetDragFiles.Call(uintptr(webView), uintptr(unsafe.Pointer(clintPos)), uintptr(unsafe.Pointer(screenPos)), uintptr(unsafe.Pointer(&files[0])), uintptr(filesCount))
}

func wkeSetDeviceParameter(webView WkeWebView, device string, paramStr string, paramInt int, paramFloat float32) {
	_wkeSetDeviceParameter.Call(uintptr(webView), CAStr(device), CAStr(paramStr), uintptr(paramInt), uintptr(paramFloat))
}

func wkeGetTempCallbackInfo(webView WkeWebView) *WkeTempCallbackInfo {
	r, _, _ := _wkeGetTempCallbackInfo.Call(uintptr(webView))
	return (*WkeTempCallbackInfo)(unsafe.Pointer(r))
}

func wkeOnMouseOverUrlChanged(webView WkeWebView, callback WkeTitleChangedCallback, callbackParam uintptr) {
	_wkeOnMouseOverUrlChanged.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnTitleChanged(webView WkeWebView, callback WkeTitleChangedCallback, callbackParam uintptr) {
	_wkeOnTitleChanged.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnURLChanged(webView WkeWebView, callback WkeURLChangedCallback, callbackParam uintptr) {
	_wkeOnURLChanged.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnURLChanged2(webView WkeWebView, callback WkeURLChangedCallback2, callbackParam uintptr) {
	_wkeOnURLChanged2.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnPaintUpdated(webView WkeWebView, callback WkePaintUpdatedCallback, callbackParam uintptr) {
	_wkeOnPaintUpdated.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnPaintBitUpdated(webView WkeWebView, callback WkePaintBitUpdatedCallback, callbackParam uintptr) {
	_wkeOnPaintBitUpdated.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnAlertBox(webView WkeWebView, callback WkeAlertBoxCallback, callbackParam uintptr) {
	_wkeOnAlertBox.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnConfirmBox(webView WkeWebView, callback WkeConfirmBoxCallback, callbackParam uintptr) {
	_wkeOnConfirmBox.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnPromptBox(webView WkeWebView, callback WkePromptBoxCallback, callbackParam uintptr) {
	_wkeOnPromptBox.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnNavigation(webView WkeWebView, callback WkeNavigationCallback, param uintptr) {
	_wkeOnNavigation.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnCreateView(webView WkeWebView, callback WkeCreateViewCallback, param uintptr) {
	_wkeOnCreateView.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnDocumentReady(webView WkeWebView, callback WkeDocumentReadyCallback, param uintptr) {
	_wkeOnDocumentReady.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnDocumentReady2(webView WkeWebView, callback WkeDocumentReady2Callback, param uintptr) {
	_wkeOnDocumentReady2.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnLoadingFinish(webView WkeWebView, callback WkeLoadingFinishCallback, param uintptr) {
	_wkeOnLoadingFinish.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnDownload(webView WkeWebView, callback WkeDownloadCallback, param uintptr) {
	_wkeOnDownload.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnDownload2(webView WkeWebView, callback WkeDownload2Callback, param uintptr) {
	_wkeOnDownload2.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnConsole(webView WkeWebView, callback WkeConsoleCallback, param uintptr) {
	_wkeOnConsole.Call(uintptr(webView), uintptr(callback), param)
}

func wkeSetUIThreadCallback(webView WkeWebView, callback WkeCallUiThread, param uintptr) {
	_wkeSetUIThreadCallback.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnLoadUrlBegin(webView WkeWebView, callback WkeLoadUrlBeginCallback, callbackParam uintptr) {
	_wkeOnLoadUrlBegin.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnLoadUrlEnd(webView WkeWebView, callback WkeLoadUrlEndCallback, callbackParam uintptr) {
	_wkeOnLoadUrlEnd.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnLoadUrlFail(webView WkeWebView, callback WkeLoadUrlFailCallback, callbackParam uintptr) {
	_wkeOnLoadUrlFail.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnDidCreateScriptContext(webView WkeWebView, callback WkeDidCreateScriptContextCallback, callbackParam uintptr) {
	_wkeOnDidCreateScriptContext.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnWillReleaseScriptContext(webView WkeWebView, callback WkeWillReleaseScriptContextCallback, callbackParam uintptr) {
	_wkeOnWillReleaseScriptContext.Call(uintptr(webView), uintptr(callback), callbackParam)
}

func wkeOnWindowClosing(webWindow WkeWebView, callback WkeWindowClosingCallback, param uintptr) {
	_wkeOnWindowClosing.Call(uintptr(webWindow), uintptr(callback), param)
}

func wkeOnWindowDestroy(webWindow WkeWebView, callback WkeWindowDestroyCallback, param uintptr) {
	_wkeOnWindowDestroy.Call(uintptr(webWindow), uintptr(callback), param)
}

func wkeOnDraggableRegionsChanged(webView WkeWebView, callback WkeDraggableRegionsChangedCallback, param uintptr) {
	_wkeOnDraggableRegionsChanged.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnWillMediaLoad(webView WkeWebView, callback WkeWillMediaLoadCallback, param uintptr) {
	_wkeOnWillMediaLoad.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnStartDragging(webView WkeWebView, callback WkeStartDraggingCallback, param uintptr) {
	_wkeOnStartDragging.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnPrint(webView WkeWebView, callback WkeOnPrintCallback, param uintptr) {
	_wkeOnPrint.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnOtherLoad(webView WkeWebView, callback WkeOnOtherLoadCallback, param uintptr) {
	_wkeOnOtherLoad.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnContextMenuItemClick(webView WkeWebView, callback WkeOnContextMenuItemClickCallback, param uintptr) {
	_wkeOnContextMenuItemClick.Call(uintptr(webView), uintptr(callback), param)
}

func wkeIsProcessingUserGesture(webView WkeWebView) bool {
	r, _, _ := _wkeIsProcessingUserGesture.Call(uintptr(webView))
	return GoBool(r)
}

func wkeNetSetMIMEType(jobPtr WkeNetJob, aType string) {
	_wkeNetSetMIMEType.Call(uintptr(jobPtr), CAStr(aType))
}

func wkeNetGetMIMEType(jobPtr WkeNetJob, mime WkeString) string {
	r, _, _ := _wkeNetGetMIMEType.Call(uintptr(jobPtr), uintptr(mime))
	return GoAStr(r)
}

func wkeNetSetHTTPHeaderField(jobPtr WkeNetJob, key string, value string, response bool) {
	_wkeNetSetHTTPHeaderField.Call(uintptr(jobPtr), CAStr(key), CAStr(value), CBool(response))
}

func wkeNetGetHTTPHeaderField(jobPtr WkeNetJob, key string) string {
	r, _, _ := _wkeNetGetHTTPHeaderField.Call(uintptr(jobPtr), CAStr(key))
	return GoAStr(r)
}

func wkeNetGetHTTPHeaderFieldFromResponse(jobPtr WkeNetJob, key string) string {
	r, _, _ := _wkeNetGetHTTPHeaderFieldFromResponse.Call(uintptr(jobPtr), CAStr(key))
	return GoAStr(r)
}

func wkeNetSetData(jobPtr WkeNetJob, buf uintptr, len int) {
	_wkeNetSetData.Call(uintptr(jobPtr), buf, uintptr(len))
}

func wkeNetHookRequest(jobPtr WkeNetJob) {
	_wkeNetHookRequest.Call(uintptr(jobPtr))
}

func wkeNetOnResponse(webView WkeWebView, callback WkeNetResponseCallback, param uintptr) {
	_wkeNetOnResponse.Call(uintptr(webView), uintptr(callback), param)
}

func wkeNetGetRequestMethod(jobPtr WkeNetJob) WkeRequestType {
	r, _, _ := _wkeNetGetRequestMethod.Call(uintptr(jobPtr))
	return WkeRequestType(r)
}

func wkeNetGetFavicon(webView WkeWebView, callback WkeOnNetGetFaviconCallback, param uintptr) int {
	r, _, _ := _wkeNetGetFavicon.Call(uintptr(webView), uintptr(callback), param)
	return int(r)
}

func wkeNetContinueJob(jobPtr WkeNetJob) {
	_wkeNetContinueJob.Call(uintptr(jobPtr))
}

func wkeNetGetUrlByJob(jobPtr WkeNetJob) string {
	r, _, _ := _wkeNetGetUrlByJob.Call(uintptr(jobPtr))
	return GoAStr(r)
}

func wkeNetCancelRequest(jobPtr WkeNetJob) {
	_wkeNetCancelRequest.Call(uintptr(jobPtr))
}

func wkeNetHoldJobToAsynCommit(jobPtr WkeNetJob) bool {
	r, _, _ := _wkeNetHoldJobToAsynCommit.Call(uintptr(jobPtr))
	return GoBool(r)
}

func wkeNetChangeRequestUrl(jobPtr WkeNetJob, url string) {
	_wkeNetChangeRequestUrl.Call(uintptr(jobPtr), CAStr(url))
}

func wkeNetCreateWebUrlRequest(url string, method string, mime string) WkeWebUrlRequestPtr {
	r, _, _ := _wkeNetCreateWebUrlRequest.Call(CAStr(url), CAStr(method), CAStr(mime))
	return WkeWebUrlRequestPtr(r)
}

func wkeNetCreateWebUrlRequest2(request BlinkWebURLRequestPtr) WkeWebUrlRequestPtr {
	r, _, _ := _wkeNetCreateWebUrlRequest2.Call(uintptr(request))
	return WkeWebUrlRequestPtr(r)
}

func wkeNetCopyWebUrlRequest(jobPtr WkeNetJob, needExtraData bool) BlinkWebURLRequestPtr {
	r, _, _ := _wkeNetCopyWebUrlRequest.Call(uintptr(jobPtr), CBool(needExtraData))
	return BlinkWebURLRequestPtr(r)
}

func wkeNetDeleteBlinkWebURLRequestPtr(request BlinkWebURLRequestPtr) {
	_wkeNetDeleteBlinkWebURLRequestPtr.Call(uintptr(request))
}

func wkeNetAddHTTPHeaderFieldToUrlRequest(request WkeWebUrlRequestPtr, name string, value string) {
	_wkeNetAddHTTPHeaderFieldToUrlRequest.Call(uintptr(request), CAStr(name), CAStr(value))
}

func wkeNetStartUrlRequest(webView WkeWebView, request WkeWebUrlRequestPtr, param uintptr, callbacks *WkeUrlRequestCallbacks) int {
	r, _, _ := _wkeNetStartUrlRequest.Call(uintptr(webView), uintptr(request), param, uintptr(unsafe.Pointer(callbacks)))
	return int(r)
}

func wkeNetGetHttpStatusCode(response WkeWebUrlResponsePtr) int {
	r, _, _ := _wkeNetGetHttpStatusCode.Call(uintptr(response))
	return int(r)
}

func wkeNetGetExpectedContentLength(response WkeWebUrlResponsePtr) int64 {
	r, _, _ := _wkeNetGetExpectedContentLength.Call(uintptr(response))
	return int64(r)
}

func wkeNetGetResponseUrl(response WkeWebUrlResponsePtr) string {
	r, _, _ := _wkeNetGetResponseUrl.Call(uintptr(response))
	return GoAStr(r)
}

func wkeNetCancelWebUrlRequest(requestId int) {
	_wkeNetCancelWebUrlRequest.Call(uintptr(requestId))
}

func wkeNetGetPostBody(jobPtr WkeNetJob) *WkePostBodyElements {
	r, _, _ := _wkeNetGetPostBody.Call(uintptr(jobPtr))
	return (*WkePostBodyElements)(unsafe.Pointer(r))
}

func wkeNetCreatePostBodyElements(webView WkeWebView, length uint) *WkePostBodyElements {
	r, _, _ := _wkeNetCreatePostBodyElements.Call(uintptr(webView), uintptr(length))
	return (*WkePostBodyElements)(unsafe.Pointer(r))
}

func wkeNetFreePostBodyElements(elements *WkePostBodyElements) {
	_wkeNetFreePostBodyElements.Call(uintptr(unsafe.Pointer(elements)))
}

func wkeNetCreatePostBodyElement(webView WkeWebView) *WkePostBodyElement {
	r, _, _ := _wkeNetCreatePostBodyElement.Call(uintptr(webView))
	return (*WkePostBodyElement)(unsafe.Pointer(r))
}

func wkeNetFreePostBodyElement(element *WkePostBodyElement) {
	_wkeNetFreePostBodyElement.Call(uintptr(unsafe.Pointer(element)))
}

func wkeIsMainFrame(webView WkeWebView, frameId WkeWebFrameHandle) bool {
	r, _, _ := _wkeIsMainFrame.Call(uintptr(webView), uintptr(frameId))
	return GoBool(r)
}

func wkeIsWebRemoteFrame(webView WkeWebView, frameId WkeWebFrameHandle) bool {
	r, _, _ := _wkeIsWebRemoteFrame.Call(uintptr(webView), uintptr(frameId))
	return GoBool(r)
}

func wkeWebFrameGetMainFrame(webView WkeWebView) WkeWebFrameHandle {
	r, _, _ := _wkeWebFrameGetMainFrame.Call(uintptr(webView))
	return WkeWebFrameHandle(r)
}

func wkeRunJsByFrame(webView WkeWebView, frameId WkeWebFrameHandle, script string, isInClosure bool) JsValue {
	r, r2, _ := _wkeRunJsByFrame.Call(uintptr(webView), uintptr(frameId), CAStr(script), CBool(isInClosure))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func wkeInsertCSSByFrame(webView WkeWebView, frameId WkeWebFrameHandle, cssText string) {
	_wkeInsertCSSByFrame.Call(uintptr(webView), uintptr(frameId), CAStr(cssText))
}

func wkeWebFrameGetMainWorldScriptContext(webView WkeWebView, webFrameId WkeWebFrameHandle, contextOut V8ContextPtr) {
	_wkeWebFrameGetMainWorldScriptContext.Call(uintptr(webView), uintptr(webFrameId), uintptr(contextOut))
}

func wkeGetBlinkMainThreadIsolate() V8Isolate {
	r, _, _ := _wkeGetBlinkMainThreadIsolate.Call()
	return V8Isolate(r)
}

func wkeCreateWebWindow(aType WkeWindowType, parent types.HWND, x int, y int, width int, height int) WkeWebView {
	r, _, _ := _wkeCreateWebWindow.Call(uintptr(aType), uintptr(parent), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return WkeWebView(r)
}

func wkeCreateWebCustomWindow(info *WkeWindowCreateInfo) WkeWebView {
	r, _, _ := _wkeCreateWebCustomWindow.Call(uintptr(unsafe.Pointer(info)))
	return WkeWebView(r)
}

func wkeDestroyWebWindow(webWindow WkeWebView) {
	_wkeDestroyWebWindow.Call(uintptr(webWindow))
}

func wkeGetWindowHandle(webWindow WkeWebView) types.HWND {
	r, _, _ := _wkeGetWindowHandle.Call(uintptr(webWindow))
	return types.HWND(r)
}

func wkeShowWindow(webWindow WkeWebView, show bool) {
	_wkeShowWindow.Call(uintptr(webWindow), CBool(show))
}

func wkeEnableWindow(webWindow WkeWebView, enable bool) {
	_wkeEnableWindow.Call(uintptr(webWindow), CBool(enable))
}

func wkeMoveWindow(webWindow WkeWebView, x int, y int, width int, height int) {
	_wkeMoveWindow.Call(uintptr(webWindow), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func wkeMoveToCenter(webWindow WkeWebView) {
	_wkeMoveToCenter.Call(uintptr(webWindow))
}

func wkeResizeWindow(webWindow WkeWebView, width int, height int) {
	_wkeResizeWindow.Call(uintptr(webWindow), uintptr(width), uintptr(height))
}

func wkeDragTargetDragEnter(webView WkeWebView, webDragData *WkeWebDragData, clientPoint *types.TPoint, screenPoint *types.TPoint, operationsAllowed WkeWebDragOperationsMask, modifiers int) WkeWebDragOperation {
	r, _, _ := _wkeDragTargetDragEnter.Call(uintptr(webView), uintptr(unsafe.Pointer(webDragData)), uintptr(unsafe.Pointer(clientPoint)), uintptr(unsafe.Pointer(screenPoint)), uintptr(operationsAllowed), uintptr(modifiers))
	return WkeWebDragOperation(r)
}

func wkeDragTargetDragOver(webView WkeWebView, clientPoint *types.TPoint, screenPoint *types.TPoint, operationsAllowed WkeWebDragOperationsMask, modifiers int) WkeWebDragOperation {
	r, _, _ := _wkeDragTargetDragOver.Call(uintptr(webView), uintptr(unsafe.Pointer(clientPoint)), uintptr(unsafe.Pointer(screenPoint)), uintptr(operationsAllowed), uintptr(modifiers))
	return WkeWebDragOperation(r)
}

func wkeDragTargetDragLeave(webView WkeWebView) {
	_wkeDragTargetDragLeave.Call(uintptr(webView))
}

func wkeDragTargetDrop(webView WkeWebView, clientPoint *types.TPoint, screenPoint *types.TPoint, modifiers int) {
	_wkeDragTargetDrop.Call(uintptr(webView), uintptr(unsafe.Pointer(clientPoint)), uintptr(unsafe.Pointer(screenPoint)), uintptr(modifiers))
}

func wkeDragTargetEnd(webView WkeWebView, clientPoint *types.TPoint, screenPoint *types.TPoint, operation WkeWebDragOperation) {
	_wkeDragTargetEnd.Call(uintptr(webView), uintptr(unsafe.Pointer(clientPoint)), uintptr(unsafe.Pointer(screenPoint)), uintptr(operation))
}

func wkeUtilSetUiCallback(callback WkeUiThreadPostTaskCallback) {
	_wkeUtilSetUiCallback.Call(uintptr(callback))
}

func wkeUtilSerializeToMHTML(webView WkeWebView) string {
	r, _, _ := _wkeUtilSerializeToMHTML.Call(uintptr(webView))
	return GoAStr(r)
}

func wkeUtilPrintToPdf(webView WkeWebView, frameId WkeWebFrameHandle, settings *WkePrintSettings) *WkePdfDatas {
	r, _, _ := _wkeUtilPrintToPdf.Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(settings)))
	return (*WkePdfDatas)(unsafe.Pointer(r))
}

func wkePrintToBitmap(webView WkeWebView, frameId WkeWebFrameHandle, settings *WkeScreenshotSettings) *WkeMemBuf {
	r, _, _ := _wkePrintToBitmap.Call(uintptr(webView), uintptr(frameId), uintptr(unsafe.Pointer(settings)))
	return (*WkeMemBuf)(unsafe.Pointer(r))
}

func wkeUtilRelasePrintPdfDatas(datas *WkePdfDatas) {
	_wkeUtilRelasePrintPdfDatas.Call(uintptr(unsafe.Pointer(datas)))
}

func wkeSetWindowTitle(webWindow WkeWebView, title string) {
	_wkeSetWindowTitle.Call(uintptr(webWindow), CAStr(title))
}

func wkeSetWindowTitleW(webWindow WkeWebView, title string) {
	_wkeSetWindowTitleW.Call(uintptr(webWindow), CWStr(title))
}

func wkeNodeOnCreateProcess(webView WkeWebView, callback WkeNodeOnCreateProcessCallback, param uintptr) {
	_wkeNodeOnCreateProcess.Call(uintptr(webView), uintptr(callback), param)
}

func wkeOnPluginFind(webView WkeWebView, mime string, callback WkeOnPluginFindCallback, param uintptr) {
	_wkeOnPluginFind.Call(uintptr(webView), CAStr(mime), uintptr(callback), param)
}

func wkeAddNpapiPlugin(webView WkeWebView, initializeFunc uintptr, getEntryPointsFunc uintptr, shutdownFunc uintptr) {
	_wkeAddNpapiPlugin.Call(uintptr(webView), initializeFunc, getEntryPointsFunc, shutdownFunc)
}

func wkeGetWebViewByNData(ndata uintptr) WkeWebView {
	r, _, _ := _wkeGetWebViewByNData.Call(ndata)
	return WkeWebView(r)
}

func wkeRegisterEmbedderCustomElement(webView WkeWebView, frameId WkeWebFrameHandle, name string, options uintptr, outResult uintptr) bool {
	r, _, _ := _wkeRegisterEmbedderCustomElement.Call(uintptr(webView), uintptr(frameId), CAStr(name), options, outResult)
	return GoBool(r)
}

func wkeSetMediaPlayerFactory(webView WkeWebView, factory WkeMediaPlayerFactory, callback WkeOnIsMediaPlayerSupportsMIMEType) {
	_wkeSetMediaPlayerFactory.Call(uintptr(webView), uintptr(factory), uintptr(callback))
}

func wkeGetContentAsMarkup(webView WkeWebView, frame WkeWebFrameHandle, size *uint) string {
	r, _, _ := _wkeGetContentAsMarkup.Call(uintptr(webView), uintptr(frame), uintptr(unsafe.Pointer(size)))
	return GoAStr(r)
}

func wkeUtilDecodeURLEscape(url string) string {
	r, _, _ := _wkeUtilDecodeURLEscape.Call(CAStr(url))
	return GoAStr(r)
}

func wkeUtilEncodeURLEscape(url string) string {
	r, _, _ := _wkeUtilEncodeURLEscape.Call(CAStr(url))
	return GoAStr(r)
}

func wkeUtilBase64Encode(str string) string {
	r, _, _ := _wkeUtilBase64Encode.Call(CAStr(str))
	return GoAStr(r)
}

func wkeUtilBase64Decode(str string) string {
	r, _, _ := _wkeUtilBase64Decode.Call(CAStr(str))
	return GoAStr(r)
}

func wkeUtilCreateV8Snapshot(str string) *WkeMemBuf {
	r, _, _ := _wkeUtilCreateV8Snapshot.Call(CAStr(str))
	return (*WkeMemBuf)(unsafe.Pointer(r))
}

func wkeRunMessageLoop() {
	_wkeRunMessageLoop.Call()
}

func wkeSaveMemoryCache(webView WkeWebView) {
	_wkeSaveMemoryCache.Call(uintptr(webView))
}
