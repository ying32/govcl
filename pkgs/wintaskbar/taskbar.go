//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
//+build windows

package wintaskbar

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/ying32/govcl/pkgs/wintaskbar/intf"

	"github.com/ying32/govcl/vcl/types/messages"

	"github.com/go-ole/go-ole"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

var (
	oldWndPrc          uintptr
	initCreateProc     func() error
	processMessageProc func(uintptr)
)

type TProcessState int32

const (
	NoProgress    TProcessState = intf.TBPF_NOPROGRESS
	Indeterminate               = intf.TBPF_INDETERMINATE
	Normal                      = intf.TBPF_NORMAL
	Error                       = intf.TBPF_ERROR
	Paused                      = intf.TBPF_PAUSED
)

type TButtonFlags uint32

const (
	Enabled        TButtonFlags = intf.THBF_ENABLED
	Disabled                    = intf.THBF_DISABLED
	DisMissonClick              = intf.THBF_DISMISSONCLICK
	NoBackGround                = intf.THBF_NOBACKGROUND
	Hidden                      = intf.THBF_HIDDEN
	Noninteractive              = intf.THBF_NONINTERACTIVE
)

type TThumbButtonClick func(index uint16)

type ThumbButton struct {
	taskBar *WinTaskBar
	button  intf.TThumbButton
}

type WinTaskBar struct {
	hWnd        types.HWND
	taskBarList intf.ITaskbarList4
	created     bool
	buttons     []*ThumbButton
	btnCallback TThumbButtonClick
}

func newThumbButton(taskBar *WinTaskBar) *ThumbButton {
	tt := new(ThumbButton)
	tt.taskBar = taskBar
	tt.button.DwMask = intf.THB_ICON | intf.THB_TOOLTIP | intf.THB_FLAGS
	tt.button.IId = uint32(len(taskBar.buttons))
	tt.button.DwFlags = uint32(Enabled)
	taskBar.buttons = append(taskBar.buttons, tt)
	return tt
}

func (tt *ThumbButton) changed() {
	if tt.taskBar.created {
		tt.taskBar.changed()
	}
}

func (tt *ThumbButton) SetHint(str string) {
	s := syscall.StringToUTF16(str)
	copy(tt.button.SzTip[:], s)
	tt.changed()
}

func (tt *ThumbButton) SetIcon(h types.HICON) {
	tt.button.HIcon = h
	tt.changed()
}

func (tt *ThumbButton) SetFlags(flags TButtonFlags) {
	tt.button.DwFlags = uint32(flags)
	tt.changed()
}

//---------------------------------------------------------------------------------------------------------------------

func NewWinTaskBar(hWnd types.HWND) *WinTaskBar {
	t := new(WinTaskBar)
	// Minimum supported client	Windows 7 [desktop apps only]
	// Minimum supported server	Windows Server 2008 R2 [desktop apps only]
	if !win.CheckWin32Version(6, 1) {
		return t
	}
	t.hWnd = hWnd
	// 只有一个窗口处理这个
	if oldWndPrc == 0 {
		initCreateProc = t.createTaskBar
		processMessageProc = t.processMessage
		oldWndPrc = win.SetWindowLongPtr(t.hWnd, win.GWL_WNDPROC, syscall.NewCallback(newWndProc))
	}
	return t
}

func (t *WinTaskBar) SetOnThumbButtonClick(event TThumbButtonClick) {
	t.btnCallback = event
}

func (t *WinTaskBar) getButtons() []intf.TThumbButton {
	r := make([]intf.TThumbButton, 0)
	for _, b := range t.buttons {
		r = append(r, b.button)
	}
	return r
}

func (t *WinTaskBar) createTaskBar() error {
	if t.created {
		return nil
	}
	ole.CoInitializeEx(0, 0)
	defer ole.CoUninitialize()

	unknown, err := ole.CreateInstance(intf.CLSID_TaskbarList, intf.SID_ITaskbarList4)
	if err != nil {
		return err
	}
	t.taskBarList = intf.ITaskbarList4(uintptr(unsafe.Pointer(unknown)))
	err = t.taskBarList.HrInit()
	if err != nil {
		return err
	}

	err = t.taskBarList.ThumbBarAddButtons(t.hWnd, t.getButtons())
	if err != nil {
		return err
	}
	t.created = true

	return nil
}

func (t *WinTaskBar) SetThumbnailTooltip(tip string) error {
	return t.taskBarList.SetThumbnailTooltip(t.hWnd, tip)
}

func (t *WinTaskBar) SetProgressValue(ullCompleted, ullTotal uint64) error {
	return t.taskBarList.SetProgressValue(t.hWnd, ullCompleted, ullTotal)
}

func (t *WinTaskBar) SetProgressState(flags TProcessState) error {
	return t.taskBarList.SetProgressState(t.hWnd, int32(flags))
}

func (t *WinTaskBar) SetOverlayIcon(hIcon types.HICON, pszDescription string) error {
	return t.taskBarList.SetOverlayIcon(t.hWnd, hIcon, pszDescription)
}

func (t *WinTaskBar) Buttons() []*ThumbButton {
	return t.buttons
}

func (t *WinTaskBar) AddButton() (*ThumbButton, error) {
	if len(t.buttons) > intf.MAX_BUTTON_COUNT {
		return nil, errors.New(fmt.Sprintf("len(buttons) >  %d", intf.MAX_BUTTON_COUNT))
	}
	if t.created {
		return nil, errors.New("已创建，不能再进行添加操作。")
	}
	return newThumbButton(t), nil
}

func (t *WinTaskBar) changed() {
	t.taskBarList.ThumbBarUpdateButtons(t.hWnd, t.getButtons())
}

func hiWord(L uint32) uint16 {
	return uint16(L >> 16)
}

func (t *WinTaskBar) processMessage(wParam uintptr) {
	if t.btnCallback != nil && hiWord(uint32(wParam)) == intf.THBN_CLICKED {
		t.btnCallback(uint16(uint32(wParam)))
	}
}

func (t *WinTaskBar) Free() {
	if oldWndPrc != 0 {
		win.SetWindowLongPtr(t.hWnd, win.GWL_WNDPROC, oldWndPrc)
		oldWndPrc = 0
	}
	t.buttons = make([]*ThumbButton, 0)
	t.taskBarList.Release()
	t.taskBarList = 0
}

// 因为lcl不知道什么原因造成WM_COMMAND消息无法收到，所以这样处理
func newWndProc(hWnd uintptr, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case intf.WM_TASKBARBUTTONCREATED:
		if initCreateProc != nil {
			initCreateProc()
		}
	case messages.WM_COMMAND:
		if processMessageProc != nil {
			processMessageProc(wParam)
		}
	}
	return win.CallWindowProc(oldWndPrc, types.HWND(hWnd), message, wParam, lParam)
}
