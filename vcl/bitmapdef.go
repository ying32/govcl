//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import . "github.com/ying32/govcl/vcl/api"

// LCL下的
func (b *TBitmap) Clear() {
	Bitmap_Clear(b.instance)
}
