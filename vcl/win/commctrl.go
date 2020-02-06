// +build windows

package win

import (
	"unsafe"

	"github.com/ying32/govcl/vcl/types"
)

const (
	LVM_FIRST = 0x1000 /* ListView messages */
	TV_FIRST  = 0x1100 /* TreeView messages */

)

const (
	// ListView
	LVM_SUBITEMHITTEST  = LVM_FIRST + 57
	LVM_GETITEMRECT     = LVM_FIRST + 14
	LVM_GETCOLUMNWIDTH  = LVM_FIRST + 29
	LVM_GETITEMPOSITION = LVM_FIRST + 16
	LVM_GETSUBITEMRECT  = LVM_FIRST + 56

	LVIR_BOUNDS       = 0
	LVIR_ICON         = 1
	LVIR_LABEL        = 2
	LVIR_SELECTBOUNDS = 3
)

const (
	TVM_GETITEMW = TV_FIRST + 62
	TVM_GETITEM  = TVM_GETITEMW

	TVM_SETITEMW = TV_FIRST + 63
	TVM_SETITEM  = TVM_SETITEMW
)

type TLVHitTestInfo struct {
	Pt       types.TPoint
	Flags    uint32
	IItem    int32
	ISubItem int32 // this is was NOT in win95.  valid only for LVM_SUBITEMHITTEST
	IGroup   int32 // readonly. index of group. only valid for owner data.
	// supports single item in multiple groups.
}

func ListView_SubItemHitTest(hwndLV types.HWND, plvhti *TLVHitTestInfo) int32 {
	return int32(SendMessage(hwndLV, LVM_SUBITEMHITTEST, 0, uintptr(unsafe.Pointer(plvhti))))
}

func ListView_GetItemRect(hWnd types.HWND, i int32, prc *types.TRect, Code int32) bool {
	if prc != nil {
		prc.Left = Code
		return SendMessage(hWnd, LVM_GETITEMRECT, uintptr(i), uintptr(unsafe.Pointer(prc))) != 0
	} else {
		return SendMessage(hWnd, LVM_GETITEMRECT, uintptr(i), 0) != 0
	}
}

func ListView_GetColumnWidth(hwnd types.HWND, iCol int32) int32 {
	return int32(SendMessage(hwnd, LVM_GETCOLUMNWIDTH, uintptr(iCol), 0))
}

func ListView_GetItemPosition(hWndLV types.HWND, i int32, ppt *types.TPoint) bool {
	return SendMessage(hWndLV, LVM_GETITEMPOSITION, uintptr(i), uintptr(unsafe.Pointer(ppt))) != 0
}

func ListView_GetSubItemRect(hwndLV types.HWND, iItem, iSubItem int32, code uint32, prc *types.TRect) bool {
	if prc != nil {
		prc.Top = iSubItem
		prc.Left = int32(code)
	}
	return SendMessage(hwndLV, LVM_GETSUBITEMRECT, uintptr(iItem), uintptr(unsafe.Pointer(prc))) != 0
}
