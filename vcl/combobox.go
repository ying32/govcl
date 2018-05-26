
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

type TComboBox struct {
    IControl
    instance uintptr
}

func NewComboBox(owner IComponent) *TComboBox {
    c := new(TComboBox)
    c.instance = ComboBox_Create(CheckPtr(owner))
    return c
}

func ComboBoxFromInst(inst uintptr) *TComboBox {
    c := new(TComboBox)
    c.instance = inst
    return c
}

func ComboBoxFromObj(obj IObject) *TComboBox {
    c := new(TComboBox)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TComboBox) Free() {
    if c.instance != 0 {
        ComboBox_Free(c.instance)
        c.instance = 0
    }
}

func (c *TComboBox) Instance() uintptr {
    return c.instance
}

func (c *TComboBox) IsValid() bool {
    return c.instance != 0
}

func (c *TComboBox) AddItem(Item string, AObject IObject) {
    ComboBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

func (c *TComboBox) Clear() {
    ComboBox_Clear(c.instance)
}

func (c *TComboBox) ClearSelection() {
    ComboBox_ClearSelection(c.instance)
}

func (c *TComboBox) DeleteSelected() {
    ComboBox_DeleteSelected(c.instance)
}

func (c *TComboBox) Focused() bool {
    return ComboBox_Focused(c.instance)
}

func (c *TComboBox) SelectAll() {
    ComboBox_SelectAll(c.instance)
}

func (c *TComboBox) CanFocus() bool {
    return ComboBox_CanFocus(c.instance)
}

func (c *TComboBox) FlipChildren(AllLevels bool) {
    ComboBox_FlipChildren(c.instance, AllLevels)
}

func (c *TComboBox) HandleAllocated() bool {
    return ComboBox_HandleAllocated(c.instance)
}

func (c *TComboBox) Invalidate() {
    ComboBox_Invalidate(c.instance)
}

func (c *TComboBox) Realign() {
    ComboBox_Realign(c.instance)
}

func (c *TComboBox) Repaint() {
    ComboBox_Repaint(c.instance)
}

func (c *TComboBox) ScaleBy(M int32, D int32) {
    ComboBox_ScaleBy(c.instance, M , D)
}

func (c *TComboBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ComboBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

func (c *TComboBox) SetFocus() {
    ComboBox_SetFocus(c.instance)
}

func (c *TComboBox) Update() {
    ComboBox_Update(c.instance)
}

func (c *TComboBox) BringToFront() {
    ComboBox_BringToFront(c.instance)
}

func (c *TComboBox) HasParent() bool {
    return ComboBox_HasParent(c.instance)
}

func (c *TComboBox) Hide() {
    ComboBox_Hide(c.instance)
}

func (c *TComboBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ComboBox_Perform(c.instance, Msg , WParam , LParam)
}

func (c *TComboBox) Refresh() {
    ComboBox_Refresh(c.instance)
}

func (c *TComboBox) SendToBack() {
    ComboBox_SendToBack(c.instance)
}

func (c *TComboBox) Show() {
    ComboBox_Show(c.instance)
}

func (c *TComboBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ComboBox_GetTextBuf(c.instance, Buffer , BufSize)
}

func (c *TComboBox) GetTextLen() int32 {
    return ComboBox_GetTextLen(c.instance)
}

func (c *TComboBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ComboBox_FindComponent(c.instance, AName))
}

func (c *TComboBox) GetNamePath() string {
    return ComboBox_GetNamePath(c.instance)
}

func (c *TComboBox) Assign(Source IObject) {
    ComboBox_Assign(c.instance, CheckPtr(Source))
}

func (c *TComboBox) ClassName() string {
    return ComboBox_ClassName(c.instance)
}

func (c *TComboBox) Equals(Obj IObject) bool {
    return ComboBox_Equals(c.instance, CheckPtr(Obj))
}

func (c *TComboBox) GetHashCode() int32 {
    return ComboBox_GetHashCode(c.instance)
}

func (c *TComboBox) ToString() string {
    return ComboBox_ToString(c.instance)
}

func (c *TComboBox) Align() TAlign {
    return ComboBox_GetAlign(c.instance)
}

func (c *TComboBox) SetAlign(value TAlign) {
    ComboBox_SetAlign(c.instance, value)
}

func (c *TComboBox) AutoComplete() bool {
    return ComboBox_GetAutoComplete(c.instance)
}

func (c *TComboBox) SetAutoComplete(value bool) {
    ComboBox_SetAutoComplete(c.instance, value)
}

func (c *TComboBox) AutoCompleteDelay() uint32 {
    return ComboBox_GetAutoCompleteDelay(c.instance)
}

func (c *TComboBox) SetAutoCompleteDelay(value uint32) {
    ComboBox_SetAutoCompleteDelay(c.instance, value)
}

func (c *TComboBox) AutoDropDown() bool {
    return ComboBox_GetAutoDropDown(c.instance)
}

func (c *TComboBox) SetAutoDropDown(value bool) {
    ComboBox_SetAutoDropDown(c.instance, value)
}

func (c *TComboBox) AutoCloseUp() bool {
    return ComboBox_GetAutoCloseUp(c.instance)
}

func (c *TComboBox) SetAutoCloseUp(value bool) {
    ComboBox_SetAutoCloseUp(c.instance, value)
}

func (c *TComboBox) BevelEdges() TBevelEdges {
    return ComboBox_GetBevelEdges(c.instance)
}

func (c *TComboBox) SetBevelEdges(value TBevelEdges) {
    ComboBox_SetBevelEdges(c.instance, value)
}

func (c *TComboBox) BevelInner() TBevelCut {
    return ComboBox_GetBevelInner(c.instance)
}

func (c *TComboBox) SetBevelInner(value TBevelCut) {
    ComboBox_SetBevelInner(c.instance, value)
}

func (c *TComboBox) BevelKind() TBevelKind {
    return ComboBox_GetBevelKind(c.instance)
}

func (c *TComboBox) SetBevelKind(value TBevelKind) {
    ComboBox_SetBevelKind(c.instance, value)
}

func (c *TComboBox) BevelOuter() TBevelCut {
    return ComboBox_GetBevelOuter(c.instance)
}

func (c *TComboBox) SetBevelOuter(value TBevelCut) {
    ComboBox_SetBevelOuter(c.instance, value)
}

func (c *TComboBox) Style() TComboBoxStyle {
    return ComboBox_GetStyle(c.instance)
}

func (c *TComboBox) SetStyle(value TComboBoxStyle) {
    ComboBox_SetStyle(c.instance, value)
}

func (c *TComboBox) Anchors() TAnchors {
    return ComboBox_GetAnchors(c.instance)
}

func (c *TComboBox) SetAnchors(value TAnchors) {
    ComboBox_SetAnchors(c.instance, value)
}

func (c *TComboBox) BiDiMode() TBiDiMode {
    return ComboBox_GetBiDiMode(c.instance)
}

func (c *TComboBox) SetBiDiMode(value TBiDiMode) {
    ComboBox_SetBiDiMode(c.instance, value)
}

func (c *TComboBox) Color() TColor {
    return ComboBox_GetColor(c.instance)
}

func (c *TComboBox) SetColor(value TColor) {
    ComboBox_SetColor(c.instance, value)
}

func (c *TComboBox) DoubleBuffered() bool {
    return ComboBox_GetDoubleBuffered(c.instance)
}

func (c *TComboBox) SetDoubleBuffered(value bool) {
    ComboBox_SetDoubleBuffered(c.instance, value)
}

func (c *TComboBox) DropDownCount() int32 {
    return ComboBox_GetDropDownCount(c.instance)
}

func (c *TComboBox) SetDropDownCount(value int32) {
    ComboBox_SetDropDownCount(c.instance, value)
}

func (c *TComboBox) Enabled() bool {
    return ComboBox_GetEnabled(c.instance)
}

func (c *TComboBox) SetEnabled(value bool) {
    ComboBox_SetEnabled(c.instance, value)
}

func (c *TComboBox) Font() *TFont {
    return FontFromInst(ComboBox_GetFont(c.instance))
}

func (c *TComboBox) SetFont(value *TFont) {
    ComboBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TComboBox) ItemHeight() int32 {
    return ComboBox_GetItemHeight(c.instance)
}

func (c *TComboBox) SetItemHeight(value int32) {
    ComboBox_SetItemHeight(c.instance, value)
}

func (c *TComboBox) ItemIndex() int32 {
    return ComboBox_GetItemIndex(c.instance)
}

func (c *TComboBox) SetItemIndex(value int32) {
    ComboBox_SetItemIndex(c.instance, value)
}

func (c *TComboBox) MaxLength() int32 {
    return ComboBox_GetMaxLength(c.instance)
}

func (c *TComboBox) SetMaxLength(value int32) {
    ComboBox_SetMaxLength(c.instance, value)
}

func (c *TComboBox) ParentColor() bool {
    return ComboBox_GetParentColor(c.instance)
}

func (c *TComboBox) SetParentColor(value bool) {
    ComboBox_SetParentColor(c.instance, value)
}

func (c *TComboBox) ParentCtl3D() bool {
    return ComboBox_GetParentCtl3D(c.instance)
}

func (c *TComboBox) SetParentCtl3D(value bool) {
    ComboBox_SetParentCtl3D(c.instance, value)
}

func (c *TComboBox) ParentDoubleBuffered() bool {
    return ComboBox_GetParentDoubleBuffered(c.instance)
}

func (c *TComboBox) SetParentDoubleBuffered(value bool) {
    ComboBox_SetParentDoubleBuffered(c.instance, value)
}

func (c *TComboBox) ParentFont() bool {
    return ComboBox_GetParentFont(c.instance)
}

func (c *TComboBox) SetParentFont(value bool) {
    ComboBox_SetParentFont(c.instance, value)
}

func (c *TComboBox) ParentShowHint() bool {
    return ComboBox_GetParentShowHint(c.instance)
}

func (c *TComboBox) SetParentShowHint(value bool) {
    ComboBox_SetParentShowHint(c.instance, value)
}

func (c *TComboBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ComboBox_GetPopupMenu(c.instance))
}

func (c *TComboBox) SetPopupMenu(value IComponent) {
    ComboBox_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TComboBox) ShowHint() bool {
    return ComboBox_GetShowHint(c.instance)
}

func (c *TComboBox) SetShowHint(value bool) {
    ComboBox_SetShowHint(c.instance, value)
}

func (c *TComboBox) Sorted() bool {
    return ComboBox_GetSorted(c.instance)
}

func (c *TComboBox) SetSorted(value bool) {
    ComboBox_SetSorted(c.instance, value)
}

func (c *TComboBox) TabOrder() uint16 {
    return ComboBox_GetTabOrder(c.instance)
}

func (c *TComboBox) SetTabOrder(value uint16) {
    ComboBox_SetTabOrder(c.instance, value)
}

func (c *TComboBox) TabStop() bool {
    return ComboBox_GetTabStop(c.instance)
}

func (c *TComboBox) SetTabStop(value bool) {
    ComboBox_SetTabStop(c.instance, value)
}

func (c *TComboBox) Text() string {
    return ComboBox_GetText(c.instance)
}

func (c *TComboBox) SetText(value string) {
    ComboBox_SetText(c.instance, value)
}

func (c *TComboBox) TextHint() string {
    return ComboBox_GetTextHint(c.instance)
}

func (c *TComboBox) SetTextHint(value string) {
    ComboBox_SetTextHint(c.instance, value)
}

func (c *TComboBox) Visible() bool {
    return ComboBox_GetVisible(c.instance)
}

func (c *TComboBox) SetVisible(value bool) {
    ComboBox_SetVisible(c.instance, value)
}

func (c *TComboBox) StyleElements() TStyleElements {
    return ComboBox_GetStyleElements(c.instance)
}

func (c *TComboBox) SetStyleElements(value TStyleElements) {
    ComboBox_SetStyleElements(c.instance, value)
}

func (c *TComboBox) SetOnChange(fn TNotifyEvent) {
    ComboBox_SetOnChange(c.instance, fn)
}

func (c *TComboBox) SetOnClick(fn TNotifyEvent) {
    ComboBox_SetOnClick(c.instance, fn)
}

func (c *TComboBox) SetOnDblClick(fn TNotifyEvent) {
    ComboBox_SetOnDblClick(c.instance, fn)
}

func (c *TComboBox) SetOnDrawItem(fn TDrawItemEvent) {
    ComboBox_SetOnDrawItem(c.instance, fn)
}

func (c *TComboBox) SetOnEnter(fn TNotifyEvent) {
    ComboBox_SetOnEnter(c.instance, fn)
}

func (c *TComboBox) SetOnExit(fn TNotifyEvent) {
    ComboBox_SetOnExit(c.instance, fn)
}

func (c *TComboBox) SetOnKeyDown(fn TKeyEvent) {
    ComboBox_SetOnKeyDown(c.instance, fn)
}

func (c *TComboBox) SetOnKeyPress(fn TKeyPressEvent) {
    ComboBox_SetOnKeyPress(c.instance, fn)
}

func (c *TComboBox) SetOnKeyUp(fn TKeyEvent) {
    ComboBox_SetOnKeyUp(c.instance, fn)
}

func (c *TComboBox) SetOnMouseEnter(fn TNotifyEvent) {
    ComboBox_SetOnMouseEnter(c.instance, fn)
}

func (c *TComboBox) SetOnMouseLeave(fn TNotifyEvent) {
    ComboBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TComboBox) Items() *TStrings {
    return StringsFromInst(ComboBox_GetItems(c.instance))
}

func (c *TComboBox) SetItems(value IObject) {
    ComboBox_SetItems(c.instance, CheckPtr(value))
}

func (c *TComboBox) SelText() string {
    return ComboBox_GetSelText(c.instance)
}

func (c *TComboBox) SetSelText(value string) {
    ComboBox_SetSelText(c.instance, value)
}

func (c *TComboBox) Canvas() *TCanvas {
    return CanvasFromInst(ComboBox_GetCanvas(c.instance))
}

func (c *TComboBox) DroppedDown() bool {
    return ComboBox_GetDroppedDown(c.instance)
}

func (c *TComboBox) SetDroppedDown(value bool) {
    ComboBox_SetDroppedDown(c.instance, value)
}

func (c *TComboBox) SelLength() int32 {
    return ComboBox_GetSelLength(c.instance)
}

func (c *TComboBox) SetSelLength(value int32) {
    ComboBox_SetSelLength(c.instance, value)
}

func (c *TComboBox) SelStart() int32 {
    return ComboBox_GetSelStart(c.instance)
}

func (c *TComboBox) SetSelStart(value int32) {
    ComboBox_SetSelStart(c.instance, value)
}

func (c *TComboBox) Brush() *TBrush {
    return BrushFromInst(ComboBox_GetBrush(c.instance))
}

func (c *TComboBox) ControlCount() int32 {
    return ComboBox_GetControlCount(c.instance)
}

func (c *TComboBox) Handle() HWND {
    return ComboBox_GetHandle(c.instance)
}

func (c *TComboBox) ParentWindow() HWND {
    return ComboBox_GetParentWindow(c.instance)
}

func (c *TComboBox) SetParentWindow(value HWND) {
    ComboBox_SetParentWindow(c.instance, value)
}

func (c *TComboBox) Action() *TAction {
    return ActionFromInst(ComboBox_GetAction(c.instance))
}

func (c *TComboBox) SetAction(value IComponent) {
    ComboBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TComboBox) BoundsRect() TRect {
    return ComboBox_GetBoundsRect(c.instance)
}

func (c *TComboBox) SetBoundsRect(value TRect) {
    ComboBox_SetBoundsRect(c.instance, value)
}

func (c *TComboBox) ClientHeight() int32 {
    return ComboBox_GetClientHeight(c.instance)
}

func (c *TComboBox) SetClientHeight(value int32) {
    ComboBox_SetClientHeight(c.instance, value)
}

func (c *TComboBox) ClientRect() TRect {
    return ComboBox_GetClientRect(c.instance)
}

func (c *TComboBox) ClientWidth() int32 {
    return ComboBox_GetClientWidth(c.instance)
}

func (c *TComboBox) SetClientWidth(value int32) {
    ComboBox_SetClientWidth(c.instance, value)
}

func (c *TComboBox) ExplicitLeft() int32 {
    return ComboBox_GetExplicitLeft(c.instance)
}

func (c *TComboBox) ExplicitTop() int32 {
    return ComboBox_GetExplicitTop(c.instance)
}

func (c *TComboBox) ExplicitWidth() int32 {
    return ComboBox_GetExplicitWidth(c.instance)
}

func (c *TComboBox) ExplicitHeight() int32 {
    return ComboBox_GetExplicitHeight(c.instance)
}

func (c *TComboBox) Parent() *TControl {
    return ControlFromInst(ComboBox_GetParent(c.instance))
}

func (c *TComboBox) SetParent(value IControl) {
    ComboBox_SetParent(c.instance, CheckPtr(value))
}

func (c *TComboBox) AlignWithMargins() bool {
    return ComboBox_GetAlignWithMargins(c.instance)
}

func (c *TComboBox) SetAlignWithMargins(value bool) {
    ComboBox_SetAlignWithMargins(c.instance, value)
}

func (c *TComboBox) Left() int32 {
    return ComboBox_GetLeft(c.instance)
}

func (c *TComboBox) SetLeft(value int32) {
    ComboBox_SetLeft(c.instance, value)
}

func (c *TComboBox) Top() int32 {
    return ComboBox_GetTop(c.instance)
}

func (c *TComboBox) SetTop(value int32) {
    ComboBox_SetTop(c.instance, value)
}

func (c *TComboBox) Width() int32 {
    return ComboBox_GetWidth(c.instance)
}

func (c *TComboBox) SetWidth(value int32) {
    ComboBox_SetWidth(c.instance, value)
}

func (c *TComboBox) Height() int32 {
    return ComboBox_GetHeight(c.instance)
}

func (c *TComboBox) SetHeight(value int32) {
    ComboBox_SetHeight(c.instance, value)
}

func (c *TComboBox) Cursor() TCursor {
    return ComboBox_GetCursor(c.instance)
}

func (c *TComboBox) SetCursor(value TCursor) {
    ComboBox_SetCursor(c.instance, value)
}

func (c *TComboBox) Hint() string {
    return ComboBox_GetHint(c.instance)
}

func (c *TComboBox) SetHint(value string) {
    ComboBox_SetHint(c.instance, value)
}

func (c *TComboBox) Margins() *TMargins {
    return MarginsFromInst(ComboBox_GetMargins(c.instance))
}

func (c *TComboBox) SetMargins(value *TMargins) {
    ComboBox_SetMargins(c.instance, CheckPtr(value))
}

func (c *TComboBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ComboBox_GetCustomHint(c.instance))
}

func (c *TComboBox) SetCustomHint(value IComponent) {
    ComboBox_SetCustomHint(c.instance, CheckPtr(value))
}

func (c *TComboBox) ComponentCount() int32 {
    return ComboBox_GetComponentCount(c.instance)
}

func (c *TComboBox) ComponentIndex() int32 {
    return ComboBox_GetComponentIndex(c.instance)
}

func (c *TComboBox) SetComponentIndex(value int32) {
    ComboBox_SetComponentIndex(c.instance, value)
}

func (c *TComboBox) Owner() *TComponent {
    return ComponentFromInst(ComboBox_GetOwner(c.instance))
}

func (c *TComboBox) Name() string {
    return ComboBox_GetName(c.instance)
}

func (c *TComboBox) SetName(value string) {
    ComboBox_SetName(c.instance, value)
}

func (c *TComboBox) Tag() int {
    return ComboBox_GetTag(c.instance)
}

func (c *TComboBox) SetTag(value int) {
    ComboBox_SetTag(c.instance, value)
}

func (c *TComboBox) Controls(Index int32) *TControl {
    return ControlFromInst(ComboBox_GetControls(c.instance, Index))
}

func (c *TComboBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ComboBox_GetComponents(c.instance, AIndex))
}

