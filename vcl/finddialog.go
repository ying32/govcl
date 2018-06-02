
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

type TFindDialog struct {
    IComponent
    instance uintptr
}

func NewFindDialog(owner IComponent) *TFindDialog {
    f := new(TFindDialog)
    f.instance = FindDialog_Create(CheckPtr(owner))
    return f
}

func FindDialogFromInst(inst uintptr) *TFindDialog {
    f := new(TFindDialog)
    f.instance = inst
    return f
}

func FindDialogFromObj(obj IObject) *TFindDialog {
    f := new(TFindDialog)
    f.instance = CheckPtr(obj)
    return f
}

func (f *TFindDialog) Free() {
    if f.instance != 0 {
        FindDialog_Free(f.instance)
        f.instance = 0
    }
}

func (f *TFindDialog) Instance() uintptr {
    return f.instance
}

func (f *TFindDialog) IsValid() bool {
    return f.instance != 0
}

func TFindDialogClass() TClass {
    return FindDialog_StaticClassType()
}

func (f *TFindDialog) CloseDialog() {
    FindDialog_CloseDialog(f.instance)
}

func (f *TFindDialog) Execute() bool {
    return FindDialog_Execute(f.instance)
}

func (f *TFindDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(FindDialog_FindComponent(f.instance, AName))
}

func (f *TFindDialog) GetNamePath() string {
    return FindDialog_GetNamePath(f.instance)
}

func (f *TFindDialog) HasParent() bool {
    return FindDialog_HasParent(f.instance)
}

func (f *TFindDialog) Assign(Source IObject) {
    FindDialog_Assign(f.instance, CheckPtr(Source))
}

func (f *TFindDialog) DisposeOf() {
    FindDialog_DisposeOf(f.instance)
}

func (f *TFindDialog) ClassType() TClass {
    return FindDialog_ClassType(f.instance)
}

func (f *TFindDialog) ClassName() string {
    return FindDialog_ClassName(f.instance)
}

func (f *TFindDialog) InstanceSize() int32 {
    return FindDialog_InstanceSize(f.instance)
}

func (f *TFindDialog) InheritsFrom(AClass TClass) bool {
    return FindDialog_InheritsFrom(f.instance, AClass)
}

func (f *TFindDialog) Equals(Obj IObject) bool {
    return FindDialog_Equals(f.instance, CheckPtr(Obj))
}

func (f *TFindDialog) GetHashCode() int32 {
    return FindDialog_GetHashCode(f.instance)
}

func (f *TFindDialog) ToString() string {
    return FindDialog_ToString(f.instance)
}

func (f *TFindDialog) Left() int32 {
    return FindDialog_GetLeft(f.instance)
}

func (f *TFindDialog) SetLeft(value int32) {
    FindDialog_SetLeft(f.instance, value)
}

func (f *TFindDialog) Position() TPoint {
    return FindDialog_GetPosition(f.instance)
}

func (f *TFindDialog) SetPosition(value TPoint) {
    FindDialog_SetPosition(f.instance, value)
}

func (f *TFindDialog) Top() int32 {
    return FindDialog_GetTop(f.instance)
}

func (f *TFindDialog) SetTop(value int32) {
    FindDialog_SetTop(f.instance, value)
}

func (f *TFindDialog) FindText() string {
    return FindDialog_GetFindText(f.instance)
}

func (f *TFindDialog) SetFindText(value string) {
    FindDialog_SetFindText(f.instance, value)
}

func (f *TFindDialog) Options() TFindOptions {
    return FindDialog_GetOptions(f.instance)
}

func (f *TFindDialog) SetOptions(value TFindOptions) {
    FindDialog_SetOptions(f.instance, value)
}

func (f *TFindDialog) SetOnFind(fn TNotifyEvent) {
    FindDialog_SetOnFind(f.instance, fn)
}

func (f *TFindDialog) Handle() HWND {
    return FindDialog_GetHandle(f.instance)
}

func (f *TFindDialog) SetOnClose(fn TNotifyEvent) {
    FindDialog_SetOnClose(f.instance, fn)
}

func (f *TFindDialog) SetOnShow(fn TNotifyEvent) {
    FindDialog_SetOnShow(f.instance, fn)
}

func (f *TFindDialog) ComponentCount() int32 {
    return FindDialog_GetComponentCount(f.instance)
}

func (f *TFindDialog) ComponentIndex() int32 {
    return FindDialog_GetComponentIndex(f.instance)
}

func (f *TFindDialog) SetComponentIndex(value int32) {
    FindDialog_SetComponentIndex(f.instance, value)
}

func (f *TFindDialog) Owner() *TComponent {
    return ComponentFromInst(FindDialog_GetOwner(f.instance))
}

func (f *TFindDialog) Name() string {
    return FindDialog_GetName(f.instance)
}

func (f *TFindDialog) SetName(value string) {
    FindDialog_SetName(f.instance, value)
}

func (f *TFindDialog) Tag() int {
    return FindDialog_GetTag(f.instance)
}

func (f *TFindDialog) SetTag(value int) {
    FindDialog_SetTag(f.instance, value)
}

func (f *TFindDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(FindDialog_GetComponents(f.instance, AIndex))
}

