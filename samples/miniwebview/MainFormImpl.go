// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TMainFormFields struct {
	webView *vcl.TMiniWebview
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	fmt.Println("ok")
	f.ScreenCenter()
	f.webView = vcl.NewMiniWebview(f)
	f.webView.SetParent(f.PnlWebview)
	f.webView.SetAlign(types.AlClient)
	f.webView.Navigate("about:blank")
	f.webView.SetOnTitleChange(f.OnWebTitleChange)
}

func (f *TMainForm) OnBtnGoForwardClick(sender vcl.IObject) {
	f.webView.GoForward()
}

func (f *TMainForm) OnBtnGoBackClick(sender vcl.IObject) {
	f.webView.GoBack()
}

func (f *TMainForm) OnBtnRefreshClick(sender vcl.IObject) {
	f.webView.Refresh()
}

func (f *TMainForm) OnButton1Click(sender vcl.IObject) {
	if f.EdtURL.Text() == "" {
		return
	}
	f.webView.Navigate(f.EdtURL.Text())
}

func (f *TMainForm) OnWebTitleChange(sender vcl.IObject, text string) {
	f.SetCaption(text + " - ying32")
}
