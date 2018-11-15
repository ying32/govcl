// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"github.com/ying32/govcl/vcl"
)

//::private::
type TMainFormFields struct {
	test string
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.test = "aaaaaa"
}

func (f *TMainForm) OnButton1Click(sender vcl.IObject) {
	About.SetCaption(f.test)
	About.Show()
}

func (f *TMainForm) OnButton2Click(sender vcl.IObject) {
	btn := vcl.ButtonFromObj(sender)
	vcl.ShowMessage("Name:" + btn.Name())
}

func (f *TMainForm) OnButton2Resize(sender vcl.IObject) {

}

func (f *TMainForm) OnPanel1MouseEnter(sender vcl.IObject) {
	pnl := vcl.PanelFromObj(sender)
	pnl.SetLeft(pnl.Left() + 5)
}

func (f *TMainForm) OnPanel1MouseLeave(sender vcl.IObject) {
	pnl := vcl.PanelFromObj(sender)
	pnl.SetLeft(pnl.Left() - 5)
}
