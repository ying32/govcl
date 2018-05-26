
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

type TPrintDialog struct {
    IComponent
    instance uintptr
}

func NewPrintDialog(owner IComponent) *TPrintDialog {
    p := new(TPrintDialog)
    p.instance = PrintDialog_Create(CheckPtr(owner))
    return p
}

func PrintDialogFromInst(inst uintptr) *TPrintDialog {
    p := new(TPrintDialog)
    p.instance = inst
    return p
}

func PrintDialogFromObj(obj IObject) *TPrintDialog {
    p := new(TPrintDialog)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPrintDialog) Free() {
    if p.instance != 0 {
        PrintDialog_Free(p.instance)
        p.instance = 0
    }
}

func (p *TPrintDialog) Instance() uintptr {
    return p.instance
}

func (p *TPrintDialog) IsValid() bool {
    return p.instance != 0
}

func (p *TPrintDialog) Execute() bool {
    return PrintDialog_Execute(p.instance)
}

func (p *TPrintDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PrintDialog_FindComponent(p.instance, AName))
}

func (p *TPrintDialog) GetNamePath() string {
    return PrintDialog_GetNamePath(p.instance)
}

func (p *TPrintDialog) HasParent() bool {
    return PrintDialog_HasParent(p.instance)
}

func (p *TPrintDialog) Assign(Source IObject) {
    PrintDialog_Assign(p.instance, CheckPtr(Source))
}

func (p *TPrintDialog) ClassName() string {
    return PrintDialog_ClassName(p.instance)
}

func (p *TPrintDialog) Equals(Obj IObject) bool {
    return PrintDialog_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPrintDialog) GetHashCode() int32 {
    return PrintDialog_GetHashCode(p.instance)
}

func (p *TPrintDialog) ToString() string {
    return PrintDialog_ToString(p.instance)
}

func (p *TPrintDialog) Collate() bool {
    return PrintDialog_GetCollate(p.instance)
}

func (p *TPrintDialog) SetCollate(value bool) {
    PrintDialog_SetCollate(p.instance, value)
}

func (p *TPrintDialog) Copies() int32 {
    return PrintDialog_GetCopies(p.instance)
}

func (p *TPrintDialog) SetCopies(value int32) {
    PrintDialog_SetCopies(p.instance, value)
}

func (p *TPrintDialog) FromPage() int32 {
    return PrintDialog_GetFromPage(p.instance)
}

func (p *TPrintDialog) SetFromPage(value int32) {
    PrintDialog_SetFromPage(p.instance, value)
}

func (p *TPrintDialog) MinPage() int32 {
    return PrintDialog_GetMinPage(p.instance)
}

func (p *TPrintDialog) SetMinPage(value int32) {
    PrintDialog_SetMinPage(p.instance, value)
}

func (p *TPrintDialog) MaxPage() int32 {
    return PrintDialog_GetMaxPage(p.instance)
}

func (p *TPrintDialog) SetMaxPage(value int32) {
    PrintDialog_SetMaxPage(p.instance, value)
}

func (p *TPrintDialog) Options() TPrintDialogOptions {
    return PrintDialog_GetOptions(p.instance)
}

func (p *TPrintDialog) SetOptions(value TPrintDialogOptions) {
    PrintDialog_SetOptions(p.instance, value)
}

func (p *TPrintDialog) PrintToFile() bool {
    return PrintDialog_GetPrintToFile(p.instance)
}

func (p *TPrintDialog) SetPrintToFile(value bool) {
    PrintDialog_SetPrintToFile(p.instance, value)
}

func (p *TPrintDialog) PrintRange() TPrintRange {
    return PrintDialog_GetPrintRange(p.instance)
}

func (p *TPrintDialog) SetPrintRange(value TPrintRange) {
    PrintDialog_SetPrintRange(p.instance, value)
}

func (p *TPrintDialog) ToPage() int32 {
    return PrintDialog_GetToPage(p.instance)
}

func (p *TPrintDialog) SetToPage(value int32) {
    PrintDialog_SetToPage(p.instance, value)
}

func (p *TPrintDialog) Handle() HWND {
    return PrintDialog_GetHandle(p.instance)
}

func (p *TPrintDialog) SetOnClose(fn TNotifyEvent) {
    PrintDialog_SetOnClose(p.instance, fn)
}

func (p *TPrintDialog) SetOnShow(fn TNotifyEvent) {
    PrintDialog_SetOnShow(p.instance, fn)
}

func (p *TPrintDialog) ComponentCount() int32 {
    return PrintDialog_GetComponentCount(p.instance)
}

func (p *TPrintDialog) ComponentIndex() int32 {
    return PrintDialog_GetComponentIndex(p.instance)
}

func (p *TPrintDialog) SetComponentIndex(value int32) {
    PrintDialog_SetComponentIndex(p.instance, value)
}

func (p *TPrintDialog) Owner() *TComponent {
    return ComponentFromInst(PrintDialog_GetOwner(p.instance))
}

func (p *TPrintDialog) Name() string {
    return PrintDialog_GetName(p.instance)
}

func (p *TPrintDialog) SetName(value string) {
    PrintDialog_SetName(p.instance, value)
}

func (p *TPrintDialog) Tag() int {
    return PrintDialog_GetTag(p.instance)
}

func (p *TPrintDialog) SetTag(value int) {
    PrintDialog_SetTag(p.instance, value)
}

func (p *TPrintDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PrintDialog_GetComponents(p.instance, AIndex))
}

