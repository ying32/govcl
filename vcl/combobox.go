
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

type TComboBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewComboBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewComboBox(owner IComponent) *TComboBox {
    c := new(TComboBox)
    c.instance = ComboBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ComboBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ComboBoxFromInst(inst uintptr) *TComboBox {
    c := new(TComboBox)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ComboBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ComboBoxFromObj(obj IObject) *TComboBox {
    c := new(TComboBox)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ComboBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ComboBoxFromUnsafePointer(ptr unsafe.Pointer) *TComboBox {
    c := new(TComboBox)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TComboBox) Free() {
    if c.instance != 0 {
        ComboBox_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TComboBox) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TComboBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TComboBox) IsValid() bool {
    return c.instance != 0
}

// TComboBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TComboBoxClass() TClass {
    return ComboBox_StaticClassType()
}

// AddItem
func (c *TComboBox) AddItem(Item string, AObject IObject) {
    ComboBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

// Clear
// CN: 清除。
// EN: .
func (c *TComboBox) Clear() {
    ComboBox_Clear(c.instance)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (c *TComboBox) ClearSelection() {
    ComboBox_ClearSelection(c.instance)
}

// DeleteSelected
func (c *TComboBox) DeleteSelected() {
    ComboBox_DeleteSelected(c.instance)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (c *TComboBox) Focused() bool {
    return ComboBox_Focused(c.instance)
}

// SelectAll
// CN: 全选。
// EN: .
func (c *TComboBox) SelectAll() {
    ComboBox_SelectAll(c.instance)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (c *TComboBox) CanFocus() bool {
    return ComboBox_CanFocus(c.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (c *TComboBox) ContainsControl(Control IControl) bool {
    return ComboBox_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (c *TComboBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(ComboBox_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (c *TComboBox) DisableAlign() {
    ComboBox_DisableAlign(c.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (c *TComboBox) EnableAlign() {
    ComboBox_EnableAlign(c.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (c *TComboBox) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(ComboBox_FindChildControl(c.instance, ControlName))
}

// FlipChildren
func (c *TComboBox) FlipChildren(AllLevels bool) {
    ComboBox_FlipChildren(c.instance, AllLevels)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TComboBox) HandleAllocated() bool {
    return ComboBox_HandleAllocated(c.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (c *TComboBox) InsertControl(AControl IControl) {
    ComboBox_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (c *TComboBox) Invalidate() {
    ComboBox_Invalidate(c.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (c *TComboBox) PaintTo(DC HDC, X int32, Y int32) {
    ComboBox_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (c *TComboBox) RemoveControl(AControl IControl) {
    ComboBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (c *TComboBox) Realign() {
    ComboBox_Realign(c.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (c *TComboBox) Repaint() {
    ComboBox_Repaint(c.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (c *TComboBox) ScaleBy(M int32, D int32) {
    ComboBox_ScaleBy(c.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (c *TComboBox) ScrollBy(DeltaX int32, DeltaY int32) {
    ComboBox_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (c *TComboBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ComboBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (c *TComboBox) SetFocus() {
    ComboBox_SetFocus(c.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (c *TComboBox) Update() {
    ComboBox_Update(c.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (c *TComboBox) UpdateControlState() {
    ComboBox_UpdateControlState(c.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (c *TComboBox) BringToFront() {
    ComboBox_BringToFront(c.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (c *TComboBox) ClientToScreen(Point TPoint) TPoint {
    return ComboBox_ClientToScreen(c.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (c *TComboBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ComboBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (c *TComboBox) Dragging() bool {
    return ComboBox_Dragging(c.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TComboBox) HasParent() bool {
    return ComboBox_HasParent(c.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (c *TComboBox) Hide() {
    ComboBox_Hide(c.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (c *TComboBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ComboBox_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (c *TComboBox) Refresh() {
    ComboBox_Refresh(c.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (c *TComboBox) ScreenToClient(Point TPoint) TPoint {
    return ComboBox_ScreenToClient(c.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (c *TComboBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ComboBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (c *TComboBox) SendToBack() {
    ComboBox_SendToBack(c.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (c *TComboBox) Show() {
    ComboBox_Show(c.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TComboBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ComboBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (c *TComboBox) GetTextLen() int32 {
    return ComboBox_GetTextLen(c.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TComboBox) SetTextBuf(Buffer string) {
    ComboBox_SetTextBuf(c.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TComboBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ComboBox_FindComponent(c.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TComboBox) GetNamePath() string {
    return ComboBox_GetNamePath(c.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TComboBox) Assign(Source IObject) {
    ComboBox_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TComboBox) DisposeOf() {
    ComboBox_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TComboBox) ClassType() TClass {
    return ComboBox_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TComboBox) ClassName() string {
    return ComboBox_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TComboBox) InstanceSize() int32 {
    return ComboBox_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TComboBox) InheritsFrom(AClass TClass) bool {
    return ComboBox_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TComboBox) Equals(Obj IObject) bool {
    return ComboBox_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TComboBox) GetHashCode() int32 {
    return ComboBox_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TComboBox) ToString() string {
    return ComboBox_ToString(c.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (c *TComboBox) Align() TAlign {
    return ComboBox_GetAlign(c.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (c *TComboBox) SetAlign(value TAlign) {
    ComboBox_SetAlign(c.instance, value)
}

// AutoComplete
func (c *TComboBox) AutoComplete() bool {
    return ComboBox_GetAutoComplete(c.instance)
}

// SetAutoComplete
func (c *TComboBox) SetAutoComplete(value bool) {
    ComboBox_SetAutoComplete(c.instance, value)
}

// AutoCompleteDelay
func (c *TComboBox) AutoCompleteDelay() uint32 {
    return ComboBox_GetAutoCompleteDelay(c.instance)
}

// SetAutoCompleteDelay
func (c *TComboBox) SetAutoCompleteDelay(value uint32) {
    ComboBox_SetAutoCompleteDelay(c.instance, value)
}

// AutoDropDown
func (c *TComboBox) AutoDropDown() bool {
    return ComboBox_GetAutoDropDown(c.instance)
}

// SetAutoDropDown
func (c *TComboBox) SetAutoDropDown(value bool) {
    ComboBox_SetAutoDropDown(c.instance, value)
}

// AutoCloseUp
func (c *TComboBox) AutoCloseUp() bool {
    return ComboBox_GetAutoCloseUp(c.instance)
}

// SetAutoCloseUp
func (c *TComboBox) SetAutoCloseUp(value bool) {
    ComboBox_SetAutoCloseUp(c.instance, value)
}

// BevelEdges
func (c *TComboBox) BevelEdges() TBevelEdges {
    return ComboBox_GetBevelEdges(c.instance)
}

// SetBevelEdges
func (c *TComboBox) SetBevelEdges(value TBevelEdges) {
    ComboBox_SetBevelEdges(c.instance, value)
}

// BevelInner
func (c *TComboBox) BevelInner() TBevelCut {
    return ComboBox_GetBevelInner(c.instance)
}

// SetBevelInner
func (c *TComboBox) SetBevelInner(value TBevelCut) {
    ComboBox_SetBevelInner(c.instance, value)
}

// BevelKind
func (c *TComboBox) BevelKind() TBevelKind {
    return ComboBox_GetBevelKind(c.instance)
}

// SetBevelKind
func (c *TComboBox) SetBevelKind(value TBevelKind) {
    ComboBox_SetBevelKind(c.instance, value)
}

// BevelOuter
func (c *TComboBox) BevelOuter() TBevelCut {
    return ComboBox_GetBevelOuter(c.instance)
}

// SetBevelOuter
func (c *TComboBox) SetBevelOuter(value TBevelCut) {
    ComboBox_SetBevelOuter(c.instance, value)
}

// Style
func (c *TComboBox) Style() TComboBoxStyle {
    return ComboBox_GetStyle(c.instance)
}

// SetStyle
func (c *TComboBox) SetStyle(value TComboBoxStyle) {
    ComboBox_SetStyle(c.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (c *TComboBox) Anchors() TAnchors {
    return ComboBox_GetAnchors(c.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (c *TComboBox) SetAnchors(value TAnchors) {
    ComboBox_SetAnchors(c.instance, value)
}

// BiDiMode
func (c *TComboBox) BiDiMode() TBiDiMode {
    return ComboBox_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TComboBox) SetBiDiMode(value TBiDiMode) {
    ComboBox_SetBiDiMode(c.instance, value)
}

// CharCase
func (c *TComboBox) CharCase() TEditCharCase {
    return ComboBox_GetCharCase(c.instance)
}

// SetCharCase
func (c *TComboBox) SetCharCase(value TEditCharCase) {
    ComboBox_SetCharCase(c.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (c *TComboBox) Color() TColor {
    return ComboBox_GetColor(c.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (c *TComboBox) SetColor(value TColor) {
    ComboBox_SetColor(c.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (c *TComboBox) DoubleBuffered() bool {
    return ComboBox_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (c *TComboBox) SetDoubleBuffered(value bool) {
    ComboBox_SetDoubleBuffered(c.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (c *TComboBox) DragCursor() TCursor {
    return ComboBox_GetDragCursor(c.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (c *TComboBox) SetDragCursor(value TCursor) {
    ComboBox_SetDragCursor(c.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (c *TComboBox) DragKind() TDragKind {
    return ComboBox_GetDragKind(c.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (c *TComboBox) SetDragKind(value TDragKind) {
    ComboBox_SetDragKind(c.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (c *TComboBox) DragMode() TDragMode {
    return ComboBox_GetDragMode(c.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (c *TComboBox) SetDragMode(value TDragMode) {
    ComboBox_SetDragMode(c.instance, value)
}

// DropDownCount
func (c *TComboBox) DropDownCount() int32 {
    return ComboBox_GetDropDownCount(c.instance)
}

// SetDropDownCount
func (c *TComboBox) SetDropDownCount(value int32) {
    ComboBox_SetDropDownCount(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TComboBox) Enabled() bool {
    return ComboBox_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TComboBox) SetEnabled(value bool) {
    ComboBox_SetEnabled(c.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (c *TComboBox) Font() *TFont {
    return FontFromInst(ComboBox_GetFont(c.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (c *TComboBox) SetFont(value *TFont) {
    ComboBox_SetFont(c.instance, CheckPtr(value))
}

// ItemHeight
func (c *TComboBox) ItemHeight() int32 {
    return ComboBox_GetItemHeight(c.instance)
}

// SetItemHeight
func (c *TComboBox) SetItemHeight(value int32) {
    ComboBox_SetItemHeight(c.instance, value)
}

// ItemIndex
func (c *TComboBox) ItemIndex() int32 {
    return ComboBox_GetItemIndex(c.instance)
}

// SetItemIndex
func (c *TComboBox) SetItemIndex(value int32) {
    ComboBox_SetItemIndex(c.instance, value)
}

// MaxLength
// CN: 获取最大长度。
// EN: .
func (c *TComboBox) MaxLength() int32 {
    return ComboBox_GetMaxLength(c.instance)
}

// SetMaxLength
// CN: 设置最大长度。
// EN: .
func (c *TComboBox) SetMaxLength(value int32) {
    ComboBox_SetMaxLength(c.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (c *TComboBox) ParentColor() bool {
    return ComboBox_GetParentColor(c.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (c *TComboBox) SetParentColor(value bool) {
    ComboBox_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TComboBox) ParentCtl3D() bool {
    return ComboBox_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TComboBox) SetParentCtl3D(value bool) {
    ComboBox_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (c *TComboBox) ParentDoubleBuffered() bool {
    return ComboBox_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (c *TComboBox) SetParentDoubleBuffered(value bool) {
    ComboBox_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (c *TComboBox) ParentFont() bool {
    return ComboBox_GetParentFont(c.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (c *TComboBox) SetParentFont(value bool) {
    ComboBox_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TComboBox) ParentShowHint() bool {
    return ComboBox_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TComboBox) SetParentShowHint(value bool) {
    ComboBox_SetParentShowHint(c.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (c *TComboBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ComboBox_GetPopupMenu(c.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (c *TComboBox) SetPopupMenu(value IComponent) {
    ComboBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (c *TComboBox) ShowHint() bool {
    return ComboBox_GetShowHint(c.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (c *TComboBox) SetShowHint(value bool) {
    ComboBox_SetShowHint(c.instance, value)
}

// Sorted
func (c *TComboBox) Sorted() bool {
    return ComboBox_GetSorted(c.instance)
}

// SetSorted
func (c *TComboBox) SetSorted(value bool) {
    ComboBox_SetSorted(c.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (c *TComboBox) TabOrder() TTabOrder {
    return ComboBox_GetTabOrder(c.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (c *TComboBox) SetTabOrder(value TTabOrder) {
    ComboBox_SetTabOrder(c.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (c *TComboBox) TabStop() bool {
    return ComboBox_GetTabStop(c.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (c *TComboBox) SetTabStop(value bool) {
    ComboBox_SetTabStop(c.instance, value)
}

// Text
// CN: 获取文本。
// EN: .
func (c *TComboBox) Text() string {
    return ComboBox_GetText(c.instance)
}

// SetText
// CN: 设置文本。
// EN: .
func (c *TComboBox) SetText(value string) {
    ComboBox_SetText(c.instance, value)
}

// TextHint
// CN: 获取提示文本。
// EN: .
func (c *TComboBox) TextHint() string {
    return ComboBox_GetTextHint(c.instance)
}

// SetTextHint
// CN: 设置提示文本。
// EN: .
func (c *TComboBox) SetTextHint(value string) {
    ComboBox_SetTextHint(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TComboBox) Visible() bool {
    return ComboBox_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TComboBox) SetVisible(value bool) {
    ComboBox_SetVisible(c.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (c *TComboBox) StyleElements() TStyleElements {
    return ComboBox_GetStyleElements(c.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (c *TComboBox) SetStyleElements(value TStyleElements) {
    ComboBox_SetStyleElements(c.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (c *TComboBox) SetOnChange(fn TNotifyEvent) {
    ComboBox_SetOnChange(c.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TComboBox) SetOnClick(fn TNotifyEvent) {
    ComboBox_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (c *TComboBox) SetOnContextPopup(fn TContextPopupEvent) {
    ComboBox_SetOnContextPopup(c.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (c *TComboBox) SetOnDblClick(fn TNotifyEvent) {
    ComboBox_SetOnDblClick(c.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (c *TComboBox) SetOnDragDrop(fn TDragDropEvent) {
    ComboBox_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (c *TComboBox) SetOnDragOver(fn TDragOverEvent) {
    ComboBox_SetOnDragOver(c.instance, fn)
}

// SetOnDrawItem
func (c *TComboBox) SetOnDrawItem(fn TDrawItemEvent) {
    ComboBox_SetOnDrawItem(c.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (c *TComboBox) SetOnEndDock(fn TEndDragEvent) {
    ComboBox_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (c *TComboBox) SetOnEndDrag(fn TEndDragEvent) {
    ComboBox_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (c *TComboBox) SetOnEnter(fn TNotifyEvent) {
    ComboBox_SetOnEnter(c.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (c *TComboBox) SetOnExit(fn TNotifyEvent) {
    ComboBox_SetOnExit(c.instance, fn)
}

// SetOnGesture
func (c *TComboBox) SetOnGesture(fn TGestureEvent) {
    ComboBox_SetOnGesture(c.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (c *TComboBox) SetOnKeyDown(fn TKeyEvent) {
    ComboBox_SetOnKeyDown(c.instance, fn)
}

// SetOnKeyPress
func (c *TComboBox) SetOnKeyPress(fn TKeyPressEvent) {
    ComboBox_SetOnKeyPress(c.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (c *TComboBox) SetOnKeyUp(fn TKeyEvent) {
    ComboBox_SetOnKeyUp(c.instance, fn)
}

// SetOnMeasureItem
func (c *TComboBox) SetOnMeasureItem(fn TMeasureItemEvent) {
    ComboBox_SetOnMeasureItem(c.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (c *TComboBox) SetOnMouseEnter(fn TNotifyEvent) {
    ComboBox_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (c *TComboBox) SetOnMouseLeave(fn TNotifyEvent) {
    ComboBox_SetOnMouseLeave(c.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (c *TComboBox) SetOnStartDock(fn TStartDockEvent) {
    ComboBox_SetOnStartDock(c.instance, fn)
}

// Items
func (c *TComboBox) Items() *TStrings {
    return StringsFromInst(ComboBox_GetItems(c.instance))
}

// SetItems
func (c *TComboBox) SetItems(value IObject) {
    ComboBox_SetItems(c.instance, CheckPtr(value))
}

// SelText
// CN: 获取选择的文本。
// EN: .
func (c *TComboBox) SelText() string {
    return ComboBox_GetSelText(c.instance)
}

// SetSelText
// CN: 设置选择的文本。
// EN: .
func (c *TComboBox) SetSelText(value string) {
    ComboBox_SetSelText(c.instance, value)
}

// Canvas
// CN: 获取画布。
// EN: .
func (c *TComboBox) Canvas() *TCanvas {
    return CanvasFromInst(ComboBox_GetCanvas(c.instance))
}

// DroppedDown
func (c *TComboBox) DroppedDown() bool {
    return ComboBox_GetDroppedDown(c.instance)
}

// SetDroppedDown
func (c *TComboBox) SetDroppedDown(value bool) {
    ComboBox_SetDroppedDown(c.instance, value)
}

// SelLength
// CN: 获取选择的长度。
// EN: .
func (c *TComboBox) SelLength() int32 {
    return ComboBox_GetSelLength(c.instance)
}

// SetSelLength
// CN: 设置选择的长度。
// EN: .
func (c *TComboBox) SetSelLength(value int32) {
    ComboBox_SetSelLength(c.instance, value)
}

// SelStart
// CN: 获取选择的启始位置。
// EN: .
func (c *TComboBox) SelStart() int32 {
    return ComboBox_GetSelStart(c.instance)
}

// SetSelStart
// CN: 设置选择的启始位置。
// EN: .
func (c *TComboBox) SetSelStart(value int32) {
    ComboBox_SetSelStart(c.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (c *TComboBox) DockClientCount() int32 {
    return ComboBox_GetDockClientCount(c.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (c *TComboBox) DockSite() bool {
    return ComboBox_GetDockSite(c.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (c *TComboBox) SetDockSite(value bool) {
    ComboBox_SetDockSite(c.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (c *TComboBox) AlignDisabled() bool {
    return ComboBox_GetAlignDisabled(c.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (c *TComboBox) MouseInClient() bool {
    return ComboBox_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (c *TComboBox) VisibleDockClientCount() int32 {
    return ComboBox_GetVisibleDockClientCount(c.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TComboBox) Brush() *TBrush {
    return BrushFromInst(ComboBox_GetBrush(c.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (c *TComboBox) ControlCount() int32 {
    return ComboBox_GetControlCount(c.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TComboBox) Handle() HWND {
    return ComboBox_GetHandle(c.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (c *TComboBox) ParentWindow() HWND {
    return ComboBox_GetParentWindow(c.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (c *TComboBox) SetParentWindow(value HWND) {
    ComboBox_SetParentWindow(c.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (c *TComboBox) UseDockManager() bool {
    return ComboBox_GetUseDockManager(c.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (c *TComboBox) SetUseDockManager(value bool) {
    ComboBox_SetUseDockManager(c.instance, value)
}

// Action
func (c *TComboBox) Action() *TAction {
    return ActionFromInst(ComboBox_GetAction(c.instance))
}

// SetAction
func (c *TComboBox) SetAction(value IComponent) {
    ComboBox_SetAction(c.instance, CheckPtr(value))
}

// BoundsRect
func (c *TComboBox) BoundsRect() TRect {
    return ComboBox_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TComboBox) SetBoundsRect(value TRect) {
    ComboBox_SetBoundsRect(c.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (c *TComboBox) ClientHeight() int32 {
    return ComboBox_GetClientHeight(c.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (c *TComboBox) SetClientHeight(value int32) {
    ComboBox_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TComboBox) ClientOrigin() TPoint {
    return ComboBox_GetClientOrigin(c.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (c *TComboBox) ClientRect() TRect {
    return ComboBox_GetClientRect(c.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (c *TComboBox) ClientWidth() int32 {
    return ComboBox_GetClientWidth(c.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (c *TComboBox) SetClientWidth(value int32) {
    ComboBox_SetClientWidth(c.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (c *TComboBox) ControlState() TControlState {
    return ComboBox_GetControlState(c.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (c *TComboBox) SetControlState(value TControlState) {
    ComboBox_SetControlState(c.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (c *TComboBox) ControlStyle() TControlStyle {
    return ComboBox_GetControlStyle(c.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (c *TComboBox) SetControlStyle(value TControlStyle) {
    ComboBox_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TComboBox) ExplicitLeft() int32 {
    return ComboBox_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TComboBox) ExplicitTop() int32 {
    return ComboBox_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TComboBox) ExplicitWidth() int32 {
    return ComboBox_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TComboBox) ExplicitHeight() int32 {
    return ComboBox_GetExplicitHeight(c.instance)
}

// Floating
func (c *TComboBox) Floating() bool {
    return ComboBox_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TComboBox) Parent() *TWinControl {
    return WinControlFromInst(ComboBox_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TComboBox) SetParent(value IWinControl) {
    ComboBox_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (c *TComboBox) AlignWithMargins() bool {
    return ComboBox_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (c *TComboBox) SetAlignWithMargins(value bool) {
    ComboBox_SetAlignWithMargins(c.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (c *TComboBox) Left() int32 {
    return ComboBox_GetLeft(c.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (c *TComboBox) SetLeft(value int32) {
    ComboBox_SetLeft(c.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (c *TComboBox) Top() int32 {
    return ComboBox_GetTop(c.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (c *TComboBox) SetTop(value int32) {
    ComboBox_SetTop(c.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (c *TComboBox) Width() int32 {
    return ComboBox_GetWidth(c.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (c *TComboBox) SetWidth(value int32) {
    ComboBox_SetWidth(c.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (c *TComboBox) Height() int32 {
    return ComboBox_GetHeight(c.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (c *TComboBox) SetHeight(value int32) {
    ComboBox_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TComboBox) Cursor() TCursor {
    return ComboBox_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TComboBox) SetCursor(value TCursor) {
    ComboBox_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (c *TComboBox) Hint() string {
    return ComboBox_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (c *TComboBox) SetHint(value string) {
    ComboBox_SetHint(c.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (c *TComboBox) Margins() *TMargins {
    return MarginsFromInst(ComboBox_GetMargins(c.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (c *TComboBox) SetMargins(value *TMargins) {
    ComboBox_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (c *TComboBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ComboBox_GetCustomHint(c.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (c *TComboBox) SetCustomHint(value IComponent) {
    ComboBox_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TComboBox) ComponentCount() int32 {
    return ComboBox_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TComboBox) ComponentIndex() int32 {
    return ComboBox_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TComboBox) SetComponentIndex(value int32) {
    ComboBox_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TComboBox) Owner() *TComponent {
    return ComponentFromInst(ComboBox_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TComboBox) Name() string {
    return ComboBox_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TComboBox) SetName(value string) {
    ComboBox_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TComboBox) Tag() int {
    return ComboBox_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TComboBox) SetTag(value int) {
    ComboBox_SetTag(c.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (c *TComboBox) DockClients(Index int32) *TControl {
    return ControlFromInst(ComboBox_GetDockClients(c.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (c *TComboBox) Controls(Index int32) *TControl {
    return ControlFromInst(ComboBox_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TComboBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ComboBox_GetComponents(c.instance, AIndex))
}

