
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

type TApplication struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewApplication
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewApplication(owner IComponent) *TApplication {
    a := new(TApplication)
    a.instance = Application_Create(CheckPtr(owner))
    a.ptr = unsafe.Pointer(a.instance)
    return a
}

// ApplicationFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ApplicationFromInst(inst uintptr) *TApplication {
    a := new(TApplication)
    a.instance = inst
    a.ptr = unsafe.Pointer(inst)
    return a
}

// ApplicationFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ApplicationFromObj(obj IObject) *TApplication {
    a := new(TApplication)
    a.instance = CheckPtr(obj)
    a.ptr = unsafe.Pointer(a.instance)
    return a
}

// ApplicationFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ApplicationFromUnsafePointer(ptr unsafe.Pointer) *TApplication {
    a := new(TApplication)
    a.instance = uintptr(ptr)
    a.ptr = ptr
    return a
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (a *TApplication) Free() {
    if a.instance != 0 {
        Application_Free(a.instance)
        a.instance = 0
        a.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (a *TApplication) Instance() uintptr {
    return a.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (a *TApplication) UnsafeAddr() unsafe.Pointer {
    return a.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (a *TApplication) IsValid() bool {
    return a.instance != 0
}

// TApplicationClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TApplicationClass() TClass {
    return Application_StaticClassType()
}

// ActivateHint
func (a *TApplication) ActivateHint(CursorPos TPoint) {
    Application_ActivateHint(a.instance, CursorPos)
}

// BringToFront
func (a *TApplication) BringToFront() {
    Application_BringToFront(a.instance)
}

// CancelHint
func (a *TApplication) CancelHint() {
    Application_CancelHint(a.instance)
}

// HandleMessage
func (a *TApplication) HandleMessage() {
    Application_HandleMessage(a.instance)
}

// HideHint
func (a *TApplication) HideHint() {
    Application_HideHint(a.instance)
}

// Minimize
func (a *TApplication) Minimize() {
    Application_Minimize(a.instance)
}

// ModalStarted
func (a *TApplication) ModalStarted() {
    Application_ModalStarted(a.instance)
}

// ModalFinished
func (a *TApplication) ModalFinished() {
    Application_ModalFinished(a.instance)
}

// NormalizeAllTopMosts
func (a *TApplication) NormalizeAllTopMosts() {
    Application_NormalizeAllTopMosts(a.instance)
}

// NormalizeTopMosts
func (a *TApplication) NormalizeTopMosts() {
    Application_NormalizeTopMosts(a.instance)
}

// ProcessMessages
func (a *TApplication) ProcessMessages() {
    Application_ProcessMessages(a.instance)
}

// Restore
func (a *TApplication) Restore() {
    Application_Restore(a.instance)
}

// RestoreTopMosts
func (a *TApplication) RestoreTopMosts() {
    Application_RestoreTopMosts(a.instance)
}

// Terminate
func (a *TApplication) Terminate() {
    Application_Terminate(a.instance)
}

// MessageBox
func (a *TApplication) MessageBox(Text string, Caption string, Flags int32) int32 {
    return Application_MessageBox(a.instance, Text , Caption , Flags)
}

// FindComponent
func (a *TApplication) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Application_FindComponent(a.instance, AName))
}

// GetNamePath
func (a *TApplication) GetNamePath() string {
    return Application_GetNamePath(a.instance)
}

// HasParent
func (a *TApplication) HasParent() bool {
    return Application_HasParent(a.instance)
}

// Assign
func (a *TApplication) Assign(Source IObject) {
    Application_Assign(a.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (a *TApplication) DisposeOf() {
    Application_DisposeOf(a.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (a *TApplication) ClassType() TClass {
    return Application_ClassType(a.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (a *TApplication) ClassName() string {
    return Application_ClassName(a.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (a *TApplication) InstanceSize() int32 {
    return Application_InstanceSize(a.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (a *TApplication) InheritsFrom(AClass TClass) bool {
    return Application_InheritsFrom(a.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (a *TApplication) Equals(Obj IObject) bool {
    return Application_Equals(a.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (a *TApplication) GetHashCode() int32 {
    return Application_GetHashCode(a.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (a *TApplication) ToString() string {
    return Application_ToString(a.instance)
}

// DialogHandle
func (a *TApplication) DialogHandle() HWND {
    return Application_GetDialogHandle(a.instance)
}

// SetDialogHandle
func (a *TApplication) SetDialogHandle(value HWND) {
    Application_SetDialogHandle(a.instance, value)
}

// ExeName
func (a *TApplication) ExeName() string {
    return Application_GetExeName(a.instance)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (a *TApplication) Hint() string {
    return Application_GetHint(a.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (a *TApplication) SetHint(value string) {
    Application_SetHint(a.instance, value)
}

// HintColor
func (a *TApplication) HintColor() TColor {
    return Application_GetHintColor(a.instance)
}

// SetHintColor
func (a *TApplication) SetHintColor(value TColor) {
    Application_SetHintColor(a.instance, value)
}

// HintHidePause
func (a *TApplication) HintHidePause() int32 {
    return Application_GetHintHidePause(a.instance)
}

// SetHintHidePause
func (a *TApplication) SetHintHidePause(value int32) {
    Application_SetHintHidePause(a.instance, value)
}

// HintPause
func (a *TApplication) HintPause() int32 {
    return Application_GetHintPause(a.instance)
}

// SetHintPause
func (a *TApplication) SetHintPause(value int32) {
    Application_SetHintPause(a.instance, value)
}

// HintShortCuts
func (a *TApplication) HintShortCuts() bool {
    return Application_GetHintShortCuts(a.instance)
}

// SetHintShortCuts
func (a *TApplication) SetHintShortCuts(value bool) {
    Application_SetHintShortCuts(a.instance, value)
}

// HintShortPause
func (a *TApplication) HintShortPause() int32 {
    return Application_GetHintShortPause(a.instance)
}

// SetHintShortPause
func (a *TApplication) SetHintShortPause(value int32) {
    Application_SetHintShortPause(a.instance, value)
}

// Icon
func (a *TApplication) Icon() *TIcon {
    return IconFromInst(Application_GetIcon(a.instance))
}

// SetIcon
func (a *TApplication) SetIcon(value *TIcon) {
    Application_SetIcon(a.instance, CheckPtr(value))
}

// IsMetropolisUI
func (a *TApplication) IsMetropolisUI() bool {
    return Application_GetIsMetropolisUI(a.instance)
}

// MainForm
func (a *TApplication) MainForm() *TForm {
    return FormFromInst(Application_GetMainForm(a.instance))
}

// MainFormHandle
func (a *TApplication) MainFormHandle() HWND {
    return Application_GetMainFormHandle(a.instance)
}

// MainFormOnTaskBar
func (a *TApplication) MainFormOnTaskBar() bool {
    return Application_GetMainFormOnTaskBar(a.instance)
}

// SetMainFormOnTaskBar
func (a *TApplication) SetMainFormOnTaskBar(value bool) {
    Application_SetMainFormOnTaskBar(a.instance, value)
}

// BiDiMode
func (a *TApplication) BiDiMode() TBiDiMode {
    return Application_GetBiDiMode(a.instance)
}

// SetBiDiMode
func (a *TApplication) SetBiDiMode(value TBiDiMode) {
    Application_SetBiDiMode(a.instance, value)
}

// ShowHint
func (a *TApplication) ShowHint() bool {
    return Application_GetShowHint(a.instance)
}

// SetShowHint
func (a *TApplication) SetShowHint(value bool) {
    Application_SetShowHint(a.instance, value)
}

// ShowMainForm
func (a *TApplication) ShowMainForm() bool {
    return Application_GetShowMainForm(a.instance)
}

// SetShowMainForm
func (a *TApplication) SetShowMainForm(value bool) {
    Application_SetShowMainForm(a.instance, value)
}

// Title
func (a *TApplication) Title() string {
    return Application_GetTitle(a.instance)
}

// SetTitle
func (a *TApplication) SetTitle(value string) {
    Application_SetTitle(a.instance, value)
}

// SetOnException
func (a *TApplication) SetOnException(fn TExceptionEvent) {
    Application_SetOnException(a.instance, fn)
}

// SetOnHelp
func (a *TApplication) SetOnHelp(fn THelpEvent) {
    Application_SetOnHelp(a.instance, fn)
}

// SetOnHint
func (a *TApplication) SetOnHint(fn TNotifyEvent) {
    Application_SetOnHint(a.instance, fn)
}

// SetOnMessage
func (a *TApplication) SetOnMessage(fn TMessageEvent) {
    Application_SetOnMessage(a.instance, fn)
}

// SetOnMinimize
func (a *TApplication) SetOnMinimize(fn TNotifyEvent) {
    Application_SetOnMinimize(a.instance, fn)
}

// SetOnRestore
func (a *TApplication) SetOnRestore(fn TNotifyEvent) {
    Application_SetOnRestore(a.instance, fn)
}

// SetOnShortCut
func (a *TApplication) SetOnShortCut(fn TShortCutEvent) {
    Application_SetOnShortCut(a.instance, fn)
}

// Handle
func (a *TApplication) Handle() HWND {
    return Application_GetHandle(a.instance)
}

// SetHandle
func (a *TApplication) SetHandle(value HWND) {
    Application_SetHandle(a.instance, value)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (a *TApplication) ComponentCount() int32 {
    return Application_GetComponentCount(a.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (a *TApplication) ComponentIndex() int32 {
    return Application_GetComponentIndex(a.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (a *TApplication) SetComponentIndex(value int32) {
    Application_SetComponentIndex(a.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (a *TApplication) Owner() *TComponent {
    return ComponentFromInst(Application_GetOwner(a.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (a *TApplication) Name() string {
    return Application_GetName(a.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (a *TApplication) SetName(value string) {
    Application_SetName(a.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (a *TApplication) Tag() int {
    return Application_GetTag(a.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (a *TApplication) SetTag(value int) {
    Application_SetTag(a.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (a *TApplication) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Application_GetComponents(a.instance, AIndex))
}

