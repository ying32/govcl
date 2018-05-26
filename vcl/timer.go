
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TTimer struct {
    IComponent
    instance uintptr
}

func NewTimer(owner IComponent) *TTimer {
    t := new(TTimer)
    t.instance = Timer_Create(CheckPtr(owner))
    return t
}

func TimerFromInst(inst uintptr) *TTimer {
    t := new(TTimer)
    t.instance = inst
    return t
}

func TimerFromObj(obj IObject) *TTimer {
    t := new(TTimer)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTimer) Free() {
    if t.instance != 0 {
        Timer_Free(t.instance)
        t.instance = 0
    }
}

func (t *TTimer) Instance() uintptr {
    return t.instance
}

func (t *TTimer) IsValid() bool {
    return t.instance != 0
}

func (t *TTimer) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Timer_FindComponent(t.instance, AName))
}

func (t *TTimer) GetNamePath() string {
    return Timer_GetNamePath(t.instance)
}

func (t *TTimer) HasParent() bool {
    return Timer_HasParent(t.instance)
}

func (t *TTimer) Assign(Source IObject) {
    Timer_Assign(t.instance, CheckPtr(Source))
}

func (t *TTimer) ClassName() string {
    return Timer_ClassName(t.instance)
}

func (t *TTimer) Equals(Obj IObject) bool {
    return Timer_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTimer) GetHashCode() int32 {
    return Timer_GetHashCode(t.instance)
}

func (t *TTimer) ToString() string {
    return Timer_ToString(t.instance)
}

func (t *TTimer) Enabled() bool {
    return Timer_GetEnabled(t.instance)
}

func (t *TTimer) SetEnabled(value bool) {
    Timer_SetEnabled(t.instance, value)
}

func (t *TTimer) Interval() uint32 {
    return Timer_GetInterval(t.instance)
}

func (t *TTimer) SetInterval(value uint32) {
    Timer_SetInterval(t.instance, value)
}

func (t *TTimer) SetOnTimer(fn TNotifyEvent) {
    Timer_SetOnTimer(t.instance, fn)
}

func (t *TTimer) ComponentCount() int32 {
    return Timer_GetComponentCount(t.instance)
}

func (t *TTimer) ComponentIndex() int32 {
    return Timer_GetComponentIndex(t.instance)
}

func (t *TTimer) SetComponentIndex(value int32) {
    Timer_SetComponentIndex(t.instance, value)
}

func (t *TTimer) Owner() *TComponent {
    return ComponentFromInst(Timer_GetOwner(t.instance))
}

func (t *TTimer) Name() string {
    return Timer_GetName(t.instance)
}

func (t *TTimer) SetName(value string) {
    Timer_SetName(t.instance, value)
}

func (t *TTimer) Tag() int {
    return Timer_GetTag(t.instance)
}

func (t *TTimer) SetTag(value int) {
    Timer_SetTag(t.instance, value)
}

func (t *TTimer) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Timer_GetComponents(t.instance, AIndex))
}

