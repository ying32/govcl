
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

type TTrackBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTrackBar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTrackBar(owner IComponent) *TTrackBar {
    t := new(TTrackBar)
    t.instance = TrackBar_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TrackBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TrackBarFromInst(inst uintptr) *TTrackBar {
    t := new(TTrackBar)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TrackBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TrackBarFromObj(obj IObject) *TTrackBar {
    t := new(TTrackBar)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TrackBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TrackBarFromUnsafePointer(ptr unsafe.Pointer) *TTrackBar {
    t := new(TTrackBar)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTrackBar) Free() {
    if t.instance != 0 {
        TrackBar_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTrackBar) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTrackBar) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTrackBar) IsValid() bool {
    return t.instance != 0
}

// TTrackBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTrackBarClass() TClass {
    return TrackBar_StaticClassType()
}

// SetTick
func (t *TTrackBar) SetTick(Value int32) {
    TrackBar_SetTick(t.instance, Value)
}

// CanFocus
func (t *TTrackBar) CanFocus() bool {
    return TrackBar_CanFocus(t.instance)
}

// ContainsControl
func (t *TTrackBar) ContainsControl(Control IControl) bool {
    return TrackBar_ContainsControl(t.instance, CheckPtr(Control))
}

// ControlAtPos
func (t *TTrackBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(TrackBar_ControlAtPos(t.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (t *TTrackBar) DisableAlign() {
    TrackBar_DisableAlign(t.instance)
}

// EnableAlign
func (t *TTrackBar) EnableAlign() {
    TrackBar_EnableAlign(t.instance)
}

// FindChildControl
func (t *TTrackBar) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(TrackBar_FindChildControl(t.instance, ControlName))
}

// FlipChildren
func (t *TTrackBar) FlipChildren(AllLevels bool) {
    TrackBar_FlipChildren(t.instance, AllLevels)
}

// Focused
func (t *TTrackBar) Focused() bool {
    return TrackBar_Focused(t.instance)
}

// HandleAllocated
func (t *TTrackBar) HandleAllocated() bool {
    return TrackBar_HandleAllocated(t.instance)
}

// InsertControl
func (t *TTrackBar) InsertControl(AControl IControl) {
    TrackBar_InsertControl(t.instance, CheckPtr(AControl))
}

// Invalidate
func (t *TTrackBar) Invalidate() {
    TrackBar_Invalidate(t.instance)
}

// PaintTo
func (t *TTrackBar) PaintTo(DC HDC, X int32, Y int32) {
    TrackBar_PaintTo(t.instance, DC , X , Y)
}

// RemoveControl
func (t *TTrackBar) RemoveControl(AControl IControl) {
    TrackBar_RemoveControl(t.instance, CheckPtr(AControl))
}

// Realign
func (t *TTrackBar) Realign() {
    TrackBar_Realign(t.instance)
}

// Repaint
func (t *TTrackBar) Repaint() {
    TrackBar_Repaint(t.instance)
}

// ScaleBy
func (t *TTrackBar) ScaleBy(M int32, D int32) {
    TrackBar_ScaleBy(t.instance, M , D)
}

// ScrollBy
func (t *TTrackBar) ScrollBy(DeltaX int32, DeltaY int32) {
    TrackBar_ScrollBy(t.instance, DeltaX , DeltaY)
}

// SetBounds
func (t *TTrackBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    TrackBar_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (t *TTrackBar) SetFocus() {
    TrackBar_SetFocus(t.instance)
}

// Update
func (t *TTrackBar) Update() {
    TrackBar_Update(t.instance)
}

// UpdateControlState
func (t *TTrackBar) UpdateControlState() {
    TrackBar_UpdateControlState(t.instance)
}

// BringToFront
func (t *TTrackBar) BringToFront() {
    TrackBar_BringToFront(t.instance)
}

// ClientToScreen
func (t *TTrackBar) ClientToScreen(Point TPoint) TPoint {
    return TrackBar_ClientToScreen(t.instance, Point)
}

// ClientToParent
func (t *TTrackBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return TrackBar_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// Dragging
func (t *TTrackBar) Dragging() bool {
    return TrackBar_Dragging(t.instance)
}

// HasParent
func (t *TTrackBar) HasParent() bool {
    return TrackBar_HasParent(t.instance)
}

// Hide
func (t *TTrackBar) Hide() {
    TrackBar_Hide(t.instance)
}

// Perform
func (t *TTrackBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return TrackBar_Perform(t.instance, Msg , WParam , LParam)
}

// Refresh
func (t *TTrackBar) Refresh() {
    TrackBar_Refresh(t.instance)
}

// ScreenToClient
func (t *TTrackBar) ScreenToClient(Point TPoint) TPoint {
    return TrackBar_ScreenToClient(t.instance, Point)
}

// ParentToClient
func (t *TTrackBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return TrackBar_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (t *TTrackBar) SendToBack() {
    TrackBar_SendToBack(t.instance)
}

// Show
func (t *TTrackBar) Show() {
    TrackBar_Show(t.instance)
}

// GetTextBuf
func (t *TTrackBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return TrackBar_GetTextBuf(t.instance, Buffer , BufSize)
}

// GetTextLen
func (t *TTrackBar) GetTextLen() int32 {
    return TrackBar_GetTextLen(t.instance)
}

// SetTextBuf
func (t *TTrackBar) SetTextBuf(Buffer string) {
    TrackBar_SetTextBuf(t.instance, Buffer)
}

// FindComponent
func (t *TTrackBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TrackBar_FindComponent(t.instance, AName))
}

// GetNamePath
func (t *TTrackBar) GetNamePath() string {
    return TrackBar_GetNamePath(t.instance)
}

// Assign
func (t *TTrackBar) Assign(Source IObject) {
    TrackBar_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTrackBar) DisposeOf() {
    TrackBar_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTrackBar) ClassType() TClass {
    return TrackBar_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTrackBar) ClassName() string {
    return TrackBar_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTrackBar) InstanceSize() int32 {
    return TrackBar_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTrackBar) InheritsFrom(AClass TClass) bool {
    return TrackBar_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTrackBar) Equals(Obj IObject) bool {
    return TrackBar_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTrackBar) GetHashCode() int32 {
    return TrackBar_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTrackBar) ToString() string {
    return TrackBar_ToString(t.instance)
}

// Align
func (t *TTrackBar) Align() TAlign {
    return TrackBar_GetAlign(t.instance)
}

// SetAlign
func (t *TTrackBar) SetAlign(value TAlign) {
    TrackBar_SetAlign(t.instance, value)
}

// Anchors
func (t *TTrackBar) Anchors() TAnchors {
    return TrackBar_GetAnchors(t.instance)
}

// SetAnchors
func (t *TTrackBar) SetAnchors(value TAnchors) {
    TrackBar_SetAnchors(t.instance, value)
}

// BorderWidth
func (t *TTrackBar) BorderWidth() int32 {
    return TrackBar_GetBorderWidth(t.instance)
}

// SetBorderWidth
func (t *TTrackBar) SetBorderWidth(value int32) {
    TrackBar_SetBorderWidth(t.instance, value)
}

// DoubleBuffered
func (t *TTrackBar) DoubleBuffered() bool {
    return TrackBar_GetDoubleBuffered(t.instance)
}

// SetDoubleBuffered
func (t *TTrackBar) SetDoubleBuffered(value bool) {
    TrackBar_SetDoubleBuffered(t.instance, value)
}

// DragCursor
func (t *TTrackBar) DragCursor() TCursor {
    return TrackBar_GetDragCursor(t.instance)
}

// SetDragCursor
func (t *TTrackBar) SetDragCursor(value TCursor) {
    TrackBar_SetDragCursor(t.instance, value)
}

// DragKind
func (t *TTrackBar) DragKind() TDragKind {
    return TrackBar_GetDragKind(t.instance)
}

// SetDragKind
func (t *TTrackBar) SetDragKind(value TDragKind) {
    TrackBar_SetDragKind(t.instance, value)
}

// DragMode
func (t *TTrackBar) DragMode() TDragMode {
    return TrackBar_GetDragMode(t.instance)
}

// SetDragMode
func (t *TTrackBar) SetDragMode(value TDragMode) {
    TrackBar_SetDragMode(t.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TTrackBar) Enabled() bool {
    return TrackBar_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TTrackBar) SetEnabled(value bool) {
    TrackBar_SetEnabled(t.instance, value)
}

// LineSize
func (t *TTrackBar) LineSize() int32 {
    return TrackBar_GetLineSize(t.instance)
}

// SetLineSize
func (t *TTrackBar) SetLineSize(value int32) {
    TrackBar_SetLineSize(t.instance, value)
}

// Max
func (t *TTrackBar) Max() int32 {
    return TrackBar_GetMax(t.instance)
}

// SetMax
func (t *TTrackBar) SetMax(value int32) {
    TrackBar_SetMax(t.instance, value)
}

// Min
func (t *TTrackBar) Min() int32 {
    return TrackBar_GetMin(t.instance)
}

// SetMin
func (t *TTrackBar) SetMin(value int32) {
    TrackBar_SetMin(t.instance, value)
}

// Orientation
func (t *TTrackBar) Orientation() TTrackBarOrientation {
    return TrackBar_GetOrientation(t.instance)
}

// SetOrientation
func (t *TTrackBar) SetOrientation(value TTrackBarOrientation) {
    TrackBar_SetOrientation(t.instance, value)
}

// ParentCtl3D
func (t *TTrackBar) ParentCtl3D() bool {
    return TrackBar_GetParentCtl3D(t.instance)
}

// SetParentCtl3D
func (t *TTrackBar) SetParentCtl3D(value bool) {
    TrackBar_SetParentCtl3D(t.instance, value)
}

// ParentDoubleBuffered
func (t *TTrackBar) ParentDoubleBuffered() bool {
    return TrackBar_GetParentDoubleBuffered(t.instance)
}

// SetParentDoubleBuffered
func (t *TTrackBar) SetParentDoubleBuffered(value bool) {
    TrackBar_SetParentDoubleBuffered(t.instance, value)
}

// ParentShowHint
func (t *TTrackBar) ParentShowHint() bool {
    return TrackBar_GetParentShowHint(t.instance)
}

// SetParentShowHint
func (t *TTrackBar) SetParentShowHint(value bool) {
    TrackBar_SetParentShowHint(t.instance, value)
}

// PageSize
func (t *TTrackBar) PageSize() int32 {
    return TrackBar_GetPageSize(t.instance)
}

// SetPageSize
func (t *TTrackBar) SetPageSize(value int32) {
    TrackBar_SetPageSize(t.instance, value)
}

// PopupMenu
func (t *TTrackBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TrackBar_GetPopupMenu(t.instance))
}

// SetPopupMenu
func (t *TTrackBar) SetPopupMenu(value IComponent) {
    TrackBar_SetPopupMenu(t.instance, CheckPtr(value))
}

// Frequency
func (t *TTrackBar) Frequency() int32 {
    return TrackBar_GetFrequency(t.instance)
}

// SetFrequency
func (t *TTrackBar) SetFrequency(value int32) {
    TrackBar_SetFrequency(t.instance, value)
}

// Position
func (t *TTrackBar) Position() int32 {
    return TrackBar_GetPosition(t.instance)
}

// SetPosition
func (t *TTrackBar) SetPosition(value int32) {
    TrackBar_SetPosition(t.instance, value)
}

// PositionToolTip
func (t *TTrackBar) PositionToolTip() TPositionToolTip {
    return TrackBar_GetPositionToolTip(t.instance)
}

// SetPositionToolTip
func (t *TTrackBar) SetPositionToolTip(value TPositionToolTip) {
    TrackBar_SetPositionToolTip(t.instance, value)
}

// SliderVisible
func (t *TTrackBar) SliderVisible() bool {
    return TrackBar_GetSliderVisible(t.instance)
}

// SetSliderVisible
func (t *TTrackBar) SetSliderVisible(value bool) {
    TrackBar_SetSliderVisible(t.instance, value)
}

// SelEnd
func (t *TTrackBar) SelEnd() int32 {
    return TrackBar_GetSelEnd(t.instance)
}

// SetSelEnd
func (t *TTrackBar) SetSelEnd(value int32) {
    TrackBar_SetSelEnd(t.instance, value)
}

// SelStart
func (t *TTrackBar) SelStart() int32 {
    return TrackBar_GetSelStart(t.instance)
}

// SetSelStart
func (t *TTrackBar) SetSelStart(value int32) {
    TrackBar_SetSelStart(t.instance, value)
}

// ShowHint
func (t *TTrackBar) ShowHint() bool {
    return TrackBar_GetShowHint(t.instance)
}

// SetShowHint
func (t *TTrackBar) SetShowHint(value bool) {
    TrackBar_SetShowHint(t.instance, value)
}

// ShowSelRange
func (t *TTrackBar) ShowSelRange() bool {
    return TrackBar_GetShowSelRange(t.instance)
}

// SetShowSelRange
func (t *TTrackBar) SetShowSelRange(value bool) {
    TrackBar_SetShowSelRange(t.instance, value)
}

// TabOrder
func (t *TTrackBar) TabOrder() TTabOrder {
    return TrackBar_GetTabOrder(t.instance)
}

// SetTabOrder
func (t *TTrackBar) SetTabOrder(value TTabOrder) {
    TrackBar_SetTabOrder(t.instance, value)
}

// TabStop
func (t *TTrackBar) TabStop() bool {
    return TrackBar_GetTabStop(t.instance)
}

// SetTabStop
func (t *TTrackBar) SetTabStop(value bool) {
    TrackBar_SetTabStop(t.instance, value)
}

// ThumbLength
func (t *TTrackBar) ThumbLength() int32 {
    return TrackBar_GetThumbLength(t.instance)
}

// SetThumbLength
func (t *TTrackBar) SetThumbLength(value int32) {
    TrackBar_SetThumbLength(t.instance, value)
}

// TickMarks
func (t *TTrackBar) TickMarks() TTickMark {
    return TrackBar_GetTickMarks(t.instance)
}

// SetTickMarks
func (t *TTrackBar) SetTickMarks(value TTickMark) {
    TrackBar_SetTickMarks(t.instance, value)
}

// TickStyle
func (t *TTrackBar) TickStyle() TTickStyle {
    return TrackBar_GetTickStyle(t.instance)
}

// SetTickStyle
func (t *TTrackBar) SetTickStyle(value TTickStyle) {
    TrackBar_SetTickStyle(t.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TTrackBar) Visible() bool {
    return TrackBar_GetVisible(t.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TTrackBar) SetVisible(value bool) {
    TrackBar_SetVisible(t.instance, value)
}

// StyleElements
func (t *TTrackBar) StyleElements() TStyleElements {
    return TrackBar_GetStyleElements(t.instance)
}

// SetStyleElements
func (t *TTrackBar) SetStyleElements(value TStyleElements) {
    TrackBar_SetStyleElements(t.instance, value)
}

// SetOnContextPopup
func (t *TTrackBar) SetOnContextPopup(fn TContextPopupEvent) {
    TrackBar_SetOnContextPopup(t.instance, fn)
}

// SetOnChange
func (t *TTrackBar) SetOnChange(fn TNotifyEvent) {
    TrackBar_SetOnChange(t.instance, fn)
}

// SetOnDragDrop
func (t *TTrackBar) SetOnDragDrop(fn TDragDropEvent) {
    TrackBar_SetOnDragDrop(t.instance, fn)
}

// SetOnDragOver
func (t *TTrackBar) SetOnDragOver(fn TDragOverEvent) {
    TrackBar_SetOnDragOver(t.instance, fn)
}

// SetOnEndDock
func (t *TTrackBar) SetOnEndDock(fn TEndDragEvent) {
    TrackBar_SetOnEndDock(t.instance, fn)
}

// SetOnEndDrag
func (t *TTrackBar) SetOnEndDrag(fn TEndDragEvent) {
    TrackBar_SetOnEndDrag(t.instance, fn)
}

// SetOnEnter
func (t *TTrackBar) SetOnEnter(fn TNotifyEvent) {
    TrackBar_SetOnEnter(t.instance, fn)
}

// SetOnExit
func (t *TTrackBar) SetOnExit(fn TNotifyEvent) {
    TrackBar_SetOnExit(t.instance, fn)
}

// SetOnKeyDown
func (t *TTrackBar) SetOnKeyDown(fn TKeyEvent) {
    TrackBar_SetOnKeyDown(t.instance, fn)
}

// SetOnKeyPress
func (t *TTrackBar) SetOnKeyPress(fn TKeyPressEvent) {
    TrackBar_SetOnKeyPress(t.instance, fn)
}

// SetOnKeyUp
func (t *TTrackBar) SetOnKeyUp(fn TKeyEvent) {
    TrackBar_SetOnKeyUp(t.instance, fn)
}

// SetOnStartDock
func (t *TTrackBar) SetOnStartDock(fn TStartDockEvent) {
    TrackBar_SetOnStartDock(t.instance, fn)
}

// DockClientCount
func (t *TTrackBar) DockClientCount() int32 {
    return TrackBar_GetDockClientCount(t.instance)
}

// DockSite
func (t *TTrackBar) DockSite() bool {
    return TrackBar_GetDockSite(t.instance)
}

// SetDockSite
func (t *TTrackBar) SetDockSite(value bool) {
    TrackBar_SetDockSite(t.instance, value)
}

// AlignDisabled
func (t *TTrackBar) AlignDisabled() bool {
    return TrackBar_GetAlignDisabled(t.instance)
}

// MouseInClient
func (t *TTrackBar) MouseInClient() bool {
    return TrackBar_GetMouseInClient(t.instance)
}

// VisibleDockClientCount
func (t *TTrackBar) VisibleDockClientCount() int32 {
    return TrackBar_GetVisibleDockClientCount(t.instance)
}

// Brush
func (t *TTrackBar) Brush() *TBrush {
    return BrushFromInst(TrackBar_GetBrush(t.instance))
}

// ControlCount
func (t *TTrackBar) ControlCount() int32 {
    return TrackBar_GetControlCount(t.instance)
}

// Handle
func (t *TTrackBar) Handle() HWND {
    return TrackBar_GetHandle(t.instance)
}

// ParentWindow
func (t *TTrackBar) ParentWindow() HWND {
    return TrackBar_GetParentWindow(t.instance)
}

// SetParentWindow
func (t *TTrackBar) SetParentWindow(value HWND) {
    TrackBar_SetParentWindow(t.instance, value)
}

// UseDockManager
func (t *TTrackBar) UseDockManager() bool {
    return TrackBar_GetUseDockManager(t.instance)
}

// SetUseDockManager
func (t *TTrackBar) SetUseDockManager(value bool) {
    TrackBar_SetUseDockManager(t.instance, value)
}

// Action
func (t *TTrackBar) Action() *TAction {
    return ActionFromInst(TrackBar_GetAction(t.instance))
}

// SetAction
func (t *TTrackBar) SetAction(value IComponent) {
    TrackBar_SetAction(t.instance, CheckPtr(value))
}

// BiDiMode
func (t *TTrackBar) BiDiMode() TBiDiMode {
    return TrackBar_GetBiDiMode(t.instance)
}

// SetBiDiMode
func (t *TTrackBar) SetBiDiMode(value TBiDiMode) {
    TrackBar_SetBiDiMode(t.instance, value)
}

// BoundsRect
func (t *TTrackBar) BoundsRect() TRect {
    return TrackBar_GetBoundsRect(t.instance)
}

// SetBoundsRect
func (t *TTrackBar) SetBoundsRect(value TRect) {
    TrackBar_SetBoundsRect(t.instance, value)
}

// ClientHeight
func (t *TTrackBar) ClientHeight() int32 {
    return TrackBar_GetClientHeight(t.instance)
}

// SetClientHeight
func (t *TTrackBar) SetClientHeight(value int32) {
    TrackBar_SetClientHeight(t.instance, value)
}

// ClientOrigin
func (t *TTrackBar) ClientOrigin() TPoint {
    return TrackBar_GetClientOrigin(t.instance)
}

// ClientRect
func (t *TTrackBar) ClientRect() TRect {
    return TrackBar_GetClientRect(t.instance)
}

// ClientWidth
func (t *TTrackBar) ClientWidth() int32 {
    return TrackBar_GetClientWidth(t.instance)
}

// SetClientWidth
func (t *TTrackBar) SetClientWidth(value int32) {
    TrackBar_SetClientWidth(t.instance, value)
}

// ControlState
func (t *TTrackBar) ControlState() TControlState {
    return TrackBar_GetControlState(t.instance)
}

// SetControlState
func (t *TTrackBar) SetControlState(value TControlState) {
    TrackBar_SetControlState(t.instance, value)
}

// ControlStyle
func (t *TTrackBar) ControlStyle() TControlStyle {
    return TrackBar_GetControlStyle(t.instance)
}

// SetControlStyle
func (t *TTrackBar) SetControlStyle(value TControlStyle) {
    TrackBar_SetControlStyle(t.instance, value)
}

// ExplicitLeft
func (t *TTrackBar) ExplicitLeft() int32 {
    return TrackBar_GetExplicitLeft(t.instance)
}

// ExplicitTop
func (t *TTrackBar) ExplicitTop() int32 {
    return TrackBar_GetExplicitTop(t.instance)
}

// ExplicitWidth
func (t *TTrackBar) ExplicitWidth() int32 {
    return TrackBar_GetExplicitWidth(t.instance)
}

// ExplicitHeight
func (t *TTrackBar) ExplicitHeight() int32 {
    return TrackBar_GetExplicitHeight(t.instance)
}

// Floating
func (t *TTrackBar) Floating() bool {
    return TrackBar_GetFloating(t.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TTrackBar) Parent() *TWinControl {
    return WinControlFromInst(TrackBar_GetParent(t.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (t *TTrackBar) SetParent(value IWinControl) {
    TrackBar_SetParent(t.instance, CheckPtr(value))
}

// AlignWithMargins
func (t *TTrackBar) AlignWithMargins() bool {
    return TrackBar_GetAlignWithMargins(t.instance)
}

// SetAlignWithMargins
func (t *TTrackBar) SetAlignWithMargins(value bool) {
    TrackBar_SetAlignWithMargins(t.instance, value)
}

// Left
func (t *TTrackBar) Left() int32 {
    return TrackBar_GetLeft(t.instance)
}

// SetLeft
func (t *TTrackBar) SetLeft(value int32) {
    TrackBar_SetLeft(t.instance, value)
}

// Top
func (t *TTrackBar) Top() int32 {
    return TrackBar_GetTop(t.instance)
}

// SetTop
func (t *TTrackBar) SetTop(value int32) {
    TrackBar_SetTop(t.instance, value)
}

// Width
func (t *TTrackBar) Width() int32 {
    return TrackBar_GetWidth(t.instance)
}

// SetWidth
func (t *TTrackBar) SetWidth(value int32) {
    TrackBar_SetWidth(t.instance, value)
}

// Height
func (t *TTrackBar) Height() int32 {
    return TrackBar_GetHeight(t.instance)
}

// SetHeight
func (t *TTrackBar) SetHeight(value int32) {
    TrackBar_SetHeight(t.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (t *TTrackBar) Cursor() TCursor {
    return TrackBar_GetCursor(t.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (t *TTrackBar) SetCursor(value TCursor) {
    TrackBar_SetCursor(t.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (t *TTrackBar) Hint() string {
    return TrackBar_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (t *TTrackBar) SetHint(value string) {
    TrackBar_SetHint(t.instance, value)
}

// Margins
func (t *TTrackBar) Margins() *TMargins {
    return MarginsFromInst(TrackBar_GetMargins(t.instance))
}

// SetMargins
func (t *TTrackBar) SetMargins(value *TMargins) {
    TrackBar_SetMargins(t.instance, CheckPtr(value))
}

// CustomHint
func (t *TTrackBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(TrackBar_GetCustomHint(t.instance))
}

// SetCustomHint
func (t *TTrackBar) SetCustomHint(value IComponent) {
    TrackBar_SetCustomHint(t.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTrackBar) ComponentCount() int32 {
    return TrackBar_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TTrackBar) ComponentIndex() int32 {
    return TrackBar_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TTrackBar) SetComponentIndex(value int32) {
    TrackBar_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTrackBar) Owner() *TComponent {
    return ComponentFromInst(TrackBar_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTrackBar) Name() string {
    return TrackBar_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTrackBar) SetName(value string) {
    TrackBar_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTrackBar) Tag() int {
    return TrackBar_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTrackBar) SetTag(value int) {
    TrackBar_SetTag(t.instance, value)
}

// DockClients
func (t *TTrackBar) DockClients(Index int32) *TControl {
    return ControlFromInst(TrackBar_GetDockClients(t.instance, Index))
}

// Controls
func (t *TTrackBar) Controls(Index int32) *TControl {
    return ControlFromInst(TrackBar_GetControls(t.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTrackBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TrackBar_GetComponents(t.instance, AIndex))
}

