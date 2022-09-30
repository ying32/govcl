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
	return NSWindow(NSWindow_FromForm(f._instance()))
}

func (n NSWindow) TitleVisibility() NSWindowTitleVisibility {
	return NSWindowTitleVisibility(NSWindow_titleVisibility(uintptr(n)))
}

func (n NSWindow) SetTitleVisibility(flag NSWindowTitleVisibility) {
	//C.NSWindow_setTitleVisibility(unsafe.Pointer(n))
	NSWindow_setTitleVisibility(uintptr(n), uintptr(flag))
}

func (n NSWindow) TitleBarAppearsTransparent() bool {
	return NSWindow_titlebarAppearsTransparent(uintptr(n))
}

func (n NSWindow) SetTitleBarAppearsTransparent(flag bool) {
	NSWindow_setTitlebarAppearsTransparent(uintptr(n), flag)
}

func (n NSWindow) SetRepresentedURL(url NSURL) {
	NSWindow_setRepresentedURL(uintptr(n), uintptr(url))
}

func (n NSWindow) StyleMask() uint {
	return NSWindow_styleMask(uintptr(n))
}

func (n NSWindow) SetStyleMask(mask uint) {
	NSWindow_setStyleMask(uintptr(n), mask)
}
