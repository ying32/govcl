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
	Button1  *vcl.TButton
	Button2  *vcl.TButton
	taskbar  *wintaskbar.TWinTaskBar
	timer    *vcl.TTimer
	progress uint64
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

	f.timer = vcl.NewTimer(f)
	f.timer.SetInterval(200)
	f.timer.SetEnabled(true)
	f.timer.SetOnTimer(f.doTimer)

	f.taskbar = wintaskbar.NewWinTaskBar(f.Handle())
	f.taskbar.SetOnThumbButtonClick(func(index uint16) {
		fmt.Println("index: ", index)
	})
	h := types.HICON(0)

	// button只能一次性添加的，然后不能再添加和删除了，只能更新，这是ms官方的说明
	// 这里的api还要重构下，方便使用
	f.taskbar.AddButton("test1", h, wintaskbar.Enabled)
	f.taskbar.AddButton("test2", h, wintaskbar.Enabled)
	f.taskbar.AddButton("test3", h, wintaskbar.Enabled)

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
	f.taskbar.SetProgressState(wintaskbar.Error)

}

func (f *TMainForm) onButton2Click(sender vcl.IObject) {

	// 叠加的icon
	f.taskbar.SetOverlayIcon(f.Icon().Handle(), "描述啊。。。")
}

func (f *TMainForm) doTimer(sender vcl.IObject) {
	f.progress++
	if f.progress > 100 {
		f.progress = 0
	}
	// 进度值
	f.taskbar.SetProgressValue(f.progress, 100)
}
