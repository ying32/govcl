
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

type TSavePictureDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewSavePictureDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewSavePictureDialog(owner IComponent) *TSavePictureDialog {
    s := new(TSavePictureDialog)
    s.instance = SavePictureDialog_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SavePictureDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func SavePictureDialogFromInst(inst uintptr) *TSavePictureDialog {
    s := new(TSavePictureDialog)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// SavePictureDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func SavePictureDialogFromObj(obj IObject) *TSavePictureDialog {
    s := new(TSavePictureDialog)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SavePictureDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func SavePictureDialogFromUnsafePointer(ptr unsafe.Pointer) *TSavePictureDialog {
    s := new(TSavePictureDialog)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TSavePictureDialog) Free() {
    if s.instance != 0 {
        SavePictureDialog_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TSavePictureDialog) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TSavePictureDialog) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TSavePictureDialog) IsValid() bool {
    return s.instance != 0
}

// TSavePictureDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TSavePictureDialogClass() TClass {
    return SavePictureDialog_StaticClassType()
}

// Execute
func (s *TSavePictureDialog) Execute() bool {
    return SavePictureDialog_Execute(s.instance)
}

// FindComponent
func (s *TSavePictureDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(SavePictureDialog_FindComponent(s.instance, AName))
}

// GetNamePath
func (s *TSavePictureDialog) GetNamePath() string {
    return SavePictureDialog_GetNamePath(s.instance)
}

// HasParent
func (s *TSavePictureDialog) HasParent() bool {
    return SavePictureDialog_HasParent(s.instance)
}

// Assign
func (s *TSavePictureDialog) Assign(Source IObject) {
    SavePictureDialog_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TSavePictureDialog) DisposeOf() {
    SavePictureDialog_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TSavePictureDialog) ClassType() TClass {
    return SavePictureDialog_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TSavePictureDialog) ClassName() string {
    return SavePictureDialog_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TSavePictureDialog) InstanceSize() int32 {
    return SavePictureDialog_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TSavePictureDialog) InheritsFrom(AClass TClass) bool {
    return SavePictureDialog_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TSavePictureDialog) Equals(Obj IObject) bool {
    return SavePictureDialog_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TSavePictureDialog) GetHashCode() int32 {
    return SavePictureDialog_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TSavePictureDialog) ToString() string {
    return SavePictureDialog_ToString(s.instance)
}

// Filter
func (s *TSavePictureDialog) Filter() string {
    return SavePictureDialog_GetFilter(s.instance)
}

// SetFilter
func (s *TSavePictureDialog) SetFilter(value string) {
    SavePictureDialog_SetFilter(s.instance, value)
}

// Files
func (s *TSavePictureDialog) Files() *TStrings {
    return StringsFromInst(SavePictureDialog_GetFiles(s.instance))
}

// DefaultExt
func (s *TSavePictureDialog) DefaultExt() string {
    return SavePictureDialog_GetDefaultExt(s.instance)
}

// SetDefaultExt
func (s *TSavePictureDialog) SetDefaultExt(value string) {
    SavePictureDialog_SetDefaultExt(s.instance, value)
}

// FileName
func (s *TSavePictureDialog) FileName() string {
    return SavePictureDialog_GetFileName(s.instance)
}

// SetFileName
func (s *TSavePictureDialog) SetFileName(value string) {
    SavePictureDialog_SetFileName(s.instance, value)
}

// FilterIndex
func (s *TSavePictureDialog) FilterIndex() int32 {
    return SavePictureDialog_GetFilterIndex(s.instance)
}

// SetFilterIndex
func (s *TSavePictureDialog) SetFilterIndex(value int32) {
    SavePictureDialog_SetFilterIndex(s.instance, value)
}

// InitialDir
func (s *TSavePictureDialog) InitialDir() string {
    return SavePictureDialog_GetInitialDir(s.instance)
}

// SetInitialDir
func (s *TSavePictureDialog) SetInitialDir(value string) {
    SavePictureDialog_SetInitialDir(s.instance, value)
}

// Options
func (s *TSavePictureDialog) Options() TOpenOptions {
    return SavePictureDialog_GetOptions(s.instance)
}

// SetOptions
func (s *TSavePictureDialog) SetOptions(value TOpenOptions) {
    SavePictureDialog_SetOptions(s.instance, value)
}

// OptionsEx
func (s *TSavePictureDialog) OptionsEx() TOpenOptionsEx {
    return SavePictureDialog_GetOptionsEx(s.instance)
}

// SetOptionsEx
func (s *TSavePictureDialog) SetOptionsEx(value TOpenOptionsEx) {
    SavePictureDialog_SetOptionsEx(s.instance, value)
}

// Title
func (s *TSavePictureDialog) Title() string {
    return SavePictureDialog_GetTitle(s.instance)
}

// SetTitle
func (s *TSavePictureDialog) SetTitle(value string) {
    SavePictureDialog_SetTitle(s.instance, value)
}

// Handle
func (s *TSavePictureDialog) Handle() HWND {
    return SavePictureDialog_GetHandle(s.instance)
}

// SetOnClose
func (s *TSavePictureDialog) SetOnClose(fn TNotifyEvent) {
    SavePictureDialog_SetOnClose(s.instance, fn)
}

// SetOnShow
func (s *TSavePictureDialog) SetOnShow(fn TNotifyEvent) {
    SavePictureDialog_SetOnShow(s.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TSavePictureDialog) ComponentCount() int32 {
    return SavePictureDialog_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TSavePictureDialog) ComponentIndex() int32 {
    return SavePictureDialog_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TSavePictureDialog) SetComponentIndex(value int32) {
    SavePictureDialog_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TSavePictureDialog) Owner() *TComponent {
    return ComponentFromInst(SavePictureDialog_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TSavePictureDialog) Name() string {
    return SavePictureDialog_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TSavePictureDialog) SetName(value string) {
    SavePictureDialog_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TSavePictureDialog) Tag() int {
    return SavePictureDialog_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TSavePictureDialog) SetTag(value int) {
    SavePictureDialog_SetTag(s.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TSavePictureDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(SavePictureDialog_GetComponents(s.instance, AIndex))
}

