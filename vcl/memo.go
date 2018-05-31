
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

type TMemo struct {
    IControl
    instance uintptr
}

func NewMemo(owner IComponent) *TMemo {
    m := new(TMemo)
    m.instance = Memo_Create(CheckPtr(owner))
    return m
}

func MemoFromInst(inst uintptr) *TMemo {
    m := new(TMemo)
    m.instance = inst
    return m
}

func MemoFromObj(obj IObject) *TMemo {
    m := new(TMemo)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMemo) Free() {
    if m.instance != 0 {
        Memo_Free(m.instance)
        m.instance = 0
    }
}

func (m *TMemo) Instance() uintptr {
    return m.instance
}

func (m *TMemo) IsValid() bool {
    return m.instance != 0
}

func (m *TMemo) Clear() {
    Memo_Clear(m.instance)
}

func (m *TMemo) ClearSelection() {
    Memo_ClearSelection(m.instance)
}

func (m *TMemo) CopyToClipboard() {
    Memo_CopyToClipboard(m.instance)
}

func (m *TMemo) CutToClipboard() {
    Memo_CutToClipboard(m.instance)
}

func (m *TMemo) PasteFromClipboard() {
    Memo_PasteFromClipboard(m.instance)
}

func (m *TMemo) Undo() {
    Memo_Undo(m.instance)
}

func (m *TMemo) ClearUndo() {
    Memo_ClearUndo(m.instance)
}

func (m *TMemo) SelectAll() {
    Memo_SelectAll(m.instance)
}

func (m *TMemo) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return Memo_GetSelTextBuf(m.instance, Buffer , BufSize)
}

func (m *TMemo) CanFocus() bool {
    return Memo_CanFocus(m.instance)
}

func (m *TMemo) FlipChildren(AllLevels bool) {
    Memo_FlipChildren(m.instance, AllLevels)
}

func (m *TMemo) Focused() bool {
    return Memo_Focused(m.instance)
}

func (m *TMemo) HandleAllocated() bool {
    return Memo_HandleAllocated(m.instance)
}

func (m *TMemo) Invalidate() {
    Memo_Invalidate(m.instance)
}

func (m *TMemo) Realign() {
    Memo_Realign(m.instance)
}

func (m *TMemo) Repaint() {
    Memo_Repaint(m.instance)
}

func (m *TMemo) ScaleBy(M int32, D int32) {
    Memo_ScaleBy(m.instance, M , D)
}

func (m *TMemo) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Memo_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

func (m *TMemo) SetFocus() {
    Memo_SetFocus(m.instance)
}

func (m *TMemo) Update() {
    Memo_Update(m.instance)
}

func (m *TMemo) BringToFront() {
    Memo_BringToFront(m.instance)
}

func (m *TMemo) Dragging() bool {
    return Memo_Dragging(m.instance)
}

func (m *TMemo) HasParent() bool {
    return Memo_HasParent(m.instance)
}

func (m *TMemo) Hide() {
    Memo_Hide(m.instance)
}

func (m *TMemo) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Memo_Perform(m.instance, Msg , WParam , LParam)
}

func (m *TMemo) Refresh() {
    Memo_Refresh(m.instance)
}

func (m *TMemo) SendToBack() {
    Memo_SendToBack(m.instance)
}

func (m *TMemo) Show() {
    Memo_Show(m.instance)
}

func (m *TMemo) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Memo_GetTextBuf(m.instance, Buffer , BufSize)
}

func (m *TMemo) GetTextLen() int32 {
    return Memo_GetTextLen(m.instance)
}

func (m *TMemo) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Memo_FindComponent(m.instance, AName))
}

func (m *TMemo) GetNamePath() string {
    return Memo_GetNamePath(m.instance)
}

func (m *TMemo) Assign(Source IObject) {
    Memo_Assign(m.instance, CheckPtr(Source))
}

func (m *TMemo) ClassName() string {
    return Memo_ClassName(m.instance)
}

func (m *TMemo) Equals(Obj IObject) bool {
    return Memo_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMemo) GetHashCode() int32 {
    return Memo_GetHashCode(m.instance)
}

func (m *TMemo) ToString() string {
    return Memo_ToString(m.instance)
}

func (m *TMemo) Align() TAlign {
    return Memo_GetAlign(m.instance)
}

func (m *TMemo) SetAlign(value TAlign) {
    Memo_SetAlign(m.instance, value)
}

func (m *TMemo) Alignment() TAlignment {
    return Memo_GetAlignment(m.instance)
}

func (m *TMemo) SetAlignment(value TAlignment) {
    Memo_SetAlignment(m.instance, value)
}

func (m *TMemo) Anchors() TAnchors {
    return Memo_GetAnchors(m.instance)
}

func (m *TMemo) SetAnchors(value TAnchors) {
    Memo_SetAnchors(m.instance, value)
}

func (m *TMemo) BevelEdges() TBevelEdges {
    return Memo_GetBevelEdges(m.instance)
}

func (m *TMemo) SetBevelEdges(value TBevelEdges) {
    Memo_SetBevelEdges(m.instance, value)
}

func (m *TMemo) BevelInner() TBevelCut {
    return Memo_GetBevelInner(m.instance)
}

func (m *TMemo) SetBevelInner(value TBevelCut) {
    Memo_SetBevelInner(m.instance, value)
}

func (m *TMemo) BevelKind() TBevelKind {
    return Memo_GetBevelKind(m.instance)
}

func (m *TMemo) SetBevelKind(value TBevelKind) {
    Memo_SetBevelKind(m.instance, value)
}

func (m *TMemo) BevelOuter() TBevelCut {
    return Memo_GetBevelOuter(m.instance)
}

func (m *TMemo) SetBevelOuter(value TBevelCut) {
    Memo_SetBevelOuter(m.instance, value)
}

func (m *TMemo) BiDiMode() TBiDiMode {
    return Memo_GetBiDiMode(m.instance)
}

func (m *TMemo) SetBiDiMode(value TBiDiMode) {
    Memo_SetBiDiMode(m.instance, value)
}

func (m *TMemo) BorderStyle() TBorderStyle {
    return Memo_GetBorderStyle(m.instance)
}

func (m *TMemo) SetBorderStyle(value TBorderStyle) {
    Memo_SetBorderStyle(m.instance, value)
}

func (m *TMemo) Color() TColor {
    return Memo_GetColor(m.instance)
}

func (m *TMemo) SetColor(value TColor) {
    Memo_SetColor(m.instance, value)
}

func (m *TMemo) DoubleBuffered() bool {
    return Memo_GetDoubleBuffered(m.instance)
}

func (m *TMemo) SetDoubleBuffered(value bool) {
    Memo_SetDoubleBuffered(m.instance, value)
}

func (m *TMemo) DragCursor() TCursor {
    return Memo_GetDragCursor(m.instance)
}

func (m *TMemo) SetDragCursor(value TCursor) {
    Memo_SetDragCursor(m.instance, value)
}

func (m *TMemo) DragKind() TDragKind {
    return Memo_GetDragKind(m.instance)
}

func (m *TMemo) SetDragKind(value TDragKind) {
    Memo_SetDragKind(m.instance, value)
}

func (m *TMemo) DragMode() TDragMode {
    return Memo_GetDragMode(m.instance)
}

func (m *TMemo) SetDragMode(value TDragMode) {
    Memo_SetDragMode(m.instance, value)
}

func (m *TMemo) Enabled() bool {
    return Memo_GetEnabled(m.instance)
}

func (m *TMemo) SetEnabled(value bool) {
    Memo_SetEnabled(m.instance, value)
}

func (m *TMemo) Font() *TFont {
    return FontFromInst(Memo_GetFont(m.instance))
}

func (m *TMemo) SetFont(value *TFont) {
    Memo_SetFont(m.instance, CheckPtr(value))
}

func (m *TMemo) HideSelection() bool {
    return Memo_GetHideSelection(m.instance)
}

func (m *TMemo) SetHideSelection(value bool) {
    Memo_SetHideSelection(m.instance, value)
}

func (m *TMemo) Lines() *TStrings {
    return StringsFromInst(Memo_GetLines(m.instance))
}

func (m *TMemo) SetLines(value IObject) {
    Memo_SetLines(m.instance, CheckPtr(value))
}

func (m *TMemo) MaxLength() int32 {
    return Memo_GetMaxLength(m.instance)
}

func (m *TMemo) SetMaxLength(value int32) {
    Memo_SetMaxLength(m.instance, value)
}

func (m *TMemo) ParentColor() bool {
    return Memo_GetParentColor(m.instance)
}

func (m *TMemo) SetParentColor(value bool) {
    Memo_SetParentColor(m.instance, value)
}

func (m *TMemo) ParentCtl3D() bool {
    return Memo_GetParentCtl3D(m.instance)
}

func (m *TMemo) SetParentCtl3D(value bool) {
    Memo_SetParentCtl3D(m.instance, value)
}

func (m *TMemo) ParentDoubleBuffered() bool {
    return Memo_GetParentDoubleBuffered(m.instance)
}

func (m *TMemo) SetParentDoubleBuffered(value bool) {
    Memo_SetParentDoubleBuffered(m.instance, value)
}

func (m *TMemo) ParentFont() bool {
    return Memo_GetParentFont(m.instance)
}

func (m *TMemo) SetParentFont(value bool) {
    Memo_SetParentFont(m.instance, value)
}

func (m *TMemo) ParentShowHint() bool {
    return Memo_GetParentShowHint(m.instance)
}

func (m *TMemo) SetParentShowHint(value bool) {
    Memo_SetParentShowHint(m.instance, value)
}

func (m *TMemo) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Memo_GetPopupMenu(m.instance))
}

func (m *TMemo) SetPopupMenu(value IComponent) {
    Memo_SetPopupMenu(m.instance, CheckPtr(value))
}

func (m *TMemo) ReadOnly() bool {
    return Memo_GetReadOnly(m.instance)
}

func (m *TMemo) SetReadOnly(value bool) {
    Memo_SetReadOnly(m.instance, value)
}

func (m *TMemo) ScrollBars() TScrollStyle {
    return Memo_GetScrollBars(m.instance)
}

func (m *TMemo) SetScrollBars(value TScrollStyle) {
    Memo_SetScrollBars(m.instance, value)
}

func (m *TMemo) ShowHint() bool {
    return Memo_GetShowHint(m.instance)
}

func (m *TMemo) SetShowHint(value bool) {
    Memo_SetShowHint(m.instance, value)
}

func (m *TMemo) TabOrder() uint16 {
    return Memo_GetTabOrder(m.instance)
}

func (m *TMemo) SetTabOrder(value uint16) {
    Memo_SetTabOrder(m.instance, value)
}

func (m *TMemo) TabStop() bool {
    return Memo_GetTabStop(m.instance)
}

func (m *TMemo) SetTabStop(value bool) {
    Memo_SetTabStop(m.instance, value)
}

func (m *TMemo) Visible() bool {
    return Memo_GetVisible(m.instance)
}

func (m *TMemo) SetVisible(value bool) {
    Memo_SetVisible(m.instance, value)
}

func (m *TMemo) WantReturns() bool {
    return Memo_GetWantReturns(m.instance)
}

func (m *TMemo) SetWantReturns(value bool) {
    Memo_SetWantReturns(m.instance, value)
}

func (m *TMemo) WantTabs() bool {
    return Memo_GetWantTabs(m.instance)
}

func (m *TMemo) SetWantTabs(value bool) {
    Memo_SetWantTabs(m.instance, value)
}

func (m *TMemo) WordWrap() bool {
    return Memo_GetWordWrap(m.instance)
}

func (m *TMemo) SetWordWrap(value bool) {
    Memo_SetWordWrap(m.instance, value)
}

func (m *TMemo) StyleElements() TStyleElements {
    return Memo_GetStyleElements(m.instance)
}

func (m *TMemo) SetStyleElements(value TStyleElements) {
    Memo_SetStyleElements(m.instance, value)
}

func (m *TMemo) SetOnChange(fn TNotifyEvent) {
    Memo_SetOnChange(m.instance, fn)
}

func (m *TMemo) SetOnClick(fn TNotifyEvent) {
    Memo_SetOnClick(m.instance, fn)
}

func (m *TMemo) SetOnContextPopup(fn TContextPopupEvent) {
    Memo_SetOnContextPopup(m.instance, fn)
}

func (m *TMemo) SetOnDblClick(fn TNotifyEvent) {
    Memo_SetOnDblClick(m.instance, fn)
}

func (m *TMemo) SetOnDragDrop(fn TDragDropEvent) {
    Memo_SetOnDragDrop(m.instance, fn)
}

func (m *TMemo) SetOnDragOver(fn TDragOverEvent) {
    Memo_SetOnDragOver(m.instance, fn)
}

func (m *TMemo) SetOnEndDock(fn TEndDragEvent) {
    Memo_SetOnEndDock(m.instance, fn)
}

func (m *TMemo) SetOnEndDrag(fn TEndDragEvent) {
    Memo_SetOnEndDrag(m.instance, fn)
}

func (m *TMemo) SetOnEnter(fn TNotifyEvent) {
    Memo_SetOnEnter(m.instance, fn)
}

func (m *TMemo) SetOnExit(fn TNotifyEvent) {
    Memo_SetOnExit(m.instance, fn)
}

func (m *TMemo) SetOnKeyDown(fn TKeyEvent) {
    Memo_SetOnKeyDown(m.instance, fn)
}

func (m *TMemo) SetOnKeyPress(fn TKeyPressEvent) {
    Memo_SetOnKeyPress(m.instance, fn)
}

func (m *TMemo) SetOnKeyUp(fn TKeyEvent) {
    Memo_SetOnKeyUp(m.instance, fn)
}

func (m *TMemo) SetOnMouseDown(fn TMouseEvent) {
    Memo_SetOnMouseDown(m.instance, fn)
}

func (m *TMemo) SetOnMouseEnter(fn TNotifyEvent) {
    Memo_SetOnMouseEnter(m.instance, fn)
}

func (m *TMemo) SetOnMouseLeave(fn TNotifyEvent) {
    Memo_SetOnMouseLeave(m.instance, fn)
}

func (m *TMemo) SetOnMouseMove(fn TMouseMoveEvent) {
    Memo_SetOnMouseMove(m.instance, fn)
}

func (m *TMemo) SetOnMouseUp(fn TMouseEvent) {
    Memo_SetOnMouseUp(m.instance, fn)
}

func (m *TMemo) SetOnStartDock(fn TStartDockEvent) {
    Memo_SetOnStartDock(m.instance, fn)
}

func (m *TMemo) CaretPos() TPoint {
    return Memo_GetCaretPos(m.instance)
}

func (m *TMemo) SetCaretPos(value TPoint) {
    Memo_SetCaretPos(m.instance, value)
}

func (m *TMemo) CanUndo() bool {
    return Memo_GetCanUndo(m.instance)
}

func (m *TMemo) Modified() bool {
    return Memo_GetModified(m.instance)
}

func (m *TMemo) SetModified(value bool) {
    Memo_SetModified(m.instance, value)
}

func (m *TMemo) SelLength() int32 {
    return Memo_GetSelLength(m.instance)
}

func (m *TMemo) SetSelLength(value int32) {
    Memo_SetSelLength(m.instance, value)
}

func (m *TMemo) SelStart() int32 {
    return Memo_GetSelStart(m.instance)
}

func (m *TMemo) SetSelStart(value int32) {
    Memo_SetSelStart(m.instance, value)
}

func (m *TMemo) SelText() string {
    return Memo_GetSelText(m.instance)
}

func (m *TMemo) SetSelText(value string) {
    Memo_SetSelText(m.instance, value)
}

func (m *TMemo) Text() string {
    return Memo_GetText(m.instance)
}

func (m *TMemo) SetText(value string) {
    Memo_SetText(m.instance, value)
}

func (m *TMemo) TextHint() string {
    return Memo_GetTextHint(m.instance)
}

func (m *TMemo) SetTextHint(value string) {
    Memo_SetTextHint(m.instance, value)
}

func (m *TMemo) DockSite() bool {
    return Memo_GetDockSite(m.instance)
}

func (m *TMemo) SetDockSite(value bool) {
    Memo_SetDockSite(m.instance, value)
}

func (m *TMemo) Brush() *TBrush {
    return BrushFromInst(Memo_GetBrush(m.instance))
}

func (m *TMemo) ControlCount() int32 {
    return Memo_GetControlCount(m.instance)
}

func (m *TMemo) Handle() HWND {
    return Memo_GetHandle(m.instance)
}

func (m *TMemo) ParentWindow() HWND {
    return Memo_GetParentWindow(m.instance)
}

func (m *TMemo) SetParentWindow(value HWND) {
    Memo_SetParentWindow(m.instance, value)
}

func (m *TMemo) UseDockManager() bool {
    return Memo_GetUseDockManager(m.instance)
}

func (m *TMemo) SetUseDockManager(value bool) {
    Memo_SetUseDockManager(m.instance, value)
}

func (m *TMemo) Action() *TAction {
    return ActionFromInst(Memo_GetAction(m.instance))
}

func (m *TMemo) SetAction(value IComponent) {
    Memo_SetAction(m.instance, CheckPtr(value))
}

func (m *TMemo) BoundsRect() TRect {
    return Memo_GetBoundsRect(m.instance)
}

func (m *TMemo) SetBoundsRect(value TRect) {
    Memo_SetBoundsRect(m.instance, value)
}

func (m *TMemo) ClientHeight() int32 {
    return Memo_GetClientHeight(m.instance)
}

func (m *TMemo) SetClientHeight(value int32) {
    Memo_SetClientHeight(m.instance, value)
}

func (m *TMemo) ClientRect() TRect {
    return Memo_GetClientRect(m.instance)
}

func (m *TMemo) ClientWidth() int32 {
    return Memo_GetClientWidth(m.instance)
}

func (m *TMemo) SetClientWidth(value int32) {
    Memo_SetClientWidth(m.instance, value)
}

func (m *TMemo) ExplicitLeft() int32 {
    return Memo_GetExplicitLeft(m.instance)
}

func (m *TMemo) ExplicitTop() int32 {
    return Memo_GetExplicitTop(m.instance)
}

func (m *TMemo) ExplicitWidth() int32 {
    return Memo_GetExplicitWidth(m.instance)
}

func (m *TMemo) ExplicitHeight() int32 {
    return Memo_GetExplicitHeight(m.instance)
}

func (m *TMemo) Floating() bool {
    return Memo_GetFloating(m.instance)
}

func (m *TMemo) Parent() *TControl {
    return ControlFromInst(Memo_GetParent(m.instance))
}

func (m *TMemo) SetParent(value IControl) {
    Memo_SetParent(m.instance, CheckPtr(value))
}

func (m *TMemo) AlignWithMargins() bool {
    return Memo_GetAlignWithMargins(m.instance)
}

func (m *TMemo) SetAlignWithMargins(value bool) {
    Memo_SetAlignWithMargins(m.instance, value)
}

func (m *TMemo) Left() int32 {
    return Memo_GetLeft(m.instance)
}

func (m *TMemo) SetLeft(value int32) {
    Memo_SetLeft(m.instance, value)
}

func (m *TMemo) Top() int32 {
    return Memo_GetTop(m.instance)
}

func (m *TMemo) SetTop(value int32) {
    Memo_SetTop(m.instance, value)
}

func (m *TMemo) Width() int32 {
    return Memo_GetWidth(m.instance)
}

func (m *TMemo) SetWidth(value int32) {
    Memo_SetWidth(m.instance, value)
}

func (m *TMemo) Height() int32 {
    return Memo_GetHeight(m.instance)
}

func (m *TMemo) SetHeight(value int32) {
    Memo_SetHeight(m.instance, value)
}

func (m *TMemo) Cursor() TCursor {
    return Memo_GetCursor(m.instance)
}

func (m *TMemo) SetCursor(value TCursor) {
    Memo_SetCursor(m.instance, value)
}

func (m *TMemo) Hint() string {
    return Memo_GetHint(m.instance)
}

func (m *TMemo) SetHint(value string) {
    Memo_SetHint(m.instance, value)
}

func (m *TMemo) Margins() *TMargins {
    return MarginsFromInst(Memo_GetMargins(m.instance))
}

func (m *TMemo) SetMargins(value *TMargins) {
    Memo_SetMargins(m.instance, CheckPtr(value))
}

func (m *TMemo) CustomHint() *TCustomHint {
    return CustomHintFromInst(Memo_GetCustomHint(m.instance))
}

func (m *TMemo) SetCustomHint(value IComponent) {
    Memo_SetCustomHint(m.instance, CheckPtr(value))
}

func (m *TMemo) ComponentCount() int32 {
    return Memo_GetComponentCount(m.instance)
}

func (m *TMemo) ComponentIndex() int32 {
    return Memo_GetComponentIndex(m.instance)
}

func (m *TMemo) SetComponentIndex(value int32) {
    Memo_SetComponentIndex(m.instance, value)
}

func (m *TMemo) Owner() *TComponent {
    return ComponentFromInst(Memo_GetOwner(m.instance))
}

func (m *TMemo) Name() string {
    return Memo_GetName(m.instance)
}

func (m *TMemo) SetName(value string) {
    Memo_SetName(m.instance, value)
}

func (m *TMemo) Tag() int {
    return Memo_GetTag(m.instance)
}

func (m *TMemo) SetTag(value int) {
    Memo_SetTag(m.instance, value)
}

func (m *TMemo) Controls(Index int32) *TControl {
    return ControlFromInst(Memo_GetControls(m.instance, Index))
}

func (m *TMemo) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Memo_GetComponents(m.instance, AIndex))
}

