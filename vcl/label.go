
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

type TLabel struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewLabel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewLabel(owner IComponent) *TLabel {
    l := new(TLabel)
    l.instance = Label_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// LabelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func LabelFromInst(inst uintptr) *TLabel {
    l := new(TLabel)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// LabelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func LabelFromObj(obj IObject) *TLabel {
    l := new(TLabel)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// LabelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func LabelFromUnsafePointer(ptr unsafe.Pointer) *TLabel {
    l := new(TLabel)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TLabel) Free() {
    if l.instance != 0 {
        Label_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TLabel) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TLabel) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TLabel) IsValid() bool {
    return l.instance != 0
}

// TLabelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TLabelClass() TClass {
    return Label_StaticClassType()
}

// BringToFront
func (l *TLabel) BringToFront() {
    Label_BringToFront(l.instance)
}

// ClientToScreen
func (l *TLabel) ClientToScreen(Point TPoint) TPoint {
    return Label_ClientToScreen(l.instance, Point)
}

// ClientToParent
func (l *TLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Label_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// Dragging
func (l *TLabel) Dragging() bool {
    return Label_Dragging(l.instance)
}

// HasParent
func (l *TLabel) HasParent() bool {
    return Label_HasParent(l.instance)
}

// Hide
func (l *TLabel) Hide() {
    Label_Hide(l.instance)
}

// Invalidate
func (l *TLabel) Invalidate() {
    Label_Invalidate(l.instance)
}

// Perform
func (l *TLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Label_Perform(l.instance, Msg , WParam , LParam)
}

// Refresh
func (l *TLabel) Refresh() {
    Label_Refresh(l.instance)
}

// Repaint
func (l *TLabel) Repaint() {
    Label_Repaint(l.instance)
}

// ScreenToClient
func (l *TLabel) ScreenToClient(Point TPoint) TPoint {
    return Label_ScreenToClient(l.instance, Point)
}

// ParentToClient
func (l *TLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Label_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (l *TLabel) SendToBack() {
    Label_SendToBack(l.instance)
}

// SetBounds
func (l *TLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Label_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
func (l *TLabel) Show() {
    Label_Show(l.instance)
}

// Update
func (l *TLabel) Update() {
    Label_Update(l.instance)
}

// GetTextBuf
func (l *TLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Label_GetTextBuf(l.instance, Buffer , BufSize)
}

// GetTextLen
func (l *TLabel) GetTextLen() int32 {
    return Label_GetTextLen(l.instance)
}

// FindComponent
func (l *TLabel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Label_FindComponent(l.instance, AName))
}

// GetNamePath
func (l *TLabel) GetNamePath() string {
    return Label_GetNamePath(l.instance)
}

// Assign
func (l *TLabel) Assign(Source IObject) {
    Label_Assign(l.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TLabel) DisposeOf() {
    Label_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TLabel) ClassType() TClass {
    return Label_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TLabel) ClassName() string {
    return Label_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TLabel) InstanceSize() int32 {
    return Label_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TLabel) InheritsFrom(AClass TClass) bool {
    return Label_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TLabel) Equals(Obj IObject) bool {
    return Label_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TLabel) GetHashCode() int32 {
    return Label_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TLabel) ToString() string {
    return Label_ToString(l.instance)
}

// Align
func (l *TLabel) Align() TAlign {
    return Label_GetAlign(l.instance)
}

// SetAlign
func (l *TLabel) SetAlign(value TAlign) {
    Label_SetAlign(l.instance, value)
}

// Alignment
func (l *TLabel) Alignment() TAlignment {
    return Label_GetAlignment(l.instance)
}

// SetAlignment
func (l *TLabel) SetAlignment(value TAlignment) {
    Label_SetAlignment(l.instance, value)
}

// Anchors
func (l *TLabel) Anchors() TAnchors {
    return Label_GetAnchors(l.instance)
}

// SetAnchors
func (l *TLabel) SetAnchors(value TAnchors) {
    Label_SetAnchors(l.instance, value)
}

// AutoSize
func (l *TLabel) AutoSize() bool {
    return Label_GetAutoSize(l.instance)
}

// SetAutoSize
func (l *TLabel) SetAutoSize(value bool) {
    Label_SetAutoSize(l.instance, value)
}

// BiDiMode
func (l *TLabel) BiDiMode() TBiDiMode {
    return Label_GetBiDiMode(l.instance)
}

// SetBiDiMode
func (l *TLabel) SetBiDiMode(value TBiDiMode) {
    Label_SetBiDiMode(l.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (l *TLabel) Caption() string {
    return Label_GetCaption(l.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (l *TLabel) SetCaption(value string) {
    Label_SetCaption(l.instance, value)
}

// Color
func (l *TLabel) Color() TColor {
    return Label_GetColor(l.instance)
}

// SetColor
func (l *TLabel) SetColor(value TColor) {
    Label_SetColor(l.instance, value)
}

// DragCursor
func (l *TLabel) DragCursor() TCursor {
    return Label_GetDragCursor(l.instance)
}

// SetDragCursor
func (l *TLabel) SetDragCursor(value TCursor) {
    Label_SetDragCursor(l.instance, value)
}

// DragKind
func (l *TLabel) DragKind() TDragKind {
    return Label_GetDragKind(l.instance)
}

// SetDragKind
func (l *TLabel) SetDragKind(value TDragKind) {
    Label_SetDragKind(l.instance, value)
}

// DragMode
func (l *TLabel) DragMode() TDragMode {
    return Label_GetDragMode(l.instance)
}

// SetDragMode
func (l *TLabel) SetDragMode(value TDragMode) {
    Label_SetDragMode(l.instance, value)
}

// EllipsisPosition
func (l *TLabel) EllipsisPosition() TEllipsisPosition {
    return Label_GetEllipsisPosition(l.instance)
}

// SetEllipsisPosition
func (l *TLabel) SetEllipsisPosition(value TEllipsisPosition) {
    Label_SetEllipsisPosition(l.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (l *TLabel) Enabled() bool {
    return Label_GetEnabled(l.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (l *TLabel) SetEnabled(value bool) {
    Label_SetEnabled(l.instance, value)
}

// Font
func (l *TLabel) Font() *TFont {
    return FontFromInst(Label_GetFont(l.instance))
}

// SetFont
func (l *TLabel) SetFont(value *TFont) {
    Label_SetFont(l.instance, CheckPtr(value))
}

// GlowSize
func (l *TLabel) GlowSize() int32 {
    return Label_GetGlowSize(l.instance)
}

// SetGlowSize
func (l *TLabel) SetGlowSize(value int32) {
    Label_SetGlowSize(l.instance, value)
}

// ParentColor
func (l *TLabel) ParentColor() bool {
    return Label_GetParentColor(l.instance)
}

// SetParentColor
func (l *TLabel) SetParentColor(value bool) {
    Label_SetParentColor(l.instance, value)
}

// ParentFont
func (l *TLabel) ParentFont() bool {
    return Label_GetParentFont(l.instance)
}

// SetParentFont
func (l *TLabel) SetParentFont(value bool) {
    Label_SetParentFont(l.instance, value)
}

// ParentShowHint
func (l *TLabel) ParentShowHint() bool {
    return Label_GetParentShowHint(l.instance)
}

// SetParentShowHint
func (l *TLabel) SetParentShowHint(value bool) {
    Label_SetParentShowHint(l.instance, value)
}

// PopupMenu
func (l *TLabel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Label_GetPopupMenu(l.instance))
}

// SetPopupMenu
func (l *TLabel) SetPopupMenu(value IComponent) {
    Label_SetPopupMenu(l.instance, CheckPtr(value))
}

// ShowAccelChar
func (l *TLabel) ShowAccelChar() bool {
    return Label_GetShowAccelChar(l.instance)
}

// SetShowAccelChar
func (l *TLabel) SetShowAccelChar(value bool) {
    Label_SetShowAccelChar(l.instance, value)
}

// ShowHint
func (l *TLabel) ShowHint() bool {
    return Label_GetShowHint(l.instance)
}

// SetShowHint
func (l *TLabel) SetShowHint(value bool) {
    Label_SetShowHint(l.instance, value)
}

// Transparent
func (l *TLabel) Transparent() bool {
    return Label_GetTransparent(l.instance)
}

// SetTransparent
func (l *TLabel) SetTransparent(value bool) {
    Label_SetTransparent(l.instance, value)
}

// Layout
func (l *TLabel) Layout() TTextLayout {
    return Label_GetLayout(l.instance)
}

// SetLayout
func (l *TLabel) SetLayout(value TTextLayout) {
    Label_SetLayout(l.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (l *TLabel) Visible() bool {
    return Label_GetVisible(l.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (l *TLabel) SetVisible(value bool) {
    Label_SetVisible(l.instance, value)
}

// WordWrap
func (l *TLabel) WordWrap() bool {
    return Label_GetWordWrap(l.instance)
}

// SetWordWrap
func (l *TLabel) SetWordWrap(value bool) {
    Label_SetWordWrap(l.instance, value)
}

// StyleElements
func (l *TLabel) StyleElements() TStyleElements {
    return Label_GetStyleElements(l.instance)
}

// SetStyleElements
func (l *TLabel) SetStyleElements(value TStyleElements) {
    Label_SetStyleElements(l.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (l *TLabel) SetOnClick(fn TNotifyEvent) {
    Label_SetOnClick(l.instance, fn)
}

// SetOnContextPopup
func (l *TLabel) SetOnContextPopup(fn TContextPopupEvent) {
    Label_SetOnContextPopup(l.instance, fn)
}

// SetOnDblClick
func (l *TLabel) SetOnDblClick(fn TNotifyEvent) {
    Label_SetOnDblClick(l.instance, fn)
}

// SetOnDragDrop
func (l *TLabel) SetOnDragDrop(fn TDragDropEvent) {
    Label_SetOnDragDrop(l.instance, fn)
}

// SetOnDragOver
func (l *TLabel) SetOnDragOver(fn TDragOverEvent) {
    Label_SetOnDragOver(l.instance, fn)
}

// SetOnEndDock
func (l *TLabel) SetOnEndDock(fn TEndDragEvent) {
    Label_SetOnEndDock(l.instance, fn)
}

// SetOnEndDrag
func (l *TLabel) SetOnEndDrag(fn TEndDragEvent) {
    Label_SetOnEndDrag(l.instance, fn)
}

// SetOnMouseDown
func (l *TLabel) SetOnMouseDown(fn TMouseEvent) {
    Label_SetOnMouseDown(l.instance, fn)
}

// SetOnMouseMove
func (l *TLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    Label_SetOnMouseMove(l.instance, fn)
}

// SetOnMouseUp
func (l *TLabel) SetOnMouseUp(fn TMouseEvent) {
    Label_SetOnMouseUp(l.instance, fn)
}

// SetOnMouseEnter
func (l *TLabel) SetOnMouseEnter(fn TNotifyEvent) {
    Label_SetOnMouseEnter(l.instance, fn)
}

// SetOnMouseLeave
func (l *TLabel) SetOnMouseLeave(fn TNotifyEvent) {
    Label_SetOnMouseLeave(l.instance, fn)
}

// SetOnStartDock
func (l *TLabel) SetOnStartDock(fn TStartDockEvent) {
    Label_SetOnStartDock(l.instance, fn)
}

// Canvas
func (l *TLabel) Canvas() *TCanvas {
    return CanvasFromInst(Label_GetCanvas(l.instance))
}

// Action
func (l *TLabel) Action() *TAction {
    return ActionFromInst(Label_GetAction(l.instance))
}

// SetAction
func (l *TLabel) SetAction(value IComponent) {
    Label_SetAction(l.instance, CheckPtr(value))
}

// BoundsRect
func (l *TLabel) BoundsRect() TRect {
    return Label_GetBoundsRect(l.instance)
}

// SetBoundsRect
func (l *TLabel) SetBoundsRect(value TRect) {
    Label_SetBoundsRect(l.instance, value)
}

// ClientHeight
func (l *TLabel) ClientHeight() int32 {
    return Label_GetClientHeight(l.instance)
}

// SetClientHeight
func (l *TLabel) SetClientHeight(value int32) {
    Label_SetClientHeight(l.instance, value)
}

// ClientRect
func (l *TLabel) ClientRect() TRect {
    return Label_GetClientRect(l.instance)
}

// ClientWidth
func (l *TLabel) ClientWidth() int32 {
    return Label_GetClientWidth(l.instance)
}

// SetClientWidth
func (l *TLabel) SetClientWidth(value int32) {
    Label_SetClientWidth(l.instance, value)
}

// ExplicitLeft
func (l *TLabel) ExplicitLeft() int32 {
    return Label_GetExplicitLeft(l.instance)
}

// ExplicitTop
func (l *TLabel) ExplicitTop() int32 {
    return Label_GetExplicitTop(l.instance)
}

// ExplicitWidth
func (l *TLabel) ExplicitWidth() int32 {
    return Label_GetExplicitWidth(l.instance)
}

// ExplicitHeight
func (l *TLabel) ExplicitHeight() int32 {
    return Label_GetExplicitHeight(l.instance)
}

// Floating
func (l *TLabel) Floating() bool {
    return Label_GetFloating(l.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (l *TLabel) Parent() *TWinControl {
    return WinControlFromInst(Label_GetParent(l.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (l *TLabel) SetParent(value IWinControl) {
    Label_SetParent(l.instance, CheckPtr(value))
}

// AlignWithMargins
func (l *TLabel) AlignWithMargins() bool {
    return Label_GetAlignWithMargins(l.instance)
}

// SetAlignWithMargins
func (l *TLabel) SetAlignWithMargins(value bool) {
    Label_SetAlignWithMargins(l.instance, value)
}

// Left
func (l *TLabel) Left() int32 {
    return Label_GetLeft(l.instance)
}

// SetLeft
func (l *TLabel) SetLeft(value int32) {
    Label_SetLeft(l.instance, value)
}

// Top
func (l *TLabel) Top() int32 {
    return Label_GetTop(l.instance)
}

// SetTop
func (l *TLabel) SetTop(value int32) {
    Label_SetTop(l.instance, value)
}

// Width
func (l *TLabel) Width() int32 {
    return Label_GetWidth(l.instance)
}

// SetWidth
func (l *TLabel) SetWidth(value int32) {
    Label_SetWidth(l.instance, value)
}

// Height
func (l *TLabel) Height() int32 {
    return Label_GetHeight(l.instance)
}

// SetHeight
func (l *TLabel) SetHeight(value int32) {
    Label_SetHeight(l.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (l *TLabel) Cursor() TCursor {
    return Label_GetCursor(l.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (l *TLabel) SetCursor(value TCursor) {
    Label_SetCursor(l.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (l *TLabel) Hint() string {
    return Label_GetHint(l.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (l *TLabel) SetHint(value string) {
    Label_SetHint(l.instance, value)
}

// Margins
func (l *TLabel) Margins() *TMargins {
    return MarginsFromInst(Label_GetMargins(l.instance))
}

// SetMargins
func (l *TLabel) SetMargins(value *TMargins) {
    Label_SetMargins(l.instance, CheckPtr(value))
}

// CustomHint
func (l *TLabel) CustomHint() *TCustomHint {
    return CustomHintFromInst(Label_GetCustomHint(l.instance))
}

// SetCustomHint
func (l *TLabel) SetCustomHint(value IComponent) {
    Label_SetCustomHint(l.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (l *TLabel) ComponentCount() int32 {
    return Label_GetComponentCount(l.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (l *TLabel) ComponentIndex() int32 {
    return Label_GetComponentIndex(l.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (l *TLabel) SetComponentIndex(value int32) {
    Label_SetComponentIndex(l.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TLabel) Owner() *TComponent {
    return ComponentFromInst(Label_GetOwner(l.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (l *TLabel) Name() string {
    return Label_GetName(l.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (l *TLabel) SetName(value string) {
    Label_SetName(l.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TLabel) Tag() int {
    return Label_GetTag(l.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TLabel) SetTag(value int) {
    Label_SetTag(l.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TLabel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Label_GetComponents(l.instance, AIndex))
}

