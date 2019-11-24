// +build windows

// miniblink及wke头文件导入
// 由ying32翻译，应用于govcl

package miniblink

import "github.com/ying32/govcl/vcl/types"

// WKE_CALL_TYPE __cdecl
type WkeRect struct {
	X, Y, W, H int
}

type WkePoint struct {
	X, Y int
}

//type Point = WkePoint

type WkeMouseFlags int

const (
	WKE_LBUTTON = 0x01
	WKE_RBUTTON = 0x02
	WKE_SHIFT   = 0x04
	WKE_CONTROL = 0x08
	WKE_MBUTTON = 0x10
)

type WkeKeyFlags int

const (
	WKE_EXTENDED = 0x0100
	WKE_REPEAT   = 0x4000
)

type WkeMouseMsg int

const (
	WKE_MSG_MOUSEMOVE     = 0x0200
	WKE_MSG_LBUTTONDOWN   = 0x0201
	WKE_MSG_LBUTTONUP     = 0x0202
	WKE_MSG_LBUTTONDBLCLK = 0x0203
	WKE_MSG_RBUTTONDOWN   = 0x0204
	WKE_MSG_RBUTTONUP     = 0x0205
	WKE_MSG_RBUTTONDBLCLK = 0x0206
	WKE_MSG_MBUTTONDOWN   = 0x0207
	WKE_MSG_MBUTTONUP     = 0x0208
	WKE_MSG_MBUTTONDBLCLK = 0x0209
	WKE_MSG_MOUSEWHEEL    = 0x020A
)

type JsExecState uintptr

type JsValue int64

type WkeWebView uintptr

type WkeString uintptr

type WkeMediaPlayer uintptr
type WkeMediaPlayerClient uintptr
type BlinkWebURLRequestPtr uintptr

type WkeProxyType int

const (
	WKE_PROXY_NONE WkeProxyType = iota + 0
	WKE_PROXY_HTTP
	WKE_PROXY_SOCKS4
	WKE_PROXY_SOCKS4A
	WKE_PROXY_SOCKS5
	WKE_PROXY_SOCKS5HOSTNAME
)

type WkeProxy struct {
	Type     WkeProxyType
	Hostname [100]uint8
	Port     uint16
	Username [50]uint8
	Password [50]uint8
}

type WkeSettingMask int

const (
	WKE_SETTING_PROXY                         = 1
	WKE_SETTING_PAINTCALLBACK_IN_OTHER_THREAD = 1 << 2
)

type WkeSettings struct {
	Proxy WkeProxy
	Mask  uint
}

type WkeViewSettings struct {
	Size    int
	BgColor uint
}

type WkeWebFrameHandle uintptr

type WkeMenuItemId int

const (
	KWkeMenuSelectedAllId      = 1 << 1
	KWkeMenuSelectedTextId     = 1 << 2
	KWkeMenuUndoId             = 1 << 3
	KWkeMenuCopyImageId        = 1 << 4
	kWkeMenuInspectElementAtId = 1 << 5
	KWkeMenuCutId              = 1 << 6
	KWkeMenuPasteId            = 1 << 7
	KWkeMenuPrintId            = 1 << 8
	KWkeMenuGoForwardId        = 1 << 9
	KWkeMenuGoBackId           = 1 << 10
	KWkeMenuReloadId           = 1 << 11
)

//typedef void* (WKE_CALL_TYPE *FILE_OPEN_) (const char* path);
type FILE_OPEN_ uintptr

//typedef void(WKE_CALL_TYPE *FILE_CLOSE_) (void* handle);
type FILE_CLOSE_ uintptr

//typedef size_t(WKE_CALL_TYPE *FILE_SIZE) (void* handle);
type FILE_SIZE uintptr

//typedef int(WKE_CALL_TYPE *FILE_READ) (void* handle, void* buffer, size_t size);
type FILE_READ uintptr

//typedef int(WKE_CALL_TYPE *FILE_SEEK) (void* handle, int offset, int origin);
type FILE_SEEK uintptr

type WKE_FILE_OPEN FILE_OPEN_
type WKE_FILE_CLOSE FILE_CLOSE_
type WKE_FILE_SIZE FILE_SIZE
type WKE_FILE_READ FILE_READ
type WKE_FILE_SEEK FILE_SEEK
type WKE_EXISTS_FILE uintptr // bool(WKE_CALL_TYPE *WKE_EXISTS_FILE)(const char * path);

type _WkeClientHandler struct { // declare warning fix

}

//typedef void(WKE_CALL_TYPE *ON_TITLE_CHANGED) (const struct _wkeClientHandler* clientHandler, const wkeString title);
type ON_TITLE_CHANGED uintptr

//typedef void(WKE_CALL_TYPE *ON_URL_CHANGED) (const struct _wkeClientHandler* clientHandler, const wkeString url);
type ON_URL_CHANGED uintptr

type WkeClientHandler struct {
	OnTitleChanged ON_TITLE_CHANGED
	OnURLChanged   ON_URL_CHANGED
}

/*
typedef bool(WKE_CALL_TYPE * wkeCookieVisitor)(
    void* params,
    const char* name,
    const char* value,
    const char* domain,
    const char* path, // If |path| is non-empty only URLs at or below the path will get the cookie value.
    int secure, // If |secure| is true the cookie will only be sent for HTTPS requests.
    int httpOnly, // If |httponly| is true the cookie will only be sent for HTTP requests.
    int* expires // The cookie expiration date is only valid if |has_expires| is true.
    );
*/
type WkeCookieVisitor = uintptr

type WkeCookieCommand int

const (
	WkeCookieCommandClearAllCookies WkeCookieCommand = iota + 0
	WkeCookieCommandClearSessionCookies
	WkeCookieCommandFlushCookiesToFile
	WkeCookieCommandReloadCookiesFromFile
)

type WkeNavigationType int

const (
	WKE_NAVIGATION_TYPE_LINKCLICK WkeNavigationType = iota + 0
	WKE_NAVIGATION_TYPE_FORMSUBMITTE
	WKE_NAVIGATION_TYPE_BACKFORWARD
	WKE_NAVIGATION_TYPE_RELOAD
	WKE_NAVIGATION_TYPE_FORMRESUBMITT
	WKE_NAVIGATION_TYPE_OTHER
)

type WkeCursorInfoType int

const (
	WkeCursorInfoPointer WkeCursorInfoType = iota + 0
	WkeCursorInfoCross
	WkeCursorInfoHand
	WkeCursorInfoIBeam
	WkeCursorInfoWait
	WkeCursorInfoHelp
	WkeCursorInfoEastResize
	WkeCursorInfoNorthResize
	WkeCursorInfoNorthEastResize
	WkeCursorInfoNorthWestResize
	WkeCursorInfoSouthResize
	WkeCursorInfoSouthEastResize
	WkeCursorInfoSouthWestResize
	WkeCursorInfoWestResize
	WkeCursorInfoNorthSouthResize
	WkeCursorInfoEastWestResize
	WkeCursorInfoNorthEastSouthWestResize
	WkeCursorInfoNorthWestSouthEastResize
	WkeCursorInfoColumnResize
	WkeCursorInfoRowResize
	WkeCursorInfoMiddlePanning
	WkeCursorInfoEastPanning
	WkeCursorInfoNorthPanning
	WkeCursorInfoNorthEastPanning
	WkeCursorInfoNorthWestPanning
	WkeCursorInfoSouthPanning
	WkeCursorInfoSouthEastPanning
	WkeCursorInfoSouthWestPanning
	WkeCursorInfoWestPanning
	WkeCursorInfoMove
	WkeCursorInfoVerticalText
	WkeCursorInfoCell
	WkeCursorInfoContextMenu
	WkeCursorInfoAlias
	WkeCursorInfoProgress
	WkeCursorInfoNoDrop
	WkeCursorInfoCopy
	WkeCursorInfoNone
	WkeCursorInfoNotAllowed
	WkeCursorInfoZoomIn
	WkeCursorInfoZoomOut
	WkeCursorInfoGrab
	WkeCursorInfoGrabbing
	WkeCursorInfoCustom
)

type WkeWindowFeatures struct {
	X      int
	Y      int
	Width  int
	Height int

	MenuBarVisible     bool
	StatusBarVisible   bool
	ToolBarVisible     bool
	LocationBarVisible bool
	ScrollbarsVisible  bool
	Resizable          bool
	Fullscreen         bool
}

type WkeMemBuf struct {
	Size   int
	Data   uintptr
	Length uint
}

type WkeStorageType int

const (
	// String data with an associated MIME type. Depending on the MIME type, there may be
	// optional metadata attributes as well.
	StorageTypeString = iota + 0
	// Stores the name of one file being dragged into the renderer.
	StorageTypeFilename
	// An image being dragged out of the renderer. Contains a buffer holding the image data
	// as well as the suggested name for saving the image to.
	StorageTypeBinaryData
	// Stores the filesystem URL of one file being dragged into the renderer.
	StorageTypeFileSystemFile
)

type StorageType WkeStorageType

type Item struct {

	// Only valid when storageType == StorageTypeString.
	StringType *WkeMemBuf
	StringData *WkeMemBuf

	// Only valid when storageType == StorageTypeFilename.
	FilenameData    *WkeMemBuf
	DisplayNameData *WkeMemBuf

	// Only valid when storageType == StorageTypeBinaryData.
	BinaryData *WkeMemBuf

	// Title associated with a link when stringType == "text/uri-list".
	// Filename when storageType == StorageTypeBinaryData.
	Title *WkeMemBuf

	// Only valid when storageType == StorageTypeFileSystemFile.
	FileSystemURL      *WkeMemBuf
	FileSystemFileSize int64

	// Only valid when stringType == "text/html".
	BaseURL *WkeMemBuf
}

type WkeWebDragData struct {
	M_itemList       *Item
	M_itemListLength int

	M_modifierKeyState int // State of Shift/Ctrl/Alt/Meta keys.
	M_filesystemId     *WkeMemBuf
}

type WkeWebDragOperation int

const (
	WkeWebDragOperationNone    = 0
	WkeWebDragOperationCopy    = 1
	WkeWebDragOperationLink    = 2
	WkeWebDragOperationGeneric = 4
	WkeWebDragOperationPrivate = 8
	WkeWebDragOperationMove    = 16
	WkeWebDragOperationDelete  = 32
	WkeWebDragOperationEvery   = 0xffffffff
)

type WkeWebDragOperationsMask WkeWebDragOperation

type WkeResourceType int

const (
	WKE_RESOURCE_TYPE_MAIN_FRAME    = 0 // top level page
	WKE_RESOURCE_TYPE_SUB_FRAME     = 1 // frame or iframe
	WKE_RESOURCE_TYPE_STYLESHEET    = 2 // a CSS stylesheet
	WKE_RESOURCE_TYPE_SCRIPT        = 3 // an external script
	WKE_RESOURCE_TYPE_IMAGE         = 4 // an image (jpg/gif/png/etc)
	WKE_RESOURCE_TYPE_FONT_RESOURCE = 5 // a font
	WKE_RESOURCE_TYPE_SUB_RESOURCE  = 6 // an "other" subresource.
	WKE_RESOURCE_TYPE_OBJECT        = 7 // an object (or embed) tag for a plugin,
	// or a resource that a plugin requested.
	WKE_RESOURCE_TYPE_MEDIA  = 8 // a media resource.
	WKE_RESOURCE_TYPE_WORKER = 9 // the main resource of a dedicated
	// worker.
	WKE_RESOURCE_TYPE_SHARED_WORKER  = 10 // the main resource of a shared worker.
	WKE_RESOURCE_TYPE_PREFETCH       = 11 // an explicitly requested prefetch
	WKE_RESOURCE_TYPE_FAVICON        = 12 // a favicon
	WKE_RESOURCE_TYPE_XHR            = 13 // a XMLHttpRequest
	WKE_RESOURCE_TYPE_PING           = 14 // a ping request for <a ping>
	WKE_RESOURCE_TYPE_SERVICE_WORKER = 15 // the main resource of a service worker.
	WKE_RESOURCE_TYPE_LAST_TYPE
)

type WkeWillSendRequestInfo struct {
	Url              WkeString
	NewUrl           WkeString
	ResourceType     WkeResourceType
	HttpResponseCode int
	Method           WkeString
	Referrer         WkeString
	Headers          uintptr
}

type WkeHttBodyElementType int

const (
	WkeHttBodyElementTypeData WkeHttBodyElementType = iota + 0
	WkeHttBodyElementTypeFile
)

type WkePostBodyElement struct {
	Size       int
	AType      WkeHttBodyElementType
	Data       *WkeMemBuf
	FilePath   WkeString
	FileStart  int64
	FileLength int64 // -1 means to the end of the file.
}

type WkePostBodyElements struct {
	Size        int
	Element     **WkePostBodyElement
	ElementSize uint
	IsDirty     bool
}

type WkeNetJob uintptr

type WkeTempCallbackInfo struct {
	Size                int
	Frame               WkeWebFrameHandle
	WillSendRequestInfo *WkeWillSendRequestInfo
	Url                 uintptr //const char* url;
	PostBody            WkePostBodyElements
	Job                 WkeNetJob
}

type WkeRequestType int

const (
	KWkeRequestTypeInvalidation WkeRequestType = iota + 0
	KWkeRequestTypeGet
	KWkeRequestTypePost
	KWkeRequestTypePut
)

type WkePdfDatas struct {
	Count int
	Sizes *uint
	Datas **uintptr
}

type WkePrintSettings struct {
	StructSize               int
	DPI                      int
	Width                    int // in px
	Height                   int
	MarginTop                int
	MarginBottom             int
	MarginLeft               int
	MarginRight              int
	IsPrintPageHeadAndFooter bool
	IsLandscape              bool
	IsPrintBackgroud         bool
}

type WkeScreenshotSettings struct {
	StructSize int
	Width      int
	Height     int
}

type (
	/*
		typedef void(WKE_CALL_TYPE*wkeTitleChangedCallback)(wkeWebView webView, void* param, const wkeString title);
		typedef void(WKE_CALL_TYPE*wkeURLChangedCallback)(wkeWebView webView, void* param, const wkeString url);
		typedef void(WKE_CALL_TYPE*wkeURLChangedCallback2)(wkeWebView webView, void* param, wkeWebFrameHandle frameId, const wkeString url);
		typedef void(WKE_CALL_TYPE*wkePaintUpdatedCallback)(wkeWebView webView, void* param, const HDC hdc, int x, int y, int cx, int cy);
		typedef void(WKE_CALL_TYPE*wkePaintBitUpdatedCallback)(wkeWebView webView, void* param, const void* buffer, const wkeRect* r, int width, int height);
		typedef void(WKE_CALL_TYPE*wkeAlertBoxCallback)(wkeWebView webView, void* param, const wkeString msg);
		typedef bool(WKE_CALL_TYPE*wkeConfirmBoxCallback)(wkeWebView webView, void* param, const wkeString msg);
		typedef bool(WKE_CALL_TYPE*wkePromptBoxCallback)(wkeWebView webView, void* param, const wkeString msg, const wkeString defaultResult, wkeString result);
		typedef bool(WKE_CALL_TYPE*wkeNavigationCallback)(wkeWebView webView, void* param, wkeNavigationType navigationType, wkeString url);
		typedef wkeWebView(WKE_CALL_TYPE*wkeCreateViewCallback)(wkeWebView webView, void* param, wkeNavigationType navigationType, const wkeString url, const wkeWindowFeatures* windowFeatures);
		typedef void(WKE_CALL_TYPE*wkeDocumentReadyCallback)(wkeWebView webView, void* param);
		typedef void(WKE_CALL_TYPE*wkeDocumentReady2Callback)(wkeWebView webView, void* param, wkeWebFrameHandle frameId);

		typedef void(WKE_CALL_TYPE*wkeOnShowDevtoolsCallback)(wkeWebView webView, void* param);

		typedef void(WKE_CALL_TYPE*wkeNodeOnCreateProcessCallback)(wkeWebView webView, void* param, const WCHAR* applicationPath, const WCHAR* arguments, STARTUPINFOW* startup);
		typedef void(WKE_CALL_TYPE*wkeOnPluginFindCallback)(wkeWebView webView, void* param, const utf8* mime, void* initializeFunc, void* getEntryPointsFunc, void* shutdownFunc);

		typedef void(WKE_CALL_TYPE*wkeOnPrintCallback)(wkeWebView webView, void* param, wkeWebFrameHandle frameId, void* printParams);
	*/
	WkeTitleChangedCallback        uintptr
	WkeURLChangedCallback          uintptr
	WkeURLChangedCallback2         uintptr
	WkePaintUpdatedCallback        uintptr
	WkePaintBitUpdatedCallback     uintptr
	WkeAlertBoxCallback            uintptr
	WkeConfirmBoxCallback          uintptr
	WkePromptBoxCallback           uintptr
	WkeNavigationCallback          uintptr
	WkeCreateViewCallback          uintptr
	WkeDocumentReadyCallback       uintptr
	WkeDocumentReady2Callback      uintptr
	WkeOnShowDevtoolsCallback      uintptr
	WkeNodeOnCreateProcessCallback uintptr
	WkeOnPluginFindCallback        uintptr
	WkeOnPrintCallback             uintptr
)

type WkeMediaLoadInfo struct {
	Size     int
	Width    int
	Height   int
	Duration float64
}

// typedef void(WKE_CALL_TYPE*wkeWillMediaLoadCallback)(wkeWebView webView, void* param, const char* url, wkeMediaLoadInfo* info);
type WkeWillMediaLoadCallback uintptr

/*
typedef void(WKE_CALL_TYPE*wkeStartDraggingCallback)(
    wkeWebView webView,
    void* param,
    wkeWebFrameHandle frame,
    const wkeWebDragData* data,
    wkeWebDragOperationsMask mask,
    const void* image,
    const wkePoint* dragImageOffset
    );
*/
type WkeStartDraggingCallback uintptr

// typedef void(WKE_CALL_TYPE*wkeUiThreadRunCallback)(HWND hWnd, void* param);
type WkeUiThreadRunCallback uintptr

// typedef int(WKE_CALL_TYPE*wkeUiThreadPostTaskCallback)(HWND hWnd, wkeUiThreadRunCallback callback, void* param);
type WkeUiThreadPostTaskCallback uintptr

type WkeOtherLoadType int

const (
	WKE_DID_START_LOADING WkeOtherLoadType = iota + 0
	WKE_DID_STOP_LOADING
	WKE_DID_NAVIGATE
	WKE_DID_NAVIGATE_IN_PAGE
	WKE_DID_GET_RESPONSE_DETAILS
	WKE_DID_GET_REDIRECT_REQUEST
	WKE_DID_POST_REQUEST
)

// typedef void(WKE_CALL_TYPE*wkeOnOtherLoadCallback)(wkeWebView webView, void* param, wkeOtherLoadType type, wkeTempCallbackInfo* info);
type WkeOnOtherLoadCallback uintptr

type WkeOnContextMenuItemClickType int

const (
	KWkeContextMenuItemClickTypePrint = 0x01
)

type WkeOnContextMenuItemClickStep int

const (
	KWkeContextMenuItemClickStepShow  = 0x01
	KWkeContextMenuItemClickStepClick = 0x02
)

/*
typedef bool(WKE_CALL_TYPE* wkeOnContextMenuItemClickCallback)(
    wkeWebView webView,
    void* param,
    wkeOnContextMenuItemClickType type,
    wkeOnContextMenuItemClickStep step,
    wkeWebFrameHandle frameId,
    void* info
    );

*/
type WkeOnContextMenuItemClickCallback uintptr

type WkeLoadingResult int

const (
	WKE_LOADING_SUCCEEDED WkeLoadingResult = iota + 0
	WKE_LOADING_FAILED
	WKE_LOADING_CANCELED
)

type WkeDownloadOpt int

const (
	KWkeDownloadOptCancel = iota + 0
	KWkeDownloadOptCacheData
)

// typedef void(WKE_CALL_TYPE*wkeNetJobDataRecvCallback)(void* ptr, wkeNetJob job, const char* data, int length);
type WkeNetJobDataRecvCallback uintptr

// typedef void(WKE_CALL_TYPE*wkeNetJobDataFinishCallback)(void* ptr, wkeNetJob job, wkeLoadingResult result);
type WkeNetJobDataFinishCallback uintptr

type WkeNetJobDataBind struct {
	Param          uintptr
	RecvCallback   WkeNetJobDataRecvCallback
	FinishCallback WkeNetJobDataFinishCallback
}

// typedef void(WKE_CALL_TYPE*wkeLoadingFinishCallback)(wkeWebView webView, void* param, const wkeString url, wkeLoadingResult result, const wkeString failedReason);
type WkeLoadingFinishCallback uintptr

// typedef bool(WKE_CALL_TYPE*wkeDownloadCallback)(wkeWebView webView, void* param, const char* url);
type WkeDownloadCallback uintptr

/*
typedef wkeDownloadOpt(WKE_CALL_TYPE*wkeDownload2Callback)(
    wkeWebView webView,
    void* param,
    size_t expectedContentLength,
    const char* url,
    const char* mime,
    const char* disposition,
    wkeNetJob job,
    wkeNetJobDataBind* dataBind);
*/
type WkeDownload2Callback uintptr

type WkeConsoleLevel int

const (
	WkeLevelDebug        = 4
	WkeLevelLog          = 1
	WkeLevelInfo         = 5
	WkeLevelWarning      = 2
	WkeLevelError        = 3
	WkeLevelRevokedError = 6
	WkeLevelLast         = WkeLevelInfo
)

// typedef void(WKE_CALL_TYPE*wkeConsoleCallback)(wkeWebView webView, void* param, wkeConsoleLevel level, const wkeString message, const wkeString sourceName, unsigned sourceLine, const wkeString stackTrace);
type WkeConsoleCallback uintptr

// typedef void(WKE_CALL_TYPE*wkeOnCallUiThread)(wkeWebView webView, void* paramOnInThread);
type WkeOnCallUiThread uintptr

// typedef void(WKE_CALL_TYPE*wkeCallUiThread)(wkeWebView webView, wkeOnCallUiThread func, void* param);
type WkeCallUiThread uintptr

// typedef wkeMediaPlayer(WKE_CALL_TYPE* wkeMediaPlayerFactory)(wkeWebView webView, wkeMediaPlayerClient client, void* npBrowserFuncs, void* npPluginFuncs);
type WkeMediaPlayerFactory uintptr

// typedef bool(WKE_CALL_TYPE* wkeOnIsMediaPlayerSupportsMIMEType)(const utf8* mime);
type WkeOnIsMediaPlayerSupportsMIMEType uintptr

//
//type WkeWebUrlRequest struct {
//}
//
//type WkeWebUrlResponse struct {
//}

type WkeWebUrlRequestPtr uintptr
type WkeWebUrlResponsePtr uintptr

//typedef void(WKE_CALL_TYPE* wkeOnUrlRequestWillRedirectCallback)(wkeWebView webView, void* param, wkeWebUrlRequestPtr oldRequest, wkeWebUrlRequestPtr request, wkeWebUrlResponsePtr redirectResponse);
type WkeOnUrlRequestWillRedirectCallback uintptr

//typedef void(WKE_CALL_TYPE* wkeOnUrlRequestDidReceiveResponseCallback)(wkeWebView webView, void* param, wkeWebUrlRequestPtr request, wkeWebUrlResponsePtr response);
type WkeOnUrlRequestDidReceiveResponseCallback uintptr

//typedef void(WKE_CALL_TYPE* wkeOnUrlRequestDidReceiveDataCallback)(wkeWebView webView, void* param, wkeWebUrlRequestPtr request, const char* data, int dataLength);
type WkeOnUrlRequestDidReceiveDataCallback uintptr

//typedef void(WKE_CALL_TYPE* wkeOnUrlRequestDidFailCallback)(wkeWebView webView, void* param, wkeWebUrlRequestPtr request, const utf8* error);
type WkeOnUrlRequestDidFailCallback uintptr

//typedef void(WKE_CALL_TYPE* wkeOnUrlRequestDidFinishLoadingCallback)(wkeWebView webView, void* param, wkeWebUrlRequestPtr request, double finishTime);
type WkeOnUrlRequestDidFinishLoadingCallback uintptr

type WkeUrlRequestCallbacks struct {
	WillRedirectCallback       WkeOnUrlRequestWillRedirectCallback
	DidReceiveResponseCallback WkeOnUrlRequestDidReceiveResponseCallback
	DidReceiveDataCallback     WkeOnUrlRequestDidReceiveDataCallback
	DidFailCallback            WkeOnUrlRequestDidFailCallback
	DidFinishLoadingCallback   WkeOnUrlRequestDidFinishLoadingCallback
}

// typedef bool(WKE_CALL_TYPE*wkeLoadUrlBeginCallback)(wkeWebView webView, void* param, const utf8* url, wkeNetJob job);
type WkeLoadUrlBeginCallback uintptr

// typedef void(WKE_CALL_TYPE*wkeLoadUrlEndCallback)(wkeWebView webView, void* param, const utf8* url, wkeNetJob job, void* buf, int len);
type WkeLoadUrlEndCallback uintptr

// typedef void(WKE_CALL_TYPE* wkeLoadUrlFailCallback)(wkeWebView webView, void* param, const utf8* url, wkeNetJob job);
type WkeLoadUrlFailCallback uintptr

// typedef void(WKE_CALL_TYPE*wkeDidCreateScriptContextCallback)(wkeWebView webView, void* param, wkeWebFrameHandle frameId, void* context, int extensionGroup, int worldId);
type WkeDidCreateScriptContextCallback uintptr

// typedef void(WKE_CALL_TYPE*wkeWillReleaseScriptContextCallback)(wkeWebView webView, void* param, wkeWebFrameHandle frameId, void* context, int worldId);
type WkeWillReleaseScriptContextCallback uintptr

// typedef bool(WKE_CALL_TYPE*wkeNetResponseCallback)(wkeWebView webView, void* param, const utf8* url, wkeNetJob job);
type WkeNetResponseCallback uintptr

// typedef void(WKE_CALL_TYPE*wkeOnNetGetFaviconCallback)(wkeWebView webView, void* param, const utf8* url, wkeMemBuf* buf);
type WkeOnNetGetFaviconCallback uintptr

type V8ContextPtr uintptr
type V8Isolate uintptr

type WkeWindowType int

const (
	WKE_WINDOW_TYPE_POPUP = iota + 0
	WKE_WINDOW_TYPE_TRANSPARENT
	WKE_WINDOW_TYPE_CONTROL
)

type WkeWindowCreateInfo struct {
	Size    int
	Parent  types.HWND
	Style   uint32
	StyleEx uint32
	X       int
	Y       int
	Width   int
	Height  int
	Color   types.COLORREF
}

// typedef bool(WKE_CALL_TYPE*wkeWindowClosingCallback)(wkeWebView webWindow, void* param);
type WkeWindowClosingCallback uintptr

// typedef void(WKE_CALL_TYPE*wkeWindowDestroyCallback)(wkeWebView webWindow, void* param);
type WkeWindowDestroyCallback uintptr

type WkeDraggableRegion struct {
	Bounds    types.TRect
	Draggable bool
}

// typedef void(WKE_CALL_TYPE*wkeDraggableRegionsChangedCallback)(wkeWebView webView, void* param, const wkeDraggableRegion* rects, int rectCount);
type WkeDraggableRegionsChangedCallback uintptr

// #define JS_CALL __fastcall
// typedef jsValue(JS_CALL* jsNativeFunction) (jsExecState es);
type JsNativeFunction uintptr

// typedef jsValue(WKE_CALL_TYPE* wkeJsNativeFunction) (jsExecState es, void* param);
type WkeJsNativeFunction uintptr

type JsType int

const (
	JSTYPE_NUMBER JsType = iota + 0
	JSTYPE_STRING
	JSTYPE_BOOLEAN
	JSTYPE_OBJECT
	JSTYPE_FUNCTION
	JSTYPE_UNDEFINED
	JSTYPE_ARRAY
	JSTYPE_NULL
)

// typedef jsValue(WKE_CALL_TYPE*jsGetPropertyCallback)(jsExecState es, jsValue object, const char* propertyName);
type JsGetPropertyCallback uintptr

// typedef bool(WKE_CALL_TYPE*jsSetPropertyCallback)(jsExecState es, jsValue object, const char* propertyName, jsValue value);
type JsSetPropertyCallback uintptr

// typedef jsValue(WKE_CALL_TYPE*jsCallAsFunctionCallback)(jsExecState es, jsValue object, jsValue* args, int argCount);
type JsCallAsFunctionCallback uintptr

// typedef void(WKE_CALL_TYPE*jsFinalizeCallback)(struct tagjsData* data);
type JsFinalizeCallback uintptr

type JsData struct {
	TypeName       [100]uint8
	PropertyGet    JsGetPropertyCallback
	PropertySet    JsSetPropertyCallback
	Finalize       JsFinalizeCallback
	CallAsFunction JsCallAsFunctionCallback
}

type JsExceptionInfo struct {
	Message            uintptr //const utf8* message; // Returns the exception message.
	SourceLine         uintptr //const utf8* sourceLine; // Returns the line of source code that the exception occurred within.
	ScriptResourceName uintptr //const utf8* scriptResourceName; // Returns the resource name for the script from where the function causing the error originates.
	LineNumber         int     // Returns the 1-based number of the line where the error occurred or 0 if the line number is unknown.
	StartPosition      int     // Returns the index within the script of the first character where the error occurred.
	EndPosition        int     // Returns the index within the script of the last character where the error occurred.
	StartColumn        int     // Returns the index within the line of the first character where the error occurred.
	EndColumn          int     // Returns the index within the line of the last character where the error occurred.
	CallstackString    uintptr //const utf8* callstackString;
}

type JsKeys struct {
	Length uint
	Keys   **uintptr
}
