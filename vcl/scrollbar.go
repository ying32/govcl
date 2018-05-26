
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

type TScrollBar struct {
    IControl
    instance uintptr
}

func NewScrollBar(owner IComponent) *TScrollBar {
    s := new(TScrollBar)
    s.instance = ScrollBar_Create(CheckPtr(owner))
    return s
}

func ScrollBarFromInst(inst uintptr) *TScrollBar {
    s := new(TScrollBar)
    s.instance = inst
    return s
}

func ScrollBarFromObj(obj IObject) *TScrollBar {
    s := new(TScrollBar)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TScrollBar) Free() {
    if s.instance != 0 {
        ScrollBar_Free(s.instance)
        s.instance = 0
    }
}

func (s *TScrollBar) Instance() uintptr {
    return s.instance
}

func (s *TScrollBar) IsValid() bool {
    return s.instance != 0
}

func (s *TScrollBar) SetParams(APosition int32, AMin int32, AMax int32) {
    ScrollBar_SetParams(s.instance, APosition , AMin , AMax)
}

func (s *TScrollBar) CanFocus() bool {
    return ScrollBar_CanFocus(s.instance)
}

func (s *TScrollBar) FlipChildren(AllLevels bool) {
    ScrollBar_FlipChildren(s.instance, AllLevels)
}

func (s *TScrollBar) Focused() bool {
    return ScrollBar_Focused(s.instance)
}

func (s *TScrollBar) HandleAllocated() bool {
    return ScrollBar_HandleAllocated(s.instance)
}

func (s *TScrollBar) Invalidate() {
    ScrollBar_Invalidate(s.instance)
}

func (s *TScrollBar) Realign() {
    ScrollBar_Realign(s.instance)
}

func (s *TScrollBar) Repaint() {
    ScrollBar_Repaint(s.instance)
}

func (s *TScrollBar) ScaleBy(M int32, D int32) {
    ScrollBar_ScaleBy(s.instance, M , D)
}

func (s *TScrollBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ScrollBar_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

func (s *TScrollBar) SetFocus() {
    ScrollBar_SetFocus(s.instance)
}

func (s *TScrollBar) Update() {
    ScrollBar_Update(s.instance)
}

func (s *TScrollBar) BringToFront() {
    ScrollBar_BringToFront(s.instance)
}

func (s *TScrollBar) HasParent() bool {
    return ScrollBar_HasParent(s.instance)
}

func (s *TScrollBar) Hide() {
    ScrollBar_Hide(s.instance)
}

func (s *TScrollBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ScrollBar_Perform(s.instance, Msg , WParam , LParam)
}

func (s *TScrollBar) Refresh() {
    ScrollBar_Refresh(s.instance)
}

func (s *TScrollBar) SendToBack() {
    ScrollBar_SendToBack(s.instance)
}

func (s *TScrollBar) Show() {
    ScrollBar_Show(s.instance)
}

func (s *TScrollBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ScrollBar_GetTextBuf(s.instance, Buffer , BufSize)
}

func (s *TScrollBar) GetTextLen() int32 {
    return ScrollBar_GetTextLen(s.instance)
}

func (s *TScrollBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ScrollBar_FindComponent(s.instance, AName))
}

func (s *TScrollBar) GetNamePath() string {
    return ScrollBar_GetNamePath(s.instance)
}

func (s *TScrollBar) Assign(Source IObject) {
    ScrollBar_Assign(s.instance, CheckPtr(Source))
}

func (s *TScrollBar) ClassName() string {
    return ScrollBar_ClassName(s.instance)
}

func (s *TScrollBar) Equals(Obj IObject) bool {
    return ScrollBar_Equals(s.instance, CheckPtr(Obj))
}

func (s *TScrollBar) GetHashCode() int32 {
    return ScrollBar_GetHashCode(s.instance)
}

func (s *TScrollBar) ToString() string {
    return ScrollBar_ToString(s.instance)
}

func (s *TScrollBar) Align() TAlign {
    return ScrollBar_GetAlign(s.instance)
}

func (s *TScrollBar) SetAlign(value TAlign) {
    ScrollBar_SetAlign(s.instance, value)
}

func (s *TScrollBar) Anchors() TAnchors {
    return ScrollBar_GetAnchors(s.instance)
}

func (s *TScrollBar) SetAnchors(value TAnchors) {
    ScrollBar_SetAnchors(s.instance, value)
}

func (s *TScrollBar) BiDiMode() TBiDiMode {
    return ScrollBar_GetBiDiMode(s.instance)
}

func (s *TScrollBar) SetBiDiMode(value TBiDiMode) {
    ScrollBar_SetBiDiMode(s.instance, value)
}

func (s *TScrollBar) DoubleBuffered() bool {
    return ScrollBar_GetDoubleBuffered(s.instance)
}

func (s *TScrollBar) SetDoubleBuffered(value bool) {
    ScrollBar_SetDoubleBuffered(s.instance, value)
}

func (s *TScrollBar) Enabled() bool {
    return ScrollBar_GetEnabled(s.instance)
}

func (s *TScrollBar) SetEnabled(value bool) {
    ScrollBar_SetEnabled(s.instance, value)
}

func (s *TScrollBar) Kind() TScrollBarKind {
    return ScrollBar_GetKind(s.instance)
}

func (s *TScrollBar) SetKind(value TScrollBarKind) {
    ScrollBar_SetKind(s.instance, value)
}

func (s *TScrollBar) LargeChange() TScrollBarInc {
    return ScrollBar_GetLargeChange(s.instance)
}

func (s *TScrollBar) SetLargeChange(value TScrollBarInc) {
    ScrollBar_SetLargeChange(s.instance, value)
}

func (s *TScrollBar) Max() int32 {
    return ScrollBar_GetMax(s.instance)
}

func (s *TScrollBar) SetMax(value int32) {
    ScrollBar_SetMax(s.instance, value)
}

func (s *TScrollBar) Min() int32 {
    return ScrollBar_GetMin(s.instance)
}

func (s *TScrollBar) SetMin(value int32) {
    ScrollBar_SetMin(s.instance, value)
}

func (s *TScrollBar) PageSize() int32 {
    return ScrollBar_GetPageSize(s.instance)
}

func (s *TScrollBar) SetPageSize(value int32) {
    ScrollBar_SetPageSize(s.instance, value)
}

func (s *TScrollBar) ParentCtl3D() bool {
    return ScrollBar_GetParentCtl3D(s.instance)
}

func (s *TScrollBar) SetParentCtl3D(value bool) {
    ScrollBar_SetParentCtl3D(s.instance, value)
}

func (s *TScrollBar) ParentDoubleBuffered() bool {
    return ScrollBar_GetParentDoubleBuffered(s.instance)
}

func (s *TScrollBar) SetParentDoubleBuffered(value bool) {
    ScrollBar_SetParentDoubleBuffered(s.instance, value)
}

func (s *TScrollBar) ParentShowHint() bool {
    return ScrollBar_GetParentShowHint(s.instance)
}

func (s *TScrollBar) SetParentShowHint(value bool) {
    ScrollBar_SetParentShowHint(s.instance, value)
}

func (s *TScrollBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ScrollBar_GetPopupMenu(s.instance))
}

func (s *TScrollBar) SetPopupMenu(value IComponent) {
    ScrollBar_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TScrollBar) Position() int32 {
    return ScrollBar_GetPosition(s.instance)
}

func (s *TScrollBar) SetPosition(value int32) {
    ScrollBar_SetPosition(s.instance, value)
}

func (s *TScrollBar) ShowHint() bool {
    return ScrollBar_GetShowHint(s.instance)
}

func (s *TScrollBar) SetShowHint(value bool) {
    ScrollBar_SetShowHint(s.instance, value)
}

func (s *TScrollBar) SmallChange() TScrollBarInc {
    return ScrollBar_GetSmallChange(s.instance)
}

func (s *TScrollBar) SetSmallChange(value TScrollBarInc) {
    ScrollBar_SetSmallChange(s.instance, value)
}

func (s *TScrollBar) TabOrder() uint16 {
    return ScrollBar_GetTabOrder(s.instance)
}

func (s *TScrollBar) SetTabOrder(value uint16) {
    ScrollBar_SetTabOrder(s.instance, value)
}

func (s *TScrollBar) TabStop() bool {
    return ScrollBar_GetTabStop(s.instance)
}

func (s *TScrollBar) SetTabStop(value bool) {
    ScrollBar_SetTabStop(s.instance, value)
}

func (s *TScrollBar) Visible() bool {
    return ScrollBar_GetVisible(s.instance)
}

func (s *TScrollBar) SetVisible(value bool) {
    ScrollBar_SetVisible(s.instance, value)
}

func (s *TScrollBar) StyleElements() TStyleElements {
    return ScrollBar_GetStyleElements(s.instance)
}

func (s *TScrollBar) SetStyleElements(value TStyleElements) {
    ScrollBar_SetStyleElements(s.instance, value)
}

func (s *TScrollBar) SetOnChange(fn TNotifyEvent) {
    ScrollBar_SetOnChange(s.instance, fn)
}

func (s *TScrollBar) SetOnEnter(fn TNotifyEvent) {
    ScrollBar_SetOnEnter(s.instance, fn)
}

func (s *TScrollBar) SetOnExit(fn TNotifyEvent) {
    ScrollBar_SetOnExit(s.instance, fn)
}

func (s *TScrollBar) SetOnKeyDown(fn TKeyEvent) {
    ScrollBar_SetOnKeyDown(s.instance, fn)
}

func (s *TScrollBar) SetOnKeyPress(fn TKeyPressEvent) {
    ScrollBar_SetOnKeyPress(s.instance, fn)
}

func (s *TScrollBar) SetOnKeyUp(fn TKeyEvent) {
    ScrollBar_SetOnKeyUp(s.instance, fn)
}

func (s *TScrollBar) SetOnMouseEnter(fn TNotifyEvent) {
    ScrollBar_SetOnMouseEnter(s.instance, fn)
}

func (s *TScrollBar) SetOnMouseLeave(fn TNotifyEvent) {
    ScrollBar_SetOnMouseLeave(s.instance, fn)
}

func (s *TScrollBar) Brush() *TBrush {
    return BrushFromInst(ScrollBar_GetBrush(s.instance))
}

func (s *TScrollBar) ControlCount() int32 {
    return ScrollBar_GetControlCount(s.instance)
}

func (s *TScrollBar) Handle() HWND {
    return ScrollBar_GetHandle(s.instance)
}

func (s *TScrollBar) ParentWindow() HWND {
    return ScrollBar_GetParentWindow(s.instance)
}

func (s *TScrollBar) SetParentWindow(value HWND) {
    ScrollBar_SetParentWindow(s.instance, value)
}

func (s *TScrollBar) Action() *TAction {
    return ActionFromInst(ScrollBar_GetAction(s.instance))
}

func (s *TScrollBar) SetAction(value IComponent) {
    ScrollBar_SetAction(s.instance, CheckPtr(value))
}

func (s *TScrollBar) BoundsRect() TRect {
    return ScrollBar_GetBoundsRect(s.instance)
}

func (s *TScrollBar) SetBoundsRect(value TRect) {
    ScrollBar_SetBoundsRect(s.instance, value)
}

func (s *TScrollBar) ClientHeight() int32 {
    return ScrollBar_GetClientHeight(s.instance)
}

func (s *TScrollBar) SetClientHeight(value int32) {
    ScrollBar_SetClientHeight(s.instance, value)
}

func (s *TScrollBar) ClientRect() TRect {
    return ScrollBar_GetClientRect(s.instance)
}

func (s *TScrollBar) ClientWidth() int32 {
    return ScrollBar_GetClientWidth(s.instance)
}

func (s *TScrollBar) SetClientWidth(value int32) {
    ScrollBar_SetClientWidth(s.instance, value)
}

func (s *TScrollBar) ExplicitLeft() int32 {
    return ScrollBar_GetExplicitLeft(s.instance)
}

func (s *TScrollBar) ExplicitTop() int32 {
    return ScrollBar_GetExplicitTop(s.instance)
}

func (s *TScrollBar) ExplicitWidth() int32 {
    return ScrollBar_GetExplicitWidth(s.instance)
}

func (s *TScrollBar) ExplicitHeight() int32 {
    return ScrollBar_GetExplicitHeight(s.instance)
}

func (s *TScrollBar) Parent() *TControl {
    return ControlFromInst(ScrollBar_GetParent(s.instance))
}

func (s *TScrollBar) SetParent(value IControl) {
    ScrollBar_SetParent(s.instance, CheckPtr(value))
}

func (s *TScrollBar) AlignWithMargins() bool {
    return ScrollBar_GetAlignWithMargins(s.instance)
}

func (s *TScrollBar) SetAlignWithMargins(value bool) {
    ScrollBar_SetAlignWithMargins(s.instance, value)
}

func (s *TScrollBar) Left() int32 {
    return ScrollBar_GetLeft(s.instance)
}

func (s *TScrollBar) SetLeft(value int32) {
    ScrollBar_SetLeft(s.instance, value)
}

func (s *TScrollBar) Top() int32 {
    return ScrollBar_GetTop(s.instance)
}

func (s *TScrollBar) SetTop(value int32) {
    ScrollBar_SetTop(s.instance, value)
}

func (s *TScrollBar) Width() int32 {
    return ScrollBar_GetWidth(s.instance)
}

func (s *TScrollBar) SetWidth(value int32) {
    ScrollBar_SetWidth(s.instance, value)
}

func (s *TScrollBar) Height() int32 {
    return ScrollBar_GetHeight(s.instance)
}

func (s *TScrollBar) SetHeight(value int32) {
    ScrollBar_SetHeight(s.instance, value)
}

func (s *TScrollBar) Cursor() TCursor {
    return ScrollBar_GetCursor(s.instance)
}

func (s *TScrollBar) SetCursor(value TCursor) {
    ScrollBar_SetCursor(s.instance, value)
}

func (s *TScrollBar) Hint() string {
    return ScrollBar_GetHint(s.instance)
}

func (s *TScrollBar) SetHint(value string) {
    ScrollBar_SetHint(s.instance, value)
}

func (s *TScrollBar) Margins() *TMargins {
    return MarginsFromInst(ScrollBar_GetMargins(s.instance))
}

func (s *TScrollBar) SetMargins(value *TMargins) {
    ScrollBar_SetMargins(s.instance, CheckPtr(value))
}

func (s *TScrollBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(ScrollBar_GetCustomHint(s.instance))
}

func (s *TScrollBar) SetCustomHint(value IComponent) {
    ScrollBar_SetCustomHint(s.instance, CheckPtr(value))
}

func (s *TScrollBar) ComponentCount() int32 {
    return ScrollBar_GetComponentCount(s.instance)
}

func (s *TScrollBar) ComponentIndex() int32 {
    return ScrollBar_GetComponentIndex(s.instance)
}

func (s *TScrollBar) SetComponentIndex(value int32) {
    ScrollBar_SetComponentIndex(s.instance, value)
}

func (s *TScrollBar) Owner() *TComponent {
    return ComponentFromInst(ScrollBar_GetOwner(s.instance))
}

func (s *TScrollBar) Name() string {
    return ScrollBar_GetName(s.instance)
}

func (s *TScrollBar) SetName(value string) {
    ScrollBar_SetName(s.instance, value)
}

func (s *TScrollBar) Tag() int {
    return ScrollBar_GetTag(s.instance)
}

func (s *TScrollBar) SetTag(value int) {
    ScrollBar_SetTag(s.instance, value)
}

func (s *TScrollBar) Controls(Index int32) *TControl {
    return ControlFromInst(ScrollBar_GetControls(s.instance, Index))
}

func (s *TScrollBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ScrollBar_GetComponents(s.instance, AIndex))
}

