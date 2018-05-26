
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TMargins struct {
    IObject
    instance uintptr
}

func NewMargins() *TMargins {
    m := new(TMargins)
    m.instance = Margins_Create()
    return m
}

func MarginsFromInst(inst uintptr) *TMargins {
    m := new(TMargins)
    m.instance = inst
    return m
}

func MarginsFromObj(obj IObject) *TMargins {
    m := new(TMargins)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMargins) Free() {
    if m.instance != 0 {
        Margins_Free(m.instance)
        m.instance = 0
    }
}

func (m *TMargins) Instance() uintptr {
    return m.instance
}

func (m *TMargins) IsValid() bool {
    return m.instance != 0
}

func (m *TMargins) SetBounds(ALeft int32, ATop int32, ARight int32, ABottom int32) {
    Margins_SetBounds(m.instance, ALeft , ATop , ARight , ABottom)
}

func (m *TMargins) Assign(Source IObject) {
    Margins_Assign(m.instance, CheckPtr(Source))
}

func (m *TMargins) GetNamePath() string {
    return Margins_GetNamePath(m.instance)
}

func (m *TMargins) ClassName() string {
    return Margins_ClassName(m.instance)
}

func (m *TMargins) Equals(Obj IObject) bool {
    return Margins_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMargins) GetHashCode() int32 {
    return Margins_GetHashCode(m.instance)
}

func (m *TMargins) ToString() string {
    return Margins_ToString(m.instance)
}

func (m *TMargins) ControlLeft() int32 {
    return Margins_GetControlLeft(m.instance)
}

func (m *TMargins) ControlTop() int32 {
    return Margins_GetControlTop(m.instance)
}

func (m *TMargins) ControlWidth() int32 {
    return Margins_GetControlWidth(m.instance)
}

func (m *TMargins) ControlHeight() int32 {
    return Margins_GetControlHeight(m.instance)
}

func (m *TMargins) ExplicitLeft() int32 {
    return Margins_GetExplicitLeft(m.instance)
}

func (m *TMargins) ExplicitTop() int32 {
    return Margins_GetExplicitTop(m.instance)
}

func (m *TMargins) ExplicitWidth() int32 {
    return Margins_GetExplicitWidth(m.instance)
}

func (m *TMargins) ExplicitHeight() int32 {
    return Margins_GetExplicitHeight(m.instance)
}

func (m *TMargins) SetOnChange(fn TNotifyEvent) {
    Margins_SetOnChange(m.instance, fn)
}

func (m *TMargins) Left() int32 {
    return Margins_GetLeft(m.instance)
}

func (m *TMargins) SetLeft(value int32) {
    Margins_SetLeft(m.instance, value)
}

func (m *TMargins) Top() int32 {
    return Margins_GetTop(m.instance)
}

func (m *TMargins) SetTop(value int32) {
    Margins_SetTop(m.instance, value)
}

func (m *TMargins) Right() int32 {
    return Margins_GetRight(m.instance)
}

func (m *TMargins) SetRight(value int32) {
    Margins_SetRight(m.instance, value)
}

func (m *TMargins) Bottom() int32 {
    return Margins_GetBottom(m.instance)
}

func (m *TMargins) SetBottom(value int32) {
    Margins_SetBottom(m.instance, value)
}

