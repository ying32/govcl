
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

type TPageSetupDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPageSetupDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPageSetupDialog(owner IComponent) *TPageSetupDialog {
    p := new(TPageSetupDialog)
    p.instance = PageSetupDialog_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PageSetupDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PageSetupDialogFromInst(inst uintptr) *TPageSetupDialog {
    p := new(TPageSetupDialog)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PageSetupDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PageSetupDialogFromObj(obj IObject) *TPageSetupDialog {
    p := new(TPageSetupDialog)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PageSetupDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PageSetupDialogFromUnsafePointer(ptr unsafe.Pointer) *TPageSetupDialog {
    p := new(TPageSetupDialog)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPageSetupDialog) Free() {
    if p.instance != 0 {
        PageSetupDialog_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPageSetupDialog) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPageSetupDialog) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPageSetupDialog) IsValid() bool {
    return p.instance != 0
}

// TPageSetupDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPageSetupDialogClass() TClass {
    return PageSetupDialog_StaticClassType()
}

// GetDefaults
func (p *TPageSetupDialog) GetDefaults() bool {
    return PageSetupDialog_GetDefaults(p.instance)
}

// Execute
func (p *TPageSetupDialog) Execute() bool {
    return PageSetupDialog_Execute(p.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (p *TPageSetupDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PageSetupDialog_FindComponent(p.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPageSetupDialog) GetNamePath() string {
    return PageSetupDialog_GetNamePath(p.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (p *TPageSetupDialog) HasParent() bool {
    return PageSetupDialog_HasParent(p.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPageSetupDialog) Assign(Source IObject) {
    PageSetupDialog_Assign(p.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPageSetupDialog) DisposeOf() {
    PageSetupDialog_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPageSetupDialog) ClassType() TClass {
    return PageSetupDialog_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPageSetupDialog) ClassName() string {
    return PageSetupDialog_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPageSetupDialog) InstanceSize() int32 {
    return PageSetupDialog_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPageSetupDialog) InheritsFrom(AClass TClass) bool {
    return PageSetupDialog_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPageSetupDialog) Equals(Obj IObject) bool {
    return PageSetupDialog_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPageSetupDialog) GetHashCode() int32 {
    return PageSetupDialog_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPageSetupDialog) ToString() string {
    return PageSetupDialog_ToString(p.instance)
}

// MinMarginLeft
func (p *TPageSetupDialog) MinMarginLeft() int32 {
    return PageSetupDialog_GetMinMarginLeft(p.instance)
}

// SetMinMarginLeft
func (p *TPageSetupDialog) SetMinMarginLeft(value int32) {
    PageSetupDialog_SetMinMarginLeft(p.instance, value)
}

// MinMarginTop
func (p *TPageSetupDialog) MinMarginTop() int32 {
    return PageSetupDialog_GetMinMarginTop(p.instance)
}

// SetMinMarginTop
func (p *TPageSetupDialog) SetMinMarginTop(value int32) {
    PageSetupDialog_SetMinMarginTop(p.instance, value)
}

// MinMarginRight
func (p *TPageSetupDialog) MinMarginRight() int32 {
    return PageSetupDialog_GetMinMarginRight(p.instance)
}

// SetMinMarginRight
func (p *TPageSetupDialog) SetMinMarginRight(value int32) {
    PageSetupDialog_SetMinMarginRight(p.instance, value)
}

// MinMarginBottom
func (p *TPageSetupDialog) MinMarginBottom() int32 {
    return PageSetupDialog_GetMinMarginBottom(p.instance)
}

// SetMinMarginBottom
func (p *TPageSetupDialog) SetMinMarginBottom(value int32) {
    PageSetupDialog_SetMinMarginBottom(p.instance, value)
}

// MarginLeft
func (p *TPageSetupDialog) MarginLeft() int32 {
    return PageSetupDialog_GetMarginLeft(p.instance)
}

// SetMarginLeft
func (p *TPageSetupDialog) SetMarginLeft(value int32) {
    PageSetupDialog_SetMarginLeft(p.instance, value)
}

// MarginTop
func (p *TPageSetupDialog) MarginTop() int32 {
    return PageSetupDialog_GetMarginTop(p.instance)
}

// SetMarginTop
func (p *TPageSetupDialog) SetMarginTop(value int32) {
    PageSetupDialog_SetMarginTop(p.instance, value)
}

// MarginRight
func (p *TPageSetupDialog) MarginRight() int32 {
    return PageSetupDialog_GetMarginRight(p.instance)
}

// SetMarginRight
func (p *TPageSetupDialog) SetMarginRight(value int32) {
    PageSetupDialog_SetMarginRight(p.instance, value)
}

// MarginBottom
func (p *TPageSetupDialog) MarginBottom() int32 {
    return PageSetupDialog_GetMarginBottom(p.instance)
}

// SetMarginBottom
func (p *TPageSetupDialog) SetMarginBottom(value int32) {
    PageSetupDialog_SetMarginBottom(p.instance, value)
}

// Options
func (p *TPageSetupDialog) Options() TPageSetupDialogOptions {
    return PageSetupDialog_GetOptions(p.instance)
}

// SetOptions
func (p *TPageSetupDialog) SetOptions(value TPageSetupDialogOptions) {
    PageSetupDialog_SetOptions(p.instance, value)
}

// PageWidth
func (p *TPageSetupDialog) PageWidth() int32 {
    return PageSetupDialog_GetPageWidth(p.instance)
}

// SetPageWidth
func (p *TPageSetupDialog) SetPageWidth(value int32) {
    PageSetupDialog_SetPageWidth(p.instance, value)
}

// PageHeight
func (p *TPageSetupDialog) PageHeight() int32 {
    return PageSetupDialog_GetPageHeight(p.instance)
}

// SetPageHeight
func (p *TPageSetupDialog) SetPageHeight(value int32) {
    PageSetupDialog_SetPageHeight(p.instance, value)
}

// Units
func (p *TPageSetupDialog) Units() TPageMeasureUnits {
    return PageSetupDialog_GetUnits(p.instance)
}

// SetUnits
func (p *TPageSetupDialog) SetUnits(value TPageMeasureUnits) {
    PageSetupDialog_SetUnits(p.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (p *TPageSetupDialog) Handle() HWND {
    return PageSetupDialog_GetHandle(p.instance)
}

// SetOnClose
func (p *TPageSetupDialog) SetOnClose(fn TNotifyEvent) {
    PageSetupDialog_SetOnClose(p.instance, fn)
}

// SetOnShow
func (p *TPageSetupDialog) SetOnShow(fn TNotifyEvent) {
    PageSetupDialog_SetOnShow(p.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPageSetupDialog) ComponentCount() int32 {
    return PageSetupDialog_GetComponentCount(p.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (p *TPageSetupDialog) ComponentIndex() int32 {
    return PageSetupDialog_GetComponentIndex(p.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (p *TPageSetupDialog) SetComponentIndex(value int32) {
    PageSetupDialog_SetComponentIndex(p.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPageSetupDialog) Owner() *TComponent {
    return ComponentFromInst(PageSetupDialog_GetOwner(p.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPageSetupDialog) Name() string {
    return PageSetupDialog_GetName(p.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPageSetupDialog) SetName(value string) {
    PageSetupDialog_SetName(p.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPageSetupDialog) Tag() int {
    return PageSetupDialog_GetTag(p.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPageSetupDialog) SetTag(value int) {
    PageSetupDialog_SetTag(p.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPageSetupDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PageSetupDialog_GetComponents(p.instance, AIndex))
}

