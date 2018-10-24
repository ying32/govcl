
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

type TPrinter struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPrinter
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPrinter() *TPrinter {
    p := new(TPrinter)
    p.instance = Printer_Create()
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PrinterFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PrinterFromInst(inst uintptr) *TPrinter {
    p := new(TPrinter)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PrinterFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PrinterFromObj(obj IObject) *TPrinter {
    p := new(TPrinter)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PrinterFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PrinterFromUnsafePointer(ptr unsafe.Pointer) *TPrinter {
    p := new(TPrinter)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPrinter) Free() {
    if p.instance != 0 {
        Printer_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPrinter) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPrinter) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPrinter) IsValid() bool {
    return p.instance != 0
}

// TPrinterClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPrinterClass() TClass {
    return Printer_StaticClassType()
}

// Abort
func (p *TPrinter) Abort() {
    Printer_Abort(p.instance)
}

// BeginDoc
func (p *TPrinter) BeginDoc() {
    Printer_BeginDoc(p.instance)
}

// EndDoc
func (p *TPrinter) EndDoc() {
    Printer_EndDoc(p.instance)
}

// NewPage
func (p *TPrinter) NewPage() {
    Printer_NewPage(p.instance)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (p *TPrinter) Refresh() {
    Printer_Refresh(p.instance)
}

// GetPrinter
func (p *TPrinter) GetPrinter(ADevice string, ADriver string, APort string, ADeviceMode *uintptr) {
    Printer_GetPrinter(p.instance, ADevice , ADriver , APort , ADeviceMode)
}

// SetPrinter
func (p *TPrinter) SetPrinter(ADevice string, ADriver string, APort string, ADeviceMode uintptr) {
    Printer_SetPrinter(p.instance, ADevice , ADriver , APort , ADeviceMode)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPrinter) DisposeOf() {
    Printer_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPrinter) ClassType() TClass {
    return Printer_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPrinter) ClassName() string {
    return Printer_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPrinter) InstanceSize() int32 {
    return Printer_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPrinter) InheritsFrom(AClass TClass) bool {
    return Printer_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPrinter) Equals(Obj IObject) bool {
    return Printer_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPrinter) GetHashCode() int32 {
    return Printer_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPrinter) ToString() string {
    return Printer_ToString(p.instance)
}

// Aborted
func (p *TPrinter) Aborted() bool {
    return Printer_GetAborted(p.instance)
}

// Canvas
// CN: 获取画布。
// EN: .
func (p *TPrinter) Canvas() *TCanvas {
    return CanvasFromInst(Printer_GetCanvas(p.instance))
}

// Capabilities
func (p *TPrinter) Capabilities() TPrinterCapabilities {
    return Printer_GetCapabilities(p.instance)
}

// Copies
func (p *TPrinter) Copies() int32 {
    return Printer_GetCopies(p.instance)
}

// SetCopies
func (p *TPrinter) SetCopies(value int32) {
    Printer_SetCopies(p.instance, value)
}

// Fonts
func (p *TPrinter) Fonts() *TStrings {
    return StringsFromInst(Printer_GetFonts(p.instance))
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (p *TPrinter) Handle() HDC {
    return Printer_GetHandle(p.instance)
}

// Orientation
func (p *TPrinter) Orientation() TPrinterOrientation {
    return Printer_GetOrientation(p.instance)
}

// SetOrientation
func (p *TPrinter) SetOrientation(value TPrinterOrientation) {
    Printer_SetOrientation(p.instance, value)
}

// PageHeight
func (p *TPrinter) PageHeight() int32 {
    return Printer_GetPageHeight(p.instance)
}

// PageWidth
func (p *TPrinter) PageWidth() int32 {
    return Printer_GetPageWidth(p.instance)
}

// PageNumber
func (p *TPrinter) PageNumber() int32 {
    return Printer_GetPageNumber(p.instance)
}

// PrinterIndex
func (p *TPrinter) PrinterIndex() int32 {
    return Printer_GetPrinterIndex(p.instance)
}

// SetPrinterIndex
func (p *TPrinter) SetPrinterIndex(value int32) {
    Printer_SetPrinterIndex(p.instance, value)
}

// Printing
func (p *TPrinter) Printing() bool {
    return Printer_GetPrinting(p.instance)
}

// Printers
func (p *TPrinter) Printers() *TStrings {
    return StringsFromInst(Printer_GetPrinters(p.instance))
}

// Title
func (p *TPrinter) Title() string {
    return Printer_GetTitle(p.instance)
}

// SetTitle
func (p *TPrinter) SetTitle(value string) {
    Printer_SetTitle(p.instance, value)
}

