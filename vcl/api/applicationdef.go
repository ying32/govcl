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
		// 开启了finalizerOn选项后，以防止关闭库后GC还没开始调用。
		callGC()
		// 运行结束后就结束close掉lib，不然他不会关掉的
		closeLib()
	}()
	application_Run.Call(app)
}

func Application_Initialize(obj uintptr) {
	application_Initialize.Call(obj)
}
