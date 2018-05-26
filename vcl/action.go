
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

type TAction struct {
    IComponent
    instance uintptr
}

func NewAction(owner IComponent) *TAction {
    a := new(TAction)
    a.instance = Action_Create(CheckPtr(owner))
    return a
}

func ActionFromInst(inst uintptr) *TAction {
    a := new(TAction)
    a.instance = inst
    return a
}

func ActionFromObj(obj IObject) *TAction {
    a := new(TAction)
    a.instance = CheckPtr(obj)
    return a
}

func (a *TAction) Free() {
    if a.instance != 0 {
        Action_Free(a.instance)
        a.instance = 0
    }
}

func (a *TAction) Instance() uintptr {
    return a.instance
}

func (a *TAction) IsValid() bool {
    return a.instance != 0
}

func (a *TAction) Execute() bool {
    return Action_Execute(a.instance)
}

func (a *TAction) Update() bool {
    return Action_Update(a.instance)
}

func (a *TAction) HasParent() bool {
    return Action_HasParent(a.instance)
}

func (a *TAction) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Action_FindComponent(a.instance, AName))
}

func (a *TAction) GetNamePath() string {
    return Action_GetNamePath(a.instance)
}

func (a *TAction) Assign(Source IObject) {
    Action_Assign(a.instance, CheckPtr(Source))
}

func (a *TAction) ClassName() string {
    return Action_ClassName(a.instance)
}

func (a *TAction) Equals(Obj IObject) bool {
    return Action_Equals(a.instance, CheckPtr(Obj))
}

func (a *TAction) GetHashCode() int32 {
    return Action_GetHashCode(a.instance)
}

func (a *TAction) ToString() string {
    return Action_ToString(a.instance)
}

func (a *TAction) Caption() string {
    return Action_GetCaption(a.instance)
}

func (a *TAction) SetCaption(value string) {
    Action_SetCaption(a.instance, value)
}

func (a *TAction) Checked() bool {
    return Action_GetChecked(a.instance)
}

func (a *TAction) SetChecked(value bool) {
    Action_SetChecked(a.instance, value)
}

func (a *TAction) Enabled() bool {
    return Action_GetEnabled(a.instance)
}

func (a *TAction) SetEnabled(value bool) {
    Action_SetEnabled(a.instance, value)
}

func (a *TAction) GroupIndex() int32 {
    return Action_GetGroupIndex(a.instance)
}

func (a *TAction) SetGroupIndex(value int32) {
    Action_SetGroupIndex(a.instance, value)
}

func (a *TAction) Hint() string {
    return Action_GetHint(a.instance)
}

func (a *TAction) SetHint(value string) {
    Action_SetHint(a.instance, value)
}

func (a *TAction) ImageIndex() int32 {
    return Action_GetImageIndex(a.instance)
}

func (a *TAction) SetImageIndex(value int32) {
    Action_SetImageIndex(a.instance, value)
}

func (a *TAction) ShortCut() TShortCut {
    return Action_GetShortCut(a.instance)
}

func (a *TAction) SetShortCut(value TShortCut) {
    Action_SetShortCut(a.instance, value)
}

func (a *TAction) Visible() bool {
    return Action_GetVisible(a.instance)
}

func (a *TAction) SetVisible(value bool) {
    Action_SetVisible(a.instance, value)
}

func (a *TAction) SetOnExecute(fn TNotifyEvent) {
    Action_SetOnExecute(a.instance, fn)
}

func (a *TAction) SetOnUpdate(fn TNotifyEvent) {
    Action_SetOnUpdate(a.instance, fn)
}

func (a *TAction) Images() *TImageList {
    return ImageListFromInst(Action_GetImages(a.instance))
}

func (a *TAction) Index() int32 {
    return Action_GetIndex(a.instance)
}

func (a *TAction) SetIndex(value int32) {
    Action_SetIndex(a.instance, value)
}

func (a *TAction) ComponentCount() int32 {
    return Action_GetComponentCount(a.instance)
}

func (a *TAction) ComponentIndex() int32 {
    return Action_GetComponentIndex(a.instance)
}

func (a *TAction) SetComponentIndex(value int32) {
    Action_SetComponentIndex(a.instance, value)
}

func (a *TAction) Owner() *TComponent {
    return ComponentFromInst(Action_GetOwner(a.instance))
}

func (a *TAction) Name() string {
    return Action_GetName(a.instance)
}

func (a *TAction) SetName(value string) {
    Action_SetName(a.instance, value)
}

func (a *TAction) Tag() int {
    return Action_GetTag(a.instance)
}

func (a *TAction) SetTag(value int) {
    Action_SetTag(a.instance, value)
}

func (a *TAction) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Action_GetComponents(a.instance, AIndex))
}

