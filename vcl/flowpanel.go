
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

type TFlowPanel struct {
    IWinControl
    instance uintptr
}

func NewFlowPanel(owner IComponent) *TFlowPanel {
    f := new(TFlowPanel)
    f.instance = FlowPanel_Create(CheckPtr(owner))
    return f
}

func FlowPanelFromInst(inst uintptr) *TFlowPanel {
    f := new(TFlowPanel)
    f.instance = inst
    return f
}

func FlowPanelFromObj(obj IObject) *TFlowPanel {
    f := new(TFlowPanel)
    f.instance = CheckPtr(obj)
    return f
}

func (f *TFlowPanel) Free() {
    if f.instance != 0 {
        FlowPanel_Free(f.instance)
        f.instance = 0
    }
}

func (f *TFlowPanel) Instance() uintptr {
    return f.instance
}

func (f *TFlowPanel) IsValid() bool {
    return f.instance != 0
}

func TFlowPanelClass() TClass {
    return FlowPanel_StaticClassType()
}

func (f *TFlowPanel) GetControlIndex(AControl IControl) int32 {
    return FlowPanel_GetControlIndex(f.instance, CheckPtr(AControl))
}

func (f *TFlowPanel) SetControlIndex(AControl IControl, Index int32) {
    FlowPanel_SetControlIndex(f.instance, CheckPtr(AControl), Index)
}

func (f *TFlowPanel) CanFocus() bool {
    return FlowPanel_CanFocus(f.instance)
}

func (f *TFlowPanel) FlipChildren(AllLevels bool) {
    FlowPanel_FlipChildren(f.instance, AllLevels)
}

func (f *TFlowPanel) Focused() bool {
    return FlowPanel_Focused(f.instance)
}

func (f *TFlowPanel) HandleAllocated() bool {
    return FlowPanel_HandleAllocated(f.instance)
}

func (f *TFlowPanel) Invalidate() {
    FlowPanel_Invalidate(f.instance)
}

func (f *TFlowPanel) Realign() {
    FlowPanel_Realign(f.instance)
}

func (f *TFlowPanel) Repaint() {
    FlowPanel_Repaint(f.instance)
}

func (f *TFlowPanel) ScaleBy(M int32, D int32) {
    FlowPanel_ScaleBy(f.instance, M , D)
}

func (f *TFlowPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    FlowPanel_SetBounds(f.instance, ALeft , ATop , AWidth , AHeight)
}

func (f *TFlowPanel) SetFocus() {
    FlowPanel_SetFocus(f.instance)
}

func (f *TFlowPanel) Update() {
    FlowPanel_Update(f.instance)
}

func (f *TFlowPanel) BringToFront() {
    FlowPanel_BringToFront(f.instance)
}

func (f *TFlowPanel) ClientToScreen(Point TPoint) TPoint {
    return FlowPanel_ClientToScreen(f.instance, Point)
}

func (f *TFlowPanel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return FlowPanel_ClientToParent(f.instance, Point , CheckPtr(AParent))
}

func (f *TFlowPanel) Dragging() bool {
    return FlowPanel_Dragging(f.instance)
}

func (f *TFlowPanel) HasParent() bool {
    return FlowPanel_HasParent(f.instance)
}

func (f *TFlowPanel) Hide() {
    FlowPanel_Hide(f.instance)
}

func (f *TFlowPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return FlowPanel_Perform(f.instance, Msg , WParam , LParam)
}

func (f *TFlowPanel) Refresh() {
    FlowPanel_Refresh(f.instance)
}

func (f *TFlowPanel) ScreenToClient(Point TPoint) TPoint {
    return FlowPanel_ScreenToClient(f.instance, Point)
}

func (f *TFlowPanel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return FlowPanel_ParentToClient(f.instance, Point , CheckPtr(AParent))
}

func (f *TFlowPanel) SendToBack() {
    FlowPanel_SendToBack(f.instance)
}

func (f *TFlowPanel) Show() {
    FlowPanel_Show(f.instance)
}

func (f *TFlowPanel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return FlowPanel_GetTextBuf(f.instance, Buffer , BufSize)
}

func (f *TFlowPanel) GetTextLen() int32 {
    return FlowPanel_GetTextLen(f.instance)
}

func (f *TFlowPanel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(FlowPanel_FindComponent(f.instance, AName))
}

func (f *TFlowPanel) GetNamePath() string {
    return FlowPanel_GetNamePath(f.instance)
}

func (f *TFlowPanel) Assign(Source IObject) {
    FlowPanel_Assign(f.instance, CheckPtr(Source))
}

func (f *TFlowPanel) DisposeOf() {
    FlowPanel_DisposeOf(f.instance)
}

func (f *TFlowPanel) ClassType() TClass {
    return FlowPanel_ClassType(f.instance)
}

func (f *TFlowPanel) ClassName() string {
    return FlowPanel_ClassName(f.instance)
}

func (f *TFlowPanel) InstanceSize() int32 {
    return FlowPanel_InstanceSize(f.instance)
}

func (f *TFlowPanel) InheritsFrom(AClass TClass) bool {
    return FlowPanel_InheritsFrom(f.instance, AClass)
}

func (f *TFlowPanel) Equals(Obj IObject) bool {
    return FlowPanel_Equals(f.instance, CheckPtr(Obj))
}

func (f *TFlowPanel) GetHashCode() int32 {
    return FlowPanel_GetHashCode(f.instance)
}

func (f *TFlowPanel) ToString() string {
    return FlowPanel_ToString(f.instance)
}

func (f *TFlowPanel) Align() TAlign {
    return FlowPanel_GetAlign(f.instance)
}

func (f *TFlowPanel) SetAlign(value TAlign) {
    FlowPanel_SetAlign(f.instance, value)
}

func (f *TFlowPanel) Alignment() TAlignment {
    return FlowPanel_GetAlignment(f.instance)
}

func (f *TFlowPanel) SetAlignment(value TAlignment) {
    FlowPanel_SetAlignment(f.instance, value)
}

func (f *TFlowPanel) Anchors() TAnchors {
    return FlowPanel_GetAnchors(f.instance)
}

func (f *TFlowPanel) SetAnchors(value TAnchors) {
    FlowPanel_SetAnchors(f.instance, value)
}

func (f *TFlowPanel) AutoSize() bool {
    return FlowPanel_GetAutoSize(f.instance)
}

func (f *TFlowPanel) SetAutoSize(value bool) {
    FlowPanel_SetAutoSize(f.instance, value)
}

func (f *TFlowPanel) AutoWrap() bool {
    return FlowPanel_GetAutoWrap(f.instance)
}

func (f *TFlowPanel) SetAutoWrap(value bool) {
    FlowPanel_SetAutoWrap(f.instance, value)
}

func (f *TFlowPanel) BevelEdges() TBevelEdges {
    return FlowPanel_GetBevelEdges(f.instance)
}

func (f *TFlowPanel) SetBevelEdges(value TBevelEdges) {
    FlowPanel_SetBevelEdges(f.instance, value)
}

func (f *TFlowPanel) BevelInner() TBevelCut {
    return FlowPanel_GetBevelInner(f.instance)
}

func (f *TFlowPanel) SetBevelInner(value TBevelCut) {
    FlowPanel_SetBevelInner(f.instance, value)
}

func (f *TFlowPanel) BevelKind() TBevelKind {
    return FlowPanel_GetBevelKind(f.instance)
}

func (f *TFlowPanel) SetBevelKind(value TBevelKind) {
    FlowPanel_SetBevelKind(f.instance, value)
}

func (f *TFlowPanel) BevelOuter() TBevelCut {
    return FlowPanel_GetBevelOuter(f.instance)
}

func (f *TFlowPanel) SetBevelOuter(value TBevelCut) {
    FlowPanel_SetBevelOuter(f.instance, value)
}

func (f *TFlowPanel) BiDiMode() TBiDiMode {
    return FlowPanel_GetBiDiMode(f.instance)
}

func (f *TFlowPanel) SetBiDiMode(value TBiDiMode) {
    FlowPanel_SetBiDiMode(f.instance, value)
}

func (f *TFlowPanel) BorderWidth() int32 {
    return FlowPanel_GetBorderWidth(f.instance)
}

func (f *TFlowPanel) SetBorderWidth(value int32) {
    FlowPanel_SetBorderWidth(f.instance, value)
}

func (f *TFlowPanel) BorderStyle() TBorderStyle {
    return FlowPanel_GetBorderStyle(f.instance)
}

func (f *TFlowPanel) SetBorderStyle(value TBorderStyle) {
    FlowPanel_SetBorderStyle(f.instance, value)
}

func (f *TFlowPanel) Caption() string {
    return FlowPanel_GetCaption(f.instance)
}

func (f *TFlowPanel) SetCaption(value string) {
    FlowPanel_SetCaption(f.instance, value)
}

func (f *TFlowPanel) Color() TColor {
    return FlowPanel_GetColor(f.instance)
}

func (f *TFlowPanel) SetColor(value TColor) {
    FlowPanel_SetColor(f.instance, value)
}

func (f *TFlowPanel) UseDockManager() bool {
    return FlowPanel_GetUseDockManager(f.instance)
}

func (f *TFlowPanel) SetUseDockManager(value bool) {
    FlowPanel_SetUseDockManager(f.instance, value)
}

func (f *TFlowPanel) DockSite() bool {
    return FlowPanel_GetDockSite(f.instance)
}

func (f *TFlowPanel) SetDockSite(value bool) {
    FlowPanel_SetDockSite(f.instance, value)
}

func (f *TFlowPanel) DoubleBuffered() bool {
    return FlowPanel_GetDoubleBuffered(f.instance)
}

func (f *TFlowPanel) SetDoubleBuffered(value bool) {
    FlowPanel_SetDoubleBuffered(f.instance, value)
}

func (f *TFlowPanel) DragCursor() TCursor {
    return FlowPanel_GetDragCursor(f.instance)
}

func (f *TFlowPanel) SetDragCursor(value TCursor) {
    FlowPanel_SetDragCursor(f.instance, value)
}

func (f *TFlowPanel) DragKind() TDragKind {
    return FlowPanel_GetDragKind(f.instance)
}

func (f *TFlowPanel) SetDragKind(value TDragKind) {
    FlowPanel_SetDragKind(f.instance, value)
}

func (f *TFlowPanel) DragMode() TDragMode {
    return FlowPanel_GetDragMode(f.instance)
}

func (f *TFlowPanel) SetDragMode(value TDragMode) {
    FlowPanel_SetDragMode(f.instance, value)
}

func (f *TFlowPanel) Enabled() bool {
    return FlowPanel_GetEnabled(f.instance)
}

func (f *TFlowPanel) SetEnabled(value bool) {
    FlowPanel_SetEnabled(f.instance, value)
}

func (f *TFlowPanel) FlowStyle() TFlowStyle {
    return FlowPanel_GetFlowStyle(f.instance)
}

func (f *TFlowPanel) SetFlowStyle(value TFlowStyle) {
    FlowPanel_SetFlowStyle(f.instance, value)
}

func (f *TFlowPanel) FullRepaint() bool {
    return FlowPanel_GetFullRepaint(f.instance)
}

func (f *TFlowPanel) SetFullRepaint(value bool) {
    FlowPanel_SetFullRepaint(f.instance, value)
}

func (f *TFlowPanel) Font() *TFont {
    return FontFromInst(FlowPanel_GetFont(f.instance))
}

func (f *TFlowPanel) SetFont(value *TFont) {
    FlowPanel_SetFont(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) Locked() bool {
    return FlowPanel_GetLocked(f.instance)
}

func (f *TFlowPanel) SetLocked(value bool) {
    FlowPanel_SetLocked(f.instance, value)
}

func (f *TFlowPanel) ParentBackground() bool {
    return FlowPanel_GetParentBackground(f.instance)
}

func (f *TFlowPanel) SetParentBackground(value bool) {
    FlowPanel_SetParentBackground(f.instance, value)
}

func (f *TFlowPanel) ParentColor() bool {
    return FlowPanel_GetParentColor(f.instance)
}

func (f *TFlowPanel) SetParentColor(value bool) {
    FlowPanel_SetParentColor(f.instance, value)
}

func (f *TFlowPanel) ParentCtl3D() bool {
    return FlowPanel_GetParentCtl3D(f.instance)
}

func (f *TFlowPanel) SetParentCtl3D(value bool) {
    FlowPanel_SetParentCtl3D(f.instance, value)
}

func (f *TFlowPanel) ParentDoubleBuffered() bool {
    return FlowPanel_GetParentDoubleBuffered(f.instance)
}

func (f *TFlowPanel) SetParentDoubleBuffered(value bool) {
    FlowPanel_SetParentDoubleBuffered(f.instance, value)
}

func (f *TFlowPanel) ParentFont() bool {
    return FlowPanel_GetParentFont(f.instance)
}

func (f *TFlowPanel) SetParentFont(value bool) {
    FlowPanel_SetParentFont(f.instance, value)
}

func (f *TFlowPanel) ParentShowHint() bool {
    return FlowPanel_GetParentShowHint(f.instance)
}

func (f *TFlowPanel) SetParentShowHint(value bool) {
    FlowPanel_SetParentShowHint(f.instance, value)
}

func (f *TFlowPanel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(FlowPanel_GetPopupMenu(f.instance))
}

func (f *TFlowPanel) SetPopupMenu(value IComponent) {
    FlowPanel_SetPopupMenu(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) ShowCaption() bool {
    return FlowPanel_GetShowCaption(f.instance)
}

func (f *TFlowPanel) SetShowCaption(value bool) {
    FlowPanel_SetShowCaption(f.instance, value)
}

func (f *TFlowPanel) ShowHint() bool {
    return FlowPanel_GetShowHint(f.instance)
}

func (f *TFlowPanel) SetShowHint(value bool) {
    FlowPanel_SetShowHint(f.instance, value)
}

func (f *TFlowPanel) TabOrder() uint16 {
    return FlowPanel_GetTabOrder(f.instance)
}

func (f *TFlowPanel) SetTabOrder(value uint16) {
    FlowPanel_SetTabOrder(f.instance, value)
}

func (f *TFlowPanel) TabStop() bool {
    return FlowPanel_GetTabStop(f.instance)
}

func (f *TFlowPanel) SetTabStop(value bool) {
    FlowPanel_SetTabStop(f.instance, value)
}

func (f *TFlowPanel) Visible() bool {
    return FlowPanel_GetVisible(f.instance)
}

func (f *TFlowPanel) SetVisible(value bool) {
    FlowPanel_SetVisible(f.instance, value)
}

func (f *TFlowPanel) StyleElements() TStyleElements {
    return FlowPanel_GetStyleElements(f.instance)
}

func (f *TFlowPanel) SetStyleElements(value TStyleElements) {
    FlowPanel_SetStyleElements(f.instance, value)
}

func (f *TFlowPanel) SetOnClick(fn TNotifyEvent) {
    FlowPanel_SetOnClick(f.instance, fn)
}

func (f *TFlowPanel) SetOnContextPopup(fn TContextPopupEvent) {
    FlowPanel_SetOnContextPopup(f.instance, fn)
}

func (f *TFlowPanel) SetOnDockDrop(fn TDockDropEvent) {
    FlowPanel_SetOnDockDrop(f.instance, fn)
}

func (f *TFlowPanel) SetOnDblClick(fn TNotifyEvent) {
    FlowPanel_SetOnDblClick(f.instance, fn)
}

func (f *TFlowPanel) SetOnDragDrop(fn TDragDropEvent) {
    FlowPanel_SetOnDragDrop(f.instance, fn)
}

func (f *TFlowPanel) SetOnDragOver(fn TDragOverEvent) {
    FlowPanel_SetOnDragOver(f.instance, fn)
}

func (f *TFlowPanel) SetOnEndDock(fn TEndDragEvent) {
    FlowPanel_SetOnEndDock(f.instance, fn)
}

func (f *TFlowPanel) SetOnEndDrag(fn TEndDragEvent) {
    FlowPanel_SetOnEndDrag(f.instance, fn)
}

func (f *TFlowPanel) SetOnEnter(fn TNotifyEvent) {
    FlowPanel_SetOnEnter(f.instance, fn)
}

func (f *TFlowPanel) SetOnExit(fn TNotifyEvent) {
    FlowPanel_SetOnExit(f.instance, fn)
}

func (f *TFlowPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    FlowPanel_SetOnGetSiteInfo(f.instance, fn)
}

func (f *TFlowPanel) SetOnMouseDown(fn TMouseEvent) {
    FlowPanel_SetOnMouseDown(f.instance, fn)
}

func (f *TFlowPanel) SetOnMouseEnter(fn TNotifyEvent) {
    FlowPanel_SetOnMouseEnter(f.instance, fn)
}

func (f *TFlowPanel) SetOnMouseLeave(fn TNotifyEvent) {
    FlowPanel_SetOnMouseLeave(f.instance, fn)
}

func (f *TFlowPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    FlowPanel_SetOnMouseMove(f.instance, fn)
}

func (f *TFlowPanel) SetOnMouseUp(fn TMouseEvent) {
    FlowPanel_SetOnMouseUp(f.instance, fn)
}

func (f *TFlowPanel) SetOnResize(fn TNotifyEvent) {
    FlowPanel_SetOnResize(f.instance, fn)
}

func (f *TFlowPanel) SetOnStartDock(fn TStartDockEvent) {
    FlowPanel_SetOnStartDock(f.instance, fn)
}

func (f *TFlowPanel) SetOnUnDock(fn TUnDockEvent) {
    FlowPanel_SetOnUnDock(f.instance, fn)
}

func (f *TFlowPanel) Brush() *TBrush {
    return BrushFromInst(FlowPanel_GetBrush(f.instance))
}

func (f *TFlowPanel) ControlCount() int32 {
    return FlowPanel_GetControlCount(f.instance)
}

func (f *TFlowPanel) Handle() HWND {
    return FlowPanel_GetHandle(f.instance)
}

func (f *TFlowPanel) ParentWindow() HWND {
    return FlowPanel_GetParentWindow(f.instance)
}

func (f *TFlowPanel) SetParentWindow(value HWND) {
    FlowPanel_SetParentWindow(f.instance, value)
}

func (f *TFlowPanel) Action() *TAction {
    return ActionFromInst(FlowPanel_GetAction(f.instance))
}

func (f *TFlowPanel) SetAction(value IComponent) {
    FlowPanel_SetAction(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) BoundsRect() TRect {
    return FlowPanel_GetBoundsRect(f.instance)
}

func (f *TFlowPanel) SetBoundsRect(value TRect) {
    FlowPanel_SetBoundsRect(f.instance, value)
}

func (f *TFlowPanel) ClientHeight() int32 {
    return FlowPanel_GetClientHeight(f.instance)
}

func (f *TFlowPanel) SetClientHeight(value int32) {
    FlowPanel_SetClientHeight(f.instance, value)
}

func (f *TFlowPanel) ClientRect() TRect {
    return FlowPanel_GetClientRect(f.instance)
}

func (f *TFlowPanel) ClientWidth() int32 {
    return FlowPanel_GetClientWidth(f.instance)
}

func (f *TFlowPanel) SetClientWidth(value int32) {
    FlowPanel_SetClientWidth(f.instance, value)
}

func (f *TFlowPanel) ExplicitLeft() int32 {
    return FlowPanel_GetExplicitLeft(f.instance)
}

func (f *TFlowPanel) ExplicitTop() int32 {
    return FlowPanel_GetExplicitTop(f.instance)
}

func (f *TFlowPanel) ExplicitWidth() int32 {
    return FlowPanel_GetExplicitWidth(f.instance)
}

func (f *TFlowPanel) ExplicitHeight() int32 {
    return FlowPanel_GetExplicitHeight(f.instance)
}

func (f *TFlowPanel) Floating() bool {
    return FlowPanel_GetFloating(f.instance)
}

func (f *TFlowPanel) Parent() *TWinControl {
    return WinControlFromInst(FlowPanel_GetParent(f.instance))
}

func (f *TFlowPanel) SetParent(value IWinControl) {
    FlowPanel_SetParent(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) AlignWithMargins() bool {
    return FlowPanel_GetAlignWithMargins(f.instance)
}

func (f *TFlowPanel) SetAlignWithMargins(value bool) {
    FlowPanel_SetAlignWithMargins(f.instance, value)
}

func (f *TFlowPanel) Left() int32 {
    return FlowPanel_GetLeft(f.instance)
}

func (f *TFlowPanel) SetLeft(value int32) {
    FlowPanel_SetLeft(f.instance, value)
}

func (f *TFlowPanel) Top() int32 {
    return FlowPanel_GetTop(f.instance)
}

func (f *TFlowPanel) SetTop(value int32) {
    FlowPanel_SetTop(f.instance, value)
}

func (f *TFlowPanel) Width() int32 {
    return FlowPanel_GetWidth(f.instance)
}

func (f *TFlowPanel) SetWidth(value int32) {
    FlowPanel_SetWidth(f.instance, value)
}

func (f *TFlowPanel) Height() int32 {
    return FlowPanel_GetHeight(f.instance)
}

func (f *TFlowPanel) SetHeight(value int32) {
    FlowPanel_SetHeight(f.instance, value)
}

func (f *TFlowPanel) Cursor() TCursor {
    return FlowPanel_GetCursor(f.instance)
}

func (f *TFlowPanel) SetCursor(value TCursor) {
    FlowPanel_SetCursor(f.instance, value)
}

func (f *TFlowPanel) Hint() string {
    return FlowPanel_GetHint(f.instance)
}

func (f *TFlowPanel) SetHint(value string) {
    FlowPanel_SetHint(f.instance, value)
}

func (f *TFlowPanel) Margins() *TMargins {
    return MarginsFromInst(FlowPanel_GetMargins(f.instance))
}

func (f *TFlowPanel) SetMargins(value *TMargins) {
    FlowPanel_SetMargins(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) CustomHint() *TCustomHint {
    return CustomHintFromInst(FlowPanel_GetCustomHint(f.instance))
}

func (f *TFlowPanel) SetCustomHint(value IComponent) {
    FlowPanel_SetCustomHint(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) ComponentCount() int32 {
    return FlowPanel_GetComponentCount(f.instance)
}

func (f *TFlowPanel) ComponentIndex() int32 {
    return FlowPanel_GetComponentIndex(f.instance)
}

func (f *TFlowPanel) SetComponentIndex(value int32) {
    FlowPanel_SetComponentIndex(f.instance, value)
}

func (f *TFlowPanel) Owner() *TComponent {
    return ComponentFromInst(FlowPanel_GetOwner(f.instance))
}

func (f *TFlowPanel) Name() string {
    return FlowPanel_GetName(f.instance)
}

func (f *TFlowPanel) SetName(value string) {
    FlowPanel_SetName(f.instance, value)
}

func (f *TFlowPanel) Tag() int {
    return FlowPanel_GetTag(f.instance)
}

func (f *TFlowPanel) SetTag(value int) {
    FlowPanel_SetTag(f.instance, value)
}

func (f *TFlowPanel) Controls(Index int32) *TControl {
    return ControlFromInst(FlowPanel_GetControls(f.instance, Index))
}

func (f *TFlowPanel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(FlowPanel_GetComponents(f.instance, AIndex))
}

