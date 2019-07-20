package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/tools"
)

func main() {
	// mac下发布时去掉，只在测试时使用
	tools.RunWithMacOSApp()

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(form1Bytes, &Form1)
	vcl.Application.CreateForm(form2Bytes, &Form2)

	vcl.Application.Run()

}
