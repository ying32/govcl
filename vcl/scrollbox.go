
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

type TScrollBox struct {
    IControl
    instance uintptr
}

func NewScrollBox(owner IComponent) *TScrollBox {
    s := new(TScrollBox)
    s.instance = ScrollBox_Create(CheckPtr(owner))
    return s
}

func ScrollBoxFromInst(inst uintptr) *TScrollBox {
    s := new(TScrollBox)
    s.instance = inst
    return s
}

func ScrollBoxFromObj(obj IObject) *TScrollBox {
    s := new(TScrollBox)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TScrollBox) Free() {
    if s.instance != 0 {
        ScrollBox_Free(s.instance)
        s.instance = 0
    }
}

func (s *TScrollBox) Instance() uintptr {
    return s.instance
}

func (s *TScrollBox) IsValid() bool {
    return s.instance != 0
}

func (s *TScrollBox) CanFocus() bool {
    return ScrollBox_CanFocus(s.instance)
}

func (s *TScrollBox) FlipChildren(AllLevels bool) {
    ScrollBox_FlipChildren(s.instance, AllLevels)
}

func (s *TScrollBox) Focused() bool {
    return ScrollBox_Focused(s.instance)
}

func (s *TScrollBox) HandleAllocated() bool {
    return ScrollBox_HandleAllocated(s.instance)
}

func (s *TScrollBox) Invalidate() {
    ScrollBox_Invalidate(s.instance)
}

func (s *TScrollBox) Realign() {
    ScrollBox_Realign(s.instance)
}

func (s *TScrollBox) Repaint() {
    ScrollBox_Repaint(s.instance)
}

func (s *TScrollBox) ScaleBy(M int32, D int32) {
    ScrollBox_ScaleBy(s.instance, M , D)
}

func (s *TScrollBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ScrollBox_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

func (s *TScrollBox) SetFocus() {
    ScrollBox_SetFocus(s.instance)
}

func (s *TScrollBox) Update() {
    ScrollBox_Update(s.instance)
}

func (s *TScrollBox) BringToFront() {
    ScrollBox_BringToFront(s.instance)
}

func (s *TScrollBox) HasParent() bool {
    return ScrollBox_HasParent(s.instance)
}

func (s *TScrollBox) Hide() {
    ScrollBox_Hide(s.instance)
}

func (s *TScrollBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ScrollBox_Perform(s.instance, Msg , WParam , LParam)
}

func (s *TScrollBox) Refresh() {
    ScrollBox_Refresh(s.instance)
}

func (s *TScrollBox) SendToBack() {
    ScrollBox_SendToBack(s.instance)
}

func (s *TScrollBox) Show() {
    ScrollBox_Show(s.instance)
}

func (s *TScrollBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ScrollBox_GetTextBuf(s.instance, Buffer , BufSize)
}

func (s *TScrollBox) GetTextLen() int32 {
    return ScrollBox_GetTextLen(s.instance)
}

func (s *TScrollBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ScrollBox_FindComponent(s.instance, AName))
}

func (s *TScrollBox) GetNamePath() string {
    return ScrollBox_GetNamePath(s.instance)
}

func (s *TScrollBox) Assign(Source IObject) {
    ScrollBox_Assign(s.instance, CheckPtr(Source))
}

func (s *TScrollBox) ClassName() string {
    return ScrollBox_ClassName(s.instance)
}

func (s *TScrollBox) Equals(Obj IObject) bool {
    return ScrollBox_Equals(s.instance, CheckPtr(Obj))
}

func (s *TScrollBox) GetHashCode() int32 {
    return ScrollBox_GetHashCode(s.instance)
}

func (s *TScrollBox) ToString() string {
    return ScrollBox_ToString(s.instance)
}

func (s *TScrollBox) Align() TAlign {
    return ScrollBox_GetAlign(s.instance)
}

func (s *TScrollBox) SetAlign(value TAlign) {
    ScrollBox_SetAlign(s.instance, value)
}

func (s *TScrollBox) Anchors() TAnchors {
    return ScrollBox_GetAnchors(s.instance)
}

func (s *TScrollBox) SetAnchors(value TAnchors) {
    ScrollBox_SetAnchors(s.instance, value)
}

func (s *TScrollBox) AutoSize() bool {
    return ScrollBox_GetAutoSize(s.instance)
}

func (s *TScrollBox) SetAutoSize(value bool) {
    ScrollBox_SetAutoSize(s.instance, value)
}

func (s *TScrollBox) BevelEdges() TBevelEdges {
    return ScrollBox_GetBevelEdges(s.instance)
}

func (s *TScrollBox) SetBevelEdges(value TBevelEdges) {
    ScrollBox_SetBevelEdges(s.instance, value)
}

func (s *TScrollBox) BevelInner() TBevelCut {
    return ScrollBox_GetBevelInner(s.instance)
}

func (s *TScrollBox) SetBevelInner(value TBevelCut) {
    ScrollBox_SetBevelInner(s.instance, value)
}

func (s *TScrollBox) BevelOuter() TBevelCut {
    return ScrollBox_GetBevelOuter(s.instance)
}

func (s *TScrollBox) SetBevelOuter(value TBevelCut) {
    ScrollBox_SetBevelOuter(s.instance, value)
}

func (s *TScrollBox) BevelKind() TBevelKind {
    return ScrollBox_GetBevelKind(s.instance)
}

func (s *TScrollBox) SetBevelKind(value TBevelKind) {
    ScrollBox_SetBevelKind(s.instance, value)
}

func (s *TScrollBox) BiDiMode() TBiDiMode {
    return ScrollBox_GetBiDiMode(s.instance)
}

func (s *TScrollBox) SetBiDiMode(value TBiDiMode) {
    ScrollBox_SetBiDiMode(s.instance, value)
}

func (s *TScrollBox) BorderStyle() TBorderStyle {
    return ScrollBox_GetBorderStyle(s.instance)
}

func (s *TScrollBox) SetBorderStyle(value TBorderStyle) {
    ScrollBox_SetBorderStyle(s.instance, value)
}

func (s *TScrollBox) DoubleBuffered() bool {
    return ScrollBox_GetDoubleBuffered(s.instance)
}

func (s *TScrollBox) SetDoubleBuffered(value bool) {
    ScrollBox_SetDoubleBuffered(s.instance, value)
}

func (s *TScrollBox) Enabled() bool {
    return ScrollBox_GetEnabled(s.instance)
}

func (s *TScrollBox) SetEnabled(value bool) {
    ScrollBox_SetEnabled(s.instance, value)
}

func (s *TScrollBox) Color() TColor {
    return ScrollBox_GetColor(s.instance)
}

func (s *TScrollBox) SetColor(value TColor) {
    ScrollBox_SetColor(s.instance, value)
}

func (s *TScrollBox) Font() *TFont {
    return FontFromInst(ScrollBox_GetFont(s.instance))
}

func (s *TScrollBox) SetFont(value *TFont) {
    ScrollBox_SetFont(s.instance, CheckPtr(value))
}

func (s *TScrollBox) ParentBackground() bool {
    return ScrollBox_GetParentBackground(s.instance)
}

func (s *TScrollBox) SetParentBackground(value bool) {
    ScrollBox_SetParentBackground(s.instance, value)
}

func (s *TScrollBox) ParentColor() bool {
    return ScrollBox_GetParentColor(s.instance)
}

func (s *TScrollBox) SetParentColor(value bool) {
    ScrollBox_SetParentColor(s.instance, value)
}

func (s *TScrollBox) ParentCtl3D() bool {
    return ScrollBox_GetParentCtl3D(s.instance)
}

func (s *TScrollBox) SetParentCtl3D(value bool) {
    ScrollBox_SetParentCtl3D(s.instance, value)
}

func (s *TScrollBox) ParentDoubleBuffered() bool {
    return ScrollBox_GetParentDoubleBuffered(s.instance)
}

func (s *TScrollBox) SetParentDoubleBuffered(value bool) {
    ScrollBox_SetParentDoubleBuffered(s.instance, value)
}

func (s *TScrollBox) ParentFont() bool {
    return ScrollBox_GetParentFont(s.instance)
}

func (s *TScrollBox) SetParentFont(value bool) {
    ScrollBox_SetParentFont(s.instance, value)
}

func (s *TScrollBox) ParentShowHint() bool {
    return ScrollBox_GetParentShowHint(s.instance)
}

func (s *TScrollBox) SetParentShowHint(value bool) {
    ScrollBox_SetParentShowHint(s.instance, value)
}

func (s *TScrollBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ScrollBox_GetPopupMenu(s.instance))
}

func (s *TScrollBox) SetPopupMenu(value IComponent) {
    ScrollBox_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TScrollBox) ShowHint() bool {
    return ScrollBox_GetShowHint(s.instance)
}

func (s *TScrollBox) SetShowHint(value bool) {
    ScrollBox_SetShowHint(s.instance, value)
}

func (s *TScrollBox) TabOrder() uint16 {
    return ScrollBox_GetTabOrder(s.instance)
}

func (s *TScrollBox) SetTabOrder(value uint16) {
    ScrollBox_SetTabOrder(s.instance, value)
}

func (s *TScrollBox) TabStop() bool {
    return ScrollBox_GetTabStop(s.instance)
}

func (s *TScrollBox) SetTabStop(value bool) {
    ScrollBox_SetTabStop(s.instance, value)
}

func (s *TScrollBox) Visible() bool {
    return ScrollBox_GetVisible(s.instance)
}

func (s *TScrollBox) SetVisible(value bool) {
    ScrollBox_SetVisible(s.instance, value)
}

func (s *TScrollBox) StyleElements() TStyleElements {
    return ScrollBox_GetStyleElements(s.instance)
}

func (s *TScrollBox) SetStyleElements(value TStyleElements) {
    ScrollBox_SetStyleElements(s.instance, value)
}

func (s *TScrollBox) SetOnClick(fn TNotifyEvent) {
    ScrollBox_SetOnClick(s.instance, fn)
}

func (s *TScrollBox) SetOnDblClick(fn TNotifyEvent) {
    ScrollBox_SetOnDblClick(s.instance, fn)
}

func (s *TScrollBox) SetOnEnter(fn TNotifyEvent) {
    ScrollBox_SetOnEnter(s.instance, fn)
}

func (s *TScrollBox) SetOnExit(fn TNotifyEvent) {
    ScrollBox_SetOnExit(s.instance, fn)
}

func (s *TScrollBox) SetOnMouseDown(fn TMouseEvent) {
    ScrollBox_SetOnMouseDown(s.instance, fn)
}

func (s *TScrollBox) SetOnMouseEnter(fn TNotifyEvent) {
    ScrollBox_SetOnMouseEnter(s.instance, fn)
}

func (s *TScrollBox) SetOnMouseLeave(fn TNotifyEvent) {
    ScrollBox_SetOnMouseLeave(s.instance, fn)
}

func (s *TScrollBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ScrollBox_SetOnMouseMove(s.instance, fn)
}

func (s *TScrollBox) SetOnMouseUp(fn TMouseEvent) {
    ScrollBox_SetOnMouseUp(s.instance, fn)
}

func (s *TScrollBox) SetOnMouseWheel(fn TMouseWheelEvent) {
    ScrollBox_SetOnMouseWheel(s.instance, fn)
}

func (s *TScrollBox) SetOnResize(fn TNotifyEvent) {
    ScrollBox_SetOnResize(s.instance, fn)
}

func (s *TScrollBox) Brush() *TBrush {
    return BrushFromInst(ScrollBox_GetBrush(s.instance))
}

func (s *TScrollBox) ControlCount() int32 {
    return ScrollBox_GetControlCount(s.instance)
}

func (s *TScrollBox) Handle() HWND {
    return ScrollBox_GetHandle(s.instance)
}

func (s *TScrollBox) ParentWindow() HWND {
    return ScrollBox_GetParentWindow(s.instance)
}

func (s *TScrollBox) SetParentWindow(value HWND) {
    ScrollBox_SetParentWindow(s.instance, value)
}

func (s *TScrollBox) Action() *TAction {
    return ActionFromInst(ScrollBox_GetAction(s.instance))
}

func (s *TScrollBox) SetAction(value IComponent) {
    ScrollBox_SetAction(s.instance, CheckPtr(value))
}

func (s *TScrollBox) BoundsRect() TRect {
    return ScrollBox_GetBoundsRect(s.instance)
}

func (s *TScrollBox) SetBoundsRect(value TRect) {
    ScrollBox_SetBoundsRect(s.instance, value)
}

func (s *TScrollBox) ClientHeight() int32 {
    return ScrollBox_GetClientHeight(s.instance)
}

func (s *TScrollBox) SetClientHeight(value int32) {
    ScrollBox_SetClientHeight(s.instance, value)
}

func (s *TScrollBox) ClientRect() TRect {
    return ScrollBox_GetClientRect(s.instance)
}

func (s *TScrollBox) ClientWidth() int32 {
    return ScrollBox_GetClientWidth(s.instance)
}

func (s *TScrollBox) SetClientWidth(value int32) {
    ScrollBox_SetClientWidth(s.instance, value)
}

func (s *TScrollBox) ExplicitLeft() int32 {
    return ScrollBox_GetExplicitLeft(s.instance)
}

func (s *TScrollBox) ExplicitTop() int32 {
    return ScrollBox_GetExplicitTop(s.instance)
}

func (s *TScrollBox) ExplicitWidth() int32 {
    return ScrollBox_GetExplicitWidth(s.instance)
}

func (s *TScrollBox) ExplicitHeight() int32 {
    return ScrollBox_GetExplicitHeight(s.instance)
}

func (s *TScrollBox) Parent() *TControl {
    return ControlFromInst(ScrollBox_GetParent(s.instance))
}

func (s *TScrollBox) SetParent(value IControl) {
    ScrollBox_SetParent(s.instance, CheckPtr(value))
}

func (s *TScrollBox) AlignWithMargins() bool {
    return ScrollBox_GetAlignWithMargins(s.instance)
}

func (s *TScrollBox) SetAlignWithMargins(value bool) {
    ScrollBox_SetAlignWithMargins(s.instance, value)
}

func (s *TScrollBox) Left() int32 {
    return ScrollBox_GetLeft(s.instance)
}

func (s *TScrollBox) SetLeft(value int32) {
    ScrollBox_SetLeft(s.instance, value)
}

func (s *TScrollBox) Top() int32 {
    return ScrollBox_GetTop(s.instance)
}

func (s *TScrollBox) SetTop(value int32) {
    ScrollBox_SetTop(s.instance, value)
}

func (s *TScrollBox) Width() int32 {
    return ScrollBox_GetWidth(s.instance)
}

func (s *TScrollBox) SetWidth(value int32) {
    ScrollBox_SetWidth(s.instance, value)
}

func (s *TScrollBox) Height() int32 {
    return ScrollBox_GetHeight(s.instance)
}

func (s *TScrollBox) SetHeight(value int32) {
    ScrollBox_SetHeight(s.instance, value)
}

func (s *TScrollBox) Cursor() TCursor {
    return ScrollBox_GetCursor(s.instance)
}

func (s *TScrollBox) SetCursor(value TCursor) {
    ScrollBox_SetCursor(s.instance, value)
}

func (s *TScrollBox) Hint() string {
    return ScrollBox_GetHint(s.instance)
}

func (s *TScrollBox) SetHint(value string) {
    ScrollBox_SetHint(s.instance, value)
}

func (s *TScrollBox) Margins() *TMargins {
    return MarginsFromInst(ScrollBox_GetMargins(s.instance))
}

func (s *TScrollBox) SetMargins(value *TMargins) {
    ScrollBox_SetMargins(s.instance, CheckPtr(value))
}

func (s *TScrollBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ScrollBox_GetCustomHint(s.instance))
}

func (s *TScrollBox) SetCustomHint(value IComponent) {
    ScrollBox_SetCustomHint(s.instance, CheckPtr(value))
}

func (s *TScrollBox) ComponentCount() int32 {
    return ScrollBox_GetComponentCount(s.instance)
}

func (s *TScrollBox) ComponentIndex() int32 {
    return ScrollBox_GetComponentIndex(s.instance)
}

func (s *TScrollBox) SetComponentIndex(value int32) {
    ScrollBox_SetComponentIndex(s.instance, value)
}

func (s *TScrollBox) Owner() *TComponent {
    return ComponentFromInst(ScrollBox_GetOwner(s.instance))
}

func (s *TScrollBox) Name() string {
    return ScrollBox_GetName(s.instance)
}

func (s *TScrollBox) SetName(value string) {
    ScrollBox_SetName(s.instance, value)
}

func (s *TScrollBox) Tag() int {
    return ScrollBox_GetTag(s.instance)
}

func (s *TScrollBox) SetTag(value int) {
    ScrollBox_SetTag(s.instance, value)
}

func (s *TScrollBox) Controls(Index int32) *TControl {
    return ControlFromInst(ScrollBox_GetControls(s.instance, Index))
}

func (s *TScrollBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ScrollBox_GetComponents(s.instance, AIndex))
}

