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
	"syscall"
	"unsafe"
)

var (

	// advapi32.dll
	advapi32dll = syscall.NewLazyDLL("advapi32.dll")

	_OpenProcessToken    = advapi32dll.NewProc("OpenProcessToken")
	_GetTokenInformation = advapi32dll.NewProc("GetTokenInformation")
)

// OpenProcessToken
func OpenProcessToken(ProcessHandle uintptr, DesiredAccess uint32, TokenHandle *uintptr) bool {
	r, _, _ := _OpenProcessToken.Call(ProcessHandle, uintptr(DesiredAccess), uintptr(unsafe.Pointer(TokenHandle)))
	return r != 0
}

// GetTokenInformation
func GetTokenInformation(TokenHandle uintptr, TokenInformationClass TTokenInformationClass, TokenInformation uintptr, TokenInformationLength uint32,
	ReturnLength *uint32) bool {
	r, _, _ := _GetTokenInformation.Call(TokenHandle, uintptr(TokenInformationClass), TokenInformation, uintptr(TokenInformationLength), uintptr(unsafe.Pointer(ReturnLength)))
	return r != 0
}
