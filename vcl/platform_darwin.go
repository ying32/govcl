//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build darwin

package vcl

/*

   #cgo CFLAGS: -x objective-c
   #cgo LDFLAGS: -framework Cocoa

   #include <Cocoa/Cocoa.h>

  static NSWindow* toNSWindow(void* ptr) {
      return ((__bridge NSWindow*)ptr); //
  }

  // test
  static void NSWindow_setTitleVisibility(void *ptr) {
     NSWindow *win = toNSWindow(ptr);
     win.TitleVisibility = NSWindowTitleHidden;
  }

*/
//import "C"

import (
	. "github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

type (
	NSObject uintptr

	NSWindow uintptr

	NSURL uintptr
)

// NSObject
func HandleToPlatformHandle(h HWND) NSObject {
	return NSObject(h)
}

func (f *TForm) PlatformWindow() NSWindow {
	r, _, _ := NSWindow_FromForm.Call(f.instance)
	return NSWindow(r)
}

func (n NSWindow) TitleVisibility() NSWindowTitleVisibility {
	r, _, _ := NSWindow_titleVisibility.Call(uintptr(n))
	return NSWindowTitleVisibility(r)
}

func (n NSWindow) SetTitleVisibility(flag NSWindowTitleVisibility) {
	//C.NSWindow_setTitleVisibility(unsafe.Pointer(n))
	NSWindow_setTitleVisibility.Call(uintptr(n), uintptr(flag))
}

func (n NSWindow) TitleBarAppearsTransparent() bool {
	r, _, _ := NSWindow_titlebarAppearsTransparent.Call(uintptr(n))
	return DBoolToGoBool(r)
}

func (n NSWindow) SetTitleBarAppearsTransparent(flag bool) {
	NSWindow_setTitlebarAppearsTransparent.Call(uintptr(n), GoBoolToDBool(flag))
}

func (n NSWindow) SetRepresentedURL(url NSURL) {
	NSWindow_setRepresentedURL.Call(uintptr(n), uintptr(url))
}

func (n NSWindow) StyleMask() uint {
	r, _, _ := NSWindow_styleMask.Call(uintptr(n))
	return uint(r)
}

func (n NSWindow) SetStyleMask(mask uint) {
	NSWindow_setStyleMask.Call(uintptr(n), uintptr(mask))
}
