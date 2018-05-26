
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

type TListItems struct {
    IObject
    instance uintptr
}

func NewListItems() *TListItems {
    l := new(TListItems)
    l.instance = ListItems_Create()
    return l
}

func ListItemsFromInst(inst uintptr) *TListItems {
    l := new(TListItems)
    l.instance = inst
    return l
}

func ListItemsFromObj(obj IObject) *TListItems {
    l := new(TListItems)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListItems) Free() {
    if l.instance != 0 {
        ListItems_Free(l.instance)
        l.instance = 0
    }
}

func (l *TListItems) Instance() uintptr {
    return l.instance
}

func (l *TListItems) IsValid() bool {
    return l.instance != 0
}

func (l *TListItems) Add() *TListItem {
    return ListItemFromInst(ListItems_Add(l.instance))
}

func (l *TListItems) AddItem(Item *TListItem, Index int32) *TListItem {
    return ListItemFromInst(ListItems_AddItem(l.instance, CheckPtr(Item), Index))
}

func (l *TListItems) Assign(Source IObject) {
    ListItems_Assign(l.instance, CheckPtr(Source))
}

func (l *TListItems) BeginUpdate() {
    ListItems_BeginUpdate(l.instance)
}

func (l *TListItems) Clear() {
    ListItems_Clear(l.instance)
}

func (l *TListItems) Delete(Index int32) {
    ListItems_Delete(l.instance, Index)
}

func (l *TListItems) EndUpdate() {
    ListItems_EndUpdate(l.instance)
}

func (l *TListItems) IndexOf(Value *TListItem) int32 {
    return ListItems_IndexOf(l.instance, CheckPtr(Value))
}

func (l *TListItems) Insert(Index int32) *TListItem {
    return ListItemFromInst(ListItems_Insert(l.instance, Index))
}

func (l *TListItems) GetNamePath() string {
    return ListItems_GetNamePath(l.instance)
}

func (l *TListItems) ClassName() string {
    return ListItems_ClassName(l.instance)
}

func (l *TListItems) Equals(Obj IObject) bool {
    return ListItems_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListItems) GetHashCode() int32 {
    return ListItems_GetHashCode(l.instance)
}

func (l *TListItems) ToString() string {
    return ListItems_ToString(l.instance)
}

func (l *TListItems) Count() int32 {
    return ListItems_GetCount(l.instance)
}

func (l *TListItems) SetCount(value int32) {
    ListItems_SetCount(l.instance, value)
}

func (l *TListItems) Handle() HWND {
    return ListItems_GetHandle(l.instance)
}

func (l *TListItems) Owner() *TControl {
    return ControlFromInst(ListItems_GetOwner(l.instance))
}

func (l *TListItems) Item(Index int32) *TListItem {
    return ListItemFromInst(ListItems_GetItem(l.instance, Index))
}

func (l *TListItems) SetItem(Index int32, value *TListItem) {
    ListItems_SetItem(l.instance, Index, CheckPtr(value))
}

