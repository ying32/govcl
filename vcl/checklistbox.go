
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

type TCheckListBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCheckListBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCheckListBox(owner IComponent) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = CheckListBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CheckListBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CheckListBoxFromInst(inst uintptr) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CheckListBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CheckListBoxFromObj(obj IObject) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CheckListBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CheckListBoxFromUnsafePointer(ptr unsafe.Pointer) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCheckListBox) Free() {
    if c.instance != 0 {
        CheckListBox_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCheckListBox) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCheckListBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCheckListBox) IsValid() bool {
    return c.instance != 0
}

// TCheckListBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCheckListBoxClass() TClass {
    return CheckListBox_StaticClassType()
}

// CheckAll
func (c *TCheckListBox) CheckAll(AState TCheckBoxState, AllowGrayed bool, AllowDisabled bool) {
    CheckListBox_CheckAll(c.instance, AState , AllowGrayed , AllowDisabled)
}

// AddItem
func (c *TCheckListBox) AddItem(Item string, AObject IObject) {
    CheckListBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

// Clear
func (c *TCheckListBox) Clear() {
    CheckListBox_Clear(c.instance)
}

// ClearSelection
func (c *TCheckListBox) ClearSelection() {
    CheckListBox_ClearSelection(c.instance)
}

// DeleteSelected
func (c *TCheckListBox) DeleteSelected() {
    CheckListBox_DeleteSelected(c.instance)
}

// SelectAll
func (c *TCheckListBox) SelectAll() {
    CheckListBox_SelectAll(c.instance)
}

// CanFocus
func (c *TCheckListBox) CanFocus() bool {
    return CheckListBox_CanFocus(c.instance)
}

// ContainsControl
func (c *TCheckListBox) ContainsControl(Control IControl) bool {
    return CheckListBox_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
func (c *TCheckListBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(CheckListBox_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (c *TCheckListBox) DisableAlign() {
    CheckListBox_DisableAlign(c.instance)
}

// EnableAlign
func (c *TCheckListBox) EnableAlign() {
    CheckListBox_EnableAlign(c.instance)
}

// FindChildControl
func (c *TCheckListBox) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(CheckListBox_FindChildControl(c.instance, ControlName))
}

// FlipChildren
func (c *TCheckListBox) FlipChildren(AllLevels bool) {
    CheckListBox_FlipChildren(c.instance, AllLevels)
}

// Focused
func (c *TCheckListBox) Focused() bool {
    return CheckListBox_Focused(c.instance)
}

// HandleAllocated
func (c *TCheckListBox) HandleAllocated() bool {
    return CheckListBox_HandleAllocated(c.instance)
}

// InsertControl
func (c *TCheckListBox) InsertControl(AControl IControl) {
    CheckListBox_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
func (c *TCheckListBox) Invalidate() {
    CheckListBox_Invalidate(c.instance)
}

// PaintTo
func (c *TCheckListBox) PaintTo(DC HDC, X int32, Y int32) {
    CheckListBox_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
func (c *TCheckListBox) RemoveControl(AControl IControl) {
    CheckListBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
func (c *TCheckListBox) Realign() {
    CheckListBox_Realign(c.instance)
}

// Repaint
func (c *TCheckListBox) Repaint() {
    CheckListBox_Repaint(c.instance)
}

// ScaleBy
func (c *TCheckListBox) ScaleBy(M int32, D int32) {
    CheckListBox_ScaleBy(c.instance, M , D)
}

// ScrollBy
func (c *TCheckListBox) ScrollBy(DeltaX int32, DeltaY int32) {
    CheckListBox_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetBounds
func (c *TCheckListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CheckListBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (c *TCheckListBox) SetFocus() {
    CheckListBox_SetFocus(c.instance)
}

// Update
func (c *TCheckListBox) Update() {
    CheckListBox_Update(c.instance)
}

// UpdateControlState
func (c *TCheckListBox) UpdateControlState() {
    CheckListBox_UpdateControlState(c.instance)
}

// BringToFront
func (c *TCheckListBox) BringToFront() {
    CheckListBox_BringToFront(c.instance)
}

// ClientToScreen
func (c *TCheckListBox) ClientToScreen(Point TPoint) TPoint {
    return CheckListBox_ClientToScreen(c.instance, Point)
}

// ClientToParent
func (c *TCheckListBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CheckListBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
func (c *TCheckListBox) Dragging() bool {
    return CheckListBox_Dragging(c.instance)
}

// HasParent
func (c *TCheckListBox) HasParent() bool {
    return CheckListBox_HasParent(c.instance)
}

// Hide
func (c *TCheckListBox) Hide() {
    CheckListBox_Hide(c.instance)
}

// Perform
func (c *TCheckListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CheckListBox_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
func (c *TCheckListBox) Refresh() {
    CheckListBox_Refresh(c.instance)
}

// ScreenToClient
func (c *TCheckListBox) ScreenToClient(Point TPoint) TPoint {
    return CheckListBox_ScreenToClient(c.instance, Point)
}

// ParentToClient
func (c *TCheckListBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CheckListBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (c *TCheckListBox) SendToBack() {
    CheckListBox_SendToBack(c.instance)
}

// Show
func (c *TCheckListBox) Show() {
    CheckListBox_Show(c.instance)
}

// GetTextBuf
func (c *TCheckListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CheckListBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
func (c *TCheckListBox) GetTextLen() int32 {
    return CheckListBox_GetTextLen(c.instance)
}

// SetTextBuf
func (c *TCheckListBox) SetTextBuf(Buffer string) {
    CheckListBox_SetTextBuf(c.instance, Buffer)
}

// FindComponent
func (c *TCheckListBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CheckListBox_FindComponent(c.instance, AName))
}

// GetNamePath
func (c *TCheckListBox) GetNamePath() string {
    return CheckListBox_GetNamePath(c.instance)
}

// Assign
func (c *TCheckListBox) Assign(Source IObject) {
    CheckListBox_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCheckListBox) DisposeOf() {
    CheckListBox_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCheckListBox) ClassType() TClass {
    return CheckListBox_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCheckListBox) ClassName() string {
    return CheckListBox_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCheckListBox) InstanceSize() int32 {
    return CheckListBox_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCheckListBox) InheritsFrom(AClass TClass) bool {
    return CheckListBox_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCheckListBox) Equals(Obj IObject) bool {
    return CheckListBox_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCheckListBox) GetHashCode() int32 {
    return CheckListBox_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCheckListBox) ToString() string {
    return CheckListBox_ToString(c.instance)
}

// SetOnClickCheck
func (c *TCheckListBox) SetOnClickCheck(fn TNotifyEvent) {
    CheckListBox_SetOnClickCheck(c.instance, fn)
}

// Align
func (c *TCheckListBox) Align() TAlign {
    return CheckListBox_GetAlign(c.instance)
}

// SetAlign
func (c *TCheckListBox) SetAlign(value TAlign) {
    CheckListBox_SetAlign(c.instance, value)
}

// AllowGrayed
func (c *TCheckListBox) AllowGrayed() bool {
    return CheckListBox_GetAllowGrayed(c.instance)
}

// SetAllowGrayed
func (c *TCheckListBox) SetAllowGrayed(value bool) {
    CheckListBox_SetAllowGrayed(c.instance, value)
}

// Anchors
func (c *TCheckListBox) Anchors() TAnchors {
    return CheckListBox_GetAnchors(c.instance)
}

// SetAnchors
func (c *TCheckListBox) SetAnchors(value TAnchors) {
    CheckListBox_SetAnchors(c.instance, value)
}

// AutoComplete
func (c *TCheckListBox) AutoComplete() bool {
    return CheckListBox_GetAutoComplete(c.instance)
}

// SetAutoComplete
func (c *TCheckListBox) SetAutoComplete(value bool) {
    CheckListBox_SetAutoComplete(c.instance, value)
}

// BevelEdges
func (c *TCheckListBox) BevelEdges() TBevelEdges {
    return CheckListBox_GetBevelEdges(c.instance)
}

// SetBevelEdges
func (c *TCheckListBox) SetBevelEdges(value TBevelEdges) {
    CheckListBox_SetBevelEdges(c.instance, value)
}

// BevelInner
func (c *TCheckListBox) BevelInner() TBevelCut {
    return CheckListBox_GetBevelInner(c.instance)
}

// SetBevelInner
func (c *TCheckListBox) SetBevelInner(value TBevelCut) {
    CheckListBox_SetBevelInner(c.instance, value)
}

// BevelOuter
func (c *TCheckListBox) BevelOuter() TBevelCut {
    return CheckListBox_GetBevelOuter(c.instance)
}

// SetBevelOuter
func (c *TCheckListBox) SetBevelOuter(value TBevelCut) {
    CheckListBox_SetBevelOuter(c.instance, value)
}

// BevelKind
func (c *TCheckListBox) BevelKind() TBevelKind {
    return CheckListBox_GetBevelKind(c.instance)
}

// SetBevelKind
func (c *TCheckListBox) SetBevelKind(value TBevelKind) {
    CheckListBox_SetBevelKind(c.instance, value)
}

// BiDiMode
func (c *TCheckListBox) BiDiMode() TBiDiMode {
    return CheckListBox_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TCheckListBox) SetBiDiMode(value TBiDiMode) {
    CheckListBox_SetBiDiMode(c.instance, value)
}

// BorderStyle
func (c *TCheckListBox) BorderStyle() TBorderStyle {
    return CheckListBox_GetBorderStyle(c.instance)
}

// SetBorderStyle
func (c *TCheckListBox) SetBorderStyle(value TBorderStyle) {
    CheckListBox_SetBorderStyle(c.instance, value)
}

// Color
func (c *TCheckListBox) Color() TColor {
    return CheckListBox_GetColor(c.instance)
}

// SetColor
func (c *TCheckListBox) SetColor(value TColor) {
    CheckListBox_SetColor(c.instance, value)
}

// Columns
func (c *TCheckListBox) Columns() int32 {
    return CheckListBox_GetColumns(c.instance)
}

// SetColumns
func (c *TCheckListBox) SetColumns(value int32) {
    CheckListBox_SetColumns(c.instance, value)
}

// DoubleBuffered
func (c *TCheckListBox) DoubleBuffered() bool {
    return CheckListBox_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
func (c *TCheckListBox) SetDoubleBuffered(value bool) {
    CheckListBox_SetDoubleBuffered(c.instance, value)
}

// DragCursor
func (c *TCheckListBox) DragCursor() TCursor {
    return CheckListBox_GetDragCursor(c.instance)
}

// SetDragCursor
func (c *TCheckListBox) SetDragCursor(value TCursor) {
    CheckListBox_SetDragCursor(c.instance, value)
}

// DragKind
func (c *TCheckListBox) DragKind() TDragKind {
    return CheckListBox_GetDragKind(c.instance)
}

// SetDragKind
func (c *TCheckListBox) SetDragKind(value TDragKind) {
    CheckListBox_SetDragKind(c.instance, value)
}

// DragMode
func (c *TCheckListBox) DragMode() TDragMode {
    return CheckListBox_GetDragMode(c.instance)
}

// SetDragMode
func (c *TCheckListBox) SetDragMode(value TDragMode) {
    CheckListBox_SetDragMode(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TCheckListBox) Enabled() bool {
    return CheckListBox_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TCheckListBox) SetEnabled(value bool) {
    CheckListBox_SetEnabled(c.instance, value)
}

// Flat
func (c *TCheckListBox) Flat() bool {
    return CheckListBox_GetFlat(c.instance)
}

// SetFlat
func (c *TCheckListBox) SetFlat(value bool) {
    CheckListBox_SetFlat(c.instance, value)
}

// Font
func (c *TCheckListBox) Font() *TFont {
    return FontFromInst(CheckListBox_GetFont(c.instance))
}

// SetFont
func (c *TCheckListBox) SetFont(value *TFont) {
    CheckListBox_SetFont(c.instance, CheckPtr(value))
}

// HeaderColor
func (c *TCheckListBox) HeaderColor() TColor {
    return CheckListBox_GetHeaderColor(c.instance)
}

// SetHeaderColor
func (c *TCheckListBox) SetHeaderColor(value TColor) {
    CheckListBox_SetHeaderColor(c.instance, value)
}

// HeaderBackgroundColor
func (c *TCheckListBox) HeaderBackgroundColor() TColor {
    return CheckListBox_GetHeaderBackgroundColor(c.instance)
}

// SetHeaderBackgroundColor
func (c *TCheckListBox) SetHeaderBackgroundColor(value TColor) {
    CheckListBox_SetHeaderBackgroundColor(c.instance, value)
}

// ItemHeight
func (c *TCheckListBox) ItemHeight() int32 {
    return CheckListBox_GetItemHeight(c.instance)
}

// SetItemHeight
func (c *TCheckListBox) SetItemHeight(value int32) {
    CheckListBox_SetItemHeight(c.instance, value)
}

// Items
func (c *TCheckListBox) Items() *TStrings {
    return StringsFromInst(CheckListBox_GetItems(c.instance))
}

// SetItems
func (c *TCheckListBox) SetItems(value IObject) {
    CheckListBox_SetItems(c.instance, CheckPtr(value))
}

// ParentColor
func (c *TCheckListBox) ParentColor() bool {
    return CheckListBox_GetParentColor(c.instance)
}

// SetParentColor
func (c *TCheckListBox) SetParentColor(value bool) {
    CheckListBox_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TCheckListBox) ParentCtl3D() bool {
    return CheckListBox_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TCheckListBox) SetParentCtl3D(value bool) {
    CheckListBox_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
func (c *TCheckListBox) ParentDoubleBuffered() bool {
    return CheckListBox_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
func (c *TCheckListBox) SetParentDoubleBuffered(value bool) {
    CheckListBox_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
func (c *TCheckListBox) ParentFont() bool {
    return CheckListBox_GetParentFont(c.instance)
}

// SetParentFont
func (c *TCheckListBox) SetParentFont(value bool) {
    CheckListBox_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TCheckListBox) ParentShowHint() bool {
    return CheckListBox_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TCheckListBox) SetParentShowHint(value bool) {
    CheckListBox_SetParentShowHint(c.instance, value)
}

// PopupMenu
func (c *TCheckListBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CheckListBox_GetPopupMenu(c.instance))
}

// SetPopupMenu
func (c *TCheckListBox) SetPopupMenu(value IComponent) {
    CheckListBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
func (c *TCheckListBox) ShowHint() bool {
    return CheckListBox_GetShowHint(c.instance)
}

// SetShowHint
func (c *TCheckListBox) SetShowHint(value bool) {
    CheckListBox_SetShowHint(c.instance, value)
}

// Sorted
func (c *TCheckListBox) Sorted() bool {
    return CheckListBox_GetSorted(c.instance)
}

// SetSorted
func (c *TCheckListBox) SetSorted(value bool) {
    CheckListBox_SetSorted(c.instance, value)
}

// Style
func (c *TCheckListBox) Style() TListBoxStyle {
    return CheckListBox_GetStyle(c.instance)
}

// SetStyle
func (c *TCheckListBox) SetStyle(value TListBoxStyle) {
    CheckListBox_SetStyle(c.instance, value)
}

// TabOrder
func (c *TCheckListBox) TabOrder() TTabOrder {
    return CheckListBox_GetTabOrder(c.instance)
}

// SetTabOrder
func (c *TCheckListBox) SetTabOrder(value TTabOrder) {
    CheckListBox_SetTabOrder(c.instance, value)
}

// TabStop
func (c *TCheckListBox) TabStop() bool {
    return CheckListBox_GetTabStop(c.instance)
}

// SetTabStop
func (c *TCheckListBox) SetTabStop(value bool) {
    CheckListBox_SetTabStop(c.instance, value)
}

// TabWidth
func (c *TCheckListBox) TabWidth() int32 {
    return CheckListBox_GetTabWidth(c.instance)
}

// SetTabWidth
func (c *TCheckListBox) SetTabWidth(value int32) {
    CheckListBox_SetTabWidth(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TCheckListBox) Visible() bool {
    return CheckListBox_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TCheckListBox) SetVisible(value bool) {
    CheckListBox_SetVisible(c.instance, value)
}

// StyleElements
func (c *TCheckListBox) StyleElements() TStyleElements {
    return CheckListBox_GetStyleElements(c.instance)
}

// SetStyleElements
func (c *TCheckListBox) SetStyleElements(value TStyleElements) {
    CheckListBox_SetStyleElements(c.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TCheckListBox) SetOnClick(fn TNotifyEvent) {
    CheckListBox_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
func (c *TCheckListBox) SetOnContextPopup(fn TContextPopupEvent) {
    CheckListBox_SetOnContextPopup(c.instance, fn)
}

// SetOnDblClick
func (c *TCheckListBox) SetOnDblClick(fn TNotifyEvent) {
    CheckListBox_SetOnDblClick(c.instance, fn)
}

// SetOnDragDrop
func (c *TCheckListBox) SetOnDragDrop(fn TDragDropEvent) {
    CheckListBox_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
func (c *TCheckListBox) SetOnDragOver(fn TDragOverEvent) {
    CheckListBox_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
func (c *TCheckListBox) SetOnEndDock(fn TEndDragEvent) {
    CheckListBox_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
func (c *TCheckListBox) SetOnEndDrag(fn TEndDragEvent) {
    CheckListBox_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
func (c *TCheckListBox) SetOnEnter(fn TNotifyEvent) {
    CheckListBox_SetOnEnter(c.instance, fn)
}

// SetOnExit
func (c *TCheckListBox) SetOnExit(fn TNotifyEvent) {
    CheckListBox_SetOnExit(c.instance, fn)
}

// SetOnKeyDown
func (c *TCheckListBox) SetOnKeyDown(fn TKeyEvent) {
    CheckListBox_SetOnKeyDown(c.instance, fn)
}

// SetOnKeyPress
func (c *TCheckListBox) SetOnKeyPress(fn TKeyPressEvent) {
    CheckListBox_SetOnKeyPress(c.instance, fn)
}

// SetOnKeyUp
func (c *TCheckListBox) SetOnKeyUp(fn TKeyEvent) {
    CheckListBox_SetOnKeyUp(c.instance, fn)
}

// SetOnMouseDown
func (c *TCheckListBox) SetOnMouseDown(fn TMouseEvent) {
    CheckListBox_SetOnMouseDown(c.instance, fn)
}

// SetOnMouseEnter
func (c *TCheckListBox) SetOnMouseEnter(fn TNotifyEvent) {
    CheckListBox_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
func (c *TCheckListBox) SetOnMouseLeave(fn TNotifyEvent) {
    CheckListBox_SetOnMouseLeave(c.instance, fn)
}

// SetOnMouseMove
func (c *TCheckListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    CheckListBox_SetOnMouseMove(c.instance, fn)
}

// SetOnMouseUp
func (c *TCheckListBox) SetOnMouseUp(fn TMouseEvent) {
    CheckListBox_SetOnMouseUp(c.instance, fn)
}

// SetOnStartDock
func (c *TCheckListBox) SetOnStartDock(fn TStartDockEvent) {
    CheckListBox_SetOnStartDock(c.instance, fn)
}

// AutoCompleteDelay
func (c *TCheckListBox) AutoCompleteDelay() uint32 {
    return CheckListBox_GetAutoCompleteDelay(c.instance)
}

// SetAutoCompleteDelay
func (c *TCheckListBox) SetAutoCompleteDelay(value uint32) {
    CheckListBox_SetAutoCompleteDelay(c.instance, value)
}

// Canvas
func (c *TCheckListBox) Canvas() *TCanvas {
    return CanvasFromInst(CheckListBox_GetCanvas(c.instance))
}

// Count
func (c *TCheckListBox) Count() int32 {
    return CheckListBox_GetCount(c.instance)
}

// SetCount
func (c *TCheckListBox) SetCount(value int32) {
    CheckListBox_SetCount(c.instance, value)
}

// MultiSelect
func (c *TCheckListBox) MultiSelect() bool {
    return CheckListBox_GetMultiSelect(c.instance)
}

// SetMultiSelect
func (c *TCheckListBox) SetMultiSelect(value bool) {
    CheckListBox_SetMultiSelect(c.instance, value)
}

// SelCount
func (c *TCheckListBox) SelCount() int32 {
    return CheckListBox_GetSelCount(c.instance)
}

// ItemIndex
func (c *TCheckListBox) ItemIndex() int32 {
    return CheckListBox_GetItemIndex(c.instance)
}

// SetItemIndex
func (c *TCheckListBox) SetItemIndex(value int32) {
    CheckListBox_SetItemIndex(c.instance, value)
}

// DockClientCount
func (c *TCheckListBox) DockClientCount() int32 {
    return CheckListBox_GetDockClientCount(c.instance)
}

// DockSite
func (c *TCheckListBox) DockSite() bool {
    return CheckListBox_GetDockSite(c.instance)
}

// SetDockSite
func (c *TCheckListBox) SetDockSite(value bool) {
    CheckListBox_SetDockSite(c.instance, value)
}

// AlignDisabled
func (c *TCheckListBox) AlignDisabled() bool {
    return CheckListBox_GetAlignDisabled(c.instance)
}

// MouseInClient
func (c *TCheckListBox) MouseInClient() bool {
    return CheckListBox_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
func (c *TCheckListBox) VisibleDockClientCount() int32 {
    return CheckListBox_GetVisibleDockClientCount(c.instance)
}

// Brush
func (c *TCheckListBox) Brush() *TBrush {
    return BrushFromInst(CheckListBox_GetBrush(c.instance))
}

// ControlCount
func (c *TCheckListBox) ControlCount() int32 {
    return CheckListBox_GetControlCount(c.instance)
}

// Handle
func (c *TCheckListBox) Handle() HWND {
    return CheckListBox_GetHandle(c.instance)
}

// ParentWindow
func (c *TCheckListBox) ParentWindow() HWND {
    return CheckListBox_GetParentWindow(c.instance)
}

// SetParentWindow
func (c *TCheckListBox) SetParentWindow(value HWND) {
    CheckListBox_SetParentWindow(c.instance, value)
}

// UseDockManager
func (c *TCheckListBox) UseDockManager() bool {
    return CheckListBox_GetUseDockManager(c.instance)
}

// SetUseDockManager
func (c *TCheckListBox) SetUseDockManager(value bool) {
    CheckListBox_SetUseDockManager(c.instance, value)
}

// Action
func (c *TCheckListBox) Action() *TAction {
    return ActionFromInst(CheckListBox_GetAction(c.instance))
}

// SetAction
func (c *TCheckListBox) SetAction(value IComponent) {
    CheckListBox_SetAction(c.instance, CheckPtr(value))
}

// BoundsRect
func (c *TCheckListBox) BoundsRect() TRect {
    return CheckListBox_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TCheckListBox) SetBoundsRect(value TRect) {
    CheckListBox_SetBoundsRect(c.instance, value)
}

// ClientHeight
func (c *TCheckListBox) ClientHeight() int32 {
    return CheckListBox_GetClientHeight(c.instance)
}

// SetClientHeight
func (c *TCheckListBox) SetClientHeight(value int32) {
    CheckListBox_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TCheckListBox) ClientOrigin() TPoint {
    return CheckListBox_GetClientOrigin(c.instance)
}

// ClientRect
func (c *TCheckListBox) ClientRect() TRect {
    return CheckListBox_GetClientRect(c.instance)
}

// ClientWidth
func (c *TCheckListBox) ClientWidth() int32 {
    return CheckListBox_GetClientWidth(c.instance)
}

// SetClientWidth
func (c *TCheckListBox) SetClientWidth(value int32) {
    CheckListBox_SetClientWidth(c.instance, value)
}

// ControlState
func (c *TCheckListBox) ControlState() TControlState {
    return CheckListBox_GetControlState(c.instance)
}

// SetControlState
func (c *TCheckListBox) SetControlState(value TControlState) {
    CheckListBox_SetControlState(c.instance, value)
}

// ControlStyle
func (c *TCheckListBox) ControlStyle() TControlStyle {
    return CheckListBox_GetControlStyle(c.instance)
}

// SetControlStyle
func (c *TCheckListBox) SetControlStyle(value TControlStyle) {
    CheckListBox_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TCheckListBox) ExplicitLeft() int32 {
    return CheckListBox_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TCheckListBox) ExplicitTop() int32 {
    return CheckListBox_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TCheckListBox) ExplicitWidth() int32 {
    return CheckListBox_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TCheckListBox) ExplicitHeight() int32 {
    return CheckListBox_GetExplicitHeight(c.instance)
}

// Floating
func (c *TCheckListBox) Floating() bool {
    return CheckListBox_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TCheckListBox) Parent() *TWinControl {
    return WinControlFromInst(CheckListBox_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TCheckListBox) SetParent(value IWinControl) {
    CheckListBox_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
func (c *TCheckListBox) AlignWithMargins() bool {
    return CheckListBox_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
func (c *TCheckListBox) SetAlignWithMargins(value bool) {
    CheckListBox_SetAlignWithMargins(c.instance, value)
}

// Left
func (c *TCheckListBox) Left() int32 {
    return CheckListBox_GetLeft(c.instance)
}

// SetLeft
func (c *TCheckListBox) SetLeft(value int32) {
    CheckListBox_SetLeft(c.instance, value)
}

// Top
func (c *TCheckListBox) Top() int32 {
    return CheckListBox_GetTop(c.instance)
}

// SetTop
func (c *TCheckListBox) SetTop(value int32) {
    CheckListBox_SetTop(c.instance, value)
}

// Width
func (c *TCheckListBox) Width() int32 {
    return CheckListBox_GetWidth(c.instance)
}

// SetWidth
func (c *TCheckListBox) SetWidth(value int32) {
    CheckListBox_SetWidth(c.instance, value)
}

// Height
func (c *TCheckListBox) Height() int32 {
    return CheckListBox_GetHeight(c.instance)
}

// SetHeight
func (c *TCheckListBox) SetHeight(value int32) {
    CheckListBox_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TCheckListBox) Cursor() TCursor {
    return CheckListBox_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TCheckListBox) SetCursor(value TCursor) {
    CheckListBox_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (c *TCheckListBox) Hint() string {
    return CheckListBox_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (c *TCheckListBox) SetHint(value string) {
    CheckListBox_SetHint(c.instance, value)
}

// Margins
func (c *TCheckListBox) Margins() *TMargins {
    return MarginsFromInst(CheckListBox_GetMargins(c.instance))
}

// SetMargins
func (c *TCheckListBox) SetMargins(value *TMargins) {
    CheckListBox_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
func (c *TCheckListBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(CheckListBox_GetCustomHint(c.instance))
}

// SetCustomHint
func (c *TCheckListBox) SetCustomHint(value IComponent) {
    CheckListBox_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TCheckListBox) ComponentCount() int32 {
    return CheckListBox_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TCheckListBox) ComponentIndex() int32 {
    return CheckListBox_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TCheckListBox) SetComponentIndex(value int32) {
    CheckListBox_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TCheckListBox) Owner() *TComponent {
    return ComponentFromInst(CheckListBox_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TCheckListBox) Name() string {
    return CheckListBox_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TCheckListBox) SetName(value string) {
    CheckListBox_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TCheckListBox) Tag() int {
    return CheckListBox_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TCheckListBox) SetTag(value int) {
    CheckListBox_SetTag(c.instance, value)
}

// Checked
func (c *TCheckListBox) Checked(Index int32) bool {
    return CheckListBox_GetChecked(c.instance, Index)
}

// Checked
func (c *TCheckListBox) SetChecked(Index int32, value bool) {
    CheckListBox_SetChecked(c.instance, Index, value)
}

// ItemEnabled
func (c *TCheckListBox) ItemEnabled(Index int32) bool {
    return CheckListBox_GetItemEnabled(c.instance, Index)
}

// ItemEnabled
func (c *TCheckListBox) SetItemEnabled(Index int32, value bool) {
    CheckListBox_SetItemEnabled(c.instance, Index, value)
}

// State
func (c *TCheckListBox) State(Index int32) TCheckBoxState {
    return CheckListBox_GetState(c.instance, Index)
}

// State
func (c *TCheckListBox) SetState(Index int32, value TCheckBoxState) {
    CheckListBox_SetState(c.instance, Index, value)
}

// Header
func (c *TCheckListBox) Header(Index int32) bool {
    return CheckListBox_GetHeader(c.instance, Index)
}

// Header
func (c *TCheckListBox) SetHeader(Index int32, value bool) {
    CheckListBox_SetHeader(c.instance, Index, value)
}

// Selected
func (c *TCheckListBox) Selected(Index int32) bool {
    return CheckListBox_GetSelected(c.instance, Index)
}

// Selected
func (c *TCheckListBox) SetSelected(Index int32, value bool) {
    CheckListBox_SetSelected(c.instance, Index, value)
}

// DockClients
func (c *TCheckListBox) DockClients(Index int32) *TControl {
    return ControlFromInst(CheckListBox_GetDockClients(c.instance, Index))
}

// Controls
func (c *TCheckListBox) Controls(Index int32) *TControl {
    return ControlFromInst(CheckListBox_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TCheckListBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CheckListBox_GetComponents(c.instance, AIndex))
}

