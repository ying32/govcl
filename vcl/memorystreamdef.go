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
)

// NewMemoryStreamFromBytes
//
// 新建FreePascal内存流来自Go字节数组。
//
// New FreePascal memory stream from Go byte array.
func NewMemoryStreamFromBytes(data []byte) *TMemoryStream {
	m := NewMemoryStream()
	m.Write(data)
	return m
}

// Read
//
// 读数据
//
// Read Data.
func (m *TMemoryStream) Read(count int32) (int32, []byte) {
	return MemoryStream_Read(m._instance(), count)
}

// Write
//
// 写数据
//
// Write Data.
func (m *TMemoryStream) Write(buffer []byte) int32 {
	return MemoryStream_Write(m._instance(), buffer)
}
