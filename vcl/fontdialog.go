
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

type TFontDialog struct {
    IComponent
    instance uintptr
}

func NewFontDialog(owner IComponent) *TFontDialog {
    f := new(TFontDialog)
    f.instance = FontDialog_Create(CheckPtr(owner))
    return f
}

func FontDialogFromInst(inst uintptr) *TFontDialog {
    f := new(TFontDialog)
    f.instance = inst
    return f
}

func FontDialogFromObj(obj IObject) *TFontDialog {
    f := new(TFontDialog)
    f.instance = CheckPtr(obj)
    return f
}

func (f *TFontDialog) Free() {
    if f.instance != 0 {
        FontDialog_Free(f.instance)
        f.instance = 0
    }
}

func (f *TFontDialog) Instance() uintptr {
    return f.instance
}

func (f *TFontDialog) IsValid() bool {
    return f.instance != 0
}

func (f *TFontDialog) Execute() bool {
    return FontDialog_Execute(f.instance)
}

func (f *TFontDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(FontDialog_FindComponent(f.instance, AName))
}

func (f *TFontDialog) GetNamePath() string {
    return FontDialog_GetNamePath(f.instance)
}

func (f *TFontDialog) HasParent() bool {
    return FontDialog_HasParent(f.instance)
}

func (f *TFontDialog) Assign(Source IObject) {
    FontDialog_Assign(f.instance, CheckPtr(Source))
}

func (f *TFontDialog) ClassName() string {
    return FontDialog_ClassName(f.instance)
}

func (f *TFontDialog) Equals(Obj IObject) bool {
    return FontDialog_Equals(f.instance, CheckPtr(Obj))
}

func (f *TFontDialog) GetHashCode() int32 {
    return FontDialog_GetHashCode(f.instance)
}

func (f *TFontDialog) ToString() string {
    return FontDialog_ToString(f.instance)
}

func (f *TFontDialog) Font() *TFont {
    return FontFromInst(FontDialog_GetFont(f.instance))
}

func (f *TFontDialog) SetFont(value *TFont) {
    FontDialog_SetFont(f.instance, CheckPtr(value))
}

func (f *TFontDialog) Options() TFontDialogOptions {
    return FontDialog_GetOptions(f.instance)
}

func (f *TFontDialog) SetOptions(value TFontDialogOptions) {
    FontDialog_SetOptions(f.instance, value)
}

func (f *TFontDialog) Handle() HWND {
    return FontDialog_GetHandle(f.instance)
}

func (f *TFontDialog) SetOnClose(fn TNotifyEvent) {
    FontDialog_SetOnClose(f.instance, fn)
}

func (f *TFontDialog) SetOnShow(fn TNotifyEvent) {
    FontDialog_SetOnShow(f.instance, fn)
}

func (f *TFontDialog) ComponentCount() int32 {
    return FontDialog_GetComponentCount(f.instance)
}

func (f *TFontDialog) ComponentIndex() int32 {
    return FontDialog_GetComponentIndex(f.instance)
}

func (f *TFontDialog) SetComponentIndex(value int32) {
    FontDialog_SetComponentIndex(f.instance, value)
}

func (f *TFontDialog) Owner() *TComponent {
    return ComponentFromInst(FontDialog_GetOwner(f.instance))
}

func (f *TFontDialog) Name() string {
    return FontDialog_GetName(f.instance)
}

func (f *TFontDialog) SetName(value string) {
    FontDialog_SetName(f.instance, value)
}

func (f *TFontDialog) Tag() int {
    return FontDialog_GetTag(f.instance)
}

func (f *TFontDialog) SetTag(value int) {
    FontDialog_SetTag(f.instance, value)
}

func (f *TFontDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(FontDialog_GetComponents(f.instance, AIndex))
}

