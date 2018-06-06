
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

type TStringGrid struct {
    IWinControl
    instance uintptr
}

func NewStringGrid(owner IComponent) *TStringGrid {
    s := new(TStringGrid)
    s.instance = StringGrid_Create(CheckPtr(owner))
    return s
}

func StringGridFromInst(inst uintptr) *TStringGrid {
    s := new(TStringGrid)
    s.instance = inst
    return s
}

func StringGridFromObj(obj IObject) *TStringGrid {
    s := new(TStringGrid)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStringGrid) Free() {
    if s.instance != 0 {
        StringGrid_Free(s.instance)
        s.instance = 0
    }
}

func (s *TStringGrid) Instance() uintptr {
    return s.instance
}

func (s *TStringGrid) IsValid() bool {
    return s.instance != 0
}

func TStringGridClass() TClass {
    return StringGrid_StaticClassType()
}

func (s *TStringGrid) CellRect(ACol int32, ARow int32) TRect {
    return StringGrid_CellRect(s.instance, ACol , ARow)
}

func (s *TStringGrid) MouseToCell(X int32, Y int32, ACol *int32, ARow *int32) {
    StringGrid_MouseToCell(s.instance, X , Y , ACol , ARow)
}

func (s *TStringGrid) MouseCoord(X int32, Y int32) TGridCoord {
    return StringGrid_MouseCoord(s.instance, X , Y)
}

func (s *TStringGrid) CanFocus() bool {
    return StringGrid_CanFocus(s.instance)
}

func (s *TStringGrid) FlipChildren(AllLevels bool) {
    StringGrid_FlipChildren(s.instance, AllLevels)
}

func (s *TStringGrid) Focused() bool {
    return StringGrid_Focused(s.instance)
}

func (s *TStringGrid) HandleAllocated() bool {
    return StringGrid_HandleAllocated(s.instance)
}

func (s *TStringGrid) Invalidate() {
    StringGrid_Invalidate(s.instance)
}

func (s *TStringGrid) Realign() {
    StringGrid_Realign(s.instance)
}

func (s *TStringGrid) Repaint() {
    StringGrid_Repaint(s.instance)
}

func (s *TStringGrid) ScaleBy(M int32, D int32) {
    StringGrid_ScaleBy(s.instance, M , D)
}

func (s *TStringGrid) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    StringGrid_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

func (s *TStringGrid) SetFocus() {
    StringGrid_SetFocus(s.instance)
}

func (s *TStringGrid) Update() {
    StringGrid_Update(s.instance)
}

func (s *TStringGrid) BringToFront() {
    StringGrid_BringToFront(s.instance)
}

func (s *TStringGrid) ClientToScreen(Point TPoint) TPoint {
    return StringGrid_ClientToScreen(s.instance, Point)
}

func (s *TStringGrid) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return StringGrid_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

func (s *TStringGrid) Dragging() bool {
    return StringGrid_Dragging(s.instance)
}

func (s *TStringGrid) HasParent() bool {
    return StringGrid_HasParent(s.instance)
}

func (s *TStringGrid) Hide() {
    StringGrid_Hide(s.instance)
}

func (s *TStringGrid) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return StringGrid_Perform(s.instance, Msg , WParam , LParam)
}

func (s *TStringGrid) Refresh() {
    StringGrid_Refresh(s.instance)
}

func (s *TStringGrid) ScreenToClient(Point TPoint) TPoint {
    return StringGrid_ScreenToClient(s.instance, Point)
}

func (s *TStringGrid) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return StringGrid_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

func (s *TStringGrid) SendToBack() {
    StringGrid_SendToBack(s.instance)
}

func (s *TStringGrid) Show() {
    StringGrid_Show(s.instance)
}

func (s *TStringGrid) GetTextBuf(Buffer string, BufSize int32) int32 {
    return StringGrid_GetTextBuf(s.instance, Buffer , BufSize)
}

func (s *TStringGrid) GetTextLen() int32 {
    return StringGrid_GetTextLen(s.instance)
}

func (s *TStringGrid) FindComponent(AName string) *TComponent {
    return ComponentFromInst(StringGrid_FindComponent(s.instance, AName))
}

func (s *TStringGrid) GetNamePath() string {
    return StringGrid_GetNamePath(s.instance)
}

func (s *TStringGrid) Assign(Source IObject) {
    StringGrid_Assign(s.instance, CheckPtr(Source))
}

func (s *TStringGrid) DisposeOf() {
    StringGrid_DisposeOf(s.instance)
}

func (s *TStringGrid) ClassType() TClass {
    return StringGrid_ClassType(s.instance)
}

func (s *TStringGrid) ClassName() string {
    return StringGrid_ClassName(s.instance)
}

func (s *TStringGrid) InstanceSize() int32 {
    return StringGrid_InstanceSize(s.instance)
}

func (s *TStringGrid) InheritsFrom(AClass TClass) bool {
    return StringGrid_InheritsFrom(s.instance, AClass)
}

func (s *TStringGrid) Equals(Obj IObject) bool {
    return StringGrid_Equals(s.instance, CheckPtr(Obj))
}

func (s *TStringGrid) GetHashCode() int32 {
    return StringGrid_GetHashCode(s.instance)
}

func (s *TStringGrid) ToString() string {
    return StringGrid_ToString(s.instance)
}

func (s *TStringGrid) Align() TAlign {
    return StringGrid_GetAlign(s.instance)
}

func (s *TStringGrid) SetAlign(value TAlign) {
    StringGrid_SetAlign(s.instance, value)
}

func (s *TStringGrid) Anchors() TAnchors {
    return StringGrid_GetAnchors(s.instance)
}

func (s *TStringGrid) SetAnchors(value TAnchors) {
    StringGrid_SetAnchors(s.instance, value)
}

func (s *TStringGrid) BevelEdges() TBevelEdges {
    return StringGrid_GetBevelEdges(s.instance)
}

func (s *TStringGrid) SetBevelEdges(value TBevelEdges) {
    StringGrid_SetBevelEdges(s.instance, value)
}

func (s *TStringGrid) BevelInner() TBevelCut {
    return StringGrid_GetBevelInner(s.instance)
}

func (s *TStringGrid) SetBevelInner(value TBevelCut) {
    StringGrid_SetBevelInner(s.instance, value)
}

func (s *TStringGrid) BevelKind() TBevelKind {
    return StringGrid_GetBevelKind(s.instance)
}

func (s *TStringGrid) SetBevelKind(value TBevelKind) {
    StringGrid_SetBevelKind(s.instance, value)
}

func (s *TStringGrid) BevelOuter() TBevelCut {
    return StringGrid_GetBevelOuter(s.instance)
}

func (s *TStringGrid) SetBevelOuter(value TBevelCut) {
    StringGrid_SetBevelOuter(s.instance, value)
}

func (s *TStringGrid) BiDiMode() TBiDiMode {
    return StringGrid_GetBiDiMode(s.instance)
}

func (s *TStringGrid) SetBiDiMode(value TBiDiMode) {
    StringGrid_SetBiDiMode(s.instance, value)
}

func (s *TStringGrid) BorderStyle() TBorderStyle {
    return StringGrid_GetBorderStyle(s.instance)
}

func (s *TStringGrid) SetBorderStyle(value TBorderStyle) {
    StringGrid_SetBorderStyle(s.instance, value)
}

func (s *TStringGrid) Color() TColor {
    return StringGrid_GetColor(s.instance)
}

func (s *TStringGrid) SetColor(value TColor) {
    StringGrid_SetColor(s.instance, value)
}

func (s *TStringGrid) ColCount() int32 {
    return StringGrid_GetColCount(s.instance)
}

func (s *TStringGrid) SetColCount(value int32) {
    StringGrid_SetColCount(s.instance, value)
}

func (s *TStringGrid) DefaultColWidth() int32 {
    return StringGrid_GetDefaultColWidth(s.instance)
}

func (s *TStringGrid) SetDefaultColWidth(value int32) {
    StringGrid_SetDefaultColWidth(s.instance, value)
}

func (s *TStringGrid) DefaultRowHeight() int32 {
    return StringGrid_GetDefaultRowHeight(s.instance)
}

func (s *TStringGrid) SetDefaultRowHeight(value int32) {
    StringGrid_SetDefaultRowHeight(s.instance, value)
}

func (s *TStringGrid) DefaultDrawing() bool {
    return StringGrid_GetDefaultDrawing(s.instance)
}

func (s *TStringGrid) SetDefaultDrawing(value bool) {
    StringGrid_SetDefaultDrawing(s.instance, value)
}

func (s *TStringGrid) DoubleBuffered() bool {
    return StringGrid_GetDoubleBuffered(s.instance)
}

func (s *TStringGrid) SetDoubleBuffered(value bool) {
    StringGrid_SetDoubleBuffered(s.instance, value)
}

func (s *TStringGrid) DragCursor() TCursor {
    return StringGrid_GetDragCursor(s.instance)
}

func (s *TStringGrid) SetDragCursor(value TCursor) {
    StringGrid_SetDragCursor(s.instance, value)
}

func (s *TStringGrid) DragKind() TDragKind {
    return StringGrid_GetDragKind(s.instance)
}

func (s *TStringGrid) SetDragKind(value TDragKind) {
    StringGrid_SetDragKind(s.instance, value)
}

func (s *TStringGrid) DragMode() TDragMode {
    return StringGrid_GetDragMode(s.instance)
}

func (s *TStringGrid) SetDragMode(value TDragMode) {
    StringGrid_SetDragMode(s.instance, value)
}

func (s *TStringGrid) DrawingStyle() TGridDrawingStyle {
    return StringGrid_GetDrawingStyle(s.instance)
}

func (s *TStringGrid) SetDrawingStyle(value TGridDrawingStyle) {
    StringGrid_SetDrawingStyle(s.instance, value)
}

func (s *TStringGrid) Enabled() bool {
    return StringGrid_GetEnabled(s.instance)
}

func (s *TStringGrid) SetEnabled(value bool) {
    StringGrid_SetEnabled(s.instance, value)
}

func (s *TStringGrid) FixedColor() TColor {
    return StringGrid_GetFixedColor(s.instance)
}

func (s *TStringGrid) SetFixedColor(value TColor) {
    StringGrid_SetFixedColor(s.instance, value)
}

func (s *TStringGrid) FixedCols() int32 {
    return StringGrid_GetFixedCols(s.instance)
}

func (s *TStringGrid) SetFixedCols(value int32) {
    StringGrid_SetFixedCols(s.instance, value)
}

func (s *TStringGrid) RowCount() int32 {
    return StringGrid_GetRowCount(s.instance)
}

func (s *TStringGrid) SetRowCount(value int32) {
    StringGrid_SetRowCount(s.instance, value)
}

func (s *TStringGrid) FixedRows() int32 {
    return StringGrid_GetFixedRows(s.instance)
}

func (s *TStringGrid) SetFixedRows(value int32) {
    StringGrid_SetFixedRows(s.instance, value)
}

func (s *TStringGrid) Font() *TFont {
    return FontFromInst(StringGrid_GetFont(s.instance))
}

func (s *TStringGrid) SetFont(value *TFont) {
    StringGrid_SetFont(s.instance, CheckPtr(value))
}

func (s *TStringGrid) GradientEndColor() TColor {
    return StringGrid_GetGradientEndColor(s.instance)
}

func (s *TStringGrid) SetGradientEndColor(value TColor) {
    StringGrid_SetGradientEndColor(s.instance, value)
}

func (s *TStringGrid) GradientStartColor() TColor {
    return StringGrid_GetGradientStartColor(s.instance)
}

func (s *TStringGrid) SetGradientStartColor(value TColor) {
    StringGrid_SetGradientStartColor(s.instance, value)
}

func (s *TStringGrid) GridLineWidth() int32 {
    return StringGrid_GetGridLineWidth(s.instance)
}

func (s *TStringGrid) SetGridLineWidth(value int32) {
    StringGrid_SetGridLineWidth(s.instance, value)
}

func (s *TStringGrid) Options() TGridOptions {
    return StringGrid_GetOptions(s.instance)
}

func (s *TStringGrid) SetOptions(value TGridOptions) {
    StringGrid_SetOptions(s.instance, value)
}

func (s *TStringGrid) ParentColor() bool {
    return StringGrid_GetParentColor(s.instance)
}

func (s *TStringGrid) SetParentColor(value bool) {
    StringGrid_SetParentColor(s.instance, value)
}

func (s *TStringGrid) ParentCtl3D() bool {
    return StringGrid_GetParentCtl3D(s.instance)
}

func (s *TStringGrid) SetParentCtl3D(value bool) {
    StringGrid_SetParentCtl3D(s.instance, value)
}

func (s *TStringGrid) ParentDoubleBuffered() bool {
    return StringGrid_GetParentDoubleBuffered(s.instance)
}

func (s *TStringGrid) SetParentDoubleBuffered(value bool) {
    StringGrid_SetParentDoubleBuffered(s.instance, value)
}

func (s *TStringGrid) ParentFont() bool {
    return StringGrid_GetParentFont(s.instance)
}

func (s *TStringGrid) SetParentFont(value bool) {
    StringGrid_SetParentFont(s.instance, value)
}

func (s *TStringGrid) ParentShowHint() bool {
    return StringGrid_GetParentShowHint(s.instance)
}

func (s *TStringGrid) SetParentShowHint(value bool) {
    StringGrid_SetParentShowHint(s.instance, value)
}

func (s *TStringGrid) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(StringGrid_GetPopupMenu(s.instance))
}

func (s *TStringGrid) SetPopupMenu(value IComponent) {
    StringGrid_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TStringGrid) ScrollBars() TScrollStyle {
    return StringGrid_GetScrollBars(s.instance)
}

func (s *TStringGrid) SetScrollBars(value TScrollStyle) {
    StringGrid_SetScrollBars(s.instance, value)
}

func (s *TStringGrid) ShowHint() bool {
    return StringGrid_GetShowHint(s.instance)
}

func (s *TStringGrid) SetShowHint(value bool) {
    StringGrid_SetShowHint(s.instance, value)
}

func (s *TStringGrid) TabOrder() uint16 {
    return StringGrid_GetTabOrder(s.instance)
}

func (s *TStringGrid) SetTabOrder(value uint16) {
    StringGrid_SetTabOrder(s.instance, value)
}

func (s *TStringGrid) Visible() bool {
    return StringGrid_GetVisible(s.instance)
}

func (s *TStringGrid) SetVisible(value bool) {
    StringGrid_SetVisible(s.instance, value)
}

func (s *TStringGrid) StyleElements() TStyleElements {
    return StringGrid_GetStyleElements(s.instance)
}

func (s *TStringGrid) SetStyleElements(value TStyleElements) {
    StringGrid_SetStyleElements(s.instance, value)
}

func (s *TStringGrid) VisibleColCount() int32 {
    return StringGrid_GetVisibleColCount(s.instance)
}

func (s *TStringGrid) VisibleRowCount() int32 {
    return StringGrid_GetVisibleRowCount(s.instance)
}

func (s *TStringGrid) SetOnClick(fn TNotifyEvent) {
    StringGrid_SetOnClick(s.instance, fn)
}

func (s *TStringGrid) SetOnColumnMoved(fn TMovedEvent) {
    StringGrid_SetOnColumnMoved(s.instance, fn)
}

func (s *TStringGrid) SetOnContextPopup(fn TContextPopupEvent) {
    StringGrid_SetOnContextPopup(s.instance, fn)
}

func (s *TStringGrid) SetOnDblClick(fn TNotifyEvent) {
    StringGrid_SetOnDblClick(s.instance, fn)
}

func (s *TStringGrid) SetOnDragDrop(fn TDragDropEvent) {
    StringGrid_SetOnDragDrop(s.instance, fn)
}

func (s *TStringGrid) SetOnDragOver(fn TDragOverEvent) {
    StringGrid_SetOnDragOver(s.instance, fn)
}

func (s *TStringGrid) SetOnDrawCell(fn TDrawCellEvent) {
    StringGrid_SetOnDrawCell(s.instance, fn)
}

func (s *TStringGrid) SetOnEndDock(fn TEndDragEvent) {
    StringGrid_SetOnEndDock(s.instance, fn)
}

func (s *TStringGrid) SetOnEndDrag(fn TEndDragEvent) {
    StringGrid_SetOnEndDrag(s.instance, fn)
}

func (s *TStringGrid) SetOnEnter(fn TNotifyEvent) {
    StringGrid_SetOnEnter(s.instance, fn)
}

func (s *TStringGrid) SetOnExit(fn TNotifyEvent) {
    StringGrid_SetOnExit(s.instance, fn)
}

func (s *TStringGrid) SetOnFixedCellClick(fn TFixedCellClickEvent) {
    StringGrid_SetOnFixedCellClick(s.instance, fn)
}

func (s *TStringGrid) SetOnGetEditMask(fn TGetEditEvent) {
    StringGrid_SetOnGetEditMask(s.instance, fn)
}

func (s *TStringGrid) SetOnGetEditText(fn TGetEditEvent) {
    StringGrid_SetOnGetEditText(s.instance, fn)
}

func (s *TStringGrid) SetOnKeyDown(fn TKeyEvent) {
    StringGrid_SetOnKeyDown(s.instance, fn)
}

func (s *TStringGrid) SetOnKeyPress(fn TKeyPressEvent) {
    StringGrid_SetOnKeyPress(s.instance, fn)
}

func (s *TStringGrid) SetOnKeyUp(fn TKeyEvent) {
    StringGrid_SetOnKeyUp(s.instance, fn)
}

func (s *TStringGrid) SetOnMouseDown(fn TMouseEvent) {
    StringGrid_SetOnMouseDown(s.instance, fn)
}

func (s *TStringGrid) SetOnMouseEnter(fn TNotifyEvent) {
    StringGrid_SetOnMouseEnter(s.instance, fn)
}

func (s *TStringGrid) SetOnMouseLeave(fn TNotifyEvent) {
    StringGrid_SetOnMouseLeave(s.instance, fn)
}

func (s *TStringGrid) SetOnMouseMove(fn TMouseMoveEvent) {
    StringGrid_SetOnMouseMove(s.instance, fn)
}

func (s *TStringGrid) SetOnMouseUp(fn TMouseEvent) {
    StringGrid_SetOnMouseUp(s.instance, fn)
}

func (s *TStringGrid) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    StringGrid_SetOnMouseWheelDown(s.instance, fn)
}

func (s *TStringGrid) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    StringGrid_SetOnMouseWheelUp(s.instance, fn)
}

func (s *TStringGrid) SetOnRowMoved(fn TMovedEvent) {
    StringGrid_SetOnRowMoved(s.instance, fn)
}

func (s *TStringGrid) SetOnSelectCell(fn TSelectCellEvent) {
    StringGrid_SetOnSelectCell(s.instance, fn)
}

func (s *TStringGrid) SetOnSetEditText(fn TSetEditEvent) {
    StringGrid_SetOnSetEditText(s.instance, fn)
}

func (s *TStringGrid) SetOnStartDock(fn TStartDockEvent) {
    StringGrid_SetOnStartDock(s.instance, fn)
}

func (s *TStringGrid) SetOnTopLeftChanged(fn TNotifyEvent) {
    StringGrid_SetOnTopLeftChanged(s.instance, fn)
}

func (s *TStringGrid) Canvas() *TCanvas {
    return CanvasFromInst(StringGrid_GetCanvas(s.instance))
}

func (s *TStringGrid) Col() int32 {
    return StringGrid_GetCol(s.instance)
}

func (s *TStringGrid) SetCol(value int32) {
    StringGrid_SetCol(s.instance, value)
}

func (s *TStringGrid) EditorMode() bool {
    return StringGrid_GetEditorMode(s.instance)
}

func (s *TStringGrid) SetEditorMode(value bool) {
    StringGrid_SetEditorMode(s.instance, value)
}

func (s *TStringGrid) GridHeight() int32 {
    return StringGrid_GetGridHeight(s.instance)
}

func (s *TStringGrid) GridWidth() int32 {
    return StringGrid_GetGridWidth(s.instance)
}

func (s *TStringGrid) LeftCol() int32 {
    return StringGrid_GetLeftCol(s.instance)
}

func (s *TStringGrid) SetLeftCol(value int32) {
    StringGrid_SetLeftCol(s.instance, value)
}

func (s *TStringGrid) Selection() TGridRect {
    return StringGrid_GetSelection(s.instance)
}

func (s *TStringGrid) SetSelection(value TGridRect) {
    StringGrid_SetSelection(s.instance, value)
}

func (s *TStringGrid) Row() int32 {
    return StringGrid_GetRow(s.instance)
}

func (s *TStringGrid) SetRow(value int32) {
    StringGrid_SetRow(s.instance, value)
}

func (s *TStringGrid) TopRow() int32 {
    return StringGrid_GetTopRow(s.instance)
}

func (s *TStringGrid) SetTopRow(value int32) {
    StringGrid_SetTopRow(s.instance, value)
}

func (s *TStringGrid) TabStop() bool {
    return StringGrid_GetTabStop(s.instance)
}

func (s *TStringGrid) SetTabStop(value bool) {
    StringGrid_SetTabStop(s.instance, value)
}

func (s *TStringGrid) DockSite() bool {
    return StringGrid_GetDockSite(s.instance)
}

func (s *TStringGrid) SetDockSite(value bool) {
    StringGrid_SetDockSite(s.instance, value)
}

func (s *TStringGrid) Brush() *TBrush {
    return BrushFromInst(StringGrid_GetBrush(s.instance))
}

func (s *TStringGrid) ControlCount() int32 {
    return StringGrid_GetControlCount(s.instance)
}

func (s *TStringGrid) Handle() HWND {
    return StringGrid_GetHandle(s.instance)
}

func (s *TStringGrid) ParentWindow() HWND {
    return StringGrid_GetParentWindow(s.instance)
}

func (s *TStringGrid) SetParentWindow(value HWND) {
    StringGrid_SetParentWindow(s.instance, value)
}

func (s *TStringGrid) UseDockManager() bool {
    return StringGrid_GetUseDockManager(s.instance)
}

func (s *TStringGrid) SetUseDockManager(value bool) {
    StringGrid_SetUseDockManager(s.instance, value)
}

func (s *TStringGrid) Action() *TAction {
    return ActionFromInst(StringGrid_GetAction(s.instance))
}

func (s *TStringGrid) SetAction(value IComponent) {
    StringGrid_SetAction(s.instance, CheckPtr(value))
}

func (s *TStringGrid) BoundsRect() TRect {
    return StringGrid_GetBoundsRect(s.instance)
}

func (s *TStringGrid) SetBoundsRect(value TRect) {
    StringGrid_SetBoundsRect(s.instance, value)
}

func (s *TStringGrid) ClientHeight() int32 {
    return StringGrid_GetClientHeight(s.instance)
}

func (s *TStringGrid) SetClientHeight(value int32) {
    StringGrid_SetClientHeight(s.instance, value)
}

func (s *TStringGrid) ClientRect() TRect {
    return StringGrid_GetClientRect(s.instance)
}

func (s *TStringGrid) ClientWidth() int32 {
    return StringGrid_GetClientWidth(s.instance)
}

func (s *TStringGrid) SetClientWidth(value int32) {
    StringGrid_SetClientWidth(s.instance, value)
}

func (s *TStringGrid) ExplicitLeft() int32 {
    return StringGrid_GetExplicitLeft(s.instance)
}

func (s *TStringGrid) ExplicitTop() int32 {
    return StringGrid_GetExplicitTop(s.instance)
}

func (s *TStringGrid) ExplicitWidth() int32 {
    return StringGrid_GetExplicitWidth(s.instance)
}

func (s *TStringGrid) ExplicitHeight() int32 {
    return StringGrid_GetExplicitHeight(s.instance)
}

func (s *TStringGrid) Floating() bool {
    return StringGrid_GetFloating(s.instance)
}

func (s *TStringGrid) Parent() *TWinControl {
    return WinControlFromInst(StringGrid_GetParent(s.instance))
}

func (s *TStringGrid) SetParent(value IWinControl) {
    StringGrid_SetParent(s.instance, CheckPtr(value))
}

func (s *TStringGrid) AlignWithMargins() bool {
    return StringGrid_GetAlignWithMargins(s.instance)
}

func (s *TStringGrid) SetAlignWithMargins(value bool) {
    StringGrid_SetAlignWithMargins(s.instance, value)
}

func (s *TStringGrid) Left() int32 {
    return StringGrid_GetLeft(s.instance)
}

func (s *TStringGrid) SetLeft(value int32) {
    StringGrid_SetLeft(s.instance, value)
}

func (s *TStringGrid) Top() int32 {
    return StringGrid_GetTop(s.instance)
}

func (s *TStringGrid) SetTop(value int32) {
    StringGrid_SetTop(s.instance, value)
}

func (s *TStringGrid) Width() int32 {
    return StringGrid_GetWidth(s.instance)
}

func (s *TStringGrid) SetWidth(value int32) {
    StringGrid_SetWidth(s.instance, value)
}

func (s *TStringGrid) Height() int32 {
    return StringGrid_GetHeight(s.instance)
}

func (s *TStringGrid) SetHeight(value int32) {
    StringGrid_SetHeight(s.instance, value)
}

func (s *TStringGrid) Cursor() TCursor {
    return StringGrid_GetCursor(s.instance)
}

func (s *TStringGrid) SetCursor(value TCursor) {
    StringGrid_SetCursor(s.instance, value)
}

func (s *TStringGrid) Hint() string {
    return StringGrid_GetHint(s.instance)
}

func (s *TStringGrid) SetHint(value string) {
    StringGrid_SetHint(s.instance, value)
}

func (s *TStringGrid) Margins() *TMargins {
    return MarginsFromInst(StringGrid_GetMargins(s.instance))
}

func (s *TStringGrid) SetMargins(value *TMargins) {
    StringGrid_SetMargins(s.instance, CheckPtr(value))
}

func (s *TStringGrid) CustomHint() *TCustomHint {
    return CustomHintFromInst(StringGrid_GetCustomHint(s.instance))
}

func (s *TStringGrid) SetCustomHint(value IComponent) {
    StringGrid_SetCustomHint(s.instance, CheckPtr(value))
}

func (s *TStringGrid) ComponentCount() int32 {
    return StringGrid_GetComponentCount(s.instance)
}

func (s *TStringGrid) ComponentIndex() int32 {
    return StringGrid_GetComponentIndex(s.instance)
}

func (s *TStringGrid) SetComponentIndex(value int32) {
    StringGrid_SetComponentIndex(s.instance, value)
}

func (s *TStringGrid) Owner() *TComponent {
    return ComponentFromInst(StringGrid_GetOwner(s.instance))
}

func (s *TStringGrid) Name() string {
    return StringGrid_GetName(s.instance)
}

func (s *TStringGrid) SetName(value string) {
    StringGrid_SetName(s.instance, value)
}

func (s *TStringGrid) Tag() int {
    return StringGrid_GetTag(s.instance)
}

func (s *TStringGrid) SetTag(value int) {
    StringGrid_SetTag(s.instance, value)
}

func (s *TStringGrid) Cells(ACol int32, ARow int32) string {
    return StringGrid_GetCells(s.instance, ACol, ARow)
}

func (s *TStringGrid) SetCells(ACol int32, ARow int32, value string) {
    StringGrid_SetCells(s.instance, ACol, ARow, value)
}

func (s *TStringGrid) Cols(Index int32) *TStrings {
    return StringsFromInst(StringGrid_GetCols(s.instance, Index))
}

func (s *TStringGrid) SetCols(Index int32, value IObject) {
    StringGrid_SetCols(s.instance, Index, CheckPtr(value))
}

func (s *TStringGrid) Objects(ACol int32, ARow int32) *TObject {
    return ObjectFromInst(StringGrid_GetObjects(s.instance, ACol, ARow))
}

func (s *TStringGrid) SetObjects(ACol int32, ARow int32, value IObject) {
    StringGrid_SetObjects(s.instance, ACol, ARow, CheckPtr(value))
}

func (s *TStringGrid) Rows(Index int32) *TStrings {
    return StringsFromInst(StringGrid_GetRows(s.instance, Index))
}

func (s *TStringGrid) SetRows(Index int32, value IObject) {
    StringGrid_SetRows(s.instance, Index, CheckPtr(value))
}

func (s *TStringGrid) ColWidths(Index int32) int32 {
    return StringGrid_GetColWidths(s.instance, Index)
}

func (s *TStringGrid) SetColWidths(Index int32, value int32) {
    StringGrid_SetColWidths(s.instance, Index, value)
}

func (s *TStringGrid) RowHeights(Index int32) int32 {
    return StringGrid_GetRowHeights(s.instance, Index)
}

func (s *TStringGrid) SetRowHeights(Index int32, value int32) {
    StringGrid_SetRowHeights(s.instance, Index, value)
}

func (s *TStringGrid) TabStops(Index int32) bool {
    return StringGrid_GetTabStops(s.instance, Index)
}

func (s *TStringGrid) SetTabStops(Index int32, value bool) {
    StringGrid_SetTabStops(s.instance, Index, value)
}

func (s *TStringGrid) Controls(Index int32) *TControl {
    return ControlFromInst(StringGrid_GetControls(s.instance, Index))
}

func (s *TStringGrid) Components(AIndex int32) *TComponent {
    return ComponentFromInst(StringGrid_GetComponents(s.instance, AIndex))
}

