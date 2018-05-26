
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

type TControl struct {
    IControl
    instance uintptr
}

func NewControl(owner IComponent) *TControl {
    c := new(TControl)
    c.instance = Control_Create(CheckPtr(owner))
    return c
}

func ControlFromInst(inst uintptr) *TControl {
    c := new(TControl)
    c.instance = inst
    return c
}

func ControlFromObj(obj IObject) *TControl {
    c := new(TControl)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TControl) Free() {
    if c.instance != 0 {
        Control_Free(c.instance)
        c.instance = 0
    }
}

func (c *TControl) Instance() uintptr {
    return c.instance
}

func (c *TControl) IsValid() bool {
    return c.instance != 0
}

func (c *TControl) BringToFront() {
    Control_BringToFront(c.instance)
}

func (c *TControl) HasParent() bool {
    return Control_HasParent(c.instance)
}

func (c *TControl) Hide() {
    Control_Hide(c.instance)
}

func (c *TControl) Invalidate() {
    Control_Invalidate(c.instance)
}

func (c *TControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Control_Perform(c.instance, Msg , WParam , LParam)
}

func (c *TControl) Refresh() {
    Control_Refresh(c.instance)
}

func (c *TControl) Repaint() {
    Control_Repaint(c.instance)
}

func (c *TControl) SendToBack() {
    Control_SendToBack(c.instance)
}

func (c *TControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Control_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

func (c *TControl) Show() {
    Control_Show(c.instance)
}

func (c *TControl) Update() {
    Control_Update(c.instance)
}

func (c *TControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Control_GetTextBuf(c.instance, Buffer , BufSize)
}

func (c *TControl) GetTextLen() int32 {
    return Control_GetTextLen(c.instance)
}

func (c *TControl) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Control_FindComponent(c.instance, AName))
}

func (c *TControl) GetNamePath() string {
    return Control_GetNamePath(c.instance)
}

func (c *TControl) Assign(Source IObject) {
    Control_Assign(c.instance, CheckPtr(Source))
}

func (c *TControl) ClassName() string {
    return Control_ClassName(c.instance)
}

func (c *TControl) Equals(Obj IObject) bool {
    return Control_Equals(c.instance, CheckPtr(Obj))
}

func (c *TControl) GetHashCode() int32 {
    return Control_GetHashCode(c.instance)
}

func (c *TControl) ToString() string {
    return Control_ToString(c.instance)
}

func (c *TControl) Enabled() bool {
    return Control_GetEnabled(c.instance)
}

func (c *TControl) SetEnabled(value bool) {
    Control_SetEnabled(c.instance, value)
}

func (c *TControl) Action() *TAction {
    return ActionFromInst(Control_GetAction(c.instance))
}

func (c *TControl) SetAction(value IComponent) {
    Control_SetAction(c.instance, CheckPtr(value))
}

func (c *TControl) Align() TAlign {
    return Control_GetAlign(c.instance)
}

func (c *TControl) SetAlign(value TAlign) {
    Control_SetAlign(c.instance, value)
}

func (c *TControl) Anchors() TAnchors {
    return Control_GetAnchors(c.instance)
}

func (c *TControl) SetAnchors(value TAnchors) {
    Control_SetAnchors(c.instance, value)
}

func (c *TControl) BiDiMode() TBiDiMode {
    return Control_GetBiDiMode(c.instance)
}

func (c *TControl) SetBiDiMode(value TBiDiMode) {
    Control_SetBiDiMode(c.instance, value)
}

func (c *TControl) BoundsRect() TRect {
    return Control_GetBoundsRect(c.instance)
}

func (c *TControl) SetBoundsRect(value TRect) {
    Control_SetBoundsRect(c.instance, value)
}

func (c *TControl) ClientHeight() int32 {
    return Control_GetClientHeight(c.instance)
}

func (c *TControl) SetClientHeight(value int32) {
    Control_SetClientHeight(c.instance, value)
}

func (c *TControl) ClientRect() TRect {
    return Control_GetClientRect(c.instance)
}

func (c *TControl) ClientWidth() int32 {
    return Control_GetClientWidth(c.instance)
}

func (c *TControl) SetClientWidth(value int32) {
    Control_SetClientWidth(c.instance, value)
}

func (c *TControl) ExplicitLeft() int32 {
    return Control_GetExplicitLeft(c.instance)
}

func (c *TControl) ExplicitTop() int32 {
    return Control_GetExplicitTop(c.instance)
}

func (c *TControl) ExplicitWidth() int32 {
    return Control_GetExplicitWidth(c.instance)
}

func (c *TControl) ExplicitHeight() int32 {
    return Control_GetExplicitHeight(c.instance)
}

func (c *TControl) ShowHint() bool {
    return Control_GetShowHint(c.instance)
}

func (c *TControl) SetShowHint(value bool) {
    Control_SetShowHint(c.instance, value)
}

func (c *TControl) Visible() bool {
    return Control_GetVisible(c.instance)
}

func (c *TControl) SetVisible(value bool) {
    Control_SetVisible(c.instance, value)
}

func (c *TControl) Parent() *TControl {
    return ControlFromInst(Control_GetParent(c.instance))
}

func (c *TControl) SetParent(value IControl) {
    Control_SetParent(c.instance, CheckPtr(value))
}

func (c *TControl) StyleElements() TStyleElements {
    return Control_GetStyleElements(c.instance)
}

func (c *TControl) SetStyleElements(value TStyleElements) {
    Control_SetStyleElements(c.instance, value)
}

func (c *TControl) AlignWithMargins() bool {
    return Control_GetAlignWithMargins(c.instance)
}

func (c *TControl) SetAlignWithMargins(value bool) {
    Control_SetAlignWithMargins(c.instance, value)
}

func (c *TControl) Left() int32 {
    return Control_GetLeft(c.instance)
}

func (c *TControl) SetLeft(value int32) {
    Control_SetLeft(c.instance, value)
}

func (c *TControl) Top() int32 {
    return Control_GetTop(c.instance)
}

func (c *TControl) SetTop(value int32) {
    Control_SetTop(c.instance, value)
}

func (c *TControl) Width() int32 {
    return Control_GetWidth(c.instance)
}

func (c *TControl) SetWidth(value int32) {
    Control_SetWidth(c.instance, value)
}

func (c *TControl) Height() int32 {
    return Control_GetHeight(c.instance)
}

func (c *TControl) SetHeight(value int32) {
    Control_SetHeight(c.instance, value)
}

func (c *TControl) Cursor() TCursor {
    return Control_GetCursor(c.instance)
}

func (c *TControl) SetCursor(value TCursor) {
    Control_SetCursor(c.instance, value)
}

func (c *TControl) Hint() string {
    return Control_GetHint(c.instance)
}

func (c *TControl) SetHint(value string) {
    Control_SetHint(c.instance, value)
}

func (c *TControl) Margins() *TMargins {
    return MarginsFromInst(Control_GetMargins(c.instance))
}

func (c *TControl) SetMargins(value *TMargins) {
    Control_SetMargins(c.instance, CheckPtr(value))
}

func (c *TControl) CustomHint() *TCustomHint {
    return CustomHintFromInst(Control_GetCustomHint(c.instance))
}

func (c *TControl) SetCustomHint(value IComponent) {
    Control_SetCustomHint(c.instance, CheckPtr(value))
}

func (c *TControl) ComponentCount() int32 {
    return Control_GetComponentCount(c.instance)
}

func (c *TControl) ComponentIndex() int32 {
    return Control_GetComponentIndex(c.instance)
}

func (c *TControl) SetComponentIndex(value int32) {
    Control_SetComponentIndex(c.instance, value)
}

func (c *TControl) Owner() *TComponent {
    return ComponentFromInst(Control_GetOwner(c.instance))
}

func (c *TControl) Name() string {
    return Control_GetName(c.instance)
}

func (c *TControl) SetName(value string) {
    Control_SetName(c.instance, value)
}

func (c *TControl) Tag() int {
    return Control_GetTag(c.instance)
}

func (c *TControl) SetTag(value int) {
    Control_SetTag(c.instance, value)
}

func (c *TControl) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Control_GetComponents(c.instance, AIndex))
}

