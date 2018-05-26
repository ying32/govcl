
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

type TGauge struct {
    IControl
    instance uintptr
}

func NewGauge(owner IComponent) *TGauge {
    g := new(TGauge)
    g.instance = Gauge_Create(CheckPtr(owner))
    return g
}

func GaugeFromInst(inst uintptr) *TGauge {
    g := new(TGauge)
    g.instance = inst
    return g
}

func GaugeFromObj(obj IObject) *TGauge {
    g := new(TGauge)
    g.instance = CheckPtr(obj)
    return g
}

func (g *TGauge) Free() {
    if g.instance != 0 {
        Gauge_Free(g.instance)
        g.instance = 0
    }
}

func (g *TGauge) Instance() uintptr {
    return g.instance
}

func (g *TGauge) IsValid() bool {
    return g.instance != 0
}

func (g *TGauge) AddProgress(Value int32) {
    Gauge_AddProgress(g.instance, Value)
}

func (g *TGauge) BringToFront() {
    Gauge_BringToFront(g.instance)
}

func (g *TGauge) HasParent() bool {
    return Gauge_HasParent(g.instance)
}

func (g *TGauge) Hide() {
    Gauge_Hide(g.instance)
}

func (g *TGauge) Invalidate() {
    Gauge_Invalidate(g.instance)
}

func (g *TGauge) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Gauge_Perform(g.instance, Msg , WParam , LParam)
}

func (g *TGauge) Refresh() {
    Gauge_Refresh(g.instance)
}

func (g *TGauge) Repaint() {
    Gauge_Repaint(g.instance)
}

func (g *TGauge) SendToBack() {
    Gauge_SendToBack(g.instance)
}

func (g *TGauge) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Gauge_SetBounds(g.instance, ALeft , ATop , AWidth , AHeight)
}

func (g *TGauge) Show() {
    Gauge_Show(g.instance)
}

func (g *TGauge) Update() {
    Gauge_Update(g.instance)
}

func (g *TGauge) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Gauge_GetTextBuf(g.instance, Buffer , BufSize)
}

func (g *TGauge) GetTextLen() int32 {
    return Gauge_GetTextLen(g.instance)
}

func (g *TGauge) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Gauge_FindComponent(g.instance, AName))
}

func (g *TGauge) GetNamePath() string {
    return Gauge_GetNamePath(g.instance)
}

func (g *TGauge) Assign(Source IObject) {
    Gauge_Assign(g.instance, CheckPtr(Source))
}

func (g *TGauge) ClassName() string {
    return Gauge_ClassName(g.instance)
}

func (g *TGauge) Equals(Obj IObject) bool {
    return Gauge_Equals(g.instance, CheckPtr(Obj))
}

func (g *TGauge) GetHashCode() int32 {
    return Gauge_GetHashCode(g.instance)
}

func (g *TGauge) ToString() string {
    return Gauge_ToString(g.instance)
}

func (g *TGauge) PercentDone() int32 {
    return Gauge_GetPercentDone(g.instance)
}

func (g *TGauge) Align() TAlign {
    return Gauge_GetAlign(g.instance)
}

func (g *TGauge) SetAlign(value TAlign) {
    Gauge_SetAlign(g.instance, value)
}

func (g *TGauge) Anchors() TAnchors {
    return Gauge_GetAnchors(g.instance)
}

func (g *TGauge) SetAnchors(value TAnchors) {
    Gauge_SetAnchors(g.instance, value)
}

func (g *TGauge) BackColor() TColor {
    return Gauge_GetBackColor(g.instance)
}

func (g *TGauge) SetBackColor(value TColor) {
    Gauge_SetBackColor(g.instance, value)
}

func (g *TGauge) BorderStyle() TBorderStyle {
    return Gauge_GetBorderStyle(g.instance)
}

func (g *TGauge) SetBorderStyle(value TBorderStyle) {
    Gauge_SetBorderStyle(g.instance, value)
}

func (g *TGauge) Color() TColor {
    return Gauge_GetColor(g.instance)
}

func (g *TGauge) SetColor(value TColor) {
    Gauge_SetColor(g.instance, value)
}

func (g *TGauge) Enabled() bool {
    return Gauge_GetEnabled(g.instance)
}

func (g *TGauge) SetEnabled(value bool) {
    Gauge_SetEnabled(g.instance, value)
}

func (g *TGauge) ForeColor() TColor {
    return Gauge_GetForeColor(g.instance)
}

func (g *TGauge) SetForeColor(value TColor) {
    Gauge_SetForeColor(g.instance, value)
}

func (g *TGauge) Font() *TFont {
    return FontFromInst(Gauge_GetFont(g.instance))
}

func (g *TGauge) SetFont(value *TFont) {
    Gauge_SetFont(g.instance, CheckPtr(value))
}

func (g *TGauge) Kind() TGaugeKind {
    return Gauge_GetKind(g.instance)
}

func (g *TGauge) SetKind(value TGaugeKind) {
    Gauge_SetKind(g.instance, value)
}

func (g *TGauge) MinValue() int32 {
    return Gauge_GetMinValue(g.instance)
}

func (g *TGauge) SetMinValue(value int32) {
    Gauge_SetMinValue(g.instance, value)
}

func (g *TGauge) MaxValue() int32 {
    return Gauge_GetMaxValue(g.instance)
}

func (g *TGauge) SetMaxValue(value int32) {
    Gauge_SetMaxValue(g.instance, value)
}

func (g *TGauge) ParentColor() bool {
    return Gauge_GetParentColor(g.instance)
}

func (g *TGauge) SetParentColor(value bool) {
    Gauge_SetParentColor(g.instance, value)
}

func (g *TGauge) ParentFont() bool {
    return Gauge_GetParentFont(g.instance)
}

func (g *TGauge) SetParentFont(value bool) {
    Gauge_SetParentFont(g.instance, value)
}

func (g *TGauge) ParentShowHint() bool {
    return Gauge_GetParentShowHint(g.instance)
}

func (g *TGauge) SetParentShowHint(value bool) {
    Gauge_SetParentShowHint(g.instance, value)
}

func (g *TGauge) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Gauge_GetPopupMenu(g.instance))
}

func (g *TGauge) SetPopupMenu(value IComponent) {
    Gauge_SetPopupMenu(g.instance, CheckPtr(value))
}

func (g *TGauge) Progress() int32 {
    return Gauge_GetProgress(g.instance)
}

func (g *TGauge) SetProgress(value int32) {
    Gauge_SetProgress(g.instance, value)
}

func (g *TGauge) ShowHint() bool {
    return Gauge_GetShowHint(g.instance)
}

func (g *TGauge) SetShowHint(value bool) {
    Gauge_SetShowHint(g.instance, value)
}

func (g *TGauge) ShowText() bool {
    return Gauge_GetShowText(g.instance)
}

func (g *TGauge) SetShowText(value bool) {
    Gauge_SetShowText(g.instance, value)
}

func (g *TGauge) Visible() bool {
    return Gauge_GetVisible(g.instance)
}

func (g *TGauge) SetVisible(value bool) {
    Gauge_SetVisible(g.instance, value)
}

func (g *TGauge) Action() *TAction {
    return ActionFromInst(Gauge_GetAction(g.instance))
}

func (g *TGauge) SetAction(value IComponent) {
    Gauge_SetAction(g.instance, CheckPtr(value))
}

func (g *TGauge) BiDiMode() TBiDiMode {
    return Gauge_GetBiDiMode(g.instance)
}

func (g *TGauge) SetBiDiMode(value TBiDiMode) {
    Gauge_SetBiDiMode(g.instance, value)
}

func (g *TGauge) BoundsRect() TRect {
    return Gauge_GetBoundsRect(g.instance)
}

func (g *TGauge) SetBoundsRect(value TRect) {
    Gauge_SetBoundsRect(g.instance, value)
}

func (g *TGauge) ClientHeight() int32 {
    return Gauge_GetClientHeight(g.instance)
}

func (g *TGauge) SetClientHeight(value int32) {
    Gauge_SetClientHeight(g.instance, value)
}

func (g *TGauge) ClientRect() TRect {
    return Gauge_GetClientRect(g.instance)
}

func (g *TGauge) ClientWidth() int32 {
    return Gauge_GetClientWidth(g.instance)
}

func (g *TGauge) SetClientWidth(value int32) {
    Gauge_SetClientWidth(g.instance, value)
}

func (g *TGauge) ExplicitLeft() int32 {
    return Gauge_GetExplicitLeft(g.instance)
}

func (g *TGauge) ExplicitTop() int32 {
    return Gauge_GetExplicitTop(g.instance)
}

func (g *TGauge) ExplicitWidth() int32 {
    return Gauge_GetExplicitWidth(g.instance)
}

func (g *TGauge) ExplicitHeight() int32 {
    return Gauge_GetExplicitHeight(g.instance)
}

func (g *TGauge) Parent() *TControl {
    return ControlFromInst(Gauge_GetParent(g.instance))
}

func (g *TGauge) SetParent(value IControl) {
    Gauge_SetParent(g.instance, CheckPtr(value))
}

func (g *TGauge) StyleElements() TStyleElements {
    return Gauge_GetStyleElements(g.instance)
}

func (g *TGauge) SetStyleElements(value TStyleElements) {
    Gauge_SetStyleElements(g.instance, value)
}

func (g *TGauge) AlignWithMargins() bool {
    return Gauge_GetAlignWithMargins(g.instance)
}

func (g *TGauge) SetAlignWithMargins(value bool) {
    Gauge_SetAlignWithMargins(g.instance, value)
}

func (g *TGauge) Left() int32 {
    return Gauge_GetLeft(g.instance)
}

func (g *TGauge) SetLeft(value int32) {
    Gauge_SetLeft(g.instance, value)
}

func (g *TGauge) Top() int32 {
    return Gauge_GetTop(g.instance)
}

func (g *TGauge) SetTop(value int32) {
    Gauge_SetTop(g.instance, value)
}

func (g *TGauge) Width() int32 {
    return Gauge_GetWidth(g.instance)
}

func (g *TGauge) SetWidth(value int32) {
    Gauge_SetWidth(g.instance, value)
}

func (g *TGauge) Height() int32 {
    return Gauge_GetHeight(g.instance)
}

func (g *TGauge) SetHeight(value int32) {
    Gauge_SetHeight(g.instance, value)
}

func (g *TGauge) Cursor() TCursor {
    return Gauge_GetCursor(g.instance)
}

func (g *TGauge) SetCursor(value TCursor) {
    Gauge_SetCursor(g.instance, value)
}

func (g *TGauge) Hint() string {
    return Gauge_GetHint(g.instance)
}

func (g *TGauge) SetHint(value string) {
    Gauge_SetHint(g.instance, value)
}

func (g *TGauge) Margins() *TMargins {
    return MarginsFromInst(Gauge_GetMargins(g.instance))
}

func (g *TGauge) SetMargins(value *TMargins) {
    Gauge_SetMargins(g.instance, CheckPtr(value))
}

func (g *TGauge) CustomHint() *TCustomHint {
    return CustomHintFromInst(Gauge_GetCustomHint(g.instance))
}

func (g *TGauge) SetCustomHint(value IComponent) {
    Gauge_SetCustomHint(g.instance, CheckPtr(value))
}

func (g *TGauge) ComponentCount() int32 {
    return Gauge_GetComponentCount(g.instance)
}

func (g *TGauge) ComponentIndex() int32 {
    return Gauge_GetComponentIndex(g.instance)
}

func (g *TGauge) SetComponentIndex(value int32) {
    Gauge_SetComponentIndex(g.instance, value)
}

func (g *TGauge) Owner() *TComponent {
    return ComponentFromInst(Gauge_GetOwner(g.instance))
}

func (g *TGauge) Name() string {
    return Gauge_GetName(g.instance)
}

func (g *TGauge) SetName(value string) {
    Gauge_SetName(g.instance, value)
}

func (g *TGauge) Tag() int {
    return Gauge_GetTag(g.instance)
}

func (g *TGauge) SetTag(value int) {
    Gauge_SetTag(g.instance, value)
}

func (g *TGauge) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Gauge_GetComponents(g.instance, AIndex))
}

