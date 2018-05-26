
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

type TPageSetupDialog struct {
    IComponent
    instance uintptr
}

func NewPageSetupDialog(owner IComponent) *TPageSetupDialog {
    p := new(TPageSetupDialog)
    p.instance = PageSetupDialog_Create(CheckPtr(owner))
    return p
}

func PageSetupDialogFromInst(inst uintptr) *TPageSetupDialog {
    p := new(TPageSetupDialog)
    p.instance = inst
    return p
}

func PageSetupDialogFromObj(obj IObject) *TPageSetupDialog {
    p := new(TPageSetupDialog)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPageSetupDialog) Free() {
    if p.instance != 0 {
        PageSetupDialog_Free(p.instance)
        p.instance = 0
    }
}

func (p *TPageSetupDialog) Instance() uintptr {
    return p.instance
}

func (p *TPageSetupDialog) IsValid() bool {
    return p.instance != 0
}

func (p *TPageSetupDialog) GetDefaults() bool {
    return PageSetupDialog_GetDefaults(p.instance)
}

func (p *TPageSetupDialog) Execute() bool {
    return PageSetupDialog_Execute(p.instance)
}

func (p *TPageSetupDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PageSetupDialog_FindComponent(p.instance, AName))
}

func (p *TPageSetupDialog) GetNamePath() string {
    return PageSetupDialog_GetNamePath(p.instance)
}

func (p *TPageSetupDialog) HasParent() bool {
    return PageSetupDialog_HasParent(p.instance)
}

func (p *TPageSetupDialog) Assign(Source IObject) {
    PageSetupDialog_Assign(p.instance, CheckPtr(Source))
}

func (p *TPageSetupDialog) ClassName() string {
    return PageSetupDialog_ClassName(p.instance)
}

func (p *TPageSetupDialog) Equals(Obj IObject) bool {
    return PageSetupDialog_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPageSetupDialog) GetHashCode() int32 {
    return PageSetupDialog_GetHashCode(p.instance)
}

func (p *TPageSetupDialog) ToString() string {
    return PageSetupDialog_ToString(p.instance)
}

func (p *TPageSetupDialog) MinMarginLeft() int32 {
    return PageSetupDialog_GetMinMarginLeft(p.instance)
}

func (p *TPageSetupDialog) SetMinMarginLeft(value int32) {
    PageSetupDialog_SetMinMarginLeft(p.instance, value)
}

func (p *TPageSetupDialog) MinMarginTop() int32 {
    return PageSetupDialog_GetMinMarginTop(p.instance)
}

func (p *TPageSetupDialog) SetMinMarginTop(value int32) {
    PageSetupDialog_SetMinMarginTop(p.instance, value)
}

func (p *TPageSetupDialog) MinMarginRight() int32 {
    return PageSetupDialog_GetMinMarginRight(p.instance)
}

func (p *TPageSetupDialog) SetMinMarginRight(value int32) {
    PageSetupDialog_SetMinMarginRight(p.instance, value)
}

func (p *TPageSetupDialog) MinMarginBottom() int32 {
    return PageSetupDialog_GetMinMarginBottom(p.instance)
}

func (p *TPageSetupDialog) SetMinMarginBottom(value int32) {
    PageSetupDialog_SetMinMarginBottom(p.instance, value)
}

func (p *TPageSetupDialog) MarginLeft() int32 {
    return PageSetupDialog_GetMarginLeft(p.instance)
}

func (p *TPageSetupDialog) SetMarginLeft(value int32) {
    PageSetupDialog_SetMarginLeft(p.instance, value)
}

func (p *TPageSetupDialog) MarginTop() int32 {
    return PageSetupDialog_GetMarginTop(p.instance)
}

func (p *TPageSetupDialog) SetMarginTop(value int32) {
    PageSetupDialog_SetMarginTop(p.instance, value)
}

func (p *TPageSetupDialog) MarginRight() int32 {
    return PageSetupDialog_GetMarginRight(p.instance)
}

func (p *TPageSetupDialog) SetMarginRight(value int32) {
    PageSetupDialog_SetMarginRight(p.instance, value)
}

func (p *TPageSetupDialog) MarginBottom() int32 {
    return PageSetupDialog_GetMarginBottom(p.instance)
}

func (p *TPageSetupDialog) SetMarginBottom(value int32) {
    PageSetupDialog_SetMarginBottom(p.instance, value)
}

func (p *TPageSetupDialog) Options() TPageSetupDialogOptions {
    return PageSetupDialog_GetOptions(p.instance)
}

func (p *TPageSetupDialog) SetOptions(value TPageSetupDialogOptions) {
    PageSetupDialog_SetOptions(p.instance, value)
}

func (p *TPageSetupDialog) PageWidth() int32 {
    return PageSetupDialog_GetPageWidth(p.instance)
}

func (p *TPageSetupDialog) SetPageWidth(value int32) {
    PageSetupDialog_SetPageWidth(p.instance, value)
}

func (p *TPageSetupDialog) PageHeight() int32 {
    return PageSetupDialog_GetPageHeight(p.instance)
}

func (p *TPageSetupDialog) SetPageHeight(value int32) {
    PageSetupDialog_SetPageHeight(p.instance, value)
}

func (p *TPageSetupDialog) Units() TPageMeasureUnits {
    return PageSetupDialog_GetUnits(p.instance)
}

func (p *TPageSetupDialog) SetUnits(value TPageMeasureUnits) {
    PageSetupDialog_SetUnits(p.instance, value)
}

func (p *TPageSetupDialog) Handle() HWND {
    return PageSetupDialog_GetHandle(p.instance)
}

func (p *TPageSetupDialog) SetOnClose(fn TNotifyEvent) {
    PageSetupDialog_SetOnClose(p.instance, fn)
}

func (p *TPageSetupDialog) SetOnShow(fn TNotifyEvent) {
    PageSetupDialog_SetOnShow(p.instance, fn)
}

func (p *TPageSetupDialog) ComponentCount() int32 {
    return PageSetupDialog_GetComponentCount(p.instance)
}

func (p *TPageSetupDialog) ComponentIndex() int32 {
    return PageSetupDialog_GetComponentIndex(p.instance)
}

func (p *TPageSetupDialog) SetComponentIndex(value int32) {
    PageSetupDialog_SetComponentIndex(p.instance, value)
}

func (p *TPageSetupDialog) Owner() *TComponent {
    return ComponentFromInst(PageSetupDialog_GetOwner(p.instance))
}

func (p *TPageSetupDialog) Name() string {
    return PageSetupDialog_GetName(p.instance)
}

func (p *TPageSetupDialog) SetName(value string) {
    PageSetupDialog_SetName(p.instance, value)
}

func (p *TPageSetupDialog) Tag() int {
    return PageSetupDialog_GetTag(p.instance)
}

func (p *TPageSetupDialog) SetTag(value int) {
    PageSetupDialog_SetTag(p.instance, value)
}

func (p *TPageSetupDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PageSetupDialog_GetComponents(p.instance, AIndex))
}

