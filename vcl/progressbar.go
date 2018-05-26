
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

type TProgressBar struct {
    IControl
    instance uintptr
}

func NewProgressBar(owner IComponent) *TProgressBar {
    p := new(TProgressBar)
    p.instance = ProgressBar_Create(CheckPtr(owner))
    return p
}

func ProgressBarFromInst(inst uintptr) *TProgressBar {
    p := new(TProgressBar)
    p.instance = inst
    return p
}

func ProgressBarFromObj(obj IObject) *TProgressBar {
    p := new(TProgressBar)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TProgressBar) Free() {
    if p.instance != 0 {
        ProgressBar_Free(p.instance)
        p.instance = 0
    }
}

func (p *TProgressBar) Instance() uintptr {
    return p.instance
}

func (p *TProgressBar) IsValid() bool {
    return p.instance != 0
}

func (p *TProgressBar) StepIt() {
    ProgressBar_StepIt(p.instance)
}

func (p *TProgressBar) StepBy(Delta int32) {
    ProgressBar_StepBy(p.instance, Delta)
}

func (p *TProgressBar) CanFocus() bool {
    return ProgressBar_CanFocus(p.instance)
}

func (p *TProgressBar) FlipChildren(AllLevels bool) {
    ProgressBar_FlipChildren(p.instance, AllLevels)
}

func (p *TProgressBar) Focused() bool {
    return ProgressBar_Focused(p.instance)
}

func (p *TProgressBar) HandleAllocated() bool {
    return ProgressBar_HandleAllocated(p.instance)
}

func (p *TProgressBar) Invalidate() {
    ProgressBar_Invalidate(p.instance)
}

func (p *TProgressBar) Realign() {
    ProgressBar_Realign(p.instance)
}

func (p *TProgressBar) Repaint() {
    ProgressBar_Repaint(p.instance)
}

func (p *TProgressBar) ScaleBy(M int32, D int32) {
    ProgressBar_ScaleBy(p.instance, M , D)
}

func (p *TProgressBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ProgressBar_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

func (p *TProgressBar) SetFocus() {
    ProgressBar_SetFocus(p.instance)
}

func (p *TProgressBar) Update() {
    ProgressBar_Update(p.instance)
}

func (p *TProgressBar) BringToFront() {
    ProgressBar_BringToFront(p.instance)
}

func (p *TProgressBar) HasParent() bool {
    return ProgressBar_HasParent(p.instance)
}

func (p *TProgressBar) Hide() {
    ProgressBar_Hide(p.instance)
}

func (p *TProgressBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ProgressBar_Perform(p.instance, Msg , WParam , LParam)
}

func (p *TProgressBar) Refresh() {
    ProgressBar_Refresh(p.instance)
}

func (p *TProgressBar) SendToBack() {
    ProgressBar_SendToBack(p.instance)
}

func (p *TProgressBar) Show() {
    ProgressBar_Show(p.instance)
}

func (p *TProgressBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ProgressBar_GetTextBuf(p.instance, Buffer , BufSize)
}

func (p *TProgressBar) GetTextLen() int32 {
    return ProgressBar_GetTextLen(p.instance)
}

func (p *TProgressBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ProgressBar_FindComponent(p.instance, AName))
}

func (p *TProgressBar) GetNamePath() string {
    return ProgressBar_GetNamePath(p.instance)
}

func (p *TProgressBar) Assign(Source IObject) {
    ProgressBar_Assign(p.instance, CheckPtr(Source))
}

func (p *TProgressBar) ClassName() string {
    return ProgressBar_ClassName(p.instance)
}

func (p *TProgressBar) Equals(Obj IObject) bool {
    return ProgressBar_Equals(p.instance, CheckPtr(Obj))
}

func (p *TProgressBar) GetHashCode() int32 {
    return ProgressBar_GetHashCode(p.instance)
}

func (p *TProgressBar) ToString() string {
    return ProgressBar_ToString(p.instance)
}

func (p *TProgressBar) Align() TAlign {
    return ProgressBar_GetAlign(p.instance)
}

func (p *TProgressBar) SetAlign(value TAlign) {
    ProgressBar_SetAlign(p.instance, value)
}

func (p *TProgressBar) Anchors() TAnchors {
    return ProgressBar_GetAnchors(p.instance)
}

func (p *TProgressBar) SetAnchors(value TAnchors) {
    ProgressBar_SetAnchors(p.instance, value)
}

func (p *TProgressBar) BorderWidth() int32 {
    return ProgressBar_GetBorderWidth(p.instance)
}

func (p *TProgressBar) SetBorderWidth(value int32) {
    ProgressBar_SetBorderWidth(p.instance, value)
}

func (p *TProgressBar) DoubleBuffered() bool {
    return ProgressBar_GetDoubleBuffered(p.instance)
}

func (p *TProgressBar) SetDoubleBuffered(value bool) {
    ProgressBar_SetDoubleBuffered(p.instance, value)
}

func (p *TProgressBar) Enabled() bool {
    return ProgressBar_GetEnabled(p.instance)
}

func (p *TProgressBar) SetEnabled(value bool) {
    ProgressBar_SetEnabled(p.instance, value)
}

func (p *TProgressBar) Hint() string {
    return ProgressBar_GetHint(p.instance)
}

func (p *TProgressBar) SetHint(value string) {
    ProgressBar_SetHint(p.instance, value)
}

func (p *TProgressBar) Min() int32 {
    return ProgressBar_GetMin(p.instance)
}

func (p *TProgressBar) SetMin(value int32) {
    ProgressBar_SetMin(p.instance, value)
}

func (p *TProgressBar) Max() int32 {
    return ProgressBar_GetMax(p.instance)
}

func (p *TProgressBar) SetMax(value int32) {
    ProgressBar_SetMax(p.instance, value)
}

func (p *TProgressBar) Orientation() TProgressBarOrientation {
    return ProgressBar_GetOrientation(p.instance)
}

func (p *TProgressBar) SetOrientation(value TProgressBarOrientation) {
    ProgressBar_SetOrientation(p.instance, value)
}

func (p *TProgressBar) ParentDoubleBuffered() bool {
    return ProgressBar_GetParentDoubleBuffered(p.instance)
}

func (p *TProgressBar) SetParentDoubleBuffered(value bool) {
    ProgressBar_SetParentDoubleBuffered(p.instance, value)
}

func (p *TProgressBar) ParentShowHint() bool {
    return ProgressBar_GetParentShowHint(p.instance)
}

func (p *TProgressBar) SetParentShowHint(value bool) {
    ProgressBar_SetParentShowHint(p.instance, value)
}

func (p *TProgressBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ProgressBar_GetPopupMenu(p.instance))
}

func (p *TProgressBar) SetPopupMenu(value IComponent) {
    ProgressBar_SetPopupMenu(p.instance, CheckPtr(value))
}

func (p *TProgressBar) Position() int32 {
    return ProgressBar_GetPosition(p.instance)
}

func (p *TProgressBar) SetPosition(value int32) {
    ProgressBar_SetPosition(p.instance, value)
}

func (p *TProgressBar) Smooth() bool {
    return ProgressBar_GetSmooth(p.instance)
}

func (p *TProgressBar) SetSmooth(value bool) {
    ProgressBar_SetSmooth(p.instance, value)
}

func (p *TProgressBar) Style() TProgressBarStyle {
    return ProgressBar_GetStyle(p.instance)
}

func (p *TProgressBar) SetStyle(value TProgressBarStyle) {
    ProgressBar_SetStyle(p.instance, value)
}

func (p *TProgressBar) MarqueeInterval() int32 {
    return ProgressBar_GetMarqueeInterval(p.instance)
}

func (p *TProgressBar) SetMarqueeInterval(value int32) {
    ProgressBar_SetMarqueeInterval(p.instance, value)
}

func (p *TProgressBar) BarColor() TColor {
    return ProgressBar_GetBarColor(p.instance)
}

func (p *TProgressBar) SetBarColor(value TColor) {
    ProgressBar_SetBarColor(p.instance, value)
}

func (p *TProgressBar) BackgroundColor() TColor {
    return ProgressBar_GetBackgroundColor(p.instance)
}

func (p *TProgressBar) SetBackgroundColor(value TColor) {
    ProgressBar_SetBackgroundColor(p.instance, value)
}

func (p *TProgressBar) SmoothReverse() bool {
    return ProgressBar_GetSmoothReverse(p.instance)
}

func (p *TProgressBar) SetSmoothReverse(value bool) {
    ProgressBar_SetSmoothReverse(p.instance, value)
}

func (p *TProgressBar) Step() int32 {
    return ProgressBar_GetStep(p.instance)
}

func (p *TProgressBar) SetStep(value int32) {
    ProgressBar_SetStep(p.instance, value)
}

func (p *TProgressBar) State() TProgressBarState {
    return ProgressBar_GetState(p.instance)
}

func (p *TProgressBar) SetState(value TProgressBarState) {
    ProgressBar_SetState(p.instance, value)
}

func (p *TProgressBar) ShowHint() bool {
    return ProgressBar_GetShowHint(p.instance)
}

func (p *TProgressBar) SetShowHint(value bool) {
    ProgressBar_SetShowHint(p.instance, value)
}

func (p *TProgressBar) TabOrder() uint16 {
    return ProgressBar_GetTabOrder(p.instance)
}

func (p *TProgressBar) SetTabOrder(value uint16) {
    ProgressBar_SetTabOrder(p.instance, value)
}

func (p *TProgressBar) TabStop() bool {
    return ProgressBar_GetTabStop(p.instance)
}

func (p *TProgressBar) SetTabStop(value bool) {
    ProgressBar_SetTabStop(p.instance, value)
}

func (p *TProgressBar) Visible() bool {
    return ProgressBar_GetVisible(p.instance)
}

func (p *TProgressBar) SetVisible(value bool) {
    ProgressBar_SetVisible(p.instance, value)
}

func (p *TProgressBar) StyleElements() TStyleElements {
    return ProgressBar_GetStyleElements(p.instance)
}

func (p *TProgressBar) SetStyleElements(value TStyleElements) {
    ProgressBar_SetStyleElements(p.instance, value)
}

func (p *TProgressBar) SetOnEnter(fn TNotifyEvent) {
    ProgressBar_SetOnEnter(p.instance, fn)
}

func (p *TProgressBar) SetOnExit(fn TNotifyEvent) {
    ProgressBar_SetOnExit(p.instance, fn)
}

func (p *TProgressBar) SetOnMouseDown(fn TMouseEvent) {
    ProgressBar_SetOnMouseDown(p.instance, fn)
}

func (p *TProgressBar) SetOnMouseEnter(fn TNotifyEvent) {
    ProgressBar_SetOnMouseEnter(p.instance, fn)
}

func (p *TProgressBar) SetOnMouseLeave(fn TNotifyEvent) {
    ProgressBar_SetOnMouseLeave(p.instance, fn)
}

func (p *TProgressBar) SetOnMouseMove(fn TMouseMoveEvent) {
    ProgressBar_SetOnMouseMove(p.instance, fn)
}

func (p *TProgressBar) SetOnMouseUp(fn TMouseEvent) {
    ProgressBar_SetOnMouseUp(p.instance, fn)
}

func (p *TProgressBar) Brush() *TBrush {
    return BrushFromInst(ProgressBar_GetBrush(p.instance))
}

func (p *TProgressBar) ControlCount() int32 {
    return ProgressBar_GetControlCount(p.instance)
}

func (p *TProgressBar) Handle() HWND {
    return ProgressBar_GetHandle(p.instance)
}

func (p *TProgressBar) ParentWindow() HWND {
    return ProgressBar_GetParentWindow(p.instance)
}

func (p *TProgressBar) SetParentWindow(value HWND) {
    ProgressBar_SetParentWindow(p.instance, value)
}

func (p *TProgressBar) Action() *TAction {
    return ActionFromInst(ProgressBar_GetAction(p.instance))
}

func (p *TProgressBar) SetAction(value IComponent) {
    ProgressBar_SetAction(p.instance, CheckPtr(value))
}

func (p *TProgressBar) BiDiMode() TBiDiMode {
    return ProgressBar_GetBiDiMode(p.instance)
}

func (p *TProgressBar) SetBiDiMode(value TBiDiMode) {
    ProgressBar_SetBiDiMode(p.instance, value)
}

func (p *TProgressBar) BoundsRect() TRect {
    return ProgressBar_GetBoundsRect(p.instance)
}

func (p *TProgressBar) SetBoundsRect(value TRect) {
    ProgressBar_SetBoundsRect(p.instance, value)
}

func (p *TProgressBar) ClientHeight() int32 {
    return ProgressBar_GetClientHeight(p.instance)
}

func (p *TProgressBar) SetClientHeight(value int32) {
    ProgressBar_SetClientHeight(p.instance, value)
}

func (p *TProgressBar) ClientRect() TRect {
    return ProgressBar_GetClientRect(p.instance)
}

func (p *TProgressBar) ClientWidth() int32 {
    return ProgressBar_GetClientWidth(p.instance)
}

func (p *TProgressBar) SetClientWidth(value int32) {
    ProgressBar_SetClientWidth(p.instance, value)
}

func (p *TProgressBar) ExplicitLeft() int32 {
    return ProgressBar_GetExplicitLeft(p.instance)
}

func (p *TProgressBar) ExplicitTop() int32 {
    return ProgressBar_GetExplicitTop(p.instance)
}

func (p *TProgressBar) ExplicitWidth() int32 {
    return ProgressBar_GetExplicitWidth(p.instance)
}

func (p *TProgressBar) ExplicitHeight() int32 {
    return ProgressBar_GetExplicitHeight(p.instance)
}

func (p *TProgressBar) Parent() *TControl {
    return ControlFromInst(ProgressBar_GetParent(p.instance))
}

func (p *TProgressBar) SetParent(value IControl) {
    ProgressBar_SetParent(p.instance, CheckPtr(value))
}

func (p *TProgressBar) AlignWithMargins() bool {
    return ProgressBar_GetAlignWithMargins(p.instance)
}

func (p *TProgressBar) SetAlignWithMargins(value bool) {
    ProgressBar_SetAlignWithMargins(p.instance, value)
}

func (p *TProgressBar) Left() int32 {
    return ProgressBar_GetLeft(p.instance)
}

func (p *TProgressBar) SetLeft(value int32) {
    ProgressBar_SetLeft(p.instance, value)
}

func (p *TProgressBar) Top() int32 {
    return ProgressBar_GetTop(p.instance)
}

func (p *TProgressBar) SetTop(value int32) {
    ProgressBar_SetTop(p.instance, value)
}

func (p *TProgressBar) Width() int32 {
    return ProgressBar_GetWidth(p.instance)
}

func (p *TProgressBar) SetWidth(value int32) {
    ProgressBar_SetWidth(p.instance, value)
}

func (p *TProgressBar) Height() int32 {
    return ProgressBar_GetHeight(p.instance)
}

func (p *TProgressBar) SetHeight(value int32) {
    ProgressBar_SetHeight(p.instance, value)
}

func (p *TProgressBar) Cursor() TCursor {
    return ProgressBar_GetCursor(p.instance)
}

func (p *TProgressBar) SetCursor(value TCursor) {
    ProgressBar_SetCursor(p.instance, value)
}

func (p *TProgressBar) Margins() *TMargins {
    return MarginsFromInst(ProgressBar_GetMargins(p.instance))
}

func (p *TProgressBar) SetMargins(value *TMargins) {
    ProgressBar_SetMargins(p.instance, CheckPtr(value))
}

func (p *TProgressBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(ProgressBar_GetCustomHint(p.instance))
}

func (p *TProgressBar) SetCustomHint(value IComponent) {
    ProgressBar_SetCustomHint(p.instance, CheckPtr(value))
}

func (p *TProgressBar) ComponentCount() int32 {
    return ProgressBar_GetComponentCount(p.instance)
}

func (p *TProgressBar) ComponentIndex() int32 {
    return ProgressBar_GetComponentIndex(p.instance)
}

func (p *TProgressBar) SetComponentIndex(value int32) {
    ProgressBar_SetComponentIndex(p.instance, value)
}

func (p *TProgressBar) Owner() *TComponent {
    return ComponentFromInst(ProgressBar_GetOwner(p.instance))
}

func (p *TProgressBar) Name() string {
    return ProgressBar_GetName(p.instance)
}

func (p *TProgressBar) SetName(value string) {
    ProgressBar_SetName(p.instance, value)
}

func (p *TProgressBar) Tag() int {
    return ProgressBar_GetTag(p.instance)
}

func (p *TProgressBar) SetTag(value int) {
    ProgressBar_SetTag(p.instance, value)
}

func (p *TProgressBar) Controls(Index int32) *TControl {
    return ControlFromInst(ProgressBar_GetControls(p.instance, Index))
}

func (p *TProgressBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ProgressBar_GetComponents(p.instance, AIndex))
}

