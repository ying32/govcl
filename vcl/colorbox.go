
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

type TColorBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewColorBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewColorBox(owner IComponent) *TColorBox {
    c := new(TColorBox)
    c.instance = ColorBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ColorBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ColorBoxFromInst(inst uintptr) *TColorBox {
    c := new(TColorBox)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ColorBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ColorBoxFromObj(obj IObject) *TColorBox {
    c := new(TColorBox)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ColorBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ColorBoxFromUnsafePointer(ptr unsafe.Pointer) *TColorBox {
    c := new(TColorBox)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TColorBox) Free() {
    if c.instance != 0 {
        ColorBox_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TColorBox) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TColorBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TColorBox) IsValid() bool {
    return c.instance != 0
}

// TColorBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TColorBoxClass() TClass {
    return ColorBox_StaticClassType()
}

// AddItem
func (c *TColorBox) AddItem(Item string, AObject IObject) {
    ColorBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

// Clear
func (c *TColorBox) Clear() {
    ColorBox_Clear(c.instance)
}

// ClearSelection
func (c *TColorBox) ClearSelection() {
    ColorBox_ClearSelection(c.instance)
}

// DeleteSelected
func (c *TColorBox) DeleteSelected() {
    ColorBox_DeleteSelected(c.instance)
}

// Focused
func (c *TColorBox) Focused() bool {
    return ColorBox_Focused(c.instance)
}

// SelectAll
func (c *TColorBox) SelectAll() {
    ColorBox_SelectAll(c.instance)
}

// CanFocus
func (c *TColorBox) CanFocus() bool {
    return ColorBox_CanFocus(c.instance)
}

// FlipChildren
func (c *TColorBox) FlipChildren(AllLevels bool) {
    ColorBox_FlipChildren(c.instance, AllLevels)
}

// HandleAllocated
func (c *TColorBox) HandleAllocated() bool {
    return ColorBox_HandleAllocated(c.instance)
}

// Invalidate
func (c *TColorBox) Invalidate() {
    ColorBox_Invalidate(c.instance)
}

// Realign
func (c *TColorBox) Realign() {
    ColorBox_Realign(c.instance)
}

// Repaint
func (c *TColorBox) Repaint() {
    ColorBox_Repaint(c.instance)
}

// ScaleBy
func (c *TColorBox) ScaleBy(M int32, D int32) {
    ColorBox_ScaleBy(c.instance, M , D)
}

// SetBounds
func (c *TColorBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ColorBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (c *TColorBox) SetFocus() {
    ColorBox_SetFocus(c.instance)
}

// Update
func (c *TColorBox) Update() {
    ColorBox_Update(c.instance)
}

// BringToFront
func (c *TColorBox) BringToFront() {
    ColorBox_BringToFront(c.instance)
}

// ClientToScreen
func (c *TColorBox) ClientToScreen(Point TPoint) TPoint {
    return ColorBox_ClientToScreen(c.instance, Point)
}

// ClientToParent
func (c *TColorBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ColorBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
func (c *TColorBox) Dragging() bool {
    return ColorBox_Dragging(c.instance)
}

// HasParent
func (c *TColorBox) HasParent() bool {
    return ColorBox_HasParent(c.instance)
}

// Hide
func (c *TColorBox) Hide() {
    ColorBox_Hide(c.instance)
}

// Perform
func (c *TColorBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ColorBox_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
func (c *TColorBox) Refresh() {
    ColorBox_Refresh(c.instance)
}

// ScreenToClient
func (c *TColorBox) ScreenToClient(Point TPoint) TPoint {
    return ColorBox_ScreenToClient(c.instance, Point)
}

// ParentToClient
func (c *TColorBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ColorBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (c *TColorBox) SendToBack() {
    ColorBox_SendToBack(c.instance)
}

// Show
func (c *TColorBox) Show() {
    ColorBox_Show(c.instance)
}

// GetTextBuf
func (c *TColorBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ColorBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
func (c *TColorBox) GetTextLen() int32 {
    return ColorBox_GetTextLen(c.instance)
}

// FindComponent
func (c *TColorBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ColorBox_FindComponent(c.instance, AName))
}

// GetNamePath
func (c *TColorBox) GetNamePath() string {
    return ColorBox_GetNamePath(c.instance)
}

// Assign
func (c *TColorBox) Assign(Source IObject) {
    ColorBox_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TColorBox) DisposeOf() {
    ColorBox_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TColorBox) ClassType() TClass {
    return ColorBox_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TColorBox) ClassName() string {
    return ColorBox_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TColorBox) InstanceSize() int32 {
    return ColorBox_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TColorBox) InheritsFrom(AClass TClass) bool {
    return ColorBox_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TColorBox) Equals(Obj IObject) bool {
    return ColorBox_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TColorBox) GetHashCode() int32 {
    return ColorBox_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TColorBox) ToString() string {
    return ColorBox_ToString(c.instance)
}

// Align
func (c *TColorBox) Align() TAlign {
    return ColorBox_GetAlign(c.instance)
}

// SetAlign
func (c *TColorBox) SetAlign(value TAlign) {
    ColorBox_SetAlign(c.instance, value)
}

// AutoComplete
func (c *TColorBox) AutoComplete() bool {
    return ColorBox_GetAutoComplete(c.instance)
}

// SetAutoComplete
func (c *TColorBox) SetAutoComplete(value bool) {
    ColorBox_SetAutoComplete(c.instance, value)
}

// AutoDropDown
func (c *TColorBox) AutoDropDown() bool {
    return ColorBox_GetAutoDropDown(c.instance)
}

// SetAutoDropDown
func (c *TColorBox) SetAutoDropDown(value bool) {
    ColorBox_SetAutoDropDown(c.instance, value)
}

// DefaultColorColor
func (c *TColorBox) DefaultColorColor() TColor {
    return ColorBox_GetDefaultColorColor(c.instance)
}

// SetDefaultColorColor
func (c *TColorBox) SetDefaultColorColor(value TColor) {
    ColorBox_SetDefaultColorColor(c.instance, value)
}

// NoneColorColor
func (c *TColorBox) NoneColorColor() TColor {
    return ColorBox_GetNoneColorColor(c.instance)
}

// SetNoneColorColor
func (c *TColorBox) SetNoneColorColor(value TColor) {
    ColorBox_SetNoneColorColor(c.instance, value)
}

// Selected
func (c *TColorBox) Selected() TColor {
    return ColorBox_GetSelected(c.instance)
}

// SetSelected
func (c *TColorBox) SetSelected(value TColor) {
    ColorBox_SetSelected(c.instance, value)
}

// Style
func (c *TColorBox) Style() TColorBoxStyle {
    return ColorBox_GetStyle(c.instance)
}

// SetStyle
func (c *TColorBox) SetStyle(value TColorBoxStyle) {
    ColorBox_SetStyle(c.instance, value)
}

// Anchors
func (c *TColorBox) Anchors() TAnchors {
    return ColorBox_GetAnchors(c.instance)
}

// SetAnchors
func (c *TColorBox) SetAnchors(value TAnchors) {
    ColorBox_SetAnchors(c.instance, value)
}

// BevelEdges
func (c *TColorBox) BevelEdges() TBevelEdges {
    return ColorBox_GetBevelEdges(c.instance)
}

// SetBevelEdges
func (c *TColorBox) SetBevelEdges(value TBevelEdges) {
    ColorBox_SetBevelEdges(c.instance, value)
}

// BevelInner
func (c *TColorBox) BevelInner() TBevelCut {
    return ColorBox_GetBevelInner(c.instance)
}

// SetBevelInner
func (c *TColorBox) SetBevelInner(value TBevelCut) {
    ColorBox_SetBevelInner(c.instance, value)
}

// BevelKind
func (c *TColorBox) BevelKind() TBevelKind {
    return ColorBox_GetBevelKind(c.instance)
}

// SetBevelKind
func (c *TColorBox) SetBevelKind(value TBevelKind) {
    ColorBox_SetBevelKind(c.instance, value)
}

// BevelOuter
func (c *TColorBox) BevelOuter() TBevelCut {
    return ColorBox_GetBevelOuter(c.instance)
}

// SetBevelOuter
func (c *TColorBox) SetBevelOuter(value TBevelCut) {
    ColorBox_SetBevelOuter(c.instance, value)
}

// BiDiMode
func (c *TColorBox) BiDiMode() TBiDiMode {
    return ColorBox_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TColorBox) SetBiDiMode(value TBiDiMode) {
    ColorBox_SetBiDiMode(c.instance, value)
}

// Color
func (c *TColorBox) Color() TColor {
    return ColorBox_GetColor(c.instance)
}

// SetColor
func (c *TColorBox) SetColor(value TColor) {
    ColorBox_SetColor(c.instance, value)
}

// DoubleBuffered
func (c *TColorBox) DoubleBuffered() bool {
    return ColorBox_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
func (c *TColorBox) SetDoubleBuffered(value bool) {
    ColorBox_SetDoubleBuffered(c.instance, value)
}

// DropDownCount
func (c *TColorBox) DropDownCount() int32 {
    return ColorBox_GetDropDownCount(c.instance)
}

// SetDropDownCount
func (c *TColorBox) SetDropDownCount(value int32) {
    ColorBox_SetDropDownCount(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TColorBox) Enabled() bool {
    return ColorBox_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TColorBox) SetEnabled(value bool) {
    ColorBox_SetEnabled(c.instance, value)
}

// Font
func (c *TColorBox) Font() *TFont {
    return FontFromInst(ColorBox_GetFont(c.instance))
}

// SetFont
func (c *TColorBox) SetFont(value *TFont) {
    ColorBox_SetFont(c.instance, CheckPtr(value))
}

// ItemHeight
func (c *TColorBox) ItemHeight() int32 {
    return ColorBox_GetItemHeight(c.instance)
}

// SetItemHeight
func (c *TColorBox) SetItemHeight(value int32) {
    ColorBox_SetItemHeight(c.instance, value)
}

// ParentColor
func (c *TColorBox) ParentColor() bool {
    return ColorBox_GetParentColor(c.instance)
}

// SetParentColor
func (c *TColorBox) SetParentColor(value bool) {
    ColorBox_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TColorBox) ParentCtl3D() bool {
    return ColorBox_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TColorBox) SetParentCtl3D(value bool) {
    ColorBox_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
func (c *TColorBox) ParentDoubleBuffered() bool {
    return ColorBox_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
func (c *TColorBox) SetParentDoubleBuffered(value bool) {
    ColorBox_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
func (c *TColorBox) ParentFont() bool {
    return ColorBox_GetParentFont(c.instance)
}

// SetParentFont
func (c *TColorBox) SetParentFont(value bool) {
    ColorBox_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TColorBox) ParentShowHint() bool {
    return ColorBox_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TColorBox) SetParentShowHint(value bool) {
    ColorBox_SetParentShowHint(c.instance, value)
}

// PopupMenu
func (c *TColorBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ColorBox_GetPopupMenu(c.instance))
}

// SetPopupMenu
func (c *TColorBox) SetPopupMenu(value IComponent) {
    ColorBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
func (c *TColorBox) ShowHint() bool {
    return ColorBox_GetShowHint(c.instance)
}

// SetShowHint
func (c *TColorBox) SetShowHint(value bool) {
    ColorBox_SetShowHint(c.instance, value)
}

// TabOrder
func (c *TColorBox) TabOrder() uint16 {
    return ColorBox_GetTabOrder(c.instance)
}

// SetTabOrder
func (c *TColorBox) SetTabOrder(value uint16) {
    ColorBox_SetTabOrder(c.instance, value)
}

// TabStop
func (c *TColorBox) TabStop() bool {
    return ColorBox_GetTabStop(c.instance)
}

// SetTabStop
func (c *TColorBox) SetTabStop(value bool) {
    ColorBox_SetTabStop(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TColorBox) Visible() bool {
    return ColorBox_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TColorBox) SetVisible(value bool) {
    ColorBox_SetVisible(c.instance, value)
}

// StyleElements
func (c *TColorBox) StyleElements() TStyleElements {
    return ColorBox_GetStyleElements(c.instance)
}

// SetStyleElements
func (c *TColorBox) SetStyleElements(value TStyleElements) {
    ColorBox_SetStyleElements(c.instance, value)
}

// SetOnChange
func (c *TColorBox) SetOnChange(fn TNotifyEvent) {
    ColorBox_SetOnChange(c.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TColorBox) SetOnClick(fn TNotifyEvent) {
    ColorBox_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
func (c *TColorBox) SetOnContextPopup(fn TContextPopupEvent) {
    ColorBox_SetOnContextPopup(c.instance, fn)
}

// SetOnDragDrop
func (c *TColorBox) SetOnDragDrop(fn TDragDropEvent) {
    ColorBox_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
func (c *TColorBox) SetOnDragOver(fn TDragOverEvent) {
    ColorBox_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
func (c *TColorBox) SetOnEndDock(fn TEndDragEvent) {
    ColorBox_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
func (c *TColorBox) SetOnEndDrag(fn TEndDragEvent) {
    ColorBox_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
func (c *TColorBox) SetOnEnter(fn TNotifyEvent) {
    ColorBox_SetOnEnter(c.instance, fn)
}

// SetOnExit
func (c *TColorBox) SetOnExit(fn TNotifyEvent) {
    ColorBox_SetOnExit(c.instance, fn)
}

// SetOnKeyDown
func (c *TColorBox) SetOnKeyDown(fn TKeyEvent) {
    ColorBox_SetOnKeyDown(c.instance, fn)
}

// SetOnKeyPress
func (c *TColorBox) SetOnKeyPress(fn TKeyPressEvent) {
    ColorBox_SetOnKeyPress(c.instance, fn)
}

// SetOnKeyUp
func (c *TColorBox) SetOnKeyUp(fn TKeyEvent) {
    ColorBox_SetOnKeyUp(c.instance, fn)
}

// SetOnMouseEnter
func (c *TColorBox) SetOnMouseEnter(fn TNotifyEvent) {
    ColorBox_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
func (c *TColorBox) SetOnMouseLeave(fn TNotifyEvent) {
    ColorBox_SetOnMouseLeave(c.instance, fn)
}

// SetOnStartDock
func (c *TColorBox) SetOnStartDock(fn TStartDockEvent) {
    ColorBox_SetOnStartDock(c.instance, fn)
}

// AutoCompleteDelay
func (c *TColorBox) AutoCompleteDelay() uint32 {
    return ColorBox_GetAutoCompleteDelay(c.instance)
}

// SetAutoCompleteDelay
func (c *TColorBox) SetAutoCompleteDelay(value uint32) {
    ColorBox_SetAutoCompleteDelay(c.instance, value)
}

// AutoCloseUp
func (c *TColorBox) AutoCloseUp() bool {
    return ColorBox_GetAutoCloseUp(c.instance)
}

// SetAutoCloseUp
func (c *TColorBox) SetAutoCloseUp(value bool) {
    ColorBox_SetAutoCloseUp(c.instance, value)
}

// CharCase
func (c *TColorBox) CharCase() TEditCharCase {
    return ColorBox_GetCharCase(c.instance)
}

// SetCharCase
func (c *TColorBox) SetCharCase(value TEditCharCase) {
    ColorBox_SetCharCase(c.instance, value)
}

// SelText
func (c *TColorBox) SelText() string {
    return ColorBox_GetSelText(c.instance)
}

// SetSelText
func (c *TColorBox) SetSelText(value string) {
    ColorBox_SetSelText(c.instance, value)
}

// TextHint
func (c *TColorBox) TextHint() string {
    return ColorBox_GetTextHint(c.instance)
}

// SetTextHint
func (c *TColorBox) SetTextHint(value string) {
    ColorBox_SetTextHint(c.instance, value)
}

// Canvas
func (c *TColorBox) Canvas() *TCanvas {
    return CanvasFromInst(ColorBox_GetCanvas(c.instance))
}

// DroppedDown
func (c *TColorBox) DroppedDown() bool {
    return ColorBox_GetDroppedDown(c.instance)
}

// SetDroppedDown
func (c *TColorBox) SetDroppedDown(value bool) {
    ColorBox_SetDroppedDown(c.instance, value)
}

// Items
func (c *TColorBox) Items() *TStrings {
    return StringsFromInst(ColorBox_GetItems(c.instance))
}

// SetItems
func (c *TColorBox) SetItems(value IObject) {
    ColorBox_SetItems(c.instance, CheckPtr(value))
}

// SelLength
func (c *TColorBox) SelLength() int32 {
    return ColorBox_GetSelLength(c.instance)
}

// SetSelLength
func (c *TColorBox) SetSelLength(value int32) {
    ColorBox_SetSelLength(c.instance, value)
}

// SelStart
func (c *TColorBox) SelStart() int32 {
    return ColorBox_GetSelStart(c.instance)
}

// SetSelStart
func (c *TColorBox) SetSelStart(value int32) {
    ColorBox_SetSelStart(c.instance, value)
}

// ItemIndex
func (c *TColorBox) ItemIndex() int32 {
    return ColorBox_GetItemIndex(c.instance)
}

// SetItemIndex
func (c *TColorBox) SetItemIndex(value int32) {
    ColorBox_SetItemIndex(c.instance, value)
}

// DockSite
func (c *TColorBox) DockSite() bool {
    return ColorBox_GetDockSite(c.instance)
}

// SetDockSite
func (c *TColorBox) SetDockSite(value bool) {
    ColorBox_SetDockSite(c.instance, value)
}

// Brush
func (c *TColorBox) Brush() *TBrush {
    return BrushFromInst(ColorBox_GetBrush(c.instance))
}

// ControlCount
func (c *TColorBox) ControlCount() int32 {
    return ColorBox_GetControlCount(c.instance)
}

// Handle
func (c *TColorBox) Handle() HWND {
    return ColorBox_GetHandle(c.instance)
}

// ParentWindow
func (c *TColorBox) ParentWindow() HWND {
    return ColorBox_GetParentWindow(c.instance)
}

// SetParentWindow
func (c *TColorBox) SetParentWindow(value HWND) {
    ColorBox_SetParentWindow(c.instance, value)
}

// UseDockManager
func (c *TColorBox) UseDockManager() bool {
    return ColorBox_GetUseDockManager(c.instance)
}

// SetUseDockManager
func (c *TColorBox) SetUseDockManager(value bool) {
    ColorBox_SetUseDockManager(c.instance, value)
}

// Action
func (c *TColorBox) Action() *TAction {
    return ActionFromInst(ColorBox_GetAction(c.instance))
}

// SetAction
func (c *TColorBox) SetAction(value IComponent) {
    ColorBox_SetAction(c.instance, CheckPtr(value))
}

// BoundsRect
func (c *TColorBox) BoundsRect() TRect {
    return ColorBox_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TColorBox) SetBoundsRect(value TRect) {
    ColorBox_SetBoundsRect(c.instance, value)
}

// ClientHeight
func (c *TColorBox) ClientHeight() int32 {
    return ColorBox_GetClientHeight(c.instance)
}

// SetClientHeight
func (c *TColorBox) SetClientHeight(value int32) {
    ColorBox_SetClientHeight(c.instance, value)
}

// ClientRect
func (c *TColorBox) ClientRect() TRect {
    return ColorBox_GetClientRect(c.instance)
}

// ClientWidth
func (c *TColorBox) ClientWidth() int32 {
    return ColorBox_GetClientWidth(c.instance)
}

// SetClientWidth
func (c *TColorBox) SetClientWidth(value int32) {
    ColorBox_SetClientWidth(c.instance, value)
}

// ExplicitLeft
func (c *TColorBox) ExplicitLeft() int32 {
    return ColorBox_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TColorBox) ExplicitTop() int32 {
    return ColorBox_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TColorBox) ExplicitWidth() int32 {
    return ColorBox_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TColorBox) ExplicitHeight() int32 {
    return ColorBox_GetExplicitHeight(c.instance)
}

// Floating
func (c *TColorBox) Floating() bool {
    return ColorBox_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TColorBox) Parent() *TWinControl {
    return WinControlFromInst(ColorBox_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TColorBox) SetParent(value IWinControl) {
    ColorBox_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
func (c *TColorBox) AlignWithMargins() bool {
    return ColorBox_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
func (c *TColorBox) SetAlignWithMargins(value bool) {
    ColorBox_SetAlignWithMargins(c.instance, value)
}

// Left
func (c *TColorBox) Left() int32 {
    return ColorBox_GetLeft(c.instance)
}

// SetLeft
func (c *TColorBox) SetLeft(value int32) {
    ColorBox_SetLeft(c.instance, value)
}

// Top
func (c *TColorBox) Top() int32 {
    return ColorBox_GetTop(c.instance)
}

// SetTop
func (c *TColorBox) SetTop(value int32) {
    ColorBox_SetTop(c.instance, value)
}

// Width
func (c *TColorBox) Width() int32 {
    return ColorBox_GetWidth(c.instance)
}

// SetWidth
func (c *TColorBox) SetWidth(value int32) {
    ColorBox_SetWidth(c.instance, value)
}

// Height
func (c *TColorBox) Height() int32 {
    return ColorBox_GetHeight(c.instance)
}

// SetHeight
func (c *TColorBox) SetHeight(value int32) {
    ColorBox_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TColorBox) Cursor() TCursor {
    return ColorBox_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TColorBox) SetCursor(value TCursor) {
    ColorBox_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (c *TColorBox) Hint() string {
    return ColorBox_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (c *TColorBox) SetHint(value string) {
    ColorBox_SetHint(c.instance, value)
}

// Margins
func (c *TColorBox) Margins() *TMargins {
    return MarginsFromInst(ColorBox_GetMargins(c.instance))
}

// SetMargins
func (c *TColorBox) SetMargins(value *TMargins) {
    ColorBox_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
func (c *TColorBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ColorBox_GetCustomHint(c.instance))
}

// SetCustomHint
func (c *TColorBox) SetCustomHint(value IComponent) {
    ColorBox_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TColorBox) ComponentCount() int32 {
    return ColorBox_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TColorBox) ComponentIndex() int32 {
    return ColorBox_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TColorBox) SetComponentIndex(value int32) {
    ColorBox_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TColorBox) Owner() *TComponent {
    return ComponentFromInst(ColorBox_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TColorBox) Name() string {
    return ColorBox_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TColorBox) SetName(value string) {
    ColorBox_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TColorBox) Tag() int {
    return ColorBox_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TColorBox) SetTag(value int) {
    ColorBox_SetTag(c.instance, value)
}

// Colors
func (c *TColorBox) Colors(Index int32) TColor {
    return ColorBox_GetColors(c.instance, Index)
}

// ColorNames
func (c *TColorBox) ColorNames(Index int32) string {
    return ColorBox_GetColorNames(c.instance, Index)
}

// Controls
func (c *TColorBox) Controls(Index int32) *TControl {
    return ControlFromInst(ColorBox_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TColorBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ColorBox_GetComponents(c.instance, AIndex))
}

