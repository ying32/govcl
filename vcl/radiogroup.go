
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

type TRadioGroup struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewRadioGroup
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewRadioGroup(owner IComponent) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = RadioGroup_Create(CheckPtr(owner))
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RadioGroupFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func RadioGroupFromInst(inst uintptr) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = inst
    r.ptr = unsafe.Pointer(inst)
    return r
}

// RadioGroupFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func RadioGroupFromObj(obj IObject) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = CheckPtr(obj)
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RadioGroupFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func RadioGroupFromUnsafePointer(ptr unsafe.Pointer) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = uintptr(ptr)
    r.ptr = ptr
    return r
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (r *TRadioGroup) Free() {
    if r.instance != 0 {
        RadioGroup_Free(r.instance)
        r.instance = 0
        r.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (r *TRadioGroup) Instance() uintptr {
    return r.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (r *TRadioGroup) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (r *TRadioGroup) IsValid() bool {
    return r.instance != 0
}

// TRadioGroupClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TRadioGroupClass() TClass {
    return RadioGroup_StaticClassType()
}

// FlipChildren
func (r *TRadioGroup) FlipChildren(AllLevels bool) {
    RadioGroup_FlipChildren(r.instance, AllLevels)
}

// CanFocus
func (r *TRadioGroup) CanFocus() bool {
    return RadioGroup_CanFocus(r.instance)
}

// Focused
func (r *TRadioGroup) Focused() bool {
    return RadioGroup_Focused(r.instance)
}

// HandleAllocated
func (r *TRadioGroup) HandleAllocated() bool {
    return RadioGroup_HandleAllocated(r.instance)
}

// Invalidate
func (r *TRadioGroup) Invalidate() {
    RadioGroup_Invalidate(r.instance)
}

// Realign
func (r *TRadioGroup) Realign() {
    RadioGroup_Realign(r.instance)
}

// Repaint
func (r *TRadioGroup) Repaint() {
    RadioGroup_Repaint(r.instance)
}

// ScaleBy
func (r *TRadioGroup) ScaleBy(M int32, D int32) {
    RadioGroup_ScaleBy(r.instance, M , D)
}

// SetBounds
func (r *TRadioGroup) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RadioGroup_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (r *TRadioGroup) SetFocus() {
    RadioGroup_SetFocus(r.instance)
}

// Update
func (r *TRadioGroup) Update() {
    RadioGroup_Update(r.instance)
}

// BringToFront
func (r *TRadioGroup) BringToFront() {
    RadioGroup_BringToFront(r.instance)
}

// ClientToScreen
func (r *TRadioGroup) ClientToScreen(Point TPoint) TPoint {
    return RadioGroup_ClientToScreen(r.instance, Point)
}

// ClientToParent
func (r *TRadioGroup) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return RadioGroup_ClientToParent(r.instance, Point , CheckPtr(AParent))
}

// Dragging
func (r *TRadioGroup) Dragging() bool {
    return RadioGroup_Dragging(r.instance)
}

// HasParent
func (r *TRadioGroup) HasParent() bool {
    return RadioGroup_HasParent(r.instance)
}

// Hide
func (r *TRadioGroup) Hide() {
    RadioGroup_Hide(r.instance)
}

// Perform
func (r *TRadioGroup) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RadioGroup_Perform(r.instance, Msg , WParam , LParam)
}

// Refresh
func (r *TRadioGroup) Refresh() {
    RadioGroup_Refresh(r.instance)
}

// ScreenToClient
func (r *TRadioGroup) ScreenToClient(Point TPoint) TPoint {
    return RadioGroup_ScreenToClient(r.instance, Point)
}

// ParentToClient
func (r *TRadioGroup) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return RadioGroup_ParentToClient(r.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (r *TRadioGroup) SendToBack() {
    RadioGroup_SendToBack(r.instance)
}

// Show
func (r *TRadioGroup) Show() {
    RadioGroup_Show(r.instance)
}

// GetTextBuf
func (r *TRadioGroup) GetTextBuf(Buffer string, BufSize int32) int32 {
    return RadioGroup_GetTextBuf(r.instance, Buffer , BufSize)
}

// GetTextLen
func (r *TRadioGroup) GetTextLen() int32 {
    return RadioGroup_GetTextLen(r.instance)
}

// FindComponent
func (r *TRadioGroup) FindComponent(AName string) *TComponent {
    return ComponentFromInst(RadioGroup_FindComponent(r.instance, AName))
}

// GetNamePath
func (r *TRadioGroup) GetNamePath() string {
    return RadioGroup_GetNamePath(r.instance)
}

// Assign
func (r *TRadioGroup) Assign(Source IObject) {
    RadioGroup_Assign(r.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (r *TRadioGroup) DisposeOf() {
    RadioGroup_DisposeOf(r.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (r *TRadioGroup) ClassType() TClass {
    return RadioGroup_ClassType(r.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (r *TRadioGroup) ClassName() string {
    return RadioGroup_ClassName(r.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (r *TRadioGroup) InstanceSize() int32 {
    return RadioGroup_InstanceSize(r.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (r *TRadioGroup) InheritsFrom(AClass TClass) bool {
    return RadioGroup_InheritsFrom(r.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (r *TRadioGroup) Equals(Obj IObject) bool {
    return RadioGroup_Equals(r.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (r *TRadioGroup) GetHashCode() int32 {
    return RadioGroup_GetHashCode(r.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (r *TRadioGroup) ToString() string {
    return RadioGroup_ToString(r.instance)
}

// Align
func (r *TRadioGroup) Align() TAlign {
    return RadioGroup_GetAlign(r.instance)
}

// SetAlign
func (r *TRadioGroup) SetAlign(value TAlign) {
    RadioGroup_SetAlign(r.instance, value)
}

// Anchors
func (r *TRadioGroup) Anchors() TAnchors {
    return RadioGroup_GetAnchors(r.instance)
}

// SetAnchors
func (r *TRadioGroup) SetAnchors(value TAnchors) {
    RadioGroup_SetAnchors(r.instance, value)
}

// BiDiMode
func (r *TRadioGroup) BiDiMode() TBiDiMode {
    return RadioGroup_GetBiDiMode(r.instance)
}

// SetBiDiMode
func (r *TRadioGroup) SetBiDiMode(value TBiDiMode) {
    RadioGroup_SetBiDiMode(r.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (r *TRadioGroup) Caption() string {
    return RadioGroup_GetCaption(r.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (r *TRadioGroup) SetCaption(value string) {
    RadioGroup_SetCaption(r.instance, value)
}

// Color
func (r *TRadioGroup) Color() TColor {
    return RadioGroup_GetColor(r.instance)
}

// SetColor
func (r *TRadioGroup) SetColor(value TColor) {
    RadioGroup_SetColor(r.instance, value)
}

// Columns
func (r *TRadioGroup) Columns() int32 {
    return RadioGroup_GetColumns(r.instance)
}

// SetColumns
func (r *TRadioGroup) SetColumns(value int32) {
    RadioGroup_SetColumns(r.instance, value)
}

// DoubleBuffered
func (r *TRadioGroup) DoubleBuffered() bool {
    return RadioGroup_GetDoubleBuffered(r.instance)
}

// SetDoubleBuffered
func (r *TRadioGroup) SetDoubleBuffered(value bool) {
    RadioGroup_SetDoubleBuffered(r.instance, value)
}

// DragCursor
func (r *TRadioGroup) DragCursor() TCursor {
    return RadioGroup_GetDragCursor(r.instance)
}

// SetDragCursor
func (r *TRadioGroup) SetDragCursor(value TCursor) {
    RadioGroup_SetDragCursor(r.instance, value)
}

// DragKind
func (r *TRadioGroup) DragKind() TDragKind {
    return RadioGroup_GetDragKind(r.instance)
}

// SetDragKind
func (r *TRadioGroup) SetDragKind(value TDragKind) {
    RadioGroup_SetDragKind(r.instance, value)
}

// DragMode
func (r *TRadioGroup) DragMode() TDragMode {
    return RadioGroup_GetDragMode(r.instance)
}

// SetDragMode
func (r *TRadioGroup) SetDragMode(value TDragMode) {
    RadioGroup_SetDragMode(r.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (r *TRadioGroup) Enabled() bool {
    return RadioGroup_GetEnabled(r.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (r *TRadioGroup) SetEnabled(value bool) {
    RadioGroup_SetEnabled(r.instance, value)
}

// Font
func (r *TRadioGroup) Font() *TFont {
    return FontFromInst(RadioGroup_GetFont(r.instance))
}

// SetFont
func (r *TRadioGroup) SetFont(value *TFont) {
    RadioGroup_SetFont(r.instance, CheckPtr(value))
}

// ItemIndex
func (r *TRadioGroup) ItemIndex() int32 {
    return RadioGroup_GetItemIndex(r.instance)
}

// SetItemIndex
func (r *TRadioGroup) SetItemIndex(value int32) {
    RadioGroup_SetItemIndex(r.instance, value)
}

// Items
func (r *TRadioGroup) Items() *TStrings {
    return StringsFromInst(RadioGroup_GetItems(r.instance))
}

// SetItems
func (r *TRadioGroup) SetItems(value IObject) {
    RadioGroup_SetItems(r.instance, CheckPtr(value))
}

// ParentBackground
func (r *TRadioGroup) ParentBackground() bool {
    return RadioGroup_GetParentBackground(r.instance)
}

// SetParentBackground
func (r *TRadioGroup) SetParentBackground(value bool) {
    RadioGroup_SetParentBackground(r.instance, value)
}

// ParentColor
func (r *TRadioGroup) ParentColor() bool {
    return RadioGroup_GetParentColor(r.instance)
}

// SetParentColor
func (r *TRadioGroup) SetParentColor(value bool) {
    RadioGroup_SetParentColor(r.instance, value)
}

// ParentCtl3D
func (r *TRadioGroup) ParentCtl3D() bool {
    return RadioGroup_GetParentCtl3D(r.instance)
}

// SetParentCtl3D
func (r *TRadioGroup) SetParentCtl3D(value bool) {
    RadioGroup_SetParentCtl3D(r.instance, value)
}

// ParentDoubleBuffered
func (r *TRadioGroup) ParentDoubleBuffered() bool {
    return RadioGroup_GetParentDoubleBuffered(r.instance)
}

// SetParentDoubleBuffered
func (r *TRadioGroup) SetParentDoubleBuffered(value bool) {
    RadioGroup_SetParentDoubleBuffered(r.instance, value)
}

// ParentFont
func (r *TRadioGroup) ParentFont() bool {
    return RadioGroup_GetParentFont(r.instance)
}

// SetParentFont
func (r *TRadioGroup) SetParentFont(value bool) {
    RadioGroup_SetParentFont(r.instance, value)
}

// ParentShowHint
func (r *TRadioGroup) ParentShowHint() bool {
    return RadioGroup_GetParentShowHint(r.instance)
}

// SetParentShowHint
func (r *TRadioGroup) SetParentShowHint(value bool) {
    RadioGroup_SetParentShowHint(r.instance, value)
}

// PopupMenu
func (r *TRadioGroup) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(RadioGroup_GetPopupMenu(r.instance))
}

// SetPopupMenu
func (r *TRadioGroup) SetPopupMenu(value IComponent) {
    RadioGroup_SetPopupMenu(r.instance, CheckPtr(value))
}

// ShowHint
func (r *TRadioGroup) ShowHint() bool {
    return RadioGroup_GetShowHint(r.instance)
}

// SetShowHint
func (r *TRadioGroup) SetShowHint(value bool) {
    RadioGroup_SetShowHint(r.instance, value)
}

// TabOrder
func (r *TRadioGroup) TabOrder() uint16 {
    return RadioGroup_GetTabOrder(r.instance)
}

// SetTabOrder
func (r *TRadioGroup) SetTabOrder(value uint16) {
    RadioGroup_SetTabOrder(r.instance, value)
}

// TabStop
func (r *TRadioGroup) TabStop() bool {
    return RadioGroup_GetTabStop(r.instance)
}

// SetTabStop
func (r *TRadioGroup) SetTabStop(value bool) {
    RadioGroup_SetTabStop(r.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (r *TRadioGroup) Visible() bool {
    return RadioGroup_GetVisible(r.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (r *TRadioGroup) SetVisible(value bool) {
    RadioGroup_SetVisible(r.instance, value)
}

// StyleElements
func (r *TRadioGroup) StyleElements() TStyleElements {
    return RadioGroup_GetStyleElements(r.instance)
}

// SetStyleElements
func (r *TRadioGroup) SetStyleElements(value TStyleElements) {
    RadioGroup_SetStyleElements(r.instance, value)
}

// WordWrap
func (r *TRadioGroup) WordWrap() bool {
    return RadioGroup_GetWordWrap(r.instance)
}

// SetWordWrap
func (r *TRadioGroup) SetWordWrap(value bool) {
    RadioGroup_SetWordWrap(r.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (r *TRadioGroup) SetOnClick(fn TNotifyEvent) {
    RadioGroup_SetOnClick(r.instance, fn)
}

// SetOnContextPopup
func (r *TRadioGroup) SetOnContextPopup(fn TContextPopupEvent) {
    RadioGroup_SetOnContextPopup(r.instance, fn)
}

// SetOnDragDrop
func (r *TRadioGroup) SetOnDragDrop(fn TDragDropEvent) {
    RadioGroup_SetOnDragDrop(r.instance, fn)
}

// SetOnDragOver
func (r *TRadioGroup) SetOnDragOver(fn TDragOverEvent) {
    RadioGroup_SetOnDragOver(r.instance, fn)
}

// SetOnEndDock
func (r *TRadioGroup) SetOnEndDock(fn TEndDragEvent) {
    RadioGroup_SetOnEndDock(r.instance, fn)
}

// SetOnEndDrag
func (r *TRadioGroup) SetOnEndDrag(fn TEndDragEvent) {
    RadioGroup_SetOnEndDrag(r.instance, fn)
}

// SetOnEnter
func (r *TRadioGroup) SetOnEnter(fn TNotifyEvent) {
    RadioGroup_SetOnEnter(r.instance, fn)
}

// SetOnExit
func (r *TRadioGroup) SetOnExit(fn TNotifyEvent) {
    RadioGroup_SetOnExit(r.instance, fn)
}

// SetOnStartDock
func (r *TRadioGroup) SetOnStartDock(fn TStartDockEvent) {
    RadioGroup_SetOnStartDock(r.instance, fn)
}

// DockSite
func (r *TRadioGroup) DockSite() bool {
    return RadioGroup_GetDockSite(r.instance)
}

// SetDockSite
func (r *TRadioGroup) SetDockSite(value bool) {
    RadioGroup_SetDockSite(r.instance, value)
}

// Brush
func (r *TRadioGroup) Brush() *TBrush {
    return BrushFromInst(RadioGroup_GetBrush(r.instance))
}

// ControlCount
func (r *TRadioGroup) ControlCount() int32 {
    return RadioGroup_GetControlCount(r.instance)
}

// Handle
func (r *TRadioGroup) Handle() HWND {
    return RadioGroup_GetHandle(r.instance)
}

// ParentWindow
func (r *TRadioGroup) ParentWindow() HWND {
    return RadioGroup_GetParentWindow(r.instance)
}

// SetParentWindow
func (r *TRadioGroup) SetParentWindow(value HWND) {
    RadioGroup_SetParentWindow(r.instance, value)
}

// UseDockManager
func (r *TRadioGroup) UseDockManager() bool {
    return RadioGroup_GetUseDockManager(r.instance)
}

// SetUseDockManager
func (r *TRadioGroup) SetUseDockManager(value bool) {
    RadioGroup_SetUseDockManager(r.instance, value)
}

// Action
func (r *TRadioGroup) Action() *TAction {
    return ActionFromInst(RadioGroup_GetAction(r.instance))
}

// SetAction
func (r *TRadioGroup) SetAction(value IComponent) {
    RadioGroup_SetAction(r.instance, CheckPtr(value))
}

// BoundsRect
func (r *TRadioGroup) BoundsRect() TRect {
    return RadioGroup_GetBoundsRect(r.instance)
}

// SetBoundsRect
func (r *TRadioGroup) SetBoundsRect(value TRect) {
    RadioGroup_SetBoundsRect(r.instance, value)
}

// ClientHeight
func (r *TRadioGroup) ClientHeight() int32 {
    return RadioGroup_GetClientHeight(r.instance)
}

// SetClientHeight
func (r *TRadioGroup) SetClientHeight(value int32) {
    RadioGroup_SetClientHeight(r.instance, value)
}

// ClientRect
func (r *TRadioGroup) ClientRect() TRect {
    return RadioGroup_GetClientRect(r.instance)
}

// ClientWidth
func (r *TRadioGroup) ClientWidth() int32 {
    return RadioGroup_GetClientWidth(r.instance)
}

// SetClientWidth
func (r *TRadioGroup) SetClientWidth(value int32) {
    RadioGroup_SetClientWidth(r.instance, value)
}

// ExplicitLeft
func (r *TRadioGroup) ExplicitLeft() int32 {
    return RadioGroup_GetExplicitLeft(r.instance)
}

// ExplicitTop
func (r *TRadioGroup) ExplicitTop() int32 {
    return RadioGroup_GetExplicitTop(r.instance)
}

// ExplicitWidth
func (r *TRadioGroup) ExplicitWidth() int32 {
    return RadioGroup_GetExplicitWidth(r.instance)
}

// ExplicitHeight
func (r *TRadioGroup) ExplicitHeight() int32 {
    return RadioGroup_GetExplicitHeight(r.instance)
}

// Floating
func (r *TRadioGroup) Floating() bool {
    return RadioGroup_GetFloating(r.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (r *TRadioGroup) Parent() *TWinControl {
    return WinControlFromInst(RadioGroup_GetParent(r.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (r *TRadioGroup) SetParent(value IWinControl) {
    RadioGroup_SetParent(r.instance, CheckPtr(value))
}

// AlignWithMargins
func (r *TRadioGroup) AlignWithMargins() bool {
    return RadioGroup_GetAlignWithMargins(r.instance)
}

// SetAlignWithMargins
func (r *TRadioGroup) SetAlignWithMargins(value bool) {
    RadioGroup_SetAlignWithMargins(r.instance, value)
}

// Left
func (r *TRadioGroup) Left() int32 {
    return RadioGroup_GetLeft(r.instance)
}

// SetLeft
func (r *TRadioGroup) SetLeft(value int32) {
    RadioGroup_SetLeft(r.instance, value)
}

// Top
func (r *TRadioGroup) Top() int32 {
    return RadioGroup_GetTop(r.instance)
}

// SetTop
func (r *TRadioGroup) SetTop(value int32) {
    RadioGroup_SetTop(r.instance, value)
}

// Width
func (r *TRadioGroup) Width() int32 {
    return RadioGroup_GetWidth(r.instance)
}

// SetWidth
func (r *TRadioGroup) SetWidth(value int32) {
    RadioGroup_SetWidth(r.instance, value)
}

// Height
func (r *TRadioGroup) Height() int32 {
    return RadioGroup_GetHeight(r.instance)
}

// SetHeight
func (r *TRadioGroup) SetHeight(value int32) {
    RadioGroup_SetHeight(r.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (r *TRadioGroup) Cursor() TCursor {
    return RadioGroup_GetCursor(r.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (r *TRadioGroup) SetCursor(value TCursor) {
    RadioGroup_SetCursor(r.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (r *TRadioGroup) Hint() string {
    return RadioGroup_GetHint(r.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (r *TRadioGroup) SetHint(value string) {
    RadioGroup_SetHint(r.instance, value)
}

// Margins
func (r *TRadioGroup) Margins() *TMargins {
    return MarginsFromInst(RadioGroup_GetMargins(r.instance))
}

// SetMargins
func (r *TRadioGroup) SetMargins(value *TMargins) {
    RadioGroup_SetMargins(r.instance, CheckPtr(value))
}

// CustomHint
func (r *TRadioGroup) CustomHint() *TCustomHint {
    return CustomHintFromInst(RadioGroup_GetCustomHint(r.instance))
}

// SetCustomHint
func (r *TRadioGroup) SetCustomHint(value IComponent) {
    RadioGroup_SetCustomHint(r.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (r *TRadioGroup) ComponentCount() int32 {
    return RadioGroup_GetComponentCount(r.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (r *TRadioGroup) ComponentIndex() int32 {
    return RadioGroup_GetComponentIndex(r.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (r *TRadioGroup) SetComponentIndex(value int32) {
    RadioGroup_SetComponentIndex(r.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (r *TRadioGroup) Owner() *TComponent {
    return ComponentFromInst(RadioGroup_GetOwner(r.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (r *TRadioGroup) Name() string {
    return RadioGroup_GetName(r.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (r *TRadioGroup) SetName(value string) {
    RadioGroup_SetName(r.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (r *TRadioGroup) Tag() int {
    return RadioGroup_GetTag(r.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (r *TRadioGroup) SetTag(value int) {
    RadioGroup_SetTag(r.instance, value)
}

// Buttons
func (r *TRadioGroup) Buttons(Index int32) *TRadioButton {
    return RadioButtonFromInst(RadioGroup_GetButtons(r.instance, Index))
}

// Controls
func (r *TRadioGroup) Controls(Index int32) *TControl {
    return ControlFromInst(RadioGroup_GetControls(r.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (r *TRadioGroup) Components(AIndex int32) *TComponent {
    return ComponentFromInst(RadioGroup_GetComponents(r.instance, AIndex))
}

