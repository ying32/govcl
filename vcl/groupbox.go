
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

type TGroupBox struct {
    IControl
    instance uintptr
}

func NewGroupBox(owner IComponent) *TGroupBox {
    g := new(TGroupBox)
    g.instance = GroupBox_Create(CheckPtr(owner))
    return g
}

func GroupBoxFromInst(inst uintptr) *TGroupBox {
    g := new(TGroupBox)
    g.instance = inst
    return g
}

func GroupBoxFromObj(obj IObject) *TGroupBox {
    g := new(TGroupBox)
    g.instance = CheckPtr(obj)
    return g
}

func (g *TGroupBox) Free() {
    if g.instance != 0 {
        GroupBox_Free(g.instance)
        g.instance = 0
    }
}

func (g *TGroupBox) Instance() uintptr {
    return g.instance
}

func (g *TGroupBox) IsValid() bool {
    return g.instance != 0
}

func (g *TGroupBox) CanFocus() bool {
    return GroupBox_CanFocus(g.instance)
}

func (g *TGroupBox) FlipChildren(AllLevels bool) {
    GroupBox_FlipChildren(g.instance, AllLevels)
}

func (g *TGroupBox) Focused() bool {
    return GroupBox_Focused(g.instance)
}

func (g *TGroupBox) HandleAllocated() bool {
    return GroupBox_HandleAllocated(g.instance)
}

func (g *TGroupBox) Invalidate() {
    GroupBox_Invalidate(g.instance)
}

func (g *TGroupBox) Realign() {
    GroupBox_Realign(g.instance)
}

func (g *TGroupBox) Repaint() {
    GroupBox_Repaint(g.instance)
}

func (g *TGroupBox) ScaleBy(M int32, D int32) {
    GroupBox_ScaleBy(g.instance, M , D)
}

func (g *TGroupBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    GroupBox_SetBounds(g.instance, ALeft , ATop , AWidth , AHeight)
}

func (g *TGroupBox) SetFocus() {
    GroupBox_SetFocus(g.instance)
}

func (g *TGroupBox) Update() {
    GroupBox_Update(g.instance)
}

func (g *TGroupBox) BringToFront() {
    GroupBox_BringToFront(g.instance)
}

func (g *TGroupBox) HasParent() bool {
    return GroupBox_HasParent(g.instance)
}

func (g *TGroupBox) Hide() {
    GroupBox_Hide(g.instance)
}

func (g *TGroupBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return GroupBox_Perform(g.instance, Msg , WParam , LParam)
}

func (g *TGroupBox) Refresh() {
    GroupBox_Refresh(g.instance)
}

func (g *TGroupBox) SendToBack() {
    GroupBox_SendToBack(g.instance)
}

func (g *TGroupBox) Show() {
    GroupBox_Show(g.instance)
}

func (g *TGroupBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return GroupBox_GetTextBuf(g.instance, Buffer , BufSize)
}

func (g *TGroupBox) GetTextLen() int32 {
    return GroupBox_GetTextLen(g.instance)
}

func (g *TGroupBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(GroupBox_FindComponent(g.instance, AName))
}

func (g *TGroupBox) GetNamePath() string {
    return GroupBox_GetNamePath(g.instance)
}

func (g *TGroupBox) Assign(Source IObject) {
    GroupBox_Assign(g.instance, CheckPtr(Source))
}

func (g *TGroupBox) ClassName() string {
    return GroupBox_ClassName(g.instance)
}

func (g *TGroupBox) Equals(Obj IObject) bool {
    return GroupBox_Equals(g.instance, CheckPtr(Obj))
}

func (g *TGroupBox) GetHashCode() int32 {
    return GroupBox_GetHashCode(g.instance)
}

func (g *TGroupBox) ToString() string {
    return GroupBox_ToString(g.instance)
}

func (g *TGroupBox) Align() TAlign {
    return GroupBox_GetAlign(g.instance)
}

func (g *TGroupBox) SetAlign(value TAlign) {
    GroupBox_SetAlign(g.instance, value)
}

func (g *TGroupBox) Anchors() TAnchors {
    return GroupBox_GetAnchors(g.instance)
}

func (g *TGroupBox) SetAnchors(value TAnchors) {
    GroupBox_SetAnchors(g.instance, value)
}

func (g *TGroupBox) BiDiMode() TBiDiMode {
    return GroupBox_GetBiDiMode(g.instance)
}

func (g *TGroupBox) SetBiDiMode(value TBiDiMode) {
    GroupBox_SetBiDiMode(g.instance, value)
}

func (g *TGroupBox) Caption() string {
    return GroupBox_GetCaption(g.instance)
}

func (g *TGroupBox) SetCaption(value string) {
    GroupBox_SetCaption(g.instance, value)
}

func (g *TGroupBox) Color() TColor {
    return GroupBox_GetColor(g.instance)
}

func (g *TGroupBox) SetColor(value TColor) {
    GroupBox_SetColor(g.instance, value)
}

func (g *TGroupBox) DoubleBuffered() bool {
    return GroupBox_GetDoubleBuffered(g.instance)
}

func (g *TGroupBox) SetDoubleBuffered(value bool) {
    GroupBox_SetDoubleBuffered(g.instance, value)
}

func (g *TGroupBox) Enabled() bool {
    return GroupBox_GetEnabled(g.instance)
}

func (g *TGroupBox) SetEnabled(value bool) {
    GroupBox_SetEnabled(g.instance, value)
}

func (g *TGroupBox) Font() *TFont {
    return FontFromInst(GroupBox_GetFont(g.instance))
}

func (g *TGroupBox) SetFont(value *TFont) {
    GroupBox_SetFont(g.instance, CheckPtr(value))
}

func (g *TGroupBox) ParentBackground() bool {
    return GroupBox_GetParentBackground(g.instance)
}

func (g *TGroupBox) SetParentBackground(value bool) {
    GroupBox_SetParentBackground(g.instance, value)
}

func (g *TGroupBox) ParentColor() bool {
    return GroupBox_GetParentColor(g.instance)
}

func (g *TGroupBox) SetParentColor(value bool) {
    GroupBox_SetParentColor(g.instance, value)
}

func (g *TGroupBox) ParentCtl3D() bool {
    return GroupBox_GetParentCtl3D(g.instance)
}

func (g *TGroupBox) SetParentCtl3D(value bool) {
    GroupBox_SetParentCtl3D(g.instance, value)
}

func (g *TGroupBox) ParentDoubleBuffered() bool {
    return GroupBox_GetParentDoubleBuffered(g.instance)
}

func (g *TGroupBox) SetParentDoubleBuffered(value bool) {
    GroupBox_SetParentDoubleBuffered(g.instance, value)
}

func (g *TGroupBox) ParentFont() bool {
    return GroupBox_GetParentFont(g.instance)
}

func (g *TGroupBox) SetParentFont(value bool) {
    GroupBox_SetParentFont(g.instance, value)
}

func (g *TGroupBox) ParentShowHint() bool {
    return GroupBox_GetParentShowHint(g.instance)
}

func (g *TGroupBox) SetParentShowHint(value bool) {
    GroupBox_SetParentShowHint(g.instance, value)
}

func (g *TGroupBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(GroupBox_GetPopupMenu(g.instance))
}

func (g *TGroupBox) SetPopupMenu(value IComponent) {
    GroupBox_SetPopupMenu(g.instance, CheckPtr(value))
}

func (g *TGroupBox) ShowHint() bool {
    return GroupBox_GetShowHint(g.instance)
}

func (g *TGroupBox) SetShowHint(value bool) {
    GroupBox_SetShowHint(g.instance, value)
}

func (g *TGroupBox) TabOrder() uint16 {
    return GroupBox_GetTabOrder(g.instance)
}

func (g *TGroupBox) SetTabOrder(value uint16) {
    GroupBox_SetTabOrder(g.instance, value)
}

func (g *TGroupBox) TabStop() bool {
    return GroupBox_GetTabStop(g.instance)
}

func (g *TGroupBox) SetTabStop(value bool) {
    GroupBox_SetTabStop(g.instance, value)
}

func (g *TGroupBox) Visible() bool {
    return GroupBox_GetVisible(g.instance)
}

func (g *TGroupBox) SetVisible(value bool) {
    GroupBox_SetVisible(g.instance, value)
}

func (g *TGroupBox) StyleElements() TStyleElements {
    return GroupBox_GetStyleElements(g.instance)
}

func (g *TGroupBox) SetStyleElements(value TStyleElements) {
    GroupBox_SetStyleElements(g.instance, value)
}

func (g *TGroupBox) SetOnClick(fn TNotifyEvent) {
    GroupBox_SetOnClick(g.instance, fn)
}

func (g *TGroupBox) SetOnDblClick(fn TNotifyEvent) {
    GroupBox_SetOnDblClick(g.instance, fn)
}

func (g *TGroupBox) SetOnEnter(fn TNotifyEvent) {
    GroupBox_SetOnEnter(g.instance, fn)
}

func (g *TGroupBox) SetOnExit(fn TNotifyEvent) {
    GroupBox_SetOnExit(g.instance, fn)
}

func (g *TGroupBox) SetOnMouseDown(fn TMouseEvent) {
    GroupBox_SetOnMouseDown(g.instance, fn)
}

func (g *TGroupBox) SetOnMouseEnter(fn TNotifyEvent) {
    GroupBox_SetOnMouseEnter(g.instance, fn)
}

func (g *TGroupBox) SetOnMouseLeave(fn TNotifyEvent) {
    GroupBox_SetOnMouseLeave(g.instance, fn)
}

func (g *TGroupBox) SetOnMouseMove(fn TMouseMoveEvent) {
    GroupBox_SetOnMouseMove(g.instance, fn)
}

func (g *TGroupBox) SetOnMouseUp(fn TMouseEvent) {
    GroupBox_SetOnMouseUp(g.instance, fn)
}

func (g *TGroupBox) Brush() *TBrush {
    return BrushFromInst(GroupBox_GetBrush(g.instance))
}

func (g *TGroupBox) ControlCount() int32 {
    return GroupBox_GetControlCount(g.instance)
}

func (g *TGroupBox) Handle() HWND {
    return GroupBox_GetHandle(g.instance)
}

func (g *TGroupBox) ParentWindow() HWND {
    return GroupBox_GetParentWindow(g.instance)
}

func (g *TGroupBox) SetParentWindow(value HWND) {
    GroupBox_SetParentWindow(g.instance, value)
}

func (g *TGroupBox) Action() *TAction {
    return ActionFromInst(GroupBox_GetAction(g.instance))
}

func (g *TGroupBox) SetAction(value IComponent) {
    GroupBox_SetAction(g.instance, CheckPtr(value))
}

func (g *TGroupBox) BoundsRect() TRect {
    return GroupBox_GetBoundsRect(g.instance)
}

func (g *TGroupBox) SetBoundsRect(value TRect) {
    GroupBox_SetBoundsRect(g.instance, value)
}

func (g *TGroupBox) ClientHeight() int32 {
    return GroupBox_GetClientHeight(g.instance)
}

func (g *TGroupBox) SetClientHeight(value int32) {
    GroupBox_SetClientHeight(g.instance, value)
}

func (g *TGroupBox) ClientRect() TRect {
    return GroupBox_GetClientRect(g.instance)
}

func (g *TGroupBox) ClientWidth() int32 {
    return GroupBox_GetClientWidth(g.instance)
}

func (g *TGroupBox) SetClientWidth(value int32) {
    GroupBox_SetClientWidth(g.instance, value)
}

func (g *TGroupBox) ExplicitLeft() int32 {
    return GroupBox_GetExplicitLeft(g.instance)
}

func (g *TGroupBox) ExplicitTop() int32 {
    return GroupBox_GetExplicitTop(g.instance)
}

func (g *TGroupBox) ExplicitWidth() int32 {
    return GroupBox_GetExplicitWidth(g.instance)
}

func (g *TGroupBox) ExplicitHeight() int32 {
    return GroupBox_GetExplicitHeight(g.instance)
}

func (g *TGroupBox) Parent() *TControl {
    return ControlFromInst(GroupBox_GetParent(g.instance))
}

func (g *TGroupBox) SetParent(value IControl) {
    GroupBox_SetParent(g.instance, CheckPtr(value))
}

func (g *TGroupBox) AlignWithMargins() bool {
    return GroupBox_GetAlignWithMargins(g.instance)
}

func (g *TGroupBox) SetAlignWithMargins(value bool) {
    GroupBox_SetAlignWithMargins(g.instance, value)
}

func (g *TGroupBox) Left() int32 {
    return GroupBox_GetLeft(g.instance)
}

func (g *TGroupBox) SetLeft(value int32) {
    GroupBox_SetLeft(g.instance, value)
}

func (g *TGroupBox) Top() int32 {
    return GroupBox_GetTop(g.instance)
}

func (g *TGroupBox) SetTop(value int32) {
    GroupBox_SetTop(g.instance, value)
}

func (g *TGroupBox) Width() int32 {
    return GroupBox_GetWidth(g.instance)
}

func (g *TGroupBox) SetWidth(value int32) {
    GroupBox_SetWidth(g.instance, value)
}

func (g *TGroupBox) Height() int32 {
    return GroupBox_GetHeight(g.instance)
}

func (g *TGroupBox) SetHeight(value int32) {
    GroupBox_SetHeight(g.instance, value)
}

func (g *TGroupBox) Cursor() TCursor {
    return GroupBox_GetCursor(g.instance)
}

func (g *TGroupBox) SetCursor(value TCursor) {
    GroupBox_SetCursor(g.instance, value)
}

func (g *TGroupBox) Hint() string {
    return GroupBox_GetHint(g.instance)
}

func (g *TGroupBox) SetHint(value string) {
    GroupBox_SetHint(g.instance, value)
}

func (g *TGroupBox) Margins() *TMargins {
    return MarginsFromInst(GroupBox_GetMargins(g.instance))
}

func (g *TGroupBox) SetMargins(value *TMargins) {
    GroupBox_SetMargins(g.instance, CheckPtr(value))
}

func (g *TGroupBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(GroupBox_GetCustomHint(g.instance))
}

func (g *TGroupBox) SetCustomHint(value IComponent) {
    GroupBox_SetCustomHint(g.instance, CheckPtr(value))
}

func (g *TGroupBox) ComponentCount() int32 {
    return GroupBox_GetComponentCount(g.instance)
}

func (g *TGroupBox) ComponentIndex() int32 {
    return GroupBox_GetComponentIndex(g.instance)
}

func (g *TGroupBox) SetComponentIndex(value int32) {
    GroupBox_SetComponentIndex(g.instance, value)
}

func (g *TGroupBox) Owner() *TComponent {
    return ComponentFromInst(GroupBox_GetOwner(g.instance))
}

func (g *TGroupBox) Name() string {
    return GroupBox_GetName(g.instance)
}

func (g *TGroupBox) SetName(value string) {
    GroupBox_SetName(g.instance, value)
}

func (g *TGroupBox) Tag() int {
    return GroupBox_GetTag(g.instance)
}

func (g *TGroupBox) SetTag(value int) {
    GroupBox_SetTag(g.instance, value)
}

func (g *TGroupBox) Controls(Index int32) *TControl {
    return ControlFromInst(GroupBox_GetControls(g.instance, Index))
}

func (g *TGroupBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(GroupBox_GetComponents(g.instance, AIndex))
}

