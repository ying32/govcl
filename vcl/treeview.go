
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

type TTreeView struct {
    IControl
    instance uintptr
}

func NewTreeView(owner IComponent) *TTreeView {
    t := new(TTreeView)
    t.instance = TreeView_Create(CheckPtr(owner))
    return t
}

func TreeViewFromInst(inst uintptr) *TTreeView {
    t := new(TTreeView)
    t.instance = inst
    return t
}

func TreeViewFromObj(obj IObject) *TTreeView {
    t := new(TTreeView)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTreeView) Free() {
    if t.instance != 0 {
        TreeView_Free(t.instance)
        t.instance = 0
    }
}

func (t *TTreeView) Instance() uintptr {
    return t.instance
}

func (t *TTreeView) IsValid() bool {
    return t.instance != 0
}

func (t *TTreeView) AlphaSort(ARecurse bool) bool {
    return TreeView_AlphaSort(t.instance, ARecurse)
}

func (t *TTreeView) FullCollapse() {
    TreeView_FullCollapse(t.instance)
}

func (t *TTreeView) FullExpand() {
    TreeView_FullExpand(t.instance)
}

func (t *TTreeView) GetNodeAt(X int32, Y int32) *TTreeNode {
    return TreeNodeFromInst(TreeView_GetNodeAt(t.instance, X , Y))
}

func (t *TTreeView) IsEditing() bool {
    return TreeView_IsEditing(t.instance)
}

func (t *TTreeView) LoadFromFile(FileName string) {
    TreeView_LoadFromFile(t.instance, FileName)
}

func (t *TTreeView) LoadFromStream(Stream IObject) {
    TreeView_LoadFromStream(t.instance, CheckPtr(Stream))
}

func (t *TTreeView) SaveToFile(FileName string) {
    TreeView_SaveToFile(t.instance, FileName)
}

func (t *TTreeView) SaveToStream(Stream IObject) {
    TreeView_SaveToStream(t.instance, CheckPtr(Stream))
}

func (t *TTreeView) Deselect(Node *TTreeNode) {
    TreeView_Deselect(t.instance, CheckPtr(Node))
}

func (t *TTreeView) Subselect(Node *TTreeNode, Validate bool) {
    TreeView_Subselect(t.instance, CheckPtr(Node), Validate)
}

func (t *TTreeView) ClearSelection(KeepPrimary bool) {
    TreeView_ClearSelection(t.instance, KeepPrimary)
}

func (t *TTreeView) FindNextToSelect() *TTreeNode {
    return TreeNodeFromInst(TreeView_FindNextToSelect(t.instance))
}

func (t *TTreeView) CustomSort(SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
    return TreeView_CustomSort(t.instance, SortProc , Data , ARecurse)
}

func (t *TTreeView) CanFocus() bool {
    return TreeView_CanFocus(t.instance)
}

func (t *TTreeView) FlipChildren(AllLevels bool) {
    TreeView_FlipChildren(t.instance, AllLevels)
}

func (t *TTreeView) Focused() bool {
    return TreeView_Focused(t.instance)
}

func (t *TTreeView) HandleAllocated() bool {
    return TreeView_HandleAllocated(t.instance)
}

func (t *TTreeView) Invalidate() {
    TreeView_Invalidate(t.instance)
}

func (t *TTreeView) Realign() {
    TreeView_Realign(t.instance)
}

func (t *TTreeView) Repaint() {
    TreeView_Repaint(t.instance)
}

func (t *TTreeView) ScaleBy(M int32, D int32) {
    TreeView_ScaleBy(t.instance, M , D)
}

func (t *TTreeView) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    TreeView_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

func (t *TTreeView) SetFocus() {
    TreeView_SetFocus(t.instance)
}

func (t *TTreeView) Update() {
    TreeView_Update(t.instance)
}

func (t *TTreeView) BringToFront() {
    TreeView_BringToFront(t.instance)
}

func (t *TTreeView) HasParent() bool {
    return TreeView_HasParent(t.instance)
}

func (t *TTreeView) Hide() {
    TreeView_Hide(t.instance)
}

func (t *TTreeView) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return TreeView_Perform(t.instance, Msg , WParam , LParam)
}

func (t *TTreeView) Refresh() {
    TreeView_Refresh(t.instance)
}

func (t *TTreeView) SendToBack() {
    TreeView_SendToBack(t.instance)
}

func (t *TTreeView) Show() {
    TreeView_Show(t.instance)
}

func (t *TTreeView) GetTextBuf(Buffer string, BufSize int32) int32 {
    return TreeView_GetTextBuf(t.instance, Buffer , BufSize)
}

func (t *TTreeView) GetTextLen() int32 {
    return TreeView_GetTextLen(t.instance)
}

func (t *TTreeView) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TreeView_FindComponent(t.instance, AName))
}

func (t *TTreeView) GetNamePath() string {
    return TreeView_GetNamePath(t.instance)
}

func (t *TTreeView) Assign(Source IObject) {
    TreeView_Assign(t.instance, CheckPtr(Source))
}

func (t *TTreeView) ClassName() string {
    return TreeView_ClassName(t.instance)
}

func (t *TTreeView) Equals(Obj IObject) bool {
    return TreeView_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTreeView) GetHashCode() int32 {
    return TreeView_GetHashCode(t.instance)
}

func (t *TTreeView) ToString() string {
    return TreeView_ToString(t.instance)
}

func (t *TTreeView) Align() TAlign {
    return TreeView_GetAlign(t.instance)
}

func (t *TTreeView) SetAlign(value TAlign) {
    TreeView_SetAlign(t.instance, value)
}

func (t *TTreeView) Anchors() TAnchors {
    return TreeView_GetAnchors(t.instance)
}

func (t *TTreeView) SetAnchors(value TAnchors) {
    TreeView_SetAnchors(t.instance, value)
}

func (t *TTreeView) AutoExpand() bool {
    return TreeView_GetAutoExpand(t.instance)
}

func (t *TTreeView) SetAutoExpand(value bool) {
    TreeView_SetAutoExpand(t.instance, value)
}

func (t *TTreeView) BevelEdges() TBevelEdges {
    return TreeView_GetBevelEdges(t.instance)
}

func (t *TTreeView) SetBevelEdges(value TBevelEdges) {
    TreeView_SetBevelEdges(t.instance, value)
}

func (t *TTreeView) BevelInner() TBevelCut {
    return TreeView_GetBevelInner(t.instance)
}

func (t *TTreeView) SetBevelInner(value TBevelCut) {
    TreeView_SetBevelInner(t.instance, value)
}

func (t *TTreeView) BevelOuter() TBevelCut {
    return TreeView_GetBevelOuter(t.instance)
}

func (t *TTreeView) SetBevelOuter(value TBevelCut) {
    TreeView_SetBevelOuter(t.instance, value)
}

func (t *TTreeView) BevelKind() TBevelKind {
    return TreeView_GetBevelKind(t.instance)
}

func (t *TTreeView) SetBevelKind(value TBevelKind) {
    TreeView_SetBevelKind(t.instance, value)
}

func (t *TTreeView) BiDiMode() TBiDiMode {
    return TreeView_GetBiDiMode(t.instance)
}

func (t *TTreeView) SetBiDiMode(value TBiDiMode) {
    TreeView_SetBiDiMode(t.instance, value)
}

func (t *TTreeView) BorderStyle() TBorderStyle {
    return TreeView_GetBorderStyle(t.instance)
}

func (t *TTreeView) SetBorderStyle(value TBorderStyle) {
    TreeView_SetBorderStyle(t.instance, value)
}

func (t *TTreeView) BorderWidth() int32 {
    return TreeView_GetBorderWidth(t.instance)
}

func (t *TTreeView) SetBorderWidth(value int32) {
    TreeView_SetBorderWidth(t.instance, value)
}

func (t *TTreeView) ChangeDelay() int32 {
    return TreeView_GetChangeDelay(t.instance)
}

func (t *TTreeView) SetChangeDelay(value int32) {
    TreeView_SetChangeDelay(t.instance, value)
}

func (t *TTreeView) Color() TColor {
    return TreeView_GetColor(t.instance)
}

func (t *TTreeView) SetColor(value TColor) {
    TreeView_SetColor(t.instance, value)
}

func (t *TTreeView) DoubleBuffered() bool {
    return TreeView_GetDoubleBuffered(t.instance)
}

func (t *TTreeView) SetDoubleBuffered(value bool) {
    TreeView_SetDoubleBuffered(t.instance, value)
}

func (t *TTreeView) Enabled() bool {
    return TreeView_GetEnabled(t.instance)
}

func (t *TTreeView) SetEnabled(value bool) {
    TreeView_SetEnabled(t.instance, value)
}

func (t *TTreeView) Font() *TFont {
    return FontFromInst(TreeView_GetFont(t.instance))
}

func (t *TTreeView) SetFont(value *TFont) {
    TreeView_SetFont(t.instance, CheckPtr(value))
}

func (t *TTreeView) HideSelection() bool {
    return TreeView_GetHideSelection(t.instance)
}

func (t *TTreeView) SetHideSelection(value bool) {
    TreeView_SetHideSelection(t.instance, value)
}

func (t *TTreeView) HotTrack() bool {
    return TreeView_GetHotTrack(t.instance)
}

func (t *TTreeView) SetHotTrack(value bool) {
    TreeView_SetHotTrack(t.instance, value)
}

func (t *TTreeView) Images() *TImageList {
    return ImageListFromInst(TreeView_GetImages(t.instance))
}

func (t *TTreeView) SetImages(value IComponent) {
    TreeView_SetImages(t.instance, CheckPtr(value))
}

func (t *TTreeView) Indent() int32 {
    return TreeView_GetIndent(t.instance)
}

func (t *TTreeView) SetIndent(value int32) {
    TreeView_SetIndent(t.instance, value)
}

func (t *TTreeView) MultiSelect() bool {
    return TreeView_GetMultiSelect(t.instance)
}

func (t *TTreeView) SetMultiSelect(value bool) {
    TreeView_SetMultiSelect(t.instance, value)
}

func (t *TTreeView) MultiSelectStyle() TMultiSelectStyle {
    return TreeView_GetMultiSelectStyle(t.instance)
}

func (t *TTreeView) SetMultiSelectStyle(value TMultiSelectStyle) {
    TreeView_SetMultiSelectStyle(t.instance, value)
}

func (t *TTreeView) ParentColor() bool {
    return TreeView_GetParentColor(t.instance)
}

func (t *TTreeView) SetParentColor(value bool) {
    TreeView_SetParentColor(t.instance, value)
}

func (t *TTreeView) ParentCtl3D() bool {
    return TreeView_GetParentCtl3D(t.instance)
}

func (t *TTreeView) SetParentCtl3D(value bool) {
    TreeView_SetParentCtl3D(t.instance, value)
}

func (t *TTreeView) ParentDoubleBuffered() bool {
    return TreeView_GetParentDoubleBuffered(t.instance)
}

func (t *TTreeView) SetParentDoubleBuffered(value bool) {
    TreeView_SetParentDoubleBuffered(t.instance, value)
}

func (t *TTreeView) ParentFont() bool {
    return TreeView_GetParentFont(t.instance)
}

func (t *TTreeView) SetParentFont(value bool) {
    TreeView_SetParentFont(t.instance, value)
}

func (t *TTreeView) ParentShowHint() bool {
    return TreeView_GetParentShowHint(t.instance)
}

func (t *TTreeView) SetParentShowHint(value bool) {
    TreeView_SetParentShowHint(t.instance, value)
}

func (t *TTreeView) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TreeView_GetPopupMenu(t.instance))
}

func (t *TTreeView) SetPopupMenu(value IComponent) {
    TreeView_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TTreeView) ReadOnly() bool {
    return TreeView_GetReadOnly(t.instance)
}

func (t *TTreeView) SetReadOnly(value bool) {
    TreeView_SetReadOnly(t.instance, value)
}

func (t *TTreeView) RightClickSelect() bool {
    return TreeView_GetRightClickSelect(t.instance)
}

func (t *TTreeView) SetRightClickSelect(value bool) {
    TreeView_SetRightClickSelect(t.instance, value)
}

func (t *TTreeView) RowSelect() bool {
    return TreeView_GetRowSelect(t.instance)
}

func (t *TTreeView) SetRowSelect(value bool) {
    TreeView_SetRowSelect(t.instance, value)
}

func (t *TTreeView) ShowButtons() bool {
    return TreeView_GetShowButtons(t.instance)
}

func (t *TTreeView) SetShowButtons(value bool) {
    TreeView_SetShowButtons(t.instance, value)
}

func (t *TTreeView) ShowHint() bool {
    return TreeView_GetShowHint(t.instance)
}

func (t *TTreeView) SetShowHint(value bool) {
    TreeView_SetShowHint(t.instance, value)
}

func (t *TTreeView) ShowLines() bool {
    return TreeView_GetShowLines(t.instance)
}

func (t *TTreeView) SetShowLines(value bool) {
    TreeView_SetShowLines(t.instance, value)
}

func (t *TTreeView) ShowRoot() bool {
    return TreeView_GetShowRoot(t.instance)
}

func (t *TTreeView) SetShowRoot(value bool) {
    TreeView_SetShowRoot(t.instance, value)
}

func (t *TTreeView) SortType() TSortType {
    return TreeView_GetSortType(t.instance)
}

func (t *TTreeView) SetSortType(value TSortType) {
    TreeView_SetSortType(t.instance, value)
}

func (t *TTreeView) StateImages() *TImageList {
    return ImageListFromInst(TreeView_GetStateImages(t.instance))
}

func (t *TTreeView) SetStateImages(value IComponent) {
    TreeView_SetStateImages(t.instance, CheckPtr(value))
}

func (t *TTreeView) TabOrder() uint16 {
    return TreeView_GetTabOrder(t.instance)
}

func (t *TTreeView) SetTabOrder(value uint16) {
    TreeView_SetTabOrder(t.instance, value)
}

func (t *TTreeView) TabStop() bool {
    return TreeView_GetTabStop(t.instance)
}

func (t *TTreeView) SetTabStop(value bool) {
    TreeView_SetTabStop(t.instance, value)
}

func (t *TTreeView) ToolTips() bool {
    return TreeView_GetToolTips(t.instance)
}

func (t *TTreeView) SetToolTips(value bool) {
    TreeView_SetToolTips(t.instance, value)
}

func (t *TTreeView) Visible() bool {
    return TreeView_GetVisible(t.instance)
}

func (t *TTreeView) SetVisible(value bool) {
    TreeView_SetVisible(t.instance, value)
}

func (t *TTreeView) StyleElements() TStyleElements {
    return TreeView_GetStyleElements(t.instance)
}

func (t *TTreeView) SetStyleElements(value TStyleElements) {
    TreeView_SetStyleElements(t.instance, value)
}

func (t *TTreeView) SetOnAdvancedCustomDraw(fn TTVAdvancedCustomDrawEvent) {
    TreeView_SetOnAdvancedCustomDraw(t.instance, fn)
}

func (t *TTreeView) SetOnAdvancedCustomDrawItem(fn TTVAdvancedCustomDrawItemEvent) {
    TreeView_SetOnAdvancedCustomDrawItem(t.instance, fn)
}

func (t *TTreeView) SetOnChange(fn TTVChangedEvent) {
    TreeView_SetOnChange(t.instance, fn)
}

func (t *TTreeView) SetOnClick(fn TNotifyEvent) {
    TreeView_SetOnClick(t.instance, fn)
}

func (t *TTreeView) SetOnCompare(fn TTVCompareEvent) {
    TreeView_SetOnCompare(t.instance, fn)
}

func (t *TTreeView) SetOnDblClick(fn TNotifyEvent) {
    TreeView_SetOnDblClick(t.instance, fn)
}

func (t *TTreeView) SetOnEnter(fn TNotifyEvent) {
    TreeView_SetOnEnter(t.instance, fn)
}

func (t *TTreeView) SetOnExit(fn TNotifyEvent) {
    TreeView_SetOnExit(t.instance, fn)
}

func (t *TTreeView) SetOnGetImageIndex(fn TTVExpandedEvent) {
    TreeView_SetOnGetImageIndex(t.instance, fn)
}

func (t *TTreeView) SetOnGetSelectedIndex(fn TTVExpandedEvent) {
    TreeView_SetOnGetSelectedIndex(t.instance, fn)
}

func (t *TTreeView) SetOnKeyDown(fn TKeyEvent) {
    TreeView_SetOnKeyDown(t.instance, fn)
}

func (t *TTreeView) SetOnKeyPress(fn TKeyPressEvent) {
    TreeView_SetOnKeyPress(t.instance, fn)
}

func (t *TTreeView) SetOnKeyUp(fn TKeyEvent) {
    TreeView_SetOnKeyUp(t.instance, fn)
}

func (t *TTreeView) SetOnMouseDown(fn TMouseEvent) {
    TreeView_SetOnMouseDown(t.instance, fn)
}

func (t *TTreeView) SetOnMouseEnter(fn TNotifyEvent) {
    TreeView_SetOnMouseEnter(t.instance, fn)
}

func (t *TTreeView) SetOnMouseLeave(fn TNotifyEvent) {
    TreeView_SetOnMouseLeave(t.instance, fn)
}

func (t *TTreeView) SetOnMouseMove(fn TMouseMoveEvent) {
    TreeView_SetOnMouseMove(t.instance, fn)
}

func (t *TTreeView) SetOnMouseUp(fn TMouseEvent) {
    TreeView_SetOnMouseUp(t.instance, fn)
}

func (t *TTreeView) Items() *TTreeNodes {
    return TreeNodesFromInst(TreeView_GetItems(t.instance))
}

func (t *TTreeView) SetItems(value *TTreeNodes) {
    TreeView_SetItems(t.instance, CheckPtr(value))
}

func (t *TTreeView) Canvas() *TCanvas {
    return CanvasFromInst(TreeView_GetCanvas(t.instance))
}

func (t *TTreeView) DropTarget() *TTreeNode {
    return TreeNodeFromInst(TreeView_GetDropTarget(t.instance))
}

func (t *TTreeView) SetDropTarget(value *TTreeNode) {
    TreeView_SetDropTarget(t.instance, CheckPtr(value))
}

func (t *TTreeView) Selected() *TTreeNode {
    return TreeNodeFromInst(TreeView_GetSelected(t.instance))
}

func (t *TTreeView) SetSelected(value *TTreeNode) {
    TreeView_SetSelected(t.instance, CheckPtr(value))
}

func (t *TTreeView) TopItem() *TTreeNode {
    return TreeNodeFromInst(TreeView_GetTopItem(t.instance))
}

func (t *TTreeView) SetTopItem(value *TTreeNode) {
    TreeView_SetTopItem(t.instance, CheckPtr(value))
}

func (t *TTreeView) SelectionCount() uint32 {
    return TreeView_GetSelectionCount(t.instance)
}

func (t *TTreeView) Brush() *TBrush {
    return BrushFromInst(TreeView_GetBrush(t.instance))
}

func (t *TTreeView) ControlCount() int32 {
    return TreeView_GetControlCount(t.instance)
}

func (t *TTreeView) Handle() HWND {
    return TreeView_GetHandle(t.instance)
}

func (t *TTreeView) ParentWindow() HWND {
    return TreeView_GetParentWindow(t.instance)
}

func (t *TTreeView) SetParentWindow(value HWND) {
    TreeView_SetParentWindow(t.instance, value)
}

func (t *TTreeView) Action() *TAction {
    return ActionFromInst(TreeView_GetAction(t.instance))
}

func (t *TTreeView) SetAction(value IComponent) {
    TreeView_SetAction(t.instance, CheckPtr(value))
}

func (t *TTreeView) BoundsRect() TRect {
    return TreeView_GetBoundsRect(t.instance)
}

func (t *TTreeView) SetBoundsRect(value TRect) {
    TreeView_SetBoundsRect(t.instance, value)
}

func (t *TTreeView) ClientHeight() int32 {
    return TreeView_GetClientHeight(t.instance)
}

func (t *TTreeView) SetClientHeight(value int32) {
    TreeView_SetClientHeight(t.instance, value)
}

func (t *TTreeView) ClientRect() TRect {
    return TreeView_GetClientRect(t.instance)
}

func (t *TTreeView) ClientWidth() int32 {
    return TreeView_GetClientWidth(t.instance)
}

func (t *TTreeView) SetClientWidth(value int32) {
    TreeView_SetClientWidth(t.instance, value)
}

func (t *TTreeView) ExplicitLeft() int32 {
    return TreeView_GetExplicitLeft(t.instance)
}

func (t *TTreeView) ExplicitTop() int32 {
    return TreeView_GetExplicitTop(t.instance)
}

func (t *TTreeView) ExplicitWidth() int32 {
    return TreeView_GetExplicitWidth(t.instance)
}

func (t *TTreeView) ExplicitHeight() int32 {
    return TreeView_GetExplicitHeight(t.instance)
}

func (t *TTreeView) Parent() *TControl {
    return ControlFromInst(TreeView_GetParent(t.instance))
}

func (t *TTreeView) SetParent(value IControl) {
    TreeView_SetParent(t.instance, CheckPtr(value))
}

func (t *TTreeView) AlignWithMargins() bool {
    return TreeView_GetAlignWithMargins(t.instance)
}

func (t *TTreeView) SetAlignWithMargins(value bool) {
    TreeView_SetAlignWithMargins(t.instance, value)
}

func (t *TTreeView) Left() int32 {
    return TreeView_GetLeft(t.instance)
}

func (t *TTreeView) SetLeft(value int32) {
    TreeView_SetLeft(t.instance, value)
}

func (t *TTreeView) Top() int32 {
    return TreeView_GetTop(t.instance)
}

func (t *TTreeView) SetTop(value int32) {
    TreeView_SetTop(t.instance, value)
}

func (t *TTreeView) Width() int32 {
    return TreeView_GetWidth(t.instance)
}

func (t *TTreeView) SetWidth(value int32) {
    TreeView_SetWidth(t.instance, value)
}

func (t *TTreeView) Height() int32 {
    return TreeView_GetHeight(t.instance)
}

func (t *TTreeView) SetHeight(value int32) {
    TreeView_SetHeight(t.instance, value)
}

func (t *TTreeView) Cursor() TCursor {
    return TreeView_GetCursor(t.instance)
}

func (t *TTreeView) SetCursor(value TCursor) {
    TreeView_SetCursor(t.instance, value)
}

func (t *TTreeView) Hint() string {
    return TreeView_GetHint(t.instance)
}

func (t *TTreeView) SetHint(value string) {
    TreeView_SetHint(t.instance, value)
}

func (t *TTreeView) Margins() *TMargins {
    return MarginsFromInst(TreeView_GetMargins(t.instance))
}

func (t *TTreeView) SetMargins(value *TMargins) {
    TreeView_SetMargins(t.instance, CheckPtr(value))
}

func (t *TTreeView) CustomHint() *TCustomHint {
    return CustomHintFromInst(TreeView_GetCustomHint(t.instance))
}

func (t *TTreeView) SetCustomHint(value IComponent) {
    TreeView_SetCustomHint(t.instance, CheckPtr(value))
}

func (t *TTreeView) ComponentCount() int32 {
    return TreeView_GetComponentCount(t.instance)
}

func (t *TTreeView) ComponentIndex() int32 {
    return TreeView_GetComponentIndex(t.instance)
}

func (t *TTreeView) SetComponentIndex(value int32) {
    TreeView_SetComponentIndex(t.instance, value)
}

func (t *TTreeView) Owner() *TComponent {
    return ComponentFromInst(TreeView_GetOwner(t.instance))
}

func (t *TTreeView) Name() string {
    return TreeView_GetName(t.instance)
}

func (t *TTreeView) SetName(value string) {
    TreeView_SetName(t.instance, value)
}

func (t *TTreeView) Tag() int {
    return TreeView_GetTag(t.instance)
}

func (t *TTreeView) SetTag(value int) {
    TreeView_SetTag(t.instance, value)
}

func (t *TTreeView) Selections(Index int32) *TTreeNode {
    return TreeNodeFromInst(TreeView_GetSelections(t.instance, Index))
}

func (t *TTreeView) Controls(Index int32) *TControl {
    return ControlFromInst(TreeView_GetControls(t.instance, Index))
}

func (t *TTreeView) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TreeView_GetComponents(t.instance, AIndex))
}

