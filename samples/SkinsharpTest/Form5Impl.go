// 在这里写你的事件。
// Write your event here.

package main

import "github.com/ying32/govcl/vcl"

func (f *TForm5) OnOpenProject1Click(sender vcl.IObject) {
	f.OpenDialog1.Execute()
}

func (f *TForm5) OnButton1Click(sender vcl.IObject) {
	f.OpenDialog1.Execute()
}

func (f *TForm5) OnButton2Click(sender vcl.IObject) {
	f.SaveDialog1.Execute()
}

func (f *TForm5) OnButton3Click(sender vcl.IObject) {
	f.OpenPictureDialog1.Execute()
}

func (f *TForm5) OnButton4Click(sender vcl.IObject) {
	f.FontDialog1.Execute()
}

func (f *TForm5) OnButton5Click(sender vcl.IObject) {
	f.ColorDialog1.Execute()
}

func (f *TForm5) OnButton7Click(sender vcl.IObject) {
	f.PrintDialog1.Execute()
}

func (f *TForm5) OnButton8Click(sender vcl.IObject) {
	f.PrinterSetupDialog1.Execute()
}

func (f *TForm5) OnButton9Click(sender vcl.IObject) {
	f.FindDialog1.Execute()
}

func (f *TForm5) OnButton10Click(sender vcl.IObject) {
	f.ReplaceDialog1.Execute()
}

func (f *TForm5) OnButton11Click(sender vcl.IObject) {
	f.PageSetupDialog1.Execute()
}
