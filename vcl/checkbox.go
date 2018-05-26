
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

type TCheckBox struct {
    IControl
    instance uintptr
}

func NewCheckBox(owner IComponent) *TCheckBox {
    c := new(TCheckBox)
    c.instance = CheckBox_Create(CheckPtr(owner))
    return c
}

func CheckBoxFromInst(inst uintptr) *TCheckBox {
    c := new(TCheckBox)
    c.instance = inst
    return c
}

func CheckBoxFromObj(obj IObject) *TCheckBox {
    c := new(TCheckBox)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCheckBox) Free() {
    if c.instance != 0 {
        CheckBox_Free(c.instance)
        c.instance = 0
    }
}

func (c *TCheckBox) Instance() uintptr {
    return c.instance
}

func (c *TCheckBox) IsValid() bool {
    return c.instance != 0
}

func (c *TCheckBox) CanFocus() bool {
    return CheckBox_CanFocus(c.instance)
}

func (c *TCheckBox) FlipChildren(AllLevels bool) {
    CheckBox_FlipChildren(c.instance, AllLevels)
}

func (c *TCheckBox) Focused() bool {
    return CheckBox_Focused(c.instance)
}

func (c *TCheckBox) HandleAllocated() bool {
    return CheckBox_HandleAllocated(c.instance)
}

func (c *TCheckBox) Invalidate() {
    CheckBox_Invalidate(c.instance)
}

func (c *TCheckBox) Realign() {
    CheckBox_Realign(c.instance)
}

func (c *TCheckBox) Repaint() {
    CheckBox_Repaint(c.instance)
}

func (c *TCheckBox) ScaleBy(M int32, D int32) {
    CheckBox_ScaleBy(c.instance, M , D)
}

func (c *TCheckBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CheckBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

func (c *TCheckBox) SetFocus() {
    CheckBox_SetFocus(c.instance)
}

func (c *TCheckBox) Update() {
    CheckBox_Update(c.instance)
}

func (c *TCheckBox) BringToFront() {
    CheckBox_BringToFront(c.instance)
}

func (c *TCheckBox) HasParent() bool {
    return CheckBox_HasParent(c.instance)
}

func (c *TCheckBox) Hide() {
    CheckBox_Hide(c.instance)
}

func (c *TCheckBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CheckBox_Perform(c.instance, Msg , WParam , LParam)
}

func (c *TCheckBox) Refresh() {
    CheckBox_Refresh(c.instance)
}

func (c *TCheckBox) SendToBack() {
    CheckBox_SendToBack(c.instance)
}

func (c *TCheckBox) Show() {
    CheckBox_Show(c.instance)
}

func (c *TCheckBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CheckBox_GetTextBuf(c.instance, Buffer , BufSize)
}

func (c *TCheckBox) GetTextLen() int32 {
    return CheckBox_GetTextLen(c.instance)
}

func (c *TCheckBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CheckBox_FindComponent(c.instance, AName))
}

func (c *TCheckBox) GetNamePath() string {
    return CheckBox_GetNamePath(c.instance)
}

func (c *TCheckBox) Assign(Source IObject) {
    CheckBox_Assign(c.instance, CheckPtr(Source))
}

func (c *TCheckBox) ClassName() string {
    return CheckBox_ClassName(c.instance)
}

func (c *TCheckBox) Equals(Obj IObject) bool {
    return CheckBox_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCheckBox) GetHashCode() int32 {
    return CheckBox_GetHashCode(c.instance)
}

func (c *TCheckBox) ToString() string {
    return CheckBox_ToString(c.instance)
}

func (c *TCheckBox) Action() *TAction {
    return ActionFromInst(CheckBox_GetAction(c.instance))
}

func (c *TCheckBox) SetAction(value IComponent) {
    CheckBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TCheckBox) Align() TAlign {
    return CheckBox_GetAlign(c.instance)
}

func (c *TCheckBox) SetAlign(value TAlign) {
    CheckBox_SetAlign(c.instance, value)
}

func (c *TCheckBox) Alignment() TLeftRight {
    return CheckBox_GetAlignment(c.instance)
}

func (c *TCheckBox) SetAlignment(value TLeftRight) {
    CheckBox_SetAlignment(c.instance, value)
}

func (c *TCheckBox) AllowGrayed() bool {
    return CheckBox_GetAllowGrayed(c.instance)
}

func (c *TCheckBox) SetAllowGrayed(value bool) {
    CheckBox_SetAllowGrayed(c.instance, value)
}

func (c *TCheckBox) Anchors() TAnchors {
    return CheckBox_GetAnchors(c.instance)
}

func (c *TCheckBox) SetAnchors(value TAnchors) {
    CheckBox_SetAnchors(c.instance, value)
}

func (c *TCheckBox) BiDiMode() TBiDiMode {
    return CheckBox_GetBiDiMode(c.instance)
}

func (c *TCheckBox) SetBiDiMode(value TBiDiMode) {
    CheckBox_SetBiDiMode(c.instance, value)
}

func (c *TCheckBox) Caption() string {
    return CheckBox_GetCaption(c.instance)
}

func (c *TCheckBox) SetCaption(value string) {
    CheckBox_SetCaption(c.instance, value)
}

func (c *TCheckBox) Checked() bool {
    return CheckBox_GetChecked(c.instance)
}

func (c *TCheckBox) SetChecked(value bool) {
    CheckBox_SetChecked(c.instance, value)
}

func (c *TCheckBox) Color() TColor {
    return CheckBox_GetColor(c.instance)
}

func (c *TCheckBox) SetColor(value TColor) {
    CheckBox_SetColor(c.instance, value)
}

func (c *TCheckBox) DoubleBuffered() bool {
    return CheckBox_GetDoubleBuffered(c.instance)
}

func (c *TCheckBox) SetDoubleBuffered(value bool) {
    CheckBox_SetDoubleBuffered(c.instance, value)
}

func (c *TCheckBox) Enabled() bool {
    return CheckBox_GetEnabled(c.instance)
}

func (c *TCheckBox) SetEnabled(value bool) {
    CheckBox_SetEnabled(c.instance, value)
}

func (c *TCheckBox) Font() *TFont {
    return FontFromInst(CheckBox_GetFont(c.instance))
}

func (c *TCheckBox) SetFont(value *TFont) {
    CheckBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TCheckBox) ParentColor() bool {
    return CheckBox_GetParentColor(c.instance)
}

func (c *TCheckBox) SetParentColor(value bool) {
    CheckBox_SetParentColor(c.instance, value)
}

func (c *TCheckBox) ParentCtl3D() bool {
    return CheckBox_GetParentCtl3D(c.instance)
}

func (c *TCheckBox) SetParentCtl3D(value bool) {
    CheckBox_SetParentCtl3D(c.instance, value)
}

func (c *TCheckBox) ParentDoubleBuffered() bool {
    return CheckBox_GetParentDoubleBuffered(c.instance)
}

func (c *TCheckBox) SetParentDoubleBuffered(value bool) {
    CheckBox_SetParentDoubleBuffered(c.instance, value)
}

func (c *TCheckBox) ParentFont() bool {
    return CheckBox_GetParentFont(c.instance)
}

func (c *TCheckBox) SetParentFont(value bool) {
    CheckBox_SetParentFont(c.instance, value)
}

func (c *TCheckBox) ParentShowHint() bool {
    return CheckBox_GetParentShowHint(c.instance)
}

func (c *TCheckBox) SetParentShowHint(value bool) {
    CheckBox_SetParentShowHint(c.instance, value)
}

func (c *TCheckBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CheckBox_GetPopupMenu(c.instance))
}

func (c *TCheckBox) SetPopupMenu(value IComponent) {
    CheckBox_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TCheckBox) ShowHint() bool {
    return CheckBox_GetShowHint(c.instance)
}

func (c *TCheckBox) SetShowHint(value bool) {
    CheckBox_SetShowHint(c.instance, value)
}

func (c *TCheckBox) State() TCheckBoxState {
    return CheckBox_GetState(c.instance)
}

func (c *TCheckBox) SetState(value TCheckBoxState) {
    CheckBox_SetState(c.instance, value)
}

func (c *TCheckBox) TabOrder() uint16 {
    return CheckBox_GetTabOrder(c.instance)
}

func (c *TCheckBox) SetTabOrder(value uint16) {
    CheckBox_SetTabOrder(c.instance, value)
}

func (c *TCheckBox) TabStop() bool {
    return CheckBox_GetTabStop(c.instance)
}

func (c *TCheckBox) SetTabStop(value bool) {
    CheckBox_SetTabStop(c.instance, value)
}

func (c *TCheckBox) Visible() bool {
    return CheckBox_GetVisible(c.instance)
}

func (c *TCheckBox) SetVisible(value bool) {
    CheckBox_SetVisible(c.instance, value)
}

func (c *TCheckBox) WordWrap() bool {
    return CheckBox_GetWordWrap(c.instance)
}

func (c *TCheckBox) SetWordWrap(value bool) {
    CheckBox_SetWordWrap(c.instance, value)
}

func (c *TCheckBox) StyleElements() TStyleElements {
    return CheckBox_GetStyleElements(c.instance)
}

func (c *TCheckBox) SetStyleElements(value TStyleElements) {
    CheckBox_SetStyleElements(c.instance, value)
}

func (c *TCheckBox) SetOnClick(fn TNotifyEvent) {
    CheckBox_SetOnClick(c.instance, fn)
}

func (c *TCheckBox) SetOnEnter(fn TNotifyEvent) {
    CheckBox_SetOnEnter(c.instance, fn)
}

func (c *TCheckBox) SetOnExit(fn TNotifyEvent) {
    CheckBox_SetOnExit(c.instance, fn)
}

func (c *TCheckBox) SetOnKeyDown(fn TKeyEvent) {
    CheckBox_SetOnKeyDown(c.instance, fn)
}

func (c *TCheckBox) SetOnKeyPress(fn TKeyPressEvent) {
    CheckBox_SetOnKeyPress(c.instance, fn)
}

func (c *TCheckBox) SetOnKeyUp(fn TKeyEvent) {
    CheckBox_SetOnKeyUp(c.instance, fn)
}

func (c *TCheckBox) SetOnMouseDown(fn TMouseEvent) {
    CheckBox_SetOnMouseDown(c.instance, fn)
}

func (c *TCheckBox) SetOnMouseEnter(fn TNotifyEvent) {
    CheckBox_SetOnMouseEnter(c.instance, fn)
}

func (c *TCheckBox) SetOnMouseLeave(fn TNotifyEvent) {
    CheckBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TCheckBox) SetOnMouseMove(fn TMouseMoveEvent) {
    CheckBox_SetOnMouseMove(c.instance, fn)
}

func (c *TCheckBox) SetOnMouseUp(fn TMouseEvent) {
    CheckBox_SetOnMouseUp(c.instance, fn)
}

func (c *TCheckBox) Brush() *TBrush {
    return BrushFromInst(CheckBox_GetBrush(c.instance))
}

func (c *TCheckBox) ControlCount() int32 {
    return CheckBox_GetControlCount(c.instance)
}

func (c *TCheckBox) Handle() HWND {
    return CheckBox_GetHandle(c.instance)
}

func (c *TCheckBox) ParentWindow() HWND {
    return CheckBox_GetParentWindow(c.instance)
}

func (c *TCheckBox) SetParentWindow(value HWND) {
    CheckBox_SetParentWindow(c.instance, value)
}

func (c *TCheckBox) BoundsRect() TRect {
    return CheckBox_GetBoundsRect(c.instance)
}

func (c *TCheckBox) SetBoundsRect(value TRect) {
    CheckBox_SetBoundsRect(c.instance, value)
}

func (c *TCheckBox) ClientHeight() int32 {
    return CheckBox_GetClientHeight(c.instance)
}

func (c *TCheckBox) SetClientHeight(value int32) {
    CheckBox_SetClientHeight(c.instance, value)
}

func (c *TCheckBox) ClientRect() TRect {
    return CheckBox_GetClientRect(c.instance)
}

func (c *TCheckBox) ClientWidth() int32 {
    return CheckBox_GetClientWidth(c.instance)
}

func (c *TCheckBox) SetClientWidth(value int32) {
    CheckBox_SetClientWidth(c.instance, value)
}

func (c *TCheckBox) ExplicitLeft() int32 {
    return CheckBox_GetExplicitLeft(c.instance)
}

func (c *TCheckBox) ExplicitTop() int32 {
    return CheckBox_GetExplicitTop(c.instance)
}

func (c *TCheckBox) ExplicitWidth() int32 {
    return CheckBox_GetExplicitWidth(c.instance)
}

func (c *TCheckBox) ExplicitHeight() int32 {
    return CheckBox_GetExplicitHeight(c.instance)
}

func (c *TCheckBox) Parent() *TControl {
    return ControlFromInst(CheckBox_GetParent(c.instance))
}

func (c *TCheckBox) SetParent(value IControl) {
    CheckBox_SetParent(c.instance, CheckPtr(value))
}

func (c *TCheckBox) AlignWithMargins() bool {
    return CheckBox_GetAlignWithMargins(c.instance)
}

func (c *TCheckBox) SetAlignWithMargins(value bool) {
    CheckBox_SetAlignWithMargins(c.instance, value)
}

func (c *TCheckBox) Left() int32 {
    return CheckBox_GetLeft(c.instance)
}

func (c *TCheckBox) SetLeft(value int32) {
    CheckBox_SetLeft(c.instance, value)
}

func (c *TCheckBox) Top() int32 {
    return CheckBox_GetTop(c.instance)
}

func (c *TCheckBox) SetTop(value int32) {
    CheckBox_SetTop(c.instance, value)
}

func (c *TCheckBox) Width() int32 {
    return CheckBox_GetWidth(c.instance)
}

func (c *TCheckBox) SetWidth(value int32) {
    CheckBox_SetWidth(c.instance, value)
}

func (c *TCheckBox) Height() int32 {
    return CheckBox_GetHeight(c.instance)
}

func (c *TCheckBox) SetHeight(value int32) {
    CheckBox_SetHeight(c.instance, value)
}

func (c *TCheckBox) Cursor() TCursor {
    return CheckBox_GetCursor(c.instance)
}

func (c *TCheckBox) SetCursor(value TCursor) {
    CheckBox_SetCursor(c.instance, value)
}

func (c *TCheckBox) Hint() string {
    return CheckBox_GetHint(c.instance)
}

func (c *TCheckBox) SetHint(value string) {
    CheckBox_SetHint(c.instance, value)
}

func (c *TCheckBox) Margins() *TMargins {
    return MarginsFromInst(CheckBox_GetMargins(c.instance))
}

func (c *TCheckBox) SetMargins(value *TMargins) {
    CheckBox_SetMargins(c.instance, CheckPtr(value))
}

func (c *TCheckBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(CheckBox_GetCustomHint(c.instance))
}

func (c *TCheckBox) SetCustomHint(value IComponent) {
    CheckBox_SetCustomHint(c.instance, CheckPtr(value))
}

func (c *TCheckBox) ComponentCount() int32 {
    return CheckBox_GetComponentCount(c.instance)
}

func (c *TCheckBox) ComponentIndex() int32 {
    return CheckBox_GetComponentIndex(c.instance)
}

func (c *TCheckBox) SetComponentIndex(value int32) {
    CheckBox_SetComponentIndex(c.instance, value)
}

func (c *TCheckBox) Owner() *TComponent {
    return ComponentFromInst(CheckBox_GetOwner(c.instance))
}

func (c *TCheckBox) Name() string {
    return CheckBox_GetName(c.instance)
}

func (c *TCheckBox) SetName(value string) {
    CheckBox_SetName(c.instance, value)
}

func (c *TCheckBox) Tag() int {
    return CheckBox_GetTag(c.instance)
}

func (c *TCheckBox) SetTag(value int) {
    CheckBox_SetTag(c.instance, value)
}

func (c *TCheckBox) Controls(Index int32) *TControl {
    return ControlFromInst(CheckBox_GetControls(c.instance, Index))
}

func (c *TCheckBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CheckBox_GetComponents(c.instance, AIndex))
}

