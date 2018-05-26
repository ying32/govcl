
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

type TIconOptions struct {
    IObject
    instance uintptr
}

func IconOptionsFromInst(inst uintptr) *TIconOptions {
    i := new(TIconOptions)
    i.instance = inst
    return i
}

func IconOptionsFromObj(obj IObject) *TIconOptions {
    i := new(TIconOptions)
    i.instance = CheckPtr(obj)
    return i
}

func (i *TIconOptions) Instance() uintptr {
    return i.instance
}

func (i *TIconOptions) IsValid() bool {
    return i.instance != 0
}

func (i *TIconOptions) Assign(Source IObject) {
    IconOptions_Assign(i.instance, CheckPtr(Source))
}

func (i *TIconOptions) GetNamePath() string {
    return IconOptions_GetNamePath(i.instance)
}

func (i *TIconOptions) ClassName() string {
    return IconOptions_ClassName(i.instance)
}

func (i *TIconOptions) Equals(Obj IObject) bool {
    return IconOptions_Equals(i.instance, CheckPtr(Obj))
}

func (i *TIconOptions) GetHashCode() int32 {
    return IconOptions_GetHashCode(i.instance)
}

func (i *TIconOptions) ToString() string {
    return IconOptions_ToString(i.instance)
}

func (i *TIconOptions) Arrangement() TIconArrangement {
    return IconOptions_GetArrangement(i.instance)
}

func (i *TIconOptions) SetArrangement(value TIconArrangement) {
    IconOptions_SetArrangement(i.instance, value)
}

func (i *TIconOptions) AutoArrange() bool {
    return IconOptions_GetAutoArrange(i.instance)
}

func (i *TIconOptions) SetAutoArrange(value bool) {
    IconOptions_SetAutoArrange(i.instance, value)
}

