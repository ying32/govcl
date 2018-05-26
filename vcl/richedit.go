
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

type TRichEdit struct {
    IControl
    instance uintptr
}

func NewRichEdit(owner IComponent) *TRichEdit {
    r := new(TRichEdit)
    r.instance = RichEdit_Create(CheckPtr(owner))
    return r
}

func RichEditFromInst(inst uintptr) *TRichEdit {
    r := new(TRichEdit)
    r.instance = inst
    return r
}

func RichEditFromObj(obj IObject) *TRichEdit {
    r := new(TRichEdit)
    r.instance = CheckPtr(obj)
    return r
}

func (r *TRichEdit) Free() {
    if r.instance != 0 {
        RichEdit_Free(r.instance)
        r.instance = 0
    }
}

func (r *TRichEdit) Instance() uintptr {
    return r.instance
}

func (r *TRichEdit) IsValid() bool {
    return r.instance != 0
}

func (r *TRichEdit) Clear() {
    RichEdit_Clear(r.instance)
}

func (r *TRichEdit) FindText(SearchStr string, StartPos int32, Length int32, Options TSearchTypes) int32 {
    return RichEdit_FindText(r.instance, SearchStr , StartPos , Length , Options)
}

func (r *TRichEdit) Print(Caption string) {
    RichEdit_Print(r.instance, Caption)
}

func (r *TRichEdit) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return RichEdit_GetSelTextBuf(r.instance, Buffer , BufSize)
}

func (r *TRichEdit) ClearSelection() {
    RichEdit_ClearSelection(r.instance)
}

func (r *TRichEdit) CopyToClipboard() {
    RichEdit_CopyToClipboard(r.instance)
}

func (r *TRichEdit) CutToClipboard() {
    RichEdit_CutToClipboard(r.instance)
}

func (r *TRichEdit) PasteFromClipboard() {
    RichEdit_PasteFromClipboard(r.instance)
}

func (r *TRichEdit) Undo() {
    RichEdit_Undo(r.instance)
}

func (r *TRichEdit) ClearUndo() {
    RichEdit_ClearUndo(r.instance)
}

func (r *TRichEdit) SelectAll() {
    RichEdit_SelectAll(r.instance)
}

func (r *TRichEdit) CanFocus() bool {
    return RichEdit_CanFocus(r.instance)
}

func (r *TRichEdit) FlipChildren(AllLevels bool) {
    RichEdit_FlipChildren(r.instance, AllLevels)
}

func (r *TRichEdit) Focused() bool {
    return RichEdit_Focused(r.instance)
}

func (r *TRichEdit) HandleAllocated() bool {
    return RichEdit_HandleAllocated(r.instance)
}

func (r *TRichEdit) Invalidate() {
    RichEdit_Invalidate(r.instance)
}

func (r *TRichEdit) Realign() {
    RichEdit_Realign(r.instance)
}

func (r *TRichEdit) Repaint() {
    RichEdit_Repaint(r.instance)
}

func (r *TRichEdit) ScaleBy(M int32, D int32) {
    RichEdit_ScaleBy(r.instance, M , D)
}

func (r *TRichEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RichEdit_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

func (r *TRichEdit) SetFocus() {
    RichEdit_SetFocus(r.instance)
}

func (r *TRichEdit) Update() {
    RichEdit_Update(r.instance)
}

func (r *TRichEdit) BringToFront() {
    RichEdit_BringToFront(r.instance)
}

func (r *TRichEdit) HasParent() bool {
    return RichEdit_HasParent(r.instance)
}

func (r *TRichEdit) Hide() {
    RichEdit_Hide(r.instance)
}

func (r *TRichEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RichEdit_Perform(r.instance, Msg , WParam , LParam)
}

func (r *TRichEdit) Refresh() {
    RichEdit_Refresh(r.instance)
}

func (r *TRichEdit) SendToBack() {
    RichEdit_SendToBack(r.instance)
}

func (r *TRichEdit) Show() {
    RichEdit_Show(r.instance)
}

func (r *TRichEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    return RichEdit_GetTextBuf(r.instance, Buffer , BufSize)
}

func (r *TRichEdit) GetTextLen() int32 {
    return RichEdit_GetTextLen(r.instance)
}

func (r *TRichEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(RichEdit_FindComponent(r.instance, AName))
}

func (r *TRichEdit) GetNamePath() string {
    return RichEdit_GetNamePath(r.instance)
}

func (r *TRichEdit) Assign(Source IObject) {
    RichEdit_Assign(r.instance, CheckPtr(Source))
}

func (r *TRichEdit) ClassName() string {
    return RichEdit_ClassName(r.instance)
}

func (r *TRichEdit) Equals(Obj IObject) bool {
    return RichEdit_Equals(r.instance, CheckPtr(Obj))
}

func (r *TRichEdit) GetHashCode() int32 {
    return RichEdit_GetHashCode(r.instance)
}

func (r *TRichEdit) ToString() string {
    return RichEdit_ToString(r.instance)
}

func (r *TRichEdit) Align() TAlign {
    return RichEdit_GetAlign(r.instance)
}

func (r *TRichEdit) SetAlign(value TAlign) {
    RichEdit_SetAlign(r.instance, value)
}

func (r *TRichEdit) Alignment() TAlignment {
    return RichEdit_GetAlignment(r.instance)
}

func (r *TRichEdit) SetAlignment(value TAlignment) {
    RichEdit_SetAlignment(r.instance, value)
}

func (r *TRichEdit) Anchors() TAnchors {
    return RichEdit_GetAnchors(r.instance)
}

func (r *TRichEdit) SetAnchors(value TAnchors) {
    RichEdit_SetAnchors(r.instance, value)
}

func (r *TRichEdit) BevelEdges() TBevelEdges {
    return RichEdit_GetBevelEdges(r.instance)
}

func (r *TRichEdit) SetBevelEdges(value TBevelEdges) {
    RichEdit_SetBevelEdges(r.instance, value)
}

func (r *TRichEdit) BevelInner() TBevelCut {
    return RichEdit_GetBevelInner(r.instance)
}

func (r *TRichEdit) SetBevelInner(value TBevelCut) {
    RichEdit_SetBevelInner(r.instance, value)
}

func (r *TRichEdit) BevelOuter() TBevelCut {
    return RichEdit_GetBevelOuter(r.instance)
}

func (r *TRichEdit) SetBevelOuter(value TBevelCut) {
    RichEdit_SetBevelOuter(r.instance, value)
}

func (r *TRichEdit) BevelKind() TBevelKind {
    return RichEdit_GetBevelKind(r.instance)
}

func (r *TRichEdit) SetBevelKind(value TBevelKind) {
    RichEdit_SetBevelKind(r.instance, value)
}

func (r *TRichEdit) BiDiMode() TBiDiMode {
    return RichEdit_GetBiDiMode(r.instance)
}

func (r *TRichEdit) SetBiDiMode(value TBiDiMode) {
    RichEdit_SetBiDiMode(r.instance, value)
}

func (r *TRichEdit) BorderStyle() TBorderStyle {
    return RichEdit_GetBorderStyle(r.instance)
}

func (r *TRichEdit) SetBorderStyle(value TBorderStyle) {
    RichEdit_SetBorderStyle(r.instance, value)
}

func (r *TRichEdit) BorderWidth() int32 {
    return RichEdit_GetBorderWidth(r.instance)
}

func (r *TRichEdit) SetBorderWidth(value int32) {
    RichEdit_SetBorderWidth(r.instance, value)
}

func (r *TRichEdit) Color() TColor {
    return RichEdit_GetColor(r.instance)
}

func (r *TRichEdit) SetColor(value TColor) {
    RichEdit_SetColor(r.instance, value)
}

func (r *TRichEdit) Enabled() bool {
    return RichEdit_GetEnabled(r.instance)
}

func (r *TRichEdit) SetEnabled(value bool) {
    RichEdit_SetEnabled(r.instance, value)
}

func (r *TRichEdit) Font() *TFont {
    return FontFromInst(RichEdit_GetFont(r.instance))
}

func (r *TRichEdit) SetFont(value *TFont) {
    RichEdit_SetFont(r.instance, CheckPtr(value))
}

func (r *TRichEdit) HideSelection() bool {
    return RichEdit_GetHideSelection(r.instance)
}

func (r *TRichEdit) SetHideSelection(value bool) {
    RichEdit_SetHideSelection(r.instance, value)
}

func (r *TRichEdit) HideScrollBars() bool {
    return RichEdit_GetHideScrollBars(r.instance)
}

func (r *TRichEdit) SetHideScrollBars(value bool) {
    RichEdit_SetHideScrollBars(r.instance, value)
}

func (r *TRichEdit) Lines() *TStrings {
    return StringsFromInst(RichEdit_GetLines(r.instance))
}

func (r *TRichEdit) SetLines(value IObject) {
    RichEdit_SetLines(r.instance, CheckPtr(value))
}

func (r *TRichEdit) MaxLength() int32 {
    return RichEdit_GetMaxLength(r.instance)
}

func (r *TRichEdit) SetMaxLength(value int32) {
    RichEdit_SetMaxLength(r.instance, value)
}

func (r *TRichEdit) ParentColor() bool {
    return RichEdit_GetParentColor(r.instance)
}

func (r *TRichEdit) SetParentColor(value bool) {
    RichEdit_SetParentColor(r.instance, value)
}

func (r *TRichEdit) ParentCtl3D() bool {
    return RichEdit_GetParentCtl3D(r.instance)
}

func (r *TRichEdit) SetParentCtl3D(value bool) {
    RichEdit_SetParentCtl3D(r.instance, value)
}

func (r *TRichEdit) ParentFont() bool {
    return RichEdit_GetParentFont(r.instance)
}

func (r *TRichEdit) SetParentFont(value bool) {
    RichEdit_SetParentFont(r.instance, value)
}

func (r *TRichEdit) ParentShowHint() bool {
    return RichEdit_GetParentShowHint(r.instance)
}

func (r *TRichEdit) SetParentShowHint(value bool) {
    RichEdit_SetParentShowHint(r.instance, value)
}

func (r *TRichEdit) PlainText() bool {
    return RichEdit_GetPlainText(r.instance)
}

func (r *TRichEdit) SetPlainText(value bool) {
    RichEdit_SetPlainText(r.instance, value)
}

func (r *TRichEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(RichEdit_GetPopupMenu(r.instance))
}

func (r *TRichEdit) SetPopupMenu(value IComponent) {
    RichEdit_SetPopupMenu(r.instance, CheckPtr(value))
}

func (r *TRichEdit) ReadOnly() bool {
    return RichEdit_GetReadOnly(r.instance)
}

func (r *TRichEdit) SetReadOnly(value bool) {
    RichEdit_SetReadOnly(r.instance, value)
}

func (r *TRichEdit) ScrollBars() TScrollStyle {
    return RichEdit_GetScrollBars(r.instance)
}

func (r *TRichEdit) SetScrollBars(value TScrollStyle) {
    RichEdit_SetScrollBars(r.instance, value)
}

func (r *TRichEdit) ShowHint() bool {
    return RichEdit_GetShowHint(r.instance)
}

func (r *TRichEdit) SetShowHint(value bool) {
    RichEdit_SetShowHint(r.instance, value)
}

func (r *TRichEdit) TabOrder() uint16 {
    return RichEdit_GetTabOrder(r.instance)
}

func (r *TRichEdit) SetTabOrder(value uint16) {
    RichEdit_SetTabOrder(r.instance, value)
}

func (r *TRichEdit) TabStop() bool {
    return RichEdit_GetTabStop(r.instance)
}

func (r *TRichEdit) SetTabStop(value bool) {
    RichEdit_SetTabStop(r.instance, value)
}

func (r *TRichEdit) Visible() bool {
    return RichEdit_GetVisible(r.instance)
}

func (r *TRichEdit) SetVisible(value bool) {
    RichEdit_SetVisible(r.instance, value)
}

func (r *TRichEdit) WantTabs() bool {
    return RichEdit_GetWantTabs(r.instance)
}

func (r *TRichEdit) SetWantTabs(value bool) {
    RichEdit_SetWantTabs(r.instance, value)
}

func (r *TRichEdit) WantReturns() bool {
    return RichEdit_GetWantReturns(r.instance)
}

func (r *TRichEdit) SetWantReturns(value bool) {
    RichEdit_SetWantReturns(r.instance, value)
}

func (r *TRichEdit) WordWrap() bool {
    return RichEdit_GetWordWrap(r.instance)
}

func (r *TRichEdit) SetWordWrap(value bool) {
    RichEdit_SetWordWrap(r.instance, value)
}

func (r *TRichEdit) StyleElements() TStyleElements {
    return RichEdit_GetStyleElements(r.instance)
}

func (r *TRichEdit) SetStyleElements(value TStyleElements) {
    RichEdit_SetStyleElements(r.instance, value)
}

func (r *TRichEdit) Zoom() int32 {
    return RichEdit_GetZoom(r.instance)
}

func (r *TRichEdit) SetZoom(value int32) {
    RichEdit_SetZoom(r.instance, value)
}

func (r *TRichEdit) SetOnChange(fn TNotifyEvent) {
    RichEdit_SetOnChange(r.instance, fn)
}

func (r *TRichEdit) SetOnClick(fn TNotifyEvent) {
    RichEdit_SetOnClick(r.instance, fn)
}

func (r *TRichEdit) SetOnDblClick(fn TNotifyEvent) {
    RichEdit_SetOnDblClick(r.instance, fn)
}

func (r *TRichEdit) SetOnEnter(fn TNotifyEvent) {
    RichEdit_SetOnEnter(r.instance, fn)
}

func (r *TRichEdit) SetOnExit(fn TNotifyEvent) {
    RichEdit_SetOnExit(r.instance, fn)
}

func (r *TRichEdit) SetOnKeyDown(fn TKeyEvent) {
    RichEdit_SetOnKeyDown(r.instance, fn)
}

func (r *TRichEdit) SetOnKeyPress(fn TKeyPressEvent) {
    RichEdit_SetOnKeyPress(r.instance, fn)
}

func (r *TRichEdit) SetOnKeyUp(fn TKeyEvent) {
    RichEdit_SetOnKeyUp(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseDown(fn TMouseEvent) {
    RichEdit_SetOnMouseDown(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseEnter(fn TNotifyEvent) {
    RichEdit_SetOnMouseEnter(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseLeave(fn TNotifyEvent) {
    RichEdit_SetOnMouseLeave(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    RichEdit_SetOnMouseMove(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseUp(fn TMouseEvent) {
    RichEdit_SetOnMouseUp(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
    RichEdit_SetOnMouseWheel(r.instance, fn)
}

func (r *TRichEdit) ActiveLineNo() uint32 {
    return RichEdit_GetActiveLineNo(r.instance)
}

func (r *TRichEdit) DefAttributes() *TTextAttributes {
    return TextAttributesFromInst(RichEdit_GetDefAttributes(r.instance))
}

func (r *TRichEdit) SetDefAttributes(value *TTextAttributes) {
    RichEdit_SetDefAttributes(r.instance, CheckPtr(value))
}

func (r *TRichEdit) SelAttributes() *TTextAttributes {
    return TextAttributesFromInst(RichEdit_GetSelAttributes(r.instance))
}

func (r *TRichEdit) SetSelAttributes(value *TTextAttributes) {
    RichEdit_SetSelAttributes(r.instance, CheckPtr(value))
}

func (r *TRichEdit) PageRect() TRect {
    return RichEdit_GetPageRect(r.instance)
}

func (r *TRichEdit) SetPageRect(value TRect) {
    RichEdit_SetPageRect(r.instance, value)
}

func (r *TRichEdit) Paragraph() *TParaAttributes {
    return ParaAttributesFromInst(RichEdit_GetParagraph(r.instance))
}

func (r *TRichEdit) CaretPos() TPoint {
    return RichEdit_GetCaretPos(r.instance)
}

func (r *TRichEdit) SetCaretPos(value TPoint) {
    RichEdit_SetCaretPos(r.instance, value)
}

func (r *TRichEdit) CanUndo() bool {
    return RichEdit_GetCanUndo(r.instance)
}

func (r *TRichEdit) Modified() bool {
    return RichEdit_GetModified(r.instance)
}

func (r *TRichEdit) SetModified(value bool) {
    RichEdit_SetModified(r.instance, value)
}

func (r *TRichEdit) SelLength() int32 {
    return RichEdit_GetSelLength(r.instance)
}

func (r *TRichEdit) SetSelLength(value int32) {
    RichEdit_SetSelLength(r.instance, value)
}

func (r *TRichEdit) SelStart() int32 {
    return RichEdit_GetSelStart(r.instance)
}

func (r *TRichEdit) SetSelStart(value int32) {
    RichEdit_SetSelStart(r.instance, value)
}

func (r *TRichEdit) SelText() string {
    return RichEdit_GetSelText(r.instance)
}

func (r *TRichEdit) SetSelText(value string) {
    RichEdit_SetSelText(r.instance, value)
}

func (r *TRichEdit) Text() string {
    return RichEdit_GetText(r.instance)
}

func (r *TRichEdit) SetText(value string) {
    RichEdit_SetText(r.instance, value)
}

func (r *TRichEdit) TextHint() string {
    return RichEdit_GetTextHint(r.instance)
}

func (r *TRichEdit) SetTextHint(value string) {
    RichEdit_SetTextHint(r.instance, value)
}

func (r *TRichEdit) DoubleBuffered() bool {
    return RichEdit_GetDoubleBuffered(r.instance)
}

func (r *TRichEdit) SetDoubleBuffered(value bool) {
    RichEdit_SetDoubleBuffered(r.instance, value)
}

func (r *TRichEdit) Brush() *TBrush {
    return BrushFromInst(RichEdit_GetBrush(r.instance))
}

func (r *TRichEdit) ControlCount() int32 {
    return RichEdit_GetControlCount(r.instance)
}

func (r *TRichEdit) Handle() HWND {
    return RichEdit_GetHandle(r.instance)
}

func (r *TRichEdit) ParentDoubleBuffered() bool {
    return RichEdit_GetParentDoubleBuffered(r.instance)
}

func (r *TRichEdit) SetParentDoubleBuffered(value bool) {
    RichEdit_SetParentDoubleBuffered(r.instance, value)
}

func (r *TRichEdit) ParentWindow() HWND {
    return RichEdit_GetParentWindow(r.instance)
}

func (r *TRichEdit) SetParentWindow(value HWND) {
    RichEdit_SetParentWindow(r.instance, value)
}

func (r *TRichEdit) Action() *TAction {
    return ActionFromInst(RichEdit_GetAction(r.instance))
}

func (r *TRichEdit) SetAction(value IComponent) {
    RichEdit_SetAction(r.instance, CheckPtr(value))
}

func (r *TRichEdit) BoundsRect() TRect {
    return RichEdit_GetBoundsRect(r.instance)
}

func (r *TRichEdit) SetBoundsRect(value TRect) {
    RichEdit_SetBoundsRect(r.instance, value)
}

func (r *TRichEdit) ClientHeight() int32 {
    return RichEdit_GetClientHeight(r.instance)
}

func (r *TRichEdit) SetClientHeight(value int32) {
    RichEdit_SetClientHeight(r.instance, value)
}

func (r *TRichEdit) ClientRect() TRect {
    return RichEdit_GetClientRect(r.instance)
}

func (r *TRichEdit) ClientWidth() int32 {
    return RichEdit_GetClientWidth(r.instance)
}

func (r *TRichEdit) SetClientWidth(value int32) {
    RichEdit_SetClientWidth(r.instance, value)
}

func (r *TRichEdit) ExplicitLeft() int32 {
    return RichEdit_GetExplicitLeft(r.instance)
}

func (r *TRichEdit) ExplicitTop() int32 {
    return RichEdit_GetExplicitTop(r.instance)
}

func (r *TRichEdit) ExplicitWidth() int32 {
    return RichEdit_GetExplicitWidth(r.instance)
}

func (r *TRichEdit) ExplicitHeight() int32 {
    return RichEdit_GetExplicitHeight(r.instance)
}

func (r *TRichEdit) Parent() *TControl {
    return ControlFromInst(RichEdit_GetParent(r.instance))
}

func (r *TRichEdit) SetParent(value IControl) {
    RichEdit_SetParent(r.instance, CheckPtr(value))
}

func (r *TRichEdit) AlignWithMargins() bool {
    return RichEdit_GetAlignWithMargins(r.instance)
}

func (r *TRichEdit) SetAlignWithMargins(value bool) {
    RichEdit_SetAlignWithMargins(r.instance, value)
}

func (r *TRichEdit) Left() int32 {
    return RichEdit_GetLeft(r.instance)
}

func (r *TRichEdit) SetLeft(value int32) {
    RichEdit_SetLeft(r.instance, value)
}

func (r *TRichEdit) Top() int32 {
    return RichEdit_GetTop(r.instance)
}

func (r *TRichEdit) SetTop(value int32) {
    RichEdit_SetTop(r.instance, value)
}

func (r *TRichEdit) Width() int32 {
    return RichEdit_GetWidth(r.instance)
}

func (r *TRichEdit) SetWidth(value int32) {
    RichEdit_SetWidth(r.instance, value)
}

func (r *TRichEdit) Height() int32 {
    return RichEdit_GetHeight(r.instance)
}

func (r *TRichEdit) SetHeight(value int32) {
    RichEdit_SetHeight(r.instance, value)
}

func (r *TRichEdit) Cursor() TCursor {
    return RichEdit_GetCursor(r.instance)
}

func (r *TRichEdit) SetCursor(value TCursor) {
    RichEdit_SetCursor(r.instance, value)
}

func (r *TRichEdit) Hint() string {
    return RichEdit_GetHint(r.instance)
}

func (r *TRichEdit) SetHint(value string) {
    RichEdit_SetHint(r.instance, value)
}

func (r *TRichEdit) Margins() *TMargins {
    return MarginsFromInst(RichEdit_GetMargins(r.instance))
}

func (r *TRichEdit) SetMargins(value *TMargins) {
    RichEdit_SetMargins(r.instance, CheckPtr(value))
}

func (r *TRichEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(RichEdit_GetCustomHint(r.instance))
}

func (r *TRichEdit) SetCustomHint(value IComponent) {
    RichEdit_SetCustomHint(r.instance, CheckPtr(value))
}

func (r *TRichEdit) ComponentCount() int32 {
    return RichEdit_GetComponentCount(r.instance)
}

func (r *TRichEdit) ComponentIndex() int32 {
    return RichEdit_GetComponentIndex(r.instance)
}

func (r *TRichEdit) SetComponentIndex(value int32) {
    RichEdit_SetComponentIndex(r.instance, value)
}

func (r *TRichEdit) Owner() *TComponent {
    return ComponentFromInst(RichEdit_GetOwner(r.instance))
}

func (r *TRichEdit) Name() string {
    return RichEdit_GetName(r.instance)
}

func (r *TRichEdit) SetName(value string) {
    RichEdit_SetName(r.instance, value)
}

func (r *TRichEdit) Tag() int {
    return RichEdit_GetTag(r.instance)
}

func (r *TRichEdit) SetTag(value int) {
    RichEdit_SetTag(r.instance, value)
}

func (r *TRichEdit) Controls(Index int32) *TControl {
    return ControlFromInst(RichEdit_GetControls(r.instance, Index))
}

func (r *TRichEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(RichEdit_GetComponents(r.instance, AIndex))
}

