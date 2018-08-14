
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

type TTextAttributes struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// TextAttributesFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TextAttributesFromInst(inst uintptr) *TTextAttributes {
    t := new(TTextAttributes)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TextAttributesFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TextAttributesFromObj(obj IObject) *TTextAttributes {
    t := new(TTextAttributes)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TextAttributesFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TextAttributesFromUnsafePointer(ptr unsafe.Pointer) *TTextAttributes {
    t := new(TTextAttributes)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTextAttributes) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTextAttributes) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTextAttributes) IsValid() bool {
    return t.instance != 0
}

// TTextAttributesClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTextAttributesClass() TClass {
    return TextAttributes_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTextAttributes) Assign(Source IObject) {
    TextAttributes_Assign(t.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTextAttributes) GetNamePath() string {
    return TextAttributes_GetNamePath(t.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTextAttributes) DisposeOf() {
    TextAttributes_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTextAttributes) ClassType() TClass {
    return TextAttributes_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTextAttributes) ClassName() string {
    return TextAttributes_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTextAttributes) InstanceSize() int32 {
    return TextAttributes_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTextAttributes) InheritsFrom(AClass TClass) bool {
    return TextAttributes_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTextAttributes) Equals(Obj IObject) bool {
    return TextAttributes_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTextAttributes) GetHashCode() int32 {
    return TextAttributes_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTextAttributes) ToString() string {
    return TextAttributes_ToString(t.instance)
}

// Charset
func (t *TTextAttributes) Charset() TFontCharset {
    return TextAttributes_GetCharset(t.instance)
}

// SetCharset
func (t *TTextAttributes) SetCharset(value TFontCharset) {
    TextAttributes_SetCharset(t.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (t *TTextAttributes) Color() TColor {
    return TextAttributes_GetColor(t.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (t *TTextAttributes) SetColor(value TColor) {
    TextAttributes_SetColor(t.instance, value)
}

// ConsistentAttributes
func (t *TTextAttributes) ConsistentAttributes() TConsistentAttributes {
    return TextAttributes_GetConsistentAttributes(t.instance)
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTextAttributes) Name() string {
    return TextAttributes_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTextAttributes) SetName(value string) {
    TextAttributes_SetName(t.instance, value)
}

// Pitch
func (t *TTextAttributes) Pitch() TFontPitch {
    return TextAttributes_GetPitch(t.instance)
}

// SetPitch
func (t *TTextAttributes) SetPitch(value TFontPitch) {
    TextAttributes_SetPitch(t.instance, value)
}

// Protected
func (t *TTextAttributes) Protected() bool {
    return TextAttributes_GetProtected(t.instance)
}

// SetProtected
func (t *TTextAttributes) SetProtected(value bool) {
    TextAttributes_SetProtected(t.instance, value)
}

// Size
func (t *TTextAttributes) Size() int32 {
    return TextAttributes_GetSize(t.instance)
}

// SetSize
func (t *TTextAttributes) SetSize(value int32) {
    TextAttributes_SetSize(t.instance, value)
}

// Style
func (t *TTextAttributes) Style() TFontStyles {
    return TextAttributes_GetStyle(t.instance)
}

// SetStyle
func (t *TTextAttributes) SetStyle(value TFontStyles) {
    TextAttributes_SetStyle(t.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (t *TTextAttributes) Height() int32 {
    return TextAttributes_GetHeight(t.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (t *TTextAttributes) SetHeight(value int32) {
    TextAttributes_SetHeight(t.instance, value)
}

