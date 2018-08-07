
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

type TImageButton struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewImageButton
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewImageButton(owner IComponent) *TImageButton {
    i := new(TImageButton)
    i.instance = ImageButton_Create(CheckPtr(owner))
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// ImageButtonFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ImageButtonFromInst(inst uintptr) *TImageButton {
    i := new(TImageButton)
    i.instance = inst
    i.ptr = unsafe.Pointer(inst)
    return i
}

// ImageButtonFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ImageButtonFromObj(obj IObject) *TImageButton {
    i := new(TImageButton)
    i.instance = CheckPtr(obj)
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// ImageButtonFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ImageButtonFromUnsafePointer(ptr unsafe.Pointer) *TImageButton {
    i := new(TImageButton)
    i.instance = uintptr(ptr)
    i.ptr = ptr
    return i
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (i *TImageButton) Free() {
    if i.instance != 0 {
        ImageButton_Free(i.instance)
        i.instance = 0
        i.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (i *TImageButton) Instance() uintptr {
    return i.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (i *TImageButton) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (i *TImageButton) IsValid() bool {
    return i.instance != 0
}

// TImageButtonClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TImageButtonClass() TClass {
    return ImageButton_StaticClassType()
}

// Click
func (i *TImageButton) Click() {
    ImageButton_Click(i.instance)
}

// BringToFront
func (i *TImageButton) BringToFront() {
    ImageButton_BringToFront(i.instance)
}

// ClientToScreen
func (i *TImageButton) ClientToScreen(Point TPoint) TPoint {
    return ImageButton_ClientToScreen(i.instance, Point)
}

// ClientToParent
func (i *TImageButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ImageButton_ClientToParent(i.instance, Point , CheckPtr(AParent))
}

// Dragging
func (i *TImageButton) Dragging() bool {
    return ImageButton_Dragging(i.instance)
}

// HasParent
func (i *TImageButton) HasParent() bool {
    return ImageButton_HasParent(i.instance)
}

// Hide
func (i *TImageButton) Hide() {
    ImageButton_Hide(i.instance)
}

// Invalidate
func (i *TImageButton) Invalidate() {
    ImageButton_Invalidate(i.instance)
}

// Perform
func (i *TImageButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ImageButton_Perform(i.instance, Msg , WParam , LParam)
}

// Refresh
func (i *TImageButton) Refresh() {
    ImageButton_Refresh(i.instance)
}

// Repaint
func (i *TImageButton) Repaint() {
    ImageButton_Repaint(i.instance)
}

// ScreenToClient
func (i *TImageButton) ScreenToClient(Point TPoint) TPoint {
    return ImageButton_ScreenToClient(i.instance, Point)
}

// ParentToClient
func (i *TImageButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ImageButton_ParentToClient(i.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (i *TImageButton) SendToBack() {
    ImageButton_SendToBack(i.instance)
}

// SetBounds
func (i *TImageButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ImageButton_SetBounds(i.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
func (i *TImageButton) Show() {
    ImageButton_Show(i.instance)
}

// Update
func (i *TImageButton) Update() {
    ImageButton_Update(i.instance)
}

// GetTextBuf
func (i *TImageButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ImageButton_GetTextBuf(i.instance, Buffer , BufSize)
}

// GetTextLen
func (i *TImageButton) GetTextLen() int32 {
    return ImageButton_GetTextLen(i.instance)
}

// FindComponent
func (i *TImageButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ImageButton_FindComponent(i.instance, AName))
}

// GetNamePath
func (i *TImageButton) GetNamePath() string {
    return ImageButton_GetNamePath(i.instance)
}

// Assign
func (i *TImageButton) Assign(Source IObject) {
    ImageButton_Assign(i.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (i *TImageButton) DisposeOf() {
    ImageButton_DisposeOf(i.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (i *TImageButton) ClassType() TClass {
    return ImageButton_ClassType(i.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (i *TImageButton) ClassName() string {
    return ImageButton_ClassName(i.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (i *TImageButton) InstanceSize() int32 {
    return ImageButton_InstanceSize(i.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (i *TImageButton) InheritsFrom(AClass TClass) bool {
    return ImageButton_InheritsFrom(i.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (i *TImageButton) Equals(Obj IObject) bool {
    return ImageButton_Equals(i.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (i *TImageButton) GetHashCode() int32 {
    return ImageButton_GetHashCode(i.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (i *TImageButton) ToString() string {
    return ImageButton_ToString(i.instance)
}

// Action
func (i *TImageButton) Action() *TAction {
    return ActionFromInst(ImageButton_GetAction(i.instance))
}

// SetAction
func (i *TImageButton) SetAction(value IComponent) {
    ImageButton_SetAction(i.instance, CheckPtr(value))
}

// Align
func (i *TImageButton) Align() TAlign {
    return ImageButton_GetAlign(i.instance)
}

// SetAlign
func (i *TImageButton) SetAlign(value TAlign) {
    ImageButton_SetAlign(i.instance, value)
}

// Anchors
func (i *TImageButton) Anchors() TAnchors {
    return ImageButton_GetAnchors(i.instance)
}

// SetAnchors
func (i *TImageButton) SetAnchors(value TAnchors) {
    ImageButton_SetAnchors(i.instance, value)
}

// AutoSize
func (i *TImageButton) AutoSize() bool {
    return ImageButton_GetAutoSize(i.instance)
}

// SetAutoSize
func (i *TImageButton) SetAutoSize(value bool) {
    ImageButton_SetAutoSize(i.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (i *TImageButton) Caption() string {
    return ImageButton_GetCaption(i.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (i *TImageButton) SetCaption(value string) {
    ImageButton_SetCaption(i.instance, value)
}

// DragCursor
func (i *TImageButton) DragCursor() TCursor {
    return ImageButton_GetDragCursor(i.instance)
}

// SetDragCursor
func (i *TImageButton) SetDragCursor(value TCursor) {
    ImageButton_SetDragCursor(i.instance, value)
}

// DragKind
func (i *TImageButton) DragKind() TDragKind {
    return ImageButton_GetDragKind(i.instance)
}

// SetDragKind
func (i *TImageButton) SetDragKind(value TDragKind) {
    ImageButton_SetDragKind(i.instance, value)
}

// DragMode
func (i *TImageButton) DragMode() TDragMode {
    return ImageButton_GetDragMode(i.instance)
}

// SetDragMode
func (i *TImageButton) SetDragMode(value TDragMode) {
    ImageButton_SetDragMode(i.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (i *TImageButton) Enabled() bool {
    return ImageButton_GetEnabled(i.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (i *TImageButton) SetEnabled(value bool) {
    ImageButton_SetEnabled(i.instance, value)
}

// Font
func (i *TImageButton) Font() *TFont {
    return FontFromInst(ImageButton_GetFont(i.instance))
}

// SetFont
func (i *TImageButton) SetFont(value *TFont) {
    ImageButton_SetFont(i.instance, CheckPtr(value))
}

// ImageCount
func (i *TImageButton) ImageCount() int32 {
    return ImageButton_GetImageCount(i.instance)
}

// SetImageCount
func (i *TImageButton) SetImageCount(value int32) {
    ImageButton_SetImageCount(i.instance, value)
}

// ModalResult
func (i *TImageButton) ModalResult() TModalResult {
    return ImageButton_GetModalResult(i.instance)
}

// SetModalResult
func (i *TImageButton) SetModalResult(value TModalResult) {
    ImageButton_SetModalResult(i.instance, value)
}

// ParentShowHint
func (i *TImageButton) ParentShowHint() bool {
    return ImageButton_GetParentShowHint(i.instance)
}

// SetParentShowHint
func (i *TImageButton) SetParentShowHint(value bool) {
    ImageButton_SetParentShowHint(i.instance, value)
}

// ParentFont
func (i *TImageButton) ParentFont() bool {
    return ImageButton_GetParentFont(i.instance)
}

// SetParentFont
func (i *TImageButton) SetParentFont(value bool) {
    ImageButton_SetParentFont(i.instance, value)
}

// Picture
func (i *TImageButton) Picture() *TPicture {
    return PictureFromInst(ImageButton_GetPicture(i.instance))
}

// SetPicture
func (i *TImageButton) SetPicture(value *TPicture) {
    ImageButton_SetPicture(i.instance, CheckPtr(value))
}

// PopupMenu
func (i *TImageButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ImageButton_GetPopupMenu(i.instance))
}

// SetPopupMenu
func (i *TImageButton) SetPopupMenu(value IComponent) {
    ImageButton_SetPopupMenu(i.instance, CheckPtr(value))
}

// ShowHint
func (i *TImageButton) ShowHint() bool {
    return ImageButton_GetShowHint(i.instance)
}

// SetShowHint
func (i *TImageButton) SetShowHint(value bool) {
    ImageButton_SetShowHint(i.instance, value)
}

// ShowCaption
func (i *TImageButton) ShowCaption() bool {
    return ImageButton_GetShowCaption(i.instance)
}

// SetShowCaption
func (i *TImageButton) SetShowCaption(value bool) {
    ImageButton_SetShowCaption(i.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (i *TImageButton) Visible() bool {
    return ImageButton_GetVisible(i.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (i *TImageButton) SetVisible(value bool) {
    ImageButton_SetVisible(i.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (i *TImageButton) SetOnClick(fn TNotifyEvent) {
    ImageButton_SetOnClick(i.instance, fn)
}

// SetOnContextPopup
func (i *TImageButton) SetOnContextPopup(fn TContextPopupEvent) {
    ImageButton_SetOnContextPopup(i.instance, fn)
}

// SetOnDblClick
func (i *TImageButton) SetOnDblClick(fn TNotifyEvent) {
    ImageButton_SetOnDblClick(i.instance, fn)
}

// SetOnDragDrop
func (i *TImageButton) SetOnDragDrop(fn TDragDropEvent) {
    ImageButton_SetOnDragDrop(i.instance, fn)
}

// SetOnDragOver
func (i *TImageButton) SetOnDragOver(fn TDragOverEvent) {
    ImageButton_SetOnDragOver(i.instance, fn)
}

// SetOnEndDock
func (i *TImageButton) SetOnEndDock(fn TEndDragEvent) {
    ImageButton_SetOnEndDock(i.instance, fn)
}

// SetOnEndDrag
func (i *TImageButton) SetOnEndDrag(fn TEndDragEvent) {
    ImageButton_SetOnEndDrag(i.instance, fn)
}

// SetOnMouseDown
func (i *TImageButton) SetOnMouseDown(fn TMouseEvent) {
    ImageButton_SetOnMouseDown(i.instance, fn)
}

// SetOnMouseEnter
func (i *TImageButton) SetOnMouseEnter(fn TNotifyEvent) {
    ImageButton_SetOnMouseEnter(i.instance, fn)
}

// SetOnMouseLeave
func (i *TImageButton) SetOnMouseLeave(fn TNotifyEvent) {
    ImageButton_SetOnMouseLeave(i.instance, fn)
}

// SetOnMouseMove
func (i *TImageButton) SetOnMouseMove(fn TMouseMoveEvent) {
    ImageButton_SetOnMouseMove(i.instance, fn)
}

// SetOnMouseUp
func (i *TImageButton) SetOnMouseUp(fn TMouseEvent) {
    ImageButton_SetOnMouseUp(i.instance, fn)
}

// BiDiMode
func (i *TImageButton) BiDiMode() TBiDiMode {
    return ImageButton_GetBiDiMode(i.instance)
}

// SetBiDiMode
func (i *TImageButton) SetBiDiMode(value TBiDiMode) {
    ImageButton_SetBiDiMode(i.instance, value)
}

// BoundsRect
func (i *TImageButton) BoundsRect() TRect {
    return ImageButton_GetBoundsRect(i.instance)
}

// SetBoundsRect
func (i *TImageButton) SetBoundsRect(value TRect) {
    ImageButton_SetBoundsRect(i.instance, value)
}

// ClientHeight
func (i *TImageButton) ClientHeight() int32 {
    return ImageButton_GetClientHeight(i.instance)
}

// SetClientHeight
func (i *TImageButton) SetClientHeight(value int32) {
    ImageButton_SetClientHeight(i.instance, value)
}

// ClientRect
func (i *TImageButton) ClientRect() TRect {
    return ImageButton_GetClientRect(i.instance)
}

// ClientWidth
func (i *TImageButton) ClientWidth() int32 {
    return ImageButton_GetClientWidth(i.instance)
}

// SetClientWidth
func (i *TImageButton) SetClientWidth(value int32) {
    ImageButton_SetClientWidth(i.instance, value)
}

// ExplicitLeft
func (i *TImageButton) ExplicitLeft() int32 {
    return ImageButton_GetExplicitLeft(i.instance)
}

// ExplicitTop
func (i *TImageButton) ExplicitTop() int32 {
    return ImageButton_GetExplicitTop(i.instance)
}

// ExplicitWidth
func (i *TImageButton) ExplicitWidth() int32 {
    return ImageButton_GetExplicitWidth(i.instance)
}

// ExplicitHeight
func (i *TImageButton) ExplicitHeight() int32 {
    return ImageButton_GetExplicitHeight(i.instance)
}

// Floating
func (i *TImageButton) Floating() bool {
    return ImageButton_GetFloating(i.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (i *TImageButton) Parent() *TWinControl {
    return WinControlFromInst(ImageButton_GetParent(i.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (i *TImageButton) SetParent(value IWinControl) {
    ImageButton_SetParent(i.instance, CheckPtr(value))
}

// StyleElements
func (i *TImageButton) StyleElements() TStyleElements {
    return ImageButton_GetStyleElements(i.instance)
}

// SetStyleElements
func (i *TImageButton) SetStyleElements(value TStyleElements) {
    ImageButton_SetStyleElements(i.instance, value)
}

// AlignWithMargins
func (i *TImageButton) AlignWithMargins() bool {
    return ImageButton_GetAlignWithMargins(i.instance)
}

// SetAlignWithMargins
func (i *TImageButton) SetAlignWithMargins(value bool) {
    ImageButton_SetAlignWithMargins(i.instance, value)
}

// Left
func (i *TImageButton) Left() int32 {
    return ImageButton_GetLeft(i.instance)
}

// SetLeft
func (i *TImageButton) SetLeft(value int32) {
    ImageButton_SetLeft(i.instance, value)
}

// Top
func (i *TImageButton) Top() int32 {
    return ImageButton_GetTop(i.instance)
}

// SetTop
func (i *TImageButton) SetTop(value int32) {
    ImageButton_SetTop(i.instance, value)
}

// Width
func (i *TImageButton) Width() int32 {
    return ImageButton_GetWidth(i.instance)
}

// SetWidth
func (i *TImageButton) SetWidth(value int32) {
    ImageButton_SetWidth(i.instance, value)
}

// Height
func (i *TImageButton) Height() int32 {
    return ImageButton_GetHeight(i.instance)
}

// SetHeight
func (i *TImageButton) SetHeight(value int32) {
    ImageButton_SetHeight(i.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (i *TImageButton) Cursor() TCursor {
    return ImageButton_GetCursor(i.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (i *TImageButton) SetCursor(value TCursor) {
    ImageButton_SetCursor(i.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (i *TImageButton) Hint() string {
    return ImageButton_GetHint(i.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (i *TImageButton) SetHint(value string) {
    ImageButton_SetHint(i.instance, value)
}

// Margins
func (i *TImageButton) Margins() *TMargins {
    return MarginsFromInst(ImageButton_GetMargins(i.instance))
}

// SetMargins
func (i *TImageButton) SetMargins(value *TMargins) {
    ImageButton_SetMargins(i.instance, CheckPtr(value))
}

// CustomHint
func (i *TImageButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(ImageButton_GetCustomHint(i.instance))
}

// SetCustomHint
func (i *TImageButton) SetCustomHint(value IComponent) {
    ImageButton_SetCustomHint(i.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (i *TImageButton) ComponentCount() int32 {
    return ImageButton_GetComponentCount(i.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (i *TImageButton) ComponentIndex() int32 {
    return ImageButton_GetComponentIndex(i.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (i *TImageButton) SetComponentIndex(value int32) {
    ImageButton_SetComponentIndex(i.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (i *TImageButton) Owner() *TComponent {
    return ComponentFromInst(ImageButton_GetOwner(i.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (i *TImageButton) Name() string {
    return ImageButton_GetName(i.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (i *TImageButton) SetName(value string) {
    ImageButton_SetName(i.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (i *TImageButton) Tag() int {
    return ImageButton_GetTag(i.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (i *TImageButton) SetTag(value int) {
    ImageButton_SetTag(i.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (i *TImageButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ImageButton_GetComponents(i.instance, AIndex))
}

