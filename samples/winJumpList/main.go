// +build windows

package main

import (
	"fmt"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
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

	// 这玩意已经有问题了，原因不明，貌似是更新了操作系统后出现的，以前其他的App也有的，现在都没有了，估计是微软改了东西
	//categ := f.JumpList1.CustomCategories().Add2()
	//categ.SetCategoryName("分类1")
	//
	//item := categ.Items().Add2()
	//item.SetFriendlyName("notepad")
	//item.SetPath("notepad.exe")
	//item.SetArguments("C:\\Windows\\System32\\drivers\\etc\\hosts")
	//item.SetIcon("notepad.exe,0")
	//
	//item = categ.Items().Add2()
	//item.SetFriendlyName("calc")
	//item.SetPath("calc.exe")
	//item.SetIcon("calc.exe,0")
	//
	//task := f.JumpList1.TaskList().Add2()
	//task.SetFriendlyName("task1")
	//task.SetPath("notepad.exe")
	//item.SetIcon("notepad.exe,0")
	//
	//task = f.JumpList1.TaskList().Add2()
	//task.SetIsSeparator(true)
	//
	//task = f.JumpList1.TaskList().Add2()
	//task.SetFriendlyName("task2")
	//task.SetPath("calc.exe")
	//item.SetIcon("calc.exe,0")
}
