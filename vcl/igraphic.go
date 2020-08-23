//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

type IGraphic interface {
	IObject
	LoadFromFile(string)
	SaveToFile(string)
	LoadFromStream(IStream)
	SaveToStream(IStream)
	Assign(IObject)
	Empty() bool
	Height() int32
	SetHeight(value int32)
	Modified() bool
	SetModified(value bool)
	Transparent() bool
	SetTransparent(value bool)
	Width() int32
	SetWidth(value int32)
}
