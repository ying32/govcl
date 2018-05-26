
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

type TImageButton struct {
    IControl
    instance uintptr
}

func NewImageButton(owner IComponent) *TImageButton {
    i := new(TImageButton)
    i.instance = ImageButton_Create(CheckPtr(owner))
    return i
}

func ImageButtonFromInst(inst uintptr) *TImageButton {
    i := new(TImageButton)
    i.instance = inst
    return i
}

func ImageButtonFromObj(obj IObject) *TImageButton {
    i := new(TImageButton)
    i.instance = CheckPtr(obj)
    return i
}

func (i *TImageButton) Free() {
    if i.instance != 0 {
        ImageButton_Free(i.instance)
        i.instance = 0
    }
}

func (i *TImageButton) Instance() uintptr {
    return i.instance
}

func (i *TImageButton) IsValid() bool {
    return i.instance != 0
}

func (i *TImageButton) Click() {
    ImageButton_Click(i.instance)
}

func (i *TImageButton) BringToFront() {
    ImageButton_BringToFront(i.instance)
}

func (i *TImageButton) HasParent() bool {
    return ImageButton_HasParent(i.instance)
}

func (i *TImageButton) Hide() {
    ImageButton_Hide(i.instance)
}

func (i *TImageButton) Invalidate() {
    ImageButton_Invalidate(i.instance)
}

func (i *TImageButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ImageButton_Perform(i.instance, Msg , WParam , LParam)
}

func (i *TImageButton) Refresh() {
    ImageButton_Refresh(i.instance)
}

func (i *TImageButton) Repaint() {
    ImageButton_Repaint(i.instance)
}

func (i *TImageButton) SendToBack() {
    ImageButton_SendToBack(i.instance)
}

func (i *TImageButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ImageButton_SetBounds(i.instance, ALeft , ATop , AWidth , AHeight)
}

func (i *TImageButton) Show() {
    ImageButton_Show(i.instance)
}

func (i *TImageButton) Update() {
    ImageButton_Update(i.instance)
}

func (i *TImageButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ImageButton_GetTextBuf(i.instance, Buffer , BufSize)
}

func (i *TImageButton) GetTextLen() int32 {
    return ImageButton_GetTextLen(i.instance)
}

func (i *TImageButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ImageButton_FindComponent(i.instance, AName))
}

func (i *TImageButton) GetNamePath() string {
    return ImageButton_GetNamePath(i.instance)
}

func (i *TImageButton) Assign(Source IObject) {
    ImageButton_Assign(i.instance, CheckPtr(Source))
}

func (i *TImageButton) ClassName() string {
    return ImageButton_ClassName(i.instance)
}

func (i *TImageButton) Equals(Obj IObject) bool {
    return ImageButton_Equals(i.instance, CheckPtr(Obj))
}

func (i *TImageButton) GetHashCode() int32 {
    return ImageButton_GetHashCode(i.instance)
}

func (i *TImageButton) ToString() string {
    return ImageButton_ToString(i.instance)
}

func (i *TImageButton) Action() *TAction {
    return ActionFromInst(ImageButton_GetAction(i.instance))
}

func (i *TImageButton) SetAction(value IComponent) {
    ImageButton_SetAction(i.instance, CheckPtr(value))
}

func (i *TImageButton) Align() TAlign {
    return ImageButton_GetAlign(i.instance)
}

func (i *TImageButton) SetAlign(value TAlign) {
    ImageButton_SetAlign(i.instance, value)
}

func (i *TImageButton) Anchors() TAnchors {
    return ImageButton_GetAnchors(i.instance)
}

func (i *TImageButton) SetAnchors(value TAnchors) {
    ImageButton_SetAnchors(i.instance, value)
}

func (i *TImageButton) AutoSize() bool {
    return ImageButton_GetAutoSize(i.instance)
}

func (i *TImageButton) SetAutoSize(value bool) {
    ImageButton_SetAutoSize(i.instance, value)
}

func (i *TImageButton) Caption() string {
    return ImageButton_GetCaption(i.instance)
}

func (i *TImageButton) SetCaption(value string) {
    ImageButton_SetCaption(i.instance, value)
}

func (i *TImageButton) Enabled() bool {
    return ImageButton_GetEnabled(i.instance)
}

func (i *TImageButton) SetEnabled(value bool) {
    ImageButton_SetEnabled(i.instance, value)
}

func (i *TImageButton) Font() *TFont {
    return FontFromInst(ImageButton_GetFont(i.instance))
}

func (i *TImageButton) SetFont(value *TFont) {
    ImageButton_SetFont(i.instance, CheckPtr(value))
}

func (i *TImageButton) ImageCount() int32 {
    return ImageButton_GetImageCount(i.instance)
}

func (i *TImageButton) SetImageCount(value int32) {
    ImageButton_SetImageCount(i.instance, value)
}

func (i *TImageButton) ModalResult() TModalResult {
    return ImageButton_GetModalResult(i.instance)
}

func (i *TImageButton) SetModalResult(value TModalResult) {
    ImageButton_SetModalResult(i.instance, value)
}

func (i *TImageButton) ParentShowHint() bool {
    return ImageButton_GetParentShowHint(i.instance)
}

func (i *TImageButton) SetParentShowHint(value bool) {
    ImageButton_SetParentShowHint(i.instance, value)
}

func (i *TImageButton) ParentFont() bool {
    return ImageButton_GetParentFont(i.instance)
}

func (i *TImageButton) SetParentFont(value bool) {
    ImageButton_SetParentFont(i.instance, value)
}

func (i *TImageButton) Picture() *TPicture {
    return PictureFromInst(ImageButton_GetPicture(i.instance))
}

func (i *TImageButton) SetPicture(value *TPicture) {
    ImageButton_SetPicture(i.instance, CheckPtr(value))
}

func (i *TImageButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ImageButton_GetPopupMenu(i.instance))
}

func (i *TImageButton) SetPopupMenu(value IComponent) {
    ImageButton_SetPopupMenu(i.instance, CheckPtr(value))
}

func (i *TImageButton) ShowHint() bool {
    return ImageButton_GetShowHint(i.instance)
}

func (i *TImageButton) SetShowHint(value bool) {
    ImageButton_SetShowHint(i.instance, value)
}

func (i *TImageButton) ShowCaption() bool {
    return ImageButton_GetShowCaption(i.instance)
}

func (i *TImageButton) SetShowCaption(value bool) {
    ImageButton_SetShowCaption(i.instance, value)
}

func (i *TImageButton) Visible() bool {
    return ImageButton_GetVisible(i.instance)
}

func (i *TImageButton) SetVisible(value bool) {
    ImageButton_SetVisible(i.instance, value)
}

func (i *TImageButton) SetOnClick(fn TNotifyEvent) {
    ImageButton_SetOnClick(i.instance, fn)
}

func (i *TImageButton) SetOnDblClick(fn TNotifyEvent) {
    ImageButton_SetOnDblClick(i.instance, fn)
}

func (i *TImageButton) SetOnMouseDown(fn TMouseEvent) {
    ImageButton_SetOnMouseDown(i.instance, fn)
}

func (i *TImageButton) SetOnMouseEnter(fn TNotifyEvent) {
    ImageButton_SetOnMouseEnter(i.instance, fn)
}

func (i *TImageButton) SetOnMouseLeave(fn TNotifyEvent) {
    ImageButton_SetOnMouseLeave(i.instance, fn)
}

func (i *TImageButton) SetOnMouseMove(fn TMouseMoveEvent) {
    ImageButton_SetOnMouseMove(i.instance, fn)
}

func (i *TImageButton) SetOnMouseUp(fn TMouseEvent) {
    ImageButton_SetOnMouseUp(i.instance, fn)
}

func (i *TImageButton) BiDiMode() TBiDiMode {
    return ImageButton_GetBiDiMode(i.instance)
}

func (i *TImageButton) SetBiDiMode(value TBiDiMode) {
    ImageButton_SetBiDiMode(i.instance, value)
}

func (i *TImageButton) BoundsRect() TRect {
    return ImageButton_GetBoundsRect(i.instance)
}

func (i *TImageButton) SetBoundsRect(value TRect) {
    ImageButton_SetBoundsRect(i.instance, value)
}

func (i *TImageButton) ClientHeight() int32 {
    return ImageButton_GetClientHeight(i.instance)
}

func (i *TImageButton) SetClientHeight(value int32) {
    ImageButton_SetClientHeight(i.instance, value)
}

func (i *TImageButton) ClientRect() TRect {
    return ImageButton_GetClientRect(i.instance)
}

func (i *TImageButton) ClientWidth() int32 {
    return ImageButton_GetClientWidth(i.instance)
}

func (i *TImageButton) SetClientWidth(value int32) {
    ImageButton_SetClientWidth(i.instance, value)
}

func (i *TImageButton) ExplicitLeft() int32 {
    return ImageButton_GetExplicitLeft(i.instance)
}

func (i *TImageButton) ExplicitTop() int32 {
    return ImageButton_GetExplicitTop(i.instance)
}

func (i *TImageButton) ExplicitWidth() int32 {
    return ImageButton_GetExplicitWidth(i.instance)
}

func (i *TImageButton) ExplicitHeight() int32 {
    return ImageButton_GetExplicitHeight(i.instance)
}

func (i *TImageButton) Parent() *TControl {
    return ControlFromInst(ImageButton_GetParent(i.instance))
}

func (i *TImageButton) SetParent(value IControl) {
    ImageButton_SetParent(i.instance, CheckPtr(value))
}

func (i *TImageButton) StyleElements() TStyleElements {
    return ImageButton_GetStyleElements(i.instance)
}

func (i *TImageButton) SetStyleElements(value TStyleElements) {
    ImageButton_SetStyleElements(i.instance, value)
}

func (i *TImageButton) AlignWithMargins() bool {
    return ImageButton_GetAlignWithMargins(i.instance)
}

func (i *TImageButton) SetAlignWithMargins(value bool) {
    ImageButton_SetAlignWithMargins(i.instance, value)
}

func (i *TImageButton) Left() int32 {
    return ImageButton_GetLeft(i.instance)
}

func (i *TImageButton) SetLeft(value int32) {
    ImageButton_SetLeft(i.instance, value)
}

func (i *TImageButton) Top() int32 {
    return ImageButton_GetTop(i.instance)
}

func (i *TImageButton) SetTop(value int32) {
    ImageButton_SetTop(i.instance, value)
}

func (i *TImageButton) Width() int32 {
    return ImageButton_GetWidth(i.instance)
}

func (i *TImageButton) SetWidth(value int32) {
    ImageButton_SetWidth(i.instance, value)
}

func (i *TImageButton) Height() int32 {
    return ImageButton_GetHeight(i.instance)
}

func (i *TImageButton) SetHeight(value int32) {
    ImageButton_SetHeight(i.instance, value)
}

func (i *TImageButton) Cursor() TCursor {
    return ImageButton_GetCursor(i.instance)
}

func (i *TImageButton) SetCursor(value TCursor) {
    ImageButton_SetCursor(i.instance, value)
}

func (i *TImageButton) Hint() string {
    return ImageButton_GetHint(i.instance)
}

func (i *TImageButton) SetHint(value string) {
    ImageButton_SetHint(i.instance, value)
}

func (i *TImageButton) Margins() *TMargins {
    return MarginsFromInst(ImageButton_GetMargins(i.instance))
}

func (i *TImageButton) SetMargins(value *TMargins) {
    ImageButton_SetMargins(i.instance, CheckPtr(value))
}

func (i *TImageButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(ImageButton_GetCustomHint(i.instance))
}

func (i *TImageButton) SetCustomHint(value IComponent) {
    ImageButton_SetCustomHint(i.instance, CheckPtr(value))
}

func (i *TImageButton) ComponentCount() int32 {
    return ImageButton_GetComponentCount(i.instance)
}

func (i *TImageButton) ComponentIndex() int32 {
    return ImageButton_GetComponentIndex(i.instance)
}

func (i *TImageButton) SetComponentIndex(value int32) {
    ImageButton_SetComponentIndex(i.instance, value)
}

func (i *TImageButton) Owner() *TComponent {
    return ComponentFromInst(ImageButton_GetOwner(i.instance))
}

func (i *TImageButton) Name() string {
    return ImageButton_GetName(i.instance)
}

func (i *TImageButton) SetName(value string) {
    ImageButton_SetName(i.instance, value)
}

func (i *TImageButton) Tag() int {
    return ImageButton_GetTag(i.instance)
}

func (i *TImageButton) SetTag(value int) {
    ImageButton_SetTag(i.instance, value)
}

func (i *TImageButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ImageButton_GetComponents(i.instance, AIndex))
}

