
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
    "time"
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TMonthCalendar struct {
    IControl
    instance uintptr
}

func NewMonthCalendar(owner IComponent) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = MonthCalendar_Create(CheckPtr(owner))
    return m
}

func MonthCalendarFromInst(inst uintptr) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = inst
    return m
}

func MonthCalendarFromObj(obj IObject) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMonthCalendar) Free() {
    if m.instance != 0 {
        MonthCalendar_Free(m.instance)
        m.instance = 0
    }
}

func (m *TMonthCalendar) Instance() uintptr {
    return m.instance
}

func (m *TMonthCalendar) IsValid() bool {
    return m.instance != 0
}

func (m *TMonthCalendar) CanFocus() bool {
    return MonthCalendar_CanFocus(m.instance)
}

func (m *TMonthCalendar) FlipChildren(AllLevels bool) {
    MonthCalendar_FlipChildren(m.instance, AllLevels)
}

func (m *TMonthCalendar) Focused() bool {
    return MonthCalendar_Focused(m.instance)
}

func (m *TMonthCalendar) HandleAllocated() bool {
    return MonthCalendar_HandleAllocated(m.instance)
}

func (m *TMonthCalendar) Invalidate() {
    MonthCalendar_Invalidate(m.instance)
}

func (m *TMonthCalendar) Realign() {
    MonthCalendar_Realign(m.instance)
}

func (m *TMonthCalendar) Repaint() {
    MonthCalendar_Repaint(m.instance)
}

func (m *TMonthCalendar) ScaleBy(M int32, D int32) {
    MonthCalendar_ScaleBy(m.instance, M , D)
}

func (m *TMonthCalendar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    MonthCalendar_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

func (m *TMonthCalendar) SetFocus() {
    MonthCalendar_SetFocus(m.instance)
}

func (m *TMonthCalendar) Update() {
    MonthCalendar_Update(m.instance)
}

func (m *TMonthCalendar) BringToFront() {
    MonthCalendar_BringToFront(m.instance)
}

func (m *TMonthCalendar) HasParent() bool {
    return MonthCalendar_HasParent(m.instance)
}

func (m *TMonthCalendar) Hide() {
    MonthCalendar_Hide(m.instance)
}

func (m *TMonthCalendar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return MonthCalendar_Perform(m.instance, Msg , WParam , LParam)
}

func (m *TMonthCalendar) Refresh() {
    MonthCalendar_Refresh(m.instance)
}

func (m *TMonthCalendar) SendToBack() {
    MonthCalendar_SendToBack(m.instance)
}

func (m *TMonthCalendar) Show() {
    MonthCalendar_Show(m.instance)
}

func (m *TMonthCalendar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return MonthCalendar_GetTextBuf(m.instance, Buffer , BufSize)
}

func (m *TMonthCalendar) GetTextLen() int32 {
    return MonthCalendar_GetTextLen(m.instance)
}

func (m *TMonthCalendar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(MonthCalendar_FindComponent(m.instance, AName))
}

func (m *TMonthCalendar) GetNamePath() string {
    return MonthCalendar_GetNamePath(m.instance)
}

func (m *TMonthCalendar) Assign(Source IObject) {
    MonthCalendar_Assign(m.instance, CheckPtr(Source))
}

func (m *TMonthCalendar) ClassName() string {
    return MonthCalendar_ClassName(m.instance)
}

func (m *TMonthCalendar) Equals(Obj IObject) bool {
    return MonthCalendar_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMonthCalendar) GetHashCode() int32 {
    return MonthCalendar_GetHashCode(m.instance)
}

func (m *TMonthCalendar) ToString() string {
    return MonthCalendar_ToString(m.instance)
}

func (m *TMonthCalendar) Align() TAlign {
    return MonthCalendar_GetAlign(m.instance)
}

func (m *TMonthCalendar) SetAlign(value TAlign) {
    MonthCalendar_SetAlign(m.instance, value)
}

func (m *TMonthCalendar) Anchors() TAnchors {
    return MonthCalendar_GetAnchors(m.instance)
}

func (m *TMonthCalendar) SetAnchors(value TAnchors) {
    MonthCalendar_SetAnchors(m.instance, value)
}

func (m *TMonthCalendar) AutoSize() bool {
    return MonthCalendar_GetAutoSize(m.instance)
}

func (m *TMonthCalendar) SetAutoSize(value bool) {
    MonthCalendar_SetAutoSize(m.instance, value)
}

func (m *TMonthCalendar) BorderWidth() int32 {
    return MonthCalendar_GetBorderWidth(m.instance)
}

func (m *TMonthCalendar) SetBorderWidth(value int32) {
    MonthCalendar_SetBorderWidth(m.instance, value)
}

func (m *TMonthCalendar) BiDiMode() TBiDiMode {
    return MonthCalendar_GetBiDiMode(m.instance)
}

func (m *TMonthCalendar) SetBiDiMode(value TBiDiMode) {
    MonthCalendar_SetBiDiMode(m.instance, value)
}

func (m *TMonthCalendar) CalColors() *TMonthCalColors {
    return MonthCalColorsFromInst(MonthCalendar_GetCalColors(m.instance))
}

func (m *TMonthCalendar) SetCalColors(value *TMonthCalColors) {
    MonthCalendar_SetCalColors(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) MultiSelect() bool {
    return MonthCalendar_GetMultiSelect(m.instance)
}

func (m *TMonthCalendar) SetMultiSelect(value bool) {
    MonthCalendar_SetMultiSelect(m.instance, value)
}

func (m *TMonthCalendar) Date() time.Time {
    return MonthCalendar_GetDate(m.instance)
}

func (m *TMonthCalendar) SetDate(value time.Time) {
    MonthCalendar_SetDate(m.instance, value)
}

func (m *TMonthCalendar) DoubleBuffered() bool {
    return MonthCalendar_GetDoubleBuffered(m.instance)
}

func (m *TMonthCalendar) SetDoubleBuffered(value bool) {
    MonthCalendar_SetDoubleBuffered(m.instance, value)
}

func (m *TMonthCalendar) Enabled() bool {
    return MonthCalendar_GetEnabled(m.instance)
}

func (m *TMonthCalendar) SetEnabled(value bool) {
    MonthCalendar_SetEnabled(m.instance, value)
}

func (m *TMonthCalendar) FirstDayOfWeek() TCalDayOfWeek {
    return MonthCalendar_GetFirstDayOfWeek(m.instance)
}

func (m *TMonthCalendar) SetFirstDayOfWeek(value TCalDayOfWeek) {
    MonthCalendar_SetFirstDayOfWeek(m.instance, value)
}

func (m *TMonthCalendar) Font() *TFont {
    return FontFromInst(MonthCalendar_GetFont(m.instance))
}

func (m *TMonthCalendar) SetFont(value *TFont) {
    MonthCalendar_SetFont(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) MaxDate() time.Time {
    return MonthCalendar_GetMaxDate(m.instance)
}

func (m *TMonthCalendar) SetMaxDate(value time.Time) {
    MonthCalendar_SetMaxDate(m.instance, value)
}

func (m *TMonthCalendar) MaxSelectRange() int32 {
    return MonthCalendar_GetMaxSelectRange(m.instance)
}

func (m *TMonthCalendar) SetMaxSelectRange(value int32) {
    MonthCalendar_SetMaxSelectRange(m.instance, value)
}

func (m *TMonthCalendar) MinDate() time.Time {
    return MonthCalendar_GetMinDate(m.instance)
}

func (m *TMonthCalendar) SetMinDate(value time.Time) {
    MonthCalendar_SetMinDate(m.instance, value)
}

func (m *TMonthCalendar) ParentDoubleBuffered() bool {
    return MonthCalendar_GetParentDoubleBuffered(m.instance)
}

func (m *TMonthCalendar) SetParentDoubleBuffered(value bool) {
    MonthCalendar_SetParentDoubleBuffered(m.instance, value)
}

func (m *TMonthCalendar) ParentFont() bool {
    return MonthCalendar_GetParentFont(m.instance)
}

func (m *TMonthCalendar) SetParentFont(value bool) {
    MonthCalendar_SetParentFont(m.instance, value)
}

func (m *TMonthCalendar) ParentShowHint() bool {
    return MonthCalendar_GetParentShowHint(m.instance)
}

func (m *TMonthCalendar) SetParentShowHint(value bool) {
    MonthCalendar_SetParentShowHint(m.instance, value)
}

func (m *TMonthCalendar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(MonthCalendar_GetPopupMenu(m.instance))
}

func (m *TMonthCalendar) SetPopupMenu(value IComponent) {
    MonthCalendar_SetPopupMenu(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) ShowHint() bool {
    return MonthCalendar_GetShowHint(m.instance)
}

func (m *TMonthCalendar) SetShowHint(value bool) {
    MonthCalendar_SetShowHint(m.instance, value)
}

func (m *TMonthCalendar) ShowToday() bool {
    return MonthCalendar_GetShowToday(m.instance)
}

func (m *TMonthCalendar) SetShowToday(value bool) {
    MonthCalendar_SetShowToday(m.instance, value)
}

func (m *TMonthCalendar) ShowTodayCircle() bool {
    return MonthCalendar_GetShowTodayCircle(m.instance)
}

func (m *TMonthCalendar) SetShowTodayCircle(value bool) {
    MonthCalendar_SetShowTodayCircle(m.instance, value)
}

func (m *TMonthCalendar) TabOrder() uint16 {
    return MonthCalendar_GetTabOrder(m.instance)
}

func (m *TMonthCalendar) SetTabOrder(value uint16) {
    MonthCalendar_SetTabOrder(m.instance, value)
}

func (m *TMonthCalendar) TabStop() bool {
    return MonthCalendar_GetTabStop(m.instance)
}

func (m *TMonthCalendar) SetTabStop(value bool) {
    MonthCalendar_SetTabStop(m.instance, value)
}

func (m *TMonthCalendar) Visible() bool {
    return MonthCalendar_GetVisible(m.instance)
}

func (m *TMonthCalendar) SetVisible(value bool) {
    MonthCalendar_SetVisible(m.instance, value)
}

func (m *TMonthCalendar) WeekNumbers() bool {
    return MonthCalendar_GetWeekNumbers(m.instance)
}

func (m *TMonthCalendar) SetWeekNumbers(value bool) {
    MonthCalendar_SetWeekNumbers(m.instance, value)
}

func (m *TMonthCalendar) SetOnClick(fn TNotifyEvent) {
    MonthCalendar_SetOnClick(m.instance, fn)
}

func (m *TMonthCalendar) SetOnDblClick(fn TNotifyEvent) {
    MonthCalendar_SetOnDblClick(m.instance, fn)
}

func (m *TMonthCalendar) SetOnEnter(fn TNotifyEvent) {
    MonthCalendar_SetOnEnter(m.instance, fn)
}

func (m *TMonthCalendar) SetOnExit(fn TNotifyEvent) {
    MonthCalendar_SetOnExit(m.instance, fn)
}

func (m *TMonthCalendar) SetOnKeyDown(fn TKeyEvent) {
    MonthCalendar_SetOnKeyDown(m.instance, fn)
}

func (m *TMonthCalendar) SetOnKeyPress(fn TKeyPressEvent) {
    MonthCalendar_SetOnKeyPress(m.instance, fn)
}

func (m *TMonthCalendar) SetOnKeyUp(fn TKeyEvent) {
    MonthCalendar_SetOnKeyUp(m.instance, fn)
}

func (m *TMonthCalendar) SetOnMouseEnter(fn TNotifyEvent) {
    MonthCalendar_SetOnMouseEnter(m.instance, fn)
}

func (m *TMonthCalendar) SetOnMouseLeave(fn TNotifyEvent) {
    MonthCalendar_SetOnMouseLeave(m.instance, fn)
}

func (m *TMonthCalendar) Brush() *TBrush {
    return BrushFromInst(MonthCalendar_GetBrush(m.instance))
}

func (m *TMonthCalendar) ControlCount() int32 {
    return MonthCalendar_GetControlCount(m.instance)
}

func (m *TMonthCalendar) Handle() HWND {
    return MonthCalendar_GetHandle(m.instance)
}

func (m *TMonthCalendar) ParentWindow() HWND {
    return MonthCalendar_GetParentWindow(m.instance)
}

func (m *TMonthCalendar) SetParentWindow(value HWND) {
    MonthCalendar_SetParentWindow(m.instance, value)
}

func (m *TMonthCalendar) Action() *TAction {
    return ActionFromInst(MonthCalendar_GetAction(m.instance))
}

func (m *TMonthCalendar) SetAction(value IComponent) {
    MonthCalendar_SetAction(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) BoundsRect() TRect {
    return MonthCalendar_GetBoundsRect(m.instance)
}

func (m *TMonthCalendar) SetBoundsRect(value TRect) {
    MonthCalendar_SetBoundsRect(m.instance, value)
}

func (m *TMonthCalendar) ClientHeight() int32 {
    return MonthCalendar_GetClientHeight(m.instance)
}

func (m *TMonthCalendar) SetClientHeight(value int32) {
    MonthCalendar_SetClientHeight(m.instance, value)
}

func (m *TMonthCalendar) ClientRect() TRect {
    return MonthCalendar_GetClientRect(m.instance)
}

func (m *TMonthCalendar) ClientWidth() int32 {
    return MonthCalendar_GetClientWidth(m.instance)
}

func (m *TMonthCalendar) SetClientWidth(value int32) {
    MonthCalendar_SetClientWidth(m.instance, value)
}

func (m *TMonthCalendar) ExplicitLeft() int32 {
    return MonthCalendar_GetExplicitLeft(m.instance)
}

func (m *TMonthCalendar) ExplicitTop() int32 {
    return MonthCalendar_GetExplicitTop(m.instance)
}

func (m *TMonthCalendar) ExplicitWidth() int32 {
    return MonthCalendar_GetExplicitWidth(m.instance)
}

func (m *TMonthCalendar) ExplicitHeight() int32 {
    return MonthCalendar_GetExplicitHeight(m.instance)
}

func (m *TMonthCalendar) Parent() *TControl {
    return ControlFromInst(MonthCalendar_GetParent(m.instance))
}

func (m *TMonthCalendar) SetParent(value IControl) {
    MonthCalendar_SetParent(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) StyleElements() TStyleElements {
    return MonthCalendar_GetStyleElements(m.instance)
}

func (m *TMonthCalendar) SetStyleElements(value TStyleElements) {
    MonthCalendar_SetStyleElements(m.instance, value)
}

func (m *TMonthCalendar) AlignWithMargins() bool {
    return MonthCalendar_GetAlignWithMargins(m.instance)
}

func (m *TMonthCalendar) SetAlignWithMargins(value bool) {
    MonthCalendar_SetAlignWithMargins(m.instance, value)
}

func (m *TMonthCalendar) Left() int32 {
    return MonthCalendar_GetLeft(m.instance)
}

func (m *TMonthCalendar) SetLeft(value int32) {
    MonthCalendar_SetLeft(m.instance, value)
}

func (m *TMonthCalendar) Top() int32 {
    return MonthCalendar_GetTop(m.instance)
}

func (m *TMonthCalendar) SetTop(value int32) {
    MonthCalendar_SetTop(m.instance, value)
}

func (m *TMonthCalendar) Width() int32 {
    return MonthCalendar_GetWidth(m.instance)
}

func (m *TMonthCalendar) SetWidth(value int32) {
    MonthCalendar_SetWidth(m.instance, value)
}

func (m *TMonthCalendar) Height() int32 {
    return MonthCalendar_GetHeight(m.instance)
}

func (m *TMonthCalendar) SetHeight(value int32) {
    MonthCalendar_SetHeight(m.instance, value)
}

func (m *TMonthCalendar) Cursor() TCursor {
    return MonthCalendar_GetCursor(m.instance)
}

func (m *TMonthCalendar) SetCursor(value TCursor) {
    MonthCalendar_SetCursor(m.instance, value)
}

func (m *TMonthCalendar) Hint() string {
    return MonthCalendar_GetHint(m.instance)
}

func (m *TMonthCalendar) SetHint(value string) {
    MonthCalendar_SetHint(m.instance, value)
}

func (m *TMonthCalendar) Margins() *TMargins {
    return MarginsFromInst(MonthCalendar_GetMargins(m.instance))
}

func (m *TMonthCalendar) SetMargins(value *TMargins) {
    MonthCalendar_SetMargins(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) CustomHint() *TCustomHint {
    return CustomHintFromInst(MonthCalendar_GetCustomHint(m.instance))
}

func (m *TMonthCalendar) SetCustomHint(value IComponent) {
    MonthCalendar_SetCustomHint(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) ComponentCount() int32 {
    return MonthCalendar_GetComponentCount(m.instance)
}

func (m *TMonthCalendar) ComponentIndex() int32 {
    return MonthCalendar_GetComponentIndex(m.instance)
}

func (m *TMonthCalendar) SetComponentIndex(value int32) {
    MonthCalendar_SetComponentIndex(m.instance, value)
}

func (m *TMonthCalendar) Owner() *TComponent {
    return ComponentFromInst(MonthCalendar_GetOwner(m.instance))
}

func (m *TMonthCalendar) Name() string {
    return MonthCalendar_GetName(m.instance)
}

func (m *TMonthCalendar) SetName(value string) {
    MonthCalendar_SetName(m.instance, value)
}

func (m *TMonthCalendar) Tag() int {
    return MonthCalendar_GetTag(m.instance)
}

func (m *TMonthCalendar) SetTag(value int) {
    MonthCalendar_SetTag(m.instance, value)
}

func (m *TMonthCalendar) Controls(Index int32) *TControl {
    return ControlFromInst(MonthCalendar_GetControls(m.instance, Index))
}

func (m *TMonthCalendar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(MonthCalendar_GetComponents(m.instance, AIndex))
}

