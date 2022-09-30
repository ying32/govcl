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

	"github.com/ying32/govcl/vcl/api/dllimports"

	. "github.com/ying32/govcl/vcl/types"
)

func DSendMessage(hWd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return defSyscallN(dllimports.DSENDMESSAGE, hWd, uintptr(msg), wParam, lParam)
}

func DPostMessage(hWd HWND, msg uint32, wParam, lParam uintptr) bool {
	return GoBool(defSyscallN(dllimports.DPOSTMESSAGE, hWd, uintptr(msg), wParam, lParam))
}

func DIsIconic(hWnd HWND) bool {
	return GoBool(defSyscallN(dllimports.DISICONIC, hWnd))
}

func DIsWindow(hWnd HWND) bool {
	return GoBool(defSyscallN(dllimports.DISWINDOW, hWnd))
}

func DIsZoomed(hWnd HWND) bool {
	return GoBool(defSyscallN(dllimports.DISZOOMED, hWnd))
}

func DIsWindowVisible(hWnd HWND) bool {
	return GoBool(defSyscallN(dllimports.DISWINDOWVISIBLE, hWnd))
}

func DGetDC(hWnd HWND) HDC {
	return defSyscallN(dllimports.DGETDC, hWnd)
}

func DReleaseDC(hWnd HWND, dc HDC) int {
	return int(defSyscallN(dllimports.DRELEASEDC, hWnd, dc))
}

func DSetForegroundWindow(hWnd HWND) bool {
	return GoBool(defSyscallN(dllimports.DSETFOREGROUNDWINDOW, hWnd))
}

func DWindowFromPoint(point TPoint) HWND {
	return defSyscallN(dllimports.DWINDOWFROMPOINT, uintptr(unsafe.Pointer(&point)))
}
