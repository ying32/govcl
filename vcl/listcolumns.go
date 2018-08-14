
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

type TListColumns struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewListColumns
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListColumns() *TListColumns {
    l := new(TListColumns)
    l.instance = ListColumns_Create()
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListColumnsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ListColumnsFromInst(inst uintptr) *TListColumns {
    l := new(TListColumns)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// ListColumnsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ListColumnsFromObj(obj IObject) *TListColumns {
    l := new(TListColumns)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListColumnsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ListColumnsFromUnsafePointer(ptr unsafe.Pointer) *TListColumns {
    l := new(TListColumns)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TListColumns) Free() {
    if l.instance != 0 {
        ListColumns_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListColumns) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListColumns) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListColumns) IsValid() bool {
    return l.instance != 0
}

// TListColumnsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListColumnsClass() TClass {
    return ListColumns_StaticClassType()
}

// Add
func (l *TListColumns) Add() *TListColumn {
    return ListColumnFromInst(ListColumns_Add(l.instance))
}

// Owner
// CN: 组件所有者。
// EN: component owner.
func (l *TListColumns) Owner() *TWinControl {
    return WinControlFromInst(ListColumns_Owner(l.instance))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TListColumns) Assign(Source IObject) {
    ListColumns_Assign(l.instance, CheckPtr(Source))
}

// BeginUpdate
func (l *TListColumns) BeginUpdate() {
    ListColumns_BeginUpdate(l.instance)
}

// Clear
// CN: 清除。
// EN: .
func (l *TListColumns) Clear() {
    ListColumns_Clear(l.instance)
}

// Delete
func (l *TListColumns) Delete(Index int32) {
    ListColumns_Delete(l.instance, Index)
}

// EndUpdate
func (l *TListColumns) EndUpdate() {
    ListColumns_EndUpdate(l.instance)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TListColumns) GetNamePath() string {
    return ListColumns_GetNamePath(l.instance)
}

// Insert
func (l *TListColumns) Insert(Index int32) *TCollectionItem {
    return CollectionItemFromInst(ListColumns_Insert(l.instance, Index))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TListColumns) DisposeOf() {
    ListColumns_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListColumns) ClassType() TClass {
    return ListColumns_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListColumns) ClassName() string {
    return ListColumns_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListColumns) InstanceSize() int32 {
    return ListColumns_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListColumns) InheritsFrom(AClass TClass) bool {
    return ListColumns_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListColumns) Equals(Obj IObject) bool {
    return ListColumns_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListColumns) GetHashCode() int32 {
    return ListColumns_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TListColumns) ToString() string {
    return ListColumns_ToString(l.instance)
}

// Capacity
func (l *TListColumns) Capacity() int32 {
    return ListColumns_GetCapacity(l.instance)
}

// SetCapacity
func (l *TListColumns) SetCapacity(value int32) {
    ListColumns_SetCapacity(l.instance, value)
}

// Count
func (l *TListColumns) Count() int32 {
    return ListColumns_GetCount(l.instance)
}

// Items
func (l *TListColumns) Items(Index int32) *TListColumn {
    return ListColumnFromInst(ListColumns_GetItems(l.instance, Index))
}

// Items
func (l *TListColumns) SetItems(Index int32, value *TListColumn) {
    ListColumns_SetItems(l.instance, Index, CheckPtr(value))
}

