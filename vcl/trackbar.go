
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

type TTrackBar struct {
    IControl
    instance uintptr
}

func NewTrackBar(owner IComponent) *TTrackBar {
    t := new(TTrackBar)
    t.instance = TrackBar_Create(CheckPtr(owner))
    return t
}

func TrackBarFromInst(inst uintptr) *TTrackBar {
    t := new(TTrackBar)
    t.instance = inst
    return t
}

func TrackBarFromObj(obj IObject) *TTrackBar {
    t := new(TTrackBar)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTrackBar) Free() {
    if t.instance != 0 {
        TrackBar_Free(t.instance)
        t.instance = 0
    }
}

func (t *TTrackBar) Instance() uintptr {
    return t.instance
}

func (t *TTrackBar) IsValid() bool {
    return t.instance != 0
}

func (t *TTrackBar) SetTick(Value int32) {
    TrackBar_SetTick(t.instance, Value)
}

func (t *TTrackBar) CanFocus() bool {
    return TrackBar_CanFocus(t.instance)
}

func (t *TTrackBar) FlipChildren(AllLevels bool) {
    TrackBar_FlipChildren(t.instance, AllLevels)
}

func (t *TTrackBar) Focused() bool {
    return TrackBar_Focused(t.instance)
}

func (t *TTrackBar) HandleAllocated() bool {
    return TrackBar_HandleAllocated(t.instance)
}

func (t *TTrackBar) Invalidate() {
    TrackBar_Invalidate(t.instance)
}

func (t *TTrackBar) Realign() {
    TrackBar_Realign(t.instance)
}

func (t *TTrackBar) Repaint() {
    TrackBar_Repaint(t.instance)
}

func (t *TTrackBar) ScaleBy(M int32, D int32) {
    TrackBar_ScaleBy(t.instance, M , D)
}

func (t *TTrackBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    TrackBar_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

func (t *TTrackBar) SetFocus() {
    TrackBar_SetFocus(t.instance)
}

func (t *TTrackBar) Update() {
    TrackBar_Update(t.instance)
}

func (t *TTrackBar) BringToFront() {
    TrackBar_BringToFront(t.instance)
}

func (t *TTrackBar) HasParent() bool {
    return TrackBar_HasParent(t.instance)
}

func (t *TTrackBar) Hide() {
    TrackBar_Hide(t.instance)
}

func (t *TTrackBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return TrackBar_Perform(t.instance, Msg , WParam , LParam)
}

func (t *TTrackBar) Refresh() {
    TrackBar_Refresh(t.instance)
}

func (t *TTrackBar) SendToBack() {
    TrackBar_SendToBack(t.instance)
}

func (t *TTrackBar) Show() {
    TrackBar_Show(t.instance)
}

func (t *TTrackBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return TrackBar_GetTextBuf(t.instance, Buffer , BufSize)
}

func (t *TTrackBar) GetTextLen() int32 {
    return TrackBar_GetTextLen(t.instance)
}

func (t *TTrackBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TrackBar_FindComponent(t.instance, AName))
}

func (t *TTrackBar) GetNamePath() string {
    return TrackBar_GetNamePath(t.instance)
}

func (t *TTrackBar) Assign(Source IObject) {
    TrackBar_Assign(t.instance, CheckPtr(Source))
}

func (t *TTrackBar) ClassName() string {
    return TrackBar_ClassName(t.instance)
}

func (t *TTrackBar) Equals(Obj IObject) bool {
    return TrackBar_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTrackBar) GetHashCode() int32 {
    return TrackBar_GetHashCode(t.instance)
}

func (t *TTrackBar) ToString() string {
    return TrackBar_ToString(t.instance)
}

func (t *TTrackBar) Align() TAlign {
    return TrackBar_GetAlign(t.instance)
}

func (t *TTrackBar) SetAlign(value TAlign) {
    TrackBar_SetAlign(t.instance, value)
}

func (t *TTrackBar) Anchors() TAnchors {
    return TrackBar_GetAnchors(t.instance)
}

func (t *TTrackBar) SetAnchors(value TAnchors) {
    TrackBar_SetAnchors(t.instance, value)
}

func (t *TTrackBar) BorderWidth() int32 {
    return TrackBar_GetBorderWidth(t.instance)
}

func (t *TTrackBar) SetBorderWidth(value int32) {
    TrackBar_SetBorderWidth(t.instance, value)
}

func (t *TTrackBar) DoubleBuffered() bool {
    return TrackBar_GetDoubleBuffered(t.instance)
}

func (t *TTrackBar) SetDoubleBuffered(value bool) {
    TrackBar_SetDoubleBuffered(t.instance, value)
}

func (t *TTrackBar) Enabled() bool {
    return TrackBar_GetEnabled(t.instance)
}

func (t *TTrackBar) SetEnabled(value bool) {
    TrackBar_SetEnabled(t.instance, value)
}

func (t *TTrackBar) LineSize() int32 {
    return TrackBar_GetLineSize(t.instance)
}

func (t *TTrackBar) SetLineSize(value int32) {
    TrackBar_SetLineSize(t.instance, value)
}

func (t *TTrackBar) Max() int32 {
    return TrackBar_GetMax(t.instance)
}

func (t *TTrackBar) SetMax(value int32) {
    TrackBar_SetMax(t.instance, value)
}

func (t *TTrackBar) Min() int32 {
    return TrackBar_GetMin(t.instance)
}

func (t *TTrackBar) SetMin(value int32) {
    TrackBar_SetMin(t.instance, value)
}

func (t *TTrackBar) Orientation() TTrackBarOrientation {
    return TrackBar_GetOrientation(t.instance)
}

func (t *TTrackBar) SetOrientation(value TTrackBarOrientation) {
    TrackBar_SetOrientation(t.instance, value)
}

func (t *TTrackBar) ParentCtl3D() bool {
    return TrackBar_GetParentCtl3D(t.instance)
}

func (t *TTrackBar) SetParentCtl3D(value bool) {
    TrackBar_SetParentCtl3D(t.instance, value)
}

func (t *TTrackBar) ParentDoubleBuffered() bool {
    return TrackBar_GetParentDoubleBuffered(t.instance)
}

func (t *TTrackBar) SetParentDoubleBuffered(value bool) {
    TrackBar_SetParentDoubleBuffered(t.instance, value)
}

func (t *TTrackBar) ParentShowHint() bool {
    return TrackBar_GetParentShowHint(t.instance)
}

func (t *TTrackBar) SetParentShowHint(value bool) {
    TrackBar_SetParentShowHint(t.instance, value)
}

func (t *TTrackBar) PageSize() int32 {
    return TrackBar_GetPageSize(t.instance)
}

func (t *TTrackBar) SetPageSize(value int32) {
    TrackBar_SetPageSize(t.instance, value)
}

func (t *TTrackBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TrackBar_GetPopupMenu(t.instance))
}

func (t *TTrackBar) SetPopupMenu(value IComponent) {
    TrackBar_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TTrackBar) Frequency() int32 {
    return TrackBar_GetFrequency(t.instance)
}

func (t *TTrackBar) SetFrequency(value int32) {
    TrackBar_SetFrequency(t.instance, value)
}

func (t *TTrackBar) Position() int32 {
    return TrackBar_GetPosition(t.instance)
}

func (t *TTrackBar) SetPosition(value int32) {
    TrackBar_SetPosition(t.instance, value)
}

func (t *TTrackBar) PositionToolTip() TPositionToolTip {
    return TrackBar_GetPositionToolTip(t.instance)
}

func (t *TTrackBar) SetPositionToolTip(value TPositionToolTip) {
    TrackBar_SetPositionToolTip(t.instance, value)
}

func (t *TTrackBar) SliderVisible() bool {
    return TrackBar_GetSliderVisible(t.instance)
}

func (t *TTrackBar) SetSliderVisible(value bool) {
    TrackBar_SetSliderVisible(t.instance, value)
}

func (t *TTrackBar) SelEnd() int32 {
    return TrackBar_GetSelEnd(t.instance)
}

func (t *TTrackBar) SetSelEnd(value int32) {
    TrackBar_SetSelEnd(t.instance, value)
}

func (t *TTrackBar) SelStart() int32 {
    return TrackBar_GetSelStart(t.instance)
}

func (t *TTrackBar) SetSelStart(value int32) {
    TrackBar_SetSelStart(t.instance, value)
}

func (t *TTrackBar) ShowHint() bool {
    return TrackBar_GetShowHint(t.instance)
}

func (t *TTrackBar) SetShowHint(value bool) {
    TrackBar_SetShowHint(t.instance, value)
}

func (t *TTrackBar) ShowSelRange() bool {
    return TrackBar_GetShowSelRange(t.instance)
}

func (t *TTrackBar) SetShowSelRange(value bool) {
    TrackBar_SetShowSelRange(t.instance, value)
}

func (t *TTrackBar) TabOrder() uint16 {
    return TrackBar_GetTabOrder(t.instance)
}

func (t *TTrackBar) SetTabOrder(value uint16) {
    TrackBar_SetTabOrder(t.instance, value)
}

func (t *TTrackBar) TabStop() bool {
    return TrackBar_GetTabStop(t.instance)
}

func (t *TTrackBar) SetTabStop(value bool) {
    TrackBar_SetTabStop(t.instance, value)
}

func (t *TTrackBar) ThumbLength() int32 {
    return TrackBar_GetThumbLength(t.instance)
}

func (t *TTrackBar) SetThumbLength(value int32) {
    TrackBar_SetThumbLength(t.instance, value)
}

func (t *TTrackBar) TickMarks() TTickMark {
    return TrackBar_GetTickMarks(t.instance)
}

func (t *TTrackBar) SetTickMarks(value TTickMark) {
    TrackBar_SetTickMarks(t.instance, value)
}

func (t *TTrackBar) TickStyle() TTickStyle {
    return TrackBar_GetTickStyle(t.instance)
}

func (t *TTrackBar) SetTickStyle(value TTickStyle) {
    TrackBar_SetTickStyle(t.instance, value)
}

func (t *TTrackBar) Visible() bool {
    return TrackBar_GetVisible(t.instance)
}

func (t *TTrackBar) SetVisible(value bool) {
    TrackBar_SetVisible(t.instance, value)
}

func (t *TTrackBar) StyleElements() TStyleElements {
    return TrackBar_GetStyleElements(t.instance)
}

func (t *TTrackBar) SetStyleElements(value TStyleElements) {
    TrackBar_SetStyleElements(t.instance, value)
}

func (t *TTrackBar) SetOnChange(fn TNotifyEvent) {
    TrackBar_SetOnChange(t.instance, fn)
}

func (t *TTrackBar) SetOnEnter(fn TNotifyEvent) {
    TrackBar_SetOnEnter(t.instance, fn)
}

func (t *TTrackBar) SetOnExit(fn TNotifyEvent) {
    TrackBar_SetOnExit(t.instance, fn)
}

func (t *TTrackBar) SetOnKeyDown(fn TKeyEvent) {
    TrackBar_SetOnKeyDown(t.instance, fn)
}

func (t *TTrackBar) SetOnKeyPress(fn TKeyPressEvent) {
    TrackBar_SetOnKeyPress(t.instance, fn)
}

func (t *TTrackBar) SetOnKeyUp(fn TKeyEvent) {
    TrackBar_SetOnKeyUp(t.instance, fn)
}

func (t *TTrackBar) Brush() *TBrush {
    return BrushFromInst(TrackBar_GetBrush(t.instance))
}

func (t *TTrackBar) ControlCount() int32 {
    return TrackBar_GetControlCount(t.instance)
}

func (t *TTrackBar) Handle() HWND {
    return TrackBar_GetHandle(t.instance)
}

func (t *TTrackBar) ParentWindow() HWND {
    return TrackBar_GetParentWindow(t.instance)
}

func (t *TTrackBar) SetParentWindow(value HWND) {
    TrackBar_SetParentWindow(t.instance, value)
}

func (t *TTrackBar) Action() *TAction {
    return ActionFromInst(TrackBar_GetAction(t.instance))
}

func (t *TTrackBar) SetAction(value IComponent) {
    TrackBar_SetAction(t.instance, CheckPtr(value))
}

func (t *TTrackBar) BiDiMode() TBiDiMode {
    return TrackBar_GetBiDiMode(t.instance)
}

func (t *TTrackBar) SetBiDiMode(value TBiDiMode) {
    TrackBar_SetBiDiMode(t.instance, value)
}

func (t *TTrackBar) BoundsRect() TRect {
    return TrackBar_GetBoundsRect(t.instance)
}

func (t *TTrackBar) SetBoundsRect(value TRect) {
    TrackBar_SetBoundsRect(t.instance, value)
}

func (t *TTrackBar) ClientHeight() int32 {
    return TrackBar_GetClientHeight(t.instance)
}

func (t *TTrackBar) SetClientHeight(value int32) {
    TrackBar_SetClientHeight(t.instance, value)
}

func (t *TTrackBar) ClientRect() TRect {
    return TrackBar_GetClientRect(t.instance)
}

func (t *TTrackBar) ClientWidth() int32 {
    return TrackBar_GetClientWidth(t.instance)
}

func (t *TTrackBar) SetClientWidth(value int32) {
    TrackBar_SetClientWidth(t.instance, value)
}

func (t *TTrackBar) ExplicitLeft() int32 {
    return TrackBar_GetExplicitLeft(t.instance)
}

func (t *TTrackBar) ExplicitTop() int32 {
    return TrackBar_GetExplicitTop(t.instance)
}

func (t *TTrackBar) ExplicitWidth() int32 {
    return TrackBar_GetExplicitWidth(t.instance)
}

func (t *TTrackBar) ExplicitHeight() int32 {
    return TrackBar_GetExplicitHeight(t.instance)
}

func (t *TTrackBar) Parent() *TControl {
    return ControlFromInst(TrackBar_GetParent(t.instance))
}

func (t *TTrackBar) SetParent(value IControl) {
    TrackBar_SetParent(t.instance, CheckPtr(value))
}

func (t *TTrackBar) AlignWithMargins() bool {
    return TrackBar_GetAlignWithMargins(t.instance)
}

func (t *TTrackBar) SetAlignWithMargins(value bool) {
    TrackBar_SetAlignWithMargins(t.instance, value)
}

func (t *TTrackBar) Left() int32 {
    return TrackBar_GetLeft(t.instance)
}

func (t *TTrackBar) SetLeft(value int32) {
    TrackBar_SetLeft(t.instance, value)
}

func (t *TTrackBar) Top() int32 {
    return TrackBar_GetTop(t.instance)
}

func (t *TTrackBar) SetTop(value int32) {
    TrackBar_SetTop(t.instance, value)
}

func (t *TTrackBar) Width() int32 {
    return TrackBar_GetWidth(t.instance)
}

func (t *TTrackBar) SetWidth(value int32) {
    TrackBar_SetWidth(t.instance, value)
}

func (t *TTrackBar) Height() int32 {
    return TrackBar_GetHeight(t.instance)
}

func (t *TTrackBar) SetHeight(value int32) {
    TrackBar_SetHeight(t.instance, value)
}

func (t *TTrackBar) Cursor() TCursor {
    return TrackBar_GetCursor(t.instance)
}

func (t *TTrackBar) SetCursor(value TCursor) {
    TrackBar_SetCursor(t.instance, value)
}

func (t *TTrackBar) Hint() string {
    return TrackBar_GetHint(t.instance)
}

func (t *TTrackBar) SetHint(value string) {
    TrackBar_SetHint(t.instance, value)
}

func (t *TTrackBar) Margins() *TMargins {
    return MarginsFromInst(TrackBar_GetMargins(t.instance))
}

func (t *TTrackBar) SetMargins(value *TMargins) {
    TrackBar_SetMargins(t.instance, CheckPtr(value))
}

func (t *TTrackBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(TrackBar_GetCustomHint(t.instance))
}

func (t *TTrackBar) SetCustomHint(value IComponent) {
    TrackBar_SetCustomHint(t.instance, CheckPtr(value))
}

func (t *TTrackBar) ComponentCount() int32 {
    return TrackBar_GetComponentCount(t.instance)
}

func (t *TTrackBar) ComponentIndex() int32 {
    return TrackBar_GetComponentIndex(t.instance)
}

func (t *TTrackBar) SetComponentIndex(value int32) {
    TrackBar_SetComponentIndex(t.instance, value)
}

func (t *TTrackBar) Owner() *TComponent {
    return ComponentFromInst(TrackBar_GetOwner(t.instance))
}

func (t *TTrackBar) Name() string {
    return TrackBar_GetName(t.instance)
}

func (t *TTrackBar) SetName(value string) {
    TrackBar_SetName(t.instance, value)
}

func (t *TTrackBar) Tag() int {
    return TrackBar_GetTag(t.instance)
}

func (t *TTrackBar) SetTag(value int) {
    TrackBar_SetTag(t.instance, value)
}

func (t *TTrackBar) Controls(Index int32) *TControl {
    return ControlFromInst(TrackBar_GetControls(t.instance, Index))
}

func (t *TTrackBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TrackBar_GetComponents(t.instance, AIndex))
}

