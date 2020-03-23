//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

func Clipboard_Instance() uintptr {
	r, _, _ := clipboard_Instance.Call()
	return r
}

func Clipboard_SetClipboard(obj uintptr) uintptr {
	r, _, _ := clipboard_SetClipboard.Call(obj)
	return r
}
