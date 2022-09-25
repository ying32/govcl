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

var (
	NSWindow_FromForm                      = newDLLProc("NSWindow_FromForm")
	NSWindow_titleVisibility               = newDLLProc("NSWindow_titleVisibility")
	NSWindow_setTitleVisibility            = newDLLProc("NSWindow_setTitleVisibility")
	NSWindow_titlebarAppearsTransparent    = newDLLProc("NSWindow_titlebarAppearsTransparent")
	NSWindow_setTitlebarAppearsTransparent = newDLLProc("NSWindow_setTitlebarAppearsTransparent")
	NSWindow_styleMask                     = newDLLProc("NSWindow_styleMask")
	NSWindow_setStyleMask                  = newDLLProc("NSWindow_setStyleMask")
	NSWindow_setRepresentedURL             = newDLLProc("NSWindow_setRepresentedURL")
	//NSWindow_release                       = newDLLProc("NSWindow_release")
)
