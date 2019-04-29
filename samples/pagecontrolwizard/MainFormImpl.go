// 在这里写你的事件

package main

import (
	"github.com/ying32/govcl/vcl"
)

//::private::
type TMainFormFields struct {
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.hideAllTab()
	f.PageControl1.SetActivePageIndex(0)
}

func (f *TMainForm) hideAllTab() {
	var i int32
	for i = 0; i < f.PageControl1.PageCount(); i++ {
		f.PageControl1.Pages(i).SetTabVisible(false)
	}
}

func (f *TMainForm) OnActPagePrevExecute(sender vcl.IObject) {
	f.PageControl1.SetActivePageIndex(f.PageControl1.ActivePageIndex() - 1)
}

func (f *TMainForm) OnActPagePrevUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.PageControl1.ActivePageIndex() > 0)
}

func (f *TMainForm) OnActPageNextExecute(sender vcl.IObject) {
	f.PageControl1.SetActivePageIndex(f.PageControl1.ActivePageIndex() + 1)
}

func (f *TMainForm) OnActPageNextUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.PageControl1.ActivePageIndex() < f.PageControl1.PageCount()-1)
}
