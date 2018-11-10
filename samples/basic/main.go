package main

import (
	"github.com/ying32/govcl/vcl"

	"fmt"

	"github.com/ying32/govcl/vcl/types"
)

type TForm1 struct {
	*vcl.TForm
	Button1 *vcl.TButton
}

var form1 *TForm1

func main() {

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)
	mainForm.SetOnCloseQuery(func(Sender vcl.IObject, CanClose *bool) {
		*CanClose = vcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
	})

	vcl.Application.CreateForm(&form1)

	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetCaption("窗口1")
	btn.SetLeft(50)
	btn.SetTop(50)
	btn.SetOnClick(func(sender vcl.IObject) {
		form1.Show()
	})

	vcl.Application.Run()
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	fmt.Println("onCreate")
	f.Button1 = vcl.NewButton(f)
	f.Button1.SetParent(f)
	f.Button1.SetCaption("我是按钮")
	f.Button1.SetOnClick(f.OnButton1Click)
}

func (f *TForm1) OnButton1Click(object vcl.IObject) {
	vcl.ShowMessage("Click")
}
