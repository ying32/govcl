package api

import (
	"syscall"
	"unsafe"
)

// ----------------------------------- 共用的
func DStrToGoStr(ustr uintptr) string {
	l := DStrLen(ustr)
	if l == 0 {
		return ""
	}
	if IsloadedLcl {
		str := make([]uint8, l)
		DMove(ustr, uintptr(unsafe.Pointer(&str[0])), l)
		return string(str)
	} else {
		str := make([]uint16, l)
		DMove(ustr, uintptr(unsafe.Pointer(&str[0])), l*2)
		return syscall.UTF16ToString(str)
	}
}

func GoStrToDStr(s string) uintptr {
	if s == "" {
		return 0
	}
	if IsloadedLcl {
		return uintptr(unsafe.Pointer(&([]byte(s + "\x00")[0])))
	} else {
		return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
	}
}

func getBuff(size int32) interface{}  {
	if IsloadedLcl {
		return make([]uint8, size + 1)
	} else {
		return make([]uint16, size + 1)
	}
}

func getBuffPtr(buff interface{}) uintptr  {
	if IsloadedLcl {
		return uintptr(unsafe.Pointer(&(buff.([]uint8))[0]))
	} else {
		return uintptr(unsafe.Pointer(&(buff.([]uint16))[0]))
	}
}

func getTextBuf(strBuff interface{}, Buffer *string) {
	if IsloadedLcl {
		*Buffer = string(strBuff.([]uint8))
	} else {
		*Buffer = syscall.UTF16ToString(strBuff.([]uint16))
	}
}

