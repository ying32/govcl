
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

type TColorBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewColorBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewColorBox(owner IComponent) *TColorBox {
    c := new(TColorBox)
    c.instance = ColorBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ColorBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ColorBoxFromInst(inst uintptr) *TColorBox {
    c := new(TColorBox)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ColorBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ColorBoxFromObj(obj IObject) *TColorBox {
    c := new(TColorBox)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ColorBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ColorBoxFromUnsafePointer(ptr unsafe.Pointer) *TColorBox {
    c := new(TColorBox)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TColorBox) Free() {
    if c.instance != 0 {
        ColorBox_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TColorBox) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TColorBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TColorBox) IsValid() bool {
    return c.instance != 0
}

// TColorBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TColorBoxClass() TClass {
    return ColorBox_StaticClassType()
}

// AddItem
func (c *TColorBox) AddItem(Item string, AObject IObject) {
    ColorBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

// Clear
// CN: 清除。
// EN: .
func (c *TColorBox) Clear() {
    ColorBox_Clear(c.instance)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (c *TColorBox) ClearSelection() {
    ColorBox_ClearSelection(c.instance)
}

// DeleteSelected
func (c *TColorBox) DeleteSelected() {
    ColorBox_DeleteSelected(c.instance)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (c *TColorBox) Focused() bool {
    return ColorBox_Focused(c.instance)
}

// SelectAll
// CN: 全选。
// EN: .
func (c *TColorBox) SelectAll() {
    ColorBox_SelectAll(c.instance)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (c *TColorBox) CanFocus() bool {
    return ColorBox_CanFocus(c.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (c *TColorBox) ContainsControl(Control IControl) bool {
    return ColorBox_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (c *TColorBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(ColorBox_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (c *TColorBox) DisableAlign() {
    ColorBox_DisableAlign(c.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (c *TColorBox) EnableAlign() {
    ColorBox_EnableAlign(c.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (c *TColorBox) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(ColorBox_FindChildControl(c.instance, ControlName))
}

// FlipChildren
func (c *TColorBox) FlipChildren(AllLevels bool) {
    ColorBox_FlipChildren(c.instance, AllLevels)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TColorBox) HandleAllocated() bool {
    return ColorBox_HandleAllocated(c.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (c *TColorBox) InsertControl(AControl IControl) {
    ColorBox_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (c *TColorBox) Invalidate() {
    ColorBox_Invalidate(c.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (c *TColorBox) PaintTo(DC HDC, X int32, Y int32) {
    ColorBox_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (c *TColorBox) RemoveControl(AControl IControl) {
    ColorBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (c *TColorBox) Realign() {
    ColorBox_Realign(c.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (c *TColorBox) Repaint() {
    ColorBox_Repaint(c.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (c *TColorBox) ScaleBy(M int32, D int32) {
    ColorBox_ScaleBy(c.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (c *TColorBox) ScrollBy(DeltaX int32, DeltaY int32) {
    ColorBox_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (c *TColorBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ColorBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (c *TColorBox) SetFocus() {
    ColorBox_SetFocus(c.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (c *TColorBox) Update() {
    ColorBox_Update(c.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (c *TColorBox) UpdateControlState() {
    ColorBox_UpdateControlState(c.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (c *TColorBox) BringToFront() {
    ColorBox_BringToFront(c.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (c *TColorBox) ClientToScreen(Point TPoint) TPoint {
    return ColorBox_ClientToScreen(c.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (c *TColorBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ColorBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (c *TColorBox) Dragging() bool {
    return ColorBox_Dragging(c.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TColorBox) HasParent() bool {
    return ColorBox_HasParent(c.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (c *TColorBox) Hide() {
    ColorBox_Hide(c.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (c *TColorBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ColorBox_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (c *TColorBox) Refresh() {
    ColorBox_Refresh(c.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (c *TColorBox) ScreenToClient(Point TPoint) TPoint {
    return ColorBox_ScreenToClient(c.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (c *TColorBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ColorBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (c *TColorBox) SendToBack() {
    ColorBox_SendToBack(c.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (c *TColorBox) Show() {
    ColorBox_Show(c.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TColorBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ColorBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (c *TColorBox) GetTextLen() int32 {
    return ColorBox_GetTextLen(c.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TColorBox) SetTextBuf(Buffer string) {
    ColorBox_SetTextBuf(c.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TColorBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ColorBox_FindComponent(c.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TColorBox) GetNamePath() string {
    return ColorBox_GetNamePath(c.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TColorBox) Assign(Source IObject) {
    ColorBox_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TColorBox) DisposeOf() {
    ColorBox_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TColorBox) ClassType() TClass {
    return ColorBox_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TColorBox) ClassName() string {
    return ColorBox_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TColorBox) InstanceSize() int32 {
    return ColorBox_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TColorBox) InheritsFrom(AClass TClass) bool {
    return ColorBox_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TColorBox) Equals(Obj IObject) bool {
    return ColorBox_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TColorBox) GetHashCode() int32 {
    return ColorBox_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TColorBox) ToString() string {
    return ColorBox_ToString(c.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (c *TColorBox) Align() TAlign {
    return ColorBox_GetAlign(c.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (c *TColorBox) SetAlign(value TAlign) {
    ColorBox_SetAlign(c.instance, value)
}

// AutoComplete
func (c *TColorBox) AutoComplete() bool {
    return ColorBox_GetAutoComplete(c.instance)
}

// SetAutoComplete
func (c *TColorBox) SetAutoComplete(value bool) {
    ColorBox_SetAutoComplete(c.instance, value)
}

// AutoDropDown
func (c *TColorBox) AutoDropDown() bool {
    return ColorBox_GetAutoDropDown(c.instance)
}

// SetAutoDropDown
func (c *TColorBox) SetAutoDropDown(value bool) {
    ColorBox_SetAutoDropDown(c.instance, value)
}

// DefaultColorColor
func (c *TColorBox) DefaultColorColor() TColor {
    return ColorBox_GetDefaultColorColor(c.instance)
}

// SetDefaultColorColor
func (c *TColorBox) SetDefaultColorColor(value TColor) {
    ColorBox_SetDefaultColorColor(c.instance, value)
}

// NoneColorColor
func (c *TColorBox) NoneColorColor() TColor {
    return ColorBox_GetNoneColorColor(c.instance)
}

// SetNoneColorColor
func (c *TColorBox) SetNoneColorColor(value TColor) {
    ColorBox_SetNoneColorColor(c.instance, value)
}

// Selected
func (c *TColorBox) Selected() TColor {
    return ColorBox_GetSelected(c.instance)
}

// SetSelected
func (c *TColorBox) SetSelected(value TColor) {
    ColorBox_SetSelected(c.instance, value)
}

// Style
func (c *TColorBox) Style() TColorBoxStyle {
    return ColorBox_GetStyle(c.instance)
}

// SetStyle
func (c *TColorBox) SetStyle(value TColorBoxStyle) {
    ColorBox_SetStyle(c.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (c *TColorBox) Anchors() TAnchors {
    return ColorBox_GetAnchors(c.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (c *TColorBox) SetAnchors(value TAnchors) {
    ColorBox_SetAnchors(c.instance, value)
}

// BevelEdges
func (c *TColorBox) BevelEdges() TBevelEdges {
    return ColorBox_GetBevelEdges(c.instance)
}

// SetBevelEdges
func (c *TColorBox) SetBevelEdges(value TBevelEdges) {
    ColorBox_SetBevelEdges(c.instance, value)
}

// BevelInner
func (c *TColorBox) BevelInner() TBevelCut {
    return ColorBox_GetBevelInner(c.instance)
}

// SetBevelInner
func (c *TColorBox) SetBevelInner(value TBevelCut) {
    ColorBox_SetBevelInner(c.instance, value)
}

// BevelKind
func (c *TColorBox) BevelKind() TBevelKind {
    return ColorBox_GetBevelKind(c.instance)
}

// SetBevelKind
func (c *TColorBox) SetBevelKind(value TBevelKind) {
    ColorBox_SetBevelKind(c.instance, value)
}

// BevelOuter
func (c *TColorBox) BevelOuter() TBevelCut {
    return ColorBox_GetBevelOuter(c.instance)
}

// SetBevelOuter
func (c *TColorBox) SetBevelOuter(value TBevelCut) {
    ColorBox_SetBevelOuter(c.instance, value)
}

// BiDiMode
func (c *TColorBox) BiDiMode() TBiDiMode {
    return ColorBox_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TColorBox) SetBiDiMode(value TBiDiMode) {
    ColorBox_SetBiDiMode(c.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (c *TColorBox) Color() TColor {
    return ColorBox_GetColor(c.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (c *TColorBox) SetColor(value TColor) {
    ColorBox_SetColor(c.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (c *TColorBox) DoubleBuffered() bool {
    return ColorBox_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (c *TColorBox) SetDoubleBuffered(value bool) {
    ColorBox_SetDoubleBuffered(c.instance, value)
}

// DropDownCount
func (c *TColorBox) DropDownCount() int32 {
    return ColorBox_GetDropDownCount(c.instance)
}

// SetDropDownCount
func (c *TColorBox) SetDropDownCount(value int32) {
    ColorBox_SetDropDownCount(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TColorBox) Enabled() bool {
    return ColorBox_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TColorBox) SetEnabled(value bool) {
    ColorBox_SetEnabled(c.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (c *TColorBox) Font() *TFont {
    return FontFromInst(ColorBox_GetFont(c.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (c *TColorBox) SetFont(value *TFont) {
    ColorBox_SetFont(c.instance, CheckPtr(value))
}

// ItemHeight
func (c *TColorBox) ItemHeight() int32 {
    return ColorBox_GetItemHeight(c.instance)
}

// SetItemHeight
func (c *TColorBox) SetItemHeight(value int32) {
    ColorBox_SetItemHeight(c.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (c *TColorBox) ParentColor() bool {
    return ColorBox_GetParentColor(c.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (c *TColorBox) SetParentColor(value bool) {
    ColorBox_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TColorBox) ParentCtl3D() bool {
    return ColorBox_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TColorBox) SetParentCtl3D(value bool) {
    ColorBox_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (c *TColorBox) ParentDoubleBuffered() bool {
    return ColorBox_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (c *TColorBox) SetParentDoubleBuffered(value bool) {
    ColorBox_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (c *TColorBox) ParentFont() bool {
    return ColorBox_GetParentFont(c.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (c *TColorBox) SetParentFont(value bool) {
    ColorBox_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TColorBox) ParentShowHint() bool {
    return ColorBox_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TColorBox) SetParentShowHint(value bool) {
    ColorBox_SetParentShowHint(c.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (c *TColorBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ColorBox_GetPopupMenu(c.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (c *TColorBox) SetPopupMenu(value IComponent) {
    ColorBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (c *TColorBox) ShowHint() bool {
    return ColorBox_GetShowHint(c.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (c *TColorBox) SetShowHint(value bool) {
    ColorBox_SetShowHint(c.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (c *TColorBox) TabOrder() TTabOrder {
    return ColorBox_GetTabOrder(c.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (c *TColorBox) SetTabOrder(value TTabOrder) {
    ColorBox_SetTabOrder(c.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (c *TColorBox) TabStop() bool {
    return ColorBox_GetTabStop(c.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (c *TColorBox) SetTabStop(value bool) {
    ColorBox_SetTabStop(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TColorBox) Visible() bool {
    return ColorBox_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TColorBox) SetVisible(value bool) {
    ColorBox_SetVisible(c.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (c *TColorBox) StyleElements() TStyleElements {
    return ColorBox_GetStyleElements(c.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (c *TColorBox) SetStyleElements(value TStyleElements) {
    ColorBox_SetStyleElements(c.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (c *TColorBox) SetOnChange(fn TNotifyEvent) {
    ColorBox_SetOnChange(c.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TColorBox) SetOnClick(fn TNotifyEvent) {
    ColorBox_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (c *TColorBox) SetOnContextPopup(fn TContextPopupEvent) {
    ColorBox_SetOnContextPopup(c.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (c *TColorBox) SetOnDragDrop(fn TDragDropEvent) {
    ColorBox_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (c *TColorBox) SetOnDragOver(fn TDragOverEvent) {
    ColorBox_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (c *TColorBox) SetOnEndDock(fn TEndDragEvent) {
    ColorBox_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (c *TColorBox) SetOnEndDrag(fn TEndDragEvent) {
    ColorBox_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (c *TColorBox) SetOnEnter(fn TNotifyEvent) {
    ColorBox_SetOnEnter(c.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (c *TColorBox) SetOnExit(fn TNotifyEvent) {
    ColorBox_SetOnExit(c.instance, fn)
}

// SetOnGesture
func (c *TColorBox) SetOnGesture(fn TGestureEvent) {
    ColorBox_SetOnGesture(c.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (c *TColorBox) SetOnKeyDown(fn TKeyEvent) {
    ColorBox_SetOnKeyDown(c.instance, fn)
}

// SetOnKeyPress
func (c *TColorBox) SetOnKeyPress(fn TKeyPressEvent) {
    ColorBox_SetOnKeyPress(c.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (c *TColorBox) SetOnKeyUp(fn TKeyEvent) {
    ColorBox_SetOnKeyUp(c.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (c *TColorBox) SetOnMouseEnter(fn TNotifyEvent) {
    ColorBox_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (c *TColorBox) SetOnMouseLeave(fn TNotifyEvent) {
    ColorBox_SetOnMouseLeave(c.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (c *TColorBox) SetOnStartDock(fn TStartDockEvent) {
    ColorBox_SetOnStartDock(c.instance, fn)
}

// AutoCompleteDelay
func (c *TColorBox) AutoCompleteDelay() uint32 {
    return ColorBox_GetAutoCompleteDelay(c.instance)
}

// SetAutoCompleteDelay
func (c *TColorBox) SetAutoCompleteDelay(value uint32) {
    ColorBox_SetAutoCompleteDelay(c.instance, value)
}

// AutoCloseUp
func (c *TColorBox) AutoCloseUp() bool {
    return ColorBox_GetAutoCloseUp(c.instance)
}

// SetAutoCloseUp
func (c *TColorBox) SetAutoCloseUp(value bool) {
    ColorBox_SetAutoCloseUp(c.instance, value)
}

// CharCase
func (c *TColorBox) CharCase() TEditCharCase {
    return ColorBox_GetCharCase(c.instance)
}

// SetCharCase
func (c *TColorBox) SetCharCase(value TEditCharCase) {
    ColorBox_SetCharCase(c.instance, value)
}

// SelText
// CN: 获取选择的文本。
// EN: .
func (c *TColorBox) SelText() string {
    return ColorBox_GetSelText(c.instance)
}

// SetSelText
// CN: 设置选择的文本。
// EN: .
func (c *TColorBox) SetSelText(value string) {
    ColorBox_SetSelText(c.instance, value)
}

// TextHint
// CN: 获取提示文本。
// EN: .
func (c *TColorBox) TextHint() string {
    return ColorBox_GetTextHint(c.instance)
}

// SetTextHint
// CN: 设置提示文本。
// EN: .
func (c *TColorBox) SetTextHint(value string) {
    ColorBox_SetTextHint(c.instance, value)
}

// Canvas
// CN: 获取画布。
// EN: .
func (c *TColorBox) Canvas() *TCanvas {
    return CanvasFromInst(ColorBox_GetCanvas(c.instance))
}

// DroppedDown
func (c *TColorBox) DroppedDown() bool {
    return ColorBox_GetDroppedDown(c.instance)
}

// SetDroppedDown
func (c *TColorBox) SetDroppedDown(value bool) {
    ColorBox_SetDroppedDown(c.instance, value)
}

// Items
func (c *TColorBox) Items() *TStrings {
    return StringsFromInst(ColorBox_GetItems(c.instance))
}

// SetItems
func (c *TColorBox) SetItems(value IObject) {
    ColorBox_SetItems(c.instance, CheckPtr(value))
}

// SelLength
// CN: 获取选择的长度。
// EN: .
func (c *TColorBox) SelLength() int32 {
    return ColorBox_GetSelLength(c.instance)
}

// SetSelLength
// CN: 设置选择的长度。
// EN: .
func (c *TColorBox) SetSelLength(value int32) {
    ColorBox_SetSelLength(c.instance, value)
}

// SelStart
// CN: 获取选择的启始位置。
// EN: .
func (c *TColorBox) SelStart() int32 {
    return ColorBox_GetSelStart(c.instance)
}

// SetSelStart
// CN: 设置选择的启始位置。
// EN: .
func (c *TColorBox) SetSelStart(value int32) {
    ColorBox_SetSelStart(c.instance, value)
}

// ItemIndex
func (c *TColorBox) ItemIndex() int32 {
    return ColorBox_GetItemIndex(c.instance)
}

// SetItemIndex
func (c *TColorBox) SetItemIndex(value int32) {
    ColorBox_SetItemIndex(c.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (c *TColorBox) DockClientCount() int32 {
    return ColorBox_GetDockClientCount(c.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (c *TColorBox) DockSite() bool {
    return ColorBox_GetDockSite(c.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (c *TColorBox) SetDockSite(value bool) {
    ColorBox_SetDockSite(c.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (c *TColorBox) AlignDisabled() bool {
    return ColorBox_GetAlignDisabled(c.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (c *TColorBox) MouseInClient() bool {
    return ColorBox_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (c *TColorBox) VisibleDockClientCount() int32 {
    return ColorBox_GetVisibleDockClientCount(c.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TColorBox) Brush() *TBrush {
    return BrushFromInst(ColorBox_GetBrush(c.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (c *TColorBox) ControlCount() int32 {
    return ColorBox_GetControlCount(c.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TColorBox) Handle() HWND {
    return ColorBox_GetHandle(c.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (c *TColorBox) ParentWindow() HWND {
    return ColorBox_GetParentWindow(c.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (c *TColorBox) SetParentWindow(value HWND) {
    ColorBox_SetParentWindow(c.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (c *TColorBox) UseDockManager() bool {
    return ColorBox_GetUseDockManager(c.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (c *TColorBox) SetUseDockManager(value bool) {
    ColorBox_SetUseDockManager(c.instance, value)
}

// Action
func (c *TColorBox) Action() *TAction {
    return ActionFromInst(ColorBox_GetAction(c.instance))
}

// SetAction
func (c *TColorBox) SetAction(value IComponent) {
    ColorBox_SetAction(c.instance, CheckPtr(value))
}

// BoundsRect
func (c *TColorBox) BoundsRect() TRect {
    return ColorBox_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TColorBox) SetBoundsRect(value TRect) {
    ColorBox_SetBoundsRect(c.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (c *TColorBox) ClientHeight() int32 {
    return ColorBox_GetClientHeight(c.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (c *TColorBox) SetClientHeight(value int32) {
    ColorBox_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TColorBox) ClientOrigin() TPoint {
    return ColorBox_GetClientOrigin(c.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (c *TColorBox) ClientRect() TRect {
    return ColorBox_GetClientRect(c.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (c *TColorBox) ClientWidth() int32 {
    return ColorBox_GetClientWidth(c.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (c *TColorBox) SetClientWidth(value int32) {
    ColorBox_SetClientWidth(c.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (c *TColorBox) ControlState() TControlState {
    return ColorBox_GetControlState(c.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (c *TColorBox) SetControlState(value TControlState) {
    ColorBox_SetControlState(c.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (c *TColorBox) ControlStyle() TControlStyle {
    return ColorBox_GetControlStyle(c.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (c *TColorBox) SetControlStyle(value TControlStyle) {
    ColorBox_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TColorBox) ExplicitLeft() int32 {
    return ColorBox_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TColorBox) ExplicitTop() int32 {
    return ColorBox_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TColorBox) ExplicitWidth() int32 {
    return ColorBox_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TColorBox) ExplicitHeight() int32 {
    return ColorBox_GetExplicitHeight(c.instance)
}

// Floating
func (c *TColorBox) Floating() bool {
    return ColorBox_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TColorBox) Parent() *TWinControl {
    return WinControlFromInst(ColorBox_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TColorBox) SetParent(value IWinControl) {
    ColorBox_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (c *TColorBox) AlignWithMargins() bool {
    return ColorBox_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (c *TColorBox) SetAlignWithMargins(value bool) {
    ColorBox_SetAlignWithMargins(c.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (c *TColorBox) Left() int32 {
    return ColorBox_GetLeft(c.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (c *TColorBox) SetLeft(value int32) {
    ColorBox_SetLeft(c.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (c *TColorBox) Top() int32 {
    return ColorBox_GetTop(c.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (c *TColorBox) SetTop(value int32) {
    ColorBox_SetTop(c.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (c *TColorBox) Width() int32 {
    return ColorBox_GetWidth(c.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (c *TColorBox) SetWidth(value int32) {
    ColorBox_SetWidth(c.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (c *TColorBox) Height() int32 {
    return ColorBox_GetHeight(c.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (c *TColorBox) SetHeight(value int32) {
    ColorBox_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TColorBox) Cursor() TCursor {
    return ColorBox_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TColorBox) SetCursor(value TCursor) {
    ColorBox_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (c *TColorBox) Hint() string {
    return ColorBox_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (c *TColorBox) SetHint(value string) {
    ColorBox_SetHint(c.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (c *TColorBox) Margins() *TMargins {
    return MarginsFromInst(ColorBox_GetMargins(c.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (c *TColorBox) SetMargins(value *TMargins) {
    ColorBox_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (c *TColorBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ColorBox_GetCustomHint(c.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (c *TColorBox) SetCustomHint(value IComponent) {
    ColorBox_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TColorBox) ComponentCount() int32 {
    return ColorBox_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TColorBox) ComponentIndex() int32 {
    return ColorBox_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TColorBox) SetComponentIndex(value int32) {
    ColorBox_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TColorBox) Owner() *TComponent {
    return ComponentFromInst(ColorBox_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TColorBox) Name() string {
    return ColorBox_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TColorBox) SetName(value string) {
    ColorBox_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TColorBox) Tag() int {
    return ColorBox_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TColorBox) SetTag(value int) {
    ColorBox_SetTag(c.instance, value)
}

// Colors
func (c *TColorBox) Colors(Index int32) TColor {
    return ColorBox_GetColors(c.instance, Index)
}

// ColorNames
func (c *TColorBox) ColorNames(Index int32) string {
    return ColorBox_GetColorNames(c.instance, Index)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (c *TColorBox) DockClients(Index int32) *TControl {
    return ControlFromInst(ColorBox_GetDockClients(c.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (c *TColorBox) Controls(Index int32) *TControl {
    return ControlFromInst(ColorBox_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TColorBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ColorBox_GetComponents(c.instance, AIndex))
}

