// 在这里写你的事件

package main

import (
	"fmt"
	"strings"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
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

	// 不知道原因，有些不设置有些反而好些。。。
	//SetIEVersion(f.webView)

	f.webView.SetOnTitleChange(f.OnWebTitleChange)
	f.webView.SetOnJSExternal(f.OnWebJsExternal)

	//f.webView.Navigate("https://github.com/ying32/govcl")

	URL := "file:///" + strings.Replace(rtl.ExtractFilePath(vcl.Application.ExeName()), "\\", "/", -1) + "test.html"
	f.EdtURL.SetText(URL)
	f.webView.Navigate(URL)

	//vcl.Application.SetOnMessage(f.onApplicationMessage)
	//f.Button3.Click()
}

//func (f *TMainForm) onApplicationMessage(msg *types.TMsg, handled *bool) {
//	if msg.Hwnd != f.Handle() {
//		if msg.Message == messages.WM_KEYDOWN {
//			fmt.Println("按下", msg.WParam)
//			f.webView.Perform(msg.Message, msg.WParam, int(msg.LParam))
//			*handled = false
//
//		}
//	}
//}

func (f *TMainForm) OnBtnGoForwardClick(sender vcl.IObject) {
	// Windows下失效了，不知道原因。。。
	f.webView.GoForward()
}

func (f *TMainForm) OnBtnGoBackClick(sender vcl.IObject) {
	// Windows下失效了，不知道原因。。。
	f.webView.GoBack()
}

func (f *TMainForm) OnBtnRefreshClick(sender vcl.IObject) {
	// Windows下失效了，不知道原因。。。
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

func (f *TMainForm) OnWebJsExternal(sender vcl.IObject, funcName, args string, retVal *string) {
	fmt.Println("js call: name:", funcName, ", args:", args)
	switch funcName {
	case "testCall":
		*retVal = "结果"
	case "runScript":
		f.webView.ExecuteJS(args)
	}

}

func (f *TMainForm) OnButton2Click(sender vcl.IObject) {
	f.webView.ExecuteJS("alert('执行脚本。');")
}

func (f *TMainForm) OnButton3Click(sender vcl.IObject) {
	f.EdtURL.SetText("about:blank")
	f.webView.LoadHTML(`
    <html>
      <head>
        <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
		<meta http-equiv="content-type" content="text/html;charset=utf-8">       
        <title>从字符串加载HTML</title>
      </head>
      <body>
         <p>这是一个从字符串加载的HTML。</p>
      </body> 
    </html>
`)
}
