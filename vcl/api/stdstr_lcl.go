//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"unsafe"
)

// ---------------- liblcl ----------------
func initLCLStdStr() {
	//fmt.Println("Library Character Encoding: UTF-8")
	GoStrToDStr = lcl_GoStrToDStr
	DStrToGoStr = lcl_DStrToGoStr
	getBuff = lcl_getBuff
	getBuffPtr = lcl_getBuffPtr
	getTextBuf = lcl_getTextBuf
}

// Go的string转换为Lazarus的string
func lcl_GoStrToDStr(s string) uintptr {
	if s == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(StringToUTF8Ptr(s)))
}

// Lazarus的string转换为Go的string
func lcl_DStrToGoStr(ustr uintptr) string {
	l := DStrLen(ustr)
	if l == 0 {
		return ""
	}
	str := make([]uint8, l)
	DMove(ustr, uintptr(unsafe.Pointer(&str[0])), l)
	return string(str)
}

func lcl_getBuff(size int32) interface{} {
	return make([]uint8, size+1)
}

func lcl_getBuffPtr(buff interface{}) uintptr {
	return uintptr(unsafe.Pointer(&(buff.([]uint8))[0]))
}

func lcl_getTextBuf(strBuff interface{}, Buffer *string, slen int) {
	*Buffer = string((strBuff.([]uint8))[:slen])
}
