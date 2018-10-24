
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

type TJPEGImage struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewJPEGImage
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewJPEGImage() *TJPEGImage {
    j := new(TJPEGImage)
    j.instance = JPEGImage_Create()
    j.ptr = unsafe.Pointer(j.instance)
    return j
}

// JPEGImageFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func JPEGImageFromInst(inst uintptr) *TJPEGImage {
    j := new(TJPEGImage)
    j.instance = inst
    j.ptr = unsafe.Pointer(inst)
    return j
}

// JPEGImageFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func JPEGImageFromObj(obj IObject) *TJPEGImage {
    j := new(TJPEGImage)
    j.instance = CheckPtr(obj)
    j.ptr = unsafe.Pointer(j.instance)
    return j
}

// JPEGImageFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func JPEGImageFromUnsafePointer(ptr unsafe.Pointer) *TJPEGImage {
    j := new(TJPEGImage)
    j.instance = uintptr(ptr)
    j.ptr = ptr
    return j
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (j *TJPEGImage) Free() {
    if j.instance != 0 {
        JPEGImage_Free(j.instance)
        j.instance = 0
        j.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (j *TJPEGImage) Instance() uintptr {
    return j.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (j *TJPEGImage) UnsafeAddr() unsafe.Pointer {
    return j.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (j *TJPEGImage) IsValid() bool {
    return j.instance != 0
}

// TJPEGImageClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TJPEGImageClass() TClass {
    return JPEGImage_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (j *TJPEGImage) Assign(Source IObject) {
    JPEGImage_Assign(j.instance, CheckPtr(Source))
}

// LoadFromStream
// CN: 文件流加载。
// EN: .
func (j *TJPEGImage) LoadFromStream(Stream IObject) {
    JPEGImage_LoadFromStream(j.instance, CheckPtr(Stream))
}

// SaveToStream
// CN: 保存至流。
// EN: .
func (j *TJPEGImage) SaveToStream(Stream IObject) {
    JPEGImage_SaveToStream(j.instance, CheckPtr(Stream))
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (j *TJPEGImage) Equals(Obj IObject) bool {
    return JPEGImage_Equals(j.instance, CheckPtr(Obj))
}

// LoadFromFile
// CN: 从文件加载。
// EN: .
func (j *TJPEGImage) LoadFromFile(Filename string) {
    JPEGImage_LoadFromFile(j.instance, Filename)
}

// SaveToFile
// CN: 保存至文件。
// EN: .
func (j *TJPEGImage) SaveToFile(Filename string) {
    JPEGImage_SaveToFile(j.instance, Filename)
}

// SetSize
func (j *TJPEGImage) SetSize(AWidth int32, AHeight int32) {
    JPEGImage_SetSize(j.instance, AWidth , AHeight)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (j *TJPEGImage) GetNamePath() string {
    return JPEGImage_GetNamePath(j.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (j *TJPEGImage) DisposeOf() {
    JPEGImage_DisposeOf(j.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (j *TJPEGImage) ClassType() TClass {
    return JPEGImage_ClassType(j.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (j *TJPEGImage) ClassName() string {
    return JPEGImage_ClassName(j.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (j *TJPEGImage) InstanceSize() int32 {
    return JPEGImage_InstanceSize(j.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (j *TJPEGImage) InheritsFrom(AClass TClass) bool {
    return JPEGImage_InheritsFrom(j.instance, AClass)
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (j *TJPEGImage) GetHashCode() int32 {
    return JPEGImage_GetHashCode(j.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (j *TJPEGImage) ToString() string {
    return JPEGImage_ToString(j.instance)
}

// PixelFormat
func (j *TJPEGImage) PixelFormat() TJPEGPixelFormat {
    return JPEGImage_GetPixelFormat(j.instance)
}

// SetPixelFormat
func (j *TJPEGImage) SetPixelFormat(value TJPEGPixelFormat) {
    JPEGImage_SetPixelFormat(j.instance, value)
}

// ProgressiveDisplay
func (j *TJPEGImage) ProgressiveDisplay() bool {
    return JPEGImage_GetProgressiveDisplay(j.instance)
}

// SetProgressiveDisplay
func (j *TJPEGImage) SetProgressiveDisplay(value bool) {
    JPEGImage_SetProgressiveDisplay(j.instance, value)
}

// Performance
func (j *TJPEGImage) Performance() TJPEGPerformance {
    return JPEGImage_GetPerformance(j.instance)
}

// SetPerformance
func (j *TJPEGImage) SetPerformance(value TJPEGPerformance) {
    JPEGImage_SetPerformance(j.instance, value)
}

// Scale
func (j *TJPEGImage) Scale() TJPEGScale {
    return JPEGImage_GetScale(j.instance)
}

// SetScale
func (j *TJPEGImage) SetScale(value TJPEGScale) {
    JPEGImage_SetScale(j.instance, value)
}

// Smoothing
func (j *TJPEGImage) Smoothing() bool {
    return JPEGImage_GetSmoothing(j.instance)
}

// SetSmoothing
func (j *TJPEGImage) SetSmoothing(value bool) {
    JPEGImage_SetSmoothing(j.instance, value)
}

// Canvas
// CN: 获取画布。
// EN: .
func (j *TJPEGImage) Canvas() *TCanvas {
    return CanvasFromInst(JPEGImage_GetCanvas(j.instance))
}

// Empty
func (j *TJPEGImage) Empty() bool {
    return JPEGImage_GetEmpty(j.instance)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (j *TJPEGImage) Height() int32 {
    return JPEGImage_GetHeight(j.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (j *TJPEGImage) SetHeight(value int32) {
    JPEGImage_SetHeight(j.instance, value)
}

// Modified
// CN: 获取修改。
// EN: Get modified.
func (j *TJPEGImage) Modified() bool {
    return JPEGImage_GetModified(j.instance)
}

// SetModified
// CN: 设置修改。
// EN: Set modified.
func (j *TJPEGImage) SetModified(value bool) {
    JPEGImage_SetModified(j.instance, value)
}

// Palette
func (j *TJPEGImage) Palette() HPALETTE {
    return JPEGImage_GetPalette(j.instance)
}

// SetPalette
func (j *TJPEGImage) SetPalette(value HPALETTE) {
    JPEGImage_SetPalette(j.instance, value)
}

// PaletteModified
func (j *TJPEGImage) PaletteModified() bool {
    return JPEGImage_GetPaletteModified(j.instance)
}

// SetPaletteModified
func (j *TJPEGImage) SetPaletteModified(value bool) {
    JPEGImage_SetPaletteModified(j.instance, value)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (j *TJPEGImage) Transparent() bool {
    return JPEGImage_GetTransparent(j.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (j *TJPEGImage) SetTransparent(value bool) {
    JPEGImage_SetTransparent(j.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (j *TJPEGImage) Width() int32 {
    return JPEGImage_GetWidth(j.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (j *TJPEGImage) SetWidth(value int32) {
    JPEGImage_SetWidth(j.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (j *TJPEGImage) SetOnChange(fn TNotifyEvent) {
    JPEGImage_SetOnChange(j.instance, fn)
}

