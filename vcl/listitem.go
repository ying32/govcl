
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

type TListItem struct {
    IObject
    instance uintptr
}

func NewListItem() *TListItem {
    l := new(TListItem)
    l.instance = ListItem_Create()
    return l
}

func ListItemFromInst(inst uintptr) *TListItem {
    l := new(TListItem)
    l.instance = inst
    return l
}

func ListItemFromObj(obj IObject) *TListItem {
    l := new(TListItem)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListItem) Free() {
    if l.instance != 0 {
        ListItem_Free(l.instance)
        l.instance = 0
    }
}

func (l *TListItem) Instance() uintptr {
    return l.instance
}

func (l *TListItem) IsValid() bool {
    return l.instance != 0
}

func (l *TListItem) Assign(Source IObject) {
    ListItem_Assign(l.instance, CheckPtr(Source))
}

func (l *TListItem) CancelEdit() {
    ListItem_CancelEdit(l.instance)
}

func (l *TListItem) Delete() {
    ListItem_Delete(l.instance)
}

func (l *TListItem) DisplayRect(Code TDisplayCode) TRect {
    return ListItem_DisplayRect(l.instance, Code)
}

func (l *TListItem) EditCaption() bool {
    return ListItem_EditCaption(l.instance)
}

func (l *TListItem) MakeVisible(PartialOK bool) {
    ListItem_MakeVisible(l.instance, PartialOK)
}

func (l *TListItem) Update() {
    ListItem_Update(l.instance)
}

func (l *TListItem) WorkArea() int32 {
    return ListItem_WorkArea(l.instance)
}

func (l *TListItem) GetNamePath() string {
    return ListItem_GetNamePath(l.instance)
}

func (l *TListItem) ClassName() string {
    return ListItem_ClassName(l.instance)
}

func (l *TListItem) Equals(Obj IObject) bool {
    return ListItem_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListItem) GetHashCode() int32 {
    return ListItem_GetHashCode(l.instance)
}

func (l *TListItem) ToString() string {
    return ListItem_ToString(l.instance)
}

func (l *TListItem) Caption() string {
    return ListItem_GetCaption(l.instance)
}

func (l *TListItem) SetCaption(value string) {
    ListItem_SetCaption(l.instance, value)
}

func (l *TListItem) Checked() bool {
    return ListItem_GetChecked(l.instance)
}

func (l *TListItem) SetChecked(value bool) {
    ListItem_SetChecked(l.instance, value)
}

func (l *TListItem) Cut() bool {
    return ListItem_GetCut(l.instance)
}

func (l *TListItem) SetCut(value bool) {
    ListItem_SetCut(l.instance, value)
}

func (l *TListItem) Data() uintptr {
    return ListItem_GetData(l.instance)
}

func (l *TListItem) SetData(value uintptr) {
    ListItem_SetData(l.instance, value)
}

func (l *TListItem) Deleting() bool {
    return ListItem_GetDeleting(l.instance)
}

func (l *TListItem) DropTarget() bool {
    return ListItem_GetDropTarget(l.instance)
}

func (l *TListItem) SetDropTarget(value bool) {
    ListItem_SetDropTarget(l.instance, value)
}

func (l *TListItem) Focused() bool {
    return ListItem_GetFocused(l.instance)
}

func (l *TListItem) SetFocused(value bool) {
    ListItem_SetFocused(l.instance, value)
}

func (l *TListItem) GroupID() int32 {
    return ListItem_GetGroupID(l.instance)
}

func (l *TListItem) SetGroupID(value int32) {
    ListItem_SetGroupID(l.instance, value)
}

func (l *TListItem) Handle() HWND {
    return ListItem_GetHandle(l.instance)
}

func (l *TListItem) ImageIndex() int32 {
    return ListItem_GetImageIndex(l.instance)
}

func (l *TListItem) SetImageIndex(value int32) {
    ListItem_SetImageIndex(l.instance, value)
}

func (l *TListItem) Indent() int32 {
    return ListItem_GetIndent(l.instance)
}

func (l *TListItem) SetIndent(value int32) {
    ListItem_SetIndent(l.instance, value)
}

func (l *TListItem) Index() int32 {
    return ListItem_GetIndex(l.instance)
}

func (l *TListItem) Left() int32 {
    return ListItem_GetLeft(l.instance)
}

func (l *TListItem) SetLeft(value int32) {
    ListItem_SetLeft(l.instance, value)
}

func (l *TListItem) Owner() *TListItems {
    return ListItemsFromInst(ListItem_GetOwner(l.instance))
}

func (l *TListItem) OverlayIndex() int32 {
    return ListItem_GetOverlayIndex(l.instance)
}

func (l *TListItem) SetOverlayIndex(value int32) {
    ListItem_SetOverlayIndex(l.instance, value)
}

func (l *TListItem) Position() TPoint {
    return ListItem_GetPosition(l.instance)
}

func (l *TListItem) SetPosition(value TPoint) {
    ListItem_SetPosition(l.instance, value)
}

func (l *TListItem) Selected() bool {
    return ListItem_GetSelected(l.instance)
}

func (l *TListItem) SetSelected(value bool) {
    ListItem_SetSelected(l.instance, value)
}

func (l *TListItem) StateIndex() int32 {
    return ListItem_GetStateIndex(l.instance)
}

func (l *TListItem) SetStateIndex(value int32) {
    ListItem_SetStateIndex(l.instance, value)
}

func (l *TListItem) SubItems() *TStrings {
    return StringsFromInst(ListItem_GetSubItems(l.instance))
}

func (l *TListItem) SetSubItems(value IObject) {
    ListItem_SetSubItems(l.instance, CheckPtr(value))
}

func (l *TListItem) Top() int32 {
    return ListItem_GetTop(l.instance)
}

func (l *TListItem) SetTop(value int32) {
    ListItem_SetTop(l.instance, value)
}

func (l *TListItem) SubItemImages(Index int32) int32 {
    return ListItem_GetSubItemImages(l.instance, Index)
}

func (l *TListItem) SetSubItemImages(Index int32, value int32) {
    ListItem_SetSubItemImages(l.instance, Index, value)
}

