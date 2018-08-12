
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

type TLabeledEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewLabeledEdit
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewLabeledEdit(owner IComponent) *TLabeledEdit {
    l := new(TLabeledEdit)
    l.instance = LabeledEdit_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// LabeledEditFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func LabeledEditFromInst(inst uintptr) *TLabeledEdit {
    l := new(TLabeledEdit)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// LabeledEditFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func LabeledEditFromObj(obj IObject) *TLabeledEdit {
    l := new(TLabeledEdit)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// LabeledEditFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func LabeledEditFromUnsafePointer(ptr unsafe.Pointer) *TLabeledEdit {
    l := new(TLabeledEdit)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TLabeledEdit) Free() {
    if l.instance != 0 {
        LabeledEdit_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TLabeledEdit) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TLabeledEdit) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TLabeledEdit) IsValid() bool {
    return l.instance != 0
}

// TLabeledEditClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TLabeledEditClass() TClass {
    return LabeledEdit_StaticClassType()
}

// SetBounds
func (l *TLabeledEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    LabeledEdit_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// SetupInternalLabel
func (l *TLabeledEdit) SetupInternalLabel() {
    LabeledEdit_SetupInternalLabel(l.instance)
}

// Clear
func (l *TLabeledEdit) Clear() {
    LabeledEdit_Clear(l.instance)
}

// ClearSelection
func (l *TLabeledEdit) ClearSelection() {
    LabeledEdit_ClearSelection(l.instance)
}

// CopyToClipboard
func (l *TLabeledEdit) CopyToClipboard() {
    LabeledEdit_CopyToClipboard(l.instance)
}

// CutToClipboard
func (l *TLabeledEdit) CutToClipboard() {
    LabeledEdit_CutToClipboard(l.instance)
}

// PasteFromClipboard
func (l *TLabeledEdit) PasteFromClipboard() {
    LabeledEdit_PasteFromClipboard(l.instance)
}

// Undo
func (l *TLabeledEdit) Undo() {
    LabeledEdit_Undo(l.instance)
}

// ClearUndo
func (l *TLabeledEdit) ClearUndo() {
    LabeledEdit_ClearUndo(l.instance)
}

// SelectAll
func (l *TLabeledEdit) SelectAll() {
    LabeledEdit_SelectAll(l.instance)
}

// GetSelTextBuf
func (l *TLabeledEdit) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return LabeledEdit_GetSelTextBuf(l.instance, Buffer , BufSize)
}

// CanFocus
func (l *TLabeledEdit) CanFocus() bool {
    return LabeledEdit_CanFocus(l.instance)
}

// ContainsControl
func (l *TLabeledEdit) ContainsControl(Control IControl) bool {
    return LabeledEdit_ContainsControl(l.instance, CheckPtr(Control))
}

// ControlAtPos
func (l *TLabeledEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(LabeledEdit_ControlAtPos(l.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (l *TLabeledEdit) DisableAlign() {
    LabeledEdit_DisableAlign(l.instance)
}

// EnableAlign
func (l *TLabeledEdit) EnableAlign() {
    LabeledEdit_EnableAlign(l.instance)
}

// FindChildControl
func (l *TLabeledEdit) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(LabeledEdit_FindChildControl(l.instance, ControlName))
}

// FlipChildren
func (l *TLabeledEdit) FlipChildren(AllLevels bool) {
    LabeledEdit_FlipChildren(l.instance, AllLevels)
}

// Focused
func (l *TLabeledEdit) Focused() bool {
    return LabeledEdit_Focused(l.instance)
}

// HandleAllocated
func (l *TLabeledEdit) HandleAllocated() bool {
    return LabeledEdit_HandleAllocated(l.instance)
}

// InsertControl
func (l *TLabeledEdit) InsertControl(AControl IControl) {
    LabeledEdit_InsertControl(l.instance, CheckPtr(AControl))
}

// Invalidate
func (l *TLabeledEdit) Invalidate() {
    LabeledEdit_Invalidate(l.instance)
}

// PaintTo
func (l *TLabeledEdit) PaintTo(DC HDC, X int32, Y int32) {
    LabeledEdit_PaintTo(l.instance, DC , X , Y)
}

// RemoveControl
func (l *TLabeledEdit) RemoveControl(AControl IControl) {
    LabeledEdit_RemoveControl(l.instance, CheckPtr(AControl))
}

// Realign
func (l *TLabeledEdit) Realign() {
    LabeledEdit_Realign(l.instance)
}

// Repaint
func (l *TLabeledEdit) Repaint() {
    LabeledEdit_Repaint(l.instance)
}

// ScaleBy
func (l *TLabeledEdit) ScaleBy(M int32, D int32) {
    LabeledEdit_ScaleBy(l.instance, M , D)
}

// ScrollBy
func (l *TLabeledEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    LabeledEdit_ScrollBy(l.instance, DeltaX , DeltaY)
}

// SetFocus
func (l *TLabeledEdit) SetFocus() {
    LabeledEdit_SetFocus(l.instance)
}

// Update
func (l *TLabeledEdit) Update() {
    LabeledEdit_Update(l.instance)
}

// UpdateControlState
func (l *TLabeledEdit) UpdateControlState() {
    LabeledEdit_UpdateControlState(l.instance)
}

// BringToFront
func (l *TLabeledEdit) BringToFront() {
    LabeledEdit_BringToFront(l.instance)
}

// ClientToScreen
func (l *TLabeledEdit) ClientToScreen(Point TPoint) TPoint {
    return LabeledEdit_ClientToScreen(l.instance, Point)
}

// ClientToParent
func (l *TLabeledEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return LabeledEdit_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// Dragging
func (l *TLabeledEdit) Dragging() bool {
    return LabeledEdit_Dragging(l.instance)
}

// HasParent
func (l *TLabeledEdit) HasParent() bool {
    return LabeledEdit_HasParent(l.instance)
}

// Hide
func (l *TLabeledEdit) Hide() {
    LabeledEdit_Hide(l.instance)
}

// Perform
func (l *TLabeledEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return LabeledEdit_Perform(l.instance, Msg , WParam , LParam)
}

// Refresh
func (l *TLabeledEdit) Refresh() {
    LabeledEdit_Refresh(l.instance)
}

// ScreenToClient
func (l *TLabeledEdit) ScreenToClient(Point TPoint) TPoint {
    return LabeledEdit_ScreenToClient(l.instance, Point)
}

// ParentToClient
func (l *TLabeledEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return LabeledEdit_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (l *TLabeledEdit) SendToBack() {
    LabeledEdit_SendToBack(l.instance)
}

// Show
func (l *TLabeledEdit) Show() {
    LabeledEdit_Show(l.instance)
}

// GetTextBuf
func (l *TLabeledEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    return LabeledEdit_GetTextBuf(l.instance, Buffer , BufSize)
}

// GetTextLen
func (l *TLabeledEdit) GetTextLen() int32 {
    return LabeledEdit_GetTextLen(l.instance)
}

// SetTextBuf
func (l *TLabeledEdit) SetTextBuf(Buffer string) {
    LabeledEdit_SetTextBuf(l.instance, Buffer)
}

// FindComponent
func (l *TLabeledEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(LabeledEdit_FindComponent(l.instance, AName))
}

// GetNamePath
func (l *TLabeledEdit) GetNamePath() string {
    return LabeledEdit_GetNamePath(l.instance)
}

// Assign
func (l *TLabeledEdit) Assign(Source IObject) {
    LabeledEdit_Assign(l.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TLabeledEdit) DisposeOf() {
    LabeledEdit_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TLabeledEdit) ClassType() TClass {
    return LabeledEdit_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TLabeledEdit) ClassName() string {
    return LabeledEdit_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TLabeledEdit) InstanceSize() int32 {
    return LabeledEdit_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TLabeledEdit) InheritsFrom(AClass TClass) bool {
    return LabeledEdit_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TLabeledEdit) Equals(Obj IObject) bool {
    return LabeledEdit_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TLabeledEdit) GetHashCode() int32 {
    return LabeledEdit_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TLabeledEdit) ToString() string {
    return LabeledEdit_ToString(l.instance)
}

// Alignment
func (l *TLabeledEdit) Alignment() TAlignment {
    return LabeledEdit_GetAlignment(l.instance)
}

// SetAlignment
func (l *TLabeledEdit) SetAlignment(value TAlignment) {
    LabeledEdit_SetAlignment(l.instance, value)
}

// Anchors
func (l *TLabeledEdit) Anchors() TAnchors {
    return LabeledEdit_GetAnchors(l.instance)
}

// SetAnchors
func (l *TLabeledEdit) SetAnchors(value TAnchors) {
    LabeledEdit_SetAnchors(l.instance, value)
}

// AutoSelect
func (l *TLabeledEdit) AutoSelect() bool {
    return LabeledEdit_GetAutoSelect(l.instance)
}

// SetAutoSelect
func (l *TLabeledEdit) SetAutoSelect(value bool) {
    LabeledEdit_SetAutoSelect(l.instance, value)
}

// AutoSize
func (l *TLabeledEdit) AutoSize() bool {
    return LabeledEdit_GetAutoSize(l.instance)
}

// SetAutoSize
func (l *TLabeledEdit) SetAutoSize(value bool) {
    LabeledEdit_SetAutoSize(l.instance, value)
}

// BevelEdges
func (l *TLabeledEdit) BevelEdges() TBevelEdges {
    return LabeledEdit_GetBevelEdges(l.instance)
}

// SetBevelEdges
func (l *TLabeledEdit) SetBevelEdges(value TBevelEdges) {
    LabeledEdit_SetBevelEdges(l.instance, value)
}

// BevelInner
func (l *TLabeledEdit) BevelInner() TBevelCut {
    return LabeledEdit_GetBevelInner(l.instance)
}

// SetBevelInner
func (l *TLabeledEdit) SetBevelInner(value TBevelCut) {
    LabeledEdit_SetBevelInner(l.instance, value)
}

// BevelKind
func (l *TLabeledEdit) BevelKind() TBevelKind {
    return LabeledEdit_GetBevelKind(l.instance)
}

// SetBevelKind
func (l *TLabeledEdit) SetBevelKind(value TBevelKind) {
    LabeledEdit_SetBevelKind(l.instance, value)
}

// BevelOuter
func (l *TLabeledEdit) BevelOuter() TBevelCut {
    return LabeledEdit_GetBevelOuter(l.instance)
}

// SetBevelOuter
func (l *TLabeledEdit) SetBevelOuter(value TBevelCut) {
    LabeledEdit_SetBevelOuter(l.instance, value)
}

// BiDiMode
func (l *TLabeledEdit) BiDiMode() TBiDiMode {
    return LabeledEdit_GetBiDiMode(l.instance)
}

// SetBiDiMode
func (l *TLabeledEdit) SetBiDiMode(value TBiDiMode) {
    LabeledEdit_SetBiDiMode(l.instance, value)
}

// BorderStyle
func (l *TLabeledEdit) BorderStyle() TBorderStyle {
    return LabeledEdit_GetBorderStyle(l.instance)
}

// SetBorderStyle
func (l *TLabeledEdit) SetBorderStyle(value TBorderStyle) {
    LabeledEdit_SetBorderStyle(l.instance, value)
}

// CharCase
func (l *TLabeledEdit) CharCase() TEditCharCase {
    return LabeledEdit_GetCharCase(l.instance)
}

// SetCharCase
func (l *TLabeledEdit) SetCharCase(value TEditCharCase) {
    LabeledEdit_SetCharCase(l.instance, value)
}

// Color
func (l *TLabeledEdit) Color() TColor {
    return LabeledEdit_GetColor(l.instance)
}

// SetColor
func (l *TLabeledEdit) SetColor(value TColor) {
    LabeledEdit_SetColor(l.instance, value)
}

// DoubleBuffered
func (l *TLabeledEdit) DoubleBuffered() bool {
    return LabeledEdit_GetDoubleBuffered(l.instance)
}

// SetDoubleBuffered
func (l *TLabeledEdit) SetDoubleBuffered(value bool) {
    LabeledEdit_SetDoubleBuffered(l.instance, value)
}

// DragCursor
func (l *TLabeledEdit) DragCursor() TCursor {
    return LabeledEdit_GetDragCursor(l.instance)
}

// SetDragCursor
func (l *TLabeledEdit) SetDragCursor(value TCursor) {
    LabeledEdit_SetDragCursor(l.instance, value)
}

// DragKind
func (l *TLabeledEdit) DragKind() TDragKind {
    return LabeledEdit_GetDragKind(l.instance)
}

// SetDragKind
func (l *TLabeledEdit) SetDragKind(value TDragKind) {
    LabeledEdit_SetDragKind(l.instance, value)
}

// DragMode
func (l *TLabeledEdit) DragMode() TDragMode {
    return LabeledEdit_GetDragMode(l.instance)
}

// SetDragMode
func (l *TLabeledEdit) SetDragMode(value TDragMode) {
    LabeledEdit_SetDragMode(l.instance, value)
}

// EditLabel
func (l *TLabeledEdit) EditLabel() *TBoundLabel {
    return BoundLabelFromInst(LabeledEdit_GetEditLabel(l.instance))
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (l *TLabeledEdit) Enabled() bool {
    return LabeledEdit_GetEnabled(l.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (l *TLabeledEdit) SetEnabled(value bool) {
    LabeledEdit_SetEnabled(l.instance, value)
}

// Font
func (l *TLabeledEdit) Font() *TFont {
    return FontFromInst(LabeledEdit_GetFont(l.instance))
}

// SetFont
func (l *TLabeledEdit) SetFont(value *TFont) {
    LabeledEdit_SetFont(l.instance, CheckPtr(value))
}

// HideSelection
func (l *TLabeledEdit) HideSelection() bool {
    return LabeledEdit_GetHideSelection(l.instance)
}

// SetHideSelection
func (l *TLabeledEdit) SetHideSelection(value bool) {
    LabeledEdit_SetHideSelection(l.instance, value)
}

// LabelPosition
func (l *TLabeledEdit) LabelPosition() TLabelPosition {
    return LabeledEdit_GetLabelPosition(l.instance)
}

// SetLabelPosition
func (l *TLabeledEdit) SetLabelPosition(value TLabelPosition) {
    LabeledEdit_SetLabelPosition(l.instance, value)
}

// LabelSpacing
func (l *TLabeledEdit) LabelSpacing() int32 {
    return LabeledEdit_GetLabelSpacing(l.instance)
}

// SetLabelSpacing
func (l *TLabeledEdit) SetLabelSpacing(value int32) {
    LabeledEdit_SetLabelSpacing(l.instance, value)
}

// MaxLength
func (l *TLabeledEdit) MaxLength() int32 {
    return LabeledEdit_GetMaxLength(l.instance)
}

// SetMaxLength
func (l *TLabeledEdit) SetMaxLength(value int32) {
    LabeledEdit_SetMaxLength(l.instance, value)
}

// NumbersOnly
func (l *TLabeledEdit) NumbersOnly() bool {
    return LabeledEdit_GetNumbersOnly(l.instance)
}

// SetNumbersOnly
func (l *TLabeledEdit) SetNumbersOnly(value bool) {
    LabeledEdit_SetNumbersOnly(l.instance, value)
}

// ParentColor
func (l *TLabeledEdit) ParentColor() bool {
    return LabeledEdit_GetParentColor(l.instance)
}

// SetParentColor
func (l *TLabeledEdit) SetParentColor(value bool) {
    LabeledEdit_SetParentColor(l.instance, value)
}

// ParentCtl3D
func (l *TLabeledEdit) ParentCtl3D() bool {
    return LabeledEdit_GetParentCtl3D(l.instance)
}

// SetParentCtl3D
func (l *TLabeledEdit) SetParentCtl3D(value bool) {
    LabeledEdit_SetParentCtl3D(l.instance, value)
}

// ParentDoubleBuffered
func (l *TLabeledEdit) ParentDoubleBuffered() bool {
    return LabeledEdit_GetParentDoubleBuffered(l.instance)
}

// SetParentDoubleBuffered
func (l *TLabeledEdit) SetParentDoubleBuffered(value bool) {
    LabeledEdit_SetParentDoubleBuffered(l.instance, value)
}

// ParentFont
func (l *TLabeledEdit) ParentFont() bool {
    return LabeledEdit_GetParentFont(l.instance)
}

// SetParentFont
func (l *TLabeledEdit) SetParentFont(value bool) {
    LabeledEdit_SetParentFont(l.instance, value)
}

// ParentShowHint
func (l *TLabeledEdit) ParentShowHint() bool {
    return LabeledEdit_GetParentShowHint(l.instance)
}

// SetParentShowHint
func (l *TLabeledEdit) SetParentShowHint(value bool) {
    LabeledEdit_SetParentShowHint(l.instance, value)
}

// PasswordChar
func (l *TLabeledEdit) PasswordChar() uint16 {
    return LabeledEdit_GetPasswordChar(l.instance)
}

// SetPasswordChar
func (l *TLabeledEdit) SetPasswordChar(value uint16) {
    LabeledEdit_SetPasswordChar(l.instance, value)
}

// PopupMenu
func (l *TLabeledEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(LabeledEdit_GetPopupMenu(l.instance))
}

// SetPopupMenu
func (l *TLabeledEdit) SetPopupMenu(value IComponent) {
    LabeledEdit_SetPopupMenu(l.instance, CheckPtr(value))
}

// ReadOnly
func (l *TLabeledEdit) ReadOnly() bool {
    return LabeledEdit_GetReadOnly(l.instance)
}

// SetReadOnly
func (l *TLabeledEdit) SetReadOnly(value bool) {
    LabeledEdit_SetReadOnly(l.instance, value)
}

// ShowHint
func (l *TLabeledEdit) ShowHint() bool {
    return LabeledEdit_GetShowHint(l.instance)
}

// SetShowHint
func (l *TLabeledEdit) SetShowHint(value bool) {
    LabeledEdit_SetShowHint(l.instance, value)
}

// TabOrder
func (l *TLabeledEdit) TabOrder() TTabOrder {
    return LabeledEdit_GetTabOrder(l.instance)
}

// SetTabOrder
func (l *TLabeledEdit) SetTabOrder(value TTabOrder) {
    LabeledEdit_SetTabOrder(l.instance, value)
}

// TabStop
func (l *TLabeledEdit) TabStop() bool {
    return LabeledEdit_GetTabStop(l.instance)
}

// SetTabStop
func (l *TLabeledEdit) SetTabStop(value bool) {
    LabeledEdit_SetTabStop(l.instance, value)
}

// Text
func (l *TLabeledEdit) Text() string {
    return LabeledEdit_GetText(l.instance)
}

// SetText
func (l *TLabeledEdit) SetText(value string) {
    LabeledEdit_SetText(l.instance, value)
}

// TextHint
func (l *TLabeledEdit) TextHint() string {
    return LabeledEdit_GetTextHint(l.instance)
}

// SetTextHint
func (l *TLabeledEdit) SetTextHint(value string) {
    LabeledEdit_SetTextHint(l.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (l *TLabeledEdit) Visible() bool {
    return LabeledEdit_GetVisible(l.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (l *TLabeledEdit) SetVisible(value bool) {
    LabeledEdit_SetVisible(l.instance, value)
}

// StyleElements
func (l *TLabeledEdit) StyleElements() TStyleElements {
    return LabeledEdit_GetStyleElements(l.instance)
}

// SetStyleElements
func (l *TLabeledEdit) SetStyleElements(value TStyleElements) {
    LabeledEdit_SetStyleElements(l.instance, value)
}

// SetOnChange
func (l *TLabeledEdit) SetOnChange(fn TNotifyEvent) {
    LabeledEdit_SetOnChange(l.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (l *TLabeledEdit) SetOnClick(fn TNotifyEvent) {
    LabeledEdit_SetOnClick(l.instance, fn)
}

// SetOnContextPopup
func (l *TLabeledEdit) SetOnContextPopup(fn TContextPopupEvent) {
    LabeledEdit_SetOnContextPopup(l.instance, fn)
}

// SetOnDblClick
func (l *TLabeledEdit) SetOnDblClick(fn TNotifyEvent) {
    LabeledEdit_SetOnDblClick(l.instance, fn)
}

// SetOnDragDrop
func (l *TLabeledEdit) SetOnDragDrop(fn TDragDropEvent) {
    LabeledEdit_SetOnDragDrop(l.instance, fn)
}

// SetOnDragOver
func (l *TLabeledEdit) SetOnDragOver(fn TDragOverEvent) {
    LabeledEdit_SetOnDragOver(l.instance, fn)
}

// SetOnEndDock
func (l *TLabeledEdit) SetOnEndDock(fn TEndDragEvent) {
    LabeledEdit_SetOnEndDock(l.instance, fn)
}

// SetOnEndDrag
func (l *TLabeledEdit) SetOnEndDrag(fn TEndDragEvent) {
    LabeledEdit_SetOnEndDrag(l.instance, fn)
}

// SetOnEnter
func (l *TLabeledEdit) SetOnEnter(fn TNotifyEvent) {
    LabeledEdit_SetOnEnter(l.instance, fn)
}

// SetOnExit
func (l *TLabeledEdit) SetOnExit(fn TNotifyEvent) {
    LabeledEdit_SetOnExit(l.instance, fn)
}

// SetOnKeyDown
func (l *TLabeledEdit) SetOnKeyDown(fn TKeyEvent) {
    LabeledEdit_SetOnKeyDown(l.instance, fn)
}

// SetOnKeyPress
func (l *TLabeledEdit) SetOnKeyPress(fn TKeyPressEvent) {
    LabeledEdit_SetOnKeyPress(l.instance, fn)
}

// SetOnKeyUp
func (l *TLabeledEdit) SetOnKeyUp(fn TKeyEvent) {
    LabeledEdit_SetOnKeyUp(l.instance, fn)
}

// SetOnMouseDown
func (l *TLabeledEdit) SetOnMouseDown(fn TMouseEvent) {
    LabeledEdit_SetOnMouseDown(l.instance, fn)
}

// SetOnMouseEnter
func (l *TLabeledEdit) SetOnMouseEnter(fn TNotifyEvent) {
    LabeledEdit_SetOnMouseEnter(l.instance, fn)
}

// SetOnMouseLeave
func (l *TLabeledEdit) SetOnMouseLeave(fn TNotifyEvent) {
    LabeledEdit_SetOnMouseLeave(l.instance, fn)
}

// SetOnMouseMove
func (l *TLabeledEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    LabeledEdit_SetOnMouseMove(l.instance, fn)
}

// SetOnMouseUp
func (l *TLabeledEdit) SetOnMouseUp(fn TMouseEvent) {
    LabeledEdit_SetOnMouseUp(l.instance, fn)
}

// SetOnStartDock
func (l *TLabeledEdit) SetOnStartDock(fn TStartDockEvent) {
    LabeledEdit_SetOnStartDock(l.instance, fn)
}

// CanUndo
func (l *TLabeledEdit) CanUndo() bool {
    return LabeledEdit_GetCanUndo(l.instance)
}

// Modified
func (l *TLabeledEdit) Modified() bool {
    return LabeledEdit_GetModified(l.instance)
}

// SetModified
func (l *TLabeledEdit) SetModified(value bool) {
    LabeledEdit_SetModified(l.instance, value)
}

// SelLength
func (l *TLabeledEdit) SelLength() int32 {
    return LabeledEdit_GetSelLength(l.instance)
}

// SetSelLength
func (l *TLabeledEdit) SetSelLength(value int32) {
    LabeledEdit_SetSelLength(l.instance, value)
}

// SelStart
func (l *TLabeledEdit) SelStart() int32 {
    return LabeledEdit_GetSelStart(l.instance)
}

// SetSelStart
func (l *TLabeledEdit) SetSelStart(value int32) {
    LabeledEdit_SetSelStart(l.instance, value)
}

// SelText
func (l *TLabeledEdit) SelText() string {
    return LabeledEdit_GetSelText(l.instance)
}

// SetSelText
func (l *TLabeledEdit) SetSelText(value string) {
    LabeledEdit_SetSelText(l.instance, value)
}

// DockClientCount
func (l *TLabeledEdit) DockClientCount() int32 {
    return LabeledEdit_GetDockClientCount(l.instance)
}

// DockSite
func (l *TLabeledEdit) DockSite() bool {
    return LabeledEdit_GetDockSite(l.instance)
}

// SetDockSite
func (l *TLabeledEdit) SetDockSite(value bool) {
    LabeledEdit_SetDockSite(l.instance, value)
}

// AlignDisabled
func (l *TLabeledEdit) AlignDisabled() bool {
    return LabeledEdit_GetAlignDisabled(l.instance)
}

// MouseInClient
func (l *TLabeledEdit) MouseInClient() bool {
    return LabeledEdit_GetMouseInClient(l.instance)
}

// VisibleDockClientCount
func (l *TLabeledEdit) VisibleDockClientCount() int32 {
    return LabeledEdit_GetVisibleDockClientCount(l.instance)
}

// Brush
func (l *TLabeledEdit) Brush() *TBrush {
    return BrushFromInst(LabeledEdit_GetBrush(l.instance))
}

// ControlCount
func (l *TLabeledEdit) ControlCount() int32 {
    return LabeledEdit_GetControlCount(l.instance)
}

// Handle
func (l *TLabeledEdit) Handle() HWND {
    return LabeledEdit_GetHandle(l.instance)
}

// ParentWindow
func (l *TLabeledEdit) ParentWindow() HWND {
    return LabeledEdit_GetParentWindow(l.instance)
}

// SetParentWindow
func (l *TLabeledEdit) SetParentWindow(value HWND) {
    LabeledEdit_SetParentWindow(l.instance, value)
}

// UseDockManager
func (l *TLabeledEdit) UseDockManager() bool {
    return LabeledEdit_GetUseDockManager(l.instance)
}

// SetUseDockManager
func (l *TLabeledEdit) SetUseDockManager(value bool) {
    LabeledEdit_SetUseDockManager(l.instance, value)
}

// Action
func (l *TLabeledEdit) Action() *TAction {
    return ActionFromInst(LabeledEdit_GetAction(l.instance))
}

// SetAction
func (l *TLabeledEdit) SetAction(value IComponent) {
    LabeledEdit_SetAction(l.instance, CheckPtr(value))
}

// Align
func (l *TLabeledEdit) Align() TAlign {
    return LabeledEdit_GetAlign(l.instance)
}

// SetAlign
func (l *TLabeledEdit) SetAlign(value TAlign) {
    LabeledEdit_SetAlign(l.instance, value)
}

// BoundsRect
func (l *TLabeledEdit) BoundsRect() TRect {
    return LabeledEdit_GetBoundsRect(l.instance)
}

// SetBoundsRect
func (l *TLabeledEdit) SetBoundsRect(value TRect) {
    LabeledEdit_SetBoundsRect(l.instance, value)
}

// ClientHeight
func (l *TLabeledEdit) ClientHeight() int32 {
    return LabeledEdit_GetClientHeight(l.instance)
}

// SetClientHeight
func (l *TLabeledEdit) SetClientHeight(value int32) {
    LabeledEdit_SetClientHeight(l.instance, value)
}

// ClientOrigin
func (l *TLabeledEdit) ClientOrigin() TPoint {
    return LabeledEdit_GetClientOrigin(l.instance)
}

// ClientRect
func (l *TLabeledEdit) ClientRect() TRect {
    return LabeledEdit_GetClientRect(l.instance)
}

// ClientWidth
func (l *TLabeledEdit) ClientWidth() int32 {
    return LabeledEdit_GetClientWidth(l.instance)
}

// SetClientWidth
func (l *TLabeledEdit) SetClientWidth(value int32) {
    LabeledEdit_SetClientWidth(l.instance, value)
}

// ControlState
func (l *TLabeledEdit) ControlState() TControlState {
    return LabeledEdit_GetControlState(l.instance)
}

// SetControlState
func (l *TLabeledEdit) SetControlState(value TControlState) {
    LabeledEdit_SetControlState(l.instance, value)
}

// ControlStyle
func (l *TLabeledEdit) ControlStyle() TControlStyle {
    return LabeledEdit_GetControlStyle(l.instance)
}

// SetControlStyle
func (l *TLabeledEdit) SetControlStyle(value TControlStyle) {
    LabeledEdit_SetControlStyle(l.instance, value)
}

// ExplicitLeft
func (l *TLabeledEdit) ExplicitLeft() int32 {
    return LabeledEdit_GetExplicitLeft(l.instance)
}

// ExplicitTop
func (l *TLabeledEdit) ExplicitTop() int32 {
    return LabeledEdit_GetExplicitTop(l.instance)
}

// ExplicitWidth
func (l *TLabeledEdit) ExplicitWidth() int32 {
    return LabeledEdit_GetExplicitWidth(l.instance)
}

// ExplicitHeight
func (l *TLabeledEdit) ExplicitHeight() int32 {
    return LabeledEdit_GetExplicitHeight(l.instance)
}

// Floating
func (l *TLabeledEdit) Floating() bool {
    return LabeledEdit_GetFloating(l.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (l *TLabeledEdit) Parent() *TWinControl {
    return WinControlFromInst(LabeledEdit_GetParent(l.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (l *TLabeledEdit) SetParent(value IWinControl) {
    LabeledEdit_SetParent(l.instance, CheckPtr(value))
}

// AlignWithMargins
func (l *TLabeledEdit) AlignWithMargins() bool {
    return LabeledEdit_GetAlignWithMargins(l.instance)
}

// SetAlignWithMargins
func (l *TLabeledEdit) SetAlignWithMargins(value bool) {
    LabeledEdit_SetAlignWithMargins(l.instance, value)
}

// Left
func (l *TLabeledEdit) Left() int32 {
    return LabeledEdit_GetLeft(l.instance)
}

// SetLeft
func (l *TLabeledEdit) SetLeft(value int32) {
    LabeledEdit_SetLeft(l.instance, value)
}

// Top
func (l *TLabeledEdit) Top() int32 {
    return LabeledEdit_GetTop(l.instance)
}

// SetTop
func (l *TLabeledEdit) SetTop(value int32) {
    LabeledEdit_SetTop(l.instance, value)
}

// Width
func (l *TLabeledEdit) Width() int32 {
    return LabeledEdit_GetWidth(l.instance)
}

// SetWidth
func (l *TLabeledEdit) SetWidth(value int32) {
    LabeledEdit_SetWidth(l.instance, value)
}

// Height
func (l *TLabeledEdit) Height() int32 {
    return LabeledEdit_GetHeight(l.instance)
}

// SetHeight
func (l *TLabeledEdit) SetHeight(value int32) {
    LabeledEdit_SetHeight(l.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (l *TLabeledEdit) Cursor() TCursor {
    return LabeledEdit_GetCursor(l.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (l *TLabeledEdit) SetCursor(value TCursor) {
    LabeledEdit_SetCursor(l.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (l *TLabeledEdit) Hint() string {
    return LabeledEdit_GetHint(l.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (l *TLabeledEdit) SetHint(value string) {
    LabeledEdit_SetHint(l.instance, value)
}

// Margins
func (l *TLabeledEdit) Margins() *TMargins {
    return MarginsFromInst(LabeledEdit_GetMargins(l.instance))
}

// SetMargins
func (l *TLabeledEdit) SetMargins(value *TMargins) {
    LabeledEdit_SetMargins(l.instance, CheckPtr(value))
}

// CustomHint
func (l *TLabeledEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(LabeledEdit_GetCustomHint(l.instance))
}

// SetCustomHint
func (l *TLabeledEdit) SetCustomHint(value IComponent) {
    LabeledEdit_SetCustomHint(l.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (l *TLabeledEdit) ComponentCount() int32 {
    return LabeledEdit_GetComponentCount(l.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (l *TLabeledEdit) ComponentIndex() int32 {
    return LabeledEdit_GetComponentIndex(l.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (l *TLabeledEdit) SetComponentIndex(value int32) {
    LabeledEdit_SetComponentIndex(l.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TLabeledEdit) Owner() *TComponent {
    return ComponentFromInst(LabeledEdit_GetOwner(l.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (l *TLabeledEdit) Name() string {
    return LabeledEdit_GetName(l.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (l *TLabeledEdit) SetName(value string) {
    LabeledEdit_SetName(l.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TLabeledEdit) Tag() int {
    return LabeledEdit_GetTag(l.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TLabeledEdit) SetTag(value int) {
    LabeledEdit_SetTag(l.instance, value)
}

// DockClients
func (l *TLabeledEdit) DockClients(Index int32) *TControl {
    return ControlFromInst(LabeledEdit_GetDockClients(l.instance, Index))
}

// Controls
func (l *TLabeledEdit) Controls(Index int32) *TControl {
    return ControlFromInst(LabeledEdit_GetControls(l.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TLabeledEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(LabeledEdit_GetComponents(l.instance, AIndex))
}

