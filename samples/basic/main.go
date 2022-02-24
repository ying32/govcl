package main

import (
	"github.com/ying32/govcl/vcl"

	"fmt"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl/types"
)

type TMainForm struct {
	*vcl.TForm
	Button1 *vcl.TButton
}

type TForm1 struct {
	*vcl.TForm
	Button1 *vcl.TButton
}

var (
	mainForm *TMainForm
	form1    *TForm1
)

func main() {
	vcl.RunApp(&mainForm, &form1)
}

// --------------MainForm -----------------
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("Hello")
	f.EnabledMaximize(false)
	f.SetWidth(300)
	f.SetHeight(200)
	f.ScreenCenter()

	f.Button1 = vcl.NewButton(f)
	f.Button1.SetParent(f)
	f.Button1.SetCaption("窗口1")
	f.Button1.SetLeft(50)
	f.Button1.SetTop(50)
	f.Button1.SetOnClick(f.OnButton1Click)

	//x := vcl.NewFloatSpinEdit(f)
	//x.SetParent(f)
	//x.SetLeft(100)
	//x.SetTop(100)
	//x.SetMaxValue(10.0)
	//x.SetIncrement(0.2)
	//x.SetMinValue(0.1)
	//x.SetValue(3.0)
	//x.SetWidth(100)
	//
	//d := vcl.NewDirectoryEdit(f)
	//d.SetParent(f)
	//d.SetLeft(100)
	//d.SetTop(150)
	//d.SetWidth(200)
	//d.SetDirectory("C:\\xxxxx")

}

func (f *TMainForm) OnFormCloseQuery(Sender vcl.IObject, CanClose *bool) {
	*CanClose = vcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
}

func (f *TMainForm) OnButton1Click(object vcl.IObject) {
	form1.Show()
}

// ---------- Form1 ----------------

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	fmt.Println("onCreate")
	f.Button1 = vcl.NewButton(f)
	fmt.Println("f.Button1:", f.Button1.Instance())
	f.Button1.SetParent(f)
	//f.Button1.SetName("Button1")
	f.Button1.SetCaption("我是按钮")
	f.Button1.SetOnClick(f.OnButton1Click)
}

func (f *TForm1) OnButton1Click(object vcl.IObject) {
	vcl.ShowMessage("Click")
}
