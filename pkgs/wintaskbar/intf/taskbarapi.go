//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
//+build windows

package intf

import (
	"github.com/go-ole/go-ole"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

type TThumbButton struct {
	DwMask  uint32
	IId     uint32
	IBitmap uint32
	HIcon   types.HICON
	SzTip   [260]uint16
	DwFlags uint32
}

const (
	MAX_BUTTON_COUNT = 7
)

// THUMBBUTTON flags
const (
	THBF_ENABLED        = 0x0000
	THBF_DISABLED       = 0x0001
	THBF_DISMISSONCLICK = 0x0002
	THBF_NOBACKGROUND   = 0x0004
	THBF_HIDDEN         = 0x0008
	THBF_NONINTERACTIVE = 0x10
	// THUMBBUTTON mask
	THB_BITMAP   = 0x0001
	THB_ICON     = 0x0002
	THB_TOOLTIP  = 0x0004
	THB_FLAGS    = 0x0008
	THBN_CLICKED = 0x1800
)

const (
	TBPF_NOPROGRESS    = 0
	TBPF_INDETERMINATE = 0x1
	TBPF_NORMAL        = 0x2
	TBPF_ERROR         = 0x4
	TBPF_PAUSED        = 0x8

	TBATF_USEMDITHUMBNAIL   = 0x1
	TBATF_USEMDILIVEPREVIEW = 0x2
)

var (
	WM_TASKBARBUTTONCREATED uint32

	CLSID_TaskbarList = ole.NewGUID("{56FDF344-FD6D-11d0-958A-006097C9A090}")
	SID_ITaskbarList4 = ole.NewGUID("{C43DC798-95D1-4BEA-9030-BB99E2983A1A}")
	SID_ITaskbarList3 = ole.NewGUID("{EA1AFB91-9E28-4B86-90E9-9E9F8A5EEFAF}")
)

func init() {
	WM_TASKBARBUTTONCREATED = win.RegisterWindowMessage("TaskbarButtonCreated")
}
