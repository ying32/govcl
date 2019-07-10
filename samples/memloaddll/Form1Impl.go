package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func (m *TForm1) OnFormCreate(sender vcl.IObject) {
	fmt.Println(Form1.Caption(), m.PixelsPerInch())
	Form1.Button1.SetOnClick(func(sender vcl.IObject) {
		//vcl.ShowMessage("Hello!")
		jpg := vcl.NewJPEGImage()
		defer jpg.Free()
		jpg.LoadFromFile("a.jpg")
	})

	Form1.CheckBox1.SetOnClick(func(sender vcl.IObject) {
		Form1.Button1.SetEnabled(Form1.CheckBox1.Checked())

	})

	Form1.SetOnCloseQuery(func(sender vcl.IObject, canClose *bool) {
		fmt.Println("关闭。")
	})

	Form1.ActFileNew.SetOnExecute(func(vcl.IObject) {
		vcl.ShowMessage("ActFileNew Execute.")
	})
	Form1.Button2.SetOnClick(func(vcl.IObject) {
		result := Form2.ShowModal()
		if result == types.MrOk {
			vcl.ShowMessage("Form2返回了OK")
		} else if result == types.MrClose || result == types.MrNone {
			vcl.ShowMessage("Form2返回了Close")
		} else if result == types.MrCancel {
			vcl.ShowMessage("Form2返回了Cancel")
		}
	})
	Form1.ActExit.SetOnExecute(func(vcl.IObject) {
		vcl.Application.Terminate()
	})
}
