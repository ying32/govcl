//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

type GtkWindow uintptr

type GtkWidget uintptr

// PGtkWidget
//func HandleToPlatformHandle(h types.HWND) GtkWidget {
//	return GtkWidget(h)
//}

func (f *TForm) PlatformWindow() GtkWindow {
	//r, _, _ := NSWindow_FromForm.Call(f.instance)
	//return GtkWindow(r)
	return GtkWindow(0)
}
