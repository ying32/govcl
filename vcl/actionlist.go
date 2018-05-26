
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

type TActionList struct {
    IComponent
    instance uintptr
}

func NewActionList(owner IComponent) *TActionList {
    a := new(TActionList)
    a.instance = ActionList_Create(CheckPtr(owner))
    return a
}

func ActionListFromInst(inst uintptr) *TActionList {
    a := new(TActionList)
    a.instance = inst
    return a
}

func ActionListFromObj(obj IObject) *TActionList {
    a := new(TActionList)
    a.instance = CheckPtr(obj)
    return a
}

func (a *TActionList) Free() {
    if a.instance != 0 {
        ActionList_Free(a.instance)
        a.instance = 0
    }
}

func (a *TActionList) Instance() uintptr {
    return a.instance
}

func (a *TActionList) IsValid() bool {
    return a.instance != 0
}

func (a *TActionList) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ActionList_FindComponent(a.instance, AName))
}

func (a *TActionList) GetNamePath() string {
    return ActionList_GetNamePath(a.instance)
}

func (a *TActionList) HasParent() bool {
    return ActionList_HasParent(a.instance)
}

func (a *TActionList) Assign(Source IObject) {
    ActionList_Assign(a.instance, CheckPtr(Source))
}

func (a *TActionList) ClassName() string {
    return ActionList_ClassName(a.instance)
}

func (a *TActionList) Equals(Obj IObject) bool {
    return ActionList_Equals(a.instance, CheckPtr(Obj))
}

func (a *TActionList) GetHashCode() int32 {
    return ActionList_GetHashCode(a.instance)
}

func (a *TActionList) ToString() string {
    return ActionList_ToString(a.instance)
}

func (a *TActionList) Images() *TImageList {
    return ImageListFromInst(ActionList_GetImages(a.instance))
}

func (a *TActionList) SetImages(value IComponent) {
    ActionList_SetImages(a.instance, CheckPtr(value))
}

func (a *TActionList) State() TActionListState {
    return ActionList_GetState(a.instance)
}

func (a *TActionList) SetState(value TActionListState) {
    ActionList_SetState(a.instance, value)
}

func (a *TActionList) SetOnChange(fn TNotifyEvent) {
    ActionList_SetOnChange(a.instance, fn)
}

func (a *TActionList) ComponentCount() int32 {
    return ActionList_GetComponentCount(a.instance)
}

func (a *TActionList) ComponentIndex() int32 {
    return ActionList_GetComponentIndex(a.instance)
}

func (a *TActionList) SetComponentIndex(value int32) {
    ActionList_SetComponentIndex(a.instance, value)
}

func (a *TActionList) Owner() *TComponent {
    return ComponentFromInst(ActionList_GetOwner(a.instance))
}

func (a *TActionList) Name() string {
    return ActionList_GetName(a.instance)
}

func (a *TActionList) SetName(value string) {
    ActionList_SetName(a.instance, value)
}

func (a *TActionList) Tag() int {
    return ActionList_GetTag(a.instance)
}

func (a *TActionList) SetTag(value int) {
    ActionList_SetTag(a.instance, value)
}

func (a *TActionList) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ActionList_GetComponents(a.instance, AIndex))
}

