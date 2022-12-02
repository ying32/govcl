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

func messageCallbackProc(f uintptr, msg uintptr) uintptr {
	v := PtrToElementValue(f)
	if v != nil {
		v.(TWndProcEvent)(
			(*TMessage)(unsafe.Pointer(msg)),
		)
	}
	return 0
}
