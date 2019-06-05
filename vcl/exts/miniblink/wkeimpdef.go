// +build windows

package miniblink

import "unsafe"

var (
	_wkeInitializeEx = wkedll.NewProc("wkeInitializeEx")
)

func InitializeEx(settings *WkeSettings) int {
	r, _, _ := _wkeInitializeEx.Call(uintptr(unsafe.Pointer(settings)))
	return int(r)
}

func Init() int {
	return InitializeEx(nil)
}

func Initialize() int {
	return InitializeEx(nil)
}

func IsInitialize() bool {
	return wkeIsInitialize()
}

func Finalize() {
	wkeFinalize()
}
