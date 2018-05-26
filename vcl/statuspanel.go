
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

type TStatusPanel struct {
    IObject
    instance uintptr
}

func NewStatusPanel() *TStatusPanel {
    s := new(TStatusPanel)
    s.instance = StatusPanel_Create()
    return s
}

func StatusPanelFromInst(inst uintptr) *TStatusPanel {
    s := new(TStatusPanel)
    s.instance = inst
    return s
}

func StatusPanelFromObj(obj IObject) *TStatusPanel {
    s := new(TStatusPanel)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStatusPanel) Free() {
    if s.instance != 0 {
        StatusPanel_Free(s.instance)
        s.instance = 0
    }
}

func (s *TStatusPanel) Instance() uintptr {
    return s.instance
}

func (s *TStatusPanel) IsValid() bool {
    return s.instance != 0
}

func (s *TStatusPanel) Assign(Source IObject) {
    StatusPanel_Assign(s.instance, CheckPtr(Source))
}

func (s *TStatusPanel) GetNamePath() string {
    return StatusPanel_GetNamePath(s.instance)
}

func (s *TStatusPanel) ClassName() string {
    return StatusPanel_ClassName(s.instance)
}

func (s *TStatusPanel) Equals(Obj IObject) bool {
    return StatusPanel_Equals(s.instance, CheckPtr(Obj))
}

func (s *TStatusPanel) GetHashCode() int32 {
    return StatusPanel_GetHashCode(s.instance)
}

func (s *TStatusPanel) ToString() string {
    return StatusPanel_ToString(s.instance)
}

func (s *TStatusPanel) Alignment() TAlignment {
    return StatusPanel_GetAlignment(s.instance)
}

func (s *TStatusPanel) SetAlignment(value TAlignment) {
    StatusPanel_SetAlignment(s.instance, value)
}

func (s *TStatusPanel) BiDiMode() TBiDiMode {
    return StatusPanel_GetBiDiMode(s.instance)
}

func (s *TStatusPanel) SetBiDiMode(value TBiDiMode) {
    StatusPanel_SetBiDiMode(s.instance, value)
}

func (s *TStatusPanel) Style() TStatusPanelStyle {
    return StatusPanel_GetStyle(s.instance)
}

func (s *TStatusPanel) SetStyle(value TStatusPanelStyle) {
    StatusPanel_SetStyle(s.instance, value)
}

func (s *TStatusPanel) Text() string {
    return StatusPanel_GetText(s.instance)
}

func (s *TStatusPanel) SetText(value string) {
    StatusPanel_SetText(s.instance, value)
}

func (s *TStatusPanel) Width() int32 {
    return StatusPanel_GetWidth(s.instance)
}

func (s *TStatusPanel) SetWidth(value int32) {
    StatusPanel_SetWidth(s.instance, value)
}

func (s *TStatusPanel) Index() int32 {
    return StatusPanel_GetIndex(s.instance)
}

func (s *TStatusPanel) SetIndex(value int32) {
    StatusPanel_SetIndex(s.instance, value)
}

