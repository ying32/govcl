//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
)

// 清除bitmap数据
//
// Clear bitmap data.
func (b *TBitmap) Clear() {
	Bitmap_Clear(b.instance)
}

// 用于ScanLine属性，aCanvasOnly 默认为 false
//
// Used for ScanLine properties, aCanvasOnly defaults to false.
func (b *TBitmap) BeginUpdate(aCanvasOnly bool) {
	Bitmap_BeginUpdate(b.instance, aCanvasOnly)
}

// 用于ScanLine属性，aStreamIsValid 默认为 false
//
// Used for ScanLine property, aStreamIsValid defaults to false.
func (b *TBitmap) EndUpdate(aStreamIsValid bool) {
	Bitmap_EndUpdate(b.instance, aStreamIsValid)
}

// 从设备驱动中加载Bitmap。
//
// Load the Bitmap from the device driver.
func (b *TBitmap) LoadFromDevice(dc types.HDC) {
	Bitmap_LoadFromDevice(b.instance, dc)
}
