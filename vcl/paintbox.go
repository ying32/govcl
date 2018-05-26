
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

type TPaintBox struct {
    IControl
    instance uintptr
}

func NewPaintBox(owner IComponent) *TPaintBox {
    p := new(TPaintBox)
    p.instance = PaintBox_Create(CheckPtr(owner))
    return p
}

func PaintBoxFromInst(inst uintptr) *TPaintBox {
    p := new(TPaintBox)
    p.instance = inst
    return p
}

func PaintBoxFromObj(obj IObject) *TPaintBox {
    p := new(TPaintBox)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPaintBox) Free() {
    if p.instance != 0 {
        PaintBox_Free(p.instance)
        p.instance = 0
    }
}

func (p *TPaintBox) Instance() uintptr {
    return p.instance
}

func (p *TPaintBox) IsValid() bool {
    return p.instance != 0
}

func (p *TPaintBox) BringToFront() {
    PaintBox_BringToFront(p.instance)
}

func (p *TPaintBox) HasParent() bool {
    return PaintBox_HasParent(p.instance)
}

func (p *TPaintBox) Hide() {
    PaintBox_Hide(p.instance)
}

func (p *TPaintBox) Invalidate() {
    PaintBox_Invalidate(p.instance)
}

func (p *TPaintBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return PaintBox_Perform(p.instance, Msg , WParam , LParam)
}

func (p *TPaintBox) Refresh() {
    PaintBox_Refresh(p.instance)
}

func (p *TPaintBox) Repaint() {
    PaintBox_Repaint(p.instance)
}

func (p *TPaintBox) SendToBack() {
    PaintBox_SendToBack(p.instance)
}

func (p *TPaintBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    PaintBox_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

func (p *TPaintBox) Show() {
    PaintBox_Show(p.instance)
}

func (p *TPaintBox) Update() {
    PaintBox_Update(p.instance)
}

func (p *TPaintBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return PaintBox_GetTextBuf(p.instance, Buffer , BufSize)
}

func (p *TPaintBox) GetTextLen() int32 {
    return PaintBox_GetTextLen(p.instance)
}

func (p *TPaintBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PaintBox_FindComponent(p.instance, AName))
}

func (p *TPaintBox) GetNamePath() string {
    return PaintBox_GetNamePath(p.instance)
}

func (p *TPaintBox) Assign(Source IObject) {
    PaintBox_Assign(p.instance, CheckPtr(Source))
}

func (p *TPaintBox) ClassName() string {
    return PaintBox_ClassName(p.instance)
}

func (p *TPaintBox) Equals(Obj IObject) bool {
    return PaintBox_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPaintBox) GetHashCode() int32 {
    return PaintBox_GetHashCode(p.instance)
}

func (p *TPaintBox) ToString() string {
    return PaintBox_ToString(p.instance)
}

func (p *TPaintBox) Canvas() *TCanvas {
    return CanvasFromInst(PaintBox_GetCanvas(p.instance))
}

func (p *TPaintBox) Align() TAlign {
    return PaintBox_GetAlign(p.instance)
}

func (p *TPaintBox) SetAlign(value TAlign) {
    PaintBox_SetAlign(p.instance, value)
}

func (p *TPaintBox) Anchors() TAnchors {
    return PaintBox_GetAnchors(p.instance)
}

func (p *TPaintBox) SetAnchors(value TAnchors) {
    PaintBox_SetAnchors(p.instance, value)
}

func (p *TPaintBox) Color() TColor {
    return PaintBox_GetColor(p.instance)
}

func (p *TPaintBox) SetColor(value TColor) {
    PaintBox_SetColor(p.instance, value)
}

func (p *TPaintBox) Enabled() bool {
    return PaintBox_GetEnabled(p.instance)
}

func (p *TPaintBox) SetEnabled(value bool) {
    PaintBox_SetEnabled(p.instance, value)
}

func (p *TPaintBox) Font() *TFont {
    return FontFromInst(PaintBox_GetFont(p.instance))
}

func (p *TPaintBox) SetFont(value *TFont) {
    PaintBox_SetFont(p.instance, CheckPtr(value))
}

func (p *TPaintBox) ParentColor() bool {
    return PaintBox_GetParentColor(p.instance)
}

func (p *TPaintBox) SetParentColor(value bool) {
    PaintBox_SetParentColor(p.instance, value)
}

func (p *TPaintBox) ParentFont() bool {
    return PaintBox_GetParentFont(p.instance)
}

func (p *TPaintBox) SetParentFont(value bool) {
    PaintBox_SetParentFont(p.instance, value)
}

func (p *TPaintBox) ParentShowHint() bool {
    return PaintBox_GetParentShowHint(p.instance)
}

func (p *TPaintBox) SetParentShowHint(value bool) {
    PaintBox_SetParentShowHint(p.instance, value)
}

func (p *TPaintBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(PaintBox_GetPopupMenu(p.instance))
}

func (p *TPaintBox) SetPopupMenu(value IComponent) {
    PaintBox_SetPopupMenu(p.instance, CheckPtr(value))
}

func (p *TPaintBox) ShowHint() bool {
    return PaintBox_GetShowHint(p.instance)
}

func (p *TPaintBox) SetShowHint(value bool) {
    PaintBox_SetShowHint(p.instance, value)
}

func (p *TPaintBox) Visible() bool {
    return PaintBox_GetVisible(p.instance)
}

func (p *TPaintBox) SetVisible(value bool) {
    PaintBox_SetVisible(p.instance, value)
}

func (p *TPaintBox) SetOnClick(fn TNotifyEvent) {
    PaintBox_SetOnClick(p.instance, fn)
}

func (p *TPaintBox) SetOnDblClick(fn TNotifyEvent) {
    PaintBox_SetOnDblClick(p.instance, fn)
}

func (p *TPaintBox) SetOnMouseDown(fn TMouseEvent) {
    PaintBox_SetOnMouseDown(p.instance, fn)
}

func (p *TPaintBox) SetOnMouseEnter(fn TNotifyEvent) {
    PaintBox_SetOnMouseEnter(p.instance, fn)
}

func (p *TPaintBox) SetOnMouseLeave(fn TNotifyEvent) {
    PaintBox_SetOnMouseLeave(p.instance, fn)
}

func (p *TPaintBox) SetOnMouseMove(fn TMouseMoveEvent) {
    PaintBox_SetOnMouseMove(p.instance, fn)
}

func (p *TPaintBox) SetOnMouseUp(fn TMouseEvent) {
    PaintBox_SetOnMouseUp(p.instance, fn)
}

func (p *TPaintBox) SetOnPaint(fn TNotifyEvent) {
    PaintBox_SetOnPaint(p.instance, fn)
}

func (p *TPaintBox) Action() *TAction {
    return ActionFromInst(PaintBox_GetAction(p.instance))
}

func (p *TPaintBox) SetAction(value IComponent) {
    PaintBox_SetAction(p.instance, CheckPtr(value))
}

func (p *TPaintBox) BiDiMode() TBiDiMode {
    return PaintBox_GetBiDiMode(p.instance)
}

func (p *TPaintBox) SetBiDiMode(value TBiDiMode) {
    PaintBox_SetBiDiMode(p.instance, value)
}

func (p *TPaintBox) BoundsRect() TRect {
    return PaintBox_GetBoundsRect(p.instance)
}

func (p *TPaintBox) SetBoundsRect(value TRect) {
    PaintBox_SetBoundsRect(p.instance, value)
}

func (p *TPaintBox) ClientHeight() int32 {
    return PaintBox_GetClientHeight(p.instance)
}

func (p *TPaintBox) SetClientHeight(value int32) {
    PaintBox_SetClientHeight(p.instance, value)
}

func (p *TPaintBox) ClientRect() TRect {
    return PaintBox_GetClientRect(p.instance)
}

func (p *TPaintBox) ClientWidth() int32 {
    return PaintBox_GetClientWidth(p.instance)
}

func (p *TPaintBox) SetClientWidth(value int32) {
    PaintBox_SetClientWidth(p.instance, value)
}

func (p *TPaintBox) ExplicitLeft() int32 {
    return PaintBox_GetExplicitLeft(p.instance)
}

func (p *TPaintBox) ExplicitTop() int32 {
    return PaintBox_GetExplicitTop(p.instance)
}

func (p *TPaintBox) ExplicitWidth() int32 {
    return PaintBox_GetExplicitWidth(p.instance)
}

func (p *TPaintBox) ExplicitHeight() int32 {
    return PaintBox_GetExplicitHeight(p.instance)
}

func (p *TPaintBox) Parent() *TControl {
    return ControlFromInst(PaintBox_GetParent(p.instance))
}

func (p *TPaintBox) SetParent(value IControl) {
    PaintBox_SetParent(p.instance, CheckPtr(value))
}

func (p *TPaintBox) StyleElements() TStyleElements {
    return PaintBox_GetStyleElements(p.instance)
}

func (p *TPaintBox) SetStyleElements(value TStyleElements) {
    PaintBox_SetStyleElements(p.instance, value)
}

func (p *TPaintBox) AlignWithMargins() bool {
    return PaintBox_GetAlignWithMargins(p.instance)
}

func (p *TPaintBox) SetAlignWithMargins(value bool) {
    PaintBox_SetAlignWithMargins(p.instance, value)
}

func (p *TPaintBox) Left() int32 {
    return PaintBox_GetLeft(p.instance)
}

func (p *TPaintBox) SetLeft(value int32) {
    PaintBox_SetLeft(p.instance, value)
}

func (p *TPaintBox) Top() int32 {
    return PaintBox_GetTop(p.instance)
}

func (p *TPaintBox) SetTop(value int32) {
    PaintBox_SetTop(p.instance, value)
}

func (p *TPaintBox) Width() int32 {
    return PaintBox_GetWidth(p.instance)
}

func (p *TPaintBox) SetWidth(value int32) {
    PaintBox_SetWidth(p.instance, value)
}

func (p *TPaintBox) Height() int32 {
    return PaintBox_GetHeight(p.instance)
}

func (p *TPaintBox) SetHeight(value int32) {
    PaintBox_SetHeight(p.instance, value)
}

func (p *TPaintBox) Cursor() TCursor {
    return PaintBox_GetCursor(p.instance)
}

func (p *TPaintBox) SetCursor(value TCursor) {
    PaintBox_SetCursor(p.instance, value)
}

func (p *TPaintBox) Hint() string {
    return PaintBox_GetHint(p.instance)
}

func (p *TPaintBox) SetHint(value string) {
    PaintBox_SetHint(p.instance, value)
}

func (p *TPaintBox) Margins() *TMargins {
    return MarginsFromInst(PaintBox_GetMargins(p.instance))
}

func (p *TPaintBox) SetMargins(value *TMargins) {
    PaintBox_SetMargins(p.instance, CheckPtr(value))
}

func (p *TPaintBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(PaintBox_GetCustomHint(p.instance))
}

func (p *TPaintBox) SetCustomHint(value IComponent) {
    PaintBox_SetCustomHint(p.instance, CheckPtr(value))
}

func (p *TPaintBox) ComponentCount() int32 {
    return PaintBox_GetComponentCount(p.instance)
}

func (p *TPaintBox) ComponentIndex() int32 {
    return PaintBox_GetComponentIndex(p.instance)
}

func (p *TPaintBox) SetComponentIndex(value int32) {
    PaintBox_SetComponentIndex(p.instance, value)
}

func (p *TPaintBox) Owner() *TComponent {
    return ComponentFromInst(PaintBox_GetOwner(p.instance))
}

func (p *TPaintBox) Name() string {
    return PaintBox_GetName(p.instance)
}

func (p *TPaintBox) SetName(value string) {
    PaintBox_SetName(p.instance, value)
}

func (p *TPaintBox) Tag() int {
    return PaintBox_GetTag(p.instance)
}

func (p *TPaintBox) SetTag(value int) {
    PaintBox_SetTag(p.instance, value)
}

func (p *TPaintBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PaintBox_GetComponents(p.instance, AIndex))
}

