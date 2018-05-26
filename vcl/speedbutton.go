
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

type TSpeedButton struct {
    IControl
    instance uintptr
}

func NewSpeedButton(owner IComponent) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = SpeedButton_Create(CheckPtr(owner))
    return s
}

func SpeedButtonFromInst(inst uintptr) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = inst
    return s
}

func SpeedButtonFromObj(obj IObject) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TSpeedButton) Free() {
    if s.instance != 0 {
        SpeedButton_Free(s.instance)
        s.instance = 0
    }
}

func (s *TSpeedButton) Instance() uintptr {
    return s.instance
}

func (s *TSpeedButton) IsValid() bool {
    return s.instance != 0
}

func (s *TSpeedButton) Click() {
    SpeedButton_Click(s.instance)
}

func (s *TSpeedButton) BringToFront() {
    SpeedButton_BringToFront(s.instance)
}

func (s *TSpeedButton) HasParent() bool {
    return SpeedButton_HasParent(s.instance)
}

func (s *TSpeedButton) Hide() {
    SpeedButton_Hide(s.instance)
}

func (s *TSpeedButton) Invalidate() {
    SpeedButton_Invalidate(s.instance)
}

func (s *TSpeedButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return SpeedButton_Perform(s.instance, Msg , WParam , LParam)
}

func (s *TSpeedButton) Refresh() {
    SpeedButton_Refresh(s.instance)
}

func (s *TSpeedButton) Repaint() {
    SpeedButton_Repaint(s.instance)
}

func (s *TSpeedButton) SendToBack() {
    SpeedButton_SendToBack(s.instance)
}

func (s *TSpeedButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    SpeedButton_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

func (s *TSpeedButton) Show() {
    SpeedButton_Show(s.instance)
}

func (s *TSpeedButton) Update() {
    SpeedButton_Update(s.instance)
}

func (s *TSpeedButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return SpeedButton_GetTextBuf(s.instance, Buffer , BufSize)
}

func (s *TSpeedButton) GetTextLen() int32 {
    return SpeedButton_GetTextLen(s.instance)
}

func (s *TSpeedButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(SpeedButton_FindComponent(s.instance, AName))
}

func (s *TSpeedButton) GetNamePath() string {
    return SpeedButton_GetNamePath(s.instance)
}

func (s *TSpeedButton) Assign(Source IObject) {
    SpeedButton_Assign(s.instance, CheckPtr(Source))
}

func (s *TSpeedButton) ClassName() string {
    return SpeedButton_ClassName(s.instance)
}

func (s *TSpeedButton) Equals(Obj IObject) bool {
    return SpeedButton_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSpeedButton) GetHashCode() int32 {
    return SpeedButton_GetHashCode(s.instance)
}

func (s *TSpeedButton) ToString() string {
    return SpeedButton_ToString(s.instance)
}

func (s *TSpeedButton) Action() *TAction {
    return ActionFromInst(SpeedButton_GetAction(s.instance))
}

func (s *TSpeedButton) SetAction(value IComponent) {
    SpeedButton_SetAction(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) Align() TAlign {
    return SpeedButton_GetAlign(s.instance)
}

func (s *TSpeedButton) SetAlign(value TAlign) {
    SpeedButton_SetAlign(s.instance, value)
}

func (s *TSpeedButton) AllowAllUp() bool {
    return SpeedButton_GetAllowAllUp(s.instance)
}

func (s *TSpeedButton) SetAllowAllUp(value bool) {
    SpeedButton_SetAllowAllUp(s.instance, value)
}

func (s *TSpeedButton) Anchors() TAnchors {
    return SpeedButton_GetAnchors(s.instance)
}

func (s *TSpeedButton) SetAnchors(value TAnchors) {
    SpeedButton_SetAnchors(s.instance, value)
}

func (s *TSpeedButton) BiDiMode() TBiDiMode {
    return SpeedButton_GetBiDiMode(s.instance)
}

func (s *TSpeedButton) SetBiDiMode(value TBiDiMode) {
    SpeedButton_SetBiDiMode(s.instance, value)
}

func (s *TSpeedButton) GroupIndex() int32 {
    return SpeedButton_GetGroupIndex(s.instance)
}

func (s *TSpeedButton) SetGroupIndex(value int32) {
    SpeedButton_SetGroupIndex(s.instance, value)
}

func (s *TSpeedButton) Down() bool {
    return SpeedButton_GetDown(s.instance)
}

func (s *TSpeedButton) SetDown(value bool) {
    SpeedButton_SetDown(s.instance, value)
}

func (s *TSpeedButton) Caption() string {
    return SpeedButton_GetCaption(s.instance)
}

func (s *TSpeedButton) SetCaption(value string) {
    SpeedButton_SetCaption(s.instance, value)
}

func (s *TSpeedButton) Enabled() bool {
    return SpeedButton_GetEnabled(s.instance)
}

func (s *TSpeedButton) SetEnabled(value bool) {
    SpeedButton_SetEnabled(s.instance, value)
}

func (s *TSpeedButton) Flat() bool {
    return SpeedButton_GetFlat(s.instance)
}

func (s *TSpeedButton) SetFlat(value bool) {
    SpeedButton_SetFlat(s.instance, value)
}

func (s *TSpeedButton) Font() *TFont {
    return FontFromInst(SpeedButton_GetFont(s.instance))
}

func (s *TSpeedButton) SetFont(value *TFont) {
    SpeedButton_SetFont(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) Glyph() *TBitmap {
    return BitmapFromInst(SpeedButton_GetGlyph(s.instance))
}

func (s *TSpeedButton) SetGlyph(value *TBitmap) {
    SpeedButton_SetGlyph(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) Layout() TButtonLayout {
    return SpeedButton_GetLayout(s.instance)
}

func (s *TSpeedButton) SetLayout(value TButtonLayout) {
    SpeedButton_SetLayout(s.instance, value)
}

func (s *TSpeedButton) NumGlyphs() TNumGlyphs {
    return SpeedButton_GetNumGlyphs(s.instance)
}

func (s *TSpeedButton) SetNumGlyphs(value TNumGlyphs) {
    SpeedButton_SetNumGlyphs(s.instance, value)
}

func (s *TSpeedButton) ParentFont() bool {
    return SpeedButton_GetParentFont(s.instance)
}

func (s *TSpeedButton) SetParentFont(value bool) {
    SpeedButton_SetParentFont(s.instance, value)
}

func (s *TSpeedButton) ParentShowHint() bool {
    return SpeedButton_GetParentShowHint(s.instance)
}

func (s *TSpeedButton) SetParentShowHint(value bool) {
    SpeedButton_SetParentShowHint(s.instance, value)
}

func (s *TSpeedButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(SpeedButton_GetPopupMenu(s.instance))
}

func (s *TSpeedButton) SetPopupMenu(value IComponent) {
    SpeedButton_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) ShowHint() bool {
    return SpeedButton_GetShowHint(s.instance)
}

func (s *TSpeedButton) SetShowHint(value bool) {
    SpeedButton_SetShowHint(s.instance, value)
}

func (s *TSpeedButton) Spacing() int32 {
    return SpeedButton_GetSpacing(s.instance)
}

func (s *TSpeedButton) SetSpacing(value int32) {
    SpeedButton_SetSpacing(s.instance, value)
}

func (s *TSpeedButton) Transparent() bool {
    return SpeedButton_GetTransparent(s.instance)
}

func (s *TSpeedButton) SetTransparent(value bool) {
    SpeedButton_SetTransparent(s.instance, value)
}

func (s *TSpeedButton) Visible() bool {
    return SpeedButton_GetVisible(s.instance)
}

func (s *TSpeedButton) SetVisible(value bool) {
    SpeedButton_SetVisible(s.instance, value)
}

func (s *TSpeedButton) StyleElements() TStyleElements {
    return SpeedButton_GetStyleElements(s.instance)
}

func (s *TSpeedButton) SetStyleElements(value TStyleElements) {
    SpeedButton_SetStyleElements(s.instance, value)
}

func (s *TSpeedButton) SetOnClick(fn TNotifyEvent) {
    SpeedButton_SetOnClick(s.instance, fn)
}

func (s *TSpeedButton) SetOnDblClick(fn TNotifyEvent) {
    SpeedButton_SetOnDblClick(s.instance, fn)
}

func (s *TSpeedButton) SetOnMouseDown(fn TMouseEvent) {
    SpeedButton_SetOnMouseDown(s.instance, fn)
}

func (s *TSpeedButton) SetOnMouseEnter(fn TNotifyEvent) {
    SpeedButton_SetOnMouseEnter(s.instance, fn)
}

func (s *TSpeedButton) SetOnMouseLeave(fn TNotifyEvent) {
    SpeedButton_SetOnMouseLeave(s.instance, fn)
}

func (s *TSpeedButton) SetOnMouseMove(fn TMouseMoveEvent) {
    SpeedButton_SetOnMouseMove(s.instance, fn)
}

func (s *TSpeedButton) SetOnMouseUp(fn TMouseEvent) {
    SpeedButton_SetOnMouseUp(s.instance, fn)
}

func (s *TSpeedButton) BoundsRect() TRect {
    return SpeedButton_GetBoundsRect(s.instance)
}

func (s *TSpeedButton) SetBoundsRect(value TRect) {
    SpeedButton_SetBoundsRect(s.instance, value)
}

func (s *TSpeedButton) ClientHeight() int32 {
    return SpeedButton_GetClientHeight(s.instance)
}

func (s *TSpeedButton) SetClientHeight(value int32) {
    SpeedButton_SetClientHeight(s.instance, value)
}

func (s *TSpeedButton) ClientRect() TRect {
    return SpeedButton_GetClientRect(s.instance)
}

func (s *TSpeedButton) ClientWidth() int32 {
    return SpeedButton_GetClientWidth(s.instance)
}

func (s *TSpeedButton) SetClientWidth(value int32) {
    SpeedButton_SetClientWidth(s.instance, value)
}

func (s *TSpeedButton) ExplicitLeft() int32 {
    return SpeedButton_GetExplicitLeft(s.instance)
}

func (s *TSpeedButton) ExplicitTop() int32 {
    return SpeedButton_GetExplicitTop(s.instance)
}

func (s *TSpeedButton) ExplicitWidth() int32 {
    return SpeedButton_GetExplicitWidth(s.instance)
}

func (s *TSpeedButton) ExplicitHeight() int32 {
    return SpeedButton_GetExplicitHeight(s.instance)
}

func (s *TSpeedButton) Parent() *TControl {
    return ControlFromInst(SpeedButton_GetParent(s.instance))
}

func (s *TSpeedButton) SetParent(value IControl) {
    SpeedButton_SetParent(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) AlignWithMargins() bool {
    return SpeedButton_GetAlignWithMargins(s.instance)
}

func (s *TSpeedButton) SetAlignWithMargins(value bool) {
    SpeedButton_SetAlignWithMargins(s.instance, value)
}

func (s *TSpeedButton) Left() int32 {
    return SpeedButton_GetLeft(s.instance)
}

func (s *TSpeedButton) SetLeft(value int32) {
    SpeedButton_SetLeft(s.instance, value)
}

func (s *TSpeedButton) Top() int32 {
    return SpeedButton_GetTop(s.instance)
}

func (s *TSpeedButton) SetTop(value int32) {
    SpeedButton_SetTop(s.instance, value)
}

func (s *TSpeedButton) Width() int32 {
    return SpeedButton_GetWidth(s.instance)
}

func (s *TSpeedButton) SetWidth(value int32) {
    SpeedButton_SetWidth(s.instance, value)
}

func (s *TSpeedButton) Height() int32 {
    return SpeedButton_GetHeight(s.instance)
}

func (s *TSpeedButton) SetHeight(value int32) {
    SpeedButton_SetHeight(s.instance, value)
}

func (s *TSpeedButton) Cursor() TCursor {
    return SpeedButton_GetCursor(s.instance)
}

func (s *TSpeedButton) SetCursor(value TCursor) {
    SpeedButton_SetCursor(s.instance, value)
}

func (s *TSpeedButton) Hint() string {
    return SpeedButton_GetHint(s.instance)
}

func (s *TSpeedButton) SetHint(value string) {
    SpeedButton_SetHint(s.instance, value)
}

func (s *TSpeedButton) Margins() *TMargins {
    return MarginsFromInst(SpeedButton_GetMargins(s.instance))
}

func (s *TSpeedButton) SetMargins(value *TMargins) {
    SpeedButton_SetMargins(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(SpeedButton_GetCustomHint(s.instance))
}

func (s *TSpeedButton) SetCustomHint(value IComponent) {
    SpeedButton_SetCustomHint(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) ComponentCount() int32 {
    return SpeedButton_GetComponentCount(s.instance)
}

func (s *TSpeedButton) ComponentIndex() int32 {
    return SpeedButton_GetComponentIndex(s.instance)
}

func (s *TSpeedButton) SetComponentIndex(value int32) {
    SpeedButton_SetComponentIndex(s.instance, value)
}

func (s *TSpeedButton) Owner() *TComponent {
    return ComponentFromInst(SpeedButton_GetOwner(s.instance))
}

func (s *TSpeedButton) Name() string {
    return SpeedButton_GetName(s.instance)
}

func (s *TSpeedButton) SetName(value string) {
    SpeedButton_SetName(s.instance, value)
}

func (s *TSpeedButton) Tag() int {
    return SpeedButton_GetTag(s.instance)
}

func (s *TSpeedButton) SetTag(value int) {
    SpeedButton_SetTag(s.instance, value)
}

func (s *TSpeedButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(SpeedButton_GetComponents(s.instance, AIndex))
}

