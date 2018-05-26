// +build windows

package win

import (
	"syscall"
	"unsafe"
)

type WKSTA_INFO_100 struct {
	Wki100_platform_id  uint32
	Wki100_computername uintptr // LPWSTR
	Wki100_langroup     uintptr // LPWSTR
	Wki100_ver_major    uint32
	Wki100_ver_minor    uint32
}

//LPWKSTA_INFO_100 = ^WKSTA_INFO_100;

const (
	netapi = "netapi32.dll"

	NERR_Success = 0
)

var (
	netapidll           = syscall.NewLazyDLL(netapi)
	_NetWkstaGetInfo100 = netapidll.NewProc("NetWkstaGetInfo")
	_NetApiBufferFree   = netapidll.NewProc("NetApiBufferFree")
)

func NetWkstaGetInfo100(ServerName string, Level uint32, BufPtr **WKSTA_INFO_100) uint32 {
	r, _, _ := _NetWkstaGetInfo100.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(ServerName))), uintptr(Level),
		uintptr(unsafe.Pointer(BufPtr)))
	return uint32(r)
}

func NetApiBufferFree(BufPtr *WKSTA_INFO_100) uint32 {
	r, _, _ := _NetApiBufferFree.Call(uintptr(unsafe.Pointer(BufPtr)))
	return uint32(r)
}
