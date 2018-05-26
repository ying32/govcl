
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

type TPageControl struct {
    IControl
    instance uintptr
}

func NewPageControl(owner IComponent) *TPageControl {
    p := new(TPageControl)
    p.instance = PageControl_Create(CheckPtr(owner))
    return p
}

func PageControlFromInst(inst uintptr) *TPageControl {
    p := new(TPageControl)
    p.instance = inst
    return p
}

func PageControlFromObj(obj IObject) *TPageControl {
    p := new(TPageControl)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPageControl) Free() {
    if p.instance != 0 {
        PageControl_Free(p.instance)
        p.instance = 0
    }
}

func (p *TPageControl) Instance() uintptr {
    return p.instance
}

func (p *TPageControl) IsValid() bool {
    return p.instance != 0
}

func (p *TPageControl) SelectNextPage(GoForward bool, CheckTabVisible bool) {
    PageControl_SelectNextPage(p.instance, GoForward , CheckTabVisible)
}

func (p *TPageControl) RowCount() int32 {
    return PageControl_RowCount(p.instance)
}

func (p *TPageControl) CanFocus() bool {
    return PageControl_CanFocus(p.instance)
}

func (p *TPageControl) FlipChildren(AllLevels bool) {
    PageControl_FlipChildren(p.instance, AllLevels)
}

func (p *TPageControl) Focused() bool {
    return PageControl_Focused(p.instance)
}

func (p *TPageControl) HandleAllocated() bool {
    return PageControl_HandleAllocated(p.instance)
}

func (p *TPageControl) Invalidate() {
    PageControl_Invalidate(p.instance)
}

func (p *TPageControl) Realign() {
    PageControl_Realign(p.instance)
}

func (p *TPageControl) Repaint() {
    PageControl_Repaint(p.instance)
}

func (p *TPageControl) ScaleBy(M int32, D int32) {
    PageControl_ScaleBy(p.instance, M , D)
}

func (p *TPageControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    PageControl_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

func (p *TPageControl) SetFocus() {
    PageControl_SetFocus(p.instance)
}

func (p *TPageControl) Update() {
    PageControl_Update(p.instance)
}

func (p *TPageControl) BringToFront() {
    PageControl_BringToFront(p.instance)
}

func (p *TPageControl) HasParent() bool {
    return PageControl_HasParent(p.instance)
}

func (p *TPageControl) Hide() {
    PageControl_Hide(p.instance)
}

func (p *TPageControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return PageControl_Perform(p.instance, Msg , WParam , LParam)
}

func (p *TPageControl) Refresh() {
    PageControl_Refresh(p.instance)
}

func (p *TPageControl) SendToBack() {
    PageControl_SendToBack(p.instance)
}

func (p *TPageControl) Show() {
    PageControl_Show(p.instance)
}

func (p *TPageControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    return PageControl_GetTextBuf(p.instance, Buffer , BufSize)
}

func (p *TPageControl) GetTextLen() int32 {
    return PageControl_GetTextLen(p.instance)
}

func (p *TPageControl) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PageControl_FindComponent(p.instance, AName))
}

func (p *TPageControl) GetNamePath() string {
    return PageControl_GetNamePath(p.instance)
}

func (p *TPageControl) Assign(Source IObject) {
    PageControl_Assign(p.instance, CheckPtr(Source))
}

func (p *TPageControl) ClassName() string {
    return PageControl_ClassName(p.instance)
}

func (p *TPageControl) Equals(Obj IObject) bool {
    return PageControl_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPageControl) GetHashCode() int32 {
    return PageControl_GetHashCode(p.instance)
}

func (p *TPageControl) ToString() string {
    return PageControl_ToString(p.instance)
}

func (p *TPageControl) ActivePageIndex() int32 {
    return PageControl_GetActivePageIndex(p.instance)
}

func (p *TPageControl) SetActivePageIndex(value int32) {
    PageControl_SetActivePageIndex(p.instance, value)
}

func (p *TPageControl) PageCount() int32 {
    return PageControl_GetPageCount(p.instance)
}

func (p *TPageControl) Align() TAlign {
    return PageControl_GetAlign(p.instance)
}

func (p *TPageControl) SetAlign(value TAlign) {
    PageControl_SetAlign(p.instance, value)
}

func (p *TPageControl) Anchors() TAnchors {
    return PageControl_GetAnchors(p.instance)
}

func (p *TPageControl) SetAnchors(value TAnchors) {
    PageControl_SetAnchors(p.instance, value)
}

func (p *TPageControl) BiDiMode() TBiDiMode {
    return PageControl_GetBiDiMode(p.instance)
}

func (p *TPageControl) SetBiDiMode(value TBiDiMode) {
    PageControl_SetBiDiMode(p.instance, value)
}

func (p *TPageControl) DoubleBuffered() bool {
    return PageControl_GetDoubleBuffered(p.instance)
}

func (p *TPageControl) SetDoubleBuffered(value bool) {
    PageControl_SetDoubleBuffered(p.instance, value)
}

func (p *TPageControl) Enabled() bool {
    return PageControl_GetEnabled(p.instance)
}

func (p *TPageControl) SetEnabled(value bool) {
    PageControl_SetEnabled(p.instance, value)
}

func (p *TPageControl) Font() *TFont {
    return FontFromInst(PageControl_GetFont(p.instance))
}

func (p *TPageControl) SetFont(value *TFont) {
    PageControl_SetFont(p.instance, CheckPtr(value))
}

func (p *TPageControl) HotTrack() bool {
    return PageControl_GetHotTrack(p.instance)
}

func (p *TPageControl) SetHotTrack(value bool) {
    PageControl_SetHotTrack(p.instance, value)
}

func (p *TPageControl) Images() *TImageList {
    return ImageListFromInst(PageControl_GetImages(p.instance))
}

func (p *TPageControl) SetImages(value IComponent) {
    PageControl_SetImages(p.instance, CheckPtr(value))
}

func (p *TPageControl) MultiLine() bool {
    return PageControl_GetMultiLine(p.instance)
}

func (p *TPageControl) SetMultiLine(value bool) {
    PageControl_SetMultiLine(p.instance, value)
}

func (p *TPageControl) ParentDoubleBuffered() bool {
    return PageControl_GetParentDoubleBuffered(p.instance)
}

func (p *TPageControl) SetParentDoubleBuffered(value bool) {
    PageControl_SetParentDoubleBuffered(p.instance, value)
}

func (p *TPageControl) ParentFont() bool {
    return PageControl_GetParentFont(p.instance)
}

func (p *TPageControl) SetParentFont(value bool) {
    PageControl_SetParentFont(p.instance, value)
}

func (p *TPageControl) ParentShowHint() bool {
    return PageControl_GetParentShowHint(p.instance)
}

func (p *TPageControl) SetParentShowHint(value bool) {
    PageControl_SetParentShowHint(p.instance, value)
}

func (p *TPageControl) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(PageControl_GetPopupMenu(p.instance))
}

func (p *TPageControl) SetPopupMenu(value IComponent) {
    PageControl_SetPopupMenu(p.instance, CheckPtr(value))
}

func (p *TPageControl) ShowHint() bool {
    return PageControl_GetShowHint(p.instance)
}

func (p *TPageControl) SetShowHint(value bool) {
    PageControl_SetShowHint(p.instance, value)
}

func (p *TPageControl) Style() TTabStyle {
    return PageControl_GetStyle(p.instance)
}

func (p *TPageControl) SetStyle(value TTabStyle) {
    PageControl_SetStyle(p.instance, value)
}

func (p *TPageControl) TabHeight() int16 {
    return PageControl_GetTabHeight(p.instance)
}

func (p *TPageControl) SetTabHeight(value int16) {
    PageControl_SetTabHeight(p.instance, value)
}

func (p *TPageControl) TabIndex() int32 {
    return PageControl_GetTabIndex(p.instance)
}

func (p *TPageControl) SetTabIndex(value int32) {
    PageControl_SetTabIndex(p.instance, value)
}

func (p *TPageControl) TabOrder() uint16 {
    return PageControl_GetTabOrder(p.instance)
}

func (p *TPageControl) SetTabOrder(value uint16) {
    PageControl_SetTabOrder(p.instance, value)
}

func (p *TPageControl) TabPosition() TTabPosition {
    return PageControl_GetTabPosition(p.instance)
}

func (p *TPageControl) SetTabPosition(value TTabPosition) {
    PageControl_SetTabPosition(p.instance, value)
}

func (p *TPageControl) TabStop() bool {
    return PageControl_GetTabStop(p.instance)
}

func (p *TPageControl) SetTabStop(value bool) {
    PageControl_SetTabStop(p.instance, value)
}

func (p *TPageControl) TabWidth() int16 {
    return PageControl_GetTabWidth(p.instance)
}

func (p *TPageControl) SetTabWidth(value int16) {
    PageControl_SetTabWidth(p.instance, value)
}

func (p *TPageControl) Visible() bool {
    return PageControl_GetVisible(p.instance)
}

func (p *TPageControl) SetVisible(value bool) {
    PageControl_SetVisible(p.instance, value)
}

func (p *TPageControl) StyleElements() TStyleElements {
    return PageControl_GetStyleElements(p.instance)
}

func (p *TPageControl) SetStyleElements(value TStyleElements) {
    PageControl_SetStyleElements(p.instance, value)
}

func (p *TPageControl) SetOnChange(fn TNotifyEvent) {
    PageControl_SetOnChange(p.instance, fn)
}

func (p *TPageControl) SetOnEnter(fn TNotifyEvent) {
    PageControl_SetOnEnter(p.instance, fn)
}

func (p *TPageControl) SetOnExit(fn TNotifyEvent) {
    PageControl_SetOnExit(p.instance, fn)
}

func (p *TPageControl) SetOnGetImageIndex(fn TTabGetImageEvent) {
    PageControl_SetOnGetImageIndex(p.instance, fn)
}

func (p *TPageControl) SetOnMouseDown(fn TMouseEvent) {
    PageControl_SetOnMouseDown(p.instance, fn)
}

func (p *TPageControl) SetOnMouseEnter(fn TNotifyEvent) {
    PageControl_SetOnMouseEnter(p.instance, fn)
}

func (p *TPageControl) SetOnMouseLeave(fn TNotifyEvent) {
    PageControl_SetOnMouseLeave(p.instance, fn)
}

func (p *TPageControl) SetOnMouseMove(fn TMouseMoveEvent) {
    PageControl_SetOnMouseMove(p.instance, fn)
}

func (p *TPageControl) SetOnMouseUp(fn TMouseEvent) {
    PageControl_SetOnMouseUp(p.instance, fn)
}

func (p *TPageControl) SetOnResize(fn TNotifyEvent) {
    PageControl_SetOnResize(p.instance, fn)
}

func (p *TPageControl) Canvas() *TCanvas {
    return CanvasFromInst(PageControl_GetCanvas(p.instance))
}

func (p *TPageControl) Brush() *TBrush {
    return BrushFromInst(PageControl_GetBrush(p.instance))
}

func (p *TPageControl) ControlCount() int32 {
    return PageControl_GetControlCount(p.instance)
}

func (p *TPageControl) Handle() HWND {
    return PageControl_GetHandle(p.instance)
}

func (p *TPageControl) ParentWindow() HWND {
    return PageControl_GetParentWindow(p.instance)
}

func (p *TPageControl) SetParentWindow(value HWND) {
    PageControl_SetParentWindow(p.instance, value)
}

func (p *TPageControl) Action() *TAction {
    return ActionFromInst(PageControl_GetAction(p.instance))
}

func (p *TPageControl) SetAction(value IComponent) {
    PageControl_SetAction(p.instance, CheckPtr(value))
}

func (p *TPageControl) BoundsRect() TRect {
    return PageControl_GetBoundsRect(p.instance)
}

func (p *TPageControl) SetBoundsRect(value TRect) {
    PageControl_SetBoundsRect(p.instance, value)
}

func (p *TPageControl) ClientHeight() int32 {
    return PageControl_GetClientHeight(p.instance)
}

func (p *TPageControl) SetClientHeight(value int32) {
    PageControl_SetClientHeight(p.instance, value)
}

func (p *TPageControl) ClientRect() TRect {
    return PageControl_GetClientRect(p.instance)
}

func (p *TPageControl) ClientWidth() int32 {
    return PageControl_GetClientWidth(p.instance)
}

func (p *TPageControl) SetClientWidth(value int32) {
    PageControl_SetClientWidth(p.instance, value)
}

func (p *TPageControl) ExplicitLeft() int32 {
    return PageControl_GetExplicitLeft(p.instance)
}

func (p *TPageControl) ExplicitTop() int32 {
    return PageControl_GetExplicitTop(p.instance)
}

func (p *TPageControl) ExplicitWidth() int32 {
    return PageControl_GetExplicitWidth(p.instance)
}

func (p *TPageControl) ExplicitHeight() int32 {
    return PageControl_GetExplicitHeight(p.instance)
}

func (p *TPageControl) Parent() *TControl {
    return ControlFromInst(PageControl_GetParent(p.instance))
}

func (p *TPageControl) SetParent(value IControl) {
    PageControl_SetParent(p.instance, CheckPtr(value))
}

func (p *TPageControl) AlignWithMargins() bool {
    return PageControl_GetAlignWithMargins(p.instance)
}

func (p *TPageControl) SetAlignWithMargins(value bool) {
    PageControl_SetAlignWithMargins(p.instance, value)
}

func (p *TPageControl) Left() int32 {
    return PageControl_GetLeft(p.instance)
}

func (p *TPageControl) SetLeft(value int32) {
    PageControl_SetLeft(p.instance, value)
}

func (p *TPageControl) Top() int32 {
    return PageControl_GetTop(p.instance)
}

func (p *TPageControl) SetTop(value int32) {
    PageControl_SetTop(p.instance, value)
}

func (p *TPageControl) Width() int32 {
    return PageControl_GetWidth(p.instance)
}

func (p *TPageControl) SetWidth(value int32) {
    PageControl_SetWidth(p.instance, value)
}

func (p *TPageControl) Height() int32 {
    return PageControl_GetHeight(p.instance)
}

func (p *TPageControl) SetHeight(value int32) {
    PageControl_SetHeight(p.instance, value)
}

func (p *TPageControl) Cursor() TCursor {
    return PageControl_GetCursor(p.instance)
}

func (p *TPageControl) SetCursor(value TCursor) {
    PageControl_SetCursor(p.instance, value)
}

func (p *TPageControl) Hint() string {
    return PageControl_GetHint(p.instance)
}

func (p *TPageControl) SetHint(value string) {
    PageControl_SetHint(p.instance, value)
}

func (p *TPageControl) Margins() *TMargins {
    return MarginsFromInst(PageControl_GetMargins(p.instance))
}

func (p *TPageControl) SetMargins(value *TMargins) {
    PageControl_SetMargins(p.instance, CheckPtr(value))
}

func (p *TPageControl) CustomHint() *TCustomHint {
    return CustomHintFromInst(PageControl_GetCustomHint(p.instance))
}

func (p *TPageControl) SetCustomHint(value IComponent) {
    PageControl_SetCustomHint(p.instance, CheckPtr(value))
}

func (p *TPageControl) ComponentCount() int32 {
    return PageControl_GetComponentCount(p.instance)
}

func (p *TPageControl) ComponentIndex() int32 {
    return PageControl_GetComponentIndex(p.instance)
}

func (p *TPageControl) SetComponentIndex(value int32) {
    PageControl_SetComponentIndex(p.instance, value)
}

func (p *TPageControl) Owner() *TComponent {
    return ComponentFromInst(PageControl_GetOwner(p.instance))
}

func (p *TPageControl) Name() string {
    return PageControl_GetName(p.instance)
}

func (p *TPageControl) SetName(value string) {
    PageControl_SetName(p.instance, value)
}

func (p *TPageControl) Tag() int {
    return PageControl_GetTag(p.instance)
}

func (p *TPageControl) SetTag(value int) {
    PageControl_SetTag(p.instance, value)
}

func (p *TPageControl) Pages(Index int32) *TTabSheet {
    return TabSheetFromInst(PageControl_GetPages(p.instance, Index))
}

func (p *TPageControl) Controls(Index int32) *TControl {
    return ControlFromInst(PageControl_GetControls(p.instance, Index))
}

func (p *TPageControl) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PageControl_GetComponents(p.instance, AIndex))
}

