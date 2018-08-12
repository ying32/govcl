
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

type TCheckBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCheckBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCheckBox(owner IComponent) *TCheckBox {
    c := new(TCheckBox)
    c.instance = CheckBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CheckBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CheckBoxFromInst(inst uintptr) *TCheckBox {
    c := new(TCheckBox)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CheckBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CheckBoxFromObj(obj IObject) *TCheckBox {
    c := new(TCheckBox)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CheckBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CheckBoxFromUnsafePointer(ptr unsafe.Pointer) *TCheckBox {
    c := new(TCheckBox)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCheckBox) Free() {
    if c.instance != 0 {
        CheckBox_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCheckBox) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCheckBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCheckBox) IsValid() bool {
    return c.instance != 0
}

// TCheckBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCheckBoxClass() TClass {
    return CheckBox_StaticClassType()
}

// CanFocus
func (c *TCheckBox) CanFocus() bool {
    return CheckBox_CanFocus(c.instance)
}

// ContainsControl
func (c *TCheckBox) ContainsControl(Control IControl) bool {
    return CheckBox_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
func (c *TCheckBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(CheckBox_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (c *TCheckBox) DisableAlign() {
    CheckBox_DisableAlign(c.instance)
}

// EnableAlign
func (c *TCheckBox) EnableAlign() {
    CheckBox_EnableAlign(c.instance)
}

// FindChildControl
func (c *TCheckBox) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(CheckBox_FindChildControl(c.instance, ControlName))
}

// FlipChildren
func (c *TCheckBox) FlipChildren(AllLevels bool) {
    CheckBox_FlipChildren(c.instance, AllLevels)
}

// Focused
func (c *TCheckBox) Focused() bool {
    return CheckBox_Focused(c.instance)
}

// HandleAllocated
func (c *TCheckBox) HandleAllocated() bool {
    return CheckBox_HandleAllocated(c.instance)
}

// InsertControl
func (c *TCheckBox) InsertControl(AControl IControl) {
    CheckBox_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
func (c *TCheckBox) Invalidate() {
    CheckBox_Invalidate(c.instance)
}

// PaintTo
func (c *TCheckBox) PaintTo(DC HDC, X int32, Y int32) {
    CheckBox_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
func (c *TCheckBox) RemoveControl(AControl IControl) {
    CheckBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
func (c *TCheckBox) Realign() {
    CheckBox_Realign(c.instance)
}

// Repaint
func (c *TCheckBox) Repaint() {
    CheckBox_Repaint(c.instance)
}

// ScaleBy
func (c *TCheckBox) ScaleBy(M int32, D int32) {
    CheckBox_ScaleBy(c.instance, M , D)
}

// ScrollBy
func (c *TCheckBox) ScrollBy(DeltaX int32, DeltaY int32) {
    CheckBox_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetBounds
func (c *TCheckBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CheckBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (c *TCheckBox) SetFocus() {
    CheckBox_SetFocus(c.instance)
}

// Update
func (c *TCheckBox) Update() {
    CheckBox_Update(c.instance)
}

// UpdateControlState
func (c *TCheckBox) UpdateControlState() {
    CheckBox_UpdateControlState(c.instance)
}

// BringToFront
func (c *TCheckBox) BringToFront() {
    CheckBox_BringToFront(c.instance)
}

// ClientToScreen
func (c *TCheckBox) ClientToScreen(Point TPoint) TPoint {
    return CheckBox_ClientToScreen(c.instance, Point)
}

// ClientToParent
func (c *TCheckBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CheckBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
func (c *TCheckBox) Dragging() bool {
    return CheckBox_Dragging(c.instance)
}

// HasParent
func (c *TCheckBox) HasParent() bool {
    return CheckBox_HasParent(c.instance)
}

// Hide
func (c *TCheckBox) Hide() {
    CheckBox_Hide(c.instance)
}

// Perform
func (c *TCheckBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CheckBox_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
func (c *TCheckBox) Refresh() {
    CheckBox_Refresh(c.instance)
}

// ScreenToClient
func (c *TCheckBox) ScreenToClient(Point TPoint) TPoint {
    return CheckBox_ScreenToClient(c.instance, Point)
}

// ParentToClient
func (c *TCheckBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CheckBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (c *TCheckBox) SendToBack() {
    CheckBox_SendToBack(c.instance)
}

// Show
func (c *TCheckBox) Show() {
    CheckBox_Show(c.instance)
}

// GetTextBuf
func (c *TCheckBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CheckBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
func (c *TCheckBox) GetTextLen() int32 {
    return CheckBox_GetTextLen(c.instance)
}

// SetTextBuf
func (c *TCheckBox) SetTextBuf(Buffer string) {
    CheckBox_SetTextBuf(c.instance, Buffer)
}

// FindComponent
func (c *TCheckBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CheckBox_FindComponent(c.instance, AName))
}

// GetNamePath
func (c *TCheckBox) GetNamePath() string {
    return CheckBox_GetNamePath(c.instance)
}

// Assign
func (c *TCheckBox) Assign(Source IObject) {
    CheckBox_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCheckBox) DisposeOf() {
    CheckBox_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCheckBox) ClassType() TClass {
    return CheckBox_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCheckBox) ClassName() string {
    return CheckBox_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCheckBox) InstanceSize() int32 {
    return CheckBox_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCheckBox) InheritsFrom(AClass TClass) bool {
    return CheckBox_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCheckBox) Equals(Obj IObject) bool {
    return CheckBox_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCheckBox) GetHashCode() int32 {
    return CheckBox_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCheckBox) ToString() string {
    return CheckBox_ToString(c.instance)
}

// Action
func (c *TCheckBox) Action() *TAction {
    return ActionFromInst(CheckBox_GetAction(c.instance))
}

// SetAction
func (c *TCheckBox) SetAction(value IComponent) {
    CheckBox_SetAction(c.instance, CheckPtr(value))
}

// Align
func (c *TCheckBox) Align() TAlign {
    return CheckBox_GetAlign(c.instance)
}

// SetAlign
func (c *TCheckBox) SetAlign(value TAlign) {
    CheckBox_SetAlign(c.instance, value)
}

// Alignment
func (c *TCheckBox) Alignment() TLeftRight {
    return CheckBox_GetAlignment(c.instance)
}

// SetAlignment
func (c *TCheckBox) SetAlignment(value TLeftRight) {
    CheckBox_SetAlignment(c.instance, value)
}

// AllowGrayed
func (c *TCheckBox) AllowGrayed() bool {
    return CheckBox_GetAllowGrayed(c.instance)
}

// SetAllowGrayed
func (c *TCheckBox) SetAllowGrayed(value bool) {
    CheckBox_SetAllowGrayed(c.instance, value)
}

// Anchors
func (c *TCheckBox) Anchors() TAnchors {
    return CheckBox_GetAnchors(c.instance)
}

// SetAnchors
func (c *TCheckBox) SetAnchors(value TAnchors) {
    CheckBox_SetAnchors(c.instance, value)
}

// BiDiMode
func (c *TCheckBox) BiDiMode() TBiDiMode {
    return CheckBox_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TCheckBox) SetBiDiMode(value TBiDiMode) {
    CheckBox_SetBiDiMode(c.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (c *TCheckBox) Caption() string {
    return CheckBox_GetCaption(c.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (c *TCheckBox) SetCaption(value string) {
    CheckBox_SetCaption(c.instance, value)
}

// Checked
func (c *TCheckBox) Checked() bool {
    return CheckBox_GetChecked(c.instance)
}

// SetChecked
func (c *TCheckBox) SetChecked(value bool) {
    CheckBox_SetChecked(c.instance, value)
}

// Color
func (c *TCheckBox) Color() TColor {
    return CheckBox_GetColor(c.instance)
}

// SetColor
func (c *TCheckBox) SetColor(value TColor) {
    CheckBox_SetColor(c.instance, value)
}

// DoubleBuffered
func (c *TCheckBox) DoubleBuffered() bool {
    return CheckBox_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
func (c *TCheckBox) SetDoubleBuffered(value bool) {
    CheckBox_SetDoubleBuffered(c.instance, value)
}

// DragCursor
func (c *TCheckBox) DragCursor() TCursor {
    return CheckBox_GetDragCursor(c.instance)
}

// SetDragCursor
func (c *TCheckBox) SetDragCursor(value TCursor) {
    CheckBox_SetDragCursor(c.instance, value)
}

// DragKind
func (c *TCheckBox) DragKind() TDragKind {
    return CheckBox_GetDragKind(c.instance)
}

// SetDragKind
func (c *TCheckBox) SetDragKind(value TDragKind) {
    CheckBox_SetDragKind(c.instance, value)
}

// DragMode
func (c *TCheckBox) DragMode() TDragMode {
    return CheckBox_GetDragMode(c.instance)
}

// SetDragMode
func (c *TCheckBox) SetDragMode(value TDragMode) {
    CheckBox_SetDragMode(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TCheckBox) Enabled() bool {
    return CheckBox_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TCheckBox) SetEnabled(value bool) {
    CheckBox_SetEnabled(c.instance, value)
}

// Font
func (c *TCheckBox) Font() *TFont {
    return FontFromInst(CheckBox_GetFont(c.instance))
}

// SetFont
func (c *TCheckBox) SetFont(value *TFont) {
    CheckBox_SetFont(c.instance, CheckPtr(value))
}

// ParentColor
func (c *TCheckBox) ParentColor() bool {
    return CheckBox_GetParentColor(c.instance)
}

// SetParentColor
func (c *TCheckBox) SetParentColor(value bool) {
    CheckBox_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TCheckBox) ParentCtl3D() bool {
    return CheckBox_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TCheckBox) SetParentCtl3D(value bool) {
    CheckBox_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
func (c *TCheckBox) ParentDoubleBuffered() bool {
    return CheckBox_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
func (c *TCheckBox) SetParentDoubleBuffered(value bool) {
    CheckBox_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
func (c *TCheckBox) ParentFont() bool {
    return CheckBox_GetParentFont(c.instance)
}

// SetParentFont
func (c *TCheckBox) SetParentFont(value bool) {
    CheckBox_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TCheckBox) ParentShowHint() bool {
    return CheckBox_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TCheckBox) SetParentShowHint(value bool) {
    CheckBox_SetParentShowHint(c.instance, value)
}

// PopupMenu
func (c *TCheckBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CheckBox_GetPopupMenu(c.instance))
}

// SetPopupMenu
func (c *TCheckBox) SetPopupMenu(value IComponent) {
    CheckBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
func (c *TCheckBox) ShowHint() bool {
    return CheckBox_GetShowHint(c.instance)
}

// SetShowHint
func (c *TCheckBox) SetShowHint(value bool) {
    CheckBox_SetShowHint(c.instance, value)
}

// State
func (c *TCheckBox) State() TCheckBoxState {
    return CheckBox_GetState(c.instance)
}

// SetState
func (c *TCheckBox) SetState(value TCheckBoxState) {
    CheckBox_SetState(c.instance, value)
}

// TabOrder
func (c *TCheckBox) TabOrder() TTabOrder {
    return CheckBox_GetTabOrder(c.instance)
}

// SetTabOrder
func (c *TCheckBox) SetTabOrder(value TTabOrder) {
    CheckBox_SetTabOrder(c.instance, value)
}

// TabStop
func (c *TCheckBox) TabStop() bool {
    return CheckBox_GetTabStop(c.instance)
}

// SetTabStop
func (c *TCheckBox) SetTabStop(value bool) {
    CheckBox_SetTabStop(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TCheckBox) Visible() bool {
    return CheckBox_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TCheckBox) SetVisible(value bool) {
    CheckBox_SetVisible(c.instance, value)
}

// WordWrap
func (c *TCheckBox) WordWrap() bool {
    return CheckBox_GetWordWrap(c.instance)
}

// SetWordWrap
func (c *TCheckBox) SetWordWrap(value bool) {
    CheckBox_SetWordWrap(c.instance, value)
}

// StyleElements
func (c *TCheckBox) StyleElements() TStyleElements {
    return CheckBox_GetStyleElements(c.instance)
}

// SetStyleElements
func (c *TCheckBox) SetStyleElements(value TStyleElements) {
    CheckBox_SetStyleElements(c.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TCheckBox) SetOnClick(fn TNotifyEvent) {
    CheckBox_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
func (c *TCheckBox) SetOnContextPopup(fn TContextPopupEvent) {
    CheckBox_SetOnContextPopup(c.instance, fn)
}

// SetOnDragDrop
func (c *TCheckBox) SetOnDragDrop(fn TDragDropEvent) {
    CheckBox_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
func (c *TCheckBox) SetOnDragOver(fn TDragOverEvent) {
    CheckBox_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
func (c *TCheckBox) SetOnEndDock(fn TEndDragEvent) {
    CheckBox_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
func (c *TCheckBox) SetOnEndDrag(fn TEndDragEvent) {
    CheckBox_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
func (c *TCheckBox) SetOnEnter(fn TNotifyEvent) {
    CheckBox_SetOnEnter(c.instance, fn)
}

// SetOnExit
func (c *TCheckBox) SetOnExit(fn TNotifyEvent) {
    CheckBox_SetOnExit(c.instance, fn)
}

// SetOnKeyDown
func (c *TCheckBox) SetOnKeyDown(fn TKeyEvent) {
    CheckBox_SetOnKeyDown(c.instance, fn)
}

// SetOnKeyPress
func (c *TCheckBox) SetOnKeyPress(fn TKeyPressEvent) {
    CheckBox_SetOnKeyPress(c.instance, fn)
}

// SetOnKeyUp
func (c *TCheckBox) SetOnKeyUp(fn TKeyEvent) {
    CheckBox_SetOnKeyUp(c.instance, fn)
}

// SetOnMouseDown
func (c *TCheckBox) SetOnMouseDown(fn TMouseEvent) {
    CheckBox_SetOnMouseDown(c.instance, fn)
}

// SetOnMouseEnter
func (c *TCheckBox) SetOnMouseEnter(fn TNotifyEvent) {
    CheckBox_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
func (c *TCheckBox) SetOnMouseLeave(fn TNotifyEvent) {
    CheckBox_SetOnMouseLeave(c.instance, fn)
}

// SetOnMouseMove
func (c *TCheckBox) SetOnMouseMove(fn TMouseMoveEvent) {
    CheckBox_SetOnMouseMove(c.instance, fn)
}

// SetOnMouseUp
func (c *TCheckBox) SetOnMouseUp(fn TMouseEvent) {
    CheckBox_SetOnMouseUp(c.instance, fn)
}

// SetOnStartDock
func (c *TCheckBox) SetOnStartDock(fn TStartDockEvent) {
    CheckBox_SetOnStartDock(c.instance, fn)
}

// DockClientCount
func (c *TCheckBox) DockClientCount() int32 {
    return CheckBox_GetDockClientCount(c.instance)
}

// DockSite
func (c *TCheckBox) DockSite() bool {
    return CheckBox_GetDockSite(c.instance)
}

// SetDockSite
func (c *TCheckBox) SetDockSite(value bool) {
    CheckBox_SetDockSite(c.instance, value)
}

// AlignDisabled
func (c *TCheckBox) AlignDisabled() bool {
    return CheckBox_GetAlignDisabled(c.instance)
}

// MouseInClient
func (c *TCheckBox) MouseInClient() bool {
    return CheckBox_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
func (c *TCheckBox) VisibleDockClientCount() int32 {
    return CheckBox_GetVisibleDockClientCount(c.instance)
}

// Brush
func (c *TCheckBox) Brush() *TBrush {
    return BrushFromInst(CheckBox_GetBrush(c.instance))
}

// ControlCount
func (c *TCheckBox) ControlCount() int32 {
    return CheckBox_GetControlCount(c.instance)
}

// Handle
func (c *TCheckBox) Handle() HWND {
    return CheckBox_GetHandle(c.instance)
}

// ParentWindow
func (c *TCheckBox) ParentWindow() HWND {
    return CheckBox_GetParentWindow(c.instance)
}

// SetParentWindow
func (c *TCheckBox) SetParentWindow(value HWND) {
    CheckBox_SetParentWindow(c.instance, value)
}

// UseDockManager
func (c *TCheckBox) UseDockManager() bool {
    return CheckBox_GetUseDockManager(c.instance)
}

// SetUseDockManager
func (c *TCheckBox) SetUseDockManager(value bool) {
    CheckBox_SetUseDockManager(c.instance, value)
}

// BoundsRect
func (c *TCheckBox) BoundsRect() TRect {
    return CheckBox_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TCheckBox) SetBoundsRect(value TRect) {
    CheckBox_SetBoundsRect(c.instance, value)
}

// ClientHeight
func (c *TCheckBox) ClientHeight() int32 {
    return CheckBox_GetClientHeight(c.instance)
}

// SetClientHeight
func (c *TCheckBox) SetClientHeight(value int32) {
    CheckBox_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TCheckBox) ClientOrigin() TPoint {
    return CheckBox_GetClientOrigin(c.instance)
}

// ClientRect
func (c *TCheckBox) ClientRect() TRect {
    return CheckBox_GetClientRect(c.instance)
}

// ClientWidth
func (c *TCheckBox) ClientWidth() int32 {
    return CheckBox_GetClientWidth(c.instance)
}

// SetClientWidth
func (c *TCheckBox) SetClientWidth(value int32) {
    CheckBox_SetClientWidth(c.instance, value)
}

// ControlState
func (c *TCheckBox) ControlState() TControlState {
    return CheckBox_GetControlState(c.instance)
}

// SetControlState
func (c *TCheckBox) SetControlState(value TControlState) {
    CheckBox_SetControlState(c.instance, value)
}

// ControlStyle
func (c *TCheckBox) ControlStyle() TControlStyle {
    return CheckBox_GetControlStyle(c.instance)
}

// SetControlStyle
func (c *TCheckBox) SetControlStyle(value TControlStyle) {
    CheckBox_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TCheckBox) ExplicitLeft() int32 {
    return CheckBox_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TCheckBox) ExplicitTop() int32 {
    return CheckBox_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TCheckBox) ExplicitWidth() int32 {
    return CheckBox_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TCheckBox) ExplicitHeight() int32 {
    return CheckBox_GetExplicitHeight(c.instance)
}

// Floating
func (c *TCheckBox) Floating() bool {
    return CheckBox_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TCheckBox) Parent() *TWinControl {
    return WinControlFromInst(CheckBox_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TCheckBox) SetParent(value IWinControl) {
    CheckBox_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
func (c *TCheckBox) AlignWithMargins() bool {
    return CheckBox_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
func (c *TCheckBox) SetAlignWithMargins(value bool) {
    CheckBox_SetAlignWithMargins(c.instance, value)
}

// Left
func (c *TCheckBox) Left() int32 {
    return CheckBox_GetLeft(c.instance)
}

// SetLeft
func (c *TCheckBox) SetLeft(value int32) {
    CheckBox_SetLeft(c.instance, value)
}

// Top
func (c *TCheckBox) Top() int32 {
    return CheckBox_GetTop(c.instance)
}

// SetTop
func (c *TCheckBox) SetTop(value int32) {
    CheckBox_SetTop(c.instance, value)
}

// Width
func (c *TCheckBox) Width() int32 {
    return CheckBox_GetWidth(c.instance)
}

// SetWidth
func (c *TCheckBox) SetWidth(value int32) {
    CheckBox_SetWidth(c.instance, value)
}

// Height
func (c *TCheckBox) Height() int32 {
    return CheckBox_GetHeight(c.instance)
}

// SetHeight
func (c *TCheckBox) SetHeight(value int32) {
    CheckBox_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TCheckBox) Cursor() TCursor {
    return CheckBox_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TCheckBox) SetCursor(value TCursor) {
    CheckBox_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (c *TCheckBox) Hint() string {
    return CheckBox_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (c *TCheckBox) SetHint(value string) {
    CheckBox_SetHint(c.instance, value)
}

// Margins
func (c *TCheckBox) Margins() *TMargins {
    return MarginsFromInst(CheckBox_GetMargins(c.instance))
}

// SetMargins
func (c *TCheckBox) SetMargins(value *TMargins) {
    CheckBox_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
func (c *TCheckBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(CheckBox_GetCustomHint(c.instance))
}

// SetCustomHint
func (c *TCheckBox) SetCustomHint(value IComponent) {
    CheckBox_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TCheckBox) ComponentCount() int32 {
    return CheckBox_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TCheckBox) ComponentIndex() int32 {
    return CheckBox_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TCheckBox) SetComponentIndex(value int32) {
    CheckBox_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TCheckBox) Owner() *TComponent {
    return ComponentFromInst(CheckBox_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TCheckBox) Name() string {
    return CheckBox_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TCheckBox) SetName(value string) {
    CheckBox_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TCheckBox) Tag() int {
    return CheckBox_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TCheckBox) SetTag(value int) {
    CheckBox_SetTag(c.instance, value)
}

// DockClients
func (c *TCheckBox) DockClients(Index int32) *TControl {
    return ControlFromInst(CheckBox_GetDockClients(c.instance, Index))
}

// Controls
func (c *TCheckBox) Controls(Index int32) *TControl {
    return ControlFromInst(CheckBox_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TCheckBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CheckBox_GetComponents(c.instance, AIndex))
}

