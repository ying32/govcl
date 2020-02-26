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
	vcl.Application.CreateForm(&Form1)
	vcl.Application.CreateForm(&Form2)

	vcl.Application.Run()

}
