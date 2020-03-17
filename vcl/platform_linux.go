//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
// +build linux

package vcl

type GdkWindow uintptr

type GtkWidget uintptr

type GtkFixed uintptr

type XID uintptr

// PGtkWidget
//func HandleToPlatformHandle(h types.HWND) GtkWidget {
//	return GtkWidget(h)
//}

func (f *TForm) PlatformWindow() GdkWindow {
	//r, _, _ := GdkWindow_FromForm.Call(f.instance)
	//return GdkWindow(r)
	return GdkWindow(0)
}

//func (g GdkWindow) XID() XID {
//	return 0
//}

// lz中首先是一个widget，然后上面用了一个fixedWidget来处理的。
func (g GtkWidget) FixedWidget() GtkFixed {
	return 0
}
