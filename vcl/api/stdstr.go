package api

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	// 一些字符串操作函数指针，这样做是因为初始的时候就确认调用VCL或者LCL库的转换函数。
	GoStrToDStr func(string) uintptr
	DStrToGoStr func(uintptr) string
	getBuff     func(int32) interface{}
	getBuffPtr  func(interface{}) uintptr
	getTextBuf  func(interface{}, *string, int)
)

// 初始系统字符串转换相关
func initStdString() error {
	if IsloadedLcl {
		fmt.Println("Library Character Encoding: UTF-8")
		GoStrToDStr = lcl_GoStrToDStr
		DStrToGoStr = lcl_DStrToGoStr
		getBuff = lcl_getBuff
		getBuffPtr = lcl_getBuffPtr
		getTextBuf = lcl_getTextBuf
	} else {
		fmt.Println("Library Character Encoding: Unicode")
		GoStrToDStr = vcl_GoStrToDStr
		DStrToGoStr = vcl_DStrToGoStr
		getBuff = vcl_getBuff
		getBuffPtr = vcl_getBuffPtr
		getTextBuf = vcl_getTextBuf
	}
	return nil
}

// 字符串到UTF8指针
func StringToUTF8Ptr(s string) *uint8 {
	temp := []byte(s)
	utf8StrArr := make([]uint8, len(temp)+1) // +1是因为Lazarus中PChar为0结尾
	copy(utf8StrArr, temp)
	return &utf8StrArr[0]
}

// ---------------- libvcl ----------------

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

// ---------------- liblcl ----------------
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
