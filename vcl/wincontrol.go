
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

type TWinControl struct {
    IWinControl
    instance uintptr
}

func NewWinControl(owner IComponent) *TWinControl {
    w := new(TWinControl)
    w.instance = WinControl_Create(CheckPtr(owner))
    return w
}

func WinControlFromInst(inst uintptr) *TWinControl {
    w := new(TWinControl)
    w.instance = inst
    return w
}

func WinControlFromObj(obj IObject) *TWinControl {
    w := new(TWinControl)
    w.instance = CheckPtr(obj)
    return w
}

func (w *TWinControl) Free() {
    if w.instance != 0 {
        WinControl_Free(w.instance)
        w.instance = 0
    }
}

func (w *TWinControl) Instance() uintptr {
    return w.instance
}

func (w *TWinControl) IsValid() bool {
    return w.instance != 0
}

func TWinControlClass() TClass {
    return WinControl_StaticClassType()
}

func (w *TWinControl) CanFocus() bool {
    return WinControl_CanFocus(w.instance)
}

func (w *TWinControl) FlipChildren(AllLevels bool) {
    WinControl_FlipChildren(w.instance, AllLevels)
}

func (w *TWinControl) Focused() bool {
    return WinControl_Focused(w.instance)
}

func (w *TWinControl) HandleAllocated() bool {
    return WinControl_HandleAllocated(w.instance)
}

func (w *TWinControl) Invalidate() {
    WinControl_Invalidate(w.instance)
}

func (w *TWinControl) Realign() {
    WinControl_Realign(w.instance)
}

func (w *TWinControl) Repaint() {
    WinControl_Repaint(w.instance)
}

func (w *TWinControl) ScaleBy(M int32, D int32) {
    WinControl_ScaleBy(w.instance, M , D)
}

func (w *TWinControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    WinControl_SetBounds(w.instance, ALeft , ATop , AWidth , AHeight)
}

func (w *TWinControl) SetFocus() {
    WinControl_SetFocus(w.instance)
}

func (w *TWinControl) Update() {
    WinControl_Update(w.instance)
}

func (w *TWinControl) BringToFront() {
    WinControl_BringToFront(w.instance)
}

func (w *TWinControl) ClientToScreen(Point TPoint) TPoint {
    return WinControl_ClientToScreen(w.instance, Point)
}

func (w *TWinControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return WinControl_ClientToParent(w.instance, Point , CheckPtr(AParent))
}

func (w *TWinControl) Dragging() bool {
    return WinControl_Dragging(w.instance)
}

func (w *TWinControl) HasParent() bool {
    return WinControl_HasParent(w.instance)
}

func (w *TWinControl) Hide() {
    WinControl_Hide(w.instance)
}

func (w *TWinControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return WinControl_Perform(w.instance, Msg , WParam , LParam)
}

func (w *TWinControl) Refresh() {
    WinControl_Refresh(w.instance)
}

func (w *TWinControl) ScreenToClient(Point TPoint) TPoint {
    return WinControl_ScreenToClient(w.instance, Point)
}

func (w *TWinControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return WinControl_ParentToClient(w.instance, Point , CheckPtr(AParent))
}

func (w *TWinControl) SendToBack() {
    WinControl_SendToBack(w.instance)
}

func (w *TWinControl) Show() {
    WinControl_Show(w.instance)
}

func (w *TWinControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    return WinControl_GetTextBuf(w.instance, Buffer , BufSize)
}

func (w *TWinControl) GetTextLen() int32 {
    return WinControl_GetTextLen(w.instance)
}

func (w *TWinControl) FindComponent(AName string) *TComponent {
    return ComponentFromInst(WinControl_FindComponent(w.instance, AName))
}

func (w *TWinControl) GetNamePath() string {
    return WinControl_GetNamePath(w.instance)
}

func (w *TWinControl) Assign(Source IObject) {
    WinControl_Assign(w.instance, CheckPtr(Source))
}

func (w *TWinControl) DisposeOf() {
    WinControl_DisposeOf(w.instance)
}

func (w *TWinControl) ClassType() TClass {
    return WinControl_ClassType(w.instance)
}

func (w *TWinControl) ClassName() string {
    return WinControl_ClassName(w.instance)
}

func (w *TWinControl) InstanceSize() int32 {
    return WinControl_InstanceSize(w.instance)
}

func (w *TWinControl) InheritsFrom(AClass TClass) bool {
    return WinControl_InheritsFrom(w.instance, AClass)
}

func (w *TWinControl) Equals(Obj IObject) bool {
    return WinControl_Equals(w.instance, CheckPtr(Obj))
}

func (w *TWinControl) GetHashCode() int32 {
    return WinControl_GetHashCode(w.instance)
}

func (w *TWinControl) ToString() string {
    return WinControl_ToString(w.instance)
}

func (w *TWinControl) DockSite() bool {
    return WinControl_GetDockSite(w.instance)
}

func (w *TWinControl) SetDockSite(value bool) {
    WinControl_SetDockSite(w.instance, value)
}

func (w *TWinControl) DoubleBuffered() bool {
    return WinControl_GetDoubleBuffered(w.instance)
}

func (w *TWinControl) SetDoubleBuffered(value bool) {
    WinControl_SetDoubleBuffered(w.instance, value)
}

func (w *TWinControl) Brush() *TBrush {
    return BrushFromInst(WinControl_GetBrush(w.instance))
}

func (w *TWinControl) ControlCount() int32 {
    return WinControl_GetControlCount(w.instance)
}

func (w *TWinControl) Handle() HWND {
    return WinControl_GetHandle(w.instance)
}

func (w *TWinControl) ParentDoubleBuffered() bool {
    return WinControl_GetParentDoubleBuffered(w.instance)
}

func (w *TWinControl) SetParentDoubleBuffered(value bool) {
    WinControl_SetParentDoubleBuffered(w.instance, value)
}

func (w *TWinControl) ParentWindow() HWND {
    return WinControl_GetParentWindow(w.instance)
}

func (w *TWinControl) SetParentWindow(value HWND) {
    WinControl_SetParentWindow(w.instance, value)
}

func (w *TWinControl) TabOrder() uint16 {
    return WinControl_GetTabOrder(w.instance)
}

func (w *TWinControl) SetTabOrder(value uint16) {
    WinControl_SetTabOrder(w.instance, value)
}

func (w *TWinControl) TabStop() bool {
    return WinControl_GetTabStop(w.instance)
}

func (w *TWinControl) SetTabStop(value bool) {
    WinControl_SetTabStop(w.instance, value)
}

func (w *TWinControl) UseDockManager() bool {
    return WinControl_GetUseDockManager(w.instance)
}

func (w *TWinControl) SetUseDockManager(value bool) {
    WinControl_SetUseDockManager(w.instance, value)
}

func (w *TWinControl) Enabled() bool {
    return WinControl_GetEnabled(w.instance)
}

func (w *TWinControl) SetEnabled(value bool) {
    WinControl_SetEnabled(w.instance, value)
}

func (w *TWinControl) Action() *TAction {
    return ActionFromInst(WinControl_GetAction(w.instance))
}

func (w *TWinControl) SetAction(value IComponent) {
    WinControl_SetAction(w.instance, CheckPtr(value))
}

func (w *TWinControl) Align() TAlign {
    return WinControl_GetAlign(w.instance)
}

func (w *TWinControl) SetAlign(value TAlign) {
    WinControl_SetAlign(w.instance, value)
}

func (w *TWinControl) Anchors() TAnchors {
    return WinControl_GetAnchors(w.instance)
}

func (w *TWinControl) SetAnchors(value TAnchors) {
    WinControl_SetAnchors(w.instance, value)
}

func (w *TWinControl) BiDiMode() TBiDiMode {
    return WinControl_GetBiDiMode(w.instance)
}

func (w *TWinControl) SetBiDiMode(value TBiDiMode) {
    WinControl_SetBiDiMode(w.instance, value)
}

func (w *TWinControl) BoundsRect() TRect {
    return WinControl_GetBoundsRect(w.instance)
}

func (w *TWinControl) SetBoundsRect(value TRect) {
    WinControl_SetBoundsRect(w.instance, value)
}

func (w *TWinControl) ClientHeight() int32 {
    return WinControl_GetClientHeight(w.instance)
}

func (w *TWinControl) SetClientHeight(value int32) {
    WinControl_SetClientHeight(w.instance, value)
}

func (w *TWinControl) ClientRect() TRect {
    return WinControl_GetClientRect(w.instance)
}

func (w *TWinControl) ClientWidth() int32 {
    return WinControl_GetClientWidth(w.instance)
}

func (w *TWinControl) SetClientWidth(value int32) {
    WinControl_SetClientWidth(w.instance, value)
}

func (w *TWinControl) ExplicitLeft() int32 {
    return WinControl_GetExplicitLeft(w.instance)
}

func (w *TWinControl) ExplicitTop() int32 {
    return WinControl_GetExplicitTop(w.instance)
}

func (w *TWinControl) ExplicitWidth() int32 {
    return WinControl_GetExplicitWidth(w.instance)
}

func (w *TWinControl) ExplicitHeight() int32 {
    return WinControl_GetExplicitHeight(w.instance)
}

func (w *TWinControl) Floating() bool {
    return WinControl_GetFloating(w.instance)
}

func (w *TWinControl) ShowHint() bool {
    return WinControl_GetShowHint(w.instance)
}

func (w *TWinControl) SetShowHint(value bool) {
    WinControl_SetShowHint(w.instance, value)
}

func (w *TWinControl) Visible() bool {
    return WinControl_GetVisible(w.instance)
}

func (w *TWinControl) SetVisible(value bool) {
    WinControl_SetVisible(w.instance, value)
}

func (w *TWinControl) Parent() *TWinControl {
    return WinControlFromInst(WinControl_GetParent(w.instance))
}

func (w *TWinControl) SetParent(value IWinControl) {
    WinControl_SetParent(w.instance, CheckPtr(value))
}

func (w *TWinControl) StyleElements() TStyleElements {
    return WinControl_GetStyleElements(w.instance)
}

func (w *TWinControl) SetStyleElements(value TStyleElements) {
    WinControl_SetStyleElements(w.instance, value)
}

func (w *TWinControl) AlignWithMargins() bool {
    return WinControl_GetAlignWithMargins(w.instance)
}

func (w *TWinControl) SetAlignWithMargins(value bool) {
    WinControl_SetAlignWithMargins(w.instance, value)
}

func (w *TWinControl) Left() int32 {
    return WinControl_GetLeft(w.instance)
}

func (w *TWinControl) SetLeft(value int32) {
    WinControl_SetLeft(w.instance, value)
}

func (w *TWinControl) Top() int32 {
    return WinControl_GetTop(w.instance)
}

func (w *TWinControl) SetTop(value int32) {
    WinControl_SetTop(w.instance, value)
}

func (w *TWinControl) Width() int32 {
    return WinControl_GetWidth(w.instance)
}

func (w *TWinControl) SetWidth(value int32) {
    WinControl_SetWidth(w.instance, value)
}

func (w *TWinControl) Height() int32 {
    return WinControl_GetHeight(w.instance)
}

func (w *TWinControl) SetHeight(value int32) {
    WinControl_SetHeight(w.instance, value)
}

func (w *TWinControl) Cursor() TCursor {
    return WinControl_GetCursor(w.instance)
}

func (w *TWinControl) SetCursor(value TCursor) {
    WinControl_SetCursor(w.instance, value)
}

func (w *TWinControl) Hint() string {
    return WinControl_GetHint(w.instance)
}

func (w *TWinControl) SetHint(value string) {
    WinControl_SetHint(w.instance, value)
}

func (w *TWinControl) Margins() *TMargins {
    return MarginsFromInst(WinControl_GetMargins(w.instance))
}

func (w *TWinControl) SetMargins(value *TMargins) {
    WinControl_SetMargins(w.instance, CheckPtr(value))
}

func (w *TWinControl) CustomHint() *TCustomHint {
    return CustomHintFromInst(WinControl_GetCustomHint(w.instance))
}

func (w *TWinControl) SetCustomHint(value IComponent) {
    WinControl_SetCustomHint(w.instance, CheckPtr(value))
}

func (w *TWinControl) ComponentCount() int32 {
    return WinControl_GetComponentCount(w.instance)
}

func (w *TWinControl) ComponentIndex() int32 {
    return WinControl_GetComponentIndex(w.instance)
}

func (w *TWinControl) SetComponentIndex(value int32) {
    WinControl_SetComponentIndex(w.instance, value)
}

func (w *TWinControl) Owner() *TComponent {
    return ComponentFromInst(WinControl_GetOwner(w.instance))
}

func (w *TWinControl) Name() string {
    return WinControl_GetName(w.instance)
}

func (w *TWinControl) SetName(value string) {
    WinControl_SetName(w.instance, value)
}

func (w *TWinControl) Tag() int {
    return WinControl_GetTag(w.instance)
}

func (w *TWinControl) SetTag(value int) {
    WinControl_SetTag(w.instance, value)
}

func (w *TWinControl) Controls(Index int32) *TControl {
    return ControlFromInst(WinControl_GetControls(w.instance, Index))
}

func (w *TWinControl) Components(AIndex int32) *TComponent {
    return ComponentFromInst(WinControl_GetComponents(w.instance, AIndex))
}

