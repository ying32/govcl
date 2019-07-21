
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

type TSaveDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewSaveDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewSaveDialog(owner IComponent) *TSaveDialog {
    s := new(TSaveDialog)
    s.instance = SaveDialog_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SaveDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func SaveDialogFromInst(inst uintptr) *TSaveDialog {
    s := new(TSaveDialog)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// SaveDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func SaveDialogFromObj(obj IObject) *TSaveDialog {
    s := new(TSaveDialog)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SaveDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func SaveDialogFromUnsafePointer(ptr unsafe.Pointer) *TSaveDialog {
    s := new(TSaveDialog)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TSaveDialog) Free() {
    if s.instance != 0 {
        SaveDialog_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TSaveDialog) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TSaveDialog) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TSaveDialog) IsValid() bool {
    return s.instance != 0
}

// TSaveDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TSaveDialogClass() TClass {
    return SaveDialog_StaticClassType()
}

// Execute
// CN: 执行。
// EN: .
func (s *TSaveDialog) Execute() bool {
    return SaveDialog_Execute(s.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TSaveDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(SaveDialog_FindComponent(s.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TSaveDialog) GetNamePath() string {
    return SaveDialog_GetNamePath(s.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TSaveDialog) HasParent() bool {
    return SaveDialog_HasParent(s.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TSaveDialog) Assign(Source IObject) {
    SaveDialog_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TSaveDialog) DisposeOf() {
    SaveDialog_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TSaveDialog) ClassType() TClass {
    return SaveDialog_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TSaveDialog) ClassName() string {
    return SaveDialog_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TSaveDialog) InstanceSize() int32 {
    return SaveDialog_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TSaveDialog) InheritsFrom(AClass TClass) bool {
    return SaveDialog_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TSaveDialog) Equals(Obj IObject) bool {
    return SaveDialog_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TSaveDialog) GetHashCode() int32 {
    return SaveDialog_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TSaveDialog) ToString() string {
    return SaveDialog_ToString(s.instance)
}

// Files
func (s *TSaveDialog) Files() *TStrings {
    return StringsFromInst(SaveDialog_GetFiles(s.instance))
}

// DefaultExt
func (s *TSaveDialog) DefaultExt() string {
    return SaveDialog_GetDefaultExt(s.instance)
}

// SetDefaultExt
func (s *TSaveDialog) SetDefaultExt(value string) {
    SaveDialog_SetDefaultExt(s.instance, value)
}

// FileName
func (s *TSaveDialog) FileName() string {
    return SaveDialog_GetFileName(s.instance)
}

// SetFileName
func (s *TSaveDialog) SetFileName(value string) {
    SaveDialog_SetFileName(s.instance, value)
}

// Filter
func (s *TSaveDialog) Filter() string {
    return SaveDialog_GetFilter(s.instance)
}

// SetFilter
func (s *TSaveDialog) SetFilter(value string) {
    SaveDialog_SetFilter(s.instance, value)
}

// FilterIndex
func (s *TSaveDialog) FilterIndex() int32 {
    return SaveDialog_GetFilterIndex(s.instance)
}

// SetFilterIndex
func (s *TSaveDialog) SetFilterIndex(value int32) {
    SaveDialog_SetFilterIndex(s.instance, value)
}

// InitialDir
func (s *TSaveDialog) InitialDir() string {
    return SaveDialog_GetInitialDir(s.instance)
}

// SetInitialDir
func (s *TSaveDialog) SetInitialDir(value string) {
    SaveDialog_SetInitialDir(s.instance, value)
}

// Options
func (s *TSaveDialog) Options() TOpenOptions {
    return SaveDialog_GetOptions(s.instance)
}

// SetOptions
func (s *TSaveDialog) SetOptions(value TOpenOptions) {
    SaveDialog_SetOptions(s.instance, value)
}

// OptionsEx
func (s *TSaveDialog) OptionsEx() TOpenOptionsEx {
    return SaveDialog_GetOptionsEx(s.instance)
}

// SetOptionsEx
func (s *TSaveDialog) SetOptionsEx(value TOpenOptionsEx) {
    SaveDialog_SetOptionsEx(s.instance, value)
}

// Title
func (s *TSaveDialog) Title() string {
    return SaveDialog_GetTitle(s.instance)
}

// SetTitle
func (s *TSaveDialog) SetTitle(value string) {
    SaveDialog_SetTitle(s.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (s *TSaveDialog) Handle() HWND {
    return SaveDialog_GetHandle(s.instance)
}

// Ctl3D
func (s *TSaveDialog) Ctl3D() bool {
    return SaveDialog_GetCtl3D(s.instance)
}

// SetCtl3D
func (s *TSaveDialog) SetCtl3D(value bool) {
    SaveDialog_SetCtl3D(s.instance, value)
}

// SetOnClose
func (s *TSaveDialog) SetOnClose(fn TNotifyEvent) {
    SaveDialog_SetOnClose(s.instance, fn)
}

// SetOnShow
// CN: 设置显示事件。
// EN: .
func (s *TSaveDialog) SetOnShow(fn TNotifyEvent) {
    SaveDialog_SetOnShow(s.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TSaveDialog) ComponentCount() int32 {
    return SaveDialog_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TSaveDialog) ComponentIndex() int32 {
    return SaveDialog_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TSaveDialog) SetComponentIndex(value int32) {
    SaveDialog_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TSaveDialog) Owner() *TComponent {
    return ComponentFromInst(SaveDialog_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TSaveDialog) Name() string {
    return SaveDialog_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TSaveDialog) SetName(value string) {
    SaveDialog_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TSaveDialog) Tag() int {
    return SaveDialog_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TSaveDialog) SetTag(value int) {
    SaveDialog_SetTag(s.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TSaveDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(SaveDialog_GetComponents(s.instance, AIndex))
}

