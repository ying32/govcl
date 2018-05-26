
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

type TOpenTextFileDialog struct {
    IComponent
    instance uintptr
}

func NewOpenTextFileDialog(owner IComponent) *TOpenTextFileDialog {
    o := new(TOpenTextFileDialog)
    o.instance = OpenTextFileDialog_Create(CheckPtr(owner))
    return o
}

func OpenTextFileDialogFromInst(inst uintptr) *TOpenTextFileDialog {
    o := new(TOpenTextFileDialog)
    o.instance = inst
    return o
}

func OpenTextFileDialogFromObj(obj IObject) *TOpenTextFileDialog {
    o := new(TOpenTextFileDialog)
    o.instance = CheckPtr(obj)
    return o
}

func (o *TOpenTextFileDialog) Free() {
    if o.instance != 0 {
        OpenTextFileDialog_Free(o.instance)
        o.instance = 0
    }
}

func (o *TOpenTextFileDialog) Instance() uintptr {
    return o.instance
}

func (o *TOpenTextFileDialog) IsValid() bool {
    return o.instance != 0
}

func (o *TOpenTextFileDialog) Execute() bool {
    return OpenTextFileDialog_Execute(o.instance)
}

func (o *TOpenTextFileDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(OpenTextFileDialog_FindComponent(o.instance, AName))
}

func (o *TOpenTextFileDialog) GetNamePath() string {
    return OpenTextFileDialog_GetNamePath(o.instance)
}

func (o *TOpenTextFileDialog) HasParent() bool {
    return OpenTextFileDialog_HasParent(o.instance)
}

func (o *TOpenTextFileDialog) Assign(Source IObject) {
    OpenTextFileDialog_Assign(o.instance, CheckPtr(Source))
}

func (o *TOpenTextFileDialog) ClassName() string {
    return OpenTextFileDialog_ClassName(o.instance)
}

func (o *TOpenTextFileDialog) Equals(Obj IObject) bool {
    return OpenTextFileDialog_Equals(o.instance, CheckPtr(Obj))
}

func (o *TOpenTextFileDialog) GetHashCode() int32 {
    return OpenTextFileDialog_GetHashCode(o.instance)
}

func (o *TOpenTextFileDialog) ToString() string {
    return OpenTextFileDialog_ToString(o.instance)
}

func (o *TOpenTextFileDialog) Files() *TStrings {
    return StringsFromInst(OpenTextFileDialog_GetFiles(o.instance))
}

func (o *TOpenTextFileDialog) DefaultExt() string {
    return OpenTextFileDialog_GetDefaultExt(o.instance)
}

func (o *TOpenTextFileDialog) SetDefaultExt(value string) {
    OpenTextFileDialog_SetDefaultExt(o.instance, value)
}

func (o *TOpenTextFileDialog) FileName() string {
    return OpenTextFileDialog_GetFileName(o.instance)
}

func (o *TOpenTextFileDialog) SetFileName(value string) {
    OpenTextFileDialog_SetFileName(o.instance, value)
}

func (o *TOpenTextFileDialog) Filter() string {
    return OpenTextFileDialog_GetFilter(o.instance)
}

func (o *TOpenTextFileDialog) SetFilter(value string) {
    OpenTextFileDialog_SetFilter(o.instance, value)
}

func (o *TOpenTextFileDialog) FilterIndex() int32 {
    return OpenTextFileDialog_GetFilterIndex(o.instance)
}

func (o *TOpenTextFileDialog) SetFilterIndex(value int32) {
    OpenTextFileDialog_SetFilterIndex(o.instance, value)
}

func (o *TOpenTextFileDialog) InitialDir() string {
    return OpenTextFileDialog_GetInitialDir(o.instance)
}

func (o *TOpenTextFileDialog) SetInitialDir(value string) {
    OpenTextFileDialog_SetInitialDir(o.instance, value)
}

func (o *TOpenTextFileDialog) Options() TOpenOptions {
    return OpenTextFileDialog_GetOptions(o.instance)
}

func (o *TOpenTextFileDialog) SetOptions(value TOpenOptions) {
    OpenTextFileDialog_SetOptions(o.instance, value)
}

func (o *TOpenTextFileDialog) OptionsEx() TOpenOptionsEx {
    return OpenTextFileDialog_GetOptionsEx(o.instance)
}

func (o *TOpenTextFileDialog) SetOptionsEx(value TOpenOptionsEx) {
    OpenTextFileDialog_SetOptionsEx(o.instance, value)
}

func (o *TOpenTextFileDialog) Title() string {
    return OpenTextFileDialog_GetTitle(o.instance)
}

func (o *TOpenTextFileDialog) SetTitle(value string) {
    OpenTextFileDialog_SetTitle(o.instance, value)
}

func (o *TOpenTextFileDialog) Handle() HWND {
    return OpenTextFileDialog_GetHandle(o.instance)
}

func (o *TOpenTextFileDialog) SetOnClose(fn TNotifyEvent) {
    OpenTextFileDialog_SetOnClose(o.instance, fn)
}

func (o *TOpenTextFileDialog) SetOnShow(fn TNotifyEvent) {
    OpenTextFileDialog_SetOnShow(o.instance, fn)
}

func (o *TOpenTextFileDialog) ComponentCount() int32 {
    return OpenTextFileDialog_GetComponentCount(o.instance)
}

func (o *TOpenTextFileDialog) ComponentIndex() int32 {
    return OpenTextFileDialog_GetComponentIndex(o.instance)
}

func (o *TOpenTextFileDialog) SetComponentIndex(value int32) {
    OpenTextFileDialog_SetComponentIndex(o.instance, value)
}

func (o *TOpenTextFileDialog) Owner() *TComponent {
    return ComponentFromInst(OpenTextFileDialog_GetOwner(o.instance))
}

func (o *TOpenTextFileDialog) Name() string {
    return OpenTextFileDialog_GetName(o.instance)
}

func (o *TOpenTextFileDialog) SetName(value string) {
    OpenTextFileDialog_SetName(o.instance, value)
}

func (o *TOpenTextFileDialog) Tag() int {
    return OpenTextFileDialog_GetTag(o.instance)
}

func (o *TOpenTextFileDialog) SetTag(value int) {
    OpenTextFileDialog_SetTag(o.instance, value)
}

func (o *TOpenTextFileDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(OpenTextFileDialog_GetComponents(o.instance, AIndex))
}

