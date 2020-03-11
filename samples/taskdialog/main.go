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
	Btn1 *vcl.TButton
}

var MainForm *TMainForm

func main() {
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&MainForm, true)
	vcl.Application.Run()
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.SetCaption("taskDialog演示")

	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetCaption("TaskDialog")
	f.Btn1.SetLeft(10)
	f.Btn1.SetTop(10)

}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {

}

func (f *TMainForm) OnBtn1Click(sender vcl.IObject) {
	taskdlg := *vcl.NewTaskDialog(f)
	defer taskdlg.Free()
	taskdlg.SetTitle("确认移除")
	taskdlg.SetCaption("询问")
	taskdlg.SetText("移除选择的项目？")
	taskdlg.SetExpandButtonCaption("展开按钮标题")
	taskdlg.SetExpandedText("展开的文本")

	taskdlg.SetFooterText("底部文本")

	rd := vcl.AsTaskDialogRadioButtonItem(taskdlg.RadioButtons().Add())
	rd.SetCaption("单选按钮1")
	rd = vcl.AsTaskDialogRadioButtonItem(taskdlg.RadioButtons().Add())
	rd.SetCaption("单选按钮2")
	rd = vcl.AsTaskDialogRadioButtonItem(taskdlg.RadioButtons().Add())
	rd.SetCaption("单选按钮3")

	taskdlg.SetCommonButtons(0) //rtl.Include(0, 0))
	btn := vcl.AsTaskDialogButtonItem(taskdlg.Buttons().Add())
	btn.SetCaption("移除")
	btn.SetModalResult(types.MrYes)

	btn = vcl.AsTaskDialogButtonItem(taskdlg.Buttons().Add())
	btn.SetCaption("保持")
	btn.SetModalResult(types.MrNo)

	if !rtl.LcLLoaded() {
		f := taskdlg.Flags()
		taskdlg.SetFlags(f.Include(types.TfShowProgressBar, types.TfShowMarqueeProgressBar))
		taskdlg.ProgressBar().SetPosition(50)
	}

	if rtl.LcLLoaded() {
		taskdlg.SetMainIcon(types.TdiQuestion)
	} else {
		taskdlg.SetMainIcon(types.TdiInformation)
	}
	if taskdlg.Execute() {
		if taskdlg.ModalResult() == types.MrYes {
			vcl.ShowMessage("项目已移除")

			fmt.Println(taskdlg.RadioButton().Caption())
		}
	}
}
