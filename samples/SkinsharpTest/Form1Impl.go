// 在这里写你的事件。
// Write your event here.

package main

import (
	"github.com/ying32/govcl/pkgs/skinh"
	"github.com/ying32/govcl/vcl"
)

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	skinh.Attach()

}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {
	skinh.Detach()
}

func (f *TForm1) OnButton6Click(sender vcl.IObject) {
	f.Close()
}

func (f *TForm1) OnButton5Click(sender vcl.IObject) {
	skinh.SetAero(0)
}

func (f *TForm1) OnButton4Click(sender vcl.IObject) {
	skinh.SetAero(1)
}

func (f *TForm1) OnButton3Click(sender vcl.IObject) {
	Form6.Show()
}

func (f *TForm1) OnButton2Click(sender vcl.IObject) {
	if f.OpenDialog1.Execute() {
		skinh.AttachEx(f.OpenDialog1.FileName(), "")
	}
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	Form2.Show()
}

func (f *TForm1) OnButton7Click(sender vcl.IObject) {
	Form3.Show()
}

func (f *TForm1) OnButton8Click(sender vcl.IObject) {
	Form4.Show()
}

func (f *TForm1) OnButton9Click(sender vcl.IObject) {
	f.Label1.SetCaption("fffff")
	f.SetAutoSize(false)
}

func (f *TForm1) OnButton10Click(sender vcl.IObject) {
	Form5.Show()
}
