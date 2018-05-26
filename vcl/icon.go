
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

type TIcon struct {
    IObject
    instance uintptr
}

func NewIcon() *TIcon {
    i := new(TIcon)
    i.instance = Icon_Create()
    return i
}

func IconFromInst(inst uintptr) *TIcon {
    i := new(TIcon)
    i.instance = inst
    return i
}

func IconFromObj(obj IObject) *TIcon {
    i := new(TIcon)
    i.instance = CheckPtr(obj)
    return i
}

func (i *TIcon) Free() {
    if i.instance != 0 {
        Icon_Free(i.instance)
        i.instance = 0
    }
}

func (i *TIcon) Instance() uintptr {
    return i.instance
}

func (i *TIcon) IsValid() bool {
    return i.instance != 0
}

func (i *TIcon) Assign(Source IObject) {
    Icon_Assign(i.instance, CheckPtr(Source))
}

func (i *TIcon) HandleAllocated() bool {
    return Icon_HandleAllocated(i.instance)
}

func (i *TIcon) LoadFromStream(Stream IObject) {
    Icon_LoadFromStream(i.instance, CheckPtr(Stream))
}

func (i *TIcon) SaveToStream(Stream IObject) {
    Icon_SaveToStream(i.instance, CheckPtr(Stream))
}

func (i *TIcon) SetSize(AWidth int32, AHeight int32) {
    Icon_SetSize(i.instance, AWidth , AHeight)
}

func (i *TIcon) LoadFromResourceName(Instance uintptr, ResName string) {
    Icon_LoadFromResourceName(i.instance, Instance , ResName)
}

func (i *TIcon) LoadFromResourceID(Instance uintptr, ResID int32) {
    Icon_LoadFromResourceID(i.instance, Instance , ResID)
}

func (i *TIcon) Equals(Obj IObject) bool {
    return Icon_Equals(i.instance, CheckPtr(Obj))
}

func (i *TIcon) LoadFromFile(Filename string) {
    Icon_LoadFromFile(i.instance, Filename)
}

func (i *TIcon) SaveToFile(Filename string) {
    Icon_SaveToFile(i.instance, Filename)
}

func (i *TIcon) GetNamePath() string {
    return Icon_GetNamePath(i.instance)
}

func (i *TIcon) ClassName() string {
    return Icon_ClassName(i.instance)
}

func (i *TIcon) GetHashCode() int32 {
    return Icon_GetHashCode(i.instance)
}

func (i *TIcon) ToString() string {
    return Icon_ToString(i.instance)
}

func (i *TIcon) Handle() HICON {
    return Icon_GetHandle(i.instance)
}

func (i *TIcon) SetHandle(value HICON) {
    Icon_SetHandle(i.instance, value)
}

func (i *TIcon) Empty() bool {
    return Icon_GetEmpty(i.instance)
}

func (i *TIcon) Height() int32 {
    return Icon_GetHeight(i.instance)
}

func (i *TIcon) SetHeight(value int32) {
    Icon_SetHeight(i.instance, value)
}

func (i *TIcon) Modified() bool {
    return Icon_GetModified(i.instance)
}

func (i *TIcon) SetModified(value bool) {
    Icon_SetModified(i.instance, value)
}

func (i *TIcon) PaletteModified() bool {
    return Icon_GetPaletteModified(i.instance)
}

func (i *TIcon) SetPaletteModified(value bool) {
    Icon_SetPaletteModified(i.instance, value)
}

func (i *TIcon) Transparent() bool {
    return Icon_GetTransparent(i.instance)
}

func (i *TIcon) SetTransparent(value bool) {
    Icon_SetTransparent(i.instance, value)
}

func (i *TIcon) Width() int32 {
    return Icon_GetWidth(i.instance)
}

func (i *TIcon) SetWidth(value int32) {
    Icon_SetWidth(i.instance, value)
}

func (i *TIcon) SetOnChange(fn TNotifyEvent) {
    Icon_SetOnChange(i.instance, fn)
}

