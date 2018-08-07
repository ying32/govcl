
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

type TCustomHint struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCustomHint
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCustomHint(owner IComponent) *TCustomHint {
    c := new(TCustomHint)
    c.instance = CustomHint_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CustomHintFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CustomHintFromInst(inst uintptr) *TCustomHint {
    c := new(TCustomHint)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CustomHintFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CustomHintFromObj(obj IObject) *TCustomHint {
    c := new(TCustomHint)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CustomHintFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CustomHintFromUnsafePointer(ptr unsafe.Pointer) *TCustomHint {
    c := new(TCustomHint)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCustomHint) Free() {
    if c.instance != 0 {
        CustomHint_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCustomHint) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCustomHint) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCustomHint) IsValid() bool {
    return c.instance != 0
}

// TCustomHintClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCustomHintClass() TClass {
    return CustomHint_StaticClassType()
}

// ShowHint
func (c *TCustomHint) ShowHint() {
    CustomHint_ShowHint(c.instance)
}

// HideHint
func (c *TCustomHint) HideHint() {
    CustomHint_HideHint(c.instance)
}

// FindComponent
func (c *TCustomHint) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CustomHint_FindComponent(c.instance, AName))
}

// GetNamePath
func (c *TCustomHint) GetNamePath() string {
    return CustomHint_GetNamePath(c.instance)
}

// HasParent
func (c *TCustomHint) HasParent() bool {
    return CustomHint_HasParent(c.instance)
}

// Assign
func (c *TCustomHint) Assign(Source IObject) {
    CustomHint_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCustomHint) DisposeOf() {
    CustomHint_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCustomHint) ClassType() TClass {
    return CustomHint_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCustomHint) ClassName() string {
    return CustomHint_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCustomHint) InstanceSize() int32 {
    return CustomHint_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCustomHint) InheritsFrom(AClass TClass) bool {
    return CustomHint_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCustomHint) Equals(Obj IObject) bool {
    return CustomHint_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCustomHint) GetHashCode() int32 {
    return CustomHint_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCustomHint) ToString() string {
    return CustomHint_ToString(c.instance)
}

// ShowingHint
func (c *TCustomHint) ShowingHint() bool {
    return CustomHint_GetShowingHint(c.instance)
}

// Title
func (c *TCustomHint) Title() string {
    return CustomHint_GetTitle(c.instance)
}

// SetTitle
func (c *TCustomHint) SetTitle(value string) {
    CustomHint_SetTitle(c.instance, value)
}

// Description
func (c *TCustomHint) Description() string {
    return CustomHint_GetDescription(c.instance)
}

// SetDescription
func (c *TCustomHint) SetDescription(value string) {
    CustomHint_SetDescription(c.instance, value)
}

// ImageIndex
func (c *TCustomHint) ImageIndex() int32 {
    return CustomHint_GetImageIndex(c.instance)
}

// SetImageIndex
func (c *TCustomHint) SetImageIndex(value int32) {
    CustomHint_SetImageIndex(c.instance, value)
}

// Images
func (c *TCustomHint) Images() *TImageList {
    return ImageListFromInst(CustomHint_GetImages(c.instance))
}

// SetImages
func (c *TCustomHint) SetImages(value IComponent) {
    CustomHint_SetImages(c.instance, CheckPtr(value))
}

// Style
func (c *TCustomHint) Style() TBalloonHintStyle {
    return CustomHint_GetStyle(c.instance)
}

// SetStyle
func (c *TCustomHint) SetStyle(value TBalloonHintStyle) {
    CustomHint_SetStyle(c.instance, value)
}

// Delay
func (c *TCustomHint) Delay() uint32 {
    return CustomHint_GetDelay(c.instance)
}

// SetDelay
func (c *TCustomHint) SetDelay(value uint32) {
    CustomHint_SetDelay(c.instance, value)
}

// HideAfter
func (c *TCustomHint) HideAfter() int32 {
    return CustomHint_GetHideAfter(c.instance)
}

// SetHideAfter
func (c *TCustomHint) SetHideAfter(value int32) {
    CustomHint_SetHideAfter(c.instance, value)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TCustomHint) ComponentCount() int32 {
    return CustomHint_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TCustomHint) ComponentIndex() int32 {
    return CustomHint_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TCustomHint) SetComponentIndex(value int32) {
    CustomHint_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TCustomHint) Owner() *TComponent {
    return ComponentFromInst(CustomHint_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TCustomHint) Name() string {
    return CustomHint_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TCustomHint) SetName(value string) {
    CustomHint_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TCustomHint) Tag() int {
    return CustomHint_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TCustomHint) SetTag(value int) {
    CustomHint_SetTag(c.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TCustomHint) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CustomHint_GetComponents(c.instance, AIndex))
}

