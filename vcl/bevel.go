
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

type TBevel struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewBevel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBevel(owner IComponent) *TBevel {
    b := new(TBevel)
    b.instance = Bevel_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BevelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func BevelFromInst(inst uintptr) *TBevel {
    b := new(TBevel)
    b.instance = inst
    b.ptr = unsafe.Pointer(inst)
    return b
}

// BevelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func BevelFromObj(obj IObject) *TBevel {
    b := new(TBevel)
    b.instance = CheckPtr(obj)
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BevelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func BevelFromUnsafePointer(ptr unsafe.Pointer) *TBevel {
    b := new(TBevel)
    b.instance = uintptr(ptr)
    b.ptr = ptr
    return b
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (b *TBevel) Free() {
    if b.instance != 0 {
        Bevel_Free(b.instance)
        b.instance = 0
        b.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBevel) Instance() uintptr {
    return b.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBevel) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBevel) IsValid() bool {
    return b.instance != 0
}

// TBevelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBevelClass() TClass {
    return Bevel_StaticClassType()
}

// BringToFront
func (b *TBevel) BringToFront() {
    Bevel_BringToFront(b.instance)
}

// ClientToScreen
func (b *TBevel) ClientToScreen(Point TPoint) TPoint {
    return Bevel_ClientToScreen(b.instance, Point)
}

// ClientToParent
func (b *TBevel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Bevel_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// Dragging
func (b *TBevel) Dragging() bool {
    return Bevel_Dragging(b.instance)
}

// HasParent
func (b *TBevel) HasParent() bool {
    return Bevel_HasParent(b.instance)
}

// Hide
func (b *TBevel) Hide() {
    Bevel_Hide(b.instance)
}

// Invalidate
func (b *TBevel) Invalidate() {
    Bevel_Invalidate(b.instance)
}

// Perform
func (b *TBevel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Bevel_Perform(b.instance, Msg , WParam , LParam)
}

// Refresh
func (b *TBevel) Refresh() {
    Bevel_Refresh(b.instance)
}

// Repaint
func (b *TBevel) Repaint() {
    Bevel_Repaint(b.instance)
}

// ScreenToClient
func (b *TBevel) ScreenToClient(Point TPoint) TPoint {
    return Bevel_ScreenToClient(b.instance, Point)
}

// ParentToClient
func (b *TBevel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Bevel_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (b *TBevel) SendToBack() {
    Bevel_SendToBack(b.instance)
}

// SetBounds
func (b *TBevel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Bevel_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
func (b *TBevel) Show() {
    Bevel_Show(b.instance)
}

// Update
func (b *TBevel) Update() {
    Bevel_Update(b.instance)
}

// GetTextBuf
func (b *TBevel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Bevel_GetTextBuf(b.instance, Buffer , BufSize)
}

// GetTextLen
func (b *TBevel) GetTextLen() int32 {
    return Bevel_GetTextLen(b.instance)
}

// FindComponent
func (b *TBevel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Bevel_FindComponent(b.instance, AName))
}

// GetNamePath
func (b *TBevel) GetNamePath() string {
    return Bevel_GetNamePath(b.instance)
}

// Assign
func (b *TBevel) Assign(Source IObject) {
    Bevel_Assign(b.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TBevel) DisposeOf() {
    Bevel_DisposeOf(b.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBevel) ClassType() TClass {
    return Bevel_ClassType(b.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBevel) ClassName() string {
    return Bevel_ClassName(b.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBevel) InstanceSize() int32 {
    return Bevel_InstanceSize(b.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBevel) InheritsFrom(AClass TClass) bool {
    return Bevel_InheritsFrom(b.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBevel) Equals(Obj IObject) bool {
    return Bevel_Equals(b.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBevel) GetHashCode() int32 {
    return Bevel_GetHashCode(b.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (b *TBevel) ToString() string {
    return Bevel_ToString(b.instance)
}

// Align
func (b *TBevel) Align() TAlign {
    return Bevel_GetAlign(b.instance)
}

// SetAlign
func (b *TBevel) SetAlign(value TAlign) {
    Bevel_SetAlign(b.instance, value)
}

// Anchors
func (b *TBevel) Anchors() TAnchors {
    return Bevel_GetAnchors(b.instance)
}

// SetAnchors
func (b *TBevel) SetAnchors(value TAnchors) {
    Bevel_SetAnchors(b.instance, value)
}

// ParentShowHint
func (b *TBevel) ParentShowHint() bool {
    return Bevel_GetParentShowHint(b.instance)
}

// SetParentShowHint
func (b *TBevel) SetParentShowHint(value bool) {
    Bevel_SetParentShowHint(b.instance, value)
}

// Shape
func (b *TBevel) Shape() TBevelShape {
    return Bevel_GetShape(b.instance)
}

// SetShape
func (b *TBevel) SetShape(value TBevelShape) {
    Bevel_SetShape(b.instance, value)
}

// ShowHint
func (b *TBevel) ShowHint() bool {
    return Bevel_GetShowHint(b.instance)
}

// SetShowHint
func (b *TBevel) SetShowHint(value bool) {
    Bevel_SetShowHint(b.instance, value)
}

// Style
func (b *TBevel) Style() TBevelStyle {
    return Bevel_GetStyle(b.instance)
}

// SetStyle
func (b *TBevel) SetStyle(value TBevelStyle) {
    Bevel_SetStyle(b.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TBevel) Visible() bool {
    return Bevel_GetVisible(b.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TBevel) SetVisible(value bool) {
    Bevel_SetVisible(b.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TBevel) Enabled() bool {
    return Bevel_GetEnabled(b.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TBevel) SetEnabled(value bool) {
    Bevel_SetEnabled(b.instance, value)
}

// Action
func (b *TBevel) Action() *TAction {
    return ActionFromInst(Bevel_GetAction(b.instance))
}

// SetAction
func (b *TBevel) SetAction(value IComponent) {
    Bevel_SetAction(b.instance, CheckPtr(value))
}

// BiDiMode
func (b *TBevel) BiDiMode() TBiDiMode {
    return Bevel_GetBiDiMode(b.instance)
}

// SetBiDiMode
func (b *TBevel) SetBiDiMode(value TBiDiMode) {
    Bevel_SetBiDiMode(b.instance, value)
}

// BoundsRect
func (b *TBevel) BoundsRect() TRect {
    return Bevel_GetBoundsRect(b.instance)
}

// SetBoundsRect
func (b *TBevel) SetBoundsRect(value TRect) {
    Bevel_SetBoundsRect(b.instance, value)
}

// ClientHeight
func (b *TBevel) ClientHeight() int32 {
    return Bevel_GetClientHeight(b.instance)
}

// SetClientHeight
func (b *TBevel) SetClientHeight(value int32) {
    Bevel_SetClientHeight(b.instance, value)
}

// ClientRect
func (b *TBevel) ClientRect() TRect {
    return Bevel_GetClientRect(b.instance)
}

// ClientWidth
func (b *TBevel) ClientWidth() int32 {
    return Bevel_GetClientWidth(b.instance)
}

// SetClientWidth
func (b *TBevel) SetClientWidth(value int32) {
    Bevel_SetClientWidth(b.instance, value)
}

// ExplicitLeft
func (b *TBevel) ExplicitLeft() int32 {
    return Bevel_GetExplicitLeft(b.instance)
}

// ExplicitTop
func (b *TBevel) ExplicitTop() int32 {
    return Bevel_GetExplicitTop(b.instance)
}

// ExplicitWidth
func (b *TBevel) ExplicitWidth() int32 {
    return Bevel_GetExplicitWidth(b.instance)
}

// ExplicitHeight
func (b *TBevel) ExplicitHeight() int32 {
    return Bevel_GetExplicitHeight(b.instance)
}

// Floating
func (b *TBevel) Floating() bool {
    return Bevel_GetFloating(b.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TBevel) Parent() *TWinControl {
    return WinControlFromInst(Bevel_GetParent(b.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TBevel) SetParent(value IWinControl) {
    Bevel_SetParent(b.instance, CheckPtr(value))
}

// StyleElements
func (b *TBevel) StyleElements() TStyleElements {
    return Bevel_GetStyleElements(b.instance)
}

// SetStyleElements
func (b *TBevel) SetStyleElements(value TStyleElements) {
    Bevel_SetStyleElements(b.instance, value)
}

// AlignWithMargins
func (b *TBevel) AlignWithMargins() bool {
    return Bevel_GetAlignWithMargins(b.instance)
}

// SetAlignWithMargins
func (b *TBevel) SetAlignWithMargins(value bool) {
    Bevel_SetAlignWithMargins(b.instance, value)
}

// Left
func (b *TBevel) Left() int32 {
    return Bevel_GetLeft(b.instance)
}

// SetLeft
func (b *TBevel) SetLeft(value int32) {
    Bevel_SetLeft(b.instance, value)
}

// Top
func (b *TBevel) Top() int32 {
    return Bevel_GetTop(b.instance)
}

// SetTop
func (b *TBevel) SetTop(value int32) {
    Bevel_SetTop(b.instance, value)
}

// Width
func (b *TBevel) Width() int32 {
    return Bevel_GetWidth(b.instance)
}

// SetWidth
func (b *TBevel) SetWidth(value int32) {
    Bevel_SetWidth(b.instance, value)
}

// Height
func (b *TBevel) Height() int32 {
    return Bevel_GetHeight(b.instance)
}

// SetHeight
func (b *TBevel) SetHeight(value int32) {
    Bevel_SetHeight(b.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TBevel) Cursor() TCursor {
    return Bevel_GetCursor(b.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TBevel) SetCursor(value TCursor) {
    Bevel_SetCursor(b.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (b *TBevel) Hint() string {
    return Bevel_GetHint(b.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (b *TBevel) SetHint(value string) {
    Bevel_SetHint(b.instance, value)
}

// Margins
func (b *TBevel) Margins() *TMargins {
    return MarginsFromInst(Bevel_GetMargins(b.instance))
}

// SetMargins
func (b *TBevel) SetMargins(value *TMargins) {
    Bevel_SetMargins(b.instance, CheckPtr(value))
}

// CustomHint
func (b *TBevel) CustomHint() *TCustomHint {
    return CustomHintFromInst(Bevel_GetCustomHint(b.instance))
}

// SetCustomHint
func (b *TBevel) SetCustomHint(value IComponent) {
    Bevel_SetCustomHint(b.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TBevel) ComponentCount() int32 {
    return Bevel_GetComponentCount(b.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (b *TBevel) ComponentIndex() int32 {
    return Bevel_GetComponentIndex(b.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (b *TBevel) SetComponentIndex(value int32) {
    Bevel_SetComponentIndex(b.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TBevel) Owner() *TComponent {
    return ComponentFromInst(Bevel_GetOwner(b.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (b *TBevel) Name() string {
    return Bevel_GetName(b.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (b *TBevel) SetName(value string) {
    Bevel_SetName(b.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TBevel) Tag() int {
    return Bevel_GetTag(b.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TBevel) SetTag(value int) {
    Bevel_SetTag(b.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TBevel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Bevel_GetComponents(b.instance, AIndex))
}

