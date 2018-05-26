
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

type TBevel struct {
    IControl
    instance uintptr
}

func NewBevel(owner IComponent) *TBevel {
    b := new(TBevel)
    b.instance = Bevel_Create(CheckPtr(owner))
    return b
}

func BevelFromInst(inst uintptr) *TBevel {
    b := new(TBevel)
    b.instance = inst
    return b
}

func BevelFromObj(obj IObject) *TBevel {
    b := new(TBevel)
    b.instance = CheckPtr(obj)
    return b
}

func (b *TBevel) Free() {
    if b.instance != 0 {
        Bevel_Free(b.instance)
        b.instance = 0
    }
}

func (b *TBevel) Instance() uintptr {
    return b.instance
}

func (b *TBevel) IsValid() bool {
    return b.instance != 0
}

func (b *TBevel) BringToFront() {
    Bevel_BringToFront(b.instance)
}

func (b *TBevel) HasParent() bool {
    return Bevel_HasParent(b.instance)
}

func (b *TBevel) Hide() {
    Bevel_Hide(b.instance)
}

func (b *TBevel) Invalidate() {
    Bevel_Invalidate(b.instance)
}

func (b *TBevel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Bevel_Perform(b.instance, Msg , WParam , LParam)
}

func (b *TBevel) Refresh() {
    Bevel_Refresh(b.instance)
}

func (b *TBevel) Repaint() {
    Bevel_Repaint(b.instance)
}

func (b *TBevel) SendToBack() {
    Bevel_SendToBack(b.instance)
}

func (b *TBevel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Bevel_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

func (b *TBevel) Show() {
    Bevel_Show(b.instance)
}

func (b *TBevel) Update() {
    Bevel_Update(b.instance)
}

func (b *TBevel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Bevel_GetTextBuf(b.instance, Buffer , BufSize)
}

func (b *TBevel) GetTextLen() int32 {
    return Bevel_GetTextLen(b.instance)
}

func (b *TBevel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Bevel_FindComponent(b.instance, AName))
}

func (b *TBevel) GetNamePath() string {
    return Bevel_GetNamePath(b.instance)
}

func (b *TBevel) Assign(Source IObject) {
    Bevel_Assign(b.instance, CheckPtr(Source))
}

func (b *TBevel) ClassName() string {
    return Bevel_ClassName(b.instance)
}

func (b *TBevel) Equals(Obj IObject) bool {
    return Bevel_Equals(b.instance, CheckPtr(Obj))
}

func (b *TBevel) GetHashCode() int32 {
    return Bevel_GetHashCode(b.instance)
}

func (b *TBevel) ToString() string {
    return Bevel_ToString(b.instance)
}

func (b *TBevel) Align() TAlign {
    return Bevel_GetAlign(b.instance)
}

func (b *TBevel) SetAlign(value TAlign) {
    Bevel_SetAlign(b.instance, value)
}

func (b *TBevel) Anchors() TAnchors {
    return Bevel_GetAnchors(b.instance)
}

func (b *TBevel) SetAnchors(value TAnchors) {
    Bevel_SetAnchors(b.instance, value)
}

func (b *TBevel) ParentShowHint() bool {
    return Bevel_GetParentShowHint(b.instance)
}

func (b *TBevel) SetParentShowHint(value bool) {
    Bevel_SetParentShowHint(b.instance, value)
}

func (b *TBevel) Shape() TBevelShape {
    return Bevel_GetShape(b.instance)
}

func (b *TBevel) SetShape(value TBevelShape) {
    Bevel_SetShape(b.instance, value)
}

func (b *TBevel) ShowHint() bool {
    return Bevel_GetShowHint(b.instance)
}

func (b *TBevel) SetShowHint(value bool) {
    Bevel_SetShowHint(b.instance, value)
}

func (b *TBevel) Style() TBevelStyle {
    return Bevel_GetStyle(b.instance)
}

func (b *TBevel) SetStyle(value TBevelStyle) {
    Bevel_SetStyle(b.instance, value)
}

func (b *TBevel) Visible() bool {
    return Bevel_GetVisible(b.instance)
}

func (b *TBevel) SetVisible(value bool) {
    Bevel_SetVisible(b.instance, value)
}

func (b *TBevel) Enabled() bool {
    return Bevel_GetEnabled(b.instance)
}

func (b *TBevel) SetEnabled(value bool) {
    Bevel_SetEnabled(b.instance, value)
}

func (b *TBevel) Action() *TAction {
    return ActionFromInst(Bevel_GetAction(b.instance))
}

func (b *TBevel) SetAction(value IComponent) {
    Bevel_SetAction(b.instance, CheckPtr(value))
}

func (b *TBevel) BiDiMode() TBiDiMode {
    return Bevel_GetBiDiMode(b.instance)
}

func (b *TBevel) SetBiDiMode(value TBiDiMode) {
    Bevel_SetBiDiMode(b.instance, value)
}

func (b *TBevel) BoundsRect() TRect {
    return Bevel_GetBoundsRect(b.instance)
}

func (b *TBevel) SetBoundsRect(value TRect) {
    Bevel_SetBoundsRect(b.instance, value)
}

func (b *TBevel) ClientHeight() int32 {
    return Bevel_GetClientHeight(b.instance)
}

func (b *TBevel) SetClientHeight(value int32) {
    Bevel_SetClientHeight(b.instance, value)
}

func (b *TBevel) ClientRect() TRect {
    return Bevel_GetClientRect(b.instance)
}

func (b *TBevel) ClientWidth() int32 {
    return Bevel_GetClientWidth(b.instance)
}

func (b *TBevel) SetClientWidth(value int32) {
    Bevel_SetClientWidth(b.instance, value)
}

func (b *TBevel) ExplicitLeft() int32 {
    return Bevel_GetExplicitLeft(b.instance)
}

func (b *TBevel) ExplicitTop() int32 {
    return Bevel_GetExplicitTop(b.instance)
}

func (b *TBevel) ExplicitWidth() int32 {
    return Bevel_GetExplicitWidth(b.instance)
}

func (b *TBevel) ExplicitHeight() int32 {
    return Bevel_GetExplicitHeight(b.instance)
}

func (b *TBevel) Parent() *TControl {
    return ControlFromInst(Bevel_GetParent(b.instance))
}

func (b *TBevel) SetParent(value IControl) {
    Bevel_SetParent(b.instance, CheckPtr(value))
}

func (b *TBevel) StyleElements() TStyleElements {
    return Bevel_GetStyleElements(b.instance)
}

func (b *TBevel) SetStyleElements(value TStyleElements) {
    Bevel_SetStyleElements(b.instance, value)
}

func (b *TBevel) AlignWithMargins() bool {
    return Bevel_GetAlignWithMargins(b.instance)
}

func (b *TBevel) SetAlignWithMargins(value bool) {
    Bevel_SetAlignWithMargins(b.instance, value)
}

func (b *TBevel) Left() int32 {
    return Bevel_GetLeft(b.instance)
}

func (b *TBevel) SetLeft(value int32) {
    Bevel_SetLeft(b.instance, value)
}

func (b *TBevel) Top() int32 {
    return Bevel_GetTop(b.instance)
}

func (b *TBevel) SetTop(value int32) {
    Bevel_SetTop(b.instance, value)
}

func (b *TBevel) Width() int32 {
    return Bevel_GetWidth(b.instance)
}

func (b *TBevel) SetWidth(value int32) {
    Bevel_SetWidth(b.instance, value)
}

func (b *TBevel) Height() int32 {
    return Bevel_GetHeight(b.instance)
}

func (b *TBevel) SetHeight(value int32) {
    Bevel_SetHeight(b.instance, value)
}

func (b *TBevel) Cursor() TCursor {
    return Bevel_GetCursor(b.instance)
}

func (b *TBevel) SetCursor(value TCursor) {
    Bevel_SetCursor(b.instance, value)
}

func (b *TBevel) Hint() string {
    return Bevel_GetHint(b.instance)
}

func (b *TBevel) SetHint(value string) {
    Bevel_SetHint(b.instance, value)
}

func (b *TBevel) Margins() *TMargins {
    return MarginsFromInst(Bevel_GetMargins(b.instance))
}

func (b *TBevel) SetMargins(value *TMargins) {
    Bevel_SetMargins(b.instance, CheckPtr(value))
}

func (b *TBevel) CustomHint() *TCustomHint {
    return CustomHintFromInst(Bevel_GetCustomHint(b.instance))
}

func (b *TBevel) SetCustomHint(value IComponent) {
    Bevel_SetCustomHint(b.instance, CheckPtr(value))
}

func (b *TBevel) ComponentCount() int32 {
    return Bevel_GetComponentCount(b.instance)
}

func (b *TBevel) ComponentIndex() int32 {
    return Bevel_GetComponentIndex(b.instance)
}

func (b *TBevel) SetComponentIndex(value int32) {
    Bevel_SetComponentIndex(b.instance, value)
}

func (b *TBevel) Owner() *TComponent {
    return ComponentFromInst(Bevel_GetOwner(b.instance))
}

func (b *TBevel) Name() string {
    return Bevel_GetName(b.instance)
}

func (b *TBevel) SetName(value string) {
    Bevel_SetName(b.instance, value)
}

func (b *TBevel) Tag() int {
    return Bevel_GetTag(b.instance)
}

func (b *TBevel) SetTag(value int) {
    Bevel_SetTag(b.instance, value)
}

func (b *TBevel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Bevel_GetComponents(b.instance, AIndex))
}

