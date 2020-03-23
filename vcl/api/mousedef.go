//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

func Mouse_Instance() uintptr {
	ret, _, _ := mouse_Instance.Call()
	return ret
}
