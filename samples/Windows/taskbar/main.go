package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/win"

	"github.com/ying32/govcl/vcl/types"

	tsbar "github.com/ying32/govcl/pkgs/wintaskbar"
	"github.com/ying32/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	Button1  *vcl.TButton
	Button2  *vcl.TButton
	taskBar  *tsbar.WinTaskBar
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

	f.taskBar = tsbar.NewWinTaskBar(f.Handle())
	f.taskBar.SetOnThumbButtonClick(f.onThumbButtonClick)

	loadIcon := func(name string) types.HICON {
		return win.LoadIcon2(rtl.MainInstance(), name)
	}

	// button只能一次性添加的，然后不能再添加和删除了，只能更新，这是ms官方的说明

	btn, _ := f.taskBar.AddButton()
	btn.SetHint("上一曲")
	btn.SetIcon(loadIcon("TASKBTN_PREV"))

	btn, _ = f.taskBar.AddButton()
	btn.SetHint("播放")
	btn.SetIcon(loadIcon("TASKBTN_PLAY"))

	btn, _ = f.taskBar.AddButton()
	btn.SetHint("暂停")
	btn.SetIcon(loadIcon("TASKBTN_PAUSE"))
	btn.SetFlags(tsbar.Hidden)

	btn, _ = f.taskBar.AddButton()
	btn.SetHint("下一曲")
	btn.SetIcon(loadIcon("TASKBTN_NEXT"))

}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {
	if f.taskBar != nil {
		f.taskBar.Free()
		f.taskBar = nil
	}
}

func (f *TMainForm) onThumbButtonClick(index uint16) {
	switch index {
	case 0: // 上一曲
		fmt.Println("上一曲")
	case 1: // 播放
		fmt.Println("播放")
		f.taskBar.Buttons()[1].SetFlags(tsbar.Hidden)
		f.taskBar.Buttons()[2].SetFlags(tsbar.Enabled)
	case 2: // 暂停
		fmt.Println("暂停")
		f.taskBar.Buttons()[2].SetFlags(tsbar.Hidden)
		f.taskBar.Buttons()[1].SetFlags(tsbar.Enabled)
	case 3: // 下一曲
		fmt.Println("下一曲")
	}
}

func (f *TMainForm) onButton1Click(sender vcl.IObject) {
	// 设置后，鼠标悬停会显示提示
	f.taskBar.SetThumbnailTooltip("HELLO!")
	// 进度状态
	f.taskBar.SetProgressState(tsbar.Error)

}

func (f *TMainForm) onButton2Click(sender vcl.IObject) {

	// 叠加的icon
	f.taskBar.SetOverlayIcon(f.Icon().Handle(), "描述啊。。。")
}

func (f *TMainForm) doTimer(sender vcl.IObject) {
	f.progress++
	if f.progress > 100 {
		f.progress = 0
	}
	// 进度值
	f.taskBar.SetProgressValue(f.progress, 100)
}
