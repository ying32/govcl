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
	//f.setPage(0)
	f.PageControl1.SetActivePageIndex(0)
}

func (f *TMainForm) hideAllTab() {
	var i int32
	for i = 0; i < f.PageControl1.PageCount(); i++ {
		sheet := f.PageControl1.Pages(i)
		sheet.SetTabVisible(false)
		sheet.SetVisible(false)
	}
}

/*func (f *TMainForm) setPage(idx int32) {
	if idx != 0 && idx != -1 && idx != 1 {
		return
	}
	if idx == 0 {
		f.PageControl1.SetActivePageIndex(0)
		sheet := f.PageControl1.Pages(0)
		sheet.SetVisible(true)
		return
	}
	curIdx := f.PageControl1.ActivePageIndex()
	sheet := f.PageControl1.Pages(curIdx)
	sheet.SetVisible(false)
	f.PageControl1.SetActivePageIndex(curIdx + idx)
	sheet = f.PageControl1.Pages(curIdx + idx)
	sheet.SetVisible(true)
}*/

func (f *TMainForm) OnActPagePrevExecute(sender vcl.IObject) {
	f.PageControl1.SetActivePageIndex(f.PageControl1.ActivePageIndex() - 1)
	//f.setPage(-1)
}

func (f *TMainForm) OnActPagePrevUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.PageControl1.ActivePageIndex() > 0)
}

func (f *TMainForm) OnActPageNextExecute(sender vcl.IObject) {
	f.PageControl1.SetActivePageIndex(f.PageControl1.ActivePageIndex() + 1)
	//f.setPage(1)
}

func (f *TMainForm) OnActPageNextUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.PageControl1.ActivePageIndex() < f.PageControl1.PageCount()-1)
}
