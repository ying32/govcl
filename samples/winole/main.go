// +build windows

package main

import (
	"fmt"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.Initialize()
	mainFrom := vcl.Application.CreateForm()
	mainFrom.SetCaption("WinOLE")
	mainFrom.ScreenCenter()

	btn := vcl.NewButton(mainFrom)
	btn.SetParent(mainFrom)
	btn.SetBounds(10, 10, 100, 30)
	btn.SetCaption("打开记事本")
	btn.SetOnClick(func(sender vcl.IObject) {
		//
		ole.CoInitializeEx(0, 0)
		defer ole.CoUninitialize()

		unknown, err := oleutil.CreateObject("WScript.Shell")
		if err != nil {
			fmt.Println(err)
			return
		}
		shell, _ := unknown.QueryInterface(ole.IID_IDispatch)
		shell.CallMethod("Run", "notepad")
	})

	vcl.Application.Run()
}
