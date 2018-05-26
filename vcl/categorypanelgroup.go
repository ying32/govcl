
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

type TCategoryPanelGroup struct {
    IControl
    instance uintptr
}

func NewCategoryPanelGroup(owner IComponent) *TCategoryPanelGroup {
    c := new(TCategoryPanelGroup)
    c.instance = CategoryPanelGroup_Create(CheckPtr(owner))
    return c
}

func CategoryPanelGroupFromInst(inst uintptr) *TCategoryPanelGroup {
    c := new(TCategoryPanelGroup)
    c.instance = inst
    return c
}

func CategoryPanelGroupFromObj(obj IObject) *TCategoryPanelGroup {
    c := new(TCategoryPanelGroup)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCategoryPanelGroup) Free() {
    if c.instance != 0 {
        CategoryPanelGroup_Free(c.instance)
        c.instance = 0
    }
}

func (c *TCategoryPanelGroup) Instance() uintptr {
    return c.instance
}

func (c *TCategoryPanelGroup) IsValid() bool {
    return c.instance != 0
}

func (c *TCategoryPanelGroup) CollapseAll() {
    CategoryPanelGroup_CollapseAll(c.instance)
}

func (c *TCategoryPanelGroup) ExpandAll() {
    CategoryPanelGroup_ExpandAll(c.instance)
}

func (c *TCategoryPanelGroup) CanFocus() bool {
    return CategoryPanelGroup_CanFocus(c.instance)
}

func (c *TCategoryPanelGroup) FlipChildren(AllLevels bool) {
    CategoryPanelGroup_FlipChildren(c.instance, AllLevels)
}

func (c *TCategoryPanelGroup) Focused() bool {
    return CategoryPanelGroup_Focused(c.instance)
}

func (c *TCategoryPanelGroup) HandleAllocated() bool {
    return CategoryPanelGroup_HandleAllocated(c.instance)
}

func (c *TCategoryPanelGroup) Invalidate() {
    CategoryPanelGroup_Invalidate(c.instance)
}

func (c *TCategoryPanelGroup) Realign() {
    CategoryPanelGroup_Realign(c.instance)
}

func (c *TCategoryPanelGroup) Repaint() {
    CategoryPanelGroup_Repaint(c.instance)
}

func (c *TCategoryPanelGroup) ScaleBy(M int32, D int32) {
    CategoryPanelGroup_ScaleBy(c.instance, M , D)
}

func (c *TCategoryPanelGroup) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CategoryPanelGroup_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

func (c *TCategoryPanelGroup) SetFocus() {
    CategoryPanelGroup_SetFocus(c.instance)
}

func (c *TCategoryPanelGroup) Update() {
    CategoryPanelGroup_Update(c.instance)
}

func (c *TCategoryPanelGroup) BringToFront() {
    CategoryPanelGroup_BringToFront(c.instance)
}

func (c *TCategoryPanelGroup) HasParent() bool {
    return CategoryPanelGroup_HasParent(c.instance)
}

func (c *TCategoryPanelGroup) Hide() {
    CategoryPanelGroup_Hide(c.instance)
}

func (c *TCategoryPanelGroup) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CategoryPanelGroup_Perform(c.instance, Msg , WParam , LParam)
}

func (c *TCategoryPanelGroup) Refresh() {
    CategoryPanelGroup_Refresh(c.instance)
}

func (c *TCategoryPanelGroup) SendToBack() {
    CategoryPanelGroup_SendToBack(c.instance)
}

func (c *TCategoryPanelGroup) Show() {
    CategoryPanelGroup_Show(c.instance)
}

func (c *TCategoryPanelGroup) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CategoryPanelGroup_GetTextBuf(c.instance, Buffer , BufSize)
}

func (c *TCategoryPanelGroup) GetTextLen() int32 {
    return CategoryPanelGroup_GetTextLen(c.instance)
}

func (c *TCategoryPanelGroup) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CategoryPanelGroup_FindComponent(c.instance, AName))
}

func (c *TCategoryPanelGroup) GetNamePath() string {
    return CategoryPanelGroup_GetNamePath(c.instance)
}

func (c *TCategoryPanelGroup) Assign(Source IObject) {
    CategoryPanelGroup_Assign(c.instance, CheckPtr(Source))
}

func (c *TCategoryPanelGroup) ClassName() string {
    return CategoryPanelGroup_ClassName(c.instance)
}

func (c *TCategoryPanelGroup) Equals(Obj IObject) bool {
    return CategoryPanelGroup_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCategoryPanelGroup) GetHashCode() int32 {
    return CategoryPanelGroup_GetHashCode(c.instance)
}

func (c *TCategoryPanelGroup) ToString() string {
    return CategoryPanelGroup_ToString(c.instance)
}

func (c *TCategoryPanelGroup) Align() TAlign {
    return CategoryPanelGroup_GetAlign(c.instance)
}

func (c *TCategoryPanelGroup) SetAlign(value TAlign) {
    CategoryPanelGroup_SetAlign(c.instance, value)
}

func (c *TCategoryPanelGroup) Anchors() TAnchors {
    return CategoryPanelGroup_GetAnchors(c.instance)
}

func (c *TCategoryPanelGroup) SetAnchors(value TAnchors) {
    CategoryPanelGroup_SetAnchors(c.instance, value)
}

func (c *TCategoryPanelGroup) BevelEdges() TBevelEdges {
    return CategoryPanelGroup_GetBevelEdges(c.instance)
}

func (c *TCategoryPanelGroup) SetBevelEdges(value TBevelEdges) {
    CategoryPanelGroup_SetBevelEdges(c.instance, value)
}

func (c *TCategoryPanelGroup) BevelInner() TBevelCut {
    return CategoryPanelGroup_GetBevelInner(c.instance)
}

func (c *TCategoryPanelGroup) SetBevelInner(value TBevelCut) {
    CategoryPanelGroup_SetBevelInner(c.instance, value)
}

func (c *TCategoryPanelGroup) BevelOuter() TBevelCut {
    return CategoryPanelGroup_GetBevelOuter(c.instance)
}

func (c *TCategoryPanelGroup) SetBevelOuter(value TBevelCut) {
    CategoryPanelGroup_SetBevelOuter(c.instance, value)
}

func (c *TCategoryPanelGroup) BevelKind() TBevelKind {
    return CategoryPanelGroup_GetBevelKind(c.instance)
}

func (c *TCategoryPanelGroup) SetBevelKind(value TBevelKind) {
    CategoryPanelGroup_SetBevelKind(c.instance, value)
}

func (c *TCategoryPanelGroup) BiDiMode() TBiDiMode {
    return CategoryPanelGroup_GetBiDiMode(c.instance)
}

func (c *TCategoryPanelGroup) SetBiDiMode(value TBiDiMode) {
    CategoryPanelGroup_SetBiDiMode(c.instance, value)
}

func (c *TCategoryPanelGroup) ChevronAlignment() TAlignment {
    return CategoryPanelGroup_GetChevronAlignment(c.instance)
}

func (c *TCategoryPanelGroup) SetChevronAlignment(value TAlignment) {
    CategoryPanelGroup_SetChevronAlignment(c.instance, value)
}

func (c *TCategoryPanelGroup) ChevronColor() TColor {
    return CategoryPanelGroup_GetChevronColor(c.instance)
}

func (c *TCategoryPanelGroup) SetChevronColor(value TColor) {
    CategoryPanelGroup_SetChevronColor(c.instance, value)
}

func (c *TCategoryPanelGroup) ChevronHotColor() TColor {
    return CategoryPanelGroup_GetChevronHotColor(c.instance)
}

func (c *TCategoryPanelGroup) SetChevronHotColor(value TColor) {
    CategoryPanelGroup_SetChevronHotColor(c.instance, value)
}

func (c *TCategoryPanelGroup) DoubleBuffered() bool {
    return CategoryPanelGroup_GetDoubleBuffered(c.instance)
}

func (c *TCategoryPanelGroup) SetDoubleBuffered(value bool) {
    CategoryPanelGroup_SetDoubleBuffered(c.instance, value)
}

func (c *TCategoryPanelGroup) Enabled() bool {
    return CategoryPanelGroup_GetEnabled(c.instance)
}

func (c *TCategoryPanelGroup) SetEnabled(value bool) {
    CategoryPanelGroup_SetEnabled(c.instance, value)
}

func (c *TCategoryPanelGroup) Color() TColor {
    return CategoryPanelGroup_GetColor(c.instance)
}

func (c *TCategoryPanelGroup) SetColor(value TColor) {
    CategoryPanelGroup_SetColor(c.instance, value)
}

func (c *TCategoryPanelGroup) Font() *TFont {
    return FontFromInst(CategoryPanelGroup_GetFont(c.instance))
}

func (c *TCategoryPanelGroup) SetFont(value *TFont) {
    CategoryPanelGroup_SetFont(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) GradientBaseColor() TColor {
    return CategoryPanelGroup_GetGradientBaseColor(c.instance)
}

func (c *TCategoryPanelGroup) SetGradientBaseColor(value TColor) {
    CategoryPanelGroup_SetGradientBaseColor(c.instance, value)
}

func (c *TCategoryPanelGroup) GradientColor() TColor {
    return CategoryPanelGroup_GetGradientColor(c.instance)
}

func (c *TCategoryPanelGroup) SetGradientColor(value TColor) {
    CategoryPanelGroup_SetGradientColor(c.instance, value)
}

func (c *TCategoryPanelGroup) GradientDirection() TGradientDirection {
    return CategoryPanelGroup_GetGradientDirection(c.instance)
}

func (c *TCategoryPanelGroup) SetGradientDirection(value TGradientDirection) {
    CategoryPanelGroup_SetGradientDirection(c.instance, value)
}

func (c *TCategoryPanelGroup) HeaderAlignment() TAlignment {
    return CategoryPanelGroup_GetHeaderAlignment(c.instance)
}

func (c *TCategoryPanelGroup) SetHeaderAlignment(value TAlignment) {
    CategoryPanelGroup_SetHeaderAlignment(c.instance, value)
}

func (c *TCategoryPanelGroup) HeaderFont() *TFont {
    return FontFromInst(CategoryPanelGroup_GetHeaderFont(c.instance))
}

func (c *TCategoryPanelGroup) SetHeaderFont(value *TFont) {
    CategoryPanelGroup_SetHeaderFont(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) HeaderHeight() int32 {
    return CategoryPanelGroup_GetHeaderHeight(c.instance)
}

func (c *TCategoryPanelGroup) SetHeaderHeight(value int32) {
    CategoryPanelGroup_SetHeaderHeight(c.instance, value)
}

func (c *TCategoryPanelGroup) HeaderImage() *TPicture {
    return PictureFromInst(CategoryPanelGroup_GetHeaderImage(c.instance))
}

func (c *TCategoryPanelGroup) SetHeaderImage(value *TPicture) {
    CategoryPanelGroup_SetHeaderImage(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) HeaderStyle() THeaderStyle {
    return CategoryPanelGroup_GetHeaderStyle(c.instance)
}

func (c *TCategoryPanelGroup) SetHeaderStyle(value THeaderStyle) {
    CategoryPanelGroup_SetHeaderStyle(c.instance, value)
}

func (c *TCategoryPanelGroup) Height() int32 {
    return CategoryPanelGroup_GetHeight(c.instance)
}

func (c *TCategoryPanelGroup) SetHeight(value int32) {
    CategoryPanelGroup_SetHeight(c.instance, value)
}

func (c *TCategoryPanelGroup) Images() *TImageList {
    return ImageListFromInst(CategoryPanelGroup_GetImages(c.instance))
}

func (c *TCategoryPanelGroup) SetImages(value IComponent) {
    CategoryPanelGroup_SetImages(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) ParentBackground() bool {
    return CategoryPanelGroup_GetParentBackground(c.instance)
}

func (c *TCategoryPanelGroup) SetParentBackground(value bool) {
    CategoryPanelGroup_SetParentBackground(c.instance, value)
}

func (c *TCategoryPanelGroup) ParentColor() bool {
    return CategoryPanelGroup_GetParentColor(c.instance)
}

func (c *TCategoryPanelGroup) SetParentColor(value bool) {
    CategoryPanelGroup_SetParentColor(c.instance, value)
}

func (c *TCategoryPanelGroup) ParentCtl3D() bool {
    return CategoryPanelGroup_GetParentCtl3D(c.instance)
}

func (c *TCategoryPanelGroup) SetParentCtl3D(value bool) {
    CategoryPanelGroup_SetParentCtl3D(c.instance, value)
}

func (c *TCategoryPanelGroup) ParentDoubleBuffered() bool {
    return CategoryPanelGroup_GetParentDoubleBuffered(c.instance)
}

func (c *TCategoryPanelGroup) SetParentDoubleBuffered(value bool) {
    CategoryPanelGroup_SetParentDoubleBuffered(c.instance, value)
}

func (c *TCategoryPanelGroup) ParentFont() bool {
    return CategoryPanelGroup_GetParentFont(c.instance)
}

func (c *TCategoryPanelGroup) SetParentFont(value bool) {
    CategoryPanelGroup_SetParentFont(c.instance, value)
}

func (c *TCategoryPanelGroup) ParentShowHint() bool {
    return CategoryPanelGroup_GetParentShowHint(c.instance)
}

func (c *TCategoryPanelGroup) SetParentShowHint(value bool) {
    CategoryPanelGroup_SetParentShowHint(c.instance, value)
}

func (c *TCategoryPanelGroup) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CategoryPanelGroup_GetPopupMenu(c.instance))
}

func (c *TCategoryPanelGroup) SetPopupMenu(value IComponent) {
    CategoryPanelGroup_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) ShowHint() bool {
    return CategoryPanelGroup_GetShowHint(c.instance)
}

func (c *TCategoryPanelGroup) SetShowHint(value bool) {
    CategoryPanelGroup_SetShowHint(c.instance, value)
}

func (c *TCategoryPanelGroup) TabOrder() uint16 {
    return CategoryPanelGroup_GetTabOrder(c.instance)
}

func (c *TCategoryPanelGroup) SetTabOrder(value uint16) {
    CategoryPanelGroup_SetTabOrder(c.instance, value)
}

func (c *TCategoryPanelGroup) TabStop() bool {
    return CategoryPanelGroup_GetTabStop(c.instance)
}

func (c *TCategoryPanelGroup) SetTabStop(value bool) {
    CategoryPanelGroup_SetTabStop(c.instance, value)
}

func (c *TCategoryPanelGroup) Visible() bool {
    return CategoryPanelGroup_GetVisible(c.instance)
}

func (c *TCategoryPanelGroup) SetVisible(value bool) {
    CategoryPanelGroup_SetVisible(c.instance, value)
}

func (c *TCategoryPanelGroup) StyleElements() TStyleElements {
    return CategoryPanelGroup_GetStyleElements(c.instance)
}

func (c *TCategoryPanelGroup) SetStyleElements(value TStyleElements) {
    CategoryPanelGroup_SetStyleElements(c.instance, value)
}

func (c *TCategoryPanelGroup) Width() int32 {
    return CategoryPanelGroup_GetWidth(c.instance)
}

func (c *TCategoryPanelGroup) SetWidth(value int32) {
    CategoryPanelGroup_SetWidth(c.instance, value)
}

func (c *TCategoryPanelGroup) SetOnClick(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnClick(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnDblClick(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnDblClick(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnEnter(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnEnter(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnExit(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnExit(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseDown(fn TMouseEvent) {
    CategoryPanelGroup_SetOnMouseDown(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseEnter(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnMouseEnter(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseLeave(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnMouseLeave(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseMove(fn TMouseMoveEvent) {
    CategoryPanelGroup_SetOnMouseMove(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseUp(fn TMouseEvent) {
    CategoryPanelGroup_SetOnMouseUp(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseWheel(fn TMouseWheelEvent) {
    CategoryPanelGroup_SetOnMouseWheel(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnResize(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnResize(c.instance, fn)
}

func (c *TCategoryPanelGroup) Panels() *TList {
    return ListFromInst(CategoryPanelGroup_GetPanels(c.instance))
}

func (c *TCategoryPanelGroup) Brush() *TBrush {
    return BrushFromInst(CategoryPanelGroup_GetBrush(c.instance))
}

func (c *TCategoryPanelGroup) ControlCount() int32 {
    return CategoryPanelGroup_GetControlCount(c.instance)
}

func (c *TCategoryPanelGroup) Handle() HWND {
    return CategoryPanelGroup_GetHandle(c.instance)
}

func (c *TCategoryPanelGroup) ParentWindow() HWND {
    return CategoryPanelGroup_GetParentWindow(c.instance)
}

func (c *TCategoryPanelGroup) SetParentWindow(value HWND) {
    CategoryPanelGroup_SetParentWindow(c.instance, value)
}

func (c *TCategoryPanelGroup) Action() *TAction {
    return ActionFromInst(CategoryPanelGroup_GetAction(c.instance))
}

func (c *TCategoryPanelGroup) SetAction(value IComponent) {
    CategoryPanelGroup_SetAction(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) BoundsRect() TRect {
    return CategoryPanelGroup_GetBoundsRect(c.instance)
}

func (c *TCategoryPanelGroup) SetBoundsRect(value TRect) {
    CategoryPanelGroup_SetBoundsRect(c.instance, value)
}

func (c *TCategoryPanelGroup) ClientHeight() int32 {
    return CategoryPanelGroup_GetClientHeight(c.instance)
}

func (c *TCategoryPanelGroup) SetClientHeight(value int32) {
    CategoryPanelGroup_SetClientHeight(c.instance, value)
}

func (c *TCategoryPanelGroup) ClientRect() TRect {
    return CategoryPanelGroup_GetClientRect(c.instance)
}

func (c *TCategoryPanelGroup) ClientWidth() int32 {
    return CategoryPanelGroup_GetClientWidth(c.instance)
}

func (c *TCategoryPanelGroup) SetClientWidth(value int32) {
    CategoryPanelGroup_SetClientWidth(c.instance, value)
}

func (c *TCategoryPanelGroup) ExplicitLeft() int32 {
    return CategoryPanelGroup_GetExplicitLeft(c.instance)
}

func (c *TCategoryPanelGroup) ExplicitTop() int32 {
    return CategoryPanelGroup_GetExplicitTop(c.instance)
}

func (c *TCategoryPanelGroup) ExplicitWidth() int32 {
    return CategoryPanelGroup_GetExplicitWidth(c.instance)
}

func (c *TCategoryPanelGroup) ExplicitHeight() int32 {
    return CategoryPanelGroup_GetExplicitHeight(c.instance)
}

func (c *TCategoryPanelGroup) Parent() *TControl {
    return ControlFromInst(CategoryPanelGroup_GetParent(c.instance))
}

func (c *TCategoryPanelGroup) SetParent(value IControl) {
    CategoryPanelGroup_SetParent(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) AlignWithMargins() bool {
    return CategoryPanelGroup_GetAlignWithMargins(c.instance)
}

func (c *TCategoryPanelGroup) SetAlignWithMargins(value bool) {
    CategoryPanelGroup_SetAlignWithMargins(c.instance, value)
}

func (c *TCategoryPanelGroup) Left() int32 {
    return CategoryPanelGroup_GetLeft(c.instance)
}

func (c *TCategoryPanelGroup) SetLeft(value int32) {
    CategoryPanelGroup_SetLeft(c.instance, value)
}

func (c *TCategoryPanelGroup) Top() int32 {
    return CategoryPanelGroup_GetTop(c.instance)
}

func (c *TCategoryPanelGroup) SetTop(value int32) {
    CategoryPanelGroup_SetTop(c.instance, value)
}

func (c *TCategoryPanelGroup) Cursor() TCursor {
    return CategoryPanelGroup_GetCursor(c.instance)
}

func (c *TCategoryPanelGroup) SetCursor(value TCursor) {
    CategoryPanelGroup_SetCursor(c.instance, value)
}

func (c *TCategoryPanelGroup) Hint() string {
    return CategoryPanelGroup_GetHint(c.instance)
}

func (c *TCategoryPanelGroup) SetHint(value string) {
    CategoryPanelGroup_SetHint(c.instance, value)
}

func (c *TCategoryPanelGroup) Margins() *TMargins {
    return MarginsFromInst(CategoryPanelGroup_GetMargins(c.instance))
}

func (c *TCategoryPanelGroup) SetMargins(value *TMargins) {
    CategoryPanelGroup_SetMargins(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) CustomHint() *TCustomHint {
    return CustomHintFromInst(CategoryPanelGroup_GetCustomHint(c.instance))
}

func (c *TCategoryPanelGroup) SetCustomHint(value IComponent) {
    CategoryPanelGroup_SetCustomHint(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) ComponentCount() int32 {
    return CategoryPanelGroup_GetComponentCount(c.instance)
}

func (c *TCategoryPanelGroup) ComponentIndex() int32 {
    return CategoryPanelGroup_GetComponentIndex(c.instance)
}

func (c *TCategoryPanelGroup) SetComponentIndex(value int32) {
    CategoryPanelGroup_SetComponentIndex(c.instance, value)
}

func (c *TCategoryPanelGroup) Owner() *TComponent {
    return ComponentFromInst(CategoryPanelGroup_GetOwner(c.instance))
}

func (c *TCategoryPanelGroup) Name() string {
    return CategoryPanelGroup_GetName(c.instance)
}

func (c *TCategoryPanelGroup) SetName(value string) {
    CategoryPanelGroup_SetName(c.instance, value)
}

func (c *TCategoryPanelGroup) Tag() int {
    return CategoryPanelGroup_GetTag(c.instance)
}

func (c *TCategoryPanelGroup) SetTag(value int) {
    CategoryPanelGroup_SetTag(c.instance, value)
}

func (c *TCategoryPanelGroup) Controls(Index int32) *TControl {
    return ControlFromInst(CategoryPanelGroup_GetControls(c.instance, Index))
}

func (c *TCategoryPanelGroup) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CategoryPanelGroup_GetComponents(c.instance, AIndex))
}

