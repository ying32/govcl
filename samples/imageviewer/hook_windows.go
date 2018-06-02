package main

import (
	"fmt"
	"syscall"

	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

// 用了另外一个方法了，这里不要了，先不管了

var oldWndProcPtr uintptr

func hookMainFormWndPrc() {
	if oldWndProcPtr == 0 {
		newWndProcPtr := syscall.NewCallback(newMainFormWndProc)
		oldWndProcPtr = win.SetWindowLongPtr(MainForm.Handle(), win.GWL_WNDPROC, newWndProcPtr)
		fmt.Println("newWndProcPtr:", newWndProcPtr, ", oldWndProcPtr:", oldWndProcPtr)
	}
}

func unHookMainFormWndPrc() {
	if oldWndProcPtr != 0 {
		win.SetWindowLongPtr(MainForm.Handle(), win.GWL_WNDPROC, oldWndProcPtr)
		oldWndProcPtr = 0
	}
}

func newMainFormWndProc(hWnd uintptr, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case win.WM_SYSCOMMAND:
		switch wParam {
		case win.SC_MAXIMIZE:

		case win.SC_MINIMIZE:

		case win.SC_RESTORE:

		case win.SC_CLOSE:

		}
	// 使用了另外的api所以上面的消息捕不到，要捕捉这个
	case win.SW_SHOWMAXIMIZED:
		if MainForm.WindowState() == types.WsMaximized {
			fmt.Println("最大化:", wParam, lParam)
		}
	}
	return win.CallWindowProc(oldWndProcPtr, types.HWND(hWnd), message, wParam, lParam)
}
