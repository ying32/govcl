
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

type TControlScrollBar struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// ControlScrollBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ControlScrollBarFromInst(inst uintptr) *TControlScrollBar {
    c := new(TControlScrollBar)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ControlScrollBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ControlScrollBarFromObj(obj IObject) *TControlScrollBar {
    c := new(TControlScrollBar)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ControlScrollBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ControlScrollBarFromUnsafePointer(ptr unsafe.Pointer) *TControlScrollBar {
    c := new(TControlScrollBar)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TControlScrollBar) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TControlScrollBar) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TControlScrollBar) IsValid() bool {
    return c.instance != 0
}

// TControlScrollBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TControlScrollBarClass() TClass {
    return ControlScrollBar_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TControlScrollBar) Assign(Source IObject) {
    ControlScrollBar_Assign(c.instance, CheckPtr(Source))
}

// ChangeBiDiPosition
func (c *TControlScrollBar) ChangeBiDiPosition() {
    ControlScrollBar_ChangeBiDiPosition(c.instance)
}

// IsScrollBarVisible
func (c *TControlScrollBar) IsScrollBarVisible() bool {
    return ControlScrollBar_IsScrollBarVisible(c.instance)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TControlScrollBar) GetNamePath() string {
    return ControlScrollBar_GetNamePath(c.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TControlScrollBar) DisposeOf() {
    ControlScrollBar_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TControlScrollBar) ClassType() TClass {
    return ControlScrollBar_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TControlScrollBar) ClassName() string {
    return ControlScrollBar_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TControlScrollBar) InstanceSize() int32 {
    return ControlScrollBar_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TControlScrollBar) InheritsFrom(AClass TClass) bool {
    return ControlScrollBar_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TControlScrollBar) Equals(Obj IObject) bool {
    return ControlScrollBar_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TControlScrollBar) GetHashCode() int32 {
    return ControlScrollBar_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TControlScrollBar) ToString() string {
    return ControlScrollBar_ToString(c.instance)
}

// Kind
func (c *TControlScrollBar) Kind() TScrollBarKind {
    return ControlScrollBar_GetKind(c.instance)
}

// ScrollPos
func (c *TControlScrollBar) ScrollPos() int32 {
    return ControlScrollBar_GetScrollPos(c.instance)
}

// ButtonSize
func (c *TControlScrollBar) ButtonSize() int32 {
    return ControlScrollBar_GetButtonSize(c.instance)
}

// SetButtonSize
func (c *TControlScrollBar) SetButtonSize(value int32) {
    ControlScrollBar_SetButtonSize(c.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (c *TControlScrollBar) Color() TColor {
    return ControlScrollBar_GetColor(c.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (c *TControlScrollBar) SetColor(value TColor) {
    ControlScrollBar_SetColor(c.instance, value)
}

// Increment
func (c *TControlScrollBar) Increment() TScrollBarInc {
    return ControlScrollBar_GetIncrement(c.instance)
}

// SetIncrement
func (c *TControlScrollBar) SetIncrement(value TScrollBarInc) {
    ControlScrollBar_SetIncrement(c.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (c *TControlScrollBar) ParentColor() bool {
    return ControlScrollBar_GetParentColor(c.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (c *TControlScrollBar) SetParentColor(value bool) {
    ControlScrollBar_SetParentColor(c.instance, value)
}

// Position
func (c *TControlScrollBar) Position() int32 {
    return ControlScrollBar_GetPosition(c.instance)
}

// SetPosition
func (c *TControlScrollBar) SetPosition(value int32) {
    ControlScrollBar_SetPosition(c.instance, value)
}

// Range
func (c *TControlScrollBar) Range() int32 {
    return ControlScrollBar_GetRange(c.instance)
}

// SetRange
func (c *TControlScrollBar) SetRange(value int32) {
    ControlScrollBar_SetRange(c.instance, value)
}

// Smooth
func (c *TControlScrollBar) Smooth() bool {
    return ControlScrollBar_GetSmooth(c.instance)
}

// SetSmooth
func (c *TControlScrollBar) SetSmooth(value bool) {
    ControlScrollBar_SetSmooth(c.instance, value)
}

// Size
func (c *TControlScrollBar) Size() int32 {
    return ControlScrollBar_GetSize(c.instance)
}

// SetSize
func (c *TControlScrollBar) SetSize(value int32) {
    ControlScrollBar_SetSize(c.instance, value)
}

// Style
func (c *TControlScrollBar) Style() TScrollBarStyle {
    return ControlScrollBar_GetStyle(c.instance)
}

// SetStyle
func (c *TControlScrollBar) SetStyle(value TScrollBarStyle) {
    ControlScrollBar_SetStyle(c.instance, value)
}

// ThumbSize
func (c *TControlScrollBar) ThumbSize() int32 {
    return ControlScrollBar_GetThumbSize(c.instance)
}

// SetThumbSize
func (c *TControlScrollBar) SetThumbSize(value int32) {
    ControlScrollBar_SetThumbSize(c.instance, value)
}

// Tracking
func (c *TControlScrollBar) Tracking() bool {
    return ControlScrollBar_GetTracking(c.instance)
}

// SetTracking
func (c *TControlScrollBar) SetTracking(value bool) {
    ControlScrollBar_SetTracking(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TControlScrollBar) Visible() bool {
    return ControlScrollBar_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TControlScrollBar) SetVisible(value bool) {
    ControlScrollBar_SetVisible(c.instance, value)
}

