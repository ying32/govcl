//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build darwin

package api

var (
	NSWindow_FromForm                      = libvcl.NewProc("NSWindow_FromForm")
	NSWindow_titleVisibility               = libvcl.NewProc("NSWindow_titleVisibility")
	NSWindow_setTitleVisibility            = libvcl.NewProc("NSWindow_setTitleVisibility")
	NSWindow_titlebarAppearsTransparent    = libvcl.NewProc("NSWindow_titlebarAppearsTransparent")
	NSWindow_setTitlebarAppearsTransparent = libvcl.NewProc("NSWindow_setTitlebarAppearsTransparent")
	NSWindow_styleMask                     = libvcl.NewProc("NSWindow_styleMask")
	NSWindow_setStyleMask                  = libvcl.NewProc("NSWindow_setStyleMask")
	NSWindow_setRepresentedURL             = libvcl.NewProc("NSWindow_setRepresentedURL")
	//NSWindow_release                       = libvcl.NewProc("NSWindow_release")
)
