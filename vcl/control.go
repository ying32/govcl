
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

type TControl struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewControl
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewControl(owner IComponent) *TControl {
    c := new(TControl)
    c.instance = Control_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ControlFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ControlFromInst(inst uintptr) *TControl {
    c := new(TControl)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ControlFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ControlFromObj(obj IObject) *TControl {
    c := new(TControl)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ControlFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ControlFromUnsafePointer(ptr unsafe.Pointer) *TControl {
    c := new(TControl)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TControl) Free() {
    if c.instance != 0 {
        Control_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TControl) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TControl) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TControl) IsValid() bool {
    return c.instance != 0
}

// TControlClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TControlClass() TClass {
    return Control_StaticClassType()
}

// BringToFront
func (c *TControl) BringToFront() {
    Control_BringToFront(c.instance)
}

// ClientToScreen
func (c *TControl) ClientToScreen(Point TPoint) TPoint {
    return Control_ClientToScreen(c.instance, Point)
}

// ClientToParent
func (c *TControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Control_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
func (c *TControl) Dragging() bool {
    return Control_Dragging(c.instance)
}

// HasParent
func (c *TControl) HasParent() bool {
    return Control_HasParent(c.instance)
}

// Hide
func (c *TControl) Hide() {
    Control_Hide(c.instance)
}

// Invalidate
func (c *TControl) Invalidate() {
    Control_Invalidate(c.instance)
}

// Perform
func (c *TControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Control_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
func (c *TControl) Refresh() {
    Control_Refresh(c.instance)
}

// Repaint
func (c *TControl) Repaint() {
    Control_Repaint(c.instance)
}

// ScreenToClient
func (c *TControl) ScreenToClient(Point TPoint) TPoint {
    return Control_ScreenToClient(c.instance, Point)
}

// ParentToClient
func (c *TControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Control_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (c *TControl) SendToBack() {
    Control_SendToBack(c.instance)
}

// SetBounds
func (c *TControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Control_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
func (c *TControl) Show() {
    Control_Show(c.instance)
}

// Update
func (c *TControl) Update() {
    Control_Update(c.instance)
}

// GetTextBuf
func (c *TControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Control_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
func (c *TControl) GetTextLen() int32 {
    return Control_GetTextLen(c.instance)
}

// FindComponent
func (c *TControl) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Control_FindComponent(c.instance, AName))
}

// GetNamePath
func (c *TControl) GetNamePath() string {
    return Control_GetNamePath(c.instance)
}

// Assign
func (c *TControl) Assign(Source IObject) {
    Control_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TControl) DisposeOf() {
    Control_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TControl) ClassType() TClass {
    return Control_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TControl) ClassName() string {
    return Control_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TControl) InstanceSize() int32 {
    return Control_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TControl) InheritsFrom(AClass TClass) bool {
    return Control_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TControl) Equals(Obj IObject) bool {
    return Control_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TControl) GetHashCode() int32 {
    return Control_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TControl) ToString() string {
    return Control_ToString(c.instance)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TControl) Enabled() bool {
    return Control_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TControl) SetEnabled(value bool) {
    Control_SetEnabled(c.instance, value)
}

// Action
func (c *TControl) Action() *TAction {
    return ActionFromInst(Control_GetAction(c.instance))
}

// SetAction
func (c *TControl) SetAction(value IComponent) {
    Control_SetAction(c.instance, CheckPtr(value))
}

// Align
func (c *TControl) Align() TAlign {
    return Control_GetAlign(c.instance)
}

// SetAlign
func (c *TControl) SetAlign(value TAlign) {
    Control_SetAlign(c.instance, value)
}

// Anchors
func (c *TControl) Anchors() TAnchors {
    return Control_GetAnchors(c.instance)
}

// SetAnchors
func (c *TControl) SetAnchors(value TAnchors) {
    Control_SetAnchors(c.instance, value)
}

// BiDiMode
func (c *TControl) BiDiMode() TBiDiMode {
    return Control_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TControl) SetBiDiMode(value TBiDiMode) {
    Control_SetBiDiMode(c.instance, value)
}

// BoundsRect
func (c *TControl) BoundsRect() TRect {
    return Control_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TControl) SetBoundsRect(value TRect) {
    Control_SetBoundsRect(c.instance, value)
}

// ClientHeight
func (c *TControl) ClientHeight() int32 {
    return Control_GetClientHeight(c.instance)
}

// SetClientHeight
func (c *TControl) SetClientHeight(value int32) {
    Control_SetClientHeight(c.instance, value)
}

// ClientRect
func (c *TControl) ClientRect() TRect {
    return Control_GetClientRect(c.instance)
}

// ClientWidth
func (c *TControl) ClientWidth() int32 {
    return Control_GetClientWidth(c.instance)
}

// SetClientWidth
func (c *TControl) SetClientWidth(value int32) {
    Control_SetClientWidth(c.instance, value)
}

// ExplicitLeft
func (c *TControl) ExplicitLeft() int32 {
    return Control_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TControl) ExplicitTop() int32 {
    return Control_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TControl) ExplicitWidth() int32 {
    return Control_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TControl) ExplicitHeight() int32 {
    return Control_GetExplicitHeight(c.instance)
}

// Floating
func (c *TControl) Floating() bool {
    return Control_GetFloating(c.instance)
}

// ShowHint
func (c *TControl) ShowHint() bool {
    return Control_GetShowHint(c.instance)
}

// SetShowHint
func (c *TControl) SetShowHint(value bool) {
    Control_SetShowHint(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TControl) Visible() bool {
    return Control_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TControl) SetVisible(value bool) {
    Control_SetVisible(c.instance, value)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TControl) Parent() *TWinControl {
    return WinControlFromInst(Control_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TControl) SetParent(value IWinControl) {
    Control_SetParent(c.instance, CheckPtr(value))
}

// StyleElements
func (c *TControl) StyleElements() TStyleElements {
    return Control_GetStyleElements(c.instance)
}

// SetStyleElements
func (c *TControl) SetStyleElements(value TStyleElements) {
    Control_SetStyleElements(c.instance, value)
}

// AlignWithMargins
func (c *TControl) AlignWithMargins() bool {
    return Control_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
func (c *TControl) SetAlignWithMargins(value bool) {
    Control_SetAlignWithMargins(c.instance, value)
}

// Left
func (c *TControl) Left() int32 {
    return Control_GetLeft(c.instance)
}

// SetLeft
func (c *TControl) SetLeft(value int32) {
    Control_SetLeft(c.instance, value)
}

// Top
func (c *TControl) Top() int32 {
    return Control_GetTop(c.instance)
}

// SetTop
func (c *TControl) SetTop(value int32) {
    Control_SetTop(c.instance, value)
}

// Width
func (c *TControl) Width() int32 {
    return Control_GetWidth(c.instance)
}

// SetWidth
func (c *TControl) SetWidth(value int32) {
    Control_SetWidth(c.instance, value)
}

// Height
func (c *TControl) Height() int32 {
    return Control_GetHeight(c.instance)
}

// SetHeight
func (c *TControl) SetHeight(value int32) {
    Control_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TControl) Cursor() TCursor {
    return Control_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TControl) SetCursor(value TCursor) {
    Control_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (c *TControl) Hint() string {
    return Control_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (c *TControl) SetHint(value string) {
    Control_SetHint(c.instance, value)
}

// Margins
func (c *TControl) Margins() *TMargins {
    return MarginsFromInst(Control_GetMargins(c.instance))
}

// SetMargins
func (c *TControl) SetMargins(value *TMargins) {
    Control_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
func (c *TControl) CustomHint() *TCustomHint {
    return CustomHintFromInst(Control_GetCustomHint(c.instance))
}

// SetCustomHint
func (c *TControl) SetCustomHint(value IComponent) {
    Control_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TControl) ComponentCount() int32 {
    return Control_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TControl) ComponentIndex() int32 {
    return Control_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TControl) SetComponentIndex(value int32) {
    Control_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TControl) Owner() *TComponent {
    return ComponentFromInst(Control_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TControl) Name() string {
    return Control_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TControl) SetName(value string) {
    Control_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TControl) Tag() int {
    return Control_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TControl) SetTag(value int) {
    Control_SetTag(c.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TControl) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Control_GetComponents(c.instance, AIndex))
}

