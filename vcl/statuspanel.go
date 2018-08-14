
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

type TStatusPanel struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewStatusPanel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewStatusPanel() *TStatusPanel {
    s := new(TStatusPanel)
    s.instance = StatusPanel_Create()
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StatusPanelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func StatusPanelFromInst(inst uintptr) *TStatusPanel {
    s := new(TStatusPanel)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// StatusPanelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func StatusPanelFromObj(obj IObject) *TStatusPanel {
    s := new(TStatusPanel)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StatusPanelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func StatusPanelFromUnsafePointer(ptr unsafe.Pointer) *TStatusPanel {
    s := new(TStatusPanel)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TStatusPanel) Free() {
    if s.instance != 0 {
        StatusPanel_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TStatusPanel) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TStatusPanel) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TStatusPanel) IsValid() bool {
    return s.instance != 0
}

// TStatusPanelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TStatusPanelClass() TClass {
    return StatusPanel_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TStatusPanel) Assign(Source IObject) {
    StatusPanel_Assign(s.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TStatusPanel) GetNamePath() string {
    return StatusPanel_GetNamePath(s.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TStatusPanel) DisposeOf() {
    StatusPanel_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TStatusPanel) ClassType() TClass {
    return StatusPanel_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TStatusPanel) ClassName() string {
    return StatusPanel_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TStatusPanel) InstanceSize() int32 {
    return StatusPanel_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TStatusPanel) InheritsFrom(AClass TClass) bool {
    return StatusPanel_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TStatusPanel) Equals(Obj IObject) bool {
    return StatusPanel_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TStatusPanel) GetHashCode() int32 {
    return StatusPanel_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TStatusPanel) ToString() string {
    return StatusPanel_ToString(s.instance)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (s *TStatusPanel) Alignment() TAlignment {
    return StatusPanel_GetAlignment(s.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (s *TStatusPanel) SetAlignment(value TAlignment) {
    StatusPanel_SetAlignment(s.instance, value)
}

// BiDiMode
func (s *TStatusPanel) BiDiMode() TBiDiMode {
    return StatusPanel_GetBiDiMode(s.instance)
}

// SetBiDiMode
func (s *TStatusPanel) SetBiDiMode(value TBiDiMode) {
    StatusPanel_SetBiDiMode(s.instance, value)
}

// Style
func (s *TStatusPanel) Style() TStatusPanelStyle {
    return StatusPanel_GetStyle(s.instance)
}

// SetStyle
func (s *TStatusPanel) SetStyle(value TStatusPanelStyle) {
    StatusPanel_SetStyle(s.instance, value)
}

// Text
// CN: 获取文本。
// EN: .
func (s *TStatusPanel) Text() string {
    return StatusPanel_GetText(s.instance)
}

// SetText
// CN: 设置文本。
// EN: .
func (s *TStatusPanel) SetText(value string) {
    StatusPanel_SetText(s.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (s *TStatusPanel) Width() int32 {
    return StatusPanel_GetWidth(s.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (s *TStatusPanel) SetWidth(value int32) {
    StatusPanel_SetWidth(s.instance, value)
}

// Index
func (s *TStatusPanel) Index() int32 {
    return StatusPanel_GetIndex(s.instance)
}

// SetIndex
func (s *TStatusPanel) SetIndex(value int32) {
    StatusPanel_SetIndex(s.instance, value)
}

