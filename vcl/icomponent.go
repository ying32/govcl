//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

type IComponent interface {
	IObject

	FindComponent(string) *TComponent
	GetNamePath() string
	HasParent() bool
	Assign(IObject)

	ComponentCount() int32
	ComponentIndex() int32
	SetComponentIndex(int32)
	Components(int32) *TComponent

	Owner() *TComponent

	Name() string
	SetName(string)

	Tag() int
	SetTag(int)
}
