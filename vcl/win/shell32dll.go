// +build windows

package win

import (
	"syscall"

	"gitee.com/ying32/govcl/vcl/types"
)

var (

	// shell32.dll
	shell32dll = syscall.NewLazyDLL("shell32.dll")

	_ShellExecute = shell32dll.NewProc("ShellExecuteW")
	_ExtractIcon  = shell32dll.NewProc("ExtractIconW")
)

// ShellExecute
func ShellExecute(hWnd types.HWND, Operation, FileName, Parameters, Directory string, ShowCmd int32) uintptr {
	r, _, _ := _ShellExecute.Call(uintptr(hWnd), CStr(Operation), CStr(FileName), CStr(Parameters), CStr(Directory), uintptr(ShowCmd))
	return r
}

// ExtractIcon
func ExtractIcon(hInst uintptr, lpszExeFileName string, nIconIndex uint32) types.HICON {
	r, _, _ := _ExtractIcon.Call(hInst, CStr(lpszExeFileName), uintptr(nIconIndex))
	return types.HICON(r)
}
