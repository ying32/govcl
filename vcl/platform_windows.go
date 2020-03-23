//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows

package vcl

import . "github.com/ying32/govcl/vcl/types"

type Window HWND

// Handle
func HandleToPlatformHandle(h HWND) Window {
	return Window(h)
}

func (f *TForm) PlatformWindow() Window {
	return Window(f.Handle())
}

//func (w Window) SendMessage(msg uint32, wParam, lParam uintptr) uintptr {
//	return win.SendMessage(HWND(w), msg, wParam, lParam)
//}
//
//func (w Window) PostMessage(msg uint32, wParam, lParam uintptr) uintptr {
//	return win.PostMessage(HWND(w), msg, wParam, lParam)
//}
