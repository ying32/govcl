
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

type TValueListEditor struct {
    IWinControl
    instance uintptr
}

func NewValueListEditor(owner IComponent) *TValueListEditor {
    v := new(TValueListEditor)
    v.instance = ValueListEditor_Create(CheckPtr(owner))
    return v
}

func ValueListEditorFromInst(inst uintptr) *TValueListEditor {
    v := new(TValueListEditor)
    v.instance = inst
    return v
}

func ValueListEditorFromObj(obj IObject) *TValueListEditor {
    v := new(TValueListEditor)
    v.instance = CheckPtr(obj)
    return v
}

func (v *TValueListEditor) Free() {
    if v.instance != 0 {
        ValueListEditor_Free(v.instance)
        v.instance = 0
    }
}

func (v *TValueListEditor) Instance() uintptr {
    return v.instance
}

func (v *TValueListEditor) IsValid() bool {
    return v.instance != 0
}

func TValueListEditorClass() TClass {
    return ValueListEditor_StaticClassType()
}

func (v *TValueListEditor) Refresh() {
    ValueListEditor_Refresh(v.instance)
}

func (v *TValueListEditor) CellRect(ACol int32, ARow int32) TRect {
    return ValueListEditor_CellRect(v.instance, ACol , ARow)
}

func (v *TValueListEditor) MouseToCell(X int32, Y int32, ACol *int32, ARow *int32) {
    ValueListEditor_MouseToCell(v.instance, X , Y , ACol , ARow)
}

func (v *TValueListEditor) MouseCoord(X int32, Y int32) TGridCoord {
    return ValueListEditor_MouseCoord(v.instance, X , Y)
}

func (v *TValueListEditor) CanFocus() bool {
    return ValueListEditor_CanFocus(v.instance)
}

func (v *TValueListEditor) FlipChildren(AllLevels bool) {
    ValueListEditor_FlipChildren(v.instance, AllLevels)
}

func (v *TValueListEditor) Focused() bool {
    return ValueListEditor_Focused(v.instance)
}

func (v *TValueListEditor) HandleAllocated() bool {
    return ValueListEditor_HandleAllocated(v.instance)
}

func (v *TValueListEditor) Invalidate() {
    ValueListEditor_Invalidate(v.instance)
}

func (v *TValueListEditor) Realign() {
    ValueListEditor_Realign(v.instance)
}

func (v *TValueListEditor) Repaint() {
    ValueListEditor_Repaint(v.instance)
}

func (v *TValueListEditor) ScaleBy(M int32, D int32) {
    ValueListEditor_ScaleBy(v.instance, M , D)
}

func (v *TValueListEditor) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ValueListEditor_SetBounds(v.instance, ALeft , ATop , AWidth , AHeight)
}

func (v *TValueListEditor) SetFocus() {
    ValueListEditor_SetFocus(v.instance)
}

func (v *TValueListEditor) Update() {
    ValueListEditor_Update(v.instance)
}

func (v *TValueListEditor) BringToFront() {
    ValueListEditor_BringToFront(v.instance)
}

func (v *TValueListEditor) ClientToScreen(Point TPoint) TPoint {
    return ValueListEditor_ClientToScreen(v.instance, Point)
}

func (v *TValueListEditor) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ValueListEditor_ClientToParent(v.instance, Point , CheckPtr(AParent))
}

func (v *TValueListEditor) Dragging() bool {
    return ValueListEditor_Dragging(v.instance)
}

func (v *TValueListEditor) HasParent() bool {
    return ValueListEditor_HasParent(v.instance)
}

func (v *TValueListEditor) Hide() {
    ValueListEditor_Hide(v.instance)
}

func (v *TValueListEditor) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ValueListEditor_Perform(v.instance, Msg , WParam , LParam)
}

func (v *TValueListEditor) ScreenToClient(Point TPoint) TPoint {
    return ValueListEditor_ScreenToClient(v.instance, Point)
}

func (v *TValueListEditor) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ValueListEditor_ParentToClient(v.instance, Point , CheckPtr(AParent))
}

func (v *TValueListEditor) SendToBack() {
    ValueListEditor_SendToBack(v.instance)
}

func (v *TValueListEditor) Show() {
    ValueListEditor_Show(v.instance)
}

func (v *TValueListEditor) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ValueListEditor_GetTextBuf(v.instance, Buffer , BufSize)
}

func (v *TValueListEditor) GetTextLen() int32 {
    return ValueListEditor_GetTextLen(v.instance)
}

func (v *TValueListEditor) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ValueListEditor_FindComponent(v.instance, AName))
}

func (v *TValueListEditor) GetNamePath() string {
    return ValueListEditor_GetNamePath(v.instance)
}

func (v *TValueListEditor) Assign(Source IObject) {
    ValueListEditor_Assign(v.instance, CheckPtr(Source))
}

func (v *TValueListEditor) DisposeOf() {
    ValueListEditor_DisposeOf(v.instance)
}

func (v *TValueListEditor) ClassType() TClass {
    return ValueListEditor_ClassType(v.instance)
}

func (v *TValueListEditor) ClassName() string {
    return ValueListEditor_ClassName(v.instance)
}

func (v *TValueListEditor) InstanceSize() int32 {
    return ValueListEditor_InstanceSize(v.instance)
}

func (v *TValueListEditor) InheritsFrom(AClass TClass) bool {
    return ValueListEditor_InheritsFrom(v.instance, AClass)
}

func (v *TValueListEditor) Equals(Obj IObject) bool {
    return ValueListEditor_Equals(v.instance, CheckPtr(Obj))
}

func (v *TValueListEditor) GetHashCode() int32 {
    return ValueListEditor_GetHashCode(v.instance)
}

func (v *TValueListEditor) ToString() string {
    return ValueListEditor_ToString(v.instance)
}

func (v *TValueListEditor) ColCount() int32 {
    return ValueListEditor_GetColCount(v.instance)
}

func (v *TValueListEditor) SetColCount(value int32) {
    ValueListEditor_SetColCount(v.instance, value)
}

func (v *TValueListEditor) RowCount() int32 {
    return ValueListEditor_GetRowCount(v.instance)
}

func (v *TValueListEditor) VisibleColCount() int32 {
    return ValueListEditor_GetVisibleColCount(v.instance)
}

func (v *TValueListEditor) VisibleRowCount() int32 {
    return ValueListEditor_GetVisibleRowCount(v.instance)
}

func (v *TValueListEditor) Align() TAlign {
    return ValueListEditor_GetAlign(v.instance)
}

func (v *TValueListEditor) SetAlign(value TAlign) {
    ValueListEditor_SetAlign(v.instance, value)
}

func (v *TValueListEditor) Anchors() TAnchors {
    return ValueListEditor_GetAnchors(v.instance)
}

func (v *TValueListEditor) SetAnchors(value TAnchors) {
    ValueListEditor_SetAnchors(v.instance, value)
}

func (v *TValueListEditor) BiDiMode() TBiDiMode {
    return ValueListEditor_GetBiDiMode(v.instance)
}

func (v *TValueListEditor) SetBiDiMode(value TBiDiMode) {
    ValueListEditor_SetBiDiMode(v.instance, value)
}

func (v *TValueListEditor) BorderStyle() TBorderStyle {
    return ValueListEditor_GetBorderStyle(v.instance)
}

func (v *TValueListEditor) SetBorderStyle(value TBorderStyle) {
    ValueListEditor_SetBorderStyle(v.instance, value)
}

func (v *TValueListEditor) Color() TColor {
    return ValueListEditor_GetColor(v.instance)
}

func (v *TValueListEditor) SetColor(value TColor) {
    ValueListEditor_SetColor(v.instance, value)
}

func (v *TValueListEditor) DefaultColWidth() int32 {
    return ValueListEditor_GetDefaultColWidth(v.instance)
}

func (v *TValueListEditor) SetDefaultColWidth(value int32) {
    ValueListEditor_SetDefaultColWidth(v.instance, value)
}

func (v *TValueListEditor) DefaultDrawing() bool {
    return ValueListEditor_GetDefaultDrawing(v.instance)
}

func (v *TValueListEditor) SetDefaultDrawing(value bool) {
    ValueListEditor_SetDefaultDrawing(v.instance, value)
}

func (v *TValueListEditor) DefaultRowHeight() int32 {
    return ValueListEditor_GetDefaultRowHeight(v.instance)
}

func (v *TValueListEditor) SetDefaultRowHeight(value int32) {
    ValueListEditor_SetDefaultRowHeight(v.instance, value)
}

func (v *TValueListEditor) DoubleBuffered() bool {
    return ValueListEditor_GetDoubleBuffered(v.instance)
}

func (v *TValueListEditor) SetDoubleBuffered(value bool) {
    ValueListEditor_SetDoubleBuffered(v.instance, value)
}

func (v *TValueListEditor) DragCursor() TCursor {
    return ValueListEditor_GetDragCursor(v.instance)
}

func (v *TValueListEditor) SetDragCursor(value TCursor) {
    ValueListEditor_SetDragCursor(v.instance, value)
}

func (v *TValueListEditor) DragKind() TDragKind {
    return ValueListEditor_GetDragKind(v.instance)
}

func (v *TValueListEditor) SetDragKind(value TDragKind) {
    ValueListEditor_SetDragKind(v.instance, value)
}

func (v *TValueListEditor) DragMode() TDragMode {
    return ValueListEditor_GetDragMode(v.instance)
}

func (v *TValueListEditor) SetDragMode(value TDragMode) {
    ValueListEditor_SetDragMode(v.instance, value)
}

func (v *TValueListEditor) DrawingStyle() TGridDrawingStyle {
    return ValueListEditor_GetDrawingStyle(v.instance)
}

func (v *TValueListEditor) SetDrawingStyle(value TGridDrawingStyle) {
    ValueListEditor_SetDrawingStyle(v.instance, value)
}

func (v *TValueListEditor) Enabled() bool {
    return ValueListEditor_GetEnabled(v.instance)
}

func (v *TValueListEditor) SetEnabled(value bool) {
    ValueListEditor_SetEnabled(v.instance, value)
}

func (v *TValueListEditor) FixedColor() TColor {
    return ValueListEditor_GetFixedColor(v.instance)
}

func (v *TValueListEditor) SetFixedColor(value TColor) {
    ValueListEditor_SetFixedColor(v.instance, value)
}

func (v *TValueListEditor) FixedCols() int32 {
    return ValueListEditor_GetFixedCols(v.instance)
}

func (v *TValueListEditor) SetFixedCols(value int32) {
    ValueListEditor_SetFixedCols(v.instance, value)
}

func (v *TValueListEditor) Font() *TFont {
    return FontFromInst(ValueListEditor_GetFont(v.instance))
}

func (v *TValueListEditor) SetFont(value *TFont) {
    ValueListEditor_SetFont(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) GradientEndColor() TColor {
    return ValueListEditor_GetGradientEndColor(v.instance)
}

func (v *TValueListEditor) SetGradientEndColor(value TColor) {
    ValueListEditor_SetGradientEndColor(v.instance, value)
}

func (v *TValueListEditor) GradientStartColor() TColor {
    return ValueListEditor_GetGradientStartColor(v.instance)
}

func (v *TValueListEditor) SetGradientStartColor(value TColor) {
    ValueListEditor_SetGradientStartColor(v.instance, value)
}

func (v *TValueListEditor) GridLineWidth() int32 {
    return ValueListEditor_GetGridLineWidth(v.instance)
}

func (v *TValueListEditor) SetGridLineWidth(value int32) {
    ValueListEditor_SetGridLineWidth(v.instance, value)
}

func (v *TValueListEditor) Options() TGridOptions {
    return ValueListEditor_GetOptions(v.instance)
}

func (v *TValueListEditor) SetOptions(value TGridOptions) {
    ValueListEditor_SetOptions(v.instance, value)
}

func (v *TValueListEditor) ParentColor() bool {
    return ValueListEditor_GetParentColor(v.instance)
}

func (v *TValueListEditor) SetParentColor(value bool) {
    ValueListEditor_SetParentColor(v.instance, value)
}

func (v *TValueListEditor) ParentCtl3D() bool {
    return ValueListEditor_GetParentCtl3D(v.instance)
}

func (v *TValueListEditor) SetParentCtl3D(value bool) {
    ValueListEditor_SetParentCtl3D(v.instance, value)
}

func (v *TValueListEditor) ParentDoubleBuffered() bool {
    return ValueListEditor_GetParentDoubleBuffered(v.instance)
}

func (v *TValueListEditor) SetParentDoubleBuffered(value bool) {
    ValueListEditor_SetParentDoubleBuffered(v.instance, value)
}

func (v *TValueListEditor) ParentFont() bool {
    return ValueListEditor_GetParentFont(v.instance)
}

func (v *TValueListEditor) SetParentFont(value bool) {
    ValueListEditor_SetParentFont(v.instance, value)
}

func (v *TValueListEditor) ParentShowHint() bool {
    return ValueListEditor_GetParentShowHint(v.instance)
}

func (v *TValueListEditor) SetParentShowHint(value bool) {
    ValueListEditor_SetParentShowHint(v.instance, value)
}

func (v *TValueListEditor) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ValueListEditor_GetPopupMenu(v.instance))
}

func (v *TValueListEditor) SetPopupMenu(value IComponent) {
    ValueListEditor_SetPopupMenu(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) ScrollBars() TScrollStyle {
    return ValueListEditor_GetScrollBars(v.instance)
}

func (v *TValueListEditor) SetScrollBars(value TScrollStyle) {
    ValueListEditor_SetScrollBars(v.instance, value)
}

func (v *TValueListEditor) ShowHint() bool {
    return ValueListEditor_GetShowHint(v.instance)
}

func (v *TValueListEditor) SetShowHint(value bool) {
    ValueListEditor_SetShowHint(v.instance, value)
}

func (v *TValueListEditor) Strings() *TStrings {
    return StringsFromInst(ValueListEditor_GetStrings(v.instance))
}

func (v *TValueListEditor) SetStrings(value IObject) {
    ValueListEditor_SetStrings(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) TabOrder() uint16 {
    return ValueListEditor_GetTabOrder(v.instance)
}

func (v *TValueListEditor) SetTabOrder(value uint16) {
    ValueListEditor_SetTabOrder(v.instance, value)
}

func (v *TValueListEditor) Visible() bool {
    return ValueListEditor_GetVisible(v.instance)
}

func (v *TValueListEditor) SetVisible(value bool) {
    ValueListEditor_SetVisible(v.instance, value)
}

func (v *TValueListEditor) StyleElements() TStyleElements {
    return ValueListEditor_GetStyleElements(v.instance)
}

func (v *TValueListEditor) SetStyleElements(value TStyleElements) {
    ValueListEditor_SetStyleElements(v.instance, value)
}

func (v *TValueListEditor) SetOnClick(fn TNotifyEvent) {
    ValueListEditor_SetOnClick(v.instance, fn)
}

func (v *TValueListEditor) SetOnContextPopup(fn TContextPopupEvent) {
    ValueListEditor_SetOnContextPopup(v.instance, fn)
}

func (v *TValueListEditor) SetOnDblClick(fn TNotifyEvent) {
    ValueListEditor_SetOnDblClick(v.instance, fn)
}

func (v *TValueListEditor) SetOnDragDrop(fn TDragDropEvent) {
    ValueListEditor_SetOnDragDrop(v.instance, fn)
}

func (v *TValueListEditor) SetOnDragOver(fn TDragOverEvent) {
    ValueListEditor_SetOnDragOver(v.instance, fn)
}

func (v *TValueListEditor) SetOnDrawCell(fn TDrawCellEvent) {
    ValueListEditor_SetOnDrawCell(v.instance, fn)
}

func (v *TValueListEditor) SetOnEndDock(fn TEndDragEvent) {
    ValueListEditor_SetOnEndDock(v.instance, fn)
}

func (v *TValueListEditor) SetOnEndDrag(fn TEndDragEvent) {
    ValueListEditor_SetOnEndDrag(v.instance, fn)
}

func (v *TValueListEditor) SetOnEnter(fn TNotifyEvent) {
    ValueListEditor_SetOnEnter(v.instance, fn)
}

func (v *TValueListEditor) SetOnExit(fn TNotifyEvent) {
    ValueListEditor_SetOnExit(v.instance, fn)
}

func (v *TValueListEditor) SetOnGetEditMask(fn TGetEditEvent) {
    ValueListEditor_SetOnGetEditMask(v.instance, fn)
}

func (v *TValueListEditor) SetOnGetEditText(fn TGetEditEvent) {
    ValueListEditor_SetOnGetEditText(v.instance, fn)
}

func (v *TValueListEditor) SetOnKeyDown(fn TKeyEvent) {
    ValueListEditor_SetOnKeyDown(v.instance, fn)
}

func (v *TValueListEditor) SetOnKeyPress(fn TKeyPressEvent) {
    ValueListEditor_SetOnKeyPress(v.instance, fn)
}

func (v *TValueListEditor) SetOnKeyUp(fn TKeyEvent) {
    ValueListEditor_SetOnKeyUp(v.instance, fn)
}

func (v *TValueListEditor) SetOnMouseDown(fn TMouseEvent) {
    ValueListEditor_SetOnMouseDown(v.instance, fn)
}

func (v *TValueListEditor) SetOnMouseEnter(fn TNotifyEvent) {
    ValueListEditor_SetOnMouseEnter(v.instance, fn)
}

func (v *TValueListEditor) SetOnMouseLeave(fn TNotifyEvent) {
    ValueListEditor_SetOnMouseLeave(v.instance, fn)
}

func (v *TValueListEditor) SetOnMouseMove(fn TMouseMoveEvent) {
    ValueListEditor_SetOnMouseMove(v.instance, fn)
}

func (v *TValueListEditor) SetOnMouseUp(fn TMouseEvent) {
    ValueListEditor_SetOnMouseUp(v.instance, fn)
}

func (v *TValueListEditor) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    ValueListEditor_SetOnMouseWheelDown(v.instance, fn)
}

func (v *TValueListEditor) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    ValueListEditor_SetOnMouseWheelUp(v.instance, fn)
}

func (v *TValueListEditor) SetOnRowMoved(fn TMovedEvent) {
    ValueListEditor_SetOnRowMoved(v.instance, fn)
}

func (v *TValueListEditor) SetOnSelectCell(fn TSelectCellEvent) {
    ValueListEditor_SetOnSelectCell(v.instance, fn)
}

func (v *TValueListEditor) SetOnSetEditText(fn TSetEditEvent) {
    ValueListEditor_SetOnSetEditText(v.instance, fn)
}

func (v *TValueListEditor) SetOnStartDock(fn TStartDockEvent) {
    ValueListEditor_SetOnStartDock(v.instance, fn)
}

func (v *TValueListEditor) SetOnTopLeftChanged(fn TNotifyEvent) {
    ValueListEditor_SetOnTopLeftChanged(v.instance, fn)
}

func (v *TValueListEditor) Canvas() *TCanvas {
    return CanvasFromInst(ValueListEditor_GetCanvas(v.instance))
}

func (v *TValueListEditor) Col() int32 {
    return ValueListEditor_GetCol(v.instance)
}

func (v *TValueListEditor) SetCol(value int32) {
    ValueListEditor_SetCol(v.instance, value)
}

func (v *TValueListEditor) EditorMode() bool {
    return ValueListEditor_GetEditorMode(v.instance)
}

func (v *TValueListEditor) SetEditorMode(value bool) {
    ValueListEditor_SetEditorMode(v.instance, value)
}

func (v *TValueListEditor) GridHeight() int32 {
    return ValueListEditor_GetGridHeight(v.instance)
}

func (v *TValueListEditor) GridWidth() int32 {
    return ValueListEditor_GetGridWidth(v.instance)
}

func (v *TValueListEditor) LeftCol() int32 {
    return ValueListEditor_GetLeftCol(v.instance)
}

func (v *TValueListEditor) SetLeftCol(value int32) {
    ValueListEditor_SetLeftCol(v.instance, value)
}

func (v *TValueListEditor) Selection() TGridRect {
    return ValueListEditor_GetSelection(v.instance)
}

func (v *TValueListEditor) SetSelection(value TGridRect) {
    ValueListEditor_SetSelection(v.instance, value)
}

func (v *TValueListEditor) Row() int32 {
    return ValueListEditor_GetRow(v.instance)
}

func (v *TValueListEditor) SetRow(value int32) {
    ValueListEditor_SetRow(v.instance, value)
}

func (v *TValueListEditor) TopRow() int32 {
    return ValueListEditor_GetTopRow(v.instance)
}

func (v *TValueListEditor) SetTopRow(value int32) {
    ValueListEditor_SetTopRow(v.instance, value)
}

func (v *TValueListEditor) TabStop() bool {
    return ValueListEditor_GetTabStop(v.instance)
}

func (v *TValueListEditor) SetTabStop(value bool) {
    ValueListEditor_SetTabStop(v.instance, value)
}

func (v *TValueListEditor) DockSite() bool {
    return ValueListEditor_GetDockSite(v.instance)
}

func (v *TValueListEditor) SetDockSite(value bool) {
    ValueListEditor_SetDockSite(v.instance, value)
}

func (v *TValueListEditor) Brush() *TBrush {
    return BrushFromInst(ValueListEditor_GetBrush(v.instance))
}

func (v *TValueListEditor) ControlCount() int32 {
    return ValueListEditor_GetControlCount(v.instance)
}

func (v *TValueListEditor) Handle() HWND {
    return ValueListEditor_GetHandle(v.instance)
}

func (v *TValueListEditor) ParentWindow() HWND {
    return ValueListEditor_GetParentWindow(v.instance)
}

func (v *TValueListEditor) SetParentWindow(value HWND) {
    ValueListEditor_SetParentWindow(v.instance, value)
}

func (v *TValueListEditor) UseDockManager() bool {
    return ValueListEditor_GetUseDockManager(v.instance)
}

func (v *TValueListEditor) SetUseDockManager(value bool) {
    ValueListEditor_SetUseDockManager(v.instance, value)
}

func (v *TValueListEditor) Action() *TAction {
    return ActionFromInst(ValueListEditor_GetAction(v.instance))
}

func (v *TValueListEditor) SetAction(value IComponent) {
    ValueListEditor_SetAction(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) BoundsRect() TRect {
    return ValueListEditor_GetBoundsRect(v.instance)
}

func (v *TValueListEditor) SetBoundsRect(value TRect) {
    ValueListEditor_SetBoundsRect(v.instance, value)
}

func (v *TValueListEditor) ClientHeight() int32 {
    return ValueListEditor_GetClientHeight(v.instance)
}

func (v *TValueListEditor) SetClientHeight(value int32) {
    ValueListEditor_SetClientHeight(v.instance, value)
}

func (v *TValueListEditor) ClientRect() TRect {
    return ValueListEditor_GetClientRect(v.instance)
}

func (v *TValueListEditor) ClientWidth() int32 {
    return ValueListEditor_GetClientWidth(v.instance)
}

func (v *TValueListEditor) SetClientWidth(value int32) {
    ValueListEditor_SetClientWidth(v.instance, value)
}

func (v *TValueListEditor) ExplicitLeft() int32 {
    return ValueListEditor_GetExplicitLeft(v.instance)
}

func (v *TValueListEditor) ExplicitTop() int32 {
    return ValueListEditor_GetExplicitTop(v.instance)
}

func (v *TValueListEditor) ExplicitWidth() int32 {
    return ValueListEditor_GetExplicitWidth(v.instance)
}

func (v *TValueListEditor) ExplicitHeight() int32 {
    return ValueListEditor_GetExplicitHeight(v.instance)
}

func (v *TValueListEditor) Floating() bool {
    return ValueListEditor_GetFloating(v.instance)
}

func (v *TValueListEditor) Parent() *TWinControl {
    return WinControlFromInst(ValueListEditor_GetParent(v.instance))
}

func (v *TValueListEditor) SetParent(value IWinControl) {
    ValueListEditor_SetParent(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) AlignWithMargins() bool {
    return ValueListEditor_GetAlignWithMargins(v.instance)
}

func (v *TValueListEditor) SetAlignWithMargins(value bool) {
    ValueListEditor_SetAlignWithMargins(v.instance, value)
}

func (v *TValueListEditor) Left() int32 {
    return ValueListEditor_GetLeft(v.instance)
}

func (v *TValueListEditor) SetLeft(value int32) {
    ValueListEditor_SetLeft(v.instance, value)
}

func (v *TValueListEditor) Top() int32 {
    return ValueListEditor_GetTop(v.instance)
}

func (v *TValueListEditor) SetTop(value int32) {
    ValueListEditor_SetTop(v.instance, value)
}

func (v *TValueListEditor) Width() int32 {
    return ValueListEditor_GetWidth(v.instance)
}

func (v *TValueListEditor) SetWidth(value int32) {
    ValueListEditor_SetWidth(v.instance, value)
}

func (v *TValueListEditor) Height() int32 {
    return ValueListEditor_GetHeight(v.instance)
}

func (v *TValueListEditor) SetHeight(value int32) {
    ValueListEditor_SetHeight(v.instance, value)
}

func (v *TValueListEditor) Cursor() TCursor {
    return ValueListEditor_GetCursor(v.instance)
}

func (v *TValueListEditor) SetCursor(value TCursor) {
    ValueListEditor_SetCursor(v.instance, value)
}

func (v *TValueListEditor) Hint() string {
    return ValueListEditor_GetHint(v.instance)
}

func (v *TValueListEditor) SetHint(value string) {
    ValueListEditor_SetHint(v.instance, value)
}

func (v *TValueListEditor) Margins() *TMargins {
    return MarginsFromInst(ValueListEditor_GetMargins(v.instance))
}

func (v *TValueListEditor) SetMargins(value *TMargins) {
    ValueListEditor_SetMargins(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) CustomHint() *TCustomHint {
    return CustomHintFromInst(ValueListEditor_GetCustomHint(v.instance))
}

func (v *TValueListEditor) SetCustomHint(value IComponent) {
    ValueListEditor_SetCustomHint(v.instance, CheckPtr(value))
}

func (v *TValueListEditor) ComponentCount() int32 {
    return ValueListEditor_GetComponentCount(v.instance)
}

func (v *TValueListEditor) ComponentIndex() int32 {
    return ValueListEditor_GetComponentIndex(v.instance)
}

func (v *TValueListEditor) SetComponentIndex(value int32) {
    ValueListEditor_SetComponentIndex(v.instance, value)
}

func (v *TValueListEditor) Owner() *TComponent {
    return ComponentFromInst(ValueListEditor_GetOwner(v.instance))
}

func (v *TValueListEditor) Name() string {
    return ValueListEditor_GetName(v.instance)
}

func (v *TValueListEditor) SetName(value string) {
    ValueListEditor_SetName(v.instance, value)
}

func (v *TValueListEditor) Tag() int {
    return ValueListEditor_GetTag(v.instance)
}

func (v *TValueListEditor) SetTag(value int) {
    ValueListEditor_SetTag(v.instance, value)
}

func (v *TValueListEditor) Cells(ACol int32, ARow int32) string {
    return ValueListEditor_GetCells(v.instance, ACol, ARow)
}

func (v *TValueListEditor) SetCells(ACol int32, ARow int32, value string) {
    ValueListEditor_SetCells(v.instance, ACol, ARow, value)
}

func (v *TValueListEditor) Values(Key string) string {
    return ValueListEditor_GetValues(v.instance, Key)
}

func (v *TValueListEditor) SetValues(Key string, value string) {
    ValueListEditor_SetValues(v.instance, Key, value)
}

func (v *TValueListEditor) ColWidths(Index int32) int32 {
    return ValueListEditor_GetColWidths(v.instance, Index)
}

func (v *TValueListEditor) SetColWidths(Index int32, value int32) {
    ValueListEditor_SetColWidths(v.instance, Index, value)
}

func (v *TValueListEditor) RowHeights(Index int32) int32 {
    return ValueListEditor_GetRowHeights(v.instance, Index)
}

func (v *TValueListEditor) SetRowHeights(Index int32, value int32) {
    ValueListEditor_SetRowHeights(v.instance, Index, value)
}

func (v *TValueListEditor) TabStops(Index int32) bool {
    return ValueListEditor_GetTabStops(v.instance, Index)
}

func (v *TValueListEditor) SetTabStops(Index int32, value bool) {
    ValueListEditor_SetTabStops(v.instance, Index, value)
}

func (v *TValueListEditor) Controls(Index int32) *TControl {
    return ControlFromInst(ValueListEditor_GetControls(v.instance, Index))
}

func (v *TValueListEditor) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ValueListEditor_GetComponents(v.instance, AIndex))
}

