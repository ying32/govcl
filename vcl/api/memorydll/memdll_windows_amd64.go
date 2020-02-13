// +build windows,amd64

package memorydll

import "github.com/ying32/govcl/vcl/api/memorydll/memorymodule"

type moduleHandle = memorymodule.MemoryModule

func memoryLoadLibary(data []byte) moduleHandle {
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
