
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

type TReplaceDialog struct {
    IComponent
    instance uintptr
}

func NewReplaceDialog(owner IComponent) *TReplaceDialog {
    r := new(TReplaceDialog)
    r.instance = ReplaceDialog_Create(CheckPtr(owner))
    return r
}

func ReplaceDialogFromInst(inst uintptr) *TReplaceDialog {
    r := new(TReplaceDialog)
    r.instance = inst
    return r
}

func ReplaceDialogFromObj(obj IObject) *TReplaceDialog {
    r := new(TReplaceDialog)
    r.instance = CheckPtr(obj)
    return r
}

func (r *TReplaceDialog) Free() {
    if r.instance != 0 {
        ReplaceDialog_Free(r.instance)
        r.instance = 0
    }
}

func (r *TReplaceDialog) Instance() uintptr {
    return r.instance
}

func (r *TReplaceDialog) IsValid() bool {
    return r.instance != 0
}

func (r *TReplaceDialog) CloseDialog() {
    ReplaceDialog_CloseDialog(r.instance)
}

func (r *TReplaceDialog) Execute() bool {
    return ReplaceDialog_Execute(r.instance)
}

func (r *TReplaceDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ReplaceDialog_FindComponent(r.instance, AName))
}

func (r *TReplaceDialog) GetNamePath() string {
    return ReplaceDialog_GetNamePath(r.instance)
}

func (r *TReplaceDialog) HasParent() bool {
    return ReplaceDialog_HasParent(r.instance)
}

func (r *TReplaceDialog) Assign(Source IObject) {
    ReplaceDialog_Assign(r.instance, CheckPtr(Source))
}

func (r *TReplaceDialog) ClassName() string {
    return ReplaceDialog_ClassName(r.instance)
}

func (r *TReplaceDialog) Equals(Obj IObject) bool {
    return ReplaceDialog_Equals(r.instance, CheckPtr(Obj))
}

func (r *TReplaceDialog) GetHashCode() int32 {
    return ReplaceDialog_GetHashCode(r.instance)
}

func (r *TReplaceDialog) ToString() string {
    return ReplaceDialog_ToString(r.instance)
}

func (r *TReplaceDialog) ReplaceText() string {
    return ReplaceDialog_GetReplaceText(r.instance)
}

func (r *TReplaceDialog) SetReplaceText(value string) {
    ReplaceDialog_SetReplaceText(r.instance, value)
}

func (r *TReplaceDialog) SetOnReplace(fn TNotifyEvent) {
    ReplaceDialog_SetOnReplace(r.instance, fn)
}

func (r *TReplaceDialog) Left() int32 {
    return ReplaceDialog_GetLeft(r.instance)
}

func (r *TReplaceDialog) SetLeft(value int32) {
    ReplaceDialog_SetLeft(r.instance, value)
}

func (r *TReplaceDialog) Position() TPoint {
    return ReplaceDialog_GetPosition(r.instance)
}

func (r *TReplaceDialog) SetPosition(value TPoint) {
    ReplaceDialog_SetPosition(r.instance, value)
}

func (r *TReplaceDialog) Top() int32 {
    return ReplaceDialog_GetTop(r.instance)
}

func (r *TReplaceDialog) SetTop(value int32) {
    ReplaceDialog_SetTop(r.instance, value)
}

func (r *TReplaceDialog) FindText() string {
    return ReplaceDialog_GetFindText(r.instance)
}

func (r *TReplaceDialog) SetFindText(value string) {
    ReplaceDialog_SetFindText(r.instance, value)
}

func (r *TReplaceDialog) Options() TFindOptions {
    return ReplaceDialog_GetOptions(r.instance)
}

func (r *TReplaceDialog) SetOptions(value TFindOptions) {
    ReplaceDialog_SetOptions(r.instance, value)
}

func (r *TReplaceDialog) SetOnFind(fn TNotifyEvent) {
    ReplaceDialog_SetOnFind(r.instance, fn)
}

func (r *TReplaceDialog) Handle() HWND {
    return ReplaceDialog_GetHandle(r.instance)
}

func (r *TReplaceDialog) SetOnClose(fn TNotifyEvent) {
    ReplaceDialog_SetOnClose(r.instance, fn)
}

func (r *TReplaceDialog) SetOnShow(fn TNotifyEvent) {
    ReplaceDialog_SetOnShow(r.instance, fn)
}

func (r *TReplaceDialog) ComponentCount() int32 {
    return ReplaceDialog_GetComponentCount(r.instance)
}

func (r *TReplaceDialog) ComponentIndex() int32 {
    return ReplaceDialog_GetComponentIndex(r.instance)
}

func (r *TReplaceDialog) SetComponentIndex(value int32) {
    ReplaceDialog_SetComponentIndex(r.instance, value)
}

func (r *TReplaceDialog) Owner() *TComponent {
    return ComponentFromInst(ReplaceDialog_GetOwner(r.instance))
}

func (r *TReplaceDialog) Name() string {
    return ReplaceDialog_GetName(r.instance)
}

func (r *TReplaceDialog) SetName(value string) {
    ReplaceDialog_SetName(r.instance, value)
}

func (r *TReplaceDialog) Tag() int {
    return ReplaceDialog_GetTag(r.instance)
}

func (r *TReplaceDialog) SetTag(value int) {
    ReplaceDialog_SetTag(r.instance, value)
}

func (r *TReplaceDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ReplaceDialog_GetComponents(r.instance, AIndex))
}

