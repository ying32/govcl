
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

type TStaticText struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewStaticText
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewStaticText(owner IComponent) *TStaticText {
    s := new(TStaticText)
    s.instance = StaticText_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StaticTextFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func StaticTextFromInst(inst uintptr) *TStaticText {
    s := new(TStaticText)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// StaticTextFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func StaticTextFromObj(obj IObject) *TStaticText {
    s := new(TStaticText)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StaticTextFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func StaticTextFromUnsafePointer(ptr unsafe.Pointer) *TStaticText {
    s := new(TStaticText)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TStaticText) Free() {
    if s.instance != 0 {
        StaticText_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TStaticText) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TStaticText) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TStaticText) IsValid() bool {
    return s.instance != 0
}

// TStaticTextClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TStaticTextClass() TClass {
    return StaticText_StaticClassType()
}

// CanFocus
func (s *TStaticText) CanFocus() bool {
    return StaticText_CanFocus(s.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (s *TStaticText) ContainsControl(Control IControl) bool {
    return StaticText_ContainsControl(s.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (s *TStaticText) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(StaticText_ControlAtPos(s.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (s *TStaticText) DisableAlign() {
    StaticText_DisableAlign(s.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (s *TStaticText) EnableAlign() {
    StaticText_EnableAlign(s.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (s *TStaticText) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(StaticText_FindChildControl(s.instance, ControlName))
}

// FlipChildren
func (s *TStaticText) FlipChildren(AllLevels bool) {
    StaticText_FlipChildren(s.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (s *TStaticText) Focused() bool {
    return StaticText_Focused(s.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (s *TStaticText) HandleAllocated() bool {
    return StaticText_HandleAllocated(s.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (s *TStaticText) InsertControl(AControl IControl) {
    StaticText_InsertControl(s.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (s *TStaticText) Invalidate() {
    StaticText_Invalidate(s.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (s *TStaticText) PaintTo(DC HDC, X int32, Y int32) {
    StaticText_PaintTo(s.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (s *TStaticText) RemoveControl(AControl IControl) {
    StaticText_RemoveControl(s.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (s *TStaticText) Realign() {
    StaticText_Realign(s.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (s *TStaticText) Repaint() {
    StaticText_Repaint(s.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (s *TStaticText) ScaleBy(M int32, D int32) {
    StaticText_ScaleBy(s.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (s *TStaticText) ScrollBy(DeltaX int32, DeltaY int32) {
    StaticText_ScrollBy(s.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (s *TStaticText) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    StaticText_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (s *TStaticText) SetFocus() {
    StaticText_SetFocus(s.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (s *TStaticText) Update() {
    StaticText_Update(s.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (s *TStaticText) UpdateControlState() {
    StaticText_UpdateControlState(s.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (s *TStaticText) BringToFront() {
    StaticText_BringToFront(s.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (s *TStaticText) ClientToScreen(Point TPoint) TPoint {
    return StaticText_ClientToScreen(s.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (s *TStaticText) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return StaticText_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (s *TStaticText) Dragging() bool {
    return StaticText_Dragging(s.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TStaticText) HasParent() bool {
    return StaticText_HasParent(s.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (s *TStaticText) Hide() {
    StaticText_Hide(s.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (s *TStaticText) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return StaticText_Perform(s.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (s *TStaticText) Refresh() {
    StaticText_Refresh(s.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (s *TStaticText) ScreenToClient(Point TPoint) TPoint {
    return StaticText_ScreenToClient(s.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (s *TStaticText) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return StaticText_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (s *TStaticText) SendToBack() {
    StaticText_SendToBack(s.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (s *TStaticText) Show() {
    StaticText_Show(s.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (s *TStaticText) GetTextBuf(Buffer string, BufSize int32) int32 {
    return StaticText_GetTextBuf(s.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (s *TStaticText) GetTextLen() int32 {
    return StaticText_GetTextLen(s.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (s *TStaticText) SetTextBuf(Buffer string) {
    StaticText_SetTextBuf(s.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TStaticText) FindComponent(AName string) *TComponent {
    return ComponentFromInst(StaticText_FindComponent(s.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TStaticText) GetNamePath() string {
    return StaticText_GetNamePath(s.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TStaticText) Assign(Source IObject) {
    StaticText_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TStaticText) DisposeOf() {
    StaticText_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TStaticText) ClassType() TClass {
    return StaticText_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TStaticText) ClassName() string {
    return StaticText_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TStaticText) InstanceSize() int32 {
    return StaticText_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TStaticText) InheritsFrom(AClass TClass) bool {
    return StaticText_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TStaticText) Equals(Obj IObject) bool {
    return StaticText_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TStaticText) GetHashCode() int32 {
    return StaticText_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TStaticText) ToString() string {
    return StaticText_ToString(s.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (s *TStaticText) Align() TAlign {
    return StaticText_GetAlign(s.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (s *TStaticText) SetAlign(value TAlign) {
    StaticText_SetAlign(s.instance, value)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (s *TStaticText) Alignment() TAlignment {
    return StaticText_GetAlignment(s.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (s *TStaticText) SetAlignment(value TAlignment) {
    StaticText_SetAlignment(s.instance, value)
}

// Anchors
func (s *TStaticText) Anchors() TAnchors {
    return StaticText_GetAnchors(s.instance)
}

// SetAnchors
func (s *TStaticText) SetAnchors(value TAnchors) {
    StaticText_SetAnchors(s.instance, value)
}

// AutoSize
func (s *TStaticText) AutoSize() bool {
    return StaticText_GetAutoSize(s.instance)
}

// SetAutoSize
func (s *TStaticText) SetAutoSize(value bool) {
    StaticText_SetAutoSize(s.instance, value)
}

// BevelEdges
func (s *TStaticText) BevelEdges() TBevelEdges {
    return StaticText_GetBevelEdges(s.instance)
}

// SetBevelEdges
func (s *TStaticText) SetBevelEdges(value TBevelEdges) {
    StaticText_SetBevelEdges(s.instance, value)
}

// BevelInner
func (s *TStaticText) BevelInner() TBevelCut {
    return StaticText_GetBevelInner(s.instance)
}

// SetBevelInner
func (s *TStaticText) SetBevelInner(value TBevelCut) {
    StaticText_SetBevelInner(s.instance, value)
}

// BevelKind
func (s *TStaticText) BevelKind() TBevelKind {
    return StaticText_GetBevelKind(s.instance)
}

// SetBevelKind
func (s *TStaticText) SetBevelKind(value TBevelKind) {
    StaticText_SetBevelKind(s.instance, value)
}

// BevelOuter
func (s *TStaticText) BevelOuter() TBevelCut {
    return StaticText_GetBevelOuter(s.instance)
}

// SetBevelOuter
func (s *TStaticText) SetBevelOuter(value TBevelCut) {
    StaticText_SetBevelOuter(s.instance, value)
}

// BiDiMode
func (s *TStaticText) BiDiMode() TBiDiMode {
    return StaticText_GetBiDiMode(s.instance)
}

// SetBiDiMode
func (s *TStaticText) SetBiDiMode(value TBiDiMode) {
    StaticText_SetBiDiMode(s.instance, value)
}

// BorderStyle
func (s *TStaticText) BorderStyle() TStaticBorderStyle {
    return StaticText_GetBorderStyle(s.instance)
}

// SetBorderStyle
func (s *TStaticText) SetBorderStyle(value TStaticBorderStyle) {
    StaticText_SetBorderStyle(s.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (s *TStaticText) Caption() string {
    return StaticText_GetCaption(s.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (s *TStaticText) SetCaption(value string) {
    StaticText_SetCaption(s.instance, value)
}

// Color
// CN: 获取设置颜色。
// EN: Get Set color.
func (s *TStaticText) Color() TColor {
    return StaticText_GetColor(s.instance)
}

// SetColor
// CN: 设置设置颜色。
// EN: Set Set color.
func (s *TStaticText) SetColor(value TColor) {
    StaticText_SetColor(s.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (s *TStaticText) DoubleBuffered() bool {
    return StaticText_GetDoubleBuffered(s.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (s *TStaticText) SetDoubleBuffered(value bool) {
    StaticText_SetDoubleBuffered(s.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (s *TStaticText) DragCursor() TCursor {
    return StaticText_GetDragCursor(s.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (s *TStaticText) SetDragCursor(value TCursor) {
    StaticText_SetDragCursor(s.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (s *TStaticText) DragKind() TDragKind {
    return StaticText_GetDragKind(s.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (s *TStaticText) SetDragKind(value TDragKind) {
    StaticText_SetDragKind(s.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (s *TStaticText) DragMode() TDragMode {
    return StaticText_GetDragMode(s.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (s *TStaticText) SetDragMode(value TDragMode) {
    StaticText_SetDragMode(s.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TStaticText) Enabled() bool {
    return StaticText_GetEnabled(s.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TStaticText) SetEnabled(value bool) {
    StaticText_SetEnabled(s.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (s *TStaticText) Font() *TFont {
    return FontFromInst(StaticText_GetFont(s.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (s *TStaticText) SetFont(value *TFont) {
    StaticText_SetFont(s.instance, CheckPtr(value))
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (s *TStaticText) ParentColor() bool {
    return StaticText_GetParentColor(s.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (s *TStaticText) SetParentColor(value bool) {
    StaticText_SetParentColor(s.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (s *TStaticText) ParentDoubleBuffered() bool {
    return StaticText_GetParentDoubleBuffered(s.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (s *TStaticText) SetParentDoubleBuffered(value bool) {
    StaticText_SetParentDoubleBuffered(s.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (s *TStaticText) ParentFont() bool {
    return StaticText_GetParentFont(s.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (s *TStaticText) SetParentFont(value bool) {
    StaticText_SetParentFont(s.instance, value)
}

// ParentShowHint
func (s *TStaticText) ParentShowHint() bool {
    return StaticText_GetParentShowHint(s.instance)
}

// SetParentShowHint
func (s *TStaticText) SetParentShowHint(value bool) {
    StaticText_SetParentShowHint(s.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (s *TStaticText) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(StaticText_GetPopupMenu(s.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (s *TStaticText) SetPopupMenu(value IComponent) {
    StaticText_SetPopupMenu(s.instance, CheckPtr(value))
}

// ShowAccelChar
func (s *TStaticText) ShowAccelChar() bool {
    return StaticText_GetShowAccelChar(s.instance)
}

// SetShowAccelChar
func (s *TStaticText) SetShowAccelChar(value bool) {
    StaticText_SetShowAccelChar(s.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (s *TStaticText) ShowHint() bool {
    return StaticText_GetShowHint(s.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (s *TStaticText) SetShowHint(value bool) {
    StaticText_SetShowHint(s.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (s *TStaticText) TabOrder() TTabOrder {
    return StaticText_GetTabOrder(s.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (s *TStaticText) SetTabOrder(value TTabOrder) {
    StaticText_SetTabOrder(s.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (s *TStaticText) TabStop() bool {
    return StaticText_GetTabStop(s.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (s *TStaticText) SetTabStop(value bool) {
    StaticText_SetTabStop(s.instance, value)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (s *TStaticText) Transparent() bool {
    return StaticText_GetTransparent(s.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (s *TStaticText) SetTransparent(value bool) {
    StaticText_SetTransparent(s.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TStaticText) Visible() bool {
    return StaticText_GetVisible(s.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TStaticText) SetVisible(value bool) {
    StaticText_SetVisible(s.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (s *TStaticText) StyleElements() TStyleElements {
    return StaticText_GetStyleElements(s.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (s *TStaticText) SetStyleElements(value TStyleElements) {
    StaticText_SetStyleElements(s.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (s *TStaticText) SetOnClick(fn TNotifyEvent) {
    StaticText_SetOnClick(s.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (s *TStaticText) SetOnContextPopup(fn TContextPopupEvent) {
    StaticText_SetOnContextPopup(s.instance, fn)
}

// SetOnDblClick
func (s *TStaticText) SetOnDblClick(fn TNotifyEvent) {
    StaticText_SetOnDblClick(s.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (s *TStaticText) SetOnDragDrop(fn TDragDropEvent) {
    StaticText_SetOnDragDrop(s.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (s *TStaticText) SetOnDragOver(fn TDragOverEvent) {
    StaticText_SetOnDragOver(s.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (s *TStaticText) SetOnEndDock(fn TEndDragEvent) {
    StaticText_SetOnEndDock(s.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (s *TStaticText) SetOnEndDrag(fn TEndDragEvent) {
    StaticText_SetOnEndDrag(s.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (s *TStaticText) SetOnMouseDown(fn TMouseEvent) {
    StaticText_SetOnMouseDown(s.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (s *TStaticText) SetOnMouseEnter(fn TNotifyEvent) {
    StaticText_SetOnMouseEnter(s.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (s *TStaticText) SetOnMouseLeave(fn TNotifyEvent) {
    StaticText_SetOnMouseLeave(s.instance, fn)
}

// SetOnMouseMove
func (s *TStaticText) SetOnMouseMove(fn TMouseMoveEvent) {
    StaticText_SetOnMouseMove(s.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (s *TStaticText) SetOnMouseUp(fn TMouseEvent) {
    StaticText_SetOnMouseUp(s.instance, fn)
}

// SetOnStartDock
func (s *TStaticText) SetOnStartDock(fn TStartDockEvent) {
    StaticText_SetOnStartDock(s.instance, fn)
}

// DockClientCount
func (s *TStaticText) DockClientCount() int32 {
    return StaticText_GetDockClientCount(s.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (s *TStaticText) DockSite() bool {
    return StaticText_GetDockSite(s.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (s *TStaticText) SetDockSite(value bool) {
    StaticText_SetDockSite(s.instance, value)
}

// AlignDisabled
func (s *TStaticText) AlignDisabled() bool {
    return StaticText_GetAlignDisabled(s.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (s *TStaticText) MouseInClient() bool {
    return StaticText_GetMouseInClient(s.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (s *TStaticText) VisibleDockClientCount() int32 {
    return StaticText_GetVisibleDockClientCount(s.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (s *TStaticText) Brush() *TBrush {
    return BrushFromInst(StaticText_GetBrush(s.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (s *TStaticText) ControlCount() int32 {
    return StaticText_GetControlCount(s.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (s *TStaticText) Handle() HWND {
    return StaticText_GetHandle(s.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (s *TStaticText) ParentWindow() HWND {
    return StaticText_GetParentWindow(s.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (s *TStaticText) SetParentWindow(value HWND) {
    StaticText_SetParentWindow(s.instance, value)
}

// UseDockManager
func (s *TStaticText) UseDockManager() bool {
    return StaticText_GetUseDockManager(s.instance)
}

// SetUseDockManager
func (s *TStaticText) SetUseDockManager(value bool) {
    StaticText_SetUseDockManager(s.instance, value)
}

// Action
func (s *TStaticText) Action() *TAction {
    return ActionFromInst(StaticText_GetAction(s.instance))
}

// SetAction
func (s *TStaticText) SetAction(value IComponent) {
    StaticText_SetAction(s.instance, CheckPtr(value))
}

// BoundsRect
func (s *TStaticText) BoundsRect() TRect {
    return StaticText_GetBoundsRect(s.instance)
}

// SetBoundsRect
func (s *TStaticText) SetBoundsRect(value TRect) {
    StaticText_SetBoundsRect(s.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (s *TStaticText) ClientHeight() int32 {
    return StaticText_GetClientHeight(s.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (s *TStaticText) SetClientHeight(value int32) {
    StaticText_SetClientHeight(s.instance, value)
}

// ClientOrigin
func (s *TStaticText) ClientOrigin() TPoint {
    return StaticText_GetClientOrigin(s.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (s *TStaticText) ClientRect() TRect {
    return StaticText_GetClientRect(s.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (s *TStaticText) ClientWidth() int32 {
    return StaticText_GetClientWidth(s.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (s *TStaticText) SetClientWidth(value int32) {
    StaticText_SetClientWidth(s.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (s *TStaticText) ControlState() TControlState {
    return StaticText_GetControlState(s.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (s *TStaticText) SetControlState(value TControlState) {
    StaticText_SetControlState(s.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (s *TStaticText) ControlStyle() TControlStyle {
    return StaticText_GetControlStyle(s.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (s *TStaticText) SetControlStyle(value TControlStyle) {
    StaticText_SetControlStyle(s.instance, value)
}

// ExplicitLeft
func (s *TStaticText) ExplicitLeft() int32 {
    return StaticText_GetExplicitLeft(s.instance)
}

// ExplicitTop
func (s *TStaticText) ExplicitTop() int32 {
    return StaticText_GetExplicitTop(s.instance)
}

// ExplicitWidth
func (s *TStaticText) ExplicitWidth() int32 {
    return StaticText_GetExplicitWidth(s.instance)
}

// ExplicitHeight
func (s *TStaticText) ExplicitHeight() int32 {
    return StaticText_GetExplicitHeight(s.instance)
}

// Floating
func (s *TStaticText) Floating() bool {
    return StaticText_GetFloating(s.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TStaticText) Parent() *TWinControl {
    return WinControlFromInst(StaticText_GetParent(s.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TStaticText) SetParent(value IWinControl) {
    StaticText_SetParent(s.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (s *TStaticText) AlignWithMargins() bool {
    return StaticText_GetAlignWithMargins(s.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (s *TStaticText) SetAlignWithMargins(value bool) {
    StaticText_SetAlignWithMargins(s.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (s *TStaticText) Left() int32 {
    return StaticText_GetLeft(s.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (s *TStaticText) SetLeft(value int32) {
    StaticText_SetLeft(s.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (s *TStaticText) Top() int32 {
    return StaticText_GetTop(s.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (s *TStaticText) SetTop(value int32) {
    StaticText_SetTop(s.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (s *TStaticText) Width() int32 {
    return StaticText_GetWidth(s.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (s *TStaticText) SetWidth(value int32) {
    StaticText_SetWidth(s.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (s *TStaticText) Height() int32 {
    return StaticText_GetHeight(s.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (s *TStaticText) SetHeight(value int32) {
    StaticText_SetHeight(s.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TStaticText) Cursor() TCursor {
    return StaticText_GetCursor(s.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TStaticText) SetCursor(value TCursor) {
    StaticText_SetCursor(s.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (s *TStaticText) Hint() string {
    return StaticText_GetHint(s.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (s *TStaticText) SetHint(value string) {
    StaticText_SetHint(s.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (s *TStaticText) Margins() *TMargins {
    return MarginsFromInst(StaticText_GetMargins(s.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (s *TStaticText) SetMargins(value *TMargins) {
    StaticText_SetMargins(s.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (s *TStaticText) CustomHint() *TCustomHint {
    return CustomHintFromInst(StaticText_GetCustomHint(s.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (s *TStaticText) SetCustomHint(value IComponent) {
    StaticText_SetCustomHint(s.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TStaticText) ComponentCount() int32 {
    return StaticText_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TStaticText) ComponentIndex() int32 {
    return StaticText_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TStaticText) SetComponentIndex(value int32) {
    StaticText_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TStaticText) Owner() *TComponent {
    return ComponentFromInst(StaticText_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TStaticText) Name() string {
    return StaticText_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TStaticText) SetName(value string) {
    StaticText_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TStaticText) Tag() int {
    return StaticText_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TStaticText) SetTag(value int) {
    StaticText_SetTag(s.instance, value)
}

// DockClients
func (s *TStaticText) DockClients(Index int32) *TControl {
    return ControlFromInst(StaticText_GetDockClients(s.instance, Index))
}

// Controls
func (s *TStaticText) Controls(Index int32) *TControl {
    return ControlFromInst(StaticText_GetControls(s.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TStaticText) Components(AIndex int32) *TComponent {
    return ComponentFromInst(StaticText_GetComponents(s.instance, AIndex))
}

