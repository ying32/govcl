
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

type TDrawGrid struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewDrawGrid
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewDrawGrid(owner IComponent) *TDrawGrid {
    d := new(TDrawGrid)
    d.instance = DrawGrid_Create(CheckPtr(owner))
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// DrawGridFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func DrawGridFromInst(inst uintptr) *TDrawGrid {
    d := new(TDrawGrid)
    d.instance = inst
    d.ptr = unsafe.Pointer(inst)
    return d
}

// DrawGridFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func DrawGridFromObj(obj IObject) *TDrawGrid {
    d := new(TDrawGrid)
    d.instance = CheckPtr(obj)
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// DrawGridFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func DrawGridFromUnsafePointer(ptr unsafe.Pointer) *TDrawGrid {
    d := new(TDrawGrid)
    d.instance = uintptr(ptr)
    d.ptr = ptr
    return d
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (d *TDrawGrid) Free() {
    if d.instance != 0 {
        DrawGrid_Free(d.instance)
        d.instance = 0
        d.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (d *TDrawGrid) Instance() uintptr {
    return d.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (d *TDrawGrid) UnsafeAddr() unsafe.Pointer {
    return d.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (d *TDrawGrid) IsValid() bool {
    return d.instance != 0
}

// TDrawGridClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TDrawGridClass() TClass {
    return DrawGrid_StaticClassType()
}

// CellRect
func (d *TDrawGrid) CellRect(ACol int32, ARow int32) TRect {
    return DrawGrid_CellRect(d.instance, ACol , ARow)
}

// MouseToCell
func (d *TDrawGrid) MouseToCell(X int32, Y int32, ACol *int32, ARow *int32) {
    DrawGrid_MouseToCell(d.instance, X , Y , ACol , ARow)
}

// MouseCoord
func (d *TDrawGrid) MouseCoord(X int32, Y int32) TGridCoord {
    return DrawGrid_MouseCoord(d.instance, X , Y)
}

// CanFocus
func (d *TDrawGrid) CanFocus() bool {
    return DrawGrid_CanFocus(d.instance)
}

// ContainsControl
func (d *TDrawGrid) ContainsControl(Control IControl) bool {
    return DrawGrid_ContainsControl(d.instance, CheckPtr(Control))
}

// ControlAtPos
func (d *TDrawGrid) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(DrawGrid_ControlAtPos(d.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (d *TDrawGrid) DisableAlign() {
    DrawGrid_DisableAlign(d.instance)
}

// EnableAlign
func (d *TDrawGrid) EnableAlign() {
    DrawGrid_EnableAlign(d.instance)
}

// FindChildControl
func (d *TDrawGrid) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(DrawGrid_FindChildControl(d.instance, ControlName))
}

// FlipChildren
func (d *TDrawGrid) FlipChildren(AllLevels bool) {
    DrawGrid_FlipChildren(d.instance, AllLevels)
}

// Focused
func (d *TDrawGrid) Focused() bool {
    return DrawGrid_Focused(d.instance)
}

// HandleAllocated
func (d *TDrawGrid) HandleAllocated() bool {
    return DrawGrid_HandleAllocated(d.instance)
}

// InsertControl
func (d *TDrawGrid) InsertControl(AControl IControl) {
    DrawGrid_InsertControl(d.instance, CheckPtr(AControl))
}

// Invalidate
func (d *TDrawGrid) Invalidate() {
    DrawGrid_Invalidate(d.instance)
}

// PaintTo
func (d *TDrawGrid) PaintTo(DC HDC, X int32, Y int32) {
    DrawGrid_PaintTo(d.instance, DC , X , Y)
}

// RemoveControl
func (d *TDrawGrid) RemoveControl(AControl IControl) {
    DrawGrid_RemoveControl(d.instance, CheckPtr(AControl))
}

// Realign
func (d *TDrawGrid) Realign() {
    DrawGrid_Realign(d.instance)
}

// Repaint
func (d *TDrawGrid) Repaint() {
    DrawGrid_Repaint(d.instance)
}

// ScaleBy
func (d *TDrawGrid) ScaleBy(M int32, D int32) {
    DrawGrid_ScaleBy(d.instance, M , D)
}

// ScrollBy
func (d *TDrawGrid) ScrollBy(DeltaX int32, DeltaY int32) {
    DrawGrid_ScrollBy(d.instance, DeltaX , DeltaY)
}

// SetBounds
func (d *TDrawGrid) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    DrawGrid_SetBounds(d.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (d *TDrawGrid) SetFocus() {
    DrawGrid_SetFocus(d.instance)
}

// Update
func (d *TDrawGrid) Update() {
    DrawGrid_Update(d.instance)
}

// UpdateControlState
func (d *TDrawGrid) UpdateControlState() {
    DrawGrid_UpdateControlState(d.instance)
}

// BringToFront
func (d *TDrawGrid) BringToFront() {
    DrawGrid_BringToFront(d.instance)
}

// ClientToScreen
func (d *TDrawGrid) ClientToScreen(Point TPoint) TPoint {
    return DrawGrid_ClientToScreen(d.instance, Point)
}

// ClientToParent
func (d *TDrawGrid) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return DrawGrid_ClientToParent(d.instance, Point , CheckPtr(AParent))
}

// Dragging
func (d *TDrawGrid) Dragging() bool {
    return DrawGrid_Dragging(d.instance)
}

// HasParent
func (d *TDrawGrid) HasParent() bool {
    return DrawGrid_HasParent(d.instance)
}

// Hide
func (d *TDrawGrid) Hide() {
    DrawGrid_Hide(d.instance)
}

// Perform
func (d *TDrawGrid) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return DrawGrid_Perform(d.instance, Msg , WParam , LParam)
}

// Refresh
func (d *TDrawGrid) Refresh() {
    DrawGrid_Refresh(d.instance)
}

// ScreenToClient
func (d *TDrawGrid) ScreenToClient(Point TPoint) TPoint {
    return DrawGrid_ScreenToClient(d.instance, Point)
}

// ParentToClient
func (d *TDrawGrid) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return DrawGrid_ParentToClient(d.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (d *TDrawGrid) SendToBack() {
    DrawGrid_SendToBack(d.instance)
}

// Show
func (d *TDrawGrid) Show() {
    DrawGrid_Show(d.instance)
}

// GetTextBuf
func (d *TDrawGrid) GetTextBuf(Buffer string, BufSize int32) int32 {
    return DrawGrid_GetTextBuf(d.instance, Buffer , BufSize)
}

// GetTextLen
func (d *TDrawGrid) GetTextLen() int32 {
    return DrawGrid_GetTextLen(d.instance)
}

// SetTextBuf
func (d *TDrawGrid) SetTextBuf(Buffer string) {
    DrawGrid_SetTextBuf(d.instance, Buffer)
}

// FindComponent
func (d *TDrawGrid) FindComponent(AName string) *TComponent {
    return ComponentFromInst(DrawGrid_FindComponent(d.instance, AName))
}

// GetNamePath
func (d *TDrawGrid) GetNamePath() string {
    return DrawGrid_GetNamePath(d.instance)
}

// Assign
func (d *TDrawGrid) Assign(Source IObject) {
    DrawGrid_Assign(d.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (d *TDrawGrid) DisposeOf() {
    DrawGrid_DisposeOf(d.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (d *TDrawGrid) ClassType() TClass {
    return DrawGrid_ClassType(d.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (d *TDrawGrid) ClassName() string {
    return DrawGrid_ClassName(d.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (d *TDrawGrid) InstanceSize() int32 {
    return DrawGrid_InstanceSize(d.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (d *TDrawGrid) InheritsFrom(AClass TClass) bool {
    return DrawGrid_InheritsFrom(d.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (d *TDrawGrid) Equals(Obj IObject) bool {
    return DrawGrid_Equals(d.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (d *TDrawGrid) GetHashCode() int32 {
    return DrawGrid_GetHashCode(d.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (d *TDrawGrid) ToString() string {
    return DrawGrid_ToString(d.instance)
}

// Align
func (d *TDrawGrid) Align() TAlign {
    return DrawGrid_GetAlign(d.instance)
}

// SetAlign
func (d *TDrawGrid) SetAlign(value TAlign) {
    DrawGrid_SetAlign(d.instance, value)
}

// Anchors
func (d *TDrawGrid) Anchors() TAnchors {
    return DrawGrid_GetAnchors(d.instance)
}

// SetAnchors
func (d *TDrawGrid) SetAnchors(value TAnchors) {
    DrawGrid_SetAnchors(d.instance, value)
}

// BevelEdges
func (d *TDrawGrid) BevelEdges() TBevelEdges {
    return DrawGrid_GetBevelEdges(d.instance)
}

// SetBevelEdges
func (d *TDrawGrid) SetBevelEdges(value TBevelEdges) {
    DrawGrid_SetBevelEdges(d.instance, value)
}

// BevelInner
func (d *TDrawGrid) BevelInner() TBevelCut {
    return DrawGrid_GetBevelInner(d.instance)
}

// SetBevelInner
func (d *TDrawGrid) SetBevelInner(value TBevelCut) {
    DrawGrid_SetBevelInner(d.instance, value)
}

// BevelKind
func (d *TDrawGrid) BevelKind() TBevelKind {
    return DrawGrid_GetBevelKind(d.instance)
}

// SetBevelKind
func (d *TDrawGrid) SetBevelKind(value TBevelKind) {
    DrawGrid_SetBevelKind(d.instance, value)
}

// BevelOuter
func (d *TDrawGrid) BevelOuter() TBevelCut {
    return DrawGrid_GetBevelOuter(d.instance)
}

// SetBevelOuter
func (d *TDrawGrid) SetBevelOuter(value TBevelCut) {
    DrawGrid_SetBevelOuter(d.instance, value)
}

// BiDiMode
func (d *TDrawGrid) BiDiMode() TBiDiMode {
    return DrawGrid_GetBiDiMode(d.instance)
}

// SetBiDiMode
func (d *TDrawGrid) SetBiDiMode(value TBiDiMode) {
    DrawGrid_SetBiDiMode(d.instance, value)
}

// BorderStyle
func (d *TDrawGrid) BorderStyle() TBorderStyle {
    return DrawGrid_GetBorderStyle(d.instance)
}

// SetBorderStyle
func (d *TDrawGrid) SetBorderStyle(value TBorderStyle) {
    DrawGrid_SetBorderStyle(d.instance, value)
}

// Color
func (d *TDrawGrid) Color() TColor {
    return DrawGrid_GetColor(d.instance)
}

// SetColor
func (d *TDrawGrid) SetColor(value TColor) {
    DrawGrid_SetColor(d.instance, value)
}

// ColCount
func (d *TDrawGrid) ColCount() int32 {
    return DrawGrid_GetColCount(d.instance)
}

// SetColCount
func (d *TDrawGrid) SetColCount(value int32) {
    DrawGrid_SetColCount(d.instance, value)
}

// DefaultColWidth
func (d *TDrawGrid) DefaultColWidth() int32 {
    return DrawGrid_GetDefaultColWidth(d.instance)
}

// SetDefaultColWidth
func (d *TDrawGrid) SetDefaultColWidth(value int32) {
    DrawGrid_SetDefaultColWidth(d.instance, value)
}

// DefaultRowHeight
func (d *TDrawGrid) DefaultRowHeight() int32 {
    return DrawGrid_GetDefaultRowHeight(d.instance)
}

// SetDefaultRowHeight
func (d *TDrawGrid) SetDefaultRowHeight(value int32) {
    DrawGrid_SetDefaultRowHeight(d.instance, value)
}

// DefaultDrawing
func (d *TDrawGrid) DefaultDrawing() bool {
    return DrawGrid_GetDefaultDrawing(d.instance)
}

// SetDefaultDrawing
func (d *TDrawGrid) SetDefaultDrawing(value bool) {
    DrawGrid_SetDefaultDrawing(d.instance, value)
}

// DoubleBuffered
func (d *TDrawGrid) DoubleBuffered() bool {
    return DrawGrid_GetDoubleBuffered(d.instance)
}

// SetDoubleBuffered
func (d *TDrawGrid) SetDoubleBuffered(value bool) {
    DrawGrid_SetDoubleBuffered(d.instance, value)
}

// DragCursor
func (d *TDrawGrid) DragCursor() TCursor {
    return DrawGrid_GetDragCursor(d.instance)
}

// SetDragCursor
func (d *TDrawGrid) SetDragCursor(value TCursor) {
    DrawGrid_SetDragCursor(d.instance, value)
}

// DragKind
func (d *TDrawGrid) DragKind() TDragKind {
    return DrawGrid_GetDragKind(d.instance)
}

// SetDragKind
func (d *TDrawGrid) SetDragKind(value TDragKind) {
    DrawGrid_SetDragKind(d.instance, value)
}

// DragMode
func (d *TDrawGrid) DragMode() TDragMode {
    return DrawGrid_GetDragMode(d.instance)
}

// SetDragMode
func (d *TDrawGrid) SetDragMode(value TDragMode) {
    DrawGrid_SetDragMode(d.instance, value)
}

// DrawingStyle
func (d *TDrawGrid) DrawingStyle() TGridDrawingStyle {
    return DrawGrid_GetDrawingStyle(d.instance)
}

// SetDrawingStyle
func (d *TDrawGrid) SetDrawingStyle(value TGridDrawingStyle) {
    DrawGrid_SetDrawingStyle(d.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (d *TDrawGrid) Enabled() bool {
    return DrawGrid_GetEnabled(d.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (d *TDrawGrid) SetEnabled(value bool) {
    DrawGrid_SetEnabled(d.instance, value)
}

// FixedColor
func (d *TDrawGrid) FixedColor() TColor {
    return DrawGrid_GetFixedColor(d.instance)
}

// SetFixedColor
func (d *TDrawGrid) SetFixedColor(value TColor) {
    DrawGrid_SetFixedColor(d.instance, value)
}

// FixedCols
func (d *TDrawGrid) FixedCols() int32 {
    return DrawGrid_GetFixedCols(d.instance)
}

// SetFixedCols
func (d *TDrawGrid) SetFixedCols(value int32) {
    DrawGrid_SetFixedCols(d.instance, value)
}

// RowCount
func (d *TDrawGrid) RowCount() int32 {
    return DrawGrid_GetRowCount(d.instance)
}

// SetRowCount
func (d *TDrawGrid) SetRowCount(value int32) {
    DrawGrid_SetRowCount(d.instance, value)
}

// FixedRows
func (d *TDrawGrid) FixedRows() int32 {
    return DrawGrid_GetFixedRows(d.instance)
}

// SetFixedRows
func (d *TDrawGrid) SetFixedRows(value int32) {
    DrawGrid_SetFixedRows(d.instance, value)
}

// Font
func (d *TDrawGrid) Font() *TFont {
    return FontFromInst(DrawGrid_GetFont(d.instance))
}

// SetFont
func (d *TDrawGrid) SetFont(value *TFont) {
    DrawGrid_SetFont(d.instance, CheckPtr(value))
}

// GradientEndColor
func (d *TDrawGrid) GradientEndColor() TColor {
    return DrawGrid_GetGradientEndColor(d.instance)
}

// SetGradientEndColor
func (d *TDrawGrid) SetGradientEndColor(value TColor) {
    DrawGrid_SetGradientEndColor(d.instance, value)
}

// GradientStartColor
func (d *TDrawGrid) GradientStartColor() TColor {
    return DrawGrid_GetGradientStartColor(d.instance)
}

// SetGradientStartColor
func (d *TDrawGrid) SetGradientStartColor(value TColor) {
    DrawGrid_SetGradientStartColor(d.instance, value)
}

// GridLineWidth
func (d *TDrawGrid) GridLineWidth() int32 {
    return DrawGrid_GetGridLineWidth(d.instance)
}

// SetGridLineWidth
func (d *TDrawGrid) SetGridLineWidth(value int32) {
    DrawGrid_SetGridLineWidth(d.instance, value)
}

// Options
func (d *TDrawGrid) Options() TGridOptions {
    return DrawGrid_GetOptions(d.instance)
}

// SetOptions
func (d *TDrawGrid) SetOptions(value TGridOptions) {
    DrawGrid_SetOptions(d.instance, value)
}

// ParentColor
func (d *TDrawGrid) ParentColor() bool {
    return DrawGrid_GetParentColor(d.instance)
}

// SetParentColor
func (d *TDrawGrid) SetParentColor(value bool) {
    DrawGrid_SetParentColor(d.instance, value)
}

// ParentCtl3D
func (d *TDrawGrid) ParentCtl3D() bool {
    return DrawGrid_GetParentCtl3D(d.instance)
}

// SetParentCtl3D
func (d *TDrawGrid) SetParentCtl3D(value bool) {
    DrawGrid_SetParentCtl3D(d.instance, value)
}

// ParentDoubleBuffered
func (d *TDrawGrid) ParentDoubleBuffered() bool {
    return DrawGrid_GetParentDoubleBuffered(d.instance)
}

// SetParentDoubleBuffered
func (d *TDrawGrid) SetParentDoubleBuffered(value bool) {
    DrawGrid_SetParentDoubleBuffered(d.instance, value)
}

// ParentFont
func (d *TDrawGrid) ParentFont() bool {
    return DrawGrid_GetParentFont(d.instance)
}

// SetParentFont
func (d *TDrawGrid) SetParentFont(value bool) {
    DrawGrid_SetParentFont(d.instance, value)
}

// ParentShowHint
func (d *TDrawGrid) ParentShowHint() bool {
    return DrawGrid_GetParentShowHint(d.instance)
}

// SetParentShowHint
func (d *TDrawGrid) SetParentShowHint(value bool) {
    DrawGrid_SetParentShowHint(d.instance, value)
}

// PopupMenu
func (d *TDrawGrid) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(DrawGrid_GetPopupMenu(d.instance))
}

// SetPopupMenu
func (d *TDrawGrid) SetPopupMenu(value IComponent) {
    DrawGrid_SetPopupMenu(d.instance, CheckPtr(value))
}

// ScrollBars
func (d *TDrawGrid) ScrollBars() TScrollStyle {
    return DrawGrid_GetScrollBars(d.instance)
}

// SetScrollBars
func (d *TDrawGrid) SetScrollBars(value TScrollStyle) {
    DrawGrid_SetScrollBars(d.instance, value)
}

// ShowHint
func (d *TDrawGrid) ShowHint() bool {
    return DrawGrid_GetShowHint(d.instance)
}

// SetShowHint
func (d *TDrawGrid) SetShowHint(value bool) {
    DrawGrid_SetShowHint(d.instance, value)
}

// TabOrder
func (d *TDrawGrid) TabOrder() TTabOrder {
    return DrawGrid_GetTabOrder(d.instance)
}

// SetTabOrder
func (d *TDrawGrid) SetTabOrder(value TTabOrder) {
    DrawGrid_SetTabOrder(d.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (d *TDrawGrid) Visible() bool {
    return DrawGrid_GetVisible(d.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (d *TDrawGrid) SetVisible(value bool) {
    DrawGrid_SetVisible(d.instance, value)
}

// StyleElements
func (d *TDrawGrid) StyleElements() TStyleElements {
    return DrawGrid_GetStyleElements(d.instance)
}

// SetStyleElements
func (d *TDrawGrid) SetStyleElements(value TStyleElements) {
    DrawGrid_SetStyleElements(d.instance, value)
}

// VisibleColCount
func (d *TDrawGrid) VisibleColCount() int32 {
    return DrawGrid_GetVisibleColCount(d.instance)
}

// VisibleRowCount
func (d *TDrawGrid) VisibleRowCount() int32 {
    return DrawGrid_GetVisibleRowCount(d.instance)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (d *TDrawGrid) SetOnClick(fn TNotifyEvent) {
    DrawGrid_SetOnClick(d.instance, fn)
}

// SetOnColumnMoved
func (d *TDrawGrid) SetOnColumnMoved(fn TMovedEvent) {
    DrawGrid_SetOnColumnMoved(d.instance, fn)
}

// SetOnContextPopup
func (d *TDrawGrid) SetOnContextPopup(fn TContextPopupEvent) {
    DrawGrid_SetOnContextPopup(d.instance, fn)
}

// SetOnDblClick
func (d *TDrawGrid) SetOnDblClick(fn TNotifyEvent) {
    DrawGrid_SetOnDblClick(d.instance, fn)
}

// SetOnDragDrop
func (d *TDrawGrid) SetOnDragDrop(fn TDragDropEvent) {
    DrawGrid_SetOnDragDrop(d.instance, fn)
}

// SetOnDragOver
func (d *TDrawGrid) SetOnDragOver(fn TDragOverEvent) {
    DrawGrid_SetOnDragOver(d.instance, fn)
}

// SetOnDrawCell
func (d *TDrawGrid) SetOnDrawCell(fn TDrawCellEvent) {
    DrawGrid_SetOnDrawCell(d.instance, fn)
}

// SetOnEndDock
func (d *TDrawGrid) SetOnEndDock(fn TEndDragEvent) {
    DrawGrid_SetOnEndDock(d.instance, fn)
}

// SetOnEndDrag
func (d *TDrawGrid) SetOnEndDrag(fn TEndDragEvent) {
    DrawGrid_SetOnEndDrag(d.instance, fn)
}

// SetOnEnter
func (d *TDrawGrid) SetOnEnter(fn TNotifyEvent) {
    DrawGrid_SetOnEnter(d.instance, fn)
}

// SetOnExit
func (d *TDrawGrid) SetOnExit(fn TNotifyEvent) {
    DrawGrid_SetOnExit(d.instance, fn)
}

// SetOnFixedCellClick
func (d *TDrawGrid) SetOnFixedCellClick(fn TFixedCellClickEvent) {
    DrawGrid_SetOnFixedCellClick(d.instance, fn)
}

// SetOnGetEditMask
func (d *TDrawGrid) SetOnGetEditMask(fn TGetEditEvent) {
    DrawGrid_SetOnGetEditMask(d.instance, fn)
}

// SetOnGetEditText
func (d *TDrawGrid) SetOnGetEditText(fn TGetEditEvent) {
    DrawGrid_SetOnGetEditText(d.instance, fn)
}

// SetOnKeyDown
func (d *TDrawGrid) SetOnKeyDown(fn TKeyEvent) {
    DrawGrid_SetOnKeyDown(d.instance, fn)
}

// SetOnKeyPress
func (d *TDrawGrid) SetOnKeyPress(fn TKeyPressEvent) {
    DrawGrid_SetOnKeyPress(d.instance, fn)
}

// SetOnKeyUp
func (d *TDrawGrid) SetOnKeyUp(fn TKeyEvent) {
    DrawGrid_SetOnKeyUp(d.instance, fn)
}

// SetOnMouseDown
func (d *TDrawGrid) SetOnMouseDown(fn TMouseEvent) {
    DrawGrid_SetOnMouseDown(d.instance, fn)
}

// SetOnMouseEnter
func (d *TDrawGrid) SetOnMouseEnter(fn TNotifyEvent) {
    DrawGrid_SetOnMouseEnter(d.instance, fn)
}

// SetOnMouseLeave
func (d *TDrawGrid) SetOnMouseLeave(fn TNotifyEvent) {
    DrawGrid_SetOnMouseLeave(d.instance, fn)
}

// SetOnMouseMove
func (d *TDrawGrid) SetOnMouseMove(fn TMouseMoveEvent) {
    DrawGrid_SetOnMouseMove(d.instance, fn)
}

// SetOnMouseUp
func (d *TDrawGrid) SetOnMouseUp(fn TMouseEvent) {
    DrawGrid_SetOnMouseUp(d.instance, fn)
}

// SetOnMouseWheelDown
func (d *TDrawGrid) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    DrawGrid_SetOnMouseWheelDown(d.instance, fn)
}

// SetOnMouseWheelUp
func (d *TDrawGrid) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    DrawGrid_SetOnMouseWheelUp(d.instance, fn)
}

// SetOnRowMoved
func (d *TDrawGrid) SetOnRowMoved(fn TMovedEvent) {
    DrawGrid_SetOnRowMoved(d.instance, fn)
}

// SetOnSelectCell
func (d *TDrawGrid) SetOnSelectCell(fn TSelectCellEvent) {
    DrawGrid_SetOnSelectCell(d.instance, fn)
}

// SetOnSetEditText
func (d *TDrawGrid) SetOnSetEditText(fn TSetEditEvent) {
    DrawGrid_SetOnSetEditText(d.instance, fn)
}

// SetOnStartDock
func (d *TDrawGrid) SetOnStartDock(fn TStartDockEvent) {
    DrawGrid_SetOnStartDock(d.instance, fn)
}

// SetOnTopLeftChanged
func (d *TDrawGrid) SetOnTopLeftChanged(fn TNotifyEvent) {
    DrawGrid_SetOnTopLeftChanged(d.instance, fn)
}

// Canvas
func (d *TDrawGrid) Canvas() *TCanvas {
    return CanvasFromInst(DrawGrid_GetCanvas(d.instance))
}

// Col
func (d *TDrawGrid) Col() int32 {
    return DrawGrid_GetCol(d.instance)
}

// SetCol
func (d *TDrawGrid) SetCol(value int32) {
    DrawGrid_SetCol(d.instance, value)
}

// EditorMode
func (d *TDrawGrid) EditorMode() bool {
    return DrawGrid_GetEditorMode(d.instance)
}

// SetEditorMode
func (d *TDrawGrid) SetEditorMode(value bool) {
    DrawGrid_SetEditorMode(d.instance, value)
}

// GridHeight
func (d *TDrawGrid) GridHeight() int32 {
    return DrawGrid_GetGridHeight(d.instance)
}

// GridWidth
func (d *TDrawGrid) GridWidth() int32 {
    return DrawGrid_GetGridWidth(d.instance)
}

// LeftCol
func (d *TDrawGrid) LeftCol() int32 {
    return DrawGrid_GetLeftCol(d.instance)
}

// SetLeftCol
func (d *TDrawGrid) SetLeftCol(value int32) {
    DrawGrid_SetLeftCol(d.instance, value)
}

// Selection
func (d *TDrawGrid) Selection() TGridRect {
    return DrawGrid_GetSelection(d.instance)
}

// SetSelection
func (d *TDrawGrid) SetSelection(value TGridRect) {
    DrawGrid_SetSelection(d.instance, value)
}

// Row
func (d *TDrawGrid) Row() int32 {
    return DrawGrid_GetRow(d.instance)
}

// SetRow
func (d *TDrawGrid) SetRow(value int32) {
    DrawGrid_SetRow(d.instance, value)
}

// TopRow
func (d *TDrawGrid) TopRow() int32 {
    return DrawGrid_GetTopRow(d.instance)
}

// SetTopRow
func (d *TDrawGrid) SetTopRow(value int32) {
    DrawGrid_SetTopRow(d.instance, value)
}

// TabStop
func (d *TDrawGrid) TabStop() bool {
    return DrawGrid_GetTabStop(d.instance)
}

// SetTabStop
func (d *TDrawGrid) SetTabStop(value bool) {
    DrawGrid_SetTabStop(d.instance, value)
}

// DockClientCount
func (d *TDrawGrid) DockClientCount() int32 {
    return DrawGrid_GetDockClientCount(d.instance)
}

// DockSite
func (d *TDrawGrid) DockSite() bool {
    return DrawGrid_GetDockSite(d.instance)
}

// SetDockSite
func (d *TDrawGrid) SetDockSite(value bool) {
    DrawGrid_SetDockSite(d.instance, value)
}

// AlignDisabled
func (d *TDrawGrid) AlignDisabled() bool {
    return DrawGrid_GetAlignDisabled(d.instance)
}

// MouseInClient
func (d *TDrawGrid) MouseInClient() bool {
    return DrawGrid_GetMouseInClient(d.instance)
}

// VisibleDockClientCount
func (d *TDrawGrid) VisibleDockClientCount() int32 {
    return DrawGrid_GetVisibleDockClientCount(d.instance)
}

// Brush
func (d *TDrawGrid) Brush() *TBrush {
    return BrushFromInst(DrawGrid_GetBrush(d.instance))
}

// ControlCount
func (d *TDrawGrid) ControlCount() int32 {
    return DrawGrid_GetControlCount(d.instance)
}

// Handle
func (d *TDrawGrid) Handle() HWND {
    return DrawGrid_GetHandle(d.instance)
}

// ParentWindow
func (d *TDrawGrid) ParentWindow() HWND {
    return DrawGrid_GetParentWindow(d.instance)
}

// SetParentWindow
func (d *TDrawGrid) SetParentWindow(value HWND) {
    DrawGrid_SetParentWindow(d.instance, value)
}

// UseDockManager
func (d *TDrawGrid) UseDockManager() bool {
    return DrawGrid_GetUseDockManager(d.instance)
}

// SetUseDockManager
func (d *TDrawGrid) SetUseDockManager(value bool) {
    DrawGrid_SetUseDockManager(d.instance, value)
}

// Action
func (d *TDrawGrid) Action() *TAction {
    return ActionFromInst(DrawGrid_GetAction(d.instance))
}

// SetAction
func (d *TDrawGrid) SetAction(value IComponent) {
    DrawGrid_SetAction(d.instance, CheckPtr(value))
}

// BoundsRect
func (d *TDrawGrid) BoundsRect() TRect {
    return DrawGrid_GetBoundsRect(d.instance)
}

// SetBoundsRect
func (d *TDrawGrid) SetBoundsRect(value TRect) {
    DrawGrid_SetBoundsRect(d.instance, value)
}

// ClientHeight
func (d *TDrawGrid) ClientHeight() int32 {
    return DrawGrid_GetClientHeight(d.instance)
}

// SetClientHeight
func (d *TDrawGrid) SetClientHeight(value int32) {
    DrawGrid_SetClientHeight(d.instance, value)
}

// ClientOrigin
func (d *TDrawGrid) ClientOrigin() TPoint {
    return DrawGrid_GetClientOrigin(d.instance)
}

// ClientRect
func (d *TDrawGrid) ClientRect() TRect {
    return DrawGrid_GetClientRect(d.instance)
}

// ClientWidth
func (d *TDrawGrid) ClientWidth() int32 {
    return DrawGrid_GetClientWidth(d.instance)
}

// SetClientWidth
func (d *TDrawGrid) SetClientWidth(value int32) {
    DrawGrid_SetClientWidth(d.instance, value)
}

// ControlState
func (d *TDrawGrid) ControlState() TControlState {
    return DrawGrid_GetControlState(d.instance)
}

// SetControlState
func (d *TDrawGrid) SetControlState(value TControlState) {
    DrawGrid_SetControlState(d.instance, value)
}

// ControlStyle
func (d *TDrawGrid) ControlStyle() TControlStyle {
    return DrawGrid_GetControlStyle(d.instance)
}

// SetControlStyle
func (d *TDrawGrid) SetControlStyle(value TControlStyle) {
    DrawGrid_SetControlStyle(d.instance, value)
}

// ExplicitLeft
func (d *TDrawGrid) ExplicitLeft() int32 {
    return DrawGrid_GetExplicitLeft(d.instance)
}

// ExplicitTop
func (d *TDrawGrid) ExplicitTop() int32 {
    return DrawGrid_GetExplicitTop(d.instance)
}

// ExplicitWidth
func (d *TDrawGrid) ExplicitWidth() int32 {
    return DrawGrid_GetExplicitWidth(d.instance)
}

// ExplicitHeight
func (d *TDrawGrid) ExplicitHeight() int32 {
    return DrawGrid_GetExplicitHeight(d.instance)
}

// Floating
func (d *TDrawGrid) Floating() bool {
    return DrawGrid_GetFloating(d.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (d *TDrawGrid) Parent() *TWinControl {
    return WinControlFromInst(DrawGrid_GetParent(d.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (d *TDrawGrid) SetParent(value IWinControl) {
    DrawGrid_SetParent(d.instance, CheckPtr(value))
}

// AlignWithMargins
func (d *TDrawGrid) AlignWithMargins() bool {
    return DrawGrid_GetAlignWithMargins(d.instance)
}

// SetAlignWithMargins
func (d *TDrawGrid) SetAlignWithMargins(value bool) {
    DrawGrid_SetAlignWithMargins(d.instance, value)
}

// Left
func (d *TDrawGrid) Left() int32 {
    return DrawGrid_GetLeft(d.instance)
}

// SetLeft
func (d *TDrawGrid) SetLeft(value int32) {
    DrawGrid_SetLeft(d.instance, value)
}

// Top
func (d *TDrawGrid) Top() int32 {
    return DrawGrid_GetTop(d.instance)
}

// SetTop
func (d *TDrawGrid) SetTop(value int32) {
    DrawGrid_SetTop(d.instance, value)
}

// Width
func (d *TDrawGrid) Width() int32 {
    return DrawGrid_GetWidth(d.instance)
}

// SetWidth
func (d *TDrawGrid) SetWidth(value int32) {
    DrawGrid_SetWidth(d.instance, value)
}

// Height
func (d *TDrawGrid) Height() int32 {
    return DrawGrid_GetHeight(d.instance)
}

// SetHeight
func (d *TDrawGrid) SetHeight(value int32) {
    DrawGrid_SetHeight(d.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (d *TDrawGrid) Cursor() TCursor {
    return DrawGrid_GetCursor(d.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (d *TDrawGrid) SetCursor(value TCursor) {
    DrawGrid_SetCursor(d.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (d *TDrawGrid) Hint() string {
    return DrawGrid_GetHint(d.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (d *TDrawGrid) SetHint(value string) {
    DrawGrid_SetHint(d.instance, value)
}

// Margins
func (d *TDrawGrid) Margins() *TMargins {
    return MarginsFromInst(DrawGrid_GetMargins(d.instance))
}

// SetMargins
func (d *TDrawGrid) SetMargins(value *TMargins) {
    DrawGrid_SetMargins(d.instance, CheckPtr(value))
}

// CustomHint
func (d *TDrawGrid) CustomHint() *TCustomHint {
    return CustomHintFromInst(DrawGrid_GetCustomHint(d.instance))
}

// SetCustomHint
func (d *TDrawGrid) SetCustomHint(value IComponent) {
    DrawGrid_SetCustomHint(d.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (d *TDrawGrid) ComponentCount() int32 {
    return DrawGrid_GetComponentCount(d.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (d *TDrawGrid) ComponentIndex() int32 {
    return DrawGrid_GetComponentIndex(d.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (d *TDrawGrid) SetComponentIndex(value int32) {
    DrawGrid_SetComponentIndex(d.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (d *TDrawGrid) Owner() *TComponent {
    return ComponentFromInst(DrawGrid_GetOwner(d.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (d *TDrawGrid) Name() string {
    return DrawGrid_GetName(d.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (d *TDrawGrid) SetName(value string) {
    DrawGrid_SetName(d.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (d *TDrawGrid) Tag() int {
    return DrawGrid_GetTag(d.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (d *TDrawGrid) SetTag(value int) {
    DrawGrid_SetTag(d.instance, value)
}

// ColWidths
func (d *TDrawGrid) ColWidths(Index int32) int32 {
    return DrawGrid_GetColWidths(d.instance, Index)
}

// ColWidths
func (d *TDrawGrid) SetColWidths(Index int32, value int32) {
    DrawGrid_SetColWidths(d.instance, Index, value)
}

// RowHeights
func (d *TDrawGrid) RowHeights(Index int32) int32 {
    return DrawGrid_GetRowHeights(d.instance, Index)
}

// RowHeights
func (d *TDrawGrid) SetRowHeights(Index int32, value int32) {
    DrawGrid_SetRowHeights(d.instance, Index, value)
}

// TabStops
func (d *TDrawGrid) TabStops(Index int32) bool {
    return DrawGrid_GetTabStops(d.instance, Index)
}

// TabStops
func (d *TDrawGrid) SetTabStops(Index int32, value bool) {
    DrawGrid_SetTabStops(d.instance, Index, value)
}

// DockClients
func (d *TDrawGrid) DockClients(Index int32) *TControl {
    return ControlFromInst(DrawGrid_GetDockClients(d.instance, Index))
}

// Controls
func (d *TDrawGrid) Controls(Index int32) *TControl {
    return ControlFromInst(DrawGrid_GetControls(d.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (d *TDrawGrid) Components(AIndex int32) *TComponent {
    return ComponentFromInst(DrawGrid_GetComponents(d.instance, AIndex))
}

