
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

type THeaderControl struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewHeaderControl
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewHeaderControl(owner IComponent) *THeaderControl {
    h := new(THeaderControl)
    h.instance = HeaderControl_Create(CheckPtr(owner))
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HeaderControlFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func HeaderControlFromInst(inst uintptr) *THeaderControl {
    h := new(THeaderControl)
    h.instance = inst
    h.ptr = unsafe.Pointer(inst)
    return h
}

// HeaderControlFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func HeaderControlFromObj(obj IObject) *THeaderControl {
    h := new(THeaderControl)
    h.instance = CheckPtr(obj)
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HeaderControlFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func HeaderControlFromUnsafePointer(ptr unsafe.Pointer) *THeaderControl {
    h := new(THeaderControl)
    h.instance = uintptr(ptr)
    h.ptr = ptr
    return h
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (h *THeaderControl) Free() {
    if h.instance != 0 {
        HeaderControl_Free(h.instance)
        h.instance = 0
        h.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (h *THeaderControl) Instance() uintptr {
    return h.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (h *THeaderControl) UnsafeAddr() unsafe.Pointer {
    return h.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (h *THeaderControl) IsValid() bool {
    return h.instance != 0
}

// THeaderControlClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func THeaderControlClass() TClass {
    return HeaderControl_StaticClassType()
}

// FlipChildren
func (h *THeaderControl) FlipChildren(AllLevels bool) {
    HeaderControl_FlipChildren(h.instance, AllLevels)
}

// CanFocus
func (h *THeaderControl) CanFocus() bool {
    return HeaderControl_CanFocus(h.instance)
}

// Focused
func (h *THeaderControl) Focused() bool {
    return HeaderControl_Focused(h.instance)
}

// HandleAllocated
func (h *THeaderControl) HandleAllocated() bool {
    return HeaderControl_HandleAllocated(h.instance)
}

// Invalidate
func (h *THeaderControl) Invalidate() {
    HeaderControl_Invalidate(h.instance)
}

// Realign
func (h *THeaderControl) Realign() {
    HeaderControl_Realign(h.instance)
}

// Repaint
func (h *THeaderControl) Repaint() {
    HeaderControl_Repaint(h.instance)
}

// ScaleBy
func (h *THeaderControl) ScaleBy(M int32, D int32) {
    HeaderControl_ScaleBy(h.instance, M , D)
}

// SetBounds
func (h *THeaderControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    HeaderControl_SetBounds(h.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (h *THeaderControl) SetFocus() {
    HeaderControl_SetFocus(h.instance)
}

// Update
func (h *THeaderControl) Update() {
    HeaderControl_Update(h.instance)
}

// BringToFront
func (h *THeaderControl) BringToFront() {
    HeaderControl_BringToFront(h.instance)
}

// ClientToScreen
func (h *THeaderControl) ClientToScreen(Point TPoint) TPoint {
    return HeaderControl_ClientToScreen(h.instance, Point)
}

// ClientToParent
func (h *THeaderControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return HeaderControl_ClientToParent(h.instance, Point , CheckPtr(AParent))
}

// Dragging
func (h *THeaderControl) Dragging() bool {
    return HeaderControl_Dragging(h.instance)
}

// HasParent
func (h *THeaderControl) HasParent() bool {
    return HeaderControl_HasParent(h.instance)
}

// Hide
func (h *THeaderControl) Hide() {
    HeaderControl_Hide(h.instance)
}

// Perform
func (h *THeaderControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return HeaderControl_Perform(h.instance, Msg , WParam , LParam)
}

// Refresh
func (h *THeaderControl) Refresh() {
    HeaderControl_Refresh(h.instance)
}

// ScreenToClient
func (h *THeaderControl) ScreenToClient(Point TPoint) TPoint {
    return HeaderControl_ScreenToClient(h.instance, Point)
}

// ParentToClient
func (h *THeaderControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return HeaderControl_ParentToClient(h.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (h *THeaderControl) SendToBack() {
    HeaderControl_SendToBack(h.instance)
}

// Show
func (h *THeaderControl) Show() {
    HeaderControl_Show(h.instance)
}

// GetTextBuf
func (h *THeaderControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    return HeaderControl_GetTextBuf(h.instance, Buffer , BufSize)
}

// GetTextLen
func (h *THeaderControl) GetTextLen() int32 {
    return HeaderControl_GetTextLen(h.instance)
}

// FindComponent
func (h *THeaderControl) FindComponent(AName string) *TComponent {
    return ComponentFromInst(HeaderControl_FindComponent(h.instance, AName))
}

// GetNamePath
func (h *THeaderControl) GetNamePath() string {
    return HeaderControl_GetNamePath(h.instance)
}

// Assign
func (h *THeaderControl) Assign(Source IObject) {
    HeaderControl_Assign(h.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (h *THeaderControl) DisposeOf() {
    HeaderControl_DisposeOf(h.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (h *THeaderControl) ClassType() TClass {
    return HeaderControl_ClassType(h.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (h *THeaderControl) ClassName() string {
    return HeaderControl_ClassName(h.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (h *THeaderControl) InstanceSize() int32 {
    return HeaderControl_InstanceSize(h.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (h *THeaderControl) InheritsFrom(AClass TClass) bool {
    return HeaderControl_InheritsFrom(h.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (h *THeaderControl) Equals(Obj IObject) bool {
    return HeaderControl_Equals(h.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (h *THeaderControl) GetHashCode() int32 {
    return HeaderControl_GetHashCode(h.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (h *THeaderControl) ToString() string {
    return HeaderControl_ToString(h.instance)
}

// Align
func (h *THeaderControl) Align() TAlign {
    return HeaderControl_GetAlign(h.instance)
}

// SetAlign
func (h *THeaderControl) SetAlign(value TAlign) {
    HeaderControl_SetAlign(h.instance, value)
}

// Anchors
func (h *THeaderControl) Anchors() TAnchors {
    return HeaderControl_GetAnchors(h.instance)
}

// SetAnchors
func (h *THeaderControl) SetAnchors(value TAnchors) {
    HeaderControl_SetAnchors(h.instance, value)
}

// BiDiMode
func (h *THeaderControl) BiDiMode() TBiDiMode {
    return HeaderControl_GetBiDiMode(h.instance)
}

// SetBiDiMode
func (h *THeaderControl) SetBiDiMode(value TBiDiMode) {
    HeaderControl_SetBiDiMode(h.instance, value)
}

// BorderWidth
func (h *THeaderControl) BorderWidth() int32 {
    return HeaderControl_GetBorderWidth(h.instance)
}

// SetBorderWidth
func (h *THeaderControl) SetBorderWidth(value int32) {
    HeaderControl_SetBorderWidth(h.instance, value)
}

// DoubleBuffered
func (h *THeaderControl) DoubleBuffered() bool {
    return HeaderControl_GetDoubleBuffered(h.instance)
}

// SetDoubleBuffered
func (h *THeaderControl) SetDoubleBuffered(value bool) {
    HeaderControl_SetDoubleBuffered(h.instance, value)
}

// DragCursor
func (h *THeaderControl) DragCursor() TCursor {
    return HeaderControl_GetDragCursor(h.instance)
}

// SetDragCursor
func (h *THeaderControl) SetDragCursor(value TCursor) {
    HeaderControl_SetDragCursor(h.instance, value)
}

// DragKind
func (h *THeaderControl) DragKind() TDragKind {
    return HeaderControl_GetDragKind(h.instance)
}

// SetDragKind
func (h *THeaderControl) SetDragKind(value TDragKind) {
    HeaderControl_SetDragKind(h.instance, value)
}

// DragMode
func (h *THeaderControl) DragMode() TDragMode {
    return HeaderControl_GetDragMode(h.instance)
}

// SetDragMode
func (h *THeaderControl) SetDragMode(value TDragMode) {
    HeaderControl_SetDragMode(h.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (h *THeaderControl) Enabled() bool {
    return HeaderControl_GetEnabled(h.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (h *THeaderControl) SetEnabled(value bool) {
    HeaderControl_SetEnabled(h.instance, value)
}

// Font
func (h *THeaderControl) Font() *TFont {
    return FontFromInst(HeaderControl_GetFont(h.instance))
}

// SetFont
func (h *THeaderControl) SetFont(value *TFont) {
    HeaderControl_SetFont(h.instance, CheckPtr(value))
}

// FullDrag
func (h *THeaderControl) FullDrag() bool {
    return HeaderControl_GetFullDrag(h.instance)
}

// SetFullDrag
func (h *THeaderControl) SetFullDrag(value bool) {
    HeaderControl_SetFullDrag(h.instance, value)
}

// HotTrack
func (h *THeaderControl) HotTrack() bool {
    return HeaderControl_GetHotTrack(h.instance)
}

// SetHotTrack
func (h *THeaderControl) SetHotTrack(value bool) {
    HeaderControl_SetHotTrack(h.instance, value)
}

// Images
func (h *THeaderControl) Images() *TImageList {
    return ImageListFromInst(HeaderControl_GetImages(h.instance))
}

// SetImages
func (h *THeaderControl) SetImages(value IComponent) {
    HeaderControl_SetImages(h.instance, CheckPtr(value))
}

// ShowHint
func (h *THeaderControl) ShowHint() bool {
    return HeaderControl_GetShowHint(h.instance)
}

// SetShowHint
func (h *THeaderControl) SetShowHint(value bool) {
    HeaderControl_SetShowHint(h.instance, value)
}

// Style
func (h *THeaderControl) Style() THeaderStyle {
    return HeaderControl_GetStyle(h.instance)
}

// SetStyle
func (h *THeaderControl) SetStyle(value THeaderStyle) {
    HeaderControl_SetStyle(h.instance, value)
}

// ParentDoubleBuffered
func (h *THeaderControl) ParentDoubleBuffered() bool {
    return HeaderControl_GetParentDoubleBuffered(h.instance)
}

// SetParentDoubleBuffered
func (h *THeaderControl) SetParentDoubleBuffered(value bool) {
    HeaderControl_SetParentDoubleBuffered(h.instance, value)
}

// ParentFont
func (h *THeaderControl) ParentFont() bool {
    return HeaderControl_GetParentFont(h.instance)
}

// SetParentFont
func (h *THeaderControl) SetParentFont(value bool) {
    HeaderControl_SetParentFont(h.instance, value)
}

// ParentShowHint
func (h *THeaderControl) ParentShowHint() bool {
    return HeaderControl_GetParentShowHint(h.instance)
}

// SetParentShowHint
func (h *THeaderControl) SetParentShowHint(value bool) {
    HeaderControl_SetParentShowHint(h.instance, value)
}

// PopupMenu
func (h *THeaderControl) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(HeaderControl_GetPopupMenu(h.instance))
}

// SetPopupMenu
func (h *THeaderControl) SetPopupMenu(value IComponent) {
    HeaderControl_SetPopupMenu(h.instance, CheckPtr(value))
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (h *THeaderControl) Visible() bool {
    return HeaderControl_GetVisible(h.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (h *THeaderControl) SetVisible(value bool) {
    HeaderControl_SetVisible(h.instance, value)
}

// StyleElements
func (h *THeaderControl) StyleElements() TStyleElements {
    return HeaderControl_GetStyleElements(h.instance)
}

// SetStyleElements
func (h *THeaderControl) SetStyleElements(value TStyleElements) {
    HeaderControl_SetStyleElements(h.instance, value)
}

// SetOnContextPopup
func (h *THeaderControl) SetOnContextPopup(fn TContextPopupEvent) {
    HeaderControl_SetOnContextPopup(h.instance, fn)
}

// SetOnDragDrop
func (h *THeaderControl) SetOnDragDrop(fn TDragDropEvent) {
    HeaderControl_SetOnDragDrop(h.instance, fn)
}

// SetOnDragOver
func (h *THeaderControl) SetOnDragOver(fn TDragOverEvent) {
    HeaderControl_SetOnDragOver(h.instance, fn)
}

// SetOnEndDock
func (h *THeaderControl) SetOnEndDock(fn TEndDragEvent) {
    HeaderControl_SetOnEndDock(h.instance, fn)
}

// SetOnEndDrag
func (h *THeaderControl) SetOnEndDrag(fn TEndDragEvent) {
    HeaderControl_SetOnEndDrag(h.instance, fn)
}

// SetOnMouseDown
func (h *THeaderControl) SetOnMouseDown(fn TMouseEvent) {
    HeaderControl_SetOnMouseDown(h.instance, fn)
}

// SetOnMouseEnter
func (h *THeaderControl) SetOnMouseEnter(fn TNotifyEvent) {
    HeaderControl_SetOnMouseEnter(h.instance, fn)
}

// SetOnMouseLeave
func (h *THeaderControl) SetOnMouseLeave(fn TNotifyEvent) {
    HeaderControl_SetOnMouseLeave(h.instance, fn)
}

// SetOnMouseMove
func (h *THeaderControl) SetOnMouseMove(fn TMouseMoveEvent) {
    HeaderControl_SetOnMouseMove(h.instance, fn)
}

// SetOnMouseUp
func (h *THeaderControl) SetOnMouseUp(fn TMouseEvent) {
    HeaderControl_SetOnMouseUp(h.instance, fn)
}

// SetOnResize
func (h *THeaderControl) SetOnResize(fn TNotifyEvent) {
    HeaderControl_SetOnResize(h.instance, fn)
}

// SetOnDrawSection
func (h *THeaderControl) SetOnDrawSection(fn TDrawSectionEvent) {
    HeaderControl_SetOnDrawSection(h.instance, fn)
}

// SetOnSectionClick
func (h *THeaderControl) SetOnSectionClick(fn TSectionNotifyEvent) {
    HeaderControl_SetOnSectionClick(h.instance, fn)
}

// SetOnSectionResize
func (h *THeaderControl) SetOnSectionResize(fn TSectionNotifyEvent) {
    HeaderControl_SetOnSectionResize(h.instance, fn)
}

// SetOnSectionTrack
func (h *THeaderControl) SetOnSectionTrack(fn TSectionTrackEvent) {
    HeaderControl_SetOnSectionTrack(h.instance, fn)
}

// SetOnSectionDrag
func (h *THeaderControl) SetOnSectionDrag(fn TSectionDragEvent) {
    HeaderControl_SetOnSectionDrag(h.instance, fn)
}

// SetOnSectionEndDrag
func (h *THeaderControl) SetOnSectionEndDrag(fn TNotifyEvent) {
    HeaderControl_SetOnSectionEndDrag(h.instance, fn)
}

// SetOnStartDock
func (h *THeaderControl) SetOnStartDock(fn TStartDockEvent) {
    HeaderControl_SetOnStartDock(h.instance, fn)
}

// Canvas
func (h *THeaderControl) Canvas() *TCanvas {
    return CanvasFromInst(HeaderControl_GetCanvas(h.instance))
}

// SetOnSectionCheck
func (h *THeaderControl) SetOnSectionCheck(fn TCustomSectionNotifyEvent) {
    HeaderControl_SetOnSectionCheck(h.instance, fn)
}

// DockSite
func (h *THeaderControl) DockSite() bool {
    return HeaderControl_GetDockSite(h.instance)
}

// SetDockSite
func (h *THeaderControl) SetDockSite(value bool) {
    HeaderControl_SetDockSite(h.instance, value)
}

// Brush
func (h *THeaderControl) Brush() *TBrush {
    return BrushFromInst(HeaderControl_GetBrush(h.instance))
}

// ControlCount
func (h *THeaderControl) ControlCount() int32 {
    return HeaderControl_GetControlCount(h.instance)
}

// Handle
func (h *THeaderControl) Handle() HWND {
    return HeaderControl_GetHandle(h.instance)
}

// ParentWindow
func (h *THeaderControl) ParentWindow() HWND {
    return HeaderControl_GetParentWindow(h.instance)
}

// SetParentWindow
func (h *THeaderControl) SetParentWindow(value HWND) {
    HeaderControl_SetParentWindow(h.instance, value)
}

// TabOrder
func (h *THeaderControl) TabOrder() uint16 {
    return HeaderControl_GetTabOrder(h.instance)
}

// SetTabOrder
func (h *THeaderControl) SetTabOrder(value uint16) {
    HeaderControl_SetTabOrder(h.instance, value)
}

// TabStop
func (h *THeaderControl) TabStop() bool {
    return HeaderControl_GetTabStop(h.instance)
}

// SetTabStop
func (h *THeaderControl) SetTabStop(value bool) {
    HeaderControl_SetTabStop(h.instance, value)
}

// UseDockManager
func (h *THeaderControl) UseDockManager() bool {
    return HeaderControl_GetUseDockManager(h.instance)
}

// SetUseDockManager
func (h *THeaderControl) SetUseDockManager(value bool) {
    HeaderControl_SetUseDockManager(h.instance, value)
}

// Action
func (h *THeaderControl) Action() *TAction {
    return ActionFromInst(HeaderControl_GetAction(h.instance))
}

// SetAction
func (h *THeaderControl) SetAction(value IComponent) {
    HeaderControl_SetAction(h.instance, CheckPtr(value))
}

// BoundsRect
func (h *THeaderControl) BoundsRect() TRect {
    return HeaderControl_GetBoundsRect(h.instance)
}

// SetBoundsRect
func (h *THeaderControl) SetBoundsRect(value TRect) {
    HeaderControl_SetBoundsRect(h.instance, value)
}

// ClientHeight
func (h *THeaderControl) ClientHeight() int32 {
    return HeaderControl_GetClientHeight(h.instance)
}

// SetClientHeight
func (h *THeaderControl) SetClientHeight(value int32) {
    HeaderControl_SetClientHeight(h.instance, value)
}

// ClientRect
func (h *THeaderControl) ClientRect() TRect {
    return HeaderControl_GetClientRect(h.instance)
}

// ClientWidth
func (h *THeaderControl) ClientWidth() int32 {
    return HeaderControl_GetClientWidth(h.instance)
}

// SetClientWidth
func (h *THeaderControl) SetClientWidth(value int32) {
    HeaderControl_SetClientWidth(h.instance, value)
}

// ExplicitLeft
func (h *THeaderControl) ExplicitLeft() int32 {
    return HeaderControl_GetExplicitLeft(h.instance)
}

// ExplicitTop
func (h *THeaderControl) ExplicitTop() int32 {
    return HeaderControl_GetExplicitTop(h.instance)
}

// ExplicitWidth
func (h *THeaderControl) ExplicitWidth() int32 {
    return HeaderControl_GetExplicitWidth(h.instance)
}

// ExplicitHeight
func (h *THeaderControl) ExplicitHeight() int32 {
    return HeaderControl_GetExplicitHeight(h.instance)
}

// Floating
func (h *THeaderControl) Floating() bool {
    return HeaderControl_GetFloating(h.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (h *THeaderControl) Parent() *TWinControl {
    return WinControlFromInst(HeaderControl_GetParent(h.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (h *THeaderControl) SetParent(value IWinControl) {
    HeaderControl_SetParent(h.instance, CheckPtr(value))
}

// AlignWithMargins
func (h *THeaderControl) AlignWithMargins() bool {
    return HeaderControl_GetAlignWithMargins(h.instance)
}

// SetAlignWithMargins
func (h *THeaderControl) SetAlignWithMargins(value bool) {
    HeaderControl_SetAlignWithMargins(h.instance, value)
}

// Left
func (h *THeaderControl) Left() int32 {
    return HeaderControl_GetLeft(h.instance)
}

// SetLeft
func (h *THeaderControl) SetLeft(value int32) {
    HeaderControl_SetLeft(h.instance, value)
}

// Top
func (h *THeaderControl) Top() int32 {
    return HeaderControl_GetTop(h.instance)
}

// SetTop
func (h *THeaderControl) SetTop(value int32) {
    HeaderControl_SetTop(h.instance, value)
}

// Width
func (h *THeaderControl) Width() int32 {
    return HeaderControl_GetWidth(h.instance)
}

// SetWidth
func (h *THeaderControl) SetWidth(value int32) {
    HeaderControl_SetWidth(h.instance, value)
}

// Height
func (h *THeaderControl) Height() int32 {
    return HeaderControl_GetHeight(h.instance)
}

// SetHeight
func (h *THeaderControl) SetHeight(value int32) {
    HeaderControl_SetHeight(h.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (h *THeaderControl) Cursor() TCursor {
    return HeaderControl_GetCursor(h.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (h *THeaderControl) SetCursor(value TCursor) {
    HeaderControl_SetCursor(h.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (h *THeaderControl) Hint() string {
    return HeaderControl_GetHint(h.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (h *THeaderControl) SetHint(value string) {
    HeaderControl_SetHint(h.instance, value)
}

// Margins
func (h *THeaderControl) Margins() *TMargins {
    return MarginsFromInst(HeaderControl_GetMargins(h.instance))
}

// SetMargins
func (h *THeaderControl) SetMargins(value *TMargins) {
    HeaderControl_SetMargins(h.instance, CheckPtr(value))
}

// CustomHint
func (h *THeaderControl) CustomHint() *TCustomHint {
    return CustomHintFromInst(HeaderControl_GetCustomHint(h.instance))
}

// SetCustomHint
func (h *THeaderControl) SetCustomHint(value IComponent) {
    HeaderControl_SetCustomHint(h.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (h *THeaderControl) ComponentCount() int32 {
    return HeaderControl_GetComponentCount(h.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (h *THeaderControl) ComponentIndex() int32 {
    return HeaderControl_GetComponentIndex(h.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (h *THeaderControl) SetComponentIndex(value int32) {
    HeaderControl_SetComponentIndex(h.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (h *THeaderControl) Owner() *TComponent {
    return ComponentFromInst(HeaderControl_GetOwner(h.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (h *THeaderControl) Name() string {
    return HeaderControl_GetName(h.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (h *THeaderControl) SetName(value string) {
    HeaderControl_SetName(h.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (h *THeaderControl) Tag() int {
    return HeaderControl_GetTag(h.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (h *THeaderControl) SetTag(value int) {
    HeaderControl_SetTag(h.instance, value)
}

// Controls
func (h *THeaderControl) Controls(Index int32) *TControl {
    return ControlFromInst(HeaderControl_GetControls(h.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (h *THeaderControl) Components(AIndex int32) *TComponent {
    return ComponentFromInst(HeaderControl_GetComponents(h.instance, AIndex))
}

