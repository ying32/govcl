package main

import (
	"github.com/ying32/govcl/vcl"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(mainFormBytes, &MainForm)
	vcl.Application.CreateForm(loginFormBytes, &LoginForm)

	// 这里可以决定是不是显示主窗口
	vcl.Application.SetShowMainForm(false)
	// 这里显示出第二个窗口
	LoginForm.Show()

	vcl.Application.Run()

}
