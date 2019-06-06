// 在这里写你的事件

package main

import (
	"fmt"
	"strings"

	"github.com/ying32/govcl/vcl/types/keys"
	"github.com/ying32/govcl/vcl/win"

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
	fmt.Println("OnFormCreate:", win.GetCurrentThreadId())
	f.web = miniblink.NewMiniBlinkWebview(f.PnlWeb.Handle())
	f.web.Show(true)
	f.web.OnTitleChanged = func(sender *miniblink.TMiniBlinkWebview, title string) {
		//fmt.Println("标题改变:", title)
		fmt.Println("onTitleChanged:", win.GetCurrentThreadId())
		f.SetCaption(title + " - ying32")
	}

	f.web.OnURLChanged = func(sender *miniblink.TMiniBlinkWebview, url string) {
		fmt.Println("URL改变:", url)
	}

	f.web.OnDocumentReady = func(sender *miniblink.TMiniBlinkWebview) {
		fmt.Println("文档已准备。")
	}

	f.web.OnLoadingFinish = func(sender *miniblink.TMiniBlinkWebview, url string, result miniblink.WkeLoadingResult, failedReason string) {
		fmt.Println("加载完成：", url, ", result:", result, ", failedReason:", failedReason)
	}
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
	f.web.Webview.Reload()
}

func (f *TMainForm) OnBtnLoadFromFileClick(sender vcl.IObject) {
	if f.OpenDialog1.Execute() {
		fileName := f.OpenDialog1.FileName()
		f.web.Webview.LoadFile(fileName)
	}
}

func (f *TMainForm) OnBtnLoadFromStringClick(sender vcl.IObject) {
	result := HTMLForm.ShowModal()
	if result == types.MrOk {
		f.web.Webview.LoadHTML(HTMLForm.Memo1.Text())
	}
}

//func (f *TMainForm) OnFormKeyPress(sender vcl.IObject, key *types.Char) {
//
//}

func (f *TMainForm) OnActGoBackExecute(sender vcl.IObject) {
	f.web.Webview.GoBack()
}

func (f *TMainForm) OnActGoBackUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.web != nil && f.web.IsValid() && f.web.Webview.CanGoBack())
}

func (f *TMainForm) OnActGoForwardExecute(sender vcl.IObject) {
	f.web.Webview.GoForward()
}

func (f *TMainForm) OnActGoForwardUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.web != nil && f.web.IsValid() && f.web.Webview.CanGoForward())
}

func (f *TMainForm) OnActNavExecute(sender vcl.IObject) {
	text := strings.TrimSpace(f.EdtURL.Text())
	if text == "" {
		vcl.ShowMessage("请输入一个URL！")
		return
	}
	f.web.Webview.LoadURL(text)
}

func (f *TMainForm) OnActNavUpdate(sender vcl.IObject) {
	//vcl.ActionFromObj(sender).SetEnabled(f.web != nil && f.web.IsValid() && f.web.Webview.CanGoForward())
}

func (f *TMainForm) OnFormKeyDown(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
	if *key == keys.VkReturn {
		f.ActNav.Execute()
		*key = 0
	}
}

func (f *TMainForm) OnBtnExecJSClick(sender vcl.IObject) {
	result := JsForm.ShowModal()
	if result == types.MrOk && f.web.IsValid() {
		val := f.web.Webview.RunJS(JsForm.Memo1.Text())
		switch val.TypeOf() {
		case miniblink.JSTYPE_NUMBER:
			fmt.Println("Number")
		case miniblink.JSTYPE_STRING:
			fmt.Println("String")
			fmt.Println("返回值：", f.web.Webview.GlobalExec().ToTempString(val))
		case miniblink.JSTYPE_BOOLEAN:
			fmt.Println("Boolean")
		case miniblink.JSTYPE_OBJECT:
			fmt.Println("Object")
		case miniblink.JSTYPE_FUNCTION:
			fmt.Println("Function")
		case miniblink.JSTYPE_UNDEFINED:
			fmt.Println("Undefined")
		case miniblink.JSTYPE_ARRAY:
			fmt.Println("Array")
		case miniblink.JSTYPE_NULL:
			fmt.Println("Null")
		default:
			fmt.Println("未知")

		}

	}
}
