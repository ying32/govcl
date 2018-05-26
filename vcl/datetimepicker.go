
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

type TDateTimePicker struct {
    IControl
    instance uintptr
}

func NewDateTimePicker(owner IComponent) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = DateTimePicker_Create(CheckPtr(owner))
    return d
}

func DateTimePickerFromInst(inst uintptr) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = inst
    return d
}

func DateTimePickerFromObj(obj IObject) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = CheckPtr(obj)
    return d
}

func (d *TDateTimePicker) Free() {
    if d.instance != 0 {
        DateTimePicker_Free(d.instance)
        d.instance = 0
    }
}

func (d *TDateTimePicker) Instance() uintptr {
    return d.instance
}

func (d *TDateTimePicker) IsValid() bool {
    return d.instance != 0
}

func (d *TDateTimePicker) CanFocus() bool {
    return DateTimePicker_CanFocus(d.instance)
}

func (d *TDateTimePicker) FlipChildren(AllLevels bool) {
    DateTimePicker_FlipChildren(d.instance, AllLevels)
}

func (d *TDateTimePicker) Focused() bool {
    return DateTimePicker_Focused(d.instance)
}

func (d *TDateTimePicker) HandleAllocated() bool {
    return DateTimePicker_HandleAllocated(d.instance)
}

func (d *TDateTimePicker) Invalidate() {
    DateTimePicker_Invalidate(d.instance)
}

func (d *TDateTimePicker) Realign() {
    DateTimePicker_Realign(d.instance)
}

func (d *TDateTimePicker) Repaint() {
    DateTimePicker_Repaint(d.instance)
}

func (d *TDateTimePicker) ScaleBy(M int32, D int32) {
    DateTimePicker_ScaleBy(d.instance, M , D)
}

func (d *TDateTimePicker) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    DateTimePicker_SetBounds(d.instance, ALeft , ATop , AWidth , AHeight)
}

func (d *TDateTimePicker) SetFocus() {
    DateTimePicker_SetFocus(d.instance)
}

func (d *TDateTimePicker) Update() {
    DateTimePicker_Update(d.instance)
}

func (d *TDateTimePicker) BringToFront() {
    DateTimePicker_BringToFront(d.instance)
}

func (d *TDateTimePicker) HasParent() bool {
    return DateTimePicker_HasParent(d.instance)
}

func (d *TDateTimePicker) Hide() {
    DateTimePicker_Hide(d.instance)
}

func (d *TDateTimePicker) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return DateTimePicker_Perform(d.instance, Msg , WParam , LParam)
}

func (d *TDateTimePicker) Refresh() {
    DateTimePicker_Refresh(d.instance)
}

func (d *TDateTimePicker) SendToBack() {
    DateTimePicker_SendToBack(d.instance)
}

func (d *TDateTimePicker) Show() {
    DateTimePicker_Show(d.instance)
}

func (d *TDateTimePicker) GetTextBuf(Buffer string, BufSize int32) int32 {
    return DateTimePicker_GetTextBuf(d.instance, Buffer , BufSize)
}

func (d *TDateTimePicker) GetTextLen() int32 {
    return DateTimePicker_GetTextLen(d.instance)
}

func (d *TDateTimePicker) FindComponent(AName string) *TComponent {
    return ComponentFromInst(DateTimePicker_FindComponent(d.instance, AName))
}

func (d *TDateTimePicker) GetNamePath() string {
    return DateTimePicker_GetNamePath(d.instance)
}

func (d *TDateTimePicker) Assign(Source IObject) {
    DateTimePicker_Assign(d.instance, CheckPtr(Source))
}

func (d *TDateTimePicker) ClassName() string {
    return DateTimePicker_ClassName(d.instance)
}

func (d *TDateTimePicker) Equals(Obj IObject) bool {
    return DateTimePicker_Equals(d.instance, CheckPtr(Obj))
}

func (d *TDateTimePicker) GetHashCode() int32 {
    return DateTimePicker_GetHashCode(d.instance)
}

func (d *TDateTimePicker) ToString() string {
    return DateTimePicker_ToString(d.instance)
}

func (d *TDateTimePicker) DateTime() time.Time {
    return DateTimePicker_GetDateTime(d.instance)
}

func (d *TDateTimePicker) SetDateTime(value time.Time) {
    DateTimePicker_SetDateTime(d.instance, value)
}

func (d *TDateTimePicker) DroppedDown() bool {
    return DateTimePicker_GetDroppedDown(d.instance)
}

func (d *TDateTimePicker) Align() TAlign {
    return DateTimePicker_GetAlign(d.instance)
}

func (d *TDateTimePicker) SetAlign(value TAlign) {
    DateTimePicker_SetAlign(d.instance, value)
}

func (d *TDateTimePicker) Anchors() TAnchors {
    return DateTimePicker_GetAnchors(d.instance)
}

func (d *TDateTimePicker) SetAnchors(value TAnchors) {
    DateTimePicker_SetAnchors(d.instance, value)
}

func (d *TDateTimePicker) BevelEdges() TBevelEdges {
    return DateTimePicker_GetBevelEdges(d.instance)
}

func (d *TDateTimePicker) SetBevelEdges(value TBevelEdges) {
    DateTimePicker_SetBevelEdges(d.instance, value)
}

func (d *TDateTimePicker) BevelInner() TBevelCut {
    return DateTimePicker_GetBevelInner(d.instance)
}

func (d *TDateTimePicker) SetBevelInner(value TBevelCut) {
    DateTimePicker_SetBevelInner(d.instance, value)
}

func (d *TDateTimePicker) BevelOuter() TBevelCut {
    return DateTimePicker_GetBevelOuter(d.instance)
}

func (d *TDateTimePicker) SetBevelOuter(value TBevelCut) {
    DateTimePicker_SetBevelOuter(d.instance, value)
}

func (d *TDateTimePicker) BevelKind() TBevelKind {
    return DateTimePicker_GetBevelKind(d.instance)
}

func (d *TDateTimePicker) SetBevelKind(value TBevelKind) {
    DateTimePicker_SetBevelKind(d.instance, value)
}

func (d *TDateTimePicker) BiDiMode() TBiDiMode {
    return DateTimePicker_GetBiDiMode(d.instance)
}

func (d *TDateTimePicker) SetBiDiMode(value TBiDiMode) {
    DateTimePicker_SetBiDiMode(d.instance, value)
}

func (d *TDateTimePicker) CalAlignment() TDTCalAlignment {
    return DateTimePicker_GetCalAlignment(d.instance)
}

func (d *TDateTimePicker) SetCalAlignment(value TDTCalAlignment) {
    DateTimePicker_SetCalAlignment(d.instance, value)
}

func (d *TDateTimePicker) CalColors() *TMonthCalColors {
    return MonthCalColorsFromInst(DateTimePicker_GetCalColors(d.instance))
}

func (d *TDateTimePicker) SetCalColors(value *TMonthCalColors) {
    DateTimePicker_SetCalColors(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) Date() time.Time {
    return DateTimePicker_GetDate(d.instance)
}

func (d *TDateTimePicker) SetDate(value time.Time) {
    DateTimePicker_SetDate(d.instance, value)
}

func (d *TDateTimePicker) Format() string {
    return DateTimePicker_GetFormat(d.instance)
}

func (d *TDateTimePicker) SetFormat(value string) {
    DateTimePicker_SetFormat(d.instance, value)
}

func (d *TDateTimePicker) Time() time.Time {
    return DateTimePicker_GetTime(d.instance)
}

func (d *TDateTimePicker) SetTime(value time.Time) {
    DateTimePicker_SetTime(d.instance, value)
}

func (d *TDateTimePicker) Checked() bool {
    return DateTimePicker_GetChecked(d.instance)
}

func (d *TDateTimePicker) SetChecked(value bool) {
    DateTimePicker_SetChecked(d.instance, value)
}

func (d *TDateTimePicker) Color() TColor {
    return DateTimePicker_GetColor(d.instance)
}

func (d *TDateTimePicker) SetColor(value TColor) {
    DateTimePicker_SetColor(d.instance, value)
}

func (d *TDateTimePicker) DateFormat() TDTDateFormat {
    return DateTimePicker_GetDateFormat(d.instance)
}

func (d *TDateTimePicker) SetDateFormat(value TDTDateFormat) {
    DateTimePicker_SetDateFormat(d.instance, value)
}

func (d *TDateTimePicker) DateMode() TDTDateMode {
    return DateTimePicker_GetDateMode(d.instance)
}

func (d *TDateTimePicker) SetDateMode(value TDTDateMode) {
    DateTimePicker_SetDateMode(d.instance, value)
}

func (d *TDateTimePicker) DoubleBuffered() bool {
    return DateTimePicker_GetDoubleBuffered(d.instance)
}

func (d *TDateTimePicker) SetDoubleBuffered(value bool) {
    DateTimePicker_SetDoubleBuffered(d.instance, value)
}

func (d *TDateTimePicker) Enabled() bool {
    return DateTimePicker_GetEnabled(d.instance)
}

func (d *TDateTimePicker) SetEnabled(value bool) {
    DateTimePicker_SetEnabled(d.instance, value)
}

func (d *TDateTimePicker) Font() *TFont {
    return FontFromInst(DateTimePicker_GetFont(d.instance))
}

func (d *TDateTimePicker) SetFont(value *TFont) {
    DateTimePicker_SetFont(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) Kind() TDateTimeKind {
    return DateTimePicker_GetKind(d.instance)
}

func (d *TDateTimePicker) SetKind(value TDateTimeKind) {
    DateTimePicker_SetKind(d.instance, value)
}

func (d *TDateTimePicker) MaxDate() time.Time {
    return DateTimePicker_GetMaxDate(d.instance)
}

func (d *TDateTimePicker) SetMaxDate(value time.Time) {
    DateTimePicker_SetMaxDate(d.instance, value)
}

func (d *TDateTimePicker) MinDate() time.Time {
    return DateTimePicker_GetMinDate(d.instance)
}

func (d *TDateTimePicker) SetMinDate(value time.Time) {
    DateTimePicker_SetMinDate(d.instance, value)
}

func (d *TDateTimePicker) ParseInput() bool {
    return DateTimePicker_GetParseInput(d.instance)
}

func (d *TDateTimePicker) SetParseInput(value bool) {
    DateTimePicker_SetParseInput(d.instance, value)
}

func (d *TDateTimePicker) ParentColor() bool {
    return DateTimePicker_GetParentColor(d.instance)
}

func (d *TDateTimePicker) SetParentColor(value bool) {
    DateTimePicker_SetParentColor(d.instance, value)
}

func (d *TDateTimePicker) ParentDoubleBuffered() bool {
    return DateTimePicker_GetParentDoubleBuffered(d.instance)
}

func (d *TDateTimePicker) SetParentDoubleBuffered(value bool) {
    DateTimePicker_SetParentDoubleBuffered(d.instance, value)
}

func (d *TDateTimePicker) ParentFont() bool {
    return DateTimePicker_GetParentFont(d.instance)
}

func (d *TDateTimePicker) SetParentFont(value bool) {
    DateTimePicker_SetParentFont(d.instance, value)
}

func (d *TDateTimePicker) ParentShowHint() bool {
    return DateTimePicker_GetParentShowHint(d.instance)
}

func (d *TDateTimePicker) SetParentShowHint(value bool) {
    DateTimePicker_SetParentShowHint(d.instance, value)
}

func (d *TDateTimePicker) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(DateTimePicker_GetPopupMenu(d.instance))
}

func (d *TDateTimePicker) SetPopupMenu(value IComponent) {
    DateTimePicker_SetPopupMenu(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) ShowHint() bool {
    return DateTimePicker_GetShowHint(d.instance)
}

func (d *TDateTimePicker) SetShowHint(value bool) {
    DateTimePicker_SetShowHint(d.instance, value)
}

func (d *TDateTimePicker) TabOrder() uint16 {
    return DateTimePicker_GetTabOrder(d.instance)
}

func (d *TDateTimePicker) SetTabOrder(value uint16) {
    DateTimePicker_SetTabOrder(d.instance, value)
}

func (d *TDateTimePicker) TabStop() bool {
    return DateTimePicker_GetTabStop(d.instance)
}

func (d *TDateTimePicker) SetTabStop(value bool) {
    DateTimePicker_SetTabStop(d.instance, value)
}

func (d *TDateTimePicker) Visible() bool {
    return DateTimePicker_GetVisible(d.instance)
}

func (d *TDateTimePicker) SetVisible(value bool) {
    DateTimePicker_SetVisible(d.instance, value)
}

func (d *TDateTimePicker) StyleElements() TStyleElements {
    return DateTimePicker_GetStyleElements(d.instance)
}

func (d *TDateTimePicker) SetStyleElements(value TStyleElements) {
    DateTimePicker_SetStyleElements(d.instance, value)
}

func (d *TDateTimePicker) SetOnClick(fn TNotifyEvent) {
    DateTimePicker_SetOnClick(d.instance, fn)
}

func (d *TDateTimePicker) SetOnChange(fn TNotifyEvent) {
    DateTimePicker_SetOnChange(d.instance, fn)
}

func (d *TDateTimePicker) SetOnEnter(fn TNotifyEvent) {
    DateTimePicker_SetOnEnter(d.instance, fn)
}

func (d *TDateTimePicker) SetOnExit(fn TNotifyEvent) {
    DateTimePicker_SetOnExit(d.instance, fn)
}

func (d *TDateTimePicker) SetOnKeyDown(fn TKeyEvent) {
    DateTimePicker_SetOnKeyDown(d.instance, fn)
}

func (d *TDateTimePicker) SetOnKeyPress(fn TKeyPressEvent) {
    DateTimePicker_SetOnKeyPress(d.instance, fn)
}

func (d *TDateTimePicker) SetOnKeyUp(fn TKeyEvent) {
    DateTimePicker_SetOnKeyUp(d.instance, fn)
}

func (d *TDateTimePicker) SetOnMouseEnter(fn TNotifyEvent) {
    DateTimePicker_SetOnMouseEnter(d.instance, fn)
}

func (d *TDateTimePicker) SetOnMouseLeave(fn TNotifyEvent) {
    DateTimePicker_SetOnMouseLeave(d.instance, fn)
}

func (d *TDateTimePicker) Brush() *TBrush {
    return BrushFromInst(DateTimePicker_GetBrush(d.instance))
}

func (d *TDateTimePicker) ControlCount() int32 {
    return DateTimePicker_GetControlCount(d.instance)
}

func (d *TDateTimePicker) Handle() HWND {
    return DateTimePicker_GetHandle(d.instance)
}

func (d *TDateTimePicker) ParentWindow() HWND {
    return DateTimePicker_GetParentWindow(d.instance)
}

func (d *TDateTimePicker) SetParentWindow(value HWND) {
    DateTimePicker_SetParentWindow(d.instance, value)
}

func (d *TDateTimePicker) Action() *TAction {
    return ActionFromInst(DateTimePicker_GetAction(d.instance))
}

func (d *TDateTimePicker) SetAction(value IComponent) {
    DateTimePicker_SetAction(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) BoundsRect() TRect {
    return DateTimePicker_GetBoundsRect(d.instance)
}

func (d *TDateTimePicker) SetBoundsRect(value TRect) {
    DateTimePicker_SetBoundsRect(d.instance, value)
}

func (d *TDateTimePicker) ClientHeight() int32 {
    return DateTimePicker_GetClientHeight(d.instance)
}

func (d *TDateTimePicker) SetClientHeight(value int32) {
    DateTimePicker_SetClientHeight(d.instance, value)
}

func (d *TDateTimePicker) ClientRect() TRect {
    return DateTimePicker_GetClientRect(d.instance)
}

func (d *TDateTimePicker) ClientWidth() int32 {
    return DateTimePicker_GetClientWidth(d.instance)
}

func (d *TDateTimePicker) SetClientWidth(value int32) {
    DateTimePicker_SetClientWidth(d.instance, value)
}

func (d *TDateTimePicker) ExplicitLeft() int32 {
    return DateTimePicker_GetExplicitLeft(d.instance)
}

func (d *TDateTimePicker) ExplicitTop() int32 {
    return DateTimePicker_GetExplicitTop(d.instance)
}

func (d *TDateTimePicker) ExplicitWidth() int32 {
    return DateTimePicker_GetExplicitWidth(d.instance)
}

func (d *TDateTimePicker) ExplicitHeight() int32 {
    return DateTimePicker_GetExplicitHeight(d.instance)
}

func (d *TDateTimePicker) Parent() *TControl {
    return ControlFromInst(DateTimePicker_GetParent(d.instance))
}

func (d *TDateTimePicker) SetParent(value IControl) {
    DateTimePicker_SetParent(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) AlignWithMargins() bool {
    return DateTimePicker_GetAlignWithMargins(d.instance)
}

func (d *TDateTimePicker) SetAlignWithMargins(value bool) {
    DateTimePicker_SetAlignWithMargins(d.instance, value)
}

func (d *TDateTimePicker) Left() int32 {
    return DateTimePicker_GetLeft(d.instance)
}

func (d *TDateTimePicker) SetLeft(value int32) {
    DateTimePicker_SetLeft(d.instance, value)
}

func (d *TDateTimePicker) Top() int32 {
    return DateTimePicker_GetTop(d.instance)
}

func (d *TDateTimePicker) SetTop(value int32) {
    DateTimePicker_SetTop(d.instance, value)
}

func (d *TDateTimePicker) Width() int32 {
    return DateTimePicker_GetWidth(d.instance)
}

func (d *TDateTimePicker) SetWidth(value int32) {
    DateTimePicker_SetWidth(d.instance, value)
}

func (d *TDateTimePicker) Height() int32 {
    return DateTimePicker_GetHeight(d.instance)
}

func (d *TDateTimePicker) SetHeight(value int32) {
    DateTimePicker_SetHeight(d.instance, value)
}

func (d *TDateTimePicker) Cursor() TCursor {
    return DateTimePicker_GetCursor(d.instance)
}

func (d *TDateTimePicker) SetCursor(value TCursor) {
    DateTimePicker_SetCursor(d.instance, value)
}

func (d *TDateTimePicker) Hint() string {
    return DateTimePicker_GetHint(d.instance)
}

func (d *TDateTimePicker) SetHint(value string) {
    DateTimePicker_SetHint(d.instance, value)
}

func (d *TDateTimePicker) Margins() *TMargins {
    return MarginsFromInst(DateTimePicker_GetMargins(d.instance))
}

func (d *TDateTimePicker) SetMargins(value *TMargins) {
    DateTimePicker_SetMargins(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) CustomHint() *TCustomHint {
    return CustomHintFromInst(DateTimePicker_GetCustomHint(d.instance))
}

func (d *TDateTimePicker) SetCustomHint(value IComponent) {
    DateTimePicker_SetCustomHint(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) ComponentCount() int32 {
    return DateTimePicker_GetComponentCount(d.instance)
}

func (d *TDateTimePicker) ComponentIndex() int32 {
    return DateTimePicker_GetComponentIndex(d.instance)
}

func (d *TDateTimePicker) SetComponentIndex(value int32) {
    DateTimePicker_SetComponentIndex(d.instance, value)
}

func (d *TDateTimePicker) Owner() *TComponent {
    return ComponentFromInst(DateTimePicker_GetOwner(d.instance))
}

func (d *TDateTimePicker) Name() string {
    return DateTimePicker_GetName(d.instance)
}

func (d *TDateTimePicker) SetName(value string) {
    DateTimePicker_SetName(d.instance, value)
}

func (d *TDateTimePicker) Tag() int {
    return DateTimePicker_GetTag(d.instance)
}

func (d *TDateTimePicker) SetTag(value int) {
    DateTimePicker_SetTag(d.instance, value)
}

func (d *TDateTimePicker) Controls(Index int32) *TControl {
    return ControlFromInst(DateTimePicker_GetControls(d.instance, Index))
}

func (d *TDateTimePicker) Components(AIndex int32) *TComponent {
    return ComponentFromInst(DateTimePicker_GetComponents(d.instance, AIndex))
}

