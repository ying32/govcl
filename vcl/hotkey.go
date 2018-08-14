
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

type THotKey struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewHotKey
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewHotKey(owner IComponent) *THotKey {
    h := new(THotKey)
    h.instance = HotKey_Create(CheckPtr(owner))
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HotKeyFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func HotKeyFromInst(inst uintptr) *THotKey {
    h := new(THotKey)
    h.instance = inst
    h.ptr = unsafe.Pointer(inst)
    return h
}

// HotKeyFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func HotKeyFromObj(obj IObject) *THotKey {
    h := new(THotKey)
    h.instance = CheckPtr(obj)
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HotKeyFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func HotKeyFromUnsafePointer(ptr unsafe.Pointer) *THotKey {
    h := new(THotKey)
    h.instance = uintptr(ptr)
    h.ptr = ptr
    return h
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (h *THotKey) Free() {
    if h.instance != 0 {
        HotKey_Free(h.instance)
        h.instance = 0
        h.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (h *THotKey) Instance() uintptr {
    return h.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (h *THotKey) UnsafeAddr() unsafe.Pointer {
    return h.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (h *THotKey) IsValid() bool {
    return h.instance != 0
}

// THotKeyClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func THotKeyClass() TClass {
    return HotKey_StaticClassType()
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (h *THotKey) CanFocus() bool {
    return HotKey_CanFocus(h.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (h *THotKey) ContainsControl(Control IControl) bool {
    return HotKey_ContainsControl(h.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (h *THotKey) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(HotKey_ControlAtPos(h.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (h *THotKey) DisableAlign() {
    HotKey_DisableAlign(h.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (h *THotKey) EnableAlign() {
    HotKey_EnableAlign(h.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (h *THotKey) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(HotKey_FindChildControl(h.instance, ControlName))
}

// FlipChildren
func (h *THotKey) FlipChildren(AllLevels bool) {
    HotKey_FlipChildren(h.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (h *THotKey) Focused() bool {
    return HotKey_Focused(h.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (h *THotKey) HandleAllocated() bool {
    return HotKey_HandleAllocated(h.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (h *THotKey) InsertControl(AControl IControl) {
    HotKey_InsertControl(h.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (h *THotKey) Invalidate() {
    HotKey_Invalidate(h.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (h *THotKey) PaintTo(DC HDC, X int32, Y int32) {
    HotKey_PaintTo(h.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (h *THotKey) RemoveControl(AControl IControl) {
    HotKey_RemoveControl(h.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (h *THotKey) Realign() {
    HotKey_Realign(h.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (h *THotKey) Repaint() {
    HotKey_Repaint(h.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (h *THotKey) ScaleBy(M int32, D int32) {
    HotKey_ScaleBy(h.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (h *THotKey) ScrollBy(DeltaX int32, DeltaY int32) {
    HotKey_ScrollBy(h.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (h *THotKey) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    HotKey_SetBounds(h.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (h *THotKey) SetFocus() {
    HotKey_SetFocus(h.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (h *THotKey) Update() {
    HotKey_Update(h.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (h *THotKey) UpdateControlState() {
    HotKey_UpdateControlState(h.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (h *THotKey) BringToFront() {
    HotKey_BringToFront(h.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (h *THotKey) ClientToScreen(Point TPoint) TPoint {
    return HotKey_ClientToScreen(h.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (h *THotKey) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return HotKey_ClientToParent(h.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (h *THotKey) Dragging() bool {
    return HotKey_Dragging(h.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (h *THotKey) HasParent() bool {
    return HotKey_HasParent(h.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (h *THotKey) Hide() {
    HotKey_Hide(h.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (h *THotKey) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return HotKey_Perform(h.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (h *THotKey) Refresh() {
    HotKey_Refresh(h.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (h *THotKey) ScreenToClient(Point TPoint) TPoint {
    return HotKey_ScreenToClient(h.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (h *THotKey) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return HotKey_ParentToClient(h.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (h *THotKey) SendToBack() {
    HotKey_SendToBack(h.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (h *THotKey) Show() {
    HotKey_Show(h.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (h *THotKey) GetTextBuf(Buffer string, BufSize int32) int32 {
    return HotKey_GetTextBuf(h.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (h *THotKey) GetTextLen() int32 {
    return HotKey_GetTextLen(h.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (h *THotKey) SetTextBuf(Buffer string) {
    HotKey_SetTextBuf(h.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (h *THotKey) FindComponent(AName string) *TComponent {
    return ComponentFromInst(HotKey_FindComponent(h.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (h *THotKey) GetNamePath() string {
    return HotKey_GetNamePath(h.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (h *THotKey) Assign(Source IObject) {
    HotKey_Assign(h.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (h *THotKey) DisposeOf() {
    HotKey_DisposeOf(h.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (h *THotKey) ClassType() TClass {
    return HotKey_ClassType(h.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (h *THotKey) ClassName() string {
    return HotKey_ClassName(h.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (h *THotKey) InstanceSize() int32 {
    return HotKey_InstanceSize(h.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (h *THotKey) InheritsFrom(AClass TClass) bool {
    return HotKey_InheritsFrom(h.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (h *THotKey) Equals(Obj IObject) bool {
    return HotKey_Equals(h.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (h *THotKey) GetHashCode() int32 {
    return HotKey_GetHashCode(h.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (h *THotKey) ToString() string {
    return HotKey_ToString(h.instance)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (h *THotKey) Anchors() TAnchors {
    return HotKey_GetAnchors(h.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (h *THotKey) SetAnchors(value TAnchors) {
    HotKey_SetAnchors(h.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (h *THotKey) AutoSize() bool {
    return HotKey_GetAutoSize(h.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (h *THotKey) SetAutoSize(value bool) {
    HotKey_SetAutoSize(h.instance, value)
}

// BiDiMode
func (h *THotKey) BiDiMode() TBiDiMode {
    return HotKey_GetBiDiMode(h.instance)
}

// SetBiDiMode
func (h *THotKey) SetBiDiMode(value TBiDiMode) {
    HotKey_SetBiDiMode(h.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (h *THotKey) Enabled() bool {
    return HotKey_GetEnabled(h.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (h *THotKey) SetEnabled(value bool) {
    HotKey_SetEnabled(h.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (h *THotKey) Hint() string {
    return HotKey_GetHint(h.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (h *THotKey) SetHint(value string) {
    HotKey_SetHint(h.instance, value)
}

// ParentShowHint
func (h *THotKey) ParentShowHint() bool {
    return HotKey_GetParentShowHint(h.instance)
}

// SetParentShowHint
func (h *THotKey) SetParentShowHint(value bool) {
    HotKey_SetParentShowHint(h.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (h *THotKey) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(HotKey_GetPopupMenu(h.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (h *THotKey) SetPopupMenu(value IComponent) {
    HotKey_SetPopupMenu(h.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (h *THotKey) ShowHint() bool {
    return HotKey_GetShowHint(h.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (h *THotKey) SetShowHint(value bool) {
    HotKey_SetShowHint(h.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (h *THotKey) TabOrder() TTabOrder {
    return HotKey_GetTabOrder(h.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (h *THotKey) SetTabOrder(value TTabOrder) {
    HotKey_SetTabOrder(h.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (h *THotKey) TabStop() bool {
    return HotKey_GetTabStop(h.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (h *THotKey) SetTabStop(value bool) {
    HotKey_SetTabStop(h.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (h *THotKey) Visible() bool {
    return HotKey_GetVisible(h.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (h *THotKey) SetVisible(value bool) {
    HotKey_SetVisible(h.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (h *THotKey) StyleElements() TStyleElements {
    return HotKey_GetStyleElements(h.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (h *THotKey) SetStyleElements(value TStyleElements) {
    HotKey_SetStyleElements(h.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (h *THotKey) SetOnChange(fn TNotifyEvent) {
    HotKey_SetOnChange(h.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (h *THotKey) SetOnContextPopup(fn TContextPopupEvent) {
    HotKey_SetOnContextPopup(h.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (h *THotKey) SetOnEnter(fn TNotifyEvent) {
    HotKey_SetOnEnter(h.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (h *THotKey) SetOnExit(fn TNotifyEvent) {
    HotKey_SetOnExit(h.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (h *THotKey) SetOnMouseDown(fn TMouseEvent) {
    HotKey_SetOnMouseDown(h.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (h *THotKey) SetOnMouseEnter(fn TNotifyEvent) {
    HotKey_SetOnMouseEnter(h.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (h *THotKey) SetOnMouseLeave(fn TNotifyEvent) {
    HotKey_SetOnMouseLeave(h.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (h *THotKey) SetOnMouseMove(fn TMouseMoveEvent) {
    HotKey_SetOnMouseMove(h.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (h *THotKey) SetOnMouseUp(fn TMouseEvent) {
    HotKey_SetOnMouseUp(h.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (h *THotKey) DockClientCount() int32 {
    return HotKey_GetDockClientCount(h.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (h *THotKey) DockSite() bool {
    return HotKey_GetDockSite(h.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (h *THotKey) SetDockSite(value bool) {
    HotKey_SetDockSite(h.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (h *THotKey) DoubleBuffered() bool {
    return HotKey_GetDoubleBuffered(h.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (h *THotKey) SetDoubleBuffered(value bool) {
    HotKey_SetDoubleBuffered(h.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (h *THotKey) AlignDisabled() bool {
    return HotKey_GetAlignDisabled(h.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (h *THotKey) MouseInClient() bool {
    return HotKey_GetMouseInClient(h.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (h *THotKey) VisibleDockClientCount() int32 {
    return HotKey_GetVisibleDockClientCount(h.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (h *THotKey) Brush() *TBrush {
    return BrushFromInst(HotKey_GetBrush(h.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (h *THotKey) ControlCount() int32 {
    return HotKey_GetControlCount(h.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (h *THotKey) Handle() HWND {
    return HotKey_GetHandle(h.instance)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (h *THotKey) ParentDoubleBuffered() bool {
    return HotKey_GetParentDoubleBuffered(h.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (h *THotKey) SetParentDoubleBuffered(value bool) {
    HotKey_SetParentDoubleBuffered(h.instance, value)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (h *THotKey) ParentWindow() HWND {
    return HotKey_GetParentWindow(h.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (h *THotKey) SetParentWindow(value HWND) {
    HotKey_SetParentWindow(h.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (h *THotKey) UseDockManager() bool {
    return HotKey_GetUseDockManager(h.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (h *THotKey) SetUseDockManager(value bool) {
    HotKey_SetUseDockManager(h.instance, value)
}

// Action
func (h *THotKey) Action() *TAction {
    return ActionFromInst(HotKey_GetAction(h.instance))
}

// SetAction
func (h *THotKey) SetAction(value IComponent) {
    HotKey_SetAction(h.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (h *THotKey) Align() TAlign {
    return HotKey_GetAlign(h.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (h *THotKey) SetAlign(value TAlign) {
    HotKey_SetAlign(h.instance, value)
}

// BoundsRect
func (h *THotKey) BoundsRect() TRect {
    return HotKey_GetBoundsRect(h.instance)
}

// SetBoundsRect
func (h *THotKey) SetBoundsRect(value TRect) {
    HotKey_SetBoundsRect(h.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (h *THotKey) ClientHeight() int32 {
    return HotKey_GetClientHeight(h.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (h *THotKey) SetClientHeight(value int32) {
    HotKey_SetClientHeight(h.instance, value)
}

// ClientOrigin
func (h *THotKey) ClientOrigin() TPoint {
    return HotKey_GetClientOrigin(h.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (h *THotKey) ClientRect() TRect {
    return HotKey_GetClientRect(h.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (h *THotKey) ClientWidth() int32 {
    return HotKey_GetClientWidth(h.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (h *THotKey) SetClientWidth(value int32) {
    HotKey_SetClientWidth(h.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (h *THotKey) ControlState() TControlState {
    return HotKey_GetControlState(h.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (h *THotKey) SetControlState(value TControlState) {
    HotKey_SetControlState(h.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (h *THotKey) ControlStyle() TControlStyle {
    return HotKey_GetControlStyle(h.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (h *THotKey) SetControlStyle(value TControlStyle) {
    HotKey_SetControlStyle(h.instance, value)
}

// ExplicitLeft
func (h *THotKey) ExplicitLeft() int32 {
    return HotKey_GetExplicitLeft(h.instance)
}

// ExplicitTop
func (h *THotKey) ExplicitTop() int32 {
    return HotKey_GetExplicitTop(h.instance)
}

// ExplicitWidth
func (h *THotKey) ExplicitWidth() int32 {
    return HotKey_GetExplicitWidth(h.instance)
}

// ExplicitHeight
func (h *THotKey) ExplicitHeight() int32 {
    return HotKey_GetExplicitHeight(h.instance)
}

// Floating
func (h *THotKey) Floating() bool {
    return HotKey_GetFloating(h.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (h *THotKey) Parent() *TWinControl {
    return WinControlFromInst(HotKey_GetParent(h.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (h *THotKey) SetParent(value IWinControl) {
    HotKey_SetParent(h.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (h *THotKey) AlignWithMargins() bool {
    return HotKey_GetAlignWithMargins(h.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (h *THotKey) SetAlignWithMargins(value bool) {
    HotKey_SetAlignWithMargins(h.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (h *THotKey) Left() int32 {
    return HotKey_GetLeft(h.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (h *THotKey) SetLeft(value int32) {
    HotKey_SetLeft(h.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (h *THotKey) Top() int32 {
    return HotKey_GetTop(h.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (h *THotKey) SetTop(value int32) {
    HotKey_SetTop(h.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (h *THotKey) Width() int32 {
    return HotKey_GetWidth(h.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (h *THotKey) SetWidth(value int32) {
    HotKey_SetWidth(h.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (h *THotKey) Height() int32 {
    return HotKey_GetHeight(h.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (h *THotKey) SetHeight(value int32) {
    HotKey_SetHeight(h.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (h *THotKey) Cursor() TCursor {
    return HotKey_GetCursor(h.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (h *THotKey) SetCursor(value TCursor) {
    HotKey_SetCursor(h.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (h *THotKey) Margins() *TMargins {
    return MarginsFromInst(HotKey_GetMargins(h.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (h *THotKey) SetMargins(value *TMargins) {
    HotKey_SetMargins(h.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (h *THotKey) CustomHint() *TCustomHint {
    return CustomHintFromInst(HotKey_GetCustomHint(h.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (h *THotKey) SetCustomHint(value IComponent) {
    HotKey_SetCustomHint(h.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (h *THotKey) ComponentCount() int32 {
    return HotKey_GetComponentCount(h.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (h *THotKey) ComponentIndex() int32 {
    return HotKey_GetComponentIndex(h.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (h *THotKey) SetComponentIndex(value int32) {
    HotKey_SetComponentIndex(h.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (h *THotKey) Owner() *TComponent {
    return ComponentFromInst(HotKey_GetOwner(h.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (h *THotKey) Name() string {
    return HotKey_GetName(h.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (h *THotKey) SetName(value string) {
    HotKey_SetName(h.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (h *THotKey) Tag() int {
    return HotKey_GetTag(h.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (h *THotKey) SetTag(value int) {
    HotKey_SetTag(h.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (h *THotKey) DockClients(Index int32) *TControl {
    return ControlFromInst(HotKey_GetDockClients(h.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (h *THotKey) Controls(Index int32) *TControl {
    return ControlFromInst(HotKey_GetControls(h.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (h *THotKey) Components(AIndex int32) *TComponent {
    return ComponentFromInst(HotKey_GetComponents(h.instance, AIndex))
}

