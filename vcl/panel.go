
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

type TPanel struct {
    IControl
    instance uintptr
}

func NewPanel(owner IComponent) *TPanel {
    p := new(TPanel)
    p.instance = Panel_Create(CheckPtr(owner))
    return p
}

func PanelFromInst(inst uintptr) *TPanel {
    p := new(TPanel)
    p.instance = inst
    return p
}

func PanelFromObj(obj IObject) *TPanel {
    p := new(TPanel)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPanel) Free() {
    if p.instance != 0 {
        Panel_Free(p.instance)
        p.instance = 0
    }
}

func (p *TPanel) Instance() uintptr {
    return p.instance
}

func (p *TPanel) IsValid() bool {
    return p.instance != 0
}

func (p *TPanel) CanFocus() bool {
    return Panel_CanFocus(p.instance)
}

func (p *TPanel) FlipChildren(AllLevels bool) {
    Panel_FlipChildren(p.instance, AllLevels)
}

func (p *TPanel) Focused() bool {
    return Panel_Focused(p.instance)
}

func (p *TPanel) HandleAllocated() bool {
    return Panel_HandleAllocated(p.instance)
}

func (p *TPanel) Invalidate() {
    Panel_Invalidate(p.instance)
}

func (p *TPanel) Realign() {
    Panel_Realign(p.instance)
}

func (p *TPanel) Repaint() {
    Panel_Repaint(p.instance)
}

func (p *TPanel) ScaleBy(M int32, D int32) {
    Panel_ScaleBy(p.instance, M , D)
}

func (p *TPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Panel_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

func (p *TPanel) SetFocus() {
    Panel_SetFocus(p.instance)
}

func (p *TPanel) Update() {
    Panel_Update(p.instance)
}

func (p *TPanel) BringToFront() {
    Panel_BringToFront(p.instance)
}

func (p *TPanel) HasParent() bool {
    return Panel_HasParent(p.instance)
}

func (p *TPanel) Hide() {
    Panel_Hide(p.instance)
}

func (p *TPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Panel_Perform(p.instance, Msg , WParam , LParam)
}

func (p *TPanel) Refresh() {
    Panel_Refresh(p.instance)
}

func (p *TPanel) SendToBack() {
    Panel_SendToBack(p.instance)
}

func (p *TPanel) Show() {
    Panel_Show(p.instance)
}

func (p *TPanel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Panel_GetTextBuf(p.instance, Buffer , BufSize)
}

func (p *TPanel) GetTextLen() int32 {
    return Panel_GetTextLen(p.instance)
}

func (p *TPanel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Panel_FindComponent(p.instance, AName))
}

func (p *TPanel) GetNamePath() string {
    return Panel_GetNamePath(p.instance)
}

func (p *TPanel) Assign(Source IObject) {
    Panel_Assign(p.instance, CheckPtr(Source))
}

func (p *TPanel) ClassName() string {
    return Panel_ClassName(p.instance)
}

func (p *TPanel) Equals(Obj IObject) bool {
    return Panel_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPanel) GetHashCode() int32 {
    return Panel_GetHashCode(p.instance)
}

func (p *TPanel) ToString() string {
    return Panel_ToString(p.instance)
}

func (p *TPanel) Align() TAlign {
    return Panel_GetAlign(p.instance)
}

func (p *TPanel) SetAlign(value TAlign) {
    Panel_SetAlign(p.instance, value)
}

func (p *TPanel) Alignment() TAlignment {
    return Panel_GetAlignment(p.instance)
}

func (p *TPanel) SetAlignment(value TAlignment) {
    Panel_SetAlignment(p.instance, value)
}

func (p *TPanel) Anchors() TAnchors {
    return Panel_GetAnchors(p.instance)
}

func (p *TPanel) SetAnchors(value TAnchors) {
    Panel_SetAnchors(p.instance, value)
}

func (p *TPanel) AutoSize() bool {
    return Panel_GetAutoSize(p.instance)
}

func (p *TPanel) SetAutoSize(value bool) {
    Panel_SetAutoSize(p.instance, value)
}

func (p *TPanel) BevelEdges() TBevelEdges {
    return Panel_GetBevelEdges(p.instance)
}

func (p *TPanel) SetBevelEdges(value TBevelEdges) {
    Panel_SetBevelEdges(p.instance, value)
}

func (p *TPanel) BevelInner() TBevelCut {
    return Panel_GetBevelInner(p.instance)
}

func (p *TPanel) SetBevelInner(value TBevelCut) {
    Panel_SetBevelInner(p.instance, value)
}

func (p *TPanel) BevelKind() TBevelKind {
    return Panel_GetBevelKind(p.instance)
}

func (p *TPanel) SetBevelKind(value TBevelKind) {
    Panel_SetBevelKind(p.instance, value)
}

func (p *TPanel) BevelOuter() TBevelCut {
    return Panel_GetBevelOuter(p.instance)
}

func (p *TPanel) SetBevelOuter(value TBevelCut) {
    Panel_SetBevelOuter(p.instance, value)
}

func (p *TPanel) BiDiMode() TBiDiMode {
    return Panel_GetBiDiMode(p.instance)
}

func (p *TPanel) SetBiDiMode(value TBiDiMode) {
    Panel_SetBiDiMode(p.instance, value)
}

func (p *TPanel) BorderWidth() int32 {
    return Panel_GetBorderWidth(p.instance)
}

func (p *TPanel) SetBorderWidth(value int32) {
    Panel_SetBorderWidth(p.instance, value)
}

func (p *TPanel) BorderStyle() TBorderStyle {
    return Panel_GetBorderStyle(p.instance)
}

func (p *TPanel) SetBorderStyle(value TBorderStyle) {
    Panel_SetBorderStyle(p.instance, value)
}

func (p *TPanel) Caption() string {
    return Panel_GetCaption(p.instance)
}

func (p *TPanel) SetCaption(value string) {
    Panel_SetCaption(p.instance, value)
}

func (p *TPanel) Color() TColor {
    return Panel_GetColor(p.instance)
}

func (p *TPanel) SetColor(value TColor) {
    Panel_SetColor(p.instance, value)
}

func (p *TPanel) DoubleBuffered() bool {
    return Panel_GetDoubleBuffered(p.instance)
}

func (p *TPanel) SetDoubleBuffered(value bool) {
    Panel_SetDoubleBuffered(p.instance, value)
}

func (p *TPanel) Enabled() bool {
    return Panel_GetEnabled(p.instance)
}

func (p *TPanel) SetEnabled(value bool) {
    Panel_SetEnabled(p.instance, value)
}

func (p *TPanel) FullRepaint() bool {
    return Panel_GetFullRepaint(p.instance)
}

func (p *TPanel) SetFullRepaint(value bool) {
    Panel_SetFullRepaint(p.instance, value)
}

func (p *TPanel) Font() *TFont {
    return FontFromInst(Panel_GetFont(p.instance))
}

func (p *TPanel) SetFont(value *TFont) {
    Panel_SetFont(p.instance, CheckPtr(value))
}

func (p *TPanel) Locked() bool {
    return Panel_GetLocked(p.instance)
}

func (p *TPanel) SetLocked(value bool) {
    Panel_SetLocked(p.instance, value)
}

func (p *TPanel) ParentBackground() bool {
    return Panel_GetParentBackground(p.instance)
}

func (p *TPanel) SetParentBackground(value bool) {
    Panel_SetParentBackground(p.instance, value)
}

func (p *TPanel) ParentColor() bool {
    return Panel_GetParentColor(p.instance)
}

func (p *TPanel) SetParentColor(value bool) {
    Panel_SetParentColor(p.instance, value)
}

func (p *TPanel) ParentCtl3D() bool {
    return Panel_GetParentCtl3D(p.instance)
}

func (p *TPanel) SetParentCtl3D(value bool) {
    Panel_SetParentCtl3D(p.instance, value)
}

func (p *TPanel) ParentDoubleBuffered() bool {
    return Panel_GetParentDoubleBuffered(p.instance)
}

func (p *TPanel) SetParentDoubleBuffered(value bool) {
    Panel_SetParentDoubleBuffered(p.instance, value)
}

func (p *TPanel) ParentFont() bool {
    return Panel_GetParentFont(p.instance)
}

func (p *TPanel) SetParentFont(value bool) {
    Panel_SetParentFont(p.instance, value)
}

func (p *TPanel) ParentShowHint() bool {
    return Panel_GetParentShowHint(p.instance)
}

func (p *TPanel) SetParentShowHint(value bool) {
    Panel_SetParentShowHint(p.instance, value)
}

func (p *TPanel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Panel_GetPopupMenu(p.instance))
}

func (p *TPanel) SetPopupMenu(value IComponent) {
    Panel_SetPopupMenu(p.instance, CheckPtr(value))
}

func (p *TPanel) ShowCaption() bool {
    return Panel_GetShowCaption(p.instance)
}

func (p *TPanel) SetShowCaption(value bool) {
    Panel_SetShowCaption(p.instance, value)
}

func (p *TPanel) ShowHint() bool {
    return Panel_GetShowHint(p.instance)
}

func (p *TPanel) SetShowHint(value bool) {
    Panel_SetShowHint(p.instance, value)
}

func (p *TPanel) TabOrder() uint16 {
    return Panel_GetTabOrder(p.instance)
}

func (p *TPanel) SetTabOrder(value uint16) {
    Panel_SetTabOrder(p.instance, value)
}

func (p *TPanel) TabStop() bool {
    return Panel_GetTabStop(p.instance)
}

func (p *TPanel) SetTabStop(value bool) {
    Panel_SetTabStop(p.instance, value)
}

func (p *TPanel) Visible() bool {
    return Panel_GetVisible(p.instance)
}

func (p *TPanel) SetVisible(value bool) {
    Panel_SetVisible(p.instance, value)
}

func (p *TPanel) StyleElements() TStyleElements {
    return Panel_GetStyleElements(p.instance)
}

func (p *TPanel) SetStyleElements(value TStyleElements) {
    Panel_SetStyleElements(p.instance, value)
}

func (p *TPanel) SetOnClick(fn TNotifyEvent) {
    Panel_SetOnClick(p.instance, fn)
}

func (p *TPanel) SetOnDblClick(fn TNotifyEvent) {
    Panel_SetOnDblClick(p.instance, fn)
}

func (p *TPanel) SetOnEnter(fn TNotifyEvent) {
    Panel_SetOnEnter(p.instance, fn)
}

func (p *TPanel) SetOnExit(fn TNotifyEvent) {
    Panel_SetOnExit(p.instance, fn)
}

func (p *TPanel) SetOnMouseDown(fn TMouseEvent) {
    Panel_SetOnMouseDown(p.instance, fn)
}

func (p *TPanel) SetOnMouseEnter(fn TNotifyEvent) {
    Panel_SetOnMouseEnter(p.instance, fn)
}

func (p *TPanel) SetOnMouseLeave(fn TNotifyEvent) {
    Panel_SetOnMouseLeave(p.instance, fn)
}

func (p *TPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    Panel_SetOnMouseMove(p.instance, fn)
}

func (p *TPanel) SetOnMouseUp(fn TMouseEvent) {
    Panel_SetOnMouseUp(p.instance, fn)
}

func (p *TPanel) SetOnResize(fn TNotifyEvent) {
    Panel_SetOnResize(p.instance, fn)
}

func (p *TPanel) Brush() *TBrush {
    return BrushFromInst(Panel_GetBrush(p.instance))
}

func (p *TPanel) ControlCount() int32 {
    return Panel_GetControlCount(p.instance)
}

func (p *TPanel) Handle() HWND {
    return Panel_GetHandle(p.instance)
}

func (p *TPanel) ParentWindow() HWND {
    return Panel_GetParentWindow(p.instance)
}

func (p *TPanel) SetParentWindow(value HWND) {
    Panel_SetParentWindow(p.instance, value)
}

func (p *TPanel) Action() *TAction {
    return ActionFromInst(Panel_GetAction(p.instance))
}

func (p *TPanel) SetAction(value IComponent) {
    Panel_SetAction(p.instance, CheckPtr(value))
}

func (p *TPanel) BoundsRect() TRect {
    return Panel_GetBoundsRect(p.instance)
}

func (p *TPanel) SetBoundsRect(value TRect) {
    Panel_SetBoundsRect(p.instance, value)
}

func (p *TPanel) ClientHeight() int32 {
    return Panel_GetClientHeight(p.instance)
}

func (p *TPanel) SetClientHeight(value int32) {
    Panel_SetClientHeight(p.instance, value)
}

func (p *TPanel) ClientRect() TRect {
    return Panel_GetClientRect(p.instance)
}

func (p *TPanel) ClientWidth() int32 {
    return Panel_GetClientWidth(p.instance)
}

func (p *TPanel) SetClientWidth(value int32) {
    Panel_SetClientWidth(p.instance, value)
}

func (p *TPanel) ExplicitLeft() int32 {
    return Panel_GetExplicitLeft(p.instance)
}

func (p *TPanel) ExplicitTop() int32 {
    return Panel_GetExplicitTop(p.instance)
}

func (p *TPanel) ExplicitWidth() int32 {
    return Panel_GetExplicitWidth(p.instance)
}

func (p *TPanel) ExplicitHeight() int32 {
    return Panel_GetExplicitHeight(p.instance)
}

func (p *TPanel) Parent() *TControl {
    return ControlFromInst(Panel_GetParent(p.instance))
}

func (p *TPanel) SetParent(value IControl) {
    Panel_SetParent(p.instance, CheckPtr(value))
}

func (p *TPanel) AlignWithMargins() bool {
    return Panel_GetAlignWithMargins(p.instance)
}

func (p *TPanel) SetAlignWithMargins(value bool) {
    Panel_SetAlignWithMargins(p.instance, value)
}

func (p *TPanel) Left() int32 {
    return Panel_GetLeft(p.instance)
}

func (p *TPanel) SetLeft(value int32) {
    Panel_SetLeft(p.instance, value)
}

func (p *TPanel) Top() int32 {
    return Panel_GetTop(p.instance)
}

func (p *TPanel) SetTop(value int32) {
    Panel_SetTop(p.instance, value)
}

func (p *TPanel) Width() int32 {
    return Panel_GetWidth(p.instance)
}

func (p *TPanel) SetWidth(value int32) {
    Panel_SetWidth(p.instance, value)
}

func (p *TPanel) Height() int32 {
    return Panel_GetHeight(p.instance)
}

func (p *TPanel) SetHeight(value int32) {
    Panel_SetHeight(p.instance, value)
}

func (p *TPanel) Cursor() TCursor {
    return Panel_GetCursor(p.instance)
}

func (p *TPanel) SetCursor(value TCursor) {
    Panel_SetCursor(p.instance, value)
}

func (p *TPanel) Hint() string {
    return Panel_GetHint(p.instance)
}

func (p *TPanel) SetHint(value string) {
    Panel_SetHint(p.instance, value)
}

func (p *TPanel) Margins() *TMargins {
    return MarginsFromInst(Panel_GetMargins(p.instance))
}

func (p *TPanel) SetMargins(value *TMargins) {
    Panel_SetMargins(p.instance, CheckPtr(value))
}

func (p *TPanel) CustomHint() *TCustomHint {
    return CustomHintFromInst(Panel_GetCustomHint(p.instance))
}

func (p *TPanel) SetCustomHint(value IComponent) {
    Panel_SetCustomHint(p.instance, CheckPtr(value))
}

func (p *TPanel) ComponentCount() int32 {
    return Panel_GetComponentCount(p.instance)
}

func (p *TPanel) ComponentIndex() int32 {
    return Panel_GetComponentIndex(p.instance)
}

func (p *TPanel) SetComponentIndex(value int32) {
    Panel_SetComponentIndex(p.instance, value)
}

func (p *TPanel) Owner() *TComponent {
    return ComponentFromInst(Panel_GetOwner(p.instance))
}

func (p *TPanel) Name() string {
    return Panel_GetName(p.instance)
}

func (p *TPanel) SetName(value string) {
    Panel_SetName(p.instance, value)
}

func (p *TPanel) Tag() int {
    return Panel_GetTag(p.instance)
}

func (p *TPanel) SetTag(value int) {
    Panel_SetTag(p.instance, value)
}

func (p *TPanel) Controls(Index int32) *TControl {
    return ControlFromInst(Panel_GetControls(p.instance, Index))
}

func (p *TPanel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Panel_GetComponents(p.instance, AIndex))
}

