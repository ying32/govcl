
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TListView struct {
    IControl
    instance uintptr
}

func NewListView(owner IComponent) *TListView {
    l := new(TListView)
    l.instance = ListView_Create(CheckPtr(owner))
    return l
}

func ListViewFromInst(inst uintptr) *TListView {
    l := new(TListView)
    l.instance = inst
    return l
}

func ListViewFromObj(obj IObject) *TListView {
    l := new(TListView)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListView) Free() {
    if l.instance != 0 {
        ListView_Free(l.instance)
        l.instance = 0
    }
}

func (l *TListView) Instance() uintptr {
    return l.instance
}

func (l *TListView) IsValid() bool {
    return l.instance != 0
}

func (l *TListView) AddItem(Item string, AObject IObject) {
    ListView_AddItem(l.instance, Item , CheckPtr(AObject))
}

func (l *TListView) AlphaSort() bool {
    return ListView_AlphaSort(l.instance)
}

func (l *TListView) Clear() {
    ListView_Clear(l.instance)
}

func (l *TListView) ClearSelection() {
    ListView_ClearSelection(l.instance)
}

func (l *TListView) DeleteSelected() {
    ListView_DeleteSelected(l.instance)
}

func (l *TListView) GetSearchString() string {
    return ListView_GetSearchString(l.instance)
}

func (l *TListView) IsEditing() bool {
    return ListView_IsEditing(l.instance)
}

func (l *TListView) SelectAll() {
    ListView_SelectAll(l.instance)
}

func (l *TListView) Scroll(DX int32, DY int32) {
    ListView_Scroll(l.instance, DX , DY)
}

func (l *TListView) CustomSort(SortProc PFNLVCOMPARE, lParam int) bool {
    return ListView_CustomSort(l.instance, SortProc , lParam)
}

func (l *TListView) CanFocus() bool {
    return ListView_CanFocus(l.instance)
}

func (l *TListView) FlipChildren(AllLevels bool) {
    ListView_FlipChildren(l.instance, AllLevels)
}

func (l *TListView) Focused() bool {
    return ListView_Focused(l.instance)
}

func (l *TListView) HandleAllocated() bool {
    return ListView_HandleAllocated(l.instance)
}

func (l *TListView) Invalidate() {
    ListView_Invalidate(l.instance)
}

func (l *TListView) Realign() {
    ListView_Realign(l.instance)
}

func (l *TListView) Repaint() {
    ListView_Repaint(l.instance)
}

func (l *TListView) ScaleBy(M int32, D int32) {
    ListView_ScaleBy(l.instance, M , D)
}

func (l *TListView) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ListView_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

func (l *TListView) SetFocus() {
    ListView_SetFocus(l.instance)
}

func (l *TListView) Update() {
    ListView_Update(l.instance)
}

func (l *TListView) BringToFront() {
    ListView_BringToFront(l.instance)
}

func (l *TListView) HasParent() bool {
    return ListView_HasParent(l.instance)
}

func (l *TListView) Hide() {
    ListView_Hide(l.instance)
}

func (l *TListView) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ListView_Perform(l.instance, Msg , WParam , LParam)
}

func (l *TListView) Refresh() {
    ListView_Refresh(l.instance)
}

func (l *TListView) SendToBack() {
    ListView_SendToBack(l.instance)
}

func (l *TListView) Show() {
    ListView_Show(l.instance)
}

func (l *TListView) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ListView_GetTextBuf(l.instance, Buffer , BufSize)
}

func (l *TListView) GetTextLen() int32 {
    return ListView_GetTextLen(l.instance)
}

func (l *TListView) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ListView_FindComponent(l.instance, AName))
}

func (l *TListView) GetNamePath() string {
    return ListView_GetNamePath(l.instance)
}

func (l *TListView) Assign(Source IObject) {
    ListView_Assign(l.instance, CheckPtr(Source))
}

func (l *TListView) ClassName() string {
    return ListView_ClassName(l.instance)
}

func (l *TListView) Equals(Obj IObject) bool {
    return ListView_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListView) GetHashCode() int32 {
    return ListView_GetHashCode(l.instance)
}

func (l *TListView) ToString() string {
    return ListView_ToString(l.instance)
}

func (l *TListView) Action() *TAction {
    return ActionFromInst(ListView_GetAction(l.instance))
}

func (l *TListView) SetAction(value IComponent) {
    ListView_SetAction(l.instance, CheckPtr(value))
}

func (l *TListView) Align() TAlign {
    return ListView_GetAlign(l.instance)
}

func (l *TListView) SetAlign(value TAlign) {
    ListView_SetAlign(l.instance, value)
}

func (l *TListView) AllocBy() int32 {
    return ListView_GetAllocBy(l.instance)
}

func (l *TListView) SetAllocBy(value int32) {
    ListView_SetAllocBy(l.instance, value)
}

func (l *TListView) Anchors() TAnchors {
    return ListView_GetAnchors(l.instance)
}

func (l *TListView) SetAnchors(value TAnchors) {
    ListView_SetAnchors(l.instance, value)
}

func (l *TListView) BevelEdges() TBevelEdges {
    return ListView_GetBevelEdges(l.instance)
}

func (l *TListView) SetBevelEdges(value TBevelEdges) {
    ListView_SetBevelEdges(l.instance, value)
}

func (l *TListView) BevelInner() TBevelCut {
    return ListView_GetBevelInner(l.instance)
}

func (l *TListView) SetBevelInner(value TBevelCut) {
    ListView_SetBevelInner(l.instance, value)
}

func (l *TListView) BevelOuter() TBevelCut {
    return ListView_GetBevelOuter(l.instance)
}

func (l *TListView) SetBevelOuter(value TBevelCut) {
    ListView_SetBevelOuter(l.instance, value)
}

func (l *TListView) BevelKind() TBevelKind {
    return ListView_GetBevelKind(l.instance)
}

func (l *TListView) SetBevelKind(value TBevelKind) {
    ListView_SetBevelKind(l.instance, value)
}

func (l *TListView) BiDiMode() TBiDiMode {
    return ListView_GetBiDiMode(l.instance)
}

func (l *TListView) SetBiDiMode(value TBiDiMode) {
    ListView_SetBiDiMode(l.instance, value)
}

func (l *TListView) BorderStyle() TBorderStyle {
    return ListView_GetBorderStyle(l.instance)
}

func (l *TListView) SetBorderStyle(value TBorderStyle) {
    ListView_SetBorderStyle(l.instance, value)
}

func (l *TListView) BorderWidth() int32 {
    return ListView_GetBorderWidth(l.instance)
}

func (l *TListView) SetBorderWidth(value int32) {
    ListView_SetBorderWidth(l.instance, value)
}

func (l *TListView) Checkboxes() bool {
    return ListView_GetCheckboxes(l.instance)
}

func (l *TListView) SetCheckboxes(value bool) {
    ListView_SetCheckboxes(l.instance, value)
}

func (l *TListView) Color() TColor {
    return ListView_GetColor(l.instance)
}

func (l *TListView) SetColor(value TColor) {
    ListView_SetColor(l.instance, value)
}

func (l *TListView) Columns() *TListColumns {
    return ListColumnsFromInst(ListView_GetColumns(l.instance))
}

func (l *TListView) SetColumns(value *TListColumns) {
    ListView_SetColumns(l.instance, CheckPtr(value))
}

func (l *TListView) ColumnClick() bool {
    return ListView_GetColumnClick(l.instance)
}

func (l *TListView) SetColumnClick(value bool) {
    ListView_SetColumnClick(l.instance, value)
}

func (l *TListView) DoubleBuffered() bool {
    return ListView_GetDoubleBuffered(l.instance)
}

func (l *TListView) SetDoubleBuffered(value bool) {
    ListView_SetDoubleBuffered(l.instance, value)
}

func (l *TListView) Enabled() bool {
    return ListView_GetEnabled(l.instance)
}

func (l *TListView) SetEnabled(value bool) {
    ListView_SetEnabled(l.instance, value)
}

func (l *TListView) Font() *TFont {
    return FontFromInst(ListView_GetFont(l.instance))
}

func (l *TListView) SetFont(value *TFont) {
    ListView_SetFont(l.instance, CheckPtr(value))
}

func (l *TListView) FlatScrollBars() bool {
    return ListView_GetFlatScrollBars(l.instance)
}

func (l *TListView) SetFlatScrollBars(value bool) {
    ListView_SetFlatScrollBars(l.instance, value)
}

func (l *TListView) FullDrag() bool {
    return ListView_GetFullDrag(l.instance)
}

func (l *TListView) SetFullDrag(value bool) {
    ListView_SetFullDrag(l.instance, value)
}

func (l *TListView) GridLines() bool {
    return ListView_GetGridLines(l.instance)
}

func (l *TListView) SetGridLines(value bool) {
    ListView_SetGridLines(l.instance, value)
}

func (l *TListView) Groups() *TListGroups {
    return ListGroupsFromInst(ListView_GetGroups(l.instance))
}

func (l *TListView) SetGroups(value *TListGroups) {
    ListView_SetGroups(l.instance, CheckPtr(value))
}

func (l *TListView) HideSelection() bool {
    return ListView_GetHideSelection(l.instance)
}

func (l *TListView) SetHideSelection(value bool) {
    ListView_SetHideSelection(l.instance, value)
}

func (l *TListView) HotTrack() bool {
    return ListView_GetHotTrack(l.instance)
}

func (l *TListView) SetHotTrack(value bool) {
    ListView_SetHotTrack(l.instance, value)
}

func (l *TListView) HoverTime() int32 {
    return ListView_GetHoverTime(l.instance)
}

func (l *TListView) SetHoverTime(value int32) {
    ListView_SetHoverTime(l.instance, value)
}

func (l *TListView) IconOptions() *TIconOptions {
    return IconOptionsFromInst(ListView_GetIconOptions(l.instance))
}

func (l *TListView) SetIconOptions(value IObject) {
    ListView_SetIconOptions(l.instance, CheckPtr(value))
}

func (l *TListView) Items() *TListItems {
    return ListItemsFromInst(ListView_GetItems(l.instance))
}

func (l *TListView) SetItems(value *TListItems) {
    ListView_SetItems(l.instance, CheckPtr(value))
}

func (l *TListView) LargeImages() *TImageList {
    return ImageListFromInst(ListView_GetLargeImages(l.instance))
}

func (l *TListView) SetLargeImages(value IComponent) {
    ListView_SetLargeImages(l.instance, CheckPtr(value))
}

func (l *TListView) MultiSelect() bool {
    return ListView_GetMultiSelect(l.instance)
}

func (l *TListView) SetMultiSelect(value bool) {
    ListView_SetMultiSelect(l.instance, value)
}

func (l *TListView) StyleElements() TStyleElements {
    return ListView_GetStyleElements(l.instance)
}

func (l *TListView) SetStyleElements(value TStyleElements) {
    ListView_SetStyleElements(l.instance, value)
}

func (l *TListView) GroupHeaderImages() *TImageList {
    return ImageListFromInst(ListView_GetGroupHeaderImages(l.instance))
}

func (l *TListView) SetGroupHeaderImages(value IComponent) {
    ListView_SetGroupHeaderImages(l.instance, CheckPtr(value))
}

func (l *TListView) GroupView() bool {
    return ListView_GetGroupView(l.instance)
}

func (l *TListView) SetGroupView(value bool) {
    ListView_SetGroupView(l.instance, value)
}

func (l *TListView) ReadOnly() bool {
    return ListView_GetReadOnly(l.instance)
}

func (l *TListView) SetReadOnly(value bool) {
    ListView_SetReadOnly(l.instance, value)
}

func (l *TListView) RowSelect() bool {
    return ListView_GetRowSelect(l.instance)
}

func (l *TListView) SetRowSelect(value bool) {
    ListView_SetRowSelect(l.instance, value)
}

func (l *TListView) ParentColor() bool {
    return ListView_GetParentColor(l.instance)
}

func (l *TListView) SetParentColor(value bool) {
    ListView_SetParentColor(l.instance, value)
}

func (l *TListView) ParentDoubleBuffered() bool {
    return ListView_GetParentDoubleBuffered(l.instance)
}

func (l *TListView) SetParentDoubleBuffered(value bool) {
    ListView_SetParentDoubleBuffered(l.instance, value)
}

func (l *TListView) ParentFont() bool {
    return ListView_GetParentFont(l.instance)
}

func (l *TListView) SetParentFont(value bool) {
    ListView_SetParentFont(l.instance, value)
}

func (l *TListView) ParentShowHint() bool {
    return ListView_GetParentShowHint(l.instance)
}

func (l *TListView) SetParentShowHint(value bool) {
    ListView_SetParentShowHint(l.instance, value)
}

func (l *TListView) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ListView_GetPopupMenu(l.instance))
}

func (l *TListView) SetPopupMenu(value IComponent) {
    ListView_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TListView) ShowColumnHeaders() bool {
    return ListView_GetShowColumnHeaders(l.instance)
}

func (l *TListView) SetShowColumnHeaders(value bool) {
    ListView_SetShowColumnHeaders(l.instance, value)
}

func (l *TListView) ShowWorkAreas() bool {
    return ListView_GetShowWorkAreas(l.instance)
}

func (l *TListView) SetShowWorkAreas(value bool) {
    ListView_SetShowWorkAreas(l.instance, value)
}

func (l *TListView) ShowHint() bool {
    return ListView_GetShowHint(l.instance)
}

func (l *TListView) SetShowHint(value bool) {
    ListView_SetShowHint(l.instance, value)
}

func (l *TListView) SmallImages() *TImageList {
    return ImageListFromInst(ListView_GetSmallImages(l.instance))
}

func (l *TListView) SetSmallImages(value IComponent) {
    ListView_SetSmallImages(l.instance, CheckPtr(value))
}

func (l *TListView) SortType() TSortType {
    return ListView_GetSortType(l.instance)
}

func (l *TListView) SetSortType(value TSortType) {
    ListView_SetSortType(l.instance, value)
}

func (l *TListView) StateImages() *TImageList {
    return ImageListFromInst(ListView_GetStateImages(l.instance))
}

func (l *TListView) SetStateImages(value IComponent) {
    ListView_SetStateImages(l.instance, CheckPtr(value))
}

func (l *TListView) TabOrder() uint16 {
    return ListView_GetTabOrder(l.instance)
}

func (l *TListView) SetTabOrder(value uint16) {
    ListView_SetTabOrder(l.instance, value)
}

func (l *TListView) TabStop() bool {
    return ListView_GetTabStop(l.instance)
}

func (l *TListView) SetTabStop(value bool) {
    ListView_SetTabStop(l.instance, value)
}

func (l *TListView) ViewStyle() TViewStyle {
    return ListView_GetViewStyle(l.instance)
}

func (l *TListView) SetViewStyle(value TViewStyle) {
    ListView_SetViewStyle(l.instance, value)
}

func (l *TListView) Visible() bool {
    return ListView_GetVisible(l.instance)
}

func (l *TListView) SetVisible(value bool) {
    ListView_SetVisible(l.instance, value)
}

func (l *TListView) SetOnAdvancedCustomDraw(fn TLVAdvancedCustomDrawEvent) {
    ListView_SetOnAdvancedCustomDraw(l.instance, fn)
}

func (l *TListView) SetOnAdvancedCustomDrawItem(fn TLVAdvancedCustomDrawItemEvent) {
    ListView_SetOnAdvancedCustomDrawItem(l.instance, fn)
}

func (l *TListView) SetOnAdvancedCustomDrawSubItem(fn TLVAdvancedCustomDrawSubItemEvent) {
    ListView_SetOnAdvancedCustomDrawSubItem(l.instance, fn)
}

func (l *TListView) SetOnChange(fn TLVChangeEvent) {
    ListView_SetOnChange(l.instance, fn)
}

func (l *TListView) SetOnClick(fn TNotifyEvent) {
    ListView_SetOnClick(l.instance, fn)
}

func (l *TListView) SetOnColumnClick(fn TLVColumnClickEvent) {
    ListView_SetOnColumnClick(l.instance, fn)
}

func (l *TListView) SetOnColumnRightClick(fn TLVColumnRClickEvent) {
    ListView_SetOnColumnRightClick(l.instance, fn)
}

func (l *TListView) SetOnCompare(fn TLVCompareEvent) {
    ListView_SetOnCompare(l.instance, fn)
}

func (l *TListView) SetOnDblClick(fn TNotifyEvent) {
    ListView_SetOnDblClick(l.instance, fn)
}

func (l *TListView) SetOnEnter(fn TNotifyEvent) {
    ListView_SetOnEnter(l.instance, fn)
}

func (l *TListView) SetOnExit(fn TNotifyEvent) {
    ListView_SetOnExit(l.instance, fn)
}

func (l *TListView) SetOnGetImageIndex(fn TLVNotifyEvent) {
    ListView_SetOnGetImageIndex(l.instance, fn)
}

func (l *TListView) SetOnKeyDown(fn TKeyEvent) {
    ListView_SetOnKeyDown(l.instance, fn)
}

func (l *TListView) SetOnKeyPress(fn TKeyPressEvent) {
    ListView_SetOnKeyPress(l.instance, fn)
}

func (l *TListView) SetOnKeyUp(fn TKeyEvent) {
    ListView_SetOnKeyUp(l.instance, fn)
}

func (l *TListView) SetOnMouseDown(fn TMouseEvent) {
    ListView_SetOnMouseDown(l.instance, fn)
}

func (l *TListView) SetOnMouseEnter(fn TNotifyEvent) {
    ListView_SetOnMouseEnter(l.instance, fn)
}

func (l *TListView) SetOnMouseLeave(fn TNotifyEvent) {
    ListView_SetOnMouseLeave(l.instance, fn)
}

func (l *TListView) SetOnMouseMove(fn TMouseMoveEvent) {
    ListView_SetOnMouseMove(l.instance, fn)
}

func (l *TListView) SetOnMouseUp(fn TMouseEvent) {
    ListView_SetOnMouseUp(l.instance, fn)
}

func (l *TListView) SetOnResize(fn TNotifyEvent) {
    ListView_SetOnResize(l.instance, fn)
}

func (l *TListView) SetOnSelectItem(fn TLVSelectItemEvent) {
    ListView_SetOnSelectItem(l.instance, fn)
}

func (l *TListView) SetOnItemChecked(fn TLVCheckedItemEvent) {
    ListView_SetOnItemChecked(l.instance, fn)
}

func (l *TListView) Canvas() *TCanvas {
    return CanvasFromInst(ListView_GetCanvas(l.instance))
}

func (l *TListView) DropTarget() *TListItem {
    return ListItemFromInst(ListView_GetDropTarget(l.instance))
}

func (l *TListView) SetDropTarget(value *TListItem) {
    ListView_SetDropTarget(l.instance, CheckPtr(value))
}

func (l *TListView) ItemFocused() *TListItem {
    return ListItemFromInst(ListView_GetItemFocused(l.instance))
}

func (l *TListView) SetItemFocused(value *TListItem) {
    ListView_SetItemFocused(l.instance, CheckPtr(value))
}

func (l *TListView) SelCount() int32 {
    return ListView_GetSelCount(l.instance)
}

func (l *TListView) Selected() *TListItem {
    return ListItemFromInst(ListView_GetSelected(l.instance))
}

func (l *TListView) SetSelected(value *TListItem) {
    ListView_SetSelected(l.instance, CheckPtr(value))
}

func (l *TListView) TopItem() *TListItem {
    return ListItemFromInst(ListView_GetTopItem(l.instance))
}

func (l *TListView) VisibleRowCount() int32 {
    return ListView_GetVisibleRowCount(l.instance)
}

func (l *TListView) ItemIndex() int32 {
    return ListView_GetItemIndex(l.instance)
}

func (l *TListView) SetItemIndex(value int32) {
    ListView_SetItemIndex(l.instance, value)
}

func (l *TListView) Brush() *TBrush {
    return BrushFromInst(ListView_GetBrush(l.instance))
}

func (l *TListView) ControlCount() int32 {
    return ListView_GetControlCount(l.instance)
}

func (l *TListView) Handle() HWND {
    return ListView_GetHandle(l.instance)
}

func (l *TListView) ParentWindow() HWND {
    return ListView_GetParentWindow(l.instance)
}

func (l *TListView) SetParentWindow(value HWND) {
    ListView_SetParentWindow(l.instance, value)
}

func (l *TListView) BoundsRect() TRect {
    return ListView_GetBoundsRect(l.instance)
}

func (l *TListView) SetBoundsRect(value TRect) {
    ListView_SetBoundsRect(l.instance, value)
}

func (l *TListView) ClientHeight() int32 {
    return ListView_GetClientHeight(l.instance)
}

func (l *TListView) SetClientHeight(value int32) {
    ListView_SetClientHeight(l.instance, value)
}

func (l *TListView) ClientRect() TRect {
    return ListView_GetClientRect(l.instance)
}

func (l *TListView) ClientWidth() int32 {
    return ListView_GetClientWidth(l.instance)
}

func (l *TListView) SetClientWidth(value int32) {
    ListView_SetClientWidth(l.instance, value)
}

func (l *TListView) ExplicitLeft() int32 {
    return ListView_GetExplicitLeft(l.instance)
}

func (l *TListView) ExplicitTop() int32 {
    return ListView_GetExplicitTop(l.instance)
}

func (l *TListView) ExplicitWidth() int32 {
    return ListView_GetExplicitWidth(l.instance)
}

func (l *TListView) ExplicitHeight() int32 {
    return ListView_GetExplicitHeight(l.instance)
}

func (l *TListView) Parent() *TControl {
    return ControlFromInst(ListView_GetParent(l.instance))
}

func (l *TListView) SetParent(value IControl) {
    ListView_SetParent(l.instance, CheckPtr(value))
}

func (l *TListView) AlignWithMargins() bool {
    return ListView_GetAlignWithMargins(l.instance)
}

func (l *TListView) SetAlignWithMargins(value bool) {
    ListView_SetAlignWithMargins(l.instance, value)
}

func (l *TListView) Left() int32 {
    return ListView_GetLeft(l.instance)
}

func (l *TListView) SetLeft(value int32) {
    ListView_SetLeft(l.instance, value)
}

func (l *TListView) Top() int32 {
    return ListView_GetTop(l.instance)
}

func (l *TListView) SetTop(value int32) {
    ListView_SetTop(l.instance, value)
}

func (l *TListView) Width() int32 {
    return ListView_GetWidth(l.instance)
}

func (l *TListView) SetWidth(value int32) {
    ListView_SetWidth(l.instance, value)
}

func (l *TListView) Height() int32 {
    return ListView_GetHeight(l.instance)
}

func (l *TListView) SetHeight(value int32) {
    ListView_SetHeight(l.instance, value)
}

func (l *TListView) Cursor() TCursor {
    return ListView_GetCursor(l.instance)
}

func (l *TListView) SetCursor(value TCursor) {
    ListView_SetCursor(l.instance, value)
}

func (l *TListView) Hint() string {
    return ListView_GetHint(l.instance)
}

func (l *TListView) SetHint(value string) {
    ListView_SetHint(l.instance, value)
}

func (l *TListView) Margins() *TMargins {
    return MarginsFromInst(ListView_GetMargins(l.instance))
}

func (l *TListView) SetMargins(value *TMargins) {
    ListView_SetMargins(l.instance, CheckPtr(value))
}

func (l *TListView) CustomHint() *TCustomHint {
    return CustomHintFromInst(ListView_GetCustomHint(l.instance))
}

func (l *TListView) SetCustomHint(value IComponent) {
    ListView_SetCustomHint(l.instance, CheckPtr(value))
}

func (l *TListView) ComponentCount() int32 {
    return ListView_GetComponentCount(l.instance)
}

func (l *TListView) ComponentIndex() int32 {
    return ListView_GetComponentIndex(l.instance)
}

func (l *TListView) SetComponentIndex(value int32) {
    ListView_SetComponentIndex(l.instance, value)
}

func (l *TListView) Owner() *TComponent {
    return ComponentFromInst(ListView_GetOwner(l.instance))
}

func (l *TListView) Name() string {
    return ListView_GetName(l.instance)
}

func (l *TListView) SetName(value string) {
    ListView_SetName(l.instance, value)
}

func (l *TListView) Tag() int {
    return ListView_GetTag(l.instance)
}

func (l *TListView) SetTag(value int) {
    ListView_SetTag(l.instance, value)
}

func (l *TListView) Column(Index int32) *TListColumn {
    return ListColumnFromInst(ListView_GetColumn(l.instance, Index))
}

func (l *TListView) Controls(Index int32) *TControl {
    return ControlFromInst(ListView_GetControls(l.instance, Index))
}

func (l *TListView) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ListView_GetComponents(l.instance, AIndex))
}

