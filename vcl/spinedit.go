
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

type TSpinEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewSpinEdit
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewSpinEdit(owner IComponent) *TSpinEdit {
    s := new(TSpinEdit)
    s.instance = SpinEdit_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SpinEditFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func SpinEditFromInst(inst uintptr) *TSpinEdit {
    s := new(TSpinEdit)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// SpinEditFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func SpinEditFromObj(obj IObject) *TSpinEdit {
    s := new(TSpinEdit)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SpinEditFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func SpinEditFromUnsafePointer(ptr unsafe.Pointer) *TSpinEdit {
    s := new(TSpinEdit)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TSpinEdit) Free() {
    if s.instance != 0 {
        SpinEdit_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TSpinEdit) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TSpinEdit) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TSpinEdit) IsValid() bool {
    return s.instance != 0
}

// TSpinEditClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TSpinEditClass() TClass {
    return SpinEdit_StaticClassType()
}

// Clear
// CN: 清除。
// EN: .
func (s *TSpinEdit) Clear() {
    SpinEdit_Clear(s.instance)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (s *TSpinEdit) ClearSelection() {
    SpinEdit_ClearSelection(s.instance)
}

// CopyToClipboard
// CN: 复制到粘贴板。
// EN: .
func (s *TSpinEdit) CopyToClipboard() {
    SpinEdit_CopyToClipboard(s.instance)
}

// CutToClipboard
// CN: 剪切到粘贴板。
// EN: .
func (s *TSpinEdit) CutToClipboard() {
    SpinEdit_CutToClipboard(s.instance)
}

// PasteFromClipboard
// CN: 从剪切板粘贴。
// EN: .
func (s *TSpinEdit) PasteFromClipboard() {
    SpinEdit_PasteFromClipboard(s.instance)
}

// Undo
// CN: 撤销上一次操作。
// EN: .
func (s *TSpinEdit) Undo() {
    SpinEdit_Undo(s.instance)
}

// ClearUndo
// CN: 清除撤销。
// EN: .
func (s *TSpinEdit) ClearUndo() {
    SpinEdit_ClearUndo(s.instance)
}

// SelectAll
// CN: 全选。
// EN: .
func (s *TSpinEdit) SelectAll() {
    SpinEdit_SelectAll(s.instance)
}

// GetSelTextBuf
func (s *TSpinEdit) GetSelTextBuf(Buffer *string, BufSize int32) int32 {
    return SpinEdit_GetSelTextBuf(s.instance, Buffer , BufSize)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (s *TSpinEdit) CanFocus() bool {
    return SpinEdit_CanFocus(s.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (s *TSpinEdit) ContainsControl(Control IControl) bool {
    return SpinEdit_ContainsControl(s.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (s *TSpinEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(SpinEdit_ControlAtPos(s.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (s *TSpinEdit) DisableAlign() {
    SpinEdit_DisableAlign(s.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (s *TSpinEdit) EnableAlign() {
    SpinEdit_EnableAlign(s.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (s *TSpinEdit) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(SpinEdit_FindChildControl(s.instance, ControlName))
}

// FlipChildren
func (s *TSpinEdit) FlipChildren(AllLevels bool) {
    SpinEdit_FlipChildren(s.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (s *TSpinEdit) Focused() bool {
    return SpinEdit_Focused(s.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (s *TSpinEdit) HandleAllocated() bool {
    return SpinEdit_HandleAllocated(s.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (s *TSpinEdit) InsertControl(AControl IControl) {
    SpinEdit_InsertControl(s.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (s *TSpinEdit) Invalidate() {
    SpinEdit_Invalidate(s.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (s *TSpinEdit) PaintTo(DC HDC, X int32, Y int32) {
    SpinEdit_PaintTo(s.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (s *TSpinEdit) RemoveControl(AControl IControl) {
    SpinEdit_RemoveControl(s.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (s *TSpinEdit) Realign() {
    SpinEdit_Realign(s.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (s *TSpinEdit) Repaint() {
    SpinEdit_Repaint(s.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (s *TSpinEdit) ScaleBy(M int32, D int32) {
    SpinEdit_ScaleBy(s.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (s *TSpinEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    SpinEdit_ScrollBy(s.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (s *TSpinEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    SpinEdit_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (s *TSpinEdit) SetFocus() {
    SpinEdit_SetFocus(s.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (s *TSpinEdit) Update() {
    SpinEdit_Update(s.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (s *TSpinEdit) UpdateControlState() {
    SpinEdit_UpdateControlState(s.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (s *TSpinEdit) BringToFront() {
    SpinEdit_BringToFront(s.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (s *TSpinEdit) ClientToScreen(Point TPoint) TPoint {
    return SpinEdit_ClientToScreen(s.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (s *TSpinEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return SpinEdit_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (s *TSpinEdit) Dragging() bool {
    return SpinEdit_Dragging(s.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TSpinEdit) HasParent() bool {
    return SpinEdit_HasParent(s.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (s *TSpinEdit) Hide() {
    SpinEdit_Hide(s.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (s *TSpinEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return SpinEdit_Perform(s.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (s *TSpinEdit) Refresh() {
    SpinEdit_Refresh(s.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (s *TSpinEdit) ScreenToClient(Point TPoint) TPoint {
    return SpinEdit_ScreenToClient(s.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (s *TSpinEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return SpinEdit_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (s *TSpinEdit) SendToBack() {
    SpinEdit_SendToBack(s.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (s *TSpinEdit) Show() {
    SpinEdit_Show(s.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (s *TSpinEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return SpinEdit_GetTextBuf(s.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (s *TSpinEdit) GetTextLen() int32 {
    return SpinEdit_GetTextLen(s.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (s *TSpinEdit) SetTextBuf(Buffer string) {
    SpinEdit_SetTextBuf(s.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TSpinEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(SpinEdit_FindComponent(s.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TSpinEdit) GetNamePath() string {
    return SpinEdit_GetNamePath(s.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TSpinEdit) Assign(Source IObject) {
    SpinEdit_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TSpinEdit) DisposeOf() {
    SpinEdit_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TSpinEdit) ClassType() TClass {
    return SpinEdit_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TSpinEdit) ClassName() string {
    return SpinEdit_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TSpinEdit) InstanceSize() int32 {
    return SpinEdit_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TSpinEdit) InheritsFrom(AClass TClass) bool {
    return SpinEdit_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TSpinEdit) Equals(Obj IObject) bool {
    return SpinEdit_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TSpinEdit) GetHashCode() int32 {
    return SpinEdit_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TSpinEdit) ToString() string {
    return SpinEdit_ToString(s.instance)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (s *TSpinEdit) Anchors() TAnchors {
    return SpinEdit_GetAnchors(s.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (s *TSpinEdit) SetAnchors(value TAnchors) {
    SpinEdit_SetAnchors(s.instance, value)
}

// AutoSelect
// CN: 获取自动选择。
// EN: .
func (s *TSpinEdit) AutoSelect() bool {
    return SpinEdit_GetAutoSelect(s.instance)
}

// SetAutoSelect
// CN: 设置自动选择。
// EN: .
func (s *TSpinEdit) SetAutoSelect(value bool) {
    SpinEdit_SetAutoSelect(s.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (s *TSpinEdit) AutoSize() bool {
    return SpinEdit_GetAutoSize(s.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (s *TSpinEdit) SetAutoSize(value bool) {
    SpinEdit_SetAutoSize(s.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (s *TSpinEdit) Color() TColor {
    return SpinEdit_GetColor(s.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (s *TSpinEdit) SetColor(value TColor) {
    SpinEdit_SetColor(s.instance, value)
}

// Constraints
func (s *TSpinEdit) Constraints() *TSizeConstraints {
    return SizeConstraintsFromInst(SpinEdit_GetConstraints(s.instance))
}

// SetConstraints
func (s *TSpinEdit) SetConstraints(value *TSizeConstraints) {
    SpinEdit_SetConstraints(s.instance, CheckPtr(value))
}

// Ctl3D
func (s *TSpinEdit) Ctl3D() bool {
    return SpinEdit_GetCtl3D(s.instance)
}

// SetCtl3D
func (s *TSpinEdit) SetCtl3D(value bool) {
    SpinEdit_SetCtl3D(s.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (s *TSpinEdit) DragCursor() TCursor {
    return SpinEdit_GetDragCursor(s.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (s *TSpinEdit) SetDragCursor(value TCursor) {
    SpinEdit_SetDragCursor(s.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (s *TSpinEdit) DragMode() TDragMode {
    return SpinEdit_GetDragMode(s.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (s *TSpinEdit) SetDragMode(value TDragMode) {
    SpinEdit_SetDragMode(s.instance, value)
}

// EditorEnabled
func (s *TSpinEdit) EditorEnabled() bool {
    return SpinEdit_GetEditorEnabled(s.instance)
}

// SetEditorEnabled
func (s *TSpinEdit) SetEditorEnabled(value bool) {
    SpinEdit_SetEditorEnabled(s.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TSpinEdit) Enabled() bool {
    return SpinEdit_GetEnabled(s.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TSpinEdit) SetEnabled(value bool) {
    SpinEdit_SetEnabled(s.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (s *TSpinEdit) Font() *TFont {
    return FontFromInst(SpinEdit_GetFont(s.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (s *TSpinEdit) SetFont(value *TFont) {
    SpinEdit_SetFont(s.instance, CheckPtr(value))
}

// Increment
func (s *TSpinEdit) Increment() int32 {
    return SpinEdit_GetIncrement(s.instance)
}

// SetIncrement
func (s *TSpinEdit) SetIncrement(value int32) {
    SpinEdit_SetIncrement(s.instance, value)
}

// MaxLength
// CN: 获取最大长度。
// EN: .
func (s *TSpinEdit) MaxLength() int32 {
    return SpinEdit_GetMaxLength(s.instance)
}

// SetMaxLength
// CN: 设置最大长度。
// EN: .
func (s *TSpinEdit) SetMaxLength(value int32) {
    SpinEdit_SetMaxLength(s.instance, value)
}

// MaxValue
func (s *TSpinEdit) MaxValue() int32 {
    return SpinEdit_GetMaxValue(s.instance)
}

// SetMaxValue
func (s *TSpinEdit) SetMaxValue(value int32) {
    SpinEdit_SetMaxValue(s.instance, value)
}

// MinValue
func (s *TSpinEdit) MinValue() int32 {
    return SpinEdit_GetMinValue(s.instance)
}

// SetMinValue
func (s *TSpinEdit) SetMinValue(value int32) {
    SpinEdit_SetMinValue(s.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (s *TSpinEdit) ParentColor() bool {
    return SpinEdit_GetParentColor(s.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (s *TSpinEdit) SetParentColor(value bool) {
    SpinEdit_SetParentColor(s.instance, value)
}

// ParentCtl3D
func (s *TSpinEdit) ParentCtl3D() bool {
    return SpinEdit_GetParentCtl3D(s.instance)
}

// SetParentCtl3D
func (s *TSpinEdit) SetParentCtl3D(value bool) {
    SpinEdit_SetParentCtl3D(s.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (s *TSpinEdit) ParentFont() bool {
    return SpinEdit_GetParentFont(s.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (s *TSpinEdit) SetParentFont(value bool) {
    SpinEdit_SetParentFont(s.instance, value)
}

// ParentShowHint
func (s *TSpinEdit) ParentShowHint() bool {
    return SpinEdit_GetParentShowHint(s.instance)
}

// SetParentShowHint
func (s *TSpinEdit) SetParentShowHint(value bool) {
    SpinEdit_SetParentShowHint(s.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (s *TSpinEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(SpinEdit_GetPopupMenu(s.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (s *TSpinEdit) SetPopupMenu(value IComponent) {
    SpinEdit_SetPopupMenu(s.instance, CheckPtr(value))
}

// ReadOnly
// CN: 获取只读。
// EN: .
func (s *TSpinEdit) ReadOnly() bool {
    return SpinEdit_GetReadOnly(s.instance)
}

// SetReadOnly
// CN: 设置只读。
// EN: .
func (s *TSpinEdit) SetReadOnly(value bool) {
    SpinEdit_SetReadOnly(s.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (s *TSpinEdit) ShowHint() bool {
    return SpinEdit_GetShowHint(s.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (s *TSpinEdit) SetShowHint(value bool) {
    SpinEdit_SetShowHint(s.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (s *TSpinEdit) TabOrder() TTabOrder {
    return SpinEdit_GetTabOrder(s.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (s *TSpinEdit) SetTabOrder(value TTabOrder) {
    SpinEdit_SetTabOrder(s.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (s *TSpinEdit) TabStop() bool {
    return SpinEdit_GetTabStop(s.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (s *TSpinEdit) SetTabStop(value bool) {
    SpinEdit_SetTabStop(s.instance, value)
}

// Value
func (s *TSpinEdit) Value() int32 {
    return SpinEdit_GetValue(s.instance)
}

// SetValue
func (s *TSpinEdit) SetValue(value int32) {
    SpinEdit_SetValue(s.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TSpinEdit) Visible() bool {
    return SpinEdit_GetVisible(s.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TSpinEdit) SetVisible(value bool) {
    SpinEdit_SetVisible(s.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (s *TSpinEdit) SetOnChange(fn TNotifyEvent) {
    SpinEdit_SetOnChange(s.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (s *TSpinEdit) SetOnClick(fn TNotifyEvent) {
    SpinEdit_SetOnClick(s.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (s *TSpinEdit) SetOnDblClick(fn TNotifyEvent) {
    SpinEdit_SetOnDblClick(s.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (s *TSpinEdit) SetOnDragDrop(fn TDragDropEvent) {
    SpinEdit_SetOnDragDrop(s.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (s *TSpinEdit) SetOnDragOver(fn TDragOverEvent) {
    SpinEdit_SetOnDragOver(s.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (s *TSpinEdit) SetOnEndDrag(fn TEndDragEvent) {
    SpinEdit_SetOnEndDrag(s.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (s *TSpinEdit) SetOnEnter(fn TNotifyEvent) {
    SpinEdit_SetOnEnter(s.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (s *TSpinEdit) SetOnExit(fn TNotifyEvent) {
    SpinEdit_SetOnExit(s.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (s *TSpinEdit) SetOnKeyDown(fn TKeyEvent) {
    SpinEdit_SetOnKeyDown(s.instance, fn)
}

// SetOnKeyPress
func (s *TSpinEdit) SetOnKeyPress(fn TKeyPressEvent) {
    SpinEdit_SetOnKeyPress(s.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (s *TSpinEdit) SetOnKeyUp(fn TKeyEvent) {
    SpinEdit_SetOnKeyUp(s.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (s *TSpinEdit) SetOnMouseDown(fn TMouseEvent) {
    SpinEdit_SetOnMouseDown(s.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (s *TSpinEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    SpinEdit_SetOnMouseMove(s.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (s *TSpinEdit) SetOnMouseUp(fn TMouseEvent) {
    SpinEdit_SetOnMouseUp(s.instance, fn)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (s *TSpinEdit) Alignment() TAlignment {
    return SpinEdit_GetAlignment(s.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (s *TSpinEdit) SetAlignment(value TAlignment) {
    SpinEdit_SetAlignment(s.instance, value)
}

// CanUndo
// CN: 获取能否撤销。
// EN: .
func (s *TSpinEdit) CanUndo() bool {
    return SpinEdit_GetCanUndo(s.instance)
}

// Modified
// CN: 获取修改。
// EN: Get modified.
func (s *TSpinEdit) Modified() bool {
    return SpinEdit_GetModified(s.instance)
}

// SetModified
// CN: 设置修改。
// EN: Set modified.
func (s *TSpinEdit) SetModified(value bool) {
    SpinEdit_SetModified(s.instance, value)
}

// SelLength
// CN: 获取选择的长度。
// EN: .
func (s *TSpinEdit) SelLength() int32 {
    return SpinEdit_GetSelLength(s.instance)
}

// SetSelLength
// CN: 设置选择的长度。
// EN: .
func (s *TSpinEdit) SetSelLength(value int32) {
    SpinEdit_SetSelLength(s.instance, value)
}

// SelStart
// CN: 获取选择的启始位置。
// EN: .
func (s *TSpinEdit) SelStart() int32 {
    return SpinEdit_GetSelStart(s.instance)
}

// SetSelStart
// CN: 设置选择的启始位置。
// EN: .
func (s *TSpinEdit) SetSelStart(value int32) {
    SpinEdit_SetSelStart(s.instance, value)
}

// SelText
// CN: 获取选择的文本。
// EN: .
func (s *TSpinEdit) SelText() string {
    return SpinEdit_GetSelText(s.instance)
}

// SetSelText
// CN: 设置选择的文本。
// EN: .
func (s *TSpinEdit) SetSelText(value string) {
    SpinEdit_SetSelText(s.instance, value)
}

// Text
// CN: 获取文本。
// EN: .
func (s *TSpinEdit) Text() string {
    if !IsloadedLcl {
        strLen := s.GetTextLen()
        var buffStr string
        if strLen != 0 {
            s.GetTextBuf(&buffStr, strLen + 1)
        }
        return buffStr
    } else { 
        return SpinEdit_GetText(s.instance)
    }
}

// SetText
// CN: 设置文本。
// EN: .
func (s *TSpinEdit) SetText(value string) {
    SpinEdit_SetText(s.instance, value)
}

// TextHint
// CN: 获取提示文本。
// EN: .
func (s *TSpinEdit) TextHint() string {
    return SpinEdit_GetTextHint(s.instance)
}

// SetTextHint
// CN: 设置提示文本。
// EN: .
func (s *TSpinEdit) SetTextHint(value string) {
    SpinEdit_SetTextHint(s.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (s *TSpinEdit) DockClientCount() int32 {
    return SpinEdit_GetDockClientCount(s.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (s *TSpinEdit) DockSite() bool {
    return SpinEdit_GetDockSite(s.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (s *TSpinEdit) SetDockSite(value bool) {
    SpinEdit_SetDockSite(s.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (s *TSpinEdit) DoubleBuffered() bool {
    return SpinEdit_GetDoubleBuffered(s.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (s *TSpinEdit) SetDoubleBuffered(value bool) {
    SpinEdit_SetDoubleBuffered(s.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (s *TSpinEdit) AlignDisabled() bool {
    return SpinEdit_GetAlignDisabled(s.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (s *TSpinEdit) MouseInClient() bool {
    return SpinEdit_GetMouseInClient(s.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (s *TSpinEdit) VisibleDockClientCount() int32 {
    return SpinEdit_GetVisibleDockClientCount(s.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (s *TSpinEdit) Brush() *TBrush {
    return BrushFromInst(SpinEdit_GetBrush(s.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (s *TSpinEdit) ControlCount() int32 {
    return SpinEdit_GetControlCount(s.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (s *TSpinEdit) Handle() HWND {
    return SpinEdit_GetHandle(s.instance)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (s *TSpinEdit) ParentDoubleBuffered() bool {
    return SpinEdit_GetParentDoubleBuffered(s.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (s *TSpinEdit) SetParentDoubleBuffered(value bool) {
    SpinEdit_SetParentDoubleBuffered(s.instance, value)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (s *TSpinEdit) ParentWindow() HWND {
    return SpinEdit_GetParentWindow(s.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (s *TSpinEdit) SetParentWindow(value HWND) {
    SpinEdit_SetParentWindow(s.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (s *TSpinEdit) UseDockManager() bool {
    return SpinEdit_GetUseDockManager(s.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (s *TSpinEdit) SetUseDockManager(value bool) {
    SpinEdit_SetUseDockManager(s.instance, value)
}

// Action
func (s *TSpinEdit) Action() *TAction {
    return ActionFromInst(SpinEdit_GetAction(s.instance))
}

// SetAction
func (s *TSpinEdit) SetAction(value IComponent) {
    SpinEdit_SetAction(s.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (s *TSpinEdit) Align() TAlign {
    return SpinEdit_GetAlign(s.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (s *TSpinEdit) SetAlign(value TAlign) {
    SpinEdit_SetAlign(s.instance, value)
}

// BiDiMode
func (s *TSpinEdit) BiDiMode() TBiDiMode {
    return SpinEdit_GetBiDiMode(s.instance)
}

// SetBiDiMode
func (s *TSpinEdit) SetBiDiMode(value TBiDiMode) {
    SpinEdit_SetBiDiMode(s.instance, value)
}

// BoundsRect
func (s *TSpinEdit) BoundsRect() TRect {
    return SpinEdit_GetBoundsRect(s.instance)
}

// SetBoundsRect
func (s *TSpinEdit) SetBoundsRect(value TRect) {
    SpinEdit_SetBoundsRect(s.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (s *TSpinEdit) ClientHeight() int32 {
    return SpinEdit_GetClientHeight(s.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (s *TSpinEdit) SetClientHeight(value int32) {
    SpinEdit_SetClientHeight(s.instance, value)
}

// ClientOrigin
func (s *TSpinEdit) ClientOrigin() TPoint {
    return SpinEdit_GetClientOrigin(s.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (s *TSpinEdit) ClientRect() TRect {
    return SpinEdit_GetClientRect(s.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (s *TSpinEdit) ClientWidth() int32 {
    return SpinEdit_GetClientWidth(s.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (s *TSpinEdit) SetClientWidth(value int32) {
    SpinEdit_SetClientWidth(s.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (s *TSpinEdit) ControlState() TControlState {
    return SpinEdit_GetControlState(s.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (s *TSpinEdit) SetControlState(value TControlState) {
    SpinEdit_SetControlState(s.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (s *TSpinEdit) ControlStyle() TControlStyle {
    return SpinEdit_GetControlStyle(s.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (s *TSpinEdit) SetControlStyle(value TControlStyle) {
    SpinEdit_SetControlStyle(s.instance, value)
}

// ExplicitLeft
func (s *TSpinEdit) ExplicitLeft() int32 {
    return SpinEdit_GetExplicitLeft(s.instance)
}

// ExplicitTop
func (s *TSpinEdit) ExplicitTop() int32 {
    return SpinEdit_GetExplicitTop(s.instance)
}

// ExplicitWidth
func (s *TSpinEdit) ExplicitWidth() int32 {
    return SpinEdit_GetExplicitWidth(s.instance)
}

// ExplicitHeight
func (s *TSpinEdit) ExplicitHeight() int32 {
    return SpinEdit_GetExplicitHeight(s.instance)
}

// Floating
func (s *TSpinEdit) Floating() bool {
    return SpinEdit_GetFloating(s.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TSpinEdit) Parent() *TWinControl {
    return WinControlFromInst(SpinEdit_GetParent(s.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TSpinEdit) SetParent(value IWinControl) {
    SpinEdit_SetParent(s.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (s *TSpinEdit) StyleElements() TStyleElements {
    return SpinEdit_GetStyleElements(s.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (s *TSpinEdit) SetStyleElements(value TStyleElements) {
    SpinEdit_SetStyleElements(s.instance, value)
}

// SetOnGesture
func (s *TSpinEdit) SetOnGesture(fn TGestureEvent) {
    SpinEdit_SetOnGesture(s.instance, fn)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (s *TSpinEdit) AlignWithMargins() bool {
    return SpinEdit_GetAlignWithMargins(s.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (s *TSpinEdit) SetAlignWithMargins(value bool) {
    SpinEdit_SetAlignWithMargins(s.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (s *TSpinEdit) Left() int32 {
    return SpinEdit_GetLeft(s.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (s *TSpinEdit) SetLeft(value int32) {
    SpinEdit_SetLeft(s.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (s *TSpinEdit) Top() int32 {
    return SpinEdit_GetTop(s.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (s *TSpinEdit) SetTop(value int32) {
    SpinEdit_SetTop(s.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (s *TSpinEdit) Width() int32 {
    return SpinEdit_GetWidth(s.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (s *TSpinEdit) SetWidth(value int32) {
    SpinEdit_SetWidth(s.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (s *TSpinEdit) Height() int32 {
    return SpinEdit_GetHeight(s.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (s *TSpinEdit) SetHeight(value int32) {
    SpinEdit_SetHeight(s.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TSpinEdit) Cursor() TCursor {
    return SpinEdit_GetCursor(s.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TSpinEdit) SetCursor(value TCursor) {
    SpinEdit_SetCursor(s.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (s *TSpinEdit) Hint() string {
    return SpinEdit_GetHint(s.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (s *TSpinEdit) SetHint(value string) {
    SpinEdit_SetHint(s.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (s *TSpinEdit) Margins() *TMargins {
    return MarginsFromInst(SpinEdit_GetMargins(s.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (s *TSpinEdit) SetMargins(value *TMargins) {
    SpinEdit_SetMargins(s.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (s *TSpinEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(SpinEdit_GetCustomHint(s.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (s *TSpinEdit) SetCustomHint(value IComponent) {
    SpinEdit_SetCustomHint(s.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TSpinEdit) ComponentCount() int32 {
    return SpinEdit_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TSpinEdit) ComponentIndex() int32 {
    return SpinEdit_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TSpinEdit) SetComponentIndex(value int32) {
    SpinEdit_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TSpinEdit) Owner() *TComponent {
    return ComponentFromInst(SpinEdit_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TSpinEdit) Name() string {
    return SpinEdit_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TSpinEdit) SetName(value string) {
    SpinEdit_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TSpinEdit) Tag() int {
    return SpinEdit_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TSpinEdit) SetTag(value int) {
    SpinEdit_SetTag(s.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (s *TSpinEdit) DockClients(Index int32) *TControl {
    return ControlFromInst(SpinEdit_GetDockClients(s.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (s *TSpinEdit) Controls(Index int32) *TControl {
    return ControlFromInst(SpinEdit_GetControls(s.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TSpinEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(SpinEdit_GetComponents(s.instance, AIndex))
}

