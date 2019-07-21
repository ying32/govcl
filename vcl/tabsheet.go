
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

type TTabSheet struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTabSheet
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTabSheet(owner IComponent) *TTabSheet {
    t := new(TTabSheet)
    t.instance = TabSheet_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TabSheetFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TabSheetFromInst(inst uintptr) *TTabSheet {
    t := new(TTabSheet)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TabSheetFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TabSheetFromObj(obj IObject) *TTabSheet {
    t := new(TTabSheet)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TabSheetFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TabSheetFromUnsafePointer(ptr unsafe.Pointer) *TTabSheet {
    t := new(TTabSheet)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTabSheet) Free() {
    if t.instance != 0 {
        TabSheet_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTabSheet) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTabSheet) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTabSheet) IsValid() bool {
    return t.instance != 0
}

// TTabSheetClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTabSheetClass() TClass {
    return TabSheet_StaticClassType()
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (t *TTabSheet) CanFocus() bool {
    return TabSheet_CanFocus(t.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (t *TTabSheet) ContainsControl(Control IControl) bool {
    return TabSheet_ContainsControl(t.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (t *TTabSheet) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(TabSheet_ControlAtPos(t.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (t *TTabSheet) DisableAlign() {
    TabSheet_DisableAlign(t.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (t *TTabSheet) EnableAlign() {
    TabSheet_EnableAlign(t.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (t *TTabSheet) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(TabSheet_FindChildControl(t.instance, ControlName))
}

// FlipChildren
func (t *TTabSheet) FlipChildren(AllLevels bool) {
    TabSheet_FlipChildren(t.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (t *TTabSheet) Focused() bool {
    return TabSheet_Focused(t.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (t *TTabSheet) HandleAllocated() bool {
    return TabSheet_HandleAllocated(t.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (t *TTabSheet) InsertControl(AControl IControl) {
    TabSheet_InsertControl(t.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (t *TTabSheet) Invalidate() {
    TabSheet_Invalidate(t.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (t *TTabSheet) PaintTo(DC HDC, X int32, Y int32) {
    TabSheet_PaintTo(t.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (t *TTabSheet) RemoveControl(AControl IControl) {
    TabSheet_RemoveControl(t.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (t *TTabSheet) Realign() {
    TabSheet_Realign(t.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (t *TTabSheet) Repaint() {
    TabSheet_Repaint(t.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (t *TTabSheet) ScaleBy(M int32, D int32) {
    TabSheet_ScaleBy(t.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (t *TTabSheet) ScrollBy(DeltaX int32, DeltaY int32) {
    TabSheet_ScrollBy(t.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (t *TTabSheet) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    TabSheet_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (t *TTabSheet) SetFocus() {
    TabSheet_SetFocus(t.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (t *TTabSheet) Update() {
    TabSheet_Update(t.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (t *TTabSheet) UpdateControlState() {
    TabSheet_UpdateControlState(t.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (t *TTabSheet) BringToFront() {
    TabSheet_BringToFront(t.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (t *TTabSheet) ClientToScreen(Point TPoint) TPoint {
    return TabSheet_ClientToScreen(t.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (t *TTabSheet) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return TabSheet_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (t *TTabSheet) Dragging() bool {
    return TabSheet_Dragging(t.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (t *TTabSheet) HasParent() bool {
    return TabSheet_HasParent(t.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (t *TTabSheet) Hide() {
    TabSheet_Hide(t.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (t *TTabSheet) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return TabSheet_Perform(t.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (t *TTabSheet) Refresh() {
    TabSheet_Refresh(t.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (t *TTabSheet) ScreenToClient(Point TPoint) TPoint {
    return TabSheet_ScreenToClient(t.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (t *TTabSheet) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return TabSheet_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (t *TTabSheet) SendToBack() {
    TabSheet_SendToBack(t.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (t *TTabSheet) Show() {
    TabSheet_Show(t.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (t *TTabSheet) GetTextBuf(Buffer string, BufSize int32) int32 {
    return TabSheet_GetTextBuf(t.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (t *TTabSheet) GetTextLen() int32 {
    return TabSheet_GetTextLen(t.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (t *TTabSheet) SetTextBuf(Buffer string) {
    TabSheet_SetTextBuf(t.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (t *TTabSheet) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TabSheet_FindComponent(t.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTabSheet) GetNamePath() string {
    return TabSheet_GetNamePath(t.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTabSheet) Assign(Source IObject) {
    TabSheet_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTabSheet) DisposeOf() {
    TabSheet_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTabSheet) ClassType() TClass {
    return TabSheet_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTabSheet) ClassName() string {
    return TabSheet_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTabSheet) InstanceSize() int32 {
    return TabSheet_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTabSheet) InheritsFrom(AClass TClass) bool {
    return TabSheet_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTabSheet) Equals(Obj IObject) bool {
    return TabSheet_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTabSheet) GetHashCode() int32 {
    return TabSheet_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTabSheet) ToString() string {
    return TabSheet_ToString(t.instance)
}

// PageControl
func (t *TTabSheet) PageControl() *TPageControl {
    return PageControlFromInst(TabSheet_GetPageControl(t.instance))
}

// SetPageControl
func (t *TTabSheet) SetPageControl(value IWinControl) {
    TabSheet_SetPageControl(t.instance, CheckPtr(value))
}

// TabIndex
func (t *TTabSheet) TabIndex() int32 {
    return TabSheet_GetTabIndex(t.instance)
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (t *TTabSheet) BorderWidth() int32 {
    return TabSheet_GetBorderWidth(t.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (t *TTabSheet) SetBorderWidth(value int32) {
    TabSheet_SetBorderWidth(t.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (t *TTabSheet) Caption() string {
    return TabSheet_GetCaption(t.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (t *TTabSheet) SetCaption(value string) {
    TabSheet_SetCaption(t.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (t *TTabSheet) DoubleBuffered() bool {
    return TabSheet_GetDoubleBuffered(t.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (t *TTabSheet) SetDoubleBuffered(value bool) {
    TabSheet_SetDoubleBuffered(t.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (t *TTabSheet) DragMode() TDragMode {
    return TabSheet_GetDragMode(t.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (t *TTabSheet) SetDragMode(value TDragMode) {
    TabSheet_SetDragMode(t.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TTabSheet) Enabled() bool {
    return TabSheet_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TTabSheet) SetEnabled(value bool) {
    TabSheet_SetEnabled(t.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (t *TTabSheet) Font() *TFont {
    return FontFromInst(TabSheet_GetFont(t.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (t *TTabSheet) SetFont(value *TFont) {
    TabSheet_SetFont(t.instance, CheckPtr(value))
}

// Height
// CN: 获取高度。
// EN: Get height.
func (t *TTabSheet) Height() int32 {
    return TabSheet_GetHeight(t.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (t *TTabSheet) SetHeight(value int32) {
    TabSheet_SetHeight(t.instance, value)
}

// Highlighted
func (t *TTabSheet) Highlighted() bool {
    return TabSheet_GetHighlighted(t.instance)
}

// SetHighlighted
func (t *TTabSheet) SetHighlighted(value bool) {
    TabSheet_SetHighlighted(t.instance, value)
}

// ImageIndex
// CN: 获取图像在images中的索引。
// EN: .
func (t *TTabSheet) ImageIndex() int32 {
    return TabSheet_GetImageIndex(t.instance)
}

// SetImageIndex
// CN: 设置图像在images中的索引。
// EN: .
func (t *TTabSheet) SetImageIndex(value int32) {
    TabSheet_SetImageIndex(t.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (t *TTabSheet) Left() int32 {
    return TabSheet_GetLeft(t.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (t *TTabSheet) SetLeft(value int32) {
    TabSheet_SetLeft(t.instance, value)
}

// Constraints
func (t *TTabSheet) Constraints() *TSizeConstraints {
    return SizeConstraintsFromInst(TabSheet_GetConstraints(t.instance))
}

// SetConstraints
func (t *TTabSheet) SetConstraints(value *TSizeConstraints) {
    TabSheet_SetConstraints(t.instance, CheckPtr(value))
}

// PageIndex
func (t *TTabSheet) PageIndex() int32 {
    return TabSheet_GetPageIndex(t.instance)
}

// SetPageIndex
func (t *TTabSheet) SetPageIndex(value int32) {
    TabSheet_SetPageIndex(t.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (t *TTabSheet) ParentDoubleBuffered() bool {
    return TabSheet_GetParentDoubleBuffered(t.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (t *TTabSheet) SetParentDoubleBuffered(value bool) {
    TabSheet_SetParentDoubleBuffered(t.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (t *TTabSheet) ParentFont() bool {
    return TabSheet_GetParentFont(t.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (t *TTabSheet) SetParentFont(value bool) {
    TabSheet_SetParentFont(t.instance, value)
}

// ParentShowHint
func (t *TTabSheet) ParentShowHint() bool {
    return TabSheet_GetParentShowHint(t.instance)
}

// SetParentShowHint
func (t *TTabSheet) SetParentShowHint(value bool) {
    TabSheet_SetParentShowHint(t.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (t *TTabSheet) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TabSheet_GetPopupMenu(t.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (t *TTabSheet) SetPopupMenu(value IComponent) {
    TabSheet_SetPopupMenu(t.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (t *TTabSheet) ShowHint() bool {
    return TabSheet_GetShowHint(t.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (t *TTabSheet) SetShowHint(value bool) {
    TabSheet_SetShowHint(t.instance, value)
}

// TabVisible
func (t *TTabSheet) TabVisible() bool {
    return TabSheet_GetTabVisible(t.instance)
}

// SetTabVisible
func (t *TTabSheet) SetTabVisible(value bool) {
    TabSheet_SetTabVisible(t.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (t *TTabSheet) Top() int32 {
    return TabSheet_GetTop(t.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (t *TTabSheet) SetTop(value int32) {
    TabSheet_SetTop(t.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TTabSheet) Visible() bool {
    return TabSheet_GetVisible(t.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TTabSheet) SetVisible(value bool) {
    TabSheet_SetVisible(t.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (t *TTabSheet) Width() int32 {
    return TabSheet_GetWidth(t.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (t *TTabSheet) SetWidth(value int32) {
    TabSheet_SetWidth(t.instance, value)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (t *TTabSheet) SetOnContextPopup(fn TContextPopupEvent) {
    TabSheet_SetOnContextPopup(t.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (t *TTabSheet) SetOnDragDrop(fn TDragDropEvent) {
    TabSheet_SetOnDragDrop(t.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (t *TTabSheet) SetOnDragOver(fn TDragOverEvent) {
    TabSheet_SetOnDragOver(t.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (t *TTabSheet) SetOnEndDrag(fn TEndDragEvent) {
    TabSheet_SetOnEndDrag(t.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (t *TTabSheet) SetOnEnter(fn TNotifyEvent) {
    TabSheet_SetOnEnter(t.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (t *TTabSheet) SetOnExit(fn TNotifyEvent) {
    TabSheet_SetOnExit(t.instance, fn)
}

// SetOnGesture
func (t *TTabSheet) SetOnGesture(fn TGestureEvent) {
    TabSheet_SetOnGesture(t.instance, fn)
}

// SetOnHide
// CN: 设置隐藏事件。
// EN: .
func (t *TTabSheet) SetOnHide(fn TNotifyEvent) {
    TabSheet_SetOnHide(t.instance, fn)
}

// SetOnMouseActivate
func (t *TTabSheet) SetOnMouseActivate(fn TMouseActivateEvent) {
    TabSheet_SetOnMouseActivate(t.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (t *TTabSheet) SetOnMouseDown(fn TMouseEvent) {
    TabSheet_SetOnMouseDown(t.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (t *TTabSheet) SetOnMouseEnter(fn TNotifyEvent) {
    TabSheet_SetOnMouseEnter(t.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (t *TTabSheet) SetOnMouseLeave(fn TNotifyEvent) {
    TabSheet_SetOnMouseLeave(t.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (t *TTabSheet) SetOnMouseMove(fn TMouseMoveEvent) {
    TabSheet_SetOnMouseMove(t.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (t *TTabSheet) SetOnMouseUp(fn TMouseEvent) {
    TabSheet_SetOnMouseUp(t.instance, fn)
}

// SetOnResize
// CN: 设置大小被改变事件。
// EN: .
func (t *TTabSheet) SetOnResize(fn TNotifyEvent) {
    TabSheet_SetOnResize(t.instance, fn)
}

// SetOnShow
// CN: 设置显示事件。
// EN: .
func (t *TTabSheet) SetOnShow(fn TNotifyEvent) {
    TabSheet_SetOnShow(t.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (t *TTabSheet) DockClientCount() int32 {
    return TabSheet_GetDockClientCount(t.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (t *TTabSheet) DockSite() bool {
    return TabSheet_GetDockSite(t.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (t *TTabSheet) SetDockSite(value bool) {
    TabSheet_SetDockSite(t.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (t *TTabSheet) AlignDisabled() bool {
    return TabSheet_GetAlignDisabled(t.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (t *TTabSheet) MouseInClient() bool {
    return TabSheet_GetMouseInClient(t.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (t *TTabSheet) VisibleDockClientCount() int32 {
    return TabSheet_GetVisibleDockClientCount(t.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (t *TTabSheet) Brush() *TBrush {
    return BrushFromInst(TabSheet_GetBrush(t.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (t *TTabSheet) ControlCount() int32 {
    return TabSheet_GetControlCount(t.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (t *TTabSheet) Handle() HWND {
    return TabSheet_GetHandle(t.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (t *TTabSheet) ParentWindow() HWND {
    return TabSheet_GetParentWindow(t.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (t *TTabSheet) SetParentWindow(value HWND) {
    TabSheet_SetParentWindow(t.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (t *TTabSheet) TabOrder() TTabOrder {
    return TabSheet_GetTabOrder(t.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (t *TTabSheet) SetTabOrder(value TTabOrder) {
    TabSheet_SetTabOrder(t.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (t *TTabSheet) TabStop() bool {
    return TabSheet_GetTabStop(t.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (t *TTabSheet) SetTabStop(value bool) {
    TabSheet_SetTabStop(t.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (t *TTabSheet) UseDockManager() bool {
    return TabSheet_GetUseDockManager(t.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (t *TTabSheet) SetUseDockManager(value bool) {
    TabSheet_SetUseDockManager(t.instance, value)
}

// Action
func (t *TTabSheet) Action() *TAction {
    return ActionFromInst(TabSheet_GetAction(t.instance))
}

// SetAction
func (t *TTabSheet) SetAction(value IComponent) {
    TabSheet_SetAction(t.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (t *TTabSheet) Align() TAlign {
    return TabSheet_GetAlign(t.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (t *TTabSheet) SetAlign(value TAlign) {
    TabSheet_SetAlign(t.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (t *TTabSheet) Anchors() TAnchors {
    return TabSheet_GetAnchors(t.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (t *TTabSheet) SetAnchors(value TAnchors) {
    TabSheet_SetAnchors(t.instance, value)
}

// BiDiMode
func (t *TTabSheet) BiDiMode() TBiDiMode {
    return TabSheet_GetBiDiMode(t.instance)
}

// SetBiDiMode
func (t *TTabSheet) SetBiDiMode(value TBiDiMode) {
    TabSheet_SetBiDiMode(t.instance, value)
}

// BoundsRect
func (t *TTabSheet) BoundsRect() TRect {
    return TabSheet_GetBoundsRect(t.instance)
}

// SetBoundsRect
func (t *TTabSheet) SetBoundsRect(value TRect) {
    TabSheet_SetBoundsRect(t.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (t *TTabSheet) ClientHeight() int32 {
    return TabSheet_GetClientHeight(t.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (t *TTabSheet) SetClientHeight(value int32) {
    TabSheet_SetClientHeight(t.instance, value)
}

// ClientOrigin
func (t *TTabSheet) ClientOrigin() TPoint {
    return TabSheet_GetClientOrigin(t.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (t *TTabSheet) ClientRect() TRect {
    return TabSheet_GetClientRect(t.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (t *TTabSheet) ClientWidth() int32 {
    return TabSheet_GetClientWidth(t.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (t *TTabSheet) SetClientWidth(value int32) {
    TabSheet_SetClientWidth(t.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (t *TTabSheet) ControlState() TControlState {
    return TabSheet_GetControlState(t.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (t *TTabSheet) SetControlState(value TControlState) {
    TabSheet_SetControlState(t.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (t *TTabSheet) ControlStyle() TControlStyle {
    return TabSheet_GetControlStyle(t.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (t *TTabSheet) SetControlStyle(value TControlStyle) {
    TabSheet_SetControlStyle(t.instance, value)
}

// ExplicitLeft
func (t *TTabSheet) ExplicitLeft() int32 {
    return TabSheet_GetExplicitLeft(t.instance)
}

// ExplicitTop
func (t *TTabSheet) ExplicitTop() int32 {
    return TabSheet_GetExplicitTop(t.instance)
}

// ExplicitWidth
func (t *TTabSheet) ExplicitWidth() int32 {
    return TabSheet_GetExplicitWidth(t.instance)
}

// ExplicitHeight
func (t *TTabSheet) ExplicitHeight() int32 {
    return TabSheet_GetExplicitHeight(t.instance)
}

// Floating
func (t *TTabSheet) Floating() bool {
    return TabSheet_GetFloating(t.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TTabSheet) Parent() *TWinControl {
    return WinControlFromInst(TabSheet_GetParent(t.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (t *TTabSheet) SetParent(value IWinControl) {
    TabSheet_SetParent(t.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (t *TTabSheet) StyleElements() TStyleElements {
    return TabSheet_GetStyleElements(t.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (t *TTabSheet) SetStyleElements(value TStyleElements) {
    TabSheet_SetStyleElements(t.instance, value)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (t *TTabSheet) AlignWithMargins() bool {
    return TabSheet_GetAlignWithMargins(t.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (t *TTabSheet) SetAlignWithMargins(value bool) {
    TabSheet_SetAlignWithMargins(t.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (t *TTabSheet) Cursor() TCursor {
    return TabSheet_GetCursor(t.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (t *TTabSheet) SetCursor(value TCursor) {
    TabSheet_SetCursor(t.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (t *TTabSheet) Hint() string {
    return TabSheet_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (t *TTabSheet) SetHint(value string) {
    TabSheet_SetHint(t.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (t *TTabSheet) Margins() *TMargins {
    return MarginsFromInst(TabSheet_GetMargins(t.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (t *TTabSheet) SetMargins(value *TMargins) {
    TabSheet_SetMargins(t.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (t *TTabSheet) CustomHint() *TCustomHint {
    return CustomHintFromInst(TabSheet_GetCustomHint(t.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (t *TTabSheet) SetCustomHint(value IComponent) {
    TabSheet_SetCustomHint(t.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTabSheet) ComponentCount() int32 {
    return TabSheet_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TTabSheet) ComponentIndex() int32 {
    return TabSheet_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TTabSheet) SetComponentIndex(value int32) {
    TabSheet_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTabSheet) Owner() *TComponent {
    return ComponentFromInst(TabSheet_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTabSheet) Name() string {
    return TabSheet_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTabSheet) SetName(value string) {
    TabSheet_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTabSheet) Tag() int {
    return TabSheet_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTabSheet) SetTag(value int) {
    TabSheet_SetTag(t.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (t *TTabSheet) DockClients(Index int32) *TControl {
    return ControlFromInst(TabSheet_GetDockClients(t.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (t *TTabSheet) Controls(Index int32) *TControl {
    return ControlFromInst(TabSheet_GetControls(t.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTabSheet) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TabSheet_GetComponents(t.instance, AIndex))
}

