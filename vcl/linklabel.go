
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

type TLinkLabel struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewLinkLabel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewLinkLabel(owner IComponent) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = LinkLabel_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// LinkLabelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func LinkLabelFromInst(inst uintptr) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// LinkLabelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func LinkLabelFromObj(obj IObject) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// LinkLabelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func LinkLabelFromUnsafePointer(ptr unsafe.Pointer) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TLinkLabel) Free() {
    if l.instance != 0 {
        LinkLabel_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TLinkLabel) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TLinkLabel) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TLinkLabel) IsValid() bool {
    return l.instance != 0
}

// TLinkLabelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TLinkLabelClass() TClass {
    return LinkLabel_StaticClassType()
}

// CanFocus
func (l *TLinkLabel) CanFocus() bool {
    return LinkLabel_CanFocus(l.instance)
}

// ContainsControl
func (l *TLinkLabel) ContainsControl(Control IControl) bool {
    return LinkLabel_ContainsControl(l.instance, CheckPtr(Control))
}

// ControlAtPos
func (l *TLinkLabel) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(LinkLabel_ControlAtPos(l.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (l *TLinkLabel) DisableAlign() {
    LinkLabel_DisableAlign(l.instance)
}

// EnableAlign
func (l *TLinkLabel) EnableAlign() {
    LinkLabel_EnableAlign(l.instance)
}

// FindChildControl
func (l *TLinkLabel) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(LinkLabel_FindChildControl(l.instance, ControlName))
}

// FlipChildren
func (l *TLinkLabel) FlipChildren(AllLevels bool) {
    LinkLabel_FlipChildren(l.instance, AllLevels)
}

// Focused
func (l *TLinkLabel) Focused() bool {
    return LinkLabel_Focused(l.instance)
}

// HandleAllocated
func (l *TLinkLabel) HandleAllocated() bool {
    return LinkLabel_HandleAllocated(l.instance)
}

// InsertControl
func (l *TLinkLabel) InsertControl(AControl IControl) {
    LinkLabel_InsertControl(l.instance, CheckPtr(AControl))
}

// Invalidate
func (l *TLinkLabel) Invalidate() {
    LinkLabel_Invalidate(l.instance)
}

// PaintTo
func (l *TLinkLabel) PaintTo(DC HDC, X int32, Y int32) {
    LinkLabel_PaintTo(l.instance, DC , X , Y)
}

// RemoveControl
func (l *TLinkLabel) RemoveControl(AControl IControl) {
    LinkLabel_RemoveControl(l.instance, CheckPtr(AControl))
}

// Realign
func (l *TLinkLabel) Realign() {
    LinkLabel_Realign(l.instance)
}

// Repaint
func (l *TLinkLabel) Repaint() {
    LinkLabel_Repaint(l.instance)
}

// ScaleBy
func (l *TLinkLabel) ScaleBy(M int32, D int32) {
    LinkLabel_ScaleBy(l.instance, M , D)
}

// ScrollBy
func (l *TLinkLabel) ScrollBy(DeltaX int32, DeltaY int32) {
    LinkLabel_ScrollBy(l.instance, DeltaX , DeltaY)
}

// SetBounds
func (l *TLinkLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    LinkLabel_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (l *TLinkLabel) SetFocus() {
    LinkLabel_SetFocus(l.instance)
}

// Update
func (l *TLinkLabel) Update() {
    LinkLabel_Update(l.instance)
}

// UpdateControlState
func (l *TLinkLabel) UpdateControlState() {
    LinkLabel_UpdateControlState(l.instance)
}

// BringToFront
func (l *TLinkLabel) BringToFront() {
    LinkLabel_BringToFront(l.instance)
}

// ClientToScreen
func (l *TLinkLabel) ClientToScreen(Point TPoint) TPoint {
    return LinkLabel_ClientToScreen(l.instance, Point)
}

// ClientToParent
func (l *TLinkLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return LinkLabel_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// Dragging
func (l *TLinkLabel) Dragging() bool {
    return LinkLabel_Dragging(l.instance)
}

// HasParent
func (l *TLinkLabel) HasParent() bool {
    return LinkLabel_HasParent(l.instance)
}

// Hide
func (l *TLinkLabel) Hide() {
    LinkLabel_Hide(l.instance)
}

// Perform
func (l *TLinkLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return LinkLabel_Perform(l.instance, Msg , WParam , LParam)
}

// Refresh
func (l *TLinkLabel) Refresh() {
    LinkLabel_Refresh(l.instance)
}

// ScreenToClient
func (l *TLinkLabel) ScreenToClient(Point TPoint) TPoint {
    return LinkLabel_ScreenToClient(l.instance, Point)
}

// ParentToClient
func (l *TLinkLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return LinkLabel_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (l *TLinkLabel) SendToBack() {
    LinkLabel_SendToBack(l.instance)
}

// Show
func (l *TLinkLabel) Show() {
    LinkLabel_Show(l.instance)
}

// GetTextBuf
func (l *TLinkLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return LinkLabel_GetTextBuf(l.instance, Buffer , BufSize)
}

// GetTextLen
func (l *TLinkLabel) GetTextLen() int32 {
    return LinkLabel_GetTextLen(l.instance)
}

// SetTextBuf
func (l *TLinkLabel) SetTextBuf(Buffer string) {
    LinkLabel_SetTextBuf(l.instance, Buffer)
}

// FindComponent
func (l *TLinkLabel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(LinkLabel_FindComponent(l.instance, AName))
}

// GetNamePath
func (l *TLinkLabel) GetNamePath() string {
    return LinkLabel_GetNamePath(l.instance)
}

// Assign
func (l *TLinkLabel) Assign(Source IObject) {
    LinkLabel_Assign(l.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TLinkLabel) DisposeOf() {
    LinkLabel_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TLinkLabel) ClassType() TClass {
    return LinkLabel_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TLinkLabel) ClassName() string {
    return LinkLabel_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TLinkLabel) InstanceSize() int32 {
    return LinkLabel_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TLinkLabel) InheritsFrom(AClass TClass) bool {
    return LinkLabel_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TLinkLabel) Equals(Obj IObject) bool {
    return LinkLabel_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TLinkLabel) GetHashCode() int32 {
    return LinkLabel_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TLinkLabel) ToString() string {
    return LinkLabel_ToString(l.instance)
}

// Align
func (l *TLinkLabel) Align() TAlign {
    return LinkLabel_GetAlign(l.instance)
}

// SetAlign
func (l *TLinkLabel) SetAlign(value TAlign) {
    LinkLabel_SetAlign(l.instance, value)
}

// Alignment
func (l *TLinkLabel) Alignment() TLinkAlignment {
    return LinkLabel_GetAlignment(l.instance)
}

// SetAlignment
func (l *TLinkLabel) SetAlignment(value TLinkAlignment) {
    LinkLabel_SetAlignment(l.instance, value)
}

// Anchors
func (l *TLinkLabel) Anchors() TAnchors {
    return LinkLabel_GetAnchors(l.instance)
}

// SetAnchors
func (l *TLinkLabel) SetAnchors(value TAnchors) {
    LinkLabel_SetAnchors(l.instance, value)
}

// AutoSize
func (l *TLinkLabel) AutoSize() bool {
    return LinkLabel_GetAutoSize(l.instance)
}

// SetAutoSize
func (l *TLinkLabel) SetAutoSize(value bool) {
    LinkLabel_SetAutoSize(l.instance, value)
}

// BevelEdges
func (l *TLinkLabel) BevelEdges() TBevelEdges {
    return LinkLabel_GetBevelEdges(l.instance)
}

// SetBevelEdges
func (l *TLinkLabel) SetBevelEdges(value TBevelEdges) {
    LinkLabel_SetBevelEdges(l.instance, value)
}

// BevelInner
func (l *TLinkLabel) BevelInner() TBevelCut {
    return LinkLabel_GetBevelInner(l.instance)
}

// SetBevelInner
func (l *TLinkLabel) SetBevelInner(value TBevelCut) {
    LinkLabel_SetBevelInner(l.instance, value)
}

// BevelKind
func (l *TLinkLabel) BevelKind() TBevelKind {
    return LinkLabel_GetBevelKind(l.instance)
}

// SetBevelKind
func (l *TLinkLabel) SetBevelKind(value TBevelKind) {
    LinkLabel_SetBevelKind(l.instance, value)
}

// BevelOuter
func (l *TLinkLabel) BevelOuter() TBevelCut {
    return LinkLabel_GetBevelOuter(l.instance)
}

// SetBevelOuter
func (l *TLinkLabel) SetBevelOuter(value TBevelCut) {
    LinkLabel_SetBevelOuter(l.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (l *TLinkLabel) Caption() string {
    return LinkLabel_GetCaption(l.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (l *TLinkLabel) SetCaption(value string) {
    LinkLabel_SetCaption(l.instance, value)
}

// Color
func (l *TLinkLabel) Color() TColor {
    return LinkLabel_GetColor(l.instance)
}

// SetColor
func (l *TLinkLabel) SetColor(value TColor) {
    LinkLabel_SetColor(l.instance, value)
}

// DragCursor
func (l *TLinkLabel) DragCursor() TCursor {
    return LinkLabel_GetDragCursor(l.instance)
}

// SetDragCursor
func (l *TLinkLabel) SetDragCursor(value TCursor) {
    LinkLabel_SetDragCursor(l.instance, value)
}

// DragKind
func (l *TLinkLabel) DragKind() TDragKind {
    return LinkLabel_GetDragKind(l.instance)
}

// SetDragKind
func (l *TLinkLabel) SetDragKind(value TDragKind) {
    LinkLabel_SetDragKind(l.instance, value)
}

// DragMode
func (l *TLinkLabel) DragMode() TDragMode {
    return LinkLabel_GetDragMode(l.instance)
}

// SetDragMode
func (l *TLinkLabel) SetDragMode(value TDragMode) {
    LinkLabel_SetDragMode(l.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (l *TLinkLabel) Enabled() bool {
    return LinkLabel_GetEnabled(l.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (l *TLinkLabel) SetEnabled(value bool) {
    LinkLabel_SetEnabled(l.instance, value)
}

// Font
func (l *TLinkLabel) Font() *TFont {
    return FontFromInst(LinkLabel_GetFont(l.instance))
}

// SetFont
func (l *TLinkLabel) SetFont(value *TFont) {
    LinkLabel_SetFont(l.instance, CheckPtr(value))
}

// ParentColor
func (l *TLinkLabel) ParentColor() bool {
    return LinkLabel_GetParentColor(l.instance)
}

// SetParentColor
func (l *TLinkLabel) SetParentColor(value bool) {
    LinkLabel_SetParentColor(l.instance, value)
}

// ParentFont
func (l *TLinkLabel) ParentFont() bool {
    return LinkLabel_GetParentFont(l.instance)
}

// SetParentFont
func (l *TLinkLabel) SetParentFont(value bool) {
    LinkLabel_SetParentFont(l.instance, value)
}

// ParentShowHint
func (l *TLinkLabel) ParentShowHint() bool {
    return LinkLabel_GetParentShowHint(l.instance)
}

// SetParentShowHint
func (l *TLinkLabel) SetParentShowHint(value bool) {
    LinkLabel_SetParentShowHint(l.instance, value)
}

// PopupMenu
func (l *TLinkLabel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(LinkLabel_GetPopupMenu(l.instance))
}

// SetPopupMenu
func (l *TLinkLabel) SetPopupMenu(value IComponent) {
    LinkLabel_SetPopupMenu(l.instance, CheckPtr(value))
}

// ShowHint
func (l *TLinkLabel) ShowHint() bool {
    return LinkLabel_GetShowHint(l.instance)
}

// SetShowHint
func (l *TLinkLabel) SetShowHint(value bool) {
    LinkLabel_SetShowHint(l.instance, value)
}

// TabOrder
func (l *TLinkLabel) TabOrder() TTabOrder {
    return LinkLabel_GetTabOrder(l.instance)
}

// SetTabOrder
func (l *TLinkLabel) SetTabOrder(value TTabOrder) {
    LinkLabel_SetTabOrder(l.instance, value)
}

// TabStop
func (l *TLinkLabel) TabStop() bool {
    return LinkLabel_GetTabStop(l.instance)
}

// SetTabStop
func (l *TLinkLabel) SetTabStop(value bool) {
    LinkLabel_SetTabStop(l.instance, value)
}

// UseVisualStyle
func (l *TLinkLabel) UseVisualStyle() bool {
    return LinkLabel_GetUseVisualStyle(l.instance)
}

// SetUseVisualStyle
func (l *TLinkLabel) SetUseVisualStyle(value bool) {
    LinkLabel_SetUseVisualStyle(l.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (l *TLinkLabel) Visible() bool {
    return LinkLabel_GetVisible(l.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (l *TLinkLabel) SetVisible(value bool) {
    LinkLabel_SetVisible(l.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (l *TLinkLabel) SetOnClick(fn TNotifyEvent) {
    LinkLabel_SetOnClick(l.instance, fn)
}

// SetOnContextPopup
func (l *TLinkLabel) SetOnContextPopup(fn TContextPopupEvent) {
    LinkLabel_SetOnContextPopup(l.instance, fn)
}

// SetOnDblClick
func (l *TLinkLabel) SetOnDblClick(fn TNotifyEvent) {
    LinkLabel_SetOnDblClick(l.instance, fn)
}

// SetOnDragDrop
func (l *TLinkLabel) SetOnDragDrop(fn TDragDropEvent) {
    LinkLabel_SetOnDragDrop(l.instance, fn)
}

// SetOnDragOver
func (l *TLinkLabel) SetOnDragOver(fn TDragOverEvent) {
    LinkLabel_SetOnDragOver(l.instance, fn)
}

// SetOnEndDock
func (l *TLinkLabel) SetOnEndDock(fn TEndDragEvent) {
    LinkLabel_SetOnEndDock(l.instance, fn)
}

// SetOnEndDrag
func (l *TLinkLabel) SetOnEndDrag(fn TEndDragEvent) {
    LinkLabel_SetOnEndDrag(l.instance, fn)
}

// SetOnMouseDown
func (l *TLinkLabel) SetOnMouseDown(fn TMouseEvent) {
    LinkLabel_SetOnMouseDown(l.instance, fn)
}

// SetOnMouseEnter
func (l *TLinkLabel) SetOnMouseEnter(fn TNotifyEvent) {
    LinkLabel_SetOnMouseEnter(l.instance, fn)
}

// SetOnMouseLeave
func (l *TLinkLabel) SetOnMouseLeave(fn TNotifyEvent) {
    LinkLabel_SetOnMouseLeave(l.instance, fn)
}

// SetOnMouseMove
func (l *TLinkLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    LinkLabel_SetOnMouseMove(l.instance, fn)
}

// SetOnMouseUp
func (l *TLinkLabel) SetOnMouseUp(fn TMouseEvent) {
    LinkLabel_SetOnMouseUp(l.instance, fn)
}

// SetOnStartDock
func (l *TLinkLabel) SetOnStartDock(fn TStartDockEvent) {
    LinkLabel_SetOnStartDock(l.instance, fn)
}

// SetOnLinkClick
func (l *TLinkLabel) SetOnLinkClick(fn TSysLinkEvent) {
    LinkLabel_SetOnLinkClick(l.instance, fn)
}

// DockClientCount
func (l *TLinkLabel) DockClientCount() int32 {
    return LinkLabel_GetDockClientCount(l.instance)
}

// DockSite
func (l *TLinkLabel) DockSite() bool {
    return LinkLabel_GetDockSite(l.instance)
}

// SetDockSite
func (l *TLinkLabel) SetDockSite(value bool) {
    LinkLabel_SetDockSite(l.instance, value)
}

// DoubleBuffered
func (l *TLinkLabel) DoubleBuffered() bool {
    return LinkLabel_GetDoubleBuffered(l.instance)
}

// SetDoubleBuffered
func (l *TLinkLabel) SetDoubleBuffered(value bool) {
    LinkLabel_SetDoubleBuffered(l.instance, value)
}

// AlignDisabled
func (l *TLinkLabel) AlignDisabled() bool {
    return LinkLabel_GetAlignDisabled(l.instance)
}

// MouseInClient
func (l *TLinkLabel) MouseInClient() bool {
    return LinkLabel_GetMouseInClient(l.instance)
}

// VisibleDockClientCount
func (l *TLinkLabel) VisibleDockClientCount() int32 {
    return LinkLabel_GetVisibleDockClientCount(l.instance)
}

// Brush
func (l *TLinkLabel) Brush() *TBrush {
    return BrushFromInst(LinkLabel_GetBrush(l.instance))
}

// ControlCount
func (l *TLinkLabel) ControlCount() int32 {
    return LinkLabel_GetControlCount(l.instance)
}

// Handle
func (l *TLinkLabel) Handle() HWND {
    return LinkLabel_GetHandle(l.instance)
}

// ParentDoubleBuffered
func (l *TLinkLabel) ParentDoubleBuffered() bool {
    return LinkLabel_GetParentDoubleBuffered(l.instance)
}

// SetParentDoubleBuffered
func (l *TLinkLabel) SetParentDoubleBuffered(value bool) {
    LinkLabel_SetParentDoubleBuffered(l.instance, value)
}

// ParentWindow
func (l *TLinkLabel) ParentWindow() HWND {
    return LinkLabel_GetParentWindow(l.instance)
}

// SetParentWindow
func (l *TLinkLabel) SetParentWindow(value HWND) {
    LinkLabel_SetParentWindow(l.instance, value)
}

// UseDockManager
func (l *TLinkLabel) UseDockManager() bool {
    return LinkLabel_GetUseDockManager(l.instance)
}

// SetUseDockManager
func (l *TLinkLabel) SetUseDockManager(value bool) {
    LinkLabel_SetUseDockManager(l.instance, value)
}

// Action
func (l *TLinkLabel) Action() *TAction {
    return ActionFromInst(LinkLabel_GetAction(l.instance))
}

// SetAction
func (l *TLinkLabel) SetAction(value IComponent) {
    LinkLabel_SetAction(l.instance, CheckPtr(value))
}

// BiDiMode
func (l *TLinkLabel) BiDiMode() TBiDiMode {
    return LinkLabel_GetBiDiMode(l.instance)
}

// SetBiDiMode
func (l *TLinkLabel) SetBiDiMode(value TBiDiMode) {
    LinkLabel_SetBiDiMode(l.instance, value)
}

// BoundsRect
func (l *TLinkLabel) BoundsRect() TRect {
    return LinkLabel_GetBoundsRect(l.instance)
}

// SetBoundsRect
func (l *TLinkLabel) SetBoundsRect(value TRect) {
    LinkLabel_SetBoundsRect(l.instance, value)
}

// ClientHeight
func (l *TLinkLabel) ClientHeight() int32 {
    return LinkLabel_GetClientHeight(l.instance)
}

// SetClientHeight
func (l *TLinkLabel) SetClientHeight(value int32) {
    LinkLabel_SetClientHeight(l.instance, value)
}

// ClientOrigin
func (l *TLinkLabel) ClientOrigin() TPoint {
    return LinkLabel_GetClientOrigin(l.instance)
}

// ClientRect
func (l *TLinkLabel) ClientRect() TRect {
    return LinkLabel_GetClientRect(l.instance)
}

// ClientWidth
func (l *TLinkLabel) ClientWidth() int32 {
    return LinkLabel_GetClientWidth(l.instance)
}

// SetClientWidth
func (l *TLinkLabel) SetClientWidth(value int32) {
    LinkLabel_SetClientWidth(l.instance, value)
}

// ControlState
func (l *TLinkLabel) ControlState() TControlState {
    return LinkLabel_GetControlState(l.instance)
}

// SetControlState
func (l *TLinkLabel) SetControlState(value TControlState) {
    LinkLabel_SetControlState(l.instance, value)
}

// ControlStyle
func (l *TLinkLabel) ControlStyle() TControlStyle {
    return LinkLabel_GetControlStyle(l.instance)
}

// SetControlStyle
func (l *TLinkLabel) SetControlStyle(value TControlStyle) {
    LinkLabel_SetControlStyle(l.instance, value)
}

// ExplicitLeft
func (l *TLinkLabel) ExplicitLeft() int32 {
    return LinkLabel_GetExplicitLeft(l.instance)
}

// ExplicitTop
func (l *TLinkLabel) ExplicitTop() int32 {
    return LinkLabel_GetExplicitTop(l.instance)
}

// ExplicitWidth
func (l *TLinkLabel) ExplicitWidth() int32 {
    return LinkLabel_GetExplicitWidth(l.instance)
}

// ExplicitHeight
func (l *TLinkLabel) ExplicitHeight() int32 {
    return LinkLabel_GetExplicitHeight(l.instance)
}

// Floating
func (l *TLinkLabel) Floating() bool {
    return LinkLabel_GetFloating(l.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (l *TLinkLabel) Parent() *TWinControl {
    return WinControlFromInst(LinkLabel_GetParent(l.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (l *TLinkLabel) SetParent(value IWinControl) {
    LinkLabel_SetParent(l.instance, CheckPtr(value))
}

// StyleElements
func (l *TLinkLabel) StyleElements() TStyleElements {
    return LinkLabel_GetStyleElements(l.instance)
}

// SetStyleElements
func (l *TLinkLabel) SetStyleElements(value TStyleElements) {
    LinkLabel_SetStyleElements(l.instance, value)
}

// AlignWithMargins
func (l *TLinkLabel) AlignWithMargins() bool {
    return LinkLabel_GetAlignWithMargins(l.instance)
}

// SetAlignWithMargins
func (l *TLinkLabel) SetAlignWithMargins(value bool) {
    LinkLabel_SetAlignWithMargins(l.instance, value)
}

// Left
func (l *TLinkLabel) Left() int32 {
    return LinkLabel_GetLeft(l.instance)
}

// SetLeft
func (l *TLinkLabel) SetLeft(value int32) {
    LinkLabel_SetLeft(l.instance, value)
}

// Top
func (l *TLinkLabel) Top() int32 {
    return LinkLabel_GetTop(l.instance)
}

// SetTop
func (l *TLinkLabel) SetTop(value int32) {
    LinkLabel_SetTop(l.instance, value)
}

// Width
func (l *TLinkLabel) Width() int32 {
    return LinkLabel_GetWidth(l.instance)
}

// SetWidth
func (l *TLinkLabel) SetWidth(value int32) {
    LinkLabel_SetWidth(l.instance, value)
}

// Height
func (l *TLinkLabel) Height() int32 {
    return LinkLabel_GetHeight(l.instance)
}

// SetHeight
func (l *TLinkLabel) SetHeight(value int32) {
    LinkLabel_SetHeight(l.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (l *TLinkLabel) Cursor() TCursor {
    return LinkLabel_GetCursor(l.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (l *TLinkLabel) SetCursor(value TCursor) {
    LinkLabel_SetCursor(l.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (l *TLinkLabel) Hint() string {
    return LinkLabel_GetHint(l.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (l *TLinkLabel) SetHint(value string) {
    LinkLabel_SetHint(l.instance, value)
}

// Margins
func (l *TLinkLabel) Margins() *TMargins {
    return MarginsFromInst(LinkLabel_GetMargins(l.instance))
}

// SetMargins
func (l *TLinkLabel) SetMargins(value *TMargins) {
    LinkLabel_SetMargins(l.instance, CheckPtr(value))
}

// CustomHint
func (l *TLinkLabel) CustomHint() *TCustomHint {
    return CustomHintFromInst(LinkLabel_GetCustomHint(l.instance))
}

// SetCustomHint
func (l *TLinkLabel) SetCustomHint(value IComponent) {
    LinkLabel_SetCustomHint(l.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (l *TLinkLabel) ComponentCount() int32 {
    return LinkLabel_GetComponentCount(l.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (l *TLinkLabel) ComponentIndex() int32 {
    return LinkLabel_GetComponentIndex(l.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (l *TLinkLabel) SetComponentIndex(value int32) {
    LinkLabel_SetComponentIndex(l.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TLinkLabel) Owner() *TComponent {
    return ComponentFromInst(LinkLabel_GetOwner(l.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (l *TLinkLabel) Name() string {
    return LinkLabel_GetName(l.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (l *TLinkLabel) SetName(value string) {
    LinkLabel_SetName(l.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TLinkLabel) Tag() int {
    return LinkLabel_GetTag(l.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TLinkLabel) SetTag(value int) {
    LinkLabel_SetTag(l.instance, value)
}

// DockClients
func (l *TLinkLabel) DockClients(Index int32) *TControl {
    return ControlFromInst(LinkLabel_GetDockClients(l.instance, Index))
}

// Controls
func (l *TLinkLabel) Controls(Index int32) *TControl {
    return ControlFromInst(LinkLabel_GetControls(l.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TLinkLabel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(LinkLabel_GetComponents(l.instance, AIndex))
}

