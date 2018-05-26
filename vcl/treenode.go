
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

type TTreeNode struct {
    IObject
    instance uintptr
}

func NewTreeNode() *TTreeNode {
    t := new(TTreeNode)
    t.instance = TreeNode_Create()
    return t
}

func TreeNodeFromInst(inst uintptr) *TTreeNode {
    t := new(TTreeNode)
    t.instance = inst
    return t
}

func TreeNodeFromObj(obj IObject) *TTreeNode {
    t := new(TTreeNode)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTreeNode) Free() {
    if t.instance != 0 {
        TreeNode_Free(t.instance)
        t.instance = 0
    }
}

func (t *TTreeNode) Instance() uintptr {
    return t.instance
}

func (t *TTreeNode) IsValid() bool {
    return t.instance != 0
}

func (t *TTreeNode) AlphaSort(ARecurse bool) bool {
    return TreeNode_AlphaSort(t.instance, ARecurse)
}

func (t *TTreeNode) Assign(Source IObject) {
    TreeNode_Assign(t.instance, CheckPtr(Source))
}

func (t *TTreeNode) Collapse(Recurse bool) {
    TreeNode_Collapse(t.instance, Recurse)
}

func (t *TTreeNode) Delete() {
    TreeNode_Delete(t.instance)
}

func (t *TTreeNode) DisplayRect(TextOnly bool) TRect {
    return TreeNode_DisplayRect(t.instance, TextOnly)
}

func (t *TTreeNode) EditText() bool {
    return TreeNode_EditText(t.instance)
}

func (t *TTreeNode) Expand(Recurse bool) {
    TreeNode_Expand(t.instance, Recurse)
}

func (t *TTreeNode) IndexOf(Value *TTreeNode) int32 {
    return TreeNode_IndexOf(t.instance, CheckPtr(Value))
}

func (t *TTreeNode) MakeVisible() {
    TreeNode_MakeVisible(t.instance)
}

func (t *TTreeNode) MoveTo(Destination *TTreeNode, Mode TNodeAttachMode) {
    TreeNode_MoveTo(t.instance, CheckPtr(Destination), Mode)
}

func (t *TTreeNode) GetNamePath() string {
    return TreeNode_GetNamePath(t.instance)
}

func (t *TTreeNode) ClassName() string {
    return TreeNode_ClassName(t.instance)
}

func (t *TTreeNode) Equals(Obj IObject) bool {
    return TreeNode_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTreeNode) GetHashCode() int32 {
    return TreeNode_GetHashCode(t.instance)
}

func (t *TTreeNode) ToString() string {
    return TreeNode_ToString(t.instance)
}

func (t *TTreeNode) AbsoluteIndex() int32 {
    return TreeNode_GetAbsoluteIndex(t.instance)
}

func (t *TTreeNode) Count() int32 {
    return TreeNode_GetCount(t.instance)
}

func (t *TTreeNode) Cut() bool {
    return TreeNode_GetCut(t.instance)
}

func (t *TTreeNode) SetCut(value bool) {
    TreeNode_SetCut(t.instance, value)
}

func (t *TTreeNode) Data() uintptr {
    return TreeNode_GetData(t.instance)
}

func (t *TTreeNode) SetData(value uintptr) {
    TreeNode_SetData(t.instance, value)
}

func (t *TTreeNode) Deleting() bool {
    return TreeNode_GetDeleting(t.instance)
}

func (t *TTreeNode) Focused() bool {
    return TreeNode_GetFocused(t.instance)
}

func (t *TTreeNode) SetFocused(value bool) {
    TreeNode_SetFocused(t.instance, value)
}

func (t *TTreeNode) DropTarget() bool {
    return TreeNode_GetDropTarget(t.instance)
}

func (t *TTreeNode) SetDropTarget(value bool) {
    TreeNode_SetDropTarget(t.instance, value)
}

func (t *TTreeNode) Selected() bool {
    return TreeNode_GetSelected(t.instance)
}

func (t *TTreeNode) SetSelected(value bool) {
    TreeNode_SetSelected(t.instance, value)
}

func (t *TTreeNode) Expanded() bool {
    return TreeNode_GetExpanded(t.instance)
}

func (t *TTreeNode) SetExpanded(value bool) {
    TreeNode_SetExpanded(t.instance, value)
}

func (t *TTreeNode) ExpandedImageIndex() int32 {
    return TreeNode_GetExpandedImageIndex(t.instance)
}

func (t *TTreeNode) SetExpandedImageIndex(value int32) {
    TreeNode_SetExpandedImageIndex(t.instance, value)
}

func (t *TTreeNode) Handle() HWND {
    return TreeNode_GetHandle(t.instance)
}

func (t *TTreeNode) HasChildren() bool {
    return TreeNode_GetHasChildren(t.instance)
}

func (t *TTreeNode) SetHasChildren(value bool) {
    TreeNode_SetHasChildren(t.instance, value)
}

func (t *TTreeNode) ImageIndex() int32 {
    return TreeNode_GetImageIndex(t.instance)
}

func (t *TTreeNode) SetImageIndex(value int32) {
    TreeNode_SetImageIndex(t.instance, value)
}

func (t *TTreeNode) Index() int32 {
    return TreeNode_GetIndex(t.instance)
}

func (t *TTreeNode) IsVisible() bool {
    return TreeNode_GetIsVisible(t.instance)
}

func (t *TTreeNode) ItemId() uintptr {
    return TreeNode_GetItemId(t.instance)
}

func (t *TTreeNode) Level() int32 {
    return TreeNode_GetLevel(t.instance)
}

func (t *TTreeNode) OverlayIndex() int32 {
    return TreeNode_GetOverlayIndex(t.instance)
}

func (t *TTreeNode) SetOverlayIndex(value int32) {
    TreeNode_SetOverlayIndex(t.instance, value)
}

func (t *TTreeNode) Owner() *TTreeNodes {
    return TreeNodesFromInst(TreeNode_GetOwner(t.instance))
}

func (t *TTreeNode) Parent() *TTreeNode {
    return TreeNodeFromInst(TreeNode_GetParent(t.instance))
}

func (t *TTreeNode) SelectedIndex() int32 {
    return TreeNode_GetSelectedIndex(t.instance)
}

func (t *TTreeNode) SetSelectedIndex(value int32) {
    TreeNode_SetSelectedIndex(t.instance, value)
}

func (t *TTreeNode) Enabled() bool {
    return TreeNode_GetEnabled(t.instance)
}

func (t *TTreeNode) SetEnabled(value bool) {
    TreeNode_SetEnabled(t.instance, value)
}

func (t *TTreeNode) StateIndex() int32 {
    return TreeNode_GetStateIndex(t.instance)
}

func (t *TTreeNode) SetStateIndex(value int32) {
    TreeNode_SetStateIndex(t.instance, value)
}

func (t *TTreeNode) Text() string {
    return TreeNode_GetText(t.instance)
}

func (t *TTreeNode) SetText(value string) {
    TreeNode_SetText(t.instance, value)
}

func (t *TTreeNode) Item(Index int32) *TTreeNode {
    return TreeNodeFromInst(TreeNode_GetItem(t.instance, Index))
}

func (t *TTreeNode) SetItem(Index int32, value *TTreeNode) {
    TreeNode_SetItem(t.instance, Index, CheckPtr(value))
}

