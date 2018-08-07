
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

type TScrollBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewScrollBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewScrollBox(owner IComponent) *TScrollBox {
    s := new(TScrollBox)
    s.instance = ScrollBox_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// ScrollBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ScrollBoxFromInst(inst uintptr) *TScrollBox {
    s := new(TScrollBox)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// ScrollBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ScrollBoxFromObj(obj IObject) *TScrollBox {
    s := new(TScrollBox)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// ScrollBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ScrollBoxFromUnsafePointer(ptr unsafe.Pointer) *TScrollBox {
    s := new(TScrollBox)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TScrollBox) Free() {
    if s.instance != 0 {
        ScrollBox_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TScrollBox) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TScrollBox) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TScrollBox) IsValid() bool {
    return s.instance != 0
}

// TScrollBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TScrollBoxClass() TClass {
    return ScrollBox_StaticClassType()
}

// CanFocus
func (s *TScrollBox) CanFocus() bool {
    return ScrollBox_CanFocus(s.instance)
}

// FlipChildren
func (s *TScrollBox) FlipChildren(AllLevels bool) {
    ScrollBox_FlipChildren(s.instance, AllLevels)
}

// Focused
func (s *TScrollBox) Focused() bool {
    return ScrollBox_Focused(s.instance)
}

// HandleAllocated
func (s *TScrollBox) HandleAllocated() bool {
    return ScrollBox_HandleAllocated(s.instance)
}

// Invalidate
func (s *TScrollBox) Invalidate() {
    ScrollBox_Invalidate(s.instance)
}

// Realign
func (s *TScrollBox) Realign() {
    ScrollBox_Realign(s.instance)
}

// Repaint
func (s *TScrollBox) Repaint() {
    ScrollBox_Repaint(s.instance)
}

// ScaleBy
func (s *TScrollBox) ScaleBy(M int32, D int32) {
    ScrollBox_ScaleBy(s.instance, M , D)
}

// SetBounds
func (s *TScrollBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ScrollBox_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (s *TScrollBox) SetFocus() {
    ScrollBox_SetFocus(s.instance)
}

// Update
func (s *TScrollBox) Update() {
    ScrollBox_Update(s.instance)
}

// BringToFront
func (s *TScrollBox) BringToFront() {
    ScrollBox_BringToFront(s.instance)
}

// ClientToScreen
func (s *TScrollBox) ClientToScreen(Point TPoint) TPoint {
    return ScrollBox_ClientToScreen(s.instance, Point)
}

// ClientToParent
func (s *TScrollBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ScrollBox_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// Dragging
func (s *TScrollBox) Dragging() bool {
    return ScrollBox_Dragging(s.instance)
}

// HasParent
func (s *TScrollBox) HasParent() bool {
    return ScrollBox_HasParent(s.instance)
}

// Hide
func (s *TScrollBox) Hide() {
    ScrollBox_Hide(s.instance)
}

// Perform
func (s *TScrollBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ScrollBox_Perform(s.instance, Msg , WParam , LParam)
}

// Refresh
func (s *TScrollBox) Refresh() {
    ScrollBox_Refresh(s.instance)
}

// ScreenToClient
func (s *TScrollBox) ScreenToClient(Point TPoint) TPoint {
    return ScrollBox_ScreenToClient(s.instance, Point)
}

// ParentToClient
func (s *TScrollBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ScrollBox_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (s *TScrollBox) SendToBack() {
    ScrollBox_SendToBack(s.instance)
}

// Show
func (s *TScrollBox) Show() {
    ScrollBox_Show(s.instance)
}

// GetTextBuf
func (s *TScrollBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ScrollBox_GetTextBuf(s.instance, Buffer , BufSize)
}

// GetTextLen
func (s *TScrollBox) GetTextLen() int32 {
    return ScrollBox_GetTextLen(s.instance)
}

// FindComponent
func (s *TScrollBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ScrollBox_FindComponent(s.instance, AName))
}

// GetNamePath
func (s *TScrollBox) GetNamePath() string {
    return ScrollBox_GetNamePath(s.instance)
}

// Assign
func (s *TScrollBox) Assign(Source IObject) {
    ScrollBox_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TScrollBox) DisposeOf() {
    ScrollBox_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TScrollBox) ClassType() TClass {
    return ScrollBox_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TScrollBox) ClassName() string {
    return ScrollBox_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TScrollBox) InstanceSize() int32 {
    return ScrollBox_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TScrollBox) InheritsFrom(AClass TClass) bool {
    return ScrollBox_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TScrollBox) Equals(Obj IObject) bool {
    return ScrollBox_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TScrollBox) GetHashCode() int32 {
    return ScrollBox_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TScrollBox) ToString() string {
    return ScrollBox_ToString(s.instance)
}

// Align
func (s *TScrollBox) Align() TAlign {
    return ScrollBox_GetAlign(s.instance)
}

// SetAlign
func (s *TScrollBox) SetAlign(value TAlign) {
    ScrollBox_SetAlign(s.instance, value)
}

// Anchors
func (s *TScrollBox) Anchors() TAnchors {
    return ScrollBox_GetAnchors(s.instance)
}

// SetAnchors
func (s *TScrollBox) SetAnchors(value TAnchors) {
    ScrollBox_SetAnchors(s.instance, value)
}

// AutoSize
func (s *TScrollBox) AutoSize() bool {
    return ScrollBox_GetAutoSize(s.instance)
}

// SetAutoSize
func (s *TScrollBox) SetAutoSize(value bool) {
    ScrollBox_SetAutoSize(s.instance, value)
}

// BevelEdges
func (s *TScrollBox) BevelEdges() TBevelEdges {
    return ScrollBox_GetBevelEdges(s.instance)
}

// SetBevelEdges
func (s *TScrollBox) SetBevelEdges(value TBevelEdges) {
    ScrollBox_SetBevelEdges(s.instance, value)
}

// BevelInner
func (s *TScrollBox) BevelInner() TBevelCut {
    return ScrollBox_GetBevelInner(s.instance)
}

// SetBevelInner
func (s *TScrollBox) SetBevelInner(value TBevelCut) {
    ScrollBox_SetBevelInner(s.instance, value)
}

// BevelOuter
func (s *TScrollBox) BevelOuter() TBevelCut {
    return ScrollBox_GetBevelOuter(s.instance)
}

// SetBevelOuter
func (s *TScrollBox) SetBevelOuter(value TBevelCut) {
    ScrollBox_SetBevelOuter(s.instance, value)
}

// BevelKind
func (s *TScrollBox) BevelKind() TBevelKind {
    return ScrollBox_GetBevelKind(s.instance)
}

// SetBevelKind
func (s *TScrollBox) SetBevelKind(value TBevelKind) {
    ScrollBox_SetBevelKind(s.instance, value)
}

// BiDiMode
func (s *TScrollBox) BiDiMode() TBiDiMode {
    return ScrollBox_GetBiDiMode(s.instance)
}

// SetBiDiMode
func (s *TScrollBox) SetBiDiMode(value TBiDiMode) {
    ScrollBox_SetBiDiMode(s.instance, value)
}

// BorderStyle
func (s *TScrollBox) BorderStyle() TBorderStyle {
    return ScrollBox_GetBorderStyle(s.instance)
}

// SetBorderStyle
func (s *TScrollBox) SetBorderStyle(value TBorderStyle) {
    ScrollBox_SetBorderStyle(s.instance, value)
}

// DockSite
func (s *TScrollBox) DockSite() bool {
    return ScrollBox_GetDockSite(s.instance)
}

// SetDockSite
func (s *TScrollBox) SetDockSite(value bool) {
    ScrollBox_SetDockSite(s.instance, value)
}

// DoubleBuffered
func (s *TScrollBox) DoubleBuffered() bool {
    return ScrollBox_GetDoubleBuffered(s.instance)
}

// SetDoubleBuffered
func (s *TScrollBox) SetDoubleBuffered(value bool) {
    ScrollBox_SetDoubleBuffered(s.instance, value)
}

// DragCursor
func (s *TScrollBox) DragCursor() TCursor {
    return ScrollBox_GetDragCursor(s.instance)
}

// SetDragCursor
func (s *TScrollBox) SetDragCursor(value TCursor) {
    ScrollBox_SetDragCursor(s.instance, value)
}

// DragKind
func (s *TScrollBox) DragKind() TDragKind {
    return ScrollBox_GetDragKind(s.instance)
}

// SetDragKind
func (s *TScrollBox) SetDragKind(value TDragKind) {
    ScrollBox_SetDragKind(s.instance, value)
}

// DragMode
func (s *TScrollBox) DragMode() TDragMode {
    return ScrollBox_GetDragMode(s.instance)
}

// SetDragMode
func (s *TScrollBox) SetDragMode(value TDragMode) {
    ScrollBox_SetDragMode(s.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TScrollBox) Enabled() bool {
    return ScrollBox_GetEnabled(s.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TScrollBox) SetEnabled(value bool) {
    ScrollBox_SetEnabled(s.instance, value)
}

// Color
func (s *TScrollBox) Color() TColor {
    return ScrollBox_GetColor(s.instance)
}

// SetColor
func (s *TScrollBox) SetColor(value TColor) {
    ScrollBox_SetColor(s.instance, value)
}

// Font
func (s *TScrollBox) Font() *TFont {
    return FontFromInst(ScrollBox_GetFont(s.instance))
}

// SetFont
func (s *TScrollBox) SetFont(value *TFont) {
    ScrollBox_SetFont(s.instance, CheckPtr(value))
}

// ParentBackground
func (s *TScrollBox) ParentBackground() bool {
    return ScrollBox_GetParentBackground(s.instance)
}

// SetParentBackground
func (s *TScrollBox) SetParentBackground(value bool) {
    ScrollBox_SetParentBackground(s.instance, value)
}

// ParentColor
func (s *TScrollBox) ParentColor() bool {
    return ScrollBox_GetParentColor(s.instance)
}

// SetParentColor
func (s *TScrollBox) SetParentColor(value bool) {
    ScrollBox_SetParentColor(s.instance, value)
}

// ParentCtl3D
func (s *TScrollBox) ParentCtl3D() bool {
    return ScrollBox_GetParentCtl3D(s.instance)
}

// SetParentCtl3D
func (s *TScrollBox) SetParentCtl3D(value bool) {
    ScrollBox_SetParentCtl3D(s.instance, value)
}

// ParentDoubleBuffered
func (s *TScrollBox) ParentDoubleBuffered() bool {
    return ScrollBox_GetParentDoubleBuffered(s.instance)
}

// SetParentDoubleBuffered
func (s *TScrollBox) SetParentDoubleBuffered(value bool) {
    ScrollBox_SetParentDoubleBuffered(s.instance, value)
}

// ParentFont
func (s *TScrollBox) ParentFont() bool {
    return ScrollBox_GetParentFont(s.instance)
}

// SetParentFont
func (s *TScrollBox) SetParentFont(value bool) {
    ScrollBox_SetParentFont(s.instance, value)
}

// ParentShowHint
func (s *TScrollBox) ParentShowHint() bool {
    return ScrollBox_GetParentShowHint(s.instance)
}

// SetParentShowHint
func (s *TScrollBox) SetParentShowHint(value bool) {
    ScrollBox_SetParentShowHint(s.instance, value)
}

// PopupMenu
func (s *TScrollBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ScrollBox_GetPopupMenu(s.instance))
}

// SetPopupMenu
func (s *TScrollBox) SetPopupMenu(value IComponent) {
    ScrollBox_SetPopupMenu(s.instance, CheckPtr(value))
}

// ShowHint
func (s *TScrollBox) ShowHint() bool {
    return ScrollBox_GetShowHint(s.instance)
}

// SetShowHint
func (s *TScrollBox) SetShowHint(value bool) {
    ScrollBox_SetShowHint(s.instance, value)
}

// TabOrder
func (s *TScrollBox) TabOrder() uint16 {
    return ScrollBox_GetTabOrder(s.instance)
}

// SetTabOrder
func (s *TScrollBox) SetTabOrder(value uint16) {
    ScrollBox_SetTabOrder(s.instance, value)
}

// TabStop
func (s *TScrollBox) TabStop() bool {
    return ScrollBox_GetTabStop(s.instance)
}

// SetTabStop
func (s *TScrollBox) SetTabStop(value bool) {
    ScrollBox_SetTabStop(s.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TScrollBox) Visible() bool {
    return ScrollBox_GetVisible(s.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TScrollBox) SetVisible(value bool) {
    ScrollBox_SetVisible(s.instance, value)
}

// StyleElements
func (s *TScrollBox) StyleElements() TStyleElements {
    return ScrollBox_GetStyleElements(s.instance)
}

// SetStyleElements
func (s *TScrollBox) SetStyleElements(value TStyleElements) {
    ScrollBox_SetStyleElements(s.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (s *TScrollBox) SetOnClick(fn TNotifyEvent) {
    ScrollBox_SetOnClick(s.instance, fn)
}

// SetOnContextPopup
func (s *TScrollBox) SetOnContextPopup(fn TContextPopupEvent) {
    ScrollBox_SetOnContextPopup(s.instance, fn)
}

// SetOnDblClick
func (s *TScrollBox) SetOnDblClick(fn TNotifyEvent) {
    ScrollBox_SetOnDblClick(s.instance, fn)
}

// SetOnDockDrop
func (s *TScrollBox) SetOnDockDrop(fn TDockDropEvent) {
    ScrollBox_SetOnDockDrop(s.instance, fn)
}

// SetOnDragDrop
func (s *TScrollBox) SetOnDragDrop(fn TDragDropEvent) {
    ScrollBox_SetOnDragDrop(s.instance, fn)
}

// SetOnDragOver
func (s *TScrollBox) SetOnDragOver(fn TDragOverEvent) {
    ScrollBox_SetOnDragOver(s.instance, fn)
}

// SetOnEndDock
func (s *TScrollBox) SetOnEndDock(fn TEndDragEvent) {
    ScrollBox_SetOnEndDock(s.instance, fn)
}

// SetOnEndDrag
func (s *TScrollBox) SetOnEndDrag(fn TEndDragEvent) {
    ScrollBox_SetOnEndDrag(s.instance, fn)
}

// SetOnEnter
func (s *TScrollBox) SetOnEnter(fn TNotifyEvent) {
    ScrollBox_SetOnEnter(s.instance, fn)
}

// SetOnExit
func (s *TScrollBox) SetOnExit(fn TNotifyEvent) {
    ScrollBox_SetOnExit(s.instance, fn)
}

// SetOnGetSiteInfo
func (s *TScrollBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    ScrollBox_SetOnGetSiteInfo(s.instance, fn)
}

// SetOnMouseDown
func (s *TScrollBox) SetOnMouseDown(fn TMouseEvent) {
    ScrollBox_SetOnMouseDown(s.instance, fn)
}

// SetOnMouseEnter
func (s *TScrollBox) SetOnMouseEnter(fn TNotifyEvent) {
    ScrollBox_SetOnMouseEnter(s.instance, fn)
}

// SetOnMouseLeave
func (s *TScrollBox) SetOnMouseLeave(fn TNotifyEvent) {
    ScrollBox_SetOnMouseLeave(s.instance, fn)
}

// SetOnMouseMove
func (s *TScrollBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ScrollBox_SetOnMouseMove(s.instance, fn)
}

// SetOnMouseUp
func (s *TScrollBox) SetOnMouseUp(fn TMouseEvent) {
    ScrollBox_SetOnMouseUp(s.instance, fn)
}

// SetOnMouseWheel
func (s *TScrollBox) SetOnMouseWheel(fn TMouseWheelEvent) {
    ScrollBox_SetOnMouseWheel(s.instance, fn)
}

// SetOnMouseWheelDown
func (s *TScrollBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    ScrollBox_SetOnMouseWheelDown(s.instance, fn)
}

// SetOnMouseWheelUp
func (s *TScrollBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    ScrollBox_SetOnMouseWheelUp(s.instance, fn)
}

// SetOnResize
func (s *TScrollBox) SetOnResize(fn TNotifyEvent) {
    ScrollBox_SetOnResize(s.instance, fn)
}

// SetOnStartDock
func (s *TScrollBox) SetOnStartDock(fn TStartDockEvent) {
    ScrollBox_SetOnStartDock(s.instance, fn)
}

// SetOnUnDock
func (s *TScrollBox) SetOnUnDock(fn TUnDockEvent) {
    ScrollBox_SetOnUnDock(s.instance, fn)
}

// Brush
func (s *TScrollBox) Brush() *TBrush {
    return BrushFromInst(ScrollBox_GetBrush(s.instance))
}

// ControlCount
func (s *TScrollBox) ControlCount() int32 {
    return ScrollBox_GetControlCount(s.instance)
}

// Handle
func (s *TScrollBox) Handle() HWND {
    return ScrollBox_GetHandle(s.instance)
}

// ParentWindow
func (s *TScrollBox) ParentWindow() HWND {
    return ScrollBox_GetParentWindow(s.instance)
}

// SetParentWindow
func (s *TScrollBox) SetParentWindow(value HWND) {
    ScrollBox_SetParentWindow(s.instance, value)
}

// UseDockManager
func (s *TScrollBox) UseDockManager() bool {
    return ScrollBox_GetUseDockManager(s.instance)
}

// SetUseDockManager
func (s *TScrollBox) SetUseDockManager(value bool) {
    ScrollBox_SetUseDockManager(s.instance, value)
}

// Action
func (s *TScrollBox) Action() *TAction {
    return ActionFromInst(ScrollBox_GetAction(s.instance))
}

// SetAction
func (s *TScrollBox) SetAction(value IComponent) {
    ScrollBox_SetAction(s.instance, CheckPtr(value))
}

// BoundsRect
func (s *TScrollBox) BoundsRect() TRect {
    return ScrollBox_GetBoundsRect(s.instance)
}

// SetBoundsRect
func (s *TScrollBox) SetBoundsRect(value TRect) {
    ScrollBox_SetBoundsRect(s.instance, value)
}

// ClientHeight
func (s *TScrollBox) ClientHeight() int32 {
    return ScrollBox_GetClientHeight(s.instance)
}

// SetClientHeight
func (s *TScrollBox) SetClientHeight(value int32) {
    ScrollBox_SetClientHeight(s.instance, value)
}

// ClientRect
func (s *TScrollBox) ClientRect() TRect {
    return ScrollBox_GetClientRect(s.instance)
}

// ClientWidth
func (s *TScrollBox) ClientWidth() int32 {
    return ScrollBox_GetClientWidth(s.instance)
}

// SetClientWidth
func (s *TScrollBox) SetClientWidth(value int32) {
    ScrollBox_SetClientWidth(s.instance, value)
}

// ExplicitLeft
func (s *TScrollBox) ExplicitLeft() int32 {
    return ScrollBox_GetExplicitLeft(s.instance)
}

// ExplicitTop
func (s *TScrollBox) ExplicitTop() int32 {
    return ScrollBox_GetExplicitTop(s.instance)
}

// ExplicitWidth
func (s *TScrollBox) ExplicitWidth() int32 {
    return ScrollBox_GetExplicitWidth(s.instance)
}

// ExplicitHeight
func (s *TScrollBox) ExplicitHeight() int32 {
    return ScrollBox_GetExplicitHeight(s.instance)
}

// Floating
func (s *TScrollBox) Floating() bool {
    return ScrollBox_GetFloating(s.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TScrollBox) Parent() *TWinControl {
    return WinControlFromInst(ScrollBox_GetParent(s.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TScrollBox) SetParent(value IWinControl) {
    ScrollBox_SetParent(s.instance, CheckPtr(value))
}

// AlignWithMargins
func (s *TScrollBox) AlignWithMargins() bool {
    return ScrollBox_GetAlignWithMargins(s.instance)
}

// SetAlignWithMargins
func (s *TScrollBox) SetAlignWithMargins(value bool) {
    ScrollBox_SetAlignWithMargins(s.instance, value)
}

// Left
func (s *TScrollBox) Left() int32 {
    return ScrollBox_GetLeft(s.instance)
}

// SetLeft
func (s *TScrollBox) SetLeft(value int32) {
    ScrollBox_SetLeft(s.instance, value)
}

// Top
func (s *TScrollBox) Top() int32 {
    return ScrollBox_GetTop(s.instance)
}

// SetTop
func (s *TScrollBox) SetTop(value int32) {
    ScrollBox_SetTop(s.instance, value)
}

// Width
func (s *TScrollBox) Width() int32 {
    return ScrollBox_GetWidth(s.instance)
}

// SetWidth
func (s *TScrollBox) SetWidth(value int32) {
    ScrollBox_SetWidth(s.instance, value)
}

// Height
func (s *TScrollBox) Height() int32 {
    return ScrollBox_GetHeight(s.instance)
}

// SetHeight
func (s *TScrollBox) SetHeight(value int32) {
    ScrollBox_SetHeight(s.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TScrollBox) Cursor() TCursor {
    return ScrollBox_GetCursor(s.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TScrollBox) SetCursor(value TCursor) {
    ScrollBox_SetCursor(s.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (s *TScrollBox) Hint() string {
    return ScrollBox_GetHint(s.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (s *TScrollBox) SetHint(value string) {
    ScrollBox_SetHint(s.instance, value)
}

// Margins
func (s *TScrollBox) Margins() *TMargins {
    return MarginsFromInst(ScrollBox_GetMargins(s.instance))
}

// SetMargins
func (s *TScrollBox) SetMargins(value *TMargins) {
    ScrollBox_SetMargins(s.instance, CheckPtr(value))
}

// CustomHint
func (s *TScrollBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ScrollBox_GetCustomHint(s.instance))
}

// SetCustomHint
func (s *TScrollBox) SetCustomHint(value IComponent) {
    ScrollBox_SetCustomHint(s.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TScrollBox) ComponentCount() int32 {
    return ScrollBox_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TScrollBox) ComponentIndex() int32 {
    return ScrollBox_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TScrollBox) SetComponentIndex(value int32) {
    ScrollBox_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TScrollBox) Owner() *TComponent {
    return ComponentFromInst(ScrollBox_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TScrollBox) Name() string {
    return ScrollBox_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TScrollBox) SetName(value string) {
    ScrollBox_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TScrollBox) Tag() int {
    return ScrollBox_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TScrollBox) SetTag(value int) {
    ScrollBox_SetTag(s.instance, value)
}

// Controls
func (s *TScrollBox) Controls(Index int32) *TControl {
    return ControlFromInst(ScrollBox_GetControls(s.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TScrollBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ScrollBox_GetComponents(s.instance, AIndex))
}

