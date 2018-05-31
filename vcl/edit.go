
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

type TEdit struct {
    IControl
    instance uintptr
}

func NewEdit(owner IComponent) *TEdit {
    e := new(TEdit)
    e.instance = Edit_Create(CheckPtr(owner))
    return e
}

func EditFromInst(inst uintptr) *TEdit {
    e := new(TEdit)
    e.instance = inst
    return e
}

func EditFromObj(obj IObject) *TEdit {
    e := new(TEdit)
    e.instance = CheckPtr(obj)
    return e
}

func (e *TEdit) Free() {
    if e.instance != 0 {
        Edit_Free(e.instance)
        e.instance = 0
    }
}

func (e *TEdit) Instance() uintptr {
    return e.instance
}

func (e *TEdit) IsValid() bool {
    return e.instance != 0
}

func (e *TEdit) Clear() {
    Edit_Clear(e.instance)
}

func (e *TEdit) ClearSelection() {
    Edit_ClearSelection(e.instance)
}

func (e *TEdit) CopyToClipboard() {
    Edit_CopyToClipboard(e.instance)
}

func (e *TEdit) CutToClipboard() {
    Edit_CutToClipboard(e.instance)
}

func (e *TEdit) PasteFromClipboard() {
    Edit_PasteFromClipboard(e.instance)
}

func (e *TEdit) Undo() {
    Edit_Undo(e.instance)
}

func (e *TEdit) ClearUndo() {
    Edit_ClearUndo(e.instance)
}

func (e *TEdit) SelectAll() {
    Edit_SelectAll(e.instance)
}

func (e *TEdit) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return Edit_GetSelTextBuf(e.instance, Buffer , BufSize)
}

func (e *TEdit) CanFocus() bool {
    return Edit_CanFocus(e.instance)
}

func (e *TEdit) FlipChildren(AllLevels bool) {
    Edit_FlipChildren(e.instance, AllLevels)
}

func (e *TEdit) Focused() bool {
    return Edit_Focused(e.instance)
}

func (e *TEdit) HandleAllocated() bool {
    return Edit_HandleAllocated(e.instance)
}

func (e *TEdit) Invalidate() {
    Edit_Invalidate(e.instance)
}

func (e *TEdit) Realign() {
    Edit_Realign(e.instance)
}

func (e *TEdit) Repaint() {
    Edit_Repaint(e.instance)
}

func (e *TEdit) ScaleBy(M int32, D int32) {
    Edit_ScaleBy(e.instance, M , D)
}

func (e *TEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Edit_SetBounds(e.instance, ALeft , ATop , AWidth , AHeight)
}

func (e *TEdit) SetFocus() {
    Edit_SetFocus(e.instance)
}

func (e *TEdit) Update() {
    Edit_Update(e.instance)
}

func (e *TEdit) BringToFront() {
    Edit_BringToFront(e.instance)
}

func (e *TEdit) Dragging() bool {
    return Edit_Dragging(e.instance)
}

func (e *TEdit) HasParent() bool {
    return Edit_HasParent(e.instance)
}

func (e *TEdit) Hide() {
    Edit_Hide(e.instance)
}

func (e *TEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Edit_Perform(e.instance, Msg , WParam , LParam)
}

func (e *TEdit) Refresh() {
    Edit_Refresh(e.instance)
}

func (e *TEdit) SendToBack() {
    Edit_SendToBack(e.instance)
}

func (e *TEdit) Show() {
    Edit_Show(e.instance)
}

func (e *TEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Edit_GetTextBuf(e.instance, Buffer , BufSize)
}

func (e *TEdit) GetTextLen() int32 {
    return Edit_GetTextLen(e.instance)
}

func (e *TEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Edit_FindComponent(e.instance, AName))
}

func (e *TEdit) GetNamePath() string {
    return Edit_GetNamePath(e.instance)
}

func (e *TEdit) Assign(Source IObject) {
    Edit_Assign(e.instance, CheckPtr(Source))
}

func (e *TEdit) ClassName() string {
    return Edit_ClassName(e.instance)
}

func (e *TEdit) Equals(Obj IObject) bool {
    return Edit_Equals(e.instance, CheckPtr(Obj))
}

func (e *TEdit) GetHashCode() int32 {
    return Edit_GetHashCode(e.instance)
}

func (e *TEdit) ToString() string {
    return Edit_ToString(e.instance)
}

func (e *TEdit) Align() TAlign {
    return Edit_GetAlign(e.instance)
}

func (e *TEdit) SetAlign(value TAlign) {
    Edit_SetAlign(e.instance, value)
}

func (e *TEdit) Alignment() TAlignment {
    return Edit_GetAlignment(e.instance)
}

func (e *TEdit) SetAlignment(value TAlignment) {
    Edit_SetAlignment(e.instance, value)
}

func (e *TEdit) Anchors() TAnchors {
    return Edit_GetAnchors(e.instance)
}

func (e *TEdit) SetAnchors(value TAnchors) {
    Edit_SetAnchors(e.instance, value)
}

func (e *TEdit) AutoSelect() bool {
    return Edit_GetAutoSelect(e.instance)
}

func (e *TEdit) SetAutoSelect(value bool) {
    Edit_SetAutoSelect(e.instance, value)
}

func (e *TEdit) AutoSize() bool {
    return Edit_GetAutoSize(e.instance)
}

func (e *TEdit) SetAutoSize(value bool) {
    Edit_SetAutoSize(e.instance, value)
}

func (e *TEdit) BevelEdges() TBevelEdges {
    return Edit_GetBevelEdges(e.instance)
}

func (e *TEdit) SetBevelEdges(value TBevelEdges) {
    Edit_SetBevelEdges(e.instance, value)
}

func (e *TEdit) BevelInner() TBevelCut {
    return Edit_GetBevelInner(e.instance)
}

func (e *TEdit) SetBevelInner(value TBevelCut) {
    Edit_SetBevelInner(e.instance, value)
}

func (e *TEdit) BevelKind() TBevelKind {
    return Edit_GetBevelKind(e.instance)
}

func (e *TEdit) SetBevelKind(value TBevelKind) {
    Edit_SetBevelKind(e.instance, value)
}

func (e *TEdit) BevelOuter() TBevelCut {
    return Edit_GetBevelOuter(e.instance)
}

func (e *TEdit) SetBevelOuter(value TBevelCut) {
    Edit_SetBevelOuter(e.instance, value)
}

func (e *TEdit) BiDiMode() TBiDiMode {
    return Edit_GetBiDiMode(e.instance)
}

func (e *TEdit) SetBiDiMode(value TBiDiMode) {
    Edit_SetBiDiMode(e.instance, value)
}

func (e *TEdit) BorderStyle() TBorderStyle {
    return Edit_GetBorderStyle(e.instance)
}

func (e *TEdit) SetBorderStyle(value TBorderStyle) {
    Edit_SetBorderStyle(e.instance, value)
}

func (e *TEdit) Color() TColor {
    return Edit_GetColor(e.instance)
}

func (e *TEdit) SetColor(value TColor) {
    Edit_SetColor(e.instance, value)
}

func (e *TEdit) DoubleBuffered() bool {
    return Edit_GetDoubleBuffered(e.instance)
}

func (e *TEdit) SetDoubleBuffered(value bool) {
    Edit_SetDoubleBuffered(e.instance, value)
}

func (e *TEdit) DragCursor() TCursor {
    return Edit_GetDragCursor(e.instance)
}

func (e *TEdit) SetDragCursor(value TCursor) {
    Edit_SetDragCursor(e.instance, value)
}

func (e *TEdit) DragKind() TDragKind {
    return Edit_GetDragKind(e.instance)
}

func (e *TEdit) SetDragKind(value TDragKind) {
    Edit_SetDragKind(e.instance, value)
}

func (e *TEdit) DragMode() TDragMode {
    return Edit_GetDragMode(e.instance)
}

func (e *TEdit) SetDragMode(value TDragMode) {
    Edit_SetDragMode(e.instance, value)
}

func (e *TEdit) Enabled() bool {
    return Edit_GetEnabled(e.instance)
}

func (e *TEdit) SetEnabled(value bool) {
    Edit_SetEnabled(e.instance, value)
}

func (e *TEdit) Font() *TFont {
    return FontFromInst(Edit_GetFont(e.instance))
}

func (e *TEdit) SetFont(value *TFont) {
    Edit_SetFont(e.instance, CheckPtr(value))
}

func (e *TEdit) HideSelection() bool {
    return Edit_GetHideSelection(e.instance)
}

func (e *TEdit) SetHideSelection(value bool) {
    Edit_SetHideSelection(e.instance, value)
}

func (e *TEdit) MaxLength() int32 {
    return Edit_GetMaxLength(e.instance)
}

func (e *TEdit) SetMaxLength(value int32) {
    Edit_SetMaxLength(e.instance, value)
}

func (e *TEdit) NumbersOnly() bool {
    return Edit_GetNumbersOnly(e.instance)
}

func (e *TEdit) SetNumbersOnly(value bool) {
    Edit_SetNumbersOnly(e.instance, value)
}

func (e *TEdit) ParentColor() bool {
    return Edit_GetParentColor(e.instance)
}

func (e *TEdit) SetParentColor(value bool) {
    Edit_SetParentColor(e.instance, value)
}

func (e *TEdit) ParentCtl3D() bool {
    return Edit_GetParentCtl3D(e.instance)
}

func (e *TEdit) SetParentCtl3D(value bool) {
    Edit_SetParentCtl3D(e.instance, value)
}

func (e *TEdit) ParentDoubleBuffered() bool {
    return Edit_GetParentDoubleBuffered(e.instance)
}

func (e *TEdit) SetParentDoubleBuffered(value bool) {
    Edit_SetParentDoubleBuffered(e.instance, value)
}

func (e *TEdit) ParentFont() bool {
    return Edit_GetParentFont(e.instance)
}

func (e *TEdit) SetParentFont(value bool) {
    Edit_SetParentFont(e.instance, value)
}

func (e *TEdit) ParentShowHint() bool {
    return Edit_GetParentShowHint(e.instance)
}

func (e *TEdit) SetParentShowHint(value bool) {
    Edit_SetParentShowHint(e.instance, value)
}

func (e *TEdit) PasswordChar() uint16 {
    return Edit_GetPasswordChar(e.instance)
}

func (e *TEdit) SetPasswordChar(value uint16) {
    Edit_SetPasswordChar(e.instance, value)
}

func (e *TEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Edit_GetPopupMenu(e.instance))
}

func (e *TEdit) SetPopupMenu(value IComponent) {
    Edit_SetPopupMenu(e.instance, CheckPtr(value))
}

func (e *TEdit) ReadOnly() bool {
    return Edit_GetReadOnly(e.instance)
}

func (e *TEdit) SetReadOnly(value bool) {
    Edit_SetReadOnly(e.instance, value)
}

func (e *TEdit) ShowHint() bool {
    return Edit_GetShowHint(e.instance)
}

func (e *TEdit) SetShowHint(value bool) {
    Edit_SetShowHint(e.instance, value)
}

func (e *TEdit) TabOrder() uint16 {
    return Edit_GetTabOrder(e.instance)
}

func (e *TEdit) SetTabOrder(value uint16) {
    Edit_SetTabOrder(e.instance, value)
}

func (e *TEdit) TabStop() bool {
    return Edit_GetTabStop(e.instance)
}

func (e *TEdit) SetTabStop(value bool) {
    Edit_SetTabStop(e.instance, value)
}

func (e *TEdit) Text() string {
    return Edit_GetText(e.instance)
}

func (e *TEdit) SetText(value string) {
    Edit_SetText(e.instance, value)
}

func (e *TEdit) TextHint() string {
    return Edit_GetTextHint(e.instance)
}

func (e *TEdit) SetTextHint(value string) {
    Edit_SetTextHint(e.instance, value)
}

func (e *TEdit) Visible() bool {
    return Edit_GetVisible(e.instance)
}

func (e *TEdit) SetVisible(value bool) {
    Edit_SetVisible(e.instance, value)
}

func (e *TEdit) StyleElements() TStyleElements {
    return Edit_GetStyleElements(e.instance)
}

func (e *TEdit) SetStyleElements(value TStyleElements) {
    Edit_SetStyleElements(e.instance, value)
}

func (e *TEdit) SetOnChange(fn TNotifyEvent) {
    Edit_SetOnChange(e.instance, fn)
}

func (e *TEdit) SetOnClick(fn TNotifyEvent) {
    Edit_SetOnClick(e.instance, fn)
}

func (e *TEdit) SetOnContextPopup(fn TContextPopupEvent) {
    Edit_SetOnContextPopup(e.instance, fn)
}

func (e *TEdit) SetOnDblClick(fn TNotifyEvent) {
    Edit_SetOnDblClick(e.instance, fn)
}

func (e *TEdit) SetOnDragDrop(fn TDragDropEvent) {
    Edit_SetOnDragDrop(e.instance, fn)
}

func (e *TEdit) SetOnDragOver(fn TDragOverEvent) {
    Edit_SetOnDragOver(e.instance, fn)
}

func (e *TEdit) SetOnEndDock(fn TEndDragEvent) {
    Edit_SetOnEndDock(e.instance, fn)
}

func (e *TEdit) SetOnEndDrag(fn TEndDragEvent) {
    Edit_SetOnEndDrag(e.instance, fn)
}

func (e *TEdit) SetOnEnter(fn TNotifyEvent) {
    Edit_SetOnEnter(e.instance, fn)
}

func (e *TEdit) SetOnExit(fn TNotifyEvent) {
    Edit_SetOnExit(e.instance, fn)
}

func (e *TEdit) SetOnKeyDown(fn TKeyEvent) {
    Edit_SetOnKeyDown(e.instance, fn)
}

func (e *TEdit) SetOnKeyPress(fn TKeyPressEvent) {
    Edit_SetOnKeyPress(e.instance, fn)
}

func (e *TEdit) SetOnKeyUp(fn TKeyEvent) {
    Edit_SetOnKeyUp(e.instance, fn)
}

func (e *TEdit) SetOnMouseDown(fn TMouseEvent) {
    Edit_SetOnMouseDown(e.instance, fn)
}

func (e *TEdit) SetOnMouseEnter(fn TNotifyEvent) {
    Edit_SetOnMouseEnter(e.instance, fn)
}

func (e *TEdit) SetOnMouseLeave(fn TNotifyEvent) {
    Edit_SetOnMouseLeave(e.instance, fn)
}

func (e *TEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    Edit_SetOnMouseMove(e.instance, fn)
}

func (e *TEdit) SetOnMouseUp(fn TMouseEvent) {
    Edit_SetOnMouseUp(e.instance, fn)
}

func (e *TEdit) SetOnStartDock(fn TStartDockEvent) {
    Edit_SetOnStartDock(e.instance, fn)
}

func (e *TEdit) CanUndo() bool {
    return Edit_GetCanUndo(e.instance)
}

func (e *TEdit) Modified() bool {
    return Edit_GetModified(e.instance)
}

func (e *TEdit) SetModified(value bool) {
    Edit_SetModified(e.instance, value)
}

func (e *TEdit) SelLength() int32 {
    return Edit_GetSelLength(e.instance)
}

func (e *TEdit) SetSelLength(value int32) {
    Edit_SetSelLength(e.instance, value)
}

func (e *TEdit) SelStart() int32 {
    return Edit_GetSelStart(e.instance)
}

func (e *TEdit) SetSelStart(value int32) {
    Edit_SetSelStart(e.instance, value)
}

func (e *TEdit) SelText() string {
    return Edit_GetSelText(e.instance)
}

func (e *TEdit) SetSelText(value string) {
    Edit_SetSelText(e.instance, value)
}

func (e *TEdit) DockSite() bool {
    return Edit_GetDockSite(e.instance)
}

func (e *TEdit) SetDockSite(value bool) {
    Edit_SetDockSite(e.instance, value)
}

func (e *TEdit) Brush() *TBrush {
    return BrushFromInst(Edit_GetBrush(e.instance))
}

func (e *TEdit) ControlCount() int32 {
    return Edit_GetControlCount(e.instance)
}

func (e *TEdit) Handle() HWND {
    return Edit_GetHandle(e.instance)
}

func (e *TEdit) ParentWindow() HWND {
    return Edit_GetParentWindow(e.instance)
}

func (e *TEdit) SetParentWindow(value HWND) {
    Edit_SetParentWindow(e.instance, value)
}

func (e *TEdit) UseDockManager() bool {
    return Edit_GetUseDockManager(e.instance)
}

func (e *TEdit) SetUseDockManager(value bool) {
    Edit_SetUseDockManager(e.instance, value)
}

func (e *TEdit) Action() *TAction {
    return ActionFromInst(Edit_GetAction(e.instance))
}

func (e *TEdit) SetAction(value IComponent) {
    Edit_SetAction(e.instance, CheckPtr(value))
}

func (e *TEdit) BoundsRect() TRect {
    return Edit_GetBoundsRect(e.instance)
}

func (e *TEdit) SetBoundsRect(value TRect) {
    Edit_SetBoundsRect(e.instance, value)
}

func (e *TEdit) ClientHeight() int32 {
    return Edit_GetClientHeight(e.instance)
}

func (e *TEdit) SetClientHeight(value int32) {
    Edit_SetClientHeight(e.instance, value)
}

func (e *TEdit) ClientRect() TRect {
    return Edit_GetClientRect(e.instance)
}

func (e *TEdit) ClientWidth() int32 {
    return Edit_GetClientWidth(e.instance)
}

func (e *TEdit) SetClientWidth(value int32) {
    Edit_SetClientWidth(e.instance, value)
}

func (e *TEdit) ExplicitLeft() int32 {
    return Edit_GetExplicitLeft(e.instance)
}

func (e *TEdit) ExplicitTop() int32 {
    return Edit_GetExplicitTop(e.instance)
}

func (e *TEdit) ExplicitWidth() int32 {
    return Edit_GetExplicitWidth(e.instance)
}

func (e *TEdit) ExplicitHeight() int32 {
    return Edit_GetExplicitHeight(e.instance)
}

func (e *TEdit) Floating() bool {
    return Edit_GetFloating(e.instance)
}

func (e *TEdit) Parent() *TControl {
    return ControlFromInst(Edit_GetParent(e.instance))
}

func (e *TEdit) SetParent(value IControl) {
    Edit_SetParent(e.instance, CheckPtr(value))
}

func (e *TEdit) AlignWithMargins() bool {
    return Edit_GetAlignWithMargins(e.instance)
}

func (e *TEdit) SetAlignWithMargins(value bool) {
    Edit_SetAlignWithMargins(e.instance, value)
}

func (e *TEdit) Left() int32 {
    return Edit_GetLeft(e.instance)
}

func (e *TEdit) SetLeft(value int32) {
    Edit_SetLeft(e.instance, value)
}

func (e *TEdit) Top() int32 {
    return Edit_GetTop(e.instance)
}

func (e *TEdit) SetTop(value int32) {
    Edit_SetTop(e.instance, value)
}

func (e *TEdit) Width() int32 {
    return Edit_GetWidth(e.instance)
}

func (e *TEdit) SetWidth(value int32) {
    Edit_SetWidth(e.instance, value)
}

func (e *TEdit) Height() int32 {
    return Edit_GetHeight(e.instance)
}

func (e *TEdit) SetHeight(value int32) {
    Edit_SetHeight(e.instance, value)
}

func (e *TEdit) Cursor() TCursor {
    return Edit_GetCursor(e.instance)
}

func (e *TEdit) SetCursor(value TCursor) {
    Edit_SetCursor(e.instance, value)
}

func (e *TEdit) Hint() string {
    return Edit_GetHint(e.instance)
}

func (e *TEdit) SetHint(value string) {
    Edit_SetHint(e.instance, value)
}

func (e *TEdit) Margins() *TMargins {
    return MarginsFromInst(Edit_GetMargins(e.instance))
}

func (e *TEdit) SetMargins(value *TMargins) {
    Edit_SetMargins(e.instance, CheckPtr(value))
}

func (e *TEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(Edit_GetCustomHint(e.instance))
}

func (e *TEdit) SetCustomHint(value IComponent) {
    Edit_SetCustomHint(e.instance, CheckPtr(value))
}

func (e *TEdit) ComponentCount() int32 {
    return Edit_GetComponentCount(e.instance)
}

func (e *TEdit) ComponentIndex() int32 {
    return Edit_GetComponentIndex(e.instance)
}

func (e *TEdit) SetComponentIndex(value int32) {
    Edit_SetComponentIndex(e.instance, value)
}

func (e *TEdit) Owner() *TComponent {
    return ComponentFromInst(Edit_GetOwner(e.instance))
}

func (e *TEdit) Name() string {
    return Edit_GetName(e.instance)
}

func (e *TEdit) SetName(value string) {
    Edit_SetName(e.instance, value)
}

func (e *TEdit) Tag() int {
    return Edit_GetTag(e.instance)
}

func (e *TEdit) SetTag(value int) {
    Edit_SetTag(e.instance, value)
}

func (e *TEdit) Controls(Index int32) *TControl {
    return ControlFromInst(Edit_GetControls(e.instance, Index))
}

func (e *TEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Edit_GetComponents(e.instance, AIndex))
}

