
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

type TButton struct {
    IControl
    instance uintptr
}

func NewButton(owner IComponent) *TButton {
    b := new(TButton)
    b.instance = Button_Create(CheckPtr(owner))
    return b
}

func ButtonFromInst(inst uintptr) *TButton {
    b := new(TButton)
    b.instance = inst
    return b
}

func ButtonFromObj(obj IObject) *TButton {
    b := new(TButton)
    b.instance = CheckPtr(obj)
    return b
}

func (b *TButton) Free() {
    if b.instance != 0 {
        Button_Free(b.instance)
        b.instance = 0
    }
}

func (b *TButton) Instance() uintptr {
    return b.instance
}

func (b *TButton) IsValid() bool {
    return b.instance != 0
}

func (b *TButton) Click() {
    Button_Click(b.instance)
}

func (b *TButton) CanFocus() bool {
    return Button_CanFocus(b.instance)
}

func (b *TButton) FlipChildren(AllLevels bool) {
    Button_FlipChildren(b.instance, AllLevels)
}

func (b *TButton) Focused() bool {
    return Button_Focused(b.instance)
}

func (b *TButton) HandleAllocated() bool {
    return Button_HandleAllocated(b.instance)
}

func (b *TButton) Invalidate() {
    Button_Invalidate(b.instance)
}

func (b *TButton) Realign() {
    Button_Realign(b.instance)
}

func (b *TButton) Repaint() {
    Button_Repaint(b.instance)
}

func (b *TButton) ScaleBy(M int32, D int32) {
    Button_ScaleBy(b.instance, M , D)
}

func (b *TButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Button_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

func (b *TButton) SetFocus() {
    Button_SetFocus(b.instance)
}

func (b *TButton) Update() {
    Button_Update(b.instance)
}

func (b *TButton) BringToFront() {
    Button_BringToFront(b.instance)
}

func (b *TButton) HasParent() bool {
    return Button_HasParent(b.instance)
}

func (b *TButton) Hide() {
    Button_Hide(b.instance)
}

func (b *TButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Button_Perform(b.instance, Msg , WParam , LParam)
}

func (b *TButton) Refresh() {
    Button_Refresh(b.instance)
}

func (b *TButton) SendToBack() {
    Button_SendToBack(b.instance)
}

func (b *TButton) Show() {
    Button_Show(b.instance)
}

func (b *TButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Button_GetTextBuf(b.instance, Buffer , BufSize)
}

func (b *TButton) GetTextLen() int32 {
    return Button_GetTextLen(b.instance)
}

func (b *TButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Button_FindComponent(b.instance, AName))
}

func (b *TButton) GetNamePath() string {
    return Button_GetNamePath(b.instance)
}

func (b *TButton) Assign(Source IObject) {
    Button_Assign(b.instance, CheckPtr(Source))
}

func (b *TButton) ClassName() string {
    return Button_ClassName(b.instance)
}

func (b *TButton) Equals(Obj IObject) bool {
    return Button_Equals(b.instance, CheckPtr(Obj))
}

func (b *TButton) GetHashCode() int32 {
    return Button_GetHashCode(b.instance)
}

func (b *TButton) ToString() string {
    return Button_ToString(b.instance)
}

func (b *TButton) Action() *TAction {
    return ActionFromInst(Button_GetAction(b.instance))
}

func (b *TButton) SetAction(value IComponent) {
    Button_SetAction(b.instance, CheckPtr(value))
}

func (b *TButton) Align() TAlign {
    return Button_GetAlign(b.instance)
}

func (b *TButton) SetAlign(value TAlign) {
    Button_SetAlign(b.instance, value)
}

func (b *TButton) Anchors() TAnchors {
    return Button_GetAnchors(b.instance)
}

func (b *TButton) SetAnchors(value TAnchors) {
    Button_SetAnchors(b.instance, value)
}

func (b *TButton) BiDiMode() TBiDiMode {
    return Button_GetBiDiMode(b.instance)
}

func (b *TButton) SetBiDiMode(value TBiDiMode) {
    Button_SetBiDiMode(b.instance, value)
}

func (b *TButton) Cancel() bool {
    return Button_GetCancel(b.instance)
}

func (b *TButton) SetCancel(value bool) {
    Button_SetCancel(b.instance, value)
}

func (b *TButton) Caption() string {
    return Button_GetCaption(b.instance)
}

func (b *TButton) SetCaption(value string) {
    Button_SetCaption(b.instance, value)
}

func (b *TButton) CommandLinkHint() string {
    return Button_GetCommandLinkHint(b.instance)
}

func (b *TButton) SetCommandLinkHint(value string) {
    Button_SetCommandLinkHint(b.instance, value)
}

func (b *TButton) Default() bool {
    return Button_GetDefault(b.instance)
}

func (b *TButton) SetDefault(value bool) {
    Button_SetDefault(b.instance, value)
}

func (b *TButton) DisabledImageIndex() int32 {
    return Button_GetDisabledImageIndex(b.instance)
}

func (b *TButton) SetDisabledImageIndex(value int32) {
    Button_SetDisabledImageIndex(b.instance, value)
}

func (b *TButton) DoubleBuffered() bool {
    return Button_GetDoubleBuffered(b.instance)
}

func (b *TButton) SetDoubleBuffered(value bool) {
    Button_SetDoubleBuffered(b.instance, value)
}

func (b *TButton) ElevationRequired() bool {
    return Button_GetElevationRequired(b.instance)
}

func (b *TButton) SetElevationRequired(value bool) {
    Button_SetElevationRequired(b.instance, value)
}

func (b *TButton) Enabled() bool {
    return Button_GetEnabled(b.instance)
}

func (b *TButton) SetEnabled(value bool) {
    Button_SetEnabled(b.instance, value)
}

func (b *TButton) Font() *TFont {
    return FontFromInst(Button_GetFont(b.instance))
}

func (b *TButton) SetFont(value *TFont) {
    Button_SetFont(b.instance, CheckPtr(value))
}

func (b *TButton) HotImageIndex() int32 {
    return Button_GetHotImageIndex(b.instance)
}

func (b *TButton) SetHotImageIndex(value int32) {
    Button_SetHotImageIndex(b.instance, value)
}

func (b *TButton) ImageAlignment() TImageAlignment {
    return Button_GetImageAlignment(b.instance)
}

func (b *TButton) SetImageAlignment(value TImageAlignment) {
    Button_SetImageAlignment(b.instance, value)
}

func (b *TButton) ImageIndex() int32 {
    return Button_GetImageIndex(b.instance)
}

func (b *TButton) SetImageIndex(value int32) {
    Button_SetImageIndex(b.instance, value)
}

func (b *TButton) Images() *TImageList {
    return ImageListFromInst(Button_GetImages(b.instance))
}

func (b *TButton) SetImages(value IComponent) {
    Button_SetImages(b.instance, CheckPtr(value))
}

func (b *TButton) ModalResult() TModalResult {
    return Button_GetModalResult(b.instance)
}

func (b *TButton) SetModalResult(value TModalResult) {
    Button_SetModalResult(b.instance, value)
}

func (b *TButton) ParentDoubleBuffered() bool {
    return Button_GetParentDoubleBuffered(b.instance)
}

func (b *TButton) SetParentDoubleBuffered(value bool) {
    Button_SetParentDoubleBuffered(b.instance, value)
}

func (b *TButton) ParentFont() bool {
    return Button_GetParentFont(b.instance)
}

func (b *TButton) SetParentFont(value bool) {
    Button_SetParentFont(b.instance, value)
}

func (b *TButton) ParentShowHint() bool {
    return Button_GetParentShowHint(b.instance)
}

func (b *TButton) SetParentShowHint(value bool) {
    Button_SetParentShowHint(b.instance, value)
}

func (b *TButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Button_GetPopupMenu(b.instance))
}

func (b *TButton) SetPopupMenu(value IComponent) {
    Button_SetPopupMenu(b.instance, CheckPtr(value))
}

func (b *TButton) PressedImageIndex() int32 {
    return Button_GetPressedImageIndex(b.instance)
}

func (b *TButton) SetPressedImageIndex(value int32) {
    Button_SetPressedImageIndex(b.instance, value)
}

func (b *TButton) SelectedImageIndex() int32 {
    return Button_GetSelectedImageIndex(b.instance)
}

func (b *TButton) SetSelectedImageIndex(value int32) {
    Button_SetSelectedImageIndex(b.instance, value)
}

func (b *TButton) ShowHint() bool {
    return Button_GetShowHint(b.instance)
}

func (b *TButton) SetShowHint(value bool) {
    Button_SetShowHint(b.instance, value)
}

func (b *TButton) Style() TButtonStyle {
    return Button_GetStyle(b.instance)
}

func (b *TButton) SetStyle(value TButtonStyle) {
    Button_SetStyle(b.instance, value)
}

func (b *TButton) StylusHotImageIndex() int32 {
    return Button_GetStylusHotImageIndex(b.instance)
}

func (b *TButton) SetStylusHotImageIndex(value int32) {
    Button_SetStylusHotImageIndex(b.instance, value)
}

func (b *TButton) TabOrder() uint16 {
    return Button_GetTabOrder(b.instance)
}

func (b *TButton) SetTabOrder(value uint16) {
    Button_SetTabOrder(b.instance, value)
}

func (b *TButton) TabStop() bool {
    return Button_GetTabStop(b.instance)
}

func (b *TButton) SetTabStop(value bool) {
    Button_SetTabStop(b.instance, value)
}

func (b *TButton) Visible() bool {
    return Button_GetVisible(b.instance)
}

func (b *TButton) SetVisible(value bool) {
    Button_SetVisible(b.instance, value)
}

func (b *TButton) WordWrap() bool {
    return Button_GetWordWrap(b.instance)
}

func (b *TButton) SetWordWrap(value bool) {
    Button_SetWordWrap(b.instance, value)
}

func (b *TButton) StyleElements() TStyleElements {
    return Button_GetStyleElements(b.instance)
}

func (b *TButton) SetStyleElements(value TStyleElements) {
    Button_SetStyleElements(b.instance, value)
}

func (b *TButton) SetOnClick(fn TNotifyEvent) {
    Button_SetOnClick(b.instance, fn)
}

func (b *TButton) SetOnEnter(fn TNotifyEvent) {
    Button_SetOnEnter(b.instance, fn)
}

func (b *TButton) SetOnExit(fn TNotifyEvent) {
    Button_SetOnExit(b.instance, fn)
}

func (b *TButton) SetOnKeyDown(fn TKeyEvent) {
    Button_SetOnKeyDown(b.instance, fn)
}

func (b *TButton) SetOnKeyPress(fn TKeyPressEvent) {
    Button_SetOnKeyPress(b.instance, fn)
}

func (b *TButton) SetOnKeyUp(fn TKeyEvent) {
    Button_SetOnKeyUp(b.instance, fn)
}

func (b *TButton) SetOnMouseDown(fn TMouseEvent) {
    Button_SetOnMouseDown(b.instance, fn)
}

func (b *TButton) SetOnMouseEnter(fn TNotifyEvent) {
    Button_SetOnMouseEnter(b.instance, fn)
}

func (b *TButton) SetOnMouseLeave(fn TNotifyEvent) {
    Button_SetOnMouseLeave(b.instance, fn)
}

func (b *TButton) SetOnMouseMove(fn TMouseMoveEvent) {
    Button_SetOnMouseMove(b.instance, fn)
}

func (b *TButton) SetOnMouseUp(fn TMouseEvent) {
    Button_SetOnMouseUp(b.instance, fn)
}

func (b *TButton) Brush() *TBrush {
    return BrushFromInst(Button_GetBrush(b.instance))
}

func (b *TButton) ControlCount() int32 {
    return Button_GetControlCount(b.instance)
}

func (b *TButton) Handle() HWND {
    return Button_GetHandle(b.instance)
}

func (b *TButton) ParentWindow() HWND {
    return Button_GetParentWindow(b.instance)
}

func (b *TButton) SetParentWindow(value HWND) {
    Button_SetParentWindow(b.instance, value)
}

func (b *TButton) BoundsRect() TRect {
    return Button_GetBoundsRect(b.instance)
}

func (b *TButton) SetBoundsRect(value TRect) {
    Button_SetBoundsRect(b.instance, value)
}

func (b *TButton) ClientHeight() int32 {
    return Button_GetClientHeight(b.instance)
}

func (b *TButton) SetClientHeight(value int32) {
    Button_SetClientHeight(b.instance, value)
}

func (b *TButton) ClientRect() TRect {
    return Button_GetClientRect(b.instance)
}

func (b *TButton) ClientWidth() int32 {
    return Button_GetClientWidth(b.instance)
}

func (b *TButton) SetClientWidth(value int32) {
    Button_SetClientWidth(b.instance, value)
}

func (b *TButton) ExplicitLeft() int32 {
    return Button_GetExplicitLeft(b.instance)
}

func (b *TButton) ExplicitTop() int32 {
    return Button_GetExplicitTop(b.instance)
}

func (b *TButton) ExplicitWidth() int32 {
    return Button_GetExplicitWidth(b.instance)
}

func (b *TButton) ExplicitHeight() int32 {
    return Button_GetExplicitHeight(b.instance)
}

func (b *TButton) Parent() *TControl {
    return ControlFromInst(Button_GetParent(b.instance))
}

func (b *TButton) SetParent(value IControl) {
    Button_SetParent(b.instance, CheckPtr(value))
}

func (b *TButton) AlignWithMargins() bool {
    return Button_GetAlignWithMargins(b.instance)
}

func (b *TButton) SetAlignWithMargins(value bool) {
    Button_SetAlignWithMargins(b.instance, value)
}

func (b *TButton) Left() int32 {
    return Button_GetLeft(b.instance)
}

func (b *TButton) SetLeft(value int32) {
    Button_SetLeft(b.instance, value)
}

func (b *TButton) Top() int32 {
    return Button_GetTop(b.instance)
}

func (b *TButton) SetTop(value int32) {
    Button_SetTop(b.instance, value)
}

func (b *TButton) Width() int32 {
    return Button_GetWidth(b.instance)
}

func (b *TButton) SetWidth(value int32) {
    Button_SetWidth(b.instance, value)
}

func (b *TButton) Height() int32 {
    return Button_GetHeight(b.instance)
}

func (b *TButton) SetHeight(value int32) {
    Button_SetHeight(b.instance, value)
}

func (b *TButton) Cursor() TCursor {
    return Button_GetCursor(b.instance)
}

func (b *TButton) SetCursor(value TCursor) {
    Button_SetCursor(b.instance, value)
}

func (b *TButton) Hint() string {
    return Button_GetHint(b.instance)
}

func (b *TButton) SetHint(value string) {
    Button_SetHint(b.instance, value)
}

func (b *TButton) Margins() *TMargins {
    return MarginsFromInst(Button_GetMargins(b.instance))
}

func (b *TButton) SetMargins(value *TMargins) {
    Button_SetMargins(b.instance, CheckPtr(value))
}

func (b *TButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(Button_GetCustomHint(b.instance))
}

func (b *TButton) SetCustomHint(value IComponent) {
    Button_SetCustomHint(b.instance, CheckPtr(value))
}

func (b *TButton) ComponentCount() int32 {
    return Button_GetComponentCount(b.instance)
}

func (b *TButton) ComponentIndex() int32 {
    return Button_GetComponentIndex(b.instance)
}

func (b *TButton) SetComponentIndex(value int32) {
    Button_SetComponentIndex(b.instance, value)
}

func (b *TButton) Owner() *TComponent {
    return ComponentFromInst(Button_GetOwner(b.instance))
}

func (b *TButton) Name() string {
    return Button_GetName(b.instance)
}

func (b *TButton) SetName(value string) {
    Button_SetName(b.instance, value)
}

func (b *TButton) Tag() int {
    return Button_GetTag(b.instance)
}

func (b *TButton) SetTag(value int) {
    Button_SetTag(b.instance, value)
}

func (b *TButton) Controls(Index int32) *TControl {
    return ControlFromInst(Button_GetControls(b.instance, Index))
}

func (b *TButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Button_GetComponents(b.instance, AIndex))
}

