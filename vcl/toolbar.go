
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

type TToolBar struct {
    IControl
    instance uintptr
}

func NewToolBar(owner IComponent) *TToolBar {
    t := new(TToolBar)
    t.instance = ToolBar_Create(CheckPtr(owner))
    return t
}

func ToolBarFromInst(inst uintptr) *TToolBar {
    t := new(TToolBar)
    t.instance = inst
    return t
}

func ToolBarFromObj(obj IObject) *TToolBar {
    t := new(TToolBar)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TToolBar) Free() {
    if t.instance != 0 {
        ToolBar_Free(t.instance)
        t.instance = 0
    }
}

func (t *TToolBar) Instance() uintptr {
    return t.instance
}

func (t *TToolBar) IsValid() bool {
    return t.instance != 0
}

func (t *TToolBar) FlipChildren(AllLevels bool) {
    ToolBar_FlipChildren(t.instance, AllLevels)
}

func (t *TToolBar) CanFocus() bool {
    return ToolBar_CanFocus(t.instance)
}

func (t *TToolBar) Focused() bool {
    return ToolBar_Focused(t.instance)
}

func (t *TToolBar) HandleAllocated() bool {
    return ToolBar_HandleAllocated(t.instance)
}

func (t *TToolBar) Invalidate() {
    ToolBar_Invalidate(t.instance)
}

func (t *TToolBar) Realign() {
    ToolBar_Realign(t.instance)
}

func (t *TToolBar) Repaint() {
    ToolBar_Repaint(t.instance)
}

func (t *TToolBar) ScaleBy(M int32, D int32) {
    ToolBar_ScaleBy(t.instance, M , D)
}

func (t *TToolBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ToolBar_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

func (t *TToolBar) SetFocus() {
    ToolBar_SetFocus(t.instance)
}

func (t *TToolBar) Update() {
    ToolBar_Update(t.instance)
}

func (t *TToolBar) BringToFront() {
    ToolBar_BringToFront(t.instance)
}

func (t *TToolBar) HasParent() bool {
    return ToolBar_HasParent(t.instance)
}

func (t *TToolBar) Hide() {
    ToolBar_Hide(t.instance)
}

func (t *TToolBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ToolBar_Perform(t.instance, Msg , WParam , LParam)
}

func (t *TToolBar) Refresh() {
    ToolBar_Refresh(t.instance)
}

func (t *TToolBar) SendToBack() {
    ToolBar_SendToBack(t.instance)
}

func (t *TToolBar) Show() {
    ToolBar_Show(t.instance)
}

func (t *TToolBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ToolBar_GetTextBuf(t.instance, Buffer , BufSize)
}

func (t *TToolBar) GetTextLen() int32 {
    return ToolBar_GetTextLen(t.instance)
}

func (t *TToolBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ToolBar_FindComponent(t.instance, AName))
}

func (t *TToolBar) GetNamePath() string {
    return ToolBar_GetNamePath(t.instance)
}

func (t *TToolBar) Assign(Source IObject) {
    ToolBar_Assign(t.instance, CheckPtr(Source))
}

func (t *TToolBar) ClassName() string {
    return ToolBar_ClassName(t.instance)
}

func (t *TToolBar) Equals(Obj IObject) bool {
    return ToolBar_Equals(t.instance, CheckPtr(Obj))
}

func (t *TToolBar) GetHashCode() int32 {
    return ToolBar_GetHashCode(t.instance)
}

func (t *TToolBar) ToString() string {
    return ToolBar_ToString(t.instance)
}

func (t *TToolBar) ButtonCount() int32 {
    return ToolBar_GetButtonCount(t.instance)
}

func (t *TToolBar) Canvas() *TCanvas {
    return CanvasFromInst(ToolBar_GetCanvas(t.instance))
}

func (t *TToolBar) CustomizeKeyName() string {
    return ToolBar_GetCustomizeKeyName(t.instance)
}

func (t *TToolBar) SetCustomizeKeyName(value string) {
    ToolBar_SetCustomizeKeyName(t.instance, value)
}

func (t *TToolBar) CustomizeValueName() string {
    return ToolBar_GetCustomizeValueName(t.instance)
}

func (t *TToolBar) SetCustomizeValueName(value string) {
    ToolBar_SetCustomizeValueName(t.instance, value)
}

func (t *TToolBar) RowCount() int32 {
    return ToolBar_GetRowCount(t.instance)
}

func (t *TToolBar) Align() TAlign {
    return ToolBar_GetAlign(t.instance)
}

func (t *TToolBar) SetAlign(value TAlign) {
    ToolBar_SetAlign(t.instance, value)
}

func (t *TToolBar) Anchors() TAnchors {
    return ToolBar_GetAnchors(t.instance)
}

func (t *TToolBar) SetAnchors(value TAnchors) {
    ToolBar_SetAnchors(t.instance, value)
}

func (t *TToolBar) AutoSize() bool {
    return ToolBar_GetAutoSize(t.instance)
}

func (t *TToolBar) SetAutoSize(value bool) {
    ToolBar_SetAutoSize(t.instance, value)
}

func (t *TToolBar) BorderWidth() int32 {
    return ToolBar_GetBorderWidth(t.instance)
}

func (t *TToolBar) SetBorderWidth(value int32) {
    ToolBar_SetBorderWidth(t.instance, value)
}

func (t *TToolBar) ButtonHeight() int32 {
    return ToolBar_GetButtonHeight(t.instance)
}

func (t *TToolBar) SetButtonHeight(value int32) {
    ToolBar_SetButtonHeight(t.instance, value)
}

func (t *TToolBar) ButtonWidth() int32 {
    return ToolBar_GetButtonWidth(t.instance)
}

func (t *TToolBar) SetButtonWidth(value int32) {
    ToolBar_SetButtonWidth(t.instance, value)
}

func (t *TToolBar) Caption() string {
    return ToolBar_GetCaption(t.instance)
}

func (t *TToolBar) SetCaption(value string) {
    ToolBar_SetCaption(t.instance, value)
}

func (t *TToolBar) Color() TColor {
    return ToolBar_GetColor(t.instance)
}

func (t *TToolBar) SetColor(value TColor) {
    ToolBar_SetColor(t.instance, value)
}

func (t *TToolBar) DoubleBuffered() bool {
    return ToolBar_GetDoubleBuffered(t.instance)
}

func (t *TToolBar) SetDoubleBuffered(value bool) {
    ToolBar_SetDoubleBuffered(t.instance, value)
}

func (t *TToolBar) DrawingStyle() TTBDrawingStyle {
    return ToolBar_GetDrawingStyle(t.instance)
}

func (t *TToolBar) SetDrawingStyle(value TTBDrawingStyle) {
    ToolBar_SetDrawingStyle(t.instance, value)
}

func (t *TToolBar) Enabled() bool {
    return ToolBar_GetEnabled(t.instance)
}

func (t *TToolBar) SetEnabled(value bool) {
    ToolBar_SetEnabled(t.instance, value)
}

func (t *TToolBar) Flat() bool {
    return ToolBar_GetFlat(t.instance)
}

func (t *TToolBar) SetFlat(value bool) {
    ToolBar_SetFlat(t.instance, value)
}

func (t *TToolBar) Font() *TFont {
    return FontFromInst(ToolBar_GetFont(t.instance))
}

func (t *TToolBar) SetFont(value *TFont) {
    ToolBar_SetFont(t.instance, CheckPtr(value))
}

func (t *TToolBar) GradientEndColor() TColor {
    return ToolBar_GetGradientEndColor(t.instance)
}

func (t *TToolBar) SetGradientEndColor(value TColor) {
    ToolBar_SetGradientEndColor(t.instance, value)
}

func (t *TToolBar) GradientStartColor() TColor {
    return ToolBar_GetGradientStartColor(t.instance)
}

func (t *TToolBar) SetGradientStartColor(value TColor) {
    ToolBar_SetGradientStartColor(t.instance, value)
}

func (t *TToolBar) Height() int32 {
    return ToolBar_GetHeight(t.instance)
}

func (t *TToolBar) SetHeight(value int32) {
    ToolBar_SetHeight(t.instance, value)
}

func (t *TToolBar) HideClippedButtons() bool {
    return ToolBar_GetHideClippedButtons(t.instance)
}

func (t *TToolBar) SetHideClippedButtons(value bool) {
    ToolBar_SetHideClippedButtons(t.instance, value)
}

func (t *TToolBar) HotImages() *TImageList {
    return ImageListFromInst(ToolBar_GetHotImages(t.instance))
}

func (t *TToolBar) SetHotImages(value IComponent) {
    ToolBar_SetHotImages(t.instance, CheckPtr(value))
}

func (t *TToolBar) Images() *TImageList {
    return ImageListFromInst(ToolBar_GetImages(t.instance))
}

func (t *TToolBar) SetImages(value IComponent) {
    ToolBar_SetImages(t.instance, CheckPtr(value))
}

func (t *TToolBar) Indent() int32 {
    return ToolBar_GetIndent(t.instance)
}

func (t *TToolBar) SetIndent(value int32) {
    ToolBar_SetIndent(t.instance, value)
}

func (t *TToolBar) List() bool {
    return ToolBar_GetList(t.instance)
}

func (t *TToolBar) SetList(value bool) {
    ToolBar_SetList(t.instance, value)
}

func (t *TToolBar) Menu() *TMainMenu {
    return MainMenuFromInst(ToolBar_GetMenu(t.instance))
}

func (t *TToolBar) SetMenu(value IComponent) {
    ToolBar_SetMenu(t.instance, CheckPtr(value))
}

func (t *TToolBar) GradientDirection() TGradientDirection {
    return ToolBar_GetGradientDirection(t.instance)
}

func (t *TToolBar) SetGradientDirection(value TGradientDirection) {
    ToolBar_SetGradientDirection(t.instance, value)
}

func (t *TToolBar) GradientDrawingOptions() TTBGradientDrawingOptions {
    return ToolBar_GetGradientDrawingOptions(t.instance)
}

func (t *TToolBar) SetGradientDrawingOptions(value TTBGradientDrawingOptions) {
    ToolBar_SetGradientDrawingOptions(t.instance, value)
}

func (t *TToolBar) ParentColor() bool {
    return ToolBar_GetParentColor(t.instance)
}

func (t *TToolBar) SetParentColor(value bool) {
    ToolBar_SetParentColor(t.instance, value)
}

func (t *TToolBar) ParentDoubleBuffered() bool {
    return ToolBar_GetParentDoubleBuffered(t.instance)
}

func (t *TToolBar) SetParentDoubleBuffered(value bool) {
    ToolBar_SetParentDoubleBuffered(t.instance, value)
}

func (t *TToolBar) ParentFont() bool {
    return ToolBar_GetParentFont(t.instance)
}

func (t *TToolBar) SetParentFont(value bool) {
    ToolBar_SetParentFont(t.instance, value)
}

func (t *TToolBar) ParentShowHint() bool {
    return ToolBar_GetParentShowHint(t.instance)
}

func (t *TToolBar) SetParentShowHint(value bool) {
    ToolBar_SetParentShowHint(t.instance, value)
}

func (t *TToolBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ToolBar_GetPopupMenu(t.instance))
}

func (t *TToolBar) SetPopupMenu(value IComponent) {
    ToolBar_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TToolBar) ShowCaptions() bool {
    return ToolBar_GetShowCaptions(t.instance)
}

func (t *TToolBar) SetShowCaptions(value bool) {
    ToolBar_SetShowCaptions(t.instance, value)
}

func (t *TToolBar) ShowHint() bool {
    return ToolBar_GetShowHint(t.instance)
}

func (t *TToolBar) SetShowHint(value bool) {
    ToolBar_SetShowHint(t.instance, value)
}

func (t *TToolBar) TabOrder() uint16 {
    return ToolBar_GetTabOrder(t.instance)
}

func (t *TToolBar) SetTabOrder(value uint16) {
    ToolBar_SetTabOrder(t.instance, value)
}

func (t *TToolBar) TabStop() bool {
    return ToolBar_GetTabStop(t.instance)
}

func (t *TToolBar) SetTabStop(value bool) {
    ToolBar_SetTabStop(t.instance, value)
}

func (t *TToolBar) Transparent() bool {
    return ToolBar_GetTransparent(t.instance)
}

func (t *TToolBar) SetTransparent(value bool) {
    ToolBar_SetTransparent(t.instance, value)
}

func (t *TToolBar) Visible() bool {
    return ToolBar_GetVisible(t.instance)
}

func (t *TToolBar) SetVisible(value bool) {
    ToolBar_SetVisible(t.instance, value)
}

func (t *TToolBar) StyleElements() TStyleElements {
    return ToolBar_GetStyleElements(t.instance)
}

func (t *TToolBar) SetStyleElements(value TStyleElements) {
    ToolBar_SetStyleElements(t.instance, value)
}

func (t *TToolBar) Wrapable() bool {
    return ToolBar_GetWrapable(t.instance)
}

func (t *TToolBar) SetWrapable(value bool) {
    ToolBar_SetWrapable(t.instance, value)
}

func (t *TToolBar) SetOnAdvancedCustomDraw(fn TTBAdvancedCustomDrawEvent) {
    ToolBar_SetOnAdvancedCustomDraw(t.instance, fn)
}

func (t *TToolBar) SetOnAdvancedCustomDrawButton(fn TTBAdvancedCustomDrawBtnEvent) {
    ToolBar_SetOnAdvancedCustomDrawButton(t.instance, fn)
}

func (t *TToolBar) SetOnClick(fn TNotifyEvent) {
    ToolBar_SetOnClick(t.instance, fn)
}

func (t *TToolBar) SetOnDblClick(fn TNotifyEvent) {
    ToolBar_SetOnDblClick(t.instance, fn)
}

func (t *TToolBar) SetOnEnter(fn TNotifyEvent) {
    ToolBar_SetOnEnter(t.instance, fn)
}

func (t *TToolBar) SetOnExit(fn TNotifyEvent) {
    ToolBar_SetOnExit(t.instance, fn)
}

func (t *TToolBar) SetOnMouseDown(fn TMouseEvent) {
    ToolBar_SetOnMouseDown(t.instance, fn)
}

func (t *TToolBar) SetOnMouseEnter(fn TNotifyEvent) {
    ToolBar_SetOnMouseEnter(t.instance, fn)
}

func (t *TToolBar) SetOnMouseLeave(fn TNotifyEvent) {
    ToolBar_SetOnMouseLeave(t.instance, fn)
}

func (t *TToolBar) SetOnMouseMove(fn TMouseMoveEvent) {
    ToolBar_SetOnMouseMove(t.instance, fn)
}

func (t *TToolBar) SetOnMouseUp(fn TMouseEvent) {
    ToolBar_SetOnMouseUp(t.instance, fn)
}

func (t *TToolBar) SetOnResize(fn TNotifyEvent) {
    ToolBar_SetOnResize(t.instance, fn)
}

func (t *TToolBar) Brush() *TBrush {
    return BrushFromInst(ToolBar_GetBrush(t.instance))
}

func (t *TToolBar) ControlCount() int32 {
    return ToolBar_GetControlCount(t.instance)
}

func (t *TToolBar) Handle() HWND {
    return ToolBar_GetHandle(t.instance)
}

func (t *TToolBar) ParentWindow() HWND {
    return ToolBar_GetParentWindow(t.instance)
}

func (t *TToolBar) SetParentWindow(value HWND) {
    ToolBar_SetParentWindow(t.instance, value)
}

func (t *TToolBar) Action() *TAction {
    return ActionFromInst(ToolBar_GetAction(t.instance))
}

func (t *TToolBar) SetAction(value IComponent) {
    ToolBar_SetAction(t.instance, CheckPtr(value))
}

func (t *TToolBar) BiDiMode() TBiDiMode {
    return ToolBar_GetBiDiMode(t.instance)
}

func (t *TToolBar) SetBiDiMode(value TBiDiMode) {
    ToolBar_SetBiDiMode(t.instance, value)
}

func (t *TToolBar) BoundsRect() TRect {
    return ToolBar_GetBoundsRect(t.instance)
}

func (t *TToolBar) SetBoundsRect(value TRect) {
    ToolBar_SetBoundsRect(t.instance, value)
}

func (t *TToolBar) ClientHeight() int32 {
    return ToolBar_GetClientHeight(t.instance)
}

func (t *TToolBar) SetClientHeight(value int32) {
    ToolBar_SetClientHeight(t.instance, value)
}

func (t *TToolBar) ClientRect() TRect {
    return ToolBar_GetClientRect(t.instance)
}

func (t *TToolBar) ClientWidth() int32 {
    return ToolBar_GetClientWidth(t.instance)
}

func (t *TToolBar) SetClientWidth(value int32) {
    ToolBar_SetClientWidth(t.instance, value)
}

func (t *TToolBar) ExplicitLeft() int32 {
    return ToolBar_GetExplicitLeft(t.instance)
}

func (t *TToolBar) ExplicitTop() int32 {
    return ToolBar_GetExplicitTop(t.instance)
}

func (t *TToolBar) ExplicitWidth() int32 {
    return ToolBar_GetExplicitWidth(t.instance)
}

func (t *TToolBar) ExplicitHeight() int32 {
    return ToolBar_GetExplicitHeight(t.instance)
}

func (t *TToolBar) Parent() *TControl {
    return ControlFromInst(ToolBar_GetParent(t.instance))
}

func (t *TToolBar) SetParent(value IControl) {
    ToolBar_SetParent(t.instance, CheckPtr(value))
}

func (t *TToolBar) AlignWithMargins() bool {
    return ToolBar_GetAlignWithMargins(t.instance)
}

func (t *TToolBar) SetAlignWithMargins(value bool) {
    ToolBar_SetAlignWithMargins(t.instance, value)
}

func (t *TToolBar) Left() int32 {
    return ToolBar_GetLeft(t.instance)
}

func (t *TToolBar) SetLeft(value int32) {
    ToolBar_SetLeft(t.instance, value)
}

func (t *TToolBar) Top() int32 {
    return ToolBar_GetTop(t.instance)
}

func (t *TToolBar) SetTop(value int32) {
    ToolBar_SetTop(t.instance, value)
}

func (t *TToolBar) Width() int32 {
    return ToolBar_GetWidth(t.instance)
}

func (t *TToolBar) SetWidth(value int32) {
    ToolBar_SetWidth(t.instance, value)
}

func (t *TToolBar) Cursor() TCursor {
    return ToolBar_GetCursor(t.instance)
}

func (t *TToolBar) SetCursor(value TCursor) {
    ToolBar_SetCursor(t.instance, value)
}

func (t *TToolBar) Hint() string {
    return ToolBar_GetHint(t.instance)
}

func (t *TToolBar) SetHint(value string) {
    ToolBar_SetHint(t.instance, value)
}

func (t *TToolBar) Margins() *TMargins {
    return MarginsFromInst(ToolBar_GetMargins(t.instance))
}

func (t *TToolBar) SetMargins(value *TMargins) {
    ToolBar_SetMargins(t.instance, CheckPtr(value))
}

func (t *TToolBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(ToolBar_GetCustomHint(t.instance))
}

func (t *TToolBar) SetCustomHint(value IComponent) {
    ToolBar_SetCustomHint(t.instance, CheckPtr(value))
}

func (t *TToolBar) ComponentCount() int32 {
    return ToolBar_GetComponentCount(t.instance)
}

func (t *TToolBar) ComponentIndex() int32 {
    return ToolBar_GetComponentIndex(t.instance)
}

func (t *TToolBar) SetComponentIndex(value int32) {
    ToolBar_SetComponentIndex(t.instance, value)
}

func (t *TToolBar) Owner() *TComponent {
    return ComponentFromInst(ToolBar_GetOwner(t.instance))
}

func (t *TToolBar) Name() string {
    return ToolBar_GetName(t.instance)
}

func (t *TToolBar) SetName(value string) {
    ToolBar_SetName(t.instance, value)
}

func (t *TToolBar) Tag() int {
    return ToolBar_GetTag(t.instance)
}

func (t *TToolBar) SetTag(value int) {
    ToolBar_SetTag(t.instance, value)
}

func (t *TToolBar) Buttons(Index int32) *TToolButton {
    return ToolButtonFromInst(ToolBar_GetButtons(t.instance, Index))
}

func (t *TToolBar) Controls(Index int32) *TControl {
    return ControlFromInst(ToolBar_GetControls(t.instance, Index))
}

func (t *TToolBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ToolBar_GetComponents(t.instance, AIndex))
}

