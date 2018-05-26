
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

type TMonitor struct {
    IObject
    instance uintptr
}

func NewMonitor() *TMonitor {
    m := new(TMonitor)
    m.instance = Monitor_Create()
    return m
}

func MonitorFromInst(inst uintptr) *TMonitor {
    m := new(TMonitor)
    m.instance = inst
    return m
}

func MonitorFromObj(obj IObject) *TMonitor {
    m := new(TMonitor)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMonitor) Free() {
    if m.instance != 0 {
        Monitor_Free(m.instance)
        m.instance = 0
    }
}

func (m *TMonitor) Instance() uintptr {
    return m.instance
}

func (m *TMonitor) IsValid() bool {
    return m.instance != 0
}

func (m *TMonitor) ClassName() string {
    return Monitor_ClassName(m.instance)
}

func (m *TMonitor) Equals(Obj IObject) bool {
    return Monitor_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMonitor) GetHashCode() int32 {
    return Monitor_GetHashCode(m.instance)
}

func (m *TMonitor) ToString() string {
    return Monitor_ToString(m.instance)
}

func (m *TMonitor) Handle() HMONITOR {
    return Monitor_GetHandle(m.instance)
}

func (m *TMonitor) MonitorNum() int32 {
    return Monitor_GetMonitorNum(m.instance)
}

func (m *TMonitor) Left() int32 {
    return Monitor_GetLeft(m.instance)
}

func (m *TMonitor) Height() int32 {
    return Monitor_GetHeight(m.instance)
}

func (m *TMonitor) Top() int32 {
    return Monitor_GetTop(m.instance)
}

func (m *TMonitor) Width() int32 {
    return Monitor_GetWidth(m.instance)
}

func (m *TMonitor) BoundsRect() TRect {
    return Monitor_GetBoundsRect(m.instance)
}

func (m *TMonitor) WorkareaRect() TRect {
    return Monitor_GetWorkareaRect(m.instance)
}

func (m *TMonitor) Primary() bool {
    return Monitor_GetPrimary(m.instance)
}

func (m *TMonitor) PixelsPerInch() int32 {
    return Monitor_GetPixelsPerInch(m.instance)
}

