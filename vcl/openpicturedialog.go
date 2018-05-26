
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

type TOpenPictureDialog struct {
    IComponent
    instance uintptr
}

func NewOpenPictureDialog(owner IComponent) *TOpenPictureDialog {
    o := new(TOpenPictureDialog)
    o.instance = OpenPictureDialog_Create(CheckPtr(owner))
    return o
}

func OpenPictureDialogFromInst(inst uintptr) *TOpenPictureDialog {
    o := new(TOpenPictureDialog)
    o.instance = inst
    return o
}

func OpenPictureDialogFromObj(obj IObject) *TOpenPictureDialog {
    o := new(TOpenPictureDialog)
    o.instance = CheckPtr(obj)
    return o
}

func (o *TOpenPictureDialog) Free() {
    if o.instance != 0 {
        OpenPictureDialog_Free(o.instance)
        o.instance = 0
    }
}

func (o *TOpenPictureDialog) Instance() uintptr {
    return o.instance
}

func (o *TOpenPictureDialog) IsValid() bool {
    return o.instance != 0
}

func (o *TOpenPictureDialog) Execute() bool {
    return OpenPictureDialog_Execute(o.instance)
}

func (o *TOpenPictureDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(OpenPictureDialog_FindComponent(o.instance, AName))
}

func (o *TOpenPictureDialog) GetNamePath() string {
    return OpenPictureDialog_GetNamePath(o.instance)
}

func (o *TOpenPictureDialog) HasParent() bool {
    return OpenPictureDialog_HasParent(o.instance)
}

func (o *TOpenPictureDialog) Assign(Source IObject) {
    OpenPictureDialog_Assign(o.instance, CheckPtr(Source))
}

func (o *TOpenPictureDialog) ClassName() string {
    return OpenPictureDialog_ClassName(o.instance)
}

func (o *TOpenPictureDialog) Equals(Obj IObject) bool {
    return OpenPictureDialog_Equals(o.instance, CheckPtr(Obj))
}

func (o *TOpenPictureDialog) GetHashCode() int32 {
    return OpenPictureDialog_GetHashCode(o.instance)
}

func (o *TOpenPictureDialog) ToString() string {
    return OpenPictureDialog_ToString(o.instance)
}

func (o *TOpenPictureDialog) Filter() string {
    return OpenPictureDialog_GetFilter(o.instance)
}

func (o *TOpenPictureDialog) SetFilter(value string) {
    OpenPictureDialog_SetFilter(o.instance, value)
}

func (o *TOpenPictureDialog) Files() *TStrings {
    return StringsFromInst(OpenPictureDialog_GetFiles(o.instance))
}

func (o *TOpenPictureDialog) DefaultExt() string {
    return OpenPictureDialog_GetDefaultExt(o.instance)
}

func (o *TOpenPictureDialog) SetDefaultExt(value string) {
    OpenPictureDialog_SetDefaultExt(o.instance, value)
}

func (o *TOpenPictureDialog) FileName() string {
    return OpenPictureDialog_GetFileName(o.instance)
}

func (o *TOpenPictureDialog) SetFileName(value string) {
    OpenPictureDialog_SetFileName(o.instance, value)
}

func (o *TOpenPictureDialog) FilterIndex() int32 {
    return OpenPictureDialog_GetFilterIndex(o.instance)
}

func (o *TOpenPictureDialog) SetFilterIndex(value int32) {
    OpenPictureDialog_SetFilterIndex(o.instance, value)
}

func (o *TOpenPictureDialog) InitialDir() string {
    return OpenPictureDialog_GetInitialDir(o.instance)
}

func (o *TOpenPictureDialog) SetInitialDir(value string) {
    OpenPictureDialog_SetInitialDir(o.instance, value)
}

func (o *TOpenPictureDialog) Options() TOpenOptions {
    return OpenPictureDialog_GetOptions(o.instance)
}

func (o *TOpenPictureDialog) SetOptions(value TOpenOptions) {
    OpenPictureDialog_SetOptions(o.instance, value)
}

func (o *TOpenPictureDialog) OptionsEx() TOpenOptionsEx {
    return OpenPictureDialog_GetOptionsEx(o.instance)
}

func (o *TOpenPictureDialog) SetOptionsEx(value TOpenOptionsEx) {
    OpenPictureDialog_SetOptionsEx(o.instance, value)
}

func (o *TOpenPictureDialog) Title() string {
    return OpenPictureDialog_GetTitle(o.instance)
}

func (o *TOpenPictureDialog) SetTitle(value string) {
    OpenPictureDialog_SetTitle(o.instance, value)
}

func (o *TOpenPictureDialog) Handle() HWND {
    return OpenPictureDialog_GetHandle(o.instance)
}

func (o *TOpenPictureDialog) SetOnClose(fn TNotifyEvent) {
    OpenPictureDialog_SetOnClose(o.instance, fn)
}

func (o *TOpenPictureDialog) SetOnShow(fn TNotifyEvent) {
    OpenPictureDialog_SetOnShow(o.instance, fn)
}

func (o *TOpenPictureDialog) ComponentCount() int32 {
    return OpenPictureDialog_GetComponentCount(o.instance)
}

func (o *TOpenPictureDialog) ComponentIndex() int32 {
    return OpenPictureDialog_GetComponentIndex(o.instance)
}

func (o *TOpenPictureDialog) SetComponentIndex(value int32) {
    OpenPictureDialog_SetComponentIndex(o.instance, value)
}

func (o *TOpenPictureDialog) Owner() *TComponent {
    return ComponentFromInst(OpenPictureDialog_GetOwner(o.instance))
}

func (o *TOpenPictureDialog) Name() string {
    return OpenPictureDialog_GetName(o.instance)
}

func (o *TOpenPictureDialog) SetName(value string) {
    OpenPictureDialog_SetName(o.instance, value)
}

func (o *TOpenPictureDialog) Tag() int {
    return OpenPictureDialog_GetTag(o.instance)
}

func (o *TOpenPictureDialog) SetTag(value int) {
    OpenPictureDialog_SetTag(o.instance, value)
}

func (o *TOpenPictureDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(OpenPictureDialog_GetComponents(o.instance, AIndex))
}

