// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (

	// kernel32.dll
	kernel32dll = syscall.NewLazyDLL("kernel32.dll")

	_GetLastError = kernel32dll.NewProc("GetLastError")
	_SetLastError = kernel32dll.NewProc("SetLastError")

	_IsWow64Process      = kernel32dll.NewProc("IsWow64Process")
	_GetCurrentProcessId = kernel32dll.NewProc("GetCurrentProcessId")
	_GetCurrentProcess   = kernel32dll.NewProc("GetCurrentProcess")
	_GetCurrentThreadId  = kernel32dll.NewProc("GetCurrentThreadId")
	_GetCurrentThread    = kernel32dll.NewProc("GetCurrentThread")

	_GetModuleHandle     = kernel32dll.NewProc("GetModuleHandleW")
	_GetVersionEx        = kernel32dll.NewProc("GetVersionExW")
	_GetNativeSystemInfo = kernel32dll.NewProc("GetNativeSystemInfo")
	_VerSetConditionMask = kernel32dll.NewProc("VerSetConditionMask")
	_VerifyVersionInfo   = kernel32dll.NewProc("VerifyVersionInfoW")

	_CloseHandle = kernel32dll.NewProc("CloseHandle")

	_OpenProcess = kernel32dll.NewProc("OpenProcess")

	_CreateMutex  = kernel32dll.NewProc("CreateMutexW")
	_ReleaseMutex = kernel32dll.NewProc("ReleaseMutex")
	_OpenMutex    = kernel32dll.NewProc("OpenMutexW")

	_MultiByteToWideChar = kernel32dll.NewProc("MultiByteToWideChar")
	_WideCharToMultiByte = kernel32dll.NewProc("WideCharToMultiByte")
)

// GetLastError
func GetLastError() uint32 {
	r, _, _ := _GetLastError.Call()
	return uint32(r)
}

// SetLastError
func SetLastError(dwErrCode uint32) {
	_SetLastError.Call(uintptr(dwErrCode))
}

// GetModuleHandle 获取当前是实例句柄，可传空
func GetModuleHandle(lpModuleName string) uintptr {
	r, _, _ := _GetModuleHandle.Call(CStr(lpModuleName))
	return r
}

// GetSelfModuleHandle 获取自身模块实例句柄
func GetSelfModuleHandle() uintptr {
	r, _, _ := _GetModuleHandle.Call(0)
	return r
}

// IsWow64Process 检测进程是否运行在64位下
func IsWow64Process(hProcess uintptr) bool {
	if _IsWow64Process.Find() != nil {
		return false
	}
	var wow64Process int32
	r, _, _ := _IsWow64Process.Call(hProcess, uintptr(unsafe.Pointer(&wow64Process)))
	if r != 0 && wow64Process != 0 {
		return true
	}
	return false
}

// GetCurrentProcess 返回当前进程伪句柄
func GetCurrentProcess() uintptr {
	r, _, _ := _GetCurrentProcess.Call()
	return r
}

// GetCurrentProcessId 返回当前进程伪句柄
func GetCurrentProcessId() uintptr {
	r, _, _ := _GetCurrentProcessId.Call()
	return r
}

// GetCurrentThreadId
func GetCurrentThreadId() uintptr {
	r, _, _ := _GetCurrentThreadId.Call()
	return r
}

// GetCurrentThread
func GetCurrentThread() uintptr {
	r, _, _ := _GetCurrentThread.Call()
	return r
}

// IsWow64 判断当前进程是否运行在64上
// 注：只有当exe为Win32并在64位系统上运行时才返回true, 否则都会返回false
func IsWow64() bool {
	return IsWow64Process(GetCurrentProcess())
}

// 获取系统版本信息
func GetVersionEx(lpVersionInformation *TOSVersionInfoEx) bool {
	if lpVersionInformation != nil {
		lpVersionInformation.OSVersionInfoSize = uint32(unsafe.Sizeof(*lpVersionInformation))
	}
	r, _, _ := _GetVersionEx.Call(uintptr(unsafe.Pointer(lpVersionInformation)))
	return r != 0
}

// GetNativeSystemInfo
func GetNativeSystemInfo(lpSystemInformation *TSystemInfo) bool {
	r, _, _ := _GetNativeSystemInfo.Call(uintptr(unsafe.Pointer(lpSystemInformation)))
	return r != 0
}

// CloseHandle
func CloseHandle(hObject uintptr) bool {
	r, _, _ := _CloseHandle.Call(hObject)
	return r != 0
}

// OpenProcess
func OpenProcess(dwDesiredAccess uint32, bInheritHandle bool, dwProcessId uint32) uintptr {
	r, _, _ := _OpenProcess.Call(uintptr(dwDesiredAccess), uintptr(CBool(bInheritHandle)), uintptr(dwProcessId))
	return r
}

// CreateMutex
func CreateMutex(lpMutexAttributes *TSecurityAttributes, bInitialOwner bool, lpName string) uintptr {
	r, _, _ := _CreateMutex.Call(uintptr(unsafe.Pointer(lpMutexAttributes)), CBool(bInitialOwner), CStr(lpName))
	return r
}

// ReleaseMutex
func ReleaseMutex(hMutex uintptr) bool {
	r, _, _ := _ReleaseMutex.Call(hMutex)
	return r != 0
}

// OpenMutex
func OpenMutex(dwDesiredAccess uint32, bInheritHandle bool, lpName string) uintptr {
	r, _, _ := _OpenMutex.Call(uintptr(dwDesiredAccess), CBool(bInheritHandle), CStr(lpName))
	return r
}

// MultiByteToWideChar
func MultiByteToWideChar(CodePage, dwFlags uint32, lpMultiByteStr uintptr, cchMultiByte int, lpWideCharStr uintptr, cchWideChar int) int {
	r, _, _ := _MultiByteToWideChar.Call(uintptr(CodePage), uintptr(dwFlags), lpMultiByteStr, uintptr(cchMultiByte), lpWideCharStr, uintptr(cchWideChar))
	return int(r)
}

// WideCharToMultiByte
func WideCharToMultiByte(CodePage, dwFlags uint32, lpWideCharStr uintptr, cchWideChar int, lpMultiByteStr uintptr, cchMultiByte int, lpDefaultChar uintptr,
	lpUsedDefaultChar *int32) int {
	r, _, _ := _WideCharToMultiByte.Call(uintptr(CodePage), uintptr(dwFlags), lpWideCharStr, uintptr(cchWideChar), lpMultiByteStr, uintptr(cchMultiByte), lpDefaultChar,
		uintptr(unsafe.Pointer(lpUsedDefaultChar)))
	return int(r)
}
