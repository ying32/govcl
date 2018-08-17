
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

type TSpeedButton struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewSpeedButton
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewSpeedButton(owner IComponent) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = SpeedButton_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SpeedButtonFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func SpeedButtonFromInst(inst uintptr) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// SpeedButtonFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func SpeedButtonFromObj(obj IObject) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SpeedButtonFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func SpeedButtonFromUnsafePointer(ptr unsafe.Pointer) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TSpeedButton) Free() {
    if s.instance != 0 {
        SpeedButton_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TSpeedButton) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TSpeedButton) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TSpeedButton) IsValid() bool {
    return s.instance != 0
}

// TSpeedButtonClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TSpeedButtonClass() TClass {
    return SpeedButton_StaticClassType()
}

// Click
// CN: 单击。
// EN: .
func (s *TSpeedButton) Click() {
    SpeedButton_Click(s.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (s *TSpeedButton) BringToFront() {
    SpeedButton_BringToFront(s.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (s *TSpeedButton) ClientToScreen(Point TPoint) TPoint {
    return SpeedButton_ClientToScreen(s.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (s *TSpeedButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return SpeedButton_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (s *TSpeedButton) Dragging() bool {
    return SpeedButton_Dragging(s.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TSpeedButton) HasParent() bool {
    return SpeedButton_HasParent(s.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (s *TSpeedButton) Hide() {
    SpeedButton_Hide(s.instance)
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (s *TSpeedButton) Invalidate() {
    SpeedButton_Invalidate(s.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (s *TSpeedButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return SpeedButton_Perform(s.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (s *TSpeedButton) Refresh() {
    SpeedButton_Refresh(s.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (s *TSpeedButton) Repaint() {
    SpeedButton_Repaint(s.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (s *TSpeedButton) ScreenToClient(Point TPoint) TPoint {
    return SpeedButton_ScreenToClient(s.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (s *TSpeedButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return SpeedButton_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (s *TSpeedButton) SendToBack() {
    SpeedButton_SendToBack(s.instance)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (s *TSpeedButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    SpeedButton_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (s *TSpeedButton) Show() {
    SpeedButton_Show(s.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (s *TSpeedButton) Update() {
    SpeedButton_Update(s.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (s *TSpeedButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return SpeedButton_GetTextBuf(s.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (s *TSpeedButton) GetTextLen() int32 {
    return SpeedButton_GetTextLen(s.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (s *TSpeedButton) SetTextBuf(Buffer string) {
    SpeedButton_SetTextBuf(s.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TSpeedButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(SpeedButton_FindComponent(s.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TSpeedButton) GetNamePath() string {
    return SpeedButton_GetNamePath(s.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TSpeedButton) Assign(Source IObject) {
    SpeedButton_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TSpeedButton) DisposeOf() {
    SpeedButton_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TSpeedButton) ClassType() TClass {
    return SpeedButton_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TSpeedButton) ClassName() string {
    return SpeedButton_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TSpeedButton) InstanceSize() int32 {
    return SpeedButton_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TSpeedButton) InheritsFrom(AClass TClass) bool {
    return SpeedButton_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TSpeedButton) Equals(Obj IObject) bool {
    return SpeedButton_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TSpeedButton) GetHashCode() int32 {
    return SpeedButton_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TSpeedButton) ToString() string {
    return SpeedButton_ToString(s.instance)
}

// Action
func (s *TSpeedButton) Action() *TAction {
    return ActionFromInst(SpeedButton_GetAction(s.instance))
}

// SetAction
func (s *TSpeedButton) SetAction(value IComponent) {
    SpeedButton_SetAction(s.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (s *TSpeedButton) Align() TAlign {
    return SpeedButton_GetAlign(s.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (s *TSpeedButton) SetAlign(value TAlign) {
    SpeedButton_SetAlign(s.instance, value)
}

// AllowAllUp
func (s *TSpeedButton) AllowAllUp() bool {
    return SpeedButton_GetAllowAllUp(s.instance)
}

// SetAllowAllUp
func (s *TSpeedButton) SetAllowAllUp(value bool) {
    SpeedButton_SetAllowAllUp(s.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (s *TSpeedButton) Anchors() TAnchors {
    return SpeedButton_GetAnchors(s.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (s *TSpeedButton) SetAnchors(value TAnchors) {
    SpeedButton_SetAnchors(s.instance, value)
}

// BiDiMode
func (s *TSpeedButton) BiDiMode() TBiDiMode {
    return SpeedButton_GetBiDiMode(s.instance)
}

// SetBiDiMode
func (s *TSpeedButton) SetBiDiMode(value TBiDiMode) {
    SpeedButton_SetBiDiMode(s.instance, value)
}

// GroupIndex
// CN: 获取团组索引。
// EN: .
func (s *TSpeedButton) GroupIndex() int32 {
    return SpeedButton_GetGroupIndex(s.instance)
}

// SetGroupIndex
// CN: 设置团组索引。
// EN: .
func (s *TSpeedButton) SetGroupIndex(value int32) {
    SpeedButton_SetGroupIndex(s.instance, value)
}

// Down
func (s *TSpeedButton) Down() bool {
    return SpeedButton_GetDown(s.instance)
}

// SetDown
func (s *TSpeedButton) SetDown(value bool) {
    SpeedButton_SetDown(s.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (s *TSpeedButton) Caption() string {
    return SpeedButton_GetCaption(s.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (s *TSpeedButton) SetCaption(value string) {
    SpeedButton_SetCaption(s.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TSpeedButton) Enabled() bool {
    return SpeedButton_GetEnabled(s.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TSpeedButton) SetEnabled(value bool) {
    SpeedButton_SetEnabled(s.instance, value)
}

// Flat
// CN: 获取平面样式。
// EN: .
func (s *TSpeedButton) Flat() bool {
    return SpeedButton_GetFlat(s.instance)
}

// SetFlat
// CN: 设置平面样式。
// EN: .
func (s *TSpeedButton) SetFlat(value bool) {
    SpeedButton_SetFlat(s.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (s *TSpeedButton) Font() *TFont {
    return FontFromInst(SpeedButton_GetFont(s.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (s *TSpeedButton) SetFont(value *TFont) {
    SpeedButton_SetFont(s.instance, CheckPtr(value))
}

// Glyph
func (s *TSpeedButton) Glyph() *TBitmap {
    return BitmapFromInst(SpeedButton_GetGlyph(s.instance))
}

// SetGlyph
func (s *TSpeedButton) SetGlyph(value *TBitmap) {
    SpeedButton_SetGlyph(s.instance, CheckPtr(value))
}

// Layout
func (s *TSpeedButton) Layout() TButtonLayout {
    return SpeedButton_GetLayout(s.instance)
}

// SetLayout
func (s *TSpeedButton) SetLayout(value TButtonLayout) {
    SpeedButton_SetLayout(s.instance, value)
}

// NumGlyphs
func (s *TSpeedButton) NumGlyphs() TNumGlyphs {
    return SpeedButton_GetNumGlyphs(s.instance)
}

// SetNumGlyphs
func (s *TSpeedButton) SetNumGlyphs(value TNumGlyphs) {
    SpeedButton_SetNumGlyphs(s.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (s *TSpeedButton) ParentFont() bool {
    return SpeedButton_GetParentFont(s.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (s *TSpeedButton) SetParentFont(value bool) {
    SpeedButton_SetParentFont(s.instance, value)
}

// ParentShowHint
func (s *TSpeedButton) ParentShowHint() bool {
    return SpeedButton_GetParentShowHint(s.instance)
}

// SetParentShowHint
func (s *TSpeedButton) SetParentShowHint(value bool) {
    SpeedButton_SetParentShowHint(s.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (s *TSpeedButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(SpeedButton_GetPopupMenu(s.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (s *TSpeedButton) SetPopupMenu(value IComponent) {
    SpeedButton_SetPopupMenu(s.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (s *TSpeedButton) ShowHint() bool {
    return SpeedButton_GetShowHint(s.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (s *TSpeedButton) SetShowHint(value bool) {
    SpeedButton_SetShowHint(s.instance, value)
}

// Spacing
func (s *TSpeedButton) Spacing() int32 {
    return SpeedButton_GetSpacing(s.instance)
}

// SetSpacing
func (s *TSpeedButton) SetSpacing(value int32) {
    SpeedButton_SetSpacing(s.instance, value)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (s *TSpeedButton) Transparent() bool {
    return SpeedButton_GetTransparent(s.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (s *TSpeedButton) SetTransparent(value bool) {
    SpeedButton_SetTransparent(s.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TSpeedButton) Visible() bool {
    return SpeedButton_GetVisible(s.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TSpeedButton) SetVisible(value bool) {
    SpeedButton_SetVisible(s.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (s *TSpeedButton) StyleElements() TStyleElements {
    return SpeedButton_GetStyleElements(s.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (s *TSpeedButton) SetStyleElements(value TStyleElements) {
    SpeedButton_SetStyleElements(s.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (s *TSpeedButton) SetOnClick(fn TNotifyEvent) {
    SpeedButton_SetOnClick(s.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (s *TSpeedButton) SetOnDblClick(fn TNotifyEvent) {
    SpeedButton_SetOnDblClick(s.instance, fn)
}

// SetOnMouseActivate
func (s *TSpeedButton) SetOnMouseActivate(fn TMouseActivateEvent) {
    SpeedButton_SetOnMouseActivate(s.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (s *TSpeedButton) SetOnMouseDown(fn TMouseEvent) {
    SpeedButton_SetOnMouseDown(s.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (s *TSpeedButton) SetOnMouseEnter(fn TNotifyEvent) {
    SpeedButton_SetOnMouseEnter(s.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (s *TSpeedButton) SetOnMouseLeave(fn TNotifyEvent) {
    SpeedButton_SetOnMouseLeave(s.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (s *TSpeedButton) SetOnMouseMove(fn TMouseMoveEvent) {
    SpeedButton_SetOnMouseMove(s.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (s *TSpeedButton) SetOnMouseUp(fn TMouseEvent) {
    SpeedButton_SetOnMouseUp(s.instance, fn)
}

// BoundsRect
func (s *TSpeedButton) BoundsRect() TRect {
    return SpeedButton_GetBoundsRect(s.instance)
}

// SetBoundsRect
func (s *TSpeedButton) SetBoundsRect(value TRect) {
    SpeedButton_SetBoundsRect(s.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (s *TSpeedButton) ClientHeight() int32 {
    return SpeedButton_GetClientHeight(s.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (s *TSpeedButton) SetClientHeight(value int32) {
    SpeedButton_SetClientHeight(s.instance, value)
}

// ClientOrigin
func (s *TSpeedButton) ClientOrigin() TPoint {
    return SpeedButton_GetClientOrigin(s.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (s *TSpeedButton) ClientRect() TRect {
    return SpeedButton_GetClientRect(s.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (s *TSpeedButton) ClientWidth() int32 {
    return SpeedButton_GetClientWidth(s.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (s *TSpeedButton) SetClientWidth(value int32) {
    SpeedButton_SetClientWidth(s.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (s *TSpeedButton) ControlState() TControlState {
    return SpeedButton_GetControlState(s.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (s *TSpeedButton) SetControlState(value TControlState) {
    SpeedButton_SetControlState(s.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (s *TSpeedButton) ControlStyle() TControlStyle {
    return SpeedButton_GetControlStyle(s.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (s *TSpeedButton) SetControlStyle(value TControlStyle) {
    SpeedButton_SetControlStyle(s.instance, value)
}

// ExplicitLeft
func (s *TSpeedButton) ExplicitLeft() int32 {
    return SpeedButton_GetExplicitLeft(s.instance)
}

// ExplicitTop
func (s *TSpeedButton) ExplicitTop() int32 {
    return SpeedButton_GetExplicitTop(s.instance)
}

// ExplicitWidth
func (s *TSpeedButton) ExplicitWidth() int32 {
    return SpeedButton_GetExplicitWidth(s.instance)
}

// ExplicitHeight
func (s *TSpeedButton) ExplicitHeight() int32 {
    return SpeedButton_GetExplicitHeight(s.instance)
}

// Floating
func (s *TSpeedButton) Floating() bool {
    return SpeedButton_GetFloating(s.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TSpeedButton) Parent() *TWinControl {
    return WinControlFromInst(SpeedButton_GetParent(s.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TSpeedButton) SetParent(value IWinControl) {
    SpeedButton_SetParent(s.instance, CheckPtr(value))
}

// SetOnGesture
func (s *TSpeedButton) SetOnGesture(fn TGestureEvent) {
    SpeedButton_SetOnGesture(s.instance, fn)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (s *TSpeedButton) AlignWithMargins() bool {
    return SpeedButton_GetAlignWithMargins(s.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (s *TSpeedButton) SetAlignWithMargins(value bool) {
    SpeedButton_SetAlignWithMargins(s.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (s *TSpeedButton) Left() int32 {
    return SpeedButton_GetLeft(s.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (s *TSpeedButton) SetLeft(value int32) {
    SpeedButton_SetLeft(s.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (s *TSpeedButton) Top() int32 {
    return SpeedButton_GetTop(s.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (s *TSpeedButton) SetTop(value int32) {
    SpeedButton_SetTop(s.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (s *TSpeedButton) Width() int32 {
    return SpeedButton_GetWidth(s.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (s *TSpeedButton) SetWidth(value int32) {
    SpeedButton_SetWidth(s.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (s *TSpeedButton) Height() int32 {
    return SpeedButton_GetHeight(s.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (s *TSpeedButton) SetHeight(value int32) {
    SpeedButton_SetHeight(s.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TSpeedButton) Cursor() TCursor {
    return SpeedButton_GetCursor(s.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TSpeedButton) SetCursor(value TCursor) {
    SpeedButton_SetCursor(s.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (s *TSpeedButton) Hint() string {
    return SpeedButton_GetHint(s.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (s *TSpeedButton) SetHint(value string) {
    SpeedButton_SetHint(s.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (s *TSpeedButton) Margins() *TMargins {
    return MarginsFromInst(SpeedButton_GetMargins(s.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (s *TSpeedButton) SetMargins(value *TMargins) {
    SpeedButton_SetMargins(s.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (s *TSpeedButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(SpeedButton_GetCustomHint(s.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (s *TSpeedButton) SetCustomHint(value IComponent) {
    SpeedButton_SetCustomHint(s.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TSpeedButton) ComponentCount() int32 {
    return SpeedButton_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TSpeedButton) ComponentIndex() int32 {
    return SpeedButton_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TSpeedButton) SetComponentIndex(value int32) {
    SpeedButton_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TSpeedButton) Owner() *TComponent {
    return ComponentFromInst(SpeedButton_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TSpeedButton) Name() string {
    return SpeedButton_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TSpeedButton) SetName(value string) {
    SpeedButton_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TSpeedButton) Tag() int {
    return SpeedButton_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TSpeedButton) SetTag(value int) {
    SpeedButton_SetTag(s.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TSpeedButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(SpeedButton_GetComponents(s.instance, AIndex))
}

