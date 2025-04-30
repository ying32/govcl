//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows

package win

import "unsafe"

// VerSetConditionMask
func VerSetConditionMask(dwlConditionMask uint64, dwTypeBitMask uint32, dwConditionMask uint8) uint64 {
	r, _, _ := _VerSetConditionMask.Call(uintptr(dwlConditionMask), uintptr(dwTypeBitMask), uintptr(dwConditionMask))
	return uint64(r)
}

// VerifyVersionInfo
func VerifyVersionInfo(lpVersionInformation *TOSVersionInfoEx, dwTypeMask uint32, dwlConditionMask uint64) bool {
	r, _, _ := _VerifyVersionInfo.Call(uintptr(unsafe.Pointer(lpVersionInformation)), uintptr(dwTypeMask), uintptr(dwlConditionMask))
	return r != 0
}
