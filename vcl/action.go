
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

type TAction struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewAction
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewAction(owner IComponent) *TAction {
    a := new(TAction)
    a.instance = Action_Create(CheckPtr(owner))
    a.ptr = unsafe.Pointer(a.instance)
    return a
}

// ActionFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ActionFromInst(inst uintptr) *TAction {
    a := new(TAction)
    a.instance = inst
    a.ptr = unsafe.Pointer(inst)
    return a
}

// ActionFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ActionFromObj(obj IObject) *TAction {
    a := new(TAction)
    a.instance = CheckPtr(obj)
    a.ptr = unsafe.Pointer(a.instance)
    return a
}

// ActionFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ActionFromUnsafePointer(ptr unsafe.Pointer) *TAction {
    a := new(TAction)
    a.instance = uintptr(ptr)
    a.ptr = ptr
    return a
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (a *TAction) Free() {
    if a.instance != 0 {
        Action_Free(a.instance)
        a.instance = 0
        a.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (a *TAction) Instance() uintptr {
    return a.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (a *TAction) UnsafeAddr() unsafe.Pointer {
    return a.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (a *TAction) IsValid() bool {
    return a.instance != 0
}

// TActionClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TActionClass() TClass {
    return Action_StaticClassType()
}

// Execute
func (a *TAction) Execute() bool {
    return Action_Execute(a.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (a *TAction) Update() bool {
    return Action_Update(a.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (a *TAction) HasParent() bool {
    return Action_HasParent(a.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (a *TAction) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Action_FindComponent(a.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (a *TAction) GetNamePath() string {
    return Action_GetNamePath(a.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (a *TAction) Assign(Source IObject) {
    Action_Assign(a.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (a *TAction) DisposeOf() {
    Action_DisposeOf(a.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (a *TAction) ClassType() TClass {
    return Action_ClassType(a.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (a *TAction) ClassName() string {
    return Action_ClassName(a.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (a *TAction) InstanceSize() int32 {
    return Action_InstanceSize(a.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (a *TAction) InheritsFrom(AClass TClass) bool {
    return Action_InheritsFrom(a.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (a *TAction) Equals(Obj IObject) bool {
    return Action_Equals(a.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (a *TAction) GetHashCode() int32 {
    return Action_GetHashCode(a.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (a *TAction) ToString() string {
    return Action_ToString(a.instance)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (a *TAction) Caption() string {
    return Action_GetCaption(a.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (a *TAction) SetCaption(value string) {
    Action_SetCaption(a.instance, value)
}

// Checked
func (a *TAction) Checked() bool {
    return Action_GetChecked(a.instance)
}

// SetChecked
func (a *TAction) SetChecked(value bool) {
    Action_SetChecked(a.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (a *TAction) Enabled() bool {
    return Action_GetEnabled(a.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (a *TAction) SetEnabled(value bool) {
    Action_SetEnabled(a.instance, value)
}

// GroupIndex
func (a *TAction) GroupIndex() int32 {
    return Action_GetGroupIndex(a.instance)
}

// SetGroupIndex
func (a *TAction) SetGroupIndex(value int32) {
    Action_SetGroupIndex(a.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (a *TAction) Hint() string {
    return Action_GetHint(a.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (a *TAction) SetHint(value string) {
    Action_SetHint(a.instance, value)
}

// ImageIndex
func (a *TAction) ImageIndex() int32 {
    return Action_GetImageIndex(a.instance)
}

// SetImageIndex
func (a *TAction) SetImageIndex(value int32) {
    Action_SetImageIndex(a.instance, value)
}

// ShortCut
func (a *TAction) ShortCut() TShortCut {
    return Action_GetShortCut(a.instance)
}

// SetShortCut
func (a *TAction) SetShortCut(value TShortCut) {
    Action_SetShortCut(a.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (a *TAction) Visible() bool {
    return Action_GetVisible(a.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (a *TAction) SetVisible(value bool) {
    Action_SetVisible(a.instance, value)
}

// SetOnExecute
func (a *TAction) SetOnExecute(fn TNotifyEvent) {
    Action_SetOnExecute(a.instance, fn)
}

// SetOnUpdate
func (a *TAction) SetOnUpdate(fn TNotifyEvent) {
    Action_SetOnUpdate(a.instance, fn)
}

// Images
func (a *TAction) Images() *TImageList {
    return ImageListFromInst(Action_GetImages(a.instance))
}

// Index
func (a *TAction) Index() int32 {
    return Action_GetIndex(a.instance)
}

// SetIndex
func (a *TAction) SetIndex(value int32) {
    Action_SetIndex(a.instance, value)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (a *TAction) ComponentCount() int32 {
    return Action_GetComponentCount(a.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (a *TAction) ComponentIndex() int32 {
    return Action_GetComponentIndex(a.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (a *TAction) SetComponentIndex(value int32) {
    Action_SetComponentIndex(a.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (a *TAction) Owner() *TComponent {
    return ComponentFromInst(Action_GetOwner(a.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (a *TAction) Name() string {
    return Action_GetName(a.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (a *TAction) SetName(value string) {
    Action_SetName(a.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (a *TAction) Tag() int {
    return Action_GetTag(a.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (a *TAction) SetTag(value int) {
    Action_SetTag(a.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (a *TAction) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Action_GetComponents(a.instance, AIndex))
}

