
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

type TTabSheet struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTabSheet
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTabSheet(owner IComponent) *TTabSheet {
    t := new(TTabSheet)
    t.instance = TabSheet_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TabSheetFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TabSheetFromInst(inst uintptr) *TTabSheet {
    t := new(TTabSheet)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TabSheetFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TabSheetFromObj(obj IObject) *TTabSheet {
    t := new(TTabSheet)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TabSheetFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TabSheetFromUnsafePointer(ptr unsafe.Pointer) *TTabSheet {
    t := new(TTabSheet)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTabSheet) Free() {
    if t.instance != 0 {
        TabSheet_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTabSheet) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTabSheet) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTabSheet) IsValid() bool {
    return t.instance != 0
}

// TTabSheetClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTabSheetClass() TClass {
    return TabSheet_StaticClassType()
}

// CanFocus
func (t *TTabSheet) CanFocus() bool {
    return TabSheet_CanFocus(t.instance)
}

// FlipChildren
func (t *TTabSheet) FlipChildren(AllLevels bool) {
    TabSheet_FlipChildren(t.instance, AllLevels)
}

// Focused
func (t *TTabSheet) Focused() bool {
    return TabSheet_Focused(t.instance)
}

// HandleAllocated
func (t *TTabSheet) HandleAllocated() bool {
    return TabSheet_HandleAllocated(t.instance)
}

// Invalidate
func (t *TTabSheet) Invalidate() {
    TabSheet_Invalidate(t.instance)
}

// Realign
func (t *TTabSheet) Realign() {
    TabSheet_Realign(t.instance)
}

// Repaint
func (t *TTabSheet) Repaint() {
    TabSheet_Repaint(t.instance)
}

// ScaleBy
func (t *TTabSheet) ScaleBy(M int32, D int32) {
    TabSheet_ScaleBy(t.instance, M , D)
}

// SetBounds
func (t *TTabSheet) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    TabSheet_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (t *TTabSheet) SetFocus() {
    TabSheet_SetFocus(t.instance)
}

// Update
func (t *TTabSheet) Update() {
    TabSheet_Update(t.instance)
}

// BringToFront
func (t *TTabSheet) BringToFront() {
    TabSheet_BringToFront(t.instance)
}

// ClientToScreen
func (t *TTabSheet) ClientToScreen(Point TPoint) TPoint {
    return TabSheet_ClientToScreen(t.instance, Point)
}

// ClientToParent
func (t *TTabSheet) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return TabSheet_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// Dragging
func (t *TTabSheet) Dragging() bool {
    return TabSheet_Dragging(t.instance)
}

// HasParent
func (t *TTabSheet) HasParent() bool {
    return TabSheet_HasParent(t.instance)
}

// Hide
func (t *TTabSheet) Hide() {
    TabSheet_Hide(t.instance)
}

// Perform
func (t *TTabSheet) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return TabSheet_Perform(t.instance, Msg , WParam , LParam)
}

// Refresh
func (t *TTabSheet) Refresh() {
    TabSheet_Refresh(t.instance)
}

// ScreenToClient
func (t *TTabSheet) ScreenToClient(Point TPoint) TPoint {
    return TabSheet_ScreenToClient(t.instance, Point)
}

// ParentToClient
func (t *TTabSheet) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return TabSheet_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (t *TTabSheet) SendToBack() {
    TabSheet_SendToBack(t.instance)
}

// Show
func (t *TTabSheet) Show() {
    TabSheet_Show(t.instance)
}

// GetTextBuf
func (t *TTabSheet) GetTextBuf(Buffer string, BufSize int32) int32 {
    return TabSheet_GetTextBuf(t.instance, Buffer , BufSize)
}

// GetTextLen
func (t *TTabSheet) GetTextLen() int32 {
    return TabSheet_GetTextLen(t.instance)
}

// FindComponent
func (t *TTabSheet) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TabSheet_FindComponent(t.instance, AName))
}

// GetNamePath
func (t *TTabSheet) GetNamePath() string {
    return TabSheet_GetNamePath(t.instance)
}

// Assign
func (t *TTabSheet) Assign(Source IObject) {
    TabSheet_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTabSheet) DisposeOf() {
    TabSheet_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTabSheet) ClassType() TClass {
    return TabSheet_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTabSheet) ClassName() string {
    return TabSheet_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTabSheet) InstanceSize() int32 {
    return TabSheet_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTabSheet) InheritsFrom(AClass TClass) bool {
    return TabSheet_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTabSheet) Equals(Obj IObject) bool {
    return TabSheet_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTabSheet) GetHashCode() int32 {
    return TabSheet_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTabSheet) ToString() string {
    return TabSheet_ToString(t.instance)
}

// PageControl
func (t *TTabSheet) PageControl() *TPageControl {
    return PageControlFromInst(TabSheet_GetPageControl(t.instance))
}

// SetPageControl
func (t *TTabSheet) SetPageControl(value IWinControl) {
    TabSheet_SetPageControl(t.instance, CheckPtr(value))
}

// TabIndex
func (t *TTabSheet) TabIndex() int32 {
    return TabSheet_GetTabIndex(t.instance)
}

// BorderWidth
func (t *TTabSheet) BorderWidth() int32 {
    return TabSheet_GetBorderWidth(t.instance)
}

// SetBorderWidth
func (t *TTabSheet) SetBorderWidth(value int32) {
    TabSheet_SetBorderWidth(t.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (t *TTabSheet) Caption() string {
    return TabSheet_GetCaption(t.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (t *TTabSheet) SetCaption(value string) {
    TabSheet_SetCaption(t.instance, value)
}

// DoubleBuffered
func (t *TTabSheet) DoubleBuffered() bool {
    return TabSheet_GetDoubleBuffered(t.instance)
}

// SetDoubleBuffered
func (t *TTabSheet) SetDoubleBuffered(value bool) {
    TabSheet_SetDoubleBuffered(t.instance, value)
}

// DragMode
func (t *TTabSheet) DragMode() TDragMode {
    return TabSheet_GetDragMode(t.instance)
}

// SetDragMode
func (t *TTabSheet) SetDragMode(value TDragMode) {
    TabSheet_SetDragMode(t.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TTabSheet) Enabled() bool {
    return TabSheet_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TTabSheet) SetEnabled(value bool) {
    TabSheet_SetEnabled(t.instance, value)
}

// Font
func (t *TTabSheet) Font() *TFont {
    return FontFromInst(TabSheet_GetFont(t.instance))
}

// SetFont
func (t *TTabSheet) SetFont(value *TFont) {
    TabSheet_SetFont(t.instance, CheckPtr(value))
}

// Height
func (t *TTabSheet) Height() int32 {
    return TabSheet_GetHeight(t.instance)
}

// SetHeight
func (t *TTabSheet) SetHeight(value int32) {
    TabSheet_SetHeight(t.instance, value)
}

// Highlighted
func (t *TTabSheet) Highlighted() bool {
    return TabSheet_GetHighlighted(t.instance)
}

// SetHighlighted
func (t *TTabSheet) SetHighlighted(value bool) {
    TabSheet_SetHighlighted(t.instance, value)
}

// ImageIndex
func (t *TTabSheet) ImageIndex() int32 {
    return TabSheet_GetImageIndex(t.instance)
}

// SetImageIndex
func (t *TTabSheet) SetImageIndex(value int32) {
    TabSheet_SetImageIndex(t.instance, value)
}

// Left
func (t *TTabSheet) Left() int32 {
    return TabSheet_GetLeft(t.instance)
}

// SetLeft
func (t *TTabSheet) SetLeft(value int32) {
    TabSheet_SetLeft(t.instance, value)
}

// PageIndex
func (t *TTabSheet) PageIndex() int32 {
    return TabSheet_GetPageIndex(t.instance)
}

// SetPageIndex
func (t *TTabSheet) SetPageIndex(value int32) {
    TabSheet_SetPageIndex(t.instance, value)
}

// ParentDoubleBuffered
func (t *TTabSheet) ParentDoubleBuffered() bool {
    return TabSheet_GetParentDoubleBuffered(t.instance)
}

// SetParentDoubleBuffered
func (t *TTabSheet) SetParentDoubleBuffered(value bool) {
    TabSheet_SetParentDoubleBuffered(t.instance, value)
}

// ParentFont
func (t *TTabSheet) ParentFont() bool {
    return TabSheet_GetParentFont(t.instance)
}

// SetParentFont
func (t *TTabSheet) SetParentFont(value bool) {
    TabSheet_SetParentFont(t.instance, value)
}

// ParentShowHint
func (t *TTabSheet) ParentShowHint() bool {
    return TabSheet_GetParentShowHint(t.instance)
}

// SetParentShowHint
func (t *TTabSheet) SetParentShowHint(value bool) {
    TabSheet_SetParentShowHint(t.instance, value)
}

// PopupMenu
func (t *TTabSheet) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TabSheet_GetPopupMenu(t.instance))
}

// SetPopupMenu
func (t *TTabSheet) SetPopupMenu(value IComponent) {
    TabSheet_SetPopupMenu(t.instance, CheckPtr(value))
}

// ShowHint
func (t *TTabSheet) ShowHint() bool {
    return TabSheet_GetShowHint(t.instance)
}

// SetShowHint
func (t *TTabSheet) SetShowHint(value bool) {
    TabSheet_SetShowHint(t.instance, value)
}

// TabVisible
func (t *TTabSheet) TabVisible() bool {
    return TabSheet_GetTabVisible(t.instance)
}

// SetTabVisible
func (t *TTabSheet) SetTabVisible(value bool) {
    TabSheet_SetTabVisible(t.instance, value)
}

// Top
func (t *TTabSheet) Top() int32 {
    return TabSheet_GetTop(t.instance)
}

// SetTop
func (t *TTabSheet) SetTop(value int32) {
    TabSheet_SetTop(t.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TTabSheet) Visible() bool {
    return TabSheet_GetVisible(t.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TTabSheet) SetVisible(value bool) {
    TabSheet_SetVisible(t.instance, value)
}

// Width
func (t *TTabSheet) Width() int32 {
    return TabSheet_GetWidth(t.instance)
}

// SetWidth
func (t *TTabSheet) SetWidth(value int32) {
    TabSheet_SetWidth(t.instance, value)
}

// SetOnContextPopup
func (t *TTabSheet) SetOnContextPopup(fn TContextPopupEvent) {
    TabSheet_SetOnContextPopup(t.instance, fn)
}

// SetOnDragDrop
func (t *TTabSheet) SetOnDragDrop(fn TDragDropEvent) {
    TabSheet_SetOnDragDrop(t.instance, fn)
}

// SetOnDragOver
func (t *TTabSheet) SetOnDragOver(fn TDragOverEvent) {
    TabSheet_SetOnDragOver(t.instance, fn)
}

// SetOnEndDrag
func (t *TTabSheet) SetOnEndDrag(fn TEndDragEvent) {
    TabSheet_SetOnEndDrag(t.instance, fn)
}

// SetOnEnter
func (t *TTabSheet) SetOnEnter(fn TNotifyEvent) {
    TabSheet_SetOnEnter(t.instance, fn)
}

// SetOnExit
func (t *TTabSheet) SetOnExit(fn TNotifyEvent) {
    TabSheet_SetOnExit(t.instance, fn)
}

// SetOnHide
func (t *TTabSheet) SetOnHide(fn TNotifyEvent) {
    TabSheet_SetOnHide(t.instance, fn)
}

// SetOnMouseDown
func (t *TTabSheet) SetOnMouseDown(fn TMouseEvent) {
    TabSheet_SetOnMouseDown(t.instance, fn)
}

// SetOnMouseEnter
func (t *TTabSheet) SetOnMouseEnter(fn TNotifyEvent) {
    TabSheet_SetOnMouseEnter(t.instance, fn)
}

// SetOnMouseLeave
func (t *TTabSheet) SetOnMouseLeave(fn TNotifyEvent) {
    TabSheet_SetOnMouseLeave(t.instance, fn)
}

// SetOnMouseMove
func (t *TTabSheet) SetOnMouseMove(fn TMouseMoveEvent) {
    TabSheet_SetOnMouseMove(t.instance, fn)
}

// SetOnMouseUp
func (t *TTabSheet) SetOnMouseUp(fn TMouseEvent) {
    TabSheet_SetOnMouseUp(t.instance, fn)
}

// SetOnResize
func (t *TTabSheet) SetOnResize(fn TNotifyEvent) {
    TabSheet_SetOnResize(t.instance, fn)
}

// SetOnShow
func (t *TTabSheet) SetOnShow(fn TNotifyEvent) {
    TabSheet_SetOnShow(t.instance, fn)
}

// DockSite
func (t *TTabSheet) DockSite() bool {
    return TabSheet_GetDockSite(t.instance)
}

// SetDockSite
func (t *TTabSheet) SetDockSite(value bool) {
    TabSheet_SetDockSite(t.instance, value)
}

// Brush
func (t *TTabSheet) Brush() *TBrush {
    return BrushFromInst(TabSheet_GetBrush(t.instance))
}

// ControlCount
func (t *TTabSheet) ControlCount() int32 {
    return TabSheet_GetControlCount(t.instance)
}

// Handle
func (t *TTabSheet) Handle() HWND {
    return TabSheet_GetHandle(t.instance)
}

// ParentWindow
func (t *TTabSheet) ParentWindow() HWND {
    return TabSheet_GetParentWindow(t.instance)
}

// SetParentWindow
func (t *TTabSheet) SetParentWindow(value HWND) {
    TabSheet_SetParentWindow(t.instance, value)
}

// TabOrder
func (t *TTabSheet) TabOrder() uint16 {
    return TabSheet_GetTabOrder(t.instance)
}

// SetTabOrder
func (t *TTabSheet) SetTabOrder(value uint16) {
    TabSheet_SetTabOrder(t.instance, value)
}

// TabStop
func (t *TTabSheet) TabStop() bool {
    return TabSheet_GetTabStop(t.instance)
}

// SetTabStop
func (t *TTabSheet) SetTabStop(value bool) {
    TabSheet_SetTabStop(t.instance, value)
}

// UseDockManager
func (t *TTabSheet) UseDockManager() bool {
    return TabSheet_GetUseDockManager(t.instance)
}

// SetUseDockManager
func (t *TTabSheet) SetUseDockManager(value bool) {
    TabSheet_SetUseDockManager(t.instance, value)
}

// Action
func (t *TTabSheet) Action() *TAction {
    return ActionFromInst(TabSheet_GetAction(t.instance))
}

// SetAction
func (t *TTabSheet) SetAction(value IComponent) {
    TabSheet_SetAction(t.instance, CheckPtr(value))
}

// Align
func (t *TTabSheet) Align() TAlign {
    return TabSheet_GetAlign(t.instance)
}

// SetAlign
func (t *TTabSheet) SetAlign(value TAlign) {
    TabSheet_SetAlign(t.instance, value)
}

// Anchors
func (t *TTabSheet) Anchors() TAnchors {
    return TabSheet_GetAnchors(t.instance)
}

// SetAnchors
func (t *TTabSheet) SetAnchors(value TAnchors) {
    TabSheet_SetAnchors(t.instance, value)
}

// BiDiMode
func (t *TTabSheet) BiDiMode() TBiDiMode {
    return TabSheet_GetBiDiMode(t.instance)
}

// SetBiDiMode
func (t *TTabSheet) SetBiDiMode(value TBiDiMode) {
    TabSheet_SetBiDiMode(t.instance, value)
}

// BoundsRect
func (t *TTabSheet) BoundsRect() TRect {
    return TabSheet_GetBoundsRect(t.instance)
}

// SetBoundsRect
func (t *TTabSheet) SetBoundsRect(value TRect) {
    TabSheet_SetBoundsRect(t.instance, value)
}

// ClientHeight
func (t *TTabSheet) ClientHeight() int32 {
    return TabSheet_GetClientHeight(t.instance)
}

// SetClientHeight
func (t *TTabSheet) SetClientHeight(value int32) {
    TabSheet_SetClientHeight(t.instance, value)
}

// ClientRect
func (t *TTabSheet) ClientRect() TRect {
    return TabSheet_GetClientRect(t.instance)
}

// ClientWidth
func (t *TTabSheet) ClientWidth() int32 {
    return TabSheet_GetClientWidth(t.instance)
}

// SetClientWidth
func (t *TTabSheet) SetClientWidth(value int32) {
    TabSheet_SetClientWidth(t.instance, value)
}

// ExplicitLeft
func (t *TTabSheet) ExplicitLeft() int32 {
    return TabSheet_GetExplicitLeft(t.instance)
}

// ExplicitTop
func (t *TTabSheet) ExplicitTop() int32 {
    return TabSheet_GetExplicitTop(t.instance)
}

// ExplicitWidth
func (t *TTabSheet) ExplicitWidth() int32 {
    return TabSheet_GetExplicitWidth(t.instance)
}

// ExplicitHeight
func (t *TTabSheet) ExplicitHeight() int32 {
    return TabSheet_GetExplicitHeight(t.instance)
}

// Floating
func (t *TTabSheet) Floating() bool {
    return TabSheet_GetFloating(t.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TTabSheet) Parent() *TWinControl {
    return WinControlFromInst(TabSheet_GetParent(t.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (t *TTabSheet) SetParent(value IWinControl) {
    TabSheet_SetParent(t.instance, CheckPtr(value))
}

// StyleElements
func (t *TTabSheet) StyleElements() TStyleElements {
    return TabSheet_GetStyleElements(t.instance)
}

// SetStyleElements
func (t *TTabSheet) SetStyleElements(value TStyleElements) {
    TabSheet_SetStyleElements(t.instance, value)
}

// AlignWithMargins
func (t *TTabSheet) AlignWithMargins() bool {
    return TabSheet_GetAlignWithMargins(t.instance)
}

// SetAlignWithMargins
func (t *TTabSheet) SetAlignWithMargins(value bool) {
    TabSheet_SetAlignWithMargins(t.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (t *TTabSheet) Cursor() TCursor {
    return TabSheet_GetCursor(t.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (t *TTabSheet) SetCursor(value TCursor) {
    TabSheet_SetCursor(t.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (t *TTabSheet) Hint() string {
    return TabSheet_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (t *TTabSheet) SetHint(value string) {
    TabSheet_SetHint(t.instance, value)
}

// Margins
func (t *TTabSheet) Margins() *TMargins {
    return MarginsFromInst(TabSheet_GetMargins(t.instance))
}

// SetMargins
func (t *TTabSheet) SetMargins(value *TMargins) {
    TabSheet_SetMargins(t.instance, CheckPtr(value))
}

// CustomHint
func (t *TTabSheet) CustomHint() *TCustomHint {
    return CustomHintFromInst(TabSheet_GetCustomHint(t.instance))
}

// SetCustomHint
func (t *TTabSheet) SetCustomHint(value IComponent) {
    TabSheet_SetCustomHint(t.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTabSheet) ComponentCount() int32 {
    return TabSheet_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TTabSheet) ComponentIndex() int32 {
    return TabSheet_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TTabSheet) SetComponentIndex(value int32) {
    TabSheet_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTabSheet) Owner() *TComponent {
    return ComponentFromInst(TabSheet_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTabSheet) Name() string {
    return TabSheet_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTabSheet) SetName(value string) {
    TabSheet_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTabSheet) Tag() int {
    return TabSheet_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTabSheet) SetTag(value int) {
    TabSheet_SetTag(t.instance, value)
}

// Controls
func (t *TTabSheet) Controls(Index int32) *TControl {
    return ControlFromInst(TabSheet_GetControls(t.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTabSheet) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TabSheet_GetComponents(t.instance, AIndex))
}

