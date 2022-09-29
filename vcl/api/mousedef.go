//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import "github.com/ying32/govcl/vcl/api/dllimports"

func Mouse_Instance() uintptr {
	return defSyscallN(dllimports.MOUSE_INSTANCE)
}
