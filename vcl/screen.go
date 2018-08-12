
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

type TScreen struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewScreen
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewScreen(owner IComponent) *TScreen {
    s := new(TScreen)
    s.instance = Screen_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// ScreenFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ScreenFromInst(inst uintptr) *TScreen {
    s := new(TScreen)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// ScreenFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ScreenFromObj(obj IObject) *TScreen {
    s := new(TScreen)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// ScreenFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ScreenFromUnsafePointer(ptr unsafe.Pointer) *TScreen {
    s := new(TScreen)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TScreen) Free() {
    if s.instance != 0 {
        Screen_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TScreen) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TScreen) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TScreen) IsValid() bool {
    return s.instance != 0
}

// TScreenClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TScreenClass() TClass {
    return Screen_StaticClassType()
}

// DisableAlign
func (s *TScreen) DisableAlign() {
    Screen_DisableAlign(s.instance)
}

// EnableAlign
func (s *TScreen) EnableAlign() {
    Screen_EnableAlign(s.instance)
}

// Realign
func (s *TScreen) Realign() {
    Screen_Realign(s.instance)
}

// FindComponent
func (s *TScreen) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Screen_FindComponent(s.instance, AName))
}

// GetNamePath
func (s *TScreen) GetNamePath() string {
    return Screen_GetNamePath(s.instance)
}

// HasParent
func (s *TScreen) HasParent() bool {
    return Screen_HasParent(s.instance)
}

// Assign
func (s *TScreen) Assign(Source IObject) {
    Screen_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TScreen) DisposeOf() {
    Screen_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TScreen) ClassType() TClass {
    return Screen_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TScreen) ClassName() string {
    return Screen_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TScreen) InstanceSize() int32 {
    return Screen_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TScreen) InheritsFrom(AClass TClass) bool {
    return Screen_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TScreen) Equals(Obj IObject) bool {
    return Screen_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TScreen) GetHashCode() int32 {
    return Screen_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TScreen) ToString() string {
    return Screen_ToString(s.instance)
}

// ActiveForm
func (s *TScreen) ActiveForm() *TForm {
    return FormFromInst(Screen_GetActiveForm(s.instance))
}

// CustomFormCount
func (s *TScreen) CustomFormCount() int32 {
    return Screen_GetCustomFormCount(s.instance)
}

// CursorCount
func (s *TScreen) CursorCount() int32 {
    return Screen_GetCursorCount(s.instance)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TScreen) Cursor() TCursor {
    return Screen_GetCursor(s.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TScreen) SetCursor(value TCursor) {
    Screen_SetCursor(s.instance, value)
}

// FocusedForm
func (s *TScreen) FocusedForm() *TForm {
    return FormFromInst(Screen_GetFocusedForm(s.instance))
}

// SetFocusedForm
func (s *TScreen) SetFocusedForm(value IWinControl) {
    Screen_SetFocusedForm(s.instance, CheckPtr(value))
}

// MonitorCount
func (s *TScreen) MonitorCount() int32 {
    return Screen_GetMonitorCount(s.instance)
}

// DesktopRect
func (s *TScreen) DesktopRect() TRect {
    return Screen_GetDesktopRect(s.instance)
}

// DesktopHeight
func (s *TScreen) DesktopHeight() int32 {
    return Screen_GetDesktopHeight(s.instance)
}

// DesktopLeft
func (s *TScreen) DesktopLeft() int32 {
    return Screen_GetDesktopLeft(s.instance)
}

// DesktopTop
func (s *TScreen) DesktopTop() int32 {
    return Screen_GetDesktopTop(s.instance)
}

// DesktopWidth
func (s *TScreen) DesktopWidth() int32 {
    return Screen_GetDesktopWidth(s.instance)
}

// WorkAreaRect
func (s *TScreen) WorkAreaRect() TRect {
    return Screen_GetWorkAreaRect(s.instance)
}

// WorkAreaHeight
func (s *TScreen) WorkAreaHeight() int32 {
    return Screen_GetWorkAreaHeight(s.instance)
}

// WorkAreaLeft
func (s *TScreen) WorkAreaLeft() int32 {
    return Screen_GetWorkAreaLeft(s.instance)
}

// WorkAreaTop
func (s *TScreen) WorkAreaTop() int32 {
    return Screen_GetWorkAreaTop(s.instance)
}

// WorkAreaWidth
func (s *TScreen) WorkAreaWidth() int32 {
    return Screen_GetWorkAreaWidth(s.instance)
}

// Fonts
func (s *TScreen) Fonts() *TStrings {
    return StringsFromInst(Screen_GetFonts(s.instance))
}

// FormCount
func (s *TScreen) FormCount() int32 {
    return Screen_GetFormCount(s.instance)
}

// Imes
func (s *TScreen) Imes() *TStrings {
    return StringsFromInst(Screen_GetImes(s.instance))
}

// DefaultIme
func (s *TScreen) DefaultIme() string {
    return Screen_GetDefaultIme(s.instance)
}

// Height
func (s *TScreen) Height() int32 {
    return Screen_GetHeight(s.instance)
}

// PixelsPerInch
func (s *TScreen) PixelsPerInch() int32 {
    return Screen_GetPixelsPerInch(s.instance)
}

// PrimaryMonitor
func (s *TScreen) PrimaryMonitor() *TMonitor {
    return MonitorFromInst(Screen_GetPrimaryMonitor(s.instance))
}

// Width
func (s *TScreen) Width() int32 {
    return Screen_GetWidth(s.instance)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TScreen) ComponentCount() int32 {
    return Screen_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TScreen) ComponentIndex() int32 {
    return Screen_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TScreen) SetComponentIndex(value int32) {
    Screen_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TScreen) Owner() *TComponent {
    return ComponentFromInst(Screen_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TScreen) Name() string {
    return Screen_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TScreen) SetName(value string) {
    Screen_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TScreen) Tag() int {
    return Screen_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TScreen) SetTag(value int) {
    Screen_SetTag(s.instance, value)
}

// Cursors
func (s *TScreen) Cursors(Index int32) HICON {
    return Screen_GetCursors(s.instance, Index)
}

// Cursors
func (s *TScreen) SetCursors(Index int32, value HICON) {
    Screen_SetCursors(s.instance, Index, value)
}

// Monitors
func (s *TScreen) Monitors(Index int32) *TMonitor {
    return MonitorFromInst(Screen_GetMonitors(s.instance, Index))
}

// Forms
func (s *TScreen) Forms(Index int32) *TForm {
    return FormFromInst(Screen_GetForms(s.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TScreen) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Screen_GetComponents(s.instance, AIndex))
}

