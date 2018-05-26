
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

type TUpDown struct {
    IControl
    instance uintptr
}

func NewUpDown(owner IComponent) *TUpDown {
    u := new(TUpDown)
    u.instance = UpDown_Create(CheckPtr(owner))
    return u
}

func UpDownFromInst(inst uintptr) *TUpDown {
    u := new(TUpDown)
    u.instance = inst
    return u
}

func UpDownFromObj(obj IObject) *TUpDown {
    u := new(TUpDown)
    u.instance = CheckPtr(obj)
    return u
}

func (u *TUpDown) Free() {
    if u.instance != 0 {
        UpDown_Free(u.instance)
        u.instance = 0
    }
}

func (u *TUpDown) Instance() uintptr {
    return u.instance
}

func (u *TUpDown) IsValid() bool {
    return u.instance != 0
}

func (u *TUpDown) CanFocus() bool {
    return UpDown_CanFocus(u.instance)
}

func (u *TUpDown) FlipChildren(AllLevels bool) {
    UpDown_FlipChildren(u.instance, AllLevels)
}

func (u *TUpDown) Focused() bool {
    return UpDown_Focused(u.instance)
}

func (u *TUpDown) HandleAllocated() bool {
    return UpDown_HandleAllocated(u.instance)
}

func (u *TUpDown) Invalidate() {
    UpDown_Invalidate(u.instance)
}

func (u *TUpDown) Realign() {
    UpDown_Realign(u.instance)
}

func (u *TUpDown) Repaint() {
    UpDown_Repaint(u.instance)
}

func (u *TUpDown) ScaleBy(M int32, D int32) {
    UpDown_ScaleBy(u.instance, M , D)
}

func (u *TUpDown) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    UpDown_SetBounds(u.instance, ALeft , ATop , AWidth , AHeight)
}

func (u *TUpDown) SetFocus() {
    UpDown_SetFocus(u.instance)
}

func (u *TUpDown) Update() {
    UpDown_Update(u.instance)
}

func (u *TUpDown) BringToFront() {
    UpDown_BringToFront(u.instance)
}

func (u *TUpDown) HasParent() bool {
    return UpDown_HasParent(u.instance)
}

func (u *TUpDown) Hide() {
    UpDown_Hide(u.instance)
}

func (u *TUpDown) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return UpDown_Perform(u.instance, Msg , WParam , LParam)
}

func (u *TUpDown) Refresh() {
    UpDown_Refresh(u.instance)
}

func (u *TUpDown) SendToBack() {
    UpDown_SendToBack(u.instance)
}

func (u *TUpDown) Show() {
    UpDown_Show(u.instance)
}

func (u *TUpDown) GetTextBuf(Buffer string, BufSize int32) int32 {
    return UpDown_GetTextBuf(u.instance, Buffer , BufSize)
}

func (u *TUpDown) GetTextLen() int32 {
    return UpDown_GetTextLen(u.instance)
}

func (u *TUpDown) FindComponent(AName string) *TComponent {
    return ComponentFromInst(UpDown_FindComponent(u.instance, AName))
}

func (u *TUpDown) GetNamePath() string {
    return UpDown_GetNamePath(u.instance)
}

func (u *TUpDown) Assign(Source IObject) {
    UpDown_Assign(u.instance, CheckPtr(Source))
}

func (u *TUpDown) ClassName() string {
    return UpDown_ClassName(u.instance)
}

func (u *TUpDown) Equals(Obj IObject) bool {
    return UpDown_Equals(u.instance, CheckPtr(Obj))
}

func (u *TUpDown) GetHashCode() int32 {
    return UpDown_GetHashCode(u.instance)
}

func (u *TUpDown) ToString() string {
    return UpDown_ToString(u.instance)
}

func (u *TUpDown) Anchors() TAnchors {
    return UpDown_GetAnchors(u.instance)
}

func (u *TUpDown) SetAnchors(value TAnchors) {
    UpDown_SetAnchors(u.instance, value)
}

func (u *TUpDown) DoubleBuffered() bool {
    return UpDown_GetDoubleBuffered(u.instance)
}

func (u *TUpDown) SetDoubleBuffered(value bool) {
    UpDown_SetDoubleBuffered(u.instance, value)
}

func (u *TUpDown) Enabled() bool {
    return UpDown_GetEnabled(u.instance)
}

func (u *TUpDown) SetEnabled(value bool) {
    UpDown_SetEnabled(u.instance, value)
}

func (u *TUpDown) Hint() string {
    return UpDown_GetHint(u.instance)
}

func (u *TUpDown) SetHint(value string) {
    UpDown_SetHint(u.instance, value)
}

func (u *TUpDown) Min() int32 {
    return UpDown_GetMin(u.instance)
}

func (u *TUpDown) SetMin(value int32) {
    UpDown_SetMin(u.instance, value)
}

func (u *TUpDown) Max() int32 {
    return UpDown_GetMax(u.instance)
}

func (u *TUpDown) SetMax(value int32) {
    UpDown_SetMax(u.instance, value)
}

func (u *TUpDown) Orientation() TUDOrientation {
    return UpDown_GetOrientation(u.instance)
}

func (u *TUpDown) SetOrientation(value TUDOrientation) {
    UpDown_SetOrientation(u.instance, value)
}

func (u *TUpDown) ParentDoubleBuffered() bool {
    return UpDown_GetParentDoubleBuffered(u.instance)
}

func (u *TUpDown) SetParentDoubleBuffered(value bool) {
    UpDown_SetParentDoubleBuffered(u.instance, value)
}

func (u *TUpDown) ParentShowHint() bool {
    return UpDown_GetParentShowHint(u.instance)
}

func (u *TUpDown) SetParentShowHint(value bool) {
    UpDown_SetParentShowHint(u.instance, value)
}

func (u *TUpDown) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(UpDown_GetPopupMenu(u.instance))
}

func (u *TUpDown) SetPopupMenu(value IComponent) {
    UpDown_SetPopupMenu(u.instance, CheckPtr(value))
}

func (u *TUpDown) Position() int32 {
    return UpDown_GetPosition(u.instance)
}

func (u *TUpDown) SetPosition(value int32) {
    UpDown_SetPosition(u.instance, value)
}

func (u *TUpDown) ShowHint() bool {
    return UpDown_GetShowHint(u.instance)
}

func (u *TUpDown) SetShowHint(value bool) {
    UpDown_SetShowHint(u.instance, value)
}

func (u *TUpDown) TabOrder() uint16 {
    return UpDown_GetTabOrder(u.instance)
}

func (u *TUpDown) SetTabOrder(value uint16) {
    UpDown_SetTabOrder(u.instance, value)
}

func (u *TUpDown) TabStop() bool {
    return UpDown_GetTabStop(u.instance)
}

func (u *TUpDown) SetTabStop(value bool) {
    UpDown_SetTabStop(u.instance, value)
}

func (u *TUpDown) Visible() bool {
    return UpDown_GetVisible(u.instance)
}

func (u *TUpDown) SetVisible(value bool) {
    UpDown_SetVisible(u.instance, value)
}

func (u *TUpDown) Wrap() bool {
    return UpDown_GetWrap(u.instance)
}

func (u *TUpDown) SetWrap(value bool) {
    UpDown_SetWrap(u.instance, value)
}

func (u *TUpDown) StyleElements() TStyleElements {
    return UpDown_GetStyleElements(u.instance)
}

func (u *TUpDown) SetStyleElements(value TStyleElements) {
    UpDown_SetStyleElements(u.instance, value)
}

func (u *TUpDown) SetOnClick(fn TUDClickEvent) {
    UpDown_SetOnClick(u.instance, fn)
}

func (u *TUpDown) SetOnEnter(fn TNotifyEvent) {
    UpDown_SetOnEnter(u.instance, fn)
}

func (u *TUpDown) SetOnExit(fn TNotifyEvent) {
    UpDown_SetOnExit(u.instance, fn)
}

func (u *TUpDown) SetOnMouseDown(fn TMouseEvent) {
    UpDown_SetOnMouseDown(u.instance, fn)
}

func (u *TUpDown) SetOnMouseEnter(fn TNotifyEvent) {
    UpDown_SetOnMouseEnter(u.instance, fn)
}

func (u *TUpDown) SetOnMouseLeave(fn TNotifyEvent) {
    UpDown_SetOnMouseLeave(u.instance, fn)
}

func (u *TUpDown) SetOnMouseMove(fn TMouseMoveEvent) {
    UpDown_SetOnMouseMove(u.instance, fn)
}

func (u *TUpDown) SetOnMouseUp(fn TMouseEvent) {
    UpDown_SetOnMouseUp(u.instance, fn)
}

func (u *TUpDown) Brush() *TBrush {
    return BrushFromInst(UpDown_GetBrush(u.instance))
}

func (u *TUpDown) ControlCount() int32 {
    return UpDown_GetControlCount(u.instance)
}

func (u *TUpDown) Handle() HWND {
    return UpDown_GetHandle(u.instance)
}

func (u *TUpDown) ParentWindow() HWND {
    return UpDown_GetParentWindow(u.instance)
}

func (u *TUpDown) SetParentWindow(value HWND) {
    UpDown_SetParentWindow(u.instance, value)
}

func (u *TUpDown) Action() *TAction {
    return ActionFromInst(UpDown_GetAction(u.instance))
}

func (u *TUpDown) SetAction(value IComponent) {
    UpDown_SetAction(u.instance, CheckPtr(value))
}

func (u *TUpDown) Align() TAlign {
    return UpDown_GetAlign(u.instance)
}

func (u *TUpDown) SetAlign(value TAlign) {
    UpDown_SetAlign(u.instance, value)
}

func (u *TUpDown) BiDiMode() TBiDiMode {
    return UpDown_GetBiDiMode(u.instance)
}

func (u *TUpDown) SetBiDiMode(value TBiDiMode) {
    UpDown_SetBiDiMode(u.instance, value)
}

func (u *TUpDown) BoundsRect() TRect {
    return UpDown_GetBoundsRect(u.instance)
}

func (u *TUpDown) SetBoundsRect(value TRect) {
    UpDown_SetBoundsRect(u.instance, value)
}

func (u *TUpDown) ClientHeight() int32 {
    return UpDown_GetClientHeight(u.instance)
}

func (u *TUpDown) SetClientHeight(value int32) {
    UpDown_SetClientHeight(u.instance, value)
}

func (u *TUpDown) ClientRect() TRect {
    return UpDown_GetClientRect(u.instance)
}

func (u *TUpDown) ClientWidth() int32 {
    return UpDown_GetClientWidth(u.instance)
}

func (u *TUpDown) SetClientWidth(value int32) {
    UpDown_SetClientWidth(u.instance, value)
}

func (u *TUpDown) ExplicitLeft() int32 {
    return UpDown_GetExplicitLeft(u.instance)
}

func (u *TUpDown) ExplicitTop() int32 {
    return UpDown_GetExplicitTop(u.instance)
}

func (u *TUpDown) ExplicitWidth() int32 {
    return UpDown_GetExplicitWidth(u.instance)
}

func (u *TUpDown) ExplicitHeight() int32 {
    return UpDown_GetExplicitHeight(u.instance)
}

func (u *TUpDown) Parent() *TControl {
    return ControlFromInst(UpDown_GetParent(u.instance))
}

func (u *TUpDown) SetParent(value IControl) {
    UpDown_SetParent(u.instance, CheckPtr(value))
}

func (u *TUpDown) AlignWithMargins() bool {
    return UpDown_GetAlignWithMargins(u.instance)
}

func (u *TUpDown) SetAlignWithMargins(value bool) {
    UpDown_SetAlignWithMargins(u.instance, value)
}

func (u *TUpDown) Left() int32 {
    return UpDown_GetLeft(u.instance)
}

func (u *TUpDown) SetLeft(value int32) {
    UpDown_SetLeft(u.instance, value)
}

func (u *TUpDown) Top() int32 {
    return UpDown_GetTop(u.instance)
}

func (u *TUpDown) SetTop(value int32) {
    UpDown_SetTop(u.instance, value)
}

func (u *TUpDown) Width() int32 {
    return UpDown_GetWidth(u.instance)
}

func (u *TUpDown) SetWidth(value int32) {
    UpDown_SetWidth(u.instance, value)
}

func (u *TUpDown) Height() int32 {
    return UpDown_GetHeight(u.instance)
}

func (u *TUpDown) SetHeight(value int32) {
    UpDown_SetHeight(u.instance, value)
}

func (u *TUpDown) Cursor() TCursor {
    return UpDown_GetCursor(u.instance)
}

func (u *TUpDown) SetCursor(value TCursor) {
    UpDown_SetCursor(u.instance, value)
}

func (u *TUpDown) Margins() *TMargins {
    return MarginsFromInst(UpDown_GetMargins(u.instance))
}

func (u *TUpDown) SetMargins(value *TMargins) {
    UpDown_SetMargins(u.instance, CheckPtr(value))
}

func (u *TUpDown) CustomHint() *TCustomHint {
    return CustomHintFromInst(UpDown_GetCustomHint(u.instance))
}

func (u *TUpDown) SetCustomHint(value IComponent) {
    UpDown_SetCustomHint(u.instance, CheckPtr(value))
}

func (u *TUpDown) ComponentCount() int32 {
    return UpDown_GetComponentCount(u.instance)
}

func (u *TUpDown) ComponentIndex() int32 {
    return UpDown_GetComponentIndex(u.instance)
}

func (u *TUpDown) SetComponentIndex(value int32) {
    UpDown_SetComponentIndex(u.instance, value)
}

func (u *TUpDown) Owner() *TComponent {
    return ComponentFromInst(UpDown_GetOwner(u.instance))
}

func (u *TUpDown) Name() string {
    return UpDown_GetName(u.instance)
}

func (u *TUpDown) SetName(value string) {
    UpDown_SetName(u.instance, value)
}

func (u *TUpDown) Tag() int {
    return UpDown_GetTag(u.instance)
}

func (u *TUpDown) SetTag(value int) {
    UpDown_SetTag(u.instance, value)
}

func (u *TUpDown) Controls(Index int32) *TControl {
    return ControlFromInst(UpDown_GetControls(u.instance, Index))
}

func (u *TUpDown) Components(AIndex int32) *TComponent {
    return ComponentFromInst(UpDown_GetComponents(u.instance, AIndex))
}

