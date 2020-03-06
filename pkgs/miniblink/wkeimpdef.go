// +build windows

package miniblink

import (
	"unsafe"
)

func InitializeEx(settings *WkeSettings) {
	if _wkeInitializeEx.Find() != nil {
		return
	}
	_wkeInitializeEx.Call(uintptr(unsafe.Pointer(settings)))
}

func Init() {
	InitializeEx(nil)
}

func Initialize() {
	InitializeEx(nil)
}

func IsInitialize() bool {
	return wkeIsInitialize()
}

func Finalize() {
	wkeFinalize()
}
