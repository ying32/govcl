
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

type TCheckListBox struct {
    IControl
    instance uintptr
}

func NewCheckListBox(owner IComponent) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = CheckListBox_Create(CheckPtr(owner))
    return c
}

func CheckListBoxFromInst(inst uintptr) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = inst
    return c
}

func CheckListBoxFromObj(obj IObject) *TCheckListBox {
    c := new(TCheckListBox)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCheckListBox) Free() {
    if c.instance != 0 {
        CheckListBox_Free(c.instance)
        c.instance = 0
    }
}

func (c *TCheckListBox) Instance() uintptr {
    return c.instance
}

func (c *TCheckListBox) IsValid() bool {
    return c.instance != 0
}

func (c *TCheckListBox) CheckAll(AState TCheckBoxState, AllowGrayed bool, AllowDisabled bool) {
    CheckListBox_CheckAll(c.instance, AState , AllowGrayed , AllowDisabled)
}

func (c *TCheckListBox) AddItem(Item string, AObject IObject) {
    CheckListBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

func (c *TCheckListBox) Clear() {
    CheckListBox_Clear(c.instance)
}

func (c *TCheckListBox) ClearSelection() {
    CheckListBox_ClearSelection(c.instance)
}

func (c *TCheckListBox) DeleteSelected() {
    CheckListBox_DeleteSelected(c.instance)
}

func (c *TCheckListBox) SelectAll() {
    CheckListBox_SelectAll(c.instance)
}

func (c *TCheckListBox) CanFocus() bool {
    return CheckListBox_CanFocus(c.instance)
}

func (c *TCheckListBox) FlipChildren(AllLevels bool) {
    CheckListBox_FlipChildren(c.instance, AllLevels)
}

func (c *TCheckListBox) Focused() bool {
    return CheckListBox_Focused(c.instance)
}

func (c *TCheckListBox) HandleAllocated() bool {
    return CheckListBox_HandleAllocated(c.instance)
}

func (c *TCheckListBox) Invalidate() {
    CheckListBox_Invalidate(c.instance)
}

func (c *TCheckListBox) Realign() {
    CheckListBox_Realign(c.instance)
}

func (c *TCheckListBox) Repaint() {
    CheckListBox_Repaint(c.instance)
}

func (c *TCheckListBox) ScaleBy(M int32, D int32) {
    CheckListBox_ScaleBy(c.instance, M , D)
}

func (c *TCheckListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CheckListBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

func (c *TCheckListBox) SetFocus() {
    CheckListBox_SetFocus(c.instance)
}

func (c *TCheckListBox) Update() {
    CheckListBox_Update(c.instance)
}

func (c *TCheckListBox) BringToFront() {
    CheckListBox_BringToFront(c.instance)
}

func (c *TCheckListBox) HasParent() bool {
    return CheckListBox_HasParent(c.instance)
}

func (c *TCheckListBox) Hide() {
    CheckListBox_Hide(c.instance)
}

func (c *TCheckListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CheckListBox_Perform(c.instance, Msg , WParam , LParam)
}

func (c *TCheckListBox) Refresh() {
    CheckListBox_Refresh(c.instance)
}

func (c *TCheckListBox) SendToBack() {
    CheckListBox_SendToBack(c.instance)
}

func (c *TCheckListBox) Show() {
    CheckListBox_Show(c.instance)
}

func (c *TCheckListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CheckListBox_GetTextBuf(c.instance, Buffer , BufSize)
}

func (c *TCheckListBox) GetTextLen() int32 {
    return CheckListBox_GetTextLen(c.instance)
}

func (c *TCheckListBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CheckListBox_FindComponent(c.instance, AName))
}

func (c *TCheckListBox) GetNamePath() string {
    return CheckListBox_GetNamePath(c.instance)
}

func (c *TCheckListBox) Assign(Source IObject) {
    CheckListBox_Assign(c.instance, CheckPtr(Source))
}

func (c *TCheckListBox) ClassName() string {
    return CheckListBox_ClassName(c.instance)
}

func (c *TCheckListBox) Equals(Obj IObject) bool {
    return CheckListBox_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCheckListBox) GetHashCode() int32 {
    return CheckListBox_GetHashCode(c.instance)
}

func (c *TCheckListBox) ToString() string {
    return CheckListBox_ToString(c.instance)
}

func (c *TCheckListBox) SetOnClickCheck(fn TNotifyEvent) {
    CheckListBox_SetOnClickCheck(c.instance, fn)
}

func (c *TCheckListBox) Align() TAlign {
    return CheckListBox_GetAlign(c.instance)
}

func (c *TCheckListBox) SetAlign(value TAlign) {
    CheckListBox_SetAlign(c.instance, value)
}

func (c *TCheckListBox) AllowGrayed() bool {
    return CheckListBox_GetAllowGrayed(c.instance)
}

func (c *TCheckListBox) SetAllowGrayed(value bool) {
    CheckListBox_SetAllowGrayed(c.instance, value)
}

func (c *TCheckListBox) Anchors() TAnchors {
    return CheckListBox_GetAnchors(c.instance)
}

func (c *TCheckListBox) SetAnchors(value TAnchors) {
    CheckListBox_SetAnchors(c.instance, value)
}

func (c *TCheckListBox) AutoComplete() bool {
    return CheckListBox_GetAutoComplete(c.instance)
}

func (c *TCheckListBox) SetAutoComplete(value bool) {
    CheckListBox_SetAutoComplete(c.instance, value)
}

func (c *TCheckListBox) BevelEdges() TBevelEdges {
    return CheckListBox_GetBevelEdges(c.instance)
}

func (c *TCheckListBox) SetBevelEdges(value TBevelEdges) {
    CheckListBox_SetBevelEdges(c.instance, value)
}

func (c *TCheckListBox) BevelInner() TBevelCut {
    return CheckListBox_GetBevelInner(c.instance)
}

func (c *TCheckListBox) SetBevelInner(value TBevelCut) {
    CheckListBox_SetBevelInner(c.instance, value)
}

func (c *TCheckListBox) BevelOuter() TBevelCut {
    return CheckListBox_GetBevelOuter(c.instance)
}

func (c *TCheckListBox) SetBevelOuter(value TBevelCut) {
    CheckListBox_SetBevelOuter(c.instance, value)
}

func (c *TCheckListBox) BevelKind() TBevelKind {
    return CheckListBox_GetBevelKind(c.instance)
}

func (c *TCheckListBox) SetBevelKind(value TBevelKind) {
    CheckListBox_SetBevelKind(c.instance, value)
}

func (c *TCheckListBox) BiDiMode() TBiDiMode {
    return CheckListBox_GetBiDiMode(c.instance)
}

func (c *TCheckListBox) SetBiDiMode(value TBiDiMode) {
    CheckListBox_SetBiDiMode(c.instance, value)
}

func (c *TCheckListBox) BorderStyle() TBorderStyle {
    return CheckListBox_GetBorderStyle(c.instance)
}

func (c *TCheckListBox) SetBorderStyle(value TBorderStyle) {
    CheckListBox_SetBorderStyle(c.instance, value)
}

func (c *TCheckListBox) Color() TColor {
    return CheckListBox_GetColor(c.instance)
}

func (c *TCheckListBox) SetColor(value TColor) {
    CheckListBox_SetColor(c.instance, value)
}

func (c *TCheckListBox) Columns() int32 {
    return CheckListBox_GetColumns(c.instance)
}

func (c *TCheckListBox) SetColumns(value int32) {
    CheckListBox_SetColumns(c.instance, value)
}

func (c *TCheckListBox) DoubleBuffered() bool {
    return CheckListBox_GetDoubleBuffered(c.instance)
}

func (c *TCheckListBox) SetDoubleBuffered(value bool) {
    CheckListBox_SetDoubleBuffered(c.instance, value)
}

func (c *TCheckListBox) Enabled() bool {
    return CheckListBox_GetEnabled(c.instance)
}

func (c *TCheckListBox) SetEnabled(value bool) {
    CheckListBox_SetEnabled(c.instance, value)
}

func (c *TCheckListBox) Flat() bool {
    return CheckListBox_GetFlat(c.instance)
}

func (c *TCheckListBox) SetFlat(value bool) {
    CheckListBox_SetFlat(c.instance, value)
}

func (c *TCheckListBox) Font() *TFont {
    return FontFromInst(CheckListBox_GetFont(c.instance))
}

func (c *TCheckListBox) SetFont(value *TFont) {
    CheckListBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TCheckListBox) HeaderColor() TColor {
    return CheckListBox_GetHeaderColor(c.instance)
}

func (c *TCheckListBox) SetHeaderColor(value TColor) {
    CheckListBox_SetHeaderColor(c.instance, value)
}

func (c *TCheckListBox) HeaderBackgroundColor() TColor {
    return CheckListBox_GetHeaderBackgroundColor(c.instance)
}

func (c *TCheckListBox) SetHeaderBackgroundColor(value TColor) {
    CheckListBox_SetHeaderBackgroundColor(c.instance, value)
}

func (c *TCheckListBox) ItemHeight() int32 {
    return CheckListBox_GetItemHeight(c.instance)
}

func (c *TCheckListBox) SetItemHeight(value int32) {
    CheckListBox_SetItemHeight(c.instance, value)
}

func (c *TCheckListBox) Items() *TStrings {
    return StringsFromInst(CheckListBox_GetItems(c.instance))
}

func (c *TCheckListBox) SetItems(value IObject) {
    CheckListBox_SetItems(c.instance, CheckPtr(value))
}

func (c *TCheckListBox) ParentColor() bool {
    return CheckListBox_GetParentColor(c.instance)
}

func (c *TCheckListBox) SetParentColor(value bool) {
    CheckListBox_SetParentColor(c.instance, value)
}

func (c *TCheckListBox) ParentCtl3D() bool {
    return CheckListBox_GetParentCtl3D(c.instance)
}

func (c *TCheckListBox) SetParentCtl3D(value bool) {
    CheckListBox_SetParentCtl3D(c.instance, value)
}

func (c *TCheckListBox) ParentDoubleBuffered() bool {
    return CheckListBox_GetParentDoubleBuffered(c.instance)
}

func (c *TCheckListBox) SetParentDoubleBuffered(value bool) {
    CheckListBox_SetParentDoubleBuffered(c.instance, value)
}

func (c *TCheckListBox) ParentFont() bool {
    return CheckListBox_GetParentFont(c.instance)
}

func (c *TCheckListBox) SetParentFont(value bool) {
    CheckListBox_SetParentFont(c.instance, value)
}

func (c *TCheckListBox) ParentShowHint() bool {
    return CheckListBox_GetParentShowHint(c.instance)
}

func (c *TCheckListBox) SetParentShowHint(value bool) {
    CheckListBox_SetParentShowHint(c.instance, value)
}

func (c *TCheckListBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CheckListBox_GetPopupMenu(c.instance))
}

func (c *TCheckListBox) SetPopupMenu(value IComponent) {
    CheckListBox_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TCheckListBox) ShowHint() bool {
    return CheckListBox_GetShowHint(c.instance)
}

func (c *TCheckListBox) SetShowHint(value bool) {
    CheckListBox_SetShowHint(c.instance, value)
}

func (c *TCheckListBox) Sorted() bool {
    return CheckListBox_GetSorted(c.instance)
}

func (c *TCheckListBox) SetSorted(value bool) {
    CheckListBox_SetSorted(c.instance, value)
}

func (c *TCheckListBox) Style() TListBoxStyle {
    return CheckListBox_GetStyle(c.instance)
}

func (c *TCheckListBox) SetStyle(value TListBoxStyle) {
    CheckListBox_SetStyle(c.instance, value)
}

func (c *TCheckListBox) TabOrder() uint16 {
    return CheckListBox_GetTabOrder(c.instance)
}

func (c *TCheckListBox) SetTabOrder(value uint16) {
    CheckListBox_SetTabOrder(c.instance, value)
}

func (c *TCheckListBox) TabStop() bool {
    return CheckListBox_GetTabStop(c.instance)
}

func (c *TCheckListBox) SetTabStop(value bool) {
    CheckListBox_SetTabStop(c.instance, value)
}

func (c *TCheckListBox) TabWidth() int32 {
    return CheckListBox_GetTabWidth(c.instance)
}

func (c *TCheckListBox) SetTabWidth(value int32) {
    CheckListBox_SetTabWidth(c.instance, value)
}

func (c *TCheckListBox) Visible() bool {
    return CheckListBox_GetVisible(c.instance)
}

func (c *TCheckListBox) SetVisible(value bool) {
    CheckListBox_SetVisible(c.instance, value)
}

func (c *TCheckListBox) StyleElements() TStyleElements {
    return CheckListBox_GetStyleElements(c.instance)
}

func (c *TCheckListBox) SetStyleElements(value TStyleElements) {
    CheckListBox_SetStyleElements(c.instance, value)
}

func (c *TCheckListBox) SetOnClick(fn TNotifyEvent) {
    CheckListBox_SetOnClick(c.instance, fn)
}

func (c *TCheckListBox) SetOnDblClick(fn TNotifyEvent) {
    CheckListBox_SetOnDblClick(c.instance, fn)
}

func (c *TCheckListBox) SetOnEnter(fn TNotifyEvent) {
    CheckListBox_SetOnEnter(c.instance, fn)
}

func (c *TCheckListBox) SetOnExit(fn TNotifyEvent) {
    CheckListBox_SetOnExit(c.instance, fn)
}

func (c *TCheckListBox) SetOnKeyDown(fn TKeyEvent) {
    CheckListBox_SetOnKeyDown(c.instance, fn)
}

func (c *TCheckListBox) SetOnKeyPress(fn TKeyPressEvent) {
    CheckListBox_SetOnKeyPress(c.instance, fn)
}

func (c *TCheckListBox) SetOnKeyUp(fn TKeyEvent) {
    CheckListBox_SetOnKeyUp(c.instance, fn)
}

func (c *TCheckListBox) SetOnMouseDown(fn TMouseEvent) {
    CheckListBox_SetOnMouseDown(c.instance, fn)
}

func (c *TCheckListBox) SetOnMouseEnter(fn TNotifyEvent) {
    CheckListBox_SetOnMouseEnter(c.instance, fn)
}

func (c *TCheckListBox) SetOnMouseLeave(fn TNotifyEvent) {
    CheckListBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TCheckListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    CheckListBox_SetOnMouseMove(c.instance, fn)
}

func (c *TCheckListBox) SetOnMouseUp(fn TMouseEvent) {
    CheckListBox_SetOnMouseUp(c.instance, fn)
}

func (c *TCheckListBox) AutoCompleteDelay() uint32 {
    return CheckListBox_GetAutoCompleteDelay(c.instance)
}

func (c *TCheckListBox) SetAutoCompleteDelay(value uint32) {
    CheckListBox_SetAutoCompleteDelay(c.instance, value)
}

func (c *TCheckListBox) Canvas() *TCanvas {
    return CanvasFromInst(CheckListBox_GetCanvas(c.instance))
}

func (c *TCheckListBox) Count() int32 {
    return CheckListBox_GetCount(c.instance)
}

func (c *TCheckListBox) SetCount(value int32) {
    CheckListBox_SetCount(c.instance, value)
}

func (c *TCheckListBox) MultiSelect() bool {
    return CheckListBox_GetMultiSelect(c.instance)
}

func (c *TCheckListBox) SetMultiSelect(value bool) {
    CheckListBox_SetMultiSelect(c.instance, value)
}

func (c *TCheckListBox) SelCount() int32 {
    return CheckListBox_GetSelCount(c.instance)
}

func (c *TCheckListBox) ItemIndex() int32 {
    return CheckListBox_GetItemIndex(c.instance)
}

func (c *TCheckListBox) SetItemIndex(value int32) {
    CheckListBox_SetItemIndex(c.instance, value)
}

func (c *TCheckListBox) Brush() *TBrush {
    return BrushFromInst(CheckListBox_GetBrush(c.instance))
}

func (c *TCheckListBox) ControlCount() int32 {
    return CheckListBox_GetControlCount(c.instance)
}

func (c *TCheckListBox) Handle() HWND {
    return CheckListBox_GetHandle(c.instance)
}

func (c *TCheckListBox) ParentWindow() HWND {
    return CheckListBox_GetParentWindow(c.instance)
}

func (c *TCheckListBox) SetParentWindow(value HWND) {
    CheckListBox_SetParentWindow(c.instance, value)
}

func (c *TCheckListBox) Action() *TAction {
    return ActionFromInst(CheckListBox_GetAction(c.instance))
}

func (c *TCheckListBox) SetAction(value IComponent) {
    CheckListBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TCheckListBox) BoundsRect() TRect {
    return CheckListBox_GetBoundsRect(c.instance)
}

func (c *TCheckListBox) SetBoundsRect(value TRect) {
    CheckListBox_SetBoundsRect(c.instance, value)
}

func (c *TCheckListBox) ClientHeight() int32 {
    return CheckListBox_GetClientHeight(c.instance)
}

func (c *TCheckListBox) SetClientHeight(value int32) {
    CheckListBox_SetClientHeight(c.instance, value)
}

func (c *TCheckListBox) ClientRect() TRect {
    return CheckListBox_GetClientRect(c.instance)
}

func (c *TCheckListBox) ClientWidth() int32 {
    return CheckListBox_GetClientWidth(c.instance)
}

func (c *TCheckListBox) SetClientWidth(value int32) {
    CheckListBox_SetClientWidth(c.instance, value)
}

func (c *TCheckListBox) ExplicitLeft() int32 {
    return CheckListBox_GetExplicitLeft(c.instance)
}

func (c *TCheckListBox) ExplicitTop() int32 {
    return CheckListBox_GetExplicitTop(c.instance)
}

func (c *TCheckListBox) ExplicitWidth() int32 {
    return CheckListBox_GetExplicitWidth(c.instance)
}

func (c *TCheckListBox) ExplicitHeight() int32 {
    return CheckListBox_GetExplicitHeight(c.instance)
}

func (c *TCheckListBox) Parent() *TControl {
    return ControlFromInst(CheckListBox_GetParent(c.instance))
}

func (c *TCheckListBox) SetParent(value IControl) {
    CheckListBox_SetParent(c.instance, CheckPtr(value))
}

func (c *TCheckListBox) AlignWithMargins() bool {
    return CheckListBox_GetAlignWithMargins(c.instance)
}

func (c *TCheckListBox) SetAlignWithMargins(value bool) {
    CheckListBox_SetAlignWithMargins(c.instance, value)
}

func (c *TCheckListBox) Left() int32 {
    return CheckListBox_GetLeft(c.instance)
}

func (c *TCheckListBox) SetLeft(value int32) {
    CheckListBox_SetLeft(c.instance, value)
}

func (c *TCheckListBox) Top() int32 {
    return CheckListBox_GetTop(c.instance)
}

func (c *TCheckListBox) SetTop(value int32) {
    CheckListBox_SetTop(c.instance, value)
}

func (c *TCheckListBox) Width() int32 {
    return CheckListBox_GetWidth(c.instance)
}

func (c *TCheckListBox) SetWidth(value int32) {
    CheckListBox_SetWidth(c.instance, value)
}

func (c *TCheckListBox) Height() int32 {
    return CheckListBox_GetHeight(c.instance)
}

func (c *TCheckListBox) SetHeight(value int32) {
    CheckListBox_SetHeight(c.instance, value)
}

func (c *TCheckListBox) Cursor() TCursor {
    return CheckListBox_GetCursor(c.instance)
}

func (c *TCheckListBox) SetCursor(value TCursor) {
    CheckListBox_SetCursor(c.instance, value)
}

func (c *TCheckListBox) Hint() string {
    return CheckListBox_GetHint(c.instance)
}

func (c *TCheckListBox) SetHint(value string) {
    CheckListBox_SetHint(c.instance, value)
}

func (c *TCheckListBox) Margins() *TMargins {
    return MarginsFromInst(CheckListBox_GetMargins(c.instance))
}

func (c *TCheckListBox) SetMargins(value *TMargins) {
    CheckListBox_SetMargins(c.instance, CheckPtr(value))
}

func (c *TCheckListBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(CheckListBox_GetCustomHint(c.instance))
}

func (c *TCheckListBox) SetCustomHint(value IComponent) {
    CheckListBox_SetCustomHint(c.instance, CheckPtr(value))
}

func (c *TCheckListBox) ComponentCount() int32 {
    return CheckListBox_GetComponentCount(c.instance)
}

func (c *TCheckListBox) ComponentIndex() int32 {
    return CheckListBox_GetComponentIndex(c.instance)
}

func (c *TCheckListBox) SetComponentIndex(value int32) {
    CheckListBox_SetComponentIndex(c.instance, value)
}

func (c *TCheckListBox) Owner() *TComponent {
    return ComponentFromInst(CheckListBox_GetOwner(c.instance))
}

func (c *TCheckListBox) Name() string {
    return CheckListBox_GetName(c.instance)
}

func (c *TCheckListBox) SetName(value string) {
    CheckListBox_SetName(c.instance, value)
}

func (c *TCheckListBox) Tag() int {
    return CheckListBox_GetTag(c.instance)
}

func (c *TCheckListBox) SetTag(value int) {
    CheckListBox_SetTag(c.instance, value)
}

func (c *TCheckListBox) Checked(Index int32) bool {
    return CheckListBox_GetChecked(c.instance, Index)
}

func (c *TCheckListBox) SetChecked(Index int32, value bool) {
    CheckListBox_SetChecked(c.instance, Index, value)
}

func (c *TCheckListBox) ItemEnabled(Index int32) bool {
    return CheckListBox_GetItemEnabled(c.instance, Index)
}

func (c *TCheckListBox) SetItemEnabled(Index int32, value bool) {
    CheckListBox_SetItemEnabled(c.instance, Index, value)
}

func (c *TCheckListBox) State(Index int32) TCheckBoxState {
    return CheckListBox_GetState(c.instance, Index)
}

func (c *TCheckListBox) SetState(Index int32, value TCheckBoxState) {
    CheckListBox_SetState(c.instance, Index, value)
}

func (c *TCheckListBox) Header(Index int32) bool {
    return CheckListBox_GetHeader(c.instance, Index)
}

func (c *TCheckListBox) SetHeader(Index int32, value bool) {
    CheckListBox_SetHeader(c.instance, Index, value)
}

func (c *TCheckListBox) Selected(Index int32) bool {
    return CheckListBox_GetSelected(c.instance, Index)
}

func (c *TCheckListBox) SetSelected(Index int32, value bool) {
    CheckListBox_SetSelected(c.instance, Index, value)
}

func (c *TCheckListBox) Controls(Index int32) *TControl {
    return ControlFromInst(CheckListBox_GetControls(c.instance, Index))
}

func (c *TCheckListBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CheckListBox_GetComponents(c.instance, AIndex))
}

