
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TMemoryStream struct {
    IObject
    instance uintptr
}

func NewMemoryStream() *TMemoryStream {
    m := new(TMemoryStream)
    m.instance = MemoryStream_Create()
    return m
}

func MemoryStreamFromInst(inst uintptr) *TMemoryStream {
    m := new(TMemoryStream)
    m.instance = inst
    return m
}

func MemoryStreamFromObj(obj IObject) *TMemoryStream {
    m := new(TMemoryStream)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMemoryStream) Free() {
    if m.instance != 0 {
        MemoryStream_Free(m.instance)
        m.instance = 0
    }
}

func (m *TMemoryStream) Instance() uintptr {
    return m.instance
}

func (m *TMemoryStream) IsValid() bool {
    return m.instance != 0
}

func (m *TMemoryStream) Clear() {
    MemoryStream_Clear(m.instance)
}

func (m *TMemoryStream) LoadFromStream(Stream IObject) {
    MemoryStream_LoadFromStream(m.instance, CheckPtr(Stream))
}

func (m *TMemoryStream) LoadFromFile(FileName string) {
    MemoryStream_LoadFromFile(m.instance, FileName)
}

func (m *TMemoryStream) Seek(Offset int64, Origin TSeekOrigin) int64 {
    return MemoryStream_Seek(m.instance, Offset , Origin)
}

func (m *TMemoryStream) SaveToStream(Stream IObject) {
    MemoryStream_SaveToStream(m.instance, CheckPtr(Stream))
}

func (m *TMemoryStream) SaveToFile(FileName string) {
    MemoryStream_SaveToFile(m.instance, FileName)
}

func (m *TMemoryStream) CopyFrom(Source IObject, Count int64) int64 {
    return MemoryStream_CopyFrom(m.instance, CheckPtr(Source), Count)
}

func (m *TMemoryStream) ClassName() string {
    return MemoryStream_ClassName(m.instance)
}

func (m *TMemoryStream) Equals(Obj IObject) bool {
    return MemoryStream_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMemoryStream) GetHashCode() int32 {
    return MemoryStream_GetHashCode(m.instance)
}

func (m *TMemoryStream) ToString() string {
    return MemoryStream_ToString(m.instance)
}

func (m *TMemoryStream) Memory() uintptr {
    return MemoryStream_GetMemory(m.instance)
}

func (m *TMemoryStream) Position() int64 {
    return MemoryStream_GetPosition(m.instance)
}

func (m *TMemoryStream) SetPosition(value int64) {
    MemoryStream_SetPosition(m.instance, value)
}

func (m *TMemoryStream) Size() int64 {
    return MemoryStream_GetSize(m.instance)
}

func (m *TMemoryStream) SetSize(value int64) {
    MemoryStream_SetSize(m.instance, value)
}

