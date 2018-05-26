
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

type TRadioButton struct {
    IControl
    instance uintptr
}

func NewRadioButton(owner IComponent) *TRadioButton {
    r := new(TRadioButton)
    r.instance = RadioButton_Create(CheckPtr(owner))
    return r
}

func RadioButtonFromInst(inst uintptr) *TRadioButton {
    r := new(TRadioButton)
    r.instance = inst
    return r
}

func RadioButtonFromObj(obj IObject) *TRadioButton {
    r := new(TRadioButton)
    r.instance = CheckPtr(obj)
    return r
}

func (r *TRadioButton) Free() {
    if r.instance != 0 {
        RadioButton_Free(r.instance)
        r.instance = 0
    }
}

func (r *TRadioButton) Instance() uintptr {
    return r.instance
}

func (r *TRadioButton) IsValid() bool {
    return r.instance != 0
}

func (r *TRadioButton) CanFocus() bool {
    return RadioButton_CanFocus(r.instance)
}

func (r *TRadioButton) FlipChildren(AllLevels bool) {
    RadioButton_FlipChildren(r.instance, AllLevels)
}

func (r *TRadioButton) Focused() bool {
    return RadioButton_Focused(r.instance)
}

func (r *TRadioButton) HandleAllocated() bool {
    return RadioButton_HandleAllocated(r.instance)
}

func (r *TRadioButton) Invalidate() {
    RadioButton_Invalidate(r.instance)
}

func (r *TRadioButton) Realign() {
    RadioButton_Realign(r.instance)
}

func (r *TRadioButton) Repaint() {
    RadioButton_Repaint(r.instance)
}

func (r *TRadioButton) ScaleBy(M int32, D int32) {
    RadioButton_ScaleBy(r.instance, M , D)
}

func (r *TRadioButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RadioButton_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

func (r *TRadioButton) SetFocus() {
    RadioButton_SetFocus(r.instance)
}

func (r *TRadioButton) Update() {
    RadioButton_Update(r.instance)
}

func (r *TRadioButton) BringToFront() {
    RadioButton_BringToFront(r.instance)
}

func (r *TRadioButton) HasParent() bool {
    return RadioButton_HasParent(r.instance)
}

func (r *TRadioButton) Hide() {
    RadioButton_Hide(r.instance)
}

func (r *TRadioButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RadioButton_Perform(r.instance, Msg , WParam , LParam)
}

func (r *TRadioButton) Refresh() {
    RadioButton_Refresh(r.instance)
}

func (r *TRadioButton) SendToBack() {
    RadioButton_SendToBack(r.instance)
}

func (r *TRadioButton) Show() {
    RadioButton_Show(r.instance)
}

func (r *TRadioButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return RadioButton_GetTextBuf(r.instance, Buffer , BufSize)
}

func (r *TRadioButton) GetTextLen() int32 {
    return RadioButton_GetTextLen(r.instance)
}

func (r *TRadioButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(RadioButton_FindComponent(r.instance, AName))
}

func (r *TRadioButton) GetNamePath() string {
    return RadioButton_GetNamePath(r.instance)
}

func (r *TRadioButton) Assign(Source IObject) {
    RadioButton_Assign(r.instance, CheckPtr(Source))
}

func (r *TRadioButton) ClassName() string {
    return RadioButton_ClassName(r.instance)
}

func (r *TRadioButton) Equals(Obj IObject) bool {
    return RadioButton_Equals(r.instance, CheckPtr(Obj))
}

func (r *TRadioButton) GetHashCode() int32 {
    return RadioButton_GetHashCode(r.instance)
}

func (r *TRadioButton) ToString() string {
    return RadioButton_ToString(r.instance)
}

func (r *TRadioButton) Action() *TAction {
    return ActionFromInst(RadioButton_GetAction(r.instance))
}

func (r *TRadioButton) SetAction(value IComponent) {
    RadioButton_SetAction(r.instance, CheckPtr(value))
}

func (r *TRadioButton) Align() TAlign {
    return RadioButton_GetAlign(r.instance)
}

func (r *TRadioButton) SetAlign(value TAlign) {
    RadioButton_SetAlign(r.instance, value)
}

func (r *TRadioButton) Alignment() TLeftRight {
    return RadioButton_GetAlignment(r.instance)
}

func (r *TRadioButton) SetAlignment(value TLeftRight) {
    RadioButton_SetAlignment(r.instance, value)
}

func (r *TRadioButton) Anchors() TAnchors {
    return RadioButton_GetAnchors(r.instance)
}

func (r *TRadioButton) SetAnchors(value TAnchors) {
    RadioButton_SetAnchors(r.instance, value)
}

func (r *TRadioButton) BiDiMode() TBiDiMode {
    return RadioButton_GetBiDiMode(r.instance)
}

func (r *TRadioButton) SetBiDiMode(value TBiDiMode) {
    RadioButton_SetBiDiMode(r.instance, value)
}

func (r *TRadioButton) Caption() string {
    return RadioButton_GetCaption(r.instance)
}

func (r *TRadioButton) SetCaption(value string) {
    RadioButton_SetCaption(r.instance, value)
}

func (r *TRadioButton) Checked() bool {
    return RadioButton_GetChecked(r.instance)
}

func (r *TRadioButton) SetChecked(value bool) {
    RadioButton_SetChecked(r.instance, value)
}

func (r *TRadioButton) Color() TColor {
    return RadioButton_GetColor(r.instance)
}

func (r *TRadioButton) SetColor(value TColor) {
    RadioButton_SetColor(r.instance, value)
}

func (r *TRadioButton) DoubleBuffered() bool {
    return RadioButton_GetDoubleBuffered(r.instance)
}

func (r *TRadioButton) SetDoubleBuffered(value bool) {
    RadioButton_SetDoubleBuffered(r.instance, value)
}

func (r *TRadioButton) Enabled() bool {
    return RadioButton_GetEnabled(r.instance)
}

func (r *TRadioButton) SetEnabled(value bool) {
    RadioButton_SetEnabled(r.instance, value)
}

func (r *TRadioButton) Font() *TFont {
    return FontFromInst(RadioButton_GetFont(r.instance))
}

func (r *TRadioButton) SetFont(value *TFont) {
    RadioButton_SetFont(r.instance, CheckPtr(value))
}

func (r *TRadioButton) ParentColor() bool {
    return RadioButton_GetParentColor(r.instance)
}

func (r *TRadioButton) SetParentColor(value bool) {
    RadioButton_SetParentColor(r.instance, value)
}

func (r *TRadioButton) ParentCtl3D() bool {
    return RadioButton_GetParentCtl3D(r.instance)
}

func (r *TRadioButton) SetParentCtl3D(value bool) {
    RadioButton_SetParentCtl3D(r.instance, value)
}

func (r *TRadioButton) ParentDoubleBuffered() bool {
    return RadioButton_GetParentDoubleBuffered(r.instance)
}

func (r *TRadioButton) SetParentDoubleBuffered(value bool) {
    RadioButton_SetParentDoubleBuffered(r.instance, value)
}

func (r *TRadioButton) ParentFont() bool {
    return RadioButton_GetParentFont(r.instance)
}

func (r *TRadioButton) SetParentFont(value bool) {
    RadioButton_SetParentFont(r.instance, value)
}

func (r *TRadioButton) ParentShowHint() bool {
    return RadioButton_GetParentShowHint(r.instance)
}

func (r *TRadioButton) SetParentShowHint(value bool) {
    RadioButton_SetParentShowHint(r.instance, value)
}

func (r *TRadioButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(RadioButton_GetPopupMenu(r.instance))
}

func (r *TRadioButton) SetPopupMenu(value IComponent) {
    RadioButton_SetPopupMenu(r.instance, CheckPtr(value))
}

func (r *TRadioButton) ShowHint() bool {
    return RadioButton_GetShowHint(r.instance)
}

func (r *TRadioButton) SetShowHint(value bool) {
    RadioButton_SetShowHint(r.instance, value)
}

func (r *TRadioButton) TabOrder() uint16 {
    return RadioButton_GetTabOrder(r.instance)
}

func (r *TRadioButton) SetTabOrder(value uint16) {
    RadioButton_SetTabOrder(r.instance, value)
}

func (r *TRadioButton) TabStop() bool {
    return RadioButton_GetTabStop(r.instance)
}

func (r *TRadioButton) SetTabStop(value bool) {
    RadioButton_SetTabStop(r.instance, value)
}

func (r *TRadioButton) Visible() bool {
    return RadioButton_GetVisible(r.instance)
}

func (r *TRadioButton) SetVisible(value bool) {
    RadioButton_SetVisible(r.instance, value)
}

func (r *TRadioButton) WordWrap() bool {
    return RadioButton_GetWordWrap(r.instance)
}

func (r *TRadioButton) SetWordWrap(value bool) {
    RadioButton_SetWordWrap(r.instance, value)
}

func (r *TRadioButton) StyleElements() TStyleElements {
    return RadioButton_GetStyleElements(r.instance)
}

func (r *TRadioButton) SetStyleElements(value TStyleElements) {
    RadioButton_SetStyleElements(r.instance, value)
}

func (r *TRadioButton) SetOnClick(fn TNotifyEvent) {
    RadioButton_SetOnClick(r.instance, fn)
}

func (r *TRadioButton) SetOnDblClick(fn TNotifyEvent) {
    RadioButton_SetOnDblClick(r.instance, fn)
}

func (r *TRadioButton) SetOnEnter(fn TNotifyEvent) {
    RadioButton_SetOnEnter(r.instance, fn)
}

func (r *TRadioButton) SetOnExit(fn TNotifyEvent) {
    RadioButton_SetOnExit(r.instance, fn)
}

func (r *TRadioButton) SetOnKeyDown(fn TKeyEvent) {
    RadioButton_SetOnKeyDown(r.instance, fn)
}

func (r *TRadioButton) SetOnKeyPress(fn TKeyPressEvent) {
    RadioButton_SetOnKeyPress(r.instance, fn)
}

func (r *TRadioButton) SetOnKeyUp(fn TKeyEvent) {
    RadioButton_SetOnKeyUp(r.instance, fn)
}

func (r *TRadioButton) SetOnMouseDown(fn TMouseEvent) {
    RadioButton_SetOnMouseDown(r.instance, fn)
}

func (r *TRadioButton) SetOnMouseEnter(fn TNotifyEvent) {
    RadioButton_SetOnMouseEnter(r.instance, fn)
}

func (r *TRadioButton) SetOnMouseLeave(fn TNotifyEvent) {
    RadioButton_SetOnMouseLeave(r.instance, fn)
}

func (r *TRadioButton) SetOnMouseMove(fn TMouseMoveEvent) {
    RadioButton_SetOnMouseMove(r.instance, fn)
}

func (r *TRadioButton) SetOnMouseUp(fn TMouseEvent) {
    RadioButton_SetOnMouseUp(r.instance, fn)
}

func (r *TRadioButton) Brush() *TBrush {
    return BrushFromInst(RadioButton_GetBrush(r.instance))
}

func (r *TRadioButton) ControlCount() int32 {
    return RadioButton_GetControlCount(r.instance)
}

func (r *TRadioButton) Handle() HWND {
    return RadioButton_GetHandle(r.instance)
}

func (r *TRadioButton) ParentWindow() HWND {
    return RadioButton_GetParentWindow(r.instance)
}

func (r *TRadioButton) SetParentWindow(value HWND) {
    RadioButton_SetParentWindow(r.instance, value)
}

func (r *TRadioButton) BoundsRect() TRect {
    return RadioButton_GetBoundsRect(r.instance)
}

func (r *TRadioButton) SetBoundsRect(value TRect) {
    RadioButton_SetBoundsRect(r.instance, value)
}

func (r *TRadioButton) ClientHeight() int32 {
    return RadioButton_GetClientHeight(r.instance)
}

func (r *TRadioButton) SetClientHeight(value int32) {
    RadioButton_SetClientHeight(r.instance, value)
}

func (r *TRadioButton) ClientRect() TRect {
    return RadioButton_GetClientRect(r.instance)
}

func (r *TRadioButton) ClientWidth() int32 {
    return RadioButton_GetClientWidth(r.instance)
}

func (r *TRadioButton) SetClientWidth(value int32) {
    RadioButton_SetClientWidth(r.instance, value)
}

func (r *TRadioButton) ExplicitLeft() int32 {
    return RadioButton_GetExplicitLeft(r.instance)
}

func (r *TRadioButton) ExplicitTop() int32 {
    return RadioButton_GetExplicitTop(r.instance)
}

func (r *TRadioButton) ExplicitWidth() int32 {
    return RadioButton_GetExplicitWidth(r.instance)
}

func (r *TRadioButton) ExplicitHeight() int32 {
    return RadioButton_GetExplicitHeight(r.instance)
}

func (r *TRadioButton) Parent() *TControl {
    return ControlFromInst(RadioButton_GetParent(r.instance))
}

func (r *TRadioButton) SetParent(value IControl) {
    RadioButton_SetParent(r.instance, CheckPtr(value))
}

func (r *TRadioButton) AlignWithMargins() bool {
    return RadioButton_GetAlignWithMargins(r.instance)
}

func (r *TRadioButton) SetAlignWithMargins(value bool) {
    RadioButton_SetAlignWithMargins(r.instance, value)
}

func (r *TRadioButton) Left() int32 {
    return RadioButton_GetLeft(r.instance)
}

func (r *TRadioButton) SetLeft(value int32) {
    RadioButton_SetLeft(r.instance, value)
}

func (r *TRadioButton) Top() int32 {
    return RadioButton_GetTop(r.instance)
}

func (r *TRadioButton) SetTop(value int32) {
    RadioButton_SetTop(r.instance, value)
}

func (r *TRadioButton) Width() int32 {
    return RadioButton_GetWidth(r.instance)
}

func (r *TRadioButton) SetWidth(value int32) {
    RadioButton_SetWidth(r.instance, value)
}

func (r *TRadioButton) Height() int32 {
    return RadioButton_GetHeight(r.instance)
}

func (r *TRadioButton) SetHeight(value int32) {
    RadioButton_SetHeight(r.instance, value)
}

func (r *TRadioButton) Cursor() TCursor {
    return RadioButton_GetCursor(r.instance)
}

func (r *TRadioButton) SetCursor(value TCursor) {
    RadioButton_SetCursor(r.instance, value)
}

func (r *TRadioButton) Hint() string {
    return RadioButton_GetHint(r.instance)
}

func (r *TRadioButton) SetHint(value string) {
    RadioButton_SetHint(r.instance, value)
}

func (r *TRadioButton) Margins() *TMargins {
    return MarginsFromInst(RadioButton_GetMargins(r.instance))
}

func (r *TRadioButton) SetMargins(value *TMargins) {
    RadioButton_SetMargins(r.instance, CheckPtr(value))
}

func (r *TRadioButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(RadioButton_GetCustomHint(r.instance))
}

func (r *TRadioButton) SetCustomHint(value IComponent) {
    RadioButton_SetCustomHint(r.instance, CheckPtr(value))
}

func (r *TRadioButton) ComponentCount() int32 {
    return RadioButton_GetComponentCount(r.instance)
}

func (r *TRadioButton) ComponentIndex() int32 {
    return RadioButton_GetComponentIndex(r.instance)
}

func (r *TRadioButton) SetComponentIndex(value int32) {
    RadioButton_SetComponentIndex(r.instance, value)
}

func (r *TRadioButton) Owner() *TComponent {
    return ComponentFromInst(RadioButton_GetOwner(r.instance))
}

func (r *TRadioButton) Name() string {
    return RadioButton_GetName(r.instance)
}

func (r *TRadioButton) SetName(value string) {
    RadioButton_SetName(r.instance, value)
}

func (r *TRadioButton) Tag() int {
    return RadioButton_GetTag(r.instance)
}

func (r *TRadioButton) SetTag(value int) {
    RadioButton_SetTag(r.instance, value)
}

func (r *TRadioButton) Controls(Index int32) *TControl {
    return ControlFromInst(RadioButton_GetControls(r.instance, Index))
}

func (r *TRadioButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(RadioButton_GetComponents(r.instance, AIndex))
}

