
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

type TLinkLabel struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewLinkLabel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewLinkLabel(owner IComponent) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = LinkLabel_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// LinkLabelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func LinkLabelFromInst(inst uintptr) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// LinkLabelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func LinkLabelFromObj(obj IObject) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// LinkLabelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func LinkLabelFromUnsafePointer(ptr unsafe.Pointer) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TLinkLabel) Free() {
    if l.instance != 0 {
        LinkLabel_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TLinkLabel) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TLinkLabel) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TLinkLabel) IsValid() bool {
    return l.instance != 0
}

// TLinkLabelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TLinkLabelClass() TClass {
    return LinkLabel_StaticClassType()
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (l *TLinkLabel) CanFocus() bool {
    return LinkLabel_CanFocus(l.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (l *TLinkLabel) ContainsControl(Control IControl) bool {
    return LinkLabel_ContainsControl(l.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (l *TLinkLabel) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(LinkLabel_ControlAtPos(l.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (l *TLinkLabel) DisableAlign() {
    LinkLabel_DisableAlign(l.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (l *TLinkLabel) EnableAlign() {
    LinkLabel_EnableAlign(l.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (l *TLinkLabel) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(LinkLabel_FindChildControl(l.instance, ControlName))
}

// FlipChildren
func (l *TLinkLabel) FlipChildren(AllLevels bool) {
    LinkLabel_FlipChildren(l.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (l *TLinkLabel) Focused() bool {
    return LinkLabel_Focused(l.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (l *TLinkLabel) HandleAllocated() bool {
    return LinkLabel_HandleAllocated(l.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (l *TLinkLabel) InsertControl(AControl IControl) {
    LinkLabel_InsertControl(l.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (l *TLinkLabel) Invalidate() {
    LinkLabel_Invalidate(l.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (l *TLinkLabel) PaintTo(DC HDC, X int32, Y int32) {
    LinkLabel_PaintTo(l.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (l *TLinkLabel) RemoveControl(AControl IControl) {
    LinkLabel_RemoveControl(l.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (l *TLinkLabel) Realign() {
    LinkLabel_Realign(l.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (l *TLinkLabel) Repaint() {
    LinkLabel_Repaint(l.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (l *TLinkLabel) ScaleBy(M int32, D int32) {
    LinkLabel_ScaleBy(l.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (l *TLinkLabel) ScrollBy(DeltaX int32, DeltaY int32) {
    LinkLabel_ScrollBy(l.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (l *TLinkLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    LinkLabel_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (l *TLinkLabel) SetFocus() {
    LinkLabel_SetFocus(l.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (l *TLinkLabel) Update() {
    LinkLabel_Update(l.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (l *TLinkLabel) UpdateControlState() {
    LinkLabel_UpdateControlState(l.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (l *TLinkLabel) BringToFront() {
    LinkLabel_BringToFront(l.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (l *TLinkLabel) ClientToScreen(Point TPoint) TPoint {
    return LinkLabel_ClientToScreen(l.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (l *TLinkLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return LinkLabel_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (l *TLinkLabel) Dragging() bool {
    return LinkLabel_Dragging(l.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (l *TLinkLabel) HasParent() bool {
    return LinkLabel_HasParent(l.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (l *TLinkLabel) Hide() {
    LinkLabel_Hide(l.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (l *TLinkLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return LinkLabel_Perform(l.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (l *TLinkLabel) Refresh() {
    LinkLabel_Refresh(l.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (l *TLinkLabel) ScreenToClient(Point TPoint) TPoint {
    return LinkLabel_ScreenToClient(l.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (l *TLinkLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return LinkLabel_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (l *TLinkLabel) SendToBack() {
    LinkLabel_SendToBack(l.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (l *TLinkLabel) Show() {
    LinkLabel_Show(l.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (l *TLinkLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return LinkLabel_GetTextBuf(l.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (l *TLinkLabel) GetTextLen() int32 {
    return LinkLabel_GetTextLen(l.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (l *TLinkLabel) SetTextBuf(Buffer string) {
    LinkLabel_SetTextBuf(l.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (l *TLinkLabel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(LinkLabel_FindComponent(l.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TLinkLabel) GetNamePath() string {
    return LinkLabel_GetNamePath(l.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TLinkLabel) Assign(Source IObject) {
    LinkLabel_Assign(l.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TLinkLabel) DisposeOf() {
    LinkLabel_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TLinkLabel) ClassType() TClass {
    return LinkLabel_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TLinkLabel) ClassName() string {
    return LinkLabel_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TLinkLabel) InstanceSize() int32 {
    return LinkLabel_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TLinkLabel) InheritsFrom(AClass TClass) bool {
    return LinkLabel_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TLinkLabel) Equals(Obj IObject) bool {
    return LinkLabel_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TLinkLabel) GetHashCode() int32 {
    return LinkLabel_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TLinkLabel) ToString() string {
    return LinkLabel_ToString(l.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (l *TLinkLabel) Align() TAlign {
    return LinkLabel_GetAlign(l.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (l *TLinkLabel) SetAlign(value TAlign) {
    LinkLabel_SetAlign(l.instance, value)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (l *TLinkLabel) Alignment() TLinkAlignment {
    return LinkLabel_GetAlignment(l.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (l *TLinkLabel) SetAlignment(value TLinkAlignment) {
    LinkLabel_SetAlignment(l.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (l *TLinkLabel) Anchors() TAnchors {
    return LinkLabel_GetAnchors(l.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (l *TLinkLabel) SetAnchors(value TAnchors) {
    LinkLabel_SetAnchors(l.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (l *TLinkLabel) AutoSize() bool {
    return LinkLabel_GetAutoSize(l.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (l *TLinkLabel) SetAutoSize(value bool) {
    LinkLabel_SetAutoSize(l.instance, value)
}

// BevelEdges
func (l *TLinkLabel) BevelEdges() TBevelEdges {
    return LinkLabel_GetBevelEdges(l.instance)
}

// SetBevelEdges
func (l *TLinkLabel) SetBevelEdges(value TBevelEdges) {
    LinkLabel_SetBevelEdges(l.instance, value)
}

// BevelInner
func (l *TLinkLabel) BevelInner() TBevelCut {
    return LinkLabel_GetBevelInner(l.instance)
}

// SetBevelInner
func (l *TLinkLabel) SetBevelInner(value TBevelCut) {
    LinkLabel_SetBevelInner(l.instance, value)
}

// BevelKind
func (l *TLinkLabel) BevelKind() TBevelKind {
    return LinkLabel_GetBevelKind(l.instance)
}

// SetBevelKind
func (l *TLinkLabel) SetBevelKind(value TBevelKind) {
    LinkLabel_SetBevelKind(l.instance, value)
}

// BevelOuter
func (l *TLinkLabel) BevelOuter() TBevelCut {
    return LinkLabel_GetBevelOuter(l.instance)
}

// SetBevelOuter
func (l *TLinkLabel) SetBevelOuter(value TBevelCut) {
    LinkLabel_SetBevelOuter(l.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (l *TLinkLabel) Caption() string {
    return LinkLabel_GetCaption(l.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (l *TLinkLabel) SetCaption(value string) {
    LinkLabel_SetCaption(l.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (l *TLinkLabel) Color() TColor {
    return LinkLabel_GetColor(l.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (l *TLinkLabel) SetColor(value TColor) {
    LinkLabel_SetColor(l.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (l *TLinkLabel) DragCursor() TCursor {
    return LinkLabel_GetDragCursor(l.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (l *TLinkLabel) SetDragCursor(value TCursor) {
    LinkLabel_SetDragCursor(l.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (l *TLinkLabel) DragKind() TDragKind {
    return LinkLabel_GetDragKind(l.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (l *TLinkLabel) SetDragKind(value TDragKind) {
    LinkLabel_SetDragKind(l.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (l *TLinkLabel) DragMode() TDragMode {
    return LinkLabel_GetDragMode(l.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (l *TLinkLabel) SetDragMode(value TDragMode) {
    LinkLabel_SetDragMode(l.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (l *TLinkLabel) Enabled() bool {
    return LinkLabel_GetEnabled(l.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (l *TLinkLabel) SetEnabled(value bool) {
    LinkLabel_SetEnabled(l.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (l *TLinkLabel) Font() *TFont {
    return FontFromInst(LinkLabel_GetFont(l.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (l *TLinkLabel) SetFont(value *TFont) {
    LinkLabel_SetFont(l.instance, CheckPtr(value))
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (l *TLinkLabel) ParentColor() bool {
    return LinkLabel_GetParentColor(l.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (l *TLinkLabel) SetParentColor(value bool) {
    LinkLabel_SetParentColor(l.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (l *TLinkLabel) ParentFont() bool {
    return LinkLabel_GetParentFont(l.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (l *TLinkLabel) SetParentFont(value bool) {
    LinkLabel_SetParentFont(l.instance, value)
}

// ParentShowHint
func (l *TLinkLabel) ParentShowHint() bool {
    return LinkLabel_GetParentShowHint(l.instance)
}

// SetParentShowHint
func (l *TLinkLabel) SetParentShowHint(value bool) {
    LinkLabel_SetParentShowHint(l.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (l *TLinkLabel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(LinkLabel_GetPopupMenu(l.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (l *TLinkLabel) SetPopupMenu(value IComponent) {
    LinkLabel_SetPopupMenu(l.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (l *TLinkLabel) ShowHint() bool {
    return LinkLabel_GetShowHint(l.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (l *TLinkLabel) SetShowHint(value bool) {
    LinkLabel_SetShowHint(l.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (l *TLinkLabel) TabOrder() TTabOrder {
    return LinkLabel_GetTabOrder(l.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (l *TLinkLabel) SetTabOrder(value TTabOrder) {
    LinkLabel_SetTabOrder(l.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (l *TLinkLabel) TabStop() bool {
    return LinkLabel_GetTabStop(l.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (l *TLinkLabel) SetTabStop(value bool) {
    LinkLabel_SetTabStop(l.instance, value)
}

// UseVisualStyle
func (l *TLinkLabel) UseVisualStyle() bool {
    return LinkLabel_GetUseVisualStyle(l.instance)
}

// SetUseVisualStyle
func (l *TLinkLabel) SetUseVisualStyle(value bool) {
    LinkLabel_SetUseVisualStyle(l.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (l *TLinkLabel) Visible() bool {
    return LinkLabel_GetVisible(l.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (l *TLinkLabel) SetVisible(value bool) {
    LinkLabel_SetVisible(l.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (l *TLinkLabel) SetOnClick(fn TNotifyEvent) {
    LinkLabel_SetOnClick(l.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (l *TLinkLabel) SetOnContextPopup(fn TContextPopupEvent) {
    LinkLabel_SetOnContextPopup(l.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (l *TLinkLabel) SetOnDblClick(fn TNotifyEvent) {
    LinkLabel_SetOnDblClick(l.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (l *TLinkLabel) SetOnDragDrop(fn TDragDropEvent) {
    LinkLabel_SetOnDragDrop(l.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (l *TLinkLabel) SetOnDragOver(fn TDragOverEvent) {
    LinkLabel_SetOnDragOver(l.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (l *TLinkLabel) SetOnEndDock(fn TEndDragEvent) {
    LinkLabel_SetOnEndDock(l.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (l *TLinkLabel) SetOnEndDrag(fn TEndDragEvent) {
    LinkLabel_SetOnEndDrag(l.instance, fn)
}

// SetOnGesture
func (l *TLinkLabel) SetOnGesture(fn TGestureEvent) {
    LinkLabel_SetOnGesture(l.instance, fn)
}

// SetOnMouseActivate
func (l *TLinkLabel) SetOnMouseActivate(fn TMouseActivateEvent) {
    LinkLabel_SetOnMouseActivate(l.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (l *TLinkLabel) SetOnMouseDown(fn TMouseEvent) {
    LinkLabel_SetOnMouseDown(l.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (l *TLinkLabel) SetOnMouseEnter(fn TNotifyEvent) {
    LinkLabel_SetOnMouseEnter(l.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (l *TLinkLabel) SetOnMouseLeave(fn TNotifyEvent) {
    LinkLabel_SetOnMouseLeave(l.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (l *TLinkLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    LinkLabel_SetOnMouseMove(l.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (l *TLinkLabel) SetOnMouseUp(fn TMouseEvent) {
    LinkLabel_SetOnMouseUp(l.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (l *TLinkLabel) SetOnStartDock(fn TStartDockEvent) {
    LinkLabel_SetOnStartDock(l.instance, fn)
}

// SetOnLinkClick
func (l *TLinkLabel) SetOnLinkClick(fn TSysLinkEvent) {
    LinkLabel_SetOnLinkClick(l.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (l *TLinkLabel) DockClientCount() int32 {
    return LinkLabel_GetDockClientCount(l.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (l *TLinkLabel) DockSite() bool {
    return LinkLabel_GetDockSite(l.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (l *TLinkLabel) SetDockSite(value bool) {
    LinkLabel_SetDockSite(l.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (l *TLinkLabel) DoubleBuffered() bool {
    return LinkLabel_GetDoubleBuffered(l.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (l *TLinkLabel) SetDoubleBuffered(value bool) {
    LinkLabel_SetDoubleBuffered(l.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (l *TLinkLabel) AlignDisabled() bool {
    return LinkLabel_GetAlignDisabled(l.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (l *TLinkLabel) MouseInClient() bool {
    return LinkLabel_GetMouseInClient(l.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (l *TLinkLabel) VisibleDockClientCount() int32 {
    return LinkLabel_GetVisibleDockClientCount(l.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (l *TLinkLabel) Brush() *TBrush {
    return BrushFromInst(LinkLabel_GetBrush(l.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (l *TLinkLabel) ControlCount() int32 {
    return LinkLabel_GetControlCount(l.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (l *TLinkLabel) Handle() HWND {
    return LinkLabel_GetHandle(l.instance)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (l *TLinkLabel) ParentDoubleBuffered() bool {
    return LinkLabel_GetParentDoubleBuffered(l.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (l *TLinkLabel) SetParentDoubleBuffered(value bool) {
    LinkLabel_SetParentDoubleBuffered(l.instance, value)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (l *TLinkLabel) ParentWindow() HWND {
    return LinkLabel_GetParentWindow(l.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (l *TLinkLabel) SetParentWindow(value HWND) {
    LinkLabel_SetParentWindow(l.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (l *TLinkLabel) UseDockManager() bool {
    return LinkLabel_GetUseDockManager(l.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (l *TLinkLabel) SetUseDockManager(value bool) {
    LinkLabel_SetUseDockManager(l.instance, value)
}

// Action
func (l *TLinkLabel) Action() *TAction {
    return ActionFromInst(LinkLabel_GetAction(l.instance))
}

// SetAction
func (l *TLinkLabel) SetAction(value IComponent) {
    LinkLabel_SetAction(l.instance, CheckPtr(value))
}

// BiDiMode
func (l *TLinkLabel) BiDiMode() TBiDiMode {
    return LinkLabel_GetBiDiMode(l.instance)
}

// SetBiDiMode
func (l *TLinkLabel) SetBiDiMode(value TBiDiMode) {
    LinkLabel_SetBiDiMode(l.instance, value)
}

// BoundsRect
func (l *TLinkLabel) BoundsRect() TRect {
    return LinkLabel_GetBoundsRect(l.instance)
}

// SetBoundsRect
func (l *TLinkLabel) SetBoundsRect(value TRect) {
    LinkLabel_SetBoundsRect(l.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (l *TLinkLabel) ClientHeight() int32 {
    return LinkLabel_GetClientHeight(l.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (l *TLinkLabel) SetClientHeight(value int32) {
    LinkLabel_SetClientHeight(l.instance, value)
}

// ClientOrigin
func (l *TLinkLabel) ClientOrigin() TPoint {
    return LinkLabel_GetClientOrigin(l.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (l *TLinkLabel) ClientRect() TRect {
    return LinkLabel_GetClientRect(l.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (l *TLinkLabel) ClientWidth() int32 {
    return LinkLabel_GetClientWidth(l.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (l *TLinkLabel) SetClientWidth(value int32) {
    LinkLabel_SetClientWidth(l.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (l *TLinkLabel) ControlState() TControlState {
    return LinkLabel_GetControlState(l.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (l *TLinkLabel) SetControlState(value TControlState) {
    LinkLabel_SetControlState(l.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (l *TLinkLabel) ControlStyle() TControlStyle {
    return LinkLabel_GetControlStyle(l.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (l *TLinkLabel) SetControlStyle(value TControlStyle) {
    LinkLabel_SetControlStyle(l.instance, value)
}

// ExplicitLeft
func (l *TLinkLabel) ExplicitLeft() int32 {
    return LinkLabel_GetExplicitLeft(l.instance)
}

// ExplicitTop
func (l *TLinkLabel) ExplicitTop() int32 {
    return LinkLabel_GetExplicitTop(l.instance)
}

// ExplicitWidth
func (l *TLinkLabel) ExplicitWidth() int32 {
    return LinkLabel_GetExplicitWidth(l.instance)
}

// ExplicitHeight
func (l *TLinkLabel) ExplicitHeight() int32 {
    return LinkLabel_GetExplicitHeight(l.instance)
}

// Floating
func (l *TLinkLabel) Floating() bool {
    return LinkLabel_GetFloating(l.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (l *TLinkLabel) Parent() *TWinControl {
    return WinControlFromInst(LinkLabel_GetParent(l.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (l *TLinkLabel) SetParent(value IWinControl) {
    LinkLabel_SetParent(l.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (l *TLinkLabel) StyleElements() TStyleElements {
    return LinkLabel_GetStyleElements(l.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (l *TLinkLabel) SetStyleElements(value TStyleElements) {
    LinkLabel_SetStyleElements(l.instance, value)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (l *TLinkLabel) AlignWithMargins() bool {
    return LinkLabel_GetAlignWithMargins(l.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (l *TLinkLabel) SetAlignWithMargins(value bool) {
    LinkLabel_SetAlignWithMargins(l.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (l *TLinkLabel) Left() int32 {
    return LinkLabel_GetLeft(l.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (l *TLinkLabel) SetLeft(value int32) {
    LinkLabel_SetLeft(l.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (l *TLinkLabel) Top() int32 {
    return LinkLabel_GetTop(l.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (l *TLinkLabel) SetTop(value int32) {
    LinkLabel_SetTop(l.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (l *TLinkLabel) Width() int32 {
    return LinkLabel_GetWidth(l.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (l *TLinkLabel) SetWidth(value int32) {
    LinkLabel_SetWidth(l.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (l *TLinkLabel) Height() int32 {
    return LinkLabel_GetHeight(l.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (l *TLinkLabel) SetHeight(value int32) {
    LinkLabel_SetHeight(l.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (l *TLinkLabel) Cursor() TCursor {
    return LinkLabel_GetCursor(l.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (l *TLinkLabel) SetCursor(value TCursor) {
    LinkLabel_SetCursor(l.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (l *TLinkLabel) Hint() string {
    return LinkLabel_GetHint(l.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (l *TLinkLabel) SetHint(value string) {
    LinkLabel_SetHint(l.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (l *TLinkLabel) Margins() *TMargins {
    return MarginsFromInst(LinkLabel_GetMargins(l.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (l *TLinkLabel) SetMargins(value *TMargins) {
    LinkLabel_SetMargins(l.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (l *TLinkLabel) CustomHint() *TCustomHint {
    return CustomHintFromInst(LinkLabel_GetCustomHint(l.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (l *TLinkLabel) SetCustomHint(value IComponent) {
    LinkLabel_SetCustomHint(l.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (l *TLinkLabel) ComponentCount() int32 {
    return LinkLabel_GetComponentCount(l.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (l *TLinkLabel) ComponentIndex() int32 {
    return LinkLabel_GetComponentIndex(l.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (l *TLinkLabel) SetComponentIndex(value int32) {
    LinkLabel_SetComponentIndex(l.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TLinkLabel) Owner() *TComponent {
    return ComponentFromInst(LinkLabel_GetOwner(l.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (l *TLinkLabel) Name() string {
    return LinkLabel_GetName(l.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (l *TLinkLabel) SetName(value string) {
    LinkLabel_SetName(l.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TLinkLabel) Tag() int {
    return LinkLabel_GetTag(l.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TLinkLabel) SetTag(value int) {
    LinkLabel_SetTag(l.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (l *TLinkLabel) DockClients(Index int32) *TControl {
    return ControlFromInst(LinkLabel_GetDockClients(l.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (l *TLinkLabel) Controls(Index int32) *TControl {
    return ControlFromInst(LinkLabel_GetControls(l.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TLinkLabel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(LinkLabel_GetComponents(l.instance, AIndex))
}

