
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

type TColorDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewColorDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewColorDialog(owner IComponent) *TColorDialog {
    c := new(TColorDialog)
    c.instance = ColorDialog_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ColorDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ColorDialogFromInst(inst uintptr) *TColorDialog {
    c := new(TColorDialog)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ColorDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ColorDialogFromObj(obj IObject) *TColorDialog {
    c := new(TColorDialog)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ColorDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ColorDialogFromUnsafePointer(ptr unsafe.Pointer) *TColorDialog {
    c := new(TColorDialog)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TColorDialog) Free() {
    if c.instance != 0 {
        ColorDialog_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TColorDialog) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TColorDialog) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TColorDialog) IsValid() bool {
    return c.instance != 0
}

// TColorDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TColorDialogClass() TClass {
    return ColorDialog_StaticClassType()
}

// Execute
// CN: 执行。
// EN: .
func (c *TColorDialog) Execute() bool {
    return ColorDialog_Execute(c.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TColorDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ColorDialog_FindComponent(c.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TColorDialog) GetNamePath() string {
    return ColorDialog_GetNamePath(c.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TColorDialog) HasParent() bool {
    return ColorDialog_HasParent(c.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TColorDialog) Assign(Source IObject) {
    ColorDialog_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TColorDialog) DisposeOf() {
    ColorDialog_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TColorDialog) ClassType() TClass {
    return ColorDialog_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TColorDialog) ClassName() string {
    return ColorDialog_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TColorDialog) InstanceSize() int32 {
    return ColorDialog_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TColorDialog) InheritsFrom(AClass TClass) bool {
    return ColorDialog_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TColorDialog) Equals(Obj IObject) bool {
    return ColorDialog_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TColorDialog) GetHashCode() int32 {
    return ColorDialog_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TColorDialog) ToString() string {
    return ColorDialog_ToString(c.instance)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (c *TColorDialog) Color() TColor {
    return ColorDialog_GetColor(c.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (c *TColorDialog) SetColor(value TColor) {
    ColorDialog_SetColor(c.instance, value)
}

// Ctl3D
func (c *TColorDialog) Ctl3D() bool {
    return ColorDialog_GetCtl3D(c.instance)
}

// SetCtl3D
func (c *TColorDialog) SetCtl3D(value bool) {
    ColorDialog_SetCtl3D(c.instance, value)
}

// Options
func (c *TColorDialog) Options() TColorDialogOptions {
    return ColorDialog_GetOptions(c.instance)
}

// SetOptions
func (c *TColorDialog) SetOptions(value TColorDialogOptions) {
    ColorDialog_SetOptions(c.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TColorDialog) Handle() HWND {
    return ColorDialog_GetHandle(c.instance)
}

// SetOnClose
func (c *TColorDialog) SetOnClose(fn TNotifyEvent) {
    ColorDialog_SetOnClose(c.instance, fn)
}

// SetOnShow
// CN: 设置显示事件。
// EN: .
func (c *TColorDialog) SetOnShow(fn TNotifyEvent) {
    ColorDialog_SetOnShow(c.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TColorDialog) ComponentCount() int32 {
    return ColorDialog_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TColorDialog) ComponentIndex() int32 {
    return ColorDialog_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TColorDialog) SetComponentIndex(value int32) {
    ColorDialog_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TColorDialog) Owner() *TComponent {
    return ComponentFromInst(ColorDialog_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TColorDialog) Name() string {
    return ColorDialog_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TColorDialog) SetName(value string) {
    ColorDialog_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TColorDialog) Tag() int {
    return ColorDialog_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TColorDialog) SetTag(value int) {
    ColorDialog_SetTag(c.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TColorDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ColorDialog_GetComponents(c.instance, AIndex))
}

