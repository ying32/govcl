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

// 字符串到UTF8指针
func StringToUTF8Ptr(s string) *uint8 {
	temp := []byte(s)
	utf8StrArr := make([]uint8, len(temp)+1) // +1是因为Lazarus中PChar为0结尾
	copy(utf8StrArr, temp)
	return &utf8StrArr[0]
}

// Go的string转换为Lazarus的string
func GoStrToDStr(s string) uintptr {
	if s == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(StringToUTF8Ptr(s)))
}

// 这种跟copyStr3基本一样，只是用go来处理了
func copyStr(src uintptr, strLen int) string {
	if strLen == 0 {
		return ""
	}
	str := make([]uint8, strLen)
	for i := 0; i < strLen; i++ {
		str[i] = *(*uint8)(unsafe.Pointer(src + uintptr(i)))
	}
	return string(str)
}

type GoStringHeader struct {
	Data uintptr
	Len  int
}

// 小点的字符适合此种方式，大了就不行了
func copyStr2(str uintptr, strLen int) string {
	if strLen == 0 {
		return ""
	}
	var ret string
	head := (*GoStringHeader)(unsafe.Pointer(&ret))
	head.Data = str
	head.Len = strLen
	return ret
}

// 最新的lz macOS下出问题了
func copyStr3(str uintptr, strLen int) string {
	if strLen == 0 {
		return ""
	}
	buffer := make([]uint8, strLen)
	DMove(str, uintptr(unsafe.Pointer(&buffer[0])), strLen)
	return string(buffer)
}

// Lazarus的string转换为Go的string
func DStrToGoStr(ustr uintptr) string {
	l := DStrLen(ustr)
	if l == 0 {
		return ""
	}
	return copyStr(ustr, int(l))
}

func getBuff(size int32) interface{} {
	return make([]uint8, size+1)
}

func getBuffPtr(buff interface{}) uintptr {
	return uintptr(unsafe.Pointer(&(buff.([]uint8))[0]))
}

func getTextBuf(strBuff interface{}, Buffer *string, slen int) {
	*Buffer = string((strBuff.([]uint8))[:slen])
}
