//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

// 从资源中创建TFrame。
//
// Create TFrame from resources.
func CreateResFrame(owner IComponent, fields ...interface{}) {
	resObjectBuild(2, owner, 0, fields...)
}
