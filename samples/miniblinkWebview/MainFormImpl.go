// 在这里写你的事件

package main

import (
	"strings"

	"github.com/ying32/govcl/vcl/types"

	"github.com/ying32/govcl/vcl/rtl"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/miniblink"
)

//::private::
type TMainFormFields struct {
	web *miniblink.TMiniBlinkWebview
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.web = miniblink.NewMiniBlinkWebview(f.PnlWeb.Handle())
	f.web.Show(true)

	//if rtl.LcLLoaded() {
	//	f.web.LoadFile("F:\\Program Files (x86)\\Golang\\go\\doc\\effective_go.html")
	//} else {
	//	f.web.LoadFileW("F:\\Program Files (x86)\\Golang\\go\\doc\\effective_go.html")
	//}
	//f.web.LoadHTMLW("<html><head><title>标题</title></head><body><p>这是一个测试</p></body></html>")
	//URL := "https://gitee.com/ying32"
	//if rtl.LcLLoaded() {
	//	//f.web.LoadURL(URL)
	//
	//} else {
	//	//f.web.LoadURLW(URL)
	//}
}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {
	if f.web != nil && f.web.IsValid() {
		f.web.Free()
	}
}

func (f *TMainForm) OnPnlWebResize(sender vcl.IObject) {
	if f.web != nil && f.web.IsValid() {
		f.web.MoveWindow(0, 0, int(f.Width()), int(f.Height()))
	}
}

func (f *TMainForm) OnBtnReloadClick(sender vcl.IObject) {
	f.web.Reload()
}

func (f *TMainForm) OnBtnNavClick(sender vcl.IObject) {
	text := strings.TrimSpace(f.EdtURL.Text())
	if text == "" {
		vcl.ShowMessage("请输入一个URL！")
		return
	}
	f.web.LoadURL(text)
}

func (f *TMainForm) OnBtnLoadFromFileClick(sender vcl.IObject) {
	if f.OpenDialog1.Execute() {
		fileName := f.OpenDialog1.FileName()
		if rtl.LcLLoaded() {
			f.web.LoadFile(fileName)
		} else {
			f.web.LoadFileW(fileName)
		}
	}
}

func (f *TMainForm) OnBtnBackClick(sender vcl.IObject) {
	f.web.GoBack()
}

func (f *TMainForm) OnBtnForwardClick(sender vcl.IObject) {
	f.web.GoForward()
}

func (f *TMainForm) OnBtnLoadFromStringClick(sender vcl.IObject) {
	result := HTMLForm.ShowModal()
	if result == types.MrOk {
		if rtl.LcLLoaded() {
			f.web.LoadHTML(HTMLForm.Memo1.Text())
		} else {
			f.web.LoadHTMLW(HTMLForm.Memo1.Text())
		}
	}
}
