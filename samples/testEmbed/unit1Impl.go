package main

import (
	"github.com/ying32/govcl/vcl"
)

// 本例程需要go1.16版本支持。

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	vcl.ShowMessage("hello")
}
