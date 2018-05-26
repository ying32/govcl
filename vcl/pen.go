
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

type TPen struct {
    IObject
    instance uintptr
}

func NewPen() *TPen {
    p := new(TPen)
    p.instance = Pen_Create()
    return p
}

func PenFromInst(inst uintptr) *TPen {
    p := new(TPen)
    p.instance = inst
    return p
}

func PenFromObj(obj IObject) *TPen {
    p := new(TPen)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPen) Free() {
    if p.instance != 0 {
        Pen_Free(p.instance)
        p.instance = 0
    }
}

func (p *TPen) Instance() uintptr {
    return p.instance
}

func (p *TPen) IsValid() bool {
    return p.instance != 0
}

func (p *TPen) Assign(Source IObject) {
    Pen_Assign(p.instance, CheckPtr(Source))
}

func (p *TPen) HandleAllocated() bool {
    return Pen_HandleAllocated(p.instance)
}

func (p *TPen) GetNamePath() string {
    return Pen_GetNamePath(p.instance)
}

func (p *TPen) ClassName() string {
    return Pen_ClassName(p.instance)
}

func (p *TPen) Equals(Obj IObject) bool {
    return Pen_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPen) GetHashCode() int32 {
    return Pen_GetHashCode(p.instance)
}

func (p *TPen) ToString() string {
    return Pen_ToString(p.instance)
}

func (p *TPen) Handle() HPEN {
    return Pen_GetHandle(p.instance)
}

func (p *TPen) SetHandle(value HPEN) {
    Pen_SetHandle(p.instance, value)
}

func (p *TPen) Color() TColor {
    return Pen_GetColor(p.instance)
}

func (p *TPen) SetColor(value TColor) {
    Pen_SetColor(p.instance, value)
}

func (p *TPen) Mode() TPenMode {
    return Pen_GetMode(p.instance)
}

func (p *TPen) SetMode(value TPenMode) {
    Pen_SetMode(p.instance, value)
}

func (p *TPen) Style() TPenStyle {
    return Pen_GetStyle(p.instance)
}

func (p *TPen) SetStyle(value TPenStyle) {
    Pen_SetStyle(p.instance, value)
}

func (p *TPen) Width() int32 {
    return Pen_GetWidth(p.instance)
}

func (p *TPen) SetWidth(value int32) {
    Pen_SetWidth(p.instance, value)
}

func (p *TPen) SetOnChange(fn TNotifyEvent) {
    Pen_SetOnChange(p.instance, fn)
}

