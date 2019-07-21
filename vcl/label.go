
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

type TLabel struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewLabel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewLabel(owner IComponent) *TLabel {
    l := new(TLabel)
    l.instance = Label_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// LabelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func LabelFromInst(inst uintptr) *TLabel {
    l := new(TLabel)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// LabelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func LabelFromObj(obj IObject) *TLabel {
    l := new(TLabel)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// LabelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func LabelFromUnsafePointer(ptr unsafe.Pointer) *TLabel {
    l := new(TLabel)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TLabel) Free() {
    if l.instance != 0 {
        Label_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TLabel) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TLabel) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TLabel) IsValid() bool {
    return l.instance != 0
}

// TLabelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TLabelClass() TClass {
    return Label_StaticClassType()
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (l *TLabel) BringToFront() {
    Label_BringToFront(l.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (l *TLabel) ClientToScreen(Point TPoint) TPoint {
    return Label_ClientToScreen(l.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (l *TLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Label_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (l *TLabel) Dragging() bool {
    return Label_Dragging(l.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (l *TLabel) HasParent() bool {
    return Label_HasParent(l.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (l *TLabel) Hide() {
    Label_Hide(l.instance)
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (l *TLabel) Invalidate() {
    Label_Invalidate(l.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (l *TLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Label_Perform(l.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (l *TLabel) Refresh() {
    Label_Refresh(l.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (l *TLabel) Repaint() {
    Label_Repaint(l.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (l *TLabel) ScreenToClient(Point TPoint) TPoint {
    return Label_ScreenToClient(l.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (l *TLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Label_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (l *TLabel) SendToBack() {
    Label_SendToBack(l.instance)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (l *TLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Label_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (l *TLabel) Show() {
    Label_Show(l.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (l *TLabel) Update() {
    Label_Update(l.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (l *TLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Label_GetTextBuf(l.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (l *TLabel) GetTextLen() int32 {
    return Label_GetTextLen(l.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (l *TLabel) SetTextBuf(Buffer string) {
    Label_SetTextBuf(l.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (l *TLabel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Label_FindComponent(l.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TLabel) GetNamePath() string {
    return Label_GetNamePath(l.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TLabel) Assign(Source IObject) {
    Label_Assign(l.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TLabel) DisposeOf() {
    Label_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TLabel) ClassType() TClass {
    return Label_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TLabel) ClassName() string {
    return Label_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TLabel) InstanceSize() int32 {
    return Label_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TLabel) InheritsFrom(AClass TClass) bool {
    return Label_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TLabel) Equals(Obj IObject) bool {
    return Label_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TLabel) GetHashCode() int32 {
    return Label_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TLabel) ToString() string {
    return Label_ToString(l.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (l *TLabel) Align() TAlign {
    return Label_GetAlign(l.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (l *TLabel) SetAlign(value TAlign) {
    Label_SetAlign(l.instance, value)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (l *TLabel) Alignment() TAlignment {
    return Label_GetAlignment(l.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (l *TLabel) SetAlignment(value TAlignment) {
    Label_SetAlignment(l.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (l *TLabel) Anchors() TAnchors {
    return Label_GetAnchors(l.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (l *TLabel) SetAnchors(value TAnchors) {
    Label_SetAnchors(l.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (l *TLabel) AutoSize() bool {
    return Label_GetAutoSize(l.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (l *TLabel) SetAutoSize(value bool) {
    Label_SetAutoSize(l.instance, value)
}

// BiDiMode
func (l *TLabel) BiDiMode() TBiDiMode {
    return Label_GetBiDiMode(l.instance)
}

// SetBiDiMode
func (l *TLabel) SetBiDiMode(value TBiDiMode) {
    Label_SetBiDiMode(l.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (l *TLabel) Caption() string {
    return Label_GetCaption(l.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (l *TLabel) SetCaption(value string) {
    Label_SetCaption(l.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (l *TLabel) Color() TColor {
    return Label_GetColor(l.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (l *TLabel) SetColor(value TColor) {
    Label_SetColor(l.instance, value)
}

// Constraints
func (l *TLabel) Constraints() *TSizeConstraints {
    return SizeConstraintsFromInst(Label_GetConstraints(l.instance))
}

// SetConstraints
func (l *TLabel) SetConstraints(value *TSizeConstraints) {
    Label_SetConstraints(l.instance, CheckPtr(value))
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (l *TLabel) DragCursor() TCursor {
    return Label_GetDragCursor(l.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (l *TLabel) SetDragCursor(value TCursor) {
    Label_SetDragCursor(l.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (l *TLabel) DragKind() TDragKind {
    return Label_GetDragKind(l.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (l *TLabel) SetDragKind(value TDragKind) {
    Label_SetDragKind(l.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (l *TLabel) DragMode() TDragMode {
    return Label_GetDragMode(l.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (l *TLabel) SetDragMode(value TDragMode) {
    Label_SetDragMode(l.instance, value)
}

// EllipsisPosition
func (l *TLabel) EllipsisPosition() TEllipsisPosition {
    return Label_GetEllipsisPosition(l.instance)
}

// SetEllipsisPosition
func (l *TLabel) SetEllipsisPosition(value TEllipsisPosition) {
    Label_SetEllipsisPosition(l.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (l *TLabel) Enabled() bool {
    return Label_GetEnabled(l.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (l *TLabel) SetEnabled(value bool) {
    Label_SetEnabled(l.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (l *TLabel) Font() *TFont {
    return FontFromInst(Label_GetFont(l.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (l *TLabel) SetFont(value *TFont) {
    Label_SetFont(l.instance, CheckPtr(value))
}

// GlowSize
func (l *TLabel) GlowSize() int32 {
    return Label_GetGlowSize(l.instance)
}

// SetGlowSize
func (l *TLabel) SetGlowSize(value int32) {
    Label_SetGlowSize(l.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (l *TLabel) ParentColor() bool {
    return Label_GetParentColor(l.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (l *TLabel) SetParentColor(value bool) {
    Label_SetParentColor(l.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (l *TLabel) ParentFont() bool {
    return Label_GetParentFont(l.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (l *TLabel) SetParentFont(value bool) {
    Label_SetParentFont(l.instance, value)
}

// ParentShowHint
func (l *TLabel) ParentShowHint() bool {
    return Label_GetParentShowHint(l.instance)
}

// SetParentShowHint
func (l *TLabel) SetParentShowHint(value bool) {
    Label_SetParentShowHint(l.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (l *TLabel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Label_GetPopupMenu(l.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (l *TLabel) SetPopupMenu(value IComponent) {
    Label_SetPopupMenu(l.instance, CheckPtr(value))
}

// ShowAccelChar
func (l *TLabel) ShowAccelChar() bool {
    return Label_GetShowAccelChar(l.instance)
}

// SetShowAccelChar
func (l *TLabel) SetShowAccelChar(value bool) {
    Label_SetShowAccelChar(l.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (l *TLabel) ShowHint() bool {
    return Label_GetShowHint(l.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (l *TLabel) SetShowHint(value bool) {
    Label_SetShowHint(l.instance, value)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (l *TLabel) Transparent() bool {
    return Label_GetTransparent(l.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (l *TLabel) SetTransparent(value bool) {
    Label_SetTransparent(l.instance, value)
}

// Layout
func (l *TLabel) Layout() TTextLayout {
    return Label_GetLayout(l.instance)
}

// SetLayout
func (l *TLabel) SetLayout(value TTextLayout) {
    Label_SetLayout(l.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (l *TLabel) Visible() bool {
    return Label_GetVisible(l.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (l *TLabel) SetVisible(value bool) {
    Label_SetVisible(l.instance, value)
}

// WordWrap
// CN: 获取自动换行。
// EN: Get Automatic line break.
func (l *TLabel) WordWrap() bool {
    return Label_GetWordWrap(l.instance)
}

// SetWordWrap
// CN: 设置自动换行。
// EN: Set Automatic line break.
func (l *TLabel) SetWordWrap(value bool) {
    Label_SetWordWrap(l.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (l *TLabel) StyleElements() TStyleElements {
    return Label_GetStyleElements(l.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (l *TLabel) SetStyleElements(value TStyleElements) {
    Label_SetStyleElements(l.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (l *TLabel) SetOnClick(fn TNotifyEvent) {
    Label_SetOnClick(l.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (l *TLabel) SetOnContextPopup(fn TContextPopupEvent) {
    Label_SetOnContextPopup(l.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (l *TLabel) SetOnDblClick(fn TNotifyEvent) {
    Label_SetOnDblClick(l.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (l *TLabel) SetOnDragDrop(fn TDragDropEvent) {
    Label_SetOnDragDrop(l.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (l *TLabel) SetOnDragOver(fn TDragOverEvent) {
    Label_SetOnDragOver(l.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (l *TLabel) SetOnEndDock(fn TEndDragEvent) {
    Label_SetOnEndDock(l.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (l *TLabel) SetOnEndDrag(fn TEndDragEvent) {
    Label_SetOnEndDrag(l.instance, fn)
}

// SetOnGesture
func (l *TLabel) SetOnGesture(fn TGestureEvent) {
    Label_SetOnGesture(l.instance, fn)
}

// SetOnMouseActivate
func (l *TLabel) SetOnMouseActivate(fn TMouseActivateEvent) {
    Label_SetOnMouseActivate(l.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (l *TLabel) SetOnMouseDown(fn TMouseEvent) {
    Label_SetOnMouseDown(l.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (l *TLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    Label_SetOnMouseMove(l.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (l *TLabel) SetOnMouseUp(fn TMouseEvent) {
    Label_SetOnMouseUp(l.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (l *TLabel) SetOnMouseEnter(fn TNotifyEvent) {
    Label_SetOnMouseEnter(l.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (l *TLabel) SetOnMouseLeave(fn TNotifyEvent) {
    Label_SetOnMouseLeave(l.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (l *TLabel) SetOnStartDock(fn TStartDockEvent) {
    Label_SetOnStartDock(l.instance, fn)
}

// Canvas
// CN: 获取画布。
// EN: .
func (l *TLabel) Canvas() *TCanvas {
    return CanvasFromInst(Label_GetCanvas(l.instance))
}

// Action
func (l *TLabel) Action() *TAction {
    return ActionFromInst(Label_GetAction(l.instance))
}

// SetAction
func (l *TLabel) SetAction(value IComponent) {
    Label_SetAction(l.instance, CheckPtr(value))
}

// BoundsRect
func (l *TLabel) BoundsRect() TRect {
    return Label_GetBoundsRect(l.instance)
}

// SetBoundsRect
func (l *TLabel) SetBoundsRect(value TRect) {
    Label_SetBoundsRect(l.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (l *TLabel) ClientHeight() int32 {
    return Label_GetClientHeight(l.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (l *TLabel) SetClientHeight(value int32) {
    Label_SetClientHeight(l.instance, value)
}

// ClientOrigin
func (l *TLabel) ClientOrigin() TPoint {
    return Label_GetClientOrigin(l.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (l *TLabel) ClientRect() TRect {
    return Label_GetClientRect(l.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (l *TLabel) ClientWidth() int32 {
    return Label_GetClientWidth(l.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (l *TLabel) SetClientWidth(value int32) {
    Label_SetClientWidth(l.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (l *TLabel) ControlState() TControlState {
    return Label_GetControlState(l.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (l *TLabel) SetControlState(value TControlState) {
    Label_SetControlState(l.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (l *TLabel) ControlStyle() TControlStyle {
    return Label_GetControlStyle(l.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (l *TLabel) SetControlStyle(value TControlStyle) {
    Label_SetControlStyle(l.instance, value)
}

// ExplicitLeft
func (l *TLabel) ExplicitLeft() int32 {
    return Label_GetExplicitLeft(l.instance)
}

// ExplicitTop
func (l *TLabel) ExplicitTop() int32 {
    return Label_GetExplicitTop(l.instance)
}

// ExplicitWidth
func (l *TLabel) ExplicitWidth() int32 {
    return Label_GetExplicitWidth(l.instance)
}

// ExplicitHeight
func (l *TLabel) ExplicitHeight() int32 {
    return Label_GetExplicitHeight(l.instance)
}

// Floating
func (l *TLabel) Floating() bool {
    return Label_GetFloating(l.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (l *TLabel) Parent() *TWinControl {
    return WinControlFromInst(Label_GetParent(l.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (l *TLabel) SetParent(value IWinControl) {
    Label_SetParent(l.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (l *TLabel) AlignWithMargins() bool {
    return Label_GetAlignWithMargins(l.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (l *TLabel) SetAlignWithMargins(value bool) {
    Label_SetAlignWithMargins(l.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (l *TLabel) Left() int32 {
    return Label_GetLeft(l.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (l *TLabel) SetLeft(value int32) {
    Label_SetLeft(l.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (l *TLabel) Top() int32 {
    return Label_GetTop(l.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (l *TLabel) SetTop(value int32) {
    Label_SetTop(l.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (l *TLabel) Width() int32 {
    return Label_GetWidth(l.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (l *TLabel) SetWidth(value int32) {
    Label_SetWidth(l.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (l *TLabel) Height() int32 {
    return Label_GetHeight(l.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (l *TLabel) SetHeight(value int32) {
    Label_SetHeight(l.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (l *TLabel) Cursor() TCursor {
    return Label_GetCursor(l.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (l *TLabel) SetCursor(value TCursor) {
    Label_SetCursor(l.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (l *TLabel) Hint() string {
    return Label_GetHint(l.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (l *TLabel) SetHint(value string) {
    Label_SetHint(l.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (l *TLabel) Margins() *TMargins {
    return MarginsFromInst(Label_GetMargins(l.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (l *TLabel) SetMargins(value *TMargins) {
    Label_SetMargins(l.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (l *TLabel) CustomHint() *TCustomHint {
    return CustomHintFromInst(Label_GetCustomHint(l.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (l *TLabel) SetCustomHint(value IComponent) {
    Label_SetCustomHint(l.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (l *TLabel) ComponentCount() int32 {
    return Label_GetComponentCount(l.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (l *TLabel) ComponentIndex() int32 {
    return Label_GetComponentIndex(l.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (l *TLabel) SetComponentIndex(value int32) {
    Label_SetComponentIndex(l.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TLabel) Owner() *TComponent {
    return ComponentFromInst(Label_GetOwner(l.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (l *TLabel) Name() string {
    return Label_GetName(l.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (l *TLabel) SetName(value string) {
    Label_SetName(l.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TLabel) Tag() int {
    return Label_GetTag(l.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TLabel) SetTag(value int) {
    Label_SetTag(l.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TLabel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Label_GetComponents(l.instance, AIndex))
}

