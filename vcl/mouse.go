
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

type TMouse struct {
    IObject
    instance uintptr
}

func NewMouse() *TMouse {
    m := new(TMouse)
    m.instance = Mouse_Create()
    return m
}

func MouseFromInst(inst uintptr) *TMouse {
    m := new(TMouse)
    m.instance = inst
    return m
}

func MouseFromObj(obj IObject) *TMouse {
    m := new(TMouse)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMouse) Free() {
    if m.instance != 0 {
        Mouse_Free(m.instance)
        m.instance = 0
    }
}

func (m *TMouse) Instance() uintptr {
    return m.instance
}

func (m *TMouse) IsValid() bool {
    return m.instance != 0
}

func (m *TMouse) ClassName() string {
    return Mouse_ClassName(m.instance)
}

func (m *TMouse) Equals(Obj IObject) bool {
    return Mouse_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMouse) GetHashCode() int32 {
    return Mouse_GetHashCode(m.instance)
}

func (m *TMouse) ToString() string {
    return Mouse_ToString(m.instance)
}

func (m *TMouse) Capture() HWND {
    return Mouse_GetCapture(m.instance)
}

func (m *TMouse) SetCapture(value HWND) {
    Mouse_SetCapture(m.instance, value)
}

func (m *TMouse) CursorPos() TPoint {
    return Mouse_GetCursorPos(m.instance)
}

func (m *TMouse) SetCursorPos(value TPoint) {
    Mouse_SetCursorPos(m.instance, value)
}

func (m *TMouse) IsDragging() bool {
    return Mouse_GetIsDragging(m.instance)
}

func (m *TMouse) IsPanning() bool {
    return Mouse_GetIsPanning(m.instance)
}

func (m *TMouse) WheelPresent() bool {
    return Mouse_GetWheelPresent(m.instance)
}

func (m *TMouse) WheelScrollLines() int32 {
    return Mouse_GetWheelScrollLines(m.instance)
}

