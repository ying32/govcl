//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import "github.com/ying32/govcl/vcl/api/dllimports"

func Printer_Instance() uintptr {
	return defSyscallN(dllimports.PRINTER_INSTANCE)
}

func Printer_SetPrinter(obj uintptr, aName string) {
	defSyscallN(dllimports.PRINTER_SETPRINTER, obj, PascalStr(aName))
}
