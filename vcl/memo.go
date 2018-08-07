
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

type TMemo struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewMemo
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMemo(owner IComponent) *TMemo {
    m := new(TMemo)
    m.instance = Memo_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MemoFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MemoFromInst(inst uintptr) *TMemo {
    m := new(TMemo)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

// MemoFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MemoFromObj(obj IObject) *TMemo {
    m := new(TMemo)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MemoFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MemoFromUnsafePointer(ptr unsafe.Pointer) *TMemo {
    m := new(TMemo)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (m *TMemo) Free() {
    if m.instance != 0 {
        Memo_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMemo) Instance() uintptr {
    return m.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMemo) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMemo) IsValid() bool {
    return m.instance != 0
}

// TMemoClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMemoClass() TClass {
    return Memo_StaticClassType()
}

// Clear
func (m *TMemo) Clear() {
    Memo_Clear(m.instance)
}

// ClearSelection
func (m *TMemo) ClearSelection() {
    Memo_ClearSelection(m.instance)
}

// CopyToClipboard
func (m *TMemo) CopyToClipboard() {
    Memo_CopyToClipboard(m.instance)
}

// CutToClipboard
func (m *TMemo) CutToClipboard() {
    Memo_CutToClipboard(m.instance)
}

// PasteFromClipboard
func (m *TMemo) PasteFromClipboard() {
    Memo_PasteFromClipboard(m.instance)
}

// Undo
func (m *TMemo) Undo() {
    Memo_Undo(m.instance)
}

// ClearUndo
func (m *TMemo) ClearUndo() {
    Memo_ClearUndo(m.instance)
}

// SelectAll
func (m *TMemo) SelectAll() {
    Memo_SelectAll(m.instance)
}

// GetSelTextBuf
func (m *TMemo) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return Memo_GetSelTextBuf(m.instance, Buffer , BufSize)
}

// CanFocus
func (m *TMemo) CanFocus() bool {
    return Memo_CanFocus(m.instance)
}

// FlipChildren
func (m *TMemo) FlipChildren(AllLevels bool) {
    Memo_FlipChildren(m.instance, AllLevels)
}

// Focused
func (m *TMemo) Focused() bool {
    return Memo_Focused(m.instance)
}

// HandleAllocated
func (m *TMemo) HandleAllocated() bool {
    return Memo_HandleAllocated(m.instance)
}

// Invalidate
func (m *TMemo) Invalidate() {
    Memo_Invalidate(m.instance)
}

// Realign
func (m *TMemo) Realign() {
    Memo_Realign(m.instance)
}

// Repaint
func (m *TMemo) Repaint() {
    Memo_Repaint(m.instance)
}

// ScaleBy
func (m *TMemo) ScaleBy(M int32, D int32) {
    Memo_ScaleBy(m.instance, M , D)
}

// SetBounds
func (m *TMemo) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Memo_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (m *TMemo) SetFocus() {
    Memo_SetFocus(m.instance)
}

// Update
func (m *TMemo) Update() {
    Memo_Update(m.instance)
}

// BringToFront
func (m *TMemo) BringToFront() {
    Memo_BringToFront(m.instance)
}

// ClientToScreen
func (m *TMemo) ClientToScreen(Point TPoint) TPoint {
    return Memo_ClientToScreen(m.instance, Point)
}

// ClientToParent
func (m *TMemo) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Memo_ClientToParent(m.instance, Point , CheckPtr(AParent))
}

// Dragging
func (m *TMemo) Dragging() bool {
    return Memo_Dragging(m.instance)
}

// HasParent
func (m *TMemo) HasParent() bool {
    return Memo_HasParent(m.instance)
}

// Hide
func (m *TMemo) Hide() {
    Memo_Hide(m.instance)
}

// Perform
func (m *TMemo) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Memo_Perform(m.instance, Msg , WParam , LParam)
}

// Refresh
func (m *TMemo) Refresh() {
    Memo_Refresh(m.instance)
}

// ScreenToClient
func (m *TMemo) ScreenToClient(Point TPoint) TPoint {
    return Memo_ScreenToClient(m.instance, Point)
}

// ParentToClient
func (m *TMemo) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Memo_ParentToClient(m.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (m *TMemo) SendToBack() {
    Memo_SendToBack(m.instance)
}

// Show
func (m *TMemo) Show() {
    Memo_Show(m.instance)
}

// GetTextBuf
func (m *TMemo) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Memo_GetTextBuf(m.instance, Buffer , BufSize)
}

// GetTextLen
func (m *TMemo) GetTextLen() int32 {
    return Memo_GetTextLen(m.instance)
}

// FindComponent
func (m *TMemo) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Memo_FindComponent(m.instance, AName))
}

// GetNamePath
func (m *TMemo) GetNamePath() string {
    return Memo_GetNamePath(m.instance)
}

// Assign
func (m *TMemo) Assign(Source IObject) {
    Memo_Assign(m.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMemo) DisposeOf() {
    Memo_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMemo) ClassType() TClass {
    return Memo_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMemo) ClassName() string {
    return Memo_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMemo) InstanceSize() int32 {
    return Memo_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMemo) InheritsFrom(AClass TClass) bool {
    return Memo_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMemo) Equals(Obj IObject) bool {
    return Memo_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMemo) GetHashCode() int32 {
    return Memo_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMemo) ToString() string {
    return Memo_ToString(m.instance)
}

// Align
func (m *TMemo) Align() TAlign {
    return Memo_GetAlign(m.instance)
}

// SetAlign
func (m *TMemo) SetAlign(value TAlign) {
    Memo_SetAlign(m.instance, value)
}

// Alignment
func (m *TMemo) Alignment() TAlignment {
    return Memo_GetAlignment(m.instance)
}

// SetAlignment
func (m *TMemo) SetAlignment(value TAlignment) {
    Memo_SetAlignment(m.instance, value)
}

// Anchors
func (m *TMemo) Anchors() TAnchors {
    return Memo_GetAnchors(m.instance)
}

// SetAnchors
func (m *TMemo) SetAnchors(value TAnchors) {
    Memo_SetAnchors(m.instance, value)
}

// BevelEdges
func (m *TMemo) BevelEdges() TBevelEdges {
    return Memo_GetBevelEdges(m.instance)
}

// SetBevelEdges
func (m *TMemo) SetBevelEdges(value TBevelEdges) {
    Memo_SetBevelEdges(m.instance, value)
}

// BevelInner
func (m *TMemo) BevelInner() TBevelCut {
    return Memo_GetBevelInner(m.instance)
}

// SetBevelInner
func (m *TMemo) SetBevelInner(value TBevelCut) {
    Memo_SetBevelInner(m.instance, value)
}

// BevelKind
func (m *TMemo) BevelKind() TBevelKind {
    return Memo_GetBevelKind(m.instance)
}

// SetBevelKind
func (m *TMemo) SetBevelKind(value TBevelKind) {
    Memo_SetBevelKind(m.instance, value)
}

// BevelOuter
func (m *TMemo) BevelOuter() TBevelCut {
    return Memo_GetBevelOuter(m.instance)
}

// SetBevelOuter
func (m *TMemo) SetBevelOuter(value TBevelCut) {
    Memo_SetBevelOuter(m.instance, value)
}

// BiDiMode
func (m *TMemo) BiDiMode() TBiDiMode {
    return Memo_GetBiDiMode(m.instance)
}

// SetBiDiMode
func (m *TMemo) SetBiDiMode(value TBiDiMode) {
    Memo_SetBiDiMode(m.instance, value)
}

// BorderStyle
func (m *TMemo) BorderStyle() TBorderStyle {
    return Memo_GetBorderStyle(m.instance)
}

// SetBorderStyle
func (m *TMemo) SetBorderStyle(value TBorderStyle) {
    Memo_SetBorderStyle(m.instance, value)
}

// CharCase
func (m *TMemo) CharCase() TEditCharCase {
    return Memo_GetCharCase(m.instance)
}

// SetCharCase
func (m *TMemo) SetCharCase(value TEditCharCase) {
    Memo_SetCharCase(m.instance, value)
}

// Color
func (m *TMemo) Color() TColor {
    return Memo_GetColor(m.instance)
}

// SetColor
func (m *TMemo) SetColor(value TColor) {
    Memo_SetColor(m.instance, value)
}

// DoubleBuffered
func (m *TMemo) DoubleBuffered() bool {
    return Memo_GetDoubleBuffered(m.instance)
}

// SetDoubleBuffered
func (m *TMemo) SetDoubleBuffered(value bool) {
    Memo_SetDoubleBuffered(m.instance, value)
}

// DragCursor
func (m *TMemo) DragCursor() TCursor {
    return Memo_GetDragCursor(m.instance)
}

// SetDragCursor
func (m *TMemo) SetDragCursor(value TCursor) {
    Memo_SetDragCursor(m.instance, value)
}

// DragKind
func (m *TMemo) DragKind() TDragKind {
    return Memo_GetDragKind(m.instance)
}

// SetDragKind
func (m *TMemo) SetDragKind(value TDragKind) {
    Memo_SetDragKind(m.instance, value)
}

// DragMode
func (m *TMemo) DragMode() TDragMode {
    return Memo_GetDragMode(m.instance)
}

// SetDragMode
func (m *TMemo) SetDragMode(value TDragMode) {
    Memo_SetDragMode(m.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (m *TMemo) Enabled() bool {
    return Memo_GetEnabled(m.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (m *TMemo) SetEnabled(value bool) {
    Memo_SetEnabled(m.instance, value)
}

// Font
func (m *TMemo) Font() *TFont {
    return FontFromInst(Memo_GetFont(m.instance))
}

// SetFont
func (m *TMemo) SetFont(value *TFont) {
    Memo_SetFont(m.instance, CheckPtr(value))
}

// HideSelection
func (m *TMemo) HideSelection() bool {
    return Memo_GetHideSelection(m.instance)
}

// SetHideSelection
func (m *TMemo) SetHideSelection(value bool) {
    Memo_SetHideSelection(m.instance, value)
}

// Lines
func (m *TMemo) Lines() *TStrings {
    return StringsFromInst(Memo_GetLines(m.instance))
}

// SetLines
func (m *TMemo) SetLines(value IObject) {
    Memo_SetLines(m.instance, CheckPtr(value))
}

// MaxLength
func (m *TMemo) MaxLength() int32 {
    return Memo_GetMaxLength(m.instance)
}

// SetMaxLength
func (m *TMemo) SetMaxLength(value int32) {
    Memo_SetMaxLength(m.instance, value)
}

// ParentColor
func (m *TMemo) ParentColor() bool {
    return Memo_GetParentColor(m.instance)
}

// SetParentColor
func (m *TMemo) SetParentColor(value bool) {
    Memo_SetParentColor(m.instance, value)
}

// ParentCtl3D
func (m *TMemo) ParentCtl3D() bool {
    return Memo_GetParentCtl3D(m.instance)
}

// SetParentCtl3D
func (m *TMemo) SetParentCtl3D(value bool) {
    Memo_SetParentCtl3D(m.instance, value)
}

// ParentDoubleBuffered
func (m *TMemo) ParentDoubleBuffered() bool {
    return Memo_GetParentDoubleBuffered(m.instance)
}

// SetParentDoubleBuffered
func (m *TMemo) SetParentDoubleBuffered(value bool) {
    Memo_SetParentDoubleBuffered(m.instance, value)
}

// ParentFont
func (m *TMemo) ParentFont() bool {
    return Memo_GetParentFont(m.instance)
}

// SetParentFont
func (m *TMemo) SetParentFont(value bool) {
    Memo_SetParentFont(m.instance, value)
}

// ParentShowHint
func (m *TMemo) ParentShowHint() bool {
    return Memo_GetParentShowHint(m.instance)
}

// SetParentShowHint
func (m *TMemo) SetParentShowHint(value bool) {
    Memo_SetParentShowHint(m.instance, value)
}

// PopupMenu
func (m *TMemo) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Memo_GetPopupMenu(m.instance))
}

// SetPopupMenu
func (m *TMemo) SetPopupMenu(value IComponent) {
    Memo_SetPopupMenu(m.instance, CheckPtr(value))
}

// ReadOnly
func (m *TMemo) ReadOnly() bool {
    return Memo_GetReadOnly(m.instance)
}

// SetReadOnly
func (m *TMemo) SetReadOnly(value bool) {
    Memo_SetReadOnly(m.instance, value)
}

// ScrollBars
func (m *TMemo) ScrollBars() TScrollStyle {
    return Memo_GetScrollBars(m.instance)
}

// SetScrollBars
func (m *TMemo) SetScrollBars(value TScrollStyle) {
    Memo_SetScrollBars(m.instance, value)
}

// ShowHint
func (m *TMemo) ShowHint() bool {
    return Memo_GetShowHint(m.instance)
}

// SetShowHint
func (m *TMemo) SetShowHint(value bool) {
    Memo_SetShowHint(m.instance, value)
}

// TabOrder
func (m *TMemo) TabOrder() uint16 {
    return Memo_GetTabOrder(m.instance)
}

// SetTabOrder
func (m *TMemo) SetTabOrder(value uint16) {
    Memo_SetTabOrder(m.instance, value)
}

// TabStop
func (m *TMemo) TabStop() bool {
    return Memo_GetTabStop(m.instance)
}

// SetTabStop
func (m *TMemo) SetTabStop(value bool) {
    Memo_SetTabStop(m.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (m *TMemo) Visible() bool {
    return Memo_GetVisible(m.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (m *TMemo) SetVisible(value bool) {
    Memo_SetVisible(m.instance, value)
}

// WantReturns
func (m *TMemo) WantReturns() bool {
    return Memo_GetWantReturns(m.instance)
}

// SetWantReturns
func (m *TMemo) SetWantReturns(value bool) {
    Memo_SetWantReturns(m.instance, value)
}

// WantTabs
func (m *TMemo) WantTabs() bool {
    return Memo_GetWantTabs(m.instance)
}

// SetWantTabs
func (m *TMemo) SetWantTabs(value bool) {
    Memo_SetWantTabs(m.instance, value)
}

// WordWrap
func (m *TMemo) WordWrap() bool {
    return Memo_GetWordWrap(m.instance)
}

// SetWordWrap
func (m *TMemo) SetWordWrap(value bool) {
    Memo_SetWordWrap(m.instance, value)
}

// StyleElements
func (m *TMemo) StyleElements() TStyleElements {
    return Memo_GetStyleElements(m.instance)
}

// SetStyleElements
func (m *TMemo) SetStyleElements(value TStyleElements) {
    Memo_SetStyleElements(m.instance, value)
}

// SetOnChange
func (m *TMemo) SetOnChange(fn TNotifyEvent) {
    Memo_SetOnChange(m.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (m *TMemo) SetOnClick(fn TNotifyEvent) {
    Memo_SetOnClick(m.instance, fn)
}

// SetOnContextPopup
func (m *TMemo) SetOnContextPopup(fn TContextPopupEvent) {
    Memo_SetOnContextPopup(m.instance, fn)
}

// SetOnDblClick
func (m *TMemo) SetOnDblClick(fn TNotifyEvent) {
    Memo_SetOnDblClick(m.instance, fn)
}

// SetOnDragDrop
func (m *TMemo) SetOnDragDrop(fn TDragDropEvent) {
    Memo_SetOnDragDrop(m.instance, fn)
}

// SetOnDragOver
func (m *TMemo) SetOnDragOver(fn TDragOverEvent) {
    Memo_SetOnDragOver(m.instance, fn)
}

// SetOnEndDock
func (m *TMemo) SetOnEndDock(fn TEndDragEvent) {
    Memo_SetOnEndDock(m.instance, fn)
}

// SetOnEndDrag
func (m *TMemo) SetOnEndDrag(fn TEndDragEvent) {
    Memo_SetOnEndDrag(m.instance, fn)
}

// SetOnEnter
func (m *TMemo) SetOnEnter(fn TNotifyEvent) {
    Memo_SetOnEnter(m.instance, fn)
}

// SetOnExit
func (m *TMemo) SetOnExit(fn TNotifyEvent) {
    Memo_SetOnExit(m.instance, fn)
}

// SetOnKeyDown
func (m *TMemo) SetOnKeyDown(fn TKeyEvent) {
    Memo_SetOnKeyDown(m.instance, fn)
}

// SetOnKeyPress
func (m *TMemo) SetOnKeyPress(fn TKeyPressEvent) {
    Memo_SetOnKeyPress(m.instance, fn)
}

// SetOnKeyUp
func (m *TMemo) SetOnKeyUp(fn TKeyEvent) {
    Memo_SetOnKeyUp(m.instance, fn)
}

// SetOnMouseDown
func (m *TMemo) SetOnMouseDown(fn TMouseEvent) {
    Memo_SetOnMouseDown(m.instance, fn)
}

// SetOnMouseEnter
func (m *TMemo) SetOnMouseEnter(fn TNotifyEvent) {
    Memo_SetOnMouseEnter(m.instance, fn)
}

// SetOnMouseLeave
func (m *TMemo) SetOnMouseLeave(fn TNotifyEvent) {
    Memo_SetOnMouseLeave(m.instance, fn)
}

// SetOnMouseMove
func (m *TMemo) SetOnMouseMove(fn TMouseMoveEvent) {
    Memo_SetOnMouseMove(m.instance, fn)
}

// SetOnMouseUp
func (m *TMemo) SetOnMouseUp(fn TMouseEvent) {
    Memo_SetOnMouseUp(m.instance, fn)
}

// SetOnStartDock
func (m *TMemo) SetOnStartDock(fn TStartDockEvent) {
    Memo_SetOnStartDock(m.instance, fn)
}

// CaretPos
func (m *TMemo) CaretPos() TPoint {
    return Memo_GetCaretPos(m.instance)
}

// SetCaretPos
func (m *TMemo) SetCaretPos(value TPoint) {
    Memo_SetCaretPos(m.instance, value)
}

// CanUndo
func (m *TMemo) CanUndo() bool {
    return Memo_GetCanUndo(m.instance)
}

// Modified
func (m *TMemo) Modified() bool {
    return Memo_GetModified(m.instance)
}

// SetModified
func (m *TMemo) SetModified(value bool) {
    Memo_SetModified(m.instance, value)
}

// SelLength
func (m *TMemo) SelLength() int32 {
    return Memo_GetSelLength(m.instance)
}

// SetSelLength
func (m *TMemo) SetSelLength(value int32) {
    Memo_SetSelLength(m.instance, value)
}

// SelStart
func (m *TMemo) SelStart() int32 {
    return Memo_GetSelStart(m.instance)
}

// SetSelStart
func (m *TMemo) SetSelStart(value int32) {
    Memo_SetSelStart(m.instance, value)
}

// SelText
func (m *TMemo) SelText() string {
    return Memo_GetSelText(m.instance)
}

// SetSelText
func (m *TMemo) SetSelText(value string) {
    Memo_SetSelText(m.instance, value)
}

// Text
func (m *TMemo) Text() string {
    return Memo_GetText(m.instance)
}

// SetText
func (m *TMemo) SetText(value string) {
    Memo_SetText(m.instance, value)
}

// TextHint
func (m *TMemo) TextHint() string {
    return Memo_GetTextHint(m.instance)
}

// SetTextHint
func (m *TMemo) SetTextHint(value string) {
    Memo_SetTextHint(m.instance, value)
}

// DockSite
func (m *TMemo) DockSite() bool {
    return Memo_GetDockSite(m.instance)
}

// SetDockSite
func (m *TMemo) SetDockSite(value bool) {
    Memo_SetDockSite(m.instance, value)
}

// Brush
func (m *TMemo) Brush() *TBrush {
    return BrushFromInst(Memo_GetBrush(m.instance))
}

// ControlCount
func (m *TMemo) ControlCount() int32 {
    return Memo_GetControlCount(m.instance)
}

// Handle
func (m *TMemo) Handle() HWND {
    return Memo_GetHandle(m.instance)
}

// ParentWindow
func (m *TMemo) ParentWindow() HWND {
    return Memo_GetParentWindow(m.instance)
}

// SetParentWindow
func (m *TMemo) SetParentWindow(value HWND) {
    Memo_SetParentWindow(m.instance, value)
}

// UseDockManager
func (m *TMemo) UseDockManager() bool {
    return Memo_GetUseDockManager(m.instance)
}

// SetUseDockManager
func (m *TMemo) SetUseDockManager(value bool) {
    Memo_SetUseDockManager(m.instance, value)
}

// Action
func (m *TMemo) Action() *TAction {
    return ActionFromInst(Memo_GetAction(m.instance))
}

// SetAction
func (m *TMemo) SetAction(value IComponent) {
    Memo_SetAction(m.instance, CheckPtr(value))
}

// BoundsRect
func (m *TMemo) BoundsRect() TRect {
    return Memo_GetBoundsRect(m.instance)
}

// SetBoundsRect
func (m *TMemo) SetBoundsRect(value TRect) {
    Memo_SetBoundsRect(m.instance, value)
}

// ClientHeight
func (m *TMemo) ClientHeight() int32 {
    return Memo_GetClientHeight(m.instance)
}

// SetClientHeight
func (m *TMemo) SetClientHeight(value int32) {
    Memo_SetClientHeight(m.instance, value)
}

// ClientRect
func (m *TMemo) ClientRect() TRect {
    return Memo_GetClientRect(m.instance)
}

// ClientWidth
func (m *TMemo) ClientWidth() int32 {
    return Memo_GetClientWidth(m.instance)
}

// SetClientWidth
func (m *TMemo) SetClientWidth(value int32) {
    Memo_SetClientWidth(m.instance, value)
}

// ExplicitLeft
func (m *TMemo) ExplicitLeft() int32 {
    return Memo_GetExplicitLeft(m.instance)
}

// ExplicitTop
func (m *TMemo) ExplicitTop() int32 {
    return Memo_GetExplicitTop(m.instance)
}

// ExplicitWidth
func (m *TMemo) ExplicitWidth() int32 {
    return Memo_GetExplicitWidth(m.instance)
}

// ExplicitHeight
func (m *TMemo) ExplicitHeight() int32 {
    return Memo_GetExplicitHeight(m.instance)
}

// Floating
func (m *TMemo) Floating() bool {
    return Memo_GetFloating(m.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (m *TMemo) Parent() *TWinControl {
    return WinControlFromInst(Memo_GetParent(m.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (m *TMemo) SetParent(value IWinControl) {
    Memo_SetParent(m.instance, CheckPtr(value))
}

// AlignWithMargins
func (m *TMemo) AlignWithMargins() bool {
    return Memo_GetAlignWithMargins(m.instance)
}

// SetAlignWithMargins
func (m *TMemo) SetAlignWithMargins(value bool) {
    Memo_SetAlignWithMargins(m.instance, value)
}

// Left
func (m *TMemo) Left() int32 {
    return Memo_GetLeft(m.instance)
}

// SetLeft
func (m *TMemo) SetLeft(value int32) {
    Memo_SetLeft(m.instance, value)
}

// Top
func (m *TMemo) Top() int32 {
    return Memo_GetTop(m.instance)
}

// SetTop
func (m *TMemo) SetTop(value int32) {
    Memo_SetTop(m.instance, value)
}

// Width
func (m *TMemo) Width() int32 {
    return Memo_GetWidth(m.instance)
}

// SetWidth
func (m *TMemo) SetWidth(value int32) {
    Memo_SetWidth(m.instance, value)
}

// Height
func (m *TMemo) Height() int32 {
    return Memo_GetHeight(m.instance)
}

// SetHeight
func (m *TMemo) SetHeight(value int32) {
    Memo_SetHeight(m.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (m *TMemo) Cursor() TCursor {
    return Memo_GetCursor(m.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (m *TMemo) SetCursor(value TCursor) {
    Memo_SetCursor(m.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (m *TMemo) Hint() string {
    return Memo_GetHint(m.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (m *TMemo) SetHint(value string) {
    Memo_SetHint(m.instance, value)
}

// Margins
func (m *TMemo) Margins() *TMargins {
    return MarginsFromInst(Memo_GetMargins(m.instance))
}

// SetMargins
func (m *TMemo) SetMargins(value *TMargins) {
    Memo_SetMargins(m.instance, CheckPtr(value))
}

// CustomHint
func (m *TMemo) CustomHint() *TCustomHint {
    return CustomHintFromInst(Memo_GetCustomHint(m.instance))
}

// SetCustomHint
func (m *TMemo) SetCustomHint(value IComponent) {
    Memo_SetCustomHint(m.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (m *TMemo) ComponentCount() int32 {
    return Memo_GetComponentCount(m.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (m *TMemo) ComponentIndex() int32 {
    return Memo_GetComponentIndex(m.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (m *TMemo) SetComponentIndex(value int32) {
    Memo_SetComponentIndex(m.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (m *TMemo) Owner() *TComponent {
    return ComponentFromInst(Memo_GetOwner(m.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (m *TMemo) Name() string {
    return Memo_GetName(m.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (m *TMemo) SetName(value string) {
    Memo_SetName(m.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (m *TMemo) Tag() int {
    return Memo_GetTag(m.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (m *TMemo) SetTag(value int) {
    Memo_SetTag(m.instance, value)
}

// Controls
func (m *TMemo) Controls(Index int32) *TControl {
    return ControlFromInst(Memo_GetControls(m.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (m *TMemo) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Memo_GetComponents(m.instance, AIndex))
}

