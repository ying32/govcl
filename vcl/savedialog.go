
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

type TSaveDialog struct {
    IComponent
    instance uintptr
}

func NewSaveDialog(owner IComponent) *TSaveDialog {
    s := new(TSaveDialog)
    s.instance = SaveDialog_Create(CheckPtr(owner))
    return s
}

func SaveDialogFromInst(inst uintptr) *TSaveDialog {
    s := new(TSaveDialog)
    s.instance = inst
    return s
}

func SaveDialogFromObj(obj IObject) *TSaveDialog {
    s := new(TSaveDialog)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TSaveDialog) Free() {
    if s.instance != 0 {
        SaveDialog_Free(s.instance)
        s.instance = 0
    }
}

func (s *TSaveDialog) Instance() uintptr {
    return s.instance
}

func (s *TSaveDialog) IsValid() bool {
    return s.instance != 0
}

func (s *TSaveDialog) Execute() bool {
    return SaveDialog_Execute(s.instance)
}

func (s *TSaveDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(SaveDialog_FindComponent(s.instance, AName))
}

func (s *TSaveDialog) GetNamePath() string {
    return SaveDialog_GetNamePath(s.instance)
}

func (s *TSaveDialog) HasParent() bool {
    return SaveDialog_HasParent(s.instance)
}

func (s *TSaveDialog) Assign(Source IObject) {
    SaveDialog_Assign(s.instance, CheckPtr(Source))
}

func (s *TSaveDialog) ClassName() string {
    return SaveDialog_ClassName(s.instance)
}

func (s *TSaveDialog) Equals(Obj IObject) bool {
    return SaveDialog_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSaveDialog) GetHashCode() int32 {
    return SaveDialog_GetHashCode(s.instance)
}

func (s *TSaveDialog) ToString() string {
    return SaveDialog_ToString(s.instance)
}

func (s *TSaveDialog) Files() *TStrings {
    return StringsFromInst(SaveDialog_GetFiles(s.instance))
}

func (s *TSaveDialog) DefaultExt() string {
    return SaveDialog_GetDefaultExt(s.instance)
}

func (s *TSaveDialog) SetDefaultExt(value string) {
    SaveDialog_SetDefaultExt(s.instance, value)
}

func (s *TSaveDialog) FileName() string {
    return SaveDialog_GetFileName(s.instance)
}

func (s *TSaveDialog) SetFileName(value string) {
    SaveDialog_SetFileName(s.instance, value)
}

func (s *TSaveDialog) Filter() string {
    return SaveDialog_GetFilter(s.instance)
}

func (s *TSaveDialog) SetFilter(value string) {
    SaveDialog_SetFilter(s.instance, value)
}

func (s *TSaveDialog) FilterIndex() int32 {
    return SaveDialog_GetFilterIndex(s.instance)
}

func (s *TSaveDialog) SetFilterIndex(value int32) {
    SaveDialog_SetFilterIndex(s.instance, value)
}

func (s *TSaveDialog) InitialDir() string {
    return SaveDialog_GetInitialDir(s.instance)
}

func (s *TSaveDialog) SetInitialDir(value string) {
    SaveDialog_SetInitialDir(s.instance, value)
}

func (s *TSaveDialog) Options() TOpenOptions {
    return SaveDialog_GetOptions(s.instance)
}

func (s *TSaveDialog) SetOptions(value TOpenOptions) {
    SaveDialog_SetOptions(s.instance, value)
}

func (s *TSaveDialog) OptionsEx() TOpenOptionsEx {
    return SaveDialog_GetOptionsEx(s.instance)
}

func (s *TSaveDialog) SetOptionsEx(value TOpenOptionsEx) {
    SaveDialog_SetOptionsEx(s.instance, value)
}

func (s *TSaveDialog) Title() string {
    return SaveDialog_GetTitle(s.instance)
}

func (s *TSaveDialog) SetTitle(value string) {
    SaveDialog_SetTitle(s.instance, value)
}

func (s *TSaveDialog) Handle() HWND {
    return SaveDialog_GetHandle(s.instance)
}

func (s *TSaveDialog) SetOnClose(fn TNotifyEvent) {
    SaveDialog_SetOnClose(s.instance, fn)
}

func (s *TSaveDialog) SetOnShow(fn TNotifyEvent) {
    SaveDialog_SetOnShow(s.instance, fn)
}

func (s *TSaveDialog) ComponentCount() int32 {
    return SaveDialog_GetComponentCount(s.instance)
}

func (s *TSaveDialog) ComponentIndex() int32 {
    return SaveDialog_GetComponentIndex(s.instance)
}

func (s *TSaveDialog) SetComponentIndex(value int32) {
    SaveDialog_SetComponentIndex(s.instance, value)
}

func (s *TSaveDialog) Owner() *TComponent {
    return ComponentFromInst(SaveDialog_GetOwner(s.instance))
}

func (s *TSaveDialog) Name() string {
    return SaveDialog_GetName(s.instance)
}

func (s *TSaveDialog) SetName(value string) {
    SaveDialog_SetName(s.instance, value)
}

func (s *TSaveDialog) Tag() int {
    return SaveDialog_GetTag(s.instance)
}

func (s *TSaveDialog) SetTag(value int) {
    SaveDialog_SetTag(s.instance, value)
}

func (s *TSaveDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(SaveDialog_GetComponents(s.instance, AIndex))
}

