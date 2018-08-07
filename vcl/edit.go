
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

type TEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewEdit
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewEdit(owner IComponent) *TEdit {
    e := new(TEdit)
    e.instance = Edit_Create(CheckPtr(owner))
    e.ptr = unsafe.Pointer(e.instance)
    return e
}

// EditFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func EditFromInst(inst uintptr) *TEdit {
    e := new(TEdit)
    e.instance = inst
    e.ptr = unsafe.Pointer(inst)
    return e
}

// EditFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func EditFromObj(obj IObject) *TEdit {
    e := new(TEdit)
    e.instance = CheckPtr(obj)
    e.ptr = unsafe.Pointer(e.instance)
    return e
}

// EditFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func EditFromUnsafePointer(ptr unsafe.Pointer) *TEdit {
    e := new(TEdit)
    e.instance = uintptr(ptr)
    e.ptr = ptr
    return e
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (e *TEdit) Free() {
    if e.instance != 0 {
        Edit_Free(e.instance)
        e.instance = 0
        e.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (e *TEdit) Instance() uintptr {
    return e.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (e *TEdit) UnsafeAddr() unsafe.Pointer {
    return e.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (e *TEdit) IsValid() bool {
    return e.instance != 0
}

// TEditClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TEditClass() TClass {
    return Edit_StaticClassType()
}

// Clear
func (e *TEdit) Clear() {
    Edit_Clear(e.instance)
}

// ClearSelection
func (e *TEdit) ClearSelection() {
    Edit_ClearSelection(e.instance)
}

// CopyToClipboard
func (e *TEdit) CopyToClipboard() {
    Edit_CopyToClipboard(e.instance)
}

// CutToClipboard
func (e *TEdit) CutToClipboard() {
    Edit_CutToClipboard(e.instance)
}

// PasteFromClipboard
func (e *TEdit) PasteFromClipboard() {
    Edit_PasteFromClipboard(e.instance)
}

// Undo
func (e *TEdit) Undo() {
    Edit_Undo(e.instance)
}

// ClearUndo
func (e *TEdit) ClearUndo() {
    Edit_ClearUndo(e.instance)
}

// SelectAll
func (e *TEdit) SelectAll() {
    Edit_SelectAll(e.instance)
}

// GetSelTextBuf
func (e *TEdit) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return Edit_GetSelTextBuf(e.instance, Buffer , BufSize)
}

// CanFocus
func (e *TEdit) CanFocus() bool {
    return Edit_CanFocus(e.instance)
}

// FlipChildren
func (e *TEdit) FlipChildren(AllLevels bool) {
    Edit_FlipChildren(e.instance, AllLevels)
}

// Focused
func (e *TEdit) Focused() bool {
    return Edit_Focused(e.instance)
}

// HandleAllocated
func (e *TEdit) HandleAllocated() bool {
    return Edit_HandleAllocated(e.instance)
}

// Invalidate
func (e *TEdit) Invalidate() {
    Edit_Invalidate(e.instance)
}

// Realign
func (e *TEdit) Realign() {
    Edit_Realign(e.instance)
}

// Repaint
func (e *TEdit) Repaint() {
    Edit_Repaint(e.instance)
}

// ScaleBy
func (e *TEdit) ScaleBy(M int32, D int32) {
    Edit_ScaleBy(e.instance, M , D)
}

// SetBounds
func (e *TEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Edit_SetBounds(e.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (e *TEdit) SetFocus() {
    Edit_SetFocus(e.instance)
}

// Update
func (e *TEdit) Update() {
    Edit_Update(e.instance)
}

// BringToFront
func (e *TEdit) BringToFront() {
    Edit_BringToFront(e.instance)
}

// ClientToScreen
func (e *TEdit) ClientToScreen(Point TPoint) TPoint {
    return Edit_ClientToScreen(e.instance, Point)
}

// ClientToParent
func (e *TEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Edit_ClientToParent(e.instance, Point , CheckPtr(AParent))
}

// Dragging
func (e *TEdit) Dragging() bool {
    return Edit_Dragging(e.instance)
}

// HasParent
func (e *TEdit) HasParent() bool {
    return Edit_HasParent(e.instance)
}

// Hide
func (e *TEdit) Hide() {
    Edit_Hide(e.instance)
}

// Perform
func (e *TEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Edit_Perform(e.instance, Msg , WParam , LParam)
}

// Refresh
func (e *TEdit) Refresh() {
    Edit_Refresh(e.instance)
}

// ScreenToClient
func (e *TEdit) ScreenToClient(Point TPoint) TPoint {
    return Edit_ScreenToClient(e.instance, Point)
}

// ParentToClient
func (e *TEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Edit_ParentToClient(e.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (e *TEdit) SendToBack() {
    Edit_SendToBack(e.instance)
}

// Show
func (e *TEdit) Show() {
    Edit_Show(e.instance)
}

// GetTextBuf
func (e *TEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Edit_GetTextBuf(e.instance, Buffer , BufSize)
}

// GetTextLen
func (e *TEdit) GetTextLen() int32 {
    return Edit_GetTextLen(e.instance)
}

// FindComponent
func (e *TEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Edit_FindComponent(e.instance, AName))
}

// GetNamePath
func (e *TEdit) GetNamePath() string {
    return Edit_GetNamePath(e.instance)
}

// Assign
func (e *TEdit) Assign(Source IObject) {
    Edit_Assign(e.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (e *TEdit) DisposeOf() {
    Edit_DisposeOf(e.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (e *TEdit) ClassType() TClass {
    return Edit_ClassType(e.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (e *TEdit) ClassName() string {
    return Edit_ClassName(e.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (e *TEdit) InstanceSize() int32 {
    return Edit_InstanceSize(e.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (e *TEdit) InheritsFrom(AClass TClass) bool {
    return Edit_InheritsFrom(e.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (e *TEdit) Equals(Obj IObject) bool {
    return Edit_Equals(e.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (e *TEdit) GetHashCode() int32 {
    return Edit_GetHashCode(e.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (e *TEdit) ToString() string {
    return Edit_ToString(e.instance)
}

// Align
func (e *TEdit) Align() TAlign {
    return Edit_GetAlign(e.instance)
}

// SetAlign
func (e *TEdit) SetAlign(value TAlign) {
    Edit_SetAlign(e.instance, value)
}

// Alignment
func (e *TEdit) Alignment() TAlignment {
    return Edit_GetAlignment(e.instance)
}

// SetAlignment
func (e *TEdit) SetAlignment(value TAlignment) {
    Edit_SetAlignment(e.instance, value)
}

// Anchors
func (e *TEdit) Anchors() TAnchors {
    return Edit_GetAnchors(e.instance)
}

// SetAnchors
func (e *TEdit) SetAnchors(value TAnchors) {
    Edit_SetAnchors(e.instance, value)
}

// AutoSelect
func (e *TEdit) AutoSelect() bool {
    return Edit_GetAutoSelect(e.instance)
}

// SetAutoSelect
func (e *TEdit) SetAutoSelect(value bool) {
    Edit_SetAutoSelect(e.instance, value)
}

// AutoSize
func (e *TEdit) AutoSize() bool {
    return Edit_GetAutoSize(e.instance)
}

// SetAutoSize
func (e *TEdit) SetAutoSize(value bool) {
    Edit_SetAutoSize(e.instance, value)
}

// BevelEdges
func (e *TEdit) BevelEdges() TBevelEdges {
    return Edit_GetBevelEdges(e.instance)
}

// SetBevelEdges
func (e *TEdit) SetBevelEdges(value TBevelEdges) {
    Edit_SetBevelEdges(e.instance, value)
}

// BevelInner
func (e *TEdit) BevelInner() TBevelCut {
    return Edit_GetBevelInner(e.instance)
}

// SetBevelInner
func (e *TEdit) SetBevelInner(value TBevelCut) {
    Edit_SetBevelInner(e.instance, value)
}

// BevelKind
func (e *TEdit) BevelKind() TBevelKind {
    return Edit_GetBevelKind(e.instance)
}

// SetBevelKind
func (e *TEdit) SetBevelKind(value TBevelKind) {
    Edit_SetBevelKind(e.instance, value)
}

// BevelOuter
func (e *TEdit) BevelOuter() TBevelCut {
    return Edit_GetBevelOuter(e.instance)
}

// SetBevelOuter
func (e *TEdit) SetBevelOuter(value TBevelCut) {
    Edit_SetBevelOuter(e.instance, value)
}

// BiDiMode
func (e *TEdit) BiDiMode() TBiDiMode {
    return Edit_GetBiDiMode(e.instance)
}

// SetBiDiMode
func (e *TEdit) SetBiDiMode(value TBiDiMode) {
    Edit_SetBiDiMode(e.instance, value)
}

// BorderStyle
func (e *TEdit) BorderStyle() TBorderStyle {
    return Edit_GetBorderStyle(e.instance)
}

// SetBorderStyle
func (e *TEdit) SetBorderStyle(value TBorderStyle) {
    Edit_SetBorderStyle(e.instance, value)
}

// CharCase
func (e *TEdit) CharCase() TEditCharCase {
    return Edit_GetCharCase(e.instance)
}

// SetCharCase
func (e *TEdit) SetCharCase(value TEditCharCase) {
    Edit_SetCharCase(e.instance, value)
}

// Color
func (e *TEdit) Color() TColor {
    return Edit_GetColor(e.instance)
}

// SetColor
func (e *TEdit) SetColor(value TColor) {
    Edit_SetColor(e.instance, value)
}

// DoubleBuffered
func (e *TEdit) DoubleBuffered() bool {
    return Edit_GetDoubleBuffered(e.instance)
}

// SetDoubleBuffered
func (e *TEdit) SetDoubleBuffered(value bool) {
    Edit_SetDoubleBuffered(e.instance, value)
}

// DragCursor
func (e *TEdit) DragCursor() TCursor {
    return Edit_GetDragCursor(e.instance)
}

// SetDragCursor
func (e *TEdit) SetDragCursor(value TCursor) {
    Edit_SetDragCursor(e.instance, value)
}

// DragKind
func (e *TEdit) DragKind() TDragKind {
    return Edit_GetDragKind(e.instance)
}

// SetDragKind
func (e *TEdit) SetDragKind(value TDragKind) {
    Edit_SetDragKind(e.instance, value)
}

// DragMode
func (e *TEdit) DragMode() TDragMode {
    return Edit_GetDragMode(e.instance)
}

// SetDragMode
func (e *TEdit) SetDragMode(value TDragMode) {
    Edit_SetDragMode(e.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (e *TEdit) Enabled() bool {
    return Edit_GetEnabled(e.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (e *TEdit) SetEnabled(value bool) {
    Edit_SetEnabled(e.instance, value)
}

// Font
func (e *TEdit) Font() *TFont {
    return FontFromInst(Edit_GetFont(e.instance))
}

// SetFont
func (e *TEdit) SetFont(value *TFont) {
    Edit_SetFont(e.instance, CheckPtr(value))
}

// HideSelection
func (e *TEdit) HideSelection() bool {
    return Edit_GetHideSelection(e.instance)
}

// SetHideSelection
func (e *TEdit) SetHideSelection(value bool) {
    Edit_SetHideSelection(e.instance, value)
}

// MaxLength
func (e *TEdit) MaxLength() int32 {
    return Edit_GetMaxLength(e.instance)
}

// SetMaxLength
func (e *TEdit) SetMaxLength(value int32) {
    Edit_SetMaxLength(e.instance, value)
}

// NumbersOnly
func (e *TEdit) NumbersOnly() bool {
    return Edit_GetNumbersOnly(e.instance)
}

// SetNumbersOnly
func (e *TEdit) SetNumbersOnly(value bool) {
    Edit_SetNumbersOnly(e.instance, value)
}

// ParentColor
func (e *TEdit) ParentColor() bool {
    return Edit_GetParentColor(e.instance)
}

// SetParentColor
func (e *TEdit) SetParentColor(value bool) {
    Edit_SetParentColor(e.instance, value)
}

// ParentCtl3D
func (e *TEdit) ParentCtl3D() bool {
    return Edit_GetParentCtl3D(e.instance)
}

// SetParentCtl3D
func (e *TEdit) SetParentCtl3D(value bool) {
    Edit_SetParentCtl3D(e.instance, value)
}

// ParentDoubleBuffered
func (e *TEdit) ParentDoubleBuffered() bool {
    return Edit_GetParentDoubleBuffered(e.instance)
}

// SetParentDoubleBuffered
func (e *TEdit) SetParentDoubleBuffered(value bool) {
    Edit_SetParentDoubleBuffered(e.instance, value)
}

// ParentFont
func (e *TEdit) ParentFont() bool {
    return Edit_GetParentFont(e.instance)
}

// SetParentFont
func (e *TEdit) SetParentFont(value bool) {
    Edit_SetParentFont(e.instance, value)
}

// ParentShowHint
func (e *TEdit) ParentShowHint() bool {
    return Edit_GetParentShowHint(e.instance)
}

// SetParentShowHint
func (e *TEdit) SetParentShowHint(value bool) {
    Edit_SetParentShowHint(e.instance, value)
}

// PasswordChar
func (e *TEdit) PasswordChar() uint16 {
    return Edit_GetPasswordChar(e.instance)
}

// SetPasswordChar
func (e *TEdit) SetPasswordChar(value uint16) {
    Edit_SetPasswordChar(e.instance, value)
}

// PopupMenu
func (e *TEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Edit_GetPopupMenu(e.instance))
}

// SetPopupMenu
func (e *TEdit) SetPopupMenu(value IComponent) {
    Edit_SetPopupMenu(e.instance, CheckPtr(value))
}

// ReadOnly
func (e *TEdit) ReadOnly() bool {
    return Edit_GetReadOnly(e.instance)
}

// SetReadOnly
func (e *TEdit) SetReadOnly(value bool) {
    Edit_SetReadOnly(e.instance, value)
}

// ShowHint
func (e *TEdit) ShowHint() bool {
    return Edit_GetShowHint(e.instance)
}

// SetShowHint
func (e *TEdit) SetShowHint(value bool) {
    Edit_SetShowHint(e.instance, value)
}

// TabOrder
func (e *TEdit) TabOrder() uint16 {
    return Edit_GetTabOrder(e.instance)
}

// SetTabOrder
func (e *TEdit) SetTabOrder(value uint16) {
    Edit_SetTabOrder(e.instance, value)
}

// TabStop
func (e *TEdit) TabStop() bool {
    return Edit_GetTabStop(e.instance)
}

// SetTabStop
func (e *TEdit) SetTabStop(value bool) {
    Edit_SetTabStop(e.instance, value)
}

// Text
func (e *TEdit) Text() string {
    return Edit_GetText(e.instance)
}

// SetText
func (e *TEdit) SetText(value string) {
    Edit_SetText(e.instance, value)
}

// TextHint
func (e *TEdit) TextHint() string {
    return Edit_GetTextHint(e.instance)
}

// SetTextHint
func (e *TEdit) SetTextHint(value string) {
    Edit_SetTextHint(e.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (e *TEdit) Visible() bool {
    return Edit_GetVisible(e.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (e *TEdit) SetVisible(value bool) {
    Edit_SetVisible(e.instance, value)
}

// StyleElements
func (e *TEdit) StyleElements() TStyleElements {
    return Edit_GetStyleElements(e.instance)
}

// SetStyleElements
func (e *TEdit) SetStyleElements(value TStyleElements) {
    Edit_SetStyleElements(e.instance, value)
}

// SetOnChange
func (e *TEdit) SetOnChange(fn TNotifyEvent) {
    Edit_SetOnChange(e.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (e *TEdit) SetOnClick(fn TNotifyEvent) {
    Edit_SetOnClick(e.instance, fn)
}

// SetOnContextPopup
func (e *TEdit) SetOnContextPopup(fn TContextPopupEvent) {
    Edit_SetOnContextPopup(e.instance, fn)
}

// SetOnDblClick
func (e *TEdit) SetOnDblClick(fn TNotifyEvent) {
    Edit_SetOnDblClick(e.instance, fn)
}

// SetOnDragDrop
func (e *TEdit) SetOnDragDrop(fn TDragDropEvent) {
    Edit_SetOnDragDrop(e.instance, fn)
}

// SetOnDragOver
func (e *TEdit) SetOnDragOver(fn TDragOverEvent) {
    Edit_SetOnDragOver(e.instance, fn)
}

// SetOnEndDock
func (e *TEdit) SetOnEndDock(fn TEndDragEvent) {
    Edit_SetOnEndDock(e.instance, fn)
}

// SetOnEndDrag
func (e *TEdit) SetOnEndDrag(fn TEndDragEvent) {
    Edit_SetOnEndDrag(e.instance, fn)
}

// SetOnEnter
func (e *TEdit) SetOnEnter(fn TNotifyEvent) {
    Edit_SetOnEnter(e.instance, fn)
}

// SetOnExit
func (e *TEdit) SetOnExit(fn TNotifyEvent) {
    Edit_SetOnExit(e.instance, fn)
}

// SetOnKeyDown
func (e *TEdit) SetOnKeyDown(fn TKeyEvent) {
    Edit_SetOnKeyDown(e.instance, fn)
}

// SetOnKeyPress
func (e *TEdit) SetOnKeyPress(fn TKeyPressEvent) {
    Edit_SetOnKeyPress(e.instance, fn)
}

// SetOnKeyUp
func (e *TEdit) SetOnKeyUp(fn TKeyEvent) {
    Edit_SetOnKeyUp(e.instance, fn)
}

// SetOnMouseDown
func (e *TEdit) SetOnMouseDown(fn TMouseEvent) {
    Edit_SetOnMouseDown(e.instance, fn)
}

// SetOnMouseEnter
func (e *TEdit) SetOnMouseEnter(fn TNotifyEvent) {
    Edit_SetOnMouseEnter(e.instance, fn)
}

// SetOnMouseLeave
func (e *TEdit) SetOnMouseLeave(fn TNotifyEvent) {
    Edit_SetOnMouseLeave(e.instance, fn)
}

// SetOnMouseMove
func (e *TEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    Edit_SetOnMouseMove(e.instance, fn)
}

// SetOnMouseUp
func (e *TEdit) SetOnMouseUp(fn TMouseEvent) {
    Edit_SetOnMouseUp(e.instance, fn)
}

// SetOnStartDock
func (e *TEdit) SetOnStartDock(fn TStartDockEvent) {
    Edit_SetOnStartDock(e.instance, fn)
}

// CanUndo
func (e *TEdit) CanUndo() bool {
    return Edit_GetCanUndo(e.instance)
}

// Modified
func (e *TEdit) Modified() bool {
    return Edit_GetModified(e.instance)
}

// SetModified
func (e *TEdit) SetModified(value bool) {
    Edit_SetModified(e.instance, value)
}

// SelLength
func (e *TEdit) SelLength() int32 {
    return Edit_GetSelLength(e.instance)
}

// SetSelLength
func (e *TEdit) SetSelLength(value int32) {
    Edit_SetSelLength(e.instance, value)
}

// SelStart
func (e *TEdit) SelStart() int32 {
    return Edit_GetSelStart(e.instance)
}

// SetSelStart
func (e *TEdit) SetSelStart(value int32) {
    Edit_SetSelStart(e.instance, value)
}

// SelText
func (e *TEdit) SelText() string {
    return Edit_GetSelText(e.instance)
}

// SetSelText
func (e *TEdit) SetSelText(value string) {
    Edit_SetSelText(e.instance, value)
}

// DockSite
func (e *TEdit) DockSite() bool {
    return Edit_GetDockSite(e.instance)
}

// SetDockSite
func (e *TEdit) SetDockSite(value bool) {
    Edit_SetDockSite(e.instance, value)
}

// Brush
func (e *TEdit) Brush() *TBrush {
    return BrushFromInst(Edit_GetBrush(e.instance))
}

// ControlCount
func (e *TEdit) ControlCount() int32 {
    return Edit_GetControlCount(e.instance)
}

// Handle
func (e *TEdit) Handle() HWND {
    return Edit_GetHandle(e.instance)
}

// ParentWindow
func (e *TEdit) ParentWindow() HWND {
    return Edit_GetParentWindow(e.instance)
}

// SetParentWindow
func (e *TEdit) SetParentWindow(value HWND) {
    Edit_SetParentWindow(e.instance, value)
}

// UseDockManager
func (e *TEdit) UseDockManager() bool {
    return Edit_GetUseDockManager(e.instance)
}

// SetUseDockManager
func (e *TEdit) SetUseDockManager(value bool) {
    Edit_SetUseDockManager(e.instance, value)
}

// Action
func (e *TEdit) Action() *TAction {
    return ActionFromInst(Edit_GetAction(e.instance))
}

// SetAction
func (e *TEdit) SetAction(value IComponent) {
    Edit_SetAction(e.instance, CheckPtr(value))
}

// BoundsRect
func (e *TEdit) BoundsRect() TRect {
    return Edit_GetBoundsRect(e.instance)
}

// SetBoundsRect
func (e *TEdit) SetBoundsRect(value TRect) {
    Edit_SetBoundsRect(e.instance, value)
}

// ClientHeight
func (e *TEdit) ClientHeight() int32 {
    return Edit_GetClientHeight(e.instance)
}

// SetClientHeight
func (e *TEdit) SetClientHeight(value int32) {
    Edit_SetClientHeight(e.instance, value)
}

// ClientRect
func (e *TEdit) ClientRect() TRect {
    return Edit_GetClientRect(e.instance)
}

// ClientWidth
func (e *TEdit) ClientWidth() int32 {
    return Edit_GetClientWidth(e.instance)
}

// SetClientWidth
func (e *TEdit) SetClientWidth(value int32) {
    Edit_SetClientWidth(e.instance, value)
}

// ExplicitLeft
func (e *TEdit) ExplicitLeft() int32 {
    return Edit_GetExplicitLeft(e.instance)
}

// ExplicitTop
func (e *TEdit) ExplicitTop() int32 {
    return Edit_GetExplicitTop(e.instance)
}

// ExplicitWidth
func (e *TEdit) ExplicitWidth() int32 {
    return Edit_GetExplicitWidth(e.instance)
}

// ExplicitHeight
func (e *TEdit) ExplicitHeight() int32 {
    return Edit_GetExplicitHeight(e.instance)
}

// Floating
func (e *TEdit) Floating() bool {
    return Edit_GetFloating(e.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (e *TEdit) Parent() *TWinControl {
    return WinControlFromInst(Edit_GetParent(e.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (e *TEdit) SetParent(value IWinControl) {
    Edit_SetParent(e.instance, CheckPtr(value))
}

// AlignWithMargins
func (e *TEdit) AlignWithMargins() bool {
    return Edit_GetAlignWithMargins(e.instance)
}

// SetAlignWithMargins
func (e *TEdit) SetAlignWithMargins(value bool) {
    Edit_SetAlignWithMargins(e.instance, value)
}

// Left
func (e *TEdit) Left() int32 {
    return Edit_GetLeft(e.instance)
}

// SetLeft
func (e *TEdit) SetLeft(value int32) {
    Edit_SetLeft(e.instance, value)
}

// Top
func (e *TEdit) Top() int32 {
    return Edit_GetTop(e.instance)
}

// SetTop
func (e *TEdit) SetTop(value int32) {
    Edit_SetTop(e.instance, value)
}

// Width
func (e *TEdit) Width() int32 {
    return Edit_GetWidth(e.instance)
}

// SetWidth
func (e *TEdit) SetWidth(value int32) {
    Edit_SetWidth(e.instance, value)
}

// Height
func (e *TEdit) Height() int32 {
    return Edit_GetHeight(e.instance)
}

// SetHeight
func (e *TEdit) SetHeight(value int32) {
    Edit_SetHeight(e.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (e *TEdit) Cursor() TCursor {
    return Edit_GetCursor(e.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (e *TEdit) SetCursor(value TCursor) {
    Edit_SetCursor(e.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (e *TEdit) Hint() string {
    return Edit_GetHint(e.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (e *TEdit) SetHint(value string) {
    Edit_SetHint(e.instance, value)
}

// Margins
func (e *TEdit) Margins() *TMargins {
    return MarginsFromInst(Edit_GetMargins(e.instance))
}

// SetMargins
func (e *TEdit) SetMargins(value *TMargins) {
    Edit_SetMargins(e.instance, CheckPtr(value))
}

// CustomHint
func (e *TEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(Edit_GetCustomHint(e.instance))
}

// SetCustomHint
func (e *TEdit) SetCustomHint(value IComponent) {
    Edit_SetCustomHint(e.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (e *TEdit) ComponentCount() int32 {
    return Edit_GetComponentCount(e.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (e *TEdit) ComponentIndex() int32 {
    return Edit_GetComponentIndex(e.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (e *TEdit) SetComponentIndex(value int32) {
    Edit_SetComponentIndex(e.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (e *TEdit) Owner() *TComponent {
    return ComponentFromInst(Edit_GetOwner(e.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (e *TEdit) Name() string {
    return Edit_GetName(e.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (e *TEdit) SetName(value string) {
    Edit_SetName(e.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (e *TEdit) Tag() int {
    return Edit_GetTag(e.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (e *TEdit) SetTag(value int) {
    Edit_SetTag(e.instance, value)
}

// Controls
func (e *TEdit) Controls(Index int32) *TControl {
    return ControlFromInst(Edit_GetControls(e.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (e *TEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Edit_GetComponents(e.instance, AIndex))
}

