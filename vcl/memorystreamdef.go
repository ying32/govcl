package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

// NewMemoryStreamFromBytes 新建Delphi内存流来自Go字节数组
func NewMemoryStreamFromBytes(data []byte) *TMemoryStream {
	m := NewMemoryStream()
	m.Write(data)
	return m
}

func (m *TMemoryStream) Read(count int32) (int32, []byte) {
	return MemoryStream_Read(m.instance, count)
}

func (m *TMemoryStream) Write(buffer []byte) int32 {
	return MemoryStream_Write(m.instance, buffer)
}
