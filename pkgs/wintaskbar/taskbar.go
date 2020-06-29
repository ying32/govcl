//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
//+build windows

/*
Minimum supported client	Windows 7 [desktop apps only]
Minimum supported server	Windows Server 2008 R2 [desktop apps only]
*/

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

	oleaut32dll   = syscall.NewLazyDLL("oleaut32.dll")
	_DispCallFunc = oleaut32dll.NewProc("DispCallFunc")

	oldWndPrc          uintptr
	initCreateProc     func() error
	processMessageProc func(uintptr)
)

type TThumbButtonClick func(index uint16)

type TThumbButton struct {
	dwMask  uint32
	iId     uint32
	iBitmap uint32
	hIcon   types.HICON
	szTip   [260]uint16
	dwFlags uint32
}

type TWinTaskBar struct {
	hWnd        types.HWND
	taskBarList *ole.IUnknown
	created     bool
	buttons     []TThumbButton
	btnCallback TThumbButtonClick
}

func NewWinTaskBar(hWnd types.HWND) *TWinTaskBar {
	t := new(TWinTaskBar)
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

func (t *TWinTaskBar) initButtons() error {
	if t.taskBarList != nil && len(t.buttons) > 0 {
		var result ole.VARIANT

		argsType := []ole.VT{
			ole.VT_UI4,
			ole.VT_UI4,
			ole.VT_PTR,
		}

		arg1 := ole.NewVariant(ole.VT_UI4, int64(t.hWnd))
		arg2 := ole.NewVariant(ole.VT_UI4, int64(len(t.buttons)))
		arg3 := ole.NewVariant(ole.VT_UI4, int64(uintptr(unsafe.Pointer(&t.buttons[0]))))
		args := []*ole.VARIANT{
			&arg1,
			&arg2,
			&arg3,
		}
		// call: function ThumbBarAddButtons(hwnd: HWND; cButtons: UINT; pButton: PThumbButton): HRESULT; stdcall;
		err := DispCallFunc(uintptr(unsafe.Pointer(t.taskBarList)), 16, ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
		if err != nil {
			return err
		}
		if !Succeeded(types.HResult(result.Val)) {
			return ole.NewError(uintptr(result.Val))
		}
	}
	return nil
}

func (t *TWinTaskBar) createTaskBar() error {
	if t.created {
		return nil
	}
	ole.CoInitializeEx(0, 0)
	defer ole.CoUninitialize()
	var err error
	t.taskBarList, err = ole.CreateInstance(CLSID_TaskbarList, SID_ITaskbarList4)

	if err != nil {
		return err
	}
	var result ole.VARIANT
	// function HrInit: HRESULT; stdcall;
	err = DispCallFunc(uintptr(unsafe.Pointer(t.taskBarList)), 4, ole.CC_STDCALL, ole.VT_EMPTY, nil, nil, &result)
	if err != nil {
		return err
	}
	if !Succeeded(types.HResult(result.Val)) {
		return ole.NewError(uintptr(result.Val))
	}
	err = t.initButtons()
	if err != nil {
		return err
	}
	t.created = true

	return nil
}

func (t *TWinTaskBar) SetThumbnailTooltip(tip string) error {
	if t.taskBarList != nil {
		var result ole.VARIANT

		argsType := []ole.VT{
			ole.VT_UI4,
			ole.VT_LPWSTR,
		}
		arg1 := ole.NewVariant(ole.VT_UI4, int64(t.hWnd))
		arg2 := ole.NewVariant(ole.VT_LPWSTR, int64(win.CStr(tip)))
		args := []*ole.VARIANT{
			&arg1,
			&arg2,
		}
		// call: function SetThumbnailTooltip(hwnd: HWND; pszTip: LPCWSTR): HRESULT; stdcall;
		err := DispCallFunc(uintptr(unsafe.Pointer(t.taskBarList)), 20, ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
		if err != nil {
			return err
		}
		if !Succeeded(types.HResult(result.Val)) {
			return ole.NewError(uintptr(result.Val))
		}
	}
	return nil
}

func (t *TWinTaskBar) ThumbBarUpdateButtons() error {
	if t.taskBarList != nil && len(t.buttons) > 0 {
		var result ole.VARIANT

		argsType := []ole.VT{
			ole.VT_UI4,
			ole.VT_UI4,
			ole.VT_PTR,
		}

		arg1 := ole.NewVariant(ole.VT_UI4, int64(t.hWnd))
		arg2 := ole.NewVariant(ole.VT_UI4, int64(len(t.buttons)))
		arg3 := ole.NewVariant(ole.VT_UI4, int64(uintptr(unsafe.Pointer(&t.buttons[0]))))
		args := []*ole.VARIANT{
			&arg1,
			&arg2,
			&arg3,
		}
		// offset: 17
		//function ThumbBarUpdateButtons(hwnd: HWND; cButtons: UINT; pButton: PThumbButton): HRESULT; stdcall;
		err := DispCallFunc(uintptr(unsafe.Pointer(t.taskBarList)), 17, ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
		if err != nil {
			return err
		}
		if !Succeeded(types.HResult(result.Val)) {
			return ole.NewError(uintptr(result.Val))
		}
	}
	return nil
}

func (t *TWinTaskBar) SetProgressValue(ullCompleted, ullTotal uint64) error {
	if t.taskBarList != nil {
		var result ole.VARIANT

		argsType := []ole.VT{
			ole.VT_UI4,
			ole.VT_UI8,
			ole.VT_UI8,
		}
		arg1 := ole.NewVariant(ole.VT_UI4, int64(t.hWnd))
		arg2 := ole.NewVariant(ole.VT_UI8, int64(ullCompleted))
		arg3 := ole.NewVariant(ole.VT_UI8, int64(ullTotal))
		args := []*ole.VARIANT{
			&arg1,
			&arg2,
			&arg3,
		}
		// offset: 10
		//function SetProgressValue(hwnd: HWND; ullCompleted: ULONGLONG; ullTotal: ULONGLONG): HRESULT; stdcall;
		err := DispCallFunc(uintptr(unsafe.Pointer(t.taskBarList)), 10, ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
		if err != nil {
			return err
		}
		if !Succeeded(types.HResult(result.Val)) {
			return ole.NewError(uintptr(result.Val))
		}
	}
	return nil
}

func (t *TWinTaskBar) SetProgressState(flags int32) error {
	if t.taskBarList != nil {
		var result ole.VARIANT

		argsType := []ole.VT{
			ole.VT_UI4,
			ole.VT_I4,
		}
		arg1 := ole.NewVariant(ole.VT_UI4, int64(t.hWnd))
		arg2 := ole.NewVariant(ole.VT_I4, int64(flags))
		args := []*ole.VARIANT{
			&arg1,
			&arg2,
		}
		// offset: 11
		//function SetProgressState(hwnd: HWND; tbpFlags: Integer): HRESULT; stdcall;
		err := DispCallFunc(uintptr(unsafe.Pointer(t.taskBarList)), 11, ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
		if err != nil {
			return err
		}
		if !Succeeded(types.HResult(result.Val)) {
			return ole.NewError(uintptr(result.Val))
		}
	}
	return nil
}

func (t *TWinTaskBar) SetOverlayIcon(hIcon types.HICON, pszDescription string) error {
	if t.taskBarList != nil {
		var result ole.VARIANT

		argsType := []ole.VT{
			ole.VT_UI4,
			ole.VT_UI4,
			ole.VT_LPWSTR,
		}
		arg1 := ole.NewVariant(ole.VT_UI4, int64(t.hWnd))
		arg2 := ole.NewVariant(ole.VT_UI4, int64(hIcon))
		arg3 := ole.NewVariant(ole.VT_LPWSTR, int64(win.CStr(pszDescription)))
		args := []*ole.VARIANT{
			&arg1,
			&arg2,
			&arg3,
		}
		// offset: 19
		// function SetOverlayIcon(hwnd: HWND; hIcon: HICON; pszDescription: LPCWSTR): HRESULT; stdcall;
		err := DispCallFunc(uintptr(unsafe.Pointer(t.taskBarList)), 11, ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
		if err != nil {
			return err
		}
		if !Succeeded(types.HResult(result.Val)) {
			return ole.NewError(uintptr(result.Val))
		}
	}
	return nil
}

func Succeeded(Status types.HResult) bool {
	return Status&types.HResult(0x80000000) == 0
}

func (t *TWinTaskBar) AddButton(tipStr string, hIconHandle types.HICON, flags uint32) error {
	if len(t.buttons) > MAX_BUTTON_COUNT {
		return errors.New(fmt.Sprintf("len(buttons) >  %d", MAX_BUTTON_COUNT))
	}
	item := TThumbButton{}
	item.dwMask = THB_ICON | THB_TOOLTIP | THB_FLAGS
	item.iId = uint32(len(t.buttons))
	item.hIcon = hIconHandle
	item.dwFlags = flags
	s := syscall.StringToUTF16(tipStr)
	copy(item.szTip[:], s)
	t.buttons = append(t.buttons, item)
	return nil
}

func (t *TWinTaskBar) changed() {
	//t.taskBarList.CallMethod("ThumbBarUpdateButtons", t.hWnd,  AIndex, @LItem);
}

func hiWord(L uint32) uint16 {
	return uint16(L >> 16)
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
	t.taskBarList = nil
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

func DispCallFunc(pvInstance uintptr, oVft uintptr, cc int32, vtReturn ole.VT, argTypes []ole.VT, args []*ole.VARIANT, result *ole.VARIANT) error {
	oVft = unsafe.Sizeof(pvInstance) * (oVft - 1)
	argLen := int32(len(argTypes))

	Argsptr := uintptr(0)
	if len(args) > 0 {
		Argsptr = uintptr(unsafe.Pointer(&args[0]))
	}
	argTypePtr := uintptr(0)
	if len(argTypes) > 0 {
		argTypePtr = uintptr(unsafe.Pointer(&argTypes[0]))
	}
	r, _, _ := _DispCallFunc.Call(pvInstance, oVft, uintptr(cc), uintptr(vtReturn), uintptr(argLen), argTypePtr, Argsptr, uintptr(unsafe.Pointer(result)))
	if !Succeeded(r) {
		return ole.NewError(r)
	}
	return nil
}

func init() {
	WM_TASKBARBUTTONCREATED = win.RegisterWindowMessage("TaskbarButtonCreated")
}
