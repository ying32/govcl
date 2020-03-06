// +build windows

package main

import (
	"fmt"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
)

type TMainForm struct {
	*vcl.TForm
	Taskbar1 *vcl.TTaskbar
	Timer1   *vcl.TTimer
	Btn1     *vcl.TButton
	Btn2     *vcl.TButton
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&mainForm)
	vcl.Application.Run()
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	fmt.Println("FromCreate")
	fmt.Println("MainForm.Handle:", vcl.Application.MainForm().Handle())
	fmt.Println("Application.Handle:", vcl.Application.Handle())
	f.Taskbar1 = vcl.NewTaskbar(f)
	f.Taskbar1.SetToolTip("这是提示")
	f.Taskbar1.SetOverlayHint("这是一个提示")
	f.Taskbar1.SetProgressMaxValue(100)
	f.Taskbar1.SetProgressValue(50)
	f.Taskbar1.OverlayIcon().LoadFromFile("pause.ico")

	f.Taskbar1.SetTabProperties(rtl.Include(0, types.CustomizedPreview))
	//f.Taskbar1.SetProgressState(types.Error)
	f.Taskbar1.SetProgressState(types.Normal)
	f.Taskbar1.SetOnThumbButtonClick(f.OnTaskbar1ThumbButtonClick)
	f.Taskbar1.SetOnThumbPreviewRequest(f.OnTaskbar1ThumbPreviewRequest)
	f.Taskbar1.SetOnWindowPreviewItemRequest(f.OnTaskbar1WindowPreviewItemRequest)

	// 预览窗口的修改
	//f.Taskbar1.PreviewClipRegion().SetWidth()

	// buttons
	tbtn := f.Taskbar1.TaskBarButtons().Add()
	tbtn.SetHint("上一曲")
	tbtn.Icon().LoadFromFile("prev.ico")
	//tbtn.Icon().LoadFromResourceName(rtl.MainInstance(), "TASKBTN_PREV")

	tbtn = f.Taskbar1.TaskBarButtons().Add()
	tbtn.SetHint("播放")
	tbtn.Icon().LoadFromFile("play.ico")
	//tbtn.Icon().LoadFromResourceName(rtl.MainInstance(), "TASKBTN_PLAY")

	tbtn = f.Taskbar1.TaskBarButtons().Add()
	tbtn.SetHint("下一曲")
	tbtn.Icon().LoadFromFile("next.ico")
	//tbtn.Icon().LoadFromResourceName(rtl.MainInstance(), "TASKBTN_NEXT")

	f.Timer1 = vcl.NewTimer(f)
	f.Timer1.SetOnTimer(f.OnTimer1Timer)
	//f.Timer1.SetInterval(1000)
	f.Timer1.SetEnabled(true)

	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 88, 28)
	f.Btn1.SetCaption("错误")
	f.Btn1.SetOnClick(f.OnButtonClick)

	f.Btn2 = vcl.NewButton(f)
	f.Btn2.SetParent(f)
	f.Btn2.SetBounds(10, f.Btn1.Top()+f.Btn1.Height()+10, 88, 28)
	f.Btn2.SetCaption("正常")
	f.Btn2.SetOnClick(f.OnButtonClick)
}

func (f *TMainForm) OnTimer1Timer(sender vcl.IObject) {
	if f.Taskbar1.ProgressValue() >= 100 {
		f.Taskbar1.SetProgressValue(0)
	}
	f.Taskbar1.SetProgressValue(f.Taskbar1.ProgressValue() + 1)
}

func (f *TMainForm) OnTaskbar1ThumbButtonClick(sender vcl.IObject, aButtonID int32) {
	fmt.Println("click:", sender.ClassName())
	switch aButtonID {
	case 0:
		fmt.Println("prev")
	case 1:
		fmt.Println("play")
	case 2:
		fmt.Println("next")
	}
}

func (f *TMainForm) OnTaskbar1ThumbPreviewRequest(sender vcl.IObject, aPreviewHeight, aPreviewWidth int32, previewBitmap *vcl.TBitmap) {
	fmt.Println("OnTaskbar1ThumbPreviewRequest")
	previewBitmap.LoadFromFile("gopher.bmp")
}

func (f *TMainForm) OnTaskbar1WindowPreviewItemRequest(sender vcl.IObject, position *types.TPoint, previewBitmap *vcl.TBitmap) {
	fmt.Println("OnTaskbar1WindowPreviewItemRequest")
}

func (f *TMainForm) OnButtonClick(sender vcl.IObject) {
	if sender.Equals(f.Btn1) {
		f.Taskbar1.SetProgressState(types.Error)
	} else if sender.Equals(f.Btn2) {
		f.Taskbar1.SetProgressState(types.Normal)
	}
}
