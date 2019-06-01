package main

import (
	"fmt"

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
	taskdlg.SetTitle("Confirm removal")
	taskdlg.SetCaption("Confirm")
	taskdlg.SetText("Remove selected item?")
	taskdlg.SetExpandButtonCaption("expand button caption")
	taskdlg.SetExpandedText("expanded text")

	taskdlg.SetFooterText("Footer Text")

	rd := vcl.TaskDialogRadioButtonItemFromObj(taskdlg.RadioButtons().Add())
	rd.SetCaption("RadioButton1")
	rd = vcl.TaskDialogRadioButtonItemFromObj(taskdlg.RadioButtons().Add())
	rd.SetCaption("RadioButton2")
	rd = vcl.TaskDialogRadioButtonItemFromObj(taskdlg.RadioButtons().Add())
	rd.SetCaption("RadioButton2")

	taskdlg.SetCommonButtons(0) //rtl.Include(0, 0))
	btn := vcl.TaskDialogButtonItemFromObj(taskdlg.Buttons().Add())
	btn.SetCaption("Remove")
	btn.SetModalResult(types.MrYes)

	btn = vcl.TaskDialogButtonItemFromObj(taskdlg.Buttons().Add())
	btn.SetCaption("Keep")
	btn.SetModalResult(types.MrNo)

	if !rtl.LcLLoaded() {
		f := taskdlg.Flags()
		taskdlg.SetFlags(rtl.Include(f, types.TfShowProgressBar, types.TfShowMarqueeProgressBar))
		taskdlg.ProgressBar().SetPosition(50)
	}

	if rtl.LcLLoaded() {
		taskdlg.SetMainIcon(types.TdiQuestion)
	} else {
		taskdlg.SetMainIcon(types.TdiInformation)
	}
	if taskdlg.Execute() {
		if taskdlg.ModalResult() == types.MrYes {
			vcl.ShowMessage("Item removed")

			fmt.Println(taskdlg.RadioButton().Caption())
		}
	}
}
