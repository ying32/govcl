
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

type TGraphic struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewGraphic
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewGraphic() *TGraphic {
    g := new(TGraphic)
    g.instance = Graphic_Create()
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GraphicFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func GraphicFromInst(inst uintptr) *TGraphic {
    g := new(TGraphic)
    g.instance = inst
    g.ptr = unsafe.Pointer(inst)
    return g
}

// GraphicFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func GraphicFromObj(obj IObject) *TGraphic {
    g := new(TGraphic)
    g.instance = CheckPtr(obj)
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GraphicFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func GraphicFromUnsafePointer(ptr unsafe.Pointer) *TGraphic {
    g := new(TGraphic)
    g.instance = uintptr(ptr)
    g.ptr = ptr
    return g
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (g *TGraphic) Free() {
    if g.instance != 0 {
        Graphic_Free(g.instance)
        g.instance = 0
        g.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (g *TGraphic) Instance() uintptr {
    return g.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (g *TGraphic) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (g *TGraphic) IsValid() bool {
    return g.instance != 0
}

// TGraphicClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TGraphicClass() TClass {
    return Graphic_StaticClassType()
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (g *TGraphic) Equals(Obj IObject) bool {
    return Graphic_Equals(g.instance, CheckPtr(Obj))
}

// LoadFromFile
func (g *TGraphic) LoadFromFile(Filename string) {
    Graphic_LoadFromFile(g.instance, Filename)
}

// SaveToFile
func (g *TGraphic) SaveToFile(Filename string) {
    Graphic_SaveToFile(g.instance, Filename)
}

// LoadFromStream
func (g *TGraphic) LoadFromStream(Stream IObject) {
    Graphic_LoadFromStream(g.instance, CheckPtr(Stream))
}

// SaveToStream
func (g *TGraphic) SaveToStream(Stream IObject) {
    Graphic_SaveToStream(g.instance, CheckPtr(Stream))
}

// SetSize
func (g *TGraphic) SetSize(AWidth int32, AHeight int32) {
    Graphic_SetSize(g.instance, AWidth , AHeight)
}

// Assign
func (g *TGraphic) Assign(Source IObject) {
    Graphic_Assign(g.instance, CheckPtr(Source))
}

// GetNamePath
func (g *TGraphic) GetNamePath() string {
    return Graphic_GetNamePath(g.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (g *TGraphic) DisposeOf() {
    Graphic_DisposeOf(g.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (g *TGraphic) ClassType() TClass {
    return Graphic_ClassType(g.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (g *TGraphic) ClassName() string {
    return Graphic_ClassName(g.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (g *TGraphic) InstanceSize() int32 {
    return Graphic_InstanceSize(g.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (g *TGraphic) InheritsFrom(AClass TClass) bool {
    return Graphic_InheritsFrom(g.instance, AClass)
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (g *TGraphic) GetHashCode() int32 {
    return Graphic_GetHashCode(g.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (g *TGraphic) ToString() string {
    return Graphic_ToString(g.instance)
}

// Empty
func (g *TGraphic) Empty() bool {
    return Graphic_GetEmpty(g.instance)
}

// Height
func (g *TGraphic) Height() int32 {
    return Graphic_GetHeight(g.instance)
}

// SetHeight
func (g *TGraphic) SetHeight(value int32) {
    Graphic_SetHeight(g.instance, value)
}

// Modified
func (g *TGraphic) Modified() bool {
    return Graphic_GetModified(g.instance)
}

// SetModified
func (g *TGraphic) SetModified(value bool) {
    Graphic_SetModified(g.instance, value)
}

// PaletteModified
func (g *TGraphic) PaletteModified() bool {
    return Graphic_GetPaletteModified(g.instance)
}

// SetPaletteModified
func (g *TGraphic) SetPaletteModified(value bool) {
    Graphic_SetPaletteModified(g.instance, value)
}

// Transparent
func (g *TGraphic) Transparent() bool {
    return Graphic_GetTransparent(g.instance)
}

// SetTransparent
func (g *TGraphic) SetTransparent(value bool) {
    Graphic_SetTransparent(g.instance, value)
}

// Width
func (g *TGraphic) Width() int32 {
    return Graphic_GetWidth(g.instance)
}

// SetWidth
func (g *TGraphic) SetWidth(value int32) {
    Graphic_SetWidth(g.instance, value)
}

// SetOnChange
func (g *TGraphic) SetOnChange(fn TNotifyEvent) {
    Graphic_SetOnChange(g.instance, fn)
}

