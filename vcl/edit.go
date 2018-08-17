
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

type TEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewEdit
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewEdit(owner IComponent) *TEdit {
    e := new(TEdit)
    e.instance = Edit_Create(CheckPtr(owner))
    e.ptr = unsafe.Pointer(e.instance)
    return e
}

// EditFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func EditFromInst(inst uintptr) *TEdit {
    e := new(TEdit)
    e.instance = inst
    e.ptr = unsafe.Pointer(inst)
    return e
}

// EditFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func EditFromObj(obj IObject) *TEdit {
    e := new(TEdit)
    e.instance = CheckPtr(obj)
    e.ptr = unsafe.Pointer(e.instance)
    return e
}

// EditFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func EditFromUnsafePointer(ptr unsafe.Pointer) *TEdit {
    e := new(TEdit)
    e.instance = uintptr(ptr)
    e.ptr = ptr
    return e
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (e *TEdit) Free() {
    if e.instance != 0 {
        Edit_Free(e.instance)
        e.instance = 0
        e.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (e *TEdit) Instance() uintptr {
    return e.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (e *TEdit) UnsafeAddr() unsafe.Pointer {
    return e.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (e *TEdit) IsValid() bool {
    return e.instance != 0
}

// TEditClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TEditClass() TClass {
    return Edit_StaticClassType()
}

// Clear
// CN: 清除。
// EN: .
func (e *TEdit) Clear() {
    Edit_Clear(e.instance)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (e *TEdit) ClearSelection() {
    Edit_ClearSelection(e.instance)
}

// CopyToClipboard
// CN: 复制到粘贴板。
// EN: .
func (e *TEdit) CopyToClipboard() {
    Edit_CopyToClipboard(e.instance)
}

// CutToClipboard
// CN: 剪切到粘贴板。
// EN: .
func (e *TEdit) CutToClipboard() {
    Edit_CutToClipboard(e.instance)
}

// PasteFromClipboard
// CN: 从剪切板粘贴。
// EN: .
func (e *TEdit) PasteFromClipboard() {
    Edit_PasteFromClipboard(e.instance)
}

// Undo
// CN: 撤销上一次操作。
// EN: .
func (e *TEdit) Undo() {
    Edit_Undo(e.instance)
}

// ClearUndo
// CN: 清除撤销。
// EN: .
func (e *TEdit) ClearUndo() {
    Edit_ClearUndo(e.instance)
}

// SelectAll
// CN: 全选。
// EN: .
func (e *TEdit) SelectAll() {
    Edit_SelectAll(e.instance)
}

// GetSelTextBuf
func (e *TEdit) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return Edit_GetSelTextBuf(e.instance, Buffer , BufSize)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (e *TEdit) CanFocus() bool {
    return Edit_CanFocus(e.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (e *TEdit) ContainsControl(Control IControl) bool {
    return Edit_ContainsControl(e.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (e *TEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(Edit_ControlAtPos(e.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (e *TEdit) DisableAlign() {
    Edit_DisableAlign(e.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (e *TEdit) EnableAlign() {
    Edit_EnableAlign(e.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (e *TEdit) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(Edit_FindChildControl(e.instance, ControlName))
}

// FlipChildren
func (e *TEdit) FlipChildren(AllLevels bool) {
    Edit_FlipChildren(e.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (e *TEdit) Focused() bool {
    return Edit_Focused(e.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (e *TEdit) HandleAllocated() bool {
    return Edit_HandleAllocated(e.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (e *TEdit) InsertControl(AControl IControl) {
    Edit_InsertControl(e.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (e *TEdit) Invalidate() {
    Edit_Invalidate(e.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (e *TEdit) PaintTo(DC HDC, X int32, Y int32) {
    Edit_PaintTo(e.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (e *TEdit) RemoveControl(AControl IControl) {
    Edit_RemoveControl(e.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (e *TEdit) Realign() {
    Edit_Realign(e.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (e *TEdit) Repaint() {
    Edit_Repaint(e.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (e *TEdit) ScaleBy(M int32, D int32) {
    Edit_ScaleBy(e.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (e *TEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    Edit_ScrollBy(e.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (e *TEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Edit_SetBounds(e.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (e *TEdit) SetFocus() {
    Edit_SetFocus(e.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (e *TEdit) Update() {
    Edit_Update(e.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (e *TEdit) UpdateControlState() {
    Edit_UpdateControlState(e.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (e *TEdit) BringToFront() {
    Edit_BringToFront(e.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (e *TEdit) ClientToScreen(Point TPoint) TPoint {
    return Edit_ClientToScreen(e.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (e *TEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Edit_ClientToParent(e.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (e *TEdit) Dragging() bool {
    return Edit_Dragging(e.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (e *TEdit) HasParent() bool {
    return Edit_HasParent(e.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (e *TEdit) Hide() {
    Edit_Hide(e.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (e *TEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Edit_Perform(e.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (e *TEdit) Refresh() {
    Edit_Refresh(e.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (e *TEdit) ScreenToClient(Point TPoint) TPoint {
    return Edit_ScreenToClient(e.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (e *TEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Edit_ParentToClient(e.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (e *TEdit) SendToBack() {
    Edit_SendToBack(e.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (e *TEdit) Show() {
    Edit_Show(e.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (e *TEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Edit_GetTextBuf(e.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (e *TEdit) GetTextLen() int32 {
    return Edit_GetTextLen(e.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (e *TEdit) SetTextBuf(Buffer string) {
    Edit_SetTextBuf(e.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (e *TEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Edit_FindComponent(e.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (e *TEdit) GetNamePath() string {
    return Edit_GetNamePath(e.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (e *TEdit) Assign(Source IObject) {
    Edit_Assign(e.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (e *TEdit) DisposeOf() {
    Edit_DisposeOf(e.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (e *TEdit) ClassType() TClass {
    return Edit_ClassType(e.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (e *TEdit) ClassName() string {
    return Edit_ClassName(e.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (e *TEdit) InstanceSize() int32 {
    return Edit_InstanceSize(e.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (e *TEdit) InheritsFrom(AClass TClass) bool {
    return Edit_InheritsFrom(e.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (e *TEdit) Equals(Obj IObject) bool {
    return Edit_Equals(e.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (e *TEdit) GetHashCode() int32 {
    return Edit_GetHashCode(e.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (e *TEdit) ToString() string {
    return Edit_ToString(e.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (e *TEdit) Align() TAlign {
    return Edit_GetAlign(e.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (e *TEdit) SetAlign(value TAlign) {
    Edit_SetAlign(e.instance, value)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (e *TEdit) Alignment() TAlignment {
    return Edit_GetAlignment(e.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (e *TEdit) SetAlignment(value TAlignment) {
    Edit_SetAlignment(e.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (e *TEdit) Anchors() TAnchors {
    return Edit_GetAnchors(e.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (e *TEdit) SetAnchors(value TAnchors) {
    Edit_SetAnchors(e.instance, value)
}

// AutoSelect
// CN: 获取自动选择。
// EN: .
func (e *TEdit) AutoSelect() bool {
    return Edit_GetAutoSelect(e.instance)
}

// SetAutoSelect
// CN: 设置自动选择。
// EN: .
func (e *TEdit) SetAutoSelect(value bool) {
    Edit_SetAutoSelect(e.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (e *TEdit) AutoSize() bool {
    return Edit_GetAutoSize(e.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (e *TEdit) SetAutoSize(value bool) {
    Edit_SetAutoSize(e.instance, value)
}

// BevelEdges
func (e *TEdit) BevelEdges() TBevelEdges {
    return Edit_GetBevelEdges(e.instance)
}

// SetBevelEdges
func (e *TEdit) SetBevelEdges(value TBevelEdges) {
    Edit_SetBevelEdges(e.instance, value)
}

// BevelInner
func (e *TEdit) BevelInner() TBevelCut {
    return Edit_GetBevelInner(e.instance)
}

// SetBevelInner
func (e *TEdit) SetBevelInner(value TBevelCut) {
    Edit_SetBevelInner(e.instance, value)
}

// BevelKind
func (e *TEdit) BevelKind() TBevelKind {
    return Edit_GetBevelKind(e.instance)
}

// SetBevelKind
func (e *TEdit) SetBevelKind(value TBevelKind) {
    Edit_SetBevelKind(e.instance, value)
}

// BevelOuter
func (e *TEdit) BevelOuter() TBevelCut {
    return Edit_GetBevelOuter(e.instance)
}

// SetBevelOuter
func (e *TEdit) SetBevelOuter(value TBevelCut) {
    Edit_SetBevelOuter(e.instance, value)
}

// BiDiMode
func (e *TEdit) BiDiMode() TBiDiMode {
    return Edit_GetBiDiMode(e.instance)
}

// SetBiDiMode
func (e *TEdit) SetBiDiMode(value TBiDiMode) {
    Edit_SetBiDiMode(e.instance, value)
}

// BorderStyle
// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (e *TEdit) BorderStyle() TBorderStyle {
    return Edit_GetBorderStyle(e.instance)
}

// SetBorderStyle
// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (e *TEdit) SetBorderStyle(value TBorderStyle) {
    Edit_SetBorderStyle(e.instance, value)
}

// CharCase
func (e *TEdit) CharCase() TEditCharCase {
    return Edit_GetCharCase(e.instance)
}

// SetCharCase
func (e *TEdit) SetCharCase(value TEditCharCase) {
    Edit_SetCharCase(e.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (e *TEdit) Color() TColor {
    return Edit_GetColor(e.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (e *TEdit) SetColor(value TColor) {
    Edit_SetColor(e.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (e *TEdit) DoubleBuffered() bool {
    return Edit_GetDoubleBuffered(e.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (e *TEdit) SetDoubleBuffered(value bool) {
    Edit_SetDoubleBuffered(e.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (e *TEdit) DragCursor() TCursor {
    return Edit_GetDragCursor(e.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (e *TEdit) SetDragCursor(value TCursor) {
    Edit_SetDragCursor(e.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (e *TEdit) DragKind() TDragKind {
    return Edit_GetDragKind(e.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (e *TEdit) SetDragKind(value TDragKind) {
    Edit_SetDragKind(e.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (e *TEdit) DragMode() TDragMode {
    return Edit_GetDragMode(e.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (e *TEdit) SetDragMode(value TDragMode) {
    Edit_SetDragMode(e.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (e *TEdit) Enabled() bool {
    return Edit_GetEnabled(e.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (e *TEdit) SetEnabled(value bool) {
    Edit_SetEnabled(e.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (e *TEdit) Font() *TFont {
    return FontFromInst(Edit_GetFont(e.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (e *TEdit) SetFont(value *TFont) {
    Edit_SetFont(e.instance, CheckPtr(value))
}

// HideSelection
// CN: 获取隐藏选择。
// EN: .
func (e *TEdit) HideSelection() bool {
    return Edit_GetHideSelection(e.instance)
}

// SetHideSelection
// CN: 设置隐藏选择。
// EN: .
func (e *TEdit) SetHideSelection(value bool) {
    Edit_SetHideSelection(e.instance, value)
}

// MaxLength
// CN: 获取最大长度。
// EN: .
func (e *TEdit) MaxLength() int32 {
    return Edit_GetMaxLength(e.instance)
}

// SetMaxLength
// CN: 设置最大长度。
// EN: .
func (e *TEdit) SetMaxLength(value int32) {
    Edit_SetMaxLength(e.instance, value)
}

// NumbersOnly
// CN: 获取只能输入数字。
// EN: .
func (e *TEdit) NumbersOnly() bool {
    return Edit_GetNumbersOnly(e.instance)
}

// SetNumbersOnly
// CN: 设置只能输入数字。
// EN: .
func (e *TEdit) SetNumbersOnly(value bool) {
    Edit_SetNumbersOnly(e.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (e *TEdit) ParentColor() bool {
    return Edit_GetParentColor(e.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (e *TEdit) SetParentColor(value bool) {
    Edit_SetParentColor(e.instance, value)
}

// ParentCtl3D
func (e *TEdit) ParentCtl3D() bool {
    return Edit_GetParentCtl3D(e.instance)
}

// SetParentCtl3D
func (e *TEdit) SetParentCtl3D(value bool) {
    Edit_SetParentCtl3D(e.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (e *TEdit) ParentDoubleBuffered() bool {
    return Edit_GetParentDoubleBuffered(e.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (e *TEdit) SetParentDoubleBuffered(value bool) {
    Edit_SetParentDoubleBuffered(e.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (e *TEdit) ParentFont() bool {
    return Edit_GetParentFont(e.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (e *TEdit) SetParentFont(value bool) {
    Edit_SetParentFont(e.instance, value)
}

// ParentShowHint
func (e *TEdit) ParentShowHint() bool {
    return Edit_GetParentShowHint(e.instance)
}

// SetParentShowHint
func (e *TEdit) SetParentShowHint(value bool) {
    Edit_SetParentShowHint(e.instance, value)
}

// PasswordChar
// CN: 获取密码掩码字符。
// EN: .
func (e *TEdit) PasswordChar() uint16 {
    return Edit_GetPasswordChar(e.instance)
}

// SetPasswordChar
// CN: 设置密码掩码字符。
// EN: .
func (e *TEdit) SetPasswordChar(value uint16) {
    Edit_SetPasswordChar(e.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (e *TEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Edit_GetPopupMenu(e.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (e *TEdit) SetPopupMenu(value IComponent) {
    Edit_SetPopupMenu(e.instance, CheckPtr(value))
}

// ReadOnly
// CN: 获取只读。
// EN: .
func (e *TEdit) ReadOnly() bool {
    return Edit_GetReadOnly(e.instance)
}

// SetReadOnly
// CN: 设置只读。
// EN: .
func (e *TEdit) SetReadOnly(value bool) {
    Edit_SetReadOnly(e.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (e *TEdit) ShowHint() bool {
    return Edit_GetShowHint(e.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (e *TEdit) SetShowHint(value bool) {
    Edit_SetShowHint(e.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (e *TEdit) TabOrder() TTabOrder {
    return Edit_GetTabOrder(e.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (e *TEdit) SetTabOrder(value TTabOrder) {
    Edit_SetTabOrder(e.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (e *TEdit) TabStop() bool {
    return Edit_GetTabStop(e.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (e *TEdit) SetTabStop(value bool) {
    Edit_SetTabStop(e.instance, value)
}

// Text
// CN: 获取文本。
// EN: .
func (e *TEdit) Text() string {
    return Edit_GetText(e.instance)
}

// SetText
// CN: 设置文本。
// EN: .
func (e *TEdit) SetText(value string) {
    Edit_SetText(e.instance, value)
}

// TextHint
// CN: 获取提示文本。
// EN: .
func (e *TEdit) TextHint() string {
    return Edit_GetTextHint(e.instance)
}

// SetTextHint
// CN: 设置提示文本。
// EN: .
func (e *TEdit) SetTextHint(value string) {
    Edit_SetTextHint(e.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (e *TEdit) Visible() bool {
    return Edit_GetVisible(e.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (e *TEdit) SetVisible(value bool) {
    Edit_SetVisible(e.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (e *TEdit) StyleElements() TStyleElements {
    return Edit_GetStyleElements(e.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (e *TEdit) SetStyleElements(value TStyleElements) {
    Edit_SetStyleElements(e.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (e *TEdit) SetOnChange(fn TNotifyEvent) {
    Edit_SetOnChange(e.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (e *TEdit) SetOnClick(fn TNotifyEvent) {
    Edit_SetOnClick(e.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (e *TEdit) SetOnContextPopup(fn TContextPopupEvent) {
    Edit_SetOnContextPopup(e.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (e *TEdit) SetOnDblClick(fn TNotifyEvent) {
    Edit_SetOnDblClick(e.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (e *TEdit) SetOnDragDrop(fn TDragDropEvent) {
    Edit_SetOnDragDrop(e.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (e *TEdit) SetOnDragOver(fn TDragOverEvent) {
    Edit_SetOnDragOver(e.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (e *TEdit) SetOnEndDock(fn TEndDragEvent) {
    Edit_SetOnEndDock(e.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (e *TEdit) SetOnEndDrag(fn TEndDragEvent) {
    Edit_SetOnEndDrag(e.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (e *TEdit) SetOnEnter(fn TNotifyEvent) {
    Edit_SetOnEnter(e.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (e *TEdit) SetOnExit(fn TNotifyEvent) {
    Edit_SetOnExit(e.instance, fn)
}

// SetOnGesture
func (e *TEdit) SetOnGesture(fn TGestureEvent) {
    Edit_SetOnGesture(e.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (e *TEdit) SetOnKeyDown(fn TKeyEvent) {
    Edit_SetOnKeyDown(e.instance, fn)
}

// SetOnKeyPress
func (e *TEdit) SetOnKeyPress(fn TKeyPressEvent) {
    Edit_SetOnKeyPress(e.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (e *TEdit) SetOnKeyUp(fn TKeyEvent) {
    Edit_SetOnKeyUp(e.instance, fn)
}

// SetOnMouseActivate
func (e *TEdit) SetOnMouseActivate(fn TMouseActivateEvent) {
    Edit_SetOnMouseActivate(e.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (e *TEdit) SetOnMouseDown(fn TMouseEvent) {
    Edit_SetOnMouseDown(e.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (e *TEdit) SetOnMouseEnter(fn TNotifyEvent) {
    Edit_SetOnMouseEnter(e.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (e *TEdit) SetOnMouseLeave(fn TNotifyEvent) {
    Edit_SetOnMouseLeave(e.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (e *TEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    Edit_SetOnMouseMove(e.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (e *TEdit) SetOnMouseUp(fn TMouseEvent) {
    Edit_SetOnMouseUp(e.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (e *TEdit) SetOnStartDock(fn TStartDockEvent) {
    Edit_SetOnStartDock(e.instance, fn)
}

// CanUndo
// CN: 获取能否撤销。
// EN: .
func (e *TEdit) CanUndo() bool {
    return Edit_GetCanUndo(e.instance)
}

// Modified
// CN: 获取修改。
// EN: Get modified.
func (e *TEdit) Modified() bool {
    return Edit_GetModified(e.instance)
}

// SetModified
// CN: 设置修改。
// EN: Set modified.
func (e *TEdit) SetModified(value bool) {
    Edit_SetModified(e.instance, value)
}

// SelLength
// CN: 获取选择的长度。
// EN: .
func (e *TEdit) SelLength() int32 {
    return Edit_GetSelLength(e.instance)
}

// SetSelLength
// CN: 设置选择的长度。
// EN: .
func (e *TEdit) SetSelLength(value int32) {
    Edit_SetSelLength(e.instance, value)
}

// SelStart
// CN: 获取选择的启始位置。
// EN: .
func (e *TEdit) SelStart() int32 {
    return Edit_GetSelStart(e.instance)
}

// SetSelStart
// CN: 设置选择的启始位置。
// EN: .
func (e *TEdit) SetSelStart(value int32) {
    Edit_SetSelStart(e.instance, value)
}

// SelText
// CN: 获取选择的文本。
// EN: .
func (e *TEdit) SelText() string {
    return Edit_GetSelText(e.instance)
}

// SetSelText
// CN: 设置选择的文本。
// EN: .
func (e *TEdit) SetSelText(value string) {
    Edit_SetSelText(e.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (e *TEdit) DockClientCount() int32 {
    return Edit_GetDockClientCount(e.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (e *TEdit) DockSite() bool {
    return Edit_GetDockSite(e.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (e *TEdit) SetDockSite(value bool) {
    Edit_SetDockSite(e.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (e *TEdit) AlignDisabled() bool {
    return Edit_GetAlignDisabled(e.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (e *TEdit) MouseInClient() bool {
    return Edit_GetMouseInClient(e.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (e *TEdit) VisibleDockClientCount() int32 {
    return Edit_GetVisibleDockClientCount(e.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (e *TEdit) Brush() *TBrush {
    return BrushFromInst(Edit_GetBrush(e.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (e *TEdit) ControlCount() int32 {
    return Edit_GetControlCount(e.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (e *TEdit) Handle() HWND {
    return Edit_GetHandle(e.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (e *TEdit) ParentWindow() HWND {
    return Edit_GetParentWindow(e.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (e *TEdit) SetParentWindow(value HWND) {
    Edit_SetParentWindow(e.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (e *TEdit) UseDockManager() bool {
    return Edit_GetUseDockManager(e.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (e *TEdit) SetUseDockManager(value bool) {
    Edit_SetUseDockManager(e.instance, value)
}

// Action
func (e *TEdit) Action() *TAction {
    return ActionFromInst(Edit_GetAction(e.instance))
}

// SetAction
func (e *TEdit) SetAction(value IComponent) {
    Edit_SetAction(e.instance, CheckPtr(value))
}

// BoundsRect
func (e *TEdit) BoundsRect() TRect {
    return Edit_GetBoundsRect(e.instance)
}

// SetBoundsRect
func (e *TEdit) SetBoundsRect(value TRect) {
    Edit_SetBoundsRect(e.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (e *TEdit) ClientHeight() int32 {
    return Edit_GetClientHeight(e.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (e *TEdit) SetClientHeight(value int32) {
    Edit_SetClientHeight(e.instance, value)
}

// ClientOrigin
func (e *TEdit) ClientOrigin() TPoint {
    return Edit_GetClientOrigin(e.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (e *TEdit) ClientRect() TRect {
    return Edit_GetClientRect(e.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (e *TEdit) ClientWidth() int32 {
    return Edit_GetClientWidth(e.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (e *TEdit) SetClientWidth(value int32) {
    Edit_SetClientWidth(e.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (e *TEdit) ControlState() TControlState {
    return Edit_GetControlState(e.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (e *TEdit) SetControlState(value TControlState) {
    Edit_SetControlState(e.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (e *TEdit) ControlStyle() TControlStyle {
    return Edit_GetControlStyle(e.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (e *TEdit) SetControlStyle(value TControlStyle) {
    Edit_SetControlStyle(e.instance, value)
}

// ExplicitLeft
func (e *TEdit) ExplicitLeft() int32 {
    return Edit_GetExplicitLeft(e.instance)
}

// ExplicitTop
func (e *TEdit) ExplicitTop() int32 {
    return Edit_GetExplicitTop(e.instance)
}

// ExplicitWidth
func (e *TEdit) ExplicitWidth() int32 {
    return Edit_GetExplicitWidth(e.instance)
}

// ExplicitHeight
func (e *TEdit) ExplicitHeight() int32 {
    return Edit_GetExplicitHeight(e.instance)
}

// Floating
func (e *TEdit) Floating() bool {
    return Edit_GetFloating(e.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (e *TEdit) Parent() *TWinControl {
    return WinControlFromInst(Edit_GetParent(e.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (e *TEdit) SetParent(value IWinControl) {
    Edit_SetParent(e.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (e *TEdit) AlignWithMargins() bool {
    return Edit_GetAlignWithMargins(e.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (e *TEdit) SetAlignWithMargins(value bool) {
    Edit_SetAlignWithMargins(e.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (e *TEdit) Left() int32 {
    return Edit_GetLeft(e.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (e *TEdit) SetLeft(value int32) {
    Edit_SetLeft(e.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (e *TEdit) Top() int32 {
    return Edit_GetTop(e.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (e *TEdit) SetTop(value int32) {
    Edit_SetTop(e.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (e *TEdit) Width() int32 {
    return Edit_GetWidth(e.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (e *TEdit) SetWidth(value int32) {
    Edit_SetWidth(e.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (e *TEdit) Height() int32 {
    return Edit_GetHeight(e.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (e *TEdit) SetHeight(value int32) {
    Edit_SetHeight(e.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (e *TEdit) Cursor() TCursor {
    return Edit_GetCursor(e.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (e *TEdit) SetCursor(value TCursor) {
    Edit_SetCursor(e.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (e *TEdit) Hint() string {
    return Edit_GetHint(e.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (e *TEdit) SetHint(value string) {
    Edit_SetHint(e.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (e *TEdit) Margins() *TMargins {
    return MarginsFromInst(Edit_GetMargins(e.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (e *TEdit) SetMargins(value *TMargins) {
    Edit_SetMargins(e.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (e *TEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(Edit_GetCustomHint(e.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (e *TEdit) SetCustomHint(value IComponent) {
    Edit_SetCustomHint(e.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (e *TEdit) ComponentCount() int32 {
    return Edit_GetComponentCount(e.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (e *TEdit) ComponentIndex() int32 {
    return Edit_GetComponentIndex(e.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (e *TEdit) SetComponentIndex(value int32) {
    Edit_SetComponentIndex(e.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (e *TEdit) Owner() *TComponent {
    return ComponentFromInst(Edit_GetOwner(e.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (e *TEdit) Name() string {
    return Edit_GetName(e.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (e *TEdit) SetName(value string) {
    Edit_SetName(e.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (e *TEdit) Tag() int {
    return Edit_GetTag(e.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (e *TEdit) SetTag(value int) {
    Edit_SetTag(e.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (e *TEdit) DockClients(Index int32) *TControl {
    return ControlFromInst(Edit_GetDockClients(e.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (e *TEdit) Controls(Index int32) *TControl {
    return ControlFromInst(Edit_GetControls(e.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (e *TEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Edit_GetComponents(e.instance, AIndex))
}

