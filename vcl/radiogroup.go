
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

type TRadioGroup struct {
    IControl
    instance uintptr
}

func NewRadioGroup(owner IComponent) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = RadioGroup_Create(CheckPtr(owner))
    return r
}

func RadioGroupFromInst(inst uintptr) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = inst
    return r
}

func RadioGroupFromObj(obj IObject) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = CheckPtr(obj)
    return r
}

func (r *TRadioGroup) Free() {
    if r.instance != 0 {
        RadioGroup_Free(r.instance)
        r.instance = 0
    }
}

func (r *TRadioGroup) Instance() uintptr {
    return r.instance
}

func (r *TRadioGroup) IsValid() bool {
    return r.instance != 0
}

func (r *TRadioGroup) FlipChildren(AllLevels bool) {
    RadioGroup_FlipChildren(r.instance, AllLevels)
}

func (r *TRadioGroup) CanFocus() bool {
    return RadioGroup_CanFocus(r.instance)
}

func (r *TRadioGroup) Focused() bool {
    return RadioGroup_Focused(r.instance)
}

func (r *TRadioGroup) HandleAllocated() bool {
    return RadioGroup_HandleAllocated(r.instance)
}

func (r *TRadioGroup) Invalidate() {
    RadioGroup_Invalidate(r.instance)
}

func (r *TRadioGroup) Realign() {
    RadioGroup_Realign(r.instance)
}

func (r *TRadioGroup) Repaint() {
    RadioGroup_Repaint(r.instance)
}

func (r *TRadioGroup) ScaleBy(M int32, D int32) {
    RadioGroup_ScaleBy(r.instance, M , D)
}

func (r *TRadioGroup) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RadioGroup_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

func (r *TRadioGroup) SetFocus() {
    RadioGroup_SetFocus(r.instance)
}

func (r *TRadioGroup) Update() {
    RadioGroup_Update(r.instance)
}

func (r *TRadioGroup) BringToFront() {
    RadioGroup_BringToFront(r.instance)
}

func (r *TRadioGroup) HasParent() bool {
    return RadioGroup_HasParent(r.instance)
}

func (r *TRadioGroup) Hide() {
    RadioGroup_Hide(r.instance)
}

func (r *TRadioGroup) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RadioGroup_Perform(r.instance, Msg , WParam , LParam)
}

func (r *TRadioGroup) Refresh() {
    RadioGroup_Refresh(r.instance)
}

func (r *TRadioGroup) SendToBack() {
    RadioGroup_SendToBack(r.instance)
}

func (r *TRadioGroup) Show() {
    RadioGroup_Show(r.instance)
}

func (r *TRadioGroup) GetTextBuf(Buffer string, BufSize int32) int32 {
    return RadioGroup_GetTextBuf(r.instance, Buffer , BufSize)
}

func (r *TRadioGroup) GetTextLen() int32 {
    return RadioGroup_GetTextLen(r.instance)
}

func (r *TRadioGroup) FindComponent(AName string) *TComponent {
    return ComponentFromInst(RadioGroup_FindComponent(r.instance, AName))
}

func (r *TRadioGroup) GetNamePath() string {
    return RadioGroup_GetNamePath(r.instance)
}

func (r *TRadioGroup) Assign(Source IObject) {
    RadioGroup_Assign(r.instance, CheckPtr(Source))
}

func (r *TRadioGroup) ClassName() string {
    return RadioGroup_ClassName(r.instance)
}

func (r *TRadioGroup) Equals(Obj IObject) bool {
    return RadioGroup_Equals(r.instance, CheckPtr(Obj))
}

func (r *TRadioGroup) GetHashCode() int32 {
    return RadioGroup_GetHashCode(r.instance)
}

func (r *TRadioGroup) ToString() string {
    return RadioGroup_ToString(r.instance)
}

func (r *TRadioGroup) Align() TAlign {
    return RadioGroup_GetAlign(r.instance)
}

func (r *TRadioGroup) SetAlign(value TAlign) {
    RadioGroup_SetAlign(r.instance, value)
}

func (r *TRadioGroup) Anchors() TAnchors {
    return RadioGroup_GetAnchors(r.instance)
}

func (r *TRadioGroup) SetAnchors(value TAnchors) {
    RadioGroup_SetAnchors(r.instance, value)
}

func (r *TRadioGroup) BiDiMode() TBiDiMode {
    return RadioGroup_GetBiDiMode(r.instance)
}

func (r *TRadioGroup) SetBiDiMode(value TBiDiMode) {
    RadioGroup_SetBiDiMode(r.instance, value)
}

func (r *TRadioGroup) Caption() string {
    return RadioGroup_GetCaption(r.instance)
}

func (r *TRadioGroup) SetCaption(value string) {
    RadioGroup_SetCaption(r.instance, value)
}

func (r *TRadioGroup) Color() TColor {
    return RadioGroup_GetColor(r.instance)
}

func (r *TRadioGroup) SetColor(value TColor) {
    RadioGroup_SetColor(r.instance, value)
}

func (r *TRadioGroup) Columns() int32 {
    return RadioGroup_GetColumns(r.instance)
}

func (r *TRadioGroup) SetColumns(value int32) {
    RadioGroup_SetColumns(r.instance, value)
}

func (r *TRadioGroup) DoubleBuffered() bool {
    return RadioGroup_GetDoubleBuffered(r.instance)
}

func (r *TRadioGroup) SetDoubleBuffered(value bool) {
    RadioGroup_SetDoubleBuffered(r.instance, value)
}

func (r *TRadioGroup) Enabled() bool {
    return RadioGroup_GetEnabled(r.instance)
}

func (r *TRadioGroup) SetEnabled(value bool) {
    RadioGroup_SetEnabled(r.instance, value)
}

func (r *TRadioGroup) Font() *TFont {
    return FontFromInst(RadioGroup_GetFont(r.instance))
}

func (r *TRadioGroup) SetFont(value *TFont) {
    RadioGroup_SetFont(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) ItemIndex() int32 {
    return RadioGroup_GetItemIndex(r.instance)
}

func (r *TRadioGroup) SetItemIndex(value int32) {
    RadioGroup_SetItemIndex(r.instance, value)
}

func (r *TRadioGroup) Items() *TStrings {
    return StringsFromInst(RadioGroup_GetItems(r.instance))
}

func (r *TRadioGroup) SetItems(value IObject) {
    RadioGroup_SetItems(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) ParentBackground() bool {
    return RadioGroup_GetParentBackground(r.instance)
}

func (r *TRadioGroup) SetParentBackground(value bool) {
    RadioGroup_SetParentBackground(r.instance, value)
}

func (r *TRadioGroup) ParentColor() bool {
    return RadioGroup_GetParentColor(r.instance)
}

func (r *TRadioGroup) SetParentColor(value bool) {
    RadioGroup_SetParentColor(r.instance, value)
}

func (r *TRadioGroup) ParentCtl3D() bool {
    return RadioGroup_GetParentCtl3D(r.instance)
}

func (r *TRadioGroup) SetParentCtl3D(value bool) {
    RadioGroup_SetParentCtl3D(r.instance, value)
}

func (r *TRadioGroup) ParentDoubleBuffered() bool {
    return RadioGroup_GetParentDoubleBuffered(r.instance)
}

func (r *TRadioGroup) SetParentDoubleBuffered(value bool) {
    RadioGroup_SetParentDoubleBuffered(r.instance, value)
}

func (r *TRadioGroup) ParentFont() bool {
    return RadioGroup_GetParentFont(r.instance)
}

func (r *TRadioGroup) SetParentFont(value bool) {
    RadioGroup_SetParentFont(r.instance, value)
}

func (r *TRadioGroup) ParentShowHint() bool {
    return RadioGroup_GetParentShowHint(r.instance)
}

func (r *TRadioGroup) SetParentShowHint(value bool) {
    RadioGroup_SetParentShowHint(r.instance, value)
}

func (r *TRadioGroup) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(RadioGroup_GetPopupMenu(r.instance))
}

func (r *TRadioGroup) SetPopupMenu(value IComponent) {
    RadioGroup_SetPopupMenu(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) ShowHint() bool {
    return RadioGroup_GetShowHint(r.instance)
}

func (r *TRadioGroup) SetShowHint(value bool) {
    RadioGroup_SetShowHint(r.instance, value)
}

func (r *TRadioGroup) TabOrder() uint16 {
    return RadioGroup_GetTabOrder(r.instance)
}

func (r *TRadioGroup) SetTabOrder(value uint16) {
    RadioGroup_SetTabOrder(r.instance, value)
}

func (r *TRadioGroup) TabStop() bool {
    return RadioGroup_GetTabStop(r.instance)
}

func (r *TRadioGroup) SetTabStop(value bool) {
    RadioGroup_SetTabStop(r.instance, value)
}

func (r *TRadioGroup) Visible() bool {
    return RadioGroup_GetVisible(r.instance)
}

func (r *TRadioGroup) SetVisible(value bool) {
    RadioGroup_SetVisible(r.instance, value)
}

func (r *TRadioGroup) StyleElements() TStyleElements {
    return RadioGroup_GetStyleElements(r.instance)
}

func (r *TRadioGroup) SetStyleElements(value TStyleElements) {
    RadioGroup_SetStyleElements(r.instance, value)
}

func (r *TRadioGroup) WordWrap() bool {
    return RadioGroup_GetWordWrap(r.instance)
}

func (r *TRadioGroup) SetWordWrap(value bool) {
    RadioGroup_SetWordWrap(r.instance, value)
}

func (r *TRadioGroup) SetOnClick(fn TNotifyEvent) {
    RadioGroup_SetOnClick(r.instance, fn)
}

func (r *TRadioGroup) SetOnEnter(fn TNotifyEvent) {
    RadioGroup_SetOnEnter(r.instance, fn)
}

func (r *TRadioGroup) SetOnExit(fn TNotifyEvent) {
    RadioGroup_SetOnExit(r.instance, fn)
}

func (r *TRadioGroup) Brush() *TBrush {
    return BrushFromInst(RadioGroup_GetBrush(r.instance))
}

func (r *TRadioGroup) ControlCount() int32 {
    return RadioGroup_GetControlCount(r.instance)
}

func (r *TRadioGroup) Handle() HWND {
    return RadioGroup_GetHandle(r.instance)
}

func (r *TRadioGroup) ParentWindow() HWND {
    return RadioGroup_GetParentWindow(r.instance)
}

func (r *TRadioGroup) SetParentWindow(value HWND) {
    RadioGroup_SetParentWindow(r.instance, value)
}

func (r *TRadioGroup) Action() *TAction {
    return ActionFromInst(RadioGroup_GetAction(r.instance))
}

func (r *TRadioGroup) SetAction(value IComponent) {
    RadioGroup_SetAction(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) BoundsRect() TRect {
    return RadioGroup_GetBoundsRect(r.instance)
}

func (r *TRadioGroup) SetBoundsRect(value TRect) {
    RadioGroup_SetBoundsRect(r.instance, value)
}

func (r *TRadioGroup) ClientHeight() int32 {
    return RadioGroup_GetClientHeight(r.instance)
}

func (r *TRadioGroup) SetClientHeight(value int32) {
    RadioGroup_SetClientHeight(r.instance, value)
}

func (r *TRadioGroup) ClientRect() TRect {
    return RadioGroup_GetClientRect(r.instance)
}

func (r *TRadioGroup) ClientWidth() int32 {
    return RadioGroup_GetClientWidth(r.instance)
}

func (r *TRadioGroup) SetClientWidth(value int32) {
    RadioGroup_SetClientWidth(r.instance, value)
}

func (r *TRadioGroup) ExplicitLeft() int32 {
    return RadioGroup_GetExplicitLeft(r.instance)
}

func (r *TRadioGroup) ExplicitTop() int32 {
    return RadioGroup_GetExplicitTop(r.instance)
}

func (r *TRadioGroup) ExplicitWidth() int32 {
    return RadioGroup_GetExplicitWidth(r.instance)
}

func (r *TRadioGroup) ExplicitHeight() int32 {
    return RadioGroup_GetExplicitHeight(r.instance)
}

func (r *TRadioGroup) Parent() *TControl {
    return ControlFromInst(RadioGroup_GetParent(r.instance))
}

func (r *TRadioGroup) SetParent(value IControl) {
    RadioGroup_SetParent(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) AlignWithMargins() bool {
    return RadioGroup_GetAlignWithMargins(r.instance)
}

func (r *TRadioGroup) SetAlignWithMargins(value bool) {
    RadioGroup_SetAlignWithMargins(r.instance, value)
}

func (r *TRadioGroup) Left() int32 {
    return RadioGroup_GetLeft(r.instance)
}

func (r *TRadioGroup) SetLeft(value int32) {
    RadioGroup_SetLeft(r.instance, value)
}

func (r *TRadioGroup) Top() int32 {
    return RadioGroup_GetTop(r.instance)
}

func (r *TRadioGroup) SetTop(value int32) {
    RadioGroup_SetTop(r.instance, value)
}

func (r *TRadioGroup) Width() int32 {
    return RadioGroup_GetWidth(r.instance)
}

func (r *TRadioGroup) SetWidth(value int32) {
    RadioGroup_SetWidth(r.instance, value)
}

func (r *TRadioGroup) Height() int32 {
    return RadioGroup_GetHeight(r.instance)
}

func (r *TRadioGroup) SetHeight(value int32) {
    RadioGroup_SetHeight(r.instance, value)
}

func (r *TRadioGroup) Cursor() TCursor {
    return RadioGroup_GetCursor(r.instance)
}

func (r *TRadioGroup) SetCursor(value TCursor) {
    RadioGroup_SetCursor(r.instance, value)
}

func (r *TRadioGroup) Hint() string {
    return RadioGroup_GetHint(r.instance)
}

func (r *TRadioGroup) SetHint(value string) {
    RadioGroup_SetHint(r.instance, value)
}

func (r *TRadioGroup) Margins() *TMargins {
    return MarginsFromInst(RadioGroup_GetMargins(r.instance))
}

func (r *TRadioGroup) SetMargins(value *TMargins) {
    RadioGroup_SetMargins(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) CustomHint() *TCustomHint {
    return CustomHintFromInst(RadioGroup_GetCustomHint(r.instance))
}

func (r *TRadioGroup) SetCustomHint(value IComponent) {
    RadioGroup_SetCustomHint(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) ComponentCount() int32 {
    return RadioGroup_GetComponentCount(r.instance)
}

func (r *TRadioGroup) ComponentIndex() int32 {
    return RadioGroup_GetComponentIndex(r.instance)
}

func (r *TRadioGroup) SetComponentIndex(value int32) {
    RadioGroup_SetComponentIndex(r.instance, value)
}

func (r *TRadioGroup) Owner() *TComponent {
    return ComponentFromInst(RadioGroup_GetOwner(r.instance))
}

func (r *TRadioGroup) Name() string {
    return RadioGroup_GetName(r.instance)
}

func (r *TRadioGroup) SetName(value string) {
    RadioGroup_SetName(r.instance, value)
}

func (r *TRadioGroup) Tag() int {
    return RadioGroup_GetTag(r.instance)
}

func (r *TRadioGroup) SetTag(value int) {
    RadioGroup_SetTag(r.instance, value)
}

func (r *TRadioGroup) Buttons(Index int32) *TRadioButton {
    return RadioButtonFromInst(RadioGroup_GetButtons(r.instance, Index))
}

func (r *TRadioGroup) Controls(Index int32) *TControl {
    return ControlFromInst(RadioGroup_GetControls(r.instance, Index))
}

func (r *TRadioGroup) Components(AIndex int32) *TComponent {
    return ComponentFromInst(RadioGroup_GetComponents(r.instance, AIndex))
}

