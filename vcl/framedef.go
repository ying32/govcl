//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

// CN: 从资源中创建TFrame。
// EN: Create TFrame from resources.
func CreateResFrame(owner IComponent, fields ...interface{}) {
	resObjtBuild(2, owner, 0, fields...)
}
