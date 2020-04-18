package main

import (
	"fmt"

	"strconv"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(object vcl.IObject) {
	fmt.Println("OnForm1Create被调用了:", object.Instance())
	vcl.Application.SetTitle("我是新标题")

	f.SetKeyPreview(true)
	f.Edit2.SetText("1")
	f.Edit2.SetNumbersOnly(true)

	f.ActFileNew.SetCaption("我是一个共享的动作")

}

func (f *TForm1) OnButton1Click(object vcl.IObject) {
	fmt.Println("f:", f.Button1.Name())
	fmt.Println("窗口1的Button1单击")
	img := vcl.NewJPEGImage()
	defer img.Free()
	img.LoadFromFile("111.jpg")
}

func (f *TForm1) OnButton2Click(object vcl.IObject) {
	fmt.Println("f:", f.Button2.Name(), "xxx:", object.ClassName())
	fmt.Println("窗口1的Button2单击")
	result := Form2.ShowModal()
	if result == types.MrOk {
		vcl.ShowMessage("Form2返回了OK")
	} else if result == types.MrClose || result == types.MrNone {
		vcl.ShowMessage("Form2返回了Close")
	} else if result == types.MrCancel {
		vcl.ShowMessage("Form2返回了Cancel")
	}
}

func (f *TForm1) OnActExitExecute(sender vcl.IObject) {
	vcl.Application.Terminate()
}

func (f *TForm1) OnActFileNewExecute(sender vcl.IObject) {
	vcl.ShowMessage("ActFileNew Execute.")
}

func (f *TForm1) OnCheckBox1Click(sender vcl.IObject) {
	f.Button1.SetEnabled(f.CheckBox1.Checked())
}

func (f *TForm1) OnFormClick(sender vcl.IObject) {
	fmt.Println("OnForm1Click")
}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {
	fmt.Println("OnForm1Destroy")
}

func (f *TForm1) OnFormKeyDown(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
	fmt.Println("OnForm1KeyDown:", *key, ", ", shift)
}

func (f *TForm1) OnFormMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	fmt.Println("OnForm1MouseDown:", button, ", ", shift, ", ", x, ", ", y)
}

func (f *TForm1) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	fmt.Println("OnForm1CloseQuery")
	//*canClose = false
}

func (f *TForm1) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	fmt.Println("OnForm1Close")
	//*canClose = false
}

func (f *TForm1) OnApplicationException(sender vcl.IObject, e *vcl.Exception) {
	fmt.Println("OnApplicationException")
	vcl.ShowMessage(e.Message())
}

func (f *TForm1) OnApplicationMinimize(sender vcl.IObject) {
	fmt.Println("OnApplicationMinimize")
}

func (f *TForm1) OnApplicationRestore(sender vcl.IObject) {
	fmt.Println("OnApplicationRestore")
}

func (f *TForm1) OnComboBox1Change(sender vcl.IObject) {
	fmt.Println("OnComboBox1Change: ", f.ComboBox1.ItemIndex())
}

func (f *TForm1) OnImageButton1Click(sender vcl.IObject) {
	vcl.ShowMessage("查杀，查杀，我查。。。")
}

func (f *TForm1) OnEdit2Change(sender vcl.IObject) {
	val, _ := strconv.Atoi(f.Edit2.Text())
	if val >= 0 && val <= 100 {
		f.Gauge1.SetProgress(int32(val))
	}
}

func (f *TForm1) OnMemo1Click(vcl.IObject) {
	fmt.Println("OnMemo1Click")
}

func (f *TForm1) OnListView1Click(vcl.IObject) {
	idx := f.ListView1.ItemIndex()
	if idx != -1 {
		item := f.ListView1.Items().Item(idx)
		fmt.Println(item.Caption())
		fmt.Println("subitem:", item.SubItems().Text())
	}
}

//func (f *TForm1) OnApplicationHint(sender vcl.IObject) {
//	fmt.Println("提示了：", vcl.Application.Hint())
//}
