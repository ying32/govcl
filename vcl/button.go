
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

type TButton struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewButton
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewButton(owner IComponent) *TButton {
    b := new(TButton)
    b.instance = Button_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// ButtonFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ButtonFromInst(inst uintptr) *TButton {
    b := new(TButton)
    b.instance = inst
    b.ptr = unsafe.Pointer(inst)
    return b
}

// ButtonFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ButtonFromObj(obj IObject) *TButton {
    b := new(TButton)
    b.instance = CheckPtr(obj)
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// ButtonFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ButtonFromUnsafePointer(ptr unsafe.Pointer) *TButton {
    b := new(TButton)
    b.instance = uintptr(ptr)
    b.ptr = ptr
    return b
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (b *TButton) Free() {
    if b.instance != 0 {
        Button_Free(b.instance)
        b.instance = 0
        b.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TButton) Instance() uintptr {
    return b.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TButton) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TButton) IsValid() bool {
    return b.instance != 0
}

// TButtonClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TButtonClass() TClass {
    return Button_StaticClassType()
}

// Click
func (b *TButton) Click() {
    Button_Click(b.instance)
}

// CanFocus
func (b *TButton) CanFocus() bool {
    return Button_CanFocus(b.instance)
}

// FlipChildren
func (b *TButton) FlipChildren(AllLevels bool) {
    Button_FlipChildren(b.instance, AllLevels)
}

// Focused
func (b *TButton) Focused() bool {
    return Button_Focused(b.instance)
}

// HandleAllocated
func (b *TButton) HandleAllocated() bool {
    return Button_HandleAllocated(b.instance)
}

// Invalidate
func (b *TButton) Invalidate() {
    Button_Invalidate(b.instance)
}

// Realign
func (b *TButton) Realign() {
    Button_Realign(b.instance)
}

// Repaint
func (b *TButton) Repaint() {
    Button_Repaint(b.instance)
}

// ScaleBy
func (b *TButton) ScaleBy(M int32, D int32) {
    Button_ScaleBy(b.instance, M , D)
}

// SetBounds
func (b *TButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Button_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (b *TButton) SetFocus() {
    Button_SetFocus(b.instance)
}

// Update
func (b *TButton) Update() {
    Button_Update(b.instance)
}

// BringToFront
func (b *TButton) BringToFront() {
    Button_BringToFront(b.instance)
}

// ClientToScreen
func (b *TButton) ClientToScreen(Point TPoint) TPoint {
    return Button_ClientToScreen(b.instance, Point)
}

// ClientToParent
func (b *TButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Button_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// Dragging
func (b *TButton) Dragging() bool {
    return Button_Dragging(b.instance)
}

// HasParent
func (b *TButton) HasParent() bool {
    return Button_HasParent(b.instance)
}

// Hide
func (b *TButton) Hide() {
    Button_Hide(b.instance)
}

// Perform
func (b *TButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Button_Perform(b.instance, Msg , WParam , LParam)
}

// Refresh
func (b *TButton) Refresh() {
    Button_Refresh(b.instance)
}

// ScreenToClient
func (b *TButton) ScreenToClient(Point TPoint) TPoint {
    return Button_ScreenToClient(b.instance, Point)
}

// ParentToClient
func (b *TButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Button_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (b *TButton) SendToBack() {
    Button_SendToBack(b.instance)
}

// Show
func (b *TButton) Show() {
    Button_Show(b.instance)
}

// GetTextBuf
func (b *TButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Button_GetTextBuf(b.instance, Buffer , BufSize)
}

// GetTextLen
func (b *TButton) GetTextLen() int32 {
    return Button_GetTextLen(b.instance)
}

// FindComponent
func (b *TButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Button_FindComponent(b.instance, AName))
}

// GetNamePath
func (b *TButton) GetNamePath() string {
    return Button_GetNamePath(b.instance)
}

// Assign
func (b *TButton) Assign(Source IObject) {
    Button_Assign(b.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TButton) DisposeOf() {
    Button_DisposeOf(b.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TButton) ClassType() TClass {
    return Button_ClassType(b.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TButton) ClassName() string {
    return Button_ClassName(b.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TButton) InstanceSize() int32 {
    return Button_InstanceSize(b.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TButton) InheritsFrom(AClass TClass) bool {
    return Button_InheritsFrom(b.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TButton) Equals(Obj IObject) bool {
    return Button_Equals(b.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TButton) GetHashCode() int32 {
    return Button_GetHashCode(b.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (b *TButton) ToString() string {
    return Button_ToString(b.instance)
}

// Action
func (b *TButton) Action() *TAction {
    return ActionFromInst(Button_GetAction(b.instance))
}

// SetAction
func (b *TButton) SetAction(value IComponent) {
    Button_SetAction(b.instance, CheckPtr(value))
}

// Align
func (b *TButton) Align() TAlign {
    return Button_GetAlign(b.instance)
}

// SetAlign
func (b *TButton) SetAlign(value TAlign) {
    Button_SetAlign(b.instance, value)
}

// Anchors
func (b *TButton) Anchors() TAnchors {
    return Button_GetAnchors(b.instance)
}

// SetAnchors
func (b *TButton) SetAnchors(value TAnchors) {
    Button_SetAnchors(b.instance, value)
}

// BiDiMode
func (b *TButton) BiDiMode() TBiDiMode {
    return Button_GetBiDiMode(b.instance)
}

// SetBiDiMode
func (b *TButton) SetBiDiMode(value TBiDiMode) {
    Button_SetBiDiMode(b.instance, value)
}

// Cancel
func (b *TButton) Cancel() bool {
    return Button_GetCancel(b.instance)
}

// SetCancel
func (b *TButton) SetCancel(value bool) {
    Button_SetCancel(b.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (b *TButton) Caption() string {
    return Button_GetCaption(b.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (b *TButton) SetCaption(value string) {
    Button_SetCaption(b.instance, value)
}

// CommandLinkHint
func (b *TButton) CommandLinkHint() string {
    return Button_GetCommandLinkHint(b.instance)
}

// SetCommandLinkHint
func (b *TButton) SetCommandLinkHint(value string) {
    Button_SetCommandLinkHint(b.instance, value)
}

// Default
func (b *TButton) Default() bool {
    return Button_GetDefault(b.instance)
}

// SetDefault
func (b *TButton) SetDefault(value bool) {
    Button_SetDefault(b.instance, value)
}

// DisabledImageIndex
func (b *TButton) DisabledImageIndex() int32 {
    return Button_GetDisabledImageIndex(b.instance)
}

// SetDisabledImageIndex
func (b *TButton) SetDisabledImageIndex(value int32) {
    Button_SetDisabledImageIndex(b.instance, value)
}

// DoubleBuffered
func (b *TButton) DoubleBuffered() bool {
    return Button_GetDoubleBuffered(b.instance)
}

// SetDoubleBuffered
func (b *TButton) SetDoubleBuffered(value bool) {
    Button_SetDoubleBuffered(b.instance, value)
}

// DragCursor
func (b *TButton) DragCursor() TCursor {
    return Button_GetDragCursor(b.instance)
}

// SetDragCursor
func (b *TButton) SetDragCursor(value TCursor) {
    Button_SetDragCursor(b.instance, value)
}

// DragKind
func (b *TButton) DragKind() TDragKind {
    return Button_GetDragKind(b.instance)
}

// SetDragKind
func (b *TButton) SetDragKind(value TDragKind) {
    Button_SetDragKind(b.instance, value)
}

// DragMode
func (b *TButton) DragMode() TDragMode {
    return Button_GetDragMode(b.instance)
}

// SetDragMode
func (b *TButton) SetDragMode(value TDragMode) {
    Button_SetDragMode(b.instance, value)
}

// ElevationRequired
func (b *TButton) ElevationRequired() bool {
    return Button_GetElevationRequired(b.instance)
}

// SetElevationRequired
func (b *TButton) SetElevationRequired(value bool) {
    Button_SetElevationRequired(b.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TButton) Enabled() bool {
    return Button_GetEnabled(b.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TButton) SetEnabled(value bool) {
    Button_SetEnabled(b.instance, value)
}

// Font
func (b *TButton) Font() *TFont {
    return FontFromInst(Button_GetFont(b.instance))
}

// SetFont
func (b *TButton) SetFont(value *TFont) {
    Button_SetFont(b.instance, CheckPtr(value))
}

// HotImageIndex
func (b *TButton) HotImageIndex() int32 {
    return Button_GetHotImageIndex(b.instance)
}

// SetHotImageIndex
func (b *TButton) SetHotImageIndex(value int32) {
    Button_SetHotImageIndex(b.instance, value)
}

// ImageAlignment
func (b *TButton) ImageAlignment() TImageAlignment {
    return Button_GetImageAlignment(b.instance)
}

// SetImageAlignment
func (b *TButton) SetImageAlignment(value TImageAlignment) {
    Button_SetImageAlignment(b.instance, value)
}

// ImageIndex
func (b *TButton) ImageIndex() int32 {
    return Button_GetImageIndex(b.instance)
}

// SetImageIndex
func (b *TButton) SetImageIndex(value int32) {
    Button_SetImageIndex(b.instance, value)
}

// Images
func (b *TButton) Images() *TImageList {
    return ImageListFromInst(Button_GetImages(b.instance))
}

// SetImages
func (b *TButton) SetImages(value IComponent) {
    Button_SetImages(b.instance, CheckPtr(value))
}

// ModalResult
func (b *TButton) ModalResult() TModalResult {
    return Button_GetModalResult(b.instance)
}

// SetModalResult
func (b *TButton) SetModalResult(value TModalResult) {
    Button_SetModalResult(b.instance, value)
}

// ParentDoubleBuffered
func (b *TButton) ParentDoubleBuffered() bool {
    return Button_GetParentDoubleBuffered(b.instance)
}

// SetParentDoubleBuffered
func (b *TButton) SetParentDoubleBuffered(value bool) {
    Button_SetParentDoubleBuffered(b.instance, value)
}

// ParentFont
func (b *TButton) ParentFont() bool {
    return Button_GetParentFont(b.instance)
}

// SetParentFont
func (b *TButton) SetParentFont(value bool) {
    Button_SetParentFont(b.instance, value)
}

// ParentShowHint
func (b *TButton) ParentShowHint() bool {
    return Button_GetParentShowHint(b.instance)
}

// SetParentShowHint
func (b *TButton) SetParentShowHint(value bool) {
    Button_SetParentShowHint(b.instance, value)
}

// PopupMenu
func (b *TButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Button_GetPopupMenu(b.instance))
}

// SetPopupMenu
func (b *TButton) SetPopupMenu(value IComponent) {
    Button_SetPopupMenu(b.instance, CheckPtr(value))
}

// PressedImageIndex
func (b *TButton) PressedImageIndex() int32 {
    return Button_GetPressedImageIndex(b.instance)
}

// SetPressedImageIndex
func (b *TButton) SetPressedImageIndex(value int32) {
    Button_SetPressedImageIndex(b.instance, value)
}

// SelectedImageIndex
func (b *TButton) SelectedImageIndex() int32 {
    return Button_GetSelectedImageIndex(b.instance)
}

// SetSelectedImageIndex
func (b *TButton) SetSelectedImageIndex(value int32) {
    Button_SetSelectedImageIndex(b.instance, value)
}

// ShowHint
func (b *TButton) ShowHint() bool {
    return Button_GetShowHint(b.instance)
}

// SetShowHint
func (b *TButton) SetShowHint(value bool) {
    Button_SetShowHint(b.instance, value)
}

// Style
func (b *TButton) Style() TButtonStyle {
    return Button_GetStyle(b.instance)
}

// SetStyle
func (b *TButton) SetStyle(value TButtonStyle) {
    Button_SetStyle(b.instance, value)
}

// StylusHotImageIndex
func (b *TButton) StylusHotImageIndex() int32 {
    return Button_GetStylusHotImageIndex(b.instance)
}

// SetStylusHotImageIndex
func (b *TButton) SetStylusHotImageIndex(value int32) {
    Button_SetStylusHotImageIndex(b.instance, value)
}

// TabOrder
func (b *TButton) TabOrder() uint16 {
    return Button_GetTabOrder(b.instance)
}

// SetTabOrder
func (b *TButton) SetTabOrder(value uint16) {
    Button_SetTabOrder(b.instance, value)
}

// TabStop
func (b *TButton) TabStop() bool {
    return Button_GetTabStop(b.instance)
}

// SetTabStop
func (b *TButton) SetTabStop(value bool) {
    Button_SetTabStop(b.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TButton) Visible() bool {
    return Button_GetVisible(b.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TButton) SetVisible(value bool) {
    Button_SetVisible(b.instance, value)
}

// WordWrap
func (b *TButton) WordWrap() bool {
    return Button_GetWordWrap(b.instance)
}

// SetWordWrap
func (b *TButton) SetWordWrap(value bool) {
    Button_SetWordWrap(b.instance, value)
}

// StyleElements
func (b *TButton) StyleElements() TStyleElements {
    return Button_GetStyleElements(b.instance)
}

// SetStyleElements
func (b *TButton) SetStyleElements(value TStyleElements) {
    Button_SetStyleElements(b.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (b *TButton) SetOnClick(fn TNotifyEvent) {
    Button_SetOnClick(b.instance, fn)
}

// SetOnContextPopup
func (b *TButton) SetOnContextPopup(fn TContextPopupEvent) {
    Button_SetOnContextPopup(b.instance, fn)
}

// SetOnDragDrop
func (b *TButton) SetOnDragDrop(fn TDragDropEvent) {
    Button_SetOnDragDrop(b.instance, fn)
}

// SetOnDragOver
func (b *TButton) SetOnDragOver(fn TDragOverEvent) {
    Button_SetOnDragOver(b.instance, fn)
}

// SetOnEndDock
func (b *TButton) SetOnEndDock(fn TEndDragEvent) {
    Button_SetOnEndDock(b.instance, fn)
}

// SetOnEndDrag
func (b *TButton) SetOnEndDrag(fn TEndDragEvent) {
    Button_SetOnEndDrag(b.instance, fn)
}

// SetOnEnter
func (b *TButton) SetOnEnter(fn TNotifyEvent) {
    Button_SetOnEnter(b.instance, fn)
}

// SetOnExit
func (b *TButton) SetOnExit(fn TNotifyEvent) {
    Button_SetOnExit(b.instance, fn)
}

// SetOnKeyDown
func (b *TButton) SetOnKeyDown(fn TKeyEvent) {
    Button_SetOnKeyDown(b.instance, fn)
}

// SetOnKeyPress
func (b *TButton) SetOnKeyPress(fn TKeyPressEvent) {
    Button_SetOnKeyPress(b.instance, fn)
}

// SetOnKeyUp
func (b *TButton) SetOnKeyUp(fn TKeyEvent) {
    Button_SetOnKeyUp(b.instance, fn)
}

// SetOnMouseDown
func (b *TButton) SetOnMouseDown(fn TMouseEvent) {
    Button_SetOnMouseDown(b.instance, fn)
}

// SetOnMouseEnter
func (b *TButton) SetOnMouseEnter(fn TNotifyEvent) {
    Button_SetOnMouseEnter(b.instance, fn)
}

// SetOnMouseLeave
func (b *TButton) SetOnMouseLeave(fn TNotifyEvent) {
    Button_SetOnMouseLeave(b.instance, fn)
}

// SetOnMouseMove
func (b *TButton) SetOnMouseMove(fn TMouseMoveEvent) {
    Button_SetOnMouseMove(b.instance, fn)
}

// SetOnMouseUp
func (b *TButton) SetOnMouseUp(fn TMouseEvent) {
    Button_SetOnMouseUp(b.instance, fn)
}

// SetOnStartDock
func (b *TButton) SetOnStartDock(fn TStartDockEvent) {
    Button_SetOnStartDock(b.instance, fn)
}

// DockSite
func (b *TButton) DockSite() bool {
    return Button_GetDockSite(b.instance)
}

// SetDockSite
func (b *TButton) SetDockSite(value bool) {
    Button_SetDockSite(b.instance, value)
}

// Brush
func (b *TButton) Brush() *TBrush {
    return BrushFromInst(Button_GetBrush(b.instance))
}

// ControlCount
func (b *TButton) ControlCount() int32 {
    return Button_GetControlCount(b.instance)
}

// Handle
func (b *TButton) Handle() HWND {
    return Button_GetHandle(b.instance)
}

// ParentWindow
func (b *TButton) ParentWindow() HWND {
    return Button_GetParentWindow(b.instance)
}

// SetParentWindow
func (b *TButton) SetParentWindow(value HWND) {
    Button_SetParentWindow(b.instance, value)
}

// UseDockManager
func (b *TButton) UseDockManager() bool {
    return Button_GetUseDockManager(b.instance)
}

// SetUseDockManager
func (b *TButton) SetUseDockManager(value bool) {
    Button_SetUseDockManager(b.instance, value)
}

// BoundsRect
func (b *TButton) BoundsRect() TRect {
    return Button_GetBoundsRect(b.instance)
}

// SetBoundsRect
func (b *TButton) SetBoundsRect(value TRect) {
    Button_SetBoundsRect(b.instance, value)
}

// ClientHeight
func (b *TButton) ClientHeight() int32 {
    return Button_GetClientHeight(b.instance)
}

// SetClientHeight
func (b *TButton) SetClientHeight(value int32) {
    Button_SetClientHeight(b.instance, value)
}

// ClientRect
func (b *TButton) ClientRect() TRect {
    return Button_GetClientRect(b.instance)
}

// ClientWidth
func (b *TButton) ClientWidth() int32 {
    return Button_GetClientWidth(b.instance)
}

// SetClientWidth
func (b *TButton) SetClientWidth(value int32) {
    Button_SetClientWidth(b.instance, value)
}

// ExplicitLeft
func (b *TButton) ExplicitLeft() int32 {
    return Button_GetExplicitLeft(b.instance)
}

// ExplicitTop
func (b *TButton) ExplicitTop() int32 {
    return Button_GetExplicitTop(b.instance)
}

// ExplicitWidth
func (b *TButton) ExplicitWidth() int32 {
    return Button_GetExplicitWidth(b.instance)
}

// ExplicitHeight
func (b *TButton) ExplicitHeight() int32 {
    return Button_GetExplicitHeight(b.instance)
}

// Floating
func (b *TButton) Floating() bool {
    return Button_GetFloating(b.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TButton) Parent() *TWinControl {
    return WinControlFromInst(Button_GetParent(b.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TButton) SetParent(value IWinControl) {
    Button_SetParent(b.instance, CheckPtr(value))
}

// AlignWithMargins
func (b *TButton) AlignWithMargins() bool {
    return Button_GetAlignWithMargins(b.instance)
}

// SetAlignWithMargins
func (b *TButton) SetAlignWithMargins(value bool) {
    Button_SetAlignWithMargins(b.instance, value)
}

// Left
func (b *TButton) Left() int32 {
    return Button_GetLeft(b.instance)
}

// SetLeft
func (b *TButton) SetLeft(value int32) {
    Button_SetLeft(b.instance, value)
}

// Top
func (b *TButton) Top() int32 {
    return Button_GetTop(b.instance)
}

// SetTop
func (b *TButton) SetTop(value int32) {
    Button_SetTop(b.instance, value)
}

// Width
func (b *TButton) Width() int32 {
    return Button_GetWidth(b.instance)
}

// SetWidth
func (b *TButton) SetWidth(value int32) {
    Button_SetWidth(b.instance, value)
}

// Height
func (b *TButton) Height() int32 {
    return Button_GetHeight(b.instance)
}

// SetHeight
func (b *TButton) SetHeight(value int32) {
    Button_SetHeight(b.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TButton) Cursor() TCursor {
    return Button_GetCursor(b.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TButton) SetCursor(value TCursor) {
    Button_SetCursor(b.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (b *TButton) Hint() string {
    return Button_GetHint(b.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (b *TButton) SetHint(value string) {
    Button_SetHint(b.instance, value)
}

// Margins
func (b *TButton) Margins() *TMargins {
    return MarginsFromInst(Button_GetMargins(b.instance))
}

// SetMargins
func (b *TButton) SetMargins(value *TMargins) {
    Button_SetMargins(b.instance, CheckPtr(value))
}

// CustomHint
func (b *TButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(Button_GetCustomHint(b.instance))
}

// SetCustomHint
func (b *TButton) SetCustomHint(value IComponent) {
    Button_SetCustomHint(b.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TButton) ComponentCount() int32 {
    return Button_GetComponentCount(b.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (b *TButton) ComponentIndex() int32 {
    return Button_GetComponentIndex(b.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (b *TButton) SetComponentIndex(value int32) {
    Button_SetComponentIndex(b.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TButton) Owner() *TComponent {
    return ComponentFromInst(Button_GetOwner(b.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (b *TButton) Name() string {
    return Button_GetName(b.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (b *TButton) SetName(value string) {
    Button_SetName(b.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TButton) Tag() int {
    return Button_GetTag(b.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TButton) SetTag(value int) {
    Button_SetTag(b.instance, value)
}

// Controls
func (b *TButton) Controls(Index int32) *TControl {
    return ControlFromInst(Button_GetControls(b.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Button_GetComponents(b.instance, AIndex))
}

