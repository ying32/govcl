
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

type THotKey struct {
    IControl
    instance uintptr
}

func NewHotKey(owner IComponent) *THotKey {
    h := new(THotKey)
    h.instance = HotKey_Create(CheckPtr(owner))
    return h
}

func HotKeyFromInst(inst uintptr) *THotKey {
    h := new(THotKey)
    h.instance = inst
    return h
}

func HotKeyFromObj(obj IObject) *THotKey {
    h := new(THotKey)
    h.instance = CheckPtr(obj)
    return h
}

func (h *THotKey) Free() {
    if h.instance != 0 {
        HotKey_Free(h.instance)
        h.instance = 0
    }
}

func (h *THotKey) Instance() uintptr {
    return h.instance
}

func (h *THotKey) IsValid() bool {
    return h.instance != 0
}

func (h *THotKey) CanFocus() bool {
    return HotKey_CanFocus(h.instance)
}

func (h *THotKey) FlipChildren(AllLevels bool) {
    HotKey_FlipChildren(h.instance, AllLevels)
}

func (h *THotKey) Focused() bool {
    return HotKey_Focused(h.instance)
}

func (h *THotKey) HandleAllocated() bool {
    return HotKey_HandleAllocated(h.instance)
}

func (h *THotKey) Invalidate() {
    HotKey_Invalidate(h.instance)
}

func (h *THotKey) Realign() {
    HotKey_Realign(h.instance)
}

func (h *THotKey) Repaint() {
    HotKey_Repaint(h.instance)
}

func (h *THotKey) ScaleBy(M int32, D int32) {
    HotKey_ScaleBy(h.instance, M , D)
}

func (h *THotKey) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    HotKey_SetBounds(h.instance, ALeft , ATop , AWidth , AHeight)
}

func (h *THotKey) SetFocus() {
    HotKey_SetFocus(h.instance)
}

func (h *THotKey) Update() {
    HotKey_Update(h.instance)
}

func (h *THotKey) BringToFront() {
    HotKey_BringToFront(h.instance)
}

func (h *THotKey) HasParent() bool {
    return HotKey_HasParent(h.instance)
}

func (h *THotKey) Hide() {
    HotKey_Hide(h.instance)
}

func (h *THotKey) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return HotKey_Perform(h.instance, Msg , WParam , LParam)
}

func (h *THotKey) Refresh() {
    HotKey_Refresh(h.instance)
}

func (h *THotKey) SendToBack() {
    HotKey_SendToBack(h.instance)
}

func (h *THotKey) Show() {
    HotKey_Show(h.instance)
}

func (h *THotKey) GetTextBuf(Buffer string, BufSize int32) int32 {
    return HotKey_GetTextBuf(h.instance, Buffer , BufSize)
}

func (h *THotKey) GetTextLen() int32 {
    return HotKey_GetTextLen(h.instance)
}

func (h *THotKey) FindComponent(AName string) *TComponent {
    return ComponentFromInst(HotKey_FindComponent(h.instance, AName))
}

func (h *THotKey) GetNamePath() string {
    return HotKey_GetNamePath(h.instance)
}

func (h *THotKey) Assign(Source IObject) {
    HotKey_Assign(h.instance, CheckPtr(Source))
}

func (h *THotKey) ClassName() string {
    return HotKey_ClassName(h.instance)
}

func (h *THotKey) Equals(Obj IObject) bool {
    return HotKey_Equals(h.instance, CheckPtr(Obj))
}

func (h *THotKey) GetHashCode() int32 {
    return HotKey_GetHashCode(h.instance)
}

func (h *THotKey) ToString() string {
    return HotKey_ToString(h.instance)
}

func (h *THotKey) Anchors() TAnchors {
    return HotKey_GetAnchors(h.instance)
}

func (h *THotKey) SetAnchors(value TAnchors) {
    HotKey_SetAnchors(h.instance, value)
}

func (h *THotKey) AutoSize() bool {
    return HotKey_GetAutoSize(h.instance)
}

func (h *THotKey) SetAutoSize(value bool) {
    HotKey_SetAutoSize(h.instance, value)
}

func (h *THotKey) BiDiMode() TBiDiMode {
    return HotKey_GetBiDiMode(h.instance)
}

func (h *THotKey) SetBiDiMode(value TBiDiMode) {
    HotKey_SetBiDiMode(h.instance, value)
}

func (h *THotKey) Enabled() bool {
    return HotKey_GetEnabled(h.instance)
}

func (h *THotKey) SetEnabled(value bool) {
    HotKey_SetEnabled(h.instance, value)
}

func (h *THotKey) Hint() string {
    return HotKey_GetHint(h.instance)
}

func (h *THotKey) SetHint(value string) {
    HotKey_SetHint(h.instance, value)
}

func (h *THotKey) ParentShowHint() bool {
    return HotKey_GetParentShowHint(h.instance)
}

func (h *THotKey) SetParentShowHint(value bool) {
    HotKey_SetParentShowHint(h.instance, value)
}

func (h *THotKey) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(HotKey_GetPopupMenu(h.instance))
}

func (h *THotKey) SetPopupMenu(value IComponent) {
    HotKey_SetPopupMenu(h.instance, CheckPtr(value))
}

func (h *THotKey) ShowHint() bool {
    return HotKey_GetShowHint(h.instance)
}

func (h *THotKey) SetShowHint(value bool) {
    HotKey_SetShowHint(h.instance, value)
}

func (h *THotKey) TabOrder() uint16 {
    return HotKey_GetTabOrder(h.instance)
}

func (h *THotKey) SetTabOrder(value uint16) {
    HotKey_SetTabOrder(h.instance, value)
}

func (h *THotKey) TabStop() bool {
    return HotKey_GetTabStop(h.instance)
}

func (h *THotKey) SetTabStop(value bool) {
    HotKey_SetTabStop(h.instance, value)
}

func (h *THotKey) Visible() bool {
    return HotKey_GetVisible(h.instance)
}

func (h *THotKey) SetVisible(value bool) {
    HotKey_SetVisible(h.instance, value)
}

func (h *THotKey) StyleElements() TStyleElements {
    return HotKey_GetStyleElements(h.instance)
}

func (h *THotKey) SetStyleElements(value TStyleElements) {
    HotKey_SetStyleElements(h.instance, value)
}

func (h *THotKey) SetOnChange(fn TNotifyEvent) {
    HotKey_SetOnChange(h.instance, fn)
}

func (h *THotKey) SetOnEnter(fn TNotifyEvent) {
    HotKey_SetOnEnter(h.instance, fn)
}

func (h *THotKey) SetOnExit(fn TNotifyEvent) {
    HotKey_SetOnExit(h.instance, fn)
}

func (h *THotKey) SetOnMouseDown(fn TMouseEvent) {
    HotKey_SetOnMouseDown(h.instance, fn)
}

func (h *THotKey) SetOnMouseEnter(fn TNotifyEvent) {
    HotKey_SetOnMouseEnter(h.instance, fn)
}

func (h *THotKey) SetOnMouseLeave(fn TNotifyEvent) {
    HotKey_SetOnMouseLeave(h.instance, fn)
}

func (h *THotKey) SetOnMouseMove(fn TMouseMoveEvent) {
    HotKey_SetOnMouseMove(h.instance, fn)
}

func (h *THotKey) SetOnMouseUp(fn TMouseEvent) {
    HotKey_SetOnMouseUp(h.instance, fn)
}

func (h *THotKey) DoubleBuffered() bool {
    return HotKey_GetDoubleBuffered(h.instance)
}

func (h *THotKey) SetDoubleBuffered(value bool) {
    HotKey_SetDoubleBuffered(h.instance, value)
}

func (h *THotKey) Brush() *TBrush {
    return BrushFromInst(HotKey_GetBrush(h.instance))
}

func (h *THotKey) ControlCount() int32 {
    return HotKey_GetControlCount(h.instance)
}

func (h *THotKey) Handle() HWND {
    return HotKey_GetHandle(h.instance)
}

func (h *THotKey) ParentDoubleBuffered() bool {
    return HotKey_GetParentDoubleBuffered(h.instance)
}

func (h *THotKey) SetParentDoubleBuffered(value bool) {
    HotKey_SetParentDoubleBuffered(h.instance, value)
}

func (h *THotKey) ParentWindow() HWND {
    return HotKey_GetParentWindow(h.instance)
}

func (h *THotKey) SetParentWindow(value HWND) {
    HotKey_SetParentWindow(h.instance, value)
}

func (h *THotKey) Action() *TAction {
    return ActionFromInst(HotKey_GetAction(h.instance))
}

func (h *THotKey) SetAction(value IComponent) {
    HotKey_SetAction(h.instance, CheckPtr(value))
}

func (h *THotKey) Align() TAlign {
    return HotKey_GetAlign(h.instance)
}

func (h *THotKey) SetAlign(value TAlign) {
    HotKey_SetAlign(h.instance, value)
}

func (h *THotKey) BoundsRect() TRect {
    return HotKey_GetBoundsRect(h.instance)
}

func (h *THotKey) SetBoundsRect(value TRect) {
    HotKey_SetBoundsRect(h.instance, value)
}

func (h *THotKey) ClientHeight() int32 {
    return HotKey_GetClientHeight(h.instance)
}

func (h *THotKey) SetClientHeight(value int32) {
    HotKey_SetClientHeight(h.instance, value)
}

func (h *THotKey) ClientRect() TRect {
    return HotKey_GetClientRect(h.instance)
}

func (h *THotKey) ClientWidth() int32 {
    return HotKey_GetClientWidth(h.instance)
}

func (h *THotKey) SetClientWidth(value int32) {
    HotKey_SetClientWidth(h.instance, value)
}

func (h *THotKey) ExplicitLeft() int32 {
    return HotKey_GetExplicitLeft(h.instance)
}

func (h *THotKey) ExplicitTop() int32 {
    return HotKey_GetExplicitTop(h.instance)
}

func (h *THotKey) ExplicitWidth() int32 {
    return HotKey_GetExplicitWidth(h.instance)
}

func (h *THotKey) ExplicitHeight() int32 {
    return HotKey_GetExplicitHeight(h.instance)
}

func (h *THotKey) Parent() *TControl {
    return ControlFromInst(HotKey_GetParent(h.instance))
}

func (h *THotKey) SetParent(value IControl) {
    HotKey_SetParent(h.instance, CheckPtr(value))
}

func (h *THotKey) AlignWithMargins() bool {
    return HotKey_GetAlignWithMargins(h.instance)
}

func (h *THotKey) SetAlignWithMargins(value bool) {
    HotKey_SetAlignWithMargins(h.instance, value)
}

func (h *THotKey) Left() int32 {
    return HotKey_GetLeft(h.instance)
}

func (h *THotKey) SetLeft(value int32) {
    HotKey_SetLeft(h.instance, value)
}

func (h *THotKey) Top() int32 {
    return HotKey_GetTop(h.instance)
}

func (h *THotKey) SetTop(value int32) {
    HotKey_SetTop(h.instance, value)
}

func (h *THotKey) Width() int32 {
    return HotKey_GetWidth(h.instance)
}

func (h *THotKey) SetWidth(value int32) {
    HotKey_SetWidth(h.instance, value)
}

func (h *THotKey) Height() int32 {
    return HotKey_GetHeight(h.instance)
}

func (h *THotKey) SetHeight(value int32) {
    HotKey_SetHeight(h.instance, value)
}

func (h *THotKey) Cursor() TCursor {
    return HotKey_GetCursor(h.instance)
}

func (h *THotKey) SetCursor(value TCursor) {
    HotKey_SetCursor(h.instance, value)
}

func (h *THotKey) Margins() *TMargins {
    return MarginsFromInst(HotKey_GetMargins(h.instance))
}

func (h *THotKey) SetMargins(value *TMargins) {
    HotKey_SetMargins(h.instance, CheckPtr(value))
}

func (h *THotKey) CustomHint() *TCustomHint {
    return CustomHintFromInst(HotKey_GetCustomHint(h.instance))
}

func (h *THotKey) SetCustomHint(value IComponent) {
    HotKey_SetCustomHint(h.instance, CheckPtr(value))
}

func (h *THotKey) ComponentCount() int32 {
    return HotKey_GetComponentCount(h.instance)
}

func (h *THotKey) ComponentIndex() int32 {
    return HotKey_GetComponentIndex(h.instance)
}

func (h *THotKey) SetComponentIndex(value int32) {
    HotKey_SetComponentIndex(h.instance, value)
}

func (h *THotKey) Owner() *TComponent {
    return ComponentFromInst(HotKey_GetOwner(h.instance))
}

func (h *THotKey) Name() string {
    return HotKey_GetName(h.instance)
}

func (h *THotKey) SetName(value string) {
    HotKey_SetName(h.instance, value)
}

func (h *THotKey) Tag() int {
    return HotKey_GetTag(h.instance)
}

func (h *THotKey) SetTag(value int) {
    HotKey_SetTag(h.instance, value)
}

func (h *THotKey) Controls(Index int32) *TControl {
    return ControlFromInst(HotKey_GetControls(h.instance, Index))
}

func (h *THotKey) Components(AIndex int32) *TComponent {
    return ComponentFromInst(HotKey_GetComponents(h.instance, AIndex))
}

