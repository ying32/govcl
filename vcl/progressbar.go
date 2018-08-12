
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

type TProgressBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewProgressBar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewProgressBar(owner IComponent) *TProgressBar {
    p := new(TProgressBar)
    p.instance = ProgressBar_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// ProgressBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ProgressBarFromInst(inst uintptr) *TProgressBar {
    p := new(TProgressBar)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// ProgressBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ProgressBarFromObj(obj IObject) *TProgressBar {
    p := new(TProgressBar)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// ProgressBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ProgressBarFromUnsafePointer(ptr unsafe.Pointer) *TProgressBar {
    p := new(TProgressBar)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TProgressBar) Free() {
    if p.instance != 0 {
        ProgressBar_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TProgressBar) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TProgressBar) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TProgressBar) IsValid() bool {
    return p.instance != 0
}

// TProgressBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TProgressBarClass() TClass {
    return ProgressBar_StaticClassType()
}

// StepIt
func (p *TProgressBar) StepIt() {
    ProgressBar_StepIt(p.instance)
}

// StepBy
func (p *TProgressBar) StepBy(Delta int32) {
    ProgressBar_StepBy(p.instance, Delta)
}

// CanFocus
func (p *TProgressBar) CanFocus() bool {
    return ProgressBar_CanFocus(p.instance)
}

// ContainsControl
func (p *TProgressBar) ContainsControl(Control IControl) bool {
    return ProgressBar_ContainsControl(p.instance, CheckPtr(Control))
}

// ControlAtPos
func (p *TProgressBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(ProgressBar_ControlAtPos(p.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (p *TProgressBar) DisableAlign() {
    ProgressBar_DisableAlign(p.instance)
}

// EnableAlign
func (p *TProgressBar) EnableAlign() {
    ProgressBar_EnableAlign(p.instance)
}

// FindChildControl
func (p *TProgressBar) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(ProgressBar_FindChildControl(p.instance, ControlName))
}

// FlipChildren
func (p *TProgressBar) FlipChildren(AllLevels bool) {
    ProgressBar_FlipChildren(p.instance, AllLevels)
}

// Focused
func (p *TProgressBar) Focused() bool {
    return ProgressBar_Focused(p.instance)
}

// HandleAllocated
func (p *TProgressBar) HandleAllocated() bool {
    return ProgressBar_HandleAllocated(p.instance)
}

// InsertControl
func (p *TProgressBar) InsertControl(AControl IControl) {
    ProgressBar_InsertControl(p.instance, CheckPtr(AControl))
}

// Invalidate
func (p *TProgressBar) Invalidate() {
    ProgressBar_Invalidate(p.instance)
}

// PaintTo
func (p *TProgressBar) PaintTo(DC HDC, X int32, Y int32) {
    ProgressBar_PaintTo(p.instance, DC , X , Y)
}

// RemoveControl
func (p *TProgressBar) RemoveControl(AControl IControl) {
    ProgressBar_RemoveControl(p.instance, CheckPtr(AControl))
}

// Realign
func (p *TProgressBar) Realign() {
    ProgressBar_Realign(p.instance)
}

// Repaint
func (p *TProgressBar) Repaint() {
    ProgressBar_Repaint(p.instance)
}

// ScaleBy
func (p *TProgressBar) ScaleBy(M int32, D int32) {
    ProgressBar_ScaleBy(p.instance, M , D)
}

// ScrollBy
func (p *TProgressBar) ScrollBy(DeltaX int32, DeltaY int32) {
    ProgressBar_ScrollBy(p.instance, DeltaX , DeltaY)
}

// SetBounds
func (p *TProgressBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ProgressBar_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (p *TProgressBar) SetFocus() {
    ProgressBar_SetFocus(p.instance)
}

// Update
func (p *TProgressBar) Update() {
    ProgressBar_Update(p.instance)
}

// UpdateControlState
func (p *TProgressBar) UpdateControlState() {
    ProgressBar_UpdateControlState(p.instance)
}

// BringToFront
func (p *TProgressBar) BringToFront() {
    ProgressBar_BringToFront(p.instance)
}

// ClientToScreen
func (p *TProgressBar) ClientToScreen(Point TPoint) TPoint {
    return ProgressBar_ClientToScreen(p.instance, Point)
}

// ClientToParent
func (p *TProgressBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ProgressBar_ClientToParent(p.instance, Point , CheckPtr(AParent))
}

// Dragging
func (p *TProgressBar) Dragging() bool {
    return ProgressBar_Dragging(p.instance)
}

// HasParent
func (p *TProgressBar) HasParent() bool {
    return ProgressBar_HasParent(p.instance)
}

// Hide
func (p *TProgressBar) Hide() {
    ProgressBar_Hide(p.instance)
}

// Perform
func (p *TProgressBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ProgressBar_Perform(p.instance, Msg , WParam , LParam)
}

// Refresh
func (p *TProgressBar) Refresh() {
    ProgressBar_Refresh(p.instance)
}

// ScreenToClient
func (p *TProgressBar) ScreenToClient(Point TPoint) TPoint {
    return ProgressBar_ScreenToClient(p.instance, Point)
}

// ParentToClient
func (p *TProgressBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ProgressBar_ParentToClient(p.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (p *TProgressBar) SendToBack() {
    ProgressBar_SendToBack(p.instance)
}

// Show
func (p *TProgressBar) Show() {
    ProgressBar_Show(p.instance)
}

// GetTextBuf
func (p *TProgressBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ProgressBar_GetTextBuf(p.instance, Buffer , BufSize)
}

// GetTextLen
func (p *TProgressBar) GetTextLen() int32 {
    return ProgressBar_GetTextLen(p.instance)
}

// SetTextBuf
func (p *TProgressBar) SetTextBuf(Buffer string) {
    ProgressBar_SetTextBuf(p.instance, Buffer)
}

// FindComponent
func (p *TProgressBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ProgressBar_FindComponent(p.instance, AName))
}

// GetNamePath
func (p *TProgressBar) GetNamePath() string {
    return ProgressBar_GetNamePath(p.instance)
}

// Assign
func (p *TProgressBar) Assign(Source IObject) {
    ProgressBar_Assign(p.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TProgressBar) DisposeOf() {
    ProgressBar_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TProgressBar) ClassType() TClass {
    return ProgressBar_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TProgressBar) ClassName() string {
    return ProgressBar_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TProgressBar) InstanceSize() int32 {
    return ProgressBar_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TProgressBar) InheritsFrom(AClass TClass) bool {
    return ProgressBar_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TProgressBar) Equals(Obj IObject) bool {
    return ProgressBar_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TProgressBar) GetHashCode() int32 {
    return ProgressBar_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TProgressBar) ToString() string {
    return ProgressBar_ToString(p.instance)
}

// Align
func (p *TProgressBar) Align() TAlign {
    return ProgressBar_GetAlign(p.instance)
}

// SetAlign
func (p *TProgressBar) SetAlign(value TAlign) {
    ProgressBar_SetAlign(p.instance, value)
}

// Anchors
func (p *TProgressBar) Anchors() TAnchors {
    return ProgressBar_GetAnchors(p.instance)
}

// SetAnchors
func (p *TProgressBar) SetAnchors(value TAnchors) {
    ProgressBar_SetAnchors(p.instance, value)
}

// BorderWidth
func (p *TProgressBar) BorderWidth() int32 {
    return ProgressBar_GetBorderWidth(p.instance)
}

// SetBorderWidth
func (p *TProgressBar) SetBorderWidth(value int32) {
    ProgressBar_SetBorderWidth(p.instance, value)
}

// DoubleBuffered
func (p *TProgressBar) DoubleBuffered() bool {
    return ProgressBar_GetDoubleBuffered(p.instance)
}

// SetDoubleBuffered
func (p *TProgressBar) SetDoubleBuffered(value bool) {
    ProgressBar_SetDoubleBuffered(p.instance, value)
}

// DragCursor
func (p *TProgressBar) DragCursor() TCursor {
    return ProgressBar_GetDragCursor(p.instance)
}

// SetDragCursor
func (p *TProgressBar) SetDragCursor(value TCursor) {
    ProgressBar_SetDragCursor(p.instance, value)
}

// DragKind
func (p *TProgressBar) DragKind() TDragKind {
    return ProgressBar_GetDragKind(p.instance)
}

// SetDragKind
func (p *TProgressBar) SetDragKind(value TDragKind) {
    ProgressBar_SetDragKind(p.instance, value)
}

// DragMode
func (p *TProgressBar) DragMode() TDragMode {
    return ProgressBar_GetDragMode(p.instance)
}

// SetDragMode
func (p *TProgressBar) SetDragMode(value TDragMode) {
    ProgressBar_SetDragMode(p.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (p *TProgressBar) Enabled() bool {
    return ProgressBar_GetEnabled(p.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (p *TProgressBar) SetEnabled(value bool) {
    ProgressBar_SetEnabled(p.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (p *TProgressBar) Hint() string {
    return ProgressBar_GetHint(p.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (p *TProgressBar) SetHint(value string) {
    ProgressBar_SetHint(p.instance, value)
}

// Min
func (p *TProgressBar) Min() int32 {
    return ProgressBar_GetMin(p.instance)
}

// SetMin
func (p *TProgressBar) SetMin(value int32) {
    ProgressBar_SetMin(p.instance, value)
}

// Max
func (p *TProgressBar) Max() int32 {
    return ProgressBar_GetMax(p.instance)
}

// SetMax
func (p *TProgressBar) SetMax(value int32) {
    ProgressBar_SetMax(p.instance, value)
}

// Orientation
func (p *TProgressBar) Orientation() TProgressBarOrientation {
    return ProgressBar_GetOrientation(p.instance)
}

// SetOrientation
func (p *TProgressBar) SetOrientation(value TProgressBarOrientation) {
    ProgressBar_SetOrientation(p.instance, value)
}

// ParentDoubleBuffered
func (p *TProgressBar) ParentDoubleBuffered() bool {
    return ProgressBar_GetParentDoubleBuffered(p.instance)
}

// SetParentDoubleBuffered
func (p *TProgressBar) SetParentDoubleBuffered(value bool) {
    ProgressBar_SetParentDoubleBuffered(p.instance, value)
}

// ParentShowHint
func (p *TProgressBar) ParentShowHint() bool {
    return ProgressBar_GetParentShowHint(p.instance)
}

// SetParentShowHint
func (p *TProgressBar) SetParentShowHint(value bool) {
    ProgressBar_SetParentShowHint(p.instance, value)
}

// PopupMenu
func (p *TProgressBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ProgressBar_GetPopupMenu(p.instance))
}

// SetPopupMenu
func (p *TProgressBar) SetPopupMenu(value IComponent) {
    ProgressBar_SetPopupMenu(p.instance, CheckPtr(value))
}

// Position
func (p *TProgressBar) Position() int32 {
    return ProgressBar_GetPosition(p.instance)
}

// SetPosition
func (p *TProgressBar) SetPosition(value int32) {
    ProgressBar_SetPosition(p.instance, value)
}

// Smooth
func (p *TProgressBar) Smooth() bool {
    return ProgressBar_GetSmooth(p.instance)
}

// SetSmooth
func (p *TProgressBar) SetSmooth(value bool) {
    ProgressBar_SetSmooth(p.instance, value)
}

// Style
func (p *TProgressBar) Style() TProgressBarStyle {
    return ProgressBar_GetStyle(p.instance)
}

// SetStyle
func (p *TProgressBar) SetStyle(value TProgressBarStyle) {
    ProgressBar_SetStyle(p.instance, value)
}

// MarqueeInterval
func (p *TProgressBar) MarqueeInterval() int32 {
    return ProgressBar_GetMarqueeInterval(p.instance)
}

// SetMarqueeInterval
func (p *TProgressBar) SetMarqueeInterval(value int32) {
    ProgressBar_SetMarqueeInterval(p.instance, value)
}

// BarColor
func (p *TProgressBar) BarColor() TColor {
    return ProgressBar_GetBarColor(p.instance)
}

// SetBarColor
func (p *TProgressBar) SetBarColor(value TColor) {
    ProgressBar_SetBarColor(p.instance, value)
}

// BackgroundColor
func (p *TProgressBar) BackgroundColor() TColor {
    return ProgressBar_GetBackgroundColor(p.instance)
}

// SetBackgroundColor
func (p *TProgressBar) SetBackgroundColor(value TColor) {
    ProgressBar_SetBackgroundColor(p.instance, value)
}

// SmoothReverse
func (p *TProgressBar) SmoothReverse() bool {
    return ProgressBar_GetSmoothReverse(p.instance)
}

// SetSmoothReverse
func (p *TProgressBar) SetSmoothReverse(value bool) {
    ProgressBar_SetSmoothReverse(p.instance, value)
}

// Step
func (p *TProgressBar) Step() int32 {
    return ProgressBar_GetStep(p.instance)
}

// SetStep
func (p *TProgressBar) SetStep(value int32) {
    ProgressBar_SetStep(p.instance, value)
}

// State
func (p *TProgressBar) State() TProgressBarState {
    return ProgressBar_GetState(p.instance)
}

// SetState
func (p *TProgressBar) SetState(value TProgressBarState) {
    ProgressBar_SetState(p.instance, value)
}

// ShowHint
func (p *TProgressBar) ShowHint() bool {
    return ProgressBar_GetShowHint(p.instance)
}

// SetShowHint
func (p *TProgressBar) SetShowHint(value bool) {
    ProgressBar_SetShowHint(p.instance, value)
}

// TabOrder
func (p *TProgressBar) TabOrder() TTabOrder {
    return ProgressBar_GetTabOrder(p.instance)
}

// SetTabOrder
func (p *TProgressBar) SetTabOrder(value TTabOrder) {
    ProgressBar_SetTabOrder(p.instance, value)
}

// TabStop
func (p *TProgressBar) TabStop() bool {
    return ProgressBar_GetTabStop(p.instance)
}

// SetTabStop
func (p *TProgressBar) SetTabStop(value bool) {
    ProgressBar_SetTabStop(p.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (p *TProgressBar) Visible() bool {
    return ProgressBar_GetVisible(p.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (p *TProgressBar) SetVisible(value bool) {
    ProgressBar_SetVisible(p.instance, value)
}

// StyleElements
func (p *TProgressBar) StyleElements() TStyleElements {
    return ProgressBar_GetStyleElements(p.instance)
}

// SetStyleElements
func (p *TProgressBar) SetStyleElements(value TStyleElements) {
    ProgressBar_SetStyleElements(p.instance, value)
}

// SetOnContextPopup
func (p *TProgressBar) SetOnContextPopup(fn TContextPopupEvent) {
    ProgressBar_SetOnContextPopup(p.instance, fn)
}

// SetOnDragDrop
func (p *TProgressBar) SetOnDragDrop(fn TDragDropEvent) {
    ProgressBar_SetOnDragDrop(p.instance, fn)
}

// SetOnDragOver
func (p *TProgressBar) SetOnDragOver(fn TDragOverEvent) {
    ProgressBar_SetOnDragOver(p.instance, fn)
}

// SetOnEndDock
func (p *TProgressBar) SetOnEndDock(fn TEndDragEvent) {
    ProgressBar_SetOnEndDock(p.instance, fn)
}

// SetOnEndDrag
func (p *TProgressBar) SetOnEndDrag(fn TEndDragEvent) {
    ProgressBar_SetOnEndDrag(p.instance, fn)
}

// SetOnEnter
func (p *TProgressBar) SetOnEnter(fn TNotifyEvent) {
    ProgressBar_SetOnEnter(p.instance, fn)
}

// SetOnExit
func (p *TProgressBar) SetOnExit(fn TNotifyEvent) {
    ProgressBar_SetOnExit(p.instance, fn)
}

// SetOnMouseDown
func (p *TProgressBar) SetOnMouseDown(fn TMouseEvent) {
    ProgressBar_SetOnMouseDown(p.instance, fn)
}

// SetOnMouseEnter
func (p *TProgressBar) SetOnMouseEnter(fn TNotifyEvent) {
    ProgressBar_SetOnMouseEnter(p.instance, fn)
}

// SetOnMouseLeave
func (p *TProgressBar) SetOnMouseLeave(fn TNotifyEvent) {
    ProgressBar_SetOnMouseLeave(p.instance, fn)
}

// SetOnMouseMove
func (p *TProgressBar) SetOnMouseMove(fn TMouseMoveEvent) {
    ProgressBar_SetOnMouseMove(p.instance, fn)
}

// SetOnMouseUp
func (p *TProgressBar) SetOnMouseUp(fn TMouseEvent) {
    ProgressBar_SetOnMouseUp(p.instance, fn)
}

// SetOnStartDock
func (p *TProgressBar) SetOnStartDock(fn TStartDockEvent) {
    ProgressBar_SetOnStartDock(p.instance, fn)
}

// DockClientCount
func (p *TProgressBar) DockClientCount() int32 {
    return ProgressBar_GetDockClientCount(p.instance)
}

// DockSite
func (p *TProgressBar) DockSite() bool {
    return ProgressBar_GetDockSite(p.instance)
}

// SetDockSite
func (p *TProgressBar) SetDockSite(value bool) {
    ProgressBar_SetDockSite(p.instance, value)
}

// AlignDisabled
func (p *TProgressBar) AlignDisabled() bool {
    return ProgressBar_GetAlignDisabled(p.instance)
}

// MouseInClient
func (p *TProgressBar) MouseInClient() bool {
    return ProgressBar_GetMouseInClient(p.instance)
}

// VisibleDockClientCount
func (p *TProgressBar) VisibleDockClientCount() int32 {
    return ProgressBar_GetVisibleDockClientCount(p.instance)
}

// Brush
func (p *TProgressBar) Brush() *TBrush {
    return BrushFromInst(ProgressBar_GetBrush(p.instance))
}

// ControlCount
func (p *TProgressBar) ControlCount() int32 {
    return ProgressBar_GetControlCount(p.instance)
}

// Handle
func (p *TProgressBar) Handle() HWND {
    return ProgressBar_GetHandle(p.instance)
}

// ParentWindow
func (p *TProgressBar) ParentWindow() HWND {
    return ProgressBar_GetParentWindow(p.instance)
}

// SetParentWindow
func (p *TProgressBar) SetParentWindow(value HWND) {
    ProgressBar_SetParentWindow(p.instance, value)
}

// UseDockManager
func (p *TProgressBar) UseDockManager() bool {
    return ProgressBar_GetUseDockManager(p.instance)
}

// SetUseDockManager
func (p *TProgressBar) SetUseDockManager(value bool) {
    ProgressBar_SetUseDockManager(p.instance, value)
}

// Action
func (p *TProgressBar) Action() *TAction {
    return ActionFromInst(ProgressBar_GetAction(p.instance))
}

// SetAction
func (p *TProgressBar) SetAction(value IComponent) {
    ProgressBar_SetAction(p.instance, CheckPtr(value))
}

// BiDiMode
func (p *TProgressBar) BiDiMode() TBiDiMode {
    return ProgressBar_GetBiDiMode(p.instance)
}

// SetBiDiMode
func (p *TProgressBar) SetBiDiMode(value TBiDiMode) {
    ProgressBar_SetBiDiMode(p.instance, value)
}

// BoundsRect
func (p *TProgressBar) BoundsRect() TRect {
    return ProgressBar_GetBoundsRect(p.instance)
}

// SetBoundsRect
func (p *TProgressBar) SetBoundsRect(value TRect) {
    ProgressBar_SetBoundsRect(p.instance, value)
}

// ClientHeight
func (p *TProgressBar) ClientHeight() int32 {
    return ProgressBar_GetClientHeight(p.instance)
}

// SetClientHeight
func (p *TProgressBar) SetClientHeight(value int32) {
    ProgressBar_SetClientHeight(p.instance, value)
}

// ClientOrigin
func (p *TProgressBar) ClientOrigin() TPoint {
    return ProgressBar_GetClientOrigin(p.instance)
}

// ClientRect
func (p *TProgressBar) ClientRect() TRect {
    return ProgressBar_GetClientRect(p.instance)
}

// ClientWidth
func (p *TProgressBar) ClientWidth() int32 {
    return ProgressBar_GetClientWidth(p.instance)
}

// SetClientWidth
func (p *TProgressBar) SetClientWidth(value int32) {
    ProgressBar_SetClientWidth(p.instance, value)
}

// ControlState
func (p *TProgressBar) ControlState() TControlState {
    return ProgressBar_GetControlState(p.instance)
}

// SetControlState
func (p *TProgressBar) SetControlState(value TControlState) {
    ProgressBar_SetControlState(p.instance, value)
}

// ControlStyle
func (p *TProgressBar) ControlStyle() TControlStyle {
    return ProgressBar_GetControlStyle(p.instance)
}

// SetControlStyle
func (p *TProgressBar) SetControlStyle(value TControlStyle) {
    ProgressBar_SetControlStyle(p.instance, value)
}

// ExplicitLeft
func (p *TProgressBar) ExplicitLeft() int32 {
    return ProgressBar_GetExplicitLeft(p.instance)
}

// ExplicitTop
func (p *TProgressBar) ExplicitTop() int32 {
    return ProgressBar_GetExplicitTop(p.instance)
}

// ExplicitWidth
func (p *TProgressBar) ExplicitWidth() int32 {
    return ProgressBar_GetExplicitWidth(p.instance)
}

// ExplicitHeight
func (p *TProgressBar) ExplicitHeight() int32 {
    return ProgressBar_GetExplicitHeight(p.instance)
}

// Floating
func (p *TProgressBar) Floating() bool {
    return ProgressBar_GetFloating(p.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (p *TProgressBar) Parent() *TWinControl {
    return WinControlFromInst(ProgressBar_GetParent(p.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (p *TProgressBar) SetParent(value IWinControl) {
    ProgressBar_SetParent(p.instance, CheckPtr(value))
}

// AlignWithMargins
func (p *TProgressBar) AlignWithMargins() bool {
    return ProgressBar_GetAlignWithMargins(p.instance)
}

// SetAlignWithMargins
func (p *TProgressBar) SetAlignWithMargins(value bool) {
    ProgressBar_SetAlignWithMargins(p.instance, value)
}

// Left
func (p *TProgressBar) Left() int32 {
    return ProgressBar_GetLeft(p.instance)
}

// SetLeft
func (p *TProgressBar) SetLeft(value int32) {
    ProgressBar_SetLeft(p.instance, value)
}

// Top
func (p *TProgressBar) Top() int32 {
    return ProgressBar_GetTop(p.instance)
}

// SetTop
func (p *TProgressBar) SetTop(value int32) {
    ProgressBar_SetTop(p.instance, value)
}

// Width
func (p *TProgressBar) Width() int32 {
    return ProgressBar_GetWidth(p.instance)
}

// SetWidth
func (p *TProgressBar) SetWidth(value int32) {
    ProgressBar_SetWidth(p.instance, value)
}

// Height
func (p *TProgressBar) Height() int32 {
    return ProgressBar_GetHeight(p.instance)
}

// SetHeight
func (p *TProgressBar) SetHeight(value int32) {
    ProgressBar_SetHeight(p.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (p *TProgressBar) Cursor() TCursor {
    return ProgressBar_GetCursor(p.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (p *TProgressBar) SetCursor(value TCursor) {
    ProgressBar_SetCursor(p.instance, value)
}

// Margins
func (p *TProgressBar) Margins() *TMargins {
    return MarginsFromInst(ProgressBar_GetMargins(p.instance))
}

// SetMargins
func (p *TProgressBar) SetMargins(value *TMargins) {
    ProgressBar_SetMargins(p.instance, CheckPtr(value))
}

// CustomHint
func (p *TProgressBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(ProgressBar_GetCustomHint(p.instance))
}

// SetCustomHint
func (p *TProgressBar) SetCustomHint(value IComponent) {
    ProgressBar_SetCustomHint(p.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TProgressBar) ComponentCount() int32 {
    return ProgressBar_GetComponentCount(p.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (p *TProgressBar) ComponentIndex() int32 {
    return ProgressBar_GetComponentIndex(p.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (p *TProgressBar) SetComponentIndex(value int32) {
    ProgressBar_SetComponentIndex(p.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TProgressBar) Owner() *TComponent {
    return ComponentFromInst(ProgressBar_GetOwner(p.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (p *TProgressBar) Name() string {
    return ProgressBar_GetName(p.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (p *TProgressBar) SetName(value string) {
    ProgressBar_SetName(p.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TProgressBar) Tag() int {
    return ProgressBar_GetTag(p.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TProgressBar) SetTag(value int) {
    ProgressBar_SetTag(p.instance, value)
}

// DockClients
func (p *TProgressBar) DockClients(Index int32) *TControl {
    return ControlFromInst(ProgressBar_GetDockClients(p.instance, Index))
}

// Controls
func (p *TProgressBar) Controls(Index int32) *TControl {
    return ControlFromInst(ProgressBar_GetControls(p.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TProgressBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ProgressBar_GetComponents(p.instance, AIndex))
}

