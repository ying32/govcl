
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
// CN: 刷新控件。
// EN: Refresh control.
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
// CN: 是否可以获得焦点。
// EN: .
func (v *TValueListEditor) CanFocus() bool {
    return ValueListEditor_CanFocus(v.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (v *TValueListEditor) ContainsControl(Control IControl) bool {
    return ValueListEditor_ContainsControl(v.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (v *TValueListEditor) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(ValueListEditor_ControlAtPos(v.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (v *TValueListEditor) DisableAlign() {
    ValueListEditor_DisableAlign(v.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (v *TValueListEditor) EnableAlign() {
    ValueListEditor_EnableAlign(v.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (v *TValueListEditor) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(ValueListEditor_FindChildControl(v.instance, ControlName))
}

// FlipChildren
func (v *TValueListEditor) FlipChildren(AllLevels bool) {
    ValueListEditor_FlipChildren(v.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (v *TValueListEditor) Focused() bool {
    return ValueListEditor_Focused(v.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (v *TValueListEditor) HandleAllocated() bool {
    return ValueListEditor_HandleAllocated(v.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (v *TValueListEditor) InsertControl(AControl IControl) {
    ValueListEditor_InsertControl(v.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (v *TValueListEditor) Invalidate() {
    ValueListEditor_Invalidate(v.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (v *TValueListEditor) PaintTo(DC HDC, X int32, Y int32) {
    ValueListEditor_PaintTo(v.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (v *TValueListEditor) RemoveControl(AControl IControl) {
    ValueListEditor_RemoveControl(v.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (v *TValueListEditor) Realign() {
    ValueListEditor_Realign(v.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (v *TValueListEditor) Repaint() {
    ValueListEditor_Repaint(v.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (v *TValueListEditor) ScaleBy(M int32, D int32) {
    ValueListEditor_ScaleBy(v.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (v *TValueListEditor) ScrollBy(DeltaX int32, DeltaY int32) {
    ValueListEditor_ScrollBy(v.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (v *TValueListEditor) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ValueListEditor_SetBounds(v.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (v *TValueListEditor) SetFocus() {
    ValueListEditor_SetFocus(v.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (v *TValueListEditor) Update() {
    ValueListEditor_Update(v.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (v *TValueListEditor) UpdateControlState() {
    ValueListEditor_UpdateControlState(v.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (v *TValueListEditor) BringToFront() {
    ValueListEditor_BringToFront(v.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (v *TValueListEditor) ClientToScreen(Point TPoint) TPoint {
    return ValueListEditor_ClientToScreen(v.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (v *TValueListEditor) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ValueListEditor_ClientToParent(v.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (v *TValueListEditor) Dragging() bool {
    return ValueListEditor_Dragging(v.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (v *TValueListEditor) HasParent() bool {
    return ValueListEditor_HasParent(v.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (v *TValueListEditor) Hide() {
    ValueListEditor_Hide(v.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (v *TValueListEditor) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ValueListEditor_Perform(v.instance, Msg , WParam , LParam)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (v *TValueListEditor) ScreenToClient(Point TPoint) TPoint {
    return ValueListEditor_ScreenToClient(v.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (v *TValueListEditor) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ValueListEditor_ParentToClient(v.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (v *TValueListEditor) SendToBack() {
    ValueListEditor_SendToBack(v.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (v *TValueListEditor) Show() {
    ValueListEditor_Show(v.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (v *TValueListEditor) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ValueListEditor_GetTextBuf(v.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (v *TValueListEditor) GetTextLen() int32 {
    return ValueListEditor_GetTextLen(v.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (v *TValueListEditor) SetTextBuf(Buffer string) {
    ValueListEditor_SetTextBuf(v.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (v *TValueListEditor) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ValueListEditor_FindComponent(v.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (v *TValueListEditor) GetNamePath() string {
    return ValueListEditor_GetNamePath(v.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
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
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (v *TValueListEditor) Align() TAlign {
    return ValueListEditor_GetAlign(v.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (v *TValueListEditor) SetAlign(value TAlign) {
    ValueListEditor_SetAlign(v.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (v *TValueListEditor) Anchors() TAnchors {
    return ValueListEditor_GetAnchors(v.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
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
// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (v *TValueListEditor) BorderStyle() TBorderStyle {
    return ValueListEditor_GetBorderStyle(v.instance)
}

// SetBorderStyle
// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (v *TValueListEditor) SetBorderStyle(value TBorderStyle) {
    ValueListEditor_SetBorderStyle(v.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (v *TValueListEditor) Color() TColor {
    return ValueListEditor_GetColor(v.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (v *TValueListEditor) SetColor(value TColor) {
    ValueListEditor_SetColor(v.instance, value)
}

// Constraints
func (v *TValueListEditor) Constraints() *TSizeConstraints {
    return SizeConstraintsFromInst(ValueListEditor_GetConstraints(v.instance))
}

// SetConstraints
func (v *TValueListEditor) SetConstraints(value *TSizeConstraints) {
    ValueListEditor_SetConstraints(v.instance, CheckPtr(value))
}

// Ctl3D
func (v *TValueListEditor) Ctl3D() bool {
    return ValueListEditor_GetCtl3D(v.instance)
}

// SetCtl3D
func (v *TValueListEditor) SetCtl3D(value bool) {
    ValueListEditor_SetCtl3D(v.instance, value)
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
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (v *TValueListEditor) DoubleBuffered() bool {
    return ValueListEditor_GetDoubleBuffered(v.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (v *TValueListEditor) SetDoubleBuffered(value bool) {
    ValueListEditor_SetDoubleBuffered(v.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (v *TValueListEditor) DragCursor() TCursor {
    return ValueListEditor_GetDragCursor(v.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (v *TValueListEditor) SetDragCursor(value TCursor) {
    ValueListEditor_SetDragCursor(v.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (v *TValueListEditor) DragKind() TDragKind {
    return ValueListEditor_GetDragKind(v.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (v *TValueListEditor) SetDragKind(value TDragKind) {
    ValueListEditor_SetDragKind(v.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (v *TValueListEditor) DragMode() TDragMode {
    return ValueListEditor_GetDragMode(v.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
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
// CN: 获取字体。
// EN: Get Font.
func (v *TValueListEditor) Font() *TFont {
    return FontFromInst(ValueListEditor_GetFont(v.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (v *TValueListEditor) SetFont(value *TFont) {
    ValueListEditor_SetFont(v.instance, CheckPtr(value))
}

// GradientEndColor
// CN: 获取渐变结束颜色, 仅VCL。
// EN: .
func (v *TValueListEditor) GradientEndColor() TColor {
    return ValueListEditor_GetGradientEndColor(v.instance)
}

// SetGradientEndColor
// CN: 设置渐变结束颜色, 仅VCL。
// EN: .
func (v *TValueListEditor) SetGradientEndColor(value TColor) {
    ValueListEditor_SetGradientEndColor(v.instance, value)
}

// GradientStartColor
// CN: 获取渐变起始颜色，仅VCL。
// EN: .
func (v *TValueListEditor) GradientStartColor() TColor {
    return ValueListEditor_GetGradientStartColor(v.instance)
}

// SetGradientStartColor
// CN: 设置渐变起始颜色，仅VCL。
// EN: .
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
// CN: 获取父容器颜色。
// EN: Get parent color.
func (v *TValueListEditor) ParentColor() bool {
    return ValueListEditor_GetParentColor(v.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
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
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (v *TValueListEditor) ParentDoubleBuffered() bool {
    return ValueListEditor_GetParentDoubleBuffered(v.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (v *TValueListEditor) SetParentDoubleBuffered(value bool) {
    ValueListEditor_SetParentDoubleBuffered(v.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (v *TValueListEditor) ParentFont() bool {
    return ValueListEditor_GetParentFont(v.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
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
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (v *TValueListEditor) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ValueListEditor_GetPopupMenu(v.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
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
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (v *TValueListEditor) ShowHint() bool {
    return ValueListEditor_GetShowHint(v.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
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
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (v *TValueListEditor) TabOrder() TTabOrder {
    return ValueListEditor_GetTabOrder(v.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
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
// CN: 获取样式元素。
// EN: Get Style element.
func (v *TValueListEditor) StyleElements() TStyleElements {
    return ValueListEditor_GetStyleElements(v.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
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
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (v *TValueListEditor) SetOnContextPopup(fn TContextPopupEvent) {
    ValueListEditor_SetOnContextPopup(v.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (v *TValueListEditor) SetOnDblClick(fn TNotifyEvent) {
    ValueListEditor_SetOnDblClick(v.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (v *TValueListEditor) SetOnDragDrop(fn TDragDropEvent) {
    ValueListEditor_SetOnDragDrop(v.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (v *TValueListEditor) SetOnDragOver(fn TDragOverEvent) {
    ValueListEditor_SetOnDragOver(v.instance, fn)
}

// SetOnDrawCell
func (v *TValueListEditor) SetOnDrawCell(fn TDrawCellEvent) {
    ValueListEditor_SetOnDrawCell(v.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (v *TValueListEditor) SetOnEndDock(fn TEndDragEvent) {
    ValueListEditor_SetOnEndDock(v.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (v *TValueListEditor) SetOnEndDrag(fn TEndDragEvent) {
    ValueListEditor_SetOnEndDrag(v.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (v *TValueListEditor) SetOnEnter(fn TNotifyEvent) {
    ValueListEditor_SetOnEnter(v.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (v *TValueListEditor) SetOnExit(fn TNotifyEvent) {
    ValueListEditor_SetOnExit(v.instance, fn)
}

// SetOnGesture
func (v *TValueListEditor) SetOnGesture(fn TGestureEvent) {
    ValueListEditor_SetOnGesture(v.instance, fn)
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
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (v *TValueListEditor) SetOnKeyDown(fn TKeyEvent) {
    ValueListEditor_SetOnKeyDown(v.instance, fn)
}

// SetOnKeyPress
func (v *TValueListEditor) SetOnKeyPress(fn TKeyPressEvent) {
    ValueListEditor_SetOnKeyPress(v.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (v *TValueListEditor) SetOnKeyUp(fn TKeyEvent) {
    ValueListEditor_SetOnKeyUp(v.instance, fn)
}

// SetOnMouseActivate
func (v *TValueListEditor) SetOnMouseActivate(fn TMouseActivateEvent) {
    ValueListEditor_SetOnMouseActivate(v.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (v *TValueListEditor) SetOnMouseDown(fn TMouseEvent) {
    ValueListEditor_SetOnMouseDown(v.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (v *TValueListEditor) SetOnMouseEnter(fn TNotifyEvent) {
    ValueListEditor_SetOnMouseEnter(v.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (v *TValueListEditor) SetOnMouseLeave(fn TNotifyEvent) {
    ValueListEditor_SetOnMouseLeave(v.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (v *TValueListEditor) SetOnMouseMove(fn TMouseMoveEvent) {
    ValueListEditor_SetOnMouseMove(v.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (v *TValueListEditor) SetOnMouseUp(fn TMouseEvent) {
    ValueListEditor_SetOnMouseUp(v.instance, fn)
}

// SetOnMouseWheelDown
// CN: 设置鼠标滚轮按下事件。
// EN: .
func (v *TValueListEditor) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    ValueListEditor_SetOnMouseWheelDown(v.instance, fn)
}

// SetOnMouseWheelUp
// CN: 设置鼠标滚轮抬起事件。
// EN: .
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
// CN: 设置启动停靠。
// EN: .
func (v *TValueListEditor) SetOnStartDock(fn TStartDockEvent) {
    ValueListEditor_SetOnStartDock(v.instance, fn)
}

// SetOnTopLeftChanged
func (v *TValueListEditor) SetOnTopLeftChanged(fn TNotifyEvent) {
    ValueListEditor_SetOnTopLeftChanged(v.instance, fn)
}

// Canvas
// CN: 获取画布。
// EN: .
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
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (v *TValueListEditor) TabStop() bool {
    return ValueListEditor_GetTabStop(v.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (v *TValueListEditor) SetTabStop(value bool) {
    ValueListEditor_SetTabStop(v.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (v *TValueListEditor) DockClientCount() int32 {
    return ValueListEditor_GetDockClientCount(v.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (v *TValueListEditor) DockSite() bool {
    return ValueListEditor_GetDockSite(v.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (v *TValueListEditor) SetDockSite(value bool) {
    ValueListEditor_SetDockSite(v.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (v *TValueListEditor) AlignDisabled() bool {
    return ValueListEditor_GetAlignDisabled(v.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (v *TValueListEditor) MouseInClient() bool {
    return ValueListEditor_GetMouseInClient(v.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (v *TValueListEditor) VisibleDockClientCount() int32 {
    return ValueListEditor_GetVisibleDockClientCount(v.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (v *TValueListEditor) Brush() *TBrush {
    return BrushFromInst(ValueListEditor_GetBrush(v.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (v *TValueListEditor) ControlCount() int32 {
    return ValueListEditor_GetControlCount(v.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (v *TValueListEditor) Handle() HWND {
    return ValueListEditor_GetHandle(v.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (v *TValueListEditor) ParentWindow() HWND {
    return ValueListEditor_GetParentWindow(v.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (v *TValueListEditor) SetParentWindow(value HWND) {
    ValueListEditor_SetParentWindow(v.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (v *TValueListEditor) UseDockManager() bool {
    return ValueListEditor_GetUseDockManager(v.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
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
// CN: 获取客户区高度。
// EN: Get client height.
func (v *TValueListEditor) ClientHeight() int32 {
    return ValueListEditor_GetClientHeight(v.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (v *TValueListEditor) SetClientHeight(value int32) {
    ValueListEditor_SetClientHeight(v.instance, value)
}

// ClientOrigin
func (v *TValueListEditor) ClientOrigin() TPoint {
    return ValueListEditor_GetClientOrigin(v.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (v *TValueListEditor) ClientRect() TRect {
    return ValueListEditor_GetClientRect(v.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (v *TValueListEditor) ClientWidth() int32 {
    return ValueListEditor_GetClientWidth(v.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (v *TValueListEditor) SetClientWidth(value int32) {
    ValueListEditor_SetClientWidth(v.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (v *TValueListEditor) ControlState() TControlState {
    return ValueListEditor_GetControlState(v.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (v *TValueListEditor) SetControlState(value TControlState) {
    ValueListEditor_SetControlState(v.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (v *TValueListEditor) ControlStyle() TControlStyle {
    return ValueListEditor_GetControlStyle(v.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
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
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (v *TValueListEditor) AlignWithMargins() bool {
    return ValueListEditor_GetAlignWithMargins(v.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (v *TValueListEditor) SetAlignWithMargins(value bool) {
    ValueListEditor_SetAlignWithMargins(v.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (v *TValueListEditor) Left() int32 {
    return ValueListEditor_GetLeft(v.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (v *TValueListEditor) SetLeft(value int32) {
    ValueListEditor_SetLeft(v.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (v *TValueListEditor) Top() int32 {
    return ValueListEditor_GetTop(v.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (v *TValueListEditor) SetTop(value int32) {
    ValueListEditor_SetTop(v.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (v *TValueListEditor) Width() int32 {
    return ValueListEditor_GetWidth(v.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (v *TValueListEditor) SetWidth(value int32) {
    ValueListEditor_SetWidth(v.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (v *TValueListEditor) Height() int32 {
    return ValueListEditor_GetHeight(v.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
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
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (v *TValueListEditor) Hint() string {
    return ValueListEditor_GetHint(v.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (v *TValueListEditor) SetHint(value string) {
    ValueListEditor_SetHint(v.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (v *TValueListEditor) Margins() *TMargins {
    return MarginsFromInst(ValueListEditor_GetMargins(v.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (v *TValueListEditor) SetMargins(value *TMargins) {
    ValueListEditor_SetMargins(v.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (v *TValueListEditor) CustomHint() *TCustomHint {
    return CustomHintFromInst(ValueListEditor_GetCustomHint(v.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
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
// CN: 获取指定索引停靠客户端。
// EN: .
func (v *TValueListEditor) DockClients(Index int32) *TControl {
    return ControlFromInst(ValueListEditor_GetDockClients(v.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (v *TValueListEditor) Controls(Index int32) *TControl {
    return ControlFromInst(ValueListEditor_GetControls(v.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (v *TValueListEditor) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ValueListEditor_GetComponents(v.instance, AIndex))
}

