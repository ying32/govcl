package main

import (
	"gitee.com/ying32/govcl/vcl"

	"gitee.com/ying32/govcl/vcl/types"
)

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

	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetCaption("按钮1")
	btn.SetLeft(50)
	btn.SetTop(50)

	vcl.Application.Run()
}
