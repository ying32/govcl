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
	val1, val2 := UInt64To(dwlConditionMask)
	r, r2, _ := _VerSetConditionMask.Call(val1, val2, uintptr(dwTypeBitMask), uintptr(dwConditionMask))
	return ToUInt64(r, r2)
}

// VerifyVersionInfo
func VerifyVersionInfo(lpVersionInformation *TOSVersionInfoEx, dwTypeMask uint32, dwlConditionMask uint64) bool {
	val1, val2 := UInt64To(dwlConditionMask)
	r, _, _ := _VerifyVersionInfo.Call(uintptr(unsafe.Pointer(lpVersionInformation)), uintptr(dwTypeMask), val1, val2)
	return r != 0
}
