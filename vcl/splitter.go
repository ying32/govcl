
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

type TSplitter struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewSplitter
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewSplitter(owner IComponent) *TSplitter {
    s := new(TSplitter)
    s.instance = Splitter_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SplitterFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func SplitterFromInst(inst uintptr) *TSplitter {
    s := new(TSplitter)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// SplitterFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func SplitterFromObj(obj IObject) *TSplitter {
    s := new(TSplitter)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SplitterFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func SplitterFromUnsafePointer(ptr unsafe.Pointer) *TSplitter {
    s := new(TSplitter)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TSplitter) Free() {
    if s.instance != 0 {
        Splitter_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TSplitter) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TSplitter) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TSplitter) IsValid() bool {
    return s.instance != 0
}

// TSplitterClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TSplitterClass() TClass {
    return Splitter_StaticClassType()
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (s *TSplitter) BringToFront() {
    Splitter_BringToFront(s.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (s *TSplitter) ClientToScreen(Point TPoint) TPoint {
    return Splitter_ClientToScreen(s.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (s *TSplitter) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Splitter_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (s *TSplitter) Dragging() bool {
    return Splitter_Dragging(s.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TSplitter) HasParent() bool {
    return Splitter_HasParent(s.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (s *TSplitter) Hide() {
    Splitter_Hide(s.instance)
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (s *TSplitter) Invalidate() {
    Splitter_Invalidate(s.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (s *TSplitter) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Splitter_Perform(s.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (s *TSplitter) Refresh() {
    Splitter_Refresh(s.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (s *TSplitter) Repaint() {
    Splitter_Repaint(s.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (s *TSplitter) ScreenToClient(Point TPoint) TPoint {
    return Splitter_ScreenToClient(s.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (s *TSplitter) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Splitter_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (s *TSplitter) SendToBack() {
    Splitter_SendToBack(s.instance)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (s *TSplitter) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Splitter_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (s *TSplitter) Show() {
    Splitter_Show(s.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (s *TSplitter) Update() {
    Splitter_Update(s.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (s *TSplitter) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Splitter_GetTextBuf(s.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (s *TSplitter) GetTextLen() int32 {
    return Splitter_GetTextLen(s.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (s *TSplitter) SetTextBuf(Buffer string) {
    Splitter_SetTextBuf(s.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TSplitter) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Splitter_FindComponent(s.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TSplitter) GetNamePath() string {
    return Splitter_GetNamePath(s.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TSplitter) Assign(Source IObject) {
    Splitter_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TSplitter) DisposeOf() {
    Splitter_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TSplitter) ClassType() TClass {
    return Splitter_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TSplitter) ClassName() string {
    return Splitter_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TSplitter) InstanceSize() int32 {
    return Splitter_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TSplitter) InheritsFrom(AClass TClass) bool {
    return Splitter_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TSplitter) Equals(Obj IObject) bool {
    return Splitter_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TSplitter) GetHashCode() int32 {
    return Splitter_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TSplitter) ToString() string {
    return Splitter_ToString(s.instance)
}

// Canvas
// CN: 获取画布。
// EN: .
func (s *TSplitter) Canvas() *TCanvas {
    return CanvasFromInst(Splitter_GetCanvas(s.instance))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (s *TSplitter) Align() TAlign {
    return Splitter_GetAlign(s.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (s *TSplitter) SetAlign(value TAlign) {
    Splitter_SetAlign(s.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (s *TSplitter) Color() TColor {
    return Splitter_GetColor(s.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (s *TSplitter) SetColor(value TColor) {
    Splitter_SetColor(s.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TSplitter) Cursor() TCursor {
    return Splitter_GetCursor(s.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TSplitter) SetCursor(value TCursor) {
    Splitter_SetCursor(s.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (s *TSplitter) ParentColor() bool {
    return Splitter_GetParentColor(s.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (s *TSplitter) SetParentColor(value bool) {
    Splitter_SetParentColor(s.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TSplitter) Visible() bool {
    return Splitter_GetVisible(s.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TSplitter) SetVisible(value bool) {
    Splitter_SetVisible(s.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (s *TSplitter) Width() int32 {
    return Splitter_GetWidth(s.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (s *TSplitter) SetWidth(value int32) {
    Splitter_SetWidth(s.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (s *TSplitter) StyleElements() TStyleElements {
    return Splitter_GetStyleElements(s.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (s *TSplitter) SetStyleElements(value TStyleElements) {
    Splitter_SetStyleElements(s.instance, value)
}

// SetOnPaint
// CN: 设置绘画事件。
// EN: .
func (s *TSplitter) SetOnPaint(fn TNotifyEvent) {
    Splitter_SetOnPaint(s.instance, fn)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TSplitter) Enabled() bool {
    return Splitter_GetEnabled(s.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TSplitter) SetEnabled(value bool) {
    Splitter_SetEnabled(s.instance, value)
}

// Action
func (s *TSplitter) Action() *TAction {
    return ActionFromInst(Splitter_GetAction(s.instance))
}

// SetAction
func (s *TSplitter) SetAction(value IComponent) {
    Splitter_SetAction(s.instance, CheckPtr(value))
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (s *TSplitter) Anchors() TAnchors {
    return Splitter_GetAnchors(s.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (s *TSplitter) SetAnchors(value TAnchors) {
    Splitter_SetAnchors(s.instance, value)
}

// BiDiMode
func (s *TSplitter) BiDiMode() TBiDiMode {
    return Splitter_GetBiDiMode(s.instance)
}

// SetBiDiMode
func (s *TSplitter) SetBiDiMode(value TBiDiMode) {
    Splitter_SetBiDiMode(s.instance, value)
}

// BoundsRect
func (s *TSplitter) BoundsRect() TRect {
    return Splitter_GetBoundsRect(s.instance)
}

// SetBoundsRect
func (s *TSplitter) SetBoundsRect(value TRect) {
    Splitter_SetBoundsRect(s.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (s *TSplitter) ClientHeight() int32 {
    return Splitter_GetClientHeight(s.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (s *TSplitter) SetClientHeight(value int32) {
    Splitter_SetClientHeight(s.instance, value)
}

// ClientOrigin
func (s *TSplitter) ClientOrigin() TPoint {
    return Splitter_GetClientOrigin(s.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (s *TSplitter) ClientRect() TRect {
    return Splitter_GetClientRect(s.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (s *TSplitter) ClientWidth() int32 {
    return Splitter_GetClientWidth(s.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (s *TSplitter) SetClientWidth(value int32) {
    Splitter_SetClientWidth(s.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (s *TSplitter) ControlState() TControlState {
    return Splitter_GetControlState(s.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (s *TSplitter) SetControlState(value TControlState) {
    Splitter_SetControlState(s.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (s *TSplitter) ControlStyle() TControlStyle {
    return Splitter_GetControlStyle(s.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (s *TSplitter) SetControlStyle(value TControlStyle) {
    Splitter_SetControlStyle(s.instance, value)
}

// ExplicitLeft
func (s *TSplitter) ExplicitLeft() int32 {
    return Splitter_GetExplicitLeft(s.instance)
}

// ExplicitTop
func (s *TSplitter) ExplicitTop() int32 {
    return Splitter_GetExplicitTop(s.instance)
}

// ExplicitWidth
func (s *TSplitter) ExplicitWidth() int32 {
    return Splitter_GetExplicitWidth(s.instance)
}

// ExplicitHeight
func (s *TSplitter) ExplicitHeight() int32 {
    return Splitter_GetExplicitHeight(s.instance)
}

// Floating
func (s *TSplitter) Floating() bool {
    return Splitter_GetFloating(s.instance)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (s *TSplitter) ShowHint() bool {
    return Splitter_GetShowHint(s.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (s *TSplitter) SetShowHint(value bool) {
    Splitter_SetShowHint(s.instance, value)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TSplitter) Parent() *TWinControl {
    return WinControlFromInst(Splitter_GetParent(s.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TSplitter) SetParent(value IWinControl) {
    Splitter_SetParent(s.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (s *TSplitter) AlignWithMargins() bool {
    return Splitter_GetAlignWithMargins(s.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (s *TSplitter) SetAlignWithMargins(value bool) {
    Splitter_SetAlignWithMargins(s.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (s *TSplitter) Left() int32 {
    return Splitter_GetLeft(s.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (s *TSplitter) SetLeft(value int32) {
    Splitter_SetLeft(s.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (s *TSplitter) Top() int32 {
    return Splitter_GetTop(s.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (s *TSplitter) SetTop(value int32) {
    Splitter_SetTop(s.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (s *TSplitter) Height() int32 {
    return Splitter_GetHeight(s.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (s *TSplitter) SetHeight(value int32) {
    Splitter_SetHeight(s.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (s *TSplitter) Hint() string {
    return Splitter_GetHint(s.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (s *TSplitter) SetHint(value string) {
    Splitter_SetHint(s.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (s *TSplitter) Margins() *TMargins {
    return MarginsFromInst(Splitter_GetMargins(s.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (s *TSplitter) SetMargins(value *TMargins) {
    Splitter_SetMargins(s.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (s *TSplitter) CustomHint() *TCustomHint {
    return CustomHintFromInst(Splitter_GetCustomHint(s.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (s *TSplitter) SetCustomHint(value IComponent) {
    Splitter_SetCustomHint(s.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TSplitter) ComponentCount() int32 {
    return Splitter_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TSplitter) ComponentIndex() int32 {
    return Splitter_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TSplitter) SetComponentIndex(value int32) {
    Splitter_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TSplitter) Owner() *TComponent {
    return ComponentFromInst(Splitter_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TSplitter) Name() string {
    return Splitter_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TSplitter) SetName(value string) {
    Splitter_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TSplitter) Tag() int {
    return Splitter_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TSplitter) SetTag(value int) {
    Splitter_SetTag(s.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TSplitter) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Splitter_GetComponents(s.instance, AIndex))
}

