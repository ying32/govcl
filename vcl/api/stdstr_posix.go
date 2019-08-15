// +build !windows

package api

import (
	"unsafe"
)

func DStrToGoStr(ustr uintptr) string {
	l := DStrLen(ustr)
	if l == 0 {
		return ""
	}
	str := make([]uint8, l)
	DMove(ustr, uintptr(unsafe.Pointer(&str[0])), l)
	return string(str)
}

func GoStrToDStr(s string) uintptr {
	if s == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(&([]byte(s + "\x00")[0])))
}

func getBuff(size int32) interface{}  {
	return make([]uint8, size + 1)
}

func getBuffPtr(buff interface{}) uintptr  {
	return uintptr(unsafe.Pointer(&(buff.([]uint8))[0]))
}

func getTextBuf(strBuff interface{}, Buffer *string) {
	*Buffer = string(strBuff.([]uint8))
}