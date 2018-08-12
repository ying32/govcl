
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

type TFlowPanel struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewFlowPanel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewFlowPanel(owner IComponent) *TFlowPanel {
    f := new(TFlowPanel)
    f.instance = FlowPanel_Create(CheckPtr(owner))
    f.ptr = unsafe.Pointer(f.instance)
    return f
}

// FlowPanelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func FlowPanelFromInst(inst uintptr) *TFlowPanel {
    f := new(TFlowPanel)
    f.instance = inst
    f.ptr = unsafe.Pointer(inst)
    return f
}

// FlowPanelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func FlowPanelFromObj(obj IObject) *TFlowPanel {
    f := new(TFlowPanel)
    f.instance = CheckPtr(obj)
    f.ptr = unsafe.Pointer(f.instance)
    return f
}

// FlowPanelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func FlowPanelFromUnsafePointer(ptr unsafe.Pointer) *TFlowPanel {
    f := new(TFlowPanel)
    f.instance = uintptr(ptr)
    f.ptr = ptr
    return f
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (f *TFlowPanel) Free() {
    if f.instance != 0 {
        FlowPanel_Free(f.instance)
        f.instance = 0
        f.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (f *TFlowPanel) Instance() uintptr {
    return f.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (f *TFlowPanel) UnsafeAddr() unsafe.Pointer {
    return f.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (f *TFlowPanel) IsValid() bool {
    return f.instance != 0
}

// TFlowPanelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TFlowPanelClass() TClass {
    return FlowPanel_StaticClassType()
}

// GetControlIndex
func (f *TFlowPanel) GetControlIndex(AControl IControl) int32 {
    return FlowPanel_GetControlIndex(f.instance, CheckPtr(AControl))
}

// SetControlIndex
func (f *TFlowPanel) SetControlIndex(AControl IControl, Index int32) {
    FlowPanel_SetControlIndex(f.instance, CheckPtr(AControl), Index)
}

// CanFocus
func (f *TFlowPanel) CanFocus() bool {
    return FlowPanel_CanFocus(f.instance)
}

// ContainsControl
func (f *TFlowPanel) ContainsControl(Control IControl) bool {
    return FlowPanel_ContainsControl(f.instance, CheckPtr(Control))
}

// ControlAtPos
func (f *TFlowPanel) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(FlowPanel_ControlAtPos(f.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (f *TFlowPanel) DisableAlign() {
    FlowPanel_DisableAlign(f.instance)
}

// EnableAlign
func (f *TFlowPanel) EnableAlign() {
    FlowPanel_EnableAlign(f.instance)
}

// FindChildControl
func (f *TFlowPanel) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(FlowPanel_FindChildControl(f.instance, ControlName))
}

// FlipChildren
func (f *TFlowPanel) FlipChildren(AllLevels bool) {
    FlowPanel_FlipChildren(f.instance, AllLevels)
}

// Focused
func (f *TFlowPanel) Focused() bool {
    return FlowPanel_Focused(f.instance)
}

// HandleAllocated
func (f *TFlowPanel) HandleAllocated() bool {
    return FlowPanel_HandleAllocated(f.instance)
}

// InsertControl
func (f *TFlowPanel) InsertControl(AControl IControl) {
    FlowPanel_InsertControl(f.instance, CheckPtr(AControl))
}

// Invalidate
func (f *TFlowPanel) Invalidate() {
    FlowPanel_Invalidate(f.instance)
}

// PaintTo
func (f *TFlowPanel) PaintTo(DC HDC, X int32, Y int32) {
    FlowPanel_PaintTo(f.instance, DC , X , Y)
}

// RemoveControl
func (f *TFlowPanel) RemoveControl(AControl IControl) {
    FlowPanel_RemoveControl(f.instance, CheckPtr(AControl))
}

// Realign
func (f *TFlowPanel) Realign() {
    FlowPanel_Realign(f.instance)
}

// Repaint
func (f *TFlowPanel) Repaint() {
    FlowPanel_Repaint(f.instance)
}

// ScaleBy
func (f *TFlowPanel) ScaleBy(M int32, D int32) {
    FlowPanel_ScaleBy(f.instance, M , D)
}

// ScrollBy
func (f *TFlowPanel) ScrollBy(DeltaX int32, DeltaY int32) {
    FlowPanel_ScrollBy(f.instance, DeltaX , DeltaY)
}

// SetBounds
func (f *TFlowPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    FlowPanel_SetBounds(f.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (f *TFlowPanel) SetFocus() {
    FlowPanel_SetFocus(f.instance)
}

// Update
func (f *TFlowPanel) Update() {
    FlowPanel_Update(f.instance)
}

// UpdateControlState
func (f *TFlowPanel) UpdateControlState() {
    FlowPanel_UpdateControlState(f.instance)
}

// BringToFront
func (f *TFlowPanel) BringToFront() {
    FlowPanel_BringToFront(f.instance)
}

// ClientToScreen
func (f *TFlowPanel) ClientToScreen(Point TPoint) TPoint {
    return FlowPanel_ClientToScreen(f.instance, Point)
}

// ClientToParent
func (f *TFlowPanel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return FlowPanel_ClientToParent(f.instance, Point , CheckPtr(AParent))
}

// Dragging
func (f *TFlowPanel) Dragging() bool {
    return FlowPanel_Dragging(f.instance)
}

// HasParent
func (f *TFlowPanel) HasParent() bool {
    return FlowPanel_HasParent(f.instance)
}

// Hide
func (f *TFlowPanel) Hide() {
    FlowPanel_Hide(f.instance)
}

// Perform
func (f *TFlowPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return FlowPanel_Perform(f.instance, Msg , WParam , LParam)
}

// Refresh
func (f *TFlowPanel) Refresh() {
    FlowPanel_Refresh(f.instance)
}

// ScreenToClient
func (f *TFlowPanel) ScreenToClient(Point TPoint) TPoint {
    return FlowPanel_ScreenToClient(f.instance, Point)
}

// ParentToClient
func (f *TFlowPanel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return FlowPanel_ParentToClient(f.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (f *TFlowPanel) SendToBack() {
    FlowPanel_SendToBack(f.instance)
}

// Show
func (f *TFlowPanel) Show() {
    FlowPanel_Show(f.instance)
}

// GetTextBuf
func (f *TFlowPanel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return FlowPanel_GetTextBuf(f.instance, Buffer , BufSize)
}

// GetTextLen
func (f *TFlowPanel) GetTextLen() int32 {
    return FlowPanel_GetTextLen(f.instance)
}

// SetTextBuf
func (f *TFlowPanel) SetTextBuf(Buffer string) {
    FlowPanel_SetTextBuf(f.instance, Buffer)
}

// FindComponent
func (f *TFlowPanel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(FlowPanel_FindComponent(f.instance, AName))
}

// GetNamePath
func (f *TFlowPanel) GetNamePath() string {
    return FlowPanel_GetNamePath(f.instance)
}

// Assign
func (f *TFlowPanel) Assign(Source IObject) {
    FlowPanel_Assign(f.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (f *TFlowPanel) DisposeOf() {
    FlowPanel_DisposeOf(f.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (f *TFlowPanel) ClassType() TClass {
    return FlowPanel_ClassType(f.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (f *TFlowPanel) ClassName() string {
    return FlowPanel_ClassName(f.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (f *TFlowPanel) InstanceSize() int32 {
    return FlowPanel_InstanceSize(f.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (f *TFlowPanel) InheritsFrom(AClass TClass) bool {
    return FlowPanel_InheritsFrom(f.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (f *TFlowPanel) Equals(Obj IObject) bool {
    return FlowPanel_Equals(f.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (f *TFlowPanel) GetHashCode() int32 {
    return FlowPanel_GetHashCode(f.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (f *TFlowPanel) ToString() string {
    return FlowPanel_ToString(f.instance)
}

// Align
func (f *TFlowPanel) Align() TAlign {
    return FlowPanel_GetAlign(f.instance)
}

// SetAlign
func (f *TFlowPanel) SetAlign(value TAlign) {
    FlowPanel_SetAlign(f.instance, value)
}

// Alignment
func (f *TFlowPanel) Alignment() TAlignment {
    return FlowPanel_GetAlignment(f.instance)
}

// SetAlignment
func (f *TFlowPanel) SetAlignment(value TAlignment) {
    FlowPanel_SetAlignment(f.instance, value)
}

// Anchors
func (f *TFlowPanel) Anchors() TAnchors {
    return FlowPanel_GetAnchors(f.instance)
}

// SetAnchors
func (f *TFlowPanel) SetAnchors(value TAnchors) {
    FlowPanel_SetAnchors(f.instance, value)
}

// AutoSize
func (f *TFlowPanel) AutoSize() bool {
    return FlowPanel_GetAutoSize(f.instance)
}

// SetAutoSize
func (f *TFlowPanel) SetAutoSize(value bool) {
    FlowPanel_SetAutoSize(f.instance, value)
}

// AutoWrap
func (f *TFlowPanel) AutoWrap() bool {
    return FlowPanel_GetAutoWrap(f.instance)
}

// SetAutoWrap
func (f *TFlowPanel) SetAutoWrap(value bool) {
    FlowPanel_SetAutoWrap(f.instance, value)
}

// BevelEdges
func (f *TFlowPanel) BevelEdges() TBevelEdges {
    return FlowPanel_GetBevelEdges(f.instance)
}

// SetBevelEdges
func (f *TFlowPanel) SetBevelEdges(value TBevelEdges) {
    FlowPanel_SetBevelEdges(f.instance, value)
}

// BevelInner
func (f *TFlowPanel) BevelInner() TBevelCut {
    return FlowPanel_GetBevelInner(f.instance)
}

// SetBevelInner
func (f *TFlowPanel) SetBevelInner(value TBevelCut) {
    FlowPanel_SetBevelInner(f.instance, value)
}

// BevelKind
func (f *TFlowPanel) BevelKind() TBevelKind {
    return FlowPanel_GetBevelKind(f.instance)
}

// SetBevelKind
func (f *TFlowPanel) SetBevelKind(value TBevelKind) {
    FlowPanel_SetBevelKind(f.instance, value)
}

// BevelOuter
func (f *TFlowPanel) BevelOuter() TBevelCut {
    return FlowPanel_GetBevelOuter(f.instance)
}

// SetBevelOuter
func (f *TFlowPanel) SetBevelOuter(value TBevelCut) {
    FlowPanel_SetBevelOuter(f.instance, value)
}

// BiDiMode
func (f *TFlowPanel) BiDiMode() TBiDiMode {
    return FlowPanel_GetBiDiMode(f.instance)
}

// SetBiDiMode
func (f *TFlowPanel) SetBiDiMode(value TBiDiMode) {
    FlowPanel_SetBiDiMode(f.instance, value)
}

// BorderWidth
func (f *TFlowPanel) BorderWidth() int32 {
    return FlowPanel_GetBorderWidth(f.instance)
}

// SetBorderWidth
func (f *TFlowPanel) SetBorderWidth(value int32) {
    FlowPanel_SetBorderWidth(f.instance, value)
}

// BorderStyle
func (f *TFlowPanel) BorderStyle() TBorderStyle {
    return FlowPanel_GetBorderStyle(f.instance)
}

// SetBorderStyle
func (f *TFlowPanel) SetBorderStyle(value TBorderStyle) {
    FlowPanel_SetBorderStyle(f.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (f *TFlowPanel) Caption() string {
    return FlowPanel_GetCaption(f.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (f *TFlowPanel) SetCaption(value string) {
    FlowPanel_SetCaption(f.instance, value)
}

// Color
func (f *TFlowPanel) Color() TColor {
    return FlowPanel_GetColor(f.instance)
}

// SetColor
func (f *TFlowPanel) SetColor(value TColor) {
    FlowPanel_SetColor(f.instance, value)
}

// UseDockManager
func (f *TFlowPanel) UseDockManager() bool {
    return FlowPanel_GetUseDockManager(f.instance)
}

// SetUseDockManager
func (f *TFlowPanel) SetUseDockManager(value bool) {
    FlowPanel_SetUseDockManager(f.instance, value)
}

// DockSite
func (f *TFlowPanel) DockSite() bool {
    return FlowPanel_GetDockSite(f.instance)
}

// SetDockSite
func (f *TFlowPanel) SetDockSite(value bool) {
    FlowPanel_SetDockSite(f.instance, value)
}

// DoubleBuffered
func (f *TFlowPanel) DoubleBuffered() bool {
    return FlowPanel_GetDoubleBuffered(f.instance)
}

// SetDoubleBuffered
func (f *TFlowPanel) SetDoubleBuffered(value bool) {
    FlowPanel_SetDoubleBuffered(f.instance, value)
}

// DragCursor
func (f *TFlowPanel) DragCursor() TCursor {
    return FlowPanel_GetDragCursor(f.instance)
}

// SetDragCursor
func (f *TFlowPanel) SetDragCursor(value TCursor) {
    FlowPanel_SetDragCursor(f.instance, value)
}

// DragKind
func (f *TFlowPanel) DragKind() TDragKind {
    return FlowPanel_GetDragKind(f.instance)
}

// SetDragKind
func (f *TFlowPanel) SetDragKind(value TDragKind) {
    FlowPanel_SetDragKind(f.instance, value)
}

// DragMode
func (f *TFlowPanel) DragMode() TDragMode {
    return FlowPanel_GetDragMode(f.instance)
}

// SetDragMode
func (f *TFlowPanel) SetDragMode(value TDragMode) {
    FlowPanel_SetDragMode(f.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (f *TFlowPanel) Enabled() bool {
    return FlowPanel_GetEnabled(f.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (f *TFlowPanel) SetEnabled(value bool) {
    FlowPanel_SetEnabled(f.instance, value)
}

// FlowStyle
func (f *TFlowPanel) FlowStyle() TFlowStyle {
    return FlowPanel_GetFlowStyle(f.instance)
}

// SetFlowStyle
func (f *TFlowPanel) SetFlowStyle(value TFlowStyle) {
    FlowPanel_SetFlowStyle(f.instance, value)
}

// FullRepaint
func (f *TFlowPanel) FullRepaint() bool {
    return FlowPanel_GetFullRepaint(f.instance)
}

// SetFullRepaint
func (f *TFlowPanel) SetFullRepaint(value bool) {
    FlowPanel_SetFullRepaint(f.instance, value)
}

// Font
func (f *TFlowPanel) Font() *TFont {
    return FontFromInst(FlowPanel_GetFont(f.instance))
}

// SetFont
func (f *TFlowPanel) SetFont(value *TFont) {
    FlowPanel_SetFont(f.instance, CheckPtr(value))
}

// Locked
func (f *TFlowPanel) Locked() bool {
    return FlowPanel_GetLocked(f.instance)
}

// SetLocked
func (f *TFlowPanel) SetLocked(value bool) {
    FlowPanel_SetLocked(f.instance, value)
}

// ParentBackground
func (f *TFlowPanel) ParentBackground() bool {
    return FlowPanel_GetParentBackground(f.instance)
}

// SetParentBackground
func (f *TFlowPanel) SetParentBackground(value bool) {
    FlowPanel_SetParentBackground(f.instance, value)
}

// ParentColor
func (f *TFlowPanel) ParentColor() bool {
    return FlowPanel_GetParentColor(f.instance)
}

// SetParentColor
func (f *TFlowPanel) SetParentColor(value bool) {
    FlowPanel_SetParentColor(f.instance, value)
}

// ParentCtl3D
func (f *TFlowPanel) ParentCtl3D() bool {
    return FlowPanel_GetParentCtl3D(f.instance)
}

// SetParentCtl3D
func (f *TFlowPanel) SetParentCtl3D(value bool) {
    FlowPanel_SetParentCtl3D(f.instance, value)
}

// ParentDoubleBuffered
func (f *TFlowPanel) ParentDoubleBuffered() bool {
    return FlowPanel_GetParentDoubleBuffered(f.instance)
}

// SetParentDoubleBuffered
func (f *TFlowPanel) SetParentDoubleBuffered(value bool) {
    FlowPanel_SetParentDoubleBuffered(f.instance, value)
}

// ParentFont
func (f *TFlowPanel) ParentFont() bool {
    return FlowPanel_GetParentFont(f.instance)
}

// SetParentFont
func (f *TFlowPanel) SetParentFont(value bool) {
    FlowPanel_SetParentFont(f.instance, value)
}

// ParentShowHint
func (f *TFlowPanel) ParentShowHint() bool {
    return FlowPanel_GetParentShowHint(f.instance)
}

// SetParentShowHint
func (f *TFlowPanel) SetParentShowHint(value bool) {
    FlowPanel_SetParentShowHint(f.instance, value)
}

// PopupMenu
func (f *TFlowPanel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(FlowPanel_GetPopupMenu(f.instance))
}

// SetPopupMenu
func (f *TFlowPanel) SetPopupMenu(value IComponent) {
    FlowPanel_SetPopupMenu(f.instance, CheckPtr(value))
}

// ShowCaption
func (f *TFlowPanel) ShowCaption() bool {
    return FlowPanel_GetShowCaption(f.instance)
}

// SetShowCaption
func (f *TFlowPanel) SetShowCaption(value bool) {
    FlowPanel_SetShowCaption(f.instance, value)
}

// ShowHint
func (f *TFlowPanel) ShowHint() bool {
    return FlowPanel_GetShowHint(f.instance)
}

// SetShowHint
func (f *TFlowPanel) SetShowHint(value bool) {
    FlowPanel_SetShowHint(f.instance, value)
}

// TabOrder
func (f *TFlowPanel) TabOrder() TTabOrder {
    return FlowPanel_GetTabOrder(f.instance)
}

// SetTabOrder
func (f *TFlowPanel) SetTabOrder(value TTabOrder) {
    FlowPanel_SetTabOrder(f.instance, value)
}

// TabStop
func (f *TFlowPanel) TabStop() bool {
    return FlowPanel_GetTabStop(f.instance)
}

// SetTabStop
func (f *TFlowPanel) SetTabStop(value bool) {
    FlowPanel_SetTabStop(f.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (f *TFlowPanel) Visible() bool {
    return FlowPanel_GetVisible(f.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (f *TFlowPanel) SetVisible(value bool) {
    FlowPanel_SetVisible(f.instance, value)
}

// StyleElements
func (f *TFlowPanel) StyleElements() TStyleElements {
    return FlowPanel_GetStyleElements(f.instance)
}

// SetStyleElements
func (f *TFlowPanel) SetStyleElements(value TStyleElements) {
    FlowPanel_SetStyleElements(f.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (f *TFlowPanel) SetOnClick(fn TNotifyEvent) {
    FlowPanel_SetOnClick(f.instance, fn)
}

// SetOnContextPopup
func (f *TFlowPanel) SetOnContextPopup(fn TContextPopupEvent) {
    FlowPanel_SetOnContextPopup(f.instance, fn)
}

// SetOnDockDrop
func (f *TFlowPanel) SetOnDockDrop(fn TDockDropEvent) {
    FlowPanel_SetOnDockDrop(f.instance, fn)
}

// SetOnDblClick
func (f *TFlowPanel) SetOnDblClick(fn TNotifyEvent) {
    FlowPanel_SetOnDblClick(f.instance, fn)
}

// SetOnDragDrop
func (f *TFlowPanel) SetOnDragDrop(fn TDragDropEvent) {
    FlowPanel_SetOnDragDrop(f.instance, fn)
}

// SetOnDragOver
func (f *TFlowPanel) SetOnDragOver(fn TDragOverEvent) {
    FlowPanel_SetOnDragOver(f.instance, fn)
}

// SetOnEndDock
func (f *TFlowPanel) SetOnEndDock(fn TEndDragEvent) {
    FlowPanel_SetOnEndDock(f.instance, fn)
}

// SetOnEndDrag
func (f *TFlowPanel) SetOnEndDrag(fn TEndDragEvent) {
    FlowPanel_SetOnEndDrag(f.instance, fn)
}

// SetOnEnter
func (f *TFlowPanel) SetOnEnter(fn TNotifyEvent) {
    FlowPanel_SetOnEnter(f.instance, fn)
}

// SetOnExit
func (f *TFlowPanel) SetOnExit(fn TNotifyEvent) {
    FlowPanel_SetOnExit(f.instance, fn)
}

// SetOnGetSiteInfo
func (f *TFlowPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    FlowPanel_SetOnGetSiteInfo(f.instance, fn)
}

// SetOnMouseDown
func (f *TFlowPanel) SetOnMouseDown(fn TMouseEvent) {
    FlowPanel_SetOnMouseDown(f.instance, fn)
}

// SetOnMouseEnter
func (f *TFlowPanel) SetOnMouseEnter(fn TNotifyEvent) {
    FlowPanel_SetOnMouseEnter(f.instance, fn)
}

// SetOnMouseLeave
func (f *TFlowPanel) SetOnMouseLeave(fn TNotifyEvent) {
    FlowPanel_SetOnMouseLeave(f.instance, fn)
}

// SetOnMouseMove
func (f *TFlowPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    FlowPanel_SetOnMouseMove(f.instance, fn)
}

// SetOnMouseUp
func (f *TFlowPanel) SetOnMouseUp(fn TMouseEvent) {
    FlowPanel_SetOnMouseUp(f.instance, fn)
}

// SetOnResize
func (f *TFlowPanel) SetOnResize(fn TNotifyEvent) {
    FlowPanel_SetOnResize(f.instance, fn)
}

// SetOnStartDock
func (f *TFlowPanel) SetOnStartDock(fn TStartDockEvent) {
    FlowPanel_SetOnStartDock(f.instance, fn)
}

// SetOnUnDock
func (f *TFlowPanel) SetOnUnDock(fn TUnDockEvent) {
    FlowPanel_SetOnUnDock(f.instance, fn)
}

// DockClientCount
func (f *TFlowPanel) DockClientCount() int32 {
    return FlowPanel_GetDockClientCount(f.instance)
}

// AlignDisabled
func (f *TFlowPanel) AlignDisabled() bool {
    return FlowPanel_GetAlignDisabled(f.instance)
}

// MouseInClient
func (f *TFlowPanel) MouseInClient() bool {
    return FlowPanel_GetMouseInClient(f.instance)
}

// VisibleDockClientCount
func (f *TFlowPanel) VisibleDockClientCount() int32 {
    return FlowPanel_GetVisibleDockClientCount(f.instance)
}

// Brush
func (f *TFlowPanel) Brush() *TBrush {
    return BrushFromInst(FlowPanel_GetBrush(f.instance))
}

// ControlCount
func (f *TFlowPanel) ControlCount() int32 {
    return FlowPanel_GetControlCount(f.instance)
}

// Handle
func (f *TFlowPanel) Handle() HWND {
    return FlowPanel_GetHandle(f.instance)
}

// ParentWindow
func (f *TFlowPanel) ParentWindow() HWND {
    return FlowPanel_GetParentWindow(f.instance)
}

// SetParentWindow
func (f *TFlowPanel) SetParentWindow(value HWND) {
    FlowPanel_SetParentWindow(f.instance, value)
}

// Action
func (f *TFlowPanel) Action() *TAction {
    return ActionFromInst(FlowPanel_GetAction(f.instance))
}

// SetAction
func (f *TFlowPanel) SetAction(value IComponent) {
    FlowPanel_SetAction(f.instance, CheckPtr(value))
}

// BoundsRect
func (f *TFlowPanel) BoundsRect() TRect {
    return FlowPanel_GetBoundsRect(f.instance)
}

// SetBoundsRect
func (f *TFlowPanel) SetBoundsRect(value TRect) {
    FlowPanel_SetBoundsRect(f.instance, value)
}

// ClientHeight
func (f *TFlowPanel) ClientHeight() int32 {
    return FlowPanel_GetClientHeight(f.instance)
}

// SetClientHeight
func (f *TFlowPanel) SetClientHeight(value int32) {
    FlowPanel_SetClientHeight(f.instance, value)
}

// ClientOrigin
func (f *TFlowPanel) ClientOrigin() TPoint {
    return FlowPanel_GetClientOrigin(f.instance)
}

// ClientRect
func (f *TFlowPanel) ClientRect() TRect {
    return FlowPanel_GetClientRect(f.instance)
}

// ClientWidth
func (f *TFlowPanel) ClientWidth() int32 {
    return FlowPanel_GetClientWidth(f.instance)
}

// SetClientWidth
func (f *TFlowPanel) SetClientWidth(value int32) {
    FlowPanel_SetClientWidth(f.instance, value)
}

// ControlState
func (f *TFlowPanel) ControlState() TControlState {
    return FlowPanel_GetControlState(f.instance)
}

// SetControlState
func (f *TFlowPanel) SetControlState(value TControlState) {
    FlowPanel_SetControlState(f.instance, value)
}

// ControlStyle
func (f *TFlowPanel) ControlStyle() TControlStyle {
    return FlowPanel_GetControlStyle(f.instance)
}

// SetControlStyle
func (f *TFlowPanel) SetControlStyle(value TControlStyle) {
    FlowPanel_SetControlStyle(f.instance, value)
}

// ExplicitLeft
func (f *TFlowPanel) ExplicitLeft() int32 {
    return FlowPanel_GetExplicitLeft(f.instance)
}

// ExplicitTop
func (f *TFlowPanel) ExplicitTop() int32 {
    return FlowPanel_GetExplicitTop(f.instance)
}

// ExplicitWidth
func (f *TFlowPanel) ExplicitWidth() int32 {
    return FlowPanel_GetExplicitWidth(f.instance)
}

// ExplicitHeight
func (f *TFlowPanel) ExplicitHeight() int32 {
    return FlowPanel_GetExplicitHeight(f.instance)
}

// Floating
func (f *TFlowPanel) Floating() bool {
    return FlowPanel_GetFloating(f.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (f *TFlowPanel) Parent() *TWinControl {
    return WinControlFromInst(FlowPanel_GetParent(f.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (f *TFlowPanel) SetParent(value IWinControl) {
    FlowPanel_SetParent(f.instance, CheckPtr(value))
}

// AlignWithMargins
func (f *TFlowPanel) AlignWithMargins() bool {
    return FlowPanel_GetAlignWithMargins(f.instance)
}

// SetAlignWithMargins
func (f *TFlowPanel) SetAlignWithMargins(value bool) {
    FlowPanel_SetAlignWithMargins(f.instance, value)
}

// Left
func (f *TFlowPanel) Left() int32 {
    return FlowPanel_GetLeft(f.instance)
}

// SetLeft
func (f *TFlowPanel) SetLeft(value int32) {
    FlowPanel_SetLeft(f.instance, value)
}

// Top
func (f *TFlowPanel) Top() int32 {
    return FlowPanel_GetTop(f.instance)
}

// SetTop
func (f *TFlowPanel) SetTop(value int32) {
    FlowPanel_SetTop(f.instance, value)
}

// Width
func (f *TFlowPanel) Width() int32 {
    return FlowPanel_GetWidth(f.instance)
}

// SetWidth
func (f *TFlowPanel) SetWidth(value int32) {
    FlowPanel_SetWidth(f.instance, value)
}

// Height
func (f *TFlowPanel) Height() int32 {
    return FlowPanel_GetHeight(f.instance)
}

// SetHeight
func (f *TFlowPanel) SetHeight(value int32) {
    FlowPanel_SetHeight(f.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (f *TFlowPanel) Cursor() TCursor {
    return FlowPanel_GetCursor(f.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (f *TFlowPanel) SetCursor(value TCursor) {
    FlowPanel_SetCursor(f.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (f *TFlowPanel) Hint() string {
    return FlowPanel_GetHint(f.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (f *TFlowPanel) SetHint(value string) {
    FlowPanel_SetHint(f.instance, value)
}

// Margins
func (f *TFlowPanel) Margins() *TMargins {
    return MarginsFromInst(FlowPanel_GetMargins(f.instance))
}

// SetMargins
func (f *TFlowPanel) SetMargins(value *TMargins) {
    FlowPanel_SetMargins(f.instance, CheckPtr(value))
}

// CustomHint
func (f *TFlowPanel) CustomHint() *TCustomHint {
    return CustomHintFromInst(FlowPanel_GetCustomHint(f.instance))
}

// SetCustomHint
func (f *TFlowPanel) SetCustomHint(value IComponent) {
    FlowPanel_SetCustomHint(f.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (f *TFlowPanel) ComponentCount() int32 {
    return FlowPanel_GetComponentCount(f.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (f *TFlowPanel) ComponentIndex() int32 {
    return FlowPanel_GetComponentIndex(f.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (f *TFlowPanel) SetComponentIndex(value int32) {
    FlowPanel_SetComponentIndex(f.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (f *TFlowPanel) Owner() *TComponent {
    return ComponentFromInst(FlowPanel_GetOwner(f.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (f *TFlowPanel) Name() string {
    return FlowPanel_GetName(f.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (f *TFlowPanel) SetName(value string) {
    FlowPanel_SetName(f.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (f *TFlowPanel) Tag() int {
    return FlowPanel_GetTag(f.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (f *TFlowPanel) SetTag(value int) {
    FlowPanel_SetTag(f.instance, value)
}

// DockClients
func (f *TFlowPanel) DockClients(Index int32) *TControl {
    return ControlFromInst(FlowPanel_GetDockClients(f.instance, Index))
}

// Controls
func (f *TFlowPanel) Controls(Index int32) *TControl {
    return ControlFromInst(FlowPanel_GetControls(f.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (f *TFlowPanel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(FlowPanel_GetComponents(f.instance, AIndex))
}

