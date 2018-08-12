
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

type TCategoryPanelGroup struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCategoryPanelGroup
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCategoryPanelGroup(owner IComponent) *TCategoryPanelGroup {
    c := new(TCategoryPanelGroup)
    c.instance = CategoryPanelGroup_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CategoryPanelGroupFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CategoryPanelGroupFromInst(inst uintptr) *TCategoryPanelGroup {
    c := new(TCategoryPanelGroup)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CategoryPanelGroupFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CategoryPanelGroupFromObj(obj IObject) *TCategoryPanelGroup {
    c := new(TCategoryPanelGroup)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CategoryPanelGroupFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CategoryPanelGroupFromUnsafePointer(ptr unsafe.Pointer) *TCategoryPanelGroup {
    c := new(TCategoryPanelGroup)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCategoryPanelGroup) Free() {
    if c.instance != 0 {
        CategoryPanelGroup_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCategoryPanelGroup) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCategoryPanelGroup) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCategoryPanelGroup) IsValid() bool {
    return c.instance != 0
}

// TCategoryPanelGroupClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCategoryPanelGroupClass() TClass {
    return CategoryPanelGroup_StaticClassType()
}

// CollapseAll
func (c *TCategoryPanelGroup) CollapseAll() {
    CategoryPanelGroup_CollapseAll(c.instance)
}

// ExpandAll
func (c *TCategoryPanelGroup) ExpandAll() {
    CategoryPanelGroup_ExpandAll(c.instance)
}

// CanFocus
func (c *TCategoryPanelGroup) CanFocus() bool {
    return CategoryPanelGroup_CanFocus(c.instance)
}

// ContainsControl
func (c *TCategoryPanelGroup) ContainsControl(Control IControl) bool {
    return CategoryPanelGroup_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
func (c *TCategoryPanelGroup) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(CategoryPanelGroup_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (c *TCategoryPanelGroup) DisableAlign() {
    CategoryPanelGroup_DisableAlign(c.instance)
}

// EnableAlign
func (c *TCategoryPanelGroup) EnableAlign() {
    CategoryPanelGroup_EnableAlign(c.instance)
}

// FindChildControl
func (c *TCategoryPanelGroup) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(CategoryPanelGroup_FindChildControl(c.instance, ControlName))
}

// FlipChildren
func (c *TCategoryPanelGroup) FlipChildren(AllLevels bool) {
    CategoryPanelGroup_FlipChildren(c.instance, AllLevels)
}

// Focused
func (c *TCategoryPanelGroup) Focused() bool {
    return CategoryPanelGroup_Focused(c.instance)
}

// HandleAllocated
func (c *TCategoryPanelGroup) HandleAllocated() bool {
    return CategoryPanelGroup_HandleAllocated(c.instance)
}

// InsertControl
func (c *TCategoryPanelGroup) InsertControl(AControl IControl) {
    CategoryPanelGroup_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
func (c *TCategoryPanelGroup) Invalidate() {
    CategoryPanelGroup_Invalidate(c.instance)
}

// PaintTo
func (c *TCategoryPanelGroup) PaintTo(DC HDC, X int32, Y int32) {
    CategoryPanelGroup_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
func (c *TCategoryPanelGroup) RemoveControl(AControl IControl) {
    CategoryPanelGroup_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
func (c *TCategoryPanelGroup) Realign() {
    CategoryPanelGroup_Realign(c.instance)
}

// Repaint
func (c *TCategoryPanelGroup) Repaint() {
    CategoryPanelGroup_Repaint(c.instance)
}

// ScaleBy
func (c *TCategoryPanelGroup) ScaleBy(M int32, D int32) {
    CategoryPanelGroup_ScaleBy(c.instance, M , D)
}

// ScrollBy
func (c *TCategoryPanelGroup) ScrollBy(DeltaX int32, DeltaY int32) {
    CategoryPanelGroup_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetBounds
func (c *TCategoryPanelGroup) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CategoryPanelGroup_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (c *TCategoryPanelGroup) SetFocus() {
    CategoryPanelGroup_SetFocus(c.instance)
}

// Update
func (c *TCategoryPanelGroup) Update() {
    CategoryPanelGroup_Update(c.instance)
}

// UpdateControlState
func (c *TCategoryPanelGroup) UpdateControlState() {
    CategoryPanelGroup_UpdateControlState(c.instance)
}

// BringToFront
func (c *TCategoryPanelGroup) BringToFront() {
    CategoryPanelGroup_BringToFront(c.instance)
}

// ClientToScreen
func (c *TCategoryPanelGroup) ClientToScreen(Point TPoint) TPoint {
    return CategoryPanelGroup_ClientToScreen(c.instance, Point)
}

// ClientToParent
func (c *TCategoryPanelGroup) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CategoryPanelGroup_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
func (c *TCategoryPanelGroup) Dragging() bool {
    return CategoryPanelGroup_Dragging(c.instance)
}

// HasParent
func (c *TCategoryPanelGroup) HasParent() bool {
    return CategoryPanelGroup_HasParent(c.instance)
}

// Hide
func (c *TCategoryPanelGroup) Hide() {
    CategoryPanelGroup_Hide(c.instance)
}

// Perform
func (c *TCategoryPanelGroup) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CategoryPanelGroup_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
func (c *TCategoryPanelGroup) Refresh() {
    CategoryPanelGroup_Refresh(c.instance)
}

// ScreenToClient
func (c *TCategoryPanelGroup) ScreenToClient(Point TPoint) TPoint {
    return CategoryPanelGroup_ScreenToClient(c.instance, Point)
}

// ParentToClient
func (c *TCategoryPanelGroup) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CategoryPanelGroup_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (c *TCategoryPanelGroup) SendToBack() {
    CategoryPanelGroup_SendToBack(c.instance)
}

// Show
func (c *TCategoryPanelGroup) Show() {
    CategoryPanelGroup_Show(c.instance)
}

// GetTextBuf
func (c *TCategoryPanelGroup) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CategoryPanelGroup_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
func (c *TCategoryPanelGroup) GetTextLen() int32 {
    return CategoryPanelGroup_GetTextLen(c.instance)
}

// SetTextBuf
func (c *TCategoryPanelGroup) SetTextBuf(Buffer string) {
    CategoryPanelGroup_SetTextBuf(c.instance, Buffer)
}

// FindComponent
func (c *TCategoryPanelGroup) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CategoryPanelGroup_FindComponent(c.instance, AName))
}

// GetNamePath
func (c *TCategoryPanelGroup) GetNamePath() string {
    return CategoryPanelGroup_GetNamePath(c.instance)
}

// Assign
func (c *TCategoryPanelGroup) Assign(Source IObject) {
    CategoryPanelGroup_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCategoryPanelGroup) DisposeOf() {
    CategoryPanelGroup_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCategoryPanelGroup) ClassType() TClass {
    return CategoryPanelGroup_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCategoryPanelGroup) ClassName() string {
    return CategoryPanelGroup_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCategoryPanelGroup) InstanceSize() int32 {
    return CategoryPanelGroup_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCategoryPanelGroup) InheritsFrom(AClass TClass) bool {
    return CategoryPanelGroup_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCategoryPanelGroup) Equals(Obj IObject) bool {
    return CategoryPanelGroup_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCategoryPanelGroup) GetHashCode() int32 {
    return CategoryPanelGroup_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCategoryPanelGroup) ToString() string {
    return CategoryPanelGroup_ToString(c.instance)
}

// Align
func (c *TCategoryPanelGroup) Align() TAlign {
    return CategoryPanelGroup_GetAlign(c.instance)
}

// SetAlign
func (c *TCategoryPanelGroup) SetAlign(value TAlign) {
    CategoryPanelGroup_SetAlign(c.instance, value)
}

// Anchors
func (c *TCategoryPanelGroup) Anchors() TAnchors {
    return CategoryPanelGroup_GetAnchors(c.instance)
}

// SetAnchors
func (c *TCategoryPanelGroup) SetAnchors(value TAnchors) {
    CategoryPanelGroup_SetAnchors(c.instance, value)
}

// BevelEdges
func (c *TCategoryPanelGroup) BevelEdges() TBevelEdges {
    return CategoryPanelGroup_GetBevelEdges(c.instance)
}

// SetBevelEdges
func (c *TCategoryPanelGroup) SetBevelEdges(value TBevelEdges) {
    CategoryPanelGroup_SetBevelEdges(c.instance, value)
}

// BevelInner
func (c *TCategoryPanelGroup) BevelInner() TBevelCut {
    return CategoryPanelGroup_GetBevelInner(c.instance)
}

// SetBevelInner
func (c *TCategoryPanelGroup) SetBevelInner(value TBevelCut) {
    CategoryPanelGroup_SetBevelInner(c.instance, value)
}

// BevelOuter
func (c *TCategoryPanelGroup) BevelOuter() TBevelCut {
    return CategoryPanelGroup_GetBevelOuter(c.instance)
}

// SetBevelOuter
func (c *TCategoryPanelGroup) SetBevelOuter(value TBevelCut) {
    CategoryPanelGroup_SetBevelOuter(c.instance, value)
}

// BevelKind
func (c *TCategoryPanelGroup) BevelKind() TBevelKind {
    return CategoryPanelGroup_GetBevelKind(c.instance)
}

// SetBevelKind
func (c *TCategoryPanelGroup) SetBevelKind(value TBevelKind) {
    CategoryPanelGroup_SetBevelKind(c.instance, value)
}

// BiDiMode
func (c *TCategoryPanelGroup) BiDiMode() TBiDiMode {
    return CategoryPanelGroup_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TCategoryPanelGroup) SetBiDiMode(value TBiDiMode) {
    CategoryPanelGroup_SetBiDiMode(c.instance, value)
}

// ChevronAlignment
func (c *TCategoryPanelGroup) ChevronAlignment() TAlignment {
    return CategoryPanelGroup_GetChevronAlignment(c.instance)
}

// SetChevronAlignment
func (c *TCategoryPanelGroup) SetChevronAlignment(value TAlignment) {
    CategoryPanelGroup_SetChevronAlignment(c.instance, value)
}

// ChevronColor
func (c *TCategoryPanelGroup) ChevronColor() TColor {
    return CategoryPanelGroup_GetChevronColor(c.instance)
}

// SetChevronColor
func (c *TCategoryPanelGroup) SetChevronColor(value TColor) {
    CategoryPanelGroup_SetChevronColor(c.instance, value)
}

// ChevronHotColor
func (c *TCategoryPanelGroup) ChevronHotColor() TColor {
    return CategoryPanelGroup_GetChevronHotColor(c.instance)
}

// SetChevronHotColor
func (c *TCategoryPanelGroup) SetChevronHotColor(value TColor) {
    CategoryPanelGroup_SetChevronHotColor(c.instance, value)
}

// DockSite
func (c *TCategoryPanelGroup) DockSite() bool {
    return CategoryPanelGroup_GetDockSite(c.instance)
}

// SetDockSite
func (c *TCategoryPanelGroup) SetDockSite(value bool) {
    CategoryPanelGroup_SetDockSite(c.instance, value)
}

// DoubleBuffered
func (c *TCategoryPanelGroup) DoubleBuffered() bool {
    return CategoryPanelGroup_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
func (c *TCategoryPanelGroup) SetDoubleBuffered(value bool) {
    CategoryPanelGroup_SetDoubleBuffered(c.instance, value)
}

// DragCursor
func (c *TCategoryPanelGroup) DragCursor() TCursor {
    return CategoryPanelGroup_GetDragCursor(c.instance)
}

// SetDragCursor
func (c *TCategoryPanelGroup) SetDragCursor(value TCursor) {
    CategoryPanelGroup_SetDragCursor(c.instance, value)
}

// DragKind
func (c *TCategoryPanelGroup) DragKind() TDragKind {
    return CategoryPanelGroup_GetDragKind(c.instance)
}

// SetDragKind
func (c *TCategoryPanelGroup) SetDragKind(value TDragKind) {
    CategoryPanelGroup_SetDragKind(c.instance, value)
}

// DragMode
func (c *TCategoryPanelGroup) DragMode() TDragMode {
    return CategoryPanelGroup_GetDragMode(c.instance)
}

// SetDragMode
func (c *TCategoryPanelGroup) SetDragMode(value TDragMode) {
    CategoryPanelGroup_SetDragMode(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TCategoryPanelGroup) Enabled() bool {
    return CategoryPanelGroup_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TCategoryPanelGroup) SetEnabled(value bool) {
    CategoryPanelGroup_SetEnabled(c.instance, value)
}

// Color
func (c *TCategoryPanelGroup) Color() TColor {
    return CategoryPanelGroup_GetColor(c.instance)
}

// SetColor
func (c *TCategoryPanelGroup) SetColor(value TColor) {
    CategoryPanelGroup_SetColor(c.instance, value)
}

// Font
func (c *TCategoryPanelGroup) Font() *TFont {
    return FontFromInst(CategoryPanelGroup_GetFont(c.instance))
}

// SetFont
func (c *TCategoryPanelGroup) SetFont(value *TFont) {
    CategoryPanelGroup_SetFont(c.instance, CheckPtr(value))
}

// GradientBaseColor
func (c *TCategoryPanelGroup) GradientBaseColor() TColor {
    return CategoryPanelGroup_GetGradientBaseColor(c.instance)
}

// SetGradientBaseColor
func (c *TCategoryPanelGroup) SetGradientBaseColor(value TColor) {
    CategoryPanelGroup_SetGradientBaseColor(c.instance, value)
}

// GradientColor
func (c *TCategoryPanelGroup) GradientColor() TColor {
    return CategoryPanelGroup_GetGradientColor(c.instance)
}

// SetGradientColor
func (c *TCategoryPanelGroup) SetGradientColor(value TColor) {
    CategoryPanelGroup_SetGradientColor(c.instance, value)
}

// GradientDirection
func (c *TCategoryPanelGroup) GradientDirection() TGradientDirection {
    return CategoryPanelGroup_GetGradientDirection(c.instance)
}

// SetGradientDirection
func (c *TCategoryPanelGroup) SetGradientDirection(value TGradientDirection) {
    CategoryPanelGroup_SetGradientDirection(c.instance, value)
}

// HeaderAlignment
func (c *TCategoryPanelGroup) HeaderAlignment() TAlignment {
    return CategoryPanelGroup_GetHeaderAlignment(c.instance)
}

// SetHeaderAlignment
func (c *TCategoryPanelGroup) SetHeaderAlignment(value TAlignment) {
    CategoryPanelGroup_SetHeaderAlignment(c.instance, value)
}

// HeaderFont
func (c *TCategoryPanelGroup) HeaderFont() *TFont {
    return FontFromInst(CategoryPanelGroup_GetHeaderFont(c.instance))
}

// SetHeaderFont
func (c *TCategoryPanelGroup) SetHeaderFont(value *TFont) {
    CategoryPanelGroup_SetHeaderFont(c.instance, CheckPtr(value))
}

// HeaderHeight
func (c *TCategoryPanelGroup) HeaderHeight() int32 {
    return CategoryPanelGroup_GetHeaderHeight(c.instance)
}

// SetHeaderHeight
func (c *TCategoryPanelGroup) SetHeaderHeight(value int32) {
    CategoryPanelGroup_SetHeaderHeight(c.instance, value)
}

// HeaderImage
func (c *TCategoryPanelGroup) HeaderImage() *TPicture {
    return PictureFromInst(CategoryPanelGroup_GetHeaderImage(c.instance))
}

// SetHeaderImage
func (c *TCategoryPanelGroup) SetHeaderImage(value *TPicture) {
    CategoryPanelGroup_SetHeaderImage(c.instance, CheckPtr(value))
}

// HeaderStyle
func (c *TCategoryPanelGroup) HeaderStyle() THeaderStyle {
    return CategoryPanelGroup_GetHeaderStyle(c.instance)
}

// SetHeaderStyle
func (c *TCategoryPanelGroup) SetHeaderStyle(value THeaderStyle) {
    CategoryPanelGroup_SetHeaderStyle(c.instance, value)
}

// Height
func (c *TCategoryPanelGroup) Height() int32 {
    return CategoryPanelGroup_GetHeight(c.instance)
}

// SetHeight
func (c *TCategoryPanelGroup) SetHeight(value int32) {
    CategoryPanelGroup_SetHeight(c.instance, value)
}

// Images
func (c *TCategoryPanelGroup) Images() *TImageList {
    return ImageListFromInst(CategoryPanelGroup_GetImages(c.instance))
}

// SetImages
func (c *TCategoryPanelGroup) SetImages(value IComponent) {
    CategoryPanelGroup_SetImages(c.instance, CheckPtr(value))
}

// ParentBackground
func (c *TCategoryPanelGroup) ParentBackground() bool {
    return CategoryPanelGroup_GetParentBackground(c.instance)
}

// SetParentBackground
func (c *TCategoryPanelGroup) SetParentBackground(value bool) {
    CategoryPanelGroup_SetParentBackground(c.instance, value)
}

// ParentColor
func (c *TCategoryPanelGroup) ParentColor() bool {
    return CategoryPanelGroup_GetParentColor(c.instance)
}

// SetParentColor
func (c *TCategoryPanelGroup) SetParentColor(value bool) {
    CategoryPanelGroup_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TCategoryPanelGroup) ParentCtl3D() bool {
    return CategoryPanelGroup_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TCategoryPanelGroup) SetParentCtl3D(value bool) {
    CategoryPanelGroup_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
func (c *TCategoryPanelGroup) ParentDoubleBuffered() bool {
    return CategoryPanelGroup_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
func (c *TCategoryPanelGroup) SetParentDoubleBuffered(value bool) {
    CategoryPanelGroup_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
func (c *TCategoryPanelGroup) ParentFont() bool {
    return CategoryPanelGroup_GetParentFont(c.instance)
}

// SetParentFont
func (c *TCategoryPanelGroup) SetParentFont(value bool) {
    CategoryPanelGroup_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TCategoryPanelGroup) ParentShowHint() bool {
    return CategoryPanelGroup_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TCategoryPanelGroup) SetParentShowHint(value bool) {
    CategoryPanelGroup_SetParentShowHint(c.instance, value)
}

// PopupMenu
func (c *TCategoryPanelGroup) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CategoryPanelGroup_GetPopupMenu(c.instance))
}

// SetPopupMenu
func (c *TCategoryPanelGroup) SetPopupMenu(value IComponent) {
    CategoryPanelGroup_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
func (c *TCategoryPanelGroup) ShowHint() bool {
    return CategoryPanelGroup_GetShowHint(c.instance)
}

// SetShowHint
func (c *TCategoryPanelGroup) SetShowHint(value bool) {
    CategoryPanelGroup_SetShowHint(c.instance, value)
}

// TabOrder
func (c *TCategoryPanelGroup) TabOrder() TTabOrder {
    return CategoryPanelGroup_GetTabOrder(c.instance)
}

// SetTabOrder
func (c *TCategoryPanelGroup) SetTabOrder(value TTabOrder) {
    CategoryPanelGroup_SetTabOrder(c.instance, value)
}

// TabStop
func (c *TCategoryPanelGroup) TabStop() bool {
    return CategoryPanelGroup_GetTabStop(c.instance)
}

// SetTabStop
func (c *TCategoryPanelGroup) SetTabStop(value bool) {
    CategoryPanelGroup_SetTabStop(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TCategoryPanelGroup) Visible() bool {
    return CategoryPanelGroup_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TCategoryPanelGroup) SetVisible(value bool) {
    CategoryPanelGroup_SetVisible(c.instance, value)
}

// StyleElements
func (c *TCategoryPanelGroup) StyleElements() TStyleElements {
    return CategoryPanelGroup_GetStyleElements(c.instance)
}

// SetStyleElements
func (c *TCategoryPanelGroup) SetStyleElements(value TStyleElements) {
    CategoryPanelGroup_SetStyleElements(c.instance, value)
}

// Width
func (c *TCategoryPanelGroup) Width() int32 {
    return CategoryPanelGroup_GetWidth(c.instance)
}

// SetWidth
func (c *TCategoryPanelGroup) SetWidth(value int32) {
    CategoryPanelGroup_SetWidth(c.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TCategoryPanelGroup) SetOnClick(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
func (c *TCategoryPanelGroup) SetOnContextPopup(fn TContextPopupEvent) {
    CategoryPanelGroup_SetOnContextPopup(c.instance, fn)
}

// SetOnDblClick
func (c *TCategoryPanelGroup) SetOnDblClick(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnDblClick(c.instance, fn)
}

// SetOnDockDrop
func (c *TCategoryPanelGroup) SetOnDockDrop(fn TDockDropEvent) {
    CategoryPanelGroup_SetOnDockDrop(c.instance, fn)
}

// SetOnDragDrop
func (c *TCategoryPanelGroup) SetOnDragDrop(fn TDragDropEvent) {
    CategoryPanelGroup_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
func (c *TCategoryPanelGroup) SetOnDragOver(fn TDragOverEvent) {
    CategoryPanelGroup_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
func (c *TCategoryPanelGroup) SetOnEndDock(fn TEndDragEvent) {
    CategoryPanelGroup_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
func (c *TCategoryPanelGroup) SetOnEndDrag(fn TEndDragEvent) {
    CategoryPanelGroup_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
func (c *TCategoryPanelGroup) SetOnEnter(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnEnter(c.instance, fn)
}

// SetOnExit
func (c *TCategoryPanelGroup) SetOnExit(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnExit(c.instance, fn)
}

// SetOnGetSiteInfo
func (c *TCategoryPanelGroup) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    CategoryPanelGroup_SetOnGetSiteInfo(c.instance, fn)
}

// SetOnMouseDown
func (c *TCategoryPanelGroup) SetOnMouseDown(fn TMouseEvent) {
    CategoryPanelGroup_SetOnMouseDown(c.instance, fn)
}

// SetOnMouseEnter
func (c *TCategoryPanelGroup) SetOnMouseEnter(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
func (c *TCategoryPanelGroup) SetOnMouseLeave(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnMouseLeave(c.instance, fn)
}

// SetOnMouseMove
func (c *TCategoryPanelGroup) SetOnMouseMove(fn TMouseMoveEvent) {
    CategoryPanelGroup_SetOnMouseMove(c.instance, fn)
}

// SetOnMouseUp
func (c *TCategoryPanelGroup) SetOnMouseUp(fn TMouseEvent) {
    CategoryPanelGroup_SetOnMouseUp(c.instance, fn)
}

// SetOnMouseWheel
func (c *TCategoryPanelGroup) SetOnMouseWheel(fn TMouseWheelEvent) {
    CategoryPanelGroup_SetOnMouseWheel(c.instance, fn)
}

// SetOnMouseWheelDown
func (c *TCategoryPanelGroup) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    CategoryPanelGroup_SetOnMouseWheelDown(c.instance, fn)
}

// SetOnMouseWheelUp
func (c *TCategoryPanelGroup) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    CategoryPanelGroup_SetOnMouseWheelUp(c.instance, fn)
}

// SetOnResize
func (c *TCategoryPanelGroup) SetOnResize(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnResize(c.instance, fn)
}

// SetOnStartDock
func (c *TCategoryPanelGroup) SetOnStartDock(fn TStartDockEvent) {
    CategoryPanelGroup_SetOnStartDock(c.instance, fn)
}

// SetOnUnDock
func (c *TCategoryPanelGroup) SetOnUnDock(fn TUnDockEvent) {
    CategoryPanelGroup_SetOnUnDock(c.instance, fn)
}

// Panels
func (c *TCategoryPanelGroup) Panels() *TList {
    return ListFromInst(CategoryPanelGroup_GetPanels(c.instance))
}

// DockClientCount
func (c *TCategoryPanelGroup) DockClientCount() int32 {
    return CategoryPanelGroup_GetDockClientCount(c.instance)
}

// AlignDisabled
func (c *TCategoryPanelGroup) AlignDisabled() bool {
    return CategoryPanelGroup_GetAlignDisabled(c.instance)
}

// MouseInClient
func (c *TCategoryPanelGroup) MouseInClient() bool {
    return CategoryPanelGroup_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
func (c *TCategoryPanelGroup) VisibleDockClientCount() int32 {
    return CategoryPanelGroup_GetVisibleDockClientCount(c.instance)
}

// Brush
func (c *TCategoryPanelGroup) Brush() *TBrush {
    return BrushFromInst(CategoryPanelGroup_GetBrush(c.instance))
}

// ControlCount
func (c *TCategoryPanelGroup) ControlCount() int32 {
    return CategoryPanelGroup_GetControlCount(c.instance)
}

// Handle
func (c *TCategoryPanelGroup) Handle() HWND {
    return CategoryPanelGroup_GetHandle(c.instance)
}

// ParentWindow
func (c *TCategoryPanelGroup) ParentWindow() HWND {
    return CategoryPanelGroup_GetParentWindow(c.instance)
}

// SetParentWindow
func (c *TCategoryPanelGroup) SetParentWindow(value HWND) {
    CategoryPanelGroup_SetParentWindow(c.instance, value)
}

// UseDockManager
func (c *TCategoryPanelGroup) UseDockManager() bool {
    return CategoryPanelGroup_GetUseDockManager(c.instance)
}

// SetUseDockManager
func (c *TCategoryPanelGroup) SetUseDockManager(value bool) {
    CategoryPanelGroup_SetUseDockManager(c.instance, value)
}

// Action
func (c *TCategoryPanelGroup) Action() *TAction {
    return ActionFromInst(CategoryPanelGroup_GetAction(c.instance))
}

// SetAction
func (c *TCategoryPanelGroup) SetAction(value IComponent) {
    CategoryPanelGroup_SetAction(c.instance, CheckPtr(value))
}

// BoundsRect
func (c *TCategoryPanelGroup) BoundsRect() TRect {
    return CategoryPanelGroup_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TCategoryPanelGroup) SetBoundsRect(value TRect) {
    CategoryPanelGroup_SetBoundsRect(c.instance, value)
}

// ClientHeight
func (c *TCategoryPanelGroup) ClientHeight() int32 {
    return CategoryPanelGroup_GetClientHeight(c.instance)
}

// SetClientHeight
func (c *TCategoryPanelGroup) SetClientHeight(value int32) {
    CategoryPanelGroup_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TCategoryPanelGroup) ClientOrigin() TPoint {
    return CategoryPanelGroup_GetClientOrigin(c.instance)
}

// ClientRect
func (c *TCategoryPanelGroup) ClientRect() TRect {
    return CategoryPanelGroup_GetClientRect(c.instance)
}

// ClientWidth
func (c *TCategoryPanelGroup) ClientWidth() int32 {
    return CategoryPanelGroup_GetClientWidth(c.instance)
}

// SetClientWidth
func (c *TCategoryPanelGroup) SetClientWidth(value int32) {
    CategoryPanelGroup_SetClientWidth(c.instance, value)
}

// ControlState
func (c *TCategoryPanelGroup) ControlState() TControlState {
    return CategoryPanelGroup_GetControlState(c.instance)
}

// SetControlState
func (c *TCategoryPanelGroup) SetControlState(value TControlState) {
    CategoryPanelGroup_SetControlState(c.instance, value)
}

// ControlStyle
func (c *TCategoryPanelGroup) ControlStyle() TControlStyle {
    return CategoryPanelGroup_GetControlStyle(c.instance)
}

// SetControlStyle
func (c *TCategoryPanelGroup) SetControlStyle(value TControlStyle) {
    CategoryPanelGroup_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TCategoryPanelGroup) ExplicitLeft() int32 {
    return CategoryPanelGroup_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TCategoryPanelGroup) ExplicitTop() int32 {
    return CategoryPanelGroup_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TCategoryPanelGroup) ExplicitWidth() int32 {
    return CategoryPanelGroup_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TCategoryPanelGroup) ExplicitHeight() int32 {
    return CategoryPanelGroup_GetExplicitHeight(c.instance)
}

// Floating
func (c *TCategoryPanelGroup) Floating() bool {
    return CategoryPanelGroup_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TCategoryPanelGroup) Parent() *TWinControl {
    return WinControlFromInst(CategoryPanelGroup_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TCategoryPanelGroup) SetParent(value IWinControl) {
    CategoryPanelGroup_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
func (c *TCategoryPanelGroup) AlignWithMargins() bool {
    return CategoryPanelGroup_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
func (c *TCategoryPanelGroup) SetAlignWithMargins(value bool) {
    CategoryPanelGroup_SetAlignWithMargins(c.instance, value)
}

// Left
func (c *TCategoryPanelGroup) Left() int32 {
    return CategoryPanelGroup_GetLeft(c.instance)
}

// SetLeft
func (c *TCategoryPanelGroup) SetLeft(value int32) {
    CategoryPanelGroup_SetLeft(c.instance, value)
}

// Top
func (c *TCategoryPanelGroup) Top() int32 {
    return CategoryPanelGroup_GetTop(c.instance)
}

// SetTop
func (c *TCategoryPanelGroup) SetTop(value int32) {
    CategoryPanelGroup_SetTop(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TCategoryPanelGroup) Cursor() TCursor {
    return CategoryPanelGroup_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TCategoryPanelGroup) SetCursor(value TCursor) {
    CategoryPanelGroup_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (c *TCategoryPanelGroup) Hint() string {
    return CategoryPanelGroup_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (c *TCategoryPanelGroup) SetHint(value string) {
    CategoryPanelGroup_SetHint(c.instance, value)
}

// Margins
func (c *TCategoryPanelGroup) Margins() *TMargins {
    return MarginsFromInst(CategoryPanelGroup_GetMargins(c.instance))
}

// SetMargins
func (c *TCategoryPanelGroup) SetMargins(value *TMargins) {
    CategoryPanelGroup_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
func (c *TCategoryPanelGroup) CustomHint() *TCustomHint {
    return CustomHintFromInst(CategoryPanelGroup_GetCustomHint(c.instance))
}

// SetCustomHint
func (c *TCategoryPanelGroup) SetCustomHint(value IComponent) {
    CategoryPanelGroup_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TCategoryPanelGroup) ComponentCount() int32 {
    return CategoryPanelGroup_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TCategoryPanelGroup) ComponentIndex() int32 {
    return CategoryPanelGroup_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TCategoryPanelGroup) SetComponentIndex(value int32) {
    CategoryPanelGroup_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TCategoryPanelGroup) Owner() *TComponent {
    return ComponentFromInst(CategoryPanelGroup_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TCategoryPanelGroup) Name() string {
    return CategoryPanelGroup_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TCategoryPanelGroup) SetName(value string) {
    CategoryPanelGroup_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TCategoryPanelGroup) Tag() int {
    return CategoryPanelGroup_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TCategoryPanelGroup) SetTag(value int) {
    CategoryPanelGroup_SetTag(c.instance, value)
}

// DockClients
func (c *TCategoryPanelGroup) DockClients(Index int32) *TControl {
    return ControlFromInst(CategoryPanelGroup_GetDockClients(c.instance, Index))
}

// Controls
func (c *TCategoryPanelGroup) Controls(Index int32) *TControl {
    return ControlFromInst(CategoryPanelGroup_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TCategoryPanelGroup) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CategoryPanelGroup_GetComponents(c.instance, AIndex))
}

