
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
// CN: 设置组件边界。
// EN: Set component boundaries.
func (l *TLabeledEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    LabeledEdit_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// SetupInternalLabel
func (l *TLabeledEdit) SetupInternalLabel() {
    LabeledEdit_SetupInternalLabel(l.instance)
}

// Clear
// CN: 清除。
// EN: .
func (l *TLabeledEdit) Clear() {
    LabeledEdit_Clear(l.instance)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (l *TLabeledEdit) ClearSelection() {
    LabeledEdit_ClearSelection(l.instance)
}

// CopyToClipboard
// CN: 复制到粘贴板。
// EN: .
func (l *TLabeledEdit) CopyToClipboard() {
    LabeledEdit_CopyToClipboard(l.instance)
}

// CutToClipboard
// CN: 剪切到粘贴板。
// EN: .
func (l *TLabeledEdit) CutToClipboard() {
    LabeledEdit_CutToClipboard(l.instance)
}

// PasteFromClipboard
// CN: 从剪切板粘贴。
// EN: .
func (l *TLabeledEdit) PasteFromClipboard() {
    LabeledEdit_PasteFromClipboard(l.instance)
}

// Undo
// CN: 撤销上一次操作。
// EN: .
func (l *TLabeledEdit) Undo() {
    LabeledEdit_Undo(l.instance)
}

// ClearUndo
// CN: 清除撤销。
// EN: .
func (l *TLabeledEdit) ClearUndo() {
    LabeledEdit_ClearUndo(l.instance)
}

// SelectAll
// CN: 全选。
// EN: .
func (l *TLabeledEdit) SelectAll() {
    LabeledEdit_SelectAll(l.instance)
}

// GetSelTextBuf
func (l *TLabeledEdit) GetSelTextBuf(Buffer *string, BufSize int32) int32 {
    return LabeledEdit_GetSelTextBuf(l.instance, Buffer , BufSize)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (l *TLabeledEdit) CanFocus() bool {
    return LabeledEdit_CanFocus(l.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (l *TLabeledEdit) ContainsControl(Control IControl) bool {
    return LabeledEdit_ContainsControl(l.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (l *TLabeledEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(LabeledEdit_ControlAtPos(l.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (l *TLabeledEdit) DisableAlign() {
    LabeledEdit_DisableAlign(l.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (l *TLabeledEdit) EnableAlign() {
    LabeledEdit_EnableAlign(l.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (l *TLabeledEdit) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(LabeledEdit_FindChildControl(l.instance, ControlName))
}

// FlipChildren
func (l *TLabeledEdit) FlipChildren(AllLevels bool) {
    LabeledEdit_FlipChildren(l.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (l *TLabeledEdit) Focused() bool {
    return LabeledEdit_Focused(l.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (l *TLabeledEdit) HandleAllocated() bool {
    return LabeledEdit_HandleAllocated(l.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (l *TLabeledEdit) InsertControl(AControl IControl) {
    LabeledEdit_InsertControl(l.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (l *TLabeledEdit) Invalidate() {
    LabeledEdit_Invalidate(l.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (l *TLabeledEdit) PaintTo(DC HDC, X int32, Y int32) {
    LabeledEdit_PaintTo(l.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (l *TLabeledEdit) RemoveControl(AControl IControl) {
    LabeledEdit_RemoveControl(l.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (l *TLabeledEdit) Realign() {
    LabeledEdit_Realign(l.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (l *TLabeledEdit) Repaint() {
    LabeledEdit_Repaint(l.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (l *TLabeledEdit) ScaleBy(M int32, D int32) {
    LabeledEdit_ScaleBy(l.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (l *TLabeledEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    LabeledEdit_ScrollBy(l.instance, DeltaX , DeltaY)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (l *TLabeledEdit) SetFocus() {
    LabeledEdit_SetFocus(l.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (l *TLabeledEdit) Update() {
    LabeledEdit_Update(l.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (l *TLabeledEdit) UpdateControlState() {
    LabeledEdit_UpdateControlState(l.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (l *TLabeledEdit) BringToFront() {
    LabeledEdit_BringToFront(l.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (l *TLabeledEdit) ClientToScreen(Point TPoint) TPoint {
    return LabeledEdit_ClientToScreen(l.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (l *TLabeledEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return LabeledEdit_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (l *TLabeledEdit) Dragging() bool {
    return LabeledEdit_Dragging(l.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (l *TLabeledEdit) HasParent() bool {
    return LabeledEdit_HasParent(l.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (l *TLabeledEdit) Hide() {
    LabeledEdit_Hide(l.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (l *TLabeledEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return LabeledEdit_Perform(l.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (l *TLabeledEdit) Refresh() {
    LabeledEdit_Refresh(l.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (l *TLabeledEdit) ScreenToClient(Point TPoint) TPoint {
    return LabeledEdit_ScreenToClient(l.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (l *TLabeledEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return LabeledEdit_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (l *TLabeledEdit) SendToBack() {
    LabeledEdit_SendToBack(l.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (l *TLabeledEdit) Show() {
    LabeledEdit_Show(l.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (l *TLabeledEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return LabeledEdit_GetTextBuf(l.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (l *TLabeledEdit) GetTextLen() int32 {
    return LabeledEdit_GetTextLen(l.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (l *TLabeledEdit) SetTextBuf(Buffer string) {
    LabeledEdit_SetTextBuf(l.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (l *TLabeledEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(LabeledEdit_FindComponent(l.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TLabeledEdit) GetNamePath() string {
    return LabeledEdit_GetNamePath(l.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
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
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (l *TLabeledEdit) Alignment() TAlignment {
    return LabeledEdit_GetAlignment(l.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (l *TLabeledEdit) SetAlignment(value TAlignment) {
    LabeledEdit_SetAlignment(l.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (l *TLabeledEdit) Anchors() TAnchors {
    return LabeledEdit_GetAnchors(l.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (l *TLabeledEdit) SetAnchors(value TAnchors) {
    LabeledEdit_SetAnchors(l.instance, value)
}

// AutoSelect
// CN: 获取自动选择。
// EN: .
func (l *TLabeledEdit) AutoSelect() bool {
    return LabeledEdit_GetAutoSelect(l.instance)
}

// SetAutoSelect
// CN: 设置自动选择。
// EN: .
func (l *TLabeledEdit) SetAutoSelect(value bool) {
    LabeledEdit_SetAutoSelect(l.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (l *TLabeledEdit) AutoSize() bool {
    return LabeledEdit_GetAutoSize(l.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
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
// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (l *TLabeledEdit) BorderStyle() TBorderStyle {
    return LabeledEdit_GetBorderStyle(l.instance)
}

// SetBorderStyle
// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
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
// CN: 获取颜色。
// EN: Get color.
func (l *TLabeledEdit) Color() TColor {
    return LabeledEdit_GetColor(l.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (l *TLabeledEdit) SetColor(value TColor) {
    LabeledEdit_SetColor(l.instance, value)
}

// Constraints
func (l *TLabeledEdit) Constraints() *TSizeConstraints {
    return SizeConstraintsFromInst(LabeledEdit_GetConstraints(l.instance))
}

// SetConstraints
func (l *TLabeledEdit) SetConstraints(value *TSizeConstraints) {
    LabeledEdit_SetConstraints(l.instance, CheckPtr(value))
}

// Ctl3D
func (l *TLabeledEdit) Ctl3D() bool {
    return LabeledEdit_GetCtl3D(l.instance)
}

// SetCtl3D
func (l *TLabeledEdit) SetCtl3D(value bool) {
    LabeledEdit_SetCtl3D(l.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (l *TLabeledEdit) DoubleBuffered() bool {
    return LabeledEdit_GetDoubleBuffered(l.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (l *TLabeledEdit) SetDoubleBuffered(value bool) {
    LabeledEdit_SetDoubleBuffered(l.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (l *TLabeledEdit) DragCursor() TCursor {
    return LabeledEdit_GetDragCursor(l.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (l *TLabeledEdit) SetDragCursor(value TCursor) {
    LabeledEdit_SetDragCursor(l.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (l *TLabeledEdit) DragKind() TDragKind {
    return LabeledEdit_GetDragKind(l.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (l *TLabeledEdit) SetDragKind(value TDragKind) {
    LabeledEdit_SetDragKind(l.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (l *TLabeledEdit) DragMode() TDragMode {
    return LabeledEdit_GetDragMode(l.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
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
// CN: 获取字体。
// EN: Get Font.
func (l *TLabeledEdit) Font() *TFont {
    return FontFromInst(LabeledEdit_GetFont(l.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (l *TLabeledEdit) SetFont(value *TFont) {
    LabeledEdit_SetFont(l.instance, CheckPtr(value))
}

// HideSelection
// CN: 获取隐藏选择。
// EN: .
func (l *TLabeledEdit) HideSelection() bool {
    return LabeledEdit_GetHideSelection(l.instance)
}

// SetHideSelection
// CN: 设置隐藏选择。
// EN: .
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
// CN: 获取最大长度。
// EN: .
func (l *TLabeledEdit) MaxLength() int32 {
    return LabeledEdit_GetMaxLength(l.instance)
}

// SetMaxLength
// CN: 设置最大长度。
// EN: .
func (l *TLabeledEdit) SetMaxLength(value int32) {
    LabeledEdit_SetMaxLength(l.instance, value)
}

// NumbersOnly
// CN: 获取只能输入数字。
// EN: .
func (l *TLabeledEdit) NumbersOnly() bool {
    return LabeledEdit_GetNumbersOnly(l.instance)
}

// SetNumbersOnly
// CN: 设置只能输入数字。
// EN: .
func (l *TLabeledEdit) SetNumbersOnly(value bool) {
    LabeledEdit_SetNumbersOnly(l.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (l *TLabeledEdit) ParentColor() bool {
    return LabeledEdit_GetParentColor(l.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
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
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (l *TLabeledEdit) ParentDoubleBuffered() bool {
    return LabeledEdit_GetParentDoubleBuffered(l.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (l *TLabeledEdit) SetParentDoubleBuffered(value bool) {
    LabeledEdit_SetParentDoubleBuffered(l.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (l *TLabeledEdit) ParentFont() bool {
    return LabeledEdit_GetParentFont(l.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
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
// CN: 获取密码掩码字符。
// EN: .
func (l *TLabeledEdit) PasswordChar() uint16 {
    return LabeledEdit_GetPasswordChar(l.instance)
}

// SetPasswordChar
// CN: 设置密码掩码字符。
// EN: .
func (l *TLabeledEdit) SetPasswordChar(value uint16) {
    LabeledEdit_SetPasswordChar(l.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (l *TLabeledEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(LabeledEdit_GetPopupMenu(l.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (l *TLabeledEdit) SetPopupMenu(value IComponent) {
    LabeledEdit_SetPopupMenu(l.instance, CheckPtr(value))
}

// ReadOnly
// CN: 获取只读。
// EN: .
func (l *TLabeledEdit) ReadOnly() bool {
    return LabeledEdit_GetReadOnly(l.instance)
}

// SetReadOnly
// CN: 设置只读。
// EN: .
func (l *TLabeledEdit) SetReadOnly(value bool) {
    LabeledEdit_SetReadOnly(l.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (l *TLabeledEdit) ShowHint() bool {
    return LabeledEdit_GetShowHint(l.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (l *TLabeledEdit) SetShowHint(value bool) {
    LabeledEdit_SetShowHint(l.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (l *TLabeledEdit) TabOrder() TTabOrder {
    return LabeledEdit_GetTabOrder(l.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (l *TLabeledEdit) SetTabOrder(value TTabOrder) {
    LabeledEdit_SetTabOrder(l.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (l *TLabeledEdit) TabStop() bool {
    return LabeledEdit_GetTabStop(l.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (l *TLabeledEdit) SetTabStop(value bool) {
    LabeledEdit_SetTabStop(l.instance, value)
}

// Text
// CN: 获取文本。
// EN: .
func (l *TLabeledEdit) Text() string {
    return LabeledEdit_GetText(l.instance)
}

// SetText
// CN: 设置文本。
// EN: .
func (l *TLabeledEdit) SetText(value string) {
    LabeledEdit_SetText(l.instance, value)
}

// TextHint
// CN: 获取提示文本。
// EN: .
func (l *TLabeledEdit) TextHint() string {
    return LabeledEdit_GetTextHint(l.instance)
}

// SetTextHint
// CN: 设置提示文本。
// EN: .
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
// CN: 获取样式元素。
// EN: Get Style element.
func (l *TLabeledEdit) StyleElements() TStyleElements {
    return LabeledEdit_GetStyleElements(l.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (l *TLabeledEdit) SetStyleElements(value TStyleElements) {
    LabeledEdit_SetStyleElements(l.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
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
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (l *TLabeledEdit) SetOnContextPopup(fn TContextPopupEvent) {
    LabeledEdit_SetOnContextPopup(l.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (l *TLabeledEdit) SetOnDblClick(fn TNotifyEvent) {
    LabeledEdit_SetOnDblClick(l.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (l *TLabeledEdit) SetOnDragDrop(fn TDragDropEvent) {
    LabeledEdit_SetOnDragDrop(l.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (l *TLabeledEdit) SetOnDragOver(fn TDragOverEvent) {
    LabeledEdit_SetOnDragOver(l.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (l *TLabeledEdit) SetOnEndDock(fn TEndDragEvent) {
    LabeledEdit_SetOnEndDock(l.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (l *TLabeledEdit) SetOnEndDrag(fn TEndDragEvent) {
    LabeledEdit_SetOnEndDrag(l.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (l *TLabeledEdit) SetOnEnter(fn TNotifyEvent) {
    LabeledEdit_SetOnEnter(l.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (l *TLabeledEdit) SetOnExit(fn TNotifyEvent) {
    LabeledEdit_SetOnExit(l.instance, fn)
}

// SetOnGesture
func (l *TLabeledEdit) SetOnGesture(fn TGestureEvent) {
    LabeledEdit_SetOnGesture(l.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (l *TLabeledEdit) SetOnKeyDown(fn TKeyEvent) {
    LabeledEdit_SetOnKeyDown(l.instance, fn)
}

// SetOnKeyPress
func (l *TLabeledEdit) SetOnKeyPress(fn TKeyPressEvent) {
    LabeledEdit_SetOnKeyPress(l.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (l *TLabeledEdit) SetOnKeyUp(fn TKeyEvent) {
    LabeledEdit_SetOnKeyUp(l.instance, fn)
}

// SetOnMouseActivate
func (l *TLabeledEdit) SetOnMouseActivate(fn TMouseActivateEvent) {
    LabeledEdit_SetOnMouseActivate(l.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (l *TLabeledEdit) SetOnMouseDown(fn TMouseEvent) {
    LabeledEdit_SetOnMouseDown(l.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (l *TLabeledEdit) SetOnMouseEnter(fn TNotifyEvent) {
    LabeledEdit_SetOnMouseEnter(l.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (l *TLabeledEdit) SetOnMouseLeave(fn TNotifyEvent) {
    LabeledEdit_SetOnMouseLeave(l.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (l *TLabeledEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    LabeledEdit_SetOnMouseMove(l.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (l *TLabeledEdit) SetOnMouseUp(fn TMouseEvent) {
    LabeledEdit_SetOnMouseUp(l.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (l *TLabeledEdit) SetOnStartDock(fn TStartDockEvent) {
    LabeledEdit_SetOnStartDock(l.instance, fn)
}

// CanUndo
// CN: 获取能否撤销。
// EN: .
func (l *TLabeledEdit) CanUndo() bool {
    return LabeledEdit_GetCanUndo(l.instance)
}

// Modified
// CN: 获取修改。
// EN: Get modified.
func (l *TLabeledEdit) Modified() bool {
    return LabeledEdit_GetModified(l.instance)
}

// SetModified
// CN: 设置修改。
// EN: Set modified.
func (l *TLabeledEdit) SetModified(value bool) {
    LabeledEdit_SetModified(l.instance, value)
}

// SelLength
// CN: 获取选择的长度。
// EN: .
func (l *TLabeledEdit) SelLength() int32 {
    return LabeledEdit_GetSelLength(l.instance)
}

// SetSelLength
// CN: 设置选择的长度。
// EN: .
func (l *TLabeledEdit) SetSelLength(value int32) {
    LabeledEdit_SetSelLength(l.instance, value)
}

// SelStart
// CN: 获取选择的启始位置。
// EN: .
func (l *TLabeledEdit) SelStart() int32 {
    return LabeledEdit_GetSelStart(l.instance)
}

// SetSelStart
// CN: 设置选择的启始位置。
// EN: .
func (l *TLabeledEdit) SetSelStart(value int32) {
    LabeledEdit_SetSelStart(l.instance, value)
}

// SelText
// CN: 获取选择的文本。
// EN: .
func (l *TLabeledEdit) SelText() string {
    return LabeledEdit_GetSelText(l.instance)
}

// SetSelText
// CN: 设置选择的文本。
// EN: .
func (l *TLabeledEdit) SetSelText(value string) {
    LabeledEdit_SetSelText(l.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (l *TLabeledEdit) DockClientCount() int32 {
    return LabeledEdit_GetDockClientCount(l.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (l *TLabeledEdit) DockSite() bool {
    return LabeledEdit_GetDockSite(l.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (l *TLabeledEdit) SetDockSite(value bool) {
    LabeledEdit_SetDockSite(l.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (l *TLabeledEdit) AlignDisabled() bool {
    return LabeledEdit_GetAlignDisabled(l.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (l *TLabeledEdit) MouseInClient() bool {
    return LabeledEdit_GetMouseInClient(l.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (l *TLabeledEdit) VisibleDockClientCount() int32 {
    return LabeledEdit_GetVisibleDockClientCount(l.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (l *TLabeledEdit) Brush() *TBrush {
    return BrushFromInst(LabeledEdit_GetBrush(l.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (l *TLabeledEdit) ControlCount() int32 {
    return LabeledEdit_GetControlCount(l.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (l *TLabeledEdit) Handle() HWND {
    return LabeledEdit_GetHandle(l.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (l *TLabeledEdit) ParentWindow() HWND {
    return LabeledEdit_GetParentWindow(l.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (l *TLabeledEdit) SetParentWindow(value HWND) {
    LabeledEdit_SetParentWindow(l.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (l *TLabeledEdit) UseDockManager() bool {
    return LabeledEdit_GetUseDockManager(l.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
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
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (l *TLabeledEdit) Align() TAlign {
    return LabeledEdit_GetAlign(l.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
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
// CN: 获取客户区高度。
// EN: Get client height.
func (l *TLabeledEdit) ClientHeight() int32 {
    return LabeledEdit_GetClientHeight(l.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (l *TLabeledEdit) SetClientHeight(value int32) {
    LabeledEdit_SetClientHeight(l.instance, value)
}

// ClientOrigin
func (l *TLabeledEdit) ClientOrigin() TPoint {
    return LabeledEdit_GetClientOrigin(l.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (l *TLabeledEdit) ClientRect() TRect {
    return LabeledEdit_GetClientRect(l.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (l *TLabeledEdit) ClientWidth() int32 {
    return LabeledEdit_GetClientWidth(l.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (l *TLabeledEdit) SetClientWidth(value int32) {
    LabeledEdit_SetClientWidth(l.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (l *TLabeledEdit) ControlState() TControlState {
    return LabeledEdit_GetControlState(l.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (l *TLabeledEdit) SetControlState(value TControlState) {
    LabeledEdit_SetControlState(l.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (l *TLabeledEdit) ControlStyle() TControlStyle {
    return LabeledEdit_GetControlStyle(l.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
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
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (l *TLabeledEdit) AlignWithMargins() bool {
    return LabeledEdit_GetAlignWithMargins(l.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (l *TLabeledEdit) SetAlignWithMargins(value bool) {
    LabeledEdit_SetAlignWithMargins(l.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (l *TLabeledEdit) Left() int32 {
    return LabeledEdit_GetLeft(l.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (l *TLabeledEdit) SetLeft(value int32) {
    LabeledEdit_SetLeft(l.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (l *TLabeledEdit) Top() int32 {
    return LabeledEdit_GetTop(l.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (l *TLabeledEdit) SetTop(value int32) {
    LabeledEdit_SetTop(l.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (l *TLabeledEdit) Width() int32 {
    return LabeledEdit_GetWidth(l.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (l *TLabeledEdit) SetWidth(value int32) {
    LabeledEdit_SetWidth(l.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (l *TLabeledEdit) Height() int32 {
    return LabeledEdit_GetHeight(l.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
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
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (l *TLabeledEdit) Hint() string {
    return LabeledEdit_GetHint(l.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (l *TLabeledEdit) SetHint(value string) {
    LabeledEdit_SetHint(l.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (l *TLabeledEdit) Margins() *TMargins {
    return MarginsFromInst(LabeledEdit_GetMargins(l.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (l *TLabeledEdit) SetMargins(value *TMargins) {
    LabeledEdit_SetMargins(l.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (l *TLabeledEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(LabeledEdit_GetCustomHint(l.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
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
// CN: 获取指定索引停靠客户端。
// EN: .
func (l *TLabeledEdit) DockClients(Index int32) *TControl {
    return ControlFromInst(LabeledEdit_GetDockClients(l.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (l *TLabeledEdit) Controls(Index int32) *TControl {
    return ControlFromInst(LabeledEdit_GetControls(l.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TLabeledEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(LabeledEdit_GetComponents(l.instance, AIndex))
}

