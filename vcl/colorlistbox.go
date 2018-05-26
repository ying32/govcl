
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

type TColorListBox struct {
    IControl
    instance uintptr
}

func NewColorListBox(owner IComponent) *TColorListBox {
    c := new(TColorListBox)
    c.instance = ColorListBox_Create(CheckPtr(owner))
    return c
}

func ColorListBoxFromInst(inst uintptr) *TColorListBox {
    c := new(TColorListBox)
    c.instance = inst
    return c
}

func ColorListBoxFromObj(obj IObject) *TColorListBox {
    c := new(TColorListBox)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TColorListBox) Free() {
    if c.instance != 0 {
        ColorListBox_Free(c.instance)
        c.instance = 0
    }
}

func (c *TColorListBox) Instance() uintptr {
    return c.instance
}

func (c *TColorListBox) IsValid() bool {
    return c.instance != 0
}

func (c *TColorListBox) AddItem(Item string, AObject IObject) {
    ColorListBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

func (c *TColorListBox) Clear() {
    ColorListBox_Clear(c.instance)
}

func (c *TColorListBox) ClearSelection() {
    ColorListBox_ClearSelection(c.instance)
}

func (c *TColorListBox) DeleteSelected() {
    ColorListBox_DeleteSelected(c.instance)
}

func (c *TColorListBox) SelectAll() {
    ColorListBox_SelectAll(c.instance)
}

func (c *TColorListBox) CanFocus() bool {
    return ColorListBox_CanFocus(c.instance)
}

func (c *TColorListBox) FlipChildren(AllLevels bool) {
    ColorListBox_FlipChildren(c.instance, AllLevels)
}

func (c *TColorListBox) Focused() bool {
    return ColorListBox_Focused(c.instance)
}

func (c *TColorListBox) HandleAllocated() bool {
    return ColorListBox_HandleAllocated(c.instance)
}

func (c *TColorListBox) Invalidate() {
    ColorListBox_Invalidate(c.instance)
}

func (c *TColorListBox) Realign() {
    ColorListBox_Realign(c.instance)
}

func (c *TColorListBox) Repaint() {
    ColorListBox_Repaint(c.instance)
}

func (c *TColorListBox) ScaleBy(M int32, D int32) {
    ColorListBox_ScaleBy(c.instance, M , D)
}

func (c *TColorListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ColorListBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

func (c *TColorListBox) SetFocus() {
    ColorListBox_SetFocus(c.instance)
}

func (c *TColorListBox) Update() {
    ColorListBox_Update(c.instance)
}

func (c *TColorListBox) BringToFront() {
    ColorListBox_BringToFront(c.instance)
}

func (c *TColorListBox) HasParent() bool {
    return ColorListBox_HasParent(c.instance)
}

func (c *TColorListBox) Hide() {
    ColorListBox_Hide(c.instance)
}

func (c *TColorListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ColorListBox_Perform(c.instance, Msg , WParam , LParam)
}

func (c *TColorListBox) Refresh() {
    ColorListBox_Refresh(c.instance)
}

func (c *TColorListBox) SendToBack() {
    ColorListBox_SendToBack(c.instance)
}

func (c *TColorListBox) Show() {
    ColorListBox_Show(c.instance)
}

func (c *TColorListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ColorListBox_GetTextBuf(c.instance, Buffer , BufSize)
}

func (c *TColorListBox) GetTextLen() int32 {
    return ColorListBox_GetTextLen(c.instance)
}

func (c *TColorListBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ColorListBox_FindComponent(c.instance, AName))
}

func (c *TColorListBox) GetNamePath() string {
    return ColorListBox_GetNamePath(c.instance)
}

func (c *TColorListBox) Assign(Source IObject) {
    ColorListBox_Assign(c.instance, CheckPtr(Source))
}

func (c *TColorListBox) ClassName() string {
    return ColorListBox_ClassName(c.instance)
}

func (c *TColorListBox) Equals(Obj IObject) bool {
    return ColorListBox_Equals(c.instance, CheckPtr(Obj))
}

func (c *TColorListBox) GetHashCode() int32 {
    return ColorListBox_GetHashCode(c.instance)
}

func (c *TColorListBox) ToString() string {
    return ColorListBox_ToString(c.instance)
}

func (c *TColorListBox) Align() TAlign {
    return ColorListBox_GetAlign(c.instance)
}

func (c *TColorListBox) SetAlign(value TAlign) {
    ColorListBox_SetAlign(c.instance, value)
}

func (c *TColorListBox) AutoComplete() bool {
    return ColorListBox_GetAutoComplete(c.instance)
}

func (c *TColorListBox) SetAutoComplete(value bool) {
    ColorListBox_SetAutoComplete(c.instance, value)
}

func (c *TColorListBox) DefaultColorColor() TColor {
    return ColorListBox_GetDefaultColorColor(c.instance)
}

func (c *TColorListBox) SetDefaultColorColor(value TColor) {
    ColorListBox_SetDefaultColorColor(c.instance, value)
}

func (c *TColorListBox) NoneColorColor() TColor {
    return ColorListBox_GetNoneColorColor(c.instance)
}

func (c *TColorListBox) SetNoneColorColor(value TColor) {
    ColorListBox_SetNoneColorColor(c.instance, value)
}

func (c *TColorListBox) Selected() TColor {
    return ColorListBox_GetSelected(c.instance)
}

func (c *TColorListBox) SetSelected(value TColor) {
    ColorListBox_SetSelected(c.instance, value)
}

func (c *TColorListBox) Style() TColorBoxStyle {
    return ColorListBox_GetStyle(c.instance)
}

func (c *TColorListBox) SetStyle(value TColorBoxStyle) {
    ColorListBox_SetStyle(c.instance, value)
}

func (c *TColorListBox) Anchors() TAnchors {
    return ColorListBox_GetAnchors(c.instance)
}

func (c *TColorListBox) SetAnchors(value TAnchors) {
    ColorListBox_SetAnchors(c.instance, value)
}

func (c *TColorListBox) BevelEdges() TBevelEdges {
    return ColorListBox_GetBevelEdges(c.instance)
}

func (c *TColorListBox) SetBevelEdges(value TBevelEdges) {
    ColorListBox_SetBevelEdges(c.instance, value)
}

func (c *TColorListBox) BevelInner() TBevelCut {
    return ColorListBox_GetBevelInner(c.instance)
}

func (c *TColorListBox) SetBevelInner(value TBevelCut) {
    ColorListBox_SetBevelInner(c.instance, value)
}

func (c *TColorListBox) BevelKind() TBevelKind {
    return ColorListBox_GetBevelKind(c.instance)
}

func (c *TColorListBox) SetBevelKind(value TBevelKind) {
    ColorListBox_SetBevelKind(c.instance, value)
}

func (c *TColorListBox) BevelOuter() TBevelCut {
    return ColorListBox_GetBevelOuter(c.instance)
}

func (c *TColorListBox) SetBevelOuter(value TBevelCut) {
    ColorListBox_SetBevelOuter(c.instance, value)
}

func (c *TColorListBox) BiDiMode() TBiDiMode {
    return ColorListBox_GetBiDiMode(c.instance)
}

func (c *TColorListBox) SetBiDiMode(value TBiDiMode) {
    ColorListBox_SetBiDiMode(c.instance, value)
}

func (c *TColorListBox) Color() TColor {
    return ColorListBox_GetColor(c.instance)
}

func (c *TColorListBox) SetColor(value TColor) {
    ColorListBox_SetColor(c.instance, value)
}

func (c *TColorListBox) DoubleBuffered() bool {
    return ColorListBox_GetDoubleBuffered(c.instance)
}

func (c *TColorListBox) SetDoubleBuffered(value bool) {
    ColorListBox_SetDoubleBuffered(c.instance, value)
}

func (c *TColorListBox) Enabled() bool {
    return ColorListBox_GetEnabled(c.instance)
}

func (c *TColorListBox) SetEnabled(value bool) {
    ColorListBox_SetEnabled(c.instance, value)
}

func (c *TColorListBox) Font() *TFont {
    return FontFromInst(ColorListBox_GetFont(c.instance))
}

func (c *TColorListBox) SetFont(value *TFont) {
    ColorListBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TColorListBox) ItemHeight() int32 {
    return ColorListBox_GetItemHeight(c.instance)
}

func (c *TColorListBox) SetItemHeight(value int32) {
    ColorListBox_SetItemHeight(c.instance, value)
}

func (c *TColorListBox) ParentColor() bool {
    return ColorListBox_GetParentColor(c.instance)
}

func (c *TColorListBox) SetParentColor(value bool) {
    ColorListBox_SetParentColor(c.instance, value)
}

func (c *TColorListBox) ParentCtl3D() bool {
    return ColorListBox_GetParentCtl3D(c.instance)
}

func (c *TColorListBox) SetParentCtl3D(value bool) {
    ColorListBox_SetParentCtl3D(c.instance, value)
}

func (c *TColorListBox) ParentDoubleBuffered() bool {
    return ColorListBox_GetParentDoubleBuffered(c.instance)
}

func (c *TColorListBox) SetParentDoubleBuffered(value bool) {
    ColorListBox_SetParentDoubleBuffered(c.instance, value)
}

func (c *TColorListBox) ParentFont() bool {
    return ColorListBox_GetParentFont(c.instance)
}

func (c *TColorListBox) SetParentFont(value bool) {
    ColorListBox_SetParentFont(c.instance, value)
}

func (c *TColorListBox) ParentShowHint() bool {
    return ColorListBox_GetParentShowHint(c.instance)
}

func (c *TColorListBox) SetParentShowHint(value bool) {
    ColorListBox_SetParentShowHint(c.instance, value)
}

func (c *TColorListBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ColorListBox_GetPopupMenu(c.instance))
}

func (c *TColorListBox) SetPopupMenu(value IComponent) {
    ColorListBox_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TColorListBox) ShowHint() bool {
    return ColorListBox_GetShowHint(c.instance)
}

func (c *TColorListBox) SetShowHint(value bool) {
    ColorListBox_SetShowHint(c.instance, value)
}

func (c *TColorListBox) TabOrder() uint16 {
    return ColorListBox_GetTabOrder(c.instance)
}

func (c *TColorListBox) SetTabOrder(value uint16) {
    ColorListBox_SetTabOrder(c.instance, value)
}

func (c *TColorListBox) TabStop() bool {
    return ColorListBox_GetTabStop(c.instance)
}

func (c *TColorListBox) SetTabStop(value bool) {
    ColorListBox_SetTabStop(c.instance, value)
}

func (c *TColorListBox) Visible() bool {
    return ColorListBox_GetVisible(c.instance)
}

func (c *TColorListBox) SetVisible(value bool) {
    ColorListBox_SetVisible(c.instance, value)
}

func (c *TColorListBox) StyleElements() TStyleElements {
    return ColorListBox_GetStyleElements(c.instance)
}

func (c *TColorListBox) SetStyleElements(value TStyleElements) {
    ColorListBox_SetStyleElements(c.instance, value)
}

func (c *TColorListBox) SetOnClick(fn TNotifyEvent) {
    ColorListBox_SetOnClick(c.instance, fn)
}

func (c *TColorListBox) SetOnDblClick(fn TNotifyEvent) {
    ColorListBox_SetOnDblClick(c.instance, fn)
}

func (c *TColorListBox) SetOnEnter(fn TNotifyEvent) {
    ColorListBox_SetOnEnter(c.instance, fn)
}

func (c *TColorListBox) SetOnExit(fn TNotifyEvent) {
    ColorListBox_SetOnExit(c.instance, fn)
}

func (c *TColorListBox) SetOnKeyDown(fn TKeyEvent) {
    ColorListBox_SetOnKeyDown(c.instance, fn)
}

func (c *TColorListBox) SetOnKeyPress(fn TKeyPressEvent) {
    ColorListBox_SetOnKeyPress(c.instance, fn)
}

func (c *TColorListBox) SetOnKeyUp(fn TKeyEvent) {
    ColorListBox_SetOnKeyUp(c.instance, fn)
}

func (c *TColorListBox) SetOnMouseDown(fn TMouseEvent) {
    ColorListBox_SetOnMouseDown(c.instance, fn)
}

func (c *TColorListBox) SetOnMouseEnter(fn TNotifyEvent) {
    ColorListBox_SetOnMouseEnter(c.instance, fn)
}

func (c *TColorListBox) SetOnMouseLeave(fn TNotifyEvent) {
    ColorListBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TColorListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ColorListBox_SetOnMouseMove(c.instance, fn)
}

func (c *TColorListBox) SetOnMouseUp(fn TMouseEvent) {
    ColorListBox_SetOnMouseUp(c.instance, fn)
}

func (c *TColorListBox) AutoCompleteDelay() uint32 {
    return ColorListBox_GetAutoCompleteDelay(c.instance)
}

func (c *TColorListBox) SetAutoCompleteDelay(value uint32) {
    ColorListBox_SetAutoCompleteDelay(c.instance, value)
}

func (c *TColorListBox) Canvas() *TCanvas {
    return CanvasFromInst(ColorListBox_GetCanvas(c.instance))
}

func (c *TColorListBox) Count() int32 {
    return ColorListBox_GetCount(c.instance)
}

func (c *TColorListBox) SetCount(value int32) {
    ColorListBox_SetCount(c.instance, value)
}

func (c *TColorListBox) Items() *TStrings {
    return StringsFromInst(ColorListBox_GetItems(c.instance))
}

func (c *TColorListBox) SetItems(value IObject) {
    ColorListBox_SetItems(c.instance, CheckPtr(value))
}

func (c *TColorListBox) MultiSelect() bool {
    return ColorListBox_GetMultiSelect(c.instance)
}

func (c *TColorListBox) SetMultiSelect(value bool) {
    ColorListBox_SetMultiSelect(c.instance, value)
}

func (c *TColorListBox) SelCount() int32 {
    return ColorListBox_GetSelCount(c.instance)
}

func (c *TColorListBox) ItemIndex() int32 {
    return ColorListBox_GetItemIndex(c.instance)
}

func (c *TColorListBox) SetItemIndex(value int32) {
    ColorListBox_SetItemIndex(c.instance, value)
}

func (c *TColorListBox) Brush() *TBrush {
    return BrushFromInst(ColorListBox_GetBrush(c.instance))
}

func (c *TColorListBox) ControlCount() int32 {
    return ColorListBox_GetControlCount(c.instance)
}

func (c *TColorListBox) Handle() HWND {
    return ColorListBox_GetHandle(c.instance)
}

func (c *TColorListBox) ParentWindow() HWND {
    return ColorListBox_GetParentWindow(c.instance)
}

func (c *TColorListBox) SetParentWindow(value HWND) {
    ColorListBox_SetParentWindow(c.instance, value)
}

func (c *TColorListBox) Action() *TAction {
    return ActionFromInst(ColorListBox_GetAction(c.instance))
}

func (c *TColorListBox) SetAction(value IComponent) {
    ColorListBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TColorListBox) BoundsRect() TRect {
    return ColorListBox_GetBoundsRect(c.instance)
}

func (c *TColorListBox) SetBoundsRect(value TRect) {
    ColorListBox_SetBoundsRect(c.instance, value)
}

func (c *TColorListBox) ClientHeight() int32 {
    return ColorListBox_GetClientHeight(c.instance)
}

func (c *TColorListBox) SetClientHeight(value int32) {
    ColorListBox_SetClientHeight(c.instance, value)
}

func (c *TColorListBox) ClientRect() TRect {
    return ColorListBox_GetClientRect(c.instance)
}

func (c *TColorListBox) ClientWidth() int32 {
    return ColorListBox_GetClientWidth(c.instance)
}

func (c *TColorListBox) SetClientWidth(value int32) {
    ColorListBox_SetClientWidth(c.instance, value)
}

func (c *TColorListBox) ExplicitLeft() int32 {
    return ColorListBox_GetExplicitLeft(c.instance)
}

func (c *TColorListBox) ExplicitTop() int32 {
    return ColorListBox_GetExplicitTop(c.instance)
}

func (c *TColorListBox) ExplicitWidth() int32 {
    return ColorListBox_GetExplicitWidth(c.instance)
}

func (c *TColorListBox) ExplicitHeight() int32 {
    return ColorListBox_GetExplicitHeight(c.instance)
}

func (c *TColorListBox) Parent() *TControl {
    return ControlFromInst(ColorListBox_GetParent(c.instance))
}

func (c *TColorListBox) SetParent(value IControl) {
    ColorListBox_SetParent(c.instance, CheckPtr(value))
}

func (c *TColorListBox) AlignWithMargins() bool {
    return ColorListBox_GetAlignWithMargins(c.instance)
}

func (c *TColorListBox) SetAlignWithMargins(value bool) {
    ColorListBox_SetAlignWithMargins(c.instance, value)
}

func (c *TColorListBox) Left() int32 {
    return ColorListBox_GetLeft(c.instance)
}

func (c *TColorListBox) SetLeft(value int32) {
    ColorListBox_SetLeft(c.instance, value)
}

func (c *TColorListBox) Top() int32 {
    return ColorListBox_GetTop(c.instance)
}

func (c *TColorListBox) SetTop(value int32) {
    ColorListBox_SetTop(c.instance, value)
}

func (c *TColorListBox) Width() int32 {
    return ColorListBox_GetWidth(c.instance)
}

func (c *TColorListBox) SetWidth(value int32) {
    ColorListBox_SetWidth(c.instance, value)
}

func (c *TColorListBox) Height() int32 {
    return ColorListBox_GetHeight(c.instance)
}

func (c *TColorListBox) SetHeight(value int32) {
    ColorListBox_SetHeight(c.instance, value)
}

func (c *TColorListBox) Cursor() TCursor {
    return ColorListBox_GetCursor(c.instance)
}

func (c *TColorListBox) SetCursor(value TCursor) {
    ColorListBox_SetCursor(c.instance, value)
}

func (c *TColorListBox) Hint() string {
    return ColorListBox_GetHint(c.instance)
}

func (c *TColorListBox) SetHint(value string) {
    ColorListBox_SetHint(c.instance, value)
}

func (c *TColorListBox) Margins() *TMargins {
    return MarginsFromInst(ColorListBox_GetMargins(c.instance))
}

func (c *TColorListBox) SetMargins(value *TMargins) {
    ColorListBox_SetMargins(c.instance, CheckPtr(value))
}

func (c *TColorListBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ColorListBox_GetCustomHint(c.instance))
}

func (c *TColorListBox) SetCustomHint(value IComponent) {
    ColorListBox_SetCustomHint(c.instance, CheckPtr(value))
}

func (c *TColorListBox) ComponentCount() int32 {
    return ColorListBox_GetComponentCount(c.instance)
}

func (c *TColorListBox) ComponentIndex() int32 {
    return ColorListBox_GetComponentIndex(c.instance)
}

func (c *TColorListBox) SetComponentIndex(value int32) {
    ColorListBox_SetComponentIndex(c.instance, value)
}

func (c *TColorListBox) Owner() *TComponent {
    return ComponentFromInst(ColorListBox_GetOwner(c.instance))
}

func (c *TColorListBox) Name() string {
    return ColorListBox_GetName(c.instance)
}

func (c *TColorListBox) SetName(value string) {
    ColorListBox_SetName(c.instance, value)
}

func (c *TColorListBox) Tag() int {
    return ColorListBox_GetTag(c.instance)
}

func (c *TColorListBox) SetTag(value int) {
    ColorListBox_SetTag(c.instance, value)
}

func (c *TColorListBox) Colors(Index int32) TColor {
    return ColorListBox_GetColors(c.instance, Index)
}

func (c *TColorListBox) ColorNames(Index int32) string {
    return ColorListBox_GetColorNames(c.instance, Index)
}

func (c *TColorListBox) Controls(Index int32) *TControl {
    return ControlFromInst(ColorListBox_GetControls(c.instance, Index))
}

func (c *TColorListBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ColorListBox_GetComponents(c.instance, AIndex))
}

