package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/types/keys"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	fmt.Println(Form1.Caption(), f.PixelsPerInch())

	// 遍历组件
	// 只要owner设置为Form的都可以通过这个方法来遍历。
	var i int32
	for i = 0; i < f.ComponentCount(); i++ {
		comp := f.Components(i)
		//fmt.Println(i, "=", comp.Name())
		if comp.Is().Memo() {
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

func (f *TForm1) OnActExitExecute(vcl.IObject) {
	vcl.Application.Terminate()
}

func (f *TForm1) OnButton2Click(vcl.IObject) {
	result := Form2.ShowModal()
	if result == types.MrOk {
		vcl.ShowMessage("Form2返回了OK")
	} else if result == types.MrClose || result == types.MrNone {
		vcl.ShowMessage("Form2返回了Close")
	} else if result == types.MrCancel {
		vcl.ShowMessage("Form2返回了Cancel")
	}
}

func (f *TForm1) OnActFileNewExecute(vcl.IObject) {
	vcl.ShowMessage("ActFileNew Execute.")
}

func (f *TForm1) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	fmt.Println("关闭。")
}

func (f *TForm1) OnCheckBox1Click(sender vcl.IObject) {
	f.Button1.SetEnabled(f.CheckBox1.Checked())
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	//vcl.ShowMessage("Hello!")
	jpg := vcl.NewJPEGImage()
	defer jpg.Free()
	jpg.LoadFromFile("a.jpg")
}
