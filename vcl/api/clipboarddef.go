//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import "github.com/ying32/govcl/vcl/types"

func Clipboard_Instance() uintptr {
	r, _, _ := clipboard_Instance.Call()
	return r
}

func Clipboard_HasFormat(obj uintptr, aFormatID types.TClipboardFormat) bool {
	r, _, _ := clipboard_HasFormat.Call(obj, uintptr(aFormatID))
	return r != 0
}

func DSetClipboard(obj uintptr) uintptr {
	r, _, _ := dSetClipboard.Call(obj)
	return r
}

func DRegisterClipboardFormat(aFormat string) types.TClipboardFormat {
	r, _, _ := dRegisterClipboardFormat.Call(GoStrToDStr(aFormat))
	return types.TClipboardFormat(r)
}
