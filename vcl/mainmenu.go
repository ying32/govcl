
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

type TMainMenu struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewMainMenu
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMainMenu(owner IComponent) *TMainMenu {
    m := new(TMainMenu)
    m.instance = MainMenu_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MainMenuFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MainMenuFromInst(inst uintptr) *TMainMenu {
    m := new(TMainMenu)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

// MainMenuFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MainMenuFromObj(obj IObject) *TMainMenu {
    m := new(TMainMenu)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MainMenuFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MainMenuFromUnsafePointer(ptr unsafe.Pointer) *TMainMenu {
    m := new(TMainMenu)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (m *TMainMenu) Free() {
    if m.instance != 0 {
        MainMenu_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMainMenu) Instance() uintptr {
    return m.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMainMenu) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMainMenu) IsValid() bool {
    return m.instance != 0
}

// TMainMenuClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMainMenuClass() TClass {
    return MainMenu_StaticClassType()
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (m *TMainMenu) FindComponent(AName string) *TComponent {
    return ComponentFromInst(MainMenu_FindComponent(m.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (m *TMainMenu) GetNamePath() string {
    return MainMenu_GetNamePath(m.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (m *TMainMenu) HasParent() bool {
    return MainMenu_HasParent(m.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (m *TMainMenu) Assign(Source IObject) {
    MainMenu_Assign(m.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMainMenu) DisposeOf() {
    MainMenu_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMainMenu) ClassType() TClass {
    return MainMenu_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMainMenu) ClassName() string {
    return MainMenu_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMainMenu) InstanceSize() int32 {
    return MainMenu_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMainMenu) InheritsFrom(AClass TClass) bool {
    return MainMenu_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMainMenu) Equals(Obj IObject) bool {
    return MainMenu_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMainMenu) GetHashCode() int32 {
    return MainMenu_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMainMenu) ToString() string {
    return MainMenu_ToString(m.instance)
}

// AutoHotkeys
func (m *TMainMenu) AutoHotkeys() TMenuAutoFlag {
    return MainMenu_GetAutoHotkeys(m.instance)
}

// SetAutoHotkeys
func (m *TMainMenu) SetAutoHotkeys(value TMenuAutoFlag) {
    MainMenu_SetAutoHotkeys(m.instance, value)
}

// BiDiMode
func (m *TMainMenu) BiDiMode() TBiDiMode {
    return MainMenu_GetBiDiMode(m.instance)
}

// SetBiDiMode
func (m *TMainMenu) SetBiDiMode(value TBiDiMode) {
    MainMenu_SetBiDiMode(m.instance, value)
}

// Images
func (m *TMainMenu) Images() *TImageList {
    return ImageListFromInst(MainMenu_GetImages(m.instance))
}

// SetImages
func (m *TMainMenu) SetImages(value IComponent) {
    MainMenu_SetImages(m.instance, CheckPtr(value))
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (m *TMainMenu) SetOnChange(fn TMenuChangeEvent) {
    MainMenu_SetOnChange(m.instance, fn)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (m *TMainMenu) Handle() HMENU {
    return MainMenu_GetHandle(m.instance)
}

// WindowHandle
func (m *TMainMenu) WindowHandle() HWND {
    return MainMenu_GetWindowHandle(m.instance)
}

// SetWindowHandle
func (m *TMainMenu) SetWindowHandle(value HWND) {
    MainMenu_SetWindowHandle(m.instance, value)
}

// Items
func (m *TMainMenu) Items() *TMenuItem {
    return MenuItemFromInst(MainMenu_GetItems(m.instance))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (m *TMainMenu) ComponentCount() int32 {
    return MainMenu_GetComponentCount(m.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (m *TMainMenu) ComponentIndex() int32 {
    return MainMenu_GetComponentIndex(m.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (m *TMainMenu) SetComponentIndex(value int32) {
    MainMenu_SetComponentIndex(m.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (m *TMainMenu) Owner() *TComponent {
    return ComponentFromInst(MainMenu_GetOwner(m.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (m *TMainMenu) Name() string {
    return MainMenu_GetName(m.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (m *TMainMenu) SetName(value string) {
    MainMenu_SetName(m.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (m *TMainMenu) Tag() int {
    return MainMenu_GetTag(m.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (m *TMainMenu) SetTag(value int) {
    MainMenu_SetTag(m.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (m *TMainMenu) Components(AIndex int32) *TComponent {
    return ComponentFromInst(MainMenu_GetComponents(m.instance, AIndex))
}

