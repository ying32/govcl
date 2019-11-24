// +build windows

// miniblink及wke头文件导入
// 由ying32翻译，应用于govcl

package miniblink

import (
	"syscall"
	"unsafe"

	"github.com/ying32/govcl/vcl/win"
)

func GoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

func CBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
}

func GoWStr(str uintptr) string {
	if str == 0 {
		return ""
	}
	l := win.LstrlenW(str)
	if l == 0 {
		return ""
	}
	buff := make([]uint16, l)
	win.Memcpy(uintptr(unsafe.Pointer(&buff[0])), str, uintptr(l*2))
	return syscall.UTF16ToString(buff)
}

func GoAStr(str uintptr) string {
	if str == 0 {
		return ""
	}
	l := win.Lstrlen(str)
	if l == 0 {
		return ""
	}
	buff := make([]byte, l)
	win.Memcpy(uintptr(unsafe.Pointer(&buff[0])), str, uintptr(l))
	return string(buff)
}

func CWStr(str string) uintptr {
	if str == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}

func CAStr(str string) uintptr {
	if str == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(&([]byte(str + "\x00"))[0]))
}

// 下面两个函数主要是给 386平台使用的，amd64不需要使用
func ToUInt64(r1, r2 uintptr) uint64 {
	ret := uint64(r2)
	ret = uint64(ret<<32) + uint64(r1)
	return ret
}

func UInt64To(val uint64) (uintptr, uintptr) {
	return uintptr(uint32(val)), uintptr(uint32(val >> 32))
}
