//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import "github.com/ying32/govcl/vcl/api/dllimports"

func Application_Instance() uintptr {
	return defSyscallN(dllimports.APPLICATION_INSTANCE)
}

func Application_CreateForm(app uintptr, initScale bool) uintptr {
	return defSyscallN(dllimports.APPLICATION_CREATEFORM, app, PascalBool(initScale))
}

func Application_Run(app uintptr) {
	defer func() {
		// 开启了finalizerOn选项后，以防止关闭库后GC还没开始调用。
		callGC()
		// 运行结束后就结束close掉lib，不然他不会关掉的
		closeLib()
	}()
	defSyscallN(dllimports.APPLICATION_RUN, app)
}

func Application_Initialize(obj uintptr) {
	defSyscallN(dllimports.APPLICATION_INITIALIZE, obj)
}
