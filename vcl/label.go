
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

type TLabel struct {
    IControl
    instance uintptr
}

func NewLabel(owner IComponent) *TLabel {
    l := new(TLabel)
    l.instance = Label_Create(CheckPtr(owner))
    return l
}

func LabelFromInst(inst uintptr) *TLabel {
    l := new(TLabel)
    l.instance = inst
    return l
}

func LabelFromObj(obj IObject) *TLabel {
    l := new(TLabel)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TLabel) Free() {
    if l.instance != 0 {
        Label_Free(l.instance)
        l.instance = 0
    }
}

func (l *TLabel) Instance() uintptr {
    return l.instance
}

func (l *TLabel) IsValid() bool {
    return l.instance != 0
}

func (l *TLabel) BringToFront() {
    Label_BringToFront(l.instance)
}

func (l *TLabel) HasParent() bool {
    return Label_HasParent(l.instance)
}

func (l *TLabel) Hide() {
    Label_Hide(l.instance)
}

func (l *TLabel) Invalidate() {
    Label_Invalidate(l.instance)
}

func (l *TLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Label_Perform(l.instance, Msg , WParam , LParam)
}

func (l *TLabel) Refresh() {
    Label_Refresh(l.instance)
}

func (l *TLabel) Repaint() {
    Label_Repaint(l.instance)
}

func (l *TLabel) SendToBack() {
    Label_SendToBack(l.instance)
}

func (l *TLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Label_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

func (l *TLabel) Show() {
    Label_Show(l.instance)
}

func (l *TLabel) Update() {
    Label_Update(l.instance)
}

func (l *TLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Label_GetTextBuf(l.instance, Buffer , BufSize)
}

func (l *TLabel) GetTextLen() int32 {
    return Label_GetTextLen(l.instance)
}

func (l *TLabel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Label_FindComponent(l.instance, AName))
}

func (l *TLabel) GetNamePath() string {
    return Label_GetNamePath(l.instance)
}

func (l *TLabel) Assign(Source IObject) {
    Label_Assign(l.instance, CheckPtr(Source))
}

func (l *TLabel) ClassName() string {
    return Label_ClassName(l.instance)
}

func (l *TLabel) Equals(Obj IObject) bool {
    return Label_Equals(l.instance, CheckPtr(Obj))
}

func (l *TLabel) GetHashCode() int32 {
    return Label_GetHashCode(l.instance)
}

func (l *TLabel) ToString() string {
    return Label_ToString(l.instance)
}

func (l *TLabel) Align() TAlign {
    return Label_GetAlign(l.instance)
}

func (l *TLabel) SetAlign(value TAlign) {
    Label_SetAlign(l.instance, value)
}

func (l *TLabel) Alignment() TAlignment {
    return Label_GetAlignment(l.instance)
}

func (l *TLabel) SetAlignment(value TAlignment) {
    Label_SetAlignment(l.instance, value)
}

func (l *TLabel) Anchors() TAnchors {
    return Label_GetAnchors(l.instance)
}

func (l *TLabel) SetAnchors(value TAnchors) {
    Label_SetAnchors(l.instance, value)
}

func (l *TLabel) AutoSize() bool {
    return Label_GetAutoSize(l.instance)
}

func (l *TLabel) SetAutoSize(value bool) {
    Label_SetAutoSize(l.instance, value)
}

func (l *TLabel) BiDiMode() TBiDiMode {
    return Label_GetBiDiMode(l.instance)
}

func (l *TLabel) SetBiDiMode(value TBiDiMode) {
    Label_SetBiDiMode(l.instance, value)
}

func (l *TLabel) Caption() string {
    return Label_GetCaption(l.instance)
}

func (l *TLabel) SetCaption(value string) {
    Label_SetCaption(l.instance, value)
}

func (l *TLabel) Color() TColor {
    return Label_GetColor(l.instance)
}

func (l *TLabel) SetColor(value TColor) {
    Label_SetColor(l.instance, value)
}

func (l *TLabel) EllipsisPosition() TEllipsisPosition {
    return Label_GetEllipsisPosition(l.instance)
}

func (l *TLabel) SetEllipsisPosition(value TEllipsisPosition) {
    Label_SetEllipsisPosition(l.instance, value)
}

func (l *TLabel) Enabled() bool {
    return Label_GetEnabled(l.instance)
}

func (l *TLabel) SetEnabled(value bool) {
    Label_SetEnabled(l.instance, value)
}

func (l *TLabel) Font() *TFont {
    return FontFromInst(Label_GetFont(l.instance))
}

func (l *TLabel) SetFont(value *TFont) {
    Label_SetFont(l.instance, CheckPtr(value))
}

func (l *TLabel) GlowSize() int32 {
    return Label_GetGlowSize(l.instance)
}

func (l *TLabel) SetGlowSize(value int32) {
    Label_SetGlowSize(l.instance, value)
}

func (l *TLabel) ParentColor() bool {
    return Label_GetParentColor(l.instance)
}

func (l *TLabel) SetParentColor(value bool) {
    Label_SetParentColor(l.instance, value)
}

func (l *TLabel) ParentFont() bool {
    return Label_GetParentFont(l.instance)
}

func (l *TLabel) SetParentFont(value bool) {
    Label_SetParentFont(l.instance, value)
}

func (l *TLabel) ParentShowHint() bool {
    return Label_GetParentShowHint(l.instance)
}

func (l *TLabel) SetParentShowHint(value bool) {
    Label_SetParentShowHint(l.instance, value)
}

func (l *TLabel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Label_GetPopupMenu(l.instance))
}

func (l *TLabel) SetPopupMenu(value IComponent) {
    Label_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TLabel) ShowAccelChar() bool {
    return Label_GetShowAccelChar(l.instance)
}

func (l *TLabel) SetShowAccelChar(value bool) {
    Label_SetShowAccelChar(l.instance, value)
}

func (l *TLabel) ShowHint() bool {
    return Label_GetShowHint(l.instance)
}

func (l *TLabel) SetShowHint(value bool) {
    Label_SetShowHint(l.instance, value)
}

func (l *TLabel) Transparent() bool {
    return Label_GetTransparent(l.instance)
}

func (l *TLabel) SetTransparent(value bool) {
    Label_SetTransparent(l.instance, value)
}

func (l *TLabel) Layout() TTextLayout {
    return Label_GetLayout(l.instance)
}

func (l *TLabel) SetLayout(value TTextLayout) {
    Label_SetLayout(l.instance, value)
}

func (l *TLabel) Visible() bool {
    return Label_GetVisible(l.instance)
}

func (l *TLabel) SetVisible(value bool) {
    Label_SetVisible(l.instance, value)
}

func (l *TLabel) WordWrap() bool {
    return Label_GetWordWrap(l.instance)
}

func (l *TLabel) SetWordWrap(value bool) {
    Label_SetWordWrap(l.instance, value)
}

func (l *TLabel) StyleElements() TStyleElements {
    return Label_GetStyleElements(l.instance)
}

func (l *TLabel) SetStyleElements(value TStyleElements) {
    Label_SetStyleElements(l.instance, value)
}

func (l *TLabel) SetOnClick(fn TNotifyEvent) {
    Label_SetOnClick(l.instance, fn)
}

func (l *TLabel) SetOnDblClick(fn TNotifyEvent) {
    Label_SetOnDblClick(l.instance, fn)
}

func (l *TLabel) SetOnMouseDown(fn TMouseEvent) {
    Label_SetOnMouseDown(l.instance, fn)
}

func (l *TLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    Label_SetOnMouseMove(l.instance, fn)
}

func (l *TLabel) SetOnMouseUp(fn TMouseEvent) {
    Label_SetOnMouseUp(l.instance, fn)
}

func (l *TLabel) SetOnMouseEnter(fn TNotifyEvent) {
    Label_SetOnMouseEnter(l.instance, fn)
}

func (l *TLabel) SetOnMouseLeave(fn TNotifyEvent) {
    Label_SetOnMouseLeave(l.instance, fn)
}

func (l *TLabel) Canvas() *TCanvas {
    return CanvasFromInst(Label_GetCanvas(l.instance))
}

func (l *TLabel) Action() *TAction {
    return ActionFromInst(Label_GetAction(l.instance))
}

func (l *TLabel) SetAction(value IComponent) {
    Label_SetAction(l.instance, CheckPtr(value))
}

func (l *TLabel) BoundsRect() TRect {
    return Label_GetBoundsRect(l.instance)
}

func (l *TLabel) SetBoundsRect(value TRect) {
    Label_SetBoundsRect(l.instance, value)
}

func (l *TLabel) ClientHeight() int32 {
    return Label_GetClientHeight(l.instance)
}

func (l *TLabel) SetClientHeight(value int32) {
    Label_SetClientHeight(l.instance, value)
}

func (l *TLabel) ClientRect() TRect {
    return Label_GetClientRect(l.instance)
}

func (l *TLabel) ClientWidth() int32 {
    return Label_GetClientWidth(l.instance)
}

func (l *TLabel) SetClientWidth(value int32) {
    Label_SetClientWidth(l.instance, value)
}

func (l *TLabel) ExplicitLeft() int32 {
    return Label_GetExplicitLeft(l.instance)
}

func (l *TLabel) ExplicitTop() int32 {
    return Label_GetExplicitTop(l.instance)
}

func (l *TLabel) ExplicitWidth() int32 {
    return Label_GetExplicitWidth(l.instance)
}

func (l *TLabel) ExplicitHeight() int32 {
    return Label_GetExplicitHeight(l.instance)
}

func (l *TLabel) Parent() *TControl {
    return ControlFromInst(Label_GetParent(l.instance))
}

func (l *TLabel) SetParent(value IControl) {
    Label_SetParent(l.instance, CheckPtr(value))
}

func (l *TLabel) AlignWithMargins() bool {
    return Label_GetAlignWithMargins(l.instance)
}

func (l *TLabel) SetAlignWithMargins(value bool) {
    Label_SetAlignWithMargins(l.instance, value)
}

func (l *TLabel) Left() int32 {
    return Label_GetLeft(l.instance)
}

func (l *TLabel) SetLeft(value int32) {
    Label_SetLeft(l.instance, value)
}

func (l *TLabel) Top() int32 {
    return Label_GetTop(l.instance)
}

func (l *TLabel) SetTop(value int32) {
    Label_SetTop(l.instance, value)
}

func (l *TLabel) Width() int32 {
    return Label_GetWidth(l.instance)
}

func (l *TLabel) SetWidth(value int32) {
    Label_SetWidth(l.instance, value)
}

func (l *TLabel) Height() int32 {
    return Label_GetHeight(l.instance)
}

func (l *TLabel) SetHeight(value int32) {
    Label_SetHeight(l.instance, value)
}

func (l *TLabel) Cursor() TCursor {
    return Label_GetCursor(l.instance)
}

func (l *TLabel) SetCursor(value TCursor) {
    Label_SetCursor(l.instance, value)
}

func (l *TLabel) Hint() string {
    return Label_GetHint(l.instance)
}

func (l *TLabel) SetHint(value string) {
    Label_SetHint(l.instance, value)
}

func (l *TLabel) Margins() *TMargins {
    return MarginsFromInst(Label_GetMargins(l.instance))
}

func (l *TLabel) SetMargins(value *TMargins) {
    Label_SetMargins(l.instance, CheckPtr(value))
}

func (l *TLabel) CustomHint() *TCustomHint {
    return CustomHintFromInst(Label_GetCustomHint(l.instance))
}

func (l *TLabel) SetCustomHint(value IComponent) {
    Label_SetCustomHint(l.instance, CheckPtr(value))
}

func (l *TLabel) ComponentCount() int32 {
    return Label_GetComponentCount(l.instance)
}

func (l *TLabel) ComponentIndex() int32 {
    return Label_GetComponentIndex(l.instance)
}

func (l *TLabel) SetComponentIndex(value int32) {
    Label_SetComponentIndex(l.instance, value)
}

func (l *TLabel) Owner() *TComponent {
    return ComponentFromInst(Label_GetOwner(l.instance))
}

func (l *TLabel) Name() string {
    return Label_GetName(l.instance)
}

func (l *TLabel) SetName(value string) {
    Label_SetName(l.instance, value)
}

func (l *TLabel) Tag() int {
    return Label_GetTag(l.instance)
}

func (l *TLabel) SetTag(value int) {
    Label_SetTag(l.instance, value)
}

func (l *TLabel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Label_GetComponents(l.instance, AIndex))
}

