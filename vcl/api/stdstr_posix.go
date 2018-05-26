// +build linux darwin

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
