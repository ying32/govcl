
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

type TPen struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPen
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPen() *TPen {
    p := new(TPen)
    p.instance = Pen_Create()
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PenFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PenFromInst(inst uintptr) *TPen {
    p := new(TPen)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PenFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PenFromObj(obj IObject) *TPen {
    p := new(TPen)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PenFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PenFromUnsafePointer(ptr unsafe.Pointer) *TPen {
    p := new(TPen)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPen) Free() {
    if p.instance != 0 {
        Pen_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPen) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPen) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPen) IsValid() bool {
    return p.instance != 0
}

// TPenClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPenClass() TClass {
    return Pen_StaticClassType()
}

// Assign
func (p *TPen) Assign(Source IObject) {
    Pen_Assign(p.instance, CheckPtr(Source))
}

// HandleAllocated
func (p *TPen) HandleAllocated() bool {
    return Pen_HandleAllocated(p.instance)
}

// GetNamePath
func (p *TPen) GetNamePath() string {
    return Pen_GetNamePath(p.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPen) DisposeOf() {
    Pen_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPen) ClassType() TClass {
    return Pen_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPen) ClassName() string {
    return Pen_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPen) InstanceSize() int32 {
    return Pen_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPen) InheritsFrom(AClass TClass) bool {
    return Pen_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPen) Equals(Obj IObject) bool {
    return Pen_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPen) GetHashCode() int32 {
    return Pen_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPen) ToString() string {
    return Pen_ToString(p.instance)
}

// Handle
func (p *TPen) Handle() HPEN {
    return Pen_GetHandle(p.instance)
}

// SetHandle
func (p *TPen) SetHandle(value HPEN) {
    Pen_SetHandle(p.instance, value)
}

// Color
func (p *TPen) Color() TColor {
    return Pen_GetColor(p.instance)
}

// SetColor
func (p *TPen) SetColor(value TColor) {
    Pen_SetColor(p.instance, value)
}

// Mode
func (p *TPen) Mode() TPenMode {
    return Pen_GetMode(p.instance)
}

// SetMode
func (p *TPen) SetMode(value TPenMode) {
    Pen_SetMode(p.instance, value)
}

// Style
func (p *TPen) Style() TPenStyle {
    return Pen_GetStyle(p.instance)
}

// SetStyle
func (p *TPen) SetStyle(value TPenStyle) {
    Pen_SetStyle(p.instance, value)
}

// Width
func (p *TPen) Width() int32 {
    return Pen_GetWidth(p.instance)
}

// SetWidth
func (p *TPen) SetWidth(value int32) {
    Pen_SetWidth(p.instance, value)
}

// SetOnChange
func (p *TPen) SetOnChange(fn TNotifyEvent) {
    Pen_SetOnChange(p.instance, fn)
}

