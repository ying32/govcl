package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/types"

	"github.com/ying32/govcl/pkgs/wintaskbar"
	"github.com/ying32/govcl/vcl"

	_ "github.com/ying32/govcl/pkgs/winappres"
)

type TMainForm struct {
	*vcl.TForm
	Button1 *vcl.TButton
	taskbar *wintaskbar.TWinTaskBar
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.RunApp(&mainForm)
}

// --------------MainForm -----------------
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

	f.SetCaption("Hello")
	f.SetWidth(300)
	f.SetHeight(200)
	f.ScreenCenter()

	f.taskbar = wintaskbar.NewWinTaskBar(f.Handle())
	f.taskbar.SetOnThumbButtonClick(func(index uint16) {
		fmt.Println("index: ", index)
	})
	h := types.HICON(0)
	f.taskbar.AddButton("test1", h, wintaskbar.THBF_ENABLED)
	f.taskbar.AddButton("test2", h, wintaskbar.THBF_ENABLED)
	f.taskbar.AddButton("test3", h, wintaskbar.THBF_ENABLED)
}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {
	if f.taskbar != nil {
		f.taskbar.Free()
		f.taskbar = nil
	}
}
