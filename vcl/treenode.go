
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TTreeNode struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTreeNode
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTreeNode() *TTreeNode {
    t := new(TTreeNode)
    t.instance = TreeNode_Create()
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TreeNodeFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TreeNodeFromInst(inst uintptr) *TTreeNode {
    t := new(TTreeNode)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TreeNodeFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TreeNodeFromObj(obj IObject) *TTreeNode {
    t := new(TTreeNode)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TreeNodeFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TreeNodeFromUnsafePointer(ptr unsafe.Pointer) *TTreeNode {
    t := new(TTreeNode)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTreeNode) Free() {
    if t.instance != 0 {
        TreeNode_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTreeNode) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTreeNode) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTreeNode) IsValid() bool {
    return t.instance != 0
}

// TTreeNodeClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTreeNodeClass() TClass {
    return TreeNode_StaticClassType()
}

// AlphaSort
func (t *TTreeNode) AlphaSort(ARecurse bool) bool {
    return TreeNode_AlphaSort(t.instance, ARecurse)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTreeNode) Assign(Source IObject) {
    TreeNode_Assign(t.instance, CheckPtr(Source))
}

// Collapse
func (t *TTreeNode) Collapse(Recurse bool) {
    TreeNode_Collapse(t.instance, Recurse)
}

// Delete
func (t *TTreeNode) Delete() {
    TreeNode_Delete(t.instance)
}

// DisplayRect
func (t *TTreeNode) DisplayRect(TextOnly bool) TRect {
    return TreeNode_DisplayRect(t.instance, TextOnly)
}

// EditText
func (t *TTreeNode) EditText() bool {
    return TreeNode_EditText(t.instance)
}

// Expand
func (t *TTreeNode) Expand(Recurse bool) {
    TreeNode_Expand(t.instance, Recurse)
}

// IndexOf
func (t *TTreeNode) IndexOf(Value *TTreeNode) int32 {
    return TreeNode_IndexOf(t.instance, CheckPtr(Value))
}

// MakeVisible
func (t *TTreeNode) MakeVisible() {
    TreeNode_MakeVisible(t.instance)
}

// MoveTo
func (t *TTreeNode) MoveTo(Destination *TTreeNode, Mode TNodeAttachMode) {
    TreeNode_MoveTo(t.instance, CheckPtr(Destination), Mode)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTreeNode) GetNamePath() string {
    return TreeNode_GetNamePath(t.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTreeNode) DisposeOf() {
    TreeNode_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTreeNode) ClassType() TClass {
    return TreeNode_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTreeNode) ClassName() string {
    return TreeNode_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTreeNode) InstanceSize() int32 {
    return TreeNode_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTreeNode) InheritsFrom(AClass TClass) bool {
    return TreeNode_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTreeNode) Equals(Obj IObject) bool {
    return TreeNode_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTreeNode) GetHashCode() int32 {
    return TreeNode_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTreeNode) ToString() string {
    return TreeNode_ToString(t.instance)
}

// AbsoluteIndex
func (t *TTreeNode) AbsoluteIndex() int32 {
    return TreeNode_GetAbsoluteIndex(t.instance)
}

// Count
func (t *TTreeNode) Count() int32 {
    return TreeNode_GetCount(t.instance)
}

// Cut
func (t *TTreeNode) Cut() bool {
    return TreeNode_GetCut(t.instance)
}

// SetCut
func (t *TTreeNode) SetCut(value bool) {
    TreeNode_SetCut(t.instance, value)
}

// Data
func (t *TTreeNode) Data() unsafe.Pointer {
    return TreeNode_GetData(t.instance)
}

// SetData
func (t *TTreeNode) SetData(value unsafe.Pointer) {
    TreeNode_SetData(t.instance, value)
}

// Deleting
func (t *TTreeNode) Deleting() bool {
    return TreeNode_GetDeleting(t.instance)
}

// Focused
// CN: 获取返回是否获取焦点。
// EN: Get Return to get focus.
func (t *TTreeNode) Focused() bool {
    return TreeNode_GetFocused(t.instance)
}

// SetFocused
// CN: 设置返回是否获取焦点。
// EN: Set Return to get focus.
func (t *TTreeNode) SetFocused(value bool) {
    TreeNode_SetFocused(t.instance, value)
}

// DropTarget
func (t *TTreeNode) DropTarget() bool {
    return TreeNode_GetDropTarget(t.instance)
}

// SetDropTarget
func (t *TTreeNode) SetDropTarget(value bool) {
    TreeNode_SetDropTarget(t.instance, value)
}

// Selected
func (t *TTreeNode) Selected() bool {
    return TreeNode_GetSelected(t.instance)
}

// SetSelected
func (t *TTreeNode) SetSelected(value bool) {
    TreeNode_SetSelected(t.instance, value)
}

// Expanded
func (t *TTreeNode) Expanded() bool {
    return TreeNode_GetExpanded(t.instance)
}

// SetExpanded
func (t *TTreeNode) SetExpanded(value bool) {
    TreeNode_SetExpanded(t.instance, value)
}

// ExpandedImageIndex
func (t *TTreeNode) ExpandedImageIndex() int32 {
    return TreeNode_GetExpandedImageIndex(t.instance)
}

// SetExpandedImageIndex
func (t *TTreeNode) SetExpandedImageIndex(value int32) {
    TreeNode_SetExpandedImageIndex(t.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (t *TTreeNode) Handle() HWND {
    return TreeNode_GetHandle(t.instance)
}

// HasChildren
func (t *TTreeNode) HasChildren() bool {
    return TreeNode_GetHasChildren(t.instance)
}

// SetHasChildren
func (t *TTreeNode) SetHasChildren(value bool) {
    TreeNode_SetHasChildren(t.instance, value)
}

// ImageIndex
func (t *TTreeNode) ImageIndex() int32 {
    return TreeNode_GetImageIndex(t.instance)
}

// SetImageIndex
func (t *TTreeNode) SetImageIndex(value int32) {
    TreeNode_SetImageIndex(t.instance, value)
}

// Index
func (t *TTreeNode) Index() int32 {
    return TreeNode_GetIndex(t.instance)
}

// IsVisible
func (t *TTreeNode) IsVisible() bool {
    return TreeNode_GetIsVisible(t.instance)
}

// ItemId
func (t *TTreeNode) ItemId() uintptr {
    return TreeNode_GetItemId(t.instance)
}

// Level
func (t *TTreeNode) Level() int32 {
    return TreeNode_GetLevel(t.instance)
}

// OverlayIndex
func (t *TTreeNode) OverlayIndex() int32 {
    return TreeNode_GetOverlayIndex(t.instance)
}

// SetOverlayIndex
func (t *TTreeNode) SetOverlayIndex(value int32) {
    TreeNode_SetOverlayIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTreeNode) Owner() *TTreeNodes {
    return TreeNodesFromInst(TreeNode_GetOwner(t.instance))
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TTreeNode) Parent() *TTreeNode {
    return TreeNodeFromInst(TreeNode_GetParent(t.instance))
}

// SelectedIndex
func (t *TTreeNode) SelectedIndex() int32 {
    return TreeNode_GetSelectedIndex(t.instance)
}

// SetSelectedIndex
func (t *TTreeNode) SetSelectedIndex(value int32) {
    TreeNode_SetSelectedIndex(t.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TTreeNode) Enabled() bool {
    return TreeNode_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TTreeNode) SetEnabled(value bool) {
    TreeNode_SetEnabled(t.instance, value)
}

// StateIndex
func (t *TTreeNode) StateIndex() int32 {
    return TreeNode_GetStateIndex(t.instance)
}

// SetStateIndex
func (t *TTreeNode) SetStateIndex(value int32) {
    TreeNode_SetStateIndex(t.instance, value)
}

// Text
func (t *TTreeNode) Text() string {
    return TreeNode_GetText(t.instance)
}

// SetText
func (t *TTreeNode) SetText(value string) {
    TreeNode_SetText(t.instance, value)
}

// Item
func (t *TTreeNode) Item(Index int32) *TTreeNode {
    return TreeNodeFromInst(TreeNode_GetItem(t.instance, Index))
}

// Item
func (t *TTreeNode) SetItem(Index int32, value *TTreeNode) {
    TreeNode_SetItem(t.instance, Index, CheckPtr(value))
}

