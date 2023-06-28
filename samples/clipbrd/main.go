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

	"github.com/ying32/govcl/vcl/types"

	"github.com/ying32/govcl/vcl"
)
import _ "github.com/ying32/govcl/pkgs/winappres"

type TMainForm struct {
	*vcl.TForm
	btn  *vcl.TButton
	btn2 *vcl.TButton
	btn3 *vcl.TButton
	btn4 *vcl.TButton
	btn5 *vcl.TButton
	img  *vcl.TImage
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

	f.btn4 = vcl.NewButton(f)
	f.btn4.SetParent(f)
	f.btn4.SetTop(f.btn3.Top() + 50)
	f.btn4.SetWidth(120)
	f.btn4.SetCaption("SetBmp")
	f.btn4.SetOnClick(f.onBtn4Click)

	f.btn5 = vcl.NewButton(f)
	f.btn5.SetParent(f)
	f.btn5.SetTop(f.btn4.Top() + 50)
	f.btn5.SetWidth(120)
	f.btn5.SetCaption("GetBmp")
	f.btn5.SetOnClick(f.onBtn5Click)

	f.img = vcl.NewImage(f)
	f.img.SetParent(f)
	f.img.SetTop(f.btn5.Top() + 30)
	f.img.SetWidth(200)
	f.img.SetHeight(200)
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

func (f *TMainForm) onBtn4Click(sender vcl.IObject) {

	//ss := vcl.NewStringList()
	//defer ss.Free()
	//vcl.Clipboard.SupportedFormats(ss)
	//vcl.ShowMessage(ss.Text())
	//return

	mem := vcl.NewMemoryStream()
	defer mem.Free()
	// only bmp
	mem.LoadFromFile("bg.bmp")
	mem.SetPosition(0)

	// 注册自定义的格式
	//vcl.RegisterClipboardFormat()

	// 预定义格式
	format := vcl.PredefinedClipboardFormat(types.PcfBitmap)
	fmt.Println("format:", format)

	if !vcl.Clipboard.SetFormat(format, mem) {
		vcl.ShowMessage("设置格式失败")
	} else {
		vcl.ShowMessage("设置成功")
	}
}

func (f *TMainForm) onBtn5Click(sender vcl.IObject) {
	if !vcl.Clipboard.HasPictureFormat() {
		return
	}
	bmpFormat := vcl.Clipboard.FindPictureFormatID()
	mem := vcl.NewMemoryStream()
	defer mem.Free()
	if vcl.Clipboard.GetFormat(bmpFormat, mem) {
		mem.SetPosition(0)
		f.img.Picture().LoadFromStream(mem)
	}
}
