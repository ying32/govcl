
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

type TForm struct {
    IControl
    instance uintptr
}

func NewForm(owner IComponent) *TForm {
    f := new(TForm)
    f.instance = Form_Create(CheckPtr(owner))
    return f
}

func FormFromInst(inst uintptr) *TForm {
    f := new(TForm)
    f.instance = inst
    return f
}

func FormFromObj(obj IObject) *TForm {
    f := new(TForm)
    f.instance = CheckPtr(obj)
    return f
}

func (f *TForm) Free() {
    if f.instance != 0 {
        Form_Free(f.instance)
        f.instance = 0
    }
}

func (f *TForm) Instance() uintptr {
    return f.instance
}

func (f *TForm) IsValid() bool {
    return f.instance != 0
}

func (f *TForm) Close() {
    Form_Close(f.instance)
}

func (f *TForm) Hide() {
    Form_Hide(f.instance)
}

func (f *TForm) Print() {
    Form_Print(f.instance)
}

func (f *TForm) SetFocus() {
    Form_SetFocus(f.instance)
}

func (f *TForm) Show() {
    Form_Show(f.instance)
}

func (f *TForm) ShowModal() int32 {
    return Form_ShowModal(f.instance)
}

func (f *TForm) CanFocus() bool {
    return Form_CanFocus(f.instance)
}

func (f *TForm) FlipChildren(AllLevels bool) {
    Form_FlipChildren(f.instance, AllLevels)
}

func (f *TForm) Focused() bool {
    return Form_Focused(f.instance)
}

func (f *TForm) HandleAllocated() bool {
    return Form_HandleAllocated(f.instance)
}

func (f *TForm) Invalidate() {
    Form_Invalidate(f.instance)
}

func (f *TForm) Realign() {
    Form_Realign(f.instance)
}

func (f *TForm) Repaint() {
    Form_Repaint(f.instance)
}

func (f *TForm) ScaleBy(M int32, D int32) {
    Form_ScaleBy(f.instance, M , D)
}

func (f *TForm) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Form_SetBounds(f.instance, ALeft , ATop , AWidth , AHeight)
}

func (f *TForm) Update() {
    Form_Update(f.instance)
}

func (f *TForm) BringToFront() {
    Form_BringToFront(f.instance)
}

func (f *TForm) HasParent() bool {
    return Form_HasParent(f.instance)
}

func (f *TForm) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Form_Perform(f.instance, Msg , WParam , LParam)
}

func (f *TForm) Refresh() {
    Form_Refresh(f.instance)
}

func (f *TForm) SendToBack() {
    Form_SendToBack(f.instance)
}

func (f *TForm) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Form_GetTextBuf(f.instance, Buffer , BufSize)
}

func (f *TForm) GetTextLen() int32 {
    return Form_GetTextLen(f.instance)
}

func (f *TForm) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Form_FindComponent(f.instance, AName))
}

func (f *TForm) GetNamePath() string {
    return Form_GetNamePath(f.instance)
}

func (f *TForm) Assign(Source IObject) {
    Form_Assign(f.instance, CheckPtr(Source))
}

func (f *TForm) ClassName() string {
    return Form_ClassName(f.instance)
}

func (f *TForm) Equals(Obj IObject) bool {
    return Form_Equals(f.instance, CheckPtr(Obj))
}

func (f *TForm) GetHashCode() int32 {
    return Form_GetHashCode(f.instance)
}

func (f *TForm) ToString() string {
    return Form_ToString(f.instance)
}

func (f *TForm) Action() *TAction {
    return ActionFromInst(Form_GetAction(f.instance))
}

func (f *TForm) SetAction(value IComponent) {
    Form_SetAction(f.instance, CheckPtr(value))
}

func (f *TForm) Align() TAlign {
    return Form_GetAlign(f.instance)
}

func (f *TForm) SetAlign(value TAlign) {
    Form_SetAlign(f.instance, value)
}

func (f *TForm) AlphaBlend() bool {
    return Form_GetAlphaBlend(f.instance)
}

func (f *TForm) SetAlphaBlend(value bool) {
    Form_SetAlphaBlend(f.instance, value)
}

func (f *TForm) AlphaBlendValue() uint8 {
    return Form_GetAlphaBlendValue(f.instance)
}

func (f *TForm) SetAlphaBlendValue(value uint8) {
    Form_SetAlphaBlendValue(f.instance, value)
}

func (f *TForm) Anchors() TAnchors {
    return Form_GetAnchors(f.instance)
}

func (f *TForm) SetAnchors(value TAnchors) {
    Form_SetAnchors(f.instance, value)
}

func (f *TForm) AutoSize() bool {
    return Form_GetAutoSize(f.instance)
}

func (f *TForm) SetAutoSize(value bool) {
    Form_SetAutoSize(f.instance, value)
}

func (f *TForm) BiDiMode() TBiDiMode {
    return Form_GetBiDiMode(f.instance)
}

func (f *TForm) SetBiDiMode(value TBiDiMode) {
    Form_SetBiDiMode(f.instance, value)
}

func (f *TForm) BorderIcons() TBorderIcons {
    return Form_GetBorderIcons(f.instance)
}

func (f *TForm) SetBorderIcons(value TBorderIcons) {
    Form_SetBorderIcons(f.instance, value)
}

func (f *TForm) BorderStyle() TFormBorderStyle {
    return Form_GetBorderStyle(f.instance)
}

func (f *TForm) SetBorderStyle(value TFormBorderStyle) {
    Form_SetBorderStyle(f.instance, value)
}

func (f *TForm) BorderWidth() int32 {
    return Form_GetBorderWidth(f.instance)
}

func (f *TForm) SetBorderWidth(value int32) {
    Form_SetBorderWidth(f.instance, value)
}

func (f *TForm) Caption() string {
    return Form_GetCaption(f.instance)
}

func (f *TForm) SetCaption(value string) {
    Form_SetCaption(f.instance, value)
}

func (f *TForm) ClientHeight() int32 {
    return Form_GetClientHeight(f.instance)
}

func (f *TForm) SetClientHeight(value int32) {
    Form_SetClientHeight(f.instance, value)
}

func (f *TForm) ClientWidth() int32 {
    return Form_GetClientWidth(f.instance)
}

func (f *TForm) SetClientWidth(value int32) {
    Form_SetClientWidth(f.instance, value)
}

func (f *TForm) Color() TColor {
    return Form_GetColor(f.instance)
}

func (f *TForm) SetColor(value TColor) {
    Form_SetColor(f.instance, value)
}

func (f *TForm) TransparentColor() bool {
    return Form_GetTransparentColor(f.instance)
}

func (f *TForm) SetTransparentColor(value bool) {
    Form_SetTransparentColor(f.instance, value)
}

func (f *TForm) TransparentColorValue() TColor {
    return Form_GetTransparentColorValue(f.instance)
}

func (f *TForm) SetTransparentColorValue(value TColor) {
    Form_SetTransparentColorValue(f.instance, value)
}

func (f *TForm) DoubleBuffered() bool {
    return Form_GetDoubleBuffered(f.instance)
}

func (f *TForm) SetDoubleBuffered(value bool) {
    Form_SetDoubleBuffered(f.instance, value)
}

func (f *TForm) Enabled() bool {
    return Form_GetEnabled(f.instance)
}

func (f *TForm) SetEnabled(value bool) {
    Form_SetEnabled(f.instance, value)
}

func (f *TForm) ParentFont() bool {
    return Form_GetParentFont(f.instance)
}

func (f *TForm) SetParentFont(value bool) {
    Form_SetParentFont(f.instance, value)
}

func (f *TForm) Font() *TFont {
    return FontFromInst(Form_GetFont(f.instance))
}

func (f *TForm) SetFont(value *TFont) {
    Form_SetFont(f.instance, CheckPtr(value))
}

func (f *TForm) FormStyle() TFormStyle {
    return Form_GetFormStyle(f.instance)
}

func (f *TForm) SetFormStyle(value TFormStyle) {
    Form_SetFormStyle(f.instance, value)
}

func (f *TForm) Height() int32 {
    return Form_GetHeight(f.instance)
}

func (f *TForm) SetHeight(value int32) {
    Form_SetHeight(f.instance, value)
}

func (f *TForm) Icon() *TIcon {
    return IconFromInst(Form_GetIcon(f.instance))
}

func (f *TForm) SetIcon(value *TIcon) {
    Form_SetIcon(f.instance, CheckPtr(value))
}

func (f *TForm) KeyPreview() bool {
    return Form_GetKeyPreview(f.instance)
}

func (f *TForm) SetKeyPreview(value bool) {
    Form_SetKeyPreview(f.instance, value)
}

func (f *TForm) Menu() *TMainMenu {
    return MainMenuFromInst(Form_GetMenu(f.instance))
}

func (f *TForm) SetMenu(value IComponent) {
    Form_SetMenu(f.instance, CheckPtr(value))
}

func (f *TForm) PixelsPerInch() int32 {
    return Form_GetPixelsPerInch(f.instance)
}

func (f *TForm) SetPixelsPerInch(value int32) {
    Form_SetPixelsPerInch(f.instance, value)
}

func (f *TForm) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Form_GetPopupMenu(f.instance))
}

func (f *TForm) SetPopupMenu(value IComponent) {
    Form_SetPopupMenu(f.instance, CheckPtr(value))
}

func (f *TForm) Position() TPosition {
    return Form_GetPosition(f.instance)
}

func (f *TForm) SetPosition(value TPosition) {
    Form_SetPosition(f.instance, value)
}

func (f *TForm) Scaled() bool {
    return Form_GetScaled(f.instance)
}

func (f *TForm) SetScaled(value bool) {
    Form_SetScaled(f.instance, value)
}

func (f *TForm) ShowHint() bool {
    return Form_GetShowHint(f.instance)
}

func (f *TForm) SetShowHint(value bool) {
    Form_SetShowHint(f.instance, value)
}

func (f *TForm) Visible() bool {
    return Form_GetVisible(f.instance)
}

func (f *TForm) SetVisible(value bool) {
    Form_SetVisible(f.instance, value)
}

func (f *TForm) Width() int32 {
    return Form_GetWidth(f.instance)
}

func (f *TForm) SetWidth(value int32) {
    Form_SetWidth(f.instance, value)
}

func (f *TForm) WindowState() TWindowState {
    return Form_GetWindowState(f.instance)
}

func (f *TForm) SetWindowState(value TWindowState) {
    Form_SetWindowState(f.instance, value)
}

func (f *TForm) StyleElements() TStyleElements {
    return Form_GetStyleElements(f.instance)
}

func (f *TForm) SetStyleElements(value TStyleElements) {
    Form_SetStyleElements(f.instance, value)
}

func (f *TForm) SetOnClick(fn TNotifyEvent) {
    Form_SetOnClick(f.instance, fn)
}

func (f *TForm) SetOnClose(fn TCloseEvent) {
    Form_SetOnClose(f.instance, fn)
}

func (f *TForm) SetOnCloseQuery(fn TCloseQueryEvent) {
    Form_SetOnCloseQuery(f.instance, fn)
}

func (f *TForm) SetOnDblClick(fn TNotifyEvent) {
    Form_SetOnDblClick(f.instance, fn)
}

func (f *TForm) SetOnHide(fn TNotifyEvent) {
    Form_SetOnHide(f.instance, fn)
}

func (f *TForm) SetOnKeyDown(fn TKeyEvent) {
    Form_SetOnKeyDown(f.instance, fn)
}

func (f *TForm) SetOnKeyPress(fn TKeyPressEvent) {
    Form_SetOnKeyPress(f.instance, fn)
}

func (f *TForm) SetOnKeyUp(fn TKeyEvent) {
    Form_SetOnKeyUp(f.instance, fn)
}

func (f *TForm) SetOnMouseDown(fn TMouseEvent) {
    Form_SetOnMouseDown(f.instance, fn)
}

func (f *TForm) SetOnMouseEnter(fn TNotifyEvent) {
    Form_SetOnMouseEnter(f.instance, fn)
}

func (f *TForm) SetOnMouseLeave(fn TNotifyEvent) {
    Form_SetOnMouseLeave(f.instance, fn)
}

func (f *TForm) SetOnMouseMove(fn TMouseMoveEvent) {
    Form_SetOnMouseMove(f.instance, fn)
}

func (f *TForm) SetOnMouseUp(fn TMouseEvent) {
    Form_SetOnMouseUp(f.instance, fn)
}

func (f *TForm) SetOnMouseWheel(fn TMouseWheelEvent) {
    Form_SetOnMouseWheel(f.instance, fn)
}

func (f *TForm) SetOnPaint(fn TNotifyEvent) {
    Form_SetOnPaint(f.instance, fn)
}

func (f *TForm) SetOnResize(fn TNotifyEvent) {
    Form_SetOnResize(f.instance, fn)
}

func (f *TForm) SetOnShow(fn TNotifyEvent) {
    Form_SetOnShow(f.instance, fn)
}

func (f *TForm) Canvas() *TCanvas {
    return CanvasFromInst(Form_GetCanvas(f.instance))
}

func (f *TForm) DropTarget() bool {
    return Form_GetDropTarget(f.instance)
}

func (f *TForm) SetDropTarget(value bool) {
    Form_SetDropTarget(f.instance, value)
}

func (f *TForm) ModalResult() TModalResult {
    return Form_GetModalResult(f.instance)
}

func (f *TForm) SetModalResult(value TModalResult) {
    Form_SetModalResult(f.instance, value)
}

func (f *TForm) Left() int32 {
    return Form_GetLeft(f.instance)
}

func (f *TForm) SetLeft(value int32) {
    Form_SetLeft(f.instance, value)
}

func (f *TForm) Top() int32 {
    return Form_GetTop(f.instance)
}

func (f *TForm) SetTop(value int32) {
    Form_SetTop(f.instance, value)
}

func (f *TForm) Brush() *TBrush {
    return BrushFromInst(Form_GetBrush(f.instance))
}

func (f *TForm) ControlCount() int32 {
    return Form_GetControlCount(f.instance)
}

func (f *TForm) Handle() HWND {
    return Form_GetHandle(f.instance)
}

func (f *TForm) ParentDoubleBuffered() bool {
    return Form_GetParentDoubleBuffered(f.instance)
}

func (f *TForm) SetParentDoubleBuffered(value bool) {
    Form_SetParentDoubleBuffered(f.instance, value)
}

func (f *TForm) ParentWindow() HWND {
    return Form_GetParentWindow(f.instance)
}

func (f *TForm) SetParentWindow(value HWND) {
    Form_SetParentWindow(f.instance, value)
}

func (f *TForm) TabOrder() uint16 {
    return Form_GetTabOrder(f.instance)
}

func (f *TForm) SetTabOrder(value uint16) {
    Form_SetTabOrder(f.instance, value)
}

func (f *TForm) TabStop() bool {
    return Form_GetTabStop(f.instance)
}

func (f *TForm) SetTabStop(value bool) {
    Form_SetTabStop(f.instance, value)
}

func (f *TForm) BoundsRect() TRect {
    return Form_GetBoundsRect(f.instance)
}

func (f *TForm) SetBoundsRect(value TRect) {
    Form_SetBoundsRect(f.instance, value)
}

func (f *TForm) ClientRect() TRect {
    return Form_GetClientRect(f.instance)
}

func (f *TForm) ExplicitLeft() int32 {
    return Form_GetExplicitLeft(f.instance)
}

func (f *TForm) ExplicitTop() int32 {
    return Form_GetExplicitTop(f.instance)
}

func (f *TForm) ExplicitWidth() int32 {
    return Form_GetExplicitWidth(f.instance)
}

func (f *TForm) ExplicitHeight() int32 {
    return Form_GetExplicitHeight(f.instance)
}

func (f *TForm) Parent() *TControl {
    return ControlFromInst(Form_GetParent(f.instance))
}

func (f *TForm) SetParent(value IControl) {
    Form_SetParent(f.instance, CheckPtr(value))
}

func (f *TForm) AlignWithMargins() bool {
    return Form_GetAlignWithMargins(f.instance)
}

func (f *TForm) SetAlignWithMargins(value bool) {
    Form_SetAlignWithMargins(f.instance, value)
}

func (f *TForm) Cursor() TCursor {
    return Form_GetCursor(f.instance)
}

func (f *TForm) SetCursor(value TCursor) {
    Form_SetCursor(f.instance, value)
}

func (f *TForm) Hint() string {
    return Form_GetHint(f.instance)
}

func (f *TForm) SetHint(value string) {
    Form_SetHint(f.instance, value)
}

func (f *TForm) Margins() *TMargins {
    return MarginsFromInst(Form_GetMargins(f.instance))
}

func (f *TForm) SetMargins(value *TMargins) {
    Form_SetMargins(f.instance, CheckPtr(value))
}

func (f *TForm) CustomHint() *TCustomHint {
    return CustomHintFromInst(Form_GetCustomHint(f.instance))
}

func (f *TForm) SetCustomHint(value IComponent) {
    Form_SetCustomHint(f.instance, CheckPtr(value))
}

func (f *TForm) ComponentCount() int32 {
    return Form_GetComponentCount(f.instance)
}

func (f *TForm) ComponentIndex() int32 {
    return Form_GetComponentIndex(f.instance)
}

func (f *TForm) SetComponentIndex(value int32) {
    Form_SetComponentIndex(f.instance, value)
}

func (f *TForm) Owner() *TComponent {
    return ComponentFromInst(Form_GetOwner(f.instance))
}

func (f *TForm) Name() string {
    return Form_GetName(f.instance)
}

func (f *TForm) SetName(value string) {
    Form_SetName(f.instance, value)
}

func (f *TForm) Tag() int {
    return Form_GetTag(f.instance)
}

func (f *TForm) SetTag(value int) {
    Form_SetTag(f.instance, value)
}

func (f *TForm) Controls(Index int32) *TControl {
    return ControlFromInst(Form_GetControls(f.instance, Index))
}

func (f *TForm) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Form_GetComponents(f.instance, AIndex))
}

