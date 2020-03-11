package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/types/keys"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	fmt.Println(Form1.Caption(), f.PixelsPerInch())
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

	// 遍历组件
	// 只要owner设置为Form的都可以通过这个方法来遍历。
	var i int32
	for i = 0; i < f.ComponentCount(); i++ {
		comp := f.Components(i)
		//fmt.Println(i, "=", comp.Name())
		if comp.InheritsFrom(vcl.TMemoClass()) {
			fmt.Println(i, "=", comp.Name(), ", 继承自TMemo")
			mem := vcl.AsMemo(comp)
			mem.SetOnKeyUp(f.memoOnKeyup)
		}
	}
}

func (f *TForm1) memoOnKeyup(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
	if shift.In(types.SsCtrl) && *key == /*keys.VkA*/ keys.VkB {
		vcl.AsMemo(sender).SelectAll()
	}
}
