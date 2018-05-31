
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

type TDragObject struct {
    IObject
    instance uintptr
}

func NewDragObject() *TDragObject {
    d := new(TDragObject)
    d.instance = DragObject_Create()
    return d
}

func DragObjectFromInst(inst uintptr) *TDragObject {
    d := new(TDragObject)
    d.instance = inst
    return d
}

func DragObjectFromObj(obj IObject) *TDragObject {
    d := new(TDragObject)
    d.instance = CheckPtr(obj)
    return d
}

func (d *TDragObject) Free() {
    if d.instance != 0 {
        DragObject_Free(d.instance)
        d.instance = 0
    }
}

func (d *TDragObject) Instance() uintptr {
    return d.instance
}

func (d *TDragObject) IsValid() bool {
    return d.instance != 0
}

func (d *TDragObject) Assign(Source *TDragObject) {
    DragObject_Assign(d.instance, CheckPtr(Source))
}

func (d *TDragObject) HideDragImage() {
    DragObject_HideDragImage(d.instance)
}

func (d *TDragObject) ShowDragImage() {
    DragObject_ShowDragImage(d.instance)
}

func (d *TDragObject) ClassName() string {
    return DragObject_ClassName(d.instance)
}

func (d *TDragObject) Equals(Obj IObject) bool {
    return DragObject_Equals(d.instance, CheckPtr(Obj))
}

func (d *TDragObject) GetHashCode() int32 {
    return DragObject_GetHashCode(d.instance)
}

func (d *TDragObject) ToString() string {
    return DragObject_ToString(d.instance)
}

func (d *TDragObject) AlwaysShowDragImages() bool {
    return DragObject_GetAlwaysShowDragImages(d.instance)
}

func (d *TDragObject) SetAlwaysShowDragImages(value bool) {
    DragObject_SetAlwaysShowDragImages(d.instance, value)
}

func (d *TDragObject) Cancelling() bool {
    return DragObject_GetCancelling(d.instance)
}

func (d *TDragObject) SetCancelling(value bool) {
    DragObject_SetCancelling(d.instance, value)
}

func (d *TDragObject) DragHandle() HWND {
    return DragObject_GetDragHandle(d.instance)
}

func (d *TDragObject) SetDragHandle(value HWND) {
    DragObject_SetDragHandle(d.instance, value)
}

func (d *TDragObject) DragPos() TPoint {
    return DragObject_GetDragPos(d.instance)
}

func (d *TDragObject) SetDragPos(value TPoint) {
    DragObject_SetDragPos(d.instance, value)
}

func (d *TDragObject) DragTarget() uintptr {
    return DragObject_GetDragTarget(d.instance)
}

func (d *TDragObject) SetDragTarget(value uintptr) {
    DragObject_SetDragTarget(d.instance, value)
}

func (d *TDragObject) DragTargetPos() TPoint {
    return DragObject_GetDragTargetPos(d.instance)
}

func (d *TDragObject) SetDragTargetPos(value TPoint) {
    DragObject_SetDragTargetPos(d.instance, value)
}

func (d *TDragObject) Dropped() bool {
    return DragObject_GetDropped(d.instance)
}

func (d *TDragObject) MouseDeltaX() float64 {
    return DragObject_GetMouseDeltaX(d.instance)
}

func (d *TDragObject) MouseDeltaY() float64 {
    return DragObject_GetMouseDeltaY(d.instance)
}

func (d *TDragObject) RightClickCancels() bool {
    return DragObject_GetRightClickCancels(d.instance)
}

func (d *TDragObject) SetRightClickCancels(value bool) {
    DragObject_SetRightClickCancels(d.instance, value)
}

