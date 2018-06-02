
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
)

type TDragDockObject struct {
    IObject
    instance uintptr
}

func NewDragDockObject() *TDragDockObject {
    d := new(TDragDockObject)
    d.instance = DragDockObject_Create()
    return d
}

func DragDockObjectFromInst(inst uintptr) *TDragDockObject {
    d := new(TDragDockObject)
    d.instance = inst
    return d
}

func DragDockObjectFromObj(obj IObject) *TDragDockObject {
    d := new(TDragDockObject)
    d.instance = CheckPtr(obj)
    return d
}

func (d *TDragDockObject) Free() {
    if d.instance != 0 {
        DragDockObject_Free(d.instance)
        d.instance = 0
    }
}

func (d *TDragDockObject) Instance() uintptr {
    return d.instance
}

func (d *TDragDockObject) IsValid() bool {
    return d.instance != 0
}

func TDragDockObjectClass() TClass {
    return DragDockObject_StaticClassType()
}

func (d *TDragDockObject) Assign(Source *TDragObject) {
    DragDockObject_Assign(d.instance, CheckPtr(Source))
}

func (d *TDragDockObject) HideDragImage() {
    DragDockObject_HideDragImage(d.instance)
}

func (d *TDragDockObject) ShowDragImage() {
    DragDockObject_ShowDragImage(d.instance)
}

func (d *TDragDockObject) DisposeOf() {
    DragDockObject_DisposeOf(d.instance)
}

func (d *TDragDockObject) ClassType() TClass {
    return DragDockObject_ClassType(d.instance)
}

func (d *TDragDockObject) ClassName() string {
    return DragDockObject_ClassName(d.instance)
}

func (d *TDragDockObject) InstanceSize() int32 {
    return DragDockObject_InstanceSize(d.instance)
}

func (d *TDragDockObject) InheritsFrom(AClass TClass) bool {
    return DragDockObject_InheritsFrom(d.instance, AClass)
}

func (d *TDragDockObject) Equals(Obj IObject) bool {
    return DragDockObject_Equals(d.instance, CheckPtr(Obj))
}

func (d *TDragDockObject) GetHashCode() int32 {
    return DragDockObject_GetHashCode(d.instance)
}

func (d *TDragDockObject) ToString() string {
    return DragDockObject_ToString(d.instance)
}

func (d *TDragDockObject) Brush() *TBrush {
    return BrushFromInst(DragDockObject_GetBrush(d.instance))
}

func (d *TDragDockObject) SetBrush(value *TBrush) {
    DragDockObject_SetBrush(d.instance, CheckPtr(value))
}

func (d *TDragDockObject) DockRect() TRect {
    return DragDockObject_GetDockRect(d.instance)
}

func (d *TDragDockObject) SetDockRect(value TRect) {
    DragDockObject_SetDockRect(d.instance, value)
}

func (d *TDragDockObject) DropAlign() TAlign {
    return DragDockObject_GetDropAlign(d.instance)
}

func (d *TDragDockObject) DropOnControl() *TControl {
    return ControlFromInst(DragDockObject_GetDropOnControl(d.instance))
}

func (d *TDragDockObject) EraseDockRect() TRect {
    return DragDockObject_GetEraseDockRect(d.instance)
}

func (d *TDragDockObject) SetEraseDockRect(value TRect) {
    DragDockObject_SetEraseDockRect(d.instance, value)
}

func (d *TDragDockObject) EraseWhenMoving() bool {
    return DragDockObject_GetEraseWhenMoving(d.instance)
}

func (d *TDragDockObject) Floating() bool {
    return DragDockObject_GetFloating(d.instance)
}

func (d *TDragDockObject) SetFloating(value bool) {
    DragDockObject_SetFloating(d.instance, value)
}

func (d *TDragDockObject) FrameWidth() int32 {
    return DragDockObject_GetFrameWidth(d.instance)
}

func (d *TDragDockObject) Control() *TControl {
    return ControlFromInst(DragDockObject_GetControl(d.instance))
}

func (d *TDragDockObject) SetControl(value IControl) {
    DragDockObject_SetControl(d.instance, CheckPtr(value))
}

func (d *TDragDockObject) AlwaysShowDragImages() bool {
    return DragDockObject_GetAlwaysShowDragImages(d.instance)
}

func (d *TDragDockObject) SetAlwaysShowDragImages(value bool) {
    DragDockObject_SetAlwaysShowDragImages(d.instance, value)
}

func (d *TDragDockObject) Cancelling() bool {
    return DragDockObject_GetCancelling(d.instance)
}

func (d *TDragDockObject) SetCancelling(value bool) {
    DragDockObject_SetCancelling(d.instance, value)
}

func (d *TDragDockObject) DragHandle() HWND {
    return DragDockObject_GetDragHandle(d.instance)
}

func (d *TDragDockObject) SetDragHandle(value HWND) {
    DragDockObject_SetDragHandle(d.instance, value)
}

func (d *TDragDockObject) DragPos() TPoint {
    return DragDockObject_GetDragPos(d.instance)
}

func (d *TDragDockObject) SetDragPos(value TPoint) {
    DragDockObject_SetDragPos(d.instance, value)
}

func (d *TDragDockObject) DragTarget() uintptr {
    return DragDockObject_GetDragTarget(d.instance)
}

func (d *TDragDockObject) SetDragTarget(value uintptr) {
    DragDockObject_SetDragTarget(d.instance, value)
}

func (d *TDragDockObject) DragTargetPos() TPoint {
    return DragDockObject_GetDragTargetPos(d.instance)
}

func (d *TDragDockObject) SetDragTargetPos(value TPoint) {
    DragDockObject_SetDragTargetPos(d.instance, value)
}

func (d *TDragDockObject) Dropped() bool {
    return DragDockObject_GetDropped(d.instance)
}

func (d *TDragDockObject) MouseDeltaX() float64 {
    return DragDockObject_GetMouseDeltaX(d.instance)
}

func (d *TDragDockObject) MouseDeltaY() float64 {
    return DragDockObject_GetMouseDeltaY(d.instance)
}

func (d *TDragDockObject) RightClickCancels() bool {
    return DragDockObject_GetRightClickCancels(d.instance)
}

func (d *TDragDockObject) SetRightClickCancels(value bool) {
    DragDockObject_SetRightClickCancels(d.instance, value)
}

