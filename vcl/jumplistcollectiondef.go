//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import . "github.com/ying32/govcl/vcl/api"

// Add2 因自动生成的不是相应对象，这里重载个函数用于简化
func (j *TJumpListCollection) Add2() *TJumpListItem {
	return AsJumpListItem(JumpListCollection_Add(j.instance))
}
