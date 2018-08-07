
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

type THotKey struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewHotKey
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewHotKey(owner IComponent) *THotKey {
    h := new(THotKey)
    h.instance = HotKey_Create(CheckPtr(owner))
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HotKeyFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func HotKeyFromInst(inst uintptr) *THotKey {
    h := new(THotKey)
    h.instance = inst
    h.ptr = unsafe.Pointer(inst)
    return h
}

// HotKeyFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func HotKeyFromObj(obj IObject) *THotKey {
    h := new(THotKey)
    h.instance = CheckPtr(obj)
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HotKeyFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func HotKeyFromUnsafePointer(ptr unsafe.Pointer) *THotKey {
    h := new(THotKey)
    h.instance = uintptr(ptr)
    h.ptr = ptr
    return h
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (h *THotKey) Free() {
    if h.instance != 0 {
        HotKey_Free(h.instance)
        h.instance = 0
        h.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (h *THotKey) Instance() uintptr {
    return h.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (h *THotKey) UnsafeAddr() unsafe.Pointer {
    return h.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (h *THotKey) IsValid() bool {
    return h.instance != 0
}

// THotKeyClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func THotKeyClass() TClass {
    return HotKey_StaticClassType()
}

// CanFocus
func (h *THotKey) CanFocus() bool {
    return HotKey_CanFocus(h.instance)
}

// FlipChildren
func (h *THotKey) FlipChildren(AllLevels bool) {
    HotKey_FlipChildren(h.instance, AllLevels)
}

// Focused
func (h *THotKey) Focused() bool {
    return HotKey_Focused(h.instance)
}

// HandleAllocated
func (h *THotKey) HandleAllocated() bool {
    return HotKey_HandleAllocated(h.instance)
}

// Invalidate
func (h *THotKey) Invalidate() {
    HotKey_Invalidate(h.instance)
}

// Realign
func (h *THotKey) Realign() {
    HotKey_Realign(h.instance)
}

// Repaint
func (h *THotKey) Repaint() {
    HotKey_Repaint(h.instance)
}

// ScaleBy
func (h *THotKey) ScaleBy(M int32, D int32) {
    HotKey_ScaleBy(h.instance, M , D)
}

// SetBounds
func (h *THotKey) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    HotKey_SetBounds(h.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (h *THotKey) SetFocus() {
    HotKey_SetFocus(h.instance)
}

// Update
func (h *THotKey) Update() {
    HotKey_Update(h.instance)
}

// BringToFront
func (h *THotKey) BringToFront() {
    HotKey_BringToFront(h.instance)
}

// ClientToScreen
func (h *THotKey) ClientToScreen(Point TPoint) TPoint {
    return HotKey_ClientToScreen(h.instance, Point)
}

// ClientToParent
func (h *THotKey) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return HotKey_ClientToParent(h.instance, Point , CheckPtr(AParent))
}

// Dragging
func (h *THotKey) Dragging() bool {
    return HotKey_Dragging(h.instance)
}

// HasParent
func (h *THotKey) HasParent() bool {
    return HotKey_HasParent(h.instance)
}

// Hide
func (h *THotKey) Hide() {
    HotKey_Hide(h.instance)
}

// Perform
func (h *THotKey) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return HotKey_Perform(h.instance, Msg , WParam , LParam)
}

// Refresh
func (h *THotKey) Refresh() {
    HotKey_Refresh(h.instance)
}

// ScreenToClient
func (h *THotKey) ScreenToClient(Point TPoint) TPoint {
    return HotKey_ScreenToClient(h.instance, Point)
}

// ParentToClient
func (h *THotKey) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return HotKey_ParentToClient(h.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (h *THotKey) SendToBack() {
    HotKey_SendToBack(h.instance)
}

// Show
func (h *THotKey) Show() {
    HotKey_Show(h.instance)
}

// GetTextBuf
func (h *THotKey) GetTextBuf(Buffer string, BufSize int32) int32 {
    return HotKey_GetTextBuf(h.instance, Buffer , BufSize)
}

// GetTextLen
func (h *THotKey) GetTextLen() int32 {
    return HotKey_GetTextLen(h.instance)
}

// FindComponent
func (h *THotKey) FindComponent(AName string) *TComponent {
    return ComponentFromInst(HotKey_FindComponent(h.instance, AName))
}

// GetNamePath
func (h *THotKey) GetNamePath() string {
    return HotKey_GetNamePath(h.instance)
}

// Assign
func (h *THotKey) Assign(Source IObject) {
    HotKey_Assign(h.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (h *THotKey) DisposeOf() {
    HotKey_DisposeOf(h.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (h *THotKey) ClassType() TClass {
    return HotKey_ClassType(h.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (h *THotKey) ClassName() string {
    return HotKey_ClassName(h.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (h *THotKey) InstanceSize() int32 {
    return HotKey_InstanceSize(h.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (h *THotKey) InheritsFrom(AClass TClass) bool {
    return HotKey_InheritsFrom(h.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (h *THotKey) Equals(Obj IObject) bool {
    return HotKey_Equals(h.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (h *THotKey) GetHashCode() int32 {
    return HotKey_GetHashCode(h.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (h *THotKey) ToString() string {
    return HotKey_ToString(h.instance)
}

// Anchors
func (h *THotKey) Anchors() TAnchors {
    return HotKey_GetAnchors(h.instance)
}

// SetAnchors
func (h *THotKey) SetAnchors(value TAnchors) {
    HotKey_SetAnchors(h.instance, value)
}

// AutoSize
func (h *THotKey) AutoSize() bool {
    return HotKey_GetAutoSize(h.instance)
}

// SetAutoSize
func (h *THotKey) SetAutoSize(value bool) {
    HotKey_SetAutoSize(h.instance, value)
}

// BiDiMode
func (h *THotKey) BiDiMode() TBiDiMode {
    return HotKey_GetBiDiMode(h.instance)
}

// SetBiDiMode
func (h *THotKey) SetBiDiMode(value TBiDiMode) {
    HotKey_SetBiDiMode(h.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (h *THotKey) Enabled() bool {
    return HotKey_GetEnabled(h.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (h *THotKey) SetEnabled(value bool) {
    HotKey_SetEnabled(h.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (h *THotKey) Hint() string {
    return HotKey_GetHint(h.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (h *THotKey) SetHint(value string) {
    HotKey_SetHint(h.instance, value)
}

// ParentShowHint
func (h *THotKey) ParentShowHint() bool {
    return HotKey_GetParentShowHint(h.instance)
}

// SetParentShowHint
func (h *THotKey) SetParentShowHint(value bool) {
    HotKey_SetParentShowHint(h.instance, value)
}

// PopupMenu
func (h *THotKey) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(HotKey_GetPopupMenu(h.instance))
}

// SetPopupMenu
func (h *THotKey) SetPopupMenu(value IComponent) {
    HotKey_SetPopupMenu(h.instance, CheckPtr(value))
}

// ShowHint
func (h *THotKey) ShowHint() bool {
    return HotKey_GetShowHint(h.instance)
}

// SetShowHint
func (h *THotKey) SetShowHint(value bool) {
    HotKey_SetShowHint(h.instance, value)
}

// TabOrder
func (h *THotKey) TabOrder() uint16 {
    return HotKey_GetTabOrder(h.instance)
}

// SetTabOrder
func (h *THotKey) SetTabOrder(value uint16) {
    HotKey_SetTabOrder(h.instance, value)
}

// TabStop
func (h *THotKey) TabStop() bool {
    return HotKey_GetTabStop(h.instance)
}

// SetTabStop
func (h *THotKey) SetTabStop(value bool) {
    HotKey_SetTabStop(h.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (h *THotKey) Visible() bool {
    return HotKey_GetVisible(h.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (h *THotKey) SetVisible(value bool) {
    HotKey_SetVisible(h.instance, value)
}

// StyleElements
func (h *THotKey) StyleElements() TStyleElements {
    return HotKey_GetStyleElements(h.instance)
}

// SetStyleElements
func (h *THotKey) SetStyleElements(value TStyleElements) {
    HotKey_SetStyleElements(h.instance, value)
}

// SetOnChange
func (h *THotKey) SetOnChange(fn TNotifyEvent) {
    HotKey_SetOnChange(h.instance, fn)
}

// SetOnContextPopup
func (h *THotKey) SetOnContextPopup(fn TContextPopupEvent) {
    HotKey_SetOnContextPopup(h.instance, fn)
}

// SetOnEnter
func (h *THotKey) SetOnEnter(fn TNotifyEvent) {
    HotKey_SetOnEnter(h.instance, fn)
}

// SetOnExit
func (h *THotKey) SetOnExit(fn TNotifyEvent) {
    HotKey_SetOnExit(h.instance, fn)
}

// SetOnMouseDown
func (h *THotKey) SetOnMouseDown(fn TMouseEvent) {
    HotKey_SetOnMouseDown(h.instance, fn)
}

// SetOnMouseEnter
func (h *THotKey) SetOnMouseEnter(fn TNotifyEvent) {
    HotKey_SetOnMouseEnter(h.instance, fn)
}

// SetOnMouseLeave
func (h *THotKey) SetOnMouseLeave(fn TNotifyEvent) {
    HotKey_SetOnMouseLeave(h.instance, fn)
}

// SetOnMouseMove
func (h *THotKey) SetOnMouseMove(fn TMouseMoveEvent) {
    HotKey_SetOnMouseMove(h.instance, fn)
}

// SetOnMouseUp
func (h *THotKey) SetOnMouseUp(fn TMouseEvent) {
    HotKey_SetOnMouseUp(h.instance, fn)
}

// DockSite
func (h *THotKey) DockSite() bool {
    return HotKey_GetDockSite(h.instance)
}

// SetDockSite
func (h *THotKey) SetDockSite(value bool) {
    HotKey_SetDockSite(h.instance, value)
}

// DoubleBuffered
func (h *THotKey) DoubleBuffered() bool {
    return HotKey_GetDoubleBuffered(h.instance)
}

// SetDoubleBuffered
func (h *THotKey) SetDoubleBuffered(value bool) {
    HotKey_SetDoubleBuffered(h.instance, value)
}

// Brush
func (h *THotKey) Brush() *TBrush {
    return BrushFromInst(HotKey_GetBrush(h.instance))
}

// ControlCount
func (h *THotKey) ControlCount() int32 {
    return HotKey_GetControlCount(h.instance)
}

// Handle
func (h *THotKey) Handle() HWND {
    return HotKey_GetHandle(h.instance)
}

// ParentDoubleBuffered
func (h *THotKey) ParentDoubleBuffered() bool {
    return HotKey_GetParentDoubleBuffered(h.instance)
}

// SetParentDoubleBuffered
func (h *THotKey) SetParentDoubleBuffered(value bool) {
    HotKey_SetParentDoubleBuffered(h.instance, value)
}

// ParentWindow
func (h *THotKey) ParentWindow() HWND {
    return HotKey_GetParentWindow(h.instance)
}

// SetParentWindow
func (h *THotKey) SetParentWindow(value HWND) {
    HotKey_SetParentWindow(h.instance, value)
}

// UseDockManager
func (h *THotKey) UseDockManager() bool {
    return HotKey_GetUseDockManager(h.instance)
}

// SetUseDockManager
func (h *THotKey) SetUseDockManager(value bool) {
    HotKey_SetUseDockManager(h.instance, value)
}

// Action
func (h *THotKey) Action() *TAction {
    return ActionFromInst(HotKey_GetAction(h.instance))
}

// SetAction
func (h *THotKey) SetAction(value IComponent) {
    HotKey_SetAction(h.instance, CheckPtr(value))
}

// Align
func (h *THotKey) Align() TAlign {
    return HotKey_GetAlign(h.instance)
}

// SetAlign
func (h *THotKey) SetAlign(value TAlign) {
    HotKey_SetAlign(h.instance, value)
}

// BoundsRect
func (h *THotKey) BoundsRect() TRect {
    return HotKey_GetBoundsRect(h.instance)
}

// SetBoundsRect
func (h *THotKey) SetBoundsRect(value TRect) {
    HotKey_SetBoundsRect(h.instance, value)
}

// ClientHeight
func (h *THotKey) ClientHeight() int32 {
    return HotKey_GetClientHeight(h.instance)
}

// SetClientHeight
func (h *THotKey) SetClientHeight(value int32) {
    HotKey_SetClientHeight(h.instance, value)
}

// ClientRect
func (h *THotKey) ClientRect() TRect {
    return HotKey_GetClientRect(h.instance)
}

// ClientWidth
func (h *THotKey) ClientWidth() int32 {
    return HotKey_GetClientWidth(h.instance)
}

// SetClientWidth
func (h *THotKey) SetClientWidth(value int32) {
    HotKey_SetClientWidth(h.instance, value)
}

// ExplicitLeft
func (h *THotKey) ExplicitLeft() int32 {
    return HotKey_GetExplicitLeft(h.instance)
}

// ExplicitTop
func (h *THotKey) ExplicitTop() int32 {
    return HotKey_GetExplicitTop(h.instance)
}

// ExplicitWidth
func (h *THotKey) ExplicitWidth() int32 {
    return HotKey_GetExplicitWidth(h.instance)
}

// ExplicitHeight
func (h *THotKey) ExplicitHeight() int32 {
    return HotKey_GetExplicitHeight(h.instance)
}

// Floating
func (h *THotKey) Floating() bool {
    return HotKey_GetFloating(h.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (h *THotKey) Parent() *TWinControl {
    return WinControlFromInst(HotKey_GetParent(h.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (h *THotKey) SetParent(value IWinControl) {
    HotKey_SetParent(h.instance, CheckPtr(value))
}

// AlignWithMargins
func (h *THotKey) AlignWithMargins() bool {
    return HotKey_GetAlignWithMargins(h.instance)
}

// SetAlignWithMargins
func (h *THotKey) SetAlignWithMargins(value bool) {
    HotKey_SetAlignWithMargins(h.instance, value)
}

// Left
func (h *THotKey) Left() int32 {
    return HotKey_GetLeft(h.instance)
}

// SetLeft
func (h *THotKey) SetLeft(value int32) {
    HotKey_SetLeft(h.instance, value)
}

// Top
func (h *THotKey) Top() int32 {
    return HotKey_GetTop(h.instance)
}

// SetTop
func (h *THotKey) SetTop(value int32) {
    HotKey_SetTop(h.instance, value)
}

// Width
func (h *THotKey) Width() int32 {
    return HotKey_GetWidth(h.instance)
}

// SetWidth
func (h *THotKey) SetWidth(value int32) {
    HotKey_SetWidth(h.instance, value)
}

// Height
func (h *THotKey) Height() int32 {
    return HotKey_GetHeight(h.instance)
}

// SetHeight
func (h *THotKey) SetHeight(value int32) {
    HotKey_SetHeight(h.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (h *THotKey) Cursor() TCursor {
    return HotKey_GetCursor(h.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (h *THotKey) SetCursor(value TCursor) {
    HotKey_SetCursor(h.instance, value)
}

// Margins
func (h *THotKey) Margins() *TMargins {
    return MarginsFromInst(HotKey_GetMargins(h.instance))
}

// SetMargins
func (h *THotKey) SetMargins(value *TMargins) {
    HotKey_SetMargins(h.instance, CheckPtr(value))
}

// CustomHint
func (h *THotKey) CustomHint() *TCustomHint {
    return CustomHintFromInst(HotKey_GetCustomHint(h.instance))
}

// SetCustomHint
func (h *THotKey) SetCustomHint(value IComponent) {
    HotKey_SetCustomHint(h.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (h *THotKey) ComponentCount() int32 {
    return HotKey_GetComponentCount(h.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (h *THotKey) ComponentIndex() int32 {
    return HotKey_GetComponentIndex(h.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (h *THotKey) SetComponentIndex(value int32) {
    HotKey_SetComponentIndex(h.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (h *THotKey) Owner() *TComponent {
    return ComponentFromInst(HotKey_GetOwner(h.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (h *THotKey) Name() string {
    return HotKey_GetName(h.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (h *THotKey) SetName(value string) {
    HotKey_SetName(h.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (h *THotKey) Tag() int {
    return HotKey_GetTag(h.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (h *THotKey) SetTag(value int) {
    HotKey_SetTag(h.instance, value)
}

// Controls
func (h *THotKey) Controls(Index int32) *TControl {
    return ControlFromInst(HotKey_GetControls(h.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (h *THotKey) Components(AIndex int32) *TComponent {
    return ComponentFromInst(HotKey_GetComponents(h.instance, AIndex))
}

