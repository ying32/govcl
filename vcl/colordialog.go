
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

type TColorDialog struct {
    IComponent
    instance uintptr
}

func NewColorDialog(owner IComponent) *TColorDialog {
    c := new(TColorDialog)
    c.instance = ColorDialog_Create(CheckPtr(owner))
    return c
}

func ColorDialogFromInst(inst uintptr) *TColorDialog {
    c := new(TColorDialog)
    c.instance = inst
    return c
}

func ColorDialogFromObj(obj IObject) *TColorDialog {
    c := new(TColorDialog)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TColorDialog) Free() {
    if c.instance != 0 {
        ColorDialog_Free(c.instance)
        c.instance = 0
    }
}

func (c *TColorDialog) Instance() uintptr {
    return c.instance
}

func (c *TColorDialog) IsValid() bool {
    return c.instance != 0
}

func (c *TColorDialog) Execute() bool {
    return ColorDialog_Execute(c.instance)
}

func (c *TColorDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ColorDialog_FindComponent(c.instance, AName))
}

func (c *TColorDialog) GetNamePath() string {
    return ColorDialog_GetNamePath(c.instance)
}

func (c *TColorDialog) HasParent() bool {
    return ColorDialog_HasParent(c.instance)
}

func (c *TColorDialog) Assign(Source IObject) {
    ColorDialog_Assign(c.instance, CheckPtr(Source))
}

func (c *TColorDialog) ClassName() string {
    return ColorDialog_ClassName(c.instance)
}

func (c *TColorDialog) Equals(Obj IObject) bool {
    return ColorDialog_Equals(c.instance, CheckPtr(Obj))
}

func (c *TColorDialog) GetHashCode() int32 {
    return ColorDialog_GetHashCode(c.instance)
}

func (c *TColorDialog) ToString() string {
    return ColorDialog_ToString(c.instance)
}

func (c *TColorDialog) Color() TColor {
    return ColorDialog_GetColor(c.instance)
}

func (c *TColorDialog) SetColor(value TColor) {
    ColorDialog_SetColor(c.instance, value)
}

func (c *TColorDialog) Options() TColorDialogOptions {
    return ColorDialog_GetOptions(c.instance)
}

func (c *TColorDialog) SetOptions(value TColorDialogOptions) {
    ColorDialog_SetOptions(c.instance, value)
}

func (c *TColorDialog) Handle() HWND {
    return ColorDialog_GetHandle(c.instance)
}

func (c *TColorDialog) SetOnClose(fn TNotifyEvent) {
    ColorDialog_SetOnClose(c.instance, fn)
}

func (c *TColorDialog) SetOnShow(fn TNotifyEvent) {
    ColorDialog_SetOnShow(c.instance, fn)
}

func (c *TColorDialog) ComponentCount() int32 {
    return ColorDialog_GetComponentCount(c.instance)
}

func (c *TColorDialog) ComponentIndex() int32 {
    return ColorDialog_GetComponentIndex(c.instance)
}

func (c *TColorDialog) SetComponentIndex(value int32) {
    ColorDialog_SetComponentIndex(c.instance, value)
}

func (c *TColorDialog) Owner() *TComponent {
    return ComponentFromInst(ColorDialog_GetOwner(c.instance))
}

func (c *TColorDialog) Name() string {
    return ColorDialog_GetName(c.instance)
}

func (c *TColorDialog) SetName(value string) {
    ColorDialog_SetName(c.instance, value)
}

func (c *TColorDialog) Tag() int {
    return ColorDialog_GetTag(c.instance)
}

func (c *TColorDialog) SetTag(value int) {
    ColorDialog_SetTag(c.instance, value)
}

func (c *TColorDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ColorDialog_GetComponents(c.instance, AIndex))
}

