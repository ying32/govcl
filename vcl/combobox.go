
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

type TComboBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewComboBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewComboBox(owner IComponent) *TComboBox {
    c := new(TComboBox)
    c.instance = ComboBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ComboBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ComboBoxFromInst(inst uintptr) *TComboBox {
    c := new(TComboBox)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ComboBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ComboBoxFromObj(obj IObject) *TComboBox {
    c := new(TComboBox)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ComboBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ComboBoxFromUnsafePointer(ptr unsafe.Pointer) *TComboBox {
    c := new(TComboBox)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TComboBox) Free() {
    if c.instance != 0 {
        ComboBox_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TComboBox) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TComboBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TComboBox) IsValid() bool {
    return c.instance != 0
}

// TComboBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TComboBoxClass() TClass {
    return ComboBox_StaticClassType()
}

// AddItem
func (c *TComboBox) AddItem(Item string, AObject IObject) {
    ComboBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

// Clear
func (c *TComboBox) Clear() {
    ComboBox_Clear(c.instance)
}

// ClearSelection
func (c *TComboBox) ClearSelection() {
    ComboBox_ClearSelection(c.instance)
}

// DeleteSelected
func (c *TComboBox) DeleteSelected() {
    ComboBox_DeleteSelected(c.instance)
}

// Focused
func (c *TComboBox) Focused() bool {
    return ComboBox_Focused(c.instance)
}

// SelectAll
func (c *TComboBox) SelectAll() {
    ComboBox_SelectAll(c.instance)
}

// CanFocus
func (c *TComboBox) CanFocus() bool {
    return ComboBox_CanFocus(c.instance)
}

// ContainsControl
func (c *TComboBox) ContainsControl(Control IControl) bool {
    return ComboBox_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
func (c *TComboBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(ComboBox_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (c *TComboBox) DisableAlign() {
    ComboBox_DisableAlign(c.instance)
}

// EnableAlign
func (c *TComboBox) EnableAlign() {
    ComboBox_EnableAlign(c.instance)
}

// FindChildControl
func (c *TComboBox) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(ComboBox_FindChildControl(c.instance, ControlName))
}

// FlipChildren
func (c *TComboBox) FlipChildren(AllLevels bool) {
    ComboBox_FlipChildren(c.instance, AllLevels)
}

// HandleAllocated
func (c *TComboBox) HandleAllocated() bool {
    return ComboBox_HandleAllocated(c.instance)
}

// InsertControl
func (c *TComboBox) InsertControl(AControl IControl) {
    ComboBox_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
func (c *TComboBox) Invalidate() {
    ComboBox_Invalidate(c.instance)
}

// PaintTo
func (c *TComboBox) PaintTo(DC HDC, X int32, Y int32) {
    ComboBox_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
func (c *TComboBox) RemoveControl(AControl IControl) {
    ComboBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
func (c *TComboBox) Realign() {
    ComboBox_Realign(c.instance)
}

// Repaint
func (c *TComboBox) Repaint() {
    ComboBox_Repaint(c.instance)
}

// ScaleBy
func (c *TComboBox) ScaleBy(M int32, D int32) {
    ComboBox_ScaleBy(c.instance, M , D)
}

// ScrollBy
func (c *TComboBox) ScrollBy(DeltaX int32, DeltaY int32) {
    ComboBox_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetBounds
func (c *TComboBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ComboBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (c *TComboBox) SetFocus() {
    ComboBox_SetFocus(c.instance)
}

// Update
func (c *TComboBox) Update() {
    ComboBox_Update(c.instance)
}

// UpdateControlState
func (c *TComboBox) UpdateControlState() {
    ComboBox_UpdateControlState(c.instance)
}

// BringToFront
func (c *TComboBox) BringToFront() {
    ComboBox_BringToFront(c.instance)
}

// ClientToScreen
func (c *TComboBox) ClientToScreen(Point TPoint) TPoint {
    return ComboBox_ClientToScreen(c.instance, Point)
}

// ClientToParent
func (c *TComboBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ComboBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
func (c *TComboBox) Dragging() bool {
    return ComboBox_Dragging(c.instance)
}

// HasParent
func (c *TComboBox) HasParent() bool {
    return ComboBox_HasParent(c.instance)
}

// Hide
func (c *TComboBox) Hide() {
    ComboBox_Hide(c.instance)
}

// Perform
func (c *TComboBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ComboBox_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
func (c *TComboBox) Refresh() {
    ComboBox_Refresh(c.instance)
}

// ScreenToClient
func (c *TComboBox) ScreenToClient(Point TPoint) TPoint {
    return ComboBox_ScreenToClient(c.instance, Point)
}

// ParentToClient
func (c *TComboBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ComboBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (c *TComboBox) SendToBack() {
    ComboBox_SendToBack(c.instance)
}

// Show
func (c *TComboBox) Show() {
    ComboBox_Show(c.instance)
}

// GetTextBuf
func (c *TComboBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ComboBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
func (c *TComboBox) GetTextLen() int32 {
    return ComboBox_GetTextLen(c.instance)
}

// SetTextBuf
func (c *TComboBox) SetTextBuf(Buffer string) {
    ComboBox_SetTextBuf(c.instance, Buffer)
}

// FindComponent
func (c *TComboBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ComboBox_FindComponent(c.instance, AName))
}

// GetNamePath
func (c *TComboBox) GetNamePath() string {
    return ComboBox_GetNamePath(c.instance)
}

// Assign
func (c *TComboBox) Assign(Source IObject) {
    ComboBox_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TComboBox) DisposeOf() {
    ComboBox_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TComboBox) ClassType() TClass {
    return ComboBox_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TComboBox) ClassName() string {
    return ComboBox_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TComboBox) InstanceSize() int32 {
    return ComboBox_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TComboBox) InheritsFrom(AClass TClass) bool {
    return ComboBox_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TComboBox) Equals(Obj IObject) bool {
    return ComboBox_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TComboBox) GetHashCode() int32 {
    return ComboBox_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TComboBox) ToString() string {
    return ComboBox_ToString(c.instance)
}

// Align
func (c *TComboBox) Align() TAlign {
    return ComboBox_GetAlign(c.instance)
}

// SetAlign
func (c *TComboBox) SetAlign(value TAlign) {
    ComboBox_SetAlign(c.instance, value)
}

// AutoComplete
func (c *TComboBox) AutoComplete() bool {
    return ComboBox_GetAutoComplete(c.instance)
}

// SetAutoComplete
func (c *TComboBox) SetAutoComplete(value bool) {
    ComboBox_SetAutoComplete(c.instance, value)
}

// AutoCompleteDelay
func (c *TComboBox) AutoCompleteDelay() uint32 {
    return ComboBox_GetAutoCompleteDelay(c.instance)
}

// SetAutoCompleteDelay
func (c *TComboBox) SetAutoCompleteDelay(value uint32) {
    ComboBox_SetAutoCompleteDelay(c.instance, value)
}

// AutoDropDown
func (c *TComboBox) AutoDropDown() bool {
    return ComboBox_GetAutoDropDown(c.instance)
}

// SetAutoDropDown
func (c *TComboBox) SetAutoDropDown(value bool) {
    ComboBox_SetAutoDropDown(c.instance, value)
}

// AutoCloseUp
func (c *TComboBox) AutoCloseUp() bool {
    return ComboBox_GetAutoCloseUp(c.instance)
}

// SetAutoCloseUp
func (c *TComboBox) SetAutoCloseUp(value bool) {
    ComboBox_SetAutoCloseUp(c.instance, value)
}

// BevelEdges
func (c *TComboBox) BevelEdges() TBevelEdges {
    return ComboBox_GetBevelEdges(c.instance)
}

// SetBevelEdges
func (c *TComboBox) SetBevelEdges(value TBevelEdges) {
    ComboBox_SetBevelEdges(c.instance, value)
}

// BevelInner
func (c *TComboBox) BevelInner() TBevelCut {
    return ComboBox_GetBevelInner(c.instance)
}

// SetBevelInner
func (c *TComboBox) SetBevelInner(value TBevelCut) {
    ComboBox_SetBevelInner(c.instance, value)
}

// BevelKind
func (c *TComboBox) BevelKind() TBevelKind {
    return ComboBox_GetBevelKind(c.instance)
}

// SetBevelKind
func (c *TComboBox) SetBevelKind(value TBevelKind) {
    ComboBox_SetBevelKind(c.instance, value)
}

// BevelOuter
func (c *TComboBox) BevelOuter() TBevelCut {
    return ComboBox_GetBevelOuter(c.instance)
}

// SetBevelOuter
func (c *TComboBox) SetBevelOuter(value TBevelCut) {
    ComboBox_SetBevelOuter(c.instance, value)
}

// Style
func (c *TComboBox) Style() TComboBoxStyle {
    return ComboBox_GetStyle(c.instance)
}

// SetStyle
func (c *TComboBox) SetStyle(value TComboBoxStyle) {
    ComboBox_SetStyle(c.instance, value)
}

// Anchors
func (c *TComboBox) Anchors() TAnchors {
    return ComboBox_GetAnchors(c.instance)
}

// SetAnchors
func (c *TComboBox) SetAnchors(value TAnchors) {
    ComboBox_SetAnchors(c.instance, value)
}

// BiDiMode
func (c *TComboBox) BiDiMode() TBiDiMode {
    return ComboBox_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TComboBox) SetBiDiMode(value TBiDiMode) {
    ComboBox_SetBiDiMode(c.instance, value)
}

// CharCase
func (c *TComboBox) CharCase() TEditCharCase {
    return ComboBox_GetCharCase(c.instance)
}

// SetCharCase
func (c *TComboBox) SetCharCase(value TEditCharCase) {
    ComboBox_SetCharCase(c.instance, value)
}

// Color
func (c *TComboBox) Color() TColor {
    return ComboBox_GetColor(c.instance)
}

// SetColor
func (c *TComboBox) SetColor(value TColor) {
    ComboBox_SetColor(c.instance, value)
}

// DoubleBuffered
func (c *TComboBox) DoubleBuffered() bool {
    return ComboBox_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
func (c *TComboBox) SetDoubleBuffered(value bool) {
    ComboBox_SetDoubleBuffered(c.instance, value)
}

// DragCursor
func (c *TComboBox) DragCursor() TCursor {
    return ComboBox_GetDragCursor(c.instance)
}

// SetDragCursor
func (c *TComboBox) SetDragCursor(value TCursor) {
    ComboBox_SetDragCursor(c.instance, value)
}

// DragKind
func (c *TComboBox) DragKind() TDragKind {
    return ComboBox_GetDragKind(c.instance)
}

// SetDragKind
func (c *TComboBox) SetDragKind(value TDragKind) {
    ComboBox_SetDragKind(c.instance, value)
}

// DragMode
func (c *TComboBox) DragMode() TDragMode {
    return ComboBox_GetDragMode(c.instance)
}

// SetDragMode
func (c *TComboBox) SetDragMode(value TDragMode) {
    ComboBox_SetDragMode(c.instance, value)
}

// DropDownCount
func (c *TComboBox) DropDownCount() int32 {
    return ComboBox_GetDropDownCount(c.instance)
}

// SetDropDownCount
func (c *TComboBox) SetDropDownCount(value int32) {
    ComboBox_SetDropDownCount(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TComboBox) Enabled() bool {
    return ComboBox_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TComboBox) SetEnabled(value bool) {
    ComboBox_SetEnabled(c.instance, value)
}

// Font
func (c *TComboBox) Font() *TFont {
    return FontFromInst(ComboBox_GetFont(c.instance))
}

// SetFont
func (c *TComboBox) SetFont(value *TFont) {
    ComboBox_SetFont(c.instance, CheckPtr(value))
}

// ItemHeight
func (c *TComboBox) ItemHeight() int32 {
    return ComboBox_GetItemHeight(c.instance)
}

// SetItemHeight
func (c *TComboBox) SetItemHeight(value int32) {
    ComboBox_SetItemHeight(c.instance, value)
}

// ItemIndex
func (c *TComboBox) ItemIndex() int32 {
    return ComboBox_GetItemIndex(c.instance)
}

// SetItemIndex
func (c *TComboBox) SetItemIndex(value int32) {
    ComboBox_SetItemIndex(c.instance, value)
}

// MaxLength
func (c *TComboBox) MaxLength() int32 {
    return ComboBox_GetMaxLength(c.instance)
}

// SetMaxLength
func (c *TComboBox) SetMaxLength(value int32) {
    ComboBox_SetMaxLength(c.instance, value)
}

// ParentColor
func (c *TComboBox) ParentColor() bool {
    return ComboBox_GetParentColor(c.instance)
}

// SetParentColor
func (c *TComboBox) SetParentColor(value bool) {
    ComboBox_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TComboBox) ParentCtl3D() bool {
    return ComboBox_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TComboBox) SetParentCtl3D(value bool) {
    ComboBox_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
func (c *TComboBox) ParentDoubleBuffered() bool {
    return ComboBox_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
func (c *TComboBox) SetParentDoubleBuffered(value bool) {
    ComboBox_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
func (c *TComboBox) ParentFont() bool {
    return ComboBox_GetParentFont(c.instance)
}

// SetParentFont
func (c *TComboBox) SetParentFont(value bool) {
    ComboBox_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TComboBox) ParentShowHint() bool {
    return ComboBox_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TComboBox) SetParentShowHint(value bool) {
    ComboBox_SetParentShowHint(c.instance, value)
}

// PopupMenu
func (c *TComboBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ComboBox_GetPopupMenu(c.instance))
}

// SetPopupMenu
func (c *TComboBox) SetPopupMenu(value IComponent) {
    ComboBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
func (c *TComboBox) ShowHint() bool {
    return ComboBox_GetShowHint(c.instance)
}

// SetShowHint
func (c *TComboBox) SetShowHint(value bool) {
    ComboBox_SetShowHint(c.instance, value)
}

// Sorted
func (c *TComboBox) Sorted() bool {
    return ComboBox_GetSorted(c.instance)
}

// SetSorted
func (c *TComboBox) SetSorted(value bool) {
    ComboBox_SetSorted(c.instance, value)
}

// TabOrder
func (c *TComboBox) TabOrder() TTabOrder {
    return ComboBox_GetTabOrder(c.instance)
}

// SetTabOrder
func (c *TComboBox) SetTabOrder(value TTabOrder) {
    ComboBox_SetTabOrder(c.instance, value)
}

// TabStop
func (c *TComboBox) TabStop() bool {
    return ComboBox_GetTabStop(c.instance)
}

// SetTabStop
func (c *TComboBox) SetTabStop(value bool) {
    ComboBox_SetTabStop(c.instance, value)
}

// Text
func (c *TComboBox) Text() string {
    return ComboBox_GetText(c.instance)
}

// SetText
func (c *TComboBox) SetText(value string) {
    ComboBox_SetText(c.instance, value)
}

// TextHint
func (c *TComboBox) TextHint() string {
    return ComboBox_GetTextHint(c.instance)
}

// SetTextHint
func (c *TComboBox) SetTextHint(value string) {
    ComboBox_SetTextHint(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TComboBox) Visible() bool {
    return ComboBox_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TComboBox) SetVisible(value bool) {
    ComboBox_SetVisible(c.instance, value)
}

// StyleElements
func (c *TComboBox) StyleElements() TStyleElements {
    return ComboBox_GetStyleElements(c.instance)
}

// SetStyleElements
func (c *TComboBox) SetStyleElements(value TStyleElements) {
    ComboBox_SetStyleElements(c.instance, value)
}

// SetOnChange
func (c *TComboBox) SetOnChange(fn TNotifyEvent) {
    ComboBox_SetOnChange(c.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TComboBox) SetOnClick(fn TNotifyEvent) {
    ComboBox_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
func (c *TComboBox) SetOnContextPopup(fn TContextPopupEvent) {
    ComboBox_SetOnContextPopup(c.instance, fn)
}

// SetOnDblClick
func (c *TComboBox) SetOnDblClick(fn TNotifyEvent) {
    ComboBox_SetOnDblClick(c.instance, fn)
}

// SetOnDragDrop
func (c *TComboBox) SetOnDragDrop(fn TDragDropEvent) {
    ComboBox_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
func (c *TComboBox) SetOnDragOver(fn TDragOverEvent) {
    ComboBox_SetOnDragOver(c.instance, fn)
}

// SetOnDrawItem
func (c *TComboBox) SetOnDrawItem(fn TDrawItemEvent) {
    ComboBox_SetOnDrawItem(c.instance, fn)
}

// SetOnEndDock
func (c *TComboBox) SetOnEndDock(fn TEndDragEvent) {
    ComboBox_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
func (c *TComboBox) SetOnEndDrag(fn TEndDragEvent) {
    ComboBox_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
func (c *TComboBox) SetOnEnter(fn TNotifyEvent) {
    ComboBox_SetOnEnter(c.instance, fn)
}

// SetOnExit
func (c *TComboBox) SetOnExit(fn TNotifyEvent) {
    ComboBox_SetOnExit(c.instance, fn)
}

// SetOnKeyDown
func (c *TComboBox) SetOnKeyDown(fn TKeyEvent) {
    ComboBox_SetOnKeyDown(c.instance, fn)
}

// SetOnKeyPress
func (c *TComboBox) SetOnKeyPress(fn TKeyPressEvent) {
    ComboBox_SetOnKeyPress(c.instance, fn)
}

// SetOnKeyUp
func (c *TComboBox) SetOnKeyUp(fn TKeyEvent) {
    ComboBox_SetOnKeyUp(c.instance, fn)
}

// SetOnMouseEnter
func (c *TComboBox) SetOnMouseEnter(fn TNotifyEvent) {
    ComboBox_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
func (c *TComboBox) SetOnMouseLeave(fn TNotifyEvent) {
    ComboBox_SetOnMouseLeave(c.instance, fn)
}

// SetOnStartDock
func (c *TComboBox) SetOnStartDock(fn TStartDockEvent) {
    ComboBox_SetOnStartDock(c.instance, fn)
}

// Items
func (c *TComboBox) Items() *TStrings {
    return StringsFromInst(ComboBox_GetItems(c.instance))
}

// SetItems
func (c *TComboBox) SetItems(value IObject) {
    ComboBox_SetItems(c.instance, CheckPtr(value))
}

// SelText
func (c *TComboBox) SelText() string {
    return ComboBox_GetSelText(c.instance)
}

// SetSelText
func (c *TComboBox) SetSelText(value string) {
    ComboBox_SetSelText(c.instance, value)
}

// Canvas
func (c *TComboBox) Canvas() *TCanvas {
    return CanvasFromInst(ComboBox_GetCanvas(c.instance))
}

// DroppedDown
func (c *TComboBox) DroppedDown() bool {
    return ComboBox_GetDroppedDown(c.instance)
}

// SetDroppedDown
func (c *TComboBox) SetDroppedDown(value bool) {
    ComboBox_SetDroppedDown(c.instance, value)
}

// SelLength
func (c *TComboBox) SelLength() int32 {
    return ComboBox_GetSelLength(c.instance)
}

// SetSelLength
func (c *TComboBox) SetSelLength(value int32) {
    ComboBox_SetSelLength(c.instance, value)
}

// SelStart
func (c *TComboBox) SelStart() int32 {
    return ComboBox_GetSelStart(c.instance)
}

// SetSelStart
func (c *TComboBox) SetSelStart(value int32) {
    ComboBox_SetSelStart(c.instance, value)
}

// DockClientCount
func (c *TComboBox) DockClientCount() int32 {
    return ComboBox_GetDockClientCount(c.instance)
}

// DockSite
func (c *TComboBox) DockSite() bool {
    return ComboBox_GetDockSite(c.instance)
}

// SetDockSite
func (c *TComboBox) SetDockSite(value bool) {
    ComboBox_SetDockSite(c.instance, value)
}

// AlignDisabled
func (c *TComboBox) AlignDisabled() bool {
    return ComboBox_GetAlignDisabled(c.instance)
}

// MouseInClient
func (c *TComboBox) MouseInClient() bool {
    return ComboBox_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
func (c *TComboBox) VisibleDockClientCount() int32 {
    return ComboBox_GetVisibleDockClientCount(c.instance)
}

// Brush
func (c *TComboBox) Brush() *TBrush {
    return BrushFromInst(ComboBox_GetBrush(c.instance))
}

// ControlCount
func (c *TComboBox) ControlCount() int32 {
    return ComboBox_GetControlCount(c.instance)
}

// Handle
func (c *TComboBox) Handle() HWND {
    return ComboBox_GetHandle(c.instance)
}

// ParentWindow
func (c *TComboBox) ParentWindow() HWND {
    return ComboBox_GetParentWindow(c.instance)
}

// SetParentWindow
func (c *TComboBox) SetParentWindow(value HWND) {
    ComboBox_SetParentWindow(c.instance, value)
}

// UseDockManager
func (c *TComboBox) UseDockManager() bool {
    return ComboBox_GetUseDockManager(c.instance)
}

// SetUseDockManager
func (c *TComboBox) SetUseDockManager(value bool) {
    ComboBox_SetUseDockManager(c.instance, value)
}

// Action
func (c *TComboBox) Action() *TAction {
    return ActionFromInst(ComboBox_GetAction(c.instance))
}

// SetAction
func (c *TComboBox) SetAction(value IComponent) {
    ComboBox_SetAction(c.instance, CheckPtr(value))
}

// BoundsRect
func (c *TComboBox) BoundsRect() TRect {
    return ComboBox_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TComboBox) SetBoundsRect(value TRect) {
    ComboBox_SetBoundsRect(c.instance, value)
}

// ClientHeight
func (c *TComboBox) ClientHeight() int32 {
    return ComboBox_GetClientHeight(c.instance)
}

// SetClientHeight
func (c *TComboBox) SetClientHeight(value int32) {
    ComboBox_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TComboBox) ClientOrigin() TPoint {
    return ComboBox_GetClientOrigin(c.instance)
}

// ClientRect
func (c *TComboBox) ClientRect() TRect {
    return ComboBox_GetClientRect(c.instance)
}

// ClientWidth
func (c *TComboBox) ClientWidth() int32 {
    return ComboBox_GetClientWidth(c.instance)
}

// SetClientWidth
func (c *TComboBox) SetClientWidth(value int32) {
    ComboBox_SetClientWidth(c.instance, value)
}

// ControlState
func (c *TComboBox) ControlState() TControlState {
    return ComboBox_GetControlState(c.instance)
}

// SetControlState
func (c *TComboBox) SetControlState(value TControlState) {
    ComboBox_SetControlState(c.instance, value)
}

// ControlStyle
func (c *TComboBox) ControlStyle() TControlStyle {
    return ComboBox_GetControlStyle(c.instance)
}

// SetControlStyle
func (c *TComboBox) SetControlStyle(value TControlStyle) {
    ComboBox_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TComboBox) ExplicitLeft() int32 {
    return ComboBox_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TComboBox) ExplicitTop() int32 {
    return ComboBox_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TComboBox) ExplicitWidth() int32 {
    return ComboBox_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TComboBox) ExplicitHeight() int32 {
    return ComboBox_GetExplicitHeight(c.instance)
}

// Floating
func (c *TComboBox) Floating() bool {
    return ComboBox_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TComboBox) Parent() *TWinControl {
    return WinControlFromInst(ComboBox_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TComboBox) SetParent(value IWinControl) {
    ComboBox_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
func (c *TComboBox) AlignWithMargins() bool {
    return ComboBox_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
func (c *TComboBox) SetAlignWithMargins(value bool) {
    ComboBox_SetAlignWithMargins(c.instance, value)
}

// Left
func (c *TComboBox) Left() int32 {
    return ComboBox_GetLeft(c.instance)
}

// SetLeft
func (c *TComboBox) SetLeft(value int32) {
    ComboBox_SetLeft(c.instance, value)
}

// Top
func (c *TComboBox) Top() int32 {
    return ComboBox_GetTop(c.instance)
}

// SetTop
func (c *TComboBox) SetTop(value int32) {
    ComboBox_SetTop(c.instance, value)
}

// Width
func (c *TComboBox) Width() int32 {
    return ComboBox_GetWidth(c.instance)
}

// SetWidth
func (c *TComboBox) SetWidth(value int32) {
    ComboBox_SetWidth(c.instance, value)
}

// Height
func (c *TComboBox) Height() int32 {
    return ComboBox_GetHeight(c.instance)
}

// SetHeight
func (c *TComboBox) SetHeight(value int32) {
    ComboBox_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TComboBox) Cursor() TCursor {
    return ComboBox_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TComboBox) SetCursor(value TCursor) {
    ComboBox_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (c *TComboBox) Hint() string {
    return ComboBox_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (c *TComboBox) SetHint(value string) {
    ComboBox_SetHint(c.instance, value)
}

// Margins
func (c *TComboBox) Margins() *TMargins {
    return MarginsFromInst(ComboBox_GetMargins(c.instance))
}

// SetMargins
func (c *TComboBox) SetMargins(value *TMargins) {
    ComboBox_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
func (c *TComboBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ComboBox_GetCustomHint(c.instance))
}

// SetCustomHint
func (c *TComboBox) SetCustomHint(value IComponent) {
    ComboBox_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TComboBox) ComponentCount() int32 {
    return ComboBox_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TComboBox) ComponentIndex() int32 {
    return ComboBox_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TComboBox) SetComponentIndex(value int32) {
    ComboBox_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TComboBox) Owner() *TComponent {
    return ComponentFromInst(ComboBox_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TComboBox) Name() string {
    return ComboBox_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TComboBox) SetName(value string) {
    ComboBox_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TComboBox) Tag() int {
    return ComboBox_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TComboBox) SetTag(value int) {
    ComboBox_SetTag(c.instance, value)
}

// DockClients
func (c *TComboBox) DockClients(Index int32) *TControl {
    return ControlFromInst(ComboBox_GetDockClients(c.instance, Index))
}

// Controls
func (c *TComboBox) Controls(Index int32) *TControl {
    return ControlFromInst(ComboBox_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TComboBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ComboBox_GetComponents(c.instance, AIndex))
}

