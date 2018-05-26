
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

type TListBox struct {
    IControl
    instance uintptr
}

func NewListBox(owner IComponent) *TListBox {
    l := new(TListBox)
    l.instance = ListBox_Create(CheckPtr(owner))
    return l
}

func ListBoxFromInst(inst uintptr) *TListBox {
    l := new(TListBox)
    l.instance = inst
    return l
}

func ListBoxFromObj(obj IObject) *TListBox {
    l := new(TListBox)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListBox) Free() {
    if l.instance != 0 {
        ListBox_Free(l.instance)
        l.instance = 0
    }
}

func (l *TListBox) Instance() uintptr {
    return l.instance
}

func (l *TListBox) IsValid() bool {
    return l.instance != 0
}

func (l *TListBox) AddItem(Item string, AObject IObject) {
    ListBox_AddItem(l.instance, Item , CheckPtr(AObject))
}

func (l *TListBox) Clear() {
    ListBox_Clear(l.instance)
}

func (l *TListBox) ClearSelection() {
    ListBox_ClearSelection(l.instance)
}

func (l *TListBox) DeleteSelected() {
    ListBox_DeleteSelected(l.instance)
}

func (l *TListBox) SelectAll() {
    ListBox_SelectAll(l.instance)
}

func (l *TListBox) CanFocus() bool {
    return ListBox_CanFocus(l.instance)
}

func (l *TListBox) FlipChildren(AllLevels bool) {
    ListBox_FlipChildren(l.instance, AllLevels)
}

func (l *TListBox) Focused() bool {
    return ListBox_Focused(l.instance)
}

func (l *TListBox) HandleAllocated() bool {
    return ListBox_HandleAllocated(l.instance)
}

func (l *TListBox) Invalidate() {
    ListBox_Invalidate(l.instance)
}

func (l *TListBox) Realign() {
    ListBox_Realign(l.instance)
}

func (l *TListBox) Repaint() {
    ListBox_Repaint(l.instance)
}

func (l *TListBox) ScaleBy(M int32, D int32) {
    ListBox_ScaleBy(l.instance, M , D)
}

func (l *TListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ListBox_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

func (l *TListBox) SetFocus() {
    ListBox_SetFocus(l.instance)
}

func (l *TListBox) Update() {
    ListBox_Update(l.instance)
}

func (l *TListBox) BringToFront() {
    ListBox_BringToFront(l.instance)
}

func (l *TListBox) HasParent() bool {
    return ListBox_HasParent(l.instance)
}

func (l *TListBox) Hide() {
    ListBox_Hide(l.instance)
}

func (l *TListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ListBox_Perform(l.instance, Msg , WParam , LParam)
}

func (l *TListBox) Refresh() {
    ListBox_Refresh(l.instance)
}

func (l *TListBox) SendToBack() {
    ListBox_SendToBack(l.instance)
}

func (l *TListBox) Show() {
    ListBox_Show(l.instance)
}

func (l *TListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ListBox_GetTextBuf(l.instance, Buffer , BufSize)
}

func (l *TListBox) GetTextLen() int32 {
    return ListBox_GetTextLen(l.instance)
}

func (l *TListBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ListBox_FindComponent(l.instance, AName))
}

func (l *TListBox) GetNamePath() string {
    return ListBox_GetNamePath(l.instance)
}

func (l *TListBox) Assign(Source IObject) {
    ListBox_Assign(l.instance, CheckPtr(Source))
}

func (l *TListBox) ClassName() string {
    return ListBox_ClassName(l.instance)
}

func (l *TListBox) Equals(Obj IObject) bool {
    return ListBox_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListBox) GetHashCode() int32 {
    return ListBox_GetHashCode(l.instance)
}

func (l *TListBox) ToString() string {
    return ListBox_ToString(l.instance)
}

func (l *TListBox) Style() TListBoxStyle {
    return ListBox_GetStyle(l.instance)
}

func (l *TListBox) SetStyle(value TListBoxStyle) {
    ListBox_SetStyle(l.instance, value)
}

func (l *TListBox) AutoComplete() bool {
    return ListBox_GetAutoComplete(l.instance)
}

func (l *TListBox) SetAutoComplete(value bool) {
    ListBox_SetAutoComplete(l.instance, value)
}

func (l *TListBox) AutoCompleteDelay() uint32 {
    return ListBox_GetAutoCompleteDelay(l.instance)
}

func (l *TListBox) SetAutoCompleteDelay(value uint32) {
    ListBox_SetAutoCompleteDelay(l.instance, value)
}

func (l *TListBox) Align() TAlign {
    return ListBox_GetAlign(l.instance)
}

func (l *TListBox) SetAlign(value TAlign) {
    ListBox_SetAlign(l.instance, value)
}

func (l *TListBox) Anchors() TAnchors {
    return ListBox_GetAnchors(l.instance)
}

func (l *TListBox) SetAnchors(value TAnchors) {
    ListBox_SetAnchors(l.instance, value)
}

func (l *TListBox) BevelEdges() TBevelEdges {
    return ListBox_GetBevelEdges(l.instance)
}

func (l *TListBox) SetBevelEdges(value TBevelEdges) {
    ListBox_SetBevelEdges(l.instance, value)
}

func (l *TListBox) BevelInner() TBevelCut {
    return ListBox_GetBevelInner(l.instance)
}

func (l *TListBox) SetBevelInner(value TBevelCut) {
    ListBox_SetBevelInner(l.instance, value)
}

func (l *TListBox) BevelKind() TBevelKind {
    return ListBox_GetBevelKind(l.instance)
}

func (l *TListBox) SetBevelKind(value TBevelKind) {
    ListBox_SetBevelKind(l.instance, value)
}

func (l *TListBox) BevelOuter() TBevelCut {
    return ListBox_GetBevelOuter(l.instance)
}

func (l *TListBox) SetBevelOuter(value TBevelCut) {
    ListBox_SetBevelOuter(l.instance, value)
}

func (l *TListBox) BiDiMode() TBiDiMode {
    return ListBox_GetBiDiMode(l.instance)
}

func (l *TListBox) SetBiDiMode(value TBiDiMode) {
    ListBox_SetBiDiMode(l.instance, value)
}

func (l *TListBox) BorderStyle() TBorderStyle {
    return ListBox_GetBorderStyle(l.instance)
}

func (l *TListBox) SetBorderStyle(value TBorderStyle) {
    ListBox_SetBorderStyle(l.instance, value)
}

func (l *TListBox) Color() TColor {
    return ListBox_GetColor(l.instance)
}

func (l *TListBox) SetColor(value TColor) {
    ListBox_SetColor(l.instance, value)
}

func (l *TListBox) Columns() int32 {
    return ListBox_GetColumns(l.instance)
}

func (l *TListBox) SetColumns(value int32) {
    ListBox_SetColumns(l.instance, value)
}

func (l *TListBox) DoubleBuffered() bool {
    return ListBox_GetDoubleBuffered(l.instance)
}

func (l *TListBox) SetDoubleBuffered(value bool) {
    ListBox_SetDoubleBuffered(l.instance, value)
}

func (l *TListBox) Enabled() bool {
    return ListBox_GetEnabled(l.instance)
}

func (l *TListBox) SetEnabled(value bool) {
    ListBox_SetEnabled(l.instance, value)
}

func (l *TListBox) Font() *TFont {
    return FontFromInst(ListBox_GetFont(l.instance))
}

func (l *TListBox) SetFont(value *TFont) {
    ListBox_SetFont(l.instance, CheckPtr(value))
}

func (l *TListBox) ItemHeight() int32 {
    return ListBox_GetItemHeight(l.instance)
}

func (l *TListBox) SetItemHeight(value int32) {
    ListBox_SetItemHeight(l.instance, value)
}

func (l *TListBox) Items() *TStrings {
    return StringsFromInst(ListBox_GetItems(l.instance))
}

func (l *TListBox) SetItems(value IObject) {
    ListBox_SetItems(l.instance, CheckPtr(value))
}

func (l *TListBox) MultiSelect() bool {
    return ListBox_GetMultiSelect(l.instance)
}

func (l *TListBox) SetMultiSelect(value bool) {
    ListBox_SetMultiSelect(l.instance, value)
}

func (l *TListBox) ParentColor() bool {
    return ListBox_GetParentColor(l.instance)
}

func (l *TListBox) SetParentColor(value bool) {
    ListBox_SetParentColor(l.instance, value)
}

func (l *TListBox) ParentCtl3D() bool {
    return ListBox_GetParentCtl3D(l.instance)
}

func (l *TListBox) SetParentCtl3D(value bool) {
    ListBox_SetParentCtl3D(l.instance, value)
}

func (l *TListBox) ParentDoubleBuffered() bool {
    return ListBox_GetParentDoubleBuffered(l.instance)
}

func (l *TListBox) SetParentDoubleBuffered(value bool) {
    ListBox_SetParentDoubleBuffered(l.instance, value)
}

func (l *TListBox) ParentFont() bool {
    return ListBox_GetParentFont(l.instance)
}

func (l *TListBox) SetParentFont(value bool) {
    ListBox_SetParentFont(l.instance, value)
}

func (l *TListBox) ParentShowHint() bool {
    return ListBox_GetParentShowHint(l.instance)
}

func (l *TListBox) SetParentShowHint(value bool) {
    ListBox_SetParentShowHint(l.instance, value)
}

func (l *TListBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ListBox_GetPopupMenu(l.instance))
}

func (l *TListBox) SetPopupMenu(value IComponent) {
    ListBox_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TListBox) ShowHint() bool {
    return ListBox_GetShowHint(l.instance)
}

func (l *TListBox) SetShowHint(value bool) {
    ListBox_SetShowHint(l.instance, value)
}

func (l *TListBox) Sorted() bool {
    return ListBox_GetSorted(l.instance)
}

func (l *TListBox) SetSorted(value bool) {
    ListBox_SetSorted(l.instance, value)
}

func (l *TListBox) TabOrder() uint16 {
    return ListBox_GetTabOrder(l.instance)
}

func (l *TListBox) SetTabOrder(value uint16) {
    ListBox_SetTabOrder(l.instance, value)
}

func (l *TListBox) TabStop() bool {
    return ListBox_GetTabStop(l.instance)
}

func (l *TListBox) SetTabStop(value bool) {
    ListBox_SetTabStop(l.instance, value)
}

func (l *TListBox) TabWidth() int32 {
    return ListBox_GetTabWidth(l.instance)
}

func (l *TListBox) SetTabWidth(value int32) {
    ListBox_SetTabWidth(l.instance, value)
}

func (l *TListBox) Visible() bool {
    return ListBox_GetVisible(l.instance)
}

func (l *TListBox) SetVisible(value bool) {
    ListBox_SetVisible(l.instance, value)
}

func (l *TListBox) StyleElements() TStyleElements {
    return ListBox_GetStyleElements(l.instance)
}

func (l *TListBox) SetStyleElements(value TStyleElements) {
    ListBox_SetStyleElements(l.instance, value)
}

func (l *TListBox) SetOnClick(fn TNotifyEvent) {
    ListBox_SetOnClick(l.instance, fn)
}

func (l *TListBox) SetOnDblClick(fn TNotifyEvent) {
    ListBox_SetOnDblClick(l.instance, fn)
}

func (l *TListBox) SetOnDrawItem(fn TDrawItemEvent) {
    ListBox_SetOnDrawItem(l.instance, fn)
}

func (l *TListBox) SetOnEnter(fn TNotifyEvent) {
    ListBox_SetOnEnter(l.instance, fn)
}

func (l *TListBox) SetOnExit(fn TNotifyEvent) {
    ListBox_SetOnExit(l.instance, fn)
}

func (l *TListBox) SetOnKeyDown(fn TKeyEvent) {
    ListBox_SetOnKeyDown(l.instance, fn)
}

func (l *TListBox) SetOnKeyPress(fn TKeyPressEvent) {
    ListBox_SetOnKeyPress(l.instance, fn)
}

func (l *TListBox) SetOnKeyUp(fn TKeyEvent) {
    ListBox_SetOnKeyUp(l.instance, fn)
}

func (l *TListBox) SetOnMouseDown(fn TMouseEvent) {
    ListBox_SetOnMouseDown(l.instance, fn)
}

func (l *TListBox) SetOnMouseEnter(fn TNotifyEvent) {
    ListBox_SetOnMouseEnter(l.instance, fn)
}

func (l *TListBox) SetOnMouseLeave(fn TNotifyEvent) {
    ListBox_SetOnMouseLeave(l.instance, fn)
}

func (l *TListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ListBox_SetOnMouseMove(l.instance, fn)
}

func (l *TListBox) SetOnMouseUp(fn TMouseEvent) {
    ListBox_SetOnMouseUp(l.instance, fn)
}

func (l *TListBox) Canvas() *TCanvas {
    return CanvasFromInst(ListBox_GetCanvas(l.instance))
}

func (l *TListBox) Count() int32 {
    return ListBox_GetCount(l.instance)
}

func (l *TListBox) SetCount(value int32) {
    ListBox_SetCount(l.instance, value)
}

func (l *TListBox) SelCount() int32 {
    return ListBox_GetSelCount(l.instance)
}

func (l *TListBox) ItemIndex() int32 {
    return ListBox_GetItemIndex(l.instance)
}

func (l *TListBox) SetItemIndex(value int32) {
    ListBox_SetItemIndex(l.instance, value)
}

func (l *TListBox) Brush() *TBrush {
    return BrushFromInst(ListBox_GetBrush(l.instance))
}

func (l *TListBox) ControlCount() int32 {
    return ListBox_GetControlCount(l.instance)
}

func (l *TListBox) Handle() HWND {
    return ListBox_GetHandle(l.instance)
}

func (l *TListBox) ParentWindow() HWND {
    return ListBox_GetParentWindow(l.instance)
}

func (l *TListBox) SetParentWindow(value HWND) {
    ListBox_SetParentWindow(l.instance, value)
}

func (l *TListBox) Action() *TAction {
    return ActionFromInst(ListBox_GetAction(l.instance))
}

func (l *TListBox) SetAction(value IComponent) {
    ListBox_SetAction(l.instance, CheckPtr(value))
}

func (l *TListBox) BoundsRect() TRect {
    return ListBox_GetBoundsRect(l.instance)
}

func (l *TListBox) SetBoundsRect(value TRect) {
    ListBox_SetBoundsRect(l.instance, value)
}

func (l *TListBox) ClientHeight() int32 {
    return ListBox_GetClientHeight(l.instance)
}

func (l *TListBox) SetClientHeight(value int32) {
    ListBox_SetClientHeight(l.instance, value)
}

func (l *TListBox) ClientRect() TRect {
    return ListBox_GetClientRect(l.instance)
}

func (l *TListBox) ClientWidth() int32 {
    return ListBox_GetClientWidth(l.instance)
}

func (l *TListBox) SetClientWidth(value int32) {
    ListBox_SetClientWidth(l.instance, value)
}

func (l *TListBox) ExplicitLeft() int32 {
    return ListBox_GetExplicitLeft(l.instance)
}

func (l *TListBox) ExplicitTop() int32 {
    return ListBox_GetExplicitTop(l.instance)
}

func (l *TListBox) ExplicitWidth() int32 {
    return ListBox_GetExplicitWidth(l.instance)
}

func (l *TListBox) ExplicitHeight() int32 {
    return ListBox_GetExplicitHeight(l.instance)
}

func (l *TListBox) Parent() *TControl {
    return ControlFromInst(ListBox_GetParent(l.instance))
}

func (l *TListBox) SetParent(value IControl) {
    ListBox_SetParent(l.instance, CheckPtr(value))
}

func (l *TListBox) AlignWithMargins() bool {
    return ListBox_GetAlignWithMargins(l.instance)
}

func (l *TListBox) SetAlignWithMargins(value bool) {
    ListBox_SetAlignWithMargins(l.instance, value)
}

func (l *TListBox) Left() int32 {
    return ListBox_GetLeft(l.instance)
}

func (l *TListBox) SetLeft(value int32) {
    ListBox_SetLeft(l.instance, value)
}

func (l *TListBox) Top() int32 {
    return ListBox_GetTop(l.instance)
}

func (l *TListBox) SetTop(value int32) {
    ListBox_SetTop(l.instance, value)
}

func (l *TListBox) Width() int32 {
    return ListBox_GetWidth(l.instance)
}

func (l *TListBox) SetWidth(value int32) {
    ListBox_SetWidth(l.instance, value)
}

func (l *TListBox) Height() int32 {
    return ListBox_GetHeight(l.instance)
}

func (l *TListBox) SetHeight(value int32) {
    ListBox_SetHeight(l.instance, value)
}

func (l *TListBox) Cursor() TCursor {
    return ListBox_GetCursor(l.instance)
}

func (l *TListBox) SetCursor(value TCursor) {
    ListBox_SetCursor(l.instance, value)
}

func (l *TListBox) Hint() string {
    return ListBox_GetHint(l.instance)
}

func (l *TListBox) SetHint(value string) {
    ListBox_SetHint(l.instance, value)
}

func (l *TListBox) Margins() *TMargins {
    return MarginsFromInst(ListBox_GetMargins(l.instance))
}

func (l *TListBox) SetMargins(value *TMargins) {
    ListBox_SetMargins(l.instance, CheckPtr(value))
}

func (l *TListBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ListBox_GetCustomHint(l.instance))
}

func (l *TListBox) SetCustomHint(value IComponent) {
    ListBox_SetCustomHint(l.instance, CheckPtr(value))
}

func (l *TListBox) ComponentCount() int32 {
    return ListBox_GetComponentCount(l.instance)
}

func (l *TListBox) ComponentIndex() int32 {
    return ListBox_GetComponentIndex(l.instance)
}

func (l *TListBox) SetComponentIndex(value int32) {
    ListBox_SetComponentIndex(l.instance, value)
}

func (l *TListBox) Owner() *TComponent {
    return ComponentFromInst(ListBox_GetOwner(l.instance))
}

func (l *TListBox) Name() string {
    return ListBox_GetName(l.instance)
}

func (l *TListBox) SetName(value string) {
    ListBox_SetName(l.instance, value)
}

func (l *TListBox) Tag() int {
    return ListBox_GetTag(l.instance)
}

func (l *TListBox) SetTag(value int) {
    ListBox_SetTag(l.instance, value)
}

func (l *TListBox) Selected(Index int32) bool {
    return ListBox_GetSelected(l.instance, Index)
}

func (l *TListBox) SetSelected(Index int32, value bool) {
    ListBox_SetSelected(l.instance, Index, value)
}

func (l *TListBox) Controls(Index int32) *TControl {
    return ControlFromInst(ListBox_GetControls(l.instance, Index))
}

func (l *TListBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ListBox_GetComponents(l.instance, AIndex))
}

