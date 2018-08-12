
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

type TStringGrid struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewStringGrid
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewStringGrid(owner IComponent) *TStringGrid {
    s := new(TStringGrid)
    s.instance = StringGrid_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StringGridFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func StringGridFromInst(inst uintptr) *TStringGrid {
    s := new(TStringGrid)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// StringGridFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func StringGridFromObj(obj IObject) *TStringGrid {
    s := new(TStringGrid)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StringGridFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func StringGridFromUnsafePointer(ptr unsafe.Pointer) *TStringGrid {
    s := new(TStringGrid)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TStringGrid) Free() {
    if s.instance != 0 {
        StringGrid_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TStringGrid) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TStringGrid) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TStringGrid) IsValid() bool {
    return s.instance != 0
}

// TStringGridClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TStringGridClass() TClass {
    return StringGrid_StaticClassType()
}

// CellRect
func (s *TStringGrid) CellRect(ACol int32, ARow int32) TRect {
    return StringGrid_CellRect(s.instance, ACol , ARow)
}

// MouseToCell
func (s *TStringGrid) MouseToCell(X int32, Y int32, ACol *int32, ARow *int32) {
    StringGrid_MouseToCell(s.instance, X , Y , ACol , ARow)
}

// MouseCoord
func (s *TStringGrid) MouseCoord(X int32, Y int32) TGridCoord {
    return StringGrid_MouseCoord(s.instance, X , Y)
}

// CanFocus
func (s *TStringGrid) CanFocus() bool {
    return StringGrid_CanFocus(s.instance)
}

// ContainsControl
func (s *TStringGrid) ContainsControl(Control IControl) bool {
    return StringGrid_ContainsControl(s.instance, CheckPtr(Control))
}

// ControlAtPos
func (s *TStringGrid) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(StringGrid_ControlAtPos(s.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (s *TStringGrid) DisableAlign() {
    StringGrid_DisableAlign(s.instance)
}

// EnableAlign
func (s *TStringGrid) EnableAlign() {
    StringGrid_EnableAlign(s.instance)
}

// FindChildControl
func (s *TStringGrid) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(StringGrid_FindChildControl(s.instance, ControlName))
}

// FlipChildren
func (s *TStringGrid) FlipChildren(AllLevels bool) {
    StringGrid_FlipChildren(s.instance, AllLevels)
}

// Focused
func (s *TStringGrid) Focused() bool {
    return StringGrid_Focused(s.instance)
}

// HandleAllocated
func (s *TStringGrid) HandleAllocated() bool {
    return StringGrid_HandleAllocated(s.instance)
}

// InsertControl
func (s *TStringGrid) InsertControl(AControl IControl) {
    StringGrid_InsertControl(s.instance, CheckPtr(AControl))
}

// Invalidate
func (s *TStringGrid) Invalidate() {
    StringGrid_Invalidate(s.instance)
}

// PaintTo
func (s *TStringGrid) PaintTo(DC HDC, X int32, Y int32) {
    StringGrid_PaintTo(s.instance, DC , X , Y)
}

// RemoveControl
func (s *TStringGrid) RemoveControl(AControl IControl) {
    StringGrid_RemoveControl(s.instance, CheckPtr(AControl))
}

// Realign
func (s *TStringGrid) Realign() {
    StringGrid_Realign(s.instance)
}

// Repaint
func (s *TStringGrid) Repaint() {
    StringGrid_Repaint(s.instance)
}

// ScaleBy
func (s *TStringGrid) ScaleBy(M int32, D int32) {
    StringGrid_ScaleBy(s.instance, M , D)
}

// ScrollBy
func (s *TStringGrid) ScrollBy(DeltaX int32, DeltaY int32) {
    StringGrid_ScrollBy(s.instance, DeltaX , DeltaY)
}

// SetBounds
func (s *TStringGrid) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    StringGrid_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (s *TStringGrid) SetFocus() {
    StringGrid_SetFocus(s.instance)
}

// Update
func (s *TStringGrid) Update() {
    StringGrid_Update(s.instance)
}

// UpdateControlState
func (s *TStringGrid) UpdateControlState() {
    StringGrid_UpdateControlState(s.instance)
}

// BringToFront
func (s *TStringGrid) BringToFront() {
    StringGrid_BringToFront(s.instance)
}

// ClientToScreen
func (s *TStringGrid) ClientToScreen(Point TPoint) TPoint {
    return StringGrid_ClientToScreen(s.instance, Point)
}

// ClientToParent
func (s *TStringGrid) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return StringGrid_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// Dragging
func (s *TStringGrid) Dragging() bool {
    return StringGrid_Dragging(s.instance)
}

// HasParent
func (s *TStringGrid) HasParent() bool {
    return StringGrid_HasParent(s.instance)
}

// Hide
func (s *TStringGrid) Hide() {
    StringGrid_Hide(s.instance)
}

// Perform
func (s *TStringGrid) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return StringGrid_Perform(s.instance, Msg , WParam , LParam)
}

// Refresh
func (s *TStringGrid) Refresh() {
    StringGrid_Refresh(s.instance)
}

// ScreenToClient
func (s *TStringGrid) ScreenToClient(Point TPoint) TPoint {
    return StringGrid_ScreenToClient(s.instance, Point)
}

// ParentToClient
func (s *TStringGrid) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return StringGrid_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (s *TStringGrid) SendToBack() {
    StringGrid_SendToBack(s.instance)
}

// Show
func (s *TStringGrid) Show() {
    StringGrid_Show(s.instance)
}

// GetTextBuf
func (s *TStringGrid) GetTextBuf(Buffer string, BufSize int32) int32 {
    return StringGrid_GetTextBuf(s.instance, Buffer , BufSize)
}

// GetTextLen
func (s *TStringGrid) GetTextLen() int32 {
    return StringGrid_GetTextLen(s.instance)
}

// SetTextBuf
func (s *TStringGrid) SetTextBuf(Buffer string) {
    StringGrid_SetTextBuf(s.instance, Buffer)
}

// FindComponent
func (s *TStringGrid) FindComponent(AName string) *TComponent {
    return ComponentFromInst(StringGrid_FindComponent(s.instance, AName))
}

// GetNamePath
func (s *TStringGrid) GetNamePath() string {
    return StringGrid_GetNamePath(s.instance)
}

// Assign
func (s *TStringGrid) Assign(Source IObject) {
    StringGrid_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TStringGrid) DisposeOf() {
    StringGrid_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TStringGrid) ClassType() TClass {
    return StringGrid_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TStringGrid) ClassName() string {
    return StringGrid_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TStringGrid) InstanceSize() int32 {
    return StringGrid_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TStringGrid) InheritsFrom(AClass TClass) bool {
    return StringGrid_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TStringGrid) Equals(Obj IObject) bool {
    return StringGrid_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TStringGrid) GetHashCode() int32 {
    return StringGrid_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TStringGrid) ToString() string {
    return StringGrid_ToString(s.instance)
}

// Align
func (s *TStringGrid) Align() TAlign {
    return StringGrid_GetAlign(s.instance)
}

// SetAlign
func (s *TStringGrid) SetAlign(value TAlign) {
    StringGrid_SetAlign(s.instance, value)
}

// Anchors
func (s *TStringGrid) Anchors() TAnchors {
    return StringGrid_GetAnchors(s.instance)
}

// SetAnchors
func (s *TStringGrid) SetAnchors(value TAnchors) {
    StringGrid_SetAnchors(s.instance, value)
}

// BevelEdges
func (s *TStringGrid) BevelEdges() TBevelEdges {
    return StringGrid_GetBevelEdges(s.instance)
}

// SetBevelEdges
func (s *TStringGrid) SetBevelEdges(value TBevelEdges) {
    StringGrid_SetBevelEdges(s.instance, value)
}

// BevelInner
func (s *TStringGrid) BevelInner() TBevelCut {
    return StringGrid_GetBevelInner(s.instance)
}

// SetBevelInner
func (s *TStringGrid) SetBevelInner(value TBevelCut) {
    StringGrid_SetBevelInner(s.instance, value)
}

// BevelKind
func (s *TStringGrid) BevelKind() TBevelKind {
    return StringGrid_GetBevelKind(s.instance)
}

// SetBevelKind
func (s *TStringGrid) SetBevelKind(value TBevelKind) {
    StringGrid_SetBevelKind(s.instance, value)
}

// BevelOuter
func (s *TStringGrid) BevelOuter() TBevelCut {
    return StringGrid_GetBevelOuter(s.instance)
}

// SetBevelOuter
func (s *TStringGrid) SetBevelOuter(value TBevelCut) {
    StringGrid_SetBevelOuter(s.instance, value)
}

// BiDiMode
func (s *TStringGrid) BiDiMode() TBiDiMode {
    return StringGrid_GetBiDiMode(s.instance)
}

// SetBiDiMode
func (s *TStringGrid) SetBiDiMode(value TBiDiMode) {
    StringGrid_SetBiDiMode(s.instance, value)
}

// BorderStyle
func (s *TStringGrid) BorderStyle() TBorderStyle {
    return StringGrid_GetBorderStyle(s.instance)
}

// SetBorderStyle
func (s *TStringGrid) SetBorderStyle(value TBorderStyle) {
    StringGrid_SetBorderStyle(s.instance, value)
}

// Color
func (s *TStringGrid) Color() TColor {
    return StringGrid_GetColor(s.instance)
}

// SetColor
func (s *TStringGrid) SetColor(value TColor) {
    StringGrid_SetColor(s.instance, value)
}

// ColCount
func (s *TStringGrid) ColCount() int32 {
    return StringGrid_GetColCount(s.instance)
}

// SetColCount
func (s *TStringGrid) SetColCount(value int32) {
    StringGrid_SetColCount(s.instance, value)
}

// DefaultColWidth
func (s *TStringGrid) DefaultColWidth() int32 {
    return StringGrid_GetDefaultColWidth(s.instance)
}

// SetDefaultColWidth
func (s *TStringGrid) SetDefaultColWidth(value int32) {
    StringGrid_SetDefaultColWidth(s.instance, value)
}

// DefaultRowHeight
func (s *TStringGrid) DefaultRowHeight() int32 {
    return StringGrid_GetDefaultRowHeight(s.instance)
}

// SetDefaultRowHeight
func (s *TStringGrid) SetDefaultRowHeight(value int32) {
    StringGrid_SetDefaultRowHeight(s.instance, value)
}

// DefaultDrawing
func (s *TStringGrid) DefaultDrawing() bool {
    return StringGrid_GetDefaultDrawing(s.instance)
}

// SetDefaultDrawing
func (s *TStringGrid) SetDefaultDrawing(value bool) {
    StringGrid_SetDefaultDrawing(s.instance, value)
}

// DoubleBuffered
func (s *TStringGrid) DoubleBuffered() bool {
    return StringGrid_GetDoubleBuffered(s.instance)
}

// SetDoubleBuffered
func (s *TStringGrid) SetDoubleBuffered(value bool) {
    StringGrid_SetDoubleBuffered(s.instance, value)
}

// DragCursor
func (s *TStringGrid) DragCursor() TCursor {
    return StringGrid_GetDragCursor(s.instance)
}

// SetDragCursor
func (s *TStringGrid) SetDragCursor(value TCursor) {
    StringGrid_SetDragCursor(s.instance, value)
}

// DragKind
func (s *TStringGrid) DragKind() TDragKind {
    return StringGrid_GetDragKind(s.instance)
}

// SetDragKind
func (s *TStringGrid) SetDragKind(value TDragKind) {
    StringGrid_SetDragKind(s.instance, value)
}

// DragMode
func (s *TStringGrid) DragMode() TDragMode {
    return StringGrid_GetDragMode(s.instance)
}

// SetDragMode
func (s *TStringGrid) SetDragMode(value TDragMode) {
    StringGrid_SetDragMode(s.instance, value)
}

// DrawingStyle
func (s *TStringGrid) DrawingStyle() TGridDrawingStyle {
    return StringGrid_GetDrawingStyle(s.instance)
}

// SetDrawingStyle
func (s *TStringGrid) SetDrawingStyle(value TGridDrawingStyle) {
    StringGrid_SetDrawingStyle(s.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TStringGrid) Enabled() bool {
    return StringGrid_GetEnabled(s.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TStringGrid) SetEnabled(value bool) {
    StringGrid_SetEnabled(s.instance, value)
}

// FixedColor
func (s *TStringGrid) FixedColor() TColor {
    return StringGrid_GetFixedColor(s.instance)
}

// SetFixedColor
func (s *TStringGrid) SetFixedColor(value TColor) {
    StringGrid_SetFixedColor(s.instance, value)
}

// FixedCols
func (s *TStringGrid) FixedCols() int32 {
    return StringGrid_GetFixedCols(s.instance)
}

// SetFixedCols
func (s *TStringGrid) SetFixedCols(value int32) {
    StringGrid_SetFixedCols(s.instance, value)
}

// RowCount
func (s *TStringGrid) RowCount() int32 {
    return StringGrid_GetRowCount(s.instance)
}

// SetRowCount
func (s *TStringGrid) SetRowCount(value int32) {
    StringGrid_SetRowCount(s.instance, value)
}

// FixedRows
func (s *TStringGrid) FixedRows() int32 {
    return StringGrid_GetFixedRows(s.instance)
}

// SetFixedRows
func (s *TStringGrid) SetFixedRows(value int32) {
    StringGrid_SetFixedRows(s.instance, value)
}

// Font
func (s *TStringGrid) Font() *TFont {
    return FontFromInst(StringGrid_GetFont(s.instance))
}

// SetFont
func (s *TStringGrid) SetFont(value *TFont) {
    StringGrid_SetFont(s.instance, CheckPtr(value))
}

// GradientEndColor
func (s *TStringGrid) GradientEndColor() TColor {
    return StringGrid_GetGradientEndColor(s.instance)
}

// SetGradientEndColor
func (s *TStringGrid) SetGradientEndColor(value TColor) {
    StringGrid_SetGradientEndColor(s.instance, value)
}

// GradientStartColor
func (s *TStringGrid) GradientStartColor() TColor {
    return StringGrid_GetGradientStartColor(s.instance)
}

// SetGradientStartColor
func (s *TStringGrid) SetGradientStartColor(value TColor) {
    StringGrid_SetGradientStartColor(s.instance, value)
}

// GridLineWidth
func (s *TStringGrid) GridLineWidth() int32 {
    return StringGrid_GetGridLineWidth(s.instance)
}

// SetGridLineWidth
func (s *TStringGrid) SetGridLineWidth(value int32) {
    StringGrid_SetGridLineWidth(s.instance, value)
}

// Options
func (s *TStringGrid) Options() TGridOptions {
    return StringGrid_GetOptions(s.instance)
}

// SetOptions
func (s *TStringGrid) SetOptions(value TGridOptions) {
    StringGrid_SetOptions(s.instance, value)
}

// ParentColor
func (s *TStringGrid) ParentColor() bool {
    return StringGrid_GetParentColor(s.instance)
}

// SetParentColor
func (s *TStringGrid) SetParentColor(value bool) {
    StringGrid_SetParentColor(s.instance, value)
}

// ParentCtl3D
func (s *TStringGrid) ParentCtl3D() bool {
    return StringGrid_GetParentCtl3D(s.instance)
}

// SetParentCtl3D
func (s *TStringGrid) SetParentCtl3D(value bool) {
    StringGrid_SetParentCtl3D(s.instance, value)
}

// ParentDoubleBuffered
func (s *TStringGrid) ParentDoubleBuffered() bool {
    return StringGrid_GetParentDoubleBuffered(s.instance)
}

// SetParentDoubleBuffered
func (s *TStringGrid) SetParentDoubleBuffered(value bool) {
    StringGrid_SetParentDoubleBuffered(s.instance, value)
}

// ParentFont
func (s *TStringGrid) ParentFont() bool {
    return StringGrid_GetParentFont(s.instance)
}

// SetParentFont
func (s *TStringGrid) SetParentFont(value bool) {
    StringGrid_SetParentFont(s.instance, value)
}

// ParentShowHint
func (s *TStringGrid) ParentShowHint() bool {
    return StringGrid_GetParentShowHint(s.instance)
}

// SetParentShowHint
func (s *TStringGrid) SetParentShowHint(value bool) {
    StringGrid_SetParentShowHint(s.instance, value)
}

// PopupMenu
func (s *TStringGrid) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(StringGrid_GetPopupMenu(s.instance))
}

// SetPopupMenu
func (s *TStringGrid) SetPopupMenu(value IComponent) {
    StringGrid_SetPopupMenu(s.instance, CheckPtr(value))
}

// ScrollBars
func (s *TStringGrid) ScrollBars() TScrollStyle {
    return StringGrid_GetScrollBars(s.instance)
}

// SetScrollBars
func (s *TStringGrid) SetScrollBars(value TScrollStyle) {
    StringGrid_SetScrollBars(s.instance, value)
}

// ShowHint
func (s *TStringGrid) ShowHint() bool {
    return StringGrid_GetShowHint(s.instance)
}

// SetShowHint
func (s *TStringGrid) SetShowHint(value bool) {
    StringGrid_SetShowHint(s.instance, value)
}

// TabOrder
func (s *TStringGrid) TabOrder() TTabOrder {
    return StringGrid_GetTabOrder(s.instance)
}

// SetTabOrder
func (s *TStringGrid) SetTabOrder(value TTabOrder) {
    StringGrid_SetTabOrder(s.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TStringGrid) Visible() bool {
    return StringGrid_GetVisible(s.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TStringGrid) SetVisible(value bool) {
    StringGrid_SetVisible(s.instance, value)
}

// StyleElements
func (s *TStringGrid) StyleElements() TStyleElements {
    return StringGrid_GetStyleElements(s.instance)
}

// SetStyleElements
func (s *TStringGrid) SetStyleElements(value TStyleElements) {
    StringGrid_SetStyleElements(s.instance, value)
}

// VisibleColCount
func (s *TStringGrid) VisibleColCount() int32 {
    return StringGrid_GetVisibleColCount(s.instance)
}

// VisibleRowCount
func (s *TStringGrid) VisibleRowCount() int32 {
    return StringGrid_GetVisibleRowCount(s.instance)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (s *TStringGrid) SetOnClick(fn TNotifyEvent) {
    StringGrid_SetOnClick(s.instance, fn)
}

// SetOnColumnMoved
func (s *TStringGrid) SetOnColumnMoved(fn TMovedEvent) {
    StringGrid_SetOnColumnMoved(s.instance, fn)
}

// SetOnContextPopup
func (s *TStringGrid) SetOnContextPopup(fn TContextPopupEvent) {
    StringGrid_SetOnContextPopup(s.instance, fn)
}

// SetOnDblClick
func (s *TStringGrid) SetOnDblClick(fn TNotifyEvent) {
    StringGrid_SetOnDblClick(s.instance, fn)
}

// SetOnDragDrop
func (s *TStringGrid) SetOnDragDrop(fn TDragDropEvent) {
    StringGrid_SetOnDragDrop(s.instance, fn)
}

// SetOnDragOver
func (s *TStringGrid) SetOnDragOver(fn TDragOverEvent) {
    StringGrid_SetOnDragOver(s.instance, fn)
}

// SetOnDrawCell
func (s *TStringGrid) SetOnDrawCell(fn TDrawCellEvent) {
    StringGrid_SetOnDrawCell(s.instance, fn)
}

// SetOnEndDock
func (s *TStringGrid) SetOnEndDock(fn TEndDragEvent) {
    StringGrid_SetOnEndDock(s.instance, fn)
}

// SetOnEndDrag
func (s *TStringGrid) SetOnEndDrag(fn TEndDragEvent) {
    StringGrid_SetOnEndDrag(s.instance, fn)
}

// SetOnEnter
func (s *TStringGrid) SetOnEnter(fn TNotifyEvent) {
    StringGrid_SetOnEnter(s.instance, fn)
}

// SetOnExit
func (s *TStringGrid) SetOnExit(fn TNotifyEvent) {
    StringGrid_SetOnExit(s.instance, fn)
}

// SetOnFixedCellClick
func (s *TStringGrid) SetOnFixedCellClick(fn TFixedCellClickEvent) {
    StringGrid_SetOnFixedCellClick(s.instance, fn)
}

// SetOnGetEditMask
func (s *TStringGrid) SetOnGetEditMask(fn TGetEditEvent) {
    StringGrid_SetOnGetEditMask(s.instance, fn)
}

// SetOnGetEditText
func (s *TStringGrid) SetOnGetEditText(fn TGetEditEvent) {
    StringGrid_SetOnGetEditText(s.instance, fn)
}

// SetOnKeyDown
func (s *TStringGrid) SetOnKeyDown(fn TKeyEvent) {
    StringGrid_SetOnKeyDown(s.instance, fn)
}

// SetOnKeyPress
func (s *TStringGrid) SetOnKeyPress(fn TKeyPressEvent) {
    StringGrid_SetOnKeyPress(s.instance, fn)
}

// SetOnKeyUp
func (s *TStringGrid) SetOnKeyUp(fn TKeyEvent) {
    StringGrid_SetOnKeyUp(s.instance, fn)
}

// SetOnMouseDown
func (s *TStringGrid) SetOnMouseDown(fn TMouseEvent) {
    StringGrid_SetOnMouseDown(s.instance, fn)
}

// SetOnMouseEnter
func (s *TStringGrid) SetOnMouseEnter(fn TNotifyEvent) {
    StringGrid_SetOnMouseEnter(s.instance, fn)
}

// SetOnMouseLeave
func (s *TStringGrid) SetOnMouseLeave(fn TNotifyEvent) {
    StringGrid_SetOnMouseLeave(s.instance, fn)
}

// SetOnMouseMove
func (s *TStringGrid) SetOnMouseMove(fn TMouseMoveEvent) {
    StringGrid_SetOnMouseMove(s.instance, fn)
}

// SetOnMouseUp
func (s *TStringGrid) SetOnMouseUp(fn TMouseEvent) {
    StringGrid_SetOnMouseUp(s.instance, fn)
}

// SetOnMouseWheelDown
func (s *TStringGrid) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    StringGrid_SetOnMouseWheelDown(s.instance, fn)
}

// SetOnMouseWheelUp
func (s *TStringGrid) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    StringGrid_SetOnMouseWheelUp(s.instance, fn)
}

// SetOnRowMoved
func (s *TStringGrid) SetOnRowMoved(fn TMovedEvent) {
    StringGrid_SetOnRowMoved(s.instance, fn)
}

// SetOnSelectCell
func (s *TStringGrid) SetOnSelectCell(fn TSelectCellEvent) {
    StringGrid_SetOnSelectCell(s.instance, fn)
}

// SetOnSetEditText
func (s *TStringGrid) SetOnSetEditText(fn TSetEditEvent) {
    StringGrid_SetOnSetEditText(s.instance, fn)
}

// SetOnStartDock
func (s *TStringGrid) SetOnStartDock(fn TStartDockEvent) {
    StringGrid_SetOnStartDock(s.instance, fn)
}

// SetOnTopLeftChanged
func (s *TStringGrid) SetOnTopLeftChanged(fn TNotifyEvent) {
    StringGrid_SetOnTopLeftChanged(s.instance, fn)
}

// Canvas
func (s *TStringGrid) Canvas() *TCanvas {
    return CanvasFromInst(StringGrid_GetCanvas(s.instance))
}

// Col
func (s *TStringGrid) Col() int32 {
    return StringGrid_GetCol(s.instance)
}

// SetCol
func (s *TStringGrid) SetCol(value int32) {
    StringGrid_SetCol(s.instance, value)
}

// EditorMode
func (s *TStringGrid) EditorMode() bool {
    return StringGrid_GetEditorMode(s.instance)
}

// SetEditorMode
func (s *TStringGrid) SetEditorMode(value bool) {
    StringGrid_SetEditorMode(s.instance, value)
}

// GridHeight
func (s *TStringGrid) GridHeight() int32 {
    return StringGrid_GetGridHeight(s.instance)
}

// GridWidth
func (s *TStringGrid) GridWidth() int32 {
    return StringGrid_GetGridWidth(s.instance)
}

// LeftCol
func (s *TStringGrid) LeftCol() int32 {
    return StringGrid_GetLeftCol(s.instance)
}

// SetLeftCol
func (s *TStringGrid) SetLeftCol(value int32) {
    StringGrid_SetLeftCol(s.instance, value)
}

// Selection
func (s *TStringGrid) Selection() TGridRect {
    return StringGrid_GetSelection(s.instance)
}

// SetSelection
func (s *TStringGrid) SetSelection(value TGridRect) {
    StringGrid_SetSelection(s.instance, value)
}

// Row
func (s *TStringGrid) Row() int32 {
    return StringGrid_GetRow(s.instance)
}

// SetRow
func (s *TStringGrid) SetRow(value int32) {
    StringGrid_SetRow(s.instance, value)
}

// TopRow
func (s *TStringGrid) TopRow() int32 {
    return StringGrid_GetTopRow(s.instance)
}

// SetTopRow
func (s *TStringGrid) SetTopRow(value int32) {
    StringGrid_SetTopRow(s.instance, value)
}

// TabStop
func (s *TStringGrid) TabStop() bool {
    return StringGrid_GetTabStop(s.instance)
}

// SetTabStop
func (s *TStringGrid) SetTabStop(value bool) {
    StringGrid_SetTabStop(s.instance, value)
}

// DockClientCount
func (s *TStringGrid) DockClientCount() int32 {
    return StringGrid_GetDockClientCount(s.instance)
}

// DockSite
func (s *TStringGrid) DockSite() bool {
    return StringGrid_GetDockSite(s.instance)
}

// SetDockSite
func (s *TStringGrid) SetDockSite(value bool) {
    StringGrid_SetDockSite(s.instance, value)
}

// AlignDisabled
func (s *TStringGrid) AlignDisabled() bool {
    return StringGrid_GetAlignDisabled(s.instance)
}

// MouseInClient
func (s *TStringGrid) MouseInClient() bool {
    return StringGrid_GetMouseInClient(s.instance)
}

// VisibleDockClientCount
func (s *TStringGrid) VisibleDockClientCount() int32 {
    return StringGrid_GetVisibleDockClientCount(s.instance)
}

// Brush
func (s *TStringGrid) Brush() *TBrush {
    return BrushFromInst(StringGrid_GetBrush(s.instance))
}

// ControlCount
func (s *TStringGrid) ControlCount() int32 {
    return StringGrid_GetControlCount(s.instance)
}

// Handle
func (s *TStringGrid) Handle() HWND {
    return StringGrid_GetHandle(s.instance)
}

// ParentWindow
func (s *TStringGrid) ParentWindow() HWND {
    return StringGrid_GetParentWindow(s.instance)
}

// SetParentWindow
func (s *TStringGrid) SetParentWindow(value HWND) {
    StringGrid_SetParentWindow(s.instance, value)
}

// UseDockManager
func (s *TStringGrid) UseDockManager() bool {
    return StringGrid_GetUseDockManager(s.instance)
}

// SetUseDockManager
func (s *TStringGrid) SetUseDockManager(value bool) {
    StringGrid_SetUseDockManager(s.instance, value)
}

// Action
func (s *TStringGrid) Action() *TAction {
    return ActionFromInst(StringGrid_GetAction(s.instance))
}

// SetAction
func (s *TStringGrid) SetAction(value IComponent) {
    StringGrid_SetAction(s.instance, CheckPtr(value))
}

// BoundsRect
func (s *TStringGrid) BoundsRect() TRect {
    return StringGrid_GetBoundsRect(s.instance)
}

// SetBoundsRect
func (s *TStringGrid) SetBoundsRect(value TRect) {
    StringGrid_SetBoundsRect(s.instance, value)
}

// ClientHeight
func (s *TStringGrid) ClientHeight() int32 {
    return StringGrid_GetClientHeight(s.instance)
}

// SetClientHeight
func (s *TStringGrid) SetClientHeight(value int32) {
    StringGrid_SetClientHeight(s.instance, value)
}

// ClientOrigin
func (s *TStringGrid) ClientOrigin() TPoint {
    return StringGrid_GetClientOrigin(s.instance)
}

// ClientRect
func (s *TStringGrid) ClientRect() TRect {
    return StringGrid_GetClientRect(s.instance)
}

// ClientWidth
func (s *TStringGrid) ClientWidth() int32 {
    return StringGrid_GetClientWidth(s.instance)
}

// SetClientWidth
func (s *TStringGrid) SetClientWidth(value int32) {
    StringGrid_SetClientWidth(s.instance, value)
}

// ControlState
func (s *TStringGrid) ControlState() TControlState {
    return StringGrid_GetControlState(s.instance)
}

// SetControlState
func (s *TStringGrid) SetControlState(value TControlState) {
    StringGrid_SetControlState(s.instance, value)
}

// ControlStyle
func (s *TStringGrid) ControlStyle() TControlStyle {
    return StringGrid_GetControlStyle(s.instance)
}

// SetControlStyle
func (s *TStringGrid) SetControlStyle(value TControlStyle) {
    StringGrid_SetControlStyle(s.instance, value)
}

// ExplicitLeft
func (s *TStringGrid) ExplicitLeft() int32 {
    return StringGrid_GetExplicitLeft(s.instance)
}

// ExplicitTop
func (s *TStringGrid) ExplicitTop() int32 {
    return StringGrid_GetExplicitTop(s.instance)
}

// ExplicitWidth
func (s *TStringGrid) ExplicitWidth() int32 {
    return StringGrid_GetExplicitWidth(s.instance)
}

// ExplicitHeight
func (s *TStringGrid) ExplicitHeight() int32 {
    return StringGrid_GetExplicitHeight(s.instance)
}

// Floating
func (s *TStringGrid) Floating() bool {
    return StringGrid_GetFloating(s.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TStringGrid) Parent() *TWinControl {
    return WinControlFromInst(StringGrid_GetParent(s.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TStringGrid) SetParent(value IWinControl) {
    StringGrid_SetParent(s.instance, CheckPtr(value))
}

// AlignWithMargins
func (s *TStringGrid) AlignWithMargins() bool {
    return StringGrid_GetAlignWithMargins(s.instance)
}

// SetAlignWithMargins
func (s *TStringGrid) SetAlignWithMargins(value bool) {
    StringGrid_SetAlignWithMargins(s.instance, value)
}

// Left
func (s *TStringGrid) Left() int32 {
    return StringGrid_GetLeft(s.instance)
}

// SetLeft
func (s *TStringGrid) SetLeft(value int32) {
    StringGrid_SetLeft(s.instance, value)
}

// Top
func (s *TStringGrid) Top() int32 {
    return StringGrid_GetTop(s.instance)
}

// SetTop
func (s *TStringGrid) SetTop(value int32) {
    StringGrid_SetTop(s.instance, value)
}

// Width
func (s *TStringGrid) Width() int32 {
    return StringGrid_GetWidth(s.instance)
}

// SetWidth
func (s *TStringGrid) SetWidth(value int32) {
    StringGrid_SetWidth(s.instance, value)
}

// Height
func (s *TStringGrid) Height() int32 {
    return StringGrid_GetHeight(s.instance)
}

// SetHeight
func (s *TStringGrid) SetHeight(value int32) {
    StringGrid_SetHeight(s.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TStringGrid) Cursor() TCursor {
    return StringGrid_GetCursor(s.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TStringGrid) SetCursor(value TCursor) {
    StringGrid_SetCursor(s.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (s *TStringGrid) Hint() string {
    return StringGrid_GetHint(s.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (s *TStringGrid) SetHint(value string) {
    StringGrid_SetHint(s.instance, value)
}

// Margins
func (s *TStringGrid) Margins() *TMargins {
    return MarginsFromInst(StringGrid_GetMargins(s.instance))
}

// SetMargins
func (s *TStringGrid) SetMargins(value *TMargins) {
    StringGrid_SetMargins(s.instance, CheckPtr(value))
}

// CustomHint
func (s *TStringGrid) CustomHint() *TCustomHint {
    return CustomHintFromInst(StringGrid_GetCustomHint(s.instance))
}

// SetCustomHint
func (s *TStringGrid) SetCustomHint(value IComponent) {
    StringGrid_SetCustomHint(s.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TStringGrid) ComponentCount() int32 {
    return StringGrid_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TStringGrid) ComponentIndex() int32 {
    return StringGrid_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TStringGrid) SetComponentIndex(value int32) {
    StringGrid_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TStringGrid) Owner() *TComponent {
    return ComponentFromInst(StringGrid_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TStringGrid) Name() string {
    return StringGrid_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TStringGrid) SetName(value string) {
    StringGrid_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TStringGrid) Tag() int {
    return StringGrid_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TStringGrid) SetTag(value int) {
    StringGrid_SetTag(s.instance, value)
}

// Cells
func (s *TStringGrid) Cells(ACol int32, ARow int32) string {
    return StringGrid_GetCells(s.instance, ACol, ARow)
}

// Cells
func (s *TStringGrid) SetCells(ACol int32, ARow int32, value string) {
    StringGrid_SetCells(s.instance, ACol, ARow, value)
}

// Cols
func (s *TStringGrid) Cols(Index int32) *TStrings {
    return StringsFromInst(StringGrid_GetCols(s.instance, Index))
}

// Cols
func (s *TStringGrid) SetCols(Index int32, value IObject) {
    StringGrid_SetCols(s.instance, Index, CheckPtr(value))
}

// Objects
func (s *TStringGrid) Objects(ACol int32, ARow int32) *TObject {
    return ObjectFromInst(StringGrid_GetObjects(s.instance, ACol, ARow))
}

// Objects
func (s *TStringGrid) SetObjects(ACol int32, ARow int32, value IObject) {
    StringGrid_SetObjects(s.instance, ACol, ARow, CheckPtr(value))
}

// Rows
func (s *TStringGrid) Rows(Index int32) *TStrings {
    return StringsFromInst(StringGrid_GetRows(s.instance, Index))
}

// Rows
func (s *TStringGrid) SetRows(Index int32, value IObject) {
    StringGrid_SetRows(s.instance, Index, CheckPtr(value))
}

// ColWidths
func (s *TStringGrid) ColWidths(Index int32) int32 {
    return StringGrid_GetColWidths(s.instance, Index)
}

// ColWidths
func (s *TStringGrid) SetColWidths(Index int32, value int32) {
    StringGrid_SetColWidths(s.instance, Index, value)
}

// RowHeights
func (s *TStringGrid) RowHeights(Index int32) int32 {
    return StringGrid_GetRowHeights(s.instance, Index)
}

// RowHeights
func (s *TStringGrid) SetRowHeights(Index int32, value int32) {
    StringGrid_SetRowHeights(s.instance, Index, value)
}

// TabStops
func (s *TStringGrid) TabStops(Index int32) bool {
    return StringGrid_GetTabStops(s.instance, Index)
}

// TabStops
func (s *TStringGrid) SetTabStops(Index int32, value bool) {
    StringGrid_SetTabStops(s.instance, Index, value)
}

// DockClients
func (s *TStringGrid) DockClients(Index int32) *TControl {
    return ControlFromInst(StringGrid_GetDockClients(s.instance, Index))
}

// Controls
func (s *TStringGrid) Controls(Index int32) *TControl {
    return ControlFromInst(StringGrid_GetControls(s.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TStringGrid) Components(AIndex int32) *TComponent {
    return ComponentFromInst(StringGrid_GetComponents(s.instance, AIndex))
}

