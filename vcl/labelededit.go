
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

type TLabeledEdit struct {
    IWinControl
    instance uintptr
}

func NewLabeledEdit(owner IComponent) *TLabeledEdit {
    l := new(TLabeledEdit)
    l.instance = LabeledEdit_Create(CheckPtr(owner))
    return l
}

func LabeledEditFromInst(inst uintptr) *TLabeledEdit {
    l := new(TLabeledEdit)
    l.instance = inst
    return l
}

func LabeledEditFromObj(obj IObject) *TLabeledEdit {
    l := new(TLabeledEdit)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TLabeledEdit) Free() {
    if l.instance != 0 {
        LabeledEdit_Free(l.instance)
        l.instance = 0
    }
}

func (l *TLabeledEdit) Instance() uintptr {
    return l.instance
}

func (l *TLabeledEdit) IsValid() bool {
    return l.instance != 0
}

func TLabeledEditClass() TClass {
    return LabeledEdit_StaticClassType()
}

func (l *TLabeledEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    LabeledEdit_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

func (l *TLabeledEdit) SetupInternalLabel() {
    LabeledEdit_SetupInternalLabel(l.instance)
}

func (l *TLabeledEdit) Clear() {
    LabeledEdit_Clear(l.instance)
}

func (l *TLabeledEdit) ClearSelection() {
    LabeledEdit_ClearSelection(l.instance)
}

func (l *TLabeledEdit) CopyToClipboard() {
    LabeledEdit_CopyToClipboard(l.instance)
}

func (l *TLabeledEdit) CutToClipboard() {
    LabeledEdit_CutToClipboard(l.instance)
}

func (l *TLabeledEdit) PasteFromClipboard() {
    LabeledEdit_PasteFromClipboard(l.instance)
}

func (l *TLabeledEdit) Undo() {
    LabeledEdit_Undo(l.instance)
}

func (l *TLabeledEdit) ClearUndo() {
    LabeledEdit_ClearUndo(l.instance)
}

func (l *TLabeledEdit) SelectAll() {
    LabeledEdit_SelectAll(l.instance)
}

func (l *TLabeledEdit) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return LabeledEdit_GetSelTextBuf(l.instance, Buffer , BufSize)
}

func (l *TLabeledEdit) CanFocus() bool {
    return LabeledEdit_CanFocus(l.instance)
}

func (l *TLabeledEdit) FlipChildren(AllLevels bool) {
    LabeledEdit_FlipChildren(l.instance, AllLevels)
}

func (l *TLabeledEdit) Focused() bool {
    return LabeledEdit_Focused(l.instance)
}

func (l *TLabeledEdit) HandleAllocated() bool {
    return LabeledEdit_HandleAllocated(l.instance)
}

func (l *TLabeledEdit) Invalidate() {
    LabeledEdit_Invalidate(l.instance)
}

func (l *TLabeledEdit) Realign() {
    LabeledEdit_Realign(l.instance)
}

func (l *TLabeledEdit) Repaint() {
    LabeledEdit_Repaint(l.instance)
}

func (l *TLabeledEdit) ScaleBy(M int32, D int32) {
    LabeledEdit_ScaleBy(l.instance, M , D)
}

func (l *TLabeledEdit) SetFocus() {
    LabeledEdit_SetFocus(l.instance)
}

func (l *TLabeledEdit) Update() {
    LabeledEdit_Update(l.instance)
}

func (l *TLabeledEdit) BringToFront() {
    LabeledEdit_BringToFront(l.instance)
}

func (l *TLabeledEdit) ClientToScreen(Point TPoint) TPoint {
    return LabeledEdit_ClientToScreen(l.instance, Point)
}

func (l *TLabeledEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return LabeledEdit_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

func (l *TLabeledEdit) Dragging() bool {
    return LabeledEdit_Dragging(l.instance)
}

func (l *TLabeledEdit) HasParent() bool {
    return LabeledEdit_HasParent(l.instance)
}

func (l *TLabeledEdit) Hide() {
    LabeledEdit_Hide(l.instance)
}

func (l *TLabeledEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return LabeledEdit_Perform(l.instance, Msg , WParam , LParam)
}

func (l *TLabeledEdit) Refresh() {
    LabeledEdit_Refresh(l.instance)
}

func (l *TLabeledEdit) ScreenToClient(Point TPoint) TPoint {
    return LabeledEdit_ScreenToClient(l.instance, Point)
}

func (l *TLabeledEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return LabeledEdit_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

func (l *TLabeledEdit) SendToBack() {
    LabeledEdit_SendToBack(l.instance)
}

func (l *TLabeledEdit) Show() {
    LabeledEdit_Show(l.instance)
}

func (l *TLabeledEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    return LabeledEdit_GetTextBuf(l.instance, Buffer , BufSize)
}

func (l *TLabeledEdit) GetTextLen() int32 {
    return LabeledEdit_GetTextLen(l.instance)
}

func (l *TLabeledEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(LabeledEdit_FindComponent(l.instance, AName))
}

func (l *TLabeledEdit) GetNamePath() string {
    return LabeledEdit_GetNamePath(l.instance)
}

func (l *TLabeledEdit) Assign(Source IObject) {
    LabeledEdit_Assign(l.instance, CheckPtr(Source))
}

func (l *TLabeledEdit) DisposeOf() {
    LabeledEdit_DisposeOf(l.instance)
}

func (l *TLabeledEdit) ClassType() TClass {
    return LabeledEdit_ClassType(l.instance)
}

func (l *TLabeledEdit) ClassName() string {
    return LabeledEdit_ClassName(l.instance)
}

func (l *TLabeledEdit) InstanceSize() int32 {
    return LabeledEdit_InstanceSize(l.instance)
}

func (l *TLabeledEdit) InheritsFrom(AClass TClass) bool {
    return LabeledEdit_InheritsFrom(l.instance, AClass)
}

func (l *TLabeledEdit) Equals(Obj IObject) bool {
    return LabeledEdit_Equals(l.instance, CheckPtr(Obj))
}

func (l *TLabeledEdit) GetHashCode() int32 {
    return LabeledEdit_GetHashCode(l.instance)
}

func (l *TLabeledEdit) ToString() string {
    return LabeledEdit_ToString(l.instance)
}

func (l *TLabeledEdit) Alignment() TAlignment {
    return LabeledEdit_GetAlignment(l.instance)
}

func (l *TLabeledEdit) SetAlignment(value TAlignment) {
    LabeledEdit_SetAlignment(l.instance, value)
}

func (l *TLabeledEdit) Anchors() TAnchors {
    return LabeledEdit_GetAnchors(l.instance)
}

func (l *TLabeledEdit) SetAnchors(value TAnchors) {
    LabeledEdit_SetAnchors(l.instance, value)
}

func (l *TLabeledEdit) AutoSelect() bool {
    return LabeledEdit_GetAutoSelect(l.instance)
}

func (l *TLabeledEdit) SetAutoSelect(value bool) {
    LabeledEdit_SetAutoSelect(l.instance, value)
}

func (l *TLabeledEdit) AutoSize() bool {
    return LabeledEdit_GetAutoSize(l.instance)
}

func (l *TLabeledEdit) SetAutoSize(value bool) {
    LabeledEdit_SetAutoSize(l.instance, value)
}

func (l *TLabeledEdit) BevelEdges() TBevelEdges {
    return LabeledEdit_GetBevelEdges(l.instance)
}

func (l *TLabeledEdit) SetBevelEdges(value TBevelEdges) {
    LabeledEdit_SetBevelEdges(l.instance, value)
}

func (l *TLabeledEdit) BevelInner() TBevelCut {
    return LabeledEdit_GetBevelInner(l.instance)
}

func (l *TLabeledEdit) SetBevelInner(value TBevelCut) {
    LabeledEdit_SetBevelInner(l.instance, value)
}

func (l *TLabeledEdit) BevelKind() TBevelKind {
    return LabeledEdit_GetBevelKind(l.instance)
}

func (l *TLabeledEdit) SetBevelKind(value TBevelKind) {
    LabeledEdit_SetBevelKind(l.instance, value)
}

func (l *TLabeledEdit) BevelOuter() TBevelCut {
    return LabeledEdit_GetBevelOuter(l.instance)
}

func (l *TLabeledEdit) SetBevelOuter(value TBevelCut) {
    LabeledEdit_SetBevelOuter(l.instance, value)
}

func (l *TLabeledEdit) BiDiMode() TBiDiMode {
    return LabeledEdit_GetBiDiMode(l.instance)
}

func (l *TLabeledEdit) SetBiDiMode(value TBiDiMode) {
    LabeledEdit_SetBiDiMode(l.instance, value)
}

func (l *TLabeledEdit) BorderStyle() TBorderStyle {
    return LabeledEdit_GetBorderStyle(l.instance)
}

func (l *TLabeledEdit) SetBorderStyle(value TBorderStyle) {
    LabeledEdit_SetBorderStyle(l.instance, value)
}

func (l *TLabeledEdit) CharCase() TEditCharCase {
    return LabeledEdit_GetCharCase(l.instance)
}

func (l *TLabeledEdit) SetCharCase(value TEditCharCase) {
    LabeledEdit_SetCharCase(l.instance, value)
}

func (l *TLabeledEdit) Color() TColor {
    return LabeledEdit_GetColor(l.instance)
}

func (l *TLabeledEdit) SetColor(value TColor) {
    LabeledEdit_SetColor(l.instance, value)
}

func (l *TLabeledEdit) DoubleBuffered() bool {
    return LabeledEdit_GetDoubleBuffered(l.instance)
}

func (l *TLabeledEdit) SetDoubleBuffered(value bool) {
    LabeledEdit_SetDoubleBuffered(l.instance, value)
}

func (l *TLabeledEdit) DragCursor() TCursor {
    return LabeledEdit_GetDragCursor(l.instance)
}

func (l *TLabeledEdit) SetDragCursor(value TCursor) {
    LabeledEdit_SetDragCursor(l.instance, value)
}

func (l *TLabeledEdit) DragKind() TDragKind {
    return LabeledEdit_GetDragKind(l.instance)
}

func (l *TLabeledEdit) SetDragKind(value TDragKind) {
    LabeledEdit_SetDragKind(l.instance, value)
}

func (l *TLabeledEdit) DragMode() TDragMode {
    return LabeledEdit_GetDragMode(l.instance)
}

func (l *TLabeledEdit) SetDragMode(value TDragMode) {
    LabeledEdit_SetDragMode(l.instance, value)
}

func (l *TLabeledEdit) EditLabel() *TBoundLabel {
    return BoundLabelFromInst(LabeledEdit_GetEditLabel(l.instance))
}

func (l *TLabeledEdit) Enabled() bool {
    return LabeledEdit_GetEnabled(l.instance)
}

func (l *TLabeledEdit) SetEnabled(value bool) {
    LabeledEdit_SetEnabled(l.instance, value)
}

func (l *TLabeledEdit) Font() *TFont {
    return FontFromInst(LabeledEdit_GetFont(l.instance))
}

func (l *TLabeledEdit) SetFont(value *TFont) {
    LabeledEdit_SetFont(l.instance, CheckPtr(value))
}

func (l *TLabeledEdit) HideSelection() bool {
    return LabeledEdit_GetHideSelection(l.instance)
}

func (l *TLabeledEdit) SetHideSelection(value bool) {
    LabeledEdit_SetHideSelection(l.instance, value)
}

func (l *TLabeledEdit) LabelPosition() TLabelPosition {
    return LabeledEdit_GetLabelPosition(l.instance)
}

func (l *TLabeledEdit) SetLabelPosition(value TLabelPosition) {
    LabeledEdit_SetLabelPosition(l.instance, value)
}

func (l *TLabeledEdit) LabelSpacing() int32 {
    return LabeledEdit_GetLabelSpacing(l.instance)
}

func (l *TLabeledEdit) SetLabelSpacing(value int32) {
    LabeledEdit_SetLabelSpacing(l.instance, value)
}

func (l *TLabeledEdit) MaxLength() int32 {
    return LabeledEdit_GetMaxLength(l.instance)
}

func (l *TLabeledEdit) SetMaxLength(value int32) {
    LabeledEdit_SetMaxLength(l.instance, value)
}

func (l *TLabeledEdit) NumbersOnly() bool {
    return LabeledEdit_GetNumbersOnly(l.instance)
}

func (l *TLabeledEdit) SetNumbersOnly(value bool) {
    LabeledEdit_SetNumbersOnly(l.instance, value)
}

func (l *TLabeledEdit) ParentColor() bool {
    return LabeledEdit_GetParentColor(l.instance)
}

func (l *TLabeledEdit) SetParentColor(value bool) {
    LabeledEdit_SetParentColor(l.instance, value)
}

func (l *TLabeledEdit) ParentCtl3D() bool {
    return LabeledEdit_GetParentCtl3D(l.instance)
}

func (l *TLabeledEdit) SetParentCtl3D(value bool) {
    LabeledEdit_SetParentCtl3D(l.instance, value)
}

func (l *TLabeledEdit) ParentDoubleBuffered() bool {
    return LabeledEdit_GetParentDoubleBuffered(l.instance)
}

func (l *TLabeledEdit) SetParentDoubleBuffered(value bool) {
    LabeledEdit_SetParentDoubleBuffered(l.instance, value)
}

func (l *TLabeledEdit) ParentFont() bool {
    return LabeledEdit_GetParentFont(l.instance)
}

func (l *TLabeledEdit) SetParentFont(value bool) {
    LabeledEdit_SetParentFont(l.instance, value)
}

func (l *TLabeledEdit) ParentShowHint() bool {
    return LabeledEdit_GetParentShowHint(l.instance)
}

func (l *TLabeledEdit) SetParentShowHint(value bool) {
    LabeledEdit_SetParentShowHint(l.instance, value)
}

func (l *TLabeledEdit) PasswordChar() uint16 {
    return LabeledEdit_GetPasswordChar(l.instance)
}

func (l *TLabeledEdit) SetPasswordChar(value uint16) {
    LabeledEdit_SetPasswordChar(l.instance, value)
}

func (l *TLabeledEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(LabeledEdit_GetPopupMenu(l.instance))
}

func (l *TLabeledEdit) SetPopupMenu(value IComponent) {
    LabeledEdit_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TLabeledEdit) ReadOnly() bool {
    return LabeledEdit_GetReadOnly(l.instance)
}

func (l *TLabeledEdit) SetReadOnly(value bool) {
    LabeledEdit_SetReadOnly(l.instance, value)
}

func (l *TLabeledEdit) ShowHint() bool {
    return LabeledEdit_GetShowHint(l.instance)
}

func (l *TLabeledEdit) SetShowHint(value bool) {
    LabeledEdit_SetShowHint(l.instance, value)
}

func (l *TLabeledEdit) TabOrder() uint16 {
    return LabeledEdit_GetTabOrder(l.instance)
}

func (l *TLabeledEdit) SetTabOrder(value uint16) {
    LabeledEdit_SetTabOrder(l.instance, value)
}

func (l *TLabeledEdit) TabStop() bool {
    return LabeledEdit_GetTabStop(l.instance)
}

func (l *TLabeledEdit) SetTabStop(value bool) {
    LabeledEdit_SetTabStop(l.instance, value)
}

func (l *TLabeledEdit) Text() string {
    return LabeledEdit_GetText(l.instance)
}

func (l *TLabeledEdit) SetText(value string) {
    LabeledEdit_SetText(l.instance, value)
}

func (l *TLabeledEdit) TextHint() string {
    return LabeledEdit_GetTextHint(l.instance)
}

func (l *TLabeledEdit) SetTextHint(value string) {
    LabeledEdit_SetTextHint(l.instance, value)
}

func (l *TLabeledEdit) Visible() bool {
    return LabeledEdit_GetVisible(l.instance)
}

func (l *TLabeledEdit) SetVisible(value bool) {
    LabeledEdit_SetVisible(l.instance, value)
}

func (l *TLabeledEdit) StyleElements() TStyleElements {
    return LabeledEdit_GetStyleElements(l.instance)
}

func (l *TLabeledEdit) SetStyleElements(value TStyleElements) {
    LabeledEdit_SetStyleElements(l.instance, value)
}

func (l *TLabeledEdit) SetOnChange(fn TNotifyEvent) {
    LabeledEdit_SetOnChange(l.instance, fn)
}

func (l *TLabeledEdit) SetOnClick(fn TNotifyEvent) {
    LabeledEdit_SetOnClick(l.instance, fn)
}

func (l *TLabeledEdit) SetOnContextPopup(fn TContextPopupEvent) {
    LabeledEdit_SetOnContextPopup(l.instance, fn)
}

func (l *TLabeledEdit) SetOnDblClick(fn TNotifyEvent) {
    LabeledEdit_SetOnDblClick(l.instance, fn)
}

func (l *TLabeledEdit) SetOnDragDrop(fn TDragDropEvent) {
    LabeledEdit_SetOnDragDrop(l.instance, fn)
}

func (l *TLabeledEdit) SetOnDragOver(fn TDragOverEvent) {
    LabeledEdit_SetOnDragOver(l.instance, fn)
}

func (l *TLabeledEdit) SetOnEndDock(fn TEndDragEvent) {
    LabeledEdit_SetOnEndDock(l.instance, fn)
}

func (l *TLabeledEdit) SetOnEndDrag(fn TEndDragEvent) {
    LabeledEdit_SetOnEndDrag(l.instance, fn)
}

func (l *TLabeledEdit) SetOnEnter(fn TNotifyEvent) {
    LabeledEdit_SetOnEnter(l.instance, fn)
}

func (l *TLabeledEdit) SetOnExit(fn TNotifyEvent) {
    LabeledEdit_SetOnExit(l.instance, fn)
}

func (l *TLabeledEdit) SetOnKeyDown(fn TKeyEvent) {
    LabeledEdit_SetOnKeyDown(l.instance, fn)
}

func (l *TLabeledEdit) SetOnKeyPress(fn TKeyPressEvent) {
    LabeledEdit_SetOnKeyPress(l.instance, fn)
}

func (l *TLabeledEdit) SetOnKeyUp(fn TKeyEvent) {
    LabeledEdit_SetOnKeyUp(l.instance, fn)
}

func (l *TLabeledEdit) SetOnMouseDown(fn TMouseEvent) {
    LabeledEdit_SetOnMouseDown(l.instance, fn)
}

func (l *TLabeledEdit) SetOnMouseEnter(fn TNotifyEvent) {
    LabeledEdit_SetOnMouseEnter(l.instance, fn)
}

func (l *TLabeledEdit) SetOnMouseLeave(fn TNotifyEvent) {
    LabeledEdit_SetOnMouseLeave(l.instance, fn)
}

func (l *TLabeledEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    LabeledEdit_SetOnMouseMove(l.instance, fn)
}

func (l *TLabeledEdit) SetOnMouseUp(fn TMouseEvent) {
    LabeledEdit_SetOnMouseUp(l.instance, fn)
}

func (l *TLabeledEdit) SetOnStartDock(fn TStartDockEvent) {
    LabeledEdit_SetOnStartDock(l.instance, fn)
}

func (l *TLabeledEdit) CanUndo() bool {
    return LabeledEdit_GetCanUndo(l.instance)
}

func (l *TLabeledEdit) Modified() bool {
    return LabeledEdit_GetModified(l.instance)
}

func (l *TLabeledEdit) SetModified(value bool) {
    LabeledEdit_SetModified(l.instance, value)
}

func (l *TLabeledEdit) SelLength() int32 {
    return LabeledEdit_GetSelLength(l.instance)
}

func (l *TLabeledEdit) SetSelLength(value int32) {
    LabeledEdit_SetSelLength(l.instance, value)
}

func (l *TLabeledEdit) SelStart() int32 {
    return LabeledEdit_GetSelStart(l.instance)
}

func (l *TLabeledEdit) SetSelStart(value int32) {
    LabeledEdit_SetSelStart(l.instance, value)
}

func (l *TLabeledEdit) SelText() string {
    return LabeledEdit_GetSelText(l.instance)
}

func (l *TLabeledEdit) SetSelText(value string) {
    LabeledEdit_SetSelText(l.instance, value)
}

func (l *TLabeledEdit) DockSite() bool {
    return LabeledEdit_GetDockSite(l.instance)
}

func (l *TLabeledEdit) SetDockSite(value bool) {
    LabeledEdit_SetDockSite(l.instance, value)
}

func (l *TLabeledEdit) Brush() *TBrush {
    return BrushFromInst(LabeledEdit_GetBrush(l.instance))
}

func (l *TLabeledEdit) ControlCount() int32 {
    return LabeledEdit_GetControlCount(l.instance)
}

func (l *TLabeledEdit) Handle() HWND {
    return LabeledEdit_GetHandle(l.instance)
}

func (l *TLabeledEdit) ParentWindow() HWND {
    return LabeledEdit_GetParentWindow(l.instance)
}

func (l *TLabeledEdit) SetParentWindow(value HWND) {
    LabeledEdit_SetParentWindow(l.instance, value)
}

func (l *TLabeledEdit) UseDockManager() bool {
    return LabeledEdit_GetUseDockManager(l.instance)
}

func (l *TLabeledEdit) SetUseDockManager(value bool) {
    LabeledEdit_SetUseDockManager(l.instance, value)
}

func (l *TLabeledEdit) Action() *TAction {
    return ActionFromInst(LabeledEdit_GetAction(l.instance))
}

func (l *TLabeledEdit) SetAction(value IComponent) {
    LabeledEdit_SetAction(l.instance, CheckPtr(value))
}

func (l *TLabeledEdit) Align() TAlign {
    return LabeledEdit_GetAlign(l.instance)
}

func (l *TLabeledEdit) SetAlign(value TAlign) {
    LabeledEdit_SetAlign(l.instance, value)
}

func (l *TLabeledEdit) BoundsRect() TRect {
    return LabeledEdit_GetBoundsRect(l.instance)
}

func (l *TLabeledEdit) SetBoundsRect(value TRect) {
    LabeledEdit_SetBoundsRect(l.instance, value)
}

func (l *TLabeledEdit) ClientHeight() int32 {
    return LabeledEdit_GetClientHeight(l.instance)
}

func (l *TLabeledEdit) SetClientHeight(value int32) {
    LabeledEdit_SetClientHeight(l.instance, value)
}

func (l *TLabeledEdit) ClientRect() TRect {
    return LabeledEdit_GetClientRect(l.instance)
}

func (l *TLabeledEdit) ClientWidth() int32 {
    return LabeledEdit_GetClientWidth(l.instance)
}

func (l *TLabeledEdit) SetClientWidth(value int32) {
    LabeledEdit_SetClientWidth(l.instance, value)
}

func (l *TLabeledEdit) ExplicitLeft() int32 {
    return LabeledEdit_GetExplicitLeft(l.instance)
}

func (l *TLabeledEdit) ExplicitTop() int32 {
    return LabeledEdit_GetExplicitTop(l.instance)
}

func (l *TLabeledEdit) ExplicitWidth() int32 {
    return LabeledEdit_GetExplicitWidth(l.instance)
}

func (l *TLabeledEdit) ExplicitHeight() int32 {
    return LabeledEdit_GetExplicitHeight(l.instance)
}

func (l *TLabeledEdit) Floating() bool {
    return LabeledEdit_GetFloating(l.instance)
}

func (l *TLabeledEdit) Parent() *TWinControl {
    return WinControlFromInst(LabeledEdit_GetParent(l.instance))
}

func (l *TLabeledEdit) SetParent(value IWinControl) {
    LabeledEdit_SetParent(l.instance, CheckPtr(value))
}

func (l *TLabeledEdit) AlignWithMargins() bool {
    return LabeledEdit_GetAlignWithMargins(l.instance)
}

func (l *TLabeledEdit) SetAlignWithMargins(value bool) {
    LabeledEdit_SetAlignWithMargins(l.instance, value)
}

func (l *TLabeledEdit) Left() int32 {
    return LabeledEdit_GetLeft(l.instance)
}

func (l *TLabeledEdit) SetLeft(value int32) {
    LabeledEdit_SetLeft(l.instance, value)
}

func (l *TLabeledEdit) Top() int32 {
    return LabeledEdit_GetTop(l.instance)
}

func (l *TLabeledEdit) SetTop(value int32) {
    LabeledEdit_SetTop(l.instance, value)
}

func (l *TLabeledEdit) Width() int32 {
    return LabeledEdit_GetWidth(l.instance)
}

func (l *TLabeledEdit) SetWidth(value int32) {
    LabeledEdit_SetWidth(l.instance, value)
}

func (l *TLabeledEdit) Height() int32 {
    return LabeledEdit_GetHeight(l.instance)
}

func (l *TLabeledEdit) SetHeight(value int32) {
    LabeledEdit_SetHeight(l.instance, value)
}

func (l *TLabeledEdit) Cursor() TCursor {
    return LabeledEdit_GetCursor(l.instance)
}

func (l *TLabeledEdit) SetCursor(value TCursor) {
    LabeledEdit_SetCursor(l.instance, value)
}

func (l *TLabeledEdit) Hint() string {
    return LabeledEdit_GetHint(l.instance)
}

func (l *TLabeledEdit) SetHint(value string) {
    LabeledEdit_SetHint(l.instance, value)
}

func (l *TLabeledEdit) Margins() *TMargins {
    return MarginsFromInst(LabeledEdit_GetMargins(l.instance))
}

func (l *TLabeledEdit) SetMargins(value *TMargins) {
    LabeledEdit_SetMargins(l.instance, CheckPtr(value))
}

func (l *TLabeledEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(LabeledEdit_GetCustomHint(l.instance))
}

func (l *TLabeledEdit) SetCustomHint(value IComponent) {
    LabeledEdit_SetCustomHint(l.instance, CheckPtr(value))
}

func (l *TLabeledEdit) ComponentCount() int32 {
    return LabeledEdit_GetComponentCount(l.instance)
}

func (l *TLabeledEdit) ComponentIndex() int32 {
    return LabeledEdit_GetComponentIndex(l.instance)
}

func (l *TLabeledEdit) SetComponentIndex(value int32) {
    LabeledEdit_SetComponentIndex(l.instance, value)
}

func (l *TLabeledEdit) Owner() *TComponent {
    return ComponentFromInst(LabeledEdit_GetOwner(l.instance))
}

func (l *TLabeledEdit) Name() string {
    return LabeledEdit_GetName(l.instance)
}

func (l *TLabeledEdit) SetName(value string) {
    LabeledEdit_SetName(l.instance, value)
}

func (l *TLabeledEdit) Tag() int {
    return LabeledEdit_GetTag(l.instance)
}

func (l *TLabeledEdit) SetTag(value int) {
    LabeledEdit_SetTag(l.instance, value)
}

func (l *TLabeledEdit) Controls(Index int32) *TControl {
    return ControlFromInst(LabeledEdit_GetControls(l.instance, Index))
}

func (l *TLabeledEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(LabeledEdit_GetComponents(l.instance, AIndex))
}

