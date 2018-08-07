
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

type TBrush struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewBrush
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBrush() *TBrush {
    b := new(TBrush)
    b.instance = Brush_Create()
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BrushFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func BrushFromInst(inst uintptr) *TBrush {
    b := new(TBrush)
    b.instance = inst
    b.ptr = unsafe.Pointer(inst)
    return b
}

// BrushFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func BrushFromObj(obj IObject) *TBrush {
    b := new(TBrush)
    b.instance = CheckPtr(obj)
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BrushFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func BrushFromUnsafePointer(ptr unsafe.Pointer) *TBrush {
    b := new(TBrush)
    b.instance = uintptr(ptr)
    b.ptr = ptr
    return b
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (b *TBrush) Free() {
    if b.instance != 0 {
        Brush_Free(b.instance)
        b.instance = 0
        b.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBrush) Instance() uintptr {
    return b.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBrush) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBrush) IsValid() bool {
    return b.instance != 0
}

// TBrushClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBrushClass() TClass {
    return Brush_StaticClassType()
}

// Assign
func (b *TBrush) Assign(Source IObject) {
    Brush_Assign(b.instance, CheckPtr(Source))
}

// HandleAllocated
func (b *TBrush) HandleAllocated() bool {
    return Brush_HandleAllocated(b.instance)
}

// GetNamePath
func (b *TBrush) GetNamePath() string {
    return Brush_GetNamePath(b.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TBrush) DisposeOf() {
    Brush_DisposeOf(b.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBrush) ClassType() TClass {
    return Brush_ClassType(b.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBrush) ClassName() string {
    return Brush_ClassName(b.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBrush) InstanceSize() int32 {
    return Brush_InstanceSize(b.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBrush) InheritsFrom(AClass TClass) bool {
    return Brush_InheritsFrom(b.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBrush) Equals(Obj IObject) bool {
    return Brush_Equals(b.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBrush) GetHashCode() int32 {
    return Brush_GetHashCode(b.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (b *TBrush) ToString() string {
    return Brush_ToString(b.instance)
}

// Bitmap
func (b *TBrush) Bitmap() *TBitmap {
    return BitmapFromInst(Brush_GetBitmap(b.instance))
}

// SetBitmap
func (b *TBrush) SetBitmap(value *TBitmap) {
    Brush_SetBitmap(b.instance, CheckPtr(value))
}

// Handle
func (b *TBrush) Handle() HBRUSH {
    return Brush_GetHandle(b.instance)
}

// SetHandle
func (b *TBrush) SetHandle(value HBRUSH) {
    Brush_SetHandle(b.instance, value)
}

// Color
func (b *TBrush) Color() TColor {
    return Brush_GetColor(b.instance)
}

// SetColor
func (b *TBrush) SetColor(value TColor) {
    Brush_SetColor(b.instance, value)
}

// Style
func (b *TBrush) Style() TBrushStyle {
    return Brush_GetStyle(b.instance)
}

// SetStyle
func (b *TBrush) SetStyle(value TBrushStyle) {
    Brush_SetStyle(b.instance, value)
}

// SetOnChange
func (b *TBrush) SetOnChange(fn TNotifyEvent) {
    Brush_SetOnChange(b.instance, fn)
}

