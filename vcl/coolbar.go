
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

type TCoolBar struct {
    IWinControl
    instance uintptr
}

func NewCoolBar(owner IComponent) *TCoolBar {
    c := new(TCoolBar)
    c.instance = CoolBar_Create(CheckPtr(owner))
    return c
}

func CoolBarFromInst(inst uintptr) *TCoolBar {
    c := new(TCoolBar)
    c.instance = inst
    return c
}

func CoolBarFromObj(obj IObject) *TCoolBar {
    c := new(TCoolBar)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCoolBar) Free() {
    if c.instance != 0 {
        CoolBar_Free(c.instance)
        c.instance = 0
    }
}

func (c *TCoolBar) Instance() uintptr {
    return c.instance
}

func (c *TCoolBar) IsValid() bool {
    return c.instance != 0
}

func TCoolBarClass() TClass {
    return CoolBar_StaticClassType()
}

func (c *TCoolBar) FlipChildren(AllLevels bool) {
    CoolBar_FlipChildren(c.instance, AllLevels)
}

func (c *TCoolBar) CanFocus() bool {
    return CoolBar_CanFocus(c.instance)
}

func (c *TCoolBar) Focused() bool {
    return CoolBar_Focused(c.instance)
}

func (c *TCoolBar) HandleAllocated() bool {
    return CoolBar_HandleAllocated(c.instance)
}

func (c *TCoolBar) Invalidate() {
    CoolBar_Invalidate(c.instance)
}

func (c *TCoolBar) Realign() {
    CoolBar_Realign(c.instance)
}

func (c *TCoolBar) Repaint() {
    CoolBar_Repaint(c.instance)
}

func (c *TCoolBar) ScaleBy(M int32, D int32) {
    CoolBar_ScaleBy(c.instance, M , D)
}

func (c *TCoolBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CoolBar_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

func (c *TCoolBar) SetFocus() {
    CoolBar_SetFocus(c.instance)
}

func (c *TCoolBar) Update() {
    CoolBar_Update(c.instance)
}

func (c *TCoolBar) BringToFront() {
    CoolBar_BringToFront(c.instance)
}

func (c *TCoolBar) ClientToScreen(Point TPoint) TPoint {
    return CoolBar_ClientToScreen(c.instance, Point)
}

func (c *TCoolBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CoolBar_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

func (c *TCoolBar) Dragging() bool {
    return CoolBar_Dragging(c.instance)
}

func (c *TCoolBar) HasParent() bool {
    return CoolBar_HasParent(c.instance)
}

func (c *TCoolBar) Hide() {
    CoolBar_Hide(c.instance)
}

func (c *TCoolBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CoolBar_Perform(c.instance, Msg , WParam , LParam)
}

func (c *TCoolBar) Refresh() {
    CoolBar_Refresh(c.instance)
}

func (c *TCoolBar) ScreenToClient(Point TPoint) TPoint {
    return CoolBar_ScreenToClient(c.instance, Point)
}

func (c *TCoolBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CoolBar_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

func (c *TCoolBar) SendToBack() {
    CoolBar_SendToBack(c.instance)
}

func (c *TCoolBar) Show() {
    CoolBar_Show(c.instance)
}

func (c *TCoolBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CoolBar_GetTextBuf(c.instance, Buffer , BufSize)
}

func (c *TCoolBar) GetTextLen() int32 {
    return CoolBar_GetTextLen(c.instance)
}

func (c *TCoolBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CoolBar_FindComponent(c.instance, AName))
}

func (c *TCoolBar) GetNamePath() string {
    return CoolBar_GetNamePath(c.instance)
}

func (c *TCoolBar) Assign(Source IObject) {
    CoolBar_Assign(c.instance, CheckPtr(Source))
}

func (c *TCoolBar) DisposeOf() {
    CoolBar_DisposeOf(c.instance)
}

func (c *TCoolBar) ClassType() TClass {
    return CoolBar_ClassType(c.instance)
}

func (c *TCoolBar) ClassName() string {
    return CoolBar_ClassName(c.instance)
}

func (c *TCoolBar) InstanceSize() int32 {
    return CoolBar_InstanceSize(c.instance)
}

func (c *TCoolBar) InheritsFrom(AClass TClass) bool {
    return CoolBar_InheritsFrom(c.instance, AClass)
}

func (c *TCoolBar) Equals(Obj IObject) bool {
    return CoolBar_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCoolBar) GetHashCode() int32 {
    return CoolBar_GetHashCode(c.instance)
}

func (c *TCoolBar) ToString() string {
    return CoolBar_ToString(c.instance)
}

func (c *TCoolBar) Align() TAlign {
    return CoolBar_GetAlign(c.instance)
}

func (c *TCoolBar) SetAlign(value TAlign) {
    CoolBar_SetAlign(c.instance, value)
}

func (c *TCoolBar) Anchors() TAnchors {
    return CoolBar_GetAnchors(c.instance)
}

func (c *TCoolBar) SetAnchors(value TAnchors) {
    CoolBar_SetAnchors(c.instance, value)
}

func (c *TCoolBar) AutoSize() bool {
    return CoolBar_GetAutoSize(c.instance)
}

func (c *TCoolBar) SetAutoSize(value bool) {
    CoolBar_SetAutoSize(c.instance, value)
}

func (c *TCoolBar) BandBorderStyle() TBorderStyle {
    return CoolBar_GetBandBorderStyle(c.instance)
}

func (c *TCoolBar) SetBandBorderStyle(value TBorderStyle) {
    CoolBar_SetBandBorderStyle(c.instance, value)
}

func (c *TCoolBar) BandMaximize() TCoolBandMaximize {
    return CoolBar_GetBandMaximize(c.instance)
}

func (c *TCoolBar) SetBandMaximize(value TCoolBandMaximize) {
    CoolBar_SetBandMaximize(c.instance, value)
}

func (c *TCoolBar) Bands() *TCoolBands {
    return CoolBandsFromInst(CoolBar_GetBands(c.instance))
}

func (c *TCoolBar) SetBands(value *TCoolBands) {
    CoolBar_SetBands(c.instance, CheckPtr(value))
}

func (c *TCoolBar) BorderWidth() int32 {
    return CoolBar_GetBorderWidth(c.instance)
}

func (c *TCoolBar) SetBorderWidth(value int32) {
    CoolBar_SetBorderWidth(c.instance, value)
}

func (c *TCoolBar) Color() TColor {
    return CoolBar_GetColor(c.instance)
}

func (c *TCoolBar) SetColor(value TColor) {
    CoolBar_SetColor(c.instance, value)
}

func (c *TCoolBar) DockSite() bool {
    return CoolBar_GetDockSite(c.instance)
}

func (c *TCoolBar) SetDockSite(value bool) {
    CoolBar_SetDockSite(c.instance, value)
}

func (c *TCoolBar) DoubleBuffered() bool {
    return CoolBar_GetDoubleBuffered(c.instance)
}

func (c *TCoolBar) SetDoubleBuffered(value bool) {
    CoolBar_SetDoubleBuffered(c.instance, value)
}

func (c *TCoolBar) DragCursor() TCursor {
    return CoolBar_GetDragCursor(c.instance)
}

func (c *TCoolBar) SetDragCursor(value TCursor) {
    CoolBar_SetDragCursor(c.instance, value)
}

func (c *TCoolBar) DragKind() TDragKind {
    return CoolBar_GetDragKind(c.instance)
}

func (c *TCoolBar) SetDragKind(value TDragKind) {
    CoolBar_SetDragKind(c.instance, value)
}

func (c *TCoolBar) DragMode() TDragMode {
    return CoolBar_GetDragMode(c.instance)
}

func (c *TCoolBar) SetDragMode(value TDragMode) {
    CoolBar_SetDragMode(c.instance, value)
}

func (c *TCoolBar) EdgeBorders() TEdgeBorders {
    return CoolBar_GetEdgeBorders(c.instance)
}

func (c *TCoolBar) SetEdgeBorders(value TEdgeBorders) {
    CoolBar_SetEdgeBorders(c.instance, value)
}

func (c *TCoolBar) EdgeInner() TEdgeStyle {
    return CoolBar_GetEdgeInner(c.instance)
}

func (c *TCoolBar) SetEdgeInner(value TEdgeStyle) {
    CoolBar_SetEdgeInner(c.instance, value)
}

func (c *TCoolBar) EdgeOuter() TEdgeStyle {
    return CoolBar_GetEdgeOuter(c.instance)
}

func (c *TCoolBar) SetEdgeOuter(value TEdgeStyle) {
    CoolBar_SetEdgeOuter(c.instance, value)
}

func (c *TCoolBar) Enabled() bool {
    return CoolBar_GetEnabled(c.instance)
}

func (c *TCoolBar) SetEnabled(value bool) {
    CoolBar_SetEnabled(c.instance, value)
}

func (c *TCoolBar) FixedSize() bool {
    return CoolBar_GetFixedSize(c.instance)
}

func (c *TCoolBar) SetFixedSize(value bool) {
    CoolBar_SetFixedSize(c.instance, value)
}

func (c *TCoolBar) FixedOrder() bool {
    return CoolBar_GetFixedOrder(c.instance)
}

func (c *TCoolBar) SetFixedOrder(value bool) {
    CoolBar_SetFixedOrder(c.instance, value)
}

func (c *TCoolBar) Font() *TFont {
    return FontFromInst(CoolBar_GetFont(c.instance))
}

func (c *TCoolBar) SetFont(value *TFont) {
    CoolBar_SetFont(c.instance, CheckPtr(value))
}

func (c *TCoolBar) Images() *TImageList {
    return ImageListFromInst(CoolBar_GetImages(c.instance))
}

func (c *TCoolBar) SetImages(value IComponent) {
    CoolBar_SetImages(c.instance, CheckPtr(value))
}

func (c *TCoolBar) ParentColor() bool {
    return CoolBar_GetParentColor(c.instance)
}

func (c *TCoolBar) SetParentColor(value bool) {
    CoolBar_SetParentColor(c.instance, value)
}

func (c *TCoolBar) ParentDoubleBuffered() bool {
    return CoolBar_GetParentDoubleBuffered(c.instance)
}

func (c *TCoolBar) SetParentDoubleBuffered(value bool) {
    CoolBar_SetParentDoubleBuffered(c.instance, value)
}

func (c *TCoolBar) ParentFont() bool {
    return CoolBar_GetParentFont(c.instance)
}

func (c *TCoolBar) SetParentFont(value bool) {
    CoolBar_SetParentFont(c.instance, value)
}

func (c *TCoolBar) ParentShowHint() bool {
    return CoolBar_GetParentShowHint(c.instance)
}

func (c *TCoolBar) SetParentShowHint(value bool) {
    CoolBar_SetParentShowHint(c.instance, value)
}

func (c *TCoolBar) Bitmap() *TBitmap {
    return BitmapFromInst(CoolBar_GetBitmap(c.instance))
}

func (c *TCoolBar) SetBitmap(value *TBitmap) {
    CoolBar_SetBitmap(c.instance, CheckPtr(value))
}

func (c *TCoolBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CoolBar_GetPopupMenu(c.instance))
}

func (c *TCoolBar) SetPopupMenu(value IComponent) {
    CoolBar_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TCoolBar) ShowHint() bool {
    return CoolBar_GetShowHint(c.instance)
}

func (c *TCoolBar) SetShowHint(value bool) {
    CoolBar_SetShowHint(c.instance, value)
}

func (c *TCoolBar) ShowText() bool {
    return CoolBar_GetShowText(c.instance)
}

func (c *TCoolBar) SetShowText(value bool) {
    CoolBar_SetShowText(c.instance, value)
}

func (c *TCoolBar) Vertical() bool {
    return CoolBar_GetVertical(c.instance)
}

func (c *TCoolBar) SetVertical(value bool) {
    CoolBar_SetVertical(c.instance, value)
}

func (c *TCoolBar) Visible() bool {
    return CoolBar_GetVisible(c.instance)
}

func (c *TCoolBar) SetVisible(value bool) {
    CoolBar_SetVisible(c.instance, value)
}

func (c *TCoolBar) StyleElements() TStyleElements {
    return CoolBar_GetStyleElements(c.instance)
}

func (c *TCoolBar) SetStyleElements(value TStyleElements) {
    CoolBar_SetStyleElements(c.instance, value)
}

func (c *TCoolBar) SetOnChange(fn TNotifyEvent) {
    CoolBar_SetOnChange(c.instance, fn)
}

func (c *TCoolBar) SetOnClick(fn TNotifyEvent) {
    CoolBar_SetOnClick(c.instance, fn)
}

func (c *TCoolBar) SetOnContextPopup(fn TContextPopupEvent) {
    CoolBar_SetOnContextPopup(c.instance, fn)
}

func (c *TCoolBar) SetOnDblClick(fn TNotifyEvent) {
    CoolBar_SetOnDblClick(c.instance, fn)
}

func (c *TCoolBar) SetOnDockDrop(fn TDockDropEvent) {
    CoolBar_SetOnDockDrop(c.instance, fn)
}

func (c *TCoolBar) SetOnDragDrop(fn TDragDropEvent) {
    CoolBar_SetOnDragDrop(c.instance, fn)
}

func (c *TCoolBar) SetOnDragOver(fn TDragOverEvent) {
    CoolBar_SetOnDragOver(c.instance, fn)
}

func (c *TCoolBar) SetOnEndDock(fn TEndDragEvent) {
    CoolBar_SetOnEndDock(c.instance, fn)
}

func (c *TCoolBar) SetOnEndDrag(fn TEndDragEvent) {
    CoolBar_SetOnEndDrag(c.instance, fn)
}

func (c *TCoolBar) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    CoolBar_SetOnGetSiteInfo(c.instance, fn)
}

func (c *TCoolBar) SetOnMouseDown(fn TMouseEvent) {
    CoolBar_SetOnMouseDown(c.instance, fn)
}

func (c *TCoolBar) SetOnMouseEnter(fn TNotifyEvent) {
    CoolBar_SetOnMouseEnter(c.instance, fn)
}

func (c *TCoolBar) SetOnMouseLeave(fn TNotifyEvent) {
    CoolBar_SetOnMouseLeave(c.instance, fn)
}

func (c *TCoolBar) SetOnMouseMove(fn TMouseMoveEvent) {
    CoolBar_SetOnMouseMove(c.instance, fn)
}

func (c *TCoolBar) SetOnMouseUp(fn TMouseEvent) {
    CoolBar_SetOnMouseUp(c.instance, fn)
}

func (c *TCoolBar) SetOnResize(fn TNotifyEvent) {
    CoolBar_SetOnResize(c.instance, fn)
}

func (c *TCoolBar) SetOnStartDock(fn TStartDockEvent) {
    CoolBar_SetOnStartDock(c.instance, fn)
}

func (c *TCoolBar) SetOnUnDock(fn TUnDockEvent) {
    CoolBar_SetOnUnDock(c.instance, fn)
}

func (c *TCoolBar) Brush() *TBrush {
    return BrushFromInst(CoolBar_GetBrush(c.instance))
}

func (c *TCoolBar) ControlCount() int32 {
    return CoolBar_GetControlCount(c.instance)
}

func (c *TCoolBar) Handle() HWND {
    return CoolBar_GetHandle(c.instance)
}

func (c *TCoolBar) ParentWindow() HWND {
    return CoolBar_GetParentWindow(c.instance)
}

func (c *TCoolBar) SetParentWindow(value HWND) {
    CoolBar_SetParentWindow(c.instance, value)
}

func (c *TCoolBar) TabOrder() uint16 {
    return CoolBar_GetTabOrder(c.instance)
}

func (c *TCoolBar) SetTabOrder(value uint16) {
    CoolBar_SetTabOrder(c.instance, value)
}

func (c *TCoolBar) TabStop() bool {
    return CoolBar_GetTabStop(c.instance)
}

func (c *TCoolBar) SetTabStop(value bool) {
    CoolBar_SetTabStop(c.instance, value)
}

func (c *TCoolBar) UseDockManager() bool {
    return CoolBar_GetUseDockManager(c.instance)
}

func (c *TCoolBar) SetUseDockManager(value bool) {
    CoolBar_SetUseDockManager(c.instance, value)
}

func (c *TCoolBar) Action() *TAction {
    return ActionFromInst(CoolBar_GetAction(c.instance))
}

func (c *TCoolBar) SetAction(value IComponent) {
    CoolBar_SetAction(c.instance, CheckPtr(value))
}

func (c *TCoolBar) BiDiMode() TBiDiMode {
    return CoolBar_GetBiDiMode(c.instance)
}

func (c *TCoolBar) SetBiDiMode(value TBiDiMode) {
    CoolBar_SetBiDiMode(c.instance, value)
}

func (c *TCoolBar) BoundsRect() TRect {
    return CoolBar_GetBoundsRect(c.instance)
}

func (c *TCoolBar) SetBoundsRect(value TRect) {
    CoolBar_SetBoundsRect(c.instance, value)
}

func (c *TCoolBar) ClientHeight() int32 {
    return CoolBar_GetClientHeight(c.instance)
}

func (c *TCoolBar) SetClientHeight(value int32) {
    CoolBar_SetClientHeight(c.instance, value)
}

func (c *TCoolBar) ClientRect() TRect {
    return CoolBar_GetClientRect(c.instance)
}

func (c *TCoolBar) ClientWidth() int32 {
    return CoolBar_GetClientWidth(c.instance)
}

func (c *TCoolBar) SetClientWidth(value int32) {
    CoolBar_SetClientWidth(c.instance, value)
}

func (c *TCoolBar) ExplicitLeft() int32 {
    return CoolBar_GetExplicitLeft(c.instance)
}

func (c *TCoolBar) ExplicitTop() int32 {
    return CoolBar_GetExplicitTop(c.instance)
}

func (c *TCoolBar) ExplicitWidth() int32 {
    return CoolBar_GetExplicitWidth(c.instance)
}

func (c *TCoolBar) ExplicitHeight() int32 {
    return CoolBar_GetExplicitHeight(c.instance)
}

func (c *TCoolBar) Floating() bool {
    return CoolBar_GetFloating(c.instance)
}

func (c *TCoolBar) Parent() *TWinControl {
    return WinControlFromInst(CoolBar_GetParent(c.instance))
}

func (c *TCoolBar) SetParent(value IWinControl) {
    CoolBar_SetParent(c.instance, CheckPtr(value))
}

func (c *TCoolBar) AlignWithMargins() bool {
    return CoolBar_GetAlignWithMargins(c.instance)
}

func (c *TCoolBar) SetAlignWithMargins(value bool) {
    CoolBar_SetAlignWithMargins(c.instance, value)
}

func (c *TCoolBar) Left() int32 {
    return CoolBar_GetLeft(c.instance)
}

func (c *TCoolBar) SetLeft(value int32) {
    CoolBar_SetLeft(c.instance, value)
}

func (c *TCoolBar) Top() int32 {
    return CoolBar_GetTop(c.instance)
}

func (c *TCoolBar) SetTop(value int32) {
    CoolBar_SetTop(c.instance, value)
}

func (c *TCoolBar) Width() int32 {
    return CoolBar_GetWidth(c.instance)
}

func (c *TCoolBar) SetWidth(value int32) {
    CoolBar_SetWidth(c.instance, value)
}

func (c *TCoolBar) Height() int32 {
    return CoolBar_GetHeight(c.instance)
}

func (c *TCoolBar) SetHeight(value int32) {
    CoolBar_SetHeight(c.instance, value)
}

func (c *TCoolBar) Cursor() TCursor {
    return CoolBar_GetCursor(c.instance)
}

func (c *TCoolBar) SetCursor(value TCursor) {
    CoolBar_SetCursor(c.instance, value)
}

func (c *TCoolBar) Hint() string {
    return CoolBar_GetHint(c.instance)
}

func (c *TCoolBar) SetHint(value string) {
    CoolBar_SetHint(c.instance, value)
}

func (c *TCoolBar) Margins() *TMargins {
    return MarginsFromInst(CoolBar_GetMargins(c.instance))
}

func (c *TCoolBar) SetMargins(value *TMargins) {
    CoolBar_SetMargins(c.instance, CheckPtr(value))
}

func (c *TCoolBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(CoolBar_GetCustomHint(c.instance))
}

func (c *TCoolBar) SetCustomHint(value IComponent) {
    CoolBar_SetCustomHint(c.instance, CheckPtr(value))
}

func (c *TCoolBar) ComponentCount() int32 {
    return CoolBar_GetComponentCount(c.instance)
}

func (c *TCoolBar) ComponentIndex() int32 {
    return CoolBar_GetComponentIndex(c.instance)
}

func (c *TCoolBar) SetComponentIndex(value int32) {
    CoolBar_SetComponentIndex(c.instance, value)
}

func (c *TCoolBar) Owner() *TComponent {
    return ComponentFromInst(CoolBar_GetOwner(c.instance))
}

func (c *TCoolBar) Name() string {
    return CoolBar_GetName(c.instance)
}

func (c *TCoolBar) SetName(value string) {
    CoolBar_SetName(c.instance, value)
}

func (c *TCoolBar) Tag() int {
    return CoolBar_GetTag(c.instance)
}

func (c *TCoolBar) SetTag(value int) {
    CoolBar_SetTag(c.instance, value)
}

func (c *TCoolBar) Controls(Index int32) *TControl {
    return ControlFromInst(CoolBar_GetControls(c.instance, Index))
}

func (c *TCoolBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CoolBar_GetComponents(c.instance, AIndex))
}

