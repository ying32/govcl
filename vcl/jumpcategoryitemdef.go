
//----------------------------------------
// 
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl

import . "github.com/ying32/govcl/vcl/api"

// Add2
func (j *TJumpCategories) Add2() *TJumpCategoryItem {
	return JumpCategoryItemFromInst(JumpCategories_Add(j.instance))
}
