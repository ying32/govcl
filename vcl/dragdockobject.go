
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

type TDragDockObject struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewDragDockObject
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewDragDockObject() *TDragDockObject {
    d := new(TDragDockObject)
    d.instance = DragDockObject_Create()
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// DragDockObjectFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func DragDockObjectFromInst(inst uintptr) *TDragDockObject {
    d := new(TDragDockObject)
    d.instance = inst
    d.ptr = unsafe.Pointer(inst)
    return d
}

// DragDockObjectFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func DragDockObjectFromObj(obj IObject) *TDragDockObject {
    d := new(TDragDockObject)
    d.instance = CheckPtr(obj)
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// DragDockObjectFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func DragDockObjectFromUnsafePointer(ptr unsafe.Pointer) *TDragDockObject {
    d := new(TDragDockObject)
    d.instance = uintptr(ptr)
    d.ptr = ptr
    return d
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (d *TDragDockObject) Free() {
    if d.instance != 0 {
        DragDockObject_Free(d.instance)
        d.instance = 0
        d.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (d *TDragDockObject) Instance() uintptr {
    return d.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (d *TDragDockObject) UnsafeAddr() unsafe.Pointer {
    return d.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (d *TDragDockObject) IsValid() bool {
    return d.instance != 0
}

// TDragDockObjectClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TDragDockObjectClass() TClass {
    return DragDockObject_StaticClassType()
}

// Assign
func (d *TDragDockObject) Assign(Source *TDragObject) {
    DragDockObject_Assign(d.instance, CheckPtr(Source))
}

// HideDragImage
func (d *TDragDockObject) HideDragImage() {
    DragDockObject_HideDragImage(d.instance)
}

// ShowDragImage
func (d *TDragDockObject) ShowDragImage() {
    DragDockObject_ShowDragImage(d.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (d *TDragDockObject) DisposeOf() {
    DragDockObject_DisposeOf(d.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (d *TDragDockObject) ClassType() TClass {
    return DragDockObject_ClassType(d.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (d *TDragDockObject) ClassName() string {
    return DragDockObject_ClassName(d.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (d *TDragDockObject) InstanceSize() int32 {
    return DragDockObject_InstanceSize(d.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (d *TDragDockObject) InheritsFrom(AClass TClass) bool {
    return DragDockObject_InheritsFrom(d.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (d *TDragDockObject) Equals(Obj IObject) bool {
    return DragDockObject_Equals(d.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (d *TDragDockObject) GetHashCode() int32 {
    return DragDockObject_GetHashCode(d.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (d *TDragDockObject) ToString() string {
    return DragDockObject_ToString(d.instance)
}

// Brush
func (d *TDragDockObject) Brush() *TBrush {
    return BrushFromInst(DragDockObject_GetBrush(d.instance))
}

// SetBrush
func (d *TDragDockObject) SetBrush(value *TBrush) {
    DragDockObject_SetBrush(d.instance, CheckPtr(value))
}

// DockRect
func (d *TDragDockObject) DockRect() TRect {
    return DragDockObject_GetDockRect(d.instance)
}

// SetDockRect
func (d *TDragDockObject) SetDockRect(value TRect) {
    DragDockObject_SetDockRect(d.instance, value)
}

// DropAlign
func (d *TDragDockObject) DropAlign() TAlign {
    return DragDockObject_GetDropAlign(d.instance)
}

// DropOnControl
func (d *TDragDockObject) DropOnControl() *TControl {
    return ControlFromInst(DragDockObject_GetDropOnControl(d.instance))
}

// EraseDockRect
func (d *TDragDockObject) EraseDockRect() TRect {
    return DragDockObject_GetEraseDockRect(d.instance)
}

// SetEraseDockRect
func (d *TDragDockObject) SetEraseDockRect(value TRect) {
    DragDockObject_SetEraseDockRect(d.instance, value)
}

// EraseWhenMoving
func (d *TDragDockObject) EraseWhenMoving() bool {
    return DragDockObject_GetEraseWhenMoving(d.instance)
}

// Floating
func (d *TDragDockObject) Floating() bool {
    return DragDockObject_GetFloating(d.instance)
}

// SetFloating
func (d *TDragDockObject) SetFloating(value bool) {
    DragDockObject_SetFloating(d.instance, value)
}

// FrameWidth
func (d *TDragDockObject) FrameWidth() int32 {
    return DragDockObject_GetFrameWidth(d.instance)
}

// Control
func (d *TDragDockObject) Control() *TControl {
    return ControlFromInst(DragDockObject_GetControl(d.instance))
}

// SetControl
func (d *TDragDockObject) SetControl(value IControl) {
    DragDockObject_SetControl(d.instance, CheckPtr(value))
}

// AlwaysShowDragImages
func (d *TDragDockObject) AlwaysShowDragImages() bool {
    return DragDockObject_GetAlwaysShowDragImages(d.instance)
}

// SetAlwaysShowDragImages
func (d *TDragDockObject) SetAlwaysShowDragImages(value bool) {
    DragDockObject_SetAlwaysShowDragImages(d.instance, value)
}

// Cancelling
func (d *TDragDockObject) Cancelling() bool {
    return DragDockObject_GetCancelling(d.instance)
}

// SetCancelling
func (d *TDragDockObject) SetCancelling(value bool) {
    DragDockObject_SetCancelling(d.instance, value)
}

// DragHandle
func (d *TDragDockObject) DragHandle() HWND {
    return DragDockObject_GetDragHandle(d.instance)
}

// SetDragHandle
func (d *TDragDockObject) SetDragHandle(value HWND) {
    DragDockObject_SetDragHandle(d.instance, value)
}

// DragPos
func (d *TDragDockObject) DragPos() TPoint {
    return DragDockObject_GetDragPos(d.instance)
}

// SetDragPos
func (d *TDragDockObject) SetDragPos(value TPoint) {
    DragDockObject_SetDragPos(d.instance, value)
}

// DragTarget
func (d *TDragDockObject) DragTarget() uintptr {
    return DragDockObject_GetDragTarget(d.instance)
}

// SetDragTarget
func (d *TDragDockObject) SetDragTarget(value uintptr) {
    DragDockObject_SetDragTarget(d.instance, value)
}

// DragTargetPos
func (d *TDragDockObject) DragTargetPos() TPoint {
    return DragDockObject_GetDragTargetPos(d.instance)
}

// SetDragTargetPos
func (d *TDragDockObject) SetDragTargetPos(value TPoint) {
    DragDockObject_SetDragTargetPos(d.instance, value)
}

// Dropped
func (d *TDragDockObject) Dropped() bool {
    return DragDockObject_GetDropped(d.instance)
}

// MouseDeltaX
func (d *TDragDockObject) MouseDeltaX() float64 {
    return DragDockObject_GetMouseDeltaX(d.instance)
}

// MouseDeltaY
func (d *TDragDockObject) MouseDeltaY() float64 {
    return DragDockObject_GetMouseDeltaY(d.instance)
}

// RightClickCancels
func (d *TDragDockObject) RightClickCancels() bool {
    return DragDockObject_GetRightClickCancels(d.instance)
}

// SetRightClickCancels
func (d *TDragDockObject) SetRightClickCancels(value bool) {
    DragDockObject_SetRightClickCancels(d.instance, value)
}

