
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
    "time"
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TDateTimePicker struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewDateTimePicker
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewDateTimePicker(owner IComponent) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = DateTimePicker_Create(CheckPtr(owner))
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// DateTimePickerFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func DateTimePickerFromInst(inst uintptr) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = inst
    d.ptr = unsafe.Pointer(inst)
    return d
}

// DateTimePickerFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func DateTimePickerFromObj(obj IObject) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = CheckPtr(obj)
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// DateTimePickerFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func DateTimePickerFromUnsafePointer(ptr unsafe.Pointer) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = uintptr(ptr)
    d.ptr = ptr
    return d
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (d *TDateTimePicker) Free() {
    if d.instance != 0 {
        DateTimePicker_Free(d.instance)
        d.instance = 0
        d.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (d *TDateTimePicker) Instance() uintptr {
    return d.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (d *TDateTimePicker) UnsafeAddr() unsafe.Pointer {
    return d.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (d *TDateTimePicker) IsValid() bool {
    return d.instance != 0
}

// TDateTimePickerClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TDateTimePickerClass() TClass {
    return DateTimePicker_StaticClassType()
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (d *TDateTimePicker) CanFocus() bool {
    return DateTimePicker_CanFocus(d.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (d *TDateTimePicker) ContainsControl(Control IControl) bool {
    return DateTimePicker_ContainsControl(d.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (d *TDateTimePicker) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(DateTimePicker_ControlAtPos(d.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (d *TDateTimePicker) DisableAlign() {
    DateTimePicker_DisableAlign(d.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (d *TDateTimePicker) EnableAlign() {
    DateTimePicker_EnableAlign(d.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (d *TDateTimePicker) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(DateTimePicker_FindChildControl(d.instance, ControlName))
}

// FlipChildren
func (d *TDateTimePicker) FlipChildren(AllLevels bool) {
    DateTimePicker_FlipChildren(d.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (d *TDateTimePicker) Focused() bool {
    return DateTimePicker_Focused(d.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (d *TDateTimePicker) HandleAllocated() bool {
    return DateTimePicker_HandleAllocated(d.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (d *TDateTimePicker) InsertControl(AControl IControl) {
    DateTimePicker_InsertControl(d.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (d *TDateTimePicker) Invalidate() {
    DateTimePicker_Invalidate(d.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (d *TDateTimePicker) PaintTo(DC HDC, X int32, Y int32) {
    DateTimePicker_PaintTo(d.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (d *TDateTimePicker) RemoveControl(AControl IControl) {
    DateTimePicker_RemoveControl(d.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (d *TDateTimePicker) Realign() {
    DateTimePicker_Realign(d.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (d *TDateTimePicker) Repaint() {
    DateTimePicker_Repaint(d.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (d *TDateTimePicker) ScaleBy(M int32, D int32) {
    DateTimePicker_ScaleBy(d.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (d *TDateTimePicker) ScrollBy(DeltaX int32, DeltaY int32) {
    DateTimePicker_ScrollBy(d.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (d *TDateTimePicker) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    DateTimePicker_SetBounds(d.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (d *TDateTimePicker) SetFocus() {
    DateTimePicker_SetFocus(d.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (d *TDateTimePicker) Update() {
    DateTimePicker_Update(d.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (d *TDateTimePicker) UpdateControlState() {
    DateTimePicker_UpdateControlState(d.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (d *TDateTimePicker) BringToFront() {
    DateTimePicker_BringToFront(d.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (d *TDateTimePicker) ClientToScreen(Point TPoint) TPoint {
    return DateTimePicker_ClientToScreen(d.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (d *TDateTimePicker) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return DateTimePicker_ClientToParent(d.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (d *TDateTimePicker) Dragging() bool {
    return DateTimePicker_Dragging(d.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (d *TDateTimePicker) HasParent() bool {
    return DateTimePicker_HasParent(d.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (d *TDateTimePicker) Hide() {
    DateTimePicker_Hide(d.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (d *TDateTimePicker) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return DateTimePicker_Perform(d.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (d *TDateTimePicker) Refresh() {
    DateTimePicker_Refresh(d.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (d *TDateTimePicker) ScreenToClient(Point TPoint) TPoint {
    return DateTimePicker_ScreenToClient(d.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (d *TDateTimePicker) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return DateTimePicker_ParentToClient(d.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (d *TDateTimePicker) SendToBack() {
    DateTimePicker_SendToBack(d.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (d *TDateTimePicker) Show() {
    DateTimePicker_Show(d.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (d *TDateTimePicker) GetTextBuf(Buffer string, BufSize int32) int32 {
    return DateTimePicker_GetTextBuf(d.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (d *TDateTimePicker) GetTextLen() int32 {
    return DateTimePicker_GetTextLen(d.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (d *TDateTimePicker) SetTextBuf(Buffer string) {
    DateTimePicker_SetTextBuf(d.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (d *TDateTimePicker) FindComponent(AName string) *TComponent {
    return ComponentFromInst(DateTimePicker_FindComponent(d.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (d *TDateTimePicker) GetNamePath() string {
    return DateTimePicker_GetNamePath(d.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (d *TDateTimePicker) Assign(Source IObject) {
    DateTimePicker_Assign(d.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (d *TDateTimePicker) DisposeOf() {
    DateTimePicker_DisposeOf(d.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (d *TDateTimePicker) ClassType() TClass {
    return DateTimePicker_ClassType(d.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (d *TDateTimePicker) ClassName() string {
    return DateTimePicker_ClassName(d.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (d *TDateTimePicker) InstanceSize() int32 {
    return DateTimePicker_InstanceSize(d.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (d *TDateTimePicker) InheritsFrom(AClass TClass) bool {
    return DateTimePicker_InheritsFrom(d.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (d *TDateTimePicker) Equals(Obj IObject) bool {
    return DateTimePicker_Equals(d.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (d *TDateTimePicker) GetHashCode() int32 {
    return DateTimePicker_GetHashCode(d.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (d *TDateTimePicker) ToString() string {
    return DateTimePicker_ToString(d.instance)
}

// DateTime
func (d *TDateTimePicker) DateTime() time.Time {
    return DateTimePicker_GetDateTime(d.instance)
}

// SetDateTime
func (d *TDateTimePicker) SetDateTime(value time.Time) {
    DateTimePicker_SetDateTime(d.instance, value)
}

// DroppedDown
func (d *TDateTimePicker) DroppedDown() bool {
    return DateTimePicker_GetDroppedDown(d.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (d *TDateTimePicker) Align() TAlign {
    return DateTimePicker_GetAlign(d.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (d *TDateTimePicker) SetAlign(value TAlign) {
    DateTimePicker_SetAlign(d.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (d *TDateTimePicker) Anchors() TAnchors {
    return DateTimePicker_GetAnchors(d.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (d *TDateTimePicker) SetAnchors(value TAnchors) {
    DateTimePicker_SetAnchors(d.instance, value)
}

// BevelEdges
func (d *TDateTimePicker) BevelEdges() TBevelEdges {
    return DateTimePicker_GetBevelEdges(d.instance)
}

// SetBevelEdges
func (d *TDateTimePicker) SetBevelEdges(value TBevelEdges) {
    DateTimePicker_SetBevelEdges(d.instance, value)
}

// BevelInner
func (d *TDateTimePicker) BevelInner() TBevelCut {
    return DateTimePicker_GetBevelInner(d.instance)
}

// SetBevelInner
func (d *TDateTimePicker) SetBevelInner(value TBevelCut) {
    DateTimePicker_SetBevelInner(d.instance, value)
}

// BevelOuter
func (d *TDateTimePicker) BevelOuter() TBevelCut {
    return DateTimePicker_GetBevelOuter(d.instance)
}

// SetBevelOuter
func (d *TDateTimePicker) SetBevelOuter(value TBevelCut) {
    DateTimePicker_SetBevelOuter(d.instance, value)
}

// BevelKind
func (d *TDateTimePicker) BevelKind() TBevelKind {
    return DateTimePicker_GetBevelKind(d.instance)
}

// SetBevelKind
func (d *TDateTimePicker) SetBevelKind(value TBevelKind) {
    DateTimePicker_SetBevelKind(d.instance, value)
}

// BiDiMode
func (d *TDateTimePicker) BiDiMode() TBiDiMode {
    return DateTimePicker_GetBiDiMode(d.instance)
}

// SetBiDiMode
func (d *TDateTimePicker) SetBiDiMode(value TBiDiMode) {
    DateTimePicker_SetBiDiMode(d.instance, value)
}

// CalAlignment
func (d *TDateTimePicker) CalAlignment() TDTCalAlignment {
    return DateTimePicker_GetCalAlignment(d.instance)
}

// SetCalAlignment
func (d *TDateTimePicker) SetCalAlignment(value TDTCalAlignment) {
    DateTimePicker_SetCalAlignment(d.instance, value)
}

// CalColors
func (d *TDateTimePicker) CalColors() *TMonthCalColors {
    return MonthCalColorsFromInst(DateTimePicker_GetCalColors(d.instance))
}

// SetCalColors
func (d *TDateTimePicker) SetCalColors(value *TMonthCalColors) {
    DateTimePicker_SetCalColors(d.instance, CheckPtr(value))
}

// Date
func (d *TDateTimePicker) Date() time.Time {
    return DateTimePicker_GetDate(d.instance)
}

// SetDate
func (d *TDateTimePicker) SetDate(value time.Time) {
    DateTimePicker_SetDate(d.instance, value)
}

// Format
func (d *TDateTimePicker) Format() string {
    return DateTimePicker_GetFormat(d.instance)
}

// SetFormat
func (d *TDateTimePicker) SetFormat(value string) {
    DateTimePicker_SetFormat(d.instance, value)
}

// Time
func (d *TDateTimePicker) Time() time.Time {
    return DateTimePicker_GetTime(d.instance)
}

// SetTime
func (d *TDateTimePicker) SetTime(value time.Time) {
    DateTimePicker_SetTime(d.instance, value)
}

// Checked
// CN: 获取是否选中。
// EN: .
func (d *TDateTimePicker) Checked() bool {
    return DateTimePicker_GetChecked(d.instance)
}

// SetChecked
// CN: 设置是否选中。
// EN: .
func (d *TDateTimePicker) SetChecked(value bool) {
    DateTimePicker_SetChecked(d.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (d *TDateTimePicker) Color() TColor {
    return DateTimePicker_GetColor(d.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (d *TDateTimePicker) SetColor(value TColor) {
    DateTimePicker_SetColor(d.instance, value)
}

// DateFormat
func (d *TDateTimePicker) DateFormat() TDTDateFormat {
    return DateTimePicker_GetDateFormat(d.instance)
}

// SetDateFormat
func (d *TDateTimePicker) SetDateFormat(value TDTDateFormat) {
    DateTimePicker_SetDateFormat(d.instance, value)
}

// DateMode
func (d *TDateTimePicker) DateMode() TDTDateMode {
    return DateTimePicker_GetDateMode(d.instance)
}

// SetDateMode
func (d *TDateTimePicker) SetDateMode(value TDTDateMode) {
    DateTimePicker_SetDateMode(d.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (d *TDateTimePicker) DoubleBuffered() bool {
    return DateTimePicker_GetDoubleBuffered(d.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (d *TDateTimePicker) SetDoubleBuffered(value bool) {
    DateTimePicker_SetDoubleBuffered(d.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (d *TDateTimePicker) DragCursor() TCursor {
    return DateTimePicker_GetDragCursor(d.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (d *TDateTimePicker) SetDragCursor(value TCursor) {
    DateTimePicker_SetDragCursor(d.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (d *TDateTimePicker) DragKind() TDragKind {
    return DateTimePicker_GetDragKind(d.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (d *TDateTimePicker) SetDragKind(value TDragKind) {
    DateTimePicker_SetDragKind(d.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (d *TDateTimePicker) DragMode() TDragMode {
    return DateTimePicker_GetDragMode(d.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (d *TDateTimePicker) SetDragMode(value TDragMode) {
    DateTimePicker_SetDragMode(d.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (d *TDateTimePicker) Enabled() bool {
    return DateTimePicker_GetEnabled(d.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (d *TDateTimePicker) SetEnabled(value bool) {
    DateTimePicker_SetEnabled(d.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (d *TDateTimePicker) Font() *TFont {
    return FontFromInst(DateTimePicker_GetFont(d.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (d *TDateTimePicker) SetFont(value *TFont) {
    DateTimePicker_SetFont(d.instance, CheckPtr(value))
}

// Kind
func (d *TDateTimePicker) Kind() TDateTimeKind {
    return DateTimePicker_GetKind(d.instance)
}

// SetKind
func (d *TDateTimePicker) SetKind(value TDateTimeKind) {
    DateTimePicker_SetKind(d.instance, value)
}

// MaxDate
func (d *TDateTimePicker) MaxDate() time.Time {
    return DateTimePicker_GetMaxDate(d.instance)
}

// SetMaxDate
func (d *TDateTimePicker) SetMaxDate(value time.Time) {
    DateTimePicker_SetMaxDate(d.instance, value)
}

// MinDate
func (d *TDateTimePicker) MinDate() time.Time {
    return DateTimePicker_GetMinDate(d.instance)
}

// SetMinDate
func (d *TDateTimePicker) SetMinDate(value time.Time) {
    DateTimePicker_SetMinDate(d.instance, value)
}

// ParseInput
func (d *TDateTimePicker) ParseInput() bool {
    return DateTimePicker_GetParseInput(d.instance)
}

// SetParseInput
func (d *TDateTimePicker) SetParseInput(value bool) {
    DateTimePicker_SetParseInput(d.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (d *TDateTimePicker) ParentColor() bool {
    return DateTimePicker_GetParentColor(d.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (d *TDateTimePicker) SetParentColor(value bool) {
    DateTimePicker_SetParentColor(d.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (d *TDateTimePicker) ParentDoubleBuffered() bool {
    return DateTimePicker_GetParentDoubleBuffered(d.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (d *TDateTimePicker) SetParentDoubleBuffered(value bool) {
    DateTimePicker_SetParentDoubleBuffered(d.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (d *TDateTimePicker) ParentFont() bool {
    return DateTimePicker_GetParentFont(d.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (d *TDateTimePicker) SetParentFont(value bool) {
    DateTimePicker_SetParentFont(d.instance, value)
}

// ParentShowHint
func (d *TDateTimePicker) ParentShowHint() bool {
    return DateTimePicker_GetParentShowHint(d.instance)
}

// SetParentShowHint
func (d *TDateTimePicker) SetParentShowHint(value bool) {
    DateTimePicker_SetParentShowHint(d.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (d *TDateTimePicker) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(DateTimePicker_GetPopupMenu(d.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (d *TDateTimePicker) SetPopupMenu(value IComponent) {
    DateTimePicker_SetPopupMenu(d.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (d *TDateTimePicker) ShowHint() bool {
    return DateTimePicker_GetShowHint(d.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (d *TDateTimePicker) SetShowHint(value bool) {
    DateTimePicker_SetShowHint(d.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (d *TDateTimePicker) TabOrder() TTabOrder {
    return DateTimePicker_GetTabOrder(d.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (d *TDateTimePicker) SetTabOrder(value TTabOrder) {
    DateTimePicker_SetTabOrder(d.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (d *TDateTimePicker) TabStop() bool {
    return DateTimePicker_GetTabStop(d.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (d *TDateTimePicker) SetTabStop(value bool) {
    DateTimePicker_SetTabStop(d.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (d *TDateTimePicker) Visible() bool {
    return DateTimePicker_GetVisible(d.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (d *TDateTimePicker) SetVisible(value bool) {
    DateTimePicker_SetVisible(d.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (d *TDateTimePicker) StyleElements() TStyleElements {
    return DateTimePicker_GetStyleElements(d.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (d *TDateTimePicker) SetStyleElements(value TStyleElements) {
    DateTimePicker_SetStyleElements(d.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (d *TDateTimePicker) SetOnClick(fn TNotifyEvent) {
    DateTimePicker_SetOnClick(d.instance, fn)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (d *TDateTimePicker) SetOnChange(fn TNotifyEvent) {
    DateTimePicker_SetOnChange(d.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (d *TDateTimePicker) SetOnContextPopup(fn TContextPopupEvent) {
    DateTimePicker_SetOnContextPopup(d.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (d *TDateTimePicker) SetOnDragDrop(fn TDragDropEvent) {
    DateTimePicker_SetOnDragDrop(d.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (d *TDateTimePicker) SetOnDragOver(fn TDragOverEvent) {
    DateTimePicker_SetOnDragOver(d.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (d *TDateTimePicker) SetOnEndDock(fn TEndDragEvent) {
    DateTimePicker_SetOnEndDock(d.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (d *TDateTimePicker) SetOnEndDrag(fn TEndDragEvent) {
    DateTimePicker_SetOnEndDrag(d.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (d *TDateTimePicker) SetOnEnter(fn TNotifyEvent) {
    DateTimePicker_SetOnEnter(d.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (d *TDateTimePicker) SetOnExit(fn TNotifyEvent) {
    DateTimePicker_SetOnExit(d.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (d *TDateTimePicker) SetOnKeyDown(fn TKeyEvent) {
    DateTimePicker_SetOnKeyDown(d.instance, fn)
}

// SetOnKeyPress
func (d *TDateTimePicker) SetOnKeyPress(fn TKeyPressEvent) {
    DateTimePicker_SetOnKeyPress(d.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (d *TDateTimePicker) SetOnKeyUp(fn TKeyEvent) {
    DateTimePicker_SetOnKeyUp(d.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (d *TDateTimePicker) SetOnMouseEnter(fn TNotifyEvent) {
    DateTimePicker_SetOnMouseEnter(d.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (d *TDateTimePicker) SetOnMouseLeave(fn TNotifyEvent) {
    DateTimePicker_SetOnMouseLeave(d.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (d *TDateTimePicker) SetOnStartDock(fn TStartDockEvent) {
    DateTimePicker_SetOnStartDock(d.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (d *TDateTimePicker) DockClientCount() int32 {
    return DateTimePicker_GetDockClientCount(d.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (d *TDateTimePicker) DockSite() bool {
    return DateTimePicker_GetDockSite(d.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (d *TDateTimePicker) SetDockSite(value bool) {
    DateTimePicker_SetDockSite(d.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (d *TDateTimePicker) AlignDisabled() bool {
    return DateTimePicker_GetAlignDisabled(d.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (d *TDateTimePicker) MouseInClient() bool {
    return DateTimePicker_GetMouseInClient(d.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (d *TDateTimePicker) VisibleDockClientCount() int32 {
    return DateTimePicker_GetVisibleDockClientCount(d.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (d *TDateTimePicker) Brush() *TBrush {
    return BrushFromInst(DateTimePicker_GetBrush(d.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (d *TDateTimePicker) ControlCount() int32 {
    return DateTimePicker_GetControlCount(d.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (d *TDateTimePicker) Handle() HWND {
    return DateTimePicker_GetHandle(d.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (d *TDateTimePicker) ParentWindow() HWND {
    return DateTimePicker_GetParentWindow(d.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (d *TDateTimePicker) SetParentWindow(value HWND) {
    DateTimePicker_SetParentWindow(d.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (d *TDateTimePicker) UseDockManager() bool {
    return DateTimePicker_GetUseDockManager(d.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (d *TDateTimePicker) SetUseDockManager(value bool) {
    DateTimePicker_SetUseDockManager(d.instance, value)
}

// Action
func (d *TDateTimePicker) Action() *TAction {
    return ActionFromInst(DateTimePicker_GetAction(d.instance))
}

// SetAction
func (d *TDateTimePicker) SetAction(value IComponent) {
    DateTimePicker_SetAction(d.instance, CheckPtr(value))
}

// BoundsRect
func (d *TDateTimePicker) BoundsRect() TRect {
    return DateTimePicker_GetBoundsRect(d.instance)
}

// SetBoundsRect
func (d *TDateTimePicker) SetBoundsRect(value TRect) {
    DateTimePicker_SetBoundsRect(d.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (d *TDateTimePicker) ClientHeight() int32 {
    return DateTimePicker_GetClientHeight(d.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (d *TDateTimePicker) SetClientHeight(value int32) {
    DateTimePicker_SetClientHeight(d.instance, value)
}

// ClientOrigin
func (d *TDateTimePicker) ClientOrigin() TPoint {
    return DateTimePicker_GetClientOrigin(d.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (d *TDateTimePicker) ClientRect() TRect {
    return DateTimePicker_GetClientRect(d.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (d *TDateTimePicker) ClientWidth() int32 {
    return DateTimePicker_GetClientWidth(d.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (d *TDateTimePicker) SetClientWidth(value int32) {
    DateTimePicker_SetClientWidth(d.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (d *TDateTimePicker) ControlState() TControlState {
    return DateTimePicker_GetControlState(d.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (d *TDateTimePicker) SetControlState(value TControlState) {
    DateTimePicker_SetControlState(d.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (d *TDateTimePicker) ControlStyle() TControlStyle {
    return DateTimePicker_GetControlStyle(d.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (d *TDateTimePicker) SetControlStyle(value TControlStyle) {
    DateTimePicker_SetControlStyle(d.instance, value)
}

// ExplicitLeft
func (d *TDateTimePicker) ExplicitLeft() int32 {
    return DateTimePicker_GetExplicitLeft(d.instance)
}

// ExplicitTop
func (d *TDateTimePicker) ExplicitTop() int32 {
    return DateTimePicker_GetExplicitTop(d.instance)
}

// ExplicitWidth
func (d *TDateTimePicker) ExplicitWidth() int32 {
    return DateTimePicker_GetExplicitWidth(d.instance)
}

// ExplicitHeight
func (d *TDateTimePicker) ExplicitHeight() int32 {
    return DateTimePicker_GetExplicitHeight(d.instance)
}

// Floating
func (d *TDateTimePicker) Floating() bool {
    return DateTimePicker_GetFloating(d.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (d *TDateTimePicker) Parent() *TWinControl {
    return WinControlFromInst(DateTimePicker_GetParent(d.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (d *TDateTimePicker) SetParent(value IWinControl) {
    DateTimePicker_SetParent(d.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (d *TDateTimePicker) AlignWithMargins() bool {
    return DateTimePicker_GetAlignWithMargins(d.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (d *TDateTimePicker) SetAlignWithMargins(value bool) {
    DateTimePicker_SetAlignWithMargins(d.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (d *TDateTimePicker) Left() int32 {
    return DateTimePicker_GetLeft(d.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (d *TDateTimePicker) SetLeft(value int32) {
    DateTimePicker_SetLeft(d.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (d *TDateTimePicker) Top() int32 {
    return DateTimePicker_GetTop(d.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (d *TDateTimePicker) SetTop(value int32) {
    DateTimePicker_SetTop(d.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (d *TDateTimePicker) Width() int32 {
    return DateTimePicker_GetWidth(d.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (d *TDateTimePicker) SetWidth(value int32) {
    DateTimePicker_SetWidth(d.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (d *TDateTimePicker) Height() int32 {
    return DateTimePicker_GetHeight(d.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (d *TDateTimePicker) SetHeight(value int32) {
    DateTimePicker_SetHeight(d.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (d *TDateTimePicker) Cursor() TCursor {
    return DateTimePicker_GetCursor(d.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (d *TDateTimePicker) SetCursor(value TCursor) {
    DateTimePicker_SetCursor(d.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (d *TDateTimePicker) Hint() string {
    return DateTimePicker_GetHint(d.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (d *TDateTimePicker) SetHint(value string) {
    DateTimePicker_SetHint(d.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (d *TDateTimePicker) Margins() *TMargins {
    return MarginsFromInst(DateTimePicker_GetMargins(d.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (d *TDateTimePicker) SetMargins(value *TMargins) {
    DateTimePicker_SetMargins(d.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (d *TDateTimePicker) CustomHint() *TCustomHint {
    return CustomHintFromInst(DateTimePicker_GetCustomHint(d.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (d *TDateTimePicker) SetCustomHint(value IComponent) {
    DateTimePicker_SetCustomHint(d.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (d *TDateTimePicker) ComponentCount() int32 {
    return DateTimePicker_GetComponentCount(d.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (d *TDateTimePicker) ComponentIndex() int32 {
    return DateTimePicker_GetComponentIndex(d.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (d *TDateTimePicker) SetComponentIndex(value int32) {
    DateTimePicker_SetComponentIndex(d.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (d *TDateTimePicker) Owner() *TComponent {
    return ComponentFromInst(DateTimePicker_GetOwner(d.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (d *TDateTimePicker) Name() string {
    return DateTimePicker_GetName(d.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (d *TDateTimePicker) SetName(value string) {
    DateTimePicker_SetName(d.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (d *TDateTimePicker) Tag() int {
    return DateTimePicker_GetTag(d.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (d *TDateTimePicker) SetTag(value int) {
    DateTimePicker_SetTag(d.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (d *TDateTimePicker) DockClients(Index int32) *TControl {
    return ControlFromInst(DateTimePicker_GetDockClients(d.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (d *TDateTimePicker) Controls(Index int32) *TControl {
    return ControlFromInst(DateTimePicker_GetControls(d.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (d *TDateTimePicker) Components(AIndex int32) *TComponent {
    return ComponentFromInst(DateTimePicker_GetComponents(d.instance, AIndex))
}

