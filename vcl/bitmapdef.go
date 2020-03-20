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

// LCL下的
func (b *TBitmap) Clear() {
	Bitmap_Clear(b.instance)
}

// LCL下的 用于ScanLine属性，aCanvasOnly 默认为 false
func (b *TBitmap) BeginUpdate(aCanvasOnly bool) {
	Bitmap_BeginUpdate(b.instance, aCanvasOnly)
}

// LCL下的 用于ScanLine属性，aStreamIsValid 默认为 false
func (b *TBitmap) EndUpdate(aStreamIsValid bool) {
	Bitmap_EndUpdate(b.instance, aStreamIsValid)
}

// LCL下的，Delphi也做相关的兼容
func (b *TBitmap) LoadFromDevice(dc types.HDC) {
	Bitmap_LoadFromDevice(b.instance, dc)
}
