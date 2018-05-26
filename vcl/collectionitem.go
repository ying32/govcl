
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TCollectionItem struct {
    IObject
    instance uintptr
}

func NewCollectionItem() *TCollectionItem {
    c := new(TCollectionItem)
    c.instance = CollectionItem_Create()
    return c
}

func CollectionItemFromInst(inst uintptr) *TCollectionItem {
    c := new(TCollectionItem)
    c.instance = inst
    return c
}

func CollectionItemFromObj(obj IObject) *TCollectionItem {
    c := new(TCollectionItem)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCollectionItem) Free() {
    if c.instance != 0 {
        CollectionItem_Free(c.instance)
        c.instance = 0
    }
}

func (c *TCollectionItem) Instance() uintptr {
    return c.instance
}

func (c *TCollectionItem) IsValid() bool {
    return c.instance != 0
}

func (c *TCollectionItem) GetNamePath() string {
    return CollectionItem_GetNamePath(c.instance)
}

func (c *TCollectionItem) Assign(Source IObject) {
    CollectionItem_Assign(c.instance, CheckPtr(Source))
}

func (c *TCollectionItem) ClassName() string {
    return CollectionItem_ClassName(c.instance)
}

func (c *TCollectionItem) Equals(Obj IObject) bool {
    return CollectionItem_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCollectionItem) GetHashCode() int32 {
    return CollectionItem_GetHashCode(c.instance)
}

func (c *TCollectionItem) ToString() string {
    return CollectionItem_ToString(c.instance)
}

func (c *TCollectionItem) Index() int32 {
    return CollectionItem_GetIndex(c.instance)
}

func (c *TCollectionItem) SetIndex(value int32) {
    CollectionItem_SetIndex(c.instance, value)
}

