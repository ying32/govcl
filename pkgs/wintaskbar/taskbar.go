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
	NoProgress    TProcessState = TBPF_NOPROGRESS
	Indeterminate               = TBPF_INDETERMINATE
	Normal                      = TBPF_NORMAL
	Error                       = TBPF_ERROR
	Paused                      = TBPF_PAUSED
)

type TButtonFlags uint32

const (
	Enabled        TButtonFlags = THBF_ENABLED
	Disabled                    = THBF_DISABLED
	DisMissonClick              = THBF_DISMISSONCLICK
	NoBackGround                = THBF_NOBACKGROUND
	Hidden                      = THBF_HIDDEN
	Noninteractive              = THBF_NONINTERACTIVE
)

type TThumbButtonClick func(index uint16)

type ThumbButton struct {
	taskBar *TWinTaskBar
	button  tThumbButton
}

type TWinTaskBar struct {
	hWnd        types.HWND
	taskBarList ITaskbarList4
	created     bool
	buttons     []*ThumbButton
	btnCallback TThumbButtonClick
}

func newThumbButton(taskBar *TWinTaskBar) *ThumbButton {
	tt := new(ThumbButton)
	tt.taskBar = taskBar
	tt.button.dwMask = THB_ICON | THB_TOOLTIP | THB_FLAGS
	tt.button.iId = uint32(len(taskBar.buttons))
	tt.button.dwFlags = uint32(Enabled)
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
	copy(tt.button.szTip[:], s)
	tt.changed()
}

func (tt *ThumbButton) SetIcon(h types.HICON) {
	tt.button.hIcon = h
	tt.changed()
}

func (tt *ThumbButton) SetFlags(flags TButtonFlags) {
	tt.button.dwFlags = uint32(flags)
	tt.changed()
}

//---------------------------------------------------------------------------------------------------------------------

func NewWinTaskBar(hWnd types.HWND) *TWinTaskBar {
	t := new(TWinTaskBar)
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

func (t *TWinTaskBar) SetOnThumbButtonClick(event TThumbButtonClick) {
	t.btnCallback = event
}

func (t *TWinTaskBar) getButtons() []tThumbButton {
	r := make([]tThumbButton, 0)
	for _, b := range t.buttons {
		r = append(r, b.button)
	}
	return r
}

func (t *TWinTaskBar) createTaskBar() error {
	if t.created {
		return nil
	}
	ole.CoInitializeEx(0, 0)
	defer ole.CoUninitialize()

	unknown, err := ole.CreateInstance(CLSID_TaskbarList, SID_ITaskbarList4)
	if err != nil {
		return err
	}
	t.taskBarList = ITaskbarList4(uintptr(unsafe.Pointer(unknown)))
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

func (t *TWinTaskBar) SetThumbnailTooltip(tip string) error {
	return t.taskBarList.SetThumbnailTooltip(t.hWnd, tip)
}

//func (t *TWinTaskBar) ThumbBarUpdateButton(index int, tipStr string, hIconHandle types.HICON, flags TButtonFlags) error {
//	if index >= 0 && index < len(t.buttons) {
//		t.setThumbButton(&t.buttons[index], tipStr, hIconHandle, flags, false)
//		return t.taskBarList.ThumbBarUpdateButtons(t.hWnd, t.buttons)
//	}
//	return errors.New("索引超出范围。")
//}

func (t *TWinTaskBar) SetProgressValue(ullCompleted, ullTotal uint64) error {
	return t.taskBarList.SetProgressValue(t.hWnd, ullCompleted, ullTotal)
}

func (t *TWinTaskBar) SetProgressState(flags TProcessState) error {
	return t.taskBarList.SetProgressState(t.hWnd, int32(flags))
}

func (t *TWinTaskBar) SetOverlayIcon(hIcon types.HICON, pszDescription string) error {
	return t.taskBarList.SetOverlayIcon(t.hWnd, hIcon, pszDescription)
}

func (t *TWinTaskBar) Buttons() []*ThumbButton {
	return t.buttons
}

//func (t *TWinTaskBar) setThumbButton(item *TThumbButton, tipStr string, hIconHandle types.HICON, dwFlags TButtonFlags, IsAdd bool) {
//	if IsAdd {
//		item.dwMask = THB_ICON | THB_TOOLTIP | THB_FLAGS
//		item.iId = uint32(len(t.buttons))
//	}
//	item.hIcon = hIconHandle
//	item.dwFlags = uint32(dwFlags)
//	s := syscall.StringToUTF16(tipStr)
//	copy(item.szTip[:], s)
//}

func (t *TWinTaskBar) AddButton() (*ThumbButton, error) {
	if len(t.buttons) > MAX_BUTTON_COUNT {
		return nil, errors.New(fmt.Sprintf("len(buttons) >  %d", MAX_BUTTON_COUNT))
	}
	if t.created {
		return nil, errors.New("已创建，不能再进行添加操作。")
	}
	return newThumbButton(t), nil
}

func (t *TWinTaskBar) changed() {
	t.taskBarList.ThumbBarUpdateButtons(t.hWnd, t.getButtons())
}

func (t *TWinTaskBar) processMessage(wParam uintptr) {
	// message == messages.WM_COMMAND &&
	if t.btnCallback != nil && hiWord(uint32(wParam)) == THBN_CLICKED {
		t.btnCallback(uint16(uint32(wParam)))
	}
}

func (t *TWinTaskBar) Free() {
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
	case WM_TASKBARBUTTONCREATED:
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
