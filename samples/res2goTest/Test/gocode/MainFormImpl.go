// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"github.com/ying32/govcl/vcl"
)

//::private::
type TMainFormFields struct {
	frame1 *TFrame2
	frame2 *TFrame3
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.frame1 = NewFrame2(f)
	f.frame2 = NewFrame3(f)

	f.frame1.Hide()
	f.frame2.Hide()

	f.frame1.SetParent(f.Panel2)
	f.frame2.SetParent(f.Panel2)
}

func (f *TMainForm) OnBtnShowAboutClick(sender vcl.IObject) {
	// 动态创建，完后释放掉，不使用Application.CreateForm
	about := NewAbout(f)
	defer about.Free()
	about.ShowModal()
}

func (f *TMainForm) OnBtnShowFrame1Click(sender vcl.IObject) {
	f.frame2.Hide()
	f.frame1.Show()
}

func (f *TMainForm) OnBtnShowFrame2Click(sender vcl.IObject) {
	f.frame1.Hide()
	f.frame2.Show()
}
