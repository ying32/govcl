package main

import (
	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl/types"
)

type TMainForm struct {
	*vcl.TForm
	lblEditLeft *vcl.TLabeledEdit
	lblEditTop *vcl.TLabeledEdit
	lblEditRight *vcl.TLabeledEdit
	lblEditBottom *vcl.TLabeledEdit
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.RunApp(&mainForm)
}

// --------------MainForm -----------------
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("Hello")
	f.EnabledMaximize(false)
	f.SetWidth(300)
	f.SetHeight(200)
	f.ScreenCenter()

	f.lblEditLeft = vcl.NewLabeledEdit(f)
	f.lblEditLeft.SetParent(f)
	f.lblEditLeft.SetLeft(50)
	f.lblEditLeft.SetTop(10)
	f.lblEditLeft.SetLabelPosition(types.LpLeft)
	f.lblEditLeft.SetLabelSpacing(6)
	f.lblEditLeft.EditLabel().SetCaption("left:")
	f.lblEditLeft.SetText("abc")
	
	
	f.lblEditTop = vcl.NewLabeledEdit(f)
	f.lblEditTop.SetParent(f)
	f.lblEditTop.SetLeft(50)
	f.lblEditTop.SetTop(140)
	//f.lblEditTop.SetLabelPosition(types.LpAbove) // 默认
	f.lblEditTop.SetLabelSpacing(6)
	f.lblEditTop.EditLabel().SetCaption("top:")
	f.lblEditTop.SetText("abc")
	
	
	f.lblEditRight = vcl.NewLabeledEdit(f)
	f.lblEditRight.SetParent(f)
	f.lblEditRight.SetLeft(170)
	f.lblEditRight.SetTop(10)
	f.lblEditRight.SetLabelPosition(types.LpRight)
	f.lblEditRight.SetLabelSpacing(6)
	f.lblEditRight.EditLabel().SetCaption("right")
	f.lblEditRight.SetText("abc")
	
	f.lblEditBottom = vcl.NewLabeledEdit(f)
	f.lblEditBottom.SetParent(f)
	f.lblEditBottom.SetLeft(200)
	f.lblEditBottom.SetTop(140)
	f.lblEditBottom.SetLabelPosition(types.LpBelow)
	f.lblEditBottom.SetLabelSpacing(6)
	f.lblEditBottom.EditLabel().SetCaption("bottom")
	f.lblEditBottom.SetText("abc")
}