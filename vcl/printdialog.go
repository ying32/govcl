
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

type TPrintDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPrintDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPrintDialog(owner IComponent) *TPrintDialog {
    p := new(TPrintDialog)
    p.instance = PrintDialog_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PrintDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PrintDialogFromInst(inst uintptr) *TPrintDialog {
    p := new(TPrintDialog)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PrintDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PrintDialogFromObj(obj IObject) *TPrintDialog {
    p := new(TPrintDialog)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PrintDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PrintDialogFromUnsafePointer(ptr unsafe.Pointer) *TPrintDialog {
    p := new(TPrintDialog)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPrintDialog) Free() {
    if p.instance != 0 {
        PrintDialog_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPrintDialog) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPrintDialog) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPrintDialog) IsValid() bool {
    return p.instance != 0
}

// TPrintDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPrintDialogClass() TClass {
    return PrintDialog_StaticClassType()
}

// Execute
// CN: 执行。
// EN: .
func (p *TPrintDialog) Execute() bool {
    return PrintDialog_Execute(p.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (p *TPrintDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PrintDialog_FindComponent(p.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPrintDialog) GetNamePath() string {
    return PrintDialog_GetNamePath(p.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (p *TPrintDialog) HasParent() bool {
    return PrintDialog_HasParent(p.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPrintDialog) Assign(Source IObject) {
    PrintDialog_Assign(p.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPrintDialog) DisposeOf() {
    PrintDialog_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPrintDialog) ClassType() TClass {
    return PrintDialog_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPrintDialog) ClassName() string {
    return PrintDialog_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPrintDialog) InstanceSize() int32 {
    return PrintDialog_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPrintDialog) InheritsFrom(AClass TClass) bool {
    return PrintDialog_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPrintDialog) Equals(Obj IObject) bool {
    return PrintDialog_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPrintDialog) GetHashCode() int32 {
    return PrintDialog_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPrintDialog) ToString() string {
    return PrintDialog_ToString(p.instance)
}

// Collate
func (p *TPrintDialog) Collate() bool {
    return PrintDialog_GetCollate(p.instance)
}

// SetCollate
func (p *TPrintDialog) SetCollate(value bool) {
    PrintDialog_SetCollate(p.instance, value)
}

// Copies
func (p *TPrintDialog) Copies() int32 {
    return PrintDialog_GetCopies(p.instance)
}

// SetCopies
func (p *TPrintDialog) SetCopies(value int32) {
    PrintDialog_SetCopies(p.instance, value)
}

// FromPage
func (p *TPrintDialog) FromPage() int32 {
    return PrintDialog_GetFromPage(p.instance)
}

// SetFromPage
func (p *TPrintDialog) SetFromPage(value int32) {
    PrintDialog_SetFromPage(p.instance, value)
}

// MinPage
func (p *TPrintDialog) MinPage() int32 {
    return PrintDialog_GetMinPage(p.instance)
}

// SetMinPage
func (p *TPrintDialog) SetMinPage(value int32) {
    PrintDialog_SetMinPage(p.instance, value)
}

// MaxPage
func (p *TPrintDialog) MaxPage() int32 {
    return PrintDialog_GetMaxPage(p.instance)
}

// SetMaxPage
func (p *TPrintDialog) SetMaxPage(value int32) {
    PrintDialog_SetMaxPage(p.instance, value)
}

// Options
func (p *TPrintDialog) Options() TPrintDialogOptions {
    return PrintDialog_GetOptions(p.instance)
}

// SetOptions
func (p *TPrintDialog) SetOptions(value TPrintDialogOptions) {
    PrintDialog_SetOptions(p.instance, value)
}

// PrintToFile
func (p *TPrintDialog) PrintToFile() bool {
    return PrintDialog_GetPrintToFile(p.instance)
}

// SetPrintToFile
func (p *TPrintDialog) SetPrintToFile(value bool) {
    PrintDialog_SetPrintToFile(p.instance, value)
}

// PrintRange
func (p *TPrintDialog) PrintRange() TPrintRange {
    return PrintDialog_GetPrintRange(p.instance)
}

// SetPrintRange
func (p *TPrintDialog) SetPrintRange(value TPrintRange) {
    PrintDialog_SetPrintRange(p.instance, value)
}

// ToPage
func (p *TPrintDialog) ToPage() int32 {
    return PrintDialog_GetToPage(p.instance)
}

// SetToPage
func (p *TPrintDialog) SetToPage(value int32) {
    PrintDialog_SetToPage(p.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (p *TPrintDialog) Handle() HWND {
    return PrintDialog_GetHandle(p.instance)
}

// SetOnClose
func (p *TPrintDialog) SetOnClose(fn TNotifyEvent) {
    PrintDialog_SetOnClose(p.instance, fn)
}

// SetOnShow
// CN: 设置显示事件。
// EN: .
func (p *TPrintDialog) SetOnShow(fn TNotifyEvent) {
    PrintDialog_SetOnShow(p.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPrintDialog) ComponentCount() int32 {
    return PrintDialog_GetComponentCount(p.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (p *TPrintDialog) ComponentIndex() int32 {
    return PrintDialog_GetComponentIndex(p.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (p *TPrintDialog) SetComponentIndex(value int32) {
    PrintDialog_SetComponentIndex(p.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPrintDialog) Owner() *TComponent {
    return ComponentFromInst(PrintDialog_GetOwner(p.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPrintDialog) Name() string {
    return PrintDialog_GetName(p.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPrintDialog) SetName(value string) {
    PrintDialog_SetName(p.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPrintDialog) Tag() int {
    return PrintDialog_GetTag(p.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPrintDialog) SetTag(value int) {
    PrintDialog_SetTag(p.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPrintDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PrintDialog_GetComponents(p.instance, AIndex))
}

