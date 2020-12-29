//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

func Form_Create2(owner uintptr, initScale bool) uintptr {
	ret, _, _ := form_Create2.Call(owner, GoBoolToDBool(initScale))
	return ret
}

func Form_SetOnWndProc(obj uintptr, fn interface{}) {
	form_SetOnWndProc.Call(obj, addMessageEventToMap(obj, fn))
}

func Form_SetGoPtr(obj uintptr, ptr uintptr) {
	form_SetGoPtr.Call(obj, ptr)
}
