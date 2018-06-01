
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

type TMaskEdit struct {
    IControl
    instance uintptr
}

func NewMaskEdit(owner IComponent) *TMaskEdit {
    m := new(TMaskEdit)
    m.instance = MaskEdit_Create(CheckPtr(owner))
    return m
}

func MaskEditFromInst(inst uintptr) *TMaskEdit {
    m := new(TMaskEdit)
    m.instance = inst
    return m
}

func MaskEditFromObj(obj IObject) *TMaskEdit {
    m := new(TMaskEdit)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMaskEdit) Free() {
    if m.instance != 0 {
        MaskEdit_Free(m.instance)
        m.instance = 0
    }
}

func (m *TMaskEdit) Instance() uintptr {
    return m.instance
}

func (m *TMaskEdit) IsValid() bool {
    return m.instance != 0
}

func (m *TMaskEdit) ValidateEdit() {
    MaskEdit_ValidateEdit(m.instance)
}

func (m *TMaskEdit) Clear() {
    MaskEdit_Clear(m.instance)
}

func (m *TMaskEdit) GetTextLen() int32 {
    return MaskEdit_GetTextLen(m.instance)
}

func (m *TMaskEdit) ClearSelection() {
    MaskEdit_ClearSelection(m.instance)
}

func (m *TMaskEdit) CopyToClipboard() {
    MaskEdit_CopyToClipboard(m.instance)
}

func (m *TMaskEdit) CutToClipboard() {
    MaskEdit_CutToClipboard(m.instance)
}

func (m *TMaskEdit) PasteFromClipboard() {
    MaskEdit_PasteFromClipboard(m.instance)
}

func (m *TMaskEdit) Undo() {
    MaskEdit_Undo(m.instance)
}

func (m *TMaskEdit) ClearUndo() {
    MaskEdit_ClearUndo(m.instance)
}

func (m *TMaskEdit) SelectAll() {
    MaskEdit_SelectAll(m.instance)
}

func (m *TMaskEdit) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return MaskEdit_GetSelTextBuf(m.instance, Buffer , BufSize)
}

func (m *TMaskEdit) CanFocus() bool {
    return MaskEdit_CanFocus(m.instance)
}

func (m *TMaskEdit) FlipChildren(AllLevels bool) {
    MaskEdit_FlipChildren(m.instance, AllLevels)
}

func (m *TMaskEdit) Focused() bool {
    return MaskEdit_Focused(m.instance)
}

func (m *TMaskEdit) HandleAllocated() bool {
    return MaskEdit_HandleAllocated(m.instance)
}

func (m *TMaskEdit) Invalidate() {
    MaskEdit_Invalidate(m.instance)
}

func (m *TMaskEdit) Realign() {
    MaskEdit_Realign(m.instance)
}

func (m *TMaskEdit) Repaint() {
    MaskEdit_Repaint(m.instance)
}

func (m *TMaskEdit) ScaleBy(M int32, D int32) {
    MaskEdit_ScaleBy(m.instance, M , D)
}

func (m *TMaskEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    MaskEdit_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

func (m *TMaskEdit) SetFocus() {
    MaskEdit_SetFocus(m.instance)
}

func (m *TMaskEdit) Update() {
    MaskEdit_Update(m.instance)
}

func (m *TMaskEdit) BringToFront() {
    MaskEdit_BringToFront(m.instance)
}

func (m *TMaskEdit) ClientToScreen(Point TPoint) TPoint {
    return MaskEdit_ClientToScreen(m.instance, Point)
}

func (m *TMaskEdit) ClientToParent(Point TPoint, AParent IControl) TPoint {
    return MaskEdit_ClientToParent(m.instance, Point , CheckPtr(AParent))
}

func (m *TMaskEdit) Dragging() bool {
    return MaskEdit_Dragging(m.instance)
}

func (m *TMaskEdit) HasParent() bool {
    return MaskEdit_HasParent(m.instance)
}

func (m *TMaskEdit) Hide() {
    MaskEdit_Hide(m.instance)
}

func (m *TMaskEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return MaskEdit_Perform(m.instance, Msg , WParam , LParam)
}

func (m *TMaskEdit) Refresh() {
    MaskEdit_Refresh(m.instance)
}

func (m *TMaskEdit) ScreenToClient(Point TPoint) TPoint {
    return MaskEdit_ScreenToClient(m.instance, Point)
}

func (m *TMaskEdit) ParentToClient(Point TPoint, AParent IControl) TPoint {
    return MaskEdit_ParentToClient(m.instance, Point , CheckPtr(AParent))
}

func (m *TMaskEdit) SendToBack() {
    MaskEdit_SendToBack(m.instance)
}

func (m *TMaskEdit) Show() {
    MaskEdit_Show(m.instance)
}

func (m *TMaskEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    return MaskEdit_GetTextBuf(m.instance, Buffer , BufSize)
}

func (m *TMaskEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(MaskEdit_FindComponent(m.instance, AName))
}

func (m *TMaskEdit) GetNamePath() string {
    return MaskEdit_GetNamePath(m.instance)
}

func (m *TMaskEdit) Assign(Source IObject) {
    MaskEdit_Assign(m.instance, CheckPtr(Source))
}

func (m *TMaskEdit) ClassName() string {
    return MaskEdit_ClassName(m.instance)
}

func (m *TMaskEdit) Equals(Obj IObject) bool {
    return MaskEdit_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMaskEdit) GetHashCode() int32 {
    return MaskEdit_GetHashCode(m.instance)
}

func (m *TMaskEdit) ToString() string {
    return MaskEdit_ToString(m.instance)
}

func (m *TMaskEdit) Align() TAlign {
    return MaskEdit_GetAlign(m.instance)
}

func (m *TMaskEdit) SetAlign(value TAlign) {
    MaskEdit_SetAlign(m.instance, value)
}

func (m *TMaskEdit) Alignment() TAlignment {
    return MaskEdit_GetAlignment(m.instance)
}

func (m *TMaskEdit) SetAlignment(value TAlignment) {
    MaskEdit_SetAlignment(m.instance, value)
}

func (m *TMaskEdit) Anchors() TAnchors {
    return MaskEdit_GetAnchors(m.instance)
}

func (m *TMaskEdit) SetAnchors(value TAnchors) {
    MaskEdit_SetAnchors(m.instance, value)
}

func (m *TMaskEdit) AutoSelect() bool {
    return MaskEdit_GetAutoSelect(m.instance)
}

func (m *TMaskEdit) SetAutoSelect(value bool) {
    MaskEdit_SetAutoSelect(m.instance, value)
}

func (m *TMaskEdit) AutoSize() bool {
    return MaskEdit_GetAutoSize(m.instance)
}

func (m *TMaskEdit) SetAutoSize(value bool) {
    MaskEdit_SetAutoSize(m.instance, value)
}

func (m *TMaskEdit) BevelEdges() TBevelEdges {
    return MaskEdit_GetBevelEdges(m.instance)
}

func (m *TMaskEdit) SetBevelEdges(value TBevelEdges) {
    MaskEdit_SetBevelEdges(m.instance, value)
}

func (m *TMaskEdit) BevelInner() TBevelCut {
    return MaskEdit_GetBevelInner(m.instance)
}

func (m *TMaskEdit) SetBevelInner(value TBevelCut) {
    MaskEdit_SetBevelInner(m.instance, value)
}

func (m *TMaskEdit) BevelOuter() TBevelCut {
    return MaskEdit_GetBevelOuter(m.instance)
}

func (m *TMaskEdit) SetBevelOuter(value TBevelCut) {
    MaskEdit_SetBevelOuter(m.instance, value)
}

func (m *TMaskEdit) BevelKind() TBevelKind {
    return MaskEdit_GetBevelKind(m.instance)
}

func (m *TMaskEdit) SetBevelKind(value TBevelKind) {
    MaskEdit_SetBevelKind(m.instance, value)
}

func (m *TMaskEdit) BiDiMode() TBiDiMode {
    return MaskEdit_GetBiDiMode(m.instance)
}

func (m *TMaskEdit) SetBiDiMode(value TBiDiMode) {
    MaskEdit_SetBiDiMode(m.instance, value)
}

func (m *TMaskEdit) BorderStyle() TBorderStyle {
    return MaskEdit_GetBorderStyle(m.instance)
}

func (m *TMaskEdit) SetBorderStyle(value TBorderStyle) {
    MaskEdit_SetBorderStyle(m.instance, value)
}

func (m *TMaskEdit) Color() TColor {
    return MaskEdit_GetColor(m.instance)
}

func (m *TMaskEdit) SetColor(value TColor) {
    MaskEdit_SetColor(m.instance, value)
}

func (m *TMaskEdit) DoubleBuffered() bool {
    return MaskEdit_GetDoubleBuffered(m.instance)
}

func (m *TMaskEdit) SetDoubleBuffered(value bool) {
    MaskEdit_SetDoubleBuffered(m.instance, value)
}

func (m *TMaskEdit) DragCursor() TCursor {
    return MaskEdit_GetDragCursor(m.instance)
}

func (m *TMaskEdit) SetDragCursor(value TCursor) {
    MaskEdit_SetDragCursor(m.instance, value)
}

func (m *TMaskEdit) DragKind() TDragKind {
    return MaskEdit_GetDragKind(m.instance)
}

func (m *TMaskEdit) SetDragKind(value TDragKind) {
    MaskEdit_SetDragKind(m.instance, value)
}

func (m *TMaskEdit) DragMode() TDragMode {
    return MaskEdit_GetDragMode(m.instance)
}

func (m *TMaskEdit) SetDragMode(value TDragMode) {
    MaskEdit_SetDragMode(m.instance, value)
}

func (m *TMaskEdit) Enabled() bool {
    return MaskEdit_GetEnabled(m.instance)
}

func (m *TMaskEdit) SetEnabled(value bool) {
    MaskEdit_SetEnabled(m.instance, value)
}

func (m *TMaskEdit) Font() *TFont {
    return FontFromInst(MaskEdit_GetFont(m.instance))
}

func (m *TMaskEdit) SetFont(value *TFont) {
    MaskEdit_SetFont(m.instance, CheckPtr(value))
}

func (m *TMaskEdit) MaxLength() int32 {
    return MaskEdit_GetMaxLength(m.instance)
}

func (m *TMaskEdit) SetMaxLength(value int32) {
    MaskEdit_SetMaxLength(m.instance, value)
}

func (m *TMaskEdit) ParentColor() bool {
    return MaskEdit_GetParentColor(m.instance)
}

func (m *TMaskEdit) SetParentColor(value bool) {
    MaskEdit_SetParentColor(m.instance, value)
}

func (m *TMaskEdit) ParentCtl3D() bool {
    return MaskEdit_GetParentCtl3D(m.instance)
}

func (m *TMaskEdit) SetParentCtl3D(value bool) {
    MaskEdit_SetParentCtl3D(m.instance, value)
}

func (m *TMaskEdit) ParentDoubleBuffered() bool {
    return MaskEdit_GetParentDoubleBuffered(m.instance)
}

func (m *TMaskEdit) SetParentDoubleBuffered(value bool) {
    MaskEdit_SetParentDoubleBuffered(m.instance, value)
}

func (m *TMaskEdit) ParentFont() bool {
    return MaskEdit_GetParentFont(m.instance)
}

func (m *TMaskEdit) SetParentFont(value bool) {
    MaskEdit_SetParentFont(m.instance, value)
}

func (m *TMaskEdit) ParentShowHint() bool {
    return MaskEdit_GetParentShowHint(m.instance)
}

func (m *TMaskEdit) SetParentShowHint(value bool) {
    MaskEdit_SetParentShowHint(m.instance, value)
}

func (m *TMaskEdit) PasswordChar() uint16 {
    return MaskEdit_GetPasswordChar(m.instance)
}

func (m *TMaskEdit) SetPasswordChar(value uint16) {
    MaskEdit_SetPasswordChar(m.instance, value)
}

func (m *TMaskEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(MaskEdit_GetPopupMenu(m.instance))
}

func (m *TMaskEdit) SetPopupMenu(value IComponent) {
    MaskEdit_SetPopupMenu(m.instance, CheckPtr(value))
}

func (m *TMaskEdit) ReadOnly() bool {
    return MaskEdit_GetReadOnly(m.instance)
}

func (m *TMaskEdit) SetReadOnly(value bool) {
    MaskEdit_SetReadOnly(m.instance, value)
}

func (m *TMaskEdit) ShowHint() bool {
    return MaskEdit_GetShowHint(m.instance)
}

func (m *TMaskEdit) SetShowHint(value bool) {
    MaskEdit_SetShowHint(m.instance, value)
}

func (m *TMaskEdit) TabOrder() uint16 {
    return MaskEdit_GetTabOrder(m.instance)
}

func (m *TMaskEdit) SetTabOrder(value uint16) {
    MaskEdit_SetTabOrder(m.instance, value)
}

func (m *TMaskEdit) TabStop() bool {
    return MaskEdit_GetTabStop(m.instance)
}

func (m *TMaskEdit) SetTabStop(value bool) {
    MaskEdit_SetTabStop(m.instance, value)
}

func (m *TMaskEdit) Text() string {
    return MaskEdit_GetText(m.instance)
}

func (m *TMaskEdit) SetText(value string) {
    MaskEdit_SetText(m.instance, value)
}

func (m *TMaskEdit) TextHint() string {
    return MaskEdit_GetTextHint(m.instance)
}

func (m *TMaskEdit) SetTextHint(value string) {
    MaskEdit_SetTextHint(m.instance, value)
}

func (m *TMaskEdit) Visible() bool {
    return MaskEdit_GetVisible(m.instance)
}

func (m *TMaskEdit) SetVisible(value bool) {
    MaskEdit_SetVisible(m.instance, value)
}

func (m *TMaskEdit) StyleElements() TStyleElements {
    return MaskEdit_GetStyleElements(m.instance)
}

func (m *TMaskEdit) SetStyleElements(value TStyleElements) {
    MaskEdit_SetStyleElements(m.instance, value)
}

func (m *TMaskEdit) SetOnChange(fn TNotifyEvent) {
    MaskEdit_SetOnChange(m.instance, fn)
}

func (m *TMaskEdit) SetOnClick(fn TNotifyEvent) {
    MaskEdit_SetOnClick(m.instance, fn)
}

func (m *TMaskEdit) SetOnDblClick(fn TNotifyEvent) {
    MaskEdit_SetOnDblClick(m.instance, fn)
}

func (m *TMaskEdit) SetOnDragDrop(fn TDragDropEvent) {
    MaskEdit_SetOnDragDrop(m.instance, fn)
}

func (m *TMaskEdit) SetOnDragOver(fn TDragOverEvent) {
    MaskEdit_SetOnDragOver(m.instance, fn)
}

func (m *TMaskEdit) SetOnEndDock(fn TEndDragEvent) {
    MaskEdit_SetOnEndDock(m.instance, fn)
}

func (m *TMaskEdit) SetOnEndDrag(fn TEndDragEvent) {
    MaskEdit_SetOnEndDrag(m.instance, fn)
}

func (m *TMaskEdit) SetOnEnter(fn TNotifyEvent) {
    MaskEdit_SetOnEnter(m.instance, fn)
}

func (m *TMaskEdit) SetOnExit(fn TNotifyEvent) {
    MaskEdit_SetOnExit(m.instance, fn)
}

func (m *TMaskEdit) SetOnKeyDown(fn TKeyEvent) {
    MaskEdit_SetOnKeyDown(m.instance, fn)
}

func (m *TMaskEdit) SetOnKeyPress(fn TKeyPressEvent) {
    MaskEdit_SetOnKeyPress(m.instance, fn)
}

func (m *TMaskEdit) SetOnKeyUp(fn TKeyEvent) {
    MaskEdit_SetOnKeyUp(m.instance, fn)
}

func (m *TMaskEdit) SetOnMouseDown(fn TMouseEvent) {
    MaskEdit_SetOnMouseDown(m.instance, fn)
}

func (m *TMaskEdit) SetOnMouseEnter(fn TNotifyEvent) {
    MaskEdit_SetOnMouseEnter(m.instance, fn)
}

func (m *TMaskEdit) SetOnMouseLeave(fn TNotifyEvent) {
    MaskEdit_SetOnMouseLeave(m.instance, fn)
}

func (m *TMaskEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    MaskEdit_SetOnMouseMove(m.instance, fn)
}

func (m *TMaskEdit) SetOnMouseUp(fn TMouseEvent) {
    MaskEdit_SetOnMouseUp(m.instance, fn)
}

func (m *TMaskEdit) SetOnStartDock(fn TStartDockEvent) {
    MaskEdit_SetOnStartDock(m.instance, fn)
}

func (m *TMaskEdit) IsMasked() bool {
    return MaskEdit_GetIsMasked(m.instance)
}

func (m *TMaskEdit) EditText() string {
    return MaskEdit_GetEditText(m.instance)
}

func (m *TMaskEdit) SetEditText(value string) {
    MaskEdit_SetEditText(m.instance, value)
}

func (m *TMaskEdit) CanUndo() bool {
    return MaskEdit_GetCanUndo(m.instance)
}

func (m *TMaskEdit) Modified() bool {
    return MaskEdit_GetModified(m.instance)
}

func (m *TMaskEdit) SetModified(value bool) {
    MaskEdit_SetModified(m.instance, value)
}

func (m *TMaskEdit) SelLength() int32 {
    return MaskEdit_GetSelLength(m.instance)
}

func (m *TMaskEdit) SetSelLength(value int32) {
    MaskEdit_SetSelLength(m.instance, value)
}

func (m *TMaskEdit) SelStart() int32 {
    return MaskEdit_GetSelStart(m.instance)
}

func (m *TMaskEdit) SetSelStart(value int32) {
    MaskEdit_SetSelStart(m.instance, value)
}

func (m *TMaskEdit) SelText() string {
    return MaskEdit_GetSelText(m.instance)
}

func (m *TMaskEdit) SetSelText(value string) {
    MaskEdit_SetSelText(m.instance, value)
}

func (m *TMaskEdit) DockSite() bool {
    return MaskEdit_GetDockSite(m.instance)
}

func (m *TMaskEdit) SetDockSite(value bool) {
    MaskEdit_SetDockSite(m.instance, value)
}

func (m *TMaskEdit) Brush() *TBrush {
    return BrushFromInst(MaskEdit_GetBrush(m.instance))
}

func (m *TMaskEdit) ControlCount() int32 {
    return MaskEdit_GetControlCount(m.instance)
}

func (m *TMaskEdit) Handle() HWND {
    return MaskEdit_GetHandle(m.instance)
}

func (m *TMaskEdit) ParentWindow() HWND {
    return MaskEdit_GetParentWindow(m.instance)
}

func (m *TMaskEdit) SetParentWindow(value HWND) {
    MaskEdit_SetParentWindow(m.instance, value)
}

func (m *TMaskEdit) UseDockManager() bool {
    return MaskEdit_GetUseDockManager(m.instance)
}

func (m *TMaskEdit) SetUseDockManager(value bool) {
    MaskEdit_SetUseDockManager(m.instance, value)
}

func (m *TMaskEdit) Action() *TAction {
    return ActionFromInst(MaskEdit_GetAction(m.instance))
}

func (m *TMaskEdit) SetAction(value IComponent) {
    MaskEdit_SetAction(m.instance, CheckPtr(value))
}

func (m *TMaskEdit) BoundsRect() TRect {
    return MaskEdit_GetBoundsRect(m.instance)
}

func (m *TMaskEdit) SetBoundsRect(value TRect) {
    MaskEdit_SetBoundsRect(m.instance, value)
}

func (m *TMaskEdit) ClientHeight() int32 {
    return MaskEdit_GetClientHeight(m.instance)
}

func (m *TMaskEdit) SetClientHeight(value int32) {
    MaskEdit_SetClientHeight(m.instance, value)
}

func (m *TMaskEdit) ClientRect() TRect {
    return MaskEdit_GetClientRect(m.instance)
}

func (m *TMaskEdit) ClientWidth() int32 {
    return MaskEdit_GetClientWidth(m.instance)
}

func (m *TMaskEdit) SetClientWidth(value int32) {
    MaskEdit_SetClientWidth(m.instance, value)
}

func (m *TMaskEdit) ExplicitLeft() int32 {
    return MaskEdit_GetExplicitLeft(m.instance)
}

func (m *TMaskEdit) ExplicitTop() int32 {
    return MaskEdit_GetExplicitTop(m.instance)
}

func (m *TMaskEdit) ExplicitWidth() int32 {
    return MaskEdit_GetExplicitWidth(m.instance)
}

func (m *TMaskEdit) ExplicitHeight() int32 {
    return MaskEdit_GetExplicitHeight(m.instance)
}

func (m *TMaskEdit) Floating() bool {
    return MaskEdit_GetFloating(m.instance)
}

func (m *TMaskEdit) Parent() *TControl {
    return ControlFromInst(MaskEdit_GetParent(m.instance))
}

func (m *TMaskEdit) SetParent(value IControl) {
    MaskEdit_SetParent(m.instance, CheckPtr(value))
}

func (m *TMaskEdit) AlignWithMargins() bool {
    return MaskEdit_GetAlignWithMargins(m.instance)
}

func (m *TMaskEdit) SetAlignWithMargins(value bool) {
    MaskEdit_SetAlignWithMargins(m.instance, value)
}

func (m *TMaskEdit) Left() int32 {
    return MaskEdit_GetLeft(m.instance)
}

func (m *TMaskEdit) SetLeft(value int32) {
    MaskEdit_SetLeft(m.instance, value)
}

func (m *TMaskEdit) Top() int32 {
    return MaskEdit_GetTop(m.instance)
}

func (m *TMaskEdit) SetTop(value int32) {
    MaskEdit_SetTop(m.instance, value)
}

func (m *TMaskEdit) Width() int32 {
    return MaskEdit_GetWidth(m.instance)
}

func (m *TMaskEdit) SetWidth(value int32) {
    MaskEdit_SetWidth(m.instance, value)
}

func (m *TMaskEdit) Height() int32 {
    return MaskEdit_GetHeight(m.instance)
}

func (m *TMaskEdit) SetHeight(value int32) {
    MaskEdit_SetHeight(m.instance, value)
}

func (m *TMaskEdit) Cursor() TCursor {
    return MaskEdit_GetCursor(m.instance)
}

func (m *TMaskEdit) SetCursor(value TCursor) {
    MaskEdit_SetCursor(m.instance, value)
}

func (m *TMaskEdit) Hint() string {
    return MaskEdit_GetHint(m.instance)
}

func (m *TMaskEdit) SetHint(value string) {
    MaskEdit_SetHint(m.instance, value)
}

func (m *TMaskEdit) Margins() *TMargins {
    return MarginsFromInst(MaskEdit_GetMargins(m.instance))
}

func (m *TMaskEdit) SetMargins(value *TMargins) {
    MaskEdit_SetMargins(m.instance, CheckPtr(value))
}

func (m *TMaskEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(MaskEdit_GetCustomHint(m.instance))
}

func (m *TMaskEdit) SetCustomHint(value IComponent) {
    MaskEdit_SetCustomHint(m.instance, CheckPtr(value))
}

func (m *TMaskEdit) ComponentCount() int32 {
    return MaskEdit_GetComponentCount(m.instance)
}

func (m *TMaskEdit) ComponentIndex() int32 {
    return MaskEdit_GetComponentIndex(m.instance)
}

func (m *TMaskEdit) SetComponentIndex(value int32) {
    MaskEdit_SetComponentIndex(m.instance, value)
}

func (m *TMaskEdit) Owner() *TComponent {
    return ComponentFromInst(MaskEdit_GetOwner(m.instance))
}

func (m *TMaskEdit) Name() string {
    return MaskEdit_GetName(m.instance)
}

func (m *TMaskEdit) SetName(value string) {
    MaskEdit_SetName(m.instance, value)
}

func (m *TMaskEdit) Tag() int {
    return MaskEdit_GetTag(m.instance)
}

func (m *TMaskEdit) SetTag(value int) {
    MaskEdit_SetTag(m.instance, value)
}

func (m *TMaskEdit) Controls(Index int32) *TControl {
    return ControlFromInst(MaskEdit_GetControls(m.instance, Index))
}

func (m *TMaskEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(MaskEdit_GetComponents(m.instance, AIndex))
}

