
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

type TOpenDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewOpenDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewOpenDialog(owner IComponent) *TOpenDialog {
    o := new(TOpenDialog)
    o.instance = OpenDialog_Create(CheckPtr(owner))
    o.ptr = unsafe.Pointer(o.instance)
    return o
}

// OpenDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func OpenDialogFromInst(inst uintptr) *TOpenDialog {
    o := new(TOpenDialog)
    o.instance = inst
    o.ptr = unsafe.Pointer(inst)
    return o
}

// OpenDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func OpenDialogFromObj(obj IObject) *TOpenDialog {
    o := new(TOpenDialog)
    o.instance = CheckPtr(obj)
    o.ptr = unsafe.Pointer(o.instance)
    return o
}

// OpenDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func OpenDialogFromUnsafePointer(ptr unsafe.Pointer) *TOpenDialog {
    o := new(TOpenDialog)
    o.instance = uintptr(ptr)
    o.ptr = ptr
    return o
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (o *TOpenDialog) Free() {
    if o.instance != 0 {
        OpenDialog_Free(o.instance)
        o.instance = 0
        o.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (o *TOpenDialog) Instance() uintptr {
    return o.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (o *TOpenDialog) UnsafeAddr() unsafe.Pointer {
    return o.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (o *TOpenDialog) IsValid() bool {
    return o.instance != 0
}

// TOpenDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TOpenDialogClass() TClass {
    return OpenDialog_StaticClassType()
}

// Execute
func (o *TOpenDialog) Execute() bool {
    return OpenDialog_Execute(o.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (o *TOpenDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(OpenDialog_FindComponent(o.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (o *TOpenDialog) GetNamePath() string {
    return OpenDialog_GetNamePath(o.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (o *TOpenDialog) HasParent() bool {
    return OpenDialog_HasParent(o.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (o *TOpenDialog) Assign(Source IObject) {
    OpenDialog_Assign(o.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (o *TOpenDialog) DisposeOf() {
    OpenDialog_DisposeOf(o.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (o *TOpenDialog) ClassType() TClass {
    return OpenDialog_ClassType(o.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (o *TOpenDialog) ClassName() string {
    return OpenDialog_ClassName(o.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (o *TOpenDialog) InstanceSize() int32 {
    return OpenDialog_InstanceSize(o.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (o *TOpenDialog) InheritsFrom(AClass TClass) bool {
    return OpenDialog_InheritsFrom(o.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (o *TOpenDialog) Equals(Obj IObject) bool {
    return OpenDialog_Equals(o.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (o *TOpenDialog) GetHashCode() int32 {
    return OpenDialog_GetHashCode(o.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (o *TOpenDialog) ToString() string {
    return OpenDialog_ToString(o.instance)
}

// Files
func (o *TOpenDialog) Files() *TStrings {
    return StringsFromInst(OpenDialog_GetFiles(o.instance))
}

// DefaultExt
func (o *TOpenDialog) DefaultExt() string {
    return OpenDialog_GetDefaultExt(o.instance)
}

// SetDefaultExt
func (o *TOpenDialog) SetDefaultExt(value string) {
    OpenDialog_SetDefaultExt(o.instance, value)
}

// FileName
func (o *TOpenDialog) FileName() string {
    return OpenDialog_GetFileName(o.instance)
}

// SetFileName
func (o *TOpenDialog) SetFileName(value string) {
    OpenDialog_SetFileName(o.instance, value)
}

// Filter
func (o *TOpenDialog) Filter() string {
    return OpenDialog_GetFilter(o.instance)
}

// SetFilter
func (o *TOpenDialog) SetFilter(value string) {
    OpenDialog_SetFilter(o.instance, value)
}

// FilterIndex
func (o *TOpenDialog) FilterIndex() int32 {
    return OpenDialog_GetFilterIndex(o.instance)
}

// SetFilterIndex
func (o *TOpenDialog) SetFilterIndex(value int32) {
    OpenDialog_SetFilterIndex(o.instance, value)
}

// InitialDir
func (o *TOpenDialog) InitialDir() string {
    return OpenDialog_GetInitialDir(o.instance)
}

// SetInitialDir
func (o *TOpenDialog) SetInitialDir(value string) {
    OpenDialog_SetInitialDir(o.instance, value)
}

// Options
func (o *TOpenDialog) Options() TOpenOptions {
    return OpenDialog_GetOptions(o.instance)
}

// SetOptions
func (o *TOpenDialog) SetOptions(value TOpenOptions) {
    OpenDialog_SetOptions(o.instance, value)
}

// OptionsEx
func (o *TOpenDialog) OptionsEx() TOpenOptionsEx {
    return OpenDialog_GetOptionsEx(o.instance)
}

// SetOptionsEx
func (o *TOpenDialog) SetOptionsEx(value TOpenOptionsEx) {
    OpenDialog_SetOptionsEx(o.instance, value)
}

// Title
func (o *TOpenDialog) Title() string {
    return OpenDialog_GetTitle(o.instance)
}

// SetTitle
func (o *TOpenDialog) SetTitle(value string) {
    OpenDialog_SetTitle(o.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (o *TOpenDialog) Handle() HWND {
    return OpenDialog_GetHandle(o.instance)
}

// SetOnClose
func (o *TOpenDialog) SetOnClose(fn TNotifyEvent) {
    OpenDialog_SetOnClose(o.instance, fn)
}

// SetOnShow
func (o *TOpenDialog) SetOnShow(fn TNotifyEvent) {
    OpenDialog_SetOnShow(o.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (o *TOpenDialog) ComponentCount() int32 {
    return OpenDialog_GetComponentCount(o.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (o *TOpenDialog) ComponentIndex() int32 {
    return OpenDialog_GetComponentIndex(o.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (o *TOpenDialog) SetComponentIndex(value int32) {
    OpenDialog_SetComponentIndex(o.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (o *TOpenDialog) Owner() *TComponent {
    return ComponentFromInst(OpenDialog_GetOwner(o.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (o *TOpenDialog) Name() string {
    return OpenDialog_GetName(o.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (o *TOpenDialog) SetName(value string) {
    OpenDialog_SetName(o.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (o *TOpenDialog) Tag() int {
    return OpenDialog_GetTag(o.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (o *TOpenDialog) SetTag(value int) {
    OpenDialog_SetTag(o.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (o *TOpenDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(OpenDialog_GetComponents(o.instance, AIndex))
}

