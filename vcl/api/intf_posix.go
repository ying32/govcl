//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package api

import (
	"unsafe"

	. "github.com/ying32/govcl/vcl/types"
)

var (
	dSendMessage         = newDLLProc("DSendMessage")
	dPostMessage         = newDLLProc("DPostMessage")
	dIsIconic            = newDLLProc("DIsIconic")
	dIsWindow            = newDLLProc("DIsWindow")
	dIsZoomed            = newDLLProc("DIsZoomed")
	dIsWindowVisible     = newDLLProc("DIsWindowVisible")
	dGetDC               = newDLLProc("DGetDC")
	dReleaseDC           = newDLLProc("DReleaseDC")
	dSetForegroundWindow = newDLLProc("DSetForegroundWindow")
	dWindowFromPoint     = newDLLProc("DWindowFromPoint")
)

func DSendMessage(hWd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	r, _, _ := dSendMessage.Call(hWd, uintptr(msg), wParam, lParam)
	return r
}

func DPostMessage(hWd HWND, msg uint32, wParam, lParam uintptr) bool {
	r, _, _ := dPostMessage.Call(hWd, uintptr(msg), wParam, lParam)
	return r != 0
}

func DIsIconic(hWnd HWND) bool {
	r, _, _ := dIsIconic.Call(uintptr(hWnd))
	return r != 0
}

func DIsWindow(hWnd HWND) bool {
	r, _, _ := dIsWindow.Call(uintptr(hWnd))
	return r != 0
}

func DIsZoomed(hWnd HWND) bool {
	r, _, _ := dIsZoomed.Call(uintptr(hWnd))
	return r != 0
}

func DIsWindowVisible(hWnd HWND) bool {
	r, _, _ := dIsWindowVisible.Call(uintptr(hWnd))
	return r != 0
}

func DGetDC(hWnd HWND) HDC {
	r, _, _ := dGetDC.Call(uintptr(hWnd))
	return HDC(r)
}

func DReleaseDC(hWnd HWND, dc HDC) int {
	r, _, _ := dReleaseDC.Call(uintptr(hWnd), uintptr(dc))
	return int(r)
}

func DSetForegroundWindow(hWnd HWND) bool {
	r, _, _ := dSetForegroundWindow.Call(uintptr(hWnd))
	return r != 0
}

func DWindowFromPoint(point TPoint) HWND {
	r, _, _ := dWindowFromPoint.Call(uintptr(unsafe.Pointer(&point)))
	return HWND(r)
}
