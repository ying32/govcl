// 在这里写你的事件

package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TMainFormFields struct {
	webview  *vcl.TMiniWebview
	isLoaded bool
}

const (
	webAddr = ":58988"
)

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.webview = vcl.NewMiniWebview(f)
	f.webview.SetParent(f.PnlWebview)
	f.webview.SetAlign(types.AlClient)
	f.webview.Navigate("about:blank")

	f.setToolBtnClickEvent()

	http.HandleFunc("/", f.httpHandleFunc)
	go func() {
		http.ListenAndServe(webAddr, nil)
	}()
}

func (f *TMainForm) onToolBtnClick(sender vcl.IObject) {
	btn := vcl.AsButton(sender)

	setSelText := func(s string) {
		f.MemoEditor.SetSelText(fmt.Sprintf("%s%s%s", s, f.MemoEditor.SelText(), s))
	}
	setSelText2 := func(s string) {
		f.MemoEditor.SetSelText(fmt.Sprintf("%s %s", s, f.MemoEditor.SelText()))
	}

	switch btn.Tag() {
	case 1:
		f.MemoEditor.Undo()
	case 2:
	case 3:
		setSelText("**")
	case 4:
		setSelText("~~")
	case 5:
		setSelText("*")
	case 6:
		setSelText2(">")

	case 10:
		setSelText2("#")
	case 11:
		setSelText2("##")
	case 12:
		setSelText2("###")
	case 13:
		setSelText2("####")
	case 14:
		setSelText2("#####")
	case 15:
		setSelText2("######")
	case 16:
		setSelText2("-")
	case 17:
		setSelText2("1.")
	case 18:
		setSelText2("\r\n------\r\n")
	case 19:
		f.MemoEditor.SetSelText(fmt.Sprintf("[%s]()", f.MemoEditor.SelText()))
	case 20:
		f.MemoEditor.SetSelText(fmt.Sprintf("[%s][]", f.MemoEditor.SelText()))
	case 21:
		f.MemoEditor.SetSelText(fmt.Sprintf("![%s]()", f.MemoEditor.SelText()))

	case 22:
		setSelText("\r\n```\r\n")
	}

}

func (f *TMainForm) setToolBtnClickEvent() {
	var i int32
	for i = 0; i < f.ToolBar1.ControlCount(); i++ {
		ctl := f.ToolBar1.Controls(i)
		if ctl.Is().ToolButton() {
			btn := vcl.AsToolButton(ctl)
			if btn.Style() == types.TbsButton {
				btn.SetOnClick(f.onToolBtnClick)
			}
		}
	}
}

func (f *TMainForm) httpHandleFunc(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html;charset=utf-8")
	var buf *bytes.Buffer
	vcl.ThreadSync(func() {
		buf = makeHtmlPage(map[string]interface{}{"MdData": f.MemoEditor.Text()})
	})
	if buf != nil {
		writer.Write(buf.Bytes())
	}
}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {

}

func (f *TMainForm) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {

}

func (f *TMainForm) OnFormResize(sender vcl.IObject) {
	if f.PnlWebview != nil && f.PnlEditor != nil {
		wHalf := (f.Width() - f.Splitter1.Width()) / 2
		f.PnlWebview.SetWidth(wHalf)
	}
}

func (f *TMainForm) OnMemoEditorChange(sender vcl.IObject) {
	if !f.isLoaded {
		f.webview.Navigate("http://127.0.0.1" + webAddr)
	} else {
		f.webview.Refresh()
	}
	f.MemoEditor.SetFocus()
}
