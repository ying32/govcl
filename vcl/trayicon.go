
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

type TTrayIcon struct {
    IComponent
    instance uintptr
}

func NewTrayIcon(owner IComponent) *TTrayIcon {
    t := new(TTrayIcon)
    t.instance = TrayIcon_Create(CheckPtr(owner))
    return t
}

func TrayIconFromInst(inst uintptr) *TTrayIcon {
    t := new(TTrayIcon)
    t.instance = inst
    return t
}

func TrayIconFromObj(obj IObject) *TTrayIcon {
    t := new(TTrayIcon)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTrayIcon) Free() {
    if t.instance != 0 {
        TrayIcon_Free(t.instance)
        t.instance = 0
    }
}

func (t *TTrayIcon) Instance() uintptr {
    return t.instance
}

func (t *TTrayIcon) IsValid() bool {
    return t.instance != 0
}

func (t *TTrayIcon) Refresh() {
    TrayIcon_Refresh(t.instance)
}

func (t *TTrayIcon) SetDefaultIcon() {
    TrayIcon_SetDefaultIcon(t.instance)
}

func (t *TTrayIcon) ShowBalloonHint() {
    TrayIcon_ShowBalloonHint(t.instance)
}

func (t *TTrayIcon) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TrayIcon_FindComponent(t.instance, AName))
}

func (t *TTrayIcon) GetNamePath() string {
    return TrayIcon_GetNamePath(t.instance)
}

func (t *TTrayIcon) HasParent() bool {
    return TrayIcon_HasParent(t.instance)
}

func (t *TTrayIcon) Assign(Source IObject) {
    TrayIcon_Assign(t.instance, CheckPtr(Source))
}

func (t *TTrayIcon) ClassName() string {
    return TrayIcon_ClassName(t.instance)
}

func (t *TTrayIcon) Equals(Obj IObject) bool {
    return TrayIcon_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTrayIcon) GetHashCode() int32 {
    return TrayIcon_GetHashCode(t.instance)
}

func (t *TTrayIcon) ToString() string {
    return TrayIcon_ToString(t.instance)
}

func (t *TTrayIcon) Animate() bool {
    return TrayIcon_GetAnimate(t.instance)
}

func (t *TTrayIcon) SetAnimate(value bool) {
    TrayIcon_SetAnimate(t.instance, value)
}

func (t *TTrayIcon) AnimateInterval() uint32 {
    return TrayIcon_GetAnimateInterval(t.instance)
}

func (t *TTrayIcon) SetAnimateInterval(value uint32) {
    TrayIcon_SetAnimateInterval(t.instance, value)
}

func (t *TTrayIcon) Hint() string {
    return TrayIcon_GetHint(t.instance)
}

func (t *TTrayIcon) SetHint(value string) {
    TrayIcon_SetHint(t.instance, value)
}

func (t *TTrayIcon) BalloonHint() string {
    return TrayIcon_GetBalloonHint(t.instance)
}

func (t *TTrayIcon) SetBalloonHint(value string) {
    TrayIcon_SetBalloonHint(t.instance, value)
}

func (t *TTrayIcon) BalloonTitle() string {
    return TrayIcon_GetBalloonTitle(t.instance)
}

func (t *TTrayIcon) SetBalloonTitle(value string) {
    TrayIcon_SetBalloonTitle(t.instance, value)
}

func (t *TTrayIcon) BalloonTimeout() int32 {
    return TrayIcon_GetBalloonTimeout(t.instance)
}

func (t *TTrayIcon) SetBalloonTimeout(value int32) {
    TrayIcon_SetBalloonTimeout(t.instance, value)
}

func (t *TTrayIcon) BalloonFlags() TBalloonFlags {
    return TrayIcon_GetBalloonFlags(t.instance)
}

func (t *TTrayIcon) SetBalloonFlags(value TBalloonFlags) {
    TrayIcon_SetBalloonFlags(t.instance, value)
}

func (t *TTrayIcon) Icon() *TIcon {
    return IconFromInst(TrayIcon_GetIcon(t.instance))
}

func (t *TTrayIcon) SetIcon(value *TIcon) {
    TrayIcon_SetIcon(t.instance, CheckPtr(value))
}

func (t *TTrayIcon) IconIndex() int32 {
    return TrayIcon_GetIconIndex(t.instance)
}

func (t *TTrayIcon) SetIconIndex(value int32) {
    TrayIcon_SetIconIndex(t.instance, value)
}

func (t *TTrayIcon) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TrayIcon_GetPopupMenu(t.instance))
}

func (t *TTrayIcon) SetPopupMenu(value IComponent) {
    TrayIcon_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TTrayIcon) Visible() bool {
    return TrayIcon_GetVisible(t.instance)
}

func (t *TTrayIcon) SetVisible(value bool) {
    TrayIcon_SetVisible(t.instance, value)
}

func (t *TTrayIcon) SetOnBalloonClick(fn TNotifyEvent) {
    TrayIcon_SetOnBalloonClick(t.instance, fn)
}

func (t *TTrayIcon) SetOnClick(fn TNotifyEvent) {
    TrayIcon_SetOnClick(t.instance, fn)
}

func (t *TTrayIcon) SetOnDblClick(fn TNotifyEvent) {
    TrayIcon_SetOnDblClick(t.instance, fn)
}

func (t *TTrayIcon) SetOnMouseMove(fn TMouseMoveEvent) {
    TrayIcon_SetOnMouseMove(t.instance, fn)
}

func (t *TTrayIcon) SetOnMouseUp(fn TMouseEvent) {
    TrayIcon_SetOnMouseUp(t.instance, fn)
}

func (t *TTrayIcon) SetOnMouseDown(fn TMouseEvent) {
    TrayIcon_SetOnMouseDown(t.instance, fn)
}

func (t *TTrayIcon) ComponentCount() int32 {
    return TrayIcon_GetComponentCount(t.instance)
}

func (t *TTrayIcon) ComponentIndex() int32 {
    return TrayIcon_GetComponentIndex(t.instance)
}

func (t *TTrayIcon) SetComponentIndex(value int32) {
    TrayIcon_SetComponentIndex(t.instance, value)
}

func (t *TTrayIcon) Owner() *TComponent {
    return ComponentFromInst(TrayIcon_GetOwner(t.instance))
}

func (t *TTrayIcon) Name() string {
    return TrayIcon_GetName(t.instance)
}

func (t *TTrayIcon) SetName(value string) {
    TrayIcon_SetName(t.instance, value)
}

func (t *TTrayIcon) Tag() int {
    return TrayIcon_GetTag(t.instance)
}

func (t *TTrayIcon) SetTag(value int) {
    TrayIcon_SetTag(t.instance, value)
}

func (t *TTrayIcon) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TrayIcon_GetComponents(t.instance, AIndex))
}

