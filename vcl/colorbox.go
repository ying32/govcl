
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

type TColorBox struct {
    IControl
    instance uintptr
}

func NewColorBox(owner IComponent) *TColorBox {
    c := new(TColorBox)
    c.instance = ColorBox_Create(CheckPtr(owner))
    return c
}

func ColorBoxFromInst(inst uintptr) *TColorBox {
    c := new(TColorBox)
    c.instance = inst
    return c
}

func ColorBoxFromObj(obj IObject) *TColorBox {
    c := new(TColorBox)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TColorBox) Free() {
    if c.instance != 0 {
        ColorBox_Free(c.instance)
        c.instance = 0
    }
}

func (c *TColorBox) Instance() uintptr {
    return c.instance
}

func (c *TColorBox) IsValid() bool {
    return c.instance != 0
}

func (c *TColorBox) AddItem(Item string, AObject IObject) {
    ColorBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

func (c *TColorBox) Clear() {
    ColorBox_Clear(c.instance)
}

func (c *TColorBox) ClearSelection() {
    ColorBox_ClearSelection(c.instance)
}

func (c *TColorBox) DeleteSelected() {
    ColorBox_DeleteSelected(c.instance)
}

func (c *TColorBox) Focused() bool {
    return ColorBox_Focused(c.instance)
}

func (c *TColorBox) SelectAll() {
    ColorBox_SelectAll(c.instance)
}

func (c *TColorBox) CanFocus() bool {
    return ColorBox_CanFocus(c.instance)
}

func (c *TColorBox) FlipChildren(AllLevels bool) {
    ColorBox_FlipChildren(c.instance, AllLevels)
}

func (c *TColorBox) HandleAllocated() bool {
    return ColorBox_HandleAllocated(c.instance)
}

func (c *TColorBox) Invalidate() {
    ColorBox_Invalidate(c.instance)
}

func (c *TColorBox) Realign() {
    ColorBox_Realign(c.instance)
}

func (c *TColorBox) Repaint() {
    ColorBox_Repaint(c.instance)
}

func (c *TColorBox) ScaleBy(M int32, D int32) {
    ColorBox_ScaleBy(c.instance, M , D)
}

func (c *TColorBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ColorBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

func (c *TColorBox) SetFocus() {
    ColorBox_SetFocus(c.instance)
}

func (c *TColorBox) Update() {
    ColorBox_Update(c.instance)
}

func (c *TColorBox) BringToFront() {
    ColorBox_BringToFront(c.instance)
}

func (c *TColorBox) HasParent() bool {
    return ColorBox_HasParent(c.instance)
}

func (c *TColorBox) Hide() {
    ColorBox_Hide(c.instance)
}

func (c *TColorBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ColorBox_Perform(c.instance, Msg , WParam , LParam)
}

func (c *TColorBox) Refresh() {
    ColorBox_Refresh(c.instance)
}

func (c *TColorBox) SendToBack() {
    ColorBox_SendToBack(c.instance)
}

func (c *TColorBox) Show() {
    ColorBox_Show(c.instance)
}

func (c *TColorBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ColorBox_GetTextBuf(c.instance, Buffer , BufSize)
}

func (c *TColorBox) GetTextLen() int32 {
    return ColorBox_GetTextLen(c.instance)
}

func (c *TColorBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ColorBox_FindComponent(c.instance, AName))
}

func (c *TColorBox) GetNamePath() string {
    return ColorBox_GetNamePath(c.instance)
}

func (c *TColorBox) Assign(Source IObject) {
    ColorBox_Assign(c.instance, CheckPtr(Source))
}

func (c *TColorBox) ClassName() string {
    return ColorBox_ClassName(c.instance)
}

func (c *TColorBox) Equals(Obj IObject) bool {
    return ColorBox_Equals(c.instance, CheckPtr(Obj))
}

func (c *TColorBox) GetHashCode() int32 {
    return ColorBox_GetHashCode(c.instance)
}

func (c *TColorBox) ToString() string {
    return ColorBox_ToString(c.instance)
}

func (c *TColorBox) Align() TAlign {
    return ColorBox_GetAlign(c.instance)
}

func (c *TColorBox) SetAlign(value TAlign) {
    ColorBox_SetAlign(c.instance, value)
}

func (c *TColorBox) AutoComplete() bool {
    return ColorBox_GetAutoComplete(c.instance)
}

func (c *TColorBox) SetAutoComplete(value bool) {
    ColorBox_SetAutoComplete(c.instance, value)
}

func (c *TColorBox) AutoDropDown() bool {
    return ColorBox_GetAutoDropDown(c.instance)
}

func (c *TColorBox) SetAutoDropDown(value bool) {
    ColorBox_SetAutoDropDown(c.instance, value)
}

func (c *TColorBox) DefaultColorColor() TColor {
    return ColorBox_GetDefaultColorColor(c.instance)
}

func (c *TColorBox) SetDefaultColorColor(value TColor) {
    ColorBox_SetDefaultColorColor(c.instance, value)
}

func (c *TColorBox) NoneColorColor() TColor {
    return ColorBox_GetNoneColorColor(c.instance)
}

func (c *TColorBox) SetNoneColorColor(value TColor) {
    ColorBox_SetNoneColorColor(c.instance, value)
}

func (c *TColorBox) Selected() TColor {
    return ColorBox_GetSelected(c.instance)
}

func (c *TColorBox) SetSelected(value TColor) {
    ColorBox_SetSelected(c.instance, value)
}

func (c *TColorBox) Style() TColorBoxStyle {
    return ColorBox_GetStyle(c.instance)
}

func (c *TColorBox) SetStyle(value TColorBoxStyle) {
    ColorBox_SetStyle(c.instance, value)
}

func (c *TColorBox) Anchors() TAnchors {
    return ColorBox_GetAnchors(c.instance)
}

func (c *TColorBox) SetAnchors(value TAnchors) {
    ColorBox_SetAnchors(c.instance, value)
}

func (c *TColorBox) BevelEdges() TBevelEdges {
    return ColorBox_GetBevelEdges(c.instance)
}

func (c *TColorBox) SetBevelEdges(value TBevelEdges) {
    ColorBox_SetBevelEdges(c.instance, value)
}

func (c *TColorBox) BevelInner() TBevelCut {
    return ColorBox_GetBevelInner(c.instance)
}

func (c *TColorBox) SetBevelInner(value TBevelCut) {
    ColorBox_SetBevelInner(c.instance, value)
}

func (c *TColorBox) BevelKind() TBevelKind {
    return ColorBox_GetBevelKind(c.instance)
}

func (c *TColorBox) SetBevelKind(value TBevelKind) {
    ColorBox_SetBevelKind(c.instance, value)
}

func (c *TColorBox) BevelOuter() TBevelCut {
    return ColorBox_GetBevelOuter(c.instance)
}

func (c *TColorBox) SetBevelOuter(value TBevelCut) {
    ColorBox_SetBevelOuter(c.instance, value)
}

func (c *TColorBox) BiDiMode() TBiDiMode {
    return ColorBox_GetBiDiMode(c.instance)
}

func (c *TColorBox) SetBiDiMode(value TBiDiMode) {
    ColorBox_SetBiDiMode(c.instance, value)
}

func (c *TColorBox) Color() TColor {
    return ColorBox_GetColor(c.instance)
}

func (c *TColorBox) SetColor(value TColor) {
    ColorBox_SetColor(c.instance, value)
}

func (c *TColorBox) DoubleBuffered() bool {
    return ColorBox_GetDoubleBuffered(c.instance)
}

func (c *TColorBox) SetDoubleBuffered(value bool) {
    ColorBox_SetDoubleBuffered(c.instance, value)
}

func (c *TColorBox) DropDownCount() int32 {
    return ColorBox_GetDropDownCount(c.instance)
}

func (c *TColorBox) SetDropDownCount(value int32) {
    ColorBox_SetDropDownCount(c.instance, value)
}

func (c *TColorBox) Enabled() bool {
    return ColorBox_GetEnabled(c.instance)
}

func (c *TColorBox) SetEnabled(value bool) {
    ColorBox_SetEnabled(c.instance, value)
}

func (c *TColorBox) Font() *TFont {
    return FontFromInst(ColorBox_GetFont(c.instance))
}

func (c *TColorBox) SetFont(value *TFont) {
    ColorBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TColorBox) ItemHeight() int32 {
    return ColorBox_GetItemHeight(c.instance)
}

func (c *TColorBox) SetItemHeight(value int32) {
    ColorBox_SetItemHeight(c.instance, value)
}

func (c *TColorBox) ParentColor() bool {
    return ColorBox_GetParentColor(c.instance)
}

func (c *TColorBox) SetParentColor(value bool) {
    ColorBox_SetParentColor(c.instance, value)
}

func (c *TColorBox) ParentCtl3D() bool {
    return ColorBox_GetParentCtl3D(c.instance)
}

func (c *TColorBox) SetParentCtl3D(value bool) {
    ColorBox_SetParentCtl3D(c.instance, value)
}

func (c *TColorBox) ParentDoubleBuffered() bool {
    return ColorBox_GetParentDoubleBuffered(c.instance)
}

func (c *TColorBox) SetParentDoubleBuffered(value bool) {
    ColorBox_SetParentDoubleBuffered(c.instance, value)
}

func (c *TColorBox) ParentFont() bool {
    return ColorBox_GetParentFont(c.instance)
}

func (c *TColorBox) SetParentFont(value bool) {
    ColorBox_SetParentFont(c.instance, value)
}

func (c *TColorBox) ParentShowHint() bool {
    return ColorBox_GetParentShowHint(c.instance)
}

func (c *TColorBox) SetParentShowHint(value bool) {
    ColorBox_SetParentShowHint(c.instance, value)
}

func (c *TColorBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ColorBox_GetPopupMenu(c.instance))
}

func (c *TColorBox) SetPopupMenu(value IComponent) {
    ColorBox_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TColorBox) ShowHint() bool {
    return ColorBox_GetShowHint(c.instance)
}

func (c *TColorBox) SetShowHint(value bool) {
    ColorBox_SetShowHint(c.instance, value)
}

func (c *TColorBox) TabOrder() uint16 {
    return ColorBox_GetTabOrder(c.instance)
}

func (c *TColorBox) SetTabOrder(value uint16) {
    ColorBox_SetTabOrder(c.instance, value)
}

func (c *TColorBox) TabStop() bool {
    return ColorBox_GetTabStop(c.instance)
}

func (c *TColorBox) SetTabStop(value bool) {
    ColorBox_SetTabStop(c.instance, value)
}

func (c *TColorBox) Visible() bool {
    return ColorBox_GetVisible(c.instance)
}

func (c *TColorBox) SetVisible(value bool) {
    ColorBox_SetVisible(c.instance, value)
}

func (c *TColorBox) StyleElements() TStyleElements {
    return ColorBox_GetStyleElements(c.instance)
}

func (c *TColorBox) SetStyleElements(value TStyleElements) {
    ColorBox_SetStyleElements(c.instance, value)
}

func (c *TColorBox) SetOnChange(fn TNotifyEvent) {
    ColorBox_SetOnChange(c.instance, fn)
}

func (c *TColorBox) SetOnClick(fn TNotifyEvent) {
    ColorBox_SetOnClick(c.instance, fn)
}

func (c *TColorBox) SetOnEnter(fn TNotifyEvent) {
    ColorBox_SetOnEnter(c.instance, fn)
}

func (c *TColorBox) SetOnExit(fn TNotifyEvent) {
    ColorBox_SetOnExit(c.instance, fn)
}

func (c *TColorBox) SetOnKeyDown(fn TKeyEvent) {
    ColorBox_SetOnKeyDown(c.instance, fn)
}

func (c *TColorBox) SetOnKeyPress(fn TKeyPressEvent) {
    ColorBox_SetOnKeyPress(c.instance, fn)
}

func (c *TColorBox) SetOnKeyUp(fn TKeyEvent) {
    ColorBox_SetOnKeyUp(c.instance, fn)
}

func (c *TColorBox) SetOnMouseEnter(fn TNotifyEvent) {
    ColorBox_SetOnMouseEnter(c.instance, fn)
}

func (c *TColorBox) SetOnMouseLeave(fn TNotifyEvent) {
    ColorBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TColorBox) AutoCompleteDelay() uint32 {
    return ColorBox_GetAutoCompleteDelay(c.instance)
}

func (c *TColorBox) SetAutoCompleteDelay(value uint32) {
    ColorBox_SetAutoCompleteDelay(c.instance, value)
}

func (c *TColorBox) AutoCloseUp() bool {
    return ColorBox_GetAutoCloseUp(c.instance)
}

func (c *TColorBox) SetAutoCloseUp(value bool) {
    ColorBox_SetAutoCloseUp(c.instance, value)
}

func (c *TColorBox) SelText() string {
    return ColorBox_GetSelText(c.instance)
}

func (c *TColorBox) SetSelText(value string) {
    ColorBox_SetSelText(c.instance, value)
}

func (c *TColorBox) TextHint() string {
    return ColorBox_GetTextHint(c.instance)
}

func (c *TColorBox) SetTextHint(value string) {
    ColorBox_SetTextHint(c.instance, value)
}

func (c *TColorBox) Canvas() *TCanvas {
    return CanvasFromInst(ColorBox_GetCanvas(c.instance))
}

func (c *TColorBox) DroppedDown() bool {
    return ColorBox_GetDroppedDown(c.instance)
}

func (c *TColorBox) SetDroppedDown(value bool) {
    ColorBox_SetDroppedDown(c.instance, value)
}

func (c *TColorBox) Items() *TStrings {
    return StringsFromInst(ColorBox_GetItems(c.instance))
}

func (c *TColorBox) SetItems(value IObject) {
    ColorBox_SetItems(c.instance, CheckPtr(value))
}

func (c *TColorBox) SelLength() int32 {
    return ColorBox_GetSelLength(c.instance)
}

func (c *TColorBox) SetSelLength(value int32) {
    ColorBox_SetSelLength(c.instance, value)
}

func (c *TColorBox) SelStart() int32 {
    return ColorBox_GetSelStart(c.instance)
}

func (c *TColorBox) SetSelStart(value int32) {
    ColorBox_SetSelStart(c.instance, value)
}

func (c *TColorBox) ItemIndex() int32 {
    return ColorBox_GetItemIndex(c.instance)
}

func (c *TColorBox) SetItemIndex(value int32) {
    ColorBox_SetItemIndex(c.instance, value)
}

func (c *TColorBox) Brush() *TBrush {
    return BrushFromInst(ColorBox_GetBrush(c.instance))
}

func (c *TColorBox) ControlCount() int32 {
    return ColorBox_GetControlCount(c.instance)
}

func (c *TColorBox) Handle() HWND {
    return ColorBox_GetHandle(c.instance)
}

func (c *TColorBox) ParentWindow() HWND {
    return ColorBox_GetParentWindow(c.instance)
}

func (c *TColorBox) SetParentWindow(value HWND) {
    ColorBox_SetParentWindow(c.instance, value)
}

func (c *TColorBox) Action() *TAction {
    return ActionFromInst(ColorBox_GetAction(c.instance))
}

func (c *TColorBox) SetAction(value IComponent) {
    ColorBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TColorBox) BoundsRect() TRect {
    return ColorBox_GetBoundsRect(c.instance)
}

func (c *TColorBox) SetBoundsRect(value TRect) {
    ColorBox_SetBoundsRect(c.instance, value)
}

func (c *TColorBox) ClientHeight() int32 {
    return ColorBox_GetClientHeight(c.instance)
}

func (c *TColorBox) SetClientHeight(value int32) {
    ColorBox_SetClientHeight(c.instance, value)
}

func (c *TColorBox) ClientRect() TRect {
    return ColorBox_GetClientRect(c.instance)
}

func (c *TColorBox) ClientWidth() int32 {
    return ColorBox_GetClientWidth(c.instance)
}

func (c *TColorBox) SetClientWidth(value int32) {
    ColorBox_SetClientWidth(c.instance, value)
}

func (c *TColorBox) ExplicitLeft() int32 {
    return ColorBox_GetExplicitLeft(c.instance)
}

func (c *TColorBox) ExplicitTop() int32 {
    return ColorBox_GetExplicitTop(c.instance)
}

func (c *TColorBox) ExplicitWidth() int32 {
    return ColorBox_GetExplicitWidth(c.instance)
}

func (c *TColorBox) ExplicitHeight() int32 {
    return ColorBox_GetExplicitHeight(c.instance)
}

func (c *TColorBox) Parent() *TControl {
    return ControlFromInst(ColorBox_GetParent(c.instance))
}

func (c *TColorBox) SetParent(value IControl) {
    ColorBox_SetParent(c.instance, CheckPtr(value))
}

func (c *TColorBox) AlignWithMargins() bool {
    return ColorBox_GetAlignWithMargins(c.instance)
}

func (c *TColorBox) SetAlignWithMargins(value bool) {
    ColorBox_SetAlignWithMargins(c.instance, value)
}

func (c *TColorBox) Left() int32 {
    return ColorBox_GetLeft(c.instance)
}

func (c *TColorBox) SetLeft(value int32) {
    ColorBox_SetLeft(c.instance, value)
}

func (c *TColorBox) Top() int32 {
    return ColorBox_GetTop(c.instance)
}

func (c *TColorBox) SetTop(value int32) {
    ColorBox_SetTop(c.instance, value)
}

func (c *TColorBox) Width() int32 {
    return ColorBox_GetWidth(c.instance)
}

func (c *TColorBox) SetWidth(value int32) {
    ColorBox_SetWidth(c.instance, value)
}

func (c *TColorBox) Height() int32 {
    return ColorBox_GetHeight(c.instance)
}

func (c *TColorBox) SetHeight(value int32) {
    ColorBox_SetHeight(c.instance, value)
}

func (c *TColorBox) Cursor() TCursor {
    return ColorBox_GetCursor(c.instance)
}

func (c *TColorBox) SetCursor(value TCursor) {
    ColorBox_SetCursor(c.instance, value)
}

func (c *TColorBox) Hint() string {
    return ColorBox_GetHint(c.instance)
}

func (c *TColorBox) SetHint(value string) {
    ColorBox_SetHint(c.instance, value)
}

func (c *TColorBox) Margins() *TMargins {
    return MarginsFromInst(ColorBox_GetMargins(c.instance))
}

func (c *TColorBox) SetMargins(value *TMargins) {
    ColorBox_SetMargins(c.instance, CheckPtr(value))
}

func (c *TColorBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ColorBox_GetCustomHint(c.instance))
}

func (c *TColorBox) SetCustomHint(value IComponent) {
    ColorBox_SetCustomHint(c.instance, CheckPtr(value))
}

func (c *TColorBox) ComponentCount() int32 {
    return ColorBox_GetComponentCount(c.instance)
}

func (c *TColorBox) ComponentIndex() int32 {
    return ColorBox_GetComponentIndex(c.instance)
}

func (c *TColorBox) SetComponentIndex(value int32) {
    ColorBox_SetComponentIndex(c.instance, value)
}

func (c *TColorBox) Owner() *TComponent {
    return ComponentFromInst(ColorBox_GetOwner(c.instance))
}

func (c *TColorBox) Name() string {
    return ColorBox_GetName(c.instance)
}

func (c *TColorBox) SetName(value string) {
    ColorBox_SetName(c.instance, value)
}

func (c *TColorBox) Tag() int {
    return ColorBox_GetTag(c.instance)
}

func (c *TColorBox) SetTag(value int) {
    ColorBox_SetTag(c.instance, value)
}

func (c *TColorBox) Colors(Index int32) TColor {
    return ColorBox_GetColors(c.instance, Index)
}

func (c *TColorBox) ColorNames(Index int32) string {
    return ColorBox_GetColorNames(c.instance, Index)
}

func (c *TColorBox) Controls(Index int32) *TControl {
    return ControlFromInst(ColorBox_GetControls(c.instance, Index))
}

func (c *TColorBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ColorBox_GetComponents(c.instance, AIndex))
}

