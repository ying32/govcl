package main

import (
	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
}

func (f *TForm1) OnButton2Click(sender vcl.IObject) {
	f.ScaleForCurrentDpi()
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	f.ScaleForPPI(125)
}
