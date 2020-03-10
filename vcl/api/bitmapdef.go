//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

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
