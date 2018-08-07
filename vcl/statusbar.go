
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

type TStatusBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewStatusBar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewStatusBar(owner IComponent) *TStatusBar {
    s := new(TStatusBar)
    s.instance = StatusBar_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StatusBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func StatusBarFromInst(inst uintptr) *TStatusBar {
    s := new(TStatusBar)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// StatusBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func StatusBarFromObj(obj IObject) *TStatusBar {
    s := new(TStatusBar)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StatusBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func StatusBarFromUnsafePointer(ptr unsafe.Pointer) *TStatusBar {
    s := new(TStatusBar)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TStatusBar) Free() {
    if s.instance != 0 {
        StatusBar_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TStatusBar) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TStatusBar) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TStatusBar) IsValid() bool {
    return s.instance != 0
}

// TStatusBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TStatusBarClass() TClass {
    return StatusBar_StaticClassType()
}

// FlipChildren
func (s *TStatusBar) FlipChildren(AllLevels bool) {
    StatusBar_FlipChildren(s.instance, AllLevels)
}

// SetBounds
func (s *TStatusBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    StatusBar_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// CanFocus
func (s *TStatusBar) CanFocus() bool {
    return StatusBar_CanFocus(s.instance)
}

// Focused
func (s *TStatusBar) Focused() bool {
    return StatusBar_Focused(s.instance)
}

// HandleAllocated
func (s *TStatusBar) HandleAllocated() bool {
    return StatusBar_HandleAllocated(s.instance)
}

// Invalidate
func (s *TStatusBar) Invalidate() {
    StatusBar_Invalidate(s.instance)
}

// Realign
func (s *TStatusBar) Realign() {
    StatusBar_Realign(s.instance)
}

// Repaint
func (s *TStatusBar) Repaint() {
    StatusBar_Repaint(s.instance)
}

// ScaleBy
func (s *TStatusBar) ScaleBy(M int32, D int32) {
    StatusBar_ScaleBy(s.instance, M , D)
}

// SetFocus
func (s *TStatusBar) SetFocus() {
    StatusBar_SetFocus(s.instance)
}

// Update
func (s *TStatusBar) Update() {
    StatusBar_Update(s.instance)
}

// BringToFront
func (s *TStatusBar) BringToFront() {
    StatusBar_BringToFront(s.instance)
}

// ClientToScreen
func (s *TStatusBar) ClientToScreen(Point TPoint) TPoint {
    return StatusBar_ClientToScreen(s.instance, Point)
}

// ClientToParent
func (s *TStatusBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return StatusBar_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// Dragging
func (s *TStatusBar) Dragging() bool {
    return StatusBar_Dragging(s.instance)
}

// HasParent
func (s *TStatusBar) HasParent() bool {
    return StatusBar_HasParent(s.instance)
}

// Hide
func (s *TStatusBar) Hide() {
    StatusBar_Hide(s.instance)
}

// Perform
func (s *TStatusBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return StatusBar_Perform(s.instance, Msg , WParam , LParam)
}

// Refresh
func (s *TStatusBar) Refresh() {
    StatusBar_Refresh(s.instance)
}

// ScreenToClient
func (s *TStatusBar) ScreenToClient(Point TPoint) TPoint {
    return StatusBar_ScreenToClient(s.instance, Point)
}

// ParentToClient
func (s *TStatusBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return StatusBar_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (s *TStatusBar) SendToBack() {
    StatusBar_SendToBack(s.instance)
}

// Show
func (s *TStatusBar) Show() {
    StatusBar_Show(s.instance)
}

// GetTextBuf
func (s *TStatusBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return StatusBar_GetTextBuf(s.instance, Buffer , BufSize)
}

// GetTextLen
func (s *TStatusBar) GetTextLen() int32 {
    return StatusBar_GetTextLen(s.instance)
}

// FindComponent
func (s *TStatusBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(StatusBar_FindComponent(s.instance, AName))
}

// GetNamePath
func (s *TStatusBar) GetNamePath() string {
    return StatusBar_GetNamePath(s.instance)
}

// Assign
func (s *TStatusBar) Assign(Source IObject) {
    StatusBar_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TStatusBar) DisposeOf() {
    StatusBar_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TStatusBar) ClassType() TClass {
    return StatusBar_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TStatusBar) ClassName() string {
    return StatusBar_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TStatusBar) InstanceSize() int32 {
    return StatusBar_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TStatusBar) InheritsFrom(AClass TClass) bool {
    return StatusBar_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TStatusBar) Equals(Obj IObject) bool {
    return StatusBar_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TStatusBar) GetHashCode() int32 {
    return StatusBar_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TStatusBar) ToString() string {
    return StatusBar_ToString(s.instance)
}

// Action
func (s *TStatusBar) Action() *TAction {
    return ActionFromInst(StatusBar_GetAction(s.instance))
}

// SetAction
func (s *TStatusBar) SetAction(value IComponent) {
    StatusBar_SetAction(s.instance, CheckPtr(value))
}

// AutoHint
func (s *TStatusBar) AutoHint() bool {
    return StatusBar_GetAutoHint(s.instance)
}

// SetAutoHint
func (s *TStatusBar) SetAutoHint(value bool) {
    StatusBar_SetAutoHint(s.instance, value)
}

// Align
func (s *TStatusBar) Align() TAlign {
    return StatusBar_GetAlign(s.instance)
}

// SetAlign
func (s *TStatusBar) SetAlign(value TAlign) {
    StatusBar_SetAlign(s.instance, value)
}

// Anchors
func (s *TStatusBar) Anchors() TAnchors {
    return StatusBar_GetAnchors(s.instance)
}

// SetAnchors
func (s *TStatusBar) SetAnchors(value TAnchors) {
    StatusBar_SetAnchors(s.instance, value)
}

// BiDiMode
func (s *TStatusBar) BiDiMode() TBiDiMode {
    return StatusBar_GetBiDiMode(s.instance)
}

// SetBiDiMode
func (s *TStatusBar) SetBiDiMode(value TBiDiMode) {
    StatusBar_SetBiDiMode(s.instance, value)
}

// BorderWidth
func (s *TStatusBar) BorderWidth() int32 {
    return StatusBar_GetBorderWidth(s.instance)
}

// SetBorderWidth
func (s *TStatusBar) SetBorderWidth(value int32) {
    StatusBar_SetBorderWidth(s.instance, value)
}

// Color
func (s *TStatusBar) Color() TColor {
    return StatusBar_GetColor(s.instance)
}

// SetColor
func (s *TStatusBar) SetColor(value TColor) {
    StatusBar_SetColor(s.instance, value)
}

// DoubleBuffered
func (s *TStatusBar) DoubleBuffered() bool {
    return StatusBar_GetDoubleBuffered(s.instance)
}

// SetDoubleBuffered
func (s *TStatusBar) SetDoubleBuffered(value bool) {
    StatusBar_SetDoubleBuffered(s.instance, value)
}

// DragCursor
func (s *TStatusBar) DragCursor() TCursor {
    return StatusBar_GetDragCursor(s.instance)
}

// SetDragCursor
func (s *TStatusBar) SetDragCursor(value TCursor) {
    StatusBar_SetDragCursor(s.instance, value)
}

// DragKind
func (s *TStatusBar) DragKind() TDragKind {
    return StatusBar_GetDragKind(s.instance)
}

// SetDragKind
func (s *TStatusBar) SetDragKind(value TDragKind) {
    StatusBar_SetDragKind(s.instance, value)
}

// DragMode
func (s *TStatusBar) DragMode() TDragMode {
    return StatusBar_GetDragMode(s.instance)
}

// SetDragMode
func (s *TStatusBar) SetDragMode(value TDragMode) {
    StatusBar_SetDragMode(s.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TStatusBar) Enabled() bool {
    return StatusBar_GetEnabled(s.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TStatusBar) SetEnabled(value bool) {
    StatusBar_SetEnabled(s.instance, value)
}

// Font
func (s *TStatusBar) Font() *TFont {
    return FontFromInst(StatusBar_GetFont(s.instance))
}

// SetFont
func (s *TStatusBar) SetFont(value *TFont) {
    StatusBar_SetFont(s.instance, CheckPtr(value))
}

// Panels
func (s *TStatusBar) Panels() *TStatusPanels {
    return StatusPanelsFromInst(StatusBar_GetPanels(s.instance))
}

// SetPanels
func (s *TStatusBar) SetPanels(value *TStatusPanels) {
    StatusBar_SetPanels(s.instance, CheckPtr(value))
}

// ParentColor
func (s *TStatusBar) ParentColor() bool {
    return StatusBar_GetParentColor(s.instance)
}

// SetParentColor
func (s *TStatusBar) SetParentColor(value bool) {
    StatusBar_SetParentColor(s.instance, value)
}

// ParentDoubleBuffered
func (s *TStatusBar) ParentDoubleBuffered() bool {
    return StatusBar_GetParentDoubleBuffered(s.instance)
}

// SetParentDoubleBuffered
func (s *TStatusBar) SetParentDoubleBuffered(value bool) {
    StatusBar_SetParentDoubleBuffered(s.instance, value)
}

// ParentFont
func (s *TStatusBar) ParentFont() bool {
    return StatusBar_GetParentFont(s.instance)
}

// SetParentFont
func (s *TStatusBar) SetParentFont(value bool) {
    StatusBar_SetParentFont(s.instance, value)
}

// ParentShowHint
func (s *TStatusBar) ParentShowHint() bool {
    return StatusBar_GetParentShowHint(s.instance)
}

// SetParentShowHint
func (s *TStatusBar) SetParentShowHint(value bool) {
    StatusBar_SetParentShowHint(s.instance, value)
}

// PopupMenu
func (s *TStatusBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(StatusBar_GetPopupMenu(s.instance))
}

// SetPopupMenu
func (s *TStatusBar) SetPopupMenu(value IComponent) {
    StatusBar_SetPopupMenu(s.instance, CheckPtr(value))
}

// ShowHint
func (s *TStatusBar) ShowHint() bool {
    return StatusBar_GetShowHint(s.instance)
}

// SetShowHint
func (s *TStatusBar) SetShowHint(value bool) {
    StatusBar_SetShowHint(s.instance, value)
}

// SimplePanel
func (s *TStatusBar) SimplePanel() bool {
    return StatusBar_GetSimplePanel(s.instance)
}

// SetSimplePanel
func (s *TStatusBar) SetSimplePanel(value bool) {
    StatusBar_SetSimplePanel(s.instance, value)
}

// SimpleText
func (s *TStatusBar) SimpleText() string {
    return StatusBar_GetSimpleText(s.instance)
}

// SetSimpleText
func (s *TStatusBar) SetSimpleText(value string) {
    StatusBar_SetSimpleText(s.instance, value)
}

// SizeGrip
func (s *TStatusBar) SizeGrip() bool {
    return StatusBar_GetSizeGrip(s.instance)
}

// SetSizeGrip
func (s *TStatusBar) SetSizeGrip(value bool) {
    StatusBar_SetSizeGrip(s.instance, value)
}

// UseSystemFont
func (s *TStatusBar) UseSystemFont() bool {
    return StatusBar_GetUseSystemFont(s.instance)
}

// SetUseSystemFont
func (s *TStatusBar) SetUseSystemFont(value bool) {
    StatusBar_SetUseSystemFont(s.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TStatusBar) Visible() bool {
    return StatusBar_GetVisible(s.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TStatusBar) SetVisible(value bool) {
    StatusBar_SetVisible(s.instance, value)
}

// StyleElements
func (s *TStatusBar) StyleElements() TStyleElements {
    return StatusBar_GetStyleElements(s.instance)
}

// SetStyleElements
func (s *TStatusBar) SetStyleElements(value TStyleElements) {
    StatusBar_SetStyleElements(s.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (s *TStatusBar) SetOnClick(fn TNotifyEvent) {
    StatusBar_SetOnClick(s.instance, fn)
}

// SetOnContextPopup
func (s *TStatusBar) SetOnContextPopup(fn TContextPopupEvent) {
    StatusBar_SetOnContextPopup(s.instance, fn)
}

// SetOnDblClick
func (s *TStatusBar) SetOnDblClick(fn TNotifyEvent) {
    StatusBar_SetOnDblClick(s.instance, fn)
}

// SetOnDragDrop
func (s *TStatusBar) SetOnDragDrop(fn TDragDropEvent) {
    StatusBar_SetOnDragDrop(s.instance, fn)
}

// SetOnDragOver
func (s *TStatusBar) SetOnDragOver(fn TDragOverEvent) {
    StatusBar_SetOnDragOver(s.instance, fn)
}

// SetOnEndDock
func (s *TStatusBar) SetOnEndDock(fn TEndDragEvent) {
    StatusBar_SetOnEndDock(s.instance, fn)
}

// SetOnEndDrag
func (s *TStatusBar) SetOnEndDrag(fn TEndDragEvent) {
    StatusBar_SetOnEndDrag(s.instance, fn)
}

// SetOnHint
func (s *TStatusBar) SetOnHint(fn TNotifyEvent) {
    StatusBar_SetOnHint(s.instance, fn)
}

// SetOnMouseDown
func (s *TStatusBar) SetOnMouseDown(fn TMouseEvent) {
    StatusBar_SetOnMouseDown(s.instance, fn)
}

// SetOnMouseEnter
func (s *TStatusBar) SetOnMouseEnter(fn TNotifyEvent) {
    StatusBar_SetOnMouseEnter(s.instance, fn)
}

// SetOnMouseLeave
func (s *TStatusBar) SetOnMouseLeave(fn TNotifyEvent) {
    StatusBar_SetOnMouseLeave(s.instance, fn)
}

// SetOnMouseMove
func (s *TStatusBar) SetOnMouseMove(fn TMouseMoveEvent) {
    StatusBar_SetOnMouseMove(s.instance, fn)
}

// SetOnMouseUp
func (s *TStatusBar) SetOnMouseUp(fn TMouseEvent) {
    StatusBar_SetOnMouseUp(s.instance, fn)
}

// SetOnResize
func (s *TStatusBar) SetOnResize(fn TNotifyEvent) {
    StatusBar_SetOnResize(s.instance, fn)
}

// SetOnStartDock
func (s *TStatusBar) SetOnStartDock(fn TStartDockEvent) {
    StatusBar_SetOnStartDock(s.instance, fn)
}

// Canvas
func (s *TStatusBar) Canvas() *TCanvas {
    return CanvasFromInst(StatusBar_GetCanvas(s.instance))
}

// DockSite
func (s *TStatusBar) DockSite() bool {
    return StatusBar_GetDockSite(s.instance)
}

// SetDockSite
func (s *TStatusBar) SetDockSite(value bool) {
    StatusBar_SetDockSite(s.instance, value)
}

// Brush
func (s *TStatusBar) Brush() *TBrush {
    return BrushFromInst(StatusBar_GetBrush(s.instance))
}

// ControlCount
func (s *TStatusBar) ControlCount() int32 {
    return StatusBar_GetControlCount(s.instance)
}

// Handle
func (s *TStatusBar) Handle() HWND {
    return StatusBar_GetHandle(s.instance)
}

// ParentWindow
func (s *TStatusBar) ParentWindow() HWND {
    return StatusBar_GetParentWindow(s.instance)
}

// SetParentWindow
func (s *TStatusBar) SetParentWindow(value HWND) {
    StatusBar_SetParentWindow(s.instance, value)
}

// TabOrder
func (s *TStatusBar) TabOrder() uint16 {
    return StatusBar_GetTabOrder(s.instance)
}

// SetTabOrder
func (s *TStatusBar) SetTabOrder(value uint16) {
    StatusBar_SetTabOrder(s.instance, value)
}

// TabStop
func (s *TStatusBar) TabStop() bool {
    return StatusBar_GetTabStop(s.instance)
}

// SetTabStop
func (s *TStatusBar) SetTabStop(value bool) {
    StatusBar_SetTabStop(s.instance, value)
}

// UseDockManager
func (s *TStatusBar) UseDockManager() bool {
    return StatusBar_GetUseDockManager(s.instance)
}

// SetUseDockManager
func (s *TStatusBar) SetUseDockManager(value bool) {
    StatusBar_SetUseDockManager(s.instance, value)
}

// BoundsRect
func (s *TStatusBar) BoundsRect() TRect {
    return StatusBar_GetBoundsRect(s.instance)
}

// SetBoundsRect
func (s *TStatusBar) SetBoundsRect(value TRect) {
    StatusBar_SetBoundsRect(s.instance, value)
}

// ClientHeight
func (s *TStatusBar) ClientHeight() int32 {
    return StatusBar_GetClientHeight(s.instance)
}

// SetClientHeight
func (s *TStatusBar) SetClientHeight(value int32) {
    StatusBar_SetClientHeight(s.instance, value)
}

// ClientRect
func (s *TStatusBar) ClientRect() TRect {
    return StatusBar_GetClientRect(s.instance)
}

// ClientWidth
func (s *TStatusBar) ClientWidth() int32 {
    return StatusBar_GetClientWidth(s.instance)
}

// SetClientWidth
func (s *TStatusBar) SetClientWidth(value int32) {
    StatusBar_SetClientWidth(s.instance, value)
}

// ExplicitLeft
func (s *TStatusBar) ExplicitLeft() int32 {
    return StatusBar_GetExplicitLeft(s.instance)
}

// ExplicitTop
func (s *TStatusBar) ExplicitTop() int32 {
    return StatusBar_GetExplicitTop(s.instance)
}

// ExplicitWidth
func (s *TStatusBar) ExplicitWidth() int32 {
    return StatusBar_GetExplicitWidth(s.instance)
}

// ExplicitHeight
func (s *TStatusBar) ExplicitHeight() int32 {
    return StatusBar_GetExplicitHeight(s.instance)
}

// Floating
func (s *TStatusBar) Floating() bool {
    return StatusBar_GetFloating(s.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TStatusBar) Parent() *TWinControl {
    return WinControlFromInst(StatusBar_GetParent(s.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TStatusBar) SetParent(value IWinControl) {
    StatusBar_SetParent(s.instance, CheckPtr(value))
}

// AlignWithMargins
func (s *TStatusBar) AlignWithMargins() bool {
    return StatusBar_GetAlignWithMargins(s.instance)
}

// SetAlignWithMargins
func (s *TStatusBar) SetAlignWithMargins(value bool) {
    StatusBar_SetAlignWithMargins(s.instance, value)
}

// Left
func (s *TStatusBar) Left() int32 {
    return StatusBar_GetLeft(s.instance)
}

// SetLeft
func (s *TStatusBar) SetLeft(value int32) {
    StatusBar_SetLeft(s.instance, value)
}

// Top
func (s *TStatusBar) Top() int32 {
    return StatusBar_GetTop(s.instance)
}

// SetTop
func (s *TStatusBar) SetTop(value int32) {
    StatusBar_SetTop(s.instance, value)
}

// Width
func (s *TStatusBar) Width() int32 {
    return StatusBar_GetWidth(s.instance)
}

// SetWidth
func (s *TStatusBar) SetWidth(value int32) {
    StatusBar_SetWidth(s.instance, value)
}

// Height
func (s *TStatusBar) Height() int32 {
    return StatusBar_GetHeight(s.instance)
}

// SetHeight
func (s *TStatusBar) SetHeight(value int32) {
    StatusBar_SetHeight(s.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TStatusBar) Cursor() TCursor {
    return StatusBar_GetCursor(s.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TStatusBar) SetCursor(value TCursor) {
    StatusBar_SetCursor(s.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (s *TStatusBar) Hint() string {
    return StatusBar_GetHint(s.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (s *TStatusBar) SetHint(value string) {
    StatusBar_SetHint(s.instance, value)
}

// Margins
func (s *TStatusBar) Margins() *TMargins {
    return MarginsFromInst(StatusBar_GetMargins(s.instance))
}

// SetMargins
func (s *TStatusBar) SetMargins(value *TMargins) {
    StatusBar_SetMargins(s.instance, CheckPtr(value))
}

// CustomHint
func (s *TStatusBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(StatusBar_GetCustomHint(s.instance))
}

// SetCustomHint
func (s *TStatusBar) SetCustomHint(value IComponent) {
    StatusBar_SetCustomHint(s.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TStatusBar) ComponentCount() int32 {
    return StatusBar_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TStatusBar) ComponentIndex() int32 {
    return StatusBar_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TStatusBar) SetComponentIndex(value int32) {
    StatusBar_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TStatusBar) Owner() *TComponent {
    return ComponentFromInst(StatusBar_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TStatusBar) Name() string {
    return StatusBar_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TStatusBar) SetName(value string) {
    StatusBar_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TStatusBar) Tag() int {
    return StatusBar_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TStatusBar) SetTag(value int) {
    StatusBar_SetTag(s.instance, value)
}

// Controls
func (s *TStatusBar) Controls(Index int32) *TControl {
    return ControlFromInst(StatusBar_GetControls(s.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TStatusBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(StatusBar_GetComponents(s.instance, AIndex))
}

