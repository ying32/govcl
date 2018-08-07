
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

type TBitBtn struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewBitBtn
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBitBtn(owner IComponent) *TBitBtn {
    b := new(TBitBtn)
    b.instance = BitBtn_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BitBtnFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func BitBtnFromInst(inst uintptr) *TBitBtn {
    b := new(TBitBtn)
    b.instance = inst
    b.ptr = unsafe.Pointer(inst)
    return b
}

// BitBtnFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func BitBtnFromObj(obj IObject) *TBitBtn {
    b := new(TBitBtn)
    b.instance = CheckPtr(obj)
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BitBtnFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func BitBtnFromUnsafePointer(ptr unsafe.Pointer) *TBitBtn {
    b := new(TBitBtn)
    b.instance = uintptr(ptr)
    b.ptr = ptr
    return b
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (b *TBitBtn) Free() {
    if b.instance != 0 {
        BitBtn_Free(b.instance)
        b.instance = 0
        b.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBitBtn) Instance() uintptr {
    return b.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBitBtn) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBitBtn) IsValid() bool {
    return b.instance != 0
}

// TBitBtnClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBitBtnClass() TClass {
    return BitBtn_StaticClassType()
}

// Click
func (b *TBitBtn) Click() {
    BitBtn_Click(b.instance)
}

// CanFocus
func (b *TBitBtn) CanFocus() bool {
    return BitBtn_CanFocus(b.instance)
}

// FlipChildren
func (b *TBitBtn) FlipChildren(AllLevels bool) {
    BitBtn_FlipChildren(b.instance, AllLevels)
}

// Focused
func (b *TBitBtn) Focused() bool {
    return BitBtn_Focused(b.instance)
}

// HandleAllocated
func (b *TBitBtn) HandleAllocated() bool {
    return BitBtn_HandleAllocated(b.instance)
}

// Invalidate
func (b *TBitBtn) Invalidate() {
    BitBtn_Invalidate(b.instance)
}

// Realign
func (b *TBitBtn) Realign() {
    BitBtn_Realign(b.instance)
}

// Repaint
func (b *TBitBtn) Repaint() {
    BitBtn_Repaint(b.instance)
}

// ScaleBy
func (b *TBitBtn) ScaleBy(M int32, D int32) {
    BitBtn_ScaleBy(b.instance, M , D)
}

// SetBounds
func (b *TBitBtn) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    BitBtn_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (b *TBitBtn) SetFocus() {
    BitBtn_SetFocus(b.instance)
}

// Update
func (b *TBitBtn) Update() {
    BitBtn_Update(b.instance)
}

// BringToFront
func (b *TBitBtn) BringToFront() {
    BitBtn_BringToFront(b.instance)
}

// ClientToScreen
func (b *TBitBtn) ClientToScreen(Point TPoint) TPoint {
    return BitBtn_ClientToScreen(b.instance, Point)
}

// ClientToParent
func (b *TBitBtn) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return BitBtn_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// Dragging
func (b *TBitBtn) Dragging() bool {
    return BitBtn_Dragging(b.instance)
}

// HasParent
func (b *TBitBtn) HasParent() bool {
    return BitBtn_HasParent(b.instance)
}

// Hide
func (b *TBitBtn) Hide() {
    BitBtn_Hide(b.instance)
}

// Perform
func (b *TBitBtn) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return BitBtn_Perform(b.instance, Msg , WParam , LParam)
}

// Refresh
func (b *TBitBtn) Refresh() {
    BitBtn_Refresh(b.instance)
}

// ScreenToClient
func (b *TBitBtn) ScreenToClient(Point TPoint) TPoint {
    return BitBtn_ScreenToClient(b.instance, Point)
}

// ParentToClient
func (b *TBitBtn) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return BitBtn_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (b *TBitBtn) SendToBack() {
    BitBtn_SendToBack(b.instance)
}

// Show
func (b *TBitBtn) Show() {
    BitBtn_Show(b.instance)
}

// GetTextBuf
func (b *TBitBtn) GetTextBuf(Buffer string, BufSize int32) int32 {
    return BitBtn_GetTextBuf(b.instance, Buffer , BufSize)
}

// GetTextLen
func (b *TBitBtn) GetTextLen() int32 {
    return BitBtn_GetTextLen(b.instance)
}

// FindComponent
func (b *TBitBtn) FindComponent(AName string) *TComponent {
    return ComponentFromInst(BitBtn_FindComponent(b.instance, AName))
}

// GetNamePath
func (b *TBitBtn) GetNamePath() string {
    return BitBtn_GetNamePath(b.instance)
}

// Assign
func (b *TBitBtn) Assign(Source IObject) {
    BitBtn_Assign(b.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TBitBtn) DisposeOf() {
    BitBtn_DisposeOf(b.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBitBtn) ClassType() TClass {
    return BitBtn_ClassType(b.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBitBtn) ClassName() string {
    return BitBtn_ClassName(b.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBitBtn) InstanceSize() int32 {
    return BitBtn_InstanceSize(b.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBitBtn) InheritsFrom(AClass TClass) bool {
    return BitBtn_InheritsFrom(b.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBitBtn) Equals(Obj IObject) bool {
    return BitBtn_Equals(b.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBitBtn) GetHashCode() int32 {
    return BitBtn_GetHashCode(b.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (b *TBitBtn) ToString() string {
    return BitBtn_ToString(b.instance)
}

// Action
func (b *TBitBtn) Action() *TAction {
    return ActionFromInst(BitBtn_GetAction(b.instance))
}

// SetAction
func (b *TBitBtn) SetAction(value IComponent) {
    BitBtn_SetAction(b.instance, CheckPtr(value))
}

// Align
func (b *TBitBtn) Align() TAlign {
    return BitBtn_GetAlign(b.instance)
}

// SetAlign
func (b *TBitBtn) SetAlign(value TAlign) {
    BitBtn_SetAlign(b.instance, value)
}

// Anchors
func (b *TBitBtn) Anchors() TAnchors {
    return BitBtn_GetAnchors(b.instance)
}

// SetAnchors
func (b *TBitBtn) SetAnchors(value TAnchors) {
    BitBtn_SetAnchors(b.instance, value)
}

// BiDiMode
func (b *TBitBtn) BiDiMode() TBiDiMode {
    return BitBtn_GetBiDiMode(b.instance)
}

// SetBiDiMode
func (b *TBitBtn) SetBiDiMode(value TBiDiMode) {
    BitBtn_SetBiDiMode(b.instance, value)
}

// Cancel
func (b *TBitBtn) Cancel() bool {
    return BitBtn_GetCancel(b.instance)
}

// SetCancel
func (b *TBitBtn) SetCancel(value bool) {
    BitBtn_SetCancel(b.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (b *TBitBtn) Caption() string {
    return BitBtn_GetCaption(b.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (b *TBitBtn) SetCaption(value string) {
    BitBtn_SetCaption(b.instance, value)
}

// Default
func (b *TBitBtn) Default() bool {
    return BitBtn_GetDefault(b.instance)
}

// SetDefault
func (b *TBitBtn) SetDefault(value bool) {
    BitBtn_SetDefault(b.instance, value)
}

// DoubleBuffered
func (b *TBitBtn) DoubleBuffered() bool {
    return BitBtn_GetDoubleBuffered(b.instance)
}

// SetDoubleBuffered
func (b *TBitBtn) SetDoubleBuffered(value bool) {
    BitBtn_SetDoubleBuffered(b.instance, value)
}

// DragCursor
func (b *TBitBtn) DragCursor() TCursor {
    return BitBtn_GetDragCursor(b.instance)
}

// SetDragCursor
func (b *TBitBtn) SetDragCursor(value TCursor) {
    BitBtn_SetDragCursor(b.instance, value)
}

// DragKind
func (b *TBitBtn) DragKind() TDragKind {
    return BitBtn_GetDragKind(b.instance)
}

// SetDragKind
func (b *TBitBtn) SetDragKind(value TDragKind) {
    BitBtn_SetDragKind(b.instance, value)
}

// DragMode
func (b *TBitBtn) DragMode() TDragMode {
    return BitBtn_GetDragMode(b.instance)
}

// SetDragMode
func (b *TBitBtn) SetDragMode(value TDragMode) {
    BitBtn_SetDragMode(b.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TBitBtn) Enabled() bool {
    return BitBtn_GetEnabled(b.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TBitBtn) SetEnabled(value bool) {
    BitBtn_SetEnabled(b.instance, value)
}

// Font
func (b *TBitBtn) Font() *TFont {
    return FontFromInst(BitBtn_GetFont(b.instance))
}

// SetFont
func (b *TBitBtn) SetFont(value *TFont) {
    BitBtn_SetFont(b.instance, CheckPtr(value))
}

// Glyph
func (b *TBitBtn) Glyph() *TBitmap {
    return BitmapFromInst(BitBtn_GetGlyph(b.instance))
}

// SetGlyph
func (b *TBitBtn) SetGlyph(value *TBitmap) {
    BitBtn_SetGlyph(b.instance, CheckPtr(value))
}

// Kind
func (b *TBitBtn) Kind() TBitBtnKind {
    return BitBtn_GetKind(b.instance)
}

// SetKind
func (b *TBitBtn) SetKind(value TBitBtnKind) {
    BitBtn_SetKind(b.instance, value)
}

// Layout
func (b *TBitBtn) Layout() TButtonLayout {
    return BitBtn_GetLayout(b.instance)
}

// SetLayout
func (b *TBitBtn) SetLayout(value TButtonLayout) {
    BitBtn_SetLayout(b.instance, value)
}

// ModalResult
func (b *TBitBtn) ModalResult() TModalResult {
    return BitBtn_GetModalResult(b.instance)
}

// SetModalResult
func (b *TBitBtn) SetModalResult(value TModalResult) {
    BitBtn_SetModalResult(b.instance, value)
}

// NumGlyphs
func (b *TBitBtn) NumGlyphs() TNumGlyphs {
    return BitBtn_GetNumGlyphs(b.instance)
}

// SetNumGlyphs
func (b *TBitBtn) SetNumGlyphs(value TNumGlyphs) {
    BitBtn_SetNumGlyphs(b.instance, value)
}

// ParentDoubleBuffered
func (b *TBitBtn) ParentDoubleBuffered() bool {
    return BitBtn_GetParentDoubleBuffered(b.instance)
}

// SetParentDoubleBuffered
func (b *TBitBtn) SetParentDoubleBuffered(value bool) {
    BitBtn_SetParentDoubleBuffered(b.instance, value)
}

// ParentFont
func (b *TBitBtn) ParentFont() bool {
    return BitBtn_GetParentFont(b.instance)
}

// SetParentFont
func (b *TBitBtn) SetParentFont(value bool) {
    BitBtn_SetParentFont(b.instance, value)
}

// ParentShowHint
func (b *TBitBtn) ParentShowHint() bool {
    return BitBtn_GetParentShowHint(b.instance)
}

// SetParentShowHint
func (b *TBitBtn) SetParentShowHint(value bool) {
    BitBtn_SetParentShowHint(b.instance, value)
}

// PopupMenu
func (b *TBitBtn) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(BitBtn_GetPopupMenu(b.instance))
}

// SetPopupMenu
func (b *TBitBtn) SetPopupMenu(value IComponent) {
    BitBtn_SetPopupMenu(b.instance, CheckPtr(value))
}

// ShowHint
func (b *TBitBtn) ShowHint() bool {
    return BitBtn_GetShowHint(b.instance)
}

// SetShowHint
func (b *TBitBtn) SetShowHint(value bool) {
    BitBtn_SetShowHint(b.instance, value)
}

// Style
func (b *TBitBtn) Style() TButtonStyle {
    return BitBtn_GetStyle(b.instance)
}

// SetStyle
func (b *TBitBtn) SetStyle(value TButtonStyle) {
    BitBtn_SetStyle(b.instance, value)
}

// Spacing
func (b *TBitBtn) Spacing() int32 {
    return BitBtn_GetSpacing(b.instance)
}

// SetSpacing
func (b *TBitBtn) SetSpacing(value int32) {
    BitBtn_SetSpacing(b.instance, value)
}

// TabOrder
func (b *TBitBtn) TabOrder() uint16 {
    return BitBtn_GetTabOrder(b.instance)
}

// SetTabOrder
func (b *TBitBtn) SetTabOrder(value uint16) {
    BitBtn_SetTabOrder(b.instance, value)
}

// TabStop
func (b *TBitBtn) TabStop() bool {
    return BitBtn_GetTabStop(b.instance)
}

// SetTabStop
func (b *TBitBtn) SetTabStop(value bool) {
    BitBtn_SetTabStop(b.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TBitBtn) Visible() bool {
    return BitBtn_GetVisible(b.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TBitBtn) SetVisible(value bool) {
    BitBtn_SetVisible(b.instance, value)
}

// WordWrap
func (b *TBitBtn) WordWrap() bool {
    return BitBtn_GetWordWrap(b.instance)
}

// SetWordWrap
func (b *TBitBtn) SetWordWrap(value bool) {
    BitBtn_SetWordWrap(b.instance, value)
}

// StyleElements
func (b *TBitBtn) StyleElements() TStyleElements {
    return BitBtn_GetStyleElements(b.instance)
}

// SetStyleElements
func (b *TBitBtn) SetStyleElements(value TStyleElements) {
    BitBtn_SetStyleElements(b.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (b *TBitBtn) SetOnClick(fn TNotifyEvent) {
    BitBtn_SetOnClick(b.instance, fn)
}

// SetOnContextPopup
func (b *TBitBtn) SetOnContextPopup(fn TContextPopupEvent) {
    BitBtn_SetOnContextPopup(b.instance, fn)
}

// SetOnDragDrop
func (b *TBitBtn) SetOnDragDrop(fn TDragDropEvent) {
    BitBtn_SetOnDragDrop(b.instance, fn)
}

// SetOnDragOver
func (b *TBitBtn) SetOnDragOver(fn TDragOverEvent) {
    BitBtn_SetOnDragOver(b.instance, fn)
}

// SetOnEndDock
func (b *TBitBtn) SetOnEndDock(fn TEndDragEvent) {
    BitBtn_SetOnEndDock(b.instance, fn)
}

// SetOnEndDrag
func (b *TBitBtn) SetOnEndDrag(fn TEndDragEvent) {
    BitBtn_SetOnEndDrag(b.instance, fn)
}

// SetOnEnter
func (b *TBitBtn) SetOnEnter(fn TNotifyEvent) {
    BitBtn_SetOnEnter(b.instance, fn)
}

// SetOnExit
func (b *TBitBtn) SetOnExit(fn TNotifyEvent) {
    BitBtn_SetOnExit(b.instance, fn)
}

// SetOnKeyDown
func (b *TBitBtn) SetOnKeyDown(fn TKeyEvent) {
    BitBtn_SetOnKeyDown(b.instance, fn)
}

// SetOnKeyPress
func (b *TBitBtn) SetOnKeyPress(fn TKeyPressEvent) {
    BitBtn_SetOnKeyPress(b.instance, fn)
}

// SetOnKeyUp
func (b *TBitBtn) SetOnKeyUp(fn TKeyEvent) {
    BitBtn_SetOnKeyUp(b.instance, fn)
}

// SetOnMouseDown
func (b *TBitBtn) SetOnMouseDown(fn TMouseEvent) {
    BitBtn_SetOnMouseDown(b.instance, fn)
}

// SetOnMouseEnter
func (b *TBitBtn) SetOnMouseEnter(fn TNotifyEvent) {
    BitBtn_SetOnMouseEnter(b.instance, fn)
}

// SetOnMouseLeave
func (b *TBitBtn) SetOnMouseLeave(fn TNotifyEvent) {
    BitBtn_SetOnMouseLeave(b.instance, fn)
}

// SetOnMouseMove
func (b *TBitBtn) SetOnMouseMove(fn TMouseMoveEvent) {
    BitBtn_SetOnMouseMove(b.instance, fn)
}

// SetOnMouseUp
func (b *TBitBtn) SetOnMouseUp(fn TMouseEvent) {
    BitBtn_SetOnMouseUp(b.instance, fn)
}

// SetOnStartDock
func (b *TBitBtn) SetOnStartDock(fn TStartDockEvent) {
    BitBtn_SetOnStartDock(b.instance, fn)
}

// CommandLinkHint
func (b *TBitBtn) CommandLinkHint() string {
    return BitBtn_GetCommandLinkHint(b.instance)
}

// SetCommandLinkHint
func (b *TBitBtn) SetCommandLinkHint(value string) {
    BitBtn_SetCommandLinkHint(b.instance, value)
}

// DisabledImageIndex
func (b *TBitBtn) DisabledImageIndex() int32 {
    return BitBtn_GetDisabledImageIndex(b.instance)
}

// SetDisabledImageIndex
func (b *TBitBtn) SetDisabledImageIndex(value int32) {
    BitBtn_SetDisabledImageIndex(b.instance, value)
}

// ElevationRequired
func (b *TBitBtn) ElevationRequired() bool {
    return BitBtn_GetElevationRequired(b.instance)
}

// SetElevationRequired
func (b *TBitBtn) SetElevationRequired(value bool) {
    BitBtn_SetElevationRequired(b.instance, value)
}

// HotImageIndex
func (b *TBitBtn) HotImageIndex() int32 {
    return BitBtn_GetHotImageIndex(b.instance)
}

// SetHotImageIndex
func (b *TBitBtn) SetHotImageIndex(value int32) {
    BitBtn_SetHotImageIndex(b.instance, value)
}

// Images
func (b *TBitBtn) Images() *TImageList {
    return ImageListFromInst(BitBtn_GetImages(b.instance))
}

// SetImages
func (b *TBitBtn) SetImages(value IComponent) {
    BitBtn_SetImages(b.instance, CheckPtr(value))
}

// ImageAlignment
func (b *TBitBtn) ImageAlignment() TImageAlignment {
    return BitBtn_GetImageAlignment(b.instance)
}

// SetImageAlignment
func (b *TBitBtn) SetImageAlignment(value TImageAlignment) {
    BitBtn_SetImageAlignment(b.instance, value)
}

// ImageIndex
func (b *TBitBtn) ImageIndex() int32 {
    return BitBtn_GetImageIndex(b.instance)
}

// SetImageIndex
func (b *TBitBtn) SetImageIndex(value int32) {
    BitBtn_SetImageIndex(b.instance, value)
}

// PressedImageIndex
func (b *TBitBtn) PressedImageIndex() int32 {
    return BitBtn_GetPressedImageIndex(b.instance)
}

// SetPressedImageIndex
func (b *TBitBtn) SetPressedImageIndex(value int32) {
    BitBtn_SetPressedImageIndex(b.instance, value)
}

// SelectedImageIndex
func (b *TBitBtn) SelectedImageIndex() int32 {
    return BitBtn_GetSelectedImageIndex(b.instance)
}

// SetSelectedImageIndex
func (b *TBitBtn) SetSelectedImageIndex(value int32) {
    BitBtn_SetSelectedImageIndex(b.instance, value)
}

// StylusHotImageIndex
func (b *TBitBtn) StylusHotImageIndex() int32 {
    return BitBtn_GetStylusHotImageIndex(b.instance)
}

// SetStylusHotImageIndex
func (b *TBitBtn) SetStylusHotImageIndex(value int32) {
    BitBtn_SetStylusHotImageIndex(b.instance, value)
}

// DockSite
func (b *TBitBtn) DockSite() bool {
    return BitBtn_GetDockSite(b.instance)
}

// SetDockSite
func (b *TBitBtn) SetDockSite(value bool) {
    BitBtn_SetDockSite(b.instance, value)
}

// Brush
func (b *TBitBtn) Brush() *TBrush {
    return BrushFromInst(BitBtn_GetBrush(b.instance))
}

// ControlCount
func (b *TBitBtn) ControlCount() int32 {
    return BitBtn_GetControlCount(b.instance)
}

// Handle
func (b *TBitBtn) Handle() HWND {
    return BitBtn_GetHandle(b.instance)
}

// ParentWindow
func (b *TBitBtn) ParentWindow() HWND {
    return BitBtn_GetParentWindow(b.instance)
}

// SetParentWindow
func (b *TBitBtn) SetParentWindow(value HWND) {
    BitBtn_SetParentWindow(b.instance, value)
}

// UseDockManager
func (b *TBitBtn) UseDockManager() bool {
    return BitBtn_GetUseDockManager(b.instance)
}

// SetUseDockManager
func (b *TBitBtn) SetUseDockManager(value bool) {
    BitBtn_SetUseDockManager(b.instance, value)
}

// BoundsRect
func (b *TBitBtn) BoundsRect() TRect {
    return BitBtn_GetBoundsRect(b.instance)
}

// SetBoundsRect
func (b *TBitBtn) SetBoundsRect(value TRect) {
    BitBtn_SetBoundsRect(b.instance, value)
}

// ClientHeight
func (b *TBitBtn) ClientHeight() int32 {
    return BitBtn_GetClientHeight(b.instance)
}

// SetClientHeight
func (b *TBitBtn) SetClientHeight(value int32) {
    BitBtn_SetClientHeight(b.instance, value)
}

// ClientRect
func (b *TBitBtn) ClientRect() TRect {
    return BitBtn_GetClientRect(b.instance)
}

// ClientWidth
func (b *TBitBtn) ClientWidth() int32 {
    return BitBtn_GetClientWidth(b.instance)
}

// SetClientWidth
func (b *TBitBtn) SetClientWidth(value int32) {
    BitBtn_SetClientWidth(b.instance, value)
}

// ExplicitLeft
func (b *TBitBtn) ExplicitLeft() int32 {
    return BitBtn_GetExplicitLeft(b.instance)
}

// ExplicitTop
func (b *TBitBtn) ExplicitTop() int32 {
    return BitBtn_GetExplicitTop(b.instance)
}

// ExplicitWidth
func (b *TBitBtn) ExplicitWidth() int32 {
    return BitBtn_GetExplicitWidth(b.instance)
}

// ExplicitHeight
func (b *TBitBtn) ExplicitHeight() int32 {
    return BitBtn_GetExplicitHeight(b.instance)
}

// Floating
func (b *TBitBtn) Floating() bool {
    return BitBtn_GetFloating(b.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TBitBtn) Parent() *TWinControl {
    return WinControlFromInst(BitBtn_GetParent(b.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TBitBtn) SetParent(value IWinControl) {
    BitBtn_SetParent(b.instance, CheckPtr(value))
}

// AlignWithMargins
func (b *TBitBtn) AlignWithMargins() bool {
    return BitBtn_GetAlignWithMargins(b.instance)
}

// SetAlignWithMargins
func (b *TBitBtn) SetAlignWithMargins(value bool) {
    BitBtn_SetAlignWithMargins(b.instance, value)
}

// Left
func (b *TBitBtn) Left() int32 {
    return BitBtn_GetLeft(b.instance)
}

// SetLeft
func (b *TBitBtn) SetLeft(value int32) {
    BitBtn_SetLeft(b.instance, value)
}

// Top
func (b *TBitBtn) Top() int32 {
    return BitBtn_GetTop(b.instance)
}

// SetTop
func (b *TBitBtn) SetTop(value int32) {
    BitBtn_SetTop(b.instance, value)
}

// Width
func (b *TBitBtn) Width() int32 {
    return BitBtn_GetWidth(b.instance)
}

// SetWidth
func (b *TBitBtn) SetWidth(value int32) {
    BitBtn_SetWidth(b.instance, value)
}

// Height
func (b *TBitBtn) Height() int32 {
    return BitBtn_GetHeight(b.instance)
}

// SetHeight
func (b *TBitBtn) SetHeight(value int32) {
    BitBtn_SetHeight(b.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TBitBtn) Cursor() TCursor {
    return BitBtn_GetCursor(b.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TBitBtn) SetCursor(value TCursor) {
    BitBtn_SetCursor(b.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (b *TBitBtn) Hint() string {
    return BitBtn_GetHint(b.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (b *TBitBtn) SetHint(value string) {
    BitBtn_SetHint(b.instance, value)
}

// Margins
func (b *TBitBtn) Margins() *TMargins {
    return MarginsFromInst(BitBtn_GetMargins(b.instance))
}

// SetMargins
func (b *TBitBtn) SetMargins(value *TMargins) {
    BitBtn_SetMargins(b.instance, CheckPtr(value))
}

// CustomHint
func (b *TBitBtn) CustomHint() *TCustomHint {
    return CustomHintFromInst(BitBtn_GetCustomHint(b.instance))
}

// SetCustomHint
func (b *TBitBtn) SetCustomHint(value IComponent) {
    BitBtn_SetCustomHint(b.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TBitBtn) ComponentCount() int32 {
    return BitBtn_GetComponentCount(b.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (b *TBitBtn) ComponentIndex() int32 {
    return BitBtn_GetComponentIndex(b.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (b *TBitBtn) SetComponentIndex(value int32) {
    BitBtn_SetComponentIndex(b.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TBitBtn) Owner() *TComponent {
    return ComponentFromInst(BitBtn_GetOwner(b.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (b *TBitBtn) Name() string {
    return BitBtn_GetName(b.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (b *TBitBtn) SetName(value string) {
    BitBtn_SetName(b.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TBitBtn) Tag() int {
    return BitBtn_GetTag(b.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TBitBtn) SetTag(value int) {
    BitBtn_SetTag(b.instance, value)
}

// Controls
func (b *TBitBtn) Controls(Index int32) *TControl {
    return ControlFromInst(BitBtn_GetControls(b.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TBitBtn) Components(AIndex int32) *TComponent {
    return ComponentFromInst(BitBtn_GetComponents(b.instance, AIndex))
}

