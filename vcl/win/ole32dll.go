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
	. "github.com/ying32/govcl/vcl/types"
)

import "syscall"

const (
	ole32 = "ole32.dll"
)

var (
	ole32dll = syscall.NewLazyDLL(ole32)

	_CoInitialize   = ole32dll.NewProc("CoInitialize")
	_CoInitializeEx = ole32dll.NewProc("CoInitializeEx")
	_CoUninitialize = ole32dll.NewProc("CoUninitialize")
)

func CoInitialize(pvReserved uintptr) HResult {
	r, _, _ := _CoInitialize.Call(pvReserved)
	return HResult(r)
}

func CoInitializeEx(pvReserved uintptr, coInit uint32) HResult {
	r, _, _ := _CoInitializeEx.Call(pvReserved, uintptr(coInit))
	return HResult(r)
}

func CoUninitialize() {
	_CoUninitialize.Call()
}
