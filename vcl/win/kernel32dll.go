//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows

package win

import (
	"fmt"
	"syscall"
	"unsafe"

	. "github.com/ying32/govcl/vcl/types"
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

	_LoadLibrary              = kernel32dll.NewProc("LoadLibraryW")
	_LoadLibraryEx            = kernel32dll.NewProc("LoadLibraryExW")
	_GetModuleFileName        = kernel32dll.NewProc("GetModuleFileNameW")
	_GetProcAddress           = kernel32dll.NewProc("GetProcAddress")
	_FreeLibrary              = kernel32dll.NewProc("FreeLibrary")
	_FreeLibraryAndExitThread = kernel32dll.NewProc("FreeLibraryAndExitThread")

	_HeapFree       = kernel32dll.NewProc("HeapFree")
	_HeapAlloc      = kernel32dll.NewProc("HeapAlloc")
	_GetProcessHeap = kernel32dll.NewProc("GetProcessHeap")
	_VirtualFree    = kernel32dll.NewProc("VirtualFree")
	_lstrlen        = kernel32dll.NewProc("lstrlenA")
	_lstrlenW       = kernel32dll.NewProc("lstrlenW")
	_VirtualAlloc   = kernel32dll.NewProc("VirtualAlloc")
	_VirtualProtect = kernel32dll.NewProc("VirtualProtect")
	_IsBadReadPtr   = kernel32dll.NewProc("IsBadReadPtr")

	_TerminateProcess    = kernel32dll.NewProc("TerminateProcess")
	_TerminateThread     = kernel32dll.NewProc("TerminateThread")
	_ExitThread          = kernel32dll.NewProc("ExitThread")
	_ExitProcess         = kernel32dll.NewProc("ExitProcess")
	_WaitForSingleObject = kernel32dll.NewProc("WaitForSingleObject")

	_FindResource   = kernel32dll.NewProc("FindResourceW")
	_LoadResource   = kernel32dll.NewProc("LoadResource")
	_LockResource   = kernel32dll.NewProc("LockResource")
	_SizeofResource = kernel32dll.NewProc("SizeofResource")
	_FreeResource   = kernel32dll.NewProc("FreeResource")

	_GlobalAddAtom    = kernel32dll.NewProc("GlobalAddAtomW")
	_GlobalDeleteAtom = kernel32dll.NewProc("GlobalDeleteAtom")

	_VirtualQueryEx     = kernel32dll.NewProc("VirtualQueryEx")
	_ReadProcessMemory  = kernel32dll.NewProc("ReadProcessMemory")
	_WriteProcessMemory = kernel32dll.NewProc("WriteProcessMemory")
	_GetSystemInfo      = kernel32dll.NewProc("GetSystemInfo")

	_OutputDebugString = kernel32dll.NewProc("OutputDebugStringW")
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

// LoadLibrary
func LoadLibrary(lpLibFileName string) uintptr {
	r, _, _ := _LoadLibrary.Call(CStr(lpLibFileName))
	return r
}

// LoadLibraryEx
func LoadLibraryEx(lpLibFileName string, hFile uintptr, dwFlags uint32) uintptr {
	r, _, _ := _LoadLibraryEx.Call(CStr(lpLibFileName), hFile, uintptr(dwFlags))
	return r
}

// GetModuleFileName
func GetModuleFileName(hModule uintptr) (string, uint32) {
	buff := make([]uint16, MAX_PATH)
	r, _, _ := _GetModuleFileName.Call(hModule, uintptr(unsafe.Pointer(&buff[0])), MAX_PATH*2)
	return GoStr(buff), uint32(r)
}

// GetProcAddress
func GetProcAddress(hModule uintptr, lpProcName string) uintptr {
	r, _, _ := _GetProcAddress.Call(hModule, uintptr(unsafe.Pointer(&UTF8ToANSI(lpProcName)[0])))
	return r
}

// FreeLibrary
func FreeLibrary(hLibModule uintptr) bool {
	r, _, _ := _FreeLibrary.Call(hLibModule)
	return r != 0
}

// FreeLibraryAndExitThread
func FreeLibraryAndExitThread(hLibModule uintptr, dwExitCode uint32) {
	_FreeLibraryAndExitThread.Call(hLibModule, uintptr(dwExitCode))
}

func HeapFree(hHeap uintptr, dwFlags uint32, lpMem uintptr) bool {
	r, _, _ := _HeapFree.Call(hHeap, uintptr(dwFlags), lpMem)
	return r != 0
}

func HeapAlloc(hHeap uintptr, dwFlags uint32, dwBytes uintptr) uintptr {
	r, _, _ := _HeapAlloc.Call(hHeap, uintptr(dwFlags), dwBytes)
	return r
}

func GetProcessHeap() uintptr {
	r, _, _ := _GetProcessHeap.Call()
	return r
}

func VirtualFree(lpAddress uintptr, dwSize uintptr, dwFreeType uint32) bool {
	r, _, _ := _VirtualFree.Call(lpAddress, dwSize, uintptr(dwFreeType))
	return r != 0
}

func Lstrlen(lpString uintptr) int32 {
	r, _, _ := _lstrlen.Call(lpString)
	return int32(r)
}

func LstrlenW(lpString uintptr) int32 {
	r, _, _ := _lstrlenW.Call(lpString)
	return int32(r)
}

func VirtualAlloc(lpvAddress uintptr, dwSize uintptr, flAllocationType, flProtect uint32) uintptr {
	r, _, _ := _VirtualAlloc.Call(lpvAddress, dwSize, uintptr(flAllocationType), uintptr(flProtect))
	return r
}

func VirtualProtect(lpAddress uintptr, dwSize uintptr, flNewProtect uint32, lpflOldProtect uintptr) bool {
	r, _, _ := _VirtualProtect.Call(lpAddress, dwSize, uintptr(flNewProtect), lpflOldProtect)
	return r != 0
}

func IsBadReadPtr(lp uintptr, ucb uintptr) bool {
	r, _, _ := _IsBadReadPtr.Call(lp, ucb)
	return r != 0
}

func TerminateProcess(hProcess uintptr, uExitCode uint32) bool {
	r, _, _ := _TerminateProcess.Call(hProcess, uintptr(uExitCode))
	return r != 0
}

func TerminateThread(hThread uintptr, dwExitCode uint32) bool {
	r, _, _ := _TerminateThread.Call(hThread, uintptr(dwExitCode))
	return r != 0
}

func ExitThread(dwExitCode uint32) {
	_ExitThread.Call(uintptr(dwExitCode))
}

func ExitProcess(uExitCode uint32) {
	_ExitProcess.Call(uintptr(uExitCode))
}

func WaitForSingleObject(hHandle uintptr, dwMilliseconds uint32) uint32 {
	r, _, _ := _WaitForSingleObject.Call(hHandle, uintptr(dwMilliseconds))
	return uint32(r)
}

func FindResource(hModule HMODULE, lpName string, lpType uintptr) HRSRC {
	r, _, _ := _FindResource.Call(uintptr(hModule), CStr(lpName), lpType)
	return HRSRC(r)
}

func LoadResource(hModule uintptr, hResInfo HRSRC) HGLOBAL {
	r, _, _ := _LoadResource.Call(hModule, uintptr(hResInfo))
	return HGLOBAL(r)
}

func LockResource(hResData HGLOBAL) uintptr {
	r, _, _ := _LockResource.Call(uintptr(hResData))
	return r
}

func SizeofResource(hModule uintptr, hResInfo HRSRC) uint32 {
	r, _, _ := _SizeofResource.Call(hModule, uintptr(hResInfo))
	return uint32(r)
}

func FreeResource(hResData HGLOBAL) bool {
	r, _, _ := _FreeResource.Call(uintptr(hResData))
	return r != 0
}

func GlobalAddAtom(lpString string) ATOM {
	r, _, _ := _GlobalAddAtom.Call(CStr(lpString))
	return ATOM(r)
}

func GlobalDeleteAtom(nAtom ATOM) ATOM {
	r, _, _ := _GlobalDeleteAtom.Call(uintptr(nAtom))
	return ATOM(r)
}

func VirtualQueryEx(hProcess uintptr, lpAddress uintptr, lpBuffer *TMemoryBasicInformation, dwLength SIZE_T) SIZE_T {
	r, _, _ := _VirtualQueryEx.Call(hProcess, lpAddress, uintptr(unsafe.Pointer(lpBuffer)), dwLength)
	return r
}

func ReadProcessMemoryBytes(hProcess uintptr, lpBaseAddress uintptr, nSize SIZE_T) ([]byte, bool) {
	var lpNumberOfBytesRead SIZE_T
	buffer := make([]byte, nSize)
	r, _, _ := _ReadProcessMemory.Call(hProcess, lpBaseAddress, uintptr(unsafe.Pointer(&buffer[0])), nSize, uintptr(unsafe.Pointer(&lpNumberOfBytesRead)))
	if r != 0 {
		return buffer[:lpNumberOfBytesRead], true
	}
	return nil, false
}

func ReadProcessMemory(hProcess uintptr, lpBaseAddress uintptr, lpBuffer uintptr, nSize SIZE_T, lpNumberOfBytesRead *SIZE_T) bool {
	r, _, _ := _ReadProcessMemory.Call(hProcess, lpBaseAddress, lpBuffer, nSize, uintptr(unsafe.Pointer(lpNumberOfBytesRead)))
	return r != 0
}

func WriteProcessMemoryBytes(hProcess uintptr, lpBaseAddress uintptr, lpBuffer []byte, nSize SIZE_T) (SIZE_T, bool) {
	var lpNumberOfBytesWritten SIZE_T
	r, _, _ := _WriteProcessMemory.Call(hProcess, lpBaseAddress, uintptr(unsafe.Pointer(&lpBuffer[0])), nSize, uintptr(unsafe.Pointer(&lpNumberOfBytesWritten)))
	if r != 0 {
		return lpNumberOfBytesWritten, true
	}
	return 0, false
}

func WriteProcessMemory(hProcess uintptr, lpBaseAddress uintptr, lpBuffer uintptr, nSize SIZE_T, lpNumberOfBytesWritten *SIZE_T) bool {
	r, _, _ := _WriteProcessMemory.Call(hProcess, lpBaseAddress, lpBuffer, nSize, uintptr(unsafe.Pointer(lpNumberOfBytesWritten)))
	return r != 0
}

func GetSystemInfo(lpSystemInfo *TSystemInfo) bool {
	r, _, _ := _GetSystemInfo.Call(uintptr(unsafe.Pointer(lpSystemInfo)))
	return r != 0
}

func OutputDebugString(v ...interface{}) {
	_OutputDebugString.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fmt.Sprintln(v...)))))
}
