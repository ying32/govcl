
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

type TPopupMenu struct {
    IComponent
    instance uintptr
}

func NewPopupMenu(owner IComponent) *TPopupMenu {
    p := new(TPopupMenu)
    p.instance = PopupMenu_Create(CheckPtr(owner))
    return p
}

func PopupMenuFromInst(inst uintptr) *TPopupMenu {
    p := new(TPopupMenu)
    p.instance = inst
    return p
}

func PopupMenuFromObj(obj IObject) *TPopupMenu {
    p := new(TPopupMenu)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPopupMenu) Free() {
    if p.instance != 0 {
        PopupMenu_Free(p.instance)
        p.instance = 0
    }
}

func (p *TPopupMenu) Instance() uintptr {
    return p.instance
}

func (p *TPopupMenu) IsValid() bool {
    return p.instance != 0
}

func (p *TPopupMenu) CloseMenu() {
    PopupMenu_CloseMenu(p.instance)
}

func (p *TPopupMenu) Popup(X int32, Y int32) {
    PopupMenu_Popup(p.instance, X , Y)
}

func (p *TPopupMenu) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PopupMenu_FindComponent(p.instance, AName))
}

func (p *TPopupMenu) GetNamePath() string {
    return PopupMenu_GetNamePath(p.instance)
}

func (p *TPopupMenu) HasParent() bool {
    return PopupMenu_HasParent(p.instance)
}

func (p *TPopupMenu) Assign(Source IObject) {
    PopupMenu_Assign(p.instance, CheckPtr(Source))
}

func (p *TPopupMenu) ClassName() string {
    return PopupMenu_ClassName(p.instance)
}

func (p *TPopupMenu) Equals(Obj IObject) bool {
    return PopupMenu_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPopupMenu) GetHashCode() int32 {
    return PopupMenu_GetHashCode(p.instance)
}

func (p *TPopupMenu) ToString() string {
    return PopupMenu_ToString(p.instance)
}

func (p *TPopupMenu) Alignment() TPopupAlignment {
    return PopupMenu_GetAlignment(p.instance)
}

func (p *TPopupMenu) SetAlignment(value TPopupAlignment) {
    PopupMenu_SetAlignment(p.instance, value)
}

func (p *TPopupMenu) AutoHotkeys() TMenuAutoFlag {
    return PopupMenu_GetAutoHotkeys(p.instance)
}

func (p *TPopupMenu) SetAutoHotkeys(value TMenuAutoFlag) {
    PopupMenu_SetAutoHotkeys(p.instance, value)
}

func (p *TPopupMenu) BiDiMode() TBiDiMode {
    return PopupMenu_GetBiDiMode(p.instance)
}

func (p *TPopupMenu) SetBiDiMode(value TBiDiMode) {
    PopupMenu_SetBiDiMode(p.instance, value)
}

func (p *TPopupMenu) Images() *TImageList {
    return ImageListFromInst(PopupMenu_GetImages(p.instance))
}

func (p *TPopupMenu) SetImages(value IComponent) {
    PopupMenu_SetImages(p.instance, CheckPtr(value))
}

func (p *TPopupMenu) SetOnChange(fn TMenuChangeEvent) {
    PopupMenu_SetOnChange(p.instance, fn)
}

func (p *TPopupMenu) SetOnPopup(fn TNotifyEvent) {
    PopupMenu_SetOnPopup(p.instance, fn)
}

func (p *TPopupMenu) Handle() HMENU {
    return PopupMenu_GetHandle(p.instance)
}

func (p *TPopupMenu) WindowHandle() HWND {
    return PopupMenu_GetWindowHandle(p.instance)
}

func (p *TPopupMenu) SetWindowHandle(value HWND) {
    PopupMenu_SetWindowHandle(p.instance, value)
}

func (p *TPopupMenu) Items() *TMenuItem {
    return MenuItemFromInst(PopupMenu_GetItems(p.instance))
}

func (p *TPopupMenu) ComponentCount() int32 {
    return PopupMenu_GetComponentCount(p.instance)
}

func (p *TPopupMenu) ComponentIndex() int32 {
    return PopupMenu_GetComponentIndex(p.instance)
}

func (p *TPopupMenu) SetComponentIndex(value int32) {
    PopupMenu_SetComponentIndex(p.instance, value)
}

func (p *TPopupMenu) Owner() *TComponent {
    return ComponentFromInst(PopupMenu_GetOwner(p.instance))
}

func (p *TPopupMenu) Name() string {
    return PopupMenu_GetName(p.instance)
}

func (p *TPopupMenu) SetName(value string) {
    PopupMenu_SetName(p.instance, value)
}

func (p *TPopupMenu) Tag() int {
    return PopupMenu_GetTag(p.instance)
}

func (p *TPopupMenu) SetTag(value int) {
    PopupMenu_SetTag(p.instance, value)
}

func (p *TPopupMenu) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PopupMenu_GetComponents(p.instance, AIndex))
}

