
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

type TShape struct {
    IControl
    instance uintptr
}

func NewShape(owner IComponent) *TShape {
    s := new(TShape)
    s.instance = Shape_Create(CheckPtr(owner))
    return s
}

func ShapeFromInst(inst uintptr) *TShape {
    s := new(TShape)
    s.instance = inst
    return s
}

func ShapeFromObj(obj IObject) *TShape {
    s := new(TShape)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TShape) Free() {
    if s.instance != 0 {
        Shape_Free(s.instance)
        s.instance = 0
    }
}

func (s *TShape) Instance() uintptr {
    return s.instance
}

func (s *TShape) IsValid() bool {
    return s.instance != 0
}

func (s *TShape) BringToFront() {
    Shape_BringToFront(s.instance)
}

func (s *TShape) HasParent() bool {
    return Shape_HasParent(s.instance)
}

func (s *TShape) Hide() {
    Shape_Hide(s.instance)
}

func (s *TShape) Invalidate() {
    Shape_Invalidate(s.instance)
}

func (s *TShape) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Shape_Perform(s.instance, Msg , WParam , LParam)
}

func (s *TShape) Refresh() {
    Shape_Refresh(s.instance)
}

func (s *TShape) Repaint() {
    Shape_Repaint(s.instance)
}

func (s *TShape) SendToBack() {
    Shape_SendToBack(s.instance)
}

func (s *TShape) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Shape_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

func (s *TShape) Show() {
    Shape_Show(s.instance)
}

func (s *TShape) Update() {
    Shape_Update(s.instance)
}

func (s *TShape) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Shape_GetTextBuf(s.instance, Buffer , BufSize)
}

func (s *TShape) GetTextLen() int32 {
    return Shape_GetTextLen(s.instance)
}

func (s *TShape) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Shape_FindComponent(s.instance, AName))
}

func (s *TShape) GetNamePath() string {
    return Shape_GetNamePath(s.instance)
}

func (s *TShape) Assign(Source IObject) {
    Shape_Assign(s.instance, CheckPtr(Source))
}

func (s *TShape) ClassName() string {
    return Shape_ClassName(s.instance)
}

func (s *TShape) Equals(Obj IObject) bool {
    return Shape_Equals(s.instance, CheckPtr(Obj))
}

func (s *TShape) GetHashCode() int32 {
    return Shape_GetHashCode(s.instance)
}

func (s *TShape) ToString() string {
    return Shape_ToString(s.instance)
}

func (s *TShape) Align() TAlign {
    return Shape_GetAlign(s.instance)
}

func (s *TShape) SetAlign(value TAlign) {
    Shape_SetAlign(s.instance, value)
}

func (s *TShape) Anchors() TAnchors {
    return Shape_GetAnchors(s.instance)
}

func (s *TShape) SetAnchors(value TAnchors) {
    Shape_SetAnchors(s.instance, value)
}

func (s *TShape) Brush() *TBrush {
    return BrushFromInst(Shape_GetBrush(s.instance))
}

func (s *TShape) SetBrush(value *TBrush) {
    Shape_SetBrush(s.instance, CheckPtr(value))
}

func (s *TShape) Enabled() bool {
    return Shape_GetEnabled(s.instance)
}

func (s *TShape) SetEnabled(value bool) {
    Shape_SetEnabled(s.instance, value)
}

func (s *TShape) ParentShowHint() bool {
    return Shape_GetParentShowHint(s.instance)
}

func (s *TShape) SetParentShowHint(value bool) {
    Shape_SetParentShowHint(s.instance, value)
}

func (s *TShape) Pen() *TPen {
    return PenFromInst(Shape_GetPen(s.instance))
}

func (s *TShape) SetPen(value *TPen) {
    Shape_SetPen(s.instance, CheckPtr(value))
}

func (s *TShape) Shape() TShapeType {
    return Shape_GetShape(s.instance)
}

func (s *TShape) SetShape(value TShapeType) {
    Shape_SetShape(s.instance, value)
}

func (s *TShape) ShowHint() bool {
    return Shape_GetShowHint(s.instance)
}

func (s *TShape) SetShowHint(value bool) {
    Shape_SetShowHint(s.instance, value)
}

func (s *TShape) Visible() bool {
    return Shape_GetVisible(s.instance)
}

func (s *TShape) SetVisible(value bool) {
    Shape_SetVisible(s.instance, value)
}

func (s *TShape) SetOnMouseDown(fn TMouseEvent) {
    Shape_SetOnMouseDown(s.instance, fn)
}

func (s *TShape) SetOnMouseEnter(fn TNotifyEvent) {
    Shape_SetOnMouseEnter(s.instance, fn)
}

func (s *TShape) SetOnMouseLeave(fn TNotifyEvent) {
    Shape_SetOnMouseLeave(s.instance, fn)
}

func (s *TShape) SetOnMouseMove(fn TMouseMoveEvent) {
    Shape_SetOnMouseMove(s.instance, fn)
}

func (s *TShape) SetOnMouseUp(fn TMouseEvent) {
    Shape_SetOnMouseUp(s.instance, fn)
}

func (s *TShape) Action() *TAction {
    return ActionFromInst(Shape_GetAction(s.instance))
}

func (s *TShape) SetAction(value IComponent) {
    Shape_SetAction(s.instance, CheckPtr(value))
}

func (s *TShape) BiDiMode() TBiDiMode {
    return Shape_GetBiDiMode(s.instance)
}

func (s *TShape) SetBiDiMode(value TBiDiMode) {
    Shape_SetBiDiMode(s.instance, value)
}

func (s *TShape) BoundsRect() TRect {
    return Shape_GetBoundsRect(s.instance)
}

func (s *TShape) SetBoundsRect(value TRect) {
    Shape_SetBoundsRect(s.instance, value)
}

func (s *TShape) ClientHeight() int32 {
    return Shape_GetClientHeight(s.instance)
}

func (s *TShape) SetClientHeight(value int32) {
    Shape_SetClientHeight(s.instance, value)
}

func (s *TShape) ClientRect() TRect {
    return Shape_GetClientRect(s.instance)
}

func (s *TShape) ClientWidth() int32 {
    return Shape_GetClientWidth(s.instance)
}

func (s *TShape) SetClientWidth(value int32) {
    Shape_SetClientWidth(s.instance, value)
}

func (s *TShape) ExplicitLeft() int32 {
    return Shape_GetExplicitLeft(s.instance)
}

func (s *TShape) ExplicitTop() int32 {
    return Shape_GetExplicitTop(s.instance)
}

func (s *TShape) ExplicitWidth() int32 {
    return Shape_GetExplicitWidth(s.instance)
}

func (s *TShape) ExplicitHeight() int32 {
    return Shape_GetExplicitHeight(s.instance)
}

func (s *TShape) Parent() *TControl {
    return ControlFromInst(Shape_GetParent(s.instance))
}

func (s *TShape) SetParent(value IControl) {
    Shape_SetParent(s.instance, CheckPtr(value))
}

func (s *TShape) StyleElements() TStyleElements {
    return Shape_GetStyleElements(s.instance)
}

func (s *TShape) SetStyleElements(value TStyleElements) {
    Shape_SetStyleElements(s.instance, value)
}

func (s *TShape) AlignWithMargins() bool {
    return Shape_GetAlignWithMargins(s.instance)
}

func (s *TShape) SetAlignWithMargins(value bool) {
    Shape_SetAlignWithMargins(s.instance, value)
}

func (s *TShape) Left() int32 {
    return Shape_GetLeft(s.instance)
}

func (s *TShape) SetLeft(value int32) {
    Shape_SetLeft(s.instance, value)
}

func (s *TShape) Top() int32 {
    return Shape_GetTop(s.instance)
}

func (s *TShape) SetTop(value int32) {
    Shape_SetTop(s.instance, value)
}

func (s *TShape) Width() int32 {
    return Shape_GetWidth(s.instance)
}

func (s *TShape) SetWidth(value int32) {
    Shape_SetWidth(s.instance, value)
}

func (s *TShape) Height() int32 {
    return Shape_GetHeight(s.instance)
}

func (s *TShape) SetHeight(value int32) {
    Shape_SetHeight(s.instance, value)
}

func (s *TShape) Cursor() TCursor {
    return Shape_GetCursor(s.instance)
}

func (s *TShape) SetCursor(value TCursor) {
    Shape_SetCursor(s.instance, value)
}

func (s *TShape) Hint() string {
    return Shape_GetHint(s.instance)
}

func (s *TShape) SetHint(value string) {
    Shape_SetHint(s.instance, value)
}

func (s *TShape) Margins() *TMargins {
    return MarginsFromInst(Shape_GetMargins(s.instance))
}

func (s *TShape) SetMargins(value *TMargins) {
    Shape_SetMargins(s.instance, CheckPtr(value))
}

func (s *TShape) CustomHint() *TCustomHint {
    return CustomHintFromInst(Shape_GetCustomHint(s.instance))
}

func (s *TShape) SetCustomHint(value IComponent) {
    Shape_SetCustomHint(s.instance, CheckPtr(value))
}

func (s *TShape) ComponentCount() int32 {
    return Shape_GetComponentCount(s.instance)
}

func (s *TShape) ComponentIndex() int32 {
    return Shape_GetComponentIndex(s.instance)
}

func (s *TShape) SetComponentIndex(value int32) {
    Shape_SetComponentIndex(s.instance, value)
}

func (s *TShape) Owner() *TComponent {
    return ComponentFromInst(Shape_GetOwner(s.instance))
}

func (s *TShape) Name() string {
    return Shape_GetName(s.instance)
}

func (s *TShape) SetName(value string) {
    Shape_SetName(s.instance, value)
}

func (s *TShape) Tag() int {
    return Shape_GetTag(s.instance)
}

func (s *TShape) SetTag(value int) {
    Shape_SetTag(s.instance, value)
}

func (s *TShape) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Shape_GetComponents(s.instance, AIndex))
}

