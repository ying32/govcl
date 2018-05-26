
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

type TCategoryPanel struct {
    IControl
    instance uintptr
}

func NewCategoryPanel(owner IComponent) *TCategoryPanel {
    c := new(TCategoryPanel)
    c.instance = CategoryPanel_Create(CheckPtr(owner))
    return c
}

func CategoryPanelFromInst(inst uintptr) *TCategoryPanel {
    c := new(TCategoryPanel)
    c.instance = inst
    return c
}

func CategoryPanelFromObj(obj IObject) *TCategoryPanel {
    c := new(TCategoryPanel)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCategoryPanel) Free() {
    if c.instance != 0 {
        CategoryPanel_Free(c.instance)
        c.instance = 0
    }
}

func (c *TCategoryPanel) Instance() uintptr {
    return c.instance
}

func (c *TCategoryPanel) IsValid() bool {
    return c.instance != 0
}

func (c *TCategoryPanel) Collapse() {
    CategoryPanel_Collapse(c.instance)
}

func (c *TCategoryPanel) Expand() {
    CategoryPanel_Expand(c.instance)
}

func (c *TCategoryPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CategoryPanel_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

func (c *TCategoryPanel) CanFocus() bool {
    return CategoryPanel_CanFocus(c.instance)
}

func (c *TCategoryPanel) FlipChildren(AllLevels bool) {
    CategoryPanel_FlipChildren(c.instance, AllLevels)
}

func (c *TCategoryPanel) Focused() bool {
    return CategoryPanel_Focused(c.instance)
}

func (c *TCategoryPanel) HandleAllocated() bool {
    return CategoryPanel_HandleAllocated(c.instance)
}

func (c *TCategoryPanel) Invalidate() {
    CategoryPanel_Invalidate(c.instance)
}

func (c *TCategoryPanel) Realign() {
    CategoryPanel_Realign(c.instance)
}

func (c *TCategoryPanel) Repaint() {
    CategoryPanel_Repaint(c.instance)
}

func (c *TCategoryPanel) ScaleBy(M int32, D int32) {
    CategoryPanel_ScaleBy(c.instance, M , D)
}

func (c *TCategoryPanel) SetFocus() {
    CategoryPanel_SetFocus(c.instance)
}

func (c *TCategoryPanel) Update() {
    CategoryPanel_Update(c.instance)
}

func (c *TCategoryPanel) BringToFront() {
    CategoryPanel_BringToFront(c.instance)
}

func (c *TCategoryPanel) HasParent() bool {
    return CategoryPanel_HasParent(c.instance)
}

func (c *TCategoryPanel) Hide() {
    CategoryPanel_Hide(c.instance)
}

func (c *TCategoryPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CategoryPanel_Perform(c.instance, Msg , WParam , LParam)
}

func (c *TCategoryPanel) Refresh() {
    CategoryPanel_Refresh(c.instance)
}

func (c *TCategoryPanel) SendToBack() {
    CategoryPanel_SendToBack(c.instance)
}

func (c *TCategoryPanel) Show() {
    CategoryPanel_Show(c.instance)
}

func (c *TCategoryPanel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CategoryPanel_GetTextBuf(c.instance, Buffer , BufSize)
}

func (c *TCategoryPanel) GetTextLen() int32 {
    return CategoryPanel_GetTextLen(c.instance)
}

func (c *TCategoryPanel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CategoryPanel_FindComponent(c.instance, AName))
}

func (c *TCategoryPanel) GetNamePath() string {
    return CategoryPanel_GetNamePath(c.instance)
}

func (c *TCategoryPanel) Assign(Source IObject) {
    CategoryPanel_Assign(c.instance, CheckPtr(Source))
}

func (c *TCategoryPanel) ClassName() string {
    return CategoryPanel_ClassName(c.instance)
}

func (c *TCategoryPanel) Equals(Obj IObject) bool {
    return CategoryPanel_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCategoryPanel) GetHashCode() int32 {
    return CategoryPanel_GetHashCode(c.instance)
}

func (c *TCategoryPanel) ToString() string {
    return CategoryPanel_ToString(c.instance)
}

func (c *TCategoryPanel) BiDiMode() TBiDiMode {
    return CategoryPanel_GetBiDiMode(c.instance)
}

func (c *TCategoryPanel) SetBiDiMode(value TBiDiMode) {
    CategoryPanel_SetBiDiMode(c.instance, value)
}

func (c *TCategoryPanel) Caption() string {
    return CategoryPanel_GetCaption(c.instance)
}

func (c *TCategoryPanel) SetCaption(value string) {
    CategoryPanel_SetCaption(c.instance, value)
}

func (c *TCategoryPanel) Color() TColor {
    return CategoryPanel_GetColor(c.instance)
}

func (c *TCategoryPanel) SetColor(value TColor) {
    CategoryPanel_SetColor(c.instance, value)
}

func (c *TCategoryPanel) Collapsed() bool {
    return CategoryPanel_GetCollapsed(c.instance)
}

func (c *TCategoryPanel) SetCollapsed(value bool) {
    CategoryPanel_SetCollapsed(c.instance, value)
}

func (c *TCategoryPanel) CollapsedHotImageIndex() int32 {
    return CategoryPanel_GetCollapsedHotImageIndex(c.instance)
}

func (c *TCategoryPanel) SetCollapsedHotImageIndex(value int32) {
    CategoryPanel_SetCollapsedHotImageIndex(c.instance, value)
}

func (c *TCategoryPanel) CollapsedImageIndex() int32 {
    return CategoryPanel_GetCollapsedImageIndex(c.instance)
}

func (c *TCategoryPanel) SetCollapsedImageIndex(value int32) {
    CategoryPanel_SetCollapsedImageIndex(c.instance, value)
}

func (c *TCategoryPanel) CollapsedPressedImageIndex() int32 {
    return CategoryPanel_GetCollapsedPressedImageIndex(c.instance)
}

func (c *TCategoryPanel) SetCollapsedPressedImageIndex(value int32) {
    CategoryPanel_SetCollapsedPressedImageIndex(c.instance, value)
}

func (c *TCategoryPanel) DoubleBuffered() bool {
    return CategoryPanel_GetDoubleBuffered(c.instance)
}

func (c *TCategoryPanel) SetDoubleBuffered(value bool) {
    CategoryPanel_SetDoubleBuffered(c.instance, value)
}

func (c *TCategoryPanel) Enabled() bool {
    return CategoryPanel_GetEnabled(c.instance)
}

func (c *TCategoryPanel) SetEnabled(value bool) {
    CategoryPanel_SetEnabled(c.instance, value)
}

func (c *TCategoryPanel) ExpandedHotImageIndex() int32 {
    return CategoryPanel_GetExpandedHotImageIndex(c.instance)
}

func (c *TCategoryPanel) SetExpandedHotImageIndex(value int32) {
    CategoryPanel_SetExpandedHotImageIndex(c.instance, value)
}

func (c *TCategoryPanel) ExpandedImageIndex() int32 {
    return CategoryPanel_GetExpandedImageIndex(c.instance)
}

func (c *TCategoryPanel) SetExpandedImageIndex(value int32) {
    CategoryPanel_SetExpandedImageIndex(c.instance, value)
}

func (c *TCategoryPanel) ExpandedPressedImageIndex() int32 {
    return CategoryPanel_GetExpandedPressedImageIndex(c.instance)
}

func (c *TCategoryPanel) SetExpandedPressedImageIndex(value int32) {
    CategoryPanel_SetExpandedPressedImageIndex(c.instance, value)
}

func (c *TCategoryPanel) FullRepaint() bool {
    return CategoryPanel_GetFullRepaint(c.instance)
}

func (c *TCategoryPanel) SetFullRepaint(value bool) {
    CategoryPanel_SetFullRepaint(c.instance, value)
}

func (c *TCategoryPanel) Font() *TFont {
    return FontFromInst(CategoryPanel_GetFont(c.instance))
}

func (c *TCategoryPanel) SetFont(value *TFont) {
    CategoryPanel_SetFont(c.instance, CheckPtr(value))
}

func (c *TCategoryPanel) Height() int32 {
    return CategoryPanel_GetHeight(c.instance)
}

func (c *TCategoryPanel) SetHeight(value int32) {
    CategoryPanel_SetHeight(c.instance, value)
}

func (c *TCategoryPanel) Left() int32 {
    return CategoryPanel_GetLeft(c.instance)
}

func (c *TCategoryPanel) SetLeft(value int32) {
    CategoryPanel_SetLeft(c.instance, value)
}

func (c *TCategoryPanel) Locked() bool {
    return CategoryPanel_GetLocked(c.instance)
}

func (c *TCategoryPanel) SetLocked(value bool) {
    CategoryPanel_SetLocked(c.instance, value)
}

func (c *TCategoryPanel) ParentBackground() bool {
    return CategoryPanel_GetParentBackground(c.instance)
}

func (c *TCategoryPanel) SetParentBackground(value bool) {
    CategoryPanel_SetParentBackground(c.instance, value)
}

func (c *TCategoryPanel) ParentColor() bool {
    return CategoryPanel_GetParentColor(c.instance)
}

func (c *TCategoryPanel) SetParentColor(value bool) {
    CategoryPanel_SetParentColor(c.instance, value)
}

func (c *TCategoryPanel) ParentCtl3D() bool {
    return CategoryPanel_GetParentCtl3D(c.instance)
}

func (c *TCategoryPanel) SetParentCtl3D(value bool) {
    CategoryPanel_SetParentCtl3D(c.instance, value)
}

func (c *TCategoryPanel) ParentDoubleBuffered() bool {
    return CategoryPanel_GetParentDoubleBuffered(c.instance)
}

func (c *TCategoryPanel) SetParentDoubleBuffered(value bool) {
    CategoryPanel_SetParentDoubleBuffered(c.instance, value)
}

func (c *TCategoryPanel) ParentFont() bool {
    return CategoryPanel_GetParentFont(c.instance)
}

func (c *TCategoryPanel) SetParentFont(value bool) {
    CategoryPanel_SetParentFont(c.instance, value)
}

func (c *TCategoryPanel) ParentShowHint() bool {
    return CategoryPanel_GetParentShowHint(c.instance)
}

func (c *TCategoryPanel) SetParentShowHint(value bool) {
    CategoryPanel_SetParentShowHint(c.instance, value)
}

func (c *TCategoryPanel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CategoryPanel_GetPopupMenu(c.instance))
}

func (c *TCategoryPanel) SetPopupMenu(value IComponent) {
    CategoryPanel_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TCategoryPanel) ShowHint() bool {
    return CategoryPanel_GetShowHint(c.instance)
}

func (c *TCategoryPanel) SetShowHint(value bool) {
    CategoryPanel_SetShowHint(c.instance, value)
}

func (c *TCategoryPanel) TabOrder() uint16 {
    return CategoryPanel_GetTabOrder(c.instance)
}

func (c *TCategoryPanel) SetTabOrder(value uint16) {
    CategoryPanel_SetTabOrder(c.instance, value)
}

func (c *TCategoryPanel) TabStop() bool {
    return CategoryPanel_GetTabStop(c.instance)
}

func (c *TCategoryPanel) SetTabStop(value bool) {
    CategoryPanel_SetTabStop(c.instance, value)
}

func (c *TCategoryPanel) Visible() bool {
    return CategoryPanel_GetVisible(c.instance)
}

func (c *TCategoryPanel) SetVisible(value bool) {
    CategoryPanel_SetVisible(c.instance, value)
}

func (c *TCategoryPanel) StyleElements() TStyleElements {
    return CategoryPanel_GetStyleElements(c.instance)
}

func (c *TCategoryPanel) SetStyleElements(value TStyleElements) {
    CategoryPanel_SetStyleElements(c.instance, value)
}

func (c *TCategoryPanel) Width() int32 {
    return CategoryPanel_GetWidth(c.instance)
}

func (c *TCategoryPanel) SetWidth(value int32) {
    CategoryPanel_SetWidth(c.instance, value)
}

func (c *TCategoryPanel) SetOnClick(fn TNotifyEvent) {
    CategoryPanel_SetOnClick(c.instance, fn)
}

func (c *TCategoryPanel) SetOnDblClick(fn TNotifyEvent) {
    CategoryPanel_SetOnDblClick(c.instance, fn)
}

func (c *TCategoryPanel) SetOnEnter(fn TNotifyEvent) {
    CategoryPanel_SetOnEnter(c.instance, fn)
}

func (c *TCategoryPanel) SetOnExit(fn TNotifyEvent) {
    CategoryPanel_SetOnExit(c.instance, fn)
}

func (c *TCategoryPanel) SetOnMouseDown(fn TMouseEvent) {
    CategoryPanel_SetOnMouseDown(c.instance, fn)
}

func (c *TCategoryPanel) SetOnMouseEnter(fn TNotifyEvent) {
    CategoryPanel_SetOnMouseEnter(c.instance, fn)
}

func (c *TCategoryPanel) SetOnMouseLeave(fn TNotifyEvent) {
    CategoryPanel_SetOnMouseLeave(c.instance, fn)
}

func (c *TCategoryPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    CategoryPanel_SetOnMouseMove(c.instance, fn)
}

func (c *TCategoryPanel) SetOnMouseUp(fn TMouseEvent) {
    CategoryPanel_SetOnMouseUp(c.instance, fn)
}

func (c *TCategoryPanel) PanelGroup() *TCategoryPanelGroup {
    return CategoryPanelGroupFromInst(CategoryPanel_GetPanelGroup(c.instance))
}

func (c *TCategoryPanel) SetPanelGroup(value IControl) {
    CategoryPanel_SetPanelGroup(c.instance, CheckPtr(value))
}

func (c *TCategoryPanel) Brush() *TBrush {
    return BrushFromInst(CategoryPanel_GetBrush(c.instance))
}

func (c *TCategoryPanel) ControlCount() int32 {
    return CategoryPanel_GetControlCount(c.instance)
}

func (c *TCategoryPanel) Handle() HWND {
    return CategoryPanel_GetHandle(c.instance)
}

func (c *TCategoryPanel) ParentWindow() HWND {
    return CategoryPanel_GetParentWindow(c.instance)
}

func (c *TCategoryPanel) SetParentWindow(value HWND) {
    CategoryPanel_SetParentWindow(c.instance, value)
}

func (c *TCategoryPanel) Action() *TAction {
    return ActionFromInst(CategoryPanel_GetAction(c.instance))
}

func (c *TCategoryPanel) SetAction(value IComponent) {
    CategoryPanel_SetAction(c.instance, CheckPtr(value))
}

func (c *TCategoryPanel) Align() TAlign {
    return CategoryPanel_GetAlign(c.instance)
}

func (c *TCategoryPanel) SetAlign(value TAlign) {
    CategoryPanel_SetAlign(c.instance, value)
}

func (c *TCategoryPanel) Anchors() TAnchors {
    return CategoryPanel_GetAnchors(c.instance)
}

func (c *TCategoryPanel) SetAnchors(value TAnchors) {
    CategoryPanel_SetAnchors(c.instance, value)
}

func (c *TCategoryPanel) BoundsRect() TRect {
    return CategoryPanel_GetBoundsRect(c.instance)
}

func (c *TCategoryPanel) SetBoundsRect(value TRect) {
    CategoryPanel_SetBoundsRect(c.instance, value)
}

func (c *TCategoryPanel) ClientHeight() int32 {
    return CategoryPanel_GetClientHeight(c.instance)
}

func (c *TCategoryPanel) SetClientHeight(value int32) {
    CategoryPanel_SetClientHeight(c.instance, value)
}

func (c *TCategoryPanel) ClientRect() TRect {
    return CategoryPanel_GetClientRect(c.instance)
}

func (c *TCategoryPanel) ClientWidth() int32 {
    return CategoryPanel_GetClientWidth(c.instance)
}

func (c *TCategoryPanel) SetClientWidth(value int32) {
    CategoryPanel_SetClientWidth(c.instance, value)
}

func (c *TCategoryPanel) ExplicitLeft() int32 {
    return CategoryPanel_GetExplicitLeft(c.instance)
}

func (c *TCategoryPanel) ExplicitTop() int32 {
    return CategoryPanel_GetExplicitTop(c.instance)
}

func (c *TCategoryPanel) ExplicitWidth() int32 {
    return CategoryPanel_GetExplicitWidth(c.instance)
}

func (c *TCategoryPanel) ExplicitHeight() int32 {
    return CategoryPanel_GetExplicitHeight(c.instance)
}

func (c *TCategoryPanel) Parent() *TControl {
    return ControlFromInst(CategoryPanel_GetParent(c.instance))
}

func (c *TCategoryPanel) SetParent(value IControl) {
    CategoryPanel_SetParent(c.instance, CheckPtr(value))
}

func (c *TCategoryPanel) AlignWithMargins() bool {
    return CategoryPanel_GetAlignWithMargins(c.instance)
}

func (c *TCategoryPanel) SetAlignWithMargins(value bool) {
    CategoryPanel_SetAlignWithMargins(c.instance, value)
}

func (c *TCategoryPanel) Top() int32 {
    return CategoryPanel_GetTop(c.instance)
}

func (c *TCategoryPanel) SetTop(value int32) {
    CategoryPanel_SetTop(c.instance, value)
}

func (c *TCategoryPanel) Cursor() TCursor {
    return CategoryPanel_GetCursor(c.instance)
}

func (c *TCategoryPanel) SetCursor(value TCursor) {
    CategoryPanel_SetCursor(c.instance, value)
}

func (c *TCategoryPanel) Hint() string {
    return CategoryPanel_GetHint(c.instance)
}

func (c *TCategoryPanel) SetHint(value string) {
    CategoryPanel_SetHint(c.instance, value)
}

func (c *TCategoryPanel) Margins() *TMargins {
    return MarginsFromInst(CategoryPanel_GetMargins(c.instance))
}

func (c *TCategoryPanel) SetMargins(value *TMargins) {
    CategoryPanel_SetMargins(c.instance, CheckPtr(value))
}

func (c *TCategoryPanel) CustomHint() *TCustomHint {
    return CustomHintFromInst(CategoryPanel_GetCustomHint(c.instance))
}

func (c *TCategoryPanel) SetCustomHint(value IComponent) {
    CategoryPanel_SetCustomHint(c.instance, CheckPtr(value))
}

func (c *TCategoryPanel) ComponentCount() int32 {
    return CategoryPanel_GetComponentCount(c.instance)
}

func (c *TCategoryPanel) ComponentIndex() int32 {
    return CategoryPanel_GetComponentIndex(c.instance)
}

func (c *TCategoryPanel) SetComponentIndex(value int32) {
    CategoryPanel_SetComponentIndex(c.instance, value)
}

func (c *TCategoryPanel) Owner() *TComponent {
    return ComponentFromInst(CategoryPanel_GetOwner(c.instance))
}

func (c *TCategoryPanel) Name() string {
    return CategoryPanel_GetName(c.instance)
}

func (c *TCategoryPanel) SetName(value string) {
    CategoryPanel_SetName(c.instance, value)
}

func (c *TCategoryPanel) Tag() int {
    return CategoryPanel_GetTag(c.instance)
}

func (c *TCategoryPanel) SetTag(value int) {
    CategoryPanel_SetTag(c.instance, value)
}

func (c *TCategoryPanel) Controls(Index int32) *TControl {
    return ControlFromInst(CategoryPanel_GetControls(c.instance, Index))
}

func (c *TCategoryPanel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CategoryPanel_GetComponents(c.instance, AIndex))
}

