//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
)
import _ "github.com/ying32/govcl/pkgs/winappres"

type TMainForm struct {
	*vcl.TForm
	btn  *vcl.TButton
	btn2 *vcl.TButton
	btn3 *vcl.TButton
}

var mainForm *TMainForm

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&mainForm)
	vcl.Application.Run()
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.btn = vcl.NewButton(f)
	f.btn.SetParent(f)
	f.btn.SetCaption("GetText")
	f.btn.SetOnClick(f.onBtnClick)

	f.btn2 = vcl.NewButton(f)
	f.btn2.SetParent(f)
	f.btn2.SetTop(40)
	f.btn2.SetWidth(120)
	f.btn2.SetCaption("GetTextBuffer")
	f.btn2.SetOnClick(f.onBtn2Click)

	f.btn3 = vcl.NewButton(f)
	f.btn3.SetParent(f)
	f.btn3.SetTop(80)
	f.btn3.SetWidth(120)
	f.btn3.SetCaption("SetText")
	f.btn3.SetOnClick(f.onBtn3Click)
}

func (f *TMainForm) onBtnClick(sender vcl.IObject) {
	str := vcl.Clipboard.AsText()
	fmt.Println("len1:", len(str))
}

func (f *TMainForm) onBtn2Click(sender vcl.IObject) {
	str := ""
	vcl.Clipboard.GetTextBuf(&str, 1000) // buff不够长
	fmt.Println("len2:", len(str))
}

func (f *TMainForm) onBtn3Click(sender vcl.IObject) {

	vcl.Clipboard.SetAsText("afdsdf")
}
