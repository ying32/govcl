//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package api

import "github.com/ying32/govcl/vcl/api/dllimports"

func NSWindow_FromForm(obj uintptr) uintptr {
	return defSyscallN(dllimports.NSWINDOW_FROMFORM, obj)
}

func NSWindow_titleVisibility(n uintptr) uintptr {
	return defSyscallN(dllimports.NSWINDOW_TITLEVISIBILITY, n)
}

func NSWindow_setTitleVisibility(n uintptr, flag uintptr) {
	defSyscallN(dllimports.NSWINDOW_SETTITLEVISIBILITY, n, flag)
}

func NSWindow_titlebarAppearsTransparent(n uintptr) bool {
	return GoBool(defSyscallN(dllimports.NSWINDOW_TITLEBARAPPEARSTRANSPARENT, n))
}

func NSWindow_setTitlebarAppearsTransparent(n uintptr, flag bool) {
	defSyscallN(dllimports.NSWINDOW_SETTITLEBARAPPEARSTRANSPARENT, n, PascalBool(flag))
}

func NSWindow_setRepresentedURL(n uintptr, url uintptr) {
	defSyscallN(dllimports.NSWINDOW_SETREPRESENTEDURL, n, url)
}

func NSWindow_styleMask(n uintptr) uint {
	return uint(defSyscallN(dllimports.NSWINDOW_STYLEMASK, n))
}

func NSWindow_setStyleMask(n uintptr, mask uint) {
	defSyscallN(dllimports.NSWINDOW_SETSTYLEMASK, n, uintptr(mask))
}
