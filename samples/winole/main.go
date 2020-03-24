// +build windows

package main

import (
	"fmt"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	Btn1 *vcl.TButton
}

var mainForm *TMainForm

func (f *TMainForm) OnFormCreate(object vcl.IObject) {
	f.SetCaption("WinOLE")
	f.ScreenCenter()

	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 100, 30)
	f.Btn1.SetCaption("打开记事本")
}

func (f *TMainForm) OnBtn1Click(object vcl.IObject) {
	ole.CoInitializeEx(0, 0)
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		fmt.Println(err)
		return
	}
	shell, _ := unknown.QueryInterface(ole.IID_IDispatch)
	shell.CallMethod("Run", "notepad")
}

func main() {
	vcl.RunApp(&mainForm)
}
