
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

type TPageControl struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPageControl
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPageControl(owner IComponent) *TPageControl {
    p := new(TPageControl)
    p.instance = PageControl_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PageControlFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PageControlFromInst(inst uintptr) *TPageControl {
    p := new(TPageControl)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PageControlFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PageControlFromObj(obj IObject) *TPageControl {
    p := new(TPageControl)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PageControlFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PageControlFromUnsafePointer(ptr unsafe.Pointer) *TPageControl {
    p := new(TPageControl)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPageControl) Free() {
    if p.instance != 0 {
        PageControl_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPageControl) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPageControl) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPageControl) IsValid() bool {
    return p.instance != 0
}

// TPageControlClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPageControlClass() TClass {
    return PageControl_StaticClassType()
}

// SelectNextPage
func (p *TPageControl) SelectNextPage(GoForward bool, CheckTabVisible bool) {
    PageControl_SelectNextPage(p.instance, GoForward , CheckTabVisible)
}

// TabRect
func (p *TPageControl) TabRect(Index int32) TRect {
    return PageControl_TabRect(p.instance, Index)
}

// RowCount
func (p *TPageControl) RowCount() int32 {
    return PageControl_RowCount(p.instance)
}

// CanFocus
func (p *TPageControl) CanFocus() bool {
    return PageControl_CanFocus(p.instance)
}

// FlipChildren
func (p *TPageControl) FlipChildren(AllLevels bool) {
    PageControl_FlipChildren(p.instance, AllLevels)
}

// Focused
func (p *TPageControl) Focused() bool {
    return PageControl_Focused(p.instance)
}

// HandleAllocated
func (p *TPageControl) HandleAllocated() bool {
    return PageControl_HandleAllocated(p.instance)
}

// Invalidate
func (p *TPageControl) Invalidate() {
    PageControl_Invalidate(p.instance)
}

// Realign
func (p *TPageControl) Realign() {
    PageControl_Realign(p.instance)
}

// Repaint
func (p *TPageControl) Repaint() {
    PageControl_Repaint(p.instance)
}

// ScaleBy
func (p *TPageControl) ScaleBy(M int32, D int32) {
    PageControl_ScaleBy(p.instance, M , D)
}

// SetBounds
func (p *TPageControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    PageControl_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (p *TPageControl) SetFocus() {
    PageControl_SetFocus(p.instance)
}

// Update
func (p *TPageControl) Update() {
    PageControl_Update(p.instance)
}

// BringToFront
func (p *TPageControl) BringToFront() {
    PageControl_BringToFront(p.instance)
}

// ClientToScreen
func (p *TPageControl) ClientToScreen(Point TPoint) TPoint {
    return PageControl_ClientToScreen(p.instance, Point)
}

// ClientToParent
func (p *TPageControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return PageControl_ClientToParent(p.instance, Point , CheckPtr(AParent))
}

// Dragging
func (p *TPageControl) Dragging() bool {
    return PageControl_Dragging(p.instance)
}

// HasParent
func (p *TPageControl) HasParent() bool {
    return PageControl_HasParent(p.instance)
}

// Hide
func (p *TPageControl) Hide() {
    PageControl_Hide(p.instance)
}

// Perform
func (p *TPageControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return PageControl_Perform(p.instance, Msg , WParam , LParam)
}

// Refresh
func (p *TPageControl) Refresh() {
    PageControl_Refresh(p.instance)
}

// ScreenToClient
func (p *TPageControl) ScreenToClient(Point TPoint) TPoint {
    return PageControl_ScreenToClient(p.instance, Point)
}

// ParentToClient
func (p *TPageControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return PageControl_ParentToClient(p.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (p *TPageControl) SendToBack() {
    PageControl_SendToBack(p.instance)
}

// Show
func (p *TPageControl) Show() {
    PageControl_Show(p.instance)
}

// GetTextBuf
func (p *TPageControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    return PageControl_GetTextBuf(p.instance, Buffer , BufSize)
}

// GetTextLen
func (p *TPageControl) GetTextLen() int32 {
    return PageControl_GetTextLen(p.instance)
}

// FindComponent
func (p *TPageControl) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PageControl_FindComponent(p.instance, AName))
}

// GetNamePath
func (p *TPageControl) GetNamePath() string {
    return PageControl_GetNamePath(p.instance)
}

// Assign
func (p *TPageControl) Assign(Source IObject) {
    PageControl_Assign(p.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPageControl) DisposeOf() {
    PageControl_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPageControl) ClassType() TClass {
    return PageControl_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPageControl) ClassName() string {
    return PageControl_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPageControl) InstanceSize() int32 {
    return PageControl_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPageControl) InheritsFrom(AClass TClass) bool {
    return PageControl_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPageControl) Equals(Obj IObject) bool {
    return PageControl_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPageControl) GetHashCode() int32 {
    return PageControl_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPageControl) ToString() string {
    return PageControl_ToString(p.instance)
}

// ActivePageIndex
func (p *TPageControl) ActivePageIndex() int32 {
    return PageControl_GetActivePageIndex(p.instance)
}

// SetActivePageIndex
func (p *TPageControl) SetActivePageIndex(value int32) {
    PageControl_SetActivePageIndex(p.instance, value)
}

// PageCount
func (p *TPageControl) PageCount() int32 {
    return PageControl_GetPageCount(p.instance)
}

// Align
func (p *TPageControl) Align() TAlign {
    return PageControl_GetAlign(p.instance)
}

// SetAlign
func (p *TPageControl) SetAlign(value TAlign) {
    PageControl_SetAlign(p.instance, value)
}

// Anchors
func (p *TPageControl) Anchors() TAnchors {
    return PageControl_GetAnchors(p.instance)
}

// SetAnchors
func (p *TPageControl) SetAnchors(value TAnchors) {
    PageControl_SetAnchors(p.instance, value)
}

// BiDiMode
func (p *TPageControl) BiDiMode() TBiDiMode {
    return PageControl_GetBiDiMode(p.instance)
}

// SetBiDiMode
func (p *TPageControl) SetBiDiMode(value TBiDiMode) {
    PageControl_SetBiDiMode(p.instance, value)
}

// DockSite
func (p *TPageControl) DockSite() bool {
    return PageControl_GetDockSite(p.instance)
}

// SetDockSite
func (p *TPageControl) SetDockSite(value bool) {
    PageControl_SetDockSite(p.instance, value)
}

// DoubleBuffered
func (p *TPageControl) DoubleBuffered() bool {
    return PageControl_GetDoubleBuffered(p.instance)
}

// SetDoubleBuffered
func (p *TPageControl) SetDoubleBuffered(value bool) {
    PageControl_SetDoubleBuffered(p.instance, value)
}

// DragCursor
func (p *TPageControl) DragCursor() TCursor {
    return PageControl_GetDragCursor(p.instance)
}

// SetDragCursor
func (p *TPageControl) SetDragCursor(value TCursor) {
    PageControl_SetDragCursor(p.instance, value)
}

// DragKind
func (p *TPageControl) DragKind() TDragKind {
    return PageControl_GetDragKind(p.instance)
}

// SetDragKind
func (p *TPageControl) SetDragKind(value TDragKind) {
    PageControl_SetDragKind(p.instance, value)
}

// DragMode
func (p *TPageControl) DragMode() TDragMode {
    return PageControl_GetDragMode(p.instance)
}

// SetDragMode
func (p *TPageControl) SetDragMode(value TDragMode) {
    PageControl_SetDragMode(p.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (p *TPageControl) Enabled() bool {
    return PageControl_GetEnabled(p.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (p *TPageControl) SetEnabled(value bool) {
    PageControl_SetEnabled(p.instance, value)
}

// Font
func (p *TPageControl) Font() *TFont {
    return FontFromInst(PageControl_GetFont(p.instance))
}

// SetFont
func (p *TPageControl) SetFont(value *TFont) {
    PageControl_SetFont(p.instance, CheckPtr(value))
}

// HotTrack
func (p *TPageControl) HotTrack() bool {
    return PageControl_GetHotTrack(p.instance)
}

// SetHotTrack
func (p *TPageControl) SetHotTrack(value bool) {
    PageControl_SetHotTrack(p.instance, value)
}

// Images
func (p *TPageControl) Images() *TImageList {
    return ImageListFromInst(PageControl_GetImages(p.instance))
}

// SetImages
func (p *TPageControl) SetImages(value IComponent) {
    PageControl_SetImages(p.instance, CheckPtr(value))
}

// MultiLine
func (p *TPageControl) MultiLine() bool {
    return PageControl_GetMultiLine(p.instance)
}

// SetMultiLine
func (p *TPageControl) SetMultiLine(value bool) {
    PageControl_SetMultiLine(p.instance, value)
}

// ParentDoubleBuffered
func (p *TPageControl) ParentDoubleBuffered() bool {
    return PageControl_GetParentDoubleBuffered(p.instance)
}

// SetParentDoubleBuffered
func (p *TPageControl) SetParentDoubleBuffered(value bool) {
    PageControl_SetParentDoubleBuffered(p.instance, value)
}

// ParentFont
func (p *TPageControl) ParentFont() bool {
    return PageControl_GetParentFont(p.instance)
}

// SetParentFont
func (p *TPageControl) SetParentFont(value bool) {
    PageControl_SetParentFont(p.instance, value)
}

// ParentShowHint
func (p *TPageControl) ParentShowHint() bool {
    return PageControl_GetParentShowHint(p.instance)
}

// SetParentShowHint
func (p *TPageControl) SetParentShowHint(value bool) {
    PageControl_SetParentShowHint(p.instance, value)
}

// PopupMenu
func (p *TPageControl) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(PageControl_GetPopupMenu(p.instance))
}

// SetPopupMenu
func (p *TPageControl) SetPopupMenu(value IComponent) {
    PageControl_SetPopupMenu(p.instance, CheckPtr(value))
}

// ShowHint
func (p *TPageControl) ShowHint() bool {
    return PageControl_GetShowHint(p.instance)
}

// SetShowHint
func (p *TPageControl) SetShowHint(value bool) {
    PageControl_SetShowHint(p.instance, value)
}

// Style
func (p *TPageControl) Style() TTabStyle {
    return PageControl_GetStyle(p.instance)
}

// SetStyle
func (p *TPageControl) SetStyle(value TTabStyle) {
    PageControl_SetStyle(p.instance, value)
}

// TabHeight
func (p *TPageControl) TabHeight() int16 {
    return PageControl_GetTabHeight(p.instance)
}

// SetTabHeight
func (p *TPageControl) SetTabHeight(value int16) {
    PageControl_SetTabHeight(p.instance, value)
}

// TabIndex
func (p *TPageControl) TabIndex() int32 {
    return PageControl_GetTabIndex(p.instance)
}

// SetTabIndex
func (p *TPageControl) SetTabIndex(value int32) {
    PageControl_SetTabIndex(p.instance, value)
}

// TabOrder
func (p *TPageControl) TabOrder() uint16 {
    return PageControl_GetTabOrder(p.instance)
}

// SetTabOrder
func (p *TPageControl) SetTabOrder(value uint16) {
    PageControl_SetTabOrder(p.instance, value)
}

// TabPosition
func (p *TPageControl) TabPosition() TTabPosition {
    return PageControl_GetTabPosition(p.instance)
}

// SetTabPosition
func (p *TPageControl) SetTabPosition(value TTabPosition) {
    PageControl_SetTabPosition(p.instance, value)
}

// TabStop
func (p *TPageControl) TabStop() bool {
    return PageControl_GetTabStop(p.instance)
}

// SetTabStop
func (p *TPageControl) SetTabStop(value bool) {
    PageControl_SetTabStop(p.instance, value)
}

// TabWidth
func (p *TPageControl) TabWidth() int16 {
    return PageControl_GetTabWidth(p.instance)
}

// SetTabWidth
func (p *TPageControl) SetTabWidth(value int16) {
    PageControl_SetTabWidth(p.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (p *TPageControl) Visible() bool {
    return PageControl_GetVisible(p.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (p *TPageControl) SetVisible(value bool) {
    PageControl_SetVisible(p.instance, value)
}

// StyleElements
func (p *TPageControl) StyleElements() TStyleElements {
    return PageControl_GetStyleElements(p.instance)
}

// SetStyleElements
func (p *TPageControl) SetStyleElements(value TStyleElements) {
    PageControl_SetStyleElements(p.instance, value)
}

// SetOnChange
func (p *TPageControl) SetOnChange(fn TNotifyEvent) {
    PageControl_SetOnChange(p.instance, fn)
}

// SetOnContextPopup
func (p *TPageControl) SetOnContextPopup(fn TContextPopupEvent) {
    PageControl_SetOnContextPopup(p.instance, fn)
}

// SetOnDockDrop
func (p *TPageControl) SetOnDockDrop(fn TDockDropEvent) {
    PageControl_SetOnDockDrop(p.instance, fn)
}

// SetOnDragDrop
func (p *TPageControl) SetOnDragDrop(fn TDragDropEvent) {
    PageControl_SetOnDragDrop(p.instance, fn)
}

// SetOnDragOver
func (p *TPageControl) SetOnDragOver(fn TDragOverEvent) {
    PageControl_SetOnDragOver(p.instance, fn)
}

// SetOnEndDock
func (p *TPageControl) SetOnEndDock(fn TEndDragEvent) {
    PageControl_SetOnEndDock(p.instance, fn)
}

// SetOnEndDrag
func (p *TPageControl) SetOnEndDrag(fn TEndDragEvent) {
    PageControl_SetOnEndDrag(p.instance, fn)
}

// SetOnEnter
func (p *TPageControl) SetOnEnter(fn TNotifyEvent) {
    PageControl_SetOnEnter(p.instance, fn)
}

// SetOnExit
func (p *TPageControl) SetOnExit(fn TNotifyEvent) {
    PageControl_SetOnExit(p.instance, fn)
}

// SetOnGetImageIndex
func (p *TPageControl) SetOnGetImageIndex(fn TTabGetImageEvent) {
    PageControl_SetOnGetImageIndex(p.instance, fn)
}

// SetOnGetSiteInfo
func (p *TPageControl) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    PageControl_SetOnGetSiteInfo(p.instance, fn)
}

// SetOnMouseDown
func (p *TPageControl) SetOnMouseDown(fn TMouseEvent) {
    PageControl_SetOnMouseDown(p.instance, fn)
}

// SetOnMouseEnter
func (p *TPageControl) SetOnMouseEnter(fn TNotifyEvent) {
    PageControl_SetOnMouseEnter(p.instance, fn)
}

// SetOnMouseLeave
func (p *TPageControl) SetOnMouseLeave(fn TNotifyEvent) {
    PageControl_SetOnMouseLeave(p.instance, fn)
}

// SetOnMouseMove
func (p *TPageControl) SetOnMouseMove(fn TMouseMoveEvent) {
    PageControl_SetOnMouseMove(p.instance, fn)
}

// SetOnMouseUp
func (p *TPageControl) SetOnMouseUp(fn TMouseEvent) {
    PageControl_SetOnMouseUp(p.instance, fn)
}

// SetOnResize
func (p *TPageControl) SetOnResize(fn TNotifyEvent) {
    PageControl_SetOnResize(p.instance, fn)
}

// SetOnStartDock
func (p *TPageControl) SetOnStartDock(fn TStartDockEvent) {
    PageControl_SetOnStartDock(p.instance, fn)
}

// SetOnUnDock
func (p *TPageControl) SetOnUnDock(fn TUnDockEvent) {
    PageControl_SetOnUnDock(p.instance, fn)
}

// Canvas
func (p *TPageControl) Canvas() *TCanvas {
    return CanvasFromInst(PageControl_GetCanvas(p.instance))
}

// Brush
func (p *TPageControl) Brush() *TBrush {
    return BrushFromInst(PageControl_GetBrush(p.instance))
}

// ControlCount
func (p *TPageControl) ControlCount() int32 {
    return PageControl_GetControlCount(p.instance)
}

// Handle
func (p *TPageControl) Handle() HWND {
    return PageControl_GetHandle(p.instance)
}

// ParentWindow
func (p *TPageControl) ParentWindow() HWND {
    return PageControl_GetParentWindow(p.instance)
}

// SetParentWindow
func (p *TPageControl) SetParentWindow(value HWND) {
    PageControl_SetParentWindow(p.instance, value)
}

// UseDockManager
func (p *TPageControl) UseDockManager() bool {
    return PageControl_GetUseDockManager(p.instance)
}

// SetUseDockManager
func (p *TPageControl) SetUseDockManager(value bool) {
    PageControl_SetUseDockManager(p.instance, value)
}

// Action
func (p *TPageControl) Action() *TAction {
    return ActionFromInst(PageControl_GetAction(p.instance))
}

// SetAction
func (p *TPageControl) SetAction(value IComponent) {
    PageControl_SetAction(p.instance, CheckPtr(value))
}

// BoundsRect
func (p *TPageControl) BoundsRect() TRect {
    return PageControl_GetBoundsRect(p.instance)
}

// SetBoundsRect
func (p *TPageControl) SetBoundsRect(value TRect) {
    PageControl_SetBoundsRect(p.instance, value)
}

// ClientHeight
func (p *TPageControl) ClientHeight() int32 {
    return PageControl_GetClientHeight(p.instance)
}

// SetClientHeight
func (p *TPageControl) SetClientHeight(value int32) {
    PageControl_SetClientHeight(p.instance, value)
}

// ClientRect
func (p *TPageControl) ClientRect() TRect {
    return PageControl_GetClientRect(p.instance)
}

// ClientWidth
func (p *TPageControl) ClientWidth() int32 {
    return PageControl_GetClientWidth(p.instance)
}

// SetClientWidth
func (p *TPageControl) SetClientWidth(value int32) {
    PageControl_SetClientWidth(p.instance, value)
}

// ExplicitLeft
func (p *TPageControl) ExplicitLeft() int32 {
    return PageControl_GetExplicitLeft(p.instance)
}

// ExplicitTop
func (p *TPageControl) ExplicitTop() int32 {
    return PageControl_GetExplicitTop(p.instance)
}

// ExplicitWidth
func (p *TPageControl) ExplicitWidth() int32 {
    return PageControl_GetExplicitWidth(p.instance)
}

// ExplicitHeight
func (p *TPageControl) ExplicitHeight() int32 {
    return PageControl_GetExplicitHeight(p.instance)
}

// Floating
func (p *TPageControl) Floating() bool {
    return PageControl_GetFloating(p.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (p *TPageControl) Parent() *TWinControl {
    return WinControlFromInst(PageControl_GetParent(p.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (p *TPageControl) SetParent(value IWinControl) {
    PageControl_SetParent(p.instance, CheckPtr(value))
}

// AlignWithMargins
func (p *TPageControl) AlignWithMargins() bool {
    return PageControl_GetAlignWithMargins(p.instance)
}

// SetAlignWithMargins
func (p *TPageControl) SetAlignWithMargins(value bool) {
    PageControl_SetAlignWithMargins(p.instance, value)
}

// Left
func (p *TPageControl) Left() int32 {
    return PageControl_GetLeft(p.instance)
}

// SetLeft
func (p *TPageControl) SetLeft(value int32) {
    PageControl_SetLeft(p.instance, value)
}

// Top
func (p *TPageControl) Top() int32 {
    return PageControl_GetTop(p.instance)
}

// SetTop
func (p *TPageControl) SetTop(value int32) {
    PageControl_SetTop(p.instance, value)
}

// Width
func (p *TPageControl) Width() int32 {
    return PageControl_GetWidth(p.instance)
}

// SetWidth
func (p *TPageControl) SetWidth(value int32) {
    PageControl_SetWidth(p.instance, value)
}

// Height
func (p *TPageControl) Height() int32 {
    return PageControl_GetHeight(p.instance)
}

// SetHeight
func (p *TPageControl) SetHeight(value int32) {
    PageControl_SetHeight(p.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (p *TPageControl) Cursor() TCursor {
    return PageControl_GetCursor(p.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (p *TPageControl) SetCursor(value TCursor) {
    PageControl_SetCursor(p.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (p *TPageControl) Hint() string {
    return PageControl_GetHint(p.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (p *TPageControl) SetHint(value string) {
    PageControl_SetHint(p.instance, value)
}

// Margins
func (p *TPageControl) Margins() *TMargins {
    return MarginsFromInst(PageControl_GetMargins(p.instance))
}

// SetMargins
func (p *TPageControl) SetMargins(value *TMargins) {
    PageControl_SetMargins(p.instance, CheckPtr(value))
}

// CustomHint
func (p *TPageControl) CustomHint() *TCustomHint {
    return CustomHintFromInst(PageControl_GetCustomHint(p.instance))
}

// SetCustomHint
func (p *TPageControl) SetCustomHint(value IComponent) {
    PageControl_SetCustomHint(p.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPageControl) ComponentCount() int32 {
    return PageControl_GetComponentCount(p.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (p *TPageControl) ComponentIndex() int32 {
    return PageControl_GetComponentIndex(p.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (p *TPageControl) SetComponentIndex(value int32) {
    PageControl_SetComponentIndex(p.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPageControl) Owner() *TComponent {
    return ComponentFromInst(PageControl_GetOwner(p.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPageControl) Name() string {
    return PageControl_GetName(p.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPageControl) SetName(value string) {
    PageControl_SetName(p.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPageControl) Tag() int {
    return PageControl_GetTag(p.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPageControl) SetTag(value int) {
    PageControl_SetTag(p.instance, value)
}

// Pages
func (p *TPageControl) Pages(Index int32) *TTabSheet {
    return TabSheetFromInst(PageControl_GetPages(p.instance, Index))
}

// Controls
func (p *TPageControl) Controls(Index int32) *TControl {
    return ControlFromInst(PageControl_GetControls(p.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPageControl) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PageControl_GetComponents(p.instance, AIndex))
}

