//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

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
	initLCLStdStr()
	return nil
}

// 字符串到UTF8指针
func StringToUTF8Ptr(s string) *uint8 {
	temp := []byte(s)
	utf8StrArr := make([]uint8, len(temp)+1) // +1是因为Lazarus中PChar为0结尾
	copy(utf8StrArr, temp)
	return &utf8StrArr[0]
}
