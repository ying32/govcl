
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

type TPrinterSetupDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPrinterSetupDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPrinterSetupDialog(owner IComponent) *TPrinterSetupDialog {
    p := new(TPrinterSetupDialog)
    p.instance = PrinterSetupDialog_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PrinterSetupDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PrinterSetupDialogFromInst(inst uintptr) *TPrinterSetupDialog {
    p := new(TPrinterSetupDialog)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PrinterSetupDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PrinterSetupDialogFromObj(obj IObject) *TPrinterSetupDialog {
    p := new(TPrinterSetupDialog)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PrinterSetupDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PrinterSetupDialogFromUnsafePointer(ptr unsafe.Pointer) *TPrinterSetupDialog {
    p := new(TPrinterSetupDialog)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPrinterSetupDialog) Free() {
    if p.instance != 0 {
        PrinterSetupDialog_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPrinterSetupDialog) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPrinterSetupDialog) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPrinterSetupDialog) IsValid() bool {
    return p.instance != 0
}

// TPrinterSetupDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPrinterSetupDialogClass() TClass {
    return PrinterSetupDialog_StaticClassType()
}

// Execute
// CN: 执行。
// EN: .
func (p *TPrinterSetupDialog) Execute() bool {
    return PrinterSetupDialog_Execute(p.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (p *TPrinterSetupDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PrinterSetupDialog_FindComponent(p.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPrinterSetupDialog) GetNamePath() string {
    return PrinterSetupDialog_GetNamePath(p.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (p *TPrinterSetupDialog) HasParent() bool {
    return PrinterSetupDialog_HasParent(p.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPrinterSetupDialog) Assign(Source IObject) {
    PrinterSetupDialog_Assign(p.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPrinterSetupDialog) DisposeOf() {
    PrinterSetupDialog_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPrinterSetupDialog) ClassType() TClass {
    return PrinterSetupDialog_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPrinterSetupDialog) ClassName() string {
    return PrinterSetupDialog_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPrinterSetupDialog) InstanceSize() int32 {
    return PrinterSetupDialog_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPrinterSetupDialog) InheritsFrom(AClass TClass) bool {
    return PrinterSetupDialog_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPrinterSetupDialog) Equals(Obj IObject) bool {
    return PrinterSetupDialog_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPrinterSetupDialog) GetHashCode() int32 {
    return PrinterSetupDialog_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPrinterSetupDialog) ToString() string {
    return PrinterSetupDialog_ToString(p.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (p *TPrinterSetupDialog) Handle() HWND {
    return PrinterSetupDialog_GetHandle(p.instance)
}

// SetOnClose
func (p *TPrinterSetupDialog) SetOnClose(fn TNotifyEvent) {
    PrinterSetupDialog_SetOnClose(p.instance, fn)
}

// SetOnShow
// CN: 设置显示事件。
// EN: .
func (p *TPrinterSetupDialog) SetOnShow(fn TNotifyEvent) {
    PrinterSetupDialog_SetOnShow(p.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPrinterSetupDialog) ComponentCount() int32 {
    return PrinterSetupDialog_GetComponentCount(p.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (p *TPrinterSetupDialog) ComponentIndex() int32 {
    return PrinterSetupDialog_GetComponentIndex(p.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (p *TPrinterSetupDialog) SetComponentIndex(value int32) {
    PrinterSetupDialog_SetComponentIndex(p.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPrinterSetupDialog) Owner() *TComponent {
    return ComponentFromInst(PrinterSetupDialog_GetOwner(p.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPrinterSetupDialog) Name() string {
    return PrinterSetupDialog_GetName(p.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPrinterSetupDialog) SetName(value string) {
    PrinterSetupDialog_SetName(p.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPrinterSetupDialog) Tag() int {
    return PrinterSetupDialog_GetTag(p.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPrinterSetupDialog) SetTag(value int) {
    PrinterSetupDialog_SetTag(p.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPrinterSetupDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PrinterSetupDialog_GetComponents(p.instance, AIndex))
}

