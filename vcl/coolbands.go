
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
)

type TCoolBands struct {
    IObject
    instance uintptr
}

func NewCoolBands() *TCoolBands {
    c := new(TCoolBands)
    c.instance = CoolBands_Create()
    return c
}

func CoolBandsFromInst(inst uintptr) *TCoolBands {
    c := new(TCoolBands)
    c.instance = inst
    return c
}

func CoolBandsFromObj(obj IObject) *TCoolBands {
    c := new(TCoolBands)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCoolBands) Free() {
    if c.instance != 0 {
        CoolBands_Free(c.instance)
        c.instance = 0
    }
}

func (c *TCoolBands) Instance() uintptr {
    return c.instance
}

func (c *TCoolBands) IsValid() bool {
    return c.instance != 0
}

func TCoolBandsClass() TClass {
    return CoolBands_StaticClassType()
}

func (c *TCoolBands) Add() *TCoolBand {
    return CoolBandFromInst(CoolBands_Add(c.instance))
}

func (c *TCoolBands) FindBand(AControl IControl) *TCoolBand {
    return CoolBandFromInst(CoolBands_FindBand(c.instance, CheckPtr(AControl)))
}

func (c *TCoolBands) Owner() *TObject {
    return ObjectFromInst(CoolBands_Owner(c.instance))
}

func (c *TCoolBands) Assign(Source IObject) {
    CoolBands_Assign(c.instance, CheckPtr(Source))
}

func (c *TCoolBands) BeginUpdate() {
    CoolBands_BeginUpdate(c.instance)
}

func (c *TCoolBands) Clear() {
    CoolBands_Clear(c.instance)
}

func (c *TCoolBands) Delete(Index int32) {
    CoolBands_Delete(c.instance, Index)
}

func (c *TCoolBands) EndUpdate() {
    CoolBands_EndUpdate(c.instance)
}

func (c *TCoolBands) GetNamePath() string {
    return CoolBands_GetNamePath(c.instance)
}

func (c *TCoolBands) Insert(Index int32) *TCollectionItem {
    return CollectionItemFromInst(CoolBands_Insert(c.instance, Index))
}

func (c *TCoolBands) DisposeOf() {
    CoolBands_DisposeOf(c.instance)
}

func (c *TCoolBands) ClassType() TClass {
    return CoolBands_ClassType(c.instance)
}

func (c *TCoolBands) ClassName() string {
    return CoolBands_ClassName(c.instance)
}

func (c *TCoolBands) InstanceSize() int32 {
    return CoolBands_InstanceSize(c.instance)
}

func (c *TCoolBands) InheritsFrom(AClass TClass) bool {
    return CoolBands_InheritsFrom(c.instance, AClass)
}

func (c *TCoolBands) Equals(Obj IObject) bool {
    return CoolBands_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCoolBands) GetHashCode() int32 {
    return CoolBands_GetHashCode(c.instance)
}

func (c *TCoolBands) ToString() string {
    return CoolBands_ToString(c.instance)
}

func (c *TCoolBands) CoolBar() *TCoolBar {
    return CoolBarFromInst(CoolBands_GetCoolBar(c.instance))
}

func (c *TCoolBands) Capacity() int32 {
    return CoolBands_GetCapacity(c.instance)
}

func (c *TCoolBands) SetCapacity(value int32) {
    CoolBands_SetCapacity(c.instance, value)
}

func (c *TCoolBands) Count() int32 {
    return CoolBands_GetCount(c.instance)
}

func (c *TCoolBands) Items(Index int32) *TCoolBand {
    return CoolBandFromInst(CoolBands_GetItems(c.instance, Index))
}

func (c *TCoolBands) SetItems(Index int32, value *TCoolBand) {
    CoolBands_SetItems(c.instance, Index, CheckPtr(value))
}

