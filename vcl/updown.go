
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

type TUpDown struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewUpDown
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewUpDown(owner IComponent) *TUpDown {
    u := new(TUpDown)
    u.instance = UpDown_Create(CheckPtr(owner))
    u.ptr = unsafe.Pointer(u.instance)
    return u
}

// UpDownFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func UpDownFromInst(inst uintptr) *TUpDown {
    u := new(TUpDown)
    u.instance = inst
    u.ptr = unsafe.Pointer(inst)
    return u
}

// UpDownFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func UpDownFromObj(obj IObject) *TUpDown {
    u := new(TUpDown)
    u.instance = CheckPtr(obj)
    u.ptr = unsafe.Pointer(u.instance)
    return u
}

// UpDownFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func UpDownFromUnsafePointer(ptr unsafe.Pointer) *TUpDown {
    u := new(TUpDown)
    u.instance = uintptr(ptr)
    u.ptr = ptr
    return u
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (u *TUpDown) Free() {
    if u.instance != 0 {
        UpDown_Free(u.instance)
        u.instance = 0
        u.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (u *TUpDown) Instance() uintptr {
    return u.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (u *TUpDown) UnsafeAddr() unsafe.Pointer {
    return u.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (u *TUpDown) IsValid() bool {
    return u.instance != 0
}

// TUpDownClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TUpDownClass() TClass {
    return UpDown_StaticClassType()
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (u *TUpDown) CanFocus() bool {
    return UpDown_CanFocus(u.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (u *TUpDown) ContainsControl(Control IControl) bool {
    return UpDown_ContainsControl(u.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (u *TUpDown) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(UpDown_ControlAtPos(u.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (u *TUpDown) DisableAlign() {
    UpDown_DisableAlign(u.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (u *TUpDown) EnableAlign() {
    UpDown_EnableAlign(u.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (u *TUpDown) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(UpDown_FindChildControl(u.instance, ControlName))
}

// FlipChildren
func (u *TUpDown) FlipChildren(AllLevels bool) {
    UpDown_FlipChildren(u.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (u *TUpDown) Focused() bool {
    return UpDown_Focused(u.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (u *TUpDown) HandleAllocated() bool {
    return UpDown_HandleAllocated(u.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (u *TUpDown) InsertControl(AControl IControl) {
    UpDown_InsertControl(u.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (u *TUpDown) Invalidate() {
    UpDown_Invalidate(u.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (u *TUpDown) PaintTo(DC HDC, X int32, Y int32) {
    UpDown_PaintTo(u.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (u *TUpDown) RemoveControl(AControl IControl) {
    UpDown_RemoveControl(u.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (u *TUpDown) Realign() {
    UpDown_Realign(u.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (u *TUpDown) Repaint() {
    UpDown_Repaint(u.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (u *TUpDown) ScaleBy(M int32, D int32) {
    UpDown_ScaleBy(u.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (u *TUpDown) ScrollBy(DeltaX int32, DeltaY int32) {
    UpDown_ScrollBy(u.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (u *TUpDown) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    UpDown_SetBounds(u.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (u *TUpDown) SetFocus() {
    UpDown_SetFocus(u.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (u *TUpDown) Update() {
    UpDown_Update(u.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (u *TUpDown) UpdateControlState() {
    UpDown_UpdateControlState(u.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (u *TUpDown) BringToFront() {
    UpDown_BringToFront(u.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (u *TUpDown) ClientToScreen(Point TPoint) TPoint {
    return UpDown_ClientToScreen(u.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (u *TUpDown) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return UpDown_ClientToParent(u.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (u *TUpDown) Dragging() bool {
    return UpDown_Dragging(u.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (u *TUpDown) HasParent() bool {
    return UpDown_HasParent(u.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (u *TUpDown) Hide() {
    UpDown_Hide(u.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (u *TUpDown) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return UpDown_Perform(u.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (u *TUpDown) Refresh() {
    UpDown_Refresh(u.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (u *TUpDown) ScreenToClient(Point TPoint) TPoint {
    return UpDown_ScreenToClient(u.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (u *TUpDown) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return UpDown_ParentToClient(u.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (u *TUpDown) SendToBack() {
    UpDown_SendToBack(u.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (u *TUpDown) Show() {
    UpDown_Show(u.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (u *TUpDown) GetTextBuf(Buffer string, BufSize int32) int32 {
    return UpDown_GetTextBuf(u.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (u *TUpDown) GetTextLen() int32 {
    return UpDown_GetTextLen(u.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (u *TUpDown) SetTextBuf(Buffer string) {
    UpDown_SetTextBuf(u.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (u *TUpDown) FindComponent(AName string) *TComponent {
    return ComponentFromInst(UpDown_FindComponent(u.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (u *TUpDown) GetNamePath() string {
    return UpDown_GetNamePath(u.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (u *TUpDown) Assign(Source IObject) {
    UpDown_Assign(u.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (u *TUpDown) DisposeOf() {
    UpDown_DisposeOf(u.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (u *TUpDown) ClassType() TClass {
    return UpDown_ClassType(u.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (u *TUpDown) ClassName() string {
    return UpDown_ClassName(u.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (u *TUpDown) InstanceSize() int32 {
    return UpDown_InstanceSize(u.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (u *TUpDown) InheritsFrom(AClass TClass) bool {
    return UpDown_InheritsFrom(u.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (u *TUpDown) Equals(Obj IObject) bool {
    return UpDown_Equals(u.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (u *TUpDown) GetHashCode() int32 {
    return UpDown_GetHashCode(u.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (u *TUpDown) ToString() string {
    return UpDown_ToString(u.instance)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (u *TUpDown) Anchors() TAnchors {
    return UpDown_GetAnchors(u.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (u *TUpDown) SetAnchors(value TAnchors) {
    UpDown_SetAnchors(u.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (u *TUpDown) DoubleBuffered() bool {
    return UpDown_GetDoubleBuffered(u.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (u *TUpDown) SetDoubleBuffered(value bool) {
    UpDown_SetDoubleBuffered(u.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (u *TUpDown) Enabled() bool {
    return UpDown_GetEnabled(u.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (u *TUpDown) SetEnabled(value bool) {
    UpDown_SetEnabled(u.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (u *TUpDown) Hint() string {
    return UpDown_GetHint(u.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (u *TUpDown) SetHint(value string) {
    UpDown_SetHint(u.instance, value)
}

// Min
func (u *TUpDown) Min() int32 {
    return UpDown_GetMin(u.instance)
}

// SetMin
func (u *TUpDown) SetMin(value int32) {
    UpDown_SetMin(u.instance, value)
}

// Max
func (u *TUpDown) Max() int32 {
    return UpDown_GetMax(u.instance)
}

// SetMax
func (u *TUpDown) SetMax(value int32) {
    UpDown_SetMax(u.instance, value)
}

// Orientation
func (u *TUpDown) Orientation() TUDOrientation {
    return UpDown_GetOrientation(u.instance)
}

// SetOrientation
func (u *TUpDown) SetOrientation(value TUDOrientation) {
    UpDown_SetOrientation(u.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (u *TUpDown) ParentDoubleBuffered() bool {
    return UpDown_GetParentDoubleBuffered(u.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (u *TUpDown) SetParentDoubleBuffered(value bool) {
    UpDown_SetParentDoubleBuffered(u.instance, value)
}

// ParentShowHint
func (u *TUpDown) ParentShowHint() bool {
    return UpDown_GetParentShowHint(u.instance)
}

// SetParentShowHint
func (u *TUpDown) SetParentShowHint(value bool) {
    UpDown_SetParentShowHint(u.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (u *TUpDown) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(UpDown_GetPopupMenu(u.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (u *TUpDown) SetPopupMenu(value IComponent) {
    UpDown_SetPopupMenu(u.instance, CheckPtr(value))
}

// Position
func (u *TUpDown) Position() int32 {
    return UpDown_GetPosition(u.instance)
}

// SetPosition
func (u *TUpDown) SetPosition(value int32) {
    UpDown_SetPosition(u.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (u *TUpDown) ShowHint() bool {
    return UpDown_GetShowHint(u.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (u *TUpDown) SetShowHint(value bool) {
    UpDown_SetShowHint(u.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (u *TUpDown) TabOrder() TTabOrder {
    return UpDown_GetTabOrder(u.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (u *TUpDown) SetTabOrder(value TTabOrder) {
    UpDown_SetTabOrder(u.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (u *TUpDown) TabStop() bool {
    return UpDown_GetTabStop(u.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (u *TUpDown) SetTabStop(value bool) {
    UpDown_SetTabStop(u.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (u *TUpDown) Visible() bool {
    return UpDown_GetVisible(u.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (u *TUpDown) SetVisible(value bool) {
    UpDown_SetVisible(u.instance, value)
}

// Wrap
func (u *TUpDown) Wrap() bool {
    return UpDown_GetWrap(u.instance)
}

// SetWrap
func (u *TUpDown) SetWrap(value bool) {
    UpDown_SetWrap(u.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (u *TUpDown) StyleElements() TStyleElements {
    return UpDown_GetStyleElements(u.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (u *TUpDown) SetStyleElements(value TStyleElements) {
    UpDown_SetStyleElements(u.instance, value)
}

// SetOnChanging
func (u *TUpDown) SetOnChanging(fn TUDChangingEvent) {
    UpDown_SetOnChanging(u.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (u *TUpDown) SetOnContextPopup(fn TContextPopupEvent) {
    UpDown_SetOnContextPopup(u.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (u *TUpDown) SetOnClick(fn TUDClickEvent) {
    UpDown_SetOnClick(u.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (u *TUpDown) SetOnEnter(fn TNotifyEvent) {
    UpDown_SetOnEnter(u.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (u *TUpDown) SetOnExit(fn TNotifyEvent) {
    UpDown_SetOnExit(u.instance, fn)
}

// SetOnMouseActivate
func (u *TUpDown) SetOnMouseActivate(fn TMouseActivateEvent) {
    UpDown_SetOnMouseActivate(u.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (u *TUpDown) SetOnMouseDown(fn TMouseEvent) {
    UpDown_SetOnMouseDown(u.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (u *TUpDown) SetOnMouseEnter(fn TNotifyEvent) {
    UpDown_SetOnMouseEnter(u.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (u *TUpDown) SetOnMouseLeave(fn TNotifyEvent) {
    UpDown_SetOnMouseLeave(u.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (u *TUpDown) SetOnMouseMove(fn TMouseMoveEvent) {
    UpDown_SetOnMouseMove(u.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (u *TUpDown) SetOnMouseUp(fn TMouseEvent) {
    UpDown_SetOnMouseUp(u.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (u *TUpDown) DockClientCount() int32 {
    return UpDown_GetDockClientCount(u.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (u *TUpDown) DockSite() bool {
    return UpDown_GetDockSite(u.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (u *TUpDown) SetDockSite(value bool) {
    UpDown_SetDockSite(u.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (u *TUpDown) AlignDisabled() bool {
    return UpDown_GetAlignDisabled(u.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (u *TUpDown) MouseInClient() bool {
    return UpDown_GetMouseInClient(u.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (u *TUpDown) VisibleDockClientCount() int32 {
    return UpDown_GetVisibleDockClientCount(u.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (u *TUpDown) Brush() *TBrush {
    return BrushFromInst(UpDown_GetBrush(u.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (u *TUpDown) ControlCount() int32 {
    return UpDown_GetControlCount(u.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (u *TUpDown) Handle() HWND {
    return UpDown_GetHandle(u.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (u *TUpDown) ParentWindow() HWND {
    return UpDown_GetParentWindow(u.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (u *TUpDown) SetParentWindow(value HWND) {
    UpDown_SetParentWindow(u.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (u *TUpDown) UseDockManager() bool {
    return UpDown_GetUseDockManager(u.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (u *TUpDown) SetUseDockManager(value bool) {
    UpDown_SetUseDockManager(u.instance, value)
}

// Action
func (u *TUpDown) Action() *TAction {
    return ActionFromInst(UpDown_GetAction(u.instance))
}

// SetAction
func (u *TUpDown) SetAction(value IComponent) {
    UpDown_SetAction(u.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (u *TUpDown) Align() TAlign {
    return UpDown_GetAlign(u.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (u *TUpDown) SetAlign(value TAlign) {
    UpDown_SetAlign(u.instance, value)
}

// BiDiMode
func (u *TUpDown) BiDiMode() TBiDiMode {
    return UpDown_GetBiDiMode(u.instance)
}

// SetBiDiMode
func (u *TUpDown) SetBiDiMode(value TBiDiMode) {
    UpDown_SetBiDiMode(u.instance, value)
}

// BoundsRect
func (u *TUpDown) BoundsRect() TRect {
    return UpDown_GetBoundsRect(u.instance)
}

// SetBoundsRect
func (u *TUpDown) SetBoundsRect(value TRect) {
    UpDown_SetBoundsRect(u.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (u *TUpDown) ClientHeight() int32 {
    return UpDown_GetClientHeight(u.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (u *TUpDown) SetClientHeight(value int32) {
    UpDown_SetClientHeight(u.instance, value)
}

// ClientOrigin
func (u *TUpDown) ClientOrigin() TPoint {
    return UpDown_GetClientOrigin(u.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (u *TUpDown) ClientRect() TRect {
    return UpDown_GetClientRect(u.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (u *TUpDown) ClientWidth() int32 {
    return UpDown_GetClientWidth(u.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (u *TUpDown) SetClientWidth(value int32) {
    UpDown_SetClientWidth(u.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (u *TUpDown) ControlState() TControlState {
    return UpDown_GetControlState(u.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (u *TUpDown) SetControlState(value TControlState) {
    UpDown_SetControlState(u.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (u *TUpDown) ControlStyle() TControlStyle {
    return UpDown_GetControlStyle(u.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (u *TUpDown) SetControlStyle(value TControlStyle) {
    UpDown_SetControlStyle(u.instance, value)
}

// ExplicitLeft
func (u *TUpDown) ExplicitLeft() int32 {
    return UpDown_GetExplicitLeft(u.instance)
}

// ExplicitTop
func (u *TUpDown) ExplicitTop() int32 {
    return UpDown_GetExplicitTop(u.instance)
}

// ExplicitWidth
func (u *TUpDown) ExplicitWidth() int32 {
    return UpDown_GetExplicitWidth(u.instance)
}

// ExplicitHeight
func (u *TUpDown) ExplicitHeight() int32 {
    return UpDown_GetExplicitHeight(u.instance)
}

// Floating
func (u *TUpDown) Floating() bool {
    return UpDown_GetFloating(u.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (u *TUpDown) Parent() *TWinControl {
    return WinControlFromInst(UpDown_GetParent(u.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (u *TUpDown) SetParent(value IWinControl) {
    UpDown_SetParent(u.instance, CheckPtr(value))
}

// SetOnGesture
func (u *TUpDown) SetOnGesture(fn TGestureEvent) {
    UpDown_SetOnGesture(u.instance, fn)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (u *TUpDown) AlignWithMargins() bool {
    return UpDown_GetAlignWithMargins(u.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (u *TUpDown) SetAlignWithMargins(value bool) {
    UpDown_SetAlignWithMargins(u.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (u *TUpDown) Left() int32 {
    return UpDown_GetLeft(u.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (u *TUpDown) SetLeft(value int32) {
    UpDown_SetLeft(u.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (u *TUpDown) Top() int32 {
    return UpDown_GetTop(u.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (u *TUpDown) SetTop(value int32) {
    UpDown_SetTop(u.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (u *TUpDown) Width() int32 {
    return UpDown_GetWidth(u.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (u *TUpDown) SetWidth(value int32) {
    UpDown_SetWidth(u.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (u *TUpDown) Height() int32 {
    return UpDown_GetHeight(u.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (u *TUpDown) SetHeight(value int32) {
    UpDown_SetHeight(u.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (u *TUpDown) Cursor() TCursor {
    return UpDown_GetCursor(u.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (u *TUpDown) SetCursor(value TCursor) {
    UpDown_SetCursor(u.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (u *TUpDown) Margins() *TMargins {
    return MarginsFromInst(UpDown_GetMargins(u.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (u *TUpDown) SetMargins(value *TMargins) {
    UpDown_SetMargins(u.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (u *TUpDown) CustomHint() *TCustomHint {
    return CustomHintFromInst(UpDown_GetCustomHint(u.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (u *TUpDown) SetCustomHint(value IComponent) {
    UpDown_SetCustomHint(u.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (u *TUpDown) ComponentCount() int32 {
    return UpDown_GetComponentCount(u.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (u *TUpDown) ComponentIndex() int32 {
    return UpDown_GetComponentIndex(u.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (u *TUpDown) SetComponentIndex(value int32) {
    UpDown_SetComponentIndex(u.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (u *TUpDown) Owner() *TComponent {
    return ComponentFromInst(UpDown_GetOwner(u.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (u *TUpDown) Name() string {
    return UpDown_GetName(u.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (u *TUpDown) SetName(value string) {
    UpDown_SetName(u.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (u *TUpDown) Tag() int {
    return UpDown_GetTag(u.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (u *TUpDown) SetTag(value int) {
    UpDown_SetTag(u.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (u *TUpDown) DockClients(Index int32) *TControl {
    return ControlFromInst(UpDown_GetDockClients(u.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (u *TUpDown) Controls(Index int32) *TControl {
    return ControlFromInst(UpDown_GetControls(u.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (u *TUpDown) Components(AIndex int32) *TComponent {
    return ComponentFromInst(UpDown_GetComponents(u.instance, AIndex))
}

