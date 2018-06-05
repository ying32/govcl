
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

type THeaderSection struct {
    IObject
    instance uintptr
}

func NewHeaderSection() *THeaderSection {
    h := new(THeaderSection)
    h.instance = HeaderSection_Create()
    return h
}

func HeaderSectionFromInst(inst uintptr) *THeaderSection {
    h := new(THeaderSection)
    h.instance = inst
    return h
}

func HeaderSectionFromObj(obj IObject) *THeaderSection {
    h := new(THeaderSection)
    h.instance = CheckPtr(obj)
    return h
}

func (h *THeaderSection) Free() {
    if h.instance != 0 {
        HeaderSection_Free(h.instance)
        h.instance = 0
    }
}

func (h *THeaderSection) Instance() uintptr {
    return h.instance
}

func (h *THeaderSection) IsValid() bool {
    return h.instance != 0
}

func THeaderSectionClass() TClass {
    return HeaderSection_StaticClassType()
}

func (h *THeaderSection) Assign(Source IObject) {
    HeaderSection_Assign(h.instance, CheckPtr(Source))
}

func (h *THeaderSection) GetNamePath() string {
    return HeaderSection_GetNamePath(h.instance)
}

func (h *THeaderSection) DisposeOf() {
    HeaderSection_DisposeOf(h.instance)
}

func (h *THeaderSection) ClassType() TClass {
    return HeaderSection_ClassType(h.instance)
}

func (h *THeaderSection) ClassName() string {
    return HeaderSection_ClassName(h.instance)
}

func (h *THeaderSection) InstanceSize() int32 {
    return HeaderSection_InstanceSize(h.instance)
}

func (h *THeaderSection) InheritsFrom(AClass TClass) bool {
    return HeaderSection_InheritsFrom(h.instance, AClass)
}

func (h *THeaderSection) Equals(Obj IObject) bool {
    return HeaderSection_Equals(h.instance, CheckPtr(Obj))
}

func (h *THeaderSection) GetHashCode() int32 {
    return HeaderSection_GetHashCode(h.instance)
}

func (h *THeaderSection) ToString() string {
    return HeaderSection_ToString(h.instance)
}

func (h *THeaderSection) Left() int32 {
    return HeaderSection_GetLeft(h.instance)
}

func (h *THeaderSection) Right() int32 {
    return HeaderSection_GetRight(h.instance)
}

func (h *THeaderSection) Alignment() TAlignment {
    return HeaderSection_GetAlignment(h.instance)
}

func (h *THeaderSection) SetAlignment(value TAlignment) {
    HeaderSection_SetAlignment(h.instance, value)
}

func (h *THeaderSection) AllowClick() bool {
    return HeaderSection_GetAllowClick(h.instance)
}

func (h *THeaderSection) SetAllowClick(value bool) {
    HeaderSection_SetAllowClick(h.instance, value)
}

func (h *THeaderSection) AutoSize() bool {
    return HeaderSection_GetAutoSize(h.instance)
}

func (h *THeaderSection) SetAutoSize(value bool) {
    HeaderSection_SetAutoSize(h.instance, value)
}

func (h *THeaderSection) BiDiMode() TBiDiMode {
    return HeaderSection_GetBiDiMode(h.instance)
}

func (h *THeaderSection) SetBiDiMode(value TBiDiMode) {
    HeaderSection_SetBiDiMode(h.instance, value)
}

func (h *THeaderSection) CheckBox() bool {
    return HeaderSection_GetCheckBox(h.instance)
}

func (h *THeaderSection) SetCheckBox(value bool) {
    HeaderSection_SetCheckBox(h.instance, value)
}

func (h *THeaderSection) Checked() bool {
    return HeaderSection_GetChecked(h.instance)
}

func (h *THeaderSection) SetChecked(value bool) {
    HeaderSection_SetChecked(h.instance, value)
}

func (h *THeaderSection) FixedWidth() bool {
    return HeaderSection_GetFixedWidth(h.instance)
}

func (h *THeaderSection) SetFixedWidth(value bool) {
    HeaderSection_SetFixedWidth(h.instance, value)
}

func (h *THeaderSection) ImageIndex() int32 {
    return HeaderSection_GetImageIndex(h.instance)
}

func (h *THeaderSection) SetImageIndex(value int32) {
    HeaderSection_SetImageIndex(h.instance, value)
}

func (h *THeaderSection) MaxWidth() int32 {
    return HeaderSection_GetMaxWidth(h.instance)
}

func (h *THeaderSection) SetMaxWidth(value int32) {
    HeaderSection_SetMaxWidth(h.instance, value)
}

func (h *THeaderSection) MinWidth() int32 {
    return HeaderSection_GetMinWidth(h.instance)
}

func (h *THeaderSection) SetMinWidth(value int32) {
    HeaderSection_SetMinWidth(h.instance, value)
}

func (h *THeaderSection) Style() THeaderSectionStyle {
    return HeaderSection_GetStyle(h.instance)
}

func (h *THeaderSection) SetStyle(value THeaderSectionStyle) {
    HeaderSection_SetStyle(h.instance, value)
}

func (h *THeaderSection) Text() string {
    return HeaderSection_GetText(h.instance)
}

func (h *THeaderSection) SetText(value string) {
    HeaderSection_SetText(h.instance, value)
}

func (h *THeaderSection) Width() int32 {
    return HeaderSection_GetWidth(h.instance)
}

func (h *THeaderSection) SetWidth(value int32) {
    HeaderSection_SetWidth(h.instance, value)
}

func (h *THeaderSection) Index() int32 {
    return HeaderSection_GetIndex(h.instance)
}

func (h *THeaderSection) SetIndex(value int32) {
    HeaderSection_SetIndex(h.instance, value)
}

