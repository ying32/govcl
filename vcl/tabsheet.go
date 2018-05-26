
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

type TTabSheet struct {
    IControl
    instance uintptr
}

func NewTabSheet(owner IComponent) *TTabSheet {
    t := new(TTabSheet)
    t.instance = TabSheet_Create(CheckPtr(owner))
    return t
}

func TabSheetFromInst(inst uintptr) *TTabSheet {
    t := new(TTabSheet)
    t.instance = inst
    return t
}

func TabSheetFromObj(obj IObject) *TTabSheet {
    t := new(TTabSheet)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTabSheet) Free() {
    if t.instance != 0 {
        TabSheet_Free(t.instance)
        t.instance = 0
    }
}

func (t *TTabSheet) Instance() uintptr {
    return t.instance
}

func (t *TTabSheet) IsValid() bool {
    return t.instance != 0
}

func (t *TTabSheet) CanFocus() bool {
    return TabSheet_CanFocus(t.instance)
}

func (t *TTabSheet) FlipChildren(AllLevels bool) {
    TabSheet_FlipChildren(t.instance, AllLevels)
}

func (t *TTabSheet) Focused() bool {
    return TabSheet_Focused(t.instance)
}

func (t *TTabSheet) HandleAllocated() bool {
    return TabSheet_HandleAllocated(t.instance)
}

func (t *TTabSheet) Invalidate() {
    TabSheet_Invalidate(t.instance)
}

func (t *TTabSheet) Realign() {
    TabSheet_Realign(t.instance)
}

func (t *TTabSheet) Repaint() {
    TabSheet_Repaint(t.instance)
}

func (t *TTabSheet) ScaleBy(M int32, D int32) {
    TabSheet_ScaleBy(t.instance, M , D)
}

func (t *TTabSheet) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    TabSheet_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

func (t *TTabSheet) SetFocus() {
    TabSheet_SetFocus(t.instance)
}

func (t *TTabSheet) Update() {
    TabSheet_Update(t.instance)
}

func (t *TTabSheet) BringToFront() {
    TabSheet_BringToFront(t.instance)
}

func (t *TTabSheet) HasParent() bool {
    return TabSheet_HasParent(t.instance)
}

func (t *TTabSheet) Hide() {
    TabSheet_Hide(t.instance)
}

func (t *TTabSheet) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return TabSheet_Perform(t.instance, Msg , WParam , LParam)
}

func (t *TTabSheet) Refresh() {
    TabSheet_Refresh(t.instance)
}

func (t *TTabSheet) SendToBack() {
    TabSheet_SendToBack(t.instance)
}

func (t *TTabSheet) Show() {
    TabSheet_Show(t.instance)
}

func (t *TTabSheet) GetTextBuf(Buffer string, BufSize int32) int32 {
    return TabSheet_GetTextBuf(t.instance, Buffer , BufSize)
}

func (t *TTabSheet) GetTextLen() int32 {
    return TabSheet_GetTextLen(t.instance)
}

func (t *TTabSheet) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TabSheet_FindComponent(t.instance, AName))
}

func (t *TTabSheet) GetNamePath() string {
    return TabSheet_GetNamePath(t.instance)
}

func (t *TTabSheet) Assign(Source IObject) {
    TabSheet_Assign(t.instance, CheckPtr(Source))
}

func (t *TTabSheet) ClassName() string {
    return TabSheet_ClassName(t.instance)
}

func (t *TTabSheet) Equals(Obj IObject) bool {
    return TabSheet_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTabSheet) GetHashCode() int32 {
    return TabSheet_GetHashCode(t.instance)
}

func (t *TTabSheet) ToString() string {
    return TabSheet_ToString(t.instance)
}

func (t *TTabSheet) PageControl() *TPageControl {
    return PageControlFromInst(TabSheet_GetPageControl(t.instance))
}

func (t *TTabSheet) SetPageControl(value IControl) {
    TabSheet_SetPageControl(t.instance, CheckPtr(value))
}

func (t *TTabSheet) TabIndex() int32 {
    return TabSheet_GetTabIndex(t.instance)
}

func (t *TTabSheet) BorderWidth() int32 {
    return TabSheet_GetBorderWidth(t.instance)
}

func (t *TTabSheet) SetBorderWidth(value int32) {
    TabSheet_SetBorderWidth(t.instance, value)
}

func (t *TTabSheet) Caption() string {
    return TabSheet_GetCaption(t.instance)
}

func (t *TTabSheet) SetCaption(value string) {
    TabSheet_SetCaption(t.instance, value)
}

func (t *TTabSheet) DoubleBuffered() bool {
    return TabSheet_GetDoubleBuffered(t.instance)
}

func (t *TTabSheet) SetDoubleBuffered(value bool) {
    TabSheet_SetDoubleBuffered(t.instance, value)
}

func (t *TTabSheet) Enabled() bool {
    return TabSheet_GetEnabled(t.instance)
}

func (t *TTabSheet) SetEnabled(value bool) {
    TabSheet_SetEnabled(t.instance, value)
}

func (t *TTabSheet) Font() *TFont {
    return FontFromInst(TabSheet_GetFont(t.instance))
}

func (t *TTabSheet) SetFont(value *TFont) {
    TabSheet_SetFont(t.instance, CheckPtr(value))
}

func (t *TTabSheet) Height() int32 {
    return TabSheet_GetHeight(t.instance)
}

func (t *TTabSheet) SetHeight(value int32) {
    TabSheet_SetHeight(t.instance, value)
}

func (t *TTabSheet) Highlighted() bool {
    return TabSheet_GetHighlighted(t.instance)
}

func (t *TTabSheet) SetHighlighted(value bool) {
    TabSheet_SetHighlighted(t.instance, value)
}

func (t *TTabSheet) ImageIndex() int32 {
    return TabSheet_GetImageIndex(t.instance)
}

func (t *TTabSheet) SetImageIndex(value int32) {
    TabSheet_SetImageIndex(t.instance, value)
}

func (t *TTabSheet) Left() int32 {
    return TabSheet_GetLeft(t.instance)
}

func (t *TTabSheet) SetLeft(value int32) {
    TabSheet_SetLeft(t.instance, value)
}

func (t *TTabSheet) PageIndex() int32 {
    return TabSheet_GetPageIndex(t.instance)
}

func (t *TTabSheet) SetPageIndex(value int32) {
    TabSheet_SetPageIndex(t.instance, value)
}

func (t *TTabSheet) ParentDoubleBuffered() bool {
    return TabSheet_GetParentDoubleBuffered(t.instance)
}

func (t *TTabSheet) SetParentDoubleBuffered(value bool) {
    TabSheet_SetParentDoubleBuffered(t.instance, value)
}

func (t *TTabSheet) ParentFont() bool {
    return TabSheet_GetParentFont(t.instance)
}

func (t *TTabSheet) SetParentFont(value bool) {
    TabSheet_SetParentFont(t.instance, value)
}

func (t *TTabSheet) ParentShowHint() bool {
    return TabSheet_GetParentShowHint(t.instance)
}

func (t *TTabSheet) SetParentShowHint(value bool) {
    TabSheet_SetParentShowHint(t.instance, value)
}

func (t *TTabSheet) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TabSheet_GetPopupMenu(t.instance))
}

func (t *TTabSheet) SetPopupMenu(value IComponent) {
    TabSheet_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TTabSheet) ShowHint() bool {
    return TabSheet_GetShowHint(t.instance)
}

func (t *TTabSheet) SetShowHint(value bool) {
    TabSheet_SetShowHint(t.instance, value)
}

func (t *TTabSheet) TabVisible() bool {
    return TabSheet_GetTabVisible(t.instance)
}

func (t *TTabSheet) SetTabVisible(value bool) {
    TabSheet_SetTabVisible(t.instance, value)
}

func (t *TTabSheet) Top() int32 {
    return TabSheet_GetTop(t.instance)
}

func (t *TTabSheet) SetTop(value int32) {
    TabSheet_SetTop(t.instance, value)
}

func (t *TTabSheet) Visible() bool {
    return TabSheet_GetVisible(t.instance)
}

func (t *TTabSheet) SetVisible(value bool) {
    TabSheet_SetVisible(t.instance, value)
}

func (t *TTabSheet) Width() int32 {
    return TabSheet_GetWidth(t.instance)
}

func (t *TTabSheet) SetWidth(value int32) {
    TabSheet_SetWidth(t.instance, value)
}

func (t *TTabSheet) SetOnEnter(fn TNotifyEvent) {
    TabSheet_SetOnEnter(t.instance, fn)
}

func (t *TTabSheet) SetOnExit(fn TNotifyEvent) {
    TabSheet_SetOnExit(t.instance, fn)
}

func (t *TTabSheet) SetOnHide(fn TNotifyEvent) {
    TabSheet_SetOnHide(t.instance, fn)
}

func (t *TTabSheet) SetOnMouseDown(fn TMouseEvent) {
    TabSheet_SetOnMouseDown(t.instance, fn)
}

func (t *TTabSheet) SetOnMouseEnter(fn TNotifyEvent) {
    TabSheet_SetOnMouseEnter(t.instance, fn)
}

func (t *TTabSheet) SetOnMouseLeave(fn TNotifyEvent) {
    TabSheet_SetOnMouseLeave(t.instance, fn)
}

func (t *TTabSheet) SetOnMouseMove(fn TMouseMoveEvent) {
    TabSheet_SetOnMouseMove(t.instance, fn)
}

func (t *TTabSheet) SetOnMouseUp(fn TMouseEvent) {
    TabSheet_SetOnMouseUp(t.instance, fn)
}

func (t *TTabSheet) SetOnResize(fn TNotifyEvent) {
    TabSheet_SetOnResize(t.instance, fn)
}

func (t *TTabSheet) SetOnShow(fn TNotifyEvent) {
    TabSheet_SetOnShow(t.instance, fn)
}

func (t *TTabSheet) Brush() *TBrush {
    return BrushFromInst(TabSheet_GetBrush(t.instance))
}

func (t *TTabSheet) ControlCount() int32 {
    return TabSheet_GetControlCount(t.instance)
}

func (t *TTabSheet) Handle() HWND {
    return TabSheet_GetHandle(t.instance)
}

func (t *TTabSheet) ParentWindow() HWND {
    return TabSheet_GetParentWindow(t.instance)
}

func (t *TTabSheet) SetParentWindow(value HWND) {
    TabSheet_SetParentWindow(t.instance, value)
}

func (t *TTabSheet) TabOrder() uint16 {
    return TabSheet_GetTabOrder(t.instance)
}

func (t *TTabSheet) SetTabOrder(value uint16) {
    TabSheet_SetTabOrder(t.instance, value)
}

func (t *TTabSheet) TabStop() bool {
    return TabSheet_GetTabStop(t.instance)
}

func (t *TTabSheet) SetTabStop(value bool) {
    TabSheet_SetTabStop(t.instance, value)
}

func (t *TTabSheet) Action() *TAction {
    return ActionFromInst(TabSheet_GetAction(t.instance))
}

func (t *TTabSheet) SetAction(value IComponent) {
    TabSheet_SetAction(t.instance, CheckPtr(value))
}

func (t *TTabSheet) Align() TAlign {
    return TabSheet_GetAlign(t.instance)
}

func (t *TTabSheet) SetAlign(value TAlign) {
    TabSheet_SetAlign(t.instance, value)
}

func (t *TTabSheet) Anchors() TAnchors {
    return TabSheet_GetAnchors(t.instance)
}

func (t *TTabSheet) SetAnchors(value TAnchors) {
    TabSheet_SetAnchors(t.instance, value)
}

func (t *TTabSheet) BiDiMode() TBiDiMode {
    return TabSheet_GetBiDiMode(t.instance)
}

func (t *TTabSheet) SetBiDiMode(value TBiDiMode) {
    TabSheet_SetBiDiMode(t.instance, value)
}

func (t *TTabSheet) BoundsRect() TRect {
    return TabSheet_GetBoundsRect(t.instance)
}

func (t *TTabSheet) SetBoundsRect(value TRect) {
    TabSheet_SetBoundsRect(t.instance, value)
}

func (t *TTabSheet) ClientHeight() int32 {
    return TabSheet_GetClientHeight(t.instance)
}

func (t *TTabSheet) SetClientHeight(value int32) {
    TabSheet_SetClientHeight(t.instance, value)
}

func (t *TTabSheet) ClientRect() TRect {
    return TabSheet_GetClientRect(t.instance)
}

func (t *TTabSheet) ClientWidth() int32 {
    return TabSheet_GetClientWidth(t.instance)
}

func (t *TTabSheet) SetClientWidth(value int32) {
    TabSheet_SetClientWidth(t.instance, value)
}

func (t *TTabSheet) ExplicitLeft() int32 {
    return TabSheet_GetExplicitLeft(t.instance)
}

func (t *TTabSheet) ExplicitTop() int32 {
    return TabSheet_GetExplicitTop(t.instance)
}

func (t *TTabSheet) ExplicitWidth() int32 {
    return TabSheet_GetExplicitWidth(t.instance)
}

func (t *TTabSheet) ExplicitHeight() int32 {
    return TabSheet_GetExplicitHeight(t.instance)
}

func (t *TTabSheet) Parent() *TControl {
    return ControlFromInst(TabSheet_GetParent(t.instance))
}

func (t *TTabSheet) SetParent(value IControl) {
    TabSheet_SetParent(t.instance, CheckPtr(value))
}

func (t *TTabSheet) StyleElements() TStyleElements {
    return TabSheet_GetStyleElements(t.instance)
}

func (t *TTabSheet) SetStyleElements(value TStyleElements) {
    TabSheet_SetStyleElements(t.instance, value)
}

func (t *TTabSheet) AlignWithMargins() bool {
    return TabSheet_GetAlignWithMargins(t.instance)
}

func (t *TTabSheet) SetAlignWithMargins(value bool) {
    TabSheet_SetAlignWithMargins(t.instance, value)
}

func (t *TTabSheet) Cursor() TCursor {
    return TabSheet_GetCursor(t.instance)
}

func (t *TTabSheet) SetCursor(value TCursor) {
    TabSheet_SetCursor(t.instance, value)
}

func (t *TTabSheet) Hint() string {
    return TabSheet_GetHint(t.instance)
}

func (t *TTabSheet) SetHint(value string) {
    TabSheet_SetHint(t.instance, value)
}

func (t *TTabSheet) Margins() *TMargins {
    return MarginsFromInst(TabSheet_GetMargins(t.instance))
}

func (t *TTabSheet) SetMargins(value *TMargins) {
    TabSheet_SetMargins(t.instance, CheckPtr(value))
}

func (t *TTabSheet) CustomHint() *TCustomHint {
    return CustomHintFromInst(TabSheet_GetCustomHint(t.instance))
}

func (t *TTabSheet) SetCustomHint(value IComponent) {
    TabSheet_SetCustomHint(t.instance, CheckPtr(value))
}

func (t *TTabSheet) ComponentCount() int32 {
    return TabSheet_GetComponentCount(t.instance)
}

func (t *TTabSheet) ComponentIndex() int32 {
    return TabSheet_GetComponentIndex(t.instance)
}

func (t *TTabSheet) SetComponentIndex(value int32) {
    TabSheet_SetComponentIndex(t.instance, value)
}

func (t *TTabSheet) Owner() *TComponent {
    return ComponentFromInst(TabSheet_GetOwner(t.instance))
}

func (t *TTabSheet) Name() string {
    return TabSheet_GetName(t.instance)
}

func (t *TTabSheet) SetName(value string) {
    TabSheet_SetName(t.instance, value)
}

func (t *TTabSheet) Tag() int {
    return TabSheet_GetTag(t.instance)
}

func (t *TTabSheet) SetTag(value int) {
    TabSheet_SetTag(t.instance, value)
}

func (t *TTabSheet) Controls(Index int32) *TControl {
    return ControlFromInst(TabSheet_GetControls(t.instance, Index))
}

func (t *TTabSheet) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TabSheet_GetComponents(t.instance, AIndex))
}

