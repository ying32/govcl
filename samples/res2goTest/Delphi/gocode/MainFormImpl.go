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
	f.test = "ffffff"
}


func (f *TMainForm) OnButton1Click(sender vcl.IObject) {
	About.SetCaption(f.test)
	About.Show()
}

