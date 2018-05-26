
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

type TCustomHint struct {
    IComponent
    instance uintptr
}

func NewCustomHint(owner IComponent) *TCustomHint {
    c := new(TCustomHint)
    c.instance = CustomHint_Create(CheckPtr(owner))
    return c
}

func CustomHintFromInst(inst uintptr) *TCustomHint {
    c := new(TCustomHint)
    c.instance = inst
    return c
}

func CustomHintFromObj(obj IObject) *TCustomHint {
    c := new(TCustomHint)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCustomHint) Free() {
    if c.instance != 0 {
        CustomHint_Free(c.instance)
        c.instance = 0
    }
}

func (c *TCustomHint) Instance() uintptr {
    return c.instance
}

func (c *TCustomHint) IsValid() bool {
    return c.instance != 0
}

func (c *TCustomHint) ShowHint() {
    CustomHint_ShowHint(c.instance)
}

func (c *TCustomHint) HideHint() {
    CustomHint_HideHint(c.instance)
}

func (c *TCustomHint) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CustomHint_FindComponent(c.instance, AName))
}

func (c *TCustomHint) GetNamePath() string {
    return CustomHint_GetNamePath(c.instance)
}

func (c *TCustomHint) HasParent() bool {
    return CustomHint_HasParent(c.instance)
}

func (c *TCustomHint) Assign(Source IObject) {
    CustomHint_Assign(c.instance, CheckPtr(Source))
}

func (c *TCustomHint) ClassName() string {
    return CustomHint_ClassName(c.instance)
}

func (c *TCustomHint) Equals(Obj IObject) bool {
    return CustomHint_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCustomHint) GetHashCode() int32 {
    return CustomHint_GetHashCode(c.instance)
}

func (c *TCustomHint) ToString() string {
    return CustomHint_ToString(c.instance)
}

func (c *TCustomHint) ShowingHint() bool {
    return CustomHint_GetShowingHint(c.instance)
}

func (c *TCustomHint) Title() string {
    return CustomHint_GetTitle(c.instance)
}

func (c *TCustomHint) SetTitle(value string) {
    CustomHint_SetTitle(c.instance, value)
}

func (c *TCustomHint) Description() string {
    return CustomHint_GetDescription(c.instance)
}

func (c *TCustomHint) SetDescription(value string) {
    CustomHint_SetDescription(c.instance, value)
}

func (c *TCustomHint) ImageIndex() int32 {
    return CustomHint_GetImageIndex(c.instance)
}

func (c *TCustomHint) SetImageIndex(value int32) {
    CustomHint_SetImageIndex(c.instance, value)
}

func (c *TCustomHint) Images() *TImageList {
    return ImageListFromInst(CustomHint_GetImages(c.instance))
}

func (c *TCustomHint) SetImages(value IComponent) {
    CustomHint_SetImages(c.instance, CheckPtr(value))
}

func (c *TCustomHint) Style() TBalloonHintStyle {
    return CustomHint_GetStyle(c.instance)
}

func (c *TCustomHint) SetStyle(value TBalloonHintStyle) {
    CustomHint_SetStyle(c.instance, value)
}

func (c *TCustomHint) Delay() uint32 {
    return CustomHint_GetDelay(c.instance)
}

func (c *TCustomHint) SetDelay(value uint32) {
    CustomHint_SetDelay(c.instance, value)
}

func (c *TCustomHint) HideAfter() int32 {
    return CustomHint_GetHideAfter(c.instance)
}

func (c *TCustomHint) SetHideAfter(value int32) {
    CustomHint_SetHideAfter(c.instance, value)
}

func (c *TCustomHint) ComponentCount() int32 {
    return CustomHint_GetComponentCount(c.instance)
}

func (c *TCustomHint) ComponentIndex() int32 {
    return CustomHint_GetComponentIndex(c.instance)
}

func (c *TCustomHint) SetComponentIndex(value int32) {
    CustomHint_SetComponentIndex(c.instance, value)
}

func (c *TCustomHint) Owner() *TComponent {
    return ComponentFromInst(CustomHint_GetOwner(c.instance))
}

func (c *TCustomHint) Name() string {
    return CustomHint_GetName(c.instance)
}

func (c *TCustomHint) SetName(value string) {
    CustomHint_SetName(c.instance, value)
}

func (c *TCustomHint) Tag() int {
    return CustomHint_GetTag(c.instance)
}

func (c *TCustomHint) SetTag(value int) {
    CustomHint_SetTag(c.instance, value)
}

func (c *TCustomHint) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CustomHint_GetComponents(c.instance, AIndex))
}

