//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------
// +build darwin

package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

type NSObject uintptr

type (
	// NSWindow
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

//func (n NSWindow) Release_() {
//		NSWindow_release.Call(uintptr(n))
//}
