//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build linux

package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
)

type GdkWindow uintptr

type GtkWidget uintptr

type GtkFixed uintptr

type XID uintptr

// PGtkWidget
func HandleToPlatformHandle(h types.HWND) GtkWidget {
	return GtkWidget(h)
}

func (f *TForm) PlatformWindow() GdkWindow {
	return GdkWindow(GdkWindow_FromForm(f._instance()))
}

func (g GdkWindow) XID() (xid XID) {
	return XID(GdkWindow_GetXId(uintptr(g)))
}

// FixedWidget
//
// lz中首先是一个widget，然后上面用了一个fixedWidget来处理的。
func (g GtkWidget) FixedWidget() GtkFixed {
	return GtkFixed(GtkWidget_GetGtkFixed(uintptr(g)))
}

func (g GtkWidget) Window() GdkWindow {
	return GdkWindow(GtkWidget_Window(uintptr(g)))
}
