//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build linux
// +build linux

package api

import (
	"unsafe"

	"github.com/ying32/govcl/vcl/api/dllimports"
)

func GdkWindow_GetXId(g uintptr) (res uintptr) {
	defSyscallN(dllimports.GDKWINDOW_GETXID, g, uintptr(unsafe.Pointer(&res)))
	return
}

func GdkWindow_FromForm(obj uintptr) uintptr {
	return defSyscallN(dllimports.GDKWINDOW_FROMFORM, obj)
}

func GtkWidget_GetGtkFixed(g uintptr) uintptr {
	return defSyscallN(dllimports.GTKWIDGET_GETGTKFIXED, g)
}

func GtkWidget_Window(g uintptr) uintptr {
	return defSyscallN(dllimports.GTKWIDGET_WINDOW, g)
}
