//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

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

	_GetCursor                  = user32dll.NewProc("GetCursor")
	_GetCursorInfo              = user32dll.NewProc("GetCursorInfo")
	_GetCursorPos               = user32dll.NewProc("GetCursorPos")
	_ShowCursor                 = user32dll.NewProc("ShowCursor")
	_SetCursorPos               = user32dll.NewProc("SetCursorPos")
	_WindowFromPoint            = user32dll.NewProc("WindowFromPoint")
	_GetClassName               = user32dll.NewProc("GetClassNameW")
	_FindWindow                 = user32dll.NewProc("FindWindowW")
	_FindWindowEx               = user32dll.NewProc("FindWindowExW")
	_GetClassInfo               = user32dll.NewProc("GetClassInfoW")
	_RegisterWindowMessage      = user32dll.NewProc("RegisterWindowMessageW")
	_GetWindowThreadProcessId   = user32dll.NewProc("GetWindowThreadProcessId")
	_GetWindowRect              = user32dll.NewProc("GetWindowRect")
	_GetClientRect              = user32dll.NewProc("GetClientRect")
	_GetWindowText              = user32dll.NewProc("GetWindowTextW")
	_GetWindowRgn               = user32dll.NewProc("GetWindowRgn")
	_IsWindow                   = user32dll.NewProc("IsWindow")
	_IsWindowEnabled            = user32dll.NewProc("IsWindowEnabled")
	_IsWindowUnicode            = user32dll.NewProc("IsWindowUnicode")
	_IsWindowVisible            = user32dll.NewProc("IsWindowVisible")
	_EnableWindow               = user32dll.NewProc("EnableWindow")
	_IsZoomed                   = user32dll.NewProc("IsZoomed")
	_IsRectEmpty                = user32dll.NewProc("IsRectEmpty")
	_IsMenu                     = user32dll.NewProc("IsMenu")
	_IsIconic                   = user32dll.NewProc("IsIconic")
	_OpenIcon                   = user32dll.NewProc("OpenIcon")
	_CloseWindow                = user32dll.NewProc("CloseWindow")
	_MoveWindow                 = user32dll.NewProc("MoveWindow")
	_SetWindowPos               = user32dll.NewProc("SetWindowPos")
	_FlashWindow                = user32dll.NewProc("FlashWindow")
	_ShowWindow                 = user32dll.NewProc("ShowWindow")
	_SetLayeredWindowAttributes = user32dll.NewProc("SetLayeredWindowAttributes")
	_ScreenToClient             = user32dll.NewProc("ScreenToClient")
	_ClientToScreen             = user32dll.NewProc("ClientToScreen")
	_ClipCursor                 = user32dll.NewProc("ClipCursor")
	_GetClipCursor              = user32dll.NewProc("GetClipCursor")
	_WindowFromPhysicalPoint    = user32dll.NewProc("WindowFromPhysicalPoint")
	_ChildWindowFromPoint       = user32dll.NewProc("ChildWindowFromPoint")
	_GetWindowDC                = user32dll.NewProc("GetWindowDC")
	_IsChild                    = user32dll.NewProc("IsChild")
	_AnimateWindow              = user32dll.NewProc("AnimateWindow")
	_ShowWindowAsync            = user32dll.NewProc("ShowWindowAsync")
	_InvalidateRect             = user32dll.NewProc("InvalidateRect")
	_InvalidateRgn              = user32dll.NewProc("InvalidateRgn")
	_RedrawWindow               = user32dll.NewProc("RedrawWindow")

	_SetForegroundWindow = user32dll.NewProc("SetForegroundWindow")

	_EnumWindows      = user32dll.NewProc("EnumWindows")
	_EnumChildWindows = user32dll.NewProc("EnumChildWindows")

	_UnregisterHotKey = user32dll.NewProc("UnregisterHotKey")
	_RegisterHotKey   = user32dll.NewProc("RegisterHotKey")

	_DrawText         = user32dll.NewProc("DrawTextW")
	_GetDesktopWindow = user32dll.NewProc("GetDesktopWindow")
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

// GetWindowLong
func GetWindowLong(hWnd HWND, nIndex int32) int {
	return GetWindowLongPtr(hWnd, nIndex)
}

// SetWindowLongPtr
func SetWindowLongPtr(hWnd HWND, nIndex int32, dwNewLong uintptr) uintptr {
	r, _, _ := _SetWindowLongPtr.Call(uintptr(hWnd), uintptr(nIndex), dwNewLong)
	return r
}

// SetWindowLong
func SetWindowLong(hWnd HWND, nIndex int32, dwNewLong uintptr) uintptr {
	return SetWindowLongPtr(hWnd, nIndex, dwNewLong)
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

const (
	// ChangeWindowMessageFilter  的常量, vista+
	MSGFLT_ADD    = 1
	MSGFLT_REMOVE = 2

	// ChangeWindowMessageFilterEx的常量, win7+
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

func GetCursor() HCURSOR {
	r, _, _ := _GetCursor.Call()
	return HCURSOR(r)
}

func GetCursorInfo(pci *TCursorInfo) bool {
	if pci != nil && pci.CbSize == 0 {
		pci.CbSize = uint32(unsafe.Sizeof(*pci))
	}
	r, _, _ := _GetCursorInfo.Call(uintptr(unsafe.Pointer(pci)))
	return r != 0
}

func GetCursorInfo2() (pci TCursorInfo, result bool) {
	result = GetCursorInfo(&pci)
	return
}

func GetCursorPos(pt *TPoint) bool {
	r, _, _ := _GetCursorPos.Call(uintptr(unsafe.Pointer(pt)))
	return r != 0
}

func GetCursorPos2() (pt TPoint, result bool) {
	result = GetCursorPos(&pt)
	return
}

func ShowCursor(bShow bool) int32 {
	r, _, _ := _ShowCursor.Call(CBool(bShow))
	return int32(r)
}

func SetCursorPos(x, y int32) bool {
	r, _, _ := _SetCursorPos.Call(uintptr(x), uintptr(y))
	return r != 0
}

func GetClassName(hWnd HWND) (string, int32) {
	buff := make([]uint16, 255)
	r, _, _ := _GetClassName.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&buff[0])), 255)
	return GoStr(buff), int32(r)
}

func FindWindow(lpClassName, lpWindowName string) HWND {
	r, _, _ := _FindWindow.Call(CStr(lpClassName), CStr(lpWindowName))
	return HWND(r)
}

func FindWindowEx(parent, child HWND, className, windowName string) HWND {
	r, _, _ := _FindWindowEx.Call(uintptr(parent), uintptr(child), CStr(className), CStr(windowName))
	return HWND(r)
}

func GetClassInfo(hInstance uintptr, lpClassName string, lpWndClass *TWndClass) bool {
	r, _, _ := _GetClassInfo.Call(hInstance, CStr(lpClassName), uintptr(unsafe.Pointer(lpWndClass)))
	return r != 0
}

func GetClassInfo2(hInstance uintptr, lpClassName string) (lpWndClass TWndClass, result bool) {
	result = GetClassInfo(hInstance, lpClassName, &lpWndClass)
	return
}

func RegisterWindowMessage(lpString string) uint32 {
	r, _, _ := _RegisterWindowMessage.Call(CStr(lpString))
	return uint32(r)
}

func GetWindowThreadProcessId(hWnd HWND, dwProcessId *uint32) uint32 {
	r, _, _ := _GetWindowThreadProcessId.Call(uintptr(hWnd), uintptr(unsafe.Pointer(dwProcessId)))
	return uint32(r)
}

func GetWindowThreadProcessId2(hWnd HWND) (dwThreadId, dwProcessId uint32) {
	dwThreadId = GetWindowThreadProcessId(hWnd, &dwProcessId)
	return
}

func GetWindowRect(hWnd HWND, lpRect *TRect) bool {
	r, _, _ := _GetWindowRect.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpRect)))
	return r != 0
}

func GetWindowRect2(hWnd HWND) (lpRect TRect, result bool) {
	result = GetWindowRect(hWnd, &lpRect)
	return
}

func GetClientRect(hWnd HWND, lpRect *TRect) bool {
	r, _, _ := _GetClientRect.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpRect)))
	return r != 0
}

func GetClientRect2(hWnd HWND) (lpRect TRect, result bool) {
	result = GetClientRect(hWnd, &lpRect)
	return
}

func GetWindowText(hWnd HWND) (string, int32) {
	buff := make([]uint16, 255)
	r, _, _ := _GetWindowText.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&buff[0])), 255)
	return GoStr(buff), int32(r)
}

func GetWindowRgn(hWnd HWND, hRgn HRGN) int32 {
	r, _, _ := _GetWindowRgn.Call(uintptr(hWnd), uintptr(hRgn))
	return int32(r)
}

func IsWindow(hWnd HWND) bool {
	r, _, _ := _IsWindow.Call(uintptr(hWnd))
	return r != 0
}

func IsWindowEnabled(hWnd HWND) bool {
	r, _, _ := _IsWindowEnabled.Call(uintptr(hWnd))
	return r != 0
}

func IsWindowUnicode(hWnd HWND) bool {
	r, _, _ := _IsWindowUnicode.Call(uintptr(hWnd))
	return r != 0
}

func IsWindowVisible(hWnd HWND) bool {
	r, _, _ := _IsWindowVisible.Call(uintptr(hWnd))
	return r != 0
}

func EnableWindow(hWnd HWND, bEnable bool) bool {
	r, _, _ := _EnableWindow.Call(uintptr(hWnd), CBool(bEnable))
	return r != 0
}

func IsZoomed(hWnd HWND) bool {
	r, _, _ := _IsZoomed.Call(uintptr(hWnd))
	return r != 0
}

func IsRectEmpty(lprc TRect) bool {
	r, _, _ := _IsRectEmpty.Call(uintptr(unsafe.Pointer(&lprc)))
	return r != 0
}

func IsMenu(hWnd HWND) bool {
	r, _, _ := _IsMenu.Call(uintptr(hWnd))
	return r != 0
}

func IsIconic(hWnd HWND) bool {
	r, _, _ := _IsIconic.Call(uintptr(hWnd))
	return r != 0
}

func OpenIcon(hWnd HWND) bool {
	r, _, _ := _OpenIcon.Call(uintptr(hWnd))
	return r != 0
}

func CloseWindow(hWnd HWND) bool {
	r, _, _ := _CloseWindow.Call(uintptr(hWnd))
	return r != 0
}

func MoveWindow(hWnd HWND, x, y, nWidth, nHeight int32, bRepaint bool) bool {
	r, _, _ := _MoveWindow.Call(uintptr(hWnd), uintptr(x), uintptr(y), uintptr(nWidth), uintptr(nHeight), CBool(bRepaint))
	return r != 0
}

func SetWindowPos(hWnd, hWndInsertAfter HWND, x, y, cx, cy int32, uFlags uint32) bool {
	r, _, _ := _SetWindowPos.Call(uintptr(hWnd), uintptr(hWndInsertAfter), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return r != 0
}

func FlashWindow(hWnd HWND, bInvert bool) bool {
	r, _, _ := _FlashWindow.Call(uintptr(hWnd), CBool(bInvert))
	return r != 0
}

func ShowWindow(hWnd HWND, nCmdShow int32) bool {
	r, _, _ := _ShowWindow.Call(uintptr(hWnd), uintptr(nCmdShow))
	return r != 0
}

func SetLayeredWindowAttributes(hWnd HWND, crKey uint32, bAlpha byte, dwFlags uint32) bool {
	if _SetLayeredWindowAttributes.Find() != nil {
		return false
	}
	r, _, _ := _SetLayeredWindowAttributes.Call(uintptr(hWnd), uintptr(crKey), uintptr(bAlpha), uintptr(dwFlags))
	return r != 0
}

func ScreenToClient(hWnd HWND, lpPoint *TPoint) bool {
	r, _, _ := _ScreenToClient.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpPoint)))
	return r != 0
}

func ClientToScreen(hWnd HWND, lpPoint *TPoint) bool {
	r, _, _ := _ClientToScreen.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpPoint)))
	return r != 0
}

func ClipCursor(lpRect *TRect) bool {
	r, _, _ := _ClipCursor.Call(uintptr(unsafe.Pointer(lpRect)))
	return r != 0
}

func GetClipCursor(lpRect *TRect) bool {
	r, _, _ := _GetClipCursor.Call(uintptr(unsafe.Pointer(lpRect)))
	return r != 0
}

func GetClipCursor2() (lpRect TRect, result bool) {
	result = GetClipCursor(&lpRect)
	return
}

func GetWindowDC(hWnd HWND) HDC {
	r, _, _ := _GetWindowDC.Call(uintptr(hWnd))
	return HDC(r)
}

func IsChild(hWndParent, hWnd HWND) bool {
	r, _, _ := _IsChild.Call(uintptr(hWndParent), uintptr(hWnd))
	return r != 0
}

func AnimateWindow(hWnd HWND, dwTime, dwFlags uint32) bool {
	r, _, _ := _AnimateWindow.Call(uintptr(hWnd), uintptr(dwTime), uintptr(dwFlags))
	return r != 0
}

func ShowWindowAsync(hWnd HWND, nCmdShow int32) bool {
	r, _, _ := _ShowWindowAsync.Call(uintptr(hWnd), uintptr(nCmdShow))
	return r != 0
}

func InvalidateRect(hWnd HWND, lpRect *TRect, bErase bool) bool {
	r, _, _ := _InvalidateRect.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpRect)), CBool(bErase))
	return r != 0
}

func InvalidateRgn(hWnd HWND, hRgn HRGN, bErase bool) bool {
	r, _, _ := _InvalidateRgn.Call(uintptr(hWnd), uintptr(hRgn), CBool(bErase))
	return r != 0
}

func RedrawWindow(hWnd HWND, lprcUpdate *TRect, hrgnUpdate HRGN, flags uint32) bool {
	r, _, _ := _RedrawWindow.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lprcUpdate)), uintptr(hrgnUpdate), uintptr(flags))
	return r != 0
}

func SetForegroundWindow(hWnd HWND) bool {
	r, _, _ := _SetForegroundWindow.Call(uintptr(hWnd))
	return r != 0
}

func EnumWindows(lpEnumFunc TFNWndEnumProc, lParam uintptr) bool {
	r, _, _ := _EnumWindows.Call(uintptr(lpEnumFunc), lParam)
	return r != 0
}

func EnumChildWindows(hWndParent HWND, lpEnumFunc TFNWndEnumProc, lParam uintptr) bool {
	r, _, _ := _EnumChildWindows.Call(uintptr(hWndParent), uintptr(lpEnumFunc), lParam)
	return r != 0
}

func UnregisterHotKey(hWnd HWND, id int32) bool {
	r, _, _ := _UnregisterHotKey.Call(uintptr(hWnd), uintptr(id))
	return r != 0
}

func RegisterHotKey(hWnd HWND, id int32, fsModifiers, vk uint32) bool {
	r, _, _ := _RegisterHotKey.Call(uintptr(hWnd), uintptr(id), uintptr(fsModifiers), uintptr(vk))
	return r != 0
}

func DrawText(hDC HDC, lpString string, nCount int32, lpRect *TRect, uFormat uint32) int32 {
	r, _, _ := _DrawText.Call(uintptr(hDC), CStr(lpString), uintptr(nCount), uintptr(unsafe.Pointer(lpRect)), uintptr(uFormat))
	return int32(r)
}

func GetDesktopWindow() HWND {
	r, _, _ := _GetDesktopWindow.Call()
	return r
}
