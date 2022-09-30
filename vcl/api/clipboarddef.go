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

	"github.com/ying32/govcl/vcl/api/dllimports"

	"github.com/ying32/govcl/vcl/types"
)

func Clipboard_Instance() uintptr {
	return defSyscallN(dllimports.CLIPBOARD_INSTANCE)
}

func Clipboard_HasFormat(obj uintptr, aFormatID types.TClipboardFormat) bool {
	return GoBool(defSyscallN(dllimports.CLIPBOARD_HASFORMAT, obj, uintptr(aFormatID)))
}

func Clipboard_GetTextBuf(obj uintptr, Buffer *string, BufSize int32) int32 {
	if BufSize <= 0 {
		return 0
	}
	buff := make([]byte, BufSize)
	ret := defSyscallN(dllimports.CLIPBOARD_GETTEXTBUF, obj, uintptr(unsafe.Pointer(&buff[0])), uintptr(BufSize))
	if int(ret) < len(buff) {
		*Buffer = string(buff[:ret])
	}
	return int32(ret)
}

func Clipboard_GetAsText(obj uintptr) string {
	return GoStr(defSyscallN(dllimports.CLIPBOARD_GETASTEXT, obj))
}

func Clipboard_SetAsText(obj uintptr, value string) {
	defSyscallN(dllimports.CLIPBOARD_SETASTEXT, obj, PascalStr(value))
}

func Clipboard_GetAsHtml(obj uintptr, ExtractFragmentOnly bool) string {
	return GoStr(defSyscallN(dllimports.CLIPBOARD_GETASHTML, obj, PascalBool(ExtractFragmentOnly)))
}

func DSetClipboard(obj uintptr) uintptr {
	return defSyscallN(dllimports.DSETCLIPBOARD, obj)
}

func DRegisterClipboardFormat(aFormat string) types.TClipboardFormat {
	return types.TClipboardFormat(defSyscallN(dllimports.DREGISTERCLIPBOARDFORMAT, PascalStr(aFormat)))
}

func DPredefinedClipboardFormat(aFormat types.TPredefinedClipboardFormat) types.TClipboardFormat {
	return types.TClipboardFormat(defSyscallN(dllimports.DPREDEFINEDCLIPBOARDFORMAT, uintptr(aFormat)))
}
