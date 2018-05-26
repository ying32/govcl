
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

type TSaveTextFileDialog struct {
    IComponent
    instance uintptr
}

func NewSaveTextFileDialog(owner IComponent) *TSaveTextFileDialog {
    s := new(TSaveTextFileDialog)
    s.instance = SaveTextFileDialog_Create(CheckPtr(owner))
    return s
}

func SaveTextFileDialogFromInst(inst uintptr) *TSaveTextFileDialog {
    s := new(TSaveTextFileDialog)
    s.instance = inst
    return s
}

func SaveTextFileDialogFromObj(obj IObject) *TSaveTextFileDialog {
    s := new(TSaveTextFileDialog)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TSaveTextFileDialog) Free() {
    if s.instance != 0 {
        SaveTextFileDialog_Free(s.instance)
        s.instance = 0
    }
}

func (s *TSaveTextFileDialog) Instance() uintptr {
    return s.instance
}

func (s *TSaveTextFileDialog) IsValid() bool {
    return s.instance != 0
}

func (s *TSaveTextFileDialog) Execute() bool {
    return SaveTextFileDialog_Execute(s.instance)
}

func (s *TSaveTextFileDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(SaveTextFileDialog_FindComponent(s.instance, AName))
}

func (s *TSaveTextFileDialog) GetNamePath() string {
    return SaveTextFileDialog_GetNamePath(s.instance)
}

func (s *TSaveTextFileDialog) HasParent() bool {
    return SaveTextFileDialog_HasParent(s.instance)
}

func (s *TSaveTextFileDialog) Assign(Source IObject) {
    SaveTextFileDialog_Assign(s.instance, CheckPtr(Source))
}

func (s *TSaveTextFileDialog) ClassName() string {
    return SaveTextFileDialog_ClassName(s.instance)
}

func (s *TSaveTextFileDialog) Equals(Obj IObject) bool {
    return SaveTextFileDialog_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSaveTextFileDialog) GetHashCode() int32 {
    return SaveTextFileDialog_GetHashCode(s.instance)
}

func (s *TSaveTextFileDialog) ToString() string {
    return SaveTextFileDialog_ToString(s.instance)
}

func (s *TSaveTextFileDialog) Files() *TStrings {
    return StringsFromInst(SaveTextFileDialog_GetFiles(s.instance))
}

func (s *TSaveTextFileDialog) DefaultExt() string {
    return SaveTextFileDialog_GetDefaultExt(s.instance)
}

func (s *TSaveTextFileDialog) SetDefaultExt(value string) {
    SaveTextFileDialog_SetDefaultExt(s.instance, value)
}

func (s *TSaveTextFileDialog) FileName() string {
    return SaveTextFileDialog_GetFileName(s.instance)
}

func (s *TSaveTextFileDialog) SetFileName(value string) {
    SaveTextFileDialog_SetFileName(s.instance, value)
}

func (s *TSaveTextFileDialog) Filter() string {
    return SaveTextFileDialog_GetFilter(s.instance)
}

func (s *TSaveTextFileDialog) SetFilter(value string) {
    SaveTextFileDialog_SetFilter(s.instance, value)
}

func (s *TSaveTextFileDialog) FilterIndex() int32 {
    return SaveTextFileDialog_GetFilterIndex(s.instance)
}

func (s *TSaveTextFileDialog) SetFilterIndex(value int32) {
    SaveTextFileDialog_SetFilterIndex(s.instance, value)
}

func (s *TSaveTextFileDialog) InitialDir() string {
    return SaveTextFileDialog_GetInitialDir(s.instance)
}

func (s *TSaveTextFileDialog) SetInitialDir(value string) {
    SaveTextFileDialog_SetInitialDir(s.instance, value)
}

func (s *TSaveTextFileDialog) Options() TOpenOptions {
    return SaveTextFileDialog_GetOptions(s.instance)
}

func (s *TSaveTextFileDialog) SetOptions(value TOpenOptions) {
    SaveTextFileDialog_SetOptions(s.instance, value)
}

func (s *TSaveTextFileDialog) OptionsEx() TOpenOptionsEx {
    return SaveTextFileDialog_GetOptionsEx(s.instance)
}

func (s *TSaveTextFileDialog) SetOptionsEx(value TOpenOptionsEx) {
    SaveTextFileDialog_SetOptionsEx(s.instance, value)
}

func (s *TSaveTextFileDialog) Title() string {
    return SaveTextFileDialog_GetTitle(s.instance)
}

func (s *TSaveTextFileDialog) SetTitle(value string) {
    SaveTextFileDialog_SetTitle(s.instance, value)
}

func (s *TSaveTextFileDialog) Handle() HWND {
    return SaveTextFileDialog_GetHandle(s.instance)
}

func (s *TSaveTextFileDialog) SetOnClose(fn TNotifyEvent) {
    SaveTextFileDialog_SetOnClose(s.instance, fn)
}

func (s *TSaveTextFileDialog) SetOnShow(fn TNotifyEvent) {
    SaveTextFileDialog_SetOnShow(s.instance, fn)
}

func (s *TSaveTextFileDialog) ComponentCount() int32 {
    return SaveTextFileDialog_GetComponentCount(s.instance)
}

func (s *TSaveTextFileDialog) ComponentIndex() int32 {
    return SaveTextFileDialog_GetComponentIndex(s.instance)
}

func (s *TSaveTextFileDialog) SetComponentIndex(value int32) {
    SaveTextFileDialog_SetComponentIndex(s.instance, value)
}

func (s *TSaveTextFileDialog) Owner() *TComponent {
    return ComponentFromInst(SaveTextFileDialog_GetOwner(s.instance))
}

func (s *TSaveTextFileDialog) Name() string {
    return SaveTextFileDialog_GetName(s.instance)
}

func (s *TSaveTextFileDialog) SetName(value string) {
    SaveTextFileDialog_SetName(s.instance, value)
}

func (s *TSaveTextFileDialog) Tag() int {
    return SaveTextFileDialog_GetTag(s.instance)
}

func (s *TSaveTextFileDialog) SetTag(value int) {
    SaveTextFileDialog_SetTag(s.instance, value)
}

func (s *TSaveTextFileDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(SaveTextFileDialog_GetComponents(s.instance, AIndex))
}

