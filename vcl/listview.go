
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

type TListView struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewListView
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListView(owner IComponent) *TListView {
    l := new(TListView)
    l.instance = ListView_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListViewFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ListViewFromInst(inst uintptr) *TListView {
    l := new(TListView)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// ListViewFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ListViewFromObj(obj IObject) *TListView {
    l := new(TListView)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListViewFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ListViewFromUnsafePointer(ptr unsafe.Pointer) *TListView {
    l := new(TListView)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TListView) Free() {
    if l.instance != 0 {
        ListView_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListView) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListView) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListView) IsValid() bool {
    return l.instance != 0
}

// TListViewClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListViewClass() TClass {
    return ListView_StaticClassType()
}

// AddItem
func (l *TListView) AddItem(Item string, AObject IObject) {
    ListView_AddItem(l.instance, Item , CheckPtr(AObject))
}

// AlphaSort
func (l *TListView) AlphaSort() bool {
    return ListView_AlphaSort(l.instance)
}

// Clear
func (l *TListView) Clear() {
    ListView_Clear(l.instance)
}

// ClearSelection
func (l *TListView) ClearSelection() {
    ListView_ClearSelection(l.instance)
}

// DeleteSelected
func (l *TListView) DeleteSelected() {
    ListView_DeleteSelected(l.instance)
}

// GetSearchString
func (l *TListView) GetSearchString() string {
    return ListView_GetSearchString(l.instance)
}

// IsEditing
func (l *TListView) IsEditing() bool {
    return ListView_IsEditing(l.instance)
}

// SelectAll
func (l *TListView) SelectAll() {
    ListView_SelectAll(l.instance)
}

// Scroll
func (l *TListView) Scroll(DX int32, DY int32) {
    ListView_Scroll(l.instance, DX , DY)
}

// CustomSort
func (l *TListView) CustomSort(SortProc PFNLVCOMPARE, lParam int) bool {
    return ListView_CustomSort(l.instance, SortProc , lParam)
}

// CanFocus
func (l *TListView) CanFocus() bool {
    return ListView_CanFocus(l.instance)
}

// FlipChildren
func (l *TListView) FlipChildren(AllLevels bool) {
    ListView_FlipChildren(l.instance, AllLevels)
}

// Focused
func (l *TListView) Focused() bool {
    return ListView_Focused(l.instance)
}

// HandleAllocated
func (l *TListView) HandleAllocated() bool {
    return ListView_HandleAllocated(l.instance)
}

// Invalidate
func (l *TListView) Invalidate() {
    ListView_Invalidate(l.instance)
}

// Realign
func (l *TListView) Realign() {
    ListView_Realign(l.instance)
}

// Repaint
func (l *TListView) Repaint() {
    ListView_Repaint(l.instance)
}

// ScaleBy
func (l *TListView) ScaleBy(M int32, D int32) {
    ListView_ScaleBy(l.instance, M , D)
}

// SetBounds
func (l *TListView) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ListView_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (l *TListView) SetFocus() {
    ListView_SetFocus(l.instance)
}

// Update
func (l *TListView) Update() {
    ListView_Update(l.instance)
}

// BringToFront
func (l *TListView) BringToFront() {
    ListView_BringToFront(l.instance)
}

// ClientToScreen
func (l *TListView) ClientToScreen(Point TPoint) TPoint {
    return ListView_ClientToScreen(l.instance, Point)
}

// ClientToParent
func (l *TListView) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ListView_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// Dragging
func (l *TListView) Dragging() bool {
    return ListView_Dragging(l.instance)
}

// HasParent
func (l *TListView) HasParent() bool {
    return ListView_HasParent(l.instance)
}

// Hide
func (l *TListView) Hide() {
    ListView_Hide(l.instance)
}

// Perform
func (l *TListView) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ListView_Perform(l.instance, Msg , WParam , LParam)
}

// Refresh
func (l *TListView) Refresh() {
    ListView_Refresh(l.instance)
}

// ScreenToClient
func (l *TListView) ScreenToClient(Point TPoint) TPoint {
    return ListView_ScreenToClient(l.instance, Point)
}

// ParentToClient
func (l *TListView) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ListView_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (l *TListView) SendToBack() {
    ListView_SendToBack(l.instance)
}

// Show
func (l *TListView) Show() {
    ListView_Show(l.instance)
}

// GetTextBuf
func (l *TListView) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ListView_GetTextBuf(l.instance, Buffer , BufSize)
}

// GetTextLen
func (l *TListView) GetTextLen() int32 {
    return ListView_GetTextLen(l.instance)
}

// FindComponent
func (l *TListView) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ListView_FindComponent(l.instance, AName))
}

// GetNamePath
func (l *TListView) GetNamePath() string {
    return ListView_GetNamePath(l.instance)
}

// Assign
func (l *TListView) Assign(Source IObject) {
    ListView_Assign(l.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TListView) DisposeOf() {
    ListView_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListView) ClassType() TClass {
    return ListView_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListView) ClassName() string {
    return ListView_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListView) InstanceSize() int32 {
    return ListView_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListView) InheritsFrom(AClass TClass) bool {
    return ListView_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListView) Equals(Obj IObject) bool {
    return ListView_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListView) GetHashCode() int32 {
    return ListView_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TListView) ToString() string {
    return ListView_ToString(l.instance)
}

// Action
func (l *TListView) Action() *TAction {
    return ActionFromInst(ListView_GetAction(l.instance))
}

// SetAction
func (l *TListView) SetAction(value IComponent) {
    ListView_SetAction(l.instance, CheckPtr(value))
}

// Align
func (l *TListView) Align() TAlign {
    return ListView_GetAlign(l.instance)
}

// SetAlign
func (l *TListView) SetAlign(value TAlign) {
    ListView_SetAlign(l.instance, value)
}

// AllocBy
func (l *TListView) AllocBy() int32 {
    return ListView_GetAllocBy(l.instance)
}

// SetAllocBy
func (l *TListView) SetAllocBy(value int32) {
    ListView_SetAllocBy(l.instance, value)
}

// Anchors
func (l *TListView) Anchors() TAnchors {
    return ListView_GetAnchors(l.instance)
}

// SetAnchors
func (l *TListView) SetAnchors(value TAnchors) {
    ListView_SetAnchors(l.instance, value)
}

// BevelEdges
func (l *TListView) BevelEdges() TBevelEdges {
    return ListView_GetBevelEdges(l.instance)
}

// SetBevelEdges
func (l *TListView) SetBevelEdges(value TBevelEdges) {
    ListView_SetBevelEdges(l.instance, value)
}

// BevelInner
func (l *TListView) BevelInner() TBevelCut {
    return ListView_GetBevelInner(l.instance)
}

// SetBevelInner
func (l *TListView) SetBevelInner(value TBevelCut) {
    ListView_SetBevelInner(l.instance, value)
}

// BevelOuter
func (l *TListView) BevelOuter() TBevelCut {
    return ListView_GetBevelOuter(l.instance)
}

// SetBevelOuter
func (l *TListView) SetBevelOuter(value TBevelCut) {
    ListView_SetBevelOuter(l.instance, value)
}

// BevelKind
func (l *TListView) BevelKind() TBevelKind {
    return ListView_GetBevelKind(l.instance)
}

// SetBevelKind
func (l *TListView) SetBevelKind(value TBevelKind) {
    ListView_SetBevelKind(l.instance, value)
}

// BiDiMode
func (l *TListView) BiDiMode() TBiDiMode {
    return ListView_GetBiDiMode(l.instance)
}

// SetBiDiMode
func (l *TListView) SetBiDiMode(value TBiDiMode) {
    ListView_SetBiDiMode(l.instance, value)
}

// BorderStyle
func (l *TListView) BorderStyle() TBorderStyle {
    return ListView_GetBorderStyle(l.instance)
}

// SetBorderStyle
func (l *TListView) SetBorderStyle(value TBorderStyle) {
    ListView_SetBorderStyle(l.instance, value)
}

// BorderWidth
func (l *TListView) BorderWidth() int32 {
    return ListView_GetBorderWidth(l.instance)
}

// SetBorderWidth
func (l *TListView) SetBorderWidth(value int32) {
    ListView_SetBorderWidth(l.instance, value)
}

// Checkboxes
func (l *TListView) Checkboxes() bool {
    return ListView_GetCheckboxes(l.instance)
}

// SetCheckboxes
func (l *TListView) SetCheckboxes(value bool) {
    ListView_SetCheckboxes(l.instance, value)
}

// Color
func (l *TListView) Color() TColor {
    return ListView_GetColor(l.instance)
}

// SetColor
func (l *TListView) SetColor(value TColor) {
    ListView_SetColor(l.instance, value)
}

// Columns
func (l *TListView) Columns() *TListColumns {
    return ListColumnsFromInst(ListView_GetColumns(l.instance))
}

// SetColumns
func (l *TListView) SetColumns(value *TListColumns) {
    ListView_SetColumns(l.instance, CheckPtr(value))
}

// ColumnClick
func (l *TListView) ColumnClick() bool {
    return ListView_GetColumnClick(l.instance)
}

// SetColumnClick
func (l *TListView) SetColumnClick(value bool) {
    ListView_SetColumnClick(l.instance, value)
}

// DoubleBuffered
func (l *TListView) DoubleBuffered() bool {
    return ListView_GetDoubleBuffered(l.instance)
}

// SetDoubleBuffered
func (l *TListView) SetDoubleBuffered(value bool) {
    ListView_SetDoubleBuffered(l.instance, value)
}

// DragCursor
func (l *TListView) DragCursor() TCursor {
    return ListView_GetDragCursor(l.instance)
}

// SetDragCursor
func (l *TListView) SetDragCursor(value TCursor) {
    ListView_SetDragCursor(l.instance, value)
}

// DragKind
func (l *TListView) DragKind() TDragKind {
    return ListView_GetDragKind(l.instance)
}

// SetDragKind
func (l *TListView) SetDragKind(value TDragKind) {
    ListView_SetDragKind(l.instance, value)
}

// DragMode
func (l *TListView) DragMode() TDragMode {
    return ListView_GetDragMode(l.instance)
}

// SetDragMode
func (l *TListView) SetDragMode(value TDragMode) {
    ListView_SetDragMode(l.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (l *TListView) Enabled() bool {
    return ListView_GetEnabled(l.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (l *TListView) SetEnabled(value bool) {
    ListView_SetEnabled(l.instance, value)
}

// Font
func (l *TListView) Font() *TFont {
    return FontFromInst(ListView_GetFont(l.instance))
}

// SetFont
func (l *TListView) SetFont(value *TFont) {
    ListView_SetFont(l.instance, CheckPtr(value))
}

// FlatScrollBars
func (l *TListView) FlatScrollBars() bool {
    return ListView_GetFlatScrollBars(l.instance)
}

// SetFlatScrollBars
func (l *TListView) SetFlatScrollBars(value bool) {
    ListView_SetFlatScrollBars(l.instance, value)
}

// FullDrag
func (l *TListView) FullDrag() bool {
    return ListView_GetFullDrag(l.instance)
}

// SetFullDrag
func (l *TListView) SetFullDrag(value bool) {
    ListView_SetFullDrag(l.instance, value)
}

// GridLines
func (l *TListView) GridLines() bool {
    return ListView_GetGridLines(l.instance)
}

// SetGridLines
func (l *TListView) SetGridLines(value bool) {
    ListView_SetGridLines(l.instance, value)
}

// Groups
func (l *TListView) Groups() *TListGroups {
    return ListGroupsFromInst(ListView_GetGroups(l.instance))
}

// SetGroups
func (l *TListView) SetGroups(value *TListGroups) {
    ListView_SetGroups(l.instance, CheckPtr(value))
}

// HideSelection
func (l *TListView) HideSelection() bool {
    return ListView_GetHideSelection(l.instance)
}

// SetHideSelection
func (l *TListView) SetHideSelection(value bool) {
    ListView_SetHideSelection(l.instance, value)
}

// HotTrack
func (l *TListView) HotTrack() bool {
    return ListView_GetHotTrack(l.instance)
}

// SetHotTrack
func (l *TListView) SetHotTrack(value bool) {
    ListView_SetHotTrack(l.instance, value)
}

// HoverTime
func (l *TListView) HoverTime() int32 {
    return ListView_GetHoverTime(l.instance)
}

// SetHoverTime
func (l *TListView) SetHoverTime(value int32) {
    ListView_SetHoverTime(l.instance, value)
}

// IconOptions
func (l *TListView) IconOptions() *TIconOptions {
    return IconOptionsFromInst(ListView_GetIconOptions(l.instance))
}

// SetIconOptions
func (l *TListView) SetIconOptions(value IObject) {
    ListView_SetIconOptions(l.instance, CheckPtr(value))
}

// Items
func (l *TListView) Items() *TListItems {
    return ListItemsFromInst(ListView_GetItems(l.instance))
}

// SetItems
func (l *TListView) SetItems(value *TListItems) {
    ListView_SetItems(l.instance, CheckPtr(value))
}

// LargeImages
func (l *TListView) LargeImages() *TImageList {
    return ImageListFromInst(ListView_GetLargeImages(l.instance))
}

// SetLargeImages
func (l *TListView) SetLargeImages(value IComponent) {
    ListView_SetLargeImages(l.instance, CheckPtr(value))
}

// MultiSelect
func (l *TListView) MultiSelect() bool {
    return ListView_GetMultiSelect(l.instance)
}

// SetMultiSelect
func (l *TListView) SetMultiSelect(value bool) {
    ListView_SetMultiSelect(l.instance, value)
}

// StyleElements
func (l *TListView) StyleElements() TStyleElements {
    return ListView_GetStyleElements(l.instance)
}

// SetStyleElements
func (l *TListView) SetStyleElements(value TStyleElements) {
    ListView_SetStyleElements(l.instance, value)
}

// GroupHeaderImages
func (l *TListView) GroupHeaderImages() *TImageList {
    return ImageListFromInst(ListView_GetGroupHeaderImages(l.instance))
}

// SetGroupHeaderImages
func (l *TListView) SetGroupHeaderImages(value IComponent) {
    ListView_SetGroupHeaderImages(l.instance, CheckPtr(value))
}

// GroupView
func (l *TListView) GroupView() bool {
    return ListView_GetGroupView(l.instance)
}

// SetGroupView
func (l *TListView) SetGroupView(value bool) {
    ListView_SetGroupView(l.instance, value)
}

// ReadOnly
func (l *TListView) ReadOnly() bool {
    return ListView_GetReadOnly(l.instance)
}

// SetReadOnly
func (l *TListView) SetReadOnly(value bool) {
    ListView_SetReadOnly(l.instance, value)
}

// RowSelect
func (l *TListView) RowSelect() bool {
    return ListView_GetRowSelect(l.instance)
}

// SetRowSelect
func (l *TListView) SetRowSelect(value bool) {
    ListView_SetRowSelect(l.instance, value)
}

// ParentColor
func (l *TListView) ParentColor() bool {
    return ListView_GetParentColor(l.instance)
}

// SetParentColor
func (l *TListView) SetParentColor(value bool) {
    ListView_SetParentColor(l.instance, value)
}

// ParentDoubleBuffered
func (l *TListView) ParentDoubleBuffered() bool {
    return ListView_GetParentDoubleBuffered(l.instance)
}

// SetParentDoubleBuffered
func (l *TListView) SetParentDoubleBuffered(value bool) {
    ListView_SetParentDoubleBuffered(l.instance, value)
}

// ParentFont
func (l *TListView) ParentFont() bool {
    return ListView_GetParentFont(l.instance)
}

// SetParentFont
func (l *TListView) SetParentFont(value bool) {
    ListView_SetParentFont(l.instance, value)
}

// ParentShowHint
func (l *TListView) ParentShowHint() bool {
    return ListView_GetParentShowHint(l.instance)
}

// SetParentShowHint
func (l *TListView) SetParentShowHint(value bool) {
    ListView_SetParentShowHint(l.instance, value)
}

// PopupMenu
func (l *TListView) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ListView_GetPopupMenu(l.instance))
}

// SetPopupMenu
func (l *TListView) SetPopupMenu(value IComponent) {
    ListView_SetPopupMenu(l.instance, CheckPtr(value))
}

// ShowColumnHeaders
func (l *TListView) ShowColumnHeaders() bool {
    return ListView_GetShowColumnHeaders(l.instance)
}

// SetShowColumnHeaders
func (l *TListView) SetShowColumnHeaders(value bool) {
    ListView_SetShowColumnHeaders(l.instance, value)
}

// ShowWorkAreas
func (l *TListView) ShowWorkAreas() bool {
    return ListView_GetShowWorkAreas(l.instance)
}

// SetShowWorkAreas
func (l *TListView) SetShowWorkAreas(value bool) {
    ListView_SetShowWorkAreas(l.instance, value)
}

// ShowHint
func (l *TListView) ShowHint() bool {
    return ListView_GetShowHint(l.instance)
}

// SetShowHint
func (l *TListView) SetShowHint(value bool) {
    ListView_SetShowHint(l.instance, value)
}

// SmallImages
func (l *TListView) SmallImages() *TImageList {
    return ImageListFromInst(ListView_GetSmallImages(l.instance))
}

// SetSmallImages
func (l *TListView) SetSmallImages(value IComponent) {
    ListView_SetSmallImages(l.instance, CheckPtr(value))
}

// SortType
func (l *TListView) SortType() TSortType {
    return ListView_GetSortType(l.instance)
}

// SetSortType
func (l *TListView) SetSortType(value TSortType) {
    ListView_SetSortType(l.instance, value)
}

// StateImages
func (l *TListView) StateImages() *TImageList {
    return ImageListFromInst(ListView_GetStateImages(l.instance))
}

// SetStateImages
func (l *TListView) SetStateImages(value IComponent) {
    ListView_SetStateImages(l.instance, CheckPtr(value))
}

// TabOrder
func (l *TListView) TabOrder() uint16 {
    return ListView_GetTabOrder(l.instance)
}

// SetTabOrder
func (l *TListView) SetTabOrder(value uint16) {
    ListView_SetTabOrder(l.instance, value)
}

// TabStop
func (l *TListView) TabStop() bool {
    return ListView_GetTabStop(l.instance)
}

// SetTabStop
func (l *TListView) SetTabStop(value bool) {
    ListView_SetTabStop(l.instance, value)
}

// ViewStyle
func (l *TListView) ViewStyle() TViewStyle {
    return ListView_GetViewStyle(l.instance)
}

// SetViewStyle
func (l *TListView) SetViewStyle(value TViewStyle) {
    ListView_SetViewStyle(l.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (l *TListView) Visible() bool {
    return ListView_GetVisible(l.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (l *TListView) SetVisible(value bool) {
    ListView_SetVisible(l.instance, value)
}

// SetOnAdvancedCustomDraw
func (l *TListView) SetOnAdvancedCustomDraw(fn TLVAdvancedCustomDrawEvent) {
    ListView_SetOnAdvancedCustomDraw(l.instance, fn)
}

// SetOnAdvancedCustomDrawItem
func (l *TListView) SetOnAdvancedCustomDrawItem(fn TLVAdvancedCustomDrawItemEvent) {
    ListView_SetOnAdvancedCustomDrawItem(l.instance, fn)
}

// SetOnAdvancedCustomDrawSubItem
func (l *TListView) SetOnAdvancedCustomDrawSubItem(fn TLVAdvancedCustomDrawSubItemEvent) {
    ListView_SetOnAdvancedCustomDrawSubItem(l.instance, fn)
}

// SetOnChange
func (l *TListView) SetOnChange(fn TLVChangeEvent) {
    ListView_SetOnChange(l.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (l *TListView) SetOnClick(fn TNotifyEvent) {
    ListView_SetOnClick(l.instance, fn)
}

// SetOnColumnClick
func (l *TListView) SetOnColumnClick(fn TLVColumnClickEvent) {
    ListView_SetOnColumnClick(l.instance, fn)
}

// SetOnColumnRightClick
func (l *TListView) SetOnColumnRightClick(fn TLVColumnRClickEvent) {
    ListView_SetOnColumnRightClick(l.instance, fn)
}

// SetOnCompare
func (l *TListView) SetOnCompare(fn TLVCompareEvent) {
    ListView_SetOnCompare(l.instance, fn)
}

// SetOnContextPopup
func (l *TListView) SetOnContextPopup(fn TContextPopupEvent) {
    ListView_SetOnContextPopup(l.instance, fn)
}

// SetOnDblClick
func (l *TListView) SetOnDblClick(fn TNotifyEvent) {
    ListView_SetOnDblClick(l.instance, fn)
}

// SetOnEndDock
func (l *TListView) SetOnEndDock(fn TEndDragEvent) {
    ListView_SetOnEndDock(l.instance, fn)
}

// SetOnEndDrag
func (l *TListView) SetOnEndDrag(fn TEndDragEvent) {
    ListView_SetOnEndDrag(l.instance, fn)
}

// SetOnEnter
func (l *TListView) SetOnEnter(fn TNotifyEvent) {
    ListView_SetOnEnter(l.instance, fn)
}

// SetOnExit
func (l *TListView) SetOnExit(fn TNotifyEvent) {
    ListView_SetOnExit(l.instance, fn)
}

// SetOnGetImageIndex
func (l *TListView) SetOnGetImageIndex(fn TLVNotifyEvent) {
    ListView_SetOnGetImageIndex(l.instance, fn)
}

// SetOnDragDrop
func (l *TListView) SetOnDragDrop(fn TDragDropEvent) {
    ListView_SetOnDragDrop(l.instance, fn)
}

// SetOnDragOver
func (l *TListView) SetOnDragOver(fn TDragOverEvent) {
    ListView_SetOnDragOver(l.instance, fn)
}

// SetOnKeyDown
func (l *TListView) SetOnKeyDown(fn TKeyEvent) {
    ListView_SetOnKeyDown(l.instance, fn)
}

// SetOnKeyPress
func (l *TListView) SetOnKeyPress(fn TKeyPressEvent) {
    ListView_SetOnKeyPress(l.instance, fn)
}

// SetOnKeyUp
func (l *TListView) SetOnKeyUp(fn TKeyEvent) {
    ListView_SetOnKeyUp(l.instance, fn)
}

// SetOnMouseDown
func (l *TListView) SetOnMouseDown(fn TMouseEvent) {
    ListView_SetOnMouseDown(l.instance, fn)
}

// SetOnMouseEnter
func (l *TListView) SetOnMouseEnter(fn TNotifyEvent) {
    ListView_SetOnMouseEnter(l.instance, fn)
}

// SetOnMouseLeave
func (l *TListView) SetOnMouseLeave(fn TNotifyEvent) {
    ListView_SetOnMouseLeave(l.instance, fn)
}

// SetOnMouseMove
func (l *TListView) SetOnMouseMove(fn TMouseMoveEvent) {
    ListView_SetOnMouseMove(l.instance, fn)
}

// SetOnMouseUp
func (l *TListView) SetOnMouseUp(fn TMouseEvent) {
    ListView_SetOnMouseUp(l.instance, fn)
}

// SetOnResize
func (l *TListView) SetOnResize(fn TNotifyEvent) {
    ListView_SetOnResize(l.instance, fn)
}

// SetOnSelectItem
func (l *TListView) SetOnSelectItem(fn TLVSelectItemEvent) {
    ListView_SetOnSelectItem(l.instance, fn)
}

// SetOnItemChecked
func (l *TListView) SetOnItemChecked(fn TLVCheckedItemEvent) {
    ListView_SetOnItemChecked(l.instance, fn)
}

// SetOnStartDock
func (l *TListView) SetOnStartDock(fn TStartDockEvent) {
    ListView_SetOnStartDock(l.instance, fn)
}

// Canvas
func (l *TListView) Canvas() *TCanvas {
    return CanvasFromInst(ListView_GetCanvas(l.instance))
}

// DropTarget
func (l *TListView) DropTarget() *TListItem {
    return ListItemFromInst(ListView_GetDropTarget(l.instance))
}

// SetDropTarget
func (l *TListView) SetDropTarget(value *TListItem) {
    ListView_SetDropTarget(l.instance, CheckPtr(value))
}

// ItemFocused
func (l *TListView) ItemFocused() *TListItem {
    return ListItemFromInst(ListView_GetItemFocused(l.instance))
}

// SetItemFocused
func (l *TListView) SetItemFocused(value *TListItem) {
    ListView_SetItemFocused(l.instance, CheckPtr(value))
}

// SelCount
func (l *TListView) SelCount() int32 {
    return ListView_GetSelCount(l.instance)
}

// Selected
func (l *TListView) Selected() *TListItem {
    return ListItemFromInst(ListView_GetSelected(l.instance))
}

// SetSelected
func (l *TListView) SetSelected(value *TListItem) {
    ListView_SetSelected(l.instance, CheckPtr(value))
}

// TopItem
func (l *TListView) TopItem() *TListItem {
    return ListItemFromInst(ListView_GetTopItem(l.instance))
}

// VisibleRowCount
func (l *TListView) VisibleRowCount() int32 {
    return ListView_GetVisibleRowCount(l.instance)
}

// ItemIndex
func (l *TListView) ItemIndex() int32 {
    return ListView_GetItemIndex(l.instance)
}

// SetItemIndex
func (l *TListView) SetItemIndex(value int32) {
    ListView_SetItemIndex(l.instance, value)
}

// DockSite
func (l *TListView) DockSite() bool {
    return ListView_GetDockSite(l.instance)
}

// SetDockSite
func (l *TListView) SetDockSite(value bool) {
    ListView_SetDockSite(l.instance, value)
}

// Brush
func (l *TListView) Brush() *TBrush {
    return BrushFromInst(ListView_GetBrush(l.instance))
}

// ControlCount
func (l *TListView) ControlCount() int32 {
    return ListView_GetControlCount(l.instance)
}

// Handle
func (l *TListView) Handle() HWND {
    return ListView_GetHandle(l.instance)
}

// ParentWindow
func (l *TListView) ParentWindow() HWND {
    return ListView_GetParentWindow(l.instance)
}

// SetParentWindow
func (l *TListView) SetParentWindow(value HWND) {
    ListView_SetParentWindow(l.instance, value)
}

// UseDockManager
func (l *TListView) UseDockManager() bool {
    return ListView_GetUseDockManager(l.instance)
}

// SetUseDockManager
func (l *TListView) SetUseDockManager(value bool) {
    ListView_SetUseDockManager(l.instance, value)
}

// BoundsRect
func (l *TListView) BoundsRect() TRect {
    return ListView_GetBoundsRect(l.instance)
}

// SetBoundsRect
func (l *TListView) SetBoundsRect(value TRect) {
    ListView_SetBoundsRect(l.instance, value)
}

// ClientHeight
func (l *TListView) ClientHeight() int32 {
    return ListView_GetClientHeight(l.instance)
}

// SetClientHeight
func (l *TListView) SetClientHeight(value int32) {
    ListView_SetClientHeight(l.instance, value)
}

// ClientRect
func (l *TListView) ClientRect() TRect {
    return ListView_GetClientRect(l.instance)
}

// ClientWidth
func (l *TListView) ClientWidth() int32 {
    return ListView_GetClientWidth(l.instance)
}

// SetClientWidth
func (l *TListView) SetClientWidth(value int32) {
    ListView_SetClientWidth(l.instance, value)
}

// ExplicitLeft
func (l *TListView) ExplicitLeft() int32 {
    return ListView_GetExplicitLeft(l.instance)
}

// ExplicitTop
func (l *TListView) ExplicitTop() int32 {
    return ListView_GetExplicitTop(l.instance)
}

// ExplicitWidth
func (l *TListView) ExplicitWidth() int32 {
    return ListView_GetExplicitWidth(l.instance)
}

// ExplicitHeight
func (l *TListView) ExplicitHeight() int32 {
    return ListView_GetExplicitHeight(l.instance)
}

// Floating
func (l *TListView) Floating() bool {
    return ListView_GetFloating(l.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (l *TListView) Parent() *TWinControl {
    return WinControlFromInst(ListView_GetParent(l.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (l *TListView) SetParent(value IWinControl) {
    ListView_SetParent(l.instance, CheckPtr(value))
}

// AlignWithMargins
func (l *TListView) AlignWithMargins() bool {
    return ListView_GetAlignWithMargins(l.instance)
}

// SetAlignWithMargins
func (l *TListView) SetAlignWithMargins(value bool) {
    ListView_SetAlignWithMargins(l.instance, value)
}

// Left
func (l *TListView) Left() int32 {
    return ListView_GetLeft(l.instance)
}

// SetLeft
func (l *TListView) SetLeft(value int32) {
    ListView_SetLeft(l.instance, value)
}

// Top
func (l *TListView) Top() int32 {
    return ListView_GetTop(l.instance)
}

// SetTop
func (l *TListView) SetTop(value int32) {
    ListView_SetTop(l.instance, value)
}

// Width
func (l *TListView) Width() int32 {
    return ListView_GetWidth(l.instance)
}

// SetWidth
func (l *TListView) SetWidth(value int32) {
    ListView_SetWidth(l.instance, value)
}

// Height
func (l *TListView) Height() int32 {
    return ListView_GetHeight(l.instance)
}

// SetHeight
func (l *TListView) SetHeight(value int32) {
    ListView_SetHeight(l.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (l *TListView) Cursor() TCursor {
    return ListView_GetCursor(l.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (l *TListView) SetCursor(value TCursor) {
    ListView_SetCursor(l.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (l *TListView) Hint() string {
    return ListView_GetHint(l.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (l *TListView) SetHint(value string) {
    ListView_SetHint(l.instance, value)
}

// Margins
func (l *TListView) Margins() *TMargins {
    return MarginsFromInst(ListView_GetMargins(l.instance))
}

// SetMargins
func (l *TListView) SetMargins(value *TMargins) {
    ListView_SetMargins(l.instance, CheckPtr(value))
}

// CustomHint
func (l *TListView) CustomHint() *TCustomHint {
    return CustomHintFromInst(ListView_GetCustomHint(l.instance))
}

// SetCustomHint
func (l *TListView) SetCustomHint(value IComponent) {
    ListView_SetCustomHint(l.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (l *TListView) ComponentCount() int32 {
    return ListView_GetComponentCount(l.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (l *TListView) ComponentIndex() int32 {
    return ListView_GetComponentIndex(l.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (l *TListView) SetComponentIndex(value int32) {
    ListView_SetComponentIndex(l.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TListView) Owner() *TComponent {
    return ComponentFromInst(ListView_GetOwner(l.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (l *TListView) Name() string {
    return ListView_GetName(l.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (l *TListView) SetName(value string) {
    ListView_SetName(l.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TListView) Tag() int {
    return ListView_GetTag(l.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TListView) SetTag(value int) {
    ListView_SetTag(l.instance, value)
}

// Column
func (l *TListView) Column(Index int32) *TListColumn {
    return ListColumnFromInst(ListView_GetColumn(l.instance, Index))
}

// Controls
func (l *TListView) Controls(Index int32) *TControl {
    return ControlFromInst(ListView_GetControls(l.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TListView) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ListView_GetComponents(l.instance, AIndex))
}

