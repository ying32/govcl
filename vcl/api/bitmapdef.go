//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import "github.com/ying32/govcl/vcl/types"

// TBitmap
func Bitmap_Clear(obj uintptr) {
	bitmap_Clear.Call(obj)
}

func Bitmap_BeginUpdate(obj uintptr, aCanvasOnly bool) {
	bitmap_BeginUpdate.Call(obj, GoBoolToDBool(aCanvasOnly))
}

func Bitmap_EndUpdate(obj uintptr, aStreamIsValid bool) {
	bitmap_EndUpdate.Call(obj, GoBoolToDBool(aStreamIsValid))
}

func Bitmap_LoadFromDevice(obj uintptr, dc types.HDC) {
	bitmap_LoadFromDevice.Call(obj, uintptr(dc))
}
