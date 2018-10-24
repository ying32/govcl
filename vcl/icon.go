
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

type TIcon struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewIcon
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewIcon() *TIcon {
    i := new(TIcon)
    i.instance = Icon_Create()
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// IconFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func IconFromInst(inst uintptr) *TIcon {
    i := new(TIcon)
    i.instance = inst
    i.ptr = unsafe.Pointer(inst)
    return i
}

// IconFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func IconFromObj(obj IObject) *TIcon {
    i := new(TIcon)
    i.instance = CheckPtr(obj)
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// IconFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func IconFromUnsafePointer(ptr unsafe.Pointer) *TIcon {
    i := new(TIcon)
    i.instance = uintptr(ptr)
    i.ptr = ptr
    return i
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (i *TIcon) Free() {
    if i.instance != 0 {
        Icon_Free(i.instance)
        i.instance = 0
        i.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (i *TIcon) Instance() uintptr {
    return i.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (i *TIcon) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (i *TIcon) IsValid() bool {
    return i.instance != 0
}

// TIconClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TIconClass() TClass {
    return Icon_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (i *TIcon) Assign(Source IObject) {
    Icon_Assign(i.instance, CheckPtr(Source))
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (i *TIcon) HandleAllocated() bool {
    return Icon_HandleAllocated(i.instance)
}

// LoadFromStream
// CN: 文件流加载。
// EN: .
func (i *TIcon) LoadFromStream(Stream IObject) {
    Icon_LoadFromStream(i.instance, CheckPtr(Stream))
}

// SaveToStream
// CN: 保存至流。
// EN: .
func (i *TIcon) SaveToStream(Stream IObject) {
    Icon_SaveToStream(i.instance, CheckPtr(Stream))
}

// SetSize
func (i *TIcon) SetSize(AWidth int32, AHeight int32) {
    Icon_SetSize(i.instance, AWidth , AHeight)
}

// LoadFromResourceName
func (i *TIcon) LoadFromResourceName(Instance uintptr, ResName string) {
    Icon_LoadFromResourceName(i.instance, Instance , ResName)
}

// LoadFromResourceID
func (i *TIcon) LoadFromResourceID(Instance uintptr, ResID int32) {
    Icon_LoadFromResourceID(i.instance, Instance , ResID)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (i *TIcon) Equals(Obj IObject) bool {
    return Icon_Equals(i.instance, CheckPtr(Obj))
}

// LoadFromFile
// CN: 从文件加载。
// EN: .
func (i *TIcon) LoadFromFile(Filename string) {
    Icon_LoadFromFile(i.instance, Filename)
}

// SaveToFile
// CN: 保存至文件。
// EN: .
func (i *TIcon) SaveToFile(Filename string) {
    Icon_SaveToFile(i.instance, Filename)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (i *TIcon) GetNamePath() string {
    return Icon_GetNamePath(i.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (i *TIcon) DisposeOf() {
    Icon_DisposeOf(i.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (i *TIcon) ClassType() TClass {
    return Icon_ClassType(i.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (i *TIcon) ClassName() string {
    return Icon_ClassName(i.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (i *TIcon) InstanceSize() int32 {
    return Icon_InstanceSize(i.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (i *TIcon) InheritsFrom(AClass TClass) bool {
    return Icon_InheritsFrom(i.instance, AClass)
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (i *TIcon) GetHashCode() int32 {
    return Icon_GetHashCode(i.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (i *TIcon) ToString() string {
    return Icon_ToString(i.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (i *TIcon) Handle() HICON {
    return Icon_GetHandle(i.instance)
}

// SetHandle
// CN: 设置控件句柄。
// EN: Set Control handle.
func (i *TIcon) SetHandle(value HICON) {
    Icon_SetHandle(i.instance, value)
}

// Empty
func (i *TIcon) Empty() bool {
    return Icon_GetEmpty(i.instance)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (i *TIcon) Height() int32 {
    return Icon_GetHeight(i.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (i *TIcon) SetHeight(value int32) {
    Icon_SetHeight(i.instance, value)
}

// Modified
// CN: 获取修改。
// EN: Get modified.
func (i *TIcon) Modified() bool {
    return Icon_GetModified(i.instance)
}

// SetModified
// CN: 设置修改。
// EN: Set modified.
func (i *TIcon) SetModified(value bool) {
    Icon_SetModified(i.instance, value)
}

// Palette
func (i *TIcon) Palette() HPALETTE {
    return Icon_GetPalette(i.instance)
}

// SetPalette
func (i *TIcon) SetPalette(value HPALETTE) {
    Icon_SetPalette(i.instance, value)
}

// PaletteModified
func (i *TIcon) PaletteModified() bool {
    return Icon_GetPaletteModified(i.instance)
}

// SetPaletteModified
func (i *TIcon) SetPaletteModified(value bool) {
    Icon_SetPaletteModified(i.instance, value)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (i *TIcon) Transparent() bool {
    return Icon_GetTransparent(i.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (i *TIcon) SetTransparent(value bool) {
    Icon_SetTransparent(i.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (i *TIcon) Width() int32 {
    return Icon_GetWidth(i.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (i *TIcon) SetWidth(value int32) {
    Icon_SetWidth(i.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (i *TIcon) SetOnChange(fn TNotifyEvent) {
    Icon_SetOnChange(i.instance, fn)
}

