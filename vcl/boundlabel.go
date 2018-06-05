
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
)

type TBoundLabel struct {
    IControl
    instance uintptr
}

func NewBoundLabel(owner IComponent) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = BoundLabel_Create(CheckPtr(owner))
    return b
}

func BoundLabelFromInst(inst uintptr) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = inst
    return b
}

func BoundLabelFromObj(obj IObject) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = CheckPtr(obj)
    return b
}

func (b *TBoundLabel) Free() {
    if b.instance != 0 {
        BoundLabel_Free(b.instance)
        b.instance = 0
    }
}

func (b *TBoundLabel) Instance() uintptr {
    return b.instance
}

func (b *TBoundLabel) IsValid() bool {
    return b.instance != 0
}

func TBoundLabelClass() TClass {
    return BoundLabel_StaticClassType()
}

func (b *TBoundLabel) BringToFront() {
    BoundLabel_BringToFront(b.instance)
}

func (b *TBoundLabel) ClientToScreen(Point TPoint) TPoint {
    return BoundLabel_ClientToScreen(b.instance, Point)
}

func (b *TBoundLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return BoundLabel_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

func (b *TBoundLabel) Dragging() bool {
    return BoundLabel_Dragging(b.instance)
}

func (b *TBoundLabel) HasParent() bool {
    return BoundLabel_HasParent(b.instance)
}

func (b *TBoundLabel) Hide() {
    BoundLabel_Hide(b.instance)
}

func (b *TBoundLabel) Invalidate() {
    BoundLabel_Invalidate(b.instance)
}

func (b *TBoundLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return BoundLabel_Perform(b.instance, Msg , WParam , LParam)
}

func (b *TBoundLabel) Refresh() {
    BoundLabel_Refresh(b.instance)
}

func (b *TBoundLabel) Repaint() {
    BoundLabel_Repaint(b.instance)
}

func (b *TBoundLabel) ScreenToClient(Point TPoint) TPoint {
    return BoundLabel_ScreenToClient(b.instance, Point)
}

func (b *TBoundLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return BoundLabel_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

func (b *TBoundLabel) SendToBack() {
    BoundLabel_SendToBack(b.instance)
}

func (b *TBoundLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    BoundLabel_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

func (b *TBoundLabel) Show() {
    BoundLabel_Show(b.instance)
}

func (b *TBoundLabel) Update() {
    BoundLabel_Update(b.instance)
}

func (b *TBoundLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return BoundLabel_GetTextBuf(b.instance, Buffer , BufSize)
}

func (b *TBoundLabel) GetTextLen() int32 {
    return BoundLabel_GetTextLen(b.instance)
}

func (b *TBoundLabel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(BoundLabel_FindComponent(b.instance, AName))
}

func (b *TBoundLabel) GetNamePath() string {
    return BoundLabel_GetNamePath(b.instance)
}

func (b *TBoundLabel) Assign(Source IObject) {
    BoundLabel_Assign(b.instance, CheckPtr(Source))
}

func (b *TBoundLabel) DisposeOf() {
    BoundLabel_DisposeOf(b.instance)
}

func (b *TBoundLabel) ClassType() TClass {
    return BoundLabel_ClassType(b.instance)
}

func (b *TBoundLabel) ClassName() string {
    return BoundLabel_ClassName(b.instance)
}

func (b *TBoundLabel) InstanceSize() int32 {
    return BoundLabel_InstanceSize(b.instance)
}

func (b *TBoundLabel) InheritsFrom(AClass TClass) bool {
    return BoundLabel_InheritsFrom(b.instance, AClass)
}

func (b *TBoundLabel) Equals(Obj IObject) bool {
    return BoundLabel_Equals(b.instance, CheckPtr(Obj))
}

func (b *TBoundLabel) GetHashCode() int32 {
    return BoundLabel_GetHashCode(b.instance)
}

func (b *TBoundLabel) ToString() string {
    return BoundLabel_ToString(b.instance)
}

func (b *TBoundLabel) BiDiMode() TBiDiMode {
    return BoundLabel_GetBiDiMode(b.instance)
}

func (b *TBoundLabel) SetBiDiMode(value TBiDiMode) {
    BoundLabel_SetBiDiMode(b.instance, value)
}

func (b *TBoundLabel) Caption() string {
    return BoundLabel_GetCaption(b.instance)
}

func (b *TBoundLabel) SetCaption(value string) {
    BoundLabel_SetCaption(b.instance, value)
}

func (b *TBoundLabel) Color() TColor {
    return BoundLabel_GetColor(b.instance)
}

func (b *TBoundLabel) SetColor(value TColor) {
    BoundLabel_SetColor(b.instance, value)
}

func (b *TBoundLabel) DragCursor() TCursor {
    return BoundLabel_GetDragCursor(b.instance)
}

func (b *TBoundLabel) SetDragCursor(value TCursor) {
    BoundLabel_SetDragCursor(b.instance, value)
}

func (b *TBoundLabel) DragKind() TDragKind {
    return BoundLabel_GetDragKind(b.instance)
}

func (b *TBoundLabel) SetDragKind(value TDragKind) {
    BoundLabel_SetDragKind(b.instance, value)
}

func (b *TBoundLabel) DragMode() TDragMode {
    return BoundLabel_GetDragMode(b.instance)
}

func (b *TBoundLabel) SetDragMode(value TDragMode) {
    BoundLabel_SetDragMode(b.instance, value)
}

func (b *TBoundLabel) Font() *TFont {
    return FontFromInst(BoundLabel_GetFont(b.instance))
}

func (b *TBoundLabel) SetFont(value *TFont) {
    BoundLabel_SetFont(b.instance, CheckPtr(value))
}

func (b *TBoundLabel) Height() int32 {
    return BoundLabel_GetHeight(b.instance)
}

func (b *TBoundLabel) SetHeight(value int32) {
    BoundLabel_SetHeight(b.instance, value)
}

func (b *TBoundLabel) Left() int32 {
    return BoundLabel_GetLeft(b.instance)
}

func (b *TBoundLabel) ParentColor() bool {
    return BoundLabel_GetParentColor(b.instance)
}

func (b *TBoundLabel) SetParentColor(value bool) {
    BoundLabel_SetParentColor(b.instance, value)
}

func (b *TBoundLabel) ParentFont() bool {
    return BoundLabel_GetParentFont(b.instance)
}

func (b *TBoundLabel) SetParentFont(value bool) {
    BoundLabel_SetParentFont(b.instance, value)
}

func (b *TBoundLabel) ParentShowHint() bool {
    return BoundLabel_GetParentShowHint(b.instance)
}

func (b *TBoundLabel) SetParentShowHint(value bool) {
    BoundLabel_SetParentShowHint(b.instance, value)
}

func (b *TBoundLabel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(BoundLabel_GetPopupMenu(b.instance))
}

func (b *TBoundLabel) SetPopupMenu(value IComponent) {
    BoundLabel_SetPopupMenu(b.instance, CheckPtr(value))
}

func (b *TBoundLabel) ShowAccelChar() bool {
    return BoundLabel_GetShowAccelChar(b.instance)
}

func (b *TBoundLabel) SetShowAccelChar(value bool) {
    BoundLabel_SetShowAccelChar(b.instance, value)
}

func (b *TBoundLabel) ShowHint() bool {
    return BoundLabel_GetShowHint(b.instance)
}

func (b *TBoundLabel) SetShowHint(value bool) {
    BoundLabel_SetShowHint(b.instance, value)
}

func (b *TBoundLabel) Top() int32 {
    return BoundLabel_GetTop(b.instance)
}

func (b *TBoundLabel) Transparent() bool {
    return BoundLabel_GetTransparent(b.instance)
}

func (b *TBoundLabel) SetTransparent(value bool) {
    BoundLabel_SetTransparent(b.instance, value)
}

func (b *TBoundLabel) Layout() TTextLayout {
    return BoundLabel_GetLayout(b.instance)
}

func (b *TBoundLabel) SetLayout(value TTextLayout) {
    BoundLabel_SetLayout(b.instance, value)
}

func (b *TBoundLabel) WordWrap() bool {
    return BoundLabel_GetWordWrap(b.instance)
}

func (b *TBoundLabel) SetWordWrap(value bool) {
    BoundLabel_SetWordWrap(b.instance, value)
}

func (b *TBoundLabel) Width() int32 {
    return BoundLabel_GetWidth(b.instance)
}

func (b *TBoundLabel) SetWidth(value int32) {
    BoundLabel_SetWidth(b.instance, value)
}

func (b *TBoundLabel) SetOnClick(fn TNotifyEvent) {
    BoundLabel_SetOnClick(b.instance, fn)
}

func (b *TBoundLabel) SetOnContextPopup(fn TContextPopupEvent) {
    BoundLabel_SetOnContextPopup(b.instance, fn)
}

func (b *TBoundLabel) SetOnDblClick(fn TNotifyEvent) {
    BoundLabel_SetOnDblClick(b.instance, fn)
}

func (b *TBoundLabel) SetOnDragDrop(fn TDragDropEvent) {
    BoundLabel_SetOnDragDrop(b.instance, fn)
}

func (b *TBoundLabel) SetOnDragOver(fn TDragOverEvent) {
    BoundLabel_SetOnDragOver(b.instance, fn)
}

func (b *TBoundLabel) SetOnEndDock(fn TEndDragEvent) {
    BoundLabel_SetOnEndDock(b.instance, fn)
}

func (b *TBoundLabel) SetOnEndDrag(fn TEndDragEvent) {
    BoundLabel_SetOnEndDrag(b.instance, fn)
}

func (b *TBoundLabel) SetOnMouseDown(fn TMouseEvent) {
    BoundLabel_SetOnMouseDown(b.instance, fn)
}

func (b *TBoundLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    BoundLabel_SetOnMouseMove(b.instance, fn)
}

func (b *TBoundLabel) SetOnMouseUp(fn TMouseEvent) {
    BoundLabel_SetOnMouseUp(b.instance, fn)
}

func (b *TBoundLabel) SetOnStartDock(fn TStartDockEvent) {
    BoundLabel_SetOnStartDock(b.instance, fn)
}

func (b *TBoundLabel) Canvas() *TCanvas {
    return CanvasFromInst(BoundLabel_GetCanvas(b.instance))
}

func (b *TBoundLabel) GlowSize() int32 {
    return BoundLabel_GetGlowSize(b.instance)
}

func (b *TBoundLabel) SetGlowSize(value int32) {
    BoundLabel_SetGlowSize(b.instance, value)
}

func (b *TBoundLabel) Enabled() bool {
    return BoundLabel_GetEnabled(b.instance)
}

func (b *TBoundLabel) SetEnabled(value bool) {
    BoundLabel_SetEnabled(b.instance, value)
}

func (b *TBoundLabel) Action() *TAction {
    return ActionFromInst(BoundLabel_GetAction(b.instance))
}

func (b *TBoundLabel) SetAction(value IComponent) {
    BoundLabel_SetAction(b.instance, CheckPtr(value))
}

func (b *TBoundLabel) Align() TAlign {
    return BoundLabel_GetAlign(b.instance)
}

func (b *TBoundLabel) SetAlign(value TAlign) {
    BoundLabel_SetAlign(b.instance, value)
}

func (b *TBoundLabel) Anchors() TAnchors {
    return BoundLabel_GetAnchors(b.instance)
}

func (b *TBoundLabel) SetAnchors(value TAnchors) {
    BoundLabel_SetAnchors(b.instance, value)
}

func (b *TBoundLabel) BoundsRect() TRect {
    return BoundLabel_GetBoundsRect(b.instance)
}

func (b *TBoundLabel) SetBoundsRect(value TRect) {
    BoundLabel_SetBoundsRect(b.instance, value)
}

func (b *TBoundLabel) ClientHeight() int32 {
    return BoundLabel_GetClientHeight(b.instance)
}

func (b *TBoundLabel) SetClientHeight(value int32) {
    BoundLabel_SetClientHeight(b.instance, value)
}

func (b *TBoundLabel) ClientRect() TRect {
    return BoundLabel_GetClientRect(b.instance)
}

func (b *TBoundLabel) ClientWidth() int32 {
    return BoundLabel_GetClientWidth(b.instance)
}

func (b *TBoundLabel) SetClientWidth(value int32) {
    BoundLabel_SetClientWidth(b.instance, value)
}

func (b *TBoundLabel) ExplicitLeft() int32 {
    return BoundLabel_GetExplicitLeft(b.instance)
}

func (b *TBoundLabel) ExplicitTop() int32 {
    return BoundLabel_GetExplicitTop(b.instance)
}

func (b *TBoundLabel) ExplicitWidth() int32 {
    return BoundLabel_GetExplicitWidth(b.instance)
}

func (b *TBoundLabel) ExplicitHeight() int32 {
    return BoundLabel_GetExplicitHeight(b.instance)
}

func (b *TBoundLabel) Floating() bool {
    return BoundLabel_GetFloating(b.instance)
}

func (b *TBoundLabel) Visible() bool {
    return BoundLabel_GetVisible(b.instance)
}

func (b *TBoundLabel) SetVisible(value bool) {
    BoundLabel_SetVisible(b.instance, value)
}

func (b *TBoundLabel) Parent() *TWinControl {
    return WinControlFromInst(BoundLabel_GetParent(b.instance))
}

func (b *TBoundLabel) SetParent(value IWinControl) {
    BoundLabel_SetParent(b.instance, CheckPtr(value))
}

func (b *TBoundLabel) StyleElements() TStyleElements {
    return BoundLabel_GetStyleElements(b.instance)
}

func (b *TBoundLabel) SetStyleElements(value TStyleElements) {
    BoundLabel_SetStyleElements(b.instance, value)
}

func (b *TBoundLabel) AlignWithMargins() bool {
    return BoundLabel_GetAlignWithMargins(b.instance)
}

func (b *TBoundLabel) SetAlignWithMargins(value bool) {
    BoundLabel_SetAlignWithMargins(b.instance, value)
}

func (b *TBoundLabel) Cursor() TCursor {
    return BoundLabel_GetCursor(b.instance)
}

func (b *TBoundLabel) SetCursor(value TCursor) {
    BoundLabel_SetCursor(b.instance, value)
}

func (b *TBoundLabel) Hint() string {
    return BoundLabel_GetHint(b.instance)
}

func (b *TBoundLabel) SetHint(value string) {
    BoundLabel_SetHint(b.instance, value)
}

func (b *TBoundLabel) Margins() *TMargins {
    return MarginsFromInst(BoundLabel_GetMargins(b.instance))
}

func (b *TBoundLabel) SetMargins(value *TMargins) {
    BoundLabel_SetMargins(b.instance, CheckPtr(value))
}

func (b *TBoundLabel) CustomHint() *TCustomHint {
    return CustomHintFromInst(BoundLabel_GetCustomHint(b.instance))
}

func (b *TBoundLabel) SetCustomHint(value IComponent) {
    BoundLabel_SetCustomHint(b.instance, CheckPtr(value))
}

func (b *TBoundLabel) ComponentCount() int32 {
    return BoundLabel_GetComponentCount(b.instance)
}

func (b *TBoundLabel) ComponentIndex() int32 {
    return BoundLabel_GetComponentIndex(b.instance)
}

func (b *TBoundLabel) SetComponentIndex(value int32) {
    BoundLabel_SetComponentIndex(b.instance, value)
}

func (b *TBoundLabel) Owner() *TComponent {
    return ComponentFromInst(BoundLabel_GetOwner(b.instance))
}

func (b *TBoundLabel) Name() string {
    return BoundLabel_GetName(b.instance)
}

func (b *TBoundLabel) SetName(value string) {
    BoundLabel_SetName(b.instance, value)
}

func (b *TBoundLabel) Tag() int {
    return BoundLabel_GetTag(b.instance)
}

func (b *TBoundLabel) SetTag(value int) {
    BoundLabel_SetTag(b.instance, value)
}

func (b *TBoundLabel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(BoundLabel_GetComponents(b.instance, AIndex))
}

