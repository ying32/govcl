
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

type TUpDown struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewUpDown
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewUpDown(owner IComponent) *TUpDown {
    u := new(TUpDown)
    u.instance = UpDown_Create(CheckPtr(owner))
    u.ptr = unsafe.Pointer(u.instance)
    return u
}

// UpDownFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func UpDownFromInst(inst uintptr) *TUpDown {
    u := new(TUpDown)
    u.instance = inst
    u.ptr = unsafe.Pointer(inst)
    return u
}

// UpDownFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func UpDownFromObj(obj IObject) *TUpDown {
    u := new(TUpDown)
    u.instance = CheckPtr(obj)
    u.ptr = unsafe.Pointer(u.instance)
    return u
}

// UpDownFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func UpDownFromUnsafePointer(ptr unsafe.Pointer) *TUpDown {
    u := new(TUpDown)
    u.instance = uintptr(ptr)
    u.ptr = ptr
    return u
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (u *TUpDown) Free() {
    if u.instance != 0 {
        UpDown_Free(u.instance)
        u.instance = 0
        u.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (u *TUpDown) Instance() uintptr {
    return u.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (u *TUpDown) UnsafeAddr() unsafe.Pointer {
    return u.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (u *TUpDown) IsValid() bool {
    return u.instance != 0
}

// TUpDownClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TUpDownClass() TClass {
    return UpDown_StaticClassType()
}

// CanFocus
func (u *TUpDown) CanFocus() bool {
    return UpDown_CanFocus(u.instance)
}

// FlipChildren
func (u *TUpDown) FlipChildren(AllLevels bool) {
    UpDown_FlipChildren(u.instance, AllLevels)
}

// Focused
func (u *TUpDown) Focused() bool {
    return UpDown_Focused(u.instance)
}

// HandleAllocated
func (u *TUpDown) HandleAllocated() bool {
    return UpDown_HandleAllocated(u.instance)
}

// Invalidate
func (u *TUpDown) Invalidate() {
    UpDown_Invalidate(u.instance)
}

// Realign
func (u *TUpDown) Realign() {
    UpDown_Realign(u.instance)
}

// Repaint
func (u *TUpDown) Repaint() {
    UpDown_Repaint(u.instance)
}

// ScaleBy
func (u *TUpDown) ScaleBy(M int32, D int32) {
    UpDown_ScaleBy(u.instance, M , D)
}

// SetBounds
func (u *TUpDown) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    UpDown_SetBounds(u.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (u *TUpDown) SetFocus() {
    UpDown_SetFocus(u.instance)
}

// Update
func (u *TUpDown) Update() {
    UpDown_Update(u.instance)
}

// BringToFront
func (u *TUpDown) BringToFront() {
    UpDown_BringToFront(u.instance)
}

// ClientToScreen
func (u *TUpDown) ClientToScreen(Point TPoint) TPoint {
    return UpDown_ClientToScreen(u.instance, Point)
}

// ClientToParent
func (u *TUpDown) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return UpDown_ClientToParent(u.instance, Point , CheckPtr(AParent))
}

// Dragging
func (u *TUpDown) Dragging() bool {
    return UpDown_Dragging(u.instance)
}

// HasParent
func (u *TUpDown) HasParent() bool {
    return UpDown_HasParent(u.instance)
}

// Hide
func (u *TUpDown) Hide() {
    UpDown_Hide(u.instance)
}

// Perform
func (u *TUpDown) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return UpDown_Perform(u.instance, Msg , WParam , LParam)
}

// Refresh
func (u *TUpDown) Refresh() {
    UpDown_Refresh(u.instance)
}

// ScreenToClient
func (u *TUpDown) ScreenToClient(Point TPoint) TPoint {
    return UpDown_ScreenToClient(u.instance, Point)
}

// ParentToClient
func (u *TUpDown) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return UpDown_ParentToClient(u.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (u *TUpDown) SendToBack() {
    UpDown_SendToBack(u.instance)
}

// Show
func (u *TUpDown) Show() {
    UpDown_Show(u.instance)
}

// GetTextBuf
func (u *TUpDown) GetTextBuf(Buffer string, BufSize int32) int32 {
    return UpDown_GetTextBuf(u.instance, Buffer , BufSize)
}

// GetTextLen
func (u *TUpDown) GetTextLen() int32 {
    return UpDown_GetTextLen(u.instance)
}

// FindComponent
func (u *TUpDown) FindComponent(AName string) *TComponent {
    return ComponentFromInst(UpDown_FindComponent(u.instance, AName))
}

// GetNamePath
func (u *TUpDown) GetNamePath() string {
    return UpDown_GetNamePath(u.instance)
}

// Assign
func (u *TUpDown) Assign(Source IObject) {
    UpDown_Assign(u.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (u *TUpDown) DisposeOf() {
    UpDown_DisposeOf(u.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (u *TUpDown) ClassType() TClass {
    return UpDown_ClassType(u.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (u *TUpDown) ClassName() string {
    return UpDown_ClassName(u.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (u *TUpDown) InstanceSize() int32 {
    return UpDown_InstanceSize(u.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (u *TUpDown) InheritsFrom(AClass TClass) bool {
    return UpDown_InheritsFrom(u.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (u *TUpDown) Equals(Obj IObject) bool {
    return UpDown_Equals(u.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (u *TUpDown) GetHashCode() int32 {
    return UpDown_GetHashCode(u.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (u *TUpDown) ToString() string {
    return UpDown_ToString(u.instance)
}

// Anchors
func (u *TUpDown) Anchors() TAnchors {
    return UpDown_GetAnchors(u.instance)
}

// SetAnchors
func (u *TUpDown) SetAnchors(value TAnchors) {
    UpDown_SetAnchors(u.instance, value)
}

// DoubleBuffered
func (u *TUpDown) DoubleBuffered() bool {
    return UpDown_GetDoubleBuffered(u.instance)
}

// SetDoubleBuffered
func (u *TUpDown) SetDoubleBuffered(value bool) {
    UpDown_SetDoubleBuffered(u.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (u *TUpDown) Enabled() bool {
    return UpDown_GetEnabled(u.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (u *TUpDown) SetEnabled(value bool) {
    UpDown_SetEnabled(u.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (u *TUpDown) Hint() string {
    return UpDown_GetHint(u.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (u *TUpDown) SetHint(value string) {
    UpDown_SetHint(u.instance, value)
}

// Min
func (u *TUpDown) Min() int32 {
    return UpDown_GetMin(u.instance)
}

// SetMin
func (u *TUpDown) SetMin(value int32) {
    UpDown_SetMin(u.instance, value)
}

// Max
func (u *TUpDown) Max() int32 {
    return UpDown_GetMax(u.instance)
}

// SetMax
func (u *TUpDown) SetMax(value int32) {
    UpDown_SetMax(u.instance, value)
}

// Orientation
func (u *TUpDown) Orientation() TUDOrientation {
    return UpDown_GetOrientation(u.instance)
}

// SetOrientation
func (u *TUpDown) SetOrientation(value TUDOrientation) {
    UpDown_SetOrientation(u.instance, value)
}

// ParentDoubleBuffered
func (u *TUpDown) ParentDoubleBuffered() bool {
    return UpDown_GetParentDoubleBuffered(u.instance)
}

// SetParentDoubleBuffered
func (u *TUpDown) SetParentDoubleBuffered(value bool) {
    UpDown_SetParentDoubleBuffered(u.instance, value)
}

// ParentShowHint
func (u *TUpDown) ParentShowHint() bool {
    return UpDown_GetParentShowHint(u.instance)
}

// SetParentShowHint
func (u *TUpDown) SetParentShowHint(value bool) {
    UpDown_SetParentShowHint(u.instance, value)
}

// PopupMenu
func (u *TUpDown) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(UpDown_GetPopupMenu(u.instance))
}

// SetPopupMenu
func (u *TUpDown) SetPopupMenu(value IComponent) {
    UpDown_SetPopupMenu(u.instance, CheckPtr(value))
}

// Position
func (u *TUpDown) Position() int32 {
    return UpDown_GetPosition(u.instance)
}

// SetPosition
func (u *TUpDown) SetPosition(value int32) {
    UpDown_SetPosition(u.instance, value)
}

// ShowHint
func (u *TUpDown) ShowHint() bool {
    return UpDown_GetShowHint(u.instance)
}

// SetShowHint
func (u *TUpDown) SetShowHint(value bool) {
    UpDown_SetShowHint(u.instance, value)
}

// TabOrder
func (u *TUpDown) TabOrder() uint16 {
    return UpDown_GetTabOrder(u.instance)
}

// SetTabOrder
func (u *TUpDown) SetTabOrder(value uint16) {
    UpDown_SetTabOrder(u.instance, value)
}

// TabStop
func (u *TUpDown) TabStop() bool {
    return UpDown_GetTabStop(u.instance)
}

// SetTabStop
func (u *TUpDown) SetTabStop(value bool) {
    UpDown_SetTabStop(u.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (u *TUpDown) Visible() bool {
    return UpDown_GetVisible(u.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (u *TUpDown) SetVisible(value bool) {
    UpDown_SetVisible(u.instance, value)
}

// Wrap
func (u *TUpDown) Wrap() bool {
    return UpDown_GetWrap(u.instance)
}

// SetWrap
func (u *TUpDown) SetWrap(value bool) {
    UpDown_SetWrap(u.instance, value)
}

// StyleElements
func (u *TUpDown) StyleElements() TStyleElements {
    return UpDown_GetStyleElements(u.instance)
}

// SetStyleElements
func (u *TUpDown) SetStyleElements(value TStyleElements) {
    UpDown_SetStyleElements(u.instance, value)
}

// SetOnContextPopup
func (u *TUpDown) SetOnContextPopup(fn TContextPopupEvent) {
    UpDown_SetOnContextPopup(u.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (u *TUpDown) SetOnClick(fn TUDClickEvent) {
    UpDown_SetOnClick(u.instance, fn)
}

// SetOnEnter
func (u *TUpDown) SetOnEnter(fn TNotifyEvent) {
    UpDown_SetOnEnter(u.instance, fn)
}

// SetOnExit
func (u *TUpDown) SetOnExit(fn TNotifyEvent) {
    UpDown_SetOnExit(u.instance, fn)
}

// SetOnMouseDown
func (u *TUpDown) SetOnMouseDown(fn TMouseEvent) {
    UpDown_SetOnMouseDown(u.instance, fn)
}

// SetOnMouseEnter
func (u *TUpDown) SetOnMouseEnter(fn TNotifyEvent) {
    UpDown_SetOnMouseEnter(u.instance, fn)
}

// SetOnMouseLeave
func (u *TUpDown) SetOnMouseLeave(fn TNotifyEvent) {
    UpDown_SetOnMouseLeave(u.instance, fn)
}

// SetOnMouseMove
func (u *TUpDown) SetOnMouseMove(fn TMouseMoveEvent) {
    UpDown_SetOnMouseMove(u.instance, fn)
}

// SetOnMouseUp
func (u *TUpDown) SetOnMouseUp(fn TMouseEvent) {
    UpDown_SetOnMouseUp(u.instance, fn)
}

// DockSite
func (u *TUpDown) DockSite() bool {
    return UpDown_GetDockSite(u.instance)
}

// SetDockSite
func (u *TUpDown) SetDockSite(value bool) {
    UpDown_SetDockSite(u.instance, value)
}

// Brush
func (u *TUpDown) Brush() *TBrush {
    return BrushFromInst(UpDown_GetBrush(u.instance))
}

// ControlCount
func (u *TUpDown) ControlCount() int32 {
    return UpDown_GetControlCount(u.instance)
}

// Handle
func (u *TUpDown) Handle() HWND {
    return UpDown_GetHandle(u.instance)
}

// ParentWindow
func (u *TUpDown) ParentWindow() HWND {
    return UpDown_GetParentWindow(u.instance)
}

// SetParentWindow
func (u *TUpDown) SetParentWindow(value HWND) {
    UpDown_SetParentWindow(u.instance, value)
}

// UseDockManager
func (u *TUpDown) UseDockManager() bool {
    return UpDown_GetUseDockManager(u.instance)
}

// SetUseDockManager
func (u *TUpDown) SetUseDockManager(value bool) {
    UpDown_SetUseDockManager(u.instance, value)
}

// Action
func (u *TUpDown) Action() *TAction {
    return ActionFromInst(UpDown_GetAction(u.instance))
}

// SetAction
func (u *TUpDown) SetAction(value IComponent) {
    UpDown_SetAction(u.instance, CheckPtr(value))
}

// Align
func (u *TUpDown) Align() TAlign {
    return UpDown_GetAlign(u.instance)
}

// SetAlign
func (u *TUpDown) SetAlign(value TAlign) {
    UpDown_SetAlign(u.instance, value)
}

// BiDiMode
func (u *TUpDown) BiDiMode() TBiDiMode {
    return UpDown_GetBiDiMode(u.instance)
}

// SetBiDiMode
func (u *TUpDown) SetBiDiMode(value TBiDiMode) {
    UpDown_SetBiDiMode(u.instance, value)
}

// BoundsRect
func (u *TUpDown) BoundsRect() TRect {
    return UpDown_GetBoundsRect(u.instance)
}

// SetBoundsRect
func (u *TUpDown) SetBoundsRect(value TRect) {
    UpDown_SetBoundsRect(u.instance, value)
}

// ClientHeight
func (u *TUpDown) ClientHeight() int32 {
    return UpDown_GetClientHeight(u.instance)
}

// SetClientHeight
func (u *TUpDown) SetClientHeight(value int32) {
    UpDown_SetClientHeight(u.instance, value)
}

// ClientRect
func (u *TUpDown) ClientRect() TRect {
    return UpDown_GetClientRect(u.instance)
}

// ClientWidth
func (u *TUpDown) ClientWidth() int32 {
    return UpDown_GetClientWidth(u.instance)
}

// SetClientWidth
func (u *TUpDown) SetClientWidth(value int32) {
    UpDown_SetClientWidth(u.instance, value)
}

// ExplicitLeft
func (u *TUpDown) ExplicitLeft() int32 {
    return UpDown_GetExplicitLeft(u.instance)
}

// ExplicitTop
func (u *TUpDown) ExplicitTop() int32 {
    return UpDown_GetExplicitTop(u.instance)
}

// ExplicitWidth
func (u *TUpDown) ExplicitWidth() int32 {
    return UpDown_GetExplicitWidth(u.instance)
}

// ExplicitHeight
func (u *TUpDown) ExplicitHeight() int32 {
    return UpDown_GetExplicitHeight(u.instance)
}

// Floating
func (u *TUpDown) Floating() bool {
    return UpDown_GetFloating(u.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (u *TUpDown) Parent() *TWinControl {
    return WinControlFromInst(UpDown_GetParent(u.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (u *TUpDown) SetParent(value IWinControl) {
    UpDown_SetParent(u.instance, CheckPtr(value))
}

// AlignWithMargins
func (u *TUpDown) AlignWithMargins() bool {
    return UpDown_GetAlignWithMargins(u.instance)
}

// SetAlignWithMargins
func (u *TUpDown) SetAlignWithMargins(value bool) {
    UpDown_SetAlignWithMargins(u.instance, value)
}

// Left
func (u *TUpDown) Left() int32 {
    return UpDown_GetLeft(u.instance)
}

// SetLeft
func (u *TUpDown) SetLeft(value int32) {
    UpDown_SetLeft(u.instance, value)
}

// Top
func (u *TUpDown) Top() int32 {
    return UpDown_GetTop(u.instance)
}

// SetTop
func (u *TUpDown) SetTop(value int32) {
    UpDown_SetTop(u.instance, value)
}

// Width
func (u *TUpDown) Width() int32 {
    return UpDown_GetWidth(u.instance)
}

// SetWidth
func (u *TUpDown) SetWidth(value int32) {
    UpDown_SetWidth(u.instance, value)
}

// Height
func (u *TUpDown) Height() int32 {
    return UpDown_GetHeight(u.instance)
}

// SetHeight
func (u *TUpDown) SetHeight(value int32) {
    UpDown_SetHeight(u.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (u *TUpDown) Cursor() TCursor {
    return UpDown_GetCursor(u.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (u *TUpDown) SetCursor(value TCursor) {
    UpDown_SetCursor(u.instance, value)
}

// Margins
func (u *TUpDown) Margins() *TMargins {
    return MarginsFromInst(UpDown_GetMargins(u.instance))
}

// SetMargins
func (u *TUpDown) SetMargins(value *TMargins) {
    UpDown_SetMargins(u.instance, CheckPtr(value))
}

// CustomHint
func (u *TUpDown) CustomHint() *TCustomHint {
    return CustomHintFromInst(UpDown_GetCustomHint(u.instance))
}

// SetCustomHint
func (u *TUpDown) SetCustomHint(value IComponent) {
    UpDown_SetCustomHint(u.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (u *TUpDown) ComponentCount() int32 {
    return UpDown_GetComponentCount(u.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (u *TUpDown) ComponentIndex() int32 {
    return UpDown_GetComponentIndex(u.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (u *TUpDown) SetComponentIndex(value int32) {
    UpDown_SetComponentIndex(u.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (u *TUpDown) Owner() *TComponent {
    return ComponentFromInst(UpDown_GetOwner(u.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (u *TUpDown) Name() string {
    return UpDown_GetName(u.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (u *TUpDown) SetName(value string) {
    UpDown_SetName(u.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (u *TUpDown) Tag() int {
    return UpDown_GetTag(u.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (u *TUpDown) SetTag(value int) {
    UpDown_SetTag(u.instance, value)
}

// Controls
func (u *TUpDown) Controls(Index int32) *TControl {
    return ControlFromInst(UpDown_GetControls(u.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (u *TUpDown) Components(AIndex int32) *TComponent {
    return ComponentFromInst(UpDown_GetComponents(u.instance, AIndex))
}

