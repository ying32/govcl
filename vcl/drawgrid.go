
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

type TDrawGrid struct {
    IWinControl
    instance uintptr
}

func NewDrawGrid(owner IComponent) *TDrawGrid {
    d := new(TDrawGrid)
    d.instance = DrawGrid_Create(CheckPtr(owner))
    return d
}

func DrawGridFromInst(inst uintptr) *TDrawGrid {
    d := new(TDrawGrid)
    d.instance = inst
    return d
}

func DrawGridFromObj(obj IObject) *TDrawGrid {
    d := new(TDrawGrid)
    d.instance = CheckPtr(obj)
    return d
}

func (d *TDrawGrid) Free() {
    if d.instance != 0 {
        DrawGrid_Free(d.instance)
        d.instance = 0
    }
}

func (d *TDrawGrid) Instance() uintptr {
    return d.instance
}

func (d *TDrawGrid) IsValid() bool {
    return d.instance != 0
}

func TDrawGridClass() TClass {
    return DrawGrid_StaticClassType()
}

func (d *TDrawGrid) CellRect(ACol int32, ARow int32) TRect {
    return DrawGrid_CellRect(d.instance, ACol , ARow)
}

func (d *TDrawGrid) MouseToCell(X int32, Y int32, ACol *int32, ARow *int32) {
    DrawGrid_MouseToCell(d.instance, X , Y , ACol , ARow)
}

func (d *TDrawGrid) MouseCoord(X int32, Y int32) TGridCoord {
    return DrawGrid_MouseCoord(d.instance, X , Y)
}

func (d *TDrawGrid) CanFocus() bool {
    return DrawGrid_CanFocus(d.instance)
}

func (d *TDrawGrid) FlipChildren(AllLevels bool) {
    DrawGrid_FlipChildren(d.instance, AllLevels)
}

func (d *TDrawGrid) Focused() bool {
    return DrawGrid_Focused(d.instance)
}

func (d *TDrawGrid) HandleAllocated() bool {
    return DrawGrid_HandleAllocated(d.instance)
}

func (d *TDrawGrid) Invalidate() {
    DrawGrid_Invalidate(d.instance)
}

func (d *TDrawGrid) Realign() {
    DrawGrid_Realign(d.instance)
}

func (d *TDrawGrid) Repaint() {
    DrawGrid_Repaint(d.instance)
}

func (d *TDrawGrid) ScaleBy(M int32, D int32) {
    DrawGrid_ScaleBy(d.instance, M , D)
}

func (d *TDrawGrid) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    DrawGrid_SetBounds(d.instance, ALeft , ATop , AWidth , AHeight)
}

func (d *TDrawGrid) SetFocus() {
    DrawGrid_SetFocus(d.instance)
}

func (d *TDrawGrid) Update() {
    DrawGrid_Update(d.instance)
}

func (d *TDrawGrid) BringToFront() {
    DrawGrid_BringToFront(d.instance)
}

func (d *TDrawGrid) ClientToScreen(Point TPoint) TPoint {
    return DrawGrid_ClientToScreen(d.instance, Point)
}

func (d *TDrawGrid) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return DrawGrid_ClientToParent(d.instance, Point , CheckPtr(AParent))
}

func (d *TDrawGrid) Dragging() bool {
    return DrawGrid_Dragging(d.instance)
}

func (d *TDrawGrid) HasParent() bool {
    return DrawGrid_HasParent(d.instance)
}

func (d *TDrawGrid) Hide() {
    DrawGrid_Hide(d.instance)
}

func (d *TDrawGrid) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return DrawGrid_Perform(d.instance, Msg , WParam , LParam)
}

func (d *TDrawGrid) Refresh() {
    DrawGrid_Refresh(d.instance)
}

func (d *TDrawGrid) ScreenToClient(Point TPoint) TPoint {
    return DrawGrid_ScreenToClient(d.instance, Point)
}

func (d *TDrawGrid) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return DrawGrid_ParentToClient(d.instance, Point , CheckPtr(AParent))
}

func (d *TDrawGrid) SendToBack() {
    DrawGrid_SendToBack(d.instance)
}

func (d *TDrawGrid) Show() {
    DrawGrid_Show(d.instance)
}

func (d *TDrawGrid) GetTextBuf(Buffer string, BufSize int32) int32 {
    return DrawGrid_GetTextBuf(d.instance, Buffer , BufSize)
}

func (d *TDrawGrid) GetTextLen() int32 {
    return DrawGrid_GetTextLen(d.instance)
}

func (d *TDrawGrid) FindComponent(AName string) *TComponent {
    return ComponentFromInst(DrawGrid_FindComponent(d.instance, AName))
}

func (d *TDrawGrid) GetNamePath() string {
    return DrawGrid_GetNamePath(d.instance)
}

func (d *TDrawGrid) Assign(Source IObject) {
    DrawGrid_Assign(d.instance, CheckPtr(Source))
}

func (d *TDrawGrid) DisposeOf() {
    DrawGrid_DisposeOf(d.instance)
}

func (d *TDrawGrid) ClassType() TClass {
    return DrawGrid_ClassType(d.instance)
}

func (d *TDrawGrid) ClassName() string {
    return DrawGrid_ClassName(d.instance)
}

func (d *TDrawGrid) InstanceSize() int32 {
    return DrawGrid_InstanceSize(d.instance)
}

func (d *TDrawGrid) InheritsFrom(AClass TClass) bool {
    return DrawGrid_InheritsFrom(d.instance, AClass)
}

func (d *TDrawGrid) Equals(Obj IObject) bool {
    return DrawGrid_Equals(d.instance, CheckPtr(Obj))
}

func (d *TDrawGrid) GetHashCode() int32 {
    return DrawGrid_GetHashCode(d.instance)
}

func (d *TDrawGrid) ToString() string {
    return DrawGrid_ToString(d.instance)
}

func (d *TDrawGrid) Align() TAlign {
    return DrawGrid_GetAlign(d.instance)
}

func (d *TDrawGrid) SetAlign(value TAlign) {
    DrawGrid_SetAlign(d.instance, value)
}

func (d *TDrawGrid) Anchors() TAnchors {
    return DrawGrid_GetAnchors(d.instance)
}

func (d *TDrawGrid) SetAnchors(value TAnchors) {
    DrawGrid_SetAnchors(d.instance, value)
}

func (d *TDrawGrid) BevelEdges() TBevelEdges {
    return DrawGrid_GetBevelEdges(d.instance)
}

func (d *TDrawGrid) SetBevelEdges(value TBevelEdges) {
    DrawGrid_SetBevelEdges(d.instance, value)
}

func (d *TDrawGrid) BevelInner() TBevelCut {
    return DrawGrid_GetBevelInner(d.instance)
}

func (d *TDrawGrid) SetBevelInner(value TBevelCut) {
    DrawGrid_SetBevelInner(d.instance, value)
}

func (d *TDrawGrid) BevelKind() TBevelKind {
    return DrawGrid_GetBevelKind(d.instance)
}

func (d *TDrawGrid) SetBevelKind(value TBevelKind) {
    DrawGrid_SetBevelKind(d.instance, value)
}

func (d *TDrawGrid) BevelOuter() TBevelCut {
    return DrawGrid_GetBevelOuter(d.instance)
}

func (d *TDrawGrid) SetBevelOuter(value TBevelCut) {
    DrawGrid_SetBevelOuter(d.instance, value)
}

func (d *TDrawGrid) BiDiMode() TBiDiMode {
    return DrawGrid_GetBiDiMode(d.instance)
}

func (d *TDrawGrid) SetBiDiMode(value TBiDiMode) {
    DrawGrid_SetBiDiMode(d.instance, value)
}

func (d *TDrawGrid) BorderStyle() TBorderStyle {
    return DrawGrid_GetBorderStyle(d.instance)
}

func (d *TDrawGrid) SetBorderStyle(value TBorderStyle) {
    DrawGrid_SetBorderStyle(d.instance, value)
}

func (d *TDrawGrid) Color() TColor {
    return DrawGrid_GetColor(d.instance)
}

func (d *TDrawGrid) SetColor(value TColor) {
    DrawGrid_SetColor(d.instance, value)
}

func (d *TDrawGrid) ColCount() int32 {
    return DrawGrid_GetColCount(d.instance)
}

func (d *TDrawGrid) SetColCount(value int32) {
    DrawGrid_SetColCount(d.instance, value)
}

func (d *TDrawGrid) DefaultColWidth() int32 {
    return DrawGrid_GetDefaultColWidth(d.instance)
}

func (d *TDrawGrid) SetDefaultColWidth(value int32) {
    DrawGrid_SetDefaultColWidth(d.instance, value)
}

func (d *TDrawGrid) DefaultRowHeight() int32 {
    return DrawGrid_GetDefaultRowHeight(d.instance)
}

func (d *TDrawGrid) SetDefaultRowHeight(value int32) {
    DrawGrid_SetDefaultRowHeight(d.instance, value)
}

func (d *TDrawGrid) DefaultDrawing() bool {
    return DrawGrid_GetDefaultDrawing(d.instance)
}

func (d *TDrawGrid) SetDefaultDrawing(value bool) {
    DrawGrid_SetDefaultDrawing(d.instance, value)
}

func (d *TDrawGrid) DoubleBuffered() bool {
    return DrawGrid_GetDoubleBuffered(d.instance)
}

func (d *TDrawGrid) SetDoubleBuffered(value bool) {
    DrawGrid_SetDoubleBuffered(d.instance, value)
}

func (d *TDrawGrid) DragCursor() TCursor {
    return DrawGrid_GetDragCursor(d.instance)
}

func (d *TDrawGrid) SetDragCursor(value TCursor) {
    DrawGrid_SetDragCursor(d.instance, value)
}

func (d *TDrawGrid) DragKind() TDragKind {
    return DrawGrid_GetDragKind(d.instance)
}

func (d *TDrawGrid) SetDragKind(value TDragKind) {
    DrawGrid_SetDragKind(d.instance, value)
}

func (d *TDrawGrid) DragMode() TDragMode {
    return DrawGrid_GetDragMode(d.instance)
}

func (d *TDrawGrid) SetDragMode(value TDragMode) {
    DrawGrid_SetDragMode(d.instance, value)
}

func (d *TDrawGrid) DrawingStyle() TGridDrawingStyle {
    return DrawGrid_GetDrawingStyle(d.instance)
}

func (d *TDrawGrid) SetDrawingStyle(value TGridDrawingStyle) {
    DrawGrid_SetDrawingStyle(d.instance, value)
}

func (d *TDrawGrid) Enabled() bool {
    return DrawGrid_GetEnabled(d.instance)
}

func (d *TDrawGrid) SetEnabled(value bool) {
    DrawGrid_SetEnabled(d.instance, value)
}

func (d *TDrawGrid) FixedColor() TColor {
    return DrawGrid_GetFixedColor(d.instance)
}

func (d *TDrawGrid) SetFixedColor(value TColor) {
    DrawGrid_SetFixedColor(d.instance, value)
}

func (d *TDrawGrid) FixedCols() int32 {
    return DrawGrid_GetFixedCols(d.instance)
}

func (d *TDrawGrid) SetFixedCols(value int32) {
    DrawGrid_SetFixedCols(d.instance, value)
}

func (d *TDrawGrid) RowCount() int32 {
    return DrawGrid_GetRowCount(d.instance)
}

func (d *TDrawGrid) SetRowCount(value int32) {
    DrawGrid_SetRowCount(d.instance, value)
}

func (d *TDrawGrid) FixedRows() int32 {
    return DrawGrid_GetFixedRows(d.instance)
}

func (d *TDrawGrid) SetFixedRows(value int32) {
    DrawGrid_SetFixedRows(d.instance, value)
}

func (d *TDrawGrid) Font() *TFont {
    return FontFromInst(DrawGrid_GetFont(d.instance))
}

func (d *TDrawGrid) SetFont(value *TFont) {
    DrawGrid_SetFont(d.instance, CheckPtr(value))
}

func (d *TDrawGrid) GradientEndColor() TColor {
    return DrawGrid_GetGradientEndColor(d.instance)
}

func (d *TDrawGrid) SetGradientEndColor(value TColor) {
    DrawGrid_SetGradientEndColor(d.instance, value)
}

func (d *TDrawGrid) GradientStartColor() TColor {
    return DrawGrid_GetGradientStartColor(d.instance)
}

func (d *TDrawGrid) SetGradientStartColor(value TColor) {
    DrawGrid_SetGradientStartColor(d.instance, value)
}

func (d *TDrawGrid) GridLineWidth() int32 {
    return DrawGrid_GetGridLineWidth(d.instance)
}

func (d *TDrawGrid) SetGridLineWidth(value int32) {
    DrawGrid_SetGridLineWidth(d.instance, value)
}

func (d *TDrawGrid) Options() TGridOptions {
    return DrawGrid_GetOptions(d.instance)
}

func (d *TDrawGrid) SetOptions(value TGridOptions) {
    DrawGrid_SetOptions(d.instance, value)
}

func (d *TDrawGrid) ParentColor() bool {
    return DrawGrid_GetParentColor(d.instance)
}

func (d *TDrawGrid) SetParentColor(value bool) {
    DrawGrid_SetParentColor(d.instance, value)
}

func (d *TDrawGrid) ParentCtl3D() bool {
    return DrawGrid_GetParentCtl3D(d.instance)
}

func (d *TDrawGrid) SetParentCtl3D(value bool) {
    DrawGrid_SetParentCtl3D(d.instance, value)
}

func (d *TDrawGrid) ParentDoubleBuffered() bool {
    return DrawGrid_GetParentDoubleBuffered(d.instance)
}

func (d *TDrawGrid) SetParentDoubleBuffered(value bool) {
    DrawGrid_SetParentDoubleBuffered(d.instance, value)
}

func (d *TDrawGrid) ParentFont() bool {
    return DrawGrid_GetParentFont(d.instance)
}

func (d *TDrawGrid) SetParentFont(value bool) {
    DrawGrid_SetParentFont(d.instance, value)
}

func (d *TDrawGrid) ParentShowHint() bool {
    return DrawGrid_GetParentShowHint(d.instance)
}

func (d *TDrawGrid) SetParentShowHint(value bool) {
    DrawGrid_SetParentShowHint(d.instance, value)
}

func (d *TDrawGrid) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(DrawGrid_GetPopupMenu(d.instance))
}

func (d *TDrawGrid) SetPopupMenu(value IComponent) {
    DrawGrid_SetPopupMenu(d.instance, CheckPtr(value))
}

func (d *TDrawGrid) ScrollBars() TScrollStyle {
    return DrawGrid_GetScrollBars(d.instance)
}

func (d *TDrawGrid) SetScrollBars(value TScrollStyle) {
    DrawGrid_SetScrollBars(d.instance, value)
}

func (d *TDrawGrid) ShowHint() bool {
    return DrawGrid_GetShowHint(d.instance)
}

func (d *TDrawGrid) SetShowHint(value bool) {
    DrawGrid_SetShowHint(d.instance, value)
}

func (d *TDrawGrid) TabOrder() uint16 {
    return DrawGrid_GetTabOrder(d.instance)
}

func (d *TDrawGrid) SetTabOrder(value uint16) {
    DrawGrid_SetTabOrder(d.instance, value)
}

func (d *TDrawGrid) Visible() bool {
    return DrawGrid_GetVisible(d.instance)
}

func (d *TDrawGrid) SetVisible(value bool) {
    DrawGrid_SetVisible(d.instance, value)
}

func (d *TDrawGrid) StyleElements() TStyleElements {
    return DrawGrid_GetStyleElements(d.instance)
}

func (d *TDrawGrid) SetStyleElements(value TStyleElements) {
    DrawGrid_SetStyleElements(d.instance, value)
}

func (d *TDrawGrid) VisibleColCount() int32 {
    return DrawGrid_GetVisibleColCount(d.instance)
}

func (d *TDrawGrid) VisibleRowCount() int32 {
    return DrawGrid_GetVisibleRowCount(d.instance)
}

func (d *TDrawGrid) SetOnClick(fn TNotifyEvent) {
    DrawGrid_SetOnClick(d.instance, fn)
}

func (d *TDrawGrid) SetOnColumnMoved(fn TMovedEvent) {
    DrawGrid_SetOnColumnMoved(d.instance, fn)
}

func (d *TDrawGrid) SetOnContextPopup(fn TContextPopupEvent) {
    DrawGrid_SetOnContextPopup(d.instance, fn)
}

func (d *TDrawGrid) SetOnDblClick(fn TNotifyEvent) {
    DrawGrid_SetOnDblClick(d.instance, fn)
}

func (d *TDrawGrid) SetOnDragDrop(fn TDragDropEvent) {
    DrawGrid_SetOnDragDrop(d.instance, fn)
}

func (d *TDrawGrid) SetOnDragOver(fn TDragOverEvent) {
    DrawGrid_SetOnDragOver(d.instance, fn)
}

func (d *TDrawGrid) SetOnDrawCell(fn TDrawCellEvent) {
    DrawGrid_SetOnDrawCell(d.instance, fn)
}

func (d *TDrawGrid) SetOnEndDock(fn TEndDragEvent) {
    DrawGrid_SetOnEndDock(d.instance, fn)
}

func (d *TDrawGrid) SetOnEndDrag(fn TEndDragEvent) {
    DrawGrid_SetOnEndDrag(d.instance, fn)
}

func (d *TDrawGrid) SetOnEnter(fn TNotifyEvent) {
    DrawGrid_SetOnEnter(d.instance, fn)
}

func (d *TDrawGrid) SetOnExit(fn TNotifyEvent) {
    DrawGrid_SetOnExit(d.instance, fn)
}

func (d *TDrawGrid) SetOnFixedCellClick(fn TFixedCellClickEvent) {
    DrawGrid_SetOnFixedCellClick(d.instance, fn)
}

func (d *TDrawGrid) SetOnGetEditMask(fn TGetEditEvent) {
    DrawGrid_SetOnGetEditMask(d.instance, fn)
}

func (d *TDrawGrid) SetOnGetEditText(fn TGetEditEvent) {
    DrawGrid_SetOnGetEditText(d.instance, fn)
}

func (d *TDrawGrid) SetOnKeyDown(fn TKeyEvent) {
    DrawGrid_SetOnKeyDown(d.instance, fn)
}

func (d *TDrawGrid) SetOnKeyPress(fn TKeyPressEvent) {
    DrawGrid_SetOnKeyPress(d.instance, fn)
}

func (d *TDrawGrid) SetOnKeyUp(fn TKeyEvent) {
    DrawGrid_SetOnKeyUp(d.instance, fn)
}

func (d *TDrawGrid) SetOnMouseDown(fn TMouseEvent) {
    DrawGrid_SetOnMouseDown(d.instance, fn)
}

func (d *TDrawGrid) SetOnMouseEnter(fn TNotifyEvent) {
    DrawGrid_SetOnMouseEnter(d.instance, fn)
}

func (d *TDrawGrid) SetOnMouseLeave(fn TNotifyEvent) {
    DrawGrid_SetOnMouseLeave(d.instance, fn)
}

func (d *TDrawGrid) SetOnMouseMove(fn TMouseMoveEvent) {
    DrawGrid_SetOnMouseMove(d.instance, fn)
}

func (d *TDrawGrid) SetOnMouseUp(fn TMouseEvent) {
    DrawGrid_SetOnMouseUp(d.instance, fn)
}

func (d *TDrawGrid) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    DrawGrid_SetOnMouseWheelDown(d.instance, fn)
}

func (d *TDrawGrid) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    DrawGrid_SetOnMouseWheelUp(d.instance, fn)
}

func (d *TDrawGrid) SetOnRowMoved(fn TMovedEvent) {
    DrawGrid_SetOnRowMoved(d.instance, fn)
}

func (d *TDrawGrid) SetOnSelectCell(fn TSelectCellEvent) {
    DrawGrid_SetOnSelectCell(d.instance, fn)
}

func (d *TDrawGrid) SetOnSetEditText(fn TSetEditEvent) {
    DrawGrid_SetOnSetEditText(d.instance, fn)
}

func (d *TDrawGrid) SetOnStartDock(fn TStartDockEvent) {
    DrawGrid_SetOnStartDock(d.instance, fn)
}

func (d *TDrawGrid) SetOnTopLeftChanged(fn TNotifyEvent) {
    DrawGrid_SetOnTopLeftChanged(d.instance, fn)
}

func (d *TDrawGrid) Canvas() *TCanvas {
    return CanvasFromInst(DrawGrid_GetCanvas(d.instance))
}

func (d *TDrawGrid) Col() int32 {
    return DrawGrid_GetCol(d.instance)
}

func (d *TDrawGrid) SetCol(value int32) {
    DrawGrid_SetCol(d.instance, value)
}

func (d *TDrawGrid) EditorMode() bool {
    return DrawGrid_GetEditorMode(d.instance)
}

func (d *TDrawGrid) SetEditorMode(value bool) {
    DrawGrid_SetEditorMode(d.instance, value)
}

func (d *TDrawGrid) GridHeight() int32 {
    return DrawGrid_GetGridHeight(d.instance)
}

func (d *TDrawGrid) GridWidth() int32 {
    return DrawGrid_GetGridWidth(d.instance)
}

func (d *TDrawGrid) LeftCol() int32 {
    return DrawGrid_GetLeftCol(d.instance)
}

func (d *TDrawGrid) SetLeftCol(value int32) {
    DrawGrid_SetLeftCol(d.instance, value)
}

func (d *TDrawGrid) Selection() TGridRect {
    return DrawGrid_GetSelection(d.instance)
}

func (d *TDrawGrid) SetSelection(value TGridRect) {
    DrawGrid_SetSelection(d.instance, value)
}

func (d *TDrawGrid) Row() int32 {
    return DrawGrid_GetRow(d.instance)
}

func (d *TDrawGrid) SetRow(value int32) {
    DrawGrid_SetRow(d.instance, value)
}

func (d *TDrawGrid) TopRow() int32 {
    return DrawGrid_GetTopRow(d.instance)
}

func (d *TDrawGrid) SetTopRow(value int32) {
    DrawGrid_SetTopRow(d.instance, value)
}

func (d *TDrawGrid) TabStop() bool {
    return DrawGrid_GetTabStop(d.instance)
}

func (d *TDrawGrid) SetTabStop(value bool) {
    DrawGrid_SetTabStop(d.instance, value)
}

func (d *TDrawGrid) DockSite() bool {
    return DrawGrid_GetDockSite(d.instance)
}

func (d *TDrawGrid) SetDockSite(value bool) {
    DrawGrid_SetDockSite(d.instance, value)
}

func (d *TDrawGrid) Brush() *TBrush {
    return BrushFromInst(DrawGrid_GetBrush(d.instance))
}

func (d *TDrawGrid) ControlCount() int32 {
    return DrawGrid_GetControlCount(d.instance)
}

func (d *TDrawGrid) Handle() HWND {
    return DrawGrid_GetHandle(d.instance)
}

func (d *TDrawGrid) ParentWindow() HWND {
    return DrawGrid_GetParentWindow(d.instance)
}

func (d *TDrawGrid) SetParentWindow(value HWND) {
    DrawGrid_SetParentWindow(d.instance, value)
}

func (d *TDrawGrid) UseDockManager() bool {
    return DrawGrid_GetUseDockManager(d.instance)
}

func (d *TDrawGrid) SetUseDockManager(value bool) {
    DrawGrid_SetUseDockManager(d.instance, value)
}

func (d *TDrawGrid) Action() *TAction {
    return ActionFromInst(DrawGrid_GetAction(d.instance))
}

func (d *TDrawGrid) SetAction(value IComponent) {
    DrawGrid_SetAction(d.instance, CheckPtr(value))
}

func (d *TDrawGrid) BoundsRect() TRect {
    return DrawGrid_GetBoundsRect(d.instance)
}

func (d *TDrawGrid) SetBoundsRect(value TRect) {
    DrawGrid_SetBoundsRect(d.instance, value)
}

func (d *TDrawGrid) ClientHeight() int32 {
    return DrawGrid_GetClientHeight(d.instance)
}

func (d *TDrawGrid) SetClientHeight(value int32) {
    DrawGrid_SetClientHeight(d.instance, value)
}

func (d *TDrawGrid) ClientRect() TRect {
    return DrawGrid_GetClientRect(d.instance)
}

func (d *TDrawGrid) ClientWidth() int32 {
    return DrawGrid_GetClientWidth(d.instance)
}

func (d *TDrawGrid) SetClientWidth(value int32) {
    DrawGrid_SetClientWidth(d.instance, value)
}

func (d *TDrawGrid) ExplicitLeft() int32 {
    return DrawGrid_GetExplicitLeft(d.instance)
}

func (d *TDrawGrid) ExplicitTop() int32 {
    return DrawGrid_GetExplicitTop(d.instance)
}

func (d *TDrawGrid) ExplicitWidth() int32 {
    return DrawGrid_GetExplicitWidth(d.instance)
}

func (d *TDrawGrid) ExplicitHeight() int32 {
    return DrawGrid_GetExplicitHeight(d.instance)
}

func (d *TDrawGrid) Floating() bool {
    return DrawGrid_GetFloating(d.instance)
}

func (d *TDrawGrid) Parent() *TWinControl {
    return WinControlFromInst(DrawGrid_GetParent(d.instance))
}

func (d *TDrawGrid) SetParent(value IWinControl) {
    DrawGrid_SetParent(d.instance, CheckPtr(value))
}

func (d *TDrawGrid) AlignWithMargins() bool {
    return DrawGrid_GetAlignWithMargins(d.instance)
}

func (d *TDrawGrid) SetAlignWithMargins(value bool) {
    DrawGrid_SetAlignWithMargins(d.instance, value)
}

func (d *TDrawGrid) Left() int32 {
    return DrawGrid_GetLeft(d.instance)
}

func (d *TDrawGrid) SetLeft(value int32) {
    DrawGrid_SetLeft(d.instance, value)
}

func (d *TDrawGrid) Top() int32 {
    return DrawGrid_GetTop(d.instance)
}

func (d *TDrawGrid) SetTop(value int32) {
    DrawGrid_SetTop(d.instance, value)
}

func (d *TDrawGrid) Width() int32 {
    return DrawGrid_GetWidth(d.instance)
}

func (d *TDrawGrid) SetWidth(value int32) {
    DrawGrid_SetWidth(d.instance, value)
}

func (d *TDrawGrid) Height() int32 {
    return DrawGrid_GetHeight(d.instance)
}

func (d *TDrawGrid) SetHeight(value int32) {
    DrawGrid_SetHeight(d.instance, value)
}

func (d *TDrawGrid) Cursor() TCursor {
    return DrawGrid_GetCursor(d.instance)
}

func (d *TDrawGrid) SetCursor(value TCursor) {
    DrawGrid_SetCursor(d.instance, value)
}

func (d *TDrawGrid) Hint() string {
    return DrawGrid_GetHint(d.instance)
}

func (d *TDrawGrid) SetHint(value string) {
    DrawGrid_SetHint(d.instance, value)
}

func (d *TDrawGrid) Margins() *TMargins {
    return MarginsFromInst(DrawGrid_GetMargins(d.instance))
}

func (d *TDrawGrid) SetMargins(value *TMargins) {
    DrawGrid_SetMargins(d.instance, CheckPtr(value))
}

func (d *TDrawGrid) CustomHint() *TCustomHint {
    return CustomHintFromInst(DrawGrid_GetCustomHint(d.instance))
}

func (d *TDrawGrid) SetCustomHint(value IComponent) {
    DrawGrid_SetCustomHint(d.instance, CheckPtr(value))
}

func (d *TDrawGrid) ComponentCount() int32 {
    return DrawGrid_GetComponentCount(d.instance)
}

func (d *TDrawGrid) ComponentIndex() int32 {
    return DrawGrid_GetComponentIndex(d.instance)
}

func (d *TDrawGrid) SetComponentIndex(value int32) {
    DrawGrid_SetComponentIndex(d.instance, value)
}

func (d *TDrawGrid) Owner() *TComponent {
    return ComponentFromInst(DrawGrid_GetOwner(d.instance))
}

func (d *TDrawGrid) Name() string {
    return DrawGrid_GetName(d.instance)
}

func (d *TDrawGrid) SetName(value string) {
    DrawGrid_SetName(d.instance, value)
}

func (d *TDrawGrid) Tag() int {
    return DrawGrid_GetTag(d.instance)
}

func (d *TDrawGrid) SetTag(value int) {
    DrawGrid_SetTag(d.instance, value)
}

func (d *TDrawGrid) ColWidths(Index int32) int32 {
    return DrawGrid_GetColWidths(d.instance, Index)
}

func (d *TDrawGrid) SetColWidths(Index int32, value int32) {
    DrawGrid_SetColWidths(d.instance, Index, value)
}

func (d *TDrawGrid) RowHeights(Index int32) int32 {
    return DrawGrid_GetRowHeights(d.instance, Index)
}

func (d *TDrawGrid) SetRowHeights(Index int32, value int32) {
    DrawGrid_SetRowHeights(d.instance, Index, value)
}

func (d *TDrawGrid) TabStops(Index int32) bool {
    return DrawGrid_GetTabStops(d.instance, Index)
}

func (d *TDrawGrid) SetTabStops(Index int32, value bool) {
    DrawGrid_SetTabStops(d.instance, Index, value)
}

func (d *TDrawGrid) Controls(Index int32) *TControl {
    return ControlFromInst(DrawGrid_GetControls(d.instance, Index))
}

func (d *TDrawGrid) Components(AIndex int32) *TComponent {
    return ComponentFromInst(DrawGrid_GetComponents(d.instance, AIndex))
}

