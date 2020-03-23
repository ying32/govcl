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
	"unsafe"

	"fmt"
	"syscall"

	. "github.com/ying32/govcl/vcl/types"
)

const MAX_VERS = 20

// FixWindowsVersion
func FixWindowsVersion(MajorVersion, MinorVersion, ServicePackMajor *int) {

	var i int
	var osvi TOSVersionInfoEx
	var dwlConditionMask uint64
	var dwTypeMask uint32

	osvi.OSVersionInfoSize = uint32(unsafe.Sizeof(osvi))
	osvi.MinorVersion = 0
	osvi.ServicePackMajor = 0

	if *ServicePackMajor != -1 {
		dwTypeMask = VER_MAJORVERSION | VER_MINORVERSION | VER_SERVICEPACKMAJOR
	} else {
		dwTypeMask = VER_MAJORVERSION | VER_MINORVERSION
	}
	dwlConditionMask = VerSetConditionMask(0, VER_MAJORVERSION, VER_GREATER_EQUAL)
	dwlConditionMask = VerSetConditionMask(dwlConditionMask, VER_MINORVERSION, VER_GREATER_EQUAL)
	if *ServicePackMajor != -1 {
		dwlConditionMask = VerSetConditionMask(dwlConditionMask, VER_SERVICEPACKMAJOR, VER_GREATER_EQUAL)

	}

	//Try to find MajorVersion
	for i = 7; i <= MAX_VERS; i++ {
		osvi.MajorVersion = uint32(i)
		if VerifyVersionInfo(&osvi, dwTypeMask, dwlConditionMask) == false {
			*MajorVersion = i - 1
			break
		}
	}
	osvi.MajorVersion = uint32(*MajorVersion)

	//Try to find MinorVersion
	for i = 1; i <= MAX_VERS; i++ {
		osvi.MinorVersion = uint32(i)
		if VerifyVersionInfo(&osvi, dwTypeMask, dwlConditionMask) == false {
			*MinorVersion = i - 1
		}
	}
	osvi.MinorVersion = uint32(*MinorVersion)

	if *ServicePackMajor != -1 {
		//Try to find ServicePackMajor
		for i = 1; i <= MAX_VERS; i++ {
			osvi.ServicePackMajor = uint16(i)
			if VerifyVersionInfo(&osvi, dwTypeMask, dwlConditionMask) == false {
				*ServicePackMajor = i - 1
				break
			}
		}
	}
}

// CheckWin32Version
func CheckWin32Version(AMajor, AMinor int) bool {
	return (Win32MajorVersion > AMajor) ||
		((Win32MajorVersion == AMajor) &&
			(Win32MinorVersion >= AMinor))
}

// IsAdministrator 判断当前进程是否以Administrator权限运行中
func IsAdministrator() bool {
	var tokenHandle uintptr
	if CheckWin32Version(6, 0) {
		if OpenProcessToken(GetCurrentProcess(), TOKEN_QUERY, &tokenHandle) {
			defer CloseHandle(tokenHandle)
			var TokenIsElevated uint32
			var ReturnLength uint32
			if GetTokenInformation(tokenHandle, TokenElevation, uintptr(unsafe.Pointer(&TokenIsElevated)), uint32(unsafe.Sizeof(TokenIsElevated)), &ReturnLength) {
				if ReturnLength == uint32(unsafe.Sizeof(TokenIsElevated)) {
					return TokenIsElevated > 0
				}
			}
		}
		return false
	}
	return true
}

// 以管理员权限运行一个程序
func RunAsAdministrator(file, params, dir string) bool {
	var sei TShellExecuteInfo
	sei.CbSize = uint32(unsafe.Sizeof(sei))
	sei.Wnd = 0
	sei.FMask = SEE_MASK_FLAG_DDEWAIT | SEE_MASK_FLAG_NO_UI
	sei.LpVerb = CStr("runas")
	sei.LpFile = CStr(file)
	sei.LpParameters = CStr(params)
	sei.LpDirectory = CStr(dir)
	sei.NShow = SW_SHOWNORMAL
	return ShellExecuteEx(&sei)
}

// OpenInExplorer 在资源管理器中定位文件
func OpenInExplorer(aFileName string) {
	ShellExecute(0, "OPEN", "Explorer.exe",
		fmt.Sprintf("/e, /select, \"%s\"", aFileName), "", SW_SHOW)
}

// 内部两个
// GoStrToCStr
func CStr(str string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}

// CStrToGoStr
func GoStr(str []uint16) string {
	return syscall.UTF16ToString(str)
}

// CStrToGoStr
func GoPtrStr(str uintptr) string {
	l := LstrlenW(str)
	if l == 0 {
		return ""
	}
	buff := make([]uint16, l)
	Memcpy(uintptr(unsafe.Pointer(&buff[0])), str, uintptr(l*2))
	return GoStr(buff)
}

// GoBoolToCBool
func CBool(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// CBoolToGoBool
func GoBool(b uintptr) bool {
	if b != 0 {
		return true
	}
	return false
}

// UTF8ToANSI 将UTF-8字符转为ANSI格式
func UTF8ToANSI(str string) []uint8 {
	if str == "" {
		return nil
	}
	utf8StrPtr := uintptr(unsafe.Pointer(&([]byte(str + "\x00")[0])))
	nLen := MultiByteToWideChar(CP_UTF8, 0, utf8StrPtr, -1, 0, 0)
	wCharBuffer := make([]uint16, nLen+1)
	wCharBufferPtr := uintptr(unsafe.Pointer(&wCharBuffer[0]))
	MultiByteToWideChar(CP_UTF8, 0, utf8StrPtr, -1, wCharBufferPtr, nLen)

	nLen = WideCharToMultiByte(CP_ACP, 0, wCharBufferPtr, -1, 0, 0, 0, nil)
	aCharBuffer := make([]uint8, nLen) // +1
	aCharBufferPtr := uintptr(unsafe.Pointer(&aCharBuffer[0]))
	WideCharToMultiByte(CP_ACP, 0, wCharBufferPtr, -1, aCharBufferPtr, nLen, 0, nil)

	return aCharBuffer
}

// GetDesktopPath 获取桌面路径
func GetDesktopPath() string {
	var lpPath string
	if SHGetSpecialFolderPath(0, &lpPath, CSIDL_DESKTOP, false) {
		return lpPath
	}
	return ""
}

// ResourceToBytes 查找指定实例中 指定名称、指定类型 资源，并返回资源字节
func ResourceToBytes(instance uintptr, resName string, resType uintptr) ([]byte, bool) {
	resInfo := FindResource(HMODULE(instance), resName, resType)
	if resInfo == 0 {
		return nil, false
	}
	hGlobal := LoadResource(instance, resInfo)
	if hGlobal == 0 {
		return nil, false
	}
	ptr := LockResource(hGlobal)
	size := SizeofResource(instance, resInfo)
	bytes := make([]byte, size)
	Memcpy(uintptr(unsafe.Pointer(&bytes[0])), ptr, uintptr(size))
	//UnlockResource(hGlobal);
	FreeResource(hGlobal)
	return bytes, true
}
