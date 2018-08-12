
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

type TToolBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewToolBar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewToolBar(owner IComponent) *TToolBar {
    t := new(TToolBar)
    t.instance = ToolBar_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// ToolBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ToolBarFromInst(inst uintptr) *TToolBar {
    t := new(TToolBar)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// ToolBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ToolBarFromObj(obj IObject) *TToolBar {
    t := new(TToolBar)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// ToolBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ToolBarFromUnsafePointer(ptr unsafe.Pointer) *TToolBar {
    t := new(TToolBar)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TToolBar) Free() {
    if t.instance != 0 {
        ToolBar_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TToolBar) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TToolBar) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TToolBar) IsValid() bool {
    return t.instance != 0
}

// TToolBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TToolBarClass() TClass {
    return ToolBar_StaticClassType()
}

// FlipChildren
func (t *TToolBar) FlipChildren(AllLevels bool) {
    ToolBar_FlipChildren(t.instance, AllLevels)
}

// CanFocus
func (t *TToolBar) CanFocus() bool {
    return ToolBar_CanFocus(t.instance)
}

// ContainsControl
func (t *TToolBar) ContainsControl(Control IControl) bool {
    return ToolBar_ContainsControl(t.instance, CheckPtr(Control))
}

// ControlAtPos
func (t *TToolBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(ToolBar_ControlAtPos(t.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (t *TToolBar) DisableAlign() {
    ToolBar_DisableAlign(t.instance)
}

// EnableAlign
func (t *TToolBar) EnableAlign() {
    ToolBar_EnableAlign(t.instance)
}

// FindChildControl
func (t *TToolBar) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(ToolBar_FindChildControl(t.instance, ControlName))
}

// Focused
func (t *TToolBar) Focused() bool {
    return ToolBar_Focused(t.instance)
}

// HandleAllocated
func (t *TToolBar) HandleAllocated() bool {
    return ToolBar_HandleAllocated(t.instance)
}

// InsertControl
func (t *TToolBar) InsertControl(AControl IControl) {
    ToolBar_InsertControl(t.instance, CheckPtr(AControl))
}

// Invalidate
func (t *TToolBar) Invalidate() {
    ToolBar_Invalidate(t.instance)
}

// PaintTo
func (t *TToolBar) PaintTo(DC HDC, X int32, Y int32) {
    ToolBar_PaintTo(t.instance, DC , X , Y)
}

// RemoveControl
func (t *TToolBar) RemoveControl(AControl IControl) {
    ToolBar_RemoveControl(t.instance, CheckPtr(AControl))
}

// Realign
func (t *TToolBar) Realign() {
    ToolBar_Realign(t.instance)
}

// Repaint
func (t *TToolBar) Repaint() {
    ToolBar_Repaint(t.instance)
}

// ScaleBy
func (t *TToolBar) ScaleBy(M int32, D int32) {
    ToolBar_ScaleBy(t.instance, M , D)
}

// ScrollBy
func (t *TToolBar) ScrollBy(DeltaX int32, DeltaY int32) {
    ToolBar_ScrollBy(t.instance, DeltaX , DeltaY)
}

// SetBounds
func (t *TToolBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ToolBar_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (t *TToolBar) SetFocus() {
    ToolBar_SetFocus(t.instance)
}

// Update
func (t *TToolBar) Update() {
    ToolBar_Update(t.instance)
}

// UpdateControlState
func (t *TToolBar) UpdateControlState() {
    ToolBar_UpdateControlState(t.instance)
}

// BringToFront
func (t *TToolBar) BringToFront() {
    ToolBar_BringToFront(t.instance)
}

// ClientToScreen
func (t *TToolBar) ClientToScreen(Point TPoint) TPoint {
    return ToolBar_ClientToScreen(t.instance, Point)
}

// ClientToParent
func (t *TToolBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ToolBar_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// Dragging
func (t *TToolBar) Dragging() bool {
    return ToolBar_Dragging(t.instance)
}

// HasParent
func (t *TToolBar) HasParent() bool {
    return ToolBar_HasParent(t.instance)
}

// Hide
func (t *TToolBar) Hide() {
    ToolBar_Hide(t.instance)
}

// Perform
func (t *TToolBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ToolBar_Perform(t.instance, Msg , WParam , LParam)
}

// Refresh
func (t *TToolBar) Refresh() {
    ToolBar_Refresh(t.instance)
}

// ScreenToClient
func (t *TToolBar) ScreenToClient(Point TPoint) TPoint {
    return ToolBar_ScreenToClient(t.instance, Point)
}

// ParentToClient
func (t *TToolBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ToolBar_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (t *TToolBar) SendToBack() {
    ToolBar_SendToBack(t.instance)
}

// Show
func (t *TToolBar) Show() {
    ToolBar_Show(t.instance)
}

// GetTextBuf
func (t *TToolBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ToolBar_GetTextBuf(t.instance, Buffer , BufSize)
}

// GetTextLen
func (t *TToolBar) GetTextLen() int32 {
    return ToolBar_GetTextLen(t.instance)
}

// SetTextBuf
func (t *TToolBar) SetTextBuf(Buffer string) {
    ToolBar_SetTextBuf(t.instance, Buffer)
}

// FindComponent
func (t *TToolBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ToolBar_FindComponent(t.instance, AName))
}

// GetNamePath
func (t *TToolBar) GetNamePath() string {
    return ToolBar_GetNamePath(t.instance)
}

// Assign
func (t *TToolBar) Assign(Source IObject) {
    ToolBar_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TToolBar) DisposeOf() {
    ToolBar_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TToolBar) ClassType() TClass {
    return ToolBar_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TToolBar) ClassName() string {
    return ToolBar_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TToolBar) InstanceSize() int32 {
    return ToolBar_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TToolBar) InheritsFrom(AClass TClass) bool {
    return ToolBar_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TToolBar) Equals(Obj IObject) bool {
    return ToolBar_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TToolBar) GetHashCode() int32 {
    return ToolBar_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TToolBar) ToString() string {
    return ToolBar_ToString(t.instance)
}

// ButtonCount
func (t *TToolBar) ButtonCount() int32 {
    return ToolBar_GetButtonCount(t.instance)
}

// Canvas
func (t *TToolBar) Canvas() *TCanvas {
    return CanvasFromInst(ToolBar_GetCanvas(t.instance))
}

// CustomizeKeyName
func (t *TToolBar) CustomizeKeyName() string {
    return ToolBar_GetCustomizeKeyName(t.instance)
}

// SetCustomizeKeyName
func (t *TToolBar) SetCustomizeKeyName(value string) {
    ToolBar_SetCustomizeKeyName(t.instance, value)
}

// CustomizeValueName
func (t *TToolBar) CustomizeValueName() string {
    return ToolBar_GetCustomizeValueName(t.instance)
}

// SetCustomizeValueName
func (t *TToolBar) SetCustomizeValueName(value string) {
    ToolBar_SetCustomizeValueName(t.instance, value)
}

// RowCount
func (t *TToolBar) RowCount() int32 {
    return ToolBar_GetRowCount(t.instance)
}

// Align
func (t *TToolBar) Align() TAlign {
    return ToolBar_GetAlign(t.instance)
}

// SetAlign
func (t *TToolBar) SetAlign(value TAlign) {
    ToolBar_SetAlign(t.instance, value)
}

// Anchors
func (t *TToolBar) Anchors() TAnchors {
    return ToolBar_GetAnchors(t.instance)
}

// SetAnchors
func (t *TToolBar) SetAnchors(value TAnchors) {
    ToolBar_SetAnchors(t.instance, value)
}

// AutoSize
func (t *TToolBar) AutoSize() bool {
    return ToolBar_GetAutoSize(t.instance)
}

// SetAutoSize
func (t *TToolBar) SetAutoSize(value bool) {
    ToolBar_SetAutoSize(t.instance, value)
}

// BorderWidth
func (t *TToolBar) BorderWidth() int32 {
    return ToolBar_GetBorderWidth(t.instance)
}

// SetBorderWidth
func (t *TToolBar) SetBorderWidth(value int32) {
    ToolBar_SetBorderWidth(t.instance, value)
}

// ButtonHeight
func (t *TToolBar) ButtonHeight() int32 {
    return ToolBar_GetButtonHeight(t.instance)
}

// SetButtonHeight
func (t *TToolBar) SetButtonHeight(value int32) {
    ToolBar_SetButtonHeight(t.instance, value)
}

// ButtonWidth
func (t *TToolBar) ButtonWidth() int32 {
    return ToolBar_GetButtonWidth(t.instance)
}

// SetButtonWidth
func (t *TToolBar) SetButtonWidth(value int32) {
    ToolBar_SetButtonWidth(t.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (t *TToolBar) Caption() string {
    return ToolBar_GetCaption(t.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (t *TToolBar) SetCaption(value string) {
    ToolBar_SetCaption(t.instance, value)
}

// Color
func (t *TToolBar) Color() TColor {
    return ToolBar_GetColor(t.instance)
}

// SetColor
func (t *TToolBar) SetColor(value TColor) {
    ToolBar_SetColor(t.instance, value)
}

// DoubleBuffered
func (t *TToolBar) DoubleBuffered() bool {
    return ToolBar_GetDoubleBuffered(t.instance)
}

// SetDoubleBuffered
func (t *TToolBar) SetDoubleBuffered(value bool) {
    ToolBar_SetDoubleBuffered(t.instance, value)
}

// DockSite
func (t *TToolBar) DockSite() bool {
    return ToolBar_GetDockSite(t.instance)
}

// SetDockSite
func (t *TToolBar) SetDockSite(value bool) {
    ToolBar_SetDockSite(t.instance, value)
}

// DragCursor
func (t *TToolBar) DragCursor() TCursor {
    return ToolBar_GetDragCursor(t.instance)
}

// SetDragCursor
func (t *TToolBar) SetDragCursor(value TCursor) {
    ToolBar_SetDragCursor(t.instance, value)
}

// DragKind
func (t *TToolBar) DragKind() TDragKind {
    return ToolBar_GetDragKind(t.instance)
}

// SetDragKind
func (t *TToolBar) SetDragKind(value TDragKind) {
    ToolBar_SetDragKind(t.instance, value)
}

// DragMode
func (t *TToolBar) DragMode() TDragMode {
    return ToolBar_GetDragMode(t.instance)
}

// SetDragMode
func (t *TToolBar) SetDragMode(value TDragMode) {
    ToolBar_SetDragMode(t.instance, value)
}

// DrawingStyle
func (t *TToolBar) DrawingStyle() TTBDrawingStyle {
    return ToolBar_GetDrawingStyle(t.instance)
}

// SetDrawingStyle
func (t *TToolBar) SetDrawingStyle(value TTBDrawingStyle) {
    ToolBar_SetDrawingStyle(t.instance, value)
}

// EdgeBorders
func (t *TToolBar) EdgeBorders() TEdgeBorders {
    return ToolBar_GetEdgeBorders(t.instance)
}

// SetEdgeBorders
func (t *TToolBar) SetEdgeBorders(value TEdgeBorders) {
    ToolBar_SetEdgeBorders(t.instance, value)
}

// EdgeInner
func (t *TToolBar) EdgeInner() TEdgeStyle {
    return ToolBar_GetEdgeInner(t.instance)
}

// SetEdgeInner
func (t *TToolBar) SetEdgeInner(value TEdgeStyle) {
    ToolBar_SetEdgeInner(t.instance, value)
}

// EdgeOuter
func (t *TToolBar) EdgeOuter() TEdgeStyle {
    return ToolBar_GetEdgeOuter(t.instance)
}

// SetEdgeOuter
func (t *TToolBar) SetEdgeOuter(value TEdgeStyle) {
    ToolBar_SetEdgeOuter(t.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TToolBar) Enabled() bool {
    return ToolBar_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TToolBar) SetEnabled(value bool) {
    ToolBar_SetEnabled(t.instance, value)
}

// Flat
func (t *TToolBar) Flat() bool {
    return ToolBar_GetFlat(t.instance)
}

// SetFlat
func (t *TToolBar) SetFlat(value bool) {
    ToolBar_SetFlat(t.instance, value)
}

// Font
func (t *TToolBar) Font() *TFont {
    return FontFromInst(ToolBar_GetFont(t.instance))
}

// SetFont
func (t *TToolBar) SetFont(value *TFont) {
    ToolBar_SetFont(t.instance, CheckPtr(value))
}

// GradientEndColor
func (t *TToolBar) GradientEndColor() TColor {
    return ToolBar_GetGradientEndColor(t.instance)
}

// SetGradientEndColor
func (t *TToolBar) SetGradientEndColor(value TColor) {
    ToolBar_SetGradientEndColor(t.instance, value)
}

// GradientStartColor
func (t *TToolBar) GradientStartColor() TColor {
    return ToolBar_GetGradientStartColor(t.instance)
}

// SetGradientStartColor
func (t *TToolBar) SetGradientStartColor(value TColor) {
    ToolBar_SetGradientStartColor(t.instance, value)
}

// Height
func (t *TToolBar) Height() int32 {
    return ToolBar_GetHeight(t.instance)
}

// SetHeight
func (t *TToolBar) SetHeight(value int32) {
    ToolBar_SetHeight(t.instance, value)
}

// HideClippedButtons
func (t *TToolBar) HideClippedButtons() bool {
    return ToolBar_GetHideClippedButtons(t.instance)
}

// SetHideClippedButtons
func (t *TToolBar) SetHideClippedButtons(value bool) {
    ToolBar_SetHideClippedButtons(t.instance, value)
}

// HotImages
func (t *TToolBar) HotImages() *TImageList {
    return ImageListFromInst(ToolBar_GetHotImages(t.instance))
}

// SetHotImages
func (t *TToolBar) SetHotImages(value IComponent) {
    ToolBar_SetHotImages(t.instance, CheckPtr(value))
}

// Images
func (t *TToolBar) Images() *TImageList {
    return ImageListFromInst(ToolBar_GetImages(t.instance))
}

// SetImages
func (t *TToolBar) SetImages(value IComponent) {
    ToolBar_SetImages(t.instance, CheckPtr(value))
}

// Indent
func (t *TToolBar) Indent() int32 {
    return ToolBar_GetIndent(t.instance)
}

// SetIndent
func (t *TToolBar) SetIndent(value int32) {
    ToolBar_SetIndent(t.instance, value)
}

// List
func (t *TToolBar) List() bool {
    return ToolBar_GetList(t.instance)
}

// SetList
func (t *TToolBar) SetList(value bool) {
    ToolBar_SetList(t.instance, value)
}

// Menu
func (t *TToolBar) Menu() *TMainMenu {
    return MainMenuFromInst(ToolBar_GetMenu(t.instance))
}

// SetMenu
func (t *TToolBar) SetMenu(value IComponent) {
    ToolBar_SetMenu(t.instance, CheckPtr(value))
}

// GradientDirection
func (t *TToolBar) GradientDirection() TGradientDirection {
    return ToolBar_GetGradientDirection(t.instance)
}

// SetGradientDirection
func (t *TToolBar) SetGradientDirection(value TGradientDirection) {
    ToolBar_SetGradientDirection(t.instance, value)
}

// GradientDrawingOptions
func (t *TToolBar) GradientDrawingOptions() TTBGradientDrawingOptions {
    return ToolBar_GetGradientDrawingOptions(t.instance)
}

// SetGradientDrawingOptions
func (t *TToolBar) SetGradientDrawingOptions(value TTBGradientDrawingOptions) {
    ToolBar_SetGradientDrawingOptions(t.instance, value)
}

// ParentColor
func (t *TToolBar) ParentColor() bool {
    return ToolBar_GetParentColor(t.instance)
}

// SetParentColor
func (t *TToolBar) SetParentColor(value bool) {
    ToolBar_SetParentColor(t.instance, value)
}

// ParentDoubleBuffered
func (t *TToolBar) ParentDoubleBuffered() bool {
    return ToolBar_GetParentDoubleBuffered(t.instance)
}

// SetParentDoubleBuffered
func (t *TToolBar) SetParentDoubleBuffered(value bool) {
    ToolBar_SetParentDoubleBuffered(t.instance, value)
}

// ParentFont
func (t *TToolBar) ParentFont() bool {
    return ToolBar_GetParentFont(t.instance)
}

// SetParentFont
func (t *TToolBar) SetParentFont(value bool) {
    ToolBar_SetParentFont(t.instance, value)
}

// ParentShowHint
func (t *TToolBar) ParentShowHint() bool {
    return ToolBar_GetParentShowHint(t.instance)
}

// SetParentShowHint
func (t *TToolBar) SetParentShowHint(value bool) {
    ToolBar_SetParentShowHint(t.instance, value)
}

// PopupMenu
func (t *TToolBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ToolBar_GetPopupMenu(t.instance))
}

// SetPopupMenu
func (t *TToolBar) SetPopupMenu(value IComponent) {
    ToolBar_SetPopupMenu(t.instance, CheckPtr(value))
}

// ShowCaptions
func (t *TToolBar) ShowCaptions() bool {
    return ToolBar_GetShowCaptions(t.instance)
}

// SetShowCaptions
func (t *TToolBar) SetShowCaptions(value bool) {
    ToolBar_SetShowCaptions(t.instance, value)
}

// ShowHint
func (t *TToolBar) ShowHint() bool {
    return ToolBar_GetShowHint(t.instance)
}

// SetShowHint
func (t *TToolBar) SetShowHint(value bool) {
    ToolBar_SetShowHint(t.instance, value)
}

// TabOrder
func (t *TToolBar) TabOrder() TTabOrder {
    return ToolBar_GetTabOrder(t.instance)
}

// SetTabOrder
func (t *TToolBar) SetTabOrder(value TTabOrder) {
    ToolBar_SetTabOrder(t.instance, value)
}

// TabStop
func (t *TToolBar) TabStop() bool {
    return ToolBar_GetTabStop(t.instance)
}

// SetTabStop
func (t *TToolBar) SetTabStop(value bool) {
    ToolBar_SetTabStop(t.instance, value)
}

// Transparent
func (t *TToolBar) Transparent() bool {
    return ToolBar_GetTransparent(t.instance)
}

// SetTransparent
func (t *TToolBar) SetTransparent(value bool) {
    ToolBar_SetTransparent(t.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TToolBar) Visible() bool {
    return ToolBar_GetVisible(t.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TToolBar) SetVisible(value bool) {
    ToolBar_SetVisible(t.instance, value)
}

// StyleElements
func (t *TToolBar) StyleElements() TStyleElements {
    return ToolBar_GetStyleElements(t.instance)
}

// SetStyleElements
func (t *TToolBar) SetStyleElements(value TStyleElements) {
    ToolBar_SetStyleElements(t.instance, value)
}

// Wrapable
func (t *TToolBar) Wrapable() bool {
    return ToolBar_GetWrapable(t.instance)
}

// SetWrapable
func (t *TToolBar) SetWrapable(value bool) {
    ToolBar_SetWrapable(t.instance, value)
}

// SetOnAdvancedCustomDraw
func (t *TToolBar) SetOnAdvancedCustomDraw(fn TTBAdvancedCustomDrawEvent) {
    ToolBar_SetOnAdvancedCustomDraw(t.instance, fn)
}

// SetOnAdvancedCustomDrawButton
func (t *TToolBar) SetOnAdvancedCustomDrawButton(fn TTBAdvancedCustomDrawBtnEvent) {
    ToolBar_SetOnAdvancedCustomDrawButton(t.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (t *TToolBar) SetOnClick(fn TNotifyEvent) {
    ToolBar_SetOnClick(t.instance, fn)
}

// SetOnContextPopup
func (t *TToolBar) SetOnContextPopup(fn TContextPopupEvent) {
    ToolBar_SetOnContextPopup(t.instance, fn)
}

// SetOnDblClick
func (t *TToolBar) SetOnDblClick(fn TNotifyEvent) {
    ToolBar_SetOnDblClick(t.instance, fn)
}

// SetOnDockDrop
func (t *TToolBar) SetOnDockDrop(fn TDockDropEvent) {
    ToolBar_SetOnDockDrop(t.instance, fn)
}

// SetOnDragDrop
func (t *TToolBar) SetOnDragDrop(fn TDragDropEvent) {
    ToolBar_SetOnDragDrop(t.instance, fn)
}

// SetOnDragOver
func (t *TToolBar) SetOnDragOver(fn TDragOverEvent) {
    ToolBar_SetOnDragOver(t.instance, fn)
}

// SetOnEndDock
func (t *TToolBar) SetOnEndDock(fn TEndDragEvent) {
    ToolBar_SetOnEndDock(t.instance, fn)
}

// SetOnEndDrag
func (t *TToolBar) SetOnEndDrag(fn TEndDragEvent) {
    ToolBar_SetOnEndDrag(t.instance, fn)
}

// SetOnEnter
func (t *TToolBar) SetOnEnter(fn TNotifyEvent) {
    ToolBar_SetOnEnter(t.instance, fn)
}

// SetOnExit
func (t *TToolBar) SetOnExit(fn TNotifyEvent) {
    ToolBar_SetOnExit(t.instance, fn)
}

// SetOnGetSiteInfo
func (t *TToolBar) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    ToolBar_SetOnGetSiteInfo(t.instance, fn)
}

// SetOnMouseDown
func (t *TToolBar) SetOnMouseDown(fn TMouseEvent) {
    ToolBar_SetOnMouseDown(t.instance, fn)
}

// SetOnMouseEnter
func (t *TToolBar) SetOnMouseEnter(fn TNotifyEvent) {
    ToolBar_SetOnMouseEnter(t.instance, fn)
}

// SetOnMouseLeave
func (t *TToolBar) SetOnMouseLeave(fn TNotifyEvent) {
    ToolBar_SetOnMouseLeave(t.instance, fn)
}

// SetOnMouseMove
func (t *TToolBar) SetOnMouseMove(fn TMouseMoveEvent) {
    ToolBar_SetOnMouseMove(t.instance, fn)
}

// SetOnMouseUp
func (t *TToolBar) SetOnMouseUp(fn TMouseEvent) {
    ToolBar_SetOnMouseUp(t.instance, fn)
}

// SetOnResize
func (t *TToolBar) SetOnResize(fn TNotifyEvent) {
    ToolBar_SetOnResize(t.instance, fn)
}

// SetOnStartDock
func (t *TToolBar) SetOnStartDock(fn TStartDockEvent) {
    ToolBar_SetOnStartDock(t.instance, fn)
}

// SetOnUnDock
func (t *TToolBar) SetOnUnDock(fn TUnDockEvent) {
    ToolBar_SetOnUnDock(t.instance, fn)
}

// DockClientCount
func (t *TToolBar) DockClientCount() int32 {
    return ToolBar_GetDockClientCount(t.instance)
}

// AlignDisabled
func (t *TToolBar) AlignDisabled() bool {
    return ToolBar_GetAlignDisabled(t.instance)
}

// MouseInClient
func (t *TToolBar) MouseInClient() bool {
    return ToolBar_GetMouseInClient(t.instance)
}

// VisibleDockClientCount
func (t *TToolBar) VisibleDockClientCount() int32 {
    return ToolBar_GetVisibleDockClientCount(t.instance)
}

// Brush
func (t *TToolBar) Brush() *TBrush {
    return BrushFromInst(ToolBar_GetBrush(t.instance))
}

// ControlCount
func (t *TToolBar) ControlCount() int32 {
    return ToolBar_GetControlCount(t.instance)
}

// Handle
func (t *TToolBar) Handle() HWND {
    return ToolBar_GetHandle(t.instance)
}

// ParentWindow
func (t *TToolBar) ParentWindow() HWND {
    return ToolBar_GetParentWindow(t.instance)
}

// SetParentWindow
func (t *TToolBar) SetParentWindow(value HWND) {
    ToolBar_SetParentWindow(t.instance, value)
}

// UseDockManager
func (t *TToolBar) UseDockManager() bool {
    return ToolBar_GetUseDockManager(t.instance)
}

// SetUseDockManager
func (t *TToolBar) SetUseDockManager(value bool) {
    ToolBar_SetUseDockManager(t.instance, value)
}

// Action
func (t *TToolBar) Action() *TAction {
    return ActionFromInst(ToolBar_GetAction(t.instance))
}

// SetAction
func (t *TToolBar) SetAction(value IComponent) {
    ToolBar_SetAction(t.instance, CheckPtr(value))
}

// BiDiMode
func (t *TToolBar) BiDiMode() TBiDiMode {
    return ToolBar_GetBiDiMode(t.instance)
}

// SetBiDiMode
func (t *TToolBar) SetBiDiMode(value TBiDiMode) {
    ToolBar_SetBiDiMode(t.instance, value)
}

// BoundsRect
func (t *TToolBar) BoundsRect() TRect {
    return ToolBar_GetBoundsRect(t.instance)
}

// SetBoundsRect
func (t *TToolBar) SetBoundsRect(value TRect) {
    ToolBar_SetBoundsRect(t.instance, value)
}

// ClientHeight
func (t *TToolBar) ClientHeight() int32 {
    return ToolBar_GetClientHeight(t.instance)
}

// SetClientHeight
func (t *TToolBar) SetClientHeight(value int32) {
    ToolBar_SetClientHeight(t.instance, value)
}

// ClientOrigin
func (t *TToolBar) ClientOrigin() TPoint {
    return ToolBar_GetClientOrigin(t.instance)
}

// ClientRect
func (t *TToolBar) ClientRect() TRect {
    return ToolBar_GetClientRect(t.instance)
}

// ClientWidth
func (t *TToolBar) ClientWidth() int32 {
    return ToolBar_GetClientWidth(t.instance)
}

// SetClientWidth
func (t *TToolBar) SetClientWidth(value int32) {
    ToolBar_SetClientWidth(t.instance, value)
}

// ControlState
func (t *TToolBar) ControlState() TControlState {
    return ToolBar_GetControlState(t.instance)
}

// SetControlState
func (t *TToolBar) SetControlState(value TControlState) {
    ToolBar_SetControlState(t.instance, value)
}

// ControlStyle
func (t *TToolBar) ControlStyle() TControlStyle {
    return ToolBar_GetControlStyle(t.instance)
}

// SetControlStyle
func (t *TToolBar) SetControlStyle(value TControlStyle) {
    ToolBar_SetControlStyle(t.instance, value)
}

// ExplicitLeft
func (t *TToolBar) ExplicitLeft() int32 {
    return ToolBar_GetExplicitLeft(t.instance)
}

// ExplicitTop
func (t *TToolBar) ExplicitTop() int32 {
    return ToolBar_GetExplicitTop(t.instance)
}

// ExplicitWidth
func (t *TToolBar) ExplicitWidth() int32 {
    return ToolBar_GetExplicitWidth(t.instance)
}

// ExplicitHeight
func (t *TToolBar) ExplicitHeight() int32 {
    return ToolBar_GetExplicitHeight(t.instance)
}

// Floating
func (t *TToolBar) Floating() bool {
    return ToolBar_GetFloating(t.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TToolBar) Parent() *TWinControl {
    return WinControlFromInst(ToolBar_GetParent(t.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (t *TToolBar) SetParent(value IWinControl) {
    ToolBar_SetParent(t.instance, CheckPtr(value))
}

// AlignWithMargins
func (t *TToolBar) AlignWithMargins() bool {
    return ToolBar_GetAlignWithMargins(t.instance)
}

// SetAlignWithMargins
func (t *TToolBar) SetAlignWithMargins(value bool) {
    ToolBar_SetAlignWithMargins(t.instance, value)
}

// Left
func (t *TToolBar) Left() int32 {
    return ToolBar_GetLeft(t.instance)
}

// SetLeft
func (t *TToolBar) SetLeft(value int32) {
    ToolBar_SetLeft(t.instance, value)
}

// Top
func (t *TToolBar) Top() int32 {
    return ToolBar_GetTop(t.instance)
}

// SetTop
func (t *TToolBar) SetTop(value int32) {
    ToolBar_SetTop(t.instance, value)
}

// Width
func (t *TToolBar) Width() int32 {
    return ToolBar_GetWidth(t.instance)
}

// SetWidth
func (t *TToolBar) SetWidth(value int32) {
    ToolBar_SetWidth(t.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (t *TToolBar) Cursor() TCursor {
    return ToolBar_GetCursor(t.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (t *TToolBar) SetCursor(value TCursor) {
    ToolBar_SetCursor(t.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (t *TToolBar) Hint() string {
    return ToolBar_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (t *TToolBar) SetHint(value string) {
    ToolBar_SetHint(t.instance, value)
}

// Margins
func (t *TToolBar) Margins() *TMargins {
    return MarginsFromInst(ToolBar_GetMargins(t.instance))
}

// SetMargins
func (t *TToolBar) SetMargins(value *TMargins) {
    ToolBar_SetMargins(t.instance, CheckPtr(value))
}

// CustomHint
func (t *TToolBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(ToolBar_GetCustomHint(t.instance))
}

// SetCustomHint
func (t *TToolBar) SetCustomHint(value IComponent) {
    ToolBar_SetCustomHint(t.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TToolBar) ComponentCount() int32 {
    return ToolBar_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TToolBar) ComponentIndex() int32 {
    return ToolBar_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TToolBar) SetComponentIndex(value int32) {
    ToolBar_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TToolBar) Owner() *TComponent {
    return ComponentFromInst(ToolBar_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TToolBar) Name() string {
    return ToolBar_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TToolBar) SetName(value string) {
    ToolBar_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TToolBar) Tag() int {
    return ToolBar_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TToolBar) SetTag(value int) {
    ToolBar_SetTag(t.instance, value)
}

// Buttons
func (t *TToolBar) Buttons(Index int32) *TToolButton {
    return ToolButtonFromInst(ToolBar_GetButtons(t.instance, Index))
}

// DockClients
func (t *TToolBar) DockClients(Index int32) *TControl {
    return ControlFromInst(ToolBar_GetDockClients(t.instance, Index))
}

// Controls
func (t *TToolBar) Controls(Index int32) *TControl {
    return ControlFromInst(ToolBar_GetControls(t.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TToolBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ToolBar_GetComponents(t.instance, AIndex))
}

