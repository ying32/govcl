//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build linux

package api

var (
	GdkWindow_GetXId      = libvcl.NewProc("GdkWindow_GetXId")
	GdkWindow_FromForm    = libvcl.NewProc("GdkWindow_FromForm")
	GtkWidget_GetGtkFixed = libvcl.NewProc("GtkWidget_GetGtkFixed")
	GtkWidget_Window      = libvcl.NewProc("GtkWidget_Window")
)
