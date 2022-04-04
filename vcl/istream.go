//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	. "github.com/ying32/govcl/vcl/types"
)

type IStream interface {
	IObject
	Size() int64
	SetSize(value int64)
	Position() int64
	SetPosition(int64)
	Read(count int32) (int32, []byte)
	Write(buffer []byte) int32
	Seek(Offset int64, Origin TSeekOrigin) int64
	CopyFrom(Source IStream, Count int64) int64
}
