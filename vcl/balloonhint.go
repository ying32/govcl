
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

type TBalloonHint struct {
    IComponent
    instance uintptr
}

func NewBalloonHint(owner IComponent) *TBalloonHint {
    b := new(TBalloonHint)
    b.instance = BalloonHint_Create(CheckPtr(owner))
    return b
}

func BalloonHintFromInst(inst uintptr) *TBalloonHint {
    b := new(TBalloonHint)
    b.instance = inst
    return b
}

func BalloonHintFromObj(obj IObject) *TBalloonHint {
    b := new(TBalloonHint)
    b.instance = CheckPtr(obj)
    return b
}

func (b *TBalloonHint) Free() {
    if b.instance != 0 {
        BalloonHint_Free(b.instance)
        b.instance = 0
    }
}

func (b *TBalloonHint) Instance() uintptr {
    return b.instance
}

func (b *TBalloonHint) IsValid() bool {
    return b.instance != 0
}

func (b *TBalloonHint) ShowHint() {
    BalloonHint_ShowHint(b.instance)
}

func (b *TBalloonHint) HideHint() {
    BalloonHint_HideHint(b.instance)
}

func (b *TBalloonHint) FindComponent(AName string) *TComponent {
    return ComponentFromInst(BalloonHint_FindComponent(b.instance, AName))
}

func (b *TBalloonHint) GetNamePath() string {
    return BalloonHint_GetNamePath(b.instance)
}

func (b *TBalloonHint) HasParent() bool {
    return BalloonHint_HasParent(b.instance)
}

func (b *TBalloonHint) Assign(Source IObject) {
    BalloonHint_Assign(b.instance, CheckPtr(Source))
}

func (b *TBalloonHint) ClassName() string {
    return BalloonHint_ClassName(b.instance)
}

func (b *TBalloonHint) Equals(Obj IObject) bool {
    return BalloonHint_Equals(b.instance, CheckPtr(Obj))
}

func (b *TBalloonHint) GetHashCode() int32 {
    return BalloonHint_GetHashCode(b.instance)
}

func (b *TBalloonHint) ToString() string {
    return BalloonHint_ToString(b.instance)
}

func (b *TBalloonHint) ShowingHint() bool {
    return BalloonHint_GetShowingHint(b.instance)
}

func (b *TBalloonHint) Title() string {
    return BalloonHint_GetTitle(b.instance)
}

func (b *TBalloonHint) SetTitle(value string) {
    BalloonHint_SetTitle(b.instance, value)
}

func (b *TBalloonHint) Description() string {
    return BalloonHint_GetDescription(b.instance)
}

func (b *TBalloonHint) SetDescription(value string) {
    BalloonHint_SetDescription(b.instance, value)
}

func (b *TBalloonHint) ImageIndex() int32 {
    return BalloonHint_GetImageIndex(b.instance)
}

func (b *TBalloonHint) SetImageIndex(value int32) {
    BalloonHint_SetImageIndex(b.instance, value)
}

func (b *TBalloonHint) Images() *TImageList {
    return ImageListFromInst(BalloonHint_GetImages(b.instance))
}

func (b *TBalloonHint) SetImages(value IComponent) {
    BalloonHint_SetImages(b.instance, CheckPtr(value))
}

func (b *TBalloonHint) Style() TBalloonHintStyle {
    return BalloonHint_GetStyle(b.instance)
}

func (b *TBalloonHint) SetStyle(value TBalloonHintStyle) {
    BalloonHint_SetStyle(b.instance, value)
}

func (b *TBalloonHint) Delay() uint32 {
    return BalloonHint_GetDelay(b.instance)
}

func (b *TBalloonHint) SetDelay(value uint32) {
    BalloonHint_SetDelay(b.instance, value)
}

func (b *TBalloonHint) HideAfter() int32 {
    return BalloonHint_GetHideAfter(b.instance)
}

func (b *TBalloonHint) SetHideAfter(value int32) {
    BalloonHint_SetHideAfter(b.instance, value)
}

func (b *TBalloonHint) ComponentCount() int32 {
    return BalloonHint_GetComponentCount(b.instance)
}

func (b *TBalloonHint) ComponentIndex() int32 {
    return BalloonHint_GetComponentIndex(b.instance)
}

func (b *TBalloonHint) SetComponentIndex(value int32) {
    BalloonHint_SetComponentIndex(b.instance, value)
}

func (b *TBalloonHint) Owner() *TComponent {
    return ComponentFromInst(BalloonHint_GetOwner(b.instance))
}

func (b *TBalloonHint) Name() string {
    return BalloonHint_GetName(b.instance)
}

func (b *TBalloonHint) SetName(value string) {
    BalloonHint_SetName(b.instance, value)
}

func (b *TBalloonHint) Tag() int {
    return BalloonHint_GetTag(b.instance)
}

func (b *TBalloonHint) SetTag(value int) {
    BalloonHint_SetTag(b.instance, value)
}

func (b *TBalloonHint) Components(AIndex int32) *TComponent {
    return ComponentFromInst(BalloonHint_GetComponents(b.instance, AIndex))
}

