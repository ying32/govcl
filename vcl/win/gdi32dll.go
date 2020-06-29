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
	gdi32dll = syscall.NewLazyDLL("gdi32.dll")

	_CreateCompatibleDC = gdi32dll.NewProc("CreateCompatibleDC")
	_CreateDIBSection   = gdi32dll.NewProc("CreateDIBSection")
	_SelectObject       = gdi32dll.NewProc("SelectObject")
	_DeleteDC           = gdi32dll.NewProc("DeleteDC")
	_DeleteObject       = gdi32dll.NewProc("DeleteObject")
	_CreatePen          = gdi32dll.NewProc("CreatePen")
	_SetROP2            = gdi32dll.NewProc("SetROP2")
	_Rectangle          = gdi32dll.NewProc("Rectangle")
	_SaveDC             = gdi32dll.NewProc("SaveDC")
	_RestoreDC          = gdi32dll.NewProc("RestoreDC")
	_SetBkMode          = gdi32dll.NewProc("SetBkMode")
	_SetTextColor       = gdi32dll.NewProc("SetTextColor")
)

func CreateCompatibleDC(dc HDC) HDC {
	r, _, _ := _CreateCompatibleDC.Call(uintptr(dc))
	return HDC(r)
}

func CreateDIBSection(dc HDC, p2 *TBitmapInfo, p3 uint32, p4 *uintptr, p5 uintptr, p6 uint32) HBITMAP {
	r, _, _ := _CreateDIBSection.Call(uintptr(dc), uintptr(unsafe.Pointer(p2)), uintptr(p3), uintptr(unsafe.Pointer(p4)), p5, uintptr(p6))
	return HBITMAP(r)
}

func SelectObject(dc HDC, p2 HGDIOBJ) HGDIOBJ {
	r, _, _ := _SelectObject.Call(uintptr(dc), uintptr(p2))
	return HGDIOBJ(r)
}

func DeleteDC(dc HDC) bool {
	r, _, _ := _DeleteDC.Call(uintptr(dc))
	return r != 0
}

func DeleteObject(p1 HGDIOBJ) bool {
	r, _, _ := _DeleteObject.Call(uintptr(p1))
	return r != 0
}

func CreatePen(style, width int32, color uint32) HPEN {
	r, _, _ := _CreatePen.Call(uintptr(style), uintptr(width), uintptr(color))
	return HPEN(r)
}

func SetROP2(dc HDC, p2 int32) int32 {
	r, _, _ := _SetROP2.Call(uintptr(dc), uintptr(p2))
	return int32(r)
}

func Rectangle(dc HDC, x1, y1, x2, y2 int32) bool {
	r, _, _ := _Rectangle.Call(uintptr(dc), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return r != 0
}

func SetBkMode(dC HDC, bkMode int32) int32 {
	r, _, _ := _SetBkMode.Call(uintptr(dC), uintptr(bkMode))
	return int32(r)
}

func SaveDC(dC HDC) int32 {
	r, _, _ := _SaveDC.Call(uintptr(dC))
	return int32(r)
}

func RestoreDC(dC HDC, savedDC int32) bool {
	r, _, _ := _RestoreDC.Call(uintptr(dC), uintptr(savedDC))
	return r != 0
}

func SetTextColor(dC HDC, color uint32) uint32 {
	r, _, _ := _SetTextColor.Call(uintptr(dC), uintptr(color))
	return uint32(r)
}
