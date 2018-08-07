
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

type TColorListBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewColorListBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewColorListBox(owner IComponent) *TColorListBox {
    c := new(TColorListBox)
    c.instance = ColorListBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ColorListBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ColorListBoxFromInst(inst uintptr) *TColorListBox {
    c := new(TColorListBox)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ColorListBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ColorListBoxFromObj(obj IObject) *TColorListBox {
    c := new(TColorListBox)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ColorListBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ColorListBoxFromUnsafePointer(ptr unsafe.Pointer) *TColorListBox {
    c := new(TColorListBox)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TColorListBox) Free() {
    if c.instance != 0 {
        ColorListBox_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TColorListBox) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TColorListBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TColorListBox) IsValid() bool {
    return c.instance != 0
}

// TColorListBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TColorListBoxClass() TClass {
    return ColorListBox_StaticClassType()
}

// AddItem
func (c *TColorListBox) AddItem(Item string, AObject IObject) {
    ColorListBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

// Clear
func (c *TColorListBox) Clear() {
    ColorListBox_Clear(c.instance)
}

// ClearSelection
func (c *TColorListBox) ClearSelection() {
    ColorListBox_ClearSelection(c.instance)
}

// DeleteSelected
func (c *TColorListBox) DeleteSelected() {
    ColorListBox_DeleteSelected(c.instance)
}

// SelectAll
func (c *TColorListBox) SelectAll() {
    ColorListBox_SelectAll(c.instance)
}

// CanFocus
func (c *TColorListBox) CanFocus() bool {
    return ColorListBox_CanFocus(c.instance)
}

// FlipChildren
func (c *TColorListBox) FlipChildren(AllLevels bool) {
    ColorListBox_FlipChildren(c.instance, AllLevels)
}

// Focused
func (c *TColorListBox) Focused() bool {
    return ColorListBox_Focused(c.instance)
}

// HandleAllocated
func (c *TColorListBox) HandleAllocated() bool {
    return ColorListBox_HandleAllocated(c.instance)
}

// Invalidate
func (c *TColorListBox) Invalidate() {
    ColorListBox_Invalidate(c.instance)
}

// Realign
func (c *TColorListBox) Realign() {
    ColorListBox_Realign(c.instance)
}

// Repaint
func (c *TColorListBox) Repaint() {
    ColorListBox_Repaint(c.instance)
}

// ScaleBy
func (c *TColorListBox) ScaleBy(M int32, D int32) {
    ColorListBox_ScaleBy(c.instance, M , D)
}

// SetBounds
func (c *TColorListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ColorListBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (c *TColorListBox) SetFocus() {
    ColorListBox_SetFocus(c.instance)
}

// Update
func (c *TColorListBox) Update() {
    ColorListBox_Update(c.instance)
}

// BringToFront
func (c *TColorListBox) BringToFront() {
    ColorListBox_BringToFront(c.instance)
}

// ClientToScreen
func (c *TColorListBox) ClientToScreen(Point TPoint) TPoint {
    return ColorListBox_ClientToScreen(c.instance, Point)
}

// ClientToParent
func (c *TColorListBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ColorListBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
func (c *TColorListBox) Dragging() bool {
    return ColorListBox_Dragging(c.instance)
}

// HasParent
func (c *TColorListBox) HasParent() bool {
    return ColorListBox_HasParent(c.instance)
}

// Hide
func (c *TColorListBox) Hide() {
    ColorListBox_Hide(c.instance)
}

// Perform
func (c *TColorListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ColorListBox_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
func (c *TColorListBox) Refresh() {
    ColorListBox_Refresh(c.instance)
}

// ScreenToClient
func (c *TColorListBox) ScreenToClient(Point TPoint) TPoint {
    return ColorListBox_ScreenToClient(c.instance, Point)
}

// ParentToClient
func (c *TColorListBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ColorListBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (c *TColorListBox) SendToBack() {
    ColorListBox_SendToBack(c.instance)
}

// Show
func (c *TColorListBox) Show() {
    ColorListBox_Show(c.instance)
}

// GetTextBuf
func (c *TColorListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ColorListBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
func (c *TColorListBox) GetTextLen() int32 {
    return ColorListBox_GetTextLen(c.instance)
}

// FindComponent
func (c *TColorListBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ColorListBox_FindComponent(c.instance, AName))
}

// GetNamePath
func (c *TColorListBox) GetNamePath() string {
    return ColorListBox_GetNamePath(c.instance)
}

// Assign
func (c *TColorListBox) Assign(Source IObject) {
    ColorListBox_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TColorListBox) DisposeOf() {
    ColorListBox_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TColorListBox) ClassType() TClass {
    return ColorListBox_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TColorListBox) ClassName() string {
    return ColorListBox_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TColorListBox) InstanceSize() int32 {
    return ColorListBox_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TColorListBox) InheritsFrom(AClass TClass) bool {
    return ColorListBox_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TColorListBox) Equals(Obj IObject) bool {
    return ColorListBox_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TColorListBox) GetHashCode() int32 {
    return ColorListBox_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TColorListBox) ToString() string {
    return ColorListBox_ToString(c.instance)
}

// Align
func (c *TColorListBox) Align() TAlign {
    return ColorListBox_GetAlign(c.instance)
}

// SetAlign
func (c *TColorListBox) SetAlign(value TAlign) {
    ColorListBox_SetAlign(c.instance, value)
}

// AutoComplete
func (c *TColorListBox) AutoComplete() bool {
    return ColorListBox_GetAutoComplete(c.instance)
}

// SetAutoComplete
func (c *TColorListBox) SetAutoComplete(value bool) {
    ColorListBox_SetAutoComplete(c.instance, value)
}

// DefaultColorColor
func (c *TColorListBox) DefaultColorColor() TColor {
    return ColorListBox_GetDefaultColorColor(c.instance)
}

// SetDefaultColorColor
func (c *TColorListBox) SetDefaultColorColor(value TColor) {
    ColorListBox_SetDefaultColorColor(c.instance, value)
}

// NoneColorColor
func (c *TColorListBox) NoneColorColor() TColor {
    return ColorListBox_GetNoneColorColor(c.instance)
}

// SetNoneColorColor
func (c *TColorListBox) SetNoneColorColor(value TColor) {
    ColorListBox_SetNoneColorColor(c.instance, value)
}

// Selected
func (c *TColorListBox) Selected() TColor {
    return ColorListBox_GetSelected(c.instance)
}

// SetSelected
func (c *TColorListBox) SetSelected(value TColor) {
    ColorListBox_SetSelected(c.instance, value)
}

// Style
func (c *TColorListBox) Style() TColorBoxStyle {
    return ColorListBox_GetStyle(c.instance)
}

// SetStyle
func (c *TColorListBox) SetStyle(value TColorBoxStyle) {
    ColorListBox_SetStyle(c.instance, value)
}

// Anchors
func (c *TColorListBox) Anchors() TAnchors {
    return ColorListBox_GetAnchors(c.instance)
}

// SetAnchors
func (c *TColorListBox) SetAnchors(value TAnchors) {
    ColorListBox_SetAnchors(c.instance, value)
}

// BevelEdges
func (c *TColorListBox) BevelEdges() TBevelEdges {
    return ColorListBox_GetBevelEdges(c.instance)
}

// SetBevelEdges
func (c *TColorListBox) SetBevelEdges(value TBevelEdges) {
    ColorListBox_SetBevelEdges(c.instance, value)
}

// BevelInner
func (c *TColorListBox) BevelInner() TBevelCut {
    return ColorListBox_GetBevelInner(c.instance)
}

// SetBevelInner
func (c *TColorListBox) SetBevelInner(value TBevelCut) {
    ColorListBox_SetBevelInner(c.instance, value)
}

// BevelKind
func (c *TColorListBox) BevelKind() TBevelKind {
    return ColorListBox_GetBevelKind(c.instance)
}

// SetBevelKind
func (c *TColorListBox) SetBevelKind(value TBevelKind) {
    ColorListBox_SetBevelKind(c.instance, value)
}

// BevelOuter
func (c *TColorListBox) BevelOuter() TBevelCut {
    return ColorListBox_GetBevelOuter(c.instance)
}

// SetBevelOuter
func (c *TColorListBox) SetBevelOuter(value TBevelCut) {
    ColorListBox_SetBevelOuter(c.instance, value)
}

// BiDiMode
func (c *TColorListBox) BiDiMode() TBiDiMode {
    return ColorListBox_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TColorListBox) SetBiDiMode(value TBiDiMode) {
    ColorListBox_SetBiDiMode(c.instance, value)
}

// Color
func (c *TColorListBox) Color() TColor {
    return ColorListBox_GetColor(c.instance)
}

// SetColor
func (c *TColorListBox) SetColor(value TColor) {
    ColorListBox_SetColor(c.instance, value)
}

// DoubleBuffered
func (c *TColorListBox) DoubleBuffered() bool {
    return ColorListBox_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
func (c *TColorListBox) SetDoubleBuffered(value bool) {
    ColorListBox_SetDoubleBuffered(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TColorListBox) Enabled() bool {
    return ColorListBox_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TColorListBox) SetEnabled(value bool) {
    ColorListBox_SetEnabled(c.instance, value)
}

// Font
func (c *TColorListBox) Font() *TFont {
    return FontFromInst(ColorListBox_GetFont(c.instance))
}

// SetFont
func (c *TColorListBox) SetFont(value *TFont) {
    ColorListBox_SetFont(c.instance, CheckPtr(value))
}

// ItemHeight
func (c *TColorListBox) ItemHeight() int32 {
    return ColorListBox_GetItemHeight(c.instance)
}

// SetItemHeight
func (c *TColorListBox) SetItemHeight(value int32) {
    ColorListBox_SetItemHeight(c.instance, value)
}

// ParentColor
func (c *TColorListBox) ParentColor() bool {
    return ColorListBox_GetParentColor(c.instance)
}

// SetParentColor
func (c *TColorListBox) SetParentColor(value bool) {
    ColorListBox_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TColorListBox) ParentCtl3D() bool {
    return ColorListBox_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TColorListBox) SetParentCtl3D(value bool) {
    ColorListBox_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
func (c *TColorListBox) ParentDoubleBuffered() bool {
    return ColorListBox_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
func (c *TColorListBox) SetParentDoubleBuffered(value bool) {
    ColorListBox_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
func (c *TColorListBox) ParentFont() bool {
    return ColorListBox_GetParentFont(c.instance)
}

// SetParentFont
func (c *TColorListBox) SetParentFont(value bool) {
    ColorListBox_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TColorListBox) ParentShowHint() bool {
    return ColorListBox_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TColorListBox) SetParentShowHint(value bool) {
    ColorListBox_SetParentShowHint(c.instance, value)
}

// PopupMenu
func (c *TColorListBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ColorListBox_GetPopupMenu(c.instance))
}

// SetPopupMenu
func (c *TColorListBox) SetPopupMenu(value IComponent) {
    ColorListBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
func (c *TColorListBox) ShowHint() bool {
    return ColorListBox_GetShowHint(c.instance)
}

// SetShowHint
func (c *TColorListBox) SetShowHint(value bool) {
    ColorListBox_SetShowHint(c.instance, value)
}

// TabOrder
func (c *TColorListBox) TabOrder() uint16 {
    return ColorListBox_GetTabOrder(c.instance)
}

// SetTabOrder
func (c *TColorListBox) SetTabOrder(value uint16) {
    ColorListBox_SetTabOrder(c.instance, value)
}

// TabStop
func (c *TColorListBox) TabStop() bool {
    return ColorListBox_GetTabStop(c.instance)
}

// SetTabStop
func (c *TColorListBox) SetTabStop(value bool) {
    ColorListBox_SetTabStop(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TColorListBox) Visible() bool {
    return ColorListBox_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TColorListBox) SetVisible(value bool) {
    ColorListBox_SetVisible(c.instance, value)
}

// StyleElements
func (c *TColorListBox) StyleElements() TStyleElements {
    return ColorListBox_GetStyleElements(c.instance)
}

// SetStyleElements
func (c *TColorListBox) SetStyleElements(value TStyleElements) {
    ColorListBox_SetStyleElements(c.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TColorListBox) SetOnClick(fn TNotifyEvent) {
    ColorListBox_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
func (c *TColorListBox) SetOnContextPopup(fn TContextPopupEvent) {
    ColorListBox_SetOnContextPopup(c.instance, fn)
}

// SetOnDblClick
func (c *TColorListBox) SetOnDblClick(fn TNotifyEvent) {
    ColorListBox_SetOnDblClick(c.instance, fn)
}

// SetOnDragDrop
func (c *TColorListBox) SetOnDragDrop(fn TDragDropEvent) {
    ColorListBox_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
func (c *TColorListBox) SetOnDragOver(fn TDragOverEvent) {
    ColorListBox_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
func (c *TColorListBox) SetOnEndDock(fn TEndDragEvent) {
    ColorListBox_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
func (c *TColorListBox) SetOnEndDrag(fn TEndDragEvent) {
    ColorListBox_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
func (c *TColorListBox) SetOnEnter(fn TNotifyEvent) {
    ColorListBox_SetOnEnter(c.instance, fn)
}

// SetOnExit
func (c *TColorListBox) SetOnExit(fn TNotifyEvent) {
    ColorListBox_SetOnExit(c.instance, fn)
}

// SetOnKeyDown
func (c *TColorListBox) SetOnKeyDown(fn TKeyEvent) {
    ColorListBox_SetOnKeyDown(c.instance, fn)
}

// SetOnKeyPress
func (c *TColorListBox) SetOnKeyPress(fn TKeyPressEvent) {
    ColorListBox_SetOnKeyPress(c.instance, fn)
}

// SetOnKeyUp
func (c *TColorListBox) SetOnKeyUp(fn TKeyEvent) {
    ColorListBox_SetOnKeyUp(c.instance, fn)
}

// SetOnMouseDown
func (c *TColorListBox) SetOnMouseDown(fn TMouseEvent) {
    ColorListBox_SetOnMouseDown(c.instance, fn)
}

// SetOnMouseEnter
func (c *TColorListBox) SetOnMouseEnter(fn TNotifyEvent) {
    ColorListBox_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
func (c *TColorListBox) SetOnMouseLeave(fn TNotifyEvent) {
    ColorListBox_SetOnMouseLeave(c.instance, fn)
}

// SetOnMouseMove
func (c *TColorListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ColorListBox_SetOnMouseMove(c.instance, fn)
}

// SetOnMouseUp
func (c *TColorListBox) SetOnMouseUp(fn TMouseEvent) {
    ColorListBox_SetOnMouseUp(c.instance, fn)
}

// SetOnStartDock
func (c *TColorListBox) SetOnStartDock(fn TStartDockEvent) {
    ColorListBox_SetOnStartDock(c.instance, fn)
}

// AutoCompleteDelay
func (c *TColorListBox) AutoCompleteDelay() uint32 {
    return ColorListBox_GetAutoCompleteDelay(c.instance)
}

// SetAutoCompleteDelay
func (c *TColorListBox) SetAutoCompleteDelay(value uint32) {
    ColorListBox_SetAutoCompleteDelay(c.instance, value)
}

// Canvas
func (c *TColorListBox) Canvas() *TCanvas {
    return CanvasFromInst(ColorListBox_GetCanvas(c.instance))
}

// Count
func (c *TColorListBox) Count() int32 {
    return ColorListBox_GetCount(c.instance)
}

// SetCount
func (c *TColorListBox) SetCount(value int32) {
    ColorListBox_SetCount(c.instance, value)
}

// Items
func (c *TColorListBox) Items() *TStrings {
    return StringsFromInst(ColorListBox_GetItems(c.instance))
}

// SetItems
func (c *TColorListBox) SetItems(value IObject) {
    ColorListBox_SetItems(c.instance, CheckPtr(value))
}

// MultiSelect
func (c *TColorListBox) MultiSelect() bool {
    return ColorListBox_GetMultiSelect(c.instance)
}

// SetMultiSelect
func (c *TColorListBox) SetMultiSelect(value bool) {
    ColorListBox_SetMultiSelect(c.instance, value)
}

// SelCount
func (c *TColorListBox) SelCount() int32 {
    return ColorListBox_GetSelCount(c.instance)
}

// ItemIndex
func (c *TColorListBox) ItemIndex() int32 {
    return ColorListBox_GetItemIndex(c.instance)
}

// SetItemIndex
func (c *TColorListBox) SetItemIndex(value int32) {
    ColorListBox_SetItemIndex(c.instance, value)
}

// DockSite
func (c *TColorListBox) DockSite() bool {
    return ColorListBox_GetDockSite(c.instance)
}

// SetDockSite
func (c *TColorListBox) SetDockSite(value bool) {
    ColorListBox_SetDockSite(c.instance, value)
}

// Brush
func (c *TColorListBox) Brush() *TBrush {
    return BrushFromInst(ColorListBox_GetBrush(c.instance))
}

// ControlCount
func (c *TColorListBox) ControlCount() int32 {
    return ColorListBox_GetControlCount(c.instance)
}

// Handle
func (c *TColorListBox) Handle() HWND {
    return ColorListBox_GetHandle(c.instance)
}

// ParentWindow
func (c *TColorListBox) ParentWindow() HWND {
    return ColorListBox_GetParentWindow(c.instance)
}

// SetParentWindow
func (c *TColorListBox) SetParentWindow(value HWND) {
    ColorListBox_SetParentWindow(c.instance, value)
}

// UseDockManager
func (c *TColorListBox) UseDockManager() bool {
    return ColorListBox_GetUseDockManager(c.instance)
}

// SetUseDockManager
func (c *TColorListBox) SetUseDockManager(value bool) {
    ColorListBox_SetUseDockManager(c.instance, value)
}

// Action
func (c *TColorListBox) Action() *TAction {
    return ActionFromInst(ColorListBox_GetAction(c.instance))
}

// SetAction
func (c *TColorListBox) SetAction(value IComponent) {
    ColorListBox_SetAction(c.instance, CheckPtr(value))
}

// BoundsRect
func (c *TColorListBox) BoundsRect() TRect {
    return ColorListBox_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TColorListBox) SetBoundsRect(value TRect) {
    ColorListBox_SetBoundsRect(c.instance, value)
}

// ClientHeight
func (c *TColorListBox) ClientHeight() int32 {
    return ColorListBox_GetClientHeight(c.instance)
}

// SetClientHeight
func (c *TColorListBox) SetClientHeight(value int32) {
    ColorListBox_SetClientHeight(c.instance, value)
}

// ClientRect
func (c *TColorListBox) ClientRect() TRect {
    return ColorListBox_GetClientRect(c.instance)
}

// ClientWidth
func (c *TColorListBox) ClientWidth() int32 {
    return ColorListBox_GetClientWidth(c.instance)
}

// SetClientWidth
func (c *TColorListBox) SetClientWidth(value int32) {
    ColorListBox_SetClientWidth(c.instance, value)
}

// ExplicitLeft
func (c *TColorListBox) ExplicitLeft() int32 {
    return ColorListBox_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TColorListBox) ExplicitTop() int32 {
    return ColorListBox_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TColorListBox) ExplicitWidth() int32 {
    return ColorListBox_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TColorListBox) ExplicitHeight() int32 {
    return ColorListBox_GetExplicitHeight(c.instance)
}

// Floating
func (c *TColorListBox) Floating() bool {
    return ColorListBox_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TColorListBox) Parent() *TWinControl {
    return WinControlFromInst(ColorListBox_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TColorListBox) SetParent(value IWinControl) {
    ColorListBox_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
func (c *TColorListBox) AlignWithMargins() bool {
    return ColorListBox_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
func (c *TColorListBox) SetAlignWithMargins(value bool) {
    ColorListBox_SetAlignWithMargins(c.instance, value)
}

// Left
func (c *TColorListBox) Left() int32 {
    return ColorListBox_GetLeft(c.instance)
}

// SetLeft
func (c *TColorListBox) SetLeft(value int32) {
    ColorListBox_SetLeft(c.instance, value)
}

// Top
func (c *TColorListBox) Top() int32 {
    return ColorListBox_GetTop(c.instance)
}

// SetTop
func (c *TColorListBox) SetTop(value int32) {
    ColorListBox_SetTop(c.instance, value)
}

// Width
func (c *TColorListBox) Width() int32 {
    return ColorListBox_GetWidth(c.instance)
}

// SetWidth
func (c *TColorListBox) SetWidth(value int32) {
    ColorListBox_SetWidth(c.instance, value)
}

// Height
func (c *TColorListBox) Height() int32 {
    return ColorListBox_GetHeight(c.instance)
}

// SetHeight
func (c *TColorListBox) SetHeight(value int32) {
    ColorListBox_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TColorListBox) Cursor() TCursor {
    return ColorListBox_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TColorListBox) SetCursor(value TCursor) {
    ColorListBox_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (c *TColorListBox) Hint() string {
    return ColorListBox_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (c *TColorListBox) SetHint(value string) {
    ColorListBox_SetHint(c.instance, value)
}

// Margins
func (c *TColorListBox) Margins() *TMargins {
    return MarginsFromInst(ColorListBox_GetMargins(c.instance))
}

// SetMargins
func (c *TColorListBox) SetMargins(value *TMargins) {
    ColorListBox_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
func (c *TColorListBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ColorListBox_GetCustomHint(c.instance))
}

// SetCustomHint
func (c *TColorListBox) SetCustomHint(value IComponent) {
    ColorListBox_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TColorListBox) ComponentCount() int32 {
    return ColorListBox_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TColorListBox) ComponentIndex() int32 {
    return ColorListBox_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TColorListBox) SetComponentIndex(value int32) {
    ColorListBox_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TColorListBox) Owner() *TComponent {
    return ComponentFromInst(ColorListBox_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TColorListBox) Name() string {
    return ColorListBox_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TColorListBox) SetName(value string) {
    ColorListBox_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TColorListBox) Tag() int {
    return ColorListBox_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TColorListBox) SetTag(value int) {
    ColorListBox_SetTag(c.instance, value)
}

// Colors
func (c *TColorListBox) Colors(Index int32) TColor {
    return ColorListBox_GetColors(c.instance, Index)
}

// ColorNames
func (c *TColorListBox) ColorNames(Index int32) string {
    return ColorListBox_GetColorNames(c.instance, Index)
}

// Controls
func (c *TColorListBox) Controls(Index int32) *TControl {
    return ControlFromInst(ColorListBox_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TColorListBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ColorListBox_GetComponents(c.instance, AIndex))
}

