
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

type TSavePictureDialog struct {
    IComponent
    instance uintptr
}

func NewSavePictureDialog(owner IComponent) *TSavePictureDialog {
    s := new(TSavePictureDialog)
    s.instance = SavePictureDialog_Create(CheckPtr(owner))
    return s
}

func SavePictureDialogFromInst(inst uintptr) *TSavePictureDialog {
    s := new(TSavePictureDialog)
    s.instance = inst
    return s
}

func SavePictureDialogFromObj(obj IObject) *TSavePictureDialog {
    s := new(TSavePictureDialog)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TSavePictureDialog) Free() {
    if s.instance != 0 {
        SavePictureDialog_Free(s.instance)
        s.instance = 0
    }
}

func (s *TSavePictureDialog) Instance() uintptr {
    return s.instance
}

func (s *TSavePictureDialog) IsValid() bool {
    return s.instance != 0
}

func (s *TSavePictureDialog) Execute() bool {
    return SavePictureDialog_Execute(s.instance)
}

func (s *TSavePictureDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(SavePictureDialog_FindComponent(s.instance, AName))
}

func (s *TSavePictureDialog) GetNamePath() string {
    return SavePictureDialog_GetNamePath(s.instance)
}

func (s *TSavePictureDialog) HasParent() bool {
    return SavePictureDialog_HasParent(s.instance)
}

func (s *TSavePictureDialog) Assign(Source IObject) {
    SavePictureDialog_Assign(s.instance, CheckPtr(Source))
}

func (s *TSavePictureDialog) ClassName() string {
    return SavePictureDialog_ClassName(s.instance)
}

func (s *TSavePictureDialog) Equals(Obj IObject) bool {
    return SavePictureDialog_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSavePictureDialog) GetHashCode() int32 {
    return SavePictureDialog_GetHashCode(s.instance)
}

func (s *TSavePictureDialog) ToString() string {
    return SavePictureDialog_ToString(s.instance)
}

func (s *TSavePictureDialog) Filter() string {
    return SavePictureDialog_GetFilter(s.instance)
}

func (s *TSavePictureDialog) SetFilter(value string) {
    SavePictureDialog_SetFilter(s.instance, value)
}

func (s *TSavePictureDialog) Files() *TStrings {
    return StringsFromInst(SavePictureDialog_GetFiles(s.instance))
}

func (s *TSavePictureDialog) DefaultExt() string {
    return SavePictureDialog_GetDefaultExt(s.instance)
}

func (s *TSavePictureDialog) SetDefaultExt(value string) {
    SavePictureDialog_SetDefaultExt(s.instance, value)
}

func (s *TSavePictureDialog) FileName() string {
    return SavePictureDialog_GetFileName(s.instance)
}

func (s *TSavePictureDialog) SetFileName(value string) {
    SavePictureDialog_SetFileName(s.instance, value)
}

func (s *TSavePictureDialog) FilterIndex() int32 {
    return SavePictureDialog_GetFilterIndex(s.instance)
}

func (s *TSavePictureDialog) SetFilterIndex(value int32) {
    SavePictureDialog_SetFilterIndex(s.instance, value)
}

func (s *TSavePictureDialog) InitialDir() string {
    return SavePictureDialog_GetInitialDir(s.instance)
}

func (s *TSavePictureDialog) SetInitialDir(value string) {
    SavePictureDialog_SetInitialDir(s.instance, value)
}

func (s *TSavePictureDialog) Options() TOpenOptions {
    return SavePictureDialog_GetOptions(s.instance)
}

func (s *TSavePictureDialog) SetOptions(value TOpenOptions) {
    SavePictureDialog_SetOptions(s.instance, value)
}

func (s *TSavePictureDialog) OptionsEx() TOpenOptionsEx {
    return SavePictureDialog_GetOptionsEx(s.instance)
}

func (s *TSavePictureDialog) SetOptionsEx(value TOpenOptionsEx) {
    SavePictureDialog_SetOptionsEx(s.instance, value)
}

func (s *TSavePictureDialog) Title() string {
    return SavePictureDialog_GetTitle(s.instance)
}

func (s *TSavePictureDialog) SetTitle(value string) {
    SavePictureDialog_SetTitle(s.instance, value)
}

func (s *TSavePictureDialog) Handle() HWND {
    return SavePictureDialog_GetHandle(s.instance)
}

func (s *TSavePictureDialog) SetOnClose(fn TNotifyEvent) {
    SavePictureDialog_SetOnClose(s.instance, fn)
}

func (s *TSavePictureDialog) SetOnShow(fn TNotifyEvent) {
    SavePictureDialog_SetOnShow(s.instance, fn)
}

func (s *TSavePictureDialog) ComponentCount() int32 {
    return SavePictureDialog_GetComponentCount(s.instance)
}

func (s *TSavePictureDialog) ComponentIndex() int32 {
    return SavePictureDialog_GetComponentIndex(s.instance)
}

func (s *TSavePictureDialog) SetComponentIndex(value int32) {
    SavePictureDialog_SetComponentIndex(s.instance, value)
}

func (s *TSavePictureDialog) Owner() *TComponent {
    return ComponentFromInst(SavePictureDialog_GetOwner(s.instance))
}

func (s *TSavePictureDialog) Name() string {
    return SavePictureDialog_GetName(s.instance)
}

func (s *TSavePictureDialog) SetName(value string) {
    SavePictureDialog_SetName(s.instance, value)
}

func (s *TSavePictureDialog) Tag() int {
    return SavePictureDialog_GetTag(s.instance)
}

func (s *TSavePictureDialog) SetTag(value int) {
    SavePictureDialog_SetTag(s.instance, value)
}

func (s *TSavePictureDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(SavePictureDialog_GetComponents(s.instance, AIndex))
}

