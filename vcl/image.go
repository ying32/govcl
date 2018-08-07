
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

type TImage struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewImage
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewImage(owner IComponent) *TImage {
    i := new(TImage)
    i.instance = Image_Create(CheckPtr(owner))
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// ImageFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ImageFromInst(inst uintptr) *TImage {
    i := new(TImage)
    i.instance = inst
    i.ptr = unsafe.Pointer(inst)
    return i
}

// ImageFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ImageFromObj(obj IObject) *TImage {
    i := new(TImage)
    i.instance = CheckPtr(obj)
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// ImageFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ImageFromUnsafePointer(ptr unsafe.Pointer) *TImage {
    i := new(TImage)
    i.instance = uintptr(ptr)
    i.ptr = ptr
    return i
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (i *TImage) Free() {
    if i.instance != 0 {
        Image_Free(i.instance)
        i.instance = 0
        i.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (i *TImage) Instance() uintptr {
    return i.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (i *TImage) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (i *TImage) IsValid() bool {
    return i.instance != 0
}

// TImageClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TImageClass() TClass {
    return Image_StaticClassType()
}

// BringToFront
func (i *TImage) BringToFront() {
    Image_BringToFront(i.instance)
}

// ClientToScreen
func (i *TImage) ClientToScreen(Point TPoint) TPoint {
    return Image_ClientToScreen(i.instance, Point)
}

// ClientToParent
func (i *TImage) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Image_ClientToParent(i.instance, Point , CheckPtr(AParent))
}

// Dragging
func (i *TImage) Dragging() bool {
    return Image_Dragging(i.instance)
}

// HasParent
func (i *TImage) HasParent() bool {
    return Image_HasParent(i.instance)
}

// Hide
func (i *TImage) Hide() {
    Image_Hide(i.instance)
}

// Invalidate
func (i *TImage) Invalidate() {
    Image_Invalidate(i.instance)
}

// Perform
func (i *TImage) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Image_Perform(i.instance, Msg , WParam , LParam)
}

// Refresh
func (i *TImage) Refresh() {
    Image_Refresh(i.instance)
}

// Repaint
func (i *TImage) Repaint() {
    Image_Repaint(i.instance)
}

// ScreenToClient
func (i *TImage) ScreenToClient(Point TPoint) TPoint {
    return Image_ScreenToClient(i.instance, Point)
}

// ParentToClient
func (i *TImage) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Image_ParentToClient(i.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (i *TImage) SendToBack() {
    Image_SendToBack(i.instance)
}

// SetBounds
func (i *TImage) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Image_SetBounds(i.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
func (i *TImage) Show() {
    Image_Show(i.instance)
}

// Update
func (i *TImage) Update() {
    Image_Update(i.instance)
}

// GetTextBuf
func (i *TImage) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Image_GetTextBuf(i.instance, Buffer , BufSize)
}

// GetTextLen
func (i *TImage) GetTextLen() int32 {
    return Image_GetTextLen(i.instance)
}

// FindComponent
func (i *TImage) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Image_FindComponent(i.instance, AName))
}

// GetNamePath
func (i *TImage) GetNamePath() string {
    return Image_GetNamePath(i.instance)
}

// Assign
func (i *TImage) Assign(Source IObject) {
    Image_Assign(i.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (i *TImage) DisposeOf() {
    Image_DisposeOf(i.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (i *TImage) ClassType() TClass {
    return Image_ClassType(i.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (i *TImage) ClassName() string {
    return Image_ClassName(i.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (i *TImage) InstanceSize() int32 {
    return Image_InstanceSize(i.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (i *TImage) InheritsFrom(AClass TClass) bool {
    return Image_InheritsFrom(i.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (i *TImage) Equals(Obj IObject) bool {
    return Image_Equals(i.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (i *TImage) GetHashCode() int32 {
    return Image_GetHashCode(i.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (i *TImage) ToString() string {
    return Image_ToString(i.instance)
}

// Canvas
func (i *TImage) Canvas() *TCanvas {
    return CanvasFromInst(Image_GetCanvas(i.instance))
}

// Align
func (i *TImage) Align() TAlign {
    return Image_GetAlign(i.instance)
}

// SetAlign
func (i *TImage) SetAlign(value TAlign) {
    Image_SetAlign(i.instance, value)
}

// Anchors
func (i *TImage) Anchors() TAnchors {
    return Image_GetAnchors(i.instance)
}

// SetAnchors
func (i *TImage) SetAnchors(value TAnchors) {
    Image_SetAnchors(i.instance, value)
}

// AutoSize
func (i *TImage) AutoSize() bool {
    return Image_GetAutoSize(i.instance)
}

// SetAutoSize
func (i *TImage) SetAutoSize(value bool) {
    Image_SetAutoSize(i.instance, value)
}

// Center
func (i *TImage) Center() bool {
    return Image_GetCenter(i.instance)
}

// SetCenter
func (i *TImage) SetCenter(value bool) {
    Image_SetCenter(i.instance, value)
}

// DragCursor
func (i *TImage) DragCursor() TCursor {
    return Image_GetDragCursor(i.instance)
}

// SetDragCursor
func (i *TImage) SetDragCursor(value TCursor) {
    Image_SetDragCursor(i.instance, value)
}

// DragKind
func (i *TImage) DragKind() TDragKind {
    return Image_GetDragKind(i.instance)
}

// SetDragKind
func (i *TImage) SetDragKind(value TDragKind) {
    Image_SetDragKind(i.instance, value)
}

// DragMode
func (i *TImage) DragMode() TDragMode {
    return Image_GetDragMode(i.instance)
}

// SetDragMode
func (i *TImage) SetDragMode(value TDragMode) {
    Image_SetDragMode(i.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (i *TImage) Enabled() bool {
    return Image_GetEnabled(i.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (i *TImage) SetEnabled(value bool) {
    Image_SetEnabled(i.instance, value)
}

// IncrementalDisplay
func (i *TImage) IncrementalDisplay() bool {
    return Image_GetIncrementalDisplay(i.instance)
}

// SetIncrementalDisplay
func (i *TImage) SetIncrementalDisplay(value bool) {
    Image_SetIncrementalDisplay(i.instance, value)
}

// ParentShowHint
func (i *TImage) ParentShowHint() bool {
    return Image_GetParentShowHint(i.instance)
}

// SetParentShowHint
func (i *TImage) SetParentShowHint(value bool) {
    Image_SetParentShowHint(i.instance, value)
}

// Picture
func (i *TImage) Picture() *TPicture {
    return PictureFromInst(Image_GetPicture(i.instance))
}

// SetPicture
func (i *TImage) SetPicture(value *TPicture) {
    Image_SetPicture(i.instance, CheckPtr(value))
}

// PopupMenu
func (i *TImage) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Image_GetPopupMenu(i.instance))
}

// SetPopupMenu
func (i *TImage) SetPopupMenu(value IComponent) {
    Image_SetPopupMenu(i.instance, CheckPtr(value))
}

// Proportional
func (i *TImage) Proportional() bool {
    return Image_GetProportional(i.instance)
}

// SetProportional
func (i *TImage) SetProportional(value bool) {
    Image_SetProportional(i.instance, value)
}

// ShowHint
func (i *TImage) ShowHint() bool {
    return Image_GetShowHint(i.instance)
}

// SetShowHint
func (i *TImage) SetShowHint(value bool) {
    Image_SetShowHint(i.instance, value)
}

// Stretch
func (i *TImage) Stretch() bool {
    return Image_GetStretch(i.instance)
}

// SetStretch
func (i *TImage) SetStretch(value bool) {
    Image_SetStretch(i.instance, value)
}

// Transparent
func (i *TImage) Transparent() bool {
    return Image_GetTransparent(i.instance)
}

// SetTransparent
func (i *TImage) SetTransparent(value bool) {
    Image_SetTransparent(i.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (i *TImage) Visible() bool {
    return Image_GetVisible(i.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (i *TImage) SetVisible(value bool) {
    Image_SetVisible(i.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (i *TImage) SetOnClick(fn TNotifyEvent) {
    Image_SetOnClick(i.instance, fn)
}

// SetOnContextPopup
func (i *TImage) SetOnContextPopup(fn TContextPopupEvent) {
    Image_SetOnContextPopup(i.instance, fn)
}

// SetOnDblClick
func (i *TImage) SetOnDblClick(fn TNotifyEvent) {
    Image_SetOnDblClick(i.instance, fn)
}

// SetOnDragDrop
func (i *TImage) SetOnDragDrop(fn TDragDropEvent) {
    Image_SetOnDragDrop(i.instance, fn)
}

// SetOnDragOver
func (i *TImage) SetOnDragOver(fn TDragOverEvent) {
    Image_SetOnDragOver(i.instance, fn)
}

// SetOnEndDock
func (i *TImage) SetOnEndDock(fn TEndDragEvent) {
    Image_SetOnEndDock(i.instance, fn)
}

// SetOnEndDrag
func (i *TImage) SetOnEndDrag(fn TEndDragEvent) {
    Image_SetOnEndDrag(i.instance, fn)
}

// SetOnMouseDown
func (i *TImage) SetOnMouseDown(fn TMouseEvent) {
    Image_SetOnMouseDown(i.instance, fn)
}

// SetOnMouseEnter
func (i *TImage) SetOnMouseEnter(fn TNotifyEvent) {
    Image_SetOnMouseEnter(i.instance, fn)
}

// SetOnMouseLeave
func (i *TImage) SetOnMouseLeave(fn TNotifyEvent) {
    Image_SetOnMouseLeave(i.instance, fn)
}

// SetOnMouseMove
func (i *TImage) SetOnMouseMove(fn TMouseMoveEvent) {
    Image_SetOnMouseMove(i.instance, fn)
}

// SetOnMouseUp
func (i *TImage) SetOnMouseUp(fn TMouseEvent) {
    Image_SetOnMouseUp(i.instance, fn)
}

// SetOnStartDock
func (i *TImage) SetOnStartDock(fn TStartDockEvent) {
    Image_SetOnStartDock(i.instance, fn)
}

// Action
func (i *TImage) Action() *TAction {
    return ActionFromInst(Image_GetAction(i.instance))
}

// SetAction
func (i *TImage) SetAction(value IComponent) {
    Image_SetAction(i.instance, CheckPtr(value))
}

// BiDiMode
func (i *TImage) BiDiMode() TBiDiMode {
    return Image_GetBiDiMode(i.instance)
}

// SetBiDiMode
func (i *TImage) SetBiDiMode(value TBiDiMode) {
    Image_SetBiDiMode(i.instance, value)
}

// BoundsRect
func (i *TImage) BoundsRect() TRect {
    return Image_GetBoundsRect(i.instance)
}

// SetBoundsRect
func (i *TImage) SetBoundsRect(value TRect) {
    Image_SetBoundsRect(i.instance, value)
}

// ClientHeight
func (i *TImage) ClientHeight() int32 {
    return Image_GetClientHeight(i.instance)
}

// SetClientHeight
func (i *TImage) SetClientHeight(value int32) {
    Image_SetClientHeight(i.instance, value)
}

// ClientRect
func (i *TImage) ClientRect() TRect {
    return Image_GetClientRect(i.instance)
}

// ClientWidth
func (i *TImage) ClientWidth() int32 {
    return Image_GetClientWidth(i.instance)
}

// SetClientWidth
func (i *TImage) SetClientWidth(value int32) {
    Image_SetClientWidth(i.instance, value)
}

// ExplicitLeft
func (i *TImage) ExplicitLeft() int32 {
    return Image_GetExplicitLeft(i.instance)
}

// ExplicitTop
func (i *TImage) ExplicitTop() int32 {
    return Image_GetExplicitTop(i.instance)
}

// ExplicitWidth
func (i *TImage) ExplicitWidth() int32 {
    return Image_GetExplicitWidth(i.instance)
}

// ExplicitHeight
func (i *TImage) ExplicitHeight() int32 {
    return Image_GetExplicitHeight(i.instance)
}

// Floating
func (i *TImage) Floating() bool {
    return Image_GetFloating(i.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (i *TImage) Parent() *TWinControl {
    return WinControlFromInst(Image_GetParent(i.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (i *TImage) SetParent(value IWinControl) {
    Image_SetParent(i.instance, CheckPtr(value))
}

// StyleElements
func (i *TImage) StyleElements() TStyleElements {
    return Image_GetStyleElements(i.instance)
}

// SetStyleElements
func (i *TImage) SetStyleElements(value TStyleElements) {
    Image_SetStyleElements(i.instance, value)
}

// AlignWithMargins
func (i *TImage) AlignWithMargins() bool {
    return Image_GetAlignWithMargins(i.instance)
}

// SetAlignWithMargins
func (i *TImage) SetAlignWithMargins(value bool) {
    Image_SetAlignWithMargins(i.instance, value)
}

// Left
func (i *TImage) Left() int32 {
    return Image_GetLeft(i.instance)
}

// SetLeft
func (i *TImage) SetLeft(value int32) {
    Image_SetLeft(i.instance, value)
}

// Top
func (i *TImage) Top() int32 {
    return Image_GetTop(i.instance)
}

// SetTop
func (i *TImage) SetTop(value int32) {
    Image_SetTop(i.instance, value)
}

// Width
func (i *TImage) Width() int32 {
    return Image_GetWidth(i.instance)
}

// SetWidth
func (i *TImage) SetWidth(value int32) {
    Image_SetWidth(i.instance, value)
}

// Height
func (i *TImage) Height() int32 {
    return Image_GetHeight(i.instance)
}

// SetHeight
func (i *TImage) SetHeight(value int32) {
    Image_SetHeight(i.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (i *TImage) Cursor() TCursor {
    return Image_GetCursor(i.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (i *TImage) SetCursor(value TCursor) {
    Image_SetCursor(i.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (i *TImage) Hint() string {
    return Image_GetHint(i.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (i *TImage) SetHint(value string) {
    Image_SetHint(i.instance, value)
}

// Margins
func (i *TImage) Margins() *TMargins {
    return MarginsFromInst(Image_GetMargins(i.instance))
}

// SetMargins
func (i *TImage) SetMargins(value *TMargins) {
    Image_SetMargins(i.instance, CheckPtr(value))
}

// CustomHint
func (i *TImage) CustomHint() *TCustomHint {
    return CustomHintFromInst(Image_GetCustomHint(i.instance))
}

// SetCustomHint
func (i *TImage) SetCustomHint(value IComponent) {
    Image_SetCustomHint(i.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (i *TImage) ComponentCount() int32 {
    return Image_GetComponentCount(i.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (i *TImage) ComponentIndex() int32 {
    return Image_GetComponentIndex(i.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (i *TImage) SetComponentIndex(value int32) {
    Image_SetComponentIndex(i.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (i *TImage) Owner() *TComponent {
    return ComponentFromInst(Image_GetOwner(i.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (i *TImage) Name() string {
    return Image_GetName(i.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (i *TImage) SetName(value string) {
    Image_SetName(i.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (i *TImage) Tag() int {
    return Image_GetTag(i.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (i *TImage) SetTag(value int) {
    Image_SetTag(i.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (i *TImage) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Image_GetComponents(i.instance, AIndex))
}

