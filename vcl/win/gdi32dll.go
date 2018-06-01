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
