//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package rtl

import (
	. "github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

// exe自身实例。
//
// Instance of an EXE.
func MainInstance() uintptr {
	return win.GetSelfModuleHandle()
}

func SendMessage(hWd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return win.SendMessage(hWd, msg, wParam, lParam)
}

func PostMessage(hWd HWND, msg uint32, wParam, lParam uintptr) bool {
	return win.PostMessage(hWd, msg, wParam, lParam) != 0
}

func IsIconic(hWnd HWND) bool {
	return win.IsIconic(hWnd)
}

func IsWindow(hWnd HWND) bool {
	return win.IsWindow(hWnd)
}

func IsZoomed(hWnd HWND) bool {
	return win.IsZoomed(hWnd)
}

func IsWindowVisible(hWnd HWND) bool {
	return win.IsWindowVisible(hWnd)
}

func GetDC(hWnd HWND) HDC {
	return win.GetDC(hWnd)
}

func ReleaseDC(hWnd HWND, dc HDC) int {
	return win.ReleaseDC(hWnd, dc)
}

func SetForegroundWindow(hWnd HWND) bool {
	return win.SetForegroundWindow(hWnd)
}

func WindowFromPoint(point TPoint) HWND {
	return win.WindowFromPoint(point)
}
