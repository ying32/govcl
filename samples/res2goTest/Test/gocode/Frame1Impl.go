// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"github.com/ying32/govcl/vcl"
)

//::private::
type TFrame1Fields struct {
}

func (f *TFrame1) OnButton1Click(sender vcl.IObject) {
	vcl.ShowMessage("Frame1")
}
