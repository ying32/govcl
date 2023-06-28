//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows
// +build windows

package win

import (
	"syscall"
	"unsafe"

	. "github.com/ying32/govcl/vcl/types"
)

var (

	// advapi32.dll
	advapi32dll = syscall.NewLazyDLL("advapi32.dll")

	_OpenProcessToken    = advapi32dll.NewProc("OpenProcessToken")
	_GetTokenInformation = advapi32dll.NewProc("GetTokenInformation")

	_RegOpenKeyEx = advapi32dll.NewProc("RegOpenKeyExW")
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

// RegOpenKeyEx
func RegOpenKeyEx(hKey HKEY, lpSubKey string, ulOptions DWORD, samDesired REGSAM, phkResult *HKEY) int32 {
	r, _, _ := _RegOpenKeyEx.Call(uintptr(hKey), CStr(lpSubKey), uintptr(ulOptions), uintptr(samDesired), uintptr(unsafe.Pointer(phkResult)))
	return int32(r)
}
