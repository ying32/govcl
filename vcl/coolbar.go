
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

type TCoolBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCoolBar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCoolBar(owner IComponent) *TCoolBar {
    c := new(TCoolBar)
    c.instance = CoolBar_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CoolBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CoolBarFromInst(inst uintptr) *TCoolBar {
    c := new(TCoolBar)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CoolBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CoolBarFromObj(obj IObject) *TCoolBar {
    c := new(TCoolBar)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CoolBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CoolBarFromUnsafePointer(ptr unsafe.Pointer) *TCoolBar {
    c := new(TCoolBar)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCoolBar) Free() {
    if c.instance != 0 {
        CoolBar_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCoolBar) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCoolBar) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCoolBar) IsValid() bool {
    return c.instance != 0
}

// TCoolBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCoolBarClass() TClass {
    return CoolBar_StaticClassType()
}

// FlipChildren
func (c *TCoolBar) FlipChildren(AllLevels bool) {
    CoolBar_FlipChildren(c.instance, AllLevels)
}

// CanFocus
func (c *TCoolBar) CanFocus() bool {
    return CoolBar_CanFocus(c.instance)
}

// Focused
func (c *TCoolBar) Focused() bool {
    return CoolBar_Focused(c.instance)
}

// HandleAllocated
func (c *TCoolBar) HandleAllocated() bool {
    return CoolBar_HandleAllocated(c.instance)
}

// Invalidate
func (c *TCoolBar) Invalidate() {
    CoolBar_Invalidate(c.instance)
}

// Realign
func (c *TCoolBar) Realign() {
    CoolBar_Realign(c.instance)
}

// Repaint
func (c *TCoolBar) Repaint() {
    CoolBar_Repaint(c.instance)
}

// ScaleBy
func (c *TCoolBar) ScaleBy(M int32, D int32) {
    CoolBar_ScaleBy(c.instance, M , D)
}

// SetBounds
func (c *TCoolBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CoolBar_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (c *TCoolBar) SetFocus() {
    CoolBar_SetFocus(c.instance)
}

// Update
func (c *TCoolBar) Update() {
    CoolBar_Update(c.instance)
}

// BringToFront
func (c *TCoolBar) BringToFront() {
    CoolBar_BringToFront(c.instance)
}

// ClientToScreen
func (c *TCoolBar) ClientToScreen(Point TPoint) TPoint {
    return CoolBar_ClientToScreen(c.instance, Point)
}

// ClientToParent
func (c *TCoolBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CoolBar_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
func (c *TCoolBar) Dragging() bool {
    return CoolBar_Dragging(c.instance)
}

// HasParent
func (c *TCoolBar) HasParent() bool {
    return CoolBar_HasParent(c.instance)
}

// Hide
func (c *TCoolBar) Hide() {
    CoolBar_Hide(c.instance)
}

// Perform
func (c *TCoolBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CoolBar_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
func (c *TCoolBar) Refresh() {
    CoolBar_Refresh(c.instance)
}

// ScreenToClient
func (c *TCoolBar) ScreenToClient(Point TPoint) TPoint {
    return CoolBar_ScreenToClient(c.instance, Point)
}

// ParentToClient
func (c *TCoolBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CoolBar_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (c *TCoolBar) SendToBack() {
    CoolBar_SendToBack(c.instance)
}

// Show
func (c *TCoolBar) Show() {
    CoolBar_Show(c.instance)
}

// GetTextBuf
func (c *TCoolBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CoolBar_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
func (c *TCoolBar) GetTextLen() int32 {
    return CoolBar_GetTextLen(c.instance)
}

// FindComponent
func (c *TCoolBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CoolBar_FindComponent(c.instance, AName))
}

// GetNamePath
func (c *TCoolBar) GetNamePath() string {
    return CoolBar_GetNamePath(c.instance)
}

// Assign
func (c *TCoolBar) Assign(Source IObject) {
    CoolBar_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCoolBar) DisposeOf() {
    CoolBar_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCoolBar) ClassType() TClass {
    return CoolBar_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCoolBar) ClassName() string {
    return CoolBar_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCoolBar) InstanceSize() int32 {
    return CoolBar_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCoolBar) InheritsFrom(AClass TClass) bool {
    return CoolBar_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCoolBar) Equals(Obj IObject) bool {
    return CoolBar_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCoolBar) GetHashCode() int32 {
    return CoolBar_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCoolBar) ToString() string {
    return CoolBar_ToString(c.instance)
}

// Align
func (c *TCoolBar) Align() TAlign {
    return CoolBar_GetAlign(c.instance)
}

// SetAlign
func (c *TCoolBar) SetAlign(value TAlign) {
    CoolBar_SetAlign(c.instance, value)
}

// Anchors
func (c *TCoolBar) Anchors() TAnchors {
    return CoolBar_GetAnchors(c.instance)
}

// SetAnchors
func (c *TCoolBar) SetAnchors(value TAnchors) {
    CoolBar_SetAnchors(c.instance, value)
}

// AutoSize
func (c *TCoolBar) AutoSize() bool {
    return CoolBar_GetAutoSize(c.instance)
}

// SetAutoSize
func (c *TCoolBar) SetAutoSize(value bool) {
    CoolBar_SetAutoSize(c.instance, value)
}

// BandBorderStyle
func (c *TCoolBar) BandBorderStyle() TBorderStyle {
    return CoolBar_GetBandBorderStyle(c.instance)
}

// SetBandBorderStyle
func (c *TCoolBar) SetBandBorderStyle(value TBorderStyle) {
    CoolBar_SetBandBorderStyle(c.instance, value)
}

// BandMaximize
func (c *TCoolBar) BandMaximize() TCoolBandMaximize {
    return CoolBar_GetBandMaximize(c.instance)
}

// SetBandMaximize
func (c *TCoolBar) SetBandMaximize(value TCoolBandMaximize) {
    CoolBar_SetBandMaximize(c.instance, value)
}

// Bands
func (c *TCoolBar) Bands() *TCoolBands {
    return CoolBandsFromInst(CoolBar_GetBands(c.instance))
}

// SetBands
func (c *TCoolBar) SetBands(value *TCoolBands) {
    CoolBar_SetBands(c.instance, CheckPtr(value))
}

// BorderWidth
func (c *TCoolBar) BorderWidth() int32 {
    return CoolBar_GetBorderWidth(c.instance)
}

// SetBorderWidth
func (c *TCoolBar) SetBorderWidth(value int32) {
    CoolBar_SetBorderWidth(c.instance, value)
}

// Color
func (c *TCoolBar) Color() TColor {
    return CoolBar_GetColor(c.instance)
}

// SetColor
func (c *TCoolBar) SetColor(value TColor) {
    CoolBar_SetColor(c.instance, value)
}

// DockSite
func (c *TCoolBar) DockSite() bool {
    return CoolBar_GetDockSite(c.instance)
}

// SetDockSite
func (c *TCoolBar) SetDockSite(value bool) {
    CoolBar_SetDockSite(c.instance, value)
}

// DoubleBuffered
func (c *TCoolBar) DoubleBuffered() bool {
    return CoolBar_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
func (c *TCoolBar) SetDoubleBuffered(value bool) {
    CoolBar_SetDoubleBuffered(c.instance, value)
}

// DragCursor
func (c *TCoolBar) DragCursor() TCursor {
    return CoolBar_GetDragCursor(c.instance)
}

// SetDragCursor
func (c *TCoolBar) SetDragCursor(value TCursor) {
    CoolBar_SetDragCursor(c.instance, value)
}

// DragKind
func (c *TCoolBar) DragKind() TDragKind {
    return CoolBar_GetDragKind(c.instance)
}

// SetDragKind
func (c *TCoolBar) SetDragKind(value TDragKind) {
    CoolBar_SetDragKind(c.instance, value)
}

// DragMode
func (c *TCoolBar) DragMode() TDragMode {
    return CoolBar_GetDragMode(c.instance)
}

// SetDragMode
func (c *TCoolBar) SetDragMode(value TDragMode) {
    CoolBar_SetDragMode(c.instance, value)
}

// EdgeBorders
func (c *TCoolBar) EdgeBorders() TEdgeBorders {
    return CoolBar_GetEdgeBorders(c.instance)
}

// SetEdgeBorders
func (c *TCoolBar) SetEdgeBorders(value TEdgeBorders) {
    CoolBar_SetEdgeBorders(c.instance, value)
}

// EdgeInner
func (c *TCoolBar) EdgeInner() TEdgeStyle {
    return CoolBar_GetEdgeInner(c.instance)
}

// SetEdgeInner
func (c *TCoolBar) SetEdgeInner(value TEdgeStyle) {
    CoolBar_SetEdgeInner(c.instance, value)
}

// EdgeOuter
func (c *TCoolBar) EdgeOuter() TEdgeStyle {
    return CoolBar_GetEdgeOuter(c.instance)
}

// SetEdgeOuter
func (c *TCoolBar) SetEdgeOuter(value TEdgeStyle) {
    CoolBar_SetEdgeOuter(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TCoolBar) Enabled() bool {
    return CoolBar_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TCoolBar) SetEnabled(value bool) {
    CoolBar_SetEnabled(c.instance, value)
}

// FixedSize
func (c *TCoolBar) FixedSize() bool {
    return CoolBar_GetFixedSize(c.instance)
}

// SetFixedSize
func (c *TCoolBar) SetFixedSize(value bool) {
    CoolBar_SetFixedSize(c.instance, value)
}

// FixedOrder
func (c *TCoolBar) FixedOrder() bool {
    return CoolBar_GetFixedOrder(c.instance)
}

// SetFixedOrder
func (c *TCoolBar) SetFixedOrder(value bool) {
    CoolBar_SetFixedOrder(c.instance, value)
}

// Font
func (c *TCoolBar) Font() *TFont {
    return FontFromInst(CoolBar_GetFont(c.instance))
}

// SetFont
func (c *TCoolBar) SetFont(value *TFont) {
    CoolBar_SetFont(c.instance, CheckPtr(value))
}

// Images
func (c *TCoolBar) Images() *TImageList {
    return ImageListFromInst(CoolBar_GetImages(c.instance))
}

// SetImages
func (c *TCoolBar) SetImages(value IComponent) {
    CoolBar_SetImages(c.instance, CheckPtr(value))
}

// ParentColor
func (c *TCoolBar) ParentColor() bool {
    return CoolBar_GetParentColor(c.instance)
}

// SetParentColor
func (c *TCoolBar) SetParentColor(value bool) {
    CoolBar_SetParentColor(c.instance, value)
}

// ParentDoubleBuffered
func (c *TCoolBar) ParentDoubleBuffered() bool {
    return CoolBar_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
func (c *TCoolBar) SetParentDoubleBuffered(value bool) {
    CoolBar_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
func (c *TCoolBar) ParentFont() bool {
    return CoolBar_GetParentFont(c.instance)
}

// SetParentFont
func (c *TCoolBar) SetParentFont(value bool) {
    CoolBar_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TCoolBar) ParentShowHint() bool {
    return CoolBar_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TCoolBar) SetParentShowHint(value bool) {
    CoolBar_SetParentShowHint(c.instance, value)
}

// Bitmap
func (c *TCoolBar) Bitmap() *TBitmap {
    return BitmapFromInst(CoolBar_GetBitmap(c.instance))
}

// SetBitmap
func (c *TCoolBar) SetBitmap(value *TBitmap) {
    CoolBar_SetBitmap(c.instance, CheckPtr(value))
}

// PopupMenu
func (c *TCoolBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CoolBar_GetPopupMenu(c.instance))
}

// SetPopupMenu
func (c *TCoolBar) SetPopupMenu(value IComponent) {
    CoolBar_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
func (c *TCoolBar) ShowHint() bool {
    return CoolBar_GetShowHint(c.instance)
}

// SetShowHint
func (c *TCoolBar) SetShowHint(value bool) {
    CoolBar_SetShowHint(c.instance, value)
}

// ShowText
func (c *TCoolBar) ShowText() bool {
    return CoolBar_GetShowText(c.instance)
}

// SetShowText
func (c *TCoolBar) SetShowText(value bool) {
    CoolBar_SetShowText(c.instance, value)
}

// Vertical
func (c *TCoolBar) Vertical() bool {
    return CoolBar_GetVertical(c.instance)
}

// SetVertical
func (c *TCoolBar) SetVertical(value bool) {
    CoolBar_SetVertical(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TCoolBar) Visible() bool {
    return CoolBar_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TCoolBar) SetVisible(value bool) {
    CoolBar_SetVisible(c.instance, value)
}

// StyleElements
func (c *TCoolBar) StyleElements() TStyleElements {
    return CoolBar_GetStyleElements(c.instance)
}

// SetStyleElements
func (c *TCoolBar) SetStyleElements(value TStyleElements) {
    CoolBar_SetStyleElements(c.instance, value)
}

// SetOnChange
func (c *TCoolBar) SetOnChange(fn TNotifyEvent) {
    CoolBar_SetOnChange(c.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TCoolBar) SetOnClick(fn TNotifyEvent) {
    CoolBar_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
func (c *TCoolBar) SetOnContextPopup(fn TContextPopupEvent) {
    CoolBar_SetOnContextPopup(c.instance, fn)
}

// SetOnDblClick
func (c *TCoolBar) SetOnDblClick(fn TNotifyEvent) {
    CoolBar_SetOnDblClick(c.instance, fn)
}

// SetOnDockDrop
func (c *TCoolBar) SetOnDockDrop(fn TDockDropEvent) {
    CoolBar_SetOnDockDrop(c.instance, fn)
}

// SetOnDragDrop
func (c *TCoolBar) SetOnDragDrop(fn TDragDropEvent) {
    CoolBar_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
func (c *TCoolBar) SetOnDragOver(fn TDragOverEvent) {
    CoolBar_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
func (c *TCoolBar) SetOnEndDock(fn TEndDragEvent) {
    CoolBar_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
func (c *TCoolBar) SetOnEndDrag(fn TEndDragEvent) {
    CoolBar_SetOnEndDrag(c.instance, fn)
}

// SetOnGetSiteInfo
func (c *TCoolBar) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    CoolBar_SetOnGetSiteInfo(c.instance, fn)
}

// SetOnMouseDown
func (c *TCoolBar) SetOnMouseDown(fn TMouseEvent) {
    CoolBar_SetOnMouseDown(c.instance, fn)
}

// SetOnMouseEnter
func (c *TCoolBar) SetOnMouseEnter(fn TNotifyEvent) {
    CoolBar_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
func (c *TCoolBar) SetOnMouseLeave(fn TNotifyEvent) {
    CoolBar_SetOnMouseLeave(c.instance, fn)
}

// SetOnMouseMove
func (c *TCoolBar) SetOnMouseMove(fn TMouseMoveEvent) {
    CoolBar_SetOnMouseMove(c.instance, fn)
}

// SetOnMouseUp
func (c *TCoolBar) SetOnMouseUp(fn TMouseEvent) {
    CoolBar_SetOnMouseUp(c.instance, fn)
}

// SetOnResize
func (c *TCoolBar) SetOnResize(fn TNotifyEvent) {
    CoolBar_SetOnResize(c.instance, fn)
}

// SetOnStartDock
func (c *TCoolBar) SetOnStartDock(fn TStartDockEvent) {
    CoolBar_SetOnStartDock(c.instance, fn)
}

// SetOnUnDock
func (c *TCoolBar) SetOnUnDock(fn TUnDockEvent) {
    CoolBar_SetOnUnDock(c.instance, fn)
}

// Brush
func (c *TCoolBar) Brush() *TBrush {
    return BrushFromInst(CoolBar_GetBrush(c.instance))
}

// ControlCount
func (c *TCoolBar) ControlCount() int32 {
    return CoolBar_GetControlCount(c.instance)
}

// Handle
func (c *TCoolBar) Handle() HWND {
    return CoolBar_GetHandle(c.instance)
}

// ParentWindow
func (c *TCoolBar) ParentWindow() HWND {
    return CoolBar_GetParentWindow(c.instance)
}

// SetParentWindow
func (c *TCoolBar) SetParentWindow(value HWND) {
    CoolBar_SetParentWindow(c.instance, value)
}

// TabOrder
func (c *TCoolBar) TabOrder() uint16 {
    return CoolBar_GetTabOrder(c.instance)
}

// SetTabOrder
func (c *TCoolBar) SetTabOrder(value uint16) {
    CoolBar_SetTabOrder(c.instance, value)
}

// TabStop
func (c *TCoolBar) TabStop() bool {
    return CoolBar_GetTabStop(c.instance)
}

// SetTabStop
func (c *TCoolBar) SetTabStop(value bool) {
    CoolBar_SetTabStop(c.instance, value)
}

// UseDockManager
func (c *TCoolBar) UseDockManager() bool {
    return CoolBar_GetUseDockManager(c.instance)
}

// SetUseDockManager
func (c *TCoolBar) SetUseDockManager(value bool) {
    CoolBar_SetUseDockManager(c.instance, value)
}

// Action
func (c *TCoolBar) Action() *TAction {
    return ActionFromInst(CoolBar_GetAction(c.instance))
}

// SetAction
func (c *TCoolBar) SetAction(value IComponent) {
    CoolBar_SetAction(c.instance, CheckPtr(value))
}

// BiDiMode
func (c *TCoolBar) BiDiMode() TBiDiMode {
    return CoolBar_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TCoolBar) SetBiDiMode(value TBiDiMode) {
    CoolBar_SetBiDiMode(c.instance, value)
}

// BoundsRect
func (c *TCoolBar) BoundsRect() TRect {
    return CoolBar_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TCoolBar) SetBoundsRect(value TRect) {
    CoolBar_SetBoundsRect(c.instance, value)
}

// ClientHeight
func (c *TCoolBar) ClientHeight() int32 {
    return CoolBar_GetClientHeight(c.instance)
}

// SetClientHeight
func (c *TCoolBar) SetClientHeight(value int32) {
    CoolBar_SetClientHeight(c.instance, value)
}

// ClientRect
func (c *TCoolBar) ClientRect() TRect {
    return CoolBar_GetClientRect(c.instance)
}

// ClientWidth
func (c *TCoolBar) ClientWidth() int32 {
    return CoolBar_GetClientWidth(c.instance)
}

// SetClientWidth
func (c *TCoolBar) SetClientWidth(value int32) {
    CoolBar_SetClientWidth(c.instance, value)
}

// ExplicitLeft
func (c *TCoolBar) ExplicitLeft() int32 {
    return CoolBar_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TCoolBar) ExplicitTop() int32 {
    return CoolBar_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TCoolBar) ExplicitWidth() int32 {
    return CoolBar_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TCoolBar) ExplicitHeight() int32 {
    return CoolBar_GetExplicitHeight(c.instance)
}

// Floating
func (c *TCoolBar) Floating() bool {
    return CoolBar_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TCoolBar) Parent() *TWinControl {
    return WinControlFromInst(CoolBar_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TCoolBar) SetParent(value IWinControl) {
    CoolBar_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
func (c *TCoolBar) AlignWithMargins() bool {
    return CoolBar_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
func (c *TCoolBar) SetAlignWithMargins(value bool) {
    CoolBar_SetAlignWithMargins(c.instance, value)
}

// Left
func (c *TCoolBar) Left() int32 {
    return CoolBar_GetLeft(c.instance)
}

// SetLeft
func (c *TCoolBar) SetLeft(value int32) {
    CoolBar_SetLeft(c.instance, value)
}

// Top
func (c *TCoolBar) Top() int32 {
    return CoolBar_GetTop(c.instance)
}

// SetTop
func (c *TCoolBar) SetTop(value int32) {
    CoolBar_SetTop(c.instance, value)
}

// Width
func (c *TCoolBar) Width() int32 {
    return CoolBar_GetWidth(c.instance)
}

// SetWidth
func (c *TCoolBar) SetWidth(value int32) {
    CoolBar_SetWidth(c.instance, value)
}

// Height
func (c *TCoolBar) Height() int32 {
    return CoolBar_GetHeight(c.instance)
}

// SetHeight
func (c *TCoolBar) SetHeight(value int32) {
    CoolBar_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TCoolBar) Cursor() TCursor {
    return CoolBar_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TCoolBar) SetCursor(value TCursor) {
    CoolBar_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (c *TCoolBar) Hint() string {
    return CoolBar_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (c *TCoolBar) SetHint(value string) {
    CoolBar_SetHint(c.instance, value)
}

// Margins
func (c *TCoolBar) Margins() *TMargins {
    return MarginsFromInst(CoolBar_GetMargins(c.instance))
}

// SetMargins
func (c *TCoolBar) SetMargins(value *TMargins) {
    CoolBar_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
func (c *TCoolBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(CoolBar_GetCustomHint(c.instance))
}

// SetCustomHint
func (c *TCoolBar) SetCustomHint(value IComponent) {
    CoolBar_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TCoolBar) ComponentCount() int32 {
    return CoolBar_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TCoolBar) ComponentIndex() int32 {
    return CoolBar_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TCoolBar) SetComponentIndex(value int32) {
    CoolBar_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TCoolBar) Owner() *TComponent {
    return ComponentFromInst(CoolBar_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TCoolBar) Name() string {
    return CoolBar_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TCoolBar) SetName(value string) {
    CoolBar_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TCoolBar) Tag() int {
    return CoolBar_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TCoolBar) SetTag(value int) {
    CoolBar_SetTag(c.instance, value)
}

// Controls
func (c *TCoolBar) Controls(Index int32) *TControl {
    return ControlFromInst(CoolBar_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TCoolBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CoolBar_GetComponents(c.instance, AIndex))
}

