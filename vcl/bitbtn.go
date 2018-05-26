
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

type TBitBtn struct {
    IControl
    instance uintptr
}

func NewBitBtn(owner IComponent) *TBitBtn {
    b := new(TBitBtn)
    b.instance = BitBtn_Create(CheckPtr(owner))
    return b
}

func BitBtnFromInst(inst uintptr) *TBitBtn {
    b := new(TBitBtn)
    b.instance = inst
    return b
}

func BitBtnFromObj(obj IObject) *TBitBtn {
    b := new(TBitBtn)
    b.instance = CheckPtr(obj)
    return b
}

func (b *TBitBtn) Free() {
    if b.instance != 0 {
        BitBtn_Free(b.instance)
        b.instance = 0
    }
}

func (b *TBitBtn) Instance() uintptr {
    return b.instance
}

func (b *TBitBtn) IsValid() bool {
    return b.instance != 0
}

func (b *TBitBtn) Click() {
    BitBtn_Click(b.instance)
}

func (b *TBitBtn) CanFocus() bool {
    return BitBtn_CanFocus(b.instance)
}

func (b *TBitBtn) FlipChildren(AllLevels bool) {
    BitBtn_FlipChildren(b.instance, AllLevels)
}

func (b *TBitBtn) Focused() bool {
    return BitBtn_Focused(b.instance)
}

func (b *TBitBtn) HandleAllocated() bool {
    return BitBtn_HandleAllocated(b.instance)
}

func (b *TBitBtn) Invalidate() {
    BitBtn_Invalidate(b.instance)
}

func (b *TBitBtn) Realign() {
    BitBtn_Realign(b.instance)
}

func (b *TBitBtn) Repaint() {
    BitBtn_Repaint(b.instance)
}

func (b *TBitBtn) ScaleBy(M int32, D int32) {
    BitBtn_ScaleBy(b.instance, M , D)
}

func (b *TBitBtn) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    BitBtn_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

func (b *TBitBtn) SetFocus() {
    BitBtn_SetFocus(b.instance)
}

func (b *TBitBtn) Update() {
    BitBtn_Update(b.instance)
}

func (b *TBitBtn) BringToFront() {
    BitBtn_BringToFront(b.instance)
}

func (b *TBitBtn) HasParent() bool {
    return BitBtn_HasParent(b.instance)
}

func (b *TBitBtn) Hide() {
    BitBtn_Hide(b.instance)
}

func (b *TBitBtn) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return BitBtn_Perform(b.instance, Msg , WParam , LParam)
}

func (b *TBitBtn) Refresh() {
    BitBtn_Refresh(b.instance)
}

func (b *TBitBtn) SendToBack() {
    BitBtn_SendToBack(b.instance)
}

func (b *TBitBtn) Show() {
    BitBtn_Show(b.instance)
}

func (b *TBitBtn) GetTextBuf(Buffer string, BufSize int32) int32 {
    return BitBtn_GetTextBuf(b.instance, Buffer , BufSize)
}

func (b *TBitBtn) GetTextLen() int32 {
    return BitBtn_GetTextLen(b.instance)
}

func (b *TBitBtn) FindComponent(AName string) *TComponent {
    return ComponentFromInst(BitBtn_FindComponent(b.instance, AName))
}

func (b *TBitBtn) GetNamePath() string {
    return BitBtn_GetNamePath(b.instance)
}

func (b *TBitBtn) Assign(Source IObject) {
    BitBtn_Assign(b.instance, CheckPtr(Source))
}

func (b *TBitBtn) ClassName() string {
    return BitBtn_ClassName(b.instance)
}

func (b *TBitBtn) Equals(Obj IObject) bool {
    return BitBtn_Equals(b.instance, CheckPtr(Obj))
}

func (b *TBitBtn) GetHashCode() int32 {
    return BitBtn_GetHashCode(b.instance)
}

func (b *TBitBtn) ToString() string {
    return BitBtn_ToString(b.instance)
}

func (b *TBitBtn) Action() *TAction {
    return ActionFromInst(BitBtn_GetAction(b.instance))
}

func (b *TBitBtn) SetAction(value IComponent) {
    BitBtn_SetAction(b.instance, CheckPtr(value))
}

func (b *TBitBtn) Align() TAlign {
    return BitBtn_GetAlign(b.instance)
}

func (b *TBitBtn) SetAlign(value TAlign) {
    BitBtn_SetAlign(b.instance, value)
}

func (b *TBitBtn) Anchors() TAnchors {
    return BitBtn_GetAnchors(b.instance)
}

func (b *TBitBtn) SetAnchors(value TAnchors) {
    BitBtn_SetAnchors(b.instance, value)
}

func (b *TBitBtn) BiDiMode() TBiDiMode {
    return BitBtn_GetBiDiMode(b.instance)
}

func (b *TBitBtn) SetBiDiMode(value TBiDiMode) {
    BitBtn_SetBiDiMode(b.instance, value)
}

func (b *TBitBtn) Cancel() bool {
    return BitBtn_GetCancel(b.instance)
}

func (b *TBitBtn) SetCancel(value bool) {
    BitBtn_SetCancel(b.instance, value)
}

func (b *TBitBtn) Caption() string {
    return BitBtn_GetCaption(b.instance)
}

func (b *TBitBtn) SetCaption(value string) {
    BitBtn_SetCaption(b.instance, value)
}

func (b *TBitBtn) Default() bool {
    return BitBtn_GetDefault(b.instance)
}

func (b *TBitBtn) SetDefault(value bool) {
    BitBtn_SetDefault(b.instance, value)
}

func (b *TBitBtn) DoubleBuffered() bool {
    return BitBtn_GetDoubleBuffered(b.instance)
}

func (b *TBitBtn) SetDoubleBuffered(value bool) {
    BitBtn_SetDoubleBuffered(b.instance, value)
}

func (b *TBitBtn) Enabled() bool {
    return BitBtn_GetEnabled(b.instance)
}

func (b *TBitBtn) SetEnabled(value bool) {
    BitBtn_SetEnabled(b.instance, value)
}

func (b *TBitBtn) Font() *TFont {
    return FontFromInst(BitBtn_GetFont(b.instance))
}

func (b *TBitBtn) SetFont(value *TFont) {
    BitBtn_SetFont(b.instance, CheckPtr(value))
}

func (b *TBitBtn) Glyph() *TBitmap {
    return BitmapFromInst(BitBtn_GetGlyph(b.instance))
}

func (b *TBitBtn) SetGlyph(value *TBitmap) {
    BitBtn_SetGlyph(b.instance, CheckPtr(value))
}

func (b *TBitBtn) Kind() TBitBtnKind {
    return BitBtn_GetKind(b.instance)
}

func (b *TBitBtn) SetKind(value TBitBtnKind) {
    BitBtn_SetKind(b.instance, value)
}

func (b *TBitBtn) Layout() TButtonLayout {
    return BitBtn_GetLayout(b.instance)
}

func (b *TBitBtn) SetLayout(value TButtonLayout) {
    BitBtn_SetLayout(b.instance, value)
}

func (b *TBitBtn) ModalResult() TModalResult {
    return BitBtn_GetModalResult(b.instance)
}

func (b *TBitBtn) SetModalResult(value TModalResult) {
    BitBtn_SetModalResult(b.instance, value)
}

func (b *TBitBtn) NumGlyphs() TNumGlyphs {
    return BitBtn_GetNumGlyphs(b.instance)
}

func (b *TBitBtn) SetNumGlyphs(value TNumGlyphs) {
    BitBtn_SetNumGlyphs(b.instance, value)
}

func (b *TBitBtn) ParentDoubleBuffered() bool {
    return BitBtn_GetParentDoubleBuffered(b.instance)
}

func (b *TBitBtn) SetParentDoubleBuffered(value bool) {
    BitBtn_SetParentDoubleBuffered(b.instance, value)
}

func (b *TBitBtn) ParentFont() bool {
    return BitBtn_GetParentFont(b.instance)
}

func (b *TBitBtn) SetParentFont(value bool) {
    BitBtn_SetParentFont(b.instance, value)
}

func (b *TBitBtn) ParentShowHint() bool {
    return BitBtn_GetParentShowHint(b.instance)
}

func (b *TBitBtn) SetParentShowHint(value bool) {
    BitBtn_SetParentShowHint(b.instance, value)
}

func (b *TBitBtn) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(BitBtn_GetPopupMenu(b.instance))
}

func (b *TBitBtn) SetPopupMenu(value IComponent) {
    BitBtn_SetPopupMenu(b.instance, CheckPtr(value))
}

func (b *TBitBtn) ShowHint() bool {
    return BitBtn_GetShowHint(b.instance)
}

func (b *TBitBtn) SetShowHint(value bool) {
    BitBtn_SetShowHint(b.instance, value)
}

func (b *TBitBtn) Style() TButtonStyle {
    return BitBtn_GetStyle(b.instance)
}

func (b *TBitBtn) SetStyle(value TButtonStyle) {
    BitBtn_SetStyle(b.instance, value)
}

func (b *TBitBtn) Spacing() int32 {
    return BitBtn_GetSpacing(b.instance)
}

func (b *TBitBtn) SetSpacing(value int32) {
    BitBtn_SetSpacing(b.instance, value)
}

func (b *TBitBtn) TabOrder() uint16 {
    return BitBtn_GetTabOrder(b.instance)
}

func (b *TBitBtn) SetTabOrder(value uint16) {
    BitBtn_SetTabOrder(b.instance, value)
}

func (b *TBitBtn) TabStop() bool {
    return BitBtn_GetTabStop(b.instance)
}

func (b *TBitBtn) SetTabStop(value bool) {
    BitBtn_SetTabStop(b.instance, value)
}

func (b *TBitBtn) Visible() bool {
    return BitBtn_GetVisible(b.instance)
}

func (b *TBitBtn) SetVisible(value bool) {
    BitBtn_SetVisible(b.instance, value)
}

func (b *TBitBtn) WordWrap() bool {
    return BitBtn_GetWordWrap(b.instance)
}

func (b *TBitBtn) SetWordWrap(value bool) {
    BitBtn_SetWordWrap(b.instance, value)
}

func (b *TBitBtn) StyleElements() TStyleElements {
    return BitBtn_GetStyleElements(b.instance)
}

func (b *TBitBtn) SetStyleElements(value TStyleElements) {
    BitBtn_SetStyleElements(b.instance, value)
}

func (b *TBitBtn) SetOnClick(fn TNotifyEvent) {
    BitBtn_SetOnClick(b.instance, fn)
}

func (b *TBitBtn) SetOnEnter(fn TNotifyEvent) {
    BitBtn_SetOnEnter(b.instance, fn)
}

func (b *TBitBtn) SetOnExit(fn TNotifyEvent) {
    BitBtn_SetOnExit(b.instance, fn)
}

func (b *TBitBtn) SetOnKeyDown(fn TKeyEvent) {
    BitBtn_SetOnKeyDown(b.instance, fn)
}

func (b *TBitBtn) SetOnKeyPress(fn TKeyPressEvent) {
    BitBtn_SetOnKeyPress(b.instance, fn)
}

func (b *TBitBtn) SetOnKeyUp(fn TKeyEvent) {
    BitBtn_SetOnKeyUp(b.instance, fn)
}

func (b *TBitBtn) SetOnMouseDown(fn TMouseEvent) {
    BitBtn_SetOnMouseDown(b.instance, fn)
}

func (b *TBitBtn) SetOnMouseEnter(fn TNotifyEvent) {
    BitBtn_SetOnMouseEnter(b.instance, fn)
}

func (b *TBitBtn) SetOnMouseLeave(fn TNotifyEvent) {
    BitBtn_SetOnMouseLeave(b.instance, fn)
}

func (b *TBitBtn) SetOnMouseMove(fn TMouseMoveEvent) {
    BitBtn_SetOnMouseMove(b.instance, fn)
}

func (b *TBitBtn) SetOnMouseUp(fn TMouseEvent) {
    BitBtn_SetOnMouseUp(b.instance, fn)
}

func (b *TBitBtn) CommandLinkHint() string {
    return BitBtn_GetCommandLinkHint(b.instance)
}

func (b *TBitBtn) SetCommandLinkHint(value string) {
    BitBtn_SetCommandLinkHint(b.instance, value)
}

func (b *TBitBtn) DisabledImageIndex() int32 {
    return BitBtn_GetDisabledImageIndex(b.instance)
}

func (b *TBitBtn) SetDisabledImageIndex(value int32) {
    BitBtn_SetDisabledImageIndex(b.instance, value)
}

func (b *TBitBtn) ElevationRequired() bool {
    return BitBtn_GetElevationRequired(b.instance)
}

func (b *TBitBtn) SetElevationRequired(value bool) {
    BitBtn_SetElevationRequired(b.instance, value)
}

func (b *TBitBtn) HotImageIndex() int32 {
    return BitBtn_GetHotImageIndex(b.instance)
}

func (b *TBitBtn) SetHotImageIndex(value int32) {
    BitBtn_SetHotImageIndex(b.instance, value)
}

func (b *TBitBtn) Images() *TImageList {
    return ImageListFromInst(BitBtn_GetImages(b.instance))
}

func (b *TBitBtn) SetImages(value IComponent) {
    BitBtn_SetImages(b.instance, CheckPtr(value))
}

func (b *TBitBtn) ImageAlignment() TImageAlignment {
    return BitBtn_GetImageAlignment(b.instance)
}

func (b *TBitBtn) SetImageAlignment(value TImageAlignment) {
    BitBtn_SetImageAlignment(b.instance, value)
}

func (b *TBitBtn) ImageIndex() int32 {
    return BitBtn_GetImageIndex(b.instance)
}

func (b *TBitBtn) SetImageIndex(value int32) {
    BitBtn_SetImageIndex(b.instance, value)
}

func (b *TBitBtn) PressedImageIndex() int32 {
    return BitBtn_GetPressedImageIndex(b.instance)
}

func (b *TBitBtn) SetPressedImageIndex(value int32) {
    BitBtn_SetPressedImageIndex(b.instance, value)
}

func (b *TBitBtn) SelectedImageIndex() int32 {
    return BitBtn_GetSelectedImageIndex(b.instance)
}

func (b *TBitBtn) SetSelectedImageIndex(value int32) {
    BitBtn_SetSelectedImageIndex(b.instance, value)
}

func (b *TBitBtn) StylusHotImageIndex() int32 {
    return BitBtn_GetStylusHotImageIndex(b.instance)
}

func (b *TBitBtn) SetStylusHotImageIndex(value int32) {
    BitBtn_SetStylusHotImageIndex(b.instance, value)
}

func (b *TBitBtn) Brush() *TBrush {
    return BrushFromInst(BitBtn_GetBrush(b.instance))
}

func (b *TBitBtn) ControlCount() int32 {
    return BitBtn_GetControlCount(b.instance)
}

func (b *TBitBtn) Handle() HWND {
    return BitBtn_GetHandle(b.instance)
}

func (b *TBitBtn) ParentWindow() HWND {
    return BitBtn_GetParentWindow(b.instance)
}

func (b *TBitBtn) SetParentWindow(value HWND) {
    BitBtn_SetParentWindow(b.instance, value)
}

func (b *TBitBtn) BoundsRect() TRect {
    return BitBtn_GetBoundsRect(b.instance)
}

func (b *TBitBtn) SetBoundsRect(value TRect) {
    BitBtn_SetBoundsRect(b.instance, value)
}

func (b *TBitBtn) ClientHeight() int32 {
    return BitBtn_GetClientHeight(b.instance)
}

func (b *TBitBtn) SetClientHeight(value int32) {
    BitBtn_SetClientHeight(b.instance, value)
}

func (b *TBitBtn) ClientRect() TRect {
    return BitBtn_GetClientRect(b.instance)
}

func (b *TBitBtn) ClientWidth() int32 {
    return BitBtn_GetClientWidth(b.instance)
}

func (b *TBitBtn) SetClientWidth(value int32) {
    BitBtn_SetClientWidth(b.instance, value)
}

func (b *TBitBtn) ExplicitLeft() int32 {
    return BitBtn_GetExplicitLeft(b.instance)
}

func (b *TBitBtn) ExplicitTop() int32 {
    return BitBtn_GetExplicitTop(b.instance)
}

func (b *TBitBtn) ExplicitWidth() int32 {
    return BitBtn_GetExplicitWidth(b.instance)
}

func (b *TBitBtn) ExplicitHeight() int32 {
    return BitBtn_GetExplicitHeight(b.instance)
}

func (b *TBitBtn) Parent() *TControl {
    return ControlFromInst(BitBtn_GetParent(b.instance))
}

func (b *TBitBtn) SetParent(value IControl) {
    BitBtn_SetParent(b.instance, CheckPtr(value))
}

func (b *TBitBtn) AlignWithMargins() bool {
    return BitBtn_GetAlignWithMargins(b.instance)
}

func (b *TBitBtn) SetAlignWithMargins(value bool) {
    BitBtn_SetAlignWithMargins(b.instance, value)
}

func (b *TBitBtn) Left() int32 {
    return BitBtn_GetLeft(b.instance)
}

func (b *TBitBtn) SetLeft(value int32) {
    BitBtn_SetLeft(b.instance, value)
}

func (b *TBitBtn) Top() int32 {
    return BitBtn_GetTop(b.instance)
}

func (b *TBitBtn) SetTop(value int32) {
    BitBtn_SetTop(b.instance, value)
}

func (b *TBitBtn) Width() int32 {
    return BitBtn_GetWidth(b.instance)
}

func (b *TBitBtn) SetWidth(value int32) {
    BitBtn_SetWidth(b.instance, value)
}

func (b *TBitBtn) Height() int32 {
    return BitBtn_GetHeight(b.instance)
}

func (b *TBitBtn) SetHeight(value int32) {
    BitBtn_SetHeight(b.instance, value)
}

func (b *TBitBtn) Cursor() TCursor {
    return BitBtn_GetCursor(b.instance)
}

func (b *TBitBtn) SetCursor(value TCursor) {
    BitBtn_SetCursor(b.instance, value)
}

func (b *TBitBtn) Hint() string {
    return BitBtn_GetHint(b.instance)
}

func (b *TBitBtn) SetHint(value string) {
    BitBtn_SetHint(b.instance, value)
}

func (b *TBitBtn) Margins() *TMargins {
    return MarginsFromInst(BitBtn_GetMargins(b.instance))
}

func (b *TBitBtn) SetMargins(value *TMargins) {
    BitBtn_SetMargins(b.instance, CheckPtr(value))
}

func (b *TBitBtn) CustomHint() *TCustomHint {
    return CustomHintFromInst(BitBtn_GetCustomHint(b.instance))
}

func (b *TBitBtn) SetCustomHint(value IComponent) {
    BitBtn_SetCustomHint(b.instance, CheckPtr(value))
}

func (b *TBitBtn) ComponentCount() int32 {
    return BitBtn_GetComponentCount(b.instance)
}

func (b *TBitBtn) ComponentIndex() int32 {
    return BitBtn_GetComponentIndex(b.instance)
}

func (b *TBitBtn) SetComponentIndex(value int32) {
    BitBtn_SetComponentIndex(b.instance, value)
}

func (b *TBitBtn) Owner() *TComponent {
    return ComponentFromInst(BitBtn_GetOwner(b.instance))
}

func (b *TBitBtn) Name() string {
    return BitBtn_GetName(b.instance)
}

func (b *TBitBtn) SetName(value string) {
    BitBtn_SetName(b.instance, value)
}

func (b *TBitBtn) Tag() int {
    return BitBtn_GetTag(b.instance)
}

func (b *TBitBtn) SetTag(value int) {
    BitBtn_SetTag(b.instance, value)
}

func (b *TBitBtn) Controls(Index int32) *TControl {
    return ControlFromInst(BitBtn_GetControls(b.instance, Index))
}

func (b *TBitBtn) Components(AIndex int32) *TComponent {
    return ComponentFromInst(BitBtn_GetComponents(b.instance, AIndex))
}

