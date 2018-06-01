// +build windows

package win

import (
	"syscall"
	"unsafe"

	. "github.com/ying32/govcl/vcl/types"
)

var (
	// user32.dll
	user32dll         = syscall.NewLazyDLL("user32.dll")
	_MessageBox       = user32dll.NewProc("MessageBoxW")
	_LoadIcon         = user32dll.NewProc("LoadIconW")
	_GetClientRect    = user32dll.NewProc("GetClientRect")
	_GetSystemMetrics = user32dll.NewProc("GetSystemMetrics")

	_CallWindowProc = user32dll.NewProc("CallWindowProcW")

	_PostMessage = user32dll.NewProc("PostMessageW")
	_SendMessage = user32dll.NewProc("SendMessageW")

	// 两个未公开的api
	// _ChangeWindowMessageFilter 适合于 Win Vista 系统，Win7及以上系统并使用下
	_ChangeWindowMessageFilter = user32dll.NewProc("ChangeWindowMessageFilter")
	// ChangeWindowMessageFilterEx 适合于 Win7 及以上系统, Vista 无此函数
	_ChangeWindowMessageFilterEx = user32dll.NewProc("ChangeWindowMessageFilterEx")

	_GetDC               = user32dll.NewProc("GetDC")
	_ReleaseDC           = user32dll.NewProc("ReleaseDC")
	_UpdateLayeredWindow = user32dll.NewProc("UpdateLayeredWindow")

	_GetCapture     = user32dll.NewProc("GetCapture")
	_SetCapture     = user32dll.NewProc("SetCapture")
	_ReleaseCapture = user32dll.NewProc("ReleaseCapture")
)

// MessageBox 消息框
func MessageBox(hWnd uintptr, lpText, lpCaption string, uType uint32) int32 {
	r, _, _ := _MessageBox.Call(hWnd, CStr(lpText), CStr(lpCaption), uintptr(uType))
	return int32(r)
}

// LoadIcon 从实例资源中加载icon
func LoadIcon(hInstance uintptr, lpIconName int) HICON {
	r, _, _ := _LoadIcon.Call(hInstance, uintptr(lpIconName))
	return HICON(r)
}

// LoadIcon2 从实例资源中加载icon
func LoadIcon2(hInstance uintptr, lpIconName string) HICON {
	r, _, _ := _LoadIcon.Call(hInstance, CStr(lpIconName))
	return HICON(r)
}

// GetClientRect 获取指定句柄客户区矩形
func GetClientRect(hWnd HWND) TRect {
	r := TRect{}
	_GetClientRect.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&r)))
	return r
}

// GetSystemMetrics
func GetSystemMetrics(nIndex int32) int32 {
	r, _, _ := _GetSystemMetrics.Call(uintptr(nIndex))
	return int32(r)
}

// GetWindowLongPtr
func GetWindowLongPtr(hWnd HWND, nIndex int32) int {
	r, _, _ := _GetWindowLongPtr.Call(uintptr(hWnd), uintptr(nIndex))
	return int(r)
}

// SetWindowLongPtr
func SetWindowLongPtr(hWnd HWND, nIndex int32, dwNewLong uintptr) uintptr {
	r, _, _ := _SetWindowLongPtr.Call(uintptr(hWnd), uintptr(nIndex), dwNewLong)
	return r
}

// CallWindowProc
func CallWindowProc(lpPrevWndFunc uintptr, hWnd HWND, Msg uint32, wParam, lParam uintptr) uintptr {
	r, _, _ := _CallWindowProc.Call(lpPrevWndFunc, uintptr(hWnd), uintptr(Msg), wParam, lParam)
	return r
}

// PostMessage
func PostMessage(hWd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	r, _, _ := _PostMessage.Call(uintptr(hWd), uintptr(msg), wParam, lParam)
	return r
}

// SendMessage
func SendMessage(hWd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	r, _, _ := _SendMessage.Call(uintptr(hWd), uintptr(msg), wParam, lParam)
	return r
}

// 两个Windows未公开函数

const (
	// ChangeWindowMessageFilter  的常量
	MSGFLT_ADD    = 1
	MSGFLT_REMOVE = 2

	// ChangeWindowMessageFilterEx的常量
	MSGFLT_RESET    = 0
	MSGFLT_ALLOW    = 1
	MSGFLT_DISALLOW = 2

	//
	WM_COPYGLOBALDATA = 0x0049
)

// 消息过滤的传入结构
type TChangeFilterStruct struct {
	CbSize    uint32
	ExtStatus uint32
}

// ChangeWindowMessageFilter
func ChangeWindowMessageFilter(message, flags uint32) bool {
	r, _, _ := _ChangeWindowMessageFilter.Call(uintptr(message), uintptr(flags))
	return r != 0
}

// ChangeWindowMessageFilterEx
func ChangeWindowMessageFilterEx(hWd HWND, message, action uint32, pChangeFilterStruct *TChangeFilterStruct) bool {
	if _ChangeWindowMessageFilterEx.Find() != nil {
		return false
	}
	r, _, _ := _ChangeWindowMessageFilterEx.Call(uintptr(hWd), uintptr(message), uintptr(action), uintptr(unsafe.Pointer(pChangeFilterStruct)))
	return r != 0
}

func GetDC(hWnd HWND) HDC {
	r, _, _ := _GetDC.Call(uintptr(hWnd))
	return HDC(r)
}

func ReleaseDC(hWnd HWND, dc HDC) int {
	r, _, _ := _ReleaseDC.Call(uintptr(hWnd), uintptr(dc))
	return int(r)
}

func UpdateLayeredWindow(handle HWND, hdcDest HDC, pptDst *TPoint, psize *TSize,
	hdcSrc HDC, pptSrc *TPoint, crKey uint32, pblend *TBlendFunction, dwFlags uint32) bool {
	r, _, _ := _UpdateLayeredWindow.Call(uintptr(handle), uintptr(hdcDest), uintptr(unsafe.Pointer(pptDst)), uintptr(unsafe.Pointer(psize)),
		uintptr(hdcSrc), uintptr(unsafe.Pointer(pptSrc)), uintptr(crKey), uintptr(unsafe.Pointer(pblend)), uintptr(dwFlags))
	return r != 0
}

func GetCapture() HWND {
	r, _, _ := _GetCapture.Call()
	return HWND(r)
}

func SetCapture(hWnd HWND) HWND {
	r, _, _ := _SetCapture.Call(uintptr(hWnd))
	return HWND(r)
}

func ReleaseCapture() bool {
	r, _, _ := _ReleaseCapture.Call()
	return r != 0
}
