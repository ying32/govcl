
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

type TPrinterSetupDialog struct {
    IComponent
    instance uintptr
}

func NewPrinterSetupDialog(owner IComponent) *TPrinterSetupDialog {
    p := new(TPrinterSetupDialog)
    p.instance = PrinterSetupDialog_Create(CheckPtr(owner))
    return p
}

func PrinterSetupDialogFromInst(inst uintptr) *TPrinterSetupDialog {
    p := new(TPrinterSetupDialog)
    p.instance = inst
    return p
}

func PrinterSetupDialogFromObj(obj IObject) *TPrinterSetupDialog {
    p := new(TPrinterSetupDialog)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPrinterSetupDialog) Free() {
    if p.instance != 0 {
        PrinterSetupDialog_Free(p.instance)
        p.instance = 0
    }
}

func (p *TPrinterSetupDialog) Instance() uintptr {
    return p.instance
}

func (p *TPrinterSetupDialog) IsValid() bool {
    return p.instance != 0
}

func (p *TPrinterSetupDialog) Execute() bool {
    return PrinterSetupDialog_Execute(p.instance)
}

func (p *TPrinterSetupDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PrinterSetupDialog_FindComponent(p.instance, AName))
}

func (p *TPrinterSetupDialog) GetNamePath() string {
    return PrinterSetupDialog_GetNamePath(p.instance)
}

func (p *TPrinterSetupDialog) HasParent() bool {
    return PrinterSetupDialog_HasParent(p.instance)
}

func (p *TPrinterSetupDialog) Assign(Source IObject) {
    PrinterSetupDialog_Assign(p.instance, CheckPtr(Source))
}

func (p *TPrinterSetupDialog) ClassName() string {
    return PrinterSetupDialog_ClassName(p.instance)
}

func (p *TPrinterSetupDialog) Equals(Obj IObject) bool {
    return PrinterSetupDialog_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPrinterSetupDialog) GetHashCode() int32 {
    return PrinterSetupDialog_GetHashCode(p.instance)
}

func (p *TPrinterSetupDialog) ToString() string {
    return PrinterSetupDialog_ToString(p.instance)
}

func (p *TPrinterSetupDialog) Handle() HWND {
    return PrinterSetupDialog_GetHandle(p.instance)
}

func (p *TPrinterSetupDialog) SetOnClose(fn TNotifyEvent) {
    PrinterSetupDialog_SetOnClose(p.instance, fn)
}

func (p *TPrinterSetupDialog) SetOnShow(fn TNotifyEvent) {
    PrinterSetupDialog_SetOnShow(p.instance, fn)
}

func (p *TPrinterSetupDialog) ComponentCount() int32 {
    return PrinterSetupDialog_GetComponentCount(p.instance)
}

func (p *TPrinterSetupDialog) ComponentIndex() int32 {
    return PrinterSetupDialog_GetComponentIndex(p.instance)
}

func (p *TPrinterSetupDialog) SetComponentIndex(value int32) {
    PrinterSetupDialog_SetComponentIndex(p.instance, value)
}

func (p *TPrinterSetupDialog) Owner() *TComponent {
    return ComponentFromInst(PrinterSetupDialog_GetOwner(p.instance))
}

func (p *TPrinterSetupDialog) Name() string {
    return PrinterSetupDialog_GetName(p.instance)
}

func (p *TPrinterSetupDialog) SetName(value string) {
    PrinterSetupDialog_SetName(p.instance, value)
}

func (p *TPrinterSetupDialog) Tag() int {
    return PrinterSetupDialog_GetTag(p.instance)
}

func (p *TPrinterSetupDialog) SetTag(value int) {
    PrinterSetupDialog_SetTag(p.instance, value)
}

func (p *TPrinterSetupDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PrinterSetupDialog_GetComponents(p.instance, AIndex))
}

