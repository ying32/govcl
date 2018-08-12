
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

type TStatusPanels struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewStatusPanels
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewStatusPanels() *TStatusPanels {
    s := new(TStatusPanels)
    s.instance = StatusPanels_Create()
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StatusPanelsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func StatusPanelsFromInst(inst uintptr) *TStatusPanels {
    s := new(TStatusPanels)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// StatusPanelsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func StatusPanelsFromObj(obj IObject) *TStatusPanels {
    s := new(TStatusPanels)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StatusPanelsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func StatusPanelsFromUnsafePointer(ptr unsafe.Pointer) *TStatusPanels {
    s := new(TStatusPanels)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TStatusPanels) Free() {
    if s.instance != 0 {
        StatusPanels_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TStatusPanels) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TStatusPanels) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TStatusPanels) IsValid() bool {
    return s.instance != 0
}

// TStatusPanelsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TStatusPanelsClass() TClass {
    return StatusPanels_StaticClassType()
}

// Add
func (s *TStatusPanels) Add() *TStatusPanel {
    return StatusPanelFromInst(StatusPanels_Add(s.instance))
}

// AddItem
func (s *TStatusPanels) AddItem(Item *TStatusPanel, Index int32) *TStatusPanel {
    return StatusPanelFromInst(StatusPanels_AddItem(s.instance, CheckPtr(Item), Index))
}

// Insert
func (s *TStatusPanels) Insert(Index int32) *TStatusPanel {
    return StatusPanelFromInst(StatusPanels_Insert(s.instance, Index))
}

// Owner
// CN: 组件所有者。
// EN: component owner.
func (s *TStatusPanels) Owner() *TObject {
    return ObjectFromInst(StatusPanels_Owner(s.instance))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TStatusPanels) Assign(Source IObject) {
    StatusPanels_Assign(s.instance, CheckPtr(Source))
}

// BeginUpdate
func (s *TStatusPanels) BeginUpdate() {
    StatusPanels_BeginUpdate(s.instance)
}

// Clear
func (s *TStatusPanels) Clear() {
    StatusPanels_Clear(s.instance)
}

// Delete
func (s *TStatusPanels) Delete(Index int32) {
    StatusPanels_Delete(s.instance, Index)
}

// EndUpdate
func (s *TStatusPanels) EndUpdate() {
    StatusPanels_EndUpdate(s.instance)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TStatusPanels) GetNamePath() string {
    return StatusPanels_GetNamePath(s.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TStatusPanels) DisposeOf() {
    StatusPanels_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TStatusPanels) ClassType() TClass {
    return StatusPanels_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TStatusPanels) ClassName() string {
    return StatusPanels_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TStatusPanels) InstanceSize() int32 {
    return StatusPanels_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TStatusPanels) InheritsFrom(AClass TClass) bool {
    return StatusPanels_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TStatusPanels) Equals(Obj IObject) bool {
    return StatusPanels_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TStatusPanels) GetHashCode() int32 {
    return StatusPanels_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TStatusPanels) ToString() string {
    return StatusPanels_ToString(s.instance)
}

// Capacity
func (s *TStatusPanels) Capacity() int32 {
    return StatusPanels_GetCapacity(s.instance)
}

// SetCapacity
func (s *TStatusPanels) SetCapacity(value int32) {
    StatusPanels_SetCapacity(s.instance, value)
}

// Count
func (s *TStatusPanels) Count() int32 {
    return StatusPanels_GetCount(s.instance)
}

// Items
func (s *TStatusPanels) Items(Index int32) *TStatusPanel {
    return StatusPanelFromInst(StatusPanels_GetItems(s.instance, Index))
}

// Items
func (s *TStatusPanels) SetItems(Index int32, value *TStatusPanel) {
    StatusPanels_SetItems(s.instance, Index, CheckPtr(value))
}

