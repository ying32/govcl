
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

type TCheckListBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCheckListBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCheckListBox(owner IComponent) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = CheckListBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CheckListBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CheckListBoxFromInst(inst uintptr) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CheckListBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CheckListBoxFromObj(obj IObject) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CheckListBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CheckListBoxFromUnsafePointer(ptr unsafe.Pointer) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCheckListBox) Free() {
    if c.instance != 0 {
        CheckListBox_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCheckListBox) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCheckListBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCheckListBox) IsValid() bool {
    return c.instance != 0
}

// TCheckListBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCheckListBoxClass() TClass {
    return CheckListBox_StaticClassType()
}

// CheckAll
func (c *TCheckListBox) CheckAll(AState TCheckBoxState, AllowGrayed bool, AllowDisabled bool) {
    CheckListBox_CheckAll(c.instance, AState , AllowGrayed , AllowDisabled)
}

// AddItem
func (c *TCheckListBox) AddItem(Item string, AObject IObject) {
    CheckListBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

// Clear
// CN: 清除。
// EN: .
func (c *TCheckListBox) Clear() {
    CheckListBox_Clear(c.instance)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (c *TCheckListBox) ClearSelection() {
    CheckListBox_ClearSelection(c.instance)
}

// DeleteSelected
func (c *TCheckListBox) DeleteSelected() {
    CheckListBox_DeleteSelected(c.instance)
}

// SelectAll
// CN: 全选。
// EN: .
func (c *TCheckListBox) SelectAll() {
    CheckListBox_SelectAll(c.instance)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (c *TCheckListBox) CanFocus() bool {
    return CheckListBox_CanFocus(c.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (c *TCheckListBox) ContainsControl(Control IControl) bool {
    return CheckListBox_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (c *TCheckListBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(CheckListBox_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (c *TCheckListBox) DisableAlign() {
    CheckListBox_DisableAlign(c.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (c *TCheckListBox) EnableAlign() {
    CheckListBox_EnableAlign(c.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (c *TCheckListBox) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(CheckListBox_FindChildControl(c.instance, ControlName))
}

// FlipChildren
func (c *TCheckListBox) FlipChildren(AllLevels bool) {
    CheckListBox_FlipChildren(c.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (c *TCheckListBox) Focused() bool {
    return CheckListBox_Focused(c.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TCheckListBox) HandleAllocated() bool {
    return CheckListBox_HandleAllocated(c.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (c *TCheckListBox) InsertControl(AControl IControl) {
    CheckListBox_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (c *TCheckListBox) Invalidate() {
    CheckListBox_Invalidate(c.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (c *TCheckListBox) PaintTo(DC HDC, X int32, Y int32) {
    CheckListBox_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (c *TCheckListBox) RemoveControl(AControl IControl) {
    CheckListBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (c *TCheckListBox) Realign() {
    CheckListBox_Realign(c.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (c *TCheckListBox) Repaint() {
    CheckListBox_Repaint(c.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (c *TCheckListBox) ScaleBy(M int32, D int32) {
    CheckListBox_ScaleBy(c.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (c *TCheckListBox) ScrollBy(DeltaX int32, DeltaY int32) {
    CheckListBox_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (c *TCheckListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CheckListBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (c *TCheckListBox) SetFocus() {
    CheckListBox_SetFocus(c.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (c *TCheckListBox) Update() {
    CheckListBox_Update(c.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (c *TCheckListBox) UpdateControlState() {
    CheckListBox_UpdateControlState(c.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (c *TCheckListBox) BringToFront() {
    CheckListBox_BringToFront(c.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (c *TCheckListBox) ClientToScreen(Point TPoint) TPoint {
    return CheckListBox_ClientToScreen(c.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (c *TCheckListBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CheckListBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (c *TCheckListBox) Dragging() bool {
    return CheckListBox_Dragging(c.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TCheckListBox) HasParent() bool {
    return CheckListBox_HasParent(c.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (c *TCheckListBox) Hide() {
    CheckListBox_Hide(c.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (c *TCheckListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CheckListBox_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (c *TCheckListBox) Refresh() {
    CheckListBox_Refresh(c.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (c *TCheckListBox) ScreenToClient(Point TPoint) TPoint {
    return CheckListBox_ScreenToClient(c.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (c *TCheckListBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CheckListBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (c *TCheckListBox) SendToBack() {
    CheckListBox_SendToBack(c.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (c *TCheckListBox) Show() {
    CheckListBox_Show(c.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TCheckListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CheckListBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (c *TCheckListBox) GetTextLen() int32 {
    return CheckListBox_GetTextLen(c.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TCheckListBox) SetTextBuf(Buffer string) {
    CheckListBox_SetTextBuf(c.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TCheckListBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CheckListBox_FindComponent(c.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TCheckListBox) GetNamePath() string {
    return CheckListBox_GetNamePath(c.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TCheckListBox) Assign(Source IObject) {
    CheckListBox_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCheckListBox) DisposeOf() {
    CheckListBox_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCheckListBox) ClassType() TClass {
    return CheckListBox_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCheckListBox) ClassName() string {
    return CheckListBox_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCheckListBox) InstanceSize() int32 {
    return CheckListBox_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCheckListBox) InheritsFrom(AClass TClass) bool {
    return CheckListBox_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCheckListBox) Equals(Obj IObject) bool {
    return CheckListBox_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCheckListBox) GetHashCode() int32 {
    return CheckListBox_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCheckListBox) ToString() string {
    return CheckListBox_ToString(c.instance)
}

// SetOnClickCheck
func (c *TCheckListBox) SetOnClickCheck(fn TNotifyEvent) {
    CheckListBox_SetOnClickCheck(c.instance, fn)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (c *TCheckListBox) Align() TAlign {
    return CheckListBox_GetAlign(c.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (c *TCheckListBox) SetAlign(value TAlign) {
    CheckListBox_SetAlign(c.instance, value)
}

// AllowGrayed
func (c *TCheckListBox) AllowGrayed() bool {
    return CheckListBox_GetAllowGrayed(c.instance)
}

// SetAllowGrayed
func (c *TCheckListBox) SetAllowGrayed(value bool) {
    CheckListBox_SetAllowGrayed(c.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (c *TCheckListBox) Anchors() TAnchors {
    return CheckListBox_GetAnchors(c.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (c *TCheckListBox) SetAnchors(value TAnchors) {
    CheckListBox_SetAnchors(c.instance, value)
}

// AutoComplete
func (c *TCheckListBox) AutoComplete() bool {
    return CheckListBox_GetAutoComplete(c.instance)
}

// SetAutoComplete
func (c *TCheckListBox) SetAutoComplete(value bool) {
    CheckListBox_SetAutoComplete(c.instance, value)
}

// BevelEdges
func (c *TCheckListBox) BevelEdges() TBevelEdges {
    return CheckListBox_GetBevelEdges(c.instance)
}

// SetBevelEdges
func (c *TCheckListBox) SetBevelEdges(value TBevelEdges) {
    CheckListBox_SetBevelEdges(c.instance, value)
}

// BevelInner
func (c *TCheckListBox) BevelInner() TBevelCut {
    return CheckListBox_GetBevelInner(c.instance)
}

// SetBevelInner
func (c *TCheckListBox) SetBevelInner(value TBevelCut) {
    CheckListBox_SetBevelInner(c.instance, value)
}

// BevelOuter
func (c *TCheckListBox) BevelOuter() TBevelCut {
    return CheckListBox_GetBevelOuter(c.instance)
}

// SetBevelOuter
func (c *TCheckListBox) SetBevelOuter(value TBevelCut) {
    CheckListBox_SetBevelOuter(c.instance, value)
}

// BevelKind
func (c *TCheckListBox) BevelKind() TBevelKind {
    return CheckListBox_GetBevelKind(c.instance)
}

// SetBevelKind
func (c *TCheckListBox) SetBevelKind(value TBevelKind) {
    CheckListBox_SetBevelKind(c.instance, value)
}

// BiDiMode
func (c *TCheckListBox) BiDiMode() TBiDiMode {
    return CheckListBox_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TCheckListBox) SetBiDiMode(value TBiDiMode) {
    CheckListBox_SetBiDiMode(c.instance, value)
}

// BorderStyle
// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (c *TCheckListBox) BorderStyle() TBorderStyle {
    return CheckListBox_GetBorderStyle(c.instance)
}

// SetBorderStyle
// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (c *TCheckListBox) SetBorderStyle(value TBorderStyle) {
    CheckListBox_SetBorderStyle(c.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (c *TCheckListBox) Color() TColor {
    return CheckListBox_GetColor(c.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (c *TCheckListBox) SetColor(value TColor) {
    CheckListBox_SetColor(c.instance, value)
}

// Columns
func (c *TCheckListBox) Columns() int32 {
    return CheckListBox_GetColumns(c.instance)
}

// SetColumns
func (c *TCheckListBox) SetColumns(value int32) {
    CheckListBox_SetColumns(c.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (c *TCheckListBox) DoubleBuffered() bool {
    return CheckListBox_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (c *TCheckListBox) SetDoubleBuffered(value bool) {
    CheckListBox_SetDoubleBuffered(c.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (c *TCheckListBox) DragCursor() TCursor {
    return CheckListBox_GetDragCursor(c.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (c *TCheckListBox) SetDragCursor(value TCursor) {
    CheckListBox_SetDragCursor(c.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (c *TCheckListBox) DragKind() TDragKind {
    return CheckListBox_GetDragKind(c.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (c *TCheckListBox) SetDragKind(value TDragKind) {
    CheckListBox_SetDragKind(c.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (c *TCheckListBox) DragMode() TDragMode {
    return CheckListBox_GetDragMode(c.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (c *TCheckListBox) SetDragMode(value TDragMode) {
    CheckListBox_SetDragMode(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TCheckListBox) Enabled() bool {
    return CheckListBox_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TCheckListBox) SetEnabled(value bool) {
    CheckListBox_SetEnabled(c.instance, value)
}

// Flat
// CN: 获取平面样式。
// EN: .
func (c *TCheckListBox) Flat() bool {
    return CheckListBox_GetFlat(c.instance)
}

// SetFlat
// CN: 设置平面样式。
// EN: .
func (c *TCheckListBox) SetFlat(value bool) {
    CheckListBox_SetFlat(c.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (c *TCheckListBox) Font() *TFont {
    return FontFromInst(CheckListBox_GetFont(c.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (c *TCheckListBox) SetFont(value *TFont) {
    CheckListBox_SetFont(c.instance, CheckPtr(value))
}

// HeaderColor
func (c *TCheckListBox) HeaderColor() TColor {
    return CheckListBox_GetHeaderColor(c.instance)
}

// SetHeaderColor
func (c *TCheckListBox) SetHeaderColor(value TColor) {
    CheckListBox_SetHeaderColor(c.instance, value)
}

// HeaderBackgroundColor
func (c *TCheckListBox) HeaderBackgroundColor() TColor {
    return CheckListBox_GetHeaderBackgroundColor(c.instance)
}

// SetHeaderBackgroundColor
func (c *TCheckListBox) SetHeaderBackgroundColor(value TColor) {
    CheckListBox_SetHeaderBackgroundColor(c.instance, value)
}

// ItemHeight
func (c *TCheckListBox) ItemHeight() int32 {
    return CheckListBox_GetItemHeight(c.instance)
}

// SetItemHeight
func (c *TCheckListBox) SetItemHeight(value int32) {
    CheckListBox_SetItemHeight(c.instance, value)
}

// Items
func (c *TCheckListBox) Items() *TStrings {
    return StringsFromInst(CheckListBox_GetItems(c.instance))
}

// SetItems
func (c *TCheckListBox) SetItems(value IObject) {
    CheckListBox_SetItems(c.instance, CheckPtr(value))
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (c *TCheckListBox) ParentColor() bool {
    return CheckListBox_GetParentColor(c.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (c *TCheckListBox) SetParentColor(value bool) {
    CheckListBox_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TCheckListBox) ParentCtl3D() bool {
    return CheckListBox_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TCheckListBox) SetParentCtl3D(value bool) {
    CheckListBox_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (c *TCheckListBox) ParentDoubleBuffered() bool {
    return CheckListBox_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (c *TCheckListBox) SetParentDoubleBuffered(value bool) {
    CheckListBox_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (c *TCheckListBox) ParentFont() bool {
    return CheckListBox_GetParentFont(c.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (c *TCheckListBox) SetParentFont(value bool) {
    CheckListBox_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TCheckListBox) ParentShowHint() bool {
    return CheckListBox_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TCheckListBox) SetParentShowHint(value bool) {
    CheckListBox_SetParentShowHint(c.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (c *TCheckListBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CheckListBox_GetPopupMenu(c.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (c *TCheckListBox) SetPopupMenu(value IComponent) {
    CheckListBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (c *TCheckListBox) ShowHint() bool {
    return CheckListBox_GetShowHint(c.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (c *TCheckListBox) SetShowHint(value bool) {
    CheckListBox_SetShowHint(c.instance, value)
}

// Sorted
func (c *TCheckListBox) Sorted() bool {
    return CheckListBox_GetSorted(c.instance)
}

// SetSorted
func (c *TCheckListBox) SetSorted(value bool) {
    CheckListBox_SetSorted(c.instance, value)
}

// Style
func (c *TCheckListBox) Style() TListBoxStyle {
    return CheckListBox_GetStyle(c.instance)
}

// SetStyle
func (c *TCheckListBox) SetStyle(value TListBoxStyle) {
    CheckListBox_SetStyle(c.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (c *TCheckListBox) TabOrder() TTabOrder {
    return CheckListBox_GetTabOrder(c.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (c *TCheckListBox) SetTabOrder(value TTabOrder) {
    CheckListBox_SetTabOrder(c.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (c *TCheckListBox) TabStop() bool {
    return CheckListBox_GetTabStop(c.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (c *TCheckListBox) SetTabStop(value bool) {
    CheckListBox_SetTabStop(c.instance, value)
}

// TabWidth
func (c *TCheckListBox) TabWidth() int32 {
    return CheckListBox_GetTabWidth(c.instance)
}

// SetTabWidth
func (c *TCheckListBox) SetTabWidth(value int32) {
    CheckListBox_SetTabWidth(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TCheckListBox) Visible() bool {
    return CheckListBox_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TCheckListBox) SetVisible(value bool) {
    CheckListBox_SetVisible(c.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (c *TCheckListBox) StyleElements() TStyleElements {
    return CheckListBox_GetStyleElements(c.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (c *TCheckListBox) SetStyleElements(value TStyleElements) {
    CheckListBox_SetStyleElements(c.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TCheckListBox) SetOnClick(fn TNotifyEvent) {
    CheckListBox_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (c *TCheckListBox) SetOnContextPopup(fn TContextPopupEvent) {
    CheckListBox_SetOnContextPopup(c.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (c *TCheckListBox) SetOnDblClick(fn TNotifyEvent) {
    CheckListBox_SetOnDblClick(c.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (c *TCheckListBox) SetOnDragDrop(fn TDragDropEvent) {
    CheckListBox_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (c *TCheckListBox) SetOnDragOver(fn TDragOverEvent) {
    CheckListBox_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (c *TCheckListBox) SetOnEndDock(fn TEndDragEvent) {
    CheckListBox_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (c *TCheckListBox) SetOnEndDrag(fn TEndDragEvent) {
    CheckListBox_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (c *TCheckListBox) SetOnEnter(fn TNotifyEvent) {
    CheckListBox_SetOnEnter(c.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (c *TCheckListBox) SetOnExit(fn TNotifyEvent) {
    CheckListBox_SetOnExit(c.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (c *TCheckListBox) SetOnKeyDown(fn TKeyEvent) {
    CheckListBox_SetOnKeyDown(c.instance, fn)
}

// SetOnKeyPress
func (c *TCheckListBox) SetOnKeyPress(fn TKeyPressEvent) {
    CheckListBox_SetOnKeyPress(c.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (c *TCheckListBox) SetOnKeyUp(fn TKeyEvent) {
    CheckListBox_SetOnKeyUp(c.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (c *TCheckListBox) SetOnMouseDown(fn TMouseEvent) {
    CheckListBox_SetOnMouseDown(c.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (c *TCheckListBox) SetOnMouseEnter(fn TNotifyEvent) {
    CheckListBox_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (c *TCheckListBox) SetOnMouseLeave(fn TNotifyEvent) {
    CheckListBox_SetOnMouseLeave(c.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (c *TCheckListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    CheckListBox_SetOnMouseMove(c.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (c *TCheckListBox) SetOnMouseUp(fn TMouseEvent) {
    CheckListBox_SetOnMouseUp(c.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (c *TCheckListBox) SetOnStartDock(fn TStartDockEvent) {
    CheckListBox_SetOnStartDock(c.instance, fn)
}

// AutoCompleteDelay
func (c *TCheckListBox) AutoCompleteDelay() uint32 {
    return CheckListBox_GetAutoCompleteDelay(c.instance)
}

// SetAutoCompleteDelay
func (c *TCheckListBox) SetAutoCompleteDelay(value uint32) {
    CheckListBox_SetAutoCompleteDelay(c.instance, value)
}

// Canvas
// CN: 获取画布。
// EN: .
func (c *TCheckListBox) Canvas() *TCanvas {
    return CanvasFromInst(CheckListBox_GetCanvas(c.instance))
}

// Count
func (c *TCheckListBox) Count() int32 {
    return CheckListBox_GetCount(c.instance)
}

// SetCount
func (c *TCheckListBox) SetCount(value int32) {
    CheckListBox_SetCount(c.instance, value)
}

// MultiSelect
func (c *TCheckListBox) MultiSelect() bool {
    return CheckListBox_GetMultiSelect(c.instance)
}

// SetMultiSelect
func (c *TCheckListBox) SetMultiSelect(value bool) {
    CheckListBox_SetMultiSelect(c.instance, value)
}

// SelCount
func (c *TCheckListBox) SelCount() int32 {
    return CheckListBox_GetSelCount(c.instance)
}

// ItemIndex
func (c *TCheckListBox) ItemIndex() int32 {
    return CheckListBox_GetItemIndex(c.instance)
}

// SetItemIndex
func (c *TCheckListBox) SetItemIndex(value int32) {
    CheckListBox_SetItemIndex(c.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (c *TCheckListBox) DockClientCount() int32 {
    return CheckListBox_GetDockClientCount(c.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (c *TCheckListBox) DockSite() bool {
    return CheckListBox_GetDockSite(c.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (c *TCheckListBox) SetDockSite(value bool) {
    CheckListBox_SetDockSite(c.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (c *TCheckListBox) AlignDisabled() bool {
    return CheckListBox_GetAlignDisabled(c.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (c *TCheckListBox) MouseInClient() bool {
    return CheckListBox_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (c *TCheckListBox) VisibleDockClientCount() int32 {
    return CheckListBox_GetVisibleDockClientCount(c.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TCheckListBox) Brush() *TBrush {
    return BrushFromInst(CheckListBox_GetBrush(c.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (c *TCheckListBox) ControlCount() int32 {
    return CheckListBox_GetControlCount(c.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TCheckListBox) Handle() HWND {
    return CheckListBox_GetHandle(c.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (c *TCheckListBox) ParentWindow() HWND {
    return CheckListBox_GetParentWindow(c.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (c *TCheckListBox) SetParentWindow(value HWND) {
    CheckListBox_SetParentWindow(c.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (c *TCheckListBox) UseDockManager() bool {
    return CheckListBox_GetUseDockManager(c.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (c *TCheckListBox) SetUseDockManager(value bool) {
    CheckListBox_SetUseDockManager(c.instance, value)
}

// Action
func (c *TCheckListBox) Action() *TAction {
    return ActionFromInst(CheckListBox_GetAction(c.instance))
}

// SetAction
func (c *TCheckListBox) SetAction(value IComponent) {
    CheckListBox_SetAction(c.instance, CheckPtr(value))
}

// BoundsRect
func (c *TCheckListBox) BoundsRect() TRect {
    return CheckListBox_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TCheckListBox) SetBoundsRect(value TRect) {
    CheckListBox_SetBoundsRect(c.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (c *TCheckListBox) ClientHeight() int32 {
    return CheckListBox_GetClientHeight(c.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (c *TCheckListBox) SetClientHeight(value int32) {
    CheckListBox_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TCheckListBox) ClientOrigin() TPoint {
    return CheckListBox_GetClientOrigin(c.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (c *TCheckListBox) ClientRect() TRect {
    return CheckListBox_GetClientRect(c.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (c *TCheckListBox) ClientWidth() int32 {
    return CheckListBox_GetClientWidth(c.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (c *TCheckListBox) SetClientWidth(value int32) {
    CheckListBox_SetClientWidth(c.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (c *TCheckListBox) ControlState() TControlState {
    return CheckListBox_GetControlState(c.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (c *TCheckListBox) SetControlState(value TControlState) {
    CheckListBox_SetControlState(c.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (c *TCheckListBox) ControlStyle() TControlStyle {
    return CheckListBox_GetControlStyle(c.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (c *TCheckListBox) SetControlStyle(value TControlStyle) {
    CheckListBox_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TCheckListBox) ExplicitLeft() int32 {
    return CheckListBox_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TCheckListBox) ExplicitTop() int32 {
    return CheckListBox_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TCheckListBox) ExplicitWidth() int32 {
    return CheckListBox_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TCheckListBox) ExplicitHeight() int32 {
    return CheckListBox_GetExplicitHeight(c.instance)
}

// Floating
func (c *TCheckListBox) Floating() bool {
    return CheckListBox_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TCheckListBox) Parent() *TWinControl {
    return WinControlFromInst(CheckListBox_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TCheckListBox) SetParent(value IWinControl) {
    CheckListBox_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (c *TCheckListBox) AlignWithMargins() bool {
    return CheckListBox_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (c *TCheckListBox) SetAlignWithMargins(value bool) {
    CheckListBox_SetAlignWithMargins(c.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (c *TCheckListBox) Left() int32 {
    return CheckListBox_GetLeft(c.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (c *TCheckListBox) SetLeft(value int32) {
    CheckListBox_SetLeft(c.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (c *TCheckListBox) Top() int32 {
    return CheckListBox_GetTop(c.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (c *TCheckListBox) SetTop(value int32) {
    CheckListBox_SetTop(c.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (c *TCheckListBox) Width() int32 {
    return CheckListBox_GetWidth(c.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (c *TCheckListBox) SetWidth(value int32) {
    CheckListBox_SetWidth(c.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (c *TCheckListBox) Height() int32 {
    return CheckListBox_GetHeight(c.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (c *TCheckListBox) SetHeight(value int32) {
    CheckListBox_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TCheckListBox) Cursor() TCursor {
    return CheckListBox_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TCheckListBox) SetCursor(value TCursor) {
    CheckListBox_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (c *TCheckListBox) Hint() string {
    return CheckListBox_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (c *TCheckListBox) SetHint(value string) {
    CheckListBox_SetHint(c.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (c *TCheckListBox) Margins() *TMargins {
    return MarginsFromInst(CheckListBox_GetMargins(c.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (c *TCheckListBox) SetMargins(value *TMargins) {
    CheckListBox_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (c *TCheckListBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(CheckListBox_GetCustomHint(c.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (c *TCheckListBox) SetCustomHint(value IComponent) {
    CheckListBox_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TCheckListBox) ComponentCount() int32 {
    return CheckListBox_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TCheckListBox) ComponentIndex() int32 {
    return CheckListBox_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TCheckListBox) SetComponentIndex(value int32) {
    CheckListBox_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TCheckListBox) Owner() *TComponent {
    return ComponentFromInst(CheckListBox_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TCheckListBox) Name() string {
    return CheckListBox_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TCheckListBox) SetName(value string) {
    CheckListBox_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TCheckListBox) Tag() int {
    return CheckListBox_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TCheckListBox) SetTag(value int) {
    CheckListBox_SetTag(c.instance, value)
}

// Checked
// CN: 获取是否选中。
// EN: .
func (c *TCheckListBox) Checked(Index int32) bool {
    return CheckListBox_GetChecked(c.instance, Index)
}

// Checked
// CN: 设置是否选中。
// EN: .
func (c *TCheckListBox) SetChecked(Index int32, value bool) {
    CheckListBox_SetChecked(c.instance, Index, value)
}

// ItemEnabled
func (c *TCheckListBox) ItemEnabled(Index int32) bool {
    return CheckListBox_GetItemEnabled(c.instance, Index)
}

// ItemEnabled
func (c *TCheckListBox) SetItemEnabled(Index int32, value bool) {
    CheckListBox_SetItemEnabled(c.instance, Index, value)
}

// State
func (c *TCheckListBox) State(Index int32) TCheckBoxState {
    return CheckListBox_GetState(c.instance, Index)
}

// State
func (c *TCheckListBox) SetState(Index int32, value TCheckBoxState) {
    CheckListBox_SetState(c.instance, Index, value)
}

// Header
func (c *TCheckListBox) Header(Index int32) bool {
    return CheckListBox_GetHeader(c.instance, Index)
}

// Header
func (c *TCheckListBox) SetHeader(Index int32, value bool) {
    CheckListBox_SetHeader(c.instance, Index, value)
}

// Selected
func (c *TCheckListBox) Selected(Index int32) bool {
    return CheckListBox_GetSelected(c.instance, Index)
}

// Selected
func (c *TCheckListBox) SetSelected(Index int32, value bool) {
    CheckListBox_SetSelected(c.instance, Index, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (c *TCheckListBox) DockClients(Index int32) *TControl {
    return ControlFromInst(CheckListBox_GetDockClients(c.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (c *TCheckListBox) Controls(Index int32) *TControl {
    return ControlFromInst(CheckListBox_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TCheckListBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CheckListBox_GetComponents(c.instance, AIndex))
}

