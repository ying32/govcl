// +build windows,amd64

//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package memorydll

import "github.com/ying32/govcl/vcl/api/memorydll/memorymodule"

type moduleHandle = memorymodule.MemoryModule

func memoryLoadLibrary(data []byte) moduleHandle {
	return memorymodule.MemoryLoadLibrary(data)
}

func memoryGetLastError() string {
	return "Load ERROR."
}

func memoryFreeLibrary(handle moduleHandle) {
	memorymodule.MemoryFreeLibrary(handle)
}

func memoryGetProcAddress(handle moduleHandle, name string) uintptr {
	return memorymodule.MemoryGetProcAddress(handle, name)
}

func handleIsValid(handle moduleHandle) bool {
	return handle != nil
}
