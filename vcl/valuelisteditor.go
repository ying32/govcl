
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

type TValueListEditor struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewValueListEditor
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewValueListEditor(owner IComponent) *TValueListEditor {
    v := new(TValueListEditor)
    v.instance = ValueListEditor_Create(CheckPtr(owner))
    v.ptr = unsafe.Pointer(v.instance)
    return v
}

// ValueListEditorFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ValueListEditorFromInst(inst uintptr) *TValueListEditor {
    v := new(TValueListEditor)
    v.instance = inst
    v.ptr = unsafe.Pointer(inst)
    return v
}

// ValueListEditorFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ValueListEditorFromObj(obj IObject) *TValueListEditor {
    v := new(TValueListEditor)
    v.instance = CheckPtr(obj)
    v.ptr = unsafe.Pointer(v.instance)
    return v
}

// ValueListEditorFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ValueListEditorFromUnsafePointer(ptr unsafe.Pointer) *TValueListEditor {
    v := new(TValueListEditor)
    v.instance = uintptr(ptr)
    v.ptr = ptr
    return v
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (v *TValueListEditor) Free() {
    if v.instance != 0 {
        ValueListEditor_Free(v.instance)
        v.instance = 0
        v.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (v *TValueListEditor) Instance() uintptr {
    return v.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (v *TValueListEditor) UnsafeAddr() unsafe.Pointer {
    return v.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (v *TValueListEditor) IsValid() bool {
    return v.instance != 0
}

// TValueListEditorClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TValueListEditorClass() TClass {
    return ValueListEditor_StaticClassType()
}

// Refresh
func (v *TValueListEditor) Refresh() {
    ValueListEditor_Refresh(v.instance)
}

// CellRect
func (v *TValueListEditor) CellRect(ACol int32, ARow int32) TRect {
    return ValueListEditor_CellRect(v.instance, ACol , ARow)
}

// MouseToCell
func (v *TValueListEditor) MouseToCell(X int32, Y int32, ACol *int32, ARow *int32) {
    ValueListEditor_MouseToCell(v.instance, X , Y , ACol , ARow)
}

// MouseCoord
func (v *TValueListEditor) MouseCoord(X int32, Y int32) TGridCoord {
    return ValueListEditor_MouseCoord(v.instance, X , Y)
}

// CanFocus
func (v *TValueListEditor) CanFocus() bool {
    return ValueListEditor_CanFocus(v.instance)
}

// ContainsControl
func (v *TValueListEditor) ContainsControl(Control IControl) bool {
    return ValueListEditor_ContainsControl(v.instance, CheckPtr(Control))
}

// ControlAtPos
func (v *TValueListEditor) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(ValueListEditor_ControlAtPos(v.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (v *TValueListEditor) DisableAlign() {
    ValueListEditor_DisableAlign(v.instance)
}

// EnableAlign
func (v *TValueListEditor) EnableAlign() {
    ValueListEditor_EnableAlign(v.instance)
}

// FindChildControl
func (v *TValueListEditor) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(ValueListEditor_FindChildControl(v.instance, ControlName))
}

// FlipChildren
func (v *TValueListEditor) FlipChildren(AllLevels bool) {
    ValueListEditor_FlipChildren(v.instance, AllLevels)
}

// Focused
func (v *TValueListEditor) Focused() bool {
    return ValueListEditor_Focused(v.instance)
}

// HandleAllocated
func (v *TValueListEditor) HandleAllocated() bool {
    return ValueListEditor_HandleAllocated(v.instance)
}

// InsertControl
func (v *TValueListEditor) InsertControl(AControl IControl) {
    ValueListEditor_InsertControl(v.instance, CheckPtr(AControl))
}

// Invalidate
func (v *TValueListEditor) Invalidate() {
    ValueListEditor_Invalidate(v.instance)
}

// PaintTo
func (v *TValueListEditor) PaintTo(DC HDC, X int32, Y int32) {
    ValueListEditor_PaintTo(v.instance, DC , X , Y)
}

// RemoveControl
func (v *TValueListEditor) RemoveControl(AControl IControl) {
    ValueListEditor_RemoveControl(v.instance, CheckPtr(AControl))
}

// Realign
func (v *TValueListEditor) Realign() {
    ValueListEditor_Realign(v.instance)
}

// Repaint
func (v *TValueListEditor) Repaint() {
    ValueListEditor_Repaint(v.instance)
}

// ScaleBy
func (v *TValueListEditor) ScaleBy(M int32, D int32) {
    ValueListEditor_ScaleBy(v.instance, M , D)
}

// ScrollBy
func (v *TValueListEditor) ScrollBy(DeltaX int32, DeltaY int32) {
    ValueListEditor_ScrollBy(v.instance, DeltaX , DeltaY)
}

// SetBounds
func (v *TValueListEditor) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ValueListEditor_SetBounds(v.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (v *TValueListEditor) SetFocus() {
    ValueListEditor_SetFocus(v.instance)
}

// Update
func (v *TValueListEditor) Update() {
    ValueListEditor_Update(v.instance)
}

// UpdateControlState
func (v *TValueListEditor) UpdateControlState() {
    ValueListEditor_UpdateControlState(v.instance)
}

// BringToFront
func (v *TValueListEditor) BringToFront() {
    ValueListEditor_BringToFront(v.instance)
}

// ClientToScreen
func (v *TValueListEditor) ClientToScreen(Point TPoint) TPoint {
    return ValueListEditor_ClientToScreen(v.instance, Point)
}

// ClientToParent
func (v *TValueListEditor) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ValueListEditor_ClientToParent(v.instance, Point , CheckPtr(AParent))
}

// Dragging
func (v *TValueListEditor) Dragging() bool {
    return ValueListEditor_Dragging(v.instance)
}

// HasParent
func (v *TValueListEditor) HasParent() bool {
    return ValueListEditor_HasParent(v.instance)
}

// Hide
func (v *TValueListEditor) Hide() {
    ValueListEditor_Hide(v.instance)
}

// Perform
func (v *TValueListEditor) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ValueListEditor_Perform(v.instance, Msg , WParam , LParam)
}

// ScreenToClient
func (v *TValueListEditor) ScreenToClient(Point TPoint) TPoint {
    return ValueListEditor_ScreenToClient(v.instance, Point)
}

// ParentToClient
func (v *TValueListEditor) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ValueListEditor_ParentToClient(v.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (v *TValueListEditor) SendToBack() {
    ValueListEditor_SendToBack(v.instance)
}

// Show
func (v *TValueListEditor) Show() {
    ValueListEditor_Show(v.instance)
}

// GetTextBuf
func (v *TValueListEditor) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ValueListEditor_GetTextBuf(v.instance, Buffer , BufSize)
}

// GetTextLen
func (v *TValueListEditor) GetTextLen() int32 {
    return ValueListEditor_GetTextLen(v.instance)
}

// SetTextBuf
func (v *TValueListEditor) SetTextBuf(Buffer string) {
    ValueListEditor_SetTextBuf(v.instance, Buffer)
}

// FindComponent
func (v *TValueListEditor) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ValueListEditor_FindComponent(v.instance, AName))
}

// GetNamePath
func (v *TValueListEditor) GetNamePath() string {
    return ValueListEditor_GetNamePath(v.instance)
}

// Assign
func (v *TValueListEditor) Assign(Source IObject) {
    ValueListEditor_Assign(v.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (v *TValueListEditor) DisposeOf() {
    ValueListEditor_DisposeOf(v.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (v *TValueListEditor) ClassType() TClass {
    return ValueListEditor_ClassType(v.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (v *TValueListEditor) ClassName() string {
    return ValueListEditor_ClassName(v.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (v *TValueListEditor) InstanceSize() int32 {
    return ValueListEditor_InstanceSize(v.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (v *TValueListEditor) InheritsFrom(AClass TClass) bool {
    return ValueListEditor_InheritsFrom(v.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (v *TValueListEditor) Equals(Obj IObject) bool {
    return ValueListEditor_Equals(v.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (v *TValueListEditor) GetHashCode() int32 {
    return ValueListEditor_GetHashCode(v.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (v *TValueListEditor) ToString() string {
    return ValueListEditor_ToString(v.instance)
}

// ColCount
func (v *TValueListEditor) ColCount() int32 {
    return ValueListEditor_GetColCount(v.instance)
}

// SetColCount
func (v *TValueListEditor) SetColCount(value int32) {
    ValueListEditor_SetColCount(v.instance, value)
}

// RowCount
func (v *TValueListEditor) RowCount() int32 {
    return ValueListEditor_GetRowCount(v.instance)
}

// VisibleColCount
func (v *TValueListEditor) VisibleColCount() int32 {
    return ValueListEditor_GetVisibleColCount(v.instance)
}

// VisibleRowCount
func (v *TValueListEditor) VisibleRowCount() int32 {
    return ValueListEditor_GetVisibleRowCount(v.instance)
}

// Align
func (v *TValueListEditor) Align() TAlign {
    return ValueListEditor_GetAlign(v.instance)
}

// SetAlign
func (v *TValueListEditor) SetAlign(value TAlign) {
    ValueListEditor_SetAlign(v.instance, value)
}

// Anchors
func (v *TValueListEditor) Anchors() TAnchors {
    return ValueListEditor_GetAnchors(v.instance)
}

// SetAnchors
func (v *TValueListEditor) SetAnchors(value TAnchors) {
    ValueListEditor_SetAnchors(v.instance, value)
}

// BiDiMode
func (v *TValueListEditor) BiDiMode() TBiDiMode {
    return ValueListEditor_GetBiDiMode(v.instance)
}

// SetBiDiMode
func (v *TValueListEditor) SetBiDiMode(value TBiDiMode) {
    ValueListEditor_SetBiDiMode(v.instance, value)
}

// BorderStyle
func (v *TValueListEditor) BorderStyle() TBorderStyle {
    return ValueListEditor_GetBorderStyle(v.instance)
}

// SetBorderStyle
func (v *TValueListEditor) SetBorderStyle(value TBorderStyle) {
    ValueListEditor_SetBorderStyle(v.instance, value)
}

// Color
func (v *TValueListEditor) Color() TColor {
    return ValueListEditor_GetColor(v.instance)
}

// SetColor
func (v *TValueListEditor) SetColor(value TColor) {
    ValueListEditor_SetColor(v.instance, value)
}

// DefaultColWidth
func (v *TValueListEditor) DefaultColWidth() int32 {
    return ValueListEditor_GetDefaultColWidth(v.instance)
}

// SetDefaultColWidth
func (v *TValueListEditor) SetDefaultColWidth(value int32) {
    ValueListEditor_SetDefaultColWidth(v.instance, value)
}

// DefaultDrawing
func (v *TValueListEditor) DefaultDrawing() bool {
    return ValueListEditor_GetDefaultDrawing(v.instance)
}

// SetDefaultDrawing
func (v *TValueListEditor) SetDefaultDrawing(value bool) {
    ValueListEditor_SetDefaultDrawing(v.instance, value)
}

// DefaultRowHeight
func (v *TValueListEditor) DefaultRowHeight() int32 {
    return ValueListEditor_GetDefaultRowHeight(v.instance)
}

// SetDefaultRowHeight
func (v *TValueListEditor) SetDefaultRowHeight(value int32) {
    ValueListEditor_SetDefaultRowHeight(v.instance, value)
}

// DoubleBuffered
func (v *TValueListEditor) DoubleBuffered() bool {
    return ValueListEditor_GetDoubleBuffered(v.instance)
}

// SetDoubleBuffered
func (v *TValueListEditor) SetDoubleBuffered(value bool) {
    ValueListEditor_SetDoubleBuffered(v.instance, value)
}

// DragCursor
func (v *TValueListEditor) DragCursor() TCursor {
    return ValueListEditor_GetDragCursor(v.instance)
}

// SetDragCursor
func (v *TValueListEditor) SetDragCursor(value TCursor) {
    ValueListEditor_SetDragCursor(v.instance, value)
}

// DragKind
func (v *TValueListEditor) DragKind() TDragKind {
    return ValueListEditor_GetDragKind(v.instance)
}

// SetDragKind
func (v *TValueListEditor) SetDragKind(value TDragKind) {
    ValueListEditor_SetDragKind(v.instance, value)
}

// DragMode
func (v *TValueListEditor) DragMode() TDragMode {
    return ValueListEditor_GetDragMode(v.instance)
}

// SetDragMode
func (v *TValueListEditor) SetDragMode(value TDragMode) {
    ValueListEditor_SetDragMode(v.instance, value)
}

// DrawingStyle
func (v *TValueListEditor) DrawingStyle() TGridDrawingStyle {
    return ValueListEditor_GetDrawingStyle(v.instance)
}

// SetDrawingStyle
func (v *TValueListEditor) SetDrawingStyle(value TGridDrawingStyle) {
    ValueListEditor_SetDrawingStyle(v.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (v *TValueListEditor) Enabled() bool {
    return ValueListEditor_GetEnabled(v.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (v *TValueListEditor) SetEnabled(value bool) {
    ValueListEditor_SetEnabled(v.instance, value)
}

// FixedColor
func (v *TValueListEditor) FixedColor() TColor {
    return ValueListEditor_GetFixedColor(v.instance)
}

// SetFixedColor
func (v *TValueListEditor) SetFixedColor(value TColor) {
    ValueListEditor_SetFixedColor(v.instance, value)
}

// FixedCols
func (v *TValueListEditor) FixedCols() int32 {
    return ValueListEditor_GetFixedCols(v.instance)
}

// SetFixedCols
func (v *TValueListEditor) SetFixedCols(value int32) {
    ValueListEditor_SetFixedCols(v.instance, value)
}

// Font
func (v *TValueListEditor) Font() *TFont {
    return FontFromInst(ValueListEditor_GetFont(v.instance))
}

// SetFont
func (v *TValueListEditor) SetFont(value *TFont) {
    ValueListEditor_SetFont(v.instance, CheckPtr(value))
}

// GradientEndColor
func (v *TValueListEditor) GradientEndColor() TColor {
    return ValueListEditor_GetGradientEndColor(v.instance)
}

// SetGradientEndColor
func (v *TValueListEditor) SetGradientEndColor(value TColor) {
    ValueListEditor_SetGradientEndColor(v.instance, value)
}

// GradientStartColor
func (v *TValueListEditor) GradientStartColor() TColor {
    return ValueListEditor_GetGradientStartColor(v.instance)
}

// SetGradientStartColor
func (v *TValueListEditor) SetGradientStartColor(value TColor) {
    ValueListEditor_SetGradientStartColor(v.instance, value)
}

// GridLineWidth
func (v *TValueListEditor) GridLineWidth() int32 {
    return ValueListEditor_GetGridLineWidth(v.instance)
}

// SetGridLineWidth
func (v *TValueListEditor) SetGridLineWidth(value int32) {
    ValueListEditor_SetGridLineWidth(v.instance, value)
}

// Options
func (v *TValueListEditor) Options() TGridOptions {
    return ValueListEditor_GetOptions(v.instance)
}

// SetOptions
func (v *TValueListEditor) SetOptions(value TGridOptions) {
    ValueListEditor_SetOptions(v.instance, value)
}

// ParentColor
func (v *TValueListEditor) ParentColor() bool {
    return ValueListEditor_GetParentColor(v.instance)
}

// SetParentColor
func (v *TValueListEditor) SetParentColor(value bool) {
    ValueListEditor_SetParentColor(v.instance, value)
}

// ParentCtl3D
func (v *TValueListEditor) ParentCtl3D() bool {
    return ValueListEditor_GetParentCtl3D(v.instance)
}

// SetParentCtl3D
func (v *TValueListEditor) SetParentCtl3D(value bool) {
    ValueListEditor_SetParentCtl3D(v.instance, value)
}

// ParentDoubleBuffered
func (v *TValueListEditor) ParentDoubleBuffered() bool {
    return ValueListEditor_GetParentDoubleBuffered(v.instance)
}

// SetParentDoubleBuffered
func (v *TValueListEditor) SetParentDoubleBuffered(value bool) {
    ValueListEditor_SetParentDoubleBuffered(v.instance, value)
}

// ParentFont
func (v *TValueListEditor) ParentFont() bool {
    return ValueListEditor_GetParentFont(v.instance)
}

// SetParentFont
func (v *TValueListEditor) SetParentFont(value bool) {
    ValueListEditor_SetParentFont(v.instance, value)
}

// ParentShowHint
func (v *TValueListEditor) ParentShowHint() bool {
    return ValueListEditor_GetParentShowHint(v.instance)
}

// SetParentShowHint
func (v *TValueListEditor) SetParentShowHint(value bool) {
    ValueListEditor_SetParentShowHint(v.instance, value)
}

// PopupMenu
func (v *TValueListEditor) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ValueListEditor_GetPopupMenu(v.instance))
}

// SetPopupMenu
func (v *TValueListEditor) SetPopupMenu(value IComponent) {
    ValueListEditor_SetPopupMenu(v.instance, CheckPtr(value))
}

// ScrollBars
func (v *TValueListEditor) ScrollBars() TScrollStyle {
    return ValueListEditor_GetScrollBars(v.instance)
}

// SetScrollBars
func (v *TValueListEditor) SetScrollBars(value TScrollStyle) {
    ValueListEditor_SetScrollBars(v.instance, value)
}

// ShowHint
func (v *TValueListEditor) ShowHint() bool {
    return ValueListEditor_GetShowHint(v.instance)
}

// SetShowHint
func (v *TValueListEditor) SetShowHint(value bool) {
    ValueListEditor_SetShowHint(v.instance, value)
}

// Strings
func (v *TValueListEditor) Strings() *TStrings {
    return StringsFromInst(ValueListEditor_GetStrings(v.instance))
}

// SetStrings
func (v *TValueListEditor) SetStrings(value IObject) {
    ValueListEditor_SetStrings(v.instance, CheckPtr(value))
}

// TabOrder
func (v *TValueListEditor) TabOrder() TTabOrder {
    return ValueListEditor_GetTabOrder(v.instance)
}

// SetTabOrder
func (v *TValueListEditor) SetTabOrder(value TTabOrder) {
    ValueListEditor_SetTabOrder(v.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (v *TValueListEditor) Visible() bool {
    return ValueListEditor_GetVisible(v.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (v *TValueListEditor) SetVisible(value bool) {
    ValueListEditor_SetVisible(v.instance, value)
}

// StyleElements
func (v *TValueListEditor) StyleElements() TStyleElements {
    return ValueListEditor_GetStyleElements(v.instance)
}

// SetStyleElements
func (v *TValueListEditor) SetStyleElements(value TStyleElements) {
    ValueListEditor_SetStyleElements(v.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (v *TValueListEditor) SetOnClick(fn TNotifyEvent) {
    ValueListEditor_SetOnClick(v.instance, fn)
}

// SetOnContextPopup
func (v *TValueListEditor) SetOnContextPopup(fn TContextPopupEvent) {
    ValueListEditor_SetOnContextPopup(v.instance, fn)
}

// SetOnDblClick
func (v *TValueListEditor) SetOnDblClick(fn TNotifyEvent) {
    ValueListEditor_SetOnDblClick(v.instance, fn)
}

// SetOnDragDrop
func (v *TValueListEditor) SetOnDragDrop(fn TDragDropEvent) {
    ValueListEditor_SetOnDragDrop(v.instance, fn)
}

// SetOnDragOver
func (v *TValueListEditor) SetOnDragOver(fn TDragOverEvent) {
    ValueListEditor_SetOnDragOver(v.instance, fn)
}

// SetOnDrawCell
func (v *TValueListEditor) SetOnDrawCell(fn TDrawCellEvent) {
    ValueListEditor_SetOnDrawCell(v.instance, fn)
}

// SetOnEndDock
func (v *TValueListEditor) SetOnEndDock(fn TEndDragEvent) {
    ValueListEditor_SetOnEndDock(v.instance, fn)
}

// SetOnEndDrag
func (v *TValueListEditor) SetOnEndDrag(fn TEndDragEvent) {
    ValueListEditor_SetOnEndDrag(v.instance, fn)
}

// SetOnEnter
func (v *TValueListEditor) SetOnEnter(fn TNotifyEvent) {
    ValueListEditor_SetOnEnter(v.instance, fn)
}

// SetOnExit
func (v *TValueListEditor) SetOnExit(fn TNotifyEvent) {
    ValueListEditor_SetOnExit(v.instance, fn)
}

// SetOnGetEditMask
func (v *TValueListEditor) SetOnGetEditMask(fn TGetEditEvent) {
    ValueListEditor_SetOnGetEditMask(v.instance, fn)
}

// SetOnGetEditText
func (v *TValueListEditor) SetOnGetEditText(fn TGetEditEvent) {
    ValueListEditor_SetOnGetEditText(v.instance, fn)
}

// SetOnKeyDown
func (v *TValueListEditor) SetOnKeyDown(fn TKeyEvent) {
    ValueListEditor_SetOnKeyDown(v.instance, fn)
}

// SetOnKeyPress
func (v *TValueListEditor) SetOnKeyPress(fn TKeyPressEvent) {
    ValueListEditor_SetOnKeyPress(v.instance, fn)
}

// SetOnKeyUp
func (v *TValueListEditor) SetOnKeyUp(fn TKeyEvent) {
    ValueListEditor_SetOnKeyUp(v.instance, fn)
}

// SetOnMouseDown
func (v *TValueListEditor) SetOnMouseDown(fn TMouseEvent) {
    ValueListEditor_SetOnMouseDown(v.instance, fn)
}

// SetOnMouseEnter
func (v *TValueListEditor) SetOnMouseEnter(fn TNotifyEvent) {
    ValueListEditor_SetOnMouseEnter(v.instance, fn)
}

// SetOnMouseLeave
func (v *TValueListEditor) SetOnMouseLeave(fn TNotifyEvent) {
    ValueListEditor_SetOnMouseLeave(v.instance, fn)
}

// SetOnMouseMove
func (v *TValueListEditor) SetOnMouseMove(fn TMouseMoveEvent) {
    ValueListEditor_SetOnMouseMove(v.instance, fn)
}

// SetOnMouseUp
func (v *TValueListEditor) SetOnMouseUp(fn TMouseEvent) {
    ValueListEditor_SetOnMouseUp(v.instance, fn)
}

// SetOnMouseWheelDown
func (v *TValueListEditor) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    ValueListEditor_SetOnMouseWheelDown(v.instance, fn)
}

// SetOnMouseWheelUp
func (v *TValueListEditor) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    ValueListEditor_SetOnMouseWheelUp(v.instance, fn)
}

// SetOnRowMoved
func (v *TValueListEditor) SetOnRowMoved(fn TMovedEvent) {
    ValueListEditor_SetOnRowMoved(v.instance, fn)
}

// SetOnSelectCell
func (v *TValueListEditor) SetOnSelectCell(fn TSelectCellEvent) {
    ValueListEditor_SetOnSelectCell(v.instance, fn)
}

// SetOnSetEditText
func (v *TValueListEditor) SetOnSetEditText(fn TSetEditEvent) {
    ValueListEditor_SetOnSetEditText(v.instance, fn)
}

// SetOnStartDock
func (v *TValueListEditor) SetOnStartDock(fn TStartDockEvent) {
    ValueListEditor_SetOnStartDock(v.instance, fn)
}

// SetOnTopLeftChanged
func (v *TValueListEditor) SetOnTopLeftChanged(fn TNotifyEvent) {
    ValueListEditor_SetOnTopLeftChanged(v.instance, fn)
}

// Canvas
func (v *TValueListEditor) Canvas() *TCanvas {
    return CanvasFromInst(ValueListEditor_GetCanvas(v.instance))
}

// Col
func (v *TValueListEditor) Col() int32 {
    return ValueListEditor_GetCol(v.instance)
}

// SetCol
func (v *TValueListEditor) SetCol(value int32) {
    ValueListEditor_SetCol(v.instance, value)
}

// EditorMode
func (v *TValueListEditor) EditorMode() bool {
    return ValueListEditor_GetEditorMode(v.instance)
}

// SetEditorMode
func (v *TValueListEditor) SetEditorMode(value bool) {
    ValueListEditor_SetEditorMode(v.instance, value)
}

// GridHeight
func (v *TValueListEditor) GridHeight() int32 {
    return ValueListEditor_GetGridHeight(v.instance)
}

// GridWidth
func (v *TValueListEditor) GridWidth() int32 {
    return ValueListEditor_GetGridWidth(v.instance)
}

// LeftCol
func (v *TValueListEditor) LeftCol() int32 {
    return ValueListEditor_GetLeftCol(v.instance)
}

// SetLeftCol
func (v *TValueListEditor) SetLeftCol(value int32) {
    ValueListEditor_SetLeftCol(v.instance, value)
}

// Selection
func (v *TValueListEditor) Selection() TGridRect {
    return ValueListEditor_GetSelection(v.instance)
}

// SetSelection
func (v *TValueListEditor) SetSelection(value TGridRect) {
    ValueListEditor_SetSelection(v.instance, value)
}

// Row
func (v *TValueListEditor) Row() int32 {
    return ValueListEditor_GetRow(v.instance)
}

// SetRow
func (v *TValueListEditor) SetRow(value int32) {
    ValueListEditor_SetRow(v.instance, value)
}

// TopRow
func (v *TValueListEditor) TopRow() int32 {
    return ValueListEditor_GetTopRow(v.instance)
}

// SetTopRow
func (v *TValueListEditor) SetTopRow(value int32) {
    ValueListEditor_SetTopRow(v.instance, value)
}

// TabStop
func (v *TValueListEditor) TabStop() bool {
    return ValueListEditor_GetTabStop(v.instance)
}

// SetTabStop
func (v *TValueListEditor) SetTabStop(value bool) {
    ValueListEditor_SetTabStop(v.instance, value)
}

// DockClientCount
func (v *TValueListEditor) DockClientCount() int32 {
    return ValueListEditor_GetDockClientCount(v.instance)
}

// DockSite
func (v *TValueListEditor) DockSite() bool {
    return ValueListEditor_GetDockSite(v.instance)
}

// SetDockSite
func (v *TValueListEditor) SetDockSite(value bool) {
    ValueListEditor_SetDockSite(v.instance, value)
}

// AlignDisabled
func (v *TValueListEditor) AlignDisabled() bool {
    return ValueListEditor_GetAlignDisabled(v.instance)
}

// MouseInClient
func (v *TValueListEditor) MouseInClient() bool {
    return ValueListEditor_GetMouseInClient(v.instance)
}

// VisibleDockClientCount
func (v *TValueListEditor) VisibleDockClientCount() int32 {
    return ValueListEditor_GetVisibleDockClientCount(v.instance)
}

// Brush
func (v *TValueListEditor) Brush() *TBrush {
    return BrushFromInst(ValueListEditor_GetBrush(v.instance))
}

// ControlCount
func (v *TValueListEditor) ControlCount() int32 {
    return ValueListEditor_GetControlCount(v.instance)
}

// Handle
func (v *TValueListEditor) Handle() HWND {
    return ValueListEditor_GetHandle(v.instance)
}

// ParentWindow
func (v *TValueListEditor) ParentWindow() HWND {
    return ValueListEditor_GetParentWindow(v.instance)
}

// SetParentWindow
func (v *TValueListEditor) SetParentWindow(value HWND) {
    ValueListEditor_SetParentWindow(v.instance, value)
}

// UseDockManager
func (v *TValueListEditor) UseDockManager() bool {
    return ValueListEditor_GetUseDockManager(v.instance)
}

// SetUseDockManager
func (v *TValueListEditor) SetUseDockManager(value bool) {
    ValueListEditor_SetUseDockManager(v.instance, value)
}

// Action
func (v *TValueListEditor) Action() *TAction {
    return ActionFromInst(ValueListEditor_GetAction(v.instance))
}

// SetAction
func (v *TValueListEditor) SetAction(value IComponent) {
    ValueListEditor_SetAction(v.instance, CheckPtr(value))
}

// BoundsRect
func (v *TValueListEditor) BoundsRect() TRect {
    return ValueListEditor_GetBoundsRect(v.instance)
}

// SetBoundsRect
func (v *TValueListEditor) SetBoundsRect(value TRect) {
    ValueListEditor_SetBoundsRect(v.instance, value)
}

// ClientHeight
func (v *TValueListEditor) ClientHeight() int32 {
    return ValueListEditor_GetClientHeight(v.instance)
}

// SetClientHeight
func (v *TValueListEditor) SetClientHeight(value int32) {
    ValueListEditor_SetClientHeight(v.instance, value)
}

// ClientOrigin
func (v *TValueListEditor) ClientOrigin() TPoint {
    return ValueListEditor_GetClientOrigin(v.instance)
}

// ClientRect
func (v *TValueListEditor) ClientRect() TRect {
    return ValueListEditor_GetClientRect(v.instance)
}

// ClientWidth
func (v *TValueListEditor) ClientWidth() int32 {
    return ValueListEditor_GetClientWidth(v.instance)
}

// SetClientWidth
func (v *TValueListEditor) SetClientWidth(value int32) {
    ValueListEditor_SetClientWidth(v.instance, value)
}

// ControlState
func (v *TValueListEditor) ControlState() TControlState {
    return ValueListEditor_GetControlState(v.instance)
}

// SetControlState
func (v *TValueListEditor) SetControlState(value TControlState) {
    ValueListEditor_SetControlState(v.instance, value)
}

// ControlStyle
func (v *TValueListEditor) ControlStyle() TControlStyle {
    return ValueListEditor_GetControlStyle(v.instance)
}

// SetControlStyle
func (v *TValueListEditor) SetControlStyle(value TControlStyle) {
    ValueListEditor_SetControlStyle(v.instance, value)
}

// ExplicitLeft
func (v *TValueListEditor) ExplicitLeft() int32 {
    return ValueListEditor_GetExplicitLeft(v.instance)
}

// ExplicitTop
func (v *TValueListEditor) ExplicitTop() int32 {
    return ValueListEditor_GetExplicitTop(v.instance)
}

// ExplicitWidth
func (v *TValueListEditor) ExplicitWidth() int32 {
    return ValueListEditor_GetExplicitWidth(v.instance)
}

// ExplicitHeight
func (v *TValueListEditor) ExplicitHeight() int32 {
    return ValueListEditor_GetExplicitHeight(v.instance)
}

// Floating
func (v *TValueListEditor) Floating() bool {
    return ValueListEditor_GetFloating(v.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (v *TValueListEditor) Parent() *TWinControl {
    return WinControlFromInst(ValueListEditor_GetParent(v.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (v *TValueListEditor) SetParent(value IWinControl) {
    ValueListEditor_SetParent(v.instance, CheckPtr(value))
}

// AlignWithMargins
func (v *TValueListEditor) AlignWithMargins() bool {
    return ValueListEditor_GetAlignWithMargins(v.instance)
}

// SetAlignWithMargins
func (v *TValueListEditor) SetAlignWithMargins(value bool) {
    ValueListEditor_SetAlignWithMargins(v.instance, value)
}

// Left
func (v *TValueListEditor) Left() int32 {
    return ValueListEditor_GetLeft(v.instance)
}

// SetLeft
func (v *TValueListEditor) SetLeft(value int32) {
    ValueListEditor_SetLeft(v.instance, value)
}

// Top
func (v *TValueListEditor) Top() int32 {
    return ValueListEditor_GetTop(v.instance)
}

// SetTop
func (v *TValueListEditor) SetTop(value int32) {
    ValueListEditor_SetTop(v.instance, value)
}

// Width
func (v *TValueListEditor) Width() int32 {
    return ValueListEditor_GetWidth(v.instance)
}

// SetWidth
func (v *TValueListEditor) SetWidth(value int32) {
    ValueListEditor_SetWidth(v.instance, value)
}

// Height
func (v *TValueListEditor) Height() int32 {
    return ValueListEditor_GetHeight(v.instance)
}

// SetHeight
func (v *TValueListEditor) SetHeight(value int32) {
    ValueListEditor_SetHeight(v.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (v *TValueListEditor) Cursor() TCursor {
    return ValueListEditor_GetCursor(v.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (v *TValueListEditor) SetCursor(value TCursor) {
    ValueListEditor_SetCursor(v.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (v *TValueListEditor) Hint() string {
    return ValueListEditor_GetHint(v.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (v *TValueListEditor) SetHint(value string) {
    ValueListEditor_SetHint(v.instance, value)
}

// Margins
func (v *TValueListEditor) Margins() *TMargins {
    return MarginsFromInst(ValueListEditor_GetMargins(v.instance))
}

// SetMargins
func (v *TValueListEditor) SetMargins(value *TMargins) {
    ValueListEditor_SetMargins(v.instance, CheckPtr(value))
}

// CustomHint
func (v *TValueListEditor) CustomHint() *TCustomHint {
    return CustomHintFromInst(ValueListEditor_GetCustomHint(v.instance))
}

// SetCustomHint
func (v *TValueListEditor) SetCustomHint(value IComponent) {
    ValueListEditor_SetCustomHint(v.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (v *TValueListEditor) ComponentCount() int32 {
    return ValueListEditor_GetComponentCount(v.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (v *TValueListEditor) ComponentIndex() int32 {
    return ValueListEditor_GetComponentIndex(v.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (v *TValueListEditor) SetComponentIndex(value int32) {
    ValueListEditor_SetComponentIndex(v.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (v *TValueListEditor) Owner() *TComponent {
    return ComponentFromInst(ValueListEditor_GetOwner(v.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (v *TValueListEditor) Name() string {
    return ValueListEditor_GetName(v.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (v *TValueListEditor) SetName(value string) {
    ValueListEditor_SetName(v.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (v *TValueListEditor) Tag() int {
    return ValueListEditor_GetTag(v.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (v *TValueListEditor) SetTag(value int) {
    ValueListEditor_SetTag(v.instance, value)
}

// Cells
func (v *TValueListEditor) Cells(ACol int32, ARow int32) string {
    return ValueListEditor_GetCells(v.instance, ACol, ARow)
}

// Cells
func (v *TValueListEditor) SetCells(ACol int32, ARow int32, value string) {
    ValueListEditor_SetCells(v.instance, ACol, ARow, value)
}

// Values
func (v *TValueListEditor) Values(Key string) string {
    return ValueListEditor_GetValues(v.instance, Key)
}

// Values
func (v *TValueListEditor) SetValues(Key string, value string) {
    ValueListEditor_SetValues(v.instance, Key, value)
}

// ColWidths
func (v *TValueListEditor) ColWidths(Index int32) int32 {
    return ValueListEditor_GetColWidths(v.instance, Index)
}

// ColWidths
func (v *TValueListEditor) SetColWidths(Index int32, value int32) {
    ValueListEditor_SetColWidths(v.instance, Index, value)
}

// RowHeights
func (v *TValueListEditor) RowHeights(Index int32) int32 {
    return ValueListEditor_GetRowHeights(v.instance, Index)
}

// RowHeights
func (v *TValueListEditor) SetRowHeights(Index int32, value int32) {
    ValueListEditor_SetRowHeights(v.instance, Index, value)
}

// TabStops
func (v *TValueListEditor) TabStops(Index int32) bool {
    return ValueListEditor_GetTabStops(v.instance, Index)
}

// TabStops
func (v *TValueListEditor) SetTabStops(Index int32, value bool) {
    ValueListEditor_SetTabStops(v.instance, Index, value)
}

// DockClients
func (v *TValueListEditor) DockClients(Index int32) *TControl {
    return ControlFromInst(ValueListEditor_GetDockClients(v.instance, Index))
}

// Controls
func (v *TValueListEditor) Controls(Index int32) *TControl {
    return ControlFromInst(ValueListEditor_GetControls(v.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (v *TValueListEditor) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ValueListEditor_GetComponents(v.instance, AIndex))
}

