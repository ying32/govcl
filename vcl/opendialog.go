
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

type TOpenDialog struct {
    IComponent
    instance uintptr
}

func NewOpenDialog(owner IComponent) *TOpenDialog {
    o := new(TOpenDialog)
    o.instance = OpenDialog_Create(CheckPtr(owner))
    return o
}

func OpenDialogFromInst(inst uintptr) *TOpenDialog {
    o := new(TOpenDialog)
    o.instance = inst
    return o
}

func OpenDialogFromObj(obj IObject) *TOpenDialog {
    o := new(TOpenDialog)
    o.instance = CheckPtr(obj)
    return o
}

func (o *TOpenDialog) Free() {
    if o.instance != 0 {
        OpenDialog_Free(o.instance)
        o.instance = 0
    }
}

func (o *TOpenDialog) Instance() uintptr {
    return o.instance
}

func (o *TOpenDialog) IsValid() bool {
    return o.instance != 0
}

func (o *TOpenDialog) Execute() bool {
    return OpenDialog_Execute(o.instance)
}

func (o *TOpenDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(OpenDialog_FindComponent(o.instance, AName))
}

func (o *TOpenDialog) GetNamePath() string {
    return OpenDialog_GetNamePath(o.instance)
}

func (o *TOpenDialog) HasParent() bool {
    return OpenDialog_HasParent(o.instance)
}

func (o *TOpenDialog) Assign(Source IObject) {
    OpenDialog_Assign(o.instance, CheckPtr(Source))
}

func (o *TOpenDialog) ClassName() string {
    return OpenDialog_ClassName(o.instance)
}

func (o *TOpenDialog) Equals(Obj IObject) bool {
    return OpenDialog_Equals(o.instance, CheckPtr(Obj))
}

func (o *TOpenDialog) GetHashCode() int32 {
    return OpenDialog_GetHashCode(o.instance)
}

func (o *TOpenDialog) ToString() string {
    return OpenDialog_ToString(o.instance)
}

func (o *TOpenDialog) Files() *TStrings {
    return StringsFromInst(OpenDialog_GetFiles(o.instance))
}

func (o *TOpenDialog) DefaultExt() string {
    return OpenDialog_GetDefaultExt(o.instance)
}

func (o *TOpenDialog) SetDefaultExt(value string) {
    OpenDialog_SetDefaultExt(o.instance, value)
}

func (o *TOpenDialog) FileName() string {
    return OpenDialog_GetFileName(o.instance)
}

func (o *TOpenDialog) SetFileName(value string) {
    OpenDialog_SetFileName(o.instance, value)
}

func (o *TOpenDialog) Filter() string {
    return OpenDialog_GetFilter(o.instance)
}

func (o *TOpenDialog) SetFilter(value string) {
    OpenDialog_SetFilter(o.instance, value)
}

func (o *TOpenDialog) FilterIndex() int32 {
    return OpenDialog_GetFilterIndex(o.instance)
}

func (o *TOpenDialog) SetFilterIndex(value int32) {
    OpenDialog_SetFilterIndex(o.instance, value)
}

func (o *TOpenDialog) InitialDir() string {
    return OpenDialog_GetInitialDir(o.instance)
}

func (o *TOpenDialog) SetInitialDir(value string) {
    OpenDialog_SetInitialDir(o.instance, value)
}

func (o *TOpenDialog) Options() TOpenOptions {
    return OpenDialog_GetOptions(o.instance)
}

func (o *TOpenDialog) SetOptions(value TOpenOptions) {
    OpenDialog_SetOptions(o.instance, value)
}

func (o *TOpenDialog) OptionsEx() TOpenOptionsEx {
    return OpenDialog_GetOptionsEx(o.instance)
}

func (o *TOpenDialog) SetOptionsEx(value TOpenOptionsEx) {
    OpenDialog_SetOptionsEx(o.instance, value)
}

func (o *TOpenDialog) Title() string {
    return OpenDialog_GetTitle(o.instance)
}

func (o *TOpenDialog) SetTitle(value string) {
    OpenDialog_SetTitle(o.instance, value)
}

func (o *TOpenDialog) Handle() HWND {
    return OpenDialog_GetHandle(o.instance)
}

func (o *TOpenDialog) SetOnClose(fn TNotifyEvent) {
    OpenDialog_SetOnClose(o.instance, fn)
}

func (o *TOpenDialog) SetOnShow(fn TNotifyEvent) {
    OpenDialog_SetOnShow(o.instance, fn)
}

func (o *TOpenDialog) ComponentCount() int32 {
    return OpenDialog_GetComponentCount(o.instance)
}

func (o *TOpenDialog) ComponentIndex() int32 {
    return OpenDialog_GetComponentIndex(o.instance)
}

func (o *TOpenDialog) SetComponentIndex(value int32) {
    OpenDialog_SetComponentIndex(o.instance, value)
}

func (o *TOpenDialog) Owner() *TComponent {
    return ComponentFromInst(OpenDialog_GetOwner(o.instance))
}

func (o *TOpenDialog) Name() string {
    return OpenDialog_GetName(o.instance)
}

func (o *TOpenDialog) SetName(value string) {
    OpenDialog_SetName(o.instance, value)
}

func (o *TOpenDialog) Tag() int {
    return OpenDialog_GetTag(o.instance)
}

func (o *TOpenDialog) SetTag(value int) {
    OpenDialog_SetTag(o.instance, value)
}

func (o *TOpenDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(OpenDialog_GetComponents(o.instance, AIndex))
}

