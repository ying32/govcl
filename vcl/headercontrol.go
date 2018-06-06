
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

type THeaderControl struct {
    IWinControl
    instance uintptr
}

func NewHeaderControl(owner IComponent) *THeaderControl {
    h := new(THeaderControl)
    h.instance = HeaderControl_Create(CheckPtr(owner))
    return h
}

func HeaderControlFromInst(inst uintptr) *THeaderControl {
    h := new(THeaderControl)
    h.instance = inst
    return h
}

func HeaderControlFromObj(obj IObject) *THeaderControl {
    h := new(THeaderControl)
    h.instance = CheckPtr(obj)
    return h
}

func (h *THeaderControl) Free() {
    if h.instance != 0 {
        HeaderControl_Free(h.instance)
        h.instance = 0
    }
}

func (h *THeaderControl) Instance() uintptr {
    return h.instance
}

func (h *THeaderControl) IsValid() bool {
    return h.instance != 0
}

func THeaderControlClass() TClass {
    return HeaderControl_StaticClassType()
}

func (h *THeaderControl) FlipChildren(AllLevels bool) {
    HeaderControl_FlipChildren(h.instance, AllLevels)
}

func (h *THeaderControl) CanFocus() bool {
    return HeaderControl_CanFocus(h.instance)
}

func (h *THeaderControl) Focused() bool {
    return HeaderControl_Focused(h.instance)
}

func (h *THeaderControl) HandleAllocated() bool {
    return HeaderControl_HandleAllocated(h.instance)
}

func (h *THeaderControl) Invalidate() {
    HeaderControl_Invalidate(h.instance)
}

func (h *THeaderControl) Realign() {
    HeaderControl_Realign(h.instance)
}

func (h *THeaderControl) Repaint() {
    HeaderControl_Repaint(h.instance)
}

func (h *THeaderControl) ScaleBy(M int32, D int32) {
    HeaderControl_ScaleBy(h.instance, M , D)
}

func (h *THeaderControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    HeaderControl_SetBounds(h.instance, ALeft , ATop , AWidth , AHeight)
}

func (h *THeaderControl) SetFocus() {
    HeaderControl_SetFocus(h.instance)
}

func (h *THeaderControl) Update() {
    HeaderControl_Update(h.instance)
}

func (h *THeaderControl) BringToFront() {
    HeaderControl_BringToFront(h.instance)
}

func (h *THeaderControl) ClientToScreen(Point TPoint) TPoint {
    return HeaderControl_ClientToScreen(h.instance, Point)
}

func (h *THeaderControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return HeaderControl_ClientToParent(h.instance, Point , CheckPtr(AParent))
}

func (h *THeaderControl) Dragging() bool {
    return HeaderControl_Dragging(h.instance)
}

func (h *THeaderControl) HasParent() bool {
    return HeaderControl_HasParent(h.instance)
}

func (h *THeaderControl) Hide() {
    HeaderControl_Hide(h.instance)
}

func (h *THeaderControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return HeaderControl_Perform(h.instance, Msg , WParam , LParam)
}

func (h *THeaderControl) Refresh() {
    HeaderControl_Refresh(h.instance)
}

func (h *THeaderControl) ScreenToClient(Point TPoint) TPoint {
    return HeaderControl_ScreenToClient(h.instance, Point)
}

func (h *THeaderControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return HeaderControl_ParentToClient(h.instance, Point , CheckPtr(AParent))
}

func (h *THeaderControl) SendToBack() {
    HeaderControl_SendToBack(h.instance)
}

func (h *THeaderControl) Show() {
    HeaderControl_Show(h.instance)
}

func (h *THeaderControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    return HeaderControl_GetTextBuf(h.instance, Buffer , BufSize)
}

func (h *THeaderControl) GetTextLen() int32 {
    return HeaderControl_GetTextLen(h.instance)
}

func (h *THeaderControl) FindComponent(AName string) *TComponent {
    return ComponentFromInst(HeaderControl_FindComponent(h.instance, AName))
}

func (h *THeaderControl) GetNamePath() string {
    return HeaderControl_GetNamePath(h.instance)
}

func (h *THeaderControl) Assign(Source IObject) {
    HeaderControl_Assign(h.instance, CheckPtr(Source))
}

func (h *THeaderControl) DisposeOf() {
    HeaderControl_DisposeOf(h.instance)
}

func (h *THeaderControl) ClassType() TClass {
    return HeaderControl_ClassType(h.instance)
}

func (h *THeaderControl) ClassName() string {
    return HeaderControl_ClassName(h.instance)
}

func (h *THeaderControl) InstanceSize() int32 {
    return HeaderControl_InstanceSize(h.instance)
}

func (h *THeaderControl) InheritsFrom(AClass TClass) bool {
    return HeaderControl_InheritsFrom(h.instance, AClass)
}

func (h *THeaderControl) Equals(Obj IObject) bool {
    return HeaderControl_Equals(h.instance, CheckPtr(Obj))
}

func (h *THeaderControl) GetHashCode() int32 {
    return HeaderControl_GetHashCode(h.instance)
}

func (h *THeaderControl) ToString() string {
    return HeaderControl_ToString(h.instance)
}

func (h *THeaderControl) Align() TAlign {
    return HeaderControl_GetAlign(h.instance)
}

func (h *THeaderControl) SetAlign(value TAlign) {
    HeaderControl_SetAlign(h.instance, value)
}

func (h *THeaderControl) Anchors() TAnchors {
    return HeaderControl_GetAnchors(h.instance)
}

func (h *THeaderControl) SetAnchors(value TAnchors) {
    HeaderControl_SetAnchors(h.instance, value)
}

func (h *THeaderControl) BiDiMode() TBiDiMode {
    return HeaderControl_GetBiDiMode(h.instance)
}

func (h *THeaderControl) SetBiDiMode(value TBiDiMode) {
    HeaderControl_SetBiDiMode(h.instance, value)
}

func (h *THeaderControl) BorderWidth() int32 {
    return HeaderControl_GetBorderWidth(h.instance)
}

func (h *THeaderControl) SetBorderWidth(value int32) {
    HeaderControl_SetBorderWidth(h.instance, value)
}

func (h *THeaderControl) DoubleBuffered() bool {
    return HeaderControl_GetDoubleBuffered(h.instance)
}

func (h *THeaderControl) SetDoubleBuffered(value bool) {
    HeaderControl_SetDoubleBuffered(h.instance, value)
}

func (h *THeaderControl) DragCursor() TCursor {
    return HeaderControl_GetDragCursor(h.instance)
}

func (h *THeaderControl) SetDragCursor(value TCursor) {
    HeaderControl_SetDragCursor(h.instance, value)
}

func (h *THeaderControl) DragKind() TDragKind {
    return HeaderControl_GetDragKind(h.instance)
}

func (h *THeaderControl) SetDragKind(value TDragKind) {
    HeaderControl_SetDragKind(h.instance, value)
}

func (h *THeaderControl) DragMode() TDragMode {
    return HeaderControl_GetDragMode(h.instance)
}

func (h *THeaderControl) SetDragMode(value TDragMode) {
    HeaderControl_SetDragMode(h.instance, value)
}

func (h *THeaderControl) Enabled() bool {
    return HeaderControl_GetEnabled(h.instance)
}

func (h *THeaderControl) SetEnabled(value bool) {
    HeaderControl_SetEnabled(h.instance, value)
}

func (h *THeaderControl) Font() *TFont {
    return FontFromInst(HeaderControl_GetFont(h.instance))
}

func (h *THeaderControl) SetFont(value *TFont) {
    HeaderControl_SetFont(h.instance, CheckPtr(value))
}

func (h *THeaderControl) FullDrag() bool {
    return HeaderControl_GetFullDrag(h.instance)
}

func (h *THeaderControl) SetFullDrag(value bool) {
    HeaderControl_SetFullDrag(h.instance, value)
}

func (h *THeaderControl) HotTrack() bool {
    return HeaderControl_GetHotTrack(h.instance)
}

func (h *THeaderControl) SetHotTrack(value bool) {
    HeaderControl_SetHotTrack(h.instance, value)
}

func (h *THeaderControl) Images() *TImageList {
    return ImageListFromInst(HeaderControl_GetImages(h.instance))
}

func (h *THeaderControl) SetImages(value IComponent) {
    HeaderControl_SetImages(h.instance, CheckPtr(value))
}

func (h *THeaderControl) ShowHint() bool {
    return HeaderControl_GetShowHint(h.instance)
}

func (h *THeaderControl) SetShowHint(value bool) {
    HeaderControl_SetShowHint(h.instance, value)
}

func (h *THeaderControl) Style() THeaderStyle {
    return HeaderControl_GetStyle(h.instance)
}

func (h *THeaderControl) SetStyle(value THeaderStyle) {
    HeaderControl_SetStyle(h.instance, value)
}

func (h *THeaderControl) ParentDoubleBuffered() bool {
    return HeaderControl_GetParentDoubleBuffered(h.instance)
}

func (h *THeaderControl) SetParentDoubleBuffered(value bool) {
    HeaderControl_SetParentDoubleBuffered(h.instance, value)
}

func (h *THeaderControl) ParentFont() bool {
    return HeaderControl_GetParentFont(h.instance)
}

func (h *THeaderControl) SetParentFont(value bool) {
    HeaderControl_SetParentFont(h.instance, value)
}

func (h *THeaderControl) ParentShowHint() bool {
    return HeaderControl_GetParentShowHint(h.instance)
}

func (h *THeaderControl) SetParentShowHint(value bool) {
    HeaderControl_SetParentShowHint(h.instance, value)
}

func (h *THeaderControl) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(HeaderControl_GetPopupMenu(h.instance))
}

func (h *THeaderControl) SetPopupMenu(value IComponent) {
    HeaderControl_SetPopupMenu(h.instance, CheckPtr(value))
}

func (h *THeaderControl) Visible() bool {
    return HeaderControl_GetVisible(h.instance)
}

func (h *THeaderControl) SetVisible(value bool) {
    HeaderControl_SetVisible(h.instance, value)
}

func (h *THeaderControl) StyleElements() TStyleElements {
    return HeaderControl_GetStyleElements(h.instance)
}

func (h *THeaderControl) SetStyleElements(value TStyleElements) {
    HeaderControl_SetStyleElements(h.instance, value)
}

func (h *THeaderControl) SetOnContextPopup(fn TContextPopupEvent) {
    HeaderControl_SetOnContextPopup(h.instance, fn)
}

func (h *THeaderControl) SetOnDragDrop(fn TDragDropEvent) {
    HeaderControl_SetOnDragDrop(h.instance, fn)
}

func (h *THeaderControl) SetOnDragOver(fn TDragOverEvent) {
    HeaderControl_SetOnDragOver(h.instance, fn)
}

func (h *THeaderControl) SetOnEndDock(fn TEndDragEvent) {
    HeaderControl_SetOnEndDock(h.instance, fn)
}

func (h *THeaderControl) SetOnEndDrag(fn TEndDragEvent) {
    HeaderControl_SetOnEndDrag(h.instance, fn)
}

func (h *THeaderControl) SetOnMouseDown(fn TMouseEvent) {
    HeaderControl_SetOnMouseDown(h.instance, fn)
}

func (h *THeaderControl) SetOnMouseEnter(fn TNotifyEvent) {
    HeaderControl_SetOnMouseEnter(h.instance, fn)
}

func (h *THeaderControl) SetOnMouseLeave(fn TNotifyEvent) {
    HeaderControl_SetOnMouseLeave(h.instance, fn)
}

func (h *THeaderControl) SetOnMouseMove(fn TMouseMoveEvent) {
    HeaderControl_SetOnMouseMove(h.instance, fn)
}

func (h *THeaderControl) SetOnMouseUp(fn TMouseEvent) {
    HeaderControl_SetOnMouseUp(h.instance, fn)
}

func (h *THeaderControl) SetOnResize(fn TNotifyEvent) {
    HeaderControl_SetOnResize(h.instance, fn)
}

func (h *THeaderControl) SetOnDrawSection(fn TDrawSectionEvent) {
    HeaderControl_SetOnDrawSection(h.instance, fn)
}

func (h *THeaderControl) SetOnSectionClick(fn TSectionNotifyEvent) {
    HeaderControl_SetOnSectionClick(h.instance, fn)
}

func (h *THeaderControl) SetOnSectionResize(fn TSectionNotifyEvent) {
    HeaderControl_SetOnSectionResize(h.instance, fn)
}

func (h *THeaderControl) SetOnSectionTrack(fn TSectionTrackEvent) {
    HeaderControl_SetOnSectionTrack(h.instance, fn)
}

func (h *THeaderControl) SetOnSectionDrag(fn TSectionDragEvent) {
    HeaderControl_SetOnSectionDrag(h.instance, fn)
}

func (h *THeaderControl) SetOnSectionEndDrag(fn TNotifyEvent) {
    HeaderControl_SetOnSectionEndDrag(h.instance, fn)
}

func (h *THeaderControl) SetOnStartDock(fn TStartDockEvent) {
    HeaderControl_SetOnStartDock(h.instance, fn)
}

func (h *THeaderControl) Canvas() *TCanvas {
    return CanvasFromInst(HeaderControl_GetCanvas(h.instance))
}

func (h *THeaderControl) SetOnSectionCheck(fn TCustomSectionNotifyEvent) {
    HeaderControl_SetOnSectionCheck(h.instance, fn)
}

func (h *THeaderControl) DockSite() bool {
    return HeaderControl_GetDockSite(h.instance)
}

func (h *THeaderControl) SetDockSite(value bool) {
    HeaderControl_SetDockSite(h.instance, value)
}

func (h *THeaderControl) Brush() *TBrush {
    return BrushFromInst(HeaderControl_GetBrush(h.instance))
}

func (h *THeaderControl) ControlCount() int32 {
    return HeaderControl_GetControlCount(h.instance)
}

func (h *THeaderControl) Handle() HWND {
    return HeaderControl_GetHandle(h.instance)
}

func (h *THeaderControl) ParentWindow() HWND {
    return HeaderControl_GetParentWindow(h.instance)
}

func (h *THeaderControl) SetParentWindow(value HWND) {
    HeaderControl_SetParentWindow(h.instance, value)
}

func (h *THeaderControl) TabOrder() uint16 {
    return HeaderControl_GetTabOrder(h.instance)
}

func (h *THeaderControl) SetTabOrder(value uint16) {
    HeaderControl_SetTabOrder(h.instance, value)
}

func (h *THeaderControl) TabStop() bool {
    return HeaderControl_GetTabStop(h.instance)
}

func (h *THeaderControl) SetTabStop(value bool) {
    HeaderControl_SetTabStop(h.instance, value)
}

func (h *THeaderControl) UseDockManager() bool {
    return HeaderControl_GetUseDockManager(h.instance)
}

func (h *THeaderControl) SetUseDockManager(value bool) {
    HeaderControl_SetUseDockManager(h.instance, value)
}

func (h *THeaderControl) Action() *TAction {
    return ActionFromInst(HeaderControl_GetAction(h.instance))
}

func (h *THeaderControl) SetAction(value IComponent) {
    HeaderControl_SetAction(h.instance, CheckPtr(value))
}

func (h *THeaderControl) BoundsRect() TRect {
    return HeaderControl_GetBoundsRect(h.instance)
}

func (h *THeaderControl) SetBoundsRect(value TRect) {
    HeaderControl_SetBoundsRect(h.instance, value)
}

func (h *THeaderControl) ClientHeight() int32 {
    return HeaderControl_GetClientHeight(h.instance)
}

func (h *THeaderControl) SetClientHeight(value int32) {
    HeaderControl_SetClientHeight(h.instance, value)
}

func (h *THeaderControl) ClientRect() TRect {
    return HeaderControl_GetClientRect(h.instance)
}

func (h *THeaderControl) ClientWidth() int32 {
    return HeaderControl_GetClientWidth(h.instance)
}

func (h *THeaderControl) SetClientWidth(value int32) {
    HeaderControl_SetClientWidth(h.instance, value)
}

func (h *THeaderControl) ExplicitLeft() int32 {
    return HeaderControl_GetExplicitLeft(h.instance)
}

func (h *THeaderControl) ExplicitTop() int32 {
    return HeaderControl_GetExplicitTop(h.instance)
}

func (h *THeaderControl) ExplicitWidth() int32 {
    return HeaderControl_GetExplicitWidth(h.instance)
}

func (h *THeaderControl) ExplicitHeight() int32 {
    return HeaderControl_GetExplicitHeight(h.instance)
}

func (h *THeaderControl) Floating() bool {
    return HeaderControl_GetFloating(h.instance)
}

func (h *THeaderControl) Parent() *TWinControl {
    return WinControlFromInst(HeaderControl_GetParent(h.instance))
}

func (h *THeaderControl) SetParent(value IWinControl) {
    HeaderControl_SetParent(h.instance, CheckPtr(value))
}

func (h *THeaderControl) AlignWithMargins() bool {
    return HeaderControl_GetAlignWithMargins(h.instance)
}

func (h *THeaderControl) SetAlignWithMargins(value bool) {
    HeaderControl_SetAlignWithMargins(h.instance, value)
}

func (h *THeaderControl) Left() int32 {
    return HeaderControl_GetLeft(h.instance)
}

func (h *THeaderControl) SetLeft(value int32) {
    HeaderControl_SetLeft(h.instance, value)
}

func (h *THeaderControl) Top() int32 {
    return HeaderControl_GetTop(h.instance)
}

func (h *THeaderControl) SetTop(value int32) {
    HeaderControl_SetTop(h.instance, value)
}

func (h *THeaderControl) Width() int32 {
    return HeaderControl_GetWidth(h.instance)
}

func (h *THeaderControl) SetWidth(value int32) {
    HeaderControl_SetWidth(h.instance, value)
}

func (h *THeaderControl) Height() int32 {
    return HeaderControl_GetHeight(h.instance)
}

func (h *THeaderControl) SetHeight(value int32) {
    HeaderControl_SetHeight(h.instance, value)
}

func (h *THeaderControl) Cursor() TCursor {
    return HeaderControl_GetCursor(h.instance)
}

func (h *THeaderControl) SetCursor(value TCursor) {
    HeaderControl_SetCursor(h.instance, value)
}

func (h *THeaderControl) Hint() string {
    return HeaderControl_GetHint(h.instance)
}

func (h *THeaderControl) SetHint(value string) {
    HeaderControl_SetHint(h.instance, value)
}

func (h *THeaderControl) Margins() *TMargins {
    return MarginsFromInst(HeaderControl_GetMargins(h.instance))
}

func (h *THeaderControl) SetMargins(value *TMargins) {
    HeaderControl_SetMargins(h.instance, CheckPtr(value))
}

func (h *THeaderControl) CustomHint() *TCustomHint {
    return CustomHintFromInst(HeaderControl_GetCustomHint(h.instance))
}

func (h *THeaderControl) SetCustomHint(value IComponent) {
    HeaderControl_SetCustomHint(h.instance, CheckPtr(value))
}

func (h *THeaderControl) ComponentCount() int32 {
    return HeaderControl_GetComponentCount(h.instance)
}

func (h *THeaderControl) ComponentIndex() int32 {
    return HeaderControl_GetComponentIndex(h.instance)
}

func (h *THeaderControl) SetComponentIndex(value int32) {
    HeaderControl_SetComponentIndex(h.instance, value)
}

func (h *THeaderControl) Owner() *TComponent {
    return ComponentFromInst(HeaderControl_GetOwner(h.instance))
}

func (h *THeaderControl) Name() string {
    return HeaderControl_GetName(h.instance)
}

func (h *THeaderControl) SetName(value string) {
    HeaderControl_SetName(h.instance, value)
}

func (h *THeaderControl) Tag() int {
    return HeaderControl_GetTag(h.instance)
}

func (h *THeaderControl) SetTag(value int) {
    HeaderControl_SetTag(h.instance, value)
}

func (h *THeaderControl) Controls(Index int32) *TControl {
    return ControlFromInst(HeaderControl_GetControls(h.instance, Index))
}

func (h *THeaderControl) Components(AIndex int32) *TComponent {
    return ComponentFromInst(HeaderControl_GetComponents(h.instance, AIndex))
}

