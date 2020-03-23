//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import . "github.com/ying32/govcl/vcl/api"

func (j *TJumpCategories) Add2() *TJumpCategoryItem {
	return AsJumpCategoryItem(JumpCategories_Add(j.instance))
}
