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

// 这种跟copyStr3基本一样，只是用go来处理了
func copyStr(src uintptr, strlen int) string {
	str := make([]uint8, strlen)
	for i := 0; i < strlen; i++ {
		str[i] = *(*uint8)(unsafe.Pointer(src + uintptr(i)))
	}
	return string(str)
}

type GoStringHeader struct {
	Data uintptr
	Len  int
}

// 小点的字符适合此种方式，大了就不行了
func copyStr2(str uintptr, strlen int) string {
	var ret string
	head := (*GoStringHeader)(unsafe.Pointer(&ret))
	head.Data = str
	head.Len = strlen
	return ret
}

// 最新的lz macOS下出问题了
func copyStr3(str uintptr, strlen int) string {
	buffer := make([]uint8, strlen)
	DMove(str, uintptr(unsafe.Pointer(&buffer[0])), strlen)
	return string(buffer)
}

// Lazarus的string转换为Go的string
func lcl_DStrToGoStr(ustr uintptr) string {
	l := DStrLen(ustr)
	if l == 0 {
		return ""
	}
	return copyStr2(ustr, int(l))
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
