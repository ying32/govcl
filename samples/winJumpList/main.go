// +build windows

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
)

type TMainForm struct {
	*vcl.TForm
	JumpList1 *vcl.TJumpList
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
	f.JumpList1 = vcl.NewJumpList(f)

	f.JumpList1.SetEnabled(true)
	f.JumpList1.SetAutoRefresh(true)
	f.JumpList1.SetApplicationID("jumplistProject")

	categ := f.JumpList1.CustomCategories().Add2()
	categ.SetCategoryName("分类1")

	item := categ.Items().Add2()
	item.SetFriendlyName("notepad")
	item.SetPath("notepad.exe")
	item.SetArguments("C:\\Windows\\System32\\drivers\\etc\\hosts")
	item.SetIcon("notepad.exe,0")

	item = categ.Items().Add2()
	item.SetFriendlyName("calc")
	item.SetPath("calc.exe")
	item.SetIcon("calc.exe,0")

	task := f.JumpList1.TaskList().Add2()
	task.SetFriendlyName("task1")
	task.SetPath("notepad.exe")
	item.SetIcon("notepad.exe,0")

	task = f.JumpList1.TaskList().Add2()
	task.SetIsSeparator(true)

	task = f.JumpList1.TaskList().Add2()
	task.SetFriendlyName("task2")
	task.SetPath("calc.exe")
	item.SetIcon("calc.exe,0")
}
