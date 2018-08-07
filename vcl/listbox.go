
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

type TListBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewListBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListBox(owner IComponent) *TListBox {
    l := new(TListBox)
    l.instance = ListBox_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ListBoxFromInst(inst uintptr) *TListBox {
    l := new(TListBox)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// ListBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ListBoxFromObj(obj IObject) *TListBox {
    l := new(TListBox)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ListBoxFromUnsafePointer(ptr unsafe.Pointer) *TListBox {
    l := new(TListBox)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TListBox) Free() {
    if l.instance != 0 {
        ListBox_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListBox) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListBox) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListBox) IsValid() bool {
    return l.instance != 0
}

// TListBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListBoxClass() TClass {
    return ListBox_StaticClassType()
}

// AddItem
func (l *TListBox) AddItem(Item string, AObject IObject) {
    ListBox_AddItem(l.instance, Item , CheckPtr(AObject))
}

// Clear
func (l *TListBox) Clear() {
    ListBox_Clear(l.instance)
}

// ClearSelection
func (l *TListBox) ClearSelection() {
    ListBox_ClearSelection(l.instance)
}

// DeleteSelected
func (l *TListBox) DeleteSelected() {
    ListBox_DeleteSelected(l.instance)
}

// SelectAll
func (l *TListBox) SelectAll() {
    ListBox_SelectAll(l.instance)
}

// CanFocus
func (l *TListBox) CanFocus() bool {
    return ListBox_CanFocus(l.instance)
}

// FlipChildren
func (l *TListBox) FlipChildren(AllLevels bool) {
    ListBox_FlipChildren(l.instance, AllLevels)
}

// Focused
func (l *TListBox) Focused() bool {
    return ListBox_Focused(l.instance)
}

// HandleAllocated
func (l *TListBox) HandleAllocated() bool {
    return ListBox_HandleAllocated(l.instance)
}

// Invalidate
func (l *TListBox) Invalidate() {
    ListBox_Invalidate(l.instance)
}

// Realign
func (l *TListBox) Realign() {
    ListBox_Realign(l.instance)
}

// Repaint
func (l *TListBox) Repaint() {
    ListBox_Repaint(l.instance)
}

// ScaleBy
func (l *TListBox) ScaleBy(M int32, D int32) {
    ListBox_ScaleBy(l.instance, M , D)
}

// SetBounds
func (l *TListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ListBox_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (l *TListBox) SetFocus() {
    ListBox_SetFocus(l.instance)
}

// Update
func (l *TListBox) Update() {
    ListBox_Update(l.instance)
}

// BringToFront
func (l *TListBox) BringToFront() {
    ListBox_BringToFront(l.instance)
}

// ClientToScreen
func (l *TListBox) ClientToScreen(Point TPoint) TPoint {
    return ListBox_ClientToScreen(l.instance, Point)
}

// ClientToParent
func (l *TListBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ListBox_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// Dragging
func (l *TListBox) Dragging() bool {
    return ListBox_Dragging(l.instance)
}

// HasParent
func (l *TListBox) HasParent() bool {
    return ListBox_HasParent(l.instance)
}

// Hide
func (l *TListBox) Hide() {
    ListBox_Hide(l.instance)
}

// Perform
func (l *TListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ListBox_Perform(l.instance, Msg , WParam , LParam)
}

// Refresh
func (l *TListBox) Refresh() {
    ListBox_Refresh(l.instance)
}

// ScreenToClient
func (l *TListBox) ScreenToClient(Point TPoint) TPoint {
    return ListBox_ScreenToClient(l.instance, Point)
}

// ParentToClient
func (l *TListBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ListBox_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (l *TListBox) SendToBack() {
    ListBox_SendToBack(l.instance)
}

// Show
func (l *TListBox) Show() {
    ListBox_Show(l.instance)
}

// GetTextBuf
func (l *TListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ListBox_GetTextBuf(l.instance, Buffer , BufSize)
}

// GetTextLen
func (l *TListBox) GetTextLen() int32 {
    return ListBox_GetTextLen(l.instance)
}

// FindComponent
func (l *TListBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ListBox_FindComponent(l.instance, AName))
}

// GetNamePath
func (l *TListBox) GetNamePath() string {
    return ListBox_GetNamePath(l.instance)
}

// Assign
func (l *TListBox) Assign(Source IObject) {
    ListBox_Assign(l.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TListBox) DisposeOf() {
    ListBox_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListBox) ClassType() TClass {
    return ListBox_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListBox) ClassName() string {
    return ListBox_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListBox) InstanceSize() int32 {
    return ListBox_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListBox) InheritsFrom(AClass TClass) bool {
    return ListBox_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListBox) Equals(Obj IObject) bool {
    return ListBox_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListBox) GetHashCode() int32 {
    return ListBox_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TListBox) ToString() string {
    return ListBox_ToString(l.instance)
}

// Style
func (l *TListBox) Style() TListBoxStyle {
    return ListBox_GetStyle(l.instance)
}

// SetStyle
func (l *TListBox) SetStyle(value TListBoxStyle) {
    ListBox_SetStyle(l.instance, value)
}

// AutoComplete
func (l *TListBox) AutoComplete() bool {
    return ListBox_GetAutoComplete(l.instance)
}

// SetAutoComplete
func (l *TListBox) SetAutoComplete(value bool) {
    ListBox_SetAutoComplete(l.instance, value)
}

// AutoCompleteDelay
func (l *TListBox) AutoCompleteDelay() uint32 {
    return ListBox_GetAutoCompleteDelay(l.instance)
}

// SetAutoCompleteDelay
func (l *TListBox) SetAutoCompleteDelay(value uint32) {
    ListBox_SetAutoCompleteDelay(l.instance, value)
}

// Align
func (l *TListBox) Align() TAlign {
    return ListBox_GetAlign(l.instance)
}

// SetAlign
func (l *TListBox) SetAlign(value TAlign) {
    ListBox_SetAlign(l.instance, value)
}

// Anchors
func (l *TListBox) Anchors() TAnchors {
    return ListBox_GetAnchors(l.instance)
}

// SetAnchors
func (l *TListBox) SetAnchors(value TAnchors) {
    ListBox_SetAnchors(l.instance, value)
}

// BevelEdges
func (l *TListBox) BevelEdges() TBevelEdges {
    return ListBox_GetBevelEdges(l.instance)
}

// SetBevelEdges
func (l *TListBox) SetBevelEdges(value TBevelEdges) {
    ListBox_SetBevelEdges(l.instance, value)
}

// BevelInner
func (l *TListBox) BevelInner() TBevelCut {
    return ListBox_GetBevelInner(l.instance)
}

// SetBevelInner
func (l *TListBox) SetBevelInner(value TBevelCut) {
    ListBox_SetBevelInner(l.instance, value)
}

// BevelKind
func (l *TListBox) BevelKind() TBevelKind {
    return ListBox_GetBevelKind(l.instance)
}

// SetBevelKind
func (l *TListBox) SetBevelKind(value TBevelKind) {
    ListBox_SetBevelKind(l.instance, value)
}

// BevelOuter
func (l *TListBox) BevelOuter() TBevelCut {
    return ListBox_GetBevelOuter(l.instance)
}

// SetBevelOuter
func (l *TListBox) SetBevelOuter(value TBevelCut) {
    ListBox_SetBevelOuter(l.instance, value)
}

// BiDiMode
func (l *TListBox) BiDiMode() TBiDiMode {
    return ListBox_GetBiDiMode(l.instance)
}

// SetBiDiMode
func (l *TListBox) SetBiDiMode(value TBiDiMode) {
    ListBox_SetBiDiMode(l.instance, value)
}

// BorderStyle
func (l *TListBox) BorderStyle() TBorderStyle {
    return ListBox_GetBorderStyle(l.instance)
}

// SetBorderStyle
func (l *TListBox) SetBorderStyle(value TBorderStyle) {
    ListBox_SetBorderStyle(l.instance, value)
}

// Color
func (l *TListBox) Color() TColor {
    return ListBox_GetColor(l.instance)
}

// SetColor
func (l *TListBox) SetColor(value TColor) {
    ListBox_SetColor(l.instance, value)
}

// Columns
func (l *TListBox) Columns() int32 {
    return ListBox_GetColumns(l.instance)
}

// SetColumns
func (l *TListBox) SetColumns(value int32) {
    ListBox_SetColumns(l.instance, value)
}

// DoubleBuffered
func (l *TListBox) DoubleBuffered() bool {
    return ListBox_GetDoubleBuffered(l.instance)
}

// SetDoubleBuffered
func (l *TListBox) SetDoubleBuffered(value bool) {
    ListBox_SetDoubleBuffered(l.instance, value)
}

// DragCursor
func (l *TListBox) DragCursor() TCursor {
    return ListBox_GetDragCursor(l.instance)
}

// SetDragCursor
func (l *TListBox) SetDragCursor(value TCursor) {
    ListBox_SetDragCursor(l.instance, value)
}

// DragKind
func (l *TListBox) DragKind() TDragKind {
    return ListBox_GetDragKind(l.instance)
}

// SetDragKind
func (l *TListBox) SetDragKind(value TDragKind) {
    ListBox_SetDragKind(l.instance, value)
}

// DragMode
func (l *TListBox) DragMode() TDragMode {
    return ListBox_GetDragMode(l.instance)
}

// SetDragMode
func (l *TListBox) SetDragMode(value TDragMode) {
    ListBox_SetDragMode(l.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (l *TListBox) Enabled() bool {
    return ListBox_GetEnabled(l.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (l *TListBox) SetEnabled(value bool) {
    ListBox_SetEnabled(l.instance, value)
}

// Font
func (l *TListBox) Font() *TFont {
    return FontFromInst(ListBox_GetFont(l.instance))
}

// SetFont
func (l *TListBox) SetFont(value *TFont) {
    ListBox_SetFont(l.instance, CheckPtr(value))
}

// ItemHeight
func (l *TListBox) ItemHeight() int32 {
    return ListBox_GetItemHeight(l.instance)
}

// SetItemHeight
func (l *TListBox) SetItemHeight(value int32) {
    ListBox_SetItemHeight(l.instance, value)
}

// Items
func (l *TListBox) Items() *TStrings {
    return StringsFromInst(ListBox_GetItems(l.instance))
}

// SetItems
func (l *TListBox) SetItems(value IObject) {
    ListBox_SetItems(l.instance, CheckPtr(value))
}

// MultiSelect
func (l *TListBox) MultiSelect() bool {
    return ListBox_GetMultiSelect(l.instance)
}

// SetMultiSelect
func (l *TListBox) SetMultiSelect(value bool) {
    ListBox_SetMultiSelect(l.instance, value)
}

// ParentColor
func (l *TListBox) ParentColor() bool {
    return ListBox_GetParentColor(l.instance)
}

// SetParentColor
func (l *TListBox) SetParentColor(value bool) {
    ListBox_SetParentColor(l.instance, value)
}

// ParentCtl3D
func (l *TListBox) ParentCtl3D() bool {
    return ListBox_GetParentCtl3D(l.instance)
}

// SetParentCtl3D
func (l *TListBox) SetParentCtl3D(value bool) {
    ListBox_SetParentCtl3D(l.instance, value)
}

// ParentDoubleBuffered
func (l *TListBox) ParentDoubleBuffered() bool {
    return ListBox_GetParentDoubleBuffered(l.instance)
}

// SetParentDoubleBuffered
func (l *TListBox) SetParentDoubleBuffered(value bool) {
    ListBox_SetParentDoubleBuffered(l.instance, value)
}

// ParentFont
func (l *TListBox) ParentFont() bool {
    return ListBox_GetParentFont(l.instance)
}

// SetParentFont
func (l *TListBox) SetParentFont(value bool) {
    ListBox_SetParentFont(l.instance, value)
}

// ParentShowHint
func (l *TListBox) ParentShowHint() bool {
    return ListBox_GetParentShowHint(l.instance)
}

// SetParentShowHint
func (l *TListBox) SetParentShowHint(value bool) {
    ListBox_SetParentShowHint(l.instance, value)
}

// PopupMenu
func (l *TListBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ListBox_GetPopupMenu(l.instance))
}

// SetPopupMenu
func (l *TListBox) SetPopupMenu(value IComponent) {
    ListBox_SetPopupMenu(l.instance, CheckPtr(value))
}

// ShowHint
func (l *TListBox) ShowHint() bool {
    return ListBox_GetShowHint(l.instance)
}

// SetShowHint
func (l *TListBox) SetShowHint(value bool) {
    ListBox_SetShowHint(l.instance, value)
}

// Sorted
func (l *TListBox) Sorted() bool {
    return ListBox_GetSorted(l.instance)
}

// SetSorted
func (l *TListBox) SetSorted(value bool) {
    ListBox_SetSorted(l.instance, value)
}

// TabOrder
func (l *TListBox) TabOrder() uint16 {
    return ListBox_GetTabOrder(l.instance)
}

// SetTabOrder
func (l *TListBox) SetTabOrder(value uint16) {
    ListBox_SetTabOrder(l.instance, value)
}

// TabStop
func (l *TListBox) TabStop() bool {
    return ListBox_GetTabStop(l.instance)
}

// SetTabStop
func (l *TListBox) SetTabStop(value bool) {
    ListBox_SetTabStop(l.instance, value)
}

// TabWidth
func (l *TListBox) TabWidth() int32 {
    return ListBox_GetTabWidth(l.instance)
}

// SetTabWidth
func (l *TListBox) SetTabWidth(value int32) {
    ListBox_SetTabWidth(l.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (l *TListBox) Visible() bool {
    return ListBox_GetVisible(l.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (l *TListBox) SetVisible(value bool) {
    ListBox_SetVisible(l.instance, value)
}

// StyleElements
func (l *TListBox) StyleElements() TStyleElements {
    return ListBox_GetStyleElements(l.instance)
}

// SetStyleElements
func (l *TListBox) SetStyleElements(value TStyleElements) {
    ListBox_SetStyleElements(l.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (l *TListBox) SetOnClick(fn TNotifyEvent) {
    ListBox_SetOnClick(l.instance, fn)
}

// SetOnContextPopup
func (l *TListBox) SetOnContextPopup(fn TContextPopupEvent) {
    ListBox_SetOnContextPopup(l.instance, fn)
}

// SetOnDblClick
func (l *TListBox) SetOnDblClick(fn TNotifyEvent) {
    ListBox_SetOnDblClick(l.instance, fn)
}

// SetOnDragDrop
func (l *TListBox) SetOnDragDrop(fn TDragDropEvent) {
    ListBox_SetOnDragDrop(l.instance, fn)
}

// SetOnDragOver
func (l *TListBox) SetOnDragOver(fn TDragOverEvent) {
    ListBox_SetOnDragOver(l.instance, fn)
}

// SetOnDrawItem
func (l *TListBox) SetOnDrawItem(fn TDrawItemEvent) {
    ListBox_SetOnDrawItem(l.instance, fn)
}

// SetOnEndDock
func (l *TListBox) SetOnEndDock(fn TEndDragEvent) {
    ListBox_SetOnEndDock(l.instance, fn)
}

// SetOnEndDrag
func (l *TListBox) SetOnEndDrag(fn TEndDragEvent) {
    ListBox_SetOnEndDrag(l.instance, fn)
}

// SetOnEnter
func (l *TListBox) SetOnEnter(fn TNotifyEvent) {
    ListBox_SetOnEnter(l.instance, fn)
}

// SetOnExit
func (l *TListBox) SetOnExit(fn TNotifyEvent) {
    ListBox_SetOnExit(l.instance, fn)
}

// SetOnKeyDown
func (l *TListBox) SetOnKeyDown(fn TKeyEvent) {
    ListBox_SetOnKeyDown(l.instance, fn)
}

// SetOnKeyPress
func (l *TListBox) SetOnKeyPress(fn TKeyPressEvent) {
    ListBox_SetOnKeyPress(l.instance, fn)
}

// SetOnKeyUp
func (l *TListBox) SetOnKeyUp(fn TKeyEvent) {
    ListBox_SetOnKeyUp(l.instance, fn)
}

// SetOnMouseDown
func (l *TListBox) SetOnMouseDown(fn TMouseEvent) {
    ListBox_SetOnMouseDown(l.instance, fn)
}

// SetOnMouseEnter
func (l *TListBox) SetOnMouseEnter(fn TNotifyEvent) {
    ListBox_SetOnMouseEnter(l.instance, fn)
}

// SetOnMouseLeave
func (l *TListBox) SetOnMouseLeave(fn TNotifyEvent) {
    ListBox_SetOnMouseLeave(l.instance, fn)
}

// SetOnMouseMove
func (l *TListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ListBox_SetOnMouseMove(l.instance, fn)
}

// SetOnMouseUp
func (l *TListBox) SetOnMouseUp(fn TMouseEvent) {
    ListBox_SetOnMouseUp(l.instance, fn)
}

// SetOnStartDock
func (l *TListBox) SetOnStartDock(fn TStartDockEvent) {
    ListBox_SetOnStartDock(l.instance, fn)
}

// Canvas
func (l *TListBox) Canvas() *TCanvas {
    return CanvasFromInst(ListBox_GetCanvas(l.instance))
}

// Count
func (l *TListBox) Count() int32 {
    return ListBox_GetCount(l.instance)
}

// SetCount
func (l *TListBox) SetCount(value int32) {
    ListBox_SetCount(l.instance, value)
}

// SelCount
func (l *TListBox) SelCount() int32 {
    return ListBox_GetSelCount(l.instance)
}

// ItemIndex
func (l *TListBox) ItemIndex() int32 {
    return ListBox_GetItemIndex(l.instance)
}

// SetItemIndex
func (l *TListBox) SetItemIndex(value int32) {
    ListBox_SetItemIndex(l.instance, value)
}

// DockSite
func (l *TListBox) DockSite() bool {
    return ListBox_GetDockSite(l.instance)
}

// SetDockSite
func (l *TListBox) SetDockSite(value bool) {
    ListBox_SetDockSite(l.instance, value)
}

// Brush
func (l *TListBox) Brush() *TBrush {
    return BrushFromInst(ListBox_GetBrush(l.instance))
}

// ControlCount
func (l *TListBox) ControlCount() int32 {
    return ListBox_GetControlCount(l.instance)
}

// Handle
func (l *TListBox) Handle() HWND {
    return ListBox_GetHandle(l.instance)
}

// ParentWindow
func (l *TListBox) ParentWindow() HWND {
    return ListBox_GetParentWindow(l.instance)
}

// SetParentWindow
func (l *TListBox) SetParentWindow(value HWND) {
    ListBox_SetParentWindow(l.instance, value)
}

// UseDockManager
func (l *TListBox) UseDockManager() bool {
    return ListBox_GetUseDockManager(l.instance)
}

// SetUseDockManager
func (l *TListBox) SetUseDockManager(value bool) {
    ListBox_SetUseDockManager(l.instance, value)
}

// Action
func (l *TListBox) Action() *TAction {
    return ActionFromInst(ListBox_GetAction(l.instance))
}

// SetAction
func (l *TListBox) SetAction(value IComponent) {
    ListBox_SetAction(l.instance, CheckPtr(value))
}

// BoundsRect
func (l *TListBox) BoundsRect() TRect {
    return ListBox_GetBoundsRect(l.instance)
}

// SetBoundsRect
func (l *TListBox) SetBoundsRect(value TRect) {
    ListBox_SetBoundsRect(l.instance, value)
}

// ClientHeight
func (l *TListBox) ClientHeight() int32 {
    return ListBox_GetClientHeight(l.instance)
}

// SetClientHeight
func (l *TListBox) SetClientHeight(value int32) {
    ListBox_SetClientHeight(l.instance, value)
}

// ClientRect
func (l *TListBox) ClientRect() TRect {
    return ListBox_GetClientRect(l.instance)
}

// ClientWidth
func (l *TListBox) ClientWidth() int32 {
    return ListBox_GetClientWidth(l.instance)
}

// SetClientWidth
func (l *TListBox) SetClientWidth(value int32) {
    ListBox_SetClientWidth(l.instance, value)
}

// ExplicitLeft
func (l *TListBox) ExplicitLeft() int32 {
    return ListBox_GetExplicitLeft(l.instance)
}

// ExplicitTop
func (l *TListBox) ExplicitTop() int32 {
    return ListBox_GetExplicitTop(l.instance)
}

// ExplicitWidth
func (l *TListBox) ExplicitWidth() int32 {
    return ListBox_GetExplicitWidth(l.instance)
}

// ExplicitHeight
func (l *TListBox) ExplicitHeight() int32 {
    return ListBox_GetExplicitHeight(l.instance)
}

// Floating
func (l *TListBox) Floating() bool {
    return ListBox_GetFloating(l.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (l *TListBox) Parent() *TWinControl {
    return WinControlFromInst(ListBox_GetParent(l.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (l *TListBox) SetParent(value IWinControl) {
    ListBox_SetParent(l.instance, CheckPtr(value))
}

// AlignWithMargins
func (l *TListBox) AlignWithMargins() bool {
    return ListBox_GetAlignWithMargins(l.instance)
}

// SetAlignWithMargins
func (l *TListBox) SetAlignWithMargins(value bool) {
    ListBox_SetAlignWithMargins(l.instance, value)
}

// Left
func (l *TListBox) Left() int32 {
    return ListBox_GetLeft(l.instance)
}

// SetLeft
func (l *TListBox) SetLeft(value int32) {
    ListBox_SetLeft(l.instance, value)
}

// Top
func (l *TListBox) Top() int32 {
    return ListBox_GetTop(l.instance)
}

// SetTop
func (l *TListBox) SetTop(value int32) {
    ListBox_SetTop(l.instance, value)
}

// Width
func (l *TListBox) Width() int32 {
    return ListBox_GetWidth(l.instance)
}

// SetWidth
func (l *TListBox) SetWidth(value int32) {
    ListBox_SetWidth(l.instance, value)
}

// Height
func (l *TListBox) Height() int32 {
    return ListBox_GetHeight(l.instance)
}

// SetHeight
func (l *TListBox) SetHeight(value int32) {
    ListBox_SetHeight(l.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (l *TListBox) Cursor() TCursor {
    return ListBox_GetCursor(l.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (l *TListBox) SetCursor(value TCursor) {
    ListBox_SetCursor(l.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (l *TListBox) Hint() string {
    return ListBox_GetHint(l.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (l *TListBox) SetHint(value string) {
    ListBox_SetHint(l.instance, value)
}

// Margins
func (l *TListBox) Margins() *TMargins {
    return MarginsFromInst(ListBox_GetMargins(l.instance))
}

// SetMargins
func (l *TListBox) SetMargins(value *TMargins) {
    ListBox_SetMargins(l.instance, CheckPtr(value))
}

// CustomHint
func (l *TListBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ListBox_GetCustomHint(l.instance))
}

// SetCustomHint
func (l *TListBox) SetCustomHint(value IComponent) {
    ListBox_SetCustomHint(l.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (l *TListBox) ComponentCount() int32 {
    return ListBox_GetComponentCount(l.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (l *TListBox) ComponentIndex() int32 {
    return ListBox_GetComponentIndex(l.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (l *TListBox) SetComponentIndex(value int32) {
    ListBox_SetComponentIndex(l.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TListBox) Owner() *TComponent {
    return ComponentFromInst(ListBox_GetOwner(l.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (l *TListBox) Name() string {
    return ListBox_GetName(l.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (l *TListBox) SetName(value string) {
    ListBox_SetName(l.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TListBox) Tag() int {
    return ListBox_GetTag(l.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TListBox) SetTag(value int) {
    ListBox_SetTag(l.instance, value)
}

// Selected
func (l *TListBox) Selected(Index int32) bool {
    return ListBox_GetSelected(l.instance, Index)
}

// Selected
func (l *TListBox) SetSelected(Index int32, value bool) {
    ListBox_SetSelected(l.instance, Index, value)
}

// Controls
func (l *TListBox) Controls(Index int32) *TControl {
    return ControlFromInst(ListBox_GetControls(l.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TListBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ListBox_GetComponents(l.instance, AIndex))
}

