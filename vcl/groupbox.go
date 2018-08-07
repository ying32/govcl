
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

type TGroupBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewGroupBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewGroupBox(owner IComponent) *TGroupBox {
    g := new(TGroupBox)
    g.instance = GroupBox_Create(CheckPtr(owner))
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GroupBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func GroupBoxFromInst(inst uintptr) *TGroupBox {
    g := new(TGroupBox)
    g.instance = inst
    g.ptr = unsafe.Pointer(inst)
    return g
}

// GroupBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func GroupBoxFromObj(obj IObject) *TGroupBox {
    g := new(TGroupBox)
    g.instance = CheckPtr(obj)
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GroupBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func GroupBoxFromUnsafePointer(ptr unsafe.Pointer) *TGroupBox {
    g := new(TGroupBox)
    g.instance = uintptr(ptr)
    g.ptr = ptr
    return g
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (g *TGroupBox) Free() {
    if g.instance != 0 {
        GroupBox_Free(g.instance)
        g.instance = 0
        g.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (g *TGroupBox) Instance() uintptr {
    return g.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (g *TGroupBox) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (g *TGroupBox) IsValid() bool {
    return g.instance != 0
}

// TGroupBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TGroupBoxClass() TClass {
    return GroupBox_StaticClassType()
}

// CanFocus
func (g *TGroupBox) CanFocus() bool {
    return GroupBox_CanFocus(g.instance)
}

// FlipChildren
func (g *TGroupBox) FlipChildren(AllLevels bool) {
    GroupBox_FlipChildren(g.instance, AllLevels)
}

// Focused
func (g *TGroupBox) Focused() bool {
    return GroupBox_Focused(g.instance)
}

// HandleAllocated
func (g *TGroupBox) HandleAllocated() bool {
    return GroupBox_HandleAllocated(g.instance)
}

// Invalidate
func (g *TGroupBox) Invalidate() {
    GroupBox_Invalidate(g.instance)
}

// Realign
func (g *TGroupBox) Realign() {
    GroupBox_Realign(g.instance)
}

// Repaint
func (g *TGroupBox) Repaint() {
    GroupBox_Repaint(g.instance)
}

// ScaleBy
func (g *TGroupBox) ScaleBy(M int32, D int32) {
    GroupBox_ScaleBy(g.instance, M , D)
}

// SetBounds
func (g *TGroupBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    GroupBox_SetBounds(g.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (g *TGroupBox) SetFocus() {
    GroupBox_SetFocus(g.instance)
}

// Update
func (g *TGroupBox) Update() {
    GroupBox_Update(g.instance)
}

// BringToFront
func (g *TGroupBox) BringToFront() {
    GroupBox_BringToFront(g.instance)
}

// ClientToScreen
func (g *TGroupBox) ClientToScreen(Point TPoint) TPoint {
    return GroupBox_ClientToScreen(g.instance, Point)
}

// ClientToParent
func (g *TGroupBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return GroupBox_ClientToParent(g.instance, Point , CheckPtr(AParent))
}

// Dragging
func (g *TGroupBox) Dragging() bool {
    return GroupBox_Dragging(g.instance)
}

// HasParent
func (g *TGroupBox) HasParent() bool {
    return GroupBox_HasParent(g.instance)
}

// Hide
func (g *TGroupBox) Hide() {
    GroupBox_Hide(g.instance)
}

// Perform
func (g *TGroupBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return GroupBox_Perform(g.instance, Msg , WParam , LParam)
}

// Refresh
func (g *TGroupBox) Refresh() {
    GroupBox_Refresh(g.instance)
}

// ScreenToClient
func (g *TGroupBox) ScreenToClient(Point TPoint) TPoint {
    return GroupBox_ScreenToClient(g.instance, Point)
}

// ParentToClient
func (g *TGroupBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return GroupBox_ParentToClient(g.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (g *TGroupBox) SendToBack() {
    GroupBox_SendToBack(g.instance)
}

// Show
func (g *TGroupBox) Show() {
    GroupBox_Show(g.instance)
}

// GetTextBuf
func (g *TGroupBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return GroupBox_GetTextBuf(g.instance, Buffer , BufSize)
}

// GetTextLen
func (g *TGroupBox) GetTextLen() int32 {
    return GroupBox_GetTextLen(g.instance)
}

// FindComponent
func (g *TGroupBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(GroupBox_FindComponent(g.instance, AName))
}

// GetNamePath
func (g *TGroupBox) GetNamePath() string {
    return GroupBox_GetNamePath(g.instance)
}

// Assign
func (g *TGroupBox) Assign(Source IObject) {
    GroupBox_Assign(g.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (g *TGroupBox) DisposeOf() {
    GroupBox_DisposeOf(g.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (g *TGroupBox) ClassType() TClass {
    return GroupBox_ClassType(g.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (g *TGroupBox) ClassName() string {
    return GroupBox_ClassName(g.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (g *TGroupBox) InstanceSize() int32 {
    return GroupBox_InstanceSize(g.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (g *TGroupBox) InheritsFrom(AClass TClass) bool {
    return GroupBox_InheritsFrom(g.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (g *TGroupBox) Equals(Obj IObject) bool {
    return GroupBox_Equals(g.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (g *TGroupBox) GetHashCode() int32 {
    return GroupBox_GetHashCode(g.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (g *TGroupBox) ToString() string {
    return GroupBox_ToString(g.instance)
}

// Align
func (g *TGroupBox) Align() TAlign {
    return GroupBox_GetAlign(g.instance)
}

// SetAlign
func (g *TGroupBox) SetAlign(value TAlign) {
    GroupBox_SetAlign(g.instance, value)
}

// Anchors
func (g *TGroupBox) Anchors() TAnchors {
    return GroupBox_GetAnchors(g.instance)
}

// SetAnchors
func (g *TGroupBox) SetAnchors(value TAnchors) {
    GroupBox_SetAnchors(g.instance, value)
}

// BiDiMode
func (g *TGroupBox) BiDiMode() TBiDiMode {
    return GroupBox_GetBiDiMode(g.instance)
}

// SetBiDiMode
func (g *TGroupBox) SetBiDiMode(value TBiDiMode) {
    GroupBox_SetBiDiMode(g.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (g *TGroupBox) Caption() string {
    return GroupBox_GetCaption(g.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (g *TGroupBox) SetCaption(value string) {
    GroupBox_SetCaption(g.instance, value)
}

// Color
func (g *TGroupBox) Color() TColor {
    return GroupBox_GetColor(g.instance)
}

// SetColor
func (g *TGroupBox) SetColor(value TColor) {
    GroupBox_SetColor(g.instance, value)
}

// DockSite
func (g *TGroupBox) DockSite() bool {
    return GroupBox_GetDockSite(g.instance)
}

// SetDockSite
func (g *TGroupBox) SetDockSite(value bool) {
    GroupBox_SetDockSite(g.instance, value)
}

// DoubleBuffered
func (g *TGroupBox) DoubleBuffered() bool {
    return GroupBox_GetDoubleBuffered(g.instance)
}

// SetDoubleBuffered
func (g *TGroupBox) SetDoubleBuffered(value bool) {
    GroupBox_SetDoubleBuffered(g.instance, value)
}

// DragCursor
func (g *TGroupBox) DragCursor() TCursor {
    return GroupBox_GetDragCursor(g.instance)
}

// SetDragCursor
func (g *TGroupBox) SetDragCursor(value TCursor) {
    GroupBox_SetDragCursor(g.instance, value)
}

// DragKind
func (g *TGroupBox) DragKind() TDragKind {
    return GroupBox_GetDragKind(g.instance)
}

// SetDragKind
func (g *TGroupBox) SetDragKind(value TDragKind) {
    GroupBox_SetDragKind(g.instance, value)
}

// DragMode
func (g *TGroupBox) DragMode() TDragMode {
    return GroupBox_GetDragMode(g.instance)
}

// SetDragMode
func (g *TGroupBox) SetDragMode(value TDragMode) {
    GroupBox_SetDragMode(g.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (g *TGroupBox) Enabled() bool {
    return GroupBox_GetEnabled(g.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (g *TGroupBox) SetEnabled(value bool) {
    GroupBox_SetEnabled(g.instance, value)
}

// Font
func (g *TGroupBox) Font() *TFont {
    return FontFromInst(GroupBox_GetFont(g.instance))
}

// SetFont
func (g *TGroupBox) SetFont(value *TFont) {
    GroupBox_SetFont(g.instance, CheckPtr(value))
}

// ParentBackground
func (g *TGroupBox) ParentBackground() bool {
    return GroupBox_GetParentBackground(g.instance)
}

// SetParentBackground
func (g *TGroupBox) SetParentBackground(value bool) {
    GroupBox_SetParentBackground(g.instance, value)
}

// ParentColor
func (g *TGroupBox) ParentColor() bool {
    return GroupBox_GetParentColor(g.instance)
}

// SetParentColor
func (g *TGroupBox) SetParentColor(value bool) {
    GroupBox_SetParentColor(g.instance, value)
}

// ParentCtl3D
func (g *TGroupBox) ParentCtl3D() bool {
    return GroupBox_GetParentCtl3D(g.instance)
}

// SetParentCtl3D
func (g *TGroupBox) SetParentCtl3D(value bool) {
    GroupBox_SetParentCtl3D(g.instance, value)
}

// ParentDoubleBuffered
func (g *TGroupBox) ParentDoubleBuffered() bool {
    return GroupBox_GetParentDoubleBuffered(g.instance)
}

// SetParentDoubleBuffered
func (g *TGroupBox) SetParentDoubleBuffered(value bool) {
    GroupBox_SetParentDoubleBuffered(g.instance, value)
}

// ParentFont
func (g *TGroupBox) ParentFont() bool {
    return GroupBox_GetParentFont(g.instance)
}

// SetParentFont
func (g *TGroupBox) SetParentFont(value bool) {
    GroupBox_SetParentFont(g.instance, value)
}

// ParentShowHint
func (g *TGroupBox) ParentShowHint() bool {
    return GroupBox_GetParentShowHint(g.instance)
}

// SetParentShowHint
func (g *TGroupBox) SetParentShowHint(value bool) {
    GroupBox_SetParentShowHint(g.instance, value)
}

// PopupMenu
func (g *TGroupBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(GroupBox_GetPopupMenu(g.instance))
}

// SetPopupMenu
func (g *TGroupBox) SetPopupMenu(value IComponent) {
    GroupBox_SetPopupMenu(g.instance, CheckPtr(value))
}

// ShowHint
func (g *TGroupBox) ShowHint() bool {
    return GroupBox_GetShowHint(g.instance)
}

// SetShowHint
func (g *TGroupBox) SetShowHint(value bool) {
    GroupBox_SetShowHint(g.instance, value)
}

// TabOrder
func (g *TGroupBox) TabOrder() uint16 {
    return GroupBox_GetTabOrder(g.instance)
}

// SetTabOrder
func (g *TGroupBox) SetTabOrder(value uint16) {
    GroupBox_SetTabOrder(g.instance, value)
}

// TabStop
func (g *TGroupBox) TabStop() bool {
    return GroupBox_GetTabStop(g.instance)
}

// SetTabStop
func (g *TGroupBox) SetTabStop(value bool) {
    GroupBox_SetTabStop(g.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (g *TGroupBox) Visible() bool {
    return GroupBox_GetVisible(g.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (g *TGroupBox) SetVisible(value bool) {
    GroupBox_SetVisible(g.instance, value)
}

// StyleElements
func (g *TGroupBox) StyleElements() TStyleElements {
    return GroupBox_GetStyleElements(g.instance)
}

// SetStyleElements
func (g *TGroupBox) SetStyleElements(value TStyleElements) {
    GroupBox_SetStyleElements(g.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (g *TGroupBox) SetOnClick(fn TNotifyEvent) {
    GroupBox_SetOnClick(g.instance, fn)
}

// SetOnContextPopup
func (g *TGroupBox) SetOnContextPopup(fn TContextPopupEvent) {
    GroupBox_SetOnContextPopup(g.instance, fn)
}

// SetOnDblClick
func (g *TGroupBox) SetOnDblClick(fn TNotifyEvent) {
    GroupBox_SetOnDblClick(g.instance, fn)
}

// SetOnDragDrop
func (g *TGroupBox) SetOnDragDrop(fn TDragDropEvent) {
    GroupBox_SetOnDragDrop(g.instance, fn)
}

// SetOnDockDrop
func (g *TGroupBox) SetOnDockDrop(fn TDockDropEvent) {
    GroupBox_SetOnDockDrop(g.instance, fn)
}

// SetOnDragOver
func (g *TGroupBox) SetOnDragOver(fn TDragOverEvent) {
    GroupBox_SetOnDragOver(g.instance, fn)
}

// SetOnEndDock
func (g *TGroupBox) SetOnEndDock(fn TEndDragEvent) {
    GroupBox_SetOnEndDock(g.instance, fn)
}

// SetOnEndDrag
func (g *TGroupBox) SetOnEndDrag(fn TEndDragEvent) {
    GroupBox_SetOnEndDrag(g.instance, fn)
}

// SetOnEnter
func (g *TGroupBox) SetOnEnter(fn TNotifyEvent) {
    GroupBox_SetOnEnter(g.instance, fn)
}

// SetOnExit
func (g *TGroupBox) SetOnExit(fn TNotifyEvent) {
    GroupBox_SetOnExit(g.instance, fn)
}

// SetOnGetSiteInfo
func (g *TGroupBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    GroupBox_SetOnGetSiteInfo(g.instance, fn)
}

// SetOnMouseDown
func (g *TGroupBox) SetOnMouseDown(fn TMouseEvent) {
    GroupBox_SetOnMouseDown(g.instance, fn)
}

// SetOnMouseEnter
func (g *TGroupBox) SetOnMouseEnter(fn TNotifyEvent) {
    GroupBox_SetOnMouseEnter(g.instance, fn)
}

// SetOnMouseLeave
func (g *TGroupBox) SetOnMouseLeave(fn TNotifyEvent) {
    GroupBox_SetOnMouseLeave(g.instance, fn)
}

// SetOnMouseMove
func (g *TGroupBox) SetOnMouseMove(fn TMouseMoveEvent) {
    GroupBox_SetOnMouseMove(g.instance, fn)
}

// SetOnMouseUp
func (g *TGroupBox) SetOnMouseUp(fn TMouseEvent) {
    GroupBox_SetOnMouseUp(g.instance, fn)
}

// SetOnStartDock
func (g *TGroupBox) SetOnStartDock(fn TStartDockEvent) {
    GroupBox_SetOnStartDock(g.instance, fn)
}

// SetOnUnDock
func (g *TGroupBox) SetOnUnDock(fn TUnDockEvent) {
    GroupBox_SetOnUnDock(g.instance, fn)
}

// Brush
func (g *TGroupBox) Brush() *TBrush {
    return BrushFromInst(GroupBox_GetBrush(g.instance))
}

// ControlCount
func (g *TGroupBox) ControlCount() int32 {
    return GroupBox_GetControlCount(g.instance)
}

// Handle
func (g *TGroupBox) Handle() HWND {
    return GroupBox_GetHandle(g.instance)
}

// ParentWindow
func (g *TGroupBox) ParentWindow() HWND {
    return GroupBox_GetParentWindow(g.instance)
}

// SetParentWindow
func (g *TGroupBox) SetParentWindow(value HWND) {
    GroupBox_SetParentWindow(g.instance, value)
}

// UseDockManager
func (g *TGroupBox) UseDockManager() bool {
    return GroupBox_GetUseDockManager(g.instance)
}

// SetUseDockManager
func (g *TGroupBox) SetUseDockManager(value bool) {
    GroupBox_SetUseDockManager(g.instance, value)
}

// Action
func (g *TGroupBox) Action() *TAction {
    return ActionFromInst(GroupBox_GetAction(g.instance))
}

// SetAction
func (g *TGroupBox) SetAction(value IComponent) {
    GroupBox_SetAction(g.instance, CheckPtr(value))
}

// BoundsRect
func (g *TGroupBox) BoundsRect() TRect {
    return GroupBox_GetBoundsRect(g.instance)
}

// SetBoundsRect
func (g *TGroupBox) SetBoundsRect(value TRect) {
    GroupBox_SetBoundsRect(g.instance, value)
}

// ClientHeight
func (g *TGroupBox) ClientHeight() int32 {
    return GroupBox_GetClientHeight(g.instance)
}

// SetClientHeight
func (g *TGroupBox) SetClientHeight(value int32) {
    GroupBox_SetClientHeight(g.instance, value)
}

// ClientRect
func (g *TGroupBox) ClientRect() TRect {
    return GroupBox_GetClientRect(g.instance)
}

// ClientWidth
func (g *TGroupBox) ClientWidth() int32 {
    return GroupBox_GetClientWidth(g.instance)
}

// SetClientWidth
func (g *TGroupBox) SetClientWidth(value int32) {
    GroupBox_SetClientWidth(g.instance, value)
}

// ExplicitLeft
func (g *TGroupBox) ExplicitLeft() int32 {
    return GroupBox_GetExplicitLeft(g.instance)
}

// ExplicitTop
func (g *TGroupBox) ExplicitTop() int32 {
    return GroupBox_GetExplicitTop(g.instance)
}

// ExplicitWidth
func (g *TGroupBox) ExplicitWidth() int32 {
    return GroupBox_GetExplicitWidth(g.instance)
}

// ExplicitHeight
func (g *TGroupBox) ExplicitHeight() int32 {
    return GroupBox_GetExplicitHeight(g.instance)
}

// Floating
func (g *TGroupBox) Floating() bool {
    return GroupBox_GetFloating(g.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (g *TGroupBox) Parent() *TWinControl {
    return WinControlFromInst(GroupBox_GetParent(g.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (g *TGroupBox) SetParent(value IWinControl) {
    GroupBox_SetParent(g.instance, CheckPtr(value))
}

// AlignWithMargins
func (g *TGroupBox) AlignWithMargins() bool {
    return GroupBox_GetAlignWithMargins(g.instance)
}

// SetAlignWithMargins
func (g *TGroupBox) SetAlignWithMargins(value bool) {
    GroupBox_SetAlignWithMargins(g.instance, value)
}

// Left
func (g *TGroupBox) Left() int32 {
    return GroupBox_GetLeft(g.instance)
}

// SetLeft
func (g *TGroupBox) SetLeft(value int32) {
    GroupBox_SetLeft(g.instance, value)
}

// Top
func (g *TGroupBox) Top() int32 {
    return GroupBox_GetTop(g.instance)
}

// SetTop
func (g *TGroupBox) SetTop(value int32) {
    GroupBox_SetTop(g.instance, value)
}

// Width
func (g *TGroupBox) Width() int32 {
    return GroupBox_GetWidth(g.instance)
}

// SetWidth
func (g *TGroupBox) SetWidth(value int32) {
    GroupBox_SetWidth(g.instance, value)
}

// Height
func (g *TGroupBox) Height() int32 {
    return GroupBox_GetHeight(g.instance)
}

// SetHeight
func (g *TGroupBox) SetHeight(value int32) {
    GroupBox_SetHeight(g.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (g *TGroupBox) Cursor() TCursor {
    return GroupBox_GetCursor(g.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (g *TGroupBox) SetCursor(value TCursor) {
    GroupBox_SetCursor(g.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (g *TGroupBox) Hint() string {
    return GroupBox_GetHint(g.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (g *TGroupBox) SetHint(value string) {
    GroupBox_SetHint(g.instance, value)
}

// Margins
func (g *TGroupBox) Margins() *TMargins {
    return MarginsFromInst(GroupBox_GetMargins(g.instance))
}

// SetMargins
func (g *TGroupBox) SetMargins(value *TMargins) {
    GroupBox_SetMargins(g.instance, CheckPtr(value))
}

// CustomHint
func (g *TGroupBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(GroupBox_GetCustomHint(g.instance))
}

// SetCustomHint
func (g *TGroupBox) SetCustomHint(value IComponent) {
    GroupBox_SetCustomHint(g.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (g *TGroupBox) ComponentCount() int32 {
    return GroupBox_GetComponentCount(g.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (g *TGroupBox) ComponentIndex() int32 {
    return GroupBox_GetComponentIndex(g.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (g *TGroupBox) SetComponentIndex(value int32) {
    GroupBox_SetComponentIndex(g.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (g *TGroupBox) Owner() *TComponent {
    return ComponentFromInst(GroupBox_GetOwner(g.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (g *TGroupBox) Name() string {
    return GroupBox_GetName(g.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (g *TGroupBox) SetName(value string) {
    GroupBox_SetName(g.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (g *TGroupBox) Tag() int {
    return GroupBox_GetTag(g.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (g *TGroupBox) SetTag(value int) {
    GroupBox_SetTag(g.instance, value)
}

// Controls
func (g *TGroupBox) Controls(Index int32) *TControl {
    return ControlFromInst(GroupBox_GetControls(g.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (g *TGroupBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(GroupBox_GetComponents(g.instance, AIndex))
}

