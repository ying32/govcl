package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/win"

	"github.com/ying32/govcl/vcl/types/keys"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	win.OutputDebugString(Form1.Caption(), f.PixelsPerInch())
}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {
	win.OutputDebugString("----------------------OnFormDestroy")
}

func (f *TForm1) memoOnKeyup(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
	if shift.In(types.SsCtrl) && *key == /*keys.VkA*/ keys.VkB {
		vcl.AsMemo(sender).SelectAll()
	}
}

func (f *TForm1) OnActExitExecute(vcl.IObject) {

}

func (f *TForm1) OnButton2Click(vcl.IObject) {
	win.OutputDebugString("currentid-1: ", win.GetCurrentThreadId())
	go func() {
		win.OutputDebugString("OnButton2Click,  currentid-2", win.GetCurrentThreadId())

		vcl.ThreadSync(func() {
			win.OutputDebugString("ThreadSync OnButton2Click-3", win.GetCurrentThreadId())
			vcl.ShowMessage("测试")
		})
	}()
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
	vcl.ShowMessage("Hello!")

}
