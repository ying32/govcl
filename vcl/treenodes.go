
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

type TTreeNodes struct {
    IObject
    instance uintptr
}

func NewTreeNodes() *TTreeNodes {
    t := new(TTreeNodes)
    t.instance = TreeNodes_Create()
    return t
}

func TreeNodesFromInst(inst uintptr) *TTreeNodes {
    t := new(TTreeNodes)
    t.instance = inst
    return t
}

func TreeNodesFromObj(obj IObject) *TTreeNodes {
    t := new(TTreeNodes)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTreeNodes) Free() {
    if t.instance != 0 {
        TreeNodes_Free(t.instance)
        t.instance = 0
    }
}

func (t *TTreeNodes) Instance() uintptr {
    return t.instance
}

func (t *TTreeNodes) IsValid() bool {
    return t.instance != 0
}

func (t *TTreeNodes) AddChildFirst(Parent *TTreeNode, S string) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddChildFirst(t.instance, CheckPtr(Parent), S))
}

func (t *TTreeNodes) AddChild(Parent *TTreeNode, S string) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddChild(t.instance, CheckPtr(Parent), S))
}

func (t *TTreeNodes) AddChildObjectFirst(Parent *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddChildObjectFirst(t.instance, CheckPtr(Parent), S , Ptr))
}

func (t *TTreeNodes) AddChildObject(Parent *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddChildObject(t.instance, CheckPtr(Parent), S , Ptr))
}

func (t *TTreeNodes) AddObjectFirst(Sibling *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddObjectFirst(t.instance, CheckPtr(Sibling), S , Ptr))
}

func (t *TTreeNodes) AddObject(Sibling *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddObject(t.instance, CheckPtr(Sibling), S , Ptr))
}

func (t *TTreeNodes) AddNode(Node *TTreeNode, Relative *TTreeNode, S string, Ptr uintptr, Method TNodeAttachMode) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddNode(t.instance, CheckPtr(Node), CheckPtr(Relative), S , Ptr , Method))
}

func (t *TTreeNodes) AddFirst(Sibling *TTreeNode, S string) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddFirst(t.instance, CheckPtr(Sibling), S))
}

func (t *TTreeNodes) Add(Sibling *TTreeNode, S string) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_Add(t.instance, CheckPtr(Sibling), S))
}

func (t *TTreeNodes) AlphaSort(ARecurse bool) bool {
    return TreeNodes_AlphaSort(t.instance, ARecurse)
}

func (t *TTreeNodes) Assign(Source IObject) {
    TreeNodes_Assign(t.instance, CheckPtr(Source))
}

func (t *TTreeNodes) BeginUpdate() {
    TreeNodes_BeginUpdate(t.instance)
}

func (t *TTreeNodes) Clear() {
    TreeNodes_Clear(t.instance)
}

func (t *TTreeNodes) Delete(Node *TTreeNode) {
    TreeNodes_Delete(t.instance, CheckPtr(Node))
}

func (t *TTreeNodes) EndUpdate() {
    TreeNodes_EndUpdate(t.instance)
}

func (t *TTreeNodes) GetFirstNode() *TTreeNode {
    return TreeNodeFromInst(TreeNodes_GetFirstNode(t.instance))
}

func (t *TTreeNodes) GetNode(ItemId uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_GetNode(t.instance, ItemId))
}

func (t *TTreeNodes) Insert(Sibling *TTreeNode, S string) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_Insert(t.instance, CheckPtr(Sibling), S))
}

func (t *TTreeNodes) InsertObject(Sibling *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_InsertObject(t.instance, CheckPtr(Sibling), S , Ptr))
}

func (t *TTreeNodes) GetNamePath() string {
    return TreeNodes_GetNamePath(t.instance)
}

func (t *TTreeNodes) ClassName() string {
    return TreeNodes_ClassName(t.instance)
}

func (t *TTreeNodes) Equals(Obj IObject) bool {
    return TreeNodes_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTreeNodes) GetHashCode() int32 {
    return TreeNodes_GetHashCode(t.instance)
}

func (t *TTreeNodes) ToString() string {
    return TreeNodes_ToString(t.instance)
}

func (t *TTreeNodes) Count() int32 {
    return TreeNodes_GetCount(t.instance)
}

func (t *TTreeNodes) Handle() HWND {
    return TreeNodes_GetHandle(t.instance)
}

func (t *TTreeNodes) Owner() *TControl {
    return ControlFromInst(TreeNodes_GetOwner(t.instance))
}

func (t *TTreeNodes) Item(Index int32) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_GetItem(t.instance, Index))
}

