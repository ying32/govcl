package main

import (
	"unsafe"

	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

const (
	LVM_SUBITEMHITTEST  = win.LVM_FIRST + 57
	LVM_GETITEMRECT     = win.LVM_FIRST + 14
	LVM_GETCOLUMNWIDTH  = win.LVM_FIRST + 29
	LVM_GETITEMPOSITION = win.LVM_FIRST + 16
)

type TLVHitTestInfo struct {
	pt       types.TPoint
	flags    uint32
	iItem    int32
	iSubItem int32 // this is was NOT in win95.  valid only for LVM_SUBITEMHITTEST
	iGroup   int32 // readonly. index of group. only valid for owner data.
	// supports single item in multiple groups.
}

func ListView_SubItemHitTest(hwndLV types.HWND, plvhti *TLVHitTestInfo) int32 {
	return int32(win.SendMessage(hwndLV, LVM_SUBITEMHITTEST, 0, uintptr(unsafe.Pointer(plvhti))))
}

func ListView_GetItemRect(hWnd types.HWND, i int32, prc *types.TRect, Code int32) bool {
	if prc != nil {
		prc.Left = Code
		return win.SendMessage(hWnd, LVM_GETITEMRECT, uintptr(i), uintptr(unsafe.Pointer(prc))) != 0
	} else {
		return win.SendMessage(hWnd, LVM_GETITEMRECT, uintptr(i), 0) != 0
	}
}

func ListView_GetColumnWidth(hwnd types.HWND, iCol int32) int32 {
	return int32(win.SendMessage(hwnd, LVM_GETCOLUMNWIDTH, uintptr(iCol), 0))
}

func ListView_GetItemPosition(hWndLV types.HWND, i int32, ppt *types.TPoint) bool {
	return win.SendMessage(hWndLV, LVM_GETITEMPOSITION, uintptr(i), uintptr(unsafe.Pointer(ppt))) != 0
}
