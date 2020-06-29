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
	Button2 *vcl.TButton
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

	f.Button1 = vcl.NewButton(f)
	f.Button1.SetParent(f)
	f.Button1.AnchorVerticalCenterTo(f)
	f.Button1.AnchorHorizontalCenterTo(f)
	f.Button1.SetOnClick(f.onButton1Click)
	f.Button1.SetCaption("test")
	f.Button1.SetWidth(120)

	f.Button2 = vcl.NewButton(f)
	f.Button2.SetParent(f)

	f.Button2.SetOnClick(f.onButton2Click)
	f.Button2.SetCaption("test2")
	f.Button2.SetWidth(120)

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

func (f *TMainForm) onButton1Click(sender vcl.IObject) {
	// 设置后，鼠标悬停会显示提示
	f.taskbar.SetThumbnailTooltip("HELLO!")
	// 进度状态
	f.taskbar.SetProgressState(wintaskbar.TBPF_ERROR)
	// 进度值
	f.taskbar.SetProgressValue(50, 100)

}

func (f *TMainForm) onButton2Click(sender vcl.IObject) {

	// 叠加的icon
	f.taskbar.SetOverlayIcon(f.Icon().Handle(), "描述啊。。。")
}
