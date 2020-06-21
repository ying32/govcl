//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

func Application_Instance() uintptr {
	ret, _, _ := application_Instance.Call()
	return ret
}

func Application_CreateForm(app uintptr, initScale bool) uintptr {
	ret, _, _ := application_CreateForm.Call(app, GoBoolToDBool(initScale))
	return ret
}

func Application_Run(app uintptr) {
	defer func() {
		// 运行完后free下
		closeLib()
	}()
	application_Run.Call(app)
}

func Application_Initialize(obj uintptr) {
	application_Initialize.Call(obj)
}

func Application_GetHandle(obj uintptr) uintptr {
	ret, _, _ := application_GetHandle.Call(obj)
	return ret
}

func Application_SetHandle(obj uintptr, handle uintptr) {
	application_SetHandle.Call(obj, handle)
}
