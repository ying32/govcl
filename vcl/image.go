
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TImage struct {
    IControl
    instance uintptr
}

func NewImage(owner IComponent) *TImage {
    i := new(TImage)
    i.instance = Image_Create(CheckPtr(owner))
    return i
}

func ImageFromInst(inst uintptr) *TImage {
    i := new(TImage)
    i.instance = inst
    return i
}

func ImageFromObj(obj IObject) *TImage {
    i := new(TImage)
    i.instance = CheckPtr(obj)
    return i
}

func (i *TImage) Free() {
    if i.instance != 0 {
        Image_Free(i.instance)
        i.instance = 0
    }
}

func (i *TImage) Instance() uintptr {
    return i.instance
}

func (i *TImage) IsValid() bool {
    return i.instance != 0
}

func (i *TImage) BringToFront() {
    Image_BringToFront(i.instance)
}

func (i *TImage) HasParent() bool {
    return Image_HasParent(i.instance)
}

func (i *TImage) Hide() {
    Image_Hide(i.instance)
}

func (i *TImage) Invalidate() {
    Image_Invalidate(i.instance)
}

func (i *TImage) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Image_Perform(i.instance, Msg , WParam , LParam)
}

func (i *TImage) Refresh() {
    Image_Refresh(i.instance)
}

func (i *TImage) Repaint() {
    Image_Repaint(i.instance)
}

func (i *TImage) SendToBack() {
    Image_SendToBack(i.instance)
}

func (i *TImage) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Image_SetBounds(i.instance, ALeft , ATop , AWidth , AHeight)
}

func (i *TImage) Show() {
    Image_Show(i.instance)
}

func (i *TImage) Update() {
    Image_Update(i.instance)
}

func (i *TImage) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Image_GetTextBuf(i.instance, Buffer , BufSize)
}

func (i *TImage) GetTextLen() int32 {
    return Image_GetTextLen(i.instance)
}

func (i *TImage) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Image_FindComponent(i.instance, AName))
}

func (i *TImage) GetNamePath() string {
    return Image_GetNamePath(i.instance)
}

func (i *TImage) Assign(Source IObject) {
    Image_Assign(i.instance, CheckPtr(Source))
}

func (i *TImage) ClassName() string {
    return Image_ClassName(i.instance)
}

func (i *TImage) Equals(Obj IObject) bool {
    return Image_Equals(i.instance, CheckPtr(Obj))
}

func (i *TImage) GetHashCode() int32 {
    return Image_GetHashCode(i.instance)
}

func (i *TImage) ToString() string {
    return Image_ToString(i.instance)
}

func (i *TImage) Canvas() *TCanvas {
    return CanvasFromInst(Image_GetCanvas(i.instance))
}

func (i *TImage) Align() TAlign {
    return Image_GetAlign(i.instance)
}

func (i *TImage) SetAlign(value TAlign) {
    Image_SetAlign(i.instance, value)
}

func (i *TImage) Anchors() TAnchors {
    return Image_GetAnchors(i.instance)
}

func (i *TImage) SetAnchors(value TAnchors) {
    Image_SetAnchors(i.instance, value)
}

func (i *TImage) AutoSize() bool {
    return Image_GetAutoSize(i.instance)
}

func (i *TImage) SetAutoSize(value bool) {
    Image_SetAutoSize(i.instance, value)
}

func (i *TImage) Center() bool {
    return Image_GetCenter(i.instance)
}

func (i *TImage) SetCenter(value bool) {
    Image_SetCenter(i.instance, value)
}

func (i *TImage) Enabled() bool {
    return Image_GetEnabled(i.instance)
}

func (i *TImage) SetEnabled(value bool) {
    Image_SetEnabled(i.instance, value)
}

func (i *TImage) IncrementalDisplay() bool {
    return Image_GetIncrementalDisplay(i.instance)
}

func (i *TImage) SetIncrementalDisplay(value bool) {
    Image_SetIncrementalDisplay(i.instance, value)
}

func (i *TImage) ParentShowHint() bool {
    return Image_GetParentShowHint(i.instance)
}

func (i *TImage) SetParentShowHint(value bool) {
    Image_SetParentShowHint(i.instance, value)
}

func (i *TImage) Picture() *TPicture {
    return PictureFromInst(Image_GetPicture(i.instance))
}

func (i *TImage) SetPicture(value *TPicture) {
    Image_SetPicture(i.instance, CheckPtr(value))
}

func (i *TImage) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Image_GetPopupMenu(i.instance))
}

func (i *TImage) SetPopupMenu(value IComponent) {
    Image_SetPopupMenu(i.instance, CheckPtr(value))
}

func (i *TImage) Proportional() bool {
    return Image_GetProportional(i.instance)
}

func (i *TImage) SetProportional(value bool) {
    Image_SetProportional(i.instance, value)
}

func (i *TImage) ShowHint() bool {
    return Image_GetShowHint(i.instance)
}

func (i *TImage) SetShowHint(value bool) {
    Image_SetShowHint(i.instance, value)
}

func (i *TImage) Stretch() bool {
    return Image_GetStretch(i.instance)
}

func (i *TImage) SetStretch(value bool) {
    Image_SetStretch(i.instance, value)
}

func (i *TImage) Transparent() bool {
    return Image_GetTransparent(i.instance)
}

func (i *TImage) SetTransparent(value bool) {
    Image_SetTransparent(i.instance, value)
}

func (i *TImage) Visible() bool {
    return Image_GetVisible(i.instance)
}

func (i *TImage) SetVisible(value bool) {
    Image_SetVisible(i.instance, value)
}

func (i *TImage) SetOnClick(fn TNotifyEvent) {
    Image_SetOnClick(i.instance, fn)
}

func (i *TImage) SetOnDblClick(fn TNotifyEvent) {
    Image_SetOnDblClick(i.instance, fn)
}

func (i *TImage) SetOnMouseDown(fn TMouseEvent) {
    Image_SetOnMouseDown(i.instance, fn)
}

func (i *TImage) SetOnMouseEnter(fn TNotifyEvent) {
    Image_SetOnMouseEnter(i.instance, fn)
}

func (i *TImage) SetOnMouseLeave(fn TNotifyEvent) {
    Image_SetOnMouseLeave(i.instance, fn)
}

func (i *TImage) SetOnMouseMove(fn TMouseMoveEvent) {
    Image_SetOnMouseMove(i.instance, fn)
}

func (i *TImage) SetOnMouseUp(fn TMouseEvent) {
    Image_SetOnMouseUp(i.instance, fn)
}

func (i *TImage) Action() *TAction {
    return ActionFromInst(Image_GetAction(i.instance))
}

func (i *TImage) SetAction(value IComponent) {
    Image_SetAction(i.instance, CheckPtr(value))
}

func (i *TImage) BiDiMode() TBiDiMode {
    return Image_GetBiDiMode(i.instance)
}

func (i *TImage) SetBiDiMode(value TBiDiMode) {
    Image_SetBiDiMode(i.instance, value)
}

func (i *TImage) BoundsRect() TRect {
    return Image_GetBoundsRect(i.instance)
}

func (i *TImage) SetBoundsRect(value TRect) {
    Image_SetBoundsRect(i.instance, value)
}

func (i *TImage) ClientHeight() int32 {
    return Image_GetClientHeight(i.instance)
}

func (i *TImage) SetClientHeight(value int32) {
    Image_SetClientHeight(i.instance, value)
}

func (i *TImage) ClientRect() TRect {
    return Image_GetClientRect(i.instance)
}

func (i *TImage) ClientWidth() int32 {
    return Image_GetClientWidth(i.instance)
}

func (i *TImage) SetClientWidth(value int32) {
    Image_SetClientWidth(i.instance, value)
}

func (i *TImage) ExplicitLeft() int32 {
    return Image_GetExplicitLeft(i.instance)
}

func (i *TImage) ExplicitTop() int32 {
    return Image_GetExplicitTop(i.instance)
}

func (i *TImage) ExplicitWidth() int32 {
    return Image_GetExplicitWidth(i.instance)
}

func (i *TImage) ExplicitHeight() int32 {
    return Image_GetExplicitHeight(i.instance)
}

func (i *TImage) Parent() *TControl {
    return ControlFromInst(Image_GetParent(i.instance))
}

func (i *TImage) SetParent(value IControl) {
    Image_SetParent(i.instance, CheckPtr(value))
}

func (i *TImage) StyleElements() TStyleElements {
    return Image_GetStyleElements(i.instance)
}

func (i *TImage) SetStyleElements(value TStyleElements) {
    Image_SetStyleElements(i.instance, value)
}

func (i *TImage) AlignWithMargins() bool {
    return Image_GetAlignWithMargins(i.instance)
}

func (i *TImage) SetAlignWithMargins(value bool) {
    Image_SetAlignWithMargins(i.instance, value)
}

func (i *TImage) Left() int32 {
    return Image_GetLeft(i.instance)
}

func (i *TImage) SetLeft(value int32) {
    Image_SetLeft(i.instance, value)
}

func (i *TImage) Top() int32 {
    return Image_GetTop(i.instance)
}

func (i *TImage) SetTop(value int32) {
    Image_SetTop(i.instance, value)
}

func (i *TImage) Width() int32 {
    return Image_GetWidth(i.instance)
}

func (i *TImage) SetWidth(value int32) {
    Image_SetWidth(i.instance, value)
}

func (i *TImage) Height() int32 {
    return Image_GetHeight(i.instance)
}

func (i *TImage) SetHeight(value int32) {
    Image_SetHeight(i.instance, value)
}

func (i *TImage) Cursor() TCursor {
    return Image_GetCursor(i.instance)
}

func (i *TImage) SetCursor(value TCursor) {
    Image_SetCursor(i.instance, value)
}

func (i *TImage) Hint() string {
    return Image_GetHint(i.instance)
}

func (i *TImage) SetHint(value string) {
    Image_SetHint(i.instance, value)
}

func (i *TImage) Margins() *TMargins {
    return MarginsFromInst(Image_GetMargins(i.instance))
}

func (i *TImage) SetMargins(value *TMargins) {
    Image_SetMargins(i.instance, CheckPtr(value))
}

func (i *TImage) CustomHint() *TCustomHint {
    return CustomHintFromInst(Image_GetCustomHint(i.instance))
}

func (i *TImage) SetCustomHint(value IComponent) {
    Image_SetCustomHint(i.instance, CheckPtr(value))
}

func (i *TImage) ComponentCount() int32 {
    return Image_GetComponentCount(i.instance)
}

func (i *TImage) ComponentIndex() int32 {
    return Image_GetComponentIndex(i.instance)
}

func (i *TImage) SetComponentIndex(value int32) {
    Image_SetComponentIndex(i.instance, value)
}

func (i *TImage) Owner() *TComponent {
    return ComponentFromInst(Image_GetOwner(i.instance))
}

func (i *TImage) Name() string {
    return Image_GetName(i.instance)
}

func (i *TImage) SetName(value string) {
    Image_SetName(i.instance, value)
}

func (i *TImage) Tag() int {
    return Image_GetTag(i.instance)
}

func (i *TImage) SetTag(value int) {
    Image_SetTag(i.instance, value)
}

func (i *TImage) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Image_GetComponents(i.instance, AIndex))
}

