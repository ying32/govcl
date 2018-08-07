
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

type TRadioButton struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewRadioButton
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewRadioButton(owner IComponent) *TRadioButton {
    r := new(TRadioButton)
    r.instance = RadioButton_Create(CheckPtr(owner))
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RadioButtonFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func RadioButtonFromInst(inst uintptr) *TRadioButton {
    r := new(TRadioButton)
    r.instance = inst
    r.ptr = unsafe.Pointer(inst)
    return r
}

// RadioButtonFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func RadioButtonFromObj(obj IObject) *TRadioButton {
    r := new(TRadioButton)
    r.instance = CheckPtr(obj)
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RadioButtonFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func RadioButtonFromUnsafePointer(ptr unsafe.Pointer) *TRadioButton {
    r := new(TRadioButton)
    r.instance = uintptr(ptr)
    r.ptr = ptr
    return r
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (r *TRadioButton) Free() {
    if r.instance != 0 {
        RadioButton_Free(r.instance)
        r.instance = 0
        r.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (r *TRadioButton) Instance() uintptr {
    return r.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (r *TRadioButton) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (r *TRadioButton) IsValid() bool {
    return r.instance != 0
}

// TRadioButtonClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TRadioButtonClass() TClass {
    return RadioButton_StaticClassType()
}

// CanFocus
func (r *TRadioButton) CanFocus() bool {
    return RadioButton_CanFocus(r.instance)
}

// FlipChildren
func (r *TRadioButton) FlipChildren(AllLevels bool) {
    RadioButton_FlipChildren(r.instance, AllLevels)
}

// Focused
func (r *TRadioButton) Focused() bool {
    return RadioButton_Focused(r.instance)
}

// HandleAllocated
func (r *TRadioButton) HandleAllocated() bool {
    return RadioButton_HandleAllocated(r.instance)
}

// Invalidate
func (r *TRadioButton) Invalidate() {
    RadioButton_Invalidate(r.instance)
}

// Realign
func (r *TRadioButton) Realign() {
    RadioButton_Realign(r.instance)
}

// Repaint
func (r *TRadioButton) Repaint() {
    RadioButton_Repaint(r.instance)
}

// ScaleBy
func (r *TRadioButton) ScaleBy(M int32, D int32) {
    RadioButton_ScaleBy(r.instance, M , D)
}

// SetBounds
func (r *TRadioButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RadioButton_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (r *TRadioButton) SetFocus() {
    RadioButton_SetFocus(r.instance)
}

// Update
func (r *TRadioButton) Update() {
    RadioButton_Update(r.instance)
}

// BringToFront
func (r *TRadioButton) BringToFront() {
    RadioButton_BringToFront(r.instance)
}

// ClientToScreen
func (r *TRadioButton) ClientToScreen(Point TPoint) TPoint {
    return RadioButton_ClientToScreen(r.instance, Point)
}

// ClientToParent
func (r *TRadioButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return RadioButton_ClientToParent(r.instance, Point , CheckPtr(AParent))
}

// Dragging
func (r *TRadioButton) Dragging() bool {
    return RadioButton_Dragging(r.instance)
}

// HasParent
func (r *TRadioButton) HasParent() bool {
    return RadioButton_HasParent(r.instance)
}

// Hide
func (r *TRadioButton) Hide() {
    RadioButton_Hide(r.instance)
}

// Perform
func (r *TRadioButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RadioButton_Perform(r.instance, Msg , WParam , LParam)
}

// Refresh
func (r *TRadioButton) Refresh() {
    RadioButton_Refresh(r.instance)
}

// ScreenToClient
func (r *TRadioButton) ScreenToClient(Point TPoint) TPoint {
    return RadioButton_ScreenToClient(r.instance, Point)
}

// ParentToClient
func (r *TRadioButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return RadioButton_ParentToClient(r.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (r *TRadioButton) SendToBack() {
    RadioButton_SendToBack(r.instance)
}

// Show
func (r *TRadioButton) Show() {
    RadioButton_Show(r.instance)
}

// GetTextBuf
func (r *TRadioButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return RadioButton_GetTextBuf(r.instance, Buffer , BufSize)
}

// GetTextLen
func (r *TRadioButton) GetTextLen() int32 {
    return RadioButton_GetTextLen(r.instance)
}

// FindComponent
func (r *TRadioButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(RadioButton_FindComponent(r.instance, AName))
}

// GetNamePath
func (r *TRadioButton) GetNamePath() string {
    return RadioButton_GetNamePath(r.instance)
}

// Assign
func (r *TRadioButton) Assign(Source IObject) {
    RadioButton_Assign(r.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (r *TRadioButton) DisposeOf() {
    RadioButton_DisposeOf(r.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (r *TRadioButton) ClassType() TClass {
    return RadioButton_ClassType(r.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (r *TRadioButton) ClassName() string {
    return RadioButton_ClassName(r.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (r *TRadioButton) InstanceSize() int32 {
    return RadioButton_InstanceSize(r.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (r *TRadioButton) InheritsFrom(AClass TClass) bool {
    return RadioButton_InheritsFrom(r.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (r *TRadioButton) Equals(Obj IObject) bool {
    return RadioButton_Equals(r.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (r *TRadioButton) GetHashCode() int32 {
    return RadioButton_GetHashCode(r.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (r *TRadioButton) ToString() string {
    return RadioButton_ToString(r.instance)
}

// Action
func (r *TRadioButton) Action() *TAction {
    return ActionFromInst(RadioButton_GetAction(r.instance))
}

// SetAction
func (r *TRadioButton) SetAction(value IComponent) {
    RadioButton_SetAction(r.instance, CheckPtr(value))
}

// Align
func (r *TRadioButton) Align() TAlign {
    return RadioButton_GetAlign(r.instance)
}

// SetAlign
func (r *TRadioButton) SetAlign(value TAlign) {
    RadioButton_SetAlign(r.instance, value)
}

// Alignment
func (r *TRadioButton) Alignment() TLeftRight {
    return RadioButton_GetAlignment(r.instance)
}

// SetAlignment
func (r *TRadioButton) SetAlignment(value TLeftRight) {
    RadioButton_SetAlignment(r.instance, value)
}

// Anchors
func (r *TRadioButton) Anchors() TAnchors {
    return RadioButton_GetAnchors(r.instance)
}

// SetAnchors
func (r *TRadioButton) SetAnchors(value TAnchors) {
    RadioButton_SetAnchors(r.instance, value)
}

// BiDiMode
func (r *TRadioButton) BiDiMode() TBiDiMode {
    return RadioButton_GetBiDiMode(r.instance)
}

// SetBiDiMode
func (r *TRadioButton) SetBiDiMode(value TBiDiMode) {
    RadioButton_SetBiDiMode(r.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (r *TRadioButton) Caption() string {
    return RadioButton_GetCaption(r.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (r *TRadioButton) SetCaption(value string) {
    RadioButton_SetCaption(r.instance, value)
}

// Checked
func (r *TRadioButton) Checked() bool {
    return RadioButton_GetChecked(r.instance)
}

// SetChecked
func (r *TRadioButton) SetChecked(value bool) {
    RadioButton_SetChecked(r.instance, value)
}

// Color
func (r *TRadioButton) Color() TColor {
    return RadioButton_GetColor(r.instance)
}

// SetColor
func (r *TRadioButton) SetColor(value TColor) {
    RadioButton_SetColor(r.instance, value)
}

// DoubleBuffered
func (r *TRadioButton) DoubleBuffered() bool {
    return RadioButton_GetDoubleBuffered(r.instance)
}

// SetDoubleBuffered
func (r *TRadioButton) SetDoubleBuffered(value bool) {
    RadioButton_SetDoubleBuffered(r.instance, value)
}

// DragCursor
func (r *TRadioButton) DragCursor() TCursor {
    return RadioButton_GetDragCursor(r.instance)
}

// SetDragCursor
func (r *TRadioButton) SetDragCursor(value TCursor) {
    RadioButton_SetDragCursor(r.instance, value)
}

// DragKind
func (r *TRadioButton) DragKind() TDragKind {
    return RadioButton_GetDragKind(r.instance)
}

// SetDragKind
func (r *TRadioButton) SetDragKind(value TDragKind) {
    RadioButton_SetDragKind(r.instance, value)
}

// DragMode
func (r *TRadioButton) DragMode() TDragMode {
    return RadioButton_GetDragMode(r.instance)
}

// SetDragMode
func (r *TRadioButton) SetDragMode(value TDragMode) {
    RadioButton_SetDragMode(r.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (r *TRadioButton) Enabled() bool {
    return RadioButton_GetEnabled(r.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (r *TRadioButton) SetEnabled(value bool) {
    RadioButton_SetEnabled(r.instance, value)
}

// Font
func (r *TRadioButton) Font() *TFont {
    return FontFromInst(RadioButton_GetFont(r.instance))
}

// SetFont
func (r *TRadioButton) SetFont(value *TFont) {
    RadioButton_SetFont(r.instance, CheckPtr(value))
}

// ParentColor
func (r *TRadioButton) ParentColor() bool {
    return RadioButton_GetParentColor(r.instance)
}

// SetParentColor
func (r *TRadioButton) SetParentColor(value bool) {
    RadioButton_SetParentColor(r.instance, value)
}

// ParentCtl3D
func (r *TRadioButton) ParentCtl3D() bool {
    return RadioButton_GetParentCtl3D(r.instance)
}

// SetParentCtl3D
func (r *TRadioButton) SetParentCtl3D(value bool) {
    RadioButton_SetParentCtl3D(r.instance, value)
}

// ParentDoubleBuffered
func (r *TRadioButton) ParentDoubleBuffered() bool {
    return RadioButton_GetParentDoubleBuffered(r.instance)
}

// SetParentDoubleBuffered
func (r *TRadioButton) SetParentDoubleBuffered(value bool) {
    RadioButton_SetParentDoubleBuffered(r.instance, value)
}

// ParentFont
func (r *TRadioButton) ParentFont() bool {
    return RadioButton_GetParentFont(r.instance)
}

// SetParentFont
func (r *TRadioButton) SetParentFont(value bool) {
    RadioButton_SetParentFont(r.instance, value)
}

// ParentShowHint
func (r *TRadioButton) ParentShowHint() bool {
    return RadioButton_GetParentShowHint(r.instance)
}

// SetParentShowHint
func (r *TRadioButton) SetParentShowHint(value bool) {
    RadioButton_SetParentShowHint(r.instance, value)
}

// PopupMenu
func (r *TRadioButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(RadioButton_GetPopupMenu(r.instance))
}

// SetPopupMenu
func (r *TRadioButton) SetPopupMenu(value IComponent) {
    RadioButton_SetPopupMenu(r.instance, CheckPtr(value))
}

// ShowHint
func (r *TRadioButton) ShowHint() bool {
    return RadioButton_GetShowHint(r.instance)
}

// SetShowHint
func (r *TRadioButton) SetShowHint(value bool) {
    RadioButton_SetShowHint(r.instance, value)
}

// TabOrder
func (r *TRadioButton) TabOrder() uint16 {
    return RadioButton_GetTabOrder(r.instance)
}

// SetTabOrder
func (r *TRadioButton) SetTabOrder(value uint16) {
    RadioButton_SetTabOrder(r.instance, value)
}

// TabStop
func (r *TRadioButton) TabStop() bool {
    return RadioButton_GetTabStop(r.instance)
}

// SetTabStop
func (r *TRadioButton) SetTabStop(value bool) {
    RadioButton_SetTabStop(r.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (r *TRadioButton) Visible() bool {
    return RadioButton_GetVisible(r.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (r *TRadioButton) SetVisible(value bool) {
    RadioButton_SetVisible(r.instance, value)
}

// WordWrap
func (r *TRadioButton) WordWrap() bool {
    return RadioButton_GetWordWrap(r.instance)
}

// SetWordWrap
func (r *TRadioButton) SetWordWrap(value bool) {
    RadioButton_SetWordWrap(r.instance, value)
}

// StyleElements
func (r *TRadioButton) StyleElements() TStyleElements {
    return RadioButton_GetStyleElements(r.instance)
}

// SetStyleElements
func (r *TRadioButton) SetStyleElements(value TStyleElements) {
    RadioButton_SetStyleElements(r.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (r *TRadioButton) SetOnClick(fn TNotifyEvent) {
    RadioButton_SetOnClick(r.instance, fn)
}

// SetOnContextPopup
func (r *TRadioButton) SetOnContextPopup(fn TContextPopupEvent) {
    RadioButton_SetOnContextPopup(r.instance, fn)
}

// SetOnDblClick
func (r *TRadioButton) SetOnDblClick(fn TNotifyEvent) {
    RadioButton_SetOnDblClick(r.instance, fn)
}

// SetOnDragDrop
func (r *TRadioButton) SetOnDragDrop(fn TDragDropEvent) {
    RadioButton_SetOnDragDrop(r.instance, fn)
}

// SetOnDragOver
func (r *TRadioButton) SetOnDragOver(fn TDragOverEvent) {
    RadioButton_SetOnDragOver(r.instance, fn)
}

// SetOnEndDock
func (r *TRadioButton) SetOnEndDock(fn TEndDragEvent) {
    RadioButton_SetOnEndDock(r.instance, fn)
}

// SetOnEndDrag
func (r *TRadioButton) SetOnEndDrag(fn TEndDragEvent) {
    RadioButton_SetOnEndDrag(r.instance, fn)
}

// SetOnEnter
func (r *TRadioButton) SetOnEnter(fn TNotifyEvent) {
    RadioButton_SetOnEnter(r.instance, fn)
}

// SetOnExit
func (r *TRadioButton) SetOnExit(fn TNotifyEvent) {
    RadioButton_SetOnExit(r.instance, fn)
}

// SetOnKeyDown
func (r *TRadioButton) SetOnKeyDown(fn TKeyEvent) {
    RadioButton_SetOnKeyDown(r.instance, fn)
}

// SetOnKeyPress
func (r *TRadioButton) SetOnKeyPress(fn TKeyPressEvent) {
    RadioButton_SetOnKeyPress(r.instance, fn)
}

// SetOnKeyUp
func (r *TRadioButton) SetOnKeyUp(fn TKeyEvent) {
    RadioButton_SetOnKeyUp(r.instance, fn)
}

// SetOnMouseDown
func (r *TRadioButton) SetOnMouseDown(fn TMouseEvent) {
    RadioButton_SetOnMouseDown(r.instance, fn)
}

// SetOnMouseEnter
func (r *TRadioButton) SetOnMouseEnter(fn TNotifyEvent) {
    RadioButton_SetOnMouseEnter(r.instance, fn)
}

// SetOnMouseLeave
func (r *TRadioButton) SetOnMouseLeave(fn TNotifyEvent) {
    RadioButton_SetOnMouseLeave(r.instance, fn)
}

// SetOnMouseMove
func (r *TRadioButton) SetOnMouseMove(fn TMouseMoveEvent) {
    RadioButton_SetOnMouseMove(r.instance, fn)
}

// SetOnMouseUp
func (r *TRadioButton) SetOnMouseUp(fn TMouseEvent) {
    RadioButton_SetOnMouseUp(r.instance, fn)
}

// SetOnStartDock
func (r *TRadioButton) SetOnStartDock(fn TStartDockEvent) {
    RadioButton_SetOnStartDock(r.instance, fn)
}

// DockSite
func (r *TRadioButton) DockSite() bool {
    return RadioButton_GetDockSite(r.instance)
}

// SetDockSite
func (r *TRadioButton) SetDockSite(value bool) {
    RadioButton_SetDockSite(r.instance, value)
}

// Brush
func (r *TRadioButton) Brush() *TBrush {
    return BrushFromInst(RadioButton_GetBrush(r.instance))
}

// ControlCount
func (r *TRadioButton) ControlCount() int32 {
    return RadioButton_GetControlCount(r.instance)
}

// Handle
func (r *TRadioButton) Handle() HWND {
    return RadioButton_GetHandle(r.instance)
}

// ParentWindow
func (r *TRadioButton) ParentWindow() HWND {
    return RadioButton_GetParentWindow(r.instance)
}

// SetParentWindow
func (r *TRadioButton) SetParentWindow(value HWND) {
    RadioButton_SetParentWindow(r.instance, value)
}

// UseDockManager
func (r *TRadioButton) UseDockManager() bool {
    return RadioButton_GetUseDockManager(r.instance)
}

// SetUseDockManager
func (r *TRadioButton) SetUseDockManager(value bool) {
    RadioButton_SetUseDockManager(r.instance, value)
}

// BoundsRect
func (r *TRadioButton) BoundsRect() TRect {
    return RadioButton_GetBoundsRect(r.instance)
}

// SetBoundsRect
func (r *TRadioButton) SetBoundsRect(value TRect) {
    RadioButton_SetBoundsRect(r.instance, value)
}

// ClientHeight
func (r *TRadioButton) ClientHeight() int32 {
    return RadioButton_GetClientHeight(r.instance)
}

// SetClientHeight
func (r *TRadioButton) SetClientHeight(value int32) {
    RadioButton_SetClientHeight(r.instance, value)
}

// ClientRect
func (r *TRadioButton) ClientRect() TRect {
    return RadioButton_GetClientRect(r.instance)
}

// ClientWidth
func (r *TRadioButton) ClientWidth() int32 {
    return RadioButton_GetClientWidth(r.instance)
}

// SetClientWidth
func (r *TRadioButton) SetClientWidth(value int32) {
    RadioButton_SetClientWidth(r.instance, value)
}

// ExplicitLeft
func (r *TRadioButton) ExplicitLeft() int32 {
    return RadioButton_GetExplicitLeft(r.instance)
}

// ExplicitTop
func (r *TRadioButton) ExplicitTop() int32 {
    return RadioButton_GetExplicitTop(r.instance)
}

// ExplicitWidth
func (r *TRadioButton) ExplicitWidth() int32 {
    return RadioButton_GetExplicitWidth(r.instance)
}

// ExplicitHeight
func (r *TRadioButton) ExplicitHeight() int32 {
    return RadioButton_GetExplicitHeight(r.instance)
}

// Floating
func (r *TRadioButton) Floating() bool {
    return RadioButton_GetFloating(r.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (r *TRadioButton) Parent() *TWinControl {
    return WinControlFromInst(RadioButton_GetParent(r.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (r *TRadioButton) SetParent(value IWinControl) {
    RadioButton_SetParent(r.instance, CheckPtr(value))
}

// AlignWithMargins
func (r *TRadioButton) AlignWithMargins() bool {
    return RadioButton_GetAlignWithMargins(r.instance)
}

// SetAlignWithMargins
func (r *TRadioButton) SetAlignWithMargins(value bool) {
    RadioButton_SetAlignWithMargins(r.instance, value)
}

// Left
func (r *TRadioButton) Left() int32 {
    return RadioButton_GetLeft(r.instance)
}

// SetLeft
func (r *TRadioButton) SetLeft(value int32) {
    RadioButton_SetLeft(r.instance, value)
}

// Top
func (r *TRadioButton) Top() int32 {
    return RadioButton_GetTop(r.instance)
}

// SetTop
func (r *TRadioButton) SetTop(value int32) {
    RadioButton_SetTop(r.instance, value)
}

// Width
func (r *TRadioButton) Width() int32 {
    return RadioButton_GetWidth(r.instance)
}

// SetWidth
func (r *TRadioButton) SetWidth(value int32) {
    RadioButton_SetWidth(r.instance, value)
}

// Height
func (r *TRadioButton) Height() int32 {
    return RadioButton_GetHeight(r.instance)
}

// SetHeight
func (r *TRadioButton) SetHeight(value int32) {
    RadioButton_SetHeight(r.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (r *TRadioButton) Cursor() TCursor {
    return RadioButton_GetCursor(r.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (r *TRadioButton) SetCursor(value TCursor) {
    RadioButton_SetCursor(r.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (r *TRadioButton) Hint() string {
    return RadioButton_GetHint(r.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (r *TRadioButton) SetHint(value string) {
    RadioButton_SetHint(r.instance, value)
}

// Margins
func (r *TRadioButton) Margins() *TMargins {
    return MarginsFromInst(RadioButton_GetMargins(r.instance))
}

// SetMargins
func (r *TRadioButton) SetMargins(value *TMargins) {
    RadioButton_SetMargins(r.instance, CheckPtr(value))
}

// CustomHint
func (r *TRadioButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(RadioButton_GetCustomHint(r.instance))
}

// SetCustomHint
func (r *TRadioButton) SetCustomHint(value IComponent) {
    RadioButton_SetCustomHint(r.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (r *TRadioButton) ComponentCount() int32 {
    return RadioButton_GetComponentCount(r.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (r *TRadioButton) ComponentIndex() int32 {
    return RadioButton_GetComponentIndex(r.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (r *TRadioButton) SetComponentIndex(value int32) {
    RadioButton_SetComponentIndex(r.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (r *TRadioButton) Owner() *TComponent {
    return ComponentFromInst(RadioButton_GetOwner(r.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (r *TRadioButton) Name() string {
    return RadioButton_GetName(r.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (r *TRadioButton) SetName(value string) {
    RadioButton_SetName(r.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (r *TRadioButton) Tag() int {
    return RadioButton_GetTag(r.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (r *TRadioButton) SetTag(value int) {
    RadioButton_SetTag(r.instance, value)
}

// Controls
func (r *TRadioButton) Controls(Index int32) *TControl {
    return ControlFromInst(RadioButton_GetControls(r.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (r *TRadioButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(RadioButton_GetComponents(r.instance, AIndex))
}

