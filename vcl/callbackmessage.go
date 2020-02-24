
//----------------------------------------
// 
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl

import (
	"unsafe"

	. "github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

func messageCallbackProc(f uintptr, msg, handled uintptr) uintptr {
	if v, ok := MessageCallbackOf(f); ok {
		v.(TWndProcEvent)(
			(*TMessage)(unsafe.Pointer(msg)),
			(*bool)(unsafe.Pointer(handled)),
		)
	}
	return 0
}
