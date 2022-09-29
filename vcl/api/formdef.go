//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import "github.com/ying32/govcl/vcl/api/dllimports"

func Form_Create2(owner uintptr, initScale bool) uintptr {
	return defSyscallN(dllimports.FORM_CREATE2, owner, PascalBool(initScale))
}

func Form_SetOnWndProc(obj uintptr, fn interface{}) {
	defSyscallN(dllimports.FORM_SETONWNDPROC, obj, MakeEventDataPtr(fn))
}

func Form_SetGoPtr(obj uintptr, ptr uintptr) {
	defSyscallN(dllimports.FORM_SETGOPTR, obj, ptr)
}
