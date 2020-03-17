//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import "github.com/ying32/govcl/vcl/types"

type NSObject uintptr

// Handle
//func HandleToPlatformHandle(h types.HWND) types.HWND {
//	return h
//}

func (f *TForm) PlatformWindow() types.HWND {
	return f.Handle()
}
