
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
    "time"
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TMonthCalendar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewMonthCalendar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMonthCalendar(owner IComponent) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = MonthCalendar_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MonthCalendarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MonthCalendarFromInst(inst uintptr) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

// MonthCalendarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MonthCalendarFromObj(obj IObject) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MonthCalendarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MonthCalendarFromUnsafePointer(ptr unsafe.Pointer) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (m *TMonthCalendar) Free() {
    if m.instance != 0 {
        MonthCalendar_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMonthCalendar) Instance() uintptr {
    return m.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMonthCalendar) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMonthCalendar) IsValid() bool {
    return m.instance != 0
}

// TMonthCalendarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMonthCalendarClass() TClass {
    return MonthCalendar_StaticClassType()
}

// CanFocus
func (m *TMonthCalendar) CanFocus() bool {
    return MonthCalendar_CanFocus(m.instance)
}

// ContainsControl
func (m *TMonthCalendar) ContainsControl(Control IControl) bool {
    return MonthCalendar_ContainsControl(m.instance, CheckPtr(Control))
}

// ControlAtPos
func (m *TMonthCalendar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(MonthCalendar_ControlAtPos(m.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
func (m *TMonthCalendar) DisableAlign() {
    MonthCalendar_DisableAlign(m.instance)
}

// EnableAlign
func (m *TMonthCalendar) EnableAlign() {
    MonthCalendar_EnableAlign(m.instance)
}

// FindChildControl
func (m *TMonthCalendar) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(MonthCalendar_FindChildControl(m.instance, ControlName))
}

// FlipChildren
func (m *TMonthCalendar) FlipChildren(AllLevels bool) {
    MonthCalendar_FlipChildren(m.instance, AllLevels)
}

// Focused
func (m *TMonthCalendar) Focused() bool {
    return MonthCalendar_Focused(m.instance)
}

// HandleAllocated
func (m *TMonthCalendar) HandleAllocated() bool {
    return MonthCalendar_HandleAllocated(m.instance)
}

// InsertControl
func (m *TMonthCalendar) InsertControl(AControl IControl) {
    MonthCalendar_InsertControl(m.instance, CheckPtr(AControl))
}

// Invalidate
func (m *TMonthCalendar) Invalidate() {
    MonthCalendar_Invalidate(m.instance)
}

// PaintTo
func (m *TMonthCalendar) PaintTo(DC HDC, X int32, Y int32) {
    MonthCalendar_PaintTo(m.instance, DC , X , Y)
}

// RemoveControl
func (m *TMonthCalendar) RemoveControl(AControl IControl) {
    MonthCalendar_RemoveControl(m.instance, CheckPtr(AControl))
}

// Realign
func (m *TMonthCalendar) Realign() {
    MonthCalendar_Realign(m.instance)
}

// Repaint
func (m *TMonthCalendar) Repaint() {
    MonthCalendar_Repaint(m.instance)
}

// ScaleBy
func (m *TMonthCalendar) ScaleBy(M int32, D int32) {
    MonthCalendar_ScaleBy(m.instance, M , D)
}

// ScrollBy
func (m *TMonthCalendar) ScrollBy(DeltaX int32, DeltaY int32) {
    MonthCalendar_ScrollBy(m.instance, DeltaX , DeltaY)
}

// SetBounds
func (m *TMonthCalendar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    MonthCalendar_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (m *TMonthCalendar) SetFocus() {
    MonthCalendar_SetFocus(m.instance)
}

// Update
func (m *TMonthCalendar) Update() {
    MonthCalendar_Update(m.instance)
}

// UpdateControlState
func (m *TMonthCalendar) UpdateControlState() {
    MonthCalendar_UpdateControlState(m.instance)
}

// BringToFront
func (m *TMonthCalendar) BringToFront() {
    MonthCalendar_BringToFront(m.instance)
}

// ClientToScreen
func (m *TMonthCalendar) ClientToScreen(Point TPoint) TPoint {
    return MonthCalendar_ClientToScreen(m.instance, Point)
}

// ClientToParent
func (m *TMonthCalendar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return MonthCalendar_ClientToParent(m.instance, Point , CheckPtr(AParent))
}

// Dragging
func (m *TMonthCalendar) Dragging() bool {
    return MonthCalendar_Dragging(m.instance)
}

// HasParent
func (m *TMonthCalendar) HasParent() bool {
    return MonthCalendar_HasParent(m.instance)
}

// Hide
func (m *TMonthCalendar) Hide() {
    MonthCalendar_Hide(m.instance)
}

// Perform
func (m *TMonthCalendar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return MonthCalendar_Perform(m.instance, Msg , WParam , LParam)
}

// Refresh
func (m *TMonthCalendar) Refresh() {
    MonthCalendar_Refresh(m.instance)
}

// ScreenToClient
func (m *TMonthCalendar) ScreenToClient(Point TPoint) TPoint {
    return MonthCalendar_ScreenToClient(m.instance, Point)
}

// ParentToClient
func (m *TMonthCalendar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return MonthCalendar_ParentToClient(m.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (m *TMonthCalendar) SendToBack() {
    MonthCalendar_SendToBack(m.instance)
}

// Show
func (m *TMonthCalendar) Show() {
    MonthCalendar_Show(m.instance)
}

// GetTextBuf
func (m *TMonthCalendar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return MonthCalendar_GetTextBuf(m.instance, Buffer , BufSize)
}

// GetTextLen
func (m *TMonthCalendar) GetTextLen() int32 {
    return MonthCalendar_GetTextLen(m.instance)
}

// SetTextBuf
func (m *TMonthCalendar) SetTextBuf(Buffer string) {
    MonthCalendar_SetTextBuf(m.instance, Buffer)
}

// FindComponent
func (m *TMonthCalendar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(MonthCalendar_FindComponent(m.instance, AName))
}

// GetNamePath
func (m *TMonthCalendar) GetNamePath() string {
    return MonthCalendar_GetNamePath(m.instance)
}

// Assign
func (m *TMonthCalendar) Assign(Source IObject) {
    MonthCalendar_Assign(m.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMonthCalendar) DisposeOf() {
    MonthCalendar_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMonthCalendar) ClassType() TClass {
    return MonthCalendar_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMonthCalendar) ClassName() string {
    return MonthCalendar_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMonthCalendar) InstanceSize() int32 {
    return MonthCalendar_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMonthCalendar) InheritsFrom(AClass TClass) bool {
    return MonthCalendar_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMonthCalendar) Equals(Obj IObject) bool {
    return MonthCalendar_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMonthCalendar) GetHashCode() int32 {
    return MonthCalendar_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMonthCalendar) ToString() string {
    return MonthCalendar_ToString(m.instance)
}

// Align
func (m *TMonthCalendar) Align() TAlign {
    return MonthCalendar_GetAlign(m.instance)
}

// SetAlign
func (m *TMonthCalendar) SetAlign(value TAlign) {
    MonthCalendar_SetAlign(m.instance, value)
}

// Anchors
func (m *TMonthCalendar) Anchors() TAnchors {
    return MonthCalendar_GetAnchors(m.instance)
}

// SetAnchors
func (m *TMonthCalendar) SetAnchors(value TAnchors) {
    MonthCalendar_SetAnchors(m.instance, value)
}

// AutoSize
func (m *TMonthCalendar) AutoSize() bool {
    return MonthCalendar_GetAutoSize(m.instance)
}

// SetAutoSize
func (m *TMonthCalendar) SetAutoSize(value bool) {
    MonthCalendar_SetAutoSize(m.instance, value)
}

// BorderWidth
func (m *TMonthCalendar) BorderWidth() int32 {
    return MonthCalendar_GetBorderWidth(m.instance)
}

// SetBorderWidth
func (m *TMonthCalendar) SetBorderWidth(value int32) {
    MonthCalendar_SetBorderWidth(m.instance, value)
}

// BiDiMode
func (m *TMonthCalendar) BiDiMode() TBiDiMode {
    return MonthCalendar_GetBiDiMode(m.instance)
}

// SetBiDiMode
func (m *TMonthCalendar) SetBiDiMode(value TBiDiMode) {
    MonthCalendar_SetBiDiMode(m.instance, value)
}

// CalColors
func (m *TMonthCalendar) CalColors() *TMonthCalColors {
    return MonthCalColorsFromInst(MonthCalendar_GetCalColors(m.instance))
}

// SetCalColors
func (m *TMonthCalendar) SetCalColors(value *TMonthCalColors) {
    MonthCalendar_SetCalColors(m.instance, CheckPtr(value))
}

// MultiSelect
func (m *TMonthCalendar) MultiSelect() bool {
    return MonthCalendar_GetMultiSelect(m.instance)
}

// SetMultiSelect
func (m *TMonthCalendar) SetMultiSelect(value bool) {
    MonthCalendar_SetMultiSelect(m.instance, value)
}

// Date
func (m *TMonthCalendar) Date() time.Time {
    return MonthCalendar_GetDate(m.instance)
}

// SetDate
func (m *TMonthCalendar) SetDate(value time.Time) {
    MonthCalendar_SetDate(m.instance, value)
}

// DoubleBuffered
func (m *TMonthCalendar) DoubleBuffered() bool {
    return MonthCalendar_GetDoubleBuffered(m.instance)
}

// SetDoubleBuffered
func (m *TMonthCalendar) SetDoubleBuffered(value bool) {
    MonthCalendar_SetDoubleBuffered(m.instance, value)
}

// DragCursor
func (m *TMonthCalendar) DragCursor() TCursor {
    return MonthCalendar_GetDragCursor(m.instance)
}

// SetDragCursor
func (m *TMonthCalendar) SetDragCursor(value TCursor) {
    MonthCalendar_SetDragCursor(m.instance, value)
}

// DragKind
func (m *TMonthCalendar) DragKind() TDragKind {
    return MonthCalendar_GetDragKind(m.instance)
}

// SetDragKind
func (m *TMonthCalendar) SetDragKind(value TDragKind) {
    MonthCalendar_SetDragKind(m.instance, value)
}

// DragMode
func (m *TMonthCalendar) DragMode() TDragMode {
    return MonthCalendar_GetDragMode(m.instance)
}

// SetDragMode
func (m *TMonthCalendar) SetDragMode(value TDragMode) {
    MonthCalendar_SetDragMode(m.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (m *TMonthCalendar) Enabled() bool {
    return MonthCalendar_GetEnabled(m.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (m *TMonthCalendar) SetEnabled(value bool) {
    MonthCalendar_SetEnabled(m.instance, value)
}

// FirstDayOfWeek
func (m *TMonthCalendar) FirstDayOfWeek() TCalDayOfWeek {
    return MonthCalendar_GetFirstDayOfWeek(m.instance)
}

// SetFirstDayOfWeek
func (m *TMonthCalendar) SetFirstDayOfWeek(value TCalDayOfWeek) {
    MonthCalendar_SetFirstDayOfWeek(m.instance, value)
}

// Font
func (m *TMonthCalendar) Font() *TFont {
    return FontFromInst(MonthCalendar_GetFont(m.instance))
}

// SetFont
func (m *TMonthCalendar) SetFont(value *TFont) {
    MonthCalendar_SetFont(m.instance, CheckPtr(value))
}

// MaxDate
func (m *TMonthCalendar) MaxDate() time.Time {
    return MonthCalendar_GetMaxDate(m.instance)
}

// SetMaxDate
func (m *TMonthCalendar) SetMaxDate(value time.Time) {
    MonthCalendar_SetMaxDate(m.instance, value)
}

// MaxSelectRange
func (m *TMonthCalendar) MaxSelectRange() int32 {
    return MonthCalendar_GetMaxSelectRange(m.instance)
}

// SetMaxSelectRange
func (m *TMonthCalendar) SetMaxSelectRange(value int32) {
    MonthCalendar_SetMaxSelectRange(m.instance, value)
}

// MinDate
func (m *TMonthCalendar) MinDate() time.Time {
    return MonthCalendar_GetMinDate(m.instance)
}

// SetMinDate
func (m *TMonthCalendar) SetMinDate(value time.Time) {
    MonthCalendar_SetMinDate(m.instance, value)
}

// ParentDoubleBuffered
func (m *TMonthCalendar) ParentDoubleBuffered() bool {
    return MonthCalendar_GetParentDoubleBuffered(m.instance)
}

// SetParentDoubleBuffered
func (m *TMonthCalendar) SetParentDoubleBuffered(value bool) {
    MonthCalendar_SetParentDoubleBuffered(m.instance, value)
}

// ParentFont
func (m *TMonthCalendar) ParentFont() bool {
    return MonthCalendar_GetParentFont(m.instance)
}

// SetParentFont
func (m *TMonthCalendar) SetParentFont(value bool) {
    MonthCalendar_SetParentFont(m.instance, value)
}

// ParentShowHint
func (m *TMonthCalendar) ParentShowHint() bool {
    return MonthCalendar_GetParentShowHint(m.instance)
}

// SetParentShowHint
func (m *TMonthCalendar) SetParentShowHint(value bool) {
    MonthCalendar_SetParentShowHint(m.instance, value)
}

// PopupMenu
func (m *TMonthCalendar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(MonthCalendar_GetPopupMenu(m.instance))
}

// SetPopupMenu
func (m *TMonthCalendar) SetPopupMenu(value IComponent) {
    MonthCalendar_SetPopupMenu(m.instance, CheckPtr(value))
}

// ShowHint
func (m *TMonthCalendar) ShowHint() bool {
    return MonthCalendar_GetShowHint(m.instance)
}

// SetShowHint
func (m *TMonthCalendar) SetShowHint(value bool) {
    MonthCalendar_SetShowHint(m.instance, value)
}

// ShowToday
func (m *TMonthCalendar) ShowToday() bool {
    return MonthCalendar_GetShowToday(m.instance)
}

// SetShowToday
func (m *TMonthCalendar) SetShowToday(value bool) {
    MonthCalendar_SetShowToday(m.instance, value)
}

// ShowTodayCircle
func (m *TMonthCalendar) ShowTodayCircle() bool {
    return MonthCalendar_GetShowTodayCircle(m.instance)
}

// SetShowTodayCircle
func (m *TMonthCalendar) SetShowTodayCircle(value bool) {
    MonthCalendar_SetShowTodayCircle(m.instance, value)
}

// TabOrder
func (m *TMonthCalendar) TabOrder() TTabOrder {
    return MonthCalendar_GetTabOrder(m.instance)
}

// SetTabOrder
func (m *TMonthCalendar) SetTabOrder(value TTabOrder) {
    MonthCalendar_SetTabOrder(m.instance, value)
}

// TabStop
func (m *TMonthCalendar) TabStop() bool {
    return MonthCalendar_GetTabStop(m.instance)
}

// SetTabStop
func (m *TMonthCalendar) SetTabStop(value bool) {
    MonthCalendar_SetTabStop(m.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (m *TMonthCalendar) Visible() bool {
    return MonthCalendar_GetVisible(m.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (m *TMonthCalendar) SetVisible(value bool) {
    MonthCalendar_SetVisible(m.instance, value)
}

// WeekNumbers
func (m *TMonthCalendar) WeekNumbers() bool {
    return MonthCalendar_GetWeekNumbers(m.instance)
}

// SetWeekNumbers
func (m *TMonthCalendar) SetWeekNumbers(value bool) {
    MonthCalendar_SetWeekNumbers(m.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (m *TMonthCalendar) SetOnClick(fn TNotifyEvent) {
    MonthCalendar_SetOnClick(m.instance, fn)
}

// SetOnContextPopup
func (m *TMonthCalendar) SetOnContextPopup(fn TContextPopupEvent) {
    MonthCalendar_SetOnContextPopup(m.instance, fn)
}

// SetOnDblClick
func (m *TMonthCalendar) SetOnDblClick(fn TNotifyEvent) {
    MonthCalendar_SetOnDblClick(m.instance, fn)
}

// SetOnDragDrop
func (m *TMonthCalendar) SetOnDragDrop(fn TDragDropEvent) {
    MonthCalendar_SetOnDragDrop(m.instance, fn)
}

// SetOnDragOver
func (m *TMonthCalendar) SetOnDragOver(fn TDragOverEvent) {
    MonthCalendar_SetOnDragOver(m.instance, fn)
}

// SetOnEndDock
func (m *TMonthCalendar) SetOnEndDock(fn TEndDragEvent) {
    MonthCalendar_SetOnEndDock(m.instance, fn)
}

// SetOnEndDrag
func (m *TMonthCalendar) SetOnEndDrag(fn TEndDragEvent) {
    MonthCalendar_SetOnEndDrag(m.instance, fn)
}

// SetOnEnter
func (m *TMonthCalendar) SetOnEnter(fn TNotifyEvent) {
    MonthCalendar_SetOnEnter(m.instance, fn)
}

// SetOnExit
func (m *TMonthCalendar) SetOnExit(fn TNotifyEvent) {
    MonthCalendar_SetOnExit(m.instance, fn)
}

// SetOnKeyDown
func (m *TMonthCalendar) SetOnKeyDown(fn TKeyEvent) {
    MonthCalendar_SetOnKeyDown(m.instance, fn)
}

// SetOnKeyPress
func (m *TMonthCalendar) SetOnKeyPress(fn TKeyPressEvent) {
    MonthCalendar_SetOnKeyPress(m.instance, fn)
}

// SetOnKeyUp
func (m *TMonthCalendar) SetOnKeyUp(fn TKeyEvent) {
    MonthCalendar_SetOnKeyUp(m.instance, fn)
}

// SetOnMouseEnter
func (m *TMonthCalendar) SetOnMouseEnter(fn TNotifyEvent) {
    MonthCalendar_SetOnMouseEnter(m.instance, fn)
}

// SetOnMouseLeave
func (m *TMonthCalendar) SetOnMouseLeave(fn TNotifyEvent) {
    MonthCalendar_SetOnMouseLeave(m.instance, fn)
}

// SetOnStartDock
func (m *TMonthCalendar) SetOnStartDock(fn TStartDockEvent) {
    MonthCalendar_SetOnStartDock(m.instance, fn)
}

// DockClientCount
func (m *TMonthCalendar) DockClientCount() int32 {
    return MonthCalendar_GetDockClientCount(m.instance)
}

// DockSite
func (m *TMonthCalendar) DockSite() bool {
    return MonthCalendar_GetDockSite(m.instance)
}

// SetDockSite
func (m *TMonthCalendar) SetDockSite(value bool) {
    MonthCalendar_SetDockSite(m.instance, value)
}

// AlignDisabled
func (m *TMonthCalendar) AlignDisabled() bool {
    return MonthCalendar_GetAlignDisabled(m.instance)
}

// MouseInClient
func (m *TMonthCalendar) MouseInClient() bool {
    return MonthCalendar_GetMouseInClient(m.instance)
}

// VisibleDockClientCount
func (m *TMonthCalendar) VisibleDockClientCount() int32 {
    return MonthCalendar_GetVisibleDockClientCount(m.instance)
}

// Brush
func (m *TMonthCalendar) Brush() *TBrush {
    return BrushFromInst(MonthCalendar_GetBrush(m.instance))
}

// ControlCount
func (m *TMonthCalendar) ControlCount() int32 {
    return MonthCalendar_GetControlCount(m.instance)
}

// Handle
func (m *TMonthCalendar) Handle() HWND {
    return MonthCalendar_GetHandle(m.instance)
}

// ParentWindow
func (m *TMonthCalendar) ParentWindow() HWND {
    return MonthCalendar_GetParentWindow(m.instance)
}

// SetParentWindow
func (m *TMonthCalendar) SetParentWindow(value HWND) {
    MonthCalendar_SetParentWindow(m.instance, value)
}

// UseDockManager
func (m *TMonthCalendar) UseDockManager() bool {
    return MonthCalendar_GetUseDockManager(m.instance)
}

// SetUseDockManager
func (m *TMonthCalendar) SetUseDockManager(value bool) {
    MonthCalendar_SetUseDockManager(m.instance, value)
}

// Action
func (m *TMonthCalendar) Action() *TAction {
    return ActionFromInst(MonthCalendar_GetAction(m.instance))
}

// SetAction
func (m *TMonthCalendar) SetAction(value IComponent) {
    MonthCalendar_SetAction(m.instance, CheckPtr(value))
}

// BoundsRect
func (m *TMonthCalendar) BoundsRect() TRect {
    return MonthCalendar_GetBoundsRect(m.instance)
}

// SetBoundsRect
func (m *TMonthCalendar) SetBoundsRect(value TRect) {
    MonthCalendar_SetBoundsRect(m.instance, value)
}

// ClientHeight
func (m *TMonthCalendar) ClientHeight() int32 {
    return MonthCalendar_GetClientHeight(m.instance)
}

// SetClientHeight
func (m *TMonthCalendar) SetClientHeight(value int32) {
    MonthCalendar_SetClientHeight(m.instance, value)
}

// ClientOrigin
func (m *TMonthCalendar) ClientOrigin() TPoint {
    return MonthCalendar_GetClientOrigin(m.instance)
}

// ClientRect
func (m *TMonthCalendar) ClientRect() TRect {
    return MonthCalendar_GetClientRect(m.instance)
}

// ClientWidth
func (m *TMonthCalendar) ClientWidth() int32 {
    return MonthCalendar_GetClientWidth(m.instance)
}

// SetClientWidth
func (m *TMonthCalendar) SetClientWidth(value int32) {
    MonthCalendar_SetClientWidth(m.instance, value)
}

// ControlState
func (m *TMonthCalendar) ControlState() TControlState {
    return MonthCalendar_GetControlState(m.instance)
}

// SetControlState
func (m *TMonthCalendar) SetControlState(value TControlState) {
    MonthCalendar_SetControlState(m.instance, value)
}

// ControlStyle
func (m *TMonthCalendar) ControlStyle() TControlStyle {
    return MonthCalendar_GetControlStyle(m.instance)
}

// SetControlStyle
func (m *TMonthCalendar) SetControlStyle(value TControlStyle) {
    MonthCalendar_SetControlStyle(m.instance, value)
}

// ExplicitLeft
func (m *TMonthCalendar) ExplicitLeft() int32 {
    return MonthCalendar_GetExplicitLeft(m.instance)
}

// ExplicitTop
func (m *TMonthCalendar) ExplicitTop() int32 {
    return MonthCalendar_GetExplicitTop(m.instance)
}

// ExplicitWidth
func (m *TMonthCalendar) ExplicitWidth() int32 {
    return MonthCalendar_GetExplicitWidth(m.instance)
}

// ExplicitHeight
func (m *TMonthCalendar) ExplicitHeight() int32 {
    return MonthCalendar_GetExplicitHeight(m.instance)
}

// Floating
func (m *TMonthCalendar) Floating() bool {
    return MonthCalendar_GetFloating(m.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (m *TMonthCalendar) Parent() *TWinControl {
    return WinControlFromInst(MonthCalendar_GetParent(m.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (m *TMonthCalendar) SetParent(value IWinControl) {
    MonthCalendar_SetParent(m.instance, CheckPtr(value))
}

// StyleElements
func (m *TMonthCalendar) StyleElements() TStyleElements {
    return MonthCalendar_GetStyleElements(m.instance)
}

// SetStyleElements
func (m *TMonthCalendar) SetStyleElements(value TStyleElements) {
    MonthCalendar_SetStyleElements(m.instance, value)
}

// AlignWithMargins
func (m *TMonthCalendar) AlignWithMargins() bool {
    return MonthCalendar_GetAlignWithMargins(m.instance)
}

// SetAlignWithMargins
func (m *TMonthCalendar) SetAlignWithMargins(value bool) {
    MonthCalendar_SetAlignWithMargins(m.instance, value)
}

// Left
func (m *TMonthCalendar) Left() int32 {
    return MonthCalendar_GetLeft(m.instance)
}

// SetLeft
func (m *TMonthCalendar) SetLeft(value int32) {
    MonthCalendar_SetLeft(m.instance, value)
}

// Top
func (m *TMonthCalendar) Top() int32 {
    return MonthCalendar_GetTop(m.instance)
}

// SetTop
func (m *TMonthCalendar) SetTop(value int32) {
    MonthCalendar_SetTop(m.instance, value)
}

// Width
func (m *TMonthCalendar) Width() int32 {
    return MonthCalendar_GetWidth(m.instance)
}

// SetWidth
func (m *TMonthCalendar) SetWidth(value int32) {
    MonthCalendar_SetWidth(m.instance, value)
}

// Height
func (m *TMonthCalendar) Height() int32 {
    return MonthCalendar_GetHeight(m.instance)
}

// SetHeight
func (m *TMonthCalendar) SetHeight(value int32) {
    MonthCalendar_SetHeight(m.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (m *TMonthCalendar) Cursor() TCursor {
    return MonthCalendar_GetCursor(m.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (m *TMonthCalendar) SetCursor(value TCursor) {
    MonthCalendar_SetCursor(m.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (m *TMonthCalendar) Hint() string {
    return MonthCalendar_GetHint(m.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (m *TMonthCalendar) SetHint(value string) {
    MonthCalendar_SetHint(m.instance, value)
}

// Margins
func (m *TMonthCalendar) Margins() *TMargins {
    return MarginsFromInst(MonthCalendar_GetMargins(m.instance))
}

// SetMargins
func (m *TMonthCalendar) SetMargins(value *TMargins) {
    MonthCalendar_SetMargins(m.instance, CheckPtr(value))
}

// CustomHint
func (m *TMonthCalendar) CustomHint() *TCustomHint {
    return CustomHintFromInst(MonthCalendar_GetCustomHint(m.instance))
}

// SetCustomHint
func (m *TMonthCalendar) SetCustomHint(value IComponent) {
    MonthCalendar_SetCustomHint(m.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (m *TMonthCalendar) ComponentCount() int32 {
    return MonthCalendar_GetComponentCount(m.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (m *TMonthCalendar) ComponentIndex() int32 {
    return MonthCalendar_GetComponentIndex(m.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (m *TMonthCalendar) SetComponentIndex(value int32) {
    MonthCalendar_SetComponentIndex(m.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (m *TMonthCalendar) Owner() *TComponent {
    return ComponentFromInst(MonthCalendar_GetOwner(m.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (m *TMonthCalendar) Name() string {
    return MonthCalendar_GetName(m.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (m *TMonthCalendar) SetName(value string) {
    MonthCalendar_SetName(m.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (m *TMonthCalendar) Tag() int {
    return MonthCalendar_GetTag(m.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (m *TMonthCalendar) SetTag(value int) {
    MonthCalendar_SetTag(m.instance, value)
}

// DockClients
func (m *TMonthCalendar) DockClients(Index int32) *TControl {
    return ControlFromInst(MonthCalendar_GetDockClients(m.instance, Index))
}

// Controls
func (m *TMonthCalendar) Controls(Index int32) *TControl {
    return ControlFromInst(MonthCalendar_GetControls(m.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (m *TMonthCalendar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(MonthCalendar_GetComponents(m.instance, AIndex))
}

