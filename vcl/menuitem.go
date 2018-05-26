
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

type TMenuItem struct {
    IComponent
    instance uintptr
}

func NewMenuItem(owner IComponent) *TMenuItem {
    m := new(TMenuItem)
    m.instance = MenuItem_Create(CheckPtr(owner))
    return m
}

func MenuItemFromInst(inst uintptr) *TMenuItem {
    m := new(TMenuItem)
    m.instance = inst
    return m
}

func MenuItemFromObj(obj IObject) *TMenuItem {
    m := new(TMenuItem)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMenuItem) Free() {
    if m.instance != 0 {
        MenuItem_Free(m.instance)
        m.instance = 0
    }
}

func (m *TMenuItem) Instance() uintptr {
    return m.instance
}

func (m *TMenuItem) IsValid() bool {
    return m.instance != 0
}

func (m *TMenuItem) Insert(Index int32, Item IComponent) {
    MenuItem_Insert(m.instance, Index , CheckPtr(Item))
}

func (m *TMenuItem) Delete(Index int32) {
    MenuItem_Delete(m.instance, Index)
}

func (m *TMenuItem) Clear() {
    MenuItem_Clear(m.instance)
}

func (m *TMenuItem) Click() {
    MenuItem_Click(m.instance)
}

func (m *TMenuItem) IndexOf(Item IComponent) int32 {
    return MenuItem_IndexOf(m.instance, CheckPtr(Item))
}

func (m *TMenuItem) HasParent() bool {
    return MenuItem_HasParent(m.instance)
}

func (m *TMenuItem) Add(Item IComponent) {
    MenuItem_Add(m.instance, CheckPtr(Item))
}

func (m *TMenuItem) FindComponent(AName string) *TComponent {
    return ComponentFromInst(MenuItem_FindComponent(m.instance, AName))
}

func (m *TMenuItem) GetNamePath() string {
    return MenuItem_GetNamePath(m.instance)
}

func (m *TMenuItem) Assign(Source IObject) {
    MenuItem_Assign(m.instance, CheckPtr(Source))
}

func (m *TMenuItem) ClassName() string {
    return MenuItem_ClassName(m.instance)
}

func (m *TMenuItem) Equals(Obj IObject) bool {
    return MenuItem_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMenuItem) GetHashCode() int32 {
    return MenuItem_GetHashCode(m.instance)
}

func (m *TMenuItem) ToString() string {
    return MenuItem_ToString(m.instance)
}

func (m *TMenuItem) Handle() HMENU {
    return MenuItem_GetHandle(m.instance)
}

func (m *TMenuItem) Count() int32 {
    return MenuItem_GetCount(m.instance)
}

func (m *TMenuItem) Parent() *TMenuItem {
    return MenuItemFromInst(MenuItem_GetParent(m.instance))
}

func (m *TMenuItem) Action() *TAction {
    return ActionFromInst(MenuItem_GetAction(m.instance))
}

func (m *TMenuItem) SetAction(value IComponent) {
    MenuItem_SetAction(m.instance, CheckPtr(value))
}

func (m *TMenuItem) AutoHotkeys() TMenuItemAutoFlag {
    return MenuItem_GetAutoHotkeys(m.instance)
}

func (m *TMenuItem) SetAutoHotkeys(value TMenuItemAutoFlag) {
    MenuItem_SetAutoHotkeys(m.instance, value)
}

func (m *TMenuItem) Bitmap() *TBitmap {
    return BitmapFromInst(MenuItem_GetBitmap(m.instance))
}

func (m *TMenuItem) SetBitmap(value *TBitmap) {
    MenuItem_SetBitmap(m.instance, CheckPtr(value))
}

func (m *TMenuItem) Caption() string {
    return MenuItem_GetCaption(m.instance)
}

func (m *TMenuItem) SetCaption(value string) {
    MenuItem_SetCaption(m.instance, value)
}

func (m *TMenuItem) Checked() bool {
    return MenuItem_GetChecked(m.instance)
}

func (m *TMenuItem) SetChecked(value bool) {
    MenuItem_SetChecked(m.instance, value)
}

func (m *TMenuItem) Default() bool {
    return MenuItem_GetDefault(m.instance)
}

func (m *TMenuItem) SetDefault(value bool) {
    MenuItem_SetDefault(m.instance, value)
}

func (m *TMenuItem) Enabled() bool {
    return MenuItem_GetEnabled(m.instance)
}

func (m *TMenuItem) SetEnabled(value bool) {
    MenuItem_SetEnabled(m.instance, value)
}

func (m *TMenuItem) GroupIndex() uint8 {
    return MenuItem_GetGroupIndex(m.instance)
}

func (m *TMenuItem) SetGroupIndex(value uint8) {
    MenuItem_SetGroupIndex(m.instance, value)
}

func (m *TMenuItem) Hint() string {
    return MenuItem_GetHint(m.instance)
}

func (m *TMenuItem) SetHint(value string) {
    MenuItem_SetHint(m.instance, value)
}

func (m *TMenuItem) ImageIndex() int32 {
    return MenuItem_GetImageIndex(m.instance)
}

func (m *TMenuItem) SetImageIndex(value int32) {
    MenuItem_SetImageIndex(m.instance, value)
}

func (m *TMenuItem) ShortCut() TShortCut {
    return MenuItem_GetShortCut(m.instance)
}

func (m *TMenuItem) SetShortCut(value TShortCut) {
    MenuItem_SetShortCut(m.instance, value)
}

func (m *TMenuItem) Visible() bool {
    return MenuItem_GetVisible(m.instance)
}

func (m *TMenuItem) SetVisible(value bool) {
    MenuItem_SetVisible(m.instance, value)
}

func (m *TMenuItem) SetOnClick(fn TNotifyEvent) {
    MenuItem_SetOnClick(m.instance, fn)
}

func (m *TMenuItem) SetOnDrawItem(fn TMenuDrawItemEvent) {
    MenuItem_SetOnDrawItem(m.instance, fn)
}

func (m *TMenuItem) ComponentCount() int32 {
    return MenuItem_GetComponentCount(m.instance)
}

func (m *TMenuItem) ComponentIndex() int32 {
    return MenuItem_GetComponentIndex(m.instance)
}

func (m *TMenuItem) SetComponentIndex(value int32) {
    MenuItem_SetComponentIndex(m.instance, value)
}

func (m *TMenuItem) Owner() *TComponent {
    return ComponentFromInst(MenuItem_GetOwner(m.instance))
}

func (m *TMenuItem) Name() string {
    return MenuItem_GetName(m.instance)
}

func (m *TMenuItem) SetName(value string) {
    MenuItem_SetName(m.instance, value)
}

func (m *TMenuItem) Tag() int {
    return MenuItem_GetTag(m.instance)
}

func (m *TMenuItem) SetTag(value int) {
    MenuItem_SetTag(m.instance, value)
}

func (m *TMenuItem) Items(Index int32) *TMenuItem {
    return MenuItemFromInst(MenuItem_GetItems(m.instance, Index))
}

func (m *TMenuItem) Components(AIndex int32) *TComponent {
    return ComponentFromInst(MenuItem_GetComponents(m.instance, AIndex))
}

