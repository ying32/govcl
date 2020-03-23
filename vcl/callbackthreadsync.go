//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"github.com/ying32/govcl/vcl/api"
)

// 线程同步的，独立出来
func threadSyncCallbackProc() uintptr {
	fn := api.ThreadSyncCallbackFn()
	if fn != nil {
		fn()
	}
	return 0
}
