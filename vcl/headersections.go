
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

type THeaderSections struct {
    IObject
    instance uintptr
}

func NewHeaderSections() *THeaderSections {
    h := new(THeaderSections)
    h.instance = HeaderSections_Create()
    return h
}

func HeaderSectionsFromInst(inst uintptr) *THeaderSections {
    h := new(THeaderSections)
    h.instance = inst
    return h
}

func HeaderSectionsFromObj(obj IObject) *THeaderSections {
    h := new(THeaderSections)
    h.instance = CheckPtr(obj)
    return h
}

func (h *THeaderSections) Free() {
    if h.instance != 0 {
        HeaderSections_Free(h.instance)
        h.instance = 0
    }
}

func (h *THeaderSections) Instance() uintptr {
    return h.instance
}

func (h *THeaderSections) IsValid() bool {
    return h.instance != 0
}

func THeaderSectionsClass() TClass {
    return HeaderSections_StaticClassType()
}

func (h *THeaderSections) Add() *THeaderSection {
    return HeaderSectionFromInst(HeaderSections_Add(h.instance))
}

func (h *THeaderSections) AddItem(Item *THeaderSection, Index int32) *THeaderSection {
    return HeaderSectionFromInst(HeaderSections_AddItem(h.instance, CheckPtr(Item), Index))
}

func (h *THeaderSections) Insert(Index int32) *THeaderSection {
    return HeaderSectionFromInst(HeaderSections_Insert(h.instance, Index))
}

func (h *THeaderSections) Owner() *TObject {
    return ObjectFromInst(HeaderSections_Owner(h.instance))
}

func (h *THeaderSections) Assign(Source IObject) {
    HeaderSections_Assign(h.instance, CheckPtr(Source))
}

func (h *THeaderSections) BeginUpdate() {
    HeaderSections_BeginUpdate(h.instance)
}

func (h *THeaderSections) Clear() {
    HeaderSections_Clear(h.instance)
}

func (h *THeaderSections) Delete(Index int32) {
    HeaderSections_Delete(h.instance, Index)
}

func (h *THeaderSections) EndUpdate() {
    HeaderSections_EndUpdate(h.instance)
}

func (h *THeaderSections) GetNamePath() string {
    return HeaderSections_GetNamePath(h.instance)
}

func (h *THeaderSections) DisposeOf() {
    HeaderSections_DisposeOf(h.instance)
}

func (h *THeaderSections) ClassType() TClass {
    return HeaderSections_ClassType(h.instance)
}

func (h *THeaderSections) ClassName() string {
    return HeaderSections_ClassName(h.instance)
}

func (h *THeaderSections) InstanceSize() int32 {
    return HeaderSections_InstanceSize(h.instance)
}

func (h *THeaderSections) InheritsFrom(AClass TClass) bool {
    return HeaderSections_InheritsFrom(h.instance, AClass)
}

func (h *THeaderSections) Equals(Obj IObject) bool {
    return HeaderSections_Equals(h.instance, CheckPtr(Obj))
}

func (h *THeaderSections) GetHashCode() int32 {
    return HeaderSections_GetHashCode(h.instance)
}

func (h *THeaderSections) ToString() string {
    return HeaderSections_ToString(h.instance)
}

func (h *THeaderSections) Capacity() int32 {
    return HeaderSections_GetCapacity(h.instance)
}

func (h *THeaderSections) SetCapacity(value int32) {
    HeaderSections_SetCapacity(h.instance, value)
}

func (h *THeaderSections) Count() int32 {
    return HeaderSections_GetCount(h.instance)
}

func (h *THeaderSections) Items(Index int32) *THeaderSection {
    return HeaderSectionFromInst(HeaderSections_GetItems(h.instance, Index))
}

func (h *THeaderSections) SetItems(Index int32, value *THeaderSection) {
    HeaderSections_SetItems(h.instance, Index, CheckPtr(value))
}

