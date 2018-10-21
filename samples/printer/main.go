package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TMainForm struct {
	*vcl.TForm
	CbbPrinters *vcl.TComboBox
	Btn1        *vcl.TButton
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&mainForm)
	vcl.Application.Run()
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 88, 28)
	f.Btn1.SetCaption("打印")
	f.Btn1.SetOnClick(f.OnButtonClick)

	f.CbbPrinters = vcl.NewComboBox(f)
	f.CbbPrinters.SetParent(f)
	f.CbbPrinters.SetBounds(f.Btn1.Left(), f.Btn1.Top()+f.Btn1.Height()+20, 200, f.CbbPrinters.Height())
	f.CbbPrinters.Items().Assign(vcl.Printer.Printers())
	f.CbbPrinters.SetOnChange(f.OnCbbChange)

	// Printer.Orientation 方向

}

func (f *TMainForm) OnButtonClick(sender vcl.IObject) {

	if f.CbbPrinters.ItemIndex() == -1 {
		vcl.ShowMessage("先设置一个打印机.")
		return
	}

	//vcl.Printer.SetTitle("标题啊。。。。 ")

	vcl.Printer.BeginDoc()
	defer vcl.Printer.EndDoc()

	canvas := vcl.Printer.Canvas()
	//canvas.Brush().SetColor(colors.ClWhite)
	canvas.Brush().SetStyle(types.BsClear)

	font := canvas.Font()
	font.SetName("微软雅黑")
	font.SetSize(16)
	font.SetColor(colors.ClGreen)

	canvas.TextOut(0, 0, "这是一个测试")
}

func (f *TMainForm) OnCbbChange(sender vcl.IObject) {
	if f.CbbPrinters.ItemIndex() != -1 {
		vcl.Printer.SetPrinterIndex(f.CbbPrinters.ItemIndex())
	}
}
