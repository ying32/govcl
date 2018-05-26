
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TStatusPanels struct {
    IObject
    instance uintptr
}

func NewStatusPanels() *TStatusPanels {
    s := new(TStatusPanels)
    s.instance = StatusPanels_Create()
    return s
}

func StatusPanelsFromInst(inst uintptr) *TStatusPanels {
    s := new(TStatusPanels)
    s.instance = inst
    return s
}

func StatusPanelsFromObj(obj IObject) *TStatusPanels {
    s := new(TStatusPanels)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStatusPanels) Free() {
    if s.instance != 0 {
        StatusPanels_Free(s.instance)
        s.instance = 0
    }
}

func (s *TStatusPanels) Instance() uintptr {
    return s.instance
}

func (s *TStatusPanels) IsValid() bool {
    return s.instance != 0
}

func (s *TStatusPanels) Add() *TStatusPanel {
    return StatusPanelFromInst(StatusPanels_Add(s.instance))
}

func (s *TStatusPanels) AddItem(Item *TStatusPanel, Index int32) *TStatusPanel {
    return StatusPanelFromInst(StatusPanels_AddItem(s.instance, CheckPtr(Item), Index))
}

func (s *TStatusPanels) Insert(Index int32) *TStatusPanel {
    return StatusPanelFromInst(StatusPanels_Insert(s.instance, Index))
}

func (s *TStatusPanels) Owner() *TObject {
    return ObjectFromInst(StatusPanels_Owner(s.instance))
}

func (s *TStatusPanels) Assign(Source IObject) {
    StatusPanels_Assign(s.instance, CheckPtr(Source))
}

func (s *TStatusPanels) BeginUpdate() {
    StatusPanels_BeginUpdate(s.instance)
}

func (s *TStatusPanels) Clear() {
    StatusPanels_Clear(s.instance)
}

func (s *TStatusPanels) Delete(Index int32) {
    StatusPanels_Delete(s.instance, Index)
}

func (s *TStatusPanels) EndUpdate() {
    StatusPanels_EndUpdate(s.instance)
}

func (s *TStatusPanels) GetNamePath() string {
    return StatusPanels_GetNamePath(s.instance)
}

func (s *TStatusPanels) ClassName() string {
    return StatusPanels_ClassName(s.instance)
}

func (s *TStatusPanels) Equals(Obj IObject) bool {
    return StatusPanels_Equals(s.instance, CheckPtr(Obj))
}

func (s *TStatusPanels) GetHashCode() int32 {
    return StatusPanels_GetHashCode(s.instance)
}

func (s *TStatusPanels) ToString() string {
    return StatusPanels_ToString(s.instance)
}

func (s *TStatusPanels) Capacity() int32 {
    return StatusPanels_GetCapacity(s.instance)
}

func (s *TStatusPanels) SetCapacity(value int32) {
    StatusPanels_SetCapacity(s.instance, value)
}

func (s *TStatusPanels) Count() int32 {
    return StatusPanels_GetCount(s.instance)
}

func (s *TStatusPanels) Items(Index int32) *TStatusPanel {
    return StatusPanelFromInst(StatusPanels_GetItems(s.instance, Index))
}

func (s *TStatusPanels) SetItems(Index int32, value *TStatusPanel) {
    StatusPanels_SetItems(s.instance, Index, CheckPtr(value))
}

