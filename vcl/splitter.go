
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

type TSplitter struct {
    IControl
    instance uintptr
}

func NewSplitter(owner IComponent) *TSplitter {
    s := new(TSplitter)
    s.instance = Splitter_Create(CheckPtr(owner))
    return s
}

func SplitterFromInst(inst uintptr) *TSplitter {
    s := new(TSplitter)
    s.instance = inst
    return s
}

func SplitterFromObj(obj IObject) *TSplitter {
    s := new(TSplitter)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TSplitter) Free() {
    if s.instance != 0 {
        Splitter_Free(s.instance)
        s.instance = 0
    }
}

func (s *TSplitter) Instance() uintptr {
    return s.instance
}

func (s *TSplitter) IsValid() bool {
    return s.instance != 0
}

func (s *TSplitter) BringToFront() {
    Splitter_BringToFront(s.instance)
}

func (s *TSplitter) HasParent() bool {
    return Splitter_HasParent(s.instance)
}

func (s *TSplitter) Hide() {
    Splitter_Hide(s.instance)
}

func (s *TSplitter) Invalidate() {
    Splitter_Invalidate(s.instance)
}

func (s *TSplitter) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Splitter_Perform(s.instance, Msg , WParam , LParam)
}

func (s *TSplitter) Refresh() {
    Splitter_Refresh(s.instance)
}

func (s *TSplitter) Repaint() {
    Splitter_Repaint(s.instance)
}

func (s *TSplitter) SendToBack() {
    Splitter_SendToBack(s.instance)
}

func (s *TSplitter) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Splitter_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

func (s *TSplitter) Show() {
    Splitter_Show(s.instance)
}

func (s *TSplitter) Update() {
    Splitter_Update(s.instance)
}

func (s *TSplitter) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Splitter_GetTextBuf(s.instance, Buffer , BufSize)
}

func (s *TSplitter) GetTextLen() int32 {
    return Splitter_GetTextLen(s.instance)
}

func (s *TSplitter) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Splitter_FindComponent(s.instance, AName))
}

func (s *TSplitter) GetNamePath() string {
    return Splitter_GetNamePath(s.instance)
}

func (s *TSplitter) Assign(Source IObject) {
    Splitter_Assign(s.instance, CheckPtr(Source))
}

func (s *TSplitter) ClassName() string {
    return Splitter_ClassName(s.instance)
}

func (s *TSplitter) Equals(Obj IObject) bool {
    return Splitter_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSplitter) GetHashCode() int32 {
    return Splitter_GetHashCode(s.instance)
}

func (s *TSplitter) ToString() string {
    return Splitter_ToString(s.instance)
}

func (s *TSplitter) Canvas() *TCanvas {
    return CanvasFromInst(Splitter_GetCanvas(s.instance))
}

func (s *TSplitter) Align() TAlign {
    return Splitter_GetAlign(s.instance)
}

func (s *TSplitter) SetAlign(value TAlign) {
    Splitter_SetAlign(s.instance, value)
}

func (s *TSplitter) Color() TColor {
    return Splitter_GetColor(s.instance)
}

func (s *TSplitter) SetColor(value TColor) {
    Splitter_SetColor(s.instance, value)
}

func (s *TSplitter) Cursor() TCursor {
    return Splitter_GetCursor(s.instance)
}

func (s *TSplitter) SetCursor(value TCursor) {
    Splitter_SetCursor(s.instance, value)
}

func (s *TSplitter) ParentColor() bool {
    return Splitter_GetParentColor(s.instance)
}

func (s *TSplitter) SetParentColor(value bool) {
    Splitter_SetParentColor(s.instance, value)
}

func (s *TSplitter) Visible() bool {
    return Splitter_GetVisible(s.instance)
}

func (s *TSplitter) SetVisible(value bool) {
    Splitter_SetVisible(s.instance, value)
}

func (s *TSplitter) Width() int32 {
    return Splitter_GetWidth(s.instance)
}

func (s *TSplitter) SetWidth(value int32) {
    Splitter_SetWidth(s.instance, value)
}

func (s *TSplitter) StyleElements() TStyleElements {
    return Splitter_GetStyleElements(s.instance)
}

func (s *TSplitter) SetStyleElements(value TStyleElements) {
    Splitter_SetStyleElements(s.instance, value)
}

func (s *TSplitter) SetOnPaint(fn TNotifyEvent) {
    Splitter_SetOnPaint(s.instance, fn)
}

func (s *TSplitter) Enabled() bool {
    return Splitter_GetEnabled(s.instance)
}

func (s *TSplitter) SetEnabled(value bool) {
    Splitter_SetEnabled(s.instance, value)
}

func (s *TSplitter) Action() *TAction {
    return ActionFromInst(Splitter_GetAction(s.instance))
}

func (s *TSplitter) SetAction(value IComponent) {
    Splitter_SetAction(s.instance, CheckPtr(value))
}

func (s *TSplitter) Anchors() TAnchors {
    return Splitter_GetAnchors(s.instance)
}

func (s *TSplitter) SetAnchors(value TAnchors) {
    Splitter_SetAnchors(s.instance, value)
}

func (s *TSplitter) BiDiMode() TBiDiMode {
    return Splitter_GetBiDiMode(s.instance)
}

func (s *TSplitter) SetBiDiMode(value TBiDiMode) {
    Splitter_SetBiDiMode(s.instance, value)
}

func (s *TSplitter) BoundsRect() TRect {
    return Splitter_GetBoundsRect(s.instance)
}

func (s *TSplitter) SetBoundsRect(value TRect) {
    Splitter_SetBoundsRect(s.instance, value)
}

func (s *TSplitter) ClientHeight() int32 {
    return Splitter_GetClientHeight(s.instance)
}

func (s *TSplitter) SetClientHeight(value int32) {
    Splitter_SetClientHeight(s.instance, value)
}

func (s *TSplitter) ClientRect() TRect {
    return Splitter_GetClientRect(s.instance)
}

func (s *TSplitter) ClientWidth() int32 {
    return Splitter_GetClientWidth(s.instance)
}

func (s *TSplitter) SetClientWidth(value int32) {
    Splitter_SetClientWidth(s.instance, value)
}

func (s *TSplitter) ExplicitLeft() int32 {
    return Splitter_GetExplicitLeft(s.instance)
}

func (s *TSplitter) ExplicitTop() int32 {
    return Splitter_GetExplicitTop(s.instance)
}

func (s *TSplitter) ExplicitWidth() int32 {
    return Splitter_GetExplicitWidth(s.instance)
}

func (s *TSplitter) ExplicitHeight() int32 {
    return Splitter_GetExplicitHeight(s.instance)
}

func (s *TSplitter) ShowHint() bool {
    return Splitter_GetShowHint(s.instance)
}

func (s *TSplitter) SetShowHint(value bool) {
    Splitter_SetShowHint(s.instance, value)
}

func (s *TSplitter) Parent() *TControl {
    return ControlFromInst(Splitter_GetParent(s.instance))
}

func (s *TSplitter) SetParent(value IControl) {
    Splitter_SetParent(s.instance, CheckPtr(value))
}

func (s *TSplitter) AlignWithMargins() bool {
    return Splitter_GetAlignWithMargins(s.instance)
}

func (s *TSplitter) SetAlignWithMargins(value bool) {
    Splitter_SetAlignWithMargins(s.instance, value)
}

func (s *TSplitter) Left() int32 {
    return Splitter_GetLeft(s.instance)
}

func (s *TSplitter) SetLeft(value int32) {
    Splitter_SetLeft(s.instance, value)
}

func (s *TSplitter) Top() int32 {
    return Splitter_GetTop(s.instance)
}

func (s *TSplitter) SetTop(value int32) {
    Splitter_SetTop(s.instance, value)
}

func (s *TSplitter) Height() int32 {
    return Splitter_GetHeight(s.instance)
}

func (s *TSplitter) SetHeight(value int32) {
    Splitter_SetHeight(s.instance, value)
}

func (s *TSplitter) Hint() string {
    return Splitter_GetHint(s.instance)
}

func (s *TSplitter) SetHint(value string) {
    Splitter_SetHint(s.instance, value)
}

func (s *TSplitter) Margins() *TMargins {
    return MarginsFromInst(Splitter_GetMargins(s.instance))
}

func (s *TSplitter) SetMargins(value *TMargins) {
    Splitter_SetMargins(s.instance, CheckPtr(value))
}

func (s *TSplitter) CustomHint() *TCustomHint {
    return CustomHintFromInst(Splitter_GetCustomHint(s.instance))
}

func (s *TSplitter) SetCustomHint(value IComponent) {
    Splitter_SetCustomHint(s.instance, CheckPtr(value))
}

func (s *TSplitter) ComponentCount() int32 {
    return Splitter_GetComponentCount(s.instance)
}

func (s *TSplitter) ComponentIndex() int32 {
    return Splitter_GetComponentIndex(s.instance)
}

func (s *TSplitter) SetComponentIndex(value int32) {
    Splitter_SetComponentIndex(s.instance, value)
}

func (s *TSplitter) Owner() *TComponent {
    return ComponentFromInst(Splitter_GetOwner(s.instance))
}

func (s *TSplitter) Name() string {
    return Splitter_GetName(s.instance)
}

func (s *TSplitter) SetName(value string) {
    Splitter_SetName(s.instance, value)
}

func (s *TSplitter) Tag() int {
    return Splitter_GetTag(s.instance)
}

func (s *TSplitter) SetTag(value int) {
    Splitter_SetTag(s.instance, value)
}

func (s *TSplitter) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Splitter_GetComponents(s.instance, AIndex))
}

