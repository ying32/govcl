
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

type TParaAttributes struct {
    IObject
    instance uintptr
}

func ParaAttributesFromInst(inst uintptr) *TParaAttributes {
    p := new(TParaAttributes)
    p.instance = inst
    return p
}

func ParaAttributesFromObj(obj IObject) *TParaAttributes {
    p := new(TParaAttributes)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TParaAttributes) Instance() uintptr {
    return p.instance
}

func (p *TParaAttributes) IsValid() bool {
    return p.instance != 0
}

func (p *TParaAttributes) Assign(Source IObject) {
    ParaAttributes_Assign(p.instance, CheckPtr(Source))
}

func (p *TParaAttributes) GetNamePath() string {
    return ParaAttributes_GetNamePath(p.instance)
}

func (p *TParaAttributes) ClassName() string {
    return ParaAttributes_ClassName(p.instance)
}

func (p *TParaAttributes) Equals(Obj IObject) bool {
    return ParaAttributes_Equals(p.instance, CheckPtr(Obj))
}

func (p *TParaAttributes) GetHashCode() int32 {
    return ParaAttributes_GetHashCode(p.instance)
}

func (p *TParaAttributes) ToString() string {
    return ParaAttributes_ToString(p.instance)
}

func (p *TParaAttributes) Alignment() TAlignment {
    return ParaAttributes_GetAlignment(p.instance)
}

func (p *TParaAttributes) SetAlignment(value TAlignment) {
    ParaAttributes_SetAlignment(p.instance, value)
}

func (p *TParaAttributes) FirstIndent() int32 {
    return ParaAttributes_GetFirstIndent(p.instance)
}

func (p *TParaAttributes) SetFirstIndent(value int32) {
    ParaAttributes_SetFirstIndent(p.instance, value)
}

func (p *TParaAttributes) LeftIndent() int32 {
    return ParaAttributes_GetLeftIndent(p.instance)
}

func (p *TParaAttributes) SetLeftIndent(value int32) {
    ParaAttributes_SetLeftIndent(p.instance, value)
}

func (p *TParaAttributes) Numbering() TNumberingStyle {
    return ParaAttributes_GetNumbering(p.instance)
}

func (p *TParaAttributes) SetNumbering(value TNumberingStyle) {
    ParaAttributes_SetNumbering(p.instance, value)
}

func (p *TParaAttributes) RightIndent() int32 {
    return ParaAttributes_GetRightIndent(p.instance)
}

func (p *TParaAttributes) SetRightIndent(value int32) {
    ParaAttributes_SetRightIndent(p.instance, value)
}

func (p *TParaAttributes) TabCount() int32 {
    return ParaAttributes_GetTabCount(p.instance)
}

func (p *TParaAttributes) SetTabCount(value int32) {
    ParaAttributes_SetTabCount(p.instance, value)
}

func (p *TParaAttributes) Tab(Index uint8) int32 {
    return ParaAttributes_GetTab(p.instance, Index)
}

func (p *TParaAttributes) SetTab(Index uint8, value int32) {
    ParaAttributes_SetTab(p.instance, Index, value)
}

