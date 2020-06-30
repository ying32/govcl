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
	"errors"
	"syscall"
	"unsafe"

	"github.com/ying32/govcl/vcl/win"

	"github.com/ying32/govcl/vcl/types"

	"github.com/go-ole/go-ole"
)

var (
	oleaut32dll   = syscall.NewLazyDLL("oleaut32.dll")
	_DispCallFunc = oleaut32dll.NewProc("DispCallFunc")

	// 虚方法表的，因为Taskbar的没有实现IDispatch接口，所以只能用这种绝对的方式做
	virtualMethods = map[string]uintptr{

		"Release": 3,

		// function HrInit: HRESULT; stdcall;
		"HrInit": 4,

		// function SetProgressValue(hwnd: HWND; ullCompleted: ULONGLONG; ullTotal: ULONGLONG): HRESULT; stdcall;
		"SetProgressValue": 10,

		//function SetProgressState(hwnd: HWND; tbpFlags: Integer): HRESULT; stdcall;
		"SetProgressState": 11,

		// function ThumbBarAddButtons(hwnd: HWND; cButtons: UINT; pButton: PThumbButton): HRESULT; stdcall;
		"ThumbBarAddButtons": 16,

		// function ThumbBarUpdateButtons(hwnd: HWND; cButtons: UINT; pButton: PThumbButton): HRESULT; stdcall;
		"ThumbBarUpdateButtons": 17,

		// function SetOverlayIcon(hwnd: HWND; hIcon: HICON; pszDescription: LPCWSTR): HRESULT; stdcall;
		"SetOverlayIcon": 19,

		// function SetThumbnailTooltip(hwnd: HWND; pszTip: LPCWSTR): HRESULT; stdcall;
		"SetThumbnailTooltip": 20}
)

func getOffset(name string) uintptr {
	if v, ok := virtualMethods[name]; ok {
		return v
	}
	return 0
}

func Succeeded(Status types.HResult) bool {
	return Status&types.HResult(0x80000000) == 0
}

type ITaskbarList4 uintptr

func (ii ITaskbarList4) Release() error {
	if ii > 0 {
		x := (*ole.IUnknown)(unsafe.Pointer(ii))
		x.Release()
	}
	//var result ole.VARIANT
	//err := DispCallFunc(uintptr(unsafe.Pointer(ii)), getOffset("Release"), ole.CC_STDCALL, ole.VT_EMPTY, nil, nil, &result)
	//if err != nil {
	//	return err
	//}
	//if !Succeeded(types.HResult(result.Val)) {
	//	return ole.NewError(uintptr(result.Val))
	//}
	return nil
}

func (ii ITaskbarList4) HrInit() error {
	var result ole.VARIANT
	err := DispCallFunc(uintptr(ii), getOffset("HrInit"), ole.CC_STDCALL, ole.VT_EMPTY, nil, nil, &result)
	if err != nil {
		return err
	}
	if !Succeeded(types.HResult(result.Val)) {
		return ole.NewError(uintptr(result.Val))
	}
	return nil
}

func (ii ITaskbarList4) addOrUpdateButtons(offset uintptr, hWnd types.HWND, buttons []TThumbButton) error {
	if len(buttons) == 0 {
		return errors.New("buttons == 0")
	}
	var result ole.VARIANT

	argsType := []ole.VT{
		ole.VT_UI4,
		ole.VT_UI4,
		ole.VT_PTR,
	}

	arg1 := ole.NewVariant(ole.VT_UI4, int64(hWnd))
	arg2 := ole.NewVariant(ole.VT_UI4, int64(len(buttons)))
	arg3 := ole.NewVariant(ole.VT_UI4, int64(uintptr(unsafe.Pointer(&buttons[0]))))
	args := []*ole.VARIANT{
		&arg1,
		&arg2,
		&arg3,
	}
	err := DispCallFunc(uintptr(ii), offset, ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
	if err != nil {
		return err
	}
	if !Succeeded(types.HResult(result.Val)) {
		return ole.NewError(uintptr(result.Val))
	}
	return nil
}

func (ii ITaskbarList4) ThumbBarAddButtons(hWnd types.HWND, buttons []TThumbButton) error {
	return ii.addOrUpdateButtons(getOffset("ThumbBarAddButtons"), hWnd, buttons)
}

func (ii ITaskbarList4) ThumbBarUpdateButtons(hWnd types.HWND, buttons []TThumbButton) error {
	return ii.addOrUpdateButtons(getOffset("ThumbBarUpdateButtons"), hWnd, buttons)
}

func (ii ITaskbarList4) SetThumbnailTooltip(hWnd types.HWND, tip string) error {
	var result ole.VARIANT

	argsType := []ole.VT{
		ole.VT_UI4,
		ole.VT_LPWSTR,
	}
	arg1 := ole.NewVariant(ole.VT_UI4, int64(hWnd))
	arg2 := ole.NewVariant(ole.VT_LPWSTR, int64(win.CStr(tip)))
	args := []*ole.VARIANT{
		&arg1,
		&arg2,
	}
	err := DispCallFunc(uintptr(ii), getOffset("SetThumbnailTooltip"), ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
	if err != nil {
		return err
	}
	if !Succeeded(types.HResult(result.Val)) {
		return ole.NewError(uintptr(result.Val))
	}
	return nil
}

func (ii ITaskbarList4) SetProgressValue(hWnd types.HWND, ullCompleted, ullTotal uint64) error {
	var result ole.VARIANT

	argsType := []ole.VT{
		ole.VT_UI4,
		ole.VT_UI8,
		ole.VT_UI8,
	}
	arg1 := ole.NewVariant(ole.VT_UI4, int64(hWnd))
	arg2 := ole.NewVariant(ole.VT_UI8, int64(ullCompleted))
	arg3 := ole.NewVariant(ole.VT_UI8, int64(ullTotal))
	args := []*ole.VARIANT{
		&arg1,
		&arg2,
		&arg3,
	}
	err := DispCallFunc(uintptr(ii), getOffset("SetProgressValue"), ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
	if err != nil {
		return err
	}
	if !Succeeded(types.HResult(result.Val)) {
		return ole.NewError(uintptr(result.Val))
	}
	return nil
}

func (ii ITaskbarList4) SetProgressState(hWnd types.HWND, flags int32) error {
	var result ole.VARIANT
	argsType := []ole.VT{
		ole.VT_UI4,
		ole.VT_I4,
	}
	arg1 := ole.NewVariant(ole.VT_UI4, int64(hWnd))
	arg2 := ole.NewVariant(ole.VT_I4, int64(flags))
	args := []*ole.VARIANT{
		&arg1,
		&arg2,
	}
	err := DispCallFunc(uintptr(ii), getOffset("SetProgressState"), ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
	if err != nil {
		return err
	}
	if !Succeeded(types.HResult(result.Val)) {
		return ole.NewError(uintptr(result.Val))
	}
	return nil
}

func (ii ITaskbarList4) SetOverlayIcon(hWnd types.HWND, hIcon types.HICON, pszDescription string) error {
	var result ole.VARIANT

	argsType := []ole.VT{
		ole.VT_UI4,
		ole.VT_UI4,
		ole.VT_LPWSTR,
	}
	arg1 := ole.NewVariant(ole.VT_UI4, int64(hWnd))
	arg2 := ole.NewVariant(ole.VT_UI4, int64(hIcon))
	arg3 := ole.NewVariant(ole.VT_LPWSTR, int64(win.CStr(pszDescription)))
	args := []*ole.VARIANT{
		&arg1,
		&arg2,
		&arg3,
	}
	err := DispCallFunc(uintptr(ii), 11, ole.CC_STDCALL, ole.VT_EMPTY, argsType, args, &result)
	if err != nil {
		return err
	}
	if !Succeeded(types.HResult(result.Val)) {
		return ole.NewError(uintptr(result.Val))
	}
	return nil
}

// -------------- call-------------------------
func DispCallFunc(pvInstance uintptr, oVft uintptr, cc int32, vtReturn ole.VT, argTypes []ole.VT, args []*ole.VARIANT, result *ole.VARIANT) error {
	if pvInstance == 0 {
		return errors.New("pvInstance == 0")
	}
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
