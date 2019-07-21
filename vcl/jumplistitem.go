
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

type TJumpListItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewJumpListItem
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewJumpListItem() *TJumpListItem {
    j := new(TJumpListItem)
    j.instance = JumpListItem_Create()
    j.ptr = unsafe.Pointer(j.instance)
    return j
}

// JumpListItemFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func JumpListItemFromInst(inst uintptr) *TJumpListItem {
    j := new(TJumpListItem)
    j.instance = inst
    j.ptr = unsafe.Pointer(inst)
    return j
}

// JumpListItemFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func JumpListItemFromObj(obj IObject) *TJumpListItem {
    j := new(TJumpListItem)
    j.instance = CheckPtr(obj)
    j.ptr = unsafe.Pointer(j.instance)
    return j
}

// JumpListItemFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func JumpListItemFromUnsafePointer(ptr unsafe.Pointer) *TJumpListItem {
    j := new(TJumpListItem)
    j.instance = uintptr(ptr)
    j.ptr = ptr
    return j
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (j *TJumpListItem) Free() {
    if j.instance != 0 {
        JumpListItem_Free(j.instance)
        j.instance = 0
        j.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (j *TJumpListItem) Instance() uintptr {
    return j.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (j *TJumpListItem) UnsafeAddr() unsafe.Pointer {
    return j.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (j *TJumpListItem) IsValid() bool {
    return j.instance != 0
}

// TJumpListItemClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TJumpListItemClass() TClass {
    return JumpListItem_StaticClassType()
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (j *TJumpListItem) GetNamePath() string {
    return JumpListItem_GetNamePath(j.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (j *TJumpListItem) Assign(Source IObject) {
    JumpListItem_Assign(j.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (j *TJumpListItem) DisposeOf() {
    JumpListItem_DisposeOf(j.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (j *TJumpListItem) ClassType() TClass {
    return JumpListItem_ClassType(j.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (j *TJumpListItem) ClassName() string {
    return JumpListItem_ClassName(j.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (j *TJumpListItem) InstanceSize() int32 {
    return JumpListItem_InstanceSize(j.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (j *TJumpListItem) InheritsFrom(AClass TClass) bool {
    return JumpListItem_InheritsFrom(j.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (j *TJumpListItem) Equals(Obj IObject) bool {
    return JumpListItem_Equals(j.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (j *TJumpListItem) GetHashCode() int32 {
    return JumpListItem_GetHashCode(j.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (j *TJumpListItem) ToString() string {
    return JumpListItem_ToString(j.instance)
}

// IsSeparator
func (j *TJumpListItem) IsSeparator() bool {
    return JumpListItem_GetIsSeparator(j.instance)
}

// SetIsSeparator
func (j *TJumpListItem) SetIsSeparator(value bool) {
    JumpListItem_SetIsSeparator(j.instance, value)
}

// Icon
// CN: 获取图标。
// EN: Get icon.
func (j *TJumpListItem) Icon() string {
    return JumpListItem_GetIcon(j.instance)
}

// SetIcon
// CN: 设置图标。
// EN: Set icon.
func (j *TJumpListItem) SetIcon(value string) {
    JumpListItem_SetIcon(j.instance, value)
}

// Path
func (j *TJumpListItem) Path() string {
    return JumpListItem_GetPath(j.instance)
}

// SetPath
func (j *TJumpListItem) SetPath(value string) {
    JumpListItem_SetPath(j.instance, value)
}

// FriendlyName
func (j *TJumpListItem) FriendlyName() string {
    return JumpListItem_GetFriendlyName(j.instance)
}

// SetFriendlyName
func (j *TJumpListItem) SetFriendlyName(value string) {
    JumpListItem_SetFriendlyName(j.instance, value)
}

// Arguments
func (j *TJumpListItem) Arguments() string {
    return JumpListItem_GetArguments(j.instance)
}

// SetArguments
func (j *TJumpListItem) SetArguments(value string) {
    JumpListItem_SetArguments(j.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (j *TJumpListItem) Visible() bool {
    return JumpListItem_GetVisible(j.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (j *TJumpListItem) SetVisible(value bool) {
    JumpListItem_SetVisible(j.instance, value)
}

// Collection
func (j *TJumpListItem) Collection() *TCollection {
    return CollectionFromInst(JumpListItem_GetCollection(j.instance))
}

// SetCollection
func (j *TJumpListItem) SetCollection(value *TCollection) {
    JumpListItem_SetCollection(j.instance, CheckPtr(value))
}

// Index
func (j *TJumpListItem) Index() int32 {
    return JumpListItem_GetIndex(j.instance)
}

// SetIndex
func (j *TJumpListItem) SetIndex(value int32) {
    JumpListItem_SetIndex(j.instance, value)
}

// DisplayName
func (j *TJumpListItem) DisplayName() string {
    return JumpListItem_GetDisplayName(j.instance)
}

// SetDisplayName
func (j *TJumpListItem) SetDisplayName(value string) {
    JumpListItem_SetDisplayName(j.instance, value)
}

