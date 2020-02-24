// +build windows


//----------------------------------------
// 
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package win

import (
	"syscall"

	"unsafe"

	. "github.com/ying32/govcl/vcl/types"
)

var (

	// shell32.dll
	shell32dll = syscall.NewLazyDLL("shell32.dll")

	_ShellExecute = shell32dll.NewProc("ShellExecuteW")
	_ExtractIcon  = shell32dll.NewProc("ExtractIconW")

	_SHChangeNotify             = shell32dll.NewProc("SHChangeNotify")
	_SHGetSpecialFolderPath     = shell32dll.NewProc("SHGetSpecialFolderPathW")
	_SHGetSpecialFolderLocation = shell32dll.NewProc("SHGetSpecialFolderLocation")
	_SHGetPathFromIDList        = shell32dll.NewProc("SHGetPathFromIDListW")
)

// ShellExecute
func ShellExecute(hWnd HWND, Operation, FileName, Parameters, Directory string, ShowCmd int32) uintptr {
	r, _, _ := _ShellExecute.Call(uintptr(hWnd), CStr(Operation), CStr(FileName), CStr(Parameters), CStr(Directory), uintptr(ShowCmd))
	return r
}

// ExtractIcon
func ExtractIcon(hInst uintptr, lpszExeFileName string, nIconIndex uint32) HICON {
	r, _, _ := _ExtractIcon.Call(hInst, CStr(lpszExeFileName), uintptr(nIconIndex))
	return HICON(r)
}

// SHChangeNotify
func SHChangeNotify(wEventId int32, uFlags uint32, dwItem1, dwItem2 uintptr) {
	_SHChangeNotify.Call(uintptr(wEventId), uintptr(uFlags), dwItem1, dwItem2)
}

// SHGetSpecialFolderPath
func SHGetSpecialFolderPath(hwndOwner HWND, lpszPath *string, nFolder int32, fCreate bool) bool {
	buff := make([]uint16, 255)
	r, _, _ := _SHGetSpecialFolderPath.Call(uintptr(hwndOwner), uintptr(unsafe.Pointer(&buff[0])), uintptr(nFolder), CBool(fCreate))
	*lpszPath = GoStr(buff)
	return r != 0
}

func SHGetSpecialFolderLocation(hwndOwner HWND, nFolder int32, ppidl **TItemIDList) HResult {
	r, _, _ := _SHGetSpecialFolderLocation.Call(uintptr(hwndOwner), uintptr(nFolder), uintptr(unsafe.Pointer(ppidl)))
	return HResult(r)
}

func SHGetPathFromIDList(pidl *TItemIDList) (bool, string) {
	szPath := make([]uint16, MAX_PATH)
	r, _, _ := _SHGetPathFromIDList.Call(uintptr(unsafe.Pointer(pidl)), uintptr(unsafe.Pointer(&szPath[0])))
	return r != 0, GoStr(szPath)
}
