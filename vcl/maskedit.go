
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
// CN: 清除。
// EN: .
func (m *TMaskEdit) Clear() {
    MaskEdit_Clear(m.instance)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (m *TMaskEdit) GetTextLen() int32 {
    return MaskEdit_GetTextLen(m.instance)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (m *TMaskEdit) ClearSelection() {
    MaskEdit_ClearSelection(m.instance)
}

// CopyToClipboard
// CN: 复制到粘贴板。
// EN: .
func (m *TMaskEdit) CopyToClipboard() {
    MaskEdit_CopyToClipboard(m.instance)
}

// CutToClipboard
// CN: 剪切到粘贴板。
// EN: .
func (m *TMaskEdit) CutToClipboard() {
    MaskEdit_CutToClipboard(m.instance)
}

// PasteFromClipboard
// CN: 从剪切板粘贴。
// EN: .
func (m *TMaskEdit) PasteFromClipboard() {
    MaskEdit_PasteFromClipboard(m.instance)
}

// Undo
// CN: 撤销上一次操作。
// EN: .
func (m *TMaskEdit) Undo() {
    MaskEdit_Undo(m.instance)
}

// ClearUndo
// CN: 清除撤销。
// EN: .
func (m *TMaskEdit) ClearUndo() {
    MaskEdit_ClearUndo(m.instance)
}

// SelectAll
// CN: 全选。
// EN: .
func (m *TMaskEdit) SelectAll() {
    MaskEdit_SelectAll(m.instance)
}

// GetSelTextBuf
func (m *TMaskEdit) GetSelTextBuf(Buffer *string, BufSize int32) int32 {
    return MaskEdit_GetSelTextBuf(m.instance, Buffer , BufSize)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (m *TMaskEdit) CanFocus() bool {
    return MaskEdit_CanFocus(m.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (m *TMaskEdit) ContainsControl(Control IControl) bool {
    return MaskEdit_ContainsControl(m.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (m *TMaskEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(MaskEdit_ControlAtPos(m.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (m *TMaskEdit) DisableAlign() {
    MaskEdit_DisableAlign(m.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (m *TMaskEdit) EnableAlign() {
    MaskEdit_EnableAlign(m.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (m *TMaskEdit) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(MaskEdit_FindChildControl(m.instance, ControlName))
}

// FlipChildren
func (m *TMaskEdit) FlipChildren(AllLevels bool) {
    MaskEdit_FlipChildren(m.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (m *TMaskEdit) Focused() bool {
    return MaskEdit_Focused(m.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (m *TMaskEdit) HandleAllocated() bool {
    return MaskEdit_HandleAllocated(m.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (m *TMaskEdit) InsertControl(AControl IControl) {
    MaskEdit_InsertControl(m.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (m *TMaskEdit) Invalidate() {
    MaskEdit_Invalidate(m.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (m *TMaskEdit) PaintTo(DC HDC, X int32, Y int32) {
    MaskEdit_PaintTo(m.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (m *TMaskEdit) RemoveControl(AControl IControl) {
    MaskEdit_RemoveControl(m.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (m *TMaskEdit) Realign() {
    MaskEdit_Realign(m.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (m *TMaskEdit) Repaint() {
    MaskEdit_Repaint(m.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (m *TMaskEdit) ScaleBy(M int32, D int32) {
    MaskEdit_ScaleBy(m.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (m *TMaskEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    MaskEdit_ScrollBy(m.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (m *TMaskEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    MaskEdit_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (m *TMaskEdit) SetFocus() {
    MaskEdit_SetFocus(m.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (m *TMaskEdit) Update() {
    MaskEdit_Update(m.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (m *TMaskEdit) UpdateControlState() {
    MaskEdit_UpdateControlState(m.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (m *TMaskEdit) BringToFront() {
    MaskEdit_BringToFront(m.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (m *TMaskEdit) ClientToScreen(Point TPoint) TPoint {
    return MaskEdit_ClientToScreen(m.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (m *TMaskEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return MaskEdit_ClientToParent(m.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (m *TMaskEdit) Dragging() bool {
    return MaskEdit_Dragging(m.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (m *TMaskEdit) HasParent() bool {
    return MaskEdit_HasParent(m.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (m *TMaskEdit) Hide() {
    MaskEdit_Hide(m.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (m *TMaskEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return MaskEdit_Perform(m.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (m *TMaskEdit) Refresh() {
    MaskEdit_Refresh(m.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (m *TMaskEdit) ScreenToClient(Point TPoint) TPoint {
    return MaskEdit_ScreenToClient(m.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (m *TMaskEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return MaskEdit_ParentToClient(m.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (m *TMaskEdit) SendToBack() {
    MaskEdit_SendToBack(m.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (m *TMaskEdit) Show() {
    MaskEdit_Show(m.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (m *TMaskEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return MaskEdit_GetTextBuf(m.instance, Buffer , BufSize)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (m *TMaskEdit) SetTextBuf(Buffer string) {
    MaskEdit_SetTextBuf(m.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (m *TMaskEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(MaskEdit_FindComponent(m.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (m *TMaskEdit) GetNamePath() string {
    return MaskEdit_GetNamePath(m.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
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
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (m *TMaskEdit) Align() TAlign {
    return MaskEdit_GetAlign(m.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (m *TMaskEdit) SetAlign(value TAlign) {
    MaskEdit_SetAlign(m.instance, value)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (m *TMaskEdit) Alignment() TAlignment {
    return MaskEdit_GetAlignment(m.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (m *TMaskEdit) SetAlignment(value TAlignment) {
    MaskEdit_SetAlignment(m.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (m *TMaskEdit) Anchors() TAnchors {
    return MaskEdit_GetAnchors(m.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (m *TMaskEdit) SetAnchors(value TAnchors) {
    MaskEdit_SetAnchors(m.instance, value)
}

// AutoSelect
// CN: 获取自动选择。
// EN: .
func (m *TMaskEdit) AutoSelect() bool {
    return MaskEdit_GetAutoSelect(m.instance)
}

// SetAutoSelect
// CN: 设置自动选择。
// EN: .
func (m *TMaskEdit) SetAutoSelect(value bool) {
    MaskEdit_SetAutoSelect(m.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (m *TMaskEdit) AutoSize() bool {
    return MaskEdit_GetAutoSize(m.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
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
// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (m *TMaskEdit) BorderStyle() TBorderStyle {
    return MaskEdit_GetBorderStyle(m.instance)
}

// SetBorderStyle
// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
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
// CN: 获取颜色。
// EN: Get color.
func (m *TMaskEdit) Color() TColor {
    return MaskEdit_GetColor(m.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (m *TMaskEdit) SetColor(value TColor) {
    MaskEdit_SetColor(m.instance, value)
}

// Constraints
func (m *TMaskEdit) Constraints() *TSizeConstraints {
    return SizeConstraintsFromInst(MaskEdit_GetConstraints(m.instance))
}

// SetConstraints
func (m *TMaskEdit) SetConstraints(value *TSizeConstraints) {
    MaskEdit_SetConstraints(m.instance, CheckPtr(value))
}

// Ctl3D
func (m *TMaskEdit) Ctl3D() bool {
    return MaskEdit_GetCtl3D(m.instance)
}

// SetCtl3D
func (m *TMaskEdit) SetCtl3D(value bool) {
    MaskEdit_SetCtl3D(m.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (m *TMaskEdit) DoubleBuffered() bool {
    return MaskEdit_GetDoubleBuffered(m.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (m *TMaskEdit) SetDoubleBuffered(value bool) {
    MaskEdit_SetDoubleBuffered(m.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (m *TMaskEdit) DragCursor() TCursor {
    return MaskEdit_GetDragCursor(m.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (m *TMaskEdit) SetDragCursor(value TCursor) {
    MaskEdit_SetDragCursor(m.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (m *TMaskEdit) DragKind() TDragKind {
    return MaskEdit_GetDragKind(m.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (m *TMaskEdit) SetDragKind(value TDragKind) {
    MaskEdit_SetDragKind(m.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (m *TMaskEdit) DragMode() TDragMode {
    return MaskEdit_GetDragMode(m.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
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
// CN: 获取字体。
// EN: Get Font.
func (m *TMaskEdit) Font() *TFont {
    return FontFromInst(MaskEdit_GetFont(m.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (m *TMaskEdit) SetFont(value *TFont) {
    MaskEdit_SetFont(m.instance, CheckPtr(value))
}

// MaxLength
// CN: 获取最大长度。
// EN: .
func (m *TMaskEdit) MaxLength() int32 {
    return MaskEdit_GetMaxLength(m.instance)
}

// SetMaxLength
// CN: 设置最大长度。
// EN: .
func (m *TMaskEdit) SetMaxLength(value int32) {
    MaskEdit_SetMaxLength(m.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (m *TMaskEdit) ParentColor() bool {
    return MaskEdit_GetParentColor(m.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
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
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (m *TMaskEdit) ParentDoubleBuffered() bool {
    return MaskEdit_GetParentDoubleBuffered(m.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (m *TMaskEdit) SetParentDoubleBuffered(value bool) {
    MaskEdit_SetParentDoubleBuffered(m.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (m *TMaskEdit) ParentFont() bool {
    return MaskEdit_GetParentFont(m.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
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
// CN: 获取密码掩码字符。
// EN: .
func (m *TMaskEdit) PasswordChar() uint16 {
    return MaskEdit_GetPasswordChar(m.instance)
}

// SetPasswordChar
// CN: 设置密码掩码字符。
// EN: .
func (m *TMaskEdit) SetPasswordChar(value uint16) {
    MaskEdit_SetPasswordChar(m.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (m *TMaskEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(MaskEdit_GetPopupMenu(m.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (m *TMaskEdit) SetPopupMenu(value IComponent) {
    MaskEdit_SetPopupMenu(m.instance, CheckPtr(value))
}

// ReadOnly
// CN: 获取只读。
// EN: .
func (m *TMaskEdit) ReadOnly() bool {
    return MaskEdit_GetReadOnly(m.instance)
}

// SetReadOnly
// CN: 设置只读。
// EN: .
func (m *TMaskEdit) SetReadOnly(value bool) {
    MaskEdit_SetReadOnly(m.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (m *TMaskEdit) ShowHint() bool {
    return MaskEdit_GetShowHint(m.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (m *TMaskEdit) SetShowHint(value bool) {
    MaskEdit_SetShowHint(m.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (m *TMaskEdit) TabOrder() TTabOrder {
    return MaskEdit_GetTabOrder(m.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (m *TMaskEdit) SetTabOrder(value TTabOrder) {
    MaskEdit_SetTabOrder(m.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (m *TMaskEdit) TabStop() bool {
    return MaskEdit_GetTabStop(m.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (m *TMaskEdit) SetTabStop(value bool) {
    MaskEdit_SetTabStop(m.instance, value)
}

// Text
// CN: 获取文本。
// EN: .
func (m *TMaskEdit) Text() string {
    return MaskEdit_GetText(m.instance)
}

// SetText
// CN: 设置文本。
// EN: .
func (m *TMaskEdit) SetText(value string) {
    MaskEdit_SetText(m.instance, value)
}

// TextHint
// CN: 获取提示文本。
// EN: .
func (m *TMaskEdit) TextHint() string {
    return MaskEdit_GetTextHint(m.instance)
}

// SetTextHint
// CN: 设置提示文本。
// EN: .
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
// CN: 获取样式元素。
// EN: Get Style element.
func (m *TMaskEdit) StyleElements() TStyleElements {
    return MaskEdit_GetStyleElements(m.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (m *TMaskEdit) SetStyleElements(value TStyleElements) {
    MaskEdit_SetStyleElements(m.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
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
// CN: 设置双击事件。
// EN: .
func (m *TMaskEdit) SetOnDblClick(fn TNotifyEvent) {
    MaskEdit_SetOnDblClick(m.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (m *TMaskEdit) SetOnDragDrop(fn TDragDropEvent) {
    MaskEdit_SetOnDragDrop(m.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (m *TMaskEdit) SetOnDragOver(fn TDragOverEvent) {
    MaskEdit_SetOnDragOver(m.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (m *TMaskEdit) SetOnEndDock(fn TEndDragEvent) {
    MaskEdit_SetOnEndDock(m.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (m *TMaskEdit) SetOnEndDrag(fn TEndDragEvent) {
    MaskEdit_SetOnEndDrag(m.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (m *TMaskEdit) SetOnEnter(fn TNotifyEvent) {
    MaskEdit_SetOnEnter(m.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (m *TMaskEdit) SetOnExit(fn TNotifyEvent) {
    MaskEdit_SetOnExit(m.instance, fn)
}

// SetOnGesture
func (m *TMaskEdit) SetOnGesture(fn TGestureEvent) {
    MaskEdit_SetOnGesture(m.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (m *TMaskEdit) SetOnKeyDown(fn TKeyEvent) {
    MaskEdit_SetOnKeyDown(m.instance, fn)
}

// SetOnKeyPress
func (m *TMaskEdit) SetOnKeyPress(fn TKeyPressEvent) {
    MaskEdit_SetOnKeyPress(m.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (m *TMaskEdit) SetOnKeyUp(fn TKeyEvent) {
    MaskEdit_SetOnKeyUp(m.instance, fn)
}

// SetOnMouseActivate
func (m *TMaskEdit) SetOnMouseActivate(fn TMouseActivateEvent) {
    MaskEdit_SetOnMouseActivate(m.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (m *TMaskEdit) SetOnMouseDown(fn TMouseEvent) {
    MaskEdit_SetOnMouseDown(m.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (m *TMaskEdit) SetOnMouseEnter(fn TNotifyEvent) {
    MaskEdit_SetOnMouseEnter(m.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (m *TMaskEdit) SetOnMouseLeave(fn TNotifyEvent) {
    MaskEdit_SetOnMouseLeave(m.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (m *TMaskEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    MaskEdit_SetOnMouseMove(m.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (m *TMaskEdit) SetOnMouseUp(fn TMouseEvent) {
    MaskEdit_SetOnMouseUp(m.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
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
// CN: 获取能否撤销。
// EN: .
func (m *TMaskEdit) CanUndo() bool {
    return MaskEdit_GetCanUndo(m.instance)
}

// Modified
// CN: 获取修改。
// EN: Get modified.
func (m *TMaskEdit) Modified() bool {
    return MaskEdit_GetModified(m.instance)
}

// SetModified
// CN: 设置修改。
// EN: Set modified.
func (m *TMaskEdit) SetModified(value bool) {
    MaskEdit_SetModified(m.instance, value)
}

// SelLength
// CN: 获取选择的长度。
// EN: .
func (m *TMaskEdit) SelLength() int32 {
    return MaskEdit_GetSelLength(m.instance)
}

// SetSelLength
// CN: 设置选择的长度。
// EN: .
func (m *TMaskEdit) SetSelLength(value int32) {
    MaskEdit_SetSelLength(m.instance, value)
}

// SelStart
// CN: 获取选择的启始位置。
// EN: .
func (m *TMaskEdit) SelStart() int32 {
    return MaskEdit_GetSelStart(m.instance)
}

// SetSelStart
// CN: 设置选择的启始位置。
// EN: .
func (m *TMaskEdit) SetSelStart(value int32) {
    MaskEdit_SetSelStart(m.instance, value)
}

// SelText
// CN: 获取选择的文本。
// EN: .
func (m *TMaskEdit) SelText() string {
    return MaskEdit_GetSelText(m.instance)
}

// SetSelText
// CN: 设置选择的文本。
// EN: .
func (m *TMaskEdit) SetSelText(value string) {
    MaskEdit_SetSelText(m.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (m *TMaskEdit) DockClientCount() int32 {
    return MaskEdit_GetDockClientCount(m.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (m *TMaskEdit) DockSite() bool {
    return MaskEdit_GetDockSite(m.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (m *TMaskEdit) SetDockSite(value bool) {
    MaskEdit_SetDockSite(m.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (m *TMaskEdit) AlignDisabled() bool {
    return MaskEdit_GetAlignDisabled(m.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (m *TMaskEdit) MouseInClient() bool {
    return MaskEdit_GetMouseInClient(m.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (m *TMaskEdit) VisibleDockClientCount() int32 {
    return MaskEdit_GetVisibleDockClientCount(m.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (m *TMaskEdit) Brush() *TBrush {
    return BrushFromInst(MaskEdit_GetBrush(m.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (m *TMaskEdit) ControlCount() int32 {
    return MaskEdit_GetControlCount(m.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (m *TMaskEdit) Handle() HWND {
    return MaskEdit_GetHandle(m.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (m *TMaskEdit) ParentWindow() HWND {
    return MaskEdit_GetParentWindow(m.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (m *TMaskEdit) SetParentWindow(value HWND) {
    MaskEdit_SetParentWindow(m.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (m *TMaskEdit) UseDockManager() bool {
    return MaskEdit_GetUseDockManager(m.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
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
// CN: 获取客户区高度。
// EN: Get client height.
func (m *TMaskEdit) ClientHeight() int32 {
    return MaskEdit_GetClientHeight(m.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (m *TMaskEdit) SetClientHeight(value int32) {
    MaskEdit_SetClientHeight(m.instance, value)
}

// ClientOrigin
func (m *TMaskEdit) ClientOrigin() TPoint {
    return MaskEdit_GetClientOrigin(m.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (m *TMaskEdit) ClientRect() TRect {
    return MaskEdit_GetClientRect(m.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (m *TMaskEdit) ClientWidth() int32 {
    return MaskEdit_GetClientWidth(m.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (m *TMaskEdit) SetClientWidth(value int32) {
    MaskEdit_SetClientWidth(m.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (m *TMaskEdit) ControlState() TControlState {
    return MaskEdit_GetControlState(m.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (m *TMaskEdit) SetControlState(value TControlState) {
    MaskEdit_SetControlState(m.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (m *TMaskEdit) ControlStyle() TControlStyle {
    return MaskEdit_GetControlStyle(m.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
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
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (m *TMaskEdit) AlignWithMargins() bool {
    return MaskEdit_GetAlignWithMargins(m.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (m *TMaskEdit) SetAlignWithMargins(value bool) {
    MaskEdit_SetAlignWithMargins(m.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (m *TMaskEdit) Left() int32 {
    return MaskEdit_GetLeft(m.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (m *TMaskEdit) SetLeft(value int32) {
    MaskEdit_SetLeft(m.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (m *TMaskEdit) Top() int32 {
    return MaskEdit_GetTop(m.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (m *TMaskEdit) SetTop(value int32) {
    MaskEdit_SetTop(m.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (m *TMaskEdit) Width() int32 {
    return MaskEdit_GetWidth(m.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (m *TMaskEdit) SetWidth(value int32) {
    MaskEdit_SetWidth(m.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (m *TMaskEdit) Height() int32 {
    return MaskEdit_GetHeight(m.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
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
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (m *TMaskEdit) Hint() string {
    return MaskEdit_GetHint(m.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (m *TMaskEdit) SetHint(value string) {
    MaskEdit_SetHint(m.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (m *TMaskEdit) Margins() *TMargins {
    return MarginsFromInst(MaskEdit_GetMargins(m.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (m *TMaskEdit) SetMargins(value *TMargins) {
    MaskEdit_SetMargins(m.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (m *TMaskEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(MaskEdit_GetCustomHint(m.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
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
// CN: 获取指定索引停靠客户端。
// EN: .
func (m *TMaskEdit) DockClients(Index int32) *TControl {
    return ControlFromInst(MaskEdit_GetDockClients(m.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (m *TMaskEdit) Controls(Index int32) *TControl {
    return ControlFromInst(MaskEdit_GetControls(m.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (m *TMaskEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(MaskEdit_GetComponents(m.instance, AIndex))
}

