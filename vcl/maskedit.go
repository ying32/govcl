
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

type TMaskEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewMaskEdit
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMaskEdit(owner IComponent) *TMaskEdit {
    m := new(TMaskEdit)
    m.instance = MaskEdit_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MaskEditFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MaskEditFromInst(inst uintptr) *TMaskEdit {
    m := new(TMaskEdit)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

// MaskEditFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MaskEditFromObj(obj IObject) *TMaskEdit {
    m := new(TMaskEdit)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MaskEditFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MaskEditFromUnsafePointer(ptr unsafe.Pointer) *TMaskEdit {
    m := new(TMaskEdit)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (m *TMaskEdit) Free() {
    if m.instance != 0 {
        MaskEdit_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMaskEdit) Instance() uintptr {
    return m.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMaskEdit) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMaskEdit) IsValid() bool {
    return m.instance != 0
}

// TMaskEditClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMaskEditClass() TClass {
    return MaskEdit_StaticClassType()
}

// ValidateEdit
func (m *TMaskEdit) ValidateEdit() {
    MaskEdit_ValidateEdit(m.instance)
}

// Clear
func (m *TMaskEdit) Clear() {
    MaskEdit_Clear(m.instance)
}

// GetTextLen
func (m *TMaskEdit) GetTextLen() int32 {
    return MaskEdit_GetTextLen(m.instance)
}

// ClearSelection
func (m *TMaskEdit) ClearSelection() {
    MaskEdit_ClearSelection(m.instance)
}

// CopyToClipboard
func (m *TMaskEdit) CopyToClipboard() {
    MaskEdit_CopyToClipboard(m.instance)
}

// CutToClipboard
func (m *TMaskEdit) CutToClipboard() {
    MaskEdit_CutToClipboard(m.instance)
}

// PasteFromClipboard
func (m *TMaskEdit) PasteFromClipboard() {
    MaskEdit_PasteFromClipboard(m.instance)
}

// Undo
func (m *TMaskEdit) Undo() {
    MaskEdit_Undo(m.instance)
}

// ClearUndo
func (m *TMaskEdit) ClearUndo() {
    MaskEdit_ClearUndo(m.instance)
}

// SelectAll
func (m *TMaskEdit) SelectAll() {
    MaskEdit_SelectAll(m.instance)
}

// GetSelTextBuf
func (m *TMaskEdit) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return MaskEdit_GetSelTextBuf(m.instance, Buffer , BufSize)
}

// CanFocus
func (m *TMaskEdit) CanFocus() bool {
    return MaskEdit_CanFocus(m.instance)
}

// ContainsControl
func (m *TMaskEdit) ContainsControl(Control IControl) bool {
    return MaskEdit_ContainsControl(m.instance, CheckPtr(Control))
}

// ControlAtPos
func (m *TMaskEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(MaskEdit_ControlAtPos(m.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (m *TMaskEdit) DisableAlign() {
    MaskEdit_DisableAlign(m.instance)
}

// EnableAlign
func (m *TMaskEdit) EnableAlign() {
    MaskEdit_EnableAlign(m.instance)
}

// FindChildControl
func (m *TMaskEdit) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(MaskEdit_FindChildControl(m.instance, ControlName))
}

// FlipChildren
func (m *TMaskEdit) FlipChildren(AllLevels bool) {
    MaskEdit_FlipChildren(m.instance, AllLevels)
}

// Focused
func (m *TMaskEdit) Focused() bool {
    return MaskEdit_Focused(m.instance)
}

// HandleAllocated
func (m *TMaskEdit) HandleAllocated() bool {
    return MaskEdit_HandleAllocated(m.instance)
}

// InsertControl
func (m *TMaskEdit) InsertControl(AControl IControl) {
    MaskEdit_InsertControl(m.instance, CheckPtr(AControl))
}

// Invalidate
func (m *TMaskEdit) Invalidate() {
    MaskEdit_Invalidate(m.instance)
}

// PaintTo
func (m *TMaskEdit) PaintTo(DC HDC, X int32, Y int32) {
    MaskEdit_PaintTo(m.instance, DC , X , Y)
}

// RemoveControl
func (m *TMaskEdit) RemoveControl(AControl IControl) {
    MaskEdit_RemoveControl(m.instance, CheckPtr(AControl))
}

// Realign
func (m *TMaskEdit) Realign() {
    MaskEdit_Realign(m.instance)
}

// Repaint
func (m *TMaskEdit) Repaint() {
    MaskEdit_Repaint(m.instance)
}

// ScaleBy
func (m *TMaskEdit) ScaleBy(M int32, D int32) {
    MaskEdit_ScaleBy(m.instance, M , D)
}

// ScrollBy
func (m *TMaskEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    MaskEdit_ScrollBy(m.instance, DeltaX , DeltaY)
}

// SetBounds
func (m *TMaskEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    MaskEdit_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (m *TMaskEdit) SetFocus() {
    MaskEdit_SetFocus(m.instance)
}

// Update
func (m *TMaskEdit) Update() {
    MaskEdit_Update(m.instance)
}

// UpdateControlState
func (m *TMaskEdit) UpdateControlState() {
    MaskEdit_UpdateControlState(m.instance)
}

// BringToFront
func (m *TMaskEdit) BringToFront() {
    MaskEdit_BringToFront(m.instance)
}

// ClientToScreen
func (m *TMaskEdit) ClientToScreen(Point TPoint) TPoint {
    return MaskEdit_ClientToScreen(m.instance, Point)
}

// ClientToParent
func (m *TMaskEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return MaskEdit_ClientToParent(m.instance, Point , CheckPtr(AParent))
}

// Dragging
func (m *TMaskEdit) Dragging() bool {
    return MaskEdit_Dragging(m.instance)
}

// HasParent
func (m *TMaskEdit) HasParent() bool {
    return MaskEdit_HasParent(m.instance)
}

// Hide
func (m *TMaskEdit) Hide() {
    MaskEdit_Hide(m.instance)
}

// Perform
func (m *TMaskEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return MaskEdit_Perform(m.instance, Msg , WParam , LParam)
}

// Refresh
func (m *TMaskEdit) Refresh() {
    MaskEdit_Refresh(m.instance)
}

// ScreenToClient
func (m *TMaskEdit) ScreenToClient(Point TPoint) TPoint {
    return MaskEdit_ScreenToClient(m.instance, Point)
}

// ParentToClient
func (m *TMaskEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return MaskEdit_ParentToClient(m.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (m *TMaskEdit) SendToBack() {
    MaskEdit_SendToBack(m.instance)
}

// Show
func (m *TMaskEdit) Show() {
    MaskEdit_Show(m.instance)
}

// GetTextBuf
func (m *TMaskEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    return MaskEdit_GetTextBuf(m.instance, Buffer , BufSize)
}

// SetTextBuf
func (m *TMaskEdit) SetTextBuf(Buffer string) {
    MaskEdit_SetTextBuf(m.instance, Buffer)
}

// FindComponent
func (m *TMaskEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(MaskEdit_FindComponent(m.instance, AName))
}

// GetNamePath
func (m *TMaskEdit) GetNamePath() string {
    return MaskEdit_GetNamePath(m.instance)
}

// Assign
func (m *TMaskEdit) Assign(Source IObject) {
    MaskEdit_Assign(m.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMaskEdit) DisposeOf() {
    MaskEdit_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMaskEdit) ClassType() TClass {
    return MaskEdit_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMaskEdit) ClassName() string {
    return MaskEdit_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMaskEdit) InstanceSize() int32 {
    return MaskEdit_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMaskEdit) InheritsFrom(AClass TClass) bool {
    return MaskEdit_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMaskEdit) Equals(Obj IObject) bool {
    return MaskEdit_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMaskEdit) GetHashCode() int32 {
    return MaskEdit_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMaskEdit) ToString() string {
    return MaskEdit_ToString(m.instance)
}

// Align
func (m *TMaskEdit) Align() TAlign {
    return MaskEdit_GetAlign(m.instance)
}

// SetAlign
func (m *TMaskEdit) SetAlign(value TAlign) {
    MaskEdit_SetAlign(m.instance, value)
}

// Alignment
func (m *TMaskEdit) Alignment() TAlignment {
    return MaskEdit_GetAlignment(m.instance)
}

// SetAlignment
func (m *TMaskEdit) SetAlignment(value TAlignment) {
    MaskEdit_SetAlignment(m.instance, value)
}

// Anchors
func (m *TMaskEdit) Anchors() TAnchors {
    return MaskEdit_GetAnchors(m.instance)
}

// SetAnchors
func (m *TMaskEdit) SetAnchors(value TAnchors) {
    MaskEdit_SetAnchors(m.instance, value)
}

// AutoSelect
func (m *TMaskEdit) AutoSelect() bool {
    return MaskEdit_GetAutoSelect(m.instance)
}

// SetAutoSelect
func (m *TMaskEdit) SetAutoSelect(value bool) {
    MaskEdit_SetAutoSelect(m.instance, value)
}

// AutoSize
func (m *TMaskEdit) AutoSize() bool {
    return MaskEdit_GetAutoSize(m.instance)
}

// SetAutoSize
func (m *TMaskEdit) SetAutoSize(value bool) {
    MaskEdit_SetAutoSize(m.instance, value)
}

// BevelEdges
func (m *TMaskEdit) BevelEdges() TBevelEdges {
    return MaskEdit_GetBevelEdges(m.instance)
}

// SetBevelEdges
func (m *TMaskEdit) SetBevelEdges(value TBevelEdges) {
    MaskEdit_SetBevelEdges(m.instance, value)
}

// BevelInner
func (m *TMaskEdit) BevelInner() TBevelCut {
    return MaskEdit_GetBevelInner(m.instance)
}

// SetBevelInner
func (m *TMaskEdit) SetBevelInner(value TBevelCut) {
    MaskEdit_SetBevelInner(m.instance, value)
}

// BevelOuter
func (m *TMaskEdit) BevelOuter() TBevelCut {
    return MaskEdit_GetBevelOuter(m.instance)
}

// SetBevelOuter
func (m *TMaskEdit) SetBevelOuter(value TBevelCut) {
    MaskEdit_SetBevelOuter(m.instance, value)
}

// BevelKind
func (m *TMaskEdit) BevelKind() TBevelKind {
    return MaskEdit_GetBevelKind(m.instance)
}

// SetBevelKind
func (m *TMaskEdit) SetBevelKind(value TBevelKind) {
    MaskEdit_SetBevelKind(m.instance, value)
}

// BiDiMode
func (m *TMaskEdit) BiDiMode() TBiDiMode {
    return MaskEdit_GetBiDiMode(m.instance)
}

// SetBiDiMode
func (m *TMaskEdit) SetBiDiMode(value TBiDiMode) {
    MaskEdit_SetBiDiMode(m.instance, value)
}

// BorderStyle
func (m *TMaskEdit) BorderStyle() TBorderStyle {
    return MaskEdit_GetBorderStyle(m.instance)
}

// SetBorderStyle
func (m *TMaskEdit) SetBorderStyle(value TBorderStyle) {
    MaskEdit_SetBorderStyle(m.instance, value)
}

// CharCase
func (m *TMaskEdit) CharCase() TEditCharCase {
    return MaskEdit_GetCharCase(m.instance)
}

// SetCharCase
func (m *TMaskEdit) SetCharCase(value TEditCharCase) {
    MaskEdit_SetCharCase(m.instance, value)
}

// Color
func (m *TMaskEdit) Color() TColor {
    return MaskEdit_GetColor(m.instance)
}

// SetColor
func (m *TMaskEdit) SetColor(value TColor) {
    MaskEdit_SetColor(m.instance, value)
}

// DoubleBuffered
func (m *TMaskEdit) DoubleBuffered() bool {
    return MaskEdit_GetDoubleBuffered(m.instance)
}

// SetDoubleBuffered
func (m *TMaskEdit) SetDoubleBuffered(value bool) {
    MaskEdit_SetDoubleBuffered(m.instance, value)
}

// DragCursor
func (m *TMaskEdit) DragCursor() TCursor {
    return MaskEdit_GetDragCursor(m.instance)
}

// SetDragCursor
func (m *TMaskEdit) SetDragCursor(value TCursor) {
    MaskEdit_SetDragCursor(m.instance, value)
}

// DragKind
func (m *TMaskEdit) DragKind() TDragKind {
    return MaskEdit_GetDragKind(m.instance)
}

// SetDragKind
func (m *TMaskEdit) SetDragKind(value TDragKind) {
    MaskEdit_SetDragKind(m.instance, value)
}

// DragMode
func (m *TMaskEdit) DragMode() TDragMode {
    return MaskEdit_GetDragMode(m.instance)
}

// SetDragMode
func (m *TMaskEdit) SetDragMode(value TDragMode) {
    MaskEdit_SetDragMode(m.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (m *TMaskEdit) Enabled() bool {
    return MaskEdit_GetEnabled(m.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (m *TMaskEdit) SetEnabled(value bool) {
    MaskEdit_SetEnabled(m.instance, value)
}

// Font
func (m *TMaskEdit) Font() *TFont {
    return FontFromInst(MaskEdit_GetFont(m.instance))
}

// SetFont
func (m *TMaskEdit) SetFont(value *TFont) {
    MaskEdit_SetFont(m.instance, CheckPtr(value))
}

// MaxLength
func (m *TMaskEdit) MaxLength() int32 {
    return MaskEdit_GetMaxLength(m.instance)
}

// SetMaxLength
func (m *TMaskEdit) SetMaxLength(value int32) {
    MaskEdit_SetMaxLength(m.instance, value)
}

// ParentColor
func (m *TMaskEdit) ParentColor() bool {
    return MaskEdit_GetParentColor(m.instance)
}

// SetParentColor
func (m *TMaskEdit) SetParentColor(value bool) {
    MaskEdit_SetParentColor(m.instance, value)
}

// ParentCtl3D
func (m *TMaskEdit) ParentCtl3D() bool {
    return MaskEdit_GetParentCtl3D(m.instance)
}

// SetParentCtl3D
func (m *TMaskEdit) SetParentCtl3D(value bool) {
    MaskEdit_SetParentCtl3D(m.instance, value)
}

// ParentDoubleBuffered
func (m *TMaskEdit) ParentDoubleBuffered() bool {
    return MaskEdit_GetParentDoubleBuffered(m.instance)
}

// SetParentDoubleBuffered
func (m *TMaskEdit) SetParentDoubleBuffered(value bool) {
    MaskEdit_SetParentDoubleBuffered(m.instance, value)
}

// ParentFont
func (m *TMaskEdit) ParentFont() bool {
    return MaskEdit_GetParentFont(m.instance)
}

// SetParentFont
func (m *TMaskEdit) SetParentFont(value bool) {
    MaskEdit_SetParentFont(m.instance, value)
}

// ParentShowHint
func (m *TMaskEdit) ParentShowHint() bool {
    return MaskEdit_GetParentShowHint(m.instance)
}

// SetParentShowHint
func (m *TMaskEdit) SetParentShowHint(value bool) {
    MaskEdit_SetParentShowHint(m.instance, value)
}

// PasswordChar
func (m *TMaskEdit) PasswordChar() uint16 {
    return MaskEdit_GetPasswordChar(m.instance)
}

// SetPasswordChar
func (m *TMaskEdit) SetPasswordChar(value uint16) {
    MaskEdit_SetPasswordChar(m.instance, value)
}

// PopupMenu
func (m *TMaskEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(MaskEdit_GetPopupMenu(m.instance))
}

// SetPopupMenu
func (m *TMaskEdit) SetPopupMenu(value IComponent) {
    MaskEdit_SetPopupMenu(m.instance, CheckPtr(value))
}

// ReadOnly
func (m *TMaskEdit) ReadOnly() bool {
    return MaskEdit_GetReadOnly(m.instance)
}

// SetReadOnly
func (m *TMaskEdit) SetReadOnly(value bool) {
    MaskEdit_SetReadOnly(m.instance, value)
}

// ShowHint
func (m *TMaskEdit) ShowHint() bool {
    return MaskEdit_GetShowHint(m.instance)
}

// SetShowHint
func (m *TMaskEdit) SetShowHint(value bool) {
    MaskEdit_SetShowHint(m.instance, value)
}

// TabOrder
func (m *TMaskEdit) TabOrder() TTabOrder {
    return MaskEdit_GetTabOrder(m.instance)
}

// SetTabOrder
func (m *TMaskEdit) SetTabOrder(value TTabOrder) {
    MaskEdit_SetTabOrder(m.instance, value)
}

// TabStop
func (m *TMaskEdit) TabStop() bool {
    return MaskEdit_GetTabStop(m.instance)
}

// SetTabStop
func (m *TMaskEdit) SetTabStop(value bool) {
    MaskEdit_SetTabStop(m.instance, value)
}

// Text
func (m *TMaskEdit) Text() string {
    return MaskEdit_GetText(m.instance)
}

// SetText
func (m *TMaskEdit) SetText(value string) {
    MaskEdit_SetText(m.instance, value)
}

// TextHint
func (m *TMaskEdit) TextHint() string {
    return MaskEdit_GetTextHint(m.instance)
}

// SetTextHint
func (m *TMaskEdit) SetTextHint(value string) {
    MaskEdit_SetTextHint(m.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (m *TMaskEdit) Visible() bool {
    return MaskEdit_GetVisible(m.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (m *TMaskEdit) SetVisible(value bool) {
    MaskEdit_SetVisible(m.instance, value)
}

// StyleElements
func (m *TMaskEdit) StyleElements() TStyleElements {
    return MaskEdit_GetStyleElements(m.instance)
}

// SetStyleElements
func (m *TMaskEdit) SetStyleElements(value TStyleElements) {
    MaskEdit_SetStyleElements(m.instance, value)
}

// SetOnChange
func (m *TMaskEdit) SetOnChange(fn TNotifyEvent) {
    MaskEdit_SetOnChange(m.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (m *TMaskEdit) SetOnClick(fn TNotifyEvent) {
    MaskEdit_SetOnClick(m.instance, fn)
}

// SetOnDblClick
func (m *TMaskEdit) SetOnDblClick(fn TNotifyEvent) {
    MaskEdit_SetOnDblClick(m.instance, fn)
}

// SetOnDragDrop
func (m *TMaskEdit) SetOnDragDrop(fn TDragDropEvent) {
    MaskEdit_SetOnDragDrop(m.instance, fn)
}

// SetOnDragOver
func (m *TMaskEdit) SetOnDragOver(fn TDragOverEvent) {
    MaskEdit_SetOnDragOver(m.instance, fn)
}

// SetOnEndDock
func (m *TMaskEdit) SetOnEndDock(fn TEndDragEvent) {
    MaskEdit_SetOnEndDock(m.instance, fn)
}

// SetOnEndDrag
func (m *TMaskEdit) SetOnEndDrag(fn TEndDragEvent) {
    MaskEdit_SetOnEndDrag(m.instance, fn)
}

// SetOnEnter
func (m *TMaskEdit) SetOnEnter(fn TNotifyEvent) {
    MaskEdit_SetOnEnter(m.instance, fn)
}

// SetOnExit
func (m *TMaskEdit) SetOnExit(fn TNotifyEvent) {
    MaskEdit_SetOnExit(m.instance, fn)
}

// SetOnKeyDown
func (m *TMaskEdit) SetOnKeyDown(fn TKeyEvent) {
    MaskEdit_SetOnKeyDown(m.instance, fn)
}

// SetOnKeyPress
func (m *TMaskEdit) SetOnKeyPress(fn TKeyPressEvent) {
    MaskEdit_SetOnKeyPress(m.instance, fn)
}

// SetOnKeyUp
func (m *TMaskEdit) SetOnKeyUp(fn TKeyEvent) {
    MaskEdit_SetOnKeyUp(m.instance, fn)
}

// SetOnMouseDown
func (m *TMaskEdit) SetOnMouseDown(fn TMouseEvent) {
    MaskEdit_SetOnMouseDown(m.instance, fn)
}

// SetOnMouseEnter
func (m *TMaskEdit) SetOnMouseEnter(fn TNotifyEvent) {
    MaskEdit_SetOnMouseEnter(m.instance, fn)
}

// SetOnMouseLeave
func (m *TMaskEdit) SetOnMouseLeave(fn TNotifyEvent) {
    MaskEdit_SetOnMouseLeave(m.instance, fn)
}

// SetOnMouseMove
func (m *TMaskEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    MaskEdit_SetOnMouseMove(m.instance, fn)
}

// SetOnMouseUp
func (m *TMaskEdit) SetOnMouseUp(fn TMouseEvent) {
    MaskEdit_SetOnMouseUp(m.instance, fn)
}

// SetOnStartDock
func (m *TMaskEdit) SetOnStartDock(fn TStartDockEvent) {
    MaskEdit_SetOnStartDock(m.instance, fn)
}

// IsMasked
func (m *TMaskEdit) IsMasked() bool {
    return MaskEdit_GetIsMasked(m.instance)
}

// EditText
func (m *TMaskEdit) EditText() string {
    return MaskEdit_GetEditText(m.instance)
}

// SetEditText
func (m *TMaskEdit) SetEditText(value string) {
    MaskEdit_SetEditText(m.instance, value)
}

// CanUndo
func (m *TMaskEdit) CanUndo() bool {
    return MaskEdit_GetCanUndo(m.instance)
}

// Modified
func (m *TMaskEdit) Modified() bool {
    return MaskEdit_GetModified(m.instance)
}

// SetModified
func (m *TMaskEdit) SetModified(value bool) {
    MaskEdit_SetModified(m.instance, value)
}

// SelLength
func (m *TMaskEdit) SelLength() int32 {
    return MaskEdit_GetSelLength(m.instance)
}

// SetSelLength
func (m *TMaskEdit) SetSelLength(value int32) {
    MaskEdit_SetSelLength(m.instance, value)
}

// SelStart
func (m *TMaskEdit) SelStart() int32 {
    return MaskEdit_GetSelStart(m.instance)
}

// SetSelStart
func (m *TMaskEdit) SetSelStart(value int32) {
    MaskEdit_SetSelStart(m.instance, value)
}

// SelText
func (m *TMaskEdit) SelText() string {
    return MaskEdit_GetSelText(m.instance)
}

// SetSelText
func (m *TMaskEdit) SetSelText(value string) {
    MaskEdit_SetSelText(m.instance, value)
}

// DockClientCount
func (m *TMaskEdit) DockClientCount() int32 {
    return MaskEdit_GetDockClientCount(m.instance)
}

// DockSite
func (m *TMaskEdit) DockSite() bool {
    return MaskEdit_GetDockSite(m.instance)
}

// SetDockSite
func (m *TMaskEdit) SetDockSite(value bool) {
    MaskEdit_SetDockSite(m.instance, value)
}

// AlignDisabled
func (m *TMaskEdit) AlignDisabled() bool {
    return MaskEdit_GetAlignDisabled(m.instance)
}

// MouseInClient
func (m *TMaskEdit) MouseInClient() bool {
    return MaskEdit_GetMouseInClient(m.instance)
}

// VisibleDockClientCount
func (m *TMaskEdit) VisibleDockClientCount() int32 {
    return MaskEdit_GetVisibleDockClientCount(m.instance)
}

// Brush
func (m *TMaskEdit) Brush() *TBrush {
    return BrushFromInst(MaskEdit_GetBrush(m.instance))
}

// ControlCount
func (m *TMaskEdit) ControlCount() int32 {
    return MaskEdit_GetControlCount(m.instance)
}

// Handle
func (m *TMaskEdit) Handle() HWND {
    return MaskEdit_GetHandle(m.instance)
}

// ParentWindow
func (m *TMaskEdit) ParentWindow() HWND {
    return MaskEdit_GetParentWindow(m.instance)
}

// SetParentWindow
func (m *TMaskEdit) SetParentWindow(value HWND) {
    MaskEdit_SetParentWindow(m.instance, value)
}

// UseDockManager
func (m *TMaskEdit) UseDockManager() bool {
    return MaskEdit_GetUseDockManager(m.instance)
}

// SetUseDockManager
func (m *TMaskEdit) SetUseDockManager(value bool) {
    MaskEdit_SetUseDockManager(m.instance, value)
}

// Action
func (m *TMaskEdit) Action() *TAction {
    return ActionFromInst(MaskEdit_GetAction(m.instance))
}

// SetAction
func (m *TMaskEdit) SetAction(value IComponent) {
    MaskEdit_SetAction(m.instance, CheckPtr(value))
}

// BoundsRect
func (m *TMaskEdit) BoundsRect() TRect {
    return MaskEdit_GetBoundsRect(m.instance)
}

// SetBoundsRect
func (m *TMaskEdit) SetBoundsRect(value TRect) {
    MaskEdit_SetBoundsRect(m.instance, value)
}

// ClientHeight
func (m *TMaskEdit) ClientHeight() int32 {
    return MaskEdit_GetClientHeight(m.instance)
}

// SetClientHeight
func (m *TMaskEdit) SetClientHeight(value int32) {
    MaskEdit_SetClientHeight(m.instance, value)
}

// ClientOrigin
func (m *TMaskEdit) ClientOrigin() TPoint {
    return MaskEdit_GetClientOrigin(m.instance)
}

// ClientRect
func (m *TMaskEdit) ClientRect() TRect {
    return MaskEdit_GetClientRect(m.instance)
}

// ClientWidth
func (m *TMaskEdit) ClientWidth() int32 {
    return MaskEdit_GetClientWidth(m.instance)
}

// SetClientWidth
func (m *TMaskEdit) SetClientWidth(value int32) {
    MaskEdit_SetClientWidth(m.instance, value)
}

// ControlState
func (m *TMaskEdit) ControlState() TControlState {
    return MaskEdit_GetControlState(m.instance)
}

// SetControlState
func (m *TMaskEdit) SetControlState(value TControlState) {
    MaskEdit_SetControlState(m.instance, value)
}

// ControlStyle
func (m *TMaskEdit) ControlStyle() TControlStyle {
    return MaskEdit_GetControlStyle(m.instance)
}

// SetControlStyle
func (m *TMaskEdit) SetControlStyle(value TControlStyle) {
    MaskEdit_SetControlStyle(m.instance, value)
}

// ExplicitLeft
func (m *TMaskEdit) ExplicitLeft() int32 {
    return MaskEdit_GetExplicitLeft(m.instance)
}

// ExplicitTop
func (m *TMaskEdit) ExplicitTop() int32 {
    return MaskEdit_GetExplicitTop(m.instance)
}

// ExplicitWidth
func (m *TMaskEdit) ExplicitWidth() int32 {
    return MaskEdit_GetExplicitWidth(m.instance)
}

// ExplicitHeight
func (m *TMaskEdit) ExplicitHeight() int32 {
    return MaskEdit_GetExplicitHeight(m.instance)
}

// Floating
func (m *TMaskEdit) Floating() bool {
    return MaskEdit_GetFloating(m.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (m *TMaskEdit) Parent() *TWinControl {
    return WinControlFromInst(MaskEdit_GetParent(m.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (m *TMaskEdit) SetParent(value IWinControl) {
    MaskEdit_SetParent(m.instance, CheckPtr(value))
}

// AlignWithMargins
func (m *TMaskEdit) AlignWithMargins() bool {
    return MaskEdit_GetAlignWithMargins(m.instance)
}

// SetAlignWithMargins
func (m *TMaskEdit) SetAlignWithMargins(value bool) {
    MaskEdit_SetAlignWithMargins(m.instance, value)
}

// Left
func (m *TMaskEdit) Left() int32 {
    return MaskEdit_GetLeft(m.instance)
}

// SetLeft
func (m *TMaskEdit) SetLeft(value int32) {
    MaskEdit_SetLeft(m.instance, value)
}

// Top
func (m *TMaskEdit) Top() int32 {
    return MaskEdit_GetTop(m.instance)
}

// SetTop
func (m *TMaskEdit) SetTop(value int32) {
    MaskEdit_SetTop(m.instance, value)
}

// Width
func (m *TMaskEdit) Width() int32 {
    return MaskEdit_GetWidth(m.instance)
}

// SetWidth
func (m *TMaskEdit) SetWidth(value int32) {
    MaskEdit_SetWidth(m.instance, value)
}

// Height
func (m *TMaskEdit) Height() int32 {
    return MaskEdit_GetHeight(m.instance)
}

// SetHeight
func (m *TMaskEdit) SetHeight(value int32) {
    MaskEdit_SetHeight(m.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (m *TMaskEdit) Cursor() TCursor {
    return MaskEdit_GetCursor(m.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (m *TMaskEdit) SetCursor(value TCursor) {
    MaskEdit_SetCursor(m.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (m *TMaskEdit) Hint() string {
    return MaskEdit_GetHint(m.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (m *TMaskEdit) SetHint(value string) {
    MaskEdit_SetHint(m.instance, value)
}

// Margins
func (m *TMaskEdit) Margins() *TMargins {
    return MarginsFromInst(MaskEdit_GetMargins(m.instance))
}

// SetMargins
func (m *TMaskEdit) SetMargins(value *TMargins) {
    MaskEdit_SetMargins(m.instance, CheckPtr(value))
}

// CustomHint
func (m *TMaskEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(MaskEdit_GetCustomHint(m.instance))
}

// SetCustomHint
func (m *TMaskEdit) SetCustomHint(value IComponent) {
    MaskEdit_SetCustomHint(m.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (m *TMaskEdit) ComponentCount() int32 {
    return MaskEdit_GetComponentCount(m.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (m *TMaskEdit) ComponentIndex() int32 {
    return MaskEdit_GetComponentIndex(m.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (m *TMaskEdit) SetComponentIndex(value int32) {
    MaskEdit_SetComponentIndex(m.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (m *TMaskEdit) Owner() *TComponent {
    return ComponentFromInst(MaskEdit_GetOwner(m.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (m *TMaskEdit) Name() string {
    return MaskEdit_GetName(m.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (m *TMaskEdit) SetName(value string) {
    MaskEdit_SetName(m.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (m *TMaskEdit) Tag() int {
    return MaskEdit_GetTag(m.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (m *TMaskEdit) SetTag(value int) {
    MaskEdit_SetTag(m.instance, value)
}

// DockClients
func (m *TMaskEdit) DockClients(Index int32) *TControl {
    return ControlFromInst(MaskEdit_GetDockClients(m.instance, Index))
}

// Controls
func (m *TMaskEdit) Controls(Index int32) *TControl {
    return ControlFromInst(MaskEdit_GetControls(m.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (m *TMaskEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(MaskEdit_GetComponents(m.instance, AIndex))
}

