
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

type TLinkLabel struct {
    IControl
    instance uintptr
}

func NewLinkLabel(owner IComponent) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = LinkLabel_Create(CheckPtr(owner))
    return l
}

func LinkLabelFromInst(inst uintptr) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = inst
    return l
}

func LinkLabelFromObj(obj IObject) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TLinkLabel) Free() {
    if l.instance != 0 {
        LinkLabel_Free(l.instance)
        l.instance = 0
    }
}

func (l *TLinkLabel) Instance() uintptr {
    return l.instance
}

func (l *TLinkLabel) IsValid() bool {
    return l.instance != 0
}

func (l *TLinkLabel) CanFocus() bool {
    return LinkLabel_CanFocus(l.instance)
}

func (l *TLinkLabel) FlipChildren(AllLevels bool) {
    LinkLabel_FlipChildren(l.instance, AllLevels)
}

func (l *TLinkLabel) Focused() bool {
    return LinkLabel_Focused(l.instance)
}

func (l *TLinkLabel) HandleAllocated() bool {
    return LinkLabel_HandleAllocated(l.instance)
}

func (l *TLinkLabel) Invalidate() {
    LinkLabel_Invalidate(l.instance)
}

func (l *TLinkLabel) Realign() {
    LinkLabel_Realign(l.instance)
}

func (l *TLinkLabel) Repaint() {
    LinkLabel_Repaint(l.instance)
}

func (l *TLinkLabel) ScaleBy(M int32, D int32) {
    LinkLabel_ScaleBy(l.instance, M , D)
}

func (l *TLinkLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    LinkLabel_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

func (l *TLinkLabel) SetFocus() {
    LinkLabel_SetFocus(l.instance)
}

func (l *TLinkLabel) Update() {
    LinkLabel_Update(l.instance)
}

func (l *TLinkLabel) BringToFront() {
    LinkLabel_BringToFront(l.instance)
}

func (l *TLinkLabel) HasParent() bool {
    return LinkLabel_HasParent(l.instance)
}

func (l *TLinkLabel) Hide() {
    LinkLabel_Hide(l.instance)
}

func (l *TLinkLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return LinkLabel_Perform(l.instance, Msg , WParam , LParam)
}

func (l *TLinkLabel) Refresh() {
    LinkLabel_Refresh(l.instance)
}

func (l *TLinkLabel) SendToBack() {
    LinkLabel_SendToBack(l.instance)
}

func (l *TLinkLabel) Show() {
    LinkLabel_Show(l.instance)
}

func (l *TLinkLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return LinkLabel_GetTextBuf(l.instance, Buffer , BufSize)
}

func (l *TLinkLabel) GetTextLen() int32 {
    return LinkLabel_GetTextLen(l.instance)
}

func (l *TLinkLabel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(LinkLabel_FindComponent(l.instance, AName))
}

func (l *TLinkLabel) GetNamePath() string {
    return LinkLabel_GetNamePath(l.instance)
}

func (l *TLinkLabel) Assign(Source IObject) {
    LinkLabel_Assign(l.instance, CheckPtr(Source))
}

func (l *TLinkLabel) ClassName() string {
    return LinkLabel_ClassName(l.instance)
}

func (l *TLinkLabel) Equals(Obj IObject) bool {
    return LinkLabel_Equals(l.instance, CheckPtr(Obj))
}

func (l *TLinkLabel) GetHashCode() int32 {
    return LinkLabel_GetHashCode(l.instance)
}

func (l *TLinkLabel) ToString() string {
    return LinkLabel_ToString(l.instance)
}

func (l *TLinkLabel) Align() TAlign {
    return LinkLabel_GetAlign(l.instance)
}

func (l *TLinkLabel) SetAlign(value TAlign) {
    LinkLabel_SetAlign(l.instance, value)
}

func (l *TLinkLabel) Alignment() TLinkAlignment {
    return LinkLabel_GetAlignment(l.instance)
}

func (l *TLinkLabel) SetAlignment(value TLinkAlignment) {
    LinkLabel_SetAlignment(l.instance, value)
}

func (l *TLinkLabel) Anchors() TAnchors {
    return LinkLabel_GetAnchors(l.instance)
}

func (l *TLinkLabel) SetAnchors(value TAnchors) {
    LinkLabel_SetAnchors(l.instance, value)
}

func (l *TLinkLabel) AutoSize() bool {
    return LinkLabel_GetAutoSize(l.instance)
}

func (l *TLinkLabel) SetAutoSize(value bool) {
    LinkLabel_SetAutoSize(l.instance, value)
}

func (l *TLinkLabel) BevelEdges() TBevelEdges {
    return LinkLabel_GetBevelEdges(l.instance)
}

func (l *TLinkLabel) SetBevelEdges(value TBevelEdges) {
    LinkLabel_SetBevelEdges(l.instance, value)
}

func (l *TLinkLabel) BevelInner() TBevelCut {
    return LinkLabel_GetBevelInner(l.instance)
}

func (l *TLinkLabel) SetBevelInner(value TBevelCut) {
    LinkLabel_SetBevelInner(l.instance, value)
}

func (l *TLinkLabel) BevelKind() TBevelKind {
    return LinkLabel_GetBevelKind(l.instance)
}

func (l *TLinkLabel) SetBevelKind(value TBevelKind) {
    LinkLabel_SetBevelKind(l.instance, value)
}

func (l *TLinkLabel) BevelOuter() TBevelCut {
    return LinkLabel_GetBevelOuter(l.instance)
}

func (l *TLinkLabel) SetBevelOuter(value TBevelCut) {
    LinkLabel_SetBevelOuter(l.instance, value)
}

func (l *TLinkLabel) Caption() string {
    return LinkLabel_GetCaption(l.instance)
}

func (l *TLinkLabel) SetCaption(value string) {
    LinkLabel_SetCaption(l.instance, value)
}

func (l *TLinkLabel) Color() TColor {
    return LinkLabel_GetColor(l.instance)
}

func (l *TLinkLabel) SetColor(value TColor) {
    LinkLabel_SetColor(l.instance, value)
}

func (l *TLinkLabel) Enabled() bool {
    return LinkLabel_GetEnabled(l.instance)
}

func (l *TLinkLabel) SetEnabled(value bool) {
    LinkLabel_SetEnabled(l.instance, value)
}

func (l *TLinkLabel) Font() *TFont {
    return FontFromInst(LinkLabel_GetFont(l.instance))
}

func (l *TLinkLabel) SetFont(value *TFont) {
    LinkLabel_SetFont(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) ParentColor() bool {
    return LinkLabel_GetParentColor(l.instance)
}

func (l *TLinkLabel) SetParentColor(value bool) {
    LinkLabel_SetParentColor(l.instance, value)
}

func (l *TLinkLabel) ParentFont() bool {
    return LinkLabel_GetParentFont(l.instance)
}

func (l *TLinkLabel) SetParentFont(value bool) {
    LinkLabel_SetParentFont(l.instance, value)
}

func (l *TLinkLabel) ParentShowHint() bool {
    return LinkLabel_GetParentShowHint(l.instance)
}

func (l *TLinkLabel) SetParentShowHint(value bool) {
    LinkLabel_SetParentShowHint(l.instance, value)
}

func (l *TLinkLabel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(LinkLabel_GetPopupMenu(l.instance))
}

func (l *TLinkLabel) SetPopupMenu(value IComponent) {
    LinkLabel_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) ShowHint() bool {
    return LinkLabel_GetShowHint(l.instance)
}

func (l *TLinkLabel) SetShowHint(value bool) {
    LinkLabel_SetShowHint(l.instance, value)
}

func (l *TLinkLabel) TabOrder() uint16 {
    return LinkLabel_GetTabOrder(l.instance)
}

func (l *TLinkLabel) SetTabOrder(value uint16) {
    LinkLabel_SetTabOrder(l.instance, value)
}

func (l *TLinkLabel) TabStop() bool {
    return LinkLabel_GetTabStop(l.instance)
}

func (l *TLinkLabel) SetTabStop(value bool) {
    LinkLabel_SetTabStop(l.instance, value)
}

func (l *TLinkLabel) UseVisualStyle() bool {
    return LinkLabel_GetUseVisualStyle(l.instance)
}

func (l *TLinkLabel) SetUseVisualStyle(value bool) {
    LinkLabel_SetUseVisualStyle(l.instance, value)
}

func (l *TLinkLabel) Visible() bool {
    return LinkLabel_GetVisible(l.instance)
}

func (l *TLinkLabel) SetVisible(value bool) {
    LinkLabel_SetVisible(l.instance, value)
}

func (l *TLinkLabel) SetOnClick(fn TNotifyEvent) {
    LinkLabel_SetOnClick(l.instance, fn)
}

func (l *TLinkLabel) SetOnDblClick(fn TNotifyEvent) {
    LinkLabel_SetOnDblClick(l.instance, fn)
}

func (l *TLinkLabel) SetOnMouseDown(fn TMouseEvent) {
    LinkLabel_SetOnMouseDown(l.instance, fn)
}

func (l *TLinkLabel) SetOnMouseEnter(fn TNotifyEvent) {
    LinkLabel_SetOnMouseEnter(l.instance, fn)
}

func (l *TLinkLabel) SetOnMouseLeave(fn TNotifyEvent) {
    LinkLabel_SetOnMouseLeave(l.instance, fn)
}

func (l *TLinkLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    LinkLabel_SetOnMouseMove(l.instance, fn)
}

func (l *TLinkLabel) SetOnMouseUp(fn TMouseEvent) {
    LinkLabel_SetOnMouseUp(l.instance, fn)
}

func (l *TLinkLabel) SetOnLinkClick(fn TSysLinkEvent) {
    LinkLabel_SetOnLinkClick(l.instance, fn)
}

func (l *TLinkLabel) DoubleBuffered() bool {
    return LinkLabel_GetDoubleBuffered(l.instance)
}

func (l *TLinkLabel) SetDoubleBuffered(value bool) {
    LinkLabel_SetDoubleBuffered(l.instance, value)
}

func (l *TLinkLabel) Brush() *TBrush {
    return BrushFromInst(LinkLabel_GetBrush(l.instance))
}

func (l *TLinkLabel) ControlCount() int32 {
    return LinkLabel_GetControlCount(l.instance)
}

func (l *TLinkLabel) Handle() HWND {
    return LinkLabel_GetHandle(l.instance)
}

func (l *TLinkLabel) ParentDoubleBuffered() bool {
    return LinkLabel_GetParentDoubleBuffered(l.instance)
}

func (l *TLinkLabel) SetParentDoubleBuffered(value bool) {
    LinkLabel_SetParentDoubleBuffered(l.instance, value)
}

func (l *TLinkLabel) ParentWindow() HWND {
    return LinkLabel_GetParentWindow(l.instance)
}

func (l *TLinkLabel) SetParentWindow(value HWND) {
    LinkLabel_SetParentWindow(l.instance, value)
}

func (l *TLinkLabel) Action() *TAction {
    return ActionFromInst(LinkLabel_GetAction(l.instance))
}

func (l *TLinkLabel) SetAction(value IComponent) {
    LinkLabel_SetAction(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) BiDiMode() TBiDiMode {
    return LinkLabel_GetBiDiMode(l.instance)
}

func (l *TLinkLabel) SetBiDiMode(value TBiDiMode) {
    LinkLabel_SetBiDiMode(l.instance, value)
}

func (l *TLinkLabel) BoundsRect() TRect {
    return LinkLabel_GetBoundsRect(l.instance)
}

func (l *TLinkLabel) SetBoundsRect(value TRect) {
    LinkLabel_SetBoundsRect(l.instance, value)
}

func (l *TLinkLabel) ClientHeight() int32 {
    return LinkLabel_GetClientHeight(l.instance)
}

func (l *TLinkLabel) SetClientHeight(value int32) {
    LinkLabel_SetClientHeight(l.instance, value)
}

func (l *TLinkLabel) ClientRect() TRect {
    return LinkLabel_GetClientRect(l.instance)
}

func (l *TLinkLabel) ClientWidth() int32 {
    return LinkLabel_GetClientWidth(l.instance)
}

func (l *TLinkLabel) SetClientWidth(value int32) {
    LinkLabel_SetClientWidth(l.instance, value)
}

func (l *TLinkLabel) ExplicitLeft() int32 {
    return LinkLabel_GetExplicitLeft(l.instance)
}

func (l *TLinkLabel) ExplicitTop() int32 {
    return LinkLabel_GetExplicitTop(l.instance)
}

func (l *TLinkLabel) ExplicitWidth() int32 {
    return LinkLabel_GetExplicitWidth(l.instance)
}

func (l *TLinkLabel) ExplicitHeight() int32 {
    return LinkLabel_GetExplicitHeight(l.instance)
}

func (l *TLinkLabel) Parent() *TControl {
    return ControlFromInst(LinkLabel_GetParent(l.instance))
}

func (l *TLinkLabel) SetParent(value IControl) {
    LinkLabel_SetParent(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) StyleElements() TStyleElements {
    return LinkLabel_GetStyleElements(l.instance)
}

func (l *TLinkLabel) SetStyleElements(value TStyleElements) {
    LinkLabel_SetStyleElements(l.instance, value)
}

func (l *TLinkLabel) AlignWithMargins() bool {
    return LinkLabel_GetAlignWithMargins(l.instance)
}

func (l *TLinkLabel) SetAlignWithMargins(value bool) {
    LinkLabel_SetAlignWithMargins(l.instance, value)
}

func (l *TLinkLabel) Left() int32 {
    return LinkLabel_GetLeft(l.instance)
}

func (l *TLinkLabel) SetLeft(value int32) {
    LinkLabel_SetLeft(l.instance, value)
}

func (l *TLinkLabel) Top() int32 {
    return LinkLabel_GetTop(l.instance)
}

func (l *TLinkLabel) SetTop(value int32) {
    LinkLabel_SetTop(l.instance, value)
}

func (l *TLinkLabel) Width() int32 {
    return LinkLabel_GetWidth(l.instance)
}

func (l *TLinkLabel) SetWidth(value int32) {
    LinkLabel_SetWidth(l.instance, value)
}

func (l *TLinkLabel) Height() int32 {
    return LinkLabel_GetHeight(l.instance)
}

func (l *TLinkLabel) SetHeight(value int32) {
    LinkLabel_SetHeight(l.instance, value)
}

func (l *TLinkLabel) Cursor() TCursor {
    return LinkLabel_GetCursor(l.instance)
}

func (l *TLinkLabel) SetCursor(value TCursor) {
    LinkLabel_SetCursor(l.instance, value)
}

func (l *TLinkLabel) Hint() string {
    return LinkLabel_GetHint(l.instance)
}

func (l *TLinkLabel) SetHint(value string) {
    LinkLabel_SetHint(l.instance, value)
}

func (l *TLinkLabel) Margins() *TMargins {
    return MarginsFromInst(LinkLabel_GetMargins(l.instance))
}

func (l *TLinkLabel) SetMargins(value *TMargins) {
    LinkLabel_SetMargins(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) CustomHint() *TCustomHint {
    return CustomHintFromInst(LinkLabel_GetCustomHint(l.instance))
}

func (l *TLinkLabel) SetCustomHint(value IComponent) {
    LinkLabel_SetCustomHint(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) ComponentCount() int32 {
    return LinkLabel_GetComponentCount(l.instance)
}

func (l *TLinkLabel) ComponentIndex() int32 {
    return LinkLabel_GetComponentIndex(l.instance)
}

func (l *TLinkLabel) SetComponentIndex(value int32) {
    LinkLabel_SetComponentIndex(l.instance, value)
}

func (l *TLinkLabel) Owner() *TComponent {
    return ComponentFromInst(LinkLabel_GetOwner(l.instance))
}

func (l *TLinkLabel) Name() string {
    return LinkLabel_GetName(l.instance)
}

func (l *TLinkLabel) SetName(value string) {
    LinkLabel_SetName(l.instance, value)
}

func (l *TLinkLabel) Tag() int {
    return LinkLabel_GetTag(l.instance)
}

func (l *TLinkLabel) SetTag(value int) {
    LinkLabel_SetTag(l.instance, value)
}

func (l *TLinkLabel) Controls(Index int32) *TControl {
    return ControlFromInst(LinkLabel_GetControls(l.instance, Index))
}

func (l *TLinkLabel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(LinkLabel_GetComponents(l.instance, AIndex))
}

