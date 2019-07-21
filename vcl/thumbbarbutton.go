
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

type TThumbBarButton struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewThumbBarButton
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewThumbBarButton() *TThumbBarButton {
    t := new(TThumbBarButton)
    t.instance = ThumbBarButton_Create()
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// ThumbBarButtonFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ThumbBarButtonFromInst(inst uintptr) *TThumbBarButton {
    t := new(TThumbBarButton)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// ThumbBarButtonFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ThumbBarButtonFromObj(obj IObject) *TThumbBarButton {
    t := new(TThumbBarButton)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// ThumbBarButtonFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ThumbBarButtonFromUnsafePointer(ptr unsafe.Pointer) *TThumbBarButton {
    t := new(TThumbBarButton)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TThumbBarButton) Free() {
    if t.instance != 0 {
        ThumbBarButton_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TThumbBarButton) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TThumbBarButton) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TThumbBarButton) IsValid() bool {
    return t.instance != 0
}

// TThumbBarButtonClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TThumbBarButtonClass() TClass {
    return ThumbBarButton_StaticClassType()
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TThumbBarButton) GetNamePath() string {
    return ThumbBarButton_GetNamePath(t.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TThumbBarButton) Assign(Source IObject) {
    ThumbBarButton_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TThumbBarButton) DisposeOf() {
    ThumbBarButton_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TThumbBarButton) ClassType() TClass {
    return ThumbBarButton_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TThumbBarButton) ClassName() string {
    return ThumbBarButton_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TThumbBarButton) InstanceSize() int32 {
    return ThumbBarButton_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TThumbBarButton) InheritsFrom(AClass TClass) bool {
    return ThumbBarButton_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TThumbBarButton) Equals(Obj IObject) bool {
    return ThumbBarButton_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TThumbBarButton) GetHashCode() int32 {
    return ThumbBarButton_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TThumbBarButton) ToString() string {
    return ThumbBarButton_ToString(t.instance)
}

// Action
func (t *TThumbBarButton) Action() *TAction {
    return ActionFromInst(ThumbBarButton_GetAction(t.instance))
}

// SetAction
func (t *TThumbBarButton) SetAction(value IComponent) {
    ThumbBarButton_SetAction(t.instance, CheckPtr(value))
}

// ButtonState
func (t *TThumbBarButton) ButtonState() TThumbButtonStates {
    return ThumbBarButton_GetButtonState(t.instance)
}

// SetButtonState
func (t *TThumbBarButton) SetButtonState(value TThumbButtonStates) {
    ThumbBarButton_SetButtonState(t.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (t *TThumbBarButton) Hint() string {
    return ThumbBarButton_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (t *TThumbBarButton) SetHint(value string) {
    ThumbBarButton_SetHint(t.instance, value)
}

// Icon
// CN: 获取图标。
// EN: Get icon.
func (t *TThumbBarButton) Icon() *TIcon {
    return IconFromInst(ThumbBarButton_GetIcon(t.instance))
}

// SetIcon
// CN: 设置图标。
// EN: Set icon.
func (t *TThumbBarButton) SetIcon(value *TIcon) {
    ThumbBarButton_SetIcon(t.instance, CheckPtr(value))
}

// Collection
func (t *TThumbBarButton) Collection() *TCollection {
    return CollectionFromInst(ThumbBarButton_GetCollection(t.instance))
}

// SetCollection
func (t *TThumbBarButton) SetCollection(value *TCollection) {
    ThumbBarButton_SetCollection(t.instance, CheckPtr(value))
}

// Index
func (t *TThumbBarButton) Index() int32 {
    return ThumbBarButton_GetIndex(t.instance)
}

// SetIndex
func (t *TThumbBarButton) SetIndex(value int32) {
    ThumbBarButton_SetIndex(t.instance, value)
}

// DisplayName
func (t *TThumbBarButton) DisplayName() string {
    return ThumbBarButton_GetDisplayName(t.instance)
}

// SetDisplayName
func (t *TThumbBarButton) SetDisplayName(value string) {
    ThumbBarButton_SetDisplayName(t.instance, value)
}

