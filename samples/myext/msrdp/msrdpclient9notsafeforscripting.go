//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

package msrdp

import (
	"unsafe"

	. "github.com/ying32/govcl/vcl"
	. "github.com/ying32/govcl/vcl/types"
)

type TMsRdpClient9NotSafeForScripting struct {
	IWinControl
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
	ptr unsafe.Pointer
}

// NewMsRdpClient9NotSafeForScripting
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMsRdpClient9NotSafeForScripting(owner IComponent) *TMsRdpClient9NotSafeForScripting {
	m := new(TMsRdpClient9NotSafeForScripting)
	m.instance = MsRdpClient9NotSafeForScripting_Create(CheckPtr(owner))
	m.ptr = unsafe.Pointer(m.instance)
	return m
}

// MsRdpClient9NotSafeForScriptingFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MsRdpClient9NotSafeForScriptingFromInst(inst uintptr) *TMsRdpClient9NotSafeForScripting {
	m := new(TMsRdpClient9NotSafeForScripting)
	m.instance = inst
	m.ptr = unsafe.Pointer(inst)
	return m
}

// MsRdpClient9NotSafeForScriptingFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MsRdpClient9NotSafeForScriptingFromObj(obj IObject) *TMsRdpClient9NotSafeForScripting {
	m := new(TMsRdpClient9NotSafeForScripting)
	m.instance = CheckPtr(obj)
	m.ptr = unsafe.Pointer(m.instance)
	return m
}

// MsRdpClient9NotSafeForScriptingFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MsRdpClient9NotSafeForScriptingFromUnsafePointer(ptr unsafe.Pointer) *TMsRdpClient9NotSafeForScripting {
	m := new(TMsRdpClient9NotSafeForScripting)
	m.instance = uintptr(ptr)
	m.ptr = ptr
	return m
}

// Free
// CN: 释放对象。
// EN: Free object.
func (m *TMsRdpClient9NotSafeForScripting) Free() {
	if m.instance != 0 {
		MsRdpClient9NotSafeForScripting_Free(m.instance)
		m.instance = 0
		m.ptr = unsafe.Pointer(uintptr(0))
	}
}

// Instance
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMsRdpClient9NotSafeForScripting) Instance() uintptr {
	return m.instance
}

// UnsafeAddr
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMsRdpClient9NotSafeForScripting) UnsafeAddr() unsafe.Pointer {
	return m.ptr
}

// IsValid
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMsRdpClient9NotSafeForScripting) IsValid() bool {
	return m.instance != 0
}

// TMsRdpClient9NotSafeForScriptingClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMsRdpClient9NotSafeForScriptingClass() TClass {
	return MsRdpClient9NotSafeForScripting_StaticClassType()
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (m *TMsRdpClient9NotSafeForScripting) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	MsRdpClient9NotSafeForScripting_SetBounds(m.instance, ALeft, ATop, AWidth, AHeight)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (m *TMsRdpClient9NotSafeForScripting) CanFocus() bool {
	return MsRdpClient9NotSafeForScripting_CanFocus(m.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (m *TMsRdpClient9NotSafeForScripting) ContainsControl(Control IControl) bool {
	return MsRdpClient9NotSafeForScripting_ContainsControl(m.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (m *TMsRdpClient9NotSafeForScripting) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
	return ControlFromInst(MsRdpClient9NotSafeForScripting_ControlAtPos(m.instance, Pos, AllowDisabled, AllowWinControls, AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (m *TMsRdpClient9NotSafeForScripting) DisableAlign() {
	MsRdpClient9NotSafeForScripting_DisableAlign(m.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (m *TMsRdpClient9NotSafeForScripting) EnableAlign() {
	MsRdpClient9NotSafeForScripting_EnableAlign(m.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (m *TMsRdpClient9NotSafeForScripting) FindChildControl(ControlName string) *TControl {
	return ControlFromInst(MsRdpClient9NotSafeForScripting_FindChildControl(m.instance, ControlName))
}

// FlipChildren
func (m *TMsRdpClient9NotSafeForScripting) FlipChildren(AllLevels bool) {
	MsRdpClient9NotSafeForScripting_FlipChildren(m.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (m *TMsRdpClient9NotSafeForScripting) Focused() bool {
	return MsRdpClient9NotSafeForScripting_Focused(m.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (m *TMsRdpClient9NotSafeForScripting) HandleAllocated() bool {
	return MsRdpClient9NotSafeForScripting_HandleAllocated(m.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (m *TMsRdpClient9NotSafeForScripting) InsertControl(AControl IControl) {
	MsRdpClient9NotSafeForScripting_InsertControl(m.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (m *TMsRdpClient9NotSafeForScripting) Invalidate() {
	MsRdpClient9NotSafeForScripting_Invalidate(m.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (m *TMsRdpClient9NotSafeForScripting) PaintTo(DC HDC, X int32, Y int32) {
	MsRdpClient9NotSafeForScripting_PaintTo(m.instance, DC, X, Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (m *TMsRdpClient9NotSafeForScripting) RemoveControl(AControl IControl) {
	MsRdpClient9NotSafeForScripting_RemoveControl(m.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (m *TMsRdpClient9NotSafeForScripting) Realign() {
	MsRdpClient9NotSafeForScripting_Realign(m.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (m *TMsRdpClient9NotSafeForScripting) Repaint() {
	MsRdpClient9NotSafeForScripting_Repaint(m.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (m *TMsRdpClient9NotSafeForScripting) ScaleBy(M int32, D int32) {
	MsRdpClient9NotSafeForScripting_ScaleBy(m.instance, M, D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (m *TMsRdpClient9NotSafeForScripting) ScrollBy(DeltaX int32, DeltaY int32) {
	MsRdpClient9NotSafeForScripting_ScrollBy(m.instance, DeltaX, DeltaY)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (m *TMsRdpClient9NotSafeForScripting) SetFocus() {
	MsRdpClient9NotSafeForScripting_SetFocus(m.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (m *TMsRdpClient9NotSafeForScripting) Update() {
	MsRdpClient9NotSafeForScripting_Update(m.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (m *TMsRdpClient9NotSafeForScripting) UpdateControlState() {
	MsRdpClient9NotSafeForScripting_UpdateControlState(m.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (m *TMsRdpClient9NotSafeForScripting) BringToFront() {
	MsRdpClient9NotSafeForScripting_BringToFront(m.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (m *TMsRdpClient9NotSafeForScripting) ClientToScreen(Point TPoint) TPoint {
	return MsRdpClient9NotSafeForScripting_ClientToScreen(m.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (m *TMsRdpClient9NotSafeForScripting) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
	return MsRdpClient9NotSafeForScripting_ClientToParent(m.instance, Point, CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (m *TMsRdpClient9NotSafeForScripting) Dragging() bool {
	return MsRdpClient9NotSafeForScripting_Dragging(m.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (m *TMsRdpClient9NotSafeForScripting) HasParent() bool {
	return MsRdpClient9NotSafeForScripting_HasParent(m.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (m *TMsRdpClient9NotSafeForScripting) Hide() {
	MsRdpClient9NotSafeForScripting_Hide(m.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (m *TMsRdpClient9NotSafeForScripting) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return MsRdpClient9NotSafeForScripting_Perform(m.instance, Msg, WParam, LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (m *TMsRdpClient9NotSafeForScripting) Refresh() {
	MsRdpClient9NotSafeForScripting_Refresh(m.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (m *TMsRdpClient9NotSafeForScripting) ScreenToClient(Point TPoint) TPoint {
	return MsRdpClient9NotSafeForScripting_ScreenToClient(m.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (m *TMsRdpClient9NotSafeForScripting) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
	return MsRdpClient9NotSafeForScripting_ParentToClient(m.instance, Point, CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (m *TMsRdpClient9NotSafeForScripting) SendToBack() {
	MsRdpClient9NotSafeForScripting_SendToBack(m.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (m *TMsRdpClient9NotSafeForScripting) Show() {
	MsRdpClient9NotSafeForScripting_Show(m.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (m *TMsRdpClient9NotSafeForScripting) GetTextBuf(Buffer string, BufSize int32) int32 {
	return MsRdpClient9NotSafeForScripting_GetTextBuf(m.instance, Buffer, BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (m *TMsRdpClient9NotSafeForScripting) GetTextLen() int32 {
	return MsRdpClient9NotSafeForScripting_GetTextLen(m.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (m *TMsRdpClient9NotSafeForScripting) SetTextBuf(Buffer string) {
	MsRdpClient9NotSafeForScripting_SetTextBuf(m.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (m *TMsRdpClient9NotSafeForScripting) FindComponent(AName string) *TComponent {
	return ComponentFromInst(MsRdpClient9NotSafeForScripting_FindComponent(m.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (m *TMsRdpClient9NotSafeForScripting) GetNamePath() string {
	return MsRdpClient9NotSafeForScripting_GetNamePath(m.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (m *TMsRdpClient9NotSafeForScripting) Assign(Source IObject) {
	MsRdpClient9NotSafeForScripting_Assign(m.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMsRdpClient9NotSafeForScripting) DisposeOf() {
	MsRdpClient9NotSafeForScripting_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMsRdpClient9NotSafeForScripting) ClassType() TClass {
	return MsRdpClient9NotSafeForScripting_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMsRdpClient9NotSafeForScripting) ClassName() string {
	return MsRdpClient9NotSafeForScripting_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMsRdpClient9NotSafeForScripting) InstanceSize() int32 {
	return MsRdpClient9NotSafeForScripting_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMsRdpClient9NotSafeForScripting) InheritsFrom(AClass TClass) bool {
	return MsRdpClient9NotSafeForScripting_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMsRdpClient9NotSafeForScripting) Equals(Obj IObject) bool {
	return MsRdpClient9NotSafeForScripting_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMsRdpClient9NotSafeForScripting) GetHashCode() int32 {
	return MsRdpClient9NotSafeForScripting_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMsRdpClient9NotSafeForScripting) ToString() string {
	return MsRdpClient9NotSafeForScripting_ToString(m.instance)
}

// ControlInterface
func (m *TMsRdpClient9NotSafeForScripting) ControlInterface() *IMsRdpClient9 {
	return IMsRdpClient9FromInst(MsRdpClient9NotSafeForScripting_GetControlInterface(m.instance))
}

// DefaultInterface
func (m *TMsRdpClient9NotSafeForScripting) DefaultInterface() *IMsRdpClient9 {
	return IMsRdpClient9FromInst(MsRdpClient9NotSafeForScripting_GetDefaultInterface(m.instance))
}

// Version
func (m *TMsRdpClient9NotSafeForScripting) Version() string {
	return MsRdpClient9NotSafeForScripting_GetVersion(m.instance)
}

// AdvancedSettings9
func (m *TMsRdpClient9NotSafeForScripting) AdvancedSettings9() *IMsRdpClientAdvancedSettings8 {
	return IMsRdpClientAdvancedSettings8FromInst(MsRdpClient9NotSafeForScripting_GetAdvancedSettings9(m.instance))
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (m *TMsRdpClient9NotSafeForScripting) Anchors() TAnchors {
	return MsRdpClient9NotSafeForScripting_GetAnchors(m.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (m *TMsRdpClient9NotSafeForScripting) SetAnchors(value TAnchors) {
	MsRdpClient9NotSafeForScripting_SetAnchors(m.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (m *TMsRdpClient9NotSafeForScripting) TabStop() bool {
	return MsRdpClient9NotSafeForScripting_GetTabStop(m.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (m *TMsRdpClient9NotSafeForScripting) SetTabStop(value bool) {
	MsRdpClient9NotSafeForScripting_SetTabStop(m.instance, value)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (m *TMsRdpClient9NotSafeForScripting) Align() TAlign {
	return MsRdpClient9NotSafeForScripting_GetAlign(m.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (m *TMsRdpClient9NotSafeForScripting) SetAlign(value TAlign) {
	MsRdpClient9NotSafeForScripting_SetAlign(m.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (m *TMsRdpClient9NotSafeForScripting) DragCursor() TCursor {
	return MsRdpClient9NotSafeForScripting_GetDragCursor(m.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (m *TMsRdpClient9NotSafeForScripting) SetDragCursor(value TCursor) {
	MsRdpClient9NotSafeForScripting_SetDragCursor(m.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (m *TMsRdpClient9NotSafeForScripting) DragMode() TDragMode {
	return MsRdpClient9NotSafeForScripting_GetDragMode(m.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (m *TMsRdpClient9NotSafeForScripting) SetDragMode(value TDragMode) {
	MsRdpClient9NotSafeForScripting_SetDragMode(m.instance, value)
}

// ParentShowHint
func (m *TMsRdpClient9NotSafeForScripting) ParentShowHint() bool {
	return MsRdpClient9NotSafeForScripting_GetParentShowHint(m.instance)
}

// SetParentShowHint
func (m *TMsRdpClient9NotSafeForScripting) SetParentShowHint(value bool) {
	MsRdpClient9NotSafeForScripting_SetParentShowHint(m.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (m *TMsRdpClient9NotSafeForScripting) PopupMenu() *TPopupMenu {
	return PopupMenuFromInst(MsRdpClient9NotSafeForScripting_GetPopupMenu(m.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (m *TMsRdpClient9NotSafeForScripting) SetPopupMenu(value IComponent) {
	MsRdpClient9NotSafeForScripting_SetPopupMenu(m.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (m *TMsRdpClient9NotSafeForScripting) ShowHint() bool {
	return MsRdpClient9NotSafeForScripting_GetShowHint(m.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (m *TMsRdpClient9NotSafeForScripting) SetShowHint(value bool) {
	MsRdpClient9NotSafeForScripting_SetShowHint(m.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (m *TMsRdpClient9NotSafeForScripting) TabOrder() TTabOrder {
	return MsRdpClient9NotSafeForScripting_GetTabOrder(m.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (m *TMsRdpClient9NotSafeForScripting) SetTabOrder(value TTabOrder) {
	MsRdpClient9NotSafeForScripting_SetTabOrder(m.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (m *TMsRdpClient9NotSafeForScripting) Visible() bool {
	return MsRdpClient9NotSafeForScripting_GetVisible(m.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (m *TMsRdpClient9NotSafeForScripting) SetVisible(value bool) {
	MsRdpClient9NotSafeForScripting_SetVisible(m.instance, value)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (m *TMsRdpClient9NotSafeForScripting) SetOnDragDrop(fn TDragDropEvent) {
	MsRdpClient9NotSafeForScripting_SetOnDragDrop(m.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (m *TMsRdpClient9NotSafeForScripting) SetOnDragOver(fn TDragOverEvent) {
	MsRdpClient9NotSafeForScripting_SetOnDragOver(m.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (m *TMsRdpClient9NotSafeForScripting) SetOnEndDrag(fn TEndDragEvent) {
	MsRdpClient9NotSafeForScripting_SetOnEndDrag(m.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (m *TMsRdpClient9NotSafeForScripting) SetOnEnter(fn TNotifyEvent) {
	MsRdpClient9NotSafeForScripting_SetOnEnter(m.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (m *TMsRdpClient9NotSafeForScripting) SetOnExit(fn TNotifyEvent) {
	MsRdpClient9NotSafeForScripting_SetOnExit(m.instance, fn)
}

// Server
func (m *TMsRdpClient9NotSafeForScripting) Server() string {
	return MsRdpClient9NotSafeForScripting_GetServer(m.instance)
}

// SetServer
func (m *TMsRdpClient9NotSafeForScripting) SetServer(value string) {
	MsRdpClient9NotSafeForScripting_SetServer(m.instance, value)
}

// UserName
func (m *TMsRdpClient9NotSafeForScripting) UserName() string {
	return MsRdpClient9NotSafeForScripting_GetUserName(m.instance)
}

// SetUserName
func (m *TMsRdpClient9NotSafeForScripting) SetUserName(value string) {
	MsRdpClient9NotSafeForScripting_SetUserName(m.instance, value)
}

// DisconnectedText
func (m *TMsRdpClient9NotSafeForScripting) DisconnectedText() string {
	return MsRdpClient9NotSafeForScripting_GetDisconnectedText(m.instance)
}

// SetDisconnectedText
func (m *TMsRdpClient9NotSafeForScripting) SetDisconnectedText(value string) {
	MsRdpClient9NotSafeForScripting_SetDisconnectedText(m.instance, value)
}

// ConnectingText
func (m *TMsRdpClient9NotSafeForScripting) ConnectingText() string {
	return MsRdpClient9NotSafeForScripting_GetConnectingText(m.instance)
}

// SetConnectingText
func (m *TMsRdpClient9NotSafeForScripting) SetConnectingText(value string) {
	MsRdpClient9NotSafeForScripting_SetConnectingText(m.instance, value)
}

// DesktopWidth
func (m *TMsRdpClient9NotSafeForScripting) DesktopWidth() int32 {
	return MsRdpClient9NotSafeForScripting_GetDesktopWidth(m.instance)
}

// SetDesktopWidth
func (m *TMsRdpClient9NotSafeForScripting) SetDesktopWidth(value int32) {
	MsRdpClient9NotSafeForScripting_SetDesktopWidth(m.instance, value)
}

// DesktopHeight
func (m *TMsRdpClient9NotSafeForScripting) DesktopHeight() int32 {
	return MsRdpClient9NotSafeForScripting_GetDesktopHeight(m.instance)
}

// SetDesktopHeight
func (m *TMsRdpClient9NotSafeForScripting) SetDesktopHeight(value int32) {
	MsRdpClient9NotSafeForScripting_SetDesktopHeight(m.instance, value)
}

// StartConnected
func (m *TMsRdpClient9NotSafeForScripting) StartConnected() int32 {
	return MsRdpClient9NotSafeForScripting_GetStartConnected(m.instance)
}

// SetStartConnected
func (m *TMsRdpClient9NotSafeForScripting) SetStartConnected(value int32) {
	MsRdpClient9NotSafeForScripting_SetStartConnected(m.instance, value)
}

// ColorDepth
func (m *TMsRdpClient9NotSafeForScripting) ColorDepth() int32 {
	return MsRdpClient9NotSafeForScripting_GetColorDepth(m.instance)
}

// SetColorDepth
func (m *TMsRdpClient9NotSafeForScripting) SetColorDepth(value int32) {
	MsRdpClient9NotSafeForScripting_SetColorDepth(m.instance, value)
}

// FullScreen
func (m *TMsRdpClient9NotSafeForScripting) FullScreen() bool {
	return MsRdpClient9NotSafeForScripting_GetFullScreen(m.instance)
}

// SetFullScreen
func (m *TMsRdpClient9NotSafeForScripting) SetFullScreen(value bool) {
	MsRdpClient9NotSafeForScripting_SetFullScreen(m.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (m *TMsRdpClient9NotSafeForScripting) DockClientCount() int32 {
	return MsRdpClient9NotSafeForScripting_GetDockClientCount(m.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (m *TMsRdpClient9NotSafeForScripting) DockSite() bool {
	return MsRdpClient9NotSafeForScripting_GetDockSite(m.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (m *TMsRdpClient9NotSafeForScripting) SetDockSite(value bool) {
	MsRdpClient9NotSafeForScripting_SetDockSite(m.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (m *TMsRdpClient9NotSafeForScripting) DoubleBuffered() bool {
	return MsRdpClient9NotSafeForScripting_GetDoubleBuffered(m.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (m *TMsRdpClient9NotSafeForScripting) SetDoubleBuffered(value bool) {
	MsRdpClient9NotSafeForScripting_SetDoubleBuffered(m.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (m *TMsRdpClient9NotSafeForScripting) AlignDisabled() bool {
	return MsRdpClient9NotSafeForScripting_GetAlignDisabled(m.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (m *TMsRdpClient9NotSafeForScripting) MouseInClient() bool {
	return MsRdpClient9NotSafeForScripting_GetMouseInClient(m.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (m *TMsRdpClient9NotSafeForScripting) VisibleDockClientCount() int32 {
	return MsRdpClient9NotSafeForScripting_GetVisibleDockClientCount(m.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (m *TMsRdpClient9NotSafeForScripting) Brush() *TBrush {
	return BrushFromInst(MsRdpClient9NotSafeForScripting_GetBrush(m.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (m *TMsRdpClient9NotSafeForScripting) ControlCount() int32 {
	return MsRdpClient9NotSafeForScripting_GetControlCount(m.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (m *TMsRdpClient9NotSafeForScripting) Handle() HWND {
	return MsRdpClient9NotSafeForScripting_GetHandle(m.instance)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (m *TMsRdpClient9NotSafeForScripting) ParentDoubleBuffered() bool {
	return MsRdpClient9NotSafeForScripting_GetParentDoubleBuffered(m.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (m *TMsRdpClient9NotSafeForScripting) SetParentDoubleBuffered(value bool) {
	MsRdpClient9NotSafeForScripting_SetParentDoubleBuffered(m.instance, value)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (m *TMsRdpClient9NotSafeForScripting) ParentWindow() HWND {
	return MsRdpClient9NotSafeForScripting_GetParentWindow(m.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (m *TMsRdpClient9NotSafeForScripting) SetParentWindow(value HWND) {
	MsRdpClient9NotSafeForScripting_SetParentWindow(m.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (m *TMsRdpClient9NotSafeForScripting) UseDockManager() bool {
	return MsRdpClient9NotSafeForScripting_GetUseDockManager(m.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (m *TMsRdpClient9NotSafeForScripting) SetUseDockManager(value bool) {
	MsRdpClient9NotSafeForScripting_SetUseDockManager(m.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (m *TMsRdpClient9NotSafeForScripting) Enabled() bool {
	return MsRdpClient9NotSafeForScripting_GetEnabled(m.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (m *TMsRdpClient9NotSafeForScripting) SetEnabled(value bool) {
	MsRdpClient9NotSafeForScripting_SetEnabled(m.instance, value)
}

// Action
func (m *TMsRdpClient9NotSafeForScripting) Action() *TAction {
	return ActionFromInst(MsRdpClient9NotSafeForScripting_GetAction(m.instance))
}

// SetAction
func (m *TMsRdpClient9NotSafeForScripting) SetAction(value IComponent) {
	MsRdpClient9NotSafeForScripting_SetAction(m.instance, CheckPtr(value))
}

// BiDiMode
func (m *TMsRdpClient9NotSafeForScripting) BiDiMode() TBiDiMode {
	return MsRdpClient9NotSafeForScripting_GetBiDiMode(m.instance)
}

// SetBiDiMode
func (m *TMsRdpClient9NotSafeForScripting) SetBiDiMode(value TBiDiMode) {
	MsRdpClient9NotSafeForScripting_SetBiDiMode(m.instance, value)
}

// BoundsRect
func (m *TMsRdpClient9NotSafeForScripting) BoundsRect() TRect {
	return MsRdpClient9NotSafeForScripting_GetBoundsRect(m.instance)
}

// SetBoundsRect
func (m *TMsRdpClient9NotSafeForScripting) SetBoundsRect(value TRect) {
	MsRdpClient9NotSafeForScripting_SetBoundsRect(m.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (m *TMsRdpClient9NotSafeForScripting) ClientHeight() int32 {
	return MsRdpClient9NotSafeForScripting_GetClientHeight(m.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (m *TMsRdpClient9NotSafeForScripting) SetClientHeight(value int32) {
	MsRdpClient9NotSafeForScripting_SetClientHeight(m.instance, value)
}

// ClientOrigin
func (m *TMsRdpClient9NotSafeForScripting) ClientOrigin() TPoint {
	return MsRdpClient9NotSafeForScripting_GetClientOrigin(m.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (m *TMsRdpClient9NotSafeForScripting) ClientRect() TRect {
	return MsRdpClient9NotSafeForScripting_GetClientRect(m.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (m *TMsRdpClient9NotSafeForScripting) ClientWidth() int32 {
	return MsRdpClient9NotSafeForScripting_GetClientWidth(m.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (m *TMsRdpClient9NotSafeForScripting) SetClientWidth(value int32) {
	MsRdpClient9NotSafeForScripting_SetClientWidth(m.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (m *TMsRdpClient9NotSafeForScripting) ControlState() TControlState {
	return MsRdpClient9NotSafeForScripting_GetControlState(m.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (m *TMsRdpClient9NotSafeForScripting) SetControlState(value TControlState) {
	MsRdpClient9NotSafeForScripting_SetControlState(m.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (m *TMsRdpClient9NotSafeForScripting) ControlStyle() TControlStyle {
	return MsRdpClient9NotSafeForScripting_GetControlStyle(m.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (m *TMsRdpClient9NotSafeForScripting) SetControlStyle(value TControlStyle) {
	MsRdpClient9NotSafeForScripting_SetControlStyle(m.instance, value)
}

// ExplicitLeft
func (m *TMsRdpClient9NotSafeForScripting) ExplicitLeft() int32 {
	return MsRdpClient9NotSafeForScripting_GetExplicitLeft(m.instance)
}

// ExplicitTop
func (m *TMsRdpClient9NotSafeForScripting) ExplicitTop() int32 {
	return MsRdpClient9NotSafeForScripting_GetExplicitTop(m.instance)
}

// ExplicitWidth
func (m *TMsRdpClient9NotSafeForScripting) ExplicitWidth() int32 {
	return MsRdpClient9NotSafeForScripting_GetExplicitWidth(m.instance)
}

// ExplicitHeight
func (m *TMsRdpClient9NotSafeForScripting) ExplicitHeight() int32 {
	return MsRdpClient9NotSafeForScripting_GetExplicitHeight(m.instance)
}

// Floating
func (m *TMsRdpClient9NotSafeForScripting) Floating() bool {
	return MsRdpClient9NotSafeForScripting_GetFloating(m.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (m *TMsRdpClient9NotSafeForScripting) Parent() *TWinControl {
	return WinControlFromInst(MsRdpClient9NotSafeForScripting_GetParent(m.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (m *TMsRdpClient9NotSafeForScripting) SetParent(value IWinControl) {
	MsRdpClient9NotSafeForScripting_SetParent(m.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (m *TMsRdpClient9NotSafeForScripting) StyleElements() TStyleElements {
	return MsRdpClient9NotSafeForScripting_GetStyleElements(m.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (m *TMsRdpClient9NotSafeForScripting) SetStyleElements(value TStyleElements) {
	MsRdpClient9NotSafeForScripting_SetStyleElements(m.instance, value)
}

// SetOnGesture
func (m *TMsRdpClient9NotSafeForScripting) SetOnGesture(fn TGestureEvent) {
	MsRdpClient9NotSafeForScripting_SetOnGesture(m.instance, fn)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (m *TMsRdpClient9NotSafeForScripting) AlignWithMargins() bool {
	return MsRdpClient9NotSafeForScripting_GetAlignWithMargins(m.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (m *TMsRdpClient9NotSafeForScripting) SetAlignWithMargins(value bool) {
	MsRdpClient9NotSafeForScripting_SetAlignWithMargins(m.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (m *TMsRdpClient9NotSafeForScripting) Left() int32 {
	return MsRdpClient9NotSafeForScripting_GetLeft(m.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (m *TMsRdpClient9NotSafeForScripting) SetLeft(value int32) {
	MsRdpClient9NotSafeForScripting_SetLeft(m.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (m *TMsRdpClient9NotSafeForScripting) Top() int32 {
	return MsRdpClient9NotSafeForScripting_GetTop(m.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (m *TMsRdpClient9NotSafeForScripting) SetTop(value int32) {
	MsRdpClient9NotSafeForScripting_SetTop(m.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (m *TMsRdpClient9NotSafeForScripting) Width() int32 {
	return MsRdpClient9NotSafeForScripting_GetWidth(m.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (m *TMsRdpClient9NotSafeForScripting) SetWidth(value int32) {
	MsRdpClient9NotSafeForScripting_SetWidth(m.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (m *TMsRdpClient9NotSafeForScripting) Height() int32 {
	return MsRdpClient9NotSafeForScripting_GetHeight(m.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (m *TMsRdpClient9NotSafeForScripting) SetHeight(value int32) {
	MsRdpClient9NotSafeForScripting_SetHeight(m.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (m *TMsRdpClient9NotSafeForScripting) Cursor() TCursor {
	return MsRdpClient9NotSafeForScripting_GetCursor(m.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (m *TMsRdpClient9NotSafeForScripting) SetCursor(value TCursor) {
	MsRdpClient9NotSafeForScripting_SetCursor(m.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (m *TMsRdpClient9NotSafeForScripting) Hint() string {
	return MsRdpClient9NotSafeForScripting_GetHint(m.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (m *TMsRdpClient9NotSafeForScripting) SetHint(value string) {
	MsRdpClient9NotSafeForScripting_SetHint(m.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (m *TMsRdpClient9NotSafeForScripting) Margins() *TMargins {
	return MarginsFromInst(MsRdpClient9NotSafeForScripting_GetMargins(m.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (m *TMsRdpClient9NotSafeForScripting) SetMargins(value *TMargins) {
	MsRdpClient9NotSafeForScripting_SetMargins(m.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (m *TMsRdpClient9NotSafeForScripting) CustomHint() *TCustomHint {
	return CustomHintFromInst(MsRdpClient9NotSafeForScripting_GetCustomHint(m.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (m *TMsRdpClient9NotSafeForScripting) SetCustomHint(value IComponent) {
	MsRdpClient9NotSafeForScripting_SetCustomHint(m.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (m *TMsRdpClient9NotSafeForScripting) ComponentCount() int32 {
	return MsRdpClient9NotSafeForScripting_GetComponentCount(m.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (m *TMsRdpClient9NotSafeForScripting) ComponentIndex() int32 {
	return MsRdpClient9NotSafeForScripting_GetComponentIndex(m.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (m *TMsRdpClient9NotSafeForScripting) SetComponentIndex(value int32) {
	MsRdpClient9NotSafeForScripting_SetComponentIndex(m.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (m *TMsRdpClient9NotSafeForScripting) Owner() *TComponent {
	return ComponentFromInst(MsRdpClient9NotSafeForScripting_GetOwner(m.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (m *TMsRdpClient9NotSafeForScripting) Name() string {
	return MsRdpClient9NotSafeForScripting_GetName(m.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (m *TMsRdpClient9NotSafeForScripting) SetName(value string) {
	MsRdpClient9NotSafeForScripting_SetName(m.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (m *TMsRdpClient9NotSafeForScripting) Tag() int {
	return MsRdpClient9NotSafeForScripting_GetTag(m.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (m *TMsRdpClient9NotSafeForScripting) SetTag(value int) {
	MsRdpClient9NotSafeForScripting_SetTag(m.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (m *TMsRdpClient9NotSafeForScripting) DockClients(Index int32) *TControl {
	return ControlFromInst(MsRdpClient9NotSafeForScripting_GetDockClients(m.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (m *TMsRdpClient9NotSafeForScripting) Controls(Index int32) *TControl {
	return ControlFromInst(MsRdpClient9NotSafeForScripting_GetControls(m.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (m *TMsRdpClient9NotSafeForScripting) Components(AIndex int32) *TComponent {
	return ComponentFromInst(MsRdpClient9NotSafeForScripting_GetComponents(m.instance, AIndex))
}
