
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

type TListGroups struct {
    IObject
    instance uintptr
}

func NewListGroups() *TListGroups {
    l := new(TListGroups)
    l.instance = ListGroups_Create()
    return l
}

func ListGroupsFromInst(inst uintptr) *TListGroups {
    l := new(TListGroups)
    l.instance = inst
    return l
}

func ListGroupsFromObj(obj IObject) *TListGroups {
    l := new(TListGroups)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListGroups) Free() {
    if l.instance != 0 {
        ListGroups_Free(l.instance)
        l.instance = 0
    }
}

func (l *TListGroups) Instance() uintptr {
    return l.instance
}

func (l *TListGroups) IsValid() bool {
    return l.instance != 0
}

func TListGroupsClass() TClass {
    return ListGroups_StaticClassType()
}

func (l *TListGroups) Add() *TListGroup {
    return ListGroupFromInst(ListGroups_Add(l.instance))
}

func (l *TListGroups) Owner() *TWinControl {
    return WinControlFromInst(ListGroups_Owner(l.instance))
}

func (l *TListGroups) Assign(Source IObject) {
    ListGroups_Assign(l.instance, CheckPtr(Source))
}

func (l *TListGroups) BeginUpdate() {
    ListGroups_BeginUpdate(l.instance)
}

func (l *TListGroups) Clear() {
    ListGroups_Clear(l.instance)
}

func (l *TListGroups) Delete(Index int32) {
    ListGroups_Delete(l.instance, Index)
}

func (l *TListGroups) EndUpdate() {
    ListGroups_EndUpdate(l.instance)
}

func (l *TListGroups) GetNamePath() string {
    return ListGroups_GetNamePath(l.instance)
}

func (l *TListGroups) Insert(Index int32) *TCollectionItem {
    return CollectionItemFromInst(ListGroups_Insert(l.instance, Index))
}

func (l *TListGroups) DisposeOf() {
    ListGroups_DisposeOf(l.instance)
}

func (l *TListGroups) ClassType() TClass {
    return ListGroups_ClassType(l.instance)
}

func (l *TListGroups) ClassName() string {
    return ListGroups_ClassName(l.instance)
}

func (l *TListGroups) InstanceSize() int32 {
    return ListGroups_InstanceSize(l.instance)
}

func (l *TListGroups) InheritsFrom(AClass TClass) bool {
    return ListGroups_InheritsFrom(l.instance, AClass)
}

func (l *TListGroups) Equals(Obj IObject) bool {
    return ListGroups_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListGroups) GetHashCode() int32 {
    return ListGroups_GetHashCode(l.instance)
}

func (l *TListGroups) ToString() string {
    return ListGroups_ToString(l.instance)
}

func (l *TListGroups) Capacity() int32 {
    return ListGroups_GetCapacity(l.instance)
}

func (l *TListGroups) SetCapacity(value int32) {
    ListGroups_SetCapacity(l.instance, value)
}

func (l *TListGroups) Count() int32 {
    return ListGroups_GetCount(l.instance)
}

func (l *TListGroups) Items(Index int32) *TListGroup {
    return ListGroupFromInst(ListGroups_GetItems(l.instance, Index))
}

func (l *TListGroups) SetItems(Index int32, value *TListGroup) {
    ListGroups_SetItems(l.instance, Index, CheckPtr(value))
}

