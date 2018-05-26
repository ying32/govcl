
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

type TStaticText struct {
    IControl
    instance uintptr
}

func NewStaticText(owner IComponent) *TStaticText {
    s := new(TStaticText)
    s.instance = StaticText_Create(CheckPtr(owner))
    return s
}

func StaticTextFromInst(inst uintptr) *TStaticText {
    s := new(TStaticText)
    s.instance = inst
    return s
}

func StaticTextFromObj(obj IObject) *TStaticText {
    s := new(TStaticText)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStaticText) Free() {
    if s.instance != 0 {
        StaticText_Free(s.instance)
        s.instance = 0
    }
}

func (s *TStaticText) Instance() uintptr {
    return s.instance
}

func (s *TStaticText) IsValid() bool {
    return s.instance != 0
}

func (s *TStaticText) CanFocus() bool {
    return StaticText_CanFocus(s.instance)
}

func (s *TStaticText) FlipChildren(AllLevels bool) {
    StaticText_FlipChildren(s.instance, AllLevels)
}

func (s *TStaticText) Focused() bool {
    return StaticText_Focused(s.instance)
}

func (s *TStaticText) HandleAllocated() bool {
    return StaticText_HandleAllocated(s.instance)
}

func (s *TStaticText) Invalidate() {
    StaticText_Invalidate(s.instance)
}

func (s *TStaticText) Realign() {
    StaticText_Realign(s.instance)
}

func (s *TStaticText) Repaint() {
    StaticText_Repaint(s.instance)
}

func (s *TStaticText) ScaleBy(M int32, D int32) {
    StaticText_ScaleBy(s.instance, M , D)
}

func (s *TStaticText) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    StaticText_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

func (s *TStaticText) SetFocus() {
    StaticText_SetFocus(s.instance)
}

func (s *TStaticText) Update() {
    StaticText_Update(s.instance)
}

func (s *TStaticText) BringToFront() {
    StaticText_BringToFront(s.instance)
}

func (s *TStaticText) HasParent() bool {
    return StaticText_HasParent(s.instance)
}

func (s *TStaticText) Hide() {
    StaticText_Hide(s.instance)
}

func (s *TStaticText) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return StaticText_Perform(s.instance, Msg , WParam , LParam)
}

func (s *TStaticText) Refresh() {
    StaticText_Refresh(s.instance)
}

func (s *TStaticText) SendToBack() {
    StaticText_SendToBack(s.instance)
}

func (s *TStaticText) Show() {
    StaticText_Show(s.instance)
}

func (s *TStaticText) GetTextBuf(Buffer string, BufSize int32) int32 {
    return StaticText_GetTextBuf(s.instance, Buffer , BufSize)
}

func (s *TStaticText) GetTextLen() int32 {
    return StaticText_GetTextLen(s.instance)
}

func (s *TStaticText) FindComponent(AName string) *TComponent {
    return ComponentFromInst(StaticText_FindComponent(s.instance, AName))
}

func (s *TStaticText) GetNamePath() string {
    return StaticText_GetNamePath(s.instance)
}

func (s *TStaticText) Assign(Source IObject) {
    StaticText_Assign(s.instance, CheckPtr(Source))
}

func (s *TStaticText) ClassName() string {
    return StaticText_ClassName(s.instance)
}

func (s *TStaticText) Equals(Obj IObject) bool {
    return StaticText_Equals(s.instance, CheckPtr(Obj))
}

func (s *TStaticText) GetHashCode() int32 {
    return StaticText_GetHashCode(s.instance)
}

func (s *TStaticText) ToString() string {
    return StaticText_ToString(s.instance)
}

func (s *TStaticText) Align() TAlign {
    return StaticText_GetAlign(s.instance)
}

func (s *TStaticText) SetAlign(value TAlign) {
    StaticText_SetAlign(s.instance, value)
}

func (s *TStaticText) Alignment() TAlignment {
    return StaticText_GetAlignment(s.instance)
}

func (s *TStaticText) SetAlignment(value TAlignment) {
    StaticText_SetAlignment(s.instance, value)
}

func (s *TStaticText) Anchors() TAnchors {
    return StaticText_GetAnchors(s.instance)
}

func (s *TStaticText) SetAnchors(value TAnchors) {
    StaticText_SetAnchors(s.instance, value)
}

func (s *TStaticText) AutoSize() bool {
    return StaticText_GetAutoSize(s.instance)
}

func (s *TStaticText) SetAutoSize(value bool) {
    StaticText_SetAutoSize(s.instance, value)
}

func (s *TStaticText) BevelEdges() TBevelEdges {
    return StaticText_GetBevelEdges(s.instance)
}

func (s *TStaticText) SetBevelEdges(value TBevelEdges) {
    StaticText_SetBevelEdges(s.instance, value)
}

func (s *TStaticText) BevelInner() TBevelCut {
    return StaticText_GetBevelInner(s.instance)
}

func (s *TStaticText) SetBevelInner(value TBevelCut) {
    StaticText_SetBevelInner(s.instance, value)
}

func (s *TStaticText) BevelKind() TBevelKind {
    return StaticText_GetBevelKind(s.instance)
}

func (s *TStaticText) SetBevelKind(value TBevelKind) {
    StaticText_SetBevelKind(s.instance, value)
}

func (s *TStaticText) BevelOuter() TBevelCut {
    return StaticText_GetBevelOuter(s.instance)
}

func (s *TStaticText) SetBevelOuter(value TBevelCut) {
    StaticText_SetBevelOuter(s.instance, value)
}

func (s *TStaticText) BiDiMode() TBiDiMode {
    return StaticText_GetBiDiMode(s.instance)
}

func (s *TStaticText) SetBiDiMode(value TBiDiMode) {
    StaticText_SetBiDiMode(s.instance, value)
}

func (s *TStaticText) BorderStyle() TStaticBorderStyle {
    return StaticText_GetBorderStyle(s.instance)
}

func (s *TStaticText) SetBorderStyle(value TStaticBorderStyle) {
    StaticText_SetBorderStyle(s.instance, value)
}

func (s *TStaticText) Caption() string {
    return StaticText_GetCaption(s.instance)
}

func (s *TStaticText) SetCaption(value string) {
    StaticText_SetCaption(s.instance, value)
}

func (s *TStaticText) Color() TColor {
    return StaticText_GetColor(s.instance)
}

func (s *TStaticText) SetColor(value TColor) {
    StaticText_SetColor(s.instance, value)
}

func (s *TStaticText) DoubleBuffered() bool {
    return StaticText_GetDoubleBuffered(s.instance)
}

func (s *TStaticText) SetDoubleBuffered(value bool) {
    StaticText_SetDoubleBuffered(s.instance, value)
}

func (s *TStaticText) Enabled() bool {
    return StaticText_GetEnabled(s.instance)
}

func (s *TStaticText) SetEnabled(value bool) {
    StaticText_SetEnabled(s.instance, value)
}

func (s *TStaticText) Font() *TFont {
    return FontFromInst(StaticText_GetFont(s.instance))
}

func (s *TStaticText) SetFont(value *TFont) {
    StaticText_SetFont(s.instance, CheckPtr(value))
}

func (s *TStaticText) ParentColor() bool {
    return StaticText_GetParentColor(s.instance)
}

func (s *TStaticText) SetParentColor(value bool) {
    StaticText_SetParentColor(s.instance, value)
}

func (s *TStaticText) ParentDoubleBuffered() bool {
    return StaticText_GetParentDoubleBuffered(s.instance)
}

func (s *TStaticText) SetParentDoubleBuffered(value bool) {
    StaticText_SetParentDoubleBuffered(s.instance, value)
}

func (s *TStaticText) ParentFont() bool {
    return StaticText_GetParentFont(s.instance)
}

func (s *TStaticText) SetParentFont(value bool) {
    StaticText_SetParentFont(s.instance, value)
}

func (s *TStaticText) ParentShowHint() bool {
    return StaticText_GetParentShowHint(s.instance)
}

func (s *TStaticText) SetParentShowHint(value bool) {
    StaticText_SetParentShowHint(s.instance, value)
}

func (s *TStaticText) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(StaticText_GetPopupMenu(s.instance))
}

func (s *TStaticText) SetPopupMenu(value IComponent) {
    StaticText_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TStaticText) ShowAccelChar() bool {
    return StaticText_GetShowAccelChar(s.instance)
}

func (s *TStaticText) SetShowAccelChar(value bool) {
    StaticText_SetShowAccelChar(s.instance, value)
}

func (s *TStaticText) ShowHint() bool {
    return StaticText_GetShowHint(s.instance)
}

func (s *TStaticText) SetShowHint(value bool) {
    StaticText_SetShowHint(s.instance, value)
}

func (s *TStaticText) TabOrder() uint16 {
    return StaticText_GetTabOrder(s.instance)
}

func (s *TStaticText) SetTabOrder(value uint16) {
    StaticText_SetTabOrder(s.instance, value)
}

func (s *TStaticText) TabStop() bool {
    return StaticText_GetTabStop(s.instance)
}

func (s *TStaticText) SetTabStop(value bool) {
    StaticText_SetTabStop(s.instance, value)
}

func (s *TStaticText) Transparent() bool {
    return StaticText_GetTransparent(s.instance)
}

func (s *TStaticText) SetTransparent(value bool) {
    StaticText_SetTransparent(s.instance, value)
}

func (s *TStaticText) Visible() bool {
    return StaticText_GetVisible(s.instance)
}

func (s *TStaticText) SetVisible(value bool) {
    StaticText_SetVisible(s.instance, value)
}

func (s *TStaticText) StyleElements() TStyleElements {
    return StaticText_GetStyleElements(s.instance)
}

func (s *TStaticText) SetStyleElements(value TStyleElements) {
    StaticText_SetStyleElements(s.instance, value)
}

func (s *TStaticText) SetOnClick(fn TNotifyEvent) {
    StaticText_SetOnClick(s.instance, fn)
}

func (s *TStaticText) SetOnDblClick(fn TNotifyEvent) {
    StaticText_SetOnDblClick(s.instance, fn)
}

func (s *TStaticText) SetOnMouseDown(fn TMouseEvent) {
    StaticText_SetOnMouseDown(s.instance, fn)
}

func (s *TStaticText) SetOnMouseEnter(fn TNotifyEvent) {
    StaticText_SetOnMouseEnter(s.instance, fn)
}

func (s *TStaticText) SetOnMouseLeave(fn TNotifyEvent) {
    StaticText_SetOnMouseLeave(s.instance, fn)
}

func (s *TStaticText) SetOnMouseMove(fn TMouseMoveEvent) {
    StaticText_SetOnMouseMove(s.instance, fn)
}

func (s *TStaticText) SetOnMouseUp(fn TMouseEvent) {
    StaticText_SetOnMouseUp(s.instance, fn)
}

func (s *TStaticText) Brush() *TBrush {
    return BrushFromInst(StaticText_GetBrush(s.instance))
}

func (s *TStaticText) ControlCount() int32 {
    return StaticText_GetControlCount(s.instance)
}

func (s *TStaticText) Handle() HWND {
    return StaticText_GetHandle(s.instance)
}

func (s *TStaticText) ParentWindow() HWND {
    return StaticText_GetParentWindow(s.instance)
}

func (s *TStaticText) SetParentWindow(value HWND) {
    StaticText_SetParentWindow(s.instance, value)
}

func (s *TStaticText) Action() *TAction {
    return ActionFromInst(StaticText_GetAction(s.instance))
}

func (s *TStaticText) SetAction(value IComponent) {
    StaticText_SetAction(s.instance, CheckPtr(value))
}

func (s *TStaticText) BoundsRect() TRect {
    return StaticText_GetBoundsRect(s.instance)
}

func (s *TStaticText) SetBoundsRect(value TRect) {
    StaticText_SetBoundsRect(s.instance, value)
}

func (s *TStaticText) ClientHeight() int32 {
    return StaticText_GetClientHeight(s.instance)
}

func (s *TStaticText) SetClientHeight(value int32) {
    StaticText_SetClientHeight(s.instance, value)
}

func (s *TStaticText) ClientRect() TRect {
    return StaticText_GetClientRect(s.instance)
}

func (s *TStaticText) ClientWidth() int32 {
    return StaticText_GetClientWidth(s.instance)
}

func (s *TStaticText) SetClientWidth(value int32) {
    StaticText_SetClientWidth(s.instance, value)
}

func (s *TStaticText) ExplicitLeft() int32 {
    return StaticText_GetExplicitLeft(s.instance)
}

func (s *TStaticText) ExplicitTop() int32 {
    return StaticText_GetExplicitTop(s.instance)
}

func (s *TStaticText) ExplicitWidth() int32 {
    return StaticText_GetExplicitWidth(s.instance)
}

func (s *TStaticText) ExplicitHeight() int32 {
    return StaticText_GetExplicitHeight(s.instance)
}

func (s *TStaticText) Parent() *TControl {
    return ControlFromInst(StaticText_GetParent(s.instance))
}

func (s *TStaticText) SetParent(value IControl) {
    StaticText_SetParent(s.instance, CheckPtr(value))
}

func (s *TStaticText) AlignWithMargins() bool {
    return StaticText_GetAlignWithMargins(s.instance)
}

func (s *TStaticText) SetAlignWithMargins(value bool) {
    StaticText_SetAlignWithMargins(s.instance, value)
}

func (s *TStaticText) Left() int32 {
    return StaticText_GetLeft(s.instance)
}

func (s *TStaticText) SetLeft(value int32) {
    StaticText_SetLeft(s.instance, value)
}

func (s *TStaticText) Top() int32 {
    return StaticText_GetTop(s.instance)
}

func (s *TStaticText) SetTop(value int32) {
    StaticText_SetTop(s.instance, value)
}

func (s *TStaticText) Width() int32 {
    return StaticText_GetWidth(s.instance)
}

func (s *TStaticText) SetWidth(value int32) {
    StaticText_SetWidth(s.instance, value)
}

func (s *TStaticText) Height() int32 {
    return StaticText_GetHeight(s.instance)
}

func (s *TStaticText) SetHeight(value int32) {
    StaticText_SetHeight(s.instance, value)
}

func (s *TStaticText) Cursor() TCursor {
    return StaticText_GetCursor(s.instance)
}

func (s *TStaticText) SetCursor(value TCursor) {
    StaticText_SetCursor(s.instance, value)
}

func (s *TStaticText) Hint() string {
    return StaticText_GetHint(s.instance)
}

func (s *TStaticText) SetHint(value string) {
    StaticText_SetHint(s.instance, value)
}

func (s *TStaticText) Margins() *TMargins {
    return MarginsFromInst(StaticText_GetMargins(s.instance))
}

func (s *TStaticText) SetMargins(value *TMargins) {
    StaticText_SetMargins(s.instance, CheckPtr(value))
}

func (s *TStaticText) CustomHint() *TCustomHint {
    return CustomHintFromInst(StaticText_GetCustomHint(s.instance))
}

func (s *TStaticText) SetCustomHint(value IComponent) {
    StaticText_SetCustomHint(s.instance, CheckPtr(value))
}

func (s *TStaticText) ComponentCount() int32 {
    return StaticText_GetComponentCount(s.instance)
}

func (s *TStaticText) ComponentIndex() int32 {
    return StaticText_GetComponentIndex(s.instance)
}

func (s *TStaticText) SetComponentIndex(value int32) {
    StaticText_SetComponentIndex(s.instance, value)
}

func (s *TStaticText) Owner() *TComponent {
    return ComponentFromInst(StaticText_GetOwner(s.instance))
}

func (s *TStaticText) Name() string {
    return StaticText_GetName(s.instance)
}

func (s *TStaticText) SetName(value string) {
    StaticText_SetName(s.instance, value)
}

func (s *TStaticText) Tag() int {
    return StaticText_GetTag(s.instance)
}

func (s *TStaticText) SetTag(value int) {
    StaticText_SetTag(s.instance, value)
}

func (s *TStaticText) Controls(Index int32) *TControl {
    return ControlFromInst(StaticText_GetControls(s.instance, Index))
}

func (s *TStaticText) Components(AIndex int32) *TComponent {
    return ComponentFromInst(StaticText_GetComponents(s.instance, AIndex))
}

