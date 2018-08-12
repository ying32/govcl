
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

type TCategoryPanel struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCategoryPanel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCategoryPanel(owner IComponent) *TCategoryPanel {
    c := new(TCategoryPanel)
    c.instance = CategoryPanel_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CategoryPanelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CategoryPanelFromInst(inst uintptr) *TCategoryPanel {
    c := new(TCategoryPanel)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CategoryPanelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CategoryPanelFromObj(obj IObject) *TCategoryPanel {
    c := new(TCategoryPanel)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CategoryPanelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CategoryPanelFromUnsafePointer(ptr unsafe.Pointer) *TCategoryPanel {
    c := new(TCategoryPanel)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCategoryPanel) Free() {
    if c.instance != 0 {
        CategoryPanel_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCategoryPanel) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCategoryPanel) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCategoryPanel) IsValid() bool {
    return c.instance != 0
}

// TCategoryPanelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCategoryPanelClass() TClass {
    return CategoryPanel_StaticClassType()
}

// Collapse
func (c *TCategoryPanel) Collapse() {
    CategoryPanel_Collapse(c.instance)
}

// Expand
func (c *TCategoryPanel) Expand() {
    CategoryPanel_Expand(c.instance)
}

// SetBounds
func (c *TCategoryPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CategoryPanel_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// CanFocus
func (c *TCategoryPanel) CanFocus() bool {
    return CategoryPanel_CanFocus(c.instance)
}

// ContainsControl
func (c *TCategoryPanel) ContainsControl(Control IControl) bool {
    return CategoryPanel_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
func (c *TCategoryPanel) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(CategoryPanel_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (c *TCategoryPanel) DisableAlign() {
    CategoryPanel_DisableAlign(c.instance)
}

// EnableAlign
func (c *TCategoryPanel) EnableAlign() {
    CategoryPanel_EnableAlign(c.instance)
}

// FindChildControl
func (c *TCategoryPanel) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(CategoryPanel_FindChildControl(c.instance, ControlName))
}

// FlipChildren
func (c *TCategoryPanel) FlipChildren(AllLevels bool) {
    CategoryPanel_FlipChildren(c.instance, AllLevels)
}

// Focused
func (c *TCategoryPanel) Focused() bool {
    return CategoryPanel_Focused(c.instance)
}

// HandleAllocated
func (c *TCategoryPanel) HandleAllocated() bool {
    return CategoryPanel_HandleAllocated(c.instance)
}

// InsertControl
func (c *TCategoryPanel) InsertControl(AControl IControl) {
    CategoryPanel_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
func (c *TCategoryPanel) Invalidate() {
    CategoryPanel_Invalidate(c.instance)
}

// PaintTo
func (c *TCategoryPanel) PaintTo(DC HDC, X int32, Y int32) {
    CategoryPanel_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
func (c *TCategoryPanel) RemoveControl(AControl IControl) {
    CategoryPanel_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
func (c *TCategoryPanel) Realign() {
    CategoryPanel_Realign(c.instance)
}

// Repaint
func (c *TCategoryPanel) Repaint() {
    CategoryPanel_Repaint(c.instance)
}

// ScaleBy
func (c *TCategoryPanel) ScaleBy(M int32, D int32) {
    CategoryPanel_ScaleBy(c.instance, M , D)
}

// ScrollBy
func (c *TCategoryPanel) ScrollBy(DeltaX int32, DeltaY int32) {
    CategoryPanel_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetFocus
func (c *TCategoryPanel) SetFocus() {
    CategoryPanel_SetFocus(c.instance)
}

// Update
func (c *TCategoryPanel) Update() {
    CategoryPanel_Update(c.instance)
}

// UpdateControlState
func (c *TCategoryPanel) UpdateControlState() {
    CategoryPanel_UpdateControlState(c.instance)
}

// BringToFront
func (c *TCategoryPanel) BringToFront() {
    CategoryPanel_BringToFront(c.instance)
}

// ClientToScreen
func (c *TCategoryPanel) ClientToScreen(Point TPoint) TPoint {
    return CategoryPanel_ClientToScreen(c.instance, Point)
}

// ClientToParent
func (c *TCategoryPanel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CategoryPanel_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
func (c *TCategoryPanel) Dragging() bool {
    return CategoryPanel_Dragging(c.instance)
}

// HasParent
func (c *TCategoryPanel) HasParent() bool {
    return CategoryPanel_HasParent(c.instance)
}

// Hide
func (c *TCategoryPanel) Hide() {
    CategoryPanel_Hide(c.instance)
}

// Perform
func (c *TCategoryPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CategoryPanel_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
func (c *TCategoryPanel) Refresh() {
    CategoryPanel_Refresh(c.instance)
}

// ScreenToClient
func (c *TCategoryPanel) ScreenToClient(Point TPoint) TPoint {
    return CategoryPanel_ScreenToClient(c.instance, Point)
}

// ParentToClient
func (c *TCategoryPanel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CategoryPanel_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (c *TCategoryPanel) SendToBack() {
    CategoryPanel_SendToBack(c.instance)
}

// Show
func (c *TCategoryPanel) Show() {
    CategoryPanel_Show(c.instance)
}

// GetTextBuf
func (c *TCategoryPanel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CategoryPanel_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
func (c *TCategoryPanel) GetTextLen() int32 {
    return CategoryPanel_GetTextLen(c.instance)
}

// SetTextBuf
func (c *TCategoryPanel) SetTextBuf(Buffer string) {
    CategoryPanel_SetTextBuf(c.instance, Buffer)
}

// FindComponent
func (c *TCategoryPanel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CategoryPanel_FindComponent(c.instance, AName))
}

// GetNamePath
func (c *TCategoryPanel) GetNamePath() string {
    return CategoryPanel_GetNamePath(c.instance)
}

// Assign
func (c *TCategoryPanel) Assign(Source IObject) {
    CategoryPanel_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCategoryPanel) DisposeOf() {
    CategoryPanel_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCategoryPanel) ClassType() TClass {
    return CategoryPanel_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCategoryPanel) ClassName() string {
    return CategoryPanel_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCategoryPanel) InstanceSize() int32 {
    return CategoryPanel_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCategoryPanel) InheritsFrom(AClass TClass) bool {
    return CategoryPanel_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCategoryPanel) Equals(Obj IObject) bool {
    return CategoryPanel_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCategoryPanel) GetHashCode() int32 {
    return CategoryPanel_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCategoryPanel) ToString() string {
    return CategoryPanel_ToString(c.instance)
}

// BiDiMode
func (c *TCategoryPanel) BiDiMode() TBiDiMode {
    return CategoryPanel_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TCategoryPanel) SetBiDiMode(value TBiDiMode) {
    CategoryPanel_SetBiDiMode(c.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (c *TCategoryPanel) Caption() string {
    return CategoryPanel_GetCaption(c.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (c *TCategoryPanel) SetCaption(value string) {
    CategoryPanel_SetCaption(c.instance, value)
}

// Color
func (c *TCategoryPanel) Color() TColor {
    return CategoryPanel_GetColor(c.instance)
}

// SetColor
func (c *TCategoryPanel) SetColor(value TColor) {
    CategoryPanel_SetColor(c.instance, value)
}

// Collapsed
func (c *TCategoryPanel) Collapsed() bool {
    return CategoryPanel_GetCollapsed(c.instance)
}

// SetCollapsed
func (c *TCategoryPanel) SetCollapsed(value bool) {
    CategoryPanel_SetCollapsed(c.instance, value)
}

// CollapsedHotImageIndex
func (c *TCategoryPanel) CollapsedHotImageIndex() int32 {
    return CategoryPanel_GetCollapsedHotImageIndex(c.instance)
}

// SetCollapsedHotImageIndex
func (c *TCategoryPanel) SetCollapsedHotImageIndex(value int32) {
    CategoryPanel_SetCollapsedHotImageIndex(c.instance, value)
}

// CollapsedImageIndex
func (c *TCategoryPanel) CollapsedImageIndex() int32 {
    return CategoryPanel_GetCollapsedImageIndex(c.instance)
}

// SetCollapsedImageIndex
func (c *TCategoryPanel) SetCollapsedImageIndex(value int32) {
    CategoryPanel_SetCollapsedImageIndex(c.instance, value)
}

// CollapsedPressedImageIndex
func (c *TCategoryPanel) CollapsedPressedImageIndex() int32 {
    return CategoryPanel_GetCollapsedPressedImageIndex(c.instance)
}

// SetCollapsedPressedImageIndex
func (c *TCategoryPanel) SetCollapsedPressedImageIndex(value int32) {
    CategoryPanel_SetCollapsedPressedImageIndex(c.instance, value)
}

// UseDockManager
func (c *TCategoryPanel) UseDockManager() bool {
    return CategoryPanel_GetUseDockManager(c.instance)
}

// SetUseDockManager
func (c *TCategoryPanel) SetUseDockManager(value bool) {
    CategoryPanel_SetUseDockManager(c.instance, value)
}

// DockSite
func (c *TCategoryPanel) DockSite() bool {
    return CategoryPanel_GetDockSite(c.instance)
}

// SetDockSite
func (c *TCategoryPanel) SetDockSite(value bool) {
    CategoryPanel_SetDockSite(c.instance, value)
}

// DoubleBuffered
func (c *TCategoryPanel) DoubleBuffered() bool {
    return CategoryPanel_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
func (c *TCategoryPanel) SetDoubleBuffered(value bool) {
    CategoryPanel_SetDoubleBuffered(c.instance, value)
}

// DragCursor
func (c *TCategoryPanel) DragCursor() TCursor {
    return CategoryPanel_GetDragCursor(c.instance)
}

// SetDragCursor
func (c *TCategoryPanel) SetDragCursor(value TCursor) {
    CategoryPanel_SetDragCursor(c.instance, value)
}

// DragKind
func (c *TCategoryPanel) DragKind() TDragKind {
    return CategoryPanel_GetDragKind(c.instance)
}

// SetDragKind
func (c *TCategoryPanel) SetDragKind(value TDragKind) {
    CategoryPanel_SetDragKind(c.instance, value)
}

// DragMode
func (c *TCategoryPanel) DragMode() TDragMode {
    return CategoryPanel_GetDragMode(c.instance)
}

// SetDragMode
func (c *TCategoryPanel) SetDragMode(value TDragMode) {
    CategoryPanel_SetDragMode(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TCategoryPanel) Enabled() bool {
    return CategoryPanel_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TCategoryPanel) SetEnabled(value bool) {
    CategoryPanel_SetEnabled(c.instance, value)
}

// ExpandedHotImageIndex
func (c *TCategoryPanel) ExpandedHotImageIndex() int32 {
    return CategoryPanel_GetExpandedHotImageIndex(c.instance)
}

// SetExpandedHotImageIndex
func (c *TCategoryPanel) SetExpandedHotImageIndex(value int32) {
    CategoryPanel_SetExpandedHotImageIndex(c.instance, value)
}

// ExpandedImageIndex
func (c *TCategoryPanel) ExpandedImageIndex() int32 {
    return CategoryPanel_GetExpandedImageIndex(c.instance)
}

// SetExpandedImageIndex
func (c *TCategoryPanel) SetExpandedImageIndex(value int32) {
    CategoryPanel_SetExpandedImageIndex(c.instance, value)
}

// ExpandedPressedImageIndex
func (c *TCategoryPanel) ExpandedPressedImageIndex() int32 {
    return CategoryPanel_GetExpandedPressedImageIndex(c.instance)
}

// SetExpandedPressedImageIndex
func (c *TCategoryPanel) SetExpandedPressedImageIndex(value int32) {
    CategoryPanel_SetExpandedPressedImageIndex(c.instance, value)
}

// FullRepaint
func (c *TCategoryPanel) FullRepaint() bool {
    return CategoryPanel_GetFullRepaint(c.instance)
}

// SetFullRepaint
func (c *TCategoryPanel) SetFullRepaint(value bool) {
    CategoryPanel_SetFullRepaint(c.instance, value)
}

// Font
func (c *TCategoryPanel) Font() *TFont {
    return FontFromInst(CategoryPanel_GetFont(c.instance))
}

// SetFont
func (c *TCategoryPanel) SetFont(value *TFont) {
    CategoryPanel_SetFont(c.instance, CheckPtr(value))
}

// Height
func (c *TCategoryPanel) Height() int32 {
    return CategoryPanel_GetHeight(c.instance)
}

// SetHeight
func (c *TCategoryPanel) SetHeight(value int32) {
    CategoryPanel_SetHeight(c.instance, value)
}

// Left
func (c *TCategoryPanel) Left() int32 {
    return CategoryPanel_GetLeft(c.instance)
}

// SetLeft
func (c *TCategoryPanel) SetLeft(value int32) {
    CategoryPanel_SetLeft(c.instance, value)
}

// Locked
func (c *TCategoryPanel) Locked() bool {
    return CategoryPanel_GetLocked(c.instance)
}

// SetLocked
func (c *TCategoryPanel) SetLocked(value bool) {
    CategoryPanel_SetLocked(c.instance, value)
}

// ParentBackground
func (c *TCategoryPanel) ParentBackground() bool {
    return CategoryPanel_GetParentBackground(c.instance)
}

// SetParentBackground
func (c *TCategoryPanel) SetParentBackground(value bool) {
    CategoryPanel_SetParentBackground(c.instance, value)
}

// ParentColor
func (c *TCategoryPanel) ParentColor() bool {
    return CategoryPanel_GetParentColor(c.instance)
}

// SetParentColor
func (c *TCategoryPanel) SetParentColor(value bool) {
    CategoryPanel_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TCategoryPanel) ParentCtl3D() bool {
    return CategoryPanel_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TCategoryPanel) SetParentCtl3D(value bool) {
    CategoryPanel_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
func (c *TCategoryPanel) ParentDoubleBuffered() bool {
    return CategoryPanel_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
func (c *TCategoryPanel) SetParentDoubleBuffered(value bool) {
    CategoryPanel_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
func (c *TCategoryPanel) ParentFont() bool {
    return CategoryPanel_GetParentFont(c.instance)
}

// SetParentFont
func (c *TCategoryPanel) SetParentFont(value bool) {
    CategoryPanel_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TCategoryPanel) ParentShowHint() bool {
    return CategoryPanel_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TCategoryPanel) SetParentShowHint(value bool) {
    CategoryPanel_SetParentShowHint(c.instance, value)
}

// PopupMenu
func (c *TCategoryPanel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CategoryPanel_GetPopupMenu(c.instance))
}

// SetPopupMenu
func (c *TCategoryPanel) SetPopupMenu(value IComponent) {
    CategoryPanel_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
func (c *TCategoryPanel) ShowHint() bool {
    return CategoryPanel_GetShowHint(c.instance)
}

// SetShowHint
func (c *TCategoryPanel) SetShowHint(value bool) {
    CategoryPanel_SetShowHint(c.instance, value)
}

// TabOrder
func (c *TCategoryPanel) TabOrder() TTabOrder {
    return CategoryPanel_GetTabOrder(c.instance)
}

// SetTabOrder
func (c *TCategoryPanel) SetTabOrder(value TTabOrder) {
    CategoryPanel_SetTabOrder(c.instance, value)
}

// TabStop
func (c *TCategoryPanel) TabStop() bool {
    return CategoryPanel_GetTabStop(c.instance)
}

// SetTabStop
func (c *TCategoryPanel) SetTabStop(value bool) {
    CategoryPanel_SetTabStop(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TCategoryPanel) Visible() bool {
    return CategoryPanel_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TCategoryPanel) SetVisible(value bool) {
    CategoryPanel_SetVisible(c.instance, value)
}

// StyleElements
func (c *TCategoryPanel) StyleElements() TStyleElements {
    return CategoryPanel_GetStyleElements(c.instance)
}

// SetStyleElements
func (c *TCategoryPanel) SetStyleElements(value TStyleElements) {
    CategoryPanel_SetStyleElements(c.instance, value)
}

// Width
func (c *TCategoryPanel) Width() int32 {
    return CategoryPanel_GetWidth(c.instance)
}

// SetWidth
func (c *TCategoryPanel) SetWidth(value int32) {
    CategoryPanel_SetWidth(c.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TCategoryPanel) SetOnClick(fn TNotifyEvent) {
    CategoryPanel_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
func (c *TCategoryPanel) SetOnContextPopup(fn TContextPopupEvent) {
    CategoryPanel_SetOnContextPopup(c.instance, fn)
}

// SetOnDockDrop
func (c *TCategoryPanel) SetOnDockDrop(fn TDockDropEvent) {
    CategoryPanel_SetOnDockDrop(c.instance, fn)
}

// SetOnDblClick
func (c *TCategoryPanel) SetOnDblClick(fn TNotifyEvent) {
    CategoryPanel_SetOnDblClick(c.instance, fn)
}

// SetOnDragDrop
func (c *TCategoryPanel) SetOnDragDrop(fn TDragDropEvent) {
    CategoryPanel_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
func (c *TCategoryPanel) SetOnDragOver(fn TDragOverEvent) {
    CategoryPanel_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
func (c *TCategoryPanel) SetOnEndDock(fn TEndDragEvent) {
    CategoryPanel_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
func (c *TCategoryPanel) SetOnEndDrag(fn TEndDragEvent) {
    CategoryPanel_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
func (c *TCategoryPanel) SetOnEnter(fn TNotifyEvent) {
    CategoryPanel_SetOnEnter(c.instance, fn)
}

// SetOnExit
func (c *TCategoryPanel) SetOnExit(fn TNotifyEvent) {
    CategoryPanel_SetOnExit(c.instance, fn)
}

// SetOnGetSiteInfo
func (c *TCategoryPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    CategoryPanel_SetOnGetSiteInfo(c.instance, fn)
}

// SetOnMouseDown
func (c *TCategoryPanel) SetOnMouseDown(fn TMouseEvent) {
    CategoryPanel_SetOnMouseDown(c.instance, fn)
}

// SetOnMouseEnter
func (c *TCategoryPanel) SetOnMouseEnter(fn TNotifyEvent) {
    CategoryPanel_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
func (c *TCategoryPanel) SetOnMouseLeave(fn TNotifyEvent) {
    CategoryPanel_SetOnMouseLeave(c.instance, fn)
}

// SetOnMouseMove
func (c *TCategoryPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    CategoryPanel_SetOnMouseMove(c.instance, fn)
}

// SetOnMouseUp
func (c *TCategoryPanel) SetOnMouseUp(fn TMouseEvent) {
    CategoryPanel_SetOnMouseUp(c.instance, fn)
}

// SetOnStartDock
func (c *TCategoryPanel) SetOnStartDock(fn TStartDockEvent) {
    CategoryPanel_SetOnStartDock(c.instance, fn)
}

// SetOnUnDock
func (c *TCategoryPanel) SetOnUnDock(fn TUnDockEvent) {
    CategoryPanel_SetOnUnDock(c.instance, fn)
}

// PanelGroup
func (c *TCategoryPanel) PanelGroup() *TCategoryPanelGroup {
    return CategoryPanelGroupFromInst(CategoryPanel_GetPanelGroup(c.instance))
}

// SetPanelGroup
func (c *TCategoryPanel) SetPanelGroup(value IWinControl) {
    CategoryPanel_SetPanelGroup(c.instance, CheckPtr(value))
}

// DockClientCount
func (c *TCategoryPanel) DockClientCount() int32 {
    return CategoryPanel_GetDockClientCount(c.instance)
}

// AlignDisabled
func (c *TCategoryPanel) AlignDisabled() bool {
    return CategoryPanel_GetAlignDisabled(c.instance)
}

// MouseInClient
func (c *TCategoryPanel) MouseInClient() bool {
    return CategoryPanel_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
func (c *TCategoryPanel) VisibleDockClientCount() int32 {
    return CategoryPanel_GetVisibleDockClientCount(c.instance)
}

// Brush
func (c *TCategoryPanel) Brush() *TBrush {
    return BrushFromInst(CategoryPanel_GetBrush(c.instance))
}

// ControlCount
func (c *TCategoryPanel) ControlCount() int32 {
    return CategoryPanel_GetControlCount(c.instance)
}

// Handle
func (c *TCategoryPanel) Handle() HWND {
    return CategoryPanel_GetHandle(c.instance)
}

// ParentWindow
func (c *TCategoryPanel) ParentWindow() HWND {
    return CategoryPanel_GetParentWindow(c.instance)
}

// SetParentWindow
func (c *TCategoryPanel) SetParentWindow(value HWND) {
    CategoryPanel_SetParentWindow(c.instance, value)
}

// Action
func (c *TCategoryPanel) Action() *TAction {
    return ActionFromInst(CategoryPanel_GetAction(c.instance))
}

// SetAction
func (c *TCategoryPanel) SetAction(value IComponent) {
    CategoryPanel_SetAction(c.instance, CheckPtr(value))
}

// Align
func (c *TCategoryPanel) Align() TAlign {
    return CategoryPanel_GetAlign(c.instance)
}

// SetAlign
func (c *TCategoryPanel) SetAlign(value TAlign) {
    CategoryPanel_SetAlign(c.instance, value)
}

// Anchors
func (c *TCategoryPanel) Anchors() TAnchors {
    return CategoryPanel_GetAnchors(c.instance)
}

// SetAnchors
func (c *TCategoryPanel) SetAnchors(value TAnchors) {
    CategoryPanel_SetAnchors(c.instance, value)
}

// BoundsRect
func (c *TCategoryPanel) BoundsRect() TRect {
    return CategoryPanel_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TCategoryPanel) SetBoundsRect(value TRect) {
    CategoryPanel_SetBoundsRect(c.instance, value)
}

// ClientHeight
func (c *TCategoryPanel) ClientHeight() int32 {
    return CategoryPanel_GetClientHeight(c.instance)
}

// SetClientHeight
func (c *TCategoryPanel) SetClientHeight(value int32) {
    CategoryPanel_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TCategoryPanel) ClientOrigin() TPoint {
    return CategoryPanel_GetClientOrigin(c.instance)
}

// ClientRect
func (c *TCategoryPanel) ClientRect() TRect {
    return CategoryPanel_GetClientRect(c.instance)
}

// ClientWidth
func (c *TCategoryPanel) ClientWidth() int32 {
    return CategoryPanel_GetClientWidth(c.instance)
}

// SetClientWidth
func (c *TCategoryPanel) SetClientWidth(value int32) {
    CategoryPanel_SetClientWidth(c.instance, value)
}

// ControlState
func (c *TCategoryPanel) ControlState() TControlState {
    return CategoryPanel_GetControlState(c.instance)
}

// SetControlState
func (c *TCategoryPanel) SetControlState(value TControlState) {
    CategoryPanel_SetControlState(c.instance, value)
}

// ControlStyle
func (c *TCategoryPanel) ControlStyle() TControlStyle {
    return CategoryPanel_GetControlStyle(c.instance)
}

// SetControlStyle
func (c *TCategoryPanel) SetControlStyle(value TControlStyle) {
    CategoryPanel_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TCategoryPanel) ExplicitLeft() int32 {
    return CategoryPanel_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TCategoryPanel) ExplicitTop() int32 {
    return CategoryPanel_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TCategoryPanel) ExplicitWidth() int32 {
    return CategoryPanel_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TCategoryPanel) ExplicitHeight() int32 {
    return CategoryPanel_GetExplicitHeight(c.instance)
}

// Floating
func (c *TCategoryPanel) Floating() bool {
    return CategoryPanel_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TCategoryPanel) Parent() *TWinControl {
    return WinControlFromInst(CategoryPanel_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TCategoryPanel) SetParent(value IWinControl) {
    CategoryPanel_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
func (c *TCategoryPanel) AlignWithMargins() bool {
    return CategoryPanel_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
func (c *TCategoryPanel) SetAlignWithMargins(value bool) {
    CategoryPanel_SetAlignWithMargins(c.instance, value)
}

// Top
func (c *TCategoryPanel) Top() int32 {
    return CategoryPanel_GetTop(c.instance)
}

// SetTop
func (c *TCategoryPanel) SetTop(value int32) {
    CategoryPanel_SetTop(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TCategoryPanel) Cursor() TCursor {
    return CategoryPanel_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TCategoryPanel) SetCursor(value TCursor) {
    CategoryPanel_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (c *TCategoryPanel) Hint() string {
    return CategoryPanel_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (c *TCategoryPanel) SetHint(value string) {
    CategoryPanel_SetHint(c.instance, value)
}

// Margins
func (c *TCategoryPanel) Margins() *TMargins {
    return MarginsFromInst(CategoryPanel_GetMargins(c.instance))
}

// SetMargins
func (c *TCategoryPanel) SetMargins(value *TMargins) {
    CategoryPanel_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
func (c *TCategoryPanel) CustomHint() *TCustomHint {
    return CustomHintFromInst(CategoryPanel_GetCustomHint(c.instance))
}

// SetCustomHint
func (c *TCategoryPanel) SetCustomHint(value IComponent) {
    CategoryPanel_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TCategoryPanel) ComponentCount() int32 {
    return CategoryPanel_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TCategoryPanel) ComponentIndex() int32 {
    return CategoryPanel_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TCategoryPanel) SetComponentIndex(value int32) {
    CategoryPanel_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TCategoryPanel) Owner() *TComponent {
    return ComponentFromInst(CategoryPanel_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TCategoryPanel) Name() string {
    return CategoryPanel_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TCategoryPanel) SetName(value string) {
    CategoryPanel_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TCategoryPanel) Tag() int {
    return CategoryPanel_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TCategoryPanel) SetTag(value int) {
    CategoryPanel_SetTag(c.instance, value)
}

// DockClients
func (c *TCategoryPanel) DockClients(Index int32) *TControl {
    return ControlFromInst(CategoryPanel_GetDockClients(c.instance, Index))
}

// Controls
func (c *TCategoryPanel) Controls(Index int32) *TControl {
    return ControlFromInst(CategoryPanel_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TCategoryPanel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CategoryPanel_GetComponents(c.instance, AIndex))
}

