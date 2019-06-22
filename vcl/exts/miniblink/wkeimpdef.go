// +build windows

package miniblink

import (
	"fmt"
	"unsafe"
)

func InitializeEx(settings *WkeSettings) {
	fmt.Println("_wkeInitializeEx find:", _wkeInitializeEx.Find())
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
