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
	. "github.com/ying32/govcl/vcl/types"
)

var (
	_GetWindowLongPtr = user32dll.NewProc("GetWindowLongW")
	_SetWindowLongPtr = user32dll.NewProc("SetWindowLongW")
)

func WindowFromPoint(point TPoint) HWND {
	r, _, _ := _WindowFromPoint.Call(uintptr(point.X), uintptr(point.Y))
	return HWND(r)
}

func WindowFromPhysicalPoint(point TPoint) HWND {
	r, _, _ := _WindowFromPhysicalPoint.Call(uintptr(point.Y), uintptr(point.Y))
	return HWND(r)
}

func ChildWindowFromPoint(hWndParent HWND, point TPoint) HWND {
	r, _, _ := _ChildWindowFromPoint.Call(uintptr(hWndParent), uintptr(point.Y), uintptr(point.Y))
	return HWND(r)
}
