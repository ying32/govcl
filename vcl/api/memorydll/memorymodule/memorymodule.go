//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows

/*

  作者：ying32，本为govcl项目中的一个dylib扩展包
  https://github.com/ying32/govcl

  直接使用cgo方式调用，因为不想再将C++的代码完全翻译成Go的了，太累了，也觉得没必要。
  https://github.com/fancycode/MemoryModule

*/

package memorymodule

/*
  #include "MemoryModule.h"
*/
import "C"
import "unsafe"

type MemoryModule C.HMEMORYMODULE

func MemoryLoadLibrary(data []byte) MemoryModule {
	return MemoryModule(C.MemoryLoadLibrary(unsafe.Pointer(&data[0]), C.size_t(len(data))))
}

func MemoryGetProcAddress(handle MemoryModule, name string) uintptr {
	if handle == nil {
		return 0
	}
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return uintptr(unsafe.Pointer(C.MemoryGetProcAddress(C.HMEMORYMODULE(handle), (*C.CHAR)(cname))))
}

func MemoryFreeLibrary(handle MemoryModule) {
	if handle != nil {
		C.MemoryFreeLibrary(C.HMEMORYMODULE(handle))
	}
}
