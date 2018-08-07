
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TGauge struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewGauge
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewGauge(owner IComponent) *TGauge {
    g := new(TGauge)
    g.instance = Gauge_Create(CheckPtr(owner))
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GaugeFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func GaugeFromInst(inst uintptr) *TGauge {
    g := new(TGauge)
    g.instance = inst
    g.ptr = unsafe.Pointer(inst)
    return g
}

// GaugeFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func GaugeFromObj(obj IObject) *TGauge {
    g := new(TGauge)
    g.instance = CheckPtr(obj)
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GaugeFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func GaugeFromUnsafePointer(ptr unsafe.Pointer) *TGauge {
    g := new(TGauge)
    g.instance = uintptr(ptr)
    g.ptr = ptr
    return g
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (g *TGauge) Free() {
    if g.instance != 0 {
        Gauge_Free(g.instance)
        g.instance = 0
        g.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (g *TGauge) Instance() uintptr {
    return g.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (g *TGauge) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (g *TGauge) IsValid() bool {
    return g.instance != 0
}

// TGaugeClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TGaugeClass() TClass {
    return Gauge_StaticClassType()
}

// AddProgress
func (g *TGauge) AddProgress(Value int32) {
    Gauge_AddProgress(g.instance, Value)
}

// BringToFront
func (g *TGauge) BringToFront() {
    Gauge_BringToFront(g.instance)
}

// ClientToScreen
func (g *TGauge) ClientToScreen(Point TPoint) TPoint {
    return Gauge_ClientToScreen(g.instance, Point)
}

// ClientToParent
func (g *TGauge) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Gauge_ClientToParent(g.instance, Point , CheckPtr(AParent))
}

// Dragging
func (g *TGauge) Dragging() bool {
    return Gauge_Dragging(g.instance)
}

// HasParent
func (g *TGauge) HasParent() bool {
    return Gauge_HasParent(g.instance)
}

// Hide
func (g *TGauge) Hide() {
    Gauge_Hide(g.instance)
}

// Invalidate
func (g *TGauge) Invalidate() {
    Gauge_Invalidate(g.instance)
}

// Perform
func (g *TGauge) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Gauge_Perform(g.instance, Msg , WParam , LParam)
}

// Refresh
func (g *TGauge) Refresh() {
    Gauge_Refresh(g.instance)
}

// Repaint
func (g *TGauge) Repaint() {
    Gauge_Repaint(g.instance)
}

// ScreenToClient
func (g *TGauge) ScreenToClient(Point TPoint) TPoint {
    return Gauge_ScreenToClient(g.instance, Point)
}

// ParentToClient
func (g *TGauge) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Gauge_ParentToClient(g.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (g *TGauge) SendToBack() {
    Gauge_SendToBack(g.instance)
}

// SetBounds
func (g *TGauge) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Gauge_SetBounds(g.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
func (g *TGauge) Show() {
    Gauge_Show(g.instance)
}

// Update
func (g *TGauge) Update() {
    Gauge_Update(g.instance)
}

// GetTextBuf
func (g *TGauge) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Gauge_GetTextBuf(g.instance, Buffer , BufSize)
}

// GetTextLen
func (g *TGauge) GetTextLen() int32 {
    return Gauge_GetTextLen(g.instance)
}

// FindComponent
func (g *TGauge) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Gauge_FindComponent(g.instance, AName))
}

// GetNamePath
func (g *TGauge) GetNamePath() string {
    return Gauge_GetNamePath(g.instance)
}

// Assign
func (g *TGauge) Assign(Source IObject) {
    Gauge_Assign(g.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (g *TGauge) DisposeOf() {
    Gauge_DisposeOf(g.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (g *TGauge) ClassType() TClass {
    return Gauge_ClassType(g.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (g *TGauge) ClassName() string {
    return Gauge_ClassName(g.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (g *TGauge) InstanceSize() int32 {
    return Gauge_InstanceSize(g.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (g *TGauge) InheritsFrom(AClass TClass) bool {
    return Gauge_InheritsFrom(g.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (g *TGauge) Equals(Obj IObject) bool {
    return Gauge_Equals(g.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (g *TGauge) GetHashCode() int32 {
    return Gauge_GetHashCode(g.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (g *TGauge) ToString() string {
    return Gauge_ToString(g.instance)
}

// PercentDone
func (g *TGauge) PercentDone() int32 {
    return Gauge_GetPercentDone(g.instance)
}

// Align
func (g *TGauge) Align() TAlign {
    return Gauge_GetAlign(g.instance)
}

// SetAlign
func (g *TGauge) SetAlign(value TAlign) {
    Gauge_SetAlign(g.instance, value)
}

// Anchors
func (g *TGauge) Anchors() TAnchors {
    return Gauge_GetAnchors(g.instance)
}

// SetAnchors
func (g *TGauge) SetAnchors(value TAnchors) {
    Gauge_SetAnchors(g.instance, value)
}

// BackColor
func (g *TGauge) BackColor() TColor {
    return Gauge_GetBackColor(g.instance)
}

// SetBackColor
func (g *TGauge) SetBackColor(value TColor) {
    Gauge_SetBackColor(g.instance, value)
}

// BorderStyle
func (g *TGauge) BorderStyle() TBorderStyle {
    return Gauge_GetBorderStyle(g.instance)
}

// SetBorderStyle
func (g *TGauge) SetBorderStyle(value TBorderStyle) {
    Gauge_SetBorderStyle(g.instance, value)
}

// Color
func (g *TGauge) Color() TColor {
    return Gauge_GetColor(g.instance)
}

// SetColor
func (g *TGauge) SetColor(value TColor) {
    Gauge_SetColor(g.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (g *TGauge) Enabled() bool {
    return Gauge_GetEnabled(g.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (g *TGauge) SetEnabled(value bool) {
    Gauge_SetEnabled(g.instance, value)
}

// ForeColor
func (g *TGauge) ForeColor() TColor {
    return Gauge_GetForeColor(g.instance)
}

// SetForeColor
func (g *TGauge) SetForeColor(value TColor) {
    Gauge_SetForeColor(g.instance, value)
}

// Font
func (g *TGauge) Font() *TFont {
    return FontFromInst(Gauge_GetFont(g.instance))
}

// SetFont
func (g *TGauge) SetFont(value *TFont) {
    Gauge_SetFont(g.instance, CheckPtr(value))
}

// Kind
func (g *TGauge) Kind() TGaugeKind {
    return Gauge_GetKind(g.instance)
}

// SetKind
func (g *TGauge) SetKind(value TGaugeKind) {
    Gauge_SetKind(g.instance, value)
}

// MinValue
func (g *TGauge) MinValue() int32 {
    return Gauge_GetMinValue(g.instance)
}

// SetMinValue
func (g *TGauge) SetMinValue(value int32) {
    Gauge_SetMinValue(g.instance, value)
}

// MaxValue
func (g *TGauge) MaxValue() int32 {
    return Gauge_GetMaxValue(g.instance)
}

// SetMaxValue
func (g *TGauge) SetMaxValue(value int32) {
    Gauge_SetMaxValue(g.instance, value)
}

// ParentColor
func (g *TGauge) ParentColor() bool {
    return Gauge_GetParentColor(g.instance)
}

// SetParentColor
func (g *TGauge) SetParentColor(value bool) {
    Gauge_SetParentColor(g.instance, value)
}

// ParentFont
func (g *TGauge) ParentFont() bool {
    return Gauge_GetParentFont(g.instance)
}

// SetParentFont
func (g *TGauge) SetParentFont(value bool) {
    Gauge_SetParentFont(g.instance, value)
}

// ParentShowHint
func (g *TGauge) ParentShowHint() bool {
    return Gauge_GetParentShowHint(g.instance)
}

// SetParentShowHint
func (g *TGauge) SetParentShowHint(value bool) {
    Gauge_SetParentShowHint(g.instance, value)
}

// PopupMenu
func (g *TGauge) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Gauge_GetPopupMenu(g.instance))
}

// SetPopupMenu
func (g *TGauge) SetPopupMenu(value IComponent) {
    Gauge_SetPopupMenu(g.instance, CheckPtr(value))
}

// Progress
func (g *TGauge) Progress() int32 {
    return Gauge_GetProgress(g.instance)
}

// SetProgress
func (g *TGauge) SetProgress(value int32) {
    Gauge_SetProgress(g.instance, value)
}

// ShowHint
func (g *TGauge) ShowHint() bool {
    return Gauge_GetShowHint(g.instance)
}

// SetShowHint
func (g *TGauge) SetShowHint(value bool) {
    Gauge_SetShowHint(g.instance, value)
}

// ShowText
func (g *TGauge) ShowText() bool {
    return Gauge_GetShowText(g.instance)
}

// SetShowText
func (g *TGauge) SetShowText(value bool) {
    Gauge_SetShowText(g.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (g *TGauge) Visible() bool {
    return Gauge_GetVisible(g.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (g *TGauge) SetVisible(value bool) {
    Gauge_SetVisible(g.instance, value)
}

// Action
func (g *TGauge) Action() *TAction {
    return ActionFromInst(Gauge_GetAction(g.instance))
}

// SetAction
func (g *TGauge) SetAction(value IComponent) {
    Gauge_SetAction(g.instance, CheckPtr(value))
}

// BiDiMode
func (g *TGauge) BiDiMode() TBiDiMode {
    return Gauge_GetBiDiMode(g.instance)
}

// SetBiDiMode
func (g *TGauge) SetBiDiMode(value TBiDiMode) {
    Gauge_SetBiDiMode(g.instance, value)
}

// BoundsRect
func (g *TGauge) BoundsRect() TRect {
    return Gauge_GetBoundsRect(g.instance)
}

// SetBoundsRect
func (g *TGauge) SetBoundsRect(value TRect) {
    Gauge_SetBoundsRect(g.instance, value)
}

// ClientHeight
func (g *TGauge) ClientHeight() int32 {
    return Gauge_GetClientHeight(g.instance)
}

// SetClientHeight
func (g *TGauge) SetClientHeight(value int32) {
    Gauge_SetClientHeight(g.instance, value)
}

// ClientRect
func (g *TGauge) ClientRect() TRect {
    return Gauge_GetClientRect(g.instance)
}

// ClientWidth
func (g *TGauge) ClientWidth() int32 {
    return Gauge_GetClientWidth(g.instance)
}

// SetClientWidth
func (g *TGauge) SetClientWidth(value int32) {
    Gauge_SetClientWidth(g.instance, value)
}

// ExplicitLeft
func (g *TGauge) ExplicitLeft() int32 {
    return Gauge_GetExplicitLeft(g.instance)
}

// ExplicitTop
func (g *TGauge) ExplicitTop() int32 {
    return Gauge_GetExplicitTop(g.instance)
}

// ExplicitWidth
func (g *TGauge) ExplicitWidth() int32 {
    return Gauge_GetExplicitWidth(g.instance)
}

// ExplicitHeight
func (g *TGauge) ExplicitHeight() int32 {
    return Gauge_GetExplicitHeight(g.instance)
}

// Floating
func (g *TGauge) Floating() bool {
    return Gauge_GetFloating(g.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (g *TGauge) Parent() *TWinControl {
    return WinControlFromInst(Gauge_GetParent(g.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (g *TGauge) SetParent(value IWinControl) {
    Gauge_SetParent(g.instance, CheckPtr(value))
}

// StyleElements
func (g *TGauge) StyleElements() TStyleElements {
    return Gauge_GetStyleElements(g.instance)
}

// SetStyleElements
func (g *TGauge) SetStyleElements(value TStyleElements) {
    Gauge_SetStyleElements(g.instance, value)
}

// AlignWithMargins
func (g *TGauge) AlignWithMargins() bool {
    return Gauge_GetAlignWithMargins(g.instance)
}

// SetAlignWithMargins
func (g *TGauge) SetAlignWithMargins(value bool) {
    Gauge_SetAlignWithMargins(g.instance, value)
}

// Left
func (g *TGauge) Left() int32 {
    return Gauge_GetLeft(g.instance)
}

// SetLeft
func (g *TGauge) SetLeft(value int32) {
    Gauge_SetLeft(g.instance, value)
}

// Top
func (g *TGauge) Top() int32 {
    return Gauge_GetTop(g.instance)
}

// SetTop
func (g *TGauge) SetTop(value int32) {
    Gauge_SetTop(g.instance, value)
}

// Width
func (g *TGauge) Width() int32 {
    return Gauge_GetWidth(g.instance)
}

// SetWidth
func (g *TGauge) SetWidth(value int32) {
    Gauge_SetWidth(g.instance, value)
}

// Height
func (g *TGauge) Height() int32 {
    return Gauge_GetHeight(g.instance)
}

// SetHeight
func (g *TGauge) SetHeight(value int32) {
    Gauge_SetHeight(g.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (g *TGauge) Cursor() TCursor {
    return Gauge_GetCursor(g.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (g *TGauge) SetCursor(value TCursor) {
    Gauge_SetCursor(g.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (g *TGauge) Hint() string {
    return Gauge_GetHint(g.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (g *TGauge) SetHint(value string) {
    Gauge_SetHint(g.instance, value)
}

// Margins
func (g *TGauge) Margins() *TMargins {
    return MarginsFromInst(Gauge_GetMargins(g.instance))
}

// SetMargins
func (g *TGauge) SetMargins(value *TMargins) {
    Gauge_SetMargins(g.instance, CheckPtr(value))
}

// CustomHint
func (g *TGauge) CustomHint() *TCustomHint {
    return CustomHintFromInst(Gauge_GetCustomHint(g.instance))
}

// SetCustomHint
func (g *TGauge) SetCustomHint(value IComponent) {
    Gauge_SetCustomHint(g.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (g *TGauge) ComponentCount() int32 {
    return Gauge_GetComponentCount(g.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (g *TGauge) ComponentIndex() int32 {
    return Gauge_GetComponentIndex(g.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (g *TGauge) SetComponentIndex(value int32) {
    Gauge_SetComponentIndex(g.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (g *TGauge) Owner() *TComponent {
    return ComponentFromInst(Gauge_GetOwner(g.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (g *TGauge) Name() string {
    return Gauge_GetName(g.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (g *TGauge) SetName(value string) {
    Gauge_SetName(g.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (g *TGauge) Tag() int {
    return Gauge_GetTag(g.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (g *TGauge) SetTag(value int) {
    Gauge_SetTag(g.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (g *TGauge) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Gauge_GetComponents(g.instance, AIndex))
}

