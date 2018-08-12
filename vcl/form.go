
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

type TForm struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewForm
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewForm(owner IComponent) *TForm {
    f := new(TForm)
    f.instance = Form_Create(CheckPtr(owner))
    f.ptr = unsafe.Pointer(f.instance)
    return f
}

// FormFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func FormFromInst(inst uintptr) *TForm {
    f := new(TForm)
    f.instance = inst
    f.ptr = unsafe.Pointer(inst)
    return f
}

// FormFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func FormFromObj(obj IObject) *TForm {
    f := new(TForm)
    f.instance = CheckPtr(obj)
    f.ptr = unsafe.Pointer(f.instance)
    return f
}

// FormFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func FormFromUnsafePointer(ptr unsafe.Pointer) *TForm {
    f := new(TForm)
    f.instance = uintptr(ptr)
    f.ptr = ptr
    return f
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (f *TForm) Free() {
    if f.instance != 0 {
        Form_Free(f.instance)
        f.instance = 0
        f.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (f *TForm) Instance() uintptr {
    return f.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (f *TForm) UnsafeAddr() unsafe.Pointer {
    return f.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (f *TForm) IsValid() bool {
    return f.instance != 0
}

// TFormClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TFormClass() TClass {
    return Form_StaticClassType()
}

// Close
func (f *TForm) Close() {
    Form_Close(f.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (f *TForm) Hide() {
    Form_Hide(f.instance)
}

// Print
func (f *TForm) Print() {
    Form_Print(f.instance)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (f *TForm) SetFocus() {
    Form_SetFocus(f.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (f *TForm) Show() {
    Form_Show(f.instance)
}

// ShowModal
func (f *TForm) ShowModal() int32 {
    return Form_ShowModal(f.instance)
}

// CanFocus
func (f *TForm) CanFocus() bool {
    return Form_CanFocus(f.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (f *TForm) ContainsControl(Control IControl) bool {
    return Form_ContainsControl(f.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (f *TForm) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(Form_ControlAtPos(f.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (f *TForm) DisableAlign() {
    Form_DisableAlign(f.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (f *TForm) EnableAlign() {
    Form_EnableAlign(f.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (f *TForm) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(Form_FindChildControl(f.instance, ControlName))
}

// FlipChildren
func (f *TForm) FlipChildren(AllLevels bool) {
    Form_FlipChildren(f.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (f *TForm) Focused() bool {
    return Form_Focused(f.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (f *TForm) HandleAllocated() bool {
    return Form_HandleAllocated(f.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (f *TForm) InsertControl(AControl IControl) {
    Form_InsertControl(f.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (f *TForm) Invalidate() {
    Form_Invalidate(f.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (f *TForm) PaintTo(DC HDC, X int32, Y int32) {
    Form_PaintTo(f.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (f *TForm) RemoveControl(AControl IControl) {
    Form_RemoveControl(f.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (f *TForm) Realign() {
    Form_Realign(f.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (f *TForm) Repaint() {
    Form_Repaint(f.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (f *TForm) ScaleBy(M int32, D int32) {
    Form_ScaleBy(f.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (f *TForm) ScrollBy(DeltaX int32, DeltaY int32) {
    Form_ScrollBy(f.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (f *TForm) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Form_SetBounds(f.instance, ALeft , ATop , AWidth , AHeight)
}

// Update
// CN: 控件更新。
// EN: Update.
func (f *TForm) Update() {
    Form_Update(f.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (f *TForm) UpdateControlState() {
    Form_UpdateControlState(f.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (f *TForm) BringToFront() {
    Form_BringToFront(f.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (f *TForm) ClientToScreen(Point TPoint) TPoint {
    return Form_ClientToScreen(f.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (f *TForm) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Form_ClientToParent(f.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (f *TForm) Dragging() bool {
    return Form_Dragging(f.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (f *TForm) HasParent() bool {
    return Form_HasParent(f.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (f *TForm) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Form_Perform(f.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (f *TForm) Refresh() {
    Form_Refresh(f.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (f *TForm) ScreenToClient(Point TPoint) TPoint {
    return Form_ScreenToClient(f.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (f *TForm) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Form_ParentToClient(f.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (f *TForm) SendToBack() {
    Form_SendToBack(f.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (f *TForm) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Form_GetTextBuf(f.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (f *TForm) GetTextLen() int32 {
    return Form_GetTextLen(f.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (f *TForm) SetTextBuf(Buffer string) {
    Form_SetTextBuf(f.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (f *TForm) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Form_FindComponent(f.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (f *TForm) GetNamePath() string {
    return Form_GetNamePath(f.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (f *TForm) Assign(Source IObject) {
    Form_Assign(f.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (f *TForm) DisposeOf() {
    Form_DisposeOf(f.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (f *TForm) ClassType() TClass {
    return Form_ClassType(f.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (f *TForm) ClassName() string {
    return Form_ClassName(f.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (f *TForm) InstanceSize() int32 {
    return Form_InstanceSize(f.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (f *TForm) InheritsFrom(AClass TClass) bool {
    return Form_InheritsFrom(f.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (f *TForm) Equals(Obj IObject) bool {
    return Form_Equals(f.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (f *TForm) GetHashCode() int32 {
    return Form_GetHashCode(f.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (f *TForm) ToString() string {
    return Form_ToString(f.instance)
}

// Action
func (f *TForm) Action() *TAction {
    return ActionFromInst(Form_GetAction(f.instance))
}

// SetAction
func (f *TForm) SetAction(value IComponent) {
    Form_SetAction(f.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (f *TForm) Align() TAlign {
    return Form_GetAlign(f.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (f *TForm) SetAlign(value TAlign) {
    Form_SetAlign(f.instance, value)
}

// AlphaBlend
func (f *TForm) AlphaBlend() bool {
    return Form_GetAlphaBlend(f.instance)
}

// SetAlphaBlend
func (f *TForm) SetAlphaBlend(value bool) {
    Form_SetAlphaBlend(f.instance, value)
}

// AlphaBlendValue
func (f *TForm) AlphaBlendValue() uint8 {
    return Form_GetAlphaBlendValue(f.instance)
}

// SetAlphaBlendValue
func (f *TForm) SetAlphaBlendValue(value uint8) {
    Form_SetAlphaBlendValue(f.instance, value)
}

// Anchors
func (f *TForm) Anchors() TAnchors {
    return Form_GetAnchors(f.instance)
}

// SetAnchors
func (f *TForm) SetAnchors(value TAnchors) {
    Form_SetAnchors(f.instance, value)
}

// AutoSize
func (f *TForm) AutoSize() bool {
    return Form_GetAutoSize(f.instance)
}

// SetAutoSize
func (f *TForm) SetAutoSize(value bool) {
    Form_SetAutoSize(f.instance, value)
}

// BiDiMode
func (f *TForm) BiDiMode() TBiDiMode {
    return Form_GetBiDiMode(f.instance)
}

// SetBiDiMode
func (f *TForm) SetBiDiMode(value TBiDiMode) {
    Form_SetBiDiMode(f.instance, value)
}

// BorderIcons
func (f *TForm) BorderIcons() TBorderIcons {
    return Form_GetBorderIcons(f.instance)
}

// SetBorderIcons
func (f *TForm) SetBorderIcons(value TBorderIcons) {
    Form_SetBorderIcons(f.instance, value)
}

// BorderStyle
func (f *TForm) BorderStyle() TFormBorderStyle {
    return Form_GetBorderStyle(f.instance)
}

// SetBorderStyle
func (f *TForm) SetBorderStyle(value TFormBorderStyle) {
    Form_SetBorderStyle(f.instance, value)
}

// BorderWidth
func (f *TForm) BorderWidth() int32 {
    return Form_GetBorderWidth(f.instance)
}

// SetBorderWidth
func (f *TForm) SetBorderWidth(value int32) {
    Form_SetBorderWidth(f.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (f *TForm) Caption() string {
    return Form_GetCaption(f.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (f *TForm) SetCaption(value string) {
    Form_SetCaption(f.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (f *TForm) ClientHeight() int32 {
    return Form_GetClientHeight(f.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (f *TForm) SetClientHeight(value int32) {
    Form_SetClientHeight(f.instance, value)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (f *TForm) ClientWidth() int32 {
    return Form_GetClientWidth(f.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (f *TForm) SetClientWidth(value int32) {
    Form_SetClientWidth(f.instance, value)
}

// Color
// CN: 获取设置颜色。
// EN: Get Set color.
func (f *TForm) Color() TColor {
    return Form_GetColor(f.instance)
}

// SetColor
// CN: 设置设置颜色。
// EN: Set Set color.
func (f *TForm) SetColor(value TColor) {
    Form_SetColor(f.instance, value)
}

// TransparentColor
func (f *TForm) TransparentColor() bool {
    return Form_GetTransparentColor(f.instance)
}

// SetTransparentColor
func (f *TForm) SetTransparentColor(value bool) {
    Form_SetTransparentColor(f.instance, value)
}

// TransparentColorValue
func (f *TForm) TransparentColorValue() TColor {
    return Form_GetTransparentColorValue(f.instance)
}

// SetTransparentColorValue
func (f *TForm) SetTransparentColorValue(value TColor) {
    Form_SetTransparentColorValue(f.instance, value)
}

// UseDockManager
func (f *TForm) UseDockManager() bool {
    return Form_GetUseDockManager(f.instance)
}

// SetUseDockManager
func (f *TForm) SetUseDockManager(value bool) {
    Form_SetUseDockManager(f.instance, value)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (f *TForm) DockSite() bool {
    return Form_GetDockSite(f.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (f *TForm) SetDockSite(value bool) {
    Form_SetDockSite(f.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (f *TForm) DoubleBuffered() bool {
    return Form_GetDoubleBuffered(f.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (f *TForm) SetDoubleBuffered(value bool) {
    Form_SetDoubleBuffered(f.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (f *TForm) DragKind() TDragKind {
    return Form_GetDragKind(f.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (f *TForm) SetDragKind(value TDragKind) {
    Form_SetDragKind(f.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (f *TForm) DragMode() TDragMode {
    return Form_GetDragMode(f.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (f *TForm) SetDragMode(value TDragMode) {
    Form_SetDragMode(f.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (f *TForm) Enabled() bool {
    return Form_GetEnabled(f.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (f *TForm) SetEnabled(value bool) {
    Form_SetEnabled(f.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (f *TForm) ParentFont() bool {
    return Form_GetParentFont(f.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (f *TForm) SetParentFont(value bool) {
    Form_SetParentFont(f.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (f *TForm) Font() *TFont {
    return FontFromInst(Form_GetFont(f.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (f *TForm) SetFont(value *TFont) {
    Form_SetFont(f.instance, CheckPtr(value))
}

// FormStyle
func (f *TForm) FormStyle() TFormStyle {
    return Form_GetFormStyle(f.instance)
}

// SetFormStyle
func (f *TForm) SetFormStyle(value TFormStyle) {
    Form_SetFormStyle(f.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (f *TForm) Height() int32 {
    return Form_GetHeight(f.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (f *TForm) SetHeight(value int32) {
    Form_SetHeight(f.instance, value)
}

// Icon
// CN: 获取图标。
// EN: Get icon.
func (f *TForm) Icon() *TIcon {
    return IconFromInst(Form_GetIcon(f.instance))
}

// SetIcon
// CN: 设置图标。
// EN: Set icon.
func (f *TForm) SetIcon(value *TIcon) {
    Form_SetIcon(f.instance, CheckPtr(value))
}

// KeyPreview
func (f *TForm) KeyPreview() bool {
    return Form_GetKeyPreview(f.instance)
}

// SetKeyPreview
func (f *TForm) SetKeyPreview(value bool) {
    Form_SetKeyPreview(f.instance, value)
}

// Menu
func (f *TForm) Menu() *TMainMenu {
    return MainMenuFromInst(Form_GetMenu(f.instance))
}

// SetMenu
func (f *TForm) SetMenu(value IComponent) {
    Form_SetMenu(f.instance, CheckPtr(value))
}

// PixelsPerInch
func (f *TForm) PixelsPerInch() int32 {
    return Form_GetPixelsPerInch(f.instance)
}

// SetPixelsPerInch
func (f *TForm) SetPixelsPerInch(value int32) {
    Form_SetPixelsPerInch(f.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (f *TForm) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Form_GetPopupMenu(f.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (f *TForm) SetPopupMenu(value IComponent) {
    Form_SetPopupMenu(f.instance, CheckPtr(value))
}

// Position
func (f *TForm) Position() TPosition {
    return Form_GetPosition(f.instance)
}

// SetPosition
func (f *TForm) SetPosition(value TPosition) {
    Form_SetPosition(f.instance, value)
}

// Scaled
func (f *TForm) Scaled() bool {
    return Form_GetScaled(f.instance)
}

// SetScaled
func (f *TForm) SetScaled(value bool) {
    Form_SetScaled(f.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (f *TForm) ShowHint() bool {
    return Form_GetShowHint(f.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (f *TForm) SetShowHint(value bool) {
    Form_SetShowHint(f.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (f *TForm) Visible() bool {
    return Form_GetVisible(f.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (f *TForm) SetVisible(value bool) {
    Form_SetVisible(f.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (f *TForm) Width() int32 {
    return Form_GetWidth(f.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (f *TForm) SetWidth(value int32) {
    Form_SetWidth(f.instance, value)
}

// WindowState
func (f *TForm) WindowState() TWindowState {
    return Form_GetWindowState(f.instance)
}

// SetWindowState
func (f *TForm) SetWindowState(value TWindowState) {
    Form_SetWindowState(f.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (f *TForm) StyleElements() TStyleElements {
    return Form_GetStyleElements(f.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (f *TForm) SetStyleElements(value TStyleElements) {
    Form_SetStyleElements(f.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (f *TForm) SetOnClick(fn TNotifyEvent) {
    Form_SetOnClick(f.instance, fn)
}

// SetOnClose
func (f *TForm) SetOnClose(fn TCloseEvent) {
    Form_SetOnClose(f.instance, fn)
}

// SetOnCloseQuery
func (f *TForm) SetOnCloseQuery(fn TCloseQueryEvent) {
    Form_SetOnCloseQuery(f.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (f *TForm) SetOnContextPopup(fn TContextPopupEvent) {
    Form_SetOnContextPopup(f.instance, fn)
}

// SetOnDblClick
func (f *TForm) SetOnDblClick(fn TNotifyEvent) {
    Form_SetOnDblClick(f.instance, fn)
}

// SetOnDockDrop
func (f *TForm) SetOnDockDrop(fn TDockDropEvent) {
    Form_SetOnDockDrop(f.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (f *TForm) SetOnDragDrop(fn TDragDropEvent) {
    Form_SetOnDragDrop(f.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (f *TForm) SetOnDragOver(fn TDragOverEvent) {
    Form_SetOnDragOver(f.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (f *TForm) SetOnEndDock(fn TEndDragEvent) {
    Form_SetOnEndDock(f.instance, fn)
}

// SetOnGetSiteInfo
func (f *TForm) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    Form_SetOnGetSiteInfo(f.instance, fn)
}

// SetOnHide
func (f *TForm) SetOnHide(fn TNotifyEvent) {
    Form_SetOnHide(f.instance, fn)
}

// SetOnHelp
func (f *TForm) SetOnHelp(fn THelpEvent) {
    Form_SetOnHelp(f.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (f *TForm) SetOnKeyDown(fn TKeyEvent) {
    Form_SetOnKeyDown(f.instance, fn)
}

// SetOnKeyPress
func (f *TForm) SetOnKeyPress(fn TKeyPressEvent) {
    Form_SetOnKeyPress(f.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (f *TForm) SetOnKeyUp(fn TKeyEvent) {
    Form_SetOnKeyUp(f.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (f *TForm) SetOnMouseDown(fn TMouseEvent) {
    Form_SetOnMouseDown(f.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (f *TForm) SetOnMouseEnter(fn TNotifyEvent) {
    Form_SetOnMouseEnter(f.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (f *TForm) SetOnMouseLeave(fn TNotifyEvent) {
    Form_SetOnMouseLeave(f.instance, fn)
}

// SetOnMouseMove
func (f *TForm) SetOnMouseMove(fn TMouseMoveEvent) {
    Form_SetOnMouseMove(f.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (f *TForm) SetOnMouseUp(fn TMouseEvent) {
    Form_SetOnMouseUp(f.instance, fn)
}

// SetOnMouseWheel
func (f *TForm) SetOnMouseWheel(fn TMouseWheelEvent) {
    Form_SetOnMouseWheel(f.instance, fn)
}

// SetOnMouseWheelDown
func (f *TForm) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    Form_SetOnMouseWheelDown(f.instance, fn)
}

// SetOnMouseWheelUp
func (f *TForm) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    Form_SetOnMouseWheelUp(f.instance, fn)
}

// SetOnPaint
func (f *TForm) SetOnPaint(fn TNotifyEvent) {
    Form_SetOnPaint(f.instance, fn)
}

// SetOnResize
func (f *TForm) SetOnResize(fn TNotifyEvent) {
    Form_SetOnResize(f.instance, fn)
}

// SetOnShortCut
func (f *TForm) SetOnShortCut(fn TShortCutEvent) {
    Form_SetOnShortCut(f.instance, fn)
}

// SetOnShow
func (f *TForm) SetOnShow(fn TNotifyEvent) {
    Form_SetOnShow(f.instance, fn)
}

// SetOnStartDock
func (f *TForm) SetOnStartDock(fn TStartDockEvent) {
    Form_SetOnStartDock(f.instance, fn)
}

// SetOnUnDock
func (f *TForm) SetOnUnDock(fn TUnDockEvent) {
    Form_SetOnUnDock(f.instance, fn)
}

// Canvas
func (f *TForm) Canvas() *TCanvas {
    return CanvasFromInst(Form_GetCanvas(f.instance))
}

// DropTarget
func (f *TForm) DropTarget() bool {
    return Form_GetDropTarget(f.instance)
}

// SetDropTarget
func (f *TForm) SetDropTarget(value bool) {
    Form_SetDropTarget(f.instance, value)
}

// ModalResult
func (f *TForm) ModalResult() TModalResult {
    return Form_GetModalResult(f.instance)
}

// SetModalResult
func (f *TForm) SetModalResult(value TModalResult) {
    Form_SetModalResult(f.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (f *TForm) Left() int32 {
    return Form_GetLeft(f.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (f *TForm) SetLeft(value int32) {
    Form_SetLeft(f.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (f *TForm) Top() int32 {
    return Form_GetTop(f.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (f *TForm) SetTop(value int32) {
    Form_SetTop(f.instance, value)
}

// DockClientCount
func (f *TForm) DockClientCount() int32 {
    return Form_GetDockClientCount(f.instance)
}

// AlignDisabled
func (f *TForm) AlignDisabled() bool {
    return Form_GetAlignDisabled(f.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (f *TForm) MouseInClient() bool {
    return Form_GetMouseInClient(f.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (f *TForm) VisibleDockClientCount() int32 {
    return Form_GetVisibleDockClientCount(f.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (f *TForm) Brush() *TBrush {
    return BrushFromInst(Form_GetBrush(f.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (f *TForm) ControlCount() int32 {
    return Form_GetControlCount(f.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (f *TForm) Handle() HWND {
    return Form_GetHandle(f.instance)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (f *TForm) ParentDoubleBuffered() bool {
    return Form_GetParentDoubleBuffered(f.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (f *TForm) SetParentDoubleBuffered(value bool) {
    Form_SetParentDoubleBuffered(f.instance, value)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (f *TForm) ParentWindow() HWND {
    return Form_GetParentWindow(f.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (f *TForm) SetParentWindow(value HWND) {
    Form_SetParentWindow(f.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (f *TForm) TabOrder() TTabOrder {
    return Form_GetTabOrder(f.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (f *TForm) SetTabOrder(value TTabOrder) {
    Form_SetTabOrder(f.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (f *TForm) TabStop() bool {
    return Form_GetTabStop(f.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (f *TForm) SetTabStop(value bool) {
    Form_SetTabStop(f.instance, value)
}

// BoundsRect
func (f *TForm) BoundsRect() TRect {
    return Form_GetBoundsRect(f.instance)
}

// SetBoundsRect
func (f *TForm) SetBoundsRect(value TRect) {
    Form_SetBoundsRect(f.instance, value)
}

// ClientOrigin
func (f *TForm) ClientOrigin() TPoint {
    return Form_GetClientOrigin(f.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (f *TForm) ClientRect() TRect {
    return Form_GetClientRect(f.instance)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (f *TForm) ControlState() TControlState {
    return Form_GetControlState(f.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (f *TForm) SetControlState(value TControlState) {
    Form_SetControlState(f.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (f *TForm) ControlStyle() TControlStyle {
    return Form_GetControlStyle(f.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (f *TForm) SetControlStyle(value TControlStyle) {
    Form_SetControlStyle(f.instance, value)
}

// ExplicitLeft
func (f *TForm) ExplicitLeft() int32 {
    return Form_GetExplicitLeft(f.instance)
}

// ExplicitTop
func (f *TForm) ExplicitTop() int32 {
    return Form_GetExplicitTop(f.instance)
}

// ExplicitWidth
func (f *TForm) ExplicitWidth() int32 {
    return Form_GetExplicitWidth(f.instance)
}

// ExplicitHeight
func (f *TForm) ExplicitHeight() int32 {
    return Form_GetExplicitHeight(f.instance)
}

// Floating
func (f *TForm) Floating() bool {
    return Form_GetFloating(f.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (f *TForm) Parent() *TWinControl {
    return WinControlFromInst(Form_GetParent(f.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (f *TForm) SetParent(value IWinControl) {
    Form_SetParent(f.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (f *TForm) AlignWithMargins() bool {
    return Form_GetAlignWithMargins(f.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (f *TForm) SetAlignWithMargins(value bool) {
    Form_SetAlignWithMargins(f.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (f *TForm) Cursor() TCursor {
    return Form_GetCursor(f.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (f *TForm) SetCursor(value TCursor) {
    Form_SetCursor(f.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (f *TForm) Hint() string {
    return Form_GetHint(f.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (f *TForm) SetHint(value string) {
    Form_SetHint(f.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (f *TForm) Margins() *TMargins {
    return MarginsFromInst(Form_GetMargins(f.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (f *TForm) SetMargins(value *TMargins) {
    Form_SetMargins(f.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (f *TForm) CustomHint() *TCustomHint {
    return CustomHintFromInst(Form_GetCustomHint(f.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (f *TForm) SetCustomHint(value IComponent) {
    Form_SetCustomHint(f.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (f *TForm) ComponentCount() int32 {
    return Form_GetComponentCount(f.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (f *TForm) ComponentIndex() int32 {
    return Form_GetComponentIndex(f.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (f *TForm) SetComponentIndex(value int32) {
    Form_SetComponentIndex(f.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (f *TForm) Owner() *TComponent {
    return ComponentFromInst(Form_GetOwner(f.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (f *TForm) Name() string {
    return Form_GetName(f.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (f *TForm) SetName(value string) {
    Form_SetName(f.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (f *TForm) Tag() int {
    return Form_GetTag(f.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (f *TForm) SetTag(value int) {
    Form_SetTag(f.instance, value)
}

// DockClients
func (f *TForm) DockClients(Index int32) *TControl {
    return ControlFromInst(Form_GetDockClients(f.instance, Index))
}

// Controls
func (f *TForm) Controls(Index int32) *TControl {
    return ControlFromInst(Form_GetControls(f.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (f *TForm) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Form_GetComponents(f.instance, AIndex))
}

