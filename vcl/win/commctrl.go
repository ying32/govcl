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
	"unsafe"

	"github.com/ying32/govcl/vcl/types"
)

const (
	LVM_FIRST = 0x1000 /* ListView messages */
	TV_FIRST  = 0x1100 /* TreeView messages */

)

// ListView
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

// TreeView
const (
	TVM_GETITEMW = TV_FIRST + 62
	TVM_GETITEM  = TVM_GETITEMW

	TVM_SETITEMW = TV_FIRST + 63
	TVM_SETITEM  = TVM_SETITEMW

	/* ====== TREEVIEW CONTROL =================== */

	TVS_HASBUTTONS      = 0x0001
	TVS_HASLINES        = 0x0002
	TVS_LINESATROOT     = 0x0004
	TVS_EDITLABELS      = 0x0008
	TVS_DISABLEDRAGDROP = 0x0010
	TVS_SHOWSELALWAYS   = 0x0020
	TVS_RTLREADING      = 0x0040
	TVS_NOTOOLTIPS      = 0x0080
	TVS_CHECKBOXES      = 0x0100
	TVS_TRACKSELECT     = 0x0200
	TVS_SINGLEEXPAND    = 0x0400
	TVS_INFOTIP         = 0x0800
	TVS_FULLROWSELECT   = 0x1000
	TVS_NOSCROLL        = 0x2000
	TVS_NONEVENHEIGHT   = 0x4000
	/* For IE >= 0x0500 */
	TVS_NOHSCROLL = 0x8000 // TVS_NOSCROLL overrides this

	/* For Windows >= Vista */
	TVS_EX_MULTISELECT         = 0x0002
	TVS_EX_DOUBLEBUFFER        = 0x0004
	TVS_EX_NOINDENTSTATE       = 0x0008
	TVS_EX_RICHTOOLTIP         = 0x0010
	TVS_EX_AUTOHSCROLL         = 0x0020
	TVS_EX_FADEINOUTEXPANDOS   = 0x0040
	TVS_EX_PARTIALCHECKBOXES   = 0x0080
	TVS_EX_EXCLUSIONCHECKBOXES = 0x0100
	TVS_EX_DIMMEDCHECKBOXES    = 0x0200
	TVS_EX_DRAWIMAGEASYNC      = 0x0400

	TVIF_TEXT          = 0x0001
	TVIF_IMAGE         = 0x0002
	TVIF_PARAM         = 0x0004
	TVIF_STATE         = 0x0008
	TVIF_HANDLE        = 0x0010
	TVIF_SELECTEDIMAGE = 0x0020
	TVIF_CHILDREN      = 0x0040
	TVIF_INTEGRAL      = 0x0080
	/* For Windows >= Vista */
	TVIF_STATEEX       = 0x0100
	TVIF_EXPANDEDIMAGE = 0x0200

	TVIS_FOCUSED       = 0x0001
	TVIS_SELECTED      = 0x0002
	TVIS_CUT           = 0x0004
	TVIS_DROPHILITED   = 0x0008
	TVIS_BOLD          = 0x0010
	TVIS_EXPANDED      = 0x0020
	TVIS_EXPANDEDONCE  = 0x0040
	TVIS_EXPANDPARTIAL = 0x0080

	TVIS_OVERLAYMASK    = 0x0F00
	TVIS_STATEIMAGEMASK = 0xF000
	TVIS_USERMASK       = 0xF000

	/* For IE >= 0x0600 */
	TVIS_EX_FLAT = 0x0001
	TVIS_EX_ALL  = 0x0002
	/* For Windows >= Vista */
	TVIS_EX_DISABLED = 0x0002

	//

	TVIS_CHECKED = 0x2000
)

// ----------- ListView

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

// -------------- TreeView

type _TREEITEM struct {
}

type HTreeItem = uintptr // _TREEITEM ptr

type TTVItem struct {
	Mask           uint32
	HItem          HTreeItem
	State          uint32
	StateMask      uint32
	PszText        uintptr
	CchTextMax     int32
	IImage         int32
	ISelectedImage int32
	CChildren      int32
	LParam         uintptr
}

// 下面这类api参考  Winapi.CommCtrl单元
// treeview api
func TreeView_GetItem(hWd types.HWND, pitem *TTVItem) bool {
	return SendMessage(hWd, TVM_GETITEM, 0, uintptr(unsafe.Pointer(pitem))) != 0
}

func TreeView_SetItem(hWd types.HWND, pitem TTVItem) bool {
	return SendMessage(hWd, TVM_SETITEM, 0, uintptr(unsafe.Pointer(&pitem))) != 0
}
