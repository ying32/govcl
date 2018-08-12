
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

type TTreeView struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTreeView
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTreeView(owner IComponent) *TTreeView {
    t := new(TTreeView)
    t.instance = TreeView_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TreeViewFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TreeViewFromInst(inst uintptr) *TTreeView {
    t := new(TTreeView)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TreeViewFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TreeViewFromObj(obj IObject) *TTreeView {
    t := new(TTreeView)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TreeViewFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TreeViewFromUnsafePointer(ptr unsafe.Pointer) *TTreeView {
    t := new(TTreeView)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTreeView) Free() {
    if t.instance != 0 {
        TreeView_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTreeView) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTreeView) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTreeView) IsValid() bool {
    return t.instance != 0
}

// TTreeViewClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTreeViewClass() TClass {
    return TreeView_StaticClassType()
}

// AlphaSort
func (t *TTreeView) AlphaSort(ARecurse bool) bool {
    return TreeView_AlphaSort(t.instance, ARecurse)
}

// FullCollapse
func (t *TTreeView) FullCollapse() {
    TreeView_FullCollapse(t.instance)
}

// FullExpand
func (t *TTreeView) FullExpand() {
    TreeView_FullExpand(t.instance)
}

// GetNodeAt
func (t *TTreeView) GetNodeAt(X int32, Y int32) *TTreeNode {
    return TreeNodeFromInst(TreeView_GetNodeAt(t.instance, X , Y))
}

// IsEditing
func (t *TTreeView) IsEditing() bool {
    return TreeView_IsEditing(t.instance)
}

// LoadFromFile
func (t *TTreeView) LoadFromFile(FileName string) {
    TreeView_LoadFromFile(t.instance, FileName)
}

// LoadFromStream
func (t *TTreeView) LoadFromStream(Stream IObject) {
    TreeView_LoadFromStream(t.instance, CheckPtr(Stream))
}

// SaveToFile
func (t *TTreeView) SaveToFile(FileName string) {
    TreeView_SaveToFile(t.instance, FileName)
}

// SaveToStream
func (t *TTreeView) SaveToStream(Stream IObject) {
    TreeView_SaveToStream(t.instance, CheckPtr(Stream))
}

// Deselect
func (t *TTreeView) Deselect(Node *TTreeNode) {
    TreeView_Deselect(t.instance, CheckPtr(Node))
}

// Subselect
func (t *TTreeView) Subselect(Node *TTreeNode, Validate bool) {
    TreeView_Subselect(t.instance, CheckPtr(Node), Validate)
}

// ClearSelection
func (t *TTreeView) ClearSelection(KeepPrimary bool) {
    TreeView_ClearSelection(t.instance, KeepPrimary)
}

// FindNextToSelect
func (t *TTreeView) FindNextToSelect() *TTreeNode {
    return TreeNodeFromInst(TreeView_FindNextToSelect(t.instance))
}

// CustomSort
func (t *TTreeView) CustomSort(SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
    return TreeView_CustomSort(t.instance, SortProc , Data , ARecurse)
}

// CanFocus
func (t *TTreeView) CanFocus() bool {
    return TreeView_CanFocus(t.instance)
}

// ContainsControl
func (t *TTreeView) ContainsControl(Control IControl) bool {
    return TreeView_ContainsControl(t.instance, CheckPtr(Control))
}

// ControlAtPos
func (t *TTreeView) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(TreeView_ControlAtPos(t.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (t *TTreeView) DisableAlign() {
    TreeView_DisableAlign(t.instance)
}

// EnableAlign
func (t *TTreeView) EnableAlign() {
    TreeView_EnableAlign(t.instance)
}

// FindChildControl
func (t *TTreeView) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(TreeView_FindChildControl(t.instance, ControlName))
}

// FlipChildren
func (t *TTreeView) FlipChildren(AllLevels bool) {
    TreeView_FlipChildren(t.instance, AllLevels)
}

// Focused
func (t *TTreeView) Focused() bool {
    return TreeView_Focused(t.instance)
}

// HandleAllocated
func (t *TTreeView) HandleAllocated() bool {
    return TreeView_HandleAllocated(t.instance)
}

// InsertControl
func (t *TTreeView) InsertControl(AControl IControl) {
    TreeView_InsertControl(t.instance, CheckPtr(AControl))
}

// Invalidate
func (t *TTreeView) Invalidate() {
    TreeView_Invalidate(t.instance)
}

// PaintTo
func (t *TTreeView) PaintTo(DC HDC, X int32, Y int32) {
    TreeView_PaintTo(t.instance, DC , X , Y)
}

// RemoveControl
func (t *TTreeView) RemoveControl(AControl IControl) {
    TreeView_RemoveControl(t.instance, CheckPtr(AControl))
}

// Realign
func (t *TTreeView) Realign() {
    TreeView_Realign(t.instance)
}

// Repaint
func (t *TTreeView) Repaint() {
    TreeView_Repaint(t.instance)
}

// ScaleBy
func (t *TTreeView) ScaleBy(M int32, D int32) {
    TreeView_ScaleBy(t.instance, M , D)
}

// ScrollBy
func (t *TTreeView) ScrollBy(DeltaX int32, DeltaY int32) {
    TreeView_ScrollBy(t.instance, DeltaX , DeltaY)
}

// SetBounds
func (t *TTreeView) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    TreeView_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (t *TTreeView) SetFocus() {
    TreeView_SetFocus(t.instance)
}

// Update
func (t *TTreeView) Update() {
    TreeView_Update(t.instance)
}

// UpdateControlState
func (t *TTreeView) UpdateControlState() {
    TreeView_UpdateControlState(t.instance)
}

// BringToFront
func (t *TTreeView) BringToFront() {
    TreeView_BringToFront(t.instance)
}

// ClientToScreen
func (t *TTreeView) ClientToScreen(Point TPoint) TPoint {
    return TreeView_ClientToScreen(t.instance, Point)
}

// ClientToParent
func (t *TTreeView) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return TreeView_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// Dragging
func (t *TTreeView) Dragging() bool {
    return TreeView_Dragging(t.instance)
}

// HasParent
func (t *TTreeView) HasParent() bool {
    return TreeView_HasParent(t.instance)
}

// Hide
func (t *TTreeView) Hide() {
    TreeView_Hide(t.instance)
}

// Perform
func (t *TTreeView) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return TreeView_Perform(t.instance, Msg , WParam , LParam)
}

// Refresh
func (t *TTreeView) Refresh() {
    TreeView_Refresh(t.instance)
}

// ScreenToClient
func (t *TTreeView) ScreenToClient(Point TPoint) TPoint {
    return TreeView_ScreenToClient(t.instance, Point)
}

// ParentToClient
func (t *TTreeView) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return TreeView_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (t *TTreeView) SendToBack() {
    TreeView_SendToBack(t.instance)
}

// Show
func (t *TTreeView) Show() {
    TreeView_Show(t.instance)
}

// GetTextBuf
func (t *TTreeView) GetTextBuf(Buffer string, BufSize int32) int32 {
    return TreeView_GetTextBuf(t.instance, Buffer , BufSize)
}

// GetTextLen
func (t *TTreeView) GetTextLen() int32 {
    return TreeView_GetTextLen(t.instance)
}

// SetTextBuf
func (t *TTreeView) SetTextBuf(Buffer string) {
    TreeView_SetTextBuf(t.instance, Buffer)
}

// FindComponent
func (t *TTreeView) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TreeView_FindComponent(t.instance, AName))
}

// GetNamePath
func (t *TTreeView) GetNamePath() string {
    return TreeView_GetNamePath(t.instance)
}

// Assign
func (t *TTreeView) Assign(Source IObject) {
    TreeView_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTreeView) DisposeOf() {
    TreeView_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTreeView) ClassType() TClass {
    return TreeView_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTreeView) ClassName() string {
    return TreeView_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTreeView) InstanceSize() int32 {
    return TreeView_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTreeView) InheritsFrom(AClass TClass) bool {
    return TreeView_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTreeView) Equals(Obj IObject) bool {
    return TreeView_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTreeView) GetHashCode() int32 {
    return TreeView_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTreeView) ToString() string {
    return TreeView_ToString(t.instance)
}

// Align
func (t *TTreeView) Align() TAlign {
    return TreeView_GetAlign(t.instance)
}

// SetAlign
func (t *TTreeView) SetAlign(value TAlign) {
    TreeView_SetAlign(t.instance, value)
}

// Anchors
func (t *TTreeView) Anchors() TAnchors {
    return TreeView_GetAnchors(t.instance)
}

// SetAnchors
func (t *TTreeView) SetAnchors(value TAnchors) {
    TreeView_SetAnchors(t.instance, value)
}

// AutoExpand
func (t *TTreeView) AutoExpand() bool {
    return TreeView_GetAutoExpand(t.instance)
}

// SetAutoExpand
func (t *TTreeView) SetAutoExpand(value bool) {
    TreeView_SetAutoExpand(t.instance, value)
}

// BevelEdges
func (t *TTreeView) BevelEdges() TBevelEdges {
    return TreeView_GetBevelEdges(t.instance)
}

// SetBevelEdges
func (t *TTreeView) SetBevelEdges(value TBevelEdges) {
    TreeView_SetBevelEdges(t.instance, value)
}

// BevelInner
func (t *TTreeView) BevelInner() TBevelCut {
    return TreeView_GetBevelInner(t.instance)
}

// SetBevelInner
func (t *TTreeView) SetBevelInner(value TBevelCut) {
    TreeView_SetBevelInner(t.instance, value)
}

// BevelOuter
func (t *TTreeView) BevelOuter() TBevelCut {
    return TreeView_GetBevelOuter(t.instance)
}

// SetBevelOuter
func (t *TTreeView) SetBevelOuter(value TBevelCut) {
    TreeView_SetBevelOuter(t.instance, value)
}

// BevelKind
func (t *TTreeView) BevelKind() TBevelKind {
    return TreeView_GetBevelKind(t.instance)
}

// SetBevelKind
func (t *TTreeView) SetBevelKind(value TBevelKind) {
    TreeView_SetBevelKind(t.instance, value)
}

// BiDiMode
func (t *TTreeView) BiDiMode() TBiDiMode {
    return TreeView_GetBiDiMode(t.instance)
}

// SetBiDiMode
func (t *TTreeView) SetBiDiMode(value TBiDiMode) {
    TreeView_SetBiDiMode(t.instance, value)
}

// BorderStyle
func (t *TTreeView) BorderStyle() TBorderStyle {
    return TreeView_GetBorderStyle(t.instance)
}

// SetBorderStyle
func (t *TTreeView) SetBorderStyle(value TBorderStyle) {
    TreeView_SetBorderStyle(t.instance, value)
}

// BorderWidth
func (t *TTreeView) BorderWidth() int32 {
    return TreeView_GetBorderWidth(t.instance)
}

// SetBorderWidth
func (t *TTreeView) SetBorderWidth(value int32) {
    TreeView_SetBorderWidth(t.instance, value)
}

// ChangeDelay
func (t *TTreeView) ChangeDelay() int32 {
    return TreeView_GetChangeDelay(t.instance)
}

// SetChangeDelay
func (t *TTreeView) SetChangeDelay(value int32) {
    TreeView_SetChangeDelay(t.instance, value)
}

// Color
func (t *TTreeView) Color() TColor {
    return TreeView_GetColor(t.instance)
}

// SetColor
func (t *TTreeView) SetColor(value TColor) {
    TreeView_SetColor(t.instance, value)
}

// DoubleBuffered
func (t *TTreeView) DoubleBuffered() bool {
    return TreeView_GetDoubleBuffered(t.instance)
}

// SetDoubleBuffered
func (t *TTreeView) SetDoubleBuffered(value bool) {
    TreeView_SetDoubleBuffered(t.instance, value)
}

// DragKind
func (t *TTreeView) DragKind() TDragKind {
    return TreeView_GetDragKind(t.instance)
}

// SetDragKind
func (t *TTreeView) SetDragKind(value TDragKind) {
    TreeView_SetDragKind(t.instance, value)
}

// DragCursor
func (t *TTreeView) DragCursor() TCursor {
    return TreeView_GetDragCursor(t.instance)
}

// SetDragCursor
func (t *TTreeView) SetDragCursor(value TCursor) {
    TreeView_SetDragCursor(t.instance, value)
}

// DragMode
func (t *TTreeView) DragMode() TDragMode {
    return TreeView_GetDragMode(t.instance)
}

// SetDragMode
func (t *TTreeView) SetDragMode(value TDragMode) {
    TreeView_SetDragMode(t.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TTreeView) Enabled() bool {
    return TreeView_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TTreeView) SetEnabled(value bool) {
    TreeView_SetEnabled(t.instance, value)
}

// Font
func (t *TTreeView) Font() *TFont {
    return FontFromInst(TreeView_GetFont(t.instance))
}

// SetFont
func (t *TTreeView) SetFont(value *TFont) {
    TreeView_SetFont(t.instance, CheckPtr(value))
}

// HideSelection
func (t *TTreeView) HideSelection() bool {
    return TreeView_GetHideSelection(t.instance)
}

// SetHideSelection
func (t *TTreeView) SetHideSelection(value bool) {
    TreeView_SetHideSelection(t.instance, value)
}

// HotTrack
func (t *TTreeView) HotTrack() bool {
    return TreeView_GetHotTrack(t.instance)
}

// SetHotTrack
func (t *TTreeView) SetHotTrack(value bool) {
    TreeView_SetHotTrack(t.instance, value)
}

// Images
func (t *TTreeView) Images() *TImageList {
    return ImageListFromInst(TreeView_GetImages(t.instance))
}

// SetImages
func (t *TTreeView) SetImages(value IComponent) {
    TreeView_SetImages(t.instance, CheckPtr(value))
}

// Indent
func (t *TTreeView) Indent() int32 {
    return TreeView_GetIndent(t.instance)
}

// SetIndent
func (t *TTreeView) SetIndent(value int32) {
    TreeView_SetIndent(t.instance, value)
}

// MultiSelect
func (t *TTreeView) MultiSelect() bool {
    return TreeView_GetMultiSelect(t.instance)
}

// SetMultiSelect
func (t *TTreeView) SetMultiSelect(value bool) {
    TreeView_SetMultiSelect(t.instance, value)
}

// MultiSelectStyle
func (t *TTreeView) MultiSelectStyle() TMultiSelectStyle {
    return TreeView_GetMultiSelectStyle(t.instance)
}

// SetMultiSelectStyle
func (t *TTreeView) SetMultiSelectStyle(value TMultiSelectStyle) {
    TreeView_SetMultiSelectStyle(t.instance, value)
}

// ParentColor
func (t *TTreeView) ParentColor() bool {
    return TreeView_GetParentColor(t.instance)
}

// SetParentColor
func (t *TTreeView) SetParentColor(value bool) {
    TreeView_SetParentColor(t.instance, value)
}

// ParentCtl3D
func (t *TTreeView) ParentCtl3D() bool {
    return TreeView_GetParentCtl3D(t.instance)
}

// SetParentCtl3D
func (t *TTreeView) SetParentCtl3D(value bool) {
    TreeView_SetParentCtl3D(t.instance, value)
}

// ParentDoubleBuffered
func (t *TTreeView) ParentDoubleBuffered() bool {
    return TreeView_GetParentDoubleBuffered(t.instance)
}

// SetParentDoubleBuffered
func (t *TTreeView) SetParentDoubleBuffered(value bool) {
    TreeView_SetParentDoubleBuffered(t.instance, value)
}

// ParentFont
func (t *TTreeView) ParentFont() bool {
    return TreeView_GetParentFont(t.instance)
}

// SetParentFont
func (t *TTreeView) SetParentFont(value bool) {
    TreeView_SetParentFont(t.instance, value)
}

// ParentShowHint
func (t *TTreeView) ParentShowHint() bool {
    return TreeView_GetParentShowHint(t.instance)
}

// SetParentShowHint
func (t *TTreeView) SetParentShowHint(value bool) {
    TreeView_SetParentShowHint(t.instance, value)
}

// PopupMenu
func (t *TTreeView) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TreeView_GetPopupMenu(t.instance))
}

// SetPopupMenu
func (t *TTreeView) SetPopupMenu(value IComponent) {
    TreeView_SetPopupMenu(t.instance, CheckPtr(value))
}

// ReadOnly
func (t *TTreeView) ReadOnly() bool {
    return TreeView_GetReadOnly(t.instance)
}

// SetReadOnly
func (t *TTreeView) SetReadOnly(value bool) {
    TreeView_SetReadOnly(t.instance, value)
}

// RightClickSelect
func (t *TTreeView) RightClickSelect() bool {
    return TreeView_GetRightClickSelect(t.instance)
}

// SetRightClickSelect
func (t *TTreeView) SetRightClickSelect(value bool) {
    TreeView_SetRightClickSelect(t.instance, value)
}

// RowSelect
func (t *TTreeView) RowSelect() bool {
    return TreeView_GetRowSelect(t.instance)
}

// SetRowSelect
func (t *TTreeView) SetRowSelect(value bool) {
    TreeView_SetRowSelect(t.instance, value)
}

// ShowButtons
func (t *TTreeView) ShowButtons() bool {
    return TreeView_GetShowButtons(t.instance)
}

// SetShowButtons
func (t *TTreeView) SetShowButtons(value bool) {
    TreeView_SetShowButtons(t.instance, value)
}

// ShowHint
func (t *TTreeView) ShowHint() bool {
    return TreeView_GetShowHint(t.instance)
}

// SetShowHint
func (t *TTreeView) SetShowHint(value bool) {
    TreeView_SetShowHint(t.instance, value)
}

// ShowLines
func (t *TTreeView) ShowLines() bool {
    return TreeView_GetShowLines(t.instance)
}

// SetShowLines
func (t *TTreeView) SetShowLines(value bool) {
    TreeView_SetShowLines(t.instance, value)
}

// ShowRoot
func (t *TTreeView) ShowRoot() bool {
    return TreeView_GetShowRoot(t.instance)
}

// SetShowRoot
func (t *TTreeView) SetShowRoot(value bool) {
    TreeView_SetShowRoot(t.instance, value)
}

// SortType
func (t *TTreeView) SortType() TSortType {
    return TreeView_GetSortType(t.instance)
}

// SetSortType
func (t *TTreeView) SetSortType(value TSortType) {
    TreeView_SetSortType(t.instance, value)
}

// StateImages
func (t *TTreeView) StateImages() *TImageList {
    return ImageListFromInst(TreeView_GetStateImages(t.instance))
}

// SetStateImages
func (t *TTreeView) SetStateImages(value IComponent) {
    TreeView_SetStateImages(t.instance, CheckPtr(value))
}

// TabOrder
func (t *TTreeView) TabOrder() TTabOrder {
    return TreeView_GetTabOrder(t.instance)
}

// SetTabOrder
func (t *TTreeView) SetTabOrder(value TTabOrder) {
    TreeView_SetTabOrder(t.instance, value)
}

// TabStop
func (t *TTreeView) TabStop() bool {
    return TreeView_GetTabStop(t.instance)
}

// SetTabStop
func (t *TTreeView) SetTabStop(value bool) {
    TreeView_SetTabStop(t.instance, value)
}

// ToolTips
func (t *TTreeView) ToolTips() bool {
    return TreeView_GetToolTips(t.instance)
}

// SetToolTips
func (t *TTreeView) SetToolTips(value bool) {
    TreeView_SetToolTips(t.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TTreeView) Visible() bool {
    return TreeView_GetVisible(t.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TTreeView) SetVisible(value bool) {
    TreeView_SetVisible(t.instance, value)
}

// StyleElements
func (t *TTreeView) StyleElements() TStyleElements {
    return TreeView_GetStyleElements(t.instance)
}

// SetStyleElements
func (t *TTreeView) SetStyleElements(value TStyleElements) {
    TreeView_SetStyleElements(t.instance, value)
}

// SetOnAdvancedCustomDraw
func (t *TTreeView) SetOnAdvancedCustomDraw(fn TTVAdvancedCustomDrawEvent) {
    TreeView_SetOnAdvancedCustomDraw(t.instance, fn)
}

// SetOnAdvancedCustomDrawItem
func (t *TTreeView) SetOnAdvancedCustomDrawItem(fn TTVAdvancedCustomDrawItemEvent) {
    TreeView_SetOnAdvancedCustomDrawItem(t.instance, fn)
}

// SetOnChange
func (t *TTreeView) SetOnChange(fn TTVChangedEvent) {
    TreeView_SetOnChange(t.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (t *TTreeView) SetOnClick(fn TNotifyEvent) {
    TreeView_SetOnClick(t.instance, fn)
}

// SetOnCompare
func (t *TTreeView) SetOnCompare(fn TTVCompareEvent) {
    TreeView_SetOnCompare(t.instance, fn)
}

// SetOnContextPopup
func (t *TTreeView) SetOnContextPopup(fn TContextPopupEvent) {
    TreeView_SetOnContextPopup(t.instance, fn)
}

// SetOnDblClick
func (t *TTreeView) SetOnDblClick(fn TNotifyEvent) {
    TreeView_SetOnDblClick(t.instance, fn)
}

// SetOnDragDrop
func (t *TTreeView) SetOnDragDrop(fn TDragDropEvent) {
    TreeView_SetOnDragDrop(t.instance, fn)
}

// SetOnDragOver
func (t *TTreeView) SetOnDragOver(fn TDragOverEvent) {
    TreeView_SetOnDragOver(t.instance, fn)
}

// SetOnEndDock
func (t *TTreeView) SetOnEndDock(fn TEndDragEvent) {
    TreeView_SetOnEndDock(t.instance, fn)
}

// SetOnEndDrag
func (t *TTreeView) SetOnEndDrag(fn TEndDragEvent) {
    TreeView_SetOnEndDrag(t.instance, fn)
}

// SetOnEnter
func (t *TTreeView) SetOnEnter(fn TNotifyEvent) {
    TreeView_SetOnEnter(t.instance, fn)
}

// SetOnExit
func (t *TTreeView) SetOnExit(fn TNotifyEvent) {
    TreeView_SetOnExit(t.instance, fn)
}

// SetOnGetImageIndex
func (t *TTreeView) SetOnGetImageIndex(fn TTVExpandedEvent) {
    TreeView_SetOnGetImageIndex(t.instance, fn)
}

// SetOnGetSelectedIndex
func (t *TTreeView) SetOnGetSelectedIndex(fn TTVExpandedEvent) {
    TreeView_SetOnGetSelectedIndex(t.instance, fn)
}

// SetOnKeyDown
func (t *TTreeView) SetOnKeyDown(fn TKeyEvent) {
    TreeView_SetOnKeyDown(t.instance, fn)
}

// SetOnKeyPress
func (t *TTreeView) SetOnKeyPress(fn TKeyPressEvent) {
    TreeView_SetOnKeyPress(t.instance, fn)
}

// SetOnKeyUp
func (t *TTreeView) SetOnKeyUp(fn TKeyEvent) {
    TreeView_SetOnKeyUp(t.instance, fn)
}

// SetOnMouseDown
func (t *TTreeView) SetOnMouseDown(fn TMouseEvent) {
    TreeView_SetOnMouseDown(t.instance, fn)
}

// SetOnMouseEnter
func (t *TTreeView) SetOnMouseEnter(fn TNotifyEvent) {
    TreeView_SetOnMouseEnter(t.instance, fn)
}

// SetOnMouseLeave
func (t *TTreeView) SetOnMouseLeave(fn TNotifyEvent) {
    TreeView_SetOnMouseLeave(t.instance, fn)
}

// SetOnMouseMove
func (t *TTreeView) SetOnMouseMove(fn TMouseMoveEvent) {
    TreeView_SetOnMouseMove(t.instance, fn)
}

// SetOnMouseUp
func (t *TTreeView) SetOnMouseUp(fn TMouseEvent) {
    TreeView_SetOnMouseUp(t.instance, fn)
}

// SetOnStartDock
func (t *TTreeView) SetOnStartDock(fn TStartDockEvent) {
    TreeView_SetOnStartDock(t.instance, fn)
}

// Items
func (t *TTreeView) Items() *TTreeNodes {
    return TreeNodesFromInst(TreeView_GetItems(t.instance))
}

// SetItems
func (t *TTreeView) SetItems(value *TTreeNodes) {
    TreeView_SetItems(t.instance, CheckPtr(value))
}

// Canvas
func (t *TTreeView) Canvas() *TCanvas {
    return CanvasFromInst(TreeView_GetCanvas(t.instance))
}

// DropTarget
func (t *TTreeView) DropTarget() *TTreeNode {
    return TreeNodeFromInst(TreeView_GetDropTarget(t.instance))
}

// SetDropTarget
func (t *TTreeView) SetDropTarget(value *TTreeNode) {
    TreeView_SetDropTarget(t.instance, CheckPtr(value))
}

// Selected
func (t *TTreeView) Selected() *TTreeNode {
    return TreeNodeFromInst(TreeView_GetSelected(t.instance))
}

// SetSelected
func (t *TTreeView) SetSelected(value *TTreeNode) {
    TreeView_SetSelected(t.instance, CheckPtr(value))
}

// TopItem
func (t *TTreeView) TopItem() *TTreeNode {
    return TreeNodeFromInst(TreeView_GetTopItem(t.instance))
}

// SetTopItem
func (t *TTreeView) SetTopItem(value *TTreeNode) {
    TreeView_SetTopItem(t.instance, CheckPtr(value))
}

// SelectionCount
func (t *TTreeView) SelectionCount() uint32 {
    return TreeView_GetSelectionCount(t.instance)
}

// DockClientCount
func (t *TTreeView) DockClientCount() int32 {
    return TreeView_GetDockClientCount(t.instance)
}

// DockSite
func (t *TTreeView) DockSite() bool {
    return TreeView_GetDockSite(t.instance)
}

// SetDockSite
func (t *TTreeView) SetDockSite(value bool) {
    TreeView_SetDockSite(t.instance, value)
}

// AlignDisabled
func (t *TTreeView) AlignDisabled() bool {
    return TreeView_GetAlignDisabled(t.instance)
}

// MouseInClient
func (t *TTreeView) MouseInClient() bool {
    return TreeView_GetMouseInClient(t.instance)
}

// VisibleDockClientCount
func (t *TTreeView) VisibleDockClientCount() int32 {
    return TreeView_GetVisibleDockClientCount(t.instance)
}

// Brush
func (t *TTreeView) Brush() *TBrush {
    return BrushFromInst(TreeView_GetBrush(t.instance))
}

// ControlCount
func (t *TTreeView) ControlCount() int32 {
    return TreeView_GetControlCount(t.instance)
}

// Handle
func (t *TTreeView) Handle() HWND {
    return TreeView_GetHandle(t.instance)
}

// ParentWindow
func (t *TTreeView) ParentWindow() HWND {
    return TreeView_GetParentWindow(t.instance)
}

// SetParentWindow
func (t *TTreeView) SetParentWindow(value HWND) {
    TreeView_SetParentWindow(t.instance, value)
}

// UseDockManager
func (t *TTreeView) UseDockManager() bool {
    return TreeView_GetUseDockManager(t.instance)
}

// SetUseDockManager
func (t *TTreeView) SetUseDockManager(value bool) {
    TreeView_SetUseDockManager(t.instance, value)
}

// Action
func (t *TTreeView) Action() *TAction {
    return ActionFromInst(TreeView_GetAction(t.instance))
}

// SetAction
func (t *TTreeView) SetAction(value IComponent) {
    TreeView_SetAction(t.instance, CheckPtr(value))
}

// BoundsRect
func (t *TTreeView) BoundsRect() TRect {
    return TreeView_GetBoundsRect(t.instance)
}

// SetBoundsRect
func (t *TTreeView) SetBoundsRect(value TRect) {
    TreeView_SetBoundsRect(t.instance, value)
}

// ClientHeight
func (t *TTreeView) ClientHeight() int32 {
    return TreeView_GetClientHeight(t.instance)
}

// SetClientHeight
func (t *TTreeView) SetClientHeight(value int32) {
    TreeView_SetClientHeight(t.instance, value)
}

// ClientOrigin
func (t *TTreeView) ClientOrigin() TPoint {
    return TreeView_GetClientOrigin(t.instance)
}

// ClientRect
func (t *TTreeView) ClientRect() TRect {
    return TreeView_GetClientRect(t.instance)
}

// ClientWidth
func (t *TTreeView) ClientWidth() int32 {
    return TreeView_GetClientWidth(t.instance)
}

// SetClientWidth
func (t *TTreeView) SetClientWidth(value int32) {
    TreeView_SetClientWidth(t.instance, value)
}

// ControlState
func (t *TTreeView) ControlState() TControlState {
    return TreeView_GetControlState(t.instance)
}

// SetControlState
func (t *TTreeView) SetControlState(value TControlState) {
    TreeView_SetControlState(t.instance, value)
}

// ControlStyle
func (t *TTreeView) ControlStyle() TControlStyle {
    return TreeView_GetControlStyle(t.instance)
}

// SetControlStyle
func (t *TTreeView) SetControlStyle(value TControlStyle) {
    TreeView_SetControlStyle(t.instance, value)
}

// ExplicitLeft
func (t *TTreeView) ExplicitLeft() int32 {
    return TreeView_GetExplicitLeft(t.instance)
}

// ExplicitTop
func (t *TTreeView) ExplicitTop() int32 {
    return TreeView_GetExplicitTop(t.instance)
}

// ExplicitWidth
func (t *TTreeView) ExplicitWidth() int32 {
    return TreeView_GetExplicitWidth(t.instance)
}

// ExplicitHeight
func (t *TTreeView) ExplicitHeight() int32 {
    return TreeView_GetExplicitHeight(t.instance)
}

// Floating
func (t *TTreeView) Floating() bool {
    return TreeView_GetFloating(t.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TTreeView) Parent() *TWinControl {
    return WinControlFromInst(TreeView_GetParent(t.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (t *TTreeView) SetParent(value IWinControl) {
    TreeView_SetParent(t.instance, CheckPtr(value))
}

// AlignWithMargins
func (t *TTreeView) AlignWithMargins() bool {
    return TreeView_GetAlignWithMargins(t.instance)
}

// SetAlignWithMargins
func (t *TTreeView) SetAlignWithMargins(value bool) {
    TreeView_SetAlignWithMargins(t.instance, value)
}

// Left
func (t *TTreeView) Left() int32 {
    return TreeView_GetLeft(t.instance)
}

// SetLeft
func (t *TTreeView) SetLeft(value int32) {
    TreeView_SetLeft(t.instance, value)
}

// Top
func (t *TTreeView) Top() int32 {
    return TreeView_GetTop(t.instance)
}

// SetTop
func (t *TTreeView) SetTop(value int32) {
    TreeView_SetTop(t.instance, value)
}

// Width
func (t *TTreeView) Width() int32 {
    return TreeView_GetWidth(t.instance)
}

// SetWidth
func (t *TTreeView) SetWidth(value int32) {
    TreeView_SetWidth(t.instance, value)
}

// Height
func (t *TTreeView) Height() int32 {
    return TreeView_GetHeight(t.instance)
}

// SetHeight
func (t *TTreeView) SetHeight(value int32) {
    TreeView_SetHeight(t.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (t *TTreeView) Cursor() TCursor {
    return TreeView_GetCursor(t.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (t *TTreeView) SetCursor(value TCursor) {
    TreeView_SetCursor(t.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (t *TTreeView) Hint() string {
    return TreeView_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (t *TTreeView) SetHint(value string) {
    TreeView_SetHint(t.instance, value)
}

// Margins
func (t *TTreeView) Margins() *TMargins {
    return MarginsFromInst(TreeView_GetMargins(t.instance))
}

// SetMargins
func (t *TTreeView) SetMargins(value *TMargins) {
    TreeView_SetMargins(t.instance, CheckPtr(value))
}

// CustomHint
func (t *TTreeView) CustomHint() *TCustomHint {
    return CustomHintFromInst(TreeView_GetCustomHint(t.instance))
}

// SetCustomHint
func (t *TTreeView) SetCustomHint(value IComponent) {
    TreeView_SetCustomHint(t.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTreeView) ComponentCount() int32 {
    return TreeView_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TTreeView) ComponentIndex() int32 {
    return TreeView_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TTreeView) SetComponentIndex(value int32) {
    TreeView_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTreeView) Owner() *TComponent {
    return ComponentFromInst(TreeView_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTreeView) Name() string {
    return TreeView_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTreeView) SetName(value string) {
    TreeView_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTreeView) Tag() int {
    return TreeView_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTreeView) SetTag(value int) {
    TreeView_SetTag(t.instance, value)
}

// Selections
func (t *TTreeView) Selections(Index int32) *TTreeNode {
    return TreeNodeFromInst(TreeView_GetSelections(t.instance, Index))
}

// DockClients
func (t *TTreeView) DockClients(Index int32) *TControl {
    return ControlFromInst(TreeView_GetDockClients(t.instance, Index))
}

// Controls
func (t *TTreeView) Controls(Index int32) *TControl {
    return ControlFromInst(TreeView_GetControls(t.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTreeView) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TreeView_GetComponents(t.instance, AIndex))
}

