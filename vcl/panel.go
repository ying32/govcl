
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

type TPanel struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPanel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPanel(owner IComponent) *TPanel {
    p := new(TPanel)
    p.instance = Panel_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PanelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PanelFromInst(inst uintptr) *TPanel {
    p := new(TPanel)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PanelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PanelFromObj(obj IObject) *TPanel {
    p := new(TPanel)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PanelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PanelFromUnsafePointer(ptr unsafe.Pointer) *TPanel {
    p := new(TPanel)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPanel) Free() {
    if p.instance != 0 {
        Panel_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPanel) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPanel) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPanel) IsValid() bool {
    return p.instance != 0
}

// TPanelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPanelClass() TClass {
    return Panel_StaticClassType()
}

// CanFocus
func (p *TPanel) CanFocus() bool {
    return Panel_CanFocus(p.instance)
}

// FlipChildren
func (p *TPanel) FlipChildren(AllLevels bool) {
    Panel_FlipChildren(p.instance, AllLevels)
}

// Focused
func (p *TPanel) Focused() bool {
    return Panel_Focused(p.instance)
}

// HandleAllocated
func (p *TPanel) HandleAllocated() bool {
    return Panel_HandleAllocated(p.instance)
}

// Invalidate
func (p *TPanel) Invalidate() {
    Panel_Invalidate(p.instance)
}

// Realign
func (p *TPanel) Realign() {
    Panel_Realign(p.instance)
}

// Repaint
func (p *TPanel) Repaint() {
    Panel_Repaint(p.instance)
}

// ScaleBy
func (p *TPanel) ScaleBy(M int32, D int32) {
    Panel_ScaleBy(p.instance, M , D)
}

// SetBounds
func (p *TPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Panel_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (p *TPanel) SetFocus() {
    Panel_SetFocus(p.instance)
}

// Update
func (p *TPanel) Update() {
    Panel_Update(p.instance)
}

// BringToFront
func (p *TPanel) BringToFront() {
    Panel_BringToFront(p.instance)
}

// ClientToScreen
func (p *TPanel) ClientToScreen(Point TPoint) TPoint {
    return Panel_ClientToScreen(p.instance, Point)
}

// ClientToParent
func (p *TPanel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Panel_ClientToParent(p.instance, Point , CheckPtr(AParent))
}

// Dragging
func (p *TPanel) Dragging() bool {
    return Panel_Dragging(p.instance)
}

// HasParent
func (p *TPanel) HasParent() bool {
    return Panel_HasParent(p.instance)
}

// Hide
func (p *TPanel) Hide() {
    Panel_Hide(p.instance)
}

// Perform
func (p *TPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Panel_Perform(p.instance, Msg , WParam , LParam)
}

// Refresh
func (p *TPanel) Refresh() {
    Panel_Refresh(p.instance)
}

// ScreenToClient
func (p *TPanel) ScreenToClient(Point TPoint) TPoint {
    return Panel_ScreenToClient(p.instance, Point)
}

// ParentToClient
func (p *TPanel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Panel_ParentToClient(p.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (p *TPanel) SendToBack() {
    Panel_SendToBack(p.instance)
}

// Show
func (p *TPanel) Show() {
    Panel_Show(p.instance)
}

// GetTextBuf
func (p *TPanel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Panel_GetTextBuf(p.instance, Buffer , BufSize)
}

// GetTextLen
func (p *TPanel) GetTextLen() int32 {
    return Panel_GetTextLen(p.instance)
}

// FindComponent
func (p *TPanel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Panel_FindComponent(p.instance, AName))
}

// GetNamePath
func (p *TPanel) GetNamePath() string {
    return Panel_GetNamePath(p.instance)
}

// Assign
func (p *TPanel) Assign(Source IObject) {
    Panel_Assign(p.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPanel) DisposeOf() {
    Panel_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPanel) ClassType() TClass {
    return Panel_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPanel) ClassName() string {
    return Panel_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPanel) InstanceSize() int32 {
    return Panel_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPanel) InheritsFrom(AClass TClass) bool {
    return Panel_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPanel) Equals(Obj IObject) bool {
    return Panel_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPanel) GetHashCode() int32 {
    return Panel_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPanel) ToString() string {
    return Panel_ToString(p.instance)
}

// Align
func (p *TPanel) Align() TAlign {
    return Panel_GetAlign(p.instance)
}

// SetAlign
func (p *TPanel) SetAlign(value TAlign) {
    Panel_SetAlign(p.instance, value)
}

// Alignment
func (p *TPanel) Alignment() TAlignment {
    return Panel_GetAlignment(p.instance)
}

// SetAlignment
func (p *TPanel) SetAlignment(value TAlignment) {
    Panel_SetAlignment(p.instance, value)
}

// Anchors
func (p *TPanel) Anchors() TAnchors {
    return Panel_GetAnchors(p.instance)
}

// SetAnchors
func (p *TPanel) SetAnchors(value TAnchors) {
    Panel_SetAnchors(p.instance, value)
}

// AutoSize
func (p *TPanel) AutoSize() bool {
    return Panel_GetAutoSize(p.instance)
}

// SetAutoSize
func (p *TPanel) SetAutoSize(value bool) {
    Panel_SetAutoSize(p.instance, value)
}

// BevelEdges
func (p *TPanel) BevelEdges() TBevelEdges {
    return Panel_GetBevelEdges(p.instance)
}

// SetBevelEdges
func (p *TPanel) SetBevelEdges(value TBevelEdges) {
    Panel_SetBevelEdges(p.instance, value)
}

// BevelInner
func (p *TPanel) BevelInner() TBevelCut {
    return Panel_GetBevelInner(p.instance)
}

// SetBevelInner
func (p *TPanel) SetBevelInner(value TBevelCut) {
    Panel_SetBevelInner(p.instance, value)
}

// BevelKind
func (p *TPanel) BevelKind() TBevelKind {
    return Panel_GetBevelKind(p.instance)
}

// SetBevelKind
func (p *TPanel) SetBevelKind(value TBevelKind) {
    Panel_SetBevelKind(p.instance, value)
}

// BevelOuter
func (p *TPanel) BevelOuter() TBevelCut {
    return Panel_GetBevelOuter(p.instance)
}

// SetBevelOuter
func (p *TPanel) SetBevelOuter(value TBevelCut) {
    Panel_SetBevelOuter(p.instance, value)
}

// BiDiMode
func (p *TPanel) BiDiMode() TBiDiMode {
    return Panel_GetBiDiMode(p.instance)
}

// SetBiDiMode
func (p *TPanel) SetBiDiMode(value TBiDiMode) {
    Panel_SetBiDiMode(p.instance, value)
}

// BorderWidth
func (p *TPanel) BorderWidth() int32 {
    return Panel_GetBorderWidth(p.instance)
}

// SetBorderWidth
func (p *TPanel) SetBorderWidth(value int32) {
    Panel_SetBorderWidth(p.instance, value)
}

// BorderStyle
func (p *TPanel) BorderStyle() TBorderStyle {
    return Panel_GetBorderStyle(p.instance)
}

// SetBorderStyle
func (p *TPanel) SetBorderStyle(value TBorderStyle) {
    Panel_SetBorderStyle(p.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (p *TPanel) Caption() string {
    return Panel_GetCaption(p.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (p *TPanel) SetCaption(value string) {
    Panel_SetCaption(p.instance, value)
}

// Color
func (p *TPanel) Color() TColor {
    return Panel_GetColor(p.instance)
}

// SetColor
func (p *TPanel) SetColor(value TColor) {
    Panel_SetColor(p.instance, value)
}

// UseDockManager
func (p *TPanel) UseDockManager() bool {
    return Panel_GetUseDockManager(p.instance)
}

// SetUseDockManager
func (p *TPanel) SetUseDockManager(value bool) {
    Panel_SetUseDockManager(p.instance, value)
}

// DockSite
func (p *TPanel) DockSite() bool {
    return Panel_GetDockSite(p.instance)
}

// SetDockSite
func (p *TPanel) SetDockSite(value bool) {
    Panel_SetDockSite(p.instance, value)
}

// DoubleBuffered
func (p *TPanel) DoubleBuffered() bool {
    return Panel_GetDoubleBuffered(p.instance)
}

// SetDoubleBuffered
func (p *TPanel) SetDoubleBuffered(value bool) {
    Panel_SetDoubleBuffered(p.instance, value)
}

// DragCursor
func (p *TPanel) DragCursor() TCursor {
    return Panel_GetDragCursor(p.instance)
}

// SetDragCursor
func (p *TPanel) SetDragCursor(value TCursor) {
    Panel_SetDragCursor(p.instance, value)
}

// DragKind
func (p *TPanel) DragKind() TDragKind {
    return Panel_GetDragKind(p.instance)
}

// SetDragKind
func (p *TPanel) SetDragKind(value TDragKind) {
    Panel_SetDragKind(p.instance, value)
}

// DragMode
func (p *TPanel) DragMode() TDragMode {
    return Panel_GetDragMode(p.instance)
}

// SetDragMode
func (p *TPanel) SetDragMode(value TDragMode) {
    Panel_SetDragMode(p.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (p *TPanel) Enabled() bool {
    return Panel_GetEnabled(p.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (p *TPanel) SetEnabled(value bool) {
    Panel_SetEnabled(p.instance, value)
}

// FullRepaint
func (p *TPanel) FullRepaint() bool {
    return Panel_GetFullRepaint(p.instance)
}

// SetFullRepaint
func (p *TPanel) SetFullRepaint(value bool) {
    Panel_SetFullRepaint(p.instance, value)
}

// Font
func (p *TPanel) Font() *TFont {
    return FontFromInst(Panel_GetFont(p.instance))
}

// SetFont
func (p *TPanel) SetFont(value *TFont) {
    Panel_SetFont(p.instance, CheckPtr(value))
}

// Locked
func (p *TPanel) Locked() bool {
    return Panel_GetLocked(p.instance)
}

// SetLocked
func (p *TPanel) SetLocked(value bool) {
    Panel_SetLocked(p.instance, value)
}

// ParentBackground
func (p *TPanel) ParentBackground() bool {
    return Panel_GetParentBackground(p.instance)
}

// SetParentBackground
func (p *TPanel) SetParentBackground(value bool) {
    Panel_SetParentBackground(p.instance, value)
}

// ParentColor
func (p *TPanel) ParentColor() bool {
    return Panel_GetParentColor(p.instance)
}

// SetParentColor
func (p *TPanel) SetParentColor(value bool) {
    Panel_SetParentColor(p.instance, value)
}

// ParentCtl3D
func (p *TPanel) ParentCtl3D() bool {
    return Panel_GetParentCtl3D(p.instance)
}

// SetParentCtl3D
func (p *TPanel) SetParentCtl3D(value bool) {
    Panel_SetParentCtl3D(p.instance, value)
}

// ParentDoubleBuffered
func (p *TPanel) ParentDoubleBuffered() bool {
    return Panel_GetParentDoubleBuffered(p.instance)
}

// SetParentDoubleBuffered
func (p *TPanel) SetParentDoubleBuffered(value bool) {
    Panel_SetParentDoubleBuffered(p.instance, value)
}

// ParentFont
func (p *TPanel) ParentFont() bool {
    return Panel_GetParentFont(p.instance)
}

// SetParentFont
func (p *TPanel) SetParentFont(value bool) {
    Panel_SetParentFont(p.instance, value)
}

// ParentShowHint
func (p *TPanel) ParentShowHint() bool {
    return Panel_GetParentShowHint(p.instance)
}

// SetParentShowHint
func (p *TPanel) SetParentShowHint(value bool) {
    Panel_SetParentShowHint(p.instance, value)
}

// PopupMenu
func (p *TPanel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Panel_GetPopupMenu(p.instance))
}

// SetPopupMenu
func (p *TPanel) SetPopupMenu(value IComponent) {
    Panel_SetPopupMenu(p.instance, CheckPtr(value))
}

// ShowCaption
func (p *TPanel) ShowCaption() bool {
    return Panel_GetShowCaption(p.instance)
}

// SetShowCaption
func (p *TPanel) SetShowCaption(value bool) {
    Panel_SetShowCaption(p.instance, value)
}

// ShowHint
func (p *TPanel) ShowHint() bool {
    return Panel_GetShowHint(p.instance)
}

// SetShowHint
func (p *TPanel) SetShowHint(value bool) {
    Panel_SetShowHint(p.instance, value)
}

// TabOrder
func (p *TPanel) TabOrder() uint16 {
    return Panel_GetTabOrder(p.instance)
}

// SetTabOrder
func (p *TPanel) SetTabOrder(value uint16) {
    Panel_SetTabOrder(p.instance, value)
}

// TabStop
func (p *TPanel) TabStop() bool {
    return Panel_GetTabStop(p.instance)
}

// SetTabStop
func (p *TPanel) SetTabStop(value bool) {
    Panel_SetTabStop(p.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (p *TPanel) Visible() bool {
    return Panel_GetVisible(p.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (p *TPanel) SetVisible(value bool) {
    Panel_SetVisible(p.instance, value)
}

// StyleElements
func (p *TPanel) StyleElements() TStyleElements {
    return Panel_GetStyleElements(p.instance)
}

// SetStyleElements
func (p *TPanel) SetStyleElements(value TStyleElements) {
    Panel_SetStyleElements(p.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (p *TPanel) SetOnClick(fn TNotifyEvent) {
    Panel_SetOnClick(p.instance, fn)
}

// SetOnContextPopup
func (p *TPanel) SetOnContextPopup(fn TContextPopupEvent) {
    Panel_SetOnContextPopup(p.instance, fn)
}

// SetOnDockDrop
func (p *TPanel) SetOnDockDrop(fn TDockDropEvent) {
    Panel_SetOnDockDrop(p.instance, fn)
}

// SetOnDblClick
func (p *TPanel) SetOnDblClick(fn TNotifyEvent) {
    Panel_SetOnDblClick(p.instance, fn)
}

// SetOnDragDrop
func (p *TPanel) SetOnDragDrop(fn TDragDropEvent) {
    Panel_SetOnDragDrop(p.instance, fn)
}

// SetOnDragOver
func (p *TPanel) SetOnDragOver(fn TDragOverEvent) {
    Panel_SetOnDragOver(p.instance, fn)
}

// SetOnEndDock
func (p *TPanel) SetOnEndDock(fn TEndDragEvent) {
    Panel_SetOnEndDock(p.instance, fn)
}

// SetOnEndDrag
func (p *TPanel) SetOnEndDrag(fn TEndDragEvent) {
    Panel_SetOnEndDrag(p.instance, fn)
}

// SetOnEnter
func (p *TPanel) SetOnEnter(fn TNotifyEvent) {
    Panel_SetOnEnter(p.instance, fn)
}

// SetOnExit
func (p *TPanel) SetOnExit(fn TNotifyEvent) {
    Panel_SetOnExit(p.instance, fn)
}

// SetOnGetSiteInfo
func (p *TPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    Panel_SetOnGetSiteInfo(p.instance, fn)
}

// SetOnMouseDown
func (p *TPanel) SetOnMouseDown(fn TMouseEvent) {
    Panel_SetOnMouseDown(p.instance, fn)
}

// SetOnMouseEnter
func (p *TPanel) SetOnMouseEnter(fn TNotifyEvent) {
    Panel_SetOnMouseEnter(p.instance, fn)
}

// SetOnMouseLeave
func (p *TPanel) SetOnMouseLeave(fn TNotifyEvent) {
    Panel_SetOnMouseLeave(p.instance, fn)
}

// SetOnMouseMove
func (p *TPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    Panel_SetOnMouseMove(p.instance, fn)
}

// SetOnMouseUp
func (p *TPanel) SetOnMouseUp(fn TMouseEvent) {
    Panel_SetOnMouseUp(p.instance, fn)
}

// SetOnResize
func (p *TPanel) SetOnResize(fn TNotifyEvent) {
    Panel_SetOnResize(p.instance, fn)
}

// SetOnStartDock
func (p *TPanel) SetOnStartDock(fn TStartDockEvent) {
    Panel_SetOnStartDock(p.instance, fn)
}

// SetOnUnDock
func (p *TPanel) SetOnUnDock(fn TUnDockEvent) {
    Panel_SetOnUnDock(p.instance, fn)
}

// Brush
func (p *TPanel) Brush() *TBrush {
    return BrushFromInst(Panel_GetBrush(p.instance))
}

// ControlCount
func (p *TPanel) ControlCount() int32 {
    return Panel_GetControlCount(p.instance)
}

// Handle
func (p *TPanel) Handle() HWND {
    return Panel_GetHandle(p.instance)
}

// ParentWindow
func (p *TPanel) ParentWindow() HWND {
    return Panel_GetParentWindow(p.instance)
}

// SetParentWindow
func (p *TPanel) SetParentWindow(value HWND) {
    Panel_SetParentWindow(p.instance, value)
}

// Action
func (p *TPanel) Action() *TAction {
    return ActionFromInst(Panel_GetAction(p.instance))
}

// SetAction
func (p *TPanel) SetAction(value IComponent) {
    Panel_SetAction(p.instance, CheckPtr(value))
}

// BoundsRect
func (p *TPanel) BoundsRect() TRect {
    return Panel_GetBoundsRect(p.instance)
}

// SetBoundsRect
func (p *TPanel) SetBoundsRect(value TRect) {
    Panel_SetBoundsRect(p.instance, value)
}

// ClientHeight
func (p *TPanel) ClientHeight() int32 {
    return Panel_GetClientHeight(p.instance)
}

// SetClientHeight
func (p *TPanel) SetClientHeight(value int32) {
    Panel_SetClientHeight(p.instance, value)
}

// ClientRect
func (p *TPanel) ClientRect() TRect {
    return Panel_GetClientRect(p.instance)
}

// ClientWidth
func (p *TPanel) ClientWidth() int32 {
    return Panel_GetClientWidth(p.instance)
}

// SetClientWidth
func (p *TPanel) SetClientWidth(value int32) {
    Panel_SetClientWidth(p.instance, value)
}

// ExplicitLeft
func (p *TPanel) ExplicitLeft() int32 {
    return Panel_GetExplicitLeft(p.instance)
}

// ExplicitTop
func (p *TPanel) ExplicitTop() int32 {
    return Panel_GetExplicitTop(p.instance)
}

// ExplicitWidth
func (p *TPanel) ExplicitWidth() int32 {
    return Panel_GetExplicitWidth(p.instance)
}

// ExplicitHeight
func (p *TPanel) ExplicitHeight() int32 {
    return Panel_GetExplicitHeight(p.instance)
}

// Floating
func (p *TPanel) Floating() bool {
    return Panel_GetFloating(p.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (p *TPanel) Parent() *TWinControl {
    return WinControlFromInst(Panel_GetParent(p.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (p *TPanel) SetParent(value IWinControl) {
    Panel_SetParent(p.instance, CheckPtr(value))
}

// AlignWithMargins
func (p *TPanel) AlignWithMargins() bool {
    return Panel_GetAlignWithMargins(p.instance)
}

// SetAlignWithMargins
func (p *TPanel) SetAlignWithMargins(value bool) {
    Panel_SetAlignWithMargins(p.instance, value)
}

// Left
func (p *TPanel) Left() int32 {
    return Panel_GetLeft(p.instance)
}

// SetLeft
func (p *TPanel) SetLeft(value int32) {
    Panel_SetLeft(p.instance, value)
}

// Top
func (p *TPanel) Top() int32 {
    return Panel_GetTop(p.instance)
}

// SetTop
func (p *TPanel) SetTop(value int32) {
    Panel_SetTop(p.instance, value)
}

// Width
func (p *TPanel) Width() int32 {
    return Panel_GetWidth(p.instance)
}

// SetWidth
func (p *TPanel) SetWidth(value int32) {
    Panel_SetWidth(p.instance, value)
}

// Height
func (p *TPanel) Height() int32 {
    return Panel_GetHeight(p.instance)
}

// SetHeight
func (p *TPanel) SetHeight(value int32) {
    Panel_SetHeight(p.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (p *TPanel) Cursor() TCursor {
    return Panel_GetCursor(p.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (p *TPanel) SetCursor(value TCursor) {
    Panel_SetCursor(p.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (p *TPanel) Hint() string {
    return Panel_GetHint(p.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (p *TPanel) SetHint(value string) {
    Panel_SetHint(p.instance, value)
}

// Margins
func (p *TPanel) Margins() *TMargins {
    return MarginsFromInst(Panel_GetMargins(p.instance))
}

// SetMargins
func (p *TPanel) SetMargins(value *TMargins) {
    Panel_SetMargins(p.instance, CheckPtr(value))
}

// CustomHint
func (p *TPanel) CustomHint() *TCustomHint {
    return CustomHintFromInst(Panel_GetCustomHint(p.instance))
}

// SetCustomHint
func (p *TPanel) SetCustomHint(value IComponent) {
    Panel_SetCustomHint(p.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPanel) ComponentCount() int32 {
    return Panel_GetComponentCount(p.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (p *TPanel) ComponentIndex() int32 {
    return Panel_GetComponentIndex(p.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (p *TPanel) SetComponentIndex(value int32) {
    Panel_SetComponentIndex(p.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPanel) Owner() *TComponent {
    return ComponentFromInst(Panel_GetOwner(p.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPanel) Name() string {
    return Panel_GetName(p.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPanel) SetName(value string) {
    Panel_SetName(p.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPanel) Tag() int {
    return Panel_GetTag(p.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPanel) SetTag(value int) {
    Panel_SetTag(p.instance, value)
}

// Controls
func (p *TPanel) Controls(Index int32) *TControl {
    return ControlFromInst(Panel_GetControls(p.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPanel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Panel_GetComponents(p.instance, AIndex))
}

