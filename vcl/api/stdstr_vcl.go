//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows

package api

import (
	"fmt"
	"syscall"
	"unsafe"
)

// ---------------- libvcl ----------------

func initVCLStdStr() {
	fmt.Println("Library Character Encoding: Unicode")
	GoStrToDStr = vcl_GoStrToDStr
	DStrToGoStr = vcl_DStrToGoStr
	getBuff = vcl_getBuff
	getBuffPtr = vcl_getBuffPtr
	getTextBuf = vcl_getTextBuf
}

// Go的string转换为Delphi的string
func vcl_GoStrToDStr(s string) uintptr {
	if s == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

// Delphi的string转换为Go的string
func vcl_DStrToGoStr(ustr uintptr) string {
	l := DStrLen(ustr)
	if l == 0 {
		return ""
	}
	str := make([]uint16, l)
	DMove(ustr, uintptr(unsafe.Pointer(&str[0])), l*2)
	return syscall.UTF16ToString(str)
}

func vcl_getBuff(size int32) interface{} {
	return make([]uint16, size+1)
}

func vcl_getBuffPtr(buff interface{}) uintptr {
	return uintptr(unsafe.Pointer(&(buff.([]uint16))[0]))
}

func vcl_getTextBuf(strBuff interface{}, Buffer *string, slen int) {
	*Buffer = syscall.UTF16ToString((strBuff.([]uint16))[:slen])
}
