
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TObject struct {
    IObject
    instance uintptr
}

func NewObject() *TObject {
    o := new(TObject)
    o.instance = Object_Create()
    return o
}

func ObjectFromInst(inst uintptr) *TObject {
    o := new(TObject)
    o.instance = inst
    return o
}

func ObjectFromObj(obj IObject) *TObject {
    o := new(TObject)
    o.instance = CheckPtr(obj)
    return o
}

func (o *TObject) Free() {
    if o.instance != 0 {
        Object_Free(o.instance)
        o.instance = 0
    }
}

func (o *TObject) Instance() uintptr {
    return o.instance
}

func (o *TObject) IsValid() bool {
    return o.instance != 0
}

func (o *TObject) ClassName() string {
    return Object_ClassName(o.instance)
}

func (o *TObject) Equals(Obj IObject) bool {
    return Object_Equals(o.instance, CheckPtr(Obj))
}

func (o *TObject) GetHashCode() int32 {
    return Object_GetHashCode(o.instance)
}

func (o *TObject) ToString() string {
    return Object_ToString(o.instance)
}

