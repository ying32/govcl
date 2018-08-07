
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

type TBoundLabel struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewBoundLabel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBoundLabel(owner IComponent) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = BoundLabel_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BoundLabelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func BoundLabelFromInst(inst uintptr) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = inst
    b.ptr = unsafe.Pointer(inst)
    return b
}

// BoundLabelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func BoundLabelFromObj(obj IObject) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = CheckPtr(obj)
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BoundLabelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func BoundLabelFromUnsafePointer(ptr unsafe.Pointer) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = uintptr(ptr)
    b.ptr = ptr
    return b
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (b *TBoundLabel) Free() {
    if b.instance != 0 {
        BoundLabel_Free(b.instance)
        b.instance = 0
        b.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBoundLabel) Instance() uintptr {
    return b.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBoundLabel) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBoundLabel) IsValid() bool {
    return b.instance != 0
}

// TBoundLabelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBoundLabelClass() TClass {
    return BoundLabel_StaticClassType()
}

// BringToFront
func (b *TBoundLabel) BringToFront() {
    BoundLabel_BringToFront(b.instance)
}

// ClientToScreen
func (b *TBoundLabel) ClientToScreen(Point TPoint) TPoint {
    return BoundLabel_ClientToScreen(b.instance, Point)
}

// ClientToParent
func (b *TBoundLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return BoundLabel_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// Dragging
func (b *TBoundLabel) Dragging() bool {
    return BoundLabel_Dragging(b.instance)
}

// HasParent
func (b *TBoundLabel) HasParent() bool {
    return BoundLabel_HasParent(b.instance)
}

// Hide
func (b *TBoundLabel) Hide() {
    BoundLabel_Hide(b.instance)
}

// Invalidate
func (b *TBoundLabel) Invalidate() {
    BoundLabel_Invalidate(b.instance)
}

// Perform
func (b *TBoundLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return BoundLabel_Perform(b.instance, Msg , WParam , LParam)
}

// Refresh
func (b *TBoundLabel) Refresh() {
    BoundLabel_Refresh(b.instance)
}

// Repaint
func (b *TBoundLabel) Repaint() {
    BoundLabel_Repaint(b.instance)
}

// ScreenToClient
func (b *TBoundLabel) ScreenToClient(Point TPoint) TPoint {
    return BoundLabel_ScreenToClient(b.instance, Point)
}

// ParentToClient
func (b *TBoundLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return BoundLabel_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (b *TBoundLabel) SendToBack() {
    BoundLabel_SendToBack(b.instance)
}

// SetBounds
func (b *TBoundLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    BoundLabel_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
func (b *TBoundLabel) Show() {
    BoundLabel_Show(b.instance)
}

// Update
func (b *TBoundLabel) Update() {
    BoundLabel_Update(b.instance)
}

// GetTextBuf
func (b *TBoundLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return BoundLabel_GetTextBuf(b.instance, Buffer , BufSize)
}

// GetTextLen
func (b *TBoundLabel) GetTextLen() int32 {
    return BoundLabel_GetTextLen(b.instance)
}

// FindComponent
func (b *TBoundLabel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(BoundLabel_FindComponent(b.instance, AName))
}

// GetNamePath
func (b *TBoundLabel) GetNamePath() string {
    return BoundLabel_GetNamePath(b.instance)
}

// Assign
func (b *TBoundLabel) Assign(Source IObject) {
    BoundLabel_Assign(b.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TBoundLabel) DisposeOf() {
    BoundLabel_DisposeOf(b.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBoundLabel) ClassType() TClass {
    return BoundLabel_ClassType(b.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBoundLabel) ClassName() string {
    return BoundLabel_ClassName(b.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBoundLabel) InstanceSize() int32 {
    return BoundLabel_InstanceSize(b.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBoundLabel) InheritsFrom(AClass TClass) bool {
    return BoundLabel_InheritsFrom(b.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBoundLabel) Equals(Obj IObject) bool {
    return BoundLabel_Equals(b.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBoundLabel) GetHashCode() int32 {
    return BoundLabel_GetHashCode(b.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (b *TBoundLabel) ToString() string {
    return BoundLabel_ToString(b.instance)
}

// BiDiMode
func (b *TBoundLabel) BiDiMode() TBiDiMode {
    return BoundLabel_GetBiDiMode(b.instance)
}

// SetBiDiMode
func (b *TBoundLabel) SetBiDiMode(value TBiDiMode) {
    BoundLabel_SetBiDiMode(b.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (b *TBoundLabel) Caption() string {
    return BoundLabel_GetCaption(b.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (b *TBoundLabel) SetCaption(value string) {
    BoundLabel_SetCaption(b.instance, value)
}

// Color
func (b *TBoundLabel) Color() TColor {
    return BoundLabel_GetColor(b.instance)
}

// SetColor
func (b *TBoundLabel) SetColor(value TColor) {
    BoundLabel_SetColor(b.instance, value)
}

// DragCursor
func (b *TBoundLabel) DragCursor() TCursor {
    return BoundLabel_GetDragCursor(b.instance)
}

// SetDragCursor
func (b *TBoundLabel) SetDragCursor(value TCursor) {
    BoundLabel_SetDragCursor(b.instance, value)
}

// DragKind
func (b *TBoundLabel) DragKind() TDragKind {
    return BoundLabel_GetDragKind(b.instance)
}

// SetDragKind
func (b *TBoundLabel) SetDragKind(value TDragKind) {
    BoundLabel_SetDragKind(b.instance, value)
}

// DragMode
func (b *TBoundLabel) DragMode() TDragMode {
    return BoundLabel_GetDragMode(b.instance)
}

// SetDragMode
func (b *TBoundLabel) SetDragMode(value TDragMode) {
    BoundLabel_SetDragMode(b.instance, value)
}

// Font
func (b *TBoundLabel) Font() *TFont {
    return FontFromInst(BoundLabel_GetFont(b.instance))
}

// SetFont
func (b *TBoundLabel) SetFont(value *TFont) {
    BoundLabel_SetFont(b.instance, CheckPtr(value))
}

// Height
func (b *TBoundLabel) Height() int32 {
    return BoundLabel_GetHeight(b.instance)
}

// SetHeight
func (b *TBoundLabel) SetHeight(value int32) {
    BoundLabel_SetHeight(b.instance, value)
}

// Left
func (b *TBoundLabel) Left() int32 {
    return BoundLabel_GetLeft(b.instance)
}

// ParentColor
func (b *TBoundLabel) ParentColor() bool {
    return BoundLabel_GetParentColor(b.instance)
}

// SetParentColor
func (b *TBoundLabel) SetParentColor(value bool) {
    BoundLabel_SetParentColor(b.instance, value)
}

// ParentFont
func (b *TBoundLabel) ParentFont() bool {
    return BoundLabel_GetParentFont(b.instance)
}

// SetParentFont
func (b *TBoundLabel) SetParentFont(value bool) {
    BoundLabel_SetParentFont(b.instance, value)
}

// ParentShowHint
func (b *TBoundLabel) ParentShowHint() bool {
    return BoundLabel_GetParentShowHint(b.instance)
}

// SetParentShowHint
func (b *TBoundLabel) SetParentShowHint(value bool) {
    BoundLabel_SetParentShowHint(b.instance, value)
}

// PopupMenu
func (b *TBoundLabel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(BoundLabel_GetPopupMenu(b.instance))
}

// SetPopupMenu
func (b *TBoundLabel) SetPopupMenu(value IComponent) {
    BoundLabel_SetPopupMenu(b.instance, CheckPtr(value))
}

// ShowAccelChar
func (b *TBoundLabel) ShowAccelChar() bool {
    return BoundLabel_GetShowAccelChar(b.instance)
}

// SetShowAccelChar
func (b *TBoundLabel) SetShowAccelChar(value bool) {
    BoundLabel_SetShowAccelChar(b.instance, value)
}

// ShowHint
func (b *TBoundLabel) ShowHint() bool {
    return BoundLabel_GetShowHint(b.instance)
}

// SetShowHint
func (b *TBoundLabel) SetShowHint(value bool) {
    BoundLabel_SetShowHint(b.instance, value)
}

// Top
func (b *TBoundLabel) Top() int32 {
    return BoundLabel_GetTop(b.instance)
}

// Transparent
func (b *TBoundLabel) Transparent() bool {
    return BoundLabel_GetTransparent(b.instance)
}

// SetTransparent
func (b *TBoundLabel) SetTransparent(value bool) {
    BoundLabel_SetTransparent(b.instance, value)
}

// Layout
func (b *TBoundLabel) Layout() TTextLayout {
    return BoundLabel_GetLayout(b.instance)
}

// SetLayout
func (b *TBoundLabel) SetLayout(value TTextLayout) {
    BoundLabel_SetLayout(b.instance, value)
}

// WordWrap
func (b *TBoundLabel) WordWrap() bool {
    return BoundLabel_GetWordWrap(b.instance)
}

// SetWordWrap
func (b *TBoundLabel) SetWordWrap(value bool) {
    BoundLabel_SetWordWrap(b.instance, value)
}

// Width
func (b *TBoundLabel) Width() int32 {
    return BoundLabel_GetWidth(b.instance)
}

// SetWidth
func (b *TBoundLabel) SetWidth(value int32) {
    BoundLabel_SetWidth(b.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (b *TBoundLabel) SetOnClick(fn TNotifyEvent) {
    BoundLabel_SetOnClick(b.instance, fn)
}

// SetOnContextPopup
func (b *TBoundLabel) SetOnContextPopup(fn TContextPopupEvent) {
    BoundLabel_SetOnContextPopup(b.instance, fn)
}

// SetOnDblClick
func (b *TBoundLabel) SetOnDblClick(fn TNotifyEvent) {
    BoundLabel_SetOnDblClick(b.instance, fn)
}

// SetOnDragDrop
func (b *TBoundLabel) SetOnDragDrop(fn TDragDropEvent) {
    BoundLabel_SetOnDragDrop(b.instance, fn)
}

// SetOnDragOver
func (b *TBoundLabel) SetOnDragOver(fn TDragOverEvent) {
    BoundLabel_SetOnDragOver(b.instance, fn)
}

// SetOnEndDock
func (b *TBoundLabel) SetOnEndDock(fn TEndDragEvent) {
    BoundLabel_SetOnEndDock(b.instance, fn)
}

// SetOnEndDrag
func (b *TBoundLabel) SetOnEndDrag(fn TEndDragEvent) {
    BoundLabel_SetOnEndDrag(b.instance, fn)
}

// SetOnMouseDown
func (b *TBoundLabel) SetOnMouseDown(fn TMouseEvent) {
    BoundLabel_SetOnMouseDown(b.instance, fn)
}

// SetOnMouseMove
func (b *TBoundLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    BoundLabel_SetOnMouseMove(b.instance, fn)
}

// SetOnMouseUp
func (b *TBoundLabel) SetOnMouseUp(fn TMouseEvent) {
    BoundLabel_SetOnMouseUp(b.instance, fn)
}

// SetOnStartDock
func (b *TBoundLabel) SetOnStartDock(fn TStartDockEvent) {
    BoundLabel_SetOnStartDock(b.instance, fn)
}

// Canvas
func (b *TBoundLabel) Canvas() *TCanvas {
    return CanvasFromInst(BoundLabel_GetCanvas(b.instance))
}

// GlowSize
func (b *TBoundLabel) GlowSize() int32 {
    return BoundLabel_GetGlowSize(b.instance)
}

// SetGlowSize
func (b *TBoundLabel) SetGlowSize(value int32) {
    BoundLabel_SetGlowSize(b.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TBoundLabel) Enabled() bool {
    return BoundLabel_GetEnabled(b.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TBoundLabel) SetEnabled(value bool) {
    BoundLabel_SetEnabled(b.instance, value)
}

// Action
func (b *TBoundLabel) Action() *TAction {
    return ActionFromInst(BoundLabel_GetAction(b.instance))
}

// SetAction
func (b *TBoundLabel) SetAction(value IComponent) {
    BoundLabel_SetAction(b.instance, CheckPtr(value))
}

// Align
func (b *TBoundLabel) Align() TAlign {
    return BoundLabel_GetAlign(b.instance)
}

// SetAlign
func (b *TBoundLabel) SetAlign(value TAlign) {
    BoundLabel_SetAlign(b.instance, value)
}

// Anchors
func (b *TBoundLabel) Anchors() TAnchors {
    return BoundLabel_GetAnchors(b.instance)
}

// SetAnchors
func (b *TBoundLabel) SetAnchors(value TAnchors) {
    BoundLabel_SetAnchors(b.instance, value)
}

// BoundsRect
func (b *TBoundLabel) BoundsRect() TRect {
    return BoundLabel_GetBoundsRect(b.instance)
}

// SetBoundsRect
func (b *TBoundLabel) SetBoundsRect(value TRect) {
    BoundLabel_SetBoundsRect(b.instance, value)
}

// ClientHeight
func (b *TBoundLabel) ClientHeight() int32 {
    return BoundLabel_GetClientHeight(b.instance)
}

// SetClientHeight
func (b *TBoundLabel) SetClientHeight(value int32) {
    BoundLabel_SetClientHeight(b.instance, value)
}

// ClientRect
func (b *TBoundLabel) ClientRect() TRect {
    return BoundLabel_GetClientRect(b.instance)
}

// ClientWidth
func (b *TBoundLabel) ClientWidth() int32 {
    return BoundLabel_GetClientWidth(b.instance)
}

// SetClientWidth
func (b *TBoundLabel) SetClientWidth(value int32) {
    BoundLabel_SetClientWidth(b.instance, value)
}

// ExplicitLeft
func (b *TBoundLabel) ExplicitLeft() int32 {
    return BoundLabel_GetExplicitLeft(b.instance)
}

// ExplicitTop
func (b *TBoundLabel) ExplicitTop() int32 {
    return BoundLabel_GetExplicitTop(b.instance)
}

// ExplicitWidth
func (b *TBoundLabel) ExplicitWidth() int32 {
    return BoundLabel_GetExplicitWidth(b.instance)
}

// ExplicitHeight
func (b *TBoundLabel) ExplicitHeight() int32 {
    return BoundLabel_GetExplicitHeight(b.instance)
}

// Floating
func (b *TBoundLabel) Floating() bool {
    return BoundLabel_GetFloating(b.instance)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TBoundLabel) Visible() bool {
    return BoundLabel_GetVisible(b.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TBoundLabel) SetVisible(value bool) {
    BoundLabel_SetVisible(b.instance, value)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TBoundLabel) Parent() *TWinControl {
    return WinControlFromInst(BoundLabel_GetParent(b.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TBoundLabel) SetParent(value IWinControl) {
    BoundLabel_SetParent(b.instance, CheckPtr(value))
}

// StyleElements
func (b *TBoundLabel) StyleElements() TStyleElements {
    return BoundLabel_GetStyleElements(b.instance)
}

// SetStyleElements
func (b *TBoundLabel) SetStyleElements(value TStyleElements) {
    BoundLabel_SetStyleElements(b.instance, value)
}

// AlignWithMargins
func (b *TBoundLabel) AlignWithMargins() bool {
    return BoundLabel_GetAlignWithMargins(b.instance)
}

// SetAlignWithMargins
func (b *TBoundLabel) SetAlignWithMargins(value bool) {
    BoundLabel_SetAlignWithMargins(b.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TBoundLabel) Cursor() TCursor {
    return BoundLabel_GetCursor(b.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TBoundLabel) SetCursor(value TCursor) {
    BoundLabel_SetCursor(b.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (b *TBoundLabel) Hint() string {
    return BoundLabel_GetHint(b.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (b *TBoundLabel) SetHint(value string) {
    BoundLabel_SetHint(b.instance, value)
}

// Margins
func (b *TBoundLabel) Margins() *TMargins {
    return MarginsFromInst(BoundLabel_GetMargins(b.instance))
}

// SetMargins
func (b *TBoundLabel) SetMargins(value *TMargins) {
    BoundLabel_SetMargins(b.instance, CheckPtr(value))
}

// CustomHint
func (b *TBoundLabel) CustomHint() *TCustomHint {
    return CustomHintFromInst(BoundLabel_GetCustomHint(b.instance))
}

// SetCustomHint
func (b *TBoundLabel) SetCustomHint(value IComponent) {
    BoundLabel_SetCustomHint(b.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TBoundLabel) ComponentCount() int32 {
    return BoundLabel_GetComponentCount(b.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (b *TBoundLabel) ComponentIndex() int32 {
    return BoundLabel_GetComponentIndex(b.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (b *TBoundLabel) SetComponentIndex(value int32) {
    BoundLabel_SetComponentIndex(b.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TBoundLabel) Owner() *TComponent {
    return ComponentFromInst(BoundLabel_GetOwner(b.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (b *TBoundLabel) Name() string {
    return BoundLabel_GetName(b.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (b *TBoundLabel) SetName(value string) {
    BoundLabel_SetName(b.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TBoundLabel) Tag() int {
    return BoundLabel_GetTag(b.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TBoundLabel) SetTag(value int) {
    BoundLabel_SetTag(b.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TBoundLabel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(BoundLabel_GetComponents(b.instance, AIndex))
}

