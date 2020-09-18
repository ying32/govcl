//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"unsafe"

	"github.com/ying32/govcl/vcl/types"
)

func Clipboard_Instance() uintptr {
	r, _, _ := clipboard_Instance.Call()
	return r
}

func Clipboard_HasFormat(obj uintptr, aFormatID types.TClipboardFormat) bool {
	r, _, _ := clipboard_HasFormat.Call(obj, uintptr(aFormatID))
	return r != 0
}

func Clipboard_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if BufSize <= 0 {
		return 0
	}
	buff := make([]byte, BufSize)
	ret, _, _ := clipboard_GetTextBuf.Call(obj, uintptr(unsafe.Pointer(&buff[0])), uintptr(BufSize))
	if int(ret) < len(buff) {
		*Buffer = string(buff[:ret])
	}
	return int32(ret)
}

func Clipboard_GetAsText(obj uintptr) string {
	ret, _, _ := clipboard_GetAsText.Call(obj)
	return DStrToGoStr(ret)
}

func Clipboard_SetAsText(obj uintptr, value string) {
	clipboard_SetAsText.Call(obj, GoStrToDStr(value))
}

func Clipboard_GetAsHtml(obj uintptr, ExtractFragmentOnly bool) string {
	ret, _, _ := clipboard_GetAsHtml.Call(obj, GoBoolToDBool(ExtractFragmentOnly))
	return DStrToGoStr(ret)
}

func DSetClipboard(obj uintptr) uintptr {
	r, _, _ := dSetClipboard.Call(obj)
	return r
}

func DRegisterClipboardFormat(aFormat string) types.TClipboardFormat {
	r, _, _ := dRegisterClipboardFormat.Call(GoStrToDStr(aFormat))
	return types.TClipboardFormat(r)
}

func DPredefinedClipboardFormat(aFormat types.TPredefinedClipboardFormat) types.TClipboardFormat {
	r, _, _ := dPredefinedClipboardFormat.Call(uintptr(aFormat))
	return types.TClipboardFormat(r)
}
