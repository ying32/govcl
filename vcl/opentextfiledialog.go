
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

type TOpenTextFileDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewOpenTextFileDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewOpenTextFileDialog(owner IComponent) *TOpenTextFileDialog {
    o := new(TOpenTextFileDialog)
    o.instance = OpenTextFileDialog_Create(CheckPtr(owner))
    o.ptr = unsafe.Pointer(o.instance)
    return o
}

// OpenTextFileDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func OpenTextFileDialogFromInst(inst uintptr) *TOpenTextFileDialog {
    o := new(TOpenTextFileDialog)
    o.instance = inst
    o.ptr = unsafe.Pointer(inst)
    return o
}

// OpenTextFileDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func OpenTextFileDialogFromObj(obj IObject) *TOpenTextFileDialog {
    o := new(TOpenTextFileDialog)
    o.instance = CheckPtr(obj)
    o.ptr = unsafe.Pointer(o.instance)
    return o
}

// OpenTextFileDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func OpenTextFileDialogFromUnsafePointer(ptr unsafe.Pointer) *TOpenTextFileDialog {
    o := new(TOpenTextFileDialog)
    o.instance = uintptr(ptr)
    o.ptr = ptr
    return o
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (o *TOpenTextFileDialog) Free() {
    if o.instance != 0 {
        OpenTextFileDialog_Free(o.instance)
        o.instance = 0
        o.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (o *TOpenTextFileDialog) Instance() uintptr {
    return o.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (o *TOpenTextFileDialog) UnsafeAddr() unsafe.Pointer {
    return o.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (o *TOpenTextFileDialog) IsValid() bool {
    return o.instance != 0
}

// TOpenTextFileDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TOpenTextFileDialogClass() TClass {
    return OpenTextFileDialog_StaticClassType()
}

// Execute
// CN: 执行。
// EN: .
func (o *TOpenTextFileDialog) Execute() bool {
    return OpenTextFileDialog_Execute(o.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (o *TOpenTextFileDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(OpenTextFileDialog_FindComponent(o.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (o *TOpenTextFileDialog) GetNamePath() string {
    return OpenTextFileDialog_GetNamePath(o.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (o *TOpenTextFileDialog) HasParent() bool {
    return OpenTextFileDialog_HasParent(o.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (o *TOpenTextFileDialog) Assign(Source IObject) {
    OpenTextFileDialog_Assign(o.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (o *TOpenTextFileDialog) DisposeOf() {
    OpenTextFileDialog_DisposeOf(o.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (o *TOpenTextFileDialog) ClassType() TClass {
    return OpenTextFileDialog_ClassType(o.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (o *TOpenTextFileDialog) ClassName() string {
    return OpenTextFileDialog_ClassName(o.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (o *TOpenTextFileDialog) InstanceSize() int32 {
    return OpenTextFileDialog_InstanceSize(o.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (o *TOpenTextFileDialog) InheritsFrom(AClass TClass) bool {
    return OpenTextFileDialog_InheritsFrom(o.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (o *TOpenTextFileDialog) Equals(Obj IObject) bool {
    return OpenTextFileDialog_Equals(o.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (o *TOpenTextFileDialog) GetHashCode() int32 {
    return OpenTextFileDialog_GetHashCode(o.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (o *TOpenTextFileDialog) ToString() string {
    return OpenTextFileDialog_ToString(o.instance)
}

// Files
func (o *TOpenTextFileDialog) Files() *TStrings {
    return StringsFromInst(OpenTextFileDialog_GetFiles(o.instance))
}

// DefaultExt
func (o *TOpenTextFileDialog) DefaultExt() string {
    return OpenTextFileDialog_GetDefaultExt(o.instance)
}

// SetDefaultExt
func (o *TOpenTextFileDialog) SetDefaultExt(value string) {
    OpenTextFileDialog_SetDefaultExt(o.instance, value)
}

// FileName
func (o *TOpenTextFileDialog) FileName() string {
    return OpenTextFileDialog_GetFileName(o.instance)
}

// SetFileName
func (o *TOpenTextFileDialog) SetFileName(value string) {
    OpenTextFileDialog_SetFileName(o.instance, value)
}

// Filter
func (o *TOpenTextFileDialog) Filter() string {
    return OpenTextFileDialog_GetFilter(o.instance)
}

// SetFilter
func (o *TOpenTextFileDialog) SetFilter(value string) {
    OpenTextFileDialog_SetFilter(o.instance, value)
}

// FilterIndex
func (o *TOpenTextFileDialog) FilterIndex() int32 {
    return OpenTextFileDialog_GetFilterIndex(o.instance)
}

// SetFilterIndex
func (o *TOpenTextFileDialog) SetFilterIndex(value int32) {
    OpenTextFileDialog_SetFilterIndex(o.instance, value)
}

// InitialDir
func (o *TOpenTextFileDialog) InitialDir() string {
    return OpenTextFileDialog_GetInitialDir(o.instance)
}

// SetInitialDir
func (o *TOpenTextFileDialog) SetInitialDir(value string) {
    OpenTextFileDialog_SetInitialDir(o.instance, value)
}

// Options
func (o *TOpenTextFileDialog) Options() TOpenOptions {
    return OpenTextFileDialog_GetOptions(o.instance)
}

// SetOptions
func (o *TOpenTextFileDialog) SetOptions(value TOpenOptions) {
    OpenTextFileDialog_SetOptions(o.instance, value)
}

// OptionsEx
func (o *TOpenTextFileDialog) OptionsEx() TOpenOptionsEx {
    return OpenTextFileDialog_GetOptionsEx(o.instance)
}

// SetOptionsEx
func (o *TOpenTextFileDialog) SetOptionsEx(value TOpenOptionsEx) {
    OpenTextFileDialog_SetOptionsEx(o.instance, value)
}

// Title
func (o *TOpenTextFileDialog) Title() string {
    return OpenTextFileDialog_GetTitle(o.instance)
}

// SetTitle
func (o *TOpenTextFileDialog) SetTitle(value string) {
    OpenTextFileDialog_SetTitle(o.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (o *TOpenTextFileDialog) Handle() HWND {
    return OpenTextFileDialog_GetHandle(o.instance)
}

// SetOnClose
func (o *TOpenTextFileDialog) SetOnClose(fn TNotifyEvent) {
    OpenTextFileDialog_SetOnClose(o.instance, fn)
}

// SetOnShow
// CN: 设置显示事件。
// EN: .
func (o *TOpenTextFileDialog) SetOnShow(fn TNotifyEvent) {
    OpenTextFileDialog_SetOnShow(o.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (o *TOpenTextFileDialog) ComponentCount() int32 {
    return OpenTextFileDialog_GetComponentCount(o.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (o *TOpenTextFileDialog) ComponentIndex() int32 {
    return OpenTextFileDialog_GetComponentIndex(o.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (o *TOpenTextFileDialog) SetComponentIndex(value int32) {
    OpenTextFileDialog_SetComponentIndex(o.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (o *TOpenTextFileDialog) Owner() *TComponent {
    return ComponentFromInst(OpenTextFileDialog_GetOwner(o.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (o *TOpenTextFileDialog) Name() string {
    return OpenTextFileDialog_GetName(o.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (o *TOpenTextFileDialog) SetName(value string) {
    OpenTextFileDialog_SetName(o.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (o *TOpenTextFileDialog) Tag() int {
    return OpenTextFileDialog_GetTag(o.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (o *TOpenTextFileDialog) SetTag(value int) {
    OpenTextFileDialog_SetTag(o.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (o *TOpenTextFileDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(OpenTextFileDialog_GetComponents(o.instance, AIndex))
}

