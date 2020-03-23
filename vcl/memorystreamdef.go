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

// CN: 新建Delphi/Lazarus内存流来自Go字节数组。
// EN: New Delphi/Lazarus memory stream from Go byte array.
func NewMemoryStreamFromBytes(data []byte) *TMemoryStream {
	m := NewMemoryStream()
	m.Write(data)
	return m
}

// CN: 读数据
// EN: Read Data.
func (m *TMemoryStream) Read(count int32) (int32, []byte) {
	return MemoryStream_Read(m.instance, count)
}

// CN: 写数据
// EN: Write Data.
func (m *TMemoryStream) Write(buffer []byte) int32 {
	return MemoryStream_Write(m.instance, buffer)
}
