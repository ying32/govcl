// +build windows

package miniblink

import (
	"fmt"
	"unsafe"
)

var (
	_wkeInitializeEx = wkedll.NewProc("wkeInitializeEx")
)

func InitializeEx(settings *WkeSettings) {
	fmt.Println(unsafe.Sizeof(settings))
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
