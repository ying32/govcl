
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

type TListBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewListBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListBox(owner IComponent) *TListBox {
    l := new(TListBox)
    l.instance = ListBox_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ListBoxFromInst(inst uintptr) *TListBox {
    l := new(TListBox)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// ListBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ListBoxFromObj(obj IObject) *TListBox {
    l := new(TListBox)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ListBoxFromUnsafePointer(ptr unsafe.Pointer) *TListBox {
    l := new(TListBox)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TListBox) Free() {
    if l.instance != 0 {
        ListBox_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListBox) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListBox) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListBox) IsValid() bool {
    return l.instance != 0
}

// TListBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListBoxClass() TClass {
    return ListBox_StaticClassType()
}

// AddItem
func (l *TListBox) AddItem(Item string, AObject IObject) {
    ListBox_AddItem(l.instance, Item , CheckPtr(AObject))
}

// Clear
// CN: 清除。
// EN: .
func (l *TListBox) Clear() {
    ListBox_Clear(l.instance)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (l *TListBox) ClearSelection() {
    ListBox_ClearSelection(l.instance)
}

// DeleteSelected
func (l *TListBox) DeleteSelected() {
    ListBox_DeleteSelected(l.instance)
}

// SelectAll
// CN: 全选。
// EN: .
func (l *TListBox) SelectAll() {
    ListBox_SelectAll(l.instance)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (l *TListBox) CanFocus() bool {
    return ListBox_CanFocus(l.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (l *TListBox) ContainsControl(Control IControl) bool {
    return ListBox_ContainsControl(l.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (l *TListBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(ListBox_ControlAtPos(l.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (l *TListBox) DisableAlign() {
    ListBox_DisableAlign(l.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (l *TListBox) EnableAlign() {
    ListBox_EnableAlign(l.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (l *TListBox) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(ListBox_FindChildControl(l.instance, ControlName))
}

// FlipChildren
func (l *TListBox) FlipChildren(AllLevels bool) {
    ListBox_FlipChildren(l.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (l *TListBox) Focused() bool {
    return ListBox_Focused(l.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (l *TListBox) HandleAllocated() bool {
    return ListBox_HandleAllocated(l.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (l *TListBox) InsertControl(AControl IControl) {
    ListBox_InsertControl(l.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (l *TListBox) Invalidate() {
    ListBox_Invalidate(l.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (l *TListBox) PaintTo(DC HDC, X int32, Y int32) {
    ListBox_PaintTo(l.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (l *TListBox) RemoveControl(AControl IControl) {
    ListBox_RemoveControl(l.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (l *TListBox) Realign() {
    ListBox_Realign(l.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (l *TListBox) Repaint() {
    ListBox_Repaint(l.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (l *TListBox) ScaleBy(M int32, D int32) {
    ListBox_ScaleBy(l.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (l *TListBox) ScrollBy(DeltaX int32, DeltaY int32) {
    ListBox_ScrollBy(l.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (l *TListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ListBox_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (l *TListBox) SetFocus() {
    ListBox_SetFocus(l.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (l *TListBox) Update() {
    ListBox_Update(l.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (l *TListBox) UpdateControlState() {
    ListBox_UpdateControlState(l.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (l *TListBox) BringToFront() {
    ListBox_BringToFront(l.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (l *TListBox) ClientToScreen(Point TPoint) TPoint {
    return ListBox_ClientToScreen(l.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (l *TListBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ListBox_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (l *TListBox) Dragging() bool {
    return ListBox_Dragging(l.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (l *TListBox) HasParent() bool {
    return ListBox_HasParent(l.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (l *TListBox) Hide() {
    ListBox_Hide(l.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (l *TListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ListBox_Perform(l.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (l *TListBox) Refresh() {
    ListBox_Refresh(l.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (l *TListBox) ScreenToClient(Point TPoint) TPoint {
    return ListBox_ScreenToClient(l.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (l *TListBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ListBox_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (l *TListBox) SendToBack() {
    ListBox_SendToBack(l.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (l *TListBox) Show() {
    ListBox_Show(l.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (l *TListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ListBox_GetTextBuf(l.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (l *TListBox) GetTextLen() int32 {
    return ListBox_GetTextLen(l.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (l *TListBox) SetTextBuf(Buffer string) {
    ListBox_SetTextBuf(l.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (l *TListBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ListBox_FindComponent(l.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TListBox) GetNamePath() string {
    return ListBox_GetNamePath(l.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TListBox) Assign(Source IObject) {
    ListBox_Assign(l.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TListBox) DisposeOf() {
    ListBox_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListBox) ClassType() TClass {
    return ListBox_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListBox) ClassName() string {
    return ListBox_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListBox) InstanceSize() int32 {
    return ListBox_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListBox) InheritsFrom(AClass TClass) bool {
    return ListBox_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListBox) Equals(Obj IObject) bool {
    return ListBox_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListBox) GetHashCode() int32 {
    return ListBox_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TListBox) ToString() string {
    return ListBox_ToString(l.instance)
}

// Style
func (l *TListBox) Style() TListBoxStyle {
    return ListBox_GetStyle(l.instance)
}

// SetStyle
func (l *TListBox) SetStyle(value TListBoxStyle) {
    ListBox_SetStyle(l.instance, value)
}

// AutoComplete
func (l *TListBox) AutoComplete() bool {
    return ListBox_GetAutoComplete(l.instance)
}

// SetAutoComplete
func (l *TListBox) SetAutoComplete(value bool) {
    ListBox_SetAutoComplete(l.instance, value)
}

// AutoCompleteDelay
func (l *TListBox) AutoCompleteDelay() uint32 {
    return ListBox_GetAutoCompleteDelay(l.instance)
}

// SetAutoCompleteDelay
func (l *TListBox) SetAutoCompleteDelay(value uint32) {
    ListBox_SetAutoCompleteDelay(l.instance, value)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (l *TListBox) Align() TAlign {
    return ListBox_GetAlign(l.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (l *TListBox) SetAlign(value TAlign) {
    ListBox_SetAlign(l.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (l *TListBox) Anchors() TAnchors {
    return ListBox_GetAnchors(l.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (l *TListBox) SetAnchors(value TAnchors) {
    ListBox_SetAnchors(l.instance, value)
}

// BevelEdges
func (l *TListBox) BevelEdges() TBevelEdges {
    return ListBox_GetBevelEdges(l.instance)
}

// SetBevelEdges
func (l *TListBox) SetBevelEdges(value TBevelEdges) {
    ListBox_SetBevelEdges(l.instance, value)
}

// BevelInner
func (l *TListBox) BevelInner() TBevelCut {
    return ListBox_GetBevelInner(l.instance)
}

// SetBevelInner
func (l *TListBox) SetBevelInner(value TBevelCut) {
    ListBox_SetBevelInner(l.instance, value)
}

// BevelKind
func (l *TListBox) BevelKind() TBevelKind {
    return ListBox_GetBevelKind(l.instance)
}

// SetBevelKind
func (l *TListBox) SetBevelKind(value TBevelKind) {
    ListBox_SetBevelKind(l.instance, value)
}

// BevelOuter
func (l *TListBox) BevelOuter() TBevelCut {
    return ListBox_GetBevelOuter(l.instance)
}

// SetBevelOuter
func (l *TListBox) SetBevelOuter(value TBevelCut) {
    ListBox_SetBevelOuter(l.instance, value)
}

// BiDiMode
func (l *TListBox) BiDiMode() TBiDiMode {
    return ListBox_GetBiDiMode(l.instance)
}

// SetBiDiMode
func (l *TListBox) SetBiDiMode(value TBiDiMode) {
    ListBox_SetBiDiMode(l.instance, value)
}

// BorderStyle
// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (l *TListBox) BorderStyle() TBorderStyle {
    return ListBox_GetBorderStyle(l.instance)
}

// SetBorderStyle
// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (l *TListBox) SetBorderStyle(value TBorderStyle) {
    ListBox_SetBorderStyle(l.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (l *TListBox) Color() TColor {
    return ListBox_GetColor(l.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (l *TListBox) SetColor(value TColor) {
    ListBox_SetColor(l.instance, value)
}

// Columns
func (l *TListBox) Columns() int32 {
    return ListBox_GetColumns(l.instance)
}

// SetColumns
func (l *TListBox) SetColumns(value int32) {
    ListBox_SetColumns(l.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (l *TListBox) DoubleBuffered() bool {
    return ListBox_GetDoubleBuffered(l.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (l *TListBox) SetDoubleBuffered(value bool) {
    ListBox_SetDoubleBuffered(l.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (l *TListBox) DragCursor() TCursor {
    return ListBox_GetDragCursor(l.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (l *TListBox) SetDragCursor(value TCursor) {
    ListBox_SetDragCursor(l.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (l *TListBox) DragKind() TDragKind {
    return ListBox_GetDragKind(l.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (l *TListBox) SetDragKind(value TDragKind) {
    ListBox_SetDragKind(l.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (l *TListBox) DragMode() TDragMode {
    return ListBox_GetDragMode(l.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (l *TListBox) SetDragMode(value TDragMode) {
    ListBox_SetDragMode(l.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (l *TListBox) Enabled() bool {
    return ListBox_GetEnabled(l.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (l *TListBox) SetEnabled(value bool) {
    ListBox_SetEnabled(l.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (l *TListBox) Font() *TFont {
    return FontFromInst(ListBox_GetFont(l.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (l *TListBox) SetFont(value *TFont) {
    ListBox_SetFont(l.instance, CheckPtr(value))
}

// ItemHeight
func (l *TListBox) ItemHeight() int32 {
    return ListBox_GetItemHeight(l.instance)
}

// SetItemHeight
func (l *TListBox) SetItemHeight(value int32) {
    ListBox_SetItemHeight(l.instance, value)
}

// Items
func (l *TListBox) Items() *TStrings {
    return StringsFromInst(ListBox_GetItems(l.instance))
}

// SetItems
func (l *TListBox) SetItems(value IObject) {
    ListBox_SetItems(l.instance, CheckPtr(value))
}

// MultiSelect
func (l *TListBox) MultiSelect() bool {
    return ListBox_GetMultiSelect(l.instance)
}

// SetMultiSelect
func (l *TListBox) SetMultiSelect(value bool) {
    ListBox_SetMultiSelect(l.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (l *TListBox) ParentColor() bool {
    return ListBox_GetParentColor(l.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (l *TListBox) SetParentColor(value bool) {
    ListBox_SetParentColor(l.instance, value)
}

// ParentCtl3D
func (l *TListBox) ParentCtl3D() bool {
    return ListBox_GetParentCtl3D(l.instance)
}

// SetParentCtl3D
func (l *TListBox) SetParentCtl3D(value bool) {
    ListBox_SetParentCtl3D(l.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (l *TListBox) ParentDoubleBuffered() bool {
    return ListBox_GetParentDoubleBuffered(l.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (l *TListBox) SetParentDoubleBuffered(value bool) {
    ListBox_SetParentDoubleBuffered(l.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (l *TListBox) ParentFont() bool {
    return ListBox_GetParentFont(l.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (l *TListBox) SetParentFont(value bool) {
    ListBox_SetParentFont(l.instance, value)
}

// ParentShowHint
func (l *TListBox) ParentShowHint() bool {
    return ListBox_GetParentShowHint(l.instance)
}

// SetParentShowHint
func (l *TListBox) SetParentShowHint(value bool) {
    ListBox_SetParentShowHint(l.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (l *TListBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ListBox_GetPopupMenu(l.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (l *TListBox) SetPopupMenu(value IComponent) {
    ListBox_SetPopupMenu(l.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (l *TListBox) ShowHint() bool {
    return ListBox_GetShowHint(l.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (l *TListBox) SetShowHint(value bool) {
    ListBox_SetShowHint(l.instance, value)
}

// Sorted
func (l *TListBox) Sorted() bool {
    return ListBox_GetSorted(l.instance)
}

// SetSorted
func (l *TListBox) SetSorted(value bool) {
    ListBox_SetSorted(l.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (l *TListBox) TabOrder() TTabOrder {
    return ListBox_GetTabOrder(l.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (l *TListBox) SetTabOrder(value TTabOrder) {
    ListBox_SetTabOrder(l.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (l *TListBox) TabStop() bool {
    return ListBox_GetTabStop(l.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (l *TListBox) SetTabStop(value bool) {
    ListBox_SetTabStop(l.instance, value)
}

// TabWidth
func (l *TListBox) TabWidth() int32 {
    return ListBox_GetTabWidth(l.instance)
}

// SetTabWidth
func (l *TListBox) SetTabWidth(value int32) {
    ListBox_SetTabWidth(l.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (l *TListBox) Visible() bool {
    return ListBox_GetVisible(l.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (l *TListBox) SetVisible(value bool) {
    ListBox_SetVisible(l.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (l *TListBox) StyleElements() TStyleElements {
    return ListBox_GetStyleElements(l.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (l *TListBox) SetStyleElements(value TStyleElements) {
    ListBox_SetStyleElements(l.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (l *TListBox) SetOnClick(fn TNotifyEvent) {
    ListBox_SetOnClick(l.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (l *TListBox) SetOnContextPopup(fn TContextPopupEvent) {
    ListBox_SetOnContextPopup(l.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (l *TListBox) SetOnDblClick(fn TNotifyEvent) {
    ListBox_SetOnDblClick(l.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (l *TListBox) SetOnDragDrop(fn TDragDropEvent) {
    ListBox_SetOnDragDrop(l.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (l *TListBox) SetOnDragOver(fn TDragOverEvent) {
    ListBox_SetOnDragOver(l.instance, fn)
}

// SetOnDrawItem
func (l *TListBox) SetOnDrawItem(fn TDrawItemEvent) {
    ListBox_SetOnDrawItem(l.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (l *TListBox) SetOnEndDock(fn TEndDragEvent) {
    ListBox_SetOnEndDock(l.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (l *TListBox) SetOnEndDrag(fn TEndDragEvent) {
    ListBox_SetOnEndDrag(l.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (l *TListBox) SetOnEnter(fn TNotifyEvent) {
    ListBox_SetOnEnter(l.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (l *TListBox) SetOnExit(fn TNotifyEvent) {
    ListBox_SetOnExit(l.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (l *TListBox) SetOnKeyDown(fn TKeyEvent) {
    ListBox_SetOnKeyDown(l.instance, fn)
}

// SetOnKeyPress
func (l *TListBox) SetOnKeyPress(fn TKeyPressEvent) {
    ListBox_SetOnKeyPress(l.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (l *TListBox) SetOnKeyUp(fn TKeyEvent) {
    ListBox_SetOnKeyUp(l.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (l *TListBox) SetOnMouseDown(fn TMouseEvent) {
    ListBox_SetOnMouseDown(l.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (l *TListBox) SetOnMouseEnter(fn TNotifyEvent) {
    ListBox_SetOnMouseEnter(l.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (l *TListBox) SetOnMouseLeave(fn TNotifyEvent) {
    ListBox_SetOnMouseLeave(l.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (l *TListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ListBox_SetOnMouseMove(l.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (l *TListBox) SetOnMouseUp(fn TMouseEvent) {
    ListBox_SetOnMouseUp(l.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (l *TListBox) SetOnStartDock(fn TStartDockEvent) {
    ListBox_SetOnStartDock(l.instance, fn)
}

// Canvas
// CN: 获取画布。
// EN: .
func (l *TListBox) Canvas() *TCanvas {
    return CanvasFromInst(ListBox_GetCanvas(l.instance))
}

// Count
func (l *TListBox) Count() int32 {
    return ListBox_GetCount(l.instance)
}

// SetCount
func (l *TListBox) SetCount(value int32) {
    ListBox_SetCount(l.instance, value)
}

// SelCount
func (l *TListBox) SelCount() int32 {
    return ListBox_GetSelCount(l.instance)
}

// ItemIndex
func (l *TListBox) ItemIndex() int32 {
    return ListBox_GetItemIndex(l.instance)
}

// SetItemIndex
func (l *TListBox) SetItemIndex(value int32) {
    ListBox_SetItemIndex(l.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (l *TListBox) DockClientCount() int32 {
    return ListBox_GetDockClientCount(l.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (l *TListBox) DockSite() bool {
    return ListBox_GetDockSite(l.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (l *TListBox) SetDockSite(value bool) {
    ListBox_SetDockSite(l.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (l *TListBox) AlignDisabled() bool {
    return ListBox_GetAlignDisabled(l.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (l *TListBox) MouseInClient() bool {
    return ListBox_GetMouseInClient(l.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (l *TListBox) VisibleDockClientCount() int32 {
    return ListBox_GetVisibleDockClientCount(l.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (l *TListBox) Brush() *TBrush {
    return BrushFromInst(ListBox_GetBrush(l.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (l *TListBox) ControlCount() int32 {
    return ListBox_GetControlCount(l.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (l *TListBox) Handle() HWND {
    return ListBox_GetHandle(l.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (l *TListBox) ParentWindow() HWND {
    return ListBox_GetParentWindow(l.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (l *TListBox) SetParentWindow(value HWND) {
    ListBox_SetParentWindow(l.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (l *TListBox) UseDockManager() bool {
    return ListBox_GetUseDockManager(l.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (l *TListBox) SetUseDockManager(value bool) {
    ListBox_SetUseDockManager(l.instance, value)
}

// Action
func (l *TListBox) Action() *TAction {
    return ActionFromInst(ListBox_GetAction(l.instance))
}

// SetAction
func (l *TListBox) SetAction(value IComponent) {
    ListBox_SetAction(l.instance, CheckPtr(value))
}

// BoundsRect
func (l *TListBox) BoundsRect() TRect {
    return ListBox_GetBoundsRect(l.instance)
}

// SetBoundsRect
func (l *TListBox) SetBoundsRect(value TRect) {
    ListBox_SetBoundsRect(l.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (l *TListBox) ClientHeight() int32 {
    return ListBox_GetClientHeight(l.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (l *TListBox) SetClientHeight(value int32) {
    ListBox_SetClientHeight(l.instance, value)
}

// ClientOrigin
func (l *TListBox) ClientOrigin() TPoint {
    return ListBox_GetClientOrigin(l.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (l *TListBox) ClientRect() TRect {
    return ListBox_GetClientRect(l.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (l *TListBox) ClientWidth() int32 {
    return ListBox_GetClientWidth(l.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (l *TListBox) SetClientWidth(value int32) {
    ListBox_SetClientWidth(l.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (l *TListBox) ControlState() TControlState {
    return ListBox_GetControlState(l.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (l *TListBox) SetControlState(value TControlState) {
    ListBox_SetControlState(l.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (l *TListBox) ControlStyle() TControlStyle {
    return ListBox_GetControlStyle(l.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (l *TListBox) SetControlStyle(value TControlStyle) {
    ListBox_SetControlStyle(l.instance, value)
}

// ExplicitLeft
func (l *TListBox) ExplicitLeft() int32 {
    return ListBox_GetExplicitLeft(l.instance)
}

// ExplicitTop
func (l *TListBox) ExplicitTop() int32 {
    return ListBox_GetExplicitTop(l.instance)
}

// ExplicitWidth
func (l *TListBox) ExplicitWidth() int32 {
    return ListBox_GetExplicitWidth(l.instance)
}

// ExplicitHeight
func (l *TListBox) ExplicitHeight() int32 {
    return ListBox_GetExplicitHeight(l.instance)
}

// Floating
func (l *TListBox) Floating() bool {
    return ListBox_GetFloating(l.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (l *TListBox) Parent() *TWinControl {
    return WinControlFromInst(ListBox_GetParent(l.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (l *TListBox) SetParent(value IWinControl) {
    ListBox_SetParent(l.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (l *TListBox) AlignWithMargins() bool {
    return ListBox_GetAlignWithMargins(l.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (l *TListBox) SetAlignWithMargins(value bool) {
    ListBox_SetAlignWithMargins(l.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (l *TListBox) Left() int32 {
    return ListBox_GetLeft(l.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (l *TListBox) SetLeft(value int32) {
    ListBox_SetLeft(l.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (l *TListBox) Top() int32 {
    return ListBox_GetTop(l.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (l *TListBox) SetTop(value int32) {
    ListBox_SetTop(l.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (l *TListBox) Width() int32 {
    return ListBox_GetWidth(l.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (l *TListBox) SetWidth(value int32) {
    ListBox_SetWidth(l.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (l *TListBox) Height() int32 {
    return ListBox_GetHeight(l.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (l *TListBox) SetHeight(value int32) {
    ListBox_SetHeight(l.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (l *TListBox) Cursor() TCursor {
    return ListBox_GetCursor(l.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (l *TListBox) SetCursor(value TCursor) {
    ListBox_SetCursor(l.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (l *TListBox) Hint() string {
    return ListBox_GetHint(l.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (l *TListBox) SetHint(value string) {
    ListBox_SetHint(l.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (l *TListBox) Margins() *TMargins {
    return MarginsFromInst(ListBox_GetMargins(l.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (l *TListBox) SetMargins(value *TMargins) {
    ListBox_SetMargins(l.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (l *TListBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ListBox_GetCustomHint(l.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (l *TListBox) SetCustomHint(value IComponent) {
    ListBox_SetCustomHint(l.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (l *TListBox) ComponentCount() int32 {
    return ListBox_GetComponentCount(l.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (l *TListBox) ComponentIndex() int32 {
    return ListBox_GetComponentIndex(l.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (l *TListBox) SetComponentIndex(value int32) {
    ListBox_SetComponentIndex(l.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TListBox) Owner() *TComponent {
    return ComponentFromInst(ListBox_GetOwner(l.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (l *TListBox) Name() string {
    return ListBox_GetName(l.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (l *TListBox) SetName(value string) {
    ListBox_SetName(l.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TListBox) Tag() int {
    return ListBox_GetTag(l.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TListBox) SetTag(value int) {
    ListBox_SetTag(l.instance, value)
}

// Selected
func (l *TListBox) Selected(Index int32) bool {
    return ListBox_GetSelected(l.instance, Index)
}

// Selected
func (l *TListBox) SetSelected(Index int32, value bool) {
    ListBox_SetSelected(l.instance, Index, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (l *TListBox) DockClients(Index int32) *TControl {
    return ControlFromInst(ListBox_GetDockClients(l.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (l *TListBox) Controls(Index int32) *TControl {
    return ControlFromInst(ListBox_GetControls(l.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TListBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ListBox_GetComponents(l.instance, AIndex))
}

