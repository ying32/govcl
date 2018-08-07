
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

type TDateTimePicker struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewDateTimePicker
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewDateTimePicker(owner IComponent) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = DateTimePicker_Create(CheckPtr(owner))
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// DateTimePickerFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func DateTimePickerFromInst(inst uintptr) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = inst
    d.ptr = unsafe.Pointer(inst)
    return d
}

// DateTimePickerFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func DateTimePickerFromObj(obj IObject) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = CheckPtr(obj)
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// DateTimePickerFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func DateTimePickerFromUnsafePointer(ptr unsafe.Pointer) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = uintptr(ptr)
    d.ptr = ptr
    return d
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (d *TDateTimePicker) Free() {
    if d.instance != 0 {
        DateTimePicker_Free(d.instance)
        d.instance = 0
        d.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (d *TDateTimePicker) Instance() uintptr {
    return d.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (d *TDateTimePicker) UnsafeAddr() unsafe.Pointer {
    return d.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (d *TDateTimePicker) IsValid() bool {
    return d.instance != 0
}

// TDateTimePickerClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TDateTimePickerClass() TClass {
    return DateTimePicker_StaticClassType()
}

// CanFocus
func (d *TDateTimePicker) CanFocus() bool {
    return DateTimePicker_CanFocus(d.instance)
}

// FlipChildren
func (d *TDateTimePicker) FlipChildren(AllLevels bool) {
    DateTimePicker_FlipChildren(d.instance, AllLevels)
}

// Focused
func (d *TDateTimePicker) Focused() bool {
    return DateTimePicker_Focused(d.instance)
}

// HandleAllocated
func (d *TDateTimePicker) HandleAllocated() bool {
    return DateTimePicker_HandleAllocated(d.instance)
}

// Invalidate
func (d *TDateTimePicker) Invalidate() {
    DateTimePicker_Invalidate(d.instance)
}

// Realign
func (d *TDateTimePicker) Realign() {
    DateTimePicker_Realign(d.instance)
}

// Repaint
func (d *TDateTimePicker) Repaint() {
    DateTimePicker_Repaint(d.instance)
}

// ScaleBy
func (d *TDateTimePicker) ScaleBy(M int32, D int32) {
    DateTimePicker_ScaleBy(d.instance, M , D)
}

// SetBounds
func (d *TDateTimePicker) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    DateTimePicker_SetBounds(d.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
func (d *TDateTimePicker) SetFocus() {
    DateTimePicker_SetFocus(d.instance)
}

// Update
func (d *TDateTimePicker) Update() {
    DateTimePicker_Update(d.instance)
}

// BringToFront
func (d *TDateTimePicker) BringToFront() {
    DateTimePicker_BringToFront(d.instance)
}

// ClientToScreen
func (d *TDateTimePicker) ClientToScreen(Point TPoint) TPoint {
    return DateTimePicker_ClientToScreen(d.instance, Point)
}

// ClientToParent
func (d *TDateTimePicker) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return DateTimePicker_ClientToParent(d.instance, Point , CheckPtr(AParent))
}

// Dragging
func (d *TDateTimePicker) Dragging() bool {
    return DateTimePicker_Dragging(d.instance)
}

// HasParent
func (d *TDateTimePicker) HasParent() bool {
    return DateTimePicker_HasParent(d.instance)
}

// Hide
func (d *TDateTimePicker) Hide() {
    DateTimePicker_Hide(d.instance)
}

// Perform
func (d *TDateTimePicker) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return DateTimePicker_Perform(d.instance, Msg , WParam , LParam)
}

// Refresh
func (d *TDateTimePicker) Refresh() {
    DateTimePicker_Refresh(d.instance)
}

// ScreenToClient
func (d *TDateTimePicker) ScreenToClient(Point TPoint) TPoint {
    return DateTimePicker_ScreenToClient(d.instance, Point)
}

// ParentToClient
func (d *TDateTimePicker) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return DateTimePicker_ParentToClient(d.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (d *TDateTimePicker) SendToBack() {
    DateTimePicker_SendToBack(d.instance)
}

// Show
func (d *TDateTimePicker) Show() {
    DateTimePicker_Show(d.instance)
}

// GetTextBuf
func (d *TDateTimePicker) GetTextBuf(Buffer string, BufSize int32) int32 {
    return DateTimePicker_GetTextBuf(d.instance, Buffer , BufSize)
}

// GetTextLen
func (d *TDateTimePicker) GetTextLen() int32 {
    return DateTimePicker_GetTextLen(d.instance)
}

// FindComponent
func (d *TDateTimePicker) FindComponent(AName string) *TComponent {
    return ComponentFromInst(DateTimePicker_FindComponent(d.instance, AName))
}

// GetNamePath
func (d *TDateTimePicker) GetNamePath() string {
    return DateTimePicker_GetNamePath(d.instance)
}

// Assign
func (d *TDateTimePicker) Assign(Source IObject) {
    DateTimePicker_Assign(d.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (d *TDateTimePicker) DisposeOf() {
    DateTimePicker_DisposeOf(d.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (d *TDateTimePicker) ClassType() TClass {
    return DateTimePicker_ClassType(d.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (d *TDateTimePicker) ClassName() string {
    return DateTimePicker_ClassName(d.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (d *TDateTimePicker) InstanceSize() int32 {
    return DateTimePicker_InstanceSize(d.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (d *TDateTimePicker) InheritsFrom(AClass TClass) bool {
    return DateTimePicker_InheritsFrom(d.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (d *TDateTimePicker) Equals(Obj IObject) bool {
    return DateTimePicker_Equals(d.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (d *TDateTimePicker) GetHashCode() int32 {
    return DateTimePicker_GetHashCode(d.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (d *TDateTimePicker) ToString() string {
    return DateTimePicker_ToString(d.instance)
}

// DateTime
func (d *TDateTimePicker) DateTime() time.Time {
    return DateTimePicker_GetDateTime(d.instance)
}

// SetDateTime
func (d *TDateTimePicker) SetDateTime(value time.Time) {
    DateTimePicker_SetDateTime(d.instance, value)
}

// DroppedDown
func (d *TDateTimePicker) DroppedDown() bool {
    return DateTimePicker_GetDroppedDown(d.instance)
}

// Align
func (d *TDateTimePicker) Align() TAlign {
    return DateTimePicker_GetAlign(d.instance)
}

// SetAlign
func (d *TDateTimePicker) SetAlign(value TAlign) {
    DateTimePicker_SetAlign(d.instance, value)
}

// Anchors
func (d *TDateTimePicker) Anchors() TAnchors {
    return DateTimePicker_GetAnchors(d.instance)
}

// SetAnchors
func (d *TDateTimePicker) SetAnchors(value TAnchors) {
    DateTimePicker_SetAnchors(d.instance, value)
}

// BevelEdges
func (d *TDateTimePicker) BevelEdges() TBevelEdges {
    return DateTimePicker_GetBevelEdges(d.instance)
}

// SetBevelEdges
func (d *TDateTimePicker) SetBevelEdges(value TBevelEdges) {
    DateTimePicker_SetBevelEdges(d.instance, value)
}

// BevelInner
func (d *TDateTimePicker) BevelInner() TBevelCut {
    return DateTimePicker_GetBevelInner(d.instance)
}

// SetBevelInner
func (d *TDateTimePicker) SetBevelInner(value TBevelCut) {
    DateTimePicker_SetBevelInner(d.instance, value)
}

// BevelOuter
func (d *TDateTimePicker) BevelOuter() TBevelCut {
    return DateTimePicker_GetBevelOuter(d.instance)
}

// SetBevelOuter
func (d *TDateTimePicker) SetBevelOuter(value TBevelCut) {
    DateTimePicker_SetBevelOuter(d.instance, value)
}

// BevelKind
func (d *TDateTimePicker) BevelKind() TBevelKind {
    return DateTimePicker_GetBevelKind(d.instance)
}

// SetBevelKind
func (d *TDateTimePicker) SetBevelKind(value TBevelKind) {
    DateTimePicker_SetBevelKind(d.instance, value)
}

// BiDiMode
func (d *TDateTimePicker) BiDiMode() TBiDiMode {
    return DateTimePicker_GetBiDiMode(d.instance)
}

// SetBiDiMode
func (d *TDateTimePicker) SetBiDiMode(value TBiDiMode) {
    DateTimePicker_SetBiDiMode(d.instance, value)
}

// CalAlignment
func (d *TDateTimePicker) CalAlignment() TDTCalAlignment {
    return DateTimePicker_GetCalAlignment(d.instance)
}

// SetCalAlignment
func (d *TDateTimePicker) SetCalAlignment(value TDTCalAlignment) {
    DateTimePicker_SetCalAlignment(d.instance, value)
}

// CalColors
func (d *TDateTimePicker) CalColors() *TMonthCalColors {
    return MonthCalColorsFromInst(DateTimePicker_GetCalColors(d.instance))
}

// SetCalColors
func (d *TDateTimePicker) SetCalColors(value *TMonthCalColors) {
    DateTimePicker_SetCalColors(d.instance, CheckPtr(value))
}

// Date
func (d *TDateTimePicker) Date() time.Time {
    return DateTimePicker_GetDate(d.instance)
}

// SetDate
func (d *TDateTimePicker) SetDate(value time.Time) {
    DateTimePicker_SetDate(d.instance, value)
}

// Format
func (d *TDateTimePicker) Format() string {
    return DateTimePicker_GetFormat(d.instance)
}

// SetFormat
func (d *TDateTimePicker) SetFormat(value string) {
    DateTimePicker_SetFormat(d.instance, value)
}

// Time
func (d *TDateTimePicker) Time() time.Time {
    return DateTimePicker_GetTime(d.instance)
}

// SetTime
func (d *TDateTimePicker) SetTime(value time.Time) {
    DateTimePicker_SetTime(d.instance, value)
}

// Checked
func (d *TDateTimePicker) Checked() bool {
    return DateTimePicker_GetChecked(d.instance)
}

// SetChecked
func (d *TDateTimePicker) SetChecked(value bool) {
    DateTimePicker_SetChecked(d.instance, value)
}

// Color
func (d *TDateTimePicker) Color() TColor {
    return DateTimePicker_GetColor(d.instance)
}

// SetColor
func (d *TDateTimePicker) SetColor(value TColor) {
    DateTimePicker_SetColor(d.instance, value)
}

// DateFormat
func (d *TDateTimePicker) DateFormat() TDTDateFormat {
    return DateTimePicker_GetDateFormat(d.instance)
}

// SetDateFormat
func (d *TDateTimePicker) SetDateFormat(value TDTDateFormat) {
    DateTimePicker_SetDateFormat(d.instance, value)
}

// DateMode
func (d *TDateTimePicker) DateMode() TDTDateMode {
    return DateTimePicker_GetDateMode(d.instance)
}

// SetDateMode
func (d *TDateTimePicker) SetDateMode(value TDTDateMode) {
    DateTimePicker_SetDateMode(d.instance, value)
}

// DoubleBuffered
func (d *TDateTimePicker) DoubleBuffered() bool {
    return DateTimePicker_GetDoubleBuffered(d.instance)
}

// SetDoubleBuffered
func (d *TDateTimePicker) SetDoubleBuffered(value bool) {
    DateTimePicker_SetDoubleBuffered(d.instance, value)
}

// DragCursor
func (d *TDateTimePicker) DragCursor() TCursor {
    return DateTimePicker_GetDragCursor(d.instance)
}

// SetDragCursor
func (d *TDateTimePicker) SetDragCursor(value TCursor) {
    DateTimePicker_SetDragCursor(d.instance, value)
}

// DragKind
func (d *TDateTimePicker) DragKind() TDragKind {
    return DateTimePicker_GetDragKind(d.instance)
}

// SetDragKind
func (d *TDateTimePicker) SetDragKind(value TDragKind) {
    DateTimePicker_SetDragKind(d.instance, value)
}

// DragMode
func (d *TDateTimePicker) DragMode() TDragMode {
    return DateTimePicker_GetDragMode(d.instance)
}

// SetDragMode
func (d *TDateTimePicker) SetDragMode(value TDragMode) {
    DateTimePicker_SetDragMode(d.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (d *TDateTimePicker) Enabled() bool {
    return DateTimePicker_GetEnabled(d.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (d *TDateTimePicker) SetEnabled(value bool) {
    DateTimePicker_SetEnabled(d.instance, value)
}

// Font
func (d *TDateTimePicker) Font() *TFont {
    return FontFromInst(DateTimePicker_GetFont(d.instance))
}

// SetFont
func (d *TDateTimePicker) SetFont(value *TFont) {
    DateTimePicker_SetFont(d.instance, CheckPtr(value))
}

// Kind
func (d *TDateTimePicker) Kind() TDateTimeKind {
    return DateTimePicker_GetKind(d.instance)
}

// SetKind
func (d *TDateTimePicker) SetKind(value TDateTimeKind) {
    DateTimePicker_SetKind(d.instance, value)
}

// MaxDate
func (d *TDateTimePicker) MaxDate() time.Time {
    return DateTimePicker_GetMaxDate(d.instance)
}

// SetMaxDate
func (d *TDateTimePicker) SetMaxDate(value time.Time) {
    DateTimePicker_SetMaxDate(d.instance, value)
}

// MinDate
func (d *TDateTimePicker) MinDate() time.Time {
    return DateTimePicker_GetMinDate(d.instance)
}

// SetMinDate
func (d *TDateTimePicker) SetMinDate(value time.Time) {
    DateTimePicker_SetMinDate(d.instance, value)
}

// ParseInput
func (d *TDateTimePicker) ParseInput() bool {
    return DateTimePicker_GetParseInput(d.instance)
}

// SetParseInput
func (d *TDateTimePicker) SetParseInput(value bool) {
    DateTimePicker_SetParseInput(d.instance, value)
}

// ParentColor
func (d *TDateTimePicker) ParentColor() bool {
    return DateTimePicker_GetParentColor(d.instance)
}

// SetParentColor
func (d *TDateTimePicker) SetParentColor(value bool) {
    DateTimePicker_SetParentColor(d.instance, value)
}

// ParentDoubleBuffered
func (d *TDateTimePicker) ParentDoubleBuffered() bool {
    return DateTimePicker_GetParentDoubleBuffered(d.instance)
}

// SetParentDoubleBuffered
func (d *TDateTimePicker) SetParentDoubleBuffered(value bool) {
    DateTimePicker_SetParentDoubleBuffered(d.instance, value)
}

// ParentFont
func (d *TDateTimePicker) ParentFont() bool {
    return DateTimePicker_GetParentFont(d.instance)
}

// SetParentFont
func (d *TDateTimePicker) SetParentFont(value bool) {
    DateTimePicker_SetParentFont(d.instance, value)
}

// ParentShowHint
func (d *TDateTimePicker) ParentShowHint() bool {
    return DateTimePicker_GetParentShowHint(d.instance)
}

// SetParentShowHint
func (d *TDateTimePicker) SetParentShowHint(value bool) {
    DateTimePicker_SetParentShowHint(d.instance, value)
}

// PopupMenu
func (d *TDateTimePicker) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(DateTimePicker_GetPopupMenu(d.instance))
}

// SetPopupMenu
func (d *TDateTimePicker) SetPopupMenu(value IComponent) {
    DateTimePicker_SetPopupMenu(d.instance, CheckPtr(value))
}

// ShowHint
func (d *TDateTimePicker) ShowHint() bool {
    return DateTimePicker_GetShowHint(d.instance)
}

// SetShowHint
func (d *TDateTimePicker) SetShowHint(value bool) {
    DateTimePicker_SetShowHint(d.instance, value)
}

// TabOrder
func (d *TDateTimePicker) TabOrder() uint16 {
    return DateTimePicker_GetTabOrder(d.instance)
}

// SetTabOrder
func (d *TDateTimePicker) SetTabOrder(value uint16) {
    DateTimePicker_SetTabOrder(d.instance, value)
}

// TabStop
func (d *TDateTimePicker) TabStop() bool {
    return DateTimePicker_GetTabStop(d.instance)
}

// SetTabStop
func (d *TDateTimePicker) SetTabStop(value bool) {
    DateTimePicker_SetTabStop(d.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (d *TDateTimePicker) Visible() bool {
    return DateTimePicker_GetVisible(d.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (d *TDateTimePicker) SetVisible(value bool) {
    DateTimePicker_SetVisible(d.instance, value)
}

// StyleElements
func (d *TDateTimePicker) StyleElements() TStyleElements {
    return DateTimePicker_GetStyleElements(d.instance)
}

// SetStyleElements
func (d *TDateTimePicker) SetStyleElements(value TStyleElements) {
    DateTimePicker_SetStyleElements(d.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (d *TDateTimePicker) SetOnClick(fn TNotifyEvent) {
    DateTimePicker_SetOnClick(d.instance, fn)
}

// SetOnChange
func (d *TDateTimePicker) SetOnChange(fn TNotifyEvent) {
    DateTimePicker_SetOnChange(d.instance, fn)
}

// SetOnContextPopup
func (d *TDateTimePicker) SetOnContextPopup(fn TContextPopupEvent) {
    DateTimePicker_SetOnContextPopup(d.instance, fn)
}

// SetOnDragDrop
func (d *TDateTimePicker) SetOnDragDrop(fn TDragDropEvent) {
    DateTimePicker_SetOnDragDrop(d.instance, fn)
}

// SetOnDragOver
func (d *TDateTimePicker) SetOnDragOver(fn TDragOverEvent) {
    DateTimePicker_SetOnDragOver(d.instance, fn)
}

// SetOnEndDock
func (d *TDateTimePicker) SetOnEndDock(fn TEndDragEvent) {
    DateTimePicker_SetOnEndDock(d.instance, fn)
}

// SetOnEndDrag
func (d *TDateTimePicker) SetOnEndDrag(fn TEndDragEvent) {
    DateTimePicker_SetOnEndDrag(d.instance, fn)
}

// SetOnEnter
func (d *TDateTimePicker) SetOnEnter(fn TNotifyEvent) {
    DateTimePicker_SetOnEnter(d.instance, fn)
}

// SetOnExit
func (d *TDateTimePicker) SetOnExit(fn TNotifyEvent) {
    DateTimePicker_SetOnExit(d.instance, fn)
}

// SetOnKeyDown
func (d *TDateTimePicker) SetOnKeyDown(fn TKeyEvent) {
    DateTimePicker_SetOnKeyDown(d.instance, fn)
}

// SetOnKeyPress
func (d *TDateTimePicker) SetOnKeyPress(fn TKeyPressEvent) {
    DateTimePicker_SetOnKeyPress(d.instance, fn)
}

// SetOnKeyUp
func (d *TDateTimePicker) SetOnKeyUp(fn TKeyEvent) {
    DateTimePicker_SetOnKeyUp(d.instance, fn)
}

// SetOnMouseEnter
func (d *TDateTimePicker) SetOnMouseEnter(fn TNotifyEvent) {
    DateTimePicker_SetOnMouseEnter(d.instance, fn)
}

// SetOnMouseLeave
func (d *TDateTimePicker) SetOnMouseLeave(fn TNotifyEvent) {
    DateTimePicker_SetOnMouseLeave(d.instance, fn)
}

// SetOnStartDock
func (d *TDateTimePicker) SetOnStartDock(fn TStartDockEvent) {
    DateTimePicker_SetOnStartDock(d.instance, fn)
}

// DockSite
func (d *TDateTimePicker) DockSite() bool {
    return DateTimePicker_GetDockSite(d.instance)
}

// SetDockSite
func (d *TDateTimePicker) SetDockSite(value bool) {
    DateTimePicker_SetDockSite(d.instance, value)
}

// Brush
func (d *TDateTimePicker) Brush() *TBrush {
    return BrushFromInst(DateTimePicker_GetBrush(d.instance))
}

// ControlCount
func (d *TDateTimePicker) ControlCount() int32 {
    return DateTimePicker_GetControlCount(d.instance)
}

// Handle
func (d *TDateTimePicker) Handle() HWND {
    return DateTimePicker_GetHandle(d.instance)
}

// ParentWindow
func (d *TDateTimePicker) ParentWindow() HWND {
    return DateTimePicker_GetParentWindow(d.instance)
}

// SetParentWindow
func (d *TDateTimePicker) SetParentWindow(value HWND) {
    DateTimePicker_SetParentWindow(d.instance, value)
}

// UseDockManager
func (d *TDateTimePicker) UseDockManager() bool {
    return DateTimePicker_GetUseDockManager(d.instance)
}

// SetUseDockManager
func (d *TDateTimePicker) SetUseDockManager(value bool) {
    DateTimePicker_SetUseDockManager(d.instance, value)
}

// Action
func (d *TDateTimePicker) Action() *TAction {
    return ActionFromInst(DateTimePicker_GetAction(d.instance))
}

// SetAction
func (d *TDateTimePicker) SetAction(value IComponent) {
    DateTimePicker_SetAction(d.instance, CheckPtr(value))
}

// BoundsRect
func (d *TDateTimePicker) BoundsRect() TRect {
    return DateTimePicker_GetBoundsRect(d.instance)
}

// SetBoundsRect
func (d *TDateTimePicker) SetBoundsRect(value TRect) {
    DateTimePicker_SetBoundsRect(d.instance, value)
}

// ClientHeight
func (d *TDateTimePicker) ClientHeight() int32 {
    return DateTimePicker_GetClientHeight(d.instance)
}

// SetClientHeight
func (d *TDateTimePicker) SetClientHeight(value int32) {
    DateTimePicker_SetClientHeight(d.instance, value)
}

// ClientRect
func (d *TDateTimePicker) ClientRect() TRect {
    return DateTimePicker_GetClientRect(d.instance)
}

// ClientWidth
func (d *TDateTimePicker) ClientWidth() int32 {
    return DateTimePicker_GetClientWidth(d.instance)
}

// SetClientWidth
func (d *TDateTimePicker) SetClientWidth(value int32) {
    DateTimePicker_SetClientWidth(d.instance, value)
}

// ExplicitLeft
func (d *TDateTimePicker) ExplicitLeft() int32 {
    return DateTimePicker_GetExplicitLeft(d.instance)
}

// ExplicitTop
func (d *TDateTimePicker) ExplicitTop() int32 {
    return DateTimePicker_GetExplicitTop(d.instance)
}

// ExplicitWidth
func (d *TDateTimePicker) ExplicitWidth() int32 {
    return DateTimePicker_GetExplicitWidth(d.instance)
}

// ExplicitHeight
func (d *TDateTimePicker) ExplicitHeight() int32 {
    return DateTimePicker_GetExplicitHeight(d.instance)
}

// Floating
func (d *TDateTimePicker) Floating() bool {
    return DateTimePicker_GetFloating(d.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (d *TDateTimePicker) Parent() *TWinControl {
    return WinControlFromInst(DateTimePicker_GetParent(d.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (d *TDateTimePicker) SetParent(value IWinControl) {
    DateTimePicker_SetParent(d.instance, CheckPtr(value))
}

// AlignWithMargins
func (d *TDateTimePicker) AlignWithMargins() bool {
    return DateTimePicker_GetAlignWithMargins(d.instance)
}

// SetAlignWithMargins
func (d *TDateTimePicker) SetAlignWithMargins(value bool) {
    DateTimePicker_SetAlignWithMargins(d.instance, value)
}

// Left
func (d *TDateTimePicker) Left() int32 {
    return DateTimePicker_GetLeft(d.instance)
}

// SetLeft
func (d *TDateTimePicker) SetLeft(value int32) {
    DateTimePicker_SetLeft(d.instance, value)
}

// Top
func (d *TDateTimePicker) Top() int32 {
    return DateTimePicker_GetTop(d.instance)
}

// SetTop
func (d *TDateTimePicker) SetTop(value int32) {
    DateTimePicker_SetTop(d.instance, value)
}

// Width
func (d *TDateTimePicker) Width() int32 {
    return DateTimePicker_GetWidth(d.instance)
}

// SetWidth
func (d *TDateTimePicker) SetWidth(value int32) {
    DateTimePicker_SetWidth(d.instance, value)
}

// Height
func (d *TDateTimePicker) Height() int32 {
    return DateTimePicker_GetHeight(d.instance)
}

// SetHeight
func (d *TDateTimePicker) SetHeight(value int32) {
    DateTimePicker_SetHeight(d.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (d *TDateTimePicker) Cursor() TCursor {
    return DateTimePicker_GetCursor(d.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (d *TDateTimePicker) SetCursor(value TCursor) {
    DateTimePicker_SetCursor(d.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (d *TDateTimePicker) Hint() string {
    return DateTimePicker_GetHint(d.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (d *TDateTimePicker) SetHint(value string) {
    DateTimePicker_SetHint(d.instance, value)
}

// Margins
func (d *TDateTimePicker) Margins() *TMargins {
    return MarginsFromInst(DateTimePicker_GetMargins(d.instance))
}

// SetMargins
func (d *TDateTimePicker) SetMargins(value *TMargins) {
    DateTimePicker_SetMargins(d.instance, CheckPtr(value))
}

// CustomHint
func (d *TDateTimePicker) CustomHint() *TCustomHint {
    return CustomHintFromInst(DateTimePicker_GetCustomHint(d.instance))
}

// SetCustomHint
func (d *TDateTimePicker) SetCustomHint(value IComponent) {
    DateTimePicker_SetCustomHint(d.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (d *TDateTimePicker) ComponentCount() int32 {
    return DateTimePicker_GetComponentCount(d.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (d *TDateTimePicker) ComponentIndex() int32 {
    return DateTimePicker_GetComponentIndex(d.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (d *TDateTimePicker) SetComponentIndex(value int32) {
    DateTimePicker_SetComponentIndex(d.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (d *TDateTimePicker) Owner() *TComponent {
    return ComponentFromInst(DateTimePicker_GetOwner(d.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (d *TDateTimePicker) Name() string {
    return DateTimePicker_GetName(d.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (d *TDateTimePicker) SetName(value string) {
    DateTimePicker_SetName(d.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (d *TDateTimePicker) Tag() int {
    return DateTimePicker_GetTag(d.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (d *TDateTimePicker) SetTag(value int) {
    DateTimePicker_SetTag(d.instance, value)
}

// Controls
func (d *TDateTimePicker) Controls(Index int32) *TControl {
    return ControlFromInst(DateTimePicker_GetControls(d.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (d *TDateTimePicker) Components(AIndex int32) *TComponent {
    return ComponentFromInst(DateTimePicker_GetComponents(d.instance, AIndex))
}

