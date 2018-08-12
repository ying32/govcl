
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

type TWinControl struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewWinControl
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewWinControl(owner IComponent) *TWinControl {
    w := new(TWinControl)
    w.instance = WinControl_Create(CheckPtr(owner))
    w.ptr = unsafe.Pointer(w.instance)
    return w
}

// WinControlFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func WinControlFromInst(inst uintptr) *TWinControl {
    w := new(TWinControl)
    w.instance = inst
    w.ptr = unsafe.Pointer(inst)
    return w
}

// WinControlFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func WinControlFromObj(obj IObject) *TWinControl {
    w := new(TWinControl)
    w.instance = CheckPtr(obj)
    w.ptr = unsafe.Pointer(w.instance)
    return w
}

// WinControlFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func WinControlFromUnsafePointer(ptr unsafe.Pointer) *TWinControl {
    w := new(TWinControl)
    w.instance = uintptr(ptr)
    w.ptr = ptr
    return w
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (w *TWinControl) Free() {
    if w.instance != 0 {
        WinControl_Free(w.instance)
        w.instance = 0
        w.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (w *TWinControl) Instance() uintptr {
    return w.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (w *TWinControl) UnsafeAddr() unsafe.Pointer {
    return w.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (w *TWinControl) IsValid() bool {
    return w.instance != 0
}

// TWinControlClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TWinControlClass() TClass {
    return WinControl_StaticClassType()
}

// CanFocus
func (w *TWinControl) CanFocus() bool {
    return WinControl_CanFocus(w.instance)
}

// ContainsControl
func (w *TWinControl) ContainsControl(Control IControl) bool {
    return WinControl_ContainsControl(w.instance, CheckPtr(Control))
}

// ControlAtPos
func (w *TWinControl) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(WinControl_ControlAtPos(w.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (w *TWinControl) DisableAlign() {
    WinControl_DisableAlign(w.instance)
}

// EnableAlign
func (w *TWinControl) EnableAlign() {
    WinControl_EnableAlign(w.instance)
}

// FindChildControl
func (w *TWinControl) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(WinControl_FindChildControl(w.instance, ControlName))
}

// FlipChildren
func (w *TWinControl) FlipChildren(AllLevels bool) {
    WinControl_FlipChildren(w.instance, AllLevels)
}

// Focused
func (w *TWinControl) Focused() bool {
    return WinControl_Focused(w.instance)
}

// HandleAllocated
func (w *TWinControl) HandleAllocated() bool {
    return WinControl_HandleAllocated(w.instance)
}

// InsertControl
func (w *TWinControl) InsertControl(AControl IControl) {
    WinControl_InsertControl(w.instance, CheckPtr(AControl))
}

// Invalidate
func (w *TWinControl) Invalidate() {
    WinControl_Invalidate(w.instance)
}

// PaintTo
func (w *TWinControl) PaintTo(DC HDC, X int32, Y int32) {
    WinControl_PaintTo(w.instance, DC , X , Y)
}

// RemoveControl
func (w *TWinControl) RemoveControl(AControl IControl) {
    WinControl_RemoveControl(w.instance, CheckPtr(AControl))
}

// Realign
func (w *TWinControl) Realign() {
    WinControl_Realign(w.instance)
}

// Repaint
func (w *TWinControl) Repaint() {
    WinControl_Repaint(w.instance)
}

// ScaleBy
func (w *TWinControl) ScaleBy(M int32, D int32) {
    WinControl_ScaleBy(w.instance, M , D)
}

// ScrollBy
func (w *TWinControl) ScrollBy(DeltaX int32, DeltaY int32) {
    WinControl_ScrollBy(w.instance, DeltaX , DeltaY)
}

// SetBounds
func (w *TWinControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    WinControl_SetBounds(w.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (w *TWinControl) SetFocus() {
    WinControl_SetFocus(w.instance)
}

// Update
func (w *TWinControl) Update() {
    WinControl_Update(w.instance)
}

// UpdateControlState
func (w *TWinControl) UpdateControlState() {
    WinControl_UpdateControlState(w.instance)
}

// BringToFront
func (w *TWinControl) BringToFront() {
    WinControl_BringToFront(w.instance)
}

// ClientToScreen
func (w *TWinControl) ClientToScreen(Point TPoint) TPoint {
    return WinControl_ClientToScreen(w.instance, Point)
}

// ClientToParent
func (w *TWinControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return WinControl_ClientToParent(w.instance, Point , CheckPtr(AParent))
}

// Dragging
func (w *TWinControl) Dragging() bool {
    return WinControl_Dragging(w.instance)
}

// HasParent
func (w *TWinControl) HasParent() bool {
    return WinControl_HasParent(w.instance)
}

// Hide
func (w *TWinControl) Hide() {
    WinControl_Hide(w.instance)
}

// Perform
func (w *TWinControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return WinControl_Perform(w.instance, Msg , WParam , LParam)
}

// Refresh
func (w *TWinControl) Refresh() {
    WinControl_Refresh(w.instance)
}

// ScreenToClient
func (w *TWinControl) ScreenToClient(Point TPoint) TPoint {
    return WinControl_ScreenToClient(w.instance, Point)
}

// ParentToClient
func (w *TWinControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return WinControl_ParentToClient(w.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (w *TWinControl) SendToBack() {
    WinControl_SendToBack(w.instance)
}

// Show
func (w *TWinControl) Show() {
    WinControl_Show(w.instance)
}

// GetTextBuf
func (w *TWinControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    return WinControl_GetTextBuf(w.instance, Buffer , BufSize)
}

// GetTextLen
func (w *TWinControl) GetTextLen() int32 {
    return WinControl_GetTextLen(w.instance)
}

// SetTextBuf
func (w *TWinControl) SetTextBuf(Buffer string) {
    WinControl_SetTextBuf(w.instance, Buffer)
}

// FindComponent
func (w *TWinControl) FindComponent(AName string) *TComponent {
    return ComponentFromInst(WinControl_FindComponent(w.instance, AName))
}

// GetNamePath
func (w *TWinControl) GetNamePath() string {
    return WinControl_GetNamePath(w.instance)
}

// Assign
func (w *TWinControl) Assign(Source IObject) {
    WinControl_Assign(w.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (w *TWinControl) DisposeOf() {
    WinControl_DisposeOf(w.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (w *TWinControl) ClassType() TClass {
    return WinControl_ClassType(w.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (w *TWinControl) ClassName() string {
    return WinControl_ClassName(w.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (w *TWinControl) InstanceSize() int32 {
    return WinControl_InstanceSize(w.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (w *TWinControl) InheritsFrom(AClass TClass) bool {
    return WinControl_InheritsFrom(w.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (w *TWinControl) Equals(Obj IObject) bool {
    return WinControl_Equals(w.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (w *TWinControl) GetHashCode() int32 {
    return WinControl_GetHashCode(w.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (w *TWinControl) ToString() string {
    return WinControl_ToString(w.instance)
}

// DockClientCount
func (w *TWinControl) DockClientCount() int32 {
    return WinControl_GetDockClientCount(w.instance)
}

// DockSite
func (w *TWinControl) DockSite() bool {
    return WinControl_GetDockSite(w.instance)
}

// SetDockSite
func (w *TWinControl) SetDockSite(value bool) {
    WinControl_SetDockSite(w.instance, value)
}

// DoubleBuffered
func (w *TWinControl) DoubleBuffered() bool {
    return WinControl_GetDoubleBuffered(w.instance)
}

// SetDoubleBuffered
func (w *TWinControl) SetDoubleBuffered(value bool) {
    WinControl_SetDoubleBuffered(w.instance, value)
}

// AlignDisabled
func (w *TWinControl) AlignDisabled() bool {
    return WinControl_GetAlignDisabled(w.instance)
}

// MouseInClient
func (w *TWinControl) MouseInClient() bool {
    return WinControl_GetMouseInClient(w.instance)
}

// VisibleDockClientCount
func (w *TWinControl) VisibleDockClientCount() int32 {
    return WinControl_GetVisibleDockClientCount(w.instance)
}

// Brush
func (w *TWinControl) Brush() *TBrush {
    return BrushFromInst(WinControl_GetBrush(w.instance))
}

// ControlCount
func (w *TWinControl) ControlCount() int32 {
    return WinControl_GetControlCount(w.instance)
}

// Handle
func (w *TWinControl) Handle() HWND {
    return WinControl_GetHandle(w.instance)
}

// ParentDoubleBuffered
func (w *TWinControl) ParentDoubleBuffered() bool {
    return WinControl_GetParentDoubleBuffered(w.instance)
}

// SetParentDoubleBuffered
func (w *TWinControl) SetParentDoubleBuffered(value bool) {
    WinControl_SetParentDoubleBuffered(w.instance, value)
}

// ParentWindow
func (w *TWinControl) ParentWindow() HWND {
    return WinControl_GetParentWindow(w.instance)
}

// SetParentWindow
func (w *TWinControl) SetParentWindow(value HWND) {
    WinControl_SetParentWindow(w.instance, value)
}

// TabOrder
func (w *TWinControl) TabOrder() TTabOrder {
    return WinControl_GetTabOrder(w.instance)
}

// SetTabOrder
func (w *TWinControl) SetTabOrder(value TTabOrder) {
    WinControl_SetTabOrder(w.instance, value)
}

// TabStop
func (w *TWinControl) TabStop() bool {
    return WinControl_GetTabStop(w.instance)
}

// SetTabStop
func (w *TWinControl) SetTabStop(value bool) {
    WinControl_SetTabStop(w.instance, value)
}

// UseDockManager
func (w *TWinControl) UseDockManager() bool {
    return WinControl_GetUseDockManager(w.instance)
}

// SetUseDockManager
func (w *TWinControl) SetUseDockManager(value bool) {
    WinControl_SetUseDockManager(w.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (w *TWinControl) Enabled() bool {
    return WinControl_GetEnabled(w.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (w *TWinControl) SetEnabled(value bool) {
    WinControl_SetEnabled(w.instance, value)
}

// Action
func (w *TWinControl) Action() *TAction {
    return ActionFromInst(WinControl_GetAction(w.instance))
}

// SetAction
func (w *TWinControl) SetAction(value IComponent) {
    WinControl_SetAction(w.instance, CheckPtr(value))
}

// Align
func (w *TWinControl) Align() TAlign {
    return WinControl_GetAlign(w.instance)
}

// SetAlign
func (w *TWinControl) SetAlign(value TAlign) {
    WinControl_SetAlign(w.instance, value)
}

// Anchors
func (w *TWinControl) Anchors() TAnchors {
    return WinControl_GetAnchors(w.instance)
}

// SetAnchors
func (w *TWinControl) SetAnchors(value TAnchors) {
    WinControl_SetAnchors(w.instance, value)
}

// BiDiMode
func (w *TWinControl) BiDiMode() TBiDiMode {
    return WinControl_GetBiDiMode(w.instance)
}

// SetBiDiMode
func (w *TWinControl) SetBiDiMode(value TBiDiMode) {
    WinControl_SetBiDiMode(w.instance, value)
}

// BoundsRect
func (w *TWinControl) BoundsRect() TRect {
    return WinControl_GetBoundsRect(w.instance)
}

// SetBoundsRect
func (w *TWinControl) SetBoundsRect(value TRect) {
    WinControl_SetBoundsRect(w.instance, value)
}

// ClientHeight
func (w *TWinControl) ClientHeight() int32 {
    return WinControl_GetClientHeight(w.instance)
}

// SetClientHeight
func (w *TWinControl) SetClientHeight(value int32) {
    WinControl_SetClientHeight(w.instance, value)
}

// ClientOrigin
func (w *TWinControl) ClientOrigin() TPoint {
    return WinControl_GetClientOrigin(w.instance)
}

// ClientRect
func (w *TWinControl) ClientRect() TRect {
    return WinControl_GetClientRect(w.instance)
}

// ClientWidth
func (w *TWinControl) ClientWidth() int32 {
    return WinControl_GetClientWidth(w.instance)
}

// SetClientWidth
func (w *TWinControl) SetClientWidth(value int32) {
    WinControl_SetClientWidth(w.instance, value)
}

// ControlState
func (w *TWinControl) ControlState() TControlState {
    return WinControl_GetControlState(w.instance)
}

// SetControlState
func (w *TWinControl) SetControlState(value TControlState) {
    WinControl_SetControlState(w.instance, value)
}

// ControlStyle
func (w *TWinControl) ControlStyle() TControlStyle {
    return WinControl_GetControlStyle(w.instance)
}

// SetControlStyle
func (w *TWinControl) SetControlStyle(value TControlStyle) {
    WinControl_SetControlStyle(w.instance, value)
}

// ExplicitLeft
func (w *TWinControl) ExplicitLeft() int32 {
    return WinControl_GetExplicitLeft(w.instance)
}

// ExplicitTop
func (w *TWinControl) ExplicitTop() int32 {
    return WinControl_GetExplicitTop(w.instance)
}

// ExplicitWidth
func (w *TWinControl) ExplicitWidth() int32 {
    return WinControl_GetExplicitWidth(w.instance)
}

// ExplicitHeight
func (w *TWinControl) ExplicitHeight() int32 {
    return WinControl_GetExplicitHeight(w.instance)
}

// Floating
func (w *TWinControl) Floating() bool {
    return WinControl_GetFloating(w.instance)
}

// ShowHint
func (w *TWinControl) ShowHint() bool {
    return WinControl_GetShowHint(w.instance)
}

// SetShowHint
func (w *TWinControl) SetShowHint(value bool) {
    WinControl_SetShowHint(w.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (w *TWinControl) Visible() bool {
    return WinControl_GetVisible(w.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (w *TWinControl) SetVisible(value bool) {
    WinControl_SetVisible(w.instance, value)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (w *TWinControl) Parent() *TWinControl {
    return WinControlFromInst(WinControl_GetParent(w.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (w *TWinControl) SetParent(value IWinControl) {
    WinControl_SetParent(w.instance, CheckPtr(value))
}

// StyleElements
func (w *TWinControl) StyleElements() TStyleElements {
    return WinControl_GetStyleElements(w.instance)
}

// SetStyleElements
func (w *TWinControl) SetStyleElements(value TStyleElements) {
    WinControl_SetStyleElements(w.instance, value)
}

// AlignWithMargins
func (w *TWinControl) AlignWithMargins() bool {
    return WinControl_GetAlignWithMargins(w.instance)
}

// SetAlignWithMargins
func (w *TWinControl) SetAlignWithMargins(value bool) {
    WinControl_SetAlignWithMargins(w.instance, value)
}

// Left
func (w *TWinControl) Left() int32 {
    return WinControl_GetLeft(w.instance)
}

// SetLeft
func (w *TWinControl) SetLeft(value int32) {
    WinControl_SetLeft(w.instance, value)
}

// Top
func (w *TWinControl) Top() int32 {
    return WinControl_GetTop(w.instance)
}

// SetTop
func (w *TWinControl) SetTop(value int32) {
    WinControl_SetTop(w.instance, value)
}

// Width
func (w *TWinControl) Width() int32 {
    return WinControl_GetWidth(w.instance)
}

// SetWidth
func (w *TWinControl) SetWidth(value int32) {
    WinControl_SetWidth(w.instance, value)
}

// Height
func (w *TWinControl) Height() int32 {
    return WinControl_GetHeight(w.instance)
}

// SetHeight
func (w *TWinControl) SetHeight(value int32) {
    WinControl_SetHeight(w.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (w *TWinControl) Cursor() TCursor {
    return WinControl_GetCursor(w.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (w *TWinControl) SetCursor(value TCursor) {
    WinControl_SetCursor(w.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (w *TWinControl) Hint() string {
    return WinControl_GetHint(w.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (w *TWinControl) SetHint(value string) {
    WinControl_SetHint(w.instance, value)
}

// Margins
func (w *TWinControl) Margins() *TMargins {
    return MarginsFromInst(WinControl_GetMargins(w.instance))
}

// SetMargins
func (w *TWinControl) SetMargins(value *TMargins) {
    WinControl_SetMargins(w.instance, CheckPtr(value))
}

// CustomHint
func (w *TWinControl) CustomHint() *TCustomHint {
    return CustomHintFromInst(WinControl_GetCustomHint(w.instance))
}

// SetCustomHint
func (w *TWinControl) SetCustomHint(value IComponent) {
    WinControl_SetCustomHint(w.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (w *TWinControl) ComponentCount() int32 {
    return WinControl_GetComponentCount(w.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (w *TWinControl) ComponentIndex() int32 {
    return WinControl_GetComponentIndex(w.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (w *TWinControl) SetComponentIndex(value int32) {
    WinControl_SetComponentIndex(w.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (w *TWinControl) Owner() *TComponent {
    return ComponentFromInst(WinControl_GetOwner(w.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (w *TWinControl) Name() string {
    return WinControl_GetName(w.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (w *TWinControl) SetName(value string) {
    WinControl_SetName(w.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (w *TWinControl) Tag() int {
    return WinControl_GetTag(w.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (w *TWinControl) SetTag(value int) {
    WinControl_SetTag(w.instance, value)
}

// DockClients
func (w *TWinControl) DockClients(Index int32) *TControl {
    return ControlFromInst(WinControl_GetDockClients(w.instance, Index))
}

// Controls
func (w *TWinControl) Controls(Index int32) *TControl {
    return ControlFromInst(WinControl_GetControls(w.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (w *TWinControl) Components(AIndex int32) *TComponent {
    return ComponentFromInst(WinControl_GetComponents(w.instance, AIndex))
}

