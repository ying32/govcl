
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

type TRichEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewRichEdit
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewRichEdit(owner IComponent) *TRichEdit {
    r := new(TRichEdit)
    r.instance = RichEdit_Create(CheckPtr(owner))
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RichEditFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func RichEditFromInst(inst uintptr) *TRichEdit {
    r := new(TRichEdit)
    r.instance = inst
    r.ptr = unsafe.Pointer(inst)
    return r
}

// RichEditFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func RichEditFromObj(obj IObject) *TRichEdit {
    r := new(TRichEdit)
    r.instance = CheckPtr(obj)
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RichEditFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func RichEditFromUnsafePointer(ptr unsafe.Pointer) *TRichEdit {
    r := new(TRichEdit)
    r.instance = uintptr(ptr)
    r.ptr = ptr
    return r
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (r *TRichEdit) Free() {
    if r.instance != 0 {
        RichEdit_Free(r.instance)
        r.instance = 0
        r.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (r *TRichEdit) Instance() uintptr {
    return r.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (r *TRichEdit) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (r *TRichEdit) IsValid() bool {
    return r.instance != 0
}

// TRichEditClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TRichEditClass() TClass {
    return RichEdit_StaticClassType()
}

// Clear
func (r *TRichEdit) Clear() {
    RichEdit_Clear(r.instance)
}

// FindText
func (r *TRichEdit) FindText(SearchStr string, StartPos int32, Length int32, Options TSearchTypes) int32 {
    return RichEdit_FindText(r.instance, SearchStr , StartPos , Length , Options)
}

// Print
func (r *TRichEdit) Print(Caption string) {
    RichEdit_Print(r.instance, Caption)
}

// GetSelTextBuf
func (r *TRichEdit) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return RichEdit_GetSelTextBuf(r.instance, Buffer , BufSize)
}

// ClearSelection
func (r *TRichEdit) ClearSelection() {
    RichEdit_ClearSelection(r.instance)
}

// CopyToClipboard
func (r *TRichEdit) CopyToClipboard() {
    RichEdit_CopyToClipboard(r.instance)
}

// CutToClipboard
func (r *TRichEdit) CutToClipboard() {
    RichEdit_CutToClipboard(r.instance)
}

// PasteFromClipboard
func (r *TRichEdit) PasteFromClipboard() {
    RichEdit_PasteFromClipboard(r.instance)
}

// Undo
func (r *TRichEdit) Undo() {
    RichEdit_Undo(r.instance)
}

// ClearUndo
func (r *TRichEdit) ClearUndo() {
    RichEdit_ClearUndo(r.instance)
}

// SelectAll
func (r *TRichEdit) SelectAll() {
    RichEdit_SelectAll(r.instance)
}

// CanFocus
func (r *TRichEdit) CanFocus() bool {
    return RichEdit_CanFocus(r.instance)
}

// ContainsControl
func (r *TRichEdit) ContainsControl(Control IControl) bool {
    return RichEdit_ContainsControl(r.instance, CheckPtr(Control))
}

// ControlAtPos
func (r *TRichEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(RichEdit_ControlAtPos(r.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (r *TRichEdit) DisableAlign() {
    RichEdit_DisableAlign(r.instance)
}

// EnableAlign
func (r *TRichEdit) EnableAlign() {
    RichEdit_EnableAlign(r.instance)
}

// FindChildControl
func (r *TRichEdit) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(RichEdit_FindChildControl(r.instance, ControlName))
}

// FlipChildren
func (r *TRichEdit) FlipChildren(AllLevels bool) {
    RichEdit_FlipChildren(r.instance, AllLevels)
}

// Focused
func (r *TRichEdit) Focused() bool {
    return RichEdit_Focused(r.instance)
}

// HandleAllocated
func (r *TRichEdit) HandleAllocated() bool {
    return RichEdit_HandleAllocated(r.instance)
}

// InsertControl
func (r *TRichEdit) InsertControl(AControl IControl) {
    RichEdit_InsertControl(r.instance, CheckPtr(AControl))
}

// Invalidate
func (r *TRichEdit) Invalidate() {
    RichEdit_Invalidate(r.instance)
}

// PaintTo
func (r *TRichEdit) PaintTo(DC HDC, X int32, Y int32) {
    RichEdit_PaintTo(r.instance, DC , X , Y)
}

// RemoveControl
func (r *TRichEdit) RemoveControl(AControl IControl) {
    RichEdit_RemoveControl(r.instance, CheckPtr(AControl))
}

// Realign
func (r *TRichEdit) Realign() {
    RichEdit_Realign(r.instance)
}

// Repaint
func (r *TRichEdit) Repaint() {
    RichEdit_Repaint(r.instance)
}

// ScaleBy
func (r *TRichEdit) ScaleBy(M int32, D int32) {
    RichEdit_ScaleBy(r.instance, M , D)
}

// ScrollBy
func (r *TRichEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    RichEdit_ScrollBy(r.instance, DeltaX , DeltaY)
}

// SetBounds
func (r *TRichEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RichEdit_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (r *TRichEdit) SetFocus() {
    RichEdit_SetFocus(r.instance)
}

// Update
func (r *TRichEdit) Update() {
    RichEdit_Update(r.instance)
}

// UpdateControlState
func (r *TRichEdit) UpdateControlState() {
    RichEdit_UpdateControlState(r.instance)
}

// BringToFront
func (r *TRichEdit) BringToFront() {
    RichEdit_BringToFront(r.instance)
}

// ClientToScreen
func (r *TRichEdit) ClientToScreen(Point TPoint) TPoint {
    return RichEdit_ClientToScreen(r.instance, Point)
}

// ClientToParent
func (r *TRichEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return RichEdit_ClientToParent(r.instance, Point , CheckPtr(AParent))
}

// Dragging
func (r *TRichEdit) Dragging() bool {
    return RichEdit_Dragging(r.instance)
}

// HasParent
func (r *TRichEdit) HasParent() bool {
    return RichEdit_HasParent(r.instance)
}

// Hide
func (r *TRichEdit) Hide() {
    RichEdit_Hide(r.instance)
}

// Perform
func (r *TRichEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RichEdit_Perform(r.instance, Msg , WParam , LParam)
}

// Refresh
func (r *TRichEdit) Refresh() {
    RichEdit_Refresh(r.instance)
}

// ScreenToClient
func (r *TRichEdit) ScreenToClient(Point TPoint) TPoint {
    return RichEdit_ScreenToClient(r.instance, Point)
}

// ParentToClient
func (r *TRichEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return RichEdit_ParentToClient(r.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (r *TRichEdit) SendToBack() {
    RichEdit_SendToBack(r.instance)
}

// Show
func (r *TRichEdit) Show() {
    RichEdit_Show(r.instance)
}

// GetTextBuf
func (r *TRichEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    return RichEdit_GetTextBuf(r.instance, Buffer , BufSize)
}

// GetTextLen
func (r *TRichEdit) GetTextLen() int32 {
    return RichEdit_GetTextLen(r.instance)
}

// SetTextBuf
func (r *TRichEdit) SetTextBuf(Buffer string) {
    RichEdit_SetTextBuf(r.instance, Buffer)
}

// FindComponent
func (r *TRichEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(RichEdit_FindComponent(r.instance, AName))
}

// GetNamePath
func (r *TRichEdit) GetNamePath() string {
    return RichEdit_GetNamePath(r.instance)
}

// Assign
func (r *TRichEdit) Assign(Source IObject) {
    RichEdit_Assign(r.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (r *TRichEdit) DisposeOf() {
    RichEdit_DisposeOf(r.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (r *TRichEdit) ClassType() TClass {
    return RichEdit_ClassType(r.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (r *TRichEdit) ClassName() string {
    return RichEdit_ClassName(r.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (r *TRichEdit) InstanceSize() int32 {
    return RichEdit_InstanceSize(r.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (r *TRichEdit) InheritsFrom(AClass TClass) bool {
    return RichEdit_InheritsFrom(r.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (r *TRichEdit) Equals(Obj IObject) bool {
    return RichEdit_Equals(r.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (r *TRichEdit) GetHashCode() int32 {
    return RichEdit_GetHashCode(r.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (r *TRichEdit) ToString() string {
    return RichEdit_ToString(r.instance)
}

// Align
func (r *TRichEdit) Align() TAlign {
    return RichEdit_GetAlign(r.instance)
}

// SetAlign
func (r *TRichEdit) SetAlign(value TAlign) {
    RichEdit_SetAlign(r.instance, value)
}

// Alignment
func (r *TRichEdit) Alignment() TAlignment {
    return RichEdit_GetAlignment(r.instance)
}

// SetAlignment
func (r *TRichEdit) SetAlignment(value TAlignment) {
    RichEdit_SetAlignment(r.instance, value)
}

// Anchors
func (r *TRichEdit) Anchors() TAnchors {
    return RichEdit_GetAnchors(r.instance)
}

// SetAnchors
func (r *TRichEdit) SetAnchors(value TAnchors) {
    RichEdit_SetAnchors(r.instance, value)
}

// BevelEdges
func (r *TRichEdit) BevelEdges() TBevelEdges {
    return RichEdit_GetBevelEdges(r.instance)
}

// SetBevelEdges
func (r *TRichEdit) SetBevelEdges(value TBevelEdges) {
    RichEdit_SetBevelEdges(r.instance, value)
}

// BevelInner
func (r *TRichEdit) BevelInner() TBevelCut {
    return RichEdit_GetBevelInner(r.instance)
}

// SetBevelInner
func (r *TRichEdit) SetBevelInner(value TBevelCut) {
    RichEdit_SetBevelInner(r.instance, value)
}

// BevelOuter
func (r *TRichEdit) BevelOuter() TBevelCut {
    return RichEdit_GetBevelOuter(r.instance)
}

// SetBevelOuter
func (r *TRichEdit) SetBevelOuter(value TBevelCut) {
    RichEdit_SetBevelOuter(r.instance, value)
}

// BevelKind
func (r *TRichEdit) BevelKind() TBevelKind {
    return RichEdit_GetBevelKind(r.instance)
}

// SetBevelKind
func (r *TRichEdit) SetBevelKind(value TBevelKind) {
    RichEdit_SetBevelKind(r.instance, value)
}

// BiDiMode
func (r *TRichEdit) BiDiMode() TBiDiMode {
    return RichEdit_GetBiDiMode(r.instance)
}

// SetBiDiMode
func (r *TRichEdit) SetBiDiMode(value TBiDiMode) {
    RichEdit_SetBiDiMode(r.instance, value)
}

// BorderStyle
func (r *TRichEdit) BorderStyle() TBorderStyle {
    return RichEdit_GetBorderStyle(r.instance)
}

// SetBorderStyle
func (r *TRichEdit) SetBorderStyle(value TBorderStyle) {
    RichEdit_SetBorderStyle(r.instance, value)
}

// BorderWidth
func (r *TRichEdit) BorderWidth() int32 {
    return RichEdit_GetBorderWidth(r.instance)
}

// SetBorderWidth
func (r *TRichEdit) SetBorderWidth(value int32) {
    RichEdit_SetBorderWidth(r.instance, value)
}

// Color
func (r *TRichEdit) Color() TColor {
    return RichEdit_GetColor(r.instance)
}

// SetColor
func (r *TRichEdit) SetColor(value TColor) {
    RichEdit_SetColor(r.instance, value)
}

// DragCursor
func (r *TRichEdit) DragCursor() TCursor {
    return RichEdit_GetDragCursor(r.instance)
}

// SetDragCursor
func (r *TRichEdit) SetDragCursor(value TCursor) {
    RichEdit_SetDragCursor(r.instance, value)
}

// DragKind
func (r *TRichEdit) DragKind() TDragKind {
    return RichEdit_GetDragKind(r.instance)
}

// SetDragKind
func (r *TRichEdit) SetDragKind(value TDragKind) {
    RichEdit_SetDragKind(r.instance, value)
}

// DragMode
func (r *TRichEdit) DragMode() TDragMode {
    return RichEdit_GetDragMode(r.instance)
}

// SetDragMode
func (r *TRichEdit) SetDragMode(value TDragMode) {
    RichEdit_SetDragMode(r.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (r *TRichEdit) Enabled() bool {
    return RichEdit_GetEnabled(r.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (r *TRichEdit) SetEnabled(value bool) {
    RichEdit_SetEnabled(r.instance, value)
}

// Font
func (r *TRichEdit) Font() *TFont {
    return FontFromInst(RichEdit_GetFont(r.instance))
}

// SetFont
func (r *TRichEdit) SetFont(value *TFont) {
    RichEdit_SetFont(r.instance, CheckPtr(value))
}

// HideSelection
func (r *TRichEdit) HideSelection() bool {
    return RichEdit_GetHideSelection(r.instance)
}

// SetHideSelection
func (r *TRichEdit) SetHideSelection(value bool) {
    RichEdit_SetHideSelection(r.instance, value)
}

// HideScrollBars
func (r *TRichEdit) HideScrollBars() bool {
    return RichEdit_GetHideScrollBars(r.instance)
}

// SetHideScrollBars
func (r *TRichEdit) SetHideScrollBars(value bool) {
    RichEdit_SetHideScrollBars(r.instance, value)
}

// Lines
func (r *TRichEdit) Lines() *TStrings {
    return StringsFromInst(RichEdit_GetLines(r.instance))
}

// SetLines
func (r *TRichEdit) SetLines(value IObject) {
    RichEdit_SetLines(r.instance, CheckPtr(value))
}

// MaxLength
func (r *TRichEdit) MaxLength() int32 {
    return RichEdit_GetMaxLength(r.instance)
}

// SetMaxLength
func (r *TRichEdit) SetMaxLength(value int32) {
    RichEdit_SetMaxLength(r.instance, value)
}

// ParentColor
func (r *TRichEdit) ParentColor() bool {
    return RichEdit_GetParentColor(r.instance)
}

// SetParentColor
func (r *TRichEdit) SetParentColor(value bool) {
    RichEdit_SetParentColor(r.instance, value)
}

// ParentCtl3D
func (r *TRichEdit) ParentCtl3D() bool {
    return RichEdit_GetParentCtl3D(r.instance)
}

// SetParentCtl3D
func (r *TRichEdit) SetParentCtl3D(value bool) {
    RichEdit_SetParentCtl3D(r.instance, value)
}

// ParentFont
func (r *TRichEdit) ParentFont() bool {
    return RichEdit_GetParentFont(r.instance)
}

// SetParentFont
func (r *TRichEdit) SetParentFont(value bool) {
    RichEdit_SetParentFont(r.instance, value)
}

// ParentShowHint
func (r *TRichEdit) ParentShowHint() bool {
    return RichEdit_GetParentShowHint(r.instance)
}

// SetParentShowHint
func (r *TRichEdit) SetParentShowHint(value bool) {
    RichEdit_SetParentShowHint(r.instance, value)
}

// PlainText
func (r *TRichEdit) PlainText() bool {
    return RichEdit_GetPlainText(r.instance)
}

// SetPlainText
func (r *TRichEdit) SetPlainText(value bool) {
    RichEdit_SetPlainText(r.instance, value)
}

// PopupMenu
func (r *TRichEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(RichEdit_GetPopupMenu(r.instance))
}

// SetPopupMenu
func (r *TRichEdit) SetPopupMenu(value IComponent) {
    RichEdit_SetPopupMenu(r.instance, CheckPtr(value))
}

// ReadOnly
func (r *TRichEdit) ReadOnly() bool {
    return RichEdit_GetReadOnly(r.instance)
}

// SetReadOnly
func (r *TRichEdit) SetReadOnly(value bool) {
    RichEdit_SetReadOnly(r.instance, value)
}

// ScrollBars
func (r *TRichEdit) ScrollBars() TScrollStyle {
    return RichEdit_GetScrollBars(r.instance)
}

// SetScrollBars
func (r *TRichEdit) SetScrollBars(value TScrollStyle) {
    RichEdit_SetScrollBars(r.instance, value)
}

// ShowHint
func (r *TRichEdit) ShowHint() bool {
    return RichEdit_GetShowHint(r.instance)
}

// SetShowHint
func (r *TRichEdit) SetShowHint(value bool) {
    RichEdit_SetShowHint(r.instance, value)
}

// TabOrder
func (r *TRichEdit) TabOrder() TTabOrder {
    return RichEdit_GetTabOrder(r.instance)
}

// SetTabOrder
func (r *TRichEdit) SetTabOrder(value TTabOrder) {
    RichEdit_SetTabOrder(r.instance, value)
}

// TabStop
func (r *TRichEdit) TabStop() bool {
    return RichEdit_GetTabStop(r.instance)
}

// SetTabStop
func (r *TRichEdit) SetTabStop(value bool) {
    RichEdit_SetTabStop(r.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (r *TRichEdit) Visible() bool {
    return RichEdit_GetVisible(r.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (r *TRichEdit) SetVisible(value bool) {
    RichEdit_SetVisible(r.instance, value)
}

// WantTabs
func (r *TRichEdit) WantTabs() bool {
    return RichEdit_GetWantTabs(r.instance)
}

// SetWantTabs
func (r *TRichEdit) SetWantTabs(value bool) {
    RichEdit_SetWantTabs(r.instance, value)
}

// WantReturns
func (r *TRichEdit) WantReturns() bool {
    return RichEdit_GetWantReturns(r.instance)
}

// SetWantReturns
func (r *TRichEdit) SetWantReturns(value bool) {
    RichEdit_SetWantReturns(r.instance, value)
}

// WordWrap
func (r *TRichEdit) WordWrap() bool {
    return RichEdit_GetWordWrap(r.instance)
}

// SetWordWrap
func (r *TRichEdit) SetWordWrap(value bool) {
    RichEdit_SetWordWrap(r.instance, value)
}

// StyleElements
func (r *TRichEdit) StyleElements() TStyleElements {
    return RichEdit_GetStyleElements(r.instance)
}

// SetStyleElements
func (r *TRichEdit) SetStyleElements(value TStyleElements) {
    RichEdit_SetStyleElements(r.instance, value)
}

// Zoom
func (r *TRichEdit) Zoom() int32 {
    return RichEdit_GetZoom(r.instance)
}

// SetZoom
func (r *TRichEdit) SetZoom(value int32) {
    RichEdit_SetZoom(r.instance, value)
}

// SetOnChange
func (r *TRichEdit) SetOnChange(fn TNotifyEvent) {
    RichEdit_SetOnChange(r.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (r *TRichEdit) SetOnClick(fn TNotifyEvent) {
    RichEdit_SetOnClick(r.instance, fn)
}

// SetOnContextPopup
func (r *TRichEdit) SetOnContextPopup(fn TContextPopupEvent) {
    RichEdit_SetOnContextPopup(r.instance, fn)
}

// SetOnDblClick
func (r *TRichEdit) SetOnDblClick(fn TNotifyEvent) {
    RichEdit_SetOnDblClick(r.instance, fn)
}

// SetOnDragDrop
func (r *TRichEdit) SetOnDragDrop(fn TDragDropEvent) {
    RichEdit_SetOnDragDrop(r.instance, fn)
}

// SetOnDragOver
func (r *TRichEdit) SetOnDragOver(fn TDragOverEvent) {
    RichEdit_SetOnDragOver(r.instance, fn)
}

// SetOnEndDock
func (r *TRichEdit) SetOnEndDock(fn TEndDragEvent) {
    RichEdit_SetOnEndDock(r.instance, fn)
}

// SetOnEndDrag
func (r *TRichEdit) SetOnEndDrag(fn TEndDragEvent) {
    RichEdit_SetOnEndDrag(r.instance, fn)
}

// SetOnEnter
func (r *TRichEdit) SetOnEnter(fn TNotifyEvent) {
    RichEdit_SetOnEnter(r.instance, fn)
}

// SetOnExit
func (r *TRichEdit) SetOnExit(fn TNotifyEvent) {
    RichEdit_SetOnExit(r.instance, fn)
}

// SetOnKeyDown
func (r *TRichEdit) SetOnKeyDown(fn TKeyEvent) {
    RichEdit_SetOnKeyDown(r.instance, fn)
}

// SetOnKeyPress
func (r *TRichEdit) SetOnKeyPress(fn TKeyPressEvent) {
    RichEdit_SetOnKeyPress(r.instance, fn)
}

// SetOnKeyUp
func (r *TRichEdit) SetOnKeyUp(fn TKeyEvent) {
    RichEdit_SetOnKeyUp(r.instance, fn)
}

// SetOnMouseDown
func (r *TRichEdit) SetOnMouseDown(fn TMouseEvent) {
    RichEdit_SetOnMouseDown(r.instance, fn)
}

// SetOnMouseEnter
func (r *TRichEdit) SetOnMouseEnter(fn TNotifyEvent) {
    RichEdit_SetOnMouseEnter(r.instance, fn)
}

// SetOnMouseLeave
func (r *TRichEdit) SetOnMouseLeave(fn TNotifyEvent) {
    RichEdit_SetOnMouseLeave(r.instance, fn)
}

// SetOnMouseMove
func (r *TRichEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    RichEdit_SetOnMouseMove(r.instance, fn)
}

// SetOnMouseUp
func (r *TRichEdit) SetOnMouseUp(fn TMouseEvent) {
    RichEdit_SetOnMouseUp(r.instance, fn)
}

// SetOnMouseWheel
func (r *TRichEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
    RichEdit_SetOnMouseWheel(r.instance, fn)
}

// SetOnMouseWheelDown
func (r *TRichEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    RichEdit_SetOnMouseWheelDown(r.instance, fn)
}

// SetOnMouseWheelUp
func (r *TRichEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    RichEdit_SetOnMouseWheelUp(r.instance, fn)
}

// SetOnStartDock
func (r *TRichEdit) SetOnStartDock(fn TStartDockEvent) {
    RichEdit_SetOnStartDock(r.instance, fn)
}

// ActiveLineNo
func (r *TRichEdit) ActiveLineNo() uint32 {
    return RichEdit_GetActiveLineNo(r.instance)
}

// DefAttributes
func (r *TRichEdit) DefAttributes() *TTextAttributes {
    return TextAttributesFromInst(RichEdit_GetDefAttributes(r.instance))
}

// SetDefAttributes
func (r *TRichEdit) SetDefAttributes(value *TTextAttributes) {
    RichEdit_SetDefAttributes(r.instance, CheckPtr(value))
}

// SelAttributes
func (r *TRichEdit) SelAttributes() *TTextAttributes {
    return TextAttributesFromInst(RichEdit_GetSelAttributes(r.instance))
}

// SetSelAttributes
func (r *TRichEdit) SetSelAttributes(value *TTextAttributes) {
    RichEdit_SetSelAttributes(r.instance, CheckPtr(value))
}

// PageRect
func (r *TRichEdit) PageRect() TRect {
    return RichEdit_GetPageRect(r.instance)
}

// SetPageRect
func (r *TRichEdit) SetPageRect(value TRect) {
    RichEdit_SetPageRect(r.instance, value)
}

// Paragraph
func (r *TRichEdit) Paragraph() *TParaAttributes {
    return ParaAttributesFromInst(RichEdit_GetParagraph(r.instance))
}

// CaretPos
func (r *TRichEdit) CaretPos() TPoint {
    return RichEdit_GetCaretPos(r.instance)
}

// SetCaretPos
func (r *TRichEdit) SetCaretPos(value TPoint) {
    RichEdit_SetCaretPos(r.instance, value)
}

// CanUndo
func (r *TRichEdit) CanUndo() bool {
    return RichEdit_GetCanUndo(r.instance)
}

// Modified
func (r *TRichEdit) Modified() bool {
    return RichEdit_GetModified(r.instance)
}

// SetModified
func (r *TRichEdit) SetModified(value bool) {
    RichEdit_SetModified(r.instance, value)
}

// SelLength
func (r *TRichEdit) SelLength() int32 {
    return RichEdit_GetSelLength(r.instance)
}

// SetSelLength
func (r *TRichEdit) SetSelLength(value int32) {
    RichEdit_SetSelLength(r.instance, value)
}

// SelStart
func (r *TRichEdit) SelStart() int32 {
    return RichEdit_GetSelStart(r.instance)
}

// SetSelStart
func (r *TRichEdit) SetSelStart(value int32) {
    RichEdit_SetSelStart(r.instance, value)
}

// SelText
func (r *TRichEdit) SelText() string {
    return RichEdit_GetSelText(r.instance)
}

// SetSelText
func (r *TRichEdit) SetSelText(value string) {
    RichEdit_SetSelText(r.instance, value)
}

// Text
func (r *TRichEdit) Text() string {
    return RichEdit_GetText(r.instance)
}

// SetText
func (r *TRichEdit) SetText(value string) {
    RichEdit_SetText(r.instance, value)
}

// TextHint
func (r *TRichEdit) TextHint() string {
    return RichEdit_GetTextHint(r.instance)
}

// SetTextHint
func (r *TRichEdit) SetTextHint(value string) {
    RichEdit_SetTextHint(r.instance, value)
}

// DockClientCount
func (r *TRichEdit) DockClientCount() int32 {
    return RichEdit_GetDockClientCount(r.instance)
}

// DockSite
func (r *TRichEdit) DockSite() bool {
    return RichEdit_GetDockSite(r.instance)
}

// SetDockSite
func (r *TRichEdit) SetDockSite(value bool) {
    RichEdit_SetDockSite(r.instance, value)
}

// DoubleBuffered
func (r *TRichEdit) DoubleBuffered() bool {
    return RichEdit_GetDoubleBuffered(r.instance)
}

// SetDoubleBuffered
func (r *TRichEdit) SetDoubleBuffered(value bool) {
    RichEdit_SetDoubleBuffered(r.instance, value)
}

// AlignDisabled
func (r *TRichEdit) AlignDisabled() bool {
    return RichEdit_GetAlignDisabled(r.instance)
}

// MouseInClient
func (r *TRichEdit) MouseInClient() bool {
    return RichEdit_GetMouseInClient(r.instance)
}

// VisibleDockClientCount
func (r *TRichEdit) VisibleDockClientCount() int32 {
    return RichEdit_GetVisibleDockClientCount(r.instance)
}

// Brush
func (r *TRichEdit) Brush() *TBrush {
    return BrushFromInst(RichEdit_GetBrush(r.instance))
}

// ControlCount
func (r *TRichEdit) ControlCount() int32 {
    return RichEdit_GetControlCount(r.instance)
}

// Handle
func (r *TRichEdit) Handle() HWND {
    return RichEdit_GetHandle(r.instance)
}

// ParentDoubleBuffered
func (r *TRichEdit) ParentDoubleBuffered() bool {
    return RichEdit_GetParentDoubleBuffered(r.instance)
}

// SetParentDoubleBuffered
func (r *TRichEdit) SetParentDoubleBuffered(value bool) {
    RichEdit_SetParentDoubleBuffered(r.instance, value)
}

// ParentWindow
func (r *TRichEdit) ParentWindow() HWND {
    return RichEdit_GetParentWindow(r.instance)
}

// SetParentWindow
func (r *TRichEdit) SetParentWindow(value HWND) {
    RichEdit_SetParentWindow(r.instance, value)
}

// UseDockManager
func (r *TRichEdit) UseDockManager() bool {
    return RichEdit_GetUseDockManager(r.instance)
}

// SetUseDockManager
func (r *TRichEdit) SetUseDockManager(value bool) {
    RichEdit_SetUseDockManager(r.instance, value)
}

// Action
func (r *TRichEdit) Action() *TAction {
    return ActionFromInst(RichEdit_GetAction(r.instance))
}

// SetAction
func (r *TRichEdit) SetAction(value IComponent) {
    RichEdit_SetAction(r.instance, CheckPtr(value))
}

// BoundsRect
func (r *TRichEdit) BoundsRect() TRect {
    return RichEdit_GetBoundsRect(r.instance)
}

// SetBoundsRect
func (r *TRichEdit) SetBoundsRect(value TRect) {
    RichEdit_SetBoundsRect(r.instance, value)
}

// ClientHeight
func (r *TRichEdit) ClientHeight() int32 {
    return RichEdit_GetClientHeight(r.instance)
}

// SetClientHeight
func (r *TRichEdit) SetClientHeight(value int32) {
    RichEdit_SetClientHeight(r.instance, value)
}

// ClientOrigin
func (r *TRichEdit) ClientOrigin() TPoint {
    return RichEdit_GetClientOrigin(r.instance)
}

// ClientRect
func (r *TRichEdit) ClientRect() TRect {
    return RichEdit_GetClientRect(r.instance)
}

// ClientWidth
func (r *TRichEdit) ClientWidth() int32 {
    return RichEdit_GetClientWidth(r.instance)
}

// SetClientWidth
func (r *TRichEdit) SetClientWidth(value int32) {
    RichEdit_SetClientWidth(r.instance, value)
}

// ControlState
func (r *TRichEdit) ControlState() TControlState {
    return RichEdit_GetControlState(r.instance)
}

// SetControlState
func (r *TRichEdit) SetControlState(value TControlState) {
    RichEdit_SetControlState(r.instance, value)
}

// ControlStyle
func (r *TRichEdit) ControlStyle() TControlStyle {
    return RichEdit_GetControlStyle(r.instance)
}

// SetControlStyle
func (r *TRichEdit) SetControlStyle(value TControlStyle) {
    RichEdit_SetControlStyle(r.instance, value)
}

// ExplicitLeft
func (r *TRichEdit) ExplicitLeft() int32 {
    return RichEdit_GetExplicitLeft(r.instance)
}

// ExplicitTop
func (r *TRichEdit) ExplicitTop() int32 {
    return RichEdit_GetExplicitTop(r.instance)
}

// ExplicitWidth
func (r *TRichEdit) ExplicitWidth() int32 {
    return RichEdit_GetExplicitWidth(r.instance)
}

// ExplicitHeight
func (r *TRichEdit) ExplicitHeight() int32 {
    return RichEdit_GetExplicitHeight(r.instance)
}

// Floating
func (r *TRichEdit) Floating() bool {
    return RichEdit_GetFloating(r.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (r *TRichEdit) Parent() *TWinControl {
    return WinControlFromInst(RichEdit_GetParent(r.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (r *TRichEdit) SetParent(value IWinControl) {
    RichEdit_SetParent(r.instance, CheckPtr(value))
}

// AlignWithMargins
func (r *TRichEdit) AlignWithMargins() bool {
    return RichEdit_GetAlignWithMargins(r.instance)
}

// SetAlignWithMargins
func (r *TRichEdit) SetAlignWithMargins(value bool) {
    RichEdit_SetAlignWithMargins(r.instance, value)
}

// Left
func (r *TRichEdit) Left() int32 {
    return RichEdit_GetLeft(r.instance)
}

// SetLeft
func (r *TRichEdit) SetLeft(value int32) {
    RichEdit_SetLeft(r.instance, value)
}

// Top
func (r *TRichEdit) Top() int32 {
    return RichEdit_GetTop(r.instance)
}

// SetTop
func (r *TRichEdit) SetTop(value int32) {
    RichEdit_SetTop(r.instance, value)
}

// Width
func (r *TRichEdit) Width() int32 {
    return RichEdit_GetWidth(r.instance)
}

// SetWidth
func (r *TRichEdit) SetWidth(value int32) {
    RichEdit_SetWidth(r.instance, value)
}

// Height
func (r *TRichEdit) Height() int32 {
    return RichEdit_GetHeight(r.instance)
}

// SetHeight
func (r *TRichEdit) SetHeight(value int32) {
    RichEdit_SetHeight(r.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (r *TRichEdit) Cursor() TCursor {
    return RichEdit_GetCursor(r.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (r *TRichEdit) SetCursor(value TCursor) {
    RichEdit_SetCursor(r.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (r *TRichEdit) Hint() string {
    return RichEdit_GetHint(r.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (r *TRichEdit) SetHint(value string) {
    RichEdit_SetHint(r.instance, value)
}

// Margins
func (r *TRichEdit) Margins() *TMargins {
    return MarginsFromInst(RichEdit_GetMargins(r.instance))
}

// SetMargins
func (r *TRichEdit) SetMargins(value *TMargins) {
    RichEdit_SetMargins(r.instance, CheckPtr(value))
}

// CustomHint
func (r *TRichEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(RichEdit_GetCustomHint(r.instance))
}

// SetCustomHint
func (r *TRichEdit) SetCustomHint(value IComponent) {
    RichEdit_SetCustomHint(r.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (r *TRichEdit) ComponentCount() int32 {
    return RichEdit_GetComponentCount(r.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (r *TRichEdit) ComponentIndex() int32 {
    return RichEdit_GetComponentIndex(r.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (r *TRichEdit) SetComponentIndex(value int32) {
    RichEdit_SetComponentIndex(r.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (r *TRichEdit) Owner() *TComponent {
    return ComponentFromInst(RichEdit_GetOwner(r.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (r *TRichEdit) Name() string {
    return RichEdit_GetName(r.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (r *TRichEdit) SetName(value string) {
    RichEdit_SetName(r.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (r *TRichEdit) Tag() int {
    return RichEdit_GetTag(r.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (r *TRichEdit) SetTag(value int) {
    RichEdit_SetTag(r.instance, value)
}

// DockClients
func (r *TRichEdit) DockClients(Index int32) *TControl {
    return ControlFromInst(RichEdit_GetDockClients(r.instance, Index))
}

// Controls
func (r *TRichEdit) Controls(Index int32) *TControl {
    return ControlFromInst(RichEdit_GetControls(r.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (r *TRichEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(RichEdit_GetComponents(r.instance, AIndex))
}

