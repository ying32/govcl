// 在这里写你的事件

package main

import (
	"fmt"
	"strings"

	"github.com/ying32/govcl/vcl/types"

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
	f.web.OnTitleChanged = f.onWkeTitleChanged
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

func (f *TMainForm) onWkeTitleChanged(sender *miniblink.TMiniBlinkWebview, title string) {
	fmt.Println("标题改变")
	f.SetCaption(title + " - ying32")
}

func (f *TMainForm) OnBtnReloadClick(sender vcl.IObject) {
	f.web.Webview.Reload()
}

func (f *TMainForm) OnBtnNavClick(sender vcl.IObject) {
	text := strings.TrimSpace(f.EdtURL.Text())
	if text == "" {
		vcl.ShowMessage("请输入一个URL！")
		return
	}
	f.web.Webview.LoadURL(text)
}

func (f *TMainForm) OnBtnLoadFromFileClick(sender vcl.IObject) {
	if f.OpenDialog1.Execute() {
		fileName := f.OpenDialog1.FileName()
		f.web.Webview.LoadFile(fileName)
	}
}

func (f *TMainForm) OnBtnBackClick(sender vcl.IObject) {
	f.web.Webview.GoBack()
}

func (f *TMainForm) OnBtnForwardClick(sender vcl.IObject) {
	f.web.Webview.GoForward()
}

func (f *TMainForm) OnBtnLoadFromStringClick(sender vcl.IObject) {
	result := HTMLForm.ShowModal()
	if result == types.MrOk {
		f.web.Webview.LoadHTML(HTMLForm.Memo1.Text())
	}
}
