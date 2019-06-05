// +build windows

package miniblink

import (
	"syscall"
	"unsafe"

	"github.com/ying32/govcl/vcl/win"
)

var (
	kWkeDllPath string = "node.dll"
)

func WkeSetWkeDllPath(dllPath string) {
	kWkeDllPath = dllPath
}

func WkeInitializeEx(settings *WkeSettings) int {
	hMod := win.LoadLibrary(kWkeDllPath)
	if hMod != 0 {
		proc := win.GetProcAddress(hMod, "wkeInitializeEx")
		if proc != 0 {
			syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(settings)), 0, 0)
			return 1
		}
	}
	return 0
}

func WkeInit() int {
	return WkeInitializeEx(nil)
}

func WkeInitialize() int {
	return WkeInitializeEx(nil)
}
