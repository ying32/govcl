
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

type TGIFImage struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewGIFImage
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewGIFImage() *TGIFImage {
    g := new(TGIFImage)
    g.instance = GIFImage_Create()
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GIFImageFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func GIFImageFromInst(inst uintptr) *TGIFImage {
    g := new(TGIFImage)
    g.instance = inst
    g.ptr = unsafe.Pointer(inst)
    return g
}

// GIFImageFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func GIFImageFromObj(obj IObject) *TGIFImage {
    g := new(TGIFImage)
    g.instance = CheckPtr(obj)
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GIFImageFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func GIFImageFromUnsafePointer(ptr unsafe.Pointer) *TGIFImage {
    g := new(TGIFImage)
    g.instance = uintptr(ptr)
    g.ptr = ptr
    return g
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (g *TGIFImage) Free() {
    if g.instance != 0 {
        GIFImage_Free(g.instance)
        g.instance = 0
        g.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (g *TGIFImage) Instance() uintptr {
    return g.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (g *TGIFImage) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (g *TGIFImage) IsValid() bool {
    return g.instance != 0
}

// TGIFImageClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TGIFImageClass() TClass {
    return GIFImage_StaticClassType()
}

// SaveToStream
func (g *TGIFImage) SaveToStream(Stream IObject) {
    GIFImage_SaveToStream(g.instance, CheckPtr(Stream))
}

// LoadFromStream
func (g *TGIFImage) LoadFromStream(Stream IObject) {
    GIFImage_LoadFromStream(g.instance, CheckPtr(Stream))
}

// Add
func (g *TGIFImage) Add(Source IObject) *TGIFFrame {
    return GIFFrameFromInst(GIFImage_Add(g.instance, CheckPtr(Source)))
}

// Clear
func (g *TGIFImage) Clear() {
    GIFImage_Clear(g.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (g *TGIFImage) Assign(Source IObject) {
    GIFImage_Assign(g.instance, CheckPtr(Source))
}

// StopDraw
func (g *TGIFImage) StopDraw() {
    GIFImage_StopDraw(g.instance)
}

// SuspendDraw
func (g *TGIFImage) SuspendDraw() {
    GIFImage_SuspendDraw(g.instance)
}

// ResumeDraw
func (g *TGIFImage) ResumeDraw() {
    GIFImage_ResumeDraw(g.instance)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (g *TGIFImage) Equals(Obj IObject) bool {
    return GIFImage_Equals(g.instance, CheckPtr(Obj))
}

// LoadFromFile
func (g *TGIFImage) LoadFromFile(Filename string) {
    GIFImage_LoadFromFile(g.instance, Filename)
}

// SaveToFile
func (g *TGIFImage) SaveToFile(Filename string) {
    GIFImage_SaveToFile(g.instance, Filename)
}

// SetSize
func (g *TGIFImage) SetSize(AWidth int32, AHeight int32) {
    GIFImage_SetSize(g.instance, AWidth , AHeight)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (g *TGIFImage) GetNamePath() string {
    return GIFImage_GetNamePath(g.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (g *TGIFImage) DisposeOf() {
    GIFImage_DisposeOf(g.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (g *TGIFImage) ClassType() TClass {
    return GIFImage_ClassType(g.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (g *TGIFImage) ClassName() string {
    return GIFImage_ClassName(g.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (g *TGIFImage) InstanceSize() int32 {
    return GIFImage_InstanceSize(g.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (g *TGIFImage) InheritsFrom(AClass TClass) bool {
    return GIFImage_InheritsFrom(g.instance, AClass)
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (g *TGIFImage) GetHashCode() int32 {
    return GIFImage_GetHashCode(g.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (g *TGIFImage) ToString() string {
    return GIFImage_ToString(g.instance)
}

// Version
func (g *TGIFImage) Version() TGIFVersion {
    return GIFImage_GetVersion(g.instance)
}

// BitsPerPixel
func (g *TGIFImage) BitsPerPixel() int32 {
    return GIFImage_GetBitsPerPixel(g.instance)
}

// BackgroundColor
func (g *TGIFImage) BackgroundColor() TColor {
    return GIFImage_GetBackgroundColor(g.instance)
}

// SetBackgroundColor
func (g *TGIFImage) SetBackgroundColor(value TColor) {
    GIFImage_SetBackgroundColor(g.instance, value)
}

// AspectRatio
func (g *TGIFImage) AspectRatio() uint8 {
    return GIFImage_GetAspectRatio(g.instance)
}

// SetAspectRatio
func (g *TGIFImage) SetAspectRatio(value uint8) {
    GIFImage_SetAspectRatio(g.instance, value)
}

// IsTransparent
func (g *TGIFImage) IsTransparent() bool {
    return GIFImage_GetIsTransparent(g.instance)
}

// AnimationSpeed
func (g *TGIFImage) AnimationSpeed() int32 {
    return GIFImage_GetAnimationSpeed(g.instance)
}

// SetAnimationSpeed
func (g *TGIFImage) SetAnimationSpeed(value int32) {
    GIFImage_SetAnimationSpeed(g.instance, value)
}

// Bitmap
func (g *TGIFImage) Bitmap() *TBitmap {
    return BitmapFromInst(GIFImage_GetBitmap(g.instance))
}

// SetOnPaint
func (g *TGIFImage) SetOnPaint(fn TNotifyEvent) {
    GIFImage_SetOnPaint(g.instance, fn)
}

// Animate
func (g *TGIFImage) Animate() bool {
    return GIFImage_GetAnimate(g.instance)
}

// SetAnimate
func (g *TGIFImage) SetAnimate(value bool) {
    GIFImage_SetAnimate(g.instance, value)
}

// AnimateLoop
func (g *TGIFImage) AnimateLoop() TGIFAnimationLoop {
    return GIFImage_GetAnimateLoop(g.instance)
}

// SetAnimateLoop
func (g *TGIFImage) SetAnimateLoop(value TGIFAnimationLoop) {
    GIFImage_SetAnimateLoop(g.instance, value)
}

// ShouldDither
func (g *TGIFImage) ShouldDither() bool {
    return GIFImage_GetShouldDither(g.instance)
}

// Empty
func (g *TGIFImage) Empty() bool {
    return GIFImage_GetEmpty(g.instance)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (g *TGIFImage) Height() int32 {
    return GIFImage_GetHeight(g.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (g *TGIFImage) SetHeight(value int32) {
    GIFImage_SetHeight(g.instance, value)
}

// Modified
// CN: 获取修改。
// EN: Get modified.
func (g *TGIFImage) Modified() bool {
    return GIFImage_GetModified(g.instance)
}

// SetModified
// CN: 设置修改。
// EN: Set modified.
func (g *TGIFImage) SetModified(value bool) {
    GIFImage_SetModified(g.instance, value)
}

// PaletteModified
func (g *TGIFImage) PaletteModified() bool {
    return GIFImage_GetPaletteModified(g.instance)
}

// SetPaletteModified
func (g *TGIFImage) SetPaletteModified(value bool) {
    GIFImage_SetPaletteModified(g.instance, value)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (g *TGIFImage) Transparent() bool {
    return GIFImage_GetTransparent(g.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (g *TGIFImage) SetTransparent(value bool) {
    GIFImage_SetTransparent(g.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (g *TGIFImage) Width() int32 {
    return GIFImage_GetWidth(g.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (g *TGIFImage) SetWidth(value int32) {
    GIFImage_SetWidth(g.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (g *TGIFImage) SetOnChange(fn TNotifyEvent) {
    GIFImage_SetOnChange(g.instance, fn)
}

