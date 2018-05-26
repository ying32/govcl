
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

type TMonthCalColors struct {
    IObject
    instance uintptr
}

func NewMonthCalColors() *TMonthCalColors {
    m := new(TMonthCalColors)
    m.instance = MonthCalColors_Create()
    return m
}

func MonthCalColorsFromInst(inst uintptr) *TMonthCalColors {
    m := new(TMonthCalColors)
    m.instance = inst
    return m
}

func MonthCalColorsFromObj(obj IObject) *TMonthCalColors {
    m := new(TMonthCalColors)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMonthCalColors) Free() {
    if m.instance != 0 {
        MonthCalColors_Free(m.instance)
        m.instance = 0
    }
}

func (m *TMonthCalColors) Instance() uintptr {
    return m.instance
}

func (m *TMonthCalColors) IsValid() bool {
    return m.instance != 0
}

func (m *TMonthCalColors) Assign(Source IObject) {
    MonthCalColors_Assign(m.instance, CheckPtr(Source))
}

func (m *TMonthCalColors) GetNamePath() string {
    return MonthCalColors_GetNamePath(m.instance)
}

func (m *TMonthCalColors) ClassName() string {
    return MonthCalColors_ClassName(m.instance)
}

func (m *TMonthCalColors) Equals(Obj IObject) bool {
    return MonthCalColors_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMonthCalColors) GetHashCode() int32 {
    return MonthCalColors_GetHashCode(m.instance)
}

func (m *TMonthCalColors) ToString() string {
    return MonthCalColors_ToString(m.instance)
}

func (m *TMonthCalColors) BackColor() TColor {
    return MonthCalColors_GetBackColor(m.instance)
}

func (m *TMonthCalColors) SetBackColor(value TColor) {
    MonthCalColors_SetBackColor(m.instance, value)
}

