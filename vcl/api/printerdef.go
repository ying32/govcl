//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

// Printer
func Printer_Instance() uintptr {
	r, _, _ := printer_Instance.Call()
	return r
}
