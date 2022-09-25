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

var (
	GdkWindow_GetXId      = newDLLProc("GdkWindow_GetXId")
	GdkWindow_FromForm    = newDLLProc("GdkWindow_FromForm")
	GtkWidget_GetGtkFixed = newDLLProc("GtkWidget_GetGtkFixed")
	GtkWidget_Window      = newDLLProc("GtkWidget_Window")
)
