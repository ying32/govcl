
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

type TStatusBar struct {
    IControl
    instance uintptr
}

func NewStatusBar(owner IComponent) *TStatusBar {
    s := new(TStatusBar)
    s.instance = StatusBar_Create(CheckPtr(owner))
    return s
}

func StatusBarFromInst(inst uintptr) *TStatusBar {
    s := new(TStatusBar)
    s.instance = inst
    return s
}

func StatusBarFromObj(obj IObject) *TStatusBar {
    s := new(TStatusBar)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStatusBar) Free() {
    if s.instance != 0 {
        StatusBar_Free(s.instance)
        s.instance = 0
    }
}

func (s *TStatusBar) Instance() uintptr {
    return s.instance
}

func (s *TStatusBar) IsValid() bool {
    return s.instance != 0
}

func (s *TStatusBar) FlipChildren(AllLevels bool) {
    StatusBar_FlipChildren(s.instance, AllLevels)
}

func (s *TStatusBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    StatusBar_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

func (s *TStatusBar) CanFocus() bool {
    return StatusBar_CanFocus(s.instance)
}

func (s *TStatusBar) Focused() bool {
    return StatusBar_Focused(s.instance)
}

func (s *TStatusBar) HandleAllocated() bool {
    return StatusBar_HandleAllocated(s.instance)
}

func (s *TStatusBar) Invalidate() {
    StatusBar_Invalidate(s.instance)
}

func (s *TStatusBar) Realign() {
    StatusBar_Realign(s.instance)
}

func (s *TStatusBar) Repaint() {
    StatusBar_Repaint(s.instance)
}

func (s *TStatusBar) ScaleBy(M int32, D int32) {
    StatusBar_ScaleBy(s.instance, M , D)
}

func (s *TStatusBar) SetFocus() {
    StatusBar_SetFocus(s.instance)
}

func (s *TStatusBar) Update() {
    StatusBar_Update(s.instance)
}

func (s *TStatusBar) BringToFront() {
    StatusBar_BringToFront(s.instance)
}

func (s *TStatusBar) HasParent() bool {
    return StatusBar_HasParent(s.instance)
}

func (s *TStatusBar) Hide() {
    StatusBar_Hide(s.instance)
}

func (s *TStatusBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return StatusBar_Perform(s.instance, Msg , WParam , LParam)
}

func (s *TStatusBar) Refresh() {
    StatusBar_Refresh(s.instance)
}

func (s *TStatusBar) SendToBack() {
    StatusBar_SendToBack(s.instance)
}

func (s *TStatusBar) Show() {
    StatusBar_Show(s.instance)
}

func (s *TStatusBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return StatusBar_GetTextBuf(s.instance, Buffer , BufSize)
}

func (s *TStatusBar) GetTextLen() int32 {
    return StatusBar_GetTextLen(s.instance)
}

func (s *TStatusBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(StatusBar_FindComponent(s.instance, AName))
}

func (s *TStatusBar) GetNamePath() string {
    return StatusBar_GetNamePath(s.instance)
}

func (s *TStatusBar) Assign(Source IObject) {
    StatusBar_Assign(s.instance, CheckPtr(Source))
}

func (s *TStatusBar) ClassName() string {
    return StatusBar_ClassName(s.instance)
}

func (s *TStatusBar) Equals(Obj IObject) bool {
    return StatusBar_Equals(s.instance, CheckPtr(Obj))
}

func (s *TStatusBar) GetHashCode() int32 {
    return StatusBar_GetHashCode(s.instance)
}

func (s *TStatusBar) ToString() string {
    return StatusBar_ToString(s.instance)
}

func (s *TStatusBar) Action() *TAction {
    return ActionFromInst(StatusBar_GetAction(s.instance))
}

func (s *TStatusBar) SetAction(value IComponent) {
    StatusBar_SetAction(s.instance, CheckPtr(value))
}

func (s *TStatusBar) AutoHint() bool {
    return StatusBar_GetAutoHint(s.instance)
}

func (s *TStatusBar) SetAutoHint(value bool) {
    StatusBar_SetAutoHint(s.instance, value)
}

func (s *TStatusBar) Align() TAlign {
    return StatusBar_GetAlign(s.instance)
}

func (s *TStatusBar) SetAlign(value TAlign) {
    StatusBar_SetAlign(s.instance, value)
}

func (s *TStatusBar) Anchors() TAnchors {
    return StatusBar_GetAnchors(s.instance)
}

func (s *TStatusBar) SetAnchors(value TAnchors) {
    StatusBar_SetAnchors(s.instance, value)
}

func (s *TStatusBar) BiDiMode() TBiDiMode {
    return StatusBar_GetBiDiMode(s.instance)
}

func (s *TStatusBar) SetBiDiMode(value TBiDiMode) {
    StatusBar_SetBiDiMode(s.instance, value)
}

func (s *TStatusBar) BorderWidth() int32 {
    return StatusBar_GetBorderWidth(s.instance)
}

func (s *TStatusBar) SetBorderWidth(value int32) {
    StatusBar_SetBorderWidth(s.instance, value)
}

func (s *TStatusBar) Color() TColor {
    return StatusBar_GetColor(s.instance)
}

func (s *TStatusBar) SetColor(value TColor) {
    StatusBar_SetColor(s.instance, value)
}

func (s *TStatusBar) DoubleBuffered() bool {
    return StatusBar_GetDoubleBuffered(s.instance)
}

func (s *TStatusBar) SetDoubleBuffered(value bool) {
    StatusBar_SetDoubleBuffered(s.instance, value)
}

func (s *TStatusBar) Enabled() bool {
    return StatusBar_GetEnabled(s.instance)
}

func (s *TStatusBar) SetEnabled(value bool) {
    StatusBar_SetEnabled(s.instance, value)
}

func (s *TStatusBar) Font() *TFont {
    return FontFromInst(StatusBar_GetFont(s.instance))
}

func (s *TStatusBar) SetFont(value *TFont) {
    StatusBar_SetFont(s.instance, CheckPtr(value))
}

func (s *TStatusBar) Panels() *TStatusPanels {
    return StatusPanelsFromInst(StatusBar_GetPanels(s.instance))
}

func (s *TStatusBar) SetPanels(value *TStatusPanels) {
    StatusBar_SetPanels(s.instance, CheckPtr(value))
}

func (s *TStatusBar) ParentColor() bool {
    return StatusBar_GetParentColor(s.instance)
}

func (s *TStatusBar) SetParentColor(value bool) {
    StatusBar_SetParentColor(s.instance, value)
}

func (s *TStatusBar) ParentDoubleBuffered() bool {
    return StatusBar_GetParentDoubleBuffered(s.instance)
}

func (s *TStatusBar) SetParentDoubleBuffered(value bool) {
    StatusBar_SetParentDoubleBuffered(s.instance, value)
}

func (s *TStatusBar) ParentFont() bool {
    return StatusBar_GetParentFont(s.instance)
}

func (s *TStatusBar) SetParentFont(value bool) {
    StatusBar_SetParentFont(s.instance, value)
}

func (s *TStatusBar) ParentShowHint() bool {
    return StatusBar_GetParentShowHint(s.instance)
}

func (s *TStatusBar) SetParentShowHint(value bool) {
    StatusBar_SetParentShowHint(s.instance, value)
}

func (s *TStatusBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(StatusBar_GetPopupMenu(s.instance))
}

func (s *TStatusBar) SetPopupMenu(value IComponent) {
    StatusBar_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TStatusBar) ShowHint() bool {
    return StatusBar_GetShowHint(s.instance)
}

func (s *TStatusBar) SetShowHint(value bool) {
    StatusBar_SetShowHint(s.instance, value)
}

func (s *TStatusBar) SimplePanel() bool {
    return StatusBar_GetSimplePanel(s.instance)
}

func (s *TStatusBar) SetSimplePanel(value bool) {
    StatusBar_SetSimplePanel(s.instance, value)
}

func (s *TStatusBar) SimpleText() string {
    return StatusBar_GetSimpleText(s.instance)
}

func (s *TStatusBar) SetSimpleText(value string) {
    StatusBar_SetSimpleText(s.instance, value)
}

func (s *TStatusBar) SizeGrip() bool {
    return StatusBar_GetSizeGrip(s.instance)
}

func (s *TStatusBar) SetSizeGrip(value bool) {
    StatusBar_SetSizeGrip(s.instance, value)
}

func (s *TStatusBar) UseSystemFont() bool {
    return StatusBar_GetUseSystemFont(s.instance)
}

func (s *TStatusBar) SetUseSystemFont(value bool) {
    StatusBar_SetUseSystemFont(s.instance, value)
}

func (s *TStatusBar) Visible() bool {
    return StatusBar_GetVisible(s.instance)
}

func (s *TStatusBar) SetVisible(value bool) {
    StatusBar_SetVisible(s.instance, value)
}

func (s *TStatusBar) StyleElements() TStyleElements {
    return StatusBar_GetStyleElements(s.instance)
}

func (s *TStatusBar) SetStyleElements(value TStyleElements) {
    StatusBar_SetStyleElements(s.instance, value)
}

func (s *TStatusBar) SetOnClick(fn TNotifyEvent) {
    StatusBar_SetOnClick(s.instance, fn)
}

func (s *TStatusBar) SetOnDblClick(fn TNotifyEvent) {
    StatusBar_SetOnDblClick(s.instance, fn)
}

func (s *TStatusBar) SetOnHint(fn TNotifyEvent) {
    StatusBar_SetOnHint(s.instance, fn)
}

func (s *TStatusBar) SetOnMouseDown(fn TMouseEvent) {
    StatusBar_SetOnMouseDown(s.instance, fn)
}

func (s *TStatusBar) SetOnMouseEnter(fn TNotifyEvent) {
    StatusBar_SetOnMouseEnter(s.instance, fn)
}

func (s *TStatusBar) SetOnMouseLeave(fn TNotifyEvent) {
    StatusBar_SetOnMouseLeave(s.instance, fn)
}

func (s *TStatusBar) SetOnMouseMove(fn TMouseMoveEvent) {
    StatusBar_SetOnMouseMove(s.instance, fn)
}

func (s *TStatusBar) SetOnMouseUp(fn TMouseEvent) {
    StatusBar_SetOnMouseUp(s.instance, fn)
}

func (s *TStatusBar) SetOnResize(fn TNotifyEvent) {
    StatusBar_SetOnResize(s.instance, fn)
}

func (s *TStatusBar) Canvas() *TCanvas {
    return CanvasFromInst(StatusBar_GetCanvas(s.instance))
}

func (s *TStatusBar) Brush() *TBrush {
    return BrushFromInst(StatusBar_GetBrush(s.instance))
}

func (s *TStatusBar) ControlCount() int32 {
    return StatusBar_GetControlCount(s.instance)
}

func (s *TStatusBar) Handle() HWND {
    return StatusBar_GetHandle(s.instance)
}

func (s *TStatusBar) ParentWindow() HWND {
    return StatusBar_GetParentWindow(s.instance)
}

func (s *TStatusBar) SetParentWindow(value HWND) {
    StatusBar_SetParentWindow(s.instance, value)
}

func (s *TStatusBar) TabOrder() uint16 {
    return StatusBar_GetTabOrder(s.instance)
}

func (s *TStatusBar) SetTabOrder(value uint16) {
    StatusBar_SetTabOrder(s.instance, value)
}

func (s *TStatusBar) TabStop() bool {
    return StatusBar_GetTabStop(s.instance)
}

func (s *TStatusBar) SetTabStop(value bool) {
    StatusBar_SetTabStop(s.instance, value)
}

func (s *TStatusBar) BoundsRect() TRect {
    return StatusBar_GetBoundsRect(s.instance)
}

func (s *TStatusBar) SetBoundsRect(value TRect) {
    StatusBar_SetBoundsRect(s.instance, value)
}

func (s *TStatusBar) ClientHeight() int32 {
    return StatusBar_GetClientHeight(s.instance)
}

func (s *TStatusBar) SetClientHeight(value int32) {
    StatusBar_SetClientHeight(s.instance, value)
}

func (s *TStatusBar) ClientRect() TRect {
    return StatusBar_GetClientRect(s.instance)
}

func (s *TStatusBar) ClientWidth() int32 {
    return StatusBar_GetClientWidth(s.instance)
}

func (s *TStatusBar) SetClientWidth(value int32) {
    StatusBar_SetClientWidth(s.instance, value)
}

func (s *TStatusBar) ExplicitLeft() int32 {
    return StatusBar_GetExplicitLeft(s.instance)
}

func (s *TStatusBar) ExplicitTop() int32 {
    return StatusBar_GetExplicitTop(s.instance)
}

func (s *TStatusBar) ExplicitWidth() int32 {
    return StatusBar_GetExplicitWidth(s.instance)
}

func (s *TStatusBar) ExplicitHeight() int32 {
    return StatusBar_GetExplicitHeight(s.instance)
}

func (s *TStatusBar) Parent() *TControl {
    return ControlFromInst(StatusBar_GetParent(s.instance))
}

func (s *TStatusBar) SetParent(value IControl) {
    StatusBar_SetParent(s.instance, CheckPtr(value))
}

func (s *TStatusBar) AlignWithMargins() bool {
    return StatusBar_GetAlignWithMargins(s.instance)
}

func (s *TStatusBar) SetAlignWithMargins(value bool) {
    StatusBar_SetAlignWithMargins(s.instance, value)
}

func (s *TStatusBar) Left() int32 {
    return StatusBar_GetLeft(s.instance)
}

func (s *TStatusBar) SetLeft(value int32) {
    StatusBar_SetLeft(s.instance, value)
}

func (s *TStatusBar) Top() int32 {
    return StatusBar_GetTop(s.instance)
}

func (s *TStatusBar) SetTop(value int32) {
    StatusBar_SetTop(s.instance, value)
}

func (s *TStatusBar) Width() int32 {
    return StatusBar_GetWidth(s.instance)
}

func (s *TStatusBar) SetWidth(value int32) {
    StatusBar_SetWidth(s.instance, value)
}

func (s *TStatusBar) Height() int32 {
    return StatusBar_GetHeight(s.instance)
}

func (s *TStatusBar) SetHeight(value int32) {
    StatusBar_SetHeight(s.instance, value)
}

func (s *TStatusBar) Cursor() TCursor {
    return StatusBar_GetCursor(s.instance)
}

func (s *TStatusBar) SetCursor(value TCursor) {
    StatusBar_SetCursor(s.instance, value)
}

func (s *TStatusBar) Hint() string {
    return StatusBar_GetHint(s.instance)
}

func (s *TStatusBar) SetHint(value string) {
    StatusBar_SetHint(s.instance, value)
}

func (s *TStatusBar) Margins() *TMargins {
    return MarginsFromInst(StatusBar_GetMargins(s.instance))
}

func (s *TStatusBar) SetMargins(value *TMargins) {
    StatusBar_SetMargins(s.instance, CheckPtr(value))
}

func (s *TStatusBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(StatusBar_GetCustomHint(s.instance))
}

func (s *TStatusBar) SetCustomHint(value IComponent) {
    StatusBar_SetCustomHint(s.instance, CheckPtr(value))
}

func (s *TStatusBar) ComponentCount() int32 {
    return StatusBar_GetComponentCount(s.instance)
}

func (s *TStatusBar) ComponentIndex() int32 {
    return StatusBar_GetComponentIndex(s.instance)
}

func (s *TStatusBar) SetComponentIndex(value int32) {
    StatusBar_SetComponentIndex(s.instance, value)
}

func (s *TStatusBar) Owner() *TComponent {
    return ComponentFromInst(StatusBar_GetOwner(s.instance))
}

func (s *TStatusBar) Name() string {
    return StatusBar_GetName(s.instance)
}

func (s *TStatusBar) SetName(value string) {
    StatusBar_SetName(s.instance, value)
}

func (s *TStatusBar) Tag() int {
    return StatusBar_GetTag(s.instance)
}

func (s *TStatusBar) SetTag(value int) {
    StatusBar_SetTag(s.instance, value)
}

func (s *TStatusBar) Controls(Index int32) *TControl {
    return ControlFromInst(StatusBar_GetControls(s.instance, Index))
}

func (s *TStatusBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(StatusBar_GetComponents(s.instance, AIndex))
}

